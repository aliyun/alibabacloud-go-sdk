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

type CodeSourceItem struct {
	// 表示代码是否是工作空间下公开的,可选值PRIVATE,PUBLIC
	Accessibility *string `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	// 代码分支
	CodeBranch *string `json:"CodeBranch,omitempty" xml:"CodeBranch,omitempty"`
	// 代码Commit ID
	CodeCommit *string `json:"CodeCommit,omitempty" xml:"CodeCommit,omitempty"`
	// 代码仓库地址
	CodeRepo *string `json:"CodeRepo,omitempty" xml:"CodeRepo,omitempty"`
	// 访问代码仓库所用的AccessToken
	CodeRepoAccessToken *string `json:"CodeRepoAccessToken,omitempty" xml:"CodeRepoAccessToken,omitempty"`
	// 访问代码仓库的用户名
	CodeRepoUserName *string `json:"CodeRepoUserName,omitempty" xml:"CodeRepoUserName,omitempty"`
	// 代码源ID
	CodeSourceId *string `json:"CodeSourceId,omitempty" xml:"CodeSourceId,omitempty"`
	// 代码源详细描述
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 代码源配置的名字
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// 创建时间
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// 修改时间
	GmtModifyTime *string `json:"GmtModifyTime,omitempty" xml:"GmtModifyTime,omitempty"`
	// 代码源配置的用户ID
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// 工作空间ID
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

func (s *CodeSourceItem) SetUserId(v string) *CodeSourceItem {
	s.UserId = &v
	return s
}

func (s *CodeSourceItem) SetWorkspaceId(v string) *CodeSourceItem {
	s.WorkspaceId = &v
	return s
}

type Dataset struct {
	// Accessibility
	Accessibility *string `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	// DataSourceType
	DataSourceType *string `json:"DataSourceType,omitempty" xml:"DataSourceType,omitempty"`
	// DataType
	DataType *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	// Dataset Id
	DatasetId *string `json:"DatasetId,omitempty" xml:"DatasetId,omitempty"`
	// Description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// GmtCreateTime
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// GmtModifiedTime
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// Labels
	Labels []*DatasetLabel `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	// Name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Options
	Options *string `json:"Options,omitempty" xml:"Options,omitempty"`
	// Owner Id
	OwnerId *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// Property
	Property *string `json:"Property,omitempty" xml:"Property,omitempty"`
	// SourceId
	SourceId *string `json:"SourceId,omitempty" xml:"SourceId,omitempty"`
	// SourceType
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// Uri
	Uri *string `json:"Uri,omitempty" xml:"Uri,omitempty"`
	// User Id
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// WorkspaceId
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

func (s *Dataset) SetLabels(v []*DatasetLabel) *Dataset {
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
	// Key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// Value
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

type AddImageRequest struct {
	// 可见性 Public 公有 Private 私有
	Accessibility *string `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	// 镜像描述
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 镜像地址
	ImageUri *string `json:"ImageUri,omitempty" xml:"ImageUri,omitempty"`
	// 镜像标签，是个数组
	Labels []*AddImageRequestLabels `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	// 镜像名称
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 工作空间id
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

func (s *AddImageRequest) SetWorkspaceId(v string) *AddImageRequest {
	s.WorkspaceId = &v
	return s
}

type AddImageRequestLabels struct {
	// Key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// Value
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
	// 镜像 id
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// Id of the request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
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
	Headers map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddImageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *AddImageResponse) SetBody(v *AddImageResponseBody) *AddImageResponse {
	s.Body = v
	return s
}

type AddImageLabelsRequest struct {
	// 标签
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
	// Key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// Value
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
	// Id of the request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
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
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddImageLabelsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *AddImageLabelsResponse) SetBody(v *AddImageLabelsResponseBody) *AddImageLabelsResponse {
	s.Body = v
	return s
}

type AddMemberRoleResponseBody struct {
	// 请求 id
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
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddMemberRoleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *AddMemberRoleResponse) SetBody(v *AddMemberRoleResponseBody) *AddMemberRoleResponse {
	s.Body = v
	return s
}

type AddWorkspaceQuotaRequest struct {
	// 模式  isolate 预付费  share 后付费  develop 开发模式
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// 产品代码
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// 产品类型，  支持PAI，MaxCompute，
	QuotaType *string `json:"QuotaType,omitempty" xml:"QuotaType,omitempty"`
}

func (s AddWorkspaceQuotaRequest) String() string {
	return tea.Prettify(s)
}

func (s AddWorkspaceQuotaRequest) GoString() string {
	return s.String()
}

func (s *AddWorkspaceQuotaRequest) SetMode(v string) *AddWorkspaceQuotaRequest {
	s.Mode = &v
	return s
}

func (s *AddWorkspaceQuotaRequest) SetProductCode(v string) *AddWorkspaceQuotaRequest {
	s.ProductCode = &v
	return s
}

func (s *AddWorkspaceQuotaRequest) SetQuotaType(v string) *AddWorkspaceQuotaRequest {
	s.QuotaType = &v
	return s
}

type AddWorkspaceQuotaResponseBody struct {
	// 请求 id
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddWorkspaceQuotaResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddWorkspaceQuotaResponseBody) GoString() string {
	return s.String()
}

func (s *AddWorkspaceQuotaResponseBody) SetRequestId(v string) *AddWorkspaceQuotaResponseBody {
	s.RequestId = &v
	return s
}

type AddWorkspaceQuotaResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddWorkspaceQuotaResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddWorkspaceQuotaResponse) String() string {
	return tea.Prettify(s)
}

func (s AddWorkspaceQuotaResponse) GoString() string {
	return s.String()
}

func (s *AddWorkspaceQuotaResponse) SetHeaders(v map[string]*string) *AddWorkspaceQuotaResponse {
	s.Headers = v
	return s
}

func (s *AddWorkspaceQuotaResponse) SetBody(v *AddWorkspaceQuotaResponseBody) *AddWorkspaceQuotaResponse {
	s.Body = v
	return s
}

type CreateCodeSourceRequest struct {
	// 代码是否是本工作空间内公开的，可选值PRIVATE、PUBLIC
	Accessibility *string `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	// 代码分支
	CodeBranch *string `json:"CodeBranch,omitempty" xml:"CodeBranch,omitempty"`
	// 代码仓库地址
	CodeRepo *string `json:"CodeRepo,omitempty" xml:"CodeRepo,omitempty"`
	// 代码仓库访问Token
	CodeRepoAccessToken *string `json:"CodeRepoAccessToken,omitempty" xml:"CodeRepoAccessToken,omitempty"`
	// 代码仓库的用户名
	CodeRepoUserName *string `json:"CodeRepoUserName,omitempty" xml:"CodeRepoUserName,omitempty"`
	// 代码源详细描述
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 代码源配置名称
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// 代码本地挂载目录，默认挂载到/root/code/下
	MountPath *string `json:"MountPath,omitempty" xml:"MountPath,omitempty"`
	// 工作空间ID
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
	// 创建的代码源配置的ID
	CodeSourceId *string `json:"CodeSourceId,omitempty" xml:"CodeSourceId,omitempty"`
	// 请求ID
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
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateCodeSourceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateCodeSourceResponse) SetBody(v *CreateCodeSourceResponseBody) *CreateCodeSourceResponse {
	s.Body = v
	return s
}

type CreateDatasetRequest struct {
	// Accessibility
	Accessibility *string `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	// DataSourceType
	DataSourceType *string `json:"DataSourceType,omitempty" xml:"DataSourceType,omitempty"`
	// DataType
	DataType *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	// Description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Labels
	Labels []*DatasetLabel `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	// Name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Options
	Options *string `json:"Options,omitempty" xml:"Options,omitempty"`
	// Property
	Property *string `json:"Property,omitempty" xml:"Property,omitempty"`
	// SourceId
	SourceId *string `json:"SourceId,omitempty" xml:"SourceId,omitempty"`
	// SourceType
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// Uri
	Uri *string `json:"Uri,omitempty" xml:"Uri,omitempty"`
	// WorkspaceId
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

func (s *CreateDatasetRequest) SetLabels(v []*DatasetLabel) *CreateDatasetRequest {
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
	// Id of the request
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
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateDatasetResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateDatasetResponse) SetBody(v *CreateDatasetResponseBody) *CreateDatasetResponse {
	s.Body = v
	return s
}

type CreateDatasetLabelsRequest struct {
	Labels []*DatasetLabel `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
}

func (s CreateDatasetLabelsRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDatasetLabelsRequest) GoString() string {
	return s.String()
}

func (s *CreateDatasetLabelsRequest) SetLabels(v []*DatasetLabel) *CreateDatasetLabelsRequest {
	s.Labels = v
	return s
}

type CreateDatasetLabelsResponseBody struct {
	// Id of the request
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
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateDatasetLabelsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateDatasetLabelsResponse) SetBody(v *CreateDatasetLabelsResponseBody) *CreateDatasetLabelsResponse {
	s.Body = v
	return s
}

type CreateDefaultWorkspaceRequest struct {
	// 描述，最多80个字符
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 环境列表
	EnvTypes []*string `json:"EnvTypes,omitempty" xml:"EnvTypes,omitempty" type:"Repeated"`
	// 资源
	Resources []*CreateDefaultWorkspaceRequestResources `json:"Resources,omitempty" xml:"Resources,omitempty" type:"Repeated"`
}

func (s CreateDefaultWorkspaceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDefaultWorkspaceRequest) GoString() string {
	return s.String()
}

func (s *CreateDefaultWorkspaceRequest) SetDescription(v string) *CreateDefaultWorkspaceRequest {
	s.Description = &v
	return s
}

func (s *CreateDefaultWorkspaceRequest) SetEnvTypes(v []*string) *CreateDefaultWorkspaceRequest {
	s.EnvTypes = v
	return s
}

func (s *CreateDefaultWorkspaceRequest) SetResources(v []*CreateDefaultWorkspaceRequestResources) *CreateDefaultWorkspaceRequest {
	s.Resources = v
	return s
}

type CreateDefaultWorkspaceRequestResources struct {
	// 产品类型
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
}

func (s CreateDefaultWorkspaceRequestResources) String() string {
	return tea.Prettify(s)
}

func (s CreateDefaultWorkspaceRequestResources) GoString() string {
	return s.String()
}

func (s *CreateDefaultWorkspaceRequestResources) SetProductType(v string) *CreateDefaultWorkspaceRequestResources {
	s.ProductType = &v
	return s
}

type CreateDefaultWorkspaceResponseBody struct {
	// 请求 id
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 工作空间 id
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s CreateDefaultWorkspaceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDefaultWorkspaceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDefaultWorkspaceResponseBody) SetRequestId(v string) *CreateDefaultWorkspaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDefaultWorkspaceResponseBody) SetWorkspaceId(v string) *CreateDefaultWorkspaceResponseBody {
	s.WorkspaceId = &v
	return s
}

type CreateDefaultWorkspaceResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateDefaultWorkspaceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateDefaultWorkspaceResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDefaultWorkspaceResponse) GoString() string {
	return s.String()
}

func (s *CreateDefaultWorkspaceResponse) SetHeaders(v map[string]*string) *CreateDefaultWorkspaceResponse {
	s.Headers = v
	return s
}

func (s *CreateDefaultWorkspaceResponse) SetBody(v *CreateDefaultWorkspaceResponseBody) *CreateDefaultWorkspaceResponse {
	s.Body = v
	return s
}

type CreateMemberRequest struct {
	// 用户列表
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
	// 角色列表
	Roles []*string `json:"Roles,omitempty" xml:"Roles,omitempty" type:"Repeated"`
	// 用户 id
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
	// 成员列表
	Members []*CreateMemberResponseBodyMembers `json:"Members,omitempty" xml:"Members,omitempty" type:"Repeated"`
	// 请求 id
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
	// 成员显示名
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// 成员 id
	MemberId *string `json:"MemberId,omitempty" xml:"MemberId,omitempty"`
	// 角色列表
	Roles []*string `json:"Roles,omitempty" xml:"Roles,omitempty" type:"Repeated"`
	// 用户 id
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
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateMemberResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateMemberResponse) SetBody(v *CreateMemberResponseBody) *CreateMemberResponse {
	s.Body = v
	return s
}

type CreateProductOrdersRequest struct {
	// 是否自动购买所有产品
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// 逗号分隔的产品
	Products *CreateProductOrdersRequestProducts `json:"Products,omitempty" xml:"Products,omitempty" type:"Struct"`
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

func (s *CreateProductOrdersRequest) SetProducts(v *CreateProductOrdersRequestProducts) *CreateProductOrdersRequest {
	s.Products = v
	return s
}

type CreateProductOrdersRequestProducts struct {
	// 是否自动续费
	AutoRenew *bool `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// 付费类型
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// 购买时长,与pricingCycle配合使用
	Duration *int64 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// 实例属性信息
	InstanceProperties []*CreateProductOrdersRequestProductsInstanceProperties `json:"InstanceProperties,omitempty" xml:"InstanceProperties,omitempty" type:"Repeated"`
	// 订单类型
	OrderType *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	// 计价单位
	PricingCycle *string `json:"PricingCycle,omitempty" xml:"PricingCycle,omitempty"`
	// 产品code
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
	// 代号
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 名
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 值
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
	// 产品购买请求id
	BuyProductRequestId *string `json:"BuyProductRequestId,omitempty" xml:"BuyProductRequestId,omitempty"`
	// 消息
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 订单id
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// Id of the request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
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
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateProductOrdersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateProductOrdersResponse) SetBody(v *CreateProductOrdersResponseBody) *CreateProductOrdersResponse {
	s.Body = v
	return s
}

type CreateUserResponseBody struct {
	// 请求 id
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 用户 id
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s CreateUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateUserResponseBody) GoString() string {
	return s.String()
}

func (s *CreateUserResponseBody) SetRequestId(v string) *CreateUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateUserResponseBody) SetUserId(v string) *CreateUserResponseBody {
	s.UserId = &v
	return s
}

type CreateUserResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateUserResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateUserResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateUserResponse) GoString() string {
	return s.String()
}

func (s *CreateUserResponse) SetHeaders(v map[string]*string) *CreateUserResponse {
	s.Headers = v
	return s
}

func (s *CreateUserResponse) SetBody(v *CreateUserResponseBody) *CreateUserResponse {
	s.Body = v
	return s
}

type CreateWorkspaceRequest struct {
	// 描述，最多80个字符
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 显示名称
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// 环境列表
	EnvTypes []*string `json:"EnvTypes,omitempty" xml:"EnvTypes,omitempty" type:"Repeated"`
	// 名字 3-23 个字符, 需要字母开头，只能包含字母下划线和数字，region内唯一
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
	// 请求 id
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 工作空间 id
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
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateWorkspaceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateWorkspaceResponse) SetBody(v *CreateWorkspaceResponseBody) *CreateWorkspaceResponse {
	s.Body = v
	return s
}

type CreateWorkspaceResourceRequest struct {
	// 资源列表
	Resources []*CreateWorkspaceResourceRequestResources `json:"Resources,omitempty" xml:"Resources,omitempty" type:"Repeated"`
}

func (s CreateWorkspaceResourceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateWorkspaceResourceRequest) GoString() string {
	return s.String()
}

func (s *CreateWorkspaceResourceRequest) SetResources(v []*CreateWorkspaceResourceRequestResources) *CreateWorkspaceResourceRequest {
	s.Resources = v
	return s
}

type CreateWorkspaceResourceRequestResources struct {
	// 环境， 支持dev（开发）、prod（生产）
	EnvType *string `json:"EnvType,omitempty" xml:"EnvType,omitempty"`
	// 分组名，主账户内唯一 一个 GroupName 下可能有一个 dev 资源和一个 prod 资源
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// 是否默认资源 每个类型都有一个默认的资源
	IsDefault *bool `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	// 资源名 长度需要在3到27个字符 region内唯一
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 产品类型， 支持PAI，MaxCompute
	ProductType *string                                          `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	Quotas      []*CreateWorkspaceResourceRequestResourcesQuotas `json:"Quotas,omitempty" xml:"Quotas,omitempty" type:"Repeated"`
	// 对于MaxCompute是个json，有如下key： Endpoint Project
	Spec *string `json:"Spec,omitempty" xml:"Spec,omitempty"`
	// 所属的工作空间 id
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

func (s *CreateWorkspaceResourceRequestResources) SetSpec(v string) *CreateWorkspaceResourceRequestResources {
	s.Spec = &v
	return s
}

func (s *CreateWorkspaceResourceRequestResources) SetWorkspaceId(v string) *CreateWorkspaceResourceRequestResources {
	s.WorkspaceId = &v
	return s
}

type CreateWorkspaceResourceRequestResourcesQuotas struct {
	// 卡类型，支持cpu、gpu
	CardType *string `json:"CardType,omitempty" xml:"CardType,omitempty"`
	// 模式 isolate 预付费 share 后付费 develop 开发模式
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// 配额名称
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 商品 code
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// 产品类型， 支持PAI，MaxCompute
	QuotaType *string `json:"QuotaType,omitempty" xml:"QuotaType,omitempty"`
	// 规格描述
	Spec *string `json:"Spec,omitempty" xml:"Spec,omitempty"`
}

func (s CreateWorkspaceResourceRequestResourcesQuotas) String() string {
	return tea.Prettify(s)
}

func (s CreateWorkspaceResourceRequestResourcesQuotas) GoString() string {
	return s.String()
}

func (s *CreateWorkspaceResourceRequestResourcesQuotas) SetCardType(v string) *CreateWorkspaceResourceRequestResourcesQuotas {
	s.CardType = &v
	return s
}

func (s *CreateWorkspaceResourceRequestResourcesQuotas) SetMode(v string) *CreateWorkspaceResourceRequestResourcesQuotas {
	s.Mode = &v
	return s
}

func (s *CreateWorkspaceResourceRequestResourcesQuotas) SetName(v string) *CreateWorkspaceResourceRequestResourcesQuotas {
	s.Name = &v
	return s
}

func (s *CreateWorkspaceResourceRequestResourcesQuotas) SetProductCode(v string) *CreateWorkspaceResourceRequestResourcesQuotas {
	s.ProductCode = &v
	return s
}

func (s *CreateWorkspaceResourceRequestResourcesQuotas) SetQuotaType(v string) *CreateWorkspaceResourceRequestResourcesQuotas {
	s.QuotaType = &v
	return s
}

func (s *CreateWorkspaceResourceRequestResourcesQuotas) SetSpec(v string) *CreateWorkspaceResourceRequestResourcesQuotas {
	s.Spec = &v
	return s
}

type CreateWorkspaceResourceResponseBody struct {
	// 请求 id
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 资源集合
	Resources []*CreateWorkspaceResourceResponseBodyResources `json:"Resources,omitempty" xml:"Resources,omitempty" type:"Repeated"`
	// 总数
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
	// 资源Id
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
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateWorkspaceResourceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateWorkspaceResourceResponse) SetBody(v *CreateWorkspaceResourceResponseBody) *CreateWorkspaceResourceResponse {
	s.Body = v
	return s
}

type DeleteCodeSourceResponseBody struct {
	// 被删除的代码源配置ID
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
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteCodeSourceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DeleteCodeSourceResponse) SetBody(v *DeleteCodeSourceResponseBody) *DeleteCodeSourceResponse {
	s.Body = v
	return s
}

type DeleteConfigResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteConfigResponseBody) SetRequestId(v string) *DeleteConfigResponseBody {
	s.RequestId = &v
	return s
}

type DeleteConfigResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteConfigResponse) GoString() string {
	return s.String()
}

func (s *DeleteConfigResponse) SetHeaders(v map[string]*string) *DeleteConfigResponse {
	s.Headers = v
	return s
}

func (s *DeleteConfigResponse) SetBody(v *DeleteConfigResponseBody) *DeleteConfigResponse {
	s.Body = v
	return s
}

type DeleteDatasetResponseBody struct {
	// Id of the request
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
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteDatasetResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DeleteDatasetResponse) SetBody(v *DeleteDatasetResponseBody) *DeleteDatasetResponse {
	s.Body = v
	return s
}

type DeleteDatasetLabelsRequest struct {
	Keys *string `json:"Keys,omitempty" xml:"Keys,omitempty"`
}

func (s DeleteDatasetLabelsRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDatasetLabelsRequest) GoString() string {
	return s.String()
}

func (s *DeleteDatasetLabelsRequest) SetKeys(v string) *DeleteDatasetLabelsRequest {
	s.Keys = &v
	return s
}

type DeleteDatasetLabelsResponseBody struct {
	// Id of the request
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
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteDatasetLabelsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DeleteDatasetLabelsResponse) SetBody(v *DeleteDatasetLabelsResponseBody) *DeleteDatasetLabelsResponse {
	s.Body = v
	return s
}

type DeleteMembersRequest struct {
	// 需要删除的成员 Id 列表，以逗号分隔
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
	// 请求 id
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteMembersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteMembersResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteMembersResponseBody) SetRequestId(v string) *DeleteMembersResponseBody {
	s.RequestId = &v
	return s
}

type DeleteMembersResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteMembersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DeleteMembersResponse) SetBody(v *DeleteMembersResponseBody) *DeleteMembersResponse {
	s.Body = v
	return s
}

type DeleteWorkspaceResponseBody struct {
	// Id of the request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
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
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteWorkspaceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DeleteWorkspaceResponse) SetBody(v *DeleteWorkspaceResponseBody) *DeleteWorkspaceResponse {
	s.Body = v
	return s
}

type GetCodeSourceResponseBody struct {
	// 代码是否是本工作空间内公开的，可选值PRIVATE、PUBLIC
	Accessibility *string `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	// 代码仓库分支
	CodeBranch *string `json:"CodeBranch,omitempty" xml:"CodeBranch,omitempty"`
	// 代码Commit
	CodeCommit *string `json:"CodeCommit,omitempty" xml:"CodeCommit,omitempty"`
	// 代码仓库地址
	CodeRepo *string `json:"CodeRepo,omitempty" xml:"CodeRepo,omitempty"`
	// 访问代码仓库的token
	CodeRepoAccessToken *string `json:"CodeRepoAccessToken,omitempty" xml:"CodeRepoAccessToken,omitempty"`
	// 代码仓库的用户名
	CodeRepoUserName *string `json:"CodeRepoUserName,omitempty" xml:"CodeRepoUserName,omitempty"`
	// 代码源配置ID
	CodeSourceId *string `json:"CodeSourceId,omitempty" xml:"CodeSourceId,omitempty"`
	// 详细描述
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 代码源配置名字
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// 创建时间
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// 修改时间
	GmtModifyTime *string `json:"GmtModifyTime,omitempty" xml:"GmtModifyTime,omitempty"`
	// 代码本地挂载目录，默认挂载到/root/code/下
	MountPath *string `json:"MountPath,omitempty" xml:"MountPath,omitempty"`
	// 请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 代码配置源的创建者ID
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// 工作空间ID
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
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetCodeSourceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetCodeSourceResponse) SetBody(v *GetCodeSourceResponseBody) *GetCodeSourceResponse {
	s.Body = v
	return s
}

type GetCodeSourcesStatisticsRequest struct {
	// 工作空间ID
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s GetCodeSourcesStatisticsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetCodeSourcesStatisticsRequest) GoString() string {
	return s.String()
}

func (s *GetCodeSourcesStatisticsRequest) SetWorkspaceId(v string) *GetCodeSourcesStatisticsRequest {
	s.WorkspaceId = &v
	return s
}

type GetCodeSourcesStatisticsResponseBody struct {
	// 此用户可以查看的代码的数目
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// Id of the request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetCodeSourcesStatisticsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetCodeSourcesStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *GetCodeSourcesStatisticsResponseBody) SetCount(v int64) *GetCodeSourcesStatisticsResponseBody {
	s.Count = &v
	return s
}

func (s *GetCodeSourcesStatisticsResponseBody) SetRequestId(v string) *GetCodeSourcesStatisticsResponseBody {
	s.RequestId = &v
	return s
}

type GetCodeSourcesStatisticsResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetCodeSourcesStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetCodeSourcesStatisticsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetCodeSourcesStatisticsResponse) GoString() string {
	return s.String()
}

func (s *GetCodeSourcesStatisticsResponse) SetHeaders(v map[string]*string) *GetCodeSourcesStatisticsResponse {
	s.Headers = v
	return s
}

func (s *GetCodeSourcesStatisticsResponse) SetBody(v *GetCodeSourcesStatisticsResponseBody) *GetCodeSourcesStatisticsResponse {
	s.Body = v
	return s
}

type GetDatasetResponseBody struct {
	Accessibility   *string         `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	DataSourceType  *string         `json:"DataSourceType,omitempty" xml:"DataSourceType,omitempty"`
	DataType        *string         `json:"DataType,omitempty" xml:"DataType,omitempty"`
	DatasetId       *string         `json:"DatasetId,omitempty" xml:"DatasetId,omitempty"`
	Description     *string         `json:"Description,omitempty" xml:"Description,omitempty"`
	GmtCreateTime   *string         `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	GmtModifiedTime *string         `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	Labels          []*DatasetLabel `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	Name            *string         `json:"Name,omitempty" xml:"Name,omitempty"`
	Options         *string         `json:"Options,omitempty" xml:"Options,omitempty"`
	OwnerId         *string         `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	Property        *string         `json:"Property,omitempty" xml:"Property,omitempty"`
	// Id of the request
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SourceId    *string `json:"SourceId,omitempty" xml:"SourceId,omitempty"`
	SourceType  *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	Uri         *string `json:"Uri,omitempty" xml:"Uri,omitempty"`
	UserId      *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
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

func (s *GetDatasetResponseBody) SetLabels(v []*DatasetLabel) *GetDatasetResponseBody {
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
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetDatasetResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetDatasetResponse) SetBody(v *GetDatasetResponseBody) *GetDatasetResponse {
	s.Body = v
	return s
}

type GetDatasetsStatisticsRequest struct {
	// WorkspaceId
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s GetDatasetsStatisticsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDatasetsStatisticsRequest) GoString() string {
	return s.String()
}

func (s *GetDatasetsStatisticsRequest) SetWorkspaceId(v string) *GetDatasetsStatisticsRequest {
	s.WorkspaceId = &v
	return s
}

type GetDatasetsStatisticsResponseBody struct {
	// Count
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetDatasetsStatisticsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDatasetsStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *GetDatasetsStatisticsResponseBody) SetCount(v int64) *GetDatasetsStatisticsResponseBody {
	s.Count = &v
	return s
}

func (s *GetDatasetsStatisticsResponseBody) SetRequestId(v string) *GetDatasetsStatisticsResponseBody {
	s.RequestId = &v
	return s
}

type GetDatasetsStatisticsResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetDatasetsStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetDatasetsStatisticsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDatasetsStatisticsResponse) GoString() string {
	return s.String()
}

func (s *GetDatasetsStatisticsResponse) SetHeaders(v map[string]*string) *GetDatasetsStatisticsResponse {
	s.Headers = v
	return s
}

func (s *GetDatasetsStatisticsResponse) SetBody(v *GetDatasetsStatisticsResponseBody) *GetDatasetsStatisticsResponse {
	s.Body = v
	return s
}

type GetDefaultWorkspaceResponseBody struct {
	// 任务详情 创建默认工作空间会有多个任务依次进行，如果一个任务未开始，不会显示在任务详情里。
	Conditions []*GetDefaultWorkspaceResponseBodyConditions `json:"Conditions,omitempty" xml:"Conditions,omitempty" type:"Repeated"`
	// 创建人
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// 描述
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 显示名称
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// 环境，用作判断简单模式还是标准模式，含义见
	EnvTypes []*string `json:"EnvTypes,omitempty" xml:"EnvTypes,omitempty" type:"Repeated"`
	// 创建 UTC 时间，日期格式 iso8601
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// 修改 UTC 时间，日期格式 iso8601
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// 拥有者
	Owner *GetDefaultWorkspaceResponseBodyOwner `json:"Owner,omitempty" xml:"Owner,omitempty" type:"Struct"`
	// 工作空间状态
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 工作空间 ID
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
	// 项目空间名称， region 内唯一
	WorkspaceName *string `json:"WorkspaceName,omitempty" xml:"WorkspaceName,omitempty"`
	// Id of the request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
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

func (s *GetDefaultWorkspaceResponseBody) SetRequestId(v string) *GetDefaultWorkspaceResponseBody {
	s.RequestId = &v
	return s
}

type GetDefaultWorkspaceResponseBodyConditions struct {
	// 返回码，正常是200，其他都是错误
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// 消息
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 类型
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
	// 用户id
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// 用户kp
	UserKp *string `json:"UserKp,omitempty" xml:"UserKp,omitempty"`
	// 用户名
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
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetDefaultWorkspaceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetDefaultWorkspaceResponse) SetBody(v *GetDefaultWorkspaceResponseBody) *GetDefaultWorkspaceResponse {
	s.Body = v
	return s
}

type GetImageRequest struct {
	// 是否显示非必要信息：Labels
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
	// 可见性
	Accessibility *string `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	// 描述
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 创建 UTC 时间，日期格式 iso8601
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// 创建 UTC 时间，日期格式 iso8601
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// 镜像地址，包含版本号
	ImageUri *string `json:"ImageUri,omitempty" xml:"ImageUri,omitempty"`
	// 镜像标签
	Labels []*GetImageResponseBodyLabels `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	// 镜像名称
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 创建人
	OperatorCreate *string `json:"OperatorCreate,omitempty" xml:"OperatorCreate,omitempty"`
	// 创建人父账户
	ParentOperatorCreate *string `json:"ParentOperatorCreate,omitempty" xml:"ParentOperatorCreate,omitempty"`
	// 工作空间id
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
	// Id of the request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
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

func (s *GetImageResponseBody) SetOperatorCreate(v string) *GetImageResponseBody {
	s.OperatorCreate = &v
	return s
}

func (s *GetImageResponseBody) SetParentOperatorCreate(v string) *GetImageResponseBody {
	s.ParentOperatorCreate = &v
	return s
}

func (s *GetImageResponseBody) SetWorkspaceId(v string) *GetImageResponseBody {
	s.WorkspaceId = &v
	return s
}

func (s *GetImageResponseBody) SetRequestId(v string) *GetImageResponseBody {
	s.RequestId = &v
	return s
}

type GetImageResponseBodyLabels struct {
	// Key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// Value
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
	Headers map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetImageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetImageResponse) SetBody(v *GetImageResponseBody) *GetImageResponse {
	s.Body = v
	return s
}

type GetImagesStatisticsRequest struct {
	// 工作空间id
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s GetImagesStatisticsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetImagesStatisticsRequest) GoString() string {
	return s.String()
}

func (s *GetImagesStatisticsRequest) SetWorkspaceId(v string) *GetImagesStatisticsRequest {
	s.WorkspaceId = &v
	return s
}

type GetImagesStatisticsResponseBody struct {
	// 镜像总数
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// Id of the request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetImagesStatisticsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetImagesStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *GetImagesStatisticsResponseBody) SetCount(v int64) *GetImagesStatisticsResponseBody {
	s.Count = &v
	return s
}

func (s *GetImagesStatisticsResponseBody) SetRequestId(v string) *GetImagesStatisticsResponseBody {
	s.RequestId = &v
	return s
}

type GetImagesStatisticsResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetImagesStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetImagesStatisticsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetImagesStatisticsResponse) GoString() string {
	return s.String()
}

func (s *GetImagesStatisticsResponse) SetHeaders(v map[string]*string) *GetImagesStatisticsResponse {
	s.Headers = v
	return s
}

func (s *GetImagesStatisticsResponse) SetBody(v *GetImagesStatisticsResponseBody) *GetImagesStatisticsResponse {
	s.Body = v
	return s
}

type GetMemberRequest struct {
	// 用户id
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s GetMemberRequest) String() string {
	return tea.Prettify(s)
}

func (s GetMemberRequest) GoString() string {
	return s.String()
}

func (s *GetMemberRequest) SetUserId(v string) *GetMemberRequest {
	s.UserId = &v
	return s
}

type GetMemberResponseBody struct {
	// 成员显示名
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// 创建 UTC 时间，日期格式 iso8601
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// 成员 id
	MemberId *string `json:"MemberId,omitempty" xml:"MemberId,omitempty"`
	// 角色列表
	Roles []*string `json:"Roles,omitempty" xml:"Roles,omitempty" type:"Repeated"`
	// 用户 id
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// 云账号用户名
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	// Id of the request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
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

func (s *GetMemberResponseBody) SetRoles(v []*string) *GetMemberResponseBody {
	s.Roles = v
	return s
}

func (s *GetMemberResponseBody) SetUserId(v string) *GetMemberResponseBody {
	s.UserId = &v
	return s
}

func (s *GetMemberResponseBody) SetUserName(v string) *GetMemberResponseBody {
	s.UserName = &v
	return s
}

func (s *GetMemberResponseBody) SetRequestId(v string) *GetMemberResponseBody {
	s.RequestId = &v
	return s
}

type GetMemberResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetMemberResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetMemberResponse) SetBody(v *GetMemberResponseBody) *GetMemberResponse {
	s.Body = v
	return s
}

type GetPermissionRequest struct {
	// 待鉴权实例的可见性 PUBLIC表示当前工作空间所有人都可以访问 PRIVATE表示只有Owner可以访问
	Accessibility *string `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	// 创建者
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
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

type GetPermissionResponseBody struct {
	// 权限
	PermissionCode  *string                                     `json:"PermissionCode,omitempty" xml:"PermissionCode,omitempty"`
	PermissionRules []*GetPermissionResponseBodyPermissionRules `json:"PermissionRules,omitempty" xml:"PermissionRules,omitempty" type:"Repeated"`
	// 请求 id
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
	// 待鉴权实例的可见性 PUBLIC表示当前工作空间所有人都可以访问 PRIVATE表示只有Owner可以访问
	Accessibility *string `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	// 存取类型 CREATOR 代表能看自已创建的 OTHERS 代表可以看别人创建的 ALL代表所有
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
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetPermissionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetPermissionResponse) SetBody(v *GetPermissionResponseBody) *GetPermissionResponse {
	s.Body = v
	return s
}

type GetRoleStatisticsRequest struct {
	// 工作空间 id
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s GetRoleStatisticsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetRoleStatisticsRequest) GoString() string {
	return s.String()
}

func (s *GetRoleStatisticsRequest) SetWorkspaceId(v string) *GetRoleStatisticsRequest {
	s.WorkspaceId = &v
	return s
}

type GetRoleStatisticsResponseBody struct {
	// 请求 id
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 角色信息统计
	Roles []*GetRoleStatisticsResponseBodyRoles `json:"Roles,omitempty" xml:"Roles,omitempty" type:"Repeated"`
	// 总数
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetRoleStatisticsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetRoleStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *GetRoleStatisticsResponseBody) SetRequestId(v string) *GetRoleStatisticsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRoleStatisticsResponseBody) SetRoles(v []*GetRoleStatisticsResponseBodyRoles) *GetRoleStatisticsResponseBody {
	s.Roles = v
	return s
}

func (s *GetRoleStatisticsResponseBody) SetTotalCount(v int64) *GetRoleStatisticsResponseBody {
	s.TotalCount = &v
	return s
}

type GetRoleStatisticsResponseBodyRoles struct {
	// 成员数量
	MemberSize *int64 `json:"MemberSize,omitempty" xml:"MemberSize,omitempty"`
	// 角色名
	RoleName *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
}

func (s GetRoleStatisticsResponseBodyRoles) String() string {
	return tea.Prettify(s)
}

func (s GetRoleStatisticsResponseBodyRoles) GoString() string {
	return s.String()
}

func (s *GetRoleStatisticsResponseBodyRoles) SetMemberSize(v int64) *GetRoleStatisticsResponseBodyRoles {
	s.MemberSize = &v
	return s
}

func (s *GetRoleStatisticsResponseBodyRoles) SetRoleName(v string) *GetRoleStatisticsResponseBodyRoles {
	s.RoleName = &v
	return s
}

type GetRoleStatisticsResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetRoleStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetRoleStatisticsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetRoleStatisticsResponse) GoString() string {
	return s.String()
}

func (s *GetRoleStatisticsResponse) SetHeaders(v map[string]*string) *GetRoleStatisticsResponse {
	s.Headers = v
	return s
}

func (s *GetRoleStatisticsResponse) SetBody(v *GetRoleStatisticsResponseBody) *GetRoleStatisticsResponse {
	s.Body = v
	return s
}

type GetWorkspaceResponseBody struct {
	// 管理员账户
	AdminNames []*string `json:"AdminNames,omitempty" xml:"AdminNames,omitempty" type:"Repeated"`
	// 创建人
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// 描述
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 显示名称
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// 环境，用作判断简单模式还是标准模式
	EnvTypes []*string `json:"EnvTypes,omitempty" xml:"EnvTypes,omitempty" type:"Repeated"`
	// 附加信息
	ExtraInfos map[string]interface{} `json:"ExtraInfos,omitempty" xml:"ExtraInfos,omitempty"`
	// 创建 UTC 时间，日期格式 iso8601
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// 修改 UTC 时间，日期格式 iso8601
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// 是否为默认工作空间
	IsDefault *bool `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	// 请求 id
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 资源数目
	ResourceCount *int32 `json:"ResourceCount,omitempty" xml:"ResourceCount,omitempty"`
	// 工作空间状态
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 工作空间 id
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
	// 项目空间名称， region 内唯一
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

func (s *GetWorkspaceResponseBody) SetRequestId(v string) *GetWorkspaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetWorkspaceResponseBody) SetResourceCount(v int32) *GetWorkspaceResponseBody {
	s.ResourceCount = &v
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

type GetWorkspaceResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetWorkspaceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetWorkspaceResponse) SetBody(v *GetWorkspaceResponseBody) *GetWorkspaceResponse {
	s.Body = v
	return s
}

type ListCodeSourcesRequest struct {
	// 代码源显示名称，支持模糊匹配
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// 排序顺序, 枚举值 desc 或者 asc
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// 取第几页的数据
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// 分页大小
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 用于排序的字段名，可选字段名："DisplayName" "GmtCreateTime" "GmtModifyTime"
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// 工作空间ID
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
	// 代码源配置列表
	CodeSources []*CodeSourceItem `json:"CodeSources,omitempty" xml:"CodeSources,omitempty" type:"Repeated"`
	// 请求Id
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 符合过滤条件的代码源配置的总数量
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
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListCodeSourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListCodeSourcesResponse) SetBody(v *ListCodeSourcesResponseBody) *ListCodeSourcesResponse {
	s.Body = v
	return s
}

type ListConfigsRequest struct {
	// 配置键
	ConfigKeys *string `json:"ConfigKeys,omitempty" xml:"ConfigKeys,omitempty"`
}

func (s ListConfigsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListConfigsRequest) GoString() string {
	return s.String()
}

func (s *ListConfigsRequest) SetConfigKeys(v string) *ListConfigsRequest {
	s.ConfigKeys = &v
	return s
}

type ListConfigsResponseBody struct {
	// 配置列表
	Configs []*ListConfigsResponseBodyConfigs `json:"Configs,omitempty" xml:"Configs,omitempty" type:"Repeated"`
	// 请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 返回数量
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListConfigsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListConfigsResponseBody) GoString() string {
	return s.String()
}

func (s *ListConfigsResponseBody) SetConfigs(v []*ListConfigsResponseBodyConfigs) *ListConfigsResponseBody {
	s.Configs = v
	return s
}

func (s *ListConfigsResponseBody) SetRequestId(v string) *ListConfigsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListConfigsResponseBody) SetTotalCount(v int64) *ListConfigsResponseBody {
	s.TotalCount = &v
	return s
}

type ListConfigsResponseBodyConfigs struct {
	// 配置键
	ConfigKey *string `json:"ConfigKey,omitempty" xml:"ConfigKey,omitempty"`
	// 配置值
	ConfigValue *string `json:"ConfigValue,omitempty" xml:"ConfigValue,omitempty"`
}

func (s ListConfigsResponseBodyConfigs) String() string {
	return tea.Prettify(s)
}

func (s ListConfigsResponseBodyConfigs) GoString() string {
	return s.String()
}

func (s *ListConfigsResponseBodyConfigs) SetConfigKey(v string) *ListConfigsResponseBodyConfigs {
	s.ConfigKey = &v
	return s
}

func (s *ListConfigsResponseBodyConfigs) SetConfigValue(v string) *ListConfigsResponseBodyConfigs {
	s.ConfigValue = &v
	return s
}

type ListConfigsResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListConfigsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListConfigsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListConfigsResponse) GoString() string {
	return s.String()
}

func (s *ListConfigsResponse) SetHeaders(v map[string]*string) *ListConfigsResponse {
	s.Headers = v
	return s
}

func (s *ListConfigsResponse) SetBody(v *ListConfigsResponseBody) *ListConfigsResponse {
	s.Body = v
	return s
}

type ListDatasetsRequest struct {
	DataSourceTypes *string `json:"DataSourceTypes,omitempty" xml:"DataSourceTypes,omitempty"`
	DataTypes       *string `json:"DataTypes,omitempty" xml:"DataTypes,omitempty"`
	LabelKeys       *string `json:"LabelKeys,omitempty" xml:"LabelKeys,omitempty"`
	LabelValues     *string `json:"LabelValues,omitempty" xml:"LabelValues,omitempty"`
	Name            *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Order           *string `json:"Order,omitempty" xml:"Order,omitempty"`
	PageNumber      *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize        *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Properties      *string `json:"Properties,omitempty" xml:"Properties,omitempty"`
	SourceTypes     *string `json:"SourceTypes,omitempty" xml:"SourceTypes,omitempty"`
	// Workspace Id
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

func (s *ListDatasetsRequest) SetLabelKeys(v string) *ListDatasetsRequest {
	s.LabelKeys = &v
	return s
}

func (s *ListDatasetsRequest) SetLabelValues(v string) *ListDatasetsRequest {
	s.LabelValues = &v
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
	// Id of the request
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int64  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListDatasetsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListDatasetsResponse) SetBody(v *ListDatasetsResponseBody) *ListDatasetsResponse {
	s.Body = v
	return s
}

type ListFeaturesRequest struct {
	// 特性名称，以逗号分隔
	Names *string `json:"Names,omitempty" xml:"Names,omitempty"`
}

func (s ListFeaturesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListFeaturesRequest) GoString() string {
	return s.String()
}

func (s *ListFeaturesRequest) SetNames(v string) *ListFeaturesRequest {
	s.Names = &v
	return s
}

type ListFeaturesResponseBody struct {
	// 特性列表
	Features []*string `json:"Features,omitempty" xml:"Features,omitempty" type:"Repeated"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 特性总数
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListFeaturesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListFeaturesResponseBody) GoString() string {
	return s.String()
}

func (s *ListFeaturesResponseBody) SetFeatures(v []*string) *ListFeaturesResponseBody {
	s.Features = v
	return s
}

func (s *ListFeaturesResponseBody) SetRequestId(v string) *ListFeaturesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListFeaturesResponseBody) SetTotalCount(v int64) *ListFeaturesResponseBody {
	s.TotalCount = &v
	return s
}

type ListFeaturesResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListFeaturesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListFeaturesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListFeaturesResponse) GoString() string {
	return s.String()
}

func (s *ListFeaturesResponse) SetHeaders(v map[string]*string) *ListFeaturesResponse {
	s.Headers = v
	return s
}

func (s *ListFeaturesResponse) SetBody(v *ListFeaturesResponseBody) *ListFeaturesResponse {
	s.Body = v
	return s
}

type ListGlobalPermissionsResponseBody struct {
	// Permissions
	Permissions []*ListGlobalPermissionsResponseBodyPermissions `json:"Permissions,omitempty" xml:"Permissions,omitempty" type:"Repeated"`
	// Id of the request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListGlobalPermissionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListGlobalPermissionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListGlobalPermissionsResponseBody) SetPermissions(v []*ListGlobalPermissionsResponseBodyPermissions) *ListGlobalPermissionsResponseBody {
	s.Permissions = v
	return s
}

func (s *ListGlobalPermissionsResponseBody) SetRequestId(v string) *ListGlobalPermissionsResponseBody {
	s.RequestId = &v
	return s
}

type ListGlobalPermissionsResponseBodyPermissions struct {
	// PermissionCode
	PermissionCode *string `json:"PermissionCode,omitempty" xml:"PermissionCode,omitempty"`
	// PermissionRules
	PermissionRules []*ListGlobalPermissionsResponseBodyPermissionsPermissionRules `json:"PermissionRules,omitempty" xml:"PermissionRules,omitempty" type:"Repeated"`
}

func (s ListGlobalPermissionsResponseBodyPermissions) String() string {
	return tea.Prettify(s)
}

func (s ListGlobalPermissionsResponseBodyPermissions) GoString() string {
	return s.String()
}

func (s *ListGlobalPermissionsResponseBodyPermissions) SetPermissionCode(v string) *ListGlobalPermissionsResponseBodyPermissions {
	s.PermissionCode = &v
	return s
}

func (s *ListGlobalPermissionsResponseBodyPermissions) SetPermissionRules(v []*ListGlobalPermissionsResponseBodyPermissionsPermissionRules) *ListGlobalPermissionsResponseBodyPermissions {
	s.PermissionRules = v
	return s
}

type ListGlobalPermissionsResponseBodyPermissionsPermissionRules struct {
	// 待鉴权实例的可见性 PUBLIC表示当前工作空间所有人都可以访问 PRIVATE表示只有Owner可以访问
	Accessibility *string `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	// 存取类型 CREATOR 代表能看自已创建的 OTHERS 代表可以看别人创建的 ALL代表所有
	EntityAccessType *string `json:"EntityAccessType,omitempty" xml:"EntityAccessType,omitempty"`
}

func (s ListGlobalPermissionsResponseBodyPermissionsPermissionRules) String() string {
	return tea.Prettify(s)
}

func (s ListGlobalPermissionsResponseBodyPermissionsPermissionRules) GoString() string {
	return s.String()
}

func (s *ListGlobalPermissionsResponseBodyPermissionsPermissionRules) SetAccessibility(v string) *ListGlobalPermissionsResponseBodyPermissionsPermissionRules {
	s.Accessibility = &v
	return s
}

func (s *ListGlobalPermissionsResponseBodyPermissionsPermissionRules) SetEntityAccessType(v string) *ListGlobalPermissionsResponseBodyPermissionsPermissionRules {
	s.EntityAccessType = &v
	return s
}

type ListGlobalPermissionsResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListGlobalPermissionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListGlobalPermissionsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListGlobalPermissionsResponse) GoString() string {
	return s.String()
}

func (s *ListGlobalPermissionsResponse) SetHeaders(v map[string]*string) *ListGlobalPermissionsResponse {
	s.Headers = v
	return s
}

func (s *ListGlobalPermissionsResponse) SetBody(v *ListGlobalPermissionsResponseBody) *ListGlobalPermissionsResponse {
	s.Body = v
	return s
}

type ListImageLabelsRequest struct {
	// 镜像id
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// image过滤条件，获取满足条件的image的所有label
	LabelFilter *string `json:"LabelFilter,omitempty" xml:"LabelFilter,omitempty"`
	// 标签列表，以逗号分隔
	LabelKeys *string `json:"LabelKeys,omitempty" xml:"LabelKeys,omitempty"`
	// 地域
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// 工作空间id
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
	// 镜像标签
	Labels []*ListImageLabelsResponseBodyLabels `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	// 符合过滤条件的数量
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// Id of the request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
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

func (s *ListImageLabelsResponseBody) SetTotalCount(v int64) *ListImageLabelsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListImageLabelsResponseBody) SetRequestId(v string) *ListImageLabelsResponseBody {
	s.RequestId = &v
	return s
}

type ListImageLabelsResponseBodyLabels struct {
	// 键
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// 值
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
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListImageLabelsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListImageLabelsResponse) SetBody(v *ListImageLabelsResponseBody) *ListImageLabelsResponse {
	s.Body = v
	return s
}

type ListImagesRequest struct {
	// 过滤值 以逗号分隔
	Labels *string `json:"Labels,omitempty" xml:"Labels,omitempty"`
	// 镜像名称，支持模糊搜索
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 创建者
	OperatorCreate *string `json:"OperatorCreate,omitempty" xml:"OperatorCreate,omitempty"`
	// 排序方向： ASC - 升序 DESC - 降序
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// 分页，从1开始，默认1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// 页大小，默认20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 排序字段
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// 是否显示非必要信息：Labels
	Verbose *bool `json:"Verbose,omitempty" xml:"Verbose,omitempty"`
	// 工作空间id
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListImagesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListImagesRequest) GoString() string {
	return s.String()
}

func (s *ListImagesRequest) SetLabels(v string) *ListImagesRequest {
	s.Labels = &v
	return s
}

func (s *ListImagesRequest) SetName(v string) *ListImagesRequest {
	s.Name = &v
	return s
}

func (s *ListImagesRequest) SetOperatorCreate(v string) *ListImagesRequest {
	s.OperatorCreate = &v
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

func (s *ListImagesRequest) SetSortBy(v string) *ListImagesRequest {
	s.SortBy = &v
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
	// 镜像列表
	Images []*ListImagesResponseBodyImages `json:"Images,omitempty" xml:"Images,omitempty" type:"Repeated"`
	// 总数
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// Id of the request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
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

func (s *ListImagesResponseBody) SetTotalCount(v int64) *ListImagesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListImagesResponseBody) SetRequestId(v string) *ListImagesResponseBody {
	s.RequestId = &v
	return s
}

type ListImagesResponseBodyImages struct {
	// 可见性
	Accessibility *string `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	// 镜像描述
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 创建 UTC 时间，日期格式 iso8601
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// 修改时间
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// 镜像id
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// 镜像地址，包含版本号
	ImageUri *string `json:"ImageUri,omitempty" xml:"ImageUri,omitempty"`
	// 镜像标签，是个map
	Labels []*ListImagesResponseBodyImagesLabels `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	// 镜像名称
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 创建人
	OperatorCreate *string `json:"OperatorCreate,omitempty" xml:"OperatorCreate,omitempty"`
	// 父创建人
	ParentOperatorCreate *string `json:"ParentOperatorCreate,omitempty" xml:"ParentOperatorCreate,omitempty"`
	// 工作空间id
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

func (s *ListImagesResponseBodyImages) SetOperatorCreate(v string) *ListImagesResponseBodyImages {
	s.OperatorCreate = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetParentOperatorCreate(v string) *ListImagesResponseBodyImages {
	s.ParentOperatorCreate = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetWorkspaceId(v string) *ListImagesResponseBodyImages {
	s.WorkspaceId = &v
	return s
}

type ListImagesResponseBodyImagesLabels struct {
	// Key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// Value
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
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListImagesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListImagesResponse) SetBody(v *ListImagesResponseBody) *ListImagesResponse {
	s.Body = v
	return s
}

type ListMembersRequest struct {
	// 成员名
	MemberName *string `json:"MemberName,omitempty" xml:"MemberName,omitempty"`
	// 分页，从1开始，默认1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// 页大小，默认20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Role 过滤列表，逗号分隔
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
	// 成员列表
	Members []*ListMembersResponseBodyMembers `json:"Members,omitempty" xml:"Members,omitempty" type:"Repeated"`
	// 请求 id
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 符合过滤条件的数量
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
	// 成员显示名
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// 创建 UTC 时间，日期格式 iso8601
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// 成员 id
	MemberId *string `json:"MemberId,omitempty" xml:"MemberId,omitempty"`
	// 角色列表
	Roles []*string `json:"Roles,omitempty" xml:"Roles,omitempty" type:"Repeated"`
	// 用户 id
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// 云账号用户名
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
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

func (s *ListMembersResponseBodyMembers) SetRoles(v []*string) *ListMembersResponseBodyMembers {
	s.Roles = v
	return s
}

func (s *ListMembersResponseBodyMembers) SetUserId(v string) *ListMembersResponseBodyMembers {
	s.UserId = &v
	return s
}

func (s *ListMembersResponseBodyMembers) SetUserName(v string) *ListMembersResponseBodyMembers {
	s.UserName = &v
	return s
}

type ListMembersResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListMembersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListMembersResponse) SetBody(v *ListMembersResponseBody) *ListMembersResponse {
	s.Body = v
	return s
}

type ListOperationLogsRequest struct {
	EntityStatus *string `json:"EntityStatus,omitempty" xml:"EntityStatus,omitempty"`
	// 以逗号分隔的日志类型，包含 Resource
	EntityTypes *string `json:"EntityTypes,omitempty" xml:"EntityTypes,omitempty"`
	// 以逗号分隔的操作状态
	OperationStatus *string `json:"OperationStatus,omitempty" xml:"OperationStatus,omitempty"`
	// 以逗号分隔的操作
	Operations *string `json:"Operations,omitempty" xml:"Operations,omitempty"`
	// 排序顺序， 顺序：ASC，倒序：DESC
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// 当前页，页码从1开始
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// 每页返回的输出数目
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 排序字段
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
}

func (s ListOperationLogsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListOperationLogsRequest) GoString() string {
	return s.String()
}

func (s *ListOperationLogsRequest) SetEntityStatus(v string) *ListOperationLogsRequest {
	s.EntityStatus = &v
	return s
}

func (s *ListOperationLogsRequest) SetEntityTypes(v string) *ListOperationLogsRequest {
	s.EntityTypes = &v
	return s
}

func (s *ListOperationLogsRequest) SetOperationStatus(v string) *ListOperationLogsRequest {
	s.OperationStatus = &v
	return s
}

func (s *ListOperationLogsRequest) SetOperations(v string) *ListOperationLogsRequest {
	s.Operations = &v
	return s
}

func (s *ListOperationLogsRequest) SetOrder(v string) *ListOperationLogsRequest {
	s.Order = &v
	return s
}

func (s *ListOperationLogsRequest) SetPageNumber(v int64) *ListOperationLogsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListOperationLogsRequest) SetPageSize(v int32) *ListOperationLogsRequest {
	s.PageSize = &v
	return s
}

func (s *ListOperationLogsRequest) SetSortBy(v string) *ListOperationLogsRequest {
	s.SortBy = &v
	return s
}

type ListOperationLogsResponseBody struct {
	// 输出日志列表
	Logs []*ListOperationLogsResponseBodyLogs `json:"Logs,omitempty" xml:"Logs,omitempty" type:"Repeated"`
	// 请求 id
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 符合过滤条件的日志数量
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListOperationLogsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListOperationLogsResponseBody) GoString() string {
	return s.String()
}

func (s *ListOperationLogsResponseBody) SetLogs(v []*ListOperationLogsResponseBodyLogs) *ListOperationLogsResponseBody {
	s.Logs = v
	return s
}

func (s *ListOperationLogsResponseBody) SetRequestId(v string) *ListOperationLogsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListOperationLogsResponseBody) SetTotalCount(v int64) *ListOperationLogsResponseBody {
	s.TotalCount = &v
	return s
}

type ListOperationLogsResponseBodyLogs struct {
	// 实体 id
	EntityId *string `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	// 实体类型，目前支持Resource
	EntityType *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	// 2021-01-30T12:51:33.028Z
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// 日志
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 操作，目前支持Create, Update, SetDefault
	Operation *string `json:"Operation,omitempty" xml:"Operation,omitempty"`
	// 操作状态，支持 Processing、Succeeded、Failed
	OperationStatus *string `json:"OperationStatus,omitempty" xml:"OperationStatus,omitempty"`
	// 操作人
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
}

func (s ListOperationLogsResponseBodyLogs) String() string {
	return tea.Prettify(s)
}

func (s ListOperationLogsResponseBodyLogs) GoString() string {
	return s.String()
}

func (s *ListOperationLogsResponseBodyLogs) SetEntityId(v string) *ListOperationLogsResponseBodyLogs {
	s.EntityId = &v
	return s
}

func (s *ListOperationLogsResponseBodyLogs) SetEntityType(v string) *ListOperationLogsResponseBodyLogs {
	s.EntityType = &v
	return s
}

func (s *ListOperationLogsResponseBodyLogs) SetGmtCreateTime(v string) *ListOperationLogsResponseBodyLogs {
	s.GmtCreateTime = &v
	return s
}

func (s *ListOperationLogsResponseBodyLogs) SetMessage(v string) *ListOperationLogsResponseBodyLogs {
	s.Message = &v
	return s
}

func (s *ListOperationLogsResponseBodyLogs) SetOperation(v string) *ListOperationLogsResponseBodyLogs {
	s.Operation = &v
	return s
}

func (s *ListOperationLogsResponseBodyLogs) SetOperationStatus(v string) *ListOperationLogsResponseBodyLogs {
	s.OperationStatus = &v
	return s
}

func (s *ListOperationLogsResponseBodyLogs) SetOperator(v string) *ListOperationLogsResponseBodyLogs {
	s.Operator = &v
	return s
}

type ListOperationLogsResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListOperationLogsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListOperationLogsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListOperationLogsResponse) GoString() string {
	return s.String()
}

func (s *ListOperationLogsResponse) SetHeaders(v map[string]*string) *ListOperationLogsResponse {
	s.Headers = v
	return s
}

func (s *ListOperationLogsResponse) SetBody(v *ListOperationLogsResponseBody) *ListOperationLogsResponse {
	s.Body = v
	return s
}

type ListPermissionsResponseBody struct {
	// 权限列表
	Permissions []*ListPermissionsResponseBodyPermissions `json:"Permissions,omitempty" xml:"Permissions,omitempty" type:"Repeated"`
	// 请求 id
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 符合过滤条件的数量
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
	// 权限 code
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
	// 待鉴权实例的可见性 PUBLIC表示当前工作空间所有人都可以访问 PRIVATE表示只有Owner可以访问
	Accessibility *string `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	// 存取类型 CREATOR 代表能看自已创建的 OTHERS 代表可以看别人创建的 ALL代表所有
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
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListPermissionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListPermissionsResponse) SetBody(v *ListPermissionsResponseBody) *ListPermissionsResponse {
	s.Body = v
	return s
}

type ListProductAuthorizationsRequest struct {
	// 逗号分隔的角色名 AliyunODPSPAIDefaultRole,AliyunPAIAccessingOSSRole,AliyunPAIDLCAccessingOSSRole,AliyunPAIDLCDefaultRole,AliyunPAIDSWDefaultRole
	RamRoleNames *string `json:"RamRoleNames,omitempty" xml:"RamRoleNames,omitempty"`
}

func (s ListProductAuthorizationsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListProductAuthorizationsRequest) GoString() string {
	return s.String()
}

func (s *ListProductAuthorizationsRequest) SetRamRoleNames(v string) *ListProductAuthorizationsRequest {
	s.RamRoleNames = &v
	return s
}

type ListProductAuthorizationsResponseBody struct {
	// 角色列表
	AuthorizationDetails []*ListProductAuthorizationsResponseBodyAuthorizationDetails `json:"AuthorizationDetails,omitempty" xml:"AuthorizationDetails,omitempty" type:"Repeated"`
	// 授权链接
	AuthorizationUrl *string `json:"AuthorizationUrl,omitempty" xml:"AuthorizationUrl,omitempty"`
	// Id of the request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListProductAuthorizationsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListProductAuthorizationsResponseBody) GoString() string {
	return s.String()
}

func (s *ListProductAuthorizationsResponseBody) SetAuthorizationDetails(v []*ListProductAuthorizationsResponseBodyAuthorizationDetails) *ListProductAuthorizationsResponseBody {
	s.AuthorizationDetails = v
	return s
}

func (s *ListProductAuthorizationsResponseBody) SetAuthorizationUrl(v string) *ListProductAuthorizationsResponseBody {
	s.AuthorizationUrl = &v
	return s
}

func (s *ListProductAuthorizationsResponseBody) SetRequestId(v string) *ListProductAuthorizationsResponseBody {
	s.RequestId = &v
	return s
}

type ListProductAuthorizationsResponseBodyAuthorizationDetails struct {
	// 快捷授权的url， 假如IsAuthorized为false时有效
	AuthorizationUrl *string `json:"AuthorizationUrl,omitempty" xml:"AuthorizationUrl,omitempty"`
	// 角色是否被授权
	IsAuthorized *bool `json:"IsAuthorized,omitempty" xml:"IsAuthorized,omitempty"`
	// RoleArn
	RamRoleARN *string `json:"RamRoleARN,omitempty" xml:"RamRoleARN,omitempty"`
	// Role的名字
	RamRoleName *string `json:"RamRoleName,omitempty" xml:"RamRoleName,omitempty"`
	// Role的类型
	RamRoleType *string `json:"RamRoleType,omitempty" xml:"RamRoleType,omitempty"`
}

func (s ListProductAuthorizationsResponseBodyAuthorizationDetails) String() string {
	return tea.Prettify(s)
}

func (s ListProductAuthorizationsResponseBodyAuthorizationDetails) GoString() string {
	return s.String()
}

func (s *ListProductAuthorizationsResponseBodyAuthorizationDetails) SetAuthorizationUrl(v string) *ListProductAuthorizationsResponseBodyAuthorizationDetails {
	s.AuthorizationUrl = &v
	return s
}

func (s *ListProductAuthorizationsResponseBodyAuthorizationDetails) SetIsAuthorized(v bool) *ListProductAuthorizationsResponseBodyAuthorizationDetails {
	s.IsAuthorized = &v
	return s
}

func (s *ListProductAuthorizationsResponseBodyAuthorizationDetails) SetRamRoleARN(v string) *ListProductAuthorizationsResponseBodyAuthorizationDetails {
	s.RamRoleARN = &v
	return s
}

func (s *ListProductAuthorizationsResponseBodyAuthorizationDetails) SetRamRoleName(v string) *ListProductAuthorizationsResponseBodyAuthorizationDetails {
	s.RamRoleName = &v
	return s
}

func (s *ListProductAuthorizationsResponseBodyAuthorizationDetails) SetRamRoleType(v string) *ListProductAuthorizationsResponseBodyAuthorizationDetails {
	s.RamRoleType = &v
	return s
}

type ListProductAuthorizationsResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListProductAuthorizationsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListProductAuthorizationsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListProductAuthorizationsResponse) GoString() string {
	return s.String()
}

func (s *ListProductAuthorizationsResponse) SetHeaders(v map[string]*string) *ListProductAuthorizationsResponse {
	s.Headers = v
	return s
}

func (s *ListProductAuthorizationsResponse) SetBody(v *ListProductAuthorizationsResponseBody) *ListProductAuthorizationsResponse {
	s.Body = v
	return s
}

type ListProductsRequest struct {
	// 逗号分割的商品 code
	ProductCodes *string `json:"ProductCodes,omitempty" xml:"ProductCodes,omitempty"`
	// 逗号分割的服务 code
	ServiceCodes *string `json:"ServiceCodes,omitempty" xml:"ServiceCodes,omitempty"`
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

type ListProductsResponseBody struct {
	// 产品列表
	Products []*ListProductsResponseBodyProducts `json:"Products,omitempty" xml:"Products,omitempty" type:"Repeated"`
	// 请求 id
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 服务列表
	Services []*ListProductsResponseBodyServices `json:"Services,omitempty" xml:"Services,omitempty" type:"Repeated"`
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
	// 购买链接
	BuyUrl *string `json:"BuyUrl,omitempty" xml:"BuyUrl,omitempty"`
	// 是否已购买
	IsPurchased *bool `json:"IsPurchased,omitempty" xml:"IsPurchased,omitempty"`
	// 商品 code
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
}

func (s ListProductsResponseBodyProducts) String() string {
	return tea.Prettify(s)
}

func (s ListProductsResponseBodyProducts) GoString() string {
	return s.String()
}

func (s *ListProductsResponseBodyProducts) SetBuyUrl(v string) *ListProductsResponseBodyProducts {
	s.BuyUrl = &v
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

type ListProductsResponseBodyServices struct {
	// 是否开通
	IsOpen *bool `json:"IsOpen,omitempty" xml:"IsOpen,omitempty"`
	// 开通链接
	OpenUrl *string `json:"OpenUrl,omitempty" xml:"OpenUrl,omitempty"`
	// 服务Code
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
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListProductsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListProductsResponse) SetBody(v *ListProductsResponseBody) *ListProductsResponse {
	s.Body = v
	return s
}

type ListQuotasRequest struct {
	// 名字，支持模糊搜索
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
	// 配额列表
	Quotas []*ListQuotasResponseBodyQuotas `json:"Quotas,omitempty" xml:"Quotas,omitempty" type:"Repeated"`
	// 请求 id
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 符合过滤条件的数量
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
	// quota的id
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// 模式  isolate 预付费  share 后付费  develop 开发模式
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// quota名字
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 产品代码
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// 产品类型， 支持PAI，MaxCompute
	QuotaType *string `json:"QuotaType,omitempty" xml:"QuotaType,omitempty"`
	// 规格描述列表
	Specs []*ListQuotasResponseBodyQuotasSpecs `json:"Specs,omitempty" xml:"Specs,omitempty" type:"Repeated"`
}

func (s ListQuotasResponseBodyQuotas) String() string {
	return tea.Prettify(s)
}

func (s ListQuotasResponseBodyQuotas) GoString() string {
	return s.String()
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
	// 规格名
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 类型，可为空
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// 规格描述
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
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListQuotasResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListQuotasResponse) SetBody(v *ListQuotasResponseBody) *ListQuotasResponse {
	s.Body = v
	return s
}

type ListResourcesRequest struct {
	// 选项，目前支持 ListResourceByWorkspace    列举某个工作空间的资源 GetResource                       获取所有工作空间下的资源 GetResourceGroup               获取所有工作空间下的资源组 ListResourceSummaryByWorkspaces  列举一组工作空间的资源汇总
	Option *string `json:"Option,omitempty" xml:"Option,omitempty"`
	// 分页，从1开始，默认1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// 页大小，默认20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 逗号分隔的产品类型，可选值 MaxCompute，DLC
	ProductTypes *string `json:"ProductTypes,omitempty" xml:"ProductTypes,omitempty"`
	// 资源的group名字
	ResourceGroupName *string `json:"ResourceGroupName,omitempty" xml:"ResourceGroupName,omitempty"`
	// 资源的名字
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	// 工作空间 id
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
	// 工作空间id列表，以逗号分隔
	WorkspaceIds *string `json:"WorkspaceIds,omitempty" xml:"WorkspaceIds,omitempty"`
}

func (s ListResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListResourcesRequest) GoString() string {
	return s.String()
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

func (s *ListResourcesRequest) SetResourceGroupName(v string) *ListResourcesRequest {
	s.ResourceGroupName = &v
	return s
}

func (s *ListResourcesRequest) SetResourceName(v string) *ListResourcesRequest {
	s.ResourceName = &v
	return s
}

func (s *ListResourcesRequest) SetWorkspaceId(v string) *ListResourcesRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListResourcesRequest) SetWorkspaceIds(v string) *ListResourcesRequest {
	s.WorkspaceIds = &v
	return s
}

type ListResourcesResponseBody struct {
	// 请求 id
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 资源列表
	Resources []*ListResourcesResponseBodyResources `json:"Resources,omitempty" xml:"Resources,omitempty" type:"Repeated"`
	// 符合过滤条件的作业数量
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
	// 环境， 支持dev（开发）、prod（生产）
	EnvType *string `json:"EnvType,omitempty" xml:"EnvType,omitempty"`
	// 创建 UTC 时间，日期格式 iso8601
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// 分组名，主账户内唯一 一个 GroupName 下可能有一个 dev 资源和一个 prod 资源
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// 是否默认资源 每个类型都有一个默认的资源
	IsDefault *bool `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	// 资源名 长度需要在3到27个字符 region内唯一
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 产品类型， 支持PAI，MaxCompute
	ProductType     *string                                              `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	Quotas          []*ListResourcesResponseBodyResourcesQuotas          `json:"Quotas,omitempty" xml:"Quotas,omitempty" type:"Repeated"`
	ResourceSummary []*ListResourcesResponseBodyResourcesResourceSummary `json:"ResourceSummary,omitempty" xml:"ResourceSummary,omitempty" type:"Repeated"`
	// 对于MaxCompute是个json，有如下key： Endpoint Project
	Spec map[string]interface{} `json:"Spec,omitempty" xml:"Spec,omitempty"`
	// 所属的工作空间 id
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
	// 资源 id
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
}

func (s ListResourcesResponseBodyResources) String() string {
	return tea.Prettify(s)
}

func (s ListResourcesResponseBodyResources) GoString() string {
	return s.String()
}

func (s *ListResourcesResponseBodyResources) SetEnvType(v string) *ListResourcesResponseBodyResources {
	s.EnvType = &v
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

func (s *ListResourcesResponseBodyResources) SetIsDefault(v bool) *ListResourcesResponseBodyResources {
	s.IsDefault = &v
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

func (s *ListResourcesResponseBodyResources) SetResourceSummary(v []*ListResourcesResponseBodyResourcesResourceSummary) *ListResourcesResponseBodyResources {
	s.ResourceSummary = v
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

func (s *ListResourcesResponseBodyResources) SetId(v string) *ListResourcesResponseBodyResources {
	s.Id = &v
	return s
}

type ListResourcesResponseBodyResourcesQuotas struct {
	// 卡类型，支持cpu、gpu
	CardType *string `json:"CardType,omitempty" xml:"CardType,omitempty"`
	// 配额id
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// 模式 isolate 预付费 share 后付费 develop 开发模式
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// 配额名称
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 商品 code
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// 产品类型， 支持PAI，MaxCompute
	QuotaType *string `json:"QuotaType,omitempty" xml:"QuotaType,omitempty"`
	// 规格描述列表
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
	// 规格名字
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 规格描述
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

type ListResourcesResponseBodyResourcesResourceSummary struct {
	// 资源个数
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// 资源类型
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
}

func (s ListResourcesResponseBodyResourcesResourceSummary) String() string {
	return tea.Prettify(s)
}

func (s ListResourcesResponseBodyResourcesResourceSummary) GoString() string {
	return s.String()
}

func (s *ListResourcesResponseBodyResourcesResourceSummary) SetCount(v int64) *ListResourcesResponseBodyResourcesResourceSummary {
	s.Count = &v
	return s
}

func (s *ListResourcesResponseBodyResourcesResourceSummary) SetProductType(v string) *ListResourcesResponseBodyResourcesResourceSummary {
	s.ProductType = &v
	return s
}

type ListResourcesResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListResourcesResponse) SetBody(v *ListResourcesResponseBody) *ListResourcesResponse {
	s.Body = v
	return s
}

type ListUsersRequest struct {
	// 账户类型列表，以逗号分隔
	AccountTypes *string `json:"AccountTypes,omitempty" xml:"AccountTypes,omitempty"`
	// 分页，从1开始，默认1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// 页大小，默认20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 用户 Id 列表，以逗号分隔
	UserIds *string `json:"UserIds,omitempty" xml:"UserIds,omitempty"`
	// 用户名
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s ListUsersRequest) String() string {
	return tea.Prettify(s)
}

func (s ListUsersRequest) GoString() string {
	return s.String()
}

func (s *ListUsersRequest) SetAccountTypes(v string) *ListUsersRequest {
	s.AccountTypes = &v
	return s
}

func (s *ListUsersRequest) SetPageNumber(v int64) *ListUsersRequest {
	s.PageNumber = &v
	return s
}

func (s *ListUsersRequest) SetPageSize(v int64) *ListUsersRequest {
	s.PageSize = &v
	return s
}

func (s *ListUsersRequest) SetUserIds(v string) *ListUsersRequest {
	s.UserIds = &v
	return s
}

func (s *ListUsersRequest) SetUserName(v string) *ListUsersRequest {
	s.UserName = &v
	return s
}

type ListUsersResponseBody struct {
	// 请求 id
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 符合过滤条件的用户数量
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// 用户列表
	Users []*ListUsersResponseBodyUsers `json:"Users,omitempty" xml:"Users,omitempty" type:"Repeated"`
}

func (s ListUsersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListUsersResponseBody) GoString() string {
	return s.String()
}

func (s *ListUsersResponseBody) SetRequestId(v string) *ListUsersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUsersResponseBody) SetTotalCount(v int64) *ListUsersResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListUsersResponseBody) SetUsers(v []*ListUsersResponseBodyUsers) *ListUsersResponseBody {
	s.Users = v
	return s
}

type ListUsersResponseBodyUsers struct {
	// 用户显示名
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// 用户 id
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// 用户名
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s ListUsersResponseBodyUsers) String() string {
	return tea.Prettify(s)
}

func (s ListUsersResponseBodyUsers) GoString() string {
	return s.String()
}

func (s *ListUsersResponseBodyUsers) SetDisplayName(v string) *ListUsersResponseBodyUsers {
	s.DisplayName = &v
	return s
}

func (s *ListUsersResponseBodyUsers) SetUserId(v string) *ListUsersResponseBodyUsers {
	s.UserId = &v
	return s
}

func (s *ListUsersResponseBodyUsers) SetUserName(v string) *ListUsersResponseBodyUsers {
	s.UserName = &v
	return s
}

type ListUsersResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListUsersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListUsersResponse) String() string {
	return tea.Prettify(s)
}

func (s ListUsersResponse) GoString() string {
	return s.String()
}

func (s *ListUsersResponse) SetHeaders(v map[string]*string) *ListUsersResponse {
	s.Headers = v
	return s
}

func (s *ListUsersResponse) SetBody(v *ListUsersResponseBody) *ListUsersResponse {
	s.Body = v
	return s
}

type ListWorkspaceUsersResponseBody struct {
	// 请求 id
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 符合过滤条件的用户数量
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// 用户列表
	Users []*ListWorkspaceUsersResponseBodyUsers `json:"Users,omitempty" xml:"Users,omitempty" type:"Repeated"`
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
	// 用户 id
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// 用户名
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
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListWorkspaceUsersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListWorkspaceUsersResponse) SetBody(v *ListWorkspaceUsersResponseBody) *ListWorkspaceUsersResponse {
	s.Body = v
	return s
}

type ListWorkspacesRequest struct {
	// GetWorkspaceFields的Field字段
	Fields *string `json:"Fields,omitempty" xml:"Fields,omitempty"`
	// 逗号分割的模块列表，目前填入PAI
	ModuleList *string `json:"ModuleList,omitempty" xml:"ModuleList,omitempty"`
	// 逗号分隔的选项
	Option *string `json:"Option,omitempty" xml:"Option,omitempty"`
	// 排序方向： ASC - 升序 DESC - 降序
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// 分页，从1开始，默认1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// 页大小，默认20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 排序字段：CreateTime
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// 状态
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 是否显示详细信息，默认true
	Verbose *bool `json:"Verbose,omitempty" xml:"Verbose,omitempty"`
	// 工作空间id列表
	WorkspaceIds *string `json:"WorkspaceIds,omitempty" xml:"WorkspaceIds,omitempty"`
	// 工作空间名字
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
	// 请求 id
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 资源限制
	ResourceLimits map[string]interface{} `json:"ResourceLimits,omitempty" xml:"ResourceLimits,omitempty"`
	// 符合过滤条件的作业数量
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// 工作空间列表
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
	// 管理员名字
	AdminNames []*string `json:"AdminNames,omitempty" xml:"AdminNames,omitempty" type:"Repeated"`
	// 创建人
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// 描述
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 环境，用作判断简单模式还是标准模式
	EnvTypes []*string `json:"EnvTypes,omitempty" xml:"EnvTypes,omitempty" type:"Repeated"`
	// 附加信息
	ExtraInfos map[string]interface{} `json:"ExtraInfos,omitempty" xml:"ExtraInfos,omitempty"`
	// 创建 UTC 时间，日期格式 iso8601
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// 修改 UTC 时间，日期格式 iso8601
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// 是否为默认工作空间
	IsDefault *bool `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	// 资源数目
	ResourceCount *int32 `json:"ResourceCount,omitempty" xml:"ResourceCount,omitempty"`
	// 工作空间状态
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 工作空间 id
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
	// 工作空间名字
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

func (s *ListWorkspacesResponseBodyWorkspaces) SetResourceCount(v int32) *ListWorkspacesResponseBodyWorkspaces {
	s.ResourceCount = &v
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
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListWorkspacesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListWorkspacesResponse) SetBody(v *ListWorkspacesResponseBody) *ListWorkspacesResponse {
	s.Body = v
	return s
}

type PublishCodeSourceResponseBody struct {
	// 被删除的代码源配置ID
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
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *PublishCodeSourceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *PublishCodeSourceResponse) SetBody(v *PublishCodeSourceResponseBody) *PublishCodeSourceResponse {
	s.Body = v
	return s
}

type PublishDatasetResponseBody struct {
	// Id of the request
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
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *PublishDatasetResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *PublishDatasetResponse) SetBody(v *PublishDatasetResponseBody) *PublishDatasetResponse {
	s.Body = v
	return s
}

type PublishImageResponseBody struct {
	// 镜像 id
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// Id of the request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
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
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *PublishImageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *PublishImageResponse) SetBody(v *PublishImageResponseBody) *PublishImageResponse {
	s.Body = v
	return s
}

type RemoveImageResponseBody struct {
	// Id of the request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
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
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RemoveImageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *RemoveImageResponse) SetBody(v *RemoveImageResponseBody) *RemoveImageResponse {
	s.Body = v
	return s
}

type RemoveImageLabelsResponseBody struct {
	// Id of the request
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
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
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RemoveImageLabelsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *RemoveImageLabelsResponse) SetBody(v *RemoveImageLabelsResponseBody) *RemoveImageLabelsResponse {
	s.Body = v
	return s
}

type RemoveMemberRoleResponseBody struct {
	// 请求 id
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
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RemoveMemberRoleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *RemoveMemberRoleResponse) SetBody(v *RemoveMemberRoleResponseBody) *RemoveMemberRoleResponse {
	s.Body = v
	return s
}

type RemoveWorkspaceQuotaResponseBody struct {
	// 请求 id
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RemoveWorkspaceQuotaResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveWorkspaceQuotaResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveWorkspaceQuotaResponseBody) SetRequestId(v string) *RemoveWorkspaceQuotaResponseBody {
	s.RequestId = &v
	return s
}

type RemoveWorkspaceQuotaResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RemoveWorkspaceQuotaResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RemoveWorkspaceQuotaResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveWorkspaceQuotaResponse) GoString() string {
	return s.String()
}

func (s *RemoveWorkspaceQuotaResponse) SetHeaders(v map[string]*string) *RemoveWorkspaceQuotaResponse {
	s.Headers = v
	return s
}

func (s *RemoveWorkspaceQuotaResponse) SetBody(v *RemoveWorkspaceQuotaResponseBody) *RemoveWorkspaceQuotaResponse {
	s.Body = v
	return s
}

type UpdateConfigsRequest struct {
	// 配置列表
	Configs []*UpdateConfigsRequestConfigs `json:"Configs,omitempty" xml:"Configs,omitempty" type:"Repeated"`
}

func (s UpdateConfigsRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateConfigsRequest) GoString() string {
	return s.String()
}

func (s *UpdateConfigsRequest) SetConfigs(v []*UpdateConfigsRequestConfigs) *UpdateConfigsRequest {
	s.Configs = v
	return s
}

type UpdateConfigsRequestConfigs struct {
	// 配置键
	ConfigKey *string `json:"ConfigKey,omitempty" xml:"ConfigKey,omitempty"`
	// 配置值
	ConfigValue *string `json:"ConfigValue,omitempty" xml:"ConfigValue,omitempty"`
}

func (s UpdateConfigsRequestConfigs) String() string {
	return tea.Prettify(s)
}

func (s UpdateConfigsRequestConfigs) GoString() string {
	return s.String()
}

func (s *UpdateConfigsRequestConfigs) SetConfigKey(v string) *UpdateConfigsRequestConfigs {
	s.ConfigKey = &v
	return s
}

func (s *UpdateConfigsRequestConfigs) SetConfigValue(v string) *UpdateConfigsRequestConfigs {
	s.ConfigValue = &v
	return s
}

type UpdateConfigsResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateConfigsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateConfigsResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateConfigsResponseBody) SetRequestId(v string) *UpdateConfigsResponseBody {
	s.RequestId = &v
	return s
}

type UpdateConfigsResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateConfigsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateConfigsResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateConfigsResponse) GoString() string {
	return s.String()
}

func (s *UpdateConfigsResponse) SetHeaders(v map[string]*string) *UpdateConfigsResponse {
	s.Headers = v
	return s
}

func (s *UpdateConfigsResponse) SetBody(v *UpdateConfigsResponseBody) *UpdateConfigsResponse {
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
	// Id of the request
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
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateDatasetResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *UpdateDatasetResponse) SetBody(v *UpdateDatasetResponseBody) *UpdateDatasetResponse {
	s.Body = v
	return s
}

type UpdateWorkspaceRequest struct {
	// 描述
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 显示名称
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
	// 请求 id
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
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateWorkspaceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *UpdateWorkspaceResponse) SetBody(v *UpdateWorkspaceResponseBody) *UpdateWorkspaceResponse {
	s.Body = v
	return s
}

type UpdateWorkspaceResourceRequest struct {
	// 是否默认资源实例，目前只能填 true，不支持填 false
	IsDefault *bool `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
}

func (s UpdateWorkspaceResourceRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateWorkspaceResourceRequest) GoString() string {
	return s.String()
}

func (s *UpdateWorkspaceResourceRequest) SetIsDefault(v bool) *UpdateWorkspaceResourceRequest {
	s.IsDefault = &v
	return s
}

type UpdateWorkspaceResourceResponseBody struct {
	// 请求 id
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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

type UpdateWorkspaceResourceResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateWorkspaceResourceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

	if !tea.BoolValue(util.IsUnset(request.ImageUri)) {
		body["ImageUri"] = request.ImageUri
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

func (client *Client) AddImageLabelsWithOptions(ImageId *string, request *AddImageLabelsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *AddImageLabelsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	ImageId = openapiutil.GetEncodeParam(ImageId)
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
		Pathname:    tea.String("/api/v1/images/" + tea.StringValue(ImageId) + "/labels"),
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

func (client *Client) AddMemberRoleWithOptions(WorkspaceId *string, MemberId *string, RoleName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *AddMemberRoleResponse, _err error) {
	WorkspaceId = openapiutil.GetEncodeParam(WorkspaceId)
	MemberId = openapiutil.GetEncodeParam(MemberId)
	RoleName = openapiutil.GetEncodeParam(RoleName)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("AddMemberRole"),
		Version:     tea.String("2021-02-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/workspaces/" + tea.StringValue(WorkspaceId) + "/members/" + tea.StringValue(MemberId) + "/roles/" + tea.StringValue(RoleName)),
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

func (client *Client) AddWorkspaceQuota(WorkspaceId *string, QuotaId *string, request *AddWorkspaceQuotaRequest) (_result *AddWorkspaceQuotaResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &AddWorkspaceQuotaResponse{}
	_body, _err := client.AddWorkspaceQuotaWithOptions(WorkspaceId, QuotaId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddWorkspaceQuotaWithOptions(WorkspaceId *string, QuotaId *string, request *AddWorkspaceQuotaRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *AddWorkspaceQuotaResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	WorkspaceId = openapiutil.GetEncodeParam(WorkspaceId)
	QuotaId = openapiutil.GetEncodeParam(QuotaId)
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Mode)) {
		body["Mode"] = request.Mode
	}

	if !tea.BoolValue(util.IsUnset(request.ProductCode)) {
		body["ProductCode"] = request.ProductCode
	}

	if !tea.BoolValue(util.IsUnset(request.QuotaType)) {
		body["QuotaType"] = request.QuotaType
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("AddWorkspaceQuota"),
		Version:     tea.String("2021-02-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/workspaces/[WorkspaceId]/quotas/[QuotaId]"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &AddWorkspaceQuotaResponse{}
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

func (client *Client) CreateDatasetLabelsWithOptions(DatasetId *string, request *CreateDatasetLabelsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateDatasetLabelsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	DatasetId = openapiutil.GetEncodeParam(DatasetId)
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
		Pathname:    tea.String("/api/v1/datasets/" + tea.StringValue(DatasetId) + "/labels"),
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

func (client *Client) CreateDefaultWorkspace(request *CreateDefaultWorkspaceRequest) (_result *CreateDefaultWorkspaceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateDefaultWorkspaceResponse{}
	_body, _err := client.CreateDefaultWorkspaceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateDefaultWorkspaceWithOptions(request *CreateDefaultWorkspaceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateDefaultWorkspaceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.EnvTypes)) {
		body["EnvTypes"] = request.EnvTypes
	}

	if !tea.BoolValue(util.IsUnset(request.Resources)) {
		body["Resources"] = request.Resources
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateDefaultWorkspace"),
		Version:     tea.String("2021-02-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/defaultWorkspaces"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateDefaultWorkspaceResponse{}
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

func (client *Client) CreateMemberWithOptions(WorkspaceId *string, request *CreateMemberRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateMemberResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	WorkspaceId = openapiutil.GetEncodeParam(WorkspaceId)
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
		Pathname:    tea.String("/api/v1/workspaces/" + tea.StringValue(WorkspaceId) + "/members"),
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

func (client *Client) CreateProductOrdersWithOptions(request *CreateProductOrdersRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateProductOrdersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AutoPay)) {
		body["AutoPay"] = request.AutoPay
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.Products))) {
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

func (client *Client) CreateUser() (_result *CreateUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateUserResponse{}
	_body, _err := client.CreateUserWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateUserWithOptions(headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateUserResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("CreateUser"),
		Version:     tea.String("2021-02-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/users"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateUserResponse{}
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

func (client *Client) CreateWorkspaceResourceWithOptions(WorkspaceId *string, request *CreateWorkspaceResourceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateWorkspaceResourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	WorkspaceId = openapiutil.GetEncodeParam(WorkspaceId)
	body := map[string]interface{}{}
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
		Pathname:    tea.String("/api/v1/workspaces/" + tea.StringValue(WorkspaceId) + "/resources"),
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

func (client *Client) DeleteCodeSourceWithOptions(CodeSourceId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteCodeSourceResponse, _err error) {
	CodeSourceId = openapiutil.GetEncodeParam(CodeSourceId)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteCodeSource"),
		Version:     tea.String("2021-02-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/codesources/" + tea.StringValue(CodeSourceId)),
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

func (client *Client) DeleteConfig(WorkspaceId *string, ConfigKey *string) (_result *DeleteConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteConfigResponse{}
	_body, _err := client.DeleteConfigWithOptions(WorkspaceId, ConfigKey, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteConfigWithOptions(WorkspaceId *string, ConfigKey *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteConfigResponse, _err error) {
	WorkspaceId = openapiutil.GetEncodeParam(WorkspaceId)
	ConfigKey = openapiutil.GetEncodeParam(ConfigKey)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteConfig"),
		Version:     tea.String("2021-02-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/workspaces/" + tea.StringValue(WorkspaceId) + "/configs/" + tea.StringValue(ConfigKey)),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteConfigResponse{}
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

func (client *Client) DeleteDatasetWithOptions(DatasetId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteDatasetResponse, _err error) {
	DatasetId = openapiutil.GetEncodeParam(DatasetId)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteDataset"),
		Version:     tea.String("2021-02-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/datasets/" + tea.StringValue(DatasetId)),
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

func (client *Client) DeleteDatasetLabelsWithOptions(DatasetId *string, request *DeleteDatasetLabelsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteDatasetLabelsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	DatasetId = openapiutil.GetEncodeParam(DatasetId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Keys)) {
		query["Keys"] = request.Keys
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteDatasetLabels"),
		Version:     tea.String("2021-02-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/datasets/" + tea.StringValue(DatasetId) + "/labels"),
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

func (client *Client) DeleteMembersWithOptions(WorkspaceId *string, request *DeleteMembersRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteMembersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	WorkspaceId = openapiutil.GetEncodeParam(WorkspaceId)
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
		Pathname:    tea.String("/api/v1/workspaces/" + tea.StringValue(WorkspaceId) + "/members"),
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

func (client *Client) DeleteWorkspaceWithOptions(WorkspaceId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteWorkspaceResponse, _err error) {
	WorkspaceId = openapiutil.GetEncodeParam(WorkspaceId)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteWorkspace"),
		Version:     tea.String("2021-02-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/workspaces/" + tea.StringValue(WorkspaceId)),
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

func (client *Client) GetCodeSourceWithOptions(CodeSourceId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetCodeSourceResponse, _err error) {
	CodeSourceId = openapiutil.GetEncodeParam(CodeSourceId)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetCodeSource"),
		Version:     tea.String("2021-02-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/codesources/" + tea.StringValue(CodeSourceId)),
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

func (client *Client) GetCodeSourcesStatistics(request *GetCodeSourcesStatisticsRequest) (_result *GetCodeSourcesStatisticsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetCodeSourcesStatisticsResponse{}
	_body, _err := client.GetCodeSourcesStatisticsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetCodeSourcesStatisticsWithOptions(request *GetCodeSourcesStatisticsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetCodeSourcesStatisticsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetCodeSourcesStatistics"),
		Version:     tea.String("2021-02-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/statistics/codesources"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetCodeSourcesStatisticsResponse{}
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

func (client *Client) GetDatasetWithOptions(DatasetId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetDatasetResponse, _err error) {
	DatasetId = openapiutil.GetEncodeParam(DatasetId)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetDataset"),
		Version:     tea.String("2021-02-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/datasets/" + tea.StringValue(DatasetId)),
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

func (client *Client) GetDatasetsStatistics(request *GetDatasetsStatisticsRequest) (_result *GetDatasetsStatisticsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetDatasetsStatisticsResponse{}
	_body, _err := client.GetDatasetsStatisticsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetDatasetsStatisticsWithOptions(request *GetDatasetsStatisticsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetDatasetsStatisticsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetDatasetsStatistics"),
		Version:     tea.String("2021-02-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/statistics/datasets"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDatasetsStatisticsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetDefaultWorkspace() (_result *GetDefaultWorkspaceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetDefaultWorkspaceResponse{}
	_body, _err := client.GetDefaultWorkspaceWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetDefaultWorkspaceWithOptions(headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetDefaultWorkspaceResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
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

func (client *Client) GetImageWithOptions(ImageId *string, request *GetImageRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetImageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	ImageId = openapiutil.GetEncodeParam(ImageId)
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
		Pathname:    tea.String("/api/v1/images/" + tea.StringValue(ImageId)),
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

func (client *Client) GetImagesStatistics(request *GetImagesStatisticsRequest) (_result *GetImagesStatisticsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetImagesStatisticsResponse{}
	_body, _err := client.GetImagesStatisticsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetImagesStatisticsWithOptions(request *GetImagesStatisticsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetImagesStatisticsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetImagesStatistics"),
		Version:     tea.String("2021-02-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/statistics/images"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetImagesStatisticsResponse{}
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

func (client *Client) GetMemberWithOptions(WorkspaceId *string, request *GetMemberRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetMemberResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	WorkspaceId = openapiutil.GetEncodeParam(WorkspaceId)
	query := map[string]interface{}{}
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
		Pathname:    tea.String("/api/v1/workspaces/" + tea.StringValue(WorkspaceId) + "/member"),
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

func (client *Client) GetPermissionWithOptions(WorkspaceId *string, PermissionCode *string, request *GetPermissionRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetPermissionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	WorkspaceId = openapiutil.GetEncodeParam(WorkspaceId)
	PermissionCode = openapiutil.GetEncodeParam(PermissionCode)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Accessibility)) {
		query["Accessibility"] = request.Accessibility
	}

	if !tea.BoolValue(util.IsUnset(request.Creator)) {
		query["Creator"] = request.Creator
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetPermission"),
		Version:     tea.String("2021-02-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/workspaces/" + tea.StringValue(WorkspaceId) + "/permissions/" + tea.StringValue(PermissionCode)),
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

func (client *Client) GetRoleStatistics(request *GetRoleStatisticsRequest) (_result *GetRoleStatisticsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetRoleStatisticsResponse{}
	_body, _err := client.GetRoleStatisticsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetRoleStatisticsWithOptions(request *GetRoleStatisticsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetRoleStatisticsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetRoleStatistics"),
		Version:     tea.String("2021-02-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/statistics/roles"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetRoleStatisticsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetWorkspace(WorkspaceId *string) (_result *GetWorkspaceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetWorkspaceResponse{}
	_body, _err := client.GetWorkspaceWithOptions(WorkspaceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetWorkspaceWithOptions(WorkspaceId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetWorkspaceResponse, _err error) {
	WorkspaceId = openapiutil.GetEncodeParam(WorkspaceId)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetWorkspace"),
		Version:     tea.String("2021-02-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/workspaces/" + tea.StringValue(WorkspaceId)),
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

func (client *Client) ListConfigs(WorkspaceId *string, request *ListConfigsRequest) (_result *ListConfigsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListConfigsResponse{}
	_body, _err := client.ListConfigsWithOptions(WorkspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListConfigsWithOptions(WorkspaceId *string, request *ListConfigsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListConfigsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	WorkspaceId = openapiutil.GetEncodeParam(WorkspaceId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConfigKeys)) {
		query["ConfigKeys"] = request.ConfigKeys
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListConfigs"),
		Version:     tea.String("2021-02-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/workspaces/" + tea.StringValue(WorkspaceId) + "/configs"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListConfigsResponse{}
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

	if !tea.BoolValue(util.IsUnset(request.LabelKeys)) {
		query["LabelKeys"] = request.LabelKeys
	}

	if !tea.BoolValue(util.IsUnset(request.LabelValues)) {
		query["LabelValues"] = request.LabelValues
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

func (client *Client) ListFeatures(request *ListFeaturesRequest) (_result *ListFeaturesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListFeaturesResponse{}
	_body, _err := client.ListFeaturesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListFeaturesWithOptions(request *ListFeaturesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListFeaturesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Names)) {
		query["Names"] = request.Names
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListFeatures"),
		Version:     tea.String("2021-02-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/features"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListFeaturesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListGlobalPermissions(WorkspaceId *string) (_result *ListGlobalPermissionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListGlobalPermissionsResponse{}
	_body, _err := client.ListGlobalPermissionsWithOptions(WorkspaceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListGlobalPermissionsWithOptions(WorkspaceId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListGlobalPermissionsResponse, _err error) {
	WorkspaceId = openapiutil.GetEncodeParam(WorkspaceId)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("ListGlobalPermissions"),
		Version:     tea.String("2021-02-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/permissions"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListGlobalPermissionsResponse{}
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

func (client *Client) ListImagesWithOptions(request *ListImagesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListImagesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Labels)) {
		query["Labels"] = request.Labels
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.OperatorCreate)) {
		query["OperatorCreate"] = request.OperatorCreate
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

func (client *Client) ListMembersWithOptions(WorkspaceId *string, request *ListMembersRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListMembersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	WorkspaceId = openapiutil.GetEncodeParam(WorkspaceId)
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
		Pathname:    tea.String("/api/v1/workspaces/" + tea.StringValue(WorkspaceId) + "/members"),
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

func (client *Client) ListOperationLogs(WorkspaceId *string, request *ListOperationLogsRequest) (_result *ListOperationLogsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListOperationLogsResponse{}
	_body, _err := client.ListOperationLogsWithOptions(WorkspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListOperationLogsWithOptions(WorkspaceId *string, request *ListOperationLogsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListOperationLogsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	WorkspaceId = openapiutil.GetEncodeParam(WorkspaceId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EntityStatus)) {
		query["EntityStatus"] = request.EntityStatus
	}

	if !tea.BoolValue(util.IsUnset(request.EntityTypes)) {
		query["EntityTypes"] = request.EntityTypes
	}

	if !tea.BoolValue(util.IsUnset(request.OperationStatus)) {
		query["OperationStatus"] = request.OperationStatus
	}

	if !tea.BoolValue(util.IsUnset(request.Operations)) {
		query["Operations"] = request.Operations
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

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListOperationLogs"),
		Version:     tea.String("2021-02-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/workspaces/" + tea.StringValue(WorkspaceId) + "/logs"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListOperationLogsResponse{}
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

func (client *Client) ListPermissionsWithOptions(WorkspaceId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListPermissionsResponse, _err error) {
	WorkspaceId = openapiutil.GetEncodeParam(WorkspaceId)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("ListPermissions"),
		Version:     tea.String("2021-02-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/workspaces/" + tea.StringValue(WorkspaceId) + "/permissions"),
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

func (client *Client) ListProductAuthorizations(request *ListProductAuthorizationsRequest) (_result *ListProductAuthorizationsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListProductAuthorizationsResponse{}
	_body, _err := client.ListProductAuthorizationsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListProductAuthorizationsWithOptions(request *ListProductAuthorizationsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListProductAuthorizationsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RamRoleNames)) {
		query["RamRoleNames"] = request.RamRoleNames
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListProductAuthorizations"),
		Version:     tea.String("2021-02-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/productauthorizations"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListProductAuthorizationsResponse{}
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

func (client *Client) ListResourcesWithOptions(request *ListResourcesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
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

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupName)) {
		query["ResourceGroupName"] = request.ResourceGroupName
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceName)) {
		query["ResourceName"] = request.ResourceName
	}

	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	if !tea.BoolValue(util.IsUnset(request.WorkspaceIds)) {
		query["WorkspaceIds"] = request.WorkspaceIds
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

func (client *Client) ListUsers(request *ListUsersRequest) (_result *ListUsersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListUsersResponse{}
	_body, _err := client.ListUsersWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListUsersWithOptions(request *ListUsersRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListUsersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountTypes)) {
		query["AccountTypes"] = request.AccountTypes
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.UserIds)) {
		query["UserIds"] = request.UserIds
	}

	if !tea.BoolValue(util.IsUnset(request.UserName)) {
		query["UserName"] = request.UserName
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListUsers"),
		Version:     tea.String("2021-02-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/users"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListUsersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListWorkspaceUsers(WorkspaceId *string) (_result *ListWorkspaceUsersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListWorkspaceUsersResponse{}
	_body, _err := client.ListWorkspaceUsersWithOptions(WorkspaceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListWorkspaceUsersWithOptions(WorkspaceId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListWorkspaceUsersResponse, _err error) {
	WorkspaceId = openapiutil.GetEncodeParam(WorkspaceId)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("ListWorkspaceUsers"),
		Version:     tea.String("2021-02-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/workspaces/" + tea.StringValue(WorkspaceId) + "/users"),
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

func (client *Client) PublishCodeSourceWithOptions(CodeSourceId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *PublishCodeSourceResponse, _err error) {
	CodeSourceId = openapiutil.GetEncodeParam(CodeSourceId)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("PublishCodeSource"),
		Version:     tea.String("2021-02-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/codesources/" + tea.StringValue(CodeSourceId) + "/publish"),
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

func (client *Client) PublishDatasetWithOptions(DatasetId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *PublishDatasetResponse, _err error) {
	DatasetId = openapiutil.GetEncodeParam(DatasetId)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("PublishDataset"),
		Version:     tea.String("2021-02-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/datasets/" + tea.StringValue(DatasetId) + "/publish"),
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

func (client *Client) PublishImageWithOptions(ImageId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *PublishImageResponse, _err error) {
	ImageId = openapiutil.GetEncodeParam(ImageId)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("PublishImage"),
		Version:     tea.String("2021-02-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/images/" + tea.StringValue(ImageId) + "/publish"),
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

func (client *Client) RemoveImageWithOptions(ImageId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *RemoveImageResponse, _err error) {
	ImageId = openapiutil.GetEncodeParam(ImageId)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("RemoveImage"),
		Version:     tea.String("2021-02-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/images/" + tea.StringValue(ImageId)),
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

func (client *Client) RemoveImageLabelsWithOptions(ImageId *string, LabelKey *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *RemoveImageLabelsResponse, _err error) {
	ImageId = openapiutil.GetEncodeParam(ImageId)
	LabelKey = openapiutil.GetEncodeParam(LabelKey)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("RemoveImageLabels"),
		Version:     tea.String("2021-02-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/images/" + tea.StringValue(ImageId) + "/labels/" + tea.StringValue(LabelKey)),
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

func (client *Client) RemoveMemberRoleWithOptions(WorkspaceId *string, MemberId *string, RoleName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *RemoveMemberRoleResponse, _err error) {
	WorkspaceId = openapiutil.GetEncodeParam(WorkspaceId)
	MemberId = openapiutil.GetEncodeParam(MemberId)
	RoleName = openapiutil.GetEncodeParam(RoleName)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("RemoveMemberRole"),
		Version:     tea.String("2021-02-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/workspaces/" + tea.StringValue(WorkspaceId) + "/members/" + tea.StringValue(MemberId) + "/roles/" + tea.StringValue(RoleName)),
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

func (client *Client) RemoveWorkspaceQuota(WorkspaceId *string, QuotaId *string) (_result *RemoveWorkspaceQuotaResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RemoveWorkspaceQuotaResponse{}
	_body, _err := client.RemoveWorkspaceQuotaWithOptions(WorkspaceId, QuotaId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RemoveWorkspaceQuotaWithOptions(WorkspaceId *string, QuotaId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *RemoveWorkspaceQuotaResponse, _err error) {
	WorkspaceId = openapiutil.GetEncodeParam(WorkspaceId)
	QuotaId = openapiutil.GetEncodeParam(QuotaId)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("RemoveWorkspaceQuota"),
		Version:     tea.String("2021-02-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/workspaces/" + tea.StringValue(WorkspaceId) + "/quotas/" + tea.StringValue(QuotaId)),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &RemoveWorkspaceQuotaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateConfigs(WorkspaceId *string, request *UpdateConfigsRequest) (_result *UpdateConfigsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateConfigsResponse{}
	_body, _err := client.UpdateConfigsWithOptions(WorkspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateConfigsWithOptions(WorkspaceId *string, request *UpdateConfigsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateConfigsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	WorkspaceId = openapiutil.GetEncodeParam(WorkspaceId)
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Configs)) {
		body["Configs"] = request.Configs
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateConfigs"),
		Version:     tea.String("2021-02-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/workspaces/" + tea.StringValue(WorkspaceId) + "/configs"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateConfigsResponse{}
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

func (client *Client) UpdateDatasetWithOptions(DatasetId *string, request *UpdateDatasetRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateDatasetResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	DatasetId = openapiutil.GetEncodeParam(DatasetId)
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
		Pathname:    tea.String("/api/v1/datasets/" + tea.StringValue(DatasetId)),
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

func (client *Client) UpdateWorkspaceWithOptions(WorkspaceId *string, request *UpdateWorkspaceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateWorkspaceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	WorkspaceId = openapiutil.GetEncodeParam(WorkspaceId)
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
		Pathname:    tea.String("/api/v1/workspaces/" + tea.StringValue(WorkspaceId)),
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

func (client *Client) UpdateWorkspaceResource(WorkspaceId *string, ResourceGroupName *string, request *UpdateWorkspaceResourceRequest) (_result *UpdateWorkspaceResourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateWorkspaceResourceResponse{}
	_body, _err := client.UpdateWorkspaceResourceWithOptions(WorkspaceId, ResourceGroupName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateWorkspaceResourceWithOptions(WorkspaceId *string, ResourceGroupName *string, request *UpdateWorkspaceResourceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateWorkspaceResourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	WorkspaceId = openapiutil.GetEncodeParam(WorkspaceId)
	ResourceGroupName = openapiutil.GetEncodeParam(ResourceGroupName)
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IsDefault)) {
		body["IsDefault"] = request.IsDefault
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateWorkspaceResource"),
		Version:     tea.String("2021-02-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/workspaces/" + tea.StringValue(WorkspaceId) + "/resources/" + tea.StringValue(ResourceGroupName)),
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
