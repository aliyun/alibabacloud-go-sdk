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

type AddAITemplateRequest struct {
	TemplateConfig *string `json:"TemplateConfig,omitempty" xml:"TemplateConfig,omitempty"`
	TemplateName   *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	TemplateType   *string `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
}

func (s AddAITemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s AddAITemplateRequest) GoString() string {
	return s.String()
}

func (s *AddAITemplateRequest) SetTemplateConfig(v string) *AddAITemplateRequest {
	s.TemplateConfig = &v
	return s
}

func (s *AddAITemplateRequest) SetTemplateName(v string) *AddAITemplateRequest {
	s.TemplateName = &v
	return s
}

func (s *AddAITemplateRequest) SetTemplateType(v string) *AddAITemplateRequest {
	s.TemplateType = &v
	return s
}

type AddAITemplateResponseBody struct {
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s AddAITemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddAITemplateResponseBody) GoString() string {
	return s.String()
}

func (s *AddAITemplateResponseBody) SetRequestId(v string) *AddAITemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddAITemplateResponseBody) SetTemplateId(v string) *AddAITemplateResponseBody {
	s.TemplateId = &v
	return s
}

type AddAITemplateResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddAITemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddAITemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s AddAITemplateResponse) GoString() string {
	return s.String()
}

func (s *AddAITemplateResponse) SetHeaders(v map[string]*string) *AddAITemplateResponse {
	s.Headers = v
	return s
}

func (s *AddAITemplateResponse) SetBody(v *AddAITemplateResponseBody) *AddAITemplateResponse {
	s.Body = v
	return s
}

type AddCategoryRequest struct {
	CateName *string `json:"CateName,omitempty" xml:"CateName,omitempty"`
	ParentId *int64  `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
	Type     *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s AddCategoryRequest) String() string {
	return tea.Prettify(s)
}

func (s AddCategoryRequest) GoString() string {
	return s.String()
}

func (s *AddCategoryRequest) SetCateName(v string) *AddCategoryRequest {
	s.CateName = &v
	return s
}

func (s *AddCategoryRequest) SetParentId(v int64) *AddCategoryRequest {
	s.ParentId = &v
	return s
}

func (s *AddCategoryRequest) SetType(v string) *AddCategoryRequest {
	s.Type = &v
	return s
}

type AddCategoryResponseBody struct {
	Category  *AddCategoryResponseBodyCategory `json:"Category,omitempty" xml:"Category,omitempty" type:"Struct"`
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddCategoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddCategoryResponseBody) GoString() string {
	return s.String()
}

func (s *AddCategoryResponseBody) SetCategory(v *AddCategoryResponseBodyCategory) *AddCategoryResponseBody {
	s.Category = v
	return s
}

func (s *AddCategoryResponseBody) SetRequestId(v string) *AddCategoryResponseBody {
	s.RequestId = &v
	return s
}

type AddCategoryResponseBodyCategory struct {
	CateId   *int64  `json:"CateId,omitempty" xml:"CateId,omitempty"`
	CateName *string `json:"CateName,omitempty" xml:"CateName,omitempty"`
	Level    *int64  `json:"Level,omitempty" xml:"Level,omitempty"`
	ParentId *int64  `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
	Type     *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s AddCategoryResponseBodyCategory) String() string {
	return tea.Prettify(s)
}

func (s AddCategoryResponseBodyCategory) GoString() string {
	return s.String()
}

func (s *AddCategoryResponseBodyCategory) SetCateId(v int64) *AddCategoryResponseBodyCategory {
	s.CateId = &v
	return s
}

func (s *AddCategoryResponseBodyCategory) SetCateName(v string) *AddCategoryResponseBodyCategory {
	s.CateName = &v
	return s
}

func (s *AddCategoryResponseBodyCategory) SetLevel(v int64) *AddCategoryResponseBodyCategory {
	s.Level = &v
	return s
}

func (s *AddCategoryResponseBodyCategory) SetParentId(v int64) *AddCategoryResponseBodyCategory {
	s.ParentId = &v
	return s
}

func (s *AddCategoryResponseBodyCategory) SetType(v string) *AddCategoryResponseBodyCategory {
	s.Type = &v
	return s
}

type AddCategoryResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddCategoryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddCategoryResponse) String() string {
	return tea.Prettify(s)
}

func (s AddCategoryResponse) GoString() string {
	return s.String()
}

func (s *AddCategoryResponse) SetHeaders(v map[string]*string) *AddCategoryResponse {
	s.Headers = v
	return s
}

func (s *AddCategoryResponse) SetBody(v *AddCategoryResponseBody) *AddCategoryResponse {
	s.Body = v
	return s
}

type AddEditingProjectRequest struct {
	CoverURL             *string `json:"CoverURL,omitempty" xml:"CoverURL,omitempty"`
	Description          *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Division             *string `json:"Division,omitempty" xml:"Division,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *string `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	Timeline             *string `json:"Timeline,omitempty" xml:"Timeline,omitempty"`
	Title                *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s AddEditingProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s AddEditingProjectRequest) GoString() string {
	return s.String()
}

func (s *AddEditingProjectRequest) SetCoverURL(v string) *AddEditingProjectRequest {
	s.CoverURL = &v
	return s
}

func (s *AddEditingProjectRequest) SetDescription(v string) *AddEditingProjectRequest {
	s.Description = &v
	return s
}

func (s *AddEditingProjectRequest) SetDivision(v string) *AddEditingProjectRequest {
	s.Division = &v
	return s
}

func (s *AddEditingProjectRequest) SetOwnerAccount(v string) *AddEditingProjectRequest {
	s.OwnerAccount = &v
	return s
}

func (s *AddEditingProjectRequest) SetOwnerId(v string) *AddEditingProjectRequest {
	s.OwnerId = &v
	return s
}

func (s *AddEditingProjectRequest) SetResourceOwnerAccount(v string) *AddEditingProjectRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AddEditingProjectRequest) SetResourceOwnerId(v string) *AddEditingProjectRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AddEditingProjectRequest) SetTimeline(v string) *AddEditingProjectRequest {
	s.Timeline = &v
	return s
}

func (s *AddEditingProjectRequest) SetTitle(v string) *AddEditingProjectRequest {
	s.Title = &v
	return s
}

type AddEditingProjectResponseBody struct {
	Project   *AddEditingProjectResponseBodyProject `json:"Project,omitempty" xml:"Project,omitempty" type:"Struct"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddEditingProjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddEditingProjectResponseBody) GoString() string {
	return s.String()
}

func (s *AddEditingProjectResponseBody) SetProject(v *AddEditingProjectResponseBodyProject) *AddEditingProjectResponseBody {
	s.Project = v
	return s
}

func (s *AddEditingProjectResponseBody) SetRequestId(v string) *AddEditingProjectResponseBody {
	s.RequestId = &v
	return s
}

type AddEditingProjectResponseBodyProject struct {
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	ProjectId    *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Title        *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s AddEditingProjectResponseBodyProject) String() string {
	return tea.Prettify(s)
}

func (s AddEditingProjectResponseBodyProject) GoString() string {
	return s.String()
}

func (s *AddEditingProjectResponseBodyProject) SetCreationTime(v string) *AddEditingProjectResponseBodyProject {
	s.CreationTime = &v
	return s
}

func (s *AddEditingProjectResponseBodyProject) SetDescription(v string) *AddEditingProjectResponseBodyProject {
	s.Description = &v
	return s
}

func (s *AddEditingProjectResponseBodyProject) SetModifiedTime(v string) *AddEditingProjectResponseBodyProject {
	s.ModifiedTime = &v
	return s
}

func (s *AddEditingProjectResponseBodyProject) SetProjectId(v string) *AddEditingProjectResponseBodyProject {
	s.ProjectId = &v
	return s
}

func (s *AddEditingProjectResponseBodyProject) SetStatus(v string) *AddEditingProjectResponseBodyProject {
	s.Status = &v
	return s
}

func (s *AddEditingProjectResponseBodyProject) SetTitle(v string) *AddEditingProjectResponseBodyProject {
	s.Title = &v
	return s
}

type AddEditingProjectResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddEditingProjectResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddEditingProjectResponse) String() string {
	return tea.Prettify(s)
}

func (s AddEditingProjectResponse) GoString() string {
	return s.String()
}

func (s *AddEditingProjectResponse) SetHeaders(v map[string]*string) *AddEditingProjectResponse {
	s.Headers = v
	return s
}

func (s *AddEditingProjectResponse) SetBody(v *AddEditingProjectResponseBody) *AddEditingProjectResponse {
	s.Body = v
	return s
}

type AddTranscodeTemplateGroupRequest struct {
	AppId                    *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	Name                     *string `json:"Name,omitempty" xml:"Name,omitempty"`
	TranscodeTemplateGroupId *string `json:"TranscodeTemplateGroupId,omitempty" xml:"TranscodeTemplateGroupId,omitempty"`
	TranscodeTemplateList    *string `json:"TranscodeTemplateList,omitempty" xml:"TranscodeTemplateList,omitempty"`
}

func (s AddTranscodeTemplateGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s AddTranscodeTemplateGroupRequest) GoString() string {
	return s.String()
}

func (s *AddTranscodeTemplateGroupRequest) SetAppId(v string) *AddTranscodeTemplateGroupRequest {
	s.AppId = &v
	return s
}

func (s *AddTranscodeTemplateGroupRequest) SetName(v string) *AddTranscodeTemplateGroupRequest {
	s.Name = &v
	return s
}

func (s *AddTranscodeTemplateGroupRequest) SetTranscodeTemplateGroupId(v string) *AddTranscodeTemplateGroupRequest {
	s.TranscodeTemplateGroupId = &v
	return s
}

func (s *AddTranscodeTemplateGroupRequest) SetTranscodeTemplateList(v string) *AddTranscodeTemplateGroupRequest {
	s.TranscodeTemplateList = &v
	return s
}

type AddTranscodeTemplateGroupResponseBody struct {
	RequestId                *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TranscodeTemplateGroupId *string `json:"TranscodeTemplateGroupId,omitempty" xml:"TranscodeTemplateGroupId,omitempty"`
}

func (s AddTranscodeTemplateGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddTranscodeTemplateGroupResponseBody) GoString() string {
	return s.String()
}

func (s *AddTranscodeTemplateGroupResponseBody) SetRequestId(v string) *AddTranscodeTemplateGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddTranscodeTemplateGroupResponseBody) SetTranscodeTemplateGroupId(v string) *AddTranscodeTemplateGroupResponseBody {
	s.TranscodeTemplateGroupId = &v
	return s
}

type AddTranscodeTemplateGroupResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddTranscodeTemplateGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddTranscodeTemplateGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s AddTranscodeTemplateGroupResponse) GoString() string {
	return s.String()
}

func (s *AddTranscodeTemplateGroupResponse) SetHeaders(v map[string]*string) *AddTranscodeTemplateGroupResponse {
	s.Headers = v
	return s
}

func (s *AddTranscodeTemplateGroupResponse) SetBody(v *AddTranscodeTemplateGroupResponseBody) *AddTranscodeTemplateGroupResponse {
	s.Body = v
	return s
}

type AddVodDomainRequest struct {
	CheckUrl       *string `json:"CheckUrl,omitempty" xml:"CheckUrl,omitempty"`
	DomainName     *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	OwnerAccount   *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId        *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	Scope          *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
	SecurityToken  *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	Sources        *string `json:"Sources,omitempty" xml:"Sources,omitempty"`
	TopLevelDomain *string `json:"TopLevelDomain,omitempty" xml:"TopLevelDomain,omitempty"`
}

func (s AddVodDomainRequest) String() string {
	return tea.Prettify(s)
}

func (s AddVodDomainRequest) GoString() string {
	return s.String()
}

func (s *AddVodDomainRequest) SetCheckUrl(v string) *AddVodDomainRequest {
	s.CheckUrl = &v
	return s
}

func (s *AddVodDomainRequest) SetDomainName(v string) *AddVodDomainRequest {
	s.DomainName = &v
	return s
}

func (s *AddVodDomainRequest) SetOwnerAccount(v string) *AddVodDomainRequest {
	s.OwnerAccount = &v
	return s
}

func (s *AddVodDomainRequest) SetOwnerId(v int64) *AddVodDomainRequest {
	s.OwnerId = &v
	return s
}

func (s *AddVodDomainRequest) SetScope(v string) *AddVodDomainRequest {
	s.Scope = &v
	return s
}

func (s *AddVodDomainRequest) SetSecurityToken(v string) *AddVodDomainRequest {
	s.SecurityToken = &v
	return s
}

func (s *AddVodDomainRequest) SetSources(v string) *AddVodDomainRequest {
	s.Sources = &v
	return s
}

func (s *AddVodDomainRequest) SetTopLevelDomain(v string) *AddVodDomainRequest {
	s.TopLevelDomain = &v
	return s
}

type AddVodDomainResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddVodDomainResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddVodDomainResponseBody) GoString() string {
	return s.String()
}

func (s *AddVodDomainResponseBody) SetRequestId(v string) *AddVodDomainResponseBody {
	s.RequestId = &v
	return s
}

type AddVodDomainResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddVodDomainResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddVodDomainResponse) String() string {
	return tea.Prettify(s)
}

func (s AddVodDomainResponse) GoString() string {
	return s.String()
}

func (s *AddVodDomainResponse) SetHeaders(v map[string]*string) *AddVodDomainResponse {
	s.Headers = v
	return s
}

func (s *AddVodDomainResponse) SetBody(v *AddVodDomainResponseBody) *AddVodDomainResponse {
	s.Body = v
	return s
}

type AddVodTemplateRequest struct {
	AppId          *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	Name           *string `json:"Name,omitempty" xml:"Name,omitempty"`
	TemplateConfig *string `json:"TemplateConfig,omitempty" xml:"TemplateConfig,omitempty"`
	TemplateType   *string `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
}

func (s AddVodTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s AddVodTemplateRequest) GoString() string {
	return s.String()
}

func (s *AddVodTemplateRequest) SetAppId(v string) *AddVodTemplateRequest {
	s.AppId = &v
	return s
}

func (s *AddVodTemplateRequest) SetName(v string) *AddVodTemplateRequest {
	s.Name = &v
	return s
}

func (s *AddVodTemplateRequest) SetTemplateConfig(v string) *AddVodTemplateRequest {
	s.TemplateConfig = &v
	return s
}

func (s *AddVodTemplateRequest) SetTemplateType(v string) *AddVodTemplateRequest {
	s.TemplateType = &v
	return s
}

type AddVodTemplateResponseBody struct {
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	VodTemplateId *string `json:"VodTemplateId,omitempty" xml:"VodTemplateId,omitempty"`
}

func (s AddVodTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddVodTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *AddVodTemplateResponseBody) SetRequestId(v string) *AddVodTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddVodTemplateResponseBody) SetVodTemplateId(v string) *AddVodTemplateResponseBody {
	s.VodTemplateId = &v
	return s
}

type AddVodTemplateResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddVodTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddVodTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s AddVodTemplateResponse) GoString() string {
	return s.String()
}

func (s *AddVodTemplateResponse) SetHeaders(v map[string]*string) *AddVodTemplateResponse {
	s.Headers = v
	return s
}

func (s *AddVodTemplateResponse) SetBody(v *AddVodTemplateResponseBody) *AddVodTemplateResponse {
	s.Body = v
	return s
}

type AddWatermarkRequest struct {
	AppId           *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	FileUrl         *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	Name            *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Type            *string `json:"Type,omitempty" xml:"Type,omitempty"`
	WatermarkConfig *string `json:"WatermarkConfig,omitempty" xml:"WatermarkConfig,omitempty"`
}

func (s AddWatermarkRequest) String() string {
	return tea.Prettify(s)
}

func (s AddWatermarkRequest) GoString() string {
	return s.String()
}

func (s *AddWatermarkRequest) SetAppId(v string) *AddWatermarkRequest {
	s.AppId = &v
	return s
}

func (s *AddWatermarkRequest) SetFileUrl(v string) *AddWatermarkRequest {
	s.FileUrl = &v
	return s
}

func (s *AddWatermarkRequest) SetName(v string) *AddWatermarkRequest {
	s.Name = &v
	return s
}

func (s *AddWatermarkRequest) SetType(v string) *AddWatermarkRequest {
	s.Type = &v
	return s
}

func (s *AddWatermarkRequest) SetWatermarkConfig(v string) *AddWatermarkRequest {
	s.WatermarkConfig = &v
	return s
}

type AddWatermarkResponseBody struct {
	RequestId     *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	WatermarkInfo *AddWatermarkResponseBodyWatermarkInfo `json:"WatermarkInfo,omitempty" xml:"WatermarkInfo,omitempty" type:"Struct"`
}

func (s AddWatermarkResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddWatermarkResponseBody) GoString() string {
	return s.String()
}

func (s *AddWatermarkResponseBody) SetRequestId(v string) *AddWatermarkResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddWatermarkResponseBody) SetWatermarkInfo(v *AddWatermarkResponseBodyWatermarkInfo) *AddWatermarkResponseBody {
	s.WatermarkInfo = v
	return s
}

type AddWatermarkResponseBodyWatermarkInfo struct {
	CreationTime    *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	FileUrl         *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	IsDefault       *string `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	Name            *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Type            *string `json:"Type,omitempty" xml:"Type,omitempty"`
	WatermarkConfig *string `json:"WatermarkConfig,omitempty" xml:"WatermarkConfig,omitempty"`
	WatermarkId     *string `json:"WatermarkId,omitempty" xml:"WatermarkId,omitempty"`
}

func (s AddWatermarkResponseBodyWatermarkInfo) String() string {
	return tea.Prettify(s)
}

func (s AddWatermarkResponseBodyWatermarkInfo) GoString() string {
	return s.String()
}

func (s *AddWatermarkResponseBodyWatermarkInfo) SetCreationTime(v string) *AddWatermarkResponseBodyWatermarkInfo {
	s.CreationTime = &v
	return s
}

func (s *AddWatermarkResponseBodyWatermarkInfo) SetFileUrl(v string) *AddWatermarkResponseBodyWatermarkInfo {
	s.FileUrl = &v
	return s
}

func (s *AddWatermarkResponseBodyWatermarkInfo) SetIsDefault(v string) *AddWatermarkResponseBodyWatermarkInfo {
	s.IsDefault = &v
	return s
}

func (s *AddWatermarkResponseBodyWatermarkInfo) SetName(v string) *AddWatermarkResponseBodyWatermarkInfo {
	s.Name = &v
	return s
}

func (s *AddWatermarkResponseBodyWatermarkInfo) SetType(v string) *AddWatermarkResponseBodyWatermarkInfo {
	s.Type = &v
	return s
}

func (s *AddWatermarkResponseBodyWatermarkInfo) SetWatermarkConfig(v string) *AddWatermarkResponseBodyWatermarkInfo {
	s.WatermarkConfig = &v
	return s
}

func (s *AddWatermarkResponseBodyWatermarkInfo) SetWatermarkId(v string) *AddWatermarkResponseBodyWatermarkInfo {
	s.WatermarkId = &v
	return s
}

type AddWatermarkResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddWatermarkResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddWatermarkResponse) String() string {
	return tea.Prettify(s)
}

func (s AddWatermarkResponse) GoString() string {
	return s.String()
}

func (s *AddWatermarkResponse) SetHeaders(v map[string]*string) *AddWatermarkResponse {
	s.Headers = v
	return s
}

func (s *AddWatermarkResponse) SetBody(v *AddWatermarkResponseBody) *AddWatermarkResponse {
	s.Body = v
	return s
}

type AttachAppPolicyToIdentityRequest struct {
	AppId        *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	IdentityName *string `json:"IdentityName,omitempty" xml:"IdentityName,omitempty"`
	IdentityType *string `json:"IdentityType,omitempty" xml:"IdentityType,omitempty"`
	PolicyNames  *string `json:"PolicyNames,omitempty" xml:"PolicyNames,omitempty"`
}

func (s AttachAppPolicyToIdentityRequest) String() string {
	return tea.Prettify(s)
}

func (s AttachAppPolicyToIdentityRequest) GoString() string {
	return s.String()
}

func (s *AttachAppPolicyToIdentityRequest) SetAppId(v string) *AttachAppPolicyToIdentityRequest {
	s.AppId = &v
	return s
}

func (s *AttachAppPolicyToIdentityRequest) SetIdentityName(v string) *AttachAppPolicyToIdentityRequest {
	s.IdentityName = &v
	return s
}

func (s *AttachAppPolicyToIdentityRequest) SetIdentityType(v string) *AttachAppPolicyToIdentityRequest {
	s.IdentityType = &v
	return s
}

func (s *AttachAppPolicyToIdentityRequest) SetPolicyNames(v string) *AttachAppPolicyToIdentityRequest {
	s.PolicyNames = &v
	return s
}

type AttachAppPolicyToIdentityResponseBody struct {
	FailedPolicyNames   []*string `json:"FailedPolicyNames,omitempty" xml:"FailedPolicyNames,omitempty" type:"Repeated"`
	NonExistPolicyNames []*string `json:"NonExistPolicyNames,omitempty" xml:"NonExistPolicyNames,omitempty" type:"Repeated"`
	RequestId           *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AttachAppPolicyToIdentityResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AttachAppPolicyToIdentityResponseBody) GoString() string {
	return s.String()
}

func (s *AttachAppPolicyToIdentityResponseBody) SetFailedPolicyNames(v []*string) *AttachAppPolicyToIdentityResponseBody {
	s.FailedPolicyNames = v
	return s
}

func (s *AttachAppPolicyToIdentityResponseBody) SetNonExistPolicyNames(v []*string) *AttachAppPolicyToIdentityResponseBody {
	s.NonExistPolicyNames = v
	return s
}

func (s *AttachAppPolicyToIdentityResponseBody) SetRequestId(v string) *AttachAppPolicyToIdentityResponseBody {
	s.RequestId = &v
	return s
}

type AttachAppPolicyToIdentityResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AttachAppPolicyToIdentityResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AttachAppPolicyToIdentityResponse) String() string {
	return tea.Prettify(s)
}

func (s AttachAppPolicyToIdentityResponse) GoString() string {
	return s.String()
}

func (s *AttachAppPolicyToIdentityResponse) SetHeaders(v map[string]*string) *AttachAppPolicyToIdentityResponse {
	s.Headers = v
	return s
}

func (s *AttachAppPolicyToIdentityResponse) SetBody(v *AttachAppPolicyToIdentityResponseBody) *AttachAppPolicyToIdentityResponse {
	s.Body = v
	return s
}

type BatchSetVodDomainConfigsRequest struct {
	DomainNames   *string `json:"DomainNames,omitempty" xml:"DomainNames,omitempty"`
	Functions     *string `json:"Functions,omitempty" xml:"Functions,omitempty"`
	OwnerAccount  *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s BatchSetVodDomainConfigsRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchSetVodDomainConfigsRequest) GoString() string {
	return s.String()
}

func (s *BatchSetVodDomainConfigsRequest) SetDomainNames(v string) *BatchSetVodDomainConfigsRequest {
	s.DomainNames = &v
	return s
}

func (s *BatchSetVodDomainConfigsRequest) SetFunctions(v string) *BatchSetVodDomainConfigsRequest {
	s.Functions = &v
	return s
}

func (s *BatchSetVodDomainConfigsRequest) SetOwnerAccount(v string) *BatchSetVodDomainConfigsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *BatchSetVodDomainConfigsRequest) SetOwnerId(v int64) *BatchSetVodDomainConfigsRequest {
	s.OwnerId = &v
	return s
}

func (s *BatchSetVodDomainConfigsRequest) SetSecurityToken(v string) *BatchSetVodDomainConfigsRequest {
	s.SecurityToken = &v
	return s
}

type BatchSetVodDomainConfigsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BatchSetVodDomainConfigsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchSetVodDomainConfigsResponseBody) GoString() string {
	return s.String()
}

func (s *BatchSetVodDomainConfigsResponseBody) SetRequestId(v string) *BatchSetVodDomainConfigsResponseBody {
	s.RequestId = &v
	return s
}

type BatchSetVodDomainConfigsResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *BatchSetVodDomainConfigsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BatchSetVodDomainConfigsResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchSetVodDomainConfigsResponse) GoString() string {
	return s.String()
}

func (s *BatchSetVodDomainConfigsResponse) SetHeaders(v map[string]*string) *BatchSetVodDomainConfigsResponse {
	s.Headers = v
	return s
}

func (s *BatchSetVodDomainConfigsResponse) SetBody(v *BatchSetVodDomainConfigsResponseBody) *BatchSetVodDomainConfigsResponse {
	s.Body = v
	return s
}

type BatchStartVodDomainRequest struct {
	DomainNames   *string `json:"DomainNames,omitempty" xml:"DomainNames,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s BatchStartVodDomainRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchStartVodDomainRequest) GoString() string {
	return s.String()
}

func (s *BatchStartVodDomainRequest) SetDomainNames(v string) *BatchStartVodDomainRequest {
	s.DomainNames = &v
	return s
}

func (s *BatchStartVodDomainRequest) SetOwnerId(v int64) *BatchStartVodDomainRequest {
	s.OwnerId = &v
	return s
}

func (s *BatchStartVodDomainRequest) SetSecurityToken(v string) *BatchStartVodDomainRequest {
	s.SecurityToken = &v
	return s
}

type BatchStartVodDomainResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BatchStartVodDomainResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchStartVodDomainResponseBody) GoString() string {
	return s.String()
}

func (s *BatchStartVodDomainResponseBody) SetRequestId(v string) *BatchStartVodDomainResponseBody {
	s.RequestId = &v
	return s
}

type BatchStartVodDomainResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *BatchStartVodDomainResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BatchStartVodDomainResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchStartVodDomainResponse) GoString() string {
	return s.String()
}

func (s *BatchStartVodDomainResponse) SetHeaders(v map[string]*string) *BatchStartVodDomainResponse {
	s.Headers = v
	return s
}

func (s *BatchStartVodDomainResponse) SetBody(v *BatchStartVodDomainResponseBody) *BatchStartVodDomainResponse {
	s.Body = v
	return s
}

type BatchStopVodDomainRequest struct {
	DomainNames   *string `json:"DomainNames,omitempty" xml:"DomainNames,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s BatchStopVodDomainRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchStopVodDomainRequest) GoString() string {
	return s.String()
}

func (s *BatchStopVodDomainRequest) SetDomainNames(v string) *BatchStopVodDomainRequest {
	s.DomainNames = &v
	return s
}

func (s *BatchStopVodDomainRequest) SetOwnerId(v int64) *BatchStopVodDomainRequest {
	s.OwnerId = &v
	return s
}

func (s *BatchStopVodDomainRequest) SetSecurityToken(v string) *BatchStopVodDomainRequest {
	s.SecurityToken = &v
	return s
}

type BatchStopVodDomainResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BatchStopVodDomainResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchStopVodDomainResponseBody) GoString() string {
	return s.String()
}

func (s *BatchStopVodDomainResponseBody) SetRequestId(v string) *BatchStopVodDomainResponseBody {
	s.RequestId = &v
	return s
}

type BatchStopVodDomainResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *BatchStopVodDomainResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BatchStopVodDomainResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchStopVodDomainResponse) GoString() string {
	return s.String()
}

func (s *BatchStopVodDomainResponse) SetHeaders(v map[string]*string) *BatchStopVodDomainResponse {
	s.Headers = v
	return s
}

func (s *BatchStopVodDomainResponse) SetBody(v *BatchStopVodDomainResponseBody) *BatchStopVodDomainResponse {
	s.Body = v
	return s
}

type CancelUrlUploadJobsRequest struct {
	JobIds     *string `json:"JobIds,omitempty" xml:"JobIds,omitempty"`
	UploadUrls *string `json:"UploadUrls,omitempty" xml:"UploadUrls,omitempty"`
}

func (s CancelUrlUploadJobsRequest) String() string {
	return tea.Prettify(s)
}

func (s CancelUrlUploadJobsRequest) GoString() string {
	return s.String()
}

func (s *CancelUrlUploadJobsRequest) SetJobIds(v string) *CancelUrlUploadJobsRequest {
	s.JobIds = &v
	return s
}

func (s *CancelUrlUploadJobsRequest) SetUploadUrls(v string) *CancelUrlUploadJobsRequest {
	s.UploadUrls = &v
	return s
}

type CancelUrlUploadJobsResponseBody struct {
	CanceledJobs []*string `json:"CanceledJobs,omitempty" xml:"CanceledJobs,omitempty" type:"Repeated"`
	NonExists    []*string `json:"NonExists,omitempty" xml:"NonExists,omitempty" type:"Repeated"`
	RequestId    *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CancelUrlUploadJobsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CancelUrlUploadJobsResponseBody) GoString() string {
	return s.String()
}

func (s *CancelUrlUploadJobsResponseBody) SetCanceledJobs(v []*string) *CancelUrlUploadJobsResponseBody {
	s.CanceledJobs = v
	return s
}

func (s *CancelUrlUploadJobsResponseBody) SetNonExists(v []*string) *CancelUrlUploadJobsResponseBody {
	s.NonExists = v
	return s
}

func (s *CancelUrlUploadJobsResponseBody) SetRequestId(v string) *CancelUrlUploadJobsResponseBody {
	s.RequestId = &v
	return s
}

type CancelUrlUploadJobsResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CancelUrlUploadJobsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CancelUrlUploadJobsResponse) String() string {
	return tea.Prettify(s)
}

func (s CancelUrlUploadJobsResponse) GoString() string {
	return s.String()
}

func (s *CancelUrlUploadJobsResponse) SetHeaders(v map[string]*string) *CancelUrlUploadJobsResponse {
	s.Headers = v
	return s
}

func (s *CancelUrlUploadJobsResponse) SetBody(v *CancelUrlUploadJobsResponseBody) *CancelUrlUploadJobsResponse {
	s.Body = v
	return s
}

type CreateAppInfoRequest struct {
	AppName     *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
}

func (s CreateAppInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAppInfoRequest) GoString() string {
	return s.String()
}

func (s *CreateAppInfoRequest) SetAppName(v string) *CreateAppInfoRequest {
	s.AppName = &v
	return s
}

func (s *CreateAppInfoRequest) SetDescription(v string) *CreateAppInfoRequest {
	s.Description = &v
	return s
}

type CreateAppInfoResponseBody struct {
	AppId     *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateAppInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAppInfoResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAppInfoResponseBody) SetAppId(v string) *CreateAppInfoResponseBody {
	s.AppId = &v
	return s
}

func (s *CreateAppInfoResponseBody) SetRequestId(v string) *CreateAppInfoResponseBody {
	s.RequestId = &v
	return s
}

type CreateAppInfoResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateAppInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateAppInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAppInfoResponse) GoString() string {
	return s.String()
}

func (s *CreateAppInfoResponse) SetHeaders(v map[string]*string) *CreateAppInfoResponse {
	s.Headers = v
	return s
}

func (s *CreateAppInfoResponse) SetBody(v *CreateAppInfoResponseBody) *CreateAppInfoResponse {
	s.Body = v
	return s
}

type CreateAuditRequest struct {
	AuditContent *string `json:"AuditContent,omitempty" xml:"AuditContent,omitempty"`
}

func (s CreateAuditRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAuditRequest) GoString() string {
	return s.String()
}

func (s *CreateAuditRequest) SetAuditContent(v string) *CreateAuditRequest {
	s.AuditContent = &v
	return s
}

type CreateAuditResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateAuditResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAuditResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAuditResponseBody) SetRequestId(v string) *CreateAuditResponseBody {
	s.RequestId = &v
	return s
}

type CreateAuditResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateAuditResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateAuditResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAuditResponse) GoString() string {
	return s.String()
}

func (s *CreateAuditResponse) SetHeaders(v map[string]*string) *CreateAuditResponse {
	s.Headers = v
	return s
}

func (s *CreateAuditResponse) SetBody(v *CreateAuditResponseBody) *CreateAuditResponse {
	s.Body = v
	return s
}

type CreateDetectionTemplateRequest struct {
	Period       *string `json:"Period,omitempty" xml:"Period,omitempty"`
	Platform     *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
}

func (s CreateDetectionTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDetectionTemplateRequest) GoString() string {
	return s.String()
}

func (s *CreateDetectionTemplateRequest) SetPeriod(v string) *CreateDetectionTemplateRequest {
	s.Period = &v
	return s
}

func (s *CreateDetectionTemplateRequest) SetPlatform(v string) *CreateDetectionTemplateRequest {
	s.Platform = &v
	return s
}

func (s *CreateDetectionTemplateRequest) SetTemplateName(v string) *CreateDetectionTemplateRequest {
	s.TemplateName = &v
	return s
}

type CreateDetectionTemplateResponseBody struct {
	DetectionTemplate *CreateDetectionTemplateResponseBodyDetectionTemplate `json:"DetectionTemplate,omitempty" xml:"DetectionTemplate,omitempty" type:"Struct"`
	RequestId         *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDetectionTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDetectionTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDetectionTemplateResponseBody) SetDetectionTemplate(v *CreateDetectionTemplateResponseBodyDetectionTemplate) *CreateDetectionTemplateResponseBody {
	s.DetectionTemplate = v
	return s
}

func (s *CreateDetectionTemplateResponseBody) SetRequestId(v string) *CreateDetectionTemplateResponseBody {
	s.RequestId = &v
	return s
}

type CreateDetectionTemplateResponseBodyDetectionTemplate struct {
	CreateTime   *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	ModifyTime   *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	Period       *string `json:"Period,omitempty" xml:"Period,omitempty"`
	Platform     *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	TemplateId   *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	UserId       *int64  `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s CreateDetectionTemplateResponseBodyDetectionTemplate) String() string {
	return tea.Prettify(s)
}

func (s CreateDetectionTemplateResponseBodyDetectionTemplate) GoString() string {
	return s.String()
}

func (s *CreateDetectionTemplateResponseBodyDetectionTemplate) SetCreateTime(v string) *CreateDetectionTemplateResponseBodyDetectionTemplate {
	s.CreateTime = &v
	return s
}

func (s *CreateDetectionTemplateResponseBodyDetectionTemplate) SetModifyTime(v string) *CreateDetectionTemplateResponseBodyDetectionTemplate {
	s.ModifyTime = &v
	return s
}

func (s *CreateDetectionTemplateResponseBodyDetectionTemplate) SetPeriod(v string) *CreateDetectionTemplateResponseBodyDetectionTemplate {
	s.Period = &v
	return s
}

func (s *CreateDetectionTemplateResponseBodyDetectionTemplate) SetPlatform(v string) *CreateDetectionTemplateResponseBodyDetectionTemplate {
	s.Platform = &v
	return s
}

func (s *CreateDetectionTemplateResponseBodyDetectionTemplate) SetTemplateId(v string) *CreateDetectionTemplateResponseBodyDetectionTemplate {
	s.TemplateId = &v
	return s
}

func (s *CreateDetectionTemplateResponseBodyDetectionTemplate) SetTemplateName(v string) *CreateDetectionTemplateResponseBodyDetectionTemplate {
	s.TemplateName = &v
	return s
}

func (s *CreateDetectionTemplateResponseBodyDetectionTemplate) SetUserId(v int64) *CreateDetectionTemplateResponseBodyDetectionTemplate {
	s.UserId = &v
	return s
}

type CreateDetectionTemplateResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateDetectionTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateDetectionTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDetectionTemplateResponse) GoString() string {
	return s.String()
}

func (s *CreateDetectionTemplateResponse) SetHeaders(v map[string]*string) *CreateDetectionTemplateResponse {
	s.Headers = v
	return s
}

func (s *CreateDetectionTemplateResponse) SetBody(v *CreateDetectionTemplateResponseBody) *CreateDetectionTemplateResponse {
	s.Body = v
	return s
}

type CreateUploadAttachedMediaRequest struct {
	AppId           *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	BusinessType    *string `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	CateIds         *string `json:"CateIds,omitempty" xml:"CateIds,omitempty"`
	Description     *string `json:"Description,omitempty" xml:"Description,omitempty"`
	FileName        *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	FileSize        *string `json:"FileSize,omitempty" xml:"FileSize,omitempty"`
	MediaExt        *string `json:"MediaExt,omitempty" xml:"MediaExt,omitempty"`
	StorageLocation *string `json:"StorageLocation,omitempty" xml:"StorageLocation,omitempty"`
	Tags            *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	Title           *string `json:"Title,omitempty" xml:"Title,omitempty"`
	UserData        *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s CreateUploadAttachedMediaRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateUploadAttachedMediaRequest) GoString() string {
	return s.String()
}

func (s *CreateUploadAttachedMediaRequest) SetAppId(v string) *CreateUploadAttachedMediaRequest {
	s.AppId = &v
	return s
}

func (s *CreateUploadAttachedMediaRequest) SetBusinessType(v string) *CreateUploadAttachedMediaRequest {
	s.BusinessType = &v
	return s
}

func (s *CreateUploadAttachedMediaRequest) SetCateIds(v string) *CreateUploadAttachedMediaRequest {
	s.CateIds = &v
	return s
}

func (s *CreateUploadAttachedMediaRequest) SetDescription(v string) *CreateUploadAttachedMediaRequest {
	s.Description = &v
	return s
}

func (s *CreateUploadAttachedMediaRequest) SetFileName(v string) *CreateUploadAttachedMediaRequest {
	s.FileName = &v
	return s
}

func (s *CreateUploadAttachedMediaRequest) SetFileSize(v string) *CreateUploadAttachedMediaRequest {
	s.FileSize = &v
	return s
}

func (s *CreateUploadAttachedMediaRequest) SetMediaExt(v string) *CreateUploadAttachedMediaRequest {
	s.MediaExt = &v
	return s
}

func (s *CreateUploadAttachedMediaRequest) SetStorageLocation(v string) *CreateUploadAttachedMediaRequest {
	s.StorageLocation = &v
	return s
}

func (s *CreateUploadAttachedMediaRequest) SetTags(v string) *CreateUploadAttachedMediaRequest {
	s.Tags = &v
	return s
}

func (s *CreateUploadAttachedMediaRequest) SetTitle(v string) *CreateUploadAttachedMediaRequest {
	s.Title = &v
	return s
}

func (s *CreateUploadAttachedMediaRequest) SetUserData(v string) *CreateUploadAttachedMediaRequest {
	s.UserData = &v
	return s
}

type CreateUploadAttachedMediaResponseBody struct {
	FileURL       *string `json:"FileURL,omitempty" xml:"FileURL,omitempty"`
	MediaId       *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	MediaURL      *string `json:"MediaURL,omitempty" xml:"MediaURL,omitempty"`
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	UploadAddress *string `json:"UploadAddress,omitempty" xml:"UploadAddress,omitempty"`
	UploadAuth    *string `json:"UploadAuth,omitempty" xml:"UploadAuth,omitempty"`
}

func (s CreateUploadAttachedMediaResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateUploadAttachedMediaResponseBody) GoString() string {
	return s.String()
}

func (s *CreateUploadAttachedMediaResponseBody) SetFileURL(v string) *CreateUploadAttachedMediaResponseBody {
	s.FileURL = &v
	return s
}

func (s *CreateUploadAttachedMediaResponseBody) SetMediaId(v string) *CreateUploadAttachedMediaResponseBody {
	s.MediaId = &v
	return s
}

func (s *CreateUploadAttachedMediaResponseBody) SetMediaURL(v string) *CreateUploadAttachedMediaResponseBody {
	s.MediaURL = &v
	return s
}

func (s *CreateUploadAttachedMediaResponseBody) SetRequestId(v string) *CreateUploadAttachedMediaResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateUploadAttachedMediaResponseBody) SetUploadAddress(v string) *CreateUploadAttachedMediaResponseBody {
	s.UploadAddress = &v
	return s
}

func (s *CreateUploadAttachedMediaResponseBody) SetUploadAuth(v string) *CreateUploadAttachedMediaResponseBody {
	s.UploadAuth = &v
	return s
}

type CreateUploadAttachedMediaResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateUploadAttachedMediaResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateUploadAttachedMediaResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateUploadAttachedMediaResponse) GoString() string {
	return s.String()
}

func (s *CreateUploadAttachedMediaResponse) SetHeaders(v map[string]*string) *CreateUploadAttachedMediaResponse {
	s.Headers = v
	return s
}

func (s *CreateUploadAttachedMediaResponse) SetBody(v *CreateUploadAttachedMediaResponseBody) *CreateUploadAttachedMediaResponse {
	s.Body = v
	return s
}

type CreateUploadImageRequest struct {
	AppId           *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	CateId          *int64  `json:"CateId,omitempty" xml:"CateId,omitempty"`
	Description     *string `json:"Description,omitempty" xml:"Description,omitempty"`
	ImageExt        *string `json:"ImageExt,omitempty" xml:"ImageExt,omitempty"`
	ImageType       *string `json:"ImageType,omitempty" xml:"ImageType,omitempty"`
	StorageLocation *string `json:"StorageLocation,omitempty" xml:"StorageLocation,omitempty"`
	Tags            *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	Title           *string `json:"Title,omitempty" xml:"Title,omitempty"`
	UserData        *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s CreateUploadImageRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateUploadImageRequest) GoString() string {
	return s.String()
}

func (s *CreateUploadImageRequest) SetAppId(v string) *CreateUploadImageRequest {
	s.AppId = &v
	return s
}

func (s *CreateUploadImageRequest) SetCateId(v int64) *CreateUploadImageRequest {
	s.CateId = &v
	return s
}

func (s *CreateUploadImageRequest) SetDescription(v string) *CreateUploadImageRequest {
	s.Description = &v
	return s
}

func (s *CreateUploadImageRequest) SetImageExt(v string) *CreateUploadImageRequest {
	s.ImageExt = &v
	return s
}

func (s *CreateUploadImageRequest) SetImageType(v string) *CreateUploadImageRequest {
	s.ImageType = &v
	return s
}

func (s *CreateUploadImageRequest) SetStorageLocation(v string) *CreateUploadImageRequest {
	s.StorageLocation = &v
	return s
}

func (s *CreateUploadImageRequest) SetTags(v string) *CreateUploadImageRequest {
	s.Tags = &v
	return s
}

func (s *CreateUploadImageRequest) SetTitle(v string) *CreateUploadImageRequest {
	s.Title = &v
	return s
}

func (s *CreateUploadImageRequest) SetUserData(v string) *CreateUploadImageRequest {
	s.UserData = &v
	return s
}

type CreateUploadImageResponseBody struct {
	FileURL       *string `json:"FileURL,omitempty" xml:"FileURL,omitempty"`
	ImageId       *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	ImageURL      *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	UploadAddress *string `json:"UploadAddress,omitempty" xml:"UploadAddress,omitempty"`
	UploadAuth    *string `json:"UploadAuth,omitempty" xml:"UploadAuth,omitempty"`
}

func (s CreateUploadImageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateUploadImageResponseBody) GoString() string {
	return s.String()
}

func (s *CreateUploadImageResponseBody) SetFileURL(v string) *CreateUploadImageResponseBody {
	s.FileURL = &v
	return s
}

func (s *CreateUploadImageResponseBody) SetImageId(v string) *CreateUploadImageResponseBody {
	s.ImageId = &v
	return s
}

func (s *CreateUploadImageResponseBody) SetImageURL(v string) *CreateUploadImageResponseBody {
	s.ImageURL = &v
	return s
}

func (s *CreateUploadImageResponseBody) SetRequestId(v string) *CreateUploadImageResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateUploadImageResponseBody) SetUploadAddress(v string) *CreateUploadImageResponseBody {
	s.UploadAddress = &v
	return s
}

func (s *CreateUploadImageResponseBody) SetUploadAuth(v string) *CreateUploadImageResponseBody {
	s.UploadAuth = &v
	return s
}

type CreateUploadImageResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateUploadImageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateUploadImageResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateUploadImageResponse) GoString() string {
	return s.String()
}

func (s *CreateUploadImageResponse) SetHeaders(v map[string]*string) *CreateUploadImageResponse {
	s.Headers = v
	return s
}

func (s *CreateUploadImageResponse) SetBody(v *CreateUploadImageResponseBody) *CreateUploadImageResponse {
	s.Body = v
	return s
}

type CreateUploadVideoRequest struct {
	AppId           *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	CateId          *int64  `json:"CateId,omitempty" xml:"CateId,omitempty"`
	CoverURL        *string `json:"CoverURL,omitempty" xml:"CoverURL,omitempty"`
	Description     *string `json:"Description,omitempty" xml:"Description,omitempty"`
	FileName        *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	FileSize        *int64  `json:"FileSize,omitempty" xml:"FileSize,omitempty"`
	StorageLocation *string `json:"StorageLocation,omitempty" xml:"StorageLocation,omitempty"`
	Tags            *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	TemplateGroupId *string `json:"TemplateGroupId,omitempty" xml:"TemplateGroupId,omitempty"`
	Title           *string `json:"Title,omitempty" xml:"Title,omitempty"`
	UserData        *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
	WorkflowId      *string `json:"WorkflowId,omitempty" xml:"WorkflowId,omitempty"`
}

func (s CreateUploadVideoRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateUploadVideoRequest) GoString() string {
	return s.String()
}

func (s *CreateUploadVideoRequest) SetAppId(v string) *CreateUploadVideoRequest {
	s.AppId = &v
	return s
}

func (s *CreateUploadVideoRequest) SetCateId(v int64) *CreateUploadVideoRequest {
	s.CateId = &v
	return s
}

func (s *CreateUploadVideoRequest) SetCoverURL(v string) *CreateUploadVideoRequest {
	s.CoverURL = &v
	return s
}

func (s *CreateUploadVideoRequest) SetDescription(v string) *CreateUploadVideoRequest {
	s.Description = &v
	return s
}

func (s *CreateUploadVideoRequest) SetFileName(v string) *CreateUploadVideoRequest {
	s.FileName = &v
	return s
}

func (s *CreateUploadVideoRequest) SetFileSize(v int64) *CreateUploadVideoRequest {
	s.FileSize = &v
	return s
}

func (s *CreateUploadVideoRequest) SetStorageLocation(v string) *CreateUploadVideoRequest {
	s.StorageLocation = &v
	return s
}

func (s *CreateUploadVideoRequest) SetTags(v string) *CreateUploadVideoRequest {
	s.Tags = &v
	return s
}

func (s *CreateUploadVideoRequest) SetTemplateGroupId(v string) *CreateUploadVideoRequest {
	s.TemplateGroupId = &v
	return s
}

func (s *CreateUploadVideoRequest) SetTitle(v string) *CreateUploadVideoRequest {
	s.Title = &v
	return s
}

func (s *CreateUploadVideoRequest) SetUserData(v string) *CreateUploadVideoRequest {
	s.UserData = &v
	return s
}

func (s *CreateUploadVideoRequest) SetWorkflowId(v string) *CreateUploadVideoRequest {
	s.WorkflowId = &v
	return s
}

type CreateUploadVideoResponseBody struct {
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	UploadAddress *string `json:"UploadAddress,omitempty" xml:"UploadAddress,omitempty"`
	UploadAuth    *string `json:"UploadAuth,omitempty" xml:"UploadAuth,omitempty"`
	VideoId       *string `json:"VideoId,omitempty" xml:"VideoId,omitempty"`
}

func (s CreateUploadVideoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateUploadVideoResponseBody) GoString() string {
	return s.String()
}

func (s *CreateUploadVideoResponseBody) SetRequestId(v string) *CreateUploadVideoResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateUploadVideoResponseBody) SetUploadAddress(v string) *CreateUploadVideoResponseBody {
	s.UploadAddress = &v
	return s
}

func (s *CreateUploadVideoResponseBody) SetUploadAuth(v string) *CreateUploadVideoResponseBody {
	s.UploadAuth = &v
	return s
}

func (s *CreateUploadVideoResponseBody) SetVideoId(v string) *CreateUploadVideoResponseBody {
	s.VideoId = &v
	return s
}

type CreateUploadVideoResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateUploadVideoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateUploadVideoResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateUploadVideoResponse) GoString() string {
	return s.String()
}

func (s *CreateUploadVideoResponse) SetHeaders(v map[string]*string) *CreateUploadVideoResponse {
	s.Headers = v
	return s
}

func (s *CreateUploadVideoResponse) SetBody(v *CreateUploadVideoResponseBody) *CreateUploadVideoResponse {
	s.Body = v
	return s
}

type CreateVodRealTimeLogDeliveryRequest struct {
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	Logstore   *string `json:"Logstore,omitempty" xml:"Logstore,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	Project    *string `json:"Project,omitempty" xml:"Project,omitempty"`
	Region     *string `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s CreateVodRealTimeLogDeliveryRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateVodRealTimeLogDeliveryRequest) GoString() string {
	return s.String()
}

func (s *CreateVodRealTimeLogDeliveryRequest) SetDomainName(v string) *CreateVodRealTimeLogDeliveryRequest {
	s.DomainName = &v
	return s
}

func (s *CreateVodRealTimeLogDeliveryRequest) SetLogstore(v string) *CreateVodRealTimeLogDeliveryRequest {
	s.Logstore = &v
	return s
}

func (s *CreateVodRealTimeLogDeliveryRequest) SetOwnerId(v int64) *CreateVodRealTimeLogDeliveryRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateVodRealTimeLogDeliveryRequest) SetProject(v string) *CreateVodRealTimeLogDeliveryRequest {
	s.Project = &v
	return s
}

func (s *CreateVodRealTimeLogDeliveryRequest) SetRegion(v string) *CreateVodRealTimeLogDeliveryRequest {
	s.Region = &v
	return s
}

type CreateVodRealTimeLogDeliveryResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateVodRealTimeLogDeliveryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateVodRealTimeLogDeliveryResponseBody) GoString() string {
	return s.String()
}

func (s *CreateVodRealTimeLogDeliveryResponseBody) SetRequestId(v string) *CreateVodRealTimeLogDeliveryResponseBody {
	s.RequestId = &v
	return s
}

type CreateVodRealTimeLogDeliveryResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateVodRealTimeLogDeliveryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateVodRealTimeLogDeliveryResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateVodRealTimeLogDeliveryResponse) GoString() string {
	return s.String()
}

func (s *CreateVodRealTimeLogDeliveryResponse) SetHeaders(v map[string]*string) *CreateVodRealTimeLogDeliveryResponse {
	s.Headers = v
	return s
}

func (s *CreateVodRealTimeLogDeliveryResponse) SetBody(v *CreateVodRealTimeLogDeliveryResponseBody) *CreateVodRealTimeLogDeliveryResponse {
	s.Body = v
	return s
}

type DeleteAIImageInfosRequest struct {
	AIImageInfoIds *string `json:"AIImageInfoIds,omitempty" xml:"AIImageInfoIds,omitempty"`
}

func (s DeleteAIImageInfosRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteAIImageInfosRequest) GoString() string {
	return s.String()
}

func (s *DeleteAIImageInfosRequest) SetAIImageInfoIds(v string) *DeleteAIImageInfosRequest {
	s.AIImageInfoIds = &v
	return s
}

type DeleteAIImageInfosResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteAIImageInfosResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteAIImageInfosResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAIImageInfosResponseBody) SetRequestId(v string) *DeleteAIImageInfosResponseBody {
	s.RequestId = &v
	return s
}

type DeleteAIImageInfosResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteAIImageInfosResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteAIImageInfosResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteAIImageInfosResponse) GoString() string {
	return s.String()
}

func (s *DeleteAIImageInfosResponse) SetHeaders(v map[string]*string) *DeleteAIImageInfosResponse {
	s.Headers = v
	return s
}

func (s *DeleteAIImageInfosResponse) SetBody(v *DeleteAIImageInfosResponseBody) *DeleteAIImageInfosResponse {
	s.Body = v
	return s
}

type DeleteAITemplateRequest struct {
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s DeleteAITemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteAITemplateRequest) GoString() string {
	return s.String()
}

func (s *DeleteAITemplateRequest) SetTemplateId(v string) *DeleteAITemplateRequest {
	s.TemplateId = &v
	return s
}

type DeleteAITemplateResponseBody struct {
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s DeleteAITemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteAITemplateResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAITemplateResponseBody) SetRequestId(v string) *DeleteAITemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAITemplateResponseBody) SetTemplateId(v string) *DeleteAITemplateResponseBody {
	s.TemplateId = &v
	return s
}

type DeleteAITemplateResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteAITemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteAITemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteAITemplateResponse) GoString() string {
	return s.String()
}

func (s *DeleteAITemplateResponse) SetHeaders(v map[string]*string) *DeleteAITemplateResponse {
	s.Headers = v
	return s
}

func (s *DeleteAITemplateResponse) SetBody(v *DeleteAITemplateResponseBody) *DeleteAITemplateResponse {
	s.Body = v
	return s
}

type DeleteAppInfoRequest struct {
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
}

func (s DeleteAppInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteAppInfoRequest) GoString() string {
	return s.String()
}

func (s *DeleteAppInfoRequest) SetAppId(v string) *DeleteAppInfoRequest {
	s.AppId = &v
	return s
}

type DeleteAppInfoResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteAppInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteAppInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAppInfoResponseBody) SetRequestId(v string) *DeleteAppInfoResponseBody {
	s.RequestId = &v
	return s
}

type DeleteAppInfoResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteAppInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteAppInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteAppInfoResponse) GoString() string {
	return s.String()
}

func (s *DeleteAppInfoResponse) SetHeaders(v map[string]*string) *DeleteAppInfoResponse {
	s.Headers = v
	return s
}

func (s *DeleteAppInfoResponse) SetBody(v *DeleteAppInfoResponseBody) *DeleteAppInfoResponse {
	s.Body = v
	return s
}

type DeleteAttachedMediaRequest struct {
	MediaIds *string `json:"MediaIds,omitempty" xml:"MediaIds,omitempty"`
}

func (s DeleteAttachedMediaRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteAttachedMediaRequest) GoString() string {
	return s.String()
}

func (s *DeleteAttachedMediaRequest) SetMediaIds(v string) *DeleteAttachedMediaRequest {
	s.MediaIds = &v
	return s
}

type DeleteAttachedMediaResponseBody struct {
	NonExistMediaIds []*string `json:"NonExistMediaIds,omitempty" xml:"NonExistMediaIds,omitempty" type:"Repeated"`
	RequestId        *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteAttachedMediaResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteAttachedMediaResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAttachedMediaResponseBody) SetNonExistMediaIds(v []*string) *DeleteAttachedMediaResponseBody {
	s.NonExistMediaIds = v
	return s
}

func (s *DeleteAttachedMediaResponseBody) SetRequestId(v string) *DeleteAttachedMediaResponseBody {
	s.RequestId = &v
	return s
}

type DeleteAttachedMediaResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteAttachedMediaResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteAttachedMediaResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteAttachedMediaResponse) GoString() string {
	return s.String()
}

func (s *DeleteAttachedMediaResponse) SetHeaders(v map[string]*string) *DeleteAttachedMediaResponse {
	s.Headers = v
	return s
}

func (s *DeleteAttachedMediaResponse) SetBody(v *DeleteAttachedMediaResponseBody) *DeleteAttachedMediaResponse {
	s.Body = v
	return s
}

type DeleteCategoryRequest struct {
	CateId *int64 `json:"CateId,omitempty" xml:"CateId,omitempty"`
}

func (s DeleteCategoryRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteCategoryRequest) GoString() string {
	return s.String()
}

func (s *DeleteCategoryRequest) SetCateId(v int64) *DeleteCategoryRequest {
	s.CateId = &v
	return s
}

type DeleteCategoryResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteCategoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteCategoryResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCategoryResponseBody) SetRequestId(v string) *DeleteCategoryResponseBody {
	s.RequestId = &v
	return s
}

type DeleteCategoryResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteCategoryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteCategoryResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteCategoryResponse) GoString() string {
	return s.String()
}

func (s *DeleteCategoryResponse) SetHeaders(v map[string]*string) *DeleteCategoryResponse {
	s.Headers = v
	return s
}

func (s *DeleteCategoryResponse) SetBody(v *DeleteCategoryResponseBody) *DeleteCategoryResponse {
	s.Body = v
	return s
}

type DeleteDetectionTemplateRequest struct {
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s DeleteDetectionTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDetectionTemplateRequest) GoString() string {
	return s.String()
}

func (s *DeleteDetectionTemplateRequest) SetTemplateId(v string) *DeleteDetectionTemplateRequest {
	s.TemplateId = &v
	return s
}

type DeleteDetectionTemplateResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDetectionTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDetectionTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDetectionTemplateResponseBody) SetRequestId(v string) *DeleteDetectionTemplateResponseBody {
	s.RequestId = &v
	return s
}

type DeleteDetectionTemplateResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteDetectionTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteDetectionTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDetectionTemplateResponse) GoString() string {
	return s.String()
}

func (s *DeleteDetectionTemplateResponse) SetHeaders(v map[string]*string) *DeleteDetectionTemplateResponse {
	s.Headers = v
	return s
}

func (s *DeleteDetectionTemplateResponse) SetBody(v *DeleteDetectionTemplateResponseBody) *DeleteDetectionTemplateResponse {
	s.Body = v
	return s
}

type DeleteDynamicImageRequest struct {
	DynamicImageIds *string `json:"DynamicImageIds,omitempty" xml:"DynamicImageIds,omitempty"`
	VideoId         *string `json:"VideoId,omitempty" xml:"VideoId,omitempty"`
}

func (s DeleteDynamicImageRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDynamicImageRequest) GoString() string {
	return s.String()
}

func (s *DeleteDynamicImageRequest) SetDynamicImageIds(v string) *DeleteDynamicImageRequest {
	s.DynamicImageIds = &v
	return s
}

func (s *DeleteDynamicImageRequest) SetVideoId(v string) *DeleteDynamicImageRequest {
	s.VideoId = &v
	return s
}

type DeleteDynamicImageResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDynamicImageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDynamicImageResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDynamicImageResponseBody) SetRequestId(v string) *DeleteDynamicImageResponseBody {
	s.RequestId = &v
	return s
}

type DeleteDynamicImageResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteDynamicImageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteDynamicImageResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDynamicImageResponse) GoString() string {
	return s.String()
}

func (s *DeleteDynamicImageResponse) SetHeaders(v map[string]*string) *DeleteDynamicImageResponse {
	s.Headers = v
	return s
}

func (s *DeleteDynamicImageResponse) SetBody(v *DeleteDynamicImageResponseBody) *DeleteDynamicImageResponse {
	s.Body = v
	return s
}

type DeleteEditingProjectRequest struct {
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ProjectIds           *string `json:"ProjectIds,omitempty" xml:"ProjectIds,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *string `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DeleteEditingProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteEditingProjectRequest) GoString() string {
	return s.String()
}

func (s *DeleteEditingProjectRequest) SetOwnerAccount(v string) *DeleteEditingProjectRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteEditingProjectRequest) SetOwnerId(v string) *DeleteEditingProjectRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteEditingProjectRequest) SetProjectIds(v string) *DeleteEditingProjectRequest {
	s.ProjectIds = &v
	return s
}

func (s *DeleteEditingProjectRequest) SetResourceOwnerAccount(v string) *DeleteEditingProjectRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteEditingProjectRequest) SetResourceOwnerId(v string) *DeleteEditingProjectRequest {
	s.ResourceOwnerId = &v
	return s
}

type DeleteEditingProjectResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteEditingProjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteEditingProjectResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteEditingProjectResponseBody) SetRequestId(v string) *DeleteEditingProjectResponseBody {
	s.RequestId = &v
	return s
}

type DeleteEditingProjectResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteEditingProjectResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteEditingProjectResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteEditingProjectResponse) GoString() string {
	return s.String()
}

func (s *DeleteEditingProjectResponse) SetHeaders(v map[string]*string) *DeleteEditingProjectResponse {
	s.Headers = v
	return s
}

func (s *DeleteEditingProjectResponse) SetBody(v *DeleteEditingProjectResponseBody) *DeleteEditingProjectResponse {
	s.Body = v
	return s
}

type DeleteImageRequest struct {
	DeleteImageType *string `json:"DeleteImageType,omitempty" xml:"DeleteImageType,omitempty"`
	ImageIds        *string `json:"ImageIds,omitempty" xml:"ImageIds,omitempty"`
	ImageType       *string `json:"ImageType,omitempty" xml:"ImageType,omitempty"`
	ImageURLs       *string `json:"ImageURLs,omitempty" xml:"ImageURLs,omitempty"`
	VideoId         *string `json:"VideoId,omitempty" xml:"VideoId,omitempty"`
}

func (s DeleteImageRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteImageRequest) GoString() string {
	return s.String()
}

func (s *DeleteImageRequest) SetDeleteImageType(v string) *DeleteImageRequest {
	s.DeleteImageType = &v
	return s
}

func (s *DeleteImageRequest) SetImageIds(v string) *DeleteImageRequest {
	s.ImageIds = &v
	return s
}

func (s *DeleteImageRequest) SetImageType(v string) *DeleteImageRequest {
	s.ImageType = &v
	return s
}

func (s *DeleteImageRequest) SetImageURLs(v string) *DeleteImageRequest {
	s.ImageURLs = &v
	return s
}

func (s *DeleteImageRequest) SetVideoId(v string) *DeleteImageRequest {
	s.VideoId = &v
	return s
}

type DeleteImageResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteImageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteImageResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteImageResponseBody) SetRequestId(v string) *DeleteImageResponseBody {
	s.RequestId = &v
	return s
}

type DeleteImageResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteImageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteImageResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteImageResponse) GoString() string {
	return s.String()
}

func (s *DeleteImageResponse) SetHeaders(v map[string]*string) *DeleteImageResponse {
	s.Headers = v
	return s
}

func (s *DeleteImageResponse) SetBody(v *DeleteImageResponseBody) *DeleteImageResponse {
	s.Body = v
	return s
}

type DeleteMessageCallbackRequest struct {
	AppId        *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
}

func (s DeleteMessageCallbackRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteMessageCallbackRequest) GoString() string {
	return s.String()
}

func (s *DeleteMessageCallbackRequest) SetAppId(v string) *DeleteMessageCallbackRequest {
	s.AppId = &v
	return s
}

func (s *DeleteMessageCallbackRequest) SetOwnerAccount(v string) *DeleteMessageCallbackRequest {
	s.OwnerAccount = &v
	return s
}

type DeleteMessageCallbackResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteMessageCallbackResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteMessageCallbackResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteMessageCallbackResponseBody) SetRequestId(v string) *DeleteMessageCallbackResponseBody {
	s.RequestId = &v
	return s
}

type DeleteMessageCallbackResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteMessageCallbackResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteMessageCallbackResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteMessageCallbackResponse) GoString() string {
	return s.String()
}

func (s *DeleteMessageCallbackResponse) SetHeaders(v map[string]*string) *DeleteMessageCallbackResponse {
	s.Headers = v
	return s
}

func (s *DeleteMessageCallbackResponse) SetBody(v *DeleteMessageCallbackResponseBody) *DeleteMessageCallbackResponse {
	s.Body = v
	return s
}

type DeleteMezzaninesRequest struct {
	Force    *bool   `json:"Force,omitempty" xml:"Force,omitempty"`
	VideoIds *string `json:"VideoIds,omitempty" xml:"VideoIds,omitempty"`
}

func (s DeleteMezzaninesRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteMezzaninesRequest) GoString() string {
	return s.String()
}

func (s *DeleteMezzaninesRequest) SetForce(v bool) *DeleteMezzaninesRequest {
	s.Force = &v
	return s
}

func (s *DeleteMezzaninesRequest) SetVideoIds(v string) *DeleteMezzaninesRequest {
	s.VideoIds = &v
	return s
}

type DeleteMezzaninesResponseBody struct {
	NonExistVideoIds     []*string `json:"NonExistVideoIds,omitempty" xml:"NonExistVideoIds,omitempty" type:"Repeated"`
	RequestId            *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	UnRemoveableVideoIds []*string `json:"UnRemoveableVideoIds,omitempty" xml:"UnRemoveableVideoIds,omitempty" type:"Repeated"`
}

func (s DeleteMezzaninesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteMezzaninesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteMezzaninesResponseBody) SetNonExistVideoIds(v []*string) *DeleteMezzaninesResponseBody {
	s.NonExistVideoIds = v
	return s
}

func (s *DeleteMezzaninesResponseBody) SetRequestId(v string) *DeleteMezzaninesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteMezzaninesResponseBody) SetUnRemoveableVideoIds(v []*string) *DeleteMezzaninesResponseBody {
	s.UnRemoveableVideoIds = v
	return s
}

type DeleteMezzaninesResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteMezzaninesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteMezzaninesResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteMezzaninesResponse) GoString() string {
	return s.String()
}

func (s *DeleteMezzaninesResponse) SetHeaders(v map[string]*string) *DeleteMezzaninesResponse {
	s.Headers = v
	return s
}

func (s *DeleteMezzaninesResponse) SetBody(v *DeleteMezzaninesResponseBody) *DeleteMezzaninesResponse {
	s.Body = v
	return s
}

type DeleteMultipartUploadRequest struct {
	MediaId      *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	MediaType    *string `json:"MediaType,omitempty" xml:"MediaType,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
}

func (s DeleteMultipartUploadRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteMultipartUploadRequest) GoString() string {
	return s.String()
}

func (s *DeleteMultipartUploadRequest) SetMediaId(v string) *DeleteMultipartUploadRequest {
	s.MediaId = &v
	return s
}

func (s *DeleteMultipartUploadRequest) SetMediaType(v string) *DeleteMultipartUploadRequest {
	s.MediaType = &v
	return s
}

func (s *DeleteMultipartUploadRequest) SetOwnerAccount(v string) *DeleteMultipartUploadRequest {
	s.OwnerAccount = &v
	return s
}

type DeleteMultipartUploadResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteMultipartUploadResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteMultipartUploadResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteMultipartUploadResponseBody) SetRequestId(v string) *DeleteMultipartUploadResponseBody {
	s.RequestId = &v
	return s
}

type DeleteMultipartUploadResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteMultipartUploadResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteMultipartUploadResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteMultipartUploadResponse) GoString() string {
	return s.String()
}

func (s *DeleteMultipartUploadResponse) SetHeaders(v map[string]*string) *DeleteMultipartUploadResponse {
	s.Headers = v
	return s
}

func (s *DeleteMultipartUploadResponse) SetBody(v *DeleteMultipartUploadResponseBody) *DeleteMultipartUploadResponse {
	s.Body = v
	return s
}

type DeleteStreamRequest struct {
	JobIds  *string `json:"JobIds,omitempty" xml:"JobIds,omitempty"`
	VideoId *string `json:"VideoId,omitempty" xml:"VideoId,omitempty"`
}

func (s DeleteStreamRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteStreamRequest) GoString() string {
	return s.String()
}

func (s *DeleteStreamRequest) SetJobIds(v string) *DeleteStreamRequest {
	s.JobIds = &v
	return s
}

func (s *DeleteStreamRequest) SetVideoId(v string) *DeleteStreamRequest {
	s.VideoId = &v
	return s
}

type DeleteStreamResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteStreamResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteStreamResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteStreamResponseBody) SetRequestId(v string) *DeleteStreamResponseBody {
	s.RequestId = &v
	return s
}

type DeleteStreamResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteStreamResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteStreamResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteStreamResponse) GoString() string {
	return s.String()
}

func (s *DeleteStreamResponse) SetHeaders(v map[string]*string) *DeleteStreamResponse {
	s.Headers = v
	return s
}

func (s *DeleteStreamResponse) SetBody(v *DeleteStreamResponseBody) *DeleteStreamResponse {
	s.Body = v
	return s
}

type DeleteTranscodeTemplateGroupRequest struct {
	ForceDelGroup            *string `json:"ForceDelGroup,omitempty" xml:"ForceDelGroup,omitempty"`
	TranscodeTemplateGroupId *string `json:"TranscodeTemplateGroupId,omitempty" xml:"TranscodeTemplateGroupId,omitempty"`
	TranscodeTemplateIds     *string `json:"TranscodeTemplateIds,omitempty" xml:"TranscodeTemplateIds,omitempty"`
}

func (s DeleteTranscodeTemplateGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteTranscodeTemplateGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteTranscodeTemplateGroupRequest) SetForceDelGroup(v string) *DeleteTranscodeTemplateGroupRequest {
	s.ForceDelGroup = &v
	return s
}

func (s *DeleteTranscodeTemplateGroupRequest) SetTranscodeTemplateGroupId(v string) *DeleteTranscodeTemplateGroupRequest {
	s.TranscodeTemplateGroupId = &v
	return s
}

func (s *DeleteTranscodeTemplateGroupRequest) SetTranscodeTemplateIds(v string) *DeleteTranscodeTemplateGroupRequest {
	s.TranscodeTemplateIds = &v
	return s
}

type DeleteTranscodeTemplateGroupResponseBody struct {
	NonExistTranscodeTemplateIds []*string `json:"NonExistTranscodeTemplateIds,omitempty" xml:"NonExistTranscodeTemplateIds,omitempty" type:"Repeated"`
	RequestId                    *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteTranscodeTemplateGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteTranscodeTemplateGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteTranscodeTemplateGroupResponseBody) SetNonExistTranscodeTemplateIds(v []*string) *DeleteTranscodeTemplateGroupResponseBody {
	s.NonExistTranscodeTemplateIds = v
	return s
}

func (s *DeleteTranscodeTemplateGroupResponseBody) SetRequestId(v string) *DeleteTranscodeTemplateGroupResponseBody {
	s.RequestId = &v
	return s
}

type DeleteTranscodeTemplateGroupResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteTranscodeTemplateGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteTranscodeTemplateGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteTranscodeTemplateGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteTranscodeTemplateGroupResponse) SetHeaders(v map[string]*string) *DeleteTranscodeTemplateGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteTranscodeTemplateGroupResponse) SetBody(v *DeleteTranscodeTemplateGroupResponseBody) *DeleteTranscodeTemplateGroupResponse {
	s.Body = v
	return s
}

type DeleteVideoRequest struct {
	VideoIds *string `json:"VideoIds,omitempty" xml:"VideoIds,omitempty"`
}

func (s DeleteVideoRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteVideoRequest) GoString() string {
	return s.String()
}

func (s *DeleteVideoRequest) SetVideoIds(v string) *DeleteVideoRequest {
	s.VideoIds = &v
	return s
}

type DeleteVideoResponseBody struct {
	ForbiddenVideoIds []*string `json:"ForbiddenVideoIds,omitempty" xml:"ForbiddenVideoIds,omitempty" type:"Repeated"`
	NonExistVideoIds  []*string `json:"NonExistVideoIds,omitempty" xml:"NonExistVideoIds,omitempty" type:"Repeated"`
	RequestId         *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteVideoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteVideoResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteVideoResponseBody) SetForbiddenVideoIds(v []*string) *DeleteVideoResponseBody {
	s.ForbiddenVideoIds = v
	return s
}

func (s *DeleteVideoResponseBody) SetNonExistVideoIds(v []*string) *DeleteVideoResponseBody {
	s.NonExistVideoIds = v
	return s
}

func (s *DeleteVideoResponseBody) SetRequestId(v string) *DeleteVideoResponseBody {
	s.RequestId = &v
	return s
}

type DeleteVideoResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteVideoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteVideoResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteVideoResponse) GoString() string {
	return s.String()
}

func (s *DeleteVideoResponse) SetHeaders(v map[string]*string) *DeleteVideoResponse {
	s.Headers = v
	return s
}

func (s *DeleteVideoResponse) SetBody(v *DeleteVideoResponseBody) *DeleteVideoResponse {
	s.Body = v
	return s
}

type DeleteVodDomainRequest struct {
	DomainName    *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	OwnerAccount  *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DeleteVodDomainRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteVodDomainRequest) GoString() string {
	return s.String()
}

func (s *DeleteVodDomainRequest) SetDomainName(v string) *DeleteVodDomainRequest {
	s.DomainName = &v
	return s
}

func (s *DeleteVodDomainRequest) SetOwnerAccount(v string) *DeleteVodDomainRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteVodDomainRequest) SetOwnerId(v int64) *DeleteVodDomainRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteVodDomainRequest) SetSecurityToken(v string) *DeleteVodDomainRequest {
	s.SecurityToken = &v
	return s
}

type DeleteVodDomainResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteVodDomainResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteVodDomainResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteVodDomainResponseBody) SetRequestId(v string) *DeleteVodDomainResponseBody {
	s.RequestId = &v
	return s
}

type DeleteVodDomainResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteVodDomainResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteVodDomainResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteVodDomainResponse) GoString() string {
	return s.String()
}

func (s *DeleteVodDomainResponse) SetHeaders(v map[string]*string) *DeleteVodDomainResponse {
	s.Headers = v
	return s
}

func (s *DeleteVodDomainResponse) SetBody(v *DeleteVodDomainResponseBody) *DeleteVodDomainResponse {
	s.Body = v
	return s
}

type DeleteVodRealtimeLogDeliveryRequest struct {
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	Logstore   *string `json:"Logstore,omitempty" xml:"Logstore,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	Project    *string `json:"Project,omitempty" xml:"Project,omitempty"`
	Region     *string `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s DeleteVodRealtimeLogDeliveryRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteVodRealtimeLogDeliveryRequest) GoString() string {
	return s.String()
}

func (s *DeleteVodRealtimeLogDeliveryRequest) SetDomainName(v string) *DeleteVodRealtimeLogDeliveryRequest {
	s.DomainName = &v
	return s
}

func (s *DeleteVodRealtimeLogDeliveryRequest) SetLogstore(v string) *DeleteVodRealtimeLogDeliveryRequest {
	s.Logstore = &v
	return s
}

func (s *DeleteVodRealtimeLogDeliveryRequest) SetOwnerId(v int64) *DeleteVodRealtimeLogDeliveryRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteVodRealtimeLogDeliveryRequest) SetProject(v string) *DeleteVodRealtimeLogDeliveryRequest {
	s.Project = &v
	return s
}

func (s *DeleteVodRealtimeLogDeliveryRequest) SetRegion(v string) *DeleteVodRealtimeLogDeliveryRequest {
	s.Region = &v
	return s
}

type DeleteVodRealtimeLogDeliveryResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteVodRealtimeLogDeliveryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteVodRealtimeLogDeliveryResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteVodRealtimeLogDeliveryResponseBody) SetRequestId(v string) *DeleteVodRealtimeLogDeliveryResponseBody {
	s.RequestId = &v
	return s
}

type DeleteVodRealtimeLogDeliveryResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteVodRealtimeLogDeliveryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteVodRealtimeLogDeliveryResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteVodRealtimeLogDeliveryResponse) GoString() string {
	return s.String()
}

func (s *DeleteVodRealtimeLogDeliveryResponse) SetHeaders(v map[string]*string) *DeleteVodRealtimeLogDeliveryResponse {
	s.Headers = v
	return s
}

func (s *DeleteVodRealtimeLogDeliveryResponse) SetBody(v *DeleteVodRealtimeLogDeliveryResponseBody) *DeleteVodRealtimeLogDeliveryResponse {
	s.Body = v
	return s
}

type DeleteVodSpecificConfigRequest struct {
	ConfigId      *string `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	DomainName    *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DeleteVodSpecificConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteVodSpecificConfigRequest) GoString() string {
	return s.String()
}

func (s *DeleteVodSpecificConfigRequest) SetConfigId(v string) *DeleteVodSpecificConfigRequest {
	s.ConfigId = &v
	return s
}

func (s *DeleteVodSpecificConfigRequest) SetDomainName(v string) *DeleteVodSpecificConfigRequest {
	s.DomainName = &v
	return s
}

func (s *DeleteVodSpecificConfigRequest) SetOwnerId(v int64) *DeleteVodSpecificConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteVodSpecificConfigRequest) SetSecurityToken(v string) *DeleteVodSpecificConfigRequest {
	s.SecurityToken = &v
	return s
}

type DeleteVodSpecificConfigResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteVodSpecificConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteVodSpecificConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteVodSpecificConfigResponseBody) SetRequestId(v string) *DeleteVodSpecificConfigResponseBody {
	s.RequestId = &v
	return s
}

type DeleteVodSpecificConfigResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteVodSpecificConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteVodSpecificConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteVodSpecificConfigResponse) GoString() string {
	return s.String()
}

func (s *DeleteVodSpecificConfigResponse) SetHeaders(v map[string]*string) *DeleteVodSpecificConfigResponse {
	s.Headers = v
	return s
}

func (s *DeleteVodSpecificConfigResponse) SetBody(v *DeleteVodSpecificConfigResponseBody) *DeleteVodSpecificConfigResponse {
	s.Body = v
	return s
}

type DeleteVodTemplateRequest struct {
	VodTemplateId *string `json:"VodTemplateId,omitempty" xml:"VodTemplateId,omitempty"`
}

func (s DeleteVodTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteVodTemplateRequest) GoString() string {
	return s.String()
}

func (s *DeleteVodTemplateRequest) SetVodTemplateId(v string) *DeleteVodTemplateRequest {
	s.VodTemplateId = &v
	return s
}

type DeleteVodTemplateResponseBody struct {
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	VodTemplateId *string `json:"VodTemplateId,omitempty" xml:"VodTemplateId,omitempty"`
}

func (s DeleteVodTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteVodTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteVodTemplateResponseBody) SetRequestId(v string) *DeleteVodTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteVodTemplateResponseBody) SetVodTemplateId(v string) *DeleteVodTemplateResponseBody {
	s.VodTemplateId = &v
	return s
}

type DeleteVodTemplateResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteVodTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteVodTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteVodTemplateResponse) GoString() string {
	return s.String()
}

func (s *DeleteVodTemplateResponse) SetHeaders(v map[string]*string) *DeleteVodTemplateResponse {
	s.Headers = v
	return s
}

func (s *DeleteVodTemplateResponse) SetBody(v *DeleteVodTemplateResponseBody) *DeleteVodTemplateResponse {
	s.Body = v
	return s
}

type DeleteWatermarkRequest struct {
	WatermarkId *string `json:"WatermarkId,omitempty" xml:"WatermarkId,omitempty"`
}

func (s DeleteWatermarkRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteWatermarkRequest) GoString() string {
	return s.String()
}

func (s *DeleteWatermarkRequest) SetWatermarkId(v string) *DeleteWatermarkRequest {
	s.WatermarkId = &v
	return s
}

type DeleteWatermarkResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteWatermarkResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteWatermarkResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteWatermarkResponseBody) SetRequestId(v string) *DeleteWatermarkResponseBody {
	s.RequestId = &v
	return s
}

type DeleteWatermarkResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteWatermarkResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteWatermarkResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteWatermarkResponse) GoString() string {
	return s.String()
}

func (s *DeleteWatermarkResponse) SetHeaders(v map[string]*string) *DeleteWatermarkResponse {
	s.Headers = v
	return s
}

func (s *DeleteWatermarkResponse) SetBody(v *DeleteWatermarkResponseBody) *DeleteWatermarkResponse {
	s.Body = v
	return s
}

type DescribePlayTopVideosRequest struct {
	BizDate  *string `json:"BizDate,omitempty" xml:"BizDate,omitempty"`
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNo   *int64  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribePlayTopVideosRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribePlayTopVideosRequest) GoString() string {
	return s.String()
}

func (s *DescribePlayTopVideosRequest) SetBizDate(v string) *DescribePlayTopVideosRequest {
	s.BizDate = &v
	return s
}

func (s *DescribePlayTopVideosRequest) SetOwnerId(v int64) *DescribePlayTopVideosRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribePlayTopVideosRequest) SetPageNo(v int64) *DescribePlayTopVideosRequest {
	s.PageNo = &v
	return s
}

func (s *DescribePlayTopVideosRequest) SetPageSize(v int64) *DescribePlayTopVideosRequest {
	s.PageSize = &v
	return s
}

type DescribePlayTopVideosResponseBody struct {
	PageNo        *int64                                          `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize      *int64                                          `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId     *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TopPlayVideos *DescribePlayTopVideosResponseBodyTopPlayVideos `json:"TopPlayVideos,omitempty" xml:"TopPlayVideos,omitempty" type:"Struct"`
	TotalNum      *int64                                          `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
}

func (s DescribePlayTopVideosResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribePlayTopVideosResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePlayTopVideosResponseBody) SetPageNo(v int64) *DescribePlayTopVideosResponseBody {
	s.PageNo = &v
	return s
}

func (s *DescribePlayTopVideosResponseBody) SetPageSize(v int64) *DescribePlayTopVideosResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribePlayTopVideosResponseBody) SetRequestId(v string) *DescribePlayTopVideosResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePlayTopVideosResponseBody) SetTopPlayVideos(v *DescribePlayTopVideosResponseBodyTopPlayVideos) *DescribePlayTopVideosResponseBody {
	s.TopPlayVideos = v
	return s
}

func (s *DescribePlayTopVideosResponseBody) SetTotalNum(v int64) *DescribePlayTopVideosResponseBody {
	s.TotalNum = &v
	return s
}

type DescribePlayTopVideosResponseBodyTopPlayVideos struct {
	TopPlayVideoStatis []*DescribePlayTopVideosResponseBodyTopPlayVideosTopPlayVideoStatis `json:"TopPlayVideoStatis,omitempty" xml:"TopPlayVideoStatis,omitempty" type:"Repeated"`
}

func (s DescribePlayTopVideosResponseBodyTopPlayVideos) String() string {
	return tea.Prettify(s)
}

func (s DescribePlayTopVideosResponseBodyTopPlayVideos) GoString() string {
	return s.String()
}

func (s *DescribePlayTopVideosResponseBodyTopPlayVideos) SetTopPlayVideoStatis(v []*DescribePlayTopVideosResponseBodyTopPlayVideosTopPlayVideoStatis) *DescribePlayTopVideosResponseBodyTopPlayVideos {
	s.TopPlayVideoStatis = v
	return s
}

type DescribePlayTopVideosResponseBodyTopPlayVideosTopPlayVideoStatis struct {
	PlayDuration *string `json:"PlayDuration,omitempty" xml:"PlayDuration,omitempty"`
	Title        *string `json:"Title,omitempty" xml:"Title,omitempty"`
	UV           *string `json:"UV,omitempty" xml:"UV,omitempty"`
	VV           *string `json:"VV,omitempty" xml:"VV,omitempty"`
	VideoId      *string `json:"VideoId,omitempty" xml:"VideoId,omitempty"`
}

func (s DescribePlayTopVideosResponseBodyTopPlayVideosTopPlayVideoStatis) String() string {
	return tea.Prettify(s)
}

func (s DescribePlayTopVideosResponseBodyTopPlayVideosTopPlayVideoStatis) GoString() string {
	return s.String()
}

func (s *DescribePlayTopVideosResponseBodyTopPlayVideosTopPlayVideoStatis) SetPlayDuration(v string) *DescribePlayTopVideosResponseBodyTopPlayVideosTopPlayVideoStatis {
	s.PlayDuration = &v
	return s
}

func (s *DescribePlayTopVideosResponseBodyTopPlayVideosTopPlayVideoStatis) SetTitle(v string) *DescribePlayTopVideosResponseBodyTopPlayVideosTopPlayVideoStatis {
	s.Title = &v
	return s
}

func (s *DescribePlayTopVideosResponseBodyTopPlayVideosTopPlayVideoStatis) SetUV(v string) *DescribePlayTopVideosResponseBodyTopPlayVideosTopPlayVideoStatis {
	s.UV = &v
	return s
}

func (s *DescribePlayTopVideosResponseBodyTopPlayVideosTopPlayVideoStatis) SetVV(v string) *DescribePlayTopVideosResponseBodyTopPlayVideosTopPlayVideoStatis {
	s.VV = &v
	return s
}

func (s *DescribePlayTopVideosResponseBodyTopPlayVideosTopPlayVideoStatis) SetVideoId(v string) *DescribePlayTopVideosResponseBodyTopPlayVideosTopPlayVideoStatis {
	s.VideoId = &v
	return s
}

type DescribePlayTopVideosResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribePlayTopVideosResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribePlayTopVideosResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribePlayTopVideosResponse) GoString() string {
	return s.String()
}

func (s *DescribePlayTopVideosResponse) SetHeaders(v map[string]*string) *DescribePlayTopVideosResponse {
	s.Headers = v
	return s
}

func (s *DescribePlayTopVideosResponse) SetBody(v *DescribePlayTopVideosResponseBody) *DescribePlayTopVideosResponse {
	s.Body = v
	return s
}

type DescribePlayUserAvgRequest struct {
	EndTime   *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OwnerId   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribePlayUserAvgRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribePlayUserAvgRequest) GoString() string {
	return s.String()
}

func (s *DescribePlayUserAvgRequest) SetEndTime(v string) *DescribePlayUserAvgRequest {
	s.EndTime = &v
	return s
}

func (s *DescribePlayUserAvgRequest) SetOwnerId(v int64) *DescribePlayUserAvgRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribePlayUserAvgRequest) SetStartTime(v string) *DescribePlayUserAvgRequest {
	s.StartTime = &v
	return s
}

type DescribePlayUserAvgResponseBody struct {
	RequestId          *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	UserPlayStatisAvgs *DescribePlayUserAvgResponseBodyUserPlayStatisAvgs `json:"UserPlayStatisAvgs,omitempty" xml:"UserPlayStatisAvgs,omitempty" type:"Struct"`
}

func (s DescribePlayUserAvgResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribePlayUserAvgResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePlayUserAvgResponseBody) SetRequestId(v string) *DescribePlayUserAvgResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePlayUserAvgResponseBody) SetUserPlayStatisAvgs(v *DescribePlayUserAvgResponseBodyUserPlayStatisAvgs) *DescribePlayUserAvgResponseBody {
	s.UserPlayStatisAvgs = v
	return s
}

type DescribePlayUserAvgResponseBodyUserPlayStatisAvgs struct {
	UserPlayStatisAvg []*DescribePlayUserAvgResponseBodyUserPlayStatisAvgsUserPlayStatisAvg `json:"UserPlayStatisAvg,omitempty" xml:"UserPlayStatisAvg,omitempty" type:"Repeated"`
}

func (s DescribePlayUserAvgResponseBodyUserPlayStatisAvgs) String() string {
	return tea.Prettify(s)
}

func (s DescribePlayUserAvgResponseBodyUserPlayStatisAvgs) GoString() string {
	return s.String()
}

func (s *DescribePlayUserAvgResponseBodyUserPlayStatisAvgs) SetUserPlayStatisAvg(v []*DescribePlayUserAvgResponseBodyUserPlayStatisAvgsUserPlayStatisAvg) *DescribePlayUserAvgResponseBodyUserPlayStatisAvgs {
	s.UserPlayStatisAvg = v
	return s
}

type DescribePlayUserAvgResponseBodyUserPlayStatisAvgsUserPlayStatisAvg struct {
	AvgPlayCount    *string `json:"AvgPlayCount,omitempty" xml:"AvgPlayCount,omitempty"`
	AvgPlayDuration *string `json:"AvgPlayDuration,omitempty" xml:"AvgPlayDuration,omitempty"`
	Date            *string `json:"Date,omitempty" xml:"Date,omitempty"`
}

func (s DescribePlayUserAvgResponseBodyUserPlayStatisAvgsUserPlayStatisAvg) String() string {
	return tea.Prettify(s)
}

func (s DescribePlayUserAvgResponseBodyUserPlayStatisAvgsUserPlayStatisAvg) GoString() string {
	return s.String()
}

func (s *DescribePlayUserAvgResponseBodyUserPlayStatisAvgsUserPlayStatisAvg) SetAvgPlayCount(v string) *DescribePlayUserAvgResponseBodyUserPlayStatisAvgsUserPlayStatisAvg {
	s.AvgPlayCount = &v
	return s
}

func (s *DescribePlayUserAvgResponseBodyUserPlayStatisAvgsUserPlayStatisAvg) SetAvgPlayDuration(v string) *DescribePlayUserAvgResponseBodyUserPlayStatisAvgsUserPlayStatisAvg {
	s.AvgPlayDuration = &v
	return s
}

func (s *DescribePlayUserAvgResponseBodyUserPlayStatisAvgsUserPlayStatisAvg) SetDate(v string) *DescribePlayUserAvgResponseBodyUserPlayStatisAvgsUserPlayStatisAvg {
	s.Date = &v
	return s
}

type DescribePlayUserAvgResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribePlayUserAvgResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribePlayUserAvgResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribePlayUserAvgResponse) GoString() string {
	return s.String()
}

func (s *DescribePlayUserAvgResponse) SetHeaders(v map[string]*string) *DescribePlayUserAvgResponse {
	s.Headers = v
	return s
}

func (s *DescribePlayUserAvgResponse) SetBody(v *DescribePlayUserAvgResponseBody) *DescribePlayUserAvgResponse {
	s.Body = v
	return s
}

type DescribePlayUserTotalRequest struct {
	EndTime   *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OwnerId   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribePlayUserTotalRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribePlayUserTotalRequest) GoString() string {
	return s.String()
}

func (s *DescribePlayUserTotalRequest) SetEndTime(v string) *DescribePlayUserTotalRequest {
	s.EndTime = &v
	return s
}

func (s *DescribePlayUserTotalRequest) SetOwnerId(v int64) *DescribePlayUserTotalRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribePlayUserTotalRequest) SetStartTime(v string) *DescribePlayUserTotalRequest {
	s.StartTime = &v
	return s
}

type DescribePlayUserTotalResponseBody struct {
	RequestId            *string                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	UserPlayStatisTotals *DescribePlayUserTotalResponseBodyUserPlayStatisTotals `json:"UserPlayStatisTotals,omitempty" xml:"UserPlayStatisTotals,omitempty" type:"Struct"`
}

func (s DescribePlayUserTotalResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribePlayUserTotalResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePlayUserTotalResponseBody) SetRequestId(v string) *DescribePlayUserTotalResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePlayUserTotalResponseBody) SetUserPlayStatisTotals(v *DescribePlayUserTotalResponseBodyUserPlayStatisTotals) *DescribePlayUserTotalResponseBody {
	s.UserPlayStatisTotals = v
	return s
}

type DescribePlayUserTotalResponseBodyUserPlayStatisTotals struct {
	UserPlayStatisTotal []*DescribePlayUserTotalResponseBodyUserPlayStatisTotalsUserPlayStatisTotal `json:"UserPlayStatisTotal,omitempty" xml:"UserPlayStatisTotal,omitempty" type:"Repeated"`
}

func (s DescribePlayUserTotalResponseBodyUserPlayStatisTotals) String() string {
	return tea.Prettify(s)
}

func (s DescribePlayUserTotalResponseBodyUserPlayStatisTotals) GoString() string {
	return s.String()
}

func (s *DescribePlayUserTotalResponseBodyUserPlayStatisTotals) SetUserPlayStatisTotal(v []*DescribePlayUserTotalResponseBodyUserPlayStatisTotalsUserPlayStatisTotal) *DescribePlayUserTotalResponseBodyUserPlayStatisTotals {
	s.UserPlayStatisTotal = v
	return s
}

type DescribePlayUserTotalResponseBodyUserPlayStatisTotalsUserPlayStatisTotal struct {
	Date         *string                                                                     `json:"Date,omitempty" xml:"Date,omitempty"`
	PlayDuration *string                                                                     `json:"PlayDuration,omitempty" xml:"PlayDuration,omitempty"`
	PlayRange    *string                                                                     `json:"PlayRange,omitempty" xml:"PlayRange,omitempty"`
	UV           *DescribePlayUserTotalResponseBodyUserPlayStatisTotalsUserPlayStatisTotalUV `json:"UV,omitempty" xml:"UV,omitempty" type:"Struct"`
	VV           *DescribePlayUserTotalResponseBodyUserPlayStatisTotalsUserPlayStatisTotalVV `json:"VV,omitempty" xml:"VV,omitempty" type:"Struct"`
}

func (s DescribePlayUserTotalResponseBodyUserPlayStatisTotalsUserPlayStatisTotal) String() string {
	return tea.Prettify(s)
}

func (s DescribePlayUserTotalResponseBodyUserPlayStatisTotalsUserPlayStatisTotal) GoString() string {
	return s.String()
}

func (s *DescribePlayUserTotalResponseBodyUserPlayStatisTotalsUserPlayStatisTotal) SetDate(v string) *DescribePlayUserTotalResponseBodyUserPlayStatisTotalsUserPlayStatisTotal {
	s.Date = &v
	return s
}

func (s *DescribePlayUserTotalResponseBodyUserPlayStatisTotalsUserPlayStatisTotal) SetPlayDuration(v string) *DescribePlayUserTotalResponseBodyUserPlayStatisTotalsUserPlayStatisTotal {
	s.PlayDuration = &v
	return s
}

func (s *DescribePlayUserTotalResponseBodyUserPlayStatisTotalsUserPlayStatisTotal) SetPlayRange(v string) *DescribePlayUserTotalResponseBodyUserPlayStatisTotalsUserPlayStatisTotal {
	s.PlayRange = &v
	return s
}

func (s *DescribePlayUserTotalResponseBodyUserPlayStatisTotalsUserPlayStatisTotal) SetUV(v *DescribePlayUserTotalResponseBodyUserPlayStatisTotalsUserPlayStatisTotalUV) *DescribePlayUserTotalResponseBodyUserPlayStatisTotalsUserPlayStatisTotal {
	s.UV = v
	return s
}

func (s *DescribePlayUserTotalResponseBodyUserPlayStatisTotalsUserPlayStatisTotal) SetVV(v *DescribePlayUserTotalResponseBodyUserPlayStatisTotalsUserPlayStatisTotalVV) *DescribePlayUserTotalResponseBodyUserPlayStatisTotalsUserPlayStatisTotal {
	s.VV = v
	return s
}

type DescribePlayUserTotalResponseBodyUserPlayStatisTotalsUserPlayStatisTotalUV struct {
	Android *string `json:"Android,omitempty" xml:"Android,omitempty"`
	Flash   *string `json:"Flash,omitempty" xml:"Flash,omitempty"`
	HTML5   *string `json:"HTML5,omitempty" xml:"HTML5,omitempty"`
	IOS     *string `json:"iOS,omitempty" xml:"iOS,omitempty"`
}

func (s DescribePlayUserTotalResponseBodyUserPlayStatisTotalsUserPlayStatisTotalUV) String() string {
	return tea.Prettify(s)
}

func (s DescribePlayUserTotalResponseBodyUserPlayStatisTotalsUserPlayStatisTotalUV) GoString() string {
	return s.String()
}

func (s *DescribePlayUserTotalResponseBodyUserPlayStatisTotalsUserPlayStatisTotalUV) SetAndroid(v string) *DescribePlayUserTotalResponseBodyUserPlayStatisTotalsUserPlayStatisTotalUV {
	s.Android = &v
	return s
}

func (s *DescribePlayUserTotalResponseBodyUserPlayStatisTotalsUserPlayStatisTotalUV) SetFlash(v string) *DescribePlayUserTotalResponseBodyUserPlayStatisTotalsUserPlayStatisTotalUV {
	s.Flash = &v
	return s
}

func (s *DescribePlayUserTotalResponseBodyUserPlayStatisTotalsUserPlayStatisTotalUV) SetHTML5(v string) *DescribePlayUserTotalResponseBodyUserPlayStatisTotalsUserPlayStatisTotalUV {
	s.HTML5 = &v
	return s
}

func (s *DescribePlayUserTotalResponseBodyUserPlayStatisTotalsUserPlayStatisTotalUV) SetIOS(v string) *DescribePlayUserTotalResponseBodyUserPlayStatisTotalsUserPlayStatisTotalUV {
	s.IOS = &v
	return s
}

type DescribePlayUserTotalResponseBodyUserPlayStatisTotalsUserPlayStatisTotalVV struct {
	Android *string `json:"Android,omitempty" xml:"Android,omitempty"`
	Flash   *string `json:"Flash,omitempty" xml:"Flash,omitempty"`
	HTML5   *string `json:"HTML5,omitempty" xml:"HTML5,omitempty"`
	IOS     *string `json:"iOS,omitempty" xml:"iOS,omitempty"`
}

func (s DescribePlayUserTotalResponseBodyUserPlayStatisTotalsUserPlayStatisTotalVV) String() string {
	return tea.Prettify(s)
}

func (s DescribePlayUserTotalResponseBodyUserPlayStatisTotalsUserPlayStatisTotalVV) GoString() string {
	return s.String()
}

func (s *DescribePlayUserTotalResponseBodyUserPlayStatisTotalsUserPlayStatisTotalVV) SetAndroid(v string) *DescribePlayUserTotalResponseBodyUserPlayStatisTotalsUserPlayStatisTotalVV {
	s.Android = &v
	return s
}

func (s *DescribePlayUserTotalResponseBodyUserPlayStatisTotalsUserPlayStatisTotalVV) SetFlash(v string) *DescribePlayUserTotalResponseBodyUserPlayStatisTotalsUserPlayStatisTotalVV {
	s.Flash = &v
	return s
}

func (s *DescribePlayUserTotalResponseBodyUserPlayStatisTotalsUserPlayStatisTotalVV) SetHTML5(v string) *DescribePlayUserTotalResponseBodyUserPlayStatisTotalsUserPlayStatisTotalVV {
	s.HTML5 = &v
	return s
}

func (s *DescribePlayUserTotalResponseBodyUserPlayStatisTotalsUserPlayStatisTotalVV) SetIOS(v string) *DescribePlayUserTotalResponseBodyUserPlayStatisTotalsUserPlayStatisTotalVV {
	s.IOS = &v
	return s
}

type DescribePlayUserTotalResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribePlayUserTotalResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribePlayUserTotalResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribePlayUserTotalResponse) GoString() string {
	return s.String()
}

func (s *DescribePlayUserTotalResponse) SetHeaders(v map[string]*string) *DescribePlayUserTotalResponse {
	s.Headers = v
	return s
}

func (s *DescribePlayUserTotalResponse) SetBody(v *DescribePlayUserTotalResponseBody) *DescribePlayUserTotalResponse {
	s.Body = v
	return s
}

type DescribePlayVideoStatisRequest struct {
	EndTime   *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OwnerId   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	VideoId   *string `json:"VideoId,omitempty" xml:"VideoId,omitempty"`
}

func (s DescribePlayVideoStatisRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribePlayVideoStatisRequest) GoString() string {
	return s.String()
}

func (s *DescribePlayVideoStatisRequest) SetEndTime(v string) *DescribePlayVideoStatisRequest {
	s.EndTime = &v
	return s
}

func (s *DescribePlayVideoStatisRequest) SetOwnerId(v int64) *DescribePlayVideoStatisRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribePlayVideoStatisRequest) SetStartTime(v string) *DescribePlayVideoStatisRequest {
	s.StartTime = &v
	return s
}

func (s *DescribePlayVideoStatisRequest) SetVideoId(v string) *DescribePlayVideoStatisRequest {
	s.VideoId = &v
	return s
}

type DescribePlayVideoStatisResponseBody struct {
	RequestId              *string                                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	VideoPlayStatisDetails *DescribePlayVideoStatisResponseBodyVideoPlayStatisDetails `json:"VideoPlayStatisDetails,omitempty" xml:"VideoPlayStatisDetails,omitempty" type:"Struct"`
}

func (s DescribePlayVideoStatisResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribePlayVideoStatisResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePlayVideoStatisResponseBody) SetRequestId(v string) *DescribePlayVideoStatisResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePlayVideoStatisResponseBody) SetVideoPlayStatisDetails(v *DescribePlayVideoStatisResponseBodyVideoPlayStatisDetails) *DescribePlayVideoStatisResponseBody {
	s.VideoPlayStatisDetails = v
	return s
}

type DescribePlayVideoStatisResponseBodyVideoPlayStatisDetails struct {
	VideoPlayStatisDetail []*DescribePlayVideoStatisResponseBodyVideoPlayStatisDetailsVideoPlayStatisDetail `json:"VideoPlayStatisDetail,omitempty" xml:"VideoPlayStatisDetail,omitempty" type:"Repeated"`
}

func (s DescribePlayVideoStatisResponseBodyVideoPlayStatisDetails) String() string {
	return tea.Prettify(s)
}

func (s DescribePlayVideoStatisResponseBodyVideoPlayStatisDetails) GoString() string {
	return s.String()
}

func (s *DescribePlayVideoStatisResponseBodyVideoPlayStatisDetails) SetVideoPlayStatisDetail(v []*DescribePlayVideoStatisResponseBodyVideoPlayStatisDetailsVideoPlayStatisDetail) *DescribePlayVideoStatisResponseBodyVideoPlayStatisDetails {
	s.VideoPlayStatisDetail = v
	return s
}

type DescribePlayVideoStatisResponseBodyVideoPlayStatisDetailsVideoPlayStatisDetail struct {
	Date         *string `json:"Date,omitempty" xml:"Date,omitempty"`
	PlayDuration *string `json:"PlayDuration,omitempty" xml:"PlayDuration,omitempty"`
	PlayRange    *string `json:"PlayRange,omitempty" xml:"PlayRange,omitempty"`
	Title        *string `json:"Title,omitempty" xml:"Title,omitempty"`
	UV           *string `json:"UV,omitempty" xml:"UV,omitempty"`
	VV           *string `json:"VV,omitempty" xml:"VV,omitempty"`
}

func (s DescribePlayVideoStatisResponseBodyVideoPlayStatisDetailsVideoPlayStatisDetail) String() string {
	return tea.Prettify(s)
}

func (s DescribePlayVideoStatisResponseBodyVideoPlayStatisDetailsVideoPlayStatisDetail) GoString() string {
	return s.String()
}

func (s *DescribePlayVideoStatisResponseBodyVideoPlayStatisDetailsVideoPlayStatisDetail) SetDate(v string) *DescribePlayVideoStatisResponseBodyVideoPlayStatisDetailsVideoPlayStatisDetail {
	s.Date = &v
	return s
}

func (s *DescribePlayVideoStatisResponseBodyVideoPlayStatisDetailsVideoPlayStatisDetail) SetPlayDuration(v string) *DescribePlayVideoStatisResponseBodyVideoPlayStatisDetailsVideoPlayStatisDetail {
	s.PlayDuration = &v
	return s
}

func (s *DescribePlayVideoStatisResponseBodyVideoPlayStatisDetailsVideoPlayStatisDetail) SetPlayRange(v string) *DescribePlayVideoStatisResponseBodyVideoPlayStatisDetailsVideoPlayStatisDetail {
	s.PlayRange = &v
	return s
}

func (s *DescribePlayVideoStatisResponseBodyVideoPlayStatisDetailsVideoPlayStatisDetail) SetTitle(v string) *DescribePlayVideoStatisResponseBodyVideoPlayStatisDetailsVideoPlayStatisDetail {
	s.Title = &v
	return s
}

func (s *DescribePlayVideoStatisResponseBodyVideoPlayStatisDetailsVideoPlayStatisDetail) SetUV(v string) *DescribePlayVideoStatisResponseBodyVideoPlayStatisDetailsVideoPlayStatisDetail {
	s.UV = &v
	return s
}

func (s *DescribePlayVideoStatisResponseBodyVideoPlayStatisDetailsVideoPlayStatisDetail) SetVV(v string) *DescribePlayVideoStatisResponseBodyVideoPlayStatisDetailsVideoPlayStatisDetail {
	s.VV = &v
	return s
}

type DescribePlayVideoStatisResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribePlayVideoStatisResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribePlayVideoStatisResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribePlayVideoStatisResponse) GoString() string {
	return s.String()
}

func (s *DescribePlayVideoStatisResponse) SetHeaders(v map[string]*string) *DescribePlayVideoStatisResponse {
	s.Headers = v
	return s
}

func (s *DescribePlayVideoStatisResponse) SetBody(v *DescribePlayVideoStatisResponseBody) *DescribePlayVideoStatisResponse {
	s.Body = v
	return s
}

type DescribeVodAIDataRequest struct {
	AIType    *string `json:"AIType,omitempty" xml:"AIType,omitempty"`
	EndTime   *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OwnerId   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	Region    *string `json:"Region,omitempty" xml:"Region,omitempty"`
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeVodAIDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeVodAIDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeVodAIDataRequest) SetAIType(v string) *DescribeVodAIDataRequest {
	s.AIType = &v
	return s
}

func (s *DescribeVodAIDataRequest) SetEndTime(v string) *DescribeVodAIDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeVodAIDataRequest) SetOwnerId(v int64) *DescribeVodAIDataRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeVodAIDataRequest) SetRegion(v string) *DescribeVodAIDataRequest {
	s.Region = &v
	return s
}

func (s *DescribeVodAIDataRequest) SetStartTime(v string) *DescribeVodAIDataRequest {
	s.StartTime = &v
	return s
}

type DescribeVodAIDataResponseBody struct {
	AIData       *DescribeVodAIDataResponseBodyAIData `json:"AIData,omitempty" xml:"AIData,omitempty" type:"Struct"`
	DataInterval *string                              `json:"DataInterval,omitempty" xml:"DataInterval,omitempty"`
	RequestId    *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeVodAIDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeVodAIDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVodAIDataResponseBody) SetAIData(v *DescribeVodAIDataResponseBodyAIData) *DescribeVodAIDataResponseBody {
	s.AIData = v
	return s
}

func (s *DescribeVodAIDataResponseBody) SetDataInterval(v string) *DescribeVodAIDataResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeVodAIDataResponseBody) SetRequestId(v string) *DescribeVodAIDataResponseBody {
	s.RequestId = &v
	return s
}

type DescribeVodAIDataResponseBodyAIData struct {
	AIDataItem []*DescribeVodAIDataResponseBodyAIDataAIDataItem `json:"AIDataItem,omitempty" xml:"AIDataItem,omitempty" type:"Repeated"`
}

func (s DescribeVodAIDataResponseBodyAIData) String() string {
	return tea.Prettify(s)
}

func (s DescribeVodAIDataResponseBodyAIData) GoString() string {
	return s.String()
}

func (s *DescribeVodAIDataResponseBodyAIData) SetAIDataItem(v []*DescribeVodAIDataResponseBodyAIDataAIDataItem) *DescribeVodAIDataResponseBodyAIData {
	s.AIDataItem = v
	return s
}

type DescribeVodAIDataResponseBodyAIDataAIDataItem struct {
	Data      *DescribeVodAIDataResponseBodyAIDataAIDataItemData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	TimeStamp *string                                            `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
}

func (s DescribeVodAIDataResponseBodyAIDataAIDataItem) String() string {
	return tea.Prettify(s)
}

func (s DescribeVodAIDataResponseBodyAIDataAIDataItem) GoString() string {
	return s.String()
}

func (s *DescribeVodAIDataResponseBodyAIDataAIDataItem) SetData(v *DescribeVodAIDataResponseBodyAIDataAIDataItemData) *DescribeVodAIDataResponseBodyAIDataAIDataItem {
	s.Data = v
	return s
}

func (s *DescribeVodAIDataResponseBodyAIDataAIDataItem) SetTimeStamp(v string) *DescribeVodAIDataResponseBodyAIDataAIDataItem {
	s.TimeStamp = &v
	return s
}

type DescribeVodAIDataResponseBodyAIDataAIDataItemData struct {
	DataItem []*DescribeVodAIDataResponseBodyAIDataAIDataItemDataDataItem `json:"DataItem,omitempty" xml:"DataItem,omitempty" type:"Repeated"`
}

func (s DescribeVodAIDataResponseBodyAIDataAIDataItemData) String() string {
	return tea.Prettify(s)
}

func (s DescribeVodAIDataResponseBodyAIDataAIDataItemData) GoString() string {
	return s.String()
}

func (s *DescribeVodAIDataResponseBodyAIDataAIDataItemData) SetDataItem(v []*DescribeVodAIDataResponseBodyAIDataAIDataItemDataDataItem) *DescribeVodAIDataResponseBodyAIDataAIDataItemData {
	s.DataItem = v
	return s
}

type DescribeVodAIDataResponseBodyAIDataAIDataItemDataDataItem struct {
	Name  *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeVodAIDataResponseBodyAIDataAIDataItemDataDataItem) String() string {
	return tea.Prettify(s)
}

func (s DescribeVodAIDataResponseBodyAIDataAIDataItemDataDataItem) GoString() string {
	return s.String()
}

func (s *DescribeVodAIDataResponseBodyAIDataAIDataItemDataDataItem) SetName(v string) *DescribeVodAIDataResponseBodyAIDataAIDataItemDataDataItem {
	s.Name = &v
	return s
}

func (s *DescribeVodAIDataResponseBodyAIDataAIDataItemDataDataItem) SetValue(v string) *DescribeVodAIDataResponseBodyAIDataAIDataItemDataDataItem {
	s.Value = &v
	return s
}

type DescribeVodAIDataResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeVodAIDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeVodAIDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeVodAIDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeVodAIDataResponse) SetHeaders(v map[string]*string) *DescribeVodAIDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeVodAIDataResponse) SetBody(v *DescribeVodAIDataResponseBody) *DescribeVodAIDataResponse {
	s.Body = v
	return s
}

type DescribeVodCertificateListRequest struct {
	DomainName    *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeVodCertificateListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeVodCertificateListRequest) GoString() string {
	return s.String()
}

func (s *DescribeVodCertificateListRequest) SetDomainName(v string) *DescribeVodCertificateListRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeVodCertificateListRequest) SetOwnerId(v int64) *DescribeVodCertificateListRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeVodCertificateListRequest) SetSecurityToken(v string) *DescribeVodCertificateListRequest {
	s.SecurityToken = &v
	return s
}

type DescribeVodCertificateListResponseBody struct {
	CertificateListModel *DescribeVodCertificateListResponseBodyCertificateListModel `json:"CertificateListModel,omitempty" xml:"CertificateListModel,omitempty" type:"Struct"`
	RequestId            *string                                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeVodCertificateListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeVodCertificateListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVodCertificateListResponseBody) SetCertificateListModel(v *DescribeVodCertificateListResponseBodyCertificateListModel) *DescribeVodCertificateListResponseBody {
	s.CertificateListModel = v
	return s
}

func (s *DescribeVodCertificateListResponseBody) SetRequestId(v string) *DescribeVodCertificateListResponseBody {
	s.RequestId = &v
	return s
}

type DescribeVodCertificateListResponseBodyCertificateListModel struct {
	CertList *DescribeVodCertificateListResponseBodyCertificateListModelCertList `json:"CertList,omitempty" xml:"CertList,omitempty" type:"Struct"`
	Count    *int32                                                              `json:"Count,omitempty" xml:"Count,omitempty"`
}

func (s DescribeVodCertificateListResponseBodyCertificateListModel) String() string {
	return tea.Prettify(s)
}

func (s DescribeVodCertificateListResponseBodyCertificateListModel) GoString() string {
	return s.String()
}

func (s *DescribeVodCertificateListResponseBodyCertificateListModel) SetCertList(v *DescribeVodCertificateListResponseBodyCertificateListModelCertList) *DescribeVodCertificateListResponseBodyCertificateListModel {
	s.CertList = v
	return s
}

func (s *DescribeVodCertificateListResponseBodyCertificateListModel) SetCount(v int32) *DescribeVodCertificateListResponseBodyCertificateListModel {
	s.Count = &v
	return s
}

type DescribeVodCertificateListResponseBodyCertificateListModelCertList struct {
	Cert []*DescribeVodCertificateListResponseBodyCertificateListModelCertListCert `json:"Cert,omitempty" xml:"Cert,omitempty" type:"Repeated"`
}

func (s DescribeVodCertificateListResponseBodyCertificateListModelCertList) String() string {
	return tea.Prettify(s)
}

func (s DescribeVodCertificateListResponseBodyCertificateListModelCertList) GoString() string {
	return s.String()
}

func (s *DescribeVodCertificateListResponseBodyCertificateListModelCertList) SetCert(v []*DescribeVodCertificateListResponseBodyCertificateListModelCertListCert) *DescribeVodCertificateListResponseBodyCertificateListModelCertList {
	s.Cert = v
	return s
}

type DescribeVodCertificateListResponseBodyCertificateListModelCertListCert struct {
	CertId      *int64  `json:"CertId,omitempty" xml:"CertId,omitempty"`
	CertName    *string `json:"CertName,omitempty" xml:"CertName,omitempty"`
	Common      *string `json:"Common,omitempty" xml:"Common,omitempty"`
	Fingerprint *string `json:"Fingerprint,omitempty" xml:"Fingerprint,omitempty"`
	Issuer      *string `json:"Issuer,omitempty" xml:"Issuer,omitempty"`
	LastTime    *int64  `json:"LastTime,omitempty" xml:"LastTime,omitempty"`
}

func (s DescribeVodCertificateListResponseBodyCertificateListModelCertListCert) String() string {
	return tea.Prettify(s)
}

func (s DescribeVodCertificateListResponseBodyCertificateListModelCertListCert) GoString() string {
	return s.String()
}

func (s *DescribeVodCertificateListResponseBodyCertificateListModelCertListCert) SetCertId(v int64) *DescribeVodCertificateListResponseBodyCertificateListModelCertListCert {
	s.CertId = &v
	return s
}

func (s *DescribeVodCertificateListResponseBodyCertificateListModelCertListCert) SetCertName(v string) *DescribeVodCertificateListResponseBodyCertificateListModelCertListCert {
	s.CertName = &v
	return s
}

func (s *DescribeVodCertificateListResponseBodyCertificateListModelCertListCert) SetCommon(v string) *DescribeVodCertificateListResponseBodyCertificateListModelCertListCert {
	s.Common = &v
	return s
}

func (s *DescribeVodCertificateListResponseBodyCertificateListModelCertListCert) SetFingerprint(v string) *DescribeVodCertificateListResponseBodyCertificateListModelCertListCert {
	s.Fingerprint = &v
	return s
}

func (s *DescribeVodCertificateListResponseBodyCertificateListModelCertListCert) SetIssuer(v string) *DescribeVodCertificateListResponseBodyCertificateListModelCertListCert {
	s.Issuer = &v
	return s
}

func (s *DescribeVodCertificateListResponseBodyCertificateListModelCertListCert) SetLastTime(v int64) *DescribeVodCertificateListResponseBodyCertificateListModelCertListCert {
	s.LastTime = &v
	return s
}

type DescribeVodCertificateListResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeVodCertificateListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeVodCertificateListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeVodCertificateListResponse) GoString() string {
	return s.String()
}

func (s *DescribeVodCertificateListResponse) SetHeaders(v map[string]*string) *DescribeVodCertificateListResponse {
	s.Headers = v
	return s
}

func (s *DescribeVodCertificateListResponse) SetBody(v *DescribeVodCertificateListResponseBody) *DescribeVodCertificateListResponse {
	s.Body = v
	return s
}

type DescribeVodDomainBpsDataRequest struct {
	DomainName     *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime        *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Interval       *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	IspNameEn      *string `json:"IspNameEn,omitempty" xml:"IspNameEn,omitempty"`
	LocationNameEn *string `json:"LocationNameEn,omitempty" xml:"LocationNameEn,omitempty"`
	OwnerId        *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	StartTime      *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeVodDomainBpsDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeVodDomainBpsDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainBpsDataRequest) SetDomainName(v string) *DescribeVodDomainBpsDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeVodDomainBpsDataRequest) SetEndTime(v string) *DescribeVodDomainBpsDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeVodDomainBpsDataRequest) SetInterval(v string) *DescribeVodDomainBpsDataRequest {
	s.Interval = &v
	return s
}

func (s *DescribeVodDomainBpsDataRequest) SetIspNameEn(v string) *DescribeVodDomainBpsDataRequest {
	s.IspNameEn = &v
	return s
}

func (s *DescribeVodDomainBpsDataRequest) SetLocationNameEn(v string) *DescribeVodDomainBpsDataRequest {
	s.LocationNameEn = &v
	return s
}

func (s *DescribeVodDomainBpsDataRequest) SetOwnerId(v int64) *DescribeVodDomainBpsDataRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeVodDomainBpsDataRequest) SetStartTime(v string) *DescribeVodDomainBpsDataRequest {
	s.StartTime = &v
	return s
}

type DescribeVodDomainBpsDataResponseBody struct {
	BpsDataPerInterval *DescribeVodDomainBpsDataResponseBodyBpsDataPerInterval `json:"BpsDataPerInterval,omitempty" xml:"BpsDataPerInterval,omitempty" type:"Struct"`
	DataInterval       *string                                                 `json:"DataInterval,omitempty" xml:"DataInterval,omitempty"`
	DomainName         *string                                                 `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime            *string                                                 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	IspNameEn          *string                                                 `json:"IspNameEn,omitempty" xml:"IspNameEn,omitempty"`
	LocationNameEn     *string                                                 `json:"LocationNameEn,omitempty" xml:"LocationNameEn,omitempty"`
	RequestId          *string                                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StartTime          *string                                                 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeVodDomainBpsDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeVodDomainBpsDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainBpsDataResponseBody) SetBpsDataPerInterval(v *DescribeVodDomainBpsDataResponseBodyBpsDataPerInterval) *DescribeVodDomainBpsDataResponseBody {
	s.BpsDataPerInterval = v
	return s
}

func (s *DescribeVodDomainBpsDataResponseBody) SetDataInterval(v string) *DescribeVodDomainBpsDataResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeVodDomainBpsDataResponseBody) SetDomainName(v string) *DescribeVodDomainBpsDataResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeVodDomainBpsDataResponseBody) SetEndTime(v string) *DescribeVodDomainBpsDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeVodDomainBpsDataResponseBody) SetIspNameEn(v string) *DescribeVodDomainBpsDataResponseBody {
	s.IspNameEn = &v
	return s
}

func (s *DescribeVodDomainBpsDataResponseBody) SetLocationNameEn(v string) *DescribeVodDomainBpsDataResponseBody {
	s.LocationNameEn = &v
	return s
}

func (s *DescribeVodDomainBpsDataResponseBody) SetRequestId(v string) *DescribeVodDomainBpsDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVodDomainBpsDataResponseBody) SetStartTime(v string) *DescribeVodDomainBpsDataResponseBody {
	s.StartTime = &v
	return s
}

type DescribeVodDomainBpsDataResponseBodyBpsDataPerInterval struct {
	DataModule []*DescribeVodDomainBpsDataResponseBodyBpsDataPerIntervalDataModule `json:"DataModule,omitempty" xml:"DataModule,omitempty" type:"Repeated"`
}

func (s DescribeVodDomainBpsDataResponseBodyBpsDataPerInterval) String() string {
	return tea.Prettify(s)
}

func (s DescribeVodDomainBpsDataResponseBodyBpsDataPerInterval) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainBpsDataResponseBodyBpsDataPerInterval) SetDataModule(v []*DescribeVodDomainBpsDataResponseBodyBpsDataPerIntervalDataModule) *DescribeVodDomainBpsDataResponseBodyBpsDataPerInterval {
	s.DataModule = v
	return s
}

type DescribeVodDomainBpsDataResponseBodyBpsDataPerIntervalDataModule struct {
	DomesticValue      *string `json:"DomesticValue,omitempty" xml:"DomesticValue,omitempty"`
	HttpsDomesticValue *string `json:"HttpsDomesticValue,omitempty" xml:"HttpsDomesticValue,omitempty"`
	HttpsOverseasValue *string `json:"HttpsOverseasValue,omitempty" xml:"HttpsOverseasValue,omitempty"`
	HttpsValue         *string `json:"HttpsValue,omitempty" xml:"HttpsValue,omitempty"`
	OverseasValue      *string `json:"OverseasValue,omitempty" xml:"OverseasValue,omitempty"`
	TimeStamp          *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	Value              *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeVodDomainBpsDataResponseBodyBpsDataPerIntervalDataModule) String() string {
	return tea.Prettify(s)
}

func (s DescribeVodDomainBpsDataResponseBodyBpsDataPerIntervalDataModule) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainBpsDataResponseBodyBpsDataPerIntervalDataModule) SetDomesticValue(v string) *DescribeVodDomainBpsDataResponseBodyBpsDataPerIntervalDataModule {
	s.DomesticValue = &v
	return s
}

func (s *DescribeVodDomainBpsDataResponseBodyBpsDataPerIntervalDataModule) SetHttpsDomesticValue(v string) *DescribeVodDomainBpsDataResponseBodyBpsDataPerIntervalDataModule {
	s.HttpsDomesticValue = &v
	return s
}

func (s *DescribeVodDomainBpsDataResponseBodyBpsDataPerIntervalDataModule) SetHttpsOverseasValue(v string) *DescribeVodDomainBpsDataResponseBodyBpsDataPerIntervalDataModule {
	s.HttpsOverseasValue = &v
	return s
}

func (s *DescribeVodDomainBpsDataResponseBodyBpsDataPerIntervalDataModule) SetHttpsValue(v string) *DescribeVodDomainBpsDataResponseBodyBpsDataPerIntervalDataModule {
	s.HttpsValue = &v
	return s
}

func (s *DescribeVodDomainBpsDataResponseBodyBpsDataPerIntervalDataModule) SetOverseasValue(v string) *DescribeVodDomainBpsDataResponseBodyBpsDataPerIntervalDataModule {
	s.OverseasValue = &v
	return s
}

func (s *DescribeVodDomainBpsDataResponseBodyBpsDataPerIntervalDataModule) SetTimeStamp(v string) *DescribeVodDomainBpsDataResponseBodyBpsDataPerIntervalDataModule {
	s.TimeStamp = &v
	return s
}

func (s *DescribeVodDomainBpsDataResponseBodyBpsDataPerIntervalDataModule) SetValue(v string) *DescribeVodDomainBpsDataResponseBodyBpsDataPerIntervalDataModule {
	s.Value = &v
	return s
}

type DescribeVodDomainBpsDataResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeVodDomainBpsDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeVodDomainBpsDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeVodDomainBpsDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainBpsDataResponse) SetHeaders(v map[string]*string) *DescribeVodDomainBpsDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeVodDomainBpsDataResponse) SetBody(v *DescribeVodDomainBpsDataResponseBody) *DescribeVodDomainBpsDataResponse {
	s.Body = v
	return s
}

type DescribeVodDomainCertificateInfoRequest struct {
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s DescribeVodDomainCertificateInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeVodDomainCertificateInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainCertificateInfoRequest) SetDomainName(v string) *DescribeVodDomainCertificateInfoRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeVodDomainCertificateInfoRequest) SetOwnerId(v int64) *DescribeVodDomainCertificateInfoRequest {
	s.OwnerId = &v
	return s
}

type DescribeVodDomainCertificateInfoResponseBody struct {
	CertInfos *DescribeVodDomainCertificateInfoResponseBodyCertInfos `json:"CertInfos,omitempty" xml:"CertInfos,omitempty" type:"Struct"`
	RequestId *string                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeVodDomainCertificateInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeVodDomainCertificateInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainCertificateInfoResponseBody) SetCertInfos(v *DescribeVodDomainCertificateInfoResponseBodyCertInfos) *DescribeVodDomainCertificateInfoResponseBody {
	s.CertInfos = v
	return s
}

func (s *DescribeVodDomainCertificateInfoResponseBody) SetRequestId(v string) *DescribeVodDomainCertificateInfoResponseBody {
	s.RequestId = &v
	return s
}

type DescribeVodDomainCertificateInfoResponseBodyCertInfos struct {
	CertInfo []*DescribeVodDomainCertificateInfoResponseBodyCertInfosCertInfo `json:"CertInfo,omitempty" xml:"CertInfo,omitempty" type:"Repeated"`
}

func (s DescribeVodDomainCertificateInfoResponseBodyCertInfos) String() string {
	return tea.Prettify(s)
}

func (s DescribeVodDomainCertificateInfoResponseBodyCertInfos) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainCertificateInfoResponseBodyCertInfos) SetCertInfo(v []*DescribeVodDomainCertificateInfoResponseBodyCertInfosCertInfo) *DescribeVodDomainCertificateInfoResponseBodyCertInfos {
	s.CertInfo = v
	return s
}

type DescribeVodDomainCertificateInfoResponseBodyCertInfosCertInfo struct {
	CertDomainName          *string `json:"CertDomainName,omitempty" xml:"CertDomainName,omitempty"`
	CertExpireTime          *string `json:"CertExpireTime,omitempty" xml:"CertExpireTime,omitempty"`
	CertLife                *string `json:"CertLife,omitempty" xml:"CertLife,omitempty"`
	CertName                *string `json:"CertName,omitempty" xml:"CertName,omitempty"`
	CertOrg                 *string `json:"CertOrg,omitempty" xml:"CertOrg,omitempty"`
	CertType                *string `json:"CertType,omitempty" xml:"CertType,omitempty"`
	DomainName              *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	ServerCertificateStatus *string `json:"ServerCertificateStatus,omitempty" xml:"ServerCertificateStatus,omitempty"`
	Status                  *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeVodDomainCertificateInfoResponseBodyCertInfosCertInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeVodDomainCertificateInfoResponseBodyCertInfosCertInfo) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainCertificateInfoResponseBodyCertInfosCertInfo) SetCertDomainName(v string) *DescribeVodDomainCertificateInfoResponseBodyCertInfosCertInfo {
	s.CertDomainName = &v
	return s
}

func (s *DescribeVodDomainCertificateInfoResponseBodyCertInfosCertInfo) SetCertExpireTime(v string) *DescribeVodDomainCertificateInfoResponseBodyCertInfosCertInfo {
	s.CertExpireTime = &v
	return s
}

func (s *DescribeVodDomainCertificateInfoResponseBodyCertInfosCertInfo) SetCertLife(v string) *DescribeVodDomainCertificateInfoResponseBodyCertInfosCertInfo {
	s.CertLife = &v
	return s
}

func (s *DescribeVodDomainCertificateInfoResponseBodyCertInfosCertInfo) SetCertName(v string) *DescribeVodDomainCertificateInfoResponseBodyCertInfosCertInfo {
	s.CertName = &v
	return s
}

func (s *DescribeVodDomainCertificateInfoResponseBodyCertInfosCertInfo) SetCertOrg(v string) *DescribeVodDomainCertificateInfoResponseBodyCertInfosCertInfo {
	s.CertOrg = &v
	return s
}

func (s *DescribeVodDomainCertificateInfoResponseBodyCertInfosCertInfo) SetCertType(v string) *DescribeVodDomainCertificateInfoResponseBodyCertInfosCertInfo {
	s.CertType = &v
	return s
}

func (s *DescribeVodDomainCertificateInfoResponseBodyCertInfosCertInfo) SetDomainName(v string) *DescribeVodDomainCertificateInfoResponseBodyCertInfosCertInfo {
	s.DomainName = &v
	return s
}

func (s *DescribeVodDomainCertificateInfoResponseBodyCertInfosCertInfo) SetServerCertificateStatus(v string) *DescribeVodDomainCertificateInfoResponseBodyCertInfosCertInfo {
	s.ServerCertificateStatus = &v
	return s
}

func (s *DescribeVodDomainCertificateInfoResponseBodyCertInfosCertInfo) SetStatus(v string) *DescribeVodDomainCertificateInfoResponseBodyCertInfosCertInfo {
	s.Status = &v
	return s
}

type DescribeVodDomainCertificateInfoResponse struct {
	Headers map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeVodDomainCertificateInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeVodDomainCertificateInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeVodDomainCertificateInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainCertificateInfoResponse) SetHeaders(v map[string]*string) *DescribeVodDomainCertificateInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeVodDomainCertificateInfoResponse) SetBody(v *DescribeVodDomainCertificateInfoResponseBody) *DescribeVodDomainCertificateInfoResponse {
	s.Body = v
	return s
}

type DescribeVodDomainConfigsRequest struct {
	DomainName    *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	FunctionNames *string `json:"FunctionNames,omitempty" xml:"FunctionNames,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeVodDomainConfigsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeVodDomainConfigsRequest) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainConfigsRequest) SetDomainName(v string) *DescribeVodDomainConfigsRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeVodDomainConfigsRequest) SetFunctionNames(v string) *DescribeVodDomainConfigsRequest {
	s.FunctionNames = &v
	return s
}

func (s *DescribeVodDomainConfigsRequest) SetOwnerId(v int64) *DescribeVodDomainConfigsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeVodDomainConfigsRequest) SetSecurityToken(v string) *DescribeVodDomainConfigsRequest {
	s.SecurityToken = &v
	return s
}

type DescribeVodDomainConfigsResponseBody struct {
	DomainConfigs *DescribeVodDomainConfigsResponseBodyDomainConfigs `json:"DomainConfigs,omitempty" xml:"DomainConfigs,omitempty" type:"Struct"`
	RequestId     *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeVodDomainConfigsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeVodDomainConfigsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainConfigsResponseBody) SetDomainConfigs(v *DescribeVodDomainConfigsResponseBodyDomainConfigs) *DescribeVodDomainConfigsResponseBody {
	s.DomainConfigs = v
	return s
}

func (s *DescribeVodDomainConfigsResponseBody) SetRequestId(v string) *DescribeVodDomainConfigsResponseBody {
	s.RequestId = &v
	return s
}

type DescribeVodDomainConfigsResponseBodyDomainConfigs struct {
	DomainConfig []*DescribeVodDomainConfigsResponseBodyDomainConfigsDomainConfig `json:"DomainConfig,omitempty" xml:"DomainConfig,omitempty" type:"Repeated"`
}

func (s DescribeVodDomainConfigsResponseBodyDomainConfigs) String() string {
	return tea.Prettify(s)
}

func (s DescribeVodDomainConfigsResponseBodyDomainConfigs) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainConfigsResponseBodyDomainConfigs) SetDomainConfig(v []*DescribeVodDomainConfigsResponseBodyDomainConfigsDomainConfig) *DescribeVodDomainConfigsResponseBodyDomainConfigs {
	s.DomainConfig = v
	return s
}

type DescribeVodDomainConfigsResponseBodyDomainConfigsDomainConfig struct {
	ConfigId     *string                                                                    `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	FunctionArgs *DescribeVodDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgs `json:"FunctionArgs,omitempty" xml:"FunctionArgs,omitempty" type:"Struct"`
	FunctionName *string                                                                    `json:"FunctionName,omitempty" xml:"FunctionName,omitempty"`
	Status       *string                                                                    `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeVodDomainConfigsResponseBodyDomainConfigsDomainConfig) String() string {
	return tea.Prettify(s)
}

func (s DescribeVodDomainConfigsResponseBodyDomainConfigsDomainConfig) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainConfigsResponseBodyDomainConfigsDomainConfig) SetConfigId(v string) *DescribeVodDomainConfigsResponseBodyDomainConfigsDomainConfig {
	s.ConfigId = &v
	return s
}

func (s *DescribeVodDomainConfigsResponseBodyDomainConfigsDomainConfig) SetFunctionArgs(v *DescribeVodDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgs) *DescribeVodDomainConfigsResponseBodyDomainConfigsDomainConfig {
	s.FunctionArgs = v
	return s
}

func (s *DescribeVodDomainConfigsResponseBodyDomainConfigsDomainConfig) SetFunctionName(v string) *DescribeVodDomainConfigsResponseBodyDomainConfigsDomainConfig {
	s.FunctionName = &v
	return s
}

func (s *DescribeVodDomainConfigsResponseBodyDomainConfigsDomainConfig) SetStatus(v string) *DescribeVodDomainConfigsResponseBodyDomainConfigsDomainConfig {
	s.Status = &v
	return s
}

type DescribeVodDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgs struct {
	FunctionArg []*DescribeVodDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgsFunctionArg `json:"FunctionArg,omitempty" xml:"FunctionArg,omitempty" type:"Repeated"`
}

func (s DescribeVodDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgs) String() string {
	return tea.Prettify(s)
}

func (s DescribeVodDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgs) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgs) SetFunctionArg(v []*DescribeVodDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgsFunctionArg) *DescribeVodDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgs {
	s.FunctionArg = v
	return s
}

type DescribeVodDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgsFunctionArg struct {
	ArgName  *string `json:"ArgName,omitempty" xml:"ArgName,omitempty"`
	ArgValue *string `json:"ArgValue,omitempty" xml:"ArgValue,omitempty"`
}

func (s DescribeVodDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgsFunctionArg) String() string {
	return tea.Prettify(s)
}

func (s DescribeVodDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgsFunctionArg) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgsFunctionArg) SetArgName(v string) *DescribeVodDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgsFunctionArg {
	s.ArgName = &v
	return s
}

func (s *DescribeVodDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgsFunctionArg) SetArgValue(v string) *DescribeVodDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgsFunctionArg {
	s.ArgValue = &v
	return s
}

type DescribeVodDomainConfigsResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeVodDomainConfigsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeVodDomainConfigsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeVodDomainConfigsResponse) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainConfigsResponse) SetHeaders(v map[string]*string) *DescribeVodDomainConfigsResponse {
	s.Headers = v
	return s
}

func (s *DescribeVodDomainConfigsResponse) SetBody(v *DescribeVodDomainConfigsResponseBody) *DescribeVodDomainConfigsResponse {
	s.Body = v
	return s
}

type DescribeVodDomainDetailRequest struct {
	DomainName    *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeVodDomainDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeVodDomainDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainDetailRequest) SetDomainName(v string) *DescribeVodDomainDetailRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeVodDomainDetailRequest) SetOwnerId(v int64) *DescribeVodDomainDetailRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeVodDomainDetailRequest) SetSecurityToken(v string) *DescribeVodDomainDetailRequest {
	s.SecurityToken = &v
	return s
}

type DescribeVodDomainDetailResponseBody struct {
	DomainDetail *DescribeVodDomainDetailResponseBodyDomainDetail `json:"DomainDetail,omitempty" xml:"DomainDetail,omitempty" type:"Struct"`
	RequestId    *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeVodDomainDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeVodDomainDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainDetailResponseBody) SetDomainDetail(v *DescribeVodDomainDetailResponseBodyDomainDetail) *DescribeVodDomainDetailResponseBody {
	s.DomainDetail = v
	return s
}

func (s *DescribeVodDomainDetailResponseBody) SetRequestId(v string) *DescribeVodDomainDetailResponseBody {
	s.RequestId = &v
	return s
}

type DescribeVodDomainDetailResponseBodyDomainDetail struct {
	CertName     *string                                                 `json:"CertName,omitempty" xml:"CertName,omitempty"`
	Cname        *string                                                 `json:"Cname,omitempty" xml:"Cname,omitempty"`
	Description  *string                                                 `json:"Description,omitempty" xml:"Description,omitempty"`
	DomainName   *string                                                 `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	DomainStatus *string                                                 `json:"DomainStatus,omitempty" xml:"DomainStatus,omitempty"`
	GmtCreated   *string                                                 `json:"GmtCreated,omitempty" xml:"GmtCreated,omitempty"`
	GmtModified  *string                                                 `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	SSLProtocol  *string                                                 `json:"SSLProtocol,omitempty" xml:"SSLProtocol,omitempty"`
	SSLPub       *string                                                 `json:"SSLPub,omitempty" xml:"SSLPub,omitempty"`
	Scope        *string                                                 `json:"Scope,omitempty" xml:"Scope,omitempty"`
	Sources      *DescribeVodDomainDetailResponseBodyDomainDetailSources `json:"Sources,omitempty" xml:"Sources,omitempty" type:"Struct"`
	Weight       *string                                                 `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s DescribeVodDomainDetailResponseBodyDomainDetail) String() string {
	return tea.Prettify(s)
}

func (s DescribeVodDomainDetailResponseBodyDomainDetail) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainDetailResponseBodyDomainDetail) SetCertName(v string) *DescribeVodDomainDetailResponseBodyDomainDetail {
	s.CertName = &v
	return s
}

func (s *DescribeVodDomainDetailResponseBodyDomainDetail) SetCname(v string) *DescribeVodDomainDetailResponseBodyDomainDetail {
	s.Cname = &v
	return s
}

func (s *DescribeVodDomainDetailResponseBodyDomainDetail) SetDescription(v string) *DescribeVodDomainDetailResponseBodyDomainDetail {
	s.Description = &v
	return s
}

func (s *DescribeVodDomainDetailResponseBodyDomainDetail) SetDomainName(v string) *DescribeVodDomainDetailResponseBodyDomainDetail {
	s.DomainName = &v
	return s
}

func (s *DescribeVodDomainDetailResponseBodyDomainDetail) SetDomainStatus(v string) *DescribeVodDomainDetailResponseBodyDomainDetail {
	s.DomainStatus = &v
	return s
}

func (s *DescribeVodDomainDetailResponseBodyDomainDetail) SetGmtCreated(v string) *DescribeVodDomainDetailResponseBodyDomainDetail {
	s.GmtCreated = &v
	return s
}

func (s *DescribeVodDomainDetailResponseBodyDomainDetail) SetGmtModified(v string) *DescribeVodDomainDetailResponseBodyDomainDetail {
	s.GmtModified = &v
	return s
}

func (s *DescribeVodDomainDetailResponseBodyDomainDetail) SetSSLProtocol(v string) *DescribeVodDomainDetailResponseBodyDomainDetail {
	s.SSLProtocol = &v
	return s
}

func (s *DescribeVodDomainDetailResponseBodyDomainDetail) SetSSLPub(v string) *DescribeVodDomainDetailResponseBodyDomainDetail {
	s.SSLPub = &v
	return s
}

func (s *DescribeVodDomainDetailResponseBodyDomainDetail) SetScope(v string) *DescribeVodDomainDetailResponseBodyDomainDetail {
	s.Scope = &v
	return s
}

func (s *DescribeVodDomainDetailResponseBodyDomainDetail) SetSources(v *DescribeVodDomainDetailResponseBodyDomainDetailSources) *DescribeVodDomainDetailResponseBodyDomainDetail {
	s.Sources = v
	return s
}

func (s *DescribeVodDomainDetailResponseBodyDomainDetail) SetWeight(v string) *DescribeVodDomainDetailResponseBodyDomainDetail {
	s.Weight = &v
	return s
}

type DescribeVodDomainDetailResponseBodyDomainDetailSources struct {
	Source []*DescribeVodDomainDetailResponseBodyDomainDetailSourcesSource `json:"Source,omitempty" xml:"Source,omitempty" type:"Repeated"`
}

func (s DescribeVodDomainDetailResponseBodyDomainDetailSources) String() string {
	return tea.Prettify(s)
}

func (s DescribeVodDomainDetailResponseBodyDomainDetailSources) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainDetailResponseBodyDomainDetailSources) SetSource(v []*DescribeVodDomainDetailResponseBodyDomainDetailSourcesSource) *DescribeVodDomainDetailResponseBodyDomainDetailSources {
	s.Source = v
	return s
}

type DescribeVodDomainDetailResponseBodyDomainDetailSourcesSource struct {
	Content  *string `json:"Content,omitempty" xml:"Content,omitempty"`
	Enabled  *string `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	Port     *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
	Priority *string `json:"Priority,omitempty" xml:"Priority,omitempty"`
	Type     *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeVodDomainDetailResponseBodyDomainDetailSourcesSource) String() string {
	return tea.Prettify(s)
}

func (s DescribeVodDomainDetailResponseBodyDomainDetailSourcesSource) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainDetailResponseBodyDomainDetailSourcesSource) SetContent(v string) *DescribeVodDomainDetailResponseBodyDomainDetailSourcesSource {
	s.Content = &v
	return s
}

func (s *DescribeVodDomainDetailResponseBodyDomainDetailSourcesSource) SetEnabled(v string) *DescribeVodDomainDetailResponseBodyDomainDetailSourcesSource {
	s.Enabled = &v
	return s
}

func (s *DescribeVodDomainDetailResponseBodyDomainDetailSourcesSource) SetPort(v int32) *DescribeVodDomainDetailResponseBodyDomainDetailSourcesSource {
	s.Port = &v
	return s
}

func (s *DescribeVodDomainDetailResponseBodyDomainDetailSourcesSource) SetPriority(v string) *DescribeVodDomainDetailResponseBodyDomainDetailSourcesSource {
	s.Priority = &v
	return s
}

func (s *DescribeVodDomainDetailResponseBodyDomainDetailSourcesSource) SetType(v string) *DescribeVodDomainDetailResponseBodyDomainDetailSourcesSource {
	s.Type = &v
	return s
}

type DescribeVodDomainDetailResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeVodDomainDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeVodDomainDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeVodDomainDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainDetailResponse) SetHeaders(v map[string]*string) *DescribeVodDomainDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribeVodDomainDetailResponse) SetBody(v *DescribeVodDomainDetailResponseBody) *DescribeVodDomainDetailResponse {
	s.Body = v
	return s
}

type DescribeVodDomainLogRequest struct {
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime    *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNumber *int64  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	StartTime  *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeVodDomainLogRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeVodDomainLogRequest) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainLogRequest) SetDomainName(v string) *DescribeVodDomainLogRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeVodDomainLogRequest) SetEndTime(v string) *DescribeVodDomainLogRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeVodDomainLogRequest) SetOwnerId(v int64) *DescribeVodDomainLogRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeVodDomainLogRequest) SetPageNumber(v int64) *DescribeVodDomainLogRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeVodDomainLogRequest) SetPageSize(v int64) *DescribeVodDomainLogRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeVodDomainLogRequest) SetStartTime(v string) *DescribeVodDomainLogRequest {
	s.StartTime = &v
	return s
}

type DescribeVodDomainLogResponseBody struct {
	DomainLogDetails *DescribeVodDomainLogResponseBodyDomainLogDetails `json:"DomainLogDetails,omitempty" xml:"DomainLogDetails,omitempty" type:"Struct"`
	RequestId        *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeVodDomainLogResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeVodDomainLogResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainLogResponseBody) SetDomainLogDetails(v *DescribeVodDomainLogResponseBodyDomainLogDetails) *DescribeVodDomainLogResponseBody {
	s.DomainLogDetails = v
	return s
}

func (s *DescribeVodDomainLogResponseBody) SetRequestId(v string) *DescribeVodDomainLogResponseBody {
	s.RequestId = &v
	return s
}

type DescribeVodDomainLogResponseBodyDomainLogDetails struct {
	DomainLogDetail []*DescribeVodDomainLogResponseBodyDomainLogDetailsDomainLogDetail `json:"DomainLogDetail,omitempty" xml:"DomainLogDetail,omitempty" type:"Repeated"`
}

func (s DescribeVodDomainLogResponseBodyDomainLogDetails) String() string {
	return tea.Prettify(s)
}

func (s DescribeVodDomainLogResponseBodyDomainLogDetails) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainLogResponseBodyDomainLogDetails) SetDomainLogDetail(v []*DescribeVodDomainLogResponseBodyDomainLogDetailsDomainLogDetail) *DescribeVodDomainLogResponseBodyDomainLogDetails {
	s.DomainLogDetail = v
	return s
}

type DescribeVodDomainLogResponseBodyDomainLogDetailsDomainLogDetail struct {
	DomainName *string                                                                   `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	LogCount   *int64                                                                    `json:"LogCount,omitempty" xml:"LogCount,omitempty"`
	LogInfos   *DescribeVodDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfos  `json:"LogInfos,omitempty" xml:"LogInfos,omitempty" type:"Struct"`
	PageInfos  *DescribeVodDomainLogResponseBodyDomainLogDetailsDomainLogDetailPageInfos `json:"PageInfos,omitempty" xml:"PageInfos,omitempty" type:"Struct"`
}

func (s DescribeVodDomainLogResponseBodyDomainLogDetailsDomainLogDetail) String() string {
	return tea.Prettify(s)
}

func (s DescribeVodDomainLogResponseBodyDomainLogDetailsDomainLogDetail) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainLogResponseBodyDomainLogDetailsDomainLogDetail) SetDomainName(v string) *DescribeVodDomainLogResponseBodyDomainLogDetailsDomainLogDetail {
	s.DomainName = &v
	return s
}

func (s *DescribeVodDomainLogResponseBodyDomainLogDetailsDomainLogDetail) SetLogCount(v int64) *DescribeVodDomainLogResponseBodyDomainLogDetailsDomainLogDetail {
	s.LogCount = &v
	return s
}

func (s *DescribeVodDomainLogResponseBodyDomainLogDetailsDomainLogDetail) SetLogInfos(v *DescribeVodDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfos) *DescribeVodDomainLogResponseBodyDomainLogDetailsDomainLogDetail {
	s.LogInfos = v
	return s
}

func (s *DescribeVodDomainLogResponseBodyDomainLogDetailsDomainLogDetail) SetPageInfos(v *DescribeVodDomainLogResponseBodyDomainLogDetailsDomainLogDetailPageInfos) *DescribeVodDomainLogResponseBodyDomainLogDetailsDomainLogDetail {
	s.PageInfos = v
	return s
}

type DescribeVodDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfos struct {
	LogInfoDetail []*DescribeVodDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail `json:"LogInfoDetail,omitempty" xml:"LogInfoDetail,omitempty" type:"Repeated"`
}

func (s DescribeVodDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfos) String() string {
	return tea.Prettify(s)
}

func (s DescribeVodDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfos) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfos) SetLogInfoDetail(v []*DescribeVodDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail) *DescribeVodDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfos {
	s.LogInfoDetail = v
	return s
}

type DescribeVodDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail struct {
	EndTime   *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	LogName   *string `json:"LogName,omitempty" xml:"LogName,omitempty"`
	LogPath   *string `json:"LogPath,omitempty" xml:"LogPath,omitempty"`
	LogSize   *int64  `json:"LogSize,omitempty" xml:"LogSize,omitempty"`
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeVodDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail) String() string {
	return tea.Prettify(s)
}

func (s DescribeVodDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail) SetEndTime(v string) *DescribeVodDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail {
	s.EndTime = &v
	return s
}

func (s *DescribeVodDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail) SetLogName(v string) *DescribeVodDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail {
	s.LogName = &v
	return s
}

func (s *DescribeVodDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail) SetLogPath(v string) *DescribeVodDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail {
	s.LogPath = &v
	return s
}

func (s *DescribeVodDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail) SetLogSize(v int64) *DescribeVodDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail {
	s.LogSize = &v
	return s
}

func (s *DescribeVodDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail) SetStartTime(v string) *DescribeVodDomainLogResponseBodyDomainLogDetailsDomainLogDetailLogInfosLogInfoDetail {
	s.StartTime = &v
	return s
}

type DescribeVodDomainLogResponseBodyDomainLogDetailsDomainLogDetailPageInfos struct {
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Total      *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeVodDomainLogResponseBodyDomainLogDetailsDomainLogDetailPageInfos) String() string {
	return tea.Prettify(s)
}

func (s DescribeVodDomainLogResponseBodyDomainLogDetailsDomainLogDetailPageInfos) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainLogResponseBodyDomainLogDetailsDomainLogDetailPageInfos) SetPageNumber(v int64) *DescribeVodDomainLogResponseBodyDomainLogDetailsDomainLogDetailPageInfos {
	s.PageNumber = &v
	return s
}

func (s *DescribeVodDomainLogResponseBodyDomainLogDetailsDomainLogDetailPageInfos) SetPageSize(v int64) *DescribeVodDomainLogResponseBodyDomainLogDetailsDomainLogDetailPageInfos {
	s.PageSize = &v
	return s
}

func (s *DescribeVodDomainLogResponseBodyDomainLogDetailsDomainLogDetailPageInfos) SetTotal(v int64) *DescribeVodDomainLogResponseBodyDomainLogDetailsDomainLogDetailPageInfos {
	s.Total = &v
	return s
}

type DescribeVodDomainLogResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeVodDomainLogResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeVodDomainLogResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeVodDomainLogResponse) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainLogResponse) SetHeaders(v map[string]*string) *DescribeVodDomainLogResponse {
	s.Headers = v
	return s
}

func (s *DescribeVodDomainLogResponse) SetBody(v *DescribeVodDomainLogResponseBody) *DescribeVodDomainLogResponse {
	s.Body = v
	return s
}

type DescribeVodDomainRealtimeLogDeliveryRequest struct {
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s DescribeVodDomainRealtimeLogDeliveryRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeVodDomainRealtimeLogDeliveryRequest) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainRealtimeLogDeliveryRequest) SetDomainName(v string) *DescribeVodDomainRealtimeLogDeliveryRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeVodDomainRealtimeLogDeliveryRequest) SetOwnerId(v int64) *DescribeVodDomainRealtimeLogDeliveryRequest {
	s.OwnerId = &v
	return s
}

type DescribeVodDomainRealtimeLogDeliveryResponseBody struct {
	Logstore  *string `json:"Logstore,omitempty" xml:"Logstore,omitempty"`
	Project   *string `json:"Project,omitempty" xml:"Project,omitempty"`
	Region    *string `json:"Region,omitempty" xml:"Region,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Status    *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeVodDomainRealtimeLogDeliveryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeVodDomainRealtimeLogDeliveryResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainRealtimeLogDeliveryResponseBody) SetLogstore(v string) *DescribeVodDomainRealtimeLogDeliveryResponseBody {
	s.Logstore = &v
	return s
}

func (s *DescribeVodDomainRealtimeLogDeliveryResponseBody) SetProject(v string) *DescribeVodDomainRealtimeLogDeliveryResponseBody {
	s.Project = &v
	return s
}

func (s *DescribeVodDomainRealtimeLogDeliveryResponseBody) SetRegion(v string) *DescribeVodDomainRealtimeLogDeliveryResponseBody {
	s.Region = &v
	return s
}

func (s *DescribeVodDomainRealtimeLogDeliveryResponseBody) SetRequestId(v string) *DescribeVodDomainRealtimeLogDeliveryResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVodDomainRealtimeLogDeliveryResponseBody) SetStatus(v string) *DescribeVodDomainRealtimeLogDeliveryResponseBody {
	s.Status = &v
	return s
}

type DescribeVodDomainRealtimeLogDeliveryResponse struct {
	Headers map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeVodDomainRealtimeLogDeliveryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeVodDomainRealtimeLogDeliveryResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeVodDomainRealtimeLogDeliveryResponse) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainRealtimeLogDeliveryResponse) SetHeaders(v map[string]*string) *DescribeVodDomainRealtimeLogDeliveryResponse {
	s.Headers = v
	return s
}

func (s *DescribeVodDomainRealtimeLogDeliveryResponse) SetBody(v *DescribeVodDomainRealtimeLogDeliveryResponseBody) *DescribeVodDomainRealtimeLogDeliveryResponse {
	s.Body = v
	return s
}

type DescribeVodDomainTrafficDataRequest struct {
	DomainName     *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime        *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Interval       *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	IspNameEn      *string `json:"IspNameEn,omitempty" xml:"IspNameEn,omitempty"`
	LocationNameEn *string `json:"LocationNameEn,omitempty" xml:"LocationNameEn,omitempty"`
	OwnerId        *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	StartTime      *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeVodDomainTrafficDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeVodDomainTrafficDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainTrafficDataRequest) SetDomainName(v string) *DescribeVodDomainTrafficDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeVodDomainTrafficDataRequest) SetEndTime(v string) *DescribeVodDomainTrafficDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeVodDomainTrafficDataRequest) SetInterval(v string) *DescribeVodDomainTrafficDataRequest {
	s.Interval = &v
	return s
}

func (s *DescribeVodDomainTrafficDataRequest) SetIspNameEn(v string) *DescribeVodDomainTrafficDataRequest {
	s.IspNameEn = &v
	return s
}

func (s *DescribeVodDomainTrafficDataRequest) SetLocationNameEn(v string) *DescribeVodDomainTrafficDataRequest {
	s.LocationNameEn = &v
	return s
}

func (s *DescribeVodDomainTrafficDataRequest) SetOwnerId(v int64) *DescribeVodDomainTrafficDataRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeVodDomainTrafficDataRequest) SetStartTime(v string) *DescribeVodDomainTrafficDataRequest {
	s.StartTime = &v
	return s
}

type DescribeVodDomainTrafficDataResponseBody struct {
	DataInterval           *string                                                         `json:"DataInterval,omitempty" xml:"DataInterval,omitempty"`
	DomainName             *string                                                         `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime                *string                                                         `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	RequestId              *string                                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StartTime              *string                                                         `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	TrafficDataPerInterval *DescribeVodDomainTrafficDataResponseBodyTrafficDataPerInterval `json:"TrafficDataPerInterval,omitempty" xml:"TrafficDataPerInterval,omitempty" type:"Struct"`
}

func (s DescribeVodDomainTrafficDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeVodDomainTrafficDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainTrafficDataResponseBody) SetDataInterval(v string) *DescribeVodDomainTrafficDataResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeVodDomainTrafficDataResponseBody) SetDomainName(v string) *DescribeVodDomainTrafficDataResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeVodDomainTrafficDataResponseBody) SetEndTime(v string) *DescribeVodDomainTrafficDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeVodDomainTrafficDataResponseBody) SetRequestId(v string) *DescribeVodDomainTrafficDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVodDomainTrafficDataResponseBody) SetStartTime(v string) *DescribeVodDomainTrafficDataResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeVodDomainTrafficDataResponseBody) SetTrafficDataPerInterval(v *DescribeVodDomainTrafficDataResponseBodyTrafficDataPerInterval) *DescribeVodDomainTrafficDataResponseBody {
	s.TrafficDataPerInterval = v
	return s
}

type DescribeVodDomainTrafficDataResponseBodyTrafficDataPerInterval struct {
	DataModule []*DescribeVodDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule `json:"DataModule,omitempty" xml:"DataModule,omitempty" type:"Repeated"`
}

func (s DescribeVodDomainTrafficDataResponseBodyTrafficDataPerInterval) String() string {
	return tea.Prettify(s)
}

func (s DescribeVodDomainTrafficDataResponseBodyTrafficDataPerInterval) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainTrafficDataResponseBodyTrafficDataPerInterval) SetDataModule(v []*DescribeVodDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule) *DescribeVodDomainTrafficDataResponseBodyTrafficDataPerInterval {
	s.DataModule = v
	return s
}

type DescribeVodDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule struct {
	DomesticValue      *string `json:"DomesticValue,omitempty" xml:"DomesticValue,omitempty"`
	HttpsDomesticValue *string `json:"HttpsDomesticValue,omitempty" xml:"HttpsDomesticValue,omitempty"`
	HttpsOverseasValue *string `json:"HttpsOverseasValue,omitempty" xml:"HttpsOverseasValue,omitempty"`
	HttpsValue         *string `json:"HttpsValue,omitempty" xml:"HttpsValue,omitempty"`
	OverseasValue      *string `json:"OverseasValue,omitempty" xml:"OverseasValue,omitempty"`
	TimeStamp          *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	Value              *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeVodDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule) String() string {
	return tea.Prettify(s)
}

func (s DescribeVodDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule) SetDomesticValue(v string) *DescribeVodDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule {
	s.DomesticValue = &v
	return s
}

func (s *DescribeVodDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule) SetHttpsDomesticValue(v string) *DescribeVodDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule {
	s.HttpsDomesticValue = &v
	return s
}

func (s *DescribeVodDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule) SetHttpsOverseasValue(v string) *DescribeVodDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule {
	s.HttpsOverseasValue = &v
	return s
}

func (s *DescribeVodDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule) SetHttpsValue(v string) *DescribeVodDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule {
	s.HttpsValue = &v
	return s
}

func (s *DescribeVodDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule) SetOverseasValue(v string) *DescribeVodDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule {
	s.OverseasValue = &v
	return s
}

func (s *DescribeVodDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule) SetTimeStamp(v string) *DescribeVodDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule {
	s.TimeStamp = &v
	return s
}

func (s *DescribeVodDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule) SetValue(v string) *DescribeVodDomainTrafficDataResponseBodyTrafficDataPerIntervalDataModule {
	s.Value = &v
	return s
}

type DescribeVodDomainTrafficDataResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeVodDomainTrafficDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeVodDomainTrafficDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeVodDomainTrafficDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainTrafficDataResponse) SetHeaders(v map[string]*string) *DescribeVodDomainTrafficDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeVodDomainTrafficDataResponse) SetBody(v *DescribeVodDomainTrafficDataResponseBody) *DescribeVodDomainTrafficDataResponse {
	s.Body = v
	return s
}

type DescribeVodDomainUsageDataRequest struct {
	Area       *string `json:"Area,omitempty" xml:"Area,omitempty"`
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime    *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Field      *string `json:"Field,omitempty" xml:"Field,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	StartTime  *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Type       *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeVodDomainUsageDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeVodDomainUsageDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainUsageDataRequest) SetArea(v string) *DescribeVodDomainUsageDataRequest {
	s.Area = &v
	return s
}

func (s *DescribeVodDomainUsageDataRequest) SetDomainName(v string) *DescribeVodDomainUsageDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeVodDomainUsageDataRequest) SetEndTime(v string) *DescribeVodDomainUsageDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeVodDomainUsageDataRequest) SetField(v string) *DescribeVodDomainUsageDataRequest {
	s.Field = &v
	return s
}

func (s *DescribeVodDomainUsageDataRequest) SetOwnerId(v int64) *DescribeVodDomainUsageDataRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeVodDomainUsageDataRequest) SetStartTime(v string) *DescribeVodDomainUsageDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeVodDomainUsageDataRequest) SetType(v string) *DescribeVodDomainUsageDataRequest {
	s.Type = &v
	return s
}

type DescribeVodDomainUsageDataResponseBody struct {
	Area                 *string                                                     `json:"Area,omitempty" xml:"Area,omitempty"`
	DataInterval         *string                                                     `json:"DataInterval,omitempty" xml:"DataInterval,omitempty"`
	DomainName           *string                                                     `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime              *string                                                     `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	RequestId            *string                                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StartTime            *string                                                     `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Type                 *string                                                     `json:"Type,omitempty" xml:"Type,omitempty"`
	UsageDataPerInterval *DescribeVodDomainUsageDataResponseBodyUsageDataPerInterval `json:"UsageDataPerInterval,omitempty" xml:"UsageDataPerInterval,omitempty" type:"Struct"`
}

func (s DescribeVodDomainUsageDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeVodDomainUsageDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainUsageDataResponseBody) SetArea(v string) *DescribeVodDomainUsageDataResponseBody {
	s.Area = &v
	return s
}

func (s *DescribeVodDomainUsageDataResponseBody) SetDataInterval(v string) *DescribeVodDomainUsageDataResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeVodDomainUsageDataResponseBody) SetDomainName(v string) *DescribeVodDomainUsageDataResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeVodDomainUsageDataResponseBody) SetEndTime(v string) *DescribeVodDomainUsageDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeVodDomainUsageDataResponseBody) SetRequestId(v string) *DescribeVodDomainUsageDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVodDomainUsageDataResponseBody) SetStartTime(v string) *DescribeVodDomainUsageDataResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeVodDomainUsageDataResponseBody) SetType(v string) *DescribeVodDomainUsageDataResponseBody {
	s.Type = &v
	return s
}

func (s *DescribeVodDomainUsageDataResponseBody) SetUsageDataPerInterval(v *DescribeVodDomainUsageDataResponseBodyUsageDataPerInterval) *DescribeVodDomainUsageDataResponseBody {
	s.UsageDataPerInterval = v
	return s
}

type DescribeVodDomainUsageDataResponseBodyUsageDataPerInterval struct {
	DataModule []*DescribeVodDomainUsageDataResponseBodyUsageDataPerIntervalDataModule `json:"DataModule,omitempty" xml:"DataModule,omitempty" type:"Repeated"`
}

func (s DescribeVodDomainUsageDataResponseBodyUsageDataPerInterval) String() string {
	return tea.Prettify(s)
}

func (s DescribeVodDomainUsageDataResponseBodyUsageDataPerInterval) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainUsageDataResponseBodyUsageDataPerInterval) SetDataModule(v []*DescribeVodDomainUsageDataResponseBodyUsageDataPerIntervalDataModule) *DescribeVodDomainUsageDataResponseBodyUsageDataPerInterval {
	s.DataModule = v
	return s
}

type DescribeVodDomainUsageDataResponseBodyUsageDataPerIntervalDataModule struct {
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	Value     *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeVodDomainUsageDataResponseBodyUsageDataPerIntervalDataModule) String() string {
	return tea.Prettify(s)
}

func (s DescribeVodDomainUsageDataResponseBodyUsageDataPerIntervalDataModule) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainUsageDataResponseBodyUsageDataPerIntervalDataModule) SetTimeStamp(v string) *DescribeVodDomainUsageDataResponseBodyUsageDataPerIntervalDataModule {
	s.TimeStamp = &v
	return s
}

func (s *DescribeVodDomainUsageDataResponseBodyUsageDataPerIntervalDataModule) SetValue(v string) *DescribeVodDomainUsageDataResponseBodyUsageDataPerIntervalDataModule {
	s.Value = &v
	return s
}

type DescribeVodDomainUsageDataResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeVodDomainUsageDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeVodDomainUsageDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeVodDomainUsageDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainUsageDataResponse) SetHeaders(v map[string]*string) *DescribeVodDomainUsageDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeVodDomainUsageDataResponse) SetBody(v *DescribeVodDomainUsageDataResponseBody) *DescribeVodDomainUsageDataResponse {
	s.Body = v
	return s
}

type DescribeVodRefreshQuotaRequest struct {
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeVodRefreshQuotaRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeVodRefreshQuotaRequest) GoString() string {
	return s.String()
}

func (s *DescribeVodRefreshQuotaRequest) SetOwnerId(v int64) *DescribeVodRefreshQuotaRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeVodRefreshQuotaRequest) SetSecurityToken(v string) *DescribeVodRefreshQuotaRequest {
	s.SecurityToken = &v
	return s
}

type DescribeVodRefreshQuotaResponseBody struct {
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

func (s DescribeVodRefreshQuotaResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeVodRefreshQuotaResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVodRefreshQuotaResponseBody) SetBlockQuota(v string) *DescribeVodRefreshQuotaResponseBody {
	s.BlockQuota = &v
	return s
}

func (s *DescribeVodRefreshQuotaResponseBody) SetDirQuota(v string) *DescribeVodRefreshQuotaResponseBody {
	s.DirQuota = &v
	return s
}

func (s *DescribeVodRefreshQuotaResponseBody) SetDirRemain(v string) *DescribeVodRefreshQuotaResponseBody {
	s.DirRemain = &v
	return s
}

func (s *DescribeVodRefreshQuotaResponseBody) SetPreloadQuota(v string) *DescribeVodRefreshQuotaResponseBody {
	s.PreloadQuota = &v
	return s
}

func (s *DescribeVodRefreshQuotaResponseBody) SetPreloadRemain(v string) *DescribeVodRefreshQuotaResponseBody {
	s.PreloadRemain = &v
	return s
}

func (s *DescribeVodRefreshQuotaResponseBody) SetRequestId(v string) *DescribeVodRefreshQuotaResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVodRefreshQuotaResponseBody) SetUrlQuota(v string) *DescribeVodRefreshQuotaResponseBody {
	s.UrlQuota = &v
	return s
}

func (s *DescribeVodRefreshQuotaResponseBody) SetUrlRemain(v string) *DescribeVodRefreshQuotaResponseBody {
	s.UrlRemain = &v
	return s
}

func (s *DescribeVodRefreshQuotaResponseBody) SetBlockRemain(v string) *DescribeVodRefreshQuotaResponseBody {
	s.BlockRemain = &v
	return s
}

type DescribeVodRefreshQuotaResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeVodRefreshQuotaResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeVodRefreshQuotaResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeVodRefreshQuotaResponse) GoString() string {
	return s.String()
}

func (s *DescribeVodRefreshQuotaResponse) SetHeaders(v map[string]*string) *DescribeVodRefreshQuotaResponse {
	s.Headers = v
	return s
}

func (s *DescribeVodRefreshQuotaResponse) SetBody(v *DescribeVodRefreshQuotaResponseBody) *DescribeVodRefreshQuotaResponse {
	s.Body = v
	return s
}

type DescribeVodRefreshTasksRequest struct {
	DomainName    *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime       *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	ObjectPath    *string `json:"ObjectPath,omitempty" xml:"ObjectPath,omitempty"`
	ObjectType    *string `json:"ObjectType,omitempty" xml:"ObjectType,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNumber    *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize      *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	StartTime     *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Status        *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TaskId        *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DescribeVodRefreshTasksRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeVodRefreshTasksRequest) GoString() string {
	return s.String()
}

func (s *DescribeVodRefreshTasksRequest) SetDomainName(v string) *DescribeVodRefreshTasksRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeVodRefreshTasksRequest) SetEndTime(v string) *DescribeVodRefreshTasksRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeVodRefreshTasksRequest) SetObjectPath(v string) *DescribeVodRefreshTasksRequest {
	s.ObjectPath = &v
	return s
}

func (s *DescribeVodRefreshTasksRequest) SetObjectType(v string) *DescribeVodRefreshTasksRequest {
	s.ObjectType = &v
	return s
}

func (s *DescribeVodRefreshTasksRequest) SetOwnerId(v int64) *DescribeVodRefreshTasksRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeVodRefreshTasksRequest) SetPageNumber(v int32) *DescribeVodRefreshTasksRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeVodRefreshTasksRequest) SetPageSize(v int32) *DescribeVodRefreshTasksRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeVodRefreshTasksRequest) SetSecurityToken(v string) *DescribeVodRefreshTasksRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeVodRefreshTasksRequest) SetStartTime(v string) *DescribeVodRefreshTasksRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeVodRefreshTasksRequest) SetStatus(v string) *DescribeVodRefreshTasksRequest {
	s.Status = &v
	return s
}

func (s *DescribeVodRefreshTasksRequest) SetTaskId(v string) *DescribeVodRefreshTasksRequest {
	s.TaskId = &v
	return s
}

type DescribeVodRefreshTasksResponseBody struct {
	PageNumber *int64                                    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int64                                    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Tasks      *DescribeVodRefreshTasksResponseBodyTasks `json:"Tasks,omitempty" xml:"Tasks,omitempty" type:"Struct"`
	TotalCount *int64                                    `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeVodRefreshTasksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeVodRefreshTasksResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVodRefreshTasksResponseBody) SetPageNumber(v int64) *DescribeVodRefreshTasksResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeVodRefreshTasksResponseBody) SetPageSize(v int64) *DescribeVodRefreshTasksResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeVodRefreshTasksResponseBody) SetRequestId(v string) *DescribeVodRefreshTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVodRefreshTasksResponseBody) SetTasks(v *DescribeVodRefreshTasksResponseBodyTasks) *DescribeVodRefreshTasksResponseBody {
	s.Tasks = v
	return s
}

func (s *DescribeVodRefreshTasksResponseBody) SetTotalCount(v int64) *DescribeVodRefreshTasksResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeVodRefreshTasksResponseBodyTasks struct {
	Task []*DescribeVodRefreshTasksResponseBodyTasksTask `json:"Task,omitempty" xml:"Task,omitempty" type:"Repeated"`
}

func (s DescribeVodRefreshTasksResponseBodyTasks) String() string {
	return tea.Prettify(s)
}

func (s DescribeVodRefreshTasksResponseBodyTasks) GoString() string {
	return s.String()
}

func (s *DescribeVodRefreshTasksResponseBodyTasks) SetTask(v []*DescribeVodRefreshTasksResponseBodyTasksTask) *DescribeVodRefreshTasksResponseBodyTasks {
	s.Task = v
	return s
}

type DescribeVodRefreshTasksResponseBodyTasksTask struct {
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
	ObjectPath   *string `json:"ObjectPath,omitempty" xml:"ObjectPath,omitempty"`
	ObjectType   *string `json:"ObjectType,omitempty" xml:"ObjectType,omitempty"`
	Process      *string `json:"Process,omitempty" xml:"Process,omitempty"`
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TaskId       *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DescribeVodRefreshTasksResponseBodyTasksTask) String() string {
	return tea.Prettify(s)
}

func (s DescribeVodRefreshTasksResponseBodyTasksTask) GoString() string {
	return s.String()
}

func (s *DescribeVodRefreshTasksResponseBodyTasksTask) SetCreationTime(v string) *DescribeVodRefreshTasksResponseBodyTasksTask {
	s.CreationTime = &v
	return s
}

func (s *DescribeVodRefreshTasksResponseBodyTasksTask) SetDescription(v string) *DescribeVodRefreshTasksResponseBodyTasksTask {
	s.Description = &v
	return s
}

func (s *DescribeVodRefreshTasksResponseBodyTasksTask) SetObjectPath(v string) *DescribeVodRefreshTasksResponseBodyTasksTask {
	s.ObjectPath = &v
	return s
}

func (s *DescribeVodRefreshTasksResponseBodyTasksTask) SetObjectType(v string) *DescribeVodRefreshTasksResponseBodyTasksTask {
	s.ObjectType = &v
	return s
}

func (s *DescribeVodRefreshTasksResponseBodyTasksTask) SetProcess(v string) *DescribeVodRefreshTasksResponseBodyTasksTask {
	s.Process = &v
	return s
}

func (s *DescribeVodRefreshTasksResponseBodyTasksTask) SetStatus(v string) *DescribeVodRefreshTasksResponseBodyTasksTask {
	s.Status = &v
	return s
}

func (s *DescribeVodRefreshTasksResponseBodyTasksTask) SetTaskId(v string) *DescribeVodRefreshTasksResponseBodyTasksTask {
	s.TaskId = &v
	return s
}

type DescribeVodRefreshTasksResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeVodRefreshTasksResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeVodRefreshTasksResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeVodRefreshTasksResponse) GoString() string {
	return s.String()
}

func (s *DescribeVodRefreshTasksResponse) SetHeaders(v map[string]*string) *DescribeVodRefreshTasksResponse {
	s.Headers = v
	return s
}

func (s *DescribeVodRefreshTasksResponse) SetBody(v *DescribeVodRefreshTasksResponseBody) *DescribeVodRefreshTasksResponse {
	s.Body = v
	return s
}

type DescribeVodStorageDataRequest struct {
	EndTime     *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OwnerId     *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	Region      *string `json:"Region,omitempty" xml:"Region,omitempty"`
	StartTime   *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Storage     *string `json:"Storage,omitempty" xml:"Storage,omitempty"`
	StorageType *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
}

func (s DescribeVodStorageDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeVodStorageDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeVodStorageDataRequest) SetEndTime(v string) *DescribeVodStorageDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeVodStorageDataRequest) SetOwnerId(v int64) *DescribeVodStorageDataRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeVodStorageDataRequest) SetRegion(v string) *DescribeVodStorageDataRequest {
	s.Region = &v
	return s
}

func (s *DescribeVodStorageDataRequest) SetStartTime(v string) *DescribeVodStorageDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeVodStorageDataRequest) SetStorage(v string) *DescribeVodStorageDataRequest {
	s.Storage = &v
	return s
}

func (s *DescribeVodStorageDataRequest) SetStorageType(v string) *DescribeVodStorageDataRequest {
	s.StorageType = &v
	return s
}

type DescribeVodStorageDataResponseBody struct {
	DataInterval *string                                        `json:"DataInterval,omitempty" xml:"DataInterval,omitempty"`
	RequestId    *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StorageData  *DescribeVodStorageDataResponseBodyStorageData `json:"StorageData,omitempty" xml:"StorageData,omitempty" type:"Struct"`
}

func (s DescribeVodStorageDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeVodStorageDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVodStorageDataResponseBody) SetDataInterval(v string) *DescribeVodStorageDataResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeVodStorageDataResponseBody) SetRequestId(v string) *DescribeVodStorageDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVodStorageDataResponseBody) SetStorageData(v *DescribeVodStorageDataResponseBodyStorageData) *DescribeVodStorageDataResponseBody {
	s.StorageData = v
	return s
}

type DescribeVodStorageDataResponseBodyStorageData struct {
	StorageDataItem []*DescribeVodStorageDataResponseBodyStorageDataStorageDataItem `json:"StorageDataItem,omitempty" xml:"StorageDataItem,omitempty" type:"Repeated"`
}

func (s DescribeVodStorageDataResponseBodyStorageData) String() string {
	return tea.Prettify(s)
}

func (s DescribeVodStorageDataResponseBodyStorageData) GoString() string {
	return s.String()
}

func (s *DescribeVodStorageDataResponseBodyStorageData) SetStorageDataItem(v []*DescribeVodStorageDataResponseBodyStorageDataStorageDataItem) *DescribeVodStorageDataResponseBodyStorageData {
	s.StorageDataItem = v
	return s
}

type DescribeVodStorageDataResponseBodyStorageDataStorageDataItem struct {
	NetworkOut         *string `json:"NetworkOut,omitempty" xml:"NetworkOut,omitempty"`
	StorageUtilization *string `json:"StorageUtilization,omitempty" xml:"StorageUtilization,omitempty"`
	TimeStamp          *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
}

func (s DescribeVodStorageDataResponseBodyStorageDataStorageDataItem) String() string {
	return tea.Prettify(s)
}

func (s DescribeVodStorageDataResponseBodyStorageDataStorageDataItem) GoString() string {
	return s.String()
}

func (s *DescribeVodStorageDataResponseBodyStorageDataStorageDataItem) SetNetworkOut(v string) *DescribeVodStorageDataResponseBodyStorageDataStorageDataItem {
	s.NetworkOut = &v
	return s
}

func (s *DescribeVodStorageDataResponseBodyStorageDataStorageDataItem) SetStorageUtilization(v string) *DescribeVodStorageDataResponseBodyStorageDataStorageDataItem {
	s.StorageUtilization = &v
	return s
}

func (s *DescribeVodStorageDataResponseBodyStorageDataStorageDataItem) SetTimeStamp(v string) *DescribeVodStorageDataResponseBodyStorageDataStorageDataItem {
	s.TimeStamp = &v
	return s
}

type DescribeVodStorageDataResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeVodStorageDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeVodStorageDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeVodStorageDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeVodStorageDataResponse) SetHeaders(v map[string]*string) *DescribeVodStorageDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeVodStorageDataResponse) SetBody(v *DescribeVodStorageDataResponseBody) *DescribeVodStorageDataResponse {
	s.Body = v
	return s
}

type DescribeVodTagResourcesRequest struct {
	OwnerId      *int64                               `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceId   []*string                            `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	ResourceType *string                              `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Tag          []*DescribeVodTagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeVodTagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeVodTagResourcesRequest) GoString() string {
	return s.String()
}

func (s *DescribeVodTagResourcesRequest) SetOwnerId(v int64) *DescribeVodTagResourcesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeVodTagResourcesRequest) SetResourceId(v []*string) *DescribeVodTagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *DescribeVodTagResourcesRequest) SetResourceType(v string) *DescribeVodTagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *DescribeVodTagResourcesRequest) SetTag(v []*DescribeVodTagResourcesRequestTag) *DescribeVodTagResourcesRequest {
	s.Tag = v
	return s
}

type DescribeVodTagResourcesRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeVodTagResourcesRequestTag) String() string {
	return tea.Prettify(s)
}

func (s DescribeVodTagResourcesRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeVodTagResourcesRequestTag) SetKey(v string) *DescribeVodTagResourcesRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeVodTagResourcesRequestTag) SetValue(v string) *DescribeVodTagResourcesRequestTag {
	s.Value = &v
	return s
}

type DescribeVodTagResourcesResponseBody struct {
	RequestId    *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TagResources []*DescribeVodTagResourcesResponseBodyTagResources `json:"TagResources,omitempty" xml:"TagResources,omitempty" type:"Repeated"`
}

func (s DescribeVodTagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeVodTagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVodTagResourcesResponseBody) SetRequestId(v string) *DescribeVodTagResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVodTagResourcesResponseBody) SetTagResources(v []*DescribeVodTagResourcesResponseBodyTagResources) *DescribeVodTagResourcesResponseBody {
	s.TagResources = v
	return s
}

type DescribeVodTagResourcesResponseBodyTagResources struct {
	ResourceId *string                                               `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	Tag        []*DescribeVodTagResourcesResponseBodyTagResourcesTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeVodTagResourcesResponseBodyTagResources) String() string {
	return tea.Prettify(s)
}

func (s DescribeVodTagResourcesResponseBodyTagResources) GoString() string {
	return s.String()
}

func (s *DescribeVodTagResourcesResponseBodyTagResources) SetResourceId(v string) *DescribeVodTagResourcesResponseBodyTagResources {
	s.ResourceId = &v
	return s
}

func (s *DescribeVodTagResourcesResponseBodyTagResources) SetTag(v []*DescribeVodTagResourcesResponseBodyTagResourcesTag) *DescribeVodTagResourcesResponseBodyTagResources {
	s.Tag = v
	return s
}

type DescribeVodTagResourcesResponseBodyTagResourcesTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeVodTagResourcesResponseBodyTagResourcesTag) String() string {
	return tea.Prettify(s)
}

func (s DescribeVodTagResourcesResponseBodyTagResourcesTag) GoString() string {
	return s.String()
}

func (s *DescribeVodTagResourcesResponseBodyTagResourcesTag) SetKey(v string) *DescribeVodTagResourcesResponseBodyTagResourcesTag {
	s.Key = &v
	return s
}

func (s *DescribeVodTagResourcesResponseBodyTagResourcesTag) SetValue(v string) *DescribeVodTagResourcesResponseBodyTagResourcesTag {
	s.Value = &v
	return s
}

type DescribeVodTagResourcesResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeVodTagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeVodTagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeVodTagResourcesResponse) GoString() string {
	return s.String()
}

func (s *DescribeVodTagResourcesResponse) SetHeaders(v map[string]*string) *DescribeVodTagResourcesResponse {
	s.Headers = v
	return s
}

func (s *DescribeVodTagResourcesResponse) SetBody(v *DescribeVodTagResourcesResponseBody) *DescribeVodTagResourcesResponse {
	s.Body = v
	return s
}

type DescribeVodTranscodeDataRequest struct {
	EndTime       *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Interval      *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	Region        *string `json:"Region,omitempty" xml:"Region,omitempty"`
	Specification *string `json:"Specification,omitempty" xml:"Specification,omitempty"`
	StartTime     *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Storage       *string `json:"Storage,omitempty" xml:"Storage,omitempty"`
}

func (s DescribeVodTranscodeDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeVodTranscodeDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeVodTranscodeDataRequest) SetEndTime(v string) *DescribeVodTranscodeDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeVodTranscodeDataRequest) SetInterval(v string) *DescribeVodTranscodeDataRequest {
	s.Interval = &v
	return s
}

func (s *DescribeVodTranscodeDataRequest) SetOwnerId(v int64) *DescribeVodTranscodeDataRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeVodTranscodeDataRequest) SetRegion(v string) *DescribeVodTranscodeDataRequest {
	s.Region = &v
	return s
}

func (s *DescribeVodTranscodeDataRequest) SetSpecification(v string) *DescribeVodTranscodeDataRequest {
	s.Specification = &v
	return s
}

func (s *DescribeVodTranscodeDataRequest) SetStartTime(v string) *DescribeVodTranscodeDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeVodTranscodeDataRequest) SetStorage(v string) *DescribeVodTranscodeDataRequest {
	s.Storage = &v
	return s
}

type DescribeVodTranscodeDataResponseBody struct {
	DataInterval  *string                                            `json:"DataInterval,omitempty" xml:"DataInterval,omitempty"`
	RequestId     *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TranscodeData *DescribeVodTranscodeDataResponseBodyTranscodeData `json:"TranscodeData,omitempty" xml:"TranscodeData,omitempty" type:"Struct"`
}

func (s DescribeVodTranscodeDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeVodTranscodeDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVodTranscodeDataResponseBody) SetDataInterval(v string) *DescribeVodTranscodeDataResponseBody {
	s.DataInterval = &v
	return s
}

func (s *DescribeVodTranscodeDataResponseBody) SetRequestId(v string) *DescribeVodTranscodeDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVodTranscodeDataResponseBody) SetTranscodeData(v *DescribeVodTranscodeDataResponseBodyTranscodeData) *DescribeVodTranscodeDataResponseBody {
	s.TranscodeData = v
	return s
}

type DescribeVodTranscodeDataResponseBodyTranscodeData struct {
	TranscodeDataItem []*DescribeVodTranscodeDataResponseBodyTranscodeDataTranscodeDataItem `json:"TranscodeDataItem,omitempty" xml:"TranscodeDataItem,omitempty" type:"Repeated"`
}

func (s DescribeVodTranscodeDataResponseBodyTranscodeData) String() string {
	return tea.Prettify(s)
}

func (s DescribeVodTranscodeDataResponseBodyTranscodeData) GoString() string {
	return s.String()
}

func (s *DescribeVodTranscodeDataResponseBodyTranscodeData) SetTranscodeDataItem(v []*DescribeVodTranscodeDataResponseBodyTranscodeDataTranscodeDataItem) *DescribeVodTranscodeDataResponseBodyTranscodeData {
	s.TranscodeDataItem = v
	return s
}

type DescribeVodTranscodeDataResponseBodyTranscodeDataTranscodeDataItem struct {
	Data      *DescribeVodTranscodeDataResponseBodyTranscodeDataTranscodeDataItemData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	TimeStamp *string                                                                 `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
}

func (s DescribeVodTranscodeDataResponseBodyTranscodeDataTranscodeDataItem) String() string {
	return tea.Prettify(s)
}

func (s DescribeVodTranscodeDataResponseBodyTranscodeDataTranscodeDataItem) GoString() string {
	return s.String()
}

func (s *DescribeVodTranscodeDataResponseBodyTranscodeDataTranscodeDataItem) SetData(v *DescribeVodTranscodeDataResponseBodyTranscodeDataTranscodeDataItemData) *DescribeVodTranscodeDataResponseBodyTranscodeDataTranscodeDataItem {
	s.Data = v
	return s
}

func (s *DescribeVodTranscodeDataResponseBodyTranscodeDataTranscodeDataItem) SetTimeStamp(v string) *DescribeVodTranscodeDataResponseBodyTranscodeDataTranscodeDataItem {
	s.TimeStamp = &v
	return s
}

type DescribeVodTranscodeDataResponseBodyTranscodeDataTranscodeDataItemData struct {
	DataItem []*DescribeVodTranscodeDataResponseBodyTranscodeDataTranscodeDataItemDataDataItem `json:"DataItem,omitempty" xml:"DataItem,omitempty" type:"Repeated"`
}

func (s DescribeVodTranscodeDataResponseBodyTranscodeDataTranscodeDataItemData) String() string {
	return tea.Prettify(s)
}

func (s DescribeVodTranscodeDataResponseBodyTranscodeDataTranscodeDataItemData) GoString() string {
	return s.String()
}

func (s *DescribeVodTranscodeDataResponseBodyTranscodeDataTranscodeDataItemData) SetDataItem(v []*DescribeVodTranscodeDataResponseBodyTranscodeDataTranscodeDataItemDataDataItem) *DescribeVodTranscodeDataResponseBodyTranscodeDataTranscodeDataItemData {
	s.DataItem = v
	return s
}

type DescribeVodTranscodeDataResponseBodyTranscodeDataTranscodeDataItemDataDataItem struct {
	Name  *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeVodTranscodeDataResponseBodyTranscodeDataTranscodeDataItemDataDataItem) String() string {
	return tea.Prettify(s)
}

func (s DescribeVodTranscodeDataResponseBodyTranscodeDataTranscodeDataItemDataDataItem) GoString() string {
	return s.String()
}

func (s *DescribeVodTranscodeDataResponseBodyTranscodeDataTranscodeDataItemDataDataItem) SetName(v string) *DescribeVodTranscodeDataResponseBodyTranscodeDataTranscodeDataItemDataDataItem {
	s.Name = &v
	return s
}

func (s *DescribeVodTranscodeDataResponseBodyTranscodeDataTranscodeDataItemDataDataItem) SetValue(v string) *DescribeVodTranscodeDataResponseBodyTranscodeDataTranscodeDataItemDataDataItem {
	s.Value = &v
	return s
}

type DescribeVodTranscodeDataResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeVodTranscodeDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeVodTranscodeDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeVodTranscodeDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeVodTranscodeDataResponse) SetHeaders(v map[string]*string) *DescribeVodTranscodeDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeVodTranscodeDataResponse) SetBody(v *DescribeVodTranscodeDataResponseBody) *DescribeVodTranscodeDataResponse {
	s.Body = v
	return s
}

type DescribeVodUserDomainsRequest struct {
	DomainName       *string                             `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	DomainSearchType *string                             `json:"DomainSearchType,omitempty" xml:"DomainSearchType,omitempty"`
	DomainStatus     *string                             `json:"DomainStatus,omitempty" xml:"DomainStatus,omitempty"`
	OwnerId          *int64                              `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNumber       *int32                              `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize         *int32                              `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SecurityToken    *string                             `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	Tag              []*DescribeVodUserDomainsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeVodUserDomainsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeVodUserDomainsRequest) GoString() string {
	return s.String()
}

func (s *DescribeVodUserDomainsRequest) SetDomainName(v string) *DescribeVodUserDomainsRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeVodUserDomainsRequest) SetDomainSearchType(v string) *DescribeVodUserDomainsRequest {
	s.DomainSearchType = &v
	return s
}

func (s *DescribeVodUserDomainsRequest) SetDomainStatus(v string) *DescribeVodUserDomainsRequest {
	s.DomainStatus = &v
	return s
}

func (s *DescribeVodUserDomainsRequest) SetOwnerId(v int64) *DescribeVodUserDomainsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeVodUserDomainsRequest) SetPageNumber(v int32) *DescribeVodUserDomainsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeVodUserDomainsRequest) SetPageSize(v int32) *DescribeVodUserDomainsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeVodUserDomainsRequest) SetSecurityToken(v string) *DescribeVodUserDomainsRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeVodUserDomainsRequest) SetTag(v []*DescribeVodUserDomainsRequestTag) *DescribeVodUserDomainsRequest {
	s.Tag = v
	return s
}

type DescribeVodUserDomainsRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeVodUserDomainsRequestTag) String() string {
	return tea.Prettify(s)
}

func (s DescribeVodUserDomainsRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeVodUserDomainsRequestTag) SetKey(v string) *DescribeVodUserDomainsRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeVodUserDomainsRequestTag) SetValue(v string) *DescribeVodUserDomainsRequestTag {
	s.Value = &v
	return s
}

type DescribeVodUserDomainsResponseBody struct {
	Domains    *DescribeVodUserDomainsResponseBodyDomains `json:"Domains,omitempty" xml:"Domains,omitempty" type:"Struct"`
	PageNumber *int64                                     `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int64                                     `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int64                                     `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeVodUserDomainsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeVodUserDomainsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVodUserDomainsResponseBody) SetDomains(v *DescribeVodUserDomainsResponseBodyDomains) *DescribeVodUserDomainsResponseBody {
	s.Domains = v
	return s
}

func (s *DescribeVodUserDomainsResponseBody) SetPageNumber(v int64) *DescribeVodUserDomainsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeVodUserDomainsResponseBody) SetPageSize(v int64) *DescribeVodUserDomainsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeVodUserDomainsResponseBody) SetRequestId(v string) *DescribeVodUserDomainsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVodUserDomainsResponseBody) SetTotalCount(v int64) *DescribeVodUserDomainsResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeVodUserDomainsResponseBodyDomains struct {
	PageData []*DescribeVodUserDomainsResponseBodyDomainsPageData `json:"PageData,omitempty" xml:"PageData,omitempty" type:"Repeated"`
}

func (s DescribeVodUserDomainsResponseBodyDomains) String() string {
	return tea.Prettify(s)
}

func (s DescribeVodUserDomainsResponseBodyDomains) GoString() string {
	return s.String()
}

func (s *DescribeVodUserDomainsResponseBodyDomains) SetPageData(v []*DescribeVodUserDomainsResponseBodyDomainsPageData) *DescribeVodUserDomainsResponseBodyDomains {
	s.PageData = v
	return s
}

type DescribeVodUserDomainsResponseBodyDomainsPageData struct {
	Cname        *string                                                   `json:"Cname,omitempty" xml:"Cname,omitempty"`
	Description  *string                                                   `json:"Description,omitempty" xml:"Description,omitempty"`
	DomainName   *string                                                   `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	DomainStatus *string                                                   `json:"DomainStatus,omitempty" xml:"DomainStatus,omitempty"`
	GmtCreated   *string                                                   `json:"GmtCreated,omitempty" xml:"GmtCreated,omitempty"`
	GmtModified  *string                                                   `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	Sandbox      *string                                                   `json:"Sandbox,omitempty" xml:"Sandbox,omitempty"`
	Sources      *DescribeVodUserDomainsResponseBodyDomainsPageDataSources `json:"Sources,omitempty" xml:"Sources,omitempty" type:"Struct"`
	SslProtocol  *string                                                   `json:"SslProtocol,omitempty" xml:"SslProtocol,omitempty"`
}

func (s DescribeVodUserDomainsResponseBodyDomainsPageData) String() string {
	return tea.Prettify(s)
}

func (s DescribeVodUserDomainsResponseBodyDomainsPageData) GoString() string {
	return s.String()
}

func (s *DescribeVodUserDomainsResponseBodyDomainsPageData) SetCname(v string) *DescribeVodUserDomainsResponseBodyDomainsPageData {
	s.Cname = &v
	return s
}

func (s *DescribeVodUserDomainsResponseBodyDomainsPageData) SetDescription(v string) *DescribeVodUserDomainsResponseBodyDomainsPageData {
	s.Description = &v
	return s
}

func (s *DescribeVodUserDomainsResponseBodyDomainsPageData) SetDomainName(v string) *DescribeVodUserDomainsResponseBodyDomainsPageData {
	s.DomainName = &v
	return s
}

func (s *DescribeVodUserDomainsResponseBodyDomainsPageData) SetDomainStatus(v string) *DescribeVodUserDomainsResponseBodyDomainsPageData {
	s.DomainStatus = &v
	return s
}

func (s *DescribeVodUserDomainsResponseBodyDomainsPageData) SetGmtCreated(v string) *DescribeVodUserDomainsResponseBodyDomainsPageData {
	s.GmtCreated = &v
	return s
}

func (s *DescribeVodUserDomainsResponseBodyDomainsPageData) SetGmtModified(v string) *DescribeVodUserDomainsResponseBodyDomainsPageData {
	s.GmtModified = &v
	return s
}

func (s *DescribeVodUserDomainsResponseBodyDomainsPageData) SetSandbox(v string) *DescribeVodUserDomainsResponseBodyDomainsPageData {
	s.Sandbox = &v
	return s
}

func (s *DescribeVodUserDomainsResponseBodyDomainsPageData) SetSources(v *DescribeVodUserDomainsResponseBodyDomainsPageDataSources) *DescribeVodUserDomainsResponseBodyDomainsPageData {
	s.Sources = v
	return s
}

func (s *DescribeVodUserDomainsResponseBodyDomainsPageData) SetSslProtocol(v string) *DescribeVodUserDomainsResponseBodyDomainsPageData {
	s.SslProtocol = &v
	return s
}

type DescribeVodUserDomainsResponseBodyDomainsPageDataSources struct {
	Source []*DescribeVodUserDomainsResponseBodyDomainsPageDataSourcesSource `json:"Source,omitempty" xml:"Source,omitempty" type:"Repeated"`
}

func (s DescribeVodUserDomainsResponseBodyDomainsPageDataSources) String() string {
	return tea.Prettify(s)
}

func (s DescribeVodUserDomainsResponseBodyDomainsPageDataSources) GoString() string {
	return s.String()
}

func (s *DescribeVodUserDomainsResponseBodyDomainsPageDataSources) SetSource(v []*DescribeVodUserDomainsResponseBodyDomainsPageDataSourcesSource) *DescribeVodUserDomainsResponseBodyDomainsPageDataSources {
	s.Source = v
	return s
}

type DescribeVodUserDomainsResponseBodyDomainsPageDataSourcesSource struct {
	Content  *string `json:"Content,omitempty" xml:"Content,omitempty"`
	Port     *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
	Priority *string `json:"Priority,omitempty" xml:"Priority,omitempty"`
	Type     *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeVodUserDomainsResponseBodyDomainsPageDataSourcesSource) String() string {
	return tea.Prettify(s)
}

func (s DescribeVodUserDomainsResponseBodyDomainsPageDataSourcesSource) GoString() string {
	return s.String()
}

func (s *DescribeVodUserDomainsResponseBodyDomainsPageDataSourcesSource) SetContent(v string) *DescribeVodUserDomainsResponseBodyDomainsPageDataSourcesSource {
	s.Content = &v
	return s
}

func (s *DescribeVodUserDomainsResponseBodyDomainsPageDataSourcesSource) SetPort(v int32) *DescribeVodUserDomainsResponseBodyDomainsPageDataSourcesSource {
	s.Port = &v
	return s
}

func (s *DescribeVodUserDomainsResponseBodyDomainsPageDataSourcesSource) SetPriority(v string) *DescribeVodUserDomainsResponseBodyDomainsPageDataSourcesSource {
	s.Priority = &v
	return s
}

func (s *DescribeVodUserDomainsResponseBodyDomainsPageDataSourcesSource) SetType(v string) *DescribeVodUserDomainsResponseBodyDomainsPageDataSourcesSource {
	s.Type = &v
	return s
}

type DescribeVodUserDomainsResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeVodUserDomainsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeVodUserDomainsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeVodUserDomainsResponse) GoString() string {
	return s.String()
}

func (s *DescribeVodUserDomainsResponse) SetHeaders(v map[string]*string) *DescribeVodUserDomainsResponse {
	s.Headers = v
	return s
}

func (s *DescribeVodUserDomainsResponse) SetBody(v *DescribeVodUserDomainsResponseBody) *DescribeVodUserDomainsResponse {
	s.Body = v
	return s
}

type DescribeVodUserTagsRequest struct {
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s DescribeVodUserTagsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeVodUserTagsRequest) GoString() string {
	return s.String()
}

func (s *DescribeVodUserTagsRequest) SetOwnerId(v int64) *DescribeVodUserTagsRequest {
	s.OwnerId = &v
	return s
}

type DescribeVodUserTagsResponseBody struct {
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Tags      []*DescribeVodUserTagsResponseBodyTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s DescribeVodUserTagsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeVodUserTagsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVodUserTagsResponseBody) SetRequestId(v string) *DescribeVodUserTagsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVodUserTagsResponseBody) SetTags(v []*DescribeVodUserTagsResponseBodyTags) *DescribeVodUserTagsResponseBody {
	s.Tags = v
	return s
}

type DescribeVodUserTagsResponseBodyTags struct {
	Key   *string   `json:"Key,omitempty" xml:"Key,omitempty"`
	Value []*string `json:"Value,omitempty" xml:"Value,omitempty" type:"Repeated"`
}

func (s DescribeVodUserTagsResponseBodyTags) String() string {
	return tea.Prettify(s)
}

func (s DescribeVodUserTagsResponseBodyTags) GoString() string {
	return s.String()
}

func (s *DescribeVodUserTagsResponseBodyTags) SetKey(v string) *DescribeVodUserTagsResponseBodyTags {
	s.Key = &v
	return s
}

func (s *DescribeVodUserTagsResponseBodyTags) SetValue(v []*string) *DescribeVodUserTagsResponseBodyTags {
	s.Value = v
	return s
}

type DescribeVodUserTagsResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeVodUserTagsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeVodUserTagsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeVodUserTagsResponse) GoString() string {
	return s.String()
}

func (s *DescribeVodUserTagsResponse) SetHeaders(v map[string]*string) *DescribeVodUserTagsResponse {
	s.Headers = v
	return s
}

func (s *DescribeVodUserTagsResponse) SetBody(v *DescribeVodUserTagsResponseBody) *DescribeVodUserTagsResponse {
	s.Body = v
	return s
}

type DescribeVodVerifyContentRequest struct {
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s DescribeVodVerifyContentRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeVodVerifyContentRequest) GoString() string {
	return s.String()
}

func (s *DescribeVodVerifyContentRequest) SetDomainName(v string) *DescribeVodVerifyContentRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeVodVerifyContentRequest) SetOwnerId(v int64) *DescribeVodVerifyContentRequest {
	s.OwnerId = &v
	return s
}

type DescribeVodVerifyContentResponseBody struct {
	Content   *string `json:"Content,omitempty" xml:"Content,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeVodVerifyContentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeVodVerifyContentResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVodVerifyContentResponseBody) SetContent(v string) *DescribeVodVerifyContentResponseBody {
	s.Content = &v
	return s
}

func (s *DescribeVodVerifyContentResponseBody) SetRequestId(v string) *DescribeVodVerifyContentResponseBody {
	s.RequestId = &v
	return s
}

type DescribeVodVerifyContentResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeVodVerifyContentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeVodVerifyContentResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeVodVerifyContentResponse) GoString() string {
	return s.String()
}

func (s *DescribeVodVerifyContentResponse) SetHeaders(v map[string]*string) *DescribeVodVerifyContentResponse {
	s.Headers = v
	return s
}

func (s *DescribeVodVerifyContentResponse) SetBody(v *DescribeVodVerifyContentResponseBody) *DescribeVodVerifyContentResponse {
	s.Body = v
	return s
}

type DetachAppPolicyFromIdentityRequest struct {
	AppId        *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	IdentityName *string `json:"IdentityName,omitempty" xml:"IdentityName,omitempty"`
	IdentityType *string `json:"IdentityType,omitempty" xml:"IdentityType,omitempty"`
	PolicyNames  *string `json:"PolicyNames,omitempty" xml:"PolicyNames,omitempty"`
}

func (s DetachAppPolicyFromIdentityRequest) String() string {
	return tea.Prettify(s)
}

func (s DetachAppPolicyFromIdentityRequest) GoString() string {
	return s.String()
}

func (s *DetachAppPolicyFromIdentityRequest) SetAppId(v string) *DetachAppPolicyFromIdentityRequest {
	s.AppId = &v
	return s
}

func (s *DetachAppPolicyFromIdentityRequest) SetIdentityName(v string) *DetachAppPolicyFromIdentityRequest {
	s.IdentityName = &v
	return s
}

func (s *DetachAppPolicyFromIdentityRequest) SetIdentityType(v string) *DetachAppPolicyFromIdentityRequest {
	s.IdentityType = &v
	return s
}

func (s *DetachAppPolicyFromIdentityRequest) SetPolicyNames(v string) *DetachAppPolicyFromIdentityRequest {
	s.PolicyNames = &v
	return s
}

type DetachAppPolicyFromIdentityResponseBody struct {
	FailedPolicyNames   []*string `json:"FailedPolicyNames,omitempty" xml:"FailedPolicyNames,omitempty" type:"Repeated"`
	NonExistPolicyNames []*string `json:"NonExistPolicyNames,omitempty" xml:"NonExistPolicyNames,omitempty" type:"Repeated"`
	RequestId           *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DetachAppPolicyFromIdentityResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DetachAppPolicyFromIdentityResponseBody) GoString() string {
	return s.String()
}

func (s *DetachAppPolicyFromIdentityResponseBody) SetFailedPolicyNames(v []*string) *DetachAppPolicyFromIdentityResponseBody {
	s.FailedPolicyNames = v
	return s
}

func (s *DetachAppPolicyFromIdentityResponseBody) SetNonExistPolicyNames(v []*string) *DetachAppPolicyFromIdentityResponseBody {
	s.NonExistPolicyNames = v
	return s
}

func (s *DetachAppPolicyFromIdentityResponseBody) SetRequestId(v string) *DetachAppPolicyFromIdentityResponseBody {
	s.RequestId = &v
	return s
}

type DetachAppPolicyFromIdentityResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DetachAppPolicyFromIdentityResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DetachAppPolicyFromIdentityResponse) String() string {
	return tea.Prettify(s)
}

func (s DetachAppPolicyFromIdentityResponse) GoString() string {
	return s.String()
}

func (s *DetachAppPolicyFromIdentityResponse) SetHeaders(v map[string]*string) *DetachAppPolicyFromIdentityResponse {
	s.Headers = v
	return s
}

func (s *DetachAppPolicyFromIdentityResponse) SetBody(v *DetachAppPolicyFromIdentityResponseBody) *DetachAppPolicyFromIdentityResponse {
	s.Body = v
	return s
}

type DisableVodRealtimeLogDeliveryRequest struct {
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s DisableVodRealtimeLogDeliveryRequest) String() string {
	return tea.Prettify(s)
}

func (s DisableVodRealtimeLogDeliveryRequest) GoString() string {
	return s.String()
}

func (s *DisableVodRealtimeLogDeliveryRequest) SetDomainName(v string) *DisableVodRealtimeLogDeliveryRequest {
	s.DomainName = &v
	return s
}

func (s *DisableVodRealtimeLogDeliveryRequest) SetOwnerId(v int64) *DisableVodRealtimeLogDeliveryRequest {
	s.OwnerId = &v
	return s
}

type DisableVodRealtimeLogDeliveryResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisableVodRealtimeLogDeliveryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DisableVodRealtimeLogDeliveryResponseBody) GoString() string {
	return s.String()
}

func (s *DisableVodRealtimeLogDeliveryResponseBody) SetRequestId(v string) *DisableVodRealtimeLogDeliveryResponseBody {
	s.RequestId = &v
	return s
}

type DisableVodRealtimeLogDeliveryResponse struct {
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DisableVodRealtimeLogDeliveryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DisableVodRealtimeLogDeliveryResponse) String() string {
	return tea.Prettify(s)
}

func (s DisableVodRealtimeLogDeliveryResponse) GoString() string {
	return s.String()
}

func (s *DisableVodRealtimeLogDeliveryResponse) SetHeaders(v map[string]*string) *DisableVodRealtimeLogDeliveryResponse {
	s.Headers = v
	return s
}

func (s *DisableVodRealtimeLogDeliveryResponse) SetBody(v *DisableVodRealtimeLogDeliveryResponseBody) *DisableVodRealtimeLogDeliveryResponse {
	s.Body = v
	return s
}

type EnableVodRealtimeLogDeliveryRequest struct {
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s EnableVodRealtimeLogDeliveryRequest) String() string {
	return tea.Prettify(s)
}

func (s EnableVodRealtimeLogDeliveryRequest) GoString() string {
	return s.String()
}

func (s *EnableVodRealtimeLogDeliveryRequest) SetDomainName(v string) *EnableVodRealtimeLogDeliveryRequest {
	s.DomainName = &v
	return s
}

func (s *EnableVodRealtimeLogDeliveryRequest) SetOwnerId(v int64) *EnableVodRealtimeLogDeliveryRequest {
	s.OwnerId = &v
	return s
}

type EnableVodRealtimeLogDeliveryResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnableVodRealtimeLogDeliveryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EnableVodRealtimeLogDeliveryResponseBody) GoString() string {
	return s.String()
}

func (s *EnableVodRealtimeLogDeliveryResponseBody) SetRequestId(v string) *EnableVodRealtimeLogDeliveryResponseBody {
	s.RequestId = &v
	return s
}

type EnableVodRealtimeLogDeliveryResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *EnableVodRealtimeLogDeliveryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s EnableVodRealtimeLogDeliveryResponse) String() string {
	return tea.Prettify(s)
}

func (s EnableVodRealtimeLogDeliveryResponse) GoString() string {
	return s.String()
}

func (s *EnableVodRealtimeLogDeliveryResponse) SetHeaders(v map[string]*string) *EnableVodRealtimeLogDeliveryResponse {
	s.Headers = v
	return s
}

func (s *EnableVodRealtimeLogDeliveryResponse) SetBody(v *EnableVodRealtimeLogDeliveryResponseBody) *EnableVodRealtimeLogDeliveryResponse {
	s.Body = v
	return s
}

type GetAICaptionExtractionJobsRequest struct {
	JobIds *string `json:"JobIds,omitempty" xml:"JobIds,omitempty"`
}

func (s GetAICaptionExtractionJobsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAICaptionExtractionJobsRequest) GoString() string {
	return s.String()
}

func (s *GetAICaptionExtractionJobsRequest) SetJobIds(v string) *GetAICaptionExtractionJobsRequest {
	s.JobIds = &v
	return s
}

type GetAICaptionExtractionJobsResponseBody struct {
	AICaptionExtractionJobList []*GetAICaptionExtractionJobsResponseBodyAICaptionExtractionJobList `json:"AICaptionExtractionJobList,omitempty" xml:"AICaptionExtractionJobList,omitempty" type:"Repeated"`
	RequestId                  *string                                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAICaptionExtractionJobsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAICaptionExtractionJobsResponseBody) GoString() string {
	return s.String()
}

func (s *GetAICaptionExtractionJobsResponseBody) SetAICaptionExtractionJobList(v []*GetAICaptionExtractionJobsResponseBodyAICaptionExtractionJobList) *GetAICaptionExtractionJobsResponseBody {
	s.AICaptionExtractionJobList = v
	return s
}

func (s *GetAICaptionExtractionJobsResponseBody) SetRequestId(v string) *GetAICaptionExtractionJobsResponseBody {
	s.RequestId = &v
	return s
}

type GetAICaptionExtractionJobsResponseBodyAICaptionExtractionJobList struct {
	AICaptionExtractionResult *string `json:"AICaptionExtractionResult,omitempty" xml:"AICaptionExtractionResult,omitempty"`
	Code                      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	CreationTime              *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	JobId                     *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	Message                   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Status                    *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TemplateConfig            *string `json:"TemplateConfig,omitempty" xml:"TemplateConfig,omitempty"`
	UserData                  *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
	VideoId                   *string `json:"VideoId,omitempty" xml:"VideoId,omitempty"`
}

func (s GetAICaptionExtractionJobsResponseBodyAICaptionExtractionJobList) String() string {
	return tea.Prettify(s)
}

func (s GetAICaptionExtractionJobsResponseBodyAICaptionExtractionJobList) GoString() string {
	return s.String()
}

func (s *GetAICaptionExtractionJobsResponseBodyAICaptionExtractionJobList) SetAICaptionExtractionResult(v string) *GetAICaptionExtractionJobsResponseBodyAICaptionExtractionJobList {
	s.AICaptionExtractionResult = &v
	return s
}

func (s *GetAICaptionExtractionJobsResponseBodyAICaptionExtractionJobList) SetCode(v string) *GetAICaptionExtractionJobsResponseBodyAICaptionExtractionJobList {
	s.Code = &v
	return s
}

func (s *GetAICaptionExtractionJobsResponseBodyAICaptionExtractionJobList) SetCreationTime(v string) *GetAICaptionExtractionJobsResponseBodyAICaptionExtractionJobList {
	s.CreationTime = &v
	return s
}

func (s *GetAICaptionExtractionJobsResponseBodyAICaptionExtractionJobList) SetJobId(v string) *GetAICaptionExtractionJobsResponseBodyAICaptionExtractionJobList {
	s.JobId = &v
	return s
}

func (s *GetAICaptionExtractionJobsResponseBodyAICaptionExtractionJobList) SetMessage(v string) *GetAICaptionExtractionJobsResponseBodyAICaptionExtractionJobList {
	s.Message = &v
	return s
}

func (s *GetAICaptionExtractionJobsResponseBodyAICaptionExtractionJobList) SetStatus(v string) *GetAICaptionExtractionJobsResponseBodyAICaptionExtractionJobList {
	s.Status = &v
	return s
}

func (s *GetAICaptionExtractionJobsResponseBodyAICaptionExtractionJobList) SetTemplateConfig(v string) *GetAICaptionExtractionJobsResponseBodyAICaptionExtractionJobList {
	s.TemplateConfig = &v
	return s
}

func (s *GetAICaptionExtractionJobsResponseBodyAICaptionExtractionJobList) SetUserData(v string) *GetAICaptionExtractionJobsResponseBodyAICaptionExtractionJobList {
	s.UserData = &v
	return s
}

func (s *GetAICaptionExtractionJobsResponseBodyAICaptionExtractionJobList) SetVideoId(v string) *GetAICaptionExtractionJobsResponseBodyAICaptionExtractionJobList {
	s.VideoId = &v
	return s
}

type GetAICaptionExtractionJobsResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetAICaptionExtractionJobsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetAICaptionExtractionJobsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAICaptionExtractionJobsResponse) GoString() string {
	return s.String()
}

func (s *GetAICaptionExtractionJobsResponse) SetHeaders(v map[string]*string) *GetAICaptionExtractionJobsResponse {
	s.Headers = v
	return s
}

func (s *GetAICaptionExtractionJobsResponse) SetBody(v *GetAICaptionExtractionJobsResponseBody) *GetAICaptionExtractionJobsResponse {
	s.Body = v
	return s
}

type GetAIImageJobsRequest struct {
	JobIds               *string `json:"JobIds,omitempty" xml:"JobIds,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *string `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s GetAIImageJobsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAIImageJobsRequest) GoString() string {
	return s.String()
}

func (s *GetAIImageJobsRequest) SetJobIds(v string) *GetAIImageJobsRequest {
	s.JobIds = &v
	return s
}

func (s *GetAIImageJobsRequest) SetOwnerAccount(v string) *GetAIImageJobsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *GetAIImageJobsRequest) SetOwnerId(v string) *GetAIImageJobsRequest {
	s.OwnerId = &v
	return s
}

func (s *GetAIImageJobsRequest) SetResourceOwnerAccount(v string) *GetAIImageJobsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetAIImageJobsRequest) SetResourceOwnerId(v string) *GetAIImageJobsRequest {
	s.ResourceOwnerId = &v
	return s
}

type GetAIImageJobsResponseBody struct {
	AIImageJobList []*GetAIImageJobsResponseBodyAIImageJobList `json:"AIImageJobList,omitempty" xml:"AIImageJobList,omitempty" type:"Repeated"`
	RequestId      *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAIImageJobsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAIImageJobsResponseBody) GoString() string {
	return s.String()
}

func (s *GetAIImageJobsResponseBody) SetAIImageJobList(v []*GetAIImageJobsResponseBodyAIImageJobList) *GetAIImageJobsResponseBody {
	s.AIImageJobList = v
	return s
}

func (s *GetAIImageJobsResponseBody) SetRequestId(v string) *GetAIImageJobsResponseBody {
	s.RequestId = &v
	return s
}

type GetAIImageJobsResponseBodyAIImageJobList struct {
	AIImageResult  *string `json:"AIImageResult,omitempty" xml:"AIImageResult,omitempty"`
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	CreationTime   *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	JobId          *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Status         *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TemplateConfig *string `json:"TemplateConfig,omitempty" xml:"TemplateConfig,omitempty"`
	TemplateId     *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	UserData       *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
	VideoId        *string `json:"VideoId,omitempty" xml:"VideoId,omitempty"`
}

func (s GetAIImageJobsResponseBodyAIImageJobList) String() string {
	return tea.Prettify(s)
}

func (s GetAIImageJobsResponseBodyAIImageJobList) GoString() string {
	return s.String()
}

func (s *GetAIImageJobsResponseBodyAIImageJobList) SetAIImageResult(v string) *GetAIImageJobsResponseBodyAIImageJobList {
	s.AIImageResult = &v
	return s
}

func (s *GetAIImageJobsResponseBodyAIImageJobList) SetCode(v string) *GetAIImageJobsResponseBodyAIImageJobList {
	s.Code = &v
	return s
}

func (s *GetAIImageJobsResponseBodyAIImageJobList) SetCreationTime(v string) *GetAIImageJobsResponseBodyAIImageJobList {
	s.CreationTime = &v
	return s
}

func (s *GetAIImageJobsResponseBodyAIImageJobList) SetJobId(v string) *GetAIImageJobsResponseBodyAIImageJobList {
	s.JobId = &v
	return s
}

func (s *GetAIImageJobsResponseBodyAIImageJobList) SetMessage(v string) *GetAIImageJobsResponseBodyAIImageJobList {
	s.Message = &v
	return s
}

func (s *GetAIImageJobsResponseBodyAIImageJobList) SetStatus(v string) *GetAIImageJobsResponseBodyAIImageJobList {
	s.Status = &v
	return s
}

func (s *GetAIImageJobsResponseBodyAIImageJobList) SetTemplateConfig(v string) *GetAIImageJobsResponseBodyAIImageJobList {
	s.TemplateConfig = &v
	return s
}

func (s *GetAIImageJobsResponseBodyAIImageJobList) SetTemplateId(v string) *GetAIImageJobsResponseBodyAIImageJobList {
	s.TemplateId = &v
	return s
}

func (s *GetAIImageJobsResponseBodyAIImageJobList) SetUserData(v string) *GetAIImageJobsResponseBodyAIImageJobList {
	s.UserData = &v
	return s
}

func (s *GetAIImageJobsResponseBodyAIImageJobList) SetVideoId(v string) *GetAIImageJobsResponseBodyAIImageJobList {
	s.VideoId = &v
	return s
}

type GetAIImageJobsResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetAIImageJobsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetAIImageJobsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAIImageJobsResponse) GoString() string {
	return s.String()
}

func (s *GetAIImageJobsResponse) SetHeaders(v map[string]*string) *GetAIImageJobsResponse {
	s.Headers = v
	return s
}

func (s *GetAIImageJobsResponse) SetBody(v *GetAIImageJobsResponseBody) *GetAIImageJobsResponse {
	s.Body = v
	return s
}

type GetAIMediaAuditJobRequest struct {
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s GetAIMediaAuditJobRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAIMediaAuditJobRequest) GoString() string {
	return s.String()
}

func (s *GetAIMediaAuditJobRequest) SetJobId(v string) *GetAIMediaAuditJobRequest {
	s.JobId = &v
	return s
}

type GetAIMediaAuditJobResponseBody struct {
	MediaAuditJob *GetAIMediaAuditJobResponseBodyMediaAuditJob `json:"MediaAuditJob,omitempty" xml:"MediaAuditJob,omitempty" type:"Struct"`
	RequestId     *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAIMediaAuditJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAIMediaAuditJobResponseBody) GoString() string {
	return s.String()
}

func (s *GetAIMediaAuditJobResponseBody) SetMediaAuditJob(v *GetAIMediaAuditJobResponseBodyMediaAuditJob) *GetAIMediaAuditJobResponseBody {
	s.MediaAuditJob = v
	return s
}

func (s *GetAIMediaAuditJobResponseBody) SetRequestId(v string) *GetAIMediaAuditJobResponseBody {
	s.RequestId = &v
	return s
}

type GetAIMediaAuditJobResponseBodyMediaAuditJob struct {
	Code         *string                                          `json:"Code,omitempty" xml:"Code,omitempty"`
	CompleteTime *string                                          `json:"CompleteTime,omitempty" xml:"CompleteTime,omitempty"`
	CreationTime *string                                          `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	Data         *GetAIMediaAuditJobResponseBodyMediaAuditJobData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	JobId        *string                                          `json:"JobId,omitempty" xml:"JobId,omitempty"`
	MediaId      *string                                          `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	Message      *string                                          `json:"Message,omitempty" xml:"Message,omitempty"`
	Status       *string                                          `json:"Status,omitempty" xml:"Status,omitempty"`
	Type         *string                                          `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetAIMediaAuditJobResponseBodyMediaAuditJob) String() string {
	return tea.Prettify(s)
}

func (s GetAIMediaAuditJobResponseBodyMediaAuditJob) GoString() string {
	return s.String()
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJob) SetCode(v string) *GetAIMediaAuditJobResponseBodyMediaAuditJob {
	s.Code = &v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJob) SetCompleteTime(v string) *GetAIMediaAuditJobResponseBodyMediaAuditJob {
	s.CompleteTime = &v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJob) SetCreationTime(v string) *GetAIMediaAuditJobResponseBodyMediaAuditJob {
	s.CreationTime = &v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJob) SetData(v *GetAIMediaAuditJobResponseBodyMediaAuditJobData) *GetAIMediaAuditJobResponseBodyMediaAuditJob {
	s.Data = v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJob) SetJobId(v string) *GetAIMediaAuditJobResponseBodyMediaAuditJob {
	s.JobId = &v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJob) SetMediaId(v string) *GetAIMediaAuditJobResponseBodyMediaAuditJob {
	s.MediaId = &v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJob) SetMessage(v string) *GetAIMediaAuditJobResponseBodyMediaAuditJob {
	s.Message = &v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJob) SetStatus(v string) *GetAIMediaAuditJobResponseBodyMediaAuditJob {
	s.Status = &v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJob) SetType(v string) *GetAIMediaAuditJobResponseBodyMediaAuditJob {
	s.Type = &v
	return s
}

type GetAIMediaAuditJobResponseBodyMediaAuditJobData struct {
	AbnormalModules *string                                                       `json:"AbnormalModules,omitempty" xml:"AbnormalModules,omitempty"`
	AudioResult     []*GetAIMediaAuditJobResponseBodyMediaAuditJobDataAudioResult `json:"AudioResult,omitempty" xml:"AudioResult,omitempty" type:"Repeated"`
	ImageResult     []*GetAIMediaAuditJobResponseBodyMediaAuditJobDataImageResult `json:"ImageResult,omitempty" xml:"ImageResult,omitempty" type:"Repeated"`
	Label           *string                                                       `json:"Label,omitempty" xml:"Label,omitempty"`
	Suggestion      *string                                                       `json:"Suggestion,omitempty" xml:"Suggestion,omitempty"`
	TextResult      []*GetAIMediaAuditJobResponseBodyMediaAuditJobDataTextResult  `json:"TextResult,omitempty" xml:"TextResult,omitempty" type:"Repeated"`
	VideoResult     *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResult   `json:"VideoResult,omitempty" xml:"VideoResult,omitempty" type:"Struct"`
}

func (s GetAIMediaAuditJobResponseBodyMediaAuditJobData) String() string {
	return tea.Prettify(s)
}

func (s GetAIMediaAuditJobResponseBodyMediaAuditJobData) GoString() string {
	return s.String()
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobData) SetAbnormalModules(v string) *GetAIMediaAuditJobResponseBodyMediaAuditJobData {
	s.AbnormalModules = &v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobData) SetAudioResult(v []*GetAIMediaAuditJobResponseBodyMediaAuditJobDataAudioResult) *GetAIMediaAuditJobResponseBodyMediaAuditJobData {
	s.AudioResult = v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobData) SetImageResult(v []*GetAIMediaAuditJobResponseBodyMediaAuditJobDataImageResult) *GetAIMediaAuditJobResponseBodyMediaAuditJobData {
	s.ImageResult = v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobData) SetLabel(v string) *GetAIMediaAuditJobResponseBodyMediaAuditJobData {
	s.Label = &v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobData) SetSuggestion(v string) *GetAIMediaAuditJobResponseBodyMediaAuditJobData {
	s.Suggestion = &v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobData) SetTextResult(v []*GetAIMediaAuditJobResponseBodyMediaAuditJobDataTextResult) *GetAIMediaAuditJobResponseBodyMediaAuditJobData {
	s.TextResult = v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobData) SetVideoResult(v *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResult) *GetAIMediaAuditJobResponseBodyMediaAuditJobData {
	s.VideoResult = v
	return s
}

type GetAIMediaAuditJobResponseBodyMediaAuditJobDataAudioResult struct {
	Label      *string `json:"Label,omitempty" xml:"Label,omitempty"`
	Scene      *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
	Score      *string `json:"Score,omitempty" xml:"Score,omitempty"`
	Suggestion *string `json:"Suggestion,omitempty" xml:"Suggestion,omitempty"`
}

func (s GetAIMediaAuditJobResponseBodyMediaAuditJobDataAudioResult) String() string {
	return tea.Prettify(s)
}

func (s GetAIMediaAuditJobResponseBodyMediaAuditJobDataAudioResult) GoString() string {
	return s.String()
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataAudioResult) SetLabel(v string) *GetAIMediaAuditJobResponseBodyMediaAuditJobDataAudioResult {
	s.Label = &v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataAudioResult) SetScene(v string) *GetAIMediaAuditJobResponseBodyMediaAuditJobDataAudioResult {
	s.Scene = &v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataAudioResult) SetScore(v string) *GetAIMediaAuditJobResponseBodyMediaAuditJobDataAudioResult {
	s.Score = &v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataAudioResult) SetSuggestion(v string) *GetAIMediaAuditJobResponseBodyMediaAuditJobDataAudioResult {
	s.Suggestion = &v
	return s
}

type GetAIMediaAuditJobResponseBodyMediaAuditJobDataImageResult struct {
	Label      *string                                                             `json:"Label,omitempty" xml:"Label,omitempty"`
	Result     []*GetAIMediaAuditJobResponseBodyMediaAuditJobDataImageResultResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	Suggestion *string                                                             `json:"Suggestion,omitempty" xml:"Suggestion,omitempty"`
	Type       *string                                                             `json:"Type,omitempty" xml:"Type,omitempty"`
	Url        *string                                                             `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s GetAIMediaAuditJobResponseBodyMediaAuditJobDataImageResult) String() string {
	return tea.Prettify(s)
}

func (s GetAIMediaAuditJobResponseBodyMediaAuditJobDataImageResult) GoString() string {
	return s.String()
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataImageResult) SetLabel(v string) *GetAIMediaAuditJobResponseBodyMediaAuditJobDataImageResult {
	s.Label = &v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataImageResult) SetResult(v []*GetAIMediaAuditJobResponseBodyMediaAuditJobDataImageResultResult) *GetAIMediaAuditJobResponseBodyMediaAuditJobDataImageResult {
	s.Result = v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataImageResult) SetSuggestion(v string) *GetAIMediaAuditJobResponseBodyMediaAuditJobDataImageResult {
	s.Suggestion = &v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataImageResult) SetType(v string) *GetAIMediaAuditJobResponseBodyMediaAuditJobDataImageResult {
	s.Type = &v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataImageResult) SetUrl(v string) *GetAIMediaAuditJobResponseBodyMediaAuditJobDataImageResult {
	s.Url = &v
	return s
}

type GetAIMediaAuditJobResponseBodyMediaAuditJobDataImageResultResult struct {
	Label      *string `json:"Label,omitempty" xml:"Label,omitempty"`
	Scene      *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
	Score      *string `json:"Score,omitempty" xml:"Score,omitempty"`
	Suggestion *string `json:"Suggestion,omitempty" xml:"Suggestion,omitempty"`
}

func (s GetAIMediaAuditJobResponseBodyMediaAuditJobDataImageResultResult) String() string {
	return tea.Prettify(s)
}

func (s GetAIMediaAuditJobResponseBodyMediaAuditJobDataImageResultResult) GoString() string {
	return s.String()
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataImageResultResult) SetLabel(v string) *GetAIMediaAuditJobResponseBodyMediaAuditJobDataImageResultResult {
	s.Label = &v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataImageResultResult) SetScene(v string) *GetAIMediaAuditJobResponseBodyMediaAuditJobDataImageResultResult {
	s.Scene = &v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataImageResultResult) SetScore(v string) *GetAIMediaAuditJobResponseBodyMediaAuditJobDataImageResultResult {
	s.Score = &v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataImageResultResult) SetSuggestion(v string) *GetAIMediaAuditJobResponseBodyMediaAuditJobDataImageResultResult {
	s.Suggestion = &v
	return s
}

type GetAIMediaAuditJobResponseBodyMediaAuditJobDataTextResult struct {
	Content    *string `json:"Content,omitempty" xml:"Content,omitempty"`
	Label      *string `json:"Label,omitempty" xml:"Label,omitempty"`
	Scene      *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
	Score      *string `json:"Score,omitempty" xml:"Score,omitempty"`
	Suggestion *string `json:"Suggestion,omitempty" xml:"Suggestion,omitempty"`
	Type       *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetAIMediaAuditJobResponseBodyMediaAuditJobDataTextResult) String() string {
	return tea.Prettify(s)
}

func (s GetAIMediaAuditJobResponseBodyMediaAuditJobDataTextResult) GoString() string {
	return s.String()
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataTextResult) SetContent(v string) *GetAIMediaAuditJobResponseBodyMediaAuditJobDataTextResult {
	s.Content = &v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataTextResult) SetLabel(v string) *GetAIMediaAuditJobResponseBodyMediaAuditJobDataTextResult {
	s.Label = &v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataTextResult) SetScene(v string) *GetAIMediaAuditJobResponseBodyMediaAuditJobDataTextResult {
	s.Scene = &v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataTextResult) SetScore(v string) *GetAIMediaAuditJobResponseBodyMediaAuditJobDataTextResult {
	s.Score = &v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataTextResult) SetSuggestion(v string) *GetAIMediaAuditJobResponseBodyMediaAuditJobDataTextResult {
	s.Suggestion = &v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataTextResult) SetType(v string) *GetAIMediaAuditJobResponseBodyMediaAuditJobDataTextResult {
	s.Type = &v
	return s
}

type GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResult struct {
	AdResult        *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultAdResult        `json:"AdResult,omitempty" xml:"AdResult,omitempty" type:"Struct"`
	Label           *string                                                                    `json:"Label,omitempty" xml:"Label,omitempty"`
	LiveResult      *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLiveResult      `json:"LiveResult,omitempty" xml:"LiveResult,omitempty" type:"Struct"`
	LogoResult      *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLogoResult      `json:"LogoResult,omitempty" xml:"LogoResult,omitempty" type:"Struct"`
	PornResult      *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultPornResult      `json:"PornResult,omitempty" xml:"PornResult,omitempty" type:"Struct"`
	Suggestion      *string                                                                    `json:"Suggestion,omitempty" xml:"Suggestion,omitempty"`
	TerrorismResult *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultTerrorismResult `json:"TerrorismResult,omitempty" xml:"TerrorismResult,omitempty" type:"Struct"`
}

func (s GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResult) String() string {
	return tea.Prettify(s)
}

func (s GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResult) GoString() string {
	return s.String()
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResult) SetAdResult(v *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultAdResult) *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResult {
	s.AdResult = v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResult) SetLabel(v string) *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResult {
	s.Label = &v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResult) SetLiveResult(v *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLiveResult) *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResult {
	s.LiveResult = v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResult) SetLogoResult(v *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLogoResult) *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResult {
	s.LogoResult = v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResult) SetPornResult(v *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultPornResult) *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResult {
	s.PornResult = v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResult) SetSuggestion(v string) *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResult {
	s.Suggestion = &v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResult) SetTerrorismResult(v *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultTerrorismResult) *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResult {
	s.TerrorismResult = v
	return s
}

type GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultAdResult struct {
	AverageScore *string                                                                          `json:"AverageScore,omitempty" xml:"AverageScore,omitempty"`
	CounterList  []*GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultAdResultCounterList `json:"CounterList,omitempty" xml:"CounterList,omitempty" type:"Repeated"`
	Label        *string                                                                          `json:"Label,omitempty" xml:"Label,omitempty"`
	MaxScore     *string                                                                          `json:"MaxScore,omitempty" xml:"MaxScore,omitempty"`
	Suggestion   *string                                                                          `json:"Suggestion,omitempty" xml:"Suggestion,omitempty"`
	TopList      []*GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultAdResultTopList     `json:"TopList,omitempty" xml:"TopList,omitempty" type:"Repeated"`
}

func (s GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultAdResult) String() string {
	return tea.Prettify(s)
}

func (s GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultAdResult) GoString() string {
	return s.String()
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultAdResult) SetAverageScore(v string) *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultAdResult {
	s.AverageScore = &v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultAdResult) SetCounterList(v []*GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultAdResultCounterList) *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultAdResult {
	s.CounterList = v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultAdResult) SetLabel(v string) *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultAdResult {
	s.Label = &v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultAdResult) SetMaxScore(v string) *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultAdResult {
	s.MaxScore = &v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultAdResult) SetSuggestion(v string) *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultAdResult {
	s.Suggestion = &v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultAdResult) SetTopList(v []*GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultAdResultTopList) *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultAdResult {
	s.TopList = v
	return s
}

type GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultAdResultCounterList struct {
	Count *int32  `json:"Count,omitempty" xml:"Count,omitempty"`
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
}

func (s GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultAdResultCounterList) String() string {
	return tea.Prettify(s)
}

func (s GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultAdResultCounterList) GoString() string {
	return s.String()
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultAdResultCounterList) SetCount(v int32) *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultAdResultCounterList {
	s.Count = &v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultAdResultCounterList) SetLabel(v string) *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultAdResultCounterList {
	s.Label = &v
	return s
}

type GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultAdResultTopList struct {
	Label     *string `json:"Label,omitempty" xml:"Label,omitempty"`
	Score     *string `json:"Score,omitempty" xml:"Score,omitempty"`
	Timestamp *string `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	Url       *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultAdResultTopList) String() string {
	return tea.Prettify(s)
}

func (s GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultAdResultTopList) GoString() string {
	return s.String()
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultAdResultTopList) SetLabel(v string) *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultAdResultTopList {
	s.Label = &v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultAdResultTopList) SetScore(v string) *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultAdResultTopList {
	s.Score = &v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultAdResultTopList) SetTimestamp(v string) *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultAdResultTopList {
	s.Timestamp = &v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultAdResultTopList) SetUrl(v string) *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultAdResultTopList {
	s.Url = &v
	return s
}

type GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLiveResult struct {
	AverageScore *string                                                                            `json:"AverageScore,omitempty" xml:"AverageScore,omitempty"`
	CounterList  []*GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLiveResultCounterList `json:"CounterList,omitempty" xml:"CounterList,omitempty" type:"Repeated"`
	Label        *string                                                                            `json:"Label,omitempty" xml:"Label,omitempty"`
	MaxScore     *string                                                                            `json:"MaxScore,omitempty" xml:"MaxScore,omitempty"`
	Suggestion   *string                                                                            `json:"Suggestion,omitempty" xml:"Suggestion,omitempty"`
	TopList      []*GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLiveResultTopList     `json:"TopList,omitempty" xml:"TopList,omitempty" type:"Repeated"`
}

func (s GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLiveResult) String() string {
	return tea.Prettify(s)
}

func (s GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLiveResult) GoString() string {
	return s.String()
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLiveResult) SetAverageScore(v string) *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLiveResult {
	s.AverageScore = &v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLiveResult) SetCounterList(v []*GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLiveResultCounterList) *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLiveResult {
	s.CounterList = v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLiveResult) SetLabel(v string) *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLiveResult {
	s.Label = &v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLiveResult) SetMaxScore(v string) *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLiveResult {
	s.MaxScore = &v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLiveResult) SetSuggestion(v string) *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLiveResult {
	s.Suggestion = &v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLiveResult) SetTopList(v []*GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLiveResultTopList) *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLiveResult {
	s.TopList = v
	return s
}

type GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLiveResultCounterList struct {
	Count *int32  `json:"Count,omitempty" xml:"Count,omitempty"`
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
}

func (s GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLiveResultCounterList) String() string {
	return tea.Prettify(s)
}

func (s GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLiveResultCounterList) GoString() string {
	return s.String()
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLiveResultCounterList) SetCount(v int32) *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLiveResultCounterList {
	s.Count = &v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLiveResultCounterList) SetLabel(v string) *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLiveResultCounterList {
	s.Label = &v
	return s
}

type GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLiveResultTopList struct {
	Label     *string `json:"Label,omitempty" xml:"Label,omitempty"`
	Score     *string `json:"Score,omitempty" xml:"Score,omitempty"`
	Timestamp *string `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	Url       *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLiveResultTopList) String() string {
	return tea.Prettify(s)
}

func (s GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLiveResultTopList) GoString() string {
	return s.String()
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLiveResultTopList) SetLabel(v string) *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLiveResultTopList {
	s.Label = &v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLiveResultTopList) SetScore(v string) *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLiveResultTopList {
	s.Score = &v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLiveResultTopList) SetTimestamp(v string) *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLiveResultTopList {
	s.Timestamp = &v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLiveResultTopList) SetUrl(v string) *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLiveResultTopList {
	s.Url = &v
	return s
}

type GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLogoResult struct {
	AverageScore *string                                                                            `json:"AverageScore,omitempty" xml:"AverageScore,omitempty"`
	CounterList  []*GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLogoResultCounterList `json:"CounterList,omitempty" xml:"CounterList,omitempty" type:"Repeated"`
	Label        *string                                                                            `json:"Label,omitempty" xml:"Label,omitempty"`
	MaxScore     *string                                                                            `json:"MaxScore,omitempty" xml:"MaxScore,omitempty"`
	Suggestion   *string                                                                            `json:"Suggestion,omitempty" xml:"Suggestion,omitempty"`
	TopList      []*GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLogoResultTopList     `json:"TopList,omitempty" xml:"TopList,omitempty" type:"Repeated"`
}

func (s GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLogoResult) String() string {
	return tea.Prettify(s)
}

func (s GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLogoResult) GoString() string {
	return s.String()
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLogoResult) SetAverageScore(v string) *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLogoResult {
	s.AverageScore = &v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLogoResult) SetCounterList(v []*GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLogoResultCounterList) *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLogoResult {
	s.CounterList = v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLogoResult) SetLabel(v string) *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLogoResult {
	s.Label = &v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLogoResult) SetMaxScore(v string) *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLogoResult {
	s.MaxScore = &v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLogoResult) SetSuggestion(v string) *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLogoResult {
	s.Suggestion = &v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLogoResult) SetTopList(v []*GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLogoResultTopList) *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLogoResult {
	s.TopList = v
	return s
}

type GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLogoResultCounterList struct {
	Count *int32  `json:"Count,omitempty" xml:"Count,omitempty"`
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
}

func (s GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLogoResultCounterList) String() string {
	return tea.Prettify(s)
}

func (s GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLogoResultCounterList) GoString() string {
	return s.String()
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLogoResultCounterList) SetCount(v int32) *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLogoResultCounterList {
	s.Count = &v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLogoResultCounterList) SetLabel(v string) *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLogoResultCounterList {
	s.Label = &v
	return s
}

type GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLogoResultTopList struct {
	Label     *string `json:"Label,omitempty" xml:"Label,omitempty"`
	Score     *string `json:"Score,omitempty" xml:"Score,omitempty"`
	Timestamp *string `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	Url       *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLogoResultTopList) String() string {
	return tea.Prettify(s)
}

func (s GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLogoResultTopList) GoString() string {
	return s.String()
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLogoResultTopList) SetLabel(v string) *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLogoResultTopList {
	s.Label = &v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLogoResultTopList) SetScore(v string) *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLogoResultTopList {
	s.Score = &v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLogoResultTopList) SetTimestamp(v string) *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLogoResultTopList {
	s.Timestamp = &v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLogoResultTopList) SetUrl(v string) *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultLogoResultTopList {
	s.Url = &v
	return s
}

type GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultPornResult struct {
	AverageScore *string                                                                            `json:"AverageScore,omitempty" xml:"AverageScore,omitempty"`
	CounterList  []*GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultPornResultCounterList `json:"CounterList,omitempty" xml:"CounterList,omitempty" type:"Repeated"`
	Label        *string                                                                            `json:"Label,omitempty" xml:"Label,omitempty"`
	MaxScore     *string                                                                            `json:"MaxScore,omitempty" xml:"MaxScore,omitempty"`
	Suggestion   *string                                                                            `json:"Suggestion,omitempty" xml:"Suggestion,omitempty"`
	TopList      []*GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultPornResultTopList     `json:"TopList,omitempty" xml:"TopList,omitempty" type:"Repeated"`
}

func (s GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultPornResult) String() string {
	return tea.Prettify(s)
}

func (s GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultPornResult) GoString() string {
	return s.String()
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultPornResult) SetAverageScore(v string) *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultPornResult {
	s.AverageScore = &v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultPornResult) SetCounterList(v []*GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultPornResultCounterList) *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultPornResult {
	s.CounterList = v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultPornResult) SetLabel(v string) *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultPornResult {
	s.Label = &v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultPornResult) SetMaxScore(v string) *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultPornResult {
	s.MaxScore = &v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultPornResult) SetSuggestion(v string) *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultPornResult {
	s.Suggestion = &v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultPornResult) SetTopList(v []*GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultPornResultTopList) *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultPornResult {
	s.TopList = v
	return s
}

type GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultPornResultCounterList struct {
	Count *int32  `json:"Count,omitempty" xml:"Count,omitempty"`
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
}

func (s GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultPornResultCounterList) String() string {
	return tea.Prettify(s)
}

func (s GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultPornResultCounterList) GoString() string {
	return s.String()
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultPornResultCounterList) SetCount(v int32) *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultPornResultCounterList {
	s.Count = &v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultPornResultCounterList) SetLabel(v string) *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultPornResultCounterList {
	s.Label = &v
	return s
}

type GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultPornResultTopList struct {
	Label     *string `json:"Label,omitempty" xml:"Label,omitempty"`
	Score     *string `json:"Score,omitempty" xml:"Score,omitempty"`
	Timestamp *string `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	Url       *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultPornResultTopList) String() string {
	return tea.Prettify(s)
}

func (s GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultPornResultTopList) GoString() string {
	return s.String()
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultPornResultTopList) SetLabel(v string) *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultPornResultTopList {
	s.Label = &v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultPornResultTopList) SetScore(v string) *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultPornResultTopList {
	s.Score = &v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultPornResultTopList) SetTimestamp(v string) *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultPornResultTopList {
	s.Timestamp = &v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultPornResultTopList) SetUrl(v string) *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultPornResultTopList {
	s.Url = &v
	return s
}

type GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultTerrorismResult struct {
	AverageScore *string                                                                                 `json:"AverageScore,omitempty" xml:"AverageScore,omitempty"`
	CounterList  []*GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultTerrorismResultCounterList `json:"CounterList,omitempty" xml:"CounterList,omitempty" type:"Repeated"`
	Label        *string                                                                                 `json:"Label,omitempty" xml:"Label,omitempty"`
	MaxScore     *string                                                                                 `json:"MaxScore,omitempty" xml:"MaxScore,omitempty"`
	Suggestion   *string                                                                                 `json:"Suggestion,omitempty" xml:"Suggestion,omitempty"`
	TopList      []*GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultTerrorismResultTopList     `json:"TopList,omitempty" xml:"TopList,omitempty" type:"Repeated"`
}

func (s GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultTerrorismResult) String() string {
	return tea.Prettify(s)
}

func (s GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultTerrorismResult) GoString() string {
	return s.String()
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultTerrorismResult) SetAverageScore(v string) *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultTerrorismResult {
	s.AverageScore = &v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultTerrorismResult) SetCounterList(v []*GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultTerrorismResultCounterList) *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultTerrorismResult {
	s.CounterList = v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultTerrorismResult) SetLabel(v string) *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultTerrorismResult {
	s.Label = &v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultTerrorismResult) SetMaxScore(v string) *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultTerrorismResult {
	s.MaxScore = &v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultTerrorismResult) SetSuggestion(v string) *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultTerrorismResult {
	s.Suggestion = &v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultTerrorismResult) SetTopList(v []*GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultTerrorismResultTopList) *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultTerrorismResult {
	s.TopList = v
	return s
}

type GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultTerrorismResultCounterList struct {
	Count *int32  `json:"Count,omitempty" xml:"Count,omitempty"`
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
}

func (s GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultTerrorismResultCounterList) String() string {
	return tea.Prettify(s)
}

func (s GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultTerrorismResultCounterList) GoString() string {
	return s.String()
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultTerrorismResultCounterList) SetCount(v int32) *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultTerrorismResultCounterList {
	s.Count = &v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultTerrorismResultCounterList) SetLabel(v string) *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultTerrorismResultCounterList {
	s.Label = &v
	return s
}

type GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultTerrorismResultTopList struct {
	Label     *string `json:"Label,omitempty" xml:"Label,omitempty"`
	Score     *string `json:"Score,omitempty" xml:"Score,omitempty"`
	Timestamp *string `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	Url       *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultTerrorismResultTopList) String() string {
	return tea.Prettify(s)
}

func (s GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultTerrorismResultTopList) GoString() string {
	return s.String()
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultTerrorismResultTopList) SetLabel(v string) *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultTerrorismResultTopList {
	s.Label = &v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultTerrorismResultTopList) SetScore(v string) *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultTerrorismResultTopList {
	s.Score = &v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultTerrorismResultTopList) SetTimestamp(v string) *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultTerrorismResultTopList {
	s.Timestamp = &v
	return s
}

func (s *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultTerrorismResultTopList) SetUrl(v string) *GetAIMediaAuditJobResponseBodyMediaAuditJobDataVideoResultTerrorismResultTopList {
	s.Url = &v
	return s
}

type GetAIMediaAuditJobResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetAIMediaAuditJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetAIMediaAuditJobResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAIMediaAuditJobResponse) GoString() string {
	return s.String()
}

func (s *GetAIMediaAuditJobResponse) SetHeaders(v map[string]*string) *GetAIMediaAuditJobResponse {
	s.Headers = v
	return s
}

func (s *GetAIMediaAuditJobResponse) SetBody(v *GetAIMediaAuditJobResponseBody) *GetAIMediaAuditJobResponse {
	s.Body = v
	return s
}

type GetAITemplateRequest struct {
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s GetAITemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAITemplateRequest) GoString() string {
	return s.String()
}

func (s *GetAITemplateRequest) SetTemplateId(v string) *GetAITemplateRequest {
	s.TemplateId = &v
	return s
}

type GetAITemplateResponseBody struct {
	RequestId    *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TemplateInfo *GetAITemplateResponseBodyTemplateInfo `json:"TemplateInfo,omitempty" xml:"TemplateInfo,omitempty" type:"Struct"`
}

func (s GetAITemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAITemplateResponseBody) GoString() string {
	return s.String()
}

func (s *GetAITemplateResponseBody) SetRequestId(v string) *GetAITemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAITemplateResponseBody) SetTemplateInfo(v *GetAITemplateResponseBodyTemplateInfo) *GetAITemplateResponseBody {
	s.TemplateInfo = v
	return s
}

type GetAITemplateResponseBodyTemplateInfo struct {
	CreationTime   *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	IsDefault      *string `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	ModifyTime     *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	Source         *string `json:"Source,omitempty" xml:"Source,omitempty"`
	TemplateConfig *string `json:"TemplateConfig,omitempty" xml:"TemplateConfig,omitempty"`
	TemplateId     *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	TemplateName   *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	TemplateType   *string `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
}

func (s GetAITemplateResponseBodyTemplateInfo) String() string {
	return tea.Prettify(s)
}

func (s GetAITemplateResponseBodyTemplateInfo) GoString() string {
	return s.String()
}

func (s *GetAITemplateResponseBodyTemplateInfo) SetCreationTime(v string) *GetAITemplateResponseBodyTemplateInfo {
	s.CreationTime = &v
	return s
}

func (s *GetAITemplateResponseBodyTemplateInfo) SetIsDefault(v string) *GetAITemplateResponseBodyTemplateInfo {
	s.IsDefault = &v
	return s
}

func (s *GetAITemplateResponseBodyTemplateInfo) SetModifyTime(v string) *GetAITemplateResponseBodyTemplateInfo {
	s.ModifyTime = &v
	return s
}

func (s *GetAITemplateResponseBodyTemplateInfo) SetSource(v string) *GetAITemplateResponseBodyTemplateInfo {
	s.Source = &v
	return s
}

func (s *GetAITemplateResponseBodyTemplateInfo) SetTemplateConfig(v string) *GetAITemplateResponseBodyTemplateInfo {
	s.TemplateConfig = &v
	return s
}

func (s *GetAITemplateResponseBodyTemplateInfo) SetTemplateId(v string) *GetAITemplateResponseBodyTemplateInfo {
	s.TemplateId = &v
	return s
}

func (s *GetAITemplateResponseBodyTemplateInfo) SetTemplateName(v string) *GetAITemplateResponseBodyTemplateInfo {
	s.TemplateName = &v
	return s
}

func (s *GetAITemplateResponseBodyTemplateInfo) SetTemplateType(v string) *GetAITemplateResponseBodyTemplateInfo {
	s.TemplateType = &v
	return s
}

type GetAITemplateResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetAITemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetAITemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAITemplateResponse) GoString() string {
	return s.String()
}

func (s *GetAITemplateResponse) SetHeaders(v map[string]*string) *GetAITemplateResponse {
	s.Headers = v
	return s
}

func (s *GetAITemplateResponse) SetBody(v *GetAITemplateResponseBody) *GetAITemplateResponse {
	s.Body = v
	return s
}

type GetAIVideoTagResultRequest struct {
	MediaId              *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *string `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s GetAIVideoTagResultRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAIVideoTagResultRequest) GoString() string {
	return s.String()
}

func (s *GetAIVideoTagResultRequest) SetMediaId(v string) *GetAIVideoTagResultRequest {
	s.MediaId = &v
	return s
}

func (s *GetAIVideoTagResultRequest) SetOwnerAccount(v string) *GetAIVideoTagResultRequest {
	s.OwnerAccount = &v
	return s
}

func (s *GetAIVideoTagResultRequest) SetOwnerId(v string) *GetAIVideoTagResultRequest {
	s.OwnerId = &v
	return s
}

func (s *GetAIVideoTagResultRequest) SetResourceOwnerAccount(v string) *GetAIVideoTagResultRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetAIVideoTagResultRequest) SetResourceOwnerId(v string) *GetAIVideoTagResultRequest {
	s.ResourceOwnerId = &v
	return s
}

type GetAIVideoTagResultResponseBody struct {
	RequestId      *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	VideoTagResult *GetAIVideoTagResultResponseBodyVideoTagResult `json:"VideoTagResult,omitempty" xml:"VideoTagResult,omitempty" type:"Struct"`
}

func (s GetAIVideoTagResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAIVideoTagResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetAIVideoTagResultResponseBody) SetRequestId(v string) *GetAIVideoTagResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAIVideoTagResultResponseBody) SetVideoTagResult(v *GetAIVideoTagResultResponseBodyVideoTagResult) *GetAIVideoTagResultResponseBody {
	s.VideoTagResult = v
	return s
}

type GetAIVideoTagResultResponseBodyVideoTagResult struct {
	Category []*GetAIVideoTagResultResponseBodyVideoTagResultCategory `json:"Category,omitempty" xml:"Category,omitempty" type:"Repeated"`
	Keyword  []*GetAIVideoTagResultResponseBodyVideoTagResultKeyword  `json:"Keyword,omitempty" xml:"Keyword,omitempty" type:"Repeated"`
	Location []*GetAIVideoTagResultResponseBodyVideoTagResultLocation `json:"Location,omitempty" xml:"Location,omitempty" type:"Repeated"`
	Person   []*GetAIVideoTagResultResponseBodyVideoTagResultPerson   `json:"Person,omitempty" xml:"Person,omitempty" type:"Repeated"`
	Time     []*GetAIVideoTagResultResponseBodyVideoTagResultTime     `json:"Time,omitempty" xml:"Time,omitempty" type:"Repeated"`
}

func (s GetAIVideoTagResultResponseBodyVideoTagResult) String() string {
	return tea.Prettify(s)
}

func (s GetAIVideoTagResultResponseBodyVideoTagResult) GoString() string {
	return s.String()
}

func (s *GetAIVideoTagResultResponseBodyVideoTagResult) SetCategory(v []*GetAIVideoTagResultResponseBodyVideoTagResultCategory) *GetAIVideoTagResultResponseBodyVideoTagResult {
	s.Category = v
	return s
}

func (s *GetAIVideoTagResultResponseBodyVideoTagResult) SetKeyword(v []*GetAIVideoTagResultResponseBodyVideoTagResultKeyword) *GetAIVideoTagResultResponseBodyVideoTagResult {
	s.Keyword = v
	return s
}

func (s *GetAIVideoTagResultResponseBodyVideoTagResult) SetLocation(v []*GetAIVideoTagResultResponseBodyVideoTagResultLocation) *GetAIVideoTagResultResponseBodyVideoTagResult {
	s.Location = v
	return s
}

func (s *GetAIVideoTagResultResponseBodyVideoTagResult) SetPerson(v []*GetAIVideoTagResultResponseBodyVideoTagResultPerson) *GetAIVideoTagResultResponseBodyVideoTagResult {
	s.Person = v
	return s
}

func (s *GetAIVideoTagResultResponseBodyVideoTagResult) SetTime(v []*GetAIVideoTagResultResponseBodyVideoTagResultTime) *GetAIVideoTagResultResponseBodyVideoTagResult {
	s.Time = v
	return s
}

type GetAIVideoTagResultResponseBodyVideoTagResultCategory struct {
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
}

func (s GetAIVideoTagResultResponseBodyVideoTagResultCategory) String() string {
	return tea.Prettify(s)
}

func (s GetAIVideoTagResultResponseBodyVideoTagResultCategory) GoString() string {
	return s.String()
}

func (s *GetAIVideoTagResultResponseBodyVideoTagResultCategory) SetTag(v string) *GetAIVideoTagResultResponseBodyVideoTagResultCategory {
	s.Tag = &v
	return s
}

type GetAIVideoTagResultResponseBodyVideoTagResultKeyword struct {
	Tag   *string   `json:"Tag,omitempty" xml:"Tag,omitempty"`
	Times []*string `json:"Times,omitempty" xml:"Times,omitempty" type:"Repeated"`
}

func (s GetAIVideoTagResultResponseBodyVideoTagResultKeyword) String() string {
	return tea.Prettify(s)
}

func (s GetAIVideoTagResultResponseBodyVideoTagResultKeyword) GoString() string {
	return s.String()
}

func (s *GetAIVideoTagResultResponseBodyVideoTagResultKeyword) SetTag(v string) *GetAIVideoTagResultResponseBodyVideoTagResultKeyword {
	s.Tag = &v
	return s
}

func (s *GetAIVideoTagResultResponseBodyVideoTagResultKeyword) SetTimes(v []*string) *GetAIVideoTagResultResponseBodyVideoTagResultKeyword {
	s.Times = v
	return s
}

type GetAIVideoTagResultResponseBodyVideoTagResultLocation struct {
	Tag   *string   `json:"Tag,omitempty" xml:"Tag,omitempty"`
	Times []*string `json:"Times,omitempty" xml:"Times,omitempty" type:"Repeated"`
}

func (s GetAIVideoTagResultResponseBodyVideoTagResultLocation) String() string {
	return tea.Prettify(s)
}

func (s GetAIVideoTagResultResponseBodyVideoTagResultLocation) GoString() string {
	return s.String()
}

func (s *GetAIVideoTagResultResponseBodyVideoTagResultLocation) SetTag(v string) *GetAIVideoTagResultResponseBodyVideoTagResultLocation {
	s.Tag = &v
	return s
}

func (s *GetAIVideoTagResultResponseBodyVideoTagResultLocation) SetTimes(v []*string) *GetAIVideoTagResultResponseBodyVideoTagResultLocation {
	s.Times = v
	return s
}

type GetAIVideoTagResultResponseBodyVideoTagResultPerson struct {
	FaceUrl *string   `json:"FaceUrl,omitempty" xml:"FaceUrl,omitempty"`
	Tag     *string   `json:"Tag,omitempty" xml:"Tag,omitempty"`
	Times   []*string `json:"Times,omitempty" xml:"Times,omitempty" type:"Repeated"`
}

func (s GetAIVideoTagResultResponseBodyVideoTagResultPerson) String() string {
	return tea.Prettify(s)
}

func (s GetAIVideoTagResultResponseBodyVideoTagResultPerson) GoString() string {
	return s.String()
}

func (s *GetAIVideoTagResultResponseBodyVideoTagResultPerson) SetFaceUrl(v string) *GetAIVideoTagResultResponseBodyVideoTagResultPerson {
	s.FaceUrl = &v
	return s
}

func (s *GetAIVideoTagResultResponseBodyVideoTagResultPerson) SetTag(v string) *GetAIVideoTagResultResponseBodyVideoTagResultPerson {
	s.Tag = &v
	return s
}

func (s *GetAIVideoTagResultResponseBodyVideoTagResultPerson) SetTimes(v []*string) *GetAIVideoTagResultResponseBodyVideoTagResultPerson {
	s.Times = v
	return s
}

type GetAIVideoTagResultResponseBodyVideoTagResultTime struct {
	Tag   *string   `json:"Tag,omitempty" xml:"Tag,omitempty"`
	Times []*string `json:"Times,omitempty" xml:"Times,omitempty" type:"Repeated"`
}

func (s GetAIVideoTagResultResponseBodyVideoTagResultTime) String() string {
	return tea.Prettify(s)
}

func (s GetAIVideoTagResultResponseBodyVideoTagResultTime) GoString() string {
	return s.String()
}

func (s *GetAIVideoTagResultResponseBodyVideoTagResultTime) SetTag(v string) *GetAIVideoTagResultResponseBodyVideoTagResultTime {
	s.Tag = &v
	return s
}

func (s *GetAIVideoTagResultResponseBodyVideoTagResultTime) SetTimes(v []*string) *GetAIVideoTagResultResponseBodyVideoTagResultTime {
	s.Times = v
	return s
}

type GetAIVideoTagResultResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetAIVideoTagResultResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetAIVideoTagResultResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAIVideoTagResultResponse) GoString() string {
	return s.String()
}

func (s *GetAIVideoTagResultResponse) SetHeaders(v map[string]*string) *GetAIVideoTagResultResponse {
	s.Headers = v
	return s
}

func (s *GetAIVideoTagResultResponse) SetBody(v *GetAIVideoTagResultResponseBody) *GetAIVideoTagResultResponse {
	s.Body = v
	return s
}

type GetAppInfosRequest struct {
	AppIds *string `json:"AppIds,omitempty" xml:"AppIds,omitempty"`
}

func (s GetAppInfosRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAppInfosRequest) GoString() string {
	return s.String()
}

func (s *GetAppInfosRequest) SetAppIds(v string) *GetAppInfosRequest {
	s.AppIds = &v
	return s
}

type GetAppInfosResponseBody struct {
	AppInfoList    []*GetAppInfosResponseBodyAppInfoList `json:"AppInfoList,omitempty" xml:"AppInfoList,omitempty" type:"Repeated"`
	NonExistAppIds []*string                             `json:"NonExistAppIds,omitempty" xml:"NonExistAppIds,omitempty" type:"Repeated"`
	RequestId      *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAppInfosResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAppInfosResponseBody) GoString() string {
	return s.String()
}

func (s *GetAppInfosResponseBody) SetAppInfoList(v []*GetAppInfosResponseBodyAppInfoList) *GetAppInfosResponseBody {
	s.AppInfoList = v
	return s
}

func (s *GetAppInfosResponseBody) SetNonExistAppIds(v []*string) *GetAppInfosResponseBody {
	s.NonExistAppIds = v
	return s
}

func (s *GetAppInfosResponseBody) SetRequestId(v string) *GetAppInfosResponseBody {
	s.RequestId = &v
	return s
}

type GetAppInfosResponseBodyAppInfoList struct {
	AppId            *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AppName          *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	CreationTime     *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	Description      *string `json:"Description,omitempty" xml:"Description,omitempty"`
	ModificationTime *string `json:"ModificationTime,omitempty" xml:"ModificationTime,omitempty"`
	Status           *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Type             *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetAppInfosResponseBodyAppInfoList) String() string {
	return tea.Prettify(s)
}

func (s GetAppInfosResponseBodyAppInfoList) GoString() string {
	return s.String()
}

func (s *GetAppInfosResponseBodyAppInfoList) SetAppId(v string) *GetAppInfosResponseBodyAppInfoList {
	s.AppId = &v
	return s
}

func (s *GetAppInfosResponseBodyAppInfoList) SetAppName(v string) *GetAppInfosResponseBodyAppInfoList {
	s.AppName = &v
	return s
}

func (s *GetAppInfosResponseBodyAppInfoList) SetCreationTime(v string) *GetAppInfosResponseBodyAppInfoList {
	s.CreationTime = &v
	return s
}

func (s *GetAppInfosResponseBodyAppInfoList) SetDescription(v string) *GetAppInfosResponseBodyAppInfoList {
	s.Description = &v
	return s
}

func (s *GetAppInfosResponseBodyAppInfoList) SetModificationTime(v string) *GetAppInfosResponseBodyAppInfoList {
	s.ModificationTime = &v
	return s
}

func (s *GetAppInfosResponseBodyAppInfoList) SetStatus(v string) *GetAppInfosResponseBodyAppInfoList {
	s.Status = &v
	return s
}

func (s *GetAppInfosResponseBodyAppInfoList) SetType(v string) *GetAppInfosResponseBodyAppInfoList {
	s.Type = &v
	return s
}

type GetAppInfosResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetAppInfosResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetAppInfosResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAppInfosResponse) GoString() string {
	return s.String()
}

func (s *GetAppInfosResponse) SetHeaders(v map[string]*string) *GetAppInfosResponse {
	s.Headers = v
	return s
}

func (s *GetAppInfosResponse) SetBody(v *GetAppInfosResponseBody) *GetAppInfosResponse {
	s.Body = v
	return s
}

type GetAttachedMediaInfoRequest struct {
	AuthTimeout *int64  `json:"AuthTimeout,omitempty" xml:"AuthTimeout,omitempty"`
	MediaIds    *string `json:"MediaIds,omitempty" xml:"MediaIds,omitempty"`
	OutputType  *string `json:"OutputType,omitempty" xml:"OutputType,omitempty"`
}

func (s GetAttachedMediaInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAttachedMediaInfoRequest) GoString() string {
	return s.String()
}

func (s *GetAttachedMediaInfoRequest) SetAuthTimeout(v int64) *GetAttachedMediaInfoRequest {
	s.AuthTimeout = &v
	return s
}

func (s *GetAttachedMediaInfoRequest) SetMediaIds(v string) *GetAttachedMediaInfoRequest {
	s.MediaIds = &v
	return s
}

func (s *GetAttachedMediaInfoRequest) SetOutputType(v string) *GetAttachedMediaInfoRequest {
	s.OutputType = &v
	return s
}

type GetAttachedMediaInfoResponseBody struct {
	AttachedMediaList []*GetAttachedMediaInfoResponseBodyAttachedMediaList `json:"AttachedMediaList,omitempty" xml:"AttachedMediaList,omitempty" type:"Repeated"`
	NonExistMediaIds  []*string                                            `json:"NonExistMediaIds,omitempty" xml:"NonExistMediaIds,omitempty" type:"Repeated"`
	RequestId         *string                                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAttachedMediaInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAttachedMediaInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetAttachedMediaInfoResponseBody) SetAttachedMediaList(v []*GetAttachedMediaInfoResponseBodyAttachedMediaList) *GetAttachedMediaInfoResponseBody {
	s.AttachedMediaList = v
	return s
}

func (s *GetAttachedMediaInfoResponseBody) SetNonExistMediaIds(v []*string) *GetAttachedMediaInfoResponseBody {
	s.NonExistMediaIds = v
	return s
}

func (s *GetAttachedMediaInfoResponseBody) SetRequestId(v string) *GetAttachedMediaInfoResponseBody {
	s.RequestId = &v
	return s
}

type GetAttachedMediaInfoResponseBodyAttachedMediaList struct {
	AppId            *string                                                        `json:"AppId,omitempty" xml:"AppId,omitempty"`
	Categories       []*GetAttachedMediaInfoResponseBodyAttachedMediaListCategories `json:"Categories,omitempty" xml:"Categories,omitempty" type:"Repeated"`
	CreationTime     *string                                                        `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	Description      *string                                                        `json:"Description,omitempty" xml:"Description,omitempty"`
	MediaId          *string                                                        `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	ModificationTime *string                                                        `json:"ModificationTime,omitempty" xml:"ModificationTime,omitempty"`
	Status           *string                                                        `json:"Status,omitempty" xml:"Status,omitempty"`
	StorageLocation  *string                                                        `json:"StorageLocation,omitempty" xml:"StorageLocation,omitempty"`
	Tags             *string                                                        `json:"Tags,omitempty" xml:"Tags,omitempty"`
	Title            *string                                                        `json:"Title,omitempty" xml:"Title,omitempty"`
	Type             *string                                                        `json:"Type,omitempty" xml:"Type,omitempty"`
	URL              *string                                                        `json:"URL,omitempty" xml:"URL,omitempty"`
}

func (s GetAttachedMediaInfoResponseBodyAttachedMediaList) String() string {
	return tea.Prettify(s)
}

func (s GetAttachedMediaInfoResponseBodyAttachedMediaList) GoString() string {
	return s.String()
}

func (s *GetAttachedMediaInfoResponseBodyAttachedMediaList) SetAppId(v string) *GetAttachedMediaInfoResponseBodyAttachedMediaList {
	s.AppId = &v
	return s
}

func (s *GetAttachedMediaInfoResponseBodyAttachedMediaList) SetCategories(v []*GetAttachedMediaInfoResponseBodyAttachedMediaListCategories) *GetAttachedMediaInfoResponseBodyAttachedMediaList {
	s.Categories = v
	return s
}

func (s *GetAttachedMediaInfoResponseBodyAttachedMediaList) SetCreationTime(v string) *GetAttachedMediaInfoResponseBodyAttachedMediaList {
	s.CreationTime = &v
	return s
}

func (s *GetAttachedMediaInfoResponseBodyAttachedMediaList) SetDescription(v string) *GetAttachedMediaInfoResponseBodyAttachedMediaList {
	s.Description = &v
	return s
}

func (s *GetAttachedMediaInfoResponseBodyAttachedMediaList) SetMediaId(v string) *GetAttachedMediaInfoResponseBodyAttachedMediaList {
	s.MediaId = &v
	return s
}

func (s *GetAttachedMediaInfoResponseBodyAttachedMediaList) SetModificationTime(v string) *GetAttachedMediaInfoResponseBodyAttachedMediaList {
	s.ModificationTime = &v
	return s
}

func (s *GetAttachedMediaInfoResponseBodyAttachedMediaList) SetStatus(v string) *GetAttachedMediaInfoResponseBodyAttachedMediaList {
	s.Status = &v
	return s
}

func (s *GetAttachedMediaInfoResponseBodyAttachedMediaList) SetStorageLocation(v string) *GetAttachedMediaInfoResponseBodyAttachedMediaList {
	s.StorageLocation = &v
	return s
}

func (s *GetAttachedMediaInfoResponseBodyAttachedMediaList) SetTags(v string) *GetAttachedMediaInfoResponseBodyAttachedMediaList {
	s.Tags = &v
	return s
}

func (s *GetAttachedMediaInfoResponseBodyAttachedMediaList) SetTitle(v string) *GetAttachedMediaInfoResponseBodyAttachedMediaList {
	s.Title = &v
	return s
}

func (s *GetAttachedMediaInfoResponseBodyAttachedMediaList) SetType(v string) *GetAttachedMediaInfoResponseBodyAttachedMediaList {
	s.Type = &v
	return s
}

func (s *GetAttachedMediaInfoResponseBodyAttachedMediaList) SetURL(v string) *GetAttachedMediaInfoResponseBodyAttachedMediaList {
	s.URL = &v
	return s
}

type GetAttachedMediaInfoResponseBodyAttachedMediaListCategories struct {
	CateId   *int64  `json:"CateId,omitempty" xml:"CateId,omitempty"`
	CateName *string `json:"CateName,omitempty" xml:"CateName,omitempty"`
	Level    *int64  `json:"Level,omitempty" xml:"Level,omitempty"`
	ParentId *int64  `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
}

func (s GetAttachedMediaInfoResponseBodyAttachedMediaListCategories) String() string {
	return tea.Prettify(s)
}

func (s GetAttachedMediaInfoResponseBodyAttachedMediaListCategories) GoString() string {
	return s.String()
}

func (s *GetAttachedMediaInfoResponseBodyAttachedMediaListCategories) SetCateId(v int64) *GetAttachedMediaInfoResponseBodyAttachedMediaListCategories {
	s.CateId = &v
	return s
}

func (s *GetAttachedMediaInfoResponseBodyAttachedMediaListCategories) SetCateName(v string) *GetAttachedMediaInfoResponseBodyAttachedMediaListCategories {
	s.CateName = &v
	return s
}

func (s *GetAttachedMediaInfoResponseBodyAttachedMediaListCategories) SetLevel(v int64) *GetAttachedMediaInfoResponseBodyAttachedMediaListCategories {
	s.Level = &v
	return s
}

func (s *GetAttachedMediaInfoResponseBodyAttachedMediaListCategories) SetParentId(v int64) *GetAttachedMediaInfoResponseBodyAttachedMediaListCategories {
	s.ParentId = &v
	return s
}

type GetAttachedMediaInfoResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetAttachedMediaInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetAttachedMediaInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAttachedMediaInfoResponse) GoString() string {
	return s.String()
}

func (s *GetAttachedMediaInfoResponse) SetHeaders(v map[string]*string) *GetAttachedMediaInfoResponse {
	s.Headers = v
	return s
}

func (s *GetAttachedMediaInfoResponse) SetBody(v *GetAttachedMediaInfoResponseBody) *GetAttachedMediaInfoResponse {
	s.Body = v
	return s
}

type GetAuditHistoryRequest struct {
	PageNo   *int64  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SortBy   *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	VideoId  *string `json:"VideoId,omitempty" xml:"VideoId,omitempty"`
}

func (s GetAuditHistoryRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAuditHistoryRequest) GoString() string {
	return s.String()
}

func (s *GetAuditHistoryRequest) SetPageNo(v int64) *GetAuditHistoryRequest {
	s.PageNo = &v
	return s
}

func (s *GetAuditHistoryRequest) SetPageSize(v int64) *GetAuditHistoryRequest {
	s.PageSize = &v
	return s
}

func (s *GetAuditHistoryRequest) SetSortBy(v string) *GetAuditHistoryRequest {
	s.SortBy = &v
	return s
}

func (s *GetAuditHistoryRequest) SetVideoId(v string) *GetAuditHistoryRequest {
	s.VideoId = &v
	return s
}

type GetAuditHistoryResponseBody struct {
	Histories []*GetAuditHistoryResponseBodyHistories `json:"Histories,omitempty" xml:"Histories,omitempty" type:"Repeated"`
	RequestId *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Status    *string                                 `json:"Status,omitempty" xml:"Status,omitempty"`
	Total     *int64                                  `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetAuditHistoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAuditHistoryResponseBody) GoString() string {
	return s.String()
}

func (s *GetAuditHistoryResponseBody) SetHistories(v []*GetAuditHistoryResponseBodyHistories) *GetAuditHistoryResponseBody {
	s.Histories = v
	return s
}

func (s *GetAuditHistoryResponseBody) SetRequestId(v string) *GetAuditHistoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAuditHistoryResponseBody) SetStatus(v string) *GetAuditHistoryResponseBody {
	s.Status = &v
	return s
}

func (s *GetAuditHistoryResponseBody) SetTotal(v int64) *GetAuditHistoryResponseBody {
	s.Total = &v
	return s
}

type GetAuditHistoryResponseBodyHistories struct {
	Auditor      *string `json:"Auditor,omitempty" xml:"Auditor,omitempty"`
	Comment      *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	Reason       *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetAuditHistoryResponseBodyHistories) String() string {
	return tea.Prettify(s)
}

func (s GetAuditHistoryResponseBodyHistories) GoString() string {
	return s.String()
}

func (s *GetAuditHistoryResponseBodyHistories) SetAuditor(v string) *GetAuditHistoryResponseBodyHistories {
	s.Auditor = &v
	return s
}

func (s *GetAuditHistoryResponseBodyHistories) SetComment(v string) *GetAuditHistoryResponseBodyHistories {
	s.Comment = &v
	return s
}

func (s *GetAuditHistoryResponseBodyHistories) SetCreationTime(v string) *GetAuditHistoryResponseBodyHistories {
	s.CreationTime = &v
	return s
}

func (s *GetAuditHistoryResponseBodyHistories) SetReason(v string) *GetAuditHistoryResponseBodyHistories {
	s.Reason = &v
	return s
}

func (s *GetAuditHistoryResponseBodyHistories) SetStatus(v string) *GetAuditHistoryResponseBodyHistories {
	s.Status = &v
	return s
}

type GetAuditHistoryResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetAuditHistoryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetAuditHistoryResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAuditHistoryResponse) GoString() string {
	return s.String()
}

func (s *GetAuditHistoryResponse) SetHeaders(v map[string]*string) *GetAuditHistoryResponse {
	s.Headers = v
	return s
}

func (s *GetAuditHistoryResponse) SetBody(v *GetAuditHistoryResponseBody) *GetAuditHistoryResponse {
	s.Body = v
	return s
}

type GetCategoriesRequest struct {
	CateId   *int64  `json:"CateId,omitempty" xml:"CateId,omitempty"`
	PageNo   *int64  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SortBy   *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	Type     *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetCategoriesRequest) String() string {
	return tea.Prettify(s)
}

func (s GetCategoriesRequest) GoString() string {
	return s.String()
}

func (s *GetCategoriesRequest) SetCateId(v int64) *GetCategoriesRequest {
	s.CateId = &v
	return s
}

func (s *GetCategoriesRequest) SetPageNo(v int64) *GetCategoriesRequest {
	s.PageNo = &v
	return s
}

func (s *GetCategoriesRequest) SetPageSize(v int64) *GetCategoriesRequest {
	s.PageSize = &v
	return s
}

func (s *GetCategoriesRequest) SetSortBy(v string) *GetCategoriesRequest {
	s.SortBy = &v
	return s
}

func (s *GetCategoriesRequest) SetType(v string) *GetCategoriesRequest {
	s.Type = &v
	return s
}

type GetCategoriesResponseBody struct {
	Category      *GetCategoriesResponseBodyCategory      `json:"Category,omitempty" xml:"Category,omitempty" type:"Struct"`
	RequestId     *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SubCategories *GetCategoriesResponseBodySubCategories `json:"SubCategories,omitempty" xml:"SubCategories,omitempty" type:"Struct"`
	SubTotal      *int64                                  `json:"SubTotal,omitempty" xml:"SubTotal,omitempty"`
}

func (s GetCategoriesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetCategoriesResponseBody) GoString() string {
	return s.String()
}

func (s *GetCategoriesResponseBody) SetCategory(v *GetCategoriesResponseBodyCategory) *GetCategoriesResponseBody {
	s.Category = v
	return s
}

func (s *GetCategoriesResponseBody) SetRequestId(v string) *GetCategoriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCategoriesResponseBody) SetSubCategories(v *GetCategoriesResponseBodySubCategories) *GetCategoriesResponseBody {
	s.SubCategories = v
	return s
}

func (s *GetCategoriesResponseBody) SetSubTotal(v int64) *GetCategoriesResponseBody {
	s.SubTotal = &v
	return s
}

type GetCategoriesResponseBodyCategory struct {
	CateId   *int64  `json:"CateId,omitempty" xml:"CateId,omitempty"`
	CateName *string `json:"CateName,omitempty" xml:"CateName,omitempty"`
	Level    *int64  `json:"Level,omitempty" xml:"Level,omitempty"`
	ParentId *int64  `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
	Type     *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetCategoriesResponseBodyCategory) String() string {
	return tea.Prettify(s)
}

func (s GetCategoriesResponseBodyCategory) GoString() string {
	return s.String()
}

func (s *GetCategoriesResponseBodyCategory) SetCateId(v int64) *GetCategoriesResponseBodyCategory {
	s.CateId = &v
	return s
}

func (s *GetCategoriesResponseBodyCategory) SetCateName(v string) *GetCategoriesResponseBodyCategory {
	s.CateName = &v
	return s
}

func (s *GetCategoriesResponseBodyCategory) SetLevel(v int64) *GetCategoriesResponseBodyCategory {
	s.Level = &v
	return s
}

func (s *GetCategoriesResponseBodyCategory) SetParentId(v int64) *GetCategoriesResponseBodyCategory {
	s.ParentId = &v
	return s
}

func (s *GetCategoriesResponseBodyCategory) SetType(v string) *GetCategoriesResponseBodyCategory {
	s.Type = &v
	return s
}

type GetCategoriesResponseBodySubCategories struct {
	Category []*GetCategoriesResponseBodySubCategoriesCategory `json:"Category,omitempty" xml:"Category,omitempty" type:"Repeated"`
}

func (s GetCategoriesResponseBodySubCategories) String() string {
	return tea.Prettify(s)
}

func (s GetCategoriesResponseBodySubCategories) GoString() string {
	return s.String()
}

func (s *GetCategoriesResponseBodySubCategories) SetCategory(v []*GetCategoriesResponseBodySubCategoriesCategory) *GetCategoriesResponseBodySubCategories {
	s.Category = v
	return s
}

type GetCategoriesResponseBodySubCategoriesCategory struct {
	CateId   *int64  `json:"CateId,omitempty" xml:"CateId,omitempty"`
	CateName *string `json:"CateName,omitempty" xml:"CateName,omitempty"`
	Level    *int64  `json:"Level,omitempty" xml:"Level,omitempty"`
	ParentId *int64  `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
	SubTotal *int64  `json:"SubTotal,omitempty" xml:"SubTotal,omitempty"`
	Type     *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetCategoriesResponseBodySubCategoriesCategory) String() string {
	return tea.Prettify(s)
}

func (s GetCategoriesResponseBodySubCategoriesCategory) GoString() string {
	return s.String()
}

func (s *GetCategoriesResponseBodySubCategoriesCategory) SetCateId(v int64) *GetCategoriesResponseBodySubCategoriesCategory {
	s.CateId = &v
	return s
}

func (s *GetCategoriesResponseBodySubCategoriesCategory) SetCateName(v string) *GetCategoriesResponseBodySubCategoriesCategory {
	s.CateName = &v
	return s
}

func (s *GetCategoriesResponseBodySubCategoriesCategory) SetLevel(v int64) *GetCategoriesResponseBodySubCategoriesCategory {
	s.Level = &v
	return s
}

func (s *GetCategoriesResponseBodySubCategoriesCategory) SetParentId(v int64) *GetCategoriesResponseBodySubCategoriesCategory {
	s.ParentId = &v
	return s
}

func (s *GetCategoriesResponseBodySubCategoriesCategory) SetSubTotal(v int64) *GetCategoriesResponseBodySubCategoriesCategory {
	s.SubTotal = &v
	return s
}

func (s *GetCategoriesResponseBodySubCategoriesCategory) SetType(v string) *GetCategoriesResponseBodySubCategoriesCategory {
	s.Type = &v
	return s
}

type GetCategoriesResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetCategoriesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetCategoriesResponse) String() string {
	return tea.Prettify(s)
}

func (s GetCategoriesResponse) GoString() string {
	return s.String()
}

func (s *GetCategoriesResponse) SetHeaders(v map[string]*string) *GetCategoriesResponse {
	s.Headers = v
	return s
}

func (s *GetCategoriesResponse) SetBody(v *GetCategoriesResponseBody) *GetCategoriesResponse {
	s.Body = v
	return s
}

type GetDefaultAITemplateRequest struct {
	TemplateType *string `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
}

func (s GetDefaultAITemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDefaultAITemplateRequest) GoString() string {
	return s.String()
}

func (s *GetDefaultAITemplateRequest) SetTemplateType(v string) *GetDefaultAITemplateRequest {
	s.TemplateType = &v
	return s
}

type GetDefaultAITemplateResponseBody struct {
	RequestId    *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TemplateInfo *GetDefaultAITemplateResponseBodyTemplateInfo `json:"TemplateInfo,omitempty" xml:"TemplateInfo,omitempty" type:"Struct"`
}

func (s GetDefaultAITemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDefaultAITemplateResponseBody) GoString() string {
	return s.String()
}

func (s *GetDefaultAITemplateResponseBody) SetRequestId(v string) *GetDefaultAITemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDefaultAITemplateResponseBody) SetTemplateInfo(v *GetDefaultAITemplateResponseBodyTemplateInfo) *GetDefaultAITemplateResponseBody {
	s.TemplateInfo = v
	return s
}

type GetDefaultAITemplateResponseBodyTemplateInfo struct {
	CreationTime   *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	IsDefault      *string `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	ModifyTime     *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	Source         *string `json:"Source,omitempty" xml:"Source,omitempty"`
	TemplateConfig *string `json:"TemplateConfig,omitempty" xml:"TemplateConfig,omitempty"`
	TemplateId     *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	TemplateName   *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	TemplateType   *string `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
}

func (s GetDefaultAITemplateResponseBodyTemplateInfo) String() string {
	return tea.Prettify(s)
}

func (s GetDefaultAITemplateResponseBodyTemplateInfo) GoString() string {
	return s.String()
}

func (s *GetDefaultAITemplateResponseBodyTemplateInfo) SetCreationTime(v string) *GetDefaultAITemplateResponseBodyTemplateInfo {
	s.CreationTime = &v
	return s
}

func (s *GetDefaultAITemplateResponseBodyTemplateInfo) SetIsDefault(v string) *GetDefaultAITemplateResponseBodyTemplateInfo {
	s.IsDefault = &v
	return s
}

func (s *GetDefaultAITemplateResponseBodyTemplateInfo) SetModifyTime(v string) *GetDefaultAITemplateResponseBodyTemplateInfo {
	s.ModifyTime = &v
	return s
}

func (s *GetDefaultAITemplateResponseBodyTemplateInfo) SetSource(v string) *GetDefaultAITemplateResponseBodyTemplateInfo {
	s.Source = &v
	return s
}

func (s *GetDefaultAITemplateResponseBodyTemplateInfo) SetTemplateConfig(v string) *GetDefaultAITemplateResponseBodyTemplateInfo {
	s.TemplateConfig = &v
	return s
}

func (s *GetDefaultAITemplateResponseBodyTemplateInfo) SetTemplateId(v string) *GetDefaultAITemplateResponseBodyTemplateInfo {
	s.TemplateId = &v
	return s
}

func (s *GetDefaultAITemplateResponseBodyTemplateInfo) SetTemplateName(v string) *GetDefaultAITemplateResponseBodyTemplateInfo {
	s.TemplateName = &v
	return s
}

func (s *GetDefaultAITemplateResponseBodyTemplateInfo) SetTemplateType(v string) *GetDefaultAITemplateResponseBodyTemplateInfo {
	s.TemplateType = &v
	return s
}

type GetDefaultAITemplateResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetDefaultAITemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetDefaultAITemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDefaultAITemplateResponse) GoString() string {
	return s.String()
}

func (s *GetDefaultAITemplateResponse) SetHeaders(v map[string]*string) *GetDefaultAITemplateResponse {
	s.Headers = v
	return s
}

func (s *GetDefaultAITemplateResponse) SetBody(v *GetDefaultAITemplateResponseBody) *GetDefaultAITemplateResponse {
	s.Body = v
	return s
}

type GetDetectionJobRequest struct {
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s GetDetectionJobRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDetectionJobRequest) GoString() string {
	return s.String()
}

func (s *GetDetectionJobRequest) SetJobId(v string) *GetDetectionJobRequest {
	s.JobId = &v
	return s
}

type GetDetectionJobResponseBody struct {
	DetectionJob *GetDetectionJobResponseBodyDetectionJob `json:"DetectionJob,omitempty" xml:"DetectionJob,omitempty" type:"Struct"`
	RequestId    *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetDetectionJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDetectionJobResponseBody) GoString() string {
	return s.String()
}

func (s *GetDetectionJobResponseBody) SetDetectionJob(v *GetDetectionJobResponseBodyDetectionJob) *GetDetectionJobResponseBody {
	s.DetectionJob = v
	return s
}

func (s *GetDetectionJobResponseBody) SetRequestId(v string) *GetDetectionJobResponseBody {
	s.RequestId = &v
	return s
}

type GetDetectionJobResponseBodyDetectionJob struct {
	BeginTime          *string `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	CopyrightBeginTime *string `json:"CopyrightBeginTime,omitempty" xml:"CopyrightBeginTime,omitempty"`
	CopyrightEndTime   *string `json:"CopyrightEndTime,omitempty" xml:"CopyrightEndTime,omitempty"`
	CopyrightFile      *string `json:"CopyrightFile,omitempty" xml:"CopyrightFile,omitempty"`
	CopyrightStatus    *string `json:"CopyrightStatus,omitempty" xml:"CopyrightStatus,omitempty"`
	CreateTime         *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	EndTime            *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	JobId              *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	ModifyTime         *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	TemplateId         *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	VideoId            *string `json:"VideoId,omitempty" xml:"VideoId,omitempty"`
	WhitelistUrls      *string `json:"WhitelistUrls,omitempty" xml:"WhitelistUrls,omitempty"`
}

func (s GetDetectionJobResponseBodyDetectionJob) String() string {
	return tea.Prettify(s)
}

func (s GetDetectionJobResponseBodyDetectionJob) GoString() string {
	return s.String()
}

func (s *GetDetectionJobResponseBodyDetectionJob) SetBeginTime(v string) *GetDetectionJobResponseBodyDetectionJob {
	s.BeginTime = &v
	return s
}

func (s *GetDetectionJobResponseBodyDetectionJob) SetCopyrightBeginTime(v string) *GetDetectionJobResponseBodyDetectionJob {
	s.CopyrightBeginTime = &v
	return s
}

func (s *GetDetectionJobResponseBodyDetectionJob) SetCopyrightEndTime(v string) *GetDetectionJobResponseBodyDetectionJob {
	s.CopyrightEndTime = &v
	return s
}

func (s *GetDetectionJobResponseBodyDetectionJob) SetCopyrightFile(v string) *GetDetectionJobResponseBodyDetectionJob {
	s.CopyrightFile = &v
	return s
}

func (s *GetDetectionJobResponseBodyDetectionJob) SetCopyrightStatus(v string) *GetDetectionJobResponseBodyDetectionJob {
	s.CopyrightStatus = &v
	return s
}

func (s *GetDetectionJobResponseBodyDetectionJob) SetCreateTime(v string) *GetDetectionJobResponseBodyDetectionJob {
	s.CreateTime = &v
	return s
}

func (s *GetDetectionJobResponseBodyDetectionJob) SetEndTime(v string) *GetDetectionJobResponseBodyDetectionJob {
	s.EndTime = &v
	return s
}

func (s *GetDetectionJobResponseBodyDetectionJob) SetJobId(v string) *GetDetectionJobResponseBodyDetectionJob {
	s.JobId = &v
	return s
}

func (s *GetDetectionJobResponseBodyDetectionJob) SetModifyTime(v string) *GetDetectionJobResponseBodyDetectionJob {
	s.ModifyTime = &v
	return s
}

func (s *GetDetectionJobResponseBodyDetectionJob) SetTemplateId(v string) *GetDetectionJobResponseBodyDetectionJob {
	s.TemplateId = &v
	return s
}

func (s *GetDetectionJobResponseBodyDetectionJob) SetVideoId(v string) *GetDetectionJobResponseBodyDetectionJob {
	s.VideoId = &v
	return s
}

func (s *GetDetectionJobResponseBodyDetectionJob) SetWhitelistUrls(v string) *GetDetectionJobResponseBodyDetectionJob {
	s.WhitelistUrls = &v
	return s
}

type GetDetectionJobResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetDetectionJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetDetectionJobResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDetectionJobResponse) GoString() string {
	return s.String()
}

func (s *GetDetectionJobResponse) SetHeaders(v map[string]*string) *GetDetectionJobResponse {
	s.Headers = v
	return s
}

func (s *GetDetectionJobResponse) SetBody(v *GetDetectionJobResponseBody) *GetDetectionJobResponse {
	s.Body = v
	return s
}

type GetDetectionResultRequest struct {
	CountByPage     *int64  `json:"CountByPage,omitempty" xml:"CountByPage,omitempty"`
	Desensitization *bool   `json:"Desensitization,omitempty" xml:"Desensitization,omitempty"`
	JobId           *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	Page            *int64  `json:"Page,omitempty" xml:"Page,omitempty"`
	Status          *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetDetectionResultRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDetectionResultRequest) GoString() string {
	return s.String()
}

func (s *GetDetectionResultRequest) SetCountByPage(v int64) *GetDetectionResultRequest {
	s.CountByPage = &v
	return s
}

func (s *GetDetectionResultRequest) SetDesensitization(v bool) *GetDetectionResultRequest {
	s.Desensitization = &v
	return s
}

func (s *GetDetectionResultRequest) SetJobId(v string) *GetDetectionResultRequest {
	s.JobId = &v
	return s
}

func (s *GetDetectionResultRequest) SetPage(v int64) *GetDetectionResultRequest {
	s.Page = &v
	return s
}

func (s *GetDetectionResultRequest) SetStatus(v string) *GetDetectionResultRequest {
	s.Status = &v
	return s
}

type GetDetectionResultResponseBody struct {
	DetectionResultList []*GetDetectionResultResponseBodyDetectionResultList `json:"DetectionResultList,omitempty" xml:"DetectionResultList,omitempty" type:"Repeated"`
	RequestId           *string                                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetDetectionResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDetectionResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetDetectionResultResponseBody) SetDetectionResultList(v []*GetDetectionResultResponseBodyDetectionResultList) *GetDetectionResultResponseBody {
	s.DetectionResultList = v
	return s
}

func (s *GetDetectionResultResponseBody) SetRequestId(v string) *GetDetectionResultResponseBody {
	s.RequestId = &v
	return s
}

type GetDetectionResultResponseBodyDetectionResultList struct {
	CollectionTitle *string `json:"CollectionTitle,omitempty" xml:"CollectionTitle,omitempty"`
	CollectionUrl   *string `json:"CollectionUrl,omitempty" xml:"CollectionUrl,omitempty"`
	ContentType     *string `json:"ContentType,omitempty" xml:"ContentType,omitempty"`
	CreateTime      *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	ModifyTime      *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	Platform        *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	Status          *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Uploader        *string `json:"Uploader,omitempty" xml:"Uploader,omitempty"`
}

func (s GetDetectionResultResponseBodyDetectionResultList) String() string {
	return tea.Prettify(s)
}

func (s GetDetectionResultResponseBodyDetectionResultList) GoString() string {
	return s.String()
}

func (s *GetDetectionResultResponseBodyDetectionResultList) SetCollectionTitle(v string) *GetDetectionResultResponseBodyDetectionResultList {
	s.CollectionTitle = &v
	return s
}

func (s *GetDetectionResultResponseBodyDetectionResultList) SetCollectionUrl(v string) *GetDetectionResultResponseBodyDetectionResultList {
	s.CollectionUrl = &v
	return s
}

func (s *GetDetectionResultResponseBodyDetectionResultList) SetContentType(v string) *GetDetectionResultResponseBodyDetectionResultList {
	s.ContentType = &v
	return s
}

func (s *GetDetectionResultResponseBodyDetectionResultList) SetCreateTime(v string) *GetDetectionResultResponseBodyDetectionResultList {
	s.CreateTime = &v
	return s
}

func (s *GetDetectionResultResponseBodyDetectionResultList) SetModifyTime(v string) *GetDetectionResultResponseBodyDetectionResultList {
	s.ModifyTime = &v
	return s
}

func (s *GetDetectionResultResponseBodyDetectionResultList) SetPlatform(v string) *GetDetectionResultResponseBodyDetectionResultList {
	s.Platform = &v
	return s
}

func (s *GetDetectionResultResponseBodyDetectionResultList) SetStatus(v string) *GetDetectionResultResponseBodyDetectionResultList {
	s.Status = &v
	return s
}

func (s *GetDetectionResultResponseBodyDetectionResultList) SetUploader(v string) *GetDetectionResultResponseBodyDetectionResultList {
	s.Uploader = &v
	return s
}

type GetDetectionResultResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetDetectionResultResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetDetectionResultResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDetectionResultResponse) GoString() string {
	return s.String()
}

func (s *GetDetectionResultResponse) SetHeaders(v map[string]*string) *GetDetectionResultResponse {
	s.Headers = v
	return s
}

func (s *GetDetectionResultResponse) SetBody(v *GetDetectionResultResponseBody) *GetDetectionResultResponse {
	s.Body = v
	return s
}

type GetDetectionTemplateRequest struct {
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s GetDetectionTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDetectionTemplateRequest) GoString() string {
	return s.String()
}

func (s *GetDetectionTemplateRequest) SetTemplateId(v string) *GetDetectionTemplateRequest {
	s.TemplateId = &v
	return s
}

type GetDetectionTemplateResponseBody struct {
	DetectionTemplate *GetDetectionTemplateResponseBodyDetectionTemplate `json:"DetectionTemplate,omitempty" xml:"DetectionTemplate,omitempty" type:"Struct"`
	RequestId         *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetDetectionTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDetectionTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *GetDetectionTemplateResponseBody) SetDetectionTemplate(v *GetDetectionTemplateResponseBodyDetectionTemplate) *GetDetectionTemplateResponseBody {
	s.DetectionTemplate = v
	return s
}

func (s *GetDetectionTemplateResponseBody) SetRequestId(v string) *GetDetectionTemplateResponseBody {
	s.RequestId = &v
	return s
}

type GetDetectionTemplateResponseBodyDetectionTemplate struct {
	CreateTime   *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	ModifyTime   *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	Period       *string `json:"Period,omitempty" xml:"Period,omitempty"`
	Platform     *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	TemplateId   *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	UserId       *int64  `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s GetDetectionTemplateResponseBodyDetectionTemplate) String() string {
	return tea.Prettify(s)
}

func (s GetDetectionTemplateResponseBodyDetectionTemplate) GoString() string {
	return s.String()
}

func (s *GetDetectionTemplateResponseBodyDetectionTemplate) SetCreateTime(v string) *GetDetectionTemplateResponseBodyDetectionTemplate {
	s.CreateTime = &v
	return s
}

func (s *GetDetectionTemplateResponseBodyDetectionTemplate) SetModifyTime(v string) *GetDetectionTemplateResponseBodyDetectionTemplate {
	s.ModifyTime = &v
	return s
}

func (s *GetDetectionTemplateResponseBodyDetectionTemplate) SetPeriod(v string) *GetDetectionTemplateResponseBodyDetectionTemplate {
	s.Period = &v
	return s
}

func (s *GetDetectionTemplateResponseBodyDetectionTemplate) SetPlatform(v string) *GetDetectionTemplateResponseBodyDetectionTemplate {
	s.Platform = &v
	return s
}

func (s *GetDetectionTemplateResponseBodyDetectionTemplate) SetTemplateId(v string) *GetDetectionTemplateResponseBodyDetectionTemplate {
	s.TemplateId = &v
	return s
}

func (s *GetDetectionTemplateResponseBodyDetectionTemplate) SetTemplateName(v string) *GetDetectionTemplateResponseBodyDetectionTemplate {
	s.TemplateName = &v
	return s
}

func (s *GetDetectionTemplateResponseBodyDetectionTemplate) SetUserId(v int64) *GetDetectionTemplateResponseBodyDetectionTemplate {
	s.UserId = &v
	return s
}

type GetDetectionTemplateResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetDetectionTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetDetectionTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDetectionTemplateResponse) GoString() string {
	return s.String()
}

func (s *GetDetectionTemplateResponse) SetHeaders(v map[string]*string) *GetDetectionTemplateResponse {
	s.Headers = v
	return s
}

func (s *GetDetectionTemplateResponse) SetBody(v *GetDetectionTemplateResponseBody) *GetDetectionTemplateResponse {
	s.Body = v
	return s
}

type GetEditingProjectRequest struct {
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ProjectId            *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *string `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s GetEditingProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s GetEditingProjectRequest) GoString() string {
	return s.String()
}

func (s *GetEditingProjectRequest) SetOwnerAccount(v string) *GetEditingProjectRequest {
	s.OwnerAccount = &v
	return s
}

func (s *GetEditingProjectRequest) SetOwnerId(v string) *GetEditingProjectRequest {
	s.OwnerId = &v
	return s
}

func (s *GetEditingProjectRequest) SetProjectId(v string) *GetEditingProjectRequest {
	s.ProjectId = &v
	return s
}

func (s *GetEditingProjectRequest) SetResourceOwnerAccount(v string) *GetEditingProjectRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetEditingProjectRequest) SetResourceOwnerId(v string) *GetEditingProjectRequest {
	s.ResourceOwnerId = &v
	return s
}

type GetEditingProjectResponseBody struct {
	Project   *GetEditingProjectResponseBodyProject `json:"Project,omitempty" xml:"Project,omitempty" type:"Struct"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetEditingProjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetEditingProjectResponseBody) GoString() string {
	return s.String()
}

func (s *GetEditingProjectResponseBody) SetProject(v *GetEditingProjectResponseBodyProject) *GetEditingProjectResponseBody {
	s.Project = v
	return s
}

func (s *GetEditingProjectResponseBody) SetRequestId(v string) *GetEditingProjectResponseBody {
	s.RequestId = &v
	return s
}

type GetEditingProjectResponseBodyProject struct {
	CoverURL        *string `json:"CoverURL,omitempty" xml:"CoverURL,omitempty"`
	CreationTime    *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	Description     *string `json:"Description,omitempty" xml:"Description,omitempty"`
	ModifiedTime    *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	ProjectId       *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Status          *string `json:"Status,omitempty" xml:"Status,omitempty"`
	StorageLocation *string `json:"StorageLocation,omitempty" xml:"StorageLocation,omitempty"`
	Timeline        *string `json:"Timeline,omitempty" xml:"Timeline,omitempty"`
	Title           *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s GetEditingProjectResponseBodyProject) String() string {
	return tea.Prettify(s)
}

func (s GetEditingProjectResponseBodyProject) GoString() string {
	return s.String()
}

func (s *GetEditingProjectResponseBodyProject) SetCoverURL(v string) *GetEditingProjectResponseBodyProject {
	s.CoverURL = &v
	return s
}

func (s *GetEditingProjectResponseBodyProject) SetCreationTime(v string) *GetEditingProjectResponseBodyProject {
	s.CreationTime = &v
	return s
}

func (s *GetEditingProjectResponseBodyProject) SetDescription(v string) *GetEditingProjectResponseBodyProject {
	s.Description = &v
	return s
}

func (s *GetEditingProjectResponseBodyProject) SetModifiedTime(v string) *GetEditingProjectResponseBodyProject {
	s.ModifiedTime = &v
	return s
}

func (s *GetEditingProjectResponseBodyProject) SetProjectId(v string) *GetEditingProjectResponseBodyProject {
	s.ProjectId = &v
	return s
}

func (s *GetEditingProjectResponseBodyProject) SetRegionId(v string) *GetEditingProjectResponseBodyProject {
	s.RegionId = &v
	return s
}

func (s *GetEditingProjectResponseBodyProject) SetStatus(v string) *GetEditingProjectResponseBodyProject {
	s.Status = &v
	return s
}

func (s *GetEditingProjectResponseBodyProject) SetStorageLocation(v string) *GetEditingProjectResponseBodyProject {
	s.StorageLocation = &v
	return s
}

func (s *GetEditingProjectResponseBodyProject) SetTimeline(v string) *GetEditingProjectResponseBodyProject {
	s.Timeline = &v
	return s
}

func (s *GetEditingProjectResponseBodyProject) SetTitle(v string) *GetEditingProjectResponseBodyProject {
	s.Title = &v
	return s
}

type GetEditingProjectResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetEditingProjectResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetEditingProjectResponse) String() string {
	return tea.Prettify(s)
}

func (s GetEditingProjectResponse) GoString() string {
	return s.String()
}

func (s *GetEditingProjectResponse) SetHeaders(v map[string]*string) *GetEditingProjectResponse {
	s.Headers = v
	return s
}

func (s *GetEditingProjectResponse) SetBody(v *GetEditingProjectResponseBody) *GetEditingProjectResponse {
	s.Body = v
	return s
}

type GetEditingProjectMaterialsRequest struct {
	MaterialType         *string `json:"MaterialType,omitempty" xml:"MaterialType,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ProjectId            *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *string `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	Type                 *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetEditingProjectMaterialsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetEditingProjectMaterialsRequest) GoString() string {
	return s.String()
}

func (s *GetEditingProjectMaterialsRequest) SetMaterialType(v string) *GetEditingProjectMaterialsRequest {
	s.MaterialType = &v
	return s
}

func (s *GetEditingProjectMaterialsRequest) SetOwnerAccount(v string) *GetEditingProjectMaterialsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *GetEditingProjectMaterialsRequest) SetOwnerId(v string) *GetEditingProjectMaterialsRequest {
	s.OwnerId = &v
	return s
}

func (s *GetEditingProjectMaterialsRequest) SetProjectId(v string) *GetEditingProjectMaterialsRequest {
	s.ProjectId = &v
	return s
}

func (s *GetEditingProjectMaterialsRequest) SetResourceOwnerAccount(v string) *GetEditingProjectMaterialsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetEditingProjectMaterialsRequest) SetResourceOwnerId(v string) *GetEditingProjectMaterialsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetEditingProjectMaterialsRequest) SetType(v string) *GetEditingProjectMaterialsRequest {
	s.Type = &v
	return s
}

type GetEditingProjectMaterialsResponseBody struct {
	MaterialList *GetEditingProjectMaterialsResponseBodyMaterialList `json:"MaterialList,omitempty" xml:"MaterialList,omitempty" type:"Struct"`
	RequestId    *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetEditingProjectMaterialsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetEditingProjectMaterialsResponseBody) GoString() string {
	return s.String()
}

func (s *GetEditingProjectMaterialsResponseBody) SetMaterialList(v *GetEditingProjectMaterialsResponseBodyMaterialList) *GetEditingProjectMaterialsResponseBody {
	s.MaterialList = v
	return s
}

func (s *GetEditingProjectMaterialsResponseBody) SetRequestId(v string) *GetEditingProjectMaterialsResponseBody {
	s.RequestId = &v
	return s
}

type GetEditingProjectMaterialsResponseBodyMaterialList struct {
	Material []*GetEditingProjectMaterialsResponseBodyMaterialListMaterial `json:"Material,omitempty" xml:"Material,omitempty" type:"Repeated"`
}

func (s GetEditingProjectMaterialsResponseBodyMaterialList) String() string {
	return tea.Prettify(s)
}

func (s GetEditingProjectMaterialsResponseBodyMaterialList) GoString() string {
	return s.String()
}

func (s *GetEditingProjectMaterialsResponseBodyMaterialList) SetMaterial(v []*GetEditingProjectMaterialsResponseBodyMaterialListMaterial) *GetEditingProjectMaterialsResponseBodyMaterialList {
	s.Material = v
	return s
}

type GetEditingProjectMaterialsResponseBodyMaterialListMaterial struct {
	CateId       *int32                                                               `json:"CateId,omitempty" xml:"CateId,omitempty"`
	CateName     *string                                                              `json:"CateName,omitempty" xml:"CateName,omitempty"`
	CoverURL     *string                                                              `json:"CoverURL,omitempty" xml:"CoverURL,omitempty"`
	CreationTime *string                                                              `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	Description  *string                                                              `json:"Description,omitempty" xml:"Description,omitempty"`
	Duration     *float32                                                             `json:"Duration,omitempty" xml:"Duration,omitempty"`
	MaterialId   *string                                                              `json:"MaterialId,omitempty" xml:"MaterialId,omitempty"`
	MaterialType *string                                                              `json:"MaterialType,omitempty" xml:"MaterialType,omitempty"`
	ModifiedTime *string                                                              `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	Size         *int64                                                               `json:"Size,omitempty" xml:"Size,omitempty"`
	Snapshots    *GetEditingProjectMaterialsResponseBodyMaterialListMaterialSnapshots `json:"Snapshots,omitempty" xml:"Snapshots,omitempty" type:"Struct"`
	Source       *string                                                              `json:"Source,omitempty" xml:"Source,omitempty"`
	SpriteConfig *string                                                              `json:"SpriteConfig,omitempty" xml:"SpriteConfig,omitempty"`
	Sprites      *GetEditingProjectMaterialsResponseBodyMaterialListMaterialSprites   `json:"Sprites,omitempty" xml:"Sprites,omitempty" type:"Struct"`
	Status       *string                                                              `json:"Status,omitempty" xml:"Status,omitempty"`
	Tags         *string                                                              `json:"Tags,omitempty" xml:"Tags,omitempty"`
	Title        *string                                                              `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s GetEditingProjectMaterialsResponseBodyMaterialListMaterial) String() string {
	return tea.Prettify(s)
}

func (s GetEditingProjectMaterialsResponseBodyMaterialListMaterial) GoString() string {
	return s.String()
}

func (s *GetEditingProjectMaterialsResponseBodyMaterialListMaterial) SetCateId(v int32) *GetEditingProjectMaterialsResponseBodyMaterialListMaterial {
	s.CateId = &v
	return s
}

func (s *GetEditingProjectMaterialsResponseBodyMaterialListMaterial) SetCateName(v string) *GetEditingProjectMaterialsResponseBodyMaterialListMaterial {
	s.CateName = &v
	return s
}

func (s *GetEditingProjectMaterialsResponseBodyMaterialListMaterial) SetCoverURL(v string) *GetEditingProjectMaterialsResponseBodyMaterialListMaterial {
	s.CoverURL = &v
	return s
}

func (s *GetEditingProjectMaterialsResponseBodyMaterialListMaterial) SetCreationTime(v string) *GetEditingProjectMaterialsResponseBodyMaterialListMaterial {
	s.CreationTime = &v
	return s
}

func (s *GetEditingProjectMaterialsResponseBodyMaterialListMaterial) SetDescription(v string) *GetEditingProjectMaterialsResponseBodyMaterialListMaterial {
	s.Description = &v
	return s
}

func (s *GetEditingProjectMaterialsResponseBodyMaterialListMaterial) SetDuration(v float32) *GetEditingProjectMaterialsResponseBodyMaterialListMaterial {
	s.Duration = &v
	return s
}

func (s *GetEditingProjectMaterialsResponseBodyMaterialListMaterial) SetMaterialId(v string) *GetEditingProjectMaterialsResponseBodyMaterialListMaterial {
	s.MaterialId = &v
	return s
}

func (s *GetEditingProjectMaterialsResponseBodyMaterialListMaterial) SetMaterialType(v string) *GetEditingProjectMaterialsResponseBodyMaterialListMaterial {
	s.MaterialType = &v
	return s
}

func (s *GetEditingProjectMaterialsResponseBodyMaterialListMaterial) SetModifiedTime(v string) *GetEditingProjectMaterialsResponseBodyMaterialListMaterial {
	s.ModifiedTime = &v
	return s
}

func (s *GetEditingProjectMaterialsResponseBodyMaterialListMaterial) SetSize(v int64) *GetEditingProjectMaterialsResponseBodyMaterialListMaterial {
	s.Size = &v
	return s
}

func (s *GetEditingProjectMaterialsResponseBodyMaterialListMaterial) SetSnapshots(v *GetEditingProjectMaterialsResponseBodyMaterialListMaterialSnapshots) *GetEditingProjectMaterialsResponseBodyMaterialListMaterial {
	s.Snapshots = v
	return s
}

func (s *GetEditingProjectMaterialsResponseBodyMaterialListMaterial) SetSource(v string) *GetEditingProjectMaterialsResponseBodyMaterialListMaterial {
	s.Source = &v
	return s
}

func (s *GetEditingProjectMaterialsResponseBodyMaterialListMaterial) SetSpriteConfig(v string) *GetEditingProjectMaterialsResponseBodyMaterialListMaterial {
	s.SpriteConfig = &v
	return s
}

func (s *GetEditingProjectMaterialsResponseBodyMaterialListMaterial) SetSprites(v *GetEditingProjectMaterialsResponseBodyMaterialListMaterialSprites) *GetEditingProjectMaterialsResponseBodyMaterialListMaterial {
	s.Sprites = v
	return s
}

func (s *GetEditingProjectMaterialsResponseBodyMaterialListMaterial) SetStatus(v string) *GetEditingProjectMaterialsResponseBodyMaterialListMaterial {
	s.Status = &v
	return s
}

func (s *GetEditingProjectMaterialsResponseBodyMaterialListMaterial) SetTags(v string) *GetEditingProjectMaterialsResponseBodyMaterialListMaterial {
	s.Tags = &v
	return s
}

func (s *GetEditingProjectMaterialsResponseBodyMaterialListMaterial) SetTitle(v string) *GetEditingProjectMaterialsResponseBodyMaterialListMaterial {
	s.Title = &v
	return s
}

type GetEditingProjectMaterialsResponseBodyMaterialListMaterialSnapshots struct {
	Snapshot []*string `json:"Snapshot,omitempty" xml:"Snapshot,omitempty" type:"Repeated"`
}

func (s GetEditingProjectMaterialsResponseBodyMaterialListMaterialSnapshots) String() string {
	return tea.Prettify(s)
}

func (s GetEditingProjectMaterialsResponseBodyMaterialListMaterialSnapshots) GoString() string {
	return s.String()
}

func (s *GetEditingProjectMaterialsResponseBodyMaterialListMaterialSnapshots) SetSnapshot(v []*string) *GetEditingProjectMaterialsResponseBodyMaterialListMaterialSnapshots {
	s.Snapshot = v
	return s
}

type GetEditingProjectMaterialsResponseBodyMaterialListMaterialSprites struct {
	Sprite []*string `json:"Sprite,omitempty" xml:"Sprite,omitempty" type:"Repeated"`
}

func (s GetEditingProjectMaterialsResponseBodyMaterialListMaterialSprites) String() string {
	return tea.Prettify(s)
}

func (s GetEditingProjectMaterialsResponseBodyMaterialListMaterialSprites) GoString() string {
	return s.String()
}

func (s *GetEditingProjectMaterialsResponseBodyMaterialListMaterialSprites) SetSprite(v []*string) *GetEditingProjectMaterialsResponseBodyMaterialListMaterialSprites {
	s.Sprite = v
	return s
}

type GetEditingProjectMaterialsResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetEditingProjectMaterialsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetEditingProjectMaterialsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetEditingProjectMaterialsResponse) GoString() string {
	return s.String()
}

func (s *GetEditingProjectMaterialsResponse) SetHeaders(v map[string]*string) *GetEditingProjectMaterialsResponse {
	s.Headers = v
	return s
}

func (s *GetEditingProjectMaterialsResponse) SetBody(v *GetEditingProjectMaterialsResponseBody) *GetEditingProjectMaterialsResponse {
	s.Body = v
	return s
}

type GetImageInfoRequest struct {
	AuthTimeout *int64  `json:"AuthTimeout,omitempty" xml:"AuthTimeout,omitempty"`
	ImageId     *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	OutputType  *string `json:"OutputType,omitempty" xml:"OutputType,omitempty"`
}

func (s GetImageInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s GetImageInfoRequest) GoString() string {
	return s.String()
}

func (s *GetImageInfoRequest) SetAuthTimeout(v int64) *GetImageInfoRequest {
	s.AuthTimeout = &v
	return s
}

func (s *GetImageInfoRequest) SetImageId(v string) *GetImageInfoRequest {
	s.ImageId = &v
	return s
}

func (s *GetImageInfoRequest) SetOutputType(v string) *GetImageInfoRequest {
	s.OutputType = &v
	return s
}

type GetImageInfoResponseBody struct {
	ImageInfo *GetImageInfoResponseBodyImageInfo `json:"ImageInfo,omitempty" xml:"ImageInfo,omitempty" type:"Struct"`
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetImageInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetImageInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetImageInfoResponseBody) SetImageInfo(v *GetImageInfoResponseBodyImageInfo) *GetImageInfoResponseBody {
	s.ImageInfo = v
	return s
}

func (s *GetImageInfoResponseBody) SetRequestId(v string) *GetImageInfoResponseBody {
	s.RequestId = &v
	return s
}

type GetImageInfoResponseBodyImageInfo struct {
	AppId           *string                                     `json:"AppId,omitempty" xml:"AppId,omitempty"`
	CateId          *int64                                      `json:"CateId,omitempty" xml:"CateId,omitempty"`
	CateName        *string                                     `json:"CateName,omitempty" xml:"CateName,omitempty"`
	CreationTime    *string                                     `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	Description     *string                                     `json:"Description,omitempty" xml:"Description,omitempty"`
	ImageId         *string                                     `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	ImageType       *string                                     `json:"ImageType,omitempty" xml:"ImageType,omitempty"`
	Mezzanine       *GetImageInfoResponseBodyImageInfoMezzanine `json:"Mezzanine,omitempty" xml:"Mezzanine,omitempty" type:"Struct"`
	Status          *string                                     `json:"Status,omitempty" xml:"Status,omitempty"`
	StorageLocation *string                                     `json:"StorageLocation,omitempty" xml:"StorageLocation,omitempty"`
	Tags            *string                                     `json:"Tags,omitempty" xml:"Tags,omitempty"`
	Title           *string                                     `json:"Title,omitempty" xml:"Title,omitempty"`
	URL             *string                                     `json:"URL,omitempty" xml:"URL,omitempty"`
}

func (s GetImageInfoResponseBodyImageInfo) String() string {
	return tea.Prettify(s)
}

func (s GetImageInfoResponseBodyImageInfo) GoString() string {
	return s.String()
}

func (s *GetImageInfoResponseBodyImageInfo) SetAppId(v string) *GetImageInfoResponseBodyImageInfo {
	s.AppId = &v
	return s
}

func (s *GetImageInfoResponseBodyImageInfo) SetCateId(v int64) *GetImageInfoResponseBodyImageInfo {
	s.CateId = &v
	return s
}

func (s *GetImageInfoResponseBodyImageInfo) SetCateName(v string) *GetImageInfoResponseBodyImageInfo {
	s.CateName = &v
	return s
}

func (s *GetImageInfoResponseBodyImageInfo) SetCreationTime(v string) *GetImageInfoResponseBodyImageInfo {
	s.CreationTime = &v
	return s
}

func (s *GetImageInfoResponseBodyImageInfo) SetDescription(v string) *GetImageInfoResponseBodyImageInfo {
	s.Description = &v
	return s
}

func (s *GetImageInfoResponseBodyImageInfo) SetImageId(v string) *GetImageInfoResponseBodyImageInfo {
	s.ImageId = &v
	return s
}

func (s *GetImageInfoResponseBodyImageInfo) SetImageType(v string) *GetImageInfoResponseBodyImageInfo {
	s.ImageType = &v
	return s
}

func (s *GetImageInfoResponseBodyImageInfo) SetMezzanine(v *GetImageInfoResponseBodyImageInfoMezzanine) *GetImageInfoResponseBodyImageInfo {
	s.Mezzanine = v
	return s
}

func (s *GetImageInfoResponseBodyImageInfo) SetStatus(v string) *GetImageInfoResponseBodyImageInfo {
	s.Status = &v
	return s
}

func (s *GetImageInfoResponseBodyImageInfo) SetStorageLocation(v string) *GetImageInfoResponseBodyImageInfo {
	s.StorageLocation = &v
	return s
}

func (s *GetImageInfoResponseBodyImageInfo) SetTags(v string) *GetImageInfoResponseBodyImageInfo {
	s.Tags = &v
	return s
}

func (s *GetImageInfoResponseBodyImageInfo) SetTitle(v string) *GetImageInfoResponseBodyImageInfo {
	s.Title = &v
	return s
}

func (s *GetImageInfoResponseBodyImageInfo) SetURL(v string) *GetImageInfoResponseBodyImageInfo {
	s.URL = &v
	return s
}

type GetImageInfoResponseBodyImageInfoMezzanine struct {
	FileSize         *string `json:"FileSize,omitempty" xml:"FileSize,omitempty"`
	FileURL          *string `json:"FileURL,omitempty" xml:"FileURL,omitempty"`
	Height           *int32  `json:"Height,omitempty" xml:"Height,omitempty"`
	OriginalFileName *string `json:"OriginalFileName,omitempty" xml:"OriginalFileName,omitempty"`
	Width            *int32  `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s GetImageInfoResponseBodyImageInfoMezzanine) String() string {
	return tea.Prettify(s)
}

func (s GetImageInfoResponseBodyImageInfoMezzanine) GoString() string {
	return s.String()
}

func (s *GetImageInfoResponseBodyImageInfoMezzanine) SetFileSize(v string) *GetImageInfoResponseBodyImageInfoMezzanine {
	s.FileSize = &v
	return s
}

func (s *GetImageInfoResponseBodyImageInfoMezzanine) SetFileURL(v string) *GetImageInfoResponseBodyImageInfoMezzanine {
	s.FileURL = &v
	return s
}

func (s *GetImageInfoResponseBodyImageInfoMezzanine) SetHeight(v int32) *GetImageInfoResponseBodyImageInfoMezzanine {
	s.Height = &v
	return s
}

func (s *GetImageInfoResponseBodyImageInfoMezzanine) SetOriginalFileName(v string) *GetImageInfoResponseBodyImageInfoMezzanine {
	s.OriginalFileName = &v
	return s
}

func (s *GetImageInfoResponseBodyImageInfoMezzanine) SetWidth(v int32) *GetImageInfoResponseBodyImageInfoMezzanine {
	s.Width = &v
	return s
}

type GetImageInfoResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetImageInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetImageInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetImageInfoResponse) GoString() string {
	return s.String()
}

func (s *GetImageInfoResponse) SetHeaders(v map[string]*string) *GetImageInfoResponse {
	s.Headers = v
	return s
}

func (s *GetImageInfoResponse) SetBody(v *GetImageInfoResponseBody) *GetImageInfoResponse {
	s.Body = v
	return s
}

type GetMediaAuditAudioResultDetailRequest struct {
	MediaId              *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNo               *int32  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *string `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s GetMediaAuditAudioResultDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s GetMediaAuditAudioResultDetailRequest) GoString() string {
	return s.String()
}

func (s *GetMediaAuditAudioResultDetailRequest) SetMediaId(v string) *GetMediaAuditAudioResultDetailRequest {
	s.MediaId = &v
	return s
}

func (s *GetMediaAuditAudioResultDetailRequest) SetOwnerAccount(v string) *GetMediaAuditAudioResultDetailRequest {
	s.OwnerAccount = &v
	return s
}

func (s *GetMediaAuditAudioResultDetailRequest) SetOwnerId(v string) *GetMediaAuditAudioResultDetailRequest {
	s.OwnerId = &v
	return s
}

func (s *GetMediaAuditAudioResultDetailRequest) SetPageNo(v int32) *GetMediaAuditAudioResultDetailRequest {
	s.PageNo = &v
	return s
}

func (s *GetMediaAuditAudioResultDetailRequest) SetResourceOwnerAccount(v string) *GetMediaAuditAudioResultDetailRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetMediaAuditAudioResultDetailRequest) SetResourceOwnerId(v string) *GetMediaAuditAudioResultDetailRequest {
	s.ResourceOwnerId = &v
	return s
}

type GetMediaAuditAudioResultDetailResponseBody struct {
	MediaAuditAudioResultDetail *GetMediaAuditAudioResultDetailResponseBodyMediaAuditAudioResultDetail `json:"MediaAuditAudioResultDetail,omitempty" xml:"MediaAuditAudioResultDetail,omitempty" type:"Struct"`
	RequestId                   *string                                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetMediaAuditAudioResultDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetMediaAuditAudioResultDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetMediaAuditAudioResultDetailResponseBody) SetMediaAuditAudioResultDetail(v *GetMediaAuditAudioResultDetailResponseBodyMediaAuditAudioResultDetail) *GetMediaAuditAudioResultDetailResponseBody {
	s.MediaAuditAudioResultDetail = v
	return s
}

func (s *GetMediaAuditAudioResultDetailResponseBody) SetRequestId(v string) *GetMediaAuditAudioResultDetailResponseBody {
	s.RequestId = &v
	return s
}

type GetMediaAuditAudioResultDetailResponseBodyMediaAuditAudioResultDetail struct {
	List      []*GetMediaAuditAudioResultDetailResponseBodyMediaAuditAudioResultDetailList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	PageTotal *int32                                                                       `json:"PageTotal,omitempty" xml:"PageTotal,omitempty"`
	Total     *int32                                                                       `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetMediaAuditAudioResultDetailResponseBodyMediaAuditAudioResultDetail) String() string {
	return tea.Prettify(s)
}

func (s GetMediaAuditAudioResultDetailResponseBodyMediaAuditAudioResultDetail) GoString() string {
	return s.String()
}

func (s *GetMediaAuditAudioResultDetailResponseBodyMediaAuditAudioResultDetail) SetList(v []*GetMediaAuditAudioResultDetailResponseBodyMediaAuditAudioResultDetailList) *GetMediaAuditAudioResultDetailResponseBodyMediaAuditAudioResultDetail {
	s.List = v
	return s
}

func (s *GetMediaAuditAudioResultDetailResponseBodyMediaAuditAudioResultDetail) SetPageTotal(v int32) *GetMediaAuditAudioResultDetailResponseBodyMediaAuditAudioResultDetail {
	s.PageTotal = &v
	return s
}

func (s *GetMediaAuditAudioResultDetailResponseBodyMediaAuditAudioResultDetail) SetTotal(v int32) *GetMediaAuditAudioResultDetailResponseBodyMediaAuditAudioResultDetail {
	s.Total = &v
	return s
}

type GetMediaAuditAudioResultDetailResponseBodyMediaAuditAudioResultDetailList struct {
	EndTime   *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Label     *string `json:"Label,omitempty" xml:"Label,omitempty"`
	StartTime *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Text      *string `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s GetMediaAuditAudioResultDetailResponseBodyMediaAuditAudioResultDetailList) String() string {
	return tea.Prettify(s)
}

func (s GetMediaAuditAudioResultDetailResponseBodyMediaAuditAudioResultDetailList) GoString() string {
	return s.String()
}

func (s *GetMediaAuditAudioResultDetailResponseBodyMediaAuditAudioResultDetailList) SetEndTime(v int64) *GetMediaAuditAudioResultDetailResponseBodyMediaAuditAudioResultDetailList {
	s.EndTime = &v
	return s
}

func (s *GetMediaAuditAudioResultDetailResponseBodyMediaAuditAudioResultDetailList) SetLabel(v string) *GetMediaAuditAudioResultDetailResponseBodyMediaAuditAudioResultDetailList {
	s.Label = &v
	return s
}

func (s *GetMediaAuditAudioResultDetailResponseBodyMediaAuditAudioResultDetailList) SetStartTime(v int64) *GetMediaAuditAudioResultDetailResponseBodyMediaAuditAudioResultDetailList {
	s.StartTime = &v
	return s
}

func (s *GetMediaAuditAudioResultDetailResponseBodyMediaAuditAudioResultDetailList) SetText(v string) *GetMediaAuditAudioResultDetailResponseBodyMediaAuditAudioResultDetailList {
	s.Text = &v
	return s
}

type GetMediaAuditAudioResultDetailResponse struct {
	Headers map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetMediaAuditAudioResultDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetMediaAuditAudioResultDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s GetMediaAuditAudioResultDetailResponse) GoString() string {
	return s.String()
}

func (s *GetMediaAuditAudioResultDetailResponse) SetHeaders(v map[string]*string) *GetMediaAuditAudioResultDetailResponse {
	s.Headers = v
	return s
}

func (s *GetMediaAuditAudioResultDetailResponse) SetBody(v *GetMediaAuditAudioResultDetailResponseBody) *GetMediaAuditAudioResultDetailResponse {
	s.Body = v
	return s
}

type GetMediaAuditResultRequest struct {
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
}

func (s GetMediaAuditResultRequest) String() string {
	return tea.Prettify(s)
}

func (s GetMediaAuditResultRequest) GoString() string {
	return s.String()
}

func (s *GetMediaAuditResultRequest) SetMediaId(v string) *GetMediaAuditResultRequest {
	s.MediaId = &v
	return s
}

type GetMediaAuditResultResponseBody struct {
	MediaAuditResult *GetMediaAuditResultResponseBodyMediaAuditResult `json:"MediaAuditResult,omitempty" xml:"MediaAuditResult,omitempty" type:"Struct"`
	RequestId        *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetMediaAuditResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetMediaAuditResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetMediaAuditResultResponseBody) SetMediaAuditResult(v *GetMediaAuditResultResponseBodyMediaAuditResult) *GetMediaAuditResultResponseBody {
	s.MediaAuditResult = v
	return s
}

func (s *GetMediaAuditResultResponseBody) SetRequestId(v string) *GetMediaAuditResultResponseBody {
	s.RequestId = &v
	return s
}

type GetMediaAuditResultResponseBodyMediaAuditResult struct {
	AbnormalModules *string                                                       `json:"AbnormalModules,omitempty" xml:"AbnormalModules,omitempty"`
	AudioResult     []*GetMediaAuditResultResponseBodyMediaAuditResultAudioResult `json:"AudioResult,omitempty" xml:"AudioResult,omitempty" type:"Repeated"`
	ImageResult     []*GetMediaAuditResultResponseBodyMediaAuditResultImageResult `json:"ImageResult,omitempty" xml:"ImageResult,omitempty" type:"Repeated"`
	Label           *string                                                       `json:"Label,omitempty" xml:"Label,omitempty"`
	Suggestion      *string                                                       `json:"Suggestion,omitempty" xml:"Suggestion,omitempty"`
	TextResult      []*GetMediaAuditResultResponseBodyMediaAuditResultTextResult  `json:"TextResult,omitempty" xml:"TextResult,omitempty" type:"Repeated"`
	VideoResult     *GetMediaAuditResultResponseBodyMediaAuditResultVideoResult   `json:"VideoResult,omitempty" xml:"VideoResult,omitempty" type:"Struct"`
}

func (s GetMediaAuditResultResponseBodyMediaAuditResult) String() string {
	return tea.Prettify(s)
}

func (s GetMediaAuditResultResponseBodyMediaAuditResult) GoString() string {
	return s.String()
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResult) SetAbnormalModules(v string) *GetMediaAuditResultResponseBodyMediaAuditResult {
	s.AbnormalModules = &v
	return s
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResult) SetAudioResult(v []*GetMediaAuditResultResponseBodyMediaAuditResultAudioResult) *GetMediaAuditResultResponseBodyMediaAuditResult {
	s.AudioResult = v
	return s
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResult) SetImageResult(v []*GetMediaAuditResultResponseBodyMediaAuditResultImageResult) *GetMediaAuditResultResponseBodyMediaAuditResult {
	s.ImageResult = v
	return s
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResult) SetLabel(v string) *GetMediaAuditResultResponseBodyMediaAuditResult {
	s.Label = &v
	return s
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResult) SetSuggestion(v string) *GetMediaAuditResultResponseBodyMediaAuditResult {
	s.Suggestion = &v
	return s
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResult) SetTextResult(v []*GetMediaAuditResultResponseBodyMediaAuditResultTextResult) *GetMediaAuditResultResponseBodyMediaAuditResult {
	s.TextResult = v
	return s
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResult) SetVideoResult(v *GetMediaAuditResultResponseBodyMediaAuditResultVideoResult) *GetMediaAuditResultResponseBodyMediaAuditResult {
	s.VideoResult = v
	return s
}

type GetMediaAuditResultResponseBodyMediaAuditResultAudioResult struct {
	Label      *string `json:"Label,omitempty" xml:"Label,omitempty"`
	Scene      *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
	Score      *string `json:"Score,omitempty" xml:"Score,omitempty"`
	Suggestion *string `json:"Suggestion,omitempty" xml:"Suggestion,omitempty"`
}

func (s GetMediaAuditResultResponseBodyMediaAuditResultAudioResult) String() string {
	return tea.Prettify(s)
}

func (s GetMediaAuditResultResponseBodyMediaAuditResultAudioResult) GoString() string {
	return s.String()
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultAudioResult) SetLabel(v string) *GetMediaAuditResultResponseBodyMediaAuditResultAudioResult {
	s.Label = &v
	return s
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultAudioResult) SetScene(v string) *GetMediaAuditResultResponseBodyMediaAuditResultAudioResult {
	s.Scene = &v
	return s
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultAudioResult) SetScore(v string) *GetMediaAuditResultResponseBodyMediaAuditResultAudioResult {
	s.Score = &v
	return s
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultAudioResult) SetSuggestion(v string) *GetMediaAuditResultResponseBodyMediaAuditResultAudioResult {
	s.Suggestion = &v
	return s
}

type GetMediaAuditResultResponseBodyMediaAuditResultImageResult struct {
	Label      *string                                                             `json:"Label,omitempty" xml:"Label,omitempty"`
	Result     []*GetMediaAuditResultResponseBodyMediaAuditResultImageResultResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	Suggestion *string                                                             `json:"Suggestion,omitempty" xml:"Suggestion,omitempty"`
	Type       *string                                                             `json:"Type,omitempty" xml:"Type,omitempty"`
	Url        *string                                                             `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s GetMediaAuditResultResponseBodyMediaAuditResultImageResult) String() string {
	return tea.Prettify(s)
}

func (s GetMediaAuditResultResponseBodyMediaAuditResultImageResult) GoString() string {
	return s.String()
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultImageResult) SetLabel(v string) *GetMediaAuditResultResponseBodyMediaAuditResultImageResult {
	s.Label = &v
	return s
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultImageResult) SetResult(v []*GetMediaAuditResultResponseBodyMediaAuditResultImageResultResult) *GetMediaAuditResultResponseBodyMediaAuditResultImageResult {
	s.Result = v
	return s
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultImageResult) SetSuggestion(v string) *GetMediaAuditResultResponseBodyMediaAuditResultImageResult {
	s.Suggestion = &v
	return s
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultImageResult) SetType(v string) *GetMediaAuditResultResponseBodyMediaAuditResultImageResult {
	s.Type = &v
	return s
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultImageResult) SetUrl(v string) *GetMediaAuditResultResponseBodyMediaAuditResultImageResult {
	s.Url = &v
	return s
}

type GetMediaAuditResultResponseBodyMediaAuditResultImageResultResult struct {
	Label      *string `json:"Label,omitempty" xml:"Label,omitempty"`
	Scene      *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
	Score      *string `json:"Score,omitempty" xml:"Score,omitempty"`
	Suggestion *string `json:"Suggestion,omitempty" xml:"Suggestion,omitempty"`
}

func (s GetMediaAuditResultResponseBodyMediaAuditResultImageResultResult) String() string {
	return tea.Prettify(s)
}

func (s GetMediaAuditResultResponseBodyMediaAuditResultImageResultResult) GoString() string {
	return s.String()
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultImageResultResult) SetLabel(v string) *GetMediaAuditResultResponseBodyMediaAuditResultImageResultResult {
	s.Label = &v
	return s
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultImageResultResult) SetScene(v string) *GetMediaAuditResultResponseBodyMediaAuditResultImageResultResult {
	s.Scene = &v
	return s
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultImageResultResult) SetScore(v string) *GetMediaAuditResultResponseBodyMediaAuditResultImageResultResult {
	s.Score = &v
	return s
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultImageResultResult) SetSuggestion(v string) *GetMediaAuditResultResponseBodyMediaAuditResultImageResultResult {
	s.Suggestion = &v
	return s
}

type GetMediaAuditResultResponseBodyMediaAuditResultTextResult struct {
	Content    *string `json:"Content,omitempty" xml:"Content,omitempty"`
	Label      *string `json:"Label,omitempty" xml:"Label,omitempty"`
	Scene      *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
	Score      *string `json:"Score,omitempty" xml:"Score,omitempty"`
	Suggestion *string `json:"Suggestion,omitempty" xml:"Suggestion,omitempty"`
	Type       *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetMediaAuditResultResponseBodyMediaAuditResultTextResult) String() string {
	return tea.Prettify(s)
}

func (s GetMediaAuditResultResponseBodyMediaAuditResultTextResult) GoString() string {
	return s.String()
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultTextResult) SetContent(v string) *GetMediaAuditResultResponseBodyMediaAuditResultTextResult {
	s.Content = &v
	return s
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultTextResult) SetLabel(v string) *GetMediaAuditResultResponseBodyMediaAuditResultTextResult {
	s.Label = &v
	return s
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultTextResult) SetScene(v string) *GetMediaAuditResultResponseBodyMediaAuditResultTextResult {
	s.Scene = &v
	return s
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultTextResult) SetScore(v string) *GetMediaAuditResultResponseBodyMediaAuditResultTextResult {
	s.Score = &v
	return s
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultTextResult) SetSuggestion(v string) *GetMediaAuditResultResponseBodyMediaAuditResultTextResult {
	s.Suggestion = &v
	return s
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultTextResult) SetType(v string) *GetMediaAuditResultResponseBodyMediaAuditResultTextResult {
	s.Type = &v
	return s
}

type GetMediaAuditResultResponseBodyMediaAuditResultVideoResult struct {
	AdResult        *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultAdResult        `json:"AdResult,omitempty" xml:"AdResult,omitempty" type:"Struct"`
	Label           *string                                                                    `json:"Label,omitempty" xml:"Label,omitempty"`
	LiveResult      *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLiveResult      `json:"LiveResult,omitempty" xml:"LiveResult,omitempty" type:"Struct"`
	LogoResult      *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLogoResult      `json:"LogoResult,omitempty" xml:"LogoResult,omitempty" type:"Struct"`
	PornResult      *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultPornResult      `json:"PornResult,omitempty" xml:"PornResult,omitempty" type:"Struct"`
	Suggestion      *string                                                                    `json:"Suggestion,omitempty" xml:"Suggestion,omitempty"`
	TerrorismResult *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultTerrorismResult `json:"TerrorismResult,omitempty" xml:"TerrorismResult,omitempty" type:"Struct"`
}

func (s GetMediaAuditResultResponseBodyMediaAuditResultVideoResult) String() string {
	return tea.Prettify(s)
}

func (s GetMediaAuditResultResponseBodyMediaAuditResultVideoResult) GoString() string {
	return s.String()
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResult) SetAdResult(v *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultAdResult) *GetMediaAuditResultResponseBodyMediaAuditResultVideoResult {
	s.AdResult = v
	return s
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResult) SetLabel(v string) *GetMediaAuditResultResponseBodyMediaAuditResultVideoResult {
	s.Label = &v
	return s
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResult) SetLiveResult(v *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLiveResult) *GetMediaAuditResultResponseBodyMediaAuditResultVideoResult {
	s.LiveResult = v
	return s
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResult) SetLogoResult(v *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLogoResult) *GetMediaAuditResultResponseBodyMediaAuditResultVideoResult {
	s.LogoResult = v
	return s
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResult) SetPornResult(v *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultPornResult) *GetMediaAuditResultResponseBodyMediaAuditResultVideoResult {
	s.PornResult = v
	return s
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResult) SetSuggestion(v string) *GetMediaAuditResultResponseBodyMediaAuditResultVideoResult {
	s.Suggestion = &v
	return s
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResult) SetTerrorismResult(v *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultTerrorismResult) *GetMediaAuditResultResponseBodyMediaAuditResultVideoResult {
	s.TerrorismResult = v
	return s
}

type GetMediaAuditResultResponseBodyMediaAuditResultVideoResultAdResult struct {
	AverageScore *string                                                                          `json:"AverageScore,omitempty" xml:"AverageScore,omitempty"`
	CounterList  []*GetMediaAuditResultResponseBodyMediaAuditResultVideoResultAdResultCounterList `json:"CounterList,omitempty" xml:"CounterList,omitempty" type:"Repeated"`
	Label        *string                                                                          `json:"Label,omitempty" xml:"Label,omitempty"`
	MaxScore     *string                                                                          `json:"MaxScore,omitempty" xml:"MaxScore,omitempty"`
	Suggestion   *string                                                                          `json:"Suggestion,omitempty" xml:"Suggestion,omitempty"`
	TopList      []*GetMediaAuditResultResponseBodyMediaAuditResultVideoResultAdResultTopList     `json:"TopList,omitempty" xml:"TopList,omitempty" type:"Repeated"`
}

func (s GetMediaAuditResultResponseBodyMediaAuditResultVideoResultAdResult) String() string {
	return tea.Prettify(s)
}

func (s GetMediaAuditResultResponseBodyMediaAuditResultVideoResultAdResult) GoString() string {
	return s.String()
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultAdResult) SetAverageScore(v string) *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultAdResult {
	s.AverageScore = &v
	return s
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultAdResult) SetCounterList(v []*GetMediaAuditResultResponseBodyMediaAuditResultVideoResultAdResultCounterList) *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultAdResult {
	s.CounterList = v
	return s
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultAdResult) SetLabel(v string) *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultAdResult {
	s.Label = &v
	return s
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultAdResult) SetMaxScore(v string) *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultAdResult {
	s.MaxScore = &v
	return s
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultAdResult) SetSuggestion(v string) *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultAdResult {
	s.Suggestion = &v
	return s
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultAdResult) SetTopList(v []*GetMediaAuditResultResponseBodyMediaAuditResultVideoResultAdResultTopList) *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultAdResult {
	s.TopList = v
	return s
}

type GetMediaAuditResultResponseBodyMediaAuditResultVideoResultAdResultCounterList struct {
	Count *int32  `json:"Count,omitempty" xml:"Count,omitempty"`
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
}

func (s GetMediaAuditResultResponseBodyMediaAuditResultVideoResultAdResultCounterList) String() string {
	return tea.Prettify(s)
}

func (s GetMediaAuditResultResponseBodyMediaAuditResultVideoResultAdResultCounterList) GoString() string {
	return s.String()
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultAdResultCounterList) SetCount(v int32) *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultAdResultCounterList {
	s.Count = &v
	return s
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultAdResultCounterList) SetLabel(v string) *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultAdResultCounterList {
	s.Label = &v
	return s
}

type GetMediaAuditResultResponseBodyMediaAuditResultVideoResultAdResultTopList struct {
	Label     *string `json:"Label,omitempty" xml:"Label,omitempty"`
	Score     *string `json:"Score,omitempty" xml:"Score,omitempty"`
	Timestamp *string `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	Url       *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s GetMediaAuditResultResponseBodyMediaAuditResultVideoResultAdResultTopList) String() string {
	return tea.Prettify(s)
}

func (s GetMediaAuditResultResponseBodyMediaAuditResultVideoResultAdResultTopList) GoString() string {
	return s.String()
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultAdResultTopList) SetLabel(v string) *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultAdResultTopList {
	s.Label = &v
	return s
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultAdResultTopList) SetScore(v string) *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultAdResultTopList {
	s.Score = &v
	return s
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultAdResultTopList) SetTimestamp(v string) *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultAdResultTopList {
	s.Timestamp = &v
	return s
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultAdResultTopList) SetUrl(v string) *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultAdResultTopList {
	s.Url = &v
	return s
}

type GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLiveResult struct {
	AverageScore *string                                                                            `json:"AverageScore,omitempty" xml:"AverageScore,omitempty"`
	CounterList  []*GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLiveResultCounterList `json:"CounterList,omitempty" xml:"CounterList,omitempty" type:"Repeated"`
	Label        *string                                                                            `json:"Label,omitempty" xml:"Label,omitempty"`
	MaxScore     *string                                                                            `json:"MaxScore,omitempty" xml:"MaxScore,omitempty"`
	Suggestion   *string                                                                            `json:"Suggestion,omitempty" xml:"Suggestion,omitempty"`
	TopList      []*GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLiveResultTopList     `json:"TopList,omitempty" xml:"TopList,omitempty" type:"Repeated"`
}

func (s GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLiveResult) String() string {
	return tea.Prettify(s)
}

func (s GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLiveResult) GoString() string {
	return s.String()
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLiveResult) SetAverageScore(v string) *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLiveResult {
	s.AverageScore = &v
	return s
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLiveResult) SetCounterList(v []*GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLiveResultCounterList) *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLiveResult {
	s.CounterList = v
	return s
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLiveResult) SetLabel(v string) *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLiveResult {
	s.Label = &v
	return s
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLiveResult) SetMaxScore(v string) *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLiveResult {
	s.MaxScore = &v
	return s
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLiveResult) SetSuggestion(v string) *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLiveResult {
	s.Suggestion = &v
	return s
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLiveResult) SetTopList(v []*GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLiveResultTopList) *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLiveResult {
	s.TopList = v
	return s
}

type GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLiveResultCounterList struct {
	Count *int32  `json:"Count,omitempty" xml:"Count,omitempty"`
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
}

func (s GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLiveResultCounterList) String() string {
	return tea.Prettify(s)
}

func (s GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLiveResultCounterList) GoString() string {
	return s.String()
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLiveResultCounterList) SetCount(v int32) *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLiveResultCounterList {
	s.Count = &v
	return s
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLiveResultCounterList) SetLabel(v string) *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLiveResultCounterList {
	s.Label = &v
	return s
}

type GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLiveResultTopList struct {
	Label     *string `json:"Label,omitempty" xml:"Label,omitempty"`
	Score     *string `json:"Score,omitempty" xml:"Score,omitempty"`
	Timestamp *string `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	Url       *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLiveResultTopList) String() string {
	return tea.Prettify(s)
}

func (s GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLiveResultTopList) GoString() string {
	return s.String()
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLiveResultTopList) SetLabel(v string) *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLiveResultTopList {
	s.Label = &v
	return s
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLiveResultTopList) SetScore(v string) *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLiveResultTopList {
	s.Score = &v
	return s
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLiveResultTopList) SetTimestamp(v string) *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLiveResultTopList {
	s.Timestamp = &v
	return s
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLiveResultTopList) SetUrl(v string) *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLiveResultTopList {
	s.Url = &v
	return s
}

type GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLogoResult struct {
	AverageScore *string                                                                            `json:"AverageScore,omitempty" xml:"AverageScore,omitempty"`
	CounterList  []*GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLogoResultCounterList `json:"CounterList,omitempty" xml:"CounterList,omitempty" type:"Repeated"`
	Label        *string                                                                            `json:"Label,omitempty" xml:"Label,omitempty"`
	MaxScore     *string                                                                            `json:"MaxScore,omitempty" xml:"MaxScore,omitempty"`
	Suggestion   *string                                                                            `json:"Suggestion,omitempty" xml:"Suggestion,omitempty"`
	TopList      []*GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLogoResultTopList     `json:"TopList,omitempty" xml:"TopList,omitempty" type:"Repeated"`
}

func (s GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLogoResult) String() string {
	return tea.Prettify(s)
}

func (s GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLogoResult) GoString() string {
	return s.String()
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLogoResult) SetAverageScore(v string) *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLogoResult {
	s.AverageScore = &v
	return s
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLogoResult) SetCounterList(v []*GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLogoResultCounterList) *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLogoResult {
	s.CounterList = v
	return s
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLogoResult) SetLabel(v string) *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLogoResult {
	s.Label = &v
	return s
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLogoResult) SetMaxScore(v string) *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLogoResult {
	s.MaxScore = &v
	return s
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLogoResult) SetSuggestion(v string) *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLogoResult {
	s.Suggestion = &v
	return s
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLogoResult) SetTopList(v []*GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLogoResultTopList) *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLogoResult {
	s.TopList = v
	return s
}

type GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLogoResultCounterList struct {
	Count *int32  `json:"Count,omitempty" xml:"Count,omitempty"`
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
}

func (s GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLogoResultCounterList) String() string {
	return tea.Prettify(s)
}

func (s GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLogoResultCounterList) GoString() string {
	return s.String()
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLogoResultCounterList) SetCount(v int32) *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLogoResultCounterList {
	s.Count = &v
	return s
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLogoResultCounterList) SetLabel(v string) *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLogoResultCounterList {
	s.Label = &v
	return s
}

type GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLogoResultTopList struct {
	Label     *string `json:"Label,omitempty" xml:"Label,omitempty"`
	Score     *string `json:"Score,omitempty" xml:"Score,omitempty"`
	Timestamp *string `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	Url       *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLogoResultTopList) String() string {
	return tea.Prettify(s)
}

func (s GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLogoResultTopList) GoString() string {
	return s.String()
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLogoResultTopList) SetLabel(v string) *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLogoResultTopList {
	s.Label = &v
	return s
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLogoResultTopList) SetScore(v string) *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLogoResultTopList {
	s.Score = &v
	return s
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLogoResultTopList) SetTimestamp(v string) *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLogoResultTopList {
	s.Timestamp = &v
	return s
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLogoResultTopList) SetUrl(v string) *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultLogoResultTopList {
	s.Url = &v
	return s
}

type GetMediaAuditResultResponseBodyMediaAuditResultVideoResultPornResult struct {
	AverageScore *string                                                                            `json:"AverageScore,omitempty" xml:"AverageScore,omitempty"`
	CounterList  []*GetMediaAuditResultResponseBodyMediaAuditResultVideoResultPornResultCounterList `json:"CounterList,omitempty" xml:"CounterList,omitempty" type:"Repeated"`
	Label        *string                                                                            `json:"Label,omitempty" xml:"Label,omitempty"`
	MaxScore     *string                                                                            `json:"MaxScore,omitempty" xml:"MaxScore,omitempty"`
	Suggestion   *string                                                                            `json:"Suggestion,omitempty" xml:"Suggestion,omitempty"`
	TopList      []*GetMediaAuditResultResponseBodyMediaAuditResultVideoResultPornResultTopList     `json:"TopList,omitempty" xml:"TopList,omitempty" type:"Repeated"`
}

func (s GetMediaAuditResultResponseBodyMediaAuditResultVideoResultPornResult) String() string {
	return tea.Prettify(s)
}

func (s GetMediaAuditResultResponseBodyMediaAuditResultVideoResultPornResult) GoString() string {
	return s.String()
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultPornResult) SetAverageScore(v string) *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultPornResult {
	s.AverageScore = &v
	return s
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultPornResult) SetCounterList(v []*GetMediaAuditResultResponseBodyMediaAuditResultVideoResultPornResultCounterList) *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultPornResult {
	s.CounterList = v
	return s
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultPornResult) SetLabel(v string) *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultPornResult {
	s.Label = &v
	return s
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultPornResult) SetMaxScore(v string) *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultPornResult {
	s.MaxScore = &v
	return s
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultPornResult) SetSuggestion(v string) *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultPornResult {
	s.Suggestion = &v
	return s
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultPornResult) SetTopList(v []*GetMediaAuditResultResponseBodyMediaAuditResultVideoResultPornResultTopList) *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultPornResult {
	s.TopList = v
	return s
}

type GetMediaAuditResultResponseBodyMediaAuditResultVideoResultPornResultCounterList struct {
	Count *int32  `json:"Count,omitempty" xml:"Count,omitempty"`
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
}

func (s GetMediaAuditResultResponseBodyMediaAuditResultVideoResultPornResultCounterList) String() string {
	return tea.Prettify(s)
}

func (s GetMediaAuditResultResponseBodyMediaAuditResultVideoResultPornResultCounterList) GoString() string {
	return s.String()
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultPornResultCounterList) SetCount(v int32) *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultPornResultCounterList {
	s.Count = &v
	return s
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultPornResultCounterList) SetLabel(v string) *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultPornResultCounterList {
	s.Label = &v
	return s
}

type GetMediaAuditResultResponseBodyMediaAuditResultVideoResultPornResultTopList struct {
	Label     *string `json:"Label,omitempty" xml:"Label,omitempty"`
	Score     *string `json:"Score,omitempty" xml:"Score,omitempty"`
	Timestamp *string `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	Url       *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s GetMediaAuditResultResponseBodyMediaAuditResultVideoResultPornResultTopList) String() string {
	return tea.Prettify(s)
}

func (s GetMediaAuditResultResponseBodyMediaAuditResultVideoResultPornResultTopList) GoString() string {
	return s.String()
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultPornResultTopList) SetLabel(v string) *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultPornResultTopList {
	s.Label = &v
	return s
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultPornResultTopList) SetScore(v string) *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultPornResultTopList {
	s.Score = &v
	return s
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultPornResultTopList) SetTimestamp(v string) *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultPornResultTopList {
	s.Timestamp = &v
	return s
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultPornResultTopList) SetUrl(v string) *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultPornResultTopList {
	s.Url = &v
	return s
}

type GetMediaAuditResultResponseBodyMediaAuditResultVideoResultTerrorismResult struct {
	AverageScore *string                                                                                 `json:"AverageScore,omitempty" xml:"AverageScore,omitempty"`
	CounterList  []*GetMediaAuditResultResponseBodyMediaAuditResultVideoResultTerrorismResultCounterList `json:"CounterList,omitempty" xml:"CounterList,omitempty" type:"Repeated"`
	Label        *string                                                                                 `json:"Label,omitempty" xml:"Label,omitempty"`
	MaxScore     *string                                                                                 `json:"MaxScore,omitempty" xml:"MaxScore,omitempty"`
	Suggestion   *string                                                                                 `json:"Suggestion,omitempty" xml:"Suggestion,omitempty"`
	TopList      []*GetMediaAuditResultResponseBodyMediaAuditResultVideoResultTerrorismResultTopList     `json:"TopList,omitempty" xml:"TopList,omitempty" type:"Repeated"`
}

func (s GetMediaAuditResultResponseBodyMediaAuditResultVideoResultTerrorismResult) String() string {
	return tea.Prettify(s)
}

func (s GetMediaAuditResultResponseBodyMediaAuditResultVideoResultTerrorismResult) GoString() string {
	return s.String()
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultTerrorismResult) SetAverageScore(v string) *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultTerrorismResult {
	s.AverageScore = &v
	return s
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultTerrorismResult) SetCounterList(v []*GetMediaAuditResultResponseBodyMediaAuditResultVideoResultTerrorismResultCounterList) *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultTerrorismResult {
	s.CounterList = v
	return s
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultTerrorismResult) SetLabel(v string) *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultTerrorismResult {
	s.Label = &v
	return s
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultTerrorismResult) SetMaxScore(v string) *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultTerrorismResult {
	s.MaxScore = &v
	return s
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultTerrorismResult) SetSuggestion(v string) *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultTerrorismResult {
	s.Suggestion = &v
	return s
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultTerrorismResult) SetTopList(v []*GetMediaAuditResultResponseBodyMediaAuditResultVideoResultTerrorismResultTopList) *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultTerrorismResult {
	s.TopList = v
	return s
}

type GetMediaAuditResultResponseBodyMediaAuditResultVideoResultTerrorismResultCounterList struct {
	Count *int32  `json:"Count,omitempty" xml:"Count,omitempty"`
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
}

func (s GetMediaAuditResultResponseBodyMediaAuditResultVideoResultTerrorismResultCounterList) String() string {
	return tea.Prettify(s)
}

func (s GetMediaAuditResultResponseBodyMediaAuditResultVideoResultTerrorismResultCounterList) GoString() string {
	return s.String()
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultTerrorismResultCounterList) SetCount(v int32) *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultTerrorismResultCounterList {
	s.Count = &v
	return s
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultTerrorismResultCounterList) SetLabel(v string) *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultTerrorismResultCounterList {
	s.Label = &v
	return s
}

type GetMediaAuditResultResponseBodyMediaAuditResultVideoResultTerrorismResultTopList struct {
	Label     *string `json:"Label,omitempty" xml:"Label,omitempty"`
	Score     *string `json:"Score,omitempty" xml:"Score,omitempty"`
	Timestamp *string `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	Url       *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s GetMediaAuditResultResponseBodyMediaAuditResultVideoResultTerrorismResultTopList) String() string {
	return tea.Prettify(s)
}

func (s GetMediaAuditResultResponseBodyMediaAuditResultVideoResultTerrorismResultTopList) GoString() string {
	return s.String()
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultTerrorismResultTopList) SetLabel(v string) *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultTerrorismResultTopList {
	s.Label = &v
	return s
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultTerrorismResultTopList) SetScore(v string) *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultTerrorismResultTopList {
	s.Score = &v
	return s
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultTerrorismResultTopList) SetTimestamp(v string) *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultTerrorismResultTopList {
	s.Timestamp = &v
	return s
}

func (s *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultTerrorismResultTopList) SetUrl(v string) *GetMediaAuditResultResponseBodyMediaAuditResultVideoResultTerrorismResultTopList {
	s.Url = &v
	return s
}

type GetMediaAuditResultResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetMediaAuditResultResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetMediaAuditResultResponse) String() string {
	return tea.Prettify(s)
}

func (s GetMediaAuditResultResponse) GoString() string {
	return s.String()
}

func (s *GetMediaAuditResultResponse) SetHeaders(v map[string]*string) *GetMediaAuditResultResponse {
	s.Headers = v
	return s
}

func (s *GetMediaAuditResultResponse) SetBody(v *GetMediaAuditResultResponseBody) *GetMediaAuditResultResponse {
	s.Body = v
	return s
}

type GetMediaAuditResultDetailRequest struct {
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	PageNo  *int32  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
}

func (s GetMediaAuditResultDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s GetMediaAuditResultDetailRequest) GoString() string {
	return s.String()
}

func (s *GetMediaAuditResultDetailRequest) SetMediaId(v string) *GetMediaAuditResultDetailRequest {
	s.MediaId = &v
	return s
}

func (s *GetMediaAuditResultDetailRequest) SetPageNo(v int32) *GetMediaAuditResultDetailRequest {
	s.PageNo = &v
	return s
}

type GetMediaAuditResultDetailResponseBody struct {
	MediaAuditResultDetail *GetMediaAuditResultDetailResponseBodyMediaAuditResultDetail `json:"MediaAuditResultDetail,omitempty" xml:"MediaAuditResultDetail,omitempty" type:"Struct"`
	RequestId              *string                                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetMediaAuditResultDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetMediaAuditResultDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetMediaAuditResultDetailResponseBody) SetMediaAuditResultDetail(v *GetMediaAuditResultDetailResponseBodyMediaAuditResultDetail) *GetMediaAuditResultDetailResponseBody {
	s.MediaAuditResultDetail = v
	return s
}

func (s *GetMediaAuditResultDetailResponseBody) SetRequestId(v string) *GetMediaAuditResultDetailResponseBody {
	s.RequestId = &v
	return s
}

type GetMediaAuditResultDetailResponseBodyMediaAuditResultDetail struct {
	List  []*GetMediaAuditResultDetailResponseBodyMediaAuditResultDetailList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	Total *int32                                                             `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetMediaAuditResultDetailResponseBodyMediaAuditResultDetail) String() string {
	return tea.Prettify(s)
}

func (s GetMediaAuditResultDetailResponseBodyMediaAuditResultDetail) GoString() string {
	return s.String()
}

func (s *GetMediaAuditResultDetailResponseBodyMediaAuditResultDetail) SetList(v []*GetMediaAuditResultDetailResponseBodyMediaAuditResultDetailList) *GetMediaAuditResultDetailResponseBodyMediaAuditResultDetail {
	s.List = v
	return s
}

func (s *GetMediaAuditResultDetailResponseBodyMediaAuditResultDetail) SetTotal(v int32) *GetMediaAuditResultDetailResponseBodyMediaAuditResultDetail {
	s.Total = &v
	return s
}

type GetMediaAuditResultDetailResponseBodyMediaAuditResultDetailList struct {
	AdLabel        *string `json:"AdLabel,omitempty" xml:"AdLabel,omitempty"`
	AdScore        *string `json:"AdScore,omitempty" xml:"AdScore,omitempty"`
	LiveLabel      *string `json:"LiveLabel,omitempty" xml:"LiveLabel,omitempty"`
	LiveScore      *string `json:"LiveScore,omitempty" xml:"LiveScore,omitempty"`
	LogoLabel      *string `json:"LogoLabel,omitempty" xml:"LogoLabel,omitempty"`
	LogoScore      *string `json:"LogoScore,omitempty" xml:"LogoScore,omitempty"`
	PornLabel      *string `json:"PornLabel,omitempty" xml:"PornLabel,omitempty"`
	PornScore      *string `json:"PornScore,omitempty" xml:"PornScore,omitempty"`
	TerrorismLabel *string `json:"TerrorismLabel,omitempty" xml:"TerrorismLabel,omitempty"`
	TerrorismScore *string `json:"TerrorismScore,omitempty" xml:"TerrorismScore,omitempty"`
	Timestamp      *string `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	Url            *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s GetMediaAuditResultDetailResponseBodyMediaAuditResultDetailList) String() string {
	return tea.Prettify(s)
}

func (s GetMediaAuditResultDetailResponseBodyMediaAuditResultDetailList) GoString() string {
	return s.String()
}

func (s *GetMediaAuditResultDetailResponseBodyMediaAuditResultDetailList) SetAdLabel(v string) *GetMediaAuditResultDetailResponseBodyMediaAuditResultDetailList {
	s.AdLabel = &v
	return s
}

func (s *GetMediaAuditResultDetailResponseBodyMediaAuditResultDetailList) SetAdScore(v string) *GetMediaAuditResultDetailResponseBodyMediaAuditResultDetailList {
	s.AdScore = &v
	return s
}

func (s *GetMediaAuditResultDetailResponseBodyMediaAuditResultDetailList) SetLiveLabel(v string) *GetMediaAuditResultDetailResponseBodyMediaAuditResultDetailList {
	s.LiveLabel = &v
	return s
}

func (s *GetMediaAuditResultDetailResponseBodyMediaAuditResultDetailList) SetLiveScore(v string) *GetMediaAuditResultDetailResponseBodyMediaAuditResultDetailList {
	s.LiveScore = &v
	return s
}

func (s *GetMediaAuditResultDetailResponseBodyMediaAuditResultDetailList) SetLogoLabel(v string) *GetMediaAuditResultDetailResponseBodyMediaAuditResultDetailList {
	s.LogoLabel = &v
	return s
}

func (s *GetMediaAuditResultDetailResponseBodyMediaAuditResultDetailList) SetLogoScore(v string) *GetMediaAuditResultDetailResponseBodyMediaAuditResultDetailList {
	s.LogoScore = &v
	return s
}

func (s *GetMediaAuditResultDetailResponseBodyMediaAuditResultDetailList) SetPornLabel(v string) *GetMediaAuditResultDetailResponseBodyMediaAuditResultDetailList {
	s.PornLabel = &v
	return s
}

func (s *GetMediaAuditResultDetailResponseBodyMediaAuditResultDetailList) SetPornScore(v string) *GetMediaAuditResultDetailResponseBodyMediaAuditResultDetailList {
	s.PornScore = &v
	return s
}

func (s *GetMediaAuditResultDetailResponseBodyMediaAuditResultDetailList) SetTerrorismLabel(v string) *GetMediaAuditResultDetailResponseBodyMediaAuditResultDetailList {
	s.TerrorismLabel = &v
	return s
}

func (s *GetMediaAuditResultDetailResponseBodyMediaAuditResultDetailList) SetTerrorismScore(v string) *GetMediaAuditResultDetailResponseBodyMediaAuditResultDetailList {
	s.TerrorismScore = &v
	return s
}

func (s *GetMediaAuditResultDetailResponseBodyMediaAuditResultDetailList) SetTimestamp(v string) *GetMediaAuditResultDetailResponseBodyMediaAuditResultDetailList {
	s.Timestamp = &v
	return s
}

func (s *GetMediaAuditResultDetailResponseBodyMediaAuditResultDetailList) SetUrl(v string) *GetMediaAuditResultDetailResponseBodyMediaAuditResultDetailList {
	s.Url = &v
	return s
}

type GetMediaAuditResultDetailResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetMediaAuditResultDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetMediaAuditResultDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s GetMediaAuditResultDetailResponse) GoString() string {
	return s.String()
}

func (s *GetMediaAuditResultDetailResponse) SetHeaders(v map[string]*string) *GetMediaAuditResultDetailResponse {
	s.Headers = v
	return s
}

func (s *GetMediaAuditResultDetailResponse) SetBody(v *GetMediaAuditResultDetailResponseBody) *GetMediaAuditResultDetailResponse {
	s.Body = v
	return s
}

type GetMediaAuditResultTimelineRequest struct {
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
}

func (s GetMediaAuditResultTimelineRequest) String() string {
	return tea.Prettify(s)
}

func (s GetMediaAuditResultTimelineRequest) GoString() string {
	return s.String()
}

func (s *GetMediaAuditResultTimelineRequest) SetMediaId(v string) *GetMediaAuditResultTimelineRequest {
	s.MediaId = &v
	return s
}

type GetMediaAuditResultTimelineResponseBody struct {
	MediaAuditResultTimeline *GetMediaAuditResultTimelineResponseBodyMediaAuditResultTimeline `json:"MediaAuditResultTimeline,omitempty" xml:"MediaAuditResultTimeline,omitempty" type:"Struct"`
	RequestId                *string                                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetMediaAuditResultTimelineResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetMediaAuditResultTimelineResponseBody) GoString() string {
	return s.String()
}

func (s *GetMediaAuditResultTimelineResponseBody) SetMediaAuditResultTimeline(v *GetMediaAuditResultTimelineResponseBodyMediaAuditResultTimeline) *GetMediaAuditResultTimelineResponseBody {
	s.MediaAuditResultTimeline = v
	return s
}

func (s *GetMediaAuditResultTimelineResponseBody) SetRequestId(v string) *GetMediaAuditResultTimelineResponseBody {
	s.RequestId = &v
	return s
}

type GetMediaAuditResultTimelineResponseBodyMediaAuditResultTimeline struct {
	Ad        []*GetMediaAuditResultTimelineResponseBodyMediaAuditResultTimelineAd        `json:"Ad,omitempty" xml:"Ad,omitempty" type:"Repeated"`
	Live      []*GetMediaAuditResultTimelineResponseBodyMediaAuditResultTimelineLive      `json:"Live,omitempty" xml:"Live,omitempty" type:"Repeated"`
	Logo      []*GetMediaAuditResultTimelineResponseBodyMediaAuditResultTimelineLogo      `json:"Logo,omitempty" xml:"Logo,omitempty" type:"Repeated"`
	Porn      []*GetMediaAuditResultTimelineResponseBodyMediaAuditResultTimelinePorn      `json:"Porn,omitempty" xml:"Porn,omitempty" type:"Repeated"`
	Terrorism []*GetMediaAuditResultTimelineResponseBodyMediaAuditResultTimelineTerrorism `json:"Terrorism,omitempty" xml:"Terrorism,omitempty" type:"Repeated"`
}

func (s GetMediaAuditResultTimelineResponseBodyMediaAuditResultTimeline) String() string {
	return tea.Prettify(s)
}

func (s GetMediaAuditResultTimelineResponseBodyMediaAuditResultTimeline) GoString() string {
	return s.String()
}

func (s *GetMediaAuditResultTimelineResponseBodyMediaAuditResultTimeline) SetAd(v []*GetMediaAuditResultTimelineResponseBodyMediaAuditResultTimelineAd) *GetMediaAuditResultTimelineResponseBodyMediaAuditResultTimeline {
	s.Ad = v
	return s
}

func (s *GetMediaAuditResultTimelineResponseBodyMediaAuditResultTimeline) SetLive(v []*GetMediaAuditResultTimelineResponseBodyMediaAuditResultTimelineLive) *GetMediaAuditResultTimelineResponseBodyMediaAuditResultTimeline {
	s.Live = v
	return s
}

func (s *GetMediaAuditResultTimelineResponseBodyMediaAuditResultTimeline) SetLogo(v []*GetMediaAuditResultTimelineResponseBodyMediaAuditResultTimelineLogo) *GetMediaAuditResultTimelineResponseBodyMediaAuditResultTimeline {
	s.Logo = v
	return s
}

func (s *GetMediaAuditResultTimelineResponseBodyMediaAuditResultTimeline) SetPorn(v []*GetMediaAuditResultTimelineResponseBodyMediaAuditResultTimelinePorn) *GetMediaAuditResultTimelineResponseBodyMediaAuditResultTimeline {
	s.Porn = v
	return s
}

func (s *GetMediaAuditResultTimelineResponseBodyMediaAuditResultTimeline) SetTerrorism(v []*GetMediaAuditResultTimelineResponseBodyMediaAuditResultTimelineTerrorism) *GetMediaAuditResultTimelineResponseBodyMediaAuditResultTimeline {
	s.Terrorism = v
	return s
}

type GetMediaAuditResultTimelineResponseBodyMediaAuditResultTimelineAd struct {
	Label     *string `json:"Label,omitempty" xml:"Label,omitempty"`
	Score     *string `json:"Score,omitempty" xml:"Score,omitempty"`
	Timestamp *string `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s GetMediaAuditResultTimelineResponseBodyMediaAuditResultTimelineAd) String() string {
	return tea.Prettify(s)
}

func (s GetMediaAuditResultTimelineResponseBodyMediaAuditResultTimelineAd) GoString() string {
	return s.String()
}

func (s *GetMediaAuditResultTimelineResponseBodyMediaAuditResultTimelineAd) SetLabel(v string) *GetMediaAuditResultTimelineResponseBodyMediaAuditResultTimelineAd {
	s.Label = &v
	return s
}

func (s *GetMediaAuditResultTimelineResponseBodyMediaAuditResultTimelineAd) SetScore(v string) *GetMediaAuditResultTimelineResponseBodyMediaAuditResultTimelineAd {
	s.Score = &v
	return s
}

func (s *GetMediaAuditResultTimelineResponseBodyMediaAuditResultTimelineAd) SetTimestamp(v string) *GetMediaAuditResultTimelineResponseBodyMediaAuditResultTimelineAd {
	s.Timestamp = &v
	return s
}

type GetMediaAuditResultTimelineResponseBodyMediaAuditResultTimelineLive struct {
	Label     *string `json:"Label,omitempty" xml:"Label,omitempty"`
	Score     *string `json:"Score,omitempty" xml:"Score,omitempty"`
	Timestamp *string `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s GetMediaAuditResultTimelineResponseBodyMediaAuditResultTimelineLive) String() string {
	return tea.Prettify(s)
}

func (s GetMediaAuditResultTimelineResponseBodyMediaAuditResultTimelineLive) GoString() string {
	return s.String()
}

func (s *GetMediaAuditResultTimelineResponseBodyMediaAuditResultTimelineLive) SetLabel(v string) *GetMediaAuditResultTimelineResponseBodyMediaAuditResultTimelineLive {
	s.Label = &v
	return s
}

func (s *GetMediaAuditResultTimelineResponseBodyMediaAuditResultTimelineLive) SetScore(v string) *GetMediaAuditResultTimelineResponseBodyMediaAuditResultTimelineLive {
	s.Score = &v
	return s
}

func (s *GetMediaAuditResultTimelineResponseBodyMediaAuditResultTimelineLive) SetTimestamp(v string) *GetMediaAuditResultTimelineResponseBodyMediaAuditResultTimelineLive {
	s.Timestamp = &v
	return s
}

type GetMediaAuditResultTimelineResponseBodyMediaAuditResultTimelineLogo struct {
	Label     *string `json:"Label,omitempty" xml:"Label,omitempty"`
	Score     *string `json:"Score,omitempty" xml:"Score,omitempty"`
	Timestamp *string `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s GetMediaAuditResultTimelineResponseBodyMediaAuditResultTimelineLogo) String() string {
	return tea.Prettify(s)
}

func (s GetMediaAuditResultTimelineResponseBodyMediaAuditResultTimelineLogo) GoString() string {
	return s.String()
}

func (s *GetMediaAuditResultTimelineResponseBodyMediaAuditResultTimelineLogo) SetLabel(v string) *GetMediaAuditResultTimelineResponseBodyMediaAuditResultTimelineLogo {
	s.Label = &v
	return s
}

func (s *GetMediaAuditResultTimelineResponseBodyMediaAuditResultTimelineLogo) SetScore(v string) *GetMediaAuditResultTimelineResponseBodyMediaAuditResultTimelineLogo {
	s.Score = &v
	return s
}

func (s *GetMediaAuditResultTimelineResponseBodyMediaAuditResultTimelineLogo) SetTimestamp(v string) *GetMediaAuditResultTimelineResponseBodyMediaAuditResultTimelineLogo {
	s.Timestamp = &v
	return s
}

type GetMediaAuditResultTimelineResponseBodyMediaAuditResultTimelinePorn struct {
	Label     *string `json:"Label,omitempty" xml:"Label,omitempty"`
	Score     *string `json:"Score,omitempty" xml:"Score,omitempty"`
	Timestamp *string `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s GetMediaAuditResultTimelineResponseBodyMediaAuditResultTimelinePorn) String() string {
	return tea.Prettify(s)
}

func (s GetMediaAuditResultTimelineResponseBodyMediaAuditResultTimelinePorn) GoString() string {
	return s.String()
}

func (s *GetMediaAuditResultTimelineResponseBodyMediaAuditResultTimelinePorn) SetLabel(v string) *GetMediaAuditResultTimelineResponseBodyMediaAuditResultTimelinePorn {
	s.Label = &v
	return s
}

func (s *GetMediaAuditResultTimelineResponseBodyMediaAuditResultTimelinePorn) SetScore(v string) *GetMediaAuditResultTimelineResponseBodyMediaAuditResultTimelinePorn {
	s.Score = &v
	return s
}

func (s *GetMediaAuditResultTimelineResponseBodyMediaAuditResultTimelinePorn) SetTimestamp(v string) *GetMediaAuditResultTimelineResponseBodyMediaAuditResultTimelinePorn {
	s.Timestamp = &v
	return s
}

type GetMediaAuditResultTimelineResponseBodyMediaAuditResultTimelineTerrorism struct {
	Label     *string `json:"Label,omitempty" xml:"Label,omitempty"`
	Score     *string `json:"Score,omitempty" xml:"Score,omitempty"`
	Timestamp *string `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s GetMediaAuditResultTimelineResponseBodyMediaAuditResultTimelineTerrorism) String() string {
	return tea.Prettify(s)
}

func (s GetMediaAuditResultTimelineResponseBodyMediaAuditResultTimelineTerrorism) GoString() string {
	return s.String()
}

func (s *GetMediaAuditResultTimelineResponseBodyMediaAuditResultTimelineTerrorism) SetLabel(v string) *GetMediaAuditResultTimelineResponseBodyMediaAuditResultTimelineTerrorism {
	s.Label = &v
	return s
}

func (s *GetMediaAuditResultTimelineResponseBodyMediaAuditResultTimelineTerrorism) SetScore(v string) *GetMediaAuditResultTimelineResponseBodyMediaAuditResultTimelineTerrorism {
	s.Score = &v
	return s
}

func (s *GetMediaAuditResultTimelineResponseBodyMediaAuditResultTimelineTerrorism) SetTimestamp(v string) *GetMediaAuditResultTimelineResponseBodyMediaAuditResultTimelineTerrorism {
	s.Timestamp = &v
	return s
}

type GetMediaAuditResultTimelineResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetMediaAuditResultTimelineResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetMediaAuditResultTimelineResponse) String() string {
	return tea.Prettify(s)
}

func (s GetMediaAuditResultTimelineResponse) GoString() string {
	return s.String()
}

func (s *GetMediaAuditResultTimelineResponse) SetHeaders(v map[string]*string) *GetMediaAuditResultTimelineResponse {
	s.Headers = v
	return s
}

func (s *GetMediaAuditResultTimelineResponse) SetBody(v *GetMediaAuditResultTimelineResponseBody) *GetMediaAuditResultTimelineResponse {
	s.Body = v
	return s
}

type GetMediaDNAResultRequest struct {
	MediaId              *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *string `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s GetMediaDNAResultRequest) String() string {
	return tea.Prettify(s)
}

func (s GetMediaDNAResultRequest) GoString() string {
	return s.String()
}

func (s *GetMediaDNAResultRequest) SetMediaId(v string) *GetMediaDNAResultRequest {
	s.MediaId = &v
	return s
}

func (s *GetMediaDNAResultRequest) SetOwnerAccount(v string) *GetMediaDNAResultRequest {
	s.OwnerAccount = &v
	return s
}

func (s *GetMediaDNAResultRequest) SetOwnerId(v string) *GetMediaDNAResultRequest {
	s.OwnerId = &v
	return s
}

func (s *GetMediaDNAResultRequest) SetResourceOwnerAccount(v string) *GetMediaDNAResultRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetMediaDNAResultRequest) SetResourceOwnerId(v string) *GetMediaDNAResultRequest {
	s.ResourceOwnerId = &v
	return s
}

type GetMediaDNAResultResponseBody struct {
	DNAResult *GetMediaDNAResultResponseBodyDNAResult `json:"DNAResult,omitempty" xml:"DNAResult,omitempty" type:"Struct"`
	RequestId *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetMediaDNAResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetMediaDNAResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetMediaDNAResultResponseBody) SetDNAResult(v *GetMediaDNAResultResponseBodyDNAResult) *GetMediaDNAResultResponseBody {
	s.DNAResult = v
	return s
}

func (s *GetMediaDNAResultResponseBody) SetRequestId(v string) *GetMediaDNAResultResponseBody {
	s.RequestId = &v
	return s
}

type GetMediaDNAResultResponseBodyDNAResult struct {
	VideoDNA []*GetMediaDNAResultResponseBodyDNAResultVideoDNA `json:"VideoDNA,omitempty" xml:"VideoDNA,omitempty" type:"Repeated"`
}

func (s GetMediaDNAResultResponseBodyDNAResult) String() string {
	return tea.Prettify(s)
}

func (s GetMediaDNAResultResponseBodyDNAResult) GoString() string {
	return s.String()
}

func (s *GetMediaDNAResultResponseBodyDNAResult) SetVideoDNA(v []*GetMediaDNAResultResponseBodyDNAResultVideoDNA) *GetMediaDNAResultResponseBodyDNAResult {
	s.VideoDNA = v
	return s
}

type GetMediaDNAResultResponseBodyDNAResultVideoDNA struct {
	Detail     []*GetMediaDNAResultResponseBodyDNAResultVideoDNADetail `json:"Detail,omitempty" xml:"Detail,omitempty" type:"Repeated"`
	PrimaryKey *string                                                 `json:"PrimaryKey,omitempty" xml:"PrimaryKey,omitempty"`
	Similarity *string                                                 `json:"Similarity,omitempty" xml:"Similarity,omitempty"`
}

func (s GetMediaDNAResultResponseBodyDNAResultVideoDNA) String() string {
	return tea.Prettify(s)
}

func (s GetMediaDNAResultResponseBodyDNAResultVideoDNA) GoString() string {
	return s.String()
}

func (s *GetMediaDNAResultResponseBodyDNAResultVideoDNA) SetDetail(v []*GetMediaDNAResultResponseBodyDNAResultVideoDNADetail) *GetMediaDNAResultResponseBodyDNAResultVideoDNA {
	s.Detail = v
	return s
}

func (s *GetMediaDNAResultResponseBodyDNAResultVideoDNA) SetPrimaryKey(v string) *GetMediaDNAResultResponseBodyDNAResultVideoDNA {
	s.PrimaryKey = &v
	return s
}

func (s *GetMediaDNAResultResponseBodyDNAResultVideoDNA) SetSimilarity(v string) *GetMediaDNAResultResponseBodyDNAResultVideoDNA {
	s.Similarity = &v
	return s
}

type GetMediaDNAResultResponseBodyDNAResultVideoDNADetail struct {
	Duplication *GetMediaDNAResultResponseBodyDNAResultVideoDNADetailDuplication `json:"Duplication,omitempty" xml:"Duplication,omitempty" type:"Struct"`
	Input       *GetMediaDNAResultResponseBodyDNAResultVideoDNADetailInput       `json:"Input,omitempty" xml:"Input,omitempty" type:"Struct"`
}

func (s GetMediaDNAResultResponseBodyDNAResultVideoDNADetail) String() string {
	return tea.Prettify(s)
}

func (s GetMediaDNAResultResponseBodyDNAResultVideoDNADetail) GoString() string {
	return s.String()
}

func (s *GetMediaDNAResultResponseBodyDNAResultVideoDNADetail) SetDuplication(v *GetMediaDNAResultResponseBodyDNAResultVideoDNADetailDuplication) *GetMediaDNAResultResponseBodyDNAResultVideoDNADetail {
	s.Duplication = v
	return s
}

func (s *GetMediaDNAResultResponseBodyDNAResultVideoDNADetail) SetInput(v *GetMediaDNAResultResponseBodyDNAResultVideoDNADetailInput) *GetMediaDNAResultResponseBodyDNAResultVideoDNADetail {
	s.Input = v
	return s
}

type GetMediaDNAResultResponseBodyDNAResultVideoDNADetailDuplication struct {
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	Start    *string `json:"Start,omitempty" xml:"Start,omitempty"`
}

func (s GetMediaDNAResultResponseBodyDNAResultVideoDNADetailDuplication) String() string {
	return tea.Prettify(s)
}

func (s GetMediaDNAResultResponseBodyDNAResultVideoDNADetailDuplication) GoString() string {
	return s.String()
}

func (s *GetMediaDNAResultResponseBodyDNAResultVideoDNADetailDuplication) SetDuration(v string) *GetMediaDNAResultResponseBodyDNAResultVideoDNADetailDuplication {
	s.Duration = &v
	return s
}

func (s *GetMediaDNAResultResponseBodyDNAResultVideoDNADetailDuplication) SetStart(v string) *GetMediaDNAResultResponseBodyDNAResultVideoDNADetailDuplication {
	s.Start = &v
	return s
}

type GetMediaDNAResultResponseBodyDNAResultVideoDNADetailInput struct {
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	Start    *string `json:"Start,omitempty" xml:"Start,omitempty"`
}

func (s GetMediaDNAResultResponseBodyDNAResultVideoDNADetailInput) String() string {
	return tea.Prettify(s)
}

func (s GetMediaDNAResultResponseBodyDNAResultVideoDNADetailInput) GoString() string {
	return s.String()
}

func (s *GetMediaDNAResultResponseBodyDNAResultVideoDNADetailInput) SetDuration(v string) *GetMediaDNAResultResponseBodyDNAResultVideoDNADetailInput {
	s.Duration = &v
	return s
}

func (s *GetMediaDNAResultResponseBodyDNAResultVideoDNADetailInput) SetStart(v string) *GetMediaDNAResultResponseBodyDNAResultVideoDNADetailInput {
	s.Start = &v
	return s
}

type GetMediaDNAResultResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetMediaDNAResultResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetMediaDNAResultResponse) String() string {
	return tea.Prettify(s)
}

func (s GetMediaDNAResultResponse) GoString() string {
	return s.String()
}

func (s *GetMediaDNAResultResponse) SetHeaders(v map[string]*string) *GetMediaDNAResultResponse {
	s.Headers = v
	return s
}

func (s *GetMediaDNAResultResponse) SetBody(v *GetMediaDNAResultResponseBody) *GetMediaDNAResultResponse {
	s.Body = v
	return s
}

type GetMessageCallbackRequest struct {
	AppId        *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
}

func (s GetMessageCallbackRequest) String() string {
	return tea.Prettify(s)
}

func (s GetMessageCallbackRequest) GoString() string {
	return s.String()
}

func (s *GetMessageCallbackRequest) SetAppId(v string) *GetMessageCallbackRequest {
	s.AppId = &v
	return s
}

func (s *GetMessageCallbackRequest) SetOwnerAccount(v string) *GetMessageCallbackRequest {
	s.OwnerAccount = &v
	return s
}

type GetMessageCallbackResponseBody struct {
	MessageCallback *GetMessageCallbackResponseBodyMessageCallback `json:"MessageCallback,omitempty" xml:"MessageCallback,omitempty" type:"Struct"`
	RequestId       *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetMessageCallbackResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetMessageCallbackResponseBody) GoString() string {
	return s.String()
}

func (s *GetMessageCallbackResponseBody) SetMessageCallback(v *GetMessageCallbackResponseBodyMessageCallback) *GetMessageCallbackResponseBody {
	s.MessageCallback = v
	return s
}

func (s *GetMessageCallbackResponseBody) SetRequestId(v string) *GetMessageCallbackResponseBody {
	s.RequestId = &v
	return s
}

type GetMessageCallbackResponseBodyMessageCallback struct {
	AppId         *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AuthKey       *string `json:"AuthKey,omitempty" xml:"AuthKey,omitempty"`
	AuthSwitch    *string `json:"AuthSwitch,omitempty" xml:"AuthSwitch,omitempty"`
	CallbackType  *string `json:"CallbackType,omitempty" xml:"CallbackType,omitempty"`
	CallbackURL   *string `json:"CallbackURL,omitempty" xml:"CallbackURL,omitempty"`
	EventTypeList *string `json:"EventTypeList,omitempty" xml:"EventTypeList,omitempty"`
	MnsEndpoint   *string `json:"MnsEndpoint,omitempty" xml:"MnsEndpoint,omitempty"`
	MnsQueueName  *string `json:"MnsQueueName,omitempty" xml:"MnsQueueName,omitempty"`
}

func (s GetMessageCallbackResponseBodyMessageCallback) String() string {
	return tea.Prettify(s)
}

func (s GetMessageCallbackResponseBodyMessageCallback) GoString() string {
	return s.String()
}

func (s *GetMessageCallbackResponseBodyMessageCallback) SetAppId(v string) *GetMessageCallbackResponseBodyMessageCallback {
	s.AppId = &v
	return s
}

func (s *GetMessageCallbackResponseBodyMessageCallback) SetAuthKey(v string) *GetMessageCallbackResponseBodyMessageCallback {
	s.AuthKey = &v
	return s
}

func (s *GetMessageCallbackResponseBodyMessageCallback) SetAuthSwitch(v string) *GetMessageCallbackResponseBodyMessageCallback {
	s.AuthSwitch = &v
	return s
}

func (s *GetMessageCallbackResponseBodyMessageCallback) SetCallbackType(v string) *GetMessageCallbackResponseBodyMessageCallback {
	s.CallbackType = &v
	return s
}

func (s *GetMessageCallbackResponseBodyMessageCallback) SetCallbackURL(v string) *GetMessageCallbackResponseBodyMessageCallback {
	s.CallbackURL = &v
	return s
}

func (s *GetMessageCallbackResponseBodyMessageCallback) SetEventTypeList(v string) *GetMessageCallbackResponseBodyMessageCallback {
	s.EventTypeList = &v
	return s
}

func (s *GetMessageCallbackResponseBodyMessageCallback) SetMnsEndpoint(v string) *GetMessageCallbackResponseBodyMessageCallback {
	s.MnsEndpoint = &v
	return s
}

func (s *GetMessageCallbackResponseBodyMessageCallback) SetMnsQueueName(v string) *GetMessageCallbackResponseBodyMessageCallback {
	s.MnsQueueName = &v
	return s
}

type GetMessageCallbackResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetMessageCallbackResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetMessageCallbackResponse) String() string {
	return tea.Prettify(s)
}

func (s GetMessageCallbackResponse) GoString() string {
	return s.String()
}

func (s *GetMessageCallbackResponse) SetHeaders(v map[string]*string) *GetMessageCallbackResponse {
	s.Headers = v
	return s
}

func (s *GetMessageCallbackResponse) SetBody(v *GetMessageCallbackResponseBody) *GetMessageCallbackResponse {
	s.Body = v
	return s
}

type GetMezzanineInfoRequest struct {
	AdditionType *string `json:"AdditionType,omitempty" xml:"AdditionType,omitempty"`
	AuthTimeout  *int64  `json:"AuthTimeout,omitempty" xml:"AuthTimeout,omitempty"`
	OutputType   *string `json:"OutputType,omitempty" xml:"OutputType,omitempty"`
	VideoId      *string `json:"VideoId,omitempty" xml:"VideoId,omitempty"`
}

func (s GetMezzanineInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s GetMezzanineInfoRequest) GoString() string {
	return s.String()
}

func (s *GetMezzanineInfoRequest) SetAdditionType(v string) *GetMezzanineInfoRequest {
	s.AdditionType = &v
	return s
}

func (s *GetMezzanineInfoRequest) SetAuthTimeout(v int64) *GetMezzanineInfoRequest {
	s.AuthTimeout = &v
	return s
}

func (s *GetMezzanineInfoRequest) SetOutputType(v string) *GetMezzanineInfoRequest {
	s.OutputType = &v
	return s
}

func (s *GetMezzanineInfoRequest) SetVideoId(v string) *GetMezzanineInfoRequest {
	s.VideoId = &v
	return s
}

type GetMezzanineInfoResponseBody struct {
	Mezzanine *GetMezzanineInfoResponseBodyMezzanine `json:"Mezzanine,omitempty" xml:"Mezzanine,omitempty" type:"Struct"`
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetMezzanineInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetMezzanineInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetMezzanineInfoResponseBody) SetMezzanine(v *GetMezzanineInfoResponseBodyMezzanine) *GetMezzanineInfoResponseBody {
	s.Mezzanine = v
	return s
}

func (s *GetMezzanineInfoResponseBody) SetRequestId(v string) *GetMezzanineInfoResponseBody {
	s.RequestId = &v
	return s
}

type GetMezzanineInfoResponseBodyMezzanine struct {
	AudioStreamList []*GetMezzanineInfoResponseBodyMezzanineAudioStreamList `json:"AudioStreamList,omitempty" xml:"AudioStreamList,omitempty" type:"Repeated"`
	Bitrate         *string                                                 `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	CreationTime    *string                                                 `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	Duration        *string                                                 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	FileName        *string                                                 `json:"FileName,omitempty" xml:"FileName,omitempty"`
	FileURL         *string                                                 `json:"FileURL,omitempty" xml:"FileURL,omitempty"`
	Fps             *string                                                 `json:"Fps,omitempty" xml:"Fps,omitempty"`
	Height          *int64                                                  `json:"Height,omitempty" xml:"Height,omitempty"`
	OutputType      *string                                                 `json:"OutputType,omitempty" xml:"OutputType,omitempty"`
	Size            *int64                                                  `json:"Size,omitempty" xml:"Size,omitempty"`
	Status          *string                                                 `json:"Status,omitempty" xml:"Status,omitempty"`
	VideoId         *string                                                 `json:"VideoId,omitempty" xml:"VideoId,omitempty"`
	VideoStreamList []*GetMezzanineInfoResponseBodyMezzanineVideoStreamList `json:"VideoStreamList,omitempty" xml:"VideoStreamList,omitempty" type:"Repeated"`
	Width           *int64                                                  `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s GetMezzanineInfoResponseBodyMezzanine) String() string {
	return tea.Prettify(s)
}

func (s GetMezzanineInfoResponseBodyMezzanine) GoString() string {
	return s.String()
}

func (s *GetMezzanineInfoResponseBodyMezzanine) SetAudioStreamList(v []*GetMezzanineInfoResponseBodyMezzanineAudioStreamList) *GetMezzanineInfoResponseBodyMezzanine {
	s.AudioStreamList = v
	return s
}

func (s *GetMezzanineInfoResponseBodyMezzanine) SetBitrate(v string) *GetMezzanineInfoResponseBodyMezzanine {
	s.Bitrate = &v
	return s
}

func (s *GetMezzanineInfoResponseBodyMezzanine) SetCreationTime(v string) *GetMezzanineInfoResponseBodyMezzanine {
	s.CreationTime = &v
	return s
}

func (s *GetMezzanineInfoResponseBodyMezzanine) SetDuration(v string) *GetMezzanineInfoResponseBodyMezzanine {
	s.Duration = &v
	return s
}

func (s *GetMezzanineInfoResponseBodyMezzanine) SetFileName(v string) *GetMezzanineInfoResponseBodyMezzanine {
	s.FileName = &v
	return s
}

func (s *GetMezzanineInfoResponseBodyMezzanine) SetFileURL(v string) *GetMezzanineInfoResponseBodyMezzanine {
	s.FileURL = &v
	return s
}

func (s *GetMezzanineInfoResponseBodyMezzanine) SetFps(v string) *GetMezzanineInfoResponseBodyMezzanine {
	s.Fps = &v
	return s
}

func (s *GetMezzanineInfoResponseBodyMezzanine) SetHeight(v int64) *GetMezzanineInfoResponseBodyMezzanine {
	s.Height = &v
	return s
}

func (s *GetMezzanineInfoResponseBodyMezzanine) SetOutputType(v string) *GetMezzanineInfoResponseBodyMezzanine {
	s.OutputType = &v
	return s
}

func (s *GetMezzanineInfoResponseBodyMezzanine) SetSize(v int64) *GetMezzanineInfoResponseBodyMezzanine {
	s.Size = &v
	return s
}

func (s *GetMezzanineInfoResponseBodyMezzanine) SetStatus(v string) *GetMezzanineInfoResponseBodyMezzanine {
	s.Status = &v
	return s
}

func (s *GetMezzanineInfoResponseBodyMezzanine) SetVideoId(v string) *GetMezzanineInfoResponseBodyMezzanine {
	s.VideoId = &v
	return s
}

func (s *GetMezzanineInfoResponseBodyMezzanine) SetVideoStreamList(v []*GetMezzanineInfoResponseBodyMezzanineVideoStreamList) *GetMezzanineInfoResponseBodyMezzanine {
	s.VideoStreamList = v
	return s
}

func (s *GetMezzanineInfoResponseBodyMezzanine) SetWidth(v int64) *GetMezzanineInfoResponseBodyMezzanine {
	s.Width = &v
	return s
}

type GetMezzanineInfoResponseBodyMezzanineAudioStreamList struct {
	Bitrate        *string `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	ChannelLayout  *string `json:"ChannelLayout,omitempty" xml:"ChannelLayout,omitempty"`
	Channels       *string `json:"Channels,omitempty" xml:"Channels,omitempty"`
	CodecLongName  *string `json:"CodecLongName,omitempty" xml:"CodecLongName,omitempty"`
	CodecName      *string `json:"CodecName,omitempty" xml:"CodecName,omitempty"`
	CodecTag       *string `json:"CodecTag,omitempty" xml:"CodecTag,omitempty"`
	CodecTagString *string `json:"CodecTagString,omitempty" xml:"CodecTagString,omitempty"`
	CodecTimeBase  *string `json:"CodecTimeBase,omitempty" xml:"CodecTimeBase,omitempty"`
	Duration       *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	Index          *string `json:"Index,omitempty" xml:"Index,omitempty"`
	Lang           *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	NumFrames      *string `json:"NumFrames,omitempty" xml:"NumFrames,omitempty"`
	SampleFmt      *string `json:"SampleFmt,omitempty" xml:"SampleFmt,omitempty"`
	SampleRate     *string `json:"SampleRate,omitempty" xml:"SampleRate,omitempty"`
	StartTime      *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Timebase       *string `json:"Timebase,omitempty" xml:"Timebase,omitempty"`
}

func (s GetMezzanineInfoResponseBodyMezzanineAudioStreamList) String() string {
	return tea.Prettify(s)
}

func (s GetMezzanineInfoResponseBodyMezzanineAudioStreamList) GoString() string {
	return s.String()
}

func (s *GetMezzanineInfoResponseBodyMezzanineAudioStreamList) SetBitrate(v string) *GetMezzanineInfoResponseBodyMezzanineAudioStreamList {
	s.Bitrate = &v
	return s
}

func (s *GetMezzanineInfoResponseBodyMezzanineAudioStreamList) SetChannelLayout(v string) *GetMezzanineInfoResponseBodyMezzanineAudioStreamList {
	s.ChannelLayout = &v
	return s
}

func (s *GetMezzanineInfoResponseBodyMezzanineAudioStreamList) SetChannels(v string) *GetMezzanineInfoResponseBodyMezzanineAudioStreamList {
	s.Channels = &v
	return s
}

func (s *GetMezzanineInfoResponseBodyMezzanineAudioStreamList) SetCodecLongName(v string) *GetMezzanineInfoResponseBodyMezzanineAudioStreamList {
	s.CodecLongName = &v
	return s
}

func (s *GetMezzanineInfoResponseBodyMezzanineAudioStreamList) SetCodecName(v string) *GetMezzanineInfoResponseBodyMezzanineAudioStreamList {
	s.CodecName = &v
	return s
}

func (s *GetMezzanineInfoResponseBodyMezzanineAudioStreamList) SetCodecTag(v string) *GetMezzanineInfoResponseBodyMezzanineAudioStreamList {
	s.CodecTag = &v
	return s
}

func (s *GetMezzanineInfoResponseBodyMezzanineAudioStreamList) SetCodecTagString(v string) *GetMezzanineInfoResponseBodyMezzanineAudioStreamList {
	s.CodecTagString = &v
	return s
}

func (s *GetMezzanineInfoResponseBodyMezzanineAudioStreamList) SetCodecTimeBase(v string) *GetMezzanineInfoResponseBodyMezzanineAudioStreamList {
	s.CodecTimeBase = &v
	return s
}

func (s *GetMezzanineInfoResponseBodyMezzanineAudioStreamList) SetDuration(v string) *GetMezzanineInfoResponseBodyMezzanineAudioStreamList {
	s.Duration = &v
	return s
}

func (s *GetMezzanineInfoResponseBodyMezzanineAudioStreamList) SetIndex(v string) *GetMezzanineInfoResponseBodyMezzanineAudioStreamList {
	s.Index = &v
	return s
}

func (s *GetMezzanineInfoResponseBodyMezzanineAudioStreamList) SetLang(v string) *GetMezzanineInfoResponseBodyMezzanineAudioStreamList {
	s.Lang = &v
	return s
}

func (s *GetMezzanineInfoResponseBodyMezzanineAudioStreamList) SetNumFrames(v string) *GetMezzanineInfoResponseBodyMezzanineAudioStreamList {
	s.NumFrames = &v
	return s
}

func (s *GetMezzanineInfoResponseBodyMezzanineAudioStreamList) SetSampleFmt(v string) *GetMezzanineInfoResponseBodyMezzanineAudioStreamList {
	s.SampleFmt = &v
	return s
}

func (s *GetMezzanineInfoResponseBodyMezzanineAudioStreamList) SetSampleRate(v string) *GetMezzanineInfoResponseBodyMezzanineAudioStreamList {
	s.SampleRate = &v
	return s
}

func (s *GetMezzanineInfoResponseBodyMezzanineAudioStreamList) SetStartTime(v string) *GetMezzanineInfoResponseBodyMezzanineAudioStreamList {
	s.StartTime = &v
	return s
}

func (s *GetMezzanineInfoResponseBodyMezzanineAudioStreamList) SetTimebase(v string) *GetMezzanineInfoResponseBodyMezzanineAudioStreamList {
	s.Timebase = &v
	return s
}

type GetMezzanineInfoResponseBodyMezzanineVideoStreamList struct {
	AvgFPS         *string `json:"AvgFPS,omitempty" xml:"AvgFPS,omitempty"`
	Bitrate        *string `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	CodecLongName  *string `json:"CodecLongName,omitempty" xml:"CodecLongName,omitempty"`
	CodecName      *string `json:"CodecName,omitempty" xml:"CodecName,omitempty"`
	CodecTag       *string `json:"CodecTag,omitempty" xml:"CodecTag,omitempty"`
	CodecTagString *string `json:"CodecTagString,omitempty" xml:"CodecTagString,omitempty"`
	CodecTimeBase  *string `json:"CodecTimeBase,omitempty" xml:"CodecTimeBase,omitempty"`
	Dar            *string `json:"Dar,omitempty" xml:"Dar,omitempty"`
	Duration       *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	Fps            *string `json:"Fps,omitempty" xml:"Fps,omitempty"`
	HasBFrames     *string `json:"HasBFrames,omitempty" xml:"HasBFrames,omitempty"`
	Height         *string `json:"Height,omitempty" xml:"Height,omitempty"`
	Index          *string `json:"Index,omitempty" xml:"Index,omitempty"`
	Lang           *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	Level          *string `json:"Level,omitempty" xml:"Level,omitempty"`
	NumFrames      *string `json:"NumFrames,omitempty" xml:"NumFrames,omitempty"`
	PixFmt         *string `json:"PixFmt,omitempty" xml:"PixFmt,omitempty"`
	Profile        *string `json:"Profile,omitempty" xml:"Profile,omitempty"`
	Rotate         *string `json:"Rotate,omitempty" xml:"Rotate,omitempty"`
	Sar            *string `json:"Sar,omitempty" xml:"Sar,omitempty"`
	StartTime      *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Timebase       *string `json:"Timebase,omitempty" xml:"Timebase,omitempty"`
	Width          *string `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s GetMezzanineInfoResponseBodyMezzanineVideoStreamList) String() string {
	return tea.Prettify(s)
}

func (s GetMezzanineInfoResponseBodyMezzanineVideoStreamList) GoString() string {
	return s.String()
}

func (s *GetMezzanineInfoResponseBodyMezzanineVideoStreamList) SetAvgFPS(v string) *GetMezzanineInfoResponseBodyMezzanineVideoStreamList {
	s.AvgFPS = &v
	return s
}

func (s *GetMezzanineInfoResponseBodyMezzanineVideoStreamList) SetBitrate(v string) *GetMezzanineInfoResponseBodyMezzanineVideoStreamList {
	s.Bitrate = &v
	return s
}

func (s *GetMezzanineInfoResponseBodyMezzanineVideoStreamList) SetCodecLongName(v string) *GetMezzanineInfoResponseBodyMezzanineVideoStreamList {
	s.CodecLongName = &v
	return s
}

func (s *GetMezzanineInfoResponseBodyMezzanineVideoStreamList) SetCodecName(v string) *GetMezzanineInfoResponseBodyMezzanineVideoStreamList {
	s.CodecName = &v
	return s
}

func (s *GetMezzanineInfoResponseBodyMezzanineVideoStreamList) SetCodecTag(v string) *GetMezzanineInfoResponseBodyMezzanineVideoStreamList {
	s.CodecTag = &v
	return s
}

func (s *GetMezzanineInfoResponseBodyMezzanineVideoStreamList) SetCodecTagString(v string) *GetMezzanineInfoResponseBodyMezzanineVideoStreamList {
	s.CodecTagString = &v
	return s
}

func (s *GetMezzanineInfoResponseBodyMezzanineVideoStreamList) SetCodecTimeBase(v string) *GetMezzanineInfoResponseBodyMezzanineVideoStreamList {
	s.CodecTimeBase = &v
	return s
}

func (s *GetMezzanineInfoResponseBodyMezzanineVideoStreamList) SetDar(v string) *GetMezzanineInfoResponseBodyMezzanineVideoStreamList {
	s.Dar = &v
	return s
}

func (s *GetMezzanineInfoResponseBodyMezzanineVideoStreamList) SetDuration(v string) *GetMezzanineInfoResponseBodyMezzanineVideoStreamList {
	s.Duration = &v
	return s
}

func (s *GetMezzanineInfoResponseBodyMezzanineVideoStreamList) SetFps(v string) *GetMezzanineInfoResponseBodyMezzanineVideoStreamList {
	s.Fps = &v
	return s
}

func (s *GetMezzanineInfoResponseBodyMezzanineVideoStreamList) SetHasBFrames(v string) *GetMezzanineInfoResponseBodyMezzanineVideoStreamList {
	s.HasBFrames = &v
	return s
}

func (s *GetMezzanineInfoResponseBodyMezzanineVideoStreamList) SetHeight(v string) *GetMezzanineInfoResponseBodyMezzanineVideoStreamList {
	s.Height = &v
	return s
}

func (s *GetMezzanineInfoResponseBodyMezzanineVideoStreamList) SetIndex(v string) *GetMezzanineInfoResponseBodyMezzanineVideoStreamList {
	s.Index = &v
	return s
}

func (s *GetMezzanineInfoResponseBodyMezzanineVideoStreamList) SetLang(v string) *GetMezzanineInfoResponseBodyMezzanineVideoStreamList {
	s.Lang = &v
	return s
}

func (s *GetMezzanineInfoResponseBodyMezzanineVideoStreamList) SetLevel(v string) *GetMezzanineInfoResponseBodyMezzanineVideoStreamList {
	s.Level = &v
	return s
}

func (s *GetMezzanineInfoResponseBodyMezzanineVideoStreamList) SetNumFrames(v string) *GetMezzanineInfoResponseBodyMezzanineVideoStreamList {
	s.NumFrames = &v
	return s
}

func (s *GetMezzanineInfoResponseBodyMezzanineVideoStreamList) SetPixFmt(v string) *GetMezzanineInfoResponseBodyMezzanineVideoStreamList {
	s.PixFmt = &v
	return s
}

func (s *GetMezzanineInfoResponseBodyMezzanineVideoStreamList) SetProfile(v string) *GetMezzanineInfoResponseBodyMezzanineVideoStreamList {
	s.Profile = &v
	return s
}

func (s *GetMezzanineInfoResponseBodyMezzanineVideoStreamList) SetRotate(v string) *GetMezzanineInfoResponseBodyMezzanineVideoStreamList {
	s.Rotate = &v
	return s
}

func (s *GetMezzanineInfoResponseBodyMezzanineVideoStreamList) SetSar(v string) *GetMezzanineInfoResponseBodyMezzanineVideoStreamList {
	s.Sar = &v
	return s
}

func (s *GetMezzanineInfoResponseBodyMezzanineVideoStreamList) SetStartTime(v string) *GetMezzanineInfoResponseBodyMezzanineVideoStreamList {
	s.StartTime = &v
	return s
}

func (s *GetMezzanineInfoResponseBodyMezzanineVideoStreamList) SetTimebase(v string) *GetMezzanineInfoResponseBodyMezzanineVideoStreamList {
	s.Timebase = &v
	return s
}

func (s *GetMezzanineInfoResponseBodyMezzanineVideoStreamList) SetWidth(v string) *GetMezzanineInfoResponseBodyMezzanineVideoStreamList {
	s.Width = &v
	return s
}

type GetMezzanineInfoResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetMezzanineInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetMezzanineInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetMezzanineInfoResponse) GoString() string {
	return s.String()
}

func (s *GetMezzanineInfoResponse) SetHeaders(v map[string]*string) *GetMezzanineInfoResponse {
	s.Headers = v
	return s
}

func (s *GetMezzanineInfoResponse) SetBody(v *GetMezzanineInfoResponseBody) *GetMezzanineInfoResponse {
	s.Body = v
	return s
}

type GetPlayInfoRequest struct {
	AdditionType *string `json:"AdditionType,omitempty" xml:"AdditionType,omitempty"`
	AuthTimeout  *int64  `json:"AuthTimeout,omitempty" xml:"AuthTimeout,omitempty"`
	Definition   *string `json:"Definition,omitempty" xml:"Definition,omitempty"`
	Formats      *string `json:"Formats,omitempty" xml:"Formats,omitempty"`
	OutputType   *string `json:"OutputType,omitempty" xml:"OutputType,omitempty"`
	PlayConfig   *string `json:"PlayConfig,omitempty" xml:"PlayConfig,omitempty"`
	ReAuthInfo   *string `json:"ReAuthInfo,omitempty" xml:"ReAuthInfo,omitempty"`
	ResultType   *string `json:"ResultType,omitempty" xml:"ResultType,omitempty"`
	StreamType   *string `json:"StreamType,omitempty" xml:"StreamType,omitempty"`
	VideoId      *string `json:"VideoId,omitempty" xml:"VideoId,omitempty"`
}

func (s GetPlayInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s GetPlayInfoRequest) GoString() string {
	return s.String()
}

func (s *GetPlayInfoRequest) SetAdditionType(v string) *GetPlayInfoRequest {
	s.AdditionType = &v
	return s
}

func (s *GetPlayInfoRequest) SetAuthTimeout(v int64) *GetPlayInfoRequest {
	s.AuthTimeout = &v
	return s
}

func (s *GetPlayInfoRequest) SetDefinition(v string) *GetPlayInfoRequest {
	s.Definition = &v
	return s
}

func (s *GetPlayInfoRequest) SetFormats(v string) *GetPlayInfoRequest {
	s.Formats = &v
	return s
}

func (s *GetPlayInfoRequest) SetOutputType(v string) *GetPlayInfoRequest {
	s.OutputType = &v
	return s
}

func (s *GetPlayInfoRequest) SetPlayConfig(v string) *GetPlayInfoRequest {
	s.PlayConfig = &v
	return s
}

func (s *GetPlayInfoRequest) SetReAuthInfo(v string) *GetPlayInfoRequest {
	s.ReAuthInfo = &v
	return s
}

func (s *GetPlayInfoRequest) SetResultType(v string) *GetPlayInfoRequest {
	s.ResultType = &v
	return s
}

func (s *GetPlayInfoRequest) SetStreamType(v string) *GetPlayInfoRequest {
	s.StreamType = &v
	return s
}

func (s *GetPlayInfoRequest) SetVideoId(v string) *GetPlayInfoRequest {
	s.VideoId = &v
	return s
}

type GetPlayInfoResponseBody struct {
	PlayInfoList *GetPlayInfoResponseBodyPlayInfoList `json:"PlayInfoList,omitempty" xml:"PlayInfoList,omitempty" type:"Struct"`
	RequestId    *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	VideoBase    *GetPlayInfoResponseBodyVideoBase    `json:"VideoBase,omitempty" xml:"VideoBase,omitempty" type:"Struct"`
}

func (s GetPlayInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetPlayInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetPlayInfoResponseBody) SetPlayInfoList(v *GetPlayInfoResponseBodyPlayInfoList) *GetPlayInfoResponseBody {
	s.PlayInfoList = v
	return s
}

func (s *GetPlayInfoResponseBody) SetRequestId(v string) *GetPlayInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPlayInfoResponseBody) SetVideoBase(v *GetPlayInfoResponseBodyVideoBase) *GetPlayInfoResponseBody {
	s.VideoBase = v
	return s
}

type GetPlayInfoResponseBodyPlayInfoList struct {
	PlayInfo []*GetPlayInfoResponseBodyPlayInfoListPlayInfo `json:"PlayInfo,omitempty" xml:"PlayInfo,omitempty" type:"Repeated"`
}

func (s GetPlayInfoResponseBodyPlayInfoList) String() string {
	return tea.Prettify(s)
}

func (s GetPlayInfoResponseBodyPlayInfoList) GoString() string {
	return s.String()
}

func (s *GetPlayInfoResponseBodyPlayInfoList) SetPlayInfo(v []*GetPlayInfoResponseBodyPlayInfoListPlayInfo) *GetPlayInfoResponseBodyPlayInfoList {
	s.PlayInfo = v
	return s
}

type GetPlayInfoResponseBodyPlayInfoListPlayInfo struct {
	Bitrate          *string `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	CreationTime     *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	Definition       *string `json:"Definition,omitempty" xml:"Definition,omitempty"`
	Duration         *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	Encrypt          *int64  `json:"Encrypt,omitempty" xml:"Encrypt,omitempty"`
	EncryptType      *string `json:"EncryptType,omitempty" xml:"EncryptType,omitempty"`
	Format           *string `json:"Format,omitempty" xml:"Format,omitempty"`
	Fps              *string `json:"Fps,omitempty" xml:"Fps,omitempty"`
	Height           *int64  `json:"Height,omitempty" xml:"Height,omitempty"`
	JobId            *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	ModificationTime *string `json:"ModificationTime,omitempty" xml:"ModificationTime,omitempty"`
	NarrowBandType   *string `json:"NarrowBandType,omitempty" xml:"NarrowBandType,omitempty"`
	PlayURL          *string `json:"PlayURL,omitempty" xml:"PlayURL,omitempty"`
	Size             *int64  `json:"Size,omitempty" xml:"Size,omitempty"`
	Specification    *string `json:"Specification,omitempty" xml:"Specification,omitempty"`
	Status           *string `json:"Status,omitempty" xml:"Status,omitempty"`
	StreamType       *string `json:"StreamType,omitempty" xml:"StreamType,omitempty"`
	WatermarkId      *string `json:"WatermarkId,omitempty" xml:"WatermarkId,omitempty"`
	Width            *int64  `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s GetPlayInfoResponseBodyPlayInfoListPlayInfo) String() string {
	return tea.Prettify(s)
}

func (s GetPlayInfoResponseBodyPlayInfoListPlayInfo) GoString() string {
	return s.String()
}

func (s *GetPlayInfoResponseBodyPlayInfoListPlayInfo) SetBitrate(v string) *GetPlayInfoResponseBodyPlayInfoListPlayInfo {
	s.Bitrate = &v
	return s
}

func (s *GetPlayInfoResponseBodyPlayInfoListPlayInfo) SetCreationTime(v string) *GetPlayInfoResponseBodyPlayInfoListPlayInfo {
	s.CreationTime = &v
	return s
}

func (s *GetPlayInfoResponseBodyPlayInfoListPlayInfo) SetDefinition(v string) *GetPlayInfoResponseBodyPlayInfoListPlayInfo {
	s.Definition = &v
	return s
}

func (s *GetPlayInfoResponseBodyPlayInfoListPlayInfo) SetDuration(v string) *GetPlayInfoResponseBodyPlayInfoListPlayInfo {
	s.Duration = &v
	return s
}

func (s *GetPlayInfoResponseBodyPlayInfoListPlayInfo) SetEncrypt(v int64) *GetPlayInfoResponseBodyPlayInfoListPlayInfo {
	s.Encrypt = &v
	return s
}

func (s *GetPlayInfoResponseBodyPlayInfoListPlayInfo) SetEncryptType(v string) *GetPlayInfoResponseBodyPlayInfoListPlayInfo {
	s.EncryptType = &v
	return s
}

func (s *GetPlayInfoResponseBodyPlayInfoListPlayInfo) SetFormat(v string) *GetPlayInfoResponseBodyPlayInfoListPlayInfo {
	s.Format = &v
	return s
}

func (s *GetPlayInfoResponseBodyPlayInfoListPlayInfo) SetFps(v string) *GetPlayInfoResponseBodyPlayInfoListPlayInfo {
	s.Fps = &v
	return s
}

func (s *GetPlayInfoResponseBodyPlayInfoListPlayInfo) SetHeight(v int64) *GetPlayInfoResponseBodyPlayInfoListPlayInfo {
	s.Height = &v
	return s
}

func (s *GetPlayInfoResponseBodyPlayInfoListPlayInfo) SetJobId(v string) *GetPlayInfoResponseBodyPlayInfoListPlayInfo {
	s.JobId = &v
	return s
}

func (s *GetPlayInfoResponseBodyPlayInfoListPlayInfo) SetModificationTime(v string) *GetPlayInfoResponseBodyPlayInfoListPlayInfo {
	s.ModificationTime = &v
	return s
}

func (s *GetPlayInfoResponseBodyPlayInfoListPlayInfo) SetNarrowBandType(v string) *GetPlayInfoResponseBodyPlayInfoListPlayInfo {
	s.NarrowBandType = &v
	return s
}

func (s *GetPlayInfoResponseBodyPlayInfoListPlayInfo) SetPlayURL(v string) *GetPlayInfoResponseBodyPlayInfoListPlayInfo {
	s.PlayURL = &v
	return s
}

func (s *GetPlayInfoResponseBodyPlayInfoListPlayInfo) SetSize(v int64) *GetPlayInfoResponseBodyPlayInfoListPlayInfo {
	s.Size = &v
	return s
}

func (s *GetPlayInfoResponseBodyPlayInfoListPlayInfo) SetSpecification(v string) *GetPlayInfoResponseBodyPlayInfoListPlayInfo {
	s.Specification = &v
	return s
}

func (s *GetPlayInfoResponseBodyPlayInfoListPlayInfo) SetStatus(v string) *GetPlayInfoResponseBodyPlayInfoListPlayInfo {
	s.Status = &v
	return s
}

func (s *GetPlayInfoResponseBodyPlayInfoListPlayInfo) SetStreamType(v string) *GetPlayInfoResponseBodyPlayInfoListPlayInfo {
	s.StreamType = &v
	return s
}

func (s *GetPlayInfoResponseBodyPlayInfoListPlayInfo) SetWatermarkId(v string) *GetPlayInfoResponseBodyPlayInfoListPlayInfo {
	s.WatermarkId = &v
	return s
}

func (s *GetPlayInfoResponseBodyPlayInfoListPlayInfo) SetWidth(v int64) *GetPlayInfoResponseBodyPlayInfoListPlayInfo {
	s.Width = &v
	return s
}

type GetPlayInfoResponseBodyVideoBase struct {
	CoverURL     *string `json:"CoverURL,omitempty" xml:"CoverURL,omitempty"`
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	DanMuURL     *string `json:"DanMuURL,omitempty" xml:"DanMuURL,omitempty"`
	Duration     *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	MediaType    *string `json:"MediaType,omitempty" xml:"MediaType,omitempty"`
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Title        *string `json:"Title,omitempty" xml:"Title,omitempty"`
	VideoId      *string `json:"VideoId,omitempty" xml:"VideoId,omitempty"`
}

func (s GetPlayInfoResponseBodyVideoBase) String() string {
	return tea.Prettify(s)
}

func (s GetPlayInfoResponseBodyVideoBase) GoString() string {
	return s.String()
}

func (s *GetPlayInfoResponseBodyVideoBase) SetCoverURL(v string) *GetPlayInfoResponseBodyVideoBase {
	s.CoverURL = &v
	return s
}

func (s *GetPlayInfoResponseBodyVideoBase) SetCreationTime(v string) *GetPlayInfoResponseBodyVideoBase {
	s.CreationTime = &v
	return s
}

func (s *GetPlayInfoResponseBodyVideoBase) SetDanMuURL(v string) *GetPlayInfoResponseBodyVideoBase {
	s.DanMuURL = &v
	return s
}

func (s *GetPlayInfoResponseBodyVideoBase) SetDuration(v string) *GetPlayInfoResponseBodyVideoBase {
	s.Duration = &v
	return s
}

func (s *GetPlayInfoResponseBodyVideoBase) SetMediaType(v string) *GetPlayInfoResponseBodyVideoBase {
	s.MediaType = &v
	return s
}

func (s *GetPlayInfoResponseBodyVideoBase) SetStatus(v string) *GetPlayInfoResponseBodyVideoBase {
	s.Status = &v
	return s
}

func (s *GetPlayInfoResponseBodyVideoBase) SetTitle(v string) *GetPlayInfoResponseBodyVideoBase {
	s.Title = &v
	return s
}

func (s *GetPlayInfoResponseBodyVideoBase) SetVideoId(v string) *GetPlayInfoResponseBodyVideoBase {
	s.VideoId = &v
	return s
}

type GetPlayInfoResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetPlayInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetPlayInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetPlayInfoResponse) GoString() string {
	return s.String()
}

func (s *GetPlayInfoResponse) SetHeaders(v map[string]*string) *GetPlayInfoResponse {
	s.Headers = v
	return s
}

func (s *GetPlayInfoResponse) SetBody(v *GetPlayInfoResponseBody) *GetPlayInfoResponse {
	s.Body = v
	return s
}

type GetTranscodeSummaryRequest struct {
	VideoIds *string `json:"VideoIds,omitempty" xml:"VideoIds,omitempty"`
}

func (s GetTranscodeSummaryRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTranscodeSummaryRequest) GoString() string {
	return s.String()
}

func (s *GetTranscodeSummaryRequest) SetVideoIds(v string) *GetTranscodeSummaryRequest {
	s.VideoIds = &v
	return s
}

type GetTranscodeSummaryResponseBody struct {
	NonExistVideoIds     []*string                                              `json:"NonExistVideoIds,omitempty" xml:"NonExistVideoIds,omitempty" type:"Repeated"`
	RequestId            *string                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TranscodeSummaryList []*GetTranscodeSummaryResponseBodyTranscodeSummaryList `json:"TranscodeSummaryList,omitempty" xml:"TranscodeSummaryList,omitempty" type:"Repeated"`
}

func (s GetTranscodeSummaryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTranscodeSummaryResponseBody) GoString() string {
	return s.String()
}

func (s *GetTranscodeSummaryResponseBody) SetNonExistVideoIds(v []*string) *GetTranscodeSummaryResponseBody {
	s.NonExistVideoIds = v
	return s
}

func (s *GetTranscodeSummaryResponseBody) SetRequestId(v string) *GetTranscodeSummaryResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTranscodeSummaryResponseBody) SetTranscodeSummaryList(v []*GetTranscodeSummaryResponseBodyTranscodeSummaryList) *GetTranscodeSummaryResponseBody {
	s.TranscodeSummaryList = v
	return s
}

type GetTranscodeSummaryResponseBodyTranscodeSummaryList struct {
	CompleteTime                *string                                                                           `json:"CompleteTime,omitempty" xml:"CompleteTime,omitempty"`
	CreationTime                *string                                                                           `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	TranscodeJobInfoSummaryList []*GetTranscodeSummaryResponseBodyTranscodeSummaryListTranscodeJobInfoSummaryList `json:"TranscodeJobInfoSummaryList,omitempty" xml:"TranscodeJobInfoSummaryList,omitempty" type:"Repeated"`
	TranscodeStatus             *string                                                                           `json:"TranscodeStatus,omitempty" xml:"TranscodeStatus,omitempty"`
	TranscodeTemplateGroupId    *string                                                                           `json:"TranscodeTemplateGroupId,omitempty" xml:"TranscodeTemplateGroupId,omitempty"`
	VideoId                     *string                                                                           `json:"VideoId,omitempty" xml:"VideoId,omitempty"`
}

func (s GetTranscodeSummaryResponseBodyTranscodeSummaryList) String() string {
	return tea.Prettify(s)
}

func (s GetTranscodeSummaryResponseBodyTranscodeSummaryList) GoString() string {
	return s.String()
}

func (s *GetTranscodeSummaryResponseBodyTranscodeSummaryList) SetCompleteTime(v string) *GetTranscodeSummaryResponseBodyTranscodeSummaryList {
	s.CompleteTime = &v
	return s
}

func (s *GetTranscodeSummaryResponseBodyTranscodeSummaryList) SetCreationTime(v string) *GetTranscodeSummaryResponseBodyTranscodeSummaryList {
	s.CreationTime = &v
	return s
}

func (s *GetTranscodeSummaryResponseBodyTranscodeSummaryList) SetTranscodeJobInfoSummaryList(v []*GetTranscodeSummaryResponseBodyTranscodeSummaryListTranscodeJobInfoSummaryList) *GetTranscodeSummaryResponseBodyTranscodeSummaryList {
	s.TranscodeJobInfoSummaryList = v
	return s
}

func (s *GetTranscodeSummaryResponseBodyTranscodeSummaryList) SetTranscodeStatus(v string) *GetTranscodeSummaryResponseBodyTranscodeSummaryList {
	s.TranscodeStatus = &v
	return s
}

func (s *GetTranscodeSummaryResponseBodyTranscodeSummaryList) SetTranscodeTemplateGroupId(v string) *GetTranscodeSummaryResponseBodyTranscodeSummaryList {
	s.TranscodeTemplateGroupId = &v
	return s
}

func (s *GetTranscodeSummaryResponseBodyTranscodeSummaryList) SetVideoId(v string) *GetTranscodeSummaryResponseBodyTranscodeSummaryList {
	s.VideoId = &v
	return s
}

type GetTranscodeSummaryResponseBodyTranscodeSummaryListTranscodeJobInfoSummaryList struct {
	Bitrate             *string   `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	CompleteTime        *string   `json:"CompleteTime,omitempty" xml:"CompleteTime,omitempty"`
	CreationTime        *string   `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	Duration            *string   `json:"Duration,omitempty" xml:"Duration,omitempty"`
	ErrorCode           *string   `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage        *string   `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Filesize            *int64    `json:"Filesize,omitempty" xml:"Filesize,omitempty"`
	Format              *string   `json:"Format,omitempty" xml:"Format,omitempty"`
	Fps                 *string   `json:"Fps,omitempty" xml:"Fps,omitempty"`
	Height              *string   `json:"Height,omitempty" xml:"Height,omitempty"`
	TranscodeJobStatus  *string   `json:"TranscodeJobStatus,omitempty" xml:"TranscodeJobStatus,omitempty"`
	TranscodeProgress   *int64    `json:"TranscodeProgress,omitempty" xml:"TranscodeProgress,omitempty"`
	TranscodeTemplateId *string   `json:"TranscodeTemplateId,omitempty" xml:"TranscodeTemplateId,omitempty"`
	WatermarkIdList     []*string `json:"WatermarkIdList,omitempty" xml:"WatermarkIdList,omitempty" type:"Repeated"`
	Width               *string   `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s GetTranscodeSummaryResponseBodyTranscodeSummaryListTranscodeJobInfoSummaryList) String() string {
	return tea.Prettify(s)
}

func (s GetTranscodeSummaryResponseBodyTranscodeSummaryListTranscodeJobInfoSummaryList) GoString() string {
	return s.String()
}

func (s *GetTranscodeSummaryResponseBodyTranscodeSummaryListTranscodeJobInfoSummaryList) SetBitrate(v string) *GetTranscodeSummaryResponseBodyTranscodeSummaryListTranscodeJobInfoSummaryList {
	s.Bitrate = &v
	return s
}

func (s *GetTranscodeSummaryResponseBodyTranscodeSummaryListTranscodeJobInfoSummaryList) SetCompleteTime(v string) *GetTranscodeSummaryResponseBodyTranscodeSummaryListTranscodeJobInfoSummaryList {
	s.CompleteTime = &v
	return s
}

func (s *GetTranscodeSummaryResponseBodyTranscodeSummaryListTranscodeJobInfoSummaryList) SetCreationTime(v string) *GetTranscodeSummaryResponseBodyTranscodeSummaryListTranscodeJobInfoSummaryList {
	s.CreationTime = &v
	return s
}

func (s *GetTranscodeSummaryResponseBodyTranscodeSummaryListTranscodeJobInfoSummaryList) SetDuration(v string) *GetTranscodeSummaryResponseBodyTranscodeSummaryListTranscodeJobInfoSummaryList {
	s.Duration = &v
	return s
}

func (s *GetTranscodeSummaryResponseBodyTranscodeSummaryListTranscodeJobInfoSummaryList) SetErrorCode(v string) *GetTranscodeSummaryResponseBodyTranscodeSummaryListTranscodeJobInfoSummaryList {
	s.ErrorCode = &v
	return s
}

func (s *GetTranscodeSummaryResponseBodyTranscodeSummaryListTranscodeJobInfoSummaryList) SetErrorMessage(v string) *GetTranscodeSummaryResponseBodyTranscodeSummaryListTranscodeJobInfoSummaryList {
	s.ErrorMessage = &v
	return s
}

func (s *GetTranscodeSummaryResponseBodyTranscodeSummaryListTranscodeJobInfoSummaryList) SetFilesize(v int64) *GetTranscodeSummaryResponseBodyTranscodeSummaryListTranscodeJobInfoSummaryList {
	s.Filesize = &v
	return s
}

func (s *GetTranscodeSummaryResponseBodyTranscodeSummaryListTranscodeJobInfoSummaryList) SetFormat(v string) *GetTranscodeSummaryResponseBodyTranscodeSummaryListTranscodeJobInfoSummaryList {
	s.Format = &v
	return s
}

func (s *GetTranscodeSummaryResponseBodyTranscodeSummaryListTranscodeJobInfoSummaryList) SetFps(v string) *GetTranscodeSummaryResponseBodyTranscodeSummaryListTranscodeJobInfoSummaryList {
	s.Fps = &v
	return s
}

func (s *GetTranscodeSummaryResponseBodyTranscodeSummaryListTranscodeJobInfoSummaryList) SetHeight(v string) *GetTranscodeSummaryResponseBodyTranscodeSummaryListTranscodeJobInfoSummaryList {
	s.Height = &v
	return s
}

func (s *GetTranscodeSummaryResponseBodyTranscodeSummaryListTranscodeJobInfoSummaryList) SetTranscodeJobStatus(v string) *GetTranscodeSummaryResponseBodyTranscodeSummaryListTranscodeJobInfoSummaryList {
	s.TranscodeJobStatus = &v
	return s
}

func (s *GetTranscodeSummaryResponseBodyTranscodeSummaryListTranscodeJobInfoSummaryList) SetTranscodeProgress(v int64) *GetTranscodeSummaryResponseBodyTranscodeSummaryListTranscodeJobInfoSummaryList {
	s.TranscodeProgress = &v
	return s
}

func (s *GetTranscodeSummaryResponseBodyTranscodeSummaryListTranscodeJobInfoSummaryList) SetTranscodeTemplateId(v string) *GetTranscodeSummaryResponseBodyTranscodeSummaryListTranscodeJobInfoSummaryList {
	s.TranscodeTemplateId = &v
	return s
}

func (s *GetTranscodeSummaryResponseBodyTranscodeSummaryListTranscodeJobInfoSummaryList) SetWatermarkIdList(v []*string) *GetTranscodeSummaryResponseBodyTranscodeSummaryListTranscodeJobInfoSummaryList {
	s.WatermarkIdList = v
	return s
}

func (s *GetTranscodeSummaryResponseBodyTranscodeSummaryListTranscodeJobInfoSummaryList) SetWidth(v string) *GetTranscodeSummaryResponseBodyTranscodeSummaryListTranscodeJobInfoSummaryList {
	s.Width = &v
	return s
}

type GetTranscodeSummaryResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetTranscodeSummaryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetTranscodeSummaryResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTranscodeSummaryResponse) GoString() string {
	return s.String()
}

func (s *GetTranscodeSummaryResponse) SetHeaders(v map[string]*string) *GetTranscodeSummaryResponse {
	s.Headers = v
	return s
}

func (s *GetTranscodeSummaryResponse) SetBody(v *GetTranscodeSummaryResponseBody) *GetTranscodeSummaryResponse {
	s.Body = v
	return s
}

type GetTranscodeTaskRequest struct {
	TranscodeTaskId *string `json:"TranscodeTaskId,omitempty" xml:"TranscodeTaskId,omitempty"`
}

func (s GetTranscodeTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTranscodeTaskRequest) GoString() string {
	return s.String()
}

func (s *GetTranscodeTaskRequest) SetTranscodeTaskId(v string) *GetTranscodeTaskRequest {
	s.TranscodeTaskId = &v
	return s
}

type GetTranscodeTaskResponseBody struct {
	RequestId     *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TranscodeTask *GetTranscodeTaskResponseBodyTranscodeTask `json:"TranscodeTask,omitempty" xml:"TranscodeTask,omitempty" type:"Struct"`
}

func (s GetTranscodeTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTranscodeTaskResponseBody) GoString() string {
	return s.String()
}

func (s *GetTranscodeTaskResponseBody) SetRequestId(v string) *GetTranscodeTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTranscodeTaskResponseBody) SetTranscodeTask(v *GetTranscodeTaskResponseBodyTranscodeTask) *GetTranscodeTaskResponseBody {
	s.TranscodeTask = v
	return s
}

type GetTranscodeTaskResponseBodyTranscodeTask struct {
	CompleteTime             *string                                                          `json:"CompleteTime,omitempty" xml:"CompleteTime,omitempty"`
	CreationTime             *string                                                          `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	TaskStatus               *string                                                          `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	TranscodeJobInfoList     []*GetTranscodeTaskResponseBodyTranscodeTaskTranscodeJobInfoList `json:"TranscodeJobInfoList,omitempty" xml:"TranscodeJobInfoList,omitempty" type:"Repeated"`
	TranscodeTaskId          *string                                                          `json:"TranscodeTaskId,omitempty" xml:"TranscodeTaskId,omitempty"`
	TranscodeTemplateGroupId *string                                                          `json:"TranscodeTemplateGroupId,omitempty" xml:"TranscodeTemplateGroupId,omitempty"`
	Trigger                  *string                                                          `json:"Trigger,omitempty" xml:"Trigger,omitempty"`
	VideoId                  *string                                                          `json:"VideoId,omitempty" xml:"VideoId,omitempty"`
}

func (s GetTranscodeTaskResponseBodyTranscodeTask) String() string {
	return tea.Prettify(s)
}

func (s GetTranscodeTaskResponseBodyTranscodeTask) GoString() string {
	return s.String()
}

func (s *GetTranscodeTaskResponseBodyTranscodeTask) SetCompleteTime(v string) *GetTranscodeTaskResponseBodyTranscodeTask {
	s.CompleteTime = &v
	return s
}

func (s *GetTranscodeTaskResponseBodyTranscodeTask) SetCreationTime(v string) *GetTranscodeTaskResponseBodyTranscodeTask {
	s.CreationTime = &v
	return s
}

func (s *GetTranscodeTaskResponseBodyTranscodeTask) SetTaskStatus(v string) *GetTranscodeTaskResponseBodyTranscodeTask {
	s.TaskStatus = &v
	return s
}

func (s *GetTranscodeTaskResponseBodyTranscodeTask) SetTranscodeJobInfoList(v []*GetTranscodeTaskResponseBodyTranscodeTaskTranscodeJobInfoList) *GetTranscodeTaskResponseBodyTranscodeTask {
	s.TranscodeJobInfoList = v
	return s
}

func (s *GetTranscodeTaskResponseBodyTranscodeTask) SetTranscodeTaskId(v string) *GetTranscodeTaskResponseBodyTranscodeTask {
	s.TranscodeTaskId = &v
	return s
}

func (s *GetTranscodeTaskResponseBodyTranscodeTask) SetTranscodeTemplateGroupId(v string) *GetTranscodeTaskResponseBodyTranscodeTask {
	s.TranscodeTemplateGroupId = &v
	return s
}

func (s *GetTranscodeTaskResponseBodyTranscodeTask) SetTrigger(v string) *GetTranscodeTaskResponseBodyTranscodeTask {
	s.Trigger = &v
	return s
}

func (s *GetTranscodeTaskResponseBodyTranscodeTask) SetVideoId(v string) *GetTranscodeTaskResponseBodyTranscodeTask {
	s.VideoId = &v
	return s
}

type GetTranscodeTaskResponseBodyTranscodeTaskTranscodeJobInfoList struct {
	CompleteTime        *string                                                                  `json:"CompleteTime,omitempty" xml:"CompleteTime,omitempty"`
	CreationTime        *string                                                                  `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	Definition          *string                                                                  `json:"Definition,omitempty" xml:"Definition,omitempty"`
	ErrorCode           *string                                                                  `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage        *string                                                                  `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	InputFileUrl        *string                                                                  `json:"InputFileUrl,omitempty" xml:"InputFileUrl,omitempty"`
	OutputFile          *GetTranscodeTaskResponseBodyTranscodeTaskTranscodeJobInfoListOutputFile `json:"OutputFile,omitempty" xml:"OutputFile,omitempty" type:"Struct"`
	Priority            *string                                                                  `json:"Priority,omitempty" xml:"Priority,omitempty"`
	TranscodeJobId      *string                                                                  `json:"TranscodeJobId,omitempty" xml:"TranscodeJobId,omitempty"`
	TranscodeJobStatus  *string                                                                  `json:"TranscodeJobStatus,omitempty" xml:"TranscodeJobStatus,omitempty"`
	TranscodeProgress   *int64                                                                   `json:"TranscodeProgress,omitempty" xml:"TranscodeProgress,omitempty"`
	TranscodeTemplateId *string                                                                  `json:"TranscodeTemplateId,omitempty" xml:"TranscodeTemplateId,omitempty"`
}

func (s GetTranscodeTaskResponseBodyTranscodeTaskTranscodeJobInfoList) String() string {
	return tea.Prettify(s)
}

func (s GetTranscodeTaskResponseBodyTranscodeTaskTranscodeJobInfoList) GoString() string {
	return s.String()
}

func (s *GetTranscodeTaskResponseBodyTranscodeTaskTranscodeJobInfoList) SetCompleteTime(v string) *GetTranscodeTaskResponseBodyTranscodeTaskTranscodeJobInfoList {
	s.CompleteTime = &v
	return s
}

func (s *GetTranscodeTaskResponseBodyTranscodeTaskTranscodeJobInfoList) SetCreationTime(v string) *GetTranscodeTaskResponseBodyTranscodeTaskTranscodeJobInfoList {
	s.CreationTime = &v
	return s
}

func (s *GetTranscodeTaskResponseBodyTranscodeTaskTranscodeJobInfoList) SetDefinition(v string) *GetTranscodeTaskResponseBodyTranscodeTaskTranscodeJobInfoList {
	s.Definition = &v
	return s
}

func (s *GetTranscodeTaskResponseBodyTranscodeTaskTranscodeJobInfoList) SetErrorCode(v string) *GetTranscodeTaskResponseBodyTranscodeTaskTranscodeJobInfoList {
	s.ErrorCode = &v
	return s
}

func (s *GetTranscodeTaskResponseBodyTranscodeTaskTranscodeJobInfoList) SetErrorMessage(v string) *GetTranscodeTaskResponseBodyTranscodeTaskTranscodeJobInfoList {
	s.ErrorMessage = &v
	return s
}

func (s *GetTranscodeTaskResponseBodyTranscodeTaskTranscodeJobInfoList) SetInputFileUrl(v string) *GetTranscodeTaskResponseBodyTranscodeTaskTranscodeJobInfoList {
	s.InputFileUrl = &v
	return s
}

func (s *GetTranscodeTaskResponseBodyTranscodeTaskTranscodeJobInfoList) SetOutputFile(v *GetTranscodeTaskResponseBodyTranscodeTaskTranscodeJobInfoListOutputFile) *GetTranscodeTaskResponseBodyTranscodeTaskTranscodeJobInfoList {
	s.OutputFile = v
	return s
}

func (s *GetTranscodeTaskResponseBodyTranscodeTaskTranscodeJobInfoList) SetPriority(v string) *GetTranscodeTaskResponseBodyTranscodeTaskTranscodeJobInfoList {
	s.Priority = &v
	return s
}

func (s *GetTranscodeTaskResponseBodyTranscodeTaskTranscodeJobInfoList) SetTranscodeJobId(v string) *GetTranscodeTaskResponseBodyTranscodeTaskTranscodeJobInfoList {
	s.TranscodeJobId = &v
	return s
}

func (s *GetTranscodeTaskResponseBodyTranscodeTaskTranscodeJobInfoList) SetTranscodeJobStatus(v string) *GetTranscodeTaskResponseBodyTranscodeTaskTranscodeJobInfoList {
	s.TranscodeJobStatus = &v
	return s
}

func (s *GetTranscodeTaskResponseBodyTranscodeTaskTranscodeJobInfoList) SetTranscodeProgress(v int64) *GetTranscodeTaskResponseBodyTranscodeTaskTranscodeJobInfoList {
	s.TranscodeProgress = &v
	return s
}

func (s *GetTranscodeTaskResponseBodyTranscodeTaskTranscodeJobInfoList) SetTranscodeTemplateId(v string) *GetTranscodeTaskResponseBodyTranscodeTaskTranscodeJobInfoList {
	s.TranscodeTemplateId = &v
	return s
}

type GetTranscodeTaskResponseBodyTranscodeTaskTranscodeJobInfoListOutputFile struct {
	AudioStreamList    *string   `json:"AudioStreamList,omitempty" xml:"AudioStreamList,omitempty"`
	Bitrate            *string   `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	Duration           *string   `json:"Duration,omitempty" xml:"Duration,omitempty"`
	Encryption         *string   `json:"Encryption,omitempty" xml:"Encryption,omitempty"`
	Filesize           *int64    `json:"Filesize,omitempty" xml:"Filesize,omitempty"`
	Format             *string   `json:"Format,omitempty" xml:"Format,omitempty"`
	Fps                *string   `json:"Fps,omitempty" xml:"Fps,omitempty"`
	Height             *string   `json:"Height,omitempty" xml:"Height,omitempty"`
	OutputFileUrl      *string   `json:"OutputFileUrl,omitempty" xml:"OutputFileUrl,omitempty"`
	SubtitleStreamList *string   `json:"SubtitleStreamList,omitempty" xml:"SubtitleStreamList,omitempty"`
	VideoStreamList    *string   `json:"VideoStreamList,omitempty" xml:"VideoStreamList,omitempty"`
	WatermarkIdList    []*string `json:"WatermarkIdList,omitempty" xml:"WatermarkIdList,omitempty" type:"Repeated"`
	Width              *string   `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s GetTranscodeTaskResponseBodyTranscodeTaskTranscodeJobInfoListOutputFile) String() string {
	return tea.Prettify(s)
}

func (s GetTranscodeTaskResponseBodyTranscodeTaskTranscodeJobInfoListOutputFile) GoString() string {
	return s.String()
}

func (s *GetTranscodeTaskResponseBodyTranscodeTaskTranscodeJobInfoListOutputFile) SetAudioStreamList(v string) *GetTranscodeTaskResponseBodyTranscodeTaskTranscodeJobInfoListOutputFile {
	s.AudioStreamList = &v
	return s
}

func (s *GetTranscodeTaskResponseBodyTranscodeTaskTranscodeJobInfoListOutputFile) SetBitrate(v string) *GetTranscodeTaskResponseBodyTranscodeTaskTranscodeJobInfoListOutputFile {
	s.Bitrate = &v
	return s
}

func (s *GetTranscodeTaskResponseBodyTranscodeTaskTranscodeJobInfoListOutputFile) SetDuration(v string) *GetTranscodeTaskResponseBodyTranscodeTaskTranscodeJobInfoListOutputFile {
	s.Duration = &v
	return s
}

func (s *GetTranscodeTaskResponseBodyTranscodeTaskTranscodeJobInfoListOutputFile) SetEncryption(v string) *GetTranscodeTaskResponseBodyTranscodeTaskTranscodeJobInfoListOutputFile {
	s.Encryption = &v
	return s
}

func (s *GetTranscodeTaskResponseBodyTranscodeTaskTranscodeJobInfoListOutputFile) SetFilesize(v int64) *GetTranscodeTaskResponseBodyTranscodeTaskTranscodeJobInfoListOutputFile {
	s.Filesize = &v
	return s
}

func (s *GetTranscodeTaskResponseBodyTranscodeTaskTranscodeJobInfoListOutputFile) SetFormat(v string) *GetTranscodeTaskResponseBodyTranscodeTaskTranscodeJobInfoListOutputFile {
	s.Format = &v
	return s
}

func (s *GetTranscodeTaskResponseBodyTranscodeTaskTranscodeJobInfoListOutputFile) SetFps(v string) *GetTranscodeTaskResponseBodyTranscodeTaskTranscodeJobInfoListOutputFile {
	s.Fps = &v
	return s
}

func (s *GetTranscodeTaskResponseBodyTranscodeTaskTranscodeJobInfoListOutputFile) SetHeight(v string) *GetTranscodeTaskResponseBodyTranscodeTaskTranscodeJobInfoListOutputFile {
	s.Height = &v
	return s
}

func (s *GetTranscodeTaskResponseBodyTranscodeTaskTranscodeJobInfoListOutputFile) SetOutputFileUrl(v string) *GetTranscodeTaskResponseBodyTranscodeTaskTranscodeJobInfoListOutputFile {
	s.OutputFileUrl = &v
	return s
}

func (s *GetTranscodeTaskResponseBodyTranscodeTaskTranscodeJobInfoListOutputFile) SetSubtitleStreamList(v string) *GetTranscodeTaskResponseBodyTranscodeTaskTranscodeJobInfoListOutputFile {
	s.SubtitleStreamList = &v
	return s
}

func (s *GetTranscodeTaskResponseBodyTranscodeTaskTranscodeJobInfoListOutputFile) SetVideoStreamList(v string) *GetTranscodeTaskResponseBodyTranscodeTaskTranscodeJobInfoListOutputFile {
	s.VideoStreamList = &v
	return s
}

func (s *GetTranscodeTaskResponseBodyTranscodeTaskTranscodeJobInfoListOutputFile) SetWatermarkIdList(v []*string) *GetTranscodeTaskResponseBodyTranscodeTaskTranscodeJobInfoListOutputFile {
	s.WatermarkIdList = v
	return s
}

func (s *GetTranscodeTaskResponseBodyTranscodeTaskTranscodeJobInfoListOutputFile) SetWidth(v string) *GetTranscodeTaskResponseBodyTranscodeTaskTranscodeJobInfoListOutputFile {
	s.Width = &v
	return s
}

type GetTranscodeTaskResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetTranscodeTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetTranscodeTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTranscodeTaskResponse) GoString() string {
	return s.String()
}

func (s *GetTranscodeTaskResponse) SetHeaders(v map[string]*string) *GetTranscodeTaskResponse {
	s.Headers = v
	return s
}

func (s *GetTranscodeTaskResponse) SetBody(v *GetTranscodeTaskResponseBody) *GetTranscodeTaskResponse {
	s.Body = v
	return s
}

type GetTranscodeTemplateGroupRequest struct {
	TranscodeTemplateGroupId *string `json:"TranscodeTemplateGroupId,omitempty" xml:"TranscodeTemplateGroupId,omitempty"`
}

func (s GetTranscodeTemplateGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTranscodeTemplateGroupRequest) GoString() string {
	return s.String()
}

func (s *GetTranscodeTemplateGroupRequest) SetTranscodeTemplateGroupId(v string) *GetTranscodeTemplateGroupRequest {
	s.TranscodeTemplateGroupId = &v
	return s
}

type GetTranscodeTemplateGroupResponseBody struct {
	RequestId              *string                                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TranscodeTemplateGroup *GetTranscodeTemplateGroupResponseBodyTranscodeTemplateGroup `json:"TranscodeTemplateGroup,omitempty" xml:"TranscodeTemplateGroup,omitempty" type:"Struct"`
}

func (s GetTranscodeTemplateGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTranscodeTemplateGroupResponseBody) GoString() string {
	return s.String()
}

func (s *GetTranscodeTemplateGroupResponseBody) SetRequestId(v string) *GetTranscodeTemplateGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTranscodeTemplateGroupResponseBody) SetTranscodeTemplateGroup(v *GetTranscodeTemplateGroupResponseBodyTranscodeTemplateGroup) *GetTranscodeTemplateGroupResponseBody {
	s.TranscodeTemplateGroup = v
	return s
}

type GetTranscodeTemplateGroupResponseBodyTranscodeTemplateGroup struct {
	AppId                    *string                                                                             `json:"AppId,omitempty" xml:"AppId,omitempty"`
	CreationTime             *string                                                                             `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	IsDefault                *string                                                                             `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	Locked                   *string                                                                             `json:"Locked,omitempty" xml:"Locked,omitempty"`
	ModifyTime               *string                                                                             `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	Name                     *string                                                                             `json:"Name,omitempty" xml:"Name,omitempty"`
	TranscodeTemplateGroupId *string                                                                             `json:"TranscodeTemplateGroupId,omitempty" xml:"TranscodeTemplateGroupId,omitempty"`
	TranscodeTemplateList    []*GetTranscodeTemplateGroupResponseBodyTranscodeTemplateGroupTranscodeTemplateList `json:"TranscodeTemplateList,omitempty" xml:"TranscodeTemplateList,omitempty" type:"Repeated"`
}

func (s GetTranscodeTemplateGroupResponseBodyTranscodeTemplateGroup) String() string {
	return tea.Prettify(s)
}

func (s GetTranscodeTemplateGroupResponseBodyTranscodeTemplateGroup) GoString() string {
	return s.String()
}

func (s *GetTranscodeTemplateGroupResponseBodyTranscodeTemplateGroup) SetAppId(v string) *GetTranscodeTemplateGroupResponseBodyTranscodeTemplateGroup {
	s.AppId = &v
	return s
}

func (s *GetTranscodeTemplateGroupResponseBodyTranscodeTemplateGroup) SetCreationTime(v string) *GetTranscodeTemplateGroupResponseBodyTranscodeTemplateGroup {
	s.CreationTime = &v
	return s
}

func (s *GetTranscodeTemplateGroupResponseBodyTranscodeTemplateGroup) SetIsDefault(v string) *GetTranscodeTemplateGroupResponseBodyTranscodeTemplateGroup {
	s.IsDefault = &v
	return s
}

func (s *GetTranscodeTemplateGroupResponseBodyTranscodeTemplateGroup) SetLocked(v string) *GetTranscodeTemplateGroupResponseBodyTranscodeTemplateGroup {
	s.Locked = &v
	return s
}

func (s *GetTranscodeTemplateGroupResponseBodyTranscodeTemplateGroup) SetModifyTime(v string) *GetTranscodeTemplateGroupResponseBodyTranscodeTemplateGroup {
	s.ModifyTime = &v
	return s
}

func (s *GetTranscodeTemplateGroupResponseBodyTranscodeTemplateGroup) SetName(v string) *GetTranscodeTemplateGroupResponseBodyTranscodeTemplateGroup {
	s.Name = &v
	return s
}

func (s *GetTranscodeTemplateGroupResponseBodyTranscodeTemplateGroup) SetTranscodeTemplateGroupId(v string) *GetTranscodeTemplateGroupResponseBodyTranscodeTemplateGroup {
	s.TranscodeTemplateGroupId = &v
	return s
}

func (s *GetTranscodeTemplateGroupResponseBodyTranscodeTemplateGroup) SetTranscodeTemplateList(v []*GetTranscodeTemplateGroupResponseBodyTranscodeTemplateGroupTranscodeTemplateList) *GetTranscodeTemplateGroupResponseBodyTranscodeTemplateGroup {
	s.TranscodeTemplateList = v
	return s
}

type GetTranscodeTemplateGroupResponseBodyTranscodeTemplateGroupTranscodeTemplateList struct {
	Audio                *string   `json:"Audio,omitempty" xml:"Audio,omitempty"`
	Clip                 *string   `json:"Clip,omitempty" xml:"Clip,omitempty"`
	Container            *string   `json:"Container,omitempty" xml:"Container,omitempty"`
	Definition           *string   `json:"Definition,omitempty" xml:"Definition,omitempty"`
	EncryptSetting       *string   `json:"EncryptSetting,omitempty" xml:"EncryptSetting,omitempty"`
	MuxConfig            *string   `json:"MuxConfig,omitempty" xml:"MuxConfig,omitempty"`
	PackageSetting       *string   `json:"PackageSetting,omitempty" xml:"PackageSetting,omitempty"`
	Rotate               *string   `json:"Rotate,omitempty" xml:"Rotate,omitempty"`
	SubtitleList         *string   `json:"SubtitleList,omitempty" xml:"SubtitleList,omitempty"`
	TemplateName         *string   `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	TransConfig          *string   `json:"TransConfig,omitempty" xml:"TransConfig,omitempty"`
	TranscodeFileRegular *string   `json:"TranscodeFileRegular,omitempty" xml:"TranscodeFileRegular,omitempty"`
	TranscodeTemplateId  *string   `json:"TranscodeTemplateId,omitempty" xml:"TranscodeTemplateId,omitempty"`
	Type                 *string   `json:"Type,omitempty" xml:"Type,omitempty"`
	Video                *string   `json:"Video,omitempty" xml:"Video,omitempty"`
	WatermarkIds         []*string `json:"WatermarkIds,omitempty" xml:"WatermarkIds,omitempty" type:"Repeated"`
}

func (s GetTranscodeTemplateGroupResponseBodyTranscodeTemplateGroupTranscodeTemplateList) String() string {
	return tea.Prettify(s)
}

func (s GetTranscodeTemplateGroupResponseBodyTranscodeTemplateGroupTranscodeTemplateList) GoString() string {
	return s.String()
}

func (s *GetTranscodeTemplateGroupResponseBodyTranscodeTemplateGroupTranscodeTemplateList) SetAudio(v string) *GetTranscodeTemplateGroupResponseBodyTranscodeTemplateGroupTranscodeTemplateList {
	s.Audio = &v
	return s
}

func (s *GetTranscodeTemplateGroupResponseBodyTranscodeTemplateGroupTranscodeTemplateList) SetClip(v string) *GetTranscodeTemplateGroupResponseBodyTranscodeTemplateGroupTranscodeTemplateList {
	s.Clip = &v
	return s
}

func (s *GetTranscodeTemplateGroupResponseBodyTranscodeTemplateGroupTranscodeTemplateList) SetContainer(v string) *GetTranscodeTemplateGroupResponseBodyTranscodeTemplateGroupTranscodeTemplateList {
	s.Container = &v
	return s
}

func (s *GetTranscodeTemplateGroupResponseBodyTranscodeTemplateGroupTranscodeTemplateList) SetDefinition(v string) *GetTranscodeTemplateGroupResponseBodyTranscodeTemplateGroupTranscodeTemplateList {
	s.Definition = &v
	return s
}

func (s *GetTranscodeTemplateGroupResponseBodyTranscodeTemplateGroupTranscodeTemplateList) SetEncryptSetting(v string) *GetTranscodeTemplateGroupResponseBodyTranscodeTemplateGroupTranscodeTemplateList {
	s.EncryptSetting = &v
	return s
}

func (s *GetTranscodeTemplateGroupResponseBodyTranscodeTemplateGroupTranscodeTemplateList) SetMuxConfig(v string) *GetTranscodeTemplateGroupResponseBodyTranscodeTemplateGroupTranscodeTemplateList {
	s.MuxConfig = &v
	return s
}

func (s *GetTranscodeTemplateGroupResponseBodyTranscodeTemplateGroupTranscodeTemplateList) SetPackageSetting(v string) *GetTranscodeTemplateGroupResponseBodyTranscodeTemplateGroupTranscodeTemplateList {
	s.PackageSetting = &v
	return s
}

func (s *GetTranscodeTemplateGroupResponseBodyTranscodeTemplateGroupTranscodeTemplateList) SetRotate(v string) *GetTranscodeTemplateGroupResponseBodyTranscodeTemplateGroupTranscodeTemplateList {
	s.Rotate = &v
	return s
}

func (s *GetTranscodeTemplateGroupResponseBodyTranscodeTemplateGroupTranscodeTemplateList) SetSubtitleList(v string) *GetTranscodeTemplateGroupResponseBodyTranscodeTemplateGroupTranscodeTemplateList {
	s.SubtitleList = &v
	return s
}

func (s *GetTranscodeTemplateGroupResponseBodyTranscodeTemplateGroupTranscodeTemplateList) SetTemplateName(v string) *GetTranscodeTemplateGroupResponseBodyTranscodeTemplateGroupTranscodeTemplateList {
	s.TemplateName = &v
	return s
}

func (s *GetTranscodeTemplateGroupResponseBodyTranscodeTemplateGroupTranscodeTemplateList) SetTransConfig(v string) *GetTranscodeTemplateGroupResponseBodyTranscodeTemplateGroupTranscodeTemplateList {
	s.TransConfig = &v
	return s
}

func (s *GetTranscodeTemplateGroupResponseBodyTranscodeTemplateGroupTranscodeTemplateList) SetTranscodeFileRegular(v string) *GetTranscodeTemplateGroupResponseBodyTranscodeTemplateGroupTranscodeTemplateList {
	s.TranscodeFileRegular = &v
	return s
}

func (s *GetTranscodeTemplateGroupResponseBodyTranscodeTemplateGroupTranscodeTemplateList) SetTranscodeTemplateId(v string) *GetTranscodeTemplateGroupResponseBodyTranscodeTemplateGroupTranscodeTemplateList {
	s.TranscodeTemplateId = &v
	return s
}

func (s *GetTranscodeTemplateGroupResponseBodyTranscodeTemplateGroupTranscodeTemplateList) SetType(v string) *GetTranscodeTemplateGroupResponseBodyTranscodeTemplateGroupTranscodeTemplateList {
	s.Type = &v
	return s
}

func (s *GetTranscodeTemplateGroupResponseBodyTranscodeTemplateGroupTranscodeTemplateList) SetVideo(v string) *GetTranscodeTemplateGroupResponseBodyTranscodeTemplateGroupTranscodeTemplateList {
	s.Video = &v
	return s
}

func (s *GetTranscodeTemplateGroupResponseBodyTranscodeTemplateGroupTranscodeTemplateList) SetWatermarkIds(v []*string) *GetTranscodeTemplateGroupResponseBodyTranscodeTemplateGroupTranscodeTemplateList {
	s.WatermarkIds = v
	return s
}

type GetTranscodeTemplateGroupResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetTranscodeTemplateGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetTranscodeTemplateGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTranscodeTemplateGroupResponse) GoString() string {
	return s.String()
}

func (s *GetTranscodeTemplateGroupResponse) SetHeaders(v map[string]*string) *GetTranscodeTemplateGroupResponse {
	s.Headers = v
	return s
}

func (s *GetTranscodeTemplateGroupResponse) SetBody(v *GetTranscodeTemplateGroupResponseBody) *GetTranscodeTemplateGroupResponse {
	s.Body = v
	return s
}

type GetURLUploadInfosRequest struct {
	JobIds     *string `json:"JobIds,omitempty" xml:"JobIds,omitempty"`
	UploadURLs *string `json:"UploadURLs,omitempty" xml:"UploadURLs,omitempty"`
}

func (s GetURLUploadInfosRequest) String() string {
	return tea.Prettify(s)
}

func (s GetURLUploadInfosRequest) GoString() string {
	return s.String()
}

func (s *GetURLUploadInfosRequest) SetJobIds(v string) *GetURLUploadInfosRequest {
	s.JobIds = &v
	return s
}

func (s *GetURLUploadInfosRequest) SetUploadURLs(v string) *GetURLUploadInfosRequest {
	s.UploadURLs = &v
	return s
}

type GetURLUploadInfosResponseBody struct {
	NonExists         []*string                                         `json:"NonExists,omitempty" xml:"NonExists,omitempty" type:"Repeated"`
	RequestId         *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	URLUploadInfoList []*GetURLUploadInfosResponseBodyURLUploadInfoList `json:"URLUploadInfoList,omitempty" xml:"URLUploadInfoList,omitempty" type:"Repeated"`
}

func (s GetURLUploadInfosResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetURLUploadInfosResponseBody) GoString() string {
	return s.String()
}

func (s *GetURLUploadInfosResponseBody) SetNonExists(v []*string) *GetURLUploadInfosResponseBody {
	s.NonExists = v
	return s
}

func (s *GetURLUploadInfosResponseBody) SetRequestId(v string) *GetURLUploadInfosResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetURLUploadInfosResponseBody) SetURLUploadInfoList(v []*GetURLUploadInfosResponseBodyURLUploadInfoList) *GetURLUploadInfosResponseBody {
	s.URLUploadInfoList = v
	return s
}

type GetURLUploadInfosResponseBodyURLUploadInfoList struct {
	CompleteTime *string `json:"CompleteTime,omitempty" xml:"CompleteTime,omitempty"`
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	ErrorCode    *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	FileSize     *string `json:"FileSize,omitempty" xml:"FileSize,omitempty"`
	JobId        *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	MediaId      *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
	UploadURL    *string `json:"UploadURL,omitempty" xml:"UploadURL,omitempty"`
	UserData     *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s GetURLUploadInfosResponseBodyURLUploadInfoList) String() string {
	return tea.Prettify(s)
}

func (s GetURLUploadInfosResponseBodyURLUploadInfoList) GoString() string {
	return s.String()
}

func (s *GetURLUploadInfosResponseBodyURLUploadInfoList) SetCompleteTime(v string) *GetURLUploadInfosResponseBodyURLUploadInfoList {
	s.CompleteTime = &v
	return s
}

func (s *GetURLUploadInfosResponseBodyURLUploadInfoList) SetCreationTime(v string) *GetURLUploadInfosResponseBodyURLUploadInfoList {
	s.CreationTime = &v
	return s
}

func (s *GetURLUploadInfosResponseBodyURLUploadInfoList) SetErrorCode(v string) *GetURLUploadInfosResponseBodyURLUploadInfoList {
	s.ErrorCode = &v
	return s
}

func (s *GetURLUploadInfosResponseBodyURLUploadInfoList) SetErrorMessage(v string) *GetURLUploadInfosResponseBodyURLUploadInfoList {
	s.ErrorMessage = &v
	return s
}

func (s *GetURLUploadInfosResponseBodyURLUploadInfoList) SetFileSize(v string) *GetURLUploadInfosResponseBodyURLUploadInfoList {
	s.FileSize = &v
	return s
}

func (s *GetURLUploadInfosResponseBodyURLUploadInfoList) SetJobId(v string) *GetURLUploadInfosResponseBodyURLUploadInfoList {
	s.JobId = &v
	return s
}

func (s *GetURLUploadInfosResponseBodyURLUploadInfoList) SetMediaId(v string) *GetURLUploadInfosResponseBodyURLUploadInfoList {
	s.MediaId = &v
	return s
}

func (s *GetURLUploadInfosResponseBodyURLUploadInfoList) SetStatus(v string) *GetURLUploadInfosResponseBodyURLUploadInfoList {
	s.Status = &v
	return s
}

func (s *GetURLUploadInfosResponseBodyURLUploadInfoList) SetUploadURL(v string) *GetURLUploadInfosResponseBodyURLUploadInfoList {
	s.UploadURL = &v
	return s
}

func (s *GetURLUploadInfosResponseBodyURLUploadInfoList) SetUserData(v string) *GetURLUploadInfosResponseBodyURLUploadInfoList {
	s.UserData = &v
	return s
}

type GetURLUploadInfosResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetURLUploadInfosResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetURLUploadInfosResponse) String() string {
	return tea.Prettify(s)
}

func (s GetURLUploadInfosResponse) GoString() string {
	return s.String()
}

func (s *GetURLUploadInfosResponse) SetHeaders(v map[string]*string) *GetURLUploadInfosResponse {
	s.Headers = v
	return s
}

func (s *GetURLUploadInfosResponse) SetBody(v *GetURLUploadInfosResponseBody) *GetURLUploadInfosResponse {
	s.Body = v
	return s
}

type GetUploadDetailsRequest struct {
	MediaIds  *string `json:"MediaIds,omitempty" xml:"MediaIds,omitempty"`
	MediaType *string `json:"MediaType,omitempty" xml:"MediaType,omitempty"`
}

func (s GetUploadDetailsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetUploadDetailsRequest) GoString() string {
	return s.String()
}

func (s *GetUploadDetailsRequest) SetMediaIds(v string) *GetUploadDetailsRequest {
	s.MediaIds = &v
	return s
}

func (s *GetUploadDetailsRequest) SetMediaType(v string) *GetUploadDetailsRequest {
	s.MediaType = &v
	return s
}

type GetUploadDetailsResponseBody struct {
	ForbiddenMediaIds []*string                                    `json:"ForbiddenMediaIds,omitempty" xml:"ForbiddenMediaIds,omitempty" type:"Repeated"`
	NonExistMediaIds  []*string                                    `json:"NonExistMediaIds,omitempty" xml:"NonExistMediaIds,omitempty" type:"Repeated"`
	RequestId         *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	UploadDetails     []*GetUploadDetailsResponseBodyUploadDetails `json:"UploadDetails,omitempty" xml:"UploadDetails,omitempty" type:"Repeated"`
}

func (s GetUploadDetailsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetUploadDetailsResponseBody) GoString() string {
	return s.String()
}

func (s *GetUploadDetailsResponseBody) SetForbiddenMediaIds(v []*string) *GetUploadDetailsResponseBody {
	s.ForbiddenMediaIds = v
	return s
}

func (s *GetUploadDetailsResponseBody) SetNonExistMediaIds(v []*string) *GetUploadDetailsResponseBody {
	s.NonExistMediaIds = v
	return s
}

func (s *GetUploadDetailsResponseBody) SetRequestId(v string) *GetUploadDetailsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetUploadDetailsResponseBody) SetUploadDetails(v []*GetUploadDetailsResponseBodyUploadDetails) *GetUploadDetailsResponseBody {
	s.UploadDetails = v
	return s
}

type GetUploadDetailsResponseBodyUploadDetails struct {
	CompletionTime   *string  `json:"CompletionTime,omitempty" xml:"CompletionTime,omitempty"`
	CreationTime     *string  `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	DeviceModel      *string  `json:"DeviceModel,omitempty" xml:"DeviceModel,omitempty"`
	FileSize         *int64   `json:"FileSize,omitempty" xml:"FileSize,omitempty"`
	MediaId          *string  `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	ModificationTime *string  `json:"ModificationTime,omitempty" xml:"ModificationTime,omitempty"`
	Status           *string  `json:"Status,omitempty" xml:"Status,omitempty"`
	Title            *string  `json:"Title,omitempty" xml:"Title,omitempty"`
	UploadIP         *string  `json:"UploadIP,omitempty" xml:"UploadIP,omitempty"`
	UploadRatio      *float32 `json:"UploadRatio,omitempty" xml:"UploadRatio,omitempty"`
	UploadSize       *int64   `json:"UploadSize,omitempty" xml:"UploadSize,omitempty"`
	UploadSource     *string  `json:"UploadSource,omitempty" xml:"UploadSource,omitempty"`
	UploadStatus     *string  `json:"UploadStatus,omitempty" xml:"UploadStatus,omitempty"`
}

func (s GetUploadDetailsResponseBodyUploadDetails) String() string {
	return tea.Prettify(s)
}

func (s GetUploadDetailsResponseBodyUploadDetails) GoString() string {
	return s.String()
}

func (s *GetUploadDetailsResponseBodyUploadDetails) SetCompletionTime(v string) *GetUploadDetailsResponseBodyUploadDetails {
	s.CompletionTime = &v
	return s
}

func (s *GetUploadDetailsResponseBodyUploadDetails) SetCreationTime(v string) *GetUploadDetailsResponseBodyUploadDetails {
	s.CreationTime = &v
	return s
}

func (s *GetUploadDetailsResponseBodyUploadDetails) SetDeviceModel(v string) *GetUploadDetailsResponseBodyUploadDetails {
	s.DeviceModel = &v
	return s
}

func (s *GetUploadDetailsResponseBodyUploadDetails) SetFileSize(v int64) *GetUploadDetailsResponseBodyUploadDetails {
	s.FileSize = &v
	return s
}

func (s *GetUploadDetailsResponseBodyUploadDetails) SetMediaId(v string) *GetUploadDetailsResponseBodyUploadDetails {
	s.MediaId = &v
	return s
}

func (s *GetUploadDetailsResponseBodyUploadDetails) SetModificationTime(v string) *GetUploadDetailsResponseBodyUploadDetails {
	s.ModificationTime = &v
	return s
}

func (s *GetUploadDetailsResponseBodyUploadDetails) SetStatus(v string) *GetUploadDetailsResponseBodyUploadDetails {
	s.Status = &v
	return s
}

func (s *GetUploadDetailsResponseBodyUploadDetails) SetTitle(v string) *GetUploadDetailsResponseBodyUploadDetails {
	s.Title = &v
	return s
}

func (s *GetUploadDetailsResponseBodyUploadDetails) SetUploadIP(v string) *GetUploadDetailsResponseBodyUploadDetails {
	s.UploadIP = &v
	return s
}

func (s *GetUploadDetailsResponseBodyUploadDetails) SetUploadRatio(v float32) *GetUploadDetailsResponseBodyUploadDetails {
	s.UploadRatio = &v
	return s
}

func (s *GetUploadDetailsResponseBodyUploadDetails) SetUploadSize(v int64) *GetUploadDetailsResponseBodyUploadDetails {
	s.UploadSize = &v
	return s
}

func (s *GetUploadDetailsResponseBodyUploadDetails) SetUploadSource(v string) *GetUploadDetailsResponseBodyUploadDetails {
	s.UploadSource = &v
	return s
}

func (s *GetUploadDetailsResponseBodyUploadDetails) SetUploadStatus(v string) *GetUploadDetailsResponseBodyUploadDetails {
	s.UploadStatus = &v
	return s
}

type GetUploadDetailsResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetUploadDetailsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetUploadDetailsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetUploadDetailsResponse) GoString() string {
	return s.String()
}

func (s *GetUploadDetailsResponse) SetHeaders(v map[string]*string) *GetUploadDetailsResponse {
	s.Headers = v
	return s
}

func (s *GetUploadDetailsResponse) SetBody(v *GetUploadDetailsResponseBody) *GetUploadDetailsResponse {
	s.Body = v
	return s
}

type GetVideoInfoRequest struct {
	VideoId *string `json:"VideoId,omitempty" xml:"VideoId,omitempty"`
}

func (s GetVideoInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s GetVideoInfoRequest) GoString() string {
	return s.String()
}

func (s *GetVideoInfoRequest) SetVideoId(v string) *GetVideoInfoRequest {
	s.VideoId = &v
	return s
}

type GetVideoInfoResponseBody struct {
	RequestId *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Video     *GetVideoInfoResponseBodyVideo `json:"Video,omitempty" xml:"Video,omitempty" type:"Struct"`
}

func (s GetVideoInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetVideoInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetVideoInfoResponseBody) SetRequestId(v string) *GetVideoInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetVideoInfoResponseBody) SetVideo(v *GetVideoInfoResponseBodyVideo) *GetVideoInfoResponseBody {
	s.Video = v
	return s
}

type GetVideoInfoResponseBodyVideo struct {
	AppId            *string                                 `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AuditStatus      *string                                 `json:"AuditStatus,omitempty" xml:"AuditStatus,omitempty"`
	CateId           *int64                                  `json:"CateId,omitempty" xml:"CateId,omitempty"`
	CateName         *string                                 `json:"CateName,omitempty" xml:"CateName,omitempty"`
	CoverURL         *string                                 `json:"CoverURL,omitempty" xml:"CoverURL,omitempty"`
	CreationTime     *string                                 `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	CustomMediaInfo  *string                                 `json:"CustomMediaInfo,omitempty" xml:"CustomMediaInfo,omitempty"`
	Description      *string                                 `json:"Description,omitempty" xml:"Description,omitempty"`
	Duration         *float32                                `json:"Duration,omitempty" xml:"Duration,omitempty"`
	ModificationTime *string                                 `json:"ModificationTime,omitempty" xml:"ModificationTime,omitempty"`
	RegionId         *string                                 `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Size             *int64                                  `json:"Size,omitempty" xml:"Size,omitempty"`
	Snapshots        *GetVideoInfoResponseBodyVideoSnapshots `json:"Snapshots,omitempty" xml:"Snapshots,omitempty" type:"Struct"`
	Status           *string                                 `json:"Status,omitempty" xml:"Status,omitempty"`
	StorageLocation  *string                                 `json:"StorageLocation,omitempty" xml:"StorageLocation,omitempty"`
	Tags             *string                                 `json:"Tags,omitempty" xml:"Tags,omitempty"`
	TemplateGroupId  *string                                 `json:"TemplateGroupId,omitempty" xml:"TemplateGroupId,omitempty"`
	Title            *string                                 `json:"Title,omitempty" xml:"Title,omitempty"`
	VideoId          *string                                 `json:"VideoId,omitempty" xml:"VideoId,omitempty"`
}

func (s GetVideoInfoResponseBodyVideo) String() string {
	return tea.Prettify(s)
}

func (s GetVideoInfoResponseBodyVideo) GoString() string {
	return s.String()
}

func (s *GetVideoInfoResponseBodyVideo) SetAppId(v string) *GetVideoInfoResponseBodyVideo {
	s.AppId = &v
	return s
}

func (s *GetVideoInfoResponseBodyVideo) SetAuditStatus(v string) *GetVideoInfoResponseBodyVideo {
	s.AuditStatus = &v
	return s
}

func (s *GetVideoInfoResponseBodyVideo) SetCateId(v int64) *GetVideoInfoResponseBodyVideo {
	s.CateId = &v
	return s
}

func (s *GetVideoInfoResponseBodyVideo) SetCateName(v string) *GetVideoInfoResponseBodyVideo {
	s.CateName = &v
	return s
}

func (s *GetVideoInfoResponseBodyVideo) SetCoverURL(v string) *GetVideoInfoResponseBodyVideo {
	s.CoverURL = &v
	return s
}

func (s *GetVideoInfoResponseBodyVideo) SetCreationTime(v string) *GetVideoInfoResponseBodyVideo {
	s.CreationTime = &v
	return s
}

func (s *GetVideoInfoResponseBodyVideo) SetCustomMediaInfo(v string) *GetVideoInfoResponseBodyVideo {
	s.CustomMediaInfo = &v
	return s
}

func (s *GetVideoInfoResponseBodyVideo) SetDescription(v string) *GetVideoInfoResponseBodyVideo {
	s.Description = &v
	return s
}

func (s *GetVideoInfoResponseBodyVideo) SetDuration(v float32) *GetVideoInfoResponseBodyVideo {
	s.Duration = &v
	return s
}

func (s *GetVideoInfoResponseBodyVideo) SetModificationTime(v string) *GetVideoInfoResponseBodyVideo {
	s.ModificationTime = &v
	return s
}

func (s *GetVideoInfoResponseBodyVideo) SetRegionId(v string) *GetVideoInfoResponseBodyVideo {
	s.RegionId = &v
	return s
}

func (s *GetVideoInfoResponseBodyVideo) SetSize(v int64) *GetVideoInfoResponseBodyVideo {
	s.Size = &v
	return s
}

func (s *GetVideoInfoResponseBodyVideo) SetSnapshots(v *GetVideoInfoResponseBodyVideoSnapshots) *GetVideoInfoResponseBodyVideo {
	s.Snapshots = v
	return s
}

func (s *GetVideoInfoResponseBodyVideo) SetStatus(v string) *GetVideoInfoResponseBodyVideo {
	s.Status = &v
	return s
}

func (s *GetVideoInfoResponseBodyVideo) SetStorageLocation(v string) *GetVideoInfoResponseBodyVideo {
	s.StorageLocation = &v
	return s
}

func (s *GetVideoInfoResponseBodyVideo) SetTags(v string) *GetVideoInfoResponseBodyVideo {
	s.Tags = &v
	return s
}

func (s *GetVideoInfoResponseBodyVideo) SetTemplateGroupId(v string) *GetVideoInfoResponseBodyVideo {
	s.TemplateGroupId = &v
	return s
}

func (s *GetVideoInfoResponseBodyVideo) SetTitle(v string) *GetVideoInfoResponseBodyVideo {
	s.Title = &v
	return s
}

func (s *GetVideoInfoResponseBodyVideo) SetVideoId(v string) *GetVideoInfoResponseBodyVideo {
	s.VideoId = &v
	return s
}

type GetVideoInfoResponseBodyVideoSnapshots struct {
	Snapshot []*string `json:"Snapshot,omitempty" xml:"Snapshot,omitempty" type:"Repeated"`
}

func (s GetVideoInfoResponseBodyVideoSnapshots) String() string {
	return tea.Prettify(s)
}

func (s GetVideoInfoResponseBodyVideoSnapshots) GoString() string {
	return s.String()
}

func (s *GetVideoInfoResponseBodyVideoSnapshots) SetSnapshot(v []*string) *GetVideoInfoResponseBodyVideoSnapshots {
	s.Snapshot = v
	return s
}

type GetVideoInfoResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetVideoInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetVideoInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetVideoInfoResponse) GoString() string {
	return s.String()
}

func (s *GetVideoInfoResponse) SetHeaders(v map[string]*string) *GetVideoInfoResponse {
	s.Headers = v
	return s
}

func (s *GetVideoInfoResponse) SetBody(v *GetVideoInfoResponseBody) *GetVideoInfoResponse {
	s.Body = v
	return s
}

type GetVideoInfosRequest struct {
	VideoIds *string `json:"VideoIds,omitempty" xml:"VideoIds,omitempty"`
}

func (s GetVideoInfosRequest) String() string {
	return tea.Prettify(s)
}

func (s GetVideoInfosRequest) GoString() string {
	return s.String()
}

func (s *GetVideoInfosRequest) SetVideoIds(v string) *GetVideoInfosRequest {
	s.VideoIds = &v
	return s
}

type GetVideoInfosResponseBody struct {
	NonExistVideoIds []*string                             `json:"NonExistVideoIds,omitempty" xml:"NonExistVideoIds,omitempty" type:"Repeated"`
	RequestId        *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	VideoList        []*GetVideoInfosResponseBodyVideoList `json:"VideoList,omitempty" xml:"VideoList,omitempty" type:"Repeated"`
}

func (s GetVideoInfosResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetVideoInfosResponseBody) GoString() string {
	return s.String()
}

func (s *GetVideoInfosResponseBody) SetNonExistVideoIds(v []*string) *GetVideoInfosResponseBody {
	s.NonExistVideoIds = v
	return s
}

func (s *GetVideoInfosResponseBody) SetRequestId(v string) *GetVideoInfosResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetVideoInfosResponseBody) SetVideoList(v []*GetVideoInfosResponseBodyVideoList) *GetVideoInfosResponseBody {
	s.VideoList = v
	return s
}

type GetVideoInfosResponseBodyVideoList struct {
	AppId            *string   `json:"AppId,omitempty" xml:"AppId,omitempty"`
	CateId           *int64    `json:"CateId,omitempty" xml:"CateId,omitempty"`
	CateName         *string   `json:"CateName,omitempty" xml:"CateName,omitempty"`
	CoverURL         *string   `json:"CoverURL,omitempty" xml:"CoverURL,omitempty"`
	CreationTime     *string   `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	Description      *string   `json:"Description,omitempty" xml:"Description,omitempty"`
	Duration         *float32  `json:"Duration,omitempty" xml:"Duration,omitempty"`
	ModificationTime *string   `json:"ModificationTime,omitempty" xml:"ModificationTime,omitempty"`
	Size             *int64    `json:"Size,omitempty" xml:"Size,omitempty"`
	Snapshots        []*string `json:"Snapshots,omitempty" xml:"Snapshots,omitempty" type:"Repeated"`
	Status           *string   `json:"Status,omitempty" xml:"Status,omitempty"`
	StorageLocation  *string   `json:"StorageLocation,omitempty" xml:"StorageLocation,omitempty"`
	Tags             *string   `json:"Tags,omitempty" xml:"Tags,omitempty"`
	TemplateGroupId  *string   `json:"TemplateGroupId,omitempty" xml:"TemplateGroupId,omitempty"`
	Title            *string   `json:"Title,omitempty" xml:"Title,omitempty"`
	VideoId          *string   `json:"VideoId,omitempty" xml:"VideoId,omitempty"`
}

func (s GetVideoInfosResponseBodyVideoList) String() string {
	return tea.Prettify(s)
}

func (s GetVideoInfosResponseBodyVideoList) GoString() string {
	return s.String()
}

func (s *GetVideoInfosResponseBodyVideoList) SetAppId(v string) *GetVideoInfosResponseBodyVideoList {
	s.AppId = &v
	return s
}

func (s *GetVideoInfosResponseBodyVideoList) SetCateId(v int64) *GetVideoInfosResponseBodyVideoList {
	s.CateId = &v
	return s
}

func (s *GetVideoInfosResponseBodyVideoList) SetCateName(v string) *GetVideoInfosResponseBodyVideoList {
	s.CateName = &v
	return s
}

func (s *GetVideoInfosResponseBodyVideoList) SetCoverURL(v string) *GetVideoInfosResponseBodyVideoList {
	s.CoverURL = &v
	return s
}

func (s *GetVideoInfosResponseBodyVideoList) SetCreationTime(v string) *GetVideoInfosResponseBodyVideoList {
	s.CreationTime = &v
	return s
}

func (s *GetVideoInfosResponseBodyVideoList) SetDescription(v string) *GetVideoInfosResponseBodyVideoList {
	s.Description = &v
	return s
}

func (s *GetVideoInfosResponseBodyVideoList) SetDuration(v float32) *GetVideoInfosResponseBodyVideoList {
	s.Duration = &v
	return s
}

func (s *GetVideoInfosResponseBodyVideoList) SetModificationTime(v string) *GetVideoInfosResponseBodyVideoList {
	s.ModificationTime = &v
	return s
}

func (s *GetVideoInfosResponseBodyVideoList) SetSize(v int64) *GetVideoInfosResponseBodyVideoList {
	s.Size = &v
	return s
}

func (s *GetVideoInfosResponseBodyVideoList) SetSnapshots(v []*string) *GetVideoInfosResponseBodyVideoList {
	s.Snapshots = v
	return s
}

func (s *GetVideoInfosResponseBodyVideoList) SetStatus(v string) *GetVideoInfosResponseBodyVideoList {
	s.Status = &v
	return s
}

func (s *GetVideoInfosResponseBodyVideoList) SetStorageLocation(v string) *GetVideoInfosResponseBodyVideoList {
	s.StorageLocation = &v
	return s
}

func (s *GetVideoInfosResponseBodyVideoList) SetTags(v string) *GetVideoInfosResponseBodyVideoList {
	s.Tags = &v
	return s
}

func (s *GetVideoInfosResponseBodyVideoList) SetTemplateGroupId(v string) *GetVideoInfosResponseBodyVideoList {
	s.TemplateGroupId = &v
	return s
}

func (s *GetVideoInfosResponseBodyVideoList) SetTitle(v string) *GetVideoInfosResponseBodyVideoList {
	s.Title = &v
	return s
}

func (s *GetVideoInfosResponseBodyVideoList) SetVideoId(v string) *GetVideoInfosResponseBodyVideoList {
	s.VideoId = &v
	return s
}

type GetVideoInfosResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetVideoInfosResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetVideoInfosResponse) String() string {
	return tea.Prettify(s)
}

func (s GetVideoInfosResponse) GoString() string {
	return s.String()
}

func (s *GetVideoInfosResponse) SetHeaders(v map[string]*string) *GetVideoInfosResponse {
	s.Headers = v
	return s
}

func (s *GetVideoInfosResponse) SetBody(v *GetVideoInfosResponseBody) *GetVideoInfosResponse {
	s.Body = v
	return s
}

type GetVideoListRequest struct {
	CateId          *int64  `json:"CateId,omitempty" xml:"CateId,omitempty"`
	EndTime         *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	PageNo          *int32  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize        *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SortBy          *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	StartTime       *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Status          *string `json:"Status,omitempty" xml:"Status,omitempty"`
	StorageLocation *string `json:"StorageLocation,omitempty" xml:"StorageLocation,omitempty"`
}

func (s GetVideoListRequest) String() string {
	return tea.Prettify(s)
}

func (s GetVideoListRequest) GoString() string {
	return s.String()
}

func (s *GetVideoListRequest) SetCateId(v int64) *GetVideoListRequest {
	s.CateId = &v
	return s
}

func (s *GetVideoListRequest) SetEndTime(v string) *GetVideoListRequest {
	s.EndTime = &v
	return s
}

func (s *GetVideoListRequest) SetPageNo(v int32) *GetVideoListRequest {
	s.PageNo = &v
	return s
}

func (s *GetVideoListRequest) SetPageSize(v int32) *GetVideoListRequest {
	s.PageSize = &v
	return s
}

func (s *GetVideoListRequest) SetSortBy(v string) *GetVideoListRequest {
	s.SortBy = &v
	return s
}

func (s *GetVideoListRequest) SetStartTime(v string) *GetVideoListRequest {
	s.StartTime = &v
	return s
}

func (s *GetVideoListRequest) SetStatus(v string) *GetVideoListRequest {
	s.Status = &v
	return s
}

func (s *GetVideoListRequest) SetStorageLocation(v string) *GetVideoListRequest {
	s.StorageLocation = &v
	return s
}

type GetVideoListResponseBody struct {
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Total     *int32                             `json:"Total,omitempty" xml:"Total,omitempty"`
	VideoList *GetVideoListResponseBodyVideoList `json:"VideoList,omitempty" xml:"VideoList,omitempty" type:"Struct"`
}

func (s GetVideoListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetVideoListResponseBody) GoString() string {
	return s.String()
}

func (s *GetVideoListResponseBody) SetRequestId(v string) *GetVideoListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetVideoListResponseBody) SetTotal(v int32) *GetVideoListResponseBody {
	s.Total = &v
	return s
}

func (s *GetVideoListResponseBody) SetVideoList(v *GetVideoListResponseBodyVideoList) *GetVideoListResponseBody {
	s.VideoList = v
	return s
}

type GetVideoListResponseBodyVideoList struct {
	Video []*GetVideoListResponseBodyVideoListVideo `json:"Video,omitempty" xml:"Video,omitempty" type:"Repeated"`
}

func (s GetVideoListResponseBodyVideoList) String() string {
	return tea.Prettify(s)
}

func (s GetVideoListResponseBodyVideoList) GoString() string {
	return s.String()
}

func (s *GetVideoListResponseBodyVideoList) SetVideo(v []*GetVideoListResponseBodyVideoListVideo) *GetVideoListResponseBodyVideoList {
	s.Video = v
	return s
}

type GetVideoListResponseBodyVideoListVideo struct {
	AppId            *string                                          `json:"AppId,omitempty" xml:"AppId,omitempty"`
	CateId           *int64                                           `json:"CateId,omitempty" xml:"CateId,omitempty"`
	CateName         *string                                          `json:"CateName,omitempty" xml:"CateName,omitempty"`
	CoverURL         *string                                          `json:"CoverURL,omitempty" xml:"CoverURL,omitempty"`
	CreationTime     *string                                          `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	Description      *string                                          `json:"Description,omitempty" xml:"Description,omitempty"`
	Duration         *float32                                         `json:"Duration,omitempty" xml:"Duration,omitempty"`
	ModificationTime *string                                          `json:"ModificationTime,omitempty" xml:"ModificationTime,omitempty"`
	Size             *int64                                           `json:"Size,omitempty" xml:"Size,omitempty"`
	Snapshots        *GetVideoListResponseBodyVideoListVideoSnapshots `json:"Snapshots,omitempty" xml:"Snapshots,omitempty" type:"Struct"`
	Status           *string                                          `json:"Status,omitempty" xml:"Status,omitempty"`
	StorageLocation  *string                                          `json:"StorageLocation,omitempty" xml:"StorageLocation,omitempty"`
	Tags             *string                                          `json:"Tags,omitempty" xml:"Tags,omitempty"`
	Title            *string                                          `json:"Title,omitempty" xml:"Title,omitempty"`
	VideoId          *string                                          `json:"VideoId,omitempty" xml:"VideoId,omitempty"`
}

func (s GetVideoListResponseBodyVideoListVideo) String() string {
	return tea.Prettify(s)
}

func (s GetVideoListResponseBodyVideoListVideo) GoString() string {
	return s.String()
}

func (s *GetVideoListResponseBodyVideoListVideo) SetAppId(v string) *GetVideoListResponseBodyVideoListVideo {
	s.AppId = &v
	return s
}

func (s *GetVideoListResponseBodyVideoListVideo) SetCateId(v int64) *GetVideoListResponseBodyVideoListVideo {
	s.CateId = &v
	return s
}

func (s *GetVideoListResponseBodyVideoListVideo) SetCateName(v string) *GetVideoListResponseBodyVideoListVideo {
	s.CateName = &v
	return s
}

func (s *GetVideoListResponseBodyVideoListVideo) SetCoverURL(v string) *GetVideoListResponseBodyVideoListVideo {
	s.CoverURL = &v
	return s
}

func (s *GetVideoListResponseBodyVideoListVideo) SetCreationTime(v string) *GetVideoListResponseBodyVideoListVideo {
	s.CreationTime = &v
	return s
}

func (s *GetVideoListResponseBodyVideoListVideo) SetDescription(v string) *GetVideoListResponseBodyVideoListVideo {
	s.Description = &v
	return s
}

func (s *GetVideoListResponseBodyVideoListVideo) SetDuration(v float32) *GetVideoListResponseBodyVideoListVideo {
	s.Duration = &v
	return s
}

func (s *GetVideoListResponseBodyVideoListVideo) SetModificationTime(v string) *GetVideoListResponseBodyVideoListVideo {
	s.ModificationTime = &v
	return s
}

func (s *GetVideoListResponseBodyVideoListVideo) SetSize(v int64) *GetVideoListResponseBodyVideoListVideo {
	s.Size = &v
	return s
}

func (s *GetVideoListResponseBodyVideoListVideo) SetSnapshots(v *GetVideoListResponseBodyVideoListVideoSnapshots) *GetVideoListResponseBodyVideoListVideo {
	s.Snapshots = v
	return s
}

func (s *GetVideoListResponseBodyVideoListVideo) SetStatus(v string) *GetVideoListResponseBodyVideoListVideo {
	s.Status = &v
	return s
}

func (s *GetVideoListResponseBodyVideoListVideo) SetStorageLocation(v string) *GetVideoListResponseBodyVideoListVideo {
	s.StorageLocation = &v
	return s
}

func (s *GetVideoListResponseBodyVideoListVideo) SetTags(v string) *GetVideoListResponseBodyVideoListVideo {
	s.Tags = &v
	return s
}

func (s *GetVideoListResponseBodyVideoListVideo) SetTitle(v string) *GetVideoListResponseBodyVideoListVideo {
	s.Title = &v
	return s
}

func (s *GetVideoListResponseBodyVideoListVideo) SetVideoId(v string) *GetVideoListResponseBodyVideoListVideo {
	s.VideoId = &v
	return s
}

type GetVideoListResponseBodyVideoListVideoSnapshots struct {
	Snapshot []*string `json:"Snapshot,omitempty" xml:"Snapshot,omitempty" type:"Repeated"`
}

func (s GetVideoListResponseBodyVideoListVideoSnapshots) String() string {
	return tea.Prettify(s)
}

func (s GetVideoListResponseBodyVideoListVideoSnapshots) GoString() string {
	return s.String()
}

func (s *GetVideoListResponseBodyVideoListVideoSnapshots) SetSnapshot(v []*string) *GetVideoListResponseBodyVideoListVideoSnapshots {
	s.Snapshot = v
	return s
}

type GetVideoListResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetVideoListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetVideoListResponse) String() string {
	return tea.Prettify(s)
}

func (s GetVideoListResponse) GoString() string {
	return s.String()
}

func (s *GetVideoListResponse) SetHeaders(v map[string]*string) *GetVideoListResponse {
	s.Headers = v
	return s
}

func (s *GetVideoListResponse) SetBody(v *GetVideoListResponseBody) *GetVideoListResponse {
	s.Body = v
	return s
}

type GetVideoPlayAuthRequest struct {
	AuthInfoTimeout *int64  `json:"AuthInfoTimeout,omitempty" xml:"AuthInfoTimeout,omitempty"`
	VideoId         *string `json:"VideoId,omitempty" xml:"VideoId,omitempty"`
}

func (s GetVideoPlayAuthRequest) String() string {
	return tea.Prettify(s)
}

func (s GetVideoPlayAuthRequest) GoString() string {
	return s.String()
}

func (s *GetVideoPlayAuthRequest) SetAuthInfoTimeout(v int64) *GetVideoPlayAuthRequest {
	s.AuthInfoTimeout = &v
	return s
}

func (s *GetVideoPlayAuthRequest) SetVideoId(v string) *GetVideoPlayAuthRequest {
	s.VideoId = &v
	return s
}

type GetVideoPlayAuthResponseBody struct {
	PlayAuth  *string                                `json:"PlayAuth,omitempty" xml:"PlayAuth,omitempty"`
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	VideoMeta *GetVideoPlayAuthResponseBodyVideoMeta `json:"VideoMeta,omitempty" xml:"VideoMeta,omitempty" type:"Struct"`
}

func (s GetVideoPlayAuthResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetVideoPlayAuthResponseBody) GoString() string {
	return s.String()
}

func (s *GetVideoPlayAuthResponseBody) SetPlayAuth(v string) *GetVideoPlayAuthResponseBody {
	s.PlayAuth = &v
	return s
}

func (s *GetVideoPlayAuthResponseBody) SetRequestId(v string) *GetVideoPlayAuthResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetVideoPlayAuthResponseBody) SetVideoMeta(v *GetVideoPlayAuthResponseBodyVideoMeta) *GetVideoPlayAuthResponseBody {
	s.VideoMeta = v
	return s
}

type GetVideoPlayAuthResponseBodyVideoMeta struct {
	CoverURL *string  `json:"CoverURL,omitempty" xml:"CoverURL,omitempty"`
	Duration *float32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	Status   *string  `json:"Status,omitempty" xml:"Status,omitempty"`
	Title    *string  `json:"Title,omitempty" xml:"Title,omitempty"`
	VideoId  *string  `json:"VideoId,omitempty" xml:"VideoId,omitempty"`
}

func (s GetVideoPlayAuthResponseBodyVideoMeta) String() string {
	return tea.Prettify(s)
}

func (s GetVideoPlayAuthResponseBodyVideoMeta) GoString() string {
	return s.String()
}

func (s *GetVideoPlayAuthResponseBodyVideoMeta) SetCoverURL(v string) *GetVideoPlayAuthResponseBodyVideoMeta {
	s.CoverURL = &v
	return s
}

func (s *GetVideoPlayAuthResponseBodyVideoMeta) SetDuration(v float32) *GetVideoPlayAuthResponseBodyVideoMeta {
	s.Duration = &v
	return s
}

func (s *GetVideoPlayAuthResponseBodyVideoMeta) SetStatus(v string) *GetVideoPlayAuthResponseBodyVideoMeta {
	s.Status = &v
	return s
}

func (s *GetVideoPlayAuthResponseBodyVideoMeta) SetTitle(v string) *GetVideoPlayAuthResponseBodyVideoMeta {
	s.Title = &v
	return s
}

func (s *GetVideoPlayAuthResponseBodyVideoMeta) SetVideoId(v string) *GetVideoPlayAuthResponseBodyVideoMeta {
	s.VideoId = &v
	return s
}

type GetVideoPlayAuthResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetVideoPlayAuthResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetVideoPlayAuthResponse) String() string {
	return tea.Prettify(s)
}

func (s GetVideoPlayAuthResponse) GoString() string {
	return s.String()
}

func (s *GetVideoPlayAuthResponse) SetHeaders(v map[string]*string) *GetVideoPlayAuthResponse {
	s.Headers = v
	return s
}

func (s *GetVideoPlayAuthResponse) SetBody(v *GetVideoPlayAuthResponseBody) *GetVideoPlayAuthResponse {
	s.Body = v
	return s
}

type GetVodTemplateRequest struct {
	VodTemplateId *string `json:"VodTemplateId,omitempty" xml:"VodTemplateId,omitempty"`
}

func (s GetVodTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s GetVodTemplateRequest) GoString() string {
	return s.String()
}

func (s *GetVodTemplateRequest) SetVodTemplateId(v string) *GetVodTemplateRequest {
	s.VodTemplateId = &v
	return s
}

type GetVodTemplateResponseBody struct {
	RequestId       *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	VodTemplateInfo *GetVodTemplateResponseBodyVodTemplateInfo `json:"VodTemplateInfo,omitempty" xml:"VodTemplateInfo,omitempty" type:"Struct"`
}

func (s GetVodTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetVodTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *GetVodTemplateResponseBody) SetRequestId(v string) *GetVodTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetVodTemplateResponseBody) SetVodTemplateInfo(v *GetVodTemplateResponseBodyVodTemplateInfo) *GetVodTemplateResponseBody {
	s.VodTemplateInfo = v
	return s
}

type GetVodTemplateResponseBodyVodTemplateInfo struct {
	CreationTime   *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	IsDefault      *string `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	ModifyTime     *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	Name           *string `json:"Name,omitempty" xml:"Name,omitempty"`
	TemplateConfig *string `json:"TemplateConfig,omitempty" xml:"TemplateConfig,omitempty"`
	TemplateType   *string `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
	VodTemplateId  *string `json:"VodTemplateId,omitempty" xml:"VodTemplateId,omitempty"`
}

func (s GetVodTemplateResponseBodyVodTemplateInfo) String() string {
	return tea.Prettify(s)
}

func (s GetVodTemplateResponseBodyVodTemplateInfo) GoString() string {
	return s.String()
}

func (s *GetVodTemplateResponseBodyVodTemplateInfo) SetCreationTime(v string) *GetVodTemplateResponseBodyVodTemplateInfo {
	s.CreationTime = &v
	return s
}

func (s *GetVodTemplateResponseBodyVodTemplateInfo) SetIsDefault(v string) *GetVodTemplateResponseBodyVodTemplateInfo {
	s.IsDefault = &v
	return s
}

func (s *GetVodTemplateResponseBodyVodTemplateInfo) SetModifyTime(v string) *GetVodTemplateResponseBodyVodTemplateInfo {
	s.ModifyTime = &v
	return s
}

func (s *GetVodTemplateResponseBodyVodTemplateInfo) SetName(v string) *GetVodTemplateResponseBodyVodTemplateInfo {
	s.Name = &v
	return s
}

func (s *GetVodTemplateResponseBodyVodTemplateInfo) SetTemplateConfig(v string) *GetVodTemplateResponseBodyVodTemplateInfo {
	s.TemplateConfig = &v
	return s
}

func (s *GetVodTemplateResponseBodyVodTemplateInfo) SetTemplateType(v string) *GetVodTemplateResponseBodyVodTemplateInfo {
	s.TemplateType = &v
	return s
}

func (s *GetVodTemplateResponseBodyVodTemplateInfo) SetVodTemplateId(v string) *GetVodTemplateResponseBodyVodTemplateInfo {
	s.VodTemplateId = &v
	return s
}

type GetVodTemplateResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetVodTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetVodTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s GetVodTemplateResponse) GoString() string {
	return s.String()
}

func (s *GetVodTemplateResponse) SetHeaders(v map[string]*string) *GetVodTemplateResponse {
	s.Headers = v
	return s
}

func (s *GetVodTemplateResponse) SetBody(v *GetVodTemplateResponseBody) *GetVodTemplateResponse {
	s.Body = v
	return s
}

type GetWatermarkRequest struct {
	WatermarkId *string `json:"WatermarkId,omitempty" xml:"WatermarkId,omitempty"`
}

func (s GetWatermarkRequest) String() string {
	return tea.Prettify(s)
}

func (s GetWatermarkRequest) GoString() string {
	return s.String()
}

func (s *GetWatermarkRequest) SetWatermarkId(v string) *GetWatermarkRequest {
	s.WatermarkId = &v
	return s
}

type GetWatermarkResponseBody struct {
	RequestId     *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	WatermarkInfo *GetWatermarkResponseBodyWatermarkInfo `json:"WatermarkInfo,omitempty" xml:"WatermarkInfo,omitempty" type:"Struct"`
}

func (s GetWatermarkResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetWatermarkResponseBody) GoString() string {
	return s.String()
}

func (s *GetWatermarkResponseBody) SetRequestId(v string) *GetWatermarkResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetWatermarkResponseBody) SetWatermarkInfo(v *GetWatermarkResponseBodyWatermarkInfo) *GetWatermarkResponseBody {
	s.WatermarkInfo = v
	return s
}

type GetWatermarkResponseBodyWatermarkInfo struct {
	AppId           *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	CreationTime    *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	FileUrl         *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	IsDefault       *string `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	Name            *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Type            *string `json:"Type,omitempty" xml:"Type,omitempty"`
	WatermarkConfig *string `json:"WatermarkConfig,omitempty" xml:"WatermarkConfig,omitempty"`
	WatermarkId     *string `json:"WatermarkId,omitempty" xml:"WatermarkId,omitempty"`
}

func (s GetWatermarkResponseBodyWatermarkInfo) String() string {
	return tea.Prettify(s)
}

func (s GetWatermarkResponseBodyWatermarkInfo) GoString() string {
	return s.String()
}

func (s *GetWatermarkResponseBodyWatermarkInfo) SetAppId(v string) *GetWatermarkResponseBodyWatermarkInfo {
	s.AppId = &v
	return s
}

func (s *GetWatermarkResponseBodyWatermarkInfo) SetCreationTime(v string) *GetWatermarkResponseBodyWatermarkInfo {
	s.CreationTime = &v
	return s
}

func (s *GetWatermarkResponseBodyWatermarkInfo) SetFileUrl(v string) *GetWatermarkResponseBodyWatermarkInfo {
	s.FileUrl = &v
	return s
}

func (s *GetWatermarkResponseBodyWatermarkInfo) SetIsDefault(v string) *GetWatermarkResponseBodyWatermarkInfo {
	s.IsDefault = &v
	return s
}

func (s *GetWatermarkResponseBodyWatermarkInfo) SetName(v string) *GetWatermarkResponseBodyWatermarkInfo {
	s.Name = &v
	return s
}

func (s *GetWatermarkResponseBodyWatermarkInfo) SetType(v string) *GetWatermarkResponseBodyWatermarkInfo {
	s.Type = &v
	return s
}

func (s *GetWatermarkResponseBodyWatermarkInfo) SetWatermarkConfig(v string) *GetWatermarkResponseBodyWatermarkInfo {
	s.WatermarkConfig = &v
	return s
}

func (s *GetWatermarkResponseBodyWatermarkInfo) SetWatermarkId(v string) *GetWatermarkResponseBodyWatermarkInfo {
	s.WatermarkId = &v
	return s
}

type GetWatermarkResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetWatermarkResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetWatermarkResponse) String() string {
	return tea.Prettify(s)
}

func (s GetWatermarkResponse) GoString() string {
	return s.String()
}

func (s *GetWatermarkResponse) SetHeaders(v map[string]*string) *GetWatermarkResponse {
	s.Headers = v
	return s
}

func (s *GetWatermarkResponse) SetBody(v *GetWatermarkResponseBody) *GetWatermarkResponse {
	s.Body = v
	return s
}

type ListAIImageInfoRequest struct {
	VideoId *string `json:"VideoId,omitempty" xml:"VideoId,omitempty"`
}

func (s ListAIImageInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAIImageInfoRequest) GoString() string {
	return s.String()
}

func (s *ListAIImageInfoRequest) SetVideoId(v string) *ListAIImageInfoRequest {
	s.VideoId = &v
	return s
}

type ListAIImageInfoResponseBody struct {
	AIImageInfoList []*ListAIImageInfoResponseBodyAIImageInfoList `json:"AIImageInfoList,omitempty" xml:"AIImageInfoList,omitempty" type:"Repeated"`
	RequestId       *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListAIImageInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAIImageInfoResponseBody) GoString() string {
	return s.String()
}

func (s *ListAIImageInfoResponseBody) SetAIImageInfoList(v []*ListAIImageInfoResponseBodyAIImageInfoList) *ListAIImageInfoResponseBody {
	s.AIImageInfoList = v
	return s
}

func (s *ListAIImageInfoResponseBody) SetRequestId(v string) *ListAIImageInfoResponseBody {
	s.RequestId = &v
	return s
}

type ListAIImageInfoResponseBodyAIImageInfoList struct {
	AIImageInfoId *string `json:"AIImageInfoId,omitempty" xml:"AIImageInfoId,omitempty"`
	CreationTime  *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	FileURL       *string `json:"FileURL,omitempty" xml:"FileURL,omitempty"`
	Format        *string `json:"Format,omitempty" xml:"Format,omitempty"`
	JobId         *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	Score         *string `json:"Score,omitempty" xml:"Score,omitempty"`
	Version       *string `json:"Version,omitempty" xml:"Version,omitempty"`
	VideoId       *string `json:"VideoId,omitempty" xml:"VideoId,omitempty"`
}

func (s ListAIImageInfoResponseBodyAIImageInfoList) String() string {
	return tea.Prettify(s)
}

func (s ListAIImageInfoResponseBodyAIImageInfoList) GoString() string {
	return s.String()
}

func (s *ListAIImageInfoResponseBodyAIImageInfoList) SetAIImageInfoId(v string) *ListAIImageInfoResponseBodyAIImageInfoList {
	s.AIImageInfoId = &v
	return s
}

func (s *ListAIImageInfoResponseBodyAIImageInfoList) SetCreationTime(v string) *ListAIImageInfoResponseBodyAIImageInfoList {
	s.CreationTime = &v
	return s
}

func (s *ListAIImageInfoResponseBodyAIImageInfoList) SetFileURL(v string) *ListAIImageInfoResponseBodyAIImageInfoList {
	s.FileURL = &v
	return s
}

func (s *ListAIImageInfoResponseBodyAIImageInfoList) SetFormat(v string) *ListAIImageInfoResponseBodyAIImageInfoList {
	s.Format = &v
	return s
}

func (s *ListAIImageInfoResponseBodyAIImageInfoList) SetJobId(v string) *ListAIImageInfoResponseBodyAIImageInfoList {
	s.JobId = &v
	return s
}

func (s *ListAIImageInfoResponseBodyAIImageInfoList) SetScore(v string) *ListAIImageInfoResponseBodyAIImageInfoList {
	s.Score = &v
	return s
}

func (s *ListAIImageInfoResponseBodyAIImageInfoList) SetVersion(v string) *ListAIImageInfoResponseBodyAIImageInfoList {
	s.Version = &v
	return s
}

func (s *ListAIImageInfoResponseBodyAIImageInfoList) SetVideoId(v string) *ListAIImageInfoResponseBodyAIImageInfoList {
	s.VideoId = &v
	return s
}

type ListAIImageInfoResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListAIImageInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListAIImageInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAIImageInfoResponse) GoString() string {
	return s.String()
}

func (s *ListAIImageInfoResponse) SetHeaders(v map[string]*string) *ListAIImageInfoResponse {
	s.Headers = v
	return s
}

func (s *ListAIImageInfoResponse) SetBody(v *ListAIImageInfoResponseBody) *ListAIImageInfoResponse {
	s.Body = v
	return s
}

type ListAIJobRequest struct {
	JobIds               *string `json:"JobIds,omitempty" xml:"JobIds,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *string `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ListAIJobRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAIJobRequest) GoString() string {
	return s.String()
}

func (s *ListAIJobRequest) SetJobIds(v string) *ListAIJobRequest {
	s.JobIds = &v
	return s
}

func (s *ListAIJobRequest) SetOwnerAccount(v string) *ListAIJobRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ListAIJobRequest) SetOwnerId(v string) *ListAIJobRequest {
	s.OwnerId = &v
	return s
}

func (s *ListAIJobRequest) SetResourceOwnerAccount(v string) *ListAIJobRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListAIJobRequest) SetResourceOwnerId(v string) *ListAIJobRequest {
	s.ResourceOwnerId = &v
	return s
}

type ListAIJobResponseBody struct {
	AIJobList        *ListAIJobResponseBodyAIJobList        `json:"AIJobList,omitempty" xml:"AIJobList,omitempty" type:"Struct"`
	NonExistAIJobIds *ListAIJobResponseBodyNonExistAIJobIds `json:"NonExistAIJobIds,omitempty" xml:"NonExistAIJobIds,omitempty" type:"Struct"`
	RequestId        *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListAIJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAIJobResponseBody) GoString() string {
	return s.String()
}

func (s *ListAIJobResponseBody) SetAIJobList(v *ListAIJobResponseBodyAIJobList) *ListAIJobResponseBody {
	s.AIJobList = v
	return s
}

func (s *ListAIJobResponseBody) SetNonExistAIJobIds(v *ListAIJobResponseBodyNonExistAIJobIds) *ListAIJobResponseBody {
	s.NonExistAIJobIds = v
	return s
}

func (s *ListAIJobResponseBody) SetRequestId(v string) *ListAIJobResponseBody {
	s.RequestId = &v
	return s
}

type ListAIJobResponseBodyAIJobList struct {
	AIJob []*ListAIJobResponseBodyAIJobListAIJob `json:"AIJob,omitempty" xml:"AIJob,omitempty" type:"Repeated"`
}

func (s ListAIJobResponseBodyAIJobList) String() string {
	return tea.Prettify(s)
}

func (s ListAIJobResponseBodyAIJobList) GoString() string {
	return s.String()
}

func (s *ListAIJobResponseBodyAIJobList) SetAIJob(v []*ListAIJobResponseBodyAIJobListAIJob) *ListAIJobResponseBodyAIJobList {
	s.AIJob = v
	return s
}

type ListAIJobResponseBodyAIJobListAIJob struct {
	Code         *string `json:"Code,omitempty" xml:"Code,omitempty"`
	CompleteTime *string `json:"CompleteTime,omitempty" xml:"CompleteTime,omitempty"`
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	Data         *string `json:"Data,omitempty" xml:"Data,omitempty"`
	JobId        *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	MediaId      *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	Message      *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Type         *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListAIJobResponseBodyAIJobListAIJob) String() string {
	return tea.Prettify(s)
}

func (s ListAIJobResponseBodyAIJobListAIJob) GoString() string {
	return s.String()
}

func (s *ListAIJobResponseBodyAIJobListAIJob) SetCode(v string) *ListAIJobResponseBodyAIJobListAIJob {
	s.Code = &v
	return s
}

func (s *ListAIJobResponseBodyAIJobListAIJob) SetCompleteTime(v string) *ListAIJobResponseBodyAIJobListAIJob {
	s.CompleteTime = &v
	return s
}

func (s *ListAIJobResponseBodyAIJobListAIJob) SetCreationTime(v string) *ListAIJobResponseBodyAIJobListAIJob {
	s.CreationTime = &v
	return s
}

func (s *ListAIJobResponseBodyAIJobListAIJob) SetData(v string) *ListAIJobResponseBodyAIJobListAIJob {
	s.Data = &v
	return s
}

func (s *ListAIJobResponseBodyAIJobListAIJob) SetJobId(v string) *ListAIJobResponseBodyAIJobListAIJob {
	s.JobId = &v
	return s
}

func (s *ListAIJobResponseBodyAIJobListAIJob) SetMediaId(v string) *ListAIJobResponseBodyAIJobListAIJob {
	s.MediaId = &v
	return s
}

func (s *ListAIJobResponseBodyAIJobListAIJob) SetMessage(v string) *ListAIJobResponseBodyAIJobListAIJob {
	s.Message = &v
	return s
}

func (s *ListAIJobResponseBodyAIJobListAIJob) SetStatus(v string) *ListAIJobResponseBodyAIJobListAIJob {
	s.Status = &v
	return s
}

func (s *ListAIJobResponseBodyAIJobListAIJob) SetType(v string) *ListAIJobResponseBodyAIJobListAIJob {
	s.Type = &v
	return s
}

type ListAIJobResponseBodyNonExistAIJobIds struct {
	String_ []*string `json:"String,omitempty" xml:"String,omitempty" type:"Repeated"`
}

func (s ListAIJobResponseBodyNonExistAIJobIds) String() string {
	return tea.Prettify(s)
}

func (s ListAIJobResponseBodyNonExistAIJobIds) GoString() string {
	return s.String()
}

func (s *ListAIJobResponseBodyNonExistAIJobIds) SetString_(v []*string) *ListAIJobResponseBodyNonExistAIJobIds {
	s.String_ = v
	return s
}

type ListAIJobResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListAIJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListAIJobResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAIJobResponse) GoString() string {
	return s.String()
}

func (s *ListAIJobResponse) SetHeaders(v map[string]*string) *ListAIJobResponse {
	s.Headers = v
	return s
}

func (s *ListAIJobResponse) SetBody(v *ListAIJobResponseBody) *ListAIJobResponse {
	s.Body = v
	return s
}

type ListAITemplateRequest struct {
	TemplateType *string `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
}

func (s ListAITemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAITemplateRequest) GoString() string {
	return s.String()
}

func (s *ListAITemplateRequest) SetTemplateType(v string) *ListAITemplateRequest {
	s.TemplateType = &v
	return s
}

type ListAITemplateResponseBody struct {
	RequestId        *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TemplateInfoList []*ListAITemplateResponseBodyTemplateInfoList `json:"TemplateInfoList,omitempty" xml:"TemplateInfoList,omitempty" type:"Repeated"`
}

func (s ListAITemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAITemplateResponseBody) GoString() string {
	return s.String()
}

func (s *ListAITemplateResponseBody) SetRequestId(v string) *ListAITemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAITemplateResponseBody) SetTemplateInfoList(v []*ListAITemplateResponseBodyTemplateInfoList) *ListAITemplateResponseBody {
	s.TemplateInfoList = v
	return s
}

type ListAITemplateResponseBodyTemplateInfoList struct {
	CreationTime   *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	IsDefault      *string `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	ModifyTime     *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	Source         *string `json:"Source,omitempty" xml:"Source,omitempty"`
	TemplateConfig *string `json:"TemplateConfig,omitempty" xml:"TemplateConfig,omitempty"`
	TemplateId     *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	TemplateName   *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	TemplateType   *string `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
}

func (s ListAITemplateResponseBodyTemplateInfoList) String() string {
	return tea.Prettify(s)
}

func (s ListAITemplateResponseBodyTemplateInfoList) GoString() string {
	return s.String()
}

func (s *ListAITemplateResponseBodyTemplateInfoList) SetCreationTime(v string) *ListAITemplateResponseBodyTemplateInfoList {
	s.CreationTime = &v
	return s
}

func (s *ListAITemplateResponseBodyTemplateInfoList) SetIsDefault(v string) *ListAITemplateResponseBodyTemplateInfoList {
	s.IsDefault = &v
	return s
}

func (s *ListAITemplateResponseBodyTemplateInfoList) SetModifyTime(v string) *ListAITemplateResponseBodyTemplateInfoList {
	s.ModifyTime = &v
	return s
}

func (s *ListAITemplateResponseBodyTemplateInfoList) SetSource(v string) *ListAITemplateResponseBodyTemplateInfoList {
	s.Source = &v
	return s
}

func (s *ListAITemplateResponseBodyTemplateInfoList) SetTemplateConfig(v string) *ListAITemplateResponseBodyTemplateInfoList {
	s.TemplateConfig = &v
	return s
}

func (s *ListAITemplateResponseBodyTemplateInfoList) SetTemplateId(v string) *ListAITemplateResponseBodyTemplateInfoList {
	s.TemplateId = &v
	return s
}

func (s *ListAITemplateResponseBodyTemplateInfoList) SetTemplateName(v string) *ListAITemplateResponseBodyTemplateInfoList {
	s.TemplateName = &v
	return s
}

func (s *ListAITemplateResponseBodyTemplateInfoList) SetTemplateType(v string) *ListAITemplateResponseBodyTemplateInfoList {
	s.TemplateType = &v
	return s
}

type ListAITemplateResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListAITemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListAITemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAITemplateResponse) GoString() string {
	return s.String()
}

func (s *ListAITemplateResponse) SetHeaders(v map[string]*string) *ListAITemplateResponse {
	s.Headers = v
	return s
}

func (s *ListAITemplateResponse) SetBody(v *ListAITemplateResponseBody) *ListAITemplateResponse {
	s.Body = v
	return s
}

type ListAppInfoRequest struct {
	PageNo   *int32  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Status   *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListAppInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAppInfoRequest) GoString() string {
	return s.String()
}

func (s *ListAppInfoRequest) SetPageNo(v int32) *ListAppInfoRequest {
	s.PageNo = &v
	return s
}

func (s *ListAppInfoRequest) SetPageSize(v int32) *ListAppInfoRequest {
	s.PageSize = &v
	return s
}

func (s *ListAppInfoRequest) SetStatus(v string) *ListAppInfoRequest {
	s.Status = &v
	return s
}

type ListAppInfoResponseBody struct {
	AppInfoList []*ListAppInfoResponseBodyAppInfoList `json:"AppInfoList,omitempty" xml:"AppInfoList,omitempty" type:"Repeated"`
	RequestId   *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Total       *int32                                `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListAppInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAppInfoResponseBody) GoString() string {
	return s.String()
}

func (s *ListAppInfoResponseBody) SetAppInfoList(v []*ListAppInfoResponseBodyAppInfoList) *ListAppInfoResponseBody {
	s.AppInfoList = v
	return s
}

func (s *ListAppInfoResponseBody) SetRequestId(v string) *ListAppInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAppInfoResponseBody) SetTotal(v int32) *ListAppInfoResponseBody {
	s.Total = &v
	return s
}

type ListAppInfoResponseBodyAppInfoList struct {
	AppId            *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AppName          *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	CreationTime     *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	Description      *string `json:"Description,omitempty" xml:"Description,omitempty"`
	ModificationTime *string `json:"ModificationTime,omitempty" xml:"ModificationTime,omitempty"`
	Status           *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Type             *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListAppInfoResponseBodyAppInfoList) String() string {
	return tea.Prettify(s)
}

func (s ListAppInfoResponseBodyAppInfoList) GoString() string {
	return s.String()
}

func (s *ListAppInfoResponseBodyAppInfoList) SetAppId(v string) *ListAppInfoResponseBodyAppInfoList {
	s.AppId = &v
	return s
}

func (s *ListAppInfoResponseBodyAppInfoList) SetAppName(v string) *ListAppInfoResponseBodyAppInfoList {
	s.AppName = &v
	return s
}

func (s *ListAppInfoResponseBodyAppInfoList) SetCreationTime(v string) *ListAppInfoResponseBodyAppInfoList {
	s.CreationTime = &v
	return s
}

func (s *ListAppInfoResponseBodyAppInfoList) SetDescription(v string) *ListAppInfoResponseBodyAppInfoList {
	s.Description = &v
	return s
}

func (s *ListAppInfoResponseBodyAppInfoList) SetModificationTime(v string) *ListAppInfoResponseBodyAppInfoList {
	s.ModificationTime = &v
	return s
}

func (s *ListAppInfoResponseBodyAppInfoList) SetStatus(v string) *ListAppInfoResponseBodyAppInfoList {
	s.Status = &v
	return s
}

func (s *ListAppInfoResponseBodyAppInfoList) SetType(v string) *ListAppInfoResponseBodyAppInfoList {
	s.Type = &v
	return s
}

type ListAppInfoResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListAppInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListAppInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAppInfoResponse) GoString() string {
	return s.String()
}

func (s *ListAppInfoResponse) SetHeaders(v map[string]*string) *ListAppInfoResponse {
	s.Headers = v
	return s
}

func (s *ListAppInfoResponse) SetBody(v *ListAppInfoResponseBody) *ListAppInfoResponse {
	s.Body = v
	return s
}

type ListAppPoliciesForIdentityRequest struct {
	AppId        *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	IdentityName *string `json:"IdentityName,omitempty" xml:"IdentityName,omitempty"`
	IdentityType *string `json:"IdentityType,omitempty" xml:"IdentityType,omitempty"`
}

func (s ListAppPoliciesForIdentityRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAppPoliciesForIdentityRequest) GoString() string {
	return s.String()
}

func (s *ListAppPoliciesForIdentityRequest) SetAppId(v string) *ListAppPoliciesForIdentityRequest {
	s.AppId = &v
	return s
}

func (s *ListAppPoliciesForIdentityRequest) SetIdentityName(v string) *ListAppPoliciesForIdentityRequest {
	s.IdentityName = &v
	return s
}

func (s *ListAppPoliciesForIdentityRequest) SetIdentityType(v string) *ListAppPoliciesForIdentityRequest {
	s.IdentityType = &v
	return s
}

type ListAppPoliciesForIdentityResponseBody struct {
	AppPolicyList []*ListAppPoliciesForIdentityResponseBodyAppPolicyList `json:"AppPolicyList,omitempty" xml:"AppPolicyList,omitempty" type:"Repeated"`
	RequestId     *string                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListAppPoliciesForIdentityResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAppPoliciesForIdentityResponseBody) GoString() string {
	return s.String()
}

func (s *ListAppPoliciesForIdentityResponseBody) SetAppPolicyList(v []*ListAppPoliciesForIdentityResponseBodyAppPolicyList) *ListAppPoliciesForIdentityResponseBody {
	s.AppPolicyList = v
	return s
}

func (s *ListAppPoliciesForIdentityResponseBody) SetRequestId(v string) *ListAppPoliciesForIdentityResponseBody {
	s.RequestId = &v
	return s
}

type ListAppPoliciesForIdentityResponseBodyAppPolicyList struct {
	AppId            *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	CreationTime     *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	Description      *string `json:"Description,omitempty" xml:"Description,omitempty"`
	ModificationTime *string `json:"ModificationTime,omitempty" xml:"ModificationTime,omitempty"`
	PolicyName       *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	PolicyType       *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
	PolicyValue      *string `json:"PolicyValue,omitempty" xml:"PolicyValue,omitempty"`
}

func (s ListAppPoliciesForIdentityResponseBodyAppPolicyList) String() string {
	return tea.Prettify(s)
}

func (s ListAppPoliciesForIdentityResponseBodyAppPolicyList) GoString() string {
	return s.String()
}

func (s *ListAppPoliciesForIdentityResponseBodyAppPolicyList) SetAppId(v string) *ListAppPoliciesForIdentityResponseBodyAppPolicyList {
	s.AppId = &v
	return s
}

func (s *ListAppPoliciesForIdentityResponseBodyAppPolicyList) SetCreationTime(v string) *ListAppPoliciesForIdentityResponseBodyAppPolicyList {
	s.CreationTime = &v
	return s
}

func (s *ListAppPoliciesForIdentityResponseBodyAppPolicyList) SetDescription(v string) *ListAppPoliciesForIdentityResponseBodyAppPolicyList {
	s.Description = &v
	return s
}

func (s *ListAppPoliciesForIdentityResponseBodyAppPolicyList) SetModificationTime(v string) *ListAppPoliciesForIdentityResponseBodyAppPolicyList {
	s.ModificationTime = &v
	return s
}

func (s *ListAppPoliciesForIdentityResponseBodyAppPolicyList) SetPolicyName(v string) *ListAppPoliciesForIdentityResponseBodyAppPolicyList {
	s.PolicyName = &v
	return s
}

func (s *ListAppPoliciesForIdentityResponseBodyAppPolicyList) SetPolicyType(v string) *ListAppPoliciesForIdentityResponseBodyAppPolicyList {
	s.PolicyType = &v
	return s
}

func (s *ListAppPoliciesForIdentityResponseBodyAppPolicyList) SetPolicyValue(v string) *ListAppPoliciesForIdentityResponseBodyAppPolicyList {
	s.PolicyValue = &v
	return s
}

type ListAppPoliciesForIdentityResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListAppPoliciesForIdentityResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListAppPoliciesForIdentityResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAppPoliciesForIdentityResponse) GoString() string {
	return s.String()
}

func (s *ListAppPoliciesForIdentityResponse) SetHeaders(v map[string]*string) *ListAppPoliciesForIdentityResponse {
	s.Headers = v
	return s
}

func (s *ListAppPoliciesForIdentityResponse) SetBody(v *ListAppPoliciesForIdentityResponseBody) *ListAppPoliciesForIdentityResponse {
	s.Body = v
	return s
}

type ListAuditSecurityIpRequest struct {
	SecurityGroupName *string `json:"SecurityGroupName,omitempty" xml:"SecurityGroupName,omitempty"`
}

func (s ListAuditSecurityIpRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAuditSecurityIpRequest) GoString() string {
	return s.String()
}

func (s *ListAuditSecurityIpRequest) SetSecurityGroupName(v string) *ListAuditSecurityIpRequest {
	s.SecurityGroupName = &v
	return s
}

type ListAuditSecurityIpResponseBody struct {
	RequestId      *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SecurityIpList []*ListAuditSecurityIpResponseBodySecurityIpList `json:"SecurityIpList,omitempty" xml:"SecurityIpList,omitempty" type:"Repeated"`
}

func (s ListAuditSecurityIpResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAuditSecurityIpResponseBody) GoString() string {
	return s.String()
}

func (s *ListAuditSecurityIpResponseBody) SetRequestId(v string) *ListAuditSecurityIpResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAuditSecurityIpResponseBody) SetSecurityIpList(v []*ListAuditSecurityIpResponseBodySecurityIpList) *ListAuditSecurityIpResponseBody {
	s.SecurityIpList = v
	return s
}

type ListAuditSecurityIpResponseBodySecurityIpList struct {
	CreationTime      *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	Ips               *string `json:"Ips,omitempty" xml:"Ips,omitempty"`
	ModificationTime  *string `json:"ModificationTime,omitempty" xml:"ModificationTime,omitempty"`
	SecurityGroupName *string `json:"SecurityGroupName,omitempty" xml:"SecurityGroupName,omitempty"`
}

func (s ListAuditSecurityIpResponseBodySecurityIpList) String() string {
	return tea.Prettify(s)
}

func (s ListAuditSecurityIpResponseBodySecurityIpList) GoString() string {
	return s.String()
}

func (s *ListAuditSecurityIpResponseBodySecurityIpList) SetCreationTime(v string) *ListAuditSecurityIpResponseBodySecurityIpList {
	s.CreationTime = &v
	return s
}

func (s *ListAuditSecurityIpResponseBodySecurityIpList) SetIps(v string) *ListAuditSecurityIpResponseBodySecurityIpList {
	s.Ips = &v
	return s
}

func (s *ListAuditSecurityIpResponseBodySecurityIpList) SetModificationTime(v string) *ListAuditSecurityIpResponseBodySecurityIpList {
	s.ModificationTime = &v
	return s
}

func (s *ListAuditSecurityIpResponseBodySecurityIpList) SetSecurityGroupName(v string) *ListAuditSecurityIpResponseBodySecurityIpList {
	s.SecurityGroupName = &v
	return s
}

type ListAuditSecurityIpResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListAuditSecurityIpResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListAuditSecurityIpResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAuditSecurityIpResponse) GoString() string {
	return s.String()
}

func (s *ListAuditSecurityIpResponse) SetHeaders(v map[string]*string) *ListAuditSecurityIpResponse {
	s.Headers = v
	return s
}

func (s *ListAuditSecurityIpResponse) SetBody(v *ListAuditSecurityIpResponseBody) *ListAuditSecurityIpResponse {
	s.Body = v
	return s
}

type ListDetectionJobRequest struct {
	VideoId *string `json:"VideoId,omitempty" xml:"VideoId,omitempty"`
}

func (s ListDetectionJobRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDetectionJobRequest) GoString() string {
	return s.String()
}

func (s *ListDetectionJobRequest) SetVideoId(v string) *ListDetectionJobRequest {
	s.VideoId = &v
	return s
}

type ListDetectionJobResponseBody struct {
	DetectionJobList []*ListDetectionJobResponseBodyDetectionJobList `json:"DetectionJobList,omitempty" xml:"DetectionJobList,omitempty" type:"Repeated"`
	RequestId        *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListDetectionJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDetectionJobResponseBody) GoString() string {
	return s.String()
}

func (s *ListDetectionJobResponseBody) SetDetectionJobList(v []*ListDetectionJobResponseBodyDetectionJobList) *ListDetectionJobResponseBody {
	s.DetectionJobList = v
	return s
}

func (s *ListDetectionJobResponseBody) SetRequestId(v string) *ListDetectionJobResponseBody {
	s.RequestId = &v
	return s
}

type ListDetectionJobResponseBodyDetectionJobList struct {
	BeginTime          *string `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	CopyrightBeginTime *string `json:"CopyrightBeginTime,omitempty" xml:"CopyrightBeginTime,omitempty"`
	CopyrightEndTime   *string `json:"CopyrightEndTime,omitempty" xml:"CopyrightEndTime,omitempty"`
	CopyrightFile      *string `json:"CopyrightFile,omitempty" xml:"CopyrightFile,omitempty"`
	CopyrightStatus    *string `json:"CopyrightStatus,omitempty" xml:"CopyrightStatus,omitempty"`
	CreateTime         *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	EndTime            *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	JobId              *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	ModifyTime         *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	TemplateId         *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	VideoId            *string `json:"VideoId,omitempty" xml:"VideoId,omitempty"`
	WhitelistUrls      *string `json:"WhitelistUrls,omitempty" xml:"WhitelistUrls,omitempty"`
}

func (s ListDetectionJobResponseBodyDetectionJobList) String() string {
	return tea.Prettify(s)
}

func (s ListDetectionJobResponseBodyDetectionJobList) GoString() string {
	return s.String()
}

func (s *ListDetectionJobResponseBodyDetectionJobList) SetBeginTime(v string) *ListDetectionJobResponseBodyDetectionJobList {
	s.BeginTime = &v
	return s
}

func (s *ListDetectionJobResponseBodyDetectionJobList) SetCopyrightBeginTime(v string) *ListDetectionJobResponseBodyDetectionJobList {
	s.CopyrightBeginTime = &v
	return s
}

func (s *ListDetectionJobResponseBodyDetectionJobList) SetCopyrightEndTime(v string) *ListDetectionJobResponseBodyDetectionJobList {
	s.CopyrightEndTime = &v
	return s
}

func (s *ListDetectionJobResponseBodyDetectionJobList) SetCopyrightFile(v string) *ListDetectionJobResponseBodyDetectionJobList {
	s.CopyrightFile = &v
	return s
}

func (s *ListDetectionJobResponseBodyDetectionJobList) SetCopyrightStatus(v string) *ListDetectionJobResponseBodyDetectionJobList {
	s.CopyrightStatus = &v
	return s
}

func (s *ListDetectionJobResponseBodyDetectionJobList) SetCreateTime(v string) *ListDetectionJobResponseBodyDetectionJobList {
	s.CreateTime = &v
	return s
}

func (s *ListDetectionJobResponseBodyDetectionJobList) SetEndTime(v string) *ListDetectionJobResponseBodyDetectionJobList {
	s.EndTime = &v
	return s
}

func (s *ListDetectionJobResponseBodyDetectionJobList) SetJobId(v string) *ListDetectionJobResponseBodyDetectionJobList {
	s.JobId = &v
	return s
}

func (s *ListDetectionJobResponseBodyDetectionJobList) SetModifyTime(v string) *ListDetectionJobResponseBodyDetectionJobList {
	s.ModifyTime = &v
	return s
}

func (s *ListDetectionJobResponseBodyDetectionJobList) SetTemplateId(v string) *ListDetectionJobResponseBodyDetectionJobList {
	s.TemplateId = &v
	return s
}

func (s *ListDetectionJobResponseBodyDetectionJobList) SetVideoId(v string) *ListDetectionJobResponseBodyDetectionJobList {
	s.VideoId = &v
	return s
}

func (s *ListDetectionJobResponseBodyDetectionJobList) SetWhitelistUrls(v string) *ListDetectionJobResponseBodyDetectionJobList {
	s.WhitelistUrls = &v
	return s
}

type ListDetectionJobResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListDetectionJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListDetectionJobResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDetectionJobResponse) GoString() string {
	return s.String()
}

func (s *ListDetectionJobResponse) SetHeaders(v map[string]*string) *ListDetectionJobResponse {
	s.Headers = v
	return s
}

func (s *ListDetectionJobResponse) SetBody(v *ListDetectionJobResponseBody) *ListDetectionJobResponse {
	s.Body = v
	return s
}

type ListDetectionTemplateRequest struct {
	Period       *string `json:"Period,omitempty" xml:"Period,omitempty"`
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
}

func (s ListDetectionTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDetectionTemplateRequest) GoString() string {
	return s.String()
}

func (s *ListDetectionTemplateRequest) SetPeriod(v string) *ListDetectionTemplateRequest {
	s.Period = &v
	return s
}

func (s *ListDetectionTemplateRequest) SetTemplateName(v string) *ListDetectionTemplateRequest {
	s.TemplateName = &v
	return s
}

type ListDetectionTemplateResponseBody struct {
	DetectionTemplateList []*ListDetectionTemplateResponseBodyDetectionTemplateList `json:"DetectionTemplateList,omitempty" xml:"DetectionTemplateList,omitempty" type:"Repeated"`
	RequestId             *string                                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListDetectionTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDetectionTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *ListDetectionTemplateResponseBody) SetDetectionTemplateList(v []*ListDetectionTemplateResponseBodyDetectionTemplateList) *ListDetectionTemplateResponseBody {
	s.DetectionTemplateList = v
	return s
}

func (s *ListDetectionTemplateResponseBody) SetRequestId(v string) *ListDetectionTemplateResponseBody {
	s.RequestId = &v
	return s
}

type ListDetectionTemplateResponseBodyDetectionTemplateList struct {
	CreateTime   *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	ModifyTime   *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	Period       *string `json:"Period,omitempty" xml:"Period,omitempty"`
	Platform     *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	TemplateId   *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	UserId       *int64  `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListDetectionTemplateResponseBodyDetectionTemplateList) String() string {
	return tea.Prettify(s)
}

func (s ListDetectionTemplateResponseBodyDetectionTemplateList) GoString() string {
	return s.String()
}

func (s *ListDetectionTemplateResponseBodyDetectionTemplateList) SetCreateTime(v string) *ListDetectionTemplateResponseBodyDetectionTemplateList {
	s.CreateTime = &v
	return s
}

func (s *ListDetectionTemplateResponseBodyDetectionTemplateList) SetModifyTime(v string) *ListDetectionTemplateResponseBodyDetectionTemplateList {
	s.ModifyTime = &v
	return s
}

func (s *ListDetectionTemplateResponseBodyDetectionTemplateList) SetPeriod(v string) *ListDetectionTemplateResponseBodyDetectionTemplateList {
	s.Period = &v
	return s
}

func (s *ListDetectionTemplateResponseBodyDetectionTemplateList) SetPlatform(v string) *ListDetectionTemplateResponseBodyDetectionTemplateList {
	s.Platform = &v
	return s
}

func (s *ListDetectionTemplateResponseBodyDetectionTemplateList) SetTemplateId(v string) *ListDetectionTemplateResponseBodyDetectionTemplateList {
	s.TemplateId = &v
	return s
}

func (s *ListDetectionTemplateResponseBodyDetectionTemplateList) SetTemplateName(v string) *ListDetectionTemplateResponseBodyDetectionTemplateList {
	s.TemplateName = &v
	return s
}

func (s *ListDetectionTemplateResponseBodyDetectionTemplateList) SetUserId(v int64) *ListDetectionTemplateResponseBodyDetectionTemplateList {
	s.UserId = &v
	return s
}

type ListDetectionTemplateResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListDetectionTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListDetectionTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDetectionTemplateResponse) GoString() string {
	return s.String()
}

func (s *ListDetectionTemplateResponse) SetHeaders(v map[string]*string) *ListDetectionTemplateResponse {
	s.Headers = v
	return s
}

func (s *ListDetectionTemplateResponse) SetBody(v *ListDetectionTemplateResponseBody) *ListDetectionTemplateResponse {
	s.Body = v
	return s
}

type ListDynamicImageRequest struct {
	VideoId *string `json:"VideoId,omitempty" xml:"VideoId,omitempty"`
}

func (s ListDynamicImageRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDynamicImageRequest) GoString() string {
	return s.String()
}

func (s *ListDynamicImageRequest) SetVideoId(v string) *ListDynamicImageRequest {
	s.VideoId = &v
	return s
}

type ListDynamicImageResponseBody struct {
	DynamicImageList []*ListDynamicImageResponseBodyDynamicImageList `json:"DynamicImageList,omitempty" xml:"DynamicImageList,omitempty" type:"Repeated"`
	RequestId        *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListDynamicImageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDynamicImageResponseBody) GoString() string {
	return s.String()
}

func (s *ListDynamicImageResponseBody) SetDynamicImageList(v []*ListDynamicImageResponseBodyDynamicImageList) *ListDynamicImageResponseBody {
	s.DynamicImageList = v
	return s
}

func (s *ListDynamicImageResponseBody) SetRequestId(v string) *ListDynamicImageResponseBody {
	s.RequestId = &v
	return s
}

type ListDynamicImageResponseBodyDynamicImageList struct {
	CreationTime   *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	Duration       *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	DynamicImageId *string `json:"DynamicImageId,omitempty" xml:"DynamicImageId,omitempty"`
	FileSize       *string `json:"FileSize,omitempty" xml:"FileSize,omitempty"`
	FileURL        *string `json:"FileURL,omitempty" xml:"FileURL,omitempty"`
	Format         *string `json:"Format,omitempty" xml:"Format,omitempty"`
	Fps            *string `json:"Fps,omitempty" xml:"Fps,omitempty"`
	Height         *string `json:"Height,omitempty" xml:"Height,omitempty"`
	JobId          *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	VideoId        *string `json:"VideoId,omitempty" xml:"VideoId,omitempty"`
	Width          *string `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s ListDynamicImageResponseBodyDynamicImageList) String() string {
	return tea.Prettify(s)
}

func (s ListDynamicImageResponseBodyDynamicImageList) GoString() string {
	return s.String()
}

func (s *ListDynamicImageResponseBodyDynamicImageList) SetCreationTime(v string) *ListDynamicImageResponseBodyDynamicImageList {
	s.CreationTime = &v
	return s
}

func (s *ListDynamicImageResponseBodyDynamicImageList) SetDuration(v string) *ListDynamicImageResponseBodyDynamicImageList {
	s.Duration = &v
	return s
}

func (s *ListDynamicImageResponseBodyDynamicImageList) SetDynamicImageId(v string) *ListDynamicImageResponseBodyDynamicImageList {
	s.DynamicImageId = &v
	return s
}

func (s *ListDynamicImageResponseBodyDynamicImageList) SetFileSize(v string) *ListDynamicImageResponseBodyDynamicImageList {
	s.FileSize = &v
	return s
}

func (s *ListDynamicImageResponseBodyDynamicImageList) SetFileURL(v string) *ListDynamicImageResponseBodyDynamicImageList {
	s.FileURL = &v
	return s
}

func (s *ListDynamicImageResponseBodyDynamicImageList) SetFormat(v string) *ListDynamicImageResponseBodyDynamicImageList {
	s.Format = &v
	return s
}

func (s *ListDynamicImageResponseBodyDynamicImageList) SetFps(v string) *ListDynamicImageResponseBodyDynamicImageList {
	s.Fps = &v
	return s
}

func (s *ListDynamicImageResponseBodyDynamicImageList) SetHeight(v string) *ListDynamicImageResponseBodyDynamicImageList {
	s.Height = &v
	return s
}

func (s *ListDynamicImageResponseBodyDynamicImageList) SetJobId(v string) *ListDynamicImageResponseBodyDynamicImageList {
	s.JobId = &v
	return s
}

func (s *ListDynamicImageResponseBodyDynamicImageList) SetVideoId(v string) *ListDynamicImageResponseBodyDynamicImageList {
	s.VideoId = &v
	return s
}

func (s *ListDynamicImageResponseBodyDynamicImageList) SetWidth(v string) *ListDynamicImageResponseBodyDynamicImageList {
	s.Width = &v
	return s
}

type ListDynamicImageResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListDynamicImageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListDynamicImageResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDynamicImageResponse) GoString() string {
	return s.String()
}

func (s *ListDynamicImageResponse) SetHeaders(v map[string]*string) *ListDynamicImageResponse {
	s.Headers = v
	return s
}

func (s *ListDynamicImageResponse) SetBody(v *ListDynamicImageResponseBody) *ListDynamicImageResponse {
	s.Body = v
	return s
}

type ListLetterSendJobRequest struct {
	DetectionId *string `json:"DetectionId,omitempty" xml:"DetectionId,omitempty"`
	TemplateId  *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	ToAddress   *string `json:"ToAddress,omitempty" xml:"ToAddress,omitempty"`
}

func (s ListLetterSendJobRequest) String() string {
	return tea.Prettify(s)
}

func (s ListLetterSendJobRequest) GoString() string {
	return s.String()
}

func (s *ListLetterSendJobRequest) SetDetectionId(v string) *ListLetterSendJobRequest {
	s.DetectionId = &v
	return s
}

func (s *ListLetterSendJobRequest) SetTemplateId(v string) *ListLetterSendJobRequest {
	s.TemplateId = &v
	return s
}

func (s *ListLetterSendJobRequest) SetToAddress(v string) *ListLetterSendJobRequest {
	s.ToAddress = &v
	return s
}

type ListLetterSendJobResponseBody struct {
	LetterJobList []*ListLetterSendJobResponseBodyLetterJobList `json:"LetterJobList,omitempty" xml:"LetterJobList,omitempty" type:"Repeated"`
	RequestId     *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListLetterSendJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListLetterSendJobResponseBody) GoString() string {
	return s.String()
}

func (s *ListLetterSendJobResponseBody) SetLetterJobList(v []*ListLetterSendJobResponseBodyLetterJobList) *ListLetterSendJobResponseBody {
	s.LetterJobList = v
	return s
}

func (s *ListLetterSendJobResponseBody) SetRequestId(v string) *ListLetterSendJobResponseBody {
	s.RequestId = &v
	return s
}

type ListLetterSendJobResponseBodyLetterJobList struct {
	BccAddress  *string `json:"BccAddress,omitempty" xml:"BccAddress,omitempty"`
	Body        *string `json:"Body,omitempty" xml:"Body,omitempty"`
	CcAddress   *string `json:"CcAddress,omitempty" xml:"CcAddress,omitempty"`
	CreateTime  *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DetectionId *string `json:"DetectionId,omitempty" xml:"DetectionId,omitempty"`
	JobId       *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	ModifyTime  *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	SendTime    *string `json:"SendTime,omitempty" xml:"SendTime,omitempty"`
	TemplateId  *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	Title       *string `json:"Title,omitempty" xml:"Title,omitempty"`
	ToAddress   *string `json:"ToAddress,omitempty" xml:"ToAddress,omitempty"`
	UserId      *int64  `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListLetterSendJobResponseBodyLetterJobList) String() string {
	return tea.Prettify(s)
}

func (s ListLetterSendJobResponseBodyLetterJobList) GoString() string {
	return s.String()
}

func (s *ListLetterSendJobResponseBodyLetterJobList) SetBccAddress(v string) *ListLetterSendJobResponseBodyLetterJobList {
	s.BccAddress = &v
	return s
}

func (s *ListLetterSendJobResponseBodyLetterJobList) SetBody(v string) *ListLetterSendJobResponseBodyLetterJobList {
	s.Body = &v
	return s
}

func (s *ListLetterSendJobResponseBodyLetterJobList) SetCcAddress(v string) *ListLetterSendJobResponseBodyLetterJobList {
	s.CcAddress = &v
	return s
}

func (s *ListLetterSendJobResponseBodyLetterJobList) SetCreateTime(v string) *ListLetterSendJobResponseBodyLetterJobList {
	s.CreateTime = &v
	return s
}

func (s *ListLetterSendJobResponseBodyLetterJobList) SetDetectionId(v string) *ListLetterSendJobResponseBodyLetterJobList {
	s.DetectionId = &v
	return s
}

func (s *ListLetterSendJobResponseBodyLetterJobList) SetJobId(v string) *ListLetterSendJobResponseBodyLetterJobList {
	s.JobId = &v
	return s
}

func (s *ListLetterSendJobResponseBodyLetterJobList) SetModifyTime(v string) *ListLetterSendJobResponseBodyLetterJobList {
	s.ModifyTime = &v
	return s
}

func (s *ListLetterSendJobResponseBodyLetterJobList) SetSendTime(v string) *ListLetterSendJobResponseBodyLetterJobList {
	s.SendTime = &v
	return s
}

func (s *ListLetterSendJobResponseBodyLetterJobList) SetTemplateId(v string) *ListLetterSendJobResponseBodyLetterJobList {
	s.TemplateId = &v
	return s
}

func (s *ListLetterSendJobResponseBodyLetterJobList) SetTitle(v string) *ListLetterSendJobResponseBodyLetterJobList {
	s.Title = &v
	return s
}

func (s *ListLetterSendJobResponseBodyLetterJobList) SetToAddress(v string) *ListLetterSendJobResponseBodyLetterJobList {
	s.ToAddress = &v
	return s
}

func (s *ListLetterSendJobResponseBodyLetterJobList) SetUserId(v int64) *ListLetterSendJobResponseBodyLetterJobList {
	s.UserId = &v
	return s
}

type ListLetterSendJobResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListLetterSendJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListLetterSendJobResponse) String() string {
	return tea.Prettify(s)
}

func (s ListLetterSendJobResponse) GoString() string {
	return s.String()
}

func (s *ListLetterSendJobResponse) SetHeaders(v map[string]*string) *ListLetterSendJobResponse {
	s.Headers = v
	return s
}

func (s *ListLetterSendJobResponse) SetBody(v *ListLetterSendJobResponseBody) *ListLetterSendJobResponse {
	s.Body = v
	return s
}

type ListLiveRecordVideoRequest struct {
	AppName    *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndTime    *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	PageNo     *int32  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SortBy     *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	StartTime  *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	StreamName *string `json:"StreamName,omitempty" xml:"StreamName,omitempty"`
}

func (s ListLiveRecordVideoRequest) String() string {
	return tea.Prettify(s)
}

func (s ListLiveRecordVideoRequest) GoString() string {
	return s.String()
}

func (s *ListLiveRecordVideoRequest) SetAppName(v string) *ListLiveRecordVideoRequest {
	s.AppName = &v
	return s
}

func (s *ListLiveRecordVideoRequest) SetDomainName(v string) *ListLiveRecordVideoRequest {
	s.DomainName = &v
	return s
}

func (s *ListLiveRecordVideoRequest) SetEndTime(v string) *ListLiveRecordVideoRequest {
	s.EndTime = &v
	return s
}

func (s *ListLiveRecordVideoRequest) SetPageNo(v int32) *ListLiveRecordVideoRequest {
	s.PageNo = &v
	return s
}

func (s *ListLiveRecordVideoRequest) SetPageSize(v int32) *ListLiveRecordVideoRequest {
	s.PageSize = &v
	return s
}

func (s *ListLiveRecordVideoRequest) SetSortBy(v string) *ListLiveRecordVideoRequest {
	s.SortBy = &v
	return s
}

func (s *ListLiveRecordVideoRequest) SetStartTime(v string) *ListLiveRecordVideoRequest {
	s.StartTime = &v
	return s
}

func (s *ListLiveRecordVideoRequest) SetStreamName(v string) *ListLiveRecordVideoRequest {
	s.StreamName = &v
	return s
}

type ListLiveRecordVideoResponseBody struct {
	LiveRecordVideoList *ListLiveRecordVideoResponseBodyLiveRecordVideoList `json:"LiveRecordVideoList,omitempty" xml:"LiveRecordVideoList,omitempty" type:"Struct"`
	RequestId           *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Total               *int32                                              `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListLiveRecordVideoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListLiveRecordVideoResponseBody) GoString() string {
	return s.String()
}

func (s *ListLiveRecordVideoResponseBody) SetLiveRecordVideoList(v *ListLiveRecordVideoResponseBodyLiveRecordVideoList) *ListLiveRecordVideoResponseBody {
	s.LiveRecordVideoList = v
	return s
}

func (s *ListLiveRecordVideoResponseBody) SetRequestId(v string) *ListLiveRecordVideoResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListLiveRecordVideoResponseBody) SetTotal(v int32) *ListLiveRecordVideoResponseBody {
	s.Total = &v
	return s
}

type ListLiveRecordVideoResponseBodyLiveRecordVideoList struct {
	LiveRecordVideo []*ListLiveRecordVideoResponseBodyLiveRecordVideoListLiveRecordVideo `json:"LiveRecordVideo,omitempty" xml:"LiveRecordVideo,omitempty" type:"Repeated"`
}

func (s ListLiveRecordVideoResponseBodyLiveRecordVideoList) String() string {
	return tea.Prettify(s)
}

func (s ListLiveRecordVideoResponseBodyLiveRecordVideoList) GoString() string {
	return s.String()
}

func (s *ListLiveRecordVideoResponseBodyLiveRecordVideoList) SetLiveRecordVideo(v []*ListLiveRecordVideoResponseBodyLiveRecordVideoListLiveRecordVideo) *ListLiveRecordVideoResponseBodyLiveRecordVideoList {
	s.LiveRecordVideo = v
	return s
}

type ListLiveRecordVideoResponseBodyLiveRecordVideoListLiveRecordVideo struct {
	AppName         *string                                                                 `json:"AppName,omitempty" xml:"AppName,omitempty"`
	DomainName      *string                                                                 `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	PlaylistId      *string                                                                 `json:"PlaylistId,omitempty" xml:"PlaylistId,omitempty"`
	RecordEndTime   *string                                                                 `json:"RecordEndTime,omitempty" xml:"RecordEndTime,omitempty"`
	RecordStartTime *string                                                                 `json:"RecordStartTime,omitempty" xml:"RecordStartTime,omitempty"`
	StreamName      *string                                                                 `json:"StreamName,omitempty" xml:"StreamName,omitempty"`
	Video           *ListLiveRecordVideoResponseBodyLiveRecordVideoListLiveRecordVideoVideo `json:"Video,omitempty" xml:"Video,omitempty" type:"Struct"`
}

func (s ListLiveRecordVideoResponseBodyLiveRecordVideoListLiveRecordVideo) String() string {
	return tea.Prettify(s)
}

func (s ListLiveRecordVideoResponseBodyLiveRecordVideoListLiveRecordVideo) GoString() string {
	return s.String()
}

func (s *ListLiveRecordVideoResponseBodyLiveRecordVideoListLiveRecordVideo) SetAppName(v string) *ListLiveRecordVideoResponseBodyLiveRecordVideoListLiveRecordVideo {
	s.AppName = &v
	return s
}

func (s *ListLiveRecordVideoResponseBodyLiveRecordVideoListLiveRecordVideo) SetDomainName(v string) *ListLiveRecordVideoResponseBodyLiveRecordVideoListLiveRecordVideo {
	s.DomainName = &v
	return s
}

func (s *ListLiveRecordVideoResponseBodyLiveRecordVideoListLiveRecordVideo) SetPlaylistId(v string) *ListLiveRecordVideoResponseBodyLiveRecordVideoListLiveRecordVideo {
	s.PlaylistId = &v
	return s
}

func (s *ListLiveRecordVideoResponseBodyLiveRecordVideoListLiveRecordVideo) SetRecordEndTime(v string) *ListLiveRecordVideoResponseBodyLiveRecordVideoListLiveRecordVideo {
	s.RecordEndTime = &v
	return s
}

func (s *ListLiveRecordVideoResponseBodyLiveRecordVideoListLiveRecordVideo) SetRecordStartTime(v string) *ListLiveRecordVideoResponseBodyLiveRecordVideoListLiveRecordVideo {
	s.RecordStartTime = &v
	return s
}

func (s *ListLiveRecordVideoResponseBodyLiveRecordVideoListLiveRecordVideo) SetStreamName(v string) *ListLiveRecordVideoResponseBodyLiveRecordVideoListLiveRecordVideo {
	s.StreamName = &v
	return s
}

func (s *ListLiveRecordVideoResponseBodyLiveRecordVideoListLiveRecordVideo) SetVideo(v *ListLiveRecordVideoResponseBodyLiveRecordVideoListLiveRecordVideoVideo) *ListLiveRecordVideoResponseBodyLiveRecordVideoListLiveRecordVideo {
	s.Video = v
	return s
}

type ListLiveRecordVideoResponseBodyLiveRecordVideoListLiveRecordVideoVideo struct {
	CateId          *int32                                                                           `json:"CateId,omitempty" xml:"CateId,omitempty"`
	CateName        *string                                                                          `json:"CateName,omitempty" xml:"CateName,omitempty"`
	CoverURL        *string                                                                          `json:"CoverURL,omitempty" xml:"CoverURL,omitempty"`
	CreationTime    *string                                                                          `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	Description     *string                                                                          `json:"Description,omitempty" xml:"Description,omitempty"`
	Duration        *float32                                                                         `json:"Duration,omitempty" xml:"Duration,omitempty"`
	ModifyTime      *string                                                                          `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	Size            *int64                                                                           `json:"Size,omitempty" xml:"Size,omitempty"`
	Snapshots       *ListLiveRecordVideoResponseBodyLiveRecordVideoListLiveRecordVideoVideoSnapshots `json:"Snapshots,omitempty" xml:"Snapshots,omitempty" type:"Struct"`
	Status          *string                                                                          `json:"Status,omitempty" xml:"Status,omitempty"`
	Tags            *string                                                                          `json:"Tags,omitempty" xml:"Tags,omitempty"`
	TemplateGroupId *string                                                                          `json:"TemplateGroupId,omitempty" xml:"TemplateGroupId,omitempty"`
	Title           *string                                                                          `json:"Title,omitempty" xml:"Title,omitempty"`
	VideoId         *string                                                                          `json:"VideoId,omitempty" xml:"VideoId,omitempty"`
}

func (s ListLiveRecordVideoResponseBodyLiveRecordVideoListLiveRecordVideoVideo) String() string {
	return tea.Prettify(s)
}

func (s ListLiveRecordVideoResponseBodyLiveRecordVideoListLiveRecordVideoVideo) GoString() string {
	return s.String()
}

func (s *ListLiveRecordVideoResponseBodyLiveRecordVideoListLiveRecordVideoVideo) SetCateId(v int32) *ListLiveRecordVideoResponseBodyLiveRecordVideoListLiveRecordVideoVideo {
	s.CateId = &v
	return s
}

func (s *ListLiveRecordVideoResponseBodyLiveRecordVideoListLiveRecordVideoVideo) SetCateName(v string) *ListLiveRecordVideoResponseBodyLiveRecordVideoListLiveRecordVideoVideo {
	s.CateName = &v
	return s
}

func (s *ListLiveRecordVideoResponseBodyLiveRecordVideoListLiveRecordVideoVideo) SetCoverURL(v string) *ListLiveRecordVideoResponseBodyLiveRecordVideoListLiveRecordVideoVideo {
	s.CoverURL = &v
	return s
}

func (s *ListLiveRecordVideoResponseBodyLiveRecordVideoListLiveRecordVideoVideo) SetCreationTime(v string) *ListLiveRecordVideoResponseBodyLiveRecordVideoListLiveRecordVideoVideo {
	s.CreationTime = &v
	return s
}

func (s *ListLiveRecordVideoResponseBodyLiveRecordVideoListLiveRecordVideoVideo) SetDescription(v string) *ListLiveRecordVideoResponseBodyLiveRecordVideoListLiveRecordVideoVideo {
	s.Description = &v
	return s
}

func (s *ListLiveRecordVideoResponseBodyLiveRecordVideoListLiveRecordVideoVideo) SetDuration(v float32) *ListLiveRecordVideoResponseBodyLiveRecordVideoListLiveRecordVideoVideo {
	s.Duration = &v
	return s
}

func (s *ListLiveRecordVideoResponseBodyLiveRecordVideoListLiveRecordVideoVideo) SetModifyTime(v string) *ListLiveRecordVideoResponseBodyLiveRecordVideoListLiveRecordVideoVideo {
	s.ModifyTime = &v
	return s
}

func (s *ListLiveRecordVideoResponseBodyLiveRecordVideoListLiveRecordVideoVideo) SetSize(v int64) *ListLiveRecordVideoResponseBodyLiveRecordVideoListLiveRecordVideoVideo {
	s.Size = &v
	return s
}

func (s *ListLiveRecordVideoResponseBodyLiveRecordVideoListLiveRecordVideoVideo) SetSnapshots(v *ListLiveRecordVideoResponseBodyLiveRecordVideoListLiveRecordVideoVideoSnapshots) *ListLiveRecordVideoResponseBodyLiveRecordVideoListLiveRecordVideoVideo {
	s.Snapshots = v
	return s
}

func (s *ListLiveRecordVideoResponseBodyLiveRecordVideoListLiveRecordVideoVideo) SetStatus(v string) *ListLiveRecordVideoResponseBodyLiveRecordVideoListLiveRecordVideoVideo {
	s.Status = &v
	return s
}

func (s *ListLiveRecordVideoResponseBodyLiveRecordVideoListLiveRecordVideoVideo) SetTags(v string) *ListLiveRecordVideoResponseBodyLiveRecordVideoListLiveRecordVideoVideo {
	s.Tags = &v
	return s
}

func (s *ListLiveRecordVideoResponseBodyLiveRecordVideoListLiveRecordVideoVideo) SetTemplateGroupId(v string) *ListLiveRecordVideoResponseBodyLiveRecordVideoListLiveRecordVideoVideo {
	s.TemplateGroupId = &v
	return s
}

func (s *ListLiveRecordVideoResponseBodyLiveRecordVideoListLiveRecordVideoVideo) SetTitle(v string) *ListLiveRecordVideoResponseBodyLiveRecordVideoListLiveRecordVideoVideo {
	s.Title = &v
	return s
}

func (s *ListLiveRecordVideoResponseBodyLiveRecordVideoListLiveRecordVideoVideo) SetVideoId(v string) *ListLiveRecordVideoResponseBodyLiveRecordVideoListLiveRecordVideoVideo {
	s.VideoId = &v
	return s
}

type ListLiveRecordVideoResponseBodyLiveRecordVideoListLiveRecordVideoVideoSnapshots struct {
	Snapshot []*string `json:"Snapshot,omitempty" xml:"Snapshot,omitempty" type:"Repeated"`
}

func (s ListLiveRecordVideoResponseBodyLiveRecordVideoListLiveRecordVideoVideoSnapshots) String() string {
	return tea.Prettify(s)
}

func (s ListLiveRecordVideoResponseBodyLiveRecordVideoListLiveRecordVideoVideoSnapshots) GoString() string {
	return s.String()
}

func (s *ListLiveRecordVideoResponseBodyLiveRecordVideoListLiveRecordVideoVideoSnapshots) SetSnapshot(v []*string) *ListLiveRecordVideoResponseBodyLiveRecordVideoListLiveRecordVideoVideoSnapshots {
	s.Snapshot = v
	return s
}

type ListLiveRecordVideoResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListLiveRecordVideoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListLiveRecordVideoResponse) String() string {
	return tea.Prettify(s)
}

func (s ListLiveRecordVideoResponse) GoString() string {
	return s.String()
}

func (s *ListLiveRecordVideoResponse) SetHeaders(v map[string]*string) *ListLiveRecordVideoResponse {
	s.Headers = v
	return s
}

func (s *ListLiveRecordVideoResponse) SetBody(v *ListLiveRecordVideoResponseBody) *ListLiveRecordVideoResponse {
	s.Body = v
	return s
}

type ListMediaDNADeleteJobRequest struct {
	JobIds               *string `json:"JobIds,omitempty" xml:"JobIds,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *string `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ListMediaDNADeleteJobRequest) String() string {
	return tea.Prettify(s)
}

func (s ListMediaDNADeleteJobRequest) GoString() string {
	return s.String()
}

func (s *ListMediaDNADeleteJobRequest) SetJobIds(v string) *ListMediaDNADeleteJobRequest {
	s.JobIds = &v
	return s
}

func (s *ListMediaDNADeleteJobRequest) SetOwnerAccount(v string) *ListMediaDNADeleteJobRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ListMediaDNADeleteJobRequest) SetOwnerId(v string) *ListMediaDNADeleteJobRequest {
	s.OwnerId = &v
	return s
}

func (s *ListMediaDNADeleteJobRequest) SetResourceOwnerAccount(v string) *ListMediaDNADeleteJobRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListMediaDNADeleteJobRequest) SetResourceOwnerId(v string) *ListMediaDNADeleteJobRequest {
	s.ResourceOwnerId = &v
	return s
}

type ListMediaDNADeleteJobResponseBody struct {
	AIJobList        *ListMediaDNADeleteJobResponseBodyAIJobList        `json:"AIJobList,omitempty" xml:"AIJobList,omitempty" type:"Struct"`
	NonExistAIJobIds *ListMediaDNADeleteJobResponseBodyNonExistAIJobIds `json:"NonExistAIJobIds,omitempty" xml:"NonExistAIJobIds,omitempty" type:"Struct"`
	RequestId        *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListMediaDNADeleteJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListMediaDNADeleteJobResponseBody) GoString() string {
	return s.String()
}

func (s *ListMediaDNADeleteJobResponseBody) SetAIJobList(v *ListMediaDNADeleteJobResponseBodyAIJobList) *ListMediaDNADeleteJobResponseBody {
	s.AIJobList = v
	return s
}

func (s *ListMediaDNADeleteJobResponseBody) SetNonExistAIJobIds(v *ListMediaDNADeleteJobResponseBodyNonExistAIJobIds) *ListMediaDNADeleteJobResponseBody {
	s.NonExistAIJobIds = v
	return s
}

func (s *ListMediaDNADeleteJobResponseBody) SetRequestId(v string) *ListMediaDNADeleteJobResponseBody {
	s.RequestId = &v
	return s
}

type ListMediaDNADeleteJobResponseBodyAIJobList struct {
	AIJob []*ListMediaDNADeleteJobResponseBodyAIJobListAIJob `json:"AIJob,omitempty" xml:"AIJob,omitempty" type:"Repeated"`
}

func (s ListMediaDNADeleteJobResponseBodyAIJobList) String() string {
	return tea.Prettify(s)
}

func (s ListMediaDNADeleteJobResponseBodyAIJobList) GoString() string {
	return s.String()
}

func (s *ListMediaDNADeleteJobResponseBodyAIJobList) SetAIJob(v []*ListMediaDNADeleteJobResponseBodyAIJobListAIJob) *ListMediaDNADeleteJobResponseBodyAIJobList {
	s.AIJob = v
	return s
}

type ListMediaDNADeleteJobResponseBodyAIJobListAIJob struct {
	Code    *string `json:"Code,omitempty" xml:"Code,omitempty"`
	FpDBId  *string `json:"FpDBId,omitempty" xml:"FpDBId,omitempty"`
	JobId   *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Status  *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListMediaDNADeleteJobResponseBodyAIJobListAIJob) String() string {
	return tea.Prettify(s)
}

func (s ListMediaDNADeleteJobResponseBodyAIJobListAIJob) GoString() string {
	return s.String()
}

func (s *ListMediaDNADeleteJobResponseBodyAIJobListAIJob) SetCode(v string) *ListMediaDNADeleteJobResponseBodyAIJobListAIJob {
	s.Code = &v
	return s
}

func (s *ListMediaDNADeleteJobResponseBodyAIJobListAIJob) SetFpDBId(v string) *ListMediaDNADeleteJobResponseBodyAIJobListAIJob {
	s.FpDBId = &v
	return s
}

func (s *ListMediaDNADeleteJobResponseBodyAIJobListAIJob) SetJobId(v string) *ListMediaDNADeleteJobResponseBodyAIJobListAIJob {
	s.JobId = &v
	return s
}

func (s *ListMediaDNADeleteJobResponseBodyAIJobListAIJob) SetMediaId(v string) *ListMediaDNADeleteJobResponseBodyAIJobListAIJob {
	s.MediaId = &v
	return s
}

func (s *ListMediaDNADeleteJobResponseBodyAIJobListAIJob) SetMessage(v string) *ListMediaDNADeleteJobResponseBodyAIJobListAIJob {
	s.Message = &v
	return s
}

func (s *ListMediaDNADeleteJobResponseBodyAIJobListAIJob) SetStatus(v string) *ListMediaDNADeleteJobResponseBodyAIJobListAIJob {
	s.Status = &v
	return s
}

type ListMediaDNADeleteJobResponseBodyNonExistAIJobIds struct {
	String_ []*string `json:"String,omitempty" xml:"String,omitempty" type:"Repeated"`
}

func (s ListMediaDNADeleteJobResponseBodyNonExistAIJobIds) String() string {
	return tea.Prettify(s)
}

func (s ListMediaDNADeleteJobResponseBodyNonExistAIJobIds) GoString() string {
	return s.String()
}

func (s *ListMediaDNADeleteJobResponseBodyNonExistAIJobIds) SetString_(v []*string) *ListMediaDNADeleteJobResponseBodyNonExistAIJobIds {
	s.String_ = v
	return s
}

type ListMediaDNADeleteJobResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListMediaDNADeleteJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListMediaDNADeleteJobResponse) String() string {
	return tea.Prettify(s)
}

func (s ListMediaDNADeleteJobResponse) GoString() string {
	return s.String()
}

func (s *ListMediaDNADeleteJobResponse) SetHeaders(v map[string]*string) *ListMediaDNADeleteJobResponse {
	s.Headers = v
	return s
}

func (s *ListMediaDNADeleteJobResponse) SetBody(v *ListMediaDNADeleteJobResponseBody) *ListMediaDNADeleteJobResponse {
	s.Body = v
	return s
}

type ListSnapshotsRequest struct {
	AuthTimeout  *string `json:"AuthTimeout,omitempty" xml:"AuthTimeout,omitempty"`
	PageNo       *string `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize     *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SnapshotType *string `json:"SnapshotType,omitempty" xml:"SnapshotType,omitempty"`
	VideoId      *string `json:"VideoId,omitempty" xml:"VideoId,omitempty"`
}

func (s ListSnapshotsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSnapshotsRequest) GoString() string {
	return s.String()
}

func (s *ListSnapshotsRequest) SetAuthTimeout(v string) *ListSnapshotsRequest {
	s.AuthTimeout = &v
	return s
}

func (s *ListSnapshotsRequest) SetPageNo(v string) *ListSnapshotsRequest {
	s.PageNo = &v
	return s
}

func (s *ListSnapshotsRequest) SetPageSize(v string) *ListSnapshotsRequest {
	s.PageSize = &v
	return s
}

func (s *ListSnapshotsRequest) SetSnapshotType(v string) *ListSnapshotsRequest {
	s.SnapshotType = &v
	return s
}

func (s *ListSnapshotsRequest) SetVideoId(v string) *ListSnapshotsRequest {
	s.VideoId = &v
	return s
}

type ListSnapshotsResponseBody struct {
	MediaSnapshot *ListSnapshotsResponseBodyMediaSnapshot `json:"MediaSnapshot,omitempty" xml:"MediaSnapshot,omitempty" type:"Struct"`
	RequestId     *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListSnapshotsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSnapshotsResponseBody) GoString() string {
	return s.String()
}

func (s *ListSnapshotsResponseBody) SetMediaSnapshot(v *ListSnapshotsResponseBodyMediaSnapshot) *ListSnapshotsResponseBody {
	s.MediaSnapshot = v
	return s
}

func (s *ListSnapshotsResponseBody) SetRequestId(v string) *ListSnapshotsResponseBody {
	s.RequestId = &v
	return s
}

type ListSnapshotsResponseBodyMediaSnapshot struct {
	CreationTime *string                                          `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	JobId        *string                                          `json:"JobId,omitempty" xml:"JobId,omitempty"`
	Regular      *string                                          `json:"Regular,omitempty" xml:"Regular,omitempty"`
	Snapshots    *ListSnapshotsResponseBodyMediaSnapshotSnapshots `json:"Snapshots,omitempty" xml:"Snapshots,omitempty" type:"Struct"`
	Total        *int64                                           `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListSnapshotsResponseBodyMediaSnapshot) String() string {
	return tea.Prettify(s)
}

func (s ListSnapshotsResponseBodyMediaSnapshot) GoString() string {
	return s.String()
}

func (s *ListSnapshotsResponseBodyMediaSnapshot) SetCreationTime(v string) *ListSnapshotsResponseBodyMediaSnapshot {
	s.CreationTime = &v
	return s
}

func (s *ListSnapshotsResponseBodyMediaSnapshot) SetJobId(v string) *ListSnapshotsResponseBodyMediaSnapshot {
	s.JobId = &v
	return s
}

func (s *ListSnapshotsResponseBodyMediaSnapshot) SetRegular(v string) *ListSnapshotsResponseBodyMediaSnapshot {
	s.Regular = &v
	return s
}

func (s *ListSnapshotsResponseBodyMediaSnapshot) SetSnapshots(v *ListSnapshotsResponseBodyMediaSnapshotSnapshots) *ListSnapshotsResponseBodyMediaSnapshot {
	s.Snapshots = v
	return s
}

func (s *ListSnapshotsResponseBodyMediaSnapshot) SetTotal(v int64) *ListSnapshotsResponseBodyMediaSnapshot {
	s.Total = &v
	return s
}

type ListSnapshotsResponseBodyMediaSnapshotSnapshots struct {
	Snapshot []*ListSnapshotsResponseBodyMediaSnapshotSnapshotsSnapshot `json:"Snapshot,omitempty" xml:"Snapshot,omitempty" type:"Repeated"`
}

func (s ListSnapshotsResponseBodyMediaSnapshotSnapshots) String() string {
	return tea.Prettify(s)
}

func (s ListSnapshotsResponseBodyMediaSnapshotSnapshots) GoString() string {
	return s.String()
}

func (s *ListSnapshotsResponseBodyMediaSnapshotSnapshots) SetSnapshot(v []*ListSnapshotsResponseBodyMediaSnapshotSnapshotsSnapshot) *ListSnapshotsResponseBodyMediaSnapshotSnapshots {
	s.Snapshot = v
	return s
}

type ListSnapshotsResponseBodyMediaSnapshotSnapshotsSnapshot struct {
	Index *int64  `json:"Index,omitempty" xml:"Index,omitempty"`
	Url   *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s ListSnapshotsResponseBodyMediaSnapshotSnapshotsSnapshot) String() string {
	return tea.Prettify(s)
}

func (s ListSnapshotsResponseBodyMediaSnapshotSnapshotsSnapshot) GoString() string {
	return s.String()
}

func (s *ListSnapshotsResponseBodyMediaSnapshotSnapshotsSnapshot) SetIndex(v int64) *ListSnapshotsResponseBodyMediaSnapshotSnapshotsSnapshot {
	s.Index = &v
	return s
}

func (s *ListSnapshotsResponseBodyMediaSnapshotSnapshotsSnapshot) SetUrl(v string) *ListSnapshotsResponseBodyMediaSnapshotSnapshotsSnapshot {
	s.Url = &v
	return s
}

type ListSnapshotsResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListSnapshotsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListSnapshotsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSnapshotsResponse) GoString() string {
	return s.String()
}

func (s *ListSnapshotsResponse) SetHeaders(v map[string]*string) *ListSnapshotsResponse {
	s.Headers = v
	return s
}

func (s *ListSnapshotsResponse) SetBody(v *ListSnapshotsResponseBody) *ListSnapshotsResponse {
	s.Body = v
	return s
}

type ListTranscodeTaskRequest struct {
	EndTime   *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	PageNo    *int32  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize  *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	VideoId   *string `json:"VideoId,omitempty" xml:"VideoId,omitempty"`
}

func (s ListTranscodeTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTranscodeTaskRequest) GoString() string {
	return s.String()
}

func (s *ListTranscodeTaskRequest) SetEndTime(v string) *ListTranscodeTaskRequest {
	s.EndTime = &v
	return s
}

func (s *ListTranscodeTaskRequest) SetPageNo(v int32) *ListTranscodeTaskRequest {
	s.PageNo = &v
	return s
}

func (s *ListTranscodeTaskRequest) SetPageSize(v int32) *ListTranscodeTaskRequest {
	s.PageSize = &v
	return s
}

func (s *ListTranscodeTaskRequest) SetStartTime(v string) *ListTranscodeTaskRequest {
	s.StartTime = &v
	return s
}

func (s *ListTranscodeTaskRequest) SetVideoId(v string) *ListTranscodeTaskRequest {
	s.VideoId = &v
	return s
}

type ListTranscodeTaskResponseBody struct {
	RequestId         *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TranscodeTaskList []*ListTranscodeTaskResponseBodyTranscodeTaskList `json:"TranscodeTaskList,omitempty" xml:"TranscodeTaskList,omitempty" type:"Repeated"`
}

func (s ListTranscodeTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTranscodeTaskResponseBody) GoString() string {
	return s.String()
}

func (s *ListTranscodeTaskResponseBody) SetRequestId(v string) *ListTranscodeTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTranscodeTaskResponseBody) SetTranscodeTaskList(v []*ListTranscodeTaskResponseBodyTranscodeTaskList) *ListTranscodeTaskResponseBody {
	s.TranscodeTaskList = v
	return s
}

type ListTranscodeTaskResponseBodyTranscodeTaskList struct {
	CompleteTime             *string `json:"CompleteTime,omitempty" xml:"CompleteTime,omitempty"`
	CreationTime             *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	TaskStatus               *string `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	TranscodeTaskId          *string `json:"TranscodeTaskId,omitempty" xml:"TranscodeTaskId,omitempty"`
	TranscodeTemplateGroupId *string `json:"TranscodeTemplateGroupId,omitempty" xml:"TranscodeTemplateGroupId,omitempty"`
	Trigger                  *string `json:"Trigger,omitempty" xml:"Trigger,omitempty"`
	VideoId                  *string `json:"VideoId,omitempty" xml:"VideoId,omitempty"`
}

func (s ListTranscodeTaskResponseBodyTranscodeTaskList) String() string {
	return tea.Prettify(s)
}

func (s ListTranscodeTaskResponseBodyTranscodeTaskList) GoString() string {
	return s.String()
}

func (s *ListTranscodeTaskResponseBodyTranscodeTaskList) SetCompleteTime(v string) *ListTranscodeTaskResponseBodyTranscodeTaskList {
	s.CompleteTime = &v
	return s
}

func (s *ListTranscodeTaskResponseBodyTranscodeTaskList) SetCreationTime(v string) *ListTranscodeTaskResponseBodyTranscodeTaskList {
	s.CreationTime = &v
	return s
}

func (s *ListTranscodeTaskResponseBodyTranscodeTaskList) SetTaskStatus(v string) *ListTranscodeTaskResponseBodyTranscodeTaskList {
	s.TaskStatus = &v
	return s
}

func (s *ListTranscodeTaskResponseBodyTranscodeTaskList) SetTranscodeTaskId(v string) *ListTranscodeTaskResponseBodyTranscodeTaskList {
	s.TranscodeTaskId = &v
	return s
}

func (s *ListTranscodeTaskResponseBodyTranscodeTaskList) SetTranscodeTemplateGroupId(v string) *ListTranscodeTaskResponseBodyTranscodeTaskList {
	s.TranscodeTemplateGroupId = &v
	return s
}

func (s *ListTranscodeTaskResponseBodyTranscodeTaskList) SetTrigger(v string) *ListTranscodeTaskResponseBodyTranscodeTaskList {
	s.Trigger = &v
	return s
}

func (s *ListTranscodeTaskResponseBodyTranscodeTaskList) SetVideoId(v string) *ListTranscodeTaskResponseBodyTranscodeTaskList {
	s.VideoId = &v
	return s
}

type ListTranscodeTaskResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListTranscodeTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListTranscodeTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTranscodeTaskResponse) GoString() string {
	return s.String()
}

func (s *ListTranscodeTaskResponse) SetHeaders(v map[string]*string) *ListTranscodeTaskResponse {
	s.Headers = v
	return s
}

func (s *ListTranscodeTaskResponse) SetBody(v *ListTranscodeTaskResponseBody) *ListTranscodeTaskResponse {
	s.Body = v
	return s
}

type ListTranscodeTemplateGroupRequest struct {
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
}

func (s ListTranscodeTemplateGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTranscodeTemplateGroupRequest) GoString() string {
	return s.String()
}

func (s *ListTranscodeTemplateGroupRequest) SetAppId(v string) *ListTranscodeTemplateGroupRequest {
	s.AppId = &v
	return s
}

type ListTranscodeTemplateGroupResponseBody struct {
	RequestId                  *string                                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TranscodeTemplateGroupList []*ListTranscodeTemplateGroupResponseBodyTranscodeTemplateGroupList `json:"TranscodeTemplateGroupList,omitempty" xml:"TranscodeTemplateGroupList,omitempty" type:"Repeated"`
}

func (s ListTranscodeTemplateGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTranscodeTemplateGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ListTranscodeTemplateGroupResponseBody) SetRequestId(v string) *ListTranscodeTemplateGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTranscodeTemplateGroupResponseBody) SetTranscodeTemplateGroupList(v []*ListTranscodeTemplateGroupResponseBodyTranscodeTemplateGroupList) *ListTranscodeTemplateGroupResponseBody {
	s.TranscodeTemplateGroupList = v
	return s
}

type ListTranscodeTemplateGroupResponseBodyTranscodeTemplateGroupList struct {
	AppId                    *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	CreationTime             *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	IsDefault                *string `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	Locked                   *string `json:"Locked,omitempty" xml:"Locked,omitempty"`
	ModifyTime               *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	Name                     *string `json:"Name,omitempty" xml:"Name,omitempty"`
	TranscodeTemplateGroupId *string `json:"TranscodeTemplateGroupId,omitempty" xml:"TranscodeTemplateGroupId,omitempty"`
}

func (s ListTranscodeTemplateGroupResponseBodyTranscodeTemplateGroupList) String() string {
	return tea.Prettify(s)
}

func (s ListTranscodeTemplateGroupResponseBodyTranscodeTemplateGroupList) GoString() string {
	return s.String()
}

func (s *ListTranscodeTemplateGroupResponseBodyTranscodeTemplateGroupList) SetAppId(v string) *ListTranscodeTemplateGroupResponseBodyTranscodeTemplateGroupList {
	s.AppId = &v
	return s
}

func (s *ListTranscodeTemplateGroupResponseBodyTranscodeTemplateGroupList) SetCreationTime(v string) *ListTranscodeTemplateGroupResponseBodyTranscodeTemplateGroupList {
	s.CreationTime = &v
	return s
}

func (s *ListTranscodeTemplateGroupResponseBodyTranscodeTemplateGroupList) SetIsDefault(v string) *ListTranscodeTemplateGroupResponseBodyTranscodeTemplateGroupList {
	s.IsDefault = &v
	return s
}

func (s *ListTranscodeTemplateGroupResponseBodyTranscodeTemplateGroupList) SetLocked(v string) *ListTranscodeTemplateGroupResponseBodyTranscodeTemplateGroupList {
	s.Locked = &v
	return s
}

func (s *ListTranscodeTemplateGroupResponseBodyTranscodeTemplateGroupList) SetModifyTime(v string) *ListTranscodeTemplateGroupResponseBodyTranscodeTemplateGroupList {
	s.ModifyTime = &v
	return s
}

func (s *ListTranscodeTemplateGroupResponseBodyTranscodeTemplateGroupList) SetName(v string) *ListTranscodeTemplateGroupResponseBodyTranscodeTemplateGroupList {
	s.Name = &v
	return s
}

func (s *ListTranscodeTemplateGroupResponseBodyTranscodeTemplateGroupList) SetTranscodeTemplateGroupId(v string) *ListTranscodeTemplateGroupResponseBodyTranscodeTemplateGroupList {
	s.TranscodeTemplateGroupId = &v
	return s
}

type ListTranscodeTemplateGroupResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListTranscodeTemplateGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListTranscodeTemplateGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTranscodeTemplateGroupResponse) GoString() string {
	return s.String()
}

func (s *ListTranscodeTemplateGroupResponse) SetHeaders(v map[string]*string) *ListTranscodeTemplateGroupResponse {
	s.Headers = v
	return s
}

func (s *ListTranscodeTemplateGroupResponse) SetBody(v *ListTranscodeTemplateGroupResponseBody) *ListTranscodeTemplateGroupResponse {
	s.Body = v
	return s
}

type ListVodRealtimeLogDeliveryDomainsRequest struct {
	Logstore *string `json:"Logstore,omitempty" xml:"Logstore,omitempty"`
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	Project  *string `json:"Project,omitempty" xml:"Project,omitempty"`
	Region   *string `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s ListVodRealtimeLogDeliveryDomainsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListVodRealtimeLogDeliveryDomainsRequest) GoString() string {
	return s.String()
}

func (s *ListVodRealtimeLogDeliveryDomainsRequest) SetLogstore(v string) *ListVodRealtimeLogDeliveryDomainsRequest {
	s.Logstore = &v
	return s
}

func (s *ListVodRealtimeLogDeliveryDomainsRequest) SetOwnerId(v int64) *ListVodRealtimeLogDeliveryDomainsRequest {
	s.OwnerId = &v
	return s
}

func (s *ListVodRealtimeLogDeliveryDomainsRequest) SetProject(v string) *ListVodRealtimeLogDeliveryDomainsRequest {
	s.Project = &v
	return s
}

func (s *ListVodRealtimeLogDeliveryDomainsRequest) SetRegion(v string) *ListVodRealtimeLogDeliveryDomainsRequest {
	s.Region = &v
	return s
}

type ListVodRealtimeLogDeliveryDomainsResponseBody struct {
	Content   *ListVodRealtimeLogDeliveryDomainsResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	RequestId *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListVodRealtimeLogDeliveryDomainsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListVodRealtimeLogDeliveryDomainsResponseBody) GoString() string {
	return s.String()
}

func (s *ListVodRealtimeLogDeliveryDomainsResponseBody) SetContent(v *ListVodRealtimeLogDeliveryDomainsResponseBodyContent) *ListVodRealtimeLogDeliveryDomainsResponseBody {
	s.Content = v
	return s
}

func (s *ListVodRealtimeLogDeliveryDomainsResponseBody) SetRequestId(v string) *ListVodRealtimeLogDeliveryDomainsResponseBody {
	s.RequestId = &v
	return s
}

type ListVodRealtimeLogDeliveryDomainsResponseBodyContent struct {
	Domains []*ListVodRealtimeLogDeliveryDomainsResponseBodyContentDomains `json:"Domains,omitempty" xml:"Domains,omitempty" type:"Repeated"`
}

func (s ListVodRealtimeLogDeliveryDomainsResponseBodyContent) String() string {
	return tea.Prettify(s)
}

func (s ListVodRealtimeLogDeliveryDomainsResponseBodyContent) GoString() string {
	return s.String()
}

func (s *ListVodRealtimeLogDeliveryDomainsResponseBodyContent) SetDomains(v []*ListVodRealtimeLogDeliveryDomainsResponseBodyContentDomains) *ListVodRealtimeLogDeliveryDomainsResponseBodyContent {
	s.Domains = v
	return s
}

type ListVodRealtimeLogDeliveryDomainsResponseBodyContentDomains struct {
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	Status     *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListVodRealtimeLogDeliveryDomainsResponseBodyContentDomains) String() string {
	return tea.Prettify(s)
}

func (s ListVodRealtimeLogDeliveryDomainsResponseBodyContentDomains) GoString() string {
	return s.String()
}

func (s *ListVodRealtimeLogDeliveryDomainsResponseBodyContentDomains) SetDomainName(v string) *ListVodRealtimeLogDeliveryDomainsResponseBodyContentDomains {
	s.DomainName = &v
	return s
}

func (s *ListVodRealtimeLogDeliveryDomainsResponseBodyContentDomains) SetStatus(v string) *ListVodRealtimeLogDeliveryDomainsResponseBodyContentDomains {
	s.Status = &v
	return s
}

type ListVodRealtimeLogDeliveryDomainsResponse struct {
	Headers map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListVodRealtimeLogDeliveryDomainsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListVodRealtimeLogDeliveryDomainsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListVodRealtimeLogDeliveryDomainsResponse) GoString() string {
	return s.String()
}

func (s *ListVodRealtimeLogDeliveryDomainsResponse) SetHeaders(v map[string]*string) *ListVodRealtimeLogDeliveryDomainsResponse {
	s.Headers = v
	return s
}

func (s *ListVodRealtimeLogDeliveryDomainsResponse) SetBody(v *ListVodRealtimeLogDeliveryDomainsResponseBody) *ListVodRealtimeLogDeliveryDomainsResponse {
	s.Body = v
	return s
}

type ListVodRealtimeLogDeliveryInfosRequest struct {
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s ListVodRealtimeLogDeliveryInfosRequest) String() string {
	return tea.Prettify(s)
}

func (s ListVodRealtimeLogDeliveryInfosRequest) GoString() string {
	return s.String()
}

func (s *ListVodRealtimeLogDeliveryInfosRequest) SetOwnerId(v int64) *ListVodRealtimeLogDeliveryInfosRequest {
	s.OwnerId = &v
	return s
}

type ListVodRealtimeLogDeliveryInfosResponseBody struct {
	Content   *ListVodRealtimeLogDeliveryInfosResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	RequestId *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListVodRealtimeLogDeliveryInfosResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListVodRealtimeLogDeliveryInfosResponseBody) GoString() string {
	return s.String()
}

func (s *ListVodRealtimeLogDeliveryInfosResponseBody) SetContent(v *ListVodRealtimeLogDeliveryInfosResponseBodyContent) *ListVodRealtimeLogDeliveryInfosResponseBody {
	s.Content = v
	return s
}

func (s *ListVodRealtimeLogDeliveryInfosResponseBody) SetRequestId(v string) *ListVodRealtimeLogDeliveryInfosResponseBody {
	s.RequestId = &v
	return s
}

type ListVodRealtimeLogDeliveryInfosResponseBodyContent struct {
	RealtimeLogDeliveryInfos []*ListVodRealtimeLogDeliveryInfosResponseBodyContentRealtimeLogDeliveryInfos `json:"RealtimeLogDeliveryInfos,omitempty" xml:"RealtimeLogDeliveryInfos,omitempty" type:"Repeated"`
}

func (s ListVodRealtimeLogDeliveryInfosResponseBodyContent) String() string {
	return tea.Prettify(s)
}

func (s ListVodRealtimeLogDeliveryInfosResponseBodyContent) GoString() string {
	return s.String()
}

func (s *ListVodRealtimeLogDeliveryInfosResponseBodyContent) SetRealtimeLogDeliveryInfos(v []*ListVodRealtimeLogDeliveryInfosResponseBodyContentRealtimeLogDeliveryInfos) *ListVodRealtimeLogDeliveryInfosResponseBodyContent {
	s.RealtimeLogDeliveryInfos = v
	return s
}

type ListVodRealtimeLogDeliveryInfosResponseBodyContentRealtimeLogDeliveryInfos struct {
	Logstore *string `json:"Logstore,omitempty" xml:"Logstore,omitempty"`
	Project  *string `json:"Project,omitempty" xml:"Project,omitempty"`
	Region   *string `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s ListVodRealtimeLogDeliveryInfosResponseBodyContentRealtimeLogDeliveryInfos) String() string {
	return tea.Prettify(s)
}

func (s ListVodRealtimeLogDeliveryInfosResponseBodyContentRealtimeLogDeliveryInfos) GoString() string {
	return s.String()
}

func (s *ListVodRealtimeLogDeliveryInfosResponseBodyContentRealtimeLogDeliveryInfos) SetLogstore(v string) *ListVodRealtimeLogDeliveryInfosResponseBodyContentRealtimeLogDeliveryInfos {
	s.Logstore = &v
	return s
}

func (s *ListVodRealtimeLogDeliveryInfosResponseBodyContentRealtimeLogDeliveryInfos) SetProject(v string) *ListVodRealtimeLogDeliveryInfosResponseBodyContentRealtimeLogDeliveryInfos {
	s.Project = &v
	return s
}

func (s *ListVodRealtimeLogDeliveryInfosResponseBodyContentRealtimeLogDeliveryInfos) SetRegion(v string) *ListVodRealtimeLogDeliveryInfosResponseBodyContentRealtimeLogDeliveryInfos {
	s.Region = &v
	return s
}

type ListVodRealtimeLogDeliveryInfosResponse struct {
	Headers map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListVodRealtimeLogDeliveryInfosResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListVodRealtimeLogDeliveryInfosResponse) String() string {
	return tea.Prettify(s)
}

func (s ListVodRealtimeLogDeliveryInfosResponse) GoString() string {
	return s.String()
}

func (s *ListVodRealtimeLogDeliveryInfosResponse) SetHeaders(v map[string]*string) *ListVodRealtimeLogDeliveryInfosResponse {
	s.Headers = v
	return s
}

func (s *ListVodRealtimeLogDeliveryInfosResponse) SetBody(v *ListVodRealtimeLogDeliveryInfosResponseBody) *ListVodRealtimeLogDeliveryInfosResponse {
	s.Body = v
	return s
}

type ListVodTemplateRequest struct {
	AppId        *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	TemplateType *string `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
}

func (s ListVodTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s ListVodTemplateRequest) GoString() string {
	return s.String()
}

func (s *ListVodTemplateRequest) SetAppId(v string) *ListVodTemplateRequest {
	s.AppId = &v
	return s
}

func (s *ListVodTemplateRequest) SetTemplateType(v string) *ListVodTemplateRequest {
	s.TemplateType = &v
	return s
}

type ListVodTemplateResponseBody struct {
	RequestId           *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	VodTemplateInfoList []*ListVodTemplateResponseBodyVodTemplateInfoList `json:"VodTemplateInfoList,omitempty" xml:"VodTemplateInfoList,omitempty" type:"Repeated"`
}

func (s ListVodTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListVodTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *ListVodTemplateResponseBody) SetRequestId(v string) *ListVodTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListVodTemplateResponseBody) SetVodTemplateInfoList(v []*ListVodTemplateResponseBodyVodTemplateInfoList) *ListVodTemplateResponseBody {
	s.VodTemplateInfoList = v
	return s
}

type ListVodTemplateResponseBodyVodTemplateInfoList struct {
	AppId          *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	CreationTime   *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	IsDefault      *string `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	ModifyTime     *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	Name           *string `json:"Name,omitempty" xml:"Name,omitempty"`
	TemplateConfig *string `json:"TemplateConfig,omitempty" xml:"TemplateConfig,omitempty"`
	TemplateType   *string `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
	VodTemplateId  *string `json:"VodTemplateId,omitempty" xml:"VodTemplateId,omitempty"`
}

func (s ListVodTemplateResponseBodyVodTemplateInfoList) String() string {
	return tea.Prettify(s)
}

func (s ListVodTemplateResponseBodyVodTemplateInfoList) GoString() string {
	return s.String()
}

func (s *ListVodTemplateResponseBodyVodTemplateInfoList) SetAppId(v string) *ListVodTemplateResponseBodyVodTemplateInfoList {
	s.AppId = &v
	return s
}

func (s *ListVodTemplateResponseBodyVodTemplateInfoList) SetCreationTime(v string) *ListVodTemplateResponseBodyVodTemplateInfoList {
	s.CreationTime = &v
	return s
}

func (s *ListVodTemplateResponseBodyVodTemplateInfoList) SetIsDefault(v string) *ListVodTemplateResponseBodyVodTemplateInfoList {
	s.IsDefault = &v
	return s
}

func (s *ListVodTemplateResponseBodyVodTemplateInfoList) SetModifyTime(v string) *ListVodTemplateResponseBodyVodTemplateInfoList {
	s.ModifyTime = &v
	return s
}

func (s *ListVodTemplateResponseBodyVodTemplateInfoList) SetName(v string) *ListVodTemplateResponseBodyVodTemplateInfoList {
	s.Name = &v
	return s
}

func (s *ListVodTemplateResponseBodyVodTemplateInfoList) SetTemplateConfig(v string) *ListVodTemplateResponseBodyVodTemplateInfoList {
	s.TemplateConfig = &v
	return s
}

func (s *ListVodTemplateResponseBodyVodTemplateInfoList) SetTemplateType(v string) *ListVodTemplateResponseBodyVodTemplateInfoList {
	s.TemplateType = &v
	return s
}

func (s *ListVodTemplateResponseBodyVodTemplateInfoList) SetVodTemplateId(v string) *ListVodTemplateResponseBodyVodTemplateInfoList {
	s.VodTemplateId = &v
	return s
}

type ListVodTemplateResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListVodTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListVodTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s ListVodTemplateResponse) GoString() string {
	return s.String()
}

func (s *ListVodTemplateResponse) SetHeaders(v map[string]*string) *ListVodTemplateResponse {
	s.Headers = v
	return s
}

func (s *ListVodTemplateResponse) SetBody(v *ListVodTemplateResponseBody) *ListVodTemplateResponse {
	s.Body = v
	return s
}

type ListWatermarkRequest struct {
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
}

func (s ListWatermarkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListWatermarkRequest) GoString() string {
	return s.String()
}

func (s *ListWatermarkRequest) SetAppId(v string) *ListWatermarkRequest {
	s.AppId = &v
	return s
}

type ListWatermarkResponseBody struct {
	RequestId      *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	WatermarkInfos []*ListWatermarkResponseBodyWatermarkInfos `json:"WatermarkInfos,omitempty" xml:"WatermarkInfos,omitempty" type:"Repeated"`
}

func (s ListWatermarkResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListWatermarkResponseBody) GoString() string {
	return s.String()
}

func (s *ListWatermarkResponseBody) SetRequestId(v string) *ListWatermarkResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListWatermarkResponseBody) SetWatermarkInfos(v []*ListWatermarkResponseBodyWatermarkInfos) *ListWatermarkResponseBody {
	s.WatermarkInfos = v
	return s
}

type ListWatermarkResponseBodyWatermarkInfos struct {
	AppId           *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	CreationTime    *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	FileUrl         *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	IsDefault       *string `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	Name            *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Type            *string `json:"Type,omitempty" xml:"Type,omitempty"`
	WatermarkConfig *string `json:"WatermarkConfig,omitempty" xml:"WatermarkConfig,omitempty"`
	WatermarkId     *string `json:"WatermarkId,omitempty" xml:"WatermarkId,omitempty"`
}

func (s ListWatermarkResponseBodyWatermarkInfos) String() string {
	return tea.Prettify(s)
}

func (s ListWatermarkResponseBodyWatermarkInfos) GoString() string {
	return s.String()
}

func (s *ListWatermarkResponseBodyWatermarkInfos) SetAppId(v string) *ListWatermarkResponseBodyWatermarkInfos {
	s.AppId = &v
	return s
}

func (s *ListWatermarkResponseBodyWatermarkInfos) SetCreationTime(v string) *ListWatermarkResponseBodyWatermarkInfos {
	s.CreationTime = &v
	return s
}

func (s *ListWatermarkResponseBodyWatermarkInfos) SetFileUrl(v string) *ListWatermarkResponseBodyWatermarkInfos {
	s.FileUrl = &v
	return s
}

func (s *ListWatermarkResponseBodyWatermarkInfos) SetIsDefault(v string) *ListWatermarkResponseBodyWatermarkInfos {
	s.IsDefault = &v
	return s
}

func (s *ListWatermarkResponseBodyWatermarkInfos) SetName(v string) *ListWatermarkResponseBodyWatermarkInfos {
	s.Name = &v
	return s
}

func (s *ListWatermarkResponseBodyWatermarkInfos) SetType(v string) *ListWatermarkResponseBodyWatermarkInfos {
	s.Type = &v
	return s
}

func (s *ListWatermarkResponseBodyWatermarkInfos) SetWatermarkConfig(v string) *ListWatermarkResponseBodyWatermarkInfos {
	s.WatermarkConfig = &v
	return s
}

func (s *ListWatermarkResponseBodyWatermarkInfos) SetWatermarkId(v string) *ListWatermarkResponseBodyWatermarkInfos {
	s.WatermarkId = &v
	return s
}

type ListWatermarkResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListWatermarkResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListWatermarkResponse) String() string {
	return tea.Prettify(s)
}

func (s ListWatermarkResponse) GoString() string {
	return s.String()
}

func (s *ListWatermarkResponse) SetHeaders(v map[string]*string) *ListWatermarkResponse {
	s.Headers = v
	return s
}

func (s *ListWatermarkResponse) SetBody(v *ListWatermarkResponseBody) *ListWatermarkResponse {
	s.Body = v
	return s
}

type MoveAppResourceRequest struct {
	ResourceIds  *string `json:"ResourceIds,omitempty" xml:"ResourceIds,omitempty"`
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	TargetAppId  *string `json:"TargetAppId,omitempty" xml:"TargetAppId,omitempty"`
}

func (s MoveAppResourceRequest) String() string {
	return tea.Prettify(s)
}

func (s MoveAppResourceRequest) GoString() string {
	return s.String()
}

func (s *MoveAppResourceRequest) SetResourceIds(v string) *MoveAppResourceRequest {
	s.ResourceIds = &v
	return s
}

func (s *MoveAppResourceRequest) SetResourceType(v string) *MoveAppResourceRequest {
	s.ResourceType = &v
	return s
}

func (s *MoveAppResourceRequest) SetTargetAppId(v string) *MoveAppResourceRequest {
	s.TargetAppId = &v
	return s
}

type MoveAppResourceResponseBody struct {
	FailedResourceIds   []*string `json:"FailedResourceIds,omitempty" xml:"FailedResourceIds,omitempty" type:"Repeated"`
	NonExistResourceIds []*string `json:"NonExistResourceIds,omitempty" xml:"NonExistResourceIds,omitempty" type:"Repeated"`
	RequestId           *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s MoveAppResourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s MoveAppResourceResponseBody) GoString() string {
	return s.String()
}

func (s *MoveAppResourceResponseBody) SetFailedResourceIds(v []*string) *MoveAppResourceResponseBody {
	s.FailedResourceIds = v
	return s
}

func (s *MoveAppResourceResponseBody) SetNonExistResourceIds(v []*string) *MoveAppResourceResponseBody {
	s.NonExistResourceIds = v
	return s
}

func (s *MoveAppResourceResponseBody) SetRequestId(v string) *MoveAppResourceResponseBody {
	s.RequestId = &v
	return s
}

type MoveAppResourceResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *MoveAppResourceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s MoveAppResourceResponse) String() string {
	return tea.Prettify(s)
}

func (s MoveAppResourceResponse) GoString() string {
	return s.String()
}

func (s *MoveAppResourceResponse) SetHeaders(v map[string]*string) *MoveAppResourceResponse {
	s.Headers = v
	return s
}

func (s *MoveAppResourceResponse) SetBody(v *MoveAppResourceResponseBody) *MoveAppResourceResponse {
	s.Body = v
	return s
}

type PreloadVodObjectCachesRequest struct {
	ObjectPath    *string `json:"ObjectPath,omitempty" xml:"ObjectPath,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s PreloadVodObjectCachesRequest) String() string {
	return tea.Prettify(s)
}

func (s PreloadVodObjectCachesRequest) GoString() string {
	return s.String()
}

func (s *PreloadVodObjectCachesRequest) SetObjectPath(v string) *PreloadVodObjectCachesRequest {
	s.ObjectPath = &v
	return s
}

func (s *PreloadVodObjectCachesRequest) SetOwnerId(v int64) *PreloadVodObjectCachesRequest {
	s.OwnerId = &v
	return s
}

func (s *PreloadVodObjectCachesRequest) SetSecurityToken(v string) *PreloadVodObjectCachesRequest {
	s.SecurityToken = &v
	return s
}

type PreloadVodObjectCachesResponseBody struct {
	PreloadTaskId *string `json:"PreloadTaskId,omitempty" xml:"PreloadTaskId,omitempty"`
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s PreloadVodObjectCachesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PreloadVodObjectCachesResponseBody) GoString() string {
	return s.String()
}

func (s *PreloadVodObjectCachesResponseBody) SetPreloadTaskId(v string) *PreloadVodObjectCachesResponseBody {
	s.PreloadTaskId = &v
	return s
}

func (s *PreloadVodObjectCachesResponseBody) SetRequestId(v string) *PreloadVodObjectCachesResponseBody {
	s.RequestId = &v
	return s
}

type PreloadVodObjectCachesResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *PreloadVodObjectCachesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PreloadVodObjectCachesResponse) String() string {
	return tea.Prettify(s)
}

func (s PreloadVodObjectCachesResponse) GoString() string {
	return s.String()
}

func (s *PreloadVodObjectCachesResponse) SetHeaders(v map[string]*string) *PreloadVodObjectCachesResponse {
	s.Headers = v
	return s
}

func (s *PreloadVodObjectCachesResponse) SetBody(v *PreloadVodObjectCachesResponseBody) *PreloadVodObjectCachesResponse {
	s.Body = v
	return s
}

type ProduceEditingProjectVideoRequest struct {
	CoverURL             *string `json:"CoverURL,omitempty" xml:"CoverURL,omitempty"`
	Description          *string `json:"Description,omitempty" xml:"Description,omitempty"`
	MediaMetadata        *string `json:"MediaMetadata,omitempty" xml:"MediaMetadata,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ProduceConfig        *string `json:"ProduceConfig,omitempty" xml:"ProduceConfig,omitempty"`
	ProjectId            *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	Timeline             *string `json:"Timeline,omitempty" xml:"Timeline,omitempty"`
	Title                *string `json:"Title,omitempty" xml:"Title,omitempty"`
	UserData             *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s ProduceEditingProjectVideoRequest) String() string {
	return tea.Prettify(s)
}

func (s ProduceEditingProjectVideoRequest) GoString() string {
	return s.String()
}

func (s *ProduceEditingProjectVideoRequest) SetCoverURL(v string) *ProduceEditingProjectVideoRequest {
	s.CoverURL = &v
	return s
}

func (s *ProduceEditingProjectVideoRequest) SetDescription(v string) *ProduceEditingProjectVideoRequest {
	s.Description = &v
	return s
}

func (s *ProduceEditingProjectVideoRequest) SetMediaMetadata(v string) *ProduceEditingProjectVideoRequest {
	s.MediaMetadata = &v
	return s
}

func (s *ProduceEditingProjectVideoRequest) SetOwnerId(v int64) *ProduceEditingProjectVideoRequest {
	s.OwnerId = &v
	return s
}

func (s *ProduceEditingProjectVideoRequest) SetProduceConfig(v string) *ProduceEditingProjectVideoRequest {
	s.ProduceConfig = &v
	return s
}

func (s *ProduceEditingProjectVideoRequest) SetProjectId(v string) *ProduceEditingProjectVideoRequest {
	s.ProjectId = &v
	return s
}

func (s *ProduceEditingProjectVideoRequest) SetResourceOwnerAccount(v string) *ProduceEditingProjectVideoRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ProduceEditingProjectVideoRequest) SetResourceOwnerId(v int64) *ProduceEditingProjectVideoRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ProduceEditingProjectVideoRequest) SetTimeline(v string) *ProduceEditingProjectVideoRequest {
	s.Timeline = &v
	return s
}

func (s *ProduceEditingProjectVideoRequest) SetTitle(v string) *ProduceEditingProjectVideoRequest {
	s.Title = &v
	return s
}

func (s *ProduceEditingProjectVideoRequest) SetUserData(v string) *ProduceEditingProjectVideoRequest {
	s.UserData = &v
	return s
}

type ProduceEditingProjectVideoResponseBody struct {
	MediaId   *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ProduceEditingProjectVideoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ProduceEditingProjectVideoResponseBody) GoString() string {
	return s.String()
}

func (s *ProduceEditingProjectVideoResponseBody) SetMediaId(v string) *ProduceEditingProjectVideoResponseBody {
	s.MediaId = &v
	return s
}

func (s *ProduceEditingProjectVideoResponseBody) SetProjectId(v string) *ProduceEditingProjectVideoResponseBody {
	s.ProjectId = &v
	return s
}

func (s *ProduceEditingProjectVideoResponseBody) SetRequestId(v string) *ProduceEditingProjectVideoResponseBody {
	s.RequestId = &v
	return s
}

type ProduceEditingProjectVideoResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ProduceEditingProjectVideoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ProduceEditingProjectVideoResponse) String() string {
	return tea.Prettify(s)
}

func (s ProduceEditingProjectVideoResponse) GoString() string {
	return s.String()
}

func (s *ProduceEditingProjectVideoResponse) SetHeaders(v map[string]*string) *ProduceEditingProjectVideoResponse {
	s.Headers = v
	return s
}

func (s *ProduceEditingProjectVideoResponse) SetBody(v *ProduceEditingProjectVideoResponseBody) *ProduceEditingProjectVideoResponse {
	s.Body = v
	return s
}

type RefreshUploadVideoRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	VideoId              *string `json:"VideoId,omitempty" xml:"VideoId,omitempty"`
}

func (s RefreshUploadVideoRequest) String() string {
	return tea.Prettify(s)
}

func (s RefreshUploadVideoRequest) GoString() string {
	return s.String()
}

func (s *RefreshUploadVideoRequest) SetOwnerId(v int64) *RefreshUploadVideoRequest {
	s.OwnerId = &v
	return s
}

func (s *RefreshUploadVideoRequest) SetResourceOwnerAccount(v string) *RefreshUploadVideoRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *RefreshUploadVideoRequest) SetResourceOwnerId(v int64) *RefreshUploadVideoRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *RefreshUploadVideoRequest) SetVideoId(v string) *RefreshUploadVideoRequest {
	s.VideoId = &v
	return s
}

type RefreshUploadVideoResponseBody struct {
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	UploadAddress *string `json:"UploadAddress,omitempty" xml:"UploadAddress,omitempty"`
	UploadAuth    *string `json:"UploadAuth,omitempty" xml:"UploadAuth,omitempty"`
	VideoId       *string `json:"VideoId,omitempty" xml:"VideoId,omitempty"`
}

func (s RefreshUploadVideoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RefreshUploadVideoResponseBody) GoString() string {
	return s.String()
}

func (s *RefreshUploadVideoResponseBody) SetRequestId(v string) *RefreshUploadVideoResponseBody {
	s.RequestId = &v
	return s
}

func (s *RefreshUploadVideoResponseBody) SetUploadAddress(v string) *RefreshUploadVideoResponseBody {
	s.UploadAddress = &v
	return s
}

func (s *RefreshUploadVideoResponseBody) SetUploadAuth(v string) *RefreshUploadVideoResponseBody {
	s.UploadAuth = &v
	return s
}

func (s *RefreshUploadVideoResponseBody) SetVideoId(v string) *RefreshUploadVideoResponseBody {
	s.VideoId = &v
	return s
}

type RefreshUploadVideoResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RefreshUploadVideoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RefreshUploadVideoResponse) String() string {
	return tea.Prettify(s)
}

func (s RefreshUploadVideoResponse) GoString() string {
	return s.String()
}

func (s *RefreshUploadVideoResponse) SetHeaders(v map[string]*string) *RefreshUploadVideoResponse {
	s.Headers = v
	return s
}

func (s *RefreshUploadVideoResponse) SetBody(v *RefreshUploadVideoResponseBody) *RefreshUploadVideoResponse {
	s.Body = v
	return s
}

type RefreshVodObjectCachesRequest struct {
	ObjectPath    *string `json:"ObjectPath,omitempty" xml:"ObjectPath,omitempty"`
	ObjectType    *string `json:"ObjectType,omitempty" xml:"ObjectType,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s RefreshVodObjectCachesRequest) String() string {
	return tea.Prettify(s)
}

func (s RefreshVodObjectCachesRequest) GoString() string {
	return s.String()
}

func (s *RefreshVodObjectCachesRequest) SetObjectPath(v string) *RefreshVodObjectCachesRequest {
	s.ObjectPath = &v
	return s
}

func (s *RefreshVodObjectCachesRequest) SetObjectType(v string) *RefreshVodObjectCachesRequest {
	s.ObjectType = &v
	return s
}

func (s *RefreshVodObjectCachesRequest) SetOwnerId(v int64) *RefreshVodObjectCachesRequest {
	s.OwnerId = &v
	return s
}

func (s *RefreshVodObjectCachesRequest) SetSecurityToken(v string) *RefreshVodObjectCachesRequest {
	s.SecurityToken = &v
	return s
}

type RefreshVodObjectCachesResponseBody struct {
	RefreshTaskId *string `json:"RefreshTaskId,omitempty" xml:"RefreshTaskId,omitempty"`
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RefreshVodObjectCachesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RefreshVodObjectCachesResponseBody) GoString() string {
	return s.String()
}

func (s *RefreshVodObjectCachesResponseBody) SetRefreshTaskId(v string) *RefreshVodObjectCachesResponseBody {
	s.RefreshTaskId = &v
	return s
}

func (s *RefreshVodObjectCachesResponseBody) SetRequestId(v string) *RefreshVodObjectCachesResponseBody {
	s.RequestId = &v
	return s
}

type RefreshVodObjectCachesResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RefreshVodObjectCachesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RefreshVodObjectCachesResponse) String() string {
	return tea.Prettify(s)
}

func (s RefreshVodObjectCachesResponse) GoString() string {
	return s.String()
}

func (s *RefreshVodObjectCachesResponse) SetHeaders(v map[string]*string) *RefreshVodObjectCachesResponse {
	s.Headers = v
	return s
}

func (s *RefreshVodObjectCachesResponse) SetBody(v *RefreshVodObjectCachesResponseBody) *RefreshVodObjectCachesResponse {
	s.Body = v
	return s
}

type RegisterMediaRequest struct {
	RegisterMetadatas *string `json:"RegisterMetadatas,omitempty" xml:"RegisterMetadatas,omitempty"`
	TemplateGroupId   *string `json:"TemplateGroupId,omitempty" xml:"TemplateGroupId,omitempty"`
	UserData          *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
	WorkflowId        *string `json:"WorkflowId,omitempty" xml:"WorkflowId,omitempty"`
}

func (s RegisterMediaRequest) String() string {
	return tea.Prettify(s)
}

func (s RegisterMediaRequest) GoString() string {
	return s.String()
}

func (s *RegisterMediaRequest) SetRegisterMetadatas(v string) *RegisterMediaRequest {
	s.RegisterMetadatas = &v
	return s
}

func (s *RegisterMediaRequest) SetTemplateGroupId(v string) *RegisterMediaRequest {
	s.TemplateGroupId = &v
	return s
}

func (s *RegisterMediaRequest) SetUserData(v string) *RegisterMediaRequest {
	s.UserData = &v
	return s
}

func (s *RegisterMediaRequest) SetWorkflowId(v string) *RegisterMediaRequest {
	s.WorkflowId = &v
	return s
}

type RegisterMediaResponseBody struct {
	FailedFileURLs      []*string                                       `json:"FailedFileURLs,omitempty" xml:"FailedFileURLs,omitempty" type:"Repeated"`
	RegisteredMediaList []*RegisterMediaResponseBodyRegisteredMediaList `json:"RegisteredMediaList,omitempty" xml:"RegisteredMediaList,omitempty" type:"Repeated"`
	RequestId           *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RegisterMediaResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RegisterMediaResponseBody) GoString() string {
	return s.String()
}

func (s *RegisterMediaResponseBody) SetFailedFileURLs(v []*string) *RegisterMediaResponseBody {
	s.FailedFileURLs = v
	return s
}

func (s *RegisterMediaResponseBody) SetRegisteredMediaList(v []*RegisterMediaResponseBodyRegisteredMediaList) *RegisterMediaResponseBody {
	s.RegisteredMediaList = v
	return s
}

func (s *RegisterMediaResponseBody) SetRequestId(v string) *RegisterMediaResponseBody {
	s.RequestId = &v
	return s
}

type RegisterMediaResponseBodyRegisteredMediaList struct {
	FileURL     *string `json:"FileURL,omitempty" xml:"FileURL,omitempty"`
	MediaId     *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	NewRegister *bool   `json:"NewRegister,omitempty" xml:"NewRegister,omitempty"`
}

func (s RegisterMediaResponseBodyRegisteredMediaList) String() string {
	return tea.Prettify(s)
}

func (s RegisterMediaResponseBodyRegisteredMediaList) GoString() string {
	return s.String()
}

func (s *RegisterMediaResponseBodyRegisteredMediaList) SetFileURL(v string) *RegisterMediaResponseBodyRegisteredMediaList {
	s.FileURL = &v
	return s
}

func (s *RegisterMediaResponseBodyRegisteredMediaList) SetMediaId(v string) *RegisterMediaResponseBodyRegisteredMediaList {
	s.MediaId = &v
	return s
}

func (s *RegisterMediaResponseBodyRegisteredMediaList) SetNewRegister(v bool) *RegisterMediaResponseBodyRegisteredMediaList {
	s.NewRegister = &v
	return s
}

type RegisterMediaResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RegisterMediaResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RegisterMediaResponse) String() string {
	return tea.Prettify(s)
}

func (s RegisterMediaResponse) GoString() string {
	return s.String()
}

func (s *RegisterMediaResponse) SetHeaders(v map[string]*string) *RegisterMediaResponse {
	s.Headers = v
	return s
}

func (s *RegisterMediaResponse) SetBody(v *RegisterMediaResponseBody) *RegisterMediaResponse {
	s.Body = v
	return s
}

type SearchEditingProjectRequest struct {
	EndTime              *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNo               *int32  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *string `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SortBy               *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	StartTime            *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Status               *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Title                *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s SearchEditingProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchEditingProjectRequest) GoString() string {
	return s.String()
}

func (s *SearchEditingProjectRequest) SetEndTime(v string) *SearchEditingProjectRequest {
	s.EndTime = &v
	return s
}

func (s *SearchEditingProjectRequest) SetOwnerAccount(v string) *SearchEditingProjectRequest {
	s.OwnerAccount = &v
	return s
}

func (s *SearchEditingProjectRequest) SetOwnerId(v string) *SearchEditingProjectRequest {
	s.OwnerId = &v
	return s
}

func (s *SearchEditingProjectRequest) SetPageNo(v int32) *SearchEditingProjectRequest {
	s.PageNo = &v
	return s
}

func (s *SearchEditingProjectRequest) SetPageSize(v int32) *SearchEditingProjectRequest {
	s.PageSize = &v
	return s
}

func (s *SearchEditingProjectRequest) SetResourceOwnerAccount(v string) *SearchEditingProjectRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SearchEditingProjectRequest) SetResourceOwnerId(v string) *SearchEditingProjectRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SearchEditingProjectRequest) SetSortBy(v string) *SearchEditingProjectRequest {
	s.SortBy = &v
	return s
}

func (s *SearchEditingProjectRequest) SetStartTime(v string) *SearchEditingProjectRequest {
	s.StartTime = &v
	return s
}

func (s *SearchEditingProjectRequest) SetStatus(v string) *SearchEditingProjectRequest {
	s.Status = &v
	return s
}

func (s *SearchEditingProjectRequest) SetTitle(v string) *SearchEditingProjectRequest {
	s.Title = &v
	return s
}

type SearchEditingProjectResponseBody struct {
	ProjectList *SearchEditingProjectResponseBodyProjectList `json:"ProjectList,omitempty" xml:"ProjectList,omitempty" type:"Struct"`
	RequestId   *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Total       *int32                                       `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s SearchEditingProjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SearchEditingProjectResponseBody) GoString() string {
	return s.String()
}

func (s *SearchEditingProjectResponseBody) SetProjectList(v *SearchEditingProjectResponseBodyProjectList) *SearchEditingProjectResponseBody {
	s.ProjectList = v
	return s
}

func (s *SearchEditingProjectResponseBody) SetRequestId(v string) *SearchEditingProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *SearchEditingProjectResponseBody) SetTotal(v int32) *SearchEditingProjectResponseBody {
	s.Total = &v
	return s
}

type SearchEditingProjectResponseBodyProjectList struct {
	Project []*SearchEditingProjectResponseBodyProjectListProject `json:"Project,omitempty" xml:"Project,omitempty" type:"Repeated"`
}

func (s SearchEditingProjectResponseBodyProjectList) String() string {
	return tea.Prettify(s)
}

func (s SearchEditingProjectResponseBodyProjectList) GoString() string {
	return s.String()
}

func (s *SearchEditingProjectResponseBodyProjectList) SetProject(v []*SearchEditingProjectResponseBodyProjectListProject) *SearchEditingProjectResponseBodyProjectList {
	s.Project = v
	return s
}

type SearchEditingProjectResponseBodyProjectListProject struct {
	CoverURL        *string  `json:"CoverURL,omitempty" xml:"CoverURL,omitempty"`
	CreationTime    *string  `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	Description     *string  `json:"Description,omitempty" xml:"Description,omitempty"`
	Duration        *float32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	ModifiedTime    *string  `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	ProjectId       *string  `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	RegionId        *string  `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Status          *string  `json:"Status,omitempty" xml:"Status,omitempty"`
	StorageLocation *string  `json:"StorageLocation,omitempty" xml:"StorageLocation,omitempty"`
	Title           *string  `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s SearchEditingProjectResponseBodyProjectListProject) String() string {
	return tea.Prettify(s)
}

func (s SearchEditingProjectResponseBodyProjectListProject) GoString() string {
	return s.String()
}

func (s *SearchEditingProjectResponseBodyProjectListProject) SetCoverURL(v string) *SearchEditingProjectResponseBodyProjectListProject {
	s.CoverURL = &v
	return s
}

func (s *SearchEditingProjectResponseBodyProjectListProject) SetCreationTime(v string) *SearchEditingProjectResponseBodyProjectListProject {
	s.CreationTime = &v
	return s
}

func (s *SearchEditingProjectResponseBodyProjectListProject) SetDescription(v string) *SearchEditingProjectResponseBodyProjectListProject {
	s.Description = &v
	return s
}

func (s *SearchEditingProjectResponseBodyProjectListProject) SetDuration(v float32) *SearchEditingProjectResponseBodyProjectListProject {
	s.Duration = &v
	return s
}

func (s *SearchEditingProjectResponseBodyProjectListProject) SetModifiedTime(v string) *SearchEditingProjectResponseBodyProjectListProject {
	s.ModifiedTime = &v
	return s
}

func (s *SearchEditingProjectResponseBodyProjectListProject) SetProjectId(v string) *SearchEditingProjectResponseBodyProjectListProject {
	s.ProjectId = &v
	return s
}

func (s *SearchEditingProjectResponseBodyProjectListProject) SetRegionId(v string) *SearchEditingProjectResponseBodyProjectListProject {
	s.RegionId = &v
	return s
}

func (s *SearchEditingProjectResponseBodyProjectListProject) SetStatus(v string) *SearchEditingProjectResponseBodyProjectListProject {
	s.Status = &v
	return s
}

func (s *SearchEditingProjectResponseBodyProjectListProject) SetStorageLocation(v string) *SearchEditingProjectResponseBodyProjectListProject {
	s.StorageLocation = &v
	return s
}

func (s *SearchEditingProjectResponseBodyProjectListProject) SetTitle(v string) *SearchEditingProjectResponseBodyProjectListProject {
	s.Title = &v
	return s
}

type SearchEditingProjectResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SearchEditingProjectResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SearchEditingProjectResponse) String() string {
	return tea.Prettify(s)
}

func (s SearchEditingProjectResponse) GoString() string {
	return s.String()
}

func (s *SearchEditingProjectResponse) SetHeaders(v map[string]*string) *SearchEditingProjectResponse {
	s.Headers = v
	return s
}

func (s *SearchEditingProjectResponse) SetBody(v *SearchEditingProjectResponseBody) *SearchEditingProjectResponse {
	s.Body = v
	return s
}

type SearchMediaRequest struct {
	Fields      *string `json:"Fields,omitempty" xml:"Fields,omitempty"`
	Match       *string `json:"Match,omitempty" xml:"Match,omitempty"`
	PageNo      *int32  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ScrollToken *string `json:"ScrollToken,omitempty" xml:"ScrollToken,omitempty"`
	SearchType  *string `json:"SearchType,omitempty" xml:"SearchType,omitempty"`
	SortBy      *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
}

func (s SearchMediaRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchMediaRequest) GoString() string {
	return s.String()
}

func (s *SearchMediaRequest) SetFields(v string) *SearchMediaRequest {
	s.Fields = &v
	return s
}

func (s *SearchMediaRequest) SetMatch(v string) *SearchMediaRequest {
	s.Match = &v
	return s
}

func (s *SearchMediaRequest) SetPageNo(v int32) *SearchMediaRequest {
	s.PageNo = &v
	return s
}

func (s *SearchMediaRequest) SetPageSize(v int32) *SearchMediaRequest {
	s.PageSize = &v
	return s
}

func (s *SearchMediaRequest) SetScrollToken(v string) *SearchMediaRequest {
	s.ScrollToken = &v
	return s
}

func (s *SearchMediaRequest) SetSearchType(v string) *SearchMediaRequest {
	s.SearchType = &v
	return s
}

func (s *SearchMediaRequest) SetSortBy(v string) *SearchMediaRequest {
	s.SortBy = &v
	return s
}

type SearchMediaResponseBody struct {
	MediaList   []*SearchMediaResponseBodyMediaList `json:"MediaList,omitempty" xml:"MediaList,omitempty" type:"Repeated"`
	RequestId   *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ScrollToken *string                             `json:"ScrollToken,omitempty" xml:"ScrollToken,omitempty"`
	Total       *int64                              `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s SearchMediaResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SearchMediaResponseBody) GoString() string {
	return s.String()
}

func (s *SearchMediaResponseBody) SetMediaList(v []*SearchMediaResponseBodyMediaList) *SearchMediaResponseBody {
	s.MediaList = v
	return s
}

func (s *SearchMediaResponseBody) SetRequestId(v string) *SearchMediaResponseBody {
	s.RequestId = &v
	return s
}

func (s *SearchMediaResponseBody) SetScrollToken(v string) *SearchMediaResponseBody {
	s.ScrollToken = &v
	return s
}

func (s *SearchMediaResponseBody) SetTotal(v int64) *SearchMediaResponseBody {
	s.Total = &v
	return s
}

type SearchMediaResponseBodyMediaList struct {
	AttachedMedia *SearchMediaResponseBodyMediaListAttachedMedia `json:"AttachedMedia,omitempty" xml:"AttachedMedia,omitempty" type:"Struct"`
	Audio         *SearchMediaResponseBodyMediaListAudio         `json:"Audio,omitempty" xml:"Audio,omitempty" type:"Struct"`
	CreationTime  *string                                        `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	Image         *SearchMediaResponseBodyMediaListImage         `json:"Image,omitempty" xml:"Image,omitempty" type:"Struct"`
	MediaId       *string                                        `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	MediaType     *string                                        `json:"MediaType,omitempty" xml:"MediaType,omitempty"`
	Video         *SearchMediaResponseBodyMediaListVideo         `json:"Video,omitempty" xml:"Video,omitempty" type:"Struct"`
}

func (s SearchMediaResponseBodyMediaList) String() string {
	return tea.Prettify(s)
}

func (s SearchMediaResponseBodyMediaList) GoString() string {
	return s.String()
}

func (s *SearchMediaResponseBodyMediaList) SetAttachedMedia(v *SearchMediaResponseBodyMediaListAttachedMedia) *SearchMediaResponseBodyMediaList {
	s.AttachedMedia = v
	return s
}

func (s *SearchMediaResponseBodyMediaList) SetAudio(v *SearchMediaResponseBodyMediaListAudio) *SearchMediaResponseBodyMediaList {
	s.Audio = v
	return s
}

func (s *SearchMediaResponseBodyMediaList) SetCreationTime(v string) *SearchMediaResponseBodyMediaList {
	s.CreationTime = &v
	return s
}

func (s *SearchMediaResponseBodyMediaList) SetImage(v *SearchMediaResponseBodyMediaListImage) *SearchMediaResponseBodyMediaList {
	s.Image = v
	return s
}

func (s *SearchMediaResponseBodyMediaList) SetMediaId(v string) *SearchMediaResponseBodyMediaList {
	s.MediaId = &v
	return s
}

func (s *SearchMediaResponseBodyMediaList) SetMediaType(v string) *SearchMediaResponseBodyMediaList {
	s.MediaType = &v
	return s
}

func (s *SearchMediaResponseBodyMediaList) SetVideo(v *SearchMediaResponseBodyMediaListVideo) *SearchMediaResponseBodyMediaList {
	s.Video = v
	return s
}

type SearchMediaResponseBodyMediaListAttachedMedia struct {
	AppId            *string                                                    `json:"AppId,omitempty" xml:"AppId,omitempty"`
	BusinessType     *string                                                    `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	Categories       []*SearchMediaResponseBodyMediaListAttachedMediaCategories `json:"Categories,omitempty" xml:"Categories,omitempty" type:"Repeated"`
	CreationTime     *string                                                    `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	Description      *string                                                    `json:"Description,omitempty" xml:"Description,omitempty"`
	MediaId          *string                                                    `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	ModificationTime *string                                                    `json:"ModificationTime,omitempty" xml:"ModificationTime,omitempty"`
	Status           *string                                                    `json:"Status,omitempty" xml:"Status,omitempty"`
	StorageLocation  *string                                                    `json:"StorageLocation,omitempty" xml:"StorageLocation,omitempty"`
	Tags             *string                                                    `json:"Tags,omitempty" xml:"Tags,omitempty"`
	Title            *string                                                    `json:"Title,omitempty" xml:"Title,omitempty"`
	URL              *string                                                    `json:"URL,omitempty" xml:"URL,omitempty"`
}

func (s SearchMediaResponseBodyMediaListAttachedMedia) String() string {
	return tea.Prettify(s)
}

func (s SearchMediaResponseBodyMediaListAttachedMedia) GoString() string {
	return s.String()
}

func (s *SearchMediaResponseBodyMediaListAttachedMedia) SetAppId(v string) *SearchMediaResponseBodyMediaListAttachedMedia {
	s.AppId = &v
	return s
}

func (s *SearchMediaResponseBodyMediaListAttachedMedia) SetBusinessType(v string) *SearchMediaResponseBodyMediaListAttachedMedia {
	s.BusinessType = &v
	return s
}

func (s *SearchMediaResponseBodyMediaListAttachedMedia) SetCategories(v []*SearchMediaResponseBodyMediaListAttachedMediaCategories) *SearchMediaResponseBodyMediaListAttachedMedia {
	s.Categories = v
	return s
}

func (s *SearchMediaResponseBodyMediaListAttachedMedia) SetCreationTime(v string) *SearchMediaResponseBodyMediaListAttachedMedia {
	s.CreationTime = &v
	return s
}

func (s *SearchMediaResponseBodyMediaListAttachedMedia) SetDescription(v string) *SearchMediaResponseBodyMediaListAttachedMedia {
	s.Description = &v
	return s
}

func (s *SearchMediaResponseBodyMediaListAttachedMedia) SetMediaId(v string) *SearchMediaResponseBodyMediaListAttachedMedia {
	s.MediaId = &v
	return s
}

func (s *SearchMediaResponseBodyMediaListAttachedMedia) SetModificationTime(v string) *SearchMediaResponseBodyMediaListAttachedMedia {
	s.ModificationTime = &v
	return s
}

func (s *SearchMediaResponseBodyMediaListAttachedMedia) SetStatus(v string) *SearchMediaResponseBodyMediaListAttachedMedia {
	s.Status = &v
	return s
}

func (s *SearchMediaResponseBodyMediaListAttachedMedia) SetStorageLocation(v string) *SearchMediaResponseBodyMediaListAttachedMedia {
	s.StorageLocation = &v
	return s
}

func (s *SearchMediaResponseBodyMediaListAttachedMedia) SetTags(v string) *SearchMediaResponseBodyMediaListAttachedMedia {
	s.Tags = &v
	return s
}

func (s *SearchMediaResponseBodyMediaListAttachedMedia) SetTitle(v string) *SearchMediaResponseBodyMediaListAttachedMedia {
	s.Title = &v
	return s
}

func (s *SearchMediaResponseBodyMediaListAttachedMedia) SetURL(v string) *SearchMediaResponseBodyMediaListAttachedMedia {
	s.URL = &v
	return s
}

type SearchMediaResponseBodyMediaListAttachedMediaCategories struct {
	CateId   *int64  `json:"CateId,omitempty" xml:"CateId,omitempty"`
	CateName *string `json:"CateName,omitempty" xml:"CateName,omitempty"`
	Level    *int64  `json:"Level,omitempty" xml:"Level,omitempty"`
	ParentId *int64  `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
}

func (s SearchMediaResponseBodyMediaListAttachedMediaCategories) String() string {
	return tea.Prettify(s)
}

func (s SearchMediaResponseBodyMediaListAttachedMediaCategories) GoString() string {
	return s.String()
}

func (s *SearchMediaResponseBodyMediaListAttachedMediaCategories) SetCateId(v int64) *SearchMediaResponseBodyMediaListAttachedMediaCategories {
	s.CateId = &v
	return s
}

func (s *SearchMediaResponseBodyMediaListAttachedMediaCategories) SetCateName(v string) *SearchMediaResponseBodyMediaListAttachedMediaCategories {
	s.CateName = &v
	return s
}

func (s *SearchMediaResponseBodyMediaListAttachedMediaCategories) SetLevel(v int64) *SearchMediaResponseBodyMediaListAttachedMediaCategories {
	s.Level = &v
	return s
}

func (s *SearchMediaResponseBodyMediaListAttachedMediaCategories) SetParentId(v int64) *SearchMediaResponseBodyMediaListAttachedMediaCategories {
	s.ParentId = &v
	return s
}

type SearchMediaResponseBodyMediaListAudio struct {
	AppId            *string   `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AudioId          *string   `json:"AudioId,omitempty" xml:"AudioId,omitempty"`
	CateId           *int64    `json:"CateId,omitempty" xml:"CateId,omitempty"`
	CateName         *string   `json:"CateName,omitempty" xml:"CateName,omitempty"`
	CoverURL         *string   `json:"CoverURL,omitempty" xml:"CoverURL,omitempty"`
	CreationTime     *string   `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	Description      *string   `json:"Description,omitempty" xml:"Description,omitempty"`
	DownloadSwitch   *string   `json:"DownloadSwitch,omitempty" xml:"DownloadSwitch,omitempty"`
	Duration         *float32  `json:"Duration,omitempty" xml:"Duration,omitempty"`
	MediaSource      *string   `json:"MediaSource,omitempty" xml:"MediaSource,omitempty"`
	ModificationTime *string   `json:"ModificationTime,omitempty" xml:"ModificationTime,omitempty"`
	PreprocessStatus *string   `json:"PreprocessStatus,omitempty" xml:"PreprocessStatus,omitempty"`
	Size             *int64    `json:"Size,omitempty" xml:"Size,omitempty"`
	Snapshots        []*string `json:"Snapshots,omitempty" xml:"Snapshots,omitempty" type:"Repeated"`
	SpriteSnapshots  []*string `json:"SpriteSnapshots,omitempty" xml:"SpriteSnapshots,omitempty" type:"Repeated"`
	Status           *string   `json:"Status,omitempty" xml:"Status,omitempty"`
	StorageLocation  *string   `json:"StorageLocation,omitempty" xml:"StorageLocation,omitempty"`
	Tags             *string   `json:"Tags,omitempty" xml:"Tags,omitempty"`
	Title            *string   `json:"Title,omitempty" xml:"Title,omitempty"`
	TranscodeMode    *string   `json:"TranscodeMode,omitempty" xml:"TranscodeMode,omitempty"`
}

func (s SearchMediaResponseBodyMediaListAudio) String() string {
	return tea.Prettify(s)
}

func (s SearchMediaResponseBodyMediaListAudio) GoString() string {
	return s.String()
}

func (s *SearchMediaResponseBodyMediaListAudio) SetAppId(v string) *SearchMediaResponseBodyMediaListAudio {
	s.AppId = &v
	return s
}

func (s *SearchMediaResponseBodyMediaListAudio) SetAudioId(v string) *SearchMediaResponseBodyMediaListAudio {
	s.AudioId = &v
	return s
}

func (s *SearchMediaResponseBodyMediaListAudio) SetCateId(v int64) *SearchMediaResponseBodyMediaListAudio {
	s.CateId = &v
	return s
}

func (s *SearchMediaResponseBodyMediaListAudio) SetCateName(v string) *SearchMediaResponseBodyMediaListAudio {
	s.CateName = &v
	return s
}

func (s *SearchMediaResponseBodyMediaListAudio) SetCoverURL(v string) *SearchMediaResponseBodyMediaListAudio {
	s.CoverURL = &v
	return s
}

func (s *SearchMediaResponseBodyMediaListAudio) SetCreationTime(v string) *SearchMediaResponseBodyMediaListAudio {
	s.CreationTime = &v
	return s
}

func (s *SearchMediaResponseBodyMediaListAudio) SetDescription(v string) *SearchMediaResponseBodyMediaListAudio {
	s.Description = &v
	return s
}

func (s *SearchMediaResponseBodyMediaListAudio) SetDownloadSwitch(v string) *SearchMediaResponseBodyMediaListAudio {
	s.DownloadSwitch = &v
	return s
}

func (s *SearchMediaResponseBodyMediaListAudio) SetDuration(v float32) *SearchMediaResponseBodyMediaListAudio {
	s.Duration = &v
	return s
}

func (s *SearchMediaResponseBodyMediaListAudio) SetMediaSource(v string) *SearchMediaResponseBodyMediaListAudio {
	s.MediaSource = &v
	return s
}

func (s *SearchMediaResponseBodyMediaListAudio) SetModificationTime(v string) *SearchMediaResponseBodyMediaListAudio {
	s.ModificationTime = &v
	return s
}

func (s *SearchMediaResponseBodyMediaListAudio) SetPreprocessStatus(v string) *SearchMediaResponseBodyMediaListAudio {
	s.PreprocessStatus = &v
	return s
}

func (s *SearchMediaResponseBodyMediaListAudio) SetSize(v int64) *SearchMediaResponseBodyMediaListAudio {
	s.Size = &v
	return s
}

func (s *SearchMediaResponseBodyMediaListAudio) SetSnapshots(v []*string) *SearchMediaResponseBodyMediaListAudio {
	s.Snapshots = v
	return s
}

func (s *SearchMediaResponseBodyMediaListAudio) SetSpriteSnapshots(v []*string) *SearchMediaResponseBodyMediaListAudio {
	s.SpriteSnapshots = v
	return s
}

func (s *SearchMediaResponseBodyMediaListAudio) SetStatus(v string) *SearchMediaResponseBodyMediaListAudio {
	s.Status = &v
	return s
}

func (s *SearchMediaResponseBodyMediaListAudio) SetStorageLocation(v string) *SearchMediaResponseBodyMediaListAudio {
	s.StorageLocation = &v
	return s
}

func (s *SearchMediaResponseBodyMediaListAudio) SetTags(v string) *SearchMediaResponseBodyMediaListAudio {
	s.Tags = &v
	return s
}

func (s *SearchMediaResponseBodyMediaListAudio) SetTitle(v string) *SearchMediaResponseBodyMediaListAudio {
	s.Title = &v
	return s
}

func (s *SearchMediaResponseBodyMediaListAudio) SetTranscodeMode(v string) *SearchMediaResponseBodyMediaListAudio {
	s.TranscodeMode = &v
	return s
}

type SearchMediaResponseBodyMediaListImage struct {
	AppId            *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	CateId           *int64  `json:"CateId,omitempty" xml:"CateId,omitempty"`
	CateName         *string `json:"CateName,omitempty" xml:"CateName,omitempty"`
	CreationTime     *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	Description      *string `json:"Description,omitempty" xml:"Description,omitempty"`
	ImageId          *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	ModificationTime *string `json:"ModificationTime,omitempty" xml:"ModificationTime,omitempty"`
	Status           *string `json:"Status,omitempty" xml:"Status,omitempty"`
	StorageLocation  *string `json:"StorageLocation,omitempty" xml:"StorageLocation,omitempty"`
	Tags             *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	Title            *string `json:"Title,omitempty" xml:"Title,omitempty"`
	URL              *string `json:"URL,omitempty" xml:"URL,omitempty"`
}

func (s SearchMediaResponseBodyMediaListImage) String() string {
	return tea.Prettify(s)
}

func (s SearchMediaResponseBodyMediaListImage) GoString() string {
	return s.String()
}

func (s *SearchMediaResponseBodyMediaListImage) SetAppId(v string) *SearchMediaResponseBodyMediaListImage {
	s.AppId = &v
	return s
}

func (s *SearchMediaResponseBodyMediaListImage) SetCateId(v int64) *SearchMediaResponseBodyMediaListImage {
	s.CateId = &v
	return s
}

func (s *SearchMediaResponseBodyMediaListImage) SetCateName(v string) *SearchMediaResponseBodyMediaListImage {
	s.CateName = &v
	return s
}

func (s *SearchMediaResponseBodyMediaListImage) SetCreationTime(v string) *SearchMediaResponseBodyMediaListImage {
	s.CreationTime = &v
	return s
}

func (s *SearchMediaResponseBodyMediaListImage) SetDescription(v string) *SearchMediaResponseBodyMediaListImage {
	s.Description = &v
	return s
}

func (s *SearchMediaResponseBodyMediaListImage) SetImageId(v string) *SearchMediaResponseBodyMediaListImage {
	s.ImageId = &v
	return s
}

func (s *SearchMediaResponseBodyMediaListImage) SetModificationTime(v string) *SearchMediaResponseBodyMediaListImage {
	s.ModificationTime = &v
	return s
}

func (s *SearchMediaResponseBodyMediaListImage) SetStatus(v string) *SearchMediaResponseBodyMediaListImage {
	s.Status = &v
	return s
}

func (s *SearchMediaResponseBodyMediaListImage) SetStorageLocation(v string) *SearchMediaResponseBodyMediaListImage {
	s.StorageLocation = &v
	return s
}

func (s *SearchMediaResponseBodyMediaListImage) SetTags(v string) *SearchMediaResponseBodyMediaListImage {
	s.Tags = &v
	return s
}

func (s *SearchMediaResponseBodyMediaListImage) SetTitle(v string) *SearchMediaResponseBodyMediaListImage {
	s.Title = &v
	return s
}

func (s *SearchMediaResponseBodyMediaListImage) SetURL(v string) *SearchMediaResponseBodyMediaListImage {
	s.URL = &v
	return s
}

type SearchMediaResponseBodyMediaListVideo struct {
	AppId            *string   `json:"AppId,omitempty" xml:"AppId,omitempty"`
	CateId           *int64    `json:"CateId,omitempty" xml:"CateId,omitempty"`
	CateName         *string   `json:"CateName,omitempty" xml:"CateName,omitempty"`
	CoverURL         *string   `json:"CoverURL,omitempty" xml:"CoverURL,omitempty"`
	CreationTime     *string   `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	Description      *string   `json:"Description,omitempty" xml:"Description,omitempty"`
	DownloadSwitch   *string   `json:"DownloadSwitch,omitempty" xml:"DownloadSwitch,omitempty"`
	Duration         *float32  `json:"Duration,omitempty" xml:"Duration,omitempty"`
	MediaSource      *string   `json:"MediaSource,omitempty" xml:"MediaSource,omitempty"`
	ModificationTime *string   `json:"ModificationTime,omitempty" xml:"ModificationTime,omitempty"`
	PreprocessStatus *string   `json:"PreprocessStatus,omitempty" xml:"PreprocessStatus,omitempty"`
	Size             *int64    `json:"Size,omitempty" xml:"Size,omitempty"`
	Snapshots        []*string `json:"Snapshots,omitempty" xml:"Snapshots,omitempty" type:"Repeated"`
	SpriteSnapshots  []*string `json:"SpriteSnapshots,omitempty" xml:"SpriteSnapshots,omitempty" type:"Repeated"`
	Status           *string   `json:"Status,omitempty" xml:"Status,omitempty"`
	StorageLocation  *string   `json:"StorageLocation,omitempty" xml:"StorageLocation,omitempty"`
	Tags             *string   `json:"Tags,omitempty" xml:"Tags,omitempty"`
	Title            *string   `json:"Title,omitempty" xml:"Title,omitempty"`
	TranscodeMode    *string   `json:"TranscodeMode,omitempty" xml:"TranscodeMode,omitempty"`
	VideoId          *string   `json:"VideoId,omitempty" xml:"VideoId,omitempty"`
}

func (s SearchMediaResponseBodyMediaListVideo) String() string {
	return tea.Prettify(s)
}

func (s SearchMediaResponseBodyMediaListVideo) GoString() string {
	return s.String()
}

func (s *SearchMediaResponseBodyMediaListVideo) SetAppId(v string) *SearchMediaResponseBodyMediaListVideo {
	s.AppId = &v
	return s
}

func (s *SearchMediaResponseBodyMediaListVideo) SetCateId(v int64) *SearchMediaResponseBodyMediaListVideo {
	s.CateId = &v
	return s
}

func (s *SearchMediaResponseBodyMediaListVideo) SetCateName(v string) *SearchMediaResponseBodyMediaListVideo {
	s.CateName = &v
	return s
}

func (s *SearchMediaResponseBodyMediaListVideo) SetCoverURL(v string) *SearchMediaResponseBodyMediaListVideo {
	s.CoverURL = &v
	return s
}

func (s *SearchMediaResponseBodyMediaListVideo) SetCreationTime(v string) *SearchMediaResponseBodyMediaListVideo {
	s.CreationTime = &v
	return s
}

func (s *SearchMediaResponseBodyMediaListVideo) SetDescription(v string) *SearchMediaResponseBodyMediaListVideo {
	s.Description = &v
	return s
}

func (s *SearchMediaResponseBodyMediaListVideo) SetDownloadSwitch(v string) *SearchMediaResponseBodyMediaListVideo {
	s.DownloadSwitch = &v
	return s
}

func (s *SearchMediaResponseBodyMediaListVideo) SetDuration(v float32) *SearchMediaResponseBodyMediaListVideo {
	s.Duration = &v
	return s
}

func (s *SearchMediaResponseBodyMediaListVideo) SetMediaSource(v string) *SearchMediaResponseBodyMediaListVideo {
	s.MediaSource = &v
	return s
}

func (s *SearchMediaResponseBodyMediaListVideo) SetModificationTime(v string) *SearchMediaResponseBodyMediaListVideo {
	s.ModificationTime = &v
	return s
}

func (s *SearchMediaResponseBodyMediaListVideo) SetPreprocessStatus(v string) *SearchMediaResponseBodyMediaListVideo {
	s.PreprocessStatus = &v
	return s
}

func (s *SearchMediaResponseBodyMediaListVideo) SetSize(v int64) *SearchMediaResponseBodyMediaListVideo {
	s.Size = &v
	return s
}

func (s *SearchMediaResponseBodyMediaListVideo) SetSnapshots(v []*string) *SearchMediaResponseBodyMediaListVideo {
	s.Snapshots = v
	return s
}

func (s *SearchMediaResponseBodyMediaListVideo) SetSpriteSnapshots(v []*string) *SearchMediaResponseBodyMediaListVideo {
	s.SpriteSnapshots = v
	return s
}

func (s *SearchMediaResponseBodyMediaListVideo) SetStatus(v string) *SearchMediaResponseBodyMediaListVideo {
	s.Status = &v
	return s
}

func (s *SearchMediaResponseBodyMediaListVideo) SetStorageLocation(v string) *SearchMediaResponseBodyMediaListVideo {
	s.StorageLocation = &v
	return s
}

func (s *SearchMediaResponseBodyMediaListVideo) SetTags(v string) *SearchMediaResponseBodyMediaListVideo {
	s.Tags = &v
	return s
}

func (s *SearchMediaResponseBodyMediaListVideo) SetTitle(v string) *SearchMediaResponseBodyMediaListVideo {
	s.Title = &v
	return s
}

func (s *SearchMediaResponseBodyMediaListVideo) SetTranscodeMode(v string) *SearchMediaResponseBodyMediaListVideo {
	s.TranscodeMode = &v
	return s
}

func (s *SearchMediaResponseBodyMediaListVideo) SetVideoId(v string) *SearchMediaResponseBodyMediaListVideo {
	s.VideoId = &v
	return s
}

type SearchMediaResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SearchMediaResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SearchMediaResponse) String() string {
	return tea.Prettify(s)
}

func (s SearchMediaResponse) GoString() string {
	return s.String()
}

func (s *SearchMediaResponse) SetHeaders(v map[string]*string) *SearchMediaResponse {
	s.Headers = v
	return s
}

func (s *SearchMediaResponse) SetBody(v *SearchMediaResponseBody) *SearchMediaResponse {
	s.Body = v
	return s
}

type SetAuditSecurityIpRequest struct {
	Ips               *string `json:"Ips,omitempty" xml:"Ips,omitempty"`
	OperateMode       *string `json:"OperateMode,omitempty" xml:"OperateMode,omitempty"`
	SecurityGroupName *string `json:"SecurityGroupName,omitempty" xml:"SecurityGroupName,omitempty"`
}

func (s SetAuditSecurityIpRequest) String() string {
	return tea.Prettify(s)
}

func (s SetAuditSecurityIpRequest) GoString() string {
	return s.String()
}

func (s *SetAuditSecurityIpRequest) SetIps(v string) *SetAuditSecurityIpRequest {
	s.Ips = &v
	return s
}

func (s *SetAuditSecurityIpRequest) SetOperateMode(v string) *SetAuditSecurityIpRequest {
	s.OperateMode = &v
	return s
}

func (s *SetAuditSecurityIpRequest) SetSecurityGroupName(v string) *SetAuditSecurityIpRequest {
	s.SecurityGroupName = &v
	return s
}

type SetAuditSecurityIpResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetAuditSecurityIpResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetAuditSecurityIpResponseBody) GoString() string {
	return s.String()
}

func (s *SetAuditSecurityIpResponseBody) SetRequestId(v string) *SetAuditSecurityIpResponseBody {
	s.RequestId = &v
	return s
}

type SetAuditSecurityIpResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SetAuditSecurityIpResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetAuditSecurityIpResponse) String() string {
	return tea.Prettify(s)
}

func (s SetAuditSecurityIpResponse) GoString() string {
	return s.String()
}

func (s *SetAuditSecurityIpResponse) SetHeaders(v map[string]*string) *SetAuditSecurityIpResponse {
	s.Headers = v
	return s
}

func (s *SetAuditSecurityIpResponse) SetBody(v *SetAuditSecurityIpResponseBody) *SetAuditSecurityIpResponse {
	s.Body = v
	return s
}

type SetCrossdomainContentRequest struct {
	Content              *string `json:"Content,omitempty" xml:"Content,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *string `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ResourceRealOwnerId  *string `json:"ResourceRealOwnerId,omitempty" xml:"ResourceRealOwnerId,omitempty"`
	StorageLocation      *string `json:"StorageLocation,omitempty" xml:"StorageLocation,omitempty"`
}

func (s SetCrossdomainContentRequest) String() string {
	return tea.Prettify(s)
}

func (s SetCrossdomainContentRequest) GoString() string {
	return s.String()
}

func (s *SetCrossdomainContentRequest) SetContent(v string) *SetCrossdomainContentRequest {
	s.Content = &v
	return s
}

func (s *SetCrossdomainContentRequest) SetOwnerAccount(v string) *SetCrossdomainContentRequest {
	s.OwnerAccount = &v
	return s
}

func (s *SetCrossdomainContentRequest) SetOwnerId(v string) *SetCrossdomainContentRequest {
	s.OwnerId = &v
	return s
}

func (s *SetCrossdomainContentRequest) SetResourceOwnerAccount(v string) *SetCrossdomainContentRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SetCrossdomainContentRequest) SetResourceOwnerId(v string) *SetCrossdomainContentRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SetCrossdomainContentRequest) SetResourceRealOwnerId(v string) *SetCrossdomainContentRequest {
	s.ResourceRealOwnerId = &v
	return s
}

func (s *SetCrossdomainContentRequest) SetStorageLocation(v string) *SetCrossdomainContentRequest {
	s.StorageLocation = &v
	return s
}

type SetCrossdomainContentResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetCrossdomainContentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetCrossdomainContentResponseBody) GoString() string {
	return s.String()
}

func (s *SetCrossdomainContentResponseBody) SetRequestId(v string) *SetCrossdomainContentResponseBody {
	s.RequestId = &v
	return s
}

type SetCrossdomainContentResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SetCrossdomainContentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetCrossdomainContentResponse) String() string {
	return tea.Prettify(s)
}

func (s SetCrossdomainContentResponse) GoString() string {
	return s.String()
}

func (s *SetCrossdomainContentResponse) SetHeaders(v map[string]*string) *SetCrossdomainContentResponse {
	s.Headers = v
	return s
}

func (s *SetCrossdomainContentResponse) SetBody(v *SetCrossdomainContentResponseBody) *SetCrossdomainContentResponse {
	s.Body = v
	return s
}

type SetDefaultAITemplateRequest struct {
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s SetDefaultAITemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s SetDefaultAITemplateRequest) GoString() string {
	return s.String()
}

func (s *SetDefaultAITemplateRequest) SetTemplateId(v string) *SetDefaultAITemplateRequest {
	s.TemplateId = &v
	return s
}

type SetDefaultAITemplateResponseBody struct {
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s SetDefaultAITemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetDefaultAITemplateResponseBody) GoString() string {
	return s.String()
}

func (s *SetDefaultAITemplateResponseBody) SetRequestId(v string) *SetDefaultAITemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetDefaultAITemplateResponseBody) SetTemplateId(v string) *SetDefaultAITemplateResponseBody {
	s.TemplateId = &v
	return s
}

type SetDefaultAITemplateResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SetDefaultAITemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetDefaultAITemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s SetDefaultAITemplateResponse) GoString() string {
	return s.String()
}

func (s *SetDefaultAITemplateResponse) SetHeaders(v map[string]*string) *SetDefaultAITemplateResponse {
	s.Headers = v
	return s
}

func (s *SetDefaultAITemplateResponse) SetBody(v *SetDefaultAITemplateResponseBody) *SetDefaultAITemplateResponse {
	s.Body = v
	return s
}

type SetDefaultTranscodeTemplateGroupRequest struct {
	TranscodeTemplateGroupId *string `json:"TranscodeTemplateGroupId,omitempty" xml:"TranscodeTemplateGroupId,omitempty"`
}

func (s SetDefaultTranscodeTemplateGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s SetDefaultTranscodeTemplateGroupRequest) GoString() string {
	return s.String()
}

func (s *SetDefaultTranscodeTemplateGroupRequest) SetTranscodeTemplateGroupId(v string) *SetDefaultTranscodeTemplateGroupRequest {
	s.TranscodeTemplateGroupId = &v
	return s
}

type SetDefaultTranscodeTemplateGroupResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetDefaultTranscodeTemplateGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetDefaultTranscodeTemplateGroupResponseBody) GoString() string {
	return s.String()
}

func (s *SetDefaultTranscodeTemplateGroupResponseBody) SetRequestId(v string) *SetDefaultTranscodeTemplateGroupResponseBody {
	s.RequestId = &v
	return s
}

type SetDefaultTranscodeTemplateGroupResponse struct {
	Headers map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SetDefaultTranscodeTemplateGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetDefaultTranscodeTemplateGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s SetDefaultTranscodeTemplateGroupResponse) GoString() string {
	return s.String()
}

func (s *SetDefaultTranscodeTemplateGroupResponse) SetHeaders(v map[string]*string) *SetDefaultTranscodeTemplateGroupResponse {
	s.Headers = v
	return s
}

func (s *SetDefaultTranscodeTemplateGroupResponse) SetBody(v *SetDefaultTranscodeTemplateGroupResponseBody) *SetDefaultTranscodeTemplateGroupResponse {
	s.Body = v
	return s
}

type SetDefaultWatermarkRequest struct {
	WatermarkId *string `json:"WatermarkId,omitempty" xml:"WatermarkId,omitempty"`
}

func (s SetDefaultWatermarkRequest) String() string {
	return tea.Prettify(s)
}

func (s SetDefaultWatermarkRequest) GoString() string {
	return s.String()
}

func (s *SetDefaultWatermarkRequest) SetWatermarkId(v string) *SetDefaultWatermarkRequest {
	s.WatermarkId = &v
	return s
}

type SetDefaultWatermarkResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetDefaultWatermarkResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetDefaultWatermarkResponseBody) GoString() string {
	return s.String()
}

func (s *SetDefaultWatermarkResponseBody) SetRequestId(v string) *SetDefaultWatermarkResponseBody {
	s.RequestId = &v
	return s
}

type SetDefaultWatermarkResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SetDefaultWatermarkResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetDefaultWatermarkResponse) String() string {
	return tea.Prettify(s)
}

func (s SetDefaultWatermarkResponse) GoString() string {
	return s.String()
}

func (s *SetDefaultWatermarkResponse) SetHeaders(v map[string]*string) *SetDefaultWatermarkResponse {
	s.Headers = v
	return s
}

func (s *SetDefaultWatermarkResponse) SetBody(v *SetDefaultWatermarkResponseBody) *SetDefaultWatermarkResponse {
	s.Body = v
	return s
}

type SetEditingProjectMaterialsRequest struct {
	MaterialIds          *string `json:"MaterialIds,omitempty" xml:"MaterialIds,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ProjectId            *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *string `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s SetEditingProjectMaterialsRequest) String() string {
	return tea.Prettify(s)
}

func (s SetEditingProjectMaterialsRequest) GoString() string {
	return s.String()
}

func (s *SetEditingProjectMaterialsRequest) SetMaterialIds(v string) *SetEditingProjectMaterialsRequest {
	s.MaterialIds = &v
	return s
}

func (s *SetEditingProjectMaterialsRequest) SetOwnerAccount(v string) *SetEditingProjectMaterialsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *SetEditingProjectMaterialsRequest) SetOwnerId(v string) *SetEditingProjectMaterialsRequest {
	s.OwnerId = &v
	return s
}

func (s *SetEditingProjectMaterialsRequest) SetProjectId(v string) *SetEditingProjectMaterialsRequest {
	s.ProjectId = &v
	return s
}

func (s *SetEditingProjectMaterialsRequest) SetResourceOwnerAccount(v string) *SetEditingProjectMaterialsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SetEditingProjectMaterialsRequest) SetResourceOwnerId(v string) *SetEditingProjectMaterialsRequest {
	s.ResourceOwnerId = &v
	return s
}

type SetEditingProjectMaterialsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetEditingProjectMaterialsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetEditingProjectMaterialsResponseBody) GoString() string {
	return s.String()
}

func (s *SetEditingProjectMaterialsResponseBody) SetRequestId(v string) *SetEditingProjectMaterialsResponseBody {
	s.RequestId = &v
	return s
}

type SetEditingProjectMaterialsResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SetEditingProjectMaterialsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetEditingProjectMaterialsResponse) String() string {
	return tea.Prettify(s)
}

func (s SetEditingProjectMaterialsResponse) GoString() string {
	return s.String()
}

func (s *SetEditingProjectMaterialsResponse) SetHeaders(v map[string]*string) *SetEditingProjectMaterialsResponse {
	s.Headers = v
	return s
}

func (s *SetEditingProjectMaterialsResponse) SetBody(v *SetEditingProjectMaterialsResponseBody) *SetEditingProjectMaterialsResponse {
	s.Body = v
	return s
}

type SetMessageCallbackRequest struct {
	AppId         *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AuthKey       *string `json:"AuthKey,omitempty" xml:"AuthKey,omitempty"`
	AuthSwitch    *string `json:"AuthSwitch,omitempty" xml:"AuthSwitch,omitempty"`
	CallbackType  *string `json:"CallbackType,omitempty" xml:"CallbackType,omitempty"`
	CallbackURL   *string `json:"CallbackURL,omitempty" xml:"CallbackURL,omitempty"`
	EventTypeList *string `json:"EventTypeList,omitempty" xml:"EventTypeList,omitempty"`
	MnsEndpoint   *string `json:"MnsEndpoint,omitempty" xml:"MnsEndpoint,omitempty"`
	MnsQueueName  *string `json:"MnsQueueName,omitempty" xml:"MnsQueueName,omitempty"`
	OwnerAccount  *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
}

func (s SetMessageCallbackRequest) String() string {
	return tea.Prettify(s)
}

func (s SetMessageCallbackRequest) GoString() string {
	return s.String()
}

func (s *SetMessageCallbackRequest) SetAppId(v string) *SetMessageCallbackRequest {
	s.AppId = &v
	return s
}

func (s *SetMessageCallbackRequest) SetAuthKey(v string) *SetMessageCallbackRequest {
	s.AuthKey = &v
	return s
}

func (s *SetMessageCallbackRequest) SetAuthSwitch(v string) *SetMessageCallbackRequest {
	s.AuthSwitch = &v
	return s
}

func (s *SetMessageCallbackRequest) SetCallbackType(v string) *SetMessageCallbackRequest {
	s.CallbackType = &v
	return s
}

func (s *SetMessageCallbackRequest) SetCallbackURL(v string) *SetMessageCallbackRequest {
	s.CallbackURL = &v
	return s
}

func (s *SetMessageCallbackRequest) SetEventTypeList(v string) *SetMessageCallbackRequest {
	s.EventTypeList = &v
	return s
}

func (s *SetMessageCallbackRequest) SetMnsEndpoint(v string) *SetMessageCallbackRequest {
	s.MnsEndpoint = &v
	return s
}

func (s *SetMessageCallbackRequest) SetMnsQueueName(v string) *SetMessageCallbackRequest {
	s.MnsQueueName = &v
	return s
}

func (s *SetMessageCallbackRequest) SetOwnerAccount(v string) *SetMessageCallbackRequest {
	s.OwnerAccount = &v
	return s
}

type SetMessageCallbackResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetMessageCallbackResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetMessageCallbackResponseBody) GoString() string {
	return s.String()
}

func (s *SetMessageCallbackResponseBody) SetRequestId(v string) *SetMessageCallbackResponseBody {
	s.RequestId = &v
	return s
}

type SetMessageCallbackResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SetMessageCallbackResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetMessageCallbackResponse) String() string {
	return tea.Prettify(s)
}

func (s SetMessageCallbackResponse) GoString() string {
	return s.String()
}

func (s *SetMessageCallbackResponse) SetHeaders(v map[string]*string) *SetMessageCallbackResponse {
	s.Headers = v
	return s
}

func (s *SetMessageCallbackResponse) SetBody(v *SetMessageCallbackResponseBody) *SetMessageCallbackResponse {
	s.Body = v
	return s
}

type SetVodDomainCertificateRequest struct {
	CertName      *string `json:"CertName,omitempty" xml:"CertName,omitempty"`
	DomainName    *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SSLPri        *string `json:"SSLPri,omitempty" xml:"SSLPri,omitempty"`
	SSLProtocol   *string `json:"SSLProtocol,omitempty" xml:"SSLProtocol,omitempty"`
	SSLPub        *string `json:"SSLPub,omitempty" xml:"SSLPub,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s SetVodDomainCertificateRequest) String() string {
	return tea.Prettify(s)
}

func (s SetVodDomainCertificateRequest) GoString() string {
	return s.String()
}

func (s *SetVodDomainCertificateRequest) SetCertName(v string) *SetVodDomainCertificateRequest {
	s.CertName = &v
	return s
}

func (s *SetVodDomainCertificateRequest) SetDomainName(v string) *SetVodDomainCertificateRequest {
	s.DomainName = &v
	return s
}

func (s *SetVodDomainCertificateRequest) SetOwnerId(v int64) *SetVodDomainCertificateRequest {
	s.OwnerId = &v
	return s
}

func (s *SetVodDomainCertificateRequest) SetSSLPri(v string) *SetVodDomainCertificateRequest {
	s.SSLPri = &v
	return s
}

func (s *SetVodDomainCertificateRequest) SetSSLProtocol(v string) *SetVodDomainCertificateRequest {
	s.SSLProtocol = &v
	return s
}

func (s *SetVodDomainCertificateRequest) SetSSLPub(v string) *SetVodDomainCertificateRequest {
	s.SSLPub = &v
	return s
}

func (s *SetVodDomainCertificateRequest) SetSecurityToken(v string) *SetVodDomainCertificateRequest {
	s.SecurityToken = &v
	return s
}

type SetVodDomainCertificateResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetVodDomainCertificateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetVodDomainCertificateResponseBody) GoString() string {
	return s.String()
}

func (s *SetVodDomainCertificateResponseBody) SetRequestId(v string) *SetVodDomainCertificateResponseBody {
	s.RequestId = &v
	return s
}

type SetVodDomainCertificateResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SetVodDomainCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetVodDomainCertificateResponse) String() string {
	return tea.Prettify(s)
}

func (s SetVodDomainCertificateResponse) GoString() string {
	return s.String()
}

func (s *SetVodDomainCertificateResponse) SetHeaders(v map[string]*string) *SetVodDomainCertificateResponse {
	s.Headers = v
	return s
}

func (s *SetVodDomainCertificateResponse) SetBody(v *SetVodDomainCertificateResponseBody) *SetVodDomainCertificateResponse {
	s.Body = v
	return s
}

type SubmitAICaptionExtractionJobRequest struct {
	AIPipelineId *string `json:"AIPipelineId,omitempty" xml:"AIPipelineId,omitempty"`
	JobConfig    *string `json:"JobConfig,omitempty" xml:"JobConfig,omitempty"`
	UserData     *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
	VideoId      *string `json:"VideoId,omitempty" xml:"VideoId,omitempty"`
}

func (s SubmitAICaptionExtractionJobRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitAICaptionExtractionJobRequest) GoString() string {
	return s.String()
}

func (s *SubmitAICaptionExtractionJobRequest) SetAIPipelineId(v string) *SubmitAICaptionExtractionJobRequest {
	s.AIPipelineId = &v
	return s
}

func (s *SubmitAICaptionExtractionJobRequest) SetJobConfig(v string) *SubmitAICaptionExtractionJobRequest {
	s.JobConfig = &v
	return s
}

func (s *SubmitAICaptionExtractionJobRequest) SetUserData(v string) *SubmitAICaptionExtractionJobRequest {
	s.UserData = &v
	return s
}

func (s *SubmitAICaptionExtractionJobRequest) SetVideoId(v string) *SubmitAICaptionExtractionJobRequest {
	s.VideoId = &v
	return s
}

type SubmitAICaptionExtractionJobResponseBody struct {
	JobId     *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SubmitAICaptionExtractionJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SubmitAICaptionExtractionJobResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitAICaptionExtractionJobResponseBody) SetJobId(v string) *SubmitAICaptionExtractionJobResponseBody {
	s.JobId = &v
	return s
}

func (s *SubmitAICaptionExtractionJobResponseBody) SetRequestId(v string) *SubmitAICaptionExtractionJobResponseBody {
	s.RequestId = &v
	return s
}

type SubmitAICaptionExtractionJobResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SubmitAICaptionExtractionJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SubmitAICaptionExtractionJobResponse) String() string {
	return tea.Prettify(s)
}

func (s SubmitAICaptionExtractionJobResponse) GoString() string {
	return s.String()
}

func (s *SubmitAICaptionExtractionJobResponse) SetHeaders(v map[string]*string) *SubmitAICaptionExtractionJobResponse {
	s.Headers = v
	return s
}

func (s *SubmitAICaptionExtractionJobResponse) SetBody(v *SubmitAICaptionExtractionJobResponseBody) *SubmitAICaptionExtractionJobResponse {
	s.Body = v
	return s
}

type SubmitAIImageAuditJobRequest struct {
	MediaAuditConfiguration *string `json:"MediaAuditConfiguration,omitempty" xml:"MediaAuditConfiguration,omitempty"`
	MediaId                 *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	OwnerAccount            *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId                 *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount    *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId         *string `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	TemplateId              *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s SubmitAIImageAuditJobRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitAIImageAuditJobRequest) GoString() string {
	return s.String()
}

func (s *SubmitAIImageAuditJobRequest) SetMediaAuditConfiguration(v string) *SubmitAIImageAuditJobRequest {
	s.MediaAuditConfiguration = &v
	return s
}

func (s *SubmitAIImageAuditJobRequest) SetMediaId(v string) *SubmitAIImageAuditJobRequest {
	s.MediaId = &v
	return s
}

func (s *SubmitAIImageAuditJobRequest) SetOwnerAccount(v string) *SubmitAIImageAuditJobRequest {
	s.OwnerAccount = &v
	return s
}

func (s *SubmitAIImageAuditJobRequest) SetOwnerId(v string) *SubmitAIImageAuditJobRequest {
	s.OwnerId = &v
	return s
}

func (s *SubmitAIImageAuditJobRequest) SetResourceOwnerAccount(v string) *SubmitAIImageAuditJobRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SubmitAIImageAuditJobRequest) SetResourceOwnerId(v string) *SubmitAIImageAuditJobRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SubmitAIImageAuditJobRequest) SetTemplateId(v string) *SubmitAIImageAuditJobRequest {
	s.TemplateId = &v
	return s
}

type SubmitAIImageAuditJobResponseBody struct {
	JobId     *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SubmitAIImageAuditJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SubmitAIImageAuditJobResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitAIImageAuditJobResponseBody) SetJobId(v string) *SubmitAIImageAuditJobResponseBody {
	s.JobId = &v
	return s
}

func (s *SubmitAIImageAuditJobResponseBody) SetRequestId(v string) *SubmitAIImageAuditJobResponseBody {
	s.RequestId = &v
	return s
}

type SubmitAIImageAuditJobResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SubmitAIImageAuditJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SubmitAIImageAuditJobResponse) String() string {
	return tea.Prettify(s)
}

func (s SubmitAIImageAuditJobResponse) GoString() string {
	return s.String()
}

func (s *SubmitAIImageAuditJobResponse) SetHeaders(v map[string]*string) *SubmitAIImageAuditJobResponse {
	s.Headers = v
	return s
}

func (s *SubmitAIImageAuditJobResponse) SetBody(v *SubmitAIImageAuditJobResponseBody) *SubmitAIImageAuditJobResponse {
	s.Body = v
	return s
}

type SubmitAIImageJobRequest struct {
	AIPipelineId         *string `json:"AIPipelineId,omitempty" xml:"AIPipelineId,omitempty"`
	AITemplateId         *string `json:"AITemplateId,omitempty" xml:"AITemplateId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *string `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	UserData             *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
	VideoId              *string `json:"VideoId,omitempty" xml:"VideoId,omitempty"`
}

func (s SubmitAIImageJobRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitAIImageJobRequest) GoString() string {
	return s.String()
}

func (s *SubmitAIImageJobRequest) SetAIPipelineId(v string) *SubmitAIImageJobRequest {
	s.AIPipelineId = &v
	return s
}

func (s *SubmitAIImageJobRequest) SetAITemplateId(v string) *SubmitAIImageJobRequest {
	s.AITemplateId = &v
	return s
}

func (s *SubmitAIImageJobRequest) SetOwnerAccount(v string) *SubmitAIImageJobRequest {
	s.OwnerAccount = &v
	return s
}

func (s *SubmitAIImageJobRequest) SetOwnerId(v string) *SubmitAIImageJobRequest {
	s.OwnerId = &v
	return s
}

func (s *SubmitAIImageJobRequest) SetResourceOwnerAccount(v string) *SubmitAIImageJobRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SubmitAIImageJobRequest) SetResourceOwnerId(v string) *SubmitAIImageJobRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SubmitAIImageJobRequest) SetUserData(v string) *SubmitAIImageJobRequest {
	s.UserData = &v
	return s
}

func (s *SubmitAIImageJobRequest) SetVideoId(v string) *SubmitAIImageJobRequest {
	s.VideoId = &v
	return s
}

type SubmitAIImageJobResponseBody struct {
	JobId     *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SubmitAIImageJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SubmitAIImageJobResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitAIImageJobResponseBody) SetJobId(v string) *SubmitAIImageJobResponseBody {
	s.JobId = &v
	return s
}

func (s *SubmitAIImageJobResponseBody) SetRequestId(v string) *SubmitAIImageJobResponseBody {
	s.RequestId = &v
	return s
}

type SubmitAIImageJobResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SubmitAIImageJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SubmitAIImageJobResponse) String() string {
	return tea.Prettify(s)
}

func (s SubmitAIImageJobResponse) GoString() string {
	return s.String()
}

func (s *SubmitAIImageJobResponse) SetHeaders(v map[string]*string) *SubmitAIImageJobResponse {
	s.Headers = v
	return s
}

func (s *SubmitAIImageJobResponse) SetBody(v *SubmitAIImageJobResponseBody) *SubmitAIImageJobResponse {
	s.Body = v
	return s
}

type SubmitAIJobRequest struct {
	Config               *string `json:"Config,omitempty" xml:"Config,omitempty"`
	MediaId              *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *string `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	Types                *string `json:"Types,omitempty" xml:"Types,omitempty"`
	UserData             *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s SubmitAIJobRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitAIJobRequest) GoString() string {
	return s.String()
}

func (s *SubmitAIJobRequest) SetConfig(v string) *SubmitAIJobRequest {
	s.Config = &v
	return s
}

func (s *SubmitAIJobRequest) SetMediaId(v string) *SubmitAIJobRequest {
	s.MediaId = &v
	return s
}

func (s *SubmitAIJobRequest) SetOwnerAccount(v string) *SubmitAIJobRequest {
	s.OwnerAccount = &v
	return s
}

func (s *SubmitAIJobRequest) SetOwnerId(v string) *SubmitAIJobRequest {
	s.OwnerId = &v
	return s
}

func (s *SubmitAIJobRequest) SetResourceOwnerAccount(v string) *SubmitAIJobRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SubmitAIJobRequest) SetResourceOwnerId(v string) *SubmitAIJobRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SubmitAIJobRequest) SetTypes(v string) *SubmitAIJobRequest {
	s.Types = &v
	return s
}

func (s *SubmitAIJobRequest) SetUserData(v string) *SubmitAIJobRequest {
	s.UserData = &v
	return s
}

type SubmitAIJobResponseBody struct {
	AIJobList *SubmitAIJobResponseBodyAIJobList `json:"AIJobList,omitempty" xml:"AIJobList,omitempty" type:"Struct"`
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SubmitAIJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SubmitAIJobResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitAIJobResponseBody) SetAIJobList(v *SubmitAIJobResponseBodyAIJobList) *SubmitAIJobResponseBody {
	s.AIJobList = v
	return s
}

func (s *SubmitAIJobResponseBody) SetRequestId(v string) *SubmitAIJobResponseBody {
	s.RequestId = &v
	return s
}

type SubmitAIJobResponseBodyAIJobList struct {
	AIJob []*SubmitAIJobResponseBodyAIJobListAIJob `json:"AIJob,omitempty" xml:"AIJob,omitempty" type:"Repeated"`
}

func (s SubmitAIJobResponseBodyAIJobList) String() string {
	return tea.Prettify(s)
}

func (s SubmitAIJobResponseBodyAIJobList) GoString() string {
	return s.String()
}

func (s *SubmitAIJobResponseBodyAIJobList) SetAIJob(v []*SubmitAIJobResponseBodyAIJobListAIJob) *SubmitAIJobResponseBodyAIJobList {
	s.AIJob = v
	return s
}

type SubmitAIJobResponseBodyAIJobListAIJob struct {
	JobId   *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	Type    *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s SubmitAIJobResponseBodyAIJobListAIJob) String() string {
	return tea.Prettify(s)
}

func (s SubmitAIJobResponseBodyAIJobListAIJob) GoString() string {
	return s.String()
}

func (s *SubmitAIJobResponseBodyAIJobListAIJob) SetJobId(v string) *SubmitAIJobResponseBodyAIJobListAIJob {
	s.JobId = &v
	return s
}

func (s *SubmitAIJobResponseBodyAIJobListAIJob) SetMediaId(v string) *SubmitAIJobResponseBodyAIJobListAIJob {
	s.MediaId = &v
	return s
}

func (s *SubmitAIJobResponseBodyAIJobListAIJob) SetType(v string) *SubmitAIJobResponseBodyAIJobListAIJob {
	s.Type = &v
	return s
}

type SubmitAIJobResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SubmitAIJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SubmitAIJobResponse) String() string {
	return tea.Prettify(s)
}

func (s SubmitAIJobResponse) GoString() string {
	return s.String()
}

func (s *SubmitAIJobResponse) SetHeaders(v map[string]*string) *SubmitAIJobResponse {
	s.Headers = v
	return s
}

func (s *SubmitAIJobResponse) SetBody(v *SubmitAIJobResponseBody) *SubmitAIJobResponse {
	s.Body = v
	return s
}

type SubmitAIMediaAuditJobRequest struct {
	MediaAuditConfiguration *string `json:"MediaAuditConfiguration,omitempty" xml:"MediaAuditConfiguration,omitempty"`
	MediaId                 *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	MediaType               *string `json:"MediaType,omitempty" xml:"MediaType,omitempty"`
	TemplateId              *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	UserData                *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s SubmitAIMediaAuditJobRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitAIMediaAuditJobRequest) GoString() string {
	return s.String()
}

func (s *SubmitAIMediaAuditJobRequest) SetMediaAuditConfiguration(v string) *SubmitAIMediaAuditJobRequest {
	s.MediaAuditConfiguration = &v
	return s
}

func (s *SubmitAIMediaAuditJobRequest) SetMediaId(v string) *SubmitAIMediaAuditJobRequest {
	s.MediaId = &v
	return s
}

func (s *SubmitAIMediaAuditJobRequest) SetMediaType(v string) *SubmitAIMediaAuditJobRequest {
	s.MediaType = &v
	return s
}

func (s *SubmitAIMediaAuditJobRequest) SetTemplateId(v string) *SubmitAIMediaAuditJobRequest {
	s.TemplateId = &v
	return s
}

func (s *SubmitAIMediaAuditJobRequest) SetUserData(v string) *SubmitAIMediaAuditJobRequest {
	s.UserData = &v
	return s
}

type SubmitAIMediaAuditJobResponseBody struct {
	JobId     *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	MediaId   *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SubmitAIMediaAuditJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SubmitAIMediaAuditJobResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitAIMediaAuditJobResponseBody) SetJobId(v string) *SubmitAIMediaAuditJobResponseBody {
	s.JobId = &v
	return s
}

func (s *SubmitAIMediaAuditJobResponseBody) SetMediaId(v string) *SubmitAIMediaAuditJobResponseBody {
	s.MediaId = &v
	return s
}

func (s *SubmitAIMediaAuditJobResponseBody) SetRequestId(v string) *SubmitAIMediaAuditJobResponseBody {
	s.RequestId = &v
	return s
}

type SubmitAIMediaAuditJobResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SubmitAIMediaAuditJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SubmitAIMediaAuditJobResponse) String() string {
	return tea.Prettify(s)
}

func (s SubmitAIMediaAuditJobResponse) GoString() string {
	return s.String()
}

func (s *SubmitAIMediaAuditJobResponse) SetHeaders(v map[string]*string) *SubmitAIMediaAuditJobResponse {
	s.Headers = v
	return s
}

func (s *SubmitAIMediaAuditJobResponse) SetBody(v *SubmitAIMediaAuditJobResponseBody) *SubmitAIMediaAuditJobResponse {
	s.Body = v
	return s
}

type SubmitDetectionJobRequest struct {
	BeginTime          *string `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	CopyrightBeginTime *string `json:"CopyrightBeginTime,omitempty" xml:"CopyrightBeginTime,omitempty"`
	CopyrightEndTime   *string `json:"CopyrightEndTime,omitempty" xml:"CopyrightEndTime,omitempty"`
	CopyrightFile      *string `json:"CopyrightFile,omitempty" xml:"CopyrightFile,omitempty"`
	CopyrightStatus    *string `json:"CopyrightStatus,omitempty" xml:"CopyrightStatus,omitempty"`
	Duration           *int32  `json:"Duration,omitempty" xml:"Duration,omitempty"`
	ShortVideo         *bool   `json:"ShortVideo,omitempty" xml:"ShortVideo,omitempty"`
	TemplateId         *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	VideoId            *string `json:"VideoId,omitempty" xml:"VideoId,omitempty"`
	WhiteListUrls      *string `json:"WhiteListUrls,omitempty" xml:"WhiteListUrls,omitempty"`
}

func (s SubmitDetectionJobRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitDetectionJobRequest) GoString() string {
	return s.String()
}

func (s *SubmitDetectionJobRequest) SetBeginTime(v string) *SubmitDetectionJobRequest {
	s.BeginTime = &v
	return s
}

func (s *SubmitDetectionJobRequest) SetCopyrightBeginTime(v string) *SubmitDetectionJobRequest {
	s.CopyrightBeginTime = &v
	return s
}

func (s *SubmitDetectionJobRequest) SetCopyrightEndTime(v string) *SubmitDetectionJobRequest {
	s.CopyrightEndTime = &v
	return s
}

func (s *SubmitDetectionJobRequest) SetCopyrightFile(v string) *SubmitDetectionJobRequest {
	s.CopyrightFile = &v
	return s
}

func (s *SubmitDetectionJobRequest) SetCopyrightStatus(v string) *SubmitDetectionJobRequest {
	s.CopyrightStatus = &v
	return s
}

func (s *SubmitDetectionJobRequest) SetDuration(v int32) *SubmitDetectionJobRequest {
	s.Duration = &v
	return s
}

func (s *SubmitDetectionJobRequest) SetShortVideo(v bool) *SubmitDetectionJobRequest {
	s.ShortVideo = &v
	return s
}

func (s *SubmitDetectionJobRequest) SetTemplateId(v string) *SubmitDetectionJobRequest {
	s.TemplateId = &v
	return s
}

func (s *SubmitDetectionJobRequest) SetVideoId(v string) *SubmitDetectionJobRequest {
	s.VideoId = &v
	return s
}

func (s *SubmitDetectionJobRequest) SetWhiteListUrls(v string) *SubmitDetectionJobRequest {
	s.WhiteListUrls = &v
	return s
}

type SubmitDetectionJobResponseBody struct {
	DetectionJob *SubmitDetectionJobResponseBodyDetectionJob `json:"DetectionJob,omitempty" xml:"DetectionJob,omitempty" type:"Struct"`
	RequestId    *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SubmitDetectionJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SubmitDetectionJobResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitDetectionJobResponseBody) SetDetectionJob(v *SubmitDetectionJobResponseBodyDetectionJob) *SubmitDetectionJobResponseBody {
	s.DetectionJob = v
	return s
}

func (s *SubmitDetectionJobResponseBody) SetRequestId(v string) *SubmitDetectionJobResponseBody {
	s.RequestId = &v
	return s
}

type SubmitDetectionJobResponseBodyDetectionJob struct {
	BeginTime          *string `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	CopyrightBeginTime *string `json:"CopyrightBeginTime,omitempty" xml:"CopyrightBeginTime,omitempty"`
	CopyrightEndTime   *string `json:"CopyrightEndTime,omitempty" xml:"CopyrightEndTime,omitempty"`
	CopyrightFile      *string `json:"CopyrightFile,omitempty" xml:"CopyrightFile,omitempty"`
	CopyrightStatus    *string `json:"CopyrightStatus,omitempty" xml:"CopyrightStatus,omitempty"`
	CreateTime         *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	EndTime            *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	JobId              *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	ModifyTime         *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	TemplateId         *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	VideoId            *string `json:"VideoId,omitempty" xml:"VideoId,omitempty"`
	WhitelistUrls      *string `json:"WhitelistUrls,omitempty" xml:"WhitelistUrls,omitempty"`
}

func (s SubmitDetectionJobResponseBodyDetectionJob) String() string {
	return tea.Prettify(s)
}

func (s SubmitDetectionJobResponseBodyDetectionJob) GoString() string {
	return s.String()
}

func (s *SubmitDetectionJobResponseBodyDetectionJob) SetBeginTime(v string) *SubmitDetectionJobResponseBodyDetectionJob {
	s.BeginTime = &v
	return s
}

func (s *SubmitDetectionJobResponseBodyDetectionJob) SetCopyrightBeginTime(v string) *SubmitDetectionJobResponseBodyDetectionJob {
	s.CopyrightBeginTime = &v
	return s
}

func (s *SubmitDetectionJobResponseBodyDetectionJob) SetCopyrightEndTime(v string) *SubmitDetectionJobResponseBodyDetectionJob {
	s.CopyrightEndTime = &v
	return s
}

func (s *SubmitDetectionJobResponseBodyDetectionJob) SetCopyrightFile(v string) *SubmitDetectionJobResponseBodyDetectionJob {
	s.CopyrightFile = &v
	return s
}

func (s *SubmitDetectionJobResponseBodyDetectionJob) SetCopyrightStatus(v string) *SubmitDetectionJobResponseBodyDetectionJob {
	s.CopyrightStatus = &v
	return s
}

func (s *SubmitDetectionJobResponseBodyDetectionJob) SetCreateTime(v string) *SubmitDetectionJobResponseBodyDetectionJob {
	s.CreateTime = &v
	return s
}

func (s *SubmitDetectionJobResponseBodyDetectionJob) SetEndTime(v string) *SubmitDetectionJobResponseBodyDetectionJob {
	s.EndTime = &v
	return s
}

func (s *SubmitDetectionJobResponseBodyDetectionJob) SetJobId(v string) *SubmitDetectionJobResponseBodyDetectionJob {
	s.JobId = &v
	return s
}

func (s *SubmitDetectionJobResponseBodyDetectionJob) SetModifyTime(v string) *SubmitDetectionJobResponseBodyDetectionJob {
	s.ModifyTime = &v
	return s
}

func (s *SubmitDetectionJobResponseBodyDetectionJob) SetTemplateId(v string) *SubmitDetectionJobResponseBodyDetectionJob {
	s.TemplateId = &v
	return s
}

func (s *SubmitDetectionJobResponseBodyDetectionJob) SetVideoId(v string) *SubmitDetectionJobResponseBodyDetectionJob {
	s.VideoId = &v
	return s
}

func (s *SubmitDetectionJobResponseBodyDetectionJob) SetWhitelistUrls(v string) *SubmitDetectionJobResponseBodyDetectionJob {
	s.WhitelistUrls = &v
	return s
}

type SubmitDetectionJobResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SubmitDetectionJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SubmitDetectionJobResponse) String() string {
	return tea.Prettify(s)
}

func (s SubmitDetectionJobResponse) GoString() string {
	return s.String()
}

func (s *SubmitDetectionJobResponse) SetHeaders(v map[string]*string) *SubmitDetectionJobResponse {
	s.Headers = v
	return s
}

func (s *SubmitDetectionJobResponse) SetBody(v *SubmitDetectionJobResponseBody) *SubmitDetectionJobResponse {
	s.Body = v
	return s
}

type SubmitDynamicImageJobRequest struct {
	DynamicImageTemplateId *string `json:"DynamicImageTemplateId,omitempty" xml:"DynamicImageTemplateId,omitempty"`
	OverrideParams         *string `json:"OverrideParams,omitempty" xml:"OverrideParams,omitempty"`
	VideoId                *string `json:"VideoId,omitempty" xml:"VideoId,omitempty"`
}

func (s SubmitDynamicImageJobRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitDynamicImageJobRequest) GoString() string {
	return s.String()
}

func (s *SubmitDynamicImageJobRequest) SetDynamicImageTemplateId(v string) *SubmitDynamicImageJobRequest {
	s.DynamicImageTemplateId = &v
	return s
}

func (s *SubmitDynamicImageJobRequest) SetOverrideParams(v string) *SubmitDynamicImageJobRequest {
	s.OverrideParams = &v
	return s
}

func (s *SubmitDynamicImageJobRequest) SetVideoId(v string) *SubmitDynamicImageJobRequest {
	s.VideoId = &v
	return s
}

type SubmitDynamicImageJobResponseBody struct {
	DynamicImageJob *SubmitDynamicImageJobResponseBodyDynamicImageJob `json:"DynamicImageJob,omitempty" xml:"DynamicImageJob,omitempty" type:"Struct"`
	RequestId       *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SubmitDynamicImageJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SubmitDynamicImageJobResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitDynamicImageJobResponseBody) SetDynamicImageJob(v *SubmitDynamicImageJobResponseBodyDynamicImageJob) *SubmitDynamicImageJobResponseBody {
	s.DynamicImageJob = v
	return s
}

func (s *SubmitDynamicImageJobResponseBody) SetRequestId(v string) *SubmitDynamicImageJobResponseBody {
	s.RequestId = &v
	return s
}

type SubmitDynamicImageJobResponseBodyDynamicImageJob struct {
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s SubmitDynamicImageJobResponseBodyDynamicImageJob) String() string {
	return tea.Prettify(s)
}

func (s SubmitDynamicImageJobResponseBodyDynamicImageJob) GoString() string {
	return s.String()
}

func (s *SubmitDynamicImageJobResponseBodyDynamicImageJob) SetJobId(v string) *SubmitDynamicImageJobResponseBodyDynamicImageJob {
	s.JobId = &v
	return s
}

type SubmitDynamicImageJobResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SubmitDynamicImageJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SubmitDynamicImageJobResponse) String() string {
	return tea.Prettify(s)
}

func (s SubmitDynamicImageJobResponse) GoString() string {
	return s.String()
}

func (s *SubmitDynamicImageJobResponse) SetHeaders(v map[string]*string) *SubmitDynamicImageJobResponse {
	s.Headers = v
	return s
}

func (s *SubmitDynamicImageJobResponse) SetBody(v *SubmitDynamicImageJobResponseBody) *SubmitDynamicImageJobResponse {
	s.Body = v
	return s
}

type SubmitLiveEditingRequest struct {
	AppName              *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	Clips                *string `json:"Clips,omitempty" xml:"Clips,omitempty"`
	CoverURL             *string `json:"CoverURL,omitempty" xml:"CoverURL,omitempty"`
	Description          *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DomainName           *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	MediaMetadata        *string `json:"MediaMetadata,omitempty" xml:"MediaMetadata,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ProduceConfig        *string `json:"ProduceConfig,omitempty" xml:"ProduceConfig,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	StreamName           *string `json:"StreamName,omitempty" xml:"StreamName,omitempty"`
	Title                *string `json:"Title,omitempty" xml:"Title,omitempty"`
	UserData             *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s SubmitLiveEditingRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitLiveEditingRequest) GoString() string {
	return s.String()
}

func (s *SubmitLiveEditingRequest) SetAppName(v string) *SubmitLiveEditingRequest {
	s.AppName = &v
	return s
}

func (s *SubmitLiveEditingRequest) SetClips(v string) *SubmitLiveEditingRequest {
	s.Clips = &v
	return s
}

func (s *SubmitLiveEditingRequest) SetCoverURL(v string) *SubmitLiveEditingRequest {
	s.CoverURL = &v
	return s
}

func (s *SubmitLiveEditingRequest) SetDescription(v string) *SubmitLiveEditingRequest {
	s.Description = &v
	return s
}

func (s *SubmitLiveEditingRequest) SetDomainName(v string) *SubmitLiveEditingRequest {
	s.DomainName = &v
	return s
}

func (s *SubmitLiveEditingRequest) SetMediaMetadata(v string) *SubmitLiveEditingRequest {
	s.MediaMetadata = &v
	return s
}

func (s *SubmitLiveEditingRequest) SetOwnerId(v int64) *SubmitLiveEditingRequest {
	s.OwnerId = &v
	return s
}

func (s *SubmitLiveEditingRequest) SetProduceConfig(v string) *SubmitLiveEditingRequest {
	s.ProduceConfig = &v
	return s
}

func (s *SubmitLiveEditingRequest) SetResourceOwnerAccount(v string) *SubmitLiveEditingRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SubmitLiveEditingRequest) SetResourceOwnerId(v int64) *SubmitLiveEditingRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SubmitLiveEditingRequest) SetStreamName(v string) *SubmitLiveEditingRequest {
	s.StreamName = &v
	return s
}

func (s *SubmitLiveEditingRequest) SetTitle(v string) *SubmitLiveEditingRequest {
	s.Title = &v
	return s
}

func (s *SubmitLiveEditingRequest) SetUserData(v string) *SubmitLiveEditingRequest {
	s.UserData = &v
	return s
}

type SubmitLiveEditingResponseBody struct {
	MediaId   *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SubmitLiveEditingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SubmitLiveEditingResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitLiveEditingResponseBody) SetMediaId(v string) *SubmitLiveEditingResponseBody {
	s.MediaId = &v
	return s
}

func (s *SubmitLiveEditingResponseBody) SetProjectId(v string) *SubmitLiveEditingResponseBody {
	s.ProjectId = &v
	return s
}

func (s *SubmitLiveEditingResponseBody) SetRequestId(v string) *SubmitLiveEditingResponseBody {
	s.RequestId = &v
	return s
}

type SubmitLiveEditingResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SubmitLiveEditingResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SubmitLiveEditingResponse) String() string {
	return tea.Prettify(s)
}

func (s SubmitLiveEditingResponse) GoString() string {
	return s.String()
}

func (s *SubmitLiveEditingResponse) SetHeaders(v map[string]*string) *SubmitLiveEditingResponse {
	s.Headers = v
	return s
}

func (s *SubmitLiveEditingResponse) SetBody(v *SubmitLiveEditingResponseBody) *SubmitLiveEditingResponse {
	s.Body = v
	return s
}

type SubmitMediaDNADeleteJobRequest struct {
	MediaId              *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *string `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s SubmitMediaDNADeleteJobRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitMediaDNADeleteJobRequest) GoString() string {
	return s.String()
}

func (s *SubmitMediaDNADeleteJobRequest) SetMediaId(v string) *SubmitMediaDNADeleteJobRequest {
	s.MediaId = &v
	return s
}

func (s *SubmitMediaDNADeleteJobRequest) SetOwnerAccount(v string) *SubmitMediaDNADeleteJobRequest {
	s.OwnerAccount = &v
	return s
}

func (s *SubmitMediaDNADeleteJobRequest) SetOwnerId(v string) *SubmitMediaDNADeleteJobRequest {
	s.OwnerId = &v
	return s
}

func (s *SubmitMediaDNADeleteJobRequest) SetResourceOwnerAccount(v string) *SubmitMediaDNADeleteJobRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SubmitMediaDNADeleteJobRequest) SetResourceOwnerId(v string) *SubmitMediaDNADeleteJobRequest {
	s.ResourceOwnerId = &v
	return s
}

type SubmitMediaDNADeleteJobResponseBody struct {
	JobId     *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SubmitMediaDNADeleteJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SubmitMediaDNADeleteJobResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitMediaDNADeleteJobResponseBody) SetJobId(v string) *SubmitMediaDNADeleteJobResponseBody {
	s.JobId = &v
	return s
}

func (s *SubmitMediaDNADeleteJobResponseBody) SetRequestId(v string) *SubmitMediaDNADeleteJobResponseBody {
	s.RequestId = &v
	return s
}

type SubmitMediaDNADeleteJobResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SubmitMediaDNADeleteJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SubmitMediaDNADeleteJobResponse) String() string {
	return tea.Prettify(s)
}

func (s SubmitMediaDNADeleteJobResponse) GoString() string {
	return s.String()
}

func (s *SubmitMediaDNADeleteJobResponse) SetHeaders(v map[string]*string) *SubmitMediaDNADeleteJobResponse {
	s.Headers = v
	return s
}

func (s *SubmitMediaDNADeleteJobResponse) SetBody(v *SubmitMediaDNADeleteJobResponseBody) *SubmitMediaDNADeleteJobResponse {
	s.Body = v
	return s
}

type SubmitPreprocessJobsRequest struct {
	PreprocessType *string `json:"PreprocessType,omitempty" xml:"PreprocessType,omitempty"`
	VideoId        *string `json:"VideoId,omitempty" xml:"VideoId,omitempty"`
}

func (s SubmitPreprocessJobsRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitPreprocessJobsRequest) GoString() string {
	return s.String()
}

func (s *SubmitPreprocessJobsRequest) SetPreprocessType(v string) *SubmitPreprocessJobsRequest {
	s.PreprocessType = &v
	return s
}

func (s *SubmitPreprocessJobsRequest) SetVideoId(v string) *SubmitPreprocessJobsRequest {
	s.VideoId = &v
	return s
}

type SubmitPreprocessJobsResponseBody struct {
	PreprocessJobs *SubmitPreprocessJobsResponseBodyPreprocessJobs `json:"PreprocessJobs,omitempty" xml:"PreprocessJobs,omitempty" type:"Struct"`
	RequestId      *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SubmitPreprocessJobsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SubmitPreprocessJobsResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitPreprocessJobsResponseBody) SetPreprocessJobs(v *SubmitPreprocessJobsResponseBodyPreprocessJobs) *SubmitPreprocessJobsResponseBody {
	s.PreprocessJobs = v
	return s
}

func (s *SubmitPreprocessJobsResponseBody) SetRequestId(v string) *SubmitPreprocessJobsResponseBody {
	s.RequestId = &v
	return s
}

type SubmitPreprocessJobsResponseBodyPreprocessJobs struct {
	PreprocessJob []*SubmitPreprocessJobsResponseBodyPreprocessJobsPreprocessJob `json:"PreprocessJob,omitempty" xml:"PreprocessJob,omitempty" type:"Repeated"`
}

func (s SubmitPreprocessJobsResponseBodyPreprocessJobs) String() string {
	return tea.Prettify(s)
}

func (s SubmitPreprocessJobsResponseBodyPreprocessJobs) GoString() string {
	return s.String()
}

func (s *SubmitPreprocessJobsResponseBodyPreprocessJobs) SetPreprocessJob(v []*SubmitPreprocessJobsResponseBodyPreprocessJobsPreprocessJob) *SubmitPreprocessJobsResponseBodyPreprocessJobs {
	s.PreprocessJob = v
	return s
}

type SubmitPreprocessJobsResponseBodyPreprocessJobsPreprocessJob struct {
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s SubmitPreprocessJobsResponseBodyPreprocessJobsPreprocessJob) String() string {
	return tea.Prettify(s)
}

func (s SubmitPreprocessJobsResponseBodyPreprocessJobsPreprocessJob) GoString() string {
	return s.String()
}

func (s *SubmitPreprocessJobsResponseBodyPreprocessJobsPreprocessJob) SetJobId(v string) *SubmitPreprocessJobsResponseBodyPreprocessJobsPreprocessJob {
	s.JobId = &v
	return s
}

type SubmitPreprocessJobsResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SubmitPreprocessJobsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SubmitPreprocessJobsResponse) String() string {
	return tea.Prettify(s)
}

func (s SubmitPreprocessJobsResponse) GoString() string {
	return s.String()
}

func (s *SubmitPreprocessJobsResponse) SetHeaders(v map[string]*string) *SubmitPreprocessJobsResponse {
	s.Headers = v
	return s
}

func (s *SubmitPreprocessJobsResponse) SetBody(v *SubmitPreprocessJobsResponseBody) *SubmitPreprocessJobsResponse {
	s.Body = v
	return s
}

type SubmitSnapshotJobRequest struct {
	Count                *int64  `json:"Count,omitempty" xml:"Count,omitempty"`
	Height               *string `json:"Height,omitempty" xml:"Height,omitempty"`
	Interval             *int64  `json:"Interval,omitempty" xml:"Interval,omitempty"`
	SnapshotTemplateId   *string `json:"SnapshotTemplateId,omitempty" xml:"SnapshotTemplateId,omitempty"`
	SpecifiedOffsetTime  *int64  `json:"SpecifiedOffsetTime,omitempty" xml:"SpecifiedOffsetTime,omitempty"`
	SpriteSnapshotConfig *string `json:"SpriteSnapshotConfig,omitempty" xml:"SpriteSnapshotConfig,omitempty"`
	UserData             *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
	VideoId              *string `json:"VideoId,omitempty" xml:"VideoId,omitempty"`
	Width                *string `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s SubmitSnapshotJobRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitSnapshotJobRequest) GoString() string {
	return s.String()
}

func (s *SubmitSnapshotJobRequest) SetCount(v int64) *SubmitSnapshotJobRequest {
	s.Count = &v
	return s
}

func (s *SubmitSnapshotJobRequest) SetHeight(v string) *SubmitSnapshotJobRequest {
	s.Height = &v
	return s
}

func (s *SubmitSnapshotJobRequest) SetInterval(v int64) *SubmitSnapshotJobRequest {
	s.Interval = &v
	return s
}

func (s *SubmitSnapshotJobRequest) SetSnapshotTemplateId(v string) *SubmitSnapshotJobRequest {
	s.SnapshotTemplateId = &v
	return s
}

func (s *SubmitSnapshotJobRequest) SetSpecifiedOffsetTime(v int64) *SubmitSnapshotJobRequest {
	s.SpecifiedOffsetTime = &v
	return s
}

func (s *SubmitSnapshotJobRequest) SetSpriteSnapshotConfig(v string) *SubmitSnapshotJobRequest {
	s.SpriteSnapshotConfig = &v
	return s
}

func (s *SubmitSnapshotJobRequest) SetUserData(v string) *SubmitSnapshotJobRequest {
	s.UserData = &v
	return s
}

func (s *SubmitSnapshotJobRequest) SetVideoId(v string) *SubmitSnapshotJobRequest {
	s.VideoId = &v
	return s
}

func (s *SubmitSnapshotJobRequest) SetWidth(v string) *SubmitSnapshotJobRequest {
	s.Width = &v
	return s
}

type SubmitSnapshotJobResponseBody struct {
	RequestId   *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SnapshotJob *SubmitSnapshotJobResponseBodySnapshotJob `json:"SnapshotJob,omitempty" xml:"SnapshotJob,omitempty" type:"Struct"`
}

func (s SubmitSnapshotJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SubmitSnapshotJobResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitSnapshotJobResponseBody) SetRequestId(v string) *SubmitSnapshotJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitSnapshotJobResponseBody) SetSnapshotJob(v *SubmitSnapshotJobResponseBodySnapshotJob) *SubmitSnapshotJobResponseBody {
	s.SnapshotJob = v
	return s
}

type SubmitSnapshotJobResponseBodySnapshotJob struct {
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s SubmitSnapshotJobResponseBodySnapshotJob) String() string {
	return tea.Prettify(s)
}

func (s SubmitSnapshotJobResponseBodySnapshotJob) GoString() string {
	return s.String()
}

func (s *SubmitSnapshotJobResponseBodySnapshotJob) SetJobId(v string) *SubmitSnapshotJobResponseBodySnapshotJob {
	s.JobId = &v
	return s
}

type SubmitSnapshotJobResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SubmitSnapshotJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SubmitSnapshotJobResponse) String() string {
	return tea.Prettify(s)
}

func (s SubmitSnapshotJobResponse) GoString() string {
	return s.String()
}

func (s *SubmitSnapshotJobResponse) SetHeaders(v map[string]*string) *SubmitSnapshotJobResponse {
	s.Headers = v
	return s
}

func (s *SubmitSnapshotJobResponse) SetBody(v *SubmitSnapshotJobResponseBody) *SubmitSnapshotJobResponse {
	s.Body = v
	return s
}

type SubmitTranscodeJobsRequest struct {
	EncryptConfig   *string `json:"EncryptConfig,omitempty" xml:"EncryptConfig,omitempty"`
	OverrideParams  *string `json:"OverrideParams,omitempty" xml:"OverrideParams,omitempty"`
	PipelineId      *string `json:"PipelineId,omitempty" xml:"PipelineId,omitempty"`
	Priority        *string `json:"Priority,omitempty" xml:"Priority,omitempty"`
	TemplateGroupId *string `json:"TemplateGroupId,omitempty" xml:"TemplateGroupId,omitempty"`
	UserData        *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
	VideoId         *string `json:"VideoId,omitempty" xml:"VideoId,omitempty"`
}

func (s SubmitTranscodeJobsRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitTranscodeJobsRequest) GoString() string {
	return s.String()
}

func (s *SubmitTranscodeJobsRequest) SetEncryptConfig(v string) *SubmitTranscodeJobsRequest {
	s.EncryptConfig = &v
	return s
}

func (s *SubmitTranscodeJobsRequest) SetOverrideParams(v string) *SubmitTranscodeJobsRequest {
	s.OverrideParams = &v
	return s
}

func (s *SubmitTranscodeJobsRequest) SetPipelineId(v string) *SubmitTranscodeJobsRequest {
	s.PipelineId = &v
	return s
}

func (s *SubmitTranscodeJobsRequest) SetPriority(v string) *SubmitTranscodeJobsRequest {
	s.Priority = &v
	return s
}

func (s *SubmitTranscodeJobsRequest) SetTemplateGroupId(v string) *SubmitTranscodeJobsRequest {
	s.TemplateGroupId = &v
	return s
}

func (s *SubmitTranscodeJobsRequest) SetUserData(v string) *SubmitTranscodeJobsRequest {
	s.UserData = &v
	return s
}

func (s *SubmitTranscodeJobsRequest) SetVideoId(v string) *SubmitTranscodeJobsRequest {
	s.VideoId = &v
	return s
}

type SubmitTranscodeJobsResponseBody struct {
	RequestId       *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TranscodeJobs   *SubmitTranscodeJobsResponseBodyTranscodeJobs `json:"TranscodeJobs,omitempty" xml:"TranscodeJobs,omitempty" type:"Struct"`
	TranscodeTaskId *string                                       `json:"TranscodeTaskId,omitempty" xml:"TranscodeTaskId,omitempty"`
}

func (s SubmitTranscodeJobsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SubmitTranscodeJobsResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitTranscodeJobsResponseBody) SetRequestId(v string) *SubmitTranscodeJobsResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitTranscodeJobsResponseBody) SetTranscodeJobs(v *SubmitTranscodeJobsResponseBodyTranscodeJobs) *SubmitTranscodeJobsResponseBody {
	s.TranscodeJobs = v
	return s
}

func (s *SubmitTranscodeJobsResponseBody) SetTranscodeTaskId(v string) *SubmitTranscodeJobsResponseBody {
	s.TranscodeTaskId = &v
	return s
}

type SubmitTranscodeJobsResponseBodyTranscodeJobs struct {
	TranscodeJob []*SubmitTranscodeJobsResponseBodyTranscodeJobsTranscodeJob `json:"TranscodeJob,omitempty" xml:"TranscodeJob,omitempty" type:"Repeated"`
}

func (s SubmitTranscodeJobsResponseBodyTranscodeJobs) String() string {
	return tea.Prettify(s)
}

func (s SubmitTranscodeJobsResponseBodyTranscodeJobs) GoString() string {
	return s.String()
}

func (s *SubmitTranscodeJobsResponseBodyTranscodeJobs) SetTranscodeJob(v []*SubmitTranscodeJobsResponseBodyTranscodeJobsTranscodeJob) *SubmitTranscodeJobsResponseBodyTranscodeJobs {
	s.TranscodeJob = v
	return s
}

type SubmitTranscodeJobsResponseBodyTranscodeJobsTranscodeJob struct {
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s SubmitTranscodeJobsResponseBodyTranscodeJobsTranscodeJob) String() string {
	return tea.Prettify(s)
}

func (s SubmitTranscodeJobsResponseBodyTranscodeJobsTranscodeJob) GoString() string {
	return s.String()
}

func (s *SubmitTranscodeJobsResponseBodyTranscodeJobsTranscodeJob) SetJobId(v string) *SubmitTranscodeJobsResponseBodyTranscodeJobsTranscodeJob {
	s.JobId = &v
	return s
}

type SubmitTranscodeJobsResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SubmitTranscodeJobsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SubmitTranscodeJobsResponse) String() string {
	return tea.Prettify(s)
}

func (s SubmitTranscodeJobsResponse) GoString() string {
	return s.String()
}

func (s *SubmitTranscodeJobsResponse) SetHeaders(v map[string]*string) *SubmitTranscodeJobsResponse {
	s.Headers = v
	return s
}

func (s *SubmitTranscodeJobsResponse) SetBody(v *SubmitTranscodeJobsResponseBody) *SubmitTranscodeJobsResponse {
	s.Body = v
	return s
}

type SubmitWorkflowJobRequest struct {
	MediaId    *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	WorkflowId *string `json:"WorkflowId,omitempty" xml:"WorkflowId,omitempty"`
}

func (s SubmitWorkflowJobRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitWorkflowJobRequest) GoString() string {
	return s.String()
}

func (s *SubmitWorkflowJobRequest) SetMediaId(v string) *SubmitWorkflowJobRequest {
	s.MediaId = &v
	return s
}

func (s *SubmitWorkflowJobRequest) SetWorkflowId(v string) *SubmitWorkflowJobRequest {
	s.WorkflowId = &v
	return s
}

type SubmitWorkflowJobResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SubmitWorkflowJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SubmitWorkflowJobResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitWorkflowJobResponseBody) SetRequestId(v string) *SubmitWorkflowJobResponseBody {
	s.RequestId = &v
	return s
}

type SubmitWorkflowJobResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SubmitWorkflowJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SubmitWorkflowJobResponse) String() string {
	return tea.Prettify(s)
}

func (s SubmitWorkflowJobResponse) GoString() string {
	return s.String()
}

func (s *SubmitWorkflowJobResponse) SetHeaders(v map[string]*string) *SubmitWorkflowJobResponse {
	s.Headers = v
	return s
}

func (s *SubmitWorkflowJobResponse) SetBody(v *SubmitWorkflowJobResponseBody) *SubmitWorkflowJobResponse {
	s.Body = v
	return s
}

type TagVodResourcesRequest struct {
	OwnerId      *int64                       `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceId   []*string                    `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	ResourceType *string                      `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Tag          []*TagVodResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s TagVodResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s TagVodResourcesRequest) GoString() string {
	return s.String()
}

func (s *TagVodResourcesRequest) SetOwnerId(v int64) *TagVodResourcesRequest {
	s.OwnerId = &v
	return s
}

func (s *TagVodResourcesRequest) SetResourceId(v []*string) *TagVodResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *TagVodResourcesRequest) SetResourceType(v string) *TagVodResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *TagVodResourcesRequest) SetTag(v []*TagVodResourcesRequestTag) *TagVodResourcesRequest {
	s.Tag = v
	return s
}

type TagVodResourcesRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s TagVodResourcesRequestTag) String() string {
	return tea.Prettify(s)
}

func (s TagVodResourcesRequestTag) GoString() string {
	return s.String()
}

func (s *TagVodResourcesRequestTag) SetKey(v string) *TagVodResourcesRequestTag {
	s.Key = &v
	return s
}

func (s *TagVodResourcesRequestTag) SetValue(v string) *TagVodResourcesRequestTag {
	s.Value = &v
	return s
}

type TagVodResourcesResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s TagVodResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TagVodResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *TagVodResourcesResponseBody) SetRequestId(v string) *TagVodResourcesResponseBody {
	s.RequestId = &v
	return s
}

type TagVodResourcesResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *TagVodResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s TagVodResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s TagVodResourcesResponse) GoString() string {
	return s.String()
}

func (s *TagVodResourcesResponse) SetHeaders(v map[string]*string) *TagVodResourcesResponse {
	s.Headers = v
	return s
}

func (s *TagVodResourcesResponse) SetBody(v *TagVodResourcesResponseBody) *TagVodResourcesResponse {
	s.Body = v
	return s
}

type UnTagVodResourcesRequest struct {
	All          *bool     `json:"All,omitempty" xml:"All,omitempty"`
	OwnerId      *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceId   []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	ResourceType *string   `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	TagKey       []*string `json:"TagKey,omitempty" xml:"TagKey,omitempty" type:"Repeated"`
}

func (s UnTagVodResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s UnTagVodResourcesRequest) GoString() string {
	return s.String()
}

func (s *UnTagVodResourcesRequest) SetAll(v bool) *UnTagVodResourcesRequest {
	s.All = &v
	return s
}

func (s *UnTagVodResourcesRequest) SetOwnerId(v int64) *UnTagVodResourcesRequest {
	s.OwnerId = &v
	return s
}

func (s *UnTagVodResourcesRequest) SetResourceId(v []*string) *UnTagVodResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *UnTagVodResourcesRequest) SetResourceType(v string) *UnTagVodResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *UnTagVodResourcesRequest) SetTagKey(v []*string) *UnTagVodResourcesRequest {
	s.TagKey = v
	return s
}

type UnTagVodResourcesResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UnTagVodResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UnTagVodResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *UnTagVodResourcesResponseBody) SetRequestId(v string) *UnTagVodResourcesResponseBody {
	s.RequestId = &v
	return s
}

type UnTagVodResourcesResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UnTagVodResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UnTagVodResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s UnTagVodResourcesResponse) GoString() string {
	return s.String()
}

func (s *UnTagVodResourcesResponse) SetHeaders(v map[string]*string) *UnTagVodResourcesResponse {
	s.Headers = v
	return s
}

func (s *UnTagVodResourcesResponse) SetBody(v *UnTagVodResourcesResponseBody) *UnTagVodResourcesResponse {
	s.Body = v
	return s
}

type UpdateAITemplateRequest struct {
	TemplateConfig *string `json:"TemplateConfig,omitempty" xml:"TemplateConfig,omitempty"`
	TemplateId     *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	TemplateName   *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
}

func (s UpdateAITemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateAITemplateRequest) GoString() string {
	return s.String()
}

func (s *UpdateAITemplateRequest) SetTemplateConfig(v string) *UpdateAITemplateRequest {
	s.TemplateConfig = &v
	return s
}

func (s *UpdateAITemplateRequest) SetTemplateId(v string) *UpdateAITemplateRequest {
	s.TemplateId = &v
	return s
}

func (s *UpdateAITemplateRequest) SetTemplateName(v string) *UpdateAITemplateRequest {
	s.TemplateName = &v
	return s
}

type UpdateAITemplateResponseBody struct {
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s UpdateAITemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateAITemplateResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAITemplateResponseBody) SetRequestId(v string) *UpdateAITemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAITemplateResponseBody) SetTemplateId(v string) *UpdateAITemplateResponseBody {
	s.TemplateId = &v
	return s
}

type UpdateAITemplateResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateAITemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateAITemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateAITemplateResponse) GoString() string {
	return s.String()
}

func (s *UpdateAITemplateResponse) SetHeaders(v map[string]*string) *UpdateAITemplateResponse {
	s.Headers = v
	return s
}

func (s *UpdateAITemplateResponse) SetBody(v *UpdateAITemplateResponseBody) *UpdateAITemplateResponse {
	s.Body = v
	return s
}

type UpdateAppInfoRequest struct {
	AppId       *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AppName     *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Status      *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s UpdateAppInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateAppInfoRequest) GoString() string {
	return s.String()
}

func (s *UpdateAppInfoRequest) SetAppId(v string) *UpdateAppInfoRequest {
	s.AppId = &v
	return s
}

func (s *UpdateAppInfoRequest) SetAppName(v string) *UpdateAppInfoRequest {
	s.AppName = &v
	return s
}

func (s *UpdateAppInfoRequest) SetDescription(v string) *UpdateAppInfoRequest {
	s.Description = &v
	return s
}

func (s *UpdateAppInfoRequest) SetStatus(v string) *UpdateAppInfoRequest {
	s.Status = &v
	return s
}

type UpdateAppInfoResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateAppInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateAppInfoResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAppInfoResponseBody) SetRequestId(v string) *UpdateAppInfoResponseBody {
	s.RequestId = &v
	return s
}

type UpdateAppInfoResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateAppInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateAppInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateAppInfoResponse) GoString() string {
	return s.String()
}

func (s *UpdateAppInfoResponse) SetHeaders(v map[string]*string) *UpdateAppInfoResponse {
	s.Headers = v
	return s
}

func (s *UpdateAppInfoResponse) SetBody(v *UpdateAppInfoResponseBody) *UpdateAppInfoResponse {
	s.Body = v
	return s
}

type UpdateAttachedMediaInfosRequest struct {
	UpdateContent *string `json:"UpdateContent,omitempty" xml:"UpdateContent,omitempty"`
}

func (s UpdateAttachedMediaInfosRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateAttachedMediaInfosRequest) GoString() string {
	return s.String()
}

func (s *UpdateAttachedMediaInfosRequest) SetUpdateContent(v string) *UpdateAttachedMediaInfosRequest {
	s.UpdateContent = &v
	return s
}

type UpdateAttachedMediaInfosResponseBody struct {
	NonExistMediaIds []*string `json:"NonExistMediaIds,omitempty" xml:"NonExistMediaIds,omitempty" type:"Repeated"`
	RequestId        *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateAttachedMediaInfosResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateAttachedMediaInfosResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAttachedMediaInfosResponseBody) SetNonExistMediaIds(v []*string) *UpdateAttachedMediaInfosResponseBody {
	s.NonExistMediaIds = v
	return s
}

func (s *UpdateAttachedMediaInfosResponseBody) SetRequestId(v string) *UpdateAttachedMediaInfosResponseBody {
	s.RequestId = &v
	return s
}

type UpdateAttachedMediaInfosResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateAttachedMediaInfosResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateAttachedMediaInfosResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateAttachedMediaInfosResponse) GoString() string {
	return s.String()
}

func (s *UpdateAttachedMediaInfosResponse) SetHeaders(v map[string]*string) *UpdateAttachedMediaInfosResponse {
	s.Headers = v
	return s
}

func (s *UpdateAttachedMediaInfosResponse) SetBody(v *UpdateAttachedMediaInfosResponseBody) *UpdateAttachedMediaInfosResponse {
	s.Body = v
	return s
}

type UpdateCategoryRequest struct {
	CateId   *int64  `json:"CateId,omitempty" xml:"CateId,omitempty"`
	CateName *string `json:"CateName,omitempty" xml:"CateName,omitempty"`
}

func (s UpdateCategoryRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateCategoryRequest) GoString() string {
	return s.String()
}

func (s *UpdateCategoryRequest) SetCateId(v int64) *UpdateCategoryRequest {
	s.CateId = &v
	return s
}

func (s *UpdateCategoryRequest) SetCateName(v string) *UpdateCategoryRequest {
	s.CateName = &v
	return s
}

type UpdateCategoryResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateCategoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateCategoryResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateCategoryResponseBody) SetRequestId(v string) *UpdateCategoryResponseBody {
	s.RequestId = &v
	return s
}

type UpdateCategoryResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateCategoryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateCategoryResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateCategoryResponse) GoString() string {
	return s.String()
}

func (s *UpdateCategoryResponse) SetHeaders(v map[string]*string) *UpdateCategoryResponse {
	s.Headers = v
	return s
}

func (s *UpdateCategoryResponse) SetBody(v *UpdateCategoryResponseBody) *UpdateCategoryResponse {
	s.Body = v
	return s
}

type UpdateDetectionJobRequest struct {
	BeginTime          *string `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	CopyrightBeginTime *string `json:"CopyrightBeginTime,omitempty" xml:"CopyrightBeginTime,omitempty"`
	CopyrightEndTime   *string `json:"CopyrightEndTime,omitempty" xml:"CopyrightEndTime,omitempty"`
	CopyrightFile      *string `json:"CopyrightFile,omitempty" xml:"CopyrightFile,omitempty"`
	CopyrightStatus    *string `json:"CopyrightStatus,omitempty" xml:"CopyrightStatus,omitempty"`
	Duration           *int32  `json:"Duration,omitempty" xml:"Duration,omitempty"`
	JobId              *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	TemplateId         *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	WhiteListUrls      *string `json:"WhiteListUrls,omitempty" xml:"WhiteListUrls,omitempty"`
}

func (s UpdateDetectionJobRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateDetectionJobRequest) GoString() string {
	return s.String()
}

func (s *UpdateDetectionJobRequest) SetBeginTime(v string) *UpdateDetectionJobRequest {
	s.BeginTime = &v
	return s
}

func (s *UpdateDetectionJobRequest) SetCopyrightBeginTime(v string) *UpdateDetectionJobRequest {
	s.CopyrightBeginTime = &v
	return s
}

func (s *UpdateDetectionJobRequest) SetCopyrightEndTime(v string) *UpdateDetectionJobRequest {
	s.CopyrightEndTime = &v
	return s
}

func (s *UpdateDetectionJobRequest) SetCopyrightFile(v string) *UpdateDetectionJobRequest {
	s.CopyrightFile = &v
	return s
}

func (s *UpdateDetectionJobRequest) SetCopyrightStatus(v string) *UpdateDetectionJobRequest {
	s.CopyrightStatus = &v
	return s
}

func (s *UpdateDetectionJobRequest) SetDuration(v int32) *UpdateDetectionJobRequest {
	s.Duration = &v
	return s
}

func (s *UpdateDetectionJobRequest) SetJobId(v string) *UpdateDetectionJobRequest {
	s.JobId = &v
	return s
}

func (s *UpdateDetectionJobRequest) SetTemplateId(v string) *UpdateDetectionJobRequest {
	s.TemplateId = &v
	return s
}

func (s *UpdateDetectionJobRequest) SetWhiteListUrls(v string) *UpdateDetectionJobRequest {
	s.WhiteListUrls = &v
	return s
}

type UpdateDetectionJobResponseBody struct {
	DetectionJob *UpdateDetectionJobResponseBodyDetectionJob `json:"DetectionJob,omitempty" xml:"DetectionJob,omitempty" type:"Struct"`
	RequestId    *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateDetectionJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateDetectionJobResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDetectionJobResponseBody) SetDetectionJob(v *UpdateDetectionJobResponseBodyDetectionJob) *UpdateDetectionJobResponseBody {
	s.DetectionJob = v
	return s
}

func (s *UpdateDetectionJobResponseBody) SetRequestId(v string) *UpdateDetectionJobResponseBody {
	s.RequestId = &v
	return s
}

type UpdateDetectionJobResponseBodyDetectionJob struct {
	BeginTime          *string `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	CopyrightBeginTime *string `json:"CopyrightBeginTime,omitempty" xml:"CopyrightBeginTime,omitempty"`
	CopyrightEndTime   *string `json:"CopyrightEndTime,omitempty" xml:"CopyrightEndTime,omitempty"`
	CopyrightFile      *string `json:"CopyrightFile,omitempty" xml:"CopyrightFile,omitempty"`
	CopyrightStatus    *string `json:"CopyrightStatus,omitempty" xml:"CopyrightStatus,omitempty"`
	CreateTime         *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	EndTime            *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	JobId              *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	ModifyTime         *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	TemplateId         *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	VideoId            *string `json:"VideoId,omitempty" xml:"VideoId,omitempty"`
	WhitelistUrls      *string `json:"WhitelistUrls,omitempty" xml:"WhitelistUrls,omitempty"`
}

func (s UpdateDetectionJobResponseBodyDetectionJob) String() string {
	return tea.Prettify(s)
}

func (s UpdateDetectionJobResponseBodyDetectionJob) GoString() string {
	return s.String()
}

func (s *UpdateDetectionJobResponseBodyDetectionJob) SetBeginTime(v string) *UpdateDetectionJobResponseBodyDetectionJob {
	s.BeginTime = &v
	return s
}

func (s *UpdateDetectionJobResponseBodyDetectionJob) SetCopyrightBeginTime(v string) *UpdateDetectionJobResponseBodyDetectionJob {
	s.CopyrightBeginTime = &v
	return s
}

func (s *UpdateDetectionJobResponseBodyDetectionJob) SetCopyrightEndTime(v string) *UpdateDetectionJobResponseBodyDetectionJob {
	s.CopyrightEndTime = &v
	return s
}

func (s *UpdateDetectionJobResponseBodyDetectionJob) SetCopyrightFile(v string) *UpdateDetectionJobResponseBodyDetectionJob {
	s.CopyrightFile = &v
	return s
}

func (s *UpdateDetectionJobResponseBodyDetectionJob) SetCopyrightStatus(v string) *UpdateDetectionJobResponseBodyDetectionJob {
	s.CopyrightStatus = &v
	return s
}

func (s *UpdateDetectionJobResponseBodyDetectionJob) SetCreateTime(v string) *UpdateDetectionJobResponseBodyDetectionJob {
	s.CreateTime = &v
	return s
}

func (s *UpdateDetectionJobResponseBodyDetectionJob) SetEndTime(v string) *UpdateDetectionJobResponseBodyDetectionJob {
	s.EndTime = &v
	return s
}

func (s *UpdateDetectionJobResponseBodyDetectionJob) SetJobId(v string) *UpdateDetectionJobResponseBodyDetectionJob {
	s.JobId = &v
	return s
}

func (s *UpdateDetectionJobResponseBodyDetectionJob) SetModifyTime(v string) *UpdateDetectionJobResponseBodyDetectionJob {
	s.ModifyTime = &v
	return s
}

func (s *UpdateDetectionJobResponseBodyDetectionJob) SetTemplateId(v string) *UpdateDetectionJobResponseBodyDetectionJob {
	s.TemplateId = &v
	return s
}

func (s *UpdateDetectionJobResponseBodyDetectionJob) SetVideoId(v string) *UpdateDetectionJobResponseBodyDetectionJob {
	s.VideoId = &v
	return s
}

func (s *UpdateDetectionJobResponseBodyDetectionJob) SetWhitelistUrls(v string) *UpdateDetectionJobResponseBodyDetectionJob {
	s.WhitelistUrls = &v
	return s
}

type UpdateDetectionJobResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateDetectionJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateDetectionJobResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateDetectionJobResponse) GoString() string {
	return s.String()
}

func (s *UpdateDetectionJobResponse) SetHeaders(v map[string]*string) *UpdateDetectionJobResponse {
	s.Headers = v
	return s
}

func (s *UpdateDetectionJobResponse) SetBody(v *UpdateDetectionJobResponseBody) *UpdateDetectionJobResponse {
	s.Body = v
	return s
}

type UpdateDetectionTemplateRequest struct {
	Period       *string `json:"Period,omitempty" xml:"Period,omitempty"`
	Platform     *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	TemplateId   *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
}

func (s UpdateDetectionTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateDetectionTemplateRequest) GoString() string {
	return s.String()
}

func (s *UpdateDetectionTemplateRequest) SetPeriod(v string) *UpdateDetectionTemplateRequest {
	s.Period = &v
	return s
}

func (s *UpdateDetectionTemplateRequest) SetPlatform(v string) *UpdateDetectionTemplateRequest {
	s.Platform = &v
	return s
}

func (s *UpdateDetectionTemplateRequest) SetTemplateId(v string) *UpdateDetectionTemplateRequest {
	s.TemplateId = &v
	return s
}

func (s *UpdateDetectionTemplateRequest) SetTemplateName(v string) *UpdateDetectionTemplateRequest {
	s.TemplateName = &v
	return s
}

type UpdateDetectionTemplateResponseBody struct {
	DetectionTemplate *UpdateDetectionTemplateResponseBodyDetectionTemplate `json:"DetectionTemplate,omitempty" xml:"DetectionTemplate,omitempty" type:"Struct"`
	RequestId         *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateDetectionTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateDetectionTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDetectionTemplateResponseBody) SetDetectionTemplate(v *UpdateDetectionTemplateResponseBodyDetectionTemplate) *UpdateDetectionTemplateResponseBody {
	s.DetectionTemplate = v
	return s
}

func (s *UpdateDetectionTemplateResponseBody) SetRequestId(v string) *UpdateDetectionTemplateResponseBody {
	s.RequestId = &v
	return s
}

type UpdateDetectionTemplateResponseBodyDetectionTemplate struct {
	CreateTime   *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	ModifyTime   *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	Period       *string `json:"Period,omitempty" xml:"Period,omitempty"`
	Platform     *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	TemplateId   *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	UserId       *int64  `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s UpdateDetectionTemplateResponseBodyDetectionTemplate) String() string {
	return tea.Prettify(s)
}

func (s UpdateDetectionTemplateResponseBodyDetectionTemplate) GoString() string {
	return s.String()
}

func (s *UpdateDetectionTemplateResponseBodyDetectionTemplate) SetCreateTime(v string) *UpdateDetectionTemplateResponseBodyDetectionTemplate {
	s.CreateTime = &v
	return s
}

func (s *UpdateDetectionTemplateResponseBodyDetectionTemplate) SetModifyTime(v string) *UpdateDetectionTemplateResponseBodyDetectionTemplate {
	s.ModifyTime = &v
	return s
}

func (s *UpdateDetectionTemplateResponseBodyDetectionTemplate) SetPeriod(v string) *UpdateDetectionTemplateResponseBodyDetectionTemplate {
	s.Period = &v
	return s
}

func (s *UpdateDetectionTemplateResponseBodyDetectionTemplate) SetPlatform(v string) *UpdateDetectionTemplateResponseBodyDetectionTemplate {
	s.Platform = &v
	return s
}

func (s *UpdateDetectionTemplateResponseBodyDetectionTemplate) SetTemplateId(v string) *UpdateDetectionTemplateResponseBodyDetectionTemplate {
	s.TemplateId = &v
	return s
}

func (s *UpdateDetectionTemplateResponseBodyDetectionTemplate) SetTemplateName(v string) *UpdateDetectionTemplateResponseBodyDetectionTemplate {
	s.TemplateName = &v
	return s
}

func (s *UpdateDetectionTemplateResponseBodyDetectionTemplate) SetUserId(v int64) *UpdateDetectionTemplateResponseBodyDetectionTemplate {
	s.UserId = &v
	return s
}

type UpdateDetectionTemplateResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateDetectionTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateDetectionTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateDetectionTemplateResponse) GoString() string {
	return s.String()
}

func (s *UpdateDetectionTemplateResponse) SetHeaders(v map[string]*string) *UpdateDetectionTemplateResponse {
	s.Headers = v
	return s
}

func (s *UpdateDetectionTemplateResponse) SetBody(v *UpdateDetectionTemplateResponseBody) *UpdateDetectionTemplateResponse {
	s.Body = v
	return s
}

type UpdateEditingProjectRequest struct {
	CoverURL             *string `json:"CoverURL,omitempty" xml:"CoverURL,omitempty"`
	Description          *string `json:"Description,omitempty" xml:"Description,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ProjectId            *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *string `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	Timeline             *string `json:"Timeline,omitempty" xml:"Timeline,omitempty"`
	Title                *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s UpdateEditingProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateEditingProjectRequest) GoString() string {
	return s.String()
}

func (s *UpdateEditingProjectRequest) SetCoverURL(v string) *UpdateEditingProjectRequest {
	s.CoverURL = &v
	return s
}

func (s *UpdateEditingProjectRequest) SetDescription(v string) *UpdateEditingProjectRequest {
	s.Description = &v
	return s
}

func (s *UpdateEditingProjectRequest) SetOwnerAccount(v string) *UpdateEditingProjectRequest {
	s.OwnerAccount = &v
	return s
}

func (s *UpdateEditingProjectRequest) SetOwnerId(v string) *UpdateEditingProjectRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateEditingProjectRequest) SetProjectId(v string) *UpdateEditingProjectRequest {
	s.ProjectId = &v
	return s
}

func (s *UpdateEditingProjectRequest) SetResourceOwnerAccount(v string) *UpdateEditingProjectRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpdateEditingProjectRequest) SetResourceOwnerId(v string) *UpdateEditingProjectRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UpdateEditingProjectRequest) SetTimeline(v string) *UpdateEditingProjectRequest {
	s.Timeline = &v
	return s
}

func (s *UpdateEditingProjectRequest) SetTitle(v string) *UpdateEditingProjectRequest {
	s.Title = &v
	return s
}

type UpdateEditingProjectResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateEditingProjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateEditingProjectResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateEditingProjectResponseBody) SetRequestId(v string) *UpdateEditingProjectResponseBody {
	s.RequestId = &v
	return s
}

type UpdateEditingProjectResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateEditingProjectResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateEditingProjectResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateEditingProjectResponse) GoString() string {
	return s.String()
}

func (s *UpdateEditingProjectResponse) SetHeaders(v map[string]*string) *UpdateEditingProjectResponse {
	s.Headers = v
	return s
}

func (s *UpdateEditingProjectResponse) SetBody(v *UpdateEditingProjectResponseBody) *UpdateEditingProjectResponse {
	s.Body = v
	return s
}

type UpdateImageInfosRequest struct {
	UpdateContent *string `json:"UpdateContent,omitempty" xml:"UpdateContent,omitempty"`
}

func (s UpdateImageInfosRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateImageInfosRequest) GoString() string {
	return s.String()
}

func (s *UpdateImageInfosRequest) SetUpdateContent(v string) *UpdateImageInfosRequest {
	s.UpdateContent = &v
	return s
}

type UpdateImageInfosResponseBody struct {
	NonExistImageIds *UpdateImageInfosResponseBodyNonExistImageIds `json:"NonExistImageIds,omitempty" xml:"NonExistImageIds,omitempty" type:"Struct"`
	RequestId        *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateImageInfosResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateImageInfosResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateImageInfosResponseBody) SetNonExistImageIds(v *UpdateImageInfosResponseBodyNonExistImageIds) *UpdateImageInfosResponseBody {
	s.NonExistImageIds = v
	return s
}

func (s *UpdateImageInfosResponseBody) SetRequestId(v string) *UpdateImageInfosResponseBody {
	s.RequestId = &v
	return s
}

type UpdateImageInfosResponseBodyNonExistImageIds struct {
	ImageId []*string `json:"ImageId,omitempty" xml:"ImageId,omitempty" type:"Repeated"`
}

func (s UpdateImageInfosResponseBodyNonExistImageIds) String() string {
	return tea.Prettify(s)
}

func (s UpdateImageInfosResponseBodyNonExistImageIds) GoString() string {
	return s.String()
}

func (s *UpdateImageInfosResponseBodyNonExistImageIds) SetImageId(v []*string) *UpdateImageInfosResponseBodyNonExistImageIds {
	s.ImageId = v
	return s
}

type UpdateImageInfosResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateImageInfosResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateImageInfosResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateImageInfosResponse) GoString() string {
	return s.String()
}

func (s *UpdateImageInfosResponse) SetHeaders(v map[string]*string) *UpdateImageInfosResponse {
	s.Headers = v
	return s
}

func (s *UpdateImageInfosResponse) SetBody(v *UpdateImageInfosResponseBody) *UpdateImageInfosResponse {
	s.Body = v
	return s
}

type UpdateTranscodeTemplateGroupRequest struct {
	Locked                   *string `json:"Locked,omitempty" xml:"Locked,omitempty"`
	Name                     *string `json:"Name,omitempty" xml:"Name,omitempty"`
	TranscodeTemplateGroupId *string `json:"TranscodeTemplateGroupId,omitempty" xml:"TranscodeTemplateGroupId,omitempty"`
	TranscodeTemplateList    *string `json:"TranscodeTemplateList,omitempty" xml:"TranscodeTemplateList,omitempty"`
}

func (s UpdateTranscodeTemplateGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateTranscodeTemplateGroupRequest) GoString() string {
	return s.String()
}

func (s *UpdateTranscodeTemplateGroupRequest) SetLocked(v string) *UpdateTranscodeTemplateGroupRequest {
	s.Locked = &v
	return s
}

func (s *UpdateTranscodeTemplateGroupRequest) SetName(v string) *UpdateTranscodeTemplateGroupRequest {
	s.Name = &v
	return s
}

func (s *UpdateTranscodeTemplateGroupRequest) SetTranscodeTemplateGroupId(v string) *UpdateTranscodeTemplateGroupRequest {
	s.TranscodeTemplateGroupId = &v
	return s
}

func (s *UpdateTranscodeTemplateGroupRequest) SetTranscodeTemplateList(v string) *UpdateTranscodeTemplateGroupRequest {
	s.TranscodeTemplateList = &v
	return s
}

type UpdateTranscodeTemplateGroupResponseBody struct {
	RequestId                *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TranscodeTemplateGroupId *string `json:"TranscodeTemplateGroupId,omitempty" xml:"TranscodeTemplateGroupId,omitempty"`
}

func (s UpdateTranscodeTemplateGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateTranscodeTemplateGroupResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateTranscodeTemplateGroupResponseBody) SetRequestId(v string) *UpdateTranscodeTemplateGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateTranscodeTemplateGroupResponseBody) SetTranscodeTemplateGroupId(v string) *UpdateTranscodeTemplateGroupResponseBody {
	s.TranscodeTemplateGroupId = &v
	return s
}

type UpdateTranscodeTemplateGroupResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateTranscodeTemplateGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateTranscodeTemplateGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateTranscodeTemplateGroupResponse) GoString() string {
	return s.String()
}

func (s *UpdateTranscodeTemplateGroupResponse) SetHeaders(v map[string]*string) *UpdateTranscodeTemplateGroupResponse {
	s.Headers = v
	return s
}

func (s *UpdateTranscodeTemplateGroupResponse) SetBody(v *UpdateTranscodeTemplateGroupResponseBody) *UpdateTranscodeTemplateGroupResponse {
	s.Body = v
	return s
}

type UpdateVideoInfoRequest struct {
	CateId      *int64  `json:"CateId,omitempty" xml:"CateId,omitempty"`
	CoverURL    *string `json:"CoverURL,omitempty" xml:"CoverURL,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Tags        *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	Title       *string `json:"Title,omitempty" xml:"Title,omitempty"`
	VideoId     *string `json:"VideoId,omitempty" xml:"VideoId,omitempty"`
}

func (s UpdateVideoInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateVideoInfoRequest) GoString() string {
	return s.String()
}

func (s *UpdateVideoInfoRequest) SetCateId(v int64) *UpdateVideoInfoRequest {
	s.CateId = &v
	return s
}

func (s *UpdateVideoInfoRequest) SetCoverURL(v string) *UpdateVideoInfoRequest {
	s.CoverURL = &v
	return s
}

func (s *UpdateVideoInfoRequest) SetDescription(v string) *UpdateVideoInfoRequest {
	s.Description = &v
	return s
}

func (s *UpdateVideoInfoRequest) SetTags(v string) *UpdateVideoInfoRequest {
	s.Tags = &v
	return s
}

func (s *UpdateVideoInfoRequest) SetTitle(v string) *UpdateVideoInfoRequest {
	s.Title = &v
	return s
}

func (s *UpdateVideoInfoRequest) SetVideoId(v string) *UpdateVideoInfoRequest {
	s.VideoId = &v
	return s
}

type UpdateVideoInfoResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateVideoInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateVideoInfoResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateVideoInfoResponseBody) SetRequestId(v string) *UpdateVideoInfoResponseBody {
	s.RequestId = &v
	return s
}

type UpdateVideoInfoResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateVideoInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateVideoInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateVideoInfoResponse) GoString() string {
	return s.String()
}

func (s *UpdateVideoInfoResponse) SetHeaders(v map[string]*string) *UpdateVideoInfoResponse {
	s.Headers = v
	return s
}

func (s *UpdateVideoInfoResponse) SetBody(v *UpdateVideoInfoResponseBody) *UpdateVideoInfoResponse {
	s.Body = v
	return s
}

type UpdateVideoInfosRequest struct {
	UpdateContent *string `json:"UpdateContent,omitempty" xml:"UpdateContent,omitempty"`
}

func (s UpdateVideoInfosRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateVideoInfosRequest) GoString() string {
	return s.String()
}

func (s *UpdateVideoInfosRequest) SetUpdateContent(v string) *UpdateVideoInfosRequest {
	s.UpdateContent = &v
	return s
}

type UpdateVideoInfosResponseBody struct {
	ForbiddenVideoIds []*string `json:"ForbiddenVideoIds,omitempty" xml:"ForbiddenVideoIds,omitempty" type:"Repeated"`
	NonExistVideoIds  []*string `json:"NonExistVideoIds,omitempty" xml:"NonExistVideoIds,omitempty" type:"Repeated"`
	RequestId         *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateVideoInfosResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateVideoInfosResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateVideoInfosResponseBody) SetForbiddenVideoIds(v []*string) *UpdateVideoInfosResponseBody {
	s.ForbiddenVideoIds = v
	return s
}

func (s *UpdateVideoInfosResponseBody) SetNonExistVideoIds(v []*string) *UpdateVideoInfosResponseBody {
	s.NonExistVideoIds = v
	return s
}

func (s *UpdateVideoInfosResponseBody) SetRequestId(v string) *UpdateVideoInfosResponseBody {
	s.RequestId = &v
	return s
}

type UpdateVideoInfosResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateVideoInfosResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateVideoInfosResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateVideoInfosResponse) GoString() string {
	return s.String()
}

func (s *UpdateVideoInfosResponse) SetHeaders(v map[string]*string) *UpdateVideoInfosResponse {
	s.Headers = v
	return s
}

func (s *UpdateVideoInfosResponse) SetBody(v *UpdateVideoInfosResponseBody) *UpdateVideoInfosResponse {
	s.Body = v
	return s
}

type UpdateVodDomainRequest struct {
	DomainName     *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	OwnerId        *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken  *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	Sources        *string `json:"Sources,omitempty" xml:"Sources,omitempty"`
	TopLevelDomain *string `json:"TopLevelDomain,omitempty" xml:"TopLevelDomain,omitempty"`
}

func (s UpdateVodDomainRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateVodDomainRequest) GoString() string {
	return s.String()
}

func (s *UpdateVodDomainRequest) SetDomainName(v string) *UpdateVodDomainRequest {
	s.DomainName = &v
	return s
}

func (s *UpdateVodDomainRequest) SetOwnerId(v int64) *UpdateVodDomainRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateVodDomainRequest) SetSecurityToken(v string) *UpdateVodDomainRequest {
	s.SecurityToken = &v
	return s
}

func (s *UpdateVodDomainRequest) SetSources(v string) *UpdateVodDomainRequest {
	s.Sources = &v
	return s
}

func (s *UpdateVodDomainRequest) SetTopLevelDomain(v string) *UpdateVodDomainRequest {
	s.TopLevelDomain = &v
	return s
}

type UpdateVodDomainResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateVodDomainResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateVodDomainResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateVodDomainResponseBody) SetRequestId(v string) *UpdateVodDomainResponseBody {
	s.RequestId = &v
	return s
}

type UpdateVodDomainResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateVodDomainResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateVodDomainResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateVodDomainResponse) GoString() string {
	return s.String()
}

func (s *UpdateVodDomainResponse) SetHeaders(v map[string]*string) *UpdateVodDomainResponse {
	s.Headers = v
	return s
}

func (s *UpdateVodDomainResponse) SetBody(v *UpdateVodDomainResponseBody) *UpdateVodDomainResponse {
	s.Body = v
	return s
}

type UpdateVodTemplateRequest struct {
	Name           *string `json:"Name,omitempty" xml:"Name,omitempty"`
	TemplateConfig *string `json:"TemplateConfig,omitempty" xml:"TemplateConfig,omitempty"`
	VodTemplateId  *string `json:"VodTemplateId,omitempty" xml:"VodTemplateId,omitempty"`
}

func (s UpdateVodTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateVodTemplateRequest) GoString() string {
	return s.String()
}

func (s *UpdateVodTemplateRequest) SetName(v string) *UpdateVodTemplateRequest {
	s.Name = &v
	return s
}

func (s *UpdateVodTemplateRequest) SetTemplateConfig(v string) *UpdateVodTemplateRequest {
	s.TemplateConfig = &v
	return s
}

func (s *UpdateVodTemplateRequest) SetVodTemplateId(v string) *UpdateVodTemplateRequest {
	s.VodTemplateId = &v
	return s
}

type UpdateVodTemplateResponseBody struct {
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	VodTemplateId *string `json:"VodTemplateId,omitempty" xml:"VodTemplateId,omitempty"`
}

func (s UpdateVodTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateVodTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateVodTemplateResponseBody) SetRequestId(v string) *UpdateVodTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateVodTemplateResponseBody) SetVodTemplateId(v string) *UpdateVodTemplateResponseBody {
	s.VodTemplateId = &v
	return s
}

type UpdateVodTemplateResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateVodTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateVodTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateVodTemplateResponse) GoString() string {
	return s.String()
}

func (s *UpdateVodTemplateResponse) SetHeaders(v map[string]*string) *UpdateVodTemplateResponse {
	s.Headers = v
	return s
}

func (s *UpdateVodTemplateResponse) SetBody(v *UpdateVodTemplateResponseBody) *UpdateVodTemplateResponse {
	s.Body = v
	return s
}

type UpdateWatermarkRequest struct {
	Name            *string `json:"Name,omitempty" xml:"Name,omitempty"`
	WatermarkConfig *string `json:"WatermarkConfig,omitempty" xml:"WatermarkConfig,omitempty"`
	WatermarkId     *string `json:"WatermarkId,omitempty" xml:"WatermarkId,omitempty"`
}

func (s UpdateWatermarkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateWatermarkRequest) GoString() string {
	return s.String()
}

func (s *UpdateWatermarkRequest) SetName(v string) *UpdateWatermarkRequest {
	s.Name = &v
	return s
}

func (s *UpdateWatermarkRequest) SetWatermarkConfig(v string) *UpdateWatermarkRequest {
	s.WatermarkConfig = &v
	return s
}

func (s *UpdateWatermarkRequest) SetWatermarkId(v string) *UpdateWatermarkRequest {
	s.WatermarkId = &v
	return s
}

type UpdateWatermarkResponseBody struct {
	RequestId     *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	WatermarkInfo *UpdateWatermarkResponseBodyWatermarkInfo `json:"WatermarkInfo,omitempty" xml:"WatermarkInfo,omitempty" type:"Struct"`
}

func (s UpdateWatermarkResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateWatermarkResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateWatermarkResponseBody) SetRequestId(v string) *UpdateWatermarkResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateWatermarkResponseBody) SetWatermarkInfo(v *UpdateWatermarkResponseBodyWatermarkInfo) *UpdateWatermarkResponseBody {
	s.WatermarkInfo = v
	return s
}

type UpdateWatermarkResponseBodyWatermarkInfo struct {
	CreationTime    *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	FileUrl         *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	IsDefault       *string `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	Name            *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Type            *string `json:"Type,omitempty" xml:"Type,omitempty"`
	WatermarkConfig *string `json:"WatermarkConfig,omitempty" xml:"WatermarkConfig,omitempty"`
	WatermarkId     *string `json:"WatermarkId,omitempty" xml:"WatermarkId,omitempty"`
}

func (s UpdateWatermarkResponseBodyWatermarkInfo) String() string {
	return tea.Prettify(s)
}

func (s UpdateWatermarkResponseBodyWatermarkInfo) GoString() string {
	return s.String()
}

func (s *UpdateWatermarkResponseBodyWatermarkInfo) SetCreationTime(v string) *UpdateWatermarkResponseBodyWatermarkInfo {
	s.CreationTime = &v
	return s
}

func (s *UpdateWatermarkResponseBodyWatermarkInfo) SetFileUrl(v string) *UpdateWatermarkResponseBodyWatermarkInfo {
	s.FileUrl = &v
	return s
}

func (s *UpdateWatermarkResponseBodyWatermarkInfo) SetIsDefault(v string) *UpdateWatermarkResponseBodyWatermarkInfo {
	s.IsDefault = &v
	return s
}

func (s *UpdateWatermarkResponseBodyWatermarkInfo) SetName(v string) *UpdateWatermarkResponseBodyWatermarkInfo {
	s.Name = &v
	return s
}

func (s *UpdateWatermarkResponseBodyWatermarkInfo) SetType(v string) *UpdateWatermarkResponseBodyWatermarkInfo {
	s.Type = &v
	return s
}

func (s *UpdateWatermarkResponseBodyWatermarkInfo) SetWatermarkConfig(v string) *UpdateWatermarkResponseBodyWatermarkInfo {
	s.WatermarkConfig = &v
	return s
}

func (s *UpdateWatermarkResponseBodyWatermarkInfo) SetWatermarkId(v string) *UpdateWatermarkResponseBodyWatermarkInfo {
	s.WatermarkId = &v
	return s
}

type UpdateWatermarkResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateWatermarkResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateWatermarkResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateWatermarkResponse) GoString() string {
	return s.String()
}

func (s *UpdateWatermarkResponse) SetHeaders(v map[string]*string) *UpdateWatermarkResponse {
	s.Headers = v
	return s
}

func (s *UpdateWatermarkResponse) SetBody(v *UpdateWatermarkResponseBody) *UpdateWatermarkResponse {
	s.Body = v
	return s
}

type UploadMediaByURLRequest struct {
	AppId           *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	StorageLocation *string `json:"StorageLocation,omitempty" xml:"StorageLocation,omitempty"`
	TemplateGroupId *string `json:"TemplateGroupId,omitempty" xml:"TemplateGroupId,omitempty"`
	UploadMetadatas *string `json:"UploadMetadatas,omitempty" xml:"UploadMetadatas,omitempty"`
	UploadURLs      *string `json:"UploadURLs,omitempty" xml:"UploadURLs,omitempty"`
	UserData        *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
	WorkflowId      *string `json:"WorkflowId,omitempty" xml:"WorkflowId,omitempty"`
}

func (s UploadMediaByURLRequest) String() string {
	return tea.Prettify(s)
}

func (s UploadMediaByURLRequest) GoString() string {
	return s.String()
}

func (s *UploadMediaByURLRequest) SetAppId(v string) *UploadMediaByURLRequest {
	s.AppId = &v
	return s
}

func (s *UploadMediaByURLRequest) SetStorageLocation(v string) *UploadMediaByURLRequest {
	s.StorageLocation = &v
	return s
}

func (s *UploadMediaByURLRequest) SetTemplateGroupId(v string) *UploadMediaByURLRequest {
	s.TemplateGroupId = &v
	return s
}

func (s *UploadMediaByURLRequest) SetUploadMetadatas(v string) *UploadMediaByURLRequest {
	s.UploadMetadatas = &v
	return s
}

func (s *UploadMediaByURLRequest) SetUploadURLs(v string) *UploadMediaByURLRequest {
	s.UploadURLs = &v
	return s
}

func (s *UploadMediaByURLRequest) SetUserData(v string) *UploadMediaByURLRequest {
	s.UserData = &v
	return s
}

func (s *UploadMediaByURLRequest) SetWorkflowId(v string) *UploadMediaByURLRequest {
	s.WorkflowId = &v
	return s
}

type UploadMediaByURLResponseBody struct {
	RequestId  *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	UploadJobs []*UploadMediaByURLResponseBodyUploadJobs `json:"UploadJobs,omitempty" xml:"UploadJobs,omitempty" type:"Repeated"`
}

func (s UploadMediaByURLResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UploadMediaByURLResponseBody) GoString() string {
	return s.String()
}

func (s *UploadMediaByURLResponseBody) SetRequestId(v string) *UploadMediaByURLResponseBody {
	s.RequestId = &v
	return s
}

func (s *UploadMediaByURLResponseBody) SetUploadJobs(v []*UploadMediaByURLResponseBodyUploadJobs) *UploadMediaByURLResponseBody {
	s.UploadJobs = v
	return s
}

type UploadMediaByURLResponseBodyUploadJobs struct {
	JobId     *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	SourceURL *string `json:"SourceURL,omitempty" xml:"SourceURL,omitempty"`
}

func (s UploadMediaByURLResponseBodyUploadJobs) String() string {
	return tea.Prettify(s)
}

func (s UploadMediaByURLResponseBodyUploadJobs) GoString() string {
	return s.String()
}

func (s *UploadMediaByURLResponseBodyUploadJobs) SetJobId(v string) *UploadMediaByURLResponseBodyUploadJobs {
	s.JobId = &v
	return s
}

func (s *UploadMediaByURLResponseBodyUploadJobs) SetSourceURL(v string) *UploadMediaByURLResponseBodyUploadJobs {
	s.SourceURL = &v
	return s
}

type UploadMediaByURLResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UploadMediaByURLResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UploadMediaByURLResponse) String() string {
	return tea.Prettify(s)
}

func (s UploadMediaByURLResponse) GoString() string {
	return s.String()
}

func (s *UploadMediaByURLResponse) SetHeaders(v map[string]*string) *UploadMediaByURLResponse {
	s.Headers = v
	return s
}

func (s *UploadMediaByURLResponse) SetBody(v *UploadMediaByURLResponseBody) *UploadMediaByURLResponse {
	s.Body = v
	return s
}

type UploadStreamByURLRequest struct {
	Definition    *string `json:"Definition,omitempty" xml:"Definition,omitempty"`
	FileExtension *string `json:"FileExtension,omitempty" xml:"FileExtension,omitempty"`
	MediaId       *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	StreamURL     *string `json:"StreamURL,omitempty" xml:"StreamURL,omitempty"`
	UserData      *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s UploadStreamByURLRequest) String() string {
	return tea.Prettify(s)
}

func (s UploadStreamByURLRequest) GoString() string {
	return s.String()
}

func (s *UploadStreamByURLRequest) SetDefinition(v string) *UploadStreamByURLRequest {
	s.Definition = &v
	return s
}

func (s *UploadStreamByURLRequest) SetFileExtension(v string) *UploadStreamByURLRequest {
	s.FileExtension = &v
	return s
}

func (s *UploadStreamByURLRequest) SetMediaId(v string) *UploadStreamByURLRequest {
	s.MediaId = &v
	return s
}

func (s *UploadStreamByURLRequest) SetStreamURL(v string) *UploadStreamByURLRequest {
	s.StreamURL = &v
	return s
}

func (s *UploadStreamByURLRequest) SetUserData(v string) *UploadStreamByURLRequest {
	s.UserData = &v
	return s
}

type UploadStreamByURLResponseBody struct {
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StreamJobId *string `json:"StreamJobId,omitempty" xml:"StreamJobId,omitempty"`
}

func (s UploadStreamByURLResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UploadStreamByURLResponseBody) GoString() string {
	return s.String()
}

func (s *UploadStreamByURLResponseBody) SetRequestId(v string) *UploadStreamByURLResponseBody {
	s.RequestId = &v
	return s
}

func (s *UploadStreamByURLResponseBody) SetStreamJobId(v string) *UploadStreamByURLResponseBody {
	s.StreamJobId = &v
	return s
}

type UploadStreamByURLResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UploadStreamByURLResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UploadStreamByURLResponse) String() string {
	return tea.Prettify(s)
}

func (s UploadStreamByURLResponse) GoString() string {
	return s.String()
}

func (s *UploadStreamByURLResponse) SetHeaders(v map[string]*string) *UploadStreamByURLResponse {
	s.Headers = v
	return s
}

func (s *UploadStreamByURLResponse) SetBody(v *UploadStreamByURLResponseBody) *UploadStreamByURLResponse {
	s.Body = v
	return s
}

type VerifyVodDomainOwnerRequest struct {
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	VerifyType *string `json:"VerifyType,omitempty" xml:"VerifyType,omitempty"`
}

func (s VerifyVodDomainOwnerRequest) String() string {
	return tea.Prettify(s)
}

func (s VerifyVodDomainOwnerRequest) GoString() string {
	return s.String()
}

func (s *VerifyVodDomainOwnerRequest) SetDomainName(v string) *VerifyVodDomainOwnerRequest {
	s.DomainName = &v
	return s
}

func (s *VerifyVodDomainOwnerRequest) SetOwnerId(v int64) *VerifyVodDomainOwnerRequest {
	s.OwnerId = &v
	return s
}

func (s *VerifyVodDomainOwnerRequest) SetVerifyType(v string) *VerifyVodDomainOwnerRequest {
	s.VerifyType = &v
	return s
}

type VerifyVodDomainOwnerResponseBody struct {
	Content   *string `json:"Content,omitempty" xml:"Content,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s VerifyVodDomainOwnerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s VerifyVodDomainOwnerResponseBody) GoString() string {
	return s.String()
}

func (s *VerifyVodDomainOwnerResponseBody) SetContent(v string) *VerifyVodDomainOwnerResponseBody {
	s.Content = &v
	return s
}

func (s *VerifyVodDomainOwnerResponseBody) SetRequestId(v string) *VerifyVodDomainOwnerResponseBody {
	s.RequestId = &v
	return s
}

type VerifyVodDomainOwnerResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *VerifyVodDomainOwnerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s VerifyVodDomainOwnerResponse) String() string {
	return tea.Prettify(s)
}

func (s VerifyVodDomainOwnerResponse) GoString() string {
	return s.String()
}

func (s *VerifyVodDomainOwnerResponse) SetHeaders(v map[string]*string) *VerifyVodDomainOwnerResponse {
	s.Headers = v
	return s
}

func (s *VerifyVodDomainOwnerResponse) SetBody(v *VerifyVodDomainOwnerResponseBody) *VerifyVodDomainOwnerResponse {
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
		"ap-northeast-2-pop":          tea.String("vod.aliyuncs.com"),
		"ap-southeast-2":              tea.String("vod.aliyuncs.com"),
		"ap-southeast-3":              tea.String("vod.aliyuncs.com"),
		"cn-beijing-finance-1":        tea.String("vod.aliyuncs.com"),
		"cn-beijing-finance-pop":      tea.String("vod.aliyuncs.com"),
		"cn-beijing-gov-1":            tea.String("vod.aliyuncs.com"),
		"cn-beijing-nu16-b01":         tea.String("vod.aliyuncs.com"),
		"cn-chengdu":                  tea.String("vod.aliyuncs.com"),
		"cn-edge-1":                   tea.String("vod.aliyuncs.com"),
		"cn-fujian":                   tea.String("vod.aliyuncs.com"),
		"cn-haidian-cm12-c01":         tea.String("vod.aliyuncs.com"),
		"cn-hangzhou":                 tea.String("vod.aliyuncs.com"),
		"cn-hangzhou-bj-b01":          tea.String("vod.aliyuncs.com"),
		"cn-hangzhou-finance":         tea.String("vod.aliyuncs.com"),
		"cn-hangzhou-internal-prod-1": tea.String("vod.aliyuncs.com"),
		"cn-hangzhou-internal-test-1": tea.String("vod.aliyuncs.com"),
		"cn-hangzhou-internal-test-2": tea.String("vod.aliyuncs.com"),
		"cn-hangzhou-internal-test-3": tea.String("vod.aliyuncs.com"),
		"cn-hangzhou-test-306":        tea.String("vod.aliyuncs.com"),
		"cn-hongkong":                 tea.String("vod.aliyuncs.com"),
		"cn-hongkong-finance-pop":     tea.String("vod.aliyuncs.com"),
		"cn-huhehaote":                tea.String("vod.aliyuncs.com"),
		"cn-huhehaote-nebula-1":       tea.String("vod.aliyuncs.com"),
		"cn-qingdao":                  tea.String("vod.aliyuncs.com"),
		"cn-qingdao-nebula":           tea.String("vod.aliyuncs.com"),
		"cn-shanghai-et15-b01":        tea.String("vod.aliyuncs.com"),
		"cn-shanghai-et2-b01":         tea.String("vod.aliyuncs.com"),
		"cn-shanghai-finance-1":       tea.String("vod.aliyuncs.com"),
		"cn-shanghai-inner":           tea.String("vod.aliyuncs.com"),
		"cn-shanghai-internal-test-1": tea.String("vod.aliyuncs.com"),
		"cn-shenzhen-finance-1":       tea.String("vod.aliyuncs.com"),
		"cn-shenzhen-inner":           tea.String("vod.aliyuncs.com"),
		"cn-shenzhen-st4-d01":         tea.String("vod.aliyuncs.com"),
		"cn-shenzhen-su18-b01":        tea.String("vod.aliyuncs.com"),
		"cn-wuhan":                    tea.String("vod.aliyuncs.com"),
		"cn-wulanchabu":               tea.String("vod.aliyuncs.com"),
		"cn-yushanfang":               tea.String("vod.aliyuncs.com"),
		"cn-zhangbei":                 tea.String("vod.aliyuncs.com"),
		"cn-zhangbei-na61-b01":        tea.String("vod.aliyuncs.com"),
		"cn-zhangjiakou":              tea.String("vod.aliyuncs.com"),
		"cn-zhangjiakou-na62-a01":     tea.String("vod.aliyuncs.com"),
		"cn-zhengzhou-nebula-1":       tea.String("vod.aliyuncs.com"),
		"eu-west-1":                   tea.String("vod.aliyuncs.com"),
		"eu-west-1-oxs":               tea.String("vod.aliyuncs.com"),
		"me-east-1":                   tea.String("vod.aliyuncs.com"),
		"rus-west-1-pop":              tea.String("vod.aliyuncs.com"),
		"us-east-1":                   tea.String("vod.aliyuncs.com"),
		"us-west-1":                   tea.String("vod.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("vod"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) AddAITemplateWithOptions(request *AddAITemplateRequest, runtime *util.RuntimeOptions) (_result *AddAITemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AddAITemplateResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AddAITemplate"), tea.String("2017-03-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddAITemplate(request *AddAITemplateRequest) (_result *AddAITemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddAITemplateResponse{}
	_body, _err := client.AddAITemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddCategoryWithOptions(request *AddCategoryRequest, runtime *util.RuntimeOptions) (_result *AddCategoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AddCategoryResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AddCategory"), tea.String("2017-03-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddCategory(request *AddCategoryRequest) (_result *AddCategoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddCategoryResponse{}
	_body, _err := client.AddCategoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddEditingProjectWithOptions(request *AddEditingProjectRequest, runtime *util.RuntimeOptions) (_result *AddEditingProjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AddEditingProjectResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AddEditingProject"), tea.String("2017-03-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddEditingProject(request *AddEditingProjectRequest) (_result *AddEditingProjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddEditingProjectResponse{}
	_body, _err := client.AddEditingProjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddTranscodeTemplateGroupWithOptions(request *AddTranscodeTemplateGroupRequest, runtime *util.RuntimeOptions) (_result *AddTranscodeTemplateGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AddTranscodeTemplateGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AddTranscodeTemplateGroup"), tea.String("2017-03-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddTranscodeTemplateGroup(request *AddTranscodeTemplateGroupRequest) (_result *AddTranscodeTemplateGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddTranscodeTemplateGroupResponse{}
	_body, _err := client.AddTranscodeTemplateGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddVodDomainWithOptions(request *AddVodDomainRequest, runtime *util.RuntimeOptions) (_result *AddVodDomainResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AddVodDomainResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AddVodDomain"), tea.String("2017-03-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddVodDomain(request *AddVodDomainRequest) (_result *AddVodDomainResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddVodDomainResponse{}
	_body, _err := client.AddVodDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddVodTemplateWithOptions(request *AddVodTemplateRequest, runtime *util.RuntimeOptions) (_result *AddVodTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AddVodTemplateResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AddVodTemplate"), tea.String("2017-03-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddVodTemplate(request *AddVodTemplateRequest) (_result *AddVodTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddVodTemplateResponse{}
	_body, _err := client.AddVodTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddWatermarkWithOptions(request *AddWatermarkRequest, runtime *util.RuntimeOptions) (_result *AddWatermarkResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AddWatermarkResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AddWatermark"), tea.String("2017-03-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddWatermark(request *AddWatermarkRequest) (_result *AddWatermarkResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddWatermarkResponse{}
	_body, _err := client.AddWatermarkWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AttachAppPolicyToIdentityWithOptions(request *AttachAppPolicyToIdentityRequest, runtime *util.RuntimeOptions) (_result *AttachAppPolicyToIdentityResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AttachAppPolicyToIdentityResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AttachAppPolicyToIdentity"), tea.String("2017-03-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AttachAppPolicyToIdentity(request *AttachAppPolicyToIdentityRequest) (_result *AttachAppPolicyToIdentityResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AttachAppPolicyToIdentityResponse{}
	_body, _err := client.AttachAppPolicyToIdentityWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BatchSetVodDomainConfigsWithOptions(request *BatchSetVodDomainConfigsRequest, runtime *util.RuntimeOptions) (_result *BatchSetVodDomainConfigsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &BatchSetVodDomainConfigsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("BatchSetVodDomainConfigs"), tea.String("2017-03-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BatchSetVodDomainConfigs(request *BatchSetVodDomainConfigsRequest) (_result *BatchSetVodDomainConfigsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BatchSetVodDomainConfigsResponse{}
	_body, _err := client.BatchSetVodDomainConfigsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BatchStartVodDomainWithOptions(request *BatchStartVodDomainRequest, runtime *util.RuntimeOptions) (_result *BatchStartVodDomainResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &BatchStartVodDomainResponse{}
	_body, _err := client.DoRPCRequest(tea.String("BatchStartVodDomain"), tea.String("2017-03-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BatchStartVodDomain(request *BatchStartVodDomainRequest) (_result *BatchStartVodDomainResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BatchStartVodDomainResponse{}
	_body, _err := client.BatchStartVodDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BatchStopVodDomainWithOptions(request *BatchStopVodDomainRequest, runtime *util.RuntimeOptions) (_result *BatchStopVodDomainResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &BatchStopVodDomainResponse{}
	_body, _err := client.DoRPCRequest(tea.String("BatchStopVodDomain"), tea.String("2017-03-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BatchStopVodDomain(request *BatchStopVodDomainRequest) (_result *BatchStopVodDomainResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BatchStopVodDomainResponse{}
	_body, _err := client.BatchStopVodDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CancelUrlUploadJobsWithOptions(request *CancelUrlUploadJobsRequest, runtime *util.RuntimeOptions) (_result *CancelUrlUploadJobsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CancelUrlUploadJobsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CancelUrlUploadJobs"), tea.String("2017-03-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CancelUrlUploadJobs(request *CancelUrlUploadJobsRequest) (_result *CancelUrlUploadJobsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CancelUrlUploadJobsResponse{}
	_body, _err := client.CancelUrlUploadJobsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateAppInfoWithOptions(request *CreateAppInfoRequest, runtime *util.RuntimeOptions) (_result *CreateAppInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateAppInfoResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateAppInfo"), tea.String("2017-03-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateAppInfo(request *CreateAppInfoRequest) (_result *CreateAppInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateAppInfoResponse{}
	_body, _err := client.CreateAppInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateAuditWithOptions(request *CreateAuditRequest, runtime *util.RuntimeOptions) (_result *CreateAuditResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateAuditResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateAudit"), tea.String("2017-03-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateAudit(request *CreateAuditRequest) (_result *CreateAuditResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateAuditResponse{}
	_body, _err := client.CreateAuditWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateDetectionTemplateWithOptions(request *CreateDetectionTemplateRequest, runtime *util.RuntimeOptions) (_result *CreateDetectionTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateDetectionTemplateResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateDetectionTemplate"), tea.String("2017-03-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateDetectionTemplate(request *CreateDetectionTemplateRequest) (_result *CreateDetectionTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateDetectionTemplateResponse{}
	_body, _err := client.CreateDetectionTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateUploadAttachedMediaWithOptions(request *CreateUploadAttachedMediaRequest, runtime *util.RuntimeOptions) (_result *CreateUploadAttachedMediaResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateUploadAttachedMediaResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateUploadAttachedMedia"), tea.String("2017-03-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateUploadAttachedMedia(request *CreateUploadAttachedMediaRequest) (_result *CreateUploadAttachedMediaResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateUploadAttachedMediaResponse{}
	_body, _err := client.CreateUploadAttachedMediaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateUploadImageWithOptions(request *CreateUploadImageRequest, runtime *util.RuntimeOptions) (_result *CreateUploadImageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateUploadImageResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateUploadImage"), tea.String("2017-03-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateUploadImage(request *CreateUploadImageRequest) (_result *CreateUploadImageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateUploadImageResponse{}
	_body, _err := client.CreateUploadImageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateUploadVideoWithOptions(request *CreateUploadVideoRequest, runtime *util.RuntimeOptions) (_result *CreateUploadVideoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateUploadVideoResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateUploadVideo"), tea.String("2017-03-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateUploadVideo(request *CreateUploadVideoRequest) (_result *CreateUploadVideoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateUploadVideoResponse{}
	_body, _err := client.CreateUploadVideoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateVodRealTimeLogDeliveryWithOptions(request *CreateVodRealTimeLogDeliveryRequest, runtime *util.RuntimeOptions) (_result *CreateVodRealTimeLogDeliveryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &CreateVodRealTimeLogDeliveryResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateVodRealTimeLogDelivery"), tea.String("2017-03-21"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateVodRealTimeLogDelivery(request *CreateVodRealTimeLogDeliveryRequest) (_result *CreateVodRealTimeLogDeliveryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateVodRealTimeLogDeliveryResponse{}
	_body, _err := client.CreateVodRealTimeLogDeliveryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteAIImageInfosWithOptions(request *DeleteAIImageInfosRequest, runtime *util.RuntimeOptions) (_result *DeleteAIImageInfosResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteAIImageInfosResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteAIImageInfos"), tea.String("2017-03-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteAIImageInfos(request *DeleteAIImageInfosRequest) (_result *DeleteAIImageInfosResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteAIImageInfosResponse{}
	_body, _err := client.DeleteAIImageInfosWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteAITemplateWithOptions(request *DeleteAITemplateRequest, runtime *util.RuntimeOptions) (_result *DeleteAITemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteAITemplateResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteAITemplate"), tea.String("2017-03-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteAITemplate(request *DeleteAITemplateRequest) (_result *DeleteAITemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteAITemplateResponse{}
	_body, _err := client.DeleteAITemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteAppInfoWithOptions(request *DeleteAppInfoRequest, runtime *util.RuntimeOptions) (_result *DeleteAppInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteAppInfoResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteAppInfo"), tea.String("2017-03-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteAppInfo(request *DeleteAppInfoRequest) (_result *DeleteAppInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteAppInfoResponse{}
	_body, _err := client.DeleteAppInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteAttachedMediaWithOptions(request *DeleteAttachedMediaRequest, runtime *util.RuntimeOptions) (_result *DeleteAttachedMediaResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteAttachedMediaResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteAttachedMedia"), tea.String("2017-03-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteAttachedMedia(request *DeleteAttachedMediaRequest) (_result *DeleteAttachedMediaResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteAttachedMediaResponse{}
	_body, _err := client.DeleteAttachedMediaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteCategoryWithOptions(request *DeleteCategoryRequest, runtime *util.RuntimeOptions) (_result *DeleteCategoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteCategoryResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteCategory"), tea.String("2017-03-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteCategory(request *DeleteCategoryRequest) (_result *DeleteCategoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteCategoryResponse{}
	_body, _err := client.DeleteCategoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteDetectionTemplateWithOptions(request *DeleteDetectionTemplateRequest, runtime *util.RuntimeOptions) (_result *DeleteDetectionTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteDetectionTemplateResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteDetectionTemplate"), tea.String("2017-03-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteDetectionTemplate(request *DeleteDetectionTemplateRequest) (_result *DeleteDetectionTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteDetectionTemplateResponse{}
	_body, _err := client.DeleteDetectionTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteDynamicImageWithOptions(request *DeleteDynamicImageRequest, runtime *util.RuntimeOptions) (_result *DeleteDynamicImageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteDynamicImageResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteDynamicImage"), tea.String("2017-03-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteDynamicImage(request *DeleteDynamicImageRequest) (_result *DeleteDynamicImageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteDynamicImageResponse{}
	_body, _err := client.DeleteDynamicImageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteEditingProjectWithOptions(request *DeleteEditingProjectRequest, runtime *util.RuntimeOptions) (_result *DeleteEditingProjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteEditingProjectResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteEditingProject"), tea.String("2017-03-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteEditingProject(request *DeleteEditingProjectRequest) (_result *DeleteEditingProjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteEditingProjectResponse{}
	_body, _err := client.DeleteEditingProjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteImageWithOptions(request *DeleteImageRequest, runtime *util.RuntimeOptions) (_result *DeleteImageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteImageResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteImage"), tea.String("2017-03-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteImage(request *DeleteImageRequest) (_result *DeleteImageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteImageResponse{}
	_body, _err := client.DeleteImageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteMessageCallbackWithOptions(request *DeleteMessageCallbackRequest, runtime *util.RuntimeOptions) (_result *DeleteMessageCallbackResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteMessageCallbackResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteMessageCallback"), tea.String("2017-03-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteMessageCallback(request *DeleteMessageCallbackRequest) (_result *DeleteMessageCallbackResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteMessageCallbackResponse{}
	_body, _err := client.DeleteMessageCallbackWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteMezzaninesWithOptions(request *DeleteMezzaninesRequest, runtime *util.RuntimeOptions) (_result *DeleteMezzaninesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteMezzaninesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteMezzanines"), tea.String("2017-03-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteMezzanines(request *DeleteMezzaninesRequest) (_result *DeleteMezzaninesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteMezzaninesResponse{}
	_body, _err := client.DeleteMezzaninesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteMultipartUploadWithOptions(request *DeleteMultipartUploadRequest, runtime *util.RuntimeOptions) (_result *DeleteMultipartUploadResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteMultipartUploadResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteMultipartUpload"), tea.String("2017-03-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteMultipartUpload(request *DeleteMultipartUploadRequest) (_result *DeleteMultipartUploadResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteMultipartUploadResponse{}
	_body, _err := client.DeleteMultipartUploadWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteStreamWithOptions(request *DeleteStreamRequest, runtime *util.RuntimeOptions) (_result *DeleteStreamResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteStreamResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteStream"), tea.String("2017-03-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteStream(request *DeleteStreamRequest) (_result *DeleteStreamResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteStreamResponse{}
	_body, _err := client.DeleteStreamWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteTranscodeTemplateGroupWithOptions(request *DeleteTranscodeTemplateGroupRequest, runtime *util.RuntimeOptions) (_result *DeleteTranscodeTemplateGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteTranscodeTemplateGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteTranscodeTemplateGroup"), tea.String("2017-03-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteTranscodeTemplateGroup(request *DeleteTranscodeTemplateGroupRequest) (_result *DeleteTranscodeTemplateGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteTranscodeTemplateGroupResponse{}
	_body, _err := client.DeleteTranscodeTemplateGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteVideoWithOptions(request *DeleteVideoRequest, runtime *util.RuntimeOptions) (_result *DeleteVideoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteVideoResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteVideo"), tea.String("2017-03-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteVideo(request *DeleteVideoRequest) (_result *DeleteVideoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteVideoResponse{}
	_body, _err := client.DeleteVideoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteVodDomainWithOptions(request *DeleteVodDomainRequest, runtime *util.RuntimeOptions) (_result *DeleteVodDomainResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteVodDomainResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteVodDomain"), tea.String("2017-03-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteVodDomain(request *DeleteVodDomainRequest) (_result *DeleteVodDomainResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteVodDomainResponse{}
	_body, _err := client.DeleteVodDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteVodRealtimeLogDeliveryWithOptions(request *DeleteVodRealtimeLogDeliveryRequest, runtime *util.RuntimeOptions) (_result *DeleteVodRealtimeLogDeliveryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &DeleteVodRealtimeLogDeliveryResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteVodRealtimeLogDelivery"), tea.String("2017-03-21"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteVodRealtimeLogDelivery(request *DeleteVodRealtimeLogDeliveryRequest) (_result *DeleteVodRealtimeLogDeliveryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteVodRealtimeLogDeliveryResponse{}
	_body, _err := client.DeleteVodRealtimeLogDeliveryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteVodSpecificConfigWithOptions(request *DeleteVodSpecificConfigRequest, runtime *util.RuntimeOptions) (_result *DeleteVodSpecificConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteVodSpecificConfigResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteVodSpecificConfig"), tea.String("2017-03-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteVodSpecificConfig(request *DeleteVodSpecificConfigRequest) (_result *DeleteVodSpecificConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteVodSpecificConfigResponse{}
	_body, _err := client.DeleteVodSpecificConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteVodTemplateWithOptions(request *DeleteVodTemplateRequest, runtime *util.RuntimeOptions) (_result *DeleteVodTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteVodTemplateResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteVodTemplate"), tea.String("2017-03-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteVodTemplate(request *DeleteVodTemplateRequest) (_result *DeleteVodTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteVodTemplateResponse{}
	_body, _err := client.DeleteVodTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteWatermarkWithOptions(request *DeleteWatermarkRequest, runtime *util.RuntimeOptions) (_result *DeleteWatermarkResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteWatermarkResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteWatermark"), tea.String("2017-03-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteWatermark(request *DeleteWatermarkRequest) (_result *DeleteWatermarkResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteWatermarkResponse{}
	_body, _err := client.DeleteWatermarkWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribePlayTopVideosWithOptions(request *DescribePlayTopVideosRequest, runtime *util.RuntimeOptions) (_result *DescribePlayTopVideosResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribePlayTopVideosResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribePlayTopVideos"), tea.String("2017-03-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribePlayTopVideos(request *DescribePlayTopVideosRequest) (_result *DescribePlayTopVideosResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribePlayTopVideosResponse{}
	_body, _err := client.DescribePlayTopVideosWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribePlayUserAvgWithOptions(request *DescribePlayUserAvgRequest, runtime *util.RuntimeOptions) (_result *DescribePlayUserAvgResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribePlayUserAvgResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribePlayUserAvg"), tea.String("2017-03-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribePlayUserAvg(request *DescribePlayUserAvgRequest) (_result *DescribePlayUserAvgResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribePlayUserAvgResponse{}
	_body, _err := client.DescribePlayUserAvgWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribePlayUserTotalWithOptions(request *DescribePlayUserTotalRequest, runtime *util.RuntimeOptions) (_result *DescribePlayUserTotalResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribePlayUserTotalResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribePlayUserTotal"), tea.String("2017-03-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribePlayUserTotal(request *DescribePlayUserTotalRequest) (_result *DescribePlayUserTotalResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribePlayUserTotalResponse{}
	_body, _err := client.DescribePlayUserTotalWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribePlayVideoStatisWithOptions(request *DescribePlayVideoStatisRequest, runtime *util.RuntimeOptions) (_result *DescribePlayVideoStatisResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribePlayVideoStatisResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribePlayVideoStatis"), tea.String("2017-03-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribePlayVideoStatis(request *DescribePlayVideoStatisRequest) (_result *DescribePlayVideoStatisResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribePlayVideoStatisResponse{}
	_body, _err := client.DescribePlayVideoStatisWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeVodAIDataWithOptions(request *DescribeVodAIDataRequest, runtime *util.RuntimeOptions) (_result *DescribeVodAIDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeVodAIDataResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeVodAIData"), tea.String("2017-03-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeVodAIData(request *DescribeVodAIDataRequest) (_result *DescribeVodAIDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeVodAIDataResponse{}
	_body, _err := client.DescribeVodAIDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeVodCertificateListWithOptions(request *DescribeVodCertificateListRequest, runtime *util.RuntimeOptions) (_result *DescribeVodCertificateListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeVodCertificateListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeVodCertificateList"), tea.String("2017-03-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeVodCertificateList(request *DescribeVodCertificateListRequest) (_result *DescribeVodCertificateListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeVodCertificateListResponse{}
	_body, _err := client.DescribeVodCertificateListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeVodDomainBpsDataWithOptions(request *DescribeVodDomainBpsDataRequest, runtime *util.RuntimeOptions) (_result *DescribeVodDomainBpsDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeVodDomainBpsDataResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeVodDomainBpsData"), tea.String("2017-03-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeVodDomainBpsData(request *DescribeVodDomainBpsDataRequest) (_result *DescribeVodDomainBpsDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeVodDomainBpsDataResponse{}
	_body, _err := client.DescribeVodDomainBpsDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeVodDomainCertificateInfoWithOptions(request *DescribeVodDomainCertificateInfoRequest, runtime *util.RuntimeOptions) (_result *DescribeVodDomainCertificateInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeVodDomainCertificateInfoResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeVodDomainCertificateInfo"), tea.String("2017-03-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeVodDomainCertificateInfo(request *DescribeVodDomainCertificateInfoRequest) (_result *DescribeVodDomainCertificateInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeVodDomainCertificateInfoResponse{}
	_body, _err := client.DescribeVodDomainCertificateInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeVodDomainConfigsWithOptions(request *DescribeVodDomainConfigsRequest, runtime *util.RuntimeOptions) (_result *DescribeVodDomainConfigsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeVodDomainConfigsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeVodDomainConfigs"), tea.String("2017-03-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeVodDomainConfigs(request *DescribeVodDomainConfigsRequest) (_result *DescribeVodDomainConfigsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeVodDomainConfigsResponse{}
	_body, _err := client.DescribeVodDomainConfigsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeVodDomainDetailWithOptions(request *DescribeVodDomainDetailRequest, runtime *util.RuntimeOptions) (_result *DescribeVodDomainDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeVodDomainDetailResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeVodDomainDetail"), tea.String("2017-03-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeVodDomainDetail(request *DescribeVodDomainDetailRequest) (_result *DescribeVodDomainDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeVodDomainDetailResponse{}
	_body, _err := client.DescribeVodDomainDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeVodDomainLogWithOptions(request *DescribeVodDomainLogRequest, runtime *util.RuntimeOptions) (_result *DescribeVodDomainLogResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeVodDomainLogResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeVodDomainLog"), tea.String("2017-03-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeVodDomainLog(request *DescribeVodDomainLogRequest) (_result *DescribeVodDomainLogResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeVodDomainLogResponse{}
	_body, _err := client.DescribeVodDomainLogWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeVodDomainRealtimeLogDeliveryWithOptions(request *DescribeVodDomainRealtimeLogDeliveryRequest, runtime *util.RuntimeOptions) (_result *DescribeVodDomainRealtimeLogDeliveryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &DescribeVodDomainRealtimeLogDeliveryResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeVodDomainRealtimeLogDelivery"), tea.String("2017-03-21"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeVodDomainRealtimeLogDelivery(request *DescribeVodDomainRealtimeLogDeliveryRequest) (_result *DescribeVodDomainRealtimeLogDeliveryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeVodDomainRealtimeLogDeliveryResponse{}
	_body, _err := client.DescribeVodDomainRealtimeLogDeliveryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeVodDomainTrafficDataWithOptions(request *DescribeVodDomainTrafficDataRequest, runtime *util.RuntimeOptions) (_result *DescribeVodDomainTrafficDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeVodDomainTrafficDataResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeVodDomainTrafficData"), tea.String("2017-03-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeVodDomainTrafficData(request *DescribeVodDomainTrafficDataRequest) (_result *DescribeVodDomainTrafficDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeVodDomainTrafficDataResponse{}
	_body, _err := client.DescribeVodDomainTrafficDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeVodDomainUsageDataWithOptions(request *DescribeVodDomainUsageDataRequest, runtime *util.RuntimeOptions) (_result *DescribeVodDomainUsageDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeVodDomainUsageDataResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeVodDomainUsageData"), tea.String("2017-03-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeVodDomainUsageData(request *DescribeVodDomainUsageDataRequest) (_result *DescribeVodDomainUsageDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeVodDomainUsageDataResponse{}
	_body, _err := client.DescribeVodDomainUsageDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeVodRefreshQuotaWithOptions(request *DescribeVodRefreshQuotaRequest, runtime *util.RuntimeOptions) (_result *DescribeVodRefreshQuotaResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeVodRefreshQuotaResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeVodRefreshQuota"), tea.String("2017-03-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeVodRefreshQuota(request *DescribeVodRefreshQuotaRequest) (_result *DescribeVodRefreshQuotaResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeVodRefreshQuotaResponse{}
	_body, _err := client.DescribeVodRefreshQuotaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeVodRefreshTasksWithOptions(request *DescribeVodRefreshTasksRequest, runtime *util.RuntimeOptions) (_result *DescribeVodRefreshTasksResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeVodRefreshTasksResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeVodRefreshTasks"), tea.String("2017-03-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeVodRefreshTasks(request *DescribeVodRefreshTasksRequest) (_result *DescribeVodRefreshTasksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeVodRefreshTasksResponse{}
	_body, _err := client.DescribeVodRefreshTasksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeVodStorageDataWithOptions(request *DescribeVodStorageDataRequest, runtime *util.RuntimeOptions) (_result *DescribeVodStorageDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeVodStorageDataResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeVodStorageData"), tea.String("2017-03-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeVodStorageData(request *DescribeVodStorageDataRequest) (_result *DescribeVodStorageDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeVodStorageDataResponse{}
	_body, _err := client.DescribeVodStorageDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeVodTagResourcesWithOptions(request *DescribeVodTagResourcesRequest, runtime *util.RuntimeOptions) (_result *DescribeVodTagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeVodTagResourcesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeVodTagResources"), tea.String("2017-03-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeVodTagResources(request *DescribeVodTagResourcesRequest) (_result *DescribeVodTagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeVodTagResourcesResponse{}
	_body, _err := client.DescribeVodTagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeVodTranscodeDataWithOptions(request *DescribeVodTranscodeDataRequest, runtime *util.RuntimeOptions) (_result *DescribeVodTranscodeDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeVodTranscodeDataResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeVodTranscodeData"), tea.String("2017-03-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeVodTranscodeData(request *DescribeVodTranscodeDataRequest) (_result *DescribeVodTranscodeDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeVodTranscodeDataResponse{}
	_body, _err := client.DescribeVodTranscodeDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeVodUserDomainsWithOptions(request *DescribeVodUserDomainsRequest, runtime *util.RuntimeOptions) (_result *DescribeVodUserDomainsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeVodUserDomainsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeVodUserDomains"), tea.String("2017-03-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeVodUserDomains(request *DescribeVodUserDomainsRequest) (_result *DescribeVodUserDomainsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeVodUserDomainsResponse{}
	_body, _err := client.DescribeVodUserDomainsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeVodUserTagsWithOptions(request *DescribeVodUserTagsRequest, runtime *util.RuntimeOptions) (_result *DescribeVodUserTagsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeVodUserTagsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeVodUserTags"), tea.String("2017-03-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeVodUserTags(request *DescribeVodUserTagsRequest) (_result *DescribeVodUserTagsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeVodUserTagsResponse{}
	_body, _err := client.DescribeVodUserTagsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeVodVerifyContentWithOptions(request *DescribeVodVerifyContentRequest, runtime *util.RuntimeOptions) (_result *DescribeVodVerifyContentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeVodVerifyContentResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeVodVerifyContent"), tea.String("2017-03-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeVodVerifyContent(request *DescribeVodVerifyContentRequest) (_result *DescribeVodVerifyContentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeVodVerifyContentResponse{}
	_body, _err := client.DescribeVodVerifyContentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DetachAppPolicyFromIdentityWithOptions(request *DetachAppPolicyFromIdentityRequest, runtime *util.RuntimeOptions) (_result *DetachAppPolicyFromIdentityResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DetachAppPolicyFromIdentityResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DetachAppPolicyFromIdentity"), tea.String("2017-03-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DetachAppPolicyFromIdentity(request *DetachAppPolicyFromIdentityRequest) (_result *DetachAppPolicyFromIdentityResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DetachAppPolicyFromIdentityResponse{}
	_body, _err := client.DetachAppPolicyFromIdentityWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DisableVodRealtimeLogDeliveryWithOptions(request *DisableVodRealtimeLogDeliveryRequest, runtime *util.RuntimeOptions) (_result *DisableVodRealtimeLogDeliveryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &DisableVodRealtimeLogDeliveryResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DisableVodRealtimeLogDelivery"), tea.String("2017-03-21"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DisableVodRealtimeLogDelivery(request *DisableVodRealtimeLogDeliveryRequest) (_result *DisableVodRealtimeLogDeliveryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DisableVodRealtimeLogDeliveryResponse{}
	_body, _err := client.DisableVodRealtimeLogDeliveryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) EnableVodRealtimeLogDeliveryWithOptions(request *EnableVodRealtimeLogDeliveryRequest, runtime *util.RuntimeOptions) (_result *EnableVodRealtimeLogDeliveryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &EnableVodRealtimeLogDeliveryResponse{}
	_body, _err := client.DoRPCRequest(tea.String("EnableVodRealtimeLogDelivery"), tea.String("2017-03-21"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EnableVodRealtimeLogDelivery(request *EnableVodRealtimeLogDeliveryRequest) (_result *EnableVodRealtimeLogDeliveryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EnableVodRealtimeLogDeliveryResponse{}
	_body, _err := client.EnableVodRealtimeLogDeliveryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAICaptionExtractionJobsWithOptions(request *GetAICaptionExtractionJobsRequest, runtime *util.RuntimeOptions) (_result *GetAICaptionExtractionJobsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetAICaptionExtractionJobsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetAICaptionExtractionJobs"), tea.String("2017-03-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAICaptionExtractionJobs(request *GetAICaptionExtractionJobsRequest) (_result *GetAICaptionExtractionJobsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAICaptionExtractionJobsResponse{}
	_body, _err := client.GetAICaptionExtractionJobsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAIImageJobsWithOptions(request *GetAIImageJobsRequest, runtime *util.RuntimeOptions) (_result *GetAIImageJobsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetAIImageJobsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetAIImageJobs"), tea.String("2017-03-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAIImageJobs(request *GetAIImageJobsRequest) (_result *GetAIImageJobsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAIImageJobsResponse{}
	_body, _err := client.GetAIImageJobsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAIMediaAuditJobWithOptions(request *GetAIMediaAuditJobRequest, runtime *util.RuntimeOptions) (_result *GetAIMediaAuditJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetAIMediaAuditJobResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetAIMediaAuditJob"), tea.String("2017-03-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAIMediaAuditJob(request *GetAIMediaAuditJobRequest) (_result *GetAIMediaAuditJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAIMediaAuditJobResponse{}
	_body, _err := client.GetAIMediaAuditJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAITemplateWithOptions(request *GetAITemplateRequest, runtime *util.RuntimeOptions) (_result *GetAITemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetAITemplateResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetAITemplate"), tea.String("2017-03-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAITemplate(request *GetAITemplateRequest) (_result *GetAITemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAITemplateResponse{}
	_body, _err := client.GetAITemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAIVideoTagResultWithOptions(request *GetAIVideoTagResultRequest, runtime *util.RuntimeOptions) (_result *GetAIVideoTagResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetAIVideoTagResultResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetAIVideoTagResult"), tea.String("2017-03-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAIVideoTagResult(request *GetAIVideoTagResultRequest) (_result *GetAIVideoTagResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAIVideoTagResultResponse{}
	_body, _err := client.GetAIVideoTagResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAppInfosWithOptions(request *GetAppInfosRequest, runtime *util.RuntimeOptions) (_result *GetAppInfosResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetAppInfosResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetAppInfos"), tea.String("2017-03-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAppInfos(request *GetAppInfosRequest) (_result *GetAppInfosResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAppInfosResponse{}
	_body, _err := client.GetAppInfosWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAttachedMediaInfoWithOptions(request *GetAttachedMediaInfoRequest, runtime *util.RuntimeOptions) (_result *GetAttachedMediaInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetAttachedMediaInfoResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetAttachedMediaInfo"), tea.String("2017-03-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAttachedMediaInfo(request *GetAttachedMediaInfoRequest) (_result *GetAttachedMediaInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAttachedMediaInfoResponse{}
	_body, _err := client.GetAttachedMediaInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAuditHistoryWithOptions(request *GetAuditHistoryRequest, runtime *util.RuntimeOptions) (_result *GetAuditHistoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetAuditHistoryResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetAuditHistory"), tea.String("2017-03-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAuditHistory(request *GetAuditHistoryRequest) (_result *GetAuditHistoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAuditHistoryResponse{}
	_body, _err := client.GetAuditHistoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetCategoriesWithOptions(request *GetCategoriesRequest, runtime *util.RuntimeOptions) (_result *GetCategoriesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetCategoriesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetCategories"), tea.String("2017-03-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetCategories(request *GetCategoriesRequest) (_result *GetCategoriesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetCategoriesResponse{}
	_body, _err := client.GetCategoriesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetDefaultAITemplateWithOptions(request *GetDefaultAITemplateRequest, runtime *util.RuntimeOptions) (_result *GetDefaultAITemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetDefaultAITemplateResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetDefaultAITemplate"), tea.String("2017-03-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetDefaultAITemplate(request *GetDefaultAITemplateRequest) (_result *GetDefaultAITemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetDefaultAITemplateResponse{}
	_body, _err := client.GetDefaultAITemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetDetectionJobWithOptions(request *GetDetectionJobRequest, runtime *util.RuntimeOptions) (_result *GetDetectionJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &GetDetectionJobResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetDetectionJob"), tea.String("2017-03-21"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetDetectionJob(request *GetDetectionJobRequest) (_result *GetDetectionJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetDetectionJobResponse{}
	_body, _err := client.GetDetectionJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetDetectionResultWithOptions(request *GetDetectionResultRequest, runtime *util.RuntimeOptions) (_result *GetDetectionResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &GetDetectionResultResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetDetectionResult"), tea.String("2017-03-21"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetDetectionResult(request *GetDetectionResultRequest) (_result *GetDetectionResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetDetectionResultResponse{}
	_body, _err := client.GetDetectionResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetDetectionTemplateWithOptions(request *GetDetectionTemplateRequest, runtime *util.RuntimeOptions) (_result *GetDetectionTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &GetDetectionTemplateResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetDetectionTemplate"), tea.String("2017-03-21"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetDetectionTemplate(request *GetDetectionTemplateRequest) (_result *GetDetectionTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetDetectionTemplateResponse{}
	_body, _err := client.GetDetectionTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetEditingProjectWithOptions(request *GetEditingProjectRequest, runtime *util.RuntimeOptions) (_result *GetEditingProjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetEditingProjectResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetEditingProject"), tea.String("2017-03-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetEditingProject(request *GetEditingProjectRequest) (_result *GetEditingProjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetEditingProjectResponse{}
	_body, _err := client.GetEditingProjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetEditingProjectMaterialsWithOptions(request *GetEditingProjectMaterialsRequest, runtime *util.RuntimeOptions) (_result *GetEditingProjectMaterialsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetEditingProjectMaterialsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetEditingProjectMaterials"), tea.String("2017-03-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetEditingProjectMaterials(request *GetEditingProjectMaterialsRequest) (_result *GetEditingProjectMaterialsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetEditingProjectMaterialsResponse{}
	_body, _err := client.GetEditingProjectMaterialsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetImageInfoWithOptions(request *GetImageInfoRequest, runtime *util.RuntimeOptions) (_result *GetImageInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetImageInfoResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetImageInfo"), tea.String("2017-03-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetImageInfo(request *GetImageInfoRequest) (_result *GetImageInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetImageInfoResponse{}
	_body, _err := client.GetImageInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetMediaAuditAudioResultDetailWithOptions(request *GetMediaAuditAudioResultDetailRequest, runtime *util.RuntimeOptions) (_result *GetMediaAuditAudioResultDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetMediaAuditAudioResultDetailResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetMediaAuditAudioResultDetail"), tea.String("2017-03-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetMediaAuditAudioResultDetail(request *GetMediaAuditAudioResultDetailRequest) (_result *GetMediaAuditAudioResultDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetMediaAuditAudioResultDetailResponse{}
	_body, _err := client.GetMediaAuditAudioResultDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetMediaAuditResultWithOptions(request *GetMediaAuditResultRequest, runtime *util.RuntimeOptions) (_result *GetMediaAuditResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetMediaAuditResultResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetMediaAuditResult"), tea.String("2017-03-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetMediaAuditResult(request *GetMediaAuditResultRequest) (_result *GetMediaAuditResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetMediaAuditResultResponse{}
	_body, _err := client.GetMediaAuditResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetMediaAuditResultDetailWithOptions(request *GetMediaAuditResultDetailRequest, runtime *util.RuntimeOptions) (_result *GetMediaAuditResultDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetMediaAuditResultDetailResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetMediaAuditResultDetail"), tea.String("2017-03-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetMediaAuditResultDetail(request *GetMediaAuditResultDetailRequest) (_result *GetMediaAuditResultDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetMediaAuditResultDetailResponse{}
	_body, _err := client.GetMediaAuditResultDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetMediaAuditResultTimelineWithOptions(request *GetMediaAuditResultTimelineRequest, runtime *util.RuntimeOptions) (_result *GetMediaAuditResultTimelineResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetMediaAuditResultTimelineResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetMediaAuditResultTimeline"), tea.String("2017-03-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetMediaAuditResultTimeline(request *GetMediaAuditResultTimelineRequest) (_result *GetMediaAuditResultTimelineResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetMediaAuditResultTimelineResponse{}
	_body, _err := client.GetMediaAuditResultTimelineWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetMediaDNAResultWithOptions(request *GetMediaDNAResultRequest, runtime *util.RuntimeOptions) (_result *GetMediaDNAResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetMediaDNAResultResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetMediaDNAResult"), tea.String("2017-03-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetMediaDNAResult(request *GetMediaDNAResultRequest) (_result *GetMediaDNAResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetMediaDNAResultResponse{}
	_body, _err := client.GetMediaDNAResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetMessageCallbackWithOptions(request *GetMessageCallbackRequest, runtime *util.RuntimeOptions) (_result *GetMessageCallbackResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetMessageCallbackResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetMessageCallback"), tea.String("2017-03-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetMessageCallback(request *GetMessageCallbackRequest) (_result *GetMessageCallbackResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetMessageCallbackResponse{}
	_body, _err := client.GetMessageCallbackWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetMezzanineInfoWithOptions(request *GetMezzanineInfoRequest, runtime *util.RuntimeOptions) (_result *GetMezzanineInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetMezzanineInfoResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetMezzanineInfo"), tea.String("2017-03-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetMezzanineInfo(request *GetMezzanineInfoRequest) (_result *GetMezzanineInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetMezzanineInfoResponse{}
	_body, _err := client.GetMezzanineInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetPlayInfoWithOptions(request *GetPlayInfoRequest, runtime *util.RuntimeOptions) (_result *GetPlayInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetPlayInfoResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetPlayInfo"), tea.String("2017-03-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetPlayInfo(request *GetPlayInfoRequest) (_result *GetPlayInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetPlayInfoResponse{}
	_body, _err := client.GetPlayInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetTranscodeSummaryWithOptions(request *GetTranscodeSummaryRequest, runtime *util.RuntimeOptions) (_result *GetTranscodeSummaryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetTranscodeSummaryResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetTranscodeSummary"), tea.String("2017-03-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetTranscodeSummary(request *GetTranscodeSummaryRequest) (_result *GetTranscodeSummaryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetTranscodeSummaryResponse{}
	_body, _err := client.GetTranscodeSummaryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetTranscodeTaskWithOptions(request *GetTranscodeTaskRequest, runtime *util.RuntimeOptions) (_result *GetTranscodeTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetTranscodeTaskResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetTranscodeTask"), tea.String("2017-03-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetTranscodeTask(request *GetTranscodeTaskRequest) (_result *GetTranscodeTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetTranscodeTaskResponse{}
	_body, _err := client.GetTranscodeTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetTranscodeTemplateGroupWithOptions(request *GetTranscodeTemplateGroupRequest, runtime *util.RuntimeOptions) (_result *GetTranscodeTemplateGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetTranscodeTemplateGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetTranscodeTemplateGroup"), tea.String("2017-03-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetTranscodeTemplateGroup(request *GetTranscodeTemplateGroupRequest) (_result *GetTranscodeTemplateGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetTranscodeTemplateGroupResponse{}
	_body, _err := client.GetTranscodeTemplateGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetURLUploadInfosWithOptions(request *GetURLUploadInfosRequest, runtime *util.RuntimeOptions) (_result *GetURLUploadInfosResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetURLUploadInfosResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetURLUploadInfos"), tea.String("2017-03-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetURLUploadInfos(request *GetURLUploadInfosRequest) (_result *GetURLUploadInfosResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetURLUploadInfosResponse{}
	_body, _err := client.GetURLUploadInfosWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetUploadDetailsWithOptions(request *GetUploadDetailsRequest, runtime *util.RuntimeOptions) (_result *GetUploadDetailsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetUploadDetailsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetUploadDetails"), tea.String("2017-03-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetUploadDetails(request *GetUploadDetailsRequest) (_result *GetUploadDetailsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetUploadDetailsResponse{}
	_body, _err := client.GetUploadDetailsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetVideoInfoWithOptions(request *GetVideoInfoRequest, runtime *util.RuntimeOptions) (_result *GetVideoInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetVideoInfoResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetVideoInfo"), tea.String("2017-03-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetVideoInfo(request *GetVideoInfoRequest) (_result *GetVideoInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetVideoInfoResponse{}
	_body, _err := client.GetVideoInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetVideoInfosWithOptions(request *GetVideoInfosRequest, runtime *util.RuntimeOptions) (_result *GetVideoInfosResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetVideoInfosResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetVideoInfos"), tea.String("2017-03-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetVideoInfos(request *GetVideoInfosRequest) (_result *GetVideoInfosResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetVideoInfosResponse{}
	_body, _err := client.GetVideoInfosWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetVideoListWithOptions(request *GetVideoListRequest, runtime *util.RuntimeOptions) (_result *GetVideoListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetVideoListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetVideoList"), tea.String("2017-03-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetVideoList(request *GetVideoListRequest) (_result *GetVideoListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetVideoListResponse{}
	_body, _err := client.GetVideoListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetVideoPlayAuthWithOptions(request *GetVideoPlayAuthRequest, runtime *util.RuntimeOptions) (_result *GetVideoPlayAuthResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetVideoPlayAuthResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetVideoPlayAuth"), tea.String("2017-03-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetVideoPlayAuth(request *GetVideoPlayAuthRequest) (_result *GetVideoPlayAuthResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetVideoPlayAuthResponse{}
	_body, _err := client.GetVideoPlayAuthWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetVodTemplateWithOptions(request *GetVodTemplateRequest, runtime *util.RuntimeOptions) (_result *GetVodTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetVodTemplateResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetVodTemplate"), tea.String("2017-03-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetVodTemplate(request *GetVodTemplateRequest) (_result *GetVodTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetVodTemplateResponse{}
	_body, _err := client.GetVodTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetWatermarkWithOptions(request *GetWatermarkRequest, runtime *util.RuntimeOptions) (_result *GetWatermarkResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetWatermarkResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetWatermark"), tea.String("2017-03-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetWatermark(request *GetWatermarkRequest) (_result *GetWatermarkResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetWatermarkResponse{}
	_body, _err := client.GetWatermarkWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListAIImageInfoWithOptions(request *ListAIImageInfoRequest, runtime *util.RuntimeOptions) (_result *ListAIImageInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListAIImageInfoResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListAIImageInfo"), tea.String("2017-03-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListAIImageInfo(request *ListAIImageInfoRequest) (_result *ListAIImageInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListAIImageInfoResponse{}
	_body, _err := client.ListAIImageInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListAIJobWithOptions(request *ListAIJobRequest, runtime *util.RuntimeOptions) (_result *ListAIJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListAIJobResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListAIJob"), tea.String("2017-03-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListAIJob(request *ListAIJobRequest) (_result *ListAIJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListAIJobResponse{}
	_body, _err := client.ListAIJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListAITemplateWithOptions(request *ListAITemplateRequest, runtime *util.RuntimeOptions) (_result *ListAITemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListAITemplateResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListAITemplate"), tea.String("2017-03-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListAITemplate(request *ListAITemplateRequest) (_result *ListAITemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListAITemplateResponse{}
	_body, _err := client.ListAITemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListAppInfoWithOptions(request *ListAppInfoRequest, runtime *util.RuntimeOptions) (_result *ListAppInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListAppInfoResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListAppInfo"), tea.String("2017-03-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListAppInfo(request *ListAppInfoRequest) (_result *ListAppInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListAppInfoResponse{}
	_body, _err := client.ListAppInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListAppPoliciesForIdentityWithOptions(request *ListAppPoliciesForIdentityRequest, runtime *util.RuntimeOptions) (_result *ListAppPoliciesForIdentityResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListAppPoliciesForIdentityResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListAppPoliciesForIdentity"), tea.String("2017-03-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListAppPoliciesForIdentity(request *ListAppPoliciesForIdentityRequest) (_result *ListAppPoliciesForIdentityResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListAppPoliciesForIdentityResponse{}
	_body, _err := client.ListAppPoliciesForIdentityWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListAuditSecurityIpWithOptions(request *ListAuditSecurityIpRequest, runtime *util.RuntimeOptions) (_result *ListAuditSecurityIpResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListAuditSecurityIpResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListAuditSecurityIp"), tea.String("2017-03-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListAuditSecurityIp(request *ListAuditSecurityIpRequest) (_result *ListAuditSecurityIpResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListAuditSecurityIpResponse{}
	_body, _err := client.ListAuditSecurityIpWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListDetectionJobWithOptions(request *ListDetectionJobRequest, runtime *util.RuntimeOptions) (_result *ListDetectionJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &ListDetectionJobResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListDetectionJob"), tea.String("2017-03-21"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListDetectionJob(request *ListDetectionJobRequest) (_result *ListDetectionJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListDetectionJobResponse{}
	_body, _err := client.ListDetectionJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListDetectionTemplateWithOptions(request *ListDetectionTemplateRequest, runtime *util.RuntimeOptions) (_result *ListDetectionTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &ListDetectionTemplateResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListDetectionTemplate"), tea.String("2017-03-21"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListDetectionTemplate(request *ListDetectionTemplateRequest) (_result *ListDetectionTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListDetectionTemplateResponse{}
	_body, _err := client.ListDetectionTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListDynamicImageWithOptions(request *ListDynamicImageRequest, runtime *util.RuntimeOptions) (_result *ListDynamicImageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListDynamicImageResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListDynamicImage"), tea.String("2017-03-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListDynamicImage(request *ListDynamicImageRequest) (_result *ListDynamicImageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListDynamicImageResponse{}
	_body, _err := client.ListDynamicImageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListLetterSendJobWithOptions(request *ListLetterSendJobRequest, runtime *util.RuntimeOptions) (_result *ListLetterSendJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &ListLetterSendJobResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListLetterSendJob"), tea.String("2017-03-21"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListLetterSendJob(request *ListLetterSendJobRequest) (_result *ListLetterSendJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListLetterSendJobResponse{}
	_body, _err := client.ListLetterSendJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListLiveRecordVideoWithOptions(request *ListLiveRecordVideoRequest, runtime *util.RuntimeOptions) (_result *ListLiveRecordVideoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListLiveRecordVideoResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListLiveRecordVideo"), tea.String("2017-03-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListLiveRecordVideo(request *ListLiveRecordVideoRequest) (_result *ListLiveRecordVideoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListLiveRecordVideoResponse{}
	_body, _err := client.ListLiveRecordVideoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListMediaDNADeleteJobWithOptions(request *ListMediaDNADeleteJobRequest, runtime *util.RuntimeOptions) (_result *ListMediaDNADeleteJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListMediaDNADeleteJobResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListMediaDNADeleteJob"), tea.String("2017-03-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListMediaDNADeleteJob(request *ListMediaDNADeleteJobRequest) (_result *ListMediaDNADeleteJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListMediaDNADeleteJobResponse{}
	_body, _err := client.ListMediaDNADeleteJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListSnapshotsWithOptions(request *ListSnapshotsRequest, runtime *util.RuntimeOptions) (_result *ListSnapshotsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListSnapshotsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListSnapshots"), tea.String("2017-03-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListSnapshots(request *ListSnapshotsRequest) (_result *ListSnapshotsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListSnapshotsResponse{}
	_body, _err := client.ListSnapshotsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListTranscodeTaskWithOptions(request *ListTranscodeTaskRequest, runtime *util.RuntimeOptions) (_result *ListTranscodeTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListTranscodeTaskResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListTranscodeTask"), tea.String("2017-03-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListTranscodeTask(request *ListTranscodeTaskRequest) (_result *ListTranscodeTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTranscodeTaskResponse{}
	_body, _err := client.ListTranscodeTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListTranscodeTemplateGroupWithOptions(request *ListTranscodeTemplateGroupRequest, runtime *util.RuntimeOptions) (_result *ListTranscodeTemplateGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListTranscodeTemplateGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListTranscodeTemplateGroup"), tea.String("2017-03-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListTranscodeTemplateGroup(request *ListTranscodeTemplateGroupRequest) (_result *ListTranscodeTemplateGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTranscodeTemplateGroupResponse{}
	_body, _err := client.ListTranscodeTemplateGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListVodRealtimeLogDeliveryDomainsWithOptions(request *ListVodRealtimeLogDeliveryDomainsRequest, runtime *util.RuntimeOptions) (_result *ListVodRealtimeLogDeliveryDomainsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &ListVodRealtimeLogDeliveryDomainsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListVodRealtimeLogDeliveryDomains"), tea.String("2017-03-21"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListVodRealtimeLogDeliveryDomains(request *ListVodRealtimeLogDeliveryDomainsRequest) (_result *ListVodRealtimeLogDeliveryDomainsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListVodRealtimeLogDeliveryDomainsResponse{}
	_body, _err := client.ListVodRealtimeLogDeliveryDomainsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListVodRealtimeLogDeliveryInfosWithOptions(request *ListVodRealtimeLogDeliveryInfosRequest, runtime *util.RuntimeOptions) (_result *ListVodRealtimeLogDeliveryInfosResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &ListVodRealtimeLogDeliveryInfosResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListVodRealtimeLogDeliveryInfos"), tea.String("2017-03-21"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListVodRealtimeLogDeliveryInfos(request *ListVodRealtimeLogDeliveryInfosRequest) (_result *ListVodRealtimeLogDeliveryInfosResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListVodRealtimeLogDeliveryInfosResponse{}
	_body, _err := client.ListVodRealtimeLogDeliveryInfosWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListVodTemplateWithOptions(request *ListVodTemplateRequest, runtime *util.RuntimeOptions) (_result *ListVodTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListVodTemplateResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListVodTemplate"), tea.String("2017-03-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListVodTemplate(request *ListVodTemplateRequest) (_result *ListVodTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListVodTemplateResponse{}
	_body, _err := client.ListVodTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListWatermarkWithOptions(request *ListWatermarkRequest, runtime *util.RuntimeOptions) (_result *ListWatermarkResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListWatermarkResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListWatermark"), tea.String("2017-03-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListWatermark(request *ListWatermarkRequest) (_result *ListWatermarkResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListWatermarkResponse{}
	_body, _err := client.ListWatermarkWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) MoveAppResourceWithOptions(request *MoveAppResourceRequest, runtime *util.RuntimeOptions) (_result *MoveAppResourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &MoveAppResourceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("MoveAppResource"), tea.String("2017-03-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) MoveAppResource(request *MoveAppResourceRequest) (_result *MoveAppResourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &MoveAppResourceResponse{}
	_body, _err := client.MoveAppResourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PreloadVodObjectCachesWithOptions(request *PreloadVodObjectCachesRequest, runtime *util.RuntimeOptions) (_result *PreloadVodObjectCachesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &PreloadVodObjectCachesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("PreloadVodObjectCaches"), tea.String("2017-03-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PreloadVodObjectCaches(request *PreloadVodObjectCachesRequest) (_result *PreloadVodObjectCachesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PreloadVodObjectCachesResponse{}
	_body, _err := client.PreloadVodObjectCachesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ProduceEditingProjectVideoWithOptions(request *ProduceEditingProjectVideoRequest, runtime *util.RuntimeOptions) (_result *ProduceEditingProjectVideoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ProduceEditingProjectVideoResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ProduceEditingProjectVideo"), tea.String("2017-03-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ProduceEditingProjectVideo(request *ProduceEditingProjectVideoRequest) (_result *ProduceEditingProjectVideoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ProduceEditingProjectVideoResponse{}
	_body, _err := client.ProduceEditingProjectVideoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RefreshUploadVideoWithOptions(request *RefreshUploadVideoRequest, runtime *util.RuntimeOptions) (_result *RefreshUploadVideoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RefreshUploadVideoResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RefreshUploadVideo"), tea.String("2017-03-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RefreshUploadVideo(request *RefreshUploadVideoRequest) (_result *RefreshUploadVideoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RefreshUploadVideoResponse{}
	_body, _err := client.RefreshUploadVideoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RefreshVodObjectCachesWithOptions(request *RefreshVodObjectCachesRequest, runtime *util.RuntimeOptions) (_result *RefreshVodObjectCachesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RefreshVodObjectCachesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RefreshVodObjectCaches"), tea.String("2017-03-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RefreshVodObjectCaches(request *RefreshVodObjectCachesRequest) (_result *RefreshVodObjectCachesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RefreshVodObjectCachesResponse{}
	_body, _err := client.RefreshVodObjectCachesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RegisterMediaWithOptions(request *RegisterMediaRequest, runtime *util.RuntimeOptions) (_result *RegisterMediaResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RegisterMediaResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RegisterMedia"), tea.String("2017-03-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RegisterMedia(request *RegisterMediaRequest) (_result *RegisterMediaResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RegisterMediaResponse{}
	_body, _err := client.RegisterMediaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SearchEditingProjectWithOptions(request *SearchEditingProjectRequest, runtime *util.RuntimeOptions) (_result *SearchEditingProjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SearchEditingProjectResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SearchEditingProject"), tea.String("2017-03-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SearchEditingProject(request *SearchEditingProjectRequest) (_result *SearchEditingProjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SearchEditingProjectResponse{}
	_body, _err := client.SearchEditingProjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SearchMediaWithOptions(request *SearchMediaRequest, runtime *util.RuntimeOptions) (_result *SearchMediaResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SearchMediaResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SearchMedia"), tea.String("2017-03-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SearchMedia(request *SearchMediaRequest) (_result *SearchMediaResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SearchMediaResponse{}
	_body, _err := client.SearchMediaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetAuditSecurityIpWithOptions(request *SetAuditSecurityIpRequest, runtime *util.RuntimeOptions) (_result *SetAuditSecurityIpResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SetAuditSecurityIpResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SetAuditSecurityIp"), tea.String("2017-03-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetAuditSecurityIp(request *SetAuditSecurityIpRequest) (_result *SetAuditSecurityIpResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetAuditSecurityIpResponse{}
	_body, _err := client.SetAuditSecurityIpWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetCrossdomainContentWithOptions(request *SetCrossdomainContentRequest, runtime *util.RuntimeOptions) (_result *SetCrossdomainContentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SetCrossdomainContentResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SetCrossdomainContent"), tea.String("2017-03-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetCrossdomainContent(request *SetCrossdomainContentRequest) (_result *SetCrossdomainContentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetCrossdomainContentResponse{}
	_body, _err := client.SetCrossdomainContentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetDefaultAITemplateWithOptions(request *SetDefaultAITemplateRequest, runtime *util.RuntimeOptions) (_result *SetDefaultAITemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SetDefaultAITemplateResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SetDefaultAITemplate"), tea.String("2017-03-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetDefaultAITemplate(request *SetDefaultAITemplateRequest) (_result *SetDefaultAITemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetDefaultAITemplateResponse{}
	_body, _err := client.SetDefaultAITemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetDefaultTranscodeTemplateGroupWithOptions(request *SetDefaultTranscodeTemplateGroupRequest, runtime *util.RuntimeOptions) (_result *SetDefaultTranscodeTemplateGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SetDefaultTranscodeTemplateGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SetDefaultTranscodeTemplateGroup"), tea.String("2017-03-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetDefaultTranscodeTemplateGroup(request *SetDefaultTranscodeTemplateGroupRequest) (_result *SetDefaultTranscodeTemplateGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetDefaultTranscodeTemplateGroupResponse{}
	_body, _err := client.SetDefaultTranscodeTemplateGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetDefaultWatermarkWithOptions(request *SetDefaultWatermarkRequest, runtime *util.RuntimeOptions) (_result *SetDefaultWatermarkResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SetDefaultWatermarkResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SetDefaultWatermark"), tea.String("2017-03-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetDefaultWatermark(request *SetDefaultWatermarkRequest) (_result *SetDefaultWatermarkResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetDefaultWatermarkResponse{}
	_body, _err := client.SetDefaultWatermarkWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetEditingProjectMaterialsWithOptions(request *SetEditingProjectMaterialsRequest, runtime *util.RuntimeOptions) (_result *SetEditingProjectMaterialsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SetEditingProjectMaterialsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SetEditingProjectMaterials"), tea.String("2017-03-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetEditingProjectMaterials(request *SetEditingProjectMaterialsRequest) (_result *SetEditingProjectMaterialsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetEditingProjectMaterialsResponse{}
	_body, _err := client.SetEditingProjectMaterialsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetMessageCallbackWithOptions(request *SetMessageCallbackRequest, runtime *util.RuntimeOptions) (_result *SetMessageCallbackResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SetMessageCallbackResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SetMessageCallback"), tea.String("2017-03-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetMessageCallback(request *SetMessageCallbackRequest) (_result *SetMessageCallbackResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetMessageCallbackResponse{}
	_body, _err := client.SetMessageCallbackWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetVodDomainCertificateWithOptions(request *SetVodDomainCertificateRequest, runtime *util.RuntimeOptions) (_result *SetVodDomainCertificateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SetVodDomainCertificateResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SetVodDomainCertificate"), tea.String("2017-03-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetVodDomainCertificate(request *SetVodDomainCertificateRequest) (_result *SetVodDomainCertificateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetVodDomainCertificateResponse{}
	_body, _err := client.SetVodDomainCertificateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SubmitAICaptionExtractionJobWithOptions(request *SubmitAICaptionExtractionJobRequest, runtime *util.RuntimeOptions) (_result *SubmitAICaptionExtractionJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SubmitAICaptionExtractionJobResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SubmitAICaptionExtractionJob"), tea.String("2017-03-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SubmitAICaptionExtractionJob(request *SubmitAICaptionExtractionJobRequest) (_result *SubmitAICaptionExtractionJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SubmitAICaptionExtractionJobResponse{}
	_body, _err := client.SubmitAICaptionExtractionJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SubmitAIImageAuditJobWithOptions(request *SubmitAIImageAuditJobRequest, runtime *util.RuntimeOptions) (_result *SubmitAIImageAuditJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SubmitAIImageAuditJobResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SubmitAIImageAuditJob"), tea.String("2017-03-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SubmitAIImageAuditJob(request *SubmitAIImageAuditJobRequest) (_result *SubmitAIImageAuditJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SubmitAIImageAuditJobResponse{}
	_body, _err := client.SubmitAIImageAuditJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SubmitAIImageJobWithOptions(request *SubmitAIImageJobRequest, runtime *util.RuntimeOptions) (_result *SubmitAIImageJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SubmitAIImageJobResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SubmitAIImageJob"), tea.String("2017-03-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SubmitAIImageJob(request *SubmitAIImageJobRequest) (_result *SubmitAIImageJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SubmitAIImageJobResponse{}
	_body, _err := client.SubmitAIImageJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SubmitAIJobWithOptions(request *SubmitAIJobRequest, runtime *util.RuntimeOptions) (_result *SubmitAIJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SubmitAIJobResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SubmitAIJob"), tea.String("2017-03-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SubmitAIJob(request *SubmitAIJobRequest) (_result *SubmitAIJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SubmitAIJobResponse{}
	_body, _err := client.SubmitAIJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SubmitAIMediaAuditJobWithOptions(request *SubmitAIMediaAuditJobRequest, runtime *util.RuntimeOptions) (_result *SubmitAIMediaAuditJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SubmitAIMediaAuditJobResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SubmitAIMediaAuditJob"), tea.String("2017-03-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SubmitAIMediaAuditJob(request *SubmitAIMediaAuditJobRequest) (_result *SubmitAIMediaAuditJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SubmitAIMediaAuditJobResponse{}
	_body, _err := client.SubmitAIMediaAuditJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SubmitDetectionJobWithOptions(request *SubmitDetectionJobRequest, runtime *util.RuntimeOptions) (_result *SubmitDetectionJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SubmitDetectionJobResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SubmitDetectionJob"), tea.String("2017-03-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SubmitDetectionJob(request *SubmitDetectionJobRequest) (_result *SubmitDetectionJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SubmitDetectionJobResponse{}
	_body, _err := client.SubmitDetectionJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SubmitDynamicImageJobWithOptions(request *SubmitDynamicImageJobRequest, runtime *util.RuntimeOptions) (_result *SubmitDynamicImageJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SubmitDynamicImageJobResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SubmitDynamicImageJob"), tea.String("2017-03-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SubmitDynamicImageJob(request *SubmitDynamicImageJobRequest) (_result *SubmitDynamicImageJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SubmitDynamicImageJobResponse{}
	_body, _err := client.SubmitDynamicImageJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SubmitLiveEditingWithOptions(request *SubmitLiveEditingRequest, runtime *util.RuntimeOptions) (_result *SubmitLiveEditingResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SubmitLiveEditingResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SubmitLiveEditing"), tea.String("2017-03-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SubmitLiveEditing(request *SubmitLiveEditingRequest) (_result *SubmitLiveEditingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SubmitLiveEditingResponse{}
	_body, _err := client.SubmitLiveEditingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SubmitMediaDNADeleteJobWithOptions(request *SubmitMediaDNADeleteJobRequest, runtime *util.RuntimeOptions) (_result *SubmitMediaDNADeleteJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SubmitMediaDNADeleteJobResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SubmitMediaDNADeleteJob"), tea.String("2017-03-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SubmitMediaDNADeleteJob(request *SubmitMediaDNADeleteJobRequest) (_result *SubmitMediaDNADeleteJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SubmitMediaDNADeleteJobResponse{}
	_body, _err := client.SubmitMediaDNADeleteJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SubmitPreprocessJobsWithOptions(request *SubmitPreprocessJobsRequest, runtime *util.RuntimeOptions) (_result *SubmitPreprocessJobsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SubmitPreprocessJobsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SubmitPreprocessJobs"), tea.String("2017-03-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SubmitPreprocessJobs(request *SubmitPreprocessJobsRequest) (_result *SubmitPreprocessJobsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SubmitPreprocessJobsResponse{}
	_body, _err := client.SubmitPreprocessJobsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SubmitSnapshotJobWithOptions(request *SubmitSnapshotJobRequest, runtime *util.RuntimeOptions) (_result *SubmitSnapshotJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SubmitSnapshotJobResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SubmitSnapshotJob"), tea.String("2017-03-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SubmitSnapshotJob(request *SubmitSnapshotJobRequest) (_result *SubmitSnapshotJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SubmitSnapshotJobResponse{}
	_body, _err := client.SubmitSnapshotJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SubmitTranscodeJobsWithOptions(request *SubmitTranscodeJobsRequest, runtime *util.RuntimeOptions) (_result *SubmitTranscodeJobsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SubmitTranscodeJobsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SubmitTranscodeJobs"), tea.String("2017-03-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SubmitTranscodeJobs(request *SubmitTranscodeJobsRequest) (_result *SubmitTranscodeJobsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SubmitTranscodeJobsResponse{}
	_body, _err := client.SubmitTranscodeJobsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SubmitWorkflowJobWithOptions(request *SubmitWorkflowJobRequest, runtime *util.RuntimeOptions) (_result *SubmitWorkflowJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SubmitWorkflowJobResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SubmitWorkflowJob"), tea.String("2017-03-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SubmitWorkflowJob(request *SubmitWorkflowJobRequest) (_result *SubmitWorkflowJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SubmitWorkflowJobResponse{}
	_body, _err := client.SubmitWorkflowJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) TagVodResourcesWithOptions(request *TagVodResourcesRequest, runtime *util.RuntimeOptions) (_result *TagVodResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &TagVodResourcesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("TagVodResources"), tea.String("2017-03-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) TagVodResources(request *TagVodResourcesRequest) (_result *TagVodResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &TagVodResourcesResponse{}
	_body, _err := client.TagVodResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UnTagVodResourcesWithOptions(request *UnTagVodResourcesRequest, runtime *util.RuntimeOptions) (_result *UnTagVodResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UnTagVodResourcesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UnTagVodResources"), tea.String("2017-03-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UnTagVodResources(request *UnTagVodResourcesRequest) (_result *UnTagVodResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UnTagVodResourcesResponse{}
	_body, _err := client.UnTagVodResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateAITemplateWithOptions(request *UpdateAITemplateRequest, runtime *util.RuntimeOptions) (_result *UpdateAITemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateAITemplateResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateAITemplate"), tea.String("2017-03-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateAITemplate(request *UpdateAITemplateRequest) (_result *UpdateAITemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateAITemplateResponse{}
	_body, _err := client.UpdateAITemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateAppInfoWithOptions(request *UpdateAppInfoRequest, runtime *util.RuntimeOptions) (_result *UpdateAppInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateAppInfoResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateAppInfo"), tea.String("2017-03-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateAppInfo(request *UpdateAppInfoRequest) (_result *UpdateAppInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateAppInfoResponse{}
	_body, _err := client.UpdateAppInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateAttachedMediaInfosWithOptions(request *UpdateAttachedMediaInfosRequest, runtime *util.RuntimeOptions) (_result *UpdateAttachedMediaInfosResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateAttachedMediaInfosResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateAttachedMediaInfos"), tea.String("2017-03-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateAttachedMediaInfos(request *UpdateAttachedMediaInfosRequest) (_result *UpdateAttachedMediaInfosResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateAttachedMediaInfosResponse{}
	_body, _err := client.UpdateAttachedMediaInfosWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateCategoryWithOptions(request *UpdateCategoryRequest, runtime *util.RuntimeOptions) (_result *UpdateCategoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateCategoryResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateCategory"), tea.String("2017-03-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateCategory(request *UpdateCategoryRequest) (_result *UpdateCategoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateCategoryResponse{}
	_body, _err := client.UpdateCategoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateDetectionJobWithOptions(request *UpdateDetectionJobRequest, runtime *util.RuntimeOptions) (_result *UpdateDetectionJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateDetectionJobResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateDetectionJob"), tea.String("2017-03-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateDetectionJob(request *UpdateDetectionJobRequest) (_result *UpdateDetectionJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateDetectionJobResponse{}
	_body, _err := client.UpdateDetectionJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateDetectionTemplateWithOptions(request *UpdateDetectionTemplateRequest, runtime *util.RuntimeOptions) (_result *UpdateDetectionTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateDetectionTemplateResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateDetectionTemplate"), tea.String("2017-03-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateDetectionTemplate(request *UpdateDetectionTemplateRequest) (_result *UpdateDetectionTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateDetectionTemplateResponse{}
	_body, _err := client.UpdateDetectionTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateEditingProjectWithOptions(request *UpdateEditingProjectRequest, runtime *util.RuntimeOptions) (_result *UpdateEditingProjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateEditingProjectResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateEditingProject"), tea.String("2017-03-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateEditingProject(request *UpdateEditingProjectRequest) (_result *UpdateEditingProjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateEditingProjectResponse{}
	_body, _err := client.UpdateEditingProjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateImageInfosWithOptions(request *UpdateImageInfosRequest, runtime *util.RuntimeOptions) (_result *UpdateImageInfosResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateImageInfosResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateImageInfos"), tea.String("2017-03-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateImageInfos(request *UpdateImageInfosRequest) (_result *UpdateImageInfosResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateImageInfosResponse{}
	_body, _err := client.UpdateImageInfosWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateTranscodeTemplateGroupWithOptions(request *UpdateTranscodeTemplateGroupRequest, runtime *util.RuntimeOptions) (_result *UpdateTranscodeTemplateGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateTranscodeTemplateGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateTranscodeTemplateGroup"), tea.String("2017-03-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateTranscodeTemplateGroup(request *UpdateTranscodeTemplateGroupRequest) (_result *UpdateTranscodeTemplateGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateTranscodeTemplateGroupResponse{}
	_body, _err := client.UpdateTranscodeTemplateGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateVideoInfoWithOptions(request *UpdateVideoInfoRequest, runtime *util.RuntimeOptions) (_result *UpdateVideoInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateVideoInfoResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateVideoInfo"), tea.String("2017-03-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateVideoInfo(request *UpdateVideoInfoRequest) (_result *UpdateVideoInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateVideoInfoResponse{}
	_body, _err := client.UpdateVideoInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateVideoInfosWithOptions(request *UpdateVideoInfosRequest, runtime *util.RuntimeOptions) (_result *UpdateVideoInfosResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateVideoInfosResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateVideoInfos"), tea.String("2017-03-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateVideoInfos(request *UpdateVideoInfosRequest) (_result *UpdateVideoInfosResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateVideoInfosResponse{}
	_body, _err := client.UpdateVideoInfosWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateVodDomainWithOptions(request *UpdateVodDomainRequest, runtime *util.RuntimeOptions) (_result *UpdateVodDomainResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateVodDomainResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateVodDomain"), tea.String("2017-03-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateVodDomain(request *UpdateVodDomainRequest) (_result *UpdateVodDomainResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateVodDomainResponse{}
	_body, _err := client.UpdateVodDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateVodTemplateWithOptions(request *UpdateVodTemplateRequest, runtime *util.RuntimeOptions) (_result *UpdateVodTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateVodTemplateResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateVodTemplate"), tea.String("2017-03-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateVodTemplate(request *UpdateVodTemplateRequest) (_result *UpdateVodTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateVodTemplateResponse{}
	_body, _err := client.UpdateVodTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateWatermarkWithOptions(request *UpdateWatermarkRequest, runtime *util.RuntimeOptions) (_result *UpdateWatermarkResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateWatermarkResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateWatermark"), tea.String("2017-03-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateWatermark(request *UpdateWatermarkRequest) (_result *UpdateWatermarkResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateWatermarkResponse{}
	_body, _err := client.UpdateWatermarkWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UploadMediaByURLWithOptions(request *UploadMediaByURLRequest, runtime *util.RuntimeOptions) (_result *UploadMediaByURLResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UploadMediaByURLResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UploadMediaByURL"), tea.String("2017-03-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UploadMediaByURL(request *UploadMediaByURLRequest) (_result *UploadMediaByURLResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UploadMediaByURLResponse{}
	_body, _err := client.UploadMediaByURLWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UploadStreamByURLWithOptions(request *UploadStreamByURLRequest, runtime *util.RuntimeOptions) (_result *UploadStreamByURLResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UploadStreamByURLResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UploadStreamByURL"), tea.String("2017-03-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UploadStreamByURL(request *UploadStreamByURLRequest) (_result *UploadStreamByURLResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UploadStreamByURLResponse{}
	_body, _err := client.UploadStreamByURLWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) VerifyVodDomainOwnerWithOptions(request *VerifyVodDomainOwnerRequest, runtime *util.RuntimeOptions) (_result *VerifyVodDomainOwnerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &VerifyVodDomainOwnerResponse{}
	_body, _err := client.DoRPCRequest(tea.String("VerifyVodDomainOwner"), tea.String("2017-03-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) VerifyVodDomainOwner(request *VerifyVodDomainOwnerRequest) (_result *VerifyVodDomainOwnerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &VerifyVodDomainOwnerResponse{}
	_body, _err := client.VerifyVodDomainOwnerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
