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

type CodeSourceItem struct {
	Accessibility       *string `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	CodeBranch          *string `json:"CodeBranch,omitempty" xml:"CodeBranch,omitempty"`
	CodeCommit          *string `json:"CodeCommit,omitempty" xml:"CodeCommit,omitempty"`
	CodeRepo            *string `json:"CodeRepo,omitempty" xml:"CodeRepo,omitempty"`
	CodeRepoAccessToken *string `json:"CodeRepoAccessToken,omitempty" xml:"CodeRepoAccessToken,omitempty"`
	CodeRepoUserName    *string `json:"CodeRepoUserName,omitempty" xml:"CodeRepoUserName,omitempty"`
	CodeSourceId        *string `json:"CodeSourceId,omitempty" xml:"CodeSourceId,omitempty"`
	Description         *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DisplayName         *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	GmtCreateTime       *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	GmtModifyTime       *string `json:"GmtModifyTime,omitempty" xml:"GmtModifyTime,omitempty"`
	MountPath           *string `json:"MountPath,omitempty" xml:"MountPath,omitempty"`
	UserId              *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	WorkspaceId         *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s CodeSourceItem) String() string {
	return tea.Prettify(s)
}

func (s CodeSourceItem) GoString() string {
	return s.String()
}

func (s *CodeSourceItem) SetAccessibility(v string) *CodeSourceItem {
	s.Accessibility = &v
	return s
}

func (s *CodeSourceItem) SetCodeBranch(v string) *CodeSourceItem {
	s.CodeBranch = &v
	return s
}

func (s *CodeSourceItem) SetCodeCommit(v string) *CodeSourceItem {
	s.CodeCommit = &v
	return s
}

func (s *CodeSourceItem) SetCodeRepo(v string) *CodeSourceItem {
	s.CodeRepo = &v
	return s
}

func (s *CodeSourceItem) SetCodeRepoAccessToken(v string) *CodeSourceItem {
	s.CodeRepoAccessToken = &v
	return s
}

func (s *CodeSourceItem) SetCodeRepoUserName(v string) *CodeSourceItem {
	s.CodeRepoUserName = &v
	return s
}

func (s *CodeSourceItem) SetCodeSourceId(v string) *CodeSourceItem {
	s.CodeSourceId = &v
	return s
}

func (s *CodeSourceItem) SetDescription(v string) *CodeSourceItem {
	s.Description = &v
	return s
}

func (s *CodeSourceItem) SetDisplayName(v string) *CodeSourceItem {
	s.DisplayName = &v
	return s
}

func (s *CodeSourceItem) SetGmtCreateTime(v string) *CodeSourceItem {
	s.GmtCreateTime = &v
	return s
}

func (s *CodeSourceItem) SetGmtModifyTime(v string) *CodeSourceItem {
	s.GmtModifyTime = &v
	return s
}

func (s *CodeSourceItem) SetMountPath(v string) *CodeSourceItem {
	s.MountPath = &v
	return s
}

func (s *CodeSourceItem) SetUserId(v string) *CodeSourceItem {
	s.UserId = &v
	return s
}

func (s *CodeSourceItem) SetWorkspaceId(v string) *CodeSourceItem {
	s.WorkspaceId = &v
	return s
}

type Collection struct {
	CollectionName  *string `json:"CollectionName,omitempty" xml:"CollectionName,omitempty"`
	GmtCreateTime   *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	OwnerId         *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	UserId          *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s Collection) String() string {
	return tea.Prettify(s)
}

func (s Collection) GoString() string {
	return s.String()
}

func (s *Collection) SetCollectionName(v string) *Collection {
	s.CollectionName = &v
	return s
}

func (s *Collection) SetGmtCreateTime(v string) *Collection {
	s.GmtCreateTime = &v
	return s
}

func (s *Collection) SetGmtModifiedTime(v string) *Collection {
	s.GmtModifiedTime = &v
	return s
}

func (s *Collection) SetOwnerId(v string) *Collection {
	s.OwnerId = &v
	return s
}

func (s *Collection) SetUserId(v string) *Collection {
	s.UserId = &v
	return s
}

type Dataset struct {
	Accessibility   *string  `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	DataSourceType  *string  `json:"DataSourceType,omitempty" xml:"DataSourceType,omitempty"`
	DataType        *string  `json:"DataType,omitempty" xml:"DataType,omitempty"`
	DatasetId       *string  `json:"DatasetId,omitempty" xml:"DatasetId,omitempty"`
	Description     *string  `json:"Description,omitempty" xml:"Description,omitempty"`
	GmtCreateTime   *string  `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	GmtModifiedTime *string  `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	Labels          []*Label `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	Name            *string  `json:"Name,omitempty" xml:"Name,omitempty"`
	Options         *string  `json:"Options,omitempty" xml:"Options,omitempty"`
	OwnerId         *string  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	Property        *string  `json:"Property,omitempty" xml:"Property,omitempty"`
	ProviderType    *string  `json:"ProviderType,omitempty" xml:"ProviderType,omitempty"`
	SourceId        *string  `json:"SourceId,omitempty" xml:"SourceId,omitempty"`
	SourceType      *string  `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	Uri             *string  `json:"Uri,omitempty" xml:"Uri,omitempty"`
	UserId          *string  `json:"UserId,omitempty" xml:"UserId,omitempty"`
	WorkspaceId     *string  `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s Dataset) String() string {
	return tea.Prettify(s)
}

func (s Dataset) GoString() string {
	return s.String()
}

func (s *Dataset) SetAccessibility(v string) *Dataset {
	s.Accessibility = &v
	return s
}

func (s *Dataset) SetDataSourceType(v string) *Dataset {
	s.DataSourceType = &v
	return s
}

func (s *Dataset) SetDataType(v string) *Dataset {
	s.DataType = &v
	return s
}

func (s *Dataset) SetDatasetId(v string) *Dataset {
	s.DatasetId = &v
	return s
}

func (s *Dataset) SetDescription(v string) *Dataset {
	s.Description = &v
	return s
}

func (s *Dataset) SetGmtCreateTime(v string) *Dataset {
	s.GmtCreateTime = &v
	return s
}

func (s *Dataset) SetGmtModifiedTime(v string) *Dataset {
	s.GmtModifiedTime = &v
	return s
}

func (s *Dataset) SetLabels(v []*Label) *Dataset {
	s.Labels = v
	return s
}

func (s *Dataset) SetName(v string) *Dataset {
	s.Name = &v
	return s
}

func (s *Dataset) SetOptions(v string) *Dataset {
	s.Options = &v
	return s
}

func (s *Dataset) SetOwnerId(v string) *Dataset {
	s.OwnerId = &v
	return s
}

func (s *Dataset) SetProperty(v string) *Dataset {
	s.Property = &v
	return s
}

func (s *Dataset) SetProviderType(v string) *Dataset {
	s.ProviderType = &v
	return s
}

func (s *Dataset) SetSourceId(v string) *Dataset {
	s.SourceId = &v
	return s
}

func (s *Dataset) SetSourceType(v string) *Dataset {
	s.SourceType = &v
	return s
}

func (s *Dataset) SetUri(v string) *Dataset {
	s.Uri = &v
	return s
}

func (s *Dataset) SetUserId(v string) *Dataset {
	s.UserId = &v
	return s
}

func (s *Dataset) SetWorkspaceId(v string) *Dataset {
	s.WorkspaceId = &v
	return s
}

type DatasetLabel struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DatasetLabel) String() string {
	return tea.Prettify(s)
}

func (s DatasetLabel) GoString() string {
	return s.String()
}

func (s *DatasetLabel) SetKey(v string) *DatasetLabel {
	s.Key = &v
	return s
}

func (s *DatasetLabel) SetValue(v string) *DatasetLabel {
	s.Value = &v
	return s
}

type Experiment struct {
	ArtifactUri       *string                  `json:"ArtifactUri,omitempty" xml:"ArtifactUri,omitempty"`
	ExperimentId      *string                  `json:"ExperimentId,omitempty" xml:"ExperimentId,omitempty"`
	GmtCreateTime     *string                  `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	GmtModifiedTime   *string                  `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	Labels            []map[string]interface{} `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	Name              *string                  `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerId           *string                  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	TensorboardLogUri *string                  `json:"TensorboardLogUri,omitempty" xml:"TensorboardLogUri,omitempty"`
	UserId            *string                  `json:"UserId,omitempty" xml:"UserId,omitempty"`
	WorkspaceId       *string                  `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s Experiment) String() string {
	return tea.Prettify(s)
}

func (s Experiment) GoString() string {
	return s.String()
}

func (s *Experiment) SetArtifactUri(v string) *Experiment {
	s.ArtifactUri = &v
	return s
}

func (s *Experiment) SetExperimentId(v string) *Experiment {
	s.ExperimentId = &v
	return s
}

func (s *Experiment) SetGmtCreateTime(v string) *Experiment {
	s.GmtCreateTime = &v
	return s
}

func (s *Experiment) SetGmtModifiedTime(v string) *Experiment {
	s.GmtModifiedTime = &v
	return s
}

func (s *Experiment) SetLabels(v []map[string]interface{}) *Experiment {
	s.Labels = v
	return s
}

func (s *Experiment) SetName(v string) *Experiment {
	s.Name = &v
	return s
}

func (s *Experiment) SetOwnerId(v string) *Experiment {
	s.OwnerId = &v
	return s
}

func (s *Experiment) SetTensorboardLogUri(v string) *Experiment {
	s.TensorboardLogUri = &v
	return s
}

func (s *Experiment) SetUserId(v string) *Experiment {
	s.UserId = &v
	return s
}

func (s *Experiment) SetWorkspaceId(v string) *Experiment {
	s.WorkspaceId = &v
	return s
}

type ExperimentLabel struct {
	ExperimentId    *string `json:"ExperimentId,omitempty" xml:"ExperimentId,omitempty"`
	GmtCreateTime   *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	Key             *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value           *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ExperimentLabel) String() string {
	return tea.Prettify(s)
}

func (s ExperimentLabel) GoString() string {
	return s.String()
}

func (s *ExperimentLabel) SetExperimentId(v string) *ExperimentLabel {
	s.ExperimentId = &v
	return s
}

func (s *ExperimentLabel) SetGmtCreateTime(v string) *ExperimentLabel {
	s.GmtCreateTime = &v
	return s
}

func (s *ExperimentLabel) SetGmtModifiedTime(v string) *ExperimentLabel {
	s.GmtModifiedTime = &v
	return s
}

func (s *ExperimentLabel) SetKey(v string) *ExperimentLabel {
	s.Key = &v
	return s
}

func (s *ExperimentLabel) SetValue(v string) *ExperimentLabel {
	s.Value = &v
	return s
}

type Label struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s Label) String() string {
	return tea.Prettify(s)
}

func (s Label) GoString() string {
	return s.String()
}

func (s *Label) SetKey(v string) *Label {
	s.Key = &v
	return s
}

func (s *Label) SetValue(v string) *Label {
	s.Value = &v
	return s
}

type LabelInfo struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s LabelInfo) String() string {
	return tea.Prettify(s)
}

func (s LabelInfo) GoString() string {
	return s.String()
}

func (s *LabelInfo) SetKey(v string) *LabelInfo {
	s.Key = &v
	return s
}

func (s *LabelInfo) SetValue(v string) *LabelInfo {
	s.Value = &v
	return s
}

type Model struct {
	Accessibility    *string                `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	Domain           *string                `json:"Domain,omitempty" xml:"Domain,omitempty"`
	ExtraInfo        map[string]interface{} `json:"ExtraInfo,omitempty" xml:"ExtraInfo,omitempty"`
	GmtCreateTime    *string                `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	GmtModifiedTime  *string                `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	Labels           []*Label               `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	LatestVersion    *ModelVersion          `json:"LatestVersion,omitempty" xml:"LatestVersion,omitempty"`
	ModelDescription *string                `json:"ModelDescription,omitempty" xml:"ModelDescription,omitempty"`
	ModelDoc         *string                `json:"ModelDoc,omitempty" xml:"ModelDoc,omitempty"`
	ModelId          *string                `json:"ModelId,omitempty" xml:"ModelId,omitempty"`
	ModelName        *string                `json:"ModelName,omitempty" xml:"ModelName,omitempty"`
	ModelType        *string                `json:"ModelType,omitempty" xml:"ModelType,omitempty"`
	OrderNumber      *int64                 `json:"OrderNumber,omitempty" xml:"OrderNumber,omitempty"`
	Origin           *string                `json:"Origin,omitempty" xml:"Origin,omitempty"`
	OwnerId          *string                `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	Provider         *string                `json:"Provider,omitempty" xml:"Provider,omitempty"`
	Task             *string                `json:"Task,omitempty" xml:"Task,omitempty"`
	UserId           *string                `json:"UserId,omitempty" xml:"UserId,omitempty"`
	WorkspaceId      *string                `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s Model) String() string {
	return tea.Prettify(s)
}

func (s Model) GoString() string {
	return s.String()
}

func (s *Model) SetAccessibility(v string) *Model {
	s.Accessibility = &v
	return s
}

func (s *Model) SetDomain(v string) *Model {
	s.Domain = &v
	return s
}

func (s *Model) SetExtraInfo(v map[string]interface{}) *Model {
	s.ExtraInfo = v
	return s
}

func (s *Model) SetGmtCreateTime(v string) *Model {
	s.GmtCreateTime = &v
	return s
}

func (s *Model) SetGmtModifiedTime(v string) *Model {
	s.GmtModifiedTime = &v
	return s
}

func (s *Model) SetLabels(v []*Label) *Model {
	s.Labels = v
	return s
}

func (s *Model) SetLatestVersion(v *ModelVersion) *Model {
	s.LatestVersion = v
	return s
}

func (s *Model) SetModelDescription(v string) *Model {
	s.ModelDescription = &v
	return s
}

func (s *Model) SetModelDoc(v string) *Model {
	s.ModelDoc = &v
	return s
}

func (s *Model) SetModelId(v string) *Model {
	s.ModelId = &v
	return s
}

func (s *Model) SetModelName(v string) *Model {
	s.ModelName = &v
	return s
}

func (s *Model) SetModelType(v string) *Model {
	s.ModelType = &v
	return s
}

func (s *Model) SetOrderNumber(v int64) *Model {
	s.OrderNumber = &v
	return s
}

func (s *Model) SetOrigin(v string) *Model {
	s.Origin = &v
	return s
}

func (s *Model) SetOwnerId(v string) *Model {
	s.OwnerId = &v
	return s
}

func (s *Model) SetProvider(v string) *Model {
	s.Provider = &v
	return s
}

func (s *Model) SetTask(v string) *Model {
	s.Task = &v
	return s
}

func (s *Model) SetUserId(v string) *Model {
	s.UserId = &v
	return s
}

func (s *Model) SetWorkspaceId(v string) *Model {
	s.WorkspaceId = &v
	return s
}

type ModelVersion struct {
	ApprovalStatus     *string                `json:"ApprovalStatus,omitempty" xml:"ApprovalStatus,omitempty"`
	EvaluationSpec     map[string]interface{} `json:"EvaluationSpec,omitempty" xml:"EvaluationSpec,omitempty"`
	ExtraInfo          map[string]interface{} `json:"ExtraInfo,omitempty" xml:"ExtraInfo,omitempty"`
	FormatType         *string                `json:"FormatType,omitempty" xml:"FormatType,omitempty"`
	FrameworkType      *string                `json:"FrameworkType,omitempty" xml:"FrameworkType,omitempty"`
	GmtCreateTime      *string                `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	GmtModifiedTime    *string                `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	InferenceSpec      map[string]interface{} `json:"InferenceSpec,omitempty" xml:"InferenceSpec,omitempty"`
	Labels             []*Label               `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	Metrics            map[string]interface{} `json:"Metrics,omitempty" xml:"Metrics,omitempty"`
	Options            *string                `json:"Options,omitempty" xml:"Options,omitempty"`
	OwnerId            *string                `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SourceId           *string                `json:"SourceId,omitempty" xml:"SourceId,omitempty"`
	SourceType         *string                `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	TrainingSpec       map[string]interface{} `json:"TrainingSpec,omitempty" xml:"TrainingSpec,omitempty"`
	Uri                *string                `json:"Uri,omitempty" xml:"Uri,omitempty"`
	UserId             *string                `json:"UserId,omitempty" xml:"UserId,omitempty"`
	VersionDescription *string                `json:"VersionDescription,omitempty" xml:"VersionDescription,omitempty"`
	VersionName        *string                `json:"VersionName,omitempty" xml:"VersionName,omitempty"`
}

func (s ModelVersion) String() string {
	return tea.Prettify(s)
}

func (s ModelVersion) GoString() string {
	return s.String()
}

func (s *ModelVersion) SetApprovalStatus(v string) *ModelVersion {
	s.ApprovalStatus = &v
	return s
}

func (s *ModelVersion) SetEvaluationSpec(v map[string]interface{}) *ModelVersion {
	s.EvaluationSpec = v
	return s
}

func (s *ModelVersion) SetExtraInfo(v map[string]interface{}) *ModelVersion {
	s.ExtraInfo = v
	return s
}

func (s *ModelVersion) SetFormatType(v string) *ModelVersion {
	s.FormatType = &v
	return s
}

func (s *ModelVersion) SetFrameworkType(v string) *ModelVersion {
	s.FrameworkType = &v
	return s
}

func (s *ModelVersion) SetGmtCreateTime(v string) *ModelVersion {
	s.GmtCreateTime = &v
	return s
}

func (s *ModelVersion) SetGmtModifiedTime(v string) *ModelVersion {
	s.GmtModifiedTime = &v
	return s
}

func (s *ModelVersion) SetInferenceSpec(v map[string]interface{}) *ModelVersion {
	s.InferenceSpec = v
	return s
}

func (s *ModelVersion) SetLabels(v []*Label) *ModelVersion {
	s.Labels = v
	return s
}

func (s *ModelVersion) SetMetrics(v map[string]interface{}) *ModelVersion {
	s.Metrics = v
	return s
}

func (s *ModelVersion) SetOptions(v string) *ModelVersion {
	s.Options = &v
	return s
}

func (s *ModelVersion) SetOwnerId(v string) *ModelVersion {
	s.OwnerId = &v
	return s
}

func (s *ModelVersion) SetSourceId(v string) *ModelVersion {
	s.SourceId = &v
	return s
}

func (s *ModelVersion) SetSourceType(v string) *ModelVersion {
	s.SourceType = &v
	return s
}

func (s *ModelVersion) SetTrainingSpec(v map[string]interface{}) *ModelVersion {
	s.TrainingSpec = v
	return s
}

func (s *ModelVersion) SetUri(v string) *ModelVersion {
	s.Uri = &v
	return s
}

func (s *ModelVersion) SetUserId(v string) *ModelVersion {
	s.UserId = &v
	return s
}

func (s *ModelVersion) SetVersionDescription(v string) *ModelVersion {
	s.VersionDescription = &v
	return s
}

func (s *ModelVersion) SetVersionName(v string) *ModelVersion {
	s.VersionName = &v
	return s
}

type ServiceTemplate struct {
	GmtCreateTime              *string                `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	GmtModifiedTime            *string                `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	InferenceSpec              map[string]interface{} `json:"InferenceSpec,omitempty" xml:"InferenceSpec,omitempty"`
	Labels                     []*Label               `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	OrderNumber                *int64                 `json:"OrderNumber,omitempty" xml:"OrderNumber,omitempty"`
	OwnerId                    *string                `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	Provider                   *string                `json:"Provider,omitempty" xml:"Provider,omitempty"`
	ServiceTemplateDescription *string                `json:"ServiceTemplateDescription,omitempty" xml:"ServiceTemplateDescription,omitempty"`
	ServiceTemplateDoc         *string                `json:"ServiceTemplateDoc,omitempty" xml:"ServiceTemplateDoc,omitempty"`
	ServiceTemplateId          *string                `json:"ServiceTemplateId,omitempty" xml:"ServiceTemplateId,omitempty"`
	ServiceTemplateName        *string                `json:"ServiceTemplateName,omitempty" xml:"ServiceTemplateName,omitempty"`
	UserId                     *string                `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ServiceTemplate) String() string {
	return tea.Prettify(s)
}

func (s ServiceTemplate) GoString() string {
	return s.String()
}

func (s *ServiceTemplate) SetGmtCreateTime(v string) *ServiceTemplate {
	s.GmtCreateTime = &v
	return s
}

func (s *ServiceTemplate) SetGmtModifiedTime(v string) *ServiceTemplate {
	s.GmtModifiedTime = &v
	return s
}

func (s *ServiceTemplate) SetInferenceSpec(v map[string]interface{}) *ServiceTemplate {
	s.InferenceSpec = v
	return s
}

func (s *ServiceTemplate) SetLabels(v []*Label) *ServiceTemplate {
	s.Labels = v
	return s
}

func (s *ServiceTemplate) SetOrderNumber(v int64) *ServiceTemplate {
	s.OrderNumber = &v
	return s
}

func (s *ServiceTemplate) SetOwnerId(v string) *ServiceTemplate {
	s.OwnerId = &v
	return s
}

func (s *ServiceTemplate) SetProvider(v string) *ServiceTemplate {
	s.Provider = &v
	return s
}

func (s *ServiceTemplate) SetServiceTemplateDescription(v string) *ServiceTemplate {
	s.ServiceTemplateDescription = &v
	return s
}

func (s *ServiceTemplate) SetServiceTemplateDoc(v string) *ServiceTemplate {
	s.ServiceTemplateDoc = &v
	return s
}

func (s *ServiceTemplate) SetServiceTemplateId(v string) *ServiceTemplate {
	s.ServiceTemplateId = &v
	return s
}

func (s *ServiceTemplate) SetServiceTemplateName(v string) *ServiceTemplate {
	s.ServiceTemplateName = &v
	return s
}

func (s *ServiceTemplate) SetUserId(v string) *ServiceTemplate {
	s.UserId = &v
	return s
}

type Trial struct {
	Accessibility   *string                  `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	ExperimentId    *string                  `json:"ExperimentId,omitempty" xml:"ExperimentId,omitempty"`
	GmtCreateTime   *string                  `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	GmtModifiedTime *string                  `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	Labels          []map[string]interface{} `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	Name            *string                  `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerId         *string                  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SourceId        *string                  `json:"SourceId,omitempty" xml:"SourceId,omitempty"`
	SourceType      *string                  `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	TrialId         *string                  `json:"TrialId,omitempty" xml:"TrialId,omitempty"`
	UserId          *string                  `json:"UserId,omitempty" xml:"UserId,omitempty"`
	WorkspaceId     *string                  `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s Trial) String() string {
	return tea.Prettify(s)
}

func (s Trial) GoString() string {
	return s.String()
}

func (s *Trial) SetAccessibility(v string) *Trial {
	s.Accessibility = &v
	return s
}

func (s *Trial) SetExperimentId(v string) *Trial {
	s.ExperimentId = &v
	return s
}

func (s *Trial) SetGmtCreateTime(v string) *Trial {
	s.GmtCreateTime = &v
	return s
}

func (s *Trial) SetGmtModifiedTime(v string) *Trial {
	s.GmtModifiedTime = &v
	return s
}

func (s *Trial) SetLabels(v []map[string]interface{}) *Trial {
	s.Labels = v
	return s
}

func (s *Trial) SetName(v string) *Trial {
	s.Name = &v
	return s
}

func (s *Trial) SetOwnerId(v string) *Trial {
	s.OwnerId = &v
	return s
}

func (s *Trial) SetSourceId(v string) *Trial {
	s.SourceId = &v
	return s
}

func (s *Trial) SetSourceType(v string) *Trial {
	s.SourceType = &v
	return s
}

func (s *Trial) SetTrialId(v string) *Trial {
	s.TrialId = &v
	return s
}

func (s *Trial) SetUserId(v string) *Trial {
	s.UserId = &v
	return s
}

func (s *Trial) SetWorkspaceId(v string) *Trial {
	s.WorkspaceId = &v
	return s
}

type TrialLabel struct {
	GmtCreateTime   *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	Key             *string `json:"Key,omitempty" xml:"Key,omitempty"`
	TrialId         *string `json:"TrialId,omitempty" xml:"TrialId,omitempty"`
	Value           *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s TrialLabel) String() string {
	return tea.Prettify(s)
}

func (s TrialLabel) GoString() string {
	return s.String()
}

func (s *TrialLabel) SetGmtCreateTime(v string) *TrialLabel {
	s.GmtCreateTime = &v
	return s
}

func (s *TrialLabel) SetGmtModifiedTime(v string) *TrialLabel {
	s.GmtModifiedTime = &v
	return s
}

func (s *TrialLabel) SetKey(v string) *TrialLabel {
	s.Key = &v
	return s
}

func (s *TrialLabel) SetTrialId(v string) *TrialLabel {
	s.TrialId = &v
	return s
}

func (s *TrialLabel) SetValue(v string) *TrialLabel {
	s.Value = &v
	return s
}

type AddImageRequest struct {
	Accessibility *string                  `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	Description   *string                  `json:"Description,omitempty" xml:"Description,omitempty"`
	ImageId       *string                  `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	ImageUri      *string                  `json:"ImageUri,omitempty" xml:"ImageUri,omitempty"`
	Labels        []*AddImageRequestLabels `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	Name          *string                  `json:"Name,omitempty" xml:"Name,omitempty"`
	Size          *int32                   `json:"Size,omitempty" xml:"Size,omitempty"`
	WorkspaceId   *string                  `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s AddImageRequest) String() string {
	return tea.Prettify(s)
}

func (s AddImageRequest) GoString() string {
	return s.String()
}

func (s *AddImageRequest) SetAccessibility(v string) *AddImageRequest {
	s.Accessibility = &v
	return s
}

func (s *AddImageRequest) SetDescription(v string) *AddImageRequest {
	s.Description = &v
	return s
}

func (s *AddImageRequest) SetImageId(v string) *AddImageRequest {
	s.ImageId = &v
	return s
}

func (s *AddImageRequest) SetImageUri(v string) *AddImageRequest {
	s.ImageUri = &v
	return s
}

func (s *AddImageRequest) SetLabels(v []*AddImageRequestLabels) *AddImageRequest {
	s.Labels = v
	return s
}

func (s *AddImageRequest) SetName(v string) *AddImageRequest {
	s.Name = &v
	return s
}

func (s *AddImageRequest) SetSize(v int32) *AddImageRequest {
	s.Size = &v
	return s
}

func (s *AddImageRequest) SetWorkspaceId(v string) *AddImageRequest {
	s.WorkspaceId = &v
	return s
}

type AddImageRequestLabels struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s AddImageRequestLabels) String() string {
	return tea.Prettify(s)
}

func (s AddImageRequestLabels) GoString() string {
	return s.String()
}

func (s *AddImageRequestLabels) SetKey(v string) *AddImageRequestLabels {
	s.Key = &v
	return s
}

func (s *AddImageRequestLabels) SetValue(v string) *AddImageRequestLabels {
	s.Value = &v
	return s
}

type AddImageResponseBody struct {
	ImageId   *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddImageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddImageResponseBody) GoString() string {
	return s.String()
}

func (s *AddImageResponseBody) SetImageId(v string) *AddImageResponseBody {
	s.ImageId = &v
	return s
}

func (s *AddImageResponseBody) SetRequestId(v string) *AddImageResponseBody {
	s.RequestId = &v
	return s
}

type AddImageResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddImageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddImageResponse) String() string {
	return tea.Prettify(s)
}

func (s AddImageResponse) GoString() string {
	return s.String()
}

func (s *AddImageResponse) SetHeaders(v map[string]*string) *AddImageResponse {
	s.Headers = v
	return s
}

func (s *AddImageResponse) SetStatusCode(v int32) *AddImageResponse {
	s.StatusCode = &v
	return s
}

func (s *AddImageResponse) SetBody(v *AddImageResponseBody) *AddImageResponse {
	s.Body = v
	return s
}

type AddImageLabelsRequest struct {
	Labels []*AddImageLabelsRequestLabels `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
}

func (s AddImageLabelsRequest) String() string {
	return tea.Prettify(s)
}

func (s AddImageLabelsRequest) GoString() string {
	return s.String()
}

func (s *AddImageLabelsRequest) SetLabels(v []*AddImageLabelsRequestLabels) *AddImageLabelsRequest {
	s.Labels = v
	return s
}

type AddImageLabelsRequestLabels struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s AddImageLabelsRequestLabels) String() string {
	return tea.Prettify(s)
}

func (s AddImageLabelsRequestLabels) GoString() string {
	return s.String()
}

func (s *AddImageLabelsRequestLabels) SetKey(v string) *AddImageLabelsRequestLabels {
	s.Key = &v
	return s
}

func (s *AddImageLabelsRequestLabels) SetValue(v string) *AddImageLabelsRequestLabels {
	s.Value = &v
	return s
}

type AddImageLabelsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddImageLabelsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddImageLabelsResponseBody) GoString() string {
	return s.String()
}

func (s *AddImageLabelsResponseBody) SetRequestId(v string) *AddImageLabelsResponseBody {
	s.RequestId = &v
	return s
}

type AddImageLabelsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddImageLabelsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddImageLabelsResponse) String() string {
	return tea.Prettify(s)
}

func (s AddImageLabelsResponse) GoString() string {
	return s.String()
}

func (s *AddImageLabelsResponse) SetHeaders(v map[string]*string) *AddImageLabelsResponse {
	s.Headers = v
	return s
}

func (s *AddImageLabelsResponse) SetStatusCode(v int32) *AddImageLabelsResponse {
	s.StatusCode = &v
	return s
}

func (s *AddImageLabelsResponse) SetBody(v *AddImageLabelsResponseBody) *AddImageLabelsResponse {
	s.Body = v
	return s
}

type AddMemberRoleResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddMemberRoleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddMemberRoleResponseBody) GoString() string {
	return s.String()
}

func (s *AddMemberRoleResponseBody) SetRequestId(v string) *AddMemberRoleResponseBody {
	s.RequestId = &v
	return s
}

type AddMemberRoleResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddMemberRoleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddMemberRoleResponse) String() string {
	return tea.Prettify(s)
}

func (s AddMemberRoleResponse) GoString() string {
	return s.String()
}

func (s *AddMemberRoleResponse) SetHeaders(v map[string]*string) *AddMemberRoleResponse {
	s.Headers = v
	return s
}

func (s *AddMemberRoleResponse) SetStatusCode(v int32) *AddMemberRoleResponse {
	s.StatusCode = &v
	return s
}

func (s *AddMemberRoleResponse) SetBody(v *AddMemberRoleResponseBody) *AddMemberRoleResponse {
	s.Body = v
	return s
}

type CreateCodeSourceRequest struct {
	Accessibility       *string `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	CodeBranch          *string `json:"CodeBranch,omitempty" xml:"CodeBranch,omitempty"`
	CodeRepo            *string `json:"CodeRepo,omitempty" xml:"CodeRepo,omitempty"`
	CodeRepoAccessToken *string `json:"CodeRepoAccessToken,omitempty" xml:"CodeRepoAccessToken,omitempty"`
	CodeRepoUserName    *string `json:"CodeRepoUserName,omitempty" xml:"CodeRepoUserName,omitempty"`
	Description         *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DisplayName         *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	MountPath           *string `json:"MountPath,omitempty" xml:"MountPath,omitempty"`
	WorkspaceId         *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s CreateCodeSourceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateCodeSourceRequest) GoString() string {
	return s.String()
}

func (s *CreateCodeSourceRequest) SetAccessibility(v string) *CreateCodeSourceRequest {
	s.Accessibility = &v
	return s
}

func (s *CreateCodeSourceRequest) SetCodeBranch(v string) *CreateCodeSourceRequest {
	s.CodeBranch = &v
	return s
}

func (s *CreateCodeSourceRequest) SetCodeRepo(v string) *CreateCodeSourceRequest {
	s.CodeRepo = &v
	return s
}

func (s *CreateCodeSourceRequest) SetCodeRepoAccessToken(v string) *CreateCodeSourceRequest {
	s.CodeRepoAccessToken = &v
	return s
}

func (s *CreateCodeSourceRequest) SetCodeRepoUserName(v string) *CreateCodeSourceRequest {
	s.CodeRepoUserName = &v
	return s
}

func (s *CreateCodeSourceRequest) SetDescription(v string) *CreateCodeSourceRequest {
	s.Description = &v
	return s
}

func (s *CreateCodeSourceRequest) SetDisplayName(v string) *CreateCodeSourceRequest {
	s.DisplayName = &v
	return s
}

func (s *CreateCodeSourceRequest) SetMountPath(v string) *CreateCodeSourceRequest {
	s.MountPath = &v
	return s
}

func (s *CreateCodeSourceRequest) SetWorkspaceId(v string) *CreateCodeSourceRequest {
	s.WorkspaceId = &v
	return s
}

type CreateCodeSourceResponseBody struct {
	CodeSourceId *string `json:"CodeSourceId,omitempty" xml:"CodeSourceId,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateCodeSourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateCodeSourceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCodeSourceResponseBody) SetCodeSourceId(v string) *CreateCodeSourceResponseBody {
	s.CodeSourceId = &v
	return s
}

func (s *CreateCodeSourceResponseBody) SetRequestId(v string) *CreateCodeSourceResponseBody {
	s.RequestId = &v
	return s
}

type CreateCodeSourceResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateCodeSourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateCodeSourceResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateCodeSourceResponse) GoString() string {
	return s.String()
}

func (s *CreateCodeSourceResponse) SetHeaders(v map[string]*string) *CreateCodeSourceResponse {
	s.Headers = v
	return s
}

func (s *CreateCodeSourceResponse) SetStatusCode(v int32) *CreateCodeSourceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCodeSourceResponse) SetBody(v *CreateCodeSourceResponseBody) *CreateCodeSourceResponse {
	s.Body = v
	return s
}

type CreateDatasetRequest struct {
	Accessibility  *string  `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	DataSourceType *string  `json:"DataSourceType,omitempty" xml:"DataSourceType,omitempty"`
	DataType       *string  `json:"DataType,omitempty" xml:"DataType,omitempty"`
	Description    *string  `json:"Description,omitempty" xml:"Description,omitempty"`
	Labels         []*Label `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	Name           *string  `json:"Name,omitempty" xml:"Name,omitempty"`
	Options        *string  `json:"Options,omitempty" xml:"Options,omitempty"`
	Property       *string  `json:"Property,omitempty" xml:"Property,omitempty"`
	Provider       *string  `json:"Provider,omitempty" xml:"Provider,omitempty"`
	ProviderType   *string  `json:"ProviderType,omitempty" xml:"ProviderType,omitempty"`
	SourceId       *string  `json:"SourceId,omitempty" xml:"SourceId,omitempty"`
	SourceType     *string  `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	Uri            *string  `json:"Uri,omitempty" xml:"Uri,omitempty"`
	WorkspaceId    *string  `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s CreateDatasetRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDatasetRequest) GoString() string {
	return s.String()
}

func (s *CreateDatasetRequest) SetAccessibility(v string) *CreateDatasetRequest {
	s.Accessibility = &v
	return s
}

func (s *CreateDatasetRequest) SetDataSourceType(v string) *CreateDatasetRequest {
	s.DataSourceType = &v
	return s
}

func (s *CreateDatasetRequest) SetDataType(v string) *CreateDatasetRequest {
	s.DataType = &v
	return s
}

func (s *CreateDatasetRequest) SetDescription(v string) *CreateDatasetRequest {
	s.Description = &v
	return s
}

func (s *CreateDatasetRequest) SetLabels(v []*Label) *CreateDatasetRequest {
	s.Labels = v
	return s
}

func (s *CreateDatasetRequest) SetName(v string) *CreateDatasetRequest {
	s.Name = &v
	return s
}

func (s *CreateDatasetRequest) SetOptions(v string) *CreateDatasetRequest {
	s.Options = &v
	return s
}

func (s *CreateDatasetRequest) SetProperty(v string) *CreateDatasetRequest {
	s.Property = &v
	return s
}

func (s *CreateDatasetRequest) SetProvider(v string) *CreateDatasetRequest {
	s.Provider = &v
	return s
}

func (s *CreateDatasetRequest) SetProviderType(v string) *CreateDatasetRequest {
	s.ProviderType = &v
	return s
}

func (s *CreateDatasetRequest) SetSourceId(v string) *CreateDatasetRequest {
	s.SourceId = &v
	return s
}

func (s *CreateDatasetRequest) SetSourceType(v string) *CreateDatasetRequest {
	s.SourceType = &v
	return s
}

func (s *CreateDatasetRequest) SetUri(v string) *CreateDatasetRequest {
	s.Uri = &v
	return s
}

func (s *CreateDatasetRequest) SetWorkspaceId(v string) *CreateDatasetRequest {
	s.WorkspaceId = &v
	return s
}

type CreateDatasetResponseBody struct {
	DatasetId *string `json:"DatasetId,omitempty" xml:"DatasetId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDatasetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDatasetResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDatasetResponseBody) SetDatasetId(v string) *CreateDatasetResponseBody {
	s.DatasetId = &v
	return s
}

func (s *CreateDatasetResponseBody) SetRequestId(v string) *CreateDatasetResponseBody {
	s.RequestId = &v
	return s
}

type CreateDatasetResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDatasetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDatasetResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDatasetResponse) GoString() string {
	return s.String()
}

func (s *CreateDatasetResponse) SetHeaders(v map[string]*string) *CreateDatasetResponse {
	s.Headers = v
	return s
}

func (s *CreateDatasetResponse) SetStatusCode(v int32) *CreateDatasetResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDatasetResponse) SetBody(v *CreateDatasetResponseBody) *CreateDatasetResponse {
	s.Body = v
	return s
}

type CreateDatasetLabelsRequest struct {
	Labels []*Label `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
}

func (s CreateDatasetLabelsRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDatasetLabelsRequest) GoString() string {
	return s.String()
}

func (s *CreateDatasetLabelsRequest) SetLabels(v []*Label) *CreateDatasetLabelsRequest {
	s.Labels = v
	return s
}

type CreateDatasetLabelsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDatasetLabelsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDatasetLabelsResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDatasetLabelsResponseBody) SetRequestId(v string) *CreateDatasetLabelsResponseBody {
	s.RequestId = &v
	return s
}

type CreateDatasetLabelsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDatasetLabelsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDatasetLabelsResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDatasetLabelsResponse) GoString() string {
	return s.String()
}

func (s *CreateDatasetLabelsResponse) SetHeaders(v map[string]*string) *CreateDatasetLabelsResponse {
	s.Headers = v
	return s
}

func (s *CreateDatasetLabelsResponse) SetStatusCode(v int32) *CreateDatasetLabelsResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDatasetLabelsResponse) SetBody(v *CreateDatasetLabelsResponseBody) *CreateDatasetLabelsResponse {
	s.Body = v
	return s
}

type CreateMemberRequest struct {
	Members []*CreateMemberRequestMembers `json:"Members,omitempty" xml:"Members,omitempty" type:"Repeated"`
}

func (s CreateMemberRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateMemberRequest) GoString() string {
	return s.String()
}

func (s *CreateMemberRequest) SetMembers(v []*CreateMemberRequestMembers) *CreateMemberRequest {
	s.Members = v
	return s
}

type CreateMemberRequestMembers struct {
	Roles  []*string `json:"Roles,omitempty" xml:"Roles,omitempty" type:"Repeated"`
	UserId *string   `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s CreateMemberRequestMembers) String() string {
	return tea.Prettify(s)
}

func (s CreateMemberRequestMembers) GoString() string {
	return s.String()
}

func (s *CreateMemberRequestMembers) SetRoles(v []*string) *CreateMemberRequestMembers {
	s.Roles = v
	return s
}

func (s *CreateMemberRequestMembers) SetUserId(v string) *CreateMemberRequestMembers {
	s.UserId = &v
	return s
}

type CreateMemberResponseBody struct {
	Members   []*CreateMemberResponseBodyMembers `json:"Members,omitempty" xml:"Members,omitempty" type:"Repeated"`
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateMemberResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateMemberResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMemberResponseBody) SetMembers(v []*CreateMemberResponseBodyMembers) *CreateMemberResponseBody {
	s.Members = v
	return s
}

func (s *CreateMemberResponseBody) SetRequestId(v string) *CreateMemberResponseBody {
	s.RequestId = &v
	return s
}

type CreateMemberResponseBodyMembers struct {
	DisplayName *string   `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	MemberId    *string   `json:"MemberId,omitempty" xml:"MemberId,omitempty"`
	Roles       []*string `json:"Roles,omitempty" xml:"Roles,omitempty" type:"Repeated"`
	UserId      *string   `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s CreateMemberResponseBodyMembers) String() string {
	return tea.Prettify(s)
}

func (s CreateMemberResponseBodyMembers) GoString() string {
	return s.String()
}

func (s *CreateMemberResponseBodyMembers) SetDisplayName(v string) *CreateMemberResponseBodyMembers {
	s.DisplayName = &v
	return s
}

func (s *CreateMemberResponseBodyMembers) SetMemberId(v string) *CreateMemberResponseBodyMembers {
	s.MemberId = &v
	return s
}

func (s *CreateMemberResponseBodyMembers) SetRoles(v []*string) *CreateMemberResponseBodyMembers {
	s.Roles = v
	return s
}

func (s *CreateMemberResponseBodyMembers) SetUserId(v string) *CreateMemberResponseBodyMembers {
	s.UserId = &v
	return s
}

type CreateMemberResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateMemberResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateMemberResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateMemberResponse) GoString() string {
	return s.String()
}

func (s *CreateMemberResponse) SetHeaders(v map[string]*string) *CreateMemberResponse {
	s.Headers = v
	return s
}

func (s *CreateMemberResponse) SetStatusCode(v int32) *CreateMemberResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateMemberResponse) SetBody(v *CreateMemberResponseBody) *CreateMemberResponse {
	s.Body = v
	return s
}

type CreateModelRequest struct {
	Accessibility    *string                `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	Domain           *string                `json:"Domain,omitempty" xml:"Domain,omitempty"`
	ExtraInfo        map[string]interface{} `json:"ExtraInfo,omitempty" xml:"ExtraInfo,omitempty"`
	Labels           []*Label               `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	ModelDescription *string                `json:"ModelDescription,omitempty" xml:"ModelDescription,omitempty"`
	ModelDoc         *string                `json:"ModelDoc,omitempty" xml:"ModelDoc,omitempty"`
	ModelName        *string                `json:"ModelName,omitempty" xml:"ModelName,omitempty"`
	ModelType        *string                `json:"ModelType,omitempty" xml:"ModelType,omitempty"`
	OrderNumber      *int64                 `json:"OrderNumber,omitempty" xml:"OrderNumber,omitempty"`
	Origin           *string                `json:"Origin,omitempty" xml:"Origin,omitempty"`
	Task             *string                `json:"Task,omitempty" xml:"Task,omitempty"`
	WorkspaceId      *string                `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s CreateModelRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateModelRequest) GoString() string {
	return s.String()
}

func (s *CreateModelRequest) SetAccessibility(v string) *CreateModelRequest {
	s.Accessibility = &v
	return s
}

func (s *CreateModelRequest) SetDomain(v string) *CreateModelRequest {
	s.Domain = &v
	return s
}

func (s *CreateModelRequest) SetExtraInfo(v map[string]interface{}) *CreateModelRequest {
	s.ExtraInfo = v
	return s
}

func (s *CreateModelRequest) SetLabels(v []*Label) *CreateModelRequest {
	s.Labels = v
	return s
}

func (s *CreateModelRequest) SetModelDescription(v string) *CreateModelRequest {
	s.ModelDescription = &v
	return s
}

func (s *CreateModelRequest) SetModelDoc(v string) *CreateModelRequest {
	s.ModelDoc = &v
	return s
}

func (s *CreateModelRequest) SetModelName(v string) *CreateModelRequest {
	s.ModelName = &v
	return s
}

func (s *CreateModelRequest) SetModelType(v string) *CreateModelRequest {
	s.ModelType = &v
	return s
}

func (s *CreateModelRequest) SetOrderNumber(v int64) *CreateModelRequest {
	s.OrderNumber = &v
	return s
}

func (s *CreateModelRequest) SetOrigin(v string) *CreateModelRequest {
	s.Origin = &v
	return s
}

func (s *CreateModelRequest) SetTask(v string) *CreateModelRequest {
	s.Task = &v
	return s
}

func (s *CreateModelRequest) SetWorkspaceId(v string) *CreateModelRequest {
	s.WorkspaceId = &v
	return s
}

type CreateModelResponseBody struct {
	ModelId   *string `json:"ModelId,omitempty" xml:"ModelId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateModelResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateModelResponseBody) GoString() string {
	return s.String()
}

func (s *CreateModelResponseBody) SetModelId(v string) *CreateModelResponseBody {
	s.ModelId = &v
	return s
}

func (s *CreateModelResponseBody) SetRequestId(v string) *CreateModelResponseBody {
	s.RequestId = &v
	return s
}

type CreateModelResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateModelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateModelResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateModelResponse) GoString() string {
	return s.String()
}

func (s *CreateModelResponse) SetHeaders(v map[string]*string) *CreateModelResponse {
	s.Headers = v
	return s
}

func (s *CreateModelResponse) SetStatusCode(v int32) *CreateModelResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateModelResponse) SetBody(v *CreateModelResponseBody) *CreateModelResponse {
	s.Body = v
	return s
}

type CreateModelLabelsRequest struct {
	Labels []*Label `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
}

func (s CreateModelLabelsRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateModelLabelsRequest) GoString() string {
	return s.String()
}

func (s *CreateModelLabelsRequest) SetLabels(v []*Label) *CreateModelLabelsRequest {
	s.Labels = v
	return s
}

type CreateModelLabelsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateModelLabelsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateModelLabelsResponseBody) GoString() string {
	return s.String()
}

func (s *CreateModelLabelsResponseBody) SetRequestId(v string) *CreateModelLabelsResponseBody {
	s.RequestId = &v
	return s
}

type CreateModelLabelsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateModelLabelsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateModelLabelsResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateModelLabelsResponse) GoString() string {
	return s.String()
}

func (s *CreateModelLabelsResponse) SetHeaders(v map[string]*string) *CreateModelLabelsResponse {
	s.Headers = v
	return s
}

func (s *CreateModelLabelsResponse) SetStatusCode(v int32) *CreateModelLabelsResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateModelLabelsResponse) SetBody(v *CreateModelLabelsResponseBody) *CreateModelLabelsResponse {
	s.Body = v
	return s
}

type CreateModelVersionRequest struct {
	ApprovalStatus     *string                `json:"ApprovalStatus,omitempty" xml:"ApprovalStatus,omitempty"`
	EvaluationSpec     map[string]interface{} `json:"EvaluationSpec,omitempty" xml:"EvaluationSpec,omitempty"`
	ExtraInfo          map[string]interface{} `json:"ExtraInfo,omitempty" xml:"ExtraInfo,omitempty"`
	FormatType         *string                `json:"FormatType,omitempty" xml:"FormatType,omitempty"`
	FrameworkType      *string                `json:"FrameworkType,omitempty" xml:"FrameworkType,omitempty"`
	InferenceSpec      map[string]interface{} `json:"InferenceSpec,omitempty" xml:"InferenceSpec,omitempty"`
	Labels             []*Label               `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	Metrics            map[string]interface{} `json:"Metrics,omitempty" xml:"Metrics,omitempty"`
	Options            *string                `json:"Options,omitempty" xml:"Options,omitempty"`
	SourceId           *string                `json:"SourceId,omitempty" xml:"SourceId,omitempty"`
	SourceType         *string                `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	TrainingSpec       map[string]interface{} `json:"TrainingSpec,omitempty" xml:"TrainingSpec,omitempty"`
	Uri                *string                `json:"Uri,omitempty" xml:"Uri,omitempty"`
	VersionDescription *string                `json:"VersionDescription,omitempty" xml:"VersionDescription,omitempty"`
	VersionName        *string                `json:"VersionName,omitempty" xml:"VersionName,omitempty"`
}

func (s CreateModelVersionRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateModelVersionRequest) GoString() string {
	return s.String()
}

func (s *CreateModelVersionRequest) SetApprovalStatus(v string) *CreateModelVersionRequest {
	s.ApprovalStatus = &v
	return s
}

func (s *CreateModelVersionRequest) SetEvaluationSpec(v map[string]interface{}) *CreateModelVersionRequest {
	s.EvaluationSpec = v
	return s
}

func (s *CreateModelVersionRequest) SetExtraInfo(v map[string]interface{}) *CreateModelVersionRequest {
	s.ExtraInfo = v
	return s
}

func (s *CreateModelVersionRequest) SetFormatType(v string) *CreateModelVersionRequest {
	s.FormatType = &v
	return s
}

func (s *CreateModelVersionRequest) SetFrameworkType(v string) *CreateModelVersionRequest {
	s.FrameworkType = &v
	return s
}

func (s *CreateModelVersionRequest) SetInferenceSpec(v map[string]interface{}) *CreateModelVersionRequest {
	s.InferenceSpec = v
	return s
}

func (s *CreateModelVersionRequest) SetLabels(v []*Label) *CreateModelVersionRequest {
	s.Labels = v
	return s
}

func (s *CreateModelVersionRequest) SetMetrics(v map[string]interface{}) *CreateModelVersionRequest {
	s.Metrics = v
	return s
}

func (s *CreateModelVersionRequest) SetOptions(v string) *CreateModelVersionRequest {
	s.Options = &v
	return s
}

func (s *CreateModelVersionRequest) SetSourceId(v string) *CreateModelVersionRequest {
	s.SourceId = &v
	return s
}

func (s *CreateModelVersionRequest) SetSourceType(v string) *CreateModelVersionRequest {
	s.SourceType = &v
	return s
}

func (s *CreateModelVersionRequest) SetTrainingSpec(v map[string]interface{}) *CreateModelVersionRequest {
	s.TrainingSpec = v
	return s
}

func (s *CreateModelVersionRequest) SetUri(v string) *CreateModelVersionRequest {
	s.Uri = &v
	return s
}

func (s *CreateModelVersionRequest) SetVersionDescription(v string) *CreateModelVersionRequest {
	s.VersionDescription = &v
	return s
}

func (s *CreateModelVersionRequest) SetVersionName(v string) *CreateModelVersionRequest {
	s.VersionName = &v
	return s
}

type CreateModelVersionResponseBody struct {
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	VersionName *string `json:"VersionName,omitempty" xml:"VersionName,omitempty"`
}

func (s CreateModelVersionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateModelVersionResponseBody) GoString() string {
	return s.String()
}

func (s *CreateModelVersionResponseBody) SetRequestId(v string) *CreateModelVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateModelVersionResponseBody) SetVersionName(v string) *CreateModelVersionResponseBody {
	s.VersionName = &v
	return s
}

type CreateModelVersionResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateModelVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateModelVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateModelVersionResponse) GoString() string {
	return s.String()
}

func (s *CreateModelVersionResponse) SetHeaders(v map[string]*string) *CreateModelVersionResponse {
	s.Headers = v
	return s
}

func (s *CreateModelVersionResponse) SetStatusCode(v int32) *CreateModelVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateModelVersionResponse) SetBody(v *CreateModelVersionResponseBody) *CreateModelVersionResponse {
	s.Body = v
	return s
}

type CreateModelVersionLabelsRequest struct {
	Labels []*Label `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
}

func (s CreateModelVersionLabelsRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateModelVersionLabelsRequest) GoString() string {
	return s.String()
}

func (s *CreateModelVersionLabelsRequest) SetLabels(v []*Label) *CreateModelVersionLabelsRequest {
	s.Labels = v
	return s
}

type CreateModelVersionLabelsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateModelVersionLabelsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateModelVersionLabelsResponseBody) GoString() string {
	return s.String()
}

func (s *CreateModelVersionLabelsResponseBody) SetRequestId(v string) *CreateModelVersionLabelsResponseBody {
	s.RequestId = &v
	return s
}

type CreateModelVersionLabelsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateModelVersionLabelsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateModelVersionLabelsResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateModelVersionLabelsResponse) GoString() string {
	return s.String()
}

func (s *CreateModelVersionLabelsResponse) SetHeaders(v map[string]*string) *CreateModelVersionLabelsResponse {
	s.Headers = v
	return s
}

func (s *CreateModelVersionLabelsResponse) SetStatusCode(v int32) *CreateModelVersionLabelsResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateModelVersionLabelsResponse) SetBody(v *CreateModelVersionLabelsResponseBody) *CreateModelVersionLabelsResponse {
	s.Body = v
	return s
}

type CreateProductOrdersRequest struct {
	AutoPay  *bool                                 `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	Products []*CreateProductOrdersRequestProducts `json:"Products,omitempty" xml:"Products,omitempty" type:"Repeated"`
}

func (s CreateProductOrdersRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateProductOrdersRequest) GoString() string {
	return s.String()
}

func (s *CreateProductOrdersRequest) SetAutoPay(v bool) *CreateProductOrdersRequest {
	s.AutoPay = &v
	return s
}

func (s *CreateProductOrdersRequest) SetProducts(v []*CreateProductOrdersRequestProducts) *CreateProductOrdersRequest {
	s.Products = v
	return s
}

type CreateProductOrdersRequestProducts struct {
	AutoRenew          *bool                                                   `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	ChargeType         *string                                                 `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	Duration           *int64                                                  `json:"Duration,omitempty" xml:"Duration,omitempty"`
	InstanceProperties []*CreateProductOrdersRequestProductsInstanceProperties `json:"InstanceProperties,omitempty" xml:"InstanceProperties,omitempty" type:"Repeated"`
	OrderType          *string                                                 `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	PricingCycle       *string                                                 `json:"PricingCycle,omitempty" xml:"PricingCycle,omitempty"`
	ProductCode        *string                                                 `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
}

func (s CreateProductOrdersRequestProducts) String() string {
	return tea.Prettify(s)
}

func (s CreateProductOrdersRequestProducts) GoString() string {
	return s.String()
}

func (s *CreateProductOrdersRequestProducts) SetAutoRenew(v bool) *CreateProductOrdersRequestProducts {
	s.AutoRenew = &v
	return s
}

func (s *CreateProductOrdersRequestProducts) SetChargeType(v string) *CreateProductOrdersRequestProducts {
	s.ChargeType = &v
	return s
}

func (s *CreateProductOrdersRequestProducts) SetDuration(v int64) *CreateProductOrdersRequestProducts {
	s.Duration = &v
	return s
}

func (s *CreateProductOrdersRequestProducts) SetInstanceProperties(v []*CreateProductOrdersRequestProductsInstanceProperties) *CreateProductOrdersRequestProducts {
	s.InstanceProperties = v
	return s
}

func (s *CreateProductOrdersRequestProducts) SetOrderType(v string) *CreateProductOrdersRequestProducts {
	s.OrderType = &v
	return s
}

func (s *CreateProductOrdersRequestProducts) SetPricingCycle(v string) *CreateProductOrdersRequestProducts {
	s.PricingCycle = &v
	return s
}

func (s *CreateProductOrdersRequestProducts) SetProductCode(v string) *CreateProductOrdersRequestProducts {
	s.ProductCode = &v
	return s
}

type CreateProductOrdersRequestProductsInstanceProperties struct {
	Code  *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Name  *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateProductOrdersRequestProductsInstanceProperties) String() string {
	return tea.Prettify(s)
}

func (s CreateProductOrdersRequestProductsInstanceProperties) GoString() string {
	return s.String()
}

func (s *CreateProductOrdersRequestProductsInstanceProperties) SetCode(v string) *CreateProductOrdersRequestProductsInstanceProperties {
	s.Code = &v
	return s
}

func (s *CreateProductOrdersRequestProductsInstanceProperties) SetName(v string) *CreateProductOrdersRequestProductsInstanceProperties {
	s.Name = &v
	return s
}

func (s *CreateProductOrdersRequestProductsInstanceProperties) SetValue(v string) *CreateProductOrdersRequestProductsInstanceProperties {
	s.Value = &v
	return s
}

type CreateProductOrdersResponseBody struct {
	BuyProductRequestId *string `json:"BuyProductRequestId,omitempty" xml:"BuyProductRequestId,omitempty"`
	Message             *string `json:"Message,omitempty" xml:"Message,omitempty"`
	OrderId             *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	RequestId           *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateProductOrdersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateProductOrdersResponseBody) GoString() string {
	return s.String()
}

func (s *CreateProductOrdersResponseBody) SetBuyProductRequestId(v string) *CreateProductOrdersResponseBody {
	s.BuyProductRequestId = &v
	return s
}

func (s *CreateProductOrdersResponseBody) SetMessage(v string) *CreateProductOrdersResponseBody {
	s.Message = &v
	return s
}

func (s *CreateProductOrdersResponseBody) SetOrderId(v string) *CreateProductOrdersResponseBody {
	s.OrderId = &v
	return s
}

func (s *CreateProductOrdersResponseBody) SetRequestId(v string) *CreateProductOrdersResponseBody {
	s.RequestId = &v
	return s
}

type CreateProductOrdersResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateProductOrdersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateProductOrdersResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateProductOrdersResponse) GoString() string {
	return s.String()
}

func (s *CreateProductOrdersResponse) SetHeaders(v map[string]*string) *CreateProductOrdersResponse {
	s.Headers = v
	return s
}

func (s *CreateProductOrdersResponse) SetStatusCode(v int32) *CreateProductOrdersResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateProductOrdersResponse) SetBody(v *CreateProductOrdersResponseBody) *CreateProductOrdersResponse {
	s.Body = v
	return s
}

type CreateWorkspaceRequest struct {
	Description   *string   `json:"Description,omitempty" xml:"Description,omitempty"`
	DisplayName   *string   `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	EnvTypes      []*string `json:"EnvTypes,omitempty" xml:"EnvTypes,omitempty" type:"Repeated"`
	WorkspaceName *string   `json:"WorkspaceName,omitempty" xml:"WorkspaceName,omitempty"`
}

func (s CreateWorkspaceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateWorkspaceRequest) GoString() string {
	return s.String()
}

func (s *CreateWorkspaceRequest) SetDescription(v string) *CreateWorkspaceRequest {
	s.Description = &v
	return s
}

func (s *CreateWorkspaceRequest) SetDisplayName(v string) *CreateWorkspaceRequest {
	s.DisplayName = &v
	return s
}

func (s *CreateWorkspaceRequest) SetEnvTypes(v []*string) *CreateWorkspaceRequest {
	s.EnvTypes = v
	return s
}

func (s *CreateWorkspaceRequest) SetWorkspaceName(v string) *CreateWorkspaceRequest {
	s.WorkspaceName = &v
	return s
}

type CreateWorkspaceResponseBody struct {
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s CreateWorkspaceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateWorkspaceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateWorkspaceResponseBody) SetRequestId(v string) *CreateWorkspaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateWorkspaceResponseBody) SetWorkspaceId(v string) *CreateWorkspaceResponseBody {
	s.WorkspaceId = &v
	return s
}

type CreateWorkspaceResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateWorkspaceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateWorkspaceResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateWorkspaceResponse) GoString() string {
	return s.String()
}

func (s *CreateWorkspaceResponse) SetHeaders(v map[string]*string) *CreateWorkspaceResponse {
	s.Headers = v
	return s
}

func (s *CreateWorkspaceResponse) SetStatusCode(v int32) *CreateWorkspaceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateWorkspaceResponse) SetBody(v *CreateWorkspaceResponseBody) *CreateWorkspaceResponse {
	s.Body = v
	return s
}

type CreateWorkspaceResourceRequest struct {
	Option    *string                                    `json:"Option,omitempty" xml:"Option,omitempty"`
	Resources []*CreateWorkspaceResourceRequestResources `json:"Resources,omitempty" xml:"Resources,omitempty" type:"Repeated"`
}

func (s CreateWorkspaceResourceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateWorkspaceResourceRequest) GoString() string {
	return s.String()
}

func (s *CreateWorkspaceResourceRequest) SetOption(v string) *CreateWorkspaceResourceRequest {
	s.Option = &v
	return s
}

func (s *CreateWorkspaceResourceRequest) SetResources(v []*CreateWorkspaceResourceRequestResources) *CreateWorkspaceResourceRequest {
	s.Resources = v
	return s
}

type CreateWorkspaceResourceRequestResources struct {
	EnvType      *string                                          `json:"EnvType,omitempty" xml:"EnvType,omitempty"`
	GroupName    *string                                          `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	IsDefault    *bool                                            `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	Labels       []*CreateWorkspaceResourceRequestResourcesLabels `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	Name         *string                                          `json:"Name,omitempty" xml:"Name,omitempty"`
	ProductType  *string                                          `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	Quotas       []*CreateWorkspaceResourceRequestResourcesQuotas `json:"Quotas,omitempty" xml:"Quotas,omitempty" type:"Repeated"`
	ResourceType *string                                          `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Spec         map[string]interface{}                           `json:"Spec,omitempty" xml:"Spec,omitempty"`
	WorkspaceId  *string                                          `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s CreateWorkspaceResourceRequestResources) String() string {
	return tea.Prettify(s)
}

func (s CreateWorkspaceResourceRequestResources) GoString() string {
	return s.String()
}

func (s *CreateWorkspaceResourceRequestResources) SetEnvType(v string) *CreateWorkspaceResourceRequestResources {
	s.EnvType = &v
	return s
}

func (s *CreateWorkspaceResourceRequestResources) SetGroupName(v string) *CreateWorkspaceResourceRequestResources {
	s.GroupName = &v
	return s
}

func (s *CreateWorkspaceResourceRequestResources) SetIsDefault(v bool) *CreateWorkspaceResourceRequestResources {
	s.IsDefault = &v
	return s
}

func (s *CreateWorkspaceResourceRequestResources) SetLabels(v []*CreateWorkspaceResourceRequestResourcesLabels) *CreateWorkspaceResourceRequestResources {
	s.Labels = v
	return s
}

func (s *CreateWorkspaceResourceRequestResources) SetName(v string) *CreateWorkspaceResourceRequestResources {
	s.Name = &v
	return s
}

func (s *CreateWorkspaceResourceRequestResources) SetProductType(v string) *CreateWorkspaceResourceRequestResources {
	s.ProductType = &v
	return s
}

func (s *CreateWorkspaceResourceRequestResources) SetQuotas(v []*CreateWorkspaceResourceRequestResourcesQuotas) *CreateWorkspaceResourceRequestResources {
	s.Quotas = v
	return s
}

func (s *CreateWorkspaceResourceRequestResources) SetResourceType(v string) *CreateWorkspaceResourceRequestResources {
	s.ResourceType = &v
	return s
}

func (s *CreateWorkspaceResourceRequestResources) SetSpec(v map[string]interface{}) *CreateWorkspaceResourceRequestResources {
	s.Spec = v
	return s
}

func (s *CreateWorkspaceResourceRequestResources) SetWorkspaceId(v string) *CreateWorkspaceResourceRequestResources {
	s.WorkspaceId = &v
	return s
}

type CreateWorkspaceResourceRequestResourcesLabels struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateWorkspaceResourceRequestResourcesLabels) String() string {
	return tea.Prettify(s)
}

func (s CreateWorkspaceResourceRequestResourcesLabels) GoString() string {
	return s.String()
}

func (s *CreateWorkspaceResourceRequestResourcesLabels) SetKey(v string) *CreateWorkspaceResourceRequestResourcesLabels {
	s.Key = &v
	return s
}

func (s *CreateWorkspaceResourceRequestResourcesLabels) SetValue(v string) *CreateWorkspaceResourceRequestResourcesLabels {
	s.Value = &v
	return s
}

type CreateWorkspaceResourceRequestResourcesQuotas struct {
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s CreateWorkspaceResourceRequestResourcesQuotas) String() string {
	return tea.Prettify(s)
}

func (s CreateWorkspaceResourceRequestResourcesQuotas) GoString() string {
	return s.String()
}

func (s *CreateWorkspaceResourceRequestResourcesQuotas) SetId(v string) *CreateWorkspaceResourceRequestResourcesQuotas {
	s.Id = &v
	return s
}

type CreateWorkspaceResourceResponseBody struct {
	RequestId  *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Resources  []*CreateWorkspaceResourceResponseBodyResources `json:"Resources,omitempty" xml:"Resources,omitempty" type:"Repeated"`
	TotalCount *int64                                          `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s CreateWorkspaceResourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateWorkspaceResourceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateWorkspaceResourceResponseBody) SetRequestId(v string) *CreateWorkspaceResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateWorkspaceResourceResponseBody) SetResources(v []*CreateWorkspaceResourceResponseBodyResources) *CreateWorkspaceResourceResponseBody {
	s.Resources = v
	return s
}

func (s *CreateWorkspaceResourceResponseBody) SetTotalCount(v int64) *CreateWorkspaceResourceResponseBody {
	s.TotalCount = &v
	return s
}

type CreateWorkspaceResourceResponseBodyResources struct {
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s CreateWorkspaceResourceResponseBodyResources) String() string {
	return tea.Prettify(s)
}

func (s CreateWorkspaceResourceResponseBodyResources) GoString() string {
	return s.String()
}

func (s *CreateWorkspaceResourceResponseBodyResources) SetId(v string) *CreateWorkspaceResourceResponseBodyResources {
	s.Id = &v
	return s
}

type CreateWorkspaceResourceResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateWorkspaceResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateWorkspaceResourceResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateWorkspaceResourceResponse) GoString() string {
	return s.String()
}

func (s *CreateWorkspaceResourceResponse) SetHeaders(v map[string]*string) *CreateWorkspaceResourceResponse {
	s.Headers = v
	return s
}

func (s *CreateWorkspaceResourceResponse) SetStatusCode(v int32) *CreateWorkspaceResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateWorkspaceResourceResponse) SetBody(v *CreateWorkspaceResourceResponseBody) *CreateWorkspaceResourceResponse {
	s.Body = v
	return s
}

type DeleteCodeSourceResponseBody struct {
	CodeSourceId *string `json:"CodeSourceId,omitempty" xml:"CodeSourceId,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteCodeSourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteCodeSourceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCodeSourceResponseBody) SetCodeSourceId(v string) *DeleteCodeSourceResponseBody {
	s.CodeSourceId = &v
	return s
}

func (s *DeleteCodeSourceResponseBody) SetRequestId(v string) *DeleteCodeSourceResponseBody {
	s.RequestId = &v
	return s
}

type DeleteCodeSourceResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteCodeSourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteCodeSourceResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteCodeSourceResponse) GoString() string {
	return s.String()
}

func (s *DeleteCodeSourceResponse) SetHeaders(v map[string]*string) *DeleteCodeSourceResponse {
	s.Headers = v
	return s
}

func (s *DeleteCodeSourceResponse) SetStatusCode(v int32) *DeleteCodeSourceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCodeSourceResponse) SetBody(v *DeleteCodeSourceResponseBody) *DeleteCodeSourceResponse {
	s.Body = v
	return s
}

type DeleteDatasetResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDatasetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDatasetResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDatasetResponseBody) SetRequestId(v string) *DeleteDatasetResponseBody {
	s.RequestId = &v
	return s
}

type DeleteDatasetResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDatasetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDatasetResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDatasetResponse) GoString() string {
	return s.String()
}

func (s *DeleteDatasetResponse) SetHeaders(v map[string]*string) *DeleteDatasetResponse {
	s.Headers = v
	return s
}

func (s *DeleteDatasetResponse) SetStatusCode(v int32) *DeleteDatasetResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDatasetResponse) SetBody(v *DeleteDatasetResponseBody) *DeleteDatasetResponse {
	s.Body = v
	return s
}

type DeleteDatasetLabelsRequest struct {
	LabelKeys *string `json:"LabelKeys,omitempty" xml:"LabelKeys,omitempty"`
}

func (s DeleteDatasetLabelsRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDatasetLabelsRequest) GoString() string {
	return s.String()
}

func (s *DeleteDatasetLabelsRequest) SetLabelKeys(v string) *DeleteDatasetLabelsRequest {
	s.LabelKeys = &v
	return s
}

type DeleteDatasetLabelsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDatasetLabelsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDatasetLabelsResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDatasetLabelsResponseBody) SetRequestId(v string) *DeleteDatasetLabelsResponseBody {
	s.RequestId = &v
	return s
}

type DeleteDatasetLabelsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDatasetLabelsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDatasetLabelsResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDatasetLabelsResponse) GoString() string {
	return s.String()
}

func (s *DeleteDatasetLabelsResponse) SetHeaders(v map[string]*string) *DeleteDatasetLabelsResponse {
	s.Headers = v
	return s
}

func (s *DeleteDatasetLabelsResponse) SetStatusCode(v int32) *DeleteDatasetLabelsResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDatasetLabelsResponse) SetBody(v *DeleteDatasetLabelsResponseBody) *DeleteDatasetLabelsResponse {
	s.Body = v
	return s
}

type DeleteMembersRequest struct {
	MemberIds *string `json:"MemberIds,omitempty" xml:"MemberIds,omitempty"`
}

func (s DeleteMembersRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteMembersRequest) GoString() string {
	return s.String()
}

func (s *DeleteMembersRequest) SetMemberIds(v string) *DeleteMembersRequest {
	s.MemberIds = &v
	return s
}

type DeleteMembersResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteMembersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteMembersResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteMembersResponseBody) SetCode(v string) *DeleteMembersResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteMembersResponseBody) SetMessage(v string) *DeleteMembersResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteMembersResponseBody) SetRequestId(v string) *DeleteMembersResponseBody {
	s.RequestId = &v
	return s
}

type DeleteMembersResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteMembersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteMembersResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteMembersResponse) GoString() string {
	return s.String()
}

func (s *DeleteMembersResponse) SetHeaders(v map[string]*string) *DeleteMembersResponse {
	s.Headers = v
	return s
}

func (s *DeleteMembersResponse) SetStatusCode(v int32) *DeleteMembersResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteMembersResponse) SetBody(v *DeleteMembersResponseBody) *DeleteMembersResponse {
	s.Body = v
	return s
}

type DeleteModelResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteModelResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteModelResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteModelResponseBody) SetRequestId(v string) *DeleteModelResponseBody {
	s.RequestId = &v
	return s
}

type DeleteModelResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteModelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteModelResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteModelResponse) GoString() string {
	return s.String()
}

func (s *DeleteModelResponse) SetHeaders(v map[string]*string) *DeleteModelResponse {
	s.Headers = v
	return s
}

func (s *DeleteModelResponse) SetStatusCode(v int32) *DeleteModelResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteModelResponse) SetBody(v *DeleteModelResponseBody) *DeleteModelResponse {
	s.Body = v
	return s
}

type DeleteModelLabelsRequest struct {
	LabelKeys *string `json:"LabelKeys,omitempty" xml:"LabelKeys,omitempty"`
}

func (s DeleteModelLabelsRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteModelLabelsRequest) GoString() string {
	return s.String()
}

func (s *DeleteModelLabelsRequest) SetLabelKeys(v string) *DeleteModelLabelsRequest {
	s.LabelKeys = &v
	return s
}

type DeleteModelLabelsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteModelLabelsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteModelLabelsResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteModelLabelsResponseBody) SetRequestId(v string) *DeleteModelLabelsResponseBody {
	s.RequestId = &v
	return s
}

type DeleteModelLabelsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteModelLabelsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteModelLabelsResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteModelLabelsResponse) GoString() string {
	return s.String()
}

func (s *DeleteModelLabelsResponse) SetHeaders(v map[string]*string) *DeleteModelLabelsResponse {
	s.Headers = v
	return s
}

func (s *DeleteModelLabelsResponse) SetStatusCode(v int32) *DeleteModelLabelsResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteModelLabelsResponse) SetBody(v *DeleteModelLabelsResponseBody) *DeleteModelLabelsResponse {
	s.Body = v
	return s
}

type DeleteModelVersionResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteModelVersionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteModelVersionResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteModelVersionResponseBody) SetRequestId(v string) *DeleteModelVersionResponseBody {
	s.RequestId = &v
	return s
}

type DeleteModelVersionResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteModelVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteModelVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteModelVersionResponse) GoString() string {
	return s.String()
}

func (s *DeleteModelVersionResponse) SetHeaders(v map[string]*string) *DeleteModelVersionResponse {
	s.Headers = v
	return s
}

func (s *DeleteModelVersionResponse) SetStatusCode(v int32) *DeleteModelVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteModelVersionResponse) SetBody(v *DeleteModelVersionResponseBody) *DeleteModelVersionResponse {
	s.Body = v
	return s
}

type DeleteModelVersionLabelsRequest struct {
	LabelKeys *string `json:"LabelKeys,omitempty" xml:"LabelKeys,omitempty"`
}

func (s DeleteModelVersionLabelsRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteModelVersionLabelsRequest) GoString() string {
	return s.String()
}

func (s *DeleteModelVersionLabelsRequest) SetLabelKeys(v string) *DeleteModelVersionLabelsRequest {
	s.LabelKeys = &v
	return s
}

type DeleteModelVersionLabelsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteModelVersionLabelsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteModelVersionLabelsResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteModelVersionLabelsResponseBody) SetRequestId(v string) *DeleteModelVersionLabelsResponseBody {
	s.RequestId = &v
	return s
}

type DeleteModelVersionLabelsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteModelVersionLabelsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteModelVersionLabelsResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteModelVersionLabelsResponse) GoString() string {
	return s.String()
}

func (s *DeleteModelVersionLabelsResponse) SetHeaders(v map[string]*string) *DeleteModelVersionLabelsResponse {
	s.Headers = v
	return s
}

func (s *DeleteModelVersionLabelsResponse) SetStatusCode(v int32) *DeleteModelVersionLabelsResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteModelVersionLabelsResponse) SetBody(v *DeleteModelVersionLabelsResponseBody) *DeleteModelVersionLabelsResponse {
	s.Body = v
	return s
}

type DeleteWorkspaceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteWorkspaceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteWorkspaceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteWorkspaceResponseBody) SetRequestId(v string) *DeleteWorkspaceResponseBody {
	s.RequestId = &v
	return s
}

type DeleteWorkspaceResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteWorkspaceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteWorkspaceResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteWorkspaceResponse) GoString() string {
	return s.String()
}

func (s *DeleteWorkspaceResponse) SetHeaders(v map[string]*string) *DeleteWorkspaceResponse {
	s.Headers = v
	return s
}

func (s *DeleteWorkspaceResponse) SetStatusCode(v int32) *DeleteWorkspaceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteWorkspaceResponse) SetBody(v *DeleteWorkspaceResponseBody) *DeleteWorkspaceResponse {
	s.Body = v
	return s
}

type DeleteWorkspaceResourceRequest struct {
	GroupName    *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	Labels       *string `json:"Labels,omitempty" xml:"Labels,omitempty"`
	Option       *string `json:"Option,omitempty" xml:"Option,omitempty"`
	ProductType  *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	ResourceIds  *string `json:"ResourceIds,omitempty" xml:"ResourceIds,omitempty"`
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s DeleteWorkspaceResourceRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteWorkspaceResourceRequest) GoString() string {
	return s.String()
}

func (s *DeleteWorkspaceResourceRequest) SetGroupName(v string) *DeleteWorkspaceResourceRequest {
	s.GroupName = &v
	return s
}

func (s *DeleteWorkspaceResourceRequest) SetLabels(v string) *DeleteWorkspaceResourceRequest {
	s.Labels = &v
	return s
}

func (s *DeleteWorkspaceResourceRequest) SetOption(v string) *DeleteWorkspaceResourceRequest {
	s.Option = &v
	return s
}

func (s *DeleteWorkspaceResourceRequest) SetProductType(v string) *DeleteWorkspaceResourceRequest {
	s.ProductType = &v
	return s
}

func (s *DeleteWorkspaceResourceRequest) SetResourceIds(v string) *DeleteWorkspaceResourceRequest {
	s.ResourceIds = &v
	return s
}

func (s *DeleteWorkspaceResourceRequest) SetResourceType(v string) *DeleteWorkspaceResourceRequest {
	s.ResourceType = &v
	return s
}

type DeleteWorkspaceResourceResponseBody struct {
	RequestId   *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResourceIds []*string `json:"ResourceIds,omitempty" xml:"ResourceIds,omitempty" type:"Repeated"`
}

func (s DeleteWorkspaceResourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteWorkspaceResourceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteWorkspaceResourceResponseBody) SetRequestId(v string) *DeleteWorkspaceResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteWorkspaceResourceResponseBody) SetResourceIds(v []*string) *DeleteWorkspaceResourceResponseBody {
	s.ResourceIds = v
	return s
}

type DeleteWorkspaceResourceResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteWorkspaceResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteWorkspaceResourceResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteWorkspaceResourceResponse) GoString() string {
	return s.String()
}

func (s *DeleteWorkspaceResourceResponse) SetHeaders(v map[string]*string) *DeleteWorkspaceResourceResponse {
	s.Headers = v
	return s
}

func (s *DeleteWorkspaceResourceResponse) SetStatusCode(v int32) *DeleteWorkspaceResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteWorkspaceResourceResponse) SetBody(v *DeleteWorkspaceResourceResponseBody) *DeleteWorkspaceResourceResponse {
	s.Body = v
	return s
}

type GetCodeSourceResponseBody struct {
	Accessibility       *string `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	CodeBranch          *string `json:"CodeBranch,omitempty" xml:"CodeBranch,omitempty"`
	CodeCommit          *string `json:"CodeCommit,omitempty" xml:"CodeCommit,omitempty"`
	CodeRepo            *string `json:"CodeRepo,omitempty" xml:"CodeRepo,omitempty"`
	CodeRepoAccessToken *string `json:"CodeRepoAccessToken,omitempty" xml:"CodeRepoAccessToken,omitempty"`
	CodeRepoUserName    *string `json:"CodeRepoUserName,omitempty" xml:"CodeRepoUserName,omitempty"`
	CodeSourceId        *string `json:"CodeSourceId,omitempty" xml:"CodeSourceId,omitempty"`
	Description         *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DisplayName         *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	GmtCreateTime       *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	GmtModifyTime       *string `json:"GmtModifyTime,omitempty" xml:"GmtModifyTime,omitempty"`
	MountPath           *string `json:"MountPath,omitempty" xml:"MountPath,omitempty"`
	RequestId           *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	UserId              *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	WorkspaceId         *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s GetCodeSourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetCodeSourceResponseBody) GoString() string {
	return s.String()
}

func (s *GetCodeSourceResponseBody) SetAccessibility(v string) *GetCodeSourceResponseBody {
	s.Accessibility = &v
	return s
}

func (s *GetCodeSourceResponseBody) SetCodeBranch(v string) *GetCodeSourceResponseBody {
	s.CodeBranch = &v
	return s
}

func (s *GetCodeSourceResponseBody) SetCodeCommit(v string) *GetCodeSourceResponseBody {
	s.CodeCommit = &v
	return s
}

func (s *GetCodeSourceResponseBody) SetCodeRepo(v string) *GetCodeSourceResponseBody {
	s.CodeRepo = &v
	return s
}

func (s *GetCodeSourceResponseBody) SetCodeRepoAccessToken(v string) *GetCodeSourceResponseBody {
	s.CodeRepoAccessToken = &v
	return s
}

func (s *GetCodeSourceResponseBody) SetCodeRepoUserName(v string) *GetCodeSourceResponseBody {
	s.CodeRepoUserName = &v
	return s
}

func (s *GetCodeSourceResponseBody) SetCodeSourceId(v string) *GetCodeSourceResponseBody {
	s.CodeSourceId = &v
	return s
}

func (s *GetCodeSourceResponseBody) SetDescription(v string) *GetCodeSourceResponseBody {
	s.Description = &v
	return s
}

func (s *GetCodeSourceResponseBody) SetDisplayName(v string) *GetCodeSourceResponseBody {
	s.DisplayName = &v
	return s
}

func (s *GetCodeSourceResponseBody) SetGmtCreateTime(v string) *GetCodeSourceResponseBody {
	s.GmtCreateTime = &v
	return s
}

func (s *GetCodeSourceResponseBody) SetGmtModifyTime(v string) *GetCodeSourceResponseBody {
	s.GmtModifyTime = &v
	return s
}

func (s *GetCodeSourceResponseBody) SetMountPath(v string) *GetCodeSourceResponseBody {
	s.MountPath = &v
	return s
}

func (s *GetCodeSourceResponseBody) SetRequestId(v string) *GetCodeSourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCodeSourceResponseBody) SetUserId(v string) *GetCodeSourceResponseBody {
	s.UserId = &v
	return s
}

func (s *GetCodeSourceResponseBody) SetWorkspaceId(v string) *GetCodeSourceResponseBody {
	s.WorkspaceId = &v
	return s
}

type GetCodeSourceResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCodeSourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCodeSourceResponse) String() string {
	return tea.Prettify(s)
}

func (s GetCodeSourceResponse) GoString() string {
	return s.String()
}

func (s *GetCodeSourceResponse) SetHeaders(v map[string]*string) *GetCodeSourceResponse {
	s.Headers = v
	return s
}

func (s *GetCodeSourceResponse) SetStatusCode(v int32) *GetCodeSourceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCodeSourceResponse) SetBody(v *GetCodeSourceResponseBody) *GetCodeSourceResponse {
	s.Body = v
	return s
}

type GetDatasetResponseBody struct {
	Accessibility   *string  `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	DataSourceType  *string  `json:"DataSourceType,omitempty" xml:"DataSourceType,omitempty"`
	DataType        *string  `json:"DataType,omitempty" xml:"DataType,omitempty"`
	DatasetId       *string  `json:"DatasetId,omitempty" xml:"DatasetId,omitempty"`
	Description     *string  `json:"Description,omitempty" xml:"Description,omitempty"`
	GmtCreateTime   *string  `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	GmtModifiedTime *string  `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	Labels          []*Label `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	Name            *string  `json:"Name,omitempty" xml:"Name,omitempty"`
	Options         *string  `json:"Options,omitempty" xml:"Options,omitempty"`
	OwnerId         *string  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	Property        *string  `json:"Property,omitempty" xml:"Property,omitempty"`
	ProviderType    *string  `json:"ProviderType,omitempty" xml:"ProviderType,omitempty"`
	RequestId       *string  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SourceId        *string  `json:"SourceId,omitempty" xml:"SourceId,omitempty"`
	SourceType      *string  `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	Uri             *string  `json:"Uri,omitempty" xml:"Uri,omitempty"`
	UserId          *string  `json:"UserId,omitempty" xml:"UserId,omitempty"`
	WorkspaceId     *string  `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s GetDatasetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDatasetResponseBody) GoString() string {
	return s.String()
}

func (s *GetDatasetResponseBody) SetAccessibility(v string) *GetDatasetResponseBody {
	s.Accessibility = &v
	return s
}

func (s *GetDatasetResponseBody) SetDataSourceType(v string) *GetDatasetResponseBody {
	s.DataSourceType = &v
	return s
}

func (s *GetDatasetResponseBody) SetDataType(v string) *GetDatasetResponseBody {
	s.DataType = &v
	return s
}

func (s *GetDatasetResponseBody) SetDatasetId(v string) *GetDatasetResponseBody {
	s.DatasetId = &v
	return s
}

func (s *GetDatasetResponseBody) SetDescription(v string) *GetDatasetResponseBody {
	s.Description = &v
	return s
}

func (s *GetDatasetResponseBody) SetGmtCreateTime(v string) *GetDatasetResponseBody {
	s.GmtCreateTime = &v
	return s
}

func (s *GetDatasetResponseBody) SetGmtModifiedTime(v string) *GetDatasetResponseBody {
	s.GmtModifiedTime = &v
	return s
}

func (s *GetDatasetResponseBody) SetLabels(v []*Label) *GetDatasetResponseBody {
	s.Labels = v
	return s
}

func (s *GetDatasetResponseBody) SetName(v string) *GetDatasetResponseBody {
	s.Name = &v
	return s
}

func (s *GetDatasetResponseBody) SetOptions(v string) *GetDatasetResponseBody {
	s.Options = &v
	return s
}

func (s *GetDatasetResponseBody) SetOwnerId(v string) *GetDatasetResponseBody {
	s.OwnerId = &v
	return s
}

func (s *GetDatasetResponseBody) SetProperty(v string) *GetDatasetResponseBody {
	s.Property = &v
	return s
}

func (s *GetDatasetResponseBody) SetProviderType(v string) *GetDatasetResponseBody {
	s.ProviderType = &v
	return s
}

func (s *GetDatasetResponseBody) SetRequestId(v string) *GetDatasetResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDatasetResponseBody) SetSourceId(v string) *GetDatasetResponseBody {
	s.SourceId = &v
	return s
}

func (s *GetDatasetResponseBody) SetSourceType(v string) *GetDatasetResponseBody {
	s.SourceType = &v
	return s
}

func (s *GetDatasetResponseBody) SetUri(v string) *GetDatasetResponseBody {
	s.Uri = &v
	return s
}

func (s *GetDatasetResponseBody) SetUserId(v string) *GetDatasetResponseBody {
	s.UserId = &v
	return s
}

func (s *GetDatasetResponseBody) SetWorkspaceId(v string) *GetDatasetResponseBody {
	s.WorkspaceId = &v
	return s
}

type GetDatasetResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDatasetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDatasetResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDatasetResponse) GoString() string {
	return s.String()
}

func (s *GetDatasetResponse) SetHeaders(v map[string]*string) *GetDatasetResponse {
	s.Headers = v
	return s
}

func (s *GetDatasetResponse) SetStatusCode(v int32) *GetDatasetResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDatasetResponse) SetBody(v *GetDatasetResponseBody) *GetDatasetResponse {
	s.Body = v
	return s
}

type GetDefaultWorkspaceRequest struct {
	Verbose *bool `json:"Verbose,omitempty" xml:"Verbose,omitempty"`
}

func (s GetDefaultWorkspaceRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDefaultWorkspaceRequest) GoString() string {
	return s.String()
}

func (s *GetDefaultWorkspaceRequest) SetVerbose(v bool) *GetDefaultWorkspaceRequest {
	s.Verbose = &v
	return s
}

type GetDefaultWorkspaceResponseBody struct {
	Conditions      []*GetDefaultWorkspaceResponseBodyConditions `json:"Conditions,omitempty" xml:"Conditions,omitempty" type:"Repeated"`
	Creator         *string                                      `json:"Creator,omitempty" xml:"Creator,omitempty"`
	Description     *string                                      `json:"Description,omitempty" xml:"Description,omitempty"`
	DisplayName     *string                                      `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	EnvTypes        []*string                                    `json:"EnvTypes,omitempty" xml:"EnvTypes,omitempty" type:"Repeated"`
	GmtCreateTime   *string                                      `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	GmtModifiedTime *string                                      `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	Owner           *GetDefaultWorkspaceResponseBodyOwner        `json:"Owner,omitempty" xml:"Owner,omitempty" type:"Struct"`
	RequestId       *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Status          *string                                      `json:"Status,omitempty" xml:"Status,omitempty"`
	WorkspaceId     *string                                      `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
	WorkspaceName   *string                                      `json:"WorkspaceName,omitempty" xml:"WorkspaceName,omitempty"`
}

func (s GetDefaultWorkspaceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDefaultWorkspaceResponseBody) GoString() string {
	return s.String()
}

func (s *GetDefaultWorkspaceResponseBody) SetConditions(v []*GetDefaultWorkspaceResponseBodyConditions) *GetDefaultWorkspaceResponseBody {
	s.Conditions = v
	return s
}

func (s *GetDefaultWorkspaceResponseBody) SetCreator(v string) *GetDefaultWorkspaceResponseBody {
	s.Creator = &v
	return s
}

func (s *GetDefaultWorkspaceResponseBody) SetDescription(v string) *GetDefaultWorkspaceResponseBody {
	s.Description = &v
	return s
}

func (s *GetDefaultWorkspaceResponseBody) SetDisplayName(v string) *GetDefaultWorkspaceResponseBody {
	s.DisplayName = &v
	return s
}

func (s *GetDefaultWorkspaceResponseBody) SetEnvTypes(v []*string) *GetDefaultWorkspaceResponseBody {
	s.EnvTypes = v
	return s
}

func (s *GetDefaultWorkspaceResponseBody) SetGmtCreateTime(v string) *GetDefaultWorkspaceResponseBody {
	s.GmtCreateTime = &v
	return s
}

func (s *GetDefaultWorkspaceResponseBody) SetGmtModifiedTime(v string) *GetDefaultWorkspaceResponseBody {
	s.GmtModifiedTime = &v
	return s
}

func (s *GetDefaultWorkspaceResponseBody) SetOwner(v *GetDefaultWorkspaceResponseBodyOwner) *GetDefaultWorkspaceResponseBody {
	s.Owner = v
	return s
}

func (s *GetDefaultWorkspaceResponseBody) SetRequestId(v string) *GetDefaultWorkspaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDefaultWorkspaceResponseBody) SetStatus(v string) *GetDefaultWorkspaceResponseBody {
	s.Status = &v
	return s
}

func (s *GetDefaultWorkspaceResponseBody) SetWorkspaceId(v string) *GetDefaultWorkspaceResponseBody {
	s.WorkspaceId = &v
	return s
}

func (s *GetDefaultWorkspaceResponseBody) SetWorkspaceName(v string) *GetDefaultWorkspaceResponseBody {
	s.WorkspaceName = &v
	return s
}

type GetDefaultWorkspaceResponseBodyConditions struct {
	Code    *int64  `json:"Code,omitempty" xml:"Code,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Type    *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetDefaultWorkspaceResponseBodyConditions) String() string {
	return tea.Prettify(s)
}

func (s GetDefaultWorkspaceResponseBodyConditions) GoString() string {
	return s.String()
}

func (s *GetDefaultWorkspaceResponseBodyConditions) SetCode(v int64) *GetDefaultWorkspaceResponseBodyConditions {
	s.Code = &v
	return s
}

func (s *GetDefaultWorkspaceResponseBodyConditions) SetMessage(v string) *GetDefaultWorkspaceResponseBodyConditions {
	s.Message = &v
	return s
}

func (s *GetDefaultWorkspaceResponseBodyConditions) SetType(v string) *GetDefaultWorkspaceResponseBodyConditions {
	s.Type = &v
	return s
}

type GetDefaultWorkspaceResponseBodyOwner struct {
	UserId   *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	UserKp   *string `json:"UserKp,omitempty" xml:"UserKp,omitempty"`
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s GetDefaultWorkspaceResponseBodyOwner) String() string {
	return tea.Prettify(s)
}

func (s GetDefaultWorkspaceResponseBodyOwner) GoString() string {
	return s.String()
}

func (s *GetDefaultWorkspaceResponseBodyOwner) SetUserId(v string) *GetDefaultWorkspaceResponseBodyOwner {
	s.UserId = &v
	return s
}

func (s *GetDefaultWorkspaceResponseBodyOwner) SetUserKp(v string) *GetDefaultWorkspaceResponseBodyOwner {
	s.UserKp = &v
	return s
}

func (s *GetDefaultWorkspaceResponseBodyOwner) SetUserName(v string) *GetDefaultWorkspaceResponseBodyOwner {
	s.UserName = &v
	return s
}

type GetDefaultWorkspaceResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDefaultWorkspaceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDefaultWorkspaceResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDefaultWorkspaceResponse) GoString() string {
	return s.String()
}

func (s *GetDefaultWorkspaceResponse) SetHeaders(v map[string]*string) *GetDefaultWorkspaceResponse {
	s.Headers = v
	return s
}

func (s *GetDefaultWorkspaceResponse) SetStatusCode(v int32) *GetDefaultWorkspaceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDefaultWorkspaceResponse) SetBody(v *GetDefaultWorkspaceResponseBody) *GetDefaultWorkspaceResponse {
	s.Body = v
	return s
}

type GetImageRequest struct {
	Verbose *bool `json:"Verbose,omitempty" xml:"Verbose,omitempty"`
}

func (s GetImageRequest) String() string {
	return tea.Prettify(s)
}

func (s GetImageRequest) GoString() string {
	return s.String()
}

func (s *GetImageRequest) SetVerbose(v bool) *GetImageRequest {
	s.Verbose = &v
	return s
}

type GetImageResponseBody struct {
	Accessibility   *string                       `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	Description     *string                       `json:"Description,omitempty" xml:"Description,omitempty"`
	GmtCreateTime   *string                       `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	GmtModifiedTime *string                       `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	ImageUri        *string                       `json:"ImageUri,omitempty" xml:"ImageUri,omitempty"`
	Labels          []*GetImageResponseBodyLabels `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	Name            *string                       `json:"Name,omitempty" xml:"Name,omitempty"`
	ParentUserId    *string                       `json:"ParentUserId,omitempty" xml:"ParentUserId,omitempty"`
	RequestId       *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Size            *int32                        `json:"Size,omitempty" xml:"Size,omitempty"`
	UserId          *string                       `json:"UserId,omitempty" xml:"UserId,omitempty"`
	WorkspaceId     *string                       `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s GetImageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetImageResponseBody) GoString() string {
	return s.String()
}

func (s *GetImageResponseBody) SetAccessibility(v string) *GetImageResponseBody {
	s.Accessibility = &v
	return s
}

func (s *GetImageResponseBody) SetDescription(v string) *GetImageResponseBody {
	s.Description = &v
	return s
}

func (s *GetImageResponseBody) SetGmtCreateTime(v string) *GetImageResponseBody {
	s.GmtCreateTime = &v
	return s
}

func (s *GetImageResponseBody) SetGmtModifiedTime(v string) *GetImageResponseBody {
	s.GmtModifiedTime = &v
	return s
}

func (s *GetImageResponseBody) SetImageUri(v string) *GetImageResponseBody {
	s.ImageUri = &v
	return s
}

func (s *GetImageResponseBody) SetLabels(v []*GetImageResponseBodyLabels) *GetImageResponseBody {
	s.Labels = v
	return s
}

func (s *GetImageResponseBody) SetName(v string) *GetImageResponseBody {
	s.Name = &v
	return s
}

func (s *GetImageResponseBody) SetParentUserId(v string) *GetImageResponseBody {
	s.ParentUserId = &v
	return s
}

func (s *GetImageResponseBody) SetRequestId(v string) *GetImageResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetImageResponseBody) SetSize(v int32) *GetImageResponseBody {
	s.Size = &v
	return s
}

func (s *GetImageResponseBody) SetUserId(v string) *GetImageResponseBody {
	s.UserId = &v
	return s
}

func (s *GetImageResponseBody) SetWorkspaceId(v string) *GetImageResponseBody {
	s.WorkspaceId = &v
	return s
}

type GetImageResponseBodyLabels struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetImageResponseBodyLabels) String() string {
	return tea.Prettify(s)
}

func (s GetImageResponseBodyLabels) GoString() string {
	return s.String()
}

func (s *GetImageResponseBodyLabels) SetKey(v string) *GetImageResponseBodyLabels {
	s.Key = &v
	return s
}

func (s *GetImageResponseBodyLabels) SetValue(v string) *GetImageResponseBodyLabels {
	s.Value = &v
	return s
}

type GetImageResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetImageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetImageResponse) String() string {
	return tea.Prettify(s)
}

func (s GetImageResponse) GoString() string {
	return s.String()
}

func (s *GetImageResponse) SetHeaders(v map[string]*string) *GetImageResponse {
	s.Headers = v
	return s
}

func (s *GetImageResponse) SetStatusCode(v int32) *GetImageResponse {
	s.StatusCode = &v
	return s
}

func (s *GetImageResponse) SetBody(v *GetImageResponseBody) *GetImageResponse {
	s.Body = v
	return s
}

type GetMemberRequest struct {
	MemberId *string `json:"MemberId,omitempty" xml:"MemberId,omitempty"`
	UserId   *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s GetMemberRequest) String() string {
	return tea.Prettify(s)
}

func (s GetMemberRequest) GoString() string {
	return s.String()
}

func (s *GetMemberRequest) SetMemberId(v string) *GetMemberRequest {
	s.MemberId = &v
	return s
}

func (s *GetMemberRequest) SetUserId(v string) *GetMemberRequest {
	s.UserId = &v
	return s
}

type GetMemberResponseBody struct {
	DisplayName   *string   `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	GmtCreateTime *string   `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	MemberId      *string   `json:"MemberId,omitempty" xml:"MemberId,omitempty"`
	MemberName    *string   `json:"MemberName,omitempty" xml:"MemberName,omitempty"`
	RequestId     *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Roles         []*string `json:"Roles,omitempty" xml:"Roles,omitempty" type:"Repeated"`
	UserId        *string   `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s GetMemberResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetMemberResponseBody) GoString() string {
	return s.String()
}

func (s *GetMemberResponseBody) SetDisplayName(v string) *GetMemberResponseBody {
	s.DisplayName = &v
	return s
}

func (s *GetMemberResponseBody) SetGmtCreateTime(v string) *GetMemberResponseBody {
	s.GmtCreateTime = &v
	return s
}

func (s *GetMemberResponseBody) SetMemberId(v string) *GetMemberResponseBody {
	s.MemberId = &v
	return s
}

func (s *GetMemberResponseBody) SetMemberName(v string) *GetMemberResponseBody {
	s.MemberName = &v
	return s
}

func (s *GetMemberResponseBody) SetRequestId(v string) *GetMemberResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMemberResponseBody) SetRoles(v []*string) *GetMemberResponseBody {
	s.Roles = v
	return s
}

func (s *GetMemberResponseBody) SetUserId(v string) *GetMemberResponseBody {
	s.UserId = &v
	return s
}

type GetMemberResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMemberResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMemberResponse) String() string {
	return tea.Prettify(s)
}

func (s GetMemberResponse) GoString() string {
	return s.String()
}

func (s *GetMemberResponse) SetHeaders(v map[string]*string) *GetMemberResponse {
	s.Headers = v
	return s
}

func (s *GetMemberResponse) SetStatusCode(v int32) *GetMemberResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMemberResponse) SetBody(v *GetMemberResponseBody) *GetMemberResponse {
	s.Body = v
	return s
}

type GetModelResponseBody struct {
	Accessibility    *string                `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	Domain           *string                `json:"Domain,omitempty" xml:"Domain,omitempty"`
	ExtraInfo        map[string]interface{} `json:"ExtraInfo,omitempty" xml:"ExtraInfo,omitempty"`
	GmtCreateTime    *string                `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	GmtModifiedTime  *string                `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	Labels           []*Label               `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	LatestVersion    *ModelVersion          `json:"LatestVersion,omitempty" xml:"LatestVersion,omitempty"`
	ModelDescription *string                `json:"ModelDescription,omitempty" xml:"ModelDescription,omitempty"`
	ModelDoc         *string                `json:"ModelDoc,omitempty" xml:"ModelDoc,omitempty"`
	ModelId          *string                `json:"ModelId,omitempty" xml:"ModelId,omitempty"`
	ModelName        *string                `json:"ModelName,omitempty" xml:"ModelName,omitempty"`
	ModelType        *string                `json:"ModelType,omitempty" xml:"ModelType,omitempty"`
	OrderNumber      *int64                 `json:"OrderNumber,omitempty" xml:"OrderNumber,omitempty"`
	Origin           *string                `json:"Origin,omitempty" xml:"Origin,omitempty"`
	OwnerId          *string                `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	Provider         *string                `json:"Provider,omitempty" xml:"Provider,omitempty"`
	RequestId        *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Task             *string                `json:"Task,omitempty" xml:"Task,omitempty"`
	UserId           *string                `json:"UserId,omitempty" xml:"UserId,omitempty"`
	WorkspaceId      *string                `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s GetModelResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetModelResponseBody) GoString() string {
	return s.String()
}

func (s *GetModelResponseBody) SetAccessibility(v string) *GetModelResponseBody {
	s.Accessibility = &v
	return s
}

func (s *GetModelResponseBody) SetDomain(v string) *GetModelResponseBody {
	s.Domain = &v
	return s
}

func (s *GetModelResponseBody) SetExtraInfo(v map[string]interface{}) *GetModelResponseBody {
	s.ExtraInfo = v
	return s
}

func (s *GetModelResponseBody) SetGmtCreateTime(v string) *GetModelResponseBody {
	s.GmtCreateTime = &v
	return s
}

func (s *GetModelResponseBody) SetGmtModifiedTime(v string) *GetModelResponseBody {
	s.GmtModifiedTime = &v
	return s
}

func (s *GetModelResponseBody) SetLabels(v []*Label) *GetModelResponseBody {
	s.Labels = v
	return s
}

func (s *GetModelResponseBody) SetLatestVersion(v *ModelVersion) *GetModelResponseBody {
	s.LatestVersion = v
	return s
}

func (s *GetModelResponseBody) SetModelDescription(v string) *GetModelResponseBody {
	s.ModelDescription = &v
	return s
}

func (s *GetModelResponseBody) SetModelDoc(v string) *GetModelResponseBody {
	s.ModelDoc = &v
	return s
}

func (s *GetModelResponseBody) SetModelId(v string) *GetModelResponseBody {
	s.ModelId = &v
	return s
}

func (s *GetModelResponseBody) SetModelName(v string) *GetModelResponseBody {
	s.ModelName = &v
	return s
}

func (s *GetModelResponseBody) SetModelType(v string) *GetModelResponseBody {
	s.ModelType = &v
	return s
}

func (s *GetModelResponseBody) SetOrderNumber(v int64) *GetModelResponseBody {
	s.OrderNumber = &v
	return s
}

func (s *GetModelResponseBody) SetOrigin(v string) *GetModelResponseBody {
	s.Origin = &v
	return s
}

func (s *GetModelResponseBody) SetOwnerId(v string) *GetModelResponseBody {
	s.OwnerId = &v
	return s
}

func (s *GetModelResponseBody) SetProvider(v string) *GetModelResponseBody {
	s.Provider = &v
	return s
}

func (s *GetModelResponseBody) SetRequestId(v string) *GetModelResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetModelResponseBody) SetTask(v string) *GetModelResponseBody {
	s.Task = &v
	return s
}

func (s *GetModelResponseBody) SetUserId(v string) *GetModelResponseBody {
	s.UserId = &v
	return s
}

func (s *GetModelResponseBody) SetWorkspaceId(v string) *GetModelResponseBody {
	s.WorkspaceId = &v
	return s
}

type GetModelResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetModelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetModelResponse) String() string {
	return tea.Prettify(s)
}

func (s GetModelResponse) GoString() string {
	return s.String()
}

func (s *GetModelResponse) SetHeaders(v map[string]*string) *GetModelResponse {
	s.Headers = v
	return s
}

func (s *GetModelResponse) SetStatusCode(v int32) *GetModelResponse {
	s.StatusCode = &v
	return s
}

func (s *GetModelResponse) SetBody(v *GetModelResponseBody) *GetModelResponse {
	s.Body = v
	return s
}

type GetModelVersionResponseBody struct {
	ApprovalStatus     *string                `json:"ApprovalStatus,omitempty" xml:"ApprovalStatus,omitempty"`
	EvaluationSpec     map[string]interface{} `json:"EvaluationSpec,omitempty" xml:"EvaluationSpec,omitempty"`
	ExtraInfo          map[string]interface{} `json:"ExtraInfo,omitempty" xml:"ExtraInfo,omitempty"`
	FormatType         *string                `json:"FormatType,omitempty" xml:"FormatType,omitempty"`
	FrameworkType      *string                `json:"FrameworkType,omitempty" xml:"FrameworkType,omitempty"`
	GmtCreateTime      *string                `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	GmtModifiedTime    *string                `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	InferenceSpec      map[string]interface{} `json:"InferenceSpec,omitempty" xml:"InferenceSpec,omitempty"`
	Labels             []*Label               `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	Metrics            map[string]interface{} `json:"Metrics,omitempty" xml:"Metrics,omitempty"`
	Options            *string                `json:"Options,omitempty" xml:"Options,omitempty"`
	OwnerId            *string                `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RequestId          *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SourceId           *string                `json:"SourceId,omitempty" xml:"SourceId,omitempty"`
	SourceType         *string                `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	TrainingSpec       map[string]interface{} `json:"TrainingSpec,omitempty" xml:"TrainingSpec,omitempty"`
	Uri                *string                `json:"Uri,omitempty" xml:"Uri,omitempty"`
	UserId             *string                `json:"UserId,omitempty" xml:"UserId,omitempty"`
	VersionDescription *string                `json:"VersionDescription,omitempty" xml:"VersionDescription,omitempty"`
	VersionName        *string                `json:"VersionName,omitempty" xml:"VersionName,omitempty"`
}

func (s GetModelVersionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetModelVersionResponseBody) GoString() string {
	return s.String()
}

func (s *GetModelVersionResponseBody) SetApprovalStatus(v string) *GetModelVersionResponseBody {
	s.ApprovalStatus = &v
	return s
}

func (s *GetModelVersionResponseBody) SetEvaluationSpec(v map[string]interface{}) *GetModelVersionResponseBody {
	s.EvaluationSpec = v
	return s
}

func (s *GetModelVersionResponseBody) SetExtraInfo(v map[string]interface{}) *GetModelVersionResponseBody {
	s.ExtraInfo = v
	return s
}

func (s *GetModelVersionResponseBody) SetFormatType(v string) *GetModelVersionResponseBody {
	s.FormatType = &v
	return s
}

func (s *GetModelVersionResponseBody) SetFrameworkType(v string) *GetModelVersionResponseBody {
	s.FrameworkType = &v
	return s
}

func (s *GetModelVersionResponseBody) SetGmtCreateTime(v string) *GetModelVersionResponseBody {
	s.GmtCreateTime = &v
	return s
}

func (s *GetModelVersionResponseBody) SetGmtModifiedTime(v string) *GetModelVersionResponseBody {
	s.GmtModifiedTime = &v
	return s
}

func (s *GetModelVersionResponseBody) SetInferenceSpec(v map[string]interface{}) *GetModelVersionResponseBody {
	s.InferenceSpec = v
	return s
}

func (s *GetModelVersionResponseBody) SetLabels(v []*Label) *GetModelVersionResponseBody {
	s.Labels = v
	return s
}

func (s *GetModelVersionResponseBody) SetMetrics(v map[string]interface{}) *GetModelVersionResponseBody {
	s.Metrics = v
	return s
}

func (s *GetModelVersionResponseBody) SetOptions(v string) *GetModelVersionResponseBody {
	s.Options = &v
	return s
}

func (s *GetModelVersionResponseBody) SetOwnerId(v string) *GetModelVersionResponseBody {
	s.OwnerId = &v
	return s
}

func (s *GetModelVersionResponseBody) SetRequestId(v string) *GetModelVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetModelVersionResponseBody) SetSourceId(v string) *GetModelVersionResponseBody {
	s.SourceId = &v
	return s
}

func (s *GetModelVersionResponseBody) SetSourceType(v string) *GetModelVersionResponseBody {
	s.SourceType = &v
	return s
}

func (s *GetModelVersionResponseBody) SetTrainingSpec(v map[string]interface{}) *GetModelVersionResponseBody {
	s.TrainingSpec = v
	return s
}

func (s *GetModelVersionResponseBody) SetUri(v string) *GetModelVersionResponseBody {
	s.Uri = &v
	return s
}

func (s *GetModelVersionResponseBody) SetUserId(v string) *GetModelVersionResponseBody {
	s.UserId = &v
	return s
}

func (s *GetModelVersionResponseBody) SetVersionDescription(v string) *GetModelVersionResponseBody {
	s.VersionDescription = &v
	return s
}

func (s *GetModelVersionResponseBody) SetVersionName(v string) *GetModelVersionResponseBody {
	s.VersionName = &v
	return s
}

type GetModelVersionResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetModelVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetModelVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s GetModelVersionResponse) GoString() string {
	return s.String()
}

func (s *GetModelVersionResponse) SetHeaders(v map[string]*string) *GetModelVersionResponse {
	s.Headers = v
	return s
}

func (s *GetModelVersionResponse) SetStatusCode(v int32) *GetModelVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *GetModelVersionResponse) SetBody(v *GetModelVersionResponseBody) *GetModelVersionResponse {
	s.Body = v
	return s
}

type GetPermissionRequest struct {
	Accessibility *string `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	Creator       *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	Option        *string `json:"Option,omitempty" xml:"Option,omitempty"`
	Resource      *string `json:"Resource,omitempty" xml:"Resource,omitempty"`
}

func (s GetPermissionRequest) String() string {
	return tea.Prettify(s)
}

func (s GetPermissionRequest) GoString() string {
	return s.String()
}

func (s *GetPermissionRequest) SetAccessibility(v string) *GetPermissionRequest {
	s.Accessibility = &v
	return s
}

func (s *GetPermissionRequest) SetCreator(v string) *GetPermissionRequest {
	s.Creator = &v
	return s
}

func (s *GetPermissionRequest) SetOption(v string) *GetPermissionRequest {
	s.Option = &v
	return s
}

func (s *GetPermissionRequest) SetResource(v string) *GetPermissionRequest {
	s.Resource = &v
	return s
}

type GetPermissionResponseBody struct {
	PermissionCode  *string                                     `json:"PermissionCode,omitempty" xml:"PermissionCode,omitempty"`
	PermissionRules []*GetPermissionResponseBodyPermissionRules `json:"PermissionRules,omitempty" xml:"PermissionRules,omitempty" type:"Repeated"`
	RequestId       *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetPermissionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetPermissionResponseBody) GoString() string {
	return s.String()
}

func (s *GetPermissionResponseBody) SetPermissionCode(v string) *GetPermissionResponseBody {
	s.PermissionCode = &v
	return s
}

func (s *GetPermissionResponseBody) SetPermissionRules(v []*GetPermissionResponseBodyPermissionRules) *GetPermissionResponseBody {
	s.PermissionRules = v
	return s
}

func (s *GetPermissionResponseBody) SetRequestId(v string) *GetPermissionResponseBody {
	s.RequestId = &v
	return s
}

type GetPermissionResponseBodyPermissionRules struct {
	Accessibility    *string `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	EntityAccessType *string `json:"EntityAccessType,omitempty" xml:"EntityAccessType,omitempty"`
}

func (s GetPermissionResponseBodyPermissionRules) String() string {
	return tea.Prettify(s)
}

func (s GetPermissionResponseBodyPermissionRules) GoString() string {
	return s.String()
}

func (s *GetPermissionResponseBodyPermissionRules) SetAccessibility(v string) *GetPermissionResponseBodyPermissionRules {
	s.Accessibility = &v
	return s
}

func (s *GetPermissionResponseBodyPermissionRules) SetEntityAccessType(v string) *GetPermissionResponseBodyPermissionRules {
	s.EntityAccessType = &v
	return s
}

type GetPermissionResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPermissionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPermissionResponse) String() string {
	return tea.Prettify(s)
}

func (s GetPermissionResponse) GoString() string {
	return s.String()
}

func (s *GetPermissionResponse) SetHeaders(v map[string]*string) *GetPermissionResponse {
	s.Headers = v
	return s
}

func (s *GetPermissionResponse) SetStatusCode(v int32) *GetPermissionResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPermissionResponse) SetBody(v *GetPermissionResponseBody) *GetPermissionResponse {
	s.Body = v
	return s
}

type GetServiceTemplateResponseBody struct {
	GmtCreateTime              *string                `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	GmtModifiedTime            *string                `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	InferenceSpec              map[string]interface{} `json:"InferenceSpec,omitempty" xml:"InferenceSpec,omitempty"`
	Labels                     []*Label               `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	OwnerId                    *string                `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	Provider                   *string                `json:"Provider,omitempty" xml:"Provider,omitempty"`
	RequestId                  *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ServiceTemplateDescription *string                `json:"ServiceTemplateDescription,omitempty" xml:"ServiceTemplateDescription,omitempty"`
	ServiceTemplateDoc         *string                `json:"ServiceTemplateDoc,omitempty" xml:"ServiceTemplateDoc,omitempty"`
	ServiceTemplateId          *string                `json:"ServiceTemplateId,omitempty" xml:"ServiceTemplateId,omitempty"`
	ServiceTemplateName        *string                `json:"ServiceTemplateName,omitempty" xml:"ServiceTemplateName,omitempty"`
	UserId                     *string                `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s GetServiceTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetServiceTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *GetServiceTemplateResponseBody) SetGmtCreateTime(v string) *GetServiceTemplateResponseBody {
	s.GmtCreateTime = &v
	return s
}

func (s *GetServiceTemplateResponseBody) SetGmtModifiedTime(v string) *GetServiceTemplateResponseBody {
	s.GmtModifiedTime = &v
	return s
}

func (s *GetServiceTemplateResponseBody) SetInferenceSpec(v map[string]interface{}) *GetServiceTemplateResponseBody {
	s.InferenceSpec = v
	return s
}

func (s *GetServiceTemplateResponseBody) SetLabels(v []*Label) *GetServiceTemplateResponseBody {
	s.Labels = v
	return s
}

func (s *GetServiceTemplateResponseBody) SetOwnerId(v string) *GetServiceTemplateResponseBody {
	s.OwnerId = &v
	return s
}

func (s *GetServiceTemplateResponseBody) SetProvider(v string) *GetServiceTemplateResponseBody {
	s.Provider = &v
	return s
}

func (s *GetServiceTemplateResponseBody) SetRequestId(v string) *GetServiceTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetServiceTemplateResponseBody) SetServiceTemplateDescription(v string) *GetServiceTemplateResponseBody {
	s.ServiceTemplateDescription = &v
	return s
}

func (s *GetServiceTemplateResponseBody) SetServiceTemplateDoc(v string) *GetServiceTemplateResponseBody {
	s.ServiceTemplateDoc = &v
	return s
}

func (s *GetServiceTemplateResponseBody) SetServiceTemplateId(v string) *GetServiceTemplateResponseBody {
	s.ServiceTemplateId = &v
	return s
}

func (s *GetServiceTemplateResponseBody) SetServiceTemplateName(v string) *GetServiceTemplateResponseBody {
	s.ServiceTemplateName = &v
	return s
}

func (s *GetServiceTemplateResponseBody) SetUserId(v string) *GetServiceTemplateResponseBody {
	s.UserId = &v
	return s
}

type GetServiceTemplateResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetServiceTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetServiceTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s GetServiceTemplateResponse) GoString() string {
	return s.String()
}

func (s *GetServiceTemplateResponse) SetHeaders(v map[string]*string) *GetServiceTemplateResponse {
	s.Headers = v
	return s
}

func (s *GetServiceTemplateResponse) SetStatusCode(v int32) *GetServiceTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *GetServiceTemplateResponse) SetBody(v *GetServiceTemplateResponseBody) *GetServiceTemplateResponse {
	s.Body = v
	return s
}

type GetWorkspaceRequest struct {
	Verbose *bool `json:"Verbose,omitempty" xml:"Verbose,omitempty"`
}

func (s GetWorkspaceRequest) String() string {
	return tea.Prettify(s)
}

func (s GetWorkspaceRequest) GoString() string {
	return s.String()
}

func (s *GetWorkspaceRequest) SetVerbose(v bool) *GetWorkspaceRequest {
	s.Verbose = &v
	return s
}

type GetWorkspaceResponseBody struct {
	AdminNames      []*string                      `json:"AdminNames,omitempty" xml:"AdminNames,omitempty" type:"Repeated"`
	Creator         *string                        `json:"Creator,omitempty" xml:"Creator,omitempty"`
	Description     *string                        `json:"Description,omitempty" xml:"Description,omitempty"`
	DisplayName     *string                        `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	EnvTypes        []*string                      `json:"EnvTypes,omitempty" xml:"EnvTypes,omitempty" type:"Repeated"`
	ExtraInfos      map[string]interface{}         `json:"ExtraInfos,omitempty" xml:"ExtraInfos,omitempty"`
	GmtCreateTime   *string                        `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	GmtModifiedTime *string                        `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	IsDefault       *bool                          `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	Owner           *GetWorkspaceResponseBodyOwner `json:"Owner,omitempty" xml:"Owner,omitempty" type:"Struct"`
	RequestId       *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Status          *string                        `json:"Status,omitempty" xml:"Status,omitempty"`
	WorkspaceId     *string                        `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
	WorkspaceName   *string                        `json:"WorkspaceName,omitempty" xml:"WorkspaceName,omitempty"`
}

func (s GetWorkspaceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetWorkspaceResponseBody) GoString() string {
	return s.String()
}

func (s *GetWorkspaceResponseBody) SetAdminNames(v []*string) *GetWorkspaceResponseBody {
	s.AdminNames = v
	return s
}

func (s *GetWorkspaceResponseBody) SetCreator(v string) *GetWorkspaceResponseBody {
	s.Creator = &v
	return s
}

func (s *GetWorkspaceResponseBody) SetDescription(v string) *GetWorkspaceResponseBody {
	s.Description = &v
	return s
}

func (s *GetWorkspaceResponseBody) SetDisplayName(v string) *GetWorkspaceResponseBody {
	s.DisplayName = &v
	return s
}

func (s *GetWorkspaceResponseBody) SetEnvTypes(v []*string) *GetWorkspaceResponseBody {
	s.EnvTypes = v
	return s
}

func (s *GetWorkspaceResponseBody) SetExtraInfos(v map[string]interface{}) *GetWorkspaceResponseBody {
	s.ExtraInfos = v
	return s
}

func (s *GetWorkspaceResponseBody) SetGmtCreateTime(v string) *GetWorkspaceResponseBody {
	s.GmtCreateTime = &v
	return s
}

func (s *GetWorkspaceResponseBody) SetGmtModifiedTime(v string) *GetWorkspaceResponseBody {
	s.GmtModifiedTime = &v
	return s
}

func (s *GetWorkspaceResponseBody) SetIsDefault(v bool) *GetWorkspaceResponseBody {
	s.IsDefault = &v
	return s
}

func (s *GetWorkspaceResponseBody) SetOwner(v *GetWorkspaceResponseBodyOwner) *GetWorkspaceResponseBody {
	s.Owner = v
	return s
}

func (s *GetWorkspaceResponseBody) SetRequestId(v string) *GetWorkspaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetWorkspaceResponseBody) SetStatus(v string) *GetWorkspaceResponseBody {
	s.Status = &v
	return s
}

func (s *GetWorkspaceResponseBody) SetWorkspaceId(v string) *GetWorkspaceResponseBody {
	s.WorkspaceId = &v
	return s
}

func (s *GetWorkspaceResponseBody) SetWorkspaceName(v string) *GetWorkspaceResponseBody {
	s.WorkspaceName = &v
	return s
}

type GetWorkspaceResponseBodyOwner struct {
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	UserId      *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	UserKp      *string `json:"UserKp,omitempty" xml:"UserKp,omitempty"`
	UserName    *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s GetWorkspaceResponseBodyOwner) String() string {
	return tea.Prettify(s)
}

func (s GetWorkspaceResponseBodyOwner) GoString() string {
	return s.String()
}

func (s *GetWorkspaceResponseBodyOwner) SetDisplayName(v string) *GetWorkspaceResponseBodyOwner {
	s.DisplayName = &v
	return s
}

func (s *GetWorkspaceResponseBodyOwner) SetUserId(v string) *GetWorkspaceResponseBodyOwner {
	s.UserId = &v
	return s
}

func (s *GetWorkspaceResponseBodyOwner) SetUserKp(v string) *GetWorkspaceResponseBodyOwner {
	s.UserKp = &v
	return s
}

func (s *GetWorkspaceResponseBodyOwner) SetUserName(v string) *GetWorkspaceResponseBodyOwner {
	s.UserName = &v
	return s
}

type GetWorkspaceResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetWorkspaceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetWorkspaceResponse) String() string {
	return tea.Prettify(s)
}

func (s GetWorkspaceResponse) GoString() string {
	return s.String()
}

func (s *GetWorkspaceResponse) SetHeaders(v map[string]*string) *GetWorkspaceResponse {
	s.Headers = v
	return s
}

func (s *GetWorkspaceResponse) SetStatusCode(v int32) *GetWorkspaceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetWorkspaceResponse) SetBody(v *GetWorkspaceResponseBody) *GetWorkspaceResponse {
	s.Body = v
	return s
}

type ListCodeSourcesRequest struct {
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	Order       *string `json:"Order,omitempty" xml:"Order,omitempty"`
	PageNumber  *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SortBy      *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListCodeSourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListCodeSourcesRequest) GoString() string {
	return s.String()
}

func (s *ListCodeSourcesRequest) SetDisplayName(v string) *ListCodeSourcesRequest {
	s.DisplayName = &v
	return s
}

func (s *ListCodeSourcesRequest) SetOrder(v string) *ListCodeSourcesRequest {
	s.Order = &v
	return s
}

func (s *ListCodeSourcesRequest) SetPageNumber(v int32) *ListCodeSourcesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListCodeSourcesRequest) SetPageSize(v int32) *ListCodeSourcesRequest {
	s.PageSize = &v
	return s
}

func (s *ListCodeSourcesRequest) SetSortBy(v string) *ListCodeSourcesRequest {
	s.SortBy = &v
	return s
}

func (s *ListCodeSourcesRequest) SetWorkspaceId(v string) *ListCodeSourcesRequest {
	s.WorkspaceId = &v
	return s
}

type ListCodeSourcesResponseBody struct {
	CodeSources []*CodeSourceItem `json:"CodeSources,omitempty" xml:"CodeSources,omitempty" type:"Repeated"`
	RequestId   *string           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount  *int64            `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListCodeSourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListCodeSourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ListCodeSourcesResponseBody) SetCodeSources(v []*CodeSourceItem) *ListCodeSourcesResponseBody {
	s.CodeSources = v
	return s
}

func (s *ListCodeSourcesResponseBody) SetRequestId(v string) *ListCodeSourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCodeSourcesResponseBody) SetTotalCount(v int64) *ListCodeSourcesResponseBody {
	s.TotalCount = &v
	return s
}

type ListCodeSourcesResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCodeSourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCodeSourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListCodeSourcesResponse) GoString() string {
	return s.String()
}

func (s *ListCodeSourcesResponse) SetHeaders(v map[string]*string) *ListCodeSourcesResponse {
	s.Headers = v
	return s
}

func (s *ListCodeSourcesResponse) SetStatusCode(v int32) *ListCodeSourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCodeSourcesResponse) SetBody(v *ListCodeSourcesResponseBody) *ListCodeSourcesResponse {
	s.Body = v
	return s
}

type ListDatasetsRequest struct {
	DataSourceTypes *string `json:"DataSourceTypes,omitempty" xml:"DataSourceTypes,omitempty"`
	DataTypes       *string `json:"DataTypes,omitempty" xml:"DataTypes,omitempty"`
	Label           *string `json:"Label,omitempty" xml:"Label,omitempty"`
	Name            *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Order           *string `json:"Order,omitempty" xml:"Order,omitempty"`
	PageNumber      *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize        *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Properties      *string `json:"Properties,omitempty" xml:"Properties,omitempty"`
	Provider        *string `json:"Provider,omitempty" xml:"Provider,omitempty"`
	SourceId        *string `json:"SourceId,omitempty" xml:"SourceId,omitempty"`
	SourceTypes     *string `json:"SourceTypes,omitempty" xml:"SourceTypes,omitempty"`
	WorkspaceId     *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListDatasetsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDatasetsRequest) GoString() string {
	return s.String()
}

func (s *ListDatasetsRequest) SetDataSourceTypes(v string) *ListDatasetsRequest {
	s.DataSourceTypes = &v
	return s
}

func (s *ListDatasetsRequest) SetDataTypes(v string) *ListDatasetsRequest {
	s.DataTypes = &v
	return s
}

func (s *ListDatasetsRequest) SetLabel(v string) *ListDatasetsRequest {
	s.Label = &v
	return s
}

func (s *ListDatasetsRequest) SetName(v string) *ListDatasetsRequest {
	s.Name = &v
	return s
}

func (s *ListDatasetsRequest) SetOrder(v string) *ListDatasetsRequest {
	s.Order = &v
	return s
}

func (s *ListDatasetsRequest) SetPageNumber(v int32) *ListDatasetsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListDatasetsRequest) SetPageSize(v int32) *ListDatasetsRequest {
	s.PageSize = &v
	return s
}

func (s *ListDatasetsRequest) SetProperties(v string) *ListDatasetsRequest {
	s.Properties = &v
	return s
}

func (s *ListDatasetsRequest) SetProvider(v string) *ListDatasetsRequest {
	s.Provider = &v
	return s
}

func (s *ListDatasetsRequest) SetSourceId(v string) *ListDatasetsRequest {
	s.SourceId = &v
	return s
}

func (s *ListDatasetsRequest) SetSourceTypes(v string) *ListDatasetsRequest {
	s.SourceTypes = &v
	return s
}

func (s *ListDatasetsRequest) SetWorkspaceId(v string) *ListDatasetsRequest {
	s.WorkspaceId = &v
	return s
}

type ListDatasetsResponseBody struct {
	Datasets   []*Dataset `json:"Datasets,omitempty" xml:"Datasets,omitempty" type:"Repeated"`
	RequestId  *string    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int64     `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListDatasetsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDatasetsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDatasetsResponseBody) SetDatasets(v []*Dataset) *ListDatasetsResponseBody {
	s.Datasets = v
	return s
}

func (s *ListDatasetsResponseBody) SetRequestId(v string) *ListDatasetsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDatasetsResponseBody) SetTotalCount(v int64) *ListDatasetsResponseBody {
	s.TotalCount = &v
	return s
}

type ListDatasetsResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDatasetsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDatasetsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDatasetsResponse) GoString() string {
	return s.String()
}

func (s *ListDatasetsResponse) SetHeaders(v map[string]*string) *ListDatasetsResponse {
	s.Headers = v
	return s
}

func (s *ListDatasetsResponse) SetStatusCode(v int32) *ListDatasetsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDatasetsResponse) SetBody(v *ListDatasetsResponseBody) *ListDatasetsResponse {
	s.Body = v
	return s
}

type ListImageLabelsRequest struct {
	ImageId     *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	LabelFilter *string `json:"LabelFilter,omitempty" xml:"LabelFilter,omitempty"`
	LabelKeys   *string `json:"LabelKeys,omitempty" xml:"LabelKeys,omitempty"`
	Region      *string `json:"Region,omitempty" xml:"Region,omitempty"`
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListImageLabelsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListImageLabelsRequest) GoString() string {
	return s.String()
}

func (s *ListImageLabelsRequest) SetImageId(v string) *ListImageLabelsRequest {
	s.ImageId = &v
	return s
}

func (s *ListImageLabelsRequest) SetLabelFilter(v string) *ListImageLabelsRequest {
	s.LabelFilter = &v
	return s
}

func (s *ListImageLabelsRequest) SetLabelKeys(v string) *ListImageLabelsRequest {
	s.LabelKeys = &v
	return s
}

func (s *ListImageLabelsRequest) SetRegion(v string) *ListImageLabelsRequest {
	s.Region = &v
	return s
}

func (s *ListImageLabelsRequest) SetWorkspaceId(v string) *ListImageLabelsRequest {
	s.WorkspaceId = &v
	return s
}

type ListImageLabelsResponseBody struct {
	Labels     []*ListImageLabelsResponseBodyLabels `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	RequestId  *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int64                               `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListImageLabelsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListImageLabelsResponseBody) GoString() string {
	return s.String()
}

func (s *ListImageLabelsResponseBody) SetLabels(v []*ListImageLabelsResponseBodyLabels) *ListImageLabelsResponseBody {
	s.Labels = v
	return s
}

func (s *ListImageLabelsResponseBody) SetRequestId(v string) *ListImageLabelsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListImageLabelsResponseBody) SetTotalCount(v int64) *ListImageLabelsResponseBody {
	s.TotalCount = &v
	return s
}

type ListImageLabelsResponseBodyLabels struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListImageLabelsResponseBodyLabels) String() string {
	return tea.Prettify(s)
}

func (s ListImageLabelsResponseBodyLabels) GoString() string {
	return s.String()
}

func (s *ListImageLabelsResponseBodyLabels) SetKey(v string) *ListImageLabelsResponseBodyLabels {
	s.Key = &v
	return s
}

func (s *ListImageLabelsResponseBodyLabels) SetValue(v string) *ListImageLabelsResponseBodyLabels {
	s.Value = &v
	return s
}

type ListImageLabelsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListImageLabelsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListImageLabelsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListImageLabelsResponse) GoString() string {
	return s.String()
}

func (s *ListImageLabelsResponse) SetHeaders(v map[string]*string) *ListImageLabelsResponse {
	s.Headers = v
	return s
}

func (s *ListImageLabelsResponse) SetStatusCode(v int32) *ListImageLabelsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListImageLabelsResponse) SetBody(v *ListImageLabelsResponseBody) *ListImageLabelsResponse {
	s.Body = v
	return s
}

type ListImagesRequest struct {
	Accessibility *string `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	Labels        *string `json:"Labels,omitempty" xml:"Labels,omitempty"`
	Name          *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Order         *string `json:"Order,omitempty" xml:"Order,omitempty"`
	PageNumber    *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize      *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ParentUserId  *string `json:"ParentUserId,omitempty" xml:"ParentUserId,omitempty"`
	Query         *string `json:"Query,omitempty" xml:"Query,omitempty"`
	SortBy        *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	UserId        *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	Verbose       *bool   `json:"Verbose,omitempty" xml:"Verbose,omitempty"`
	WorkspaceId   *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListImagesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListImagesRequest) GoString() string {
	return s.String()
}

func (s *ListImagesRequest) SetAccessibility(v string) *ListImagesRequest {
	s.Accessibility = &v
	return s
}

func (s *ListImagesRequest) SetLabels(v string) *ListImagesRequest {
	s.Labels = &v
	return s
}

func (s *ListImagesRequest) SetName(v string) *ListImagesRequest {
	s.Name = &v
	return s
}

func (s *ListImagesRequest) SetOrder(v string) *ListImagesRequest {
	s.Order = &v
	return s
}

func (s *ListImagesRequest) SetPageNumber(v int32) *ListImagesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListImagesRequest) SetPageSize(v int32) *ListImagesRequest {
	s.PageSize = &v
	return s
}

func (s *ListImagesRequest) SetParentUserId(v string) *ListImagesRequest {
	s.ParentUserId = &v
	return s
}

func (s *ListImagesRequest) SetQuery(v string) *ListImagesRequest {
	s.Query = &v
	return s
}

func (s *ListImagesRequest) SetSortBy(v string) *ListImagesRequest {
	s.SortBy = &v
	return s
}

func (s *ListImagesRequest) SetUserId(v string) *ListImagesRequest {
	s.UserId = &v
	return s
}

func (s *ListImagesRequest) SetVerbose(v bool) *ListImagesRequest {
	s.Verbose = &v
	return s
}

func (s *ListImagesRequest) SetWorkspaceId(v string) *ListImagesRequest {
	s.WorkspaceId = &v
	return s
}

type ListImagesResponseBody struct {
	Images     []*ListImagesResponseBodyImages `json:"Images,omitempty" xml:"Images,omitempty" type:"Repeated"`
	RequestId  *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int64                          `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListImagesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListImagesResponseBody) GoString() string {
	return s.String()
}

func (s *ListImagesResponseBody) SetImages(v []*ListImagesResponseBodyImages) *ListImagesResponseBody {
	s.Images = v
	return s
}

func (s *ListImagesResponseBody) SetRequestId(v string) *ListImagesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListImagesResponseBody) SetTotalCount(v int64) *ListImagesResponseBody {
	s.TotalCount = &v
	return s
}

type ListImagesResponseBodyImages struct {
	Accessibility   *string                               `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	Description     *string                               `json:"Description,omitempty" xml:"Description,omitempty"`
	GmtCreateTime   *string                               `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	GmtModifiedTime *string                               `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	ImageId         *string                               `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	ImageUri        *string                               `json:"ImageUri,omitempty" xml:"ImageUri,omitempty"`
	Labels          []*ListImagesResponseBodyImagesLabels `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	Name            *string                               `json:"Name,omitempty" xml:"Name,omitempty"`
	ParentUserId    *string                               `json:"ParentUserId,omitempty" xml:"ParentUserId,omitempty"`
	Size            *int32                                `json:"Size,omitempty" xml:"Size,omitempty"`
	UserId          *string                               `json:"UserId,omitempty" xml:"UserId,omitempty"`
	WorkspaceId     *string                               `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListImagesResponseBodyImages) String() string {
	return tea.Prettify(s)
}

func (s ListImagesResponseBodyImages) GoString() string {
	return s.String()
}

func (s *ListImagesResponseBodyImages) SetAccessibility(v string) *ListImagesResponseBodyImages {
	s.Accessibility = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetDescription(v string) *ListImagesResponseBodyImages {
	s.Description = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetGmtCreateTime(v string) *ListImagesResponseBodyImages {
	s.GmtCreateTime = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetGmtModifiedTime(v string) *ListImagesResponseBodyImages {
	s.GmtModifiedTime = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetImageId(v string) *ListImagesResponseBodyImages {
	s.ImageId = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetImageUri(v string) *ListImagesResponseBodyImages {
	s.ImageUri = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetLabels(v []*ListImagesResponseBodyImagesLabels) *ListImagesResponseBodyImages {
	s.Labels = v
	return s
}

func (s *ListImagesResponseBodyImages) SetName(v string) *ListImagesResponseBodyImages {
	s.Name = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetParentUserId(v string) *ListImagesResponseBodyImages {
	s.ParentUserId = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetSize(v int32) *ListImagesResponseBodyImages {
	s.Size = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetUserId(v string) *ListImagesResponseBodyImages {
	s.UserId = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetWorkspaceId(v string) *ListImagesResponseBodyImages {
	s.WorkspaceId = &v
	return s
}

type ListImagesResponseBodyImagesLabels struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListImagesResponseBodyImagesLabels) String() string {
	return tea.Prettify(s)
}

func (s ListImagesResponseBodyImagesLabels) GoString() string {
	return s.String()
}

func (s *ListImagesResponseBodyImagesLabels) SetKey(v string) *ListImagesResponseBodyImagesLabels {
	s.Key = &v
	return s
}

func (s *ListImagesResponseBodyImagesLabels) SetValue(v string) *ListImagesResponseBodyImagesLabels {
	s.Value = &v
	return s
}

type ListImagesResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListImagesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListImagesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListImagesResponse) GoString() string {
	return s.String()
}

func (s *ListImagesResponse) SetHeaders(v map[string]*string) *ListImagesResponse {
	s.Headers = v
	return s
}

func (s *ListImagesResponse) SetStatusCode(v int32) *ListImagesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListImagesResponse) SetBody(v *ListImagesResponseBody) *ListImagesResponse {
	s.Body = v
	return s
}

type ListMembersRequest struct {
	MemberName *string `json:"MemberName,omitempty" xml:"MemberName,omitempty"`
	PageNumber *int64  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Roles      *string `json:"Roles,omitempty" xml:"Roles,omitempty"`
}

func (s ListMembersRequest) String() string {
	return tea.Prettify(s)
}

func (s ListMembersRequest) GoString() string {
	return s.String()
}

func (s *ListMembersRequest) SetMemberName(v string) *ListMembersRequest {
	s.MemberName = &v
	return s
}

func (s *ListMembersRequest) SetPageNumber(v int64) *ListMembersRequest {
	s.PageNumber = &v
	return s
}

func (s *ListMembersRequest) SetPageSize(v int32) *ListMembersRequest {
	s.PageSize = &v
	return s
}

func (s *ListMembersRequest) SetRoles(v string) *ListMembersRequest {
	s.Roles = &v
	return s
}

type ListMembersResponseBody struct {
	Members    []*ListMembersResponseBodyMembers `json:"Members,omitempty" xml:"Members,omitempty" type:"Repeated"`
	RequestId  *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int64                            `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListMembersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListMembersResponseBody) GoString() string {
	return s.String()
}

func (s *ListMembersResponseBody) SetMembers(v []*ListMembersResponseBodyMembers) *ListMembersResponseBody {
	s.Members = v
	return s
}

func (s *ListMembersResponseBody) SetRequestId(v string) *ListMembersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMembersResponseBody) SetTotalCount(v int64) *ListMembersResponseBody {
	s.TotalCount = &v
	return s
}

type ListMembersResponseBodyMembers struct {
	DisplayName   *string   `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	GmtCreateTime *string   `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	MemberId      *string   `json:"MemberId,omitempty" xml:"MemberId,omitempty"`
	MemberName    *string   `json:"MemberName,omitempty" xml:"MemberName,omitempty"`
	Roles         []*string `json:"Roles,omitempty" xml:"Roles,omitempty" type:"Repeated"`
	UserId        *string   `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListMembersResponseBodyMembers) String() string {
	return tea.Prettify(s)
}

func (s ListMembersResponseBodyMembers) GoString() string {
	return s.String()
}

func (s *ListMembersResponseBodyMembers) SetDisplayName(v string) *ListMembersResponseBodyMembers {
	s.DisplayName = &v
	return s
}

func (s *ListMembersResponseBodyMembers) SetGmtCreateTime(v string) *ListMembersResponseBodyMembers {
	s.GmtCreateTime = &v
	return s
}

func (s *ListMembersResponseBodyMembers) SetMemberId(v string) *ListMembersResponseBodyMembers {
	s.MemberId = &v
	return s
}

func (s *ListMembersResponseBodyMembers) SetMemberName(v string) *ListMembersResponseBodyMembers {
	s.MemberName = &v
	return s
}

func (s *ListMembersResponseBodyMembers) SetRoles(v []*string) *ListMembersResponseBodyMembers {
	s.Roles = v
	return s
}

func (s *ListMembersResponseBodyMembers) SetUserId(v string) *ListMembersResponseBodyMembers {
	s.UserId = &v
	return s
}

type ListMembersResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMembersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMembersResponse) String() string {
	return tea.Prettify(s)
}

func (s ListMembersResponse) GoString() string {
	return s.String()
}

func (s *ListMembersResponse) SetHeaders(v map[string]*string) *ListMembersResponse {
	s.Headers = v
	return s
}

func (s *ListMembersResponse) SetStatusCode(v int32) *ListMembersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMembersResponse) SetBody(v *ListMembersResponseBody) *ListMembersResponse {
	s.Body = v
	return s
}

type ListModelVersionsRequest struct {
	ApprovalStatus *string `json:"ApprovalStatus,omitempty" xml:"ApprovalStatus,omitempty"`
	FormatType     *string `json:"FormatType,omitempty" xml:"FormatType,omitempty"`
	FrameworkType  *string `json:"FrameworkType,omitempty" xml:"FrameworkType,omitempty"`
	Label          *string `json:"Label,omitempty" xml:"Label,omitempty"`
	Order          *string `json:"Order,omitempty" xml:"Order,omitempty"`
	PageNumber     *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize       *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SortBy         *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	SourceId       *string `json:"SourceId,omitempty" xml:"SourceId,omitempty"`
	SourceType     *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	VersionName    *string `json:"VersionName,omitempty" xml:"VersionName,omitempty"`
}

func (s ListModelVersionsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListModelVersionsRequest) GoString() string {
	return s.String()
}

func (s *ListModelVersionsRequest) SetApprovalStatus(v string) *ListModelVersionsRequest {
	s.ApprovalStatus = &v
	return s
}

func (s *ListModelVersionsRequest) SetFormatType(v string) *ListModelVersionsRequest {
	s.FormatType = &v
	return s
}

func (s *ListModelVersionsRequest) SetFrameworkType(v string) *ListModelVersionsRequest {
	s.FrameworkType = &v
	return s
}

func (s *ListModelVersionsRequest) SetLabel(v string) *ListModelVersionsRequest {
	s.Label = &v
	return s
}

func (s *ListModelVersionsRequest) SetOrder(v string) *ListModelVersionsRequest {
	s.Order = &v
	return s
}

func (s *ListModelVersionsRequest) SetPageNumber(v int32) *ListModelVersionsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListModelVersionsRequest) SetPageSize(v int32) *ListModelVersionsRequest {
	s.PageSize = &v
	return s
}

func (s *ListModelVersionsRequest) SetSortBy(v string) *ListModelVersionsRequest {
	s.SortBy = &v
	return s
}

func (s *ListModelVersionsRequest) SetSourceId(v string) *ListModelVersionsRequest {
	s.SourceId = &v
	return s
}

func (s *ListModelVersionsRequest) SetSourceType(v string) *ListModelVersionsRequest {
	s.SourceType = &v
	return s
}

func (s *ListModelVersionsRequest) SetVersionName(v string) *ListModelVersionsRequest {
	s.VersionName = &v
	return s
}

type ListModelVersionsResponseBody struct {
	RequestId  *string         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int64          `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	Versions   []*ModelVersion `json:"Versions,omitempty" xml:"Versions,omitempty" type:"Repeated"`
}

func (s ListModelVersionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListModelVersionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListModelVersionsResponseBody) SetRequestId(v string) *ListModelVersionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListModelVersionsResponseBody) SetTotalCount(v int64) *ListModelVersionsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListModelVersionsResponseBody) SetVersions(v []*ModelVersion) *ListModelVersionsResponseBody {
	s.Versions = v
	return s
}

type ListModelVersionsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListModelVersionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListModelVersionsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListModelVersionsResponse) GoString() string {
	return s.String()
}

func (s *ListModelVersionsResponse) SetHeaders(v map[string]*string) *ListModelVersionsResponse {
	s.Headers = v
	return s
}

func (s *ListModelVersionsResponse) SetStatusCode(v int32) *ListModelVersionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListModelVersionsResponse) SetBody(v *ListModelVersionsResponseBody) *ListModelVersionsResponse {
	s.Body = v
	return s
}

type ListModelsRequest struct {
	Collections *string `json:"Collections,omitempty" xml:"Collections,omitempty"`
	Domain      *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	Label       *string `json:"Label,omitempty" xml:"Label,omitempty"`
	ModelName   *string `json:"ModelName,omitempty" xml:"ModelName,omitempty"`
	ModelType   *string `json:"ModelType,omitempty" xml:"ModelType,omitempty"`
	Order       *string `json:"Order,omitempty" xml:"Order,omitempty"`
	Origin      *string `json:"Origin,omitempty" xml:"Origin,omitempty"`
	PageNumber  *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Provider    *string `json:"Provider,omitempty" xml:"Provider,omitempty"`
	Query       *string `json:"Query,omitempty" xml:"Query,omitempty"`
	SortBy      *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	Task        *string `json:"Task,omitempty" xml:"Task,omitempty"`
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListModelsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListModelsRequest) GoString() string {
	return s.String()
}

func (s *ListModelsRequest) SetCollections(v string) *ListModelsRequest {
	s.Collections = &v
	return s
}

func (s *ListModelsRequest) SetDomain(v string) *ListModelsRequest {
	s.Domain = &v
	return s
}

func (s *ListModelsRequest) SetLabel(v string) *ListModelsRequest {
	s.Label = &v
	return s
}

func (s *ListModelsRequest) SetModelName(v string) *ListModelsRequest {
	s.ModelName = &v
	return s
}

func (s *ListModelsRequest) SetModelType(v string) *ListModelsRequest {
	s.ModelType = &v
	return s
}

func (s *ListModelsRequest) SetOrder(v string) *ListModelsRequest {
	s.Order = &v
	return s
}

func (s *ListModelsRequest) SetOrigin(v string) *ListModelsRequest {
	s.Origin = &v
	return s
}

func (s *ListModelsRequest) SetPageNumber(v int32) *ListModelsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListModelsRequest) SetPageSize(v int32) *ListModelsRequest {
	s.PageSize = &v
	return s
}

func (s *ListModelsRequest) SetProvider(v string) *ListModelsRequest {
	s.Provider = &v
	return s
}

func (s *ListModelsRequest) SetQuery(v string) *ListModelsRequest {
	s.Query = &v
	return s
}

func (s *ListModelsRequest) SetSortBy(v string) *ListModelsRequest {
	s.SortBy = &v
	return s
}

func (s *ListModelsRequest) SetTask(v string) *ListModelsRequest {
	s.Task = &v
	return s
}

func (s *ListModelsRequest) SetWorkspaceId(v string) *ListModelsRequest {
	s.WorkspaceId = &v
	return s
}

type ListModelsResponseBody struct {
	Models     []*Model `json:"Models,omitempty" xml:"Models,omitempty" type:"Repeated"`
	RequestId  *string  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int64   `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListModelsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListModelsResponseBody) GoString() string {
	return s.String()
}

func (s *ListModelsResponseBody) SetModels(v []*Model) *ListModelsResponseBody {
	s.Models = v
	return s
}

func (s *ListModelsResponseBody) SetRequestId(v string) *ListModelsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListModelsResponseBody) SetTotalCount(v int64) *ListModelsResponseBody {
	s.TotalCount = &v
	return s
}

type ListModelsResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListModelsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListModelsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListModelsResponse) GoString() string {
	return s.String()
}

func (s *ListModelsResponse) SetHeaders(v map[string]*string) *ListModelsResponse {
	s.Headers = v
	return s
}

func (s *ListModelsResponse) SetStatusCode(v int32) *ListModelsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListModelsResponse) SetBody(v *ListModelsResponseBody) *ListModelsResponse {
	s.Body = v
	return s
}

type ListPermissionsResponseBody struct {
	Permissions []*ListPermissionsResponseBodyPermissions `json:"Permissions,omitempty" xml:"Permissions,omitempty" type:"Repeated"`
	RequestId   *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount  *int64                                    `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListPermissionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListPermissionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListPermissionsResponseBody) SetPermissions(v []*ListPermissionsResponseBodyPermissions) *ListPermissionsResponseBody {
	s.Permissions = v
	return s
}

func (s *ListPermissionsResponseBody) SetRequestId(v string) *ListPermissionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPermissionsResponseBody) SetTotalCount(v int64) *ListPermissionsResponseBody {
	s.TotalCount = &v
	return s
}

type ListPermissionsResponseBodyPermissions struct {
	PermissionCode  *string                                                  `json:"PermissionCode,omitempty" xml:"PermissionCode,omitempty"`
	PermissionRules []*ListPermissionsResponseBodyPermissionsPermissionRules `json:"PermissionRules,omitempty" xml:"PermissionRules,omitempty" type:"Repeated"`
}

func (s ListPermissionsResponseBodyPermissions) String() string {
	return tea.Prettify(s)
}

func (s ListPermissionsResponseBodyPermissions) GoString() string {
	return s.String()
}

func (s *ListPermissionsResponseBodyPermissions) SetPermissionCode(v string) *ListPermissionsResponseBodyPermissions {
	s.PermissionCode = &v
	return s
}

func (s *ListPermissionsResponseBodyPermissions) SetPermissionRules(v []*ListPermissionsResponseBodyPermissionsPermissionRules) *ListPermissionsResponseBodyPermissions {
	s.PermissionRules = v
	return s
}

type ListPermissionsResponseBodyPermissionsPermissionRules struct {
	Accessibility    *string `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	EntityAccessType *string `json:"EntityAccessType,omitempty" xml:"EntityAccessType,omitempty"`
}

func (s ListPermissionsResponseBodyPermissionsPermissionRules) String() string {
	return tea.Prettify(s)
}

func (s ListPermissionsResponseBodyPermissionsPermissionRules) GoString() string {
	return s.String()
}

func (s *ListPermissionsResponseBodyPermissionsPermissionRules) SetAccessibility(v string) *ListPermissionsResponseBodyPermissionsPermissionRules {
	s.Accessibility = &v
	return s
}

func (s *ListPermissionsResponseBodyPermissionsPermissionRules) SetEntityAccessType(v string) *ListPermissionsResponseBodyPermissionsPermissionRules {
	s.EntityAccessType = &v
	return s
}

type ListPermissionsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPermissionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPermissionsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListPermissionsResponse) GoString() string {
	return s.String()
}

func (s *ListPermissionsResponse) SetHeaders(v map[string]*string) *ListPermissionsResponse {
	s.Headers = v
	return s
}

func (s *ListPermissionsResponse) SetStatusCode(v int32) *ListPermissionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPermissionsResponse) SetBody(v *ListPermissionsResponseBody) *ListPermissionsResponse {
	s.Body = v
	return s
}

type ListProductsRequest struct {
	ProductCodes *string `json:"ProductCodes,omitempty" xml:"ProductCodes,omitempty"`
	ServiceCodes *string `json:"ServiceCodes,omitempty" xml:"ServiceCodes,omitempty"`
	Verbose      *bool   `json:"Verbose,omitempty" xml:"Verbose,omitempty"`
}

func (s ListProductsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListProductsRequest) GoString() string {
	return s.String()
}

func (s *ListProductsRequest) SetProductCodes(v string) *ListProductsRequest {
	s.ProductCodes = &v
	return s
}

func (s *ListProductsRequest) SetServiceCodes(v string) *ListProductsRequest {
	s.ServiceCodes = &v
	return s
}

func (s *ListProductsRequest) SetVerbose(v bool) *ListProductsRequest {
	s.Verbose = &v
	return s
}

type ListProductsResponseBody struct {
	Products  []*ListProductsResponseBodyProducts `json:"Products,omitempty" xml:"Products,omitempty" type:"Repeated"`
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Services  []*ListProductsResponseBodyServices `json:"Services,omitempty" xml:"Services,omitempty" type:"Repeated"`
}

func (s ListProductsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListProductsResponseBody) GoString() string {
	return s.String()
}

func (s *ListProductsResponseBody) SetProducts(v []*ListProductsResponseBodyProducts) *ListProductsResponseBody {
	s.Products = v
	return s
}

func (s *ListProductsResponseBody) SetRequestId(v string) *ListProductsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListProductsResponseBody) SetServices(v []*ListProductsResponseBodyServices) *ListProductsResponseBody {
	s.Services = v
	return s
}

type ListProductsResponseBodyProducts struct {
	HasPermissionToPurchase *bool   `json:"HasPermissionToPurchase,omitempty" xml:"HasPermissionToPurchase,omitempty"`
	IsPurchased             *bool   `json:"IsPurchased,omitempty" xml:"IsPurchased,omitempty"`
	ProductCode             *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	ProductInstanceId       *string `json:"ProductInstanceId,omitempty" xml:"ProductInstanceId,omitempty"`
	PurchaseUrl             *string `json:"PurchaseUrl,omitempty" xml:"PurchaseUrl,omitempty"`
}

func (s ListProductsResponseBodyProducts) String() string {
	return tea.Prettify(s)
}

func (s ListProductsResponseBodyProducts) GoString() string {
	return s.String()
}

func (s *ListProductsResponseBodyProducts) SetHasPermissionToPurchase(v bool) *ListProductsResponseBodyProducts {
	s.HasPermissionToPurchase = &v
	return s
}

func (s *ListProductsResponseBodyProducts) SetIsPurchased(v bool) *ListProductsResponseBodyProducts {
	s.IsPurchased = &v
	return s
}

func (s *ListProductsResponseBodyProducts) SetProductCode(v string) *ListProductsResponseBodyProducts {
	s.ProductCode = &v
	return s
}

func (s *ListProductsResponseBodyProducts) SetProductInstanceId(v string) *ListProductsResponseBodyProducts {
	s.ProductInstanceId = &v
	return s
}

func (s *ListProductsResponseBodyProducts) SetPurchaseUrl(v string) *ListProductsResponseBodyProducts {
	s.PurchaseUrl = &v
	return s
}

type ListProductsResponseBodyServices struct {
	IsOpen      *bool   `json:"IsOpen,omitempty" xml:"IsOpen,omitempty"`
	OpenUrl     *string `json:"OpenUrl,omitempty" xml:"OpenUrl,omitempty"`
	ServiceCode *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
}

func (s ListProductsResponseBodyServices) String() string {
	return tea.Prettify(s)
}

func (s ListProductsResponseBodyServices) GoString() string {
	return s.String()
}

func (s *ListProductsResponseBodyServices) SetIsOpen(v bool) *ListProductsResponseBodyServices {
	s.IsOpen = &v
	return s
}

func (s *ListProductsResponseBodyServices) SetOpenUrl(v string) *ListProductsResponseBodyServices {
	s.OpenUrl = &v
	return s
}

func (s *ListProductsResponseBodyServices) SetServiceCode(v string) *ListProductsResponseBodyServices {
	s.ServiceCode = &v
	return s
}

type ListProductsResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListProductsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListProductsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListProductsResponse) GoString() string {
	return s.String()
}

func (s *ListProductsResponse) SetHeaders(v map[string]*string) *ListProductsResponse {
	s.Headers = v
	return s
}

func (s *ListProductsResponse) SetStatusCode(v int32) *ListProductsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListProductsResponse) SetBody(v *ListProductsResponseBody) *ListProductsResponse {
	s.Body = v
	return s
}

type ListQuotasRequest struct {
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListQuotasRequest) String() string {
	return tea.Prettify(s)
}

func (s ListQuotasRequest) GoString() string {
	return s.String()
}

func (s *ListQuotasRequest) SetName(v string) *ListQuotasRequest {
	s.Name = &v
	return s
}

type ListQuotasResponseBody struct {
	Quotas     []*ListQuotasResponseBodyQuotas `json:"Quotas,omitempty" xml:"Quotas,omitempty" type:"Repeated"`
	RequestId  *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int64                          `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListQuotasResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListQuotasResponseBody) GoString() string {
	return s.String()
}

func (s *ListQuotasResponseBody) SetQuotas(v []*ListQuotasResponseBodyQuotas) *ListQuotasResponseBody {
	s.Quotas = v
	return s
}

func (s *ListQuotasResponseBody) SetRequestId(v string) *ListQuotasResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListQuotasResponseBody) SetTotalCount(v int64) *ListQuotasResponseBody {
	s.TotalCount = &v
	return s
}

type ListQuotasResponseBodyQuotas struct {
	DisplayName *string                              `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	Id          *string                              `json:"Id,omitempty" xml:"Id,omitempty"`
	Mode        *string                              `json:"Mode,omitempty" xml:"Mode,omitempty"`
	Name        *string                              `json:"Name,omitempty" xml:"Name,omitempty"`
	ProductCode *string                              `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	QuotaType   *string                              `json:"QuotaType,omitempty" xml:"QuotaType,omitempty"`
	Specs       []*ListQuotasResponseBodyQuotasSpecs `json:"Specs,omitempty" xml:"Specs,omitempty" type:"Repeated"`
}

func (s ListQuotasResponseBodyQuotas) String() string {
	return tea.Prettify(s)
}

func (s ListQuotasResponseBodyQuotas) GoString() string {
	return s.String()
}

func (s *ListQuotasResponseBodyQuotas) SetDisplayName(v string) *ListQuotasResponseBodyQuotas {
	s.DisplayName = &v
	return s
}

func (s *ListQuotasResponseBodyQuotas) SetId(v string) *ListQuotasResponseBodyQuotas {
	s.Id = &v
	return s
}

func (s *ListQuotasResponseBodyQuotas) SetMode(v string) *ListQuotasResponseBodyQuotas {
	s.Mode = &v
	return s
}

func (s *ListQuotasResponseBodyQuotas) SetName(v string) *ListQuotasResponseBodyQuotas {
	s.Name = &v
	return s
}

func (s *ListQuotasResponseBodyQuotas) SetProductCode(v string) *ListQuotasResponseBodyQuotas {
	s.ProductCode = &v
	return s
}

func (s *ListQuotasResponseBodyQuotas) SetQuotaType(v string) *ListQuotasResponseBodyQuotas {
	s.QuotaType = &v
	return s
}

func (s *ListQuotasResponseBodyQuotas) SetSpecs(v []*ListQuotasResponseBodyQuotasSpecs) *ListQuotasResponseBodyQuotas {
	s.Specs = v
	return s
}

type ListQuotasResponseBodyQuotasSpecs struct {
	Name  *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Type  *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListQuotasResponseBodyQuotasSpecs) String() string {
	return tea.Prettify(s)
}

func (s ListQuotasResponseBodyQuotasSpecs) GoString() string {
	return s.String()
}

func (s *ListQuotasResponseBodyQuotasSpecs) SetName(v string) *ListQuotasResponseBodyQuotasSpecs {
	s.Name = &v
	return s
}

func (s *ListQuotasResponseBodyQuotasSpecs) SetType(v string) *ListQuotasResponseBodyQuotasSpecs {
	s.Type = &v
	return s
}

func (s *ListQuotasResponseBodyQuotasSpecs) SetValue(v string) *ListQuotasResponseBodyQuotasSpecs {
	s.Value = &v
	return s
}

type ListQuotasResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListQuotasResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListQuotasResponse) String() string {
	return tea.Prettify(s)
}

func (s ListQuotasResponse) GoString() string {
	return s.String()
}

func (s *ListQuotasResponse) SetHeaders(v map[string]*string) *ListQuotasResponse {
	s.Headers = v
	return s
}

func (s *ListQuotasResponse) SetStatusCode(v int32) *ListQuotasResponse {
	s.StatusCode = &v
	return s
}

func (s *ListQuotasResponse) SetBody(v *ListQuotasResponseBody) *ListQuotasResponse {
	s.Body = v
	return s
}

type ListResourcesRequest struct {
	GroupName     *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	Labels        *string `json:"Labels,omitempty" xml:"Labels,omitempty"`
	Option        *string `json:"Option,omitempty" xml:"Option,omitempty"`
	PageNumber    *int64  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize      *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ProductTypes  *string `json:"ProductTypes,omitempty" xml:"ProductTypes,omitempty"`
	QuotaIds      *string `json:"QuotaIds,omitempty" xml:"QuotaIds,omitempty"`
	ResourceName  *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	ResourceTypes *string `json:"ResourceTypes,omitempty" xml:"ResourceTypes,omitempty"`
	Verbose       *bool   `json:"Verbose,omitempty" xml:"Verbose,omitempty"`
	VerboseFields *string `json:"VerboseFields,omitempty" xml:"VerboseFields,omitempty"`
	WorkspaceId   *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListResourcesRequest) GoString() string {
	return s.String()
}

func (s *ListResourcesRequest) SetGroupName(v string) *ListResourcesRequest {
	s.GroupName = &v
	return s
}

func (s *ListResourcesRequest) SetLabels(v string) *ListResourcesRequest {
	s.Labels = &v
	return s
}

func (s *ListResourcesRequest) SetOption(v string) *ListResourcesRequest {
	s.Option = &v
	return s
}

func (s *ListResourcesRequest) SetPageNumber(v int64) *ListResourcesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListResourcesRequest) SetPageSize(v int32) *ListResourcesRequest {
	s.PageSize = &v
	return s
}

func (s *ListResourcesRequest) SetProductTypes(v string) *ListResourcesRequest {
	s.ProductTypes = &v
	return s
}

func (s *ListResourcesRequest) SetQuotaIds(v string) *ListResourcesRequest {
	s.QuotaIds = &v
	return s
}

func (s *ListResourcesRequest) SetResourceName(v string) *ListResourcesRequest {
	s.ResourceName = &v
	return s
}

func (s *ListResourcesRequest) SetResourceTypes(v string) *ListResourcesRequest {
	s.ResourceTypes = &v
	return s
}

func (s *ListResourcesRequest) SetVerbose(v bool) *ListResourcesRequest {
	s.Verbose = &v
	return s
}

func (s *ListResourcesRequest) SetVerboseFields(v string) *ListResourcesRequest {
	s.VerboseFields = &v
	return s
}

func (s *ListResourcesRequest) SetWorkspaceId(v string) *ListResourcesRequest {
	s.WorkspaceId = &v
	return s
}

type ListResourcesResponseBody struct {
	RequestId  *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Resources  []*ListResourcesResponseBodyResources `json:"Resources,omitempty" xml:"Resources,omitempty" type:"Repeated"`
	TotalCount *int64                                `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ListResourcesResponseBody) SetRequestId(v string) *ListResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListResourcesResponseBody) SetResources(v []*ListResourcesResponseBodyResources) *ListResourcesResponseBody {
	s.Resources = v
	return s
}

func (s *ListResourcesResponseBody) SetTotalCount(v int64) *ListResourcesResponseBody {
	s.TotalCount = &v
	return s
}

type ListResourcesResponseBodyResources struct {
	Encryption    *ListResourcesResponseBodyResourcesEncryption `json:"Encryption,omitempty" xml:"Encryption,omitempty" type:"Struct"`
	EnvType       *string                                       `json:"EnvType,omitempty" xml:"EnvType,omitempty"`
	Executor      *ListResourcesResponseBodyResourcesExecutor   `json:"Executor,omitempty" xml:"Executor,omitempty" type:"Struct"`
	GmtCreateTime *string                                       `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	GroupName     *string                                       `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	Id            *string                                       `json:"Id,omitempty" xml:"Id,omitempty"`
	IsDefault     *bool                                         `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	Labels        []*ListResourcesResponseBodyResourcesLabels   `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	Name          *string                                       `json:"Name,omitempty" xml:"Name,omitempty"`
	ProductType   *string                                       `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	Quotas        []*ListResourcesResponseBodyResourcesQuotas   `json:"Quotas,omitempty" xml:"Quotas,omitempty" type:"Repeated"`
	ResourceType  *string                                       `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Spec          map[string]interface{}                        `json:"Spec,omitempty" xml:"Spec,omitempty"`
	WorkspaceId   *string                                       `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListResourcesResponseBodyResources) String() string {
	return tea.Prettify(s)
}

func (s ListResourcesResponseBodyResources) GoString() string {
	return s.String()
}

func (s *ListResourcesResponseBodyResources) SetEncryption(v *ListResourcesResponseBodyResourcesEncryption) *ListResourcesResponseBodyResources {
	s.Encryption = v
	return s
}

func (s *ListResourcesResponseBodyResources) SetEnvType(v string) *ListResourcesResponseBodyResources {
	s.EnvType = &v
	return s
}

func (s *ListResourcesResponseBodyResources) SetExecutor(v *ListResourcesResponseBodyResourcesExecutor) *ListResourcesResponseBodyResources {
	s.Executor = v
	return s
}

func (s *ListResourcesResponseBodyResources) SetGmtCreateTime(v string) *ListResourcesResponseBodyResources {
	s.GmtCreateTime = &v
	return s
}

func (s *ListResourcesResponseBodyResources) SetGroupName(v string) *ListResourcesResponseBodyResources {
	s.GroupName = &v
	return s
}

func (s *ListResourcesResponseBodyResources) SetId(v string) *ListResourcesResponseBodyResources {
	s.Id = &v
	return s
}

func (s *ListResourcesResponseBodyResources) SetIsDefault(v bool) *ListResourcesResponseBodyResources {
	s.IsDefault = &v
	return s
}

func (s *ListResourcesResponseBodyResources) SetLabels(v []*ListResourcesResponseBodyResourcesLabels) *ListResourcesResponseBodyResources {
	s.Labels = v
	return s
}

func (s *ListResourcesResponseBodyResources) SetName(v string) *ListResourcesResponseBodyResources {
	s.Name = &v
	return s
}

func (s *ListResourcesResponseBodyResources) SetProductType(v string) *ListResourcesResponseBodyResources {
	s.ProductType = &v
	return s
}

func (s *ListResourcesResponseBodyResources) SetQuotas(v []*ListResourcesResponseBodyResourcesQuotas) *ListResourcesResponseBodyResources {
	s.Quotas = v
	return s
}

func (s *ListResourcesResponseBodyResources) SetResourceType(v string) *ListResourcesResponseBodyResources {
	s.ResourceType = &v
	return s
}

func (s *ListResourcesResponseBodyResources) SetSpec(v map[string]interface{}) *ListResourcesResponseBodyResources {
	s.Spec = v
	return s
}

func (s *ListResourcesResponseBodyResources) SetWorkspaceId(v string) *ListResourcesResponseBodyResources {
	s.WorkspaceId = &v
	return s
}

type ListResourcesResponseBodyResourcesEncryption struct {
	Algorithm *string `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	Enabled   *bool   `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	Key       *string `json:"Key,omitempty" xml:"Key,omitempty"`
}

func (s ListResourcesResponseBodyResourcesEncryption) String() string {
	return tea.Prettify(s)
}

func (s ListResourcesResponseBodyResourcesEncryption) GoString() string {
	return s.String()
}

func (s *ListResourcesResponseBodyResourcesEncryption) SetAlgorithm(v string) *ListResourcesResponseBodyResourcesEncryption {
	s.Algorithm = &v
	return s
}

func (s *ListResourcesResponseBodyResourcesEncryption) SetEnabled(v bool) *ListResourcesResponseBodyResourcesEncryption {
	s.Enabled = &v
	return s
}

func (s *ListResourcesResponseBodyResourcesEncryption) SetKey(v string) *ListResourcesResponseBodyResourcesEncryption {
	s.Key = &v
	return s
}

type ListResourcesResponseBodyResourcesExecutor struct {
	OwnerId *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s ListResourcesResponseBodyResourcesExecutor) String() string {
	return tea.Prettify(s)
}

func (s ListResourcesResponseBodyResourcesExecutor) GoString() string {
	return s.String()
}

func (s *ListResourcesResponseBodyResourcesExecutor) SetOwnerId(v string) *ListResourcesResponseBodyResourcesExecutor {
	s.OwnerId = &v
	return s
}

type ListResourcesResponseBodyResourcesLabels struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListResourcesResponseBodyResourcesLabels) String() string {
	return tea.Prettify(s)
}

func (s ListResourcesResponseBodyResourcesLabels) GoString() string {
	return s.String()
}

func (s *ListResourcesResponseBodyResourcesLabels) SetKey(v string) *ListResourcesResponseBodyResourcesLabels {
	s.Key = &v
	return s
}

func (s *ListResourcesResponseBodyResourcesLabels) SetValue(v string) *ListResourcesResponseBodyResourcesLabels {
	s.Value = &v
	return s
}

type ListResourcesResponseBodyResourcesQuotas struct {
	CardType    *string                                          `json:"CardType,omitempty" xml:"CardType,omitempty"`
	DisplayName *string                                          `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	Id          *string                                          `json:"Id,omitempty" xml:"Id,omitempty"`
	Mode        *string                                          `json:"Mode,omitempty" xml:"Mode,omitempty"`
	Name        *string                                          `json:"Name,omitempty" xml:"Name,omitempty"`
	ProductCode *string                                          `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	QuotaType   *string                                          `json:"QuotaType,omitempty" xml:"QuotaType,omitempty"`
	Specs       []*ListResourcesResponseBodyResourcesQuotasSpecs `json:"Specs,omitempty" xml:"Specs,omitempty" type:"Repeated"`
}

func (s ListResourcesResponseBodyResourcesQuotas) String() string {
	return tea.Prettify(s)
}

func (s ListResourcesResponseBodyResourcesQuotas) GoString() string {
	return s.String()
}

func (s *ListResourcesResponseBodyResourcesQuotas) SetCardType(v string) *ListResourcesResponseBodyResourcesQuotas {
	s.CardType = &v
	return s
}

func (s *ListResourcesResponseBodyResourcesQuotas) SetDisplayName(v string) *ListResourcesResponseBodyResourcesQuotas {
	s.DisplayName = &v
	return s
}

func (s *ListResourcesResponseBodyResourcesQuotas) SetId(v string) *ListResourcesResponseBodyResourcesQuotas {
	s.Id = &v
	return s
}

func (s *ListResourcesResponseBodyResourcesQuotas) SetMode(v string) *ListResourcesResponseBodyResourcesQuotas {
	s.Mode = &v
	return s
}

func (s *ListResourcesResponseBodyResourcesQuotas) SetName(v string) *ListResourcesResponseBodyResourcesQuotas {
	s.Name = &v
	return s
}

func (s *ListResourcesResponseBodyResourcesQuotas) SetProductCode(v string) *ListResourcesResponseBodyResourcesQuotas {
	s.ProductCode = &v
	return s
}

func (s *ListResourcesResponseBodyResourcesQuotas) SetQuotaType(v string) *ListResourcesResponseBodyResourcesQuotas {
	s.QuotaType = &v
	return s
}

func (s *ListResourcesResponseBodyResourcesQuotas) SetSpecs(v []*ListResourcesResponseBodyResourcesQuotasSpecs) *ListResourcesResponseBodyResourcesQuotas {
	s.Specs = v
	return s
}

type ListResourcesResponseBodyResourcesQuotasSpecs struct {
	Name  *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListResourcesResponseBodyResourcesQuotasSpecs) String() string {
	return tea.Prettify(s)
}

func (s ListResourcesResponseBodyResourcesQuotasSpecs) GoString() string {
	return s.String()
}

func (s *ListResourcesResponseBodyResourcesQuotasSpecs) SetName(v string) *ListResourcesResponseBodyResourcesQuotasSpecs {
	s.Name = &v
	return s
}

func (s *ListResourcesResponseBodyResourcesQuotasSpecs) SetValue(v string) *ListResourcesResponseBodyResourcesQuotasSpecs {
	s.Value = &v
	return s
}

type ListResourcesResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListResourcesResponse) GoString() string {
	return s.String()
}

func (s *ListResourcesResponse) SetHeaders(v map[string]*string) *ListResourcesResponse {
	s.Headers = v
	return s
}

func (s *ListResourcesResponse) SetStatusCode(v int32) *ListResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListResourcesResponse) SetBody(v *ListResourcesResponseBody) *ListResourcesResponse {
	s.Body = v
	return s
}

type ListServiceTemplatesRequest struct {
	Label               *string `json:"Label,omitempty" xml:"Label,omitempty"`
	Order               *string `json:"Order,omitempty" xml:"Order,omitempty"`
	PageNumber          *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize            *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Provider            *string `json:"Provider,omitempty" xml:"Provider,omitempty"`
	Query               *string `json:"Query,omitempty" xml:"Query,omitempty"`
	ServiceTemplateName *string `json:"ServiceTemplateName,omitempty" xml:"ServiceTemplateName,omitempty"`
	SortBy              *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
}

func (s ListServiceTemplatesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListServiceTemplatesRequest) GoString() string {
	return s.String()
}

func (s *ListServiceTemplatesRequest) SetLabel(v string) *ListServiceTemplatesRequest {
	s.Label = &v
	return s
}

func (s *ListServiceTemplatesRequest) SetOrder(v string) *ListServiceTemplatesRequest {
	s.Order = &v
	return s
}

func (s *ListServiceTemplatesRequest) SetPageNumber(v int32) *ListServiceTemplatesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListServiceTemplatesRequest) SetPageSize(v int32) *ListServiceTemplatesRequest {
	s.PageSize = &v
	return s
}

func (s *ListServiceTemplatesRequest) SetProvider(v string) *ListServiceTemplatesRequest {
	s.Provider = &v
	return s
}

func (s *ListServiceTemplatesRequest) SetQuery(v string) *ListServiceTemplatesRequest {
	s.Query = &v
	return s
}

func (s *ListServiceTemplatesRequest) SetServiceTemplateName(v string) *ListServiceTemplatesRequest {
	s.ServiceTemplateName = &v
	return s
}

func (s *ListServiceTemplatesRequest) SetSortBy(v string) *ListServiceTemplatesRequest {
	s.SortBy = &v
	return s
}

type ListServiceTemplatesResponseBody struct {
	RequestId        *string            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ServiceTemplates []*ServiceTemplate `json:"ServiceTemplates,omitempty" xml:"ServiceTemplates,omitempty" type:"Repeated"`
	TotalCount       *int64             `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListServiceTemplatesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListServiceTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *ListServiceTemplatesResponseBody) SetRequestId(v string) *ListServiceTemplatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListServiceTemplatesResponseBody) SetServiceTemplates(v []*ServiceTemplate) *ListServiceTemplatesResponseBody {
	s.ServiceTemplates = v
	return s
}

func (s *ListServiceTemplatesResponseBody) SetTotalCount(v int64) *ListServiceTemplatesResponseBody {
	s.TotalCount = &v
	return s
}

type ListServiceTemplatesResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListServiceTemplatesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListServiceTemplatesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListServiceTemplatesResponse) GoString() string {
	return s.String()
}

func (s *ListServiceTemplatesResponse) SetHeaders(v map[string]*string) *ListServiceTemplatesResponse {
	s.Headers = v
	return s
}

func (s *ListServiceTemplatesResponse) SetStatusCode(v int32) *ListServiceTemplatesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListServiceTemplatesResponse) SetBody(v *ListServiceTemplatesResponseBody) *ListServiceTemplatesResponse {
	s.Body = v
	return s
}

type ListWorkspaceUsersRequest struct {
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s ListWorkspaceUsersRequest) String() string {
	return tea.Prettify(s)
}

func (s ListWorkspaceUsersRequest) GoString() string {
	return s.String()
}

func (s *ListWorkspaceUsersRequest) SetUserName(v string) *ListWorkspaceUsersRequest {
	s.UserName = &v
	return s
}

type ListWorkspaceUsersResponseBody struct {
	RequestId  *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int64                                 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	Users      []*ListWorkspaceUsersResponseBodyUsers `json:"Users,omitempty" xml:"Users,omitempty" type:"Repeated"`
}

func (s ListWorkspaceUsersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListWorkspaceUsersResponseBody) GoString() string {
	return s.String()
}

func (s *ListWorkspaceUsersResponseBody) SetRequestId(v string) *ListWorkspaceUsersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListWorkspaceUsersResponseBody) SetTotalCount(v int64) *ListWorkspaceUsersResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListWorkspaceUsersResponseBody) SetUsers(v []*ListWorkspaceUsersResponseBodyUsers) *ListWorkspaceUsersResponseBody {
	s.Users = v
	return s
}

type ListWorkspaceUsersResponseBodyUsers struct {
	UserId   *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s ListWorkspaceUsersResponseBodyUsers) String() string {
	return tea.Prettify(s)
}

func (s ListWorkspaceUsersResponseBodyUsers) GoString() string {
	return s.String()
}

func (s *ListWorkspaceUsersResponseBodyUsers) SetUserId(v string) *ListWorkspaceUsersResponseBodyUsers {
	s.UserId = &v
	return s
}

func (s *ListWorkspaceUsersResponseBodyUsers) SetUserName(v string) *ListWorkspaceUsersResponseBodyUsers {
	s.UserName = &v
	return s
}

type ListWorkspaceUsersResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListWorkspaceUsersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListWorkspaceUsersResponse) String() string {
	return tea.Prettify(s)
}

func (s ListWorkspaceUsersResponse) GoString() string {
	return s.String()
}

func (s *ListWorkspaceUsersResponse) SetHeaders(v map[string]*string) *ListWorkspaceUsersResponse {
	s.Headers = v
	return s
}

func (s *ListWorkspaceUsersResponse) SetStatusCode(v int32) *ListWorkspaceUsersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListWorkspaceUsersResponse) SetBody(v *ListWorkspaceUsersResponseBody) *ListWorkspaceUsersResponse {
	s.Body = v
	return s
}

type ListWorkspacesRequest struct {
	Fields        *string `json:"Fields,omitempty" xml:"Fields,omitempty"`
	ModuleList    *string `json:"ModuleList,omitempty" xml:"ModuleList,omitempty"`
	Option        *string `json:"Option,omitempty" xml:"Option,omitempty"`
	Order         *string `json:"Order,omitempty" xml:"Order,omitempty"`
	PageNumber    *int64  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize      *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SortBy        *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	Status        *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Verbose       *bool   `json:"Verbose,omitempty" xml:"Verbose,omitempty"`
	WorkspaceIds  *string `json:"WorkspaceIds,omitempty" xml:"WorkspaceIds,omitempty"`
	WorkspaceName *string `json:"WorkspaceName,omitempty" xml:"WorkspaceName,omitempty"`
}

func (s ListWorkspacesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListWorkspacesRequest) GoString() string {
	return s.String()
}

func (s *ListWorkspacesRequest) SetFields(v string) *ListWorkspacesRequest {
	s.Fields = &v
	return s
}

func (s *ListWorkspacesRequest) SetModuleList(v string) *ListWorkspacesRequest {
	s.ModuleList = &v
	return s
}

func (s *ListWorkspacesRequest) SetOption(v string) *ListWorkspacesRequest {
	s.Option = &v
	return s
}

func (s *ListWorkspacesRequest) SetOrder(v string) *ListWorkspacesRequest {
	s.Order = &v
	return s
}

func (s *ListWorkspacesRequest) SetPageNumber(v int64) *ListWorkspacesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListWorkspacesRequest) SetPageSize(v int32) *ListWorkspacesRequest {
	s.PageSize = &v
	return s
}

func (s *ListWorkspacesRequest) SetSortBy(v string) *ListWorkspacesRequest {
	s.SortBy = &v
	return s
}

func (s *ListWorkspacesRequest) SetStatus(v string) *ListWorkspacesRequest {
	s.Status = &v
	return s
}

func (s *ListWorkspacesRequest) SetVerbose(v bool) *ListWorkspacesRequest {
	s.Verbose = &v
	return s
}

func (s *ListWorkspacesRequest) SetWorkspaceIds(v string) *ListWorkspacesRequest {
	s.WorkspaceIds = &v
	return s
}

func (s *ListWorkspacesRequest) SetWorkspaceName(v string) *ListWorkspacesRequest {
	s.WorkspaceName = &v
	return s
}

type ListWorkspacesResponseBody struct {
	RequestId      *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResourceLimits map[string]interface{}                  `json:"ResourceLimits,omitempty" xml:"ResourceLimits,omitempty"`
	TotalCount     *int64                                  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	Workspaces     []*ListWorkspacesResponseBodyWorkspaces `json:"Workspaces,omitempty" xml:"Workspaces,omitempty" type:"Repeated"`
}

func (s ListWorkspacesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListWorkspacesResponseBody) GoString() string {
	return s.String()
}

func (s *ListWorkspacesResponseBody) SetRequestId(v string) *ListWorkspacesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListWorkspacesResponseBody) SetResourceLimits(v map[string]interface{}) *ListWorkspacesResponseBody {
	s.ResourceLimits = v
	return s
}

func (s *ListWorkspacesResponseBody) SetTotalCount(v int64) *ListWorkspacesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListWorkspacesResponseBody) SetWorkspaces(v []*ListWorkspacesResponseBodyWorkspaces) *ListWorkspacesResponseBody {
	s.Workspaces = v
	return s
}

type ListWorkspacesResponseBodyWorkspaces struct {
	AdminNames      []*string              `json:"AdminNames,omitempty" xml:"AdminNames,omitempty" type:"Repeated"`
	Creator         *string                `json:"Creator,omitempty" xml:"Creator,omitempty"`
	Description     *string                `json:"Description,omitempty" xml:"Description,omitempty"`
	EnvTypes        []*string              `json:"EnvTypes,omitempty" xml:"EnvTypes,omitempty" type:"Repeated"`
	ExtraInfos      map[string]interface{} `json:"ExtraInfos,omitempty" xml:"ExtraInfos,omitempty"`
	GmtCreateTime   *string                `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	GmtModifiedTime *string                `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	IsDefault       *bool                  `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	Status          *string                `json:"Status,omitempty" xml:"Status,omitempty"`
	WorkspaceId     *string                `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
	WorkspaceName   *string                `json:"WorkspaceName,omitempty" xml:"WorkspaceName,omitempty"`
}

func (s ListWorkspacesResponseBodyWorkspaces) String() string {
	return tea.Prettify(s)
}

func (s ListWorkspacesResponseBodyWorkspaces) GoString() string {
	return s.String()
}

func (s *ListWorkspacesResponseBodyWorkspaces) SetAdminNames(v []*string) *ListWorkspacesResponseBodyWorkspaces {
	s.AdminNames = v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspaces) SetCreator(v string) *ListWorkspacesResponseBodyWorkspaces {
	s.Creator = &v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspaces) SetDescription(v string) *ListWorkspacesResponseBodyWorkspaces {
	s.Description = &v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspaces) SetEnvTypes(v []*string) *ListWorkspacesResponseBodyWorkspaces {
	s.EnvTypes = v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspaces) SetExtraInfos(v map[string]interface{}) *ListWorkspacesResponseBodyWorkspaces {
	s.ExtraInfos = v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspaces) SetGmtCreateTime(v string) *ListWorkspacesResponseBodyWorkspaces {
	s.GmtCreateTime = &v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspaces) SetGmtModifiedTime(v string) *ListWorkspacesResponseBodyWorkspaces {
	s.GmtModifiedTime = &v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspaces) SetIsDefault(v bool) *ListWorkspacesResponseBodyWorkspaces {
	s.IsDefault = &v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspaces) SetStatus(v string) *ListWorkspacesResponseBodyWorkspaces {
	s.Status = &v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspaces) SetWorkspaceId(v string) *ListWorkspacesResponseBodyWorkspaces {
	s.WorkspaceId = &v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspaces) SetWorkspaceName(v string) *ListWorkspacesResponseBodyWorkspaces {
	s.WorkspaceName = &v
	return s
}

type ListWorkspacesResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListWorkspacesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListWorkspacesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListWorkspacesResponse) GoString() string {
	return s.String()
}

func (s *ListWorkspacesResponse) SetHeaders(v map[string]*string) *ListWorkspacesResponse {
	s.Headers = v
	return s
}

func (s *ListWorkspacesResponse) SetStatusCode(v int32) *ListWorkspacesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListWorkspacesResponse) SetBody(v *ListWorkspacesResponseBody) *ListWorkspacesResponse {
	s.Body = v
	return s
}

type PublishCodeSourceResponseBody struct {
	CodeSourceId *string `json:"CodeSourceId,omitempty" xml:"CodeSourceId,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s PublishCodeSourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PublishCodeSourceResponseBody) GoString() string {
	return s.String()
}

func (s *PublishCodeSourceResponseBody) SetCodeSourceId(v string) *PublishCodeSourceResponseBody {
	s.CodeSourceId = &v
	return s
}

func (s *PublishCodeSourceResponseBody) SetRequestId(v string) *PublishCodeSourceResponseBody {
	s.RequestId = &v
	return s
}

type PublishCodeSourceResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PublishCodeSourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PublishCodeSourceResponse) String() string {
	return tea.Prettify(s)
}

func (s PublishCodeSourceResponse) GoString() string {
	return s.String()
}

func (s *PublishCodeSourceResponse) SetHeaders(v map[string]*string) *PublishCodeSourceResponse {
	s.Headers = v
	return s
}

func (s *PublishCodeSourceResponse) SetStatusCode(v int32) *PublishCodeSourceResponse {
	s.StatusCode = &v
	return s
}

func (s *PublishCodeSourceResponse) SetBody(v *PublishCodeSourceResponseBody) *PublishCodeSourceResponse {
	s.Body = v
	return s
}

type PublishDatasetResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s PublishDatasetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PublishDatasetResponseBody) GoString() string {
	return s.String()
}

func (s *PublishDatasetResponseBody) SetRequestId(v string) *PublishDatasetResponseBody {
	s.RequestId = &v
	return s
}

type PublishDatasetResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PublishDatasetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PublishDatasetResponse) String() string {
	return tea.Prettify(s)
}

func (s PublishDatasetResponse) GoString() string {
	return s.String()
}

func (s *PublishDatasetResponse) SetHeaders(v map[string]*string) *PublishDatasetResponse {
	s.Headers = v
	return s
}

func (s *PublishDatasetResponse) SetStatusCode(v int32) *PublishDatasetResponse {
	s.StatusCode = &v
	return s
}

func (s *PublishDatasetResponse) SetBody(v *PublishDatasetResponseBody) *PublishDatasetResponse {
	s.Body = v
	return s
}

type PublishImageResponseBody struct {
	ImageId   *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s PublishImageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PublishImageResponseBody) GoString() string {
	return s.String()
}

func (s *PublishImageResponseBody) SetImageId(v string) *PublishImageResponseBody {
	s.ImageId = &v
	return s
}

func (s *PublishImageResponseBody) SetRequestId(v string) *PublishImageResponseBody {
	s.RequestId = &v
	return s
}

type PublishImageResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PublishImageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PublishImageResponse) String() string {
	return tea.Prettify(s)
}

func (s PublishImageResponse) GoString() string {
	return s.String()
}

func (s *PublishImageResponse) SetHeaders(v map[string]*string) *PublishImageResponse {
	s.Headers = v
	return s
}

func (s *PublishImageResponse) SetStatusCode(v int32) *PublishImageResponse {
	s.StatusCode = &v
	return s
}

func (s *PublishImageResponse) SetBody(v *PublishImageResponseBody) *PublishImageResponse {
	s.Body = v
	return s
}

type RemoveImageResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RemoveImageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveImageResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveImageResponseBody) SetRequestId(v string) *RemoveImageResponseBody {
	s.RequestId = &v
	return s
}

type RemoveImageResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveImageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveImageResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveImageResponse) GoString() string {
	return s.String()
}

func (s *RemoveImageResponse) SetHeaders(v map[string]*string) *RemoveImageResponse {
	s.Headers = v
	return s
}

func (s *RemoveImageResponse) SetStatusCode(v int32) *RemoveImageResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveImageResponse) SetBody(v *RemoveImageResponseBody) *RemoveImageResponse {
	s.Body = v
	return s
}

type RemoveImageLabelsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RemoveImageLabelsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveImageLabelsResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveImageLabelsResponseBody) SetRequestId(v string) *RemoveImageLabelsResponseBody {
	s.RequestId = &v
	return s
}

type RemoveImageLabelsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveImageLabelsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveImageLabelsResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveImageLabelsResponse) GoString() string {
	return s.String()
}

func (s *RemoveImageLabelsResponse) SetHeaders(v map[string]*string) *RemoveImageLabelsResponse {
	s.Headers = v
	return s
}

func (s *RemoveImageLabelsResponse) SetStatusCode(v int32) *RemoveImageLabelsResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveImageLabelsResponse) SetBody(v *RemoveImageLabelsResponseBody) *RemoveImageLabelsResponse {
	s.Body = v
	return s
}

type RemoveMemberRoleResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RemoveMemberRoleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveMemberRoleResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveMemberRoleResponseBody) SetRequestId(v string) *RemoveMemberRoleResponseBody {
	s.RequestId = &v
	return s
}

type RemoveMemberRoleResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveMemberRoleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveMemberRoleResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveMemberRoleResponse) GoString() string {
	return s.String()
}

func (s *RemoveMemberRoleResponse) SetHeaders(v map[string]*string) *RemoveMemberRoleResponse {
	s.Headers = v
	return s
}

func (s *RemoveMemberRoleResponse) SetStatusCode(v int32) *RemoveMemberRoleResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveMemberRoleResponse) SetBody(v *RemoveMemberRoleResponseBody) *RemoveMemberRoleResponse {
	s.Body = v
	return s
}

type UpdateDatasetRequest struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Options     *string `json:"Options,omitempty" xml:"Options,omitempty"`
}

func (s UpdateDatasetRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateDatasetRequest) GoString() string {
	return s.String()
}

func (s *UpdateDatasetRequest) SetDescription(v string) *UpdateDatasetRequest {
	s.Description = &v
	return s
}

func (s *UpdateDatasetRequest) SetName(v string) *UpdateDatasetRequest {
	s.Name = &v
	return s
}

func (s *UpdateDatasetRequest) SetOptions(v string) *UpdateDatasetRequest {
	s.Options = &v
	return s
}

type UpdateDatasetResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateDatasetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateDatasetResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDatasetResponseBody) SetRequestId(v string) *UpdateDatasetResponseBody {
	s.RequestId = &v
	return s
}

type UpdateDatasetResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateDatasetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateDatasetResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateDatasetResponse) GoString() string {
	return s.String()
}

func (s *UpdateDatasetResponse) SetHeaders(v map[string]*string) *UpdateDatasetResponse {
	s.Headers = v
	return s
}

func (s *UpdateDatasetResponse) SetStatusCode(v int32) *UpdateDatasetResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateDatasetResponse) SetBody(v *UpdateDatasetResponseBody) *UpdateDatasetResponse {
	s.Body = v
	return s
}

type UpdateDefaultWorkspaceRequest struct {
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s UpdateDefaultWorkspaceRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateDefaultWorkspaceRequest) GoString() string {
	return s.String()
}

func (s *UpdateDefaultWorkspaceRequest) SetWorkspaceId(v string) *UpdateDefaultWorkspaceRequest {
	s.WorkspaceId = &v
	return s
}

type UpdateDefaultWorkspaceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateDefaultWorkspaceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateDefaultWorkspaceResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDefaultWorkspaceResponseBody) SetRequestId(v string) *UpdateDefaultWorkspaceResponseBody {
	s.RequestId = &v
	return s
}

type UpdateDefaultWorkspaceResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateDefaultWorkspaceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateDefaultWorkspaceResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateDefaultWorkspaceResponse) GoString() string {
	return s.String()
}

func (s *UpdateDefaultWorkspaceResponse) SetHeaders(v map[string]*string) *UpdateDefaultWorkspaceResponse {
	s.Headers = v
	return s
}

func (s *UpdateDefaultWorkspaceResponse) SetStatusCode(v int32) *UpdateDefaultWorkspaceResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateDefaultWorkspaceResponse) SetBody(v *UpdateDefaultWorkspaceResponseBody) *UpdateDefaultWorkspaceResponse {
	s.Body = v
	return s
}

type UpdateModelRequest struct {
	Accessibility    *string                `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	Domain           *string                `json:"Domain,omitempty" xml:"Domain,omitempty"`
	ExtraInfo        map[string]interface{} `json:"ExtraInfo,omitempty" xml:"ExtraInfo,omitempty"`
	ModelDescription *string                `json:"ModelDescription,omitempty" xml:"ModelDescription,omitempty"`
	ModelDoc         *string                `json:"ModelDoc,omitempty" xml:"ModelDoc,omitempty"`
	ModelName        *string                `json:"ModelName,omitempty" xml:"ModelName,omitempty"`
	ModelType        *string                `json:"ModelType,omitempty" xml:"ModelType,omitempty"`
	OrderNumber      *int64                 `json:"OrderNumber,omitempty" xml:"OrderNumber,omitempty"`
	Origin           *string                `json:"Origin,omitempty" xml:"Origin,omitempty"`
	Task             *string                `json:"Task,omitempty" xml:"Task,omitempty"`
}

func (s UpdateModelRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateModelRequest) GoString() string {
	return s.String()
}

func (s *UpdateModelRequest) SetAccessibility(v string) *UpdateModelRequest {
	s.Accessibility = &v
	return s
}

func (s *UpdateModelRequest) SetDomain(v string) *UpdateModelRequest {
	s.Domain = &v
	return s
}

func (s *UpdateModelRequest) SetExtraInfo(v map[string]interface{}) *UpdateModelRequest {
	s.ExtraInfo = v
	return s
}

func (s *UpdateModelRequest) SetModelDescription(v string) *UpdateModelRequest {
	s.ModelDescription = &v
	return s
}

func (s *UpdateModelRequest) SetModelDoc(v string) *UpdateModelRequest {
	s.ModelDoc = &v
	return s
}

func (s *UpdateModelRequest) SetModelName(v string) *UpdateModelRequest {
	s.ModelName = &v
	return s
}

func (s *UpdateModelRequest) SetModelType(v string) *UpdateModelRequest {
	s.ModelType = &v
	return s
}

func (s *UpdateModelRequest) SetOrderNumber(v int64) *UpdateModelRequest {
	s.OrderNumber = &v
	return s
}

func (s *UpdateModelRequest) SetOrigin(v string) *UpdateModelRequest {
	s.Origin = &v
	return s
}

func (s *UpdateModelRequest) SetTask(v string) *UpdateModelRequest {
	s.Task = &v
	return s
}

type UpdateModelResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateModelResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateModelResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateModelResponseBody) SetRequestId(v string) *UpdateModelResponseBody {
	s.RequestId = &v
	return s
}

type UpdateModelResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateModelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateModelResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateModelResponse) GoString() string {
	return s.String()
}

func (s *UpdateModelResponse) SetHeaders(v map[string]*string) *UpdateModelResponse {
	s.Headers = v
	return s
}

func (s *UpdateModelResponse) SetStatusCode(v int32) *UpdateModelResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateModelResponse) SetBody(v *UpdateModelResponseBody) *UpdateModelResponse {
	s.Body = v
	return s
}

type UpdateModelVersionRequest struct {
	ApprovalStatus     *string                `json:"ApprovalStatus,omitempty" xml:"ApprovalStatus,omitempty"`
	EvaluationSpec     map[string]interface{} `json:"EvaluationSpec,omitempty" xml:"EvaluationSpec,omitempty"`
	ExtraInfo          map[string]interface{} `json:"ExtraInfo,omitempty" xml:"ExtraInfo,omitempty"`
	InferenceSpec      map[string]interface{} `json:"InferenceSpec,omitempty" xml:"InferenceSpec,omitempty"`
	Metrics            map[string]interface{} `json:"Metrics,omitempty" xml:"Metrics,omitempty"`
	Options            *string                `json:"Options,omitempty" xml:"Options,omitempty"`
	SourceId           *string                `json:"SourceId,omitempty" xml:"SourceId,omitempty"`
	SourceType         *string                `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	TrainingSpec       map[string]interface{} `json:"TrainingSpec,omitempty" xml:"TrainingSpec,omitempty"`
	VersionDescription *string                `json:"VersionDescription,omitempty" xml:"VersionDescription,omitempty"`
}

func (s UpdateModelVersionRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateModelVersionRequest) GoString() string {
	return s.String()
}

func (s *UpdateModelVersionRequest) SetApprovalStatus(v string) *UpdateModelVersionRequest {
	s.ApprovalStatus = &v
	return s
}

func (s *UpdateModelVersionRequest) SetEvaluationSpec(v map[string]interface{}) *UpdateModelVersionRequest {
	s.EvaluationSpec = v
	return s
}

func (s *UpdateModelVersionRequest) SetExtraInfo(v map[string]interface{}) *UpdateModelVersionRequest {
	s.ExtraInfo = v
	return s
}

func (s *UpdateModelVersionRequest) SetInferenceSpec(v map[string]interface{}) *UpdateModelVersionRequest {
	s.InferenceSpec = v
	return s
}

func (s *UpdateModelVersionRequest) SetMetrics(v map[string]interface{}) *UpdateModelVersionRequest {
	s.Metrics = v
	return s
}

func (s *UpdateModelVersionRequest) SetOptions(v string) *UpdateModelVersionRequest {
	s.Options = &v
	return s
}

func (s *UpdateModelVersionRequest) SetSourceId(v string) *UpdateModelVersionRequest {
	s.SourceId = &v
	return s
}

func (s *UpdateModelVersionRequest) SetSourceType(v string) *UpdateModelVersionRequest {
	s.SourceType = &v
	return s
}

func (s *UpdateModelVersionRequest) SetTrainingSpec(v map[string]interface{}) *UpdateModelVersionRequest {
	s.TrainingSpec = v
	return s
}

func (s *UpdateModelVersionRequest) SetVersionDescription(v string) *UpdateModelVersionRequest {
	s.VersionDescription = &v
	return s
}

type UpdateModelVersionResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateModelVersionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateModelVersionResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateModelVersionResponseBody) SetRequestId(v string) *UpdateModelVersionResponseBody {
	s.RequestId = &v
	return s
}

type UpdateModelVersionResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateModelVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateModelVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateModelVersionResponse) GoString() string {
	return s.String()
}

func (s *UpdateModelVersionResponse) SetHeaders(v map[string]*string) *UpdateModelVersionResponse {
	s.Headers = v
	return s
}

func (s *UpdateModelVersionResponse) SetStatusCode(v int32) *UpdateModelVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateModelVersionResponse) SetBody(v *UpdateModelVersionResponseBody) *UpdateModelVersionResponse {
	s.Body = v
	return s
}

type UpdateWorkspaceRequest struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
}

func (s UpdateWorkspaceRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateWorkspaceRequest) GoString() string {
	return s.String()
}

func (s *UpdateWorkspaceRequest) SetDescription(v string) *UpdateWorkspaceRequest {
	s.Description = &v
	return s
}

func (s *UpdateWorkspaceRequest) SetDisplayName(v string) *UpdateWorkspaceRequest {
	s.DisplayName = &v
	return s
}

type UpdateWorkspaceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateWorkspaceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateWorkspaceResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateWorkspaceResponseBody) SetRequestId(v string) *UpdateWorkspaceResponseBody {
	s.RequestId = &v
	return s
}

type UpdateWorkspaceResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateWorkspaceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateWorkspaceResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateWorkspaceResponse) GoString() string {
	return s.String()
}

func (s *UpdateWorkspaceResponse) SetHeaders(v map[string]*string) *UpdateWorkspaceResponse {
	s.Headers = v
	return s
}

func (s *UpdateWorkspaceResponse) SetStatusCode(v int32) *UpdateWorkspaceResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateWorkspaceResponse) SetBody(v *UpdateWorkspaceResponseBody) *UpdateWorkspaceResponse {
	s.Body = v
	return s
}

type UpdateWorkspaceResourceRequest struct {
	GroupName    *string                                 `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	IsDefault    *bool                                   `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	Labels       []*UpdateWorkspaceResourceRequestLabels `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	ProductType  *string                                 `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	ResourceIds  []*string                               `json:"ResourceIds,omitempty" xml:"ResourceIds,omitempty" type:"Repeated"`
	ResourceType *string                                 `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Spec         map[string]interface{}                  `json:"Spec,omitempty" xml:"Spec,omitempty"`
}

func (s UpdateWorkspaceResourceRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateWorkspaceResourceRequest) GoString() string {
	return s.String()
}

func (s *UpdateWorkspaceResourceRequest) SetGroupName(v string) *UpdateWorkspaceResourceRequest {
	s.GroupName = &v
	return s
}

func (s *UpdateWorkspaceResourceRequest) SetIsDefault(v bool) *UpdateWorkspaceResourceRequest {
	s.IsDefault = &v
	return s
}

func (s *UpdateWorkspaceResourceRequest) SetLabels(v []*UpdateWorkspaceResourceRequestLabels) *UpdateWorkspaceResourceRequest {
	s.Labels = v
	return s
}

func (s *UpdateWorkspaceResourceRequest) SetProductType(v string) *UpdateWorkspaceResourceRequest {
	s.ProductType = &v
	return s
}

func (s *UpdateWorkspaceResourceRequest) SetResourceIds(v []*string) *UpdateWorkspaceResourceRequest {
	s.ResourceIds = v
	return s
}

func (s *UpdateWorkspaceResourceRequest) SetResourceType(v string) *UpdateWorkspaceResourceRequest {
	s.ResourceType = &v
	return s
}

func (s *UpdateWorkspaceResourceRequest) SetSpec(v map[string]interface{}) *UpdateWorkspaceResourceRequest {
	s.Spec = v
	return s
}

type UpdateWorkspaceResourceRequestLabels struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateWorkspaceResourceRequestLabels) String() string {
	return tea.Prettify(s)
}

func (s UpdateWorkspaceResourceRequestLabels) GoString() string {
	return s.String()
}

func (s *UpdateWorkspaceResourceRequestLabels) SetKey(v string) *UpdateWorkspaceResourceRequestLabels {
	s.Key = &v
	return s
}

func (s *UpdateWorkspaceResourceRequestLabels) SetValue(v string) *UpdateWorkspaceResourceRequestLabels {
	s.Value = &v
	return s
}

type UpdateWorkspaceResourceResponseBody struct {
	RequestId   *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResourceIds []*string `json:"ResourceIds,omitempty" xml:"ResourceIds,omitempty" type:"Repeated"`
}

func (s UpdateWorkspaceResourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateWorkspaceResourceResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateWorkspaceResourceResponseBody) SetRequestId(v string) *UpdateWorkspaceResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateWorkspaceResourceResponseBody) SetResourceIds(v []*string) *UpdateWorkspaceResourceResponseBody {
	s.ResourceIds = v
	return s
}

type UpdateWorkspaceResourceResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateWorkspaceResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateWorkspaceResourceResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateWorkspaceResourceResponse) GoString() string {
	return s.String()
}

func (s *UpdateWorkspaceResourceResponse) SetHeaders(v map[string]*string) *UpdateWorkspaceResourceResponse {
	s.Headers = v
	return s
}

func (s *UpdateWorkspaceResourceResponse) SetStatusCode(v int32) *UpdateWorkspaceResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateWorkspaceResourceResponse) SetBody(v *UpdateWorkspaceResourceResponseBody) *UpdateWorkspaceResourceResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("aiworkspace"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) AddImageWithOptions(request *AddImageRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *AddImageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Accessibility)) {
		body["Accessibility"] = request.Accessibility
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.ImageId)) {
		body["ImageId"] = request.ImageId
	}

	if !tea.BoolValue(util.IsUnset(request.ImageUri)) {
		body["ImageUri"] = request.ImageUri
	}

	if !tea.BoolValue(util.IsUnset(request.Labels)) {
		body["Labels"] = request.Labels
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Size)) {
		body["Size"] = request.Size
	}

	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("AddImage"),
		Version:     tea.String("2021-02-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/images"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &AddImageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddImage(request *AddImageRequest) (_result *AddImageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &AddImageResponse{}
	_body, _err := client.AddImageWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddImageLabelsWithOptions(ImageId *string, request *AddImageLabelsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *AddImageLabelsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Labels)) {
		body["Labels"] = request.Labels
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("AddImageLabels"),
		Version:     tea.String("2021-02-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/images/" + tea.StringValue(openapiutil.GetEncodeParam(ImageId)) + "/labels"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &AddImageLabelsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddImageLabels(ImageId *string, request *AddImageLabelsRequest) (_result *AddImageLabelsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &AddImageLabelsResponse{}
	_body, _err := client.AddImageLabelsWithOptions(ImageId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddMemberRoleWithOptions(WorkspaceId *string, MemberId *string, RoleName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *AddMemberRoleResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("AddMemberRole"),
		Version:     tea.String("2021-02-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/workspaces/" + tea.StringValue(openapiutil.GetEncodeParam(WorkspaceId)) + "/members/" + tea.StringValue(openapiutil.GetEncodeParam(MemberId)) + "/roles/" + tea.StringValue(openapiutil.GetEncodeParam(RoleName))),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &AddMemberRoleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddMemberRole(WorkspaceId *string, MemberId *string, RoleName *string) (_result *AddMemberRoleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &AddMemberRoleResponse{}
	_body, _err := client.AddMemberRoleWithOptions(WorkspaceId, MemberId, RoleName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateCodeSourceWithOptions(request *CreateCodeSourceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateCodeSourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Accessibility)) {
		body["Accessibility"] = request.Accessibility
	}

	if !tea.BoolValue(util.IsUnset(request.CodeBranch)) {
		body["CodeBranch"] = request.CodeBranch
	}

	if !tea.BoolValue(util.IsUnset(request.CodeRepo)) {
		body["CodeRepo"] = request.CodeRepo
	}

	if !tea.BoolValue(util.IsUnset(request.CodeRepoAccessToken)) {
		body["CodeRepoAccessToken"] = request.CodeRepoAccessToken
	}

	if !tea.BoolValue(util.IsUnset(request.CodeRepoUserName)) {
		body["CodeRepoUserName"] = request.CodeRepoUserName
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.DisplayName)) {
		body["DisplayName"] = request.DisplayName
	}

	if !tea.BoolValue(util.IsUnset(request.MountPath)) {
		body["MountPath"] = request.MountPath
	}

	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateCodeSource"),
		Version:     tea.String("2021-02-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/codesources"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateCodeSourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateCodeSource(request *CreateCodeSourceRequest) (_result *CreateCodeSourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateCodeSourceResponse{}
	_body, _err := client.CreateCodeSourceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateDatasetWithOptions(request *CreateDatasetRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateDatasetResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Accessibility)) {
		body["Accessibility"] = request.Accessibility
	}

	if !tea.BoolValue(util.IsUnset(request.DataSourceType)) {
		body["DataSourceType"] = request.DataSourceType
	}

	if !tea.BoolValue(util.IsUnset(request.DataType)) {
		body["DataType"] = request.DataType
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Labels)) {
		body["Labels"] = request.Labels
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Options)) {
		body["Options"] = request.Options
	}

	if !tea.BoolValue(util.IsUnset(request.Property)) {
		body["Property"] = request.Property
	}

	if !tea.BoolValue(util.IsUnset(request.Provider)) {
		body["Provider"] = request.Provider
	}

	if !tea.BoolValue(util.IsUnset(request.ProviderType)) {
		body["ProviderType"] = request.ProviderType
	}

	if !tea.BoolValue(util.IsUnset(request.SourceId)) {
		body["SourceId"] = request.SourceId
	}

	if !tea.BoolValue(util.IsUnset(request.SourceType)) {
		body["SourceType"] = request.SourceType
	}

	if !tea.BoolValue(util.IsUnset(request.Uri)) {
		body["Uri"] = request.Uri
	}

	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateDataset"),
		Version:     tea.String("2021-02-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/datasets"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateDatasetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateDataset(request *CreateDatasetRequest) (_result *CreateDatasetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateDatasetResponse{}
	_body, _err := client.CreateDatasetWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateDatasetLabelsWithOptions(DatasetId *string, request *CreateDatasetLabelsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateDatasetLabelsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Labels)) {
		body["Labels"] = request.Labels
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateDatasetLabels"),
		Version:     tea.String("2021-02-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/datasets/" + tea.StringValue(openapiutil.GetEncodeParam(DatasetId)) + "/labels"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateDatasetLabelsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateDatasetLabels(DatasetId *string, request *CreateDatasetLabelsRequest) (_result *CreateDatasetLabelsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateDatasetLabelsResponse{}
	_body, _err := client.CreateDatasetLabelsWithOptions(DatasetId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateMemberWithOptions(WorkspaceId *string, request *CreateMemberRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateMemberResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Members)) {
		body["Members"] = request.Members
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateMember"),
		Version:     tea.String("2021-02-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/workspaces/" + tea.StringValue(openapiutil.GetEncodeParam(WorkspaceId)) + "/members"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateMemberResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateMember(WorkspaceId *string, request *CreateMemberRequest) (_result *CreateMemberResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateMemberResponse{}
	_body, _err := client.CreateMemberWithOptions(WorkspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateModelWithOptions(request *CreateModelRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateModelResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Accessibility)) {
		body["Accessibility"] = request.Accessibility
	}

	if !tea.BoolValue(util.IsUnset(request.Domain)) {
		body["Domain"] = request.Domain
	}

	if !tea.BoolValue(util.IsUnset(request.ExtraInfo)) {
		body["ExtraInfo"] = request.ExtraInfo
	}

	if !tea.BoolValue(util.IsUnset(request.Labels)) {
		body["Labels"] = request.Labels
	}

	if !tea.BoolValue(util.IsUnset(request.ModelDescription)) {
		body["ModelDescription"] = request.ModelDescription
	}

	if !tea.BoolValue(util.IsUnset(request.ModelDoc)) {
		body["ModelDoc"] = request.ModelDoc
	}

	if !tea.BoolValue(util.IsUnset(request.ModelName)) {
		body["ModelName"] = request.ModelName
	}

	if !tea.BoolValue(util.IsUnset(request.ModelType)) {
		body["ModelType"] = request.ModelType
	}

	if !tea.BoolValue(util.IsUnset(request.OrderNumber)) {
		body["OrderNumber"] = request.OrderNumber
	}

	if !tea.BoolValue(util.IsUnset(request.Origin)) {
		body["Origin"] = request.Origin
	}

	if !tea.BoolValue(util.IsUnset(request.Task)) {
		body["Task"] = request.Task
	}

	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateModel"),
		Version:     tea.String("2021-02-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/models"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateModelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateModel(request *CreateModelRequest) (_result *CreateModelResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateModelResponse{}
	_body, _err := client.CreateModelWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateModelLabelsWithOptions(ModelId *string, request *CreateModelLabelsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateModelLabelsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Labels)) {
		body["Labels"] = request.Labels
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateModelLabels"),
		Version:     tea.String("2021-02-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/models/" + tea.StringValue(openapiutil.GetEncodeParam(ModelId)) + "/labels"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateModelLabelsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateModelLabels(ModelId *string, request *CreateModelLabelsRequest) (_result *CreateModelLabelsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateModelLabelsResponse{}
	_body, _err := client.CreateModelLabelsWithOptions(ModelId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateModelVersionWithOptions(ModelId *string, request *CreateModelVersionRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateModelVersionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApprovalStatus)) {
		body["ApprovalStatus"] = request.ApprovalStatus
	}

	if !tea.BoolValue(util.IsUnset(request.EvaluationSpec)) {
		body["EvaluationSpec"] = request.EvaluationSpec
	}

	if !tea.BoolValue(util.IsUnset(request.ExtraInfo)) {
		body["ExtraInfo"] = request.ExtraInfo
	}

	if !tea.BoolValue(util.IsUnset(request.FormatType)) {
		body["FormatType"] = request.FormatType
	}

	if !tea.BoolValue(util.IsUnset(request.FrameworkType)) {
		body["FrameworkType"] = request.FrameworkType
	}

	if !tea.BoolValue(util.IsUnset(request.InferenceSpec)) {
		body["InferenceSpec"] = request.InferenceSpec
	}

	if !tea.BoolValue(util.IsUnset(request.Labels)) {
		body["Labels"] = request.Labels
	}

	if !tea.BoolValue(util.IsUnset(request.Metrics)) {
		body["Metrics"] = request.Metrics
	}

	if !tea.BoolValue(util.IsUnset(request.Options)) {
		body["Options"] = request.Options
	}

	if !tea.BoolValue(util.IsUnset(request.SourceId)) {
		body["SourceId"] = request.SourceId
	}

	if !tea.BoolValue(util.IsUnset(request.SourceType)) {
		body["SourceType"] = request.SourceType
	}

	if !tea.BoolValue(util.IsUnset(request.TrainingSpec)) {
		body["TrainingSpec"] = request.TrainingSpec
	}

	if !tea.BoolValue(util.IsUnset(request.Uri)) {
		body["Uri"] = request.Uri
	}

	if !tea.BoolValue(util.IsUnset(request.VersionDescription)) {
		body["VersionDescription"] = request.VersionDescription
	}

	if !tea.BoolValue(util.IsUnset(request.VersionName)) {
		body["VersionName"] = request.VersionName
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateModelVersion"),
		Version:     tea.String("2021-02-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/models/" + tea.StringValue(openapiutil.GetEncodeParam(ModelId)) + "/versions"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateModelVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateModelVersion(ModelId *string, request *CreateModelVersionRequest) (_result *CreateModelVersionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateModelVersionResponse{}
	_body, _err := client.CreateModelVersionWithOptions(ModelId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateModelVersionLabelsWithOptions(ModelId *string, VersionName *string, request *CreateModelVersionLabelsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateModelVersionLabelsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Labels)) {
		body["Labels"] = request.Labels
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateModelVersionLabels"),
		Version:     tea.String("2021-02-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/models/" + tea.StringValue(openapiutil.GetEncodeParam(ModelId)) + "/versions/" + tea.StringValue(openapiutil.GetEncodeParam(VersionName)) + "/labels"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateModelVersionLabelsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateModelVersionLabels(ModelId *string, VersionName *string, request *CreateModelVersionLabelsRequest) (_result *CreateModelVersionLabelsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateModelVersionLabelsResponse{}
	_body, _err := client.CreateModelVersionLabelsWithOptions(ModelId, VersionName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateProductOrdersWithOptions(request *CreateProductOrdersRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateProductOrdersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AutoPay)) {
		body["AutoPay"] = request.AutoPay
	}

	if !tea.BoolValue(util.IsUnset(request.Products)) {
		body["Products"] = request.Products
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateProductOrders"),
		Version:     tea.String("2021-02-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/productorders"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateProductOrdersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateProductOrders(request *CreateProductOrdersRequest) (_result *CreateProductOrdersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateProductOrdersResponse{}
	_body, _err := client.CreateProductOrdersWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateWorkspaceWithOptions(request *CreateWorkspaceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateWorkspaceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.DisplayName)) {
		body["DisplayName"] = request.DisplayName
	}

	if !tea.BoolValue(util.IsUnset(request.EnvTypes)) {
		body["EnvTypes"] = request.EnvTypes
	}

	if !tea.BoolValue(util.IsUnset(request.WorkspaceName)) {
		body["WorkspaceName"] = request.WorkspaceName
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateWorkspace"),
		Version:     tea.String("2021-02-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/workspaces"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateWorkspaceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateWorkspace(request *CreateWorkspaceRequest) (_result *CreateWorkspaceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateWorkspaceResponse{}
	_body, _err := client.CreateWorkspaceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateWorkspaceResourceWithOptions(WorkspaceId *string, request *CreateWorkspaceResourceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateWorkspaceResourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Option)) {
		body["Option"] = request.Option
	}

	if !tea.BoolValue(util.IsUnset(request.Resources)) {
		body["Resources"] = request.Resources
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateWorkspaceResource"),
		Version:     tea.String("2021-02-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/workspaces/" + tea.StringValue(openapiutil.GetEncodeParam(WorkspaceId)) + "/resources"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateWorkspaceResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateWorkspaceResource(WorkspaceId *string, request *CreateWorkspaceResourceRequest) (_result *CreateWorkspaceResourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateWorkspaceResourceResponse{}
	_body, _err := client.CreateWorkspaceResourceWithOptions(WorkspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteCodeSourceWithOptions(CodeSourceId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteCodeSourceResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteCodeSource"),
		Version:     tea.String("2021-02-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/codesources/" + tea.StringValue(openapiutil.GetEncodeParam(CodeSourceId))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteCodeSourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteCodeSource(CodeSourceId *string) (_result *DeleteCodeSourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteCodeSourceResponse{}
	_body, _err := client.DeleteCodeSourceWithOptions(CodeSourceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteDatasetWithOptions(DatasetId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteDatasetResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteDataset"),
		Version:     tea.String("2021-02-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/datasets/" + tea.StringValue(openapiutil.GetEncodeParam(DatasetId))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteDatasetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteDataset(DatasetId *string) (_result *DeleteDatasetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteDatasetResponse{}
	_body, _err := client.DeleteDatasetWithOptions(DatasetId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteDatasetLabelsWithOptions(DatasetId *string, request *DeleteDatasetLabelsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteDatasetLabelsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.LabelKeys)) {
		query["LabelKeys"] = request.LabelKeys
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteDatasetLabels"),
		Version:     tea.String("2021-02-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/datasets/" + tea.StringValue(openapiutil.GetEncodeParam(DatasetId)) + "/labels"),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteDatasetLabelsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteDatasetLabels(DatasetId *string, request *DeleteDatasetLabelsRequest) (_result *DeleteDatasetLabelsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteDatasetLabelsResponse{}
	_body, _err := client.DeleteDatasetLabelsWithOptions(DatasetId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteMembersWithOptions(WorkspaceId *string, request *DeleteMembersRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteMembersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MemberIds)) {
		query["MemberIds"] = request.MemberIds
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteMembers"),
		Version:     tea.String("2021-02-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/workspaces/" + tea.StringValue(openapiutil.GetEncodeParam(WorkspaceId)) + "/members"),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteMembersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteMembers(WorkspaceId *string, request *DeleteMembersRequest) (_result *DeleteMembersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteMembersResponse{}
	_body, _err := client.DeleteMembersWithOptions(WorkspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteModelWithOptions(ModelId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteModelResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteModel"),
		Version:     tea.String("2021-02-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/models/" + tea.StringValue(openapiutil.GetEncodeParam(ModelId))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteModelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteModel(ModelId *string) (_result *DeleteModelResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteModelResponse{}
	_body, _err := client.DeleteModelWithOptions(ModelId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteModelLabelsWithOptions(ModelId *string, request *DeleteModelLabelsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteModelLabelsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.LabelKeys)) {
		query["LabelKeys"] = request.LabelKeys
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteModelLabels"),
		Version:     tea.String("2021-02-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/models/" + tea.StringValue(openapiutil.GetEncodeParam(ModelId)) + "/labels"),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteModelLabelsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteModelLabels(ModelId *string, request *DeleteModelLabelsRequest) (_result *DeleteModelLabelsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteModelLabelsResponse{}
	_body, _err := client.DeleteModelLabelsWithOptions(ModelId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteModelVersionWithOptions(ModelId *string, VersionName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteModelVersionResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteModelVersion"),
		Version:     tea.String("2021-02-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/models/" + tea.StringValue(openapiutil.GetEncodeParam(ModelId)) + "/versions/" + tea.StringValue(openapiutil.GetEncodeParam(VersionName))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteModelVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteModelVersion(ModelId *string, VersionName *string) (_result *DeleteModelVersionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteModelVersionResponse{}
	_body, _err := client.DeleteModelVersionWithOptions(ModelId, VersionName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteModelVersionLabelsWithOptions(ModelId *string, VersionName *string, request *DeleteModelVersionLabelsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteModelVersionLabelsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.LabelKeys)) {
		query["LabelKeys"] = request.LabelKeys
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteModelVersionLabels"),
		Version:     tea.String("2021-02-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/models/" + tea.StringValue(openapiutil.GetEncodeParam(ModelId)) + "/versions/" + tea.StringValue(openapiutil.GetEncodeParam(VersionName)) + "/labels"),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteModelVersionLabelsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteModelVersionLabels(ModelId *string, VersionName *string, request *DeleteModelVersionLabelsRequest) (_result *DeleteModelVersionLabelsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteModelVersionLabelsResponse{}
	_body, _err := client.DeleteModelVersionLabelsWithOptions(ModelId, VersionName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteWorkspaceWithOptions(WorkspaceId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteWorkspaceResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteWorkspace"),
		Version:     tea.String("2021-02-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/workspaces/" + tea.StringValue(openapiutil.GetEncodeParam(WorkspaceId))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteWorkspaceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteWorkspace(WorkspaceId *string) (_result *DeleteWorkspaceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteWorkspaceResponse{}
	_body, _err := client.DeleteWorkspaceWithOptions(WorkspaceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteWorkspaceResourceWithOptions(WorkspaceId *string, request *DeleteWorkspaceResourceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteWorkspaceResourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GroupName)) {
		query["GroupName"] = request.GroupName
	}

	if !tea.BoolValue(util.IsUnset(request.Labels)) {
		query["Labels"] = request.Labels
	}

	if !tea.BoolValue(util.IsUnset(request.Option)) {
		query["Option"] = request.Option
	}

	if !tea.BoolValue(util.IsUnset(request.ProductType)) {
		query["ProductType"] = request.ProductType
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceIds)) {
		query["ResourceIds"] = request.ResourceIds
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteWorkspaceResource"),
		Version:     tea.String("2021-02-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/workspaces/" + tea.StringValue(openapiutil.GetEncodeParam(WorkspaceId)) + "/resources"),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteWorkspaceResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteWorkspaceResource(WorkspaceId *string, request *DeleteWorkspaceResourceRequest) (_result *DeleteWorkspaceResourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteWorkspaceResourceResponse{}
	_body, _err := client.DeleteWorkspaceResourceWithOptions(WorkspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetCodeSourceWithOptions(CodeSourceId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetCodeSourceResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetCodeSource"),
		Version:     tea.String("2021-02-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/codesources/" + tea.StringValue(openapiutil.GetEncodeParam(CodeSourceId))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetCodeSourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetCodeSource(CodeSourceId *string) (_result *GetCodeSourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetCodeSourceResponse{}
	_body, _err := client.GetCodeSourceWithOptions(CodeSourceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetDatasetWithOptions(DatasetId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetDatasetResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetDataset"),
		Version:     tea.String("2021-02-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/datasets/" + tea.StringValue(openapiutil.GetEncodeParam(DatasetId))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDatasetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetDataset(DatasetId *string) (_result *GetDatasetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetDatasetResponse{}
	_body, _err := client.GetDatasetWithOptions(DatasetId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetDefaultWorkspaceWithOptions(request *GetDefaultWorkspaceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetDefaultWorkspaceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Verbose)) {
		query["Verbose"] = request.Verbose
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetDefaultWorkspace"),
		Version:     tea.String("2021-02-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/defaultWorkspaces"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDefaultWorkspaceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetDefaultWorkspace(request *GetDefaultWorkspaceRequest) (_result *GetDefaultWorkspaceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetDefaultWorkspaceResponse{}
	_body, _err := client.GetDefaultWorkspaceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetImageWithOptions(ImageId *string, request *GetImageRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetImageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Verbose)) {
		query["Verbose"] = request.Verbose
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetImage"),
		Version:     tea.String("2021-02-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/images/" + tea.StringValue(openapiutil.GetEncodeParam(ImageId))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetImageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetImage(ImageId *string, request *GetImageRequest) (_result *GetImageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetImageResponse{}
	_body, _err := client.GetImageWithOptions(ImageId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetMemberWithOptions(WorkspaceId *string, request *GetMemberRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetMemberResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MemberId)) {
		query["MemberId"] = request.MemberId
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetMember"),
		Version:     tea.String("2021-02-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/workspaces/" + tea.StringValue(openapiutil.GetEncodeParam(WorkspaceId)) + "/member"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetMemberResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetMember(WorkspaceId *string, request *GetMemberRequest) (_result *GetMemberResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetMemberResponse{}
	_body, _err := client.GetMemberWithOptions(WorkspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetModelWithOptions(ModelId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetModelResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetModel"),
		Version:     tea.String("2021-02-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/models/" + tea.StringValue(openapiutil.GetEncodeParam(ModelId))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetModelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetModel(ModelId *string) (_result *GetModelResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetModelResponse{}
	_body, _err := client.GetModelWithOptions(ModelId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetModelVersionWithOptions(ModelId *string, VersionName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetModelVersionResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetModelVersion"),
		Version:     tea.String("2021-02-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/models/" + tea.StringValue(openapiutil.GetEncodeParam(ModelId)) + "/versions/" + tea.StringValue(openapiutil.GetEncodeParam(VersionName))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetModelVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetModelVersion(ModelId *string, VersionName *string) (_result *GetModelVersionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetModelVersionResponse{}
	_body, _err := client.GetModelVersionWithOptions(ModelId, VersionName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetPermissionWithOptions(WorkspaceId *string, PermissionCode *string, request *GetPermissionRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetPermissionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Accessibility)) {
		query["Accessibility"] = request.Accessibility
	}

	if !tea.BoolValue(util.IsUnset(request.Creator)) {
		query["Creator"] = request.Creator
	}

	if !tea.BoolValue(util.IsUnset(request.Option)) {
		query["Option"] = request.Option
	}

	if !tea.BoolValue(util.IsUnset(request.Resource)) {
		query["Resource"] = request.Resource
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetPermission"),
		Version:     tea.String("2021-02-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/workspaces/" + tea.StringValue(openapiutil.GetEncodeParam(WorkspaceId)) + "/permissions/" + tea.StringValue(openapiutil.GetEncodeParam(PermissionCode))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetPermissionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetPermission(WorkspaceId *string, PermissionCode *string, request *GetPermissionRequest) (_result *GetPermissionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetPermissionResponse{}
	_body, _err := client.GetPermissionWithOptions(WorkspaceId, PermissionCode, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetServiceTemplateWithOptions(ServiceTemplateId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetServiceTemplateResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetServiceTemplate"),
		Version:     tea.String("2021-02-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/servicetemplates/" + tea.StringValue(openapiutil.GetEncodeParam(ServiceTemplateId))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetServiceTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetServiceTemplate(ServiceTemplateId *string) (_result *GetServiceTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetServiceTemplateResponse{}
	_body, _err := client.GetServiceTemplateWithOptions(ServiceTemplateId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetWorkspaceWithOptions(WorkspaceId *string, request *GetWorkspaceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetWorkspaceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Verbose)) {
		query["Verbose"] = request.Verbose
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetWorkspace"),
		Version:     tea.String("2021-02-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/workspaces/" + tea.StringValue(openapiutil.GetEncodeParam(WorkspaceId))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetWorkspaceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetWorkspace(WorkspaceId *string, request *GetWorkspaceRequest) (_result *GetWorkspaceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetWorkspaceResponse{}
	_body, _err := client.GetWorkspaceWithOptions(WorkspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListCodeSourcesWithOptions(request *ListCodeSourcesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListCodeSourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DisplayName)) {
		query["DisplayName"] = request.DisplayName
	}

	if !tea.BoolValue(util.IsUnset(request.Order)) {
		query["Order"] = request.Order
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SortBy)) {
		query["SortBy"] = request.SortBy
	}

	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListCodeSources"),
		Version:     tea.String("2021-02-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/codesources"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListCodeSourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListCodeSources(request *ListCodeSourcesRequest) (_result *ListCodeSourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListCodeSourcesResponse{}
	_body, _err := client.ListCodeSourcesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListDatasetsWithOptions(request *ListDatasetsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListDatasetsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DataSourceTypes)) {
		query["DataSourceTypes"] = request.DataSourceTypes
	}

	if !tea.BoolValue(util.IsUnset(request.DataTypes)) {
		query["DataTypes"] = request.DataTypes
	}

	if !tea.BoolValue(util.IsUnset(request.Label)) {
		query["Label"] = request.Label
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Order)) {
		query["Order"] = request.Order
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Properties)) {
		query["Properties"] = request.Properties
	}

	if !tea.BoolValue(util.IsUnset(request.Provider)) {
		query["Provider"] = request.Provider
	}

	if !tea.BoolValue(util.IsUnset(request.SourceId)) {
		query["SourceId"] = request.SourceId
	}

	if !tea.BoolValue(util.IsUnset(request.SourceTypes)) {
		query["SourceTypes"] = request.SourceTypes
	}

	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListDatasets"),
		Version:     tea.String("2021-02-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/datasets"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListDatasetsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListDatasets(request *ListDatasetsRequest) (_result *ListDatasetsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListDatasetsResponse{}
	_body, _err := client.ListDatasetsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListImageLabelsWithOptions(request *ListImageLabelsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListImageLabelsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ImageId)) {
		query["ImageId"] = request.ImageId
	}

	if !tea.BoolValue(util.IsUnset(request.LabelFilter)) {
		query["LabelFilter"] = request.LabelFilter
	}

	if !tea.BoolValue(util.IsUnset(request.LabelKeys)) {
		query["LabelKeys"] = request.LabelKeys
	}

	if !tea.BoolValue(util.IsUnset(request.Region)) {
		query["Region"] = request.Region
	}

	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListImageLabels"),
		Version:     tea.String("2021-02-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/image/labels"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListImageLabelsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListImageLabels(request *ListImageLabelsRequest) (_result *ListImageLabelsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListImageLabelsResponse{}
	_body, _err := client.ListImageLabelsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListImagesWithOptions(request *ListImagesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListImagesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Accessibility)) {
		query["Accessibility"] = request.Accessibility
	}

	if !tea.BoolValue(util.IsUnset(request.Labels)) {
		query["Labels"] = request.Labels
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Order)) {
		query["Order"] = request.Order
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ParentUserId)) {
		query["ParentUserId"] = request.ParentUserId
	}

	if !tea.BoolValue(util.IsUnset(request.Query)) {
		query["Query"] = request.Query
	}

	if !tea.BoolValue(util.IsUnset(request.SortBy)) {
		query["SortBy"] = request.SortBy
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["UserId"] = request.UserId
	}

	if !tea.BoolValue(util.IsUnset(request.Verbose)) {
		query["Verbose"] = request.Verbose
	}

	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListImages"),
		Version:     tea.String("2021-02-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/images"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListImagesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListImages(request *ListImagesRequest) (_result *ListImagesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListImagesResponse{}
	_body, _err := client.ListImagesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListMembersWithOptions(WorkspaceId *string, request *ListMembersRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListMembersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MemberName)) {
		query["MemberName"] = request.MemberName
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Roles)) {
		query["Roles"] = request.Roles
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListMembers"),
		Version:     tea.String("2021-02-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/workspaces/" + tea.StringValue(openapiutil.GetEncodeParam(WorkspaceId)) + "/members"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListMembersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListMembers(WorkspaceId *string, request *ListMembersRequest) (_result *ListMembersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListMembersResponse{}
	_body, _err := client.ListMembersWithOptions(WorkspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListModelVersionsWithOptions(ModelId *string, request *ListModelVersionsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListModelVersionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApprovalStatus)) {
		query["ApprovalStatus"] = request.ApprovalStatus
	}

	if !tea.BoolValue(util.IsUnset(request.FormatType)) {
		query["FormatType"] = request.FormatType
	}

	if !tea.BoolValue(util.IsUnset(request.FrameworkType)) {
		query["FrameworkType"] = request.FrameworkType
	}

	if !tea.BoolValue(util.IsUnset(request.Label)) {
		query["Label"] = request.Label
	}

	if !tea.BoolValue(util.IsUnset(request.Order)) {
		query["Order"] = request.Order
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SortBy)) {
		query["SortBy"] = request.SortBy
	}

	if !tea.BoolValue(util.IsUnset(request.SourceId)) {
		query["SourceId"] = request.SourceId
	}

	if !tea.BoolValue(util.IsUnset(request.SourceType)) {
		query["SourceType"] = request.SourceType
	}

	if !tea.BoolValue(util.IsUnset(request.VersionName)) {
		query["VersionName"] = request.VersionName
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListModelVersions"),
		Version:     tea.String("2021-02-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/models/" + tea.StringValue(openapiutil.GetEncodeParam(ModelId)) + "/versions"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListModelVersionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListModelVersions(ModelId *string, request *ListModelVersionsRequest) (_result *ListModelVersionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListModelVersionsResponse{}
	_body, _err := client.ListModelVersionsWithOptions(ModelId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListModelsWithOptions(request *ListModelsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListModelsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Collections)) {
		query["Collections"] = request.Collections
	}

	if !tea.BoolValue(util.IsUnset(request.Domain)) {
		query["Domain"] = request.Domain
	}

	if !tea.BoolValue(util.IsUnset(request.Label)) {
		query["Label"] = request.Label
	}

	if !tea.BoolValue(util.IsUnset(request.ModelName)) {
		query["ModelName"] = request.ModelName
	}

	if !tea.BoolValue(util.IsUnset(request.ModelType)) {
		query["ModelType"] = request.ModelType
	}

	if !tea.BoolValue(util.IsUnset(request.Order)) {
		query["Order"] = request.Order
	}

	if !tea.BoolValue(util.IsUnset(request.Origin)) {
		query["Origin"] = request.Origin
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Provider)) {
		query["Provider"] = request.Provider
	}

	if !tea.BoolValue(util.IsUnset(request.Query)) {
		query["Query"] = request.Query
	}

	if !tea.BoolValue(util.IsUnset(request.SortBy)) {
		query["SortBy"] = request.SortBy
	}

	if !tea.BoolValue(util.IsUnset(request.Task)) {
		query["Task"] = request.Task
	}

	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListModels"),
		Version:     tea.String("2021-02-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/models"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListModelsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListModels(request *ListModelsRequest) (_result *ListModelsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListModelsResponse{}
	_body, _err := client.ListModelsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListPermissionsWithOptions(WorkspaceId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListPermissionsResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("ListPermissions"),
		Version:     tea.String("2021-02-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/workspaces/" + tea.StringValue(openapiutil.GetEncodeParam(WorkspaceId)) + "/permissions"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListPermissionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListPermissions(WorkspaceId *string) (_result *ListPermissionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListPermissionsResponse{}
	_body, _err := client.ListPermissionsWithOptions(WorkspaceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListProductsWithOptions(request *ListProductsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListProductsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ProductCodes)) {
		query["ProductCodes"] = request.ProductCodes
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceCodes)) {
		query["ServiceCodes"] = request.ServiceCodes
	}

	if !tea.BoolValue(util.IsUnset(request.Verbose)) {
		query["Verbose"] = request.Verbose
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListProducts"),
		Version:     tea.String("2021-02-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/products"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListProductsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListProducts(request *ListProductsRequest) (_result *ListProductsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListProductsResponse{}
	_body, _err := client.ListProductsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListQuotasWithOptions(request *ListQuotasRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListQuotasResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListQuotas"),
		Version:     tea.String("2021-02-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/quotas"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListQuotasResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListQuotas(request *ListQuotasRequest) (_result *ListQuotasResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListQuotasResponse{}
	_body, _err := client.ListQuotasWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListResourcesWithOptions(request *ListResourcesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GroupName)) {
		query["GroupName"] = request.GroupName
	}

	if !tea.BoolValue(util.IsUnset(request.Labels)) {
		query["Labels"] = request.Labels
	}

	if !tea.BoolValue(util.IsUnset(request.Option)) {
		query["Option"] = request.Option
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ProductTypes)) {
		query["ProductTypes"] = request.ProductTypes
	}

	if !tea.BoolValue(util.IsUnset(request.QuotaIds)) {
		query["QuotaIds"] = request.QuotaIds
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceName)) {
		query["ResourceName"] = request.ResourceName
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceTypes)) {
		query["ResourceTypes"] = request.ResourceTypes
	}

	if !tea.BoolValue(util.IsUnset(request.Verbose)) {
		query["Verbose"] = request.Verbose
	}

	if !tea.BoolValue(util.IsUnset(request.VerboseFields)) {
		query["VerboseFields"] = request.VerboseFields
	}

	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListResources"),
		Version:     tea.String("2021-02-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/resources"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListResources(request *ListResourcesRequest) (_result *ListResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListResourcesResponse{}
	_body, _err := client.ListResourcesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListServiceTemplatesWithOptions(request *ListServiceTemplatesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListServiceTemplatesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Label)) {
		query["Label"] = request.Label
	}

	if !tea.BoolValue(util.IsUnset(request.Order)) {
		query["Order"] = request.Order
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Provider)) {
		query["Provider"] = request.Provider
	}

	if !tea.BoolValue(util.IsUnset(request.Query)) {
		query["Query"] = request.Query
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceTemplateName)) {
		query["ServiceTemplateName"] = request.ServiceTemplateName
	}

	if !tea.BoolValue(util.IsUnset(request.SortBy)) {
		query["SortBy"] = request.SortBy
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListServiceTemplates"),
		Version:     tea.String("2021-02-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/servicetemplates"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListServiceTemplatesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListServiceTemplates(request *ListServiceTemplatesRequest) (_result *ListServiceTemplatesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListServiceTemplatesResponse{}
	_body, _err := client.ListServiceTemplatesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListWorkspaceUsersWithOptions(WorkspaceId *string, request *ListWorkspaceUsersRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListWorkspaceUsersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UserName)) {
		query["UserName"] = request.UserName
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListWorkspaceUsers"),
		Version:     tea.String("2021-02-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/workspaces/" + tea.StringValue(openapiutil.GetEncodeParam(WorkspaceId)) + "/users"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListWorkspaceUsersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListWorkspaceUsers(WorkspaceId *string, request *ListWorkspaceUsersRequest) (_result *ListWorkspaceUsersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListWorkspaceUsersResponse{}
	_body, _err := client.ListWorkspaceUsersWithOptions(WorkspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListWorkspacesWithOptions(request *ListWorkspacesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListWorkspacesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Fields)) {
		query["Fields"] = request.Fields
	}

	if !tea.BoolValue(util.IsUnset(request.ModuleList)) {
		query["ModuleList"] = request.ModuleList
	}

	if !tea.BoolValue(util.IsUnset(request.Option)) {
		query["Option"] = request.Option
	}

	if !tea.BoolValue(util.IsUnset(request.Order)) {
		query["Order"] = request.Order
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SortBy)) {
		query["SortBy"] = request.SortBy
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.Verbose)) {
		query["Verbose"] = request.Verbose
	}

	if !tea.BoolValue(util.IsUnset(request.WorkspaceIds)) {
		query["WorkspaceIds"] = request.WorkspaceIds
	}

	if !tea.BoolValue(util.IsUnset(request.WorkspaceName)) {
		query["WorkspaceName"] = request.WorkspaceName
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListWorkspaces"),
		Version:     tea.String("2021-02-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/workspaces"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListWorkspacesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListWorkspaces(request *ListWorkspacesRequest) (_result *ListWorkspacesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListWorkspacesResponse{}
	_body, _err := client.ListWorkspacesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PublishCodeSourceWithOptions(CodeSourceId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *PublishCodeSourceResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("PublishCodeSource"),
		Version:     tea.String("2021-02-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/codesources/" + tea.StringValue(openapiutil.GetEncodeParam(CodeSourceId)) + "/publish"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &PublishCodeSourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PublishCodeSource(CodeSourceId *string) (_result *PublishCodeSourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &PublishCodeSourceResponse{}
	_body, _err := client.PublishCodeSourceWithOptions(CodeSourceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PublishDatasetWithOptions(DatasetId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *PublishDatasetResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("PublishDataset"),
		Version:     tea.String("2021-02-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/datasets/" + tea.StringValue(openapiutil.GetEncodeParam(DatasetId)) + "/publish"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &PublishDatasetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PublishDataset(DatasetId *string) (_result *PublishDatasetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &PublishDatasetResponse{}
	_body, _err := client.PublishDatasetWithOptions(DatasetId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PublishImageWithOptions(ImageId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *PublishImageResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("PublishImage"),
		Version:     tea.String("2021-02-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/images/" + tea.StringValue(openapiutil.GetEncodeParam(ImageId)) + "/publish"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &PublishImageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PublishImage(ImageId *string) (_result *PublishImageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &PublishImageResponse{}
	_body, _err := client.PublishImageWithOptions(ImageId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RemoveImageWithOptions(ImageId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *RemoveImageResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("RemoveImage"),
		Version:     tea.String("2021-02-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/images/" + tea.StringValue(openapiutil.GetEncodeParam(ImageId))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &RemoveImageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RemoveImage(ImageId *string) (_result *RemoveImageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RemoveImageResponse{}
	_body, _err := client.RemoveImageWithOptions(ImageId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RemoveImageLabelsWithOptions(ImageId *string, LabelKey *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *RemoveImageLabelsResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("RemoveImageLabels"),
		Version:     tea.String("2021-02-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/images/" + tea.StringValue(openapiutil.GetEncodeParam(ImageId)) + "/labels/" + tea.StringValue(openapiutil.GetEncodeParam(LabelKey))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &RemoveImageLabelsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RemoveImageLabels(ImageId *string, LabelKey *string) (_result *RemoveImageLabelsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RemoveImageLabelsResponse{}
	_body, _err := client.RemoveImageLabelsWithOptions(ImageId, LabelKey, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RemoveMemberRoleWithOptions(WorkspaceId *string, MemberId *string, RoleName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *RemoveMemberRoleResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("RemoveMemberRole"),
		Version:     tea.String("2021-02-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/workspaces/" + tea.StringValue(openapiutil.GetEncodeParam(WorkspaceId)) + "/members/" + tea.StringValue(openapiutil.GetEncodeParam(MemberId)) + "/roles/" + tea.StringValue(openapiutil.GetEncodeParam(RoleName))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &RemoveMemberRoleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RemoveMemberRole(WorkspaceId *string, MemberId *string, RoleName *string) (_result *RemoveMemberRoleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RemoveMemberRoleResponse{}
	_body, _err := client.RemoveMemberRoleWithOptions(WorkspaceId, MemberId, RoleName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateDatasetWithOptions(DatasetId *string, request *UpdateDatasetRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateDatasetResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Options)) {
		body["Options"] = request.Options
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateDataset"),
		Version:     tea.String("2021-02-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/datasets/" + tea.StringValue(openapiutil.GetEncodeParam(DatasetId))),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateDatasetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateDataset(DatasetId *string, request *UpdateDatasetRequest) (_result *UpdateDatasetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateDatasetResponse{}
	_body, _err := client.UpdateDatasetWithOptions(DatasetId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateDefaultWorkspaceWithOptions(request *UpdateDefaultWorkspaceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateDefaultWorkspaceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateDefaultWorkspace"),
		Version:     tea.String("2021-02-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/defaultWorkspaces"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateDefaultWorkspaceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateDefaultWorkspace(request *UpdateDefaultWorkspaceRequest) (_result *UpdateDefaultWorkspaceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateDefaultWorkspaceResponse{}
	_body, _err := client.UpdateDefaultWorkspaceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateModelWithOptions(ModelId *string, request *UpdateModelRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateModelResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Accessibility)) {
		body["Accessibility"] = request.Accessibility
	}

	if !tea.BoolValue(util.IsUnset(request.Domain)) {
		body["Domain"] = request.Domain
	}

	if !tea.BoolValue(util.IsUnset(request.ExtraInfo)) {
		body["ExtraInfo"] = request.ExtraInfo
	}

	if !tea.BoolValue(util.IsUnset(request.ModelDescription)) {
		body["ModelDescription"] = request.ModelDescription
	}

	if !tea.BoolValue(util.IsUnset(request.ModelDoc)) {
		body["ModelDoc"] = request.ModelDoc
	}

	if !tea.BoolValue(util.IsUnset(request.ModelName)) {
		body["ModelName"] = request.ModelName
	}

	if !tea.BoolValue(util.IsUnset(request.ModelType)) {
		body["ModelType"] = request.ModelType
	}

	if !tea.BoolValue(util.IsUnset(request.OrderNumber)) {
		body["OrderNumber"] = request.OrderNumber
	}

	if !tea.BoolValue(util.IsUnset(request.Origin)) {
		body["Origin"] = request.Origin
	}

	if !tea.BoolValue(util.IsUnset(request.Task)) {
		body["Task"] = request.Task
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateModel"),
		Version:     tea.String("2021-02-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/models/" + tea.StringValue(openapiutil.GetEncodeParam(ModelId))),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateModelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateModel(ModelId *string, request *UpdateModelRequest) (_result *UpdateModelResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateModelResponse{}
	_body, _err := client.UpdateModelWithOptions(ModelId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateModelVersionWithOptions(ModelId *string, VersionName *string, request *UpdateModelVersionRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateModelVersionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApprovalStatus)) {
		body["ApprovalStatus"] = request.ApprovalStatus
	}

	if !tea.BoolValue(util.IsUnset(request.EvaluationSpec)) {
		body["EvaluationSpec"] = request.EvaluationSpec
	}

	if !tea.BoolValue(util.IsUnset(request.ExtraInfo)) {
		body["ExtraInfo"] = request.ExtraInfo
	}

	if !tea.BoolValue(util.IsUnset(request.InferenceSpec)) {
		body["InferenceSpec"] = request.InferenceSpec
	}

	if !tea.BoolValue(util.IsUnset(request.Metrics)) {
		body["Metrics"] = request.Metrics
	}

	if !tea.BoolValue(util.IsUnset(request.Options)) {
		body["Options"] = request.Options
	}

	if !tea.BoolValue(util.IsUnset(request.SourceId)) {
		body["SourceId"] = request.SourceId
	}

	if !tea.BoolValue(util.IsUnset(request.SourceType)) {
		body["SourceType"] = request.SourceType
	}

	if !tea.BoolValue(util.IsUnset(request.TrainingSpec)) {
		body["TrainingSpec"] = request.TrainingSpec
	}

	if !tea.BoolValue(util.IsUnset(request.VersionDescription)) {
		body["VersionDescription"] = request.VersionDescription
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateModelVersion"),
		Version:     tea.String("2021-02-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/models/" + tea.StringValue(openapiutil.GetEncodeParam(ModelId)) + "/versions/" + tea.StringValue(openapiutil.GetEncodeParam(VersionName))),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateModelVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateModelVersion(ModelId *string, VersionName *string, request *UpdateModelVersionRequest) (_result *UpdateModelVersionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateModelVersionResponse{}
	_body, _err := client.UpdateModelVersionWithOptions(ModelId, VersionName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateWorkspaceWithOptions(WorkspaceId *string, request *UpdateWorkspaceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateWorkspaceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.DisplayName)) {
		body["DisplayName"] = request.DisplayName
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateWorkspace"),
		Version:     tea.String("2021-02-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/workspaces/" + tea.StringValue(openapiutil.GetEncodeParam(WorkspaceId))),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateWorkspaceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateWorkspace(WorkspaceId *string, request *UpdateWorkspaceRequest) (_result *UpdateWorkspaceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateWorkspaceResponse{}
	_body, _err := client.UpdateWorkspaceWithOptions(WorkspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateWorkspaceResourceWithOptions(WorkspaceId *string, request *UpdateWorkspaceResourceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateWorkspaceResourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GroupName)) {
		body["GroupName"] = request.GroupName
	}

	if !tea.BoolValue(util.IsUnset(request.IsDefault)) {
		body["IsDefault"] = request.IsDefault
	}

	if !tea.BoolValue(util.IsUnset(request.Labels)) {
		body["Labels"] = request.Labels
	}

	if !tea.BoolValue(util.IsUnset(request.ProductType)) {
		body["ProductType"] = request.ProductType
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceIds)) {
		body["ResourceIds"] = request.ResourceIds
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		body["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.Spec)) {
		body["Spec"] = request.Spec
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateWorkspaceResource"),
		Version:     tea.String("2021-02-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/workspaces/" + tea.StringValue(openapiutil.GetEncodeParam(WorkspaceId)) + "/resources"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateWorkspaceResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateWorkspaceResource(WorkspaceId *string, request *UpdateWorkspaceResourceRequest) (_result *UpdateWorkspaceResourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateWorkspaceResourceResponse{}
	_body, _err := client.UpdateWorkspaceResourceWithOptions(WorkspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
