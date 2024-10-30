// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type CodeSourceItem struct {
	// example:
	//
	// PUBLIC
	Accessibility *string `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	// example:
	//
	// master
	CodeBranch *string `json:"CodeBranch,omitempty" xml:"CodeBranch,omitempty"`
	// example:
	//
	// 44da109b59f8596152987eaa8f3b2487bb72ea63
	CodeCommit *string `json:"CodeCommit,omitempty" xml:"CodeCommit,omitempty"`
	// example:
	//
	// https://code.aliyun.com/pai-dlc/examples.git
	CodeRepo            *string `json:"CodeRepo,omitempty" xml:"CodeRepo,omitempty"`
	CodeRepoAccessToken *string `json:"CodeRepoAccessToken,omitempty" xml:"CodeRepoAccessToken,omitempty"`
	// example:
	//
	// user
	CodeRepoUserName *string `json:"CodeRepoUserName,omitempty" xml:"CodeRepoUserName,omitempty"`
	// example:
	//
	// code-20210111103721-85qz78ia96lu
	CodeSourceId *string `json:"CodeSourceId,omitempty" xml:"CodeSourceId,omitempty"`
	// example:
	//
	// code source of dlc examples
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// MyCodeSourceName1
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// example:
	//
	// 2021-01-18T12:52:15Z
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// example:
	//
	// 2021-01-18T12:52:15Z
	GmtModifyTime *string `json:"GmtModifyTime,omitempty" xml:"GmtModifyTime,omitempty"`
	// example:
	//
	// /root/code/
	MountPath *string `json:"MountPath,omitempty" xml:"MountPath,omitempty"`
	// example:
	//
	// 1157290171663117
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// example:
	//
	// 1234
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
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
	// example:
	//
	// AI4D
	CollectionName *string `json:"CollectionName,omitempty" xml:"CollectionName,omitempty"`
	// example:
	//
	// 2021-01-21T17:12:35Z
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// example:
	//
	// 2021-01-21T17:12:35Z
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// example:
	//
	// 155770209******
	OwnerId *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// 155770209******
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
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
	// example:
	//
	// PRIVATE PUBLIC
	Accessibility *string `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	// example:
	//
	// OSS URL
	DataSourceType *string `json:"DataSourceType,omitempty" xml:"DataSourceType,omitempty"`
	// example:
	//
	// COMMON PIC TEXT VIDEO AUDIO
	DataType *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	// example:
	//
	// d-c0h44g3wlwkj8o4348
	DatasetId *string `json:"DatasetId,omitempty" xml:"DatasetId,omitempty"`
	// example:
	//
	// Animal images.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 2021-01-30T12:51:33.028Z
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// example:
	//
	// 2021-01-30T12:51:33.028Z
	GmtModifiedTime *string         `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	Labels          []*Label        `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	LatestVersion   *DatasetVersion `json:"LatestVersion,omitempty" xml:"LatestVersion,omitempty"`
	// example:
	//
	// AnimalDataset
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// jsonstring
	Options *string `json:"Options,omitempty" xml:"Options,omitempty"`
	// example:
	//
	// 1004110000006048
	OwnerId *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// FILE DIRECTORY TABULAR
	Property *string `json:"Property,omitempty" xml:"Property,omitempty"`
	// example:
	//
	// Ecs
	ProviderType *string `json:"ProviderType,omitempty" xml:"ProviderType,omitempty"`
	// example:
	//
	// d-bvfasdf4wxxj8o411
	SourceDatasetId *string `json:"SourceDatasetId,omitempty" xml:"SourceDatasetId,omitempty"`
	// example:
	//
	// v2
	SourceDatasetVersion *string `json:"SourceDatasetVersion,omitempty" xml:"SourceDatasetVersion,omitempty"`
	// example:
	//
	// Source Id
	SourceId *string `json:"SourceId,omitempty" xml:"SourceId,omitempty"`
	// example:
	//
	// USER ITAG  PAI_PUBLIC_DATASET
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// example:
	//
	// text-classification
	TagTemplateType *string `json:"TagTemplateType,omitempty" xml:"TagTemplateType,omitempty"`
	// example:
	//
	// oss://xxx
	Uri *string `json:"Uri,omitempty" xml:"Uri,omitempty"`
	// example:
	//
	// 2004110000006048
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// example:
	//
	// Workspace Id
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
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

func (s *Dataset) SetLatestVersion(v *DatasetVersion) *Dataset {
	s.LatestVersion = v
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

func (s *Dataset) SetSourceDatasetId(v string) *Dataset {
	s.SourceDatasetId = &v
	return s
}

func (s *Dataset) SetSourceDatasetVersion(v string) *Dataset {
	s.SourceDatasetVersion = &v
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

func (s *Dataset) SetTagTemplateType(v string) *Dataset {
	s.TagTemplateType = &v
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

type DatasetVersion struct {
	DataCount *int64 `json:"DataCount,omitempty" xml:"DataCount,omitempty"`
	DataSize  *int64 `json:"DataSize,omitempty" xml:"DataSize,omitempty"`
	// example:
	//
	// OSS
	DataSourceType  *string  `json:"DataSourceType,omitempty" xml:"DataSourceType,omitempty"`
	Description     *string  `json:"Description,omitempty" xml:"Description,omitempty"`
	GmtCreateTime   *string  `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	GmtModifiedTime *string  `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	Labels          []*Label `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	Options         *string  `json:"Options,omitempty" xml:"Options,omitempty"`
	// example:
	//
	// FILE
	Property   *string `json:"Property,omitempty" xml:"Property,omitempty"`
	SourceId   *string `json:"SourceId,omitempty" xml:"SourceId,omitempty"`
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// example:
	//
	// OSS://xxx
	Uri *string `json:"Uri,omitempty" xml:"Uri,omitempty"`
	// example:
	//
	// v1
	VersionName *string `json:"VersionName,omitempty" xml:"VersionName,omitempty"`
}

func (s DatasetVersion) String() string {
	return tea.Prettify(s)
}

func (s DatasetVersion) GoString() string {
	return s.String()
}

func (s *DatasetVersion) SetDataCount(v int64) *DatasetVersion {
	s.DataCount = &v
	return s
}

func (s *DatasetVersion) SetDataSize(v int64) *DatasetVersion {
	s.DataSize = &v
	return s
}

func (s *DatasetVersion) SetDataSourceType(v string) *DatasetVersion {
	s.DataSourceType = &v
	return s
}

func (s *DatasetVersion) SetDescription(v string) *DatasetVersion {
	s.Description = &v
	return s
}

func (s *DatasetVersion) SetGmtCreateTime(v string) *DatasetVersion {
	s.GmtCreateTime = &v
	return s
}

func (s *DatasetVersion) SetGmtModifiedTime(v string) *DatasetVersion {
	s.GmtModifiedTime = &v
	return s
}

func (s *DatasetVersion) SetLabels(v []*Label) *DatasetVersion {
	s.Labels = v
	return s
}

func (s *DatasetVersion) SetOptions(v string) *DatasetVersion {
	s.Options = &v
	return s
}

func (s *DatasetVersion) SetProperty(v string) *DatasetVersion {
	s.Property = &v
	return s
}

func (s *DatasetVersion) SetSourceId(v string) *DatasetVersion {
	s.SourceId = &v
	return s
}

func (s *DatasetVersion) SetSourceType(v string) *DatasetVersion {
	s.SourceType = &v
	return s
}

func (s *DatasetVersion) SetUri(v string) *DatasetVersion {
	s.Uri = &v
	return s
}

func (s *DatasetVersion) SetVersionName(v string) *DatasetVersion {
	s.VersionName = &v
	return s
}

type Experiment struct {
	Accessibility     *string            `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	ArtifactUri       *string            `json:"ArtifactUri,omitempty" xml:"ArtifactUri,omitempty"`
	ExperimentId      *string            `json:"ExperimentId,omitempty" xml:"ExperimentId,omitempty"`
	GmtCreateTime     *string            `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	GmtModifiedTime   *string            `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	Labels            []*ExperimentLabel `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	LatestRun         *Run               `json:"LatestRun,omitempty" xml:"LatestRun,omitempty"`
	Name              *string            `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerId           *string            `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RequestId         *string            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TensorboardLogUri *string            `json:"TensorboardLogUri,omitempty" xml:"TensorboardLogUri,omitempty"`
	UserId            *string            `json:"UserId,omitempty" xml:"UserId,omitempty"`
	WorkspaceId       *string            `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s Experiment) String() string {
	return tea.Prettify(s)
}

func (s Experiment) GoString() string {
	return s.String()
}

func (s *Experiment) SetAccessibility(v string) *Experiment {
	s.Accessibility = &v
	return s
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

func (s *Experiment) SetLabels(v []*ExperimentLabel) *Experiment {
	s.Labels = v
	return s
}

func (s *Experiment) SetLatestRun(v *Run) *Experiment {
	s.LatestRun = v
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

func (s *Experiment) SetRequestId(v string) *Experiment {
	s.RequestId = &v
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
	// example:
	//
	// exp-890waerw09a0f
	ExperimentId *string `json:"ExperimentId,omitempty" xml:"ExperimentId,omitempty"`
	// example:
	//
	// 2023-12-27T03:30:04Z
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// example:
	//
	// 2023-12-27T03:30:04Z
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// example:
	//
	// key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// value
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
	// example:
	//
	// key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// value
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
	// example:
	//
	// PUBLIC
	Accessibility *string `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	// example:
	//
	// cv
	Domain    *string                `json:"Domain,omitempty" xml:"Domain,omitempty"`
	ExtraInfo map[string]interface{} `json:"ExtraInfo,omitempty" xml:"ExtraInfo,omitempty"`
	// example:
	//
	// 2021-01-21T17:12:35Z
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// example:
	//
	// 2021-01-21T17:12:35Z
	GmtModifiedTime  *string       `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	Labels           []*Label      `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	LatestVersion    *ModelVersion `json:"LatestVersion,omitempty" xml:"LatestVersion,omitempty"`
	ModelDescription *string       `json:"ModelDescription,omitempty" xml:"ModelDescription,omitempty"`
	// example:
	//
	// https://***.md
	ModelDoc *string `json:"ModelDoc,omitempty" xml:"ModelDoc,omitempty"`
	// example:
	//
	// model-1123*****
	ModelId   *string `json:"ModelId,omitempty" xml:"ModelId,omitempty"`
	ModelName *string `json:"ModelName,omitempty" xml:"ModelName,omitempty"`
	// example:
	//
	// Checkpoint
	ModelType *string `json:"ModelType,omitempty" xml:"ModelType,omitempty"`
	// example:
	//
	// 101
	OrderNumber *int64 `json:"OrderNumber,omitempty" xml:"OrderNumber,omitempty"`
	// example:
	//
	// ModelScope
	Origin *string `json:"Origin,omitempty" xml:"Origin,omitempty"`
	// example:
	//
	// 1557702098******
	OwnerId *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// pai
	Provider *string `json:"Provider,omitempty" xml:"Provider,omitempty"`
	// example:
	//
	// text-classifiaction
	Task *string `json:"Task,omitempty" xml:"Task,omitempty"`
	// example:
	//
	// 1557702098******
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// example:
	//
	// 234**
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
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
	ApprovalStatus  *string                `json:"ApprovalStatus,omitempty" xml:"ApprovalStatus,omitempty"`
	CompressionSpec map[string]interface{} `json:"CompressionSpec,omitempty" xml:"CompressionSpec,omitempty"`
	EvaluationSpec  map[string]interface{} `json:"EvaluationSpec,omitempty" xml:"EvaluationSpec,omitempty"`
	ExtraInfo       map[string]interface{} `json:"ExtraInfo,omitempty" xml:"ExtraInfo,omitempty"`
	// example:
	//
	// SavedModel
	FormatType *string `json:"FormatType,omitempty" xml:"FormatType,omitempty"`
	// example:
	//
	// TensorFlow
	FrameworkType *string `json:"FrameworkType,omitempty" xml:"FrameworkType,omitempty"`
	// example:
	//
	// 2021-01-21T17:12:35Z
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// example:
	//
	// 2021-01-21T17:12:35Z
	GmtModifiedTime *string                `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	InferenceSpec   map[string]interface{} `json:"InferenceSpec,omitempty" xml:"InferenceSpec,omitempty"`
	Labels          []*Label               `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	Metrics         map[string]interface{} `json:"Metrics,omitempty" xml:"Metrics,omitempty"`
	Options         *string                `json:"Options,omitempty" xml:"Options,omitempty"`
	// example:
	//
	// 155770209******
	OwnerId      *string                `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SourceId     *string                `json:"SourceId,omitempty" xml:"SourceId,omitempty"`
	SourceType   *string                `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	TrainingSpec map[string]interface{} `json:"TrainingSpec,omitempty" xml:"TrainingSpec,omitempty"`
	// example:
	//
	// oss://bucket/path-to-model
	Uri *string `json:"Uri,omitempty" xml:"Uri,omitempty"`
	// example:
	//
	// 155770209******
	UserId             *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	VersionDescription *string `json:"VersionDescription,omitempty" xml:"VersionDescription,omitempty"`
	// example:
	//
	// 1.0.0
	VersionName *string `json:"VersionName,omitempty" xml:"VersionName,omitempty"`
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

func (s *ModelVersion) SetCompressionSpec(v map[string]interface{}) *ModelVersion {
	s.CompressionSpec = v
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

type Run struct {
	Accessibility   *string      `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	ExperimentId    *string      `json:"ExperimentId,omitempty" xml:"ExperimentId,omitempty"`
	GmtCreateTime   *string      `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	GmtModifiedTime *string      `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	Labels          []*RunLabel  `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	Metrics         []*RunMetric `json:"Metrics,omitempty" xml:"Metrics,omitempty" type:"Repeated"`
	Name            *string      `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerId         *string      `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	Params          []*RunParam  `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	RequestId       *string      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RunId           *string      `json:"RunId,omitempty" xml:"RunId,omitempty"`
	SourceId        *string      `json:"SourceId,omitempty" xml:"SourceId,omitempty"`
	SourceType      *string      `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	UserId          *string      `json:"UserId,omitempty" xml:"UserId,omitempty"`
	WorkspaceId     *string      `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s Run) String() string {
	return tea.Prettify(s)
}

func (s Run) GoString() string {
	return s.String()
}

func (s *Run) SetAccessibility(v string) *Run {
	s.Accessibility = &v
	return s
}

func (s *Run) SetExperimentId(v string) *Run {
	s.ExperimentId = &v
	return s
}

func (s *Run) SetGmtCreateTime(v string) *Run {
	s.GmtCreateTime = &v
	return s
}

func (s *Run) SetGmtModifiedTime(v string) *Run {
	s.GmtModifiedTime = &v
	return s
}

func (s *Run) SetLabels(v []*RunLabel) *Run {
	s.Labels = v
	return s
}

func (s *Run) SetMetrics(v []*RunMetric) *Run {
	s.Metrics = v
	return s
}

func (s *Run) SetName(v string) *Run {
	s.Name = &v
	return s
}

func (s *Run) SetOwnerId(v string) *Run {
	s.OwnerId = &v
	return s
}

func (s *Run) SetParams(v []*RunParam) *Run {
	s.Params = v
	return s
}

func (s *Run) SetRequestId(v string) *Run {
	s.RequestId = &v
	return s
}

func (s *Run) SetRunId(v string) *Run {
	s.RunId = &v
	return s
}

func (s *Run) SetSourceId(v string) *Run {
	s.SourceId = &v
	return s
}

func (s *Run) SetSourceType(v string) *Run {
	s.SourceType = &v
	return s
}

func (s *Run) SetUserId(v string) *Run {
	s.UserId = &v
	return s
}

func (s *Run) SetWorkspaceId(v string) *Run {
	s.WorkspaceId = &v
	return s
}

type RunLabel struct {
	GmtCreateTime   *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// This parameter is required.
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	RunId *string `json:"RunId,omitempty" xml:"RunId,omitempty"`
	// This parameter is required.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s RunLabel) String() string {
	return tea.Prettify(s)
}

func (s RunLabel) GoString() string {
	return s.String()
}

func (s *RunLabel) SetGmtCreateTime(v string) *RunLabel {
	s.GmtCreateTime = &v
	return s
}

func (s *RunLabel) SetGmtModifiedTime(v string) *RunLabel {
	s.GmtModifiedTime = &v
	return s
}

func (s *RunLabel) SetKey(v string) *RunLabel {
	s.Key = &v
	return s
}

func (s *RunLabel) SetRunId(v string) *RunLabel {
	s.RunId = &v
	return s
}

func (s *RunLabel) SetValue(v string) *RunLabel {
	s.Value = &v
	return s
}

type RunMetric struct {
	// This parameter is required.
	Key       *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Step      *int64  `json:"Step,omitempty" xml:"Step,omitempty"`
	Timestamp *int64  `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	// This parameter is required.
	Value *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s RunMetric) String() string {
	return tea.Prettify(s)
}

func (s RunMetric) GoString() string {
	return s.String()
}

func (s *RunMetric) SetKey(v string) *RunMetric {
	s.Key = &v
	return s
}

func (s *RunMetric) SetStep(v int64) *RunMetric {
	s.Step = &v
	return s
}

func (s *RunMetric) SetTimestamp(v int64) *RunMetric {
	s.Timestamp = &v
	return s
}

func (s *RunMetric) SetValue(v float32) *RunMetric {
	s.Value = &v
	return s
}

type RunParam struct {
	// This parameter is required.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// This parameter is required.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s RunParam) String() string {
	return tea.Prettify(s)
}

func (s RunParam) GoString() string {
	return s.String()
}

func (s *RunParam) SetKey(v string) *RunParam {
	s.Key = &v
	return s
}

func (s *RunParam) SetValue(v string) *RunParam {
	s.Value = &v
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
	// example:
	//
	// 2023-12-27T03:30:04Z
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// example:
	//
	// 2023-12-27T03:30:04Z
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// example:
	//
	// key
	Key     *string `json:"Key,omitempty" xml:"Key,omitempty"`
	TrialId *string `json:"TrialId,omitempty" xml:"TrialId,omitempty"`
	// example:
	//
	// value
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
	// example:
	//
	// PUBLIC
	Accessibility *string `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	Description   *string `json:"Description,omitempty" xml:"Description,omitempty"`
	ImageId       *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// registry.cn-hangzhou.aliyuncs.com/pai-compression/nlp:gpu
	ImageUri *string                  `json:"ImageUri,omitempty" xml:"ImageUri,omitempty"`
	Labels   []*AddImageRequestLabels `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// nlp-compression
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Size *int32  `json:"Size,omitempty" xml:"Size,omitempty"`
	// example:
	//
	// 15******45
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
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
	// example:
	//
	// system.chipType
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// GPU
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
	// example:
	//
	// image-4c62******53uor
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// example:
	//
	// 5A14FA81-DD4E-******-6343FE44B941
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
	// This parameter is required.
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
	// example:
	//
	// system.chipType
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// GPU
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
	// example:
	//
	// 5A14FA81-DD4E-******-6343FE44B941
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
	// example:
	//
	// 5A14FA81-DD4E-******-6343FE44B941
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
	// example:
	//
	// PRIVATE
	Accessibility *string `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	// example:
	//
	// master
	CodeBranch *string `json:"CodeBranch,omitempty" xml:"CodeBranch,omitempty"`
	// example:
	//
	// https://code.aliyun.com/******
	CodeRepo *string `json:"CodeRepo,omitempty" xml:"CodeRepo,omitempty"`
	// example:
	//
	// ***
	CodeRepoAccessToken *string `json:"CodeRepoAccessToken,omitempty" xml:"CodeRepoAccessToken,omitempty"`
	// example:
	//
	// use***
	CodeRepoUserName *string `json:"CodeRepoUserName,omitempty" xml:"CodeRepoUserName,omitempty"`
	// example:
	//
	// code source of dlc examples
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// MyCodeSource1
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// example:
	//
	// /root/code/code-source-1
	MountPath *string `json:"MountPath,omitempty" xml:"MountPath,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1234
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
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
	// example:
	//
	// code-20********
	CodeSourceId *string `json:"CodeSourceId,omitempty" xml:"CodeSourceId,omitempty"`
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3**********
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// example:
	//
	// PRIVATE
	Accessibility *string `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	DataCount     *int64  `json:"DataCount,omitempty" xml:"DataCount,omitempty"`
	DataSize      *int64  `json:"DataSize,omitempty" xml:"DataSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// NAS
	DataSourceType *string `json:"DataSourceType,omitempty" xml:"DataSourceType,omitempty"`
	// example:
	//
	// COMMON
	DataType    *string  `json:"DataType,omitempty" xml:"DataType,omitempty"`
	Description *string  `json:"Description,omitempty" xml:"Description,omitempty"`
	Labels      []*Label `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// myName
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// {
	//
	//   "mountPath": "/mnt/data/"
	//
	// }
	Options *string `json:"Options,omitempty" xml:"Options,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// DIRECTORY
	Property *string `json:"Property,omitempty" xml:"Property,omitempty"`
	Provider *string `json:"Provider,omitempty" xml:"Provider,omitempty"`
	// example:
	//
	// Ecs
	ProviderType         *string `json:"ProviderType,omitempty" xml:"ProviderType,omitempty"`
	SourceDatasetId      *string `json:"SourceDatasetId,omitempty" xml:"SourceDatasetId,omitempty"`
	SourceDatasetVersion *string `json:"SourceDatasetVersion,omitempty" xml:"SourceDatasetVersion,omitempty"`
	// example:
	//
	// jdnhf***fnrimv
	SourceId *string `json:"SourceId,omitempty" xml:"SourceId,omitempty"`
	// example:
	//
	// USER
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// nas://09f****f2.cn-hangzhou/
	Uri *string `json:"Uri,omitempty" xml:"Uri,omitempty"`
	// example:
	//
	// 29884000000186970
	UserId             *string  `json:"UserId,omitempty" xml:"UserId,omitempty"`
	VersionDescription *string  `json:"VersionDescription,omitempty" xml:"VersionDescription,omitempty"`
	VersionLabels      []*Label `json:"VersionLabels,omitempty" xml:"VersionLabels,omitempty" type:"Repeated"`
	// example:
	//
	// 478**
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
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

func (s *CreateDatasetRequest) SetDataCount(v int64) *CreateDatasetRequest {
	s.DataCount = &v
	return s
}

func (s *CreateDatasetRequest) SetDataSize(v int64) *CreateDatasetRequest {
	s.DataSize = &v
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

func (s *CreateDatasetRequest) SetSourceDatasetId(v string) *CreateDatasetRequest {
	s.SourceDatasetId = &v
	return s
}

func (s *CreateDatasetRequest) SetSourceDatasetVersion(v string) *CreateDatasetRequest {
	s.SourceDatasetVersion = &v
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

func (s *CreateDatasetRequest) SetUserId(v string) *CreateDatasetRequest {
	s.UserId = &v
	return s
}

func (s *CreateDatasetRequest) SetVersionDescription(v string) *CreateDatasetRequest {
	s.VersionDescription = &v
	return s
}

func (s *CreateDatasetRequest) SetVersionLabels(v []*Label) *CreateDatasetRequest {
	s.VersionLabels = v
	return s
}

func (s *CreateDatasetRequest) SetWorkspaceId(v string) *CreateDatasetRequest {
	s.WorkspaceId = &v
	return s
}

type CreateDatasetResponseBody struct {
	// example:
	//
	// d-rbvg5*****jhc9ks92
	DatasetId *string `json:"DatasetId,omitempty" xml:"DatasetId,omitempty"`
	// example:
	//
	// B2C51F93-1C07-5477-9705-5FDB****F19F
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
	// example:
	//
	// A083731B-4973-54D1-B324-E53****4DD44
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

type CreateDatasetVersionRequest struct {
	// example:
	//
	// 300
	DataCount *int64 `json:"DataCount,omitempty" xml:"DataCount,omitempty"`
	// example:
	//
	// 19000
	DataSize *int64 `json:"DataSize,omitempty" xml:"DataSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// OSS
	DataSourceType *string  `json:"DataSourceType,omitempty" xml:"DataSourceType,omitempty"`
	Description    *string  `json:"Description,omitempty" xml:"Description,omitempty"`
	Labels         []*Label `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	// example:
	//
	// {
	//
	//   "mountPath": "/mnt/data/"
	//
	// }
	Options *string `json:"Options,omitempty" xml:"Options,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// DIRECTORY
	Property *string `json:"Property,omitempty" xml:"Property,omitempty"`
	// example:
	//
	// d-a0xbe5n03bhqof46ce
	SourceId *string `json:"SourceId,omitempty" xml:"SourceId,omitempty"`
	// example:
	//
	// USER
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// oss://mybucket.oss-cn-beijing.aliyuncs.com/mypath/
	Uri *string `json:"Uri,omitempty" xml:"Uri,omitempty"`
}

func (s CreateDatasetVersionRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDatasetVersionRequest) GoString() string {
	return s.String()
}

func (s *CreateDatasetVersionRequest) SetDataCount(v int64) *CreateDatasetVersionRequest {
	s.DataCount = &v
	return s
}

func (s *CreateDatasetVersionRequest) SetDataSize(v int64) *CreateDatasetVersionRequest {
	s.DataSize = &v
	return s
}

func (s *CreateDatasetVersionRequest) SetDataSourceType(v string) *CreateDatasetVersionRequest {
	s.DataSourceType = &v
	return s
}

func (s *CreateDatasetVersionRequest) SetDescription(v string) *CreateDatasetVersionRequest {
	s.Description = &v
	return s
}

func (s *CreateDatasetVersionRequest) SetLabels(v []*Label) *CreateDatasetVersionRequest {
	s.Labels = v
	return s
}

func (s *CreateDatasetVersionRequest) SetOptions(v string) *CreateDatasetVersionRequest {
	s.Options = &v
	return s
}

func (s *CreateDatasetVersionRequest) SetProperty(v string) *CreateDatasetVersionRequest {
	s.Property = &v
	return s
}

func (s *CreateDatasetVersionRequest) SetSourceId(v string) *CreateDatasetVersionRequest {
	s.SourceId = &v
	return s
}

func (s *CreateDatasetVersionRequest) SetSourceType(v string) *CreateDatasetVersionRequest {
	s.SourceType = &v
	return s
}

func (s *CreateDatasetVersionRequest) SetUri(v string) *CreateDatasetVersionRequest {
	s.Uri = &v
	return s
}

type CreateDatasetVersionResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// v1
	VersionName *string `json:"VersionName,omitempty" xml:"VersionName,omitempty"`
}

func (s CreateDatasetVersionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDatasetVersionResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDatasetVersionResponseBody) SetRequestId(v string) *CreateDatasetVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDatasetVersionResponseBody) SetVersionName(v string) *CreateDatasetVersionResponseBody {
	s.VersionName = &v
	return s
}

type CreateDatasetVersionResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDatasetVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDatasetVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDatasetVersionResponse) GoString() string {
	return s.String()
}

func (s *CreateDatasetVersionResponse) SetHeaders(v map[string]*string) *CreateDatasetVersionResponse {
	s.Headers = v
	return s
}

func (s *CreateDatasetVersionResponse) SetStatusCode(v int32) *CreateDatasetVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDatasetVersionResponse) SetBody(v *CreateDatasetVersionResponseBody) *CreateDatasetVersionResponse {
	s.Body = v
	return s
}

type CreateDatasetVersionLabelsRequest struct {
	// This parameter is required.
	Labels []*Label `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
}

func (s CreateDatasetVersionLabelsRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDatasetVersionLabelsRequest) GoString() string {
	return s.String()
}

func (s *CreateDatasetVersionLabelsRequest) SetLabels(v []*Label) *CreateDatasetVersionLabelsRequest {
	s.Labels = v
	return s
}

type CreateDatasetVersionLabelsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDatasetVersionLabelsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDatasetVersionLabelsResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDatasetVersionLabelsResponseBody) SetRequestId(v string) *CreateDatasetVersionLabelsResponseBody {
	s.RequestId = &v
	return s
}

type CreateDatasetVersionLabelsResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDatasetVersionLabelsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDatasetVersionLabelsResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDatasetVersionLabelsResponse) GoString() string {
	return s.String()
}

func (s *CreateDatasetVersionLabelsResponse) SetHeaders(v map[string]*string) *CreateDatasetVersionLabelsResponse {
	s.Headers = v
	return s
}

func (s *CreateDatasetVersionLabelsResponse) SetStatusCode(v int32) *CreateDatasetVersionLabelsResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDatasetVersionLabelsResponse) SetBody(v *CreateDatasetVersionLabelsResponseBody) *CreateDatasetVersionLabelsResponse {
	s.Body = v
	return s
}

type CreateExperimentRequest struct {
	// example:
	//
	// PRIVATE
	Accessibility *string `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	// Artifact的OSS存储路径
	//
	// example:
	//
	// oss://test-bucket.oss-cn-hangzhou.aliyuncs.com/test
	ArtifactUri *string `json:"ArtifactUri,omitempty" xml:"ArtifactUri,omitempty"`
	// 标签
	Labels []*LabelInfo `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	// 名称
	//
	// This parameter is required.
	//
	// example:
	//
	// exp-test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 工作空间ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 478**
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s CreateExperimentRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateExperimentRequest) GoString() string {
	return s.String()
}

func (s *CreateExperimentRequest) SetAccessibility(v string) *CreateExperimentRequest {
	s.Accessibility = &v
	return s
}

func (s *CreateExperimentRequest) SetArtifactUri(v string) *CreateExperimentRequest {
	s.ArtifactUri = &v
	return s
}

func (s *CreateExperimentRequest) SetLabels(v []*LabelInfo) *CreateExperimentRequest {
	s.Labels = v
	return s
}

func (s *CreateExperimentRequest) SetName(v string) *CreateExperimentRequest {
	s.Name = &v
	return s
}

func (s *CreateExperimentRequest) SetWorkspaceId(v string) *CreateExperimentRequest {
	s.WorkspaceId = &v
	return s
}

type CreateExperimentResponseBody struct {
	ExperimentId *string `json:"ExperimentId,omitempty" xml:"ExperimentId,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateExperimentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateExperimentResponseBody) GoString() string {
	return s.String()
}

func (s *CreateExperimentResponseBody) SetExperimentId(v string) *CreateExperimentResponseBody {
	s.ExperimentId = &v
	return s
}

func (s *CreateExperimentResponseBody) SetRequestId(v string) *CreateExperimentResponseBody {
	s.RequestId = &v
	return s
}

type CreateExperimentResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateExperimentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateExperimentResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateExperimentResponse) GoString() string {
	return s.String()
}

func (s *CreateExperimentResponse) SetHeaders(v map[string]*string) *CreateExperimentResponse {
	s.Headers = v
	return s
}

func (s *CreateExperimentResponse) SetStatusCode(v int32) *CreateExperimentResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateExperimentResponse) SetBody(v *CreateExperimentResponseBody) *CreateExperimentResponse {
	s.Body = v
	return s
}

type CreateMemberRequest struct {
	// This parameter is required.
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
	// This parameter is required.
	Roles []*string `json:"Roles,omitempty" xml:"Roles,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 21513926******88039
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
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
	Members []*CreateMemberResponseBodyMembers `json:"Members,omitempty" xml:"Members,omitempty" type:"Repeated"`
	// example:
	//
	// DA869D1B-035A-43B2-ACC1-C56681BD9FAA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// example:
	//
	// myDisplayName
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// example:
	//
	// 145883-21513926******88039
	MemberId *string   `json:"MemberId,omitempty" xml:"MemberId,omitempty"`
	Roles    []*string `json:"Roles,omitempty" xml:"Roles,omitempty" type:"Repeated"`
	// example:
	//
	// 21513926******88039
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
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
	// example:
	//
	// PRIVATE
	Accessibility *string `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	// example:
	//
	// nlp
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// if can be null:
	// true
	ExtraInfo        map[string]interface{} `json:"ExtraInfo,omitempty" xml:"ExtraInfo,omitempty"`
	Labels           []*Label               `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	ModelDescription *string                `json:"ModelDescription,omitempty" xml:"ModelDescription,omitempty"`
	// example:
	//
	// https://*.md
	ModelDoc *string `json:"ModelDoc,omitempty" xml:"ModelDoc,omitempty"`
	// This parameter is required.
	ModelName *string `json:"ModelName,omitempty" xml:"ModelName,omitempty"`
	// example:
	//
	// Checkpoint
	ModelType *string `json:"ModelType,omitempty" xml:"ModelType,omitempty"`
	// example:
	//
	// 101
	OrderNumber *int64 `json:"OrderNumber,omitempty" xml:"OrderNumber,omitempty"`
	// example:
	//
	// ModelScope
	Origin *string `json:"Origin,omitempty" xml:"Origin,omitempty"`
	// example:
	//
	// text-classification
	Task *string `json:"Task,omitempty" xml:"Task,omitempty"`
	// example:
	//
	// 796**
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
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
	// example:
	//
	// model-rbvg5wzljz****ks92
	ModelId *string `json:"ModelId,omitempty" xml:"ModelId,omitempty"`
	// example:
	//
	// 9DAD3112-AE22-5563-9A02-5C7E8****E35
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
	// example:
	//
	// F81D9EC0-1872-50F5-A96C-A0647D****1D
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
	// example:
	//
	// Approved
	ApprovalStatus *string `json:"ApprovalStatus,omitempty" xml:"ApprovalStatus,omitempty"`
	// example:
	//
	// {}
	CompressionSpec map[string]interface{} `json:"CompressionSpec,omitempty" xml:"CompressionSpec,omitempty"`
	// example:
	//
	// {}
	EvaluationSpec map[string]interface{} `json:"EvaluationSpec,omitempty" xml:"EvaluationSpec,omitempty"`
	// example:
	//
	// {}
	ExtraInfo map[string]interface{} `json:"ExtraInfo,omitempty" xml:"ExtraInfo,omitempty"`
	// example:
	//
	// SavedModel
	FormatType *string `json:"FormatType,omitempty" xml:"FormatType,omitempty"`
	// example:
	//
	// TensorFlow
	FrameworkType *string `json:"FrameworkType,omitempty" xml:"FrameworkType,omitempty"`
	// example:
	//
	// {
	//
	//     "processor": "tensorflow_gpu_1.12"
	//
	// }
	InferenceSpec map[string]interface{} `json:"InferenceSpec,omitempty" xml:"InferenceSpec,omitempty"`
	Labels        []*Label               `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	// example:
	//
	// {}
	Metrics map[string]interface{} `json:"Metrics,omitempty" xml:"Metrics,omitempty"`
	// example:
	//
	// {}
	Options *string `json:"Options,omitempty" xml:"Options,omitempty"`
	// example:
	//
	// region=cn-shanghai,workspaceId=13**,kind=PipelineRun,id=run-sakdb****jdf
	SourceId *string `json:"SourceId,omitempty" xml:"SourceId,omitempty"`
	// example:
	//
	// PAIFlow
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// example:
	//
	// {}
	TrainingSpec map[string]interface{} `json:"TrainingSpec,omitempty" xml:"TrainingSpec,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// oss://mybucket.oss-cn-beijing.aliyuncs.com/mypath/
	Uri                *string `json:"Uri,omitempty" xml:"Uri,omitempty"`
	VersionDescription *string `json:"VersionDescription,omitempty" xml:"VersionDescription,omitempty"`
	// example:
	//
	// 0.1.0
	VersionName *string `json:"VersionName,omitempty" xml:"VersionName,omitempty"`
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

func (s *CreateModelVersionRequest) SetCompressionSpec(v map[string]interface{}) *CreateModelVersionRequest {
	s.CompressionSpec = v
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
	// example:
	//
	// 21645FCD-BAB9-5742-89AE-AEB27****B2E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 0.1.0
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
	// example:
	//
	// 5A14FA81-DD4E-******-6343FE44B941
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
	// example:
	//
	// true
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
	// example:
	//
	// true
	AutoRenew *bool `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// example:
	//
	// POSTPAY
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// example:
	//
	// 1
	Duration           *int64                                                  `json:"Duration,omitempty" xml:"Duration,omitempty"`
	InstanceProperties []*CreateProductOrdersRequestProductsInstanceProperties `json:"InstanceProperties,omitempty" xml:"InstanceProperties,omitempty" type:"Repeated"`
	// example:
	//
	// BUY
	OrderType *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	// example:
	//
	// Month
	PricingCycle *string `json:"PricingCycle,omitempty" xml:"PricingCycle,omitempty"`
	// example:
	//
	// DataWorks_share
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
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
	// example:
	//
	// commodity_type。
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// oss。
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
	// example:
	//
	// 3ed6a882-0d85-4dd8-ad36-cd8d74ab9fdb
	BuyProductRequestId *string `json:"BuyProductRequestId,omitempty" xml:"BuyProductRequestId,omitempty"`
	Message             *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 210292536260646
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// example:
	//
	// ksdjf-jksd-*****slkdjf
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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

type CreateRunRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// exp-6thbb5xrbmp*****
	ExperimentId *string  `json:"ExperimentId,omitempty" xml:"ExperimentId,omitempty"`
	Labels       []*Label `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	// example:
	//
	// myName
	Name   *string     `json:"Name,omitempty" xml:"Name,omitempty"`
	Params []*RunParam `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	// example:
	//
	// job-jdnhf***fnrimv
	SourceId *string `json:"SourceId,omitempty" xml:"SourceId,omitempty"`
	// example:
	//
	// DLC
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
}

func (s CreateRunRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateRunRequest) GoString() string {
	return s.String()
}

func (s *CreateRunRequest) SetExperimentId(v string) *CreateRunRequest {
	s.ExperimentId = &v
	return s
}

func (s *CreateRunRequest) SetLabels(v []*Label) *CreateRunRequest {
	s.Labels = v
	return s
}

func (s *CreateRunRequest) SetName(v string) *CreateRunRequest {
	s.Name = &v
	return s
}

func (s *CreateRunRequest) SetParams(v []*RunParam) *CreateRunRequest {
	s.Params = v
	return s
}

func (s *CreateRunRequest) SetSourceId(v string) *CreateRunRequest {
	s.SourceId = &v
	return s
}

func (s *CreateRunRequest) SetSourceType(v string) *CreateRunRequest {
	s.SourceType = &v
	return s
}

type CreateRunResponseBody struct {
	// example:
	//
	// run-1meoz7VJd2C6f****
	RunId *string `json:"RunId,omitempty" xml:"RunId,omitempty"`
	// Id of the request
	//
	// example:
	//
	// ADF6D849-*****-7E7030F0CE53
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateRunResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateRunResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRunResponseBody) SetRunId(v string) *CreateRunResponseBody {
	s.RunId = &v
	return s
}

func (s *CreateRunResponseBody) SetRequestId(v string) *CreateRunResponseBody {
	s.RequestId = &v
	return s
}

type CreateRunResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateRunResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateRunResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateRunResponse) GoString() string {
	return s.String()
}

func (s *CreateRunResponse) SetHeaders(v map[string]*string) *CreateRunResponse {
	s.Headers = v
	return s
}

func (s *CreateRunResponse) SetStatusCode(v int32) *CreateRunResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateRunResponse) SetBody(v *CreateRunResponseBody) *CreateRunResponse {
	s.Body = v
	return s
}

type CreateWorkspaceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// display name
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// This parameter is required.
	EnvTypes []*string `json:"EnvTypes,omitempty" xml:"EnvTypes,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// workspace_example
	WorkspaceName *string `json:"WorkspaceName,omitempty" xml:"WorkspaceName,omitempty"`
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
	// example:
	//
	// 1e195c5116124202371861018d5bde
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1234
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
	// example:
	//
	// CreateAndAttach
	Option *string `json:"Option,omitempty" xml:"Option,omitempty"`
	// This parameter is required.
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
	// This parameter is required.
	//
	// example:
	//
	// prod
	EnvType *string `json:"EnvType,omitempty" xml:"EnvType,omitempty"`
	// example:
	//
	// groupName
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// example:
	//
	// false
	IsDefault *bool                                            `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	Labels    []*CreateWorkspaceResourceRequestResourcesLabels `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// ResourceName
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// MaxCompute
	ProductType  *string                                          `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	Quotas       []*CreateWorkspaceResourceRequestResourcesQuotas `json:"Quotas,omitempty" xml:"Quotas,omitempty" type:"Repeated"`
	ResourceType *string                                          `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Spec         map[string]interface{}                           `json:"Spec,omitempty" xml:"Spec,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// 232892******92912
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
	// example:
	//
	// 1e195c5116124202371861018d5bde
	RequestId *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Resources []*CreateWorkspaceResourceResponseBodyResources `json:"Resources,omitempty" xml:"Resources,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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
	// example:
	//
	// 1234
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
	// example:
	//
	// code-20210111103721-85qz78ia96lu
	CodeSourceId *string `json:"CodeSourceId,omitempty" xml:"CodeSourceId,omitempty"`
	// example:
	//
	// 5A14FA81-DD4E-******-6343FE44B941
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// example:
	//
	// A0F049F0-8D69-5BAC-8F10-B******A34C
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
	// example:
	//
	// key1,key2
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
	// example:
	//
	// 64B50C1D-D4C2-560C-86A3-A6ED****16D
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

type DeleteDatasetVersionResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDatasetVersionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDatasetVersionResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDatasetVersionResponseBody) SetRequestId(v string) *DeleteDatasetVersionResponseBody {
	s.RequestId = &v
	return s
}

type DeleteDatasetVersionResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDatasetVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDatasetVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDatasetVersionResponse) GoString() string {
	return s.String()
}

func (s *DeleteDatasetVersionResponse) SetHeaders(v map[string]*string) *DeleteDatasetVersionResponse {
	s.Headers = v
	return s
}

func (s *DeleteDatasetVersionResponse) SetStatusCode(v int32) *DeleteDatasetVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDatasetVersionResponse) SetBody(v *DeleteDatasetVersionResponseBody) *DeleteDatasetVersionResponse {
	s.Body = v
	return s
}

type DeleteDatasetVersionLabelsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// key1,key2
	Keys *string `json:"Keys,omitempty" xml:"Keys,omitempty"`
}

func (s DeleteDatasetVersionLabelsRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDatasetVersionLabelsRequest) GoString() string {
	return s.String()
}

func (s *DeleteDatasetVersionLabelsRequest) SetKeys(v string) *DeleteDatasetVersionLabelsRequest {
	s.Keys = &v
	return s
}

type DeleteDatasetVersionLabelsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDatasetVersionLabelsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDatasetVersionLabelsResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDatasetVersionLabelsResponseBody) SetRequestId(v string) *DeleteDatasetVersionLabelsResponseBody {
	s.RequestId = &v
	return s
}

type DeleteDatasetVersionLabelsResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDatasetVersionLabelsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDatasetVersionLabelsResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDatasetVersionLabelsResponse) GoString() string {
	return s.String()
}

func (s *DeleteDatasetVersionLabelsResponse) SetHeaders(v map[string]*string) *DeleteDatasetVersionLabelsResponse {
	s.Headers = v
	return s
}

func (s *DeleteDatasetVersionLabelsResponse) SetStatusCode(v int32) *DeleteDatasetVersionLabelsResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDatasetVersionLabelsResponse) SetBody(v *DeleteDatasetVersionLabelsResponseBody) *DeleteDatasetVersionLabelsResponse {
	s.Body = v
	return s
}

type DeleteExperimentResponseBody struct {
	// example:
	//
	// 8D7B2E70-F770-505B-A672-09F1D8F2EC1E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteExperimentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteExperimentResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteExperimentResponseBody) SetRequestId(v string) *DeleteExperimentResponseBody {
	s.RequestId = &v
	return s
}

type DeleteExperimentResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteExperimentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteExperimentResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteExperimentResponse) GoString() string {
	return s.String()
}

func (s *DeleteExperimentResponse) SetHeaders(v map[string]*string) *DeleteExperimentResponse {
	s.Headers = v
	return s
}

func (s *DeleteExperimentResponse) SetStatusCode(v int32) *DeleteExperimentResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteExperimentResponse) SetBody(v *DeleteExperimentResponseBody) *DeleteExperimentResponse {
	s.Body = v
	return s
}

type DeleteExperimentLabelResponseBody struct {
	// example:
	//
	// 5A14FA81-DD4E-******-6343FE44B941
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteExperimentLabelResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteExperimentLabelResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteExperimentLabelResponseBody) SetRequestId(v string) *DeleteExperimentLabelResponseBody {
	s.RequestId = &v
	return s
}

type DeleteExperimentLabelResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteExperimentLabelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteExperimentLabelResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteExperimentLabelResponse) GoString() string {
	return s.String()
}

func (s *DeleteExperimentLabelResponse) SetHeaders(v map[string]*string) *DeleteExperimentLabelResponse {
	s.Headers = v
	return s
}

func (s *DeleteExperimentLabelResponse) SetStatusCode(v int32) *DeleteExperimentLabelResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteExperimentLabelResponse) SetBody(v *DeleteExperimentLabelResponseBody) *DeleteExperimentLabelResponse {
	s.Body = v
	return s
}

type DeleteMembersRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 145883-21513926******88039,145883-2769726******87513
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
	Code    *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// D5BFFEE3-6025-443F-8A03-02D619B5C4B9
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
	// example:
	//
	// 5A14FA81-DD4E-******-6343FE44B941
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
	// example:
	//
	// key1,key2
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
	// example:
	//
	// 5A14FA81-DD4E-******-6343FE44B941
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
	// example:
	//
	// 5A14FA81-DD4E-******-6343FE44B941
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
	// example:
	//
	// key1,key2
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
	// example:
	//
	// 5A14FA81-DD4E-******-6343FE44B941
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

type DeleteRunResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// ADF6D849-*****-7E7030F0CE53
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteRunResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteRunResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRunResponseBody) SetRequestId(v string) *DeleteRunResponseBody {
	s.RequestId = &v
	return s
}

type DeleteRunResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteRunResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteRunResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteRunResponse) GoString() string {
	return s.String()
}

func (s *DeleteRunResponse) SetHeaders(v map[string]*string) *DeleteRunResponse {
	s.Headers = v
	return s
}

func (s *DeleteRunResponse) SetStatusCode(v int32) *DeleteRunResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteRunResponse) SetBody(v *DeleteRunResponseBody) *DeleteRunResponse {
	s.Body = v
	return s
}

type DeleteRunLabelResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// ADF6D849-*****-7E7030F0CE53
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteRunLabelResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteRunLabelResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRunLabelResponseBody) SetRequestId(v string) *DeleteRunLabelResponseBody {
	s.RequestId = &v
	return s
}

type DeleteRunLabelResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteRunLabelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteRunLabelResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteRunLabelResponse) GoString() string {
	return s.String()
}

func (s *DeleteRunLabelResponse) SetHeaders(v map[string]*string) *DeleteRunLabelResponse {
	s.Headers = v
	return s
}

func (s *DeleteRunLabelResponse) SetStatusCode(v int32) *DeleteRunLabelResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteRunLabelResponse) SetBody(v *DeleteRunLabelResponseBody) *DeleteRunLabelResponse {
	s.Body = v
	return s
}

type DeleteWorkspaceResponseBody struct {
	// example:
	//
	// 5A14FA81-DD4E-******-6343FE44B941
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
	// example:
	//
	// group
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	Labels    *string `json:"Labels,omitempty" xml:"Labels,omitempty"`
	// example:
	//
	// DetachAndDelete
	Option *string `json:"Option,omitempty" xml:"Option,omitempty"`
	// example:
	//
	// DLC
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
	// example:
	//
	// 5A14FA81-DD4E-******-6343FE44B941
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
	// example:
	//
	// PRIVATE
	Accessibility *string `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	// example:
	//
	// master
	CodeBranch *string `json:"CodeBranch,omitempty" xml:"CodeBranch,omitempty"`
	// example:
	//
	// 44da10***********
	CodeCommit *string `json:"CodeCommit,omitempty" xml:"CodeCommit,omitempty"`
	// example:
	//
	// https://code.aliyun.com/pai-dlc/examples.git
	CodeRepo *string `json:"CodeRepo,omitempty" xml:"CodeRepo,omitempty"`
	// example:
	//
	// xxxx
	CodeRepoAccessToken *string `json:"CodeRepoAccessToken,omitempty" xml:"CodeRepoAccessToken,omitempty"`
	// example:
	//
	// user1
	CodeRepoUserName *string `json:"CodeRepoUserName,omitempty" xml:"CodeRepoUserName,omitempty"`
	// example:
	//
	// code-202**********
	CodeSourceId *string `json:"CodeSourceId,omitempty" xml:"CodeSourceId,omitempty"`
	// example:
	//
	// This is my data source 1.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// MyCodeSource1
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// example:
	//
	// 2021-01-12T23:36:01.123Z
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// example:
	//
	// 2021-01-12T23:36:01.123Z
	GmtModifyTime *string `json:"GmtModifyTime,omitempty" xml:"GmtModifyTime,omitempty"`
	// example:
	//
	// /root/code
	MountPath *string `json:"MountPath,omitempty" xml:"MountPath,omitempty"`
	// example:
	//
	// 5A14FA81-DD4E-******-6343FE44B941
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1722********
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// example:
	//
	// 1234
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
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
	// example:
	//
	// PRIVATE
	Accessibility *string `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	// example:
	//
	// NAS
	DataSourceType *string `json:"DataSourceType,omitempty" xml:"DataSourceType,omitempty"`
	// example:
	//
	// COMMON
	DataType *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	// example:
	//
	// d-rbvg5wz****c9ks92
	DatasetId   *string `json:"DatasetId,omitempty" xml:"DatasetId,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 2021-01-30T12:51:33.028Z
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// example:
	//
	// 2021-01-30T12:51:33.028Z
	GmtModifiedTime *string         `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	Labels          []*Label        `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	LatestVersion   *DatasetVersion `json:"LatestVersion,omitempty" xml:"LatestVersion,omitempty"`
	// example:
	//
	// myName
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// {
	//
	//   "mountPath": "/mnt/data/"
	//
	// }
	Options *string `json:"Options,omitempty" xml:"Options,omitempty"`
	// example:
	//
	// 1631044****3440
	OwnerId *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// DIRECTORY
	Property     *string `json:"Property,omitempty" xml:"Property,omitempty"`
	Provider     *string `json:"Provider,omitempty" xml:"Provider,omitempty"`
	ProviderType *string `json:"ProviderType,omitempty" xml:"ProviderType,omitempty"`
	// example:
	//
	// 5A14FA81-DD4E-******-6343FE44B941
	RequestId            *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SourceDatasetId      *string `json:"SourceDatasetId,omitempty" xml:"SourceDatasetId,omitempty"`
	SourceDatasetVersion *string `json:"SourceDatasetVersion,omitempty" xml:"SourceDatasetVersion,omitempty"`
	// example:
	//
	// jdnhf***fnrimv
	SourceId *string `json:"SourceId,omitempty" xml:"SourceId,omitempty"`
	// example:
	//
	// USER
	SourceType      *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	TagTemplateType *string `json:"TagTemplateType,omitempty" xml:"TagTemplateType,omitempty"`
	// example:
	//
	// nas://09f****f2.cn-hangzhou/
	Uri *string `json:"Uri,omitempty" xml:"Uri,omitempty"`
	// example:
	//
	// 2485765****023475
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// example:
	//
	// 478**
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
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

func (s *GetDatasetResponseBody) SetLatestVersion(v *DatasetVersion) *GetDatasetResponseBody {
	s.LatestVersion = v
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

func (s *GetDatasetResponseBody) SetProvider(v string) *GetDatasetResponseBody {
	s.Provider = &v
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

func (s *GetDatasetResponseBody) SetSourceDatasetId(v string) *GetDatasetResponseBody {
	s.SourceDatasetId = &v
	return s
}

func (s *GetDatasetResponseBody) SetSourceDatasetVersion(v string) *GetDatasetResponseBody {
	s.SourceDatasetVersion = &v
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

func (s *GetDatasetResponseBody) SetTagTemplateType(v string) *GetDatasetResponseBody {
	s.TagTemplateType = &v
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

type GetDatasetVersionResponseBody struct {
	// 数据集的数据量
	DataCount *int64 `json:"DataCount,omitempty" xml:"DataCount,omitempty"`
	// 数据集版本的数据大小。
	DataSize *int64 `json:"DataSize,omitempty" xml:"DataSize,omitempty"`
	// 数据源类型。支持以下取值：
	//
	// - OSS：阿里云对象存储（OSS）。
	//
	// - NAS：阿里云文件存储（NAS）。
	//
	// This parameter is required.
	DataSourceType *string `json:"DataSourceType,omitempty" xml:"DataSourceType,omitempty"`
	// 代表资源一级ID的资源属性字段
	DatasetId *string `json:"DatasetId,omitempty" xml:"DatasetId,omitempty"`
	// 数据集版本的描述信息。
	Description   *string `json:"Description,omitempty" xml:"Description,omitempty"`
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// 创建时间。
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// 代表资源标签的资源属性字段
	Labels []*Label `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	// 扩展字段，JsonString类型。
	//
	// 当DLC使用数据集时，可通过配置mountPath字段指定数据集默认挂载路径。
	Options *string `json:"Options,omitempty" xml:"Options,omitempty"`
	// 数据集的属性。支持以下取值：
	//
	// - FILE：文件。
	//
	// - DIRECTORY：文件夹。
	//
	// This parameter is required.
	Property  *string `json:"Property,omitempty" xml:"Property,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 数据来源ID。
	SourceId *string `json:"SourceId,omitempty" xml:"SourceId,omitempty"`
	// 数据来源类型，默认为USER。支持以下取值：
	//
	// - PAI-PUBLIC-DATASET：PAI公共数据集。
	//
	// - ITAG：iTAG模块标注结果生成的数据集。
	//
	// - USER：用户注册的数据集。
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// Uri配置样例如下：
	//
	// - 数据源类型为OSS：`oss://bucket.endpoint/object`
	//
	// - 数据源类型为NAS：
	//
	// 通用型NAS格式为：`nas://<nasfisid>.region/subpath/to/dir/`；
	//
	// CPFS1.0：`nas://<cpfs-fsid>.region/subpath/to/dir/`；
	//
	// CPFS2.0：`nas://<cpfs-fsid>.region/<protocolserviceid>/`。
	//
	// CPFS1.0和CPFS2.0根据fsid的格式来区分：CPFS1.0 格式为cpfs-<8位ascii字符>；CPFS2.0 格式为cpfs-<16为ascii字符>。
	//
	// This parameter is required.
	Uri *string `json:"Uri,omitempty" xml:"Uri,omitempty"`
	// 代表资源名称的资源属性字段
	VersionName *string `json:"VersionName,omitempty" xml:"VersionName,omitempty"`
}

func (s GetDatasetVersionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDatasetVersionResponseBody) GoString() string {
	return s.String()
}

func (s *GetDatasetVersionResponseBody) SetDataCount(v int64) *GetDatasetVersionResponseBody {
	s.DataCount = &v
	return s
}

func (s *GetDatasetVersionResponseBody) SetDataSize(v int64) *GetDatasetVersionResponseBody {
	s.DataSize = &v
	return s
}

func (s *GetDatasetVersionResponseBody) SetDataSourceType(v string) *GetDatasetVersionResponseBody {
	s.DataSourceType = &v
	return s
}

func (s *GetDatasetVersionResponseBody) SetDatasetId(v string) *GetDatasetVersionResponseBody {
	s.DatasetId = &v
	return s
}

func (s *GetDatasetVersionResponseBody) SetDescription(v string) *GetDatasetVersionResponseBody {
	s.Description = &v
	return s
}

func (s *GetDatasetVersionResponseBody) SetGmtCreateTime(v string) *GetDatasetVersionResponseBody {
	s.GmtCreateTime = &v
	return s
}

func (s *GetDatasetVersionResponseBody) SetGmtModifiedTime(v string) *GetDatasetVersionResponseBody {
	s.GmtModifiedTime = &v
	return s
}

func (s *GetDatasetVersionResponseBody) SetLabels(v []*Label) *GetDatasetVersionResponseBody {
	s.Labels = v
	return s
}

func (s *GetDatasetVersionResponseBody) SetOptions(v string) *GetDatasetVersionResponseBody {
	s.Options = &v
	return s
}

func (s *GetDatasetVersionResponseBody) SetProperty(v string) *GetDatasetVersionResponseBody {
	s.Property = &v
	return s
}

func (s *GetDatasetVersionResponseBody) SetRequestId(v string) *GetDatasetVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDatasetVersionResponseBody) SetSourceId(v string) *GetDatasetVersionResponseBody {
	s.SourceId = &v
	return s
}

func (s *GetDatasetVersionResponseBody) SetSourceType(v string) *GetDatasetVersionResponseBody {
	s.SourceType = &v
	return s
}

func (s *GetDatasetVersionResponseBody) SetUri(v string) *GetDatasetVersionResponseBody {
	s.Uri = &v
	return s
}

func (s *GetDatasetVersionResponseBody) SetVersionName(v string) *GetDatasetVersionResponseBody {
	s.VersionName = &v
	return s
}

type GetDatasetVersionResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDatasetVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDatasetVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDatasetVersionResponse) GoString() string {
	return s.String()
}

func (s *GetDatasetVersionResponse) SetHeaders(v map[string]*string) *GetDatasetVersionResponse {
	s.Headers = v
	return s
}

func (s *GetDatasetVersionResponse) SetStatusCode(v int32) *GetDatasetVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDatasetVersionResponse) SetBody(v *GetDatasetVersionResponseBody) *GetDatasetVersionResponse {
	s.Body = v
	return s
}

type GetDefaultWorkspaceRequest struct {
	// example:
	//
	// false
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
	Conditions []*GetDefaultWorkspaceResponseBodyConditions `json:"Conditions,omitempty" xml:"Conditions,omitempty" type:"Repeated"`
	// example:
	//
	// 17915******4216
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// example:
	//
	// workspace description example
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// workspace-example
	DisplayName *string   `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	EnvTypes    []*string `json:"EnvTypes,omitempty" xml:"EnvTypes,omitempty" type:"Repeated"`
	// example:
	//
	// 2021-01-21T17:12:35.232Z
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// example:
	//
	// 2021-01-21T17:12:35.232Z
	GmtModifiedTime *string                               `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	Owner           *GetDefaultWorkspaceResponseBodyOwner `json:"Owner,omitempty" xml:"Owner,omitempty" type:"Struct"`
	// example:
	//
	// 5A14FA81-DD4E-******-6343FE44B941
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// ENABLED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 1234
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
	// example:
	//
	// workspace-example
	WorkspaceName *string `json:"WorkspaceName,omitempty" xml:"WorkspaceName,omitempty"`
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
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// Create Failed
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// CREATING
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
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
	// example:
	//
	// 17915******4216
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// example:
	//
	// 17915******4216
	UserKp *string `json:"UserKp,omitempty" xml:"UserKp,omitempty"`
	// example:
	//
	// username
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

type GetExperimentRequest struct {
	// example:
	//
	// false
	Verbose *bool `json:"Verbose,omitempty" xml:"Verbose,omitempty"`
}

func (s GetExperimentRequest) String() string {
	return tea.Prettify(s)
}

func (s GetExperimentRequest) GoString() string {
	return s.String()
}

func (s *GetExperimentRequest) SetVerbose(v bool) *GetExperimentRequest {
	s.Verbose = &v
	return s
}

type GetExperimentResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *Experiment        `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetExperimentResponse) String() string {
	return tea.Prettify(s)
}

func (s GetExperimentResponse) GoString() string {
	return s.String()
}

func (s *GetExperimentResponse) SetHeaders(v map[string]*string) *GetExperimentResponse {
	s.Headers = v
	return s
}

func (s *GetExperimentResponse) SetStatusCode(v int32) *GetExperimentResponse {
	s.StatusCode = &v
	return s
}

func (s *GetExperimentResponse) SetBody(v *Experiment) *GetExperimentResponse {
	s.Body = v
	return s
}

type GetImageRequest struct {
	// example:
	//
	// false
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
	// example:
	//
	// PUBLIC
	Accessibility *string `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	Description   *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 2021-01-21T17:12:35.232Z
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// example:
	//
	// 2021-01-21T17:12:35.232Z
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// example:
	//
	// registry.cn-hangzhou.aliyuncs.******ession/nlp:gpu
	ImageUri *string                       `json:"ImageUri,omitempty" xml:"ImageUri,omitempty"`
	Labels   []*GetImageResponseBodyLabels `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	// example:
	//
	// nlp-compression
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 15577******8921
	ParentUserId *string `json:"ParentUserId,omitempty" xml:"ParentUserId,omitempty"`
	// example:
	//
	// 5A14FA81-DD4E-******-6343FE44B941
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Size      *int32  `json:"Size,omitempty" xml:"Size,omitempty"`
	// example:
	//
	// 15577******8921
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// example:
	//
	// 15945
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
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
	// example:
	//
	// system.chipType
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// GPU
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
	// example:
	//
	// 21513926******88039
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
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
	// example:
	//
	// myDisplayName
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// example:
	//
	// 2021-01-21T17:12:35.232Z
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// example:
	//
	// 145883-21513926******88039
	MemberId *string `json:"MemberId,omitempty" xml:"MemberId,omitempty"`
	// example:
	//
	// user1
	MemberName *string `json:"MemberName,omitempty" xml:"MemberName,omitempty"`
	// example:
	//
	// 5A14FA81-DD4E-******-6343FE44B941
	RequestId *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Roles     []*string `json:"Roles,omitempty" xml:"Roles,omitempty" type:"Repeated"`
	// example:
	//
	// 21513926******88039
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
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
	// example:
	//
	// PUBLIC
	Accessibility *string `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	// example:
	//
	// cv
	Domain    *string                `json:"Domain,omitempty" xml:"Domain,omitempty"`
	ExtraInfo map[string]interface{} `json:"ExtraInfo,omitempty" xml:"ExtraInfo,omitempty"`
	// example:
	//
	// 2021-01-30T12:51:33.028Z
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// example:
	//
	// 2021-01-30T12:51:33.028Z
	GmtModifiedTime  *string       `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	Labels           []*Label      `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	LatestVersion    *ModelVersion `json:"LatestVersion,omitempty" xml:"LatestVersion,omitempty"`
	ModelDescription *string       `json:"ModelDescription,omitempty" xml:"ModelDescription,omitempty"`
	// example:
	//
	// https://***.md
	ModelDoc *string `json:"ModelDoc,omitempty" xml:"ModelDoc,omitempty"`
	// example:
	//
	// model-rbvg5wzljz****ks92
	ModelId   *string `json:"ModelId,omitempty" xml:"ModelId,omitempty"`
	ModelName *string `json:"ModelName,omitempty" xml:"ModelName,omitempty"`
	// example:
	//
	// Checkpoint
	ModelType *string `json:"ModelType,omitempty" xml:"ModelType,omitempty"`
	// example:
	//
	// 1
	OrderNumber *int64 `json:"OrderNumber,omitempty" xml:"OrderNumber,omitempty"`
	// example:
	//
	// ModelScope
	Origin *string `json:"Origin,omitempty" xml:"Origin,omitempty"`
	// example:
	//
	// 1234567890******
	OwnerId *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// pai
	Provider *string `json:"Provider,omitempty" xml:"Provider,omitempty"`
	// example:
	//
	// 5A14FA81-DD4E-******-6343FE44B941
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// text-classification
	Task *string `json:"Task,omitempty" xml:"Task,omitempty"`
	// example:
	//
	// 1234567890******
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// example:
	//
	// 234**
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
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
	// example:
	//
	// Approved
	ApprovalStatus *string `json:"ApprovalStatus,omitempty" xml:"ApprovalStatus,omitempty"`
	// example:
	//
	// {}
	CompressionSpec map[string]interface{} `json:"CompressionSpec,omitempty" xml:"CompressionSpec,omitempty"`
	// example:
	//
	// {}
	EvaluationSpec map[string]interface{} `json:"EvaluationSpec,omitempty" xml:"EvaluationSpec,omitempty"`
	// example:
	//
	// {}
	ExtraInfo map[string]interface{} `json:"ExtraInfo,omitempty" xml:"ExtraInfo,omitempty"`
	// example:
	//
	// SavedModel
	FormatType *string `json:"FormatType,omitempty" xml:"FormatType,omitempty"`
	// example:
	//
	// TensorFlow
	FrameworkType *string `json:"FrameworkType,omitempty" xml:"FrameworkType,omitempty"`
	// example:
	//
	// 2021-01-30T12:51:33.028Z
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// example:
	//
	// 2021-01-30T12:51:33.028Z
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// example:
	//
	// {
	//
	//     "Processor": "tensorflow_gpu_1.12"
	//
	// }
	InferenceSpec map[string]interface{} `json:"InferenceSpec,omitempty" xml:"InferenceSpec,omitempty"`
	Labels        []*Label               `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	// example:
	//
	// {}
	Metrics map[string]interface{} `json:"Metrics,omitempty" xml:"Metrics,omitempty"`
	// example:
	//
	// {}
	Options *string `json:"Options,omitempty" xml:"Options,omitempty"`
	// example:
	//
	// 1234567890******
	OwnerId *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// 5A14FA81-DD4E-******-6343FE44B941
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// region=cn-shanghai,workspaceId=13**,kind=PipelineRun,id=run-sakdb****jdf
	SourceId *string `json:"SourceId,omitempty" xml:"SourceId,omitempty"`
	// example:
	//
	// PAIFlow
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// example:
	//
	// {}
	TrainingSpec map[string]interface{} `json:"TrainingSpec,omitempty" xml:"TrainingSpec,omitempty"`
	Uri          *string                `json:"Uri,omitempty" xml:"Uri,omitempty"`
	// example:
	//
	// 1234567890******
	UserId             *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	VersionDescription *string `json:"VersionDescription,omitempty" xml:"VersionDescription,omitempty"`
	// example:
	//
	// 0.1.0
	VersionName *string `json:"VersionName,omitempty" xml:"VersionName,omitempty"`
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

func (s *GetModelVersionResponseBody) SetCompressionSpec(v map[string]interface{}) *GetModelVersionResponseBody {
	s.CompressionSpec = v
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
	// example:
	//
	// PUBLIC
	Accessibility *string `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	// example:
	//
	// 17915******4216
	Creator  *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	Option   *string `json:"Option,omitempty" xml:"Option,omitempty"`
	Resource *string `json:"Resource,omitempty" xml:"Resource,omitempty"`
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
	// example:
	//
	// PaiDLC:ListJobs
	PermissionCode  *string                                     `json:"PermissionCode,omitempty" xml:"PermissionCode,omitempty"`
	PermissionRules []*GetPermissionResponseBodyPermissionRules `json:"PermissionRules,omitempty" xml:"PermissionRules,omitempty" type:"Repeated"`
	// example:
	//
	// 5A14FA81-DD4E-******-6343FE44B941
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// example:
	//
	// PRIVATE
	Accessibility *string `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	// example:
	//
	// CREATOR
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

type GetRunRequest struct {
	// example:
	//
	// true
	Verbose *bool `json:"Verbose,omitempty" xml:"Verbose,omitempty"`
}

func (s GetRunRequest) String() string {
	return tea.Prettify(s)
}

func (s GetRunRequest) GoString() string {
	return s.String()
}

func (s *GetRunRequest) SetVerbose(v bool) *GetRunRequest {
	s.Verbose = &v
	return s
}

type GetRunResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *Run               `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRunResponse) String() string {
	return tea.Prettify(s)
}

func (s GetRunResponse) GoString() string {
	return s.String()
}

func (s *GetRunResponse) SetHeaders(v map[string]*string) *GetRunResponse {
	s.Headers = v
	return s
}

func (s *GetRunResponse) SetStatusCode(v int32) *GetRunResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRunResponse) SetBody(v *Run) *GetRunResponse {
	s.Body = v
	return s
}

type GetWorkspaceRequest struct {
	// example:
	//
	// true
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
	AdminNames []*string `json:"AdminNames,omitempty" xml:"AdminNames,omitempty" type:"Repeated"`
	// example:
	//
	// 1157******94123
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// example:
	//
	// workspace description example
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// workspace-example
	DisplayName *string   `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	EnvTypes    []*string `json:"EnvTypes,omitempty" xml:"EnvTypes,omitempty" type:"Repeated"`
	// example:
	//
	// {"TenantId": "4286******98"}
	ExtraInfos map[string]interface{} `json:"ExtraInfos,omitempty" xml:"ExtraInfos,omitempty"`
	// example:
	//
	// 2021-01-21T17:12:35.232Z
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// example:
	//
	// 2021-01-21T17:12:35.232Z
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// example:
	//
	// true
	IsDefault *bool                          `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	Owner     *GetWorkspaceResponseBodyOwner `json:"Owner,omitempty" xml:"Owner,omitempty" type:"Struct"`
	// example:
	//
	// A0F049F0-8D69-5BAC-8F10-B4DED1B5A34C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// ENABLED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 1234
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
	// example:
	//
	// workspace-example
	WorkspaceName *string `json:"WorkspaceName,omitempty" xml:"WorkspaceName,omitempty"`
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
	// example:
	//
	// mings****t
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// example:
	//
	// 1157******94123
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// example:
	//
	// 1157******94123
	UserKp *string `json:"UserKp,omitempty" xml:"UserKp,omitempty"`
	// example:
	//
	// mings****t
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
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
	// example:
	//
	// MyDataSource
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// example:
	//
	// desc
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// GmtModifyTime
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// example:
	//
	// 1234
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
	// example:
	//
	// 5A14FA81-DD4E-******-6343FE44B941
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 2
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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

type ListDatasetVersionsRequest struct {
	// example:
	//
	// OSS
	DataSourcesTypes *string `json:"DataSourcesTypes,omitempty" xml:"DataSourcesTypes,omitempty"`
	// example:
	//
	// key1,key2
	LabelKeys *string `json:"LabelKeys,omitempty" xml:"LabelKeys,omitempty"`
	// example:
	//
	// value1,value2
	LableValues *string `json:"LableValues,omitempty" xml:"LableValues,omitempty"`
	// example:
	//
	// ASC
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// DIRECTORY
	Properties *string `json:"Properties,omitempty" xml:"Properties,omitempty"`
	// example:
	//
	// GmtCreateTime
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// example:
	//
	// d-a0xbe5n03bhqof46ce
	SourceId *string `json:"SourceId,omitempty" xml:"SourceId,omitempty"`
	// example:
	//
	// USER
	SourceTypes *string `json:"SourceTypes,omitempty" xml:"SourceTypes,omitempty"`
}

func (s ListDatasetVersionsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDatasetVersionsRequest) GoString() string {
	return s.String()
}

func (s *ListDatasetVersionsRequest) SetDataSourcesTypes(v string) *ListDatasetVersionsRequest {
	s.DataSourcesTypes = &v
	return s
}

func (s *ListDatasetVersionsRequest) SetLabelKeys(v string) *ListDatasetVersionsRequest {
	s.LabelKeys = &v
	return s
}

func (s *ListDatasetVersionsRequest) SetLableValues(v string) *ListDatasetVersionsRequest {
	s.LableValues = &v
	return s
}

func (s *ListDatasetVersionsRequest) SetOrder(v string) *ListDatasetVersionsRequest {
	s.Order = &v
	return s
}

func (s *ListDatasetVersionsRequest) SetPageNumber(v int32) *ListDatasetVersionsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListDatasetVersionsRequest) SetPageSize(v int32) *ListDatasetVersionsRequest {
	s.PageSize = &v
	return s
}

func (s *ListDatasetVersionsRequest) SetProperties(v string) *ListDatasetVersionsRequest {
	s.Properties = &v
	return s
}

func (s *ListDatasetVersionsRequest) SetSortBy(v string) *ListDatasetVersionsRequest {
	s.SortBy = &v
	return s
}

func (s *ListDatasetVersionsRequest) SetSourceId(v string) *ListDatasetVersionsRequest {
	s.SourceId = &v
	return s
}

func (s *ListDatasetVersionsRequest) SetSourceTypes(v string) *ListDatasetVersionsRequest {
	s.SourceTypes = &v
	return s
}

type ListDatasetVersionsResponseBody struct {
	DatasetVersions []*DatasetVersion `json:"DatasetVersions,omitempty" xml:"DatasetVersions,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize  *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 5
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListDatasetVersionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDatasetVersionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDatasetVersionsResponseBody) SetDatasetVersions(v []*DatasetVersion) *ListDatasetVersionsResponseBody {
	s.DatasetVersions = v
	return s
}

func (s *ListDatasetVersionsResponseBody) SetPageNumber(v int32) *ListDatasetVersionsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListDatasetVersionsResponseBody) SetPageSize(v int32) *ListDatasetVersionsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListDatasetVersionsResponseBody) SetRequestId(v string) *ListDatasetVersionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDatasetVersionsResponseBody) SetTotalCount(v int32) *ListDatasetVersionsResponseBody {
	s.TotalCount = &v
	return s
}

type ListDatasetVersionsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDatasetVersionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDatasetVersionsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDatasetVersionsResponse) GoString() string {
	return s.String()
}

func (s *ListDatasetVersionsResponse) SetHeaders(v map[string]*string) *ListDatasetVersionsResponse {
	s.Headers = v
	return s
}

func (s *ListDatasetVersionsResponse) SetStatusCode(v int32) *ListDatasetVersionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDatasetVersionsResponse) SetBody(v *ListDatasetVersionsResponseBody) *ListDatasetVersionsResponse {
	s.Body = v
	return s
}

type ListDatasetsRequest struct {
	// example:
	//
	// OSS
	DataSourceTypes *string `json:"DataSourceTypes,omitempty" xml:"DataSourceTypes,omitempty"`
	// example:
	//
	// COMMON,TEXT
	DataTypes *string `json:"DataTypes,omitempty" xml:"DataTypes,omitempty"`
	// example:
	//
	// test
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// example:
	//
	// myName
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// ASC
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// FILE
	Properties      *string `json:"Properties,omitempty" xml:"Properties,omitempty"`
	Provider        *string `json:"Provider,omitempty" xml:"Provider,omitempty"`
	SourceDatasetId *string `json:"SourceDatasetId,omitempty" xml:"SourceDatasetId,omitempty"`
	// example:
	//
	// d-rbvg5wzljzjhc9ks92
	SourceId *string `json:"SourceId,omitempty" xml:"SourceId,omitempty"`
	// example:
	//
	// USER,ITAG
	SourceTypes *string `json:"SourceTypes,omitempty" xml:"SourceTypes,omitempty"`
	// example:
	//
	// 324**
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
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

func (s *ListDatasetsRequest) SetSourceDatasetId(v string) *ListDatasetsRequest {
	s.SourceDatasetId = &v
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
	Datasets []*Dataset `json:"Datasets,omitempty" xml:"Datasets,omitempty" type:"Repeated"`
	// example:
	//
	// 5A14FA81-DD4E-******-6343FE44B941
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 15
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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

type ListExperimentRequest struct {
	// example:
	//
	// is_evaluation:true
	Labels *string `json:"Labels,omitempty" xml:"Labels,omitempty"`
	// example:
	//
	// 10
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// exp-test
	Name    *string                       `json:"Name,omitempty" xml:"Name,omitempty"`
	Options *ListExperimentRequestOptions `json:"Options,omitempty" xml:"Options,omitempty" type:"Struct"`
	// example:
	//
	// DESC
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// example:
	//
	// GmtCreateTime DESC,Name ASC
	OrderBy *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 0
	PageToken *int64 `json:"PageToken,omitempty" xml:"PageToken,omitempty"`
	// example:
	//
	// GmtCreateTime
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// example:
	//
	// false
	Verbose *bool `json:"Verbose,omitempty" xml:"Verbose,omitempty"`
	// example:
	//
	// 151739
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListExperimentRequest) String() string {
	return tea.Prettify(s)
}

func (s ListExperimentRequest) GoString() string {
	return s.String()
}

func (s *ListExperimentRequest) SetLabels(v string) *ListExperimentRequest {
	s.Labels = &v
	return s
}

func (s *ListExperimentRequest) SetMaxResults(v int64) *ListExperimentRequest {
	s.MaxResults = &v
	return s
}

func (s *ListExperimentRequest) SetName(v string) *ListExperimentRequest {
	s.Name = &v
	return s
}

func (s *ListExperimentRequest) SetOptions(v *ListExperimentRequestOptions) *ListExperimentRequest {
	s.Options = v
	return s
}

func (s *ListExperimentRequest) SetOrder(v string) *ListExperimentRequest {
	s.Order = &v
	return s
}

func (s *ListExperimentRequest) SetOrderBy(v string) *ListExperimentRequest {
	s.OrderBy = &v
	return s
}

func (s *ListExperimentRequest) SetPageNumber(v int32) *ListExperimentRequest {
	s.PageNumber = &v
	return s
}

func (s *ListExperimentRequest) SetPageSize(v int32) *ListExperimentRequest {
	s.PageSize = &v
	return s
}

func (s *ListExperimentRequest) SetPageToken(v int64) *ListExperimentRequest {
	s.PageToken = &v
	return s
}

func (s *ListExperimentRequest) SetSortBy(v string) *ListExperimentRequest {
	s.SortBy = &v
	return s
}

func (s *ListExperimentRequest) SetVerbose(v bool) *ListExperimentRequest {
	s.Verbose = &v
	return s
}

func (s *ListExperimentRequest) SetWorkspaceId(v string) *ListExperimentRequest {
	s.WorkspaceId = &v
	return s
}

type ListExperimentRequestOptions struct {
	// example:
	//
	// true
	MatchNameExactly *string `json:"match_name_exactly,omitempty" xml:"match_name_exactly,omitempty"`
}

func (s ListExperimentRequestOptions) String() string {
	return tea.Prettify(s)
}

func (s ListExperimentRequestOptions) GoString() string {
	return s.String()
}

func (s *ListExperimentRequestOptions) SetMatchNameExactly(v string) *ListExperimentRequestOptions {
	s.MatchNameExactly = &v
	return s
}

type ListExperimentShrinkRequest struct {
	// example:
	//
	// is_evaluation:true
	Labels *string `json:"Labels,omitempty" xml:"Labels,omitempty"`
	// example:
	//
	// 10
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// exp-test
	Name          *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OptionsShrink *string `json:"Options,omitempty" xml:"Options,omitempty"`
	// example:
	//
	// DESC
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// example:
	//
	// GmtCreateTime DESC,Name ASC
	OrderBy *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 0
	PageToken *int64 `json:"PageToken,omitempty" xml:"PageToken,omitempty"`
	// example:
	//
	// GmtCreateTime
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// example:
	//
	// false
	Verbose *bool `json:"Verbose,omitempty" xml:"Verbose,omitempty"`
	// example:
	//
	// 151739
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListExperimentShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListExperimentShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListExperimentShrinkRequest) SetLabels(v string) *ListExperimentShrinkRequest {
	s.Labels = &v
	return s
}

func (s *ListExperimentShrinkRequest) SetMaxResults(v int64) *ListExperimentShrinkRequest {
	s.MaxResults = &v
	return s
}

func (s *ListExperimentShrinkRequest) SetName(v string) *ListExperimentShrinkRequest {
	s.Name = &v
	return s
}

func (s *ListExperimentShrinkRequest) SetOptionsShrink(v string) *ListExperimentShrinkRequest {
	s.OptionsShrink = &v
	return s
}

func (s *ListExperimentShrinkRequest) SetOrder(v string) *ListExperimentShrinkRequest {
	s.Order = &v
	return s
}

func (s *ListExperimentShrinkRequest) SetOrderBy(v string) *ListExperimentShrinkRequest {
	s.OrderBy = &v
	return s
}

func (s *ListExperimentShrinkRequest) SetPageNumber(v int32) *ListExperimentShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListExperimentShrinkRequest) SetPageSize(v int32) *ListExperimentShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListExperimentShrinkRequest) SetPageToken(v int64) *ListExperimentShrinkRequest {
	s.PageToken = &v
	return s
}

func (s *ListExperimentShrinkRequest) SetSortBy(v string) *ListExperimentShrinkRequest {
	s.SortBy = &v
	return s
}

func (s *ListExperimentShrinkRequest) SetVerbose(v bool) *ListExperimentShrinkRequest {
	s.Verbose = &v
	return s
}

func (s *ListExperimentShrinkRequest) SetWorkspaceId(v string) *ListExperimentShrinkRequest {
	s.WorkspaceId = &v
	return s
}

type ListExperimentResponseBody struct {
	Experiments []*Experiment `json:"Experiments,omitempty" xml:"Experiments,omitempty" type:"Repeated"`
	// example:
	//
	// 0
	NextPageToken *int64 `json:"NextPageToken,omitempty" xml:"NextPageToken,omitempty"`
	// example:
	//
	// 5
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// example:
	//
	// 0C6835C5-A424-5AFB-ACC2-F1E3CA1ABF7C
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListExperimentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListExperimentResponseBody) GoString() string {
	return s.String()
}

func (s *ListExperimentResponseBody) SetExperiments(v []*Experiment) *ListExperimentResponseBody {
	s.Experiments = v
	return s
}

func (s *ListExperimentResponseBody) SetNextPageToken(v int64) *ListExperimentResponseBody {
	s.NextPageToken = &v
	return s
}

func (s *ListExperimentResponseBody) SetTotalCount(v int64) *ListExperimentResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListExperimentResponseBody) SetRequestId(v string) *ListExperimentResponseBody {
	s.RequestId = &v
	return s
}

type ListExperimentResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListExperimentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListExperimentResponse) String() string {
	return tea.Prettify(s)
}

func (s ListExperimentResponse) GoString() string {
	return s.String()
}

func (s *ListExperimentResponse) SetHeaders(v map[string]*string) *ListExperimentResponse {
	s.Headers = v
	return s
}

func (s *ListExperimentResponse) SetStatusCode(v int32) *ListExperimentResponse {
	s.StatusCode = &v
	return s
}

func (s *ListExperimentResponse) SetBody(v *ListExperimentResponseBody) *ListExperimentResponse {
	s.Body = v
	return s
}

type ListImageLabelsRequest struct {
	// example:
	//
	// image-4c62******53uor
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// example:
	//
	// system.framework=XGBoost 1.6.0,system.official=true
	LabelFilter *string `json:"LabelFilter,omitempty" xml:"LabelFilter,omitempty"`
	// example:
	//
	// system.framework,system.official
	LabelKeys *string `json:"LabelKeys,omitempty" xml:"LabelKeys,omitempty"`
	// example:
	//
	// cn-shanghai
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// example:
	//
	// 12345
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
	Labels []*ListImageLabelsResponseBodyLabels `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	// example:
	//
	// 5A14FA81-DD4E-******-6343FE44B941
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 2
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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
	// example:
	//
	// system.chipType
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// GPU
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
	ImageUri      *string `json:"ImageUri,omitempty" xml:"ImageUri,omitempty"`
	// example:
	//
	// system.framework=XGBoost 1.6.0,system.official=true
	Labels *string `json:"Labels,omitempty" xml:"Labels,omitempty"`
	// example:
	//
	// tensorflow_2.9
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// DESC
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 155**********904
	ParentUserId *string `json:"ParentUserId,omitempty" xml:"ParentUserId,omitempty"`
	Query        *string `json:"Query,omitempty" xml:"Query,omitempty"`
	// example:
	//
	// GmtCreateTime
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// example:
	//
	// 155**********904
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// example:
	//
	// true
	Verbose *bool `json:"Verbose,omitempty" xml:"Verbose,omitempty"`
	// example:
	//
	// 20******55
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
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

func (s *ListImagesRequest) SetImageUri(v string) *ListImagesRequest {
	s.ImageUri = &v
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
	Images []*ListImagesResponseBodyImages `json:"Images,omitempty" xml:"Images,omitempty" type:"Repeated"`
	// example:
	//
	// 5A14FA81-DD4E-******-6343FE44B941
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 2
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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
	// example:
	//
	// PUBLIC
	Accessibility *string `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	// example:
	//
	// desc
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 2021-01-21T17:12:35.232Z
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// example:
	//
	// 2021-01-21T17:12:35.232Z
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// example:
	//
	// image-tzi7f9******s45t
	ImageId  *string                               `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	ImageUri *string                               `json:"ImageUri,omitempty" xml:"ImageUri,omitempty"`
	Labels   []*ListImagesResponseBodyImagesLabels `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	// example:
	//
	// tensorflow_2.9
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 15577******82932
	ParentUserId *string `json:"ParentUserId,omitempty" xml:"ParentUserId,omitempty"`
	Size         *int32  `json:"Size,omitempty" xml:"Size,omitempty"`
	// example:
	//
	// 15577******82932
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// example:
	//
	// 20******55
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
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
	// example:
	//
	// system.chipType
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// GPU
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
	// example:
	//
	// zhangsan
	MemberName *string `json:"MemberName,omitempty" xml:"MemberName,omitempty"`
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// PAI.AlgoDeveloper
	Roles *string `json:"Roles,omitempty" xml:"Roles,omitempty"`
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
	Members []*ListMembersResponseBodyMembers `json:"Members,omitempty" xml:"Members,omitempty" type:"Repeated"`
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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
	// example:
	//
	// myDisplayName
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// example:
	//
	// 2021-01-21T17:12:35.232Z
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// example:
	//
	// 14588*****51688039
	MemberId *string `json:"MemberId,omitempty" xml:"MemberId,omitempty"`
	// example:
	//
	// user1
	MemberName *string   `json:"MemberName,omitempty" xml:"MemberName,omitempty"`
	Roles      []*string `json:"Roles,omitempty" xml:"Roles,omitempty" type:"Repeated"`
	// example:
	//
	// 215139******88039
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
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
	// example:
	//
	// Approved
	ApprovalStatus *string `json:"ApprovalStatus,omitempty" xml:"ApprovalStatus,omitempty"`
	// example:
	//
	// SavedModel
	FormatType *string `json:"FormatType,omitempty" xml:"FormatType,omitempty"`
	// example:
	//
	// TensorFlow
	FrameworkType *string `json:"FrameworkType,omitempty" xml:"FrameworkType,omitempty"`
	// example:
	//
	// key1
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// example:
	//
	// DESC
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// GmtCreateTime
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// example:
	//
	// region=cn-shanghai,workspaceId=13**,kind=PipelineRun,id=run-sakdb****jdf
	SourceId *string `json:"SourceId,omitempty" xml:"SourceId,omitempty"`
	// example:
	//
	// PAIFlow
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// example:
	//
	// 1.0.1
	VersionName *string `json:"VersionName,omitempty" xml:"VersionName,omitempty"`
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
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC***3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 15
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
	// example:
	//
	// AI4D,QuickStart
	Collections *string `json:"Collections,omitempty" xml:"Collections,omitempty"`
	// example:
	//
	// nlp
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// example:
	//
	// key1
	Label     *string `json:"Label,omitempty" xml:"Label,omitempty"`
	ModelName *string `json:"ModelName,omitempty" xml:"ModelName,omitempty"`
	// example:
	//
	// Endpoint
	ModelType *string `json:"ModelType,omitempty" xml:"ModelType,omitempty"`
	// example:
	//
	// DESC
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// example:
	//
	// ModelScope
	Origin *string `json:"Origin,omitempty" xml:"Origin,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// pai
	Provider *string `json:"Provider,omitempty" xml:"Provider,omitempty"`
	// example:
	//
	// nlp
	Query *string `json:"Query,omitempty" xml:"Query,omitempty"`
	// example:
	//
	// GmtCreateTime
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// example:
	//
	// text-classification
	Task *string `json:"Task,omitempty" xml:"Task,omitempty"`
	// example:
	//
	// 324**
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
	Models []*Model `json:"Models,omitempty" xml:"Models,omitempty" type:"Repeated"`
	// example:
	//
	// 5A14FA81-DD4E-******-6343FE44B941
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 15
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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
	// example:
	//
	// 2AE63638-5420-56DC-B******8174039A0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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
	// example:
	//
	// PaiDLC:GetTensorboard
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
	// example:
	//
	// PRIVATE
	Accessibility *string `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	// example:
	//
	// CREATOR
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
	// example:
	//
	// PAI_isolate
	ProductCodes *string `json:"ProductCodes,omitempty" xml:"ProductCodes,omitempty"`
	// example:
	//
	// oss
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
	Products []*ListProductsResponseBodyProducts `json:"Products,omitempty" xml:"Products,omitempty" type:"Repeated"`
	// example:
	//
	// 1e195c5116124202371861018d5bde
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
	HasPermissionToPurchase *bool `json:"HasPermissionToPurchase,omitempty" xml:"HasPermissionToPurchase,omitempty"`
	// example:
	//
	// true
	IsPurchased *bool `json:"IsPurchased,omitempty" xml:"IsPurchased,omitempty"`
	// example:
	//
	// DataWorks_isolate
	ProductCode       *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	ProductInstanceId *string `json:"ProductInstanceId,omitempty" xml:"ProductInstanceId,omitempty"`
	// example:
	//
	// https://common-buy.aliy
	PurchaseUrl *string `json:"PurchaseUrl,omitempty" xml:"PurchaseUrl,omitempty"`
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
	// example:
	//
	// true
	IsOpen  *bool   `json:"IsOpen,omitempty" xml:"IsOpen,omitempty"`
	OpenUrl *string `json:"OpenUrl,omitempty" xml:"OpenUrl,omitempty"`
	// example:
	//
	// oss
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
	// example:
	//
	// quota-name
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
	Quotas []*ListQuotasResponseBodyQuotas `json:"Quotas,omitempty" xml:"Quotas,omitempty" type:"Repeated"`
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// example:
	//
	// 1828233
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// isolate
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// example:
	//
	// quota-name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// MaxCompute_share
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// example:
	//
	// MaxCompute
	QuotaType *string `json:"QuotaType,omitempty" xml:"QuotaType,omitempty"`
	// example:
	//
	// {\\"cu\\":\\"11500\\",\\"minCu\\":\\"2300\\",\\"parentId\\":\\"0\\"}
	Specs []*ListQuotasResponseBodyQuotasSpecs `json:"Specs,omitempty" xml:"Specs,omitempty" type:"Repeated"`
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
	// example:
	//
	// cu
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// string
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// 11500
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
	// example:
	//
	// group
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	Labels    *string `json:"Labels,omitempty" xml:"Labels,omitempty"`
	// example:
	//
	// ListResourceByWorkspace
	Option *string `json:"Option,omitempty" xml:"Option,omitempty"`
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// MaxCompute
	ProductTypes *string `json:"ProductTypes,omitempty" xml:"ProductTypes,omitempty"`
	QuotaIds     *string `json:"QuotaIds,omitempty" xml:"QuotaIds,omitempty"`
	// example:
	//
	// resource
	ResourceName  *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	ResourceTypes *string `json:"ResourceTypes,omitempty" xml:"ResourceTypes,omitempty"`
	// example:
	//
	// true
	Verbose       *bool   `json:"Verbose,omitempty" xml:"Verbose,omitempty"`
	VerboseFields *string `json:"VerboseFields,omitempty" xml:"VerboseFields,omitempty"`
	// example:
	//
	// 123
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
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
	// example:
	//
	// 1e195c5116124202371861018d5bde
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Resources []*ListResourcesResponseBodyResources `json:"Resources,omitempty" xml:"Resources,omitempty" type:"Repeated"`
	// example:
	//
	// 2
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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
	Encryption *ListResourcesResponseBodyResourcesEncryption `json:"Encryption,omitempty" xml:"Encryption,omitempty" type:"Struct"`
	// example:
	//
	// prod
	EnvType  *string                                     `json:"EnvType,omitempty" xml:"EnvType,omitempty"`
	Executor *ListResourcesResponseBodyResourcesExecutor `json:"Executor,omitempty" xml:"Executor,omitempty" type:"Struct"`
	// example:
	//
	// 2021-01-21T17:12:35.232Z
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// example:
	//
	// groupName
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// example:
	//
	// 123
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// true
	IsDefault *bool                                       `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	Labels    []*ListResourcesResponseBodyResourcesLabels `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	// example:
	//
	// ResourceName
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// MaxCompute
	ProductType  *string                                     `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	Quotas       []*ListResourcesResponseBodyResourcesQuotas `json:"Quotas,omitempty" xml:"Quotas,omitempty" type:"Repeated"`
	ResourceType *string                                     `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// example:
	//
	// 对于MaxCompute {"Endpoint": "odps.alibaba-inc.com", "Project": "mignshi"}
	Spec map[string]interface{} `json:"Spec,omitempty" xml:"Spec,omitempty"`
	// example:
	//
	// 123
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
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
	// example:
	//
	// cpu
	CardType    *string `json:"CardType,omitempty" xml:"CardType,omitempty"`
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// example:
	//
	// 123
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// develop
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// example:
	//
	// QuotaName
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// MaxCompute_isolate
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// example:
	//
	// MaxCompute
	QuotaType *string `json:"QuotaType,omitempty" xml:"QuotaType,omitempty"`
	// example:
	//
	// {\\"cu\\":\\"11500\\",\\"minCu\\":\\"2300\\",\\"parentId\\":\\"0\\"}
	Specs []*ListResourcesResponseBodyResourcesQuotasSpecs `json:"Specs,omitempty" xml:"Specs,omitempty" type:"Repeated"`
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
	// example:
	//
	// cu
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 11500
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

type ListRunMetricsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// loss
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// 100
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// 0
	PageToken *int64 `json:"PageToken,omitempty" xml:"PageToken,omitempty"`
}

func (s ListRunMetricsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListRunMetricsRequest) GoString() string {
	return s.String()
}

func (s *ListRunMetricsRequest) SetKey(v string) *ListRunMetricsRequest {
	s.Key = &v
	return s
}

func (s *ListRunMetricsRequest) SetMaxResults(v int64) *ListRunMetricsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListRunMetricsRequest) SetPageToken(v int64) *ListRunMetricsRequest {
	s.PageToken = &v
	return s
}

type ListRunMetricsResponseBody struct {
	Metrics []*RunMetric `json:"Metrics,omitempty" xml:"Metrics,omitempty" type:"Repeated"`
	// example:
	//
	// 0
	NextPageToken *int64 `json:"NextPageToken,omitempty" xml:"NextPageToken,omitempty"`
	// Id of the request
	//
	// example:
	//
	// ADF6D849-*****-7E7030F0CE53
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListRunMetricsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListRunMetricsResponseBody) GoString() string {
	return s.String()
}

func (s *ListRunMetricsResponseBody) SetMetrics(v []*RunMetric) *ListRunMetricsResponseBody {
	s.Metrics = v
	return s
}

func (s *ListRunMetricsResponseBody) SetNextPageToken(v int64) *ListRunMetricsResponseBody {
	s.NextPageToken = &v
	return s
}

func (s *ListRunMetricsResponseBody) SetRequestId(v string) *ListRunMetricsResponseBody {
	s.RequestId = &v
	return s
}

type ListRunMetricsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListRunMetricsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListRunMetricsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListRunMetricsResponse) GoString() string {
	return s.String()
}

func (s *ListRunMetricsResponse) SetHeaders(v map[string]*string) *ListRunMetricsResponse {
	s.Headers = v
	return s
}

func (s *ListRunMetricsResponse) SetStatusCode(v int32) *ListRunMetricsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRunMetricsResponse) SetBody(v *ListRunMetricsResponseBody) *ListRunMetricsResponse {
	s.Body = v
	return s
}

type ListRunsRequest struct {
	// example:
	//
	// exp-1zpfthdx******
	ExperimentId *string `json:"ExperimentId,omitempty" xml:"ExperimentId,omitempty"`
	// example:
	//
	// 2021-01-30T12:51:33.028Z
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// example:
	//
	// is_evaluation:true
	Labels *string `json:"Labels,omitempty" xml:"Labels,omitempty"`
	// example:
	//
	// 10
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// myName
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// DESC
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// example:
	//
	// GmtCreateTime DESC,Name ASC
	OrderBy *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 0
	PageToken *int64 `json:"PageToken,omitempty" xml:"PageToken,omitempty"`
	// example:
	//
	// GmtCreateTime
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// example:
	//
	// job-rbvg5wzlj****
	SourceId *string `json:"SourceId,omitempty" xml:"SourceId,omitempty"`
	// example:
	//
	// TrainingService
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// example:
	//
	// true
	Verbose *bool `json:"Verbose,omitempty" xml:"Verbose,omitempty"`
	// example:
	//
	// 22840
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListRunsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListRunsRequest) GoString() string {
	return s.String()
}

func (s *ListRunsRequest) SetExperimentId(v string) *ListRunsRequest {
	s.ExperimentId = &v
	return s
}

func (s *ListRunsRequest) SetGmtCreateTime(v string) *ListRunsRequest {
	s.GmtCreateTime = &v
	return s
}

func (s *ListRunsRequest) SetLabels(v string) *ListRunsRequest {
	s.Labels = &v
	return s
}

func (s *ListRunsRequest) SetMaxResults(v int64) *ListRunsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListRunsRequest) SetName(v string) *ListRunsRequest {
	s.Name = &v
	return s
}

func (s *ListRunsRequest) SetOrder(v string) *ListRunsRequest {
	s.Order = &v
	return s
}

func (s *ListRunsRequest) SetOrderBy(v string) *ListRunsRequest {
	s.OrderBy = &v
	return s
}

func (s *ListRunsRequest) SetPageNumber(v int64) *ListRunsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListRunsRequest) SetPageSize(v int64) *ListRunsRequest {
	s.PageSize = &v
	return s
}

func (s *ListRunsRequest) SetPageToken(v int64) *ListRunsRequest {
	s.PageToken = &v
	return s
}

func (s *ListRunsRequest) SetSortBy(v string) *ListRunsRequest {
	s.SortBy = &v
	return s
}

func (s *ListRunsRequest) SetSourceId(v string) *ListRunsRequest {
	s.SourceId = &v
	return s
}

func (s *ListRunsRequest) SetSourceType(v string) *ListRunsRequest {
	s.SourceType = &v
	return s
}

func (s *ListRunsRequest) SetVerbose(v bool) *ListRunsRequest {
	s.Verbose = &v
	return s
}

func (s *ListRunsRequest) SetWorkspaceId(v string) *ListRunsRequest {
	s.WorkspaceId = &v
	return s
}

type ListRunsResponseBody struct {
	// example:
	//
	// 0
	NextPageToken *int64 `json:"NextPageToken,omitempty" xml:"NextPageToken,omitempty"`
	Runs          []*Run `json:"Runs,omitempty" xml:"Runs,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// Id of the request
	//
	// example:
	//
	// ADF6D849-*****-7E7030F0CE53
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListRunsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListRunsResponseBody) GoString() string {
	return s.String()
}

func (s *ListRunsResponseBody) SetNextPageToken(v int64) *ListRunsResponseBody {
	s.NextPageToken = &v
	return s
}

func (s *ListRunsResponseBody) SetRuns(v []*Run) *ListRunsResponseBody {
	s.Runs = v
	return s
}

func (s *ListRunsResponseBody) SetTotalCount(v int64) *ListRunsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListRunsResponseBody) SetRequestId(v string) *ListRunsResponseBody {
	s.RequestId = &v
	return s
}

type ListRunsResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListRunsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListRunsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListRunsResponse) GoString() string {
	return s.String()
}

func (s *ListRunsResponse) SetHeaders(v map[string]*string) *ListRunsResponse {
	s.Headers = v
	return s
}

func (s *ListRunsResponse) SetStatusCode(v int32) *ListRunsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRunsResponse) SetBody(v *ListRunsResponseBody) *ListRunsResponse {
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
	// example:
	//
	// 1e195c5116124202371861018d5bde
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 2
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
	// example:
	//
	// 1611******3000
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// example:
	//
	// she******mo
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
	// example:
	//
	// Id
	Fields *string `json:"Fields,omitempty" xml:"Fields,omitempty"`
	// example:
	//
	// PAI
	ModuleList *string `json:"ModuleList,omitempty" xml:"ModuleList,omitempty"`
	// example:
	//
	// GetWorkspaces
	Option *string `json:"Option,omitempty" xml:"Option,omitempty"`
	// example:
	//
	// ASC
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// GmtCreateTime
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// example:
	//
	// ENABLED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// true
	Verbose *bool `json:"Verbose,omitempty" xml:"Verbose,omitempty"`
	// example:
	//
	// 123,234
	WorkspaceIds *string `json:"WorkspaceIds,omitempty" xml:"WorkspaceIds,omitempty"`
	// example:
	//
	// abc
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
	// example:
	//
	// 8D7B2E70-F770-505B-A672-09F1D8F2EC1E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// {
	//
	//    "MaxCompute_share": 1,
	//
	//    "MaxCompute_isolate": 1,
	//
	//    "DLC_share": 1
	//
	// }
	ResourceLimits map[string]interface{} `json:"ResourceLimits,omitempty" xml:"ResourceLimits,omitempty"`
	// example:
	//
	// 1
	TotalCount *int64                                  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	Workspaces []*ListWorkspacesResponseBodyWorkspaces `json:"Workspaces,omitempty" xml:"Workspaces,omitempty" type:"Repeated"`
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
	AdminNames []*string `json:"AdminNames,omitempty" xml:"AdminNames,omitempty" type:"Repeated"`
	// example:
	//
	// 122424353535
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// example:
	//
	// workspace description example
	Description *string   `json:"Description,omitempty" xml:"Description,omitempty"`
	EnvTypes    []*string `json:"EnvTypes,omitempty" xml:"EnvTypes,omitempty" type:"Repeated"`
	// example:
	//
	// {"TenantId": "4286******98"}
	ExtraInfos map[string]interface{} `json:"ExtraInfos,omitempty" xml:"ExtraInfos,omitempty"`
	// example:
	//
	// 2021-01-21T17:12:35.232Z
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// example:
	//
	// 2021-01-21T17:12:35.232Z
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// example:
	//
	// false
	IsDefault *bool `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	// example:
	//
	// ENABLED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 123
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
	// example:
	//
	// workspace-example
	WorkspaceName *string `json:"WorkspaceName,omitempty" xml:"WorkspaceName,omitempty"`
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

type LogRunMetricsRequest struct {
	Metrics []*RunMetric `json:"Metrics,omitempty" xml:"Metrics,omitempty" type:"Repeated"`
}

func (s LogRunMetricsRequest) String() string {
	return tea.Prettify(s)
}

func (s LogRunMetricsRequest) GoString() string {
	return s.String()
}

func (s *LogRunMetricsRequest) SetMetrics(v []*RunMetric) *LogRunMetricsRequest {
	s.Metrics = v
	return s
}

type LogRunMetricsResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// ADF6D849-*****-7E7030F0CE53
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s LogRunMetricsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s LogRunMetricsResponseBody) GoString() string {
	return s.String()
}

func (s *LogRunMetricsResponseBody) SetRequestId(v string) *LogRunMetricsResponseBody {
	s.RequestId = &v
	return s
}

type LogRunMetricsResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *LogRunMetricsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s LogRunMetricsResponse) String() string {
	return tea.Prettify(s)
}

func (s LogRunMetricsResponse) GoString() string {
	return s.String()
}

func (s *LogRunMetricsResponse) SetHeaders(v map[string]*string) *LogRunMetricsResponse {
	s.Headers = v
	return s
}

func (s *LogRunMetricsResponse) SetStatusCode(v int32) *LogRunMetricsResponse {
	s.StatusCode = &v
	return s
}

func (s *LogRunMetricsResponse) SetBody(v *LogRunMetricsResponseBody) *LogRunMetricsResponse {
	s.Body = v
	return s
}

type PublishCodeSourceResponseBody struct {
	// example:
	//
	// code-a797*******
	CodeSourceId *string `json:"CodeSourceId,omitempty" xml:"CodeSourceId,omitempty"`
	// example:
	//
	// 5A14FA81-DD4E-******-6343FE44B941
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// example:
	//
	// A0F049F0-8D69-5BAC-8F10-B******A34C
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
	// example:
	//
	// image-dk******fa
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// example:
	//
	// A0F049F0-8D69-5BAC-8F10-B******A34C
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
	// example:
	//
	// 5A14FA81-DD4E-******-6343FE44B941
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
	// example:
	//
	// 5A14FA81-DD4E-******-6343FE44B941
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
	// example:
	//
	// 5A14FA81-DD4E-******-6343FE44B941
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

type SetExperimentLabelsRequest struct {
	Labels []*LabelInfo `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
}

func (s SetExperimentLabelsRequest) String() string {
	return tea.Prettify(s)
}

func (s SetExperimentLabelsRequest) GoString() string {
	return s.String()
}

func (s *SetExperimentLabelsRequest) SetLabels(v []*LabelInfo) *SetExperimentLabelsRequest {
	s.Labels = v
	return s
}

type SetExperimentLabelsResponseBody struct {
	// example:
	//
	// 5A14FA81-DD4E-******-6343FE44B941
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetExperimentLabelsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetExperimentLabelsResponseBody) GoString() string {
	return s.String()
}

func (s *SetExperimentLabelsResponseBody) SetRequestId(v string) *SetExperimentLabelsResponseBody {
	s.RequestId = &v
	return s
}

type SetExperimentLabelsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetExperimentLabelsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetExperimentLabelsResponse) String() string {
	return tea.Prettify(s)
}

func (s SetExperimentLabelsResponse) GoString() string {
	return s.String()
}

func (s *SetExperimentLabelsResponse) SetHeaders(v map[string]*string) *SetExperimentLabelsResponse {
	s.Headers = v
	return s
}

func (s *SetExperimentLabelsResponse) SetStatusCode(v int32) *SetExperimentLabelsResponse {
	s.StatusCode = &v
	return s
}

func (s *SetExperimentLabelsResponse) SetBody(v *SetExperimentLabelsResponseBody) *SetExperimentLabelsResponse {
	s.Body = v
	return s
}

type UpdateDatasetRequest struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// myName
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// {
	//
	//   "mountPath": "/mnt/data/"
	//
	// }
	Options *string `json:"Options,omitempty" xml:"Options,omitempty"`
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
	// example:
	//
	// 5A14FA81-DD4E-******-6343FE44B941
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

type UpdateDatasetVersionRequest struct {
	// example:
	//
	// 100
	DataCount *int64 `json:"DataCount,omitempty" xml:"DataCount,omitempty"`
	// example:
	//
	// 100000
	DataSize    *int64  `json:"DataSize,omitempty" xml:"DataSize,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// {
	//
	//   "mountPath": "/mnt/data/"
	//
	// }
	Options *string `json:"Options,omitempty" xml:"Options,omitempty"`
}

func (s UpdateDatasetVersionRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateDatasetVersionRequest) GoString() string {
	return s.String()
}

func (s *UpdateDatasetVersionRequest) SetDataCount(v int64) *UpdateDatasetVersionRequest {
	s.DataCount = &v
	return s
}

func (s *UpdateDatasetVersionRequest) SetDataSize(v int64) *UpdateDatasetVersionRequest {
	s.DataSize = &v
	return s
}

func (s *UpdateDatasetVersionRequest) SetDescription(v string) *UpdateDatasetVersionRequest {
	s.Description = &v
	return s
}

func (s *UpdateDatasetVersionRequest) SetOptions(v string) *UpdateDatasetVersionRequest {
	s.Options = &v
	return s
}

type UpdateDatasetVersionResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateDatasetVersionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateDatasetVersionResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDatasetVersionResponseBody) SetRequestId(v string) *UpdateDatasetVersionResponseBody {
	s.RequestId = &v
	return s
}

type UpdateDatasetVersionResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateDatasetVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateDatasetVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateDatasetVersionResponse) GoString() string {
	return s.String()
}

func (s *UpdateDatasetVersionResponse) SetHeaders(v map[string]*string) *UpdateDatasetVersionResponse {
	s.Headers = v
	return s
}

func (s *UpdateDatasetVersionResponse) SetStatusCode(v int32) *UpdateDatasetVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateDatasetVersionResponse) SetBody(v *UpdateDatasetVersionResponseBody) *UpdateDatasetVersionResponse {
	s.Body = v
	return s
}

type UpdateDefaultWorkspaceRequest struct {
	// example:
	//
	// 12345
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
	// example:
	//
	// 17915******4216
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

type UpdateExperimentRequest struct {
	// example:
	//
	// PRIVATE
	Accessibility *string `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	// 名称
	//
	// example:
	//
	// myName
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s UpdateExperimentRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateExperimentRequest) GoString() string {
	return s.String()
}

func (s *UpdateExperimentRequest) SetAccessibility(v string) *UpdateExperimentRequest {
	s.Accessibility = &v
	return s
}

func (s *UpdateExperimentRequest) SetName(v string) *UpdateExperimentRequest {
	s.Name = &v
	return s
}

type UpdateExperimentResponseBody struct {
	// example:
	//
	// 5A14FA81-DD4E-******-6343FE44B941
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateExperimentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateExperimentResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateExperimentResponseBody) SetRequestId(v string) *UpdateExperimentResponseBody {
	s.RequestId = &v
	return s
}

type UpdateExperimentResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateExperimentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateExperimentResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateExperimentResponse) GoString() string {
	return s.String()
}

func (s *UpdateExperimentResponse) SetHeaders(v map[string]*string) *UpdateExperimentResponse {
	s.Headers = v
	return s
}

func (s *UpdateExperimentResponse) SetStatusCode(v int32) *UpdateExperimentResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateExperimentResponse) SetBody(v *UpdateExperimentResponseBody) *UpdateExperimentResponse {
	s.Body = v
	return s
}

type UpdateModelRequest struct {
	// example:
	//
	// PUBLIC
	Accessibility *string `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	// example:
	//
	// nlp
	Domain           *string                `json:"Domain,omitempty" xml:"Domain,omitempty"`
	ExtraInfo        map[string]interface{} `json:"ExtraInfo,omitempty" xml:"ExtraInfo,omitempty"`
	ModelDescription *string                `json:"ModelDescription,omitempty" xml:"ModelDescription,omitempty"`
	// example:
	//
	// https://*.md
	ModelDoc  *string `json:"ModelDoc,omitempty" xml:"ModelDoc,omitempty"`
	ModelName *string `json:"ModelName,omitempty" xml:"ModelName,omitempty"`
	// example:
	//
	// Checkpoint
	ModelType *string `json:"ModelType,omitempty" xml:"ModelType,omitempty"`
	// example:
	//
	// 101
	OrderNumber *int64 `json:"OrderNumber,omitempty" xml:"OrderNumber,omitempty"`
	// example:
	//
	// ModelScope
	Origin *string `json:"Origin,omitempty" xml:"Origin,omitempty"`
	// example:
	//
	// text-classification
	Task *string `json:"Task,omitempty" xml:"Task,omitempty"`
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
	// example:
	//
	// A0F049F0-8D69-5BAC-8F10-B******A34C
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
	// example:
	//
	// Approved
	ApprovalStatus *string `json:"ApprovalStatus,omitempty" xml:"ApprovalStatus,omitempty"`
	// example:
	//
	// {}
	CompressionSpec map[string]interface{} `json:"CompressionSpec,omitempty" xml:"CompressionSpec,omitempty"`
	// example:
	//
	// {}
	EvaluationSpec map[string]interface{} `json:"EvaluationSpec,omitempty" xml:"EvaluationSpec,omitempty"`
	// example:
	//
	// {}
	ExtraInfo map[string]interface{} `json:"ExtraInfo,omitempty" xml:"ExtraInfo,omitempty"`
	// example:
	//
	// {     "processor": "tensorflow_gpu_1.12" }
	InferenceSpec map[string]interface{} `json:"InferenceSpec,omitempty" xml:"InferenceSpec,omitempty"`
	// example:
	//
	// {}
	Metrics map[string]interface{} `json:"Metrics,omitempty" xml:"Metrics,omitempty"`
	// example:
	//
	// {}
	Options *string `json:"Options,omitempty" xml:"Options,omitempty"`
	// example:
	//
	// region=cn-shanghai,workspaceId=13**,kind=PipelineRun,id=run-sakdb****jdf
	SourceId *string `json:"SourceId,omitempty" xml:"SourceId,omitempty"`
	// example:
	//
	// PAIFlow
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// example:
	//
	// {}
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

func (s *UpdateModelVersionRequest) SetCompressionSpec(v map[string]interface{}) *UpdateModelVersionRequest {
	s.CompressionSpec = v
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
	// example:
	//
	// D5BFFEE3-6025-443F-8A03-02D61***C4B9
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

type UpdateRunRequest struct {
	Labels []*Label `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	// example:
	//
	// myName
	Name   *string     `json:"Name,omitempty" xml:"Name,omitempty"`
	Params []*RunParam `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
}

func (s UpdateRunRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateRunRequest) GoString() string {
	return s.String()
}

func (s *UpdateRunRequest) SetLabels(v []*Label) *UpdateRunRequest {
	s.Labels = v
	return s
}

func (s *UpdateRunRequest) SetName(v string) *UpdateRunRequest {
	s.Name = &v
	return s
}

func (s *UpdateRunRequest) SetParams(v []*RunParam) *UpdateRunRequest {
	s.Params = v
	return s
}

type UpdateRunResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// ADF6D849-*****-7E7030F0CE53
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateRunResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateRunResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateRunResponseBody) SetRequestId(v string) *UpdateRunResponseBody {
	s.RequestId = &v
	return s
}

type UpdateRunResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateRunResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateRunResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateRunResponse) GoString() string {
	return s.String()
}

func (s *UpdateRunResponse) SetHeaders(v map[string]*string) *UpdateRunResponse {
	s.Headers = v
	return s
}

func (s *UpdateRunResponse) SetStatusCode(v int32) *UpdateRunResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateRunResponse) SetBody(v *UpdateRunResponseBody) *UpdateRunResponse {
	s.Body = v
	return s
}

type UpdateWorkspaceRequest struct {
	// example:
	//
	// Description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// workspace-example
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
	// example:
	//
	// 5A14FA81-DD4E-******-6343FE44B941
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
	// example:
	//
	// group-kjds******sd
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// example:
	//
	// true
	IsDefault *bool                                   `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	Labels    []*UpdateWorkspaceResourceRequestLabels `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	// example:
	//
	// MaxCompute
	ProductType  *string                `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	ResourceIds  []*string              `json:"ResourceIds,omitempty" xml:"ResourceIds,omitempty" type:"Repeated"`
	ResourceType *string                `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Spec         map[string]interface{} `json:"Spec,omitempty" xml:"Spec,omitempty"`
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
	// example:
	//
	// 5A14FA81-DD4E-******-6343FE44B941
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

// Summary:
//
// 增加 Image
//
// @param request - AddImageRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddImageResponse
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

// Summary:
//
// 增加 Image
//
// @param request - AddImageRequest
//
// @return AddImageResponse
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

// Summary:
//
// 增加 Image 的标签
//
// @param request - AddImageLabelsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddImageLabelsResponse
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

// Summary:
//
// 增加 Image 的标签
//
// @param request - AddImageLabelsRequest
//
// @return AddImageLabelsResponse
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

// Summary:
//
// 增加成员角色
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddMemberRoleResponse
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

// Summary:
//
// 增加成员角色
//
// @return AddMemberRoleResponse
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

// Summary:
//
// 创建一个代码源配置
//
// @param request - CreateCodeSourceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCodeSourceResponse
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

// Summary:
//
// 创建一个代码源配置
//
// @param request - CreateCodeSourceRequest
//
// @return CreateCodeSourceResponse
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

// Summary:
//
// 创建数据集
//
// @param request - CreateDatasetRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDatasetResponse
func (client *Client) CreateDatasetWithOptions(request *CreateDatasetRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateDatasetResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Accessibility)) {
		body["Accessibility"] = request.Accessibility
	}

	if !tea.BoolValue(util.IsUnset(request.DataCount)) {
		body["DataCount"] = request.DataCount
	}

	if !tea.BoolValue(util.IsUnset(request.DataSize)) {
		body["DataSize"] = request.DataSize
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

	if !tea.BoolValue(util.IsUnset(request.SourceDatasetId)) {
		body["SourceDatasetId"] = request.SourceDatasetId
	}

	if !tea.BoolValue(util.IsUnset(request.SourceDatasetVersion)) {
		body["SourceDatasetVersion"] = request.SourceDatasetVersion
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

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["UserId"] = request.UserId
	}

	if !tea.BoolValue(util.IsUnset(request.VersionDescription)) {
		body["VersionDescription"] = request.VersionDescription
	}

	if !tea.BoolValue(util.IsUnset(request.VersionLabels)) {
		body["VersionLabels"] = request.VersionLabels
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

// Summary:
//
// 创建数据集
//
// @param request - CreateDatasetRequest
//
// @return CreateDatasetResponse
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

// Summary:
//
// 创建或更新 Dataset 的标签
//
// @param request - CreateDatasetLabelsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDatasetLabelsResponse
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

// Summary:
//
// 创建或更新 Dataset 的标签
//
// @param request - CreateDatasetLabelsRequest
//
// @return CreateDatasetLabelsResponse
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

// Summary:
//
// 创建数据集版本
//
// @param request - CreateDatasetVersionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDatasetVersionResponse
func (client *Client) CreateDatasetVersionWithOptions(DatasetId *string, request *CreateDatasetVersionRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateDatasetVersionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DataCount)) {
		body["DataCount"] = request.DataCount
	}

	if !tea.BoolValue(util.IsUnset(request.DataSize)) {
		body["DataSize"] = request.DataSize
	}

	if !tea.BoolValue(util.IsUnset(request.DataSourceType)) {
		body["DataSourceType"] = request.DataSourceType
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Labels)) {
		body["Labels"] = request.Labels
	}

	if !tea.BoolValue(util.IsUnset(request.Options)) {
		body["Options"] = request.Options
	}

	if !tea.BoolValue(util.IsUnset(request.Property)) {
		body["Property"] = request.Property
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

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateDatasetVersion"),
		Version:     tea.String("2021-02-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/datasets/" + tea.StringValue(openapiutil.GetEncodeParam(DatasetId)) + "/versions"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateDatasetVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建数据集版本
//
// @param request - CreateDatasetVersionRequest
//
// @return CreateDatasetVersionResponse
func (client *Client) CreateDatasetVersion(DatasetId *string, request *CreateDatasetVersionRequest) (_result *CreateDatasetVersionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateDatasetVersionResponse{}
	_body, _err := client.CreateDatasetVersionWithOptions(DatasetId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建数据集版本的标签
//
// @param request - CreateDatasetVersionLabelsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDatasetVersionLabelsResponse
func (client *Client) CreateDatasetVersionLabelsWithOptions(DatasetId *string, VersionName *string, request *CreateDatasetVersionLabelsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateDatasetVersionLabelsResponse, _err error) {
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
		Action:      tea.String("CreateDatasetVersionLabels"),
		Version:     tea.String("2021-02-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/datasets/" + tea.StringValue(openapiutil.GetEncodeParam(DatasetId)) + "/versions/" + tea.StringValue(openapiutil.GetEncodeParam(VersionName)) + "/labels"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateDatasetVersionLabelsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建数据集版本的标签
//
// @param request - CreateDatasetVersionLabelsRequest
//
// @return CreateDatasetVersionLabelsResponse
func (client *Client) CreateDatasetVersionLabels(DatasetId *string, VersionName *string, request *CreateDatasetVersionLabelsRequest) (_result *CreateDatasetVersionLabelsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateDatasetVersionLabelsResponse{}
	_body, _err := client.CreateDatasetVersionLabelsWithOptions(DatasetId, VersionName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建实验
//
// @param request - CreateExperimentRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateExperimentResponse
func (client *Client) CreateExperimentWithOptions(request *CreateExperimentRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateExperimentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Accessibility)) {
		body["Accessibility"] = request.Accessibility
	}

	if !tea.BoolValue(util.IsUnset(request.ArtifactUri)) {
		body["ArtifactUri"] = request.ArtifactUri
	}

	if !tea.BoolValue(util.IsUnset(request.Labels)) {
		body["Labels"] = request.Labels
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateExperiment"),
		Version:     tea.String("2021-02-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/experiments"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateExperimentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建实验
//
// @param request - CreateExperimentRequest
//
// @return CreateExperimentResponse
func (client *Client) CreateExperiment(request *CreateExperimentRequest) (_result *CreateExperimentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateExperimentResponse{}
	_body, _err := client.CreateExperimentWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建成员
//
// @param request - CreateMemberRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateMemberResponse
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

// Summary:
//
// 创建成员
//
// @param request - CreateMemberRequest
//
// @return CreateMemberResponse
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

// Summary:
//
// 创建模型
//
// @param request - CreateModelRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateModelResponse
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

// Summary:
//
// 创建模型
//
// @param request - CreateModelRequest
//
// @return CreateModelResponse
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

// Summary:
//
// 创建或更新模型的标签
//
// @param request - CreateModelLabelsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateModelLabelsResponse
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

// Summary:
//
// 创建或更新模型的标签
//
// @param request - CreateModelLabelsRequest
//
// @return CreateModelLabelsResponse
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

// Summary:
//
// 创建模型版本
//
// @param request - CreateModelVersionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateModelVersionResponse
func (client *Client) CreateModelVersionWithOptions(ModelId *string, request *CreateModelVersionRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateModelVersionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApprovalStatus)) {
		body["ApprovalStatus"] = request.ApprovalStatus
	}

	if !tea.BoolValue(util.IsUnset(request.CompressionSpec)) {
		body["CompressionSpec"] = request.CompressionSpec
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

// Summary:
//
// 创建模型版本
//
// @param request - CreateModelVersionRequest
//
// @return CreateModelVersionResponse
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

// Summary:
//
// 创建或更新模型版本的标签
//
// @param request - CreateModelVersionLabelsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateModelVersionLabelsResponse
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

// Summary:
//
// 创建或更新模型版本的标签
//
// @param request - CreateModelVersionLabelsRequest
//
// @return CreateModelVersionLabelsResponse
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

// Summary:
//
// 创建产品订单
//
// @param request - CreateProductOrdersRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateProductOrdersResponse
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

// Summary:
//
// 创建产品订单
//
// @param request - CreateProductOrdersRequest
//
// @return CreateProductOrdersResponse
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

// Summary:
//
// 创建一次运行
//
// @param request - CreateRunRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateRunResponse
func (client *Client) CreateRunWithOptions(request *CreateRunRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateRunResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ExperimentId)) {
		body["ExperimentId"] = request.ExperimentId
	}

	if !tea.BoolValue(util.IsUnset(request.Labels)) {
		body["Labels"] = request.Labels
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Params)) {
		body["Params"] = request.Params
	}

	if !tea.BoolValue(util.IsUnset(request.SourceId)) {
		body["SourceId"] = request.SourceId
	}

	if !tea.BoolValue(util.IsUnset(request.SourceType)) {
		body["SourceType"] = request.SourceType
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateRun"),
		Version:     tea.String("2021-02-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/runs"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateRunResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建一次运行
//
// @param request - CreateRunRequest
//
// @return CreateRunResponse
func (client *Client) CreateRun(request *CreateRunRequest) (_result *CreateRunResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateRunResponse{}
	_body, _err := client.CreateRunWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建工作空间
//
// @param request - CreateWorkspaceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateWorkspaceResponse
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

// Summary:
//
// 创建工作空间
//
// @param request - CreateWorkspaceRequest
//
// @return CreateWorkspaceResponse
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

// Summary:
//
// 创建资源
//
// @param request - CreateWorkspaceResourceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateWorkspaceResourceResponse
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

// Summary:
//
// 创建资源
//
// @param request - CreateWorkspaceResourceRequest
//
// @return CreateWorkspaceResourceResponse
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

// Summary:
//
// 删除一个代码源配置
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCodeSourceResponse
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

// Summary:
//
// 删除一个代码源配置
//
// @return DeleteCodeSourceResponse
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

// Summary:
//
// 删除数据集
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDatasetResponse
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

// Summary:
//
// 删除数据集
//
// @return DeleteDatasetResponse
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

// Summary:
//
// 删除 Dataset 的标签
//
// @param request - DeleteDatasetLabelsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDatasetLabelsResponse
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

// Summary:
//
// 删除 Dataset 的标签
//
// @param request - DeleteDatasetLabelsRequest
//
// @return DeleteDatasetLabelsResponse
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

// Summary:
//
// 删除指定版本的数据集信息，如果删除的版本是该数据集的仅存版本，版本删除后会联动删除dataset 表中的数据集信息
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDatasetVersionResponse
func (client *Client) DeleteDatasetVersionWithOptions(DatasetId *string, VersionName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteDatasetVersionResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteDatasetVersion"),
		Version:     tea.String("2021-02-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/datasets/" + tea.StringValue(openapiutil.GetEncodeParam(DatasetId)) + "/versions/" + tea.StringValue(openapiutil.GetEncodeParam(VersionName))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteDatasetVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除指定版本的数据集信息，如果删除的版本是该数据集的仅存版本，版本删除后会联动删除dataset 表中的数据集信息
//
// @return DeleteDatasetVersionResponse
func (client *Client) DeleteDatasetVersion(DatasetId *string, VersionName *string) (_result *DeleteDatasetVersionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteDatasetVersionResponse{}
	_body, _err := client.DeleteDatasetVersionWithOptions(DatasetId, VersionName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除数据集版本的标签。
//
// @param request - DeleteDatasetVersionLabelsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDatasetVersionLabelsResponse
func (client *Client) DeleteDatasetVersionLabelsWithOptions(DatasetId *string, VersionName *string, request *DeleteDatasetVersionLabelsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteDatasetVersionLabelsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Keys)) {
		query["Keys"] = request.Keys
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteDatasetVersionLabels"),
		Version:     tea.String("2021-02-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/datasets/" + tea.StringValue(openapiutil.GetEncodeParam(DatasetId)) + "/versions/" + tea.StringValue(openapiutil.GetEncodeParam(VersionName)) + "/labels"),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteDatasetVersionLabelsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除数据集版本的标签。
//
// @param request - DeleteDatasetVersionLabelsRequest
//
// @return DeleteDatasetVersionLabelsResponse
func (client *Client) DeleteDatasetVersionLabels(DatasetId *string, VersionName *string, request *DeleteDatasetVersionLabelsRequest) (_result *DeleteDatasetVersionLabelsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteDatasetVersionLabelsResponse{}
	_body, _err := client.DeleteDatasetVersionLabelsWithOptions(DatasetId, VersionName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除实验
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteExperimentResponse
func (client *Client) DeleteExperimentWithOptions(ExperimentId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteExperimentResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteExperiment"),
		Version:     tea.String("2021-02-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/experiments/" + tea.StringValue(openapiutil.GetEncodeParam(ExperimentId))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteExperimentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除实验
//
// @return DeleteExperimentResponse
func (client *Client) DeleteExperiment(ExperimentId *string) (_result *DeleteExperimentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteExperimentResponse{}
	_body, _err := client.DeleteExperimentWithOptions(ExperimentId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除实验标签
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteExperimentLabelResponse
func (client *Client) DeleteExperimentLabelWithOptions(ExperimentId *string, Key *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteExperimentLabelResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteExperimentLabel"),
		Version:     tea.String("2021-02-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/experiments/" + tea.StringValue(openapiutil.GetEncodeParam(ExperimentId)) + "/labels/" + tea.StringValue(openapiutil.GetEncodeParam(Key))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteExperimentLabelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除实验标签
//
// @return DeleteExperimentLabelResponse
func (client *Client) DeleteExperimentLabel(ExperimentId *string, Key *string) (_result *DeleteExperimentLabelResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteExperimentLabelResponse{}
	_body, _err := client.DeleteExperimentLabelWithOptions(ExperimentId, Key, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除工作空间成员
//
// @param request - DeleteMembersRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteMembersResponse
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

// Summary:
//
// 删除工作空间成员
//
// @param request - DeleteMembersRequest
//
// @return DeleteMembersResponse
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

// Summary:
//
// 删除模型
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteModelResponse
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

// Summary:
//
// 删除模型
//
// @return DeleteModelResponse
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

// Summary:
//
// 删除模型的标签
//
// @param request - DeleteModelLabelsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteModelLabelsResponse
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

// Summary:
//
// 删除模型的标签
//
// @param request - DeleteModelLabelsRequest
//
// @return DeleteModelLabelsResponse
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

// Summary:
//
// 删除模型版本
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteModelVersionResponse
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

// Summary:
//
// 删除模型版本
//
// @return DeleteModelVersionResponse
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

// Summary:
//
// 删除模型版本的标签
//
// @param request - DeleteModelVersionLabelsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteModelVersionLabelsResponse
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

// Summary:
//
// 删除模型版本的标签
//
// @param request - DeleteModelVersionLabelsRequest
//
// @return DeleteModelVersionLabelsResponse
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

// Summary:
//
// 删除Run
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteRunResponse
func (client *Client) DeleteRunWithOptions(RunId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteRunResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteRun"),
		Version:     tea.String("2021-02-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/runs/" + tea.StringValue(openapiutil.GetEncodeParam(RunId))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteRunResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除Run
//
// @return DeleteRunResponse
func (client *Client) DeleteRun(RunId *string) (_result *DeleteRunResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteRunResponse{}
	_body, _err := client.DeleteRunWithOptions(RunId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除Run标签
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteRunLabelResponse
func (client *Client) DeleteRunLabelWithOptions(RunId *string, Key *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteRunLabelResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteRunLabel"),
		Version:     tea.String("2021-02-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/runs/" + tea.StringValue(openapiutil.GetEncodeParam(RunId)) + "/labels/" + tea.StringValue(openapiutil.GetEncodeParam(Key))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteRunLabelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除Run标签
//
// @return DeleteRunLabelResponse
func (client *Client) DeleteRunLabel(RunId *string, Key *string) (_result *DeleteRunLabelResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteRunLabelResponse{}
	_body, _err := client.DeleteRunLabelWithOptions(RunId, Key, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除工作空间
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteWorkspaceResponse
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

// Summary:
//
// 删除工作空间
//
// @return DeleteWorkspaceResponse
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

// Summary:
//
// 删除工作空间资源
//
// @param request - DeleteWorkspaceResourceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteWorkspaceResourceResponse
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

// Summary:
//
// 删除工作空间资源
//
// @param request - DeleteWorkspaceResourceRequest
//
// @return DeleteWorkspaceResourceResponse
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

// Summary:
//
// 获取一个代码源配置
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCodeSourceResponse
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

// Summary:
//
// 获取一个代码源配置
//
// @return GetCodeSourceResponse
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

// Summary:
//
// 获取数据集
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDatasetResponse
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

// Summary:
//
// 获取数据集
//
// @return GetDatasetResponse
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

// Summary:
//
// 获取指定版本的数据集信息
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDatasetVersionResponse
func (client *Client) GetDatasetVersionWithOptions(DatasetId *string, VersionName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetDatasetVersionResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetDatasetVersion"),
		Version:     tea.String("2021-02-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/datasets/" + tea.StringValue(openapiutil.GetEncodeParam(DatasetId)) + "/versions/" + tea.StringValue(openapiutil.GetEncodeParam(VersionName))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDatasetVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取指定版本的数据集信息
//
// @return GetDatasetVersionResponse
func (client *Client) GetDatasetVersion(DatasetId *string, VersionName *string) (_result *GetDatasetVersionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetDatasetVersionResponse{}
	_body, _err := client.GetDatasetVersionWithOptions(DatasetId, VersionName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取默认工作空间
//
// @param request - GetDefaultWorkspaceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDefaultWorkspaceResponse
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

// Summary:
//
// 获取默认工作空间
//
// @param request - GetDefaultWorkspaceRequest
//
// @return GetDefaultWorkspaceResponse
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

// Summary:
//
// 获取实验
//
// @param request - GetExperimentRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetExperimentResponse
func (client *Client) GetExperimentWithOptions(ExperimentId *string, request *GetExperimentRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetExperimentResponse, _err error) {
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
		Action:      tea.String("GetExperiment"),
		Version:     tea.String("2021-02-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/experiments/" + tea.StringValue(openapiutil.GetEncodeParam(ExperimentId))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetExperimentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取实验
//
// @param request - GetExperimentRequest
//
// @return GetExperimentResponse
func (client *Client) GetExperiment(ExperimentId *string, request *GetExperimentRequest) (_result *GetExperimentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetExperimentResponse{}
	_body, _err := client.GetExperimentWithOptions(ExperimentId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取镜像
//
// @param request - GetImageRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetImageResponse
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

// Summary:
//
// 获取镜像
//
// @param request - GetImageRequest
//
// @return GetImageResponse
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

// Summary:
//
// 获取成员
//
// @param request - GetMemberRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMemberResponse
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

// Summary:
//
// 获取成员
//
// @param request - GetMemberRequest
//
// @return GetMemberResponse
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

// Summary:
//
// 获取模型
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetModelResponse
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

// Summary:
//
// 获取模型
//
// @return GetModelResponse
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

// Summary:
//
// 获取模型版本
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetModelVersionResponse
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

// Summary:
//
// 获取模型版本
//
// @return GetModelVersionResponse
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

// Summary:
//
// 获取权限，若无权限则返回错误
//
// @param request - GetPermissionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPermissionResponse
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

// Summary:
//
// 获取权限，若无权限则返回错误
//
// @param request - GetPermissionRequest
//
// @return GetPermissionResponse
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

// Summary:
//
// 获取Run详情
//
// @param request - GetRunRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRunResponse
func (client *Client) GetRunWithOptions(RunId *string, request *GetRunRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetRunResponse, _err error) {
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
		Action:      tea.String("GetRun"),
		Version:     tea.String("2021-02-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/runs/" + tea.StringValue(openapiutil.GetEncodeParam(RunId))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetRunResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取Run详情
//
// @param request - GetRunRequest
//
// @return GetRunResponse
func (client *Client) GetRun(RunId *string, request *GetRunRequest) (_result *GetRunResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetRunResponse{}
	_body, _err := client.GetRunWithOptions(RunId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取工作空间
//
// @param request - GetWorkspaceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetWorkspaceResponse
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

// Summary:
//
// 获取工作空间
//
// @param request - GetWorkspaceRequest
//
// @return GetWorkspaceResponse
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

// Summary:
//
// 获取代码源配置列表
//
// @param request - ListCodeSourcesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCodeSourcesResponse
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

// Summary:
//
// 获取代码源配置列表
//
// @param request - ListCodeSourcesRequest
//
// @return ListCodeSourcesResponse
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

// Summary:
//
// 获取数据集版本列表
//
// @param request - ListDatasetVersionsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDatasetVersionsResponse
func (client *Client) ListDatasetVersionsWithOptions(DatasetId *string, request *ListDatasetVersionsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListDatasetVersionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DataSourcesTypes)) {
		query["DataSourcesTypes"] = request.DataSourcesTypes
	}

	if !tea.BoolValue(util.IsUnset(request.LabelKeys)) {
		query["LabelKeys"] = request.LabelKeys
	}

	if !tea.BoolValue(util.IsUnset(request.LableValues)) {
		query["LableValues"] = request.LableValues
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

	if !tea.BoolValue(util.IsUnset(request.SortBy)) {
		query["SortBy"] = request.SortBy
	}

	if !tea.BoolValue(util.IsUnset(request.SourceId)) {
		query["SourceId"] = request.SourceId
	}

	if !tea.BoolValue(util.IsUnset(request.SourceTypes)) {
		query["SourceTypes"] = request.SourceTypes
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListDatasetVersions"),
		Version:     tea.String("2021-02-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/datasets/" + tea.StringValue(openapiutil.GetEncodeParam(DatasetId)) + "/versions"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListDatasetVersionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取数据集版本列表
//
// @param request - ListDatasetVersionsRequest
//
// @return ListDatasetVersionsResponse
func (client *Client) ListDatasetVersions(DatasetId *string, request *ListDatasetVersionsRequest) (_result *ListDatasetVersionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListDatasetVersionsResponse{}
	_body, _err := client.ListDatasetVersionsWithOptions(DatasetId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取数据集列表
//
// @param request - ListDatasetsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDatasetsResponse
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

	if !tea.BoolValue(util.IsUnset(request.SourceDatasetId)) {
		query["SourceDatasetId"] = request.SourceDatasetId
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

// Summary:
//
// 获取数据集列表
//
// @param request - ListDatasetsRequest
//
// @return ListDatasetsResponse
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

// Summary:
//
// 获取实验列表
//
// @param tmpReq - ListExperimentRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListExperimentResponse
func (client *Client) ListExperimentWithOptions(tmpReq *ListExperimentRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListExperimentResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListExperimentShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Options)) {
		request.OptionsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Options, tea.String("Options"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Labels)) {
		query["Labels"] = request.Labels
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.OptionsShrink)) {
		query["Options"] = request.OptionsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Order)) {
		query["Order"] = request.Order
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

	if !tea.BoolValue(util.IsUnset(request.PageToken)) {
		query["PageToken"] = request.PageToken
	}

	if !tea.BoolValue(util.IsUnset(request.SortBy)) {
		query["SortBy"] = request.SortBy
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
		Action:      tea.String("ListExperiment"),
		Version:     tea.String("2021-02-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/experiments"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListExperimentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取实验列表
//
// @param request - ListExperimentRequest
//
// @return ListExperimentResponse
func (client *Client) ListExperiment(request *ListExperimentRequest) (_result *ListExperimentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListExperimentResponse{}
	_body, _err := client.ListExperimentWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 列举标签
//
// @param request - ListImageLabelsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListImageLabelsResponse
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

// Summary:
//
// 列举标签
//
// @param request - ListImageLabelsRequest
//
// @return ListImageLabelsResponse
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

// Summary:
//
// 列举已注册镜像
//
// @param request - ListImagesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListImagesResponse
func (client *Client) ListImagesWithOptions(request *ListImagesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListImagesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Accessibility)) {
		query["Accessibility"] = request.Accessibility
	}

	if !tea.BoolValue(util.IsUnset(request.ImageUri)) {
		query["ImageUri"] = request.ImageUri
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

// Summary:
//
// 列举已注册镜像
//
// @param request - ListImagesRequest
//
// @return ListImagesResponse
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

// Summary:
//
// 列举工作空间成员
//
// @param request - ListMembersRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMembersResponse
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

// Summary:
//
// 列举工作空间成员
//
// @param request - ListMembersRequest
//
// @return ListMembersResponse
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

// Summary:
//
// 获取模型版本列表
//
// @param request - ListModelVersionsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListModelVersionsResponse
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

// Summary:
//
// 获取模型版本列表
//
// @param request - ListModelVersionsRequest
//
// @return ListModelVersionsResponse
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

// Summary:
//
// 获取模型列表
//
// @param request - ListModelsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListModelsResponse
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

// Summary:
//
// 获取模型列表
//
// @param request - ListModelsRequest
//
// @return ListModelsResponse
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

// Summary:
//
// 列举权限
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPermissionsResponse
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

// Summary:
//
// 列举权限
//
// @return ListPermissionsResponse
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

// Summary:
//
// 列举产品
//
// @param request - ListProductsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListProductsResponse
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

// Summary:
//
// 列举产品
//
// @param request - ListProductsRequest
//
// @return ListProductsResponse
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

// Summary:
//
// 获取已有配额列表
//
// @param request - ListQuotasRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListQuotasResponse
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

// Summary:
//
// 获取已有配额列表
//
// @param request - ListQuotasRequest
//
// @return ListQuotasResponse
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

// Summary:
//
// 列举工作空间资源
//
// @param request - ListResourcesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListResourcesResponse
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

// Summary:
//
// 列举工作空间资源
//
// @param request - ListResourcesRequest
//
// @return ListResourcesResponse
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

// Summary:
//
// 获取Run的指标记录列表
//
// @param request - ListRunMetricsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRunMetricsResponse
func (client *Client) ListRunMetricsWithOptions(RunId *string, request *ListRunMetricsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListRunMetricsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Key)) {
		query["Key"] = request.Key
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.PageToken)) {
		query["PageToken"] = request.PageToken
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListRunMetrics"),
		Version:     tea.String("2021-02-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/runs/" + tea.StringValue(openapiutil.GetEncodeParam(RunId)) + "/metrics"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListRunMetricsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取Run的指标记录列表
//
// @param request - ListRunMetricsRequest
//
// @return ListRunMetricsResponse
func (client *Client) ListRunMetrics(RunId *string, request *ListRunMetricsRequest) (_result *ListRunMetricsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListRunMetricsResponse{}
	_body, _err := client.ListRunMetricsWithOptions(RunId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取Run列表
//
// @param request - ListRunsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRunsResponse
func (client *Client) ListRunsWithOptions(request *ListRunsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListRunsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ExperimentId)) {
		query["ExperimentId"] = request.ExperimentId
	}

	if !tea.BoolValue(util.IsUnset(request.GmtCreateTime)) {
		query["GmtCreateTime"] = request.GmtCreateTime
	}

	if !tea.BoolValue(util.IsUnset(request.Labels)) {
		query["Labels"] = request.Labels
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Order)) {
		query["Order"] = request.Order
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

	if !tea.BoolValue(util.IsUnset(request.PageToken)) {
		query["PageToken"] = request.PageToken
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
		Action:      tea.String("ListRuns"),
		Version:     tea.String("2021-02-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/runs"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListRunsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取Run列表
//
// @param request - ListRunsRequest
//
// @return ListRunsResponse
func (client *Client) ListRuns(request *ListRunsRequest) (_result *ListRunsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListRunsResponse{}
	_body, _err := client.ListRunsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 列出工作空间的可变为成员的用户
//
// @param request - ListWorkspaceUsersRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListWorkspaceUsersResponse
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

// Summary:
//
// 列出工作空间的可变为成员的用户
//
// @param request - ListWorkspaceUsersRequest
//
// @return ListWorkspaceUsersResponse
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

// Summary:
//
// 获得工作空间列表
//
// @param request - ListWorkspacesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListWorkspacesResponse
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

// Summary:
//
// 获得工作空间列表
//
// @param request - ListWorkspacesRequest
//
// @return ListWorkspacesResponse
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

// Summary:
//
// 批量记录Run的指标
//
// @param request - LogRunMetricsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return LogRunMetricsResponse
func (client *Client) LogRunMetricsWithOptions(RunId *string, request *LogRunMetricsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *LogRunMetricsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Metrics)) {
		body["Metrics"] = request.Metrics
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("LogRunMetrics"),
		Version:     tea.String("2021-02-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/runs/" + tea.StringValue(openapiutil.GetEncodeParam(RunId)) + "/metrics/action/log"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &LogRunMetricsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量记录Run的指标
//
// @param request - LogRunMetricsRequest
//
// @return LogRunMetricsResponse
func (client *Client) LogRunMetrics(RunId *string, request *LogRunMetricsRequest) (_result *LogRunMetricsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &LogRunMetricsResponse{}
	_body, _err := client.LogRunMetricsWithOptions(RunId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 发布一个代码源配置为本工作空间下所有人可见
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PublishCodeSourceResponse
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

// Summary:
//
// 发布一个代码源配置为本工作空间下所有人可见
//
// @return PublishCodeSourceResponse
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

// Summary:
//
// 更新数据集
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PublishDatasetResponse
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

// Summary:
//
// 更新数据集
//
// @return PublishDatasetResponse
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

// Summary:
//
// 发布 Image
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PublishImageResponse
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

// Summary:
//
// 发布 Image
//
// @return PublishImageResponse
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

// Summary:
//
// 删除 Image
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveImageResponse
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

// Summary:
//
// 删除 Image
//
// @return RemoveImageResponse
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

// Summary:
//
// 删除 Image 的标签
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveImageLabelsResponse
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

// Summary:
//
// 删除 Image 的标签
//
// @return RemoveImageLabelsResponse
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

// Summary:
//
// 删除成员角色
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveMemberRoleResponse
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

// Summary:
//
// 删除成员角色
//
// @return RemoveMemberRoleResponse
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

// Summary:
//
// 更新实验标签
//
// @param request - SetExperimentLabelsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetExperimentLabelsResponse
func (client *Client) SetExperimentLabelsWithOptions(ExperimentId *string, request *SetExperimentLabelsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *SetExperimentLabelsResponse, _err error) {
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
		Action:      tea.String("SetExperimentLabels"),
		Version:     tea.String("2021-02-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/experiments/" + tea.StringValue(openapiutil.GetEncodeParam(ExperimentId)) + "/labels"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &SetExperimentLabelsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新实验标签
//
// @param request - SetExperimentLabelsRequest
//
// @return SetExperimentLabelsResponse
func (client *Client) SetExperimentLabels(ExperimentId *string, request *SetExperimentLabelsRequest) (_result *SetExperimentLabelsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &SetExperimentLabelsResponse{}
	_body, _err := client.SetExperimentLabelsWithOptions(ExperimentId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新数据集
//
// @param request - UpdateDatasetRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDatasetResponse
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

// Summary:
//
// 更新数据集
//
// @param request - UpdateDatasetRequest
//
// @return UpdateDatasetResponse
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

// Summary:
//
// 更新指定版本的数据集信息
//
// @param request - UpdateDatasetVersionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDatasetVersionResponse
func (client *Client) UpdateDatasetVersionWithOptions(DatasetId *string, VersionName *string, request *UpdateDatasetVersionRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateDatasetVersionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DataCount)) {
		body["DataCount"] = request.DataCount
	}

	if !tea.BoolValue(util.IsUnset(request.DataSize)) {
		body["DataSize"] = request.DataSize
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Options)) {
		body["Options"] = request.Options
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateDatasetVersion"),
		Version:     tea.String("2021-02-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/datasets/" + tea.StringValue(openapiutil.GetEncodeParam(DatasetId)) + "/versions/" + tea.StringValue(openapiutil.GetEncodeParam(VersionName))),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateDatasetVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新指定版本的数据集信息
//
// @param request - UpdateDatasetVersionRequest
//
// @return UpdateDatasetVersionResponse
func (client *Client) UpdateDatasetVersion(DatasetId *string, VersionName *string, request *UpdateDatasetVersionRequest) (_result *UpdateDatasetVersionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateDatasetVersionResponse{}
	_body, _err := client.UpdateDatasetVersionWithOptions(DatasetId, VersionName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新默认工作空间
//
// @param request - UpdateDefaultWorkspaceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDefaultWorkspaceResponse
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

// Summary:
//
// 更新默认工作空间
//
// @param request - UpdateDefaultWorkspaceRequest
//
// @return UpdateDefaultWorkspaceResponse
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

// Summary:
//
// 更新实验
//
// @param request - UpdateExperimentRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateExperimentResponse
func (client *Client) UpdateExperimentWithOptions(ExperimentId *string, request *UpdateExperimentRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateExperimentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Accessibility)) {
		body["Accessibility"] = request.Accessibility
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateExperiment"),
		Version:     tea.String("2021-02-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/experiments/" + tea.StringValue(openapiutil.GetEncodeParam(ExperimentId))),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateExperimentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新实验
//
// @param request - UpdateExperimentRequest
//
// @return UpdateExperimentResponse
func (client *Client) UpdateExperiment(ExperimentId *string, request *UpdateExperimentRequest) (_result *UpdateExperimentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateExperimentResponse{}
	_body, _err := client.UpdateExperimentWithOptions(ExperimentId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新模型
//
// @param request - UpdateModelRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateModelResponse
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

// Summary:
//
// 更新模型
//
// @param request - UpdateModelRequest
//
// @return UpdateModelResponse
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

// Summary:
//
// 更新模型版本
//
// @param request - UpdateModelVersionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateModelVersionResponse
func (client *Client) UpdateModelVersionWithOptions(ModelId *string, VersionName *string, request *UpdateModelVersionRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateModelVersionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApprovalStatus)) {
		body["ApprovalStatus"] = request.ApprovalStatus
	}

	if !tea.BoolValue(util.IsUnset(request.CompressionSpec)) {
		body["CompressionSpec"] = request.CompressionSpec
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

// Summary:
//
// 更新模型版本
//
// @param request - UpdateModelVersionRequest
//
// @return UpdateModelVersionResponse
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

// Summary:
//
// 更新Run
//
// @param request - UpdateRunRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateRunResponse
func (client *Client) UpdateRunWithOptions(RunId *string, request *UpdateRunRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateRunResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Labels)) {
		body["Labels"] = request.Labels
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Params)) {
		body["Params"] = request.Params
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateRun"),
		Version:     tea.String("2021-02-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/runs/" + tea.StringValue(openapiutil.GetEncodeParam(RunId))),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateRunResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新Run
//
// @param request - UpdateRunRequest
//
// @return UpdateRunResponse
func (client *Client) UpdateRun(RunId *string, request *UpdateRunRequest) (_result *UpdateRunResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateRunResponse{}
	_body, _err := client.UpdateRunWithOptions(RunId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新工作空间
//
// @param request - UpdateWorkspaceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateWorkspaceResponse
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

// Summary:
//
// 更新工作空间
//
// @param request - UpdateWorkspaceRequest
//
// @return UpdateWorkspaceResponse
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

// Summary:
//
// 更新工作空间资源
//
// @param request - UpdateWorkspaceResourceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateWorkspaceResourceResponse
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

// Summary:
//
// 更新工作空间资源
//
// @param request - UpdateWorkspaceResourceRequest
//
// @return UpdateWorkspaceResourceResponse
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
