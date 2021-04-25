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

type AbortRunRequest struct {
	// 工作空间名称
	Workspace *string `json:"Workspace,omitempty" xml:"Workspace,omitempty"`
	// 任务ID
	RunId *string `json:"RunId,omitempty" xml:"RunId,omitempty"`
}

func (s AbortRunRequest) String() string {
	return tea.Prettify(s)
}

func (s AbortRunRequest) GoString() string {
	return s.String()
}

func (s *AbortRunRequest) SetWorkspace(v string) *AbortRunRequest {
	s.Workspace = &v
	return s
}

func (s *AbortRunRequest) SetRunId(v string) *AbortRunRequest {
	s.RunId = &v
	return s
}

type AbortRunResponseBody struct {
	// 主机ID
	HostId *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	// 请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AbortRunResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AbortRunResponseBody) GoString() string {
	return s.String()
}

func (s *AbortRunResponseBody) SetHostId(v string) *AbortRunResponseBody {
	s.HostId = &v
	return s
}

func (s *AbortRunResponseBody) SetRequestId(v string) *AbortRunResponseBody {
	s.RequestId = &v
	return s
}

type AbortRunResponse struct {
	Headers map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AbortRunResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AbortRunResponse) String() string {
	return tea.Prettify(s)
}

func (s AbortRunResponse) GoString() string {
	return s.String()
}

func (s *AbortRunResponse) SetHeaders(v map[string]*string) *AbortRunResponse {
	s.Headers = v
	return s
}

func (s *AbortRunResponse) SetBody(v *AbortRunResponseBody) *AbortRunResponse {
	s.Body = v
	return s
}

type AbortSubmissionRequest struct {
	// 工作空间名称
	Workspace *string `json:"Workspace,omitempty" xml:"Workspace,omitempty"`
	// 投递ID
	SubmissionId *string `json:"SubmissionId,omitempty" xml:"SubmissionId,omitempty"`
}

func (s AbortSubmissionRequest) String() string {
	return tea.Prettify(s)
}

func (s AbortSubmissionRequest) GoString() string {
	return s.String()
}

func (s *AbortSubmissionRequest) SetWorkspace(v string) *AbortSubmissionRequest {
	s.Workspace = &v
	return s
}

func (s *AbortSubmissionRequest) SetSubmissionId(v string) *AbortSubmissionRequest {
	s.SubmissionId = &v
	return s
}

type AbortSubmissionResponseBody struct {
	// 主机ID
	HostId *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	// 请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AbortSubmissionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AbortSubmissionResponseBody) GoString() string {
	return s.String()
}

func (s *AbortSubmissionResponseBody) SetHostId(v string) *AbortSubmissionResponseBody {
	s.HostId = &v
	return s
}

func (s *AbortSubmissionResponseBody) SetRequestId(v string) *AbortSubmissionResponseBody {
	s.RequestId = &v
	return s
}

type AbortSubmissionResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AbortSubmissionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AbortSubmissionResponse) String() string {
	return tea.Prettify(s)
}

func (s AbortSubmissionResponse) GoString() string {
	return s.String()
}

func (s *AbortSubmissionResponse) SetHeaders(v map[string]*string) *AbortSubmissionResponse {
	s.Headers = v
	return s
}

func (s *AbortSubmissionResponse) SetBody(v *AbortSubmissionResponseBody) *AbortSubmissionResponse {
	s.Body = v
	return s
}

type CreateAppRequest struct {
	// 工作空间名称
	Workspace *string `json:"Workspace,omitempty" xml:"Workspace,omitempty"`
	// 应用名称
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// 应用描述
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 应用定义
	Definition *string `json:"Definition,omitempty" xml:"Definition,omitempty"`
	// 应用描述语言
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// 应用描述语语言版本
	LanguageVersion *string `json:"LanguageVersion,omitempty" xml:"LanguageVersion,omitempty"`
	// 应用类型
	AppType *string `json:"AppType,omitempty" xml:"AppType,omitempty"`
	// 应用使用文档
	Documentation *string `json:"Documentation,omitempty" xml:"Documentation,omitempty"`
	// 应用当前版本说明
	RevisionComment *string `json:"RevisionComment,omitempty" xml:"RevisionComment,omitempty"`
	// 应用标签
	Labels *string `json:"Labels,omitempty" xml:"Labels,omitempty"`
	// 幂等Token
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// 依赖应用
	Dependencies []*CreateAppRequestDependencies `json:"Dependencies,omitempty" xml:"Dependencies,omitempty" type:"Repeated"`
	// 参考输入
	Configs []*CreateAppRequestConfigs `json:"Configs,omitempty" xml:"Configs,omitempty" type:"Repeated"`
	// 主WDL路径
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
}

func (s CreateAppRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAppRequest) GoString() string {
	return s.String()
}

func (s *CreateAppRequest) SetWorkspace(v string) *CreateAppRequest {
	s.Workspace = &v
	return s
}

func (s *CreateAppRequest) SetAppName(v string) *CreateAppRequest {
	s.AppName = &v
	return s
}

func (s *CreateAppRequest) SetDescription(v string) *CreateAppRequest {
	s.Description = &v
	return s
}

func (s *CreateAppRequest) SetDefinition(v string) *CreateAppRequest {
	s.Definition = &v
	return s
}

func (s *CreateAppRequest) SetLanguage(v string) *CreateAppRequest {
	s.Language = &v
	return s
}

func (s *CreateAppRequest) SetLanguageVersion(v string) *CreateAppRequest {
	s.LanguageVersion = &v
	return s
}

func (s *CreateAppRequest) SetAppType(v string) *CreateAppRequest {
	s.AppType = &v
	return s
}

func (s *CreateAppRequest) SetDocumentation(v string) *CreateAppRequest {
	s.Documentation = &v
	return s
}

func (s *CreateAppRequest) SetRevisionComment(v string) *CreateAppRequest {
	s.RevisionComment = &v
	return s
}

func (s *CreateAppRequest) SetLabels(v string) *CreateAppRequest {
	s.Labels = &v
	return s
}

func (s *CreateAppRequest) SetClientToken(v string) *CreateAppRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateAppRequest) SetDependencies(v []*CreateAppRequestDependencies) *CreateAppRequest {
	s.Dependencies = v
	return s
}

func (s *CreateAppRequest) SetConfigs(v []*CreateAppRequestConfigs) *CreateAppRequest {
	s.Configs = v
	return s
}

func (s *CreateAppRequest) SetPath(v string) *CreateAppRequest {
	s.Path = &v
	return s
}

type CreateAppRequestDependencies struct {
	// 依赖路径
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// 依赖内容
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
}

func (s CreateAppRequestDependencies) String() string {
	return tea.Prettify(s)
}

func (s CreateAppRequestDependencies) GoString() string {
	return s.String()
}

func (s *CreateAppRequestDependencies) SetPath(v string) *CreateAppRequestDependencies {
	s.Path = &v
	return s
}

func (s *CreateAppRequestDependencies) SetContent(v string) *CreateAppRequestDependencies {
	s.Content = &v
	return s
}

type CreateAppRequestConfigs struct {
	Path    *string `json:"Path,omitempty" xml:"Path,omitempty"`
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
}

func (s CreateAppRequestConfigs) String() string {
	return tea.Prettify(s)
}

func (s CreateAppRequestConfigs) GoString() string {
	return s.String()
}

func (s *CreateAppRequestConfigs) SetPath(v string) *CreateAppRequestConfigs {
	s.Path = &v
	return s
}

func (s *CreateAppRequestConfigs) SetContent(v string) *CreateAppRequestConfigs {
	s.Content = &v
	return s
}

type CreateAppShrinkRequest struct {
	// 工作空间名称
	Workspace *string `json:"Workspace,omitempty" xml:"Workspace,omitempty"`
	// 应用名称
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// 应用描述
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 应用定义
	Definition *string `json:"Definition,omitempty" xml:"Definition,omitempty"`
	// 应用描述语言
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// 应用描述语语言版本
	LanguageVersion *string `json:"LanguageVersion,omitempty" xml:"LanguageVersion,omitempty"`
	// 应用类型
	AppType *string `json:"AppType,omitempty" xml:"AppType,omitempty"`
	// 应用使用文档
	Documentation *string `json:"Documentation,omitempty" xml:"Documentation,omitempty"`
	// 应用当前版本说明
	RevisionComment *string `json:"RevisionComment,omitempty" xml:"RevisionComment,omitempty"`
	// 应用标签
	Labels *string `json:"Labels,omitempty" xml:"Labels,omitempty"`
	// 幂等Token
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// 依赖应用
	DependenciesShrink *string `json:"Dependencies,omitempty" xml:"Dependencies,omitempty"`
	// 参考输入
	ConfigsShrink *string `json:"Configs,omitempty" xml:"Configs,omitempty"`
	// 主WDL路径
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
}

func (s CreateAppShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAppShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateAppShrinkRequest) SetWorkspace(v string) *CreateAppShrinkRequest {
	s.Workspace = &v
	return s
}

func (s *CreateAppShrinkRequest) SetAppName(v string) *CreateAppShrinkRequest {
	s.AppName = &v
	return s
}

func (s *CreateAppShrinkRequest) SetDescription(v string) *CreateAppShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreateAppShrinkRequest) SetDefinition(v string) *CreateAppShrinkRequest {
	s.Definition = &v
	return s
}

func (s *CreateAppShrinkRequest) SetLanguage(v string) *CreateAppShrinkRequest {
	s.Language = &v
	return s
}

func (s *CreateAppShrinkRequest) SetLanguageVersion(v string) *CreateAppShrinkRequest {
	s.LanguageVersion = &v
	return s
}

func (s *CreateAppShrinkRequest) SetAppType(v string) *CreateAppShrinkRequest {
	s.AppType = &v
	return s
}

func (s *CreateAppShrinkRequest) SetDocumentation(v string) *CreateAppShrinkRequest {
	s.Documentation = &v
	return s
}

func (s *CreateAppShrinkRequest) SetRevisionComment(v string) *CreateAppShrinkRequest {
	s.RevisionComment = &v
	return s
}

func (s *CreateAppShrinkRequest) SetLabels(v string) *CreateAppShrinkRequest {
	s.Labels = &v
	return s
}

func (s *CreateAppShrinkRequest) SetClientToken(v string) *CreateAppShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateAppShrinkRequest) SetDependenciesShrink(v string) *CreateAppShrinkRequest {
	s.DependenciesShrink = &v
	return s
}

func (s *CreateAppShrinkRequest) SetConfigsShrink(v string) *CreateAppShrinkRequest {
	s.ConfigsShrink = &v
	return s
}

func (s *CreateAppShrinkRequest) SetPath(v string) *CreateAppShrinkRequest {
	s.Path = &v
	return s
}

type CreateAppResponseBody struct {
	// 工作空间
	Workspace *string `json:"Workspace,omitempty" xml:"Workspace,omitempty"`
	// 应用名称
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// 主机 ID
	HostId *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	// 请求 ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 应用版本号
	Revision *string `json:"Revision,omitempty" xml:"Revision,omitempty"`
}

func (s CreateAppResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAppResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAppResponseBody) SetWorkspace(v string) *CreateAppResponseBody {
	s.Workspace = &v
	return s
}

func (s *CreateAppResponseBody) SetAppName(v string) *CreateAppResponseBody {
	s.AppName = &v
	return s
}

func (s *CreateAppResponseBody) SetHostId(v string) *CreateAppResponseBody {
	s.HostId = &v
	return s
}

func (s *CreateAppResponseBody) SetRequestId(v string) *CreateAppResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAppResponseBody) SetRevision(v string) *CreateAppResponseBody {
	s.Revision = &v
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

type CreateEntityRequest struct {
	Workspace   *string                           `json:"Workspace,omitempty" xml:"Workspace,omitempty"`
	EntityType  *string                           `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	EntityItems []*CreateEntityRequestEntityItems `json:"EntityItems,omitempty" xml:"EntityItems,omitempty" type:"Repeated"`
	ClientToken *string                           `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s CreateEntityRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateEntityRequest) GoString() string {
	return s.String()
}

func (s *CreateEntityRequest) SetWorkspace(v string) *CreateEntityRequest {
	s.Workspace = &v
	return s
}

func (s *CreateEntityRequest) SetEntityType(v string) *CreateEntityRequest {
	s.EntityType = &v
	return s
}

func (s *CreateEntityRequest) SetEntityItems(v []*CreateEntityRequestEntityItems) *CreateEntityRequest {
	s.EntityItems = v
	return s
}

func (s *CreateEntityRequest) SetClientToken(v string) *CreateEntityRequest {
	s.ClientToken = &v
	return s
}

type CreateEntityRequestEntityItems struct {
	EntityName *string            `json:"EntityName,omitempty" xml:"EntityName,omitempty"`
	EntityData map[string]*string `json:"EntityData,omitempty" xml:"EntityData,omitempty"`
}

func (s CreateEntityRequestEntityItems) String() string {
	return tea.Prettify(s)
}

func (s CreateEntityRequestEntityItems) GoString() string {
	return s.String()
}

func (s *CreateEntityRequestEntityItems) SetEntityName(v string) *CreateEntityRequestEntityItems {
	s.EntityName = &v
	return s
}

func (s *CreateEntityRequestEntityItems) SetEntityData(v map[string]*string) *CreateEntityRequestEntityItems {
	s.EntityData = v
	return s
}

type CreateEntityShrinkRequest struct {
	Workspace         *string `json:"Workspace,omitempty" xml:"Workspace,omitempty"`
	EntityType        *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	EntityItemsShrink *string `json:"EntityItems,omitempty" xml:"EntityItems,omitempty"`
	ClientToken       *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s CreateEntityShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateEntityShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateEntityShrinkRequest) SetWorkspace(v string) *CreateEntityShrinkRequest {
	s.Workspace = &v
	return s
}

func (s *CreateEntityShrinkRequest) SetEntityType(v string) *CreateEntityShrinkRequest {
	s.EntityType = &v
	return s
}

func (s *CreateEntityShrinkRequest) SetEntityItemsShrink(v string) *CreateEntityShrinkRequest {
	s.EntityItemsShrink = &v
	return s
}

func (s *CreateEntityShrinkRequest) SetClientToken(v string) *CreateEntityShrinkRequest {
	s.ClientToken = &v
	return s
}

type CreateEntityResponseBody struct {
	// 主机ID
	HostId *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	// 请求ID
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Workspace  *string `json:"Workspace,omitempty" xml:"Workspace,omitempty"`
	EntityType *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
}

func (s CreateEntityResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateEntityResponseBody) GoString() string {
	return s.String()
}

func (s *CreateEntityResponseBody) SetHostId(v string) *CreateEntityResponseBody {
	s.HostId = &v
	return s
}

func (s *CreateEntityResponseBody) SetRequestId(v string) *CreateEntityResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateEntityResponseBody) SetWorkspace(v string) *CreateEntityResponseBody {
	s.Workspace = &v
	return s
}

func (s *CreateEntityResponseBody) SetEntityType(v string) *CreateEntityResponseBody {
	s.EntityType = &v
	return s
}

type CreateEntityResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateEntityResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateEntityResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateEntityResponse) GoString() string {
	return s.String()
}

func (s *CreateEntityResponse) SetHeaders(v map[string]*string) *CreateEntityResponse {
	s.Headers = v
	return s
}

func (s *CreateEntityResponse) SetBody(v *CreateEntityResponseBody) *CreateEntityResponse {
	s.Body = v
	return s
}

type CreateRunRequest struct {
	// 工作空间名字
	Workspace *string `json:"Workspace,omitempty" xml:"Workspace,omitempty"`
	// 任务名称
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 应用名称
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// 应用版本号
	Revision *string `json:"Revision,omitempty" xml:"Revision,omitempty"`
	// 任务描述
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 任务标签
	Labels *string `json:"Labels,omitempty" xml:"Labels,omitempty"`
	// 任务配置
	ExecuteOptions *string `json:"ExecuteOptions,omitempty" xml:"ExecuteOptions,omitempty"`
	// 任务执行目录
	ExecuteDirectory *string `json:"ExecuteDirectory,omitempty" xml:"ExecuteDirectory,omitempty"`
	// 任务输出拷贝目录
	OutputFolder *string `json:"OutputFolder,omitempty" xml:"OutputFolder,omitempty"`
	// 默认运行时
	DefaultRuntime *string `json:"DefaultRuntime,omitempty" xml:"DefaultRuntime,omitempty"`
	// 任务输入
	Inputs *string `json:"Inputs,omitempty" xml:"Inputs,omitempty"`
	// 任务幂等token
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s CreateRunRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateRunRequest) GoString() string {
	return s.String()
}

func (s *CreateRunRequest) SetWorkspace(v string) *CreateRunRequest {
	s.Workspace = &v
	return s
}

func (s *CreateRunRequest) SetName(v string) *CreateRunRequest {
	s.Name = &v
	return s
}

func (s *CreateRunRequest) SetAppName(v string) *CreateRunRequest {
	s.AppName = &v
	return s
}

func (s *CreateRunRequest) SetRevision(v string) *CreateRunRequest {
	s.Revision = &v
	return s
}

func (s *CreateRunRequest) SetDescription(v string) *CreateRunRequest {
	s.Description = &v
	return s
}

func (s *CreateRunRequest) SetLabels(v string) *CreateRunRequest {
	s.Labels = &v
	return s
}

func (s *CreateRunRequest) SetExecuteOptions(v string) *CreateRunRequest {
	s.ExecuteOptions = &v
	return s
}

func (s *CreateRunRequest) SetExecuteDirectory(v string) *CreateRunRequest {
	s.ExecuteDirectory = &v
	return s
}

func (s *CreateRunRequest) SetOutputFolder(v string) *CreateRunRequest {
	s.OutputFolder = &v
	return s
}

func (s *CreateRunRequest) SetDefaultRuntime(v string) *CreateRunRequest {
	s.DefaultRuntime = &v
	return s
}

func (s *CreateRunRequest) SetInputs(v string) *CreateRunRequest {
	s.Inputs = &v
	return s
}

func (s *CreateRunRequest) SetClientToken(v string) *CreateRunRequest {
	s.ClientToken = &v
	return s
}

type CreateRunResponseBody struct {
	// 请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 主机ID
	HostId *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	// 工作空间
	Workspace *string `json:"Workspace,omitempty" xml:"Workspace,omitempty"`
	// 任务ID
	RunId *string `json:"RunId,omitempty" xml:"RunId,omitempty"`
}

func (s CreateRunResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateRunResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRunResponseBody) SetRequestId(v string) *CreateRunResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateRunResponseBody) SetHostId(v string) *CreateRunResponseBody {
	s.HostId = &v
	return s
}

func (s *CreateRunResponseBody) SetWorkspace(v string) *CreateRunResponseBody {
	s.Workspace = &v
	return s
}

func (s *CreateRunResponseBody) SetRunId(v string) *CreateRunResponseBody {
	s.RunId = &v
	return s
}

type CreateRunResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateRunResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateRunResponse) SetBody(v *CreateRunResponseBody) *CreateRunResponse {
	s.Body = v
	return s
}

type CreateSubmissionRequest struct {
	// 工作空间名字
	Workspace *string `json:"Workspace,omitempty" xml:"Workspace,omitempty"`
	// 应用名称
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// 应用版本号
	Revision *string `json:"Revision,omitempty" xml:"Revision,omitempty"`
	// 实体类型
	EntityType  *string   `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	EntityNames []*string `json:"EntityNames,omitempty" xml:"EntityNames,omitempty" type:"Repeated"`
	// 任务配置
	ExecuteOptions *string `json:"ExecuteOptions,omitempty" xml:"ExecuteOptions,omitempty"`
	// 任务执行目录
	ExecuteDirectory *string `json:"ExecuteDirectory,omitempty" xml:"ExecuteDirectory,omitempty"`
	// 任务输出拷贝目录
	OutputFolder *string `json:"OutputFolder,omitempty" xml:"OutputFolder,omitempty"`
	// 默认运行时
	DefaultRuntime *string `json:"DefaultRuntime,omitempty" xml:"DefaultRuntime,omitempty"`
	// 任务输入
	Inputs *string `json:"Inputs,omitempty" xml:"Inputs,omitempty"`
	// 任务输出
	Outputs *string `json:"Outputs,omitempty" xml:"Outputs,omitempty"`
	// 任务幂等token
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s CreateSubmissionRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateSubmissionRequest) GoString() string {
	return s.String()
}

func (s *CreateSubmissionRequest) SetWorkspace(v string) *CreateSubmissionRequest {
	s.Workspace = &v
	return s
}

func (s *CreateSubmissionRequest) SetAppName(v string) *CreateSubmissionRequest {
	s.AppName = &v
	return s
}

func (s *CreateSubmissionRequest) SetRevision(v string) *CreateSubmissionRequest {
	s.Revision = &v
	return s
}

func (s *CreateSubmissionRequest) SetEntityType(v string) *CreateSubmissionRequest {
	s.EntityType = &v
	return s
}

func (s *CreateSubmissionRequest) SetEntityNames(v []*string) *CreateSubmissionRequest {
	s.EntityNames = v
	return s
}

func (s *CreateSubmissionRequest) SetExecuteOptions(v string) *CreateSubmissionRequest {
	s.ExecuteOptions = &v
	return s
}

func (s *CreateSubmissionRequest) SetExecuteDirectory(v string) *CreateSubmissionRequest {
	s.ExecuteDirectory = &v
	return s
}

func (s *CreateSubmissionRequest) SetOutputFolder(v string) *CreateSubmissionRequest {
	s.OutputFolder = &v
	return s
}

func (s *CreateSubmissionRequest) SetDefaultRuntime(v string) *CreateSubmissionRequest {
	s.DefaultRuntime = &v
	return s
}

func (s *CreateSubmissionRequest) SetInputs(v string) *CreateSubmissionRequest {
	s.Inputs = &v
	return s
}

func (s *CreateSubmissionRequest) SetOutputs(v string) *CreateSubmissionRequest {
	s.Outputs = &v
	return s
}

func (s *CreateSubmissionRequest) SetClientToken(v string) *CreateSubmissionRequest {
	s.ClientToken = &v
	return s
}

type CreateSubmissionShrinkRequest struct {
	// 工作空间名字
	Workspace *string `json:"Workspace,omitempty" xml:"Workspace,omitempty"`
	// 应用名称
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// 应用版本号
	Revision *string `json:"Revision,omitempty" xml:"Revision,omitempty"`
	// 实体类型
	EntityType        *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	EntityNamesShrink *string `json:"EntityNames,omitempty" xml:"EntityNames,omitempty"`
	// 任务配置
	ExecuteOptions *string `json:"ExecuteOptions,omitempty" xml:"ExecuteOptions,omitempty"`
	// 任务执行目录
	ExecuteDirectory *string `json:"ExecuteDirectory,omitempty" xml:"ExecuteDirectory,omitempty"`
	// 任务输出拷贝目录
	OutputFolder *string `json:"OutputFolder,omitempty" xml:"OutputFolder,omitempty"`
	// 默认运行时
	DefaultRuntime *string `json:"DefaultRuntime,omitempty" xml:"DefaultRuntime,omitempty"`
	// 任务输入
	Inputs *string `json:"Inputs,omitempty" xml:"Inputs,omitempty"`
	// 任务输出
	Outputs *string `json:"Outputs,omitempty" xml:"Outputs,omitempty"`
	// 任务幂等token
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s CreateSubmissionShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateSubmissionShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateSubmissionShrinkRequest) SetWorkspace(v string) *CreateSubmissionShrinkRequest {
	s.Workspace = &v
	return s
}

func (s *CreateSubmissionShrinkRequest) SetAppName(v string) *CreateSubmissionShrinkRequest {
	s.AppName = &v
	return s
}

func (s *CreateSubmissionShrinkRequest) SetRevision(v string) *CreateSubmissionShrinkRequest {
	s.Revision = &v
	return s
}

func (s *CreateSubmissionShrinkRequest) SetEntityType(v string) *CreateSubmissionShrinkRequest {
	s.EntityType = &v
	return s
}

func (s *CreateSubmissionShrinkRequest) SetEntityNamesShrink(v string) *CreateSubmissionShrinkRequest {
	s.EntityNamesShrink = &v
	return s
}

func (s *CreateSubmissionShrinkRequest) SetExecuteOptions(v string) *CreateSubmissionShrinkRequest {
	s.ExecuteOptions = &v
	return s
}

func (s *CreateSubmissionShrinkRequest) SetExecuteDirectory(v string) *CreateSubmissionShrinkRequest {
	s.ExecuteDirectory = &v
	return s
}

func (s *CreateSubmissionShrinkRequest) SetOutputFolder(v string) *CreateSubmissionShrinkRequest {
	s.OutputFolder = &v
	return s
}

func (s *CreateSubmissionShrinkRequest) SetDefaultRuntime(v string) *CreateSubmissionShrinkRequest {
	s.DefaultRuntime = &v
	return s
}

func (s *CreateSubmissionShrinkRequest) SetInputs(v string) *CreateSubmissionShrinkRequest {
	s.Inputs = &v
	return s
}

func (s *CreateSubmissionShrinkRequest) SetOutputs(v string) *CreateSubmissionShrinkRequest {
	s.Outputs = &v
	return s
}

func (s *CreateSubmissionShrinkRequest) SetClientToken(v string) *CreateSubmissionShrinkRequest {
	s.ClientToken = &v
	return s
}

type CreateSubmissionResponseBody struct {
	// 请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 主机ID
	HostId *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	// 工作空间
	Workspace *string `json:"Workspace,omitempty" xml:"Workspace,omitempty"`
	// 投递ID
	SubmissionId *string `json:"SubmissionId,omitempty" xml:"SubmissionId,omitempty"`
}

func (s CreateSubmissionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateSubmissionResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSubmissionResponseBody) SetRequestId(v string) *CreateSubmissionResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSubmissionResponseBody) SetHostId(v string) *CreateSubmissionResponseBody {
	s.HostId = &v
	return s
}

func (s *CreateSubmissionResponseBody) SetWorkspace(v string) *CreateSubmissionResponseBody {
	s.Workspace = &v
	return s
}

func (s *CreateSubmissionResponseBody) SetSubmissionId(v string) *CreateSubmissionResponseBody {
	s.SubmissionId = &v
	return s
}

type CreateSubmissionResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateSubmissionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateSubmissionResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateSubmissionResponse) GoString() string {
	return s.String()
}

func (s *CreateSubmissionResponse) SetHeaders(v map[string]*string) *CreateSubmissionResponse {
	s.Headers = v
	return s
}

func (s *CreateSubmissionResponse) SetBody(v *CreateSubmissionResponseBody) *CreateSubmissionResponse {
	s.Body = v
	return s
}

type CreateTemplateRequest struct {
	// 工作空间名称
	Workspace *string `json:"Workspace,omitempty" xml:"Workspace,omitempty"`
	// 应用模板名称
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// 应用模板描述
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 应用的名称
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// 应用的版本
	Revision *string `json:"Revision,omitempty" xml:"Revision,omitempty"`
	// 根实体类型
	RootEntity *string `json:"RootEntity,omitempty" xml:"RootEntity,omitempty"`
	// 应用的输入
	Inputs []*CreateTemplateRequestInputs `json:"Inputs,omitempty" xml:"Inputs,omitempty" type:"Repeated"`
	// 应用的输出
	Outputs []*CreateTemplateRequestOutputs `json:"Outputs,omitempty" xml:"Outputs,omitempty" type:"Repeated"`
	// 应用标签
	Labels *string `json:"Labels,omitempty" xml:"Labels,omitempty"`
	// 幂等Token
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s CreateTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateTemplateRequest) GoString() string {
	return s.String()
}

func (s *CreateTemplateRequest) SetWorkspace(v string) *CreateTemplateRequest {
	s.Workspace = &v
	return s
}

func (s *CreateTemplateRequest) SetTemplateName(v string) *CreateTemplateRequest {
	s.TemplateName = &v
	return s
}

func (s *CreateTemplateRequest) SetDescription(v string) *CreateTemplateRequest {
	s.Description = &v
	return s
}

func (s *CreateTemplateRequest) SetAppName(v string) *CreateTemplateRequest {
	s.AppName = &v
	return s
}

func (s *CreateTemplateRequest) SetRevision(v string) *CreateTemplateRequest {
	s.Revision = &v
	return s
}

func (s *CreateTemplateRequest) SetRootEntity(v string) *CreateTemplateRequest {
	s.RootEntity = &v
	return s
}

func (s *CreateTemplateRequest) SetInputs(v []*CreateTemplateRequestInputs) *CreateTemplateRequest {
	s.Inputs = v
	return s
}

func (s *CreateTemplateRequest) SetOutputs(v []*CreateTemplateRequestOutputs) *CreateTemplateRequest {
	s.Outputs = v
	return s
}

func (s *CreateTemplateRequest) SetLabels(v string) *CreateTemplateRequest {
	s.Labels = &v
	return s
}

func (s *CreateTemplateRequest) SetClientToken(v string) *CreateTemplateRequest {
	s.ClientToken = &v
	return s
}

type CreateTemplateRequestInputs struct {
	// 任务名称
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// 变量名
	VariableName *string `json:"VariableName,omitempty" xml:"VariableName,omitempty"`
	// 变量类型
	VariableType *string `json:"VariableType,omitempty" xml:"VariableType,omitempty"`
	// 变量值
	VariableValue *string `json:"VariableValue,omitempty" xml:"VariableValue,omitempty"`
	// 是否必填
	Required *bool `json:"Required,omitempty" xml:"Required,omitempty"`
	// 帮助信息
	Help *string `json:"Help,omitempty" xml:"Help,omitempty"`
	// 步骤顺序
	StepOrder *int32 `json:"StepOrder,omitempty" xml:"StepOrder,omitempty"`
}

func (s CreateTemplateRequestInputs) String() string {
	return tea.Prettify(s)
}

func (s CreateTemplateRequestInputs) GoString() string {
	return s.String()
}

func (s *CreateTemplateRequestInputs) SetTaskName(v string) *CreateTemplateRequestInputs {
	s.TaskName = &v
	return s
}

func (s *CreateTemplateRequestInputs) SetVariableName(v string) *CreateTemplateRequestInputs {
	s.VariableName = &v
	return s
}

func (s *CreateTemplateRequestInputs) SetVariableType(v string) *CreateTemplateRequestInputs {
	s.VariableType = &v
	return s
}

func (s *CreateTemplateRequestInputs) SetVariableValue(v string) *CreateTemplateRequestInputs {
	s.VariableValue = &v
	return s
}

func (s *CreateTemplateRequestInputs) SetRequired(v bool) *CreateTemplateRequestInputs {
	s.Required = &v
	return s
}

func (s *CreateTemplateRequestInputs) SetHelp(v string) *CreateTemplateRequestInputs {
	s.Help = &v
	return s
}

func (s *CreateTemplateRequestInputs) SetStepOrder(v int32) *CreateTemplateRequestInputs {
	s.StepOrder = &v
	return s
}

type CreateTemplateRequestOutputs struct {
	// 任务名称
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// 变量名
	VariableName *string `json:"VariableName,omitempty" xml:"VariableName,omitempty"`
	// 变量类型
	VariableType *string `json:"VariableType,omitempty" xml:"VariableType,omitempty"`
	// 变量值
	VariableValue *string `json:"VariableValue,omitempty" xml:"VariableValue,omitempty"`
	// 是否必填
	Required *bool `json:"Required,omitempty" xml:"Required,omitempty"`
	// 帮助信息
	Help *string `json:"Help,omitempty" xml:"Help,omitempty"`
	// 步骤顺序
	StepOrder *int32 `json:"StepOrder,omitempty" xml:"StepOrder,omitempty"`
}

func (s CreateTemplateRequestOutputs) String() string {
	return tea.Prettify(s)
}

func (s CreateTemplateRequestOutputs) GoString() string {
	return s.String()
}

func (s *CreateTemplateRequestOutputs) SetTaskName(v string) *CreateTemplateRequestOutputs {
	s.TaskName = &v
	return s
}

func (s *CreateTemplateRequestOutputs) SetVariableName(v string) *CreateTemplateRequestOutputs {
	s.VariableName = &v
	return s
}

func (s *CreateTemplateRequestOutputs) SetVariableType(v string) *CreateTemplateRequestOutputs {
	s.VariableType = &v
	return s
}

func (s *CreateTemplateRequestOutputs) SetVariableValue(v string) *CreateTemplateRequestOutputs {
	s.VariableValue = &v
	return s
}

func (s *CreateTemplateRequestOutputs) SetRequired(v bool) *CreateTemplateRequestOutputs {
	s.Required = &v
	return s
}

func (s *CreateTemplateRequestOutputs) SetHelp(v string) *CreateTemplateRequestOutputs {
	s.Help = &v
	return s
}

func (s *CreateTemplateRequestOutputs) SetStepOrder(v int32) *CreateTemplateRequestOutputs {
	s.StepOrder = &v
	return s
}

type CreateTemplateShrinkRequest struct {
	// 工作空间名称
	Workspace *string `json:"Workspace,omitempty" xml:"Workspace,omitempty"`
	// 应用模板名称
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// 应用模板描述
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 应用的名称
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// 应用的版本
	Revision *string `json:"Revision,omitempty" xml:"Revision,omitempty"`
	// 根实体类型
	RootEntity *string `json:"RootEntity,omitempty" xml:"RootEntity,omitempty"`
	// 应用的输入
	InputsShrink *string `json:"Inputs,omitempty" xml:"Inputs,omitempty"`
	// 应用的输出
	OutputsShrink *string `json:"Outputs,omitempty" xml:"Outputs,omitempty"`
	// 应用标签
	Labels *string `json:"Labels,omitempty" xml:"Labels,omitempty"`
	// 幂等Token
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s CreateTemplateShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateTemplateShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateTemplateShrinkRequest) SetWorkspace(v string) *CreateTemplateShrinkRequest {
	s.Workspace = &v
	return s
}

func (s *CreateTemplateShrinkRequest) SetTemplateName(v string) *CreateTemplateShrinkRequest {
	s.TemplateName = &v
	return s
}

func (s *CreateTemplateShrinkRequest) SetDescription(v string) *CreateTemplateShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreateTemplateShrinkRequest) SetAppName(v string) *CreateTemplateShrinkRequest {
	s.AppName = &v
	return s
}

func (s *CreateTemplateShrinkRequest) SetRevision(v string) *CreateTemplateShrinkRequest {
	s.Revision = &v
	return s
}

func (s *CreateTemplateShrinkRequest) SetRootEntity(v string) *CreateTemplateShrinkRequest {
	s.RootEntity = &v
	return s
}

func (s *CreateTemplateShrinkRequest) SetInputsShrink(v string) *CreateTemplateShrinkRequest {
	s.InputsShrink = &v
	return s
}

func (s *CreateTemplateShrinkRequest) SetOutputsShrink(v string) *CreateTemplateShrinkRequest {
	s.OutputsShrink = &v
	return s
}

func (s *CreateTemplateShrinkRequest) SetLabels(v string) *CreateTemplateShrinkRequest {
	s.Labels = &v
	return s
}

func (s *CreateTemplateShrinkRequest) SetClientToken(v string) *CreateTemplateShrinkRequest {
	s.ClientToken = &v
	return s
}

type CreateTemplateResponseBody struct {
	// 工作空间名称
	Workspace *string `json:"Workspace,omitempty" xml:"Workspace,omitempty"`
	// 应用模板名称
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// 主机 ID
	HostId *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	// 请求 ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTemplateResponseBody) SetWorkspace(v string) *CreateTemplateResponseBody {
	s.Workspace = &v
	return s
}

func (s *CreateTemplateResponseBody) SetTemplateName(v string) *CreateTemplateResponseBody {
	s.TemplateName = &v
	return s
}

func (s *CreateTemplateResponseBody) SetHostId(v string) *CreateTemplateResponseBody {
	s.HostId = &v
	return s
}

func (s *CreateTemplateResponseBody) SetRequestId(v string) *CreateTemplateResponseBody {
	s.RequestId = &v
	return s
}

type CreateTemplateResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateTemplateResponse) GoString() string {
	return s.String()
}

func (s *CreateTemplateResponse) SetHeaders(v map[string]*string) *CreateTemplateResponse {
	s.Headers = v
	return s
}

func (s *CreateTemplateResponse) SetBody(v *CreateTemplateResponseBody) *CreateTemplateResponse {
	s.Body = v
	return s
}

type CreateWorkspaceRequest struct {
	// 工作空间名称
	Workspace *string `json:"Workspace,omitempty" xml:"Workspace,omitempty"`
	// 工作空间描述
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 幂等Token
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// 工作空间的OSS工作路径
	OssRoot *string `json:"OssRoot,omitempty" xml:"OssRoot,omitempty"`
	// 工作空间任务生命周期
	JobLifecycle *int32 `json:"JobLifecycle,omitempty" xml:"JobLifecycle,omitempty"`
	// 工作空间内的ram角色
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// 工作空间标签
	Labels *string `json:"Labels,omitempty" xml:"Labels,omitempty"`
}

func (s CreateWorkspaceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateWorkspaceRequest) GoString() string {
	return s.String()
}

func (s *CreateWorkspaceRequest) SetWorkspace(v string) *CreateWorkspaceRequest {
	s.Workspace = &v
	return s
}

func (s *CreateWorkspaceRequest) SetDescription(v string) *CreateWorkspaceRequest {
	s.Description = &v
	return s
}

func (s *CreateWorkspaceRequest) SetClientToken(v string) *CreateWorkspaceRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateWorkspaceRequest) SetOssRoot(v string) *CreateWorkspaceRequest {
	s.OssRoot = &v
	return s
}

func (s *CreateWorkspaceRequest) SetJobLifecycle(v int32) *CreateWorkspaceRequest {
	s.JobLifecycle = &v
	return s
}

func (s *CreateWorkspaceRequest) SetRole(v string) *CreateWorkspaceRequest {
	s.Role = &v
	return s
}

func (s *CreateWorkspaceRequest) SetLabels(v string) *CreateWorkspaceRequest {
	s.Labels = &v
	return s
}

type CreateWorkspaceResponseBody struct {
	// 主机ID
	HostId *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	// 请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 创建成功的工作空间名称
	Workspace *string `json:"Workspace,omitempty" xml:"Workspace,omitempty"`
}

func (s CreateWorkspaceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateWorkspaceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateWorkspaceResponseBody) SetHostId(v string) *CreateWorkspaceResponseBody {
	s.HostId = &v
	return s
}

func (s *CreateWorkspaceResponseBody) SetRequestId(v string) *CreateWorkspaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateWorkspaceResponseBody) SetWorkspace(v string) *CreateWorkspaceResponseBody {
	s.Workspace = &v
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

type DeleteAppRequest struct {
	// 工作空间名称
	Workspace *string `json:"Workspace,omitempty" xml:"Workspace,omitempty"`
	// 应用名称
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// 应用版本
	Revision *string `json:"Revision,omitempty" xml:"Revision,omitempty"`
}

func (s DeleteAppRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteAppRequest) GoString() string {
	return s.String()
}

func (s *DeleteAppRequest) SetWorkspace(v string) *DeleteAppRequest {
	s.Workspace = &v
	return s
}

func (s *DeleteAppRequest) SetAppName(v string) *DeleteAppRequest {
	s.AppName = &v
	return s
}

func (s *DeleteAppRequest) SetRevision(v string) *DeleteAppRequest {
	s.Revision = &v
	return s
}

type DeleteAppResponseBody struct {
	// 主机ID
	HostId *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	// 请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteAppResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteAppResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAppResponseBody) SetHostId(v string) *DeleteAppResponseBody {
	s.HostId = &v
	return s
}

func (s *DeleteAppResponseBody) SetRequestId(v string) *DeleteAppResponseBody {
	s.RequestId = &v
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

type DeleteEntityItemsRequest struct {
	Workspace   *string   `json:"Workspace,omitempty" xml:"Workspace,omitempty"`
	EntityType  *string   `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	EntityNames []*string `json:"EntityNames,omitempty" xml:"EntityNames,omitempty" type:"Repeated"`
}

func (s DeleteEntityItemsRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteEntityItemsRequest) GoString() string {
	return s.String()
}

func (s *DeleteEntityItemsRequest) SetWorkspace(v string) *DeleteEntityItemsRequest {
	s.Workspace = &v
	return s
}

func (s *DeleteEntityItemsRequest) SetEntityType(v string) *DeleteEntityItemsRequest {
	s.EntityType = &v
	return s
}

func (s *DeleteEntityItemsRequest) SetEntityNames(v []*string) *DeleteEntityItemsRequest {
	s.EntityNames = v
	return s
}

type DeleteEntityItemsShrinkRequest struct {
	Workspace         *string `json:"Workspace,omitempty" xml:"Workspace,omitempty"`
	EntityType        *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	EntityNamesShrink *string `json:"EntityNames,omitempty" xml:"EntityNames,omitempty"`
}

func (s DeleteEntityItemsShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteEntityItemsShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteEntityItemsShrinkRequest) SetWorkspace(v string) *DeleteEntityItemsShrinkRequest {
	s.Workspace = &v
	return s
}

func (s *DeleteEntityItemsShrinkRequest) SetEntityType(v string) *DeleteEntityItemsShrinkRequest {
	s.EntityType = &v
	return s
}

func (s *DeleteEntityItemsShrinkRequest) SetEntityNamesShrink(v string) *DeleteEntityItemsShrinkRequest {
	s.EntityNamesShrink = &v
	return s
}

type DeleteEntityItemsResponseBody struct {
	HostId    *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteEntityItemsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteEntityItemsResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteEntityItemsResponseBody) SetHostId(v string) *DeleteEntityItemsResponseBody {
	s.HostId = &v
	return s
}

func (s *DeleteEntityItemsResponseBody) SetRequestId(v string) *DeleteEntityItemsResponseBody {
	s.RequestId = &v
	return s
}

type DeleteEntityItemsResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteEntityItemsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteEntityItemsResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteEntityItemsResponse) GoString() string {
	return s.String()
}

func (s *DeleteEntityItemsResponse) SetHeaders(v map[string]*string) *DeleteEntityItemsResponse {
	s.Headers = v
	return s
}

func (s *DeleteEntityItemsResponse) SetBody(v *DeleteEntityItemsResponseBody) *DeleteEntityItemsResponse {
	s.Body = v
	return s
}

type DeleteRunRequest struct {
	// 工作空间名称
	Workspace *string `json:"Workspace,omitempty" xml:"Workspace,omitempty"`
	// 任务ID
	RunId *string `json:"RunId,omitempty" xml:"RunId,omitempty"`
}

func (s DeleteRunRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteRunRequest) GoString() string {
	return s.String()
}

func (s *DeleteRunRequest) SetWorkspace(v string) *DeleteRunRequest {
	s.Workspace = &v
	return s
}

func (s *DeleteRunRequest) SetRunId(v string) *DeleteRunRequest {
	s.RunId = &v
	return s
}

type DeleteRunResponseBody struct {
	// 主机ID
	HostId *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	// 请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteRunResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteRunResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRunResponseBody) SetHostId(v string) *DeleteRunResponseBody {
	s.HostId = &v
	return s
}

func (s *DeleteRunResponseBody) SetRequestId(v string) *DeleteRunResponseBody {
	s.RequestId = &v
	return s
}

type DeleteRunResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteRunResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DeleteRunResponse) SetBody(v *DeleteRunResponseBody) *DeleteRunResponse {
	s.Body = v
	return s
}

type DeleteSubmissionRequest struct {
	// 工作空间名称
	Workspace *string `json:"Workspace,omitempty" xml:"Workspace,omitempty"`
	// 投递ID
	SubmissionId *string `json:"SubmissionId,omitempty" xml:"SubmissionId,omitempty"`
}

func (s DeleteSubmissionRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteSubmissionRequest) GoString() string {
	return s.String()
}

func (s *DeleteSubmissionRequest) SetWorkspace(v string) *DeleteSubmissionRequest {
	s.Workspace = &v
	return s
}

func (s *DeleteSubmissionRequest) SetSubmissionId(v string) *DeleteSubmissionRequest {
	s.SubmissionId = &v
	return s
}

type DeleteSubmissionResponseBody struct {
	// 主机ID
	HostId *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	// 请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteSubmissionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteSubmissionResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSubmissionResponseBody) SetHostId(v string) *DeleteSubmissionResponseBody {
	s.HostId = &v
	return s
}

func (s *DeleteSubmissionResponseBody) SetRequestId(v string) *DeleteSubmissionResponseBody {
	s.RequestId = &v
	return s
}

type DeleteSubmissionResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteSubmissionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteSubmissionResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteSubmissionResponse) GoString() string {
	return s.String()
}

func (s *DeleteSubmissionResponse) SetHeaders(v map[string]*string) *DeleteSubmissionResponse {
	s.Headers = v
	return s
}

func (s *DeleteSubmissionResponse) SetBody(v *DeleteSubmissionResponseBody) *DeleteSubmissionResponse {
	s.Body = v
	return s
}

type DeleteTemplateRequest struct {
	// 工作空间名称
	Workspace *string `json:"Workspace,omitempty" xml:"Workspace,omitempty"`
	// 应用名称
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
}

func (s DeleteTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteTemplateRequest) GoString() string {
	return s.String()
}

func (s *DeleteTemplateRequest) SetWorkspace(v string) *DeleteTemplateRequest {
	s.Workspace = &v
	return s
}

func (s *DeleteTemplateRequest) SetTemplateName(v string) *DeleteTemplateRequest {
	s.TemplateName = &v
	return s
}

type DeleteTemplateResponseBody struct {
	// 主机ID
	HostId *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	// 请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteTemplateResponseBody) SetHostId(v string) *DeleteTemplateResponseBody {
	s.HostId = &v
	return s
}

func (s *DeleteTemplateResponseBody) SetRequestId(v string) *DeleteTemplateResponseBody {
	s.RequestId = &v
	return s
}

type DeleteTemplateResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteTemplateResponse) GoString() string {
	return s.String()
}

func (s *DeleteTemplateResponse) SetHeaders(v map[string]*string) *DeleteTemplateResponse {
	s.Headers = v
	return s
}

func (s *DeleteTemplateResponse) SetBody(v *DeleteTemplateResponseBody) *DeleteTemplateResponse {
	s.Body = v
	return s
}

type DeleteWorkspaceRequest struct {
	// 工作空间名称
	Workspace *string `json:"Workspace,omitempty" xml:"Workspace,omitempty"`
}

func (s DeleteWorkspaceRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteWorkspaceRequest) GoString() string {
	return s.String()
}

func (s *DeleteWorkspaceRequest) SetWorkspace(v string) *DeleteWorkspaceRequest {
	s.Workspace = &v
	return s
}

type DeleteWorkspaceResponseBody struct {
	// 主机ID
	HostId *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	// 请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteWorkspaceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteWorkspaceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteWorkspaceResponseBody) SetHostId(v string) *DeleteWorkspaceResponseBody {
	s.HostId = &v
	return s
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

type DownloadEntityRequest struct {
	Workspace   *string   `json:"Workspace,omitempty" xml:"Workspace,omitempty"`
	EntityType  *string   `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	EntityNames []*string `json:"EntityNames,omitempty" xml:"EntityNames,omitempty" type:"Repeated"`
}

func (s DownloadEntityRequest) String() string {
	return tea.Prettify(s)
}

func (s DownloadEntityRequest) GoString() string {
	return s.String()
}

func (s *DownloadEntityRequest) SetWorkspace(v string) *DownloadEntityRequest {
	s.Workspace = &v
	return s
}

func (s *DownloadEntityRequest) SetEntityType(v string) *DownloadEntityRequest {
	s.EntityType = &v
	return s
}

func (s *DownloadEntityRequest) SetEntityNames(v []*string) *DownloadEntityRequest {
	s.EntityNames = v
	return s
}

type DownloadEntityShrinkRequest struct {
	Workspace         *string `json:"Workspace,omitempty" xml:"Workspace,omitempty"`
	EntityType        *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	EntityNamesShrink *string `json:"EntityNames,omitempty" xml:"EntityNames,omitempty"`
}

func (s DownloadEntityShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s DownloadEntityShrinkRequest) GoString() string {
	return s.String()
}

func (s *DownloadEntityShrinkRequest) SetWorkspace(v string) *DownloadEntityShrinkRequest {
	s.Workspace = &v
	return s
}

func (s *DownloadEntityShrinkRequest) SetEntityType(v string) *DownloadEntityShrinkRequest {
	s.EntityType = &v
	return s
}

func (s *DownloadEntityShrinkRequest) SetEntityNamesShrink(v string) *DownloadEntityShrinkRequest {
	s.EntityNamesShrink = &v
	return s
}

type DownloadEntityResponseBody struct {
	// 主机ID
	HostId *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	// 请求ID
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	EntityTSVFile *string `json:"EntityTSVFile,omitempty" xml:"EntityTSVFile,omitempty"`
}

func (s DownloadEntityResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DownloadEntityResponseBody) GoString() string {
	return s.String()
}

func (s *DownloadEntityResponseBody) SetHostId(v string) *DownloadEntityResponseBody {
	s.HostId = &v
	return s
}

func (s *DownloadEntityResponseBody) SetRequestId(v string) *DownloadEntityResponseBody {
	s.RequestId = &v
	return s
}

func (s *DownloadEntityResponseBody) SetEntityTSVFile(v string) *DownloadEntityResponseBody {
	s.EntityTSVFile = &v
	return s
}

type DownloadEntityResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DownloadEntityResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DownloadEntityResponse) String() string {
	return tea.Prettify(s)
}

func (s DownloadEntityResponse) GoString() string {
	return s.String()
}

func (s *DownloadEntityResponse) SetHeaders(v map[string]*string) *DownloadEntityResponse {
	s.Headers = v
	return s
}

func (s *DownloadEntityResponse) SetBody(v *DownloadEntityResponseBody) *DownloadEntityResponse {
	s.Body = v
	return s
}

type GetAppRequest struct {
	// 工作空间名称
	Workspace *string `json:"Workspace,omitempty" xml:"Workspace,omitempty"`
	// 应用名称
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// 应用版本号
	Revision *string `json:"Revision,omitempty" xml:"Revision,omitempty"`
}

func (s GetAppRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAppRequest) GoString() string {
	return s.String()
}

func (s *GetAppRequest) SetWorkspace(v string) *GetAppRequest {
	s.Workspace = &v
	return s
}

func (s *GetAppRequest) SetAppName(v string) *GetAppRequest {
	s.AppName = &v
	return s
}

func (s *GetAppRequest) SetRevision(v string) *GetAppRequest {
	s.Revision = &v
	return s
}

type GetAppResponseBody struct {
	// 工作流名称
	WorkflowName *string `json:"WorkflowName,omitempty" xml:"WorkflowName,omitempty"`
	// 应用名称
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 应用定义
	Definition *string `json:"Definition,omitempty" xml:"Definition,omitempty"`
	// 应用简要描述
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 应用详细文档
	Documentation *string `json:"Documentation,omitempty" xml:"Documentation,omitempty"`
	// 主机ID
	HostId *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	// 应用输入
	Inputs []*GetAppResponseBodyInputs `json:"Inputs,omitempty" xml:"Inputs,omitempty" type:"Repeated"`
	// 应用标签
	Labels map[string]*string `json:"Labels,omitempty" xml:"Labels,omitempty"`
	// 应用描述语言
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// 应用描述语言版本
	LanguageVersion *string `json:"LanguageVersion,omitempty" xml:"LanguageVersion,omitempty"`
	// 应用最后修改时间
	LastModifiedTime *string `json:"LastModifiedTime,omitempty" xml:"LastModifiedTime,omitempty"`
	// 应用的输出参数
	Outputs []*GetAppResponseBodyOutputs `json:"Outputs,omitempty" xml:"Outputs,omitempty" type:"Repeated"`
	// 请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 应用版本号
	Revision *string `json:"Revision,omitempty" xml:"Revision,omitempty"`
	// 应用当前版本修改
	RevisionComment *string `json:"RevisionComment,omitempty" xml:"RevisionComment,omitempty"`
	// 应用的所有版本号
	Revisions []*GetAppResponseBodyRevisions `json:"Revisions,omitempty" xml:"Revisions,omitempty" type:"Repeated"`
	// 应用可见范围
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
	// 应用URL
	URL *string `json:"URL,omitempty" xml:"URL,omitempty"`
	// 工作空间名称
	Workspace *string `json:"Workspace,omitempty" xml:"Workspace,omitempty"`
	// 应用来源
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// 命名空间
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// 应用原名
	AppOrigName *string `json:"AppOrigName,omitempty" xml:"AppOrigName,omitempty"`
	// 实体类型
	AppType *string `json:"AppType,omitempty" xml:"AppType,omitempty"`
	// 依赖应用
	Dependencies []*GetAppResponseBodyDependencies `json:"Dependencies,omitempty" xml:"Dependencies,omitempty" type:"Repeated"`
	// 参考输入
	Configs []*GetAppResponseBodyConfigs `json:"Configs,omitempty" xml:"Configs,omitempty" type:"Repeated"`
	// 主WDL路径
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
}

func (s GetAppResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAppResponseBody) GoString() string {
	return s.String()
}

func (s *GetAppResponseBody) SetWorkflowName(v string) *GetAppResponseBody {
	s.WorkflowName = &v
	return s
}

func (s *GetAppResponseBody) SetAppName(v string) *GetAppResponseBody {
	s.AppName = &v
	return s
}

func (s *GetAppResponseBody) SetCreateTime(v string) *GetAppResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetAppResponseBody) SetDefinition(v string) *GetAppResponseBody {
	s.Definition = &v
	return s
}

func (s *GetAppResponseBody) SetDescription(v string) *GetAppResponseBody {
	s.Description = &v
	return s
}

func (s *GetAppResponseBody) SetDocumentation(v string) *GetAppResponseBody {
	s.Documentation = &v
	return s
}

func (s *GetAppResponseBody) SetHostId(v string) *GetAppResponseBody {
	s.HostId = &v
	return s
}

func (s *GetAppResponseBody) SetInputs(v []*GetAppResponseBodyInputs) *GetAppResponseBody {
	s.Inputs = v
	return s
}

func (s *GetAppResponseBody) SetLabels(v map[string]*string) *GetAppResponseBody {
	s.Labels = v
	return s
}

func (s *GetAppResponseBody) SetLanguage(v string) *GetAppResponseBody {
	s.Language = &v
	return s
}

func (s *GetAppResponseBody) SetLanguageVersion(v string) *GetAppResponseBody {
	s.LanguageVersion = &v
	return s
}

func (s *GetAppResponseBody) SetLastModifiedTime(v string) *GetAppResponseBody {
	s.LastModifiedTime = &v
	return s
}

func (s *GetAppResponseBody) SetOutputs(v []*GetAppResponseBodyOutputs) *GetAppResponseBody {
	s.Outputs = v
	return s
}

func (s *GetAppResponseBody) SetRequestId(v string) *GetAppResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAppResponseBody) SetRevision(v string) *GetAppResponseBody {
	s.Revision = &v
	return s
}

func (s *GetAppResponseBody) SetRevisionComment(v string) *GetAppResponseBody {
	s.RevisionComment = &v
	return s
}

func (s *GetAppResponseBody) SetRevisions(v []*GetAppResponseBodyRevisions) *GetAppResponseBody {
	s.Revisions = v
	return s
}

func (s *GetAppResponseBody) SetScope(v string) *GetAppResponseBody {
	s.Scope = &v
	return s
}

func (s *GetAppResponseBody) SetURL(v string) *GetAppResponseBody {
	s.URL = &v
	return s
}

func (s *GetAppResponseBody) SetWorkspace(v string) *GetAppResponseBody {
	s.Workspace = &v
	return s
}

func (s *GetAppResponseBody) SetSource(v string) *GetAppResponseBody {
	s.Source = &v
	return s
}

func (s *GetAppResponseBody) SetNamespace(v string) *GetAppResponseBody {
	s.Namespace = &v
	return s
}

func (s *GetAppResponseBody) SetAppOrigName(v string) *GetAppResponseBody {
	s.AppOrigName = &v
	return s
}

func (s *GetAppResponseBody) SetAppType(v string) *GetAppResponseBody {
	s.AppType = &v
	return s
}

func (s *GetAppResponseBody) SetDependencies(v []*GetAppResponseBodyDependencies) *GetAppResponseBody {
	s.Dependencies = v
	return s
}

func (s *GetAppResponseBody) SetConfigs(v []*GetAppResponseBodyConfigs) *GetAppResponseBody {
	s.Configs = v
	return s
}

func (s *GetAppResponseBody) SetPath(v string) *GetAppResponseBody {
	s.Path = &v
	return s
}

type GetAppResponseBodyInputs struct {
	TaskName      *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	VariableName  *string `json:"VariableName,omitempty" xml:"VariableName,omitempty"`
	VariableType  *string `json:"VariableType,omitempty" xml:"VariableType,omitempty"`
	VariableValue *string `json:"VariableValue,omitempty" xml:"VariableValue,omitempty"`
	Required      *bool   `json:"Required,omitempty" xml:"Required,omitempty"`
	Help          *string `json:"Help,omitempty" xml:"Help,omitempty"`
	StepOrder     *int64  `json:"StepOrder,omitempty" xml:"StepOrder,omitempty"`
}

func (s GetAppResponseBodyInputs) String() string {
	return tea.Prettify(s)
}

func (s GetAppResponseBodyInputs) GoString() string {
	return s.String()
}

func (s *GetAppResponseBodyInputs) SetTaskName(v string) *GetAppResponseBodyInputs {
	s.TaskName = &v
	return s
}

func (s *GetAppResponseBodyInputs) SetVariableName(v string) *GetAppResponseBodyInputs {
	s.VariableName = &v
	return s
}

func (s *GetAppResponseBodyInputs) SetVariableType(v string) *GetAppResponseBodyInputs {
	s.VariableType = &v
	return s
}

func (s *GetAppResponseBodyInputs) SetVariableValue(v string) *GetAppResponseBodyInputs {
	s.VariableValue = &v
	return s
}

func (s *GetAppResponseBodyInputs) SetRequired(v bool) *GetAppResponseBodyInputs {
	s.Required = &v
	return s
}

func (s *GetAppResponseBodyInputs) SetHelp(v string) *GetAppResponseBodyInputs {
	s.Help = &v
	return s
}

func (s *GetAppResponseBodyInputs) SetStepOrder(v int64) *GetAppResponseBodyInputs {
	s.StepOrder = &v
	return s
}

type GetAppResponseBodyOutputs struct {
	TaskName      *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	VariableName  *string `json:"VariableName,omitempty" xml:"VariableName,omitempty"`
	VariableType  *string `json:"VariableType,omitempty" xml:"VariableType,omitempty"`
	VariableValue *string `json:"VariableValue,omitempty" xml:"VariableValue,omitempty"`
	Required      *bool   `json:"Required,omitempty" xml:"Required,omitempty"`
	Help          *string `json:"Help,omitempty" xml:"Help,omitempty"`
	StepOrder     *int64  `json:"StepOrder,omitempty" xml:"StepOrder,omitempty"`
}

func (s GetAppResponseBodyOutputs) String() string {
	return tea.Prettify(s)
}

func (s GetAppResponseBodyOutputs) GoString() string {
	return s.String()
}

func (s *GetAppResponseBodyOutputs) SetTaskName(v string) *GetAppResponseBodyOutputs {
	s.TaskName = &v
	return s
}

func (s *GetAppResponseBodyOutputs) SetVariableName(v string) *GetAppResponseBodyOutputs {
	s.VariableName = &v
	return s
}

func (s *GetAppResponseBodyOutputs) SetVariableType(v string) *GetAppResponseBodyOutputs {
	s.VariableType = &v
	return s
}

func (s *GetAppResponseBodyOutputs) SetVariableValue(v string) *GetAppResponseBodyOutputs {
	s.VariableValue = &v
	return s
}

func (s *GetAppResponseBodyOutputs) SetRequired(v bool) *GetAppResponseBodyOutputs {
	s.Required = &v
	return s
}

func (s *GetAppResponseBodyOutputs) SetHelp(v string) *GetAppResponseBodyOutputs {
	s.Help = &v
	return s
}

func (s *GetAppResponseBodyOutputs) SetStepOrder(v int64) *GetAppResponseBodyOutputs {
	s.StepOrder = &v
	return s
}

type GetAppResponseBodyRevisions struct {
	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 版本号
	Revision *string `json:"Revision,omitempty" xml:"Revision,omitempty"`
	// 版本描述
	RevisionComment *string `json:"RevisionComment,omitempty" xml:"RevisionComment,omitempty"`
}

func (s GetAppResponseBodyRevisions) String() string {
	return tea.Prettify(s)
}

func (s GetAppResponseBodyRevisions) GoString() string {
	return s.String()
}

func (s *GetAppResponseBodyRevisions) SetCreateTime(v string) *GetAppResponseBodyRevisions {
	s.CreateTime = &v
	return s
}

func (s *GetAppResponseBodyRevisions) SetRevision(v string) *GetAppResponseBodyRevisions {
	s.Revision = &v
	return s
}

func (s *GetAppResponseBodyRevisions) SetRevisionComment(v string) *GetAppResponseBodyRevisions {
	s.RevisionComment = &v
	return s
}

type GetAppResponseBodyDependencies struct {
	// 依赖路径
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// wdl内容
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
}

func (s GetAppResponseBodyDependencies) String() string {
	return tea.Prettify(s)
}

func (s GetAppResponseBodyDependencies) GoString() string {
	return s.String()
}

func (s *GetAppResponseBodyDependencies) SetPath(v string) *GetAppResponseBodyDependencies {
	s.Path = &v
	return s
}

func (s *GetAppResponseBodyDependencies) SetContent(v string) *GetAppResponseBodyDependencies {
	s.Content = &v
	return s
}

type GetAppResponseBodyConfigs struct {
	Path    *string `json:"Path,omitempty" xml:"Path,omitempty"`
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
}

func (s GetAppResponseBodyConfigs) String() string {
	return tea.Prettify(s)
}

func (s GetAppResponseBodyConfigs) GoString() string {
	return s.String()
}

func (s *GetAppResponseBodyConfigs) SetPath(v string) *GetAppResponseBodyConfigs {
	s.Path = &v
	return s
}

func (s *GetAppResponseBodyConfigs) SetContent(v string) *GetAppResponseBodyConfigs {
	s.Content = &v
	return s
}

type GetAppResponse struct {
	Headers map[string]*string  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetAppResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetAppResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAppResponse) GoString() string {
	return s.String()
}

func (s *GetAppResponse) SetHeaders(v map[string]*string) *GetAppResponse {
	s.Headers = v
	return s
}

func (s *GetAppResponse) SetBody(v *GetAppResponseBody) *GetAppResponse {
	s.Body = v
	return s
}

type GetEntityRequest struct {
	Workspace  *string `json:"Workspace,omitempty" xml:"Workspace,omitempty"`
	EntityType *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
}

func (s GetEntityRequest) String() string {
	return tea.Prettify(s)
}

func (s GetEntityRequest) GoString() string {
	return s.String()
}

func (s *GetEntityRequest) SetWorkspace(v string) *GetEntityRequest {
	s.Workspace = &v
	return s
}

func (s *GetEntityRequest) SetEntityType(v string) *GetEntityRequest {
	s.EntityType = &v
	return s
}

type GetEntityResponseBody struct {
	// 主机ID
	HostId *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	// 请求ID
	RequestId  *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Workspace  *string   `json:"Workspace,omitempty" xml:"Workspace,omitempty"`
	EntityType *string   `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	Attributes []*string `json:"Attributes,omitempty" xml:"Attributes,omitempty" type:"Repeated"`
	TotalCount *int32    `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetEntityResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetEntityResponseBody) GoString() string {
	return s.String()
}

func (s *GetEntityResponseBody) SetHostId(v string) *GetEntityResponseBody {
	s.HostId = &v
	return s
}

func (s *GetEntityResponseBody) SetRequestId(v string) *GetEntityResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetEntityResponseBody) SetWorkspace(v string) *GetEntityResponseBody {
	s.Workspace = &v
	return s
}

func (s *GetEntityResponseBody) SetEntityType(v string) *GetEntityResponseBody {
	s.EntityType = &v
	return s
}

func (s *GetEntityResponseBody) SetAttributes(v []*string) *GetEntityResponseBody {
	s.Attributes = v
	return s
}

func (s *GetEntityResponseBody) SetTotalCount(v int32) *GetEntityResponseBody {
	s.TotalCount = &v
	return s
}

type GetEntityResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetEntityResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetEntityResponse) String() string {
	return tea.Prettify(s)
}

func (s GetEntityResponse) GoString() string {
	return s.String()
}

func (s *GetEntityResponse) SetHeaders(v map[string]*string) *GetEntityResponse {
	s.Headers = v
	return s
}

func (s *GetEntityResponse) SetBody(v *GetEntityResponseBody) *GetEntityResponse {
	s.Body = v
	return s
}

type GetGlobalAppRequest struct {
	// 查询字段信息
	Attributes *string `json:"Attributes,omitempty" xml:"Attributes,omitempty"`
	// 命名空间
	NamespaceName *string `json:"NamespaceName,omitempty" xml:"NamespaceName,omitempty"`
	// 应用名称
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// 应用版本
	AppVersion *string `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
	// 应用可用区域
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s GetGlobalAppRequest) String() string {
	return tea.Prettify(s)
}

func (s GetGlobalAppRequest) GoString() string {
	return s.String()
}

func (s *GetGlobalAppRequest) SetAttributes(v string) *GetGlobalAppRequest {
	s.Attributes = &v
	return s
}

func (s *GetGlobalAppRequest) SetNamespaceName(v string) *GetGlobalAppRequest {
	s.NamespaceName = &v
	return s
}

func (s *GetGlobalAppRequest) SetAppName(v string) *GetGlobalAppRequest {
	s.AppName = &v
	return s
}

func (s *GetGlobalAppRequest) SetAppVersion(v string) *GetGlobalAppRequest {
	s.AppVersion = &v
	return s
}

func (s *GetGlobalAppRequest) SetRegion(v string) *GetGlobalAppRequest {
	s.Region = &v
	return s
}

type GetGlobalAppResponseBody struct {
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Message     *string `json:"Message,omitempty" xml:"Message,omitempty"`
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// 更新时间
	LastModified *string `json:"LastModified,omitempty" xml:"LastModified,omitempty"`
	// 命名空间名称
	NamespaceName *string `json:"NamespaceName,omitempty" xml:"NamespaceName,omitempty"`
	// 应用权限
	AppScope *string `json:"AppScope,omitempty" xml:"AppScope,omitempty"`
	// 应用名称
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// 应用支持的区域
	RegionIds []*string `json:"RegionIds,omitempty" xml:"RegionIds,omitempty" type:"Repeated"`
	// 应用版本
	AppVersion *string `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
	// 应用类型
	AppType *string `json:"AppType,omitempty" xml:"AppType,omitempty"`
	// 应用计费说明
	AppFee *string `json:"AppFee,omitempty" xml:"AppFee,omitempty"`
	// 应用描述
	AppDescription *string `json:"AppDescription,omitempty" xml:"AppDescription,omitempty"`
	// 应用所属分类
	Categories []*string `json:"Categories,omitempty" xml:"Categories,omitempty" type:"Repeated"`
	// 应用所属工具合集
	Toolkit *string `json:"Toolkit,omitempty" xml:"Toolkit,omitempty"`
	// 应用联系人信息
	Contact *string `json:"Contact,omitempty" xml:"Contact,omitempty"`
	// 应用主页信息
	Links []*string `json:"Links,omitempty" xml:"Links,omitempty" type:"Repeated"`
	// 应用的所有版本列表
	AppVersions []*GetGlobalAppResponseBodyAppVersions `json:"AppVersions,omitempty" xml:"AppVersions,omitempty" type:"Repeated"`
	// 默认版本
	AppDefaultVersion *string `json:"AppDefaultVersion,omitempty" xml:"AppDefaultVersion,omitempty"`
	// 应用描述语言标准(WDL/CWL)
	AppDescriptorType *string `json:"AppDescriptorType,omitempty" xml:"AppDescriptorType,omitempty"`
	// 应用的描述文件内容
	AppDescriptorFiles []*GetGlobalAppResponseBodyAppDescriptorFiles `json:"AppDescriptorFiles,omitempty" xml:"AppDescriptorFiles,omitempty" type:"Repeated"`
	// 应用的帮助文档
	Document *string `json:"Document,omitempty" xml:"Document,omitempty"`
	// 应用的备注信息
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// 应用的DAG信息
	DAG            *string `json:"DAG,omitempty" xml:"DAG,omitempty"`
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	HostId         *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	// 应用收藏置顶标记
	Pinned *bool `json:"Pinned,omitempty" xml:"Pinned,omitempty"`
}

func (s GetGlobalAppResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetGlobalAppResponseBody) GoString() string {
	return s.String()
}

func (s *GetGlobalAppResponseBody) SetRequestId(v string) *GetGlobalAppResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetGlobalAppResponseBody) SetMessage(v string) *GetGlobalAppResponseBody {
	s.Message = &v
	return s
}

func (s *GetGlobalAppResponseBody) SetDynamicCode(v string) *GetGlobalAppResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *GetGlobalAppResponseBody) SetLastModified(v string) *GetGlobalAppResponseBody {
	s.LastModified = &v
	return s
}

func (s *GetGlobalAppResponseBody) SetNamespaceName(v string) *GetGlobalAppResponseBody {
	s.NamespaceName = &v
	return s
}

func (s *GetGlobalAppResponseBody) SetAppScope(v string) *GetGlobalAppResponseBody {
	s.AppScope = &v
	return s
}

func (s *GetGlobalAppResponseBody) SetAppName(v string) *GetGlobalAppResponseBody {
	s.AppName = &v
	return s
}

func (s *GetGlobalAppResponseBody) SetRegionIds(v []*string) *GetGlobalAppResponseBody {
	s.RegionIds = v
	return s
}

func (s *GetGlobalAppResponseBody) SetAppVersion(v string) *GetGlobalAppResponseBody {
	s.AppVersion = &v
	return s
}

func (s *GetGlobalAppResponseBody) SetAppType(v string) *GetGlobalAppResponseBody {
	s.AppType = &v
	return s
}

func (s *GetGlobalAppResponseBody) SetAppFee(v string) *GetGlobalAppResponseBody {
	s.AppFee = &v
	return s
}

func (s *GetGlobalAppResponseBody) SetAppDescription(v string) *GetGlobalAppResponseBody {
	s.AppDescription = &v
	return s
}

func (s *GetGlobalAppResponseBody) SetCategories(v []*string) *GetGlobalAppResponseBody {
	s.Categories = v
	return s
}

func (s *GetGlobalAppResponseBody) SetToolkit(v string) *GetGlobalAppResponseBody {
	s.Toolkit = &v
	return s
}

func (s *GetGlobalAppResponseBody) SetContact(v string) *GetGlobalAppResponseBody {
	s.Contact = &v
	return s
}

func (s *GetGlobalAppResponseBody) SetLinks(v []*string) *GetGlobalAppResponseBody {
	s.Links = v
	return s
}

func (s *GetGlobalAppResponseBody) SetAppVersions(v []*GetGlobalAppResponseBodyAppVersions) *GetGlobalAppResponseBody {
	s.AppVersions = v
	return s
}

func (s *GetGlobalAppResponseBody) SetAppDefaultVersion(v string) *GetGlobalAppResponseBody {
	s.AppDefaultVersion = &v
	return s
}

func (s *GetGlobalAppResponseBody) SetAppDescriptorType(v string) *GetGlobalAppResponseBody {
	s.AppDescriptorType = &v
	return s
}

func (s *GetGlobalAppResponseBody) SetAppDescriptorFiles(v []*GetGlobalAppResponseBodyAppDescriptorFiles) *GetGlobalAppResponseBody {
	s.AppDescriptorFiles = v
	return s
}

func (s *GetGlobalAppResponseBody) SetDocument(v string) *GetGlobalAppResponseBody {
	s.Document = &v
	return s
}

func (s *GetGlobalAppResponseBody) SetComment(v string) *GetGlobalAppResponseBody {
	s.Comment = &v
	return s
}

func (s *GetGlobalAppResponseBody) SetDAG(v string) *GetGlobalAppResponseBody {
	s.DAG = &v
	return s
}

func (s *GetGlobalAppResponseBody) SetDynamicMessage(v string) *GetGlobalAppResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *GetGlobalAppResponseBody) SetCode(v string) *GetGlobalAppResponseBody {
	s.Code = &v
	return s
}

func (s *GetGlobalAppResponseBody) SetHostId(v string) *GetGlobalAppResponseBody {
	s.HostId = &v
	return s
}

func (s *GetGlobalAppResponseBody) SetPinned(v bool) *GetGlobalAppResponseBody {
	s.Pinned = &v
	return s
}

type GetGlobalAppResponseBodyAppVersions struct {
	// 应用版本
	AppVersion *string `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
	// 版本描述
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// 更新时间
	LastModified *string `json:"LastModified,omitempty" xml:"LastModified,omitempty"`
}

func (s GetGlobalAppResponseBodyAppVersions) String() string {
	return tea.Prettify(s)
}

func (s GetGlobalAppResponseBodyAppVersions) GoString() string {
	return s.String()
}

func (s *GetGlobalAppResponseBodyAppVersions) SetAppVersion(v string) *GetGlobalAppResponseBodyAppVersions {
	s.AppVersion = &v
	return s
}

func (s *GetGlobalAppResponseBodyAppVersions) SetComment(v string) *GetGlobalAppResponseBodyAppVersions {
	s.Comment = &v
	return s
}

func (s *GetGlobalAppResponseBodyAppVersions) SetLastModified(v string) *GetGlobalAppResponseBodyAppVersions {
	s.LastModified = &v
	return s
}

type GetGlobalAppResponseBodyAppDescriptorFiles struct {
	// 应用文件的类型
	FileType *string `json:"FileType,omitempty" xml:"FileType,omitempty"`
	// 应用文件的路径，除PRIMARY_DESCRIPTOR外，其他均为相对于PRIMARY_DESCRIPTOR的相对路径
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// 应用文件内容
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// 应用文件链接
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
	// 应用文件内容的完整性校验码，如MD5值
	Checksum *string `json:"Checksum,omitempty" xml:"Checksum,omitempty"`
}

func (s GetGlobalAppResponseBodyAppDescriptorFiles) String() string {
	return tea.Prettify(s)
}

func (s GetGlobalAppResponseBodyAppDescriptorFiles) GoString() string {
	return s.String()
}

func (s *GetGlobalAppResponseBodyAppDescriptorFiles) SetFileType(v string) *GetGlobalAppResponseBodyAppDescriptorFiles {
	s.FileType = &v
	return s
}

func (s *GetGlobalAppResponseBodyAppDescriptorFiles) SetPath(v string) *GetGlobalAppResponseBodyAppDescriptorFiles {
	s.Path = &v
	return s
}

func (s *GetGlobalAppResponseBodyAppDescriptorFiles) SetContent(v string) *GetGlobalAppResponseBodyAppDescriptorFiles {
	s.Content = &v
	return s
}

func (s *GetGlobalAppResponseBodyAppDescriptorFiles) SetUrl(v string) *GetGlobalAppResponseBodyAppDescriptorFiles {
	s.Url = &v
	return s
}

func (s *GetGlobalAppResponseBodyAppDescriptorFiles) SetChecksum(v string) *GetGlobalAppResponseBodyAppDescriptorFiles {
	s.Checksum = &v
	return s
}

type GetGlobalAppResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetGlobalAppResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetGlobalAppResponse) String() string {
	return tea.Prettify(s)
}

func (s GetGlobalAppResponse) GoString() string {
	return s.String()
}

func (s *GetGlobalAppResponse) SetHeaders(v map[string]*string) *GetGlobalAppResponse {
	s.Headers = v
	return s
}

func (s *GetGlobalAppResponse) SetBody(v *GetGlobalAppResponseBody) *GetGlobalAppResponse {
	s.Body = v
	return s
}

type GetPublicDatasetRequest struct {
	// 数据集名称
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	// 查询的字段名
	Attributes []*string `json:"Attributes,omitempty" xml:"Attributes,omitempty" type:"Repeated"`
}

func (s GetPublicDatasetRequest) String() string {
	return tea.Prettify(s)
}

func (s GetPublicDatasetRequest) GoString() string {
	return s.String()
}

func (s *GetPublicDatasetRequest) SetDatasetName(v string) *GetPublicDatasetRequest {
	s.DatasetName = &v
	return s
}

func (s *GetPublicDatasetRequest) SetAttributes(v []*string) *GetPublicDatasetRequest {
	s.Attributes = v
	return s
}

type GetPublicDatasetShrinkRequest struct {
	// 数据集名称
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	// 查询的字段名
	AttributesShrink *string `json:"Attributes,omitempty" xml:"Attributes,omitempty"`
}

func (s GetPublicDatasetShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s GetPublicDatasetShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetPublicDatasetShrinkRequest) SetDatasetName(v string) *GetPublicDatasetShrinkRequest {
	s.DatasetName = &v
	return s
}

func (s *GetPublicDatasetShrinkRequest) SetAttributesShrink(v string) *GetPublicDatasetShrinkRequest {
	s.AttributesShrink = &v
	return s
}

type GetPublicDatasetResponseBody struct {
	// 主机ID
	HostId *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	// 请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	HttpCode  *int32  `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	// 错误码
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 错误消息
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 是否调用成功
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// 公共数据集描述
	DatasetDescription *string `json:"DatasetDescription,omitempty" xml:"DatasetDescription,omitempty"`
	// 公共数据集版权信息
	Copyright *string `json:"Copyright,omitempty" xml:"Copyright,omitempty"`
	// 公共数据集访问要求
	AccessRequirements *string `json:"AccessRequirements,omitempty" xml:"AccessRequirements,omitempty"`
	// 关于公共数据集
	About *string `json:"About,omitempty" xml:"About,omitempty"`
	// 公共数据集标签
	Tags []*string `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// 公共数据集名称
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	// 公共数据集更新频率
	UpdateFrequency *string `json:"UpdateFrequency,omitempty" xml:"UpdateFrequency,omitempty"`
	// 公共数据集可用区域
	RegionIds []*string `json:"RegionIds,omitempty" xml:"RegionIds,omitempty" type:"Repeated"`
	// 最后更新时间
	LastModified *string `json:"LastModified,omitempty" xml:"LastModified,omitempty"`
}

func (s GetPublicDatasetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetPublicDatasetResponseBody) GoString() string {
	return s.String()
}

func (s *GetPublicDatasetResponseBody) SetHostId(v string) *GetPublicDatasetResponseBody {
	s.HostId = &v
	return s
}

func (s *GetPublicDatasetResponseBody) SetRequestId(v string) *GetPublicDatasetResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPublicDatasetResponseBody) SetHttpCode(v int32) *GetPublicDatasetResponseBody {
	s.HttpCode = &v
	return s
}

func (s *GetPublicDatasetResponseBody) SetCode(v string) *GetPublicDatasetResponseBody {
	s.Code = &v
	return s
}

func (s *GetPublicDatasetResponseBody) SetMessage(v string) *GetPublicDatasetResponseBody {
	s.Message = &v
	return s
}

func (s *GetPublicDatasetResponseBody) SetSuccess(v bool) *GetPublicDatasetResponseBody {
	s.Success = &v
	return s
}

func (s *GetPublicDatasetResponseBody) SetDatasetDescription(v string) *GetPublicDatasetResponseBody {
	s.DatasetDescription = &v
	return s
}

func (s *GetPublicDatasetResponseBody) SetCopyright(v string) *GetPublicDatasetResponseBody {
	s.Copyright = &v
	return s
}

func (s *GetPublicDatasetResponseBody) SetAccessRequirements(v string) *GetPublicDatasetResponseBody {
	s.AccessRequirements = &v
	return s
}

func (s *GetPublicDatasetResponseBody) SetAbout(v string) *GetPublicDatasetResponseBody {
	s.About = &v
	return s
}

func (s *GetPublicDatasetResponseBody) SetTags(v []*string) *GetPublicDatasetResponseBody {
	s.Tags = v
	return s
}

func (s *GetPublicDatasetResponseBody) SetDatasetName(v string) *GetPublicDatasetResponseBody {
	s.DatasetName = &v
	return s
}

func (s *GetPublicDatasetResponseBody) SetUpdateFrequency(v string) *GetPublicDatasetResponseBody {
	s.UpdateFrequency = &v
	return s
}

func (s *GetPublicDatasetResponseBody) SetRegionIds(v []*string) *GetPublicDatasetResponseBody {
	s.RegionIds = v
	return s
}

func (s *GetPublicDatasetResponseBody) SetLastModified(v string) *GetPublicDatasetResponseBody {
	s.LastModified = &v
	return s
}

type GetPublicDatasetResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetPublicDatasetResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetPublicDatasetResponse) String() string {
	return tea.Prettify(s)
}

func (s GetPublicDatasetResponse) GoString() string {
	return s.String()
}

func (s *GetPublicDatasetResponse) SetHeaders(v map[string]*string) *GetPublicDatasetResponse {
	s.Headers = v
	return s
}

func (s *GetPublicDatasetResponse) SetBody(v *GetPublicDatasetResponseBody) *GetPublicDatasetResponse {
	s.Body = v
	return s
}

type GetPublicDatasetEntityRequest struct {
	// 数据集名称
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	// 实体类型
	EntityType *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
}

func (s GetPublicDatasetEntityRequest) String() string {
	return tea.Prettify(s)
}

func (s GetPublicDatasetEntityRequest) GoString() string {
	return s.String()
}

func (s *GetPublicDatasetEntityRequest) SetDatasetName(v string) *GetPublicDatasetEntityRequest {
	s.DatasetName = &v
	return s
}

func (s *GetPublicDatasetEntityRequest) SetEntityType(v string) *GetPublicDatasetEntityRequest {
	s.EntityType = &v
	return s
}

type GetPublicDatasetEntityResponseBody struct {
	// 主机ID
	HostId *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	// 请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	HttpCode  *int32  `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	// 错误码
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 错误消息
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 公共数据集名称
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	// 实体类型
	EntityType *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	// 实体属性名称列表
	Attributes []*string `json:"Attributes,omitempty" xml:"Attributes,omitempty" type:"Repeated"`
	// 该类型实体总数
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetPublicDatasetEntityResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetPublicDatasetEntityResponseBody) GoString() string {
	return s.String()
}

func (s *GetPublicDatasetEntityResponseBody) SetHostId(v string) *GetPublicDatasetEntityResponseBody {
	s.HostId = &v
	return s
}

func (s *GetPublicDatasetEntityResponseBody) SetRequestId(v string) *GetPublicDatasetEntityResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPublicDatasetEntityResponseBody) SetHttpCode(v int32) *GetPublicDatasetEntityResponseBody {
	s.HttpCode = &v
	return s
}

func (s *GetPublicDatasetEntityResponseBody) SetCode(v string) *GetPublicDatasetEntityResponseBody {
	s.Code = &v
	return s
}

func (s *GetPublicDatasetEntityResponseBody) SetMessage(v string) *GetPublicDatasetEntityResponseBody {
	s.Message = &v
	return s
}

func (s *GetPublicDatasetEntityResponseBody) SetDatasetName(v string) *GetPublicDatasetEntityResponseBody {
	s.DatasetName = &v
	return s
}

func (s *GetPublicDatasetEntityResponseBody) SetEntityType(v string) *GetPublicDatasetEntityResponseBody {
	s.EntityType = &v
	return s
}

func (s *GetPublicDatasetEntityResponseBody) SetAttributes(v []*string) *GetPublicDatasetEntityResponseBody {
	s.Attributes = v
	return s
}

func (s *GetPublicDatasetEntityResponseBody) SetTotalCount(v int32) *GetPublicDatasetEntityResponseBody {
	s.TotalCount = &v
	return s
}

type GetPublicDatasetEntityResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetPublicDatasetEntityResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetPublicDatasetEntityResponse) String() string {
	return tea.Prettify(s)
}

func (s GetPublicDatasetEntityResponse) GoString() string {
	return s.String()
}

func (s *GetPublicDatasetEntityResponse) SetHeaders(v map[string]*string) *GetPublicDatasetEntityResponse {
	s.Headers = v
	return s
}

func (s *GetPublicDatasetEntityResponse) SetBody(v *GetPublicDatasetEntityResponseBody) *GetPublicDatasetEntityResponse {
	s.Body = v
	return s
}

type GetRunRequest struct {
	// 工作空间名字
	Workspace *string `json:"Workspace,omitempty" xml:"Workspace,omitempty"`
	// 任务ID
	RunId      *string `json:"RunId,omitempty" xml:"RunId,omitempty"`
	Attributes *string `json:"Attributes,omitempty" xml:"Attributes,omitempty"`
}

func (s GetRunRequest) String() string {
	return tea.Prettify(s)
}

func (s GetRunRequest) GoString() string {
	return s.String()
}

func (s *GetRunRequest) SetWorkspace(v string) *GetRunRequest {
	s.Workspace = &v
	return s
}

func (s *GetRunRequest) SetRunId(v string) *GetRunRequest {
	s.RunId = &v
	return s
}

func (s *GetRunRequest) SetAttributes(v string) *GetRunRequest {
	s.Attributes = &v
	return s
}

type GetRunResponseBody struct {
	// 请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 主机ID
	HostId *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	// 工作空间名字
	Workspace *string `json:"Workspace,omitempty" xml:"Workspace,omitempty"`
	// 任务ID
	RunId *string `json:"RunId,omitempty" xml:"RunId,omitempty"`
	// 任务名称
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 提交ID
	SubmissionId *string `json:"SubmissionId,omitempty" xml:"SubmissionId,omitempty"`
	// 应用来源
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// 命名空间
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// 应用原名
	AppOrigName *string `json:"AppOrigName,omitempty" xml:"AppOrigName,omitempty"`
	// 应用名称
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// 应用版本
	Revision *string `json:"Revision,omitempty" xml:"Revision,omitempty"`
	// 实体类型
	EntityType *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	// 实体对象名称
	EntityName *string `json:"EntityName,omitempty" xml:"EntityName,omitempty"`
	// 用户ID
	User *string `json:"User,omitempty" xml:"User,omitempty"`
	// 任务状态
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 提交时间
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 开始时间
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// 结束时间
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// 任务配置
	ExecuteOptions *GetRunResponseBodyExecuteOptions `json:"ExecuteOptions,omitempty" xml:"ExecuteOptions,omitempty" type:"Struct"`
	Inputs         *string                           `json:"Inputs,omitempty" xml:"Inputs,omitempty"`
	Outputs        *string                           `json:"Outputs,omitempty" xml:"Outputs,omitempty"`
	// 任务标签
	Labels map[string]*string `json:"Labels,omitempty" xml:"Labels,omitempty"`
	// 输出拷贝目录
	OutputFolder *string `json:"OutputFolder,omitempty" xml:"OutputFolder,omitempty"`
	// 任务执行目录
	ExecuteDirectory *string `json:"ExecuteDirectory,omitempty" xml:"ExecuteDirectory,omitempty"`
	// 默认runtime值
	DefaultRuntime *string `json:"DefaultRuntime,omitempty" xml:"DefaultRuntime,omitempty"`
	// 任务描述
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 时序信息
	Timing   *string `json:"Timing,omitempty" xml:"Timing,omitempty"`
	Calls    *string `json:"Calls,omitempty" xml:"Calls,omitempty"`
	Failures *string `json:"Failures,omitempty" xml:"Failures,omitempty"`
}

func (s GetRunResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetRunResponseBody) GoString() string {
	return s.String()
}

func (s *GetRunResponseBody) SetRequestId(v string) *GetRunResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRunResponseBody) SetHostId(v string) *GetRunResponseBody {
	s.HostId = &v
	return s
}

func (s *GetRunResponseBody) SetWorkspace(v string) *GetRunResponseBody {
	s.Workspace = &v
	return s
}

func (s *GetRunResponseBody) SetRunId(v string) *GetRunResponseBody {
	s.RunId = &v
	return s
}

func (s *GetRunResponseBody) SetName(v string) *GetRunResponseBody {
	s.Name = &v
	return s
}

func (s *GetRunResponseBody) SetSubmissionId(v string) *GetRunResponseBody {
	s.SubmissionId = &v
	return s
}

func (s *GetRunResponseBody) SetSource(v string) *GetRunResponseBody {
	s.Source = &v
	return s
}

func (s *GetRunResponseBody) SetNamespace(v string) *GetRunResponseBody {
	s.Namespace = &v
	return s
}

func (s *GetRunResponseBody) SetAppOrigName(v string) *GetRunResponseBody {
	s.AppOrigName = &v
	return s
}

func (s *GetRunResponseBody) SetAppName(v string) *GetRunResponseBody {
	s.AppName = &v
	return s
}

func (s *GetRunResponseBody) SetRevision(v string) *GetRunResponseBody {
	s.Revision = &v
	return s
}

func (s *GetRunResponseBody) SetEntityType(v string) *GetRunResponseBody {
	s.EntityType = &v
	return s
}

func (s *GetRunResponseBody) SetEntityName(v string) *GetRunResponseBody {
	s.EntityName = &v
	return s
}

func (s *GetRunResponseBody) SetUser(v string) *GetRunResponseBody {
	s.User = &v
	return s
}

func (s *GetRunResponseBody) SetStatus(v string) *GetRunResponseBody {
	s.Status = &v
	return s
}

func (s *GetRunResponseBody) SetCreateTime(v string) *GetRunResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetRunResponseBody) SetStartTime(v string) *GetRunResponseBody {
	s.StartTime = &v
	return s
}

func (s *GetRunResponseBody) SetEndTime(v string) *GetRunResponseBody {
	s.EndTime = &v
	return s
}

func (s *GetRunResponseBody) SetExecuteOptions(v *GetRunResponseBodyExecuteOptions) *GetRunResponseBody {
	s.ExecuteOptions = v
	return s
}

func (s *GetRunResponseBody) SetInputs(v string) *GetRunResponseBody {
	s.Inputs = &v
	return s
}

func (s *GetRunResponseBody) SetOutputs(v string) *GetRunResponseBody {
	s.Outputs = &v
	return s
}

func (s *GetRunResponseBody) SetLabels(v map[string]*string) *GetRunResponseBody {
	s.Labels = v
	return s
}

func (s *GetRunResponseBody) SetOutputFolder(v string) *GetRunResponseBody {
	s.OutputFolder = &v
	return s
}

func (s *GetRunResponseBody) SetExecuteDirectory(v string) *GetRunResponseBody {
	s.ExecuteDirectory = &v
	return s
}

func (s *GetRunResponseBody) SetDefaultRuntime(v string) *GetRunResponseBody {
	s.DefaultRuntime = &v
	return s
}

func (s *GetRunResponseBody) SetDescription(v string) *GetRunResponseBody {
	s.Description = &v
	return s
}

func (s *GetRunResponseBody) SetTiming(v string) *GetRunResponseBody {
	s.Timing = &v
	return s
}

func (s *GetRunResponseBody) SetCalls(v string) *GetRunResponseBody {
	s.Calls = &v
	return s
}

func (s *GetRunResponseBody) SetFailures(v string) *GetRunResponseBody {
	s.Failures = &v
	return s
}

type GetRunResponseBodyExecuteOptions struct {
	CallCaching               *bool   `json:"CallCaching,omitempty" xml:"CallCaching,omitempty"`
	DeleteIntermediateResults *bool   `json:"DeleteIntermediateResults,omitempty" xml:"DeleteIntermediateResults,omitempty"`
	FailureMode               *string `json:"FailureMode,omitempty" xml:"FailureMode,omitempty"`
}

func (s GetRunResponseBodyExecuteOptions) String() string {
	return tea.Prettify(s)
}

func (s GetRunResponseBodyExecuteOptions) GoString() string {
	return s.String()
}

func (s *GetRunResponseBodyExecuteOptions) SetCallCaching(v bool) *GetRunResponseBodyExecuteOptions {
	s.CallCaching = &v
	return s
}

func (s *GetRunResponseBodyExecuteOptions) SetDeleteIntermediateResults(v bool) *GetRunResponseBodyExecuteOptions {
	s.DeleteIntermediateResults = &v
	return s
}

func (s *GetRunResponseBodyExecuteOptions) SetFailureMode(v string) *GetRunResponseBodyExecuteOptions {
	s.FailureMode = &v
	return s
}

type GetRunResponse struct {
	Headers map[string]*string  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetRunResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetRunResponse) SetBody(v *GetRunResponseBody) *GetRunResponse {
	s.Body = v
	return s
}

type GetSubmissionRequest struct {
	// 工作空间
	Workspace *string `json:"Workspace,omitempty" xml:"Workspace,omitempty"`
	// 投递ID
	SubmissionId *string `json:"SubmissionId,omitempty" xml:"SubmissionId,omitempty"`
}

func (s GetSubmissionRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSubmissionRequest) GoString() string {
	return s.String()
}

func (s *GetSubmissionRequest) SetWorkspace(v string) *GetSubmissionRequest {
	s.Workspace = &v
	return s
}

func (s *GetSubmissionRequest) SetSubmissionId(v string) *GetSubmissionRequest {
	s.SubmissionId = &v
	return s
}

type GetSubmissionResponseBody struct {
	// 请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 主机ID
	HostId *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	// 投递列表
	Submission *GetSubmissionResponseBodySubmission `json:"Submission,omitempty" xml:"Submission,omitempty" type:"Struct"`
}

func (s GetSubmissionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSubmissionResponseBody) GoString() string {
	return s.String()
}

func (s *GetSubmissionResponseBody) SetRequestId(v string) *GetSubmissionResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSubmissionResponseBody) SetHostId(v string) *GetSubmissionResponseBody {
	s.HostId = &v
	return s
}

func (s *GetSubmissionResponseBody) SetSubmission(v *GetSubmissionResponseBodySubmission) *GetSubmissionResponseBody {
	s.Submission = v
	return s
}

type GetSubmissionResponseBodySubmission struct {
	// 提交ID
	Workspace *string `json:"Workspace,omitempty" xml:"Workspace,omitempty"`
	// 提交ID
	SubmissionId *string `json:"SubmissionId,omitempty" xml:"SubmissionId,omitempty"`
	// 任务状态
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 提交时间
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 开始时间
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// 结束时间
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// 实体类型
	EntityType *string                                      `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	RunStats   *GetSubmissionResponseBodySubmissionRunStats `json:"RunStats,omitempty" xml:"RunStats,omitempty" type:"Struct"`
}

func (s GetSubmissionResponseBodySubmission) String() string {
	return tea.Prettify(s)
}

func (s GetSubmissionResponseBodySubmission) GoString() string {
	return s.String()
}

func (s *GetSubmissionResponseBodySubmission) SetWorkspace(v string) *GetSubmissionResponseBodySubmission {
	s.Workspace = &v
	return s
}

func (s *GetSubmissionResponseBodySubmission) SetSubmissionId(v string) *GetSubmissionResponseBodySubmission {
	s.SubmissionId = &v
	return s
}

func (s *GetSubmissionResponseBodySubmission) SetStatus(v string) *GetSubmissionResponseBodySubmission {
	s.Status = &v
	return s
}

func (s *GetSubmissionResponseBodySubmission) SetCreateTime(v string) *GetSubmissionResponseBodySubmission {
	s.CreateTime = &v
	return s
}

func (s *GetSubmissionResponseBodySubmission) SetStartTime(v string) *GetSubmissionResponseBodySubmission {
	s.StartTime = &v
	return s
}

func (s *GetSubmissionResponseBodySubmission) SetEndTime(v string) *GetSubmissionResponseBodySubmission {
	s.EndTime = &v
	return s
}

func (s *GetSubmissionResponseBodySubmission) SetEntityType(v string) *GetSubmissionResponseBodySubmission {
	s.EntityType = &v
	return s
}

func (s *GetSubmissionResponseBodySubmission) SetRunStats(v *GetSubmissionResponseBodySubmissionRunStats) *GetSubmissionResponseBodySubmission {
	s.RunStats = v
	return s
}

type GetSubmissionResponseBodySubmissionRunStats struct {
	Submitted *int64 `json:"Submitted,omitempty" xml:"Submitted,omitempty"`
	Pending   *int64 `json:"Pending,omitempty" xml:"Pending,omitempty"`
	Running   *int64 `json:"Running,omitempty" xml:"Running,omitempty"`
	Succeeded *int64 `json:"Succeeded,omitempty" xml:"Succeeded,omitempty"`
	Failed    *int64 `json:"Failed,omitempty" xml:"Failed,omitempty"`
	Aborting  *int64 `json:"Aborting,omitempty" xml:"Aborting,omitempty"`
	Aborted   *int64 `json:"Aborted,omitempty" xml:"Aborted,omitempty"`
}

func (s GetSubmissionResponseBodySubmissionRunStats) String() string {
	return tea.Prettify(s)
}

func (s GetSubmissionResponseBodySubmissionRunStats) GoString() string {
	return s.String()
}

func (s *GetSubmissionResponseBodySubmissionRunStats) SetSubmitted(v int64) *GetSubmissionResponseBodySubmissionRunStats {
	s.Submitted = &v
	return s
}

func (s *GetSubmissionResponseBodySubmissionRunStats) SetPending(v int64) *GetSubmissionResponseBodySubmissionRunStats {
	s.Pending = &v
	return s
}

func (s *GetSubmissionResponseBodySubmissionRunStats) SetRunning(v int64) *GetSubmissionResponseBodySubmissionRunStats {
	s.Running = &v
	return s
}

func (s *GetSubmissionResponseBodySubmissionRunStats) SetSucceeded(v int64) *GetSubmissionResponseBodySubmissionRunStats {
	s.Succeeded = &v
	return s
}

func (s *GetSubmissionResponseBodySubmissionRunStats) SetFailed(v int64) *GetSubmissionResponseBodySubmissionRunStats {
	s.Failed = &v
	return s
}

func (s *GetSubmissionResponseBodySubmissionRunStats) SetAborting(v int64) *GetSubmissionResponseBodySubmissionRunStats {
	s.Aborting = &v
	return s
}

func (s *GetSubmissionResponseBodySubmissionRunStats) SetAborted(v int64) *GetSubmissionResponseBodySubmissionRunStats {
	s.Aborted = &v
	return s
}

type GetSubmissionResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetSubmissionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetSubmissionResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSubmissionResponse) GoString() string {
	return s.String()
}

func (s *GetSubmissionResponse) SetHeaders(v map[string]*string) *GetSubmissionResponse {
	s.Headers = v
	return s
}

func (s *GetSubmissionResponse) SetBody(v *GetSubmissionResponseBody) *GetSubmissionResponse {
	s.Body = v
	return s
}

type GetTemplateRequest struct {
	// 工作空间名称
	Workspace *string `json:"Workspace,omitempty" xml:"Workspace,omitempty"`
	// 应用模板名称
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
}

func (s GetTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTemplateRequest) GoString() string {
	return s.String()
}

func (s *GetTemplateRequest) SetWorkspace(v string) *GetTemplateRequest {
	s.Workspace = &v
	return s
}

func (s *GetTemplateRequest) SetTemplateName(v string) *GetTemplateRequest {
	s.TemplateName = &v
	return s
}

type GetTemplateResponseBody struct {
	// 应用模板名称
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 应用简要描述
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 应用的名称
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// 应用的版本
	Revision *string `json:"Revision,omitempty" xml:"Revision,omitempty"`
	// 主机ID
	HostId *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	// 应用输入
	Inputs []*GetTemplateResponseBodyInputs `json:"Inputs,omitempty" xml:"Inputs,omitempty" type:"Repeated"`
	// 应用标签
	Labels map[string]*string `json:"Labels,omitempty" xml:"Labels,omitempty"`
	// 应用最后修改时间
	LastModifiedTime *string `json:"LastModifiedTime,omitempty" xml:"LastModifiedTime,omitempty"`
	// 应用的输出参数
	Outputs []*GetTemplateResponseBodyOutputs `json:"Outputs,omitempty" xml:"Outputs,omitempty" type:"Repeated"`
	// 请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 应用来源
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// 工作空间名称
	Workspace *string `json:"Workspace,omitempty" xml:"Workspace,omitempty"`
	// 实体类型
	RootEntity *string `json:"RootEntity,omitempty" xml:"RootEntity,omitempty"`
}

func (s GetTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *GetTemplateResponseBody) SetTemplateName(v string) *GetTemplateResponseBody {
	s.TemplateName = &v
	return s
}

func (s *GetTemplateResponseBody) SetCreateTime(v string) *GetTemplateResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetTemplateResponseBody) SetDescription(v string) *GetTemplateResponseBody {
	s.Description = &v
	return s
}

func (s *GetTemplateResponseBody) SetAppName(v string) *GetTemplateResponseBody {
	s.AppName = &v
	return s
}

func (s *GetTemplateResponseBody) SetRevision(v string) *GetTemplateResponseBody {
	s.Revision = &v
	return s
}

func (s *GetTemplateResponseBody) SetHostId(v string) *GetTemplateResponseBody {
	s.HostId = &v
	return s
}

func (s *GetTemplateResponseBody) SetInputs(v []*GetTemplateResponseBodyInputs) *GetTemplateResponseBody {
	s.Inputs = v
	return s
}

func (s *GetTemplateResponseBody) SetLabels(v map[string]*string) *GetTemplateResponseBody {
	s.Labels = v
	return s
}

func (s *GetTemplateResponseBody) SetLastModifiedTime(v string) *GetTemplateResponseBody {
	s.LastModifiedTime = &v
	return s
}

func (s *GetTemplateResponseBody) SetOutputs(v []*GetTemplateResponseBodyOutputs) *GetTemplateResponseBody {
	s.Outputs = v
	return s
}

func (s *GetTemplateResponseBody) SetRequestId(v string) *GetTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTemplateResponseBody) SetSource(v string) *GetTemplateResponseBody {
	s.Source = &v
	return s
}

func (s *GetTemplateResponseBody) SetWorkspace(v string) *GetTemplateResponseBody {
	s.Workspace = &v
	return s
}

func (s *GetTemplateResponseBody) SetRootEntity(v string) *GetTemplateResponseBody {
	s.RootEntity = &v
	return s
}

type GetTemplateResponseBodyInputs struct {
	// 任务名称
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// 变量名称
	VariableName *string `json:"VariableName,omitempty" xml:"VariableName,omitempty"`
	// 变量类型
	VariableType *string `json:"VariableType,omitempty" xml:"VariableType,omitempty"`
	// 变量值
	VariableValue *string `json:"VariableValue,omitempty" xml:"VariableValue,omitempty"`
	// 是否必须参数
	Required *bool `json:"Required,omitempty" xml:"Required,omitempty"`
	// 帮助信息
	Help *string `json:"Help,omitempty" xml:"Help,omitempty"`
	// 步骤顺序
	StepOrder *int64 `json:"StepOrder,omitempty" xml:"StepOrder,omitempty"`
}

func (s GetTemplateResponseBodyInputs) String() string {
	return tea.Prettify(s)
}

func (s GetTemplateResponseBodyInputs) GoString() string {
	return s.String()
}

func (s *GetTemplateResponseBodyInputs) SetTaskName(v string) *GetTemplateResponseBodyInputs {
	s.TaskName = &v
	return s
}

func (s *GetTemplateResponseBodyInputs) SetVariableName(v string) *GetTemplateResponseBodyInputs {
	s.VariableName = &v
	return s
}

func (s *GetTemplateResponseBodyInputs) SetVariableType(v string) *GetTemplateResponseBodyInputs {
	s.VariableType = &v
	return s
}

func (s *GetTemplateResponseBodyInputs) SetVariableValue(v string) *GetTemplateResponseBodyInputs {
	s.VariableValue = &v
	return s
}

func (s *GetTemplateResponseBodyInputs) SetRequired(v bool) *GetTemplateResponseBodyInputs {
	s.Required = &v
	return s
}

func (s *GetTemplateResponseBodyInputs) SetHelp(v string) *GetTemplateResponseBodyInputs {
	s.Help = &v
	return s
}

func (s *GetTemplateResponseBodyInputs) SetStepOrder(v int64) *GetTemplateResponseBodyInputs {
	s.StepOrder = &v
	return s
}

type GetTemplateResponseBodyOutputs struct {
	// 任务名称
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// 变量名称
	VariableName *string `json:"VariableName,omitempty" xml:"VariableName,omitempty"`
	// 变量类型
	VariableType *string `json:"VariableType,omitempty" xml:"VariableType,omitempty"`
	// 变量值
	VariableValue *string `json:"VariableValue,omitempty" xml:"VariableValue,omitempty"`
	// 是否必须参数
	Required *bool `json:"Required,omitempty" xml:"Required,omitempty"`
	// 帮助信息
	Help *string `json:"Help,omitempty" xml:"Help,omitempty"`
	// 步骤顺序
	StepOrder *int64 `json:"StepOrder,omitempty" xml:"StepOrder,omitempty"`
}

func (s GetTemplateResponseBodyOutputs) String() string {
	return tea.Prettify(s)
}

func (s GetTemplateResponseBodyOutputs) GoString() string {
	return s.String()
}

func (s *GetTemplateResponseBodyOutputs) SetTaskName(v string) *GetTemplateResponseBodyOutputs {
	s.TaskName = &v
	return s
}

func (s *GetTemplateResponseBodyOutputs) SetVariableName(v string) *GetTemplateResponseBodyOutputs {
	s.VariableName = &v
	return s
}

func (s *GetTemplateResponseBodyOutputs) SetVariableType(v string) *GetTemplateResponseBodyOutputs {
	s.VariableType = &v
	return s
}

func (s *GetTemplateResponseBodyOutputs) SetVariableValue(v string) *GetTemplateResponseBodyOutputs {
	s.VariableValue = &v
	return s
}

func (s *GetTemplateResponseBodyOutputs) SetRequired(v bool) *GetTemplateResponseBodyOutputs {
	s.Required = &v
	return s
}

func (s *GetTemplateResponseBodyOutputs) SetHelp(v string) *GetTemplateResponseBodyOutputs {
	s.Help = &v
	return s
}

func (s *GetTemplateResponseBodyOutputs) SetStepOrder(v int64) *GetTemplateResponseBodyOutputs {
	s.StepOrder = &v
	return s
}

type GetTemplateResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTemplateResponse) GoString() string {
	return s.String()
}

func (s *GetTemplateResponse) SetHeaders(v map[string]*string) *GetTemplateResponse {
	s.Headers = v
	return s
}

func (s *GetTemplateResponse) SetBody(v *GetTemplateResponseBody) *GetTemplateResponse {
	s.Body = v
	return s
}

type GetWorkspaceRequest struct {
	Workspace *string `json:"Workspace,omitempty" xml:"Workspace,omitempty"`
}

func (s GetWorkspaceRequest) String() string {
	return tea.Prettify(s)
}

func (s GetWorkspaceRequest) GoString() string {
	return s.String()
}

func (s *GetWorkspaceRequest) SetWorkspace(v string) *GetWorkspaceRequest {
	s.Workspace = &v
	return s
}

type GetWorkspaceResponseBody struct {
	// 请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 工作空间简要描述
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 主机ID
	HostId *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	// 工作空间内作业生命周期
	JobLifecycle *int32 `json:"JobLifecycle,omitempty" xml:"JobLifecycle,omitempty"`
	// 工作空间标签
	Labels map[string]*string `json:"Labels,omitempty" xml:"Labels,omitempty"`
	// 工作空间内OSS上的工作路径
	OssRoot *string `json:"OssRoot,omitempty" xml:"OssRoot,omitempty"`
	// 工作空间Bucket
	BucketName *string `json:"BucketName,omitempty" xml:"BucketName,omitempty"`
	// 工作空间内默认的RAM服务角色
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 最后修改时间
	LastModifiedTime *string `json:"LastModifiedTime,omitempty" xml:"LastModifiedTime,omitempty"`
	// 工作空间名称
	Workspace *string `json:"Workspace,omitempty" xml:"Workspace,omitempty"`
	// 工作空间状态
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 地域ID
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetWorkspaceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetWorkspaceResponseBody) GoString() string {
	return s.String()
}

func (s *GetWorkspaceResponseBody) SetRequestId(v string) *GetWorkspaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetWorkspaceResponseBody) SetDescription(v string) *GetWorkspaceResponseBody {
	s.Description = &v
	return s
}

func (s *GetWorkspaceResponseBody) SetHostId(v string) *GetWorkspaceResponseBody {
	s.HostId = &v
	return s
}

func (s *GetWorkspaceResponseBody) SetJobLifecycle(v int32) *GetWorkspaceResponseBody {
	s.JobLifecycle = &v
	return s
}

func (s *GetWorkspaceResponseBody) SetLabels(v map[string]*string) *GetWorkspaceResponseBody {
	s.Labels = v
	return s
}

func (s *GetWorkspaceResponseBody) SetOssRoot(v string) *GetWorkspaceResponseBody {
	s.OssRoot = &v
	return s
}

func (s *GetWorkspaceResponseBody) SetBucketName(v string) *GetWorkspaceResponseBody {
	s.BucketName = &v
	return s
}

func (s *GetWorkspaceResponseBody) SetRole(v string) *GetWorkspaceResponseBody {
	s.Role = &v
	return s
}

func (s *GetWorkspaceResponseBody) SetCreateTime(v string) *GetWorkspaceResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetWorkspaceResponseBody) SetLastModifiedTime(v string) *GetWorkspaceResponseBody {
	s.LastModifiedTime = &v
	return s
}

func (s *GetWorkspaceResponseBody) SetWorkspace(v string) *GetWorkspaceResponseBody {
	s.Workspace = &v
	return s
}

func (s *GetWorkspaceResponseBody) SetStatus(v string) *GetWorkspaceResponseBody {
	s.Status = &v
	return s
}

func (s *GetWorkspaceResponseBody) SetRegionId(v string) *GetWorkspaceResponseBody {
	s.RegionId = &v
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

type InstallGlobalAppRequest struct {
	// 来源
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// 命名空间名称
	NamespaceName *string `json:"NamespaceName,omitempty" xml:"NamespaceName,omitempty"`
	// 应用描述
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// 工作空间
	Workspace *string `json:"Workspace,omitempty" xml:"Workspace,omitempty"`
	// 安装后应用名
	InstalledAppName *string `json:"InstalledAppName,omitempty" xml:"InstalledAppName,omitempty"`
}

func (s InstallGlobalAppRequest) String() string {
	return tea.Prettify(s)
}

func (s InstallGlobalAppRequest) GoString() string {
	return s.String()
}

func (s *InstallGlobalAppRequest) SetSource(v string) *InstallGlobalAppRequest {
	s.Source = &v
	return s
}

func (s *InstallGlobalAppRequest) SetNamespaceName(v string) *InstallGlobalAppRequest {
	s.NamespaceName = &v
	return s
}

func (s *InstallGlobalAppRequest) SetAppName(v string) *InstallGlobalAppRequest {
	s.AppName = &v
	return s
}

func (s *InstallGlobalAppRequest) SetWorkspace(v string) *InstallGlobalAppRequest {
	s.Workspace = &v
	return s
}

func (s *InstallGlobalAppRequest) SetInstalledAppName(v string) *InstallGlobalAppRequest {
	s.InstalledAppName = &v
	return s
}

type InstallGlobalAppResponseBody struct {
	// 主机 ID
	HostId *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	// 请求 ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 安装后应用名
	InstalledAppName *string `json:"InstalledAppName,omitempty" xml:"InstalledAppName,omitempty"`
	// 工作空间
	Workspace *string `json:"Workspace,omitempty" xml:"Workspace,omitempty"`
	// 区域名
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s InstallGlobalAppResponseBody) String() string {
	return tea.Prettify(s)
}

func (s InstallGlobalAppResponseBody) GoString() string {
	return s.String()
}

func (s *InstallGlobalAppResponseBody) SetHostId(v string) *InstallGlobalAppResponseBody {
	s.HostId = &v
	return s
}

func (s *InstallGlobalAppResponseBody) SetRequestId(v string) *InstallGlobalAppResponseBody {
	s.RequestId = &v
	return s
}

func (s *InstallGlobalAppResponseBody) SetInstalledAppName(v string) *InstallGlobalAppResponseBody {
	s.InstalledAppName = &v
	return s
}

func (s *InstallGlobalAppResponseBody) SetWorkspace(v string) *InstallGlobalAppResponseBody {
	s.Workspace = &v
	return s
}

func (s *InstallGlobalAppResponseBody) SetRegionId(v string) *InstallGlobalAppResponseBody {
	s.RegionId = &v
	return s
}

type InstallGlobalAppResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *InstallGlobalAppResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s InstallGlobalAppResponse) String() string {
	return tea.Prettify(s)
}

func (s InstallGlobalAppResponse) GoString() string {
	return s.String()
}

func (s *InstallGlobalAppResponse) SetHeaders(v map[string]*string) *InstallGlobalAppResponse {
	s.Headers = v
	return s
}

func (s *InstallGlobalAppResponse) SetBody(v *InstallGlobalAppResponseBody) *InstallGlobalAppResponse {
	s.Body = v
	return s
}

type ListAppsRequest struct {
	// 工作空间
	Workspace *string `json:"Workspace,omitempty" xml:"Workspace,omitempty"`
	// Next Token
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// 最大返回结果数
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// 排序依据
	OrderBy *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	// 是否逆序
	IsReversed *bool `json:"IsReversed,omitempty" xml:"IsReversed,omitempty"`
	// Label 选择器
	LabelSelector *string `json:"LabelSelector,omitempty" xml:"LabelSelector,omitempty"`
	// 应用范围
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
	// 应用描述语言
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// 应用类型
	AppType *string `json:"AppType,omitempty" xml:"AppType,omitempty"`
	// 按照名字匹配
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListAppsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAppsRequest) GoString() string {
	return s.String()
}

func (s *ListAppsRequest) SetWorkspace(v string) *ListAppsRequest {
	s.Workspace = &v
	return s
}

func (s *ListAppsRequest) SetNextToken(v string) *ListAppsRequest {
	s.NextToken = &v
	return s
}

func (s *ListAppsRequest) SetMaxResults(v int32) *ListAppsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListAppsRequest) SetOrderBy(v string) *ListAppsRequest {
	s.OrderBy = &v
	return s
}

func (s *ListAppsRequest) SetIsReversed(v bool) *ListAppsRequest {
	s.IsReversed = &v
	return s
}

func (s *ListAppsRequest) SetLabelSelector(v string) *ListAppsRequest {
	s.LabelSelector = &v
	return s
}

func (s *ListAppsRequest) SetScope(v string) *ListAppsRequest {
	s.Scope = &v
	return s
}

func (s *ListAppsRequest) SetLanguage(v string) *ListAppsRequest {
	s.Language = &v
	return s
}

func (s *ListAppsRequest) SetAppType(v string) *ListAppsRequest {
	s.AppType = &v
	return s
}

func (s *ListAppsRequest) SetName(v string) *ListAppsRequest {
	s.Name = &v
	return s
}

type ListAppsResponseBody struct {
	// 应用数组
	Apps []*ListAppsResponseBodyApps `json:"Apps,omitempty" xml:"Apps,omitempty" type:"Repeated"`
	// 主机ID
	HostId *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	// 最大返回个数
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// Next Token
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// 请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 返回个数
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListAppsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAppsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAppsResponseBody) SetApps(v []*ListAppsResponseBodyApps) *ListAppsResponseBody {
	s.Apps = v
	return s
}

func (s *ListAppsResponseBody) SetHostId(v string) *ListAppsResponseBody {
	s.HostId = &v
	return s
}

func (s *ListAppsResponseBody) SetMaxResults(v int32) *ListAppsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListAppsResponseBody) SetNextToken(v string) *ListAppsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListAppsResponseBody) SetRequestId(v string) *ListAppsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAppsResponseBody) SetTotalCount(v int32) *ListAppsResponseBody {
	s.TotalCount = &v
	return s
}

type ListAppsResponseBodyApps struct {
	// 应用名称
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 应用描述
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 应用类型
	AppType *string `json:"AppType,omitempty" xml:"AppType,omitempty"`
	// 应用描述语言
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// 应用可见范围
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
	// 应用所在工作空间
	Workspace *string `json:"Workspace,omitempty" xml:"Workspace,omitempty"`
	// 应用来源
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// 命名空间
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// 应用原名
	AppOrigName *string `json:"AppOrigName,omitempty" xml:"AppOrigName,omitempty"`
	// 默认版本
	AppDefaultVersion *string `json:"AppDefaultVersion,omitempty" xml:"AppDefaultVersion,omitempty"`
	// 标签
	Labels map[string]*string `json:"Labels,omitempty" xml:"Labels,omitempty"`
}

func (s ListAppsResponseBodyApps) String() string {
	return tea.Prettify(s)
}

func (s ListAppsResponseBodyApps) GoString() string {
	return s.String()
}

func (s *ListAppsResponseBodyApps) SetAppName(v string) *ListAppsResponseBodyApps {
	s.AppName = &v
	return s
}

func (s *ListAppsResponseBodyApps) SetCreateTime(v string) *ListAppsResponseBodyApps {
	s.CreateTime = &v
	return s
}

func (s *ListAppsResponseBodyApps) SetDescription(v string) *ListAppsResponseBodyApps {
	s.Description = &v
	return s
}

func (s *ListAppsResponseBodyApps) SetAppType(v string) *ListAppsResponseBodyApps {
	s.AppType = &v
	return s
}

func (s *ListAppsResponseBodyApps) SetLanguage(v string) *ListAppsResponseBodyApps {
	s.Language = &v
	return s
}

func (s *ListAppsResponseBodyApps) SetScope(v string) *ListAppsResponseBodyApps {
	s.Scope = &v
	return s
}

func (s *ListAppsResponseBodyApps) SetWorkspace(v string) *ListAppsResponseBodyApps {
	s.Workspace = &v
	return s
}

func (s *ListAppsResponseBodyApps) SetSource(v string) *ListAppsResponseBodyApps {
	s.Source = &v
	return s
}

func (s *ListAppsResponseBodyApps) SetNamespace(v string) *ListAppsResponseBodyApps {
	s.Namespace = &v
	return s
}

func (s *ListAppsResponseBodyApps) SetAppOrigName(v string) *ListAppsResponseBodyApps {
	s.AppOrigName = &v
	return s
}

func (s *ListAppsResponseBodyApps) SetAppDefaultVersion(v string) *ListAppsResponseBodyApps {
	s.AppDefaultVersion = &v
	return s
}

func (s *ListAppsResponseBodyApps) SetLabels(v map[string]*string) *ListAppsResponseBodyApps {
	s.Labels = v
	return s
}

type ListAppsResponse struct {
	Headers map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListAppsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListAppsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAppsResponse) GoString() string {
	return s.String()
}

func (s *ListAppsResponse) SetHeaders(v map[string]*string) *ListAppsResponse {
	s.Headers = v
	return s
}

func (s *ListAppsResponse) SetBody(v *ListAppsResponseBody) *ListAppsResponse {
	s.Body = v
	return s
}

type ListAuthorizedSoftwareRequest struct {
	// 区域
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// 软件名称、软件长名称中搜索的关键字
	Search *string `json:"Search,omitempty" xml:"Search,omitempty"`
	// 排序字段
	OrderBy *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	// 是否反转
	IsReversed *bool `json:"IsReversed,omitempty" xml:"IsReversed,omitempty"`
	// 翻页Token
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// 分页数量
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
}

func (s ListAuthorizedSoftwareRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAuthorizedSoftwareRequest) GoString() string {
	return s.String()
}

func (s *ListAuthorizedSoftwareRequest) SetRegion(v string) *ListAuthorizedSoftwareRequest {
	s.Region = &v
	return s
}

func (s *ListAuthorizedSoftwareRequest) SetSearch(v string) *ListAuthorizedSoftwareRequest {
	s.Search = &v
	return s
}

func (s *ListAuthorizedSoftwareRequest) SetOrderBy(v string) *ListAuthorizedSoftwareRequest {
	s.OrderBy = &v
	return s
}

func (s *ListAuthorizedSoftwareRequest) SetIsReversed(v bool) *ListAuthorizedSoftwareRequest {
	s.IsReversed = &v
	return s
}

func (s *ListAuthorizedSoftwareRequest) SetNextToken(v string) *ListAuthorizedSoftwareRequest {
	s.NextToken = &v
	return s
}

func (s *ListAuthorizedSoftwareRequest) SetMaxResults(v int32) *ListAuthorizedSoftwareRequest {
	s.MaxResults = &v
	return s
}

type ListAuthorizedSoftwareResponseBody struct {
	// 主机ID
	HostId *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	// 请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	HttpCode  *int32  `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	// 错误码
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 错误消息
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 是否调用成功
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	DynamicCode    *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// 翻页TOKEN
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// 分页数
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// 总记录数
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// 软件信息
	Softwares []*ListAuthorizedSoftwareResponseBodySoftwares `json:"Softwares,omitempty" xml:"Softwares,omitempty" type:"Repeated"`
}

func (s ListAuthorizedSoftwareResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAuthorizedSoftwareResponseBody) GoString() string {
	return s.String()
}

func (s *ListAuthorizedSoftwareResponseBody) SetHostId(v string) *ListAuthorizedSoftwareResponseBody {
	s.HostId = &v
	return s
}

func (s *ListAuthorizedSoftwareResponseBody) SetRequestId(v string) *ListAuthorizedSoftwareResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAuthorizedSoftwareResponseBody) SetHttpCode(v int32) *ListAuthorizedSoftwareResponseBody {
	s.HttpCode = &v
	return s
}

func (s *ListAuthorizedSoftwareResponseBody) SetCode(v string) *ListAuthorizedSoftwareResponseBody {
	s.Code = &v
	return s
}

func (s *ListAuthorizedSoftwareResponseBody) SetMessage(v string) *ListAuthorizedSoftwareResponseBody {
	s.Message = &v
	return s
}

func (s *ListAuthorizedSoftwareResponseBody) SetSuccess(v bool) *ListAuthorizedSoftwareResponseBody {
	s.Success = &v
	return s
}

func (s *ListAuthorizedSoftwareResponseBody) SetDynamicCode(v string) *ListAuthorizedSoftwareResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *ListAuthorizedSoftwareResponseBody) SetDynamicMessage(v string) *ListAuthorizedSoftwareResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *ListAuthorizedSoftwareResponseBody) SetNextToken(v string) *ListAuthorizedSoftwareResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListAuthorizedSoftwareResponseBody) SetMaxResults(v int32) *ListAuthorizedSoftwareResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListAuthorizedSoftwareResponseBody) SetTotalCount(v int32) *ListAuthorizedSoftwareResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListAuthorizedSoftwareResponseBody) SetSoftwares(v []*ListAuthorizedSoftwareResponseBodySoftwares) *ListAuthorizedSoftwareResponseBody {
	s.Softwares = v
	return s
}

type ListAuthorizedSoftwareResponseBodySoftwares struct {
	// 软件名称
	SoftwareName *string `json:"SoftwareName,omitempty" xml:"SoftwareName,omitempty"`
	// 软件长名称
	SoftwareLongName *string `json:"SoftwareLongName,omitempty" xml:"SoftwareLongName,omitempty"`
	// 软件描述
	SoftwareDescription *string `json:"SoftwareDescription,omitempty" xml:"SoftwareDescription,omitempty"`
	// 软件默认版本
	SoftwareDefaultVersion *string `json:"SoftwareDefaultVersion,omitempty" xml:"SoftwareDefaultVersion,omitempty"`
	// 软件所有版本
	SoftwareVersions []*string `json:"SoftwareVersions,omitempty" xml:"SoftwareVersions,omitempty" type:"Repeated"`
	// 软件使用费用
	SoftwareLicenseFee *float32 `json:"SoftwareLicenseFee,omitempty" xml:"SoftwareLicenseFee,omitempty"`
	// 软件可用区域
	RegionIds []*string `json:"RegionIds,omitempty" xml:"RegionIds,omitempty" type:"Repeated"`
	// 帮助链接
	HelpLink *string `json:"HelpLink,omitempty" xml:"HelpLink,omitempty"`
	// 软件图标链接
	SoftwareIcon *string `json:"SoftwareIcon,omitempty" xml:"SoftwareIcon,omitempty"`
	// 最后更新时间
	LastModified *string `json:"LastModified,omitempty" xml:"LastModified,omitempty"`
}

func (s ListAuthorizedSoftwareResponseBodySoftwares) String() string {
	return tea.Prettify(s)
}

func (s ListAuthorizedSoftwareResponseBodySoftwares) GoString() string {
	return s.String()
}

func (s *ListAuthorizedSoftwareResponseBodySoftwares) SetSoftwareName(v string) *ListAuthorizedSoftwareResponseBodySoftwares {
	s.SoftwareName = &v
	return s
}

func (s *ListAuthorizedSoftwareResponseBodySoftwares) SetSoftwareLongName(v string) *ListAuthorizedSoftwareResponseBodySoftwares {
	s.SoftwareLongName = &v
	return s
}

func (s *ListAuthorizedSoftwareResponseBodySoftwares) SetSoftwareDescription(v string) *ListAuthorizedSoftwareResponseBodySoftwares {
	s.SoftwareDescription = &v
	return s
}

func (s *ListAuthorizedSoftwareResponseBodySoftwares) SetSoftwareDefaultVersion(v string) *ListAuthorizedSoftwareResponseBodySoftwares {
	s.SoftwareDefaultVersion = &v
	return s
}

func (s *ListAuthorizedSoftwareResponseBodySoftwares) SetSoftwareVersions(v []*string) *ListAuthorizedSoftwareResponseBodySoftwares {
	s.SoftwareVersions = v
	return s
}

func (s *ListAuthorizedSoftwareResponseBodySoftwares) SetSoftwareLicenseFee(v float32) *ListAuthorizedSoftwareResponseBodySoftwares {
	s.SoftwareLicenseFee = &v
	return s
}

func (s *ListAuthorizedSoftwareResponseBodySoftwares) SetRegionIds(v []*string) *ListAuthorizedSoftwareResponseBodySoftwares {
	s.RegionIds = v
	return s
}

func (s *ListAuthorizedSoftwareResponseBodySoftwares) SetHelpLink(v string) *ListAuthorizedSoftwareResponseBodySoftwares {
	s.HelpLink = &v
	return s
}

func (s *ListAuthorizedSoftwareResponseBodySoftwares) SetSoftwareIcon(v string) *ListAuthorizedSoftwareResponseBodySoftwares {
	s.SoftwareIcon = &v
	return s
}

func (s *ListAuthorizedSoftwareResponseBodySoftwares) SetLastModified(v string) *ListAuthorizedSoftwareResponseBodySoftwares {
	s.LastModified = &v
	return s
}

type ListAuthorizedSoftwareResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListAuthorizedSoftwareResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListAuthorizedSoftwareResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAuthorizedSoftwareResponse) GoString() string {
	return s.String()
}

func (s *ListAuthorizedSoftwareResponse) SetHeaders(v map[string]*string) *ListAuthorizedSoftwareResponse {
	s.Headers = v
	return s
}

func (s *ListAuthorizedSoftwareResponse) SetBody(v *ListAuthorizedSoftwareResponseBody) *ListAuthorizedSoftwareResponse {
	s.Body = v
	return s
}

type ListContainerImagesRequest struct {
	// 区域
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// 翻页Token
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// 分页数量
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
}

func (s ListContainerImagesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListContainerImagesRequest) GoString() string {
	return s.String()
}

func (s *ListContainerImagesRequest) SetRegion(v string) *ListContainerImagesRequest {
	s.Region = &v
	return s
}

func (s *ListContainerImagesRequest) SetNextToken(v string) *ListContainerImagesRequest {
	s.NextToken = &v
	return s
}

func (s *ListContainerImagesRequest) SetMaxResults(v int32) *ListContainerImagesRequest {
	s.MaxResults = &v
	return s
}

type ListContainerImagesResponseBody struct {
	HostId     *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	HttpCode   *int32  `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	Code       *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message    *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Success    *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	NextToken  *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	MaxResults *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	TotalCount *int32  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// 容器镜像
	ContainerImages []*ListContainerImagesResponseBodyContainerImages `json:"ContainerImages,omitempty" xml:"ContainerImages,omitempty" type:"Repeated"`
}

func (s ListContainerImagesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListContainerImagesResponseBody) GoString() string {
	return s.String()
}

func (s *ListContainerImagesResponseBody) SetHostId(v string) *ListContainerImagesResponseBody {
	s.HostId = &v
	return s
}

func (s *ListContainerImagesResponseBody) SetRequestId(v string) *ListContainerImagesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListContainerImagesResponseBody) SetHttpCode(v int32) *ListContainerImagesResponseBody {
	s.HttpCode = &v
	return s
}

func (s *ListContainerImagesResponseBody) SetCode(v string) *ListContainerImagesResponseBody {
	s.Code = &v
	return s
}

func (s *ListContainerImagesResponseBody) SetMessage(v string) *ListContainerImagesResponseBody {
	s.Message = &v
	return s
}

func (s *ListContainerImagesResponseBody) SetSuccess(v bool) *ListContainerImagesResponseBody {
	s.Success = &v
	return s
}

func (s *ListContainerImagesResponseBody) SetNextToken(v string) *ListContainerImagesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListContainerImagesResponseBody) SetMaxResults(v int32) *ListContainerImagesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListContainerImagesResponseBody) SetTotalCount(v int32) *ListContainerImagesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListContainerImagesResponseBody) SetContainerImages(v []*ListContainerImagesResponseBodyContainerImages) *ListContainerImagesResponseBody {
	s.ContainerImages = v
	return s
}

type ListContainerImagesResponseBodyContainerImages struct {
	// 容器镜像名称空间名称
	ContainerImageNamespace *string `json:"ContainerImageNamespace,omitempty" xml:"ContainerImageNamespace,omitempty"`
	// 容器镜像名称
	ContainerImageName *string `json:"ContainerImageName,omitempty" xml:"ContainerImageName,omitempty"`
	// 容器镜像描述
	ContainerImageDescription *string `json:"ContainerImageDescription,omitempty" xml:"ContainerImageDescription,omitempty"`
	// 容器镜像版本
	ContainerImageVersions []*string `json:"ContainerImageVersions,omitempty" xml:"ContainerImageVersions,omitempty" type:"Repeated"`
	// 容器镜像所在区域
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// 容器镜像仓库名称
	ContainerRegistry *string `json:"ContainerRegistry,omitempty" xml:"ContainerRegistry,omitempty"`
	// 最后更新时间
	LastModified *string `json:"LastModified,omitempty" xml:"LastModified,omitempty"`
}

func (s ListContainerImagesResponseBodyContainerImages) String() string {
	return tea.Prettify(s)
}

func (s ListContainerImagesResponseBodyContainerImages) GoString() string {
	return s.String()
}

func (s *ListContainerImagesResponseBodyContainerImages) SetContainerImageNamespace(v string) *ListContainerImagesResponseBodyContainerImages {
	s.ContainerImageNamespace = &v
	return s
}

func (s *ListContainerImagesResponseBodyContainerImages) SetContainerImageName(v string) *ListContainerImagesResponseBodyContainerImages {
	s.ContainerImageName = &v
	return s
}

func (s *ListContainerImagesResponseBodyContainerImages) SetContainerImageDescription(v string) *ListContainerImagesResponseBodyContainerImages {
	s.ContainerImageDescription = &v
	return s
}

func (s *ListContainerImagesResponseBodyContainerImages) SetContainerImageVersions(v []*string) *ListContainerImagesResponseBodyContainerImages {
	s.ContainerImageVersions = v
	return s
}

func (s *ListContainerImagesResponseBodyContainerImages) SetRegionId(v string) *ListContainerImagesResponseBodyContainerImages {
	s.RegionId = &v
	return s
}

func (s *ListContainerImagesResponseBodyContainerImages) SetContainerRegistry(v string) *ListContainerImagesResponseBodyContainerImages {
	s.ContainerRegistry = &v
	return s
}

func (s *ListContainerImagesResponseBodyContainerImages) SetLastModified(v string) *ListContainerImagesResponseBodyContainerImages {
	s.LastModified = &v
	return s
}

type ListContainerImagesResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListContainerImagesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListContainerImagesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListContainerImagesResponse) GoString() string {
	return s.String()
}

func (s *ListContainerImagesResponse) SetHeaders(v map[string]*string) *ListContainerImagesResponse {
	s.Headers = v
	return s
}

func (s *ListContainerImagesResponse) SetBody(v *ListContainerImagesResponseBody) *ListContainerImagesResponse {
	s.Body = v
	return s
}

type ListEntitiesRequest struct {
	Workspace  *string `json:"Workspace,omitempty" xml:"Workspace,omitempty"`
	NextToken  *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	MaxResults *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	OrderBy    *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	IsReversed *bool   `json:"IsReversed,omitempty" xml:"IsReversed,omitempty"`
}

func (s ListEntitiesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListEntitiesRequest) GoString() string {
	return s.String()
}

func (s *ListEntitiesRequest) SetWorkspace(v string) *ListEntitiesRequest {
	s.Workspace = &v
	return s
}

func (s *ListEntitiesRequest) SetNextToken(v string) *ListEntitiesRequest {
	s.NextToken = &v
	return s
}

func (s *ListEntitiesRequest) SetMaxResults(v int32) *ListEntitiesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListEntitiesRequest) SetOrderBy(v string) *ListEntitiesRequest {
	s.OrderBy = &v
	return s
}

func (s *ListEntitiesRequest) SetIsReversed(v bool) *ListEntitiesRequest {
	s.IsReversed = &v
	return s
}

type ListEntitiesResponseBody struct {
	// 主机ID
	HostId *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	// 请求ID
	RequestId  *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	NextToken  *string                             `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	MaxResults *int32                              `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	TotalCount *int32                              `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	Entities   []*ListEntitiesResponseBodyEntities `json:"Entities,omitempty" xml:"Entities,omitempty" type:"Repeated"`
}

func (s ListEntitiesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListEntitiesResponseBody) GoString() string {
	return s.String()
}

func (s *ListEntitiesResponseBody) SetHostId(v string) *ListEntitiesResponseBody {
	s.HostId = &v
	return s
}

func (s *ListEntitiesResponseBody) SetRequestId(v string) *ListEntitiesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListEntitiesResponseBody) SetNextToken(v string) *ListEntitiesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListEntitiesResponseBody) SetMaxResults(v int32) *ListEntitiesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListEntitiesResponseBody) SetTotalCount(v int32) *ListEntitiesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListEntitiesResponseBody) SetEntities(v []*ListEntitiesResponseBodyEntities) *ListEntitiesResponseBody {
	s.Entities = v
	return s
}

type ListEntitiesResponseBodyEntities struct {
	// 实体类型
	EntityType *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
}

func (s ListEntitiesResponseBodyEntities) String() string {
	return tea.Prettify(s)
}

func (s ListEntitiesResponseBodyEntities) GoString() string {
	return s.String()
}

func (s *ListEntitiesResponseBodyEntities) SetEntityType(v string) *ListEntitiesResponseBodyEntities {
	s.EntityType = &v
	return s
}

type ListEntitiesResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListEntitiesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListEntitiesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListEntitiesResponse) GoString() string {
	return s.String()
}

func (s *ListEntitiesResponse) SetHeaders(v map[string]*string) *ListEntitiesResponse {
	s.Headers = v
	return s
}

func (s *ListEntitiesResponse) SetBody(v *ListEntitiesResponseBody) *ListEntitiesResponse {
	s.Body = v
	return s
}

type ListEntityItemsRequest struct {
	Workspace  *string `json:"Workspace,omitempty" xml:"Workspace,omitempty"`
	EntityType *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	Search     *string `json:"Search,omitempty" xml:"Search,omitempty"`
	OrderBy    *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	IsReversed *bool   `json:"IsReversed,omitempty" xml:"IsReversed,omitempty"`
	NextToken  *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	MaxResults *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
}

func (s ListEntityItemsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListEntityItemsRequest) GoString() string {
	return s.String()
}

func (s *ListEntityItemsRequest) SetWorkspace(v string) *ListEntityItemsRequest {
	s.Workspace = &v
	return s
}

func (s *ListEntityItemsRequest) SetEntityType(v string) *ListEntityItemsRequest {
	s.EntityType = &v
	return s
}

func (s *ListEntityItemsRequest) SetSearch(v string) *ListEntityItemsRequest {
	s.Search = &v
	return s
}

func (s *ListEntityItemsRequest) SetOrderBy(v string) *ListEntityItemsRequest {
	s.OrderBy = &v
	return s
}

func (s *ListEntityItemsRequest) SetIsReversed(v bool) *ListEntityItemsRequest {
	s.IsReversed = &v
	return s
}

func (s *ListEntityItemsRequest) SetNextToken(v string) *ListEntityItemsRequest {
	s.NextToken = &v
	return s
}

func (s *ListEntityItemsRequest) SetMaxResults(v int32) *ListEntityItemsRequest {
	s.MaxResults = &v
	return s
}

type ListEntityItemsResponseBody struct {
	// 主机ID
	HostId *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	// 请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 请求的最大结果数
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// 下次查询的起始Token
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// 返回总个数
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// 实体类型数组
	EntityItems []*ListEntityItemsResponseBodyEntityItems `json:"EntityItems,omitempty" xml:"EntityItems,omitempty" type:"Repeated"`
}

func (s ListEntityItemsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListEntityItemsResponseBody) GoString() string {
	return s.String()
}

func (s *ListEntityItemsResponseBody) SetHostId(v string) *ListEntityItemsResponseBody {
	s.HostId = &v
	return s
}

func (s *ListEntityItemsResponseBody) SetRequestId(v string) *ListEntityItemsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListEntityItemsResponseBody) SetMaxResults(v int32) *ListEntityItemsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListEntityItemsResponseBody) SetNextToken(v string) *ListEntityItemsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListEntityItemsResponseBody) SetTotalCount(v int32) *ListEntityItemsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListEntityItemsResponseBody) SetEntityItems(v []*ListEntityItemsResponseBodyEntityItems) *ListEntityItemsResponseBody {
	s.EntityItems = v
	return s
}

type ListEntityItemsResponseBodyEntityItems struct {
	EntityName *string            `json:"EntityName,omitempty" xml:"EntityName,omitempty"`
	EntityData map[string]*string `json:"EntityData,omitempty" xml:"EntityData,omitempty"`
}

func (s ListEntityItemsResponseBodyEntityItems) String() string {
	return tea.Prettify(s)
}

func (s ListEntityItemsResponseBodyEntityItems) GoString() string {
	return s.String()
}

func (s *ListEntityItemsResponseBodyEntityItems) SetEntityName(v string) *ListEntityItemsResponseBodyEntityItems {
	s.EntityName = &v
	return s
}

func (s *ListEntityItemsResponseBodyEntityItems) SetEntityData(v map[string]*string) *ListEntityItemsResponseBodyEntityItems {
	s.EntityData = v
	return s
}

type ListEntityItemsResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListEntityItemsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListEntityItemsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListEntityItemsResponse) GoString() string {
	return s.String()
}

func (s *ListEntityItemsResponse) SetHeaders(v map[string]*string) *ListEntityItemsResponse {
	s.Headers = v
	return s
}

func (s *ListEntityItemsResponse) SetBody(v *ListEntityItemsResponseBody) *ListEntityItemsResponse {
	s.Body = v
	return s
}

type ListGlobalAppsRequest struct {
	// 模糊查询字段：NamesapceName  AppName  Categories AppDescription
	Search *string `json:"Search,omitempty" xml:"Search,omitempty"`
	// 排序
	OrderBy *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	// 是否倒序，默认倒序排列
	IsReversed *bool `json:"IsReversed,omitempty" xml:"IsReversed,omitempty"`
	// 用来标记当前开始读取的位置，置空表示从头开始。
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// 一批返回的最大数据量
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// 可见范围
	AppScope *string `json:"AppScope,omitempty" xml:"AppScope,omitempty"`
	// 分类
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// 工具集
	Toolkit *string `json:"Toolkit,omitempty" xml:"Toolkit,omitempty"`
	// 区域Id
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s ListGlobalAppsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListGlobalAppsRequest) GoString() string {
	return s.String()
}

func (s *ListGlobalAppsRequest) SetSearch(v string) *ListGlobalAppsRequest {
	s.Search = &v
	return s
}

func (s *ListGlobalAppsRequest) SetOrderBy(v string) *ListGlobalAppsRequest {
	s.OrderBy = &v
	return s
}

func (s *ListGlobalAppsRequest) SetIsReversed(v bool) *ListGlobalAppsRequest {
	s.IsReversed = &v
	return s
}

func (s *ListGlobalAppsRequest) SetNextToken(v string) *ListGlobalAppsRequest {
	s.NextToken = &v
	return s
}

func (s *ListGlobalAppsRequest) SetMaxResults(v int32) *ListGlobalAppsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListGlobalAppsRequest) SetAppScope(v string) *ListGlobalAppsRequest {
	s.AppScope = &v
	return s
}

func (s *ListGlobalAppsRequest) SetCategory(v string) *ListGlobalAppsRequest {
	s.Category = &v
	return s
}

func (s *ListGlobalAppsRequest) SetToolkit(v string) *ListGlobalAppsRequest {
	s.Toolkit = &v
	return s
}

func (s *ListGlobalAppsRequest) SetRegion(v string) *ListGlobalAppsRequest {
	s.Region = &v
	return s
}

type ListGlobalAppsResponseBody struct {
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Message     *string `json:"Message,omitempty" xml:"Message,omitempty"`
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// 公共应用集合
	GlobalApps []*ListGlobalAppsResponseBodyGlobalApps `json:"GlobalApps,omitempty" xml:"GlobalApps,omitempty" type:"Repeated"`
	// 用来表示当前调用返回读取到的位置，空代表数据已经读取完毕
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// 本次请求所返回的最大记录条数
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// 本次请求条件下的数据总量
	TotalCount     *int64  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	HostId         *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
}

func (s ListGlobalAppsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListGlobalAppsResponseBody) GoString() string {
	return s.String()
}

func (s *ListGlobalAppsResponseBody) SetRequestId(v string) *ListGlobalAppsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListGlobalAppsResponseBody) SetMessage(v string) *ListGlobalAppsResponseBody {
	s.Message = &v
	return s
}

func (s *ListGlobalAppsResponseBody) SetDynamicCode(v string) *ListGlobalAppsResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *ListGlobalAppsResponseBody) SetGlobalApps(v []*ListGlobalAppsResponseBodyGlobalApps) *ListGlobalAppsResponseBody {
	s.GlobalApps = v
	return s
}

func (s *ListGlobalAppsResponseBody) SetNextToken(v string) *ListGlobalAppsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListGlobalAppsResponseBody) SetMaxResults(v int32) *ListGlobalAppsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListGlobalAppsResponseBody) SetTotalCount(v int64) *ListGlobalAppsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListGlobalAppsResponseBody) SetDynamicMessage(v string) *ListGlobalAppsResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *ListGlobalAppsResponseBody) SetCode(v string) *ListGlobalAppsResponseBody {
	s.Code = &v
	return s
}

func (s *ListGlobalAppsResponseBody) SetSuccess(v bool) *ListGlobalAppsResponseBody {
	s.Success = &v
	return s
}

func (s *ListGlobalAppsResponseBody) SetHostId(v string) *ListGlobalAppsResponseBody {
	s.HostId = &v
	return s
}

type ListGlobalAppsResponseBodyGlobalApps struct {
	// 命名空间名称
	NamespaceName *string `json:"NamespaceName,omitempty" xml:"NamespaceName,omitempty"`
	// 应用权限
	AppScope *string `json:"AppScope,omitempty" xml:"AppScope,omitempty"`
	// 应用名称
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// 应用描述
	AppDescription *string `json:"AppDescription,omitempty" xml:"AppDescription,omitempty"`
	// 应用工具合集
	Toolkit *string `json:"Toolkit,omitempty" xml:"Toolkit,omitempty"`
	// 应用计费说明
	AppFee *string `json:"AppFee,omitempty" xml:"AppFee,omitempty"`
	// 应用默认版本
	AppDefaultVersion *string `json:"AppDefaultVersion,omitempty" xml:"AppDefaultVersion,omitempty"`
	// 应用收藏置顶标记
	Pinned *bool `json:"Pinned,omitempty" xml:"Pinned,omitempty"`
	// 更新时间
	LastModified *string `json:"LastModified,omitempty" xml:"LastModified,omitempty"`
	// 应用支持的区域ids
	RegionIds []*string `json:"RegionIds,omitempty" xml:"RegionIds,omitempty" type:"Repeated"`
	// 应用所属分类
	Categories []*string `json:"Categories,omitempty" xml:"Categories,omitempty" type:"Repeated"`
}

func (s ListGlobalAppsResponseBodyGlobalApps) String() string {
	return tea.Prettify(s)
}

func (s ListGlobalAppsResponseBodyGlobalApps) GoString() string {
	return s.String()
}

func (s *ListGlobalAppsResponseBodyGlobalApps) SetNamespaceName(v string) *ListGlobalAppsResponseBodyGlobalApps {
	s.NamespaceName = &v
	return s
}

func (s *ListGlobalAppsResponseBodyGlobalApps) SetAppScope(v string) *ListGlobalAppsResponseBodyGlobalApps {
	s.AppScope = &v
	return s
}

func (s *ListGlobalAppsResponseBodyGlobalApps) SetAppName(v string) *ListGlobalAppsResponseBodyGlobalApps {
	s.AppName = &v
	return s
}

func (s *ListGlobalAppsResponseBodyGlobalApps) SetAppDescription(v string) *ListGlobalAppsResponseBodyGlobalApps {
	s.AppDescription = &v
	return s
}

func (s *ListGlobalAppsResponseBodyGlobalApps) SetToolkit(v string) *ListGlobalAppsResponseBodyGlobalApps {
	s.Toolkit = &v
	return s
}

func (s *ListGlobalAppsResponseBodyGlobalApps) SetAppFee(v string) *ListGlobalAppsResponseBodyGlobalApps {
	s.AppFee = &v
	return s
}

func (s *ListGlobalAppsResponseBodyGlobalApps) SetAppDefaultVersion(v string) *ListGlobalAppsResponseBodyGlobalApps {
	s.AppDefaultVersion = &v
	return s
}

func (s *ListGlobalAppsResponseBodyGlobalApps) SetPinned(v bool) *ListGlobalAppsResponseBodyGlobalApps {
	s.Pinned = &v
	return s
}

func (s *ListGlobalAppsResponseBodyGlobalApps) SetLastModified(v string) *ListGlobalAppsResponseBodyGlobalApps {
	s.LastModified = &v
	return s
}

func (s *ListGlobalAppsResponseBodyGlobalApps) SetRegionIds(v []*string) *ListGlobalAppsResponseBodyGlobalApps {
	s.RegionIds = v
	return s
}

func (s *ListGlobalAppsResponseBodyGlobalApps) SetCategories(v []*string) *ListGlobalAppsResponseBodyGlobalApps {
	s.Categories = v
	return s
}

type ListGlobalAppsResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListGlobalAppsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListGlobalAppsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListGlobalAppsResponse) GoString() string {
	return s.String()
}

func (s *ListGlobalAppsResponse) SetHeaders(v map[string]*string) *ListGlobalAppsResponse {
	s.Headers = v
	return s
}

func (s *ListGlobalAppsResponse) SetBody(v *ListGlobalAppsResponseBody) *ListGlobalAppsResponse {
	s.Body = v
	return s
}

type ListPublicDatasetRequest struct {
	// 名称、描述中搜索的关键字
	Search *string `json:"Search,omitempty" xml:"Search,omitempty"`
	// 公共数据集标签名
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	// 排序字段
	OrderBy *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	// 排序是否反转
	IsReversed *bool `json:"IsReversed,omitempty" xml:"IsReversed,omitempty"`
	// 翻页Token
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// 分页数量
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
}

func (s ListPublicDatasetRequest) String() string {
	return tea.Prettify(s)
}

func (s ListPublicDatasetRequest) GoString() string {
	return s.String()
}

func (s *ListPublicDatasetRequest) SetSearch(v string) *ListPublicDatasetRequest {
	s.Search = &v
	return s
}

func (s *ListPublicDatasetRequest) SetTag(v string) *ListPublicDatasetRequest {
	s.Tag = &v
	return s
}

func (s *ListPublicDatasetRequest) SetOrderBy(v string) *ListPublicDatasetRequest {
	s.OrderBy = &v
	return s
}

func (s *ListPublicDatasetRequest) SetIsReversed(v bool) *ListPublicDatasetRequest {
	s.IsReversed = &v
	return s
}

func (s *ListPublicDatasetRequest) SetNextToken(v string) *ListPublicDatasetRequest {
	s.NextToken = &v
	return s
}

func (s *ListPublicDatasetRequest) SetMaxResults(v int32) *ListPublicDatasetRequest {
	s.MaxResults = &v
	return s
}

type ListPublicDatasetResponseBody struct {
	// 主机ID
	HostId *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	// 请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	HttpCode  *int32  `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	// 错误码
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 错误消息
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 是否调用成功
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// 翻页Token
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// 分页数
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// 总记录数
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// 公共数据集信息
	Datasets []*ListPublicDatasetResponseBodyDatasets `json:"Datasets,omitempty" xml:"Datasets,omitempty" type:"Repeated"`
}

func (s ListPublicDatasetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListPublicDatasetResponseBody) GoString() string {
	return s.String()
}

func (s *ListPublicDatasetResponseBody) SetHostId(v string) *ListPublicDatasetResponseBody {
	s.HostId = &v
	return s
}

func (s *ListPublicDatasetResponseBody) SetRequestId(v string) *ListPublicDatasetResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPublicDatasetResponseBody) SetHttpCode(v int32) *ListPublicDatasetResponseBody {
	s.HttpCode = &v
	return s
}

func (s *ListPublicDatasetResponseBody) SetCode(v string) *ListPublicDatasetResponseBody {
	s.Code = &v
	return s
}

func (s *ListPublicDatasetResponseBody) SetMessage(v string) *ListPublicDatasetResponseBody {
	s.Message = &v
	return s
}

func (s *ListPublicDatasetResponseBody) SetSuccess(v bool) *ListPublicDatasetResponseBody {
	s.Success = &v
	return s
}

func (s *ListPublicDatasetResponseBody) SetNextToken(v string) *ListPublicDatasetResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListPublicDatasetResponseBody) SetMaxResults(v int32) *ListPublicDatasetResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListPublicDatasetResponseBody) SetTotalCount(v int32) *ListPublicDatasetResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListPublicDatasetResponseBody) SetDatasets(v []*ListPublicDatasetResponseBodyDatasets) *ListPublicDatasetResponseBody {
	s.Datasets = v
	return s
}

type ListPublicDatasetResponseBodyDatasets struct {
	// 最后更新时间
	LastModified *string `json:"LastModified,omitempty" xml:"LastModified,omitempty"`
	// 公共数据集版权信息
	Copyright *string `json:"Copyright,omitempty" xml:"Copyright,omitempty"`
	// 公共数据集访问要求
	AccessRequirements *string `json:"AccessRequirements,omitempty" xml:"AccessRequirements,omitempty"`
	// 关于公共数据集
	About *string `json:"About,omitempty" xml:"About,omitempty"`
	// 公共数据集标签
	Tags []*string `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// 公共数据集名称
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	// 公共数据集描述
	DatasetDescription *string `json:"DatasetDescription,omitempty" xml:"DatasetDescription,omitempty"`
	// 公共数据集更新频率
	UpdateFrequency *string `json:"UpdateFrequency,omitempty" xml:"UpdateFrequency,omitempty"`
	// 公共数据集可用区域
	RegionIds []*string `json:"RegionIds,omitempty" xml:"RegionIds,omitempty" type:"Repeated"`
}

func (s ListPublicDatasetResponseBodyDatasets) String() string {
	return tea.Prettify(s)
}

func (s ListPublicDatasetResponseBodyDatasets) GoString() string {
	return s.String()
}

func (s *ListPublicDatasetResponseBodyDatasets) SetLastModified(v string) *ListPublicDatasetResponseBodyDatasets {
	s.LastModified = &v
	return s
}

func (s *ListPublicDatasetResponseBodyDatasets) SetCopyright(v string) *ListPublicDatasetResponseBodyDatasets {
	s.Copyright = &v
	return s
}

func (s *ListPublicDatasetResponseBodyDatasets) SetAccessRequirements(v string) *ListPublicDatasetResponseBodyDatasets {
	s.AccessRequirements = &v
	return s
}

func (s *ListPublicDatasetResponseBodyDatasets) SetAbout(v string) *ListPublicDatasetResponseBodyDatasets {
	s.About = &v
	return s
}

func (s *ListPublicDatasetResponseBodyDatasets) SetTags(v []*string) *ListPublicDatasetResponseBodyDatasets {
	s.Tags = v
	return s
}

func (s *ListPublicDatasetResponseBodyDatasets) SetDatasetName(v string) *ListPublicDatasetResponseBodyDatasets {
	s.DatasetName = &v
	return s
}

func (s *ListPublicDatasetResponseBodyDatasets) SetDatasetDescription(v string) *ListPublicDatasetResponseBodyDatasets {
	s.DatasetDescription = &v
	return s
}

func (s *ListPublicDatasetResponseBodyDatasets) SetUpdateFrequency(v string) *ListPublicDatasetResponseBodyDatasets {
	s.UpdateFrequency = &v
	return s
}

func (s *ListPublicDatasetResponseBodyDatasets) SetRegionIds(v []*string) *ListPublicDatasetResponseBodyDatasets {
	s.RegionIds = v
	return s
}

type ListPublicDatasetResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListPublicDatasetResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListPublicDatasetResponse) String() string {
	return tea.Prettify(s)
}

func (s ListPublicDatasetResponse) GoString() string {
	return s.String()
}

func (s *ListPublicDatasetResponse) SetHeaders(v map[string]*string) *ListPublicDatasetResponse {
	s.Headers = v
	return s
}

func (s *ListPublicDatasetResponse) SetBody(v *ListPublicDatasetResponseBody) *ListPublicDatasetResponse {
	s.Body = v
	return s
}

type ListPublicDatasetEntitiesRequest struct {
	// 数据集名称
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	// 排序字段
	OrderBy *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	// 排序是否反转
	IsReversed *bool `json:"IsReversed,omitempty" xml:"IsReversed,omitempty"`
	// 翻页Token
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// 分页数量
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
}

func (s ListPublicDatasetEntitiesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListPublicDatasetEntitiesRequest) GoString() string {
	return s.String()
}

func (s *ListPublicDatasetEntitiesRequest) SetDatasetName(v string) *ListPublicDatasetEntitiesRequest {
	s.DatasetName = &v
	return s
}

func (s *ListPublicDatasetEntitiesRequest) SetOrderBy(v string) *ListPublicDatasetEntitiesRequest {
	s.OrderBy = &v
	return s
}

func (s *ListPublicDatasetEntitiesRequest) SetIsReversed(v bool) *ListPublicDatasetEntitiesRequest {
	s.IsReversed = &v
	return s
}

func (s *ListPublicDatasetEntitiesRequest) SetNextToken(v string) *ListPublicDatasetEntitiesRequest {
	s.NextToken = &v
	return s
}

func (s *ListPublicDatasetEntitiesRequest) SetMaxResults(v int32) *ListPublicDatasetEntitiesRequest {
	s.MaxResults = &v
	return s
}

type ListPublicDatasetEntitiesResponseBody struct {
	// 主机ID
	HostId *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	// 请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	HttpCode  *int32  `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	// 错误码
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 错误消息
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	DynamicCode    *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// 翻页Token
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// 分页数
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// 总记录数
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// 公共数据集名称
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	// 该实体包含的所有类型
	Entities []*ListPublicDatasetEntitiesResponseBodyEntities `json:"Entities,omitempty" xml:"Entities,omitempty" type:"Repeated"`
}

func (s ListPublicDatasetEntitiesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListPublicDatasetEntitiesResponseBody) GoString() string {
	return s.String()
}

func (s *ListPublicDatasetEntitiesResponseBody) SetHostId(v string) *ListPublicDatasetEntitiesResponseBody {
	s.HostId = &v
	return s
}

func (s *ListPublicDatasetEntitiesResponseBody) SetRequestId(v string) *ListPublicDatasetEntitiesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPublicDatasetEntitiesResponseBody) SetHttpCode(v int32) *ListPublicDatasetEntitiesResponseBody {
	s.HttpCode = &v
	return s
}

func (s *ListPublicDatasetEntitiesResponseBody) SetCode(v string) *ListPublicDatasetEntitiesResponseBody {
	s.Code = &v
	return s
}

func (s *ListPublicDatasetEntitiesResponseBody) SetMessage(v string) *ListPublicDatasetEntitiesResponseBody {
	s.Message = &v
	return s
}

func (s *ListPublicDatasetEntitiesResponseBody) SetDynamicCode(v string) *ListPublicDatasetEntitiesResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *ListPublicDatasetEntitiesResponseBody) SetDynamicMessage(v string) *ListPublicDatasetEntitiesResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *ListPublicDatasetEntitiesResponseBody) SetNextToken(v string) *ListPublicDatasetEntitiesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListPublicDatasetEntitiesResponseBody) SetMaxResults(v int32) *ListPublicDatasetEntitiesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListPublicDatasetEntitiesResponseBody) SetTotalCount(v int32) *ListPublicDatasetEntitiesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListPublicDatasetEntitiesResponseBody) SetDatasetName(v string) *ListPublicDatasetEntitiesResponseBody {
	s.DatasetName = &v
	return s
}

func (s *ListPublicDatasetEntitiesResponseBody) SetEntities(v []*ListPublicDatasetEntitiesResponseBodyEntities) *ListPublicDatasetEntitiesResponseBody {
	s.Entities = v
	return s
}

type ListPublicDatasetEntitiesResponseBodyEntities struct {
	// 实体类型
	EntityType *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
}

func (s ListPublicDatasetEntitiesResponseBodyEntities) String() string {
	return tea.Prettify(s)
}

func (s ListPublicDatasetEntitiesResponseBodyEntities) GoString() string {
	return s.String()
}

func (s *ListPublicDatasetEntitiesResponseBodyEntities) SetEntityType(v string) *ListPublicDatasetEntitiesResponseBodyEntities {
	s.EntityType = &v
	return s
}

type ListPublicDatasetEntitiesResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListPublicDatasetEntitiesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListPublicDatasetEntitiesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListPublicDatasetEntitiesResponse) GoString() string {
	return s.String()
}

func (s *ListPublicDatasetEntitiesResponse) SetHeaders(v map[string]*string) *ListPublicDatasetEntitiesResponse {
	s.Headers = v
	return s
}

func (s *ListPublicDatasetEntitiesResponse) SetBody(v *ListPublicDatasetEntitiesResponseBody) *ListPublicDatasetEntitiesResponse {
	s.Body = v
	return s
}

type ListPublicDatasetEntityItemsRequest struct {
	// 数据集名称
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	// 实体类型
	EntityType *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	// 实体名中搜索的关键字
	Search *string `json:"Search,omitempty" xml:"Search,omitempty"`
	// 排序字段
	OrderBy *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	// 是否反转
	IsReversed *bool `json:"IsReversed,omitempty" xml:"IsReversed,omitempty"`
	// 翻页Token
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// 分页数量
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
}

func (s ListPublicDatasetEntityItemsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListPublicDatasetEntityItemsRequest) GoString() string {
	return s.String()
}

func (s *ListPublicDatasetEntityItemsRequest) SetDatasetName(v string) *ListPublicDatasetEntityItemsRequest {
	s.DatasetName = &v
	return s
}

func (s *ListPublicDatasetEntityItemsRequest) SetEntityType(v string) *ListPublicDatasetEntityItemsRequest {
	s.EntityType = &v
	return s
}

func (s *ListPublicDatasetEntityItemsRequest) SetSearch(v string) *ListPublicDatasetEntityItemsRequest {
	s.Search = &v
	return s
}

func (s *ListPublicDatasetEntityItemsRequest) SetOrderBy(v string) *ListPublicDatasetEntityItemsRequest {
	s.OrderBy = &v
	return s
}

func (s *ListPublicDatasetEntityItemsRequest) SetIsReversed(v bool) *ListPublicDatasetEntityItemsRequest {
	s.IsReversed = &v
	return s
}

func (s *ListPublicDatasetEntityItemsRequest) SetNextToken(v string) *ListPublicDatasetEntityItemsRequest {
	s.NextToken = &v
	return s
}

func (s *ListPublicDatasetEntityItemsRequest) SetMaxResults(v int32) *ListPublicDatasetEntityItemsRequest {
	s.MaxResults = &v
	return s
}

type ListPublicDatasetEntityItemsResponseBody struct {
	// 主机ID
	HostId *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	// 请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	HttpCode  *int32  `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	// 错误码
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 错误消息
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	DynamicCode    *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// 翻页Token
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// 分页数
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// 总记录数
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// 公共数据集名称
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	// 该实体包含的所有对象
	EntityItems []*ListPublicDatasetEntityItemsResponseBodyEntityItems `json:"EntityItems,omitempty" xml:"EntityItems,omitempty" type:"Repeated"`
}

func (s ListPublicDatasetEntityItemsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListPublicDatasetEntityItemsResponseBody) GoString() string {
	return s.String()
}

func (s *ListPublicDatasetEntityItemsResponseBody) SetHostId(v string) *ListPublicDatasetEntityItemsResponseBody {
	s.HostId = &v
	return s
}

func (s *ListPublicDatasetEntityItemsResponseBody) SetRequestId(v string) *ListPublicDatasetEntityItemsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPublicDatasetEntityItemsResponseBody) SetHttpCode(v int32) *ListPublicDatasetEntityItemsResponseBody {
	s.HttpCode = &v
	return s
}

func (s *ListPublicDatasetEntityItemsResponseBody) SetCode(v string) *ListPublicDatasetEntityItemsResponseBody {
	s.Code = &v
	return s
}

func (s *ListPublicDatasetEntityItemsResponseBody) SetMessage(v string) *ListPublicDatasetEntityItemsResponseBody {
	s.Message = &v
	return s
}

func (s *ListPublicDatasetEntityItemsResponseBody) SetDynamicCode(v string) *ListPublicDatasetEntityItemsResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *ListPublicDatasetEntityItemsResponseBody) SetDynamicMessage(v string) *ListPublicDatasetEntityItemsResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *ListPublicDatasetEntityItemsResponseBody) SetNextToken(v string) *ListPublicDatasetEntityItemsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListPublicDatasetEntityItemsResponseBody) SetMaxResults(v int32) *ListPublicDatasetEntityItemsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListPublicDatasetEntityItemsResponseBody) SetTotalCount(v int32) *ListPublicDatasetEntityItemsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListPublicDatasetEntityItemsResponseBody) SetDatasetName(v string) *ListPublicDatasetEntityItemsResponseBody {
	s.DatasetName = &v
	return s
}

func (s *ListPublicDatasetEntityItemsResponseBody) SetEntityItems(v []*ListPublicDatasetEntityItemsResponseBodyEntityItems) *ListPublicDatasetEntityItemsResponseBody {
	s.EntityItems = v
	return s
}

type ListPublicDatasetEntityItemsResponseBodyEntityItems struct {
	// 实体名称
	EntityName *string `json:"EntityName,omitempty" xml:"EntityName,omitempty"`
	// 实体属性值
	EntityData map[string]*string `json:"EntityData,omitempty" xml:"EntityData,omitempty"`
}

func (s ListPublicDatasetEntityItemsResponseBodyEntityItems) String() string {
	return tea.Prettify(s)
}

func (s ListPublicDatasetEntityItemsResponseBodyEntityItems) GoString() string {
	return s.String()
}

func (s *ListPublicDatasetEntityItemsResponseBodyEntityItems) SetEntityName(v string) *ListPublicDatasetEntityItemsResponseBodyEntityItems {
	s.EntityName = &v
	return s
}

func (s *ListPublicDatasetEntityItemsResponseBodyEntityItems) SetEntityData(v map[string]*string) *ListPublicDatasetEntityItemsResponseBodyEntityItems {
	s.EntityData = v
	return s
}

type ListPublicDatasetEntityItemsResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListPublicDatasetEntityItemsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListPublicDatasetEntityItemsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListPublicDatasetEntityItemsResponse) GoString() string {
	return s.String()
}

func (s *ListPublicDatasetEntityItemsResponse) SetHeaders(v map[string]*string) *ListPublicDatasetEntityItemsResponse {
	s.Headers = v
	return s
}

func (s *ListPublicDatasetEntityItemsResponse) SetBody(v *ListPublicDatasetEntityItemsResponseBody) *ListPublicDatasetEntityItemsResponse {
	s.Body = v
	return s
}

type ListPublicDatasetTagsRequest struct {
	// 标签名中搜索的关键字
	Search *string `json:"Search,omitempty" xml:"Search,omitempty"`
	// 排序字段
	OrderBy *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	// 是否反转
	IsReversed *bool `json:"IsReversed,omitempty" xml:"IsReversed,omitempty"`
	// 翻页Token
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// 分页数量
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
}

func (s ListPublicDatasetTagsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListPublicDatasetTagsRequest) GoString() string {
	return s.String()
}

func (s *ListPublicDatasetTagsRequest) SetSearch(v string) *ListPublicDatasetTagsRequest {
	s.Search = &v
	return s
}

func (s *ListPublicDatasetTagsRequest) SetOrderBy(v string) *ListPublicDatasetTagsRequest {
	s.OrderBy = &v
	return s
}

func (s *ListPublicDatasetTagsRequest) SetIsReversed(v bool) *ListPublicDatasetTagsRequest {
	s.IsReversed = &v
	return s
}

func (s *ListPublicDatasetTagsRequest) SetNextToken(v string) *ListPublicDatasetTagsRequest {
	s.NextToken = &v
	return s
}

func (s *ListPublicDatasetTagsRequest) SetMaxResults(v int32) *ListPublicDatasetTagsRequest {
	s.MaxResults = &v
	return s
}

type ListPublicDatasetTagsResponseBody struct {
	// 主机ID
	HostId *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	// 请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	HttpCode  *int32  `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	// 错误码
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 错误消息
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 是否调用成功
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// 翻页Token
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// 分页数
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// 总记录数
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// 公共数据集标签
	Tags []*string `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s ListPublicDatasetTagsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListPublicDatasetTagsResponseBody) GoString() string {
	return s.String()
}

func (s *ListPublicDatasetTagsResponseBody) SetHostId(v string) *ListPublicDatasetTagsResponseBody {
	s.HostId = &v
	return s
}

func (s *ListPublicDatasetTagsResponseBody) SetRequestId(v string) *ListPublicDatasetTagsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPublicDatasetTagsResponseBody) SetHttpCode(v int32) *ListPublicDatasetTagsResponseBody {
	s.HttpCode = &v
	return s
}

func (s *ListPublicDatasetTagsResponseBody) SetCode(v string) *ListPublicDatasetTagsResponseBody {
	s.Code = &v
	return s
}

func (s *ListPublicDatasetTagsResponseBody) SetMessage(v string) *ListPublicDatasetTagsResponseBody {
	s.Message = &v
	return s
}

func (s *ListPublicDatasetTagsResponseBody) SetSuccess(v bool) *ListPublicDatasetTagsResponseBody {
	s.Success = &v
	return s
}

func (s *ListPublicDatasetTagsResponseBody) SetNextToken(v string) *ListPublicDatasetTagsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListPublicDatasetTagsResponseBody) SetMaxResults(v int32) *ListPublicDatasetTagsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListPublicDatasetTagsResponseBody) SetTotalCount(v int32) *ListPublicDatasetTagsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListPublicDatasetTagsResponseBody) SetTags(v []*string) *ListPublicDatasetTagsResponseBody {
	s.Tags = v
	return s
}

type ListPublicDatasetTagsResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListPublicDatasetTagsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListPublicDatasetTagsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListPublicDatasetTagsResponse) GoString() string {
	return s.String()
}

func (s *ListPublicDatasetTagsResponse) SetHeaders(v map[string]*string) *ListPublicDatasetTagsResponse {
	s.Headers = v
	return s
}

func (s *ListPublicDatasetTagsResponse) SetBody(v *ListPublicDatasetTagsResponseBody) *ListPublicDatasetTagsResponse {
	s.Body = v
	return s
}

type ListRegionsResponseBody struct {
	// 主机ID
	HostId *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	// 请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 错误码
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 错误消息
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 是否调用成功
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// 基因云产品上线区域
	Regions []*ListRegionsResponseBodyRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Repeated"`
}

func (s ListRegionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListRegionsResponseBody) SetHostId(v string) *ListRegionsResponseBody {
	s.HostId = &v
	return s
}

func (s *ListRegionsResponseBody) SetRequestId(v string) *ListRegionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRegionsResponseBody) SetCode(v string) *ListRegionsResponseBody {
	s.Code = &v
	return s
}

func (s *ListRegionsResponseBody) SetMessage(v string) *ListRegionsResponseBody {
	s.Message = &v
	return s
}

func (s *ListRegionsResponseBody) SetSuccess(v bool) *ListRegionsResponseBody {
	s.Success = &v
	return s
}

func (s *ListRegionsResponseBody) SetRegions(v []*ListRegionsResponseBodyRegions) *ListRegionsResponseBody {
	s.Regions = v
	return s
}

type ListRegionsResponseBodyRegions struct {
	LocalName      *string `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	RegionEndpoint *string `json:"RegionEndpoint,omitempty" xml:"RegionEndpoint,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListRegionsResponseBodyRegions) String() string {
	return tea.Prettify(s)
}

func (s ListRegionsResponseBodyRegions) GoString() string {
	return s.String()
}

func (s *ListRegionsResponseBodyRegions) SetLocalName(v string) *ListRegionsResponseBodyRegions {
	s.LocalName = &v
	return s
}

func (s *ListRegionsResponseBodyRegions) SetRegionEndpoint(v string) *ListRegionsResponseBodyRegions {
	s.RegionEndpoint = &v
	return s
}

func (s *ListRegionsResponseBodyRegions) SetRegionId(v string) *ListRegionsResponseBodyRegions {
	s.RegionId = &v
	return s
}

type ListRegionsResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListRegionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListRegionsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListRegionsResponse) GoString() string {
	return s.String()
}

func (s *ListRegionsResponse) SetHeaders(v map[string]*string) *ListRegionsResponse {
	s.Headers = v
	return s
}

func (s *ListRegionsResponse) SetBody(v *ListRegionsResponseBody) *ListRegionsResponse {
	s.Body = v
	return s
}

type ListRunsRequest struct {
	// 工作空间
	Workspace *string `json:"Workspace,omitempty" xml:"Workspace,omitempty"`
	// 任务名字
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 状态
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 提交ID
	SubmissionId *string `json:"SubmissionId,omitempty" xml:"SubmissionId,omitempty"`
	// 搜索ID
	Search *string `json:"Search,omitempty" xml:"Search,omitempty"`
	// 标签选择
	LabelSelector *string `json:"LabelSelector,omitempty" xml:"LabelSelector,omitempty"`
	// 应用名称
	AppName    *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	NextToken  *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	MaxResults *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	OrderBy    *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	IsReversed *bool   `json:"IsReversed,omitempty" xml:"IsReversed,omitempty"`
	GetTotal   *bool   `json:"GetTotal,omitempty" xml:"GetTotal,omitempty"`
}

func (s ListRunsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListRunsRequest) GoString() string {
	return s.String()
}

func (s *ListRunsRequest) SetWorkspace(v string) *ListRunsRequest {
	s.Workspace = &v
	return s
}

func (s *ListRunsRequest) SetName(v string) *ListRunsRequest {
	s.Name = &v
	return s
}

func (s *ListRunsRequest) SetStatus(v string) *ListRunsRequest {
	s.Status = &v
	return s
}

func (s *ListRunsRequest) SetSubmissionId(v string) *ListRunsRequest {
	s.SubmissionId = &v
	return s
}

func (s *ListRunsRequest) SetSearch(v string) *ListRunsRequest {
	s.Search = &v
	return s
}

func (s *ListRunsRequest) SetLabelSelector(v string) *ListRunsRequest {
	s.LabelSelector = &v
	return s
}

func (s *ListRunsRequest) SetAppName(v string) *ListRunsRequest {
	s.AppName = &v
	return s
}

func (s *ListRunsRequest) SetNextToken(v string) *ListRunsRequest {
	s.NextToken = &v
	return s
}

func (s *ListRunsRequest) SetMaxResults(v int32) *ListRunsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListRunsRequest) SetOrderBy(v string) *ListRunsRequest {
	s.OrderBy = &v
	return s
}

func (s *ListRunsRequest) SetIsReversed(v bool) *ListRunsRequest {
	s.IsReversed = &v
	return s
}

func (s *ListRunsRequest) SetGetTotal(v bool) *ListRunsRequest {
	s.GetTotal = &v
	return s
}

type ListRunsResponseBody struct {
	// 请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 主机ID
	HostId *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	// 最大返回结果
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// 下次查询Token
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// 返回个数
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// 任务列表
	Runs []*ListRunsResponseBodyRuns `json:"Runs,omitempty" xml:"Runs,omitempty" type:"Repeated"`
}

func (s ListRunsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListRunsResponseBody) GoString() string {
	return s.String()
}

func (s *ListRunsResponseBody) SetRequestId(v string) *ListRunsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRunsResponseBody) SetHostId(v string) *ListRunsResponseBody {
	s.HostId = &v
	return s
}

func (s *ListRunsResponseBody) SetMaxResults(v int32) *ListRunsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListRunsResponseBody) SetNextToken(v string) *ListRunsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListRunsResponseBody) SetTotalCount(v int32) *ListRunsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListRunsResponseBody) SetRuns(v []*ListRunsResponseBodyRuns) *ListRunsResponseBody {
	s.Runs = v
	return s
}

type ListRunsResponseBodyRuns struct {
	// 应用名称
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// 应用来源
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// 应用命名空间
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// 应用原名
	AppOrigName *string `json:"AppOrigName,omitempty" xml:"AppOrigName,omitempty"`
	// 任务标签
	Labels map[string]*string `json:"Labels,omitempty" xml:"Labels,omitempty"`
	// 任务名称
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 应用版本号
	Revision *string `json:"Revision,omitempty" xml:"Revision,omitempty"`
	// 任务ID
	RunId *string `json:"RunId,omitempty" xml:"RunId,omitempty"`
	// 任务状态
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 提交时间
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 开始时间
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// 结束时间
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// 提交ID
	SubmissionId *string `json:"SubmissionId,omitempty" xml:"SubmissionId,omitempty"`
	// 实体名称
	EntityName *string `json:"EntityName,omitempty" xml:"EntityName,omitempty"`
	// 实体对象ID
	EntityType *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	// 运行目录
	ExecuteDirectory *string                                 `json:"ExecuteDirectory,omitempty" xml:"ExecuteDirectory,omitempty"`
	ExecuteOptions   *ListRunsResponseBodyRunsExecuteOptions `json:"ExecuteOptions,omitempty" xml:"ExecuteOptions,omitempty" type:"Struct"`
	// 输入参数
	Inputs         *string `json:"Inputs,omitempty" xml:"Inputs,omitempty"`
	DefaultRuntime *string `json:"DefaultRuntime,omitempty" xml:"DefaultRuntime,omitempty"`
	Workspace      *string `json:"Workspace,omitempty" xml:"Workspace,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListRunsResponseBodyRuns) String() string {
	return tea.Prettify(s)
}

func (s ListRunsResponseBodyRuns) GoString() string {
	return s.String()
}

func (s *ListRunsResponseBodyRuns) SetAppName(v string) *ListRunsResponseBodyRuns {
	s.AppName = &v
	return s
}

func (s *ListRunsResponseBodyRuns) SetSource(v string) *ListRunsResponseBodyRuns {
	s.Source = &v
	return s
}

func (s *ListRunsResponseBodyRuns) SetNamespace(v string) *ListRunsResponseBodyRuns {
	s.Namespace = &v
	return s
}

func (s *ListRunsResponseBodyRuns) SetAppOrigName(v string) *ListRunsResponseBodyRuns {
	s.AppOrigName = &v
	return s
}

func (s *ListRunsResponseBodyRuns) SetLabels(v map[string]*string) *ListRunsResponseBodyRuns {
	s.Labels = v
	return s
}

func (s *ListRunsResponseBodyRuns) SetName(v string) *ListRunsResponseBodyRuns {
	s.Name = &v
	return s
}

func (s *ListRunsResponseBodyRuns) SetRevision(v string) *ListRunsResponseBodyRuns {
	s.Revision = &v
	return s
}

func (s *ListRunsResponseBodyRuns) SetRunId(v string) *ListRunsResponseBodyRuns {
	s.RunId = &v
	return s
}

func (s *ListRunsResponseBodyRuns) SetStatus(v string) *ListRunsResponseBodyRuns {
	s.Status = &v
	return s
}

func (s *ListRunsResponseBodyRuns) SetCreateTime(v string) *ListRunsResponseBodyRuns {
	s.CreateTime = &v
	return s
}

func (s *ListRunsResponseBodyRuns) SetStartTime(v string) *ListRunsResponseBodyRuns {
	s.StartTime = &v
	return s
}

func (s *ListRunsResponseBodyRuns) SetEndTime(v string) *ListRunsResponseBodyRuns {
	s.EndTime = &v
	return s
}

func (s *ListRunsResponseBodyRuns) SetSubmissionId(v string) *ListRunsResponseBodyRuns {
	s.SubmissionId = &v
	return s
}

func (s *ListRunsResponseBodyRuns) SetEntityName(v string) *ListRunsResponseBodyRuns {
	s.EntityName = &v
	return s
}

func (s *ListRunsResponseBodyRuns) SetEntityType(v string) *ListRunsResponseBodyRuns {
	s.EntityType = &v
	return s
}

func (s *ListRunsResponseBodyRuns) SetExecuteDirectory(v string) *ListRunsResponseBodyRuns {
	s.ExecuteDirectory = &v
	return s
}

func (s *ListRunsResponseBodyRuns) SetExecuteOptions(v *ListRunsResponseBodyRunsExecuteOptions) *ListRunsResponseBodyRuns {
	s.ExecuteOptions = v
	return s
}

func (s *ListRunsResponseBodyRuns) SetInputs(v string) *ListRunsResponseBodyRuns {
	s.Inputs = &v
	return s
}

func (s *ListRunsResponseBodyRuns) SetDefaultRuntime(v string) *ListRunsResponseBodyRuns {
	s.DefaultRuntime = &v
	return s
}

func (s *ListRunsResponseBodyRuns) SetWorkspace(v string) *ListRunsResponseBodyRuns {
	s.Workspace = &v
	return s
}

func (s *ListRunsResponseBodyRuns) SetRegionId(v string) *ListRunsResponseBodyRuns {
	s.RegionId = &v
	return s
}

type ListRunsResponseBodyRunsExecuteOptions struct {
	CallCaching               *bool   `json:"CallCaching,omitempty" xml:"CallCaching,omitempty"`
	DeleteIntermediateResults *bool   `json:"DeleteIntermediateResults,omitempty" xml:"DeleteIntermediateResults,omitempty"`
	FailureMode               *string `json:"FailureMode,omitempty" xml:"FailureMode,omitempty"`
}

func (s ListRunsResponseBodyRunsExecuteOptions) String() string {
	return tea.Prettify(s)
}

func (s ListRunsResponseBodyRunsExecuteOptions) GoString() string {
	return s.String()
}

func (s *ListRunsResponseBodyRunsExecuteOptions) SetCallCaching(v bool) *ListRunsResponseBodyRunsExecuteOptions {
	s.CallCaching = &v
	return s
}

func (s *ListRunsResponseBodyRunsExecuteOptions) SetDeleteIntermediateResults(v bool) *ListRunsResponseBodyRunsExecuteOptions {
	s.DeleteIntermediateResults = &v
	return s
}

func (s *ListRunsResponseBodyRunsExecuteOptions) SetFailureMode(v string) *ListRunsResponseBodyRunsExecuteOptions {
	s.FailureMode = &v
	return s
}

type ListRunsResponse struct {
	Headers map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListRunsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListRunsResponse) SetBody(v *ListRunsResponseBody) *ListRunsResponse {
	s.Body = v
	return s
}

type ListSubmissionsRequest struct {
	// 工作空间
	Workspace *string `json:"Workspace,omitempty" xml:"Workspace,omitempty"`
	// 状态
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 搜索ID
	Search     *string `json:"Search,omitempty" xml:"Search,omitempty"`
	NextToken  *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	MaxResults *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	OrderBy    *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	IsReversed *bool   `json:"IsReversed,omitempty" xml:"IsReversed,omitempty"`
}

func (s ListSubmissionsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSubmissionsRequest) GoString() string {
	return s.String()
}

func (s *ListSubmissionsRequest) SetWorkspace(v string) *ListSubmissionsRequest {
	s.Workspace = &v
	return s
}

func (s *ListSubmissionsRequest) SetStatus(v string) *ListSubmissionsRequest {
	s.Status = &v
	return s
}

func (s *ListSubmissionsRequest) SetSearch(v string) *ListSubmissionsRequest {
	s.Search = &v
	return s
}

func (s *ListSubmissionsRequest) SetNextToken(v string) *ListSubmissionsRequest {
	s.NextToken = &v
	return s
}

func (s *ListSubmissionsRequest) SetMaxResults(v int32) *ListSubmissionsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListSubmissionsRequest) SetOrderBy(v string) *ListSubmissionsRequest {
	s.OrderBy = &v
	return s
}

func (s *ListSubmissionsRequest) SetIsReversed(v bool) *ListSubmissionsRequest {
	s.IsReversed = &v
	return s
}

type ListSubmissionsResponseBody struct {
	// 请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 主机ID
	HostId *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	// 最大返回结果
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// 下次查询Token
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// 返回个数
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// 投递列表
	Submissions []*ListSubmissionsResponseBodySubmissions `json:"Submissions,omitempty" xml:"Submissions,omitempty" type:"Repeated"`
}

func (s ListSubmissionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSubmissionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListSubmissionsResponseBody) SetRequestId(v string) *ListSubmissionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSubmissionsResponseBody) SetHostId(v string) *ListSubmissionsResponseBody {
	s.HostId = &v
	return s
}

func (s *ListSubmissionsResponseBody) SetMaxResults(v int32) *ListSubmissionsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListSubmissionsResponseBody) SetNextToken(v string) *ListSubmissionsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListSubmissionsResponseBody) SetTotalCount(v int32) *ListSubmissionsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListSubmissionsResponseBody) SetSubmissions(v []*ListSubmissionsResponseBodySubmissions) *ListSubmissionsResponseBody {
	s.Submissions = v
	return s
}

type ListSubmissionsResponseBodySubmissions struct {
	// 提交ID
	Workspace *string `json:"Workspace,omitempty" xml:"Workspace,omitempty"`
	// 提交ID
	SubmissionId *string `json:"SubmissionId,omitempty" xml:"SubmissionId,omitempty"`
	// 任务状态
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 提交时间
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 开始时间
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// 结束时间
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// 实体类型
	EntityType *string                                         `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	RunStats   *ListSubmissionsResponseBodySubmissionsRunStats `json:"RunStats,omitempty" xml:"RunStats,omitempty" type:"Struct"`
}

func (s ListSubmissionsResponseBodySubmissions) String() string {
	return tea.Prettify(s)
}

func (s ListSubmissionsResponseBodySubmissions) GoString() string {
	return s.String()
}

func (s *ListSubmissionsResponseBodySubmissions) SetWorkspace(v string) *ListSubmissionsResponseBodySubmissions {
	s.Workspace = &v
	return s
}

func (s *ListSubmissionsResponseBodySubmissions) SetSubmissionId(v string) *ListSubmissionsResponseBodySubmissions {
	s.SubmissionId = &v
	return s
}

func (s *ListSubmissionsResponseBodySubmissions) SetStatus(v string) *ListSubmissionsResponseBodySubmissions {
	s.Status = &v
	return s
}

func (s *ListSubmissionsResponseBodySubmissions) SetCreateTime(v string) *ListSubmissionsResponseBodySubmissions {
	s.CreateTime = &v
	return s
}

func (s *ListSubmissionsResponseBodySubmissions) SetStartTime(v string) *ListSubmissionsResponseBodySubmissions {
	s.StartTime = &v
	return s
}

func (s *ListSubmissionsResponseBodySubmissions) SetEndTime(v string) *ListSubmissionsResponseBodySubmissions {
	s.EndTime = &v
	return s
}

func (s *ListSubmissionsResponseBodySubmissions) SetEntityType(v string) *ListSubmissionsResponseBodySubmissions {
	s.EntityType = &v
	return s
}

func (s *ListSubmissionsResponseBodySubmissions) SetRunStats(v *ListSubmissionsResponseBodySubmissionsRunStats) *ListSubmissionsResponseBodySubmissions {
	s.RunStats = v
	return s
}

type ListSubmissionsResponseBodySubmissionsRunStats struct {
	Submitted *int64 `json:"Submitted,omitempty" xml:"Submitted,omitempty"`
	Pending   *int64 `json:"Pending,omitempty" xml:"Pending,omitempty"`
	Running   *int64 `json:"Running,omitempty" xml:"Running,omitempty"`
	Succeeded *int64 `json:"Succeeded,omitempty" xml:"Succeeded,omitempty"`
	Failed    *int64 `json:"Failed,omitempty" xml:"Failed,omitempty"`
	Aborting  *int64 `json:"Aborting,omitempty" xml:"Aborting,omitempty"`
	Aborted   *int64 `json:"Aborted,omitempty" xml:"Aborted,omitempty"`
}

func (s ListSubmissionsResponseBodySubmissionsRunStats) String() string {
	return tea.Prettify(s)
}

func (s ListSubmissionsResponseBodySubmissionsRunStats) GoString() string {
	return s.String()
}

func (s *ListSubmissionsResponseBodySubmissionsRunStats) SetSubmitted(v int64) *ListSubmissionsResponseBodySubmissionsRunStats {
	s.Submitted = &v
	return s
}

func (s *ListSubmissionsResponseBodySubmissionsRunStats) SetPending(v int64) *ListSubmissionsResponseBodySubmissionsRunStats {
	s.Pending = &v
	return s
}

func (s *ListSubmissionsResponseBodySubmissionsRunStats) SetRunning(v int64) *ListSubmissionsResponseBodySubmissionsRunStats {
	s.Running = &v
	return s
}

func (s *ListSubmissionsResponseBodySubmissionsRunStats) SetSucceeded(v int64) *ListSubmissionsResponseBodySubmissionsRunStats {
	s.Succeeded = &v
	return s
}

func (s *ListSubmissionsResponseBodySubmissionsRunStats) SetFailed(v int64) *ListSubmissionsResponseBodySubmissionsRunStats {
	s.Failed = &v
	return s
}

func (s *ListSubmissionsResponseBodySubmissionsRunStats) SetAborting(v int64) *ListSubmissionsResponseBodySubmissionsRunStats {
	s.Aborting = &v
	return s
}

func (s *ListSubmissionsResponseBodySubmissionsRunStats) SetAborted(v int64) *ListSubmissionsResponseBodySubmissionsRunStats {
	s.Aborted = &v
	return s
}

type ListSubmissionsResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListSubmissionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListSubmissionsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSubmissionsResponse) GoString() string {
	return s.String()
}

func (s *ListSubmissionsResponse) SetHeaders(v map[string]*string) *ListSubmissionsResponse {
	s.Headers = v
	return s
}

func (s *ListSubmissionsResponse) SetBody(v *ListSubmissionsResponseBody) *ListSubmissionsResponse {
	s.Body = v
	return s
}

type ListTemplatesRequest struct {
	// 工作空间
	Workspace *string `json:"Workspace,omitempty" xml:"Workspace,omitempty"`
	// Next Token
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// 最大返回结果数
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// 排序依据
	OrderBy *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	// 是否逆序
	IsReversed *bool `json:"IsReversed,omitempty" xml:"IsReversed,omitempty"`
	// Label 选择器
	LabelSelector *string `json:"LabelSelector,omitempty" xml:"LabelSelector,omitempty"`
	// 查找条件
	Search *string `json:"Search,omitempty" xml:"Search,omitempty"`
}

func (s ListTemplatesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTemplatesRequest) GoString() string {
	return s.String()
}

func (s *ListTemplatesRequest) SetWorkspace(v string) *ListTemplatesRequest {
	s.Workspace = &v
	return s
}

func (s *ListTemplatesRequest) SetNextToken(v string) *ListTemplatesRequest {
	s.NextToken = &v
	return s
}

func (s *ListTemplatesRequest) SetMaxResults(v int32) *ListTemplatesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListTemplatesRequest) SetOrderBy(v string) *ListTemplatesRequest {
	s.OrderBy = &v
	return s
}

func (s *ListTemplatesRequest) SetIsReversed(v bool) *ListTemplatesRequest {
	s.IsReversed = &v
	return s
}

func (s *ListTemplatesRequest) SetLabelSelector(v string) *ListTemplatesRequest {
	s.LabelSelector = &v
	return s
}

func (s *ListTemplatesRequest) SetSearch(v string) *ListTemplatesRequest {
	s.Search = &v
	return s
}

type ListTemplatesResponseBody struct {
	// 主机ID
	HostId *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	// 最大返回结果
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// 下次查询Token
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// 请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 应用模板列表
	Templates []*ListTemplatesResponseBodyTemplates `json:"Templates,omitempty" xml:"Templates,omitempty" type:"Repeated"`
	// 返回个数
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListTemplatesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *ListTemplatesResponseBody) SetHostId(v string) *ListTemplatesResponseBody {
	s.HostId = &v
	return s
}

func (s *ListTemplatesResponseBody) SetMaxResults(v int32) *ListTemplatesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListTemplatesResponseBody) SetNextToken(v string) *ListTemplatesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTemplatesResponseBody) SetRequestId(v string) *ListTemplatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTemplatesResponseBody) SetTemplates(v []*ListTemplatesResponseBodyTemplates) *ListTemplatesResponseBody {
	s.Templates = v
	return s
}

func (s *ListTemplatesResponseBody) SetTotalCount(v int32) *ListTemplatesResponseBody {
	s.TotalCount = &v
	return s
}

type ListTemplatesResponseBodyTemplates struct {
	// 工作空间
	Workspace *string `json:"Workspace,omitempty" xml:"Workspace,omitempty"`
	// 应用模板名称
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// 标签
	Labels map[string]*string `json:"Labels,omitempty" xml:"Labels,omitempty"`
	// 应用名称
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// 应用版本
	Revision *string `json:"Revision,omitempty" xml:"Revision,omitempty"`
	// 实体类型
	RootEntity *string `json:"RootEntity,omitempty" xml:"RootEntity,omitempty"`
	// 应用输入
	Inputs []*ListTemplatesResponseBodyTemplatesInputs `json:"Inputs,omitempty" xml:"Inputs,omitempty" type:"Repeated"`
	// 应用的输出参数
	Outputs []*ListTemplatesResponseBodyTemplatesOutputs `json:"Outputs,omitempty" xml:"Outputs,omitempty" type:"Repeated"`
	// 模板描述信息
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 最后修改时间
	LastModifiedTime *string `json:"LastModifiedTime,omitempty" xml:"LastModifiedTime,omitempty"`
}

func (s ListTemplatesResponseBodyTemplates) String() string {
	return tea.Prettify(s)
}

func (s ListTemplatesResponseBodyTemplates) GoString() string {
	return s.String()
}

func (s *ListTemplatesResponseBodyTemplates) SetWorkspace(v string) *ListTemplatesResponseBodyTemplates {
	s.Workspace = &v
	return s
}

func (s *ListTemplatesResponseBodyTemplates) SetTemplateName(v string) *ListTemplatesResponseBodyTemplates {
	s.TemplateName = &v
	return s
}

func (s *ListTemplatesResponseBodyTemplates) SetLabels(v map[string]*string) *ListTemplatesResponseBodyTemplates {
	s.Labels = v
	return s
}

func (s *ListTemplatesResponseBodyTemplates) SetAppName(v string) *ListTemplatesResponseBodyTemplates {
	s.AppName = &v
	return s
}

func (s *ListTemplatesResponseBodyTemplates) SetRevision(v string) *ListTemplatesResponseBodyTemplates {
	s.Revision = &v
	return s
}

func (s *ListTemplatesResponseBodyTemplates) SetRootEntity(v string) *ListTemplatesResponseBodyTemplates {
	s.RootEntity = &v
	return s
}

func (s *ListTemplatesResponseBodyTemplates) SetInputs(v []*ListTemplatesResponseBodyTemplatesInputs) *ListTemplatesResponseBodyTemplates {
	s.Inputs = v
	return s
}

func (s *ListTemplatesResponseBodyTemplates) SetOutputs(v []*ListTemplatesResponseBodyTemplatesOutputs) *ListTemplatesResponseBodyTemplates {
	s.Outputs = v
	return s
}

func (s *ListTemplatesResponseBodyTemplates) SetDescription(v string) *ListTemplatesResponseBodyTemplates {
	s.Description = &v
	return s
}

func (s *ListTemplatesResponseBodyTemplates) SetCreateTime(v string) *ListTemplatesResponseBodyTemplates {
	s.CreateTime = &v
	return s
}

func (s *ListTemplatesResponseBodyTemplates) SetLastModifiedTime(v string) *ListTemplatesResponseBodyTemplates {
	s.LastModifiedTime = &v
	return s
}

type ListTemplatesResponseBodyTemplatesInputs struct {
	// 任务名称
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// 变量名称
	VariableName *string `json:"VariableName,omitempty" xml:"VariableName,omitempty"`
	// 变量类型
	VariableType *string `json:"VariableType,omitempty" xml:"VariableType,omitempty"`
	// 变量值
	VariableValue *string `json:"VariableValue,omitempty" xml:"VariableValue,omitempty"`
	// 是否必须参数
	Required *bool `json:"Required,omitempty" xml:"Required,omitempty"`
	// 帮助信息
	Help *string `json:"Help,omitempty" xml:"Help,omitempty"`
	// 步骤顺序
	StepOrder *int64 `json:"StepOrder,omitempty" xml:"StepOrder,omitempty"`
}

func (s ListTemplatesResponseBodyTemplatesInputs) String() string {
	return tea.Prettify(s)
}

func (s ListTemplatesResponseBodyTemplatesInputs) GoString() string {
	return s.String()
}

func (s *ListTemplatesResponseBodyTemplatesInputs) SetTaskName(v string) *ListTemplatesResponseBodyTemplatesInputs {
	s.TaskName = &v
	return s
}

func (s *ListTemplatesResponseBodyTemplatesInputs) SetVariableName(v string) *ListTemplatesResponseBodyTemplatesInputs {
	s.VariableName = &v
	return s
}

func (s *ListTemplatesResponseBodyTemplatesInputs) SetVariableType(v string) *ListTemplatesResponseBodyTemplatesInputs {
	s.VariableType = &v
	return s
}

func (s *ListTemplatesResponseBodyTemplatesInputs) SetVariableValue(v string) *ListTemplatesResponseBodyTemplatesInputs {
	s.VariableValue = &v
	return s
}

func (s *ListTemplatesResponseBodyTemplatesInputs) SetRequired(v bool) *ListTemplatesResponseBodyTemplatesInputs {
	s.Required = &v
	return s
}

func (s *ListTemplatesResponseBodyTemplatesInputs) SetHelp(v string) *ListTemplatesResponseBodyTemplatesInputs {
	s.Help = &v
	return s
}

func (s *ListTemplatesResponseBodyTemplatesInputs) SetStepOrder(v int64) *ListTemplatesResponseBodyTemplatesInputs {
	s.StepOrder = &v
	return s
}

type ListTemplatesResponseBodyTemplatesOutputs struct {
	// 任务名称
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// 变量名称
	VariableName *string `json:"VariableName,omitempty" xml:"VariableName,omitempty"`
	// 变量类型
	VariableType *string `json:"VariableType,omitempty" xml:"VariableType,omitempty"`
	// 变量值
	VariableValue *string `json:"VariableValue,omitempty" xml:"VariableValue,omitempty"`
	// 是否必须参数
	Required *bool `json:"Required,omitempty" xml:"Required,omitempty"`
	// 帮助信息
	Help *string `json:"Help,omitempty" xml:"Help,omitempty"`
	// 步骤顺序
	StepOrder *int64 `json:"StepOrder,omitempty" xml:"StepOrder,omitempty"`
}

func (s ListTemplatesResponseBodyTemplatesOutputs) String() string {
	return tea.Prettify(s)
}

func (s ListTemplatesResponseBodyTemplatesOutputs) GoString() string {
	return s.String()
}

func (s *ListTemplatesResponseBodyTemplatesOutputs) SetTaskName(v string) *ListTemplatesResponseBodyTemplatesOutputs {
	s.TaskName = &v
	return s
}

func (s *ListTemplatesResponseBodyTemplatesOutputs) SetVariableName(v string) *ListTemplatesResponseBodyTemplatesOutputs {
	s.VariableName = &v
	return s
}

func (s *ListTemplatesResponseBodyTemplatesOutputs) SetVariableType(v string) *ListTemplatesResponseBodyTemplatesOutputs {
	s.VariableType = &v
	return s
}

func (s *ListTemplatesResponseBodyTemplatesOutputs) SetVariableValue(v string) *ListTemplatesResponseBodyTemplatesOutputs {
	s.VariableValue = &v
	return s
}

func (s *ListTemplatesResponseBodyTemplatesOutputs) SetRequired(v bool) *ListTemplatesResponseBodyTemplatesOutputs {
	s.Required = &v
	return s
}

func (s *ListTemplatesResponseBodyTemplatesOutputs) SetHelp(v string) *ListTemplatesResponseBodyTemplatesOutputs {
	s.Help = &v
	return s
}

func (s *ListTemplatesResponseBodyTemplatesOutputs) SetStepOrder(v int64) *ListTemplatesResponseBodyTemplatesOutputs {
	s.StepOrder = &v
	return s
}

type ListTemplatesResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListTemplatesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListTemplatesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTemplatesResponse) GoString() string {
	return s.String()
}

func (s *ListTemplatesResponse) SetHeaders(v map[string]*string) *ListTemplatesResponse {
	s.Headers = v
	return s
}

func (s *ListTemplatesResponse) SetBody(v *ListTemplatesResponseBody) *ListTemplatesResponse {
	s.Body = v
	return s
}

type ListUserActiveRunsRequest struct {
	// 工作空间
	Workspace *string `json:"Workspace,omitempty" xml:"Workspace,omitempty"`
	// 任务名字
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 状态
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 提交ID
	SubmissionId *string `json:"SubmissionId,omitempty" xml:"SubmissionId,omitempty"`
	// 搜索ID
	Search *string `json:"Search,omitempty" xml:"Search,omitempty"`
	// 标签选择
	LabelSelector *string `json:"LabelSelector,omitempty" xml:"LabelSelector,omitempty"`
	// 应用名称
	AppName    *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	NextToken  *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	MaxResults *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	OrderBy    *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	IsReversed *bool   `json:"IsReversed,omitempty" xml:"IsReversed,omitempty"`
}

func (s ListUserActiveRunsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListUserActiveRunsRequest) GoString() string {
	return s.String()
}

func (s *ListUserActiveRunsRequest) SetWorkspace(v string) *ListUserActiveRunsRequest {
	s.Workspace = &v
	return s
}

func (s *ListUserActiveRunsRequest) SetName(v string) *ListUserActiveRunsRequest {
	s.Name = &v
	return s
}

func (s *ListUserActiveRunsRequest) SetStatus(v string) *ListUserActiveRunsRequest {
	s.Status = &v
	return s
}

func (s *ListUserActiveRunsRequest) SetSubmissionId(v string) *ListUserActiveRunsRequest {
	s.SubmissionId = &v
	return s
}

func (s *ListUserActiveRunsRequest) SetSearch(v string) *ListUserActiveRunsRequest {
	s.Search = &v
	return s
}

func (s *ListUserActiveRunsRequest) SetLabelSelector(v string) *ListUserActiveRunsRequest {
	s.LabelSelector = &v
	return s
}

func (s *ListUserActiveRunsRequest) SetAppName(v string) *ListUserActiveRunsRequest {
	s.AppName = &v
	return s
}

func (s *ListUserActiveRunsRequest) SetNextToken(v string) *ListUserActiveRunsRequest {
	s.NextToken = &v
	return s
}

func (s *ListUserActiveRunsRequest) SetMaxResults(v int32) *ListUserActiveRunsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListUserActiveRunsRequest) SetOrderBy(v string) *ListUserActiveRunsRequest {
	s.OrderBy = &v
	return s
}

func (s *ListUserActiveRunsRequest) SetIsReversed(v bool) *ListUserActiveRunsRequest {
	s.IsReversed = &v
	return s
}

type ListUserActiveRunsResponseBody struct {
	// 请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 主机ID
	HostId *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	// 最大返回结果
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// 下次查询Token
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// 返回个数
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// 任务列表
	Runs []*ListUserActiveRunsResponseBodyRuns `json:"Runs,omitempty" xml:"Runs,omitempty" type:"Repeated"`
}

func (s ListUserActiveRunsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListUserActiveRunsResponseBody) GoString() string {
	return s.String()
}

func (s *ListUserActiveRunsResponseBody) SetRequestId(v string) *ListUserActiveRunsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUserActiveRunsResponseBody) SetHostId(v string) *ListUserActiveRunsResponseBody {
	s.HostId = &v
	return s
}

func (s *ListUserActiveRunsResponseBody) SetMaxResults(v int32) *ListUserActiveRunsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListUserActiveRunsResponseBody) SetNextToken(v string) *ListUserActiveRunsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListUserActiveRunsResponseBody) SetTotalCount(v int32) *ListUserActiveRunsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListUserActiveRunsResponseBody) SetRuns(v []*ListUserActiveRunsResponseBodyRuns) *ListUserActiveRunsResponseBody {
	s.Runs = v
	return s
}

type ListUserActiveRunsResponseBodyRuns struct {
	// 应用名称
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// 应用来源
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// 应用命名空间
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// 应用原名
	AppOrigName *string `json:"AppOrigName,omitempty" xml:"AppOrigName,omitempty"`
	// 任务标签
	Labels map[string]*string `json:"Labels,omitempty" xml:"Labels,omitempty"`
	// 任务名称
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 应用版本号
	Revision *string `json:"Revision,omitempty" xml:"Revision,omitempty"`
	// 任务ID
	RunId *string `json:"RunId,omitempty" xml:"RunId,omitempty"`
	// 任务状态
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 提交时间
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 开始时间
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// 结束时间
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// 提交ID
	SubmissionId *string `json:"SubmissionId,omitempty" xml:"SubmissionId,omitempty"`
	// 实体名称
	EntityName *string `json:"EntityName,omitempty" xml:"EntityName,omitempty"`
	// 实体对象ID
	EntityType *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	// 运行目录
	ExecuteDirectory *string                                           `json:"ExecuteDirectory,omitempty" xml:"ExecuteDirectory,omitempty"`
	ExecuteOptions   *ListUserActiveRunsResponseBodyRunsExecuteOptions `json:"ExecuteOptions,omitempty" xml:"ExecuteOptions,omitempty" type:"Struct"`
	// 输入参数
	Inputs         *string `json:"Inputs,omitempty" xml:"Inputs,omitempty"`
	DefaultRuntime *string `json:"DefaultRuntime,omitempty" xml:"DefaultRuntime,omitempty"`
	Workspace      *string `json:"Workspace,omitempty" xml:"Workspace,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListUserActiveRunsResponseBodyRuns) String() string {
	return tea.Prettify(s)
}

func (s ListUserActiveRunsResponseBodyRuns) GoString() string {
	return s.String()
}

func (s *ListUserActiveRunsResponseBodyRuns) SetAppName(v string) *ListUserActiveRunsResponseBodyRuns {
	s.AppName = &v
	return s
}

func (s *ListUserActiveRunsResponseBodyRuns) SetSource(v string) *ListUserActiveRunsResponseBodyRuns {
	s.Source = &v
	return s
}

func (s *ListUserActiveRunsResponseBodyRuns) SetNamespace(v string) *ListUserActiveRunsResponseBodyRuns {
	s.Namespace = &v
	return s
}

func (s *ListUserActiveRunsResponseBodyRuns) SetAppOrigName(v string) *ListUserActiveRunsResponseBodyRuns {
	s.AppOrigName = &v
	return s
}

func (s *ListUserActiveRunsResponseBodyRuns) SetLabels(v map[string]*string) *ListUserActiveRunsResponseBodyRuns {
	s.Labels = v
	return s
}

func (s *ListUserActiveRunsResponseBodyRuns) SetName(v string) *ListUserActiveRunsResponseBodyRuns {
	s.Name = &v
	return s
}

func (s *ListUserActiveRunsResponseBodyRuns) SetRevision(v string) *ListUserActiveRunsResponseBodyRuns {
	s.Revision = &v
	return s
}

func (s *ListUserActiveRunsResponseBodyRuns) SetRunId(v string) *ListUserActiveRunsResponseBodyRuns {
	s.RunId = &v
	return s
}

func (s *ListUserActiveRunsResponseBodyRuns) SetStatus(v string) *ListUserActiveRunsResponseBodyRuns {
	s.Status = &v
	return s
}

func (s *ListUserActiveRunsResponseBodyRuns) SetCreateTime(v string) *ListUserActiveRunsResponseBodyRuns {
	s.CreateTime = &v
	return s
}

func (s *ListUserActiveRunsResponseBodyRuns) SetStartTime(v string) *ListUserActiveRunsResponseBodyRuns {
	s.StartTime = &v
	return s
}

func (s *ListUserActiveRunsResponseBodyRuns) SetEndTime(v string) *ListUserActiveRunsResponseBodyRuns {
	s.EndTime = &v
	return s
}

func (s *ListUserActiveRunsResponseBodyRuns) SetSubmissionId(v string) *ListUserActiveRunsResponseBodyRuns {
	s.SubmissionId = &v
	return s
}

func (s *ListUserActiveRunsResponseBodyRuns) SetEntityName(v string) *ListUserActiveRunsResponseBodyRuns {
	s.EntityName = &v
	return s
}

func (s *ListUserActiveRunsResponseBodyRuns) SetEntityType(v string) *ListUserActiveRunsResponseBodyRuns {
	s.EntityType = &v
	return s
}

func (s *ListUserActiveRunsResponseBodyRuns) SetExecuteDirectory(v string) *ListUserActiveRunsResponseBodyRuns {
	s.ExecuteDirectory = &v
	return s
}

func (s *ListUserActiveRunsResponseBodyRuns) SetExecuteOptions(v *ListUserActiveRunsResponseBodyRunsExecuteOptions) *ListUserActiveRunsResponseBodyRuns {
	s.ExecuteOptions = v
	return s
}

func (s *ListUserActiveRunsResponseBodyRuns) SetInputs(v string) *ListUserActiveRunsResponseBodyRuns {
	s.Inputs = &v
	return s
}

func (s *ListUserActiveRunsResponseBodyRuns) SetDefaultRuntime(v string) *ListUserActiveRunsResponseBodyRuns {
	s.DefaultRuntime = &v
	return s
}

func (s *ListUserActiveRunsResponseBodyRuns) SetWorkspace(v string) *ListUserActiveRunsResponseBodyRuns {
	s.Workspace = &v
	return s
}

func (s *ListUserActiveRunsResponseBodyRuns) SetRegionId(v string) *ListUserActiveRunsResponseBodyRuns {
	s.RegionId = &v
	return s
}

type ListUserActiveRunsResponseBodyRunsExecuteOptions struct {
	CallCaching               *bool   `json:"CallCaching,omitempty" xml:"CallCaching,omitempty"`
	DeleteIntermediateResults *bool   `json:"DeleteIntermediateResults,omitempty" xml:"DeleteIntermediateResults,omitempty"`
	FailureMode               *string `json:"FailureMode,omitempty" xml:"FailureMode,omitempty"`
}

func (s ListUserActiveRunsResponseBodyRunsExecuteOptions) String() string {
	return tea.Prettify(s)
}

func (s ListUserActiveRunsResponseBodyRunsExecuteOptions) GoString() string {
	return s.String()
}

func (s *ListUserActiveRunsResponseBodyRunsExecuteOptions) SetCallCaching(v bool) *ListUserActiveRunsResponseBodyRunsExecuteOptions {
	s.CallCaching = &v
	return s
}

func (s *ListUserActiveRunsResponseBodyRunsExecuteOptions) SetDeleteIntermediateResults(v bool) *ListUserActiveRunsResponseBodyRunsExecuteOptions {
	s.DeleteIntermediateResults = &v
	return s
}

func (s *ListUserActiveRunsResponseBodyRunsExecuteOptions) SetFailureMode(v string) *ListUserActiveRunsResponseBodyRunsExecuteOptions {
	s.FailureMode = &v
	return s
}

type ListUserActiveRunsResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListUserActiveRunsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListUserActiveRunsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListUserActiveRunsResponse) GoString() string {
	return s.String()
}

func (s *ListUserActiveRunsResponse) SetHeaders(v map[string]*string) *ListUserActiveRunsResponse {
	s.Headers = v
	return s
}

func (s *ListUserActiveRunsResponse) SetBody(v *ListUserActiveRunsResponseBody) *ListUserActiveRunsResponse {
	s.Body = v
	return s
}

type ListWorkspacesRequest struct {
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OrderBy    *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	IsReversed *bool   `json:"IsReversed,omitempty" xml:"IsReversed,omitempty"`
	NextToken  *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	MaxResults *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// Label选择器
	LabelSelector *string `json:"LabelSelector,omitempty" xml:"LabelSelector,omitempty"`
}

func (s ListWorkspacesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListWorkspacesRequest) GoString() string {
	return s.String()
}

func (s *ListWorkspacesRequest) SetName(v string) *ListWorkspacesRequest {
	s.Name = &v
	return s
}

func (s *ListWorkspacesRequest) SetOrderBy(v string) *ListWorkspacesRequest {
	s.OrderBy = &v
	return s
}

func (s *ListWorkspacesRequest) SetIsReversed(v bool) *ListWorkspacesRequest {
	s.IsReversed = &v
	return s
}

func (s *ListWorkspacesRequest) SetNextToken(v string) *ListWorkspacesRequest {
	s.NextToken = &v
	return s
}

func (s *ListWorkspacesRequest) SetMaxResults(v int32) *ListWorkspacesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListWorkspacesRequest) SetLabelSelector(v string) *ListWorkspacesRequest {
	s.LabelSelector = &v
	return s
}

type ListWorkspacesResponseBody struct {
	// 主机ID
	HostId *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	// 最大结果数
	MaxResults *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 下次查询的起始Token
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// 返回总个数
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// 工作空间数组
	Workspaces []*ListWorkspacesResponseBodyWorkspaces `json:"Workspaces,omitempty" xml:"Workspaces,omitempty" type:"Repeated"`
}

func (s ListWorkspacesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListWorkspacesResponseBody) GoString() string {
	return s.String()
}

func (s *ListWorkspacesResponseBody) SetHostId(v string) *ListWorkspacesResponseBody {
	s.HostId = &v
	return s
}

func (s *ListWorkspacesResponseBody) SetMaxResults(v int32) *ListWorkspacesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListWorkspacesResponseBody) SetRequestId(v string) *ListWorkspacesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListWorkspacesResponseBody) SetNextToken(v string) *ListWorkspacesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListWorkspacesResponseBody) SetTotalCount(v int32) *ListWorkspacesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListWorkspacesResponseBody) SetWorkspaces(v []*ListWorkspacesResponseBodyWorkspaces) *ListWorkspacesResponseBody {
	s.Workspaces = v
	return s
}

type ListWorkspacesResponseBodyWorkspaces struct {
	// 工作空间描述
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 任务生命周期
	JobLifecycle *int32 `json:"JobLifecycle,omitempty" xml:"JobLifecycle,omitempty"`
	// 工作空间标签
	Labels map[string]*string `json:"Labels,omitempty" xml:"Labels,omitempty"`
	// OSS工作路径
	OssRoot *string `json:"OssRoot,omitempty" xml:"OssRoot,omitempty"`
	// 工作空间Bucket名字
	BucketName *string `json:"BucketName,omitempty" xml:"BucketName,omitempty"`
	// RAM Role
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 最后修改时间
	LastModifiedTime *string `json:"LastModifiedTime,omitempty" xml:"LastModifiedTime,omitempty"`
	// 工作空间名称
	Workspace *string `json:"Workspace,omitempty" xml:"Workspace,omitempty"`
	// 工作空间状态
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 地域ID
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListWorkspacesResponseBodyWorkspaces) String() string {
	return tea.Prettify(s)
}

func (s ListWorkspacesResponseBodyWorkspaces) GoString() string {
	return s.String()
}

func (s *ListWorkspacesResponseBodyWorkspaces) SetDescription(v string) *ListWorkspacesResponseBodyWorkspaces {
	s.Description = &v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspaces) SetJobLifecycle(v int32) *ListWorkspacesResponseBodyWorkspaces {
	s.JobLifecycle = &v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspaces) SetLabels(v map[string]*string) *ListWorkspacesResponseBodyWorkspaces {
	s.Labels = v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspaces) SetOssRoot(v string) *ListWorkspacesResponseBodyWorkspaces {
	s.OssRoot = &v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspaces) SetBucketName(v string) *ListWorkspacesResponseBodyWorkspaces {
	s.BucketName = &v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspaces) SetRole(v string) *ListWorkspacesResponseBodyWorkspaces {
	s.Role = &v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspaces) SetCreateTime(v string) *ListWorkspacesResponseBodyWorkspaces {
	s.CreateTime = &v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspaces) SetLastModifiedTime(v string) *ListWorkspacesResponseBodyWorkspaces {
	s.LastModifiedTime = &v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspaces) SetWorkspace(v string) *ListWorkspacesResponseBodyWorkspaces {
	s.Workspace = &v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspaces) SetStatus(v string) *ListWorkspacesResponseBodyWorkspaces {
	s.Status = &v
	return s
}

func (s *ListWorkspacesResponseBodyWorkspaces) SetRegionId(v string) *ListWorkspacesResponseBodyWorkspaces {
	s.RegionId = &v
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

type ParseAppInputsRequest struct {
	// 工作空间名称
	Workspace *string `json:"Workspace,omitempty" xml:"Workspace,omitempty"`
	// 应用名称
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// 文件路径
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// 应用定义
	Definition *string `json:"Definition,omitempty" xml:"Definition,omitempty"`
	// 依赖应用
	Dependencies []*ParseAppInputsRequestDependencies `json:"Dependencies,omitempty" xml:"Dependencies,omitempty" type:"Repeated"`
}

func (s ParseAppInputsRequest) String() string {
	return tea.Prettify(s)
}

func (s ParseAppInputsRequest) GoString() string {
	return s.String()
}

func (s *ParseAppInputsRequest) SetWorkspace(v string) *ParseAppInputsRequest {
	s.Workspace = &v
	return s
}

func (s *ParseAppInputsRequest) SetAppName(v string) *ParseAppInputsRequest {
	s.AppName = &v
	return s
}

func (s *ParseAppInputsRequest) SetPath(v string) *ParseAppInputsRequest {
	s.Path = &v
	return s
}

func (s *ParseAppInputsRequest) SetDefinition(v string) *ParseAppInputsRequest {
	s.Definition = &v
	return s
}

func (s *ParseAppInputsRequest) SetDependencies(v []*ParseAppInputsRequestDependencies) *ParseAppInputsRequest {
	s.Dependencies = v
	return s
}

type ParseAppInputsRequestDependencies struct {
	// 依赖路径
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// 依赖内容
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
}

func (s ParseAppInputsRequestDependencies) String() string {
	return tea.Prettify(s)
}

func (s ParseAppInputsRequestDependencies) GoString() string {
	return s.String()
}

func (s *ParseAppInputsRequestDependencies) SetPath(v string) *ParseAppInputsRequestDependencies {
	s.Path = &v
	return s
}

func (s *ParseAppInputsRequestDependencies) SetContent(v string) *ParseAppInputsRequestDependencies {
	s.Content = &v
	return s
}

type ParseAppInputsShrinkRequest struct {
	// 工作空间名称
	Workspace *string `json:"Workspace,omitempty" xml:"Workspace,omitempty"`
	// 应用名称
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// 文件路径
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// 应用定义
	Definition *string `json:"Definition,omitempty" xml:"Definition,omitempty"`
	// 依赖应用
	DependenciesShrink *string `json:"Dependencies,omitempty" xml:"Dependencies,omitempty"`
}

func (s ParseAppInputsShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ParseAppInputsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ParseAppInputsShrinkRequest) SetWorkspace(v string) *ParseAppInputsShrinkRequest {
	s.Workspace = &v
	return s
}

func (s *ParseAppInputsShrinkRequest) SetAppName(v string) *ParseAppInputsShrinkRequest {
	s.AppName = &v
	return s
}

func (s *ParseAppInputsShrinkRequest) SetPath(v string) *ParseAppInputsShrinkRequest {
	s.Path = &v
	return s
}

func (s *ParseAppInputsShrinkRequest) SetDefinition(v string) *ParseAppInputsShrinkRequest {
	s.Definition = &v
	return s
}

func (s *ParseAppInputsShrinkRequest) SetDependenciesShrink(v string) *ParseAppInputsShrinkRequest {
	s.DependenciesShrink = &v
	return s
}

type ParseAppInputsResponseBody struct {
	// 主机 ID
	HostId *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	// 请求 ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 应用输入参数
	Inputs []*ParseAppInputsResponseBodyInputs `json:"Inputs,omitempty" xml:"Inputs,omitempty" type:"Repeated"`
}

func (s ParseAppInputsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ParseAppInputsResponseBody) GoString() string {
	return s.String()
}

func (s *ParseAppInputsResponseBody) SetHostId(v string) *ParseAppInputsResponseBody {
	s.HostId = &v
	return s
}

func (s *ParseAppInputsResponseBody) SetRequestId(v string) *ParseAppInputsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ParseAppInputsResponseBody) SetInputs(v []*ParseAppInputsResponseBodyInputs) *ParseAppInputsResponseBody {
	s.Inputs = v
	return s
}

type ParseAppInputsResponseBodyInputs struct {
	// 参数名称
	TaskName      *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	VariableName  *string `json:"VariableName,omitempty" xml:"VariableName,omitempty"`
	VariableType  *string `json:"VariableType,omitempty" xml:"VariableType,omitempty"`
	VariableValue *string `json:"VariableValue,omitempty" xml:"VariableValue,omitempty"`
	Required      *bool   `json:"Required,omitempty" xml:"Required,omitempty"`
	Help          *string `json:"Help,omitempty" xml:"Help,omitempty"`
	StepOrder     *int64  `json:"StepOrder,omitempty" xml:"StepOrder,omitempty"`
}

func (s ParseAppInputsResponseBodyInputs) String() string {
	return tea.Prettify(s)
}

func (s ParseAppInputsResponseBodyInputs) GoString() string {
	return s.String()
}

func (s *ParseAppInputsResponseBodyInputs) SetTaskName(v string) *ParseAppInputsResponseBodyInputs {
	s.TaskName = &v
	return s
}

func (s *ParseAppInputsResponseBodyInputs) SetVariableName(v string) *ParseAppInputsResponseBodyInputs {
	s.VariableName = &v
	return s
}

func (s *ParseAppInputsResponseBodyInputs) SetVariableType(v string) *ParseAppInputsResponseBodyInputs {
	s.VariableType = &v
	return s
}

func (s *ParseAppInputsResponseBodyInputs) SetVariableValue(v string) *ParseAppInputsResponseBodyInputs {
	s.VariableValue = &v
	return s
}

func (s *ParseAppInputsResponseBodyInputs) SetRequired(v bool) *ParseAppInputsResponseBodyInputs {
	s.Required = &v
	return s
}

func (s *ParseAppInputsResponseBodyInputs) SetHelp(v string) *ParseAppInputsResponseBodyInputs {
	s.Help = &v
	return s
}

func (s *ParseAppInputsResponseBodyInputs) SetStepOrder(v int64) *ParseAppInputsResponseBodyInputs {
	s.StepOrder = &v
	return s
}

type ParseAppInputsResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ParseAppInputsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ParseAppInputsResponse) String() string {
	return tea.Prettify(s)
}

func (s ParseAppInputsResponse) GoString() string {
	return s.String()
}

func (s *ParseAppInputsResponse) SetHeaders(v map[string]*string) *ParseAppInputsResponse {
	s.Headers = v
	return s
}

func (s *ParseAppInputsResponse) SetBody(v *ParseAppInputsResponseBody) *ParseAppInputsResponse {
	s.Body = v
	return s
}

type ResumeSubmissionRequest struct {
	// 工作空间名称
	Workspace *string `json:"Workspace,omitempty" xml:"Workspace,omitempty"`
	// 投递ID
	SubmissionId *string `json:"SubmissionId,omitempty" xml:"SubmissionId,omitempty"`
}

func (s ResumeSubmissionRequest) String() string {
	return tea.Prettify(s)
}

func (s ResumeSubmissionRequest) GoString() string {
	return s.String()
}

func (s *ResumeSubmissionRequest) SetWorkspace(v string) *ResumeSubmissionRequest {
	s.Workspace = &v
	return s
}

func (s *ResumeSubmissionRequest) SetSubmissionId(v string) *ResumeSubmissionRequest {
	s.SubmissionId = &v
	return s
}

type ResumeSubmissionResponseBody struct {
	// 主机ID
	HostId *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	// 请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ResumeSubmissionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ResumeSubmissionResponseBody) GoString() string {
	return s.String()
}

func (s *ResumeSubmissionResponseBody) SetHostId(v string) *ResumeSubmissionResponseBody {
	s.HostId = &v
	return s
}

func (s *ResumeSubmissionResponseBody) SetRequestId(v string) *ResumeSubmissionResponseBody {
	s.RequestId = &v
	return s
}

type ResumeSubmissionResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ResumeSubmissionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ResumeSubmissionResponse) String() string {
	return tea.Prettify(s)
}

func (s ResumeSubmissionResponse) GoString() string {
	return s.String()
}

func (s *ResumeSubmissionResponse) SetHeaders(v map[string]*string) *ResumeSubmissionResponse {
	s.Headers = v
	return s
}

func (s *ResumeSubmissionResponse) SetBody(v *ResumeSubmissionResponseBody) *ResumeSubmissionResponse {
	s.Body = v
	return s
}

type UpdateEntityRequest struct {
	Workspace   *string                           `json:"Workspace,omitempty" xml:"Workspace,omitempty"`
	EntityType  *string                           `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	EntityItems []*UpdateEntityRequestEntityItems `json:"EntityItems,omitempty" xml:"EntityItems,omitempty" type:"Repeated"`
}

func (s UpdateEntityRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateEntityRequest) GoString() string {
	return s.String()
}

func (s *UpdateEntityRequest) SetWorkspace(v string) *UpdateEntityRequest {
	s.Workspace = &v
	return s
}

func (s *UpdateEntityRequest) SetEntityType(v string) *UpdateEntityRequest {
	s.EntityType = &v
	return s
}

func (s *UpdateEntityRequest) SetEntityItems(v []*UpdateEntityRequestEntityItems) *UpdateEntityRequest {
	s.EntityItems = v
	return s
}

type UpdateEntityRequestEntityItems struct {
	EntityName *string            `json:"EntityName,omitempty" xml:"EntityName,omitempty"`
	EntityData map[string]*string `json:"EntityData,omitempty" xml:"EntityData,omitempty"`
}

func (s UpdateEntityRequestEntityItems) String() string {
	return tea.Prettify(s)
}

func (s UpdateEntityRequestEntityItems) GoString() string {
	return s.String()
}

func (s *UpdateEntityRequestEntityItems) SetEntityName(v string) *UpdateEntityRequestEntityItems {
	s.EntityName = &v
	return s
}

func (s *UpdateEntityRequestEntityItems) SetEntityData(v map[string]*string) *UpdateEntityRequestEntityItems {
	s.EntityData = v
	return s
}

type UpdateEntityShrinkRequest struct {
	Workspace         *string `json:"Workspace,omitempty" xml:"Workspace,omitempty"`
	EntityType        *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	EntityItemsShrink *string `json:"EntityItems,omitempty" xml:"EntityItems,omitempty"`
}

func (s UpdateEntityShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateEntityShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateEntityShrinkRequest) SetWorkspace(v string) *UpdateEntityShrinkRequest {
	s.Workspace = &v
	return s
}

func (s *UpdateEntityShrinkRequest) SetEntityType(v string) *UpdateEntityShrinkRequest {
	s.EntityType = &v
	return s
}

func (s *UpdateEntityShrinkRequest) SetEntityItemsShrink(v string) *UpdateEntityShrinkRequest {
	s.EntityItemsShrink = &v
	return s
}

type UpdateEntityResponseBody struct {
	// 主机ID
	HostId *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	// 请求ID
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Workspace  *string `json:"Workspace,omitempty" xml:"Workspace,omitempty"`
	EntityType *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
}

func (s UpdateEntityResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateEntityResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateEntityResponseBody) SetHostId(v string) *UpdateEntityResponseBody {
	s.HostId = &v
	return s
}

func (s *UpdateEntityResponseBody) SetRequestId(v string) *UpdateEntityResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateEntityResponseBody) SetWorkspace(v string) *UpdateEntityResponseBody {
	s.Workspace = &v
	return s
}

func (s *UpdateEntityResponseBody) SetEntityType(v string) *UpdateEntityResponseBody {
	s.EntityType = &v
	return s
}

type UpdateEntityResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateEntityResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateEntityResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateEntityResponse) GoString() string {
	return s.String()
}

func (s *UpdateEntityResponse) SetHeaders(v map[string]*string) *UpdateEntityResponse {
	s.Headers = v
	return s
}

func (s *UpdateEntityResponse) SetBody(v *UpdateEntityResponseBody) *UpdateEntityResponse {
	s.Body = v
	return s
}

type UpdateTemplateRequest struct {
	// 工作空间名称
	Workspace *string `json:"Workspace,omitempty" xml:"Workspace,omitempty"`
	// 应用模板名称
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// 工作空间描述
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 实体类型
	RootEntity *string `json:"RootEntity,omitempty" xml:"RootEntity,omitempty"`
	// 应用的输入
	Inputs []*UpdateTemplateRequestInputs `json:"Inputs,omitempty" xml:"Inputs,omitempty" type:"Repeated"`
	// 应用的输出
	Outputs []*UpdateTemplateRequestOutputs `json:"Outputs,omitempty" xml:"Outputs,omitempty" type:"Repeated"`
	// 工作空间标签
	Labels *string `json:"Labels,omitempty" xml:"Labels,omitempty"`
}

func (s UpdateTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateTemplateRequest) GoString() string {
	return s.String()
}

func (s *UpdateTemplateRequest) SetWorkspace(v string) *UpdateTemplateRequest {
	s.Workspace = &v
	return s
}

func (s *UpdateTemplateRequest) SetTemplateName(v string) *UpdateTemplateRequest {
	s.TemplateName = &v
	return s
}

func (s *UpdateTemplateRequest) SetDescription(v string) *UpdateTemplateRequest {
	s.Description = &v
	return s
}

func (s *UpdateTemplateRequest) SetRootEntity(v string) *UpdateTemplateRequest {
	s.RootEntity = &v
	return s
}

func (s *UpdateTemplateRequest) SetInputs(v []*UpdateTemplateRequestInputs) *UpdateTemplateRequest {
	s.Inputs = v
	return s
}

func (s *UpdateTemplateRequest) SetOutputs(v []*UpdateTemplateRequestOutputs) *UpdateTemplateRequest {
	s.Outputs = v
	return s
}

func (s *UpdateTemplateRequest) SetLabels(v string) *UpdateTemplateRequest {
	s.Labels = &v
	return s
}

type UpdateTemplateRequestInputs struct {
	// 任务名称
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// 变量名
	VariableName *string `json:"VariableName,omitempty" xml:"VariableName,omitempty"`
	// 变量类型
	VariableType *string `json:"VariableType,omitempty" xml:"VariableType,omitempty"`
	// 变量值
	VariableValue *string `json:"VariableValue,omitempty" xml:"VariableValue,omitempty"`
	// 是否必填
	Required *bool `json:"Required,omitempty" xml:"Required,omitempty"`
	// 帮助信息
	Help *string `json:"Help,omitempty" xml:"Help,omitempty"`
	// 步骤顺序
	StepOrder *int32 `json:"StepOrder,omitempty" xml:"StepOrder,omitempty"`
}

func (s UpdateTemplateRequestInputs) String() string {
	return tea.Prettify(s)
}

func (s UpdateTemplateRequestInputs) GoString() string {
	return s.String()
}

func (s *UpdateTemplateRequestInputs) SetTaskName(v string) *UpdateTemplateRequestInputs {
	s.TaskName = &v
	return s
}

func (s *UpdateTemplateRequestInputs) SetVariableName(v string) *UpdateTemplateRequestInputs {
	s.VariableName = &v
	return s
}

func (s *UpdateTemplateRequestInputs) SetVariableType(v string) *UpdateTemplateRequestInputs {
	s.VariableType = &v
	return s
}

func (s *UpdateTemplateRequestInputs) SetVariableValue(v string) *UpdateTemplateRequestInputs {
	s.VariableValue = &v
	return s
}

func (s *UpdateTemplateRequestInputs) SetRequired(v bool) *UpdateTemplateRequestInputs {
	s.Required = &v
	return s
}

func (s *UpdateTemplateRequestInputs) SetHelp(v string) *UpdateTemplateRequestInputs {
	s.Help = &v
	return s
}

func (s *UpdateTemplateRequestInputs) SetStepOrder(v int32) *UpdateTemplateRequestInputs {
	s.StepOrder = &v
	return s
}

type UpdateTemplateRequestOutputs struct {
	// 任务名称
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// 变量名
	VariableName *string `json:"VariableName,omitempty" xml:"VariableName,omitempty"`
	// 变量类型
	VariableType *string `json:"VariableType,omitempty" xml:"VariableType,omitempty"`
	// 变量值
	VariableValue *string `json:"VariableValue,omitempty" xml:"VariableValue,omitempty"`
	// 是否必填
	Required *bool `json:"Required,omitempty" xml:"Required,omitempty"`
	// 帮助信息
	Help *string `json:"Help,omitempty" xml:"Help,omitempty"`
	// 步骤顺序
	StepOrder *int32 `json:"StepOrder,omitempty" xml:"StepOrder,omitempty"`
}

func (s UpdateTemplateRequestOutputs) String() string {
	return tea.Prettify(s)
}

func (s UpdateTemplateRequestOutputs) GoString() string {
	return s.String()
}

func (s *UpdateTemplateRequestOutputs) SetTaskName(v string) *UpdateTemplateRequestOutputs {
	s.TaskName = &v
	return s
}

func (s *UpdateTemplateRequestOutputs) SetVariableName(v string) *UpdateTemplateRequestOutputs {
	s.VariableName = &v
	return s
}

func (s *UpdateTemplateRequestOutputs) SetVariableType(v string) *UpdateTemplateRequestOutputs {
	s.VariableType = &v
	return s
}

func (s *UpdateTemplateRequestOutputs) SetVariableValue(v string) *UpdateTemplateRequestOutputs {
	s.VariableValue = &v
	return s
}

func (s *UpdateTemplateRequestOutputs) SetRequired(v bool) *UpdateTemplateRequestOutputs {
	s.Required = &v
	return s
}

func (s *UpdateTemplateRequestOutputs) SetHelp(v string) *UpdateTemplateRequestOutputs {
	s.Help = &v
	return s
}

func (s *UpdateTemplateRequestOutputs) SetStepOrder(v int32) *UpdateTemplateRequestOutputs {
	s.StepOrder = &v
	return s
}

type UpdateTemplateShrinkRequest struct {
	// 工作空间名称
	Workspace *string `json:"Workspace,omitempty" xml:"Workspace,omitempty"`
	// 应用模板名称
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// 工作空间描述
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 实体类型
	RootEntity *string `json:"RootEntity,omitempty" xml:"RootEntity,omitempty"`
	// 应用的输入
	InputsShrink *string `json:"Inputs,omitempty" xml:"Inputs,omitempty"`
	// 应用的输出
	OutputsShrink *string `json:"Outputs,omitempty" xml:"Outputs,omitempty"`
	// 工作空间标签
	Labels *string `json:"Labels,omitempty" xml:"Labels,omitempty"`
}

func (s UpdateTemplateShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateTemplateShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateTemplateShrinkRequest) SetWorkspace(v string) *UpdateTemplateShrinkRequest {
	s.Workspace = &v
	return s
}

func (s *UpdateTemplateShrinkRequest) SetTemplateName(v string) *UpdateTemplateShrinkRequest {
	s.TemplateName = &v
	return s
}

func (s *UpdateTemplateShrinkRequest) SetDescription(v string) *UpdateTemplateShrinkRequest {
	s.Description = &v
	return s
}

func (s *UpdateTemplateShrinkRequest) SetRootEntity(v string) *UpdateTemplateShrinkRequest {
	s.RootEntity = &v
	return s
}

func (s *UpdateTemplateShrinkRequest) SetInputsShrink(v string) *UpdateTemplateShrinkRequest {
	s.InputsShrink = &v
	return s
}

func (s *UpdateTemplateShrinkRequest) SetOutputsShrink(v string) *UpdateTemplateShrinkRequest {
	s.OutputsShrink = &v
	return s
}

func (s *UpdateTemplateShrinkRequest) SetLabels(v string) *UpdateTemplateShrinkRequest {
	s.Labels = &v
	return s
}

type UpdateTemplateResponseBody struct {
	// 主机ID
	HostId *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	// 请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateTemplateResponseBody) SetHostId(v string) *UpdateTemplateResponseBody {
	s.HostId = &v
	return s
}

func (s *UpdateTemplateResponseBody) SetRequestId(v string) *UpdateTemplateResponseBody {
	s.RequestId = &v
	return s
}

type UpdateTemplateResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateTemplateResponse) GoString() string {
	return s.String()
}

func (s *UpdateTemplateResponse) SetHeaders(v map[string]*string) *UpdateTemplateResponse {
	s.Headers = v
	return s
}

func (s *UpdateTemplateResponse) SetBody(v *UpdateTemplateResponseBody) *UpdateTemplateResponse {
	s.Body = v
	return s
}

type UpdateWorkspaceRequest struct {
	// 工作空间名称
	Workspace *string `json:"Workspace,omitempty" xml:"Workspace,omitempty"`
	// 工作空间描述
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 工作空间OSS工作路径
	OssRoot *string `json:"OssRoot,omitempty" xml:"OssRoot,omitempty"`
	// 工作空间内任务生命周期
	JobLifecycle *int32 `json:"JobLifecycle,omitempty" xml:"JobLifecycle,omitempty"`
	// 工作空间内Ram角色
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// 工作空间标签
	Labels *string `json:"Labels,omitempty" xml:"Labels,omitempty"`
}

func (s UpdateWorkspaceRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateWorkspaceRequest) GoString() string {
	return s.String()
}

func (s *UpdateWorkspaceRequest) SetWorkspace(v string) *UpdateWorkspaceRequest {
	s.Workspace = &v
	return s
}

func (s *UpdateWorkspaceRequest) SetDescription(v string) *UpdateWorkspaceRequest {
	s.Description = &v
	return s
}

func (s *UpdateWorkspaceRequest) SetOssRoot(v string) *UpdateWorkspaceRequest {
	s.OssRoot = &v
	return s
}

func (s *UpdateWorkspaceRequest) SetJobLifecycle(v int32) *UpdateWorkspaceRequest {
	s.JobLifecycle = &v
	return s
}

func (s *UpdateWorkspaceRequest) SetRole(v string) *UpdateWorkspaceRequest {
	s.Role = &v
	return s
}

func (s *UpdateWorkspaceRequest) SetLabels(v string) *UpdateWorkspaceRequest {
	s.Labels = &v
	return s
}

type UpdateWorkspaceResponseBody struct {
	// 主机ID
	HostId *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	// 请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateWorkspaceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateWorkspaceResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateWorkspaceResponseBody) SetHostId(v string) *UpdateWorkspaceResponseBody {
	s.HostId = &v
	return s
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

type UploadEntityRequest struct {
	Workspace     *string `json:"Workspace,omitempty" xml:"Workspace,omitempty"`
	EntityTSVFile *string `json:"EntityTSVFile,omitempty" xml:"EntityTSVFile,omitempty"`
}

func (s UploadEntityRequest) String() string {
	return tea.Prettify(s)
}

func (s UploadEntityRequest) GoString() string {
	return s.String()
}

func (s *UploadEntityRequest) SetWorkspace(v string) *UploadEntityRequest {
	s.Workspace = &v
	return s
}

func (s *UploadEntityRequest) SetEntityTSVFile(v string) *UploadEntityRequest {
	s.EntityTSVFile = &v
	return s
}

type UploadEntityResponseBody struct {
	// 主机ID
	HostId *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	// 请求ID
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Workspace  *string `json:"Workspace,omitempty" xml:"Workspace,omitempty"`
	EntityType *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
}

func (s UploadEntityResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UploadEntityResponseBody) GoString() string {
	return s.String()
}

func (s *UploadEntityResponseBody) SetHostId(v string) *UploadEntityResponseBody {
	s.HostId = &v
	return s
}

func (s *UploadEntityResponseBody) SetRequestId(v string) *UploadEntityResponseBody {
	s.RequestId = &v
	return s
}

func (s *UploadEntityResponseBody) SetWorkspace(v string) *UploadEntityResponseBody {
	s.Workspace = &v
	return s
}

func (s *UploadEntityResponseBody) SetEntityType(v string) *UploadEntityResponseBody {
	s.EntityType = &v
	return s
}

type UploadEntityResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UploadEntityResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UploadEntityResponse) String() string {
	return tea.Prettify(s)
}

func (s UploadEntityResponse) GoString() string {
	return s.String()
}

func (s *UploadEntityResponse) SetHeaders(v map[string]*string) *UploadEntityResponse {
	s.Headers = v
	return s
}

func (s *UploadEntityResponse) SetBody(v *UploadEntityResponseBody) *UploadEntityResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("easygene"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) AbortRunWithOptions(request *AbortRunRequest, runtime *util.RuntimeOptions) (_result *AbortRunResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AbortRunResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AbortRun"), tea.String("2021-03-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AbortRun(request *AbortRunRequest) (_result *AbortRunResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AbortRunResponse{}
	_body, _err := client.AbortRunWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AbortSubmissionWithOptions(request *AbortSubmissionRequest, runtime *util.RuntimeOptions) (_result *AbortSubmissionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AbortSubmissionResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AbortSubmission"), tea.String("2021-03-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AbortSubmission(request *AbortSubmissionRequest) (_result *AbortSubmissionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AbortSubmissionResponse{}
	_body, _err := client.AbortSubmissionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateAppWithOptions(tmpReq *CreateAppRequest, runtime *util.RuntimeOptions) (_result *CreateAppResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateAppShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Dependencies)) {
		request.DependenciesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Dependencies, tea.String("Dependencies"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Configs)) {
		request.ConfigsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Configs, tea.String("Configs"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateAppResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateApp"), tea.String("2021-03-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) CreateEntityWithOptions(tmpReq *CreateEntityRequest, runtime *util.RuntimeOptions) (_result *CreateEntityResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateEntityShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.EntityItems)) {
		request.EntityItemsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.EntityItems, tea.String("EntityItems"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateEntityResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateEntity"), tea.String("2021-03-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateEntity(request *CreateEntityRequest) (_result *CreateEntityResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateEntityResponse{}
	_body, _err := client.CreateEntityWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateRunWithOptions(request *CreateRunRequest, runtime *util.RuntimeOptions) (_result *CreateRunResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateRunResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateRun"), tea.String("2021-03-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateRun(request *CreateRunRequest) (_result *CreateRunResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateRunResponse{}
	_body, _err := client.CreateRunWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateSubmissionWithOptions(tmpReq *CreateSubmissionRequest, runtime *util.RuntimeOptions) (_result *CreateSubmissionResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateSubmissionShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.EntityNames)) {
		request.EntityNamesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.EntityNames, tea.String("EntityNames"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateSubmissionResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateSubmission"), tea.String("2021-03-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateSubmission(request *CreateSubmissionRequest) (_result *CreateSubmissionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateSubmissionResponse{}
	_body, _err := client.CreateSubmissionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateTemplateWithOptions(tmpReq *CreateTemplateRequest, runtime *util.RuntimeOptions) (_result *CreateTemplateResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateTemplateShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Inputs)) {
		request.InputsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Inputs, tea.String("Inputs"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Outputs)) {
		request.OutputsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Outputs, tea.String("Outputs"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateTemplateResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateTemplate"), tea.String("2021-03-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateTemplate(request *CreateTemplateRequest) (_result *CreateTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateTemplateResponse{}
	_body, _err := client.CreateTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateWorkspaceWithOptions(request *CreateWorkspaceRequest, runtime *util.RuntimeOptions) (_result *CreateWorkspaceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateWorkspaceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateWorkspace"), tea.String("2021-03-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateWorkspace(request *CreateWorkspaceRequest) (_result *CreateWorkspaceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateWorkspaceResponse{}
	_body, _err := client.CreateWorkspaceWithOptions(request, runtime)
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
	_body, _err := client.DoRPCRequest(tea.String("DeleteApp"), tea.String("2021-03-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) DeleteEntityItemsWithOptions(tmpReq *DeleteEntityItemsRequest, runtime *util.RuntimeOptions) (_result *DeleteEntityItemsResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &DeleteEntityItemsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.EntityNames)) {
		request.EntityNamesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.EntityNames, tea.String("EntityNames"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteEntityItemsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteEntityItems"), tea.String("2021-03-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteEntityItems(request *DeleteEntityItemsRequest) (_result *DeleteEntityItemsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteEntityItemsResponse{}
	_body, _err := client.DeleteEntityItemsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteRunWithOptions(request *DeleteRunRequest, runtime *util.RuntimeOptions) (_result *DeleteRunResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteRunResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteRun"), tea.String("2021-03-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteRun(request *DeleteRunRequest) (_result *DeleteRunResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteRunResponse{}
	_body, _err := client.DeleteRunWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteSubmissionWithOptions(request *DeleteSubmissionRequest, runtime *util.RuntimeOptions) (_result *DeleteSubmissionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteSubmissionResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteSubmission"), tea.String("2021-03-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteSubmission(request *DeleteSubmissionRequest) (_result *DeleteSubmissionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteSubmissionResponse{}
	_body, _err := client.DeleteSubmissionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteTemplateWithOptions(request *DeleteTemplateRequest, runtime *util.RuntimeOptions) (_result *DeleteTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteTemplateResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteTemplate"), tea.String("2021-03-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteTemplate(request *DeleteTemplateRequest) (_result *DeleteTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteTemplateResponse{}
	_body, _err := client.DeleteTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteWorkspaceWithOptions(request *DeleteWorkspaceRequest, runtime *util.RuntimeOptions) (_result *DeleteWorkspaceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteWorkspaceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteWorkspace"), tea.String("2021-03-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteWorkspace(request *DeleteWorkspaceRequest) (_result *DeleteWorkspaceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteWorkspaceResponse{}
	_body, _err := client.DeleteWorkspaceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DownloadEntityWithOptions(tmpReq *DownloadEntityRequest, runtime *util.RuntimeOptions) (_result *DownloadEntityResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &DownloadEntityShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.EntityNames)) {
		request.EntityNamesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.EntityNames, tea.String("EntityNames"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DownloadEntityResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DownloadEntity"), tea.String("2021-03-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DownloadEntity(request *DownloadEntityRequest) (_result *DownloadEntityResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DownloadEntityResponse{}
	_body, _err := client.DownloadEntityWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAppWithOptions(request *GetAppRequest, runtime *util.RuntimeOptions) (_result *GetAppResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetAppResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetApp"), tea.String("2021-03-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetApp(request *GetAppRequest) (_result *GetAppResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAppResponse{}
	_body, _err := client.GetAppWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetEntityWithOptions(request *GetEntityRequest, runtime *util.RuntimeOptions) (_result *GetEntityResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetEntityResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetEntity"), tea.String("2021-03-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetEntity(request *GetEntityRequest) (_result *GetEntityResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetEntityResponse{}
	_body, _err := client.GetEntityWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetGlobalAppWithOptions(request *GetGlobalAppRequest, runtime *util.RuntimeOptions) (_result *GetGlobalAppResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &GetGlobalAppResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetGlobalApp"), tea.String("2021-03-15"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetGlobalApp(request *GetGlobalAppRequest) (_result *GetGlobalAppResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetGlobalAppResponse{}
	_body, _err := client.GetGlobalAppWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetPublicDatasetWithOptions(tmpReq *GetPublicDatasetRequest, runtime *util.RuntimeOptions) (_result *GetPublicDatasetResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &GetPublicDatasetShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Attributes)) {
		request.AttributesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Attributes, tea.String("Attributes"), tea.String("json"))
	}

	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &GetPublicDatasetResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetPublicDataset"), tea.String("2021-03-15"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetPublicDataset(request *GetPublicDatasetRequest) (_result *GetPublicDatasetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetPublicDatasetResponse{}
	_body, _err := client.GetPublicDatasetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetPublicDatasetEntityWithOptions(request *GetPublicDatasetEntityRequest, runtime *util.RuntimeOptions) (_result *GetPublicDatasetEntityResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetPublicDatasetEntityResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetPublicDatasetEntity"), tea.String("2021-03-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetPublicDatasetEntity(request *GetPublicDatasetEntityRequest) (_result *GetPublicDatasetEntityResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetPublicDatasetEntityResponse{}
	_body, _err := client.GetPublicDatasetEntityWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetRunWithOptions(request *GetRunRequest, runtime *util.RuntimeOptions) (_result *GetRunResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &GetRunResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetRun"), tea.String("2021-03-15"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetRun(request *GetRunRequest) (_result *GetRunResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetRunResponse{}
	_body, _err := client.GetRunWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetSubmissionWithOptions(request *GetSubmissionRequest, runtime *util.RuntimeOptions) (_result *GetSubmissionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &GetSubmissionResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetSubmission"), tea.String("2021-03-15"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetSubmission(request *GetSubmissionRequest) (_result *GetSubmissionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetSubmissionResponse{}
	_body, _err := client.GetSubmissionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetTemplateWithOptions(request *GetTemplateRequest, runtime *util.RuntimeOptions) (_result *GetTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetTemplateResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetTemplate"), tea.String("2021-03-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetTemplate(request *GetTemplateRequest) (_result *GetTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetTemplateResponse{}
	_body, _err := client.GetTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetWorkspaceWithOptions(request *GetWorkspaceRequest, runtime *util.RuntimeOptions) (_result *GetWorkspaceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetWorkspaceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetWorkspace"), tea.String("2021-03-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetWorkspace(request *GetWorkspaceRequest) (_result *GetWorkspaceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetWorkspaceResponse{}
	_body, _err := client.GetWorkspaceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) InstallGlobalAppWithOptions(request *InstallGlobalAppRequest, runtime *util.RuntimeOptions) (_result *InstallGlobalAppResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &InstallGlobalAppResponse{}
	_body, _err := client.DoRPCRequest(tea.String("InstallGlobalApp"), tea.String("2021-03-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) InstallGlobalApp(request *InstallGlobalAppRequest) (_result *InstallGlobalAppResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &InstallGlobalAppResponse{}
	_body, _err := client.InstallGlobalAppWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListAppsWithOptions(request *ListAppsRequest, runtime *util.RuntimeOptions) (_result *ListAppsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListAppsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListApps"), tea.String("2021-03-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListApps(request *ListAppsRequest) (_result *ListAppsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListAppsResponse{}
	_body, _err := client.ListAppsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListAuthorizedSoftwareWithOptions(request *ListAuthorizedSoftwareRequest, runtime *util.RuntimeOptions) (_result *ListAuthorizedSoftwareResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &ListAuthorizedSoftwareResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListAuthorizedSoftware"), tea.String("2021-03-15"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListAuthorizedSoftware(request *ListAuthorizedSoftwareRequest) (_result *ListAuthorizedSoftwareResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListAuthorizedSoftwareResponse{}
	_body, _err := client.ListAuthorizedSoftwareWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListContainerImagesWithOptions(request *ListContainerImagesRequest, runtime *util.RuntimeOptions) (_result *ListContainerImagesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &ListContainerImagesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListContainerImages"), tea.String("2021-03-15"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListContainerImages(request *ListContainerImagesRequest) (_result *ListContainerImagesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListContainerImagesResponse{}
	_body, _err := client.ListContainerImagesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListEntitiesWithOptions(request *ListEntitiesRequest, runtime *util.RuntimeOptions) (_result *ListEntitiesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &ListEntitiesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListEntities"), tea.String("2021-03-15"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListEntities(request *ListEntitiesRequest) (_result *ListEntitiesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListEntitiesResponse{}
	_body, _err := client.ListEntitiesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListEntityItemsWithOptions(request *ListEntityItemsRequest, runtime *util.RuntimeOptions) (_result *ListEntityItemsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListEntityItemsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListEntityItems"), tea.String("2021-03-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListEntityItems(request *ListEntityItemsRequest) (_result *ListEntityItemsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListEntityItemsResponse{}
	_body, _err := client.ListEntityItemsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListGlobalAppsWithOptions(request *ListGlobalAppsRequest, runtime *util.RuntimeOptions) (_result *ListGlobalAppsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &ListGlobalAppsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListGlobalApps"), tea.String("2021-03-15"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListGlobalApps(request *ListGlobalAppsRequest) (_result *ListGlobalAppsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListGlobalAppsResponse{}
	_body, _err := client.ListGlobalAppsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListPublicDatasetWithOptions(request *ListPublicDatasetRequest, runtime *util.RuntimeOptions) (_result *ListPublicDatasetResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &ListPublicDatasetResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListPublicDataset"), tea.String("2021-03-15"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListPublicDataset(request *ListPublicDatasetRequest) (_result *ListPublicDatasetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListPublicDatasetResponse{}
	_body, _err := client.ListPublicDatasetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListPublicDatasetEntitiesWithOptions(request *ListPublicDatasetEntitiesRequest, runtime *util.RuntimeOptions) (_result *ListPublicDatasetEntitiesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &ListPublicDatasetEntitiesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListPublicDatasetEntities"), tea.String("2021-03-15"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListPublicDatasetEntities(request *ListPublicDatasetEntitiesRequest) (_result *ListPublicDatasetEntitiesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListPublicDatasetEntitiesResponse{}
	_body, _err := client.ListPublicDatasetEntitiesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListPublicDatasetEntityItemsWithOptions(request *ListPublicDatasetEntityItemsRequest, runtime *util.RuntimeOptions) (_result *ListPublicDatasetEntityItemsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &ListPublicDatasetEntityItemsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListPublicDatasetEntityItems"), tea.String("2021-03-15"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListPublicDatasetEntityItems(request *ListPublicDatasetEntityItemsRequest) (_result *ListPublicDatasetEntityItemsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListPublicDatasetEntityItemsResponse{}
	_body, _err := client.ListPublicDatasetEntityItemsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListPublicDatasetTagsWithOptions(request *ListPublicDatasetTagsRequest, runtime *util.RuntimeOptions) (_result *ListPublicDatasetTagsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &ListPublicDatasetTagsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListPublicDatasetTags"), tea.String("2021-03-15"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListPublicDatasetTags(request *ListPublicDatasetTagsRequest) (_result *ListPublicDatasetTagsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListPublicDatasetTagsResponse{}
	_body, _err := client.ListPublicDatasetTagsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListRegionsWithOptions(runtime *util.RuntimeOptions) (_result *ListRegionsResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	_result = &ListRegionsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListRegions"), tea.String("2021-03-15"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListRegions() (_result *ListRegionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListRegionsResponse{}
	_body, _err := client.ListRegionsWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListRunsWithOptions(request *ListRunsRequest, runtime *util.RuntimeOptions) (_result *ListRunsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &ListRunsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListRuns"), tea.String("2021-03-15"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListRuns(request *ListRunsRequest) (_result *ListRunsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListRunsResponse{}
	_body, _err := client.ListRunsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListSubmissionsWithOptions(request *ListSubmissionsRequest, runtime *util.RuntimeOptions) (_result *ListSubmissionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &ListSubmissionsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListSubmissions"), tea.String("2021-03-15"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListSubmissions(request *ListSubmissionsRequest) (_result *ListSubmissionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListSubmissionsResponse{}
	_body, _err := client.ListSubmissionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListTemplatesWithOptions(request *ListTemplatesRequest, runtime *util.RuntimeOptions) (_result *ListTemplatesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListTemplatesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListTemplates"), tea.String("2021-03-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListTemplates(request *ListTemplatesRequest) (_result *ListTemplatesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTemplatesResponse{}
	_body, _err := client.ListTemplatesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListUserActiveRunsWithOptions(request *ListUserActiveRunsRequest, runtime *util.RuntimeOptions) (_result *ListUserActiveRunsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &ListUserActiveRunsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListUserActiveRuns"), tea.String("2021-03-15"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListUserActiveRuns(request *ListUserActiveRunsRequest) (_result *ListUserActiveRunsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListUserActiveRunsResponse{}
	_body, _err := client.ListUserActiveRunsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListWorkspacesWithOptions(request *ListWorkspacesRequest, runtime *util.RuntimeOptions) (_result *ListWorkspacesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &ListWorkspacesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListWorkspaces"), tea.String("2021-03-15"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListWorkspaces(request *ListWorkspacesRequest) (_result *ListWorkspacesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListWorkspacesResponse{}
	_body, _err := client.ListWorkspacesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ParseAppInputsWithOptions(tmpReq *ParseAppInputsRequest, runtime *util.RuntimeOptions) (_result *ParseAppInputsResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ParseAppInputsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Dependencies)) {
		request.DependenciesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Dependencies, tea.String("Dependencies"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ParseAppInputsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ParseAppInputs"), tea.String("2021-03-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ParseAppInputs(request *ParseAppInputsRequest) (_result *ParseAppInputsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ParseAppInputsResponse{}
	_body, _err := client.ParseAppInputsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ResumeSubmissionWithOptions(request *ResumeSubmissionRequest, runtime *util.RuntimeOptions) (_result *ResumeSubmissionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ResumeSubmissionResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ResumeSubmission"), tea.String("2021-03-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ResumeSubmission(request *ResumeSubmissionRequest) (_result *ResumeSubmissionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ResumeSubmissionResponse{}
	_body, _err := client.ResumeSubmissionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateEntityWithOptions(tmpReq *UpdateEntityRequest, runtime *util.RuntimeOptions) (_result *UpdateEntityResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &UpdateEntityShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.EntityItems)) {
		request.EntityItemsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.EntityItems, tea.String("EntityItems"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateEntityResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateEntity"), tea.String("2021-03-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateEntity(request *UpdateEntityRequest) (_result *UpdateEntityResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateEntityResponse{}
	_body, _err := client.UpdateEntityWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateTemplateWithOptions(tmpReq *UpdateTemplateRequest, runtime *util.RuntimeOptions) (_result *UpdateTemplateResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &UpdateTemplateShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Inputs)) {
		request.InputsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Inputs, tea.String("Inputs"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Outputs)) {
		request.OutputsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Outputs, tea.String("Outputs"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateTemplateResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateTemplate"), tea.String("2021-03-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateTemplate(request *UpdateTemplateRequest) (_result *UpdateTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateTemplateResponse{}
	_body, _err := client.UpdateTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateWorkspaceWithOptions(request *UpdateWorkspaceRequest, runtime *util.RuntimeOptions) (_result *UpdateWorkspaceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateWorkspaceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateWorkspace"), tea.String("2021-03-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateWorkspace(request *UpdateWorkspaceRequest) (_result *UpdateWorkspaceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateWorkspaceResponse{}
	_body, _err := client.UpdateWorkspaceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UploadEntityWithOptions(request *UploadEntityRequest, runtime *util.RuntimeOptions) (_result *UploadEntityResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UploadEntityResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UploadEntity"), tea.String("2021-03-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UploadEntity(request *UploadEntityRequest) (_result *UploadEntityResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UploadEntityResponse{}
	_body, _err := client.UploadEntityWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
