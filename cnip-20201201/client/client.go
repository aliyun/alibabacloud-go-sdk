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

type ProductComponentRelationDetail struct {
	// appVersion
	AppVersion *string `json:"appVersion,omitempty" xml:"appVersion,omitempty"`
	// category
	Category *string `json:"category,omitempty" xml:"category,omitempty"`
	// class
	Class *string `json:"class,omitempty" xml:"class,omitempty"`
	// componentName
	ComponentName *string `json:"componentName,omitempty" xml:"componentName,omitempty"`
	// componentOrchestrationValues
	ComponentOrchestrationValues *string `json:"componentOrchestrationValues,omitempty" xml:"componentOrchestrationValues,omitempty"`
	// componentUID
	ComponentUID *string `json:"componentUID,omitempty" xml:"componentUID,omitempty"`
	// componentVersionUID
	ComponentVersionUID *string `json:"componentVersionUID,omitempty" xml:"componentVersionUID,omitempty"`
	// createdAt
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// description
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// documents
	Documents *string `json:"documents,omitempty" xml:"documents,omitempty"`
	// enable
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
	// imagesMapping
	ImagesMapping *string `json:"imagesMapping,omitempty" xml:"imagesMapping,omitempty"`
	// namespace
	Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
	// orchestrationType
	OrchestrationType *string `json:"orchestrationType,omitempty" xml:"orchestrationType,omitempty"`
	// parentComponent
	ParentComponent *bool `json:"parentComponent,omitempty" xml:"parentComponent,omitempty"`
	// parentComponentVersionRelationUID
	ParentComponentVersionRelationUID *string `json:"parentComponentVersionRelationUID,omitempty" xml:"parentComponentVersionRelationUID,omitempty"`
	// parentComponentVersionUID
	ParentComponentVersionUID *string `json:"parentComponentVersionUID,omitempty" xml:"parentComponentVersionUID,omitempty"`
	// priority
	Priority *int32 `json:"priority,omitempty" xml:"priority,omitempty"`
	// productVersionUID
	ProductVersionUID *string `json:"productVersionUID,omitempty" xml:"productVersionUID,omitempty"`
	// provider
	Provider *string `json:"provider,omitempty" xml:"provider,omitempty"`
	// public
	Public *bool `json:"public,omitempty" xml:"public,omitempty"`
	// readme
	Readme *string `json:"readme,omitempty" xml:"readme,omitempty"`
	// relationUID
	RelationUID *string `json:"relationUID,omitempty" xml:"relationUID,omitempty"`
	// releaseName
	ReleaseName *string `json:"releaseName,omitempty" xml:"releaseName,omitempty"`
	// resources
	Resources *string `json:"resources,omitempty" xml:"resources,omitempty"`
	// sequence
	Sequence *int32 `json:"sequence,omitempty" xml:"sequence,omitempty"`
	// singleton
	Singleton *bool `json:"singleton,omitempty" xml:"singleton,omitempty"`
	// source
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
	// version
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s ProductComponentRelationDetail) String() string {
	return tea.Prettify(s)
}

func (s ProductComponentRelationDetail) GoString() string {
	return s.String()
}

func (s *ProductComponentRelationDetail) SetAppVersion(v string) *ProductComponentRelationDetail {
	s.AppVersion = &v
	return s
}

func (s *ProductComponentRelationDetail) SetCategory(v string) *ProductComponentRelationDetail {
	s.Category = &v
	return s
}

func (s *ProductComponentRelationDetail) SetClass(v string) *ProductComponentRelationDetail {
	s.Class = &v
	return s
}

func (s *ProductComponentRelationDetail) SetComponentName(v string) *ProductComponentRelationDetail {
	s.ComponentName = &v
	return s
}

func (s *ProductComponentRelationDetail) SetComponentOrchestrationValues(v string) *ProductComponentRelationDetail {
	s.ComponentOrchestrationValues = &v
	return s
}

func (s *ProductComponentRelationDetail) SetComponentUID(v string) *ProductComponentRelationDetail {
	s.ComponentUID = &v
	return s
}

func (s *ProductComponentRelationDetail) SetComponentVersionUID(v string) *ProductComponentRelationDetail {
	s.ComponentVersionUID = &v
	return s
}

func (s *ProductComponentRelationDetail) SetCreatedAt(v string) *ProductComponentRelationDetail {
	s.CreatedAt = &v
	return s
}

func (s *ProductComponentRelationDetail) SetDescription(v string) *ProductComponentRelationDetail {
	s.Description = &v
	return s
}

func (s *ProductComponentRelationDetail) SetDocuments(v string) *ProductComponentRelationDetail {
	s.Documents = &v
	return s
}

func (s *ProductComponentRelationDetail) SetEnable(v bool) *ProductComponentRelationDetail {
	s.Enable = &v
	return s
}

func (s *ProductComponentRelationDetail) SetImagesMapping(v string) *ProductComponentRelationDetail {
	s.ImagesMapping = &v
	return s
}

func (s *ProductComponentRelationDetail) SetNamespace(v string) *ProductComponentRelationDetail {
	s.Namespace = &v
	return s
}

func (s *ProductComponentRelationDetail) SetOrchestrationType(v string) *ProductComponentRelationDetail {
	s.OrchestrationType = &v
	return s
}

func (s *ProductComponentRelationDetail) SetParentComponent(v bool) *ProductComponentRelationDetail {
	s.ParentComponent = &v
	return s
}

func (s *ProductComponentRelationDetail) SetParentComponentVersionRelationUID(v string) *ProductComponentRelationDetail {
	s.ParentComponentVersionRelationUID = &v
	return s
}

func (s *ProductComponentRelationDetail) SetParentComponentVersionUID(v string) *ProductComponentRelationDetail {
	s.ParentComponentVersionUID = &v
	return s
}

func (s *ProductComponentRelationDetail) SetPriority(v int32) *ProductComponentRelationDetail {
	s.Priority = &v
	return s
}

func (s *ProductComponentRelationDetail) SetProductVersionUID(v string) *ProductComponentRelationDetail {
	s.ProductVersionUID = &v
	return s
}

func (s *ProductComponentRelationDetail) SetProvider(v string) *ProductComponentRelationDetail {
	s.Provider = &v
	return s
}

func (s *ProductComponentRelationDetail) SetPublic(v bool) *ProductComponentRelationDetail {
	s.Public = &v
	return s
}

func (s *ProductComponentRelationDetail) SetReadme(v string) *ProductComponentRelationDetail {
	s.Readme = &v
	return s
}

func (s *ProductComponentRelationDetail) SetRelationUID(v string) *ProductComponentRelationDetail {
	s.RelationUID = &v
	return s
}

func (s *ProductComponentRelationDetail) SetReleaseName(v string) *ProductComponentRelationDetail {
	s.ReleaseName = &v
	return s
}

func (s *ProductComponentRelationDetail) SetResources(v string) *ProductComponentRelationDetail {
	s.Resources = &v
	return s
}

func (s *ProductComponentRelationDetail) SetSequence(v int32) *ProductComponentRelationDetail {
	s.Sequence = &v
	return s
}

func (s *ProductComponentRelationDetail) SetSingleton(v bool) *ProductComponentRelationDetail {
	s.Singleton = &v
	return s
}

func (s *ProductComponentRelationDetail) SetSource(v string) *ProductComponentRelationDetail {
	s.Source = &v
	return s
}

func (s *ProductComponentRelationDetail) SetVersion(v string) *ProductComponentRelationDetail {
	s.Version = &v
	return s
}

type FoundationVersion struct {
	// description
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// documents
	Documents *string `json:"documents,omitempty" xml:"documents,omitempty"`
	// name
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// platforms
	Platforms *FoundationVersionPlatforms `json:"platforms,omitempty" xml:"platforms,omitempty" type:"Struct"`
	// status
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// uid
	Uid *string `json:"uid,omitempty" xml:"uid,omitempty"`
	// version
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s FoundationVersion) String() string {
	return tea.Prettify(s)
}

func (s FoundationVersion) GoString() string {
	return s.String()
}

func (s *FoundationVersion) SetDescription(v string) *FoundationVersion {
	s.Description = &v
	return s
}

func (s *FoundationVersion) SetDocuments(v string) *FoundationVersion {
	s.Documents = &v
	return s
}

func (s *FoundationVersion) SetName(v string) *FoundationVersion {
	s.Name = &v
	return s
}

func (s *FoundationVersion) SetPlatforms(v *FoundationVersionPlatforms) *FoundationVersion {
	s.Platforms = v
	return s
}

func (s *FoundationVersion) SetStatus(v string) *FoundationVersion {
	s.Status = &v
	return s
}

func (s *FoundationVersion) SetUid(v string) *FoundationVersion {
	s.Uid = &v
	return s
}

func (s *FoundationVersion) SetVersion(v string) *FoundationVersion {
	s.Version = &v
	return s
}

type FoundationVersionPlatforms struct {
	// architecture
	Architecture *string `json:"architecture,omitempty" xml:"architecture,omitempty"`
	// os
	Os *string `json:"os,omitempty" xml:"os,omitempty"`
}

func (s FoundationVersionPlatforms) String() string {
	return tea.Prettify(s)
}

func (s FoundationVersionPlatforms) GoString() string {
	return s.String()
}

func (s *FoundationVersionPlatforms) SetArchitecture(v string) *FoundationVersionPlatforms {
	s.Architecture = &v
	return s
}

func (s *FoundationVersionPlatforms) SetOs(v string) *FoundationVersionPlatforms {
	s.Os = &v
	return s
}

type Platform struct {
	// architecture
	Architecture *string `json:"architecture,omitempty" xml:"architecture,omitempty"`
	// os
	Os *string `json:"os,omitempty" xml:"os,omitempty"`
}

func (s Platform) String() string {
	return tea.Prettify(s)
}

func (s Platform) GoString() string {
	return s.String()
}

func (s *Platform) SetArchitecture(v string) *Platform {
	s.Architecture = &v
	return s
}

func (s *Platform) SetOs(v string) *Platform {
	s.Os = &v
	return s
}

type ComponentVersion struct {
	// appVersion
	AppVersion *string `json:"appVersion,omitempty" xml:"appVersion,omitempty"`
	// componentName
	ComponentName *string `json:"componentName,omitempty" xml:"componentName,omitempty"`
	// componentUID
	ComponentUID *string `json:"componentUID,omitempty" xml:"componentUID,omitempty"`
	// description
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// documents
	Documents *string `json:"documents,omitempty" xml:"documents,omitempty"`
	// imagesMapping
	ImagesMapping *string `json:"imagesMapping,omitempty" xml:"imagesMapping,omitempty"`
	// namespace
	Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
	// orchestrationType
	OrchestrationType *string `json:"orchestrationType,omitempty" xml:"orchestrationType,omitempty"`
	// orchestrationValues
	OrchestrationValues *string `json:"orchestrationValues,omitempty" xml:"orchestrationValues,omitempty"`
	// packageURL
	PackageURL *string `json:"packageURL,omitempty" xml:"packageURL,omitempty"`
	// parentComponent
	ParentComponent *bool `json:"parentComponent,omitempty" xml:"parentComponent,omitempty"`
	// platforms
	Platforms []*Platform `json:"platforms,omitempty" xml:"platforms,omitempty" type:"Repeated"`
	// readme
	Readme *string `json:"readme,omitempty" xml:"readme,omitempty"`
	// resources
	Resources *string `json:"resources,omitempty" xml:"resources,omitempty"`
	// source
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
	// uid
	Uid *string `json:"uid,omitempty" xml:"uid,omitempty"`
	// version
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s ComponentVersion) String() string {
	return tea.Prettify(s)
}

func (s ComponentVersion) GoString() string {
	return s.String()
}

func (s *ComponentVersion) SetAppVersion(v string) *ComponentVersion {
	s.AppVersion = &v
	return s
}

func (s *ComponentVersion) SetComponentName(v string) *ComponentVersion {
	s.ComponentName = &v
	return s
}

func (s *ComponentVersion) SetComponentUID(v string) *ComponentVersion {
	s.ComponentUID = &v
	return s
}

func (s *ComponentVersion) SetDescription(v string) *ComponentVersion {
	s.Description = &v
	return s
}

func (s *ComponentVersion) SetDocuments(v string) *ComponentVersion {
	s.Documents = &v
	return s
}

func (s *ComponentVersion) SetImagesMapping(v string) *ComponentVersion {
	s.ImagesMapping = &v
	return s
}

func (s *ComponentVersion) SetNamespace(v string) *ComponentVersion {
	s.Namespace = &v
	return s
}

func (s *ComponentVersion) SetOrchestrationType(v string) *ComponentVersion {
	s.OrchestrationType = &v
	return s
}

func (s *ComponentVersion) SetOrchestrationValues(v string) *ComponentVersion {
	s.OrchestrationValues = &v
	return s
}

func (s *ComponentVersion) SetPackageURL(v string) *ComponentVersion {
	s.PackageURL = &v
	return s
}

func (s *ComponentVersion) SetParentComponent(v bool) *ComponentVersion {
	s.ParentComponent = &v
	return s
}

func (s *ComponentVersion) SetPlatforms(v []*Platform) *ComponentVersion {
	s.Platforms = v
	return s
}

func (s *ComponentVersion) SetReadme(v string) *ComponentVersion {
	s.Readme = &v
	return s
}

func (s *ComponentVersion) SetResources(v string) *ComponentVersion {
	s.Resources = &v
	return s
}

func (s *ComponentVersion) SetSource(v string) *ComponentVersion {
	s.Source = &v
	return s
}

func (s *ComponentVersion) SetUid(v string) *ComponentVersion {
	s.Uid = &v
	return s
}

func (s *ComponentVersion) SetVersion(v string) *ComponentVersion {
	s.Version = &v
	return s
}

type GetEnvironmentResponseBody struct {
	Data    *GetEnvironmentResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	ErrCode *string                         `json:"errCode,omitempty" xml:"errCode,omitempty"`
	ErrMsg  *string                         `json:"errMsg,omitempty" xml:"errMsg,omitempty"`
	Success *bool                           `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetEnvironmentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetEnvironmentResponseBody) GoString() string {
	return s.String()
}

func (s *GetEnvironmentResponseBody) SetData(v *GetEnvironmentResponseBodyData) *GetEnvironmentResponseBody {
	s.Data = v
	return s
}

func (s *GetEnvironmentResponseBody) SetErrCode(v string) *GetEnvironmentResponseBody {
	s.ErrCode = &v
	return s
}

func (s *GetEnvironmentResponseBody) SetErrMsg(v string) *GetEnvironmentResponseBody {
	s.ErrMsg = &v
	return s
}

func (s *GetEnvironmentResponseBody) SetSuccess(v bool) *GetEnvironmentResponseBody {
	s.Success = &v
	return s
}

type GetEnvironmentResponseBodyData struct {
	ClusterId      *string                                 `json:"clusterId,omitempty" xml:"clusterId,omitempty"`
	CreatedAt      *string                                 `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	Description    *string                                 `json:"description,omitempty" xml:"description,omitempty"`
	InstanceList   *string                                 `json:"instanceList,omitempty" xml:"instanceList,omitempty"`
	Location       *string                                 `json:"location,omitempty" xml:"location,omitempty"`
	Name           *string                                 `json:"name,omitempty" xml:"name,omitempty"`
	Platform       *GetEnvironmentResponseBodyDataPlatform `json:"platform,omitempty" xml:"platform,omitempty" type:"Struct"`
	ProductName    *string                                 `json:"productName,omitempty" xml:"productName,omitempty"`
	ProductVersion *string                                 `json:"productVersion,omitempty" xml:"productVersion,omitempty"`
	Uid            *string                                 `json:"uid,omitempty" xml:"uid,omitempty"`
	VendorConfig   *string                                 `json:"vendorConfig,omitempty" xml:"vendorConfig,omitempty"`
	VendorType     *string                                 `json:"vendorType,omitempty" xml:"vendorType,omitempty"`
	InstanceStatus *string                                 `json:"instanceStatus,omitempty" xml:"instanceStatus,omitempty"`
}

func (s GetEnvironmentResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetEnvironmentResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetEnvironmentResponseBodyData) SetClusterId(v string) *GetEnvironmentResponseBodyData {
	s.ClusterId = &v
	return s
}

func (s *GetEnvironmentResponseBodyData) SetCreatedAt(v string) *GetEnvironmentResponseBodyData {
	s.CreatedAt = &v
	return s
}

func (s *GetEnvironmentResponseBodyData) SetDescription(v string) *GetEnvironmentResponseBodyData {
	s.Description = &v
	return s
}

func (s *GetEnvironmentResponseBodyData) SetInstanceList(v string) *GetEnvironmentResponseBodyData {
	s.InstanceList = &v
	return s
}

func (s *GetEnvironmentResponseBodyData) SetLocation(v string) *GetEnvironmentResponseBodyData {
	s.Location = &v
	return s
}

func (s *GetEnvironmentResponseBodyData) SetName(v string) *GetEnvironmentResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetEnvironmentResponseBodyData) SetPlatform(v *GetEnvironmentResponseBodyDataPlatform) *GetEnvironmentResponseBodyData {
	s.Platform = v
	return s
}

func (s *GetEnvironmentResponseBodyData) SetProductName(v string) *GetEnvironmentResponseBodyData {
	s.ProductName = &v
	return s
}

func (s *GetEnvironmentResponseBodyData) SetProductVersion(v string) *GetEnvironmentResponseBodyData {
	s.ProductVersion = &v
	return s
}

func (s *GetEnvironmentResponseBodyData) SetUid(v string) *GetEnvironmentResponseBodyData {
	s.Uid = &v
	return s
}

func (s *GetEnvironmentResponseBodyData) SetVendorConfig(v string) *GetEnvironmentResponseBodyData {
	s.VendorConfig = &v
	return s
}

func (s *GetEnvironmentResponseBodyData) SetVendorType(v string) *GetEnvironmentResponseBodyData {
	s.VendorType = &v
	return s
}

func (s *GetEnvironmentResponseBodyData) SetInstanceStatus(v string) *GetEnvironmentResponseBodyData {
	s.InstanceStatus = &v
	return s
}

type GetEnvironmentResponseBodyDataPlatform struct {
	Architecture *string `json:"architecture,omitempty" xml:"architecture,omitempty"`
	Os           *string `json:"os,omitempty" xml:"os,omitempty"`
}

func (s GetEnvironmentResponseBodyDataPlatform) String() string {
	return tea.Prettify(s)
}

func (s GetEnvironmentResponseBodyDataPlatform) GoString() string {
	return s.String()
}

func (s *GetEnvironmentResponseBodyDataPlatform) SetArchitecture(v string) *GetEnvironmentResponseBodyDataPlatform {
	s.Architecture = &v
	return s
}

func (s *GetEnvironmentResponseBodyDataPlatform) SetOs(v string) *GetEnvironmentResponseBodyDataPlatform {
	s.Os = &v
	return s
}

type GetEnvironmentResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetEnvironmentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetEnvironmentResponse) String() string {
	return tea.Prettify(s)
}

func (s GetEnvironmentResponse) GoString() string {
	return s.String()
}

func (s *GetEnvironmentResponse) SetHeaders(v map[string]*string) *GetEnvironmentResponse {
	s.Headers = v
	return s
}

func (s *GetEnvironmentResponse) SetBody(v *GetEnvironmentResponseBody) *GetEnvironmentResponse {
	s.Body = v
	return s
}

type ListProductVersionRelatedComponentVersionsResponseBody struct {
	Data    []*ProductComponentRelationDetail `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	ErrCode *string                           `json:"errCode,omitempty" xml:"errCode,omitempty"`
	ErrMsg  *string                           `json:"errMsg,omitempty" xml:"errMsg,omitempty"`
	Success *bool                             `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ListProductVersionRelatedComponentVersionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListProductVersionRelatedComponentVersionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListProductVersionRelatedComponentVersionsResponseBody) SetData(v []*ProductComponentRelationDetail) *ListProductVersionRelatedComponentVersionsResponseBody {
	s.Data = v
	return s
}

func (s *ListProductVersionRelatedComponentVersionsResponseBody) SetErrCode(v string) *ListProductVersionRelatedComponentVersionsResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ListProductVersionRelatedComponentVersionsResponseBody) SetErrMsg(v string) *ListProductVersionRelatedComponentVersionsResponseBody {
	s.ErrMsg = &v
	return s
}

func (s *ListProductVersionRelatedComponentVersionsResponseBody) SetSuccess(v bool) *ListProductVersionRelatedComponentVersionsResponseBody {
	s.Success = &v
	return s
}

type ListProductVersionRelatedComponentVersionsResponse struct {
	Headers map[string]*string                                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListProductVersionRelatedComponentVersionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListProductVersionRelatedComponentVersionsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListProductVersionRelatedComponentVersionsResponse) GoString() string {
	return s.String()
}

func (s *ListProductVersionRelatedComponentVersionsResponse) SetHeaders(v map[string]*string) *ListProductVersionRelatedComponentVersionsResponse {
	s.Headers = v
	return s
}

func (s *ListProductVersionRelatedComponentVersionsResponse) SetBody(v *ListProductVersionRelatedComponentVersionsResponseBody) *ListProductVersionRelatedComponentVersionsResponse {
	s.Body = v
	return s
}

type GetComponentVersionChildrenResponseBody struct {
	Data    []*GetComponentVersionChildrenResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	ErrCode *string                                        `json:"errCode,omitempty" xml:"errCode,omitempty"`
	ErrMsg  *string                                        `json:"errMsg,omitempty" xml:"errMsg,omitempty"`
	Status  *string                                        `json:"status,omitempty" xml:"status,omitempty"`
	Success *bool                                          `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetComponentVersionChildrenResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetComponentVersionChildrenResponseBody) GoString() string {
	return s.String()
}

func (s *GetComponentVersionChildrenResponseBody) SetData(v []*GetComponentVersionChildrenResponseBodyData) *GetComponentVersionChildrenResponseBody {
	s.Data = v
	return s
}

func (s *GetComponentVersionChildrenResponseBody) SetErrCode(v string) *GetComponentVersionChildrenResponseBody {
	s.ErrCode = &v
	return s
}

func (s *GetComponentVersionChildrenResponseBody) SetErrMsg(v string) *GetComponentVersionChildrenResponseBody {
	s.ErrMsg = &v
	return s
}

func (s *GetComponentVersionChildrenResponseBody) SetStatus(v string) *GetComponentVersionChildrenResponseBody {
	s.Status = &v
	return s
}

func (s *GetComponentVersionChildrenResponseBody) SetSuccess(v bool) *GetComponentVersionChildrenResponseBody {
	s.Success = &v
	return s
}

type GetComponentVersionChildrenResponseBodyData struct {
	ComponentName              *string   `json:"componentName,omitempty" xml:"componentName,omitempty"`
	ComponentUID               *string   `json:"componentUID,omitempty" xml:"componentUID,omitempty"`
	Description                *string   `json:"description,omitempty" xml:"description,omitempty"`
	Documents                  []*string `json:"documents,omitempty" xml:"documents,omitempty" type:"Repeated"`
	OrchestrationValues        *string   `json:"orchestrationValues,omitempty" xml:"orchestrationValues,omitempty"`
	PackageURL                 *string   `json:"packageURL,omitempty" xml:"packageURL,omitempty"`
	ParentComponent            *bool     `json:"parentComponent,omitempty" xml:"parentComponent,omitempty"`
	ProductComponentVersionUID *string   `json:"productComponentVersionUID,omitempty" xml:"productComponentVersionUID,omitempty"`
	Provider                   *string   `json:"provider,omitempty" xml:"provider,omitempty"`
	Readme                     *string   `json:"readme,omitempty" xml:"readme,omitempty"`
	Resources                  *string   `json:"resources,omitempty" xml:"resources,omitempty"`
	Uid                        *string   `json:"uid,omitempty" xml:"uid,omitempty"`
	Version                    *string   `json:"version,omitempty" xml:"version,omitempty"`
}

func (s GetComponentVersionChildrenResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetComponentVersionChildrenResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetComponentVersionChildrenResponseBodyData) SetComponentName(v string) *GetComponentVersionChildrenResponseBodyData {
	s.ComponentName = &v
	return s
}

func (s *GetComponentVersionChildrenResponseBodyData) SetComponentUID(v string) *GetComponentVersionChildrenResponseBodyData {
	s.ComponentUID = &v
	return s
}

func (s *GetComponentVersionChildrenResponseBodyData) SetDescription(v string) *GetComponentVersionChildrenResponseBodyData {
	s.Description = &v
	return s
}

func (s *GetComponentVersionChildrenResponseBodyData) SetDocuments(v []*string) *GetComponentVersionChildrenResponseBodyData {
	s.Documents = v
	return s
}

func (s *GetComponentVersionChildrenResponseBodyData) SetOrchestrationValues(v string) *GetComponentVersionChildrenResponseBodyData {
	s.OrchestrationValues = &v
	return s
}

func (s *GetComponentVersionChildrenResponseBodyData) SetPackageURL(v string) *GetComponentVersionChildrenResponseBodyData {
	s.PackageURL = &v
	return s
}

func (s *GetComponentVersionChildrenResponseBodyData) SetParentComponent(v bool) *GetComponentVersionChildrenResponseBodyData {
	s.ParentComponent = &v
	return s
}

func (s *GetComponentVersionChildrenResponseBodyData) SetProductComponentVersionUID(v string) *GetComponentVersionChildrenResponseBodyData {
	s.ProductComponentVersionUID = &v
	return s
}

func (s *GetComponentVersionChildrenResponseBodyData) SetProvider(v string) *GetComponentVersionChildrenResponseBodyData {
	s.Provider = &v
	return s
}

func (s *GetComponentVersionChildrenResponseBodyData) SetReadme(v string) *GetComponentVersionChildrenResponseBodyData {
	s.Readme = &v
	return s
}

func (s *GetComponentVersionChildrenResponseBodyData) SetResources(v string) *GetComponentVersionChildrenResponseBodyData {
	s.Resources = &v
	return s
}

func (s *GetComponentVersionChildrenResponseBodyData) SetUid(v string) *GetComponentVersionChildrenResponseBodyData {
	s.Uid = &v
	return s
}

func (s *GetComponentVersionChildrenResponseBodyData) SetVersion(v string) *GetComponentVersionChildrenResponseBodyData {
	s.Version = &v
	return s
}

type GetComponentVersionChildrenResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetComponentVersionChildrenResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetComponentVersionChildrenResponse) String() string {
	return tea.Prettify(s)
}

func (s GetComponentVersionChildrenResponse) GoString() string {
	return s.String()
}

func (s *GetComponentVersionChildrenResponse) SetHeaders(v map[string]*string) *GetComponentVersionChildrenResponse {
	s.Headers = v
	return s
}

func (s *GetComponentVersionChildrenResponse) SetBody(v *GetComponentVersionChildrenResponseBody) *GetComponentVersionChildrenResponse {
	s.Body = v
	return s
}

type GetProductEnvironmentResponseBody struct {
	Data    *GetProductEnvironmentResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	ErrCode *string                                `json:"errCode,omitempty" xml:"errCode,omitempty"`
	ErrMsg  *string                                `json:"errMsg,omitempty" xml:"errMsg,omitempty"`
	Success *bool                                  `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetProductEnvironmentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetProductEnvironmentResponseBody) GoString() string {
	return s.String()
}

func (s *GetProductEnvironmentResponseBody) SetData(v *GetProductEnvironmentResponseBodyData) *GetProductEnvironmentResponseBody {
	s.Data = v
	return s
}

func (s *GetProductEnvironmentResponseBody) SetErrCode(v string) *GetProductEnvironmentResponseBody {
	s.ErrCode = &v
	return s
}

func (s *GetProductEnvironmentResponseBody) SetErrMsg(v string) *GetProductEnvironmentResponseBody {
	s.ErrMsg = &v
	return s
}

func (s *GetProductEnvironmentResponseBody) SetSuccess(v bool) *GetProductEnvironmentResponseBody {
	s.Success = &v
	return s
}

type GetProductEnvironmentResponseBodyData struct {
	CreatedAt      *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	ClusterUID     *string `json:"clusterUID,omitempty" xml:"clusterUID,omitempty"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty"`
	InstanceList   *string `json:"instanceList,omitempty" xml:"instanceList,omitempty"`
	InstanceStatus *string `json:"instanceStatus,omitempty" xml:"instanceStatus,omitempty"`
	Name           *string `json:"name,omitempty" xml:"name,omitempty"`
	ProductName    *string `json:"productName,omitempty" xml:"productName,omitempty"`
	ProductVersion *string `json:"productVersion,omitempty" xml:"productVersion,omitempty"`
	Type           *string `json:"type,omitempty" xml:"type,omitempty"`
	Uid            *string `json:"uid,omitempty" xml:"uid,omitempty"`
	VendorConfig   *string `json:"vendorConfig,omitempty" xml:"vendorConfig,omitempty"`
	VendorType     *string `json:"vendorType,omitempty" xml:"vendorType,omitempty"`
}

func (s GetProductEnvironmentResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetProductEnvironmentResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetProductEnvironmentResponseBodyData) SetCreatedAt(v string) *GetProductEnvironmentResponseBodyData {
	s.CreatedAt = &v
	return s
}

func (s *GetProductEnvironmentResponseBodyData) SetClusterUID(v string) *GetProductEnvironmentResponseBodyData {
	s.ClusterUID = &v
	return s
}

func (s *GetProductEnvironmentResponseBodyData) SetDescription(v string) *GetProductEnvironmentResponseBodyData {
	s.Description = &v
	return s
}

func (s *GetProductEnvironmentResponseBodyData) SetInstanceList(v string) *GetProductEnvironmentResponseBodyData {
	s.InstanceList = &v
	return s
}

func (s *GetProductEnvironmentResponseBodyData) SetInstanceStatus(v string) *GetProductEnvironmentResponseBodyData {
	s.InstanceStatus = &v
	return s
}

func (s *GetProductEnvironmentResponseBodyData) SetName(v string) *GetProductEnvironmentResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetProductEnvironmentResponseBodyData) SetProductName(v string) *GetProductEnvironmentResponseBodyData {
	s.ProductName = &v
	return s
}

func (s *GetProductEnvironmentResponseBodyData) SetProductVersion(v string) *GetProductEnvironmentResponseBodyData {
	s.ProductVersion = &v
	return s
}

func (s *GetProductEnvironmentResponseBodyData) SetType(v string) *GetProductEnvironmentResponseBodyData {
	s.Type = &v
	return s
}

func (s *GetProductEnvironmentResponseBodyData) SetUid(v string) *GetProductEnvironmentResponseBodyData {
	s.Uid = &v
	return s
}

func (s *GetProductEnvironmentResponseBodyData) SetVendorConfig(v string) *GetProductEnvironmentResponseBodyData {
	s.VendorConfig = &v
	return s
}

func (s *GetProductEnvironmentResponseBodyData) SetVendorType(v string) *GetProductEnvironmentResponseBodyData {
	s.VendorType = &v
	return s
}

type GetProductEnvironmentResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetProductEnvironmentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetProductEnvironmentResponse) String() string {
	return tea.Prettify(s)
}

func (s GetProductEnvironmentResponse) GoString() string {
	return s.String()
}

func (s *GetProductEnvironmentResponse) SetHeaders(v map[string]*string) *GetProductEnvironmentResponse {
	s.Headers = v
	return s
}

func (s *GetProductEnvironmentResponse) SetBody(v *GetProductEnvironmentResponseBody) *GetProductEnvironmentResponse {
	s.Body = v
	return s
}

type GetProductVersionPackageRequest struct {
	Platform             *string `json:"platform,omitempty" xml:"platform,omitempty"`
	PackageType          *string `json:"packageType,omitempty" xml:"packageType,omitempty"`
	OldProductVersionUID *string `json:"oldProductVersionUID,omitempty" xml:"oldProductVersionUID,omitempty"`
}

func (s GetProductVersionPackageRequest) String() string {
	return tea.Prettify(s)
}

func (s GetProductVersionPackageRequest) GoString() string {
	return s.String()
}

func (s *GetProductVersionPackageRequest) SetPlatform(v string) *GetProductVersionPackageRequest {
	s.Platform = &v
	return s
}

func (s *GetProductVersionPackageRequest) SetPackageType(v string) *GetProductVersionPackageRequest {
	s.PackageType = &v
	return s
}

func (s *GetProductVersionPackageRequest) SetOldProductVersionUID(v string) *GetProductVersionPackageRequest {
	s.OldProductVersionUID = &v
	return s
}

type GetProductVersionPackageResponseBody struct {
	Data    *GetProductVersionPackageResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	ErrCode *string                                   `json:"errCode,omitempty" xml:"errCode,omitempty"`
	ErrMsg  *string                                   `json:"errMsg,omitempty" xml:"errMsg,omitempty"`
	Success *bool                                     `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetProductVersionPackageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetProductVersionPackageResponseBody) GoString() string {
	return s.String()
}

func (s *GetProductVersionPackageResponseBody) SetData(v *GetProductVersionPackageResponseBodyData) *GetProductVersionPackageResponseBody {
	s.Data = v
	return s
}

func (s *GetProductVersionPackageResponseBody) SetErrCode(v string) *GetProductVersionPackageResponseBody {
	s.ErrCode = &v
	return s
}

func (s *GetProductVersionPackageResponseBody) SetErrMsg(v string) *GetProductVersionPackageResponseBody {
	s.ErrMsg = &v
	return s
}

func (s *GetProductVersionPackageResponseBody) SetSuccess(v bool) *GetProductVersionPackageResponseBody {
	s.Success = &v
	return s
}

type GetProductVersionPackageResponseBodyData struct {
	PackageURL *string `json:"packageURL,omitempty" xml:"packageURL,omitempty"`
}

func (s GetProductVersionPackageResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetProductVersionPackageResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetProductVersionPackageResponseBodyData) SetPackageURL(v string) *GetProductVersionPackageResponseBodyData {
	s.PackageURL = &v
	return s
}

type GetProductVersionPackageResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetProductVersionPackageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetProductVersionPackageResponse) String() string {
	return tea.Prettify(s)
}

func (s GetProductVersionPackageResponse) GoString() string {
	return s.String()
}

func (s *GetProductVersionPackageResponse) SetHeaders(v map[string]*string) *GetProductVersionPackageResponse {
	s.Headers = v
	return s
}

func (s *GetProductVersionPackageResponse) SetBody(v *GetProductVersionPackageResponseBody) *GetProductVersionPackageResponse {
	s.Body = v
	return s
}

type ListAlicloudRegionResponseBody struct {
	Data    *ListAlicloudRegionResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	Success *bool                               `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ListAlicloudRegionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAlicloudRegionResponseBody) GoString() string {
	return s.String()
}

func (s *ListAlicloudRegionResponseBody) SetData(v *ListAlicloudRegionResponseBodyData) *ListAlicloudRegionResponseBody {
	s.Data = v
	return s
}

func (s *ListAlicloudRegionResponseBody) SetSuccess(v bool) *ListAlicloudRegionResponseBody {
	s.Success = &v
	return s
}

type ListAlicloudRegionResponseBodyData struct {
	Regions   *ListAlicloudRegionResponseBodyDataRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Struct"`
	RequestId *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListAlicloudRegionResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListAlicloudRegionResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListAlicloudRegionResponseBodyData) SetRegions(v *ListAlicloudRegionResponseBodyDataRegions) *ListAlicloudRegionResponseBodyData {
	s.Regions = v
	return s
}

func (s *ListAlicloudRegionResponseBodyData) SetRequestId(v string) *ListAlicloudRegionResponseBodyData {
	s.RequestId = &v
	return s
}

type ListAlicloudRegionResponseBodyDataRegions struct {
	Region []*ListAlicloudRegionResponseBodyDataRegionsRegion `json:"Region,omitempty" xml:"Region,omitempty" type:"Repeated"`
}

func (s ListAlicloudRegionResponseBodyDataRegions) String() string {
	return tea.Prettify(s)
}

func (s ListAlicloudRegionResponseBodyDataRegions) GoString() string {
	return s.String()
}

func (s *ListAlicloudRegionResponseBodyDataRegions) SetRegion(v []*ListAlicloudRegionResponseBodyDataRegionsRegion) *ListAlicloudRegionResponseBodyDataRegions {
	s.Region = v
	return s
}

type ListAlicloudRegionResponseBodyDataRegionsRegion struct {
	LocalName      *string `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	RegionEndpoint *string `json:"RegionEndpoint,omitempty" xml:"RegionEndpoint,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Status         *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListAlicloudRegionResponseBodyDataRegionsRegion) String() string {
	return tea.Prettify(s)
}

func (s ListAlicloudRegionResponseBodyDataRegionsRegion) GoString() string {
	return s.String()
}

func (s *ListAlicloudRegionResponseBodyDataRegionsRegion) SetLocalName(v string) *ListAlicloudRegionResponseBodyDataRegionsRegion {
	s.LocalName = &v
	return s
}

func (s *ListAlicloudRegionResponseBodyDataRegionsRegion) SetRegionEndpoint(v string) *ListAlicloudRegionResponseBodyDataRegionsRegion {
	s.RegionEndpoint = &v
	return s
}

func (s *ListAlicloudRegionResponseBodyDataRegionsRegion) SetRegionId(v string) *ListAlicloudRegionResponseBodyDataRegionsRegion {
	s.RegionId = &v
	return s
}

func (s *ListAlicloudRegionResponseBodyDataRegionsRegion) SetStatus(v string) *ListAlicloudRegionResponseBodyDataRegionsRegion {
	s.Status = &v
	return s
}

type ListAlicloudRegionResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListAlicloudRegionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListAlicloudRegionResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAlicloudRegionResponse) GoString() string {
	return s.String()
}

func (s *ListAlicloudRegionResponse) SetHeaders(v map[string]*string) *ListAlicloudRegionResponse {
	s.Headers = v
	return s
}

func (s *ListAlicloudRegionResponse) SetBody(v *ListAlicloudRegionResponseBody) *ListAlicloudRegionResponse {
	s.Body = v
	return s
}

type ListComponentVersionsRequest struct {
	PageNum   *int32                                   `json:"pageNum,omitempty" xml:"pageNum,omitempty"`
	PageSize  *int32                                   `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	Platforms []*ListComponentVersionsRequestPlatforms `json:"platforms,omitempty" xml:"platforms,omitempty" type:"Repeated"`
}

func (s ListComponentVersionsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListComponentVersionsRequest) GoString() string {
	return s.String()
}

func (s *ListComponentVersionsRequest) SetPageNum(v int32) *ListComponentVersionsRequest {
	s.PageNum = &v
	return s
}

func (s *ListComponentVersionsRequest) SetPageSize(v int32) *ListComponentVersionsRequest {
	s.PageSize = &v
	return s
}

func (s *ListComponentVersionsRequest) SetPlatforms(v []*ListComponentVersionsRequestPlatforms) *ListComponentVersionsRequest {
	s.Platforms = v
	return s
}

type ListComponentVersionsRequestPlatforms struct {
	Architecture *string `json:"architecture,omitempty" xml:"architecture,omitempty"`
	Os           *string `json:"os,omitempty" xml:"os,omitempty"`
}

func (s ListComponentVersionsRequestPlatforms) String() string {
	return tea.Prettify(s)
}

func (s ListComponentVersionsRequestPlatforms) GoString() string {
	return s.String()
}

func (s *ListComponentVersionsRequestPlatforms) SetArchitecture(v string) *ListComponentVersionsRequestPlatforms {
	s.Architecture = &v
	return s
}

func (s *ListComponentVersionsRequestPlatforms) SetOs(v string) *ListComponentVersionsRequestPlatforms {
	s.Os = &v
	return s
}

type ListComponentVersionsShrinkRequest struct {
	PageNum         *int32  `json:"pageNum,omitempty" xml:"pageNum,omitempty"`
	PageSize        *int32  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	PlatformsShrink *string `json:"platforms,omitempty" xml:"platforms,omitempty"`
}

func (s ListComponentVersionsShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListComponentVersionsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListComponentVersionsShrinkRequest) SetPageNum(v int32) *ListComponentVersionsShrinkRequest {
	s.PageNum = &v
	return s
}

func (s *ListComponentVersionsShrinkRequest) SetPageSize(v int32) *ListComponentVersionsShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListComponentVersionsShrinkRequest) SetPlatformsShrink(v string) *ListComponentVersionsShrinkRequest {
	s.PlatformsShrink = &v
	return s
}

type ListComponentVersionsResponseBody struct {
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *ListComponentVersionsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	ErrCode   *string                                `json:"errCode,omitempty" xml:"errCode,omitempty"`
	ErrMsg    *string                                `json:"errMsg,omitempty" xml:"errMsg,omitempty"`
	Success   *bool                                  `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ListComponentVersionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListComponentVersionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListComponentVersionsResponseBody) SetRequestId(v string) *ListComponentVersionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListComponentVersionsResponseBody) SetData(v *ListComponentVersionsResponseBodyData) *ListComponentVersionsResponseBody {
	s.Data = v
	return s
}

func (s *ListComponentVersionsResponseBody) SetErrCode(v string) *ListComponentVersionsResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ListComponentVersionsResponseBody) SetErrMsg(v string) *ListComponentVersionsResponseBody {
	s.ErrMsg = &v
	return s
}

func (s *ListComponentVersionsResponseBody) SetSuccess(v bool) *ListComponentVersionsResponseBody {
	s.Success = &v
	return s
}

type ListComponentVersionsResponseBodyData struct {
	List     []*ListComponentVersionsResponseBodyDataList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	PageNum  *int32                                       `json:"pageNum,omitempty" xml:"pageNum,omitempty"`
	PageSize *int32                                       `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	Total    *int32                                       `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListComponentVersionsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListComponentVersionsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListComponentVersionsResponseBodyData) SetList(v []*ListComponentVersionsResponseBodyDataList) *ListComponentVersionsResponseBodyData {
	s.List = v
	return s
}

func (s *ListComponentVersionsResponseBodyData) SetPageNum(v int32) *ListComponentVersionsResponseBodyData {
	s.PageNum = &v
	return s
}

func (s *ListComponentVersionsResponseBodyData) SetPageSize(v int32) *ListComponentVersionsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListComponentVersionsResponseBodyData) SetTotal(v int32) *ListComponentVersionsResponseBodyData {
	s.Total = &v
	return s
}

type ListComponentVersionsResponseBodyDataList struct {
	AppVersion          *string `json:"appVersion,omitempty" xml:"appVersion,omitempty"`
	ComponentName       *string `json:"componentName,omitempty" xml:"componentName,omitempty"`
	ComponentUID        *string `json:"componentUID,omitempty" xml:"componentUID,omitempty"`
	Description         *string `json:"description,omitempty" xml:"description,omitempty"`
	Documents           *string `json:"documents,omitempty" xml:"documents,omitempty"`
	ImagesMapping       *string `json:"imagesMapping,omitempty" xml:"imagesMapping,omitempty"`
	OrchestrationValues *string `json:"orchestrationValues,omitempty" xml:"orchestrationValues,omitempty"`
	PackageURL          *string `json:"packageURL,omitempty" xml:"packageURL,omitempty"`
	ParentComponent     *bool   `json:"parentComponent,omitempty" xml:"parentComponent,omitempty"`
	Provider            *string `json:"provider,omitempty" xml:"provider,omitempty"`
	Readme              *string `json:"readme,omitempty" xml:"readme,omitempty"`
	Resources           *string `json:"resources,omitempty" xml:"resources,omitempty"`
	Uid                 *string `json:"uid,omitempty" xml:"uid,omitempty"`
	Version             *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s ListComponentVersionsResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s ListComponentVersionsResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *ListComponentVersionsResponseBodyDataList) SetAppVersion(v string) *ListComponentVersionsResponseBodyDataList {
	s.AppVersion = &v
	return s
}

func (s *ListComponentVersionsResponseBodyDataList) SetComponentName(v string) *ListComponentVersionsResponseBodyDataList {
	s.ComponentName = &v
	return s
}

func (s *ListComponentVersionsResponseBodyDataList) SetComponentUID(v string) *ListComponentVersionsResponseBodyDataList {
	s.ComponentUID = &v
	return s
}

func (s *ListComponentVersionsResponseBodyDataList) SetDescription(v string) *ListComponentVersionsResponseBodyDataList {
	s.Description = &v
	return s
}

func (s *ListComponentVersionsResponseBodyDataList) SetDocuments(v string) *ListComponentVersionsResponseBodyDataList {
	s.Documents = &v
	return s
}

func (s *ListComponentVersionsResponseBodyDataList) SetImagesMapping(v string) *ListComponentVersionsResponseBodyDataList {
	s.ImagesMapping = &v
	return s
}

func (s *ListComponentVersionsResponseBodyDataList) SetOrchestrationValues(v string) *ListComponentVersionsResponseBodyDataList {
	s.OrchestrationValues = &v
	return s
}

func (s *ListComponentVersionsResponseBodyDataList) SetPackageURL(v string) *ListComponentVersionsResponseBodyDataList {
	s.PackageURL = &v
	return s
}

func (s *ListComponentVersionsResponseBodyDataList) SetParentComponent(v bool) *ListComponentVersionsResponseBodyDataList {
	s.ParentComponent = &v
	return s
}

func (s *ListComponentVersionsResponseBodyDataList) SetProvider(v string) *ListComponentVersionsResponseBodyDataList {
	s.Provider = &v
	return s
}

func (s *ListComponentVersionsResponseBodyDataList) SetReadme(v string) *ListComponentVersionsResponseBodyDataList {
	s.Readme = &v
	return s
}

func (s *ListComponentVersionsResponseBodyDataList) SetResources(v string) *ListComponentVersionsResponseBodyDataList {
	s.Resources = &v
	return s
}

func (s *ListComponentVersionsResponseBodyDataList) SetUid(v string) *ListComponentVersionsResponseBodyDataList {
	s.Uid = &v
	return s
}

func (s *ListComponentVersionsResponseBodyDataList) SetVersion(v string) *ListComponentVersionsResponseBodyDataList {
	s.Version = &v
	return s
}

type ListComponentVersionsResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListComponentVersionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListComponentVersionsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListComponentVersionsResponse) GoString() string {
	return s.String()
}

func (s *ListComponentVersionsResponse) SetHeaders(v map[string]*string) *ListComponentVersionsResponse {
	s.Headers = v
	return s
}

func (s *ListComponentVersionsResponse) SetBody(v *ListComponentVersionsResponseBody) *ListComponentVersionsResponse {
	s.Body = v
	return s
}

type UpdateSnapshotInstanceJoinOptionWithBatchRequest struct {
	InstanceUIDs *string `json:"instanceUIDs,omitempty" xml:"instanceUIDs,omitempty"`
	JoinSnapshot *bool   `json:"joinSnapshot,omitempty" xml:"joinSnapshot,omitempty"`
	RootPassword *string `json:"rootPassword,omitempty" xml:"rootPassword,omitempty"`
}

func (s UpdateSnapshotInstanceJoinOptionWithBatchRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateSnapshotInstanceJoinOptionWithBatchRequest) GoString() string {
	return s.String()
}

func (s *UpdateSnapshotInstanceJoinOptionWithBatchRequest) SetInstanceUIDs(v string) *UpdateSnapshotInstanceJoinOptionWithBatchRequest {
	s.InstanceUIDs = &v
	return s
}

func (s *UpdateSnapshotInstanceJoinOptionWithBatchRequest) SetJoinSnapshot(v bool) *UpdateSnapshotInstanceJoinOptionWithBatchRequest {
	s.JoinSnapshot = &v
	return s
}

func (s *UpdateSnapshotInstanceJoinOptionWithBatchRequest) SetRootPassword(v string) *UpdateSnapshotInstanceJoinOptionWithBatchRequest {
	s.RootPassword = &v
	return s
}

type UpdateSnapshotInstanceJoinOptionWithBatchResponseBody struct {
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	ErrMsg  *string `json:"errMsg,omitempty" xml:"errMsg,omitempty"`
	Success *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateSnapshotInstanceJoinOptionWithBatchResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateSnapshotInstanceJoinOptionWithBatchResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSnapshotInstanceJoinOptionWithBatchResponseBody) SetErrCode(v string) *UpdateSnapshotInstanceJoinOptionWithBatchResponseBody {
	s.ErrCode = &v
	return s
}

func (s *UpdateSnapshotInstanceJoinOptionWithBatchResponseBody) SetErrMsg(v string) *UpdateSnapshotInstanceJoinOptionWithBatchResponseBody {
	s.ErrMsg = &v
	return s
}

func (s *UpdateSnapshotInstanceJoinOptionWithBatchResponseBody) SetSuccess(v bool) *UpdateSnapshotInstanceJoinOptionWithBatchResponseBody {
	s.Success = &v
	return s
}

type UpdateSnapshotInstanceJoinOptionWithBatchResponse struct {
	Headers map[string]*string                                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateSnapshotInstanceJoinOptionWithBatchResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateSnapshotInstanceJoinOptionWithBatchResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateSnapshotInstanceJoinOptionWithBatchResponse) GoString() string {
	return s.String()
}

func (s *UpdateSnapshotInstanceJoinOptionWithBatchResponse) SetHeaders(v map[string]*string) *UpdateSnapshotInstanceJoinOptionWithBatchResponse {
	s.Headers = v
	return s
}

func (s *UpdateSnapshotInstanceJoinOptionWithBatchResponse) SetBody(v *UpdateSnapshotInstanceJoinOptionWithBatchResponseBody) *UpdateSnapshotInstanceJoinOptionWithBatchResponse {
	s.Body = v
	return s
}

type GenerateVendorConfigTemplateResponseBody struct {
	Data    *GenerateVendorConfigTemplateResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	ErrCode *string                                       `json:"errCode,omitempty" xml:"errCode,omitempty"`
	ErrMsg  *string                                       `json:"errMsg,omitempty" xml:"errMsg,omitempty"`
	Success *bool                                         `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GenerateVendorConfigTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GenerateVendorConfigTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateVendorConfigTemplateResponseBody) SetData(v *GenerateVendorConfigTemplateResponseBodyData) *GenerateVendorConfigTemplateResponseBody {
	s.Data = v
	return s
}

func (s *GenerateVendorConfigTemplateResponseBody) SetErrCode(v string) *GenerateVendorConfigTemplateResponseBody {
	s.ErrCode = &v
	return s
}

func (s *GenerateVendorConfigTemplateResponseBody) SetErrMsg(v string) *GenerateVendorConfigTemplateResponseBody {
	s.ErrMsg = &v
	return s
}

func (s *GenerateVendorConfigTemplateResponseBody) SetSuccess(v bool) *GenerateVendorConfigTemplateResponseBody {
	s.Success = &v
	return s
}

type GenerateVendorConfigTemplateResponseBodyData struct {
	VendorConfig *string `json:"vendorConfig,omitempty" xml:"vendorConfig,omitempty"`
}

func (s GenerateVendorConfigTemplateResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GenerateVendorConfigTemplateResponseBodyData) GoString() string {
	return s.String()
}

func (s *GenerateVendorConfigTemplateResponseBodyData) SetVendorConfig(v string) *GenerateVendorConfigTemplateResponseBodyData {
	s.VendorConfig = &v
	return s
}

type GenerateVendorConfigTemplateResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GenerateVendorConfigTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GenerateVendorConfigTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s GenerateVendorConfigTemplateResponse) GoString() string {
	return s.String()
}

func (s *GenerateVendorConfigTemplateResponse) SetHeaders(v map[string]*string) *GenerateVendorConfigTemplateResponse {
	s.Headers = v
	return s
}

func (s *GenerateVendorConfigTemplateResponse) SetBody(v *GenerateVendorConfigTemplateResponseBody) *GenerateVendorConfigTemplateResponse {
	s.Body = v
	return s
}

type UpdateProductComponentRequest struct {
	ComponentOrchestrationValues *string `json:"componentOrchestrationValues,omitempty" xml:"componentOrchestrationValues,omitempty"`
	Enable                       *bool   `json:"enable,omitempty" xml:"enable,omitempty"`
	ReleaseName                  *string `json:"releaseName,omitempty" xml:"releaseName,omitempty"`
}

func (s UpdateProductComponentRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateProductComponentRequest) GoString() string {
	return s.String()
}

func (s *UpdateProductComponentRequest) SetComponentOrchestrationValues(v string) *UpdateProductComponentRequest {
	s.ComponentOrchestrationValues = &v
	return s
}

func (s *UpdateProductComponentRequest) SetEnable(v bool) *UpdateProductComponentRequest {
	s.Enable = &v
	return s
}

func (s *UpdateProductComponentRequest) SetReleaseName(v string) *UpdateProductComponentRequest {
	s.ReleaseName = &v
	return s
}

type UpdateProductComponentResponseBody struct {
	ErrMsg  *string `json:"errMsg,omitempty" xml:"errMsg,omitempty"`
	Success *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateProductComponentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateProductComponentResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateProductComponentResponseBody) SetErrMsg(v string) *UpdateProductComponentResponseBody {
	s.ErrMsg = &v
	return s
}

func (s *UpdateProductComponentResponseBody) SetSuccess(v bool) *UpdateProductComponentResponseBody {
	s.Success = &v
	return s
}

type UpdateProductComponentResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateProductComponentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateProductComponentResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateProductComponentResponse) GoString() string {
	return s.String()
}

func (s *UpdateProductComponentResponse) SetHeaders(v map[string]*string) *UpdateProductComponentResponse {
	s.Headers = v
	return s
}

func (s *UpdateProductComponentResponse) SetBody(v *UpdateProductComponentResponseBody) *UpdateProductComponentResponse {
	s.Body = v
	return s
}

type UpdateEnvironmentNodesRequest struct {
	Cpu          *int32                                      `json:"cpu,omitempty" xml:"cpu,omitempty"`
	DataDisk     []*string                                   `json:"dataDisk,omitempty" xml:"dataDisk,omitempty" type:"Repeated"`
	DataDisk2    []*UpdateEnvironmentNodesRequestDataDisk2   `json:"dataDisk2,omitempty" xml:"dataDisk2,omitempty" type:"Repeated"`
	EnvUID       *string                                     `json:"envUID,omitempty" xml:"envUID,omitempty"`
	Identifier   *string                                     `json:"identifier,omitempty" xml:"identifier,omitempty"`
	Labels       map[string]interface{}                      `json:"labels,omitempty" xml:"labels,omitempty"`
	Memory       *int32                                      `json:"memory,omitempty" xml:"memory,omitempty"`
	NodeIP       *string                                     `json:"nodeIP,omitempty" xml:"nodeIP,omitempty"`
	RootPassword *string                                     `json:"rootPassword,omitempty" xml:"rootPassword,omitempty"`
	SystemDisk   []*string                                   `json:"systemDisk,omitempty" xml:"systemDisk,omitempty" type:"Repeated"`
	SystemDisk2  []*UpdateEnvironmentNodesRequestSystemDisk2 `json:"systemDisk2,omitempty" xml:"systemDisk2,omitempty" type:"Repeated"`
	Taints       []*UpdateEnvironmentNodesRequestTaints      `json:"taints,omitempty" xml:"taints,omitempty" type:"Repeated"`
}

func (s UpdateEnvironmentNodesRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateEnvironmentNodesRequest) GoString() string {
	return s.String()
}

func (s *UpdateEnvironmentNodesRequest) SetCpu(v int32) *UpdateEnvironmentNodesRequest {
	s.Cpu = &v
	return s
}

func (s *UpdateEnvironmentNodesRequest) SetDataDisk(v []*string) *UpdateEnvironmentNodesRequest {
	s.DataDisk = v
	return s
}

func (s *UpdateEnvironmentNodesRequest) SetDataDisk2(v []*UpdateEnvironmentNodesRequestDataDisk2) *UpdateEnvironmentNodesRequest {
	s.DataDisk2 = v
	return s
}

func (s *UpdateEnvironmentNodesRequest) SetEnvUID(v string) *UpdateEnvironmentNodesRequest {
	s.EnvUID = &v
	return s
}

func (s *UpdateEnvironmentNodesRequest) SetIdentifier(v string) *UpdateEnvironmentNodesRequest {
	s.Identifier = &v
	return s
}

func (s *UpdateEnvironmentNodesRequest) SetLabels(v map[string]interface{}) *UpdateEnvironmentNodesRequest {
	s.Labels = v
	return s
}

func (s *UpdateEnvironmentNodesRequest) SetMemory(v int32) *UpdateEnvironmentNodesRequest {
	s.Memory = &v
	return s
}

func (s *UpdateEnvironmentNodesRequest) SetNodeIP(v string) *UpdateEnvironmentNodesRequest {
	s.NodeIP = &v
	return s
}

func (s *UpdateEnvironmentNodesRequest) SetRootPassword(v string) *UpdateEnvironmentNodesRequest {
	s.RootPassword = &v
	return s
}

func (s *UpdateEnvironmentNodesRequest) SetSystemDisk(v []*string) *UpdateEnvironmentNodesRequest {
	s.SystemDisk = v
	return s
}

func (s *UpdateEnvironmentNodesRequest) SetSystemDisk2(v []*UpdateEnvironmentNodesRequestSystemDisk2) *UpdateEnvironmentNodesRequest {
	s.SystemDisk2 = v
	return s
}

func (s *UpdateEnvironmentNodesRequest) SetTaints(v []*UpdateEnvironmentNodesRequestTaints) *UpdateEnvironmentNodesRequest {
	s.Taints = v
	return s
}

type UpdateEnvironmentNodesRequestDataDisk2 struct {
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	Size *string `json:"size,omitempty" xml:"size,omitempty"`
}

func (s UpdateEnvironmentNodesRequestDataDisk2) String() string {
	return tea.Prettify(s)
}

func (s UpdateEnvironmentNodesRequestDataDisk2) GoString() string {
	return s.String()
}

func (s *UpdateEnvironmentNodesRequestDataDisk2) SetName(v string) *UpdateEnvironmentNodesRequestDataDisk2 {
	s.Name = &v
	return s
}

func (s *UpdateEnvironmentNodesRequestDataDisk2) SetSize(v string) *UpdateEnvironmentNodesRequestDataDisk2 {
	s.Size = &v
	return s
}

type UpdateEnvironmentNodesRequestSystemDisk2 struct {
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	Size *string `json:"size,omitempty" xml:"size,omitempty"`
}

func (s UpdateEnvironmentNodesRequestSystemDisk2) String() string {
	return tea.Prettify(s)
}

func (s UpdateEnvironmentNodesRequestSystemDisk2) GoString() string {
	return s.String()
}

func (s *UpdateEnvironmentNodesRequestSystemDisk2) SetName(v string) *UpdateEnvironmentNodesRequestSystemDisk2 {
	s.Name = &v
	return s
}

func (s *UpdateEnvironmentNodesRequestSystemDisk2) SetSize(v string) *UpdateEnvironmentNodesRequestSystemDisk2 {
	s.Size = &v
	return s
}

type UpdateEnvironmentNodesRequestTaints struct {
	Effect *string `json:"effect,omitempty" xml:"effect,omitempty"`
	Key    *string `json:"key,omitempty" xml:"key,omitempty"`
	Value  *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s UpdateEnvironmentNodesRequestTaints) String() string {
	return tea.Prettify(s)
}

func (s UpdateEnvironmentNodesRequestTaints) GoString() string {
	return s.String()
}

func (s *UpdateEnvironmentNodesRequestTaints) SetEffect(v string) *UpdateEnvironmentNodesRequestTaints {
	s.Effect = &v
	return s
}

func (s *UpdateEnvironmentNodesRequestTaints) SetKey(v string) *UpdateEnvironmentNodesRequestTaints {
	s.Key = &v
	return s
}

func (s *UpdateEnvironmentNodesRequestTaints) SetValue(v string) *UpdateEnvironmentNodesRequestTaints {
	s.Value = &v
	return s
}

type UpdateEnvironmentNodesResponseBody struct {
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	ErrMsg  *string `json:"errMsg,omitempty" xml:"errMsg,omitempty"`
	Success *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateEnvironmentNodesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateEnvironmentNodesResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateEnvironmentNodesResponseBody) SetErrCode(v string) *UpdateEnvironmentNodesResponseBody {
	s.ErrCode = &v
	return s
}

func (s *UpdateEnvironmentNodesResponseBody) SetErrMsg(v string) *UpdateEnvironmentNodesResponseBody {
	s.ErrMsg = &v
	return s
}

func (s *UpdateEnvironmentNodesResponseBody) SetSuccess(v bool) *UpdateEnvironmentNodesResponseBody {
	s.Success = &v
	return s
}

type UpdateEnvironmentNodesResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateEnvironmentNodesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateEnvironmentNodesResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateEnvironmentNodesResponse) GoString() string {
	return s.String()
}

func (s *UpdateEnvironmentNodesResponse) SetHeaders(v map[string]*string) *UpdateEnvironmentNodesResponse {
	s.Headers = v
	return s
}

func (s *UpdateEnvironmentNodesResponse) SetBody(v *UpdateEnvironmentNodesResponseBody) *UpdateEnvironmentNodesResponse {
	s.Body = v
	return s
}

type ListEnvironmentPackagesRequest struct {
	PageNum  *int32 `json:"pageNum,omitempty" xml:"pageNum,omitempty"`
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s ListEnvironmentPackagesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListEnvironmentPackagesRequest) GoString() string {
	return s.String()
}

func (s *ListEnvironmentPackagesRequest) SetPageNum(v int32) *ListEnvironmentPackagesRequest {
	s.PageNum = &v
	return s
}

func (s *ListEnvironmentPackagesRequest) SetPageSize(v int32) *ListEnvironmentPackagesRequest {
	s.PageSize = &v
	return s
}

type ListEnvironmentPackagesResponseBody struct {
	Data    *ListEnvironmentPackagesResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	ErrCode *string                                  `json:"errCode,omitempty" xml:"errCode,omitempty"`
	ErrMsg  *string                                  `json:"errMsg,omitempty" xml:"errMsg,omitempty"`
	Success *bool                                    `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ListEnvironmentPackagesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListEnvironmentPackagesResponseBody) GoString() string {
	return s.String()
}

func (s *ListEnvironmentPackagesResponseBody) SetData(v *ListEnvironmentPackagesResponseBodyData) *ListEnvironmentPackagesResponseBody {
	s.Data = v
	return s
}

func (s *ListEnvironmentPackagesResponseBody) SetErrCode(v string) *ListEnvironmentPackagesResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ListEnvironmentPackagesResponseBody) SetErrMsg(v string) *ListEnvironmentPackagesResponseBody {
	s.ErrMsg = &v
	return s
}

func (s *ListEnvironmentPackagesResponseBody) SetSuccess(v bool) *ListEnvironmentPackagesResponseBody {
	s.Success = &v
	return s
}

type ListEnvironmentPackagesResponseBodyData struct {
	List     []*ListEnvironmentPackagesResponseBodyDataList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	PageNum  *int32                                         `json:"pageNum,omitempty" xml:"pageNum,omitempty"`
	PageSize *int32                                         `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	Total    *int32                                         `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListEnvironmentPackagesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListEnvironmentPackagesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListEnvironmentPackagesResponseBodyData) SetList(v []*ListEnvironmentPackagesResponseBodyDataList) *ListEnvironmentPackagesResponseBodyData {
	s.List = v
	return s
}

func (s *ListEnvironmentPackagesResponseBodyData) SetPageNum(v int32) *ListEnvironmentPackagesResponseBodyData {
	s.PageNum = &v
	return s
}

func (s *ListEnvironmentPackagesResponseBodyData) SetPageSize(v int32) *ListEnvironmentPackagesResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListEnvironmentPackagesResponseBodyData) SetTotal(v int32) *ListEnvironmentPackagesResponseBodyData {
	s.Total = &v
	return s
}

type ListEnvironmentPackagesResponseBodyDataList struct {
	EnvUID   *string `json:"envUID,omitempty" xml:"envUID,omitempty"`
	Provider *string `json:"provider,omitempty" xml:"provider,omitempty"`
	Status   *string `json:"status,omitempty" xml:"status,omitempty"`
	Uid      *string `json:"uid,omitempty" xml:"uid,omitempty"`
	Url      *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s ListEnvironmentPackagesResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s ListEnvironmentPackagesResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *ListEnvironmentPackagesResponseBodyDataList) SetEnvUID(v string) *ListEnvironmentPackagesResponseBodyDataList {
	s.EnvUID = &v
	return s
}

func (s *ListEnvironmentPackagesResponseBodyDataList) SetProvider(v string) *ListEnvironmentPackagesResponseBodyDataList {
	s.Provider = &v
	return s
}

func (s *ListEnvironmentPackagesResponseBodyDataList) SetStatus(v string) *ListEnvironmentPackagesResponseBodyDataList {
	s.Status = &v
	return s
}

func (s *ListEnvironmentPackagesResponseBodyDataList) SetUid(v string) *ListEnvironmentPackagesResponseBodyDataList {
	s.Uid = &v
	return s
}

func (s *ListEnvironmentPackagesResponseBodyDataList) SetUrl(v string) *ListEnvironmentPackagesResponseBodyDataList {
	s.Url = &v
	return s
}

type ListEnvironmentPackagesResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListEnvironmentPackagesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListEnvironmentPackagesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListEnvironmentPackagesResponse) GoString() string {
	return s.String()
}

func (s *ListEnvironmentPackagesResponse) SetHeaders(v map[string]*string) *ListEnvironmentPackagesResponse {
	s.Headers = v
	return s
}

func (s *ListEnvironmentPackagesResponse) SetBody(v *ListEnvironmentPackagesResponseBody) *ListEnvironmentPackagesResponse {
	s.Body = v
	return s
}

type CreateEnvironmentHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	ClientToken   *string            `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s CreateEnvironmentHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateEnvironmentHeaders) GoString() string {
	return s.String()
}

func (s *CreateEnvironmentHeaders) SetCommonHeaders(v map[string]*string) *CreateEnvironmentHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateEnvironmentHeaders) SetClientToken(v string) *CreateEnvironmentHeaders {
	s.ClientToken = &v
	return s
}

type CreateEnvironmentRequest struct {
	Annotations *string                           `json:"annotations,omitempty" xml:"annotations,omitempty"`
	Description *string                           `json:"description,omitempty" xml:"description,omitempty"`
	Location    *string                           `json:"location,omitempty" xml:"location,omitempty"`
	Name        *string                           `json:"name,omitempty" xml:"name,omitempty"`
	Platform    *CreateEnvironmentRequestPlatform `json:"platform,omitempty" xml:"platform,omitempty" type:"Struct"`
	VendorType  *string                           `json:"vendorType,omitempty" xml:"vendorType,omitempty"`
}

func (s CreateEnvironmentRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateEnvironmentRequest) GoString() string {
	return s.String()
}

func (s *CreateEnvironmentRequest) SetAnnotations(v string) *CreateEnvironmentRequest {
	s.Annotations = &v
	return s
}

func (s *CreateEnvironmentRequest) SetDescription(v string) *CreateEnvironmentRequest {
	s.Description = &v
	return s
}

func (s *CreateEnvironmentRequest) SetLocation(v string) *CreateEnvironmentRequest {
	s.Location = &v
	return s
}

func (s *CreateEnvironmentRequest) SetName(v string) *CreateEnvironmentRequest {
	s.Name = &v
	return s
}

func (s *CreateEnvironmentRequest) SetPlatform(v *CreateEnvironmentRequestPlatform) *CreateEnvironmentRequest {
	s.Platform = v
	return s
}

func (s *CreateEnvironmentRequest) SetVendorType(v string) *CreateEnvironmentRequest {
	s.VendorType = &v
	return s
}

type CreateEnvironmentRequestPlatform struct {
	Architecture *string `json:"architecture,omitempty" xml:"architecture,omitempty"`
	Os           *string `json:"os,omitempty" xml:"os,omitempty"`
}

func (s CreateEnvironmentRequestPlatform) String() string {
	return tea.Prettify(s)
}

func (s CreateEnvironmentRequestPlatform) GoString() string {
	return s.String()
}

func (s *CreateEnvironmentRequestPlatform) SetArchitecture(v string) *CreateEnvironmentRequestPlatform {
	s.Architecture = &v
	return s
}

func (s *CreateEnvironmentRequestPlatform) SetOs(v string) *CreateEnvironmentRequestPlatform {
	s.Os = &v
	return s
}

type CreateEnvironmentResponseBody struct {
	Data    *CreateEnvironmentResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	ErrCode *string                            `json:"errCode,omitempty" xml:"errCode,omitempty"`
	ErrMsg  *string                            `json:"errMsg,omitempty" xml:"errMsg,omitempty"`
	Success *bool                              `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreateEnvironmentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateEnvironmentResponseBody) GoString() string {
	return s.String()
}

func (s *CreateEnvironmentResponseBody) SetData(v *CreateEnvironmentResponseBodyData) *CreateEnvironmentResponseBody {
	s.Data = v
	return s
}

func (s *CreateEnvironmentResponseBody) SetErrCode(v string) *CreateEnvironmentResponseBody {
	s.ErrCode = &v
	return s
}

func (s *CreateEnvironmentResponseBody) SetErrMsg(v string) *CreateEnvironmentResponseBody {
	s.ErrMsg = &v
	return s
}

func (s *CreateEnvironmentResponseBody) SetSuccess(v bool) *CreateEnvironmentResponseBody {
	s.Success = &v
	return s
}

type CreateEnvironmentResponseBodyData struct {
	Uid *string `json:"uid,omitempty" xml:"uid,omitempty"`
}

func (s CreateEnvironmentResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateEnvironmentResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateEnvironmentResponseBodyData) SetUid(v string) *CreateEnvironmentResponseBodyData {
	s.Uid = &v
	return s
}

type CreateEnvironmentResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateEnvironmentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateEnvironmentResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateEnvironmentResponse) GoString() string {
	return s.String()
}

func (s *CreateEnvironmentResponse) SetHeaders(v map[string]*string) *CreateEnvironmentResponse {
	s.Headers = v
	return s
}

func (s *CreateEnvironmentResponse) SetBody(v *CreateEnvironmentResponseBody) *CreateEnvironmentResponse {
	s.Body = v
	return s
}

type GetEnvironmentLogResponseBody struct {
	Data    *GetEnvironmentLogResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	ErrCode *string                            `json:"errCode,omitempty" xml:"errCode,omitempty"`
	ErrMsg  *string                            `json:"errMsg,omitempty" xml:"errMsg,omitempty"`
	Success *bool                              `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetEnvironmentLogResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetEnvironmentLogResponseBody) GoString() string {
	return s.String()
}

func (s *GetEnvironmentLogResponseBody) SetData(v *GetEnvironmentLogResponseBodyData) *GetEnvironmentLogResponseBody {
	s.Data = v
	return s
}

func (s *GetEnvironmentLogResponseBody) SetErrCode(v string) *GetEnvironmentLogResponseBody {
	s.ErrCode = &v
	return s
}

func (s *GetEnvironmentLogResponseBody) SetErrMsg(v string) *GetEnvironmentLogResponseBody {
	s.ErrMsg = &v
	return s
}

func (s *GetEnvironmentLogResponseBody) SetSuccess(v bool) *GetEnvironmentLogResponseBody {
	s.Success = &v
	return s
}

type GetEnvironmentLogResponseBodyData struct {
	End     *bool   `json:"end,omitempty" xml:"end,omitempty"`
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
}

func (s GetEnvironmentLogResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetEnvironmentLogResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetEnvironmentLogResponseBodyData) SetEnd(v bool) *GetEnvironmentLogResponseBodyData {
	s.End = &v
	return s
}

func (s *GetEnvironmentLogResponseBodyData) SetMessage(v string) *GetEnvironmentLogResponseBodyData {
	s.Message = &v
	return s
}

type GetEnvironmentLogResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetEnvironmentLogResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetEnvironmentLogResponse) String() string {
	return tea.Prettify(s)
}

func (s GetEnvironmentLogResponse) GoString() string {
	return s.String()
}

func (s *GetEnvironmentLogResponse) SetHeaders(v map[string]*string) *GetEnvironmentLogResponse {
	s.Headers = v
	return s
}

func (s *GetEnvironmentLogResponse) SetBody(v *GetEnvironmentLogResponseBody) *GetEnvironmentLogResponse {
	s.Body = v
	return s
}

type ListEnvironmentNodeRequest struct {
	PageNum  *int32  `json:"pageNum,omitempty" xml:"pageNum,omitempty"`
	PageSize *int32  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	Name     *string `json:"name,omitempty" xml:"name,omitempty"`
	Fuzzy    *string `json:"fuzzy,omitempty" xml:"fuzzy,omitempty"`
	NodeIp   *string `json:"nodeIp,omitempty" xml:"nodeIp,omitempty"`
}

func (s ListEnvironmentNodeRequest) String() string {
	return tea.Prettify(s)
}

func (s ListEnvironmentNodeRequest) GoString() string {
	return s.String()
}

func (s *ListEnvironmentNodeRequest) SetPageNum(v int32) *ListEnvironmentNodeRequest {
	s.PageNum = &v
	return s
}

func (s *ListEnvironmentNodeRequest) SetPageSize(v int32) *ListEnvironmentNodeRequest {
	s.PageSize = &v
	return s
}

func (s *ListEnvironmentNodeRequest) SetName(v string) *ListEnvironmentNodeRequest {
	s.Name = &v
	return s
}

func (s *ListEnvironmentNodeRequest) SetFuzzy(v string) *ListEnvironmentNodeRequest {
	s.Fuzzy = &v
	return s
}

func (s *ListEnvironmentNodeRequest) SetNodeIp(v string) *ListEnvironmentNodeRequest {
	s.NodeIp = &v
	return s
}

type ListEnvironmentNodeResponseBody struct {
	Data    *ListEnvironmentNodeResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	ErrCode *string                              `json:"errCode,omitempty" xml:"errCode,omitempty"`
	ErrMsg  *string                              `json:"errMsg,omitempty" xml:"errMsg,omitempty"`
	Success *bool                                `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ListEnvironmentNodeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListEnvironmentNodeResponseBody) GoString() string {
	return s.String()
}

func (s *ListEnvironmentNodeResponseBody) SetData(v *ListEnvironmentNodeResponseBodyData) *ListEnvironmentNodeResponseBody {
	s.Data = v
	return s
}

func (s *ListEnvironmentNodeResponseBody) SetErrCode(v string) *ListEnvironmentNodeResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ListEnvironmentNodeResponseBody) SetErrMsg(v string) *ListEnvironmentNodeResponseBody {
	s.ErrMsg = &v
	return s
}

func (s *ListEnvironmentNodeResponseBody) SetSuccess(v bool) *ListEnvironmentNodeResponseBody {
	s.Success = &v
	return s
}

type ListEnvironmentNodeResponseBodyData struct {
	List     []*ListEnvironmentNodeResponseBodyDataList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	PageNum  *int32                                     `json:"pageNum,omitempty" xml:"pageNum,omitempty"`
	PageSize *int32                                     `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	Total    *int32                                     `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListEnvironmentNodeResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListEnvironmentNodeResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListEnvironmentNodeResponseBodyData) SetList(v []*ListEnvironmentNodeResponseBodyDataList) *ListEnvironmentNodeResponseBodyData {
	s.List = v
	return s
}

func (s *ListEnvironmentNodeResponseBodyData) SetPageNum(v int32) *ListEnvironmentNodeResponseBodyData {
	s.PageNum = &v
	return s
}

func (s *ListEnvironmentNodeResponseBodyData) SetPageSize(v int32) *ListEnvironmentNodeResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListEnvironmentNodeResponseBodyData) SetTotal(v int32) *ListEnvironmentNodeResponseBodyData {
	s.Total = &v
	return s
}

type ListEnvironmentNodeResponseBodyDataList struct {
	CreatedAt         *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	Description       *string `json:"description,omitempty" xml:"description,omitempty"`
	Id                *int32  `json:"id,omitempty" xml:"id,omitempty"`
	Name              *string `json:"name,omitempty" xml:"name,omitempty"`
	ProductName       *string `json:"productName,omitempty" xml:"productName,omitempty"`
	ProductVersion    *string `json:"productVersion,omitempty" xml:"productVersion,omitempty"`
	ProductVersionUID *string `json:"productVersionUID,omitempty" xml:"productVersionUID,omitempty"`
	Uid               *string `json:"uid,omitempty" xml:"uid,omitempty"`
	VendorType        *string `json:"vendorType,omitempty" xml:"vendorType,omitempty"`
}

func (s ListEnvironmentNodeResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s ListEnvironmentNodeResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *ListEnvironmentNodeResponseBodyDataList) SetCreatedAt(v string) *ListEnvironmentNodeResponseBodyDataList {
	s.CreatedAt = &v
	return s
}

func (s *ListEnvironmentNodeResponseBodyDataList) SetDescription(v string) *ListEnvironmentNodeResponseBodyDataList {
	s.Description = &v
	return s
}

func (s *ListEnvironmentNodeResponseBodyDataList) SetId(v int32) *ListEnvironmentNodeResponseBodyDataList {
	s.Id = &v
	return s
}

func (s *ListEnvironmentNodeResponseBodyDataList) SetName(v string) *ListEnvironmentNodeResponseBodyDataList {
	s.Name = &v
	return s
}

func (s *ListEnvironmentNodeResponseBodyDataList) SetProductName(v string) *ListEnvironmentNodeResponseBodyDataList {
	s.ProductName = &v
	return s
}

func (s *ListEnvironmentNodeResponseBodyDataList) SetProductVersion(v string) *ListEnvironmentNodeResponseBodyDataList {
	s.ProductVersion = &v
	return s
}

func (s *ListEnvironmentNodeResponseBodyDataList) SetProductVersionUID(v string) *ListEnvironmentNodeResponseBodyDataList {
	s.ProductVersionUID = &v
	return s
}

func (s *ListEnvironmentNodeResponseBodyDataList) SetUid(v string) *ListEnvironmentNodeResponseBodyDataList {
	s.Uid = &v
	return s
}

func (s *ListEnvironmentNodeResponseBodyDataList) SetVendorType(v string) *ListEnvironmentNodeResponseBodyDataList {
	s.VendorType = &v
	return s
}

type ListEnvironmentNodeResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListEnvironmentNodeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListEnvironmentNodeResponse) String() string {
	return tea.Prettify(s)
}

func (s ListEnvironmentNodeResponse) GoString() string {
	return s.String()
}

func (s *ListEnvironmentNodeResponse) SetHeaders(v map[string]*string) *ListEnvironmentNodeResponse {
	s.Headers = v
	return s
}

func (s *ListEnvironmentNodeResponse) SetBody(v *ListEnvironmentNodeResponseBody) *ListEnvironmentNodeResponse {
	s.Body = v
	return s
}

type ListProductVersionRelatedFoundationComponentVersionsResponseBody struct {
	Data    []*ProductComponentRelationDetail `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	ErrCode *string                           `json:"errCode,omitempty" xml:"errCode,omitempty"`
	ErrMsg  *string                           `json:"errMsg,omitempty" xml:"errMsg,omitempty"`
	Success *bool                             `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ListProductVersionRelatedFoundationComponentVersionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListProductVersionRelatedFoundationComponentVersionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListProductVersionRelatedFoundationComponentVersionsResponseBody) SetData(v []*ProductComponentRelationDetail) *ListProductVersionRelatedFoundationComponentVersionsResponseBody {
	s.Data = v
	return s
}

func (s *ListProductVersionRelatedFoundationComponentVersionsResponseBody) SetErrCode(v string) *ListProductVersionRelatedFoundationComponentVersionsResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ListProductVersionRelatedFoundationComponentVersionsResponseBody) SetErrMsg(v string) *ListProductVersionRelatedFoundationComponentVersionsResponseBody {
	s.ErrMsg = &v
	return s
}

func (s *ListProductVersionRelatedFoundationComponentVersionsResponseBody) SetSuccess(v bool) *ListProductVersionRelatedFoundationComponentVersionsResponseBody {
	s.Success = &v
	return s
}

type ListProductVersionRelatedFoundationComponentVersionsResponse struct {
	Headers map[string]*string                                                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListProductVersionRelatedFoundationComponentVersionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListProductVersionRelatedFoundationComponentVersionsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListProductVersionRelatedFoundationComponentVersionsResponse) GoString() string {
	return s.String()
}

func (s *ListProductVersionRelatedFoundationComponentVersionsResponse) SetHeaders(v map[string]*string) *ListProductVersionRelatedFoundationComponentVersionsResponse {
	s.Headers = v
	return s
}

func (s *ListProductVersionRelatedFoundationComponentVersionsResponse) SetBody(v *ListProductVersionRelatedFoundationComponentVersionsResponseBody) *ListProductVersionRelatedFoundationComponentVersionsResponse {
	s.Body = v
	return s
}

type SyncComponentRequest struct {
	Region     *string `json:"region,omitempty" xml:"region,omitempty"`
	BucketName *string `json:"bucketName,omitempty" xml:"bucketName,omitempty"`
}

func (s SyncComponentRequest) String() string {
	return tea.Prettify(s)
}

func (s SyncComponentRequest) GoString() string {
	return s.String()
}

func (s *SyncComponentRequest) SetRegion(v string) *SyncComponentRequest {
	s.Region = &v
	return s
}

func (s *SyncComponentRequest) SetBucketName(v string) *SyncComponentRequest {
	s.BucketName = &v
	return s
}

type SyncComponentResponseBody struct {
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	ErrMsg  *string `json:"errMsg,omitempty" xml:"errMsg,omitempty"`
	Status  *string `json:"status,omitempty" xml:"status,omitempty"`
	Success *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s SyncComponentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SyncComponentResponseBody) GoString() string {
	return s.String()
}

func (s *SyncComponentResponseBody) SetErrCode(v string) *SyncComponentResponseBody {
	s.ErrCode = &v
	return s
}

func (s *SyncComponentResponseBody) SetErrMsg(v string) *SyncComponentResponseBody {
	s.ErrMsg = &v
	return s
}

func (s *SyncComponentResponseBody) SetStatus(v string) *SyncComponentResponseBody {
	s.Status = &v
	return s
}

func (s *SyncComponentResponseBody) SetSuccess(v bool) *SyncComponentResponseBody {
	s.Success = &v
	return s
}

type SyncComponentResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SyncComponentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SyncComponentResponse) String() string {
	return tea.Prettify(s)
}

func (s SyncComponentResponse) GoString() string {
	return s.String()
}

func (s *SyncComponentResponse) SetHeaders(v map[string]*string) *SyncComponentResponse {
	s.Headers = v
	return s
}

func (s *SyncComponentResponse) SetBody(v *SyncComponentResponseBody) *SyncComponentResponse {
	s.Body = v
	return s
}

type UpdateComponentToProductRequest struct {
	// the component Version ID
	ComponentVersionID *string `json:"componentVersionID,omitempty" xml:"componentVersionID,omitempty"`
}

func (s UpdateComponentToProductRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateComponentToProductRequest) GoString() string {
	return s.String()
}

func (s *UpdateComponentToProductRequest) SetComponentVersionID(v string) *UpdateComponentToProductRequest {
	s.ComponentVersionID = &v
	return s
}

type UpdateComponentToProductResponseBody struct {
	// Id of the request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateComponentToProductResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateComponentToProductResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateComponentToProductResponseBody) SetRequestId(v string) *UpdateComponentToProductResponseBody {
	s.RequestId = &v
	return s
}

type UpdateComponentToProductResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateComponentToProductResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateComponentToProductResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateComponentToProductResponse) GoString() string {
	return s.String()
}

func (s *UpdateComponentToProductResponse) SetHeaders(v map[string]*string) *UpdateComponentToProductResponse {
	s.Headers = v
	return s
}

func (s *UpdateComponentToProductResponse) SetBody(v *UpdateComponentToProductResponseBody) *UpdateComponentToProductResponse {
	s.Body = v
	return s
}

type CreateEnvironmentNodeRequest struct {
	Cpu          *int32                                    `json:"cpu,omitempty" xml:"cpu,omitempty"`
	DataDisk     []*CreateEnvironmentNodeRequestDataDisk   `json:"dataDisk,omitempty" xml:"dataDisk,omitempty" type:"Repeated"`
	HostName     *string                                   `json:"hostName,omitempty" xml:"hostName,omitempty"`
	Identifier   *string                                   `json:"identifier,omitempty" xml:"identifier,omitempty"`
	Labels       map[string]interface{}                    `json:"labels,omitempty" xml:"labels,omitempty"`
	Memory       *int32                                    `json:"memory,omitempty" xml:"memory,omitempty"`
	Os           *string                                   `json:"os,omitempty" xml:"os,omitempty"`
	PrivateIP    *string                                   `json:"privateIP,omitempty" xml:"privateIP,omitempty"`
	Provider     *string                                   `json:"provider,omitempty" xml:"provider,omitempty"`
	RootPassword *string                                   `json:"rootPassword,omitempty" xml:"rootPassword,omitempty"`
	SystemDisk   []*CreateEnvironmentNodeRequestSystemDisk `json:"systemDisk,omitempty" xml:"systemDisk,omitempty" type:"Repeated"`
	Taints       []*CreateEnvironmentNodeRequestTaints     `json:"taints,omitempty" xml:"taints,omitempty" type:"Repeated"`
}

func (s CreateEnvironmentNodeRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateEnvironmentNodeRequest) GoString() string {
	return s.String()
}

func (s *CreateEnvironmentNodeRequest) SetCpu(v int32) *CreateEnvironmentNodeRequest {
	s.Cpu = &v
	return s
}

func (s *CreateEnvironmentNodeRequest) SetDataDisk(v []*CreateEnvironmentNodeRequestDataDisk) *CreateEnvironmentNodeRequest {
	s.DataDisk = v
	return s
}

func (s *CreateEnvironmentNodeRequest) SetHostName(v string) *CreateEnvironmentNodeRequest {
	s.HostName = &v
	return s
}

func (s *CreateEnvironmentNodeRequest) SetIdentifier(v string) *CreateEnvironmentNodeRequest {
	s.Identifier = &v
	return s
}

func (s *CreateEnvironmentNodeRequest) SetLabels(v map[string]interface{}) *CreateEnvironmentNodeRequest {
	s.Labels = v
	return s
}

func (s *CreateEnvironmentNodeRequest) SetMemory(v int32) *CreateEnvironmentNodeRequest {
	s.Memory = &v
	return s
}

func (s *CreateEnvironmentNodeRequest) SetOs(v string) *CreateEnvironmentNodeRequest {
	s.Os = &v
	return s
}

func (s *CreateEnvironmentNodeRequest) SetPrivateIP(v string) *CreateEnvironmentNodeRequest {
	s.PrivateIP = &v
	return s
}

func (s *CreateEnvironmentNodeRequest) SetProvider(v string) *CreateEnvironmentNodeRequest {
	s.Provider = &v
	return s
}

func (s *CreateEnvironmentNodeRequest) SetRootPassword(v string) *CreateEnvironmentNodeRequest {
	s.RootPassword = &v
	return s
}

func (s *CreateEnvironmentNodeRequest) SetSystemDisk(v []*CreateEnvironmentNodeRequestSystemDisk) *CreateEnvironmentNodeRequest {
	s.SystemDisk = v
	return s
}

func (s *CreateEnvironmentNodeRequest) SetTaints(v []*CreateEnvironmentNodeRequestTaints) *CreateEnvironmentNodeRequest {
	s.Taints = v
	return s
}

type CreateEnvironmentNodeRequestDataDisk struct {
	Name     *string `json:"name,omitempty" xml:"name,omitempty"`
	Required *int32  `json:"required,omitempty" xml:"required,omitempty"`
}

func (s CreateEnvironmentNodeRequestDataDisk) String() string {
	return tea.Prettify(s)
}

func (s CreateEnvironmentNodeRequestDataDisk) GoString() string {
	return s.String()
}

func (s *CreateEnvironmentNodeRequestDataDisk) SetName(v string) *CreateEnvironmentNodeRequestDataDisk {
	s.Name = &v
	return s
}

func (s *CreateEnvironmentNodeRequestDataDisk) SetRequired(v int32) *CreateEnvironmentNodeRequestDataDisk {
	s.Required = &v
	return s
}

type CreateEnvironmentNodeRequestSystemDisk struct {
	Name     *string `json:"name,omitempty" xml:"name,omitempty"`
	Required *int32  `json:"required,omitempty" xml:"required,omitempty"`
}

func (s CreateEnvironmentNodeRequestSystemDisk) String() string {
	return tea.Prettify(s)
}

func (s CreateEnvironmentNodeRequestSystemDisk) GoString() string {
	return s.String()
}

func (s *CreateEnvironmentNodeRequestSystemDisk) SetName(v string) *CreateEnvironmentNodeRequestSystemDisk {
	s.Name = &v
	return s
}

func (s *CreateEnvironmentNodeRequestSystemDisk) SetRequired(v int32) *CreateEnvironmentNodeRequestSystemDisk {
	s.Required = &v
	return s
}

type CreateEnvironmentNodeRequestTaints struct {
	Effect *string `json:"effect,omitempty" xml:"effect,omitempty"`
	Key    *string `json:"key,omitempty" xml:"key,omitempty"`
	Value  *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s CreateEnvironmentNodeRequestTaints) String() string {
	return tea.Prettify(s)
}

func (s CreateEnvironmentNodeRequestTaints) GoString() string {
	return s.String()
}

func (s *CreateEnvironmentNodeRequestTaints) SetEffect(v string) *CreateEnvironmentNodeRequestTaints {
	s.Effect = &v
	return s
}

func (s *CreateEnvironmentNodeRequestTaints) SetKey(v string) *CreateEnvironmentNodeRequestTaints {
	s.Key = &v
	return s
}

func (s *CreateEnvironmentNodeRequestTaints) SetValue(v string) *CreateEnvironmentNodeRequestTaints {
	s.Value = &v
	return s
}

type CreateEnvironmentNodeResponseBody struct {
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	ErrMsg  *string `json:"errMsg,omitempty" xml:"errMsg,omitempty"`
	Success *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreateEnvironmentNodeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateEnvironmentNodeResponseBody) GoString() string {
	return s.String()
}

func (s *CreateEnvironmentNodeResponseBody) SetErrCode(v string) *CreateEnvironmentNodeResponseBody {
	s.ErrCode = &v
	return s
}

func (s *CreateEnvironmentNodeResponseBody) SetErrMsg(v string) *CreateEnvironmentNodeResponseBody {
	s.ErrMsg = &v
	return s
}

func (s *CreateEnvironmentNodeResponseBody) SetSuccess(v bool) *CreateEnvironmentNodeResponseBody {
	s.Success = &v
	return s
}

type CreateEnvironmentNodeResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateEnvironmentNodeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateEnvironmentNodeResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateEnvironmentNodeResponse) GoString() string {
	return s.String()
}

func (s *CreateEnvironmentNodeResponse) SetHeaders(v map[string]*string) *CreateEnvironmentNodeResponse {
	s.Headers = v
	return s
}

func (s *CreateEnvironmentNodeResponse) SetBody(v *CreateEnvironmentNodeResponseBody) *CreateEnvironmentNodeResponse {
	s.Body = v
	return s
}

type GetComponentResponseBody struct {
	Data    *GetComponentResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	ErrCode *string                       `json:"errCode,omitempty" xml:"errCode,omitempty"`
	ErrMsg  *string                       `json:"errMsg,omitempty" xml:"errMsg,omitempty"`
	Success *bool                         `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetComponentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetComponentResponseBody) GoString() string {
	return s.String()
}

func (s *GetComponentResponseBody) SetData(v *GetComponentResponseBodyData) *GetComponentResponseBody {
	s.Data = v
	return s
}

func (s *GetComponentResponseBody) SetErrCode(v string) *GetComponentResponseBody {
	s.ErrCode = &v
	return s
}

func (s *GetComponentResponseBody) SetErrMsg(v string) *GetComponentResponseBody {
	s.ErrMsg = &v
	return s
}

func (s *GetComponentResponseBody) SetSuccess(v bool) *GetComponentResponseBody {
	s.Success = &v
	return s
}

type GetComponentResponseBodyData struct {
	Annotations *GetComponentResponseBodyDataAnnotations `json:"annotations,omitempty" xml:"annotations,omitempty" type:"Struct"`
	Category    *string                                  `json:"category,omitempty" xml:"category,omitempty"`
	Description *string                                  `json:"description,omitempty" xml:"description,omitempty"`
	Documents   *string                                  `json:"documents,omitempty" xml:"documents,omitempty"`
	Name        *string                                  `json:"name,omitempty" xml:"name,omitempty"`
	Provider    *string                                  `json:"provider,omitempty" xml:"provider,omitempty"`
	Public      *bool                                    `json:"public,omitempty" xml:"public,omitempty"`
	Singleton   *bool                                    `json:"singleton,omitempty" xml:"singleton,omitempty"`
	Source      *string                                  `json:"source,omitempty" xml:"source,omitempty"`
	Uid         *string                                  `json:"uid,omitempty" xml:"uid,omitempty"`
}

func (s GetComponentResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetComponentResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetComponentResponseBodyData) SetAnnotations(v *GetComponentResponseBodyDataAnnotations) *GetComponentResponseBodyData {
	s.Annotations = v
	return s
}

func (s *GetComponentResponseBodyData) SetCategory(v string) *GetComponentResponseBodyData {
	s.Category = &v
	return s
}

func (s *GetComponentResponseBodyData) SetDescription(v string) *GetComponentResponseBodyData {
	s.Description = &v
	return s
}

func (s *GetComponentResponseBodyData) SetDocuments(v string) *GetComponentResponseBodyData {
	s.Documents = &v
	return s
}

func (s *GetComponentResponseBodyData) SetName(v string) *GetComponentResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetComponentResponseBodyData) SetProvider(v string) *GetComponentResponseBodyData {
	s.Provider = &v
	return s
}

func (s *GetComponentResponseBodyData) SetPublic(v bool) *GetComponentResponseBodyData {
	s.Public = &v
	return s
}

func (s *GetComponentResponseBodyData) SetSingleton(v bool) *GetComponentResponseBodyData {
	s.Singleton = &v
	return s
}

func (s *GetComponentResponseBodyData) SetSource(v string) *GetComponentResponseBodyData {
	s.Source = &v
	return s
}

func (s *GetComponentResponseBodyData) SetUid(v string) *GetComponentResponseBodyData {
	s.Uid = &v
	return s
}

type GetComponentResponseBodyDataAnnotations struct {
	AdditionalProp1 *string `json:"additionalProp1,omitempty" xml:"additionalProp1,omitempty"`
	AdditionalProp2 *string `json:"additionalProp2,omitempty" xml:"additionalProp2,omitempty"`
	AdditionalProp3 *string `json:"additionalProp3,omitempty" xml:"additionalProp3,omitempty"`
}

func (s GetComponentResponseBodyDataAnnotations) String() string {
	return tea.Prettify(s)
}

func (s GetComponentResponseBodyDataAnnotations) GoString() string {
	return s.String()
}

func (s *GetComponentResponseBodyDataAnnotations) SetAdditionalProp1(v string) *GetComponentResponseBodyDataAnnotations {
	s.AdditionalProp1 = &v
	return s
}

func (s *GetComponentResponseBodyDataAnnotations) SetAdditionalProp2(v string) *GetComponentResponseBodyDataAnnotations {
	s.AdditionalProp2 = &v
	return s
}

func (s *GetComponentResponseBodyDataAnnotations) SetAdditionalProp3(v string) *GetComponentResponseBodyDataAnnotations {
	s.AdditionalProp3 = &v
	return s
}

type GetComponentResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetComponentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetComponentResponse) String() string {
	return tea.Prettify(s)
}

func (s GetComponentResponse) GoString() string {
	return s.String()
}

func (s *GetComponentResponse) SetHeaders(v map[string]*string) *GetComponentResponse {
	s.Headers = v
	return s
}

func (s *GetComponentResponse) SetBody(v *GetComponentResponseBody) *GetComponentResponse {
	s.Body = v
	return s
}

type ListFoundationVersionRelatedComponentVersionsResponseBody struct {
	RequestId *string                                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *ListFoundationVersionRelatedComponentVersionsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	ErrCode   *string                                                        `json:"errCode,omitempty" xml:"errCode,omitempty"`
	ErrMsg    *string                                                        `json:"errMsg,omitempty" xml:"errMsg,omitempty"`
	Success   *bool                                                          `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ListFoundationVersionRelatedComponentVersionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListFoundationVersionRelatedComponentVersionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListFoundationVersionRelatedComponentVersionsResponseBody) SetRequestId(v string) *ListFoundationVersionRelatedComponentVersionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListFoundationVersionRelatedComponentVersionsResponseBody) SetData(v *ListFoundationVersionRelatedComponentVersionsResponseBodyData) *ListFoundationVersionRelatedComponentVersionsResponseBody {
	s.Data = v
	return s
}

func (s *ListFoundationVersionRelatedComponentVersionsResponseBody) SetErrCode(v string) *ListFoundationVersionRelatedComponentVersionsResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ListFoundationVersionRelatedComponentVersionsResponseBody) SetErrMsg(v string) *ListFoundationVersionRelatedComponentVersionsResponseBody {
	s.ErrMsg = &v
	return s
}

func (s *ListFoundationVersionRelatedComponentVersionsResponseBody) SetSuccess(v bool) *ListFoundationVersionRelatedComponentVersionsResponseBody {
	s.Success = &v
	return s
}

type ListFoundationVersionRelatedComponentVersionsResponseBodyData struct {
	List     []*ComponentVersion `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	PageNum  *int32              `json:"pageNum,omitempty" xml:"pageNum,omitempty"`
	PageSize *int32              `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	Total    *int32              `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListFoundationVersionRelatedComponentVersionsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListFoundationVersionRelatedComponentVersionsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListFoundationVersionRelatedComponentVersionsResponseBodyData) SetList(v []*ComponentVersion) *ListFoundationVersionRelatedComponentVersionsResponseBodyData {
	s.List = v
	return s
}

func (s *ListFoundationVersionRelatedComponentVersionsResponseBodyData) SetPageNum(v int32) *ListFoundationVersionRelatedComponentVersionsResponseBodyData {
	s.PageNum = &v
	return s
}

func (s *ListFoundationVersionRelatedComponentVersionsResponseBodyData) SetPageSize(v int32) *ListFoundationVersionRelatedComponentVersionsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListFoundationVersionRelatedComponentVersionsResponseBodyData) SetTotal(v int32) *ListFoundationVersionRelatedComponentVersionsResponseBodyData {
	s.Total = &v
	return s
}

type ListFoundationVersionRelatedComponentVersionsResponse struct {
	Headers map[string]*string                                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListFoundationVersionRelatedComponentVersionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListFoundationVersionRelatedComponentVersionsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListFoundationVersionRelatedComponentVersionsResponse) GoString() string {
	return s.String()
}

func (s *ListFoundationVersionRelatedComponentVersionsResponse) SetHeaders(v map[string]*string) *ListFoundationVersionRelatedComponentVersionsResponse {
	s.Headers = v
	return s
}

func (s *ListFoundationVersionRelatedComponentVersionsResponse) SetBody(v *ListFoundationVersionRelatedComponentVersionsResponseBody) *ListFoundationVersionRelatedComponentVersionsResponse {
	s.Body = v
	return s
}

type GetSnapshotResponseBody struct {
	Data    *GetSnapshotResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	ErrCode *string                      `json:"errCode,omitempty" xml:"errCode,omitempty"`
	ErrMsg  *string                      `json:"errMsg,omitempty" xml:"errMsg,omitempty"`
	Success *bool                        `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetSnapshotResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSnapshotResponseBody) GoString() string {
	return s.String()
}

func (s *GetSnapshotResponseBody) SetData(v *GetSnapshotResponseBodyData) *GetSnapshotResponseBody {
	s.Data = v
	return s
}

func (s *GetSnapshotResponseBody) SetErrCode(v string) *GetSnapshotResponseBody {
	s.ErrCode = &v
	return s
}

func (s *GetSnapshotResponseBody) SetErrMsg(v string) *GetSnapshotResponseBody {
	s.ErrMsg = &v
	return s
}

func (s *GetSnapshotResponseBody) SetSuccess(v bool) *GetSnapshotResponseBody {
	s.Success = &v
	return s
}

type GetSnapshotResponseBodyData struct {
	Description          *string `json:"description,omitempty" xml:"description,omitempty"`
	InstanceCIDR         *string `json:"instanceCIDR,omitempty" xml:"instanceCIDR,omitempty"`
	Name                 *string `json:"name,omitempty" xml:"name,omitempty"`
	ProductName          *string `json:"productName,omitempty" xml:"productName,omitempty"`
	ProductVersion       *string `json:"productVersion,omitempty" xml:"productVersion,omitempty"`
	ProductVersionDesc   *string `json:"productVersionDesc,omitempty" xml:"productVersionDesc,omitempty"`
	Region               *string `json:"region,omitempty" xml:"region,omitempty"`
	SnapshotRegion       *string `json:"snapshotRegion,omitempty" xml:"snapshotRegion,omitempty"`
	SnapshotStatus       *string `json:"snapshotStatus,omitempty" xml:"snapshotStatus,omitempty"`
	SourceEnvironmentUID *string `json:"sourceEnvironmentUID,omitempty" xml:"sourceEnvironmentUID,omitempty"`
	SourceType           *string `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
	Uid                  *string `json:"uid,omitempty" xml:"uid,omitempty"`
	Vpcid                *string `json:"vpcid,omitempty" xml:"vpcid,omitempty"`
}

func (s GetSnapshotResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetSnapshotResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetSnapshotResponseBodyData) SetDescription(v string) *GetSnapshotResponseBodyData {
	s.Description = &v
	return s
}

func (s *GetSnapshotResponseBodyData) SetInstanceCIDR(v string) *GetSnapshotResponseBodyData {
	s.InstanceCIDR = &v
	return s
}

func (s *GetSnapshotResponseBodyData) SetName(v string) *GetSnapshotResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetSnapshotResponseBodyData) SetProductName(v string) *GetSnapshotResponseBodyData {
	s.ProductName = &v
	return s
}

func (s *GetSnapshotResponseBodyData) SetProductVersion(v string) *GetSnapshotResponseBodyData {
	s.ProductVersion = &v
	return s
}

func (s *GetSnapshotResponseBodyData) SetProductVersionDesc(v string) *GetSnapshotResponseBodyData {
	s.ProductVersionDesc = &v
	return s
}

func (s *GetSnapshotResponseBodyData) SetRegion(v string) *GetSnapshotResponseBodyData {
	s.Region = &v
	return s
}

func (s *GetSnapshotResponseBodyData) SetSnapshotRegion(v string) *GetSnapshotResponseBodyData {
	s.SnapshotRegion = &v
	return s
}

func (s *GetSnapshotResponseBodyData) SetSnapshotStatus(v string) *GetSnapshotResponseBodyData {
	s.SnapshotStatus = &v
	return s
}

func (s *GetSnapshotResponseBodyData) SetSourceEnvironmentUID(v string) *GetSnapshotResponseBodyData {
	s.SourceEnvironmentUID = &v
	return s
}

func (s *GetSnapshotResponseBodyData) SetSourceType(v string) *GetSnapshotResponseBodyData {
	s.SourceType = &v
	return s
}

func (s *GetSnapshotResponseBodyData) SetUid(v string) *GetSnapshotResponseBodyData {
	s.Uid = &v
	return s
}

func (s *GetSnapshotResponseBodyData) SetVpcid(v string) *GetSnapshotResponseBodyData {
	s.Vpcid = &v
	return s
}

type GetSnapshotResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetSnapshotResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetSnapshotResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSnapshotResponse) GoString() string {
	return s.String()
}

func (s *GetSnapshotResponse) SetHeaders(v map[string]*string) *GetSnapshotResponse {
	s.Headers = v
	return s
}

func (s *GetSnapshotResponse) SetBody(v *GetSnapshotResponseBody) *GetSnapshotResponse {
	s.Body = v
	return s
}

type GetLicenseResponseBody struct {
	// Id of the request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetLicenseResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetLicenseResponseBody) GoString() string {
	return s.String()
}

func (s *GetLicenseResponseBody) SetRequestId(v string) *GetLicenseResponseBody {
	s.RequestId = &v
	return s
}

type GetLicenseResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetLicenseResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetLicenseResponse) String() string {
	return tea.Prettify(s)
}

func (s GetLicenseResponse) GoString() string {
	return s.String()
}

func (s *GetLicenseResponse) SetHeaders(v map[string]*string) *GetLicenseResponse {
	s.Headers = v
	return s
}

func (s *GetLicenseResponse) SetBody(v *GetLicenseResponseBody) *GetLicenseResponse {
	s.Body = v
	return s
}

type CreateLatestProductVersionHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	ClientToken   *string            `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s CreateLatestProductVersionHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateLatestProductVersionHeaders) GoString() string {
	return s.String()
}

func (s *CreateLatestProductVersionHeaders) SetCommonHeaders(v map[string]*string) *CreateLatestProductVersionHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateLatestProductVersionHeaders) SetClientToken(v string) *CreateLatestProductVersionHeaders {
	s.ClientToken = &v
	return s
}

type CreateLatestProductVersionResponseBody struct {
	Data    *CreateLatestProductVersionResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	ErrCode *string                                     `json:"errCode,omitempty" xml:"errCode,omitempty"`
	ErrMsg  *string                                     `json:"errMsg,omitempty" xml:"errMsg,omitempty"`
	Success *bool                                       `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreateLatestProductVersionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateLatestProductVersionResponseBody) GoString() string {
	return s.String()
}

func (s *CreateLatestProductVersionResponseBody) SetData(v *CreateLatestProductVersionResponseBodyData) *CreateLatestProductVersionResponseBody {
	s.Data = v
	return s
}

func (s *CreateLatestProductVersionResponseBody) SetErrCode(v string) *CreateLatestProductVersionResponseBody {
	s.ErrCode = &v
	return s
}

func (s *CreateLatestProductVersionResponseBody) SetErrMsg(v string) *CreateLatestProductVersionResponseBody {
	s.ErrMsg = &v
	return s
}

func (s *CreateLatestProductVersionResponseBody) SetSuccess(v bool) *CreateLatestProductVersionResponseBody {
	s.Success = &v
	return s
}

type CreateLatestProductVersionResponseBodyData struct {
	Uid *string `json:"uid,omitempty" xml:"uid,omitempty"`
}

func (s CreateLatestProductVersionResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateLatestProductVersionResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateLatestProductVersionResponseBodyData) SetUid(v string) *CreateLatestProductVersionResponseBodyData {
	s.Uid = &v
	return s
}

type CreateLatestProductVersionResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateLatestProductVersionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateLatestProductVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateLatestProductVersionResponse) GoString() string {
	return s.String()
}

func (s *CreateLatestProductVersionResponse) SetHeaders(v map[string]*string) *CreateLatestProductVersionResponse {
	s.Headers = v
	return s
}

func (s *CreateLatestProductVersionResponse) SetBody(v *CreateLatestProductVersionResponseBody) *CreateLatestProductVersionResponse {
	s.Body = v
	return s
}

type ListAlicloudVPCResponseBody struct {
	Data    *ListAlicloudVPCResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	Success *bool                            `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ListAlicloudVPCResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAlicloudVPCResponseBody) GoString() string {
	return s.String()
}

func (s *ListAlicloudVPCResponseBody) SetData(v *ListAlicloudVPCResponseBodyData) *ListAlicloudVPCResponseBody {
	s.Data = v
	return s
}

func (s *ListAlicloudVPCResponseBody) SetSuccess(v bool) *ListAlicloudVPCResponseBody {
	s.Success = &v
	return s
}

type ListAlicloudVPCResponseBodyData struct {
	PageNumber *int32                               `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                               `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                               `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	Vpcs       *ListAlicloudVPCResponseBodyDataVpcs `json:"Vpcs,omitempty" xml:"Vpcs,omitempty" type:"Struct"`
}

func (s ListAlicloudVPCResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListAlicloudVPCResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListAlicloudVPCResponseBodyData) SetPageNumber(v int32) *ListAlicloudVPCResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListAlicloudVPCResponseBodyData) SetPageSize(v int32) *ListAlicloudVPCResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListAlicloudVPCResponseBodyData) SetRequestId(v string) *ListAlicloudVPCResponseBodyData {
	s.RequestId = &v
	return s
}

func (s *ListAlicloudVPCResponseBodyData) SetTotalCount(v int32) *ListAlicloudVPCResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListAlicloudVPCResponseBodyData) SetVpcs(v *ListAlicloudVPCResponseBodyDataVpcs) *ListAlicloudVPCResponseBodyData {
	s.Vpcs = v
	return s
}

type ListAlicloudVPCResponseBodyDataVpcs struct {
	Vpc []*ListAlicloudVPCResponseBodyDataVpcsVpc `json:"Vpc,omitempty" xml:"Vpc,omitempty" type:"Repeated"`
}

func (s ListAlicloudVPCResponseBodyDataVpcs) String() string {
	return tea.Prettify(s)
}

func (s ListAlicloudVPCResponseBodyDataVpcs) GoString() string {
	return s.String()
}

func (s *ListAlicloudVPCResponseBodyDataVpcs) SetVpc(v []*ListAlicloudVPCResponseBodyDataVpcsVpc) *ListAlicloudVPCResponseBodyDataVpcs {
	s.Vpc = v
	return s
}

type ListAlicloudVPCResponseBodyDataVpcsVpc struct {
	CidrBlock    *string                                           `json:"CidrBlock,omitempty" xml:"CidrBlock,omitempty"`
	CreationTime *string                                           `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	Description  *string                                           `json:"Description,omitempty" xml:"Description,omitempty"`
	IsDefault    *bool                                             `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	RegionId     *string                                           `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Status       *string                                           `json:"Status,omitempty" xml:"Status,omitempty"`
	UserCidrs    *ListAlicloudVPCResponseBodyDataVpcsVpcUserCidrs  `json:"UserCidrs,omitempty" xml:"UserCidrs,omitempty" type:"Struct"`
	VRouterId    *string                                           `json:"VRouterId,omitempty" xml:"VRouterId,omitempty"`
	VSwitchIds   *ListAlicloudVPCResponseBodyDataVpcsVpcVSwitchIds `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty" type:"Struct"`
	VpcId        *string                                           `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	VpcName      *string                                           `json:"VpcName,omitempty" xml:"VpcName,omitempty"`
}

func (s ListAlicloudVPCResponseBodyDataVpcsVpc) String() string {
	return tea.Prettify(s)
}

func (s ListAlicloudVPCResponseBodyDataVpcsVpc) GoString() string {
	return s.String()
}

func (s *ListAlicloudVPCResponseBodyDataVpcsVpc) SetCidrBlock(v string) *ListAlicloudVPCResponseBodyDataVpcsVpc {
	s.CidrBlock = &v
	return s
}

func (s *ListAlicloudVPCResponseBodyDataVpcsVpc) SetCreationTime(v string) *ListAlicloudVPCResponseBodyDataVpcsVpc {
	s.CreationTime = &v
	return s
}

func (s *ListAlicloudVPCResponseBodyDataVpcsVpc) SetDescription(v string) *ListAlicloudVPCResponseBodyDataVpcsVpc {
	s.Description = &v
	return s
}

func (s *ListAlicloudVPCResponseBodyDataVpcsVpc) SetIsDefault(v bool) *ListAlicloudVPCResponseBodyDataVpcsVpc {
	s.IsDefault = &v
	return s
}

func (s *ListAlicloudVPCResponseBodyDataVpcsVpc) SetRegionId(v string) *ListAlicloudVPCResponseBodyDataVpcsVpc {
	s.RegionId = &v
	return s
}

func (s *ListAlicloudVPCResponseBodyDataVpcsVpc) SetStatus(v string) *ListAlicloudVPCResponseBodyDataVpcsVpc {
	s.Status = &v
	return s
}

func (s *ListAlicloudVPCResponseBodyDataVpcsVpc) SetUserCidrs(v *ListAlicloudVPCResponseBodyDataVpcsVpcUserCidrs) *ListAlicloudVPCResponseBodyDataVpcsVpc {
	s.UserCidrs = v
	return s
}

func (s *ListAlicloudVPCResponseBodyDataVpcsVpc) SetVRouterId(v string) *ListAlicloudVPCResponseBodyDataVpcsVpc {
	s.VRouterId = &v
	return s
}

func (s *ListAlicloudVPCResponseBodyDataVpcsVpc) SetVSwitchIds(v *ListAlicloudVPCResponseBodyDataVpcsVpcVSwitchIds) *ListAlicloudVPCResponseBodyDataVpcsVpc {
	s.VSwitchIds = v
	return s
}

func (s *ListAlicloudVPCResponseBodyDataVpcsVpc) SetVpcId(v string) *ListAlicloudVPCResponseBodyDataVpcsVpc {
	s.VpcId = &v
	return s
}

func (s *ListAlicloudVPCResponseBodyDataVpcsVpc) SetVpcName(v string) *ListAlicloudVPCResponseBodyDataVpcsVpc {
	s.VpcName = &v
	return s
}

type ListAlicloudVPCResponseBodyDataVpcsVpcUserCidrs struct {
	UserCidr *string `json:"UserCidr,omitempty" xml:"UserCidr,omitempty"`
}

func (s ListAlicloudVPCResponseBodyDataVpcsVpcUserCidrs) String() string {
	return tea.Prettify(s)
}

func (s ListAlicloudVPCResponseBodyDataVpcsVpcUserCidrs) GoString() string {
	return s.String()
}

func (s *ListAlicloudVPCResponseBodyDataVpcsVpcUserCidrs) SetUserCidr(v string) *ListAlicloudVPCResponseBodyDataVpcsVpcUserCidrs {
	s.UserCidr = &v
	return s
}

type ListAlicloudVPCResponseBodyDataVpcsVpcVSwitchIds struct {
	VSwitchId []*string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty" type:"Repeated"`
}

func (s ListAlicloudVPCResponseBodyDataVpcsVpcVSwitchIds) String() string {
	return tea.Prettify(s)
}

func (s ListAlicloudVPCResponseBodyDataVpcsVpcVSwitchIds) GoString() string {
	return s.String()
}

func (s *ListAlicloudVPCResponseBodyDataVpcsVpcVSwitchIds) SetVSwitchId(v []*string) *ListAlicloudVPCResponseBodyDataVpcsVpcVSwitchIds {
	s.VSwitchId = v
	return s
}

type ListAlicloudVPCResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListAlicloudVPCResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListAlicloudVPCResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAlicloudVPCResponse) GoString() string {
	return s.String()
}

func (s *ListAlicloudVPCResponse) SetHeaders(v map[string]*string) *ListAlicloudVPCResponse {
	s.Headers = v
	return s
}

func (s *ListAlicloudVPCResponse) SetBody(v *ListAlicloudVPCResponseBody) *ListAlicloudVPCResponse {
	s.Body = v
	return s
}

type CreateProductHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	ClientToken   *string            `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s CreateProductHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateProductHeaders) GoString() string {
	return s.String()
}

func (s *CreateProductHeaders) SetCommonHeaders(v map[string]*string) *CreateProductHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateProductHeaders) SetClientToken(v string) *CreateProductHeaders {
	s.ClientToken = &v
	return s
}

type CreateProductRequest struct {
	Annotations          *string `json:"annotations,omitempty" xml:"annotations,omitempty"`
	ComponentVersionUID  *string `json:"componentVersionUID,omitempty" xml:"componentVersionUID,omitempty"`
	Description          *string `json:"description,omitempty" xml:"description,omitempty"`
	FoundationVersionUID *string `json:"foundationVersionUID,omitempty" xml:"foundationVersionUID,omitempty"`
	ProductName          *string `json:"productName,omitempty" xml:"productName,omitempty"`
	PrometheusUID        *string `json:"prometheusUID,omitempty" xml:"prometheusUID,omitempty"`
}

func (s CreateProductRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateProductRequest) GoString() string {
	return s.String()
}

func (s *CreateProductRequest) SetAnnotations(v string) *CreateProductRequest {
	s.Annotations = &v
	return s
}

func (s *CreateProductRequest) SetComponentVersionUID(v string) *CreateProductRequest {
	s.ComponentVersionUID = &v
	return s
}

func (s *CreateProductRequest) SetDescription(v string) *CreateProductRequest {
	s.Description = &v
	return s
}

func (s *CreateProductRequest) SetFoundationVersionUID(v string) *CreateProductRequest {
	s.FoundationVersionUID = &v
	return s
}

func (s *CreateProductRequest) SetProductName(v string) *CreateProductRequest {
	s.ProductName = &v
	return s
}

func (s *CreateProductRequest) SetPrometheusUID(v string) *CreateProductRequest {
	s.PrometheusUID = &v
	return s
}

type CreateProductResponseBody struct {
	Data    *CreateProductResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	ErrCode *string                        `json:"errCode,omitempty" xml:"errCode,omitempty"`
	ErrMsg  *string                        `json:"errMsg,omitempty" xml:"errMsg,omitempty"`
	Success *bool                          `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreateProductResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateProductResponseBody) GoString() string {
	return s.String()
}

func (s *CreateProductResponseBody) SetData(v *CreateProductResponseBodyData) *CreateProductResponseBody {
	s.Data = v
	return s
}

func (s *CreateProductResponseBody) SetErrCode(v string) *CreateProductResponseBody {
	s.ErrCode = &v
	return s
}

func (s *CreateProductResponseBody) SetErrMsg(v string) *CreateProductResponseBody {
	s.ErrMsg = &v
	return s
}

func (s *CreateProductResponseBody) SetSuccess(v bool) *CreateProductResponseBody {
	s.Success = &v
	return s
}

type CreateProductResponseBodyData struct {
	Uid *string `json:"uid,omitempty" xml:"uid,omitempty"`
}

func (s CreateProductResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateProductResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateProductResponseBodyData) SetUid(v string) *CreateProductResponseBodyData {
	s.Uid = &v
	return s
}

type CreateProductResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateProductResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateProductResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateProductResponse) GoString() string {
	return s.String()
}

func (s *CreateProductResponse) SetHeaders(v map[string]*string) *CreateProductResponse {
	s.Headers = v
	return s
}

func (s *CreateProductResponse) SetBody(v *CreateProductResponseBody) *CreateProductResponse {
	s.Body = v
	return s
}

type GetProductEnvironmentsRequest struct {
	ProductUID *string                                   `json:"productUID,omitempty" xml:"productUID,omitempty"`
	EnvType    *string                                   `json:"envType,omitempty" xml:"envType,omitempty"`
	Platforms  []*GetProductEnvironmentsRequestPlatforms `json:"platforms,omitempty" xml:"platforms,omitempty" type:"Repeated"`
}

func (s GetProductEnvironmentsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetProductEnvironmentsRequest) GoString() string {
	return s.String()
}

func (s *GetProductEnvironmentsRequest) SetProductUID(v string) *GetProductEnvironmentsRequest {
	s.ProductUID = &v
	return s
}

func (s *GetProductEnvironmentsRequest) SetEnvType(v string) *GetProductEnvironmentsRequest {
	s.EnvType = &v
	return s
}

func (s *GetProductEnvironmentsRequest) SetPlatforms(v []*GetProductEnvironmentsRequestPlatforms) *GetProductEnvironmentsRequest {
	s.Platforms = v
	return s
}

type GetProductEnvironmentsRequestPlatforms struct {
	Architecture *string `json:"architecture,omitempty" xml:"architecture,omitempty"`
	Os           *string `json:"os,omitempty" xml:"os,omitempty"`
}

func (s GetProductEnvironmentsRequestPlatforms) String() string {
	return tea.Prettify(s)
}

func (s GetProductEnvironmentsRequestPlatforms) GoString() string {
	return s.String()
}

func (s *GetProductEnvironmentsRequestPlatforms) SetArchitecture(v string) *GetProductEnvironmentsRequestPlatforms {
	s.Architecture = &v
	return s
}

func (s *GetProductEnvironmentsRequestPlatforms) SetOs(v string) *GetProductEnvironmentsRequestPlatforms {
	s.Os = &v
	return s
}

type GetProductEnvironmentsShrinkRequest struct {
	ProductUID      *string `json:"productUID,omitempty" xml:"productUID,omitempty"`
	EnvType         *string `json:"envType,omitempty" xml:"envType,omitempty"`
	PlatformsShrink *string `json:"platforms,omitempty" xml:"platforms,omitempty"`
}

func (s GetProductEnvironmentsShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s GetProductEnvironmentsShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetProductEnvironmentsShrinkRequest) SetProductUID(v string) *GetProductEnvironmentsShrinkRequest {
	s.ProductUID = &v
	return s
}

func (s *GetProductEnvironmentsShrinkRequest) SetEnvType(v string) *GetProductEnvironmentsShrinkRequest {
	s.EnvType = &v
	return s
}

func (s *GetProductEnvironmentsShrinkRequest) SetPlatformsShrink(v string) *GetProductEnvironmentsShrinkRequest {
	s.PlatformsShrink = &v
	return s
}

type GetProductEnvironmentsResponseBody struct {
	Data    *GetProductEnvironmentsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	ErrCode *string                                 `json:"errCode,omitempty" xml:"errCode,omitempty"`
	ErrMsg  *string                                 `json:"errMsg,omitempty" xml:"errMsg,omitempty"`
	Success *bool                                   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetProductEnvironmentsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetProductEnvironmentsResponseBody) GoString() string {
	return s.String()
}

func (s *GetProductEnvironmentsResponseBody) SetData(v *GetProductEnvironmentsResponseBodyData) *GetProductEnvironmentsResponseBody {
	s.Data = v
	return s
}

func (s *GetProductEnvironmentsResponseBody) SetErrCode(v string) *GetProductEnvironmentsResponseBody {
	s.ErrCode = &v
	return s
}

func (s *GetProductEnvironmentsResponseBody) SetErrMsg(v string) *GetProductEnvironmentsResponseBody {
	s.ErrMsg = &v
	return s
}

func (s *GetProductEnvironmentsResponseBody) SetSuccess(v bool) *GetProductEnvironmentsResponseBody {
	s.Success = &v
	return s
}

type GetProductEnvironmentsResponseBodyData struct {
	ClusterId      *string `json:"clusterId,omitempty" xml:"clusterId,omitempty"`
	CreatedAt      *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty"`
	InstanceList   *string `json:"instanceList,omitempty" xml:"instanceList,omitempty"`
	Name           *string `json:"name,omitempty" xml:"name,omitempty"`
	ProductName    *string `json:"productName,omitempty" xml:"productName,omitempty"`
	ProductVersion *string `json:"productVersion,omitempty" xml:"productVersion,omitempty"`
	Uid            *string `json:"uid,omitempty" xml:"uid,omitempty"`
	VendorConfig   *string `json:"vendorConfig,omitempty" xml:"vendorConfig,omitempty"`
	VendorType     *string `json:"vendorType,omitempty" xml:"vendorType,omitempty"`
}

func (s GetProductEnvironmentsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetProductEnvironmentsResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetProductEnvironmentsResponseBodyData) SetClusterId(v string) *GetProductEnvironmentsResponseBodyData {
	s.ClusterId = &v
	return s
}

func (s *GetProductEnvironmentsResponseBodyData) SetCreatedAt(v string) *GetProductEnvironmentsResponseBodyData {
	s.CreatedAt = &v
	return s
}

func (s *GetProductEnvironmentsResponseBodyData) SetDescription(v string) *GetProductEnvironmentsResponseBodyData {
	s.Description = &v
	return s
}

func (s *GetProductEnvironmentsResponseBodyData) SetInstanceList(v string) *GetProductEnvironmentsResponseBodyData {
	s.InstanceList = &v
	return s
}

func (s *GetProductEnvironmentsResponseBodyData) SetName(v string) *GetProductEnvironmentsResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetProductEnvironmentsResponseBodyData) SetProductName(v string) *GetProductEnvironmentsResponseBodyData {
	s.ProductName = &v
	return s
}

func (s *GetProductEnvironmentsResponseBodyData) SetProductVersion(v string) *GetProductEnvironmentsResponseBodyData {
	s.ProductVersion = &v
	return s
}

func (s *GetProductEnvironmentsResponseBodyData) SetUid(v string) *GetProductEnvironmentsResponseBodyData {
	s.Uid = &v
	return s
}

func (s *GetProductEnvironmentsResponseBodyData) SetVendorConfig(v string) *GetProductEnvironmentsResponseBodyData {
	s.VendorConfig = &v
	return s
}

func (s *GetProductEnvironmentsResponseBodyData) SetVendorType(v string) *GetProductEnvironmentsResponseBodyData {
	s.VendorType = &v
	return s
}

type GetProductEnvironmentsResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetProductEnvironmentsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetProductEnvironmentsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetProductEnvironmentsResponse) GoString() string {
	return s.String()
}

func (s *GetProductEnvironmentsResponse) SetHeaders(v map[string]*string) *GetProductEnvironmentsResponse {
	s.Headers = v
	return s
}

func (s *GetProductEnvironmentsResponse) SetBody(v *GetProductEnvironmentsResponseBody) *GetProductEnvironmentsResponse {
	s.Body = v
	return s
}

type DeleteComponentResponseBody struct {
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	Success    *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrCode    *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
}

func (s DeleteComponentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteComponentResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteComponentResponseBody) SetErrMessage(v string) *DeleteComponentResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DeleteComponentResponseBody) SetSuccess(v bool) *DeleteComponentResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteComponentResponseBody) SetErrCode(v string) *DeleteComponentResponseBody {
	s.ErrCode = &v
	return s
}

type DeleteComponentResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteComponentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteComponentResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteComponentResponse) GoString() string {
	return s.String()
}

func (s *DeleteComponentResponse) SetHeaders(v map[string]*string) *DeleteComponentResponse {
	s.Headers = v
	return s
}

func (s *DeleteComponentResponse) SetBody(v *DeleteComponentResponseBody) *DeleteComponentResponse {
	s.Body = v
	return s
}

type DeleteProductComponentResponseBody struct {
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	ErrMsg  *string `json:"errMsg,omitempty" xml:"errMsg,omitempty"`
	Status  *string `json:"status,omitempty" xml:"status,omitempty"`
	Success *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeleteProductComponentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteProductComponentResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteProductComponentResponseBody) SetErrCode(v string) *DeleteProductComponentResponseBody {
	s.ErrCode = &v
	return s
}

func (s *DeleteProductComponentResponseBody) SetErrMsg(v string) *DeleteProductComponentResponseBody {
	s.ErrMsg = &v
	return s
}

func (s *DeleteProductComponentResponseBody) SetStatus(v string) *DeleteProductComponentResponseBody {
	s.Status = &v
	return s
}

func (s *DeleteProductComponentResponseBody) SetSuccess(v bool) *DeleteProductComponentResponseBody {
	s.Success = &v
	return s
}

type DeleteProductComponentResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteProductComponentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteProductComponentResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteProductComponentResponse) GoString() string {
	return s.String()
}

func (s *DeleteProductComponentResponse) SetHeaders(v map[string]*string) *DeleteProductComponentResponse {
	s.Headers = v
	return s
}

func (s *DeleteProductComponentResponse) SetBody(v *DeleteProductComponentResponseBody) *DeleteProductComponentResponse {
	s.Body = v
	return s
}

type CreateEnvironmentWithSnapshotRequest struct {
	EnvironmentDesc *string `json:"environmentDesc,omitempty" xml:"environmentDesc,omitempty"`
	EnvironmentName *string `json:"environmentName,omitempty" xml:"environmentName,omitempty"`
}

func (s CreateEnvironmentWithSnapshotRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateEnvironmentWithSnapshotRequest) GoString() string {
	return s.String()
}

func (s *CreateEnvironmentWithSnapshotRequest) SetEnvironmentDesc(v string) *CreateEnvironmentWithSnapshotRequest {
	s.EnvironmentDesc = &v
	return s
}

func (s *CreateEnvironmentWithSnapshotRequest) SetEnvironmentName(v string) *CreateEnvironmentWithSnapshotRequest {
	s.EnvironmentName = &v
	return s
}

type CreateEnvironmentWithSnapshotResponseBody struct {
	Data    *CreateEnvironmentWithSnapshotResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	ErrCode *string                                        `json:"errCode,omitempty" xml:"errCode,omitempty"`
	ErrMsg  *string                                        `json:"errMsg,omitempty" xml:"errMsg,omitempty"`
	Success *bool                                          `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreateEnvironmentWithSnapshotResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateEnvironmentWithSnapshotResponseBody) GoString() string {
	return s.String()
}

func (s *CreateEnvironmentWithSnapshotResponseBody) SetData(v *CreateEnvironmentWithSnapshotResponseBodyData) *CreateEnvironmentWithSnapshotResponseBody {
	s.Data = v
	return s
}

func (s *CreateEnvironmentWithSnapshotResponseBody) SetErrCode(v string) *CreateEnvironmentWithSnapshotResponseBody {
	s.ErrCode = &v
	return s
}

func (s *CreateEnvironmentWithSnapshotResponseBody) SetErrMsg(v string) *CreateEnvironmentWithSnapshotResponseBody {
	s.ErrMsg = &v
	return s
}

func (s *CreateEnvironmentWithSnapshotResponseBody) SetSuccess(v bool) *CreateEnvironmentWithSnapshotResponseBody {
	s.Success = &v
	return s
}

type CreateEnvironmentWithSnapshotResponseBodyData struct {
	Uid *string `json:"uid,omitempty" xml:"uid,omitempty"`
}

func (s CreateEnvironmentWithSnapshotResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateEnvironmentWithSnapshotResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateEnvironmentWithSnapshotResponseBodyData) SetUid(v string) *CreateEnvironmentWithSnapshotResponseBodyData {
	s.Uid = &v
	return s
}

type CreateEnvironmentWithSnapshotResponse struct {
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateEnvironmentWithSnapshotResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateEnvironmentWithSnapshotResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateEnvironmentWithSnapshotResponse) GoString() string {
	return s.String()
}

func (s *CreateEnvironmentWithSnapshotResponse) SetHeaders(v map[string]*string) *CreateEnvironmentWithSnapshotResponse {
	s.Headers = v
	return s
}

func (s *CreateEnvironmentWithSnapshotResponse) SetBody(v *CreateEnvironmentWithSnapshotResponseBody) *CreateEnvironmentWithSnapshotResponse {
	s.Body = v
	return s
}

type DeleteEnvironmentResponseBody struct {
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	Success    *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrCode    *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
}

func (s DeleteEnvironmentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteEnvironmentResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteEnvironmentResponseBody) SetErrMessage(v string) *DeleteEnvironmentResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DeleteEnvironmentResponseBody) SetSuccess(v bool) *DeleteEnvironmentResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteEnvironmentResponseBody) SetErrCode(v string) *DeleteEnvironmentResponseBody {
	s.ErrCode = &v
	return s
}

type DeleteEnvironmentResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteEnvironmentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteEnvironmentResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteEnvironmentResponse) GoString() string {
	return s.String()
}

func (s *DeleteEnvironmentResponse) SetHeaders(v map[string]*string) *DeleteEnvironmentResponse {
	s.Headers = v
	return s
}

func (s *DeleteEnvironmentResponse) SetBody(v *DeleteEnvironmentResponseBody) *DeleteEnvironmentResponse {
	s.Body = v
	return s
}

type UpdateProductVersionRequest struct {
	CompatibleVersions *string `json:"compatibleVersions,omitempty" xml:"compatibleVersions,omitempty"`
	Description        *string `json:"description,omitempty" xml:"description,omitempty"`
	Version            *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s UpdateProductVersionRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateProductVersionRequest) GoString() string {
	return s.String()
}

func (s *UpdateProductVersionRequest) SetCompatibleVersions(v string) *UpdateProductVersionRequest {
	s.CompatibleVersions = &v
	return s
}

func (s *UpdateProductVersionRequest) SetDescription(v string) *UpdateProductVersionRequest {
	s.Description = &v
	return s
}

func (s *UpdateProductVersionRequest) SetVersion(v string) *UpdateProductVersionRequest {
	s.Version = &v
	return s
}

type UpdateProductVersionResponseBody struct {
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	ErrMsg  *string `json:"errMsg,omitempty" xml:"errMsg,omitempty"`
	Success *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateProductVersionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateProductVersionResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateProductVersionResponseBody) SetErrCode(v string) *UpdateProductVersionResponseBody {
	s.ErrCode = &v
	return s
}

func (s *UpdateProductVersionResponseBody) SetErrMsg(v string) *UpdateProductVersionResponseBody {
	s.ErrMsg = &v
	return s
}

func (s *UpdateProductVersionResponseBody) SetSuccess(v bool) *UpdateProductVersionResponseBody {
	s.Success = &v
	return s
}

type UpdateProductVersionResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateProductVersionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateProductVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateProductVersionResponse) GoString() string {
	return s.String()
}

func (s *UpdateProductVersionResponse) SetHeaders(v map[string]*string) *UpdateProductVersionResponse {
	s.Headers = v
	return s
}

func (s *UpdateProductVersionResponse) SetBody(v *UpdateProductVersionResponseBody) *UpdateProductVersionResponse {
	s.Body = v
	return s
}

type GetChildrenComponentVersionParametersListRequest struct {
	RelationId *string `json:"relation_id,omitempty" xml:"relation_id,omitempty"`
}

func (s GetChildrenComponentVersionParametersListRequest) String() string {
	return tea.Prettify(s)
}

func (s GetChildrenComponentVersionParametersListRequest) GoString() string {
	return s.String()
}

func (s *GetChildrenComponentVersionParametersListRequest) SetRelationId(v string) *GetChildrenComponentVersionParametersListRequest {
	s.RelationId = &v
	return s
}

type GetChildrenComponentVersionParametersListResponseBody struct {
	Data    *GetChildrenComponentVersionParametersListResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	ErrCode *string                                                    `json:"errCode,omitempty" xml:"errCode,omitempty"`
	ErrMsg  *string                                                    `json:"errMsg,omitempty" xml:"errMsg,omitempty"`
	Success *bool                                                      `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetChildrenComponentVersionParametersListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetChildrenComponentVersionParametersListResponseBody) GoString() string {
	return s.String()
}

func (s *GetChildrenComponentVersionParametersListResponseBody) SetData(v *GetChildrenComponentVersionParametersListResponseBodyData) *GetChildrenComponentVersionParametersListResponseBody {
	s.Data = v
	return s
}

func (s *GetChildrenComponentVersionParametersListResponseBody) SetErrCode(v string) *GetChildrenComponentVersionParametersListResponseBody {
	s.ErrCode = &v
	return s
}

func (s *GetChildrenComponentVersionParametersListResponseBody) SetErrMsg(v string) *GetChildrenComponentVersionParametersListResponseBody {
	s.ErrMsg = &v
	return s
}

func (s *GetChildrenComponentVersionParametersListResponseBody) SetSuccess(v bool) *GetChildrenComponentVersionParametersListResponseBody {
	s.Success = &v
	return s
}

type GetChildrenComponentVersionParametersListResponseBodyData struct {
	Annotations *GetChildrenComponentVersionParametersListResponseBodyDataAnnotations `json:"annotations,omitempty" xml:"annotations,omitempty" type:"Struct"`
	Category    *string                                                               `json:"category,omitempty" xml:"category,omitempty"`
	Class       *string                                                               `json:"class,omitempty" xml:"class,omitempty"`
	Description *string                                                               `json:"description,omitempty" xml:"description,omitempty"`
	Documents   []*string                                                             `json:"documents,omitempty" xml:"documents,omitempty" type:"Repeated"`
	Name        *string                                                               `json:"name,omitempty" xml:"name,omitempty"`
	Provider    *string                                                               `json:"provider,omitempty" xml:"provider,omitempty"`
	Uid         *string                                                               `json:"uid,omitempty" xml:"uid,omitempty"`
}

func (s GetChildrenComponentVersionParametersListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetChildrenComponentVersionParametersListResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetChildrenComponentVersionParametersListResponseBodyData) SetAnnotations(v *GetChildrenComponentVersionParametersListResponseBodyDataAnnotations) *GetChildrenComponentVersionParametersListResponseBodyData {
	s.Annotations = v
	return s
}

func (s *GetChildrenComponentVersionParametersListResponseBodyData) SetCategory(v string) *GetChildrenComponentVersionParametersListResponseBodyData {
	s.Category = &v
	return s
}

func (s *GetChildrenComponentVersionParametersListResponseBodyData) SetClass(v string) *GetChildrenComponentVersionParametersListResponseBodyData {
	s.Class = &v
	return s
}

func (s *GetChildrenComponentVersionParametersListResponseBodyData) SetDescription(v string) *GetChildrenComponentVersionParametersListResponseBodyData {
	s.Description = &v
	return s
}

func (s *GetChildrenComponentVersionParametersListResponseBodyData) SetDocuments(v []*string) *GetChildrenComponentVersionParametersListResponseBodyData {
	s.Documents = v
	return s
}

func (s *GetChildrenComponentVersionParametersListResponseBodyData) SetName(v string) *GetChildrenComponentVersionParametersListResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetChildrenComponentVersionParametersListResponseBodyData) SetProvider(v string) *GetChildrenComponentVersionParametersListResponseBodyData {
	s.Provider = &v
	return s
}

func (s *GetChildrenComponentVersionParametersListResponseBodyData) SetUid(v string) *GetChildrenComponentVersionParametersListResponseBodyData {
	s.Uid = &v
	return s
}

type GetChildrenComponentVersionParametersListResponseBodyDataAnnotations struct {
	AdditionalProp1 *string `json:"additionalProp1,omitempty" xml:"additionalProp1,omitempty"`
	AdditionalProp2 *string `json:"additionalProp2,omitempty" xml:"additionalProp2,omitempty"`
	AdditionalProp3 *string `json:"additionalProp3,omitempty" xml:"additionalProp3,omitempty"`
}

func (s GetChildrenComponentVersionParametersListResponseBodyDataAnnotations) String() string {
	return tea.Prettify(s)
}

func (s GetChildrenComponentVersionParametersListResponseBodyDataAnnotations) GoString() string {
	return s.String()
}

func (s *GetChildrenComponentVersionParametersListResponseBodyDataAnnotations) SetAdditionalProp1(v string) *GetChildrenComponentVersionParametersListResponseBodyDataAnnotations {
	s.AdditionalProp1 = &v
	return s
}

func (s *GetChildrenComponentVersionParametersListResponseBodyDataAnnotations) SetAdditionalProp2(v string) *GetChildrenComponentVersionParametersListResponseBodyDataAnnotations {
	s.AdditionalProp2 = &v
	return s
}

func (s *GetChildrenComponentVersionParametersListResponseBodyDataAnnotations) SetAdditionalProp3(v string) *GetChildrenComponentVersionParametersListResponseBodyDataAnnotations {
	s.AdditionalProp3 = &v
	return s
}

type GetChildrenComponentVersionParametersListResponse struct {
	Headers map[string]*string                                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetChildrenComponentVersionParametersListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetChildrenComponentVersionParametersListResponse) String() string {
	return tea.Prettify(s)
}

func (s GetChildrenComponentVersionParametersListResponse) GoString() string {
	return s.String()
}

func (s *GetChildrenComponentVersionParametersListResponse) SetHeaders(v map[string]*string) *GetChildrenComponentVersionParametersListResponse {
	s.Headers = v
	return s
}

func (s *GetChildrenComponentVersionParametersListResponse) SetBody(v *GetChildrenComponentVersionParametersListResponseBody) *GetChildrenComponentVersionParametersListResponse {
	s.Body = v
	return s
}

type CreateSnapshotRequest struct {
	Description        *string `json:"description,omitempty" xml:"description,omitempty"`
	Name               *string `json:"name,omitempty" xml:"name,omitempty"`
	ProductName        *string `json:"productName,omitempty" xml:"productName,omitempty"`
	ProductVersion     *string `json:"productVersion,omitempty" xml:"productVersion,omitempty"`
	ProductVersionDesc *string `json:"productVersionDesc,omitempty" xml:"productVersionDesc,omitempty"`
	Region             *string `json:"region,omitempty" xml:"region,omitempty"`
	Vpcid              *string `json:"vpcid,omitempty" xml:"vpcid,omitempty"`
}

func (s CreateSnapshotRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateSnapshotRequest) GoString() string {
	return s.String()
}

func (s *CreateSnapshotRequest) SetDescription(v string) *CreateSnapshotRequest {
	s.Description = &v
	return s
}

func (s *CreateSnapshotRequest) SetName(v string) *CreateSnapshotRequest {
	s.Name = &v
	return s
}

func (s *CreateSnapshotRequest) SetProductName(v string) *CreateSnapshotRequest {
	s.ProductName = &v
	return s
}

func (s *CreateSnapshotRequest) SetProductVersion(v string) *CreateSnapshotRequest {
	s.ProductVersion = &v
	return s
}

func (s *CreateSnapshotRequest) SetProductVersionDesc(v string) *CreateSnapshotRequest {
	s.ProductVersionDesc = &v
	return s
}

func (s *CreateSnapshotRequest) SetRegion(v string) *CreateSnapshotRequest {
	s.Region = &v
	return s
}

func (s *CreateSnapshotRequest) SetVpcid(v string) *CreateSnapshotRequest {
	s.Vpcid = &v
	return s
}

type CreateSnapshotResponseBody struct {
	Data    *CreateSnapshotResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	ErrCode *string                         `json:"errCode,omitempty" xml:"errCode,omitempty"`
	ErrMsg  *string                         `json:"errMsg,omitempty" xml:"errMsg,omitempty"`
	Success *bool                           `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreateSnapshotResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateSnapshotResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSnapshotResponseBody) SetData(v *CreateSnapshotResponseBodyData) *CreateSnapshotResponseBody {
	s.Data = v
	return s
}

func (s *CreateSnapshotResponseBody) SetErrCode(v string) *CreateSnapshotResponseBody {
	s.ErrCode = &v
	return s
}

func (s *CreateSnapshotResponseBody) SetErrMsg(v string) *CreateSnapshotResponseBody {
	s.ErrMsg = &v
	return s
}

func (s *CreateSnapshotResponseBody) SetSuccess(v bool) *CreateSnapshotResponseBody {
	s.Success = &v
	return s
}

type CreateSnapshotResponseBodyData struct {
	Uid *string `json:"uid,omitempty" xml:"uid,omitempty"`
}

func (s CreateSnapshotResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateSnapshotResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateSnapshotResponseBodyData) SetUid(v string) *CreateSnapshotResponseBodyData {
	s.Uid = &v
	return s
}

type CreateSnapshotResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateSnapshotResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateSnapshotResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateSnapshotResponse) GoString() string {
	return s.String()
}

func (s *CreateSnapshotResponse) SetHeaders(v map[string]*string) *CreateSnapshotResponse {
	s.Headers = v
	return s
}

func (s *CreateSnapshotResponse) SetBody(v *CreateSnapshotResponseBody) *CreateSnapshotResponse {
	s.Body = v
	return s
}

type GetLatestVersionDifferencesRequest struct {
	// 上一个产品版本id
	PreVersionID *string `json:"preVersionID,omitempty" xml:"preVersionID,omitempty"`
}

func (s GetLatestVersionDifferencesRequest) String() string {
	return tea.Prettify(s)
}

func (s GetLatestVersionDifferencesRequest) GoString() string {
	return s.String()
}

func (s *GetLatestVersionDifferencesRequest) SetPreVersionID(v string) *GetLatestVersionDifferencesRequest {
	s.PreVersionID = &v
	return s
}

type GetLatestVersionDifferencesResponseBody struct {
	// Id of the request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetLatestVersionDifferencesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetLatestVersionDifferencesResponseBody) GoString() string {
	return s.String()
}

func (s *GetLatestVersionDifferencesResponseBody) SetRequestId(v string) *GetLatestVersionDifferencesResponseBody {
	s.RequestId = &v
	return s
}

type GetLatestVersionDifferencesResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetLatestVersionDifferencesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetLatestVersionDifferencesResponse) String() string {
	return tea.Prettify(s)
}

func (s GetLatestVersionDifferencesResponse) GoString() string {
	return s.String()
}

func (s *GetLatestVersionDifferencesResponse) SetHeaders(v map[string]*string) *GetLatestVersionDifferencesResponse {
	s.Headers = v
	return s
}

func (s *GetLatestVersionDifferencesResponse) SetBody(v *GetLatestVersionDifferencesResponseBody) *GetLatestVersionDifferencesResponse {
	s.Body = v
	return s
}

type DeleteEnvironmentNodeRequest struct {
	EnvUID *string `json:"envUID,omitempty" xml:"envUID,omitempty"`
}

func (s DeleteEnvironmentNodeRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteEnvironmentNodeRequest) GoString() string {
	return s.String()
}

func (s *DeleteEnvironmentNodeRequest) SetEnvUID(v string) *DeleteEnvironmentNodeRequest {
	s.EnvUID = &v
	return s
}

type DeleteEnvironmentNodeResponseBody struct {
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	Success    *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrCode    *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
}

func (s DeleteEnvironmentNodeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteEnvironmentNodeResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteEnvironmentNodeResponseBody) SetErrMessage(v string) *DeleteEnvironmentNodeResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DeleteEnvironmentNodeResponseBody) SetSuccess(v bool) *DeleteEnvironmentNodeResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteEnvironmentNodeResponseBody) SetErrCode(v string) *DeleteEnvironmentNodeResponseBody {
	s.ErrCode = &v
	return s
}

type DeleteEnvironmentNodeResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteEnvironmentNodeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteEnvironmentNodeResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteEnvironmentNodeResponse) GoString() string {
	return s.String()
}

func (s *DeleteEnvironmentNodeResponse) SetHeaders(v map[string]*string) *DeleteEnvironmentNodeResponse {
	s.Headers = v
	return s
}

func (s *DeleteEnvironmentNodeResponse) SetBody(v *DeleteEnvironmentNodeResponseBody) *DeleteEnvironmentNodeResponse {
	s.Body = v
	return s
}

type ApplyComponentRequest struct {
	Annotations         *string                           `json:"annotations,omitempty" xml:"annotations,omitempty"`
	AppVersion          *string                           `json:"appVersion,omitempty" xml:"appVersion,omitempty"`
	Category            *string                           `json:"category,omitempty" xml:"category,omitempty"`
	ComponentClass      *string                           `json:"componentClass,omitempty" xml:"componentClass,omitempty"`
	Description         *string                           `json:"description,omitempty" xml:"description,omitempty"`
	Documents           *string                           `json:"documents,omitempty" xml:"documents,omitempty"`
	ImagesMapping       *string                           `json:"imagesMapping,omitempty" xml:"imagesMapping,omitempty"`
	Name                *string                           `json:"name,omitempty" xml:"name,omitempty"`
	Namespace           *string                           `json:"namespace,omitempty" xml:"namespace,omitempty"`
	OrchestrationType   *string                           `json:"orchestrationType,omitempty" xml:"orchestrationType,omitempty"`
	OrchestrationValues *string                           `json:"orchestrationValues,omitempty" xml:"orchestrationValues,omitempty"`
	PackageURL          *string                           `json:"packageURL,omitempty" xml:"packageURL,omitempty"`
	ParentComponent     *bool                             `json:"parentComponent,omitempty" xml:"parentComponent,omitempty"`
	Platforms           []*ApplyComponentRequestPlatforms `json:"platforms,omitempty" xml:"platforms,omitempty" type:"Repeated"`
	Priority            *int32                            `json:"priority,omitempty" xml:"priority,omitempty"`
	Provider            *string                           `json:"provider,omitempty" xml:"provider,omitempty"`
	Public              *bool                             `json:"public,omitempty" xml:"public,omitempty"`
	Readme              *string                           `json:"readme,omitempty" xml:"readme,omitempty"`
	Resources           *string                           `json:"resources,omitempty" xml:"resources,omitempty"`
	Singleton           *bool                             `json:"singleton,omitempty" xml:"singleton,omitempty"`
	Version             *string                           `json:"version,omitempty" xml:"version,omitempty"`
}

func (s ApplyComponentRequest) String() string {
	return tea.Prettify(s)
}

func (s ApplyComponentRequest) GoString() string {
	return s.String()
}

func (s *ApplyComponentRequest) SetAnnotations(v string) *ApplyComponentRequest {
	s.Annotations = &v
	return s
}

func (s *ApplyComponentRequest) SetAppVersion(v string) *ApplyComponentRequest {
	s.AppVersion = &v
	return s
}

func (s *ApplyComponentRequest) SetCategory(v string) *ApplyComponentRequest {
	s.Category = &v
	return s
}

func (s *ApplyComponentRequest) SetComponentClass(v string) *ApplyComponentRequest {
	s.ComponentClass = &v
	return s
}

func (s *ApplyComponentRequest) SetDescription(v string) *ApplyComponentRequest {
	s.Description = &v
	return s
}

func (s *ApplyComponentRequest) SetDocuments(v string) *ApplyComponentRequest {
	s.Documents = &v
	return s
}

func (s *ApplyComponentRequest) SetImagesMapping(v string) *ApplyComponentRequest {
	s.ImagesMapping = &v
	return s
}

func (s *ApplyComponentRequest) SetName(v string) *ApplyComponentRequest {
	s.Name = &v
	return s
}

func (s *ApplyComponentRequest) SetNamespace(v string) *ApplyComponentRequest {
	s.Namespace = &v
	return s
}

func (s *ApplyComponentRequest) SetOrchestrationType(v string) *ApplyComponentRequest {
	s.OrchestrationType = &v
	return s
}

func (s *ApplyComponentRequest) SetOrchestrationValues(v string) *ApplyComponentRequest {
	s.OrchestrationValues = &v
	return s
}

func (s *ApplyComponentRequest) SetPackageURL(v string) *ApplyComponentRequest {
	s.PackageURL = &v
	return s
}

func (s *ApplyComponentRequest) SetParentComponent(v bool) *ApplyComponentRequest {
	s.ParentComponent = &v
	return s
}

func (s *ApplyComponentRequest) SetPlatforms(v []*ApplyComponentRequestPlatforms) *ApplyComponentRequest {
	s.Platforms = v
	return s
}

func (s *ApplyComponentRequest) SetPriority(v int32) *ApplyComponentRequest {
	s.Priority = &v
	return s
}

func (s *ApplyComponentRequest) SetProvider(v string) *ApplyComponentRequest {
	s.Provider = &v
	return s
}

func (s *ApplyComponentRequest) SetPublic(v bool) *ApplyComponentRequest {
	s.Public = &v
	return s
}

func (s *ApplyComponentRequest) SetReadme(v string) *ApplyComponentRequest {
	s.Readme = &v
	return s
}

func (s *ApplyComponentRequest) SetResources(v string) *ApplyComponentRequest {
	s.Resources = &v
	return s
}

func (s *ApplyComponentRequest) SetSingleton(v bool) *ApplyComponentRequest {
	s.Singleton = &v
	return s
}

func (s *ApplyComponentRequest) SetVersion(v string) *ApplyComponentRequest {
	s.Version = &v
	return s
}

type ApplyComponentRequestPlatforms struct {
	Architecture *string `json:"architecture,omitempty" xml:"architecture,omitempty"`
	Os           *string `json:"os,omitempty" xml:"os,omitempty"`
}

func (s ApplyComponentRequestPlatforms) String() string {
	return tea.Prettify(s)
}

func (s ApplyComponentRequestPlatforms) GoString() string {
	return s.String()
}

func (s *ApplyComponentRequestPlatforms) SetArchitecture(v string) *ApplyComponentRequestPlatforms {
	s.Architecture = &v
	return s
}

func (s *ApplyComponentRequestPlatforms) SetOs(v string) *ApplyComponentRequestPlatforms {
	s.Os = &v
	return s
}

type ApplyComponentResponseBody struct {
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	ErrMsg  *string `json:"errMsg,omitempty" xml:"errMsg,omitempty"`
	Status  *string `json:"status,omitempty" xml:"status,omitempty"`
	Success *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ApplyComponentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ApplyComponentResponseBody) GoString() string {
	return s.String()
}

func (s *ApplyComponentResponseBody) SetErrCode(v string) *ApplyComponentResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ApplyComponentResponseBody) SetErrMsg(v string) *ApplyComponentResponseBody {
	s.ErrMsg = &v
	return s
}

func (s *ApplyComponentResponseBody) SetStatus(v string) *ApplyComponentResponseBody {
	s.Status = &v
	return s
}

func (s *ApplyComponentResponseBody) SetSuccess(v bool) *ApplyComponentResponseBody {
	s.Success = &v
	return s
}

type ApplyComponentResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ApplyComponentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ApplyComponentResponse) String() string {
	return tea.Prettify(s)
}

func (s ApplyComponentResponse) GoString() string {
	return s.String()
}

func (s *ApplyComponentResponse) SetHeaders(v map[string]*string) *ApplyComponentResponse {
	s.Headers = v
	return s
}

func (s *ApplyComponentResponse) SetBody(v *ApplyComponentResponseBody) *ApplyComponentResponse {
	s.Body = v
	return s
}

type GetSnapshotInstancesRequest struct {
	PageSize   *int32  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	PageNum    *int32  `json:"pageNum,omitempty" xml:"pageNum,omitempty"`
	SortKey    *string `json:"sortKey,omitempty" xml:"sortKey,omitempty"`
	SortDirect *string `json:"sortDirect,omitempty" xml:"sortDirect,omitempty"`
}

func (s GetSnapshotInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSnapshotInstancesRequest) GoString() string {
	return s.String()
}

func (s *GetSnapshotInstancesRequest) SetPageSize(v int32) *GetSnapshotInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *GetSnapshotInstancesRequest) SetPageNum(v int32) *GetSnapshotInstancesRequest {
	s.PageNum = &v
	return s
}

func (s *GetSnapshotInstancesRequest) SetSortKey(v string) *GetSnapshotInstancesRequest {
	s.SortKey = &v
	return s
}

func (s *GetSnapshotInstancesRequest) SetSortDirect(v string) *GetSnapshotInstancesRequest {
	s.SortDirect = &v
	return s
}

type GetSnapshotInstancesResponseBody struct {
	Data    *GetSnapshotInstancesResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	ErrCode *string                               `json:"errCode,omitempty" xml:"errCode,omitempty"`
	ErrMsg  *string                               `json:"errMsg,omitempty" xml:"errMsg,omitempty"`
	Success *bool                                 `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetSnapshotInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSnapshotInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *GetSnapshotInstancesResponseBody) SetData(v *GetSnapshotInstancesResponseBodyData) *GetSnapshotInstancesResponseBody {
	s.Data = v
	return s
}

func (s *GetSnapshotInstancesResponseBody) SetErrCode(v string) *GetSnapshotInstancesResponseBody {
	s.ErrCode = &v
	return s
}

func (s *GetSnapshotInstancesResponseBody) SetErrMsg(v string) *GetSnapshotInstancesResponseBody {
	s.ErrMsg = &v
	return s
}

func (s *GetSnapshotInstancesResponseBody) SetSuccess(v bool) *GetSnapshotInstancesResponseBody {
	s.Success = &v
	return s
}

type GetSnapshotInstancesResponseBodyData struct {
	List     []*GetSnapshotInstancesResponseBodyDataList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	PageNum  *int32                                      `json:"pageNum,omitempty" xml:"pageNum,omitempty"`
	PageSize *int32                                      `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	Total    *int32                                      `json:"total,omitempty" xml:"total,omitempty"`
}

func (s GetSnapshotInstancesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetSnapshotInstancesResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetSnapshotInstancesResponseBodyData) SetList(v []*GetSnapshotInstancesResponseBodyDataList) *GetSnapshotInstancesResponseBodyData {
	s.List = v
	return s
}

func (s *GetSnapshotInstancesResponseBodyData) SetPageNum(v int32) *GetSnapshotInstancesResponseBodyData {
	s.PageNum = &v
	return s
}

func (s *GetSnapshotInstancesResponseBodyData) SetPageSize(v int32) *GetSnapshotInstancesResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *GetSnapshotInstancesResponseBodyData) SetTotal(v int32) *GetSnapshotInstancesResponseBodyData {
	s.Total = &v
	return s
}

type GetSnapshotInstancesResponseBodyDataList struct {
	Annotations       *GetSnapshotInstancesResponseBodyDataListAnnotations `json:"annotations,omitempty" xml:"annotations,omitempty" type:"Struct"`
	Cpu               *int32                                               `json:"cpu,omitempty" xml:"cpu,omitempty"`
	EcsInstanceID     *string                                              `json:"ecsInstanceID,omitempty" xml:"ecsInstanceID,omitempty"`
	HostName          *string                                              `json:"hostName,omitempty" xml:"hostName,omitempty"`
	Identifier        *string                                              `json:"identifier,omitempty" xml:"identifier,omitempty"`
	ImageID           *string                                              `json:"imageID,omitempty" xml:"imageID,omitempty"`
	InstanceType      *string                                              `json:"instanceType,omitempty" xml:"instanceType,omitempty"`
	InternetBandwidth *int32                                               `json:"internetBandwidth,omitempty" xml:"internetBandwidth,omitempty"`
	JoinSnapshot      *bool                                                `json:"joinSnapshot,omitempty" xml:"joinSnapshot,omitempty"`
	Memory            *int32                                               `json:"memory,omitempty" xml:"memory,omitempty"`
	PrivateIP         *string                                              `json:"privateIP,omitempty" xml:"privateIP,omitempty"`
	PublicIP          *string                                              `json:"publicIP,omitempty" xml:"publicIP,omitempty"`
	RootPassword      *string                                              `json:"rootPassword,omitempty" xml:"rootPassword,omitempty"`
	StorageTotalSize  *int32                                               `json:"storageTotalSize,omitempty" xml:"storageTotalSize,omitempty"`
	Uid               *string                                              `json:"uid,omitempty" xml:"uid,omitempty"`
}

func (s GetSnapshotInstancesResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s GetSnapshotInstancesResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *GetSnapshotInstancesResponseBodyDataList) SetAnnotations(v *GetSnapshotInstancesResponseBodyDataListAnnotations) *GetSnapshotInstancesResponseBodyDataList {
	s.Annotations = v
	return s
}

func (s *GetSnapshotInstancesResponseBodyDataList) SetCpu(v int32) *GetSnapshotInstancesResponseBodyDataList {
	s.Cpu = &v
	return s
}

func (s *GetSnapshotInstancesResponseBodyDataList) SetEcsInstanceID(v string) *GetSnapshotInstancesResponseBodyDataList {
	s.EcsInstanceID = &v
	return s
}

func (s *GetSnapshotInstancesResponseBodyDataList) SetHostName(v string) *GetSnapshotInstancesResponseBodyDataList {
	s.HostName = &v
	return s
}

func (s *GetSnapshotInstancesResponseBodyDataList) SetIdentifier(v string) *GetSnapshotInstancesResponseBodyDataList {
	s.Identifier = &v
	return s
}

func (s *GetSnapshotInstancesResponseBodyDataList) SetImageID(v string) *GetSnapshotInstancesResponseBodyDataList {
	s.ImageID = &v
	return s
}

func (s *GetSnapshotInstancesResponseBodyDataList) SetInstanceType(v string) *GetSnapshotInstancesResponseBodyDataList {
	s.InstanceType = &v
	return s
}

func (s *GetSnapshotInstancesResponseBodyDataList) SetInternetBandwidth(v int32) *GetSnapshotInstancesResponseBodyDataList {
	s.InternetBandwidth = &v
	return s
}

func (s *GetSnapshotInstancesResponseBodyDataList) SetJoinSnapshot(v bool) *GetSnapshotInstancesResponseBodyDataList {
	s.JoinSnapshot = &v
	return s
}

func (s *GetSnapshotInstancesResponseBodyDataList) SetMemory(v int32) *GetSnapshotInstancesResponseBodyDataList {
	s.Memory = &v
	return s
}

func (s *GetSnapshotInstancesResponseBodyDataList) SetPrivateIP(v string) *GetSnapshotInstancesResponseBodyDataList {
	s.PrivateIP = &v
	return s
}

func (s *GetSnapshotInstancesResponseBodyDataList) SetPublicIP(v string) *GetSnapshotInstancesResponseBodyDataList {
	s.PublicIP = &v
	return s
}

func (s *GetSnapshotInstancesResponseBodyDataList) SetRootPassword(v string) *GetSnapshotInstancesResponseBodyDataList {
	s.RootPassword = &v
	return s
}

func (s *GetSnapshotInstancesResponseBodyDataList) SetStorageTotalSize(v int32) *GetSnapshotInstancesResponseBodyDataList {
	s.StorageTotalSize = &v
	return s
}

func (s *GetSnapshotInstancesResponseBodyDataList) SetUid(v string) *GetSnapshotInstancesResponseBodyDataList {
	s.Uid = &v
	return s
}

type GetSnapshotInstancesResponseBodyDataListAnnotations struct {
	AdditionalProp1 *string `json:"additionalProp1,omitempty" xml:"additionalProp1,omitempty"`
	AdditionalProp2 *string `json:"additionalProp2,omitempty" xml:"additionalProp2,omitempty"`
	AdditionalProp3 *string `json:"additionalProp3,omitempty" xml:"additionalProp3,omitempty"`
}

func (s GetSnapshotInstancesResponseBodyDataListAnnotations) String() string {
	return tea.Prettify(s)
}

func (s GetSnapshotInstancesResponseBodyDataListAnnotations) GoString() string {
	return s.String()
}

func (s *GetSnapshotInstancesResponseBodyDataListAnnotations) SetAdditionalProp1(v string) *GetSnapshotInstancesResponseBodyDataListAnnotations {
	s.AdditionalProp1 = &v
	return s
}

func (s *GetSnapshotInstancesResponseBodyDataListAnnotations) SetAdditionalProp2(v string) *GetSnapshotInstancesResponseBodyDataListAnnotations {
	s.AdditionalProp2 = &v
	return s
}

func (s *GetSnapshotInstancesResponseBodyDataListAnnotations) SetAdditionalProp3(v string) *GetSnapshotInstancesResponseBodyDataListAnnotations {
	s.AdditionalProp3 = &v
	return s
}

type GetSnapshotInstancesResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetSnapshotInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetSnapshotInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSnapshotInstancesResponse) GoString() string {
	return s.String()
}

func (s *GetSnapshotInstancesResponse) SetHeaders(v map[string]*string) *GetSnapshotInstancesResponse {
	s.Headers = v
	return s
}

func (s *GetSnapshotInstancesResponse) SetBody(v *GetSnapshotInstancesResponseBody) *GetSnapshotInstancesResponse {
	s.Body = v
	return s
}

type ListEnvironmentsRequest struct {
	PageNum        *int32  `json:"pageNum,omitempty" xml:"pageNum,omitempty"`
	PageSize       *int32  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	Name           *string `json:"name,omitempty" xml:"name,omitempty"`
	Fuzzy          *string `json:"fuzzy,omitempty" xml:"fuzzy,omitempty"`
	InstanceStatus *string `json:"instanceStatus,omitempty" xml:"instanceStatus,omitempty"`
	VendorType     *string `json:"vendorType,omitempty" xml:"vendorType,omitempty"`
}

func (s ListEnvironmentsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListEnvironmentsRequest) GoString() string {
	return s.String()
}

func (s *ListEnvironmentsRequest) SetPageNum(v int32) *ListEnvironmentsRequest {
	s.PageNum = &v
	return s
}

func (s *ListEnvironmentsRequest) SetPageSize(v int32) *ListEnvironmentsRequest {
	s.PageSize = &v
	return s
}

func (s *ListEnvironmentsRequest) SetName(v string) *ListEnvironmentsRequest {
	s.Name = &v
	return s
}

func (s *ListEnvironmentsRequest) SetFuzzy(v string) *ListEnvironmentsRequest {
	s.Fuzzy = &v
	return s
}

func (s *ListEnvironmentsRequest) SetInstanceStatus(v string) *ListEnvironmentsRequest {
	s.InstanceStatus = &v
	return s
}

func (s *ListEnvironmentsRequest) SetVendorType(v string) *ListEnvironmentsRequest {
	s.VendorType = &v
	return s
}

type ListEnvironmentsResponseBody struct {
	Data    *ListEnvironmentsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	ErrCode *string                           `json:"errCode,omitempty" xml:"errCode,omitempty"`
	ErrMsg  *string                           `json:"errMsg,omitempty" xml:"errMsg,omitempty"`
	Success *bool                             `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ListEnvironmentsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListEnvironmentsResponseBody) GoString() string {
	return s.String()
}

func (s *ListEnvironmentsResponseBody) SetData(v *ListEnvironmentsResponseBodyData) *ListEnvironmentsResponseBody {
	s.Data = v
	return s
}

func (s *ListEnvironmentsResponseBody) SetErrCode(v string) *ListEnvironmentsResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ListEnvironmentsResponseBody) SetErrMsg(v string) *ListEnvironmentsResponseBody {
	s.ErrMsg = &v
	return s
}

func (s *ListEnvironmentsResponseBody) SetSuccess(v bool) *ListEnvironmentsResponseBody {
	s.Success = &v
	return s
}

type ListEnvironmentsResponseBodyData struct {
	List     []*ListEnvironmentsResponseBodyDataList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	PageNum  *int32                                  `json:"pageNum,omitempty" xml:"pageNum,omitempty"`
	PageSize *int32                                  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	Total    *int32                                  `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListEnvironmentsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListEnvironmentsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListEnvironmentsResponseBodyData) SetList(v []*ListEnvironmentsResponseBodyDataList) *ListEnvironmentsResponseBodyData {
	s.List = v
	return s
}

func (s *ListEnvironmentsResponseBodyData) SetPageNum(v int32) *ListEnvironmentsResponseBodyData {
	s.PageNum = &v
	return s
}

func (s *ListEnvironmentsResponseBodyData) SetPageSize(v int32) *ListEnvironmentsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListEnvironmentsResponseBodyData) SetTotal(v int32) *ListEnvironmentsResponseBodyData {
	s.Total = &v
	return s
}

type ListEnvironmentsResponseBodyDataList struct {
	CreatedAt         *string                                       `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	Description       *string                                       `json:"description,omitempty" xml:"description,omitempty"`
	Id                *int32                                        `json:"id,omitempty" xml:"id,omitempty"`
	Location          *string                                       `json:"location,omitempty" xml:"location,omitempty"`
	Name              *string                                       `json:"name,omitempty" xml:"name,omitempty"`
	Platform          *ListEnvironmentsResponseBodyDataListPlatform `json:"platform,omitempty" xml:"platform,omitempty" type:"Struct"`
	ProductName       *string                                       `json:"productName,omitempty" xml:"productName,omitempty"`
	ProductVersion    *string                                       `json:"productVersion,omitempty" xml:"productVersion,omitempty"`
	ProductVersionUID *string                                       `json:"productVersionUID,omitempty" xml:"productVersionUID,omitempty"`
	Uid               *string                                       `json:"uid,omitempty" xml:"uid,omitempty"`
	VendorType        *string                                       `json:"vendorType,omitempty" xml:"vendorType,omitempty"`
}

func (s ListEnvironmentsResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s ListEnvironmentsResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *ListEnvironmentsResponseBodyDataList) SetCreatedAt(v string) *ListEnvironmentsResponseBodyDataList {
	s.CreatedAt = &v
	return s
}

func (s *ListEnvironmentsResponseBodyDataList) SetDescription(v string) *ListEnvironmentsResponseBodyDataList {
	s.Description = &v
	return s
}

func (s *ListEnvironmentsResponseBodyDataList) SetId(v int32) *ListEnvironmentsResponseBodyDataList {
	s.Id = &v
	return s
}

func (s *ListEnvironmentsResponseBodyDataList) SetLocation(v string) *ListEnvironmentsResponseBodyDataList {
	s.Location = &v
	return s
}

func (s *ListEnvironmentsResponseBodyDataList) SetName(v string) *ListEnvironmentsResponseBodyDataList {
	s.Name = &v
	return s
}

func (s *ListEnvironmentsResponseBodyDataList) SetPlatform(v *ListEnvironmentsResponseBodyDataListPlatform) *ListEnvironmentsResponseBodyDataList {
	s.Platform = v
	return s
}

func (s *ListEnvironmentsResponseBodyDataList) SetProductName(v string) *ListEnvironmentsResponseBodyDataList {
	s.ProductName = &v
	return s
}

func (s *ListEnvironmentsResponseBodyDataList) SetProductVersion(v string) *ListEnvironmentsResponseBodyDataList {
	s.ProductVersion = &v
	return s
}

func (s *ListEnvironmentsResponseBodyDataList) SetProductVersionUID(v string) *ListEnvironmentsResponseBodyDataList {
	s.ProductVersionUID = &v
	return s
}

func (s *ListEnvironmentsResponseBodyDataList) SetUid(v string) *ListEnvironmentsResponseBodyDataList {
	s.Uid = &v
	return s
}

func (s *ListEnvironmentsResponseBodyDataList) SetVendorType(v string) *ListEnvironmentsResponseBodyDataList {
	s.VendorType = &v
	return s
}

type ListEnvironmentsResponseBodyDataListPlatform struct {
	Architecture *string `json:"architecture,omitempty" xml:"architecture,omitempty"`
	Os           *string `json:"os,omitempty" xml:"os,omitempty"`
}

func (s ListEnvironmentsResponseBodyDataListPlatform) String() string {
	return tea.Prettify(s)
}

func (s ListEnvironmentsResponseBodyDataListPlatform) GoString() string {
	return s.String()
}

func (s *ListEnvironmentsResponseBodyDataListPlatform) SetArchitecture(v string) *ListEnvironmentsResponseBodyDataListPlatform {
	s.Architecture = &v
	return s
}

func (s *ListEnvironmentsResponseBodyDataListPlatform) SetOs(v string) *ListEnvironmentsResponseBodyDataListPlatform {
	s.Os = &v
	return s
}

type ListEnvironmentsResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListEnvironmentsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListEnvironmentsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListEnvironmentsResponse) GoString() string {
	return s.String()
}

func (s *ListEnvironmentsResponse) SetHeaders(v map[string]*string) *ListEnvironmentsResponse {
	s.Headers = v
	return s
}

func (s *ListEnvironmentsResponse) SetBody(v *ListEnvironmentsResponseBody) *ListEnvironmentsResponse {
	s.Body = v
	return s
}

type CheckSLRRequest struct {
	Uid *string `json:"uid,omitempty" xml:"uid,omitempty"`
}

func (s CheckSLRRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckSLRRequest) GoString() string {
	return s.String()
}

func (s *CheckSLRRequest) SetUid(v string) *CheckSLRRequest {
	s.Uid = &v
	return s
}

type CheckSLRResponseBody struct {
	Data      *string `json:"data,omitempty" xml:"data,omitempty"`
	ErrCode   *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	ErrMsg    *string `json:"errMsg,omitempty" xml:"errMsg,omitempty"`
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	Success   *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CheckSLRResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckSLRResponseBody) GoString() string {
	return s.String()
}

func (s *CheckSLRResponseBody) SetData(v string) *CheckSLRResponseBody {
	s.Data = &v
	return s
}

func (s *CheckSLRResponseBody) SetErrCode(v string) *CheckSLRResponseBody {
	s.ErrCode = &v
	return s
}

func (s *CheckSLRResponseBody) SetErrMsg(v string) *CheckSLRResponseBody {
	s.ErrMsg = &v
	return s
}

func (s *CheckSLRResponseBody) SetErrorCode(v string) *CheckSLRResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CheckSLRResponseBody) SetSuccess(v bool) *CheckSLRResponseBody {
	s.Success = &v
	return s
}

type CheckSLRResponse struct {
	Headers map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CheckSLRResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CheckSLRResponse) String() string {
	return tea.Prettify(s)
}

func (s CheckSLRResponse) GoString() string {
	return s.String()
}

func (s *CheckSLRResponse) SetHeaders(v map[string]*string) *CheckSLRResponse {
	s.Headers = v
	return s
}

func (s *CheckSLRResponse) SetBody(v *CheckSLRResponseBody) *CheckSLRResponse {
	s.Body = v
	return s
}

type UpdateProductRequest struct {
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
}

func (s UpdateProductRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateProductRequest) GoString() string {
	return s.String()
}

func (s *UpdateProductRequest) SetDescription(v string) *UpdateProductRequest {
	s.Description = &v
	return s
}

type UpdateProductResponseBody struct {
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	ErrMsg  *string `json:"errMsg,omitempty" xml:"errMsg,omitempty"`
	Success *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateProductResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateProductResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateProductResponseBody) SetErrCode(v string) *UpdateProductResponseBody {
	s.ErrCode = &v
	return s
}

func (s *UpdateProductResponseBody) SetErrMsg(v string) *UpdateProductResponseBody {
	s.ErrMsg = &v
	return s
}

func (s *UpdateProductResponseBody) SetSuccess(v bool) *UpdateProductResponseBody {
	s.Success = &v
	return s
}

type UpdateProductResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateProductResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateProductResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateProductResponse) GoString() string {
	return s.String()
}

func (s *UpdateProductResponse) SetHeaders(v map[string]*string) *UpdateProductResponse {
	s.Headers = v
	return s
}

func (s *UpdateProductResponse) SetBody(v *UpdateProductResponseBody) *UpdateProductResponse {
	s.Body = v
	return s
}

type ApplyComponentsRequest struct {
	ChildrenList []*ApplyComponentsRequestChildrenList `json:"childrenList,omitempty" xml:"childrenList,omitempty" type:"Repeated"`
	Component    *string                               `json:"component,omitempty" xml:"component,omitempty"`
}

func (s ApplyComponentsRequest) String() string {
	return tea.Prettify(s)
}

func (s ApplyComponentsRequest) GoString() string {
	return s.String()
}

func (s *ApplyComponentsRequest) SetChildrenList(v []*ApplyComponentsRequestChildrenList) *ApplyComponentsRequest {
	s.ChildrenList = v
	return s
}

func (s *ApplyComponentsRequest) SetComponent(v string) *ApplyComponentsRequest {
	s.Component = &v
	return s
}

type ApplyComponentsRequestChildrenList struct {
	Annotations         *string                                        `json:"annotations,omitempty" xml:"annotations,omitempty"`
	AppVersion          *string                                        `json:"appVersion,omitempty" xml:"appVersion,omitempty"`
	Category            *string                                        `json:"category,omitempty" xml:"category,omitempty"`
	ComponentClass      *string                                        `json:"componentClass,omitempty" xml:"componentClass,omitempty"`
	Description         *string                                        `json:"description,omitempty" xml:"description,omitempty"`
	Documents           *string                                        `json:"documents,omitempty" xml:"documents,omitempty"`
	ImagesMapping       *string                                        `json:"imagesMapping,omitempty" xml:"imagesMapping,omitempty"`
	Name                *string                                        `json:"name,omitempty" xml:"name,omitempty"`
	Namespace           *string                                        `json:"namespace,omitempty" xml:"namespace,omitempty"`
	OrchestrationType   *string                                        `json:"orchestrationType,omitempty" xml:"orchestrationType,omitempty"`
	OrchestrationValues *string                                        `json:"orchestrationValues,omitempty" xml:"orchestrationValues,omitempty"`
	PackageURL          *string                                        `json:"packageURL,omitempty" xml:"packageURL,omitempty"`
	ParentComponent     *bool                                          `json:"parentComponent,omitempty" xml:"parentComponent,omitempty"`
	Platforms           []*ApplyComponentsRequestChildrenListPlatforms `json:"platforms,omitempty" xml:"platforms,omitempty" type:"Repeated"`
	Priority            *int32                                         `json:"priority,omitempty" xml:"priority,omitempty"`
	Provider            *string                                        `json:"provider,omitempty" xml:"provider,omitempty"`
	Public              *bool                                          `json:"public,omitempty" xml:"public,omitempty"`
	Readme              *string                                        `json:"readme,omitempty" xml:"readme,omitempty"`
	Resources           *string                                        `json:"resources,omitempty" xml:"resources,omitempty"`
	Singleton           *bool                                          `json:"singleton,omitempty" xml:"singleton,omitempty"`
	Version             *string                                        `json:"version,omitempty" xml:"version,omitempty"`
}

func (s ApplyComponentsRequestChildrenList) String() string {
	return tea.Prettify(s)
}

func (s ApplyComponentsRequestChildrenList) GoString() string {
	return s.String()
}

func (s *ApplyComponentsRequestChildrenList) SetAnnotations(v string) *ApplyComponentsRequestChildrenList {
	s.Annotations = &v
	return s
}

func (s *ApplyComponentsRequestChildrenList) SetAppVersion(v string) *ApplyComponentsRequestChildrenList {
	s.AppVersion = &v
	return s
}

func (s *ApplyComponentsRequestChildrenList) SetCategory(v string) *ApplyComponentsRequestChildrenList {
	s.Category = &v
	return s
}

func (s *ApplyComponentsRequestChildrenList) SetComponentClass(v string) *ApplyComponentsRequestChildrenList {
	s.ComponentClass = &v
	return s
}

func (s *ApplyComponentsRequestChildrenList) SetDescription(v string) *ApplyComponentsRequestChildrenList {
	s.Description = &v
	return s
}

func (s *ApplyComponentsRequestChildrenList) SetDocuments(v string) *ApplyComponentsRequestChildrenList {
	s.Documents = &v
	return s
}

func (s *ApplyComponentsRequestChildrenList) SetImagesMapping(v string) *ApplyComponentsRequestChildrenList {
	s.ImagesMapping = &v
	return s
}

func (s *ApplyComponentsRequestChildrenList) SetName(v string) *ApplyComponentsRequestChildrenList {
	s.Name = &v
	return s
}

func (s *ApplyComponentsRequestChildrenList) SetNamespace(v string) *ApplyComponentsRequestChildrenList {
	s.Namespace = &v
	return s
}

func (s *ApplyComponentsRequestChildrenList) SetOrchestrationType(v string) *ApplyComponentsRequestChildrenList {
	s.OrchestrationType = &v
	return s
}

func (s *ApplyComponentsRequestChildrenList) SetOrchestrationValues(v string) *ApplyComponentsRequestChildrenList {
	s.OrchestrationValues = &v
	return s
}

func (s *ApplyComponentsRequestChildrenList) SetPackageURL(v string) *ApplyComponentsRequestChildrenList {
	s.PackageURL = &v
	return s
}

func (s *ApplyComponentsRequestChildrenList) SetParentComponent(v bool) *ApplyComponentsRequestChildrenList {
	s.ParentComponent = &v
	return s
}

func (s *ApplyComponentsRequestChildrenList) SetPlatforms(v []*ApplyComponentsRequestChildrenListPlatforms) *ApplyComponentsRequestChildrenList {
	s.Platforms = v
	return s
}

func (s *ApplyComponentsRequestChildrenList) SetPriority(v int32) *ApplyComponentsRequestChildrenList {
	s.Priority = &v
	return s
}

func (s *ApplyComponentsRequestChildrenList) SetProvider(v string) *ApplyComponentsRequestChildrenList {
	s.Provider = &v
	return s
}

func (s *ApplyComponentsRequestChildrenList) SetPublic(v bool) *ApplyComponentsRequestChildrenList {
	s.Public = &v
	return s
}

func (s *ApplyComponentsRequestChildrenList) SetReadme(v string) *ApplyComponentsRequestChildrenList {
	s.Readme = &v
	return s
}

func (s *ApplyComponentsRequestChildrenList) SetResources(v string) *ApplyComponentsRequestChildrenList {
	s.Resources = &v
	return s
}

func (s *ApplyComponentsRequestChildrenList) SetSingleton(v bool) *ApplyComponentsRequestChildrenList {
	s.Singleton = &v
	return s
}

func (s *ApplyComponentsRequestChildrenList) SetVersion(v string) *ApplyComponentsRequestChildrenList {
	s.Version = &v
	return s
}

type ApplyComponentsRequestChildrenListPlatforms struct {
	Architecture *string `json:"architecture,omitempty" xml:"architecture,omitempty"`
	Os           *string `json:"os,omitempty" xml:"os,omitempty"`
}

func (s ApplyComponentsRequestChildrenListPlatforms) String() string {
	return tea.Prettify(s)
}

func (s ApplyComponentsRequestChildrenListPlatforms) GoString() string {
	return s.String()
}

func (s *ApplyComponentsRequestChildrenListPlatforms) SetArchitecture(v string) *ApplyComponentsRequestChildrenListPlatforms {
	s.Architecture = &v
	return s
}

func (s *ApplyComponentsRequestChildrenListPlatforms) SetOs(v string) *ApplyComponentsRequestChildrenListPlatforms {
	s.Os = &v
	return s
}

type ApplyComponentsResponseBody struct {
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	ErrMsg  *string `json:"errMsg,omitempty" xml:"errMsg,omitempty"`
	Status  *string `json:"status,omitempty" xml:"status,omitempty"`
	Success *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ApplyComponentsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ApplyComponentsResponseBody) GoString() string {
	return s.String()
}

func (s *ApplyComponentsResponseBody) SetErrCode(v string) *ApplyComponentsResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ApplyComponentsResponseBody) SetErrMsg(v string) *ApplyComponentsResponseBody {
	s.ErrMsg = &v
	return s
}

func (s *ApplyComponentsResponseBody) SetStatus(v string) *ApplyComponentsResponseBody {
	s.Status = &v
	return s
}

func (s *ApplyComponentsResponseBody) SetSuccess(v bool) *ApplyComponentsResponseBody {
	s.Success = &v
	return s
}

type ApplyComponentsResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ApplyComponentsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ApplyComponentsResponse) String() string {
	return tea.Prettify(s)
}

func (s ApplyComponentsResponse) GoString() string {
	return s.String()
}

func (s *ApplyComponentsResponse) SetHeaders(v map[string]*string) *ApplyComponentsResponse {
	s.Headers = v
	return s
}

func (s *ApplyComponentsResponse) SetBody(v *ApplyComponentsResponseBody) *ApplyComponentsResponse {
	s.Body = v
	return s
}

type CreatePackageConfigResponseBody struct {
	Data    *CreatePackageConfigResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	ErrCode *string                              `json:"errCode,omitempty" xml:"errCode,omitempty"`
	ErrMsg  *string                              `json:"errMsg,omitempty" xml:"errMsg,omitempty"`
	Success *bool                                `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreatePackageConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreatePackageConfigResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePackageConfigResponseBody) SetData(v *CreatePackageConfigResponseBodyData) *CreatePackageConfigResponseBody {
	s.Data = v
	return s
}

func (s *CreatePackageConfigResponseBody) SetErrCode(v string) *CreatePackageConfigResponseBody {
	s.ErrCode = &v
	return s
}

func (s *CreatePackageConfigResponseBody) SetErrMsg(v string) *CreatePackageConfigResponseBody {
	s.ErrMsg = &v
	return s
}

func (s *CreatePackageConfigResponseBody) SetSuccess(v bool) *CreatePackageConfigResponseBody {
	s.Success = &v
	return s
}

type CreatePackageConfigResponseBodyData struct {
	Uid *string `json:"uid,omitempty" xml:"uid,omitempty"`
}

func (s CreatePackageConfigResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreatePackageConfigResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreatePackageConfigResponseBodyData) SetUid(v string) *CreatePackageConfigResponseBodyData {
	s.Uid = &v
	return s
}

type CreatePackageConfigResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreatePackageConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreatePackageConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s CreatePackageConfigResponse) GoString() string {
	return s.String()
}

func (s *CreatePackageConfigResponse) SetHeaders(v map[string]*string) *CreatePackageConfigResponse {
	s.Headers = v
	return s
}

func (s *CreatePackageConfigResponse) SetBody(v *CreatePackageConfigResponseBody) *CreatePackageConfigResponse {
	s.Body = v
	return s
}

type AddProductComponentRequest struct {
	ReleaseName *string `json:"releaseName,omitempty" xml:"releaseName,omitempty"`
}

func (s AddProductComponentRequest) String() string {
	return tea.Prettify(s)
}

func (s AddProductComponentRequest) GoString() string {
	return s.String()
}

func (s *AddProductComponentRequest) SetReleaseName(v string) *AddProductComponentRequest {
	s.ReleaseName = &v
	return s
}

type AddProductComponentResponseBody struct {
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	ErrMsg  *string `json:"errMsg,omitempty" xml:"errMsg,omitempty"`
	Status  *string `json:"status,omitempty" xml:"status,omitempty"`
	Success *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s AddProductComponentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddProductComponentResponseBody) GoString() string {
	return s.String()
}

func (s *AddProductComponentResponseBody) SetErrCode(v string) *AddProductComponentResponseBody {
	s.ErrCode = &v
	return s
}

func (s *AddProductComponentResponseBody) SetErrMsg(v string) *AddProductComponentResponseBody {
	s.ErrMsg = &v
	return s
}

func (s *AddProductComponentResponseBody) SetStatus(v string) *AddProductComponentResponseBody {
	s.Status = &v
	return s
}

func (s *AddProductComponentResponseBody) SetSuccess(v bool) *AddProductComponentResponseBody {
	s.Success = &v
	return s
}

type AddProductComponentResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddProductComponentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddProductComponentResponse) String() string {
	return tea.Prettify(s)
}

func (s AddProductComponentResponse) GoString() string {
	return s.String()
}

func (s *AddProductComponentResponse) SetHeaders(v map[string]*string) *AddProductComponentResponse {
	s.Headers = v
	return s
}

func (s *AddProductComponentResponse) SetBody(v *AddProductComponentResponseBody) *AddProductComponentResponse {
	s.Body = v
	return s
}

type DeleteSnapshotResponseBody struct {
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	ErrMsg  *string `json:"errMsg,omitempty" xml:"errMsg,omitempty"`
	Success *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeleteSnapshotResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteSnapshotResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSnapshotResponseBody) SetErrCode(v string) *DeleteSnapshotResponseBody {
	s.ErrCode = &v
	return s
}

func (s *DeleteSnapshotResponseBody) SetErrMsg(v string) *DeleteSnapshotResponseBody {
	s.ErrMsg = &v
	return s
}

func (s *DeleteSnapshotResponseBody) SetSuccess(v bool) *DeleteSnapshotResponseBody {
	s.Success = &v
	return s
}

type DeleteSnapshotResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteSnapshotResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteSnapshotResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteSnapshotResponse) GoString() string {
	return s.String()
}

func (s *DeleteSnapshotResponse) SetHeaders(v map[string]*string) *DeleteSnapshotResponse {
	s.Headers = v
	return s
}

func (s *DeleteSnapshotResponse) SetBody(v *DeleteSnapshotResponseBody) *DeleteSnapshotResponse {
	s.Body = v
	return s
}

type ListEnvironmentsWithSnapshotResponseBody struct {
	Data    *ListEnvironmentsWithSnapshotResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	ErrCode *string                                       `json:"errCode,omitempty" xml:"errCode,omitempty"`
	ErrMsg  *string                                       `json:"errMsg,omitempty" xml:"errMsg,omitempty"`
	Success *bool                                         `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ListEnvironmentsWithSnapshotResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListEnvironmentsWithSnapshotResponseBody) GoString() string {
	return s.String()
}

func (s *ListEnvironmentsWithSnapshotResponseBody) SetData(v *ListEnvironmentsWithSnapshotResponseBodyData) *ListEnvironmentsWithSnapshotResponseBody {
	s.Data = v
	return s
}

func (s *ListEnvironmentsWithSnapshotResponseBody) SetErrCode(v string) *ListEnvironmentsWithSnapshotResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ListEnvironmentsWithSnapshotResponseBody) SetErrMsg(v string) *ListEnvironmentsWithSnapshotResponseBody {
	s.ErrMsg = &v
	return s
}

func (s *ListEnvironmentsWithSnapshotResponseBody) SetSuccess(v bool) *ListEnvironmentsWithSnapshotResponseBody {
	s.Success = &v
	return s
}

type ListEnvironmentsWithSnapshotResponseBodyData struct {
	List     []*ListEnvironmentsWithSnapshotResponseBodyDataList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	PageNum  *int32                                              `json:"pageNum,omitempty" xml:"pageNum,omitempty"`
	PageSize *int32                                              `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	Total    *int32                                              `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListEnvironmentsWithSnapshotResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListEnvironmentsWithSnapshotResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListEnvironmentsWithSnapshotResponseBodyData) SetList(v []*ListEnvironmentsWithSnapshotResponseBodyDataList) *ListEnvironmentsWithSnapshotResponseBodyData {
	s.List = v
	return s
}

func (s *ListEnvironmentsWithSnapshotResponseBodyData) SetPageNum(v int32) *ListEnvironmentsWithSnapshotResponseBodyData {
	s.PageNum = &v
	return s
}

func (s *ListEnvironmentsWithSnapshotResponseBodyData) SetPageSize(v int32) *ListEnvironmentsWithSnapshotResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListEnvironmentsWithSnapshotResponseBodyData) SetTotal(v int32) *ListEnvironmentsWithSnapshotResponseBodyData {
	s.Total = &v
	return s
}

type ListEnvironmentsWithSnapshotResponseBodyDataList struct {
	CreatedAt         *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	Description       *string `json:"description,omitempty" xml:"description,omitempty"`
	Id                *int32  `json:"id,omitempty" xml:"id,omitempty"`
	InstanceStatus    *string `json:"instanceStatus,omitempty" xml:"instanceStatus,omitempty"`
	Name              *string `json:"name,omitempty" xml:"name,omitempty"`
	ProductName       *string `json:"productName,omitempty" xml:"productName,omitempty"`
	ProductVersion    *string `json:"productVersion,omitempty" xml:"productVersion,omitempty"`
	ProductVersionUID *string `json:"productVersionUID,omitempty" xml:"productVersionUID,omitempty"`
	Uid               *string `json:"uid,omitempty" xml:"uid,omitempty"`
	VendorType        *string `json:"vendorType,omitempty" xml:"vendorType,omitempty"`
}

func (s ListEnvironmentsWithSnapshotResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s ListEnvironmentsWithSnapshotResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *ListEnvironmentsWithSnapshotResponseBodyDataList) SetCreatedAt(v string) *ListEnvironmentsWithSnapshotResponseBodyDataList {
	s.CreatedAt = &v
	return s
}

func (s *ListEnvironmentsWithSnapshotResponseBodyDataList) SetDescription(v string) *ListEnvironmentsWithSnapshotResponseBodyDataList {
	s.Description = &v
	return s
}

func (s *ListEnvironmentsWithSnapshotResponseBodyDataList) SetId(v int32) *ListEnvironmentsWithSnapshotResponseBodyDataList {
	s.Id = &v
	return s
}

func (s *ListEnvironmentsWithSnapshotResponseBodyDataList) SetInstanceStatus(v string) *ListEnvironmentsWithSnapshotResponseBodyDataList {
	s.InstanceStatus = &v
	return s
}

func (s *ListEnvironmentsWithSnapshotResponseBodyDataList) SetName(v string) *ListEnvironmentsWithSnapshotResponseBodyDataList {
	s.Name = &v
	return s
}

func (s *ListEnvironmentsWithSnapshotResponseBodyDataList) SetProductName(v string) *ListEnvironmentsWithSnapshotResponseBodyDataList {
	s.ProductName = &v
	return s
}

func (s *ListEnvironmentsWithSnapshotResponseBodyDataList) SetProductVersion(v string) *ListEnvironmentsWithSnapshotResponseBodyDataList {
	s.ProductVersion = &v
	return s
}

func (s *ListEnvironmentsWithSnapshotResponseBodyDataList) SetProductVersionUID(v string) *ListEnvironmentsWithSnapshotResponseBodyDataList {
	s.ProductVersionUID = &v
	return s
}

func (s *ListEnvironmentsWithSnapshotResponseBodyDataList) SetUid(v string) *ListEnvironmentsWithSnapshotResponseBodyDataList {
	s.Uid = &v
	return s
}

func (s *ListEnvironmentsWithSnapshotResponseBodyDataList) SetVendorType(v string) *ListEnvironmentsWithSnapshotResponseBodyDataList {
	s.VendorType = &v
	return s
}

type ListEnvironmentsWithSnapshotResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListEnvironmentsWithSnapshotResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListEnvironmentsWithSnapshotResponse) String() string {
	return tea.Prettify(s)
}

func (s ListEnvironmentsWithSnapshotResponse) GoString() string {
	return s.String()
}

func (s *ListEnvironmentsWithSnapshotResponse) SetHeaders(v map[string]*string) *ListEnvironmentsWithSnapshotResponse {
	s.Headers = v
	return s
}

func (s *ListEnvironmentsWithSnapshotResponse) SetBody(v *ListEnvironmentsWithSnapshotResponseBody) *ListEnvironmentsWithSnapshotResponse {
	s.Body = v
	return s
}

type GetEnvironmentNodeResponseBody struct {
	Data    *GetEnvironmentNodeResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	ErrCode *string                             `json:"errCode,omitempty" xml:"errCode,omitempty"`
	ErrMsg  *string                             `json:"errMsg,omitempty" xml:"errMsg,omitempty"`
	Success *bool                               `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetEnvironmentNodeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetEnvironmentNodeResponseBody) GoString() string {
	return s.String()
}

func (s *GetEnvironmentNodeResponseBody) SetData(v *GetEnvironmentNodeResponseBodyData) *GetEnvironmentNodeResponseBody {
	s.Data = v
	return s
}

func (s *GetEnvironmentNodeResponseBody) SetErrCode(v string) *GetEnvironmentNodeResponseBody {
	s.ErrCode = &v
	return s
}

func (s *GetEnvironmentNodeResponseBody) SetErrMsg(v string) *GetEnvironmentNodeResponseBody {
	s.ErrMsg = &v
	return s
}

func (s *GetEnvironmentNodeResponseBody) SetSuccess(v bool) *GetEnvironmentNodeResponseBody {
	s.Success = &v
	return s
}

type GetEnvironmentNodeResponseBodyData struct {
	ClusterId      *string `json:"clusterId,omitempty" xml:"clusterId,omitempty"`
	CreatedAt      *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty"`
	InstanceList   *string `json:"instanceList,omitempty" xml:"instanceList,omitempty"`
	Name           *string `json:"name,omitempty" xml:"name,omitempty"`
	ProductName    *string `json:"productName,omitempty" xml:"productName,omitempty"`
	ProductVersion *string `json:"productVersion,omitempty" xml:"productVersion,omitempty"`
	Uid            *string `json:"uid,omitempty" xml:"uid,omitempty"`
	VendorConfig   *string `json:"vendorConfig,omitempty" xml:"vendorConfig,omitempty"`
	VendorType     *string `json:"vendorType,omitempty" xml:"vendorType,omitempty"`
}

func (s GetEnvironmentNodeResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetEnvironmentNodeResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetEnvironmentNodeResponseBodyData) SetClusterId(v string) *GetEnvironmentNodeResponseBodyData {
	s.ClusterId = &v
	return s
}

func (s *GetEnvironmentNodeResponseBodyData) SetCreatedAt(v string) *GetEnvironmentNodeResponseBodyData {
	s.CreatedAt = &v
	return s
}

func (s *GetEnvironmentNodeResponseBodyData) SetDescription(v string) *GetEnvironmentNodeResponseBodyData {
	s.Description = &v
	return s
}

func (s *GetEnvironmentNodeResponseBodyData) SetInstanceList(v string) *GetEnvironmentNodeResponseBodyData {
	s.InstanceList = &v
	return s
}

func (s *GetEnvironmentNodeResponseBodyData) SetName(v string) *GetEnvironmentNodeResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetEnvironmentNodeResponseBodyData) SetProductName(v string) *GetEnvironmentNodeResponseBodyData {
	s.ProductName = &v
	return s
}

func (s *GetEnvironmentNodeResponseBodyData) SetProductVersion(v string) *GetEnvironmentNodeResponseBodyData {
	s.ProductVersion = &v
	return s
}

func (s *GetEnvironmentNodeResponseBodyData) SetUid(v string) *GetEnvironmentNodeResponseBodyData {
	s.Uid = &v
	return s
}

func (s *GetEnvironmentNodeResponseBodyData) SetVendorConfig(v string) *GetEnvironmentNodeResponseBodyData {
	s.VendorConfig = &v
	return s
}

func (s *GetEnvironmentNodeResponseBodyData) SetVendorType(v string) *GetEnvironmentNodeResponseBodyData {
	s.VendorType = &v
	return s
}

type GetEnvironmentNodeResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetEnvironmentNodeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetEnvironmentNodeResponse) String() string {
	return tea.Prettify(s)
}

func (s GetEnvironmentNodeResponse) GoString() string {
	return s.String()
}

func (s *GetEnvironmentNodeResponse) SetHeaders(v map[string]*string) *GetEnvironmentNodeResponse {
	s.Headers = v
	return s
}

func (s *GetEnvironmentNodeResponse) SetBody(v *GetEnvironmentNodeResponseBody) *GetEnvironmentNodeResponse {
	s.Body = v
	return s
}

type UpdateSnapshotRequest struct {
	Description        *string `json:"description,omitempty" xml:"description,omitempty"`
	ProductName        *string `json:"productName,omitempty" xml:"productName,omitempty"`
	ProductVersion     *string `json:"productVersion,omitempty" xml:"productVersion,omitempty"`
	ProductVersionDesc *string `json:"productVersionDesc,omitempty" xml:"productVersionDesc,omitempty"`
	UpdateKey          *string `json:"updateKey,omitempty" xml:"updateKey,omitempty"`
}

func (s UpdateSnapshotRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateSnapshotRequest) GoString() string {
	return s.String()
}

func (s *UpdateSnapshotRequest) SetDescription(v string) *UpdateSnapshotRequest {
	s.Description = &v
	return s
}

func (s *UpdateSnapshotRequest) SetProductName(v string) *UpdateSnapshotRequest {
	s.ProductName = &v
	return s
}

func (s *UpdateSnapshotRequest) SetProductVersion(v string) *UpdateSnapshotRequest {
	s.ProductVersion = &v
	return s
}

func (s *UpdateSnapshotRequest) SetProductVersionDesc(v string) *UpdateSnapshotRequest {
	s.ProductVersionDesc = &v
	return s
}

func (s *UpdateSnapshotRequest) SetUpdateKey(v string) *UpdateSnapshotRequest {
	s.UpdateKey = &v
	return s
}

type UpdateSnapshotResponseBody struct {
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	ErrMsg  *string `json:"errMsg,omitempty" xml:"errMsg,omitempty"`
	Success *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateSnapshotResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateSnapshotResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSnapshotResponseBody) SetErrCode(v string) *UpdateSnapshotResponseBody {
	s.ErrCode = &v
	return s
}

func (s *UpdateSnapshotResponseBody) SetErrMsg(v string) *UpdateSnapshotResponseBody {
	s.ErrMsg = &v
	return s
}

func (s *UpdateSnapshotResponseBody) SetSuccess(v bool) *UpdateSnapshotResponseBody {
	s.Success = &v
	return s
}

type UpdateSnapshotResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateSnapshotResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateSnapshotResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateSnapshotResponse) GoString() string {
	return s.String()
}

func (s *UpdateSnapshotResponse) SetHeaders(v map[string]*string) *UpdateSnapshotResponse {
	s.Headers = v
	return s
}

func (s *UpdateSnapshotResponse) SetBody(v *UpdateSnapshotResponseBody) *UpdateSnapshotResponse {
	s.Body = v
	return s
}

type CreateEnvironmentAndGenerateVendorConfigHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	ClientToken   *string            `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s CreateEnvironmentAndGenerateVendorConfigHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateEnvironmentAndGenerateVendorConfigHeaders) GoString() string {
	return s.String()
}

func (s *CreateEnvironmentAndGenerateVendorConfigHeaders) SetCommonHeaders(v map[string]*string) *CreateEnvironmentAndGenerateVendorConfigHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateEnvironmentAndGenerateVendorConfigHeaders) SetClientToken(v string) *CreateEnvironmentAndGenerateVendorConfigHeaders {
	s.ClientToken = &v
	return s
}

type CreateEnvironmentAndGenerateVendorConfigRequest struct {
	EnvUID            *string                                                  `json:"envUID,omitempty" xml:"envUID,omitempty"`
	Platform          *CreateEnvironmentAndGenerateVendorConfigRequestPlatform `json:"platform,omitempty" xml:"platform,omitempty" type:"Struct"`
	ProductName       *string                                                  `json:"productName,omitempty" xml:"productName,omitempty"`
	ProductUID        *string                                                  `json:"productUID,omitempty" xml:"productUID,omitempty"`
	ProductVersion    *string                                                  `json:"productVersion,omitempty" xml:"productVersion,omitempty"`
	ProductVersionUID *string                                                  `json:"productVersionUID,omitempty" xml:"productVersionUID,omitempty"`
	VendorType        *string                                                  `json:"vendorType,omitempty" xml:"vendorType,omitempty"`
}

func (s CreateEnvironmentAndGenerateVendorConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateEnvironmentAndGenerateVendorConfigRequest) GoString() string {
	return s.String()
}

func (s *CreateEnvironmentAndGenerateVendorConfigRequest) SetEnvUID(v string) *CreateEnvironmentAndGenerateVendorConfigRequest {
	s.EnvUID = &v
	return s
}

func (s *CreateEnvironmentAndGenerateVendorConfigRequest) SetPlatform(v *CreateEnvironmentAndGenerateVendorConfigRequestPlatform) *CreateEnvironmentAndGenerateVendorConfigRequest {
	s.Platform = v
	return s
}

func (s *CreateEnvironmentAndGenerateVendorConfigRequest) SetProductName(v string) *CreateEnvironmentAndGenerateVendorConfigRequest {
	s.ProductName = &v
	return s
}

func (s *CreateEnvironmentAndGenerateVendorConfigRequest) SetProductUID(v string) *CreateEnvironmentAndGenerateVendorConfigRequest {
	s.ProductUID = &v
	return s
}

func (s *CreateEnvironmentAndGenerateVendorConfigRequest) SetProductVersion(v string) *CreateEnvironmentAndGenerateVendorConfigRequest {
	s.ProductVersion = &v
	return s
}

func (s *CreateEnvironmentAndGenerateVendorConfigRequest) SetProductVersionUID(v string) *CreateEnvironmentAndGenerateVendorConfigRequest {
	s.ProductVersionUID = &v
	return s
}

func (s *CreateEnvironmentAndGenerateVendorConfigRequest) SetVendorType(v string) *CreateEnvironmentAndGenerateVendorConfigRequest {
	s.VendorType = &v
	return s
}

type CreateEnvironmentAndGenerateVendorConfigRequestPlatform struct {
	Architecture *string `json:"architecture,omitempty" xml:"architecture,omitempty"`
	Os           *string `json:"os,omitempty" xml:"os,omitempty"`
}

func (s CreateEnvironmentAndGenerateVendorConfigRequestPlatform) String() string {
	return tea.Prettify(s)
}

func (s CreateEnvironmentAndGenerateVendorConfigRequestPlatform) GoString() string {
	return s.String()
}

func (s *CreateEnvironmentAndGenerateVendorConfigRequestPlatform) SetArchitecture(v string) *CreateEnvironmentAndGenerateVendorConfigRequestPlatform {
	s.Architecture = &v
	return s
}

func (s *CreateEnvironmentAndGenerateVendorConfigRequestPlatform) SetOs(v string) *CreateEnvironmentAndGenerateVendorConfigRequestPlatform {
	s.Os = &v
	return s
}

type CreateEnvironmentAndGenerateVendorConfigResponseBody struct {
	Data    *CreateEnvironmentAndGenerateVendorConfigResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	ErrCode *string                                                   `json:"errCode,omitempty" xml:"errCode,omitempty"`
	ErrMsg  *string                                                   `json:"errMsg,omitempty" xml:"errMsg,omitempty"`
	Success *bool                                                     `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreateEnvironmentAndGenerateVendorConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateEnvironmentAndGenerateVendorConfigResponseBody) GoString() string {
	return s.String()
}

func (s *CreateEnvironmentAndGenerateVendorConfigResponseBody) SetData(v *CreateEnvironmentAndGenerateVendorConfigResponseBodyData) *CreateEnvironmentAndGenerateVendorConfigResponseBody {
	s.Data = v
	return s
}

func (s *CreateEnvironmentAndGenerateVendorConfigResponseBody) SetErrCode(v string) *CreateEnvironmentAndGenerateVendorConfigResponseBody {
	s.ErrCode = &v
	return s
}

func (s *CreateEnvironmentAndGenerateVendorConfigResponseBody) SetErrMsg(v string) *CreateEnvironmentAndGenerateVendorConfigResponseBody {
	s.ErrMsg = &v
	return s
}

func (s *CreateEnvironmentAndGenerateVendorConfigResponseBody) SetSuccess(v bool) *CreateEnvironmentAndGenerateVendorConfigResponseBody {
	s.Success = &v
	return s
}

type CreateEnvironmentAndGenerateVendorConfigResponseBodyData struct {
	EnvUID       *string `json:"envUID,omitempty" xml:"envUID,omitempty"`
	VendorConfig *string `json:"vendorConfig,omitempty" xml:"vendorConfig,omitempty"`
}

func (s CreateEnvironmentAndGenerateVendorConfigResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateEnvironmentAndGenerateVendorConfigResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateEnvironmentAndGenerateVendorConfigResponseBodyData) SetEnvUID(v string) *CreateEnvironmentAndGenerateVendorConfigResponseBodyData {
	s.EnvUID = &v
	return s
}

func (s *CreateEnvironmentAndGenerateVendorConfigResponseBodyData) SetVendorConfig(v string) *CreateEnvironmentAndGenerateVendorConfigResponseBodyData {
	s.VendorConfig = &v
	return s
}

type CreateEnvironmentAndGenerateVendorConfigResponse struct {
	Headers map[string]*string                                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateEnvironmentAndGenerateVendorConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateEnvironmentAndGenerateVendorConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateEnvironmentAndGenerateVendorConfigResponse) GoString() string {
	return s.String()
}

func (s *CreateEnvironmentAndGenerateVendorConfigResponse) SetHeaders(v map[string]*string) *CreateEnvironmentAndGenerateVendorConfigResponse {
	s.Headers = v
	return s
}

func (s *CreateEnvironmentAndGenerateVendorConfigResponse) SetBody(v *CreateEnvironmentAndGenerateVendorConfigResponseBody) *CreateEnvironmentAndGenerateVendorConfigResponse {
	s.Body = v
	return s
}

type CreateEnvironmentSnapshotRequest struct {
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	Name        *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s CreateEnvironmentSnapshotRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateEnvironmentSnapshotRequest) GoString() string {
	return s.String()
}

func (s *CreateEnvironmentSnapshotRequest) SetDescription(v string) *CreateEnvironmentSnapshotRequest {
	s.Description = &v
	return s
}

func (s *CreateEnvironmentSnapshotRequest) SetName(v string) *CreateEnvironmentSnapshotRequest {
	s.Name = &v
	return s
}

type CreateEnvironmentSnapshotResponseBody struct {
	Data    *CreateEnvironmentSnapshotResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	ErrCode *string                                    `json:"errCode,omitempty" xml:"errCode,omitempty"`
	ErrMsg  *string                                    `json:"errMsg,omitempty" xml:"errMsg,omitempty"`
	Success *bool                                      `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreateEnvironmentSnapshotResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateEnvironmentSnapshotResponseBody) GoString() string {
	return s.String()
}

func (s *CreateEnvironmentSnapshotResponseBody) SetData(v *CreateEnvironmentSnapshotResponseBodyData) *CreateEnvironmentSnapshotResponseBody {
	s.Data = v
	return s
}

func (s *CreateEnvironmentSnapshotResponseBody) SetErrCode(v string) *CreateEnvironmentSnapshotResponseBody {
	s.ErrCode = &v
	return s
}

func (s *CreateEnvironmentSnapshotResponseBody) SetErrMsg(v string) *CreateEnvironmentSnapshotResponseBody {
	s.ErrMsg = &v
	return s
}

func (s *CreateEnvironmentSnapshotResponseBody) SetSuccess(v bool) *CreateEnvironmentSnapshotResponseBody {
	s.Success = &v
	return s
}

type CreateEnvironmentSnapshotResponseBodyData struct {
	Uid *string `json:"uid,omitempty" xml:"uid,omitempty"`
}

func (s CreateEnvironmentSnapshotResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateEnvironmentSnapshotResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateEnvironmentSnapshotResponseBodyData) SetUid(v string) *CreateEnvironmentSnapshotResponseBodyData {
	s.Uid = &v
	return s
}

type CreateEnvironmentSnapshotResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateEnvironmentSnapshotResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateEnvironmentSnapshotResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateEnvironmentSnapshotResponse) GoString() string {
	return s.String()
}

func (s *CreateEnvironmentSnapshotResponse) SetHeaders(v map[string]*string) *CreateEnvironmentSnapshotResponse {
	s.Headers = v
	return s
}

func (s *CreateEnvironmentSnapshotResponse) SetBody(v *CreateEnvironmentSnapshotResponseBody) *CreateEnvironmentSnapshotResponse {
	s.Body = v
	return s
}

type InitSnapshotInstanceResponseBody struct {
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	ErrMsg  *string `json:"errMsg,omitempty" xml:"errMsg,omitempty"`
	Success *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s InitSnapshotInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s InitSnapshotInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *InitSnapshotInstanceResponseBody) SetErrCode(v string) *InitSnapshotInstanceResponseBody {
	s.ErrCode = &v
	return s
}

func (s *InitSnapshotInstanceResponseBody) SetErrMsg(v string) *InitSnapshotInstanceResponseBody {
	s.ErrMsg = &v
	return s
}

func (s *InitSnapshotInstanceResponseBody) SetSuccess(v bool) *InitSnapshotInstanceResponseBody {
	s.Success = &v
	return s
}

type InitSnapshotInstanceResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *InitSnapshotInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s InitSnapshotInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s InitSnapshotInstanceResponse) GoString() string {
	return s.String()
}

func (s *InitSnapshotInstanceResponse) SetHeaders(v map[string]*string) *InitSnapshotInstanceResponse {
	s.Headers = v
	return s
}

func (s *InitSnapshotInstanceResponse) SetBody(v *InitSnapshotInstanceResponseBody) *InitSnapshotInstanceResponse {
	s.Body = v
	return s
}

type UpdateProductVersionRelatedFoundationVersionRequest struct {
	FoundationVersionUID *string `json:"foundationVersionUID,omitempty" xml:"foundationVersionUID,omitempty"`
}

func (s UpdateProductVersionRelatedFoundationVersionRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateProductVersionRelatedFoundationVersionRequest) GoString() string {
	return s.String()
}

func (s *UpdateProductVersionRelatedFoundationVersionRequest) SetFoundationVersionUID(v string) *UpdateProductVersionRelatedFoundationVersionRequest {
	s.FoundationVersionUID = &v
	return s
}

type UpdateProductVersionRelatedFoundationVersionResponseBody struct {
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	ErrMsg  *string `json:"errMsg,omitempty" xml:"errMsg,omitempty"`
	Success *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateProductVersionRelatedFoundationVersionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateProductVersionRelatedFoundationVersionResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateProductVersionRelatedFoundationVersionResponseBody) SetErrCode(v string) *UpdateProductVersionRelatedFoundationVersionResponseBody {
	s.ErrCode = &v
	return s
}

func (s *UpdateProductVersionRelatedFoundationVersionResponseBody) SetErrMsg(v string) *UpdateProductVersionRelatedFoundationVersionResponseBody {
	s.ErrMsg = &v
	return s
}

func (s *UpdateProductVersionRelatedFoundationVersionResponseBody) SetSuccess(v bool) *UpdateProductVersionRelatedFoundationVersionResponseBody {
	s.Success = &v
	return s
}

type UpdateProductVersionRelatedFoundationVersionResponse struct {
	Headers map[string]*string                                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateProductVersionRelatedFoundationVersionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateProductVersionRelatedFoundationVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateProductVersionRelatedFoundationVersionResponse) GoString() string {
	return s.String()
}

func (s *UpdateProductVersionRelatedFoundationVersionResponse) SetHeaders(v map[string]*string) *UpdateProductVersionRelatedFoundationVersionResponse {
	s.Headers = v
	return s
}

func (s *UpdateProductVersionRelatedFoundationVersionResponse) SetBody(v *UpdateProductVersionRelatedFoundationVersionResponseBody) *UpdateProductVersionRelatedFoundationVersionResponse {
	s.Body = v
	return s
}

type ListEnvironmentParamsRequest struct {
	PageNum  *int32  `json:"pageNum,omitempty" xml:"pageNum,omitempty"`
	PageSize *int32  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	Name     *string `json:"name,omitempty" xml:"name,omitempty"`
	Fuzzy    *string `json:"fuzzy,omitempty" xml:"fuzzy,omitempty"`
}

func (s ListEnvironmentParamsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListEnvironmentParamsRequest) GoString() string {
	return s.String()
}

func (s *ListEnvironmentParamsRequest) SetPageNum(v int32) *ListEnvironmentParamsRequest {
	s.PageNum = &v
	return s
}

func (s *ListEnvironmentParamsRequest) SetPageSize(v int32) *ListEnvironmentParamsRequest {
	s.PageSize = &v
	return s
}

func (s *ListEnvironmentParamsRequest) SetName(v string) *ListEnvironmentParamsRequest {
	s.Name = &v
	return s
}

func (s *ListEnvironmentParamsRequest) SetFuzzy(v string) *ListEnvironmentParamsRequest {
	s.Fuzzy = &v
	return s
}

type ListEnvironmentParamsResponseBody struct {
	Data    *ListEnvironmentParamsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	ErrCode *string                                `json:"errCode,omitempty" xml:"errCode,omitempty"`
	ErrMsg  *string                                `json:"errMsg,omitempty" xml:"errMsg,omitempty"`
	Success *bool                                  `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ListEnvironmentParamsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListEnvironmentParamsResponseBody) GoString() string {
	return s.String()
}

func (s *ListEnvironmentParamsResponseBody) SetData(v *ListEnvironmentParamsResponseBodyData) *ListEnvironmentParamsResponseBody {
	s.Data = v
	return s
}

func (s *ListEnvironmentParamsResponseBody) SetErrCode(v string) *ListEnvironmentParamsResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ListEnvironmentParamsResponseBody) SetErrMsg(v string) *ListEnvironmentParamsResponseBody {
	s.ErrMsg = &v
	return s
}

func (s *ListEnvironmentParamsResponseBody) SetSuccess(v bool) *ListEnvironmentParamsResponseBody {
	s.Success = &v
	return s
}

type ListEnvironmentParamsResponseBodyData struct {
	List     []*ListEnvironmentParamsResponseBodyDataList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	PageNum  *int32                                       `json:"pageNum,omitempty" xml:"pageNum,omitempty"`
	PageSize *int32                                       `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	Total    *int32                                       `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListEnvironmentParamsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListEnvironmentParamsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListEnvironmentParamsResponseBodyData) SetList(v []*ListEnvironmentParamsResponseBodyDataList) *ListEnvironmentParamsResponseBodyData {
	s.List = v
	return s
}

func (s *ListEnvironmentParamsResponseBodyData) SetPageNum(v int32) *ListEnvironmentParamsResponseBodyData {
	s.PageNum = &v
	return s
}

func (s *ListEnvironmentParamsResponseBodyData) SetPageSize(v int32) *ListEnvironmentParamsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListEnvironmentParamsResponseBodyData) SetTotal(v int32) *ListEnvironmentParamsResponseBodyData {
	s.Total = &v
	return s
}

type ListEnvironmentParamsResponseBodyDataList struct {
	CreatedAt         *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	Description       *string `json:"description,omitempty" xml:"description,omitempty"`
	Id                *int32  `json:"id,omitempty" xml:"id,omitempty"`
	Name              *string `json:"name,omitempty" xml:"name,omitempty"`
	ProductName       *string `json:"productName,omitempty" xml:"productName,omitempty"`
	ProductVersion    *string `json:"productVersion,omitempty" xml:"productVersion,omitempty"`
	ProductVersionUID *string `json:"productVersionUID,omitempty" xml:"productVersionUID,omitempty"`
	Uid               *string `json:"uid,omitempty" xml:"uid,omitempty"`
	VendorType        *string `json:"vendorType,omitempty" xml:"vendorType,omitempty"`
}

func (s ListEnvironmentParamsResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s ListEnvironmentParamsResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *ListEnvironmentParamsResponseBodyDataList) SetCreatedAt(v string) *ListEnvironmentParamsResponseBodyDataList {
	s.CreatedAt = &v
	return s
}

func (s *ListEnvironmentParamsResponseBodyDataList) SetDescription(v string) *ListEnvironmentParamsResponseBodyDataList {
	s.Description = &v
	return s
}

func (s *ListEnvironmentParamsResponseBodyDataList) SetId(v int32) *ListEnvironmentParamsResponseBodyDataList {
	s.Id = &v
	return s
}

func (s *ListEnvironmentParamsResponseBodyDataList) SetName(v string) *ListEnvironmentParamsResponseBodyDataList {
	s.Name = &v
	return s
}

func (s *ListEnvironmentParamsResponseBodyDataList) SetProductName(v string) *ListEnvironmentParamsResponseBodyDataList {
	s.ProductName = &v
	return s
}

func (s *ListEnvironmentParamsResponseBodyDataList) SetProductVersion(v string) *ListEnvironmentParamsResponseBodyDataList {
	s.ProductVersion = &v
	return s
}

func (s *ListEnvironmentParamsResponseBodyDataList) SetProductVersionUID(v string) *ListEnvironmentParamsResponseBodyDataList {
	s.ProductVersionUID = &v
	return s
}

func (s *ListEnvironmentParamsResponseBodyDataList) SetUid(v string) *ListEnvironmentParamsResponseBodyDataList {
	s.Uid = &v
	return s
}

func (s *ListEnvironmentParamsResponseBodyDataList) SetVendorType(v string) *ListEnvironmentParamsResponseBodyDataList {
	s.VendorType = &v
	return s
}

type ListEnvironmentParamsResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListEnvironmentParamsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListEnvironmentParamsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListEnvironmentParamsResponse) GoString() string {
	return s.String()
}

func (s *ListEnvironmentParamsResponse) SetHeaders(v map[string]*string) *ListEnvironmentParamsResponse {
	s.Headers = v
	return s
}

func (s *ListEnvironmentParamsResponse) SetBody(v *ListEnvironmentParamsResponseBody) *ListEnvironmentParamsResponse {
	s.Body = v
	return s
}

type GetFoundationVersionResponseBody struct {
	Data    []*FoundationVersion `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	ErrCode *string              `json:"errCode,omitempty" xml:"errCode,omitempty"`
	ErrMsg  *string              `json:"errMsg,omitempty" xml:"errMsg,omitempty"`
	Success *bool                `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetFoundationVersionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetFoundationVersionResponseBody) GoString() string {
	return s.String()
}

func (s *GetFoundationVersionResponseBody) SetData(v []*FoundationVersion) *GetFoundationVersionResponseBody {
	s.Data = v
	return s
}

func (s *GetFoundationVersionResponseBody) SetErrCode(v string) *GetFoundationVersionResponseBody {
	s.ErrCode = &v
	return s
}

func (s *GetFoundationVersionResponseBody) SetErrMsg(v string) *GetFoundationVersionResponseBody {
	s.ErrMsg = &v
	return s
}

func (s *GetFoundationVersionResponseBody) SetSuccess(v bool) *GetFoundationVersionResponseBody {
	s.Success = &v
	return s
}

type GetFoundationVersionResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetFoundationVersionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetFoundationVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s GetFoundationVersionResponse) GoString() string {
	return s.String()
}

func (s *GetFoundationVersionResponse) SetHeaders(v map[string]*string) *GetFoundationVersionResponse {
	s.Headers = v
	return s
}

func (s *GetFoundationVersionResponse) SetBody(v *GetFoundationVersionResponseBody) *GetFoundationVersionResponse {
	s.Body = v
	return s
}

type DeleteProductResponseBody struct {
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	Success    *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrCode    *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
}

func (s DeleteProductResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteProductResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteProductResponseBody) SetErrMessage(v string) *DeleteProductResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DeleteProductResponseBody) SetSuccess(v bool) *DeleteProductResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteProductResponseBody) SetErrCode(v string) *DeleteProductResponseBody {
	s.ErrCode = &v
	return s
}

type DeleteProductResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteProductResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteProductResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteProductResponse) GoString() string {
	return s.String()
}

func (s *DeleteProductResponse) SetHeaders(v map[string]*string) *DeleteProductResponse {
	s.Headers = v
	return s
}

func (s *DeleteProductResponse) SetBody(v *DeleteProductResponseBody) *DeleteProductResponse {
	s.Body = v
	return s
}

type UpdateEnvironmentRequest struct {
	Description  *string `json:"description,omitempty" xml:"description,omitempty"`
	Location     *string `json:"location,omitempty" xml:"location,omitempty"`
	VendorConfig *string `json:"vendorConfig,omitempty" xml:"vendorConfig,omitempty"`
}

func (s UpdateEnvironmentRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateEnvironmentRequest) GoString() string {
	return s.String()
}

func (s *UpdateEnvironmentRequest) SetDescription(v string) *UpdateEnvironmentRequest {
	s.Description = &v
	return s
}

func (s *UpdateEnvironmentRequest) SetLocation(v string) *UpdateEnvironmentRequest {
	s.Location = &v
	return s
}

func (s *UpdateEnvironmentRequest) SetVendorConfig(v string) *UpdateEnvironmentRequest {
	s.VendorConfig = &v
	return s
}

type UpdateEnvironmentResponseBody struct {
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	ErrMsg  *string `json:"errMsg,omitempty" xml:"errMsg,omitempty"`
	Success *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateEnvironmentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateEnvironmentResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateEnvironmentResponseBody) SetErrCode(v string) *UpdateEnvironmentResponseBody {
	s.ErrCode = &v
	return s
}

func (s *UpdateEnvironmentResponseBody) SetErrMsg(v string) *UpdateEnvironmentResponseBody {
	s.ErrMsg = &v
	return s
}

func (s *UpdateEnvironmentResponseBody) SetSuccess(v bool) *UpdateEnvironmentResponseBody {
	s.Success = &v
	return s
}

type UpdateEnvironmentResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateEnvironmentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateEnvironmentResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateEnvironmentResponse) GoString() string {
	return s.String()
}

func (s *UpdateEnvironmentResponse) SetHeaders(v map[string]*string) *UpdateEnvironmentResponse {
	s.Headers = v
	return s
}

func (s *UpdateEnvironmentResponse) SetBody(v *UpdateEnvironmentResponseBody) *UpdateEnvironmentResponse {
	s.Body = v
	return s
}

type GetEnvironmentPackageResponseBody struct {
	Data    *GetEnvironmentPackageResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	ErrCode *string                                `json:"errCode,omitempty" xml:"errCode,omitempty"`
	ErrMsg  *string                                `json:"errMsg,omitempty" xml:"errMsg,omitempty"`
	Success *bool                                  `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetEnvironmentPackageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetEnvironmentPackageResponseBody) GoString() string {
	return s.String()
}

func (s *GetEnvironmentPackageResponseBody) SetData(v *GetEnvironmentPackageResponseBodyData) *GetEnvironmentPackageResponseBody {
	s.Data = v
	return s
}

func (s *GetEnvironmentPackageResponseBody) SetErrCode(v string) *GetEnvironmentPackageResponseBody {
	s.ErrCode = &v
	return s
}

func (s *GetEnvironmentPackageResponseBody) SetErrMsg(v string) *GetEnvironmentPackageResponseBody {
	s.ErrMsg = &v
	return s
}

func (s *GetEnvironmentPackageResponseBody) SetSuccess(v bool) *GetEnvironmentPackageResponseBody {
	s.Success = &v
	return s
}

type GetEnvironmentPackageResponseBodyData struct {
	PackageURL *string `json:"packageURL,omitempty" xml:"packageURL,omitempty"`
}

func (s GetEnvironmentPackageResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetEnvironmentPackageResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetEnvironmentPackageResponseBodyData) SetPackageURL(v string) *GetEnvironmentPackageResponseBodyData {
	s.PackageURL = &v
	return s
}

type GetEnvironmentPackageResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetEnvironmentPackageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetEnvironmentPackageResponse) String() string {
	return tea.Prettify(s)
}

func (s GetEnvironmentPackageResponse) GoString() string {
	return s.String()
}

func (s *GetEnvironmentPackageResponse) SetHeaders(v map[string]*string) *GetEnvironmentPackageResponse {
	s.Headers = v
	return s
}

func (s *GetEnvironmentPackageResponse) SetBody(v *GetEnvironmentPackageResponseBody) *GetEnvironmentPackageResponse {
	s.Body = v
	return s
}

type GetProductComponentDetailResponseBody struct {
	Data    *GetProductComponentDetailResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	ErrCode *string                                    `json:"errCode,omitempty" xml:"errCode,omitempty"`
	Success *bool                                      `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetProductComponentDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetProductComponentDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetProductComponentDetailResponseBody) SetData(v *GetProductComponentDetailResponseBodyData) *GetProductComponentDetailResponseBody {
	s.Data = v
	return s
}

func (s *GetProductComponentDetailResponseBody) SetErrCode(v string) *GetProductComponentDetailResponseBody {
	s.ErrCode = &v
	return s
}

func (s *GetProductComponentDetailResponseBody) SetSuccess(v bool) *GetProductComponentDetailResponseBody {
	s.Success = &v
	return s
}

type GetProductComponentDetailResponseBodyData struct {
	AppVersion                   *string                                                                  `json:"appVersion,omitempty" xml:"appVersion,omitempty"`
	Category                     *string                                                                  `json:"category,omitempty" xml:"category,omitempty"`
	ChildrenComponentVersionList []*GetProductComponentDetailResponseBodyDataChildrenComponentVersionList `json:"childrenComponentVersionList,omitempty" xml:"childrenComponentVersionList,omitempty" type:"Repeated"`
	Class                        *string                                                                  `json:"class,omitempty" xml:"class,omitempty"`
	ComponentName                *string                                                                  `json:"componentName,omitempty" xml:"componentName,omitempty"`
	ComponentUID                 *string                                                                  `json:"componentUID,omitempty" xml:"componentUID,omitempty"`
	Description                  *string                                                                  `json:"description,omitempty" xml:"description,omitempty"`
	Documents                    []*string                                                                `json:"documents,omitempty" xml:"documents,omitempty" type:"Repeated"`
	Enable                       *bool                                                                    `json:"enable,omitempty" xml:"enable,omitempty"`
	HasDependency                *bool                                                                    `json:"hasDependency,omitempty" xml:"hasDependency,omitempty"`
	ImagesMapping                *string                                                                  `json:"imagesMapping,omitempty" xml:"imagesMapping,omitempty"`
	Namespace                    *string                                                                  `json:"namespace,omitempty" xml:"namespace,omitempty"`
	OrchestrationValues          *string                                                                  `json:"orchestrationValues,omitempty" xml:"orchestrationValues,omitempty"`
	PackageURL                   *string                                                                  `json:"packageURL,omitempty" xml:"packageURL,omitempty"`
	ParentComponent              *bool                                                                    `json:"parentComponent,omitempty" xml:"parentComponent,omitempty"`
	Priority                     *int32                                                                   `json:"priority,omitempty" xml:"priority,omitempty"`
	ProductComponentVersionUID   *string                                                                  `json:"productComponentVersionUID,omitempty" xml:"productComponentVersionUID,omitempty"`
	Provider                     *string                                                                  `json:"provider,omitempty" xml:"provider,omitempty"`
	Readme                       *string                                                                  `json:"readme,omitempty" xml:"readme,omitempty"`
	Resources                    *string                                                                  `json:"resources,omitempty" xml:"resources,omitempty"`
	Singleton                    *bool                                                                    `json:"singleton,omitempty" xml:"singleton,omitempty"`
	Uid                          *string                                                                  `json:"uid,omitempty" xml:"uid,omitempty"`
	Version                      *string                                                                  `json:"version,omitempty" xml:"version,omitempty"`
}

func (s GetProductComponentDetailResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetProductComponentDetailResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetProductComponentDetailResponseBodyData) SetAppVersion(v string) *GetProductComponentDetailResponseBodyData {
	s.AppVersion = &v
	return s
}

func (s *GetProductComponentDetailResponseBodyData) SetCategory(v string) *GetProductComponentDetailResponseBodyData {
	s.Category = &v
	return s
}

func (s *GetProductComponentDetailResponseBodyData) SetChildrenComponentVersionList(v []*GetProductComponentDetailResponseBodyDataChildrenComponentVersionList) *GetProductComponentDetailResponseBodyData {
	s.ChildrenComponentVersionList = v
	return s
}

func (s *GetProductComponentDetailResponseBodyData) SetClass(v string) *GetProductComponentDetailResponseBodyData {
	s.Class = &v
	return s
}

func (s *GetProductComponentDetailResponseBodyData) SetComponentName(v string) *GetProductComponentDetailResponseBodyData {
	s.ComponentName = &v
	return s
}

func (s *GetProductComponentDetailResponseBodyData) SetComponentUID(v string) *GetProductComponentDetailResponseBodyData {
	s.ComponentUID = &v
	return s
}

func (s *GetProductComponentDetailResponseBodyData) SetDescription(v string) *GetProductComponentDetailResponseBodyData {
	s.Description = &v
	return s
}

func (s *GetProductComponentDetailResponseBodyData) SetDocuments(v []*string) *GetProductComponentDetailResponseBodyData {
	s.Documents = v
	return s
}

func (s *GetProductComponentDetailResponseBodyData) SetEnable(v bool) *GetProductComponentDetailResponseBodyData {
	s.Enable = &v
	return s
}

func (s *GetProductComponentDetailResponseBodyData) SetHasDependency(v bool) *GetProductComponentDetailResponseBodyData {
	s.HasDependency = &v
	return s
}

func (s *GetProductComponentDetailResponseBodyData) SetImagesMapping(v string) *GetProductComponentDetailResponseBodyData {
	s.ImagesMapping = &v
	return s
}

func (s *GetProductComponentDetailResponseBodyData) SetNamespace(v string) *GetProductComponentDetailResponseBodyData {
	s.Namespace = &v
	return s
}

func (s *GetProductComponentDetailResponseBodyData) SetOrchestrationValues(v string) *GetProductComponentDetailResponseBodyData {
	s.OrchestrationValues = &v
	return s
}

func (s *GetProductComponentDetailResponseBodyData) SetPackageURL(v string) *GetProductComponentDetailResponseBodyData {
	s.PackageURL = &v
	return s
}

func (s *GetProductComponentDetailResponseBodyData) SetParentComponent(v bool) *GetProductComponentDetailResponseBodyData {
	s.ParentComponent = &v
	return s
}

func (s *GetProductComponentDetailResponseBodyData) SetPriority(v int32) *GetProductComponentDetailResponseBodyData {
	s.Priority = &v
	return s
}

func (s *GetProductComponentDetailResponseBodyData) SetProductComponentVersionUID(v string) *GetProductComponentDetailResponseBodyData {
	s.ProductComponentVersionUID = &v
	return s
}

func (s *GetProductComponentDetailResponseBodyData) SetProvider(v string) *GetProductComponentDetailResponseBodyData {
	s.Provider = &v
	return s
}

func (s *GetProductComponentDetailResponseBodyData) SetReadme(v string) *GetProductComponentDetailResponseBodyData {
	s.Readme = &v
	return s
}

func (s *GetProductComponentDetailResponseBodyData) SetResources(v string) *GetProductComponentDetailResponseBodyData {
	s.Resources = &v
	return s
}

func (s *GetProductComponentDetailResponseBodyData) SetSingleton(v bool) *GetProductComponentDetailResponseBodyData {
	s.Singleton = &v
	return s
}

func (s *GetProductComponentDetailResponseBodyData) SetUid(v string) *GetProductComponentDetailResponseBodyData {
	s.Uid = &v
	return s
}

func (s *GetProductComponentDetailResponseBodyData) SetVersion(v string) *GetProductComponentDetailResponseBodyData {
	s.Version = &v
	return s
}

type GetProductComponentDetailResponseBodyDataChildrenComponentVersionList struct {
	AppVersion                 *string   `json:"appVersion,omitempty" xml:"appVersion,omitempty"`
	Category                   *string   `json:"category,omitempty" xml:"category,omitempty"`
	Class                      *string   `json:"class,omitempty" xml:"class,omitempty"`
	ComponentName              *string   `json:"componentName,omitempty" xml:"componentName,omitempty"`
	ComponentUID               *string   `json:"componentUID,omitempty" xml:"componentUID,omitempty"`
	Description                *string   `json:"description,omitempty" xml:"description,omitempty"`
	Documents                  []*string `json:"documents,omitempty" xml:"documents,omitempty" type:"Repeated"`
	Enable                     *bool     `json:"enable,omitempty" xml:"enable,omitempty"`
	ImagesMapping              *string   `json:"imagesMapping,omitempty" xml:"imagesMapping,omitempty"`
	Namespace                  *string   `json:"namespace,omitempty" xml:"namespace,omitempty"`
	OrchestrationValues        *string   `json:"orchestrationValues,omitempty" xml:"orchestrationValues,omitempty"`
	PackageURL                 *string   `json:"packageURL,omitempty" xml:"packageURL,omitempty"`
	ParentComponent            *bool     `json:"parentComponent,omitempty" xml:"parentComponent,omitempty"`
	Priority                   *int32    `json:"priority,omitempty" xml:"priority,omitempty"`
	ProductComponentVersionUID *string   `json:"productComponentVersionUID,omitempty" xml:"productComponentVersionUID,omitempty"`
	Provider                   *string   `json:"provider,omitempty" xml:"provider,omitempty"`
	Readme                     *string   `json:"readme,omitempty" xml:"readme,omitempty"`
	Resources                  *string   `json:"resources,omitempty" xml:"resources,omitempty"`
	Singleton                  *bool     `json:"singleton,omitempty" xml:"singleton,omitempty"`
	Uid                        *string   `json:"uid,omitempty" xml:"uid,omitempty"`
	Version                    *string   `json:"version,omitempty" xml:"version,omitempty"`
}

func (s GetProductComponentDetailResponseBodyDataChildrenComponentVersionList) String() string {
	return tea.Prettify(s)
}

func (s GetProductComponentDetailResponseBodyDataChildrenComponentVersionList) GoString() string {
	return s.String()
}

func (s *GetProductComponentDetailResponseBodyDataChildrenComponentVersionList) SetAppVersion(v string) *GetProductComponentDetailResponseBodyDataChildrenComponentVersionList {
	s.AppVersion = &v
	return s
}

func (s *GetProductComponentDetailResponseBodyDataChildrenComponentVersionList) SetCategory(v string) *GetProductComponentDetailResponseBodyDataChildrenComponentVersionList {
	s.Category = &v
	return s
}

func (s *GetProductComponentDetailResponseBodyDataChildrenComponentVersionList) SetClass(v string) *GetProductComponentDetailResponseBodyDataChildrenComponentVersionList {
	s.Class = &v
	return s
}

func (s *GetProductComponentDetailResponseBodyDataChildrenComponentVersionList) SetComponentName(v string) *GetProductComponentDetailResponseBodyDataChildrenComponentVersionList {
	s.ComponentName = &v
	return s
}

func (s *GetProductComponentDetailResponseBodyDataChildrenComponentVersionList) SetComponentUID(v string) *GetProductComponentDetailResponseBodyDataChildrenComponentVersionList {
	s.ComponentUID = &v
	return s
}

func (s *GetProductComponentDetailResponseBodyDataChildrenComponentVersionList) SetDescription(v string) *GetProductComponentDetailResponseBodyDataChildrenComponentVersionList {
	s.Description = &v
	return s
}

func (s *GetProductComponentDetailResponseBodyDataChildrenComponentVersionList) SetDocuments(v []*string) *GetProductComponentDetailResponseBodyDataChildrenComponentVersionList {
	s.Documents = v
	return s
}

func (s *GetProductComponentDetailResponseBodyDataChildrenComponentVersionList) SetEnable(v bool) *GetProductComponentDetailResponseBodyDataChildrenComponentVersionList {
	s.Enable = &v
	return s
}

func (s *GetProductComponentDetailResponseBodyDataChildrenComponentVersionList) SetImagesMapping(v string) *GetProductComponentDetailResponseBodyDataChildrenComponentVersionList {
	s.ImagesMapping = &v
	return s
}

func (s *GetProductComponentDetailResponseBodyDataChildrenComponentVersionList) SetNamespace(v string) *GetProductComponentDetailResponseBodyDataChildrenComponentVersionList {
	s.Namespace = &v
	return s
}

func (s *GetProductComponentDetailResponseBodyDataChildrenComponentVersionList) SetOrchestrationValues(v string) *GetProductComponentDetailResponseBodyDataChildrenComponentVersionList {
	s.OrchestrationValues = &v
	return s
}

func (s *GetProductComponentDetailResponseBodyDataChildrenComponentVersionList) SetPackageURL(v string) *GetProductComponentDetailResponseBodyDataChildrenComponentVersionList {
	s.PackageURL = &v
	return s
}

func (s *GetProductComponentDetailResponseBodyDataChildrenComponentVersionList) SetParentComponent(v bool) *GetProductComponentDetailResponseBodyDataChildrenComponentVersionList {
	s.ParentComponent = &v
	return s
}

func (s *GetProductComponentDetailResponseBodyDataChildrenComponentVersionList) SetPriority(v int32) *GetProductComponentDetailResponseBodyDataChildrenComponentVersionList {
	s.Priority = &v
	return s
}

func (s *GetProductComponentDetailResponseBodyDataChildrenComponentVersionList) SetProductComponentVersionUID(v string) *GetProductComponentDetailResponseBodyDataChildrenComponentVersionList {
	s.ProductComponentVersionUID = &v
	return s
}

func (s *GetProductComponentDetailResponseBodyDataChildrenComponentVersionList) SetProvider(v string) *GetProductComponentDetailResponseBodyDataChildrenComponentVersionList {
	s.Provider = &v
	return s
}

func (s *GetProductComponentDetailResponseBodyDataChildrenComponentVersionList) SetReadme(v string) *GetProductComponentDetailResponseBodyDataChildrenComponentVersionList {
	s.Readme = &v
	return s
}

func (s *GetProductComponentDetailResponseBodyDataChildrenComponentVersionList) SetResources(v string) *GetProductComponentDetailResponseBodyDataChildrenComponentVersionList {
	s.Resources = &v
	return s
}

func (s *GetProductComponentDetailResponseBodyDataChildrenComponentVersionList) SetSingleton(v bool) *GetProductComponentDetailResponseBodyDataChildrenComponentVersionList {
	s.Singleton = &v
	return s
}

func (s *GetProductComponentDetailResponseBodyDataChildrenComponentVersionList) SetUid(v string) *GetProductComponentDetailResponseBodyDataChildrenComponentVersionList {
	s.Uid = &v
	return s
}

func (s *GetProductComponentDetailResponseBodyDataChildrenComponentVersionList) SetVersion(v string) *GetProductComponentDetailResponseBodyDataChildrenComponentVersionList {
	s.Version = &v
	return s
}

type GetProductComponentDetailResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetProductComponentDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetProductComponentDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s GetProductComponentDetailResponse) GoString() string {
	return s.String()
}

func (s *GetProductComponentDetailResponse) SetHeaders(v map[string]*string) *GetProductComponentDetailResponse {
	s.Headers = v
	return s
}

func (s *GetProductComponentDetailResponse) SetBody(v *GetProductComponentDetailResponseBody) *GetProductComponentDetailResponse {
	s.Body = v
	return s
}

type ImportEnvironmentNodesRequest struct {
	NodeListYaml *string `json:"nodeListYaml,omitempty" xml:"nodeListYaml,omitempty"`
}

func (s ImportEnvironmentNodesRequest) String() string {
	return tea.Prettify(s)
}

func (s ImportEnvironmentNodesRequest) GoString() string {
	return s.String()
}

func (s *ImportEnvironmentNodesRequest) SetNodeListYaml(v string) *ImportEnvironmentNodesRequest {
	s.NodeListYaml = &v
	return s
}

type ImportEnvironmentNodesResponseBody struct {
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	ErrMsg  *string `json:"errMsg,omitempty" xml:"errMsg,omitempty"`
	Success *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ImportEnvironmentNodesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ImportEnvironmentNodesResponseBody) GoString() string {
	return s.String()
}

func (s *ImportEnvironmentNodesResponseBody) SetErrCode(v string) *ImportEnvironmentNodesResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ImportEnvironmentNodesResponseBody) SetErrMsg(v string) *ImportEnvironmentNodesResponseBody {
	s.ErrMsg = &v
	return s
}

func (s *ImportEnvironmentNodesResponseBody) SetSuccess(v bool) *ImportEnvironmentNodesResponseBody {
	s.Success = &v
	return s
}

type ImportEnvironmentNodesResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ImportEnvironmentNodesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ImportEnvironmentNodesResponse) String() string {
	return tea.Prettify(s)
}

func (s ImportEnvironmentNodesResponse) GoString() string {
	return s.String()
}

func (s *ImportEnvironmentNodesResponse) SetHeaders(v map[string]*string) *ImportEnvironmentNodesResponse {
	s.Headers = v
	return s
}

func (s *ImportEnvironmentNodesResponse) SetBody(v *ImportEnvironmentNodesResponseBody) *ImportEnvironmentNodesResponse {
	s.Body = v
	return s
}

type ListComponentsRequest struct {
	Name     *string `json:"name,omitempty" xml:"name,omitempty"`
	Category *string `json:"category,omitempty" xml:"category,omitempty"`
	PageNum  *int32  `json:"pageNum,omitempty" xml:"pageNum,omitempty"`
	PageSize *int32  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	Fuzzy    *string `json:"fuzzy,omitempty" xml:"fuzzy,omitempty"`
	Public   *bool   `json:"public,omitempty" xml:"public,omitempty"`
}

func (s ListComponentsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListComponentsRequest) GoString() string {
	return s.String()
}

func (s *ListComponentsRequest) SetName(v string) *ListComponentsRequest {
	s.Name = &v
	return s
}

func (s *ListComponentsRequest) SetCategory(v string) *ListComponentsRequest {
	s.Category = &v
	return s
}

func (s *ListComponentsRequest) SetPageNum(v int32) *ListComponentsRequest {
	s.PageNum = &v
	return s
}

func (s *ListComponentsRequest) SetPageSize(v int32) *ListComponentsRequest {
	s.PageSize = &v
	return s
}

func (s *ListComponentsRequest) SetFuzzy(v string) *ListComponentsRequest {
	s.Fuzzy = &v
	return s
}

func (s *ListComponentsRequest) SetPublic(v bool) *ListComponentsRequest {
	s.Public = &v
	return s
}

type ListComponentsResponseBody struct {
	Data    *ListComponentsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	ErrCode *string                         `json:"errCode,omitempty" xml:"errCode,omitempty"`
	ErrMsg  *string                         `json:"errMsg,omitempty" xml:"errMsg,omitempty"`
	Success *bool                           `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ListComponentsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListComponentsResponseBody) GoString() string {
	return s.String()
}

func (s *ListComponentsResponseBody) SetData(v *ListComponentsResponseBodyData) *ListComponentsResponseBody {
	s.Data = v
	return s
}

func (s *ListComponentsResponseBody) SetErrCode(v string) *ListComponentsResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ListComponentsResponseBody) SetErrMsg(v string) *ListComponentsResponseBody {
	s.ErrMsg = &v
	return s
}

func (s *ListComponentsResponseBody) SetSuccess(v bool) *ListComponentsResponseBody {
	s.Success = &v
	return s
}

type ListComponentsResponseBodyData struct {
	List     []*ListComponentsResponseBodyDataList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	PageNum  *int32                                `json:"pageNum,omitempty" xml:"pageNum,omitempty"`
	PageSize *int32                                `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	Total    *int32                                `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListComponentsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListComponentsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListComponentsResponseBodyData) SetList(v []*ListComponentsResponseBodyDataList) *ListComponentsResponseBodyData {
	s.List = v
	return s
}

func (s *ListComponentsResponseBodyData) SetPageNum(v int32) *ListComponentsResponseBodyData {
	s.PageNum = &v
	return s
}

func (s *ListComponentsResponseBodyData) SetPageSize(v int32) *ListComponentsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListComponentsResponseBodyData) SetTotal(v int32) *ListComponentsResponseBodyData {
	s.Total = &v
	return s
}

type ListComponentsResponseBodyDataList struct {
	Annotations *ListComponentsResponseBodyDataListAnnotations `json:"annotations,omitempty" xml:"annotations,omitempty" type:"Struct"`
	Category    *string                                        `json:"category,omitempty" xml:"category,omitempty"`
	Description *string                                        `json:"description,omitempty" xml:"description,omitempty"`
	Documents   *string                                        `json:"documents,omitempty" xml:"documents,omitempty"`
	Name        *string                                        `json:"name,omitempty" xml:"name,omitempty"`
	Provider    *string                                        `json:"provider,omitempty" xml:"provider,omitempty"`
	Public      *bool                                          `json:"public,omitempty" xml:"public,omitempty"`
	Singleton   *bool                                          `json:"singleton,omitempty" xml:"singleton,omitempty"`
	Source      *string                                        `json:"source,omitempty" xml:"source,omitempty"`
	Uid         *string                                        `json:"uid,omitempty" xml:"uid,omitempty"`
}

func (s ListComponentsResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s ListComponentsResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *ListComponentsResponseBodyDataList) SetAnnotations(v *ListComponentsResponseBodyDataListAnnotations) *ListComponentsResponseBodyDataList {
	s.Annotations = v
	return s
}

func (s *ListComponentsResponseBodyDataList) SetCategory(v string) *ListComponentsResponseBodyDataList {
	s.Category = &v
	return s
}

func (s *ListComponentsResponseBodyDataList) SetDescription(v string) *ListComponentsResponseBodyDataList {
	s.Description = &v
	return s
}

func (s *ListComponentsResponseBodyDataList) SetDocuments(v string) *ListComponentsResponseBodyDataList {
	s.Documents = &v
	return s
}

func (s *ListComponentsResponseBodyDataList) SetName(v string) *ListComponentsResponseBodyDataList {
	s.Name = &v
	return s
}

func (s *ListComponentsResponseBodyDataList) SetProvider(v string) *ListComponentsResponseBodyDataList {
	s.Provider = &v
	return s
}

func (s *ListComponentsResponseBodyDataList) SetPublic(v bool) *ListComponentsResponseBodyDataList {
	s.Public = &v
	return s
}

func (s *ListComponentsResponseBodyDataList) SetSingleton(v bool) *ListComponentsResponseBodyDataList {
	s.Singleton = &v
	return s
}

func (s *ListComponentsResponseBodyDataList) SetSource(v string) *ListComponentsResponseBodyDataList {
	s.Source = &v
	return s
}

func (s *ListComponentsResponseBodyDataList) SetUid(v string) *ListComponentsResponseBodyDataList {
	s.Uid = &v
	return s
}

type ListComponentsResponseBodyDataListAnnotations struct {
	Annotations *string `json:"annotations,omitempty" xml:"annotations,omitempty"`
}

func (s ListComponentsResponseBodyDataListAnnotations) String() string {
	return tea.Prettify(s)
}

func (s ListComponentsResponseBodyDataListAnnotations) GoString() string {
	return s.String()
}

func (s *ListComponentsResponseBodyDataListAnnotations) SetAnnotations(v string) *ListComponentsResponseBodyDataListAnnotations {
	s.Annotations = &v
	return s
}

type ListComponentsResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListComponentsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListComponentsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListComponentsResponse) GoString() string {
	return s.String()
}

func (s *ListComponentsResponse) SetHeaders(v map[string]*string) *ListComponentsResponse {
	s.Headers = v
	return s
}

func (s *ListComponentsResponse) SetBody(v *ListComponentsResponseBody) *ListComponentsResponse {
	s.Body = v
	return s
}

type AddEnvironmentProductVersionHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	ClientToken   *string            `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s AddEnvironmentProductVersionHeaders) String() string {
	return tea.Prettify(s)
}

func (s AddEnvironmentProductVersionHeaders) GoString() string {
	return s.String()
}

func (s *AddEnvironmentProductVersionHeaders) SetCommonHeaders(v map[string]*string) *AddEnvironmentProductVersionHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AddEnvironmentProductVersionHeaders) SetClientToken(v string) *AddEnvironmentProductVersionHeaders {
	s.ClientToken = &v
	return s
}

type AddEnvironmentProductVersionRequest struct {
	ProductName       *string `json:"productName,omitempty" xml:"productName,omitempty"`
	ProductUID        *string `json:"productUID,omitempty" xml:"productUID,omitempty"`
	ProductVersion    *string `json:"productVersion,omitempty" xml:"productVersion,omitempty"`
	ProductVersionUID *string `json:"productVersionUID,omitempty" xml:"productVersionUID,omitempty"`
}

func (s AddEnvironmentProductVersionRequest) String() string {
	return tea.Prettify(s)
}

func (s AddEnvironmentProductVersionRequest) GoString() string {
	return s.String()
}

func (s *AddEnvironmentProductVersionRequest) SetProductName(v string) *AddEnvironmentProductVersionRequest {
	s.ProductName = &v
	return s
}

func (s *AddEnvironmentProductVersionRequest) SetProductUID(v string) *AddEnvironmentProductVersionRequest {
	s.ProductUID = &v
	return s
}

func (s *AddEnvironmentProductVersionRequest) SetProductVersion(v string) *AddEnvironmentProductVersionRequest {
	s.ProductVersion = &v
	return s
}

func (s *AddEnvironmentProductVersionRequest) SetProductVersionUID(v string) *AddEnvironmentProductVersionRequest {
	s.ProductVersionUID = &v
	return s
}

type AddEnvironmentProductVersionResponseBody struct {
	Data    *AddEnvironmentProductVersionResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	ErrCode *string                                       `json:"errCode,omitempty" xml:"errCode,omitempty"`
	ErrMsg  *string                                       `json:"errMsg,omitempty" xml:"errMsg,omitempty"`
	Success *bool                                         `json:"success,omitempty" xml:"success,omitempty"`
}

func (s AddEnvironmentProductVersionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddEnvironmentProductVersionResponseBody) GoString() string {
	return s.String()
}

func (s *AddEnvironmentProductVersionResponseBody) SetData(v *AddEnvironmentProductVersionResponseBodyData) *AddEnvironmentProductVersionResponseBody {
	s.Data = v
	return s
}

func (s *AddEnvironmentProductVersionResponseBody) SetErrCode(v string) *AddEnvironmentProductVersionResponseBody {
	s.ErrCode = &v
	return s
}

func (s *AddEnvironmentProductVersionResponseBody) SetErrMsg(v string) *AddEnvironmentProductVersionResponseBody {
	s.ErrMsg = &v
	return s
}

func (s *AddEnvironmentProductVersionResponseBody) SetSuccess(v bool) *AddEnvironmentProductVersionResponseBody {
	s.Success = &v
	return s
}

type AddEnvironmentProductVersionResponseBodyData struct {
	Uid *string `json:"uid,omitempty" xml:"uid,omitempty"`
}

func (s AddEnvironmentProductVersionResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s AddEnvironmentProductVersionResponseBodyData) GoString() string {
	return s.String()
}

func (s *AddEnvironmentProductVersionResponseBodyData) SetUid(v string) *AddEnvironmentProductVersionResponseBodyData {
	s.Uid = &v
	return s
}

type AddEnvironmentProductVersionResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddEnvironmentProductVersionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddEnvironmentProductVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s AddEnvironmentProductVersionResponse) GoString() string {
	return s.String()
}

func (s *AddEnvironmentProductVersionResponse) SetHeaders(v map[string]*string) *AddEnvironmentProductVersionResponse {
	s.Headers = v
	return s
}

func (s *AddEnvironmentProductVersionResponse) SetBody(v *AddEnvironmentProductVersionResponseBody) *AddEnvironmentProductVersionResponse {
	s.Body = v
	return s
}

type ListProductVersionsRequest struct {
	Released  *bool                                  `json:"released,omitempty" xml:"released,omitempty"`
	Platforms []*ListProductVersionsRequestPlatforms `json:"platforms,omitempty" xml:"platforms,omitempty" type:"Repeated"`
}

func (s ListProductVersionsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListProductVersionsRequest) GoString() string {
	return s.String()
}

func (s *ListProductVersionsRequest) SetReleased(v bool) *ListProductVersionsRequest {
	s.Released = &v
	return s
}

func (s *ListProductVersionsRequest) SetPlatforms(v []*ListProductVersionsRequestPlatforms) *ListProductVersionsRequest {
	s.Platforms = v
	return s
}

type ListProductVersionsRequestPlatforms struct {
	Architecture *string `json:"architecture,omitempty" xml:"architecture,omitempty"`
	Os           *string `json:"os,omitempty" xml:"os,omitempty"`
}

func (s ListProductVersionsRequestPlatforms) String() string {
	return tea.Prettify(s)
}

func (s ListProductVersionsRequestPlatforms) GoString() string {
	return s.String()
}

func (s *ListProductVersionsRequestPlatforms) SetArchitecture(v string) *ListProductVersionsRequestPlatforms {
	s.Architecture = &v
	return s
}

func (s *ListProductVersionsRequestPlatforms) SetOs(v string) *ListProductVersionsRequestPlatforms {
	s.Os = &v
	return s
}

type ListProductVersionsShrinkRequest struct {
	Released        *bool   `json:"released,omitempty" xml:"released,omitempty"`
	PlatformsShrink *string `json:"platforms,omitempty" xml:"platforms,omitempty"`
}

func (s ListProductVersionsShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListProductVersionsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListProductVersionsShrinkRequest) SetReleased(v bool) *ListProductVersionsShrinkRequest {
	s.Released = &v
	return s
}

func (s *ListProductVersionsShrinkRequest) SetPlatformsShrink(v string) *ListProductVersionsShrinkRequest {
	s.PlatformsShrink = &v
	return s
}

type ListProductVersionsResponseBody struct {
	Data    *ListProductVersionsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	ErrCode *string                              `json:"errCode,omitempty" xml:"errCode,omitempty"`
	ErrMsg  *string                              `json:"errMsg,omitempty" xml:"errMsg,omitempty"`
	Success *bool                                `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ListProductVersionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListProductVersionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListProductVersionsResponseBody) SetData(v *ListProductVersionsResponseBodyData) *ListProductVersionsResponseBody {
	s.Data = v
	return s
}

func (s *ListProductVersionsResponseBody) SetErrCode(v string) *ListProductVersionsResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ListProductVersionsResponseBody) SetErrMsg(v string) *ListProductVersionsResponseBody {
	s.ErrMsg = &v
	return s
}

func (s *ListProductVersionsResponseBody) SetSuccess(v bool) *ListProductVersionsResponseBody {
	s.Success = &v
	return s
}

type ListProductVersionsResponseBodyData struct {
	List     []*ListProductVersionsResponseBodyDataList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	PageNum  *int32                                     `json:"pageNum,omitempty" xml:"pageNum,omitempty"`
	PageSize *int32                                     `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	Total    *int32                                     `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListProductVersionsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListProductVersionsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListProductVersionsResponseBodyData) SetList(v []*ListProductVersionsResponseBodyDataList) *ListProductVersionsResponseBodyData {
	s.List = v
	return s
}

func (s *ListProductVersionsResponseBodyData) SetPageNum(v int32) *ListProductVersionsResponseBodyData {
	s.PageNum = &v
	return s
}

func (s *ListProductVersionsResponseBodyData) SetPageSize(v int32) *ListProductVersionsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListProductVersionsResponseBodyData) SetTotal(v int32) *ListProductVersionsResponseBodyData {
	s.Total = &v
	return s
}

type ListProductVersionsResponseBodyDataList struct {
	Annotations *ListProductVersionsResponseBodyDataListAnnotations `json:"annotations,omitempty" xml:"annotations,omitempty" type:"Struct"`
	Description *string                                             `json:"description,omitempty" xml:"description,omitempty"`
	PackageURL  *string                                             `json:"packageURL,omitempty" xml:"packageURL,omitempty"`
	ProductName *string                                             `json:"productName,omitempty" xml:"productName,omitempty"`
	ProductUID  *string                                             `json:"productUID,omitempty" xml:"productUID,omitempty"`
	Provider    *string                                             `json:"provider,omitempty" xml:"provider,omitempty"`
	Uid         *string                                             `json:"uid,omitempty" xml:"uid,omitempty"`
	Version     *string                                             `json:"version,omitempty" xml:"version,omitempty"`
}

func (s ListProductVersionsResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s ListProductVersionsResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *ListProductVersionsResponseBodyDataList) SetAnnotations(v *ListProductVersionsResponseBodyDataListAnnotations) *ListProductVersionsResponseBodyDataList {
	s.Annotations = v
	return s
}

func (s *ListProductVersionsResponseBodyDataList) SetDescription(v string) *ListProductVersionsResponseBodyDataList {
	s.Description = &v
	return s
}

func (s *ListProductVersionsResponseBodyDataList) SetPackageURL(v string) *ListProductVersionsResponseBodyDataList {
	s.PackageURL = &v
	return s
}

func (s *ListProductVersionsResponseBodyDataList) SetProductName(v string) *ListProductVersionsResponseBodyDataList {
	s.ProductName = &v
	return s
}

func (s *ListProductVersionsResponseBodyDataList) SetProductUID(v string) *ListProductVersionsResponseBodyDataList {
	s.ProductUID = &v
	return s
}

func (s *ListProductVersionsResponseBodyDataList) SetProvider(v string) *ListProductVersionsResponseBodyDataList {
	s.Provider = &v
	return s
}

func (s *ListProductVersionsResponseBodyDataList) SetUid(v string) *ListProductVersionsResponseBodyDataList {
	s.Uid = &v
	return s
}

func (s *ListProductVersionsResponseBodyDataList) SetVersion(v string) *ListProductVersionsResponseBodyDataList {
	s.Version = &v
	return s
}

type ListProductVersionsResponseBodyDataListAnnotations struct {
	AdditionalProp1 *string `json:"additionalProp1,omitempty" xml:"additionalProp1,omitempty"`
	AdditionalProp2 *string `json:"additionalProp2,omitempty" xml:"additionalProp2,omitempty"`
	AdditionalProp3 *string `json:"additionalProp3,omitempty" xml:"additionalProp3,omitempty"`
}

func (s ListProductVersionsResponseBodyDataListAnnotations) String() string {
	return tea.Prettify(s)
}

func (s ListProductVersionsResponseBodyDataListAnnotations) GoString() string {
	return s.String()
}

func (s *ListProductVersionsResponseBodyDataListAnnotations) SetAdditionalProp1(v string) *ListProductVersionsResponseBodyDataListAnnotations {
	s.AdditionalProp1 = &v
	return s
}

func (s *ListProductVersionsResponseBodyDataListAnnotations) SetAdditionalProp2(v string) *ListProductVersionsResponseBodyDataListAnnotations {
	s.AdditionalProp2 = &v
	return s
}

func (s *ListProductVersionsResponseBodyDataListAnnotations) SetAdditionalProp3(v string) *ListProductVersionsResponseBodyDataListAnnotations {
	s.AdditionalProp3 = &v
	return s
}

type ListProductVersionsResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListProductVersionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListProductVersionsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListProductVersionsResponse) GoString() string {
	return s.String()
}

func (s *ListProductVersionsResponse) SetHeaders(v map[string]*string) *ListProductVersionsResponse {
	s.Headers = v
	return s
}

func (s *ListProductVersionsResponse) SetBody(v *ListProductVersionsResponseBody) *ListProductVersionsResponse {
	s.Body = v
	return s
}

type GetChildrenComponentVersionListResponseBody struct {
	Data    *GetChildrenComponentVersionListResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	ErrCode *string                                          `json:"errCode,omitempty" xml:"errCode,omitempty"`
	ErrMsg  *string                                          `json:"errMsg,omitempty" xml:"errMsg,omitempty"`
	Success *bool                                            `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetChildrenComponentVersionListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetChildrenComponentVersionListResponseBody) GoString() string {
	return s.String()
}

func (s *GetChildrenComponentVersionListResponseBody) SetData(v *GetChildrenComponentVersionListResponseBodyData) *GetChildrenComponentVersionListResponseBody {
	s.Data = v
	return s
}

func (s *GetChildrenComponentVersionListResponseBody) SetErrCode(v string) *GetChildrenComponentVersionListResponseBody {
	s.ErrCode = &v
	return s
}

func (s *GetChildrenComponentVersionListResponseBody) SetErrMsg(v string) *GetChildrenComponentVersionListResponseBody {
	s.ErrMsg = &v
	return s
}

func (s *GetChildrenComponentVersionListResponseBody) SetSuccess(v bool) *GetChildrenComponentVersionListResponseBody {
	s.Success = &v
	return s
}

type GetChildrenComponentVersionListResponseBodyData struct {
	ClusterId      *string `json:"clusterId,omitempty" xml:"clusterId,omitempty"`
	CreatedAt      *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty"`
	InstanceList   *string `json:"instanceList,omitempty" xml:"instanceList,omitempty"`
	Name           *string `json:"name,omitempty" xml:"name,omitempty"`
	ProductName    *string `json:"productName,omitempty" xml:"productName,omitempty"`
	ProductVersion *string `json:"productVersion,omitempty" xml:"productVersion,omitempty"`
	Uid            *string `json:"uid,omitempty" xml:"uid,omitempty"`
	VendorConfig   *string `json:"vendorConfig,omitempty" xml:"vendorConfig,omitempty"`
	VendorType     *string `json:"vendorType,omitempty" xml:"vendorType,omitempty"`
}

func (s GetChildrenComponentVersionListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetChildrenComponentVersionListResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetChildrenComponentVersionListResponseBodyData) SetClusterId(v string) *GetChildrenComponentVersionListResponseBodyData {
	s.ClusterId = &v
	return s
}

func (s *GetChildrenComponentVersionListResponseBodyData) SetCreatedAt(v string) *GetChildrenComponentVersionListResponseBodyData {
	s.CreatedAt = &v
	return s
}

func (s *GetChildrenComponentVersionListResponseBodyData) SetDescription(v string) *GetChildrenComponentVersionListResponseBodyData {
	s.Description = &v
	return s
}

func (s *GetChildrenComponentVersionListResponseBodyData) SetInstanceList(v string) *GetChildrenComponentVersionListResponseBodyData {
	s.InstanceList = &v
	return s
}

func (s *GetChildrenComponentVersionListResponseBodyData) SetName(v string) *GetChildrenComponentVersionListResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetChildrenComponentVersionListResponseBodyData) SetProductName(v string) *GetChildrenComponentVersionListResponseBodyData {
	s.ProductName = &v
	return s
}

func (s *GetChildrenComponentVersionListResponseBodyData) SetProductVersion(v string) *GetChildrenComponentVersionListResponseBodyData {
	s.ProductVersion = &v
	return s
}

func (s *GetChildrenComponentVersionListResponseBodyData) SetUid(v string) *GetChildrenComponentVersionListResponseBodyData {
	s.Uid = &v
	return s
}

func (s *GetChildrenComponentVersionListResponseBodyData) SetVendorConfig(v string) *GetChildrenComponentVersionListResponseBodyData {
	s.VendorConfig = &v
	return s
}

func (s *GetChildrenComponentVersionListResponseBodyData) SetVendorType(v string) *GetChildrenComponentVersionListResponseBodyData {
	s.VendorType = &v
	return s
}

type GetChildrenComponentVersionListResponse struct {
	Headers map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetChildrenComponentVersionListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetChildrenComponentVersionListResponse) String() string {
	return tea.Prettify(s)
}

func (s GetChildrenComponentVersionListResponse) GoString() string {
	return s.String()
}

func (s *GetChildrenComponentVersionListResponse) SetHeaders(v map[string]*string) *GetChildrenComponentVersionListResponse {
	s.Headers = v
	return s
}

func (s *GetChildrenComponentVersionListResponse) SetBody(v *GetChildrenComponentVersionListResponseBody) *GetChildrenComponentVersionListResponse {
	s.Body = v
	return s
}

type CreateSLRHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	ClientToken   *string            `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s CreateSLRHeaders) String() string {
	return tea.Prettify(s)
}

func (s CreateSLRHeaders) GoString() string {
	return s.String()
}

func (s *CreateSLRHeaders) SetCommonHeaders(v map[string]*string) *CreateSLRHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateSLRHeaders) SetClientToken(v string) *CreateSLRHeaders {
	s.ClientToken = &v
	return s
}

type CreateSLRResponseBody struct {
	Data      *string `json:"data,omitempty" xml:"data,omitempty"`
	ErrCode   *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	ErrMsg    *string `json:"errMsg,omitempty" xml:"errMsg,omitempty"`
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	Success   *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreateSLRResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateSLRResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSLRResponseBody) SetData(v string) *CreateSLRResponseBody {
	s.Data = &v
	return s
}

func (s *CreateSLRResponseBody) SetErrCode(v string) *CreateSLRResponseBody {
	s.ErrCode = &v
	return s
}

func (s *CreateSLRResponseBody) SetErrMsg(v string) *CreateSLRResponseBody {
	s.ErrMsg = &v
	return s
}

func (s *CreateSLRResponseBody) SetErrorCode(v string) *CreateSLRResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateSLRResponseBody) SetSuccess(v bool) *CreateSLRResponseBody {
	s.Success = &v
	return s
}

type CreateSLRResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateSLRResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateSLRResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateSLRResponse) GoString() string {
	return s.String()
}

func (s *CreateSLRResponse) SetHeaders(v map[string]*string) *CreateSLRResponse {
	s.Headers = v
	return s
}

func (s *CreateSLRResponse) SetBody(v *CreateSLRResponseBody) *CreateSLRResponse {
	s.Body = v
	return s
}

type GetProductVersionRelatedComponentVersionDetailResponseBody struct {
	Data    []*ProductComponentRelationDetail `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	ErrCode *string                           `json:"errCode,omitempty" xml:"errCode,omitempty"`
	ErrMsg  *string                           `json:"errMsg,omitempty" xml:"errMsg,omitempty"`
	Success *bool                             `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetProductVersionRelatedComponentVersionDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetProductVersionRelatedComponentVersionDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetProductVersionRelatedComponentVersionDetailResponseBody) SetData(v []*ProductComponentRelationDetail) *GetProductVersionRelatedComponentVersionDetailResponseBody {
	s.Data = v
	return s
}

func (s *GetProductVersionRelatedComponentVersionDetailResponseBody) SetErrCode(v string) *GetProductVersionRelatedComponentVersionDetailResponseBody {
	s.ErrCode = &v
	return s
}

func (s *GetProductVersionRelatedComponentVersionDetailResponseBody) SetErrMsg(v string) *GetProductVersionRelatedComponentVersionDetailResponseBody {
	s.ErrMsg = &v
	return s
}

func (s *GetProductVersionRelatedComponentVersionDetailResponseBody) SetSuccess(v bool) *GetProductVersionRelatedComponentVersionDetailResponseBody {
	s.Success = &v
	return s
}

type GetProductVersionRelatedComponentVersionDetailResponse struct {
	Headers map[string]*string                                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetProductVersionRelatedComponentVersionDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetProductVersionRelatedComponentVersionDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s GetProductVersionRelatedComponentVersionDetailResponse) GoString() string {
	return s.String()
}

func (s *GetProductVersionRelatedComponentVersionDetailResponse) SetHeaders(v map[string]*string) *GetProductVersionRelatedComponentVersionDetailResponse {
	s.Headers = v
	return s
}

func (s *GetProductVersionRelatedComponentVersionDetailResponse) SetBody(v *GetProductVersionRelatedComponentVersionDetailResponseBody) *GetProductVersionRelatedComponentVersionDetailResponse {
	s.Body = v
	return s
}

type AddEnvironmentPackageHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	ClientToken   *string            `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s AddEnvironmentPackageHeaders) String() string {
	return tea.Prettify(s)
}

func (s AddEnvironmentPackageHeaders) GoString() string {
	return s.String()
}

func (s *AddEnvironmentPackageHeaders) SetCommonHeaders(v map[string]*string) *AddEnvironmentPackageHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AddEnvironmentPackageHeaders) SetClientToken(v string) *AddEnvironmentPackageHeaders {
	s.ClientToken = &v
	return s
}

type AddEnvironmentPackageRequest struct {
	PackageType *string `json:"packageType,omitempty" xml:"packageType,omitempty"`
}

func (s AddEnvironmentPackageRequest) String() string {
	return tea.Prettify(s)
}

func (s AddEnvironmentPackageRequest) GoString() string {
	return s.String()
}

func (s *AddEnvironmentPackageRequest) SetPackageType(v string) *AddEnvironmentPackageRequest {
	s.PackageType = &v
	return s
}

type AddEnvironmentPackageResponseBody struct {
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	ErrMsg  *string `json:"errMsg,omitempty" xml:"errMsg,omitempty"`
	Success *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s AddEnvironmentPackageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddEnvironmentPackageResponseBody) GoString() string {
	return s.String()
}

func (s *AddEnvironmentPackageResponseBody) SetErrCode(v string) *AddEnvironmentPackageResponseBody {
	s.ErrCode = &v
	return s
}

func (s *AddEnvironmentPackageResponseBody) SetErrMsg(v string) *AddEnvironmentPackageResponseBody {
	s.ErrMsg = &v
	return s
}

func (s *AddEnvironmentPackageResponseBody) SetSuccess(v bool) *AddEnvironmentPackageResponseBody {
	s.Success = &v
	return s
}

type AddEnvironmentPackageResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddEnvironmentPackageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddEnvironmentPackageResponse) String() string {
	return tea.Prettify(s)
}

func (s AddEnvironmentPackageResponse) GoString() string {
	return s.String()
}

func (s *AddEnvironmentPackageResponse) SetHeaders(v map[string]*string) *AddEnvironmentPackageResponse {
	s.Headers = v
	return s
}

func (s *AddEnvironmentPackageResponse) SetBody(v *AddEnvironmentPackageResponseBody) *AddEnvironmentPackageResponse {
	s.Body = v
	return s
}

type UpdateEnvironmentProductVersionRequest struct {
	CompatibleVersions   *string `json:"compatibleVersions,omitempty" xml:"compatibleVersions,omitempty"`
	OldProductVersion    *string `json:"oldProductVersion,omitempty" xml:"oldProductVersion,omitempty"`
	OldProductVersionUID *string `json:"oldProductVersionUID,omitempty" xml:"oldProductVersionUID,omitempty"`
	ProductName          *string `json:"productName,omitempty" xml:"productName,omitempty"`
	ProductUID           *string `json:"productUID,omitempty" xml:"productUID,omitempty"`
	ProductVersion       *string `json:"productVersion,omitempty" xml:"productVersion,omitempty"`
	ProductVersionUID    *string `json:"productVersionUID,omitempty" xml:"productVersionUID,omitempty"`
}

func (s UpdateEnvironmentProductVersionRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateEnvironmentProductVersionRequest) GoString() string {
	return s.String()
}

func (s *UpdateEnvironmentProductVersionRequest) SetCompatibleVersions(v string) *UpdateEnvironmentProductVersionRequest {
	s.CompatibleVersions = &v
	return s
}

func (s *UpdateEnvironmentProductVersionRequest) SetOldProductVersion(v string) *UpdateEnvironmentProductVersionRequest {
	s.OldProductVersion = &v
	return s
}

func (s *UpdateEnvironmentProductVersionRequest) SetOldProductVersionUID(v string) *UpdateEnvironmentProductVersionRequest {
	s.OldProductVersionUID = &v
	return s
}

func (s *UpdateEnvironmentProductVersionRequest) SetProductName(v string) *UpdateEnvironmentProductVersionRequest {
	s.ProductName = &v
	return s
}

func (s *UpdateEnvironmentProductVersionRequest) SetProductUID(v string) *UpdateEnvironmentProductVersionRequest {
	s.ProductUID = &v
	return s
}

func (s *UpdateEnvironmentProductVersionRequest) SetProductVersion(v string) *UpdateEnvironmentProductVersionRequest {
	s.ProductVersion = &v
	return s
}

func (s *UpdateEnvironmentProductVersionRequest) SetProductVersionUID(v string) *UpdateEnvironmentProductVersionRequest {
	s.ProductVersionUID = &v
	return s
}

type UpdateEnvironmentProductVersionResponseBody struct {
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	ErrMsg  *string `json:"errMsg,omitempty" xml:"errMsg,omitempty"`
	Success *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateEnvironmentProductVersionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateEnvironmentProductVersionResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateEnvironmentProductVersionResponseBody) SetErrCode(v string) *UpdateEnvironmentProductVersionResponseBody {
	s.ErrCode = &v
	return s
}

func (s *UpdateEnvironmentProductVersionResponseBody) SetErrMsg(v string) *UpdateEnvironmentProductVersionResponseBody {
	s.ErrMsg = &v
	return s
}

func (s *UpdateEnvironmentProductVersionResponseBody) SetSuccess(v bool) *UpdateEnvironmentProductVersionResponseBody {
	s.Success = &v
	return s
}

type UpdateEnvironmentProductVersionResponse struct {
	Headers map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateEnvironmentProductVersionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateEnvironmentProductVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateEnvironmentProductVersionResponse) GoString() string {
	return s.String()
}

func (s *UpdateEnvironmentProductVersionResponse) SetHeaders(v map[string]*string) *UpdateEnvironmentProductVersionResponse {
	s.Headers = v
	return s
}

func (s *UpdateEnvironmentProductVersionResponse) SetBody(v *UpdateEnvironmentProductVersionResponseBody) *UpdateEnvironmentProductVersionResponse {
	s.Body = v
	return s
}

type GetProductVersionPlatformsResponseBody struct {
	Data    *GetProductVersionPlatformsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	ErrCode *string                                     `json:"errCode,omitempty" xml:"errCode,omitempty"`
	ErrMsg  *string                                     `json:"errMsg,omitempty" xml:"errMsg,omitempty"`
	Success *bool                                       `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetProductVersionPlatformsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetProductVersionPlatformsResponseBody) GoString() string {
	return s.String()
}

func (s *GetProductVersionPlatformsResponseBody) SetData(v *GetProductVersionPlatformsResponseBodyData) *GetProductVersionPlatformsResponseBody {
	s.Data = v
	return s
}

func (s *GetProductVersionPlatformsResponseBody) SetErrCode(v string) *GetProductVersionPlatformsResponseBody {
	s.ErrCode = &v
	return s
}

func (s *GetProductVersionPlatformsResponseBody) SetErrMsg(v string) *GetProductVersionPlatformsResponseBody {
	s.ErrMsg = &v
	return s
}

func (s *GetProductVersionPlatformsResponseBody) SetSuccess(v bool) *GetProductVersionPlatformsResponseBody {
	s.Success = &v
	return s
}

type GetProductVersionPlatformsResponseBodyData struct {
	List []*GetProductVersionPlatformsResponseBodyDataList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
}

func (s GetProductVersionPlatformsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetProductVersionPlatformsResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetProductVersionPlatformsResponseBodyData) SetList(v []*GetProductVersionPlatformsResponseBodyDataList) *GetProductVersionPlatformsResponseBodyData {
	s.List = v
	return s
}

type GetProductVersionPlatformsResponseBodyDataList struct {
	Architecture *string `json:"architecture,omitempty" xml:"architecture,omitempty"`
	Os           *string `json:"os,omitempty" xml:"os,omitempty"`
}

func (s GetProductVersionPlatformsResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s GetProductVersionPlatformsResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *GetProductVersionPlatformsResponseBodyDataList) SetArchitecture(v string) *GetProductVersionPlatformsResponseBodyDataList {
	s.Architecture = &v
	return s
}

func (s *GetProductVersionPlatformsResponseBodyDataList) SetOs(v string) *GetProductVersionPlatformsResponseBodyDataList {
	s.Os = &v
	return s
}

type GetProductVersionPlatformsResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetProductVersionPlatformsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetProductVersionPlatformsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetProductVersionPlatformsResponse) GoString() string {
	return s.String()
}

func (s *GetProductVersionPlatformsResponse) SetHeaders(v map[string]*string) *GetProductVersionPlatformsResponse {
	s.Headers = v
	return s
}

func (s *GetProductVersionPlatformsResponse) SetBody(v *GetProductVersionPlatformsResponseBody) *GetProductVersionPlatformsResponse {
	s.Body = v
	return s
}

type SaveEnvironmentParamRequest struct {
	ComponentUID        *string   `json:"componentUID,omitempty" xml:"componentUID,omitempty"`
	ComponentVersionUID *string   `json:"componentVersionUID,omitempty" xml:"componentVersionUID,omitempty"`
	Name                *string   `json:"name,omitempty" xml:"name,omitempty"`
	ParamUID            *string   `json:"paramUID,omitempty" xml:"paramUID,omitempty"`
	Provider            *string   `json:"provider,omitempty" xml:"provider,omitempty"`
	ReleaseName         *string   `json:"releaseName,omitempty" xml:"releaseName,omitempty"`
	Scope               []*string `json:"scope,omitempty" xml:"scope,omitempty" type:"Repeated"`
	Value               *string   `json:"value,omitempty" xml:"value,omitempty"`
}

func (s SaveEnvironmentParamRequest) String() string {
	return tea.Prettify(s)
}

func (s SaveEnvironmentParamRequest) GoString() string {
	return s.String()
}

func (s *SaveEnvironmentParamRequest) SetComponentUID(v string) *SaveEnvironmentParamRequest {
	s.ComponentUID = &v
	return s
}

func (s *SaveEnvironmentParamRequest) SetComponentVersionUID(v string) *SaveEnvironmentParamRequest {
	s.ComponentVersionUID = &v
	return s
}

func (s *SaveEnvironmentParamRequest) SetName(v string) *SaveEnvironmentParamRequest {
	s.Name = &v
	return s
}

func (s *SaveEnvironmentParamRequest) SetParamUID(v string) *SaveEnvironmentParamRequest {
	s.ParamUID = &v
	return s
}

func (s *SaveEnvironmentParamRequest) SetProvider(v string) *SaveEnvironmentParamRequest {
	s.Provider = &v
	return s
}

func (s *SaveEnvironmentParamRequest) SetReleaseName(v string) *SaveEnvironmentParamRequest {
	s.ReleaseName = &v
	return s
}

func (s *SaveEnvironmentParamRequest) SetScope(v []*string) *SaveEnvironmentParamRequest {
	s.Scope = v
	return s
}

func (s *SaveEnvironmentParamRequest) SetValue(v string) *SaveEnvironmentParamRequest {
	s.Value = &v
	return s
}

type SaveEnvironmentParamResponseBody struct {
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	ErrMsg  *string `json:"errMsg,omitempty" xml:"errMsg,omitempty"`
	Success *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s SaveEnvironmentParamResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SaveEnvironmentParamResponseBody) GoString() string {
	return s.String()
}

func (s *SaveEnvironmentParamResponseBody) SetErrCode(v string) *SaveEnvironmentParamResponseBody {
	s.ErrCode = &v
	return s
}

func (s *SaveEnvironmentParamResponseBody) SetErrMsg(v string) *SaveEnvironmentParamResponseBody {
	s.ErrMsg = &v
	return s
}

func (s *SaveEnvironmentParamResponseBody) SetSuccess(v bool) *SaveEnvironmentParamResponseBody {
	s.Success = &v
	return s
}

type SaveEnvironmentParamResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SaveEnvironmentParamResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SaveEnvironmentParamResponse) String() string {
	return tea.Prettify(s)
}

func (s SaveEnvironmentParamResponse) GoString() string {
	return s.String()
}

func (s *SaveEnvironmentParamResponse) SetHeaders(v map[string]*string) *SaveEnvironmentParamResponse {
	s.Headers = v
	return s
}

func (s *SaveEnvironmentParamResponse) SetBody(v *SaveEnvironmentParamResponseBody) *SaveEnvironmentParamResponse {
	s.Body = v
	return s
}

type UpdateSnapshotInstanceJoinOptionRequest struct {
	JoinSnapshot *bool   `json:"joinSnapshot,omitempty" xml:"joinSnapshot,omitempty"`
	RootPassword *string `json:"rootPassword,omitempty" xml:"rootPassword,omitempty"`
}

func (s UpdateSnapshotInstanceJoinOptionRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateSnapshotInstanceJoinOptionRequest) GoString() string {
	return s.String()
}

func (s *UpdateSnapshotInstanceJoinOptionRequest) SetJoinSnapshot(v bool) *UpdateSnapshotInstanceJoinOptionRequest {
	s.JoinSnapshot = &v
	return s
}

func (s *UpdateSnapshotInstanceJoinOptionRequest) SetRootPassword(v string) *UpdateSnapshotInstanceJoinOptionRequest {
	s.RootPassword = &v
	return s
}

type UpdateSnapshotInstanceJoinOptionResponseBody struct {
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	ErrMsg  *string `json:"errMsg,omitempty" xml:"errMsg,omitempty"`
	Success *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateSnapshotInstanceJoinOptionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateSnapshotInstanceJoinOptionResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSnapshotInstanceJoinOptionResponseBody) SetErrCode(v string) *UpdateSnapshotInstanceJoinOptionResponseBody {
	s.ErrCode = &v
	return s
}

func (s *UpdateSnapshotInstanceJoinOptionResponseBody) SetErrMsg(v string) *UpdateSnapshotInstanceJoinOptionResponseBody {
	s.ErrMsg = &v
	return s
}

func (s *UpdateSnapshotInstanceJoinOptionResponseBody) SetSuccess(v bool) *UpdateSnapshotInstanceJoinOptionResponseBody {
	s.Success = &v
	return s
}

type UpdateSnapshotInstanceJoinOptionResponse struct {
	Headers map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateSnapshotInstanceJoinOptionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateSnapshotInstanceJoinOptionResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateSnapshotInstanceJoinOptionResponse) GoString() string {
	return s.String()
}

func (s *UpdateSnapshotInstanceJoinOptionResponse) SetHeaders(v map[string]*string) *UpdateSnapshotInstanceJoinOptionResponse {
	s.Headers = v
	return s
}

func (s *UpdateSnapshotInstanceJoinOptionResponse) SetBody(v *UpdateSnapshotInstanceJoinOptionResponseBody) *UpdateSnapshotInstanceJoinOptionResponse {
	s.Body = v
	return s
}

type GetProductVersionResourceResponseBody struct {
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	ErrMsg  *string `json:"errMsg,omitempty" xml:"errMsg,omitempty"`
	Success *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetProductVersionResourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetProductVersionResourceResponseBody) GoString() string {
	return s.String()
}

func (s *GetProductVersionResourceResponseBody) SetErrCode(v string) *GetProductVersionResourceResponseBody {
	s.ErrCode = &v
	return s
}

func (s *GetProductVersionResourceResponseBody) SetErrMsg(v string) *GetProductVersionResourceResponseBody {
	s.ErrMsg = &v
	return s
}

func (s *GetProductVersionResourceResponseBody) SetSuccess(v bool) *GetProductVersionResourceResponseBody {
	s.Success = &v
	return s
}

type GetProductVersionResourceResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetProductVersionResourceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetProductVersionResourceResponse) String() string {
	return tea.Prettify(s)
}

func (s GetProductVersionResourceResponse) GoString() string {
	return s.String()
}

func (s *GetProductVersionResourceResponse) SetHeaders(v map[string]*string) *GetProductVersionResourceResponse {
	s.Headers = v
	return s
}

func (s *GetProductVersionResourceResponse) SetBody(v *GetProductVersionResourceResponseBody) *GetProductVersionResourceResponse {
	s.Body = v
	return s
}

type CreateLicenseRequest struct {
	// expire time
	EffectiveYear *int64 `json:"effectiveYear,omitempty" xml:"effectiveYear,omitempty"`
}

func (s CreateLicenseRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateLicenseRequest) GoString() string {
	return s.String()
}

func (s *CreateLicenseRequest) SetEffectiveYear(v int64) *CreateLicenseRequest {
	s.EffectiveYear = &v
	return s
}

type CreateLicenseResponseBody struct {
	// Id of the request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateLicenseResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateLicenseResponseBody) GoString() string {
	return s.String()
}

func (s *CreateLicenseResponseBody) SetRequestId(v string) *CreateLicenseResponseBody {
	s.RequestId = &v
	return s
}

type CreateLicenseResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateLicenseResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateLicenseResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateLicenseResponse) GoString() string {
	return s.String()
}

func (s *CreateLicenseResponse) SetHeaders(v map[string]*string) *CreateLicenseResponse {
	s.Headers = v
	return s
}

func (s *CreateLicenseResponse) SetBody(v *CreateLicenseResponseBody) *CreateLicenseResponse {
	s.Body = v
	return s
}

type ShareSnapshotRequest struct {
	TargetProvider *string `json:"targetProvider,omitempty" xml:"targetProvider,omitempty"`
}

func (s ShareSnapshotRequest) String() string {
	return tea.Prettify(s)
}

func (s ShareSnapshotRequest) GoString() string {
	return s.String()
}

func (s *ShareSnapshotRequest) SetTargetProvider(v string) *ShareSnapshotRequest {
	s.TargetProvider = &v
	return s
}

type ShareSnapshotResponseBody struct {
	Data    *ShareSnapshotResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	ErrCode *string                        `json:"errCode,omitempty" xml:"errCode,omitempty"`
	ErrMsg  *string                        `json:"errMsg,omitempty" xml:"errMsg,omitempty"`
	Success *bool                          `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ShareSnapshotResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ShareSnapshotResponseBody) GoString() string {
	return s.String()
}

func (s *ShareSnapshotResponseBody) SetData(v *ShareSnapshotResponseBodyData) *ShareSnapshotResponseBody {
	s.Data = v
	return s
}

func (s *ShareSnapshotResponseBody) SetErrCode(v string) *ShareSnapshotResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ShareSnapshotResponseBody) SetErrMsg(v string) *ShareSnapshotResponseBody {
	s.ErrMsg = &v
	return s
}

func (s *ShareSnapshotResponseBody) SetSuccess(v bool) *ShareSnapshotResponseBody {
	s.Success = &v
	return s
}

type ShareSnapshotResponseBodyData struct {
	NewSnapshotUID *string `json:"newSnapshotUID,omitempty" xml:"newSnapshotUID,omitempty"`
}

func (s ShareSnapshotResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ShareSnapshotResponseBodyData) GoString() string {
	return s.String()
}

func (s *ShareSnapshotResponseBodyData) SetNewSnapshotUID(v string) *ShareSnapshotResponseBodyData {
	s.NewSnapshotUID = &v
	return s
}

type ShareSnapshotResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ShareSnapshotResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ShareSnapshotResponse) String() string {
	return tea.Prettify(s)
}

func (s ShareSnapshotResponse) GoString() string {
	return s.String()
}

func (s *ShareSnapshotResponse) SetHeaders(v map[string]*string) *ShareSnapshotResponse {
	s.Headers = v
	return s
}

func (s *ShareSnapshotResponse) SetBody(v *ShareSnapshotResponseBody) *ShareSnapshotResponse {
	s.Body = v
	return s
}

type DeleteEnvironmentParamResponseBody struct {
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	Success    *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrCode    *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
}

func (s DeleteEnvironmentParamResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteEnvironmentParamResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteEnvironmentParamResponseBody) SetErrMessage(v string) *DeleteEnvironmentParamResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DeleteEnvironmentParamResponseBody) SetSuccess(v bool) *DeleteEnvironmentParamResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteEnvironmentParamResponseBody) SetErrCode(v string) *DeleteEnvironmentParamResponseBody {
	s.ErrCode = &v
	return s
}

type DeleteEnvironmentParamResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteEnvironmentParamResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteEnvironmentParamResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteEnvironmentParamResponse) GoString() string {
	return s.String()
}

func (s *DeleteEnvironmentParamResponse) SetHeaders(v map[string]*string) *DeleteEnvironmentParamResponse {
	s.Headers = v
	return s
}

func (s *DeleteEnvironmentParamResponse) SetBody(v *DeleteEnvironmentParamResponseBody) *DeleteEnvironmentParamResponse {
	s.Body = v
	return s
}

type GetProductResponseBody struct {
	Data    *GetProductResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	ErrCode *string                     `json:"errCode,omitempty" xml:"errCode,omitempty"`
	ErrMsg  *string                     `json:"errMsg,omitempty" xml:"errMsg,omitempty"`
	Success *bool                       `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetProductResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetProductResponseBody) GoString() string {
	return s.String()
}

func (s *GetProductResponseBody) SetData(v *GetProductResponseBodyData) *GetProductResponseBody {
	s.Data = v
	return s
}

func (s *GetProductResponseBody) SetErrCode(v string) *GetProductResponseBody {
	s.ErrCode = &v
	return s
}

func (s *GetProductResponseBody) SetErrMsg(v string) *GetProductResponseBody {
	s.ErrMsg = &v
	return s
}

func (s *GetProductResponseBody) SetSuccess(v bool) *GetProductResponseBody {
	s.Success = &v
	return s
}

type GetProductResponseBodyData struct {
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	Name        *string `json:"name,omitempty" xml:"name,omitempty"`
	Provider    *string `json:"provider,omitempty" xml:"provider,omitempty"`
	Uid         *string `json:"uid,omitempty" xml:"uid,omitempty"`
}

func (s GetProductResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetProductResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetProductResponseBodyData) SetDescription(v string) *GetProductResponseBodyData {
	s.Description = &v
	return s
}

func (s *GetProductResponseBodyData) SetName(v string) *GetProductResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetProductResponseBodyData) SetProvider(v string) *GetProductResponseBodyData {
	s.Provider = &v
	return s
}

func (s *GetProductResponseBodyData) SetUid(v string) *GetProductResponseBodyData {
	s.Uid = &v
	return s
}

type GetProductResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetProductResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetProductResponse) String() string {
	return tea.Prettify(s)
}

func (s GetProductResponse) GoString() string {
	return s.String()
}

func (s *GetProductResponse) SetHeaders(v map[string]*string) *GetProductResponse {
	s.Headers = v
	return s
}

func (s *GetProductResponse) SetBody(v *GetProductResponseBody) *GetProductResponse {
	s.Body = v
	return s
}

type DeleteComponentVersionResponseBody struct {
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	Success    *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteComponentVersionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteComponentVersionResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteComponentVersionResponseBody) SetErrMessage(v string) *DeleteComponentVersionResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DeleteComponentVersionResponseBody) SetSuccess(v bool) *DeleteComponentVersionResponseBody {
	s.Success = &v
	return s
}

type DeleteComponentVersionResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteComponentVersionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteComponentVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteComponentVersionResponse) GoString() string {
	return s.String()
}

func (s *DeleteComponentVersionResponse) SetHeaders(v map[string]*string) *DeleteComponentVersionResponse {
	s.Headers = v
	return s
}

func (s *DeleteComponentVersionResponse) SetBody(v *DeleteComponentVersionResponseBody) *DeleteComponentVersionResponse {
	s.Body = v
	return s
}

type DeployEnvironmentProductResponseBody struct {
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	ErrMsg  *string `json:"errMsg,omitempty" xml:"errMsg,omitempty"`
	Success *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeployEnvironmentProductResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeployEnvironmentProductResponseBody) GoString() string {
	return s.String()
}

func (s *DeployEnvironmentProductResponseBody) SetErrCode(v string) *DeployEnvironmentProductResponseBody {
	s.ErrCode = &v
	return s
}

func (s *DeployEnvironmentProductResponseBody) SetErrMsg(v string) *DeployEnvironmentProductResponseBody {
	s.ErrMsg = &v
	return s
}

func (s *DeployEnvironmentProductResponseBody) SetSuccess(v bool) *DeployEnvironmentProductResponseBody {
	s.Success = &v
	return s
}

type DeployEnvironmentProductResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeployEnvironmentProductResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeployEnvironmentProductResponse) String() string {
	return tea.Prettify(s)
}

func (s DeployEnvironmentProductResponse) GoString() string {
	return s.String()
}

func (s *DeployEnvironmentProductResponse) SetHeaders(v map[string]*string) *DeployEnvironmentProductResponse {
	s.Headers = v
	return s
}

func (s *DeployEnvironmentProductResponse) SetBody(v *DeployEnvironmentProductResponseBody) *DeployEnvironmentProductResponse {
	s.Body = v
	return s
}

type InitEnvironmentResourceRequest struct {
	AccessKeyID     *string `json:"accessKeyID,omitempty" xml:"accessKeyID,omitempty"`
	AccessKeySecret *string `json:"accessKeySecret,omitempty" xml:"accessKeySecret,omitempty"`
	SecurityToken   *string `json:"securityToken,omitempty" xml:"securityToken,omitempty"`
}

func (s InitEnvironmentResourceRequest) String() string {
	return tea.Prettify(s)
}

func (s InitEnvironmentResourceRequest) GoString() string {
	return s.String()
}

func (s *InitEnvironmentResourceRequest) SetAccessKeyID(v string) *InitEnvironmentResourceRequest {
	s.AccessKeyID = &v
	return s
}

func (s *InitEnvironmentResourceRequest) SetAccessKeySecret(v string) *InitEnvironmentResourceRequest {
	s.AccessKeySecret = &v
	return s
}

func (s *InitEnvironmentResourceRequest) SetSecurityToken(v string) *InitEnvironmentResourceRequest {
	s.SecurityToken = &v
	return s
}

type InitEnvironmentResourceResponseBody struct {
	Data    *InitEnvironmentResourceResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	ErrCode *string                                  `json:"errCode,omitempty" xml:"errCode,omitempty"`
	ErrMsg  *string                                  `json:"errMsg,omitempty" xml:"errMsg,omitempty"`
	Success *bool                                    `json:"success,omitempty" xml:"success,omitempty"`
}

func (s InitEnvironmentResourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s InitEnvironmentResourceResponseBody) GoString() string {
	return s.String()
}

func (s *InitEnvironmentResourceResponseBody) SetData(v *InitEnvironmentResourceResponseBodyData) *InitEnvironmentResourceResponseBody {
	s.Data = v
	return s
}

func (s *InitEnvironmentResourceResponseBody) SetErrCode(v string) *InitEnvironmentResourceResponseBody {
	s.ErrCode = &v
	return s
}

func (s *InitEnvironmentResourceResponseBody) SetErrMsg(v string) *InitEnvironmentResourceResponseBody {
	s.ErrMsg = &v
	return s
}

func (s *InitEnvironmentResourceResponseBody) SetSuccess(v bool) *InitEnvironmentResourceResponseBody {
	s.Success = &v
	return s
}

type InitEnvironmentResourceResponseBodyData struct {
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s InitEnvironmentResourceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s InitEnvironmentResourceResponseBodyData) GoString() string {
	return s.String()
}

func (s *InitEnvironmentResourceResponseBodyData) SetStatus(v string) *InitEnvironmentResourceResponseBodyData {
	s.Status = &v
	return s
}

type InitEnvironmentResourceResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *InitEnvironmentResourceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s InitEnvironmentResourceResponse) String() string {
	return tea.Prettify(s)
}

func (s InitEnvironmentResourceResponse) GoString() string {
	return s.String()
}

func (s *InitEnvironmentResourceResponse) SetHeaders(v map[string]*string) *InitEnvironmentResourceResponse {
	s.Headers = v
	return s
}

func (s *InitEnvironmentResourceResponse) SetBody(v *InitEnvironmentResourceResponseBody) *InitEnvironmentResourceResponse {
	s.Body = v
	return s
}

type ListFoundationVersionsResponseBody struct {
	Data    []*FoundationVersion `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	ErrCode *string              `json:"errCode,omitempty" xml:"errCode,omitempty"`
	ErrMsg  *string              `json:"errMsg,omitempty" xml:"errMsg,omitempty"`
	Success *bool                `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ListFoundationVersionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListFoundationVersionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListFoundationVersionsResponseBody) SetData(v []*FoundationVersion) *ListFoundationVersionsResponseBody {
	s.Data = v
	return s
}

func (s *ListFoundationVersionsResponseBody) SetErrCode(v string) *ListFoundationVersionsResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ListFoundationVersionsResponseBody) SetErrMsg(v string) *ListFoundationVersionsResponseBody {
	s.ErrMsg = &v
	return s
}

func (s *ListFoundationVersionsResponseBody) SetSuccess(v bool) *ListFoundationVersionsResponseBody {
	s.Success = &v
	return s
}

type ListFoundationVersionsResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListFoundationVersionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListFoundationVersionsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListFoundationVersionsResponse) GoString() string {
	return s.String()
}

func (s *ListFoundationVersionsResponse) SetHeaders(v map[string]*string) *ListFoundationVersionsResponse {
	s.Headers = v
	return s
}

func (s *ListFoundationVersionsResponse) SetBody(v *ListFoundationVersionsResponseBody) *ListFoundationVersionsResponse {
	s.Body = v
	return s
}

type ListEnvironmentDeployRecordRequest struct {
	PageNum  *int32 `json:"pageNum,omitempty" xml:"pageNum,omitempty"`
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s ListEnvironmentDeployRecordRequest) String() string {
	return tea.Prettify(s)
}

func (s ListEnvironmentDeployRecordRequest) GoString() string {
	return s.String()
}

func (s *ListEnvironmentDeployRecordRequest) SetPageNum(v int32) *ListEnvironmentDeployRecordRequest {
	s.PageNum = &v
	return s
}

func (s *ListEnvironmentDeployRecordRequest) SetPageSize(v int32) *ListEnvironmentDeployRecordRequest {
	s.PageSize = &v
	return s
}

type ListEnvironmentDeployRecordResponseBody struct {
	Data    *ListEnvironmentDeployRecordResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	ErrCode *string                                      `json:"errCode,omitempty" xml:"errCode,omitempty"`
	ErrMsg  *string                                      `json:"errMsg,omitempty" xml:"errMsg,omitempty"`
	Success *bool                                        `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ListEnvironmentDeployRecordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListEnvironmentDeployRecordResponseBody) GoString() string {
	return s.String()
}

func (s *ListEnvironmentDeployRecordResponseBody) SetData(v *ListEnvironmentDeployRecordResponseBodyData) *ListEnvironmentDeployRecordResponseBody {
	s.Data = v
	return s
}

func (s *ListEnvironmentDeployRecordResponseBody) SetErrCode(v string) *ListEnvironmentDeployRecordResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ListEnvironmentDeployRecordResponseBody) SetErrMsg(v string) *ListEnvironmentDeployRecordResponseBody {
	s.ErrMsg = &v
	return s
}

func (s *ListEnvironmentDeployRecordResponseBody) SetSuccess(v bool) *ListEnvironmentDeployRecordResponseBody {
	s.Success = &v
	return s
}

type ListEnvironmentDeployRecordResponseBodyData struct {
	List     []*ListEnvironmentDeployRecordResponseBodyDataList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	PageNum  *int32                                             `json:"pageNum,omitempty" xml:"pageNum,omitempty"`
	PageSize *int32                                             `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	Total    *int32                                             `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListEnvironmentDeployRecordResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListEnvironmentDeployRecordResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListEnvironmentDeployRecordResponseBodyData) SetList(v []*ListEnvironmentDeployRecordResponseBodyDataList) *ListEnvironmentDeployRecordResponseBodyData {
	s.List = v
	return s
}

func (s *ListEnvironmentDeployRecordResponseBodyData) SetPageNum(v int32) *ListEnvironmentDeployRecordResponseBodyData {
	s.PageNum = &v
	return s
}

func (s *ListEnvironmentDeployRecordResponseBodyData) SetPageSize(v int32) *ListEnvironmentDeployRecordResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListEnvironmentDeployRecordResponseBodyData) SetTotal(v int32) *ListEnvironmentDeployRecordResponseBodyData {
	s.Total = &v
	return s
}

type ListEnvironmentDeployRecordResponseBodyDataList struct {
	EnvUID   *string `json:"envUID,omitempty" xml:"envUID,omitempty"`
	Provider *string `json:"provider,omitempty" xml:"provider,omitempty"`
	Status   *string `json:"status,omitempty" xml:"status,omitempty"`
	Uid      *string `json:"uid,omitempty" xml:"uid,omitempty"`
}

func (s ListEnvironmentDeployRecordResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s ListEnvironmentDeployRecordResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *ListEnvironmentDeployRecordResponseBodyDataList) SetEnvUID(v string) *ListEnvironmentDeployRecordResponseBodyDataList {
	s.EnvUID = &v
	return s
}

func (s *ListEnvironmentDeployRecordResponseBodyDataList) SetProvider(v string) *ListEnvironmentDeployRecordResponseBodyDataList {
	s.Provider = &v
	return s
}

func (s *ListEnvironmentDeployRecordResponseBodyDataList) SetStatus(v string) *ListEnvironmentDeployRecordResponseBodyDataList {
	s.Status = &v
	return s
}

func (s *ListEnvironmentDeployRecordResponseBodyDataList) SetUid(v string) *ListEnvironmentDeployRecordResponseBodyDataList {
	s.Uid = &v
	return s
}

type ListEnvironmentDeployRecordResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListEnvironmentDeployRecordResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListEnvironmentDeployRecordResponse) String() string {
	return tea.Prettify(s)
}

func (s ListEnvironmentDeployRecordResponse) GoString() string {
	return s.String()
}

func (s *ListEnvironmentDeployRecordResponse) SetHeaders(v map[string]*string) *ListEnvironmentDeployRecordResponse {
	s.Headers = v
	return s
}

func (s *ListEnvironmentDeployRecordResponse) SetBody(v *ListEnvironmentDeployRecordResponseBody) *ListEnvironmentDeployRecordResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("cnip"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) GetEnvironment(uid *string) (_result *GetEnvironmentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetEnvironmentResponse{}
	_body, _err := client.GetEnvironmentWithOptions(uid, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetEnvironmentWithOptions(uid *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetEnvironmentResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	_result = &GetEnvironmentResponse{}
	_body, _err := client.DoROARequest(tea.String("GetEnvironment"), tea.String("2020-12-01"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/api/v1/environments/"+tea.StringValue(uid)), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListProductVersionRelatedComponentVersions(uid *string) (_result *ListProductVersionRelatedComponentVersionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListProductVersionRelatedComponentVersionsResponse{}
	_body, _err := client.ListProductVersionRelatedComponentVersionsWithOptions(uid, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListProductVersionRelatedComponentVersionsWithOptions(uid *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListProductVersionRelatedComponentVersionsResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	_result = &ListProductVersionRelatedComponentVersionsResponse{}
	_body, _err := client.DoROARequest(tea.String("ListProductVersionRelatedComponentVersions"), tea.String("2020-12-01"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/api/v1/product_versions/"+tea.StringValue(uid)+"/component_versions"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetComponentVersionChildren(uid *string, versionUID *string) (_result *GetComponentVersionChildrenResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetComponentVersionChildrenResponse{}
	_body, _err := client.GetComponentVersionChildrenWithOptions(uid, versionUID, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetComponentVersionChildrenWithOptions(uid *string, versionUID *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetComponentVersionChildrenResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	_result = &GetComponentVersionChildrenResponse{}
	_body, _err := client.DoROARequest(tea.String("GetComponentVersionChildren"), tea.String("2020-12-01"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/integration/api/v1/components/"+tea.StringValue(uid)+"/versions/"+tea.StringValue(versionUID)+"/children"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetProductEnvironment(uid *string) (_result *GetProductEnvironmentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetProductEnvironmentResponse{}
	_body, _err := client.GetProductEnvironmentWithOptions(uid, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetProductEnvironmentWithOptions(uid *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetProductEnvironmentResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	_result = &GetProductEnvironmentResponse{}
	_body, _err := client.DoROARequest(tea.String("GetProductEnvironment"), tea.String("2020-12-01"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/api/v1/product_envs/"+tea.StringValue(uid)), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetProductVersionPackage(uid *string, request *GetProductVersionPackageRequest) (_result *GetProductVersionPackageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetProductVersionPackageResponse{}
	_body, _err := client.GetProductVersionPackageWithOptions(uid, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetProductVersionPackageWithOptions(uid *string, request *GetProductVersionPackageRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetProductVersionPackageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Platform)) {
		query["platform"] = request.Platform
	}

	if !tea.BoolValue(util.IsUnset(request.PackageType)) {
		query["packageType"] = request.PackageType
	}

	if !tea.BoolValue(util.IsUnset(request.OldProductVersionUID)) {
		query["oldProductVersionUID"] = request.OldProductVersionUID
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &GetProductVersionPackageResponse{}
	_body, _err := client.DoROARequest(tea.String("GetProductVersionPackage"), tea.String("2020-12-01"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/api/v1/hosting/product_versions/"+tea.StringValue(uid)+"/packages"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListAlicloudRegion() (_result *ListAlicloudRegionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListAlicloudRegionResponse{}
	_body, _err := client.ListAlicloudRegionWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListAlicloudRegionWithOptions(headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListAlicloudRegionResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	_result = &ListAlicloudRegionResponse{}
	_body, _err := client.DoROARequest(tea.String("ListAlicloudRegion"), tea.String("2020-12-01"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/api/v1/alicloud/regions"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListComponentVersions(uid *string, request *ListComponentVersionsRequest) (_result *ListComponentVersionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListComponentVersionsResponse{}
	_body, _err := client.ListComponentVersionsWithOptions(uid, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListComponentVersionsWithOptions(uid *string, tmpReq *ListComponentVersionsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListComponentVersionsResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListComponentVersionsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Platforms)) {
		request.PlatformsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Platforms, tea.String("platforms"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNum)) {
		query["pageNum"] = request.PageNum
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.PlatformsShrink)) {
		query["platforms"] = request.PlatformsShrink
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &ListComponentVersionsResponse{}
	_body, _err := client.DoROARequest(tea.String("ListComponentVersions"), tea.String("2020-12-01"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/api/v1/components/"+tea.StringValue(uid)+"/versions"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateSnapshotInstanceJoinOptionWithBatch(uid *string, request *UpdateSnapshotInstanceJoinOptionWithBatchRequest) (_result *UpdateSnapshotInstanceJoinOptionWithBatchResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateSnapshotInstanceJoinOptionWithBatchResponse{}
	_body, _err := client.UpdateSnapshotInstanceJoinOptionWithBatchWithOptions(uid, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateSnapshotInstanceJoinOptionWithBatchWithOptions(uid *string, request *UpdateSnapshotInstanceJoinOptionWithBatchRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateSnapshotInstanceJoinOptionWithBatchResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceUIDs)) {
		query["instanceUIDs"] = request.InstanceUIDs
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JoinSnapshot)) {
		body["joinSnapshot"] = request.JoinSnapshot
	}

	if !tea.BoolValue(util.IsUnset(request.RootPassword)) {
		body["rootPassword"] = request.RootPassword
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &UpdateSnapshotInstanceJoinOptionWithBatchResponse{}
	_body, _err := client.DoROARequest(tea.String("UpdateSnapshotInstanceJoinOptionWithBatch"), tea.String("2020-12-01"), tea.String("HTTPS"), tea.String("PUT"), tea.String("AK"), tea.String("/api/v1/snapshots/"+tea.StringValue(uid)+"/instances"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GenerateVendorConfigTemplate(uid *string) (_result *GenerateVendorConfigTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GenerateVendorConfigTemplateResponse{}
	_body, _err := client.GenerateVendorConfigTemplateWithOptions(uid, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GenerateVendorConfigTemplateWithOptions(uid *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GenerateVendorConfigTemplateResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	_result = &GenerateVendorConfigTemplateResponse{}
	_body, _err := client.DoROARequest(tea.String("GenerateVendorConfigTemplate"), tea.String("2020-12-01"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/api/v1/environments/"+tea.StringValue(uid)+"/vendorConfigTmpl"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateProductComponent(uid *string, request *UpdateProductComponentRequest) (_result *UpdateProductComponentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateProductComponentResponse{}
	_body, _err := client.UpdateProductComponentWithOptions(uid, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateProductComponentWithOptions(uid *string, request *UpdateProductComponentRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateProductComponentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ComponentOrchestrationValues)) {
		body["componentOrchestrationValues"] = request.ComponentOrchestrationValues
	}

	if !tea.BoolValue(util.IsUnset(request.Enable)) {
		body["enable"] = request.Enable
	}

	if !tea.BoolValue(util.IsUnset(request.ReleaseName)) {
		body["releaseName"] = request.ReleaseName
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &UpdateProductComponentResponse{}
	_body, _err := client.DoROARequest(tea.String("UpdateProductComponent"), tea.String("2020-12-01"), tea.String("HTTPS"), tea.String("PUT"), tea.String("AK"), tea.String("/api/v1/productComponentVersions/"+tea.StringValue(uid)), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateEnvironmentNodes(uid *string, request *UpdateEnvironmentNodesRequest) (_result *UpdateEnvironmentNodesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateEnvironmentNodesResponse{}
	_body, _err := client.UpdateEnvironmentNodesWithOptions(uid, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateEnvironmentNodesWithOptions(uid *string, request *UpdateEnvironmentNodesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateEnvironmentNodesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Cpu)) {
		body["cpu"] = request.Cpu
	}

	if !tea.BoolValue(util.IsUnset(request.DataDisk)) {
		body["dataDisk"] = request.DataDisk
	}

	if !tea.BoolValue(util.IsUnset(request.DataDisk2)) {
		body["dataDisk2"] = request.DataDisk2
	}

	if !tea.BoolValue(util.IsUnset(request.EnvUID)) {
		body["envUID"] = request.EnvUID
	}

	if !tea.BoolValue(util.IsUnset(request.Identifier)) {
		body["identifier"] = request.Identifier
	}

	if !tea.BoolValue(util.IsUnset(request.Labels)) {
		body["labels"] = request.Labels
	}

	if !tea.BoolValue(util.IsUnset(request.Memory)) {
		body["memory"] = request.Memory
	}

	if !tea.BoolValue(util.IsUnset(request.NodeIP)) {
		body["nodeIP"] = request.NodeIP
	}

	if !tea.BoolValue(util.IsUnset(request.RootPassword)) {
		body["rootPassword"] = request.RootPassword
	}

	if !tea.BoolValue(util.IsUnset(request.SystemDisk)) {
		body["systemDisk"] = request.SystemDisk
	}

	if !tea.BoolValue(util.IsUnset(request.SystemDisk2)) {
		body["systemDisk2"] = request.SystemDisk2
	}

	if !tea.BoolValue(util.IsUnset(request.Taints)) {
		body["taints"] = request.Taints
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &UpdateEnvironmentNodesResponse{}
	_body, _err := client.DoROARequest(tea.String("UpdateEnvironmentNodes"), tea.String("2020-12-01"), tea.String("HTTPS"), tea.String("PUT"), tea.String("AK"), tea.String("/api/v1/environmentnodes/"+tea.StringValue(uid)), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListEnvironmentPackages(uid *string, request *ListEnvironmentPackagesRequest) (_result *ListEnvironmentPackagesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListEnvironmentPackagesResponse{}
	_body, _err := client.ListEnvironmentPackagesWithOptions(uid, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListEnvironmentPackagesWithOptions(uid *string, request *ListEnvironmentPackagesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListEnvironmentPackagesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNum)) {
		query["pageNum"] = request.PageNum
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &ListEnvironmentPackagesResponse{}
	_body, _err := client.DoROARequest(tea.String("ListEnvironmentPackages"), tea.String("2020-12-01"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/api/v1/environments/"+tea.StringValue(uid)+"/packages"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateEnvironment(request *CreateEnvironmentRequest) (_result *CreateEnvironmentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateEnvironmentHeaders{}
	_result = &CreateEnvironmentResponse{}
	_body, _err := client.CreateEnvironmentWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateEnvironmentWithOptions(request *CreateEnvironmentRequest, headers *CreateEnvironmentHeaders, runtime *util.RuntimeOptions) (_result *CreateEnvironmentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Annotations)) {
		body["annotations"] = request.Annotations
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Location)) {
		body["location"] = request.Location
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.Platform))) {
		body["platform"] = request.Platform
	}

	if !tea.BoolValue(util.IsUnset(request.VendorType)) {
		body["vendorType"] = request.VendorType
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.ClientToken)) {
		realHeaders["ClientToken"] = headers.ClientToken
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &CreateEnvironmentResponse{}
	_body, _err := client.DoROARequest(tea.String("CreateEnvironment"), tea.String("2020-12-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/api/v1/environments"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetEnvironmentLog(uid *string) (_result *GetEnvironmentLogResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetEnvironmentLogResponse{}
	_body, _err := client.GetEnvironmentLogWithOptions(uid, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetEnvironmentLogWithOptions(uid *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetEnvironmentLogResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	_result = &GetEnvironmentLogResponse{}
	_body, _err := client.DoROARequest(tea.String("GetEnvironmentLog"), tea.String("2020-12-01"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/api/v1/envLogs/"+tea.StringValue(uid)), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListEnvironmentNode(uid *string, request *ListEnvironmentNodeRequest) (_result *ListEnvironmentNodeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListEnvironmentNodeResponse{}
	_body, _err := client.ListEnvironmentNodeWithOptions(uid, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListEnvironmentNodeWithOptions(uid *string, request *ListEnvironmentNodeRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListEnvironmentNodeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNum)) {
		query["pageNum"] = request.PageNum
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Fuzzy)) {
		query["fuzzy"] = request.Fuzzy
	}

	if !tea.BoolValue(util.IsUnset(request.NodeIp)) {
		query["nodeIp"] = request.NodeIp
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &ListEnvironmentNodeResponse{}
	_body, _err := client.DoROARequest(tea.String("ListEnvironmentNode"), tea.String("2020-12-01"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/api/v1/environments/"+tea.StringValue(uid)+"/nodes"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListProductVersionRelatedFoundationComponentVersions(uid *string) (_result *ListProductVersionRelatedFoundationComponentVersionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListProductVersionRelatedFoundationComponentVersionsResponse{}
	_body, _err := client.ListProductVersionRelatedFoundationComponentVersionsWithOptions(uid, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListProductVersionRelatedFoundationComponentVersionsWithOptions(uid *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListProductVersionRelatedFoundationComponentVersionsResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	_result = &ListProductVersionRelatedFoundationComponentVersionsResponse{}
	_body, _err := client.DoROARequest(tea.String("ListProductVersionRelatedFoundationComponentVersions"), tea.String("2020-12-01"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/api/v1/product_versions/"+tea.StringValue(uid)+"/foundation/component_versions"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SyncComponent(request *SyncComponentRequest) (_result *SyncComponentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &SyncComponentResponse{}
	_body, _err := client.SyncComponentWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SyncComponentWithOptions(request *SyncComponentRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *SyncComponentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Region)) {
		query["region"] = request.Region
	}

	if !tea.BoolValue(util.IsUnset(request.BucketName)) {
		query["bucketName"] = request.BucketName
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &SyncComponentResponse{}
	_body, _err := client.DoROARequest(tea.String("SyncComponent"), tea.String("2020-12-01"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/integration/api/v1/oss/sync/apps"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateComponentToProduct(id *string, versionId *string, productComponentVersionRelationId *string, request *UpdateComponentToProductRequest) (_result *UpdateComponentToProductResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateComponentToProductResponse{}
	_body, _err := client.UpdateComponentToProductWithOptions(id, versionId, productComponentVersionRelationId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateComponentToProductWithOptions(id *string, versionId *string, productComponentVersionRelationId *string, request *UpdateComponentToProductRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateComponentToProductResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ComponentVersionID)) {
		query["componentVersionID"] = request.ComponentVersionID
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &UpdateComponentToProductResponse{}
	_body, _err := client.DoROARequest(tea.String("UpdateComponentToProduct"), tea.String("2020-12-01"), tea.String("HTTPS"), tea.String("PUT"), tea.String("AK"), tea.String("/integration/api/v1/products/"+tea.StringValue(id)+"/versions/"+tea.StringValue(versionId)+"/componentVersionRelations/"+tea.StringValue(productComponentVersionRelationId)), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateEnvironmentNode(uid *string, request *CreateEnvironmentNodeRequest) (_result *CreateEnvironmentNodeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateEnvironmentNodeResponse{}
	_body, _err := client.CreateEnvironmentNodeWithOptions(uid, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateEnvironmentNodeWithOptions(uid *string, request *CreateEnvironmentNodeRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateEnvironmentNodeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Cpu)) {
		body["cpu"] = request.Cpu
	}

	if !tea.BoolValue(util.IsUnset(request.DataDisk)) {
		body["dataDisk"] = request.DataDisk
	}

	if !tea.BoolValue(util.IsUnset(request.HostName)) {
		body["hostName"] = request.HostName
	}

	if !tea.BoolValue(util.IsUnset(request.Identifier)) {
		body["identifier"] = request.Identifier
	}

	if !tea.BoolValue(util.IsUnset(request.Labels)) {
		body["labels"] = request.Labels
	}

	if !tea.BoolValue(util.IsUnset(request.Memory)) {
		body["memory"] = request.Memory
	}

	if !tea.BoolValue(util.IsUnset(request.Os)) {
		body["os"] = request.Os
	}

	if !tea.BoolValue(util.IsUnset(request.PrivateIP)) {
		body["privateIP"] = request.PrivateIP
	}

	if !tea.BoolValue(util.IsUnset(request.Provider)) {
		body["provider"] = request.Provider
	}

	if !tea.BoolValue(util.IsUnset(request.RootPassword)) {
		body["rootPassword"] = request.RootPassword
	}

	if !tea.BoolValue(util.IsUnset(request.SystemDisk)) {
		body["systemDisk"] = request.SystemDisk
	}

	if !tea.BoolValue(util.IsUnset(request.Taints)) {
		body["taints"] = request.Taints
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &CreateEnvironmentNodeResponse{}
	_body, _err := client.DoROARequest(tea.String("CreateEnvironmentNode"), tea.String("2020-12-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/api/v1/environments/"+tea.StringValue(uid)+"/nodes"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetComponent(uid *string) (_result *GetComponentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetComponentResponse{}
	_body, _err := client.GetComponentWithOptions(uid, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetComponentWithOptions(uid *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetComponentResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	_result = &GetComponentResponse{}
	_body, _err := client.DoROARequest(tea.String("GetComponent"), tea.String("2020-12-01"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/api/v1/components/"+tea.StringValue(uid)), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListFoundationVersionRelatedComponentVersions(uid *string) (_result *ListFoundationVersionRelatedComponentVersionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListFoundationVersionRelatedComponentVersionsResponse{}
	_body, _err := client.ListFoundationVersionRelatedComponentVersionsWithOptions(uid, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListFoundationVersionRelatedComponentVersionsWithOptions(uid *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListFoundationVersionRelatedComponentVersionsResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	_result = &ListFoundationVersionRelatedComponentVersionsResponse{}
	_body, _err := client.DoROARequest(tea.String("ListFoundationVersionRelatedComponentVersions"), tea.String("2020-12-01"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/api/v1/foundation/versions/"+tea.StringValue(uid)+"/component_versions"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetSnapshot(uid *string) (_result *GetSnapshotResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetSnapshotResponse{}
	_body, _err := client.GetSnapshotWithOptions(uid, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetSnapshotWithOptions(uid *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetSnapshotResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	_result = &GetSnapshotResponse{}
	_body, _err := client.DoROARequest(tea.String("GetSnapshot"), tea.String("2020-12-01"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/api/v1/snapshots/"+tea.StringValue(uid)), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetLicense(uid *string) (_result *GetLicenseResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetLicenseResponse{}
	_body, _err := client.GetLicenseWithOptions(uid, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetLicenseWithOptions(uid *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetLicenseResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	_result = &GetLicenseResponse{}
	_body, _err := client.DoROARequest(tea.String("GetLicense"), tea.String("2020-12-01"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/api/v1/environments/"+tea.StringValue(uid)+"/licenses"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateLatestProductVersion(uid *string, versionUID *string) (_result *CreateLatestProductVersionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateLatestProductVersionHeaders{}
	_result = &CreateLatestProductVersionResponse{}
	_body, _err := client.CreateLatestProductVersionWithOptions(uid, versionUID, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateLatestProductVersionWithOptions(uid *string, versionUID *string, headers *CreateLatestProductVersionHeaders, runtime *util.RuntimeOptions) (_result *CreateLatestProductVersionResponse, _err error) {
	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.ClientToken)) {
		realHeaders["ClientToken"] = headers.ClientToken
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
	}
	_result = &CreateLatestProductVersionResponse{}
	_body, _err := client.DoROARequest(tea.String("CreateLatestProductVersion"), tea.String("2020-12-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/integration/api/v1/products/"+tea.StringValue(uid)+"/versions/"+tea.StringValue(versionUID)), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListAlicloudVPC(regionid *string) (_result *ListAlicloudVPCResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListAlicloudVPCResponse{}
	_body, _err := client.ListAlicloudVPCWithOptions(regionid, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListAlicloudVPCWithOptions(regionid *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListAlicloudVPCResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	_result = &ListAlicloudVPCResponse{}
	_body, _err := client.DoROARequest(tea.String("ListAlicloudVPC"), tea.String("2020-12-01"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/api/v1/alicloud/regions/"+tea.StringValue(regionid)+"/vpcs"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateProduct(request *CreateProductRequest) (_result *CreateProductResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateProductHeaders{}
	_result = &CreateProductResponse{}
	_body, _err := client.CreateProductWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateProductWithOptions(request *CreateProductRequest, headers *CreateProductHeaders, runtime *util.RuntimeOptions) (_result *CreateProductResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Annotations)) {
		body["annotations"] = request.Annotations
	}

	if !tea.BoolValue(util.IsUnset(request.ComponentVersionUID)) {
		body["componentVersionUID"] = request.ComponentVersionUID
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.FoundationVersionUID)) {
		body["foundationVersionUID"] = request.FoundationVersionUID
	}

	if !tea.BoolValue(util.IsUnset(request.ProductName)) {
		body["productName"] = request.ProductName
	}

	if !tea.BoolValue(util.IsUnset(request.PrometheusUID)) {
		body["prometheusUID"] = request.PrometheusUID
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.ClientToken)) {
		realHeaders["ClientToken"] = headers.ClientToken
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &CreateProductResponse{}
	_body, _err := client.DoROARequest(tea.String("CreateProduct"), tea.String("2020-12-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/integration/api/v1/products"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetProductEnvironments(request *GetProductEnvironmentsRequest) (_result *GetProductEnvironmentsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetProductEnvironmentsResponse{}
	_body, _err := client.GetProductEnvironmentsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetProductEnvironmentsWithOptions(tmpReq *GetProductEnvironmentsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetProductEnvironmentsResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &GetProductEnvironmentsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Platforms)) {
		request.PlatformsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Platforms, tea.String("platforms"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ProductUID)) {
		query["productUID"] = request.ProductUID
	}

	if !tea.BoolValue(util.IsUnset(request.EnvType)) {
		query["envType"] = request.EnvType
	}

	if !tea.BoolValue(util.IsUnset(request.PlatformsShrink)) {
		query["platforms"] = request.PlatformsShrink
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &GetProductEnvironmentsResponse{}
	_body, _err := client.DoROARequest(tea.String("GetProductEnvironments"), tea.String("2020-12-01"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/api/v1/product_envs"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteComponent(uid *string) (_result *DeleteComponentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteComponentResponse{}
	_body, _err := client.DeleteComponentWithOptions(uid, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteComponentWithOptions(uid *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteComponentResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	_result = &DeleteComponentResponse{}
	_body, _err := client.DoROARequest(tea.String("DeleteComponent"), tea.String("2020-12-01"), tea.String("HTTPS"), tea.String("DELETE"), tea.String("AK"), tea.String("/integration/api/v1/components/"+tea.StringValue(uid)), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteProductComponent(id *string, versionId *string, productComponentVersionRelationId *string) (_result *DeleteProductComponentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteProductComponentResponse{}
	_body, _err := client.DeleteProductComponentWithOptions(id, versionId, productComponentVersionRelationId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteProductComponentWithOptions(id *string, versionId *string, productComponentVersionRelationId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteProductComponentResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	_result = &DeleteProductComponentResponse{}
	_body, _err := client.DoROARequest(tea.String("DeleteProductComponent"), tea.String("2020-12-01"), tea.String("HTTPS"), tea.String("DELETE"), tea.String("AK"), tea.String("/integration/api/v1/products/"+tea.StringValue(id)+"/versions/"+tea.StringValue(versionId)+"/componentVersionRelations/"+tea.StringValue(productComponentVersionRelationId)), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateEnvironmentWithSnapshot(uid *string, request *CreateEnvironmentWithSnapshotRequest) (_result *CreateEnvironmentWithSnapshotResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateEnvironmentWithSnapshotResponse{}
	_body, _err := client.CreateEnvironmentWithSnapshotWithOptions(uid, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateEnvironmentWithSnapshotWithOptions(uid *string, request *CreateEnvironmentWithSnapshotRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateEnvironmentWithSnapshotResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EnvironmentDesc)) {
		body["environmentDesc"] = request.EnvironmentDesc
	}

	if !tea.BoolValue(util.IsUnset(request.EnvironmentName)) {
		body["environmentName"] = request.EnvironmentName
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &CreateEnvironmentWithSnapshotResponse{}
	_body, _err := client.DoROARequest(tea.String("CreateEnvironmentWithSnapshot"), tea.String("2020-12-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/api/v1/snapshots/"+tea.StringValue(uid)+"/environments"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteEnvironment(uid *string) (_result *DeleteEnvironmentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteEnvironmentResponse{}
	_body, _err := client.DeleteEnvironmentWithOptions(uid, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteEnvironmentWithOptions(uid *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteEnvironmentResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	_result = &DeleteEnvironmentResponse{}
	_body, _err := client.DoROARequest(tea.String("DeleteEnvironment"), tea.String("2020-12-01"), tea.String("HTTPS"), tea.String("DELETE"), tea.String("AK"), tea.String("/api/v1/environments/"+tea.StringValue(uid)), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateProductVersion(uid *string, request *UpdateProductVersionRequest) (_result *UpdateProductVersionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateProductVersionResponse{}
	_body, _err := client.UpdateProductVersionWithOptions(uid, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateProductVersionWithOptions(uid *string, request *UpdateProductVersionRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateProductVersionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CompatibleVersions)) {
		body["compatibleVersions"] = request.CompatibleVersions
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Version)) {
		body["version"] = request.Version
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &UpdateProductVersionResponse{}
	_body, _err := client.DoROARequest(tea.String("UpdateProductVersion"), tea.String("2020-12-01"), tea.String("HTTPS"), tea.String("PUT"), tea.String("AK"), tea.String("/api/v1/product_versions/"+tea.StringValue(uid)), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetChildrenComponentVersionParametersList(id *string, versionId *string, request *GetChildrenComponentVersionParametersListRequest) (_result *GetChildrenComponentVersionParametersListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetChildrenComponentVersionParametersListResponse{}
	_body, _err := client.GetChildrenComponentVersionParametersListWithOptions(id, versionId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetChildrenComponentVersionParametersListWithOptions(id *string, versionId *string, request *GetChildrenComponentVersionParametersListRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetChildrenComponentVersionParametersListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RelationId)) {
		query["relation_id"] = request.RelationId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &GetChildrenComponentVersionParametersListResponse{}
	_body, _err := client.DoROARequest(tea.String("GetChildrenComponentVersionParametersList"), tea.String("2020-12-01"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/integration/api/v1/components/"+tea.StringValue(id)+"/versions/"+tea.StringValue(versionId)+"/parameters"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateSnapshot(request *CreateSnapshotRequest) (_result *CreateSnapshotResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateSnapshotResponse{}
	_body, _err := client.CreateSnapshotWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateSnapshotWithOptions(request *CreateSnapshotRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateSnapshotResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.ProductName)) {
		body["productName"] = request.ProductName
	}

	if !tea.BoolValue(util.IsUnset(request.ProductVersion)) {
		body["productVersion"] = request.ProductVersion
	}

	if !tea.BoolValue(util.IsUnset(request.ProductVersionDesc)) {
		body["productVersionDesc"] = request.ProductVersionDesc
	}

	if !tea.BoolValue(util.IsUnset(request.Region)) {
		body["region"] = request.Region
	}

	if !tea.BoolValue(util.IsUnset(request.Vpcid)) {
		body["vpcid"] = request.Vpcid
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &CreateSnapshotResponse{}
	_body, _err := client.DoROARequest(tea.String("CreateSnapshot"), tea.String("2020-12-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/api/v1/snapshots"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetLatestVersionDifferences(id *string, versionID *string, request *GetLatestVersionDifferencesRequest) (_result *GetLatestVersionDifferencesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetLatestVersionDifferencesResponse{}
	_body, _err := client.GetLatestVersionDifferencesWithOptions(id, versionID, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetLatestVersionDifferencesWithOptions(id *string, versionID *string, request *GetLatestVersionDifferencesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetLatestVersionDifferencesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PreVersionID)) {
		query["preVersionID"] = request.PreVersionID
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &GetLatestVersionDifferencesResponse{}
	_body, _err := client.DoROARequest(tea.String("GetLatestVersionDifferences"), tea.String("2020-12-01"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/integration/api/v1/products/"+tea.StringValue(id)+"/versions/"+tea.StringValue(versionID)+"/differences"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteEnvironmentNode(uid *string, request *DeleteEnvironmentNodeRequest) (_result *DeleteEnvironmentNodeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteEnvironmentNodeResponse{}
	_body, _err := client.DeleteEnvironmentNodeWithOptions(uid, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteEnvironmentNodeWithOptions(uid *string, request *DeleteEnvironmentNodeRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteEnvironmentNodeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EnvUID)) {
		query["envUID"] = request.EnvUID
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &DeleteEnvironmentNodeResponse{}
	_body, _err := client.DoROARequest(tea.String("DeleteEnvironmentNode"), tea.String("2020-12-01"), tea.String("HTTPS"), tea.String("DELETE"), tea.String("AK"), tea.String("/api/v1/environmentnodes/"+tea.StringValue(uid)), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ApplyComponent(request *ApplyComponentRequest) (_result *ApplyComponentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ApplyComponentResponse{}
	_body, _err := client.ApplyComponentWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ApplyComponentWithOptions(request *ApplyComponentRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ApplyComponentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Annotations)) {
		body["annotations"] = request.Annotations
	}

	if !tea.BoolValue(util.IsUnset(request.AppVersion)) {
		body["appVersion"] = request.AppVersion
	}

	if !tea.BoolValue(util.IsUnset(request.Category)) {
		body["category"] = request.Category
	}

	if !tea.BoolValue(util.IsUnset(request.ComponentClass)) {
		body["componentClass"] = request.ComponentClass
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Documents)) {
		body["documents"] = request.Documents
	}

	if !tea.BoolValue(util.IsUnset(request.ImagesMapping)) {
		body["imagesMapping"] = request.ImagesMapping
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Namespace)) {
		body["namespace"] = request.Namespace
	}

	if !tea.BoolValue(util.IsUnset(request.OrchestrationType)) {
		body["orchestrationType"] = request.OrchestrationType
	}

	if !tea.BoolValue(util.IsUnset(request.OrchestrationValues)) {
		body["orchestrationValues"] = request.OrchestrationValues
	}

	if !tea.BoolValue(util.IsUnset(request.PackageURL)) {
		body["packageURL"] = request.PackageURL
	}

	if !tea.BoolValue(util.IsUnset(request.ParentComponent)) {
		body["parentComponent"] = request.ParentComponent
	}

	if !tea.BoolValue(util.IsUnset(request.Platforms)) {
		body["platforms"] = request.Platforms
	}

	if !tea.BoolValue(util.IsUnset(request.Priority)) {
		body["priority"] = request.Priority
	}

	if !tea.BoolValue(util.IsUnset(request.Provider)) {
		body["provider"] = request.Provider
	}

	if !tea.BoolValue(util.IsUnset(request.Public)) {
		body["public"] = request.Public
	}

	if !tea.BoolValue(util.IsUnset(request.Readme)) {
		body["readme"] = request.Readme
	}

	if !tea.BoolValue(util.IsUnset(request.Resources)) {
		body["resources"] = request.Resources
	}

	if !tea.BoolValue(util.IsUnset(request.Singleton)) {
		body["singleton"] = request.Singleton
	}

	if !tea.BoolValue(util.IsUnset(request.Version)) {
		body["version"] = request.Version
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &ApplyComponentResponse{}
	_body, _err := client.DoROARequest(tea.String("ApplyComponent"), tea.String("2020-12-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/integration/api/v1/apps"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetSnapshotInstances(uid *string, request *GetSnapshotInstancesRequest) (_result *GetSnapshotInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetSnapshotInstancesResponse{}
	_body, _err := client.GetSnapshotInstancesWithOptions(uid, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetSnapshotInstancesWithOptions(uid *string, request *GetSnapshotInstancesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetSnapshotInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.PageNum)) {
		query["pageNum"] = request.PageNum
	}

	if !tea.BoolValue(util.IsUnset(request.SortKey)) {
		query["sortKey"] = request.SortKey
	}

	if !tea.BoolValue(util.IsUnset(request.SortDirect)) {
		query["sortDirect"] = request.SortDirect
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &GetSnapshotInstancesResponse{}
	_body, _err := client.DoROARequest(tea.String("GetSnapshotInstances"), tea.String("2020-12-01"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/api/v1/snapshots/"+tea.StringValue(uid)+"/instances"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListEnvironments(request *ListEnvironmentsRequest) (_result *ListEnvironmentsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListEnvironmentsResponse{}
	_body, _err := client.ListEnvironmentsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListEnvironmentsWithOptions(request *ListEnvironmentsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListEnvironmentsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNum)) {
		query["pageNum"] = request.PageNum
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Fuzzy)) {
		query["fuzzy"] = request.Fuzzy
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceStatus)) {
		query["instanceStatus"] = request.InstanceStatus
	}

	if !tea.BoolValue(util.IsUnset(request.VendorType)) {
		query["vendorType"] = request.VendorType
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &ListEnvironmentsResponse{}
	_body, _err := client.DoROARequest(tea.String("ListEnvironments"), tea.String("2020-12-01"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/api/v1/environments"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CheckSLR(request *CheckSLRRequest) (_result *CheckSLRResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CheckSLRResponse{}
	_body, _err := client.CheckSLRWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CheckSLRWithOptions(request *CheckSLRRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CheckSLRResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Uid)) {
		query["uid"] = request.Uid
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &CheckSLRResponse{}
	_body, _err := client.DoROARequest(tea.String("CheckSLR"), tea.String("2020-12-01"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/api/v1/slr"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateProduct(uid *string, request *UpdateProductRequest) (_result *UpdateProductResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateProductResponse{}
	_body, _err := client.UpdateProductWithOptions(uid, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateProductWithOptions(uid *string, request *UpdateProductRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateProductResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    request.Description,
	}
	_result = &UpdateProductResponse{}
	_body, _err := client.DoROARequest(tea.String("UpdateProduct"), tea.String("2020-12-01"), tea.String("HTTPS"), tea.String("PUT"), tea.String("AK"), tea.String("/api/v1/products/"+tea.StringValue(uid)), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ApplyComponents(request *ApplyComponentsRequest) (_result *ApplyComponentsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ApplyComponentsResponse{}
	_body, _err := client.ApplyComponentsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ApplyComponentsWithOptions(request *ApplyComponentsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ApplyComponentsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ChildrenList)) {
		body["childrenList"] = request.ChildrenList
	}

	if !tea.BoolValue(util.IsUnset(request.Component)) {
		body["component"] = request.Component
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &ApplyComponentsResponse{}
	_body, _err := client.DoROARequest(tea.String("ApplyComponents"), tea.String("2020-12-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/integration/api/v1/components"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreatePackageConfig(uid *string) (_result *CreatePackageConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreatePackageConfigResponse{}
	_body, _err := client.CreatePackageConfigWithOptions(uid, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreatePackageConfigWithOptions(uid *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreatePackageConfigResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	_result = &CreatePackageConfigResponse{}
	_body, _err := client.DoROARequest(tea.String("CreatePackageConfig"), tea.String("2020-12-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/api/v1/environments/"+tea.StringValue(uid)+"/package_config"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddProductComponent(id *string, versionId *string, componentVersionId *string, ClientToken *string, request *AddProductComponentRequest) (_result *AddProductComponentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &AddProductComponentResponse{}
	_body, _err := client.AddProductComponentWithOptions(id, versionId, componentVersionId, ClientToken, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddProductComponentWithOptions(id *string, versionId *string, componentVersionId *string, ClientToken *string, request *AddProductComponentRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *AddProductComponentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    request.ReleaseName,
	}
	_result = &AddProductComponentResponse{}
	_body, _err := client.DoROARequest(tea.String("AddProductComponent"), tea.String("2020-12-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/integration/api/v1/products/"+tea.StringValue(id)+"/versions/"+tea.StringValue(versionId)+"/componentVersions/"+tea.StringValue(componentVersionId)), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteSnapshot(uid *string) (_result *DeleteSnapshotResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteSnapshotResponse{}
	_body, _err := client.DeleteSnapshotWithOptions(uid, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteSnapshotWithOptions(uid *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteSnapshotResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	_result = &DeleteSnapshotResponse{}
	_body, _err := client.DoROARequest(tea.String("DeleteSnapshot"), tea.String("2020-12-01"), tea.String("HTTPS"), tea.String("DELETE"), tea.String("AK"), tea.String("/api/v1/snapshots/"+tea.StringValue(uid)), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListEnvironmentsWithSnapshot(uid *string) (_result *ListEnvironmentsWithSnapshotResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListEnvironmentsWithSnapshotResponse{}
	_body, _err := client.ListEnvironmentsWithSnapshotWithOptions(uid, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListEnvironmentsWithSnapshotWithOptions(uid *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListEnvironmentsWithSnapshotResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	_result = &ListEnvironmentsWithSnapshotResponse{}
	_body, _err := client.DoROARequest(tea.String("ListEnvironmentsWithSnapshot"), tea.String("2020-12-01"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/api/v1/snapshots/"+tea.StringValue(uid)+"/environments"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetEnvironmentNode(uid *string) (_result *GetEnvironmentNodeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetEnvironmentNodeResponse{}
	_body, _err := client.GetEnvironmentNodeWithOptions(uid, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetEnvironmentNodeWithOptions(uid *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetEnvironmentNodeResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	_result = &GetEnvironmentNodeResponse{}
	_body, _err := client.DoROARequest(tea.String("GetEnvironmentNode"), tea.String("2020-12-01"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/api/v1/environmentnodes/"+tea.StringValue(uid)), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateSnapshot(uid *string, request *UpdateSnapshotRequest) (_result *UpdateSnapshotResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateSnapshotResponse{}
	_body, _err := client.UpdateSnapshotWithOptions(uid, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateSnapshotWithOptions(uid *string, request *UpdateSnapshotRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateSnapshotResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.ProductName)) {
		body["productName"] = request.ProductName
	}

	if !tea.BoolValue(util.IsUnset(request.ProductVersion)) {
		body["productVersion"] = request.ProductVersion
	}

	if !tea.BoolValue(util.IsUnset(request.ProductVersionDesc)) {
		body["productVersionDesc"] = request.ProductVersionDesc
	}

	if !tea.BoolValue(util.IsUnset(request.UpdateKey)) {
		body["updateKey"] = request.UpdateKey
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &UpdateSnapshotResponse{}
	_body, _err := client.DoROARequest(tea.String("UpdateSnapshot"), tea.String("2020-12-01"), tea.String("HTTPS"), tea.String("PUT"), tea.String("AK"), tea.String("/api/v1/snapshots/"+tea.StringValue(uid)), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateEnvironmentAndGenerateVendorConfig(request *CreateEnvironmentAndGenerateVendorConfigRequest) (_result *CreateEnvironmentAndGenerateVendorConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateEnvironmentAndGenerateVendorConfigHeaders{}
	_result = &CreateEnvironmentAndGenerateVendorConfigResponse{}
	_body, _err := client.CreateEnvironmentAndGenerateVendorConfigWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateEnvironmentAndGenerateVendorConfigWithOptions(request *CreateEnvironmentAndGenerateVendorConfigRequest, headers *CreateEnvironmentAndGenerateVendorConfigHeaders, runtime *util.RuntimeOptions) (_result *CreateEnvironmentAndGenerateVendorConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EnvUID)) {
		body["envUID"] = request.EnvUID
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.Platform))) {
		body["platform"] = request.Platform
	}

	if !tea.BoolValue(util.IsUnset(request.ProductName)) {
		body["productName"] = request.ProductName
	}

	if !tea.BoolValue(util.IsUnset(request.ProductUID)) {
		body["productUID"] = request.ProductUID
	}

	if !tea.BoolValue(util.IsUnset(request.ProductVersion)) {
		body["productVersion"] = request.ProductVersion
	}

	if !tea.BoolValue(util.IsUnset(request.ProductVersionUID)) {
		body["productVersionUID"] = request.ProductVersionUID
	}

	if !tea.BoolValue(util.IsUnset(request.VendorType)) {
		body["vendorType"] = request.VendorType
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.ClientToken)) {
		realHeaders["ClientToken"] = headers.ClientToken
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &CreateEnvironmentAndGenerateVendorConfigResponse{}
	_body, _err := client.DoROARequest(tea.String("CreateEnvironmentAndGenerateVendorConfig"), tea.String("2020-12-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/api/v1/product_envs/vendor_config"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateEnvironmentSnapshot(uid *string, request *CreateEnvironmentSnapshotRequest) (_result *CreateEnvironmentSnapshotResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateEnvironmentSnapshotResponse{}
	_body, _err := client.CreateEnvironmentSnapshotWithOptions(uid, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateEnvironmentSnapshotWithOptions(uid *string, request *CreateEnvironmentSnapshotRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateEnvironmentSnapshotResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &CreateEnvironmentSnapshotResponse{}
	_body, _err := client.DoROARequest(tea.String("CreateEnvironmentSnapshot"), tea.String("2020-12-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/api/v1/environments/"+tea.StringValue(uid)+"/snapshots"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) InitSnapshotInstance(uid *string) (_result *InitSnapshotInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &InitSnapshotInstanceResponse{}
	_body, _err := client.InitSnapshotInstanceWithOptions(uid, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) InitSnapshotInstanceWithOptions(uid *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *InitSnapshotInstanceResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	_result = &InitSnapshotInstanceResponse{}
	_body, _err := client.DoROARequest(tea.String("InitSnapshotInstance"), tea.String("2020-12-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/api/v1/snapshots/"+tea.StringValue(uid)+"/instances"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateProductVersionRelatedFoundationVersion(uid *string, request *UpdateProductVersionRelatedFoundationVersionRequest) (_result *UpdateProductVersionRelatedFoundationVersionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateProductVersionRelatedFoundationVersionResponse{}
	_body, _err := client.UpdateProductVersionRelatedFoundationVersionWithOptions(uid, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateProductVersionRelatedFoundationVersionWithOptions(uid *string, request *UpdateProductVersionRelatedFoundationVersionRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateProductVersionRelatedFoundationVersionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    request.FoundationVersionUID,
	}
	_result = &UpdateProductVersionRelatedFoundationVersionResponse{}
	_body, _err := client.DoROARequest(tea.String("UpdateProductVersionRelatedFoundationVersion"), tea.String("2020-12-01"), tea.String("HTTPS"), tea.String("PUT"), tea.String("AK"), tea.String("/api/v1/product_versions/"+tea.StringValue(uid)+"/foundation"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListEnvironmentParams(uid *string, request *ListEnvironmentParamsRequest) (_result *ListEnvironmentParamsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListEnvironmentParamsResponse{}
	_body, _err := client.ListEnvironmentParamsWithOptions(uid, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListEnvironmentParamsWithOptions(uid *string, request *ListEnvironmentParamsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListEnvironmentParamsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNum)) {
		query["pageNum"] = request.PageNum
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Fuzzy)) {
		query["fuzzy"] = request.Fuzzy
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &ListEnvironmentParamsResponse{}
	_body, _err := client.DoROARequest(tea.String("ListEnvironmentParams"), tea.String("2020-12-01"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/api/v1/environments/"+tea.StringValue(uid)+"/params"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetFoundationVersion(uid *string) (_result *GetFoundationVersionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetFoundationVersionResponse{}
	_body, _err := client.GetFoundationVersionWithOptions(uid, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetFoundationVersionWithOptions(uid *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetFoundationVersionResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	_result = &GetFoundationVersionResponse{}
	_body, _err := client.DoROARequest(tea.String("GetFoundationVersion"), tea.String("2020-12-01"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/api/v1/foundation/versions/"+tea.StringValue(uid)), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteProduct(uid *string) (_result *DeleteProductResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteProductResponse{}
	_body, _err := client.DeleteProductWithOptions(uid, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteProductWithOptions(uid *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteProductResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	_result = &DeleteProductResponse{}
	_body, _err := client.DoROARequest(tea.String("DeleteProduct"), tea.String("2020-12-01"), tea.String("HTTPS"), tea.String("DELETE"), tea.String("AK"), tea.String("/integration/api/v1/products/"+tea.StringValue(uid)), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateEnvironment(uid *string, request *UpdateEnvironmentRequest) (_result *UpdateEnvironmentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateEnvironmentResponse{}
	_body, _err := client.UpdateEnvironmentWithOptions(uid, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateEnvironmentWithOptions(uid *string, request *UpdateEnvironmentRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateEnvironmentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Location)) {
		body["location"] = request.Location
	}

	if !tea.BoolValue(util.IsUnset(request.VendorConfig)) {
		body["vendorConfig"] = request.VendorConfig
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &UpdateEnvironmentResponse{}
	_body, _err := client.DoROARequest(tea.String("UpdateEnvironment"), tea.String("2020-12-01"), tea.String("HTTPS"), tea.String("PUT"), tea.String("AK"), tea.String("/api/v1/environments/"+tea.StringValue(uid)), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetEnvironmentPackage(uid *string) (_result *GetEnvironmentPackageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetEnvironmentPackageResponse{}
	_body, _err := client.GetEnvironmentPackageWithOptions(uid, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetEnvironmentPackageWithOptions(uid *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetEnvironmentPackageResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	_result = &GetEnvironmentPackageResponse{}
	_body, _err := client.DoROARequest(tea.String("GetEnvironmentPackage"), tea.String("2020-12-01"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/api/v1/envPackages/"+tea.StringValue(uid)), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetProductComponentDetail(uid *string, versionUID *string, productComponentVersionRelationUID *string) (_result *GetProductComponentDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetProductComponentDetailResponse{}
	_body, _err := client.GetProductComponentDetailWithOptions(uid, versionUID, productComponentVersionRelationUID, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetProductComponentDetailWithOptions(uid *string, versionUID *string, productComponentVersionRelationUID *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetProductComponentDetailResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	_result = &GetProductComponentDetailResponse{}
	_body, _err := client.DoROARequest(tea.String("GetProductComponentDetail"), tea.String("2020-12-01"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/integration/api/v1/products/"+tea.StringValue(uid)+"/versions/"+tea.StringValue(versionUID)+"/productComponentVersionRelations/"+tea.StringValue(productComponentVersionRelationUID)+"/detail"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ImportEnvironmentNodes(uid *string, request *ImportEnvironmentNodesRequest) (_result *ImportEnvironmentNodesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ImportEnvironmentNodesResponse{}
	_body, _err := client.ImportEnvironmentNodesWithOptions(uid, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ImportEnvironmentNodesWithOptions(uid *string, request *ImportEnvironmentNodesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ImportEnvironmentNodesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    request.NodeListYaml,
	}
	_result = &ImportEnvironmentNodesResponse{}
	_body, _err := client.DoROARequest(tea.String("ImportEnvironmentNodes"), tea.String("2020-12-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/api/v1/environments/"+tea.StringValue(uid)+"/importnodes"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListComponents(request *ListComponentsRequest) (_result *ListComponentsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListComponentsResponse{}
	_body, _err := client.ListComponentsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListComponentsWithOptions(request *ListComponentsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListComponentsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Category)) {
		query["category"] = request.Category
	}

	if !tea.BoolValue(util.IsUnset(request.PageNum)) {
		query["pageNum"] = request.PageNum
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Fuzzy)) {
		query["fuzzy"] = request.Fuzzy
	}

	if !tea.BoolValue(util.IsUnset(request.Public)) {
		query["public"] = request.Public
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &ListComponentsResponse{}
	_body, _err := client.DoROARequest(tea.String("ListComponents"), tea.String("2020-12-01"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/api/v1/components"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddEnvironmentProductVersion(uid *string, request *AddEnvironmentProductVersionRequest) (_result *AddEnvironmentProductVersionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &AddEnvironmentProductVersionHeaders{}
	_result = &AddEnvironmentProductVersionResponse{}
	_body, _err := client.AddEnvironmentProductVersionWithOptions(uid, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddEnvironmentProductVersionWithOptions(uid *string, request *AddEnvironmentProductVersionRequest, headers *AddEnvironmentProductVersionHeaders, runtime *util.RuntimeOptions) (_result *AddEnvironmentProductVersionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ProductName)) {
		body["productName"] = request.ProductName
	}

	if !tea.BoolValue(util.IsUnset(request.ProductUID)) {
		body["productUID"] = request.ProductUID
	}

	if !tea.BoolValue(util.IsUnset(request.ProductVersion)) {
		body["productVersion"] = request.ProductVersion
	}

	if !tea.BoolValue(util.IsUnset(request.ProductVersionUID)) {
		body["productVersionUID"] = request.ProductVersionUID
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.ClientToken)) {
		realHeaders["ClientToken"] = headers.ClientToken
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &AddEnvironmentProductVersionResponse{}
	_body, _err := client.DoROARequest(tea.String("AddEnvironmentProductVersion"), tea.String("2020-12-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/api/v1/environments/"+tea.StringValue(uid)+"/productVersions"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListProductVersions(uid *string, request *ListProductVersionsRequest) (_result *ListProductVersionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListProductVersionsResponse{}
	_body, _err := client.ListProductVersionsWithOptions(uid, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListProductVersionsWithOptions(uid *string, tmpReq *ListProductVersionsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListProductVersionsResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListProductVersionsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Platforms)) {
		request.PlatformsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Platforms, tea.String("platforms"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Released)) {
		query["released"] = request.Released
	}

	if !tea.BoolValue(util.IsUnset(request.PlatformsShrink)) {
		query["platforms"] = request.PlatformsShrink
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &ListProductVersionsResponse{}
	_body, _err := client.DoROARequest(tea.String("ListProductVersions"), tea.String("2020-12-01"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/api/v1/products/"+tea.StringValue(uid)+"/versions"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetChildrenComponentVersionList(id *string, versionId *string) (_result *GetChildrenComponentVersionListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetChildrenComponentVersionListResponse{}
	_body, _err := client.GetChildrenComponentVersionListWithOptions(id, versionId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetChildrenComponentVersionListWithOptions(id *string, versionId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetChildrenComponentVersionListResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	_result = &GetChildrenComponentVersionListResponse{}
	_body, _err := client.DoROARequest(tea.String("GetChildrenComponentVersionList"), tea.String("2020-12-01"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/integration/api/v1/products/"+tea.StringValue(id)+"/versions/"+tea.StringValue(versionId)+"/children"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateSLR() (_result *CreateSLRResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CreateSLRHeaders{}
	_result = &CreateSLRResponse{}
	_body, _err := client.CreateSLRWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateSLRWithOptions(headers *CreateSLRHeaders, runtime *util.RuntimeOptions) (_result *CreateSLRResponse, _err error) {
	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.ClientToken)) {
		realHeaders["ClientToken"] = headers.ClientToken
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
	}
	_result = &CreateSLRResponse{}
	_body, _err := client.DoROARequest(tea.String("CreateSLR"), tea.String("2020-12-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/api/v1/slr"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetProductVersionRelatedComponentVersionDetail(uid *string, relationUID *string) (_result *GetProductVersionRelatedComponentVersionDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetProductVersionRelatedComponentVersionDetailResponse{}
	_body, _err := client.GetProductVersionRelatedComponentVersionDetailWithOptions(uid, relationUID, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetProductVersionRelatedComponentVersionDetailWithOptions(uid *string, relationUID *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetProductVersionRelatedComponentVersionDetailResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	_result = &GetProductVersionRelatedComponentVersionDetailResponse{}
	_body, _err := client.DoROARequest(tea.String("GetProductVersionRelatedComponentVersionDetail"), tea.String("2020-12-01"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/api/v1/product_versions/"+tea.StringValue(uid)+"/relations/"+tea.StringValue(relationUID)), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddEnvironmentPackage(uid *string, request *AddEnvironmentPackageRequest) (_result *AddEnvironmentPackageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &AddEnvironmentPackageHeaders{}
	_result = &AddEnvironmentPackageResponse{}
	_body, _err := client.AddEnvironmentPackageWithOptions(uid, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddEnvironmentPackageWithOptions(uid *string, request *AddEnvironmentPackageRequest, headers *AddEnvironmentPackageHeaders, runtime *util.RuntimeOptions) (_result *AddEnvironmentPackageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PackageType)) {
		query["packageType"] = request.PackageType
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.ClientToken)) {
		realHeaders["ClientToken"] = headers.ClientToken
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	_result = &AddEnvironmentPackageResponse{}
	_body, _err := client.DoROARequest(tea.String("AddEnvironmentPackage"), tea.String("2020-12-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/api/v1/environments/"+tea.StringValue(uid)+"/package"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateEnvironmentProductVersion(uid *string, request *UpdateEnvironmentProductVersionRequest) (_result *UpdateEnvironmentProductVersionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateEnvironmentProductVersionResponse{}
	_body, _err := client.UpdateEnvironmentProductVersionWithOptions(uid, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateEnvironmentProductVersionWithOptions(uid *string, request *UpdateEnvironmentProductVersionRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateEnvironmentProductVersionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CompatibleVersions)) {
		body["compatibleVersions"] = request.CompatibleVersions
	}

	if !tea.BoolValue(util.IsUnset(request.OldProductVersion)) {
		body["oldProductVersion"] = request.OldProductVersion
	}

	if !tea.BoolValue(util.IsUnset(request.OldProductVersionUID)) {
		body["oldProductVersionUID"] = request.OldProductVersionUID
	}

	if !tea.BoolValue(util.IsUnset(request.ProductName)) {
		body["productName"] = request.ProductName
	}

	if !tea.BoolValue(util.IsUnset(request.ProductUID)) {
		body["productUID"] = request.ProductUID
	}

	if !tea.BoolValue(util.IsUnset(request.ProductVersion)) {
		body["productVersion"] = request.ProductVersion
	}

	if !tea.BoolValue(util.IsUnset(request.ProductVersionUID)) {
		body["productVersionUID"] = request.ProductVersionUID
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &UpdateEnvironmentProductVersionResponse{}
	_body, _err := client.DoROARequest(tea.String("UpdateEnvironmentProductVersion"), tea.String("2020-12-01"), tea.String("HTTPS"), tea.String("PUT"), tea.String("AK"), tea.String("/api/v1/environments/"+tea.StringValue(uid)+"/product_versions"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetProductVersionPlatforms(uid *string) (_result *GetProductVersionPlatformsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetProductVersionPlatformsResponse{}
	_body, _err := client.GetProductVersionPlatformsWithOptions(uid, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetProductVersionPlatformsWithOptions(uid *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetProductVersionPlatformsResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	_result = &GetProductVersionPlatformsResponse{}
	_body, _err := client.DoROARequest(tea.String("GetProductVersionPlatforms"), tea.String("2020-12-01"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/api/v1/product_versions/"+tea.StringValue(uid)+"/platforms"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SaveEnvironmentParam(uid *string, request *SaveEnvironmentParamRequest) (_result *SaveEnvironmentParamResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &SaveEnvironmentParamResponse{}
	_body, _err := client.SaveEnvironmentParamWithOptions(uid, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SaveEnvironmentParamWithOptions(uid *string, request *SaveEnvironmentParamRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *SaveEnvironmentParamResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ComponentUID)) {
		body["componentUID"] = request.ComponentUID
	}

	if !tea.BoolValue(util.IsUnset(request.ComponentVersionUID)) {
		body["componentVersionUID"] = request.ComponentVersionUID
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.ParamUID)) {
		body["paramUID"] = request.ParamUID
	}

	if !tea.BoolValue(util.IsUnset(request.Provider)) {
		body["provider"] = request.Provider
	}

	if !tea.BoolValue(util.IsUnset(request.ReleaseName)) {
		body["releaseName"] = request.ReleaseName
	}

	if !tea.BoolValue(util.IsUnset(request.Scope)) {
		body["scope"] = request.Scope
	}

	if !tea.BoolValue(util.IsUnset(request.Value)) {
		body["value"] = request.Value
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &SaveEnvironmentParamResponse{}
	_body, _err := client.DoROARequest(tea.String("SaveEnvironmentParam"), tea.String("2020-12-01"), tea.String("HTTPS"), tea.String("PUT"), tea.String("AK"), tea.String("/api/v1/environments/"+tea.StringValue(uid)+"/params"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateSnapshotInstanceJoinOption(instanceuid *string, uid *string, request *UpdateSnapshotInstanceJoinOptionRequest) (_result *UpdateSnapshotInstanceJoinOptionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateSnapshotInstanceJoinOptionResponse{}
	_body, _err := client.UpdateSnapshotInstanceJoinOptionWithOptions(instanceuid, uid, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateSnapshotInstanceJoinOptionWithOptions(instanceuid *string, uid *string, request *UpdateSnapshotInstanceJoinOptionRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateSnapshotInstanceJoinOptionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JoinSnapshot)) {
		body["joinSnapshot"] = request.JoinSnapshot
	}

	if !tea.BoolValue(util.IsUnset(request.RootPassword)) {
		body["rootPassword"] = request.RootPassword
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &UpdateSnapshotInstanceJoinOptionResponse{}
	_body, _err := client.DoROARequest(tea.String("UpdateSnapshotInstanceJoinOption"), tea.String("2020-12-01"), tea.String("HTTPS"), tea.String("PUT"), tea.String("AK"), tea.String("/api/v1/snapshots/"+tea.StringValue(uid)+"/instances/"+tea.StringValue(instanceuid)), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetProductVersionResource(uid *string) (_result *GetProductVersionResourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetProductVersionResourceResponse{}
	_body, _err := client.GetProductVersionResourceWithOptions(uid, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetProductVersionResourceWithOptions(uid *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetProductVersionResourceResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	_result = &GetProductVersionResourceResponse{}
	_body, _err := client.DoROARequest(tea.String("GetProductVersionResource"), tea.String("2020-12-01"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/api/v1/product_versions/"+tea.StringValue(uid)+"/resources"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateLicense(uid *string, request *CreateLicenseRequest) (_result *CreateLicenseResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateLicenseResponse{}
	_body, _err := client.CreateLicenseWithOptions(uid, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateLicenseWithOptions(uid *string, request *CreateLicenseRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateLicenseResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EffectiveYear)) {
		query["effectiveYear"] = request.EffectiveYear
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &CreateLicenseResponse{}
	_body, _err := client.DoROARequest(tea.String("CreateLicense"), tea.String("2020-12-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/api/v1/environments/"+tea.StringValue(uid)+"/licenses"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ShareSnapshot(uid *string, request *ShareSnapshotRequest) (_result *ShareSnapshotResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ShareSnapshotResponse{}
	_body, _err := client.ShareSnapshotWithOptions(uid, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ShareSnapshotWithOptions(uid *string, request *ShareSnapshotRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ShareSnapshotResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TargetProvider)) {
		body["targetProvider"] = request.TargetProvider
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &ShareSnapshotResponse{}
	_body, _err := client.DoROARequest(tea.String("ShareSnapshot"), tea.String("2020-12-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/api/v1/snapshots/"+tea.StringValue(uid)+"/snapshots"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteEnvironmentParam(uid *string) (_result *DeleteEnvironmentParamResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteEnvironmentParamResponse{}
	_body, _err := client.DeleteEnvironmentParamWithOptions(uid, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteEnvironmentParamWithOptions(uid *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteEnvironmentParamResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	_result = &DeleteEnvironmentParamResponse{}
	_body, _err := client.DoROARequest(tea.String("DeleteEnvironmentParam"), tea.String("2020-12-01"), tea.String("HTTPS"), tea.String("DELETE"), tea.String("AK"), tea.String("/api/v1/environmentparams/"+tea.StringValue(uid)), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetProduct(uid *string) (_result *GetProductResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetProductResponse{}
	_body, _err := client.GetProductWithOptions(uid, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetProductWithOptions(uid *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetProductResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	_result = &GetProductResponse{}
	_body, _err := client.DoROARequest(tea.String("GetProduct"), tea.String("2020-12-01"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/api/v1/products/"+tea.StringValue(uid)), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteComponentVersion(uid *string, versionUid *string) (_result *DeleteComponentVersionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteComponentVersionResponse{}
	_body, _err := client.DeleteComponentVersionWithOptions(uid, versionUid, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteComponentVersionWithOptions(uid *string, versionUid *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteComponentVersionResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	_result = &DeleteComponentVersionResponse{}
	_body, _err := client.DoROARequest(tea.String("DeleteComponentVersion"), tea.String("2020-12-01"), tea.String("HTTPS"), tea.String("DELETE"), tea.String("AK"), tea.String("/integration/api/v1/components/"+tea.StringValue(uid)+"/versions/"+tea.StringValue(versionUid)), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeployEnvironmentProduct(uid *string) (_result *DeployEnvironmentProductResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeployEnvironmentProductResponse{}
	_body, _err := client.DeployEnvironmentProductWithOptions(uid, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeployEnvironmentProductWithOptions(uid *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeployEnvironmentProductResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	_result = &DeployEnvironmentProductResponse{}
	_body, _err := client.DoROARequest(tea.String("DeployEnvironmentProduct"), tea.String("2020-12-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/api/v1/environments/"+tea.StringValue(uid)+"/deployment"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) InitEnvironmentResource(uid *string, request *InitEnvironmentResourceRequest) (_result *InitEnvironmentResourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &InitEnvironmentResourceResponse{}
	_body, _err := client.InitEnvironmentResourceWithOptions(uid, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) InitEnvironmentResourceWithOptions(uid *string, request *InitEnvironmentResourceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *InitEnvironmentResourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessKeyID)) {
		body["accessKeyID"] = request.AccessKeyID
	}

	if !tea.BoolValue(util.IsUnset(request.AccessKeySecret)) {
		body["accessKeySecret"] = request.AccessKeySecret
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		body["securityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &InitEnvironmentResourceResponse{}
	_body, _err := client.DoROARequest(tea.String("InitEnvironmentResource"), tea.String("2020-12-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/api/v1/environments/"+tea.StringValue(uid)+"/resources"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListFoundationVersions() (_result *ListFoundationVersionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListFoundationVersionsResponse{}
	_body, _err := client.ListFoundationVersionsWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListFoundationVersionsWithOptions(headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListFoundationVersionsResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	_result = &ListFoundationVersionsResponse{}
	_body, _err := client.DoROARequest(tea.String("ListFoundationVersions"), tea.String("2020-12-01"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/api/v1/foundation/versions"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListEnvironmentDeployRecord(uid *string, request *ListEnvironmentDeployRecordRequest) (_result *ListEnvironmentDeployRecordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListEnvironmentDeployRecordResponse{}
	_body, _err := client.ListEnvironmentDeployRecordWithOptions(uid, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListEnvironmentDeployRecordWithOptions(uid *string, request *ListEnvironmentDeployRecordRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListEnvironmentDeployRecordResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNum)) {
		query["pageNum"] = request.PageNum
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &ListEnvironmentDeployRecordResponse{}
	_body, _err := client.DoROARequest(tea.String("ListEnvironmentDeployRecord"), tea.String("2020-12-01"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/api/v1/environments/"+tea.StringValue(uid)+"/deployments"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}
