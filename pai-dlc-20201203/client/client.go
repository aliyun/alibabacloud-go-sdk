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

type AIMasterMessage struct {
	Extended        *string `json:"Extended,omitempty" xml:"Extended,omitempty"`
	JobRestartCount *int32  `json:"JobRestartCount,omitempty" xml:"JobRestartCount,omitempty"`
	MessageContent  *string `json:"MessageContent,omitempty" xml:"MessageContent,omitempty"`
	MessageEvent    *string `json:"MessageEvent,omitempty" xml:"MessageEvent,omitempty"`
	MessageVersion  *int32  `json:"MessageVersion,omitempty" xml:"MessageVersion,omitempty"`
	RestartType     *string `json:"RestartType,omitempty" xml:"RestartType,omitempty"`
}

func (s AIMasterMessage) String() string {
	return tea.Prettify(s)
}

func (s AIMasterMessage) GoString() string {
	return s.String()
}

func (s *AIMasterMessage) SetExtended(v string) *AIMasterMessage {
	s.Extended = &v
	return s
}

func (s *AIMasterMessage) SetJobRestartCount(v int32) *AIMasterMessage {
	s.JobRestartCount = &v
	return s
}

func (s *AIMasterMessage) SetMessageContent(v string) *AIMasterMessage {
	s.MessageContent = &v
	return s
}

func (s *AIMasterMessage) SetMessageEvent(v string) *AIMasterMessage {
	s.MessageEvent = &v
	return s
}

func (s *AIMasterMessage) SetMessageVersion(v int32) *AIMasterMessage {
	s.MessageVersion = &v
	return s
}

func (s *AIMasterMessage) SetRestartType(v string) *AIMasterMessage {
	s.RestartType = &v
	return s
}

type AliyunAccounts struct {
	AliyunUid     *string `json:"AliyunUid,omitempty" xml:"AliyunUid,omitempty"`
	EmployeeId    *string `json:"EmployeeId,omitempty" xml:"EmployeeId,omitempty"`
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	GmtModifyTime *string `json:"GmtModifyTime,omitempty" xml:"GmtModifyTime,omitempty"`
}

func (s AliyunAccounts) String() string {
	return tea.Prettify(s)
}

func (s AliyunAccounts) GoString() string {
	return s.String()
}

func (s *AliyunAccounts) SetAliyunUid(v string) *AliyunAccounts {
	s.AliyunUid = &v
	return s
}

func (s *AliyunAccounts) SetEmployeeId(v string) *AliyunAccounts {
	s.EmployeeId = &v
	return s
}

func (s *AliyunAccounts) SetGmtCreateTime(v string) *AliyunAccounts {
	s.GmtCreateTime = &v
	return s
}

func (s *AliyunAccounts) SetGmtModifyTime(v string) *AliyunAccounts {
	s.GmtModifyTime = &v
	return s
}

type CodeSourceItem struct {
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
	UserId              *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s CodeSourceItem) String() string {
	return tea.Prettify(s)
}

func (s CodeSourceItem) GoString() string {
	return s.String()
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

type ContainerSpec struct {
	Args       []*string             `json:"Args,omitempty" xml:"Args,omitempty" type:"Repeated"`
	Command    []*string             `json:"Command,omitempty" xml:"Command,omitempty" type:"Repeated"`
	Env        []*EnvVar             `json:"Env,omitempty" xml:"Env,omitempty" type:"Repeated"`
	Image      *string               `json:"Image,omitempty" xml:"Image,omitempty"`
	Name       *string               `json:"Name,omitempty" xml:"Name,omitempty"`
	Resources  *ResourceRequirements `json:"Resources,omitempty" xml:"Resources,omitempty"`
	WorkingDir *string               `json:"WorkingDir,omitempty" xml:"WorkingDir,omitempty"`
}

func (s ContainerSpec) String() string {
	return tea.Prettify(s)
}

func (s ContainerSpec) GoString() string {
	return s.String()
}

func (s *ContainerSpec) SetArgs(v []*string) *ContainerSpec {
	s.Args = v
	return s
}

func (s *ContainerSpec) SetCommand(v []*string) *ContainerSpec {
	s.Command = v
	return s
}

func (s *ContainerSpec) SetEnv(v []*EnvVar) *ContainerSpec {
	s.Env = v
	return s
}

func (s *ContainerSpec) SetImage(v string) *ContainerSpec {
	s.Image = &v
	return s
}

func (s *ContainerSpec) SetName(v string) *ContainerSpec {
	s.Name = &v
	return s
}

func (s *ContainerSpec) SetResources(v *ResourceRequirements) *ContainerSpec {
	s.Resources = v
	return s
}

func (s *ContainerSpec) SetWorkingDir(v string) *ContainerSpec {
	s.WorkingDir = &v
	return s
}

type DataSourceItem struct {
	DataSourceId   *string `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	DataSourceType *string `json:"DataSourceType,omitempty" xml:"DataSourceType,omitempty"`
	Description    *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DisplayName    *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	Endpoint       *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	FileSystemId   *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	GmtCreateTime  *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	GmtModifyTime  *string `json:"GmtModifyTime,omitempty" xml:"GmtModifyTime,omitempty"`
	MountPath      *string `json:"MountPath,omitempty" xml:"MountPath,omitempty"`
	Options        *string `json:"Options,omitempty" xml:"Options,omitempty"`
	Path           *string `json:"Path,omitempty" xml:"Path,omitempty"`
	UserId         *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s DataSourceItem) String() string {
	return tea.Prettify(s)
}

func (s DataSourceItem) GoString() string {
	return s.String()
}

func (s *DataSourceItem) SetDataSourceId(v string) *DataSourceItem {
	s.DataSourceId = &v
	return s
}

func (s *DataSourceItem) SetDataSourceType(v string) *DataSourceItem {
	s.DataSourceType = &v
	return s
}

func (s *DataSourceItem) SetDescription(v string) *DataSourceItem {
	s.Description = &v
	return s
}

func (s *DataSourceItem) SetDisplayName(v string) *DataSourceItem {
	s.DisplayName = &v
	return s
}

func (s *DataSourceItem) SetEndpoint(v string) *DataSourceItem {
	s.Endpoint = &v
	return s
}

func (s *DataSourceItem) SetFileSystemId(v string) *DataSourceItem {
	s.FileSystemId = &v
	return s
}

func (s *DataSourceItem) SetGmtCreateTime(v string) *DataSourceItem {
	s.GmtCreateTime = &v
	return s
}

func (s *DataSourceItem) SetGmtModifyTime(v string) *DataSourceItem {
	s.GmtModifyTime = &v
	return s
}

func (s *DataSourceItem) SetMountPath(v string) *DataSourceItem {
	s.MountPath = &v
	return s
}

func (s *DataSourceItem) SetOptions(v string) *DataSourceItem {
	s.Options = &v
	return s
}

func (s *DataSourceItem) SetPath(v string) *DataSourceItem {
	s.Path = &v
	return s
}

func (s *DataSourceItem) SetUserId(v string) *DataSourceItem {
	s.UserId = &v
	return s
}

type DebuggerConfig struct {
	Content          *string `json:"Content,omitempty" xml:"Content,omitempty"`
	DebuggerConfigId *string `json:"DebuggerConfigId,omitempty" xml:"DebuggerConfigId,omitempty"`
	Description      *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DisplayName      *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	GmtCreateTime    *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	GmtModifyTime    *string `json:"GmtModifyTime,omitempty" xml:"GmtModifyTime,omitempty"`
}

func (s DebuggerConfig) String() string {
	return tea.Prettify(s)
}

func (s DebuggerConfig) GoString() string {
	return s.String()
}

func (s *DebuggerConfig) SetContent(v string) *DebuggerConfig {
	s.Content = &v
	return s
}

func (s *DebuggerConfig) SetDebuggerConfigId(v string) *DebuggerConfig {
	s.DebuggerConfigId = &v
	return s
}

func (s *DebuggerConfig) SetDescription(v string) *DebuggerConfig {
	s.Description = &v
	return s
}

func (s *DebuggerConfig) SetDisplayName(v string) *DebuggerConfig {
	s.DisplayName = &v
	return s
}

func (s *DebuggerConfig) SetGmtCreateTime(v string) *DebuggerConfig {
	s.GmtCreateTime = &v
	return s
}

func (s *DebuggerConfig) SetGmtModifyTime(v string) *DebuggerConfig {
	s.GmtModifyTime = &v
	return s
}

type DebuggerJob struct {
	DebuggerJobId    *string `json:"DebuggerJobId,omitempty" xml:"DebuggerJobId,omitempty"`
	DisplayName      *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	Duration         *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	GmtCreateTime    *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	GmtFailedTime    *string `json:"GmtFailedTime,omitempty" xml:"GmtFailedTime,omitempty"`
	GmtFinishTime    *string `json:"GmtFinishTime,omitempty" xml:"GmtFinishTime,omitempty"`
	GmtRunningTime   *string `json:"GmtRunningTime,omitempty" xml:"GmtRunningTime,omitempty"`
	GmtStoppedTime   *string `json:"GmtStoppedTime,omitempty" xml:"GmtStoppedTime,omitempty"`
	GmtSubmittedTime *string `json:"GmtSubmittedTime,omitempty" xml:"GmtSubmittedTime,omitempty"`
	GmtSucceedTime   *string `json:"GmtSucceedTime,omitempty" xml:"GmtSucceedTime,omitempty"`
	Status           *string `json:"Status,omitempty" xml:"Status,omitempty"`
	UserId           *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	WorkspaceId      *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
	WorkspaceName    *string `json:"WorkspaceName,omitempty" xml:"WorkspaceName,omitempty"`
}

func (s DebuggerJob) String() string {
	return tea.Prettify(s)
}

func (s DebuggerJob) GoString() string {
	return s.String()
}

func (s *DebuggerJob) SetDebuggerJobId(v string) *DebuggerJob {
	s.DebuggerJobId = &v
	return s
}

func (s *DebuggerJob) SetDisplayName(v string) *DebuggerJob {
	s.DisplayName = &v
	return s
}

func (s *DebuggerJob) SetDuration(v string) *DebuggerJob {
	s.Duration = &v
	return s
}

func (s *DebuggerJob) SetGmtCreateTime(v string) *DebuggerJob {
	s.GmtCreateTime = &v
	return s
}

func (s *DebuggerJob) SetGmtFailedTime(v string) *DebuggerJob {
	s.GmtFailedTime = &v
	return s
}

func (s *DebuggerJob) SetGmtFinishTime(v string) *DebuggerJob {
	s.GmtFinishTime = &v
	return s
}

func (s *DebuggerJob) SetGmtRunningTime(v string) *DebuggerJob {
	s.GmtRunningTime = &v
	return s
}

func (s *DebuggerJob) SetGmtStoppedTime(v string) *DebuggerJob {
	s.GmtStoppedTime = &v
	return s
}

func (s *DebuggerJob) SetGmtSubmittedTime(v string) *DebuggerJob {
	s.GmtSubmittedTime = &v
	return s
}

func (s *DebuggerJob) SetGmtSucceedTime(v string) *DebuggerJob {
	s.GmtSucceedTime = &v
	return s
}

func (s *DebuggerJob) SetStatus(v string) *DebuggerJob {
	s.Status = &v
	return s
}

func (s *DebuggerJob) SetUserId(v string) *DebuggerJob {
	s.UserId = &v
	return s
}

func (s *DebuggerJob) SetWorkspaceId(v string) *DebuggerJob {
	s.WorkspaceId = &v
	return s
}

func (s *DebuggerJob) SetWorkspaceName(v string) *DebuggerJob {
	s.WorkspaceName = &v
	return s
}

type DebuggerJobIssue struct {
	DebuggerJobIssue   *string `json:"DebuggerJobIssue,omitempty" xml:"DebuggerJobIssue,omitempty"`
	GmtCreateTime      *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	JobDebuggerIssueId *string `json:"JobDebuggerIssueId,omitempty" xml:"JobDebuggerIssueId,omitempty"`
	JobId              *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	ReasonCode         *string `json:"ReasonCode,omitempty" xml:"ReasonCode,omitempty"`
	ReasonMessage      *string `json:"ReasonMessage,omitempty" xml:"ReasonMessage,omitempty"`
	RuleName           *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
}

func (s DebuggerJobIssue) String() string {
	return tea.Prettify(s)
}

func (s DebuggerJobIssue) GoString() string {
	return s.String()
}

func (s *DebuggerJobIssue) SetDebuggerJobIssue(v string) *DebuggerJobIssue {
	s.DebuggerJobIssue = &v
	return s
}

func (s *DebuggerJobIssue) SetGmtCreateTime(v string) *DebuggerJobIssue {
	s.GmtCreateTime = &v
	return s
}

func (s *DebuggerJobIssue) SetJobDebuggerIssueId(v string) *DebuggerJobIssue {
	s.JobDebuggerIssueId = &v
	return s
}

func (s *DebuggerJobIssue) SetJobId(v string) *DebuggerJobIssue {
	s.JobId = &v
	return s
}

func (s *DebuggerJobIssue) SetReasonCode(v string) *DebuggerJobIssue {
	s.ReasonCode = &v
	return s
}

func (s *DebuggerJobIssue) SetReasonMessage(v string) *DebuggerJobIssue {
	s.ReasonMessage = &v
	return s
}

func (s *DebuggerJobIssue) SetRuleName(v string) *DebuggerJobIssue {
	s.RuleName = &v
	return s
}

type DebuggerResult struct {
	DebuggerConfigContent *string `json:"DebuggerConfigContent,omitempty" xml:"DebuggerConfigContent,omitempty"`
	DebuggerJobIssues     *string `json:"DebuggerJobIssues,omitempty" xml:"DebuggerJobIssues,omitempty"`
	DebuggerJobStatus     *string `json:"DebuggerJobStatus,omitempty" xml:"DebuggerJobStatus,omitempty"`
	DebuggerReportURL     *string `json:"DebuggerReportURL,omitempty" xml:"DebuggerReportURL,omitempty"`
	JobDisplayName        *string `json:"JobDisplayName,omitempty" xml:"JobDisplayName,omitempty"`
	JobId                 *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	JobUserId             *string `json:"JobUserId,omitempty" xml:"JobUserId,omitempty"`
}

func (s DebuggerResult) String() string {
	return tea.Prettify(s)
}

func (s DebuggerResult) GoString() string {
	return s.String()
}

func (s *DebuggerResult) SetDebuggerConfigContent(v string) *DebuggerResult {
	s.DebuggerConfigContent = &v
	return s
}

func (s *DebuggerResult) SetDebuggerJobIssues(v string) *DebuggerResult {
	s.DebuggerJobIssues = &v
	return s
}

func (s *DebuggerResult) SetDebuggerJobStatus(v string) *DebuggerResult {
	s.DebuggerJobStatus = &v
	return s
}

func (s *DebuggerResult) SetDebuggerReportURL(v string) *DebuggerResult {
	s.DebuggerReportURL = &v
	return s
}

func (s *DebuggerResult) SetJobDisplayName(v string) *DebuggerResult {
	s.JobDisplayName = &v
	return s
}

func (s *DebuggerResult) SetJobId(v string) *DebuggerResult {
	s.JobId = &v
	return s
}

func (s *DebuggerResult) SetJobUserId(v string) *DebuggerResult {
	s.JobUserId = &v
	return s
}

type EcsSpec struct {
	AcceleratorType        *string   `json:"AcceleratorType,omitempty" xml:"AcceleratorType,omitempty"`
	Cpu                    *int32    `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	DefaultGPUDriver       *string   `json:"DefaultGPUDriver,omitempty" xml:"DefaultGPUDriver,omitempty"`
	Gpu                    *int32    `json:"Gpu,omitempty" xml:"Gpu,omitempty"`
	GpuType                *string   `json:"GpuType,omitempty" xml:"GpuType,omitempty"`
	InstanceType           *string   `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	IsAvailable            *bool     `json:"IsAvailable,omitempty" xml:"IsAvailable,omitempty"`
	Memory                 *int32    `json:"Memory,omitempty" xml:"Memory,omitempty"`
	NonProtectSpotDiscount *float32  `json:"NonProtectSpotDiscount,omitempty" xml:"NonProtectSpotDiscount,omitempty"`
	PaymentTypes           []*string `json:"PaymentTypes,omitempty" xml:"PaymentTypes,omitempty" type:"Repeated"`
	ResourceType           *string   `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	SpotStockStatus        *string   `json:"SpotStockStatus,omitempty" xml:"SpotStockStatus,omitempty"`
	SupportedGPUDrivers    []*string `json:"SupportedGPUDrivers,omitempty" xml:"SupportedGPUDrivers,omitempty" type:"Repeated"`
}

func (s EcsSpec) String() string {
	return tea.Prettify(s)
}

func (s EcsSpec) GoString() string {
	return s.String()
}

func (s *EcsSpec) SetAcceleratorType(v string) *EcsSpec {
	s.AcceleratorType = &v
	return s
}

func (s *EcsSpec) SetCpu(v int32) *EcsSpec {
	s.Cpu = &v
	return s
}

func (s *EcsSpec) SetDefaultGPUDriver(v string) *EcsSpec {
	s.DefaultGPUDriver = &v
	return s
}

func (s *EcsSpec) SetGpu(v int32) *EcsSpec {
	s.Gpu = &v
	return s
}

func (s *EcsSpec) SetGpuType(v string) *EcsSpec {
	s.GpuType = &v
	return s
}

func (s *EcsSpec) SetInstanceType(v string) *EcsSpec {
	s.InstanceType = &v
	return s
}

func (s *EcsSpec) SetIsAvailable(v bool) *EcsSpec {
	s.IsAvailable = &v
	return s
}

func (s *EcsSpec) SetMemory(v int32) *EcsSpec {
	s.Memory = &v
	return s
}

func (s *EcsSpec) SetNonProtectSpotDiscount(v float32) *EcsSpec {
	s.NonProtectSpotDiscount = &v
	return s
}

func (s *EcsSpec) SetPaymentTypes(v []*string) *EcsSpec {
	s.PaymentTypes = v
	return s
}

func (s *EcsSpec) SetResourceType(v string) *EcsSpec {
	s.ResourceType = &v
	return s
}

func (s *EcsSpec) SetSpotStockStatus(v string) *EcsSpec {
	s.SpotStockStatus = &v
	return s
}

func (s *EcsSpec) SetSupportedGPUDrivers(v []*string) *EcsSpec {
	s.SupportedGPUDrivers = v
	return s
}

type EnvVar struct {
	Name  *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s EnvVar) String() string {
	return tea.Prettify(s)
}

func (s EnvVar) GoString() string {
	return s.String()
}

func (s *EnvVar) SetName(v string) *EnvVar {
	s.Name = &v
	return s
}

func (s *EnvVar) SetValue(v string) *EnvVar {
	s.Value = &v
	return s
}

type EventInfo struct {
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	Id      *string `json:"Id,omitempty" xml:"Id,omitempty"`
	PodId   *string `json:"PodId,omitempty" xml:"PodId,omitempty"`
	PodUid  *string `json:"PodUid,omitempty" xml:"PodUid,omitempty"`
	Time    *string `json:"Time,omitempty" xml:"Time,omitempty"`
}

func (s EventInfo) String() string {
	return tea.Prettify(s)
}

func (s EventInfo) GoString() string {
	return s.String()
}

func (s *EventInfo) SetContent(v string) *EventInfo {
	s.Content = &v
	return s
}

func (s *EventInfo) SetId(v string) *EventInfo {
	s.Id = &v
	return s
}

func (s *EventInfo) SetPodId(v string) *EventInfo {
	s.PodId = &v
	return s
}

func (s *EventInfo) SetPodUid(v string) *EventInfo {
	s.PodUid = &v
	return s
}

func (s *EventInfo) SetTime(v string) *EventInfo {
	s.Time = &v
	return s
}

type ExtraPodSpec struct {
	InitContainers         []*ContainerSpec   `json:"InitContainers,omitempty" xml:"InitContainers,omitempty" type:"Repeated"`
	PodAnnotations         map[string]*string `json:"PodAnnotations,omitempty" xml:"PodAnnotations,omitempty"`
	PodLabels              map[string]*string `json:"PodLabels,omitempty" xml:"PodLabels,omitempty"`
	SharedVolumeMountPaths []*string          `json:"SharedVolumeMountPaths,omitempty" xml:"SharedVolumeMountPaths,omitempty" type:"Repeated"`
	SideCarContainers      []*ContainerSpec   `json:"SideCarContainers,omitempty" xml:"SideCarContainers,omitempty" type:"Repeated"`
}

func (s ExtraPodSpec) String() string {
	return tea.Prettify(s)
}

func (s ExtraPodSpec) GoString() string {
	return s.String()
}

func (s *ExtraPodSpec) SetInitContainers(v []*ContainerSpec) *ExtraPodSpec {
	s.InitContainers = v
	return s
}

func (s *ExtraPodSpec) SetPodAnnotations(v map[string]*string) *ExtraPodSpec {
	s.PodAnnotations = v
	return s
}

func (s *ExtraPodSpec) SetPodLabels(v map[string]*string) *ExtraPodSpec {
	s.PodLabels = v
	return s
}

func (s *ExtraPodSpec) SetSharedVolumeMountPaths(v []*string) *ExtraPodSpec {
	s.SharedVolumeMountPaths = v
	return s
}

func (s *ExtraPodSpec) SetSideCarContainers(v []*ContainerSpec) *ExtraPodSpec {
	s.SideCarContainers = v
	return s
}

type FreeResourceClusterControlItem struct {
	ClusterID                    *string `json:"ClusterID,omitempty" xml:"ClusterID,omitempty"`
	ClusterName                  *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	CrossClusters                *bool   `json:"CrossClusters,omitempty" xml:"CrossClusters,omitempty"`
	EnableFreeResource           *bool   `json:"EnableFreeResource,omitempty" xml:"EnableFreeResource,omitempty"`
	FreeResourceClusterControlId *string `json:"FreeResourceClusterControlId,omitempty" xml:"FreeResourceClusterControlId,omitempty"`
	GmtCreateTime                *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	GmtModifyTime                *string `json:"GmtModifyTime,omitempty" xml:"GmtModifyTime,omitempty"`
	RegionID                     *string `json:"RegionID,omitempty" xml:"RegionID,omitempty"`
}

func (s FreeResourceClusterControlItem) String() string {
	return tea.Prettify(s)
}

func (s FreeResourceClusterControlItem) GoString() string {
	return s.String()
}

func (s *FreeResourceClusterControlItem) SetClusterID(v string) *FreeResourceClusterControlItem {
	s.ClusterID = &v
	return s
}

func (s *FreeResourceClusterControlItem) SetClusterName(v string) *FreeResourceClusterControlItem {
	s.ClusterName = &v
	return s
}

func (s *FreeResourceClusterControlItem) SetCrossClusters(v bool) *FreeResourceClusterControlItem {
	s.CrossClusters = &v
	return s
}

func (s *FreeResourceClusterControlItem) SetEnableFreeResource(v bool) *FreeResourceClusterControlItem {
	s.EnableFreeResource = &v
	return s
}

func (s *FreeResourceClusterControlItem) SetFreeResourceClusterControlId(v string) *FreeResourceClusterControlItem {
	s.FreeResourceClusterControlId = &v
	return s
}

func (s *FreeResourceClusterControlItem) SetGmtCreateTime(v string) *FreeResourceClusterControlItem {
	s.GmtCreateTime = &v
	return s
}

func (s *FreeResourceClusterControlItem) SetGmtModifyTime(v string) *FreeResourceClusterControlItem {
	s.GmtModifyTime = &v
	return s
}

func (s *FreeResourceClusterControlItem) SetRegionID(v string) *FreeResourceClusterControlItem {
	s.RegionID = &v
	return s
}

type FreeResourceDetail struct {
	Amount       *int32  `json:"Amount,omitempty" xml:"Amount,omitempty"`
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s FreeResourceDetail) String() string {
	return tea.Prettify(s)
}

func (s FreeResourceDetail) GoString() string {
	return s.String()
}

func (s *FreeResourceDetail) SetAmount(v int32) *FreeResourceDetail {
	s.Amount = &v
	return s
}

func (s *FreeResourceDetail) SetResourceType(v string) *FreeResourceDetail {
	s.ResourceType = &v
	return s
}

type FreeResourceItem struct {
	AvailableNumber *int64  `json:"AvailableNumber,omitempty" xml:"AvailableNumber,omitempty"`
	ClusterID       *string `json:"ClusterID,omitempty" xml:"ClusterID,omitempty"`
	ClusterName     *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	FreeResourceId  *string `json:"FreeResourceId,omitempty" xml:"FreeResourceId,omitempty"`
	GmtCreateTime   *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	GmtModifyTime   *string `json:"GmtModifyTime,omitempty" xml:"GmtModifyTime,omitempty"`
	RegionID        *string `json:"RegionID,omitempty" xml:"RegionID,omitempty"`
	ResourceType    *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s FreeResourceItem) String() string {
	return tea.Prettify(s)
}

func (s FreeResourceItem) GoString() string {
	return s.String()
}

func (s *FreeResourceItem) SetAvailableNumber(v int64) *FreeResourceItem {
	s.AvailableNumber = &v
	return s
}

func (s *FreeResourceItem) SetClusterID(v string) *FreeResourceItem {
	s.ClusterID = &v
	return s
}

func (s *FreeResourceItem) SetClusterName(v string) *FreeResourceItem {
	s.ClusterName = &v
	return s
}

func (s *FreeResourceItem) SetFreeResourceId(v string) *FreeResourceItem {
	s.FreeResourceId = &v
	return s
}

func (s *FreeResourceItem) SetGmtCreateTime(v string) *FreeResourceItem {
	s.GmtCreateTime = &v
	return s
}

func (s *FreeResourceItem) SetGmtModifyTime(v string) *FreeResourceItem {
	s.GmtModifyTime = &v
	return s
}

func (s *FreeResourceItem) SetRegionID(v string) *FreeResourceItem {
	s.RegionID = &v
	return s
}

func (s *FreeResourceItem) SetResourceType(v string) *FreeResourceItem {
	s.ResourceType = &v
	return s
}

type GPUDetail struct {
	GPU             *string `json:"GPU,omitempty" xml:"GPU,omitempty"`
	GPUType         *string `json:"GPUType,omitempty" xml:"GPUType,omitempty"`
	GPUTypeFullName *string `json:"GPUTypeFullName,omitempty" xml:"GPUTypeFullName,omitempty"`
}

func (s GPUDetail) String() string {
	return tea.Prettify(s)
}

func (s GPUDetail) GoString() string {
	return s.String()
}

func (s *GPUDetail) SetGPU(v string) *GPUDetail {
	s.GPU = &v
	return s
}

func (s *GPUDetail) SetGPUType(v string) *GPUDetail {
	s.GPUType = &v
	return s
}

func (s *GPUDetail) SetGPUTypeFullName(v string) *GPUDetail {
	s.GPUTypeFullName = &v
	return s
}

type ImageConfig struct {
	Auth           *string `json:"Auth,omitempty" xml:"Auth,omitempty"`
	DockerRegistry *string `json:"DockerRegistry,omitempty" xml:"DockerRegistry,omitempty"`
	Password       *string `json:"Password,omitempty" xml:"Password,omitempty"`
	Username       *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s ImageConfig) String() string {
	return tea.Prettify(s)
}

func (s ImageConfig) GoString() string {
	return s.String()
}

func (s *ImageConfig) SetAuth(v string) *ImageConfig {
	s.Auth = &v
	return s
}

func (s *ImageConfig) SetDockerRegistry(v string) *ImageConfig {
	s.DockerRegistry = &v
	return s
}

func (s *ImageConfig) SetPassword(v string) *ImageConfig {
	s.Password = &v
	return s
}

func (s *ImageConfig) SetUsername(v string) *ImageConfig {
	s.Username = &v
	return s
}

type ImageItem struct {
	AcceleratorType   *string `json:"AcceleratorType,omitempty" xml:"AcceleratorType,omitempty"`
	AuthorId          *string `json:"AuthorId,omitempty" xml:"AuthorId,omitempty"`
	Framework         *string `json:"Framework,omitempty" xml:"Framework,omitempty"`
	ImageProviderType *string `json:"ImageProviderType,omitempty" xml:"ImageProviderType,omitempty"`
	ImageTag          *string `json:"ImageTag,omitempty" xml:"ImageTag,omitempty"`
	ImageUrl          *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	ImageUrlVpc       *string `json:"ImageUrlVpc,omitempty" xml:"ImageUrlVpc,omitempty"`
}

func (s ImageItem) String() string {
	return tea.Prettify(s)
}

func (s ImageItem) GoString() string {
	return s.String()
}

func (s *ImageItem) SetAcceleratorType(v string) *ImageItem {
	s.AcceleratorType = &v
	return s
}

func (s *ImageItem) SetAuthorId(v string) *ImageItem {
	s.AuthorId = &v
	return s
}

func (s *ImageItem) SetFramework(v string) *ImageItem {
	s.Framework = &v
	return s
}

func (s *ImageItem) SetImageProviderType(v string) *ImageItem {
	s.ImageProviderType = &v
	return s
}

func (s *ImageItem) SetImageTag(v string) *ImageItem {
	s.ImageTag = &v
	return s
}

func (s *ImageItem) SetImageUrl(v string) *ImageItem {
	s.ImageUrl = &v
	return s
}

func (s *ImageItem) SetImageUrlVpc(v string) *ImageItem {
	s.ImageUrlVpc = &v
	return s
}

type JobDebuggerConfig struct {
	DebuggerConfigContent *string `json:"DebuggerConfigContent,omitempty" xml:"DebuggerConfigContent,omitempty"`
	DebuggerConfigId      *string `json:"DebuggerConfigId,omitempty" xml:"DebuggerConfigId,omitempty"`
	GmtCreateTime         *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	JobId                 *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s JobDebuggerConfig) String() string {
	return tea.Prettify(s)
}

func (s JobDebuggerConfig) GoString() string {
	return s.String()
}

func (s *JobDebuggerConfig) SetDebuggerConfigContent(v string) *JobDebuggerConfig {
	s.DebuggerConfigContent = &v
	return s
}

func (s *JobDebuggerConfig) SetDebuggerConfigId(v string) *JobDebuggerConfig {
	s.DebuggerConfigId = &v
	return s
}

func (s *JobDebuggerConfig) SetGmtCreateTime(v string) *JobDebuggerConfig {
	s.GmtCreateTime = &v
	return s
}

func (s *JobDebuggerConfig) SetJobId(v string) *JobDebuggerConfig {
	s.JobId = &v
	return s
}

type JobElasticSpec struct {
	AIMasterDockerImage      *string `json:"AIMasterDockerImage,omitempty" xml:"AIMasterDockerImage,omitempty"`
	AIMasterType             *string `json:"AIMasterType,omitempty" xml:"AIMasterType,omitempty"`
	EDPMaxParallelism        *int32  `json:"EDPMaxParallelism,omitempty" xml:"EDPMaxParallelism,omitempty"`
	EDPMinParallelism        *int32  `json:"EDPMinParallelism,omitempty" xml:"EDPMinParallelism,omitempty"`
	ElasticStrategy          *string `json:"ElasticStrategy,omitempty" xml:"ElasticStrategy,omitempty"`
	EnableAIMaster           *bool   `json:"EnableAIMaster,omitempty" xml:"EnableAIMaster,omitempty"`
	EnableEDP                *bool   `json:"EnableEDP,omitempty" xml:"EnableEDP,omitempty"`
	EnableElasticTraining    *bool   `json:"EnableElasticTraining,omitempty" xml:"EnableElasticTraining,omitempty"`
	EnablePsJobElasticPS     *bool   `json:"EnablePsJobElasticPS,omitempty" xml:"EnablePsJobElasticPS,omitempty"`
	EnablePsJobElasticWorker *bool   `json:"EnablePsJobElasticWorker,omitempty" xml:"EnablePsJobElasticWorker,omitempty"`
	EnablePsResourceEstimate *bool   `json:"EnablePsResourceEstimate,omitempty" xml:"EnablePsResourceEstimate,omitempty"`
	MaxParallelism           *int32  `json:"MaxParallelism,omitempty" xml:"MaxParallelism,omitempty"`
	MinParallelism           *int32  `json:"MinParallelism,omitempty" xml:"MinParallelism,omitempty"`
	PSMaxParallelism         *int32  `json:"PSMaxParallelism,omitempty" xml:"PSMaxParallelism,omitempty"`
	PSMinParallelism         *int32  `json:"PSMinParallelism,omitempty" xml:"PSMinParallelism,omitempty"`
}

func (s JobElasticSpec) String() string {
	return tea.Prettify(s)
}

func (s JobElasticSpec) GoString() string {
	return s.String()
}

func (s *JobElasticSpec) SetAIMasterDockerImage(v string) *JobElasticSpec {
	s.AIMasterDockerImage = &v
	return s
}

func (s *JobElasticSpec) SetAIMasterType(v string) *JobElasticSpec {
	s.AIMasterType = &v
	return s
}

func (s *JobElasticSpec) SetEDPMaxParallelism(v int32) *JobElasticSpec {
	s.EDPMaxParallelism = &v
	return s
}

func (s *JobElasticSpec) SetEDPMinParallelism(v int32) *JobElasticSpec {
	s.EDPMinParallelism = &v
	return s
}

func (s *JobElasticSpec) SetElasticStrategy(v string) *JobElasticSpec {
	s.ElasticStrategy = &v
	return s
}

func (s *JobElasticSpec) SetEnableAIMaster(v bool) *JobElasticSpec {
	s.EnableAIMaster = &v
	return s
}

func (s *JobElasticSpec) SetEnableEDP(v bool) *JobElasticSpec {
	s.EnableEDP = &v
	return s
}

func (s *JobElasticSpec) SetEnableElasticTraining(v bool) *JobElasticSpec {
	s.EnableElasticTraining = &v
	return s
}

func (s *JobElasticSpec) SetEnablePsJobElasticPS(v bool) *JobElasticSpec {
	s.EnablePsJobElasticPS = &v
	return s
}

func (s *JobElasticSpec) SetEnablePsJobElasticWorker(v bool) *JobElasticSpec {
	s.EnablePsJobElasticWorker = &v
	return s
}

func (s *JobElasticSpec) SetEnablePsResourceEstimate(v bool) *JobElasticSpec {
	s.EnablePsResourceEstimate = &v
	return s
}

func (s *JobElasticSpec) SetMaxParallelism(v int32) *JobElasticSpec {
	s.MaxParallelism = &v
	return s
}

func (s *JobElasticSpec) SetMinParallelism(v int32) *JobElasticSpec {
	s.MinParallelism = &v
	return s
}

func (s *JobElasticSpec) SetPSMaxParallelism(v int32) *JobElasticSpec {
	s.PSMaxParallelism = &v
	return s
}

func (s *JobElasticSpec) SetPSMinParallelism(v int32) *JobElasticSpec {
	s.PSMinParallelism = &v
	return s
}

type JobItem struct {
	CodeSource          *JobItemCodeSource    `json:"CodeSource,omitempty" xml:"CodeSource,omitempty" type:"Struct"`
	DataSources         []*JobItemDataSources `json:"DataSources,omitempty" xml:"DataSources,omitempty" type:"Repeated"`
	DisplayName         *string               `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	Duration            *int64                `json:"Duration,omitempty" xml:"Duration,omitempty"`
	EnabledDebugger     *bool                 `json:"EnabledDebugger,omitempty" xml:"EnabledDebugger,omitempty"`
	Envs                map[string]*string    `json:"Envs,omitempty" xml:"Envs,omitempty"`
	GmtCreateTime       *string               `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	GmtFailedTime       *string               `json:"GmtFailedTime,omitempty" xml:"GmtFailedTime,omitempty"`
	GmtFinishTime       *string               `json:"GmtFinishTime,omitempty" xml:"GmtFinishTime,omitempty"`
	GmtRunningTime      *string               `json:"GmtRunningTime,omitempty" xml:"GmtRunningTime,omitempty"`
	GmtStoppedTime      *string               `json:"GmtStoppedTime,omitempty" xml:"GmtStoppedTime,omitempty"`
	GmtSubmittedTime    *string               `json:"GmtSubmittedTime,omitempty" xml:"GmtSubmittedTime,omitempty"`
	GmtSuccessedTime    *string               `json:"GmtSuccessedTime,omitempty" xml:"GmtSuccessedTime,omitempty"`
	JobId               *string               `json:"JobId,omitempty" xml:"JobId,omitempty"`
	JobSpecs            []*JobSpec            `json:"JobSpecs,omitempty" xml:"JobSpecs,omitempty" type:"Repeated"`
	JobType             *string               `json:"JobType,omitempty" xml:"JobType,omitempty"`
	Priority            *int32                `json:"Priority,omitempty" xml:"Priority,omitempty"`
	ReasonCode          *string               `json:"ReasonCode,omitempty" xml:"ReasonCode,omitempty"`
	ReasonMessage       *string               `json:"ReasonMessage,omitempty" xml:"ReasonMessage,omitempty"`
	ResourceId          *string               `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceLevel       *string               `json:"ResourceLevel,omitempty" xml:"ResourceLevel,omitempty"`
	ResourceName        *string               `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	ResourceType        *string               `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Settings            *JobSettings          `json:"Settings,omitempty" xml:"Settings,omitempty"`
	Status              *string               `json:"Status,omitempty" xml:"Status,omitempty"`
	SubStatus           *string               `json:"SubStatus,omitempty" xml:"SubStatus,omitempty"`
	ThirdpartyLibDir    *string               `json:"ThirdpartyLibDir,omitempty" xml:"ThirdpartyLibDir,omitempty"`
	ThirdpartyLibs      []*string             `json:"ThirdpartyLibs,omitempty" xml:"ThirdpartyLibs,omitempty" type:"Repeated"`
	UseOversoldResource *bool                 `json:"UseOversoldResource,omitempty" xml:"UseOversoldResource,omitempty"`
	UserCommand         *string               `json:"UserCommand,omitempty" xml:"UserCommand,omitempty"`
	UserId              *string               `json:"UserId,omitempty" xml:"UserId,omitempty"`
	Username            *string               `json:"Username,omitempty" xml:"Username,omitempty"`
	WorkspaceId         *string               `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
	WorkspaceName       *string               `json:"WorkspaceName,omitempty" xml:"WorkspaceName,omitempty"`
}

func (s JobItem) String() string {
	return tea.Prettify(s)
}

func (s JobItem) GoString() string {
	return s.String()
}

func (s *JobItem) SetCodeSource(v *JobItemCodeSource) *JobItem {
	s.CodeSource = v
	return s
}

func (s *JobItem) SetDataSources(v []*JobItemDataSources) *JobItem {
	s.DataSources = v
	return s
}

func (s *JobItem) SetDisplayName(v string) *JobItem {
	s.DisplayName = &v
	return s
}

func (s *JobItem) SetDuration(v int64) *JobItem {
	s.Duration = &v
	return s
}

func (s *JobItem) SetEnabledDebugger(v bool) *JobItem {
	s.EnabledDebugger = &v
	return s
}

func (s *JobItem) SetEnvs(v map[string]*string) *JobItem {
	s.Envs = v
	return s
}

func (s *JobItem) SetGmtCreateTime(v string) *JobItem {
	s.GmtCreateTime = &v
	return s
}

func (s *JobItem) SetGmtFailedTime(v string) *JobItem {
	s.GmtFailedTime = &v
	return s
}

func (s *JobItem) SetGmtFinishTime(v string) *JobItem {
	s.GmtFinishTime = &v
	return s
}

func (s *JobItem) SetGmtRunningTime(v string) *JobItem {
	s.GmtRunningTime = &v
	return s
}

func (s *JobItem) SetGmtStoppedTime(v string) *JobItem {
	s.GmtStoppedTime = &v
	return s
}

func (s *JobItem) SetGmtSubmittedTime(v string) *JobItem {
	s.GmtSubmittedTime = &v
	return s
}

func (s *JobItem) SetGmtSuccessedTime(v string) *JobItem {
	s.GmtSuccessedTime = &v
	return s
}

func (s *JobItem) SetJobId(v string) *JobItem {
	s.JobId = &v
	return s
}

func (s *JobItem) SetJobSpecs(v []*JobSpec) *JobItem {
	s.JobSpecs = v
	return s
}

func (s *JobItem) SetJobType(v string) *JobItem {
	s.JobType = &v
	return s
}

func (s *JobItem) SetPriority(v int32) *JobItem {
	s.Priority = &v
	return s
}

func (s *JobItem) SetReasonCode(v string) *JobItem {
	s.ReasonCode = &v
	return s
}

func (s *JobItem) SetReasonMessage(v string) *JobItem {
	s.ReasonMessage = &v
	return s
}

func (s *JobItem) SetResourceId(v string) *JobItem {
	s.ResourceId = &v
	return s
}

func (s *JobItem) SetResourceLevel(v string) *JobItem {
	s.ResourceLevel = &v
	return s
}

func (s *JobItem) SetResourceName(v string) *JobItem {
	s.ResourceName = &v
	return s
}

func (s *JobItem) SetResourceType(v string) *JobItem {
	s.ResourceType = &v
	return s
}

func (s *JobItem) SetSettings(v *JobSettings) *JobItem {
	s.Settings = v
	return s
}

func (s *JobItem) SetStatus(v string) *JobItem {
	s.Status = &v
	return s
}

func (s *JobItem) SetSubStatus(v string) *JobItem {
	s.SubStatus = &v
	return s
}

func (s *JobItem) SetThirdpartyLibDir(v string) *JobItem {
	s.ThirdpartyLibDir = &v
	return s
}

func (s *JobItem) SetThirdpartyLibs(v []*string) *JobItem {
	s.ThirdpartyLibs = v
	return s
}

func (s *JobItem) SetUseOversoldResource(v bool) *JobItem {
	s.UseOversoldResource = &v
	return s
}

func (s *JobItem) SetUserCommand(v string) *JobItem {
	s.UserCommand = &v
	return s
}

func (s *JobItem) SetUserId(v string) *JobItem {
	s.UserId = &v
	return s
}

func (s *JobItem) SetUsername(v string) *JobItem {
	s.Username = &v
	return s
}

func (s *JobItem) SetWorkspaceId(v string) *JobItem {
	s.WorkspaceId = &v
	return s
}

func (s *JobItem) SetWorkspaceName(v string) *JobItem {
	s.WorkspaceName = &v
	return s
}

type JobItemCodeSource struct {
	Branch       *string `json:"Branch,omitempty" xml:"Branch,omitempty"`
	CodeSourceId *string `json:"CodeSourceId,omitempty" xml:"CodeSourceId,omitempty"`
	Commit       *string `json:"Commit,omitempty" xml:"Commit,omitempty"`
	MountPath    *string `json:"MountPath,omitempty" xml:"MountPath,omitempty"`
}

func (s JobItemCodeSource) String() string {
	return tea.Prettify(s)
}

func (s JobItemCodeSource) GoString() string {
	return s.String()
}

func (s *JobItemCodeSource) SetBranch(v string) *JobItemCodeSource {
	s.Branch = &v
	return s
}

func (s *JobItemCodeSource) SetCodeSourceId(v string) *JobItemCodeSource {
	s.CodeSourceId = &v
	return s
}

func (s *JobItemCodeSource) SetCommit(v string) *JobItemCodeSource {
	s.Commit = &v
	return s
}

func (s *JobItemCodeSource) SetMountPath(v string) *JobItemCodeSource {
	s.MountPath = &v
	return s
}

type JobItemDataSources struct {
	DataSourceId *string `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	MountPath    *string `json:"MountPath,omitempty" xml:"MountPath,omitempty"`
}

func (s JobItemDataSources) String() string {
	return tea.Prettify(s)
}

func (s JobItemDataSources) GoString() string {
	return s.String()
}

func (s *JobItemDataSources) SetDataSourceId(v string) *JobItemDataSources {
	s.DataSourceId = &v
	return s
}

func (s *JobItemDataSources) SetMountPath(v string) *JobItemDataSources {
	s.MountPath = &v
	return s
}

type JobSettings struct {
	AdvancedSettings                map[string]interface{} `json:"AdvancedSettings,omitempty" xml:"AdvancedSettings,omitempty"`
	BusinessUserId                  *string                `json:"BusinessUserId,omitempty" xml:"BusinessUserId,omitempty"`
	Caller                          *string                `json:"Caller,omitempty" xml:"Caller,omitempty"`
	Driver                          *string                `json:"Driver,omitempty" xml:"Driver,omitempty"`
	EnableErrorMonitoringInAIMaster *bool                  `json:"EnableErrorMonitoringInAIMaster,omitempty" xml:"EnableErrorMonitoringInAIMaster,omitempty"`
	EnableOssAppend                 *bool                  `json:"EnableOssAppend,omitempty" xml:"EnableOssAppend,omitempty"`
	EnableRDMA                      *bool                  `json:"EnableRDMA,omitempty" xml:"EnableRDMA,omitempty"`
	EnableSanityCheck               *bool                  `json:"EnableSanityCheck,omitempty" xml:"EnableSanityCheck,omitempty"`
	EnableTideResource              *bool                  `json:"EnableTideResource,omitempty" xml:"EnableTideResource,omitempty"`
	ErrorMonitoringArgs             *string                `json:"ErrorMonitoringArgs,omitempty" xml:"ErrorMonitoringArgs,omitempty"`
	JobReservedMinutes              *int32                 `json:"JobReservedMinutes,omitempty" xml:"JobReservedMinutes,omitempty"`
	JobReservedPolicy               *string                `json:"JobReservedPolicy,omitempty" xml:"JobReservedPolicy,omitempty"`
	OversoldType                    *string                `json:"OversoldType,omitempty" xml:"OversoldType,omitempty"`
	PipelineId                      *string                `json:"PipelineId,omitempty" xml:"PipelineId,omitempty"`
	SanityCheckArgs                 *string                `json:"SanityCheckArgs,omitempty" xml:"SanityCheckArgs,omitempty"`
	Tags                            map[string]*string     `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s JobSettings) String() string {
	return tea.Prettify(s)
}

func (s JobSettings) GoString() string {
	return s.String()
}

func (s *JobSettings) SetAdvancedSettings(v map[string]interface{}) *JobSettings {
	s.AdvancedSettings = v
	return s
}

func (s *JobSettings) SetBusinessUserId(v string) *JobSettings {
	s.BusinessUserId = &v
	return s
}

func (s *JobSettings) SetCaller(v string) *JobSettings {
	s.Caller = &v
	return s
}

func (s *JobSettings) SetDriver(v string) *JobSettings {
	s.Driver = &v
	return s
}

func (s *JobSettings) SetEnableErrorMonitoringInAIMaster(v bool) *JobSettings {
	s.EnableErrorMonitoringInAIMaster = &v
	return s
}

func (s *JobSettings) SetEnableOssAppend(v bool) *JobSettings {
	s.EnableOssAppend = &v
	return s
}

func (s *JobSettings) SetEnableRDMA(v bool) *JobSettings {
	s.EnableRDMA = &v
	return s
}

func (s *JobSettings) SetEnableSanityCheck(v bool) *JobSettings {
	s.EnableSanityCheck = &v
	return s
}

func (s *JobSettings) SetEnableTideResource(v bool) *JobSettings {
	s.EnableTideResource = &v
	return s
}

func (s *JobSettings) SetErrorMonitoringArgs(v string) *JobSettings {
	s.ErrorMonitoringArgs = &v
	return s
}

func (s *JobSettings) SetJobReservedMinutes(v int32) *JobSettings {
	s.JobReservedMinutes = &v
	return s
}

func (s *JobSettings) SetJobReservedPolicy(v string) *JobSettings {
	s.JobReservedPolicy = &v
	return s
}

func (s *JobSettings) SetOversoldType(v string) *JobSettings {
	s.OversoldType = &v
	return s
}

func (s *JobSettings) SetPipelineId(v string) *JobSettings {
	s.PipelineId = &v
	return s
}

func (s *JobSettings) SetSanityCheckArgs(v string) *JobSettings {
	s.SanityCheckArgs = &v
	return s
}

func (s *JobSettings) SetTags(v map[string]*string) *JobSettings {
	s.Tags = v
	return s
}

type JobSpec struct {
	EcsSpec         *string         `json:"EcsSpec,omitempty" xml:"EcsSpec,omitempty"`
	ExtraPodSpec    *ExtraPodSpec   `json:"ExtraPodSpec,omitempty" xml:"ExtraPodSpec,omitempty"`
	Image           *string         `json:"Image,omitempty" xml:"Image,omitempty"`
	ImageConfig     *ImageConfig    `json:"ImageConfig,omitempty" xml:"ImageConfig,omitempty"`
	PodCount        *int64          `json:"PodCount,omitempty" xml:"PodCount,omitempty"`
	ResourceConfig  *ResourceConfig `json:"ResourceConfig,omitempty" xml:"ResourceConfig,omitempty"`
	SpotSpec        *SpotSpec       `json:"SpotSpec,omitempty" xml:"SpotSpec,omitempty"`
	Type            *string         `json:"Type,omitempty" xml:"Type,omitempty"`
	UseSpotInstance *bool           `json:"UseSpotInstance,omitempty" xml:"UseSpotInstance,omitempty"`
}

func (s JobSpec) String() string {
	return tea.Prettify(s)
}

func (s JobSpec) GoString() string {
	return s.String()
}

func (s *JobSpec) SetEcsSpec(v string) *JobSpec {
	s.EcsSpec = &v
	return s
}

func (s *JobSpec) SetExtraPodSpec(v *ExtraPodSpec) *JobSpec {
	s.ExtraPodSpec = v
	return s
}

func (s *JobSpec) SetImage(v string) *JobSpec {
	s.Image = &v
	return s
}

func (s *JobSpec) SetImageConfig(v *ImageConfig) *JobSpec {
	s.ImageConfig = v
	return s
}

func (s *JobSpec) SetPodCount(v int64) *JobSpec {
	s.PodCount = &v
	return s
}

func (s *JobSpec) SetResourceConfig(v *ResourceConfig) *JobSpec {
	s.ResourceConfig = v
	return s
}

func (s *JobSpec) SetSpotSpec(v *SpotSpec) *JobSpec {
	s.SpotSpec = v
	return s
}

func (s *JobSpec) SetType(v string) *JobSpec {
	s.Type = &v
	return s
}

func (s *JobSpec) SetUseSpotInstance(v bool) *JobSpec {
	s.UseSpotInstance = &v
	return s
}

type LogInfo struct {
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	Id      *string `json:"Id,omitempty" xml:"Id,omitempty"`
	PodId   *string `json:"PodId,omitempty" xml:"PodId,omitempty"`
	PodUid  *string `json:"PodUid,omitempty" xml:"PodUid,omitempty"`
	Source  *string `json:"Source,omitempty" xml:"Source,omitempty"`
	Time    *string `json:"Time,omitempty" xml:"Time,omitempty"`
}

func (s LogInfo) String() string {
	return tea.Prettify(s)
}

func (s LogInfo) GoString() string {
	return s.String()
}

func (s *LogInfo) SetContent(v string) *LogInfo {
	s.Content = &v
	return s
}

func (s *LogInfo) SetId(v string) *LogInfo {
	s.Id = &v
	return s
}

func (s *LogInfo) SetPodId(v string) *LogInfo {
	s.PodId = &v
	return s
}

func (s *LogInfo) SetPodUid(v string) *LogInfo {
	s.PodUid = &v
	return s
}

func (s *LogInfo) SetSource(v string) *LogInfo {
	s.Source = &v
	return s
}

func (s *LogInfo) SetTime(v string) *LogInfo {
	s.Time = &v
	return s
}

type Member struct {
	MemberId   *string `json:"MemberId,omitempty" xml:"MemberId,omitempty"`
	MemberType *string `json:"MemberType,omitempty" xml:"MemberType,omitempty"`
}

func (s Member) String() string {
	return tea.Prettify(s)
}

func (s Member) GoString() string {
	return s.String()
}

func (s *Member) SetMemberId(v string) *Member {
	s.MemberId = &v
	return s
}

func (s *Member) SetMemberType(v string) *Member {
	s.MemberType = &v
	return s
}

type Metric struct {
	Time  *int64   `json:"Time,omitempty" xml:"Time,omitempty"`
	Value *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s Metric) String() string {
	return tea.Prettify(s)
}

func (s Metric) GoString() string {
	return s.String()
}

func (s *Metric) SetTime(v int64) *Metric {
	s.Time = &v
	return s
}

func (s *Metric) SetValue(v float32) *Metric {
	s.Value = &v
	return s
}

type NodeMetric struct {
	Metrics  []*Metric `json:"Metrics,omitempty" xml:"Metrics,omitempty" type:"Repeated"`
	NodeName *string   `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
}

func (s NodeMetric) String() string {
	return tea.Prettify(s)
}

func (s NodeMetric) GoString() string {
	return s.String()
}

func (s *NodeMetric) SetMetrics(v []*Metric) *NodeMetric {
	s.Metrics = v
	return s
}

func (s *NodeMetric) SetNodeName(v string) *NodeMetric {
	s.NodeName = &v
	return s
}

type PodItem struct {
	GmtCreateTime *string    `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	GmtFinishTime *string    `json:"GmtFinishTime,omitempty" xml:"GmtFinishTime,omitempty"`
	GmtStartTime  *string    `json:"GmtStartTime,omitempty" xml:"GmtStartTime,omitempty"`
	HistoryPods   []*PodItem `json:"HistoryPods,omitempty" xml:"HistoryPods,omitempty" type:"Repeated"`
	Ip            *string    `json:"Ip,omitempty" xml:"Ip,omitempty"`
	PodId         *string    `json:"PodId,omitempty" xml:"PodId,omitempty"`
	PodUid        *string    `json:"PodUid,omitempty" xml:"PodUid,omitempty"`
	Status        *string    `json:"Status,omitempty" xml:"Status,omitempty"`
	Type          *string    `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s PodItem) String() string {
	return tea.Prettify(s)
}

func (s PodItem) GoString() string {
	return s.String()
}

func (s *PodItem) SetGmtCreateTime(v string) *PodItem {
	s.GmtCreateTime = &v
	return s
}

func (s *PodItem) SetGmtFinishTime(v string) *PodItem {
	s.GmtFinishTime = &v
	return s
}

func (s *PodItem) SetGmtStartTime(v string) *PodItem {
	s.GmtStartTime = &v
	return s
}

func (s *PodItem) SetHistoryPods(v []*PodItem) *PodItem {
	s.HistoryPods = v
	return s
}

func (s *PodItem) SetIp(v string) *PodItem {
	s.Ip = &v
	return s
}

func (s *PodItem) SetPodId(v string) *PodItem {
	s.PodId = &v
	return s
}

func (s *PodItem) SetPodUid(v string) *PodItem {
	s.PodUid = &v
	return s
}

func (s *PodItem) SetStatus(v string) *PodItem {
	s.Status = &v
	return s
}

func (s *PodItem) SetType(v string) *PodItem {
	s.Type = &v
	return s
}

type PodMetric struct {
	Metrics []*Metric `json:"Metrics,omitempty" xml:"Metrics,omitempty" type:"Repeated"`
	PodId   *string   `json:"PodId,omitempty" xml:"PodId,omitempty"`
}

func (s PodMetric) String() string {
	return tea.Prettify(s)
}

func (s PodMetric) GoString() string {
	return s.String()
}

func (s *PodMetric) SetMetrics(v []*Metric) *PodMetric {
	s.Metrics = v
	return s
}

func (s *PodMetric) SetPodId(v string) *PodMetric {
	s.PodId = &v
	return s
}

type Quota struct {
	ClusterId      *string      `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	ClusterName    *string      `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	QuotaConfig    *QuotaConfig `json:"QuotaConfig,omitempty" xml:"QuotaConfig,omitempty"`
	QuotaId        *string      `json:"QuotaId,omitempty" xml:"QuotaId,omitempty"`
	QuotaName      *string      `json:"QuotaName,omitempty" xml:"QuotaName,omitempty"`
	QuotaType      *string      `json:"QuotaType,omitempty" xml:"QuotaType,omitempty"`
	TotalQuota     *QuotaDetail `json:"TotalQuota,omitempty" xml:"TotalQuota,omitempty"`
	TotalTideQuota *QuotaDetail `json:"TotalTideQuota,omitempty" xml:"TotalTideQuota,omitempty"`
	UsedQuota      *QuotaDetail `json:"UsedQuota,omitempty" xml:"UsedQuota,omitempty"`
	UsedTideQuota  *QuotaDetail `json:"UsedTideQuota,omitempty" xml:"UsedTideQuota,omitempty"`
}

func (s Quota) String() string {
	return tea.Prettify(s)
}

func (s Quota) GoString() string {
	return s.String()
}

func (s *Quota) SetClusterId(v string) *Quota {
	s.ClusterId = &v
	return s
}

func (s *Quota) SetClusterName(v string) *Quota {
	s.ClusterName = &v
	return s
}

func (s *Quota) SetQuotaConfig(v *QuotaConfig) *Quota {
	s.QuotaConfig = v
	return s
}

func (s *Quota) SetQuotaId(v string) *Quota {
	s.QuotaId = &v
	return s
}

func (s *Quota) SetQuotaName(v string) *Quota {
	s.QuotaName = &v
	return s
}

func (s *Quota) SetQuotaType(v string) *Quota {
	s.QuotaType = &v
	return s
}

func (s *Quota) SetTotalQuota(v *QuotaDetail) *Quota {
	s.TotalQuota = v
	return s
}

func (s *Quota) SetTotalTideQuota(v *QuotaDetail) *Quota {
	s.TotalTideQuota = v
	return s
}

func (s *Quota) SetUsedQuota(v *QuotaDetail) *Quota {
	s.UsedQuota = v
	return s
}

func (s *Quota) SetUsedTideQuota(v *QuotaDetail) *Quota {
	s.UsedTideQuota = v
	return s
}

type QuotaConfig struct {
	AllowedMaxPriority *int32  `json:"AllowedMaxPriority,omitempty" xml:"AllowedMaxPriority,omitempty"`
	EnableDLC          *bool   `json:"EnableDLC,omitempty" xml:"EnableDLC,omitempty"`
	EnableDSW          *bool   `json:"EnableDSW,omitempty" xml:"EnableDSW,omitempty"`
	EnableTideResource *bool   `json:"EnableTideResource,omitempty" xml:"EnableTideResource,omitempty"`
	ResourceLevel      *string `json:"ResourceLevel,omitempty" xml:"ResourceLevel,omitempty"`
}

func (s QuotaConfig) String() string {
	return tea.Prettify(s)
}

func (s QuotaConfig) GoString() string {
	return s.String()
}

func (s *QuotaConfig) SetAllowedMaxPriority(v int32) *QuotaConfig {
	s.AllowedMaxPriority = &v
	return s
}

func (s *QuotaConfig) SetEnableDLC(v bool) *QuotaConfig {
	s.EnableDLC = &v
	return s
}

func (s *QuotaConfig) SetEnableDSW(v bool) *QuotaConfig {
	s.EnableDSW = &v
	return s
}

func (s *QuotaConfig) SetEnableTideResource(v bool) *QuotaConfig {
	s.EnableTideResource = &v
	return s
}

func (s *QuotaConfig) SetResourceLevel(v string) *QuotaConfig {
	s.ResourceLevel = &v
	return s
}

type QuotaDetail struct {
	CPU             *string      `json:"CPU,omitempty" xml:"CPU,omitempty"`
	GPU             *string      `json:"GPU,omitempty" xml:"GPU,omitempty"`
	GPUDetails      []*GPUDetail `json:"GPUDetails,omitempty" xml:"GPUDetails,omitempty" type:"Repeated"`
	GPUType         *string      `json:"GPUType,omitempty" xml:"GPUType,omitempty"`
	GPUTypeFullName *string      `json:"GPUTypeFullName,omitempty" xml:"GPUTypeFullName,omitempty"`
	Memory          *string      `json:"Memory,omitempty" xml:"Memory,omitempty"`
}

func (s QuotaDetail) String() string {
	return tea.Prettify(s)
}

func (s QuotaDetail) GoString() string {
	return s.String()
}

func (s *QuotaDetail) SetCPU(v string) *QuotaDetail {
	s.CPU = &v
	return s
}

func (s *QuotaDetail) SetGPU(v string) *QuotaDetail {
	s.GPU = &v
	return s
}

func (s *QuotaDetail) SetGPUDetails(v []*GPUDetail) *QuotaDetail {
	s.GPUDetails = v
	return s
}

func (s *QuotaDetail) SetGPUType(v string) *QuotaDetail {
	s.GPUType = &v
	return s
}

func (s *QuotaDetail) SetGPUTypeFullName(v string) *QuotaDetail {
	s.GPUTypeFullName = &v
	return s
}

func (s *QuotaDetail) SetMemory(v string) *QuotaDetail {
	s.Memory = &v
	return s
}

type ResourceConfig struct {
	CPU          *string `json:"CPU,omitempty" xml:"CPU,omitempty"`
	GPU          *string `json:"GPU,omitempty" xml:"GPU,omitempty"`
	GPUType      *string `json:"GPUType,omitempty" xml:"GPUType,omitempty"`
	Memory       *string `json:"Memory,omitempty" xml:"Memory,omitempty"`
	SharedMemory *string `json:"SharedMemory,omitempty" xml:"SharedMemory,omitempty"`
}

func (s ResourceConfig) String() string {
	return tea.Prettify(s)
}

func (s ResourceConfig) GoString() string {
	return s.String()
}

func (s *ResourceConfig) SetCPU(v string) *ResourceConfig {
	s.CPU = &v
	return s
}

func (s *ResourceConfig) SetGPU(v string) *ResourceConfig {
	s.GPU = &v
	return s
}

func (s *ResourceConfig) SetGPUType(v string) *ResourceConfig {
	s.GPUType = &v
	return s
}

func (s *ResourceConfig) SetMemory(v string) *ResourceConfig {
	s.Memory = &v
	return s
}

func (s *ResourceConfig) SetSharedMemory(v string) *ResourceConfig {
	s.SharedMemory = &v
	return s
}

type ResourceRequirements struct {
	Limits   map[string]*string `json:"Limits,omitempty" xml:"Limits,omitempty"`
	Requests map[string]*string `json:"Requests,omitempty" xml:"Requests,omitempty"`
}

func (s ResourceRequirements) String() string {
	return tea.Prettify(s)
}

func (s ResourceRequirements) GoString() string {
	return s.String()
}

func (s *ResourceRequirements) SetLimits(v map[string]*string) *ResourceRequirements {
	s.Limits = v
	return s
}

func (s *ResourceRequirements) SetRequests(v map[string]*string) *ResourceRequirements {
	s.Requests = v
	return s
}

type Resources struct {
	CPU    *string `json:"CPU,omitempty" xml:"CPU,omitempty"`
	GPU    *string `json:"GPU,omitempty" xml:"GPU,omitempty"`
	Memory *string `json:"Memory,omitempty" xml:"Memory,omitempty"`
}

func (s Resources) String() string {
	return tea.Prettify(s)
}

func (s Resources) GoString() string {
	return s.String()
}

func (s *Resources) SetCPU(v string) *Resources {
	s.CPU = &v
	return s
}

func (s *Resources) SetGPU(v string) *Resources {
	s.GPU = &v
	return s
}

func (s *Resources) SetMemory(v string) *Resources {
	s.Memory = &v
	return s
}

type SanityCheckResultItem struct {
	CheckNumber *int32  `json:"CheckNumber,omitempty" xml:"CheckNumber,omitempty"`
	FinishedAt  *string `json:"FinishedAt,omitempty" xml:"FinishedAt,omitempty"`
	Message     *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Phase       *string `json:"Phase,omitempty" xml:"Phase,omitempty"`
	StartedAt   *string `json:"StartedAt,omitempty" xml:"StartedAt,omitempty"`
	Status      *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s SanityCheckResultItem) String() string {
	return tea.Prettify(s)
}

func (s SanityCheckResultItem) GoString() string {
	return s.String()
}

func (s *SanityCheckResultItem) SetCheckNumber(v int32) *SanityCheckResultItem {
	s.CheckNumber = &v
	return s
}

func (s *SanityCheckResultItem) SetFinishedAt(v string) *SanityCheckResultItem {
	s.FinishedAt = &v
	return s
}

func (s *SanityCheckResultItem) SetMessage(v string) *SanityCheckResultItem {
	s.Message = &v
	return s
}

func (s *SanityCheckResultItem) SetPhase(v string) *SanityCheckResultItem {
	s.Phase = &v
	return s
}

func (s *SanityCheckResultItem) SetStartedAt(v string) *SanityCheckResultItem {
	s.StartedAt = &v
	return s
}

func (s *SanityCheckResultItem) SetStatus(v string) *SanityCheckResultItem {
	s.Status = &v
	return s
}

type SmartCache struct {
	CacheWorkerNum  *int64  `json:"CacheWorkerNum,omitempty" xml:"CacheWorkerNum,omitempty"`
	CacheWorkerSize *int64  `json:"CacheWorkerSize,omitempty" xml:"CacheWorkerSize,omitempty"`
	Description     *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DisplayName     *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	Duration        *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	Endpoint        *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	FileSystemId    *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	GmtCreateTime   *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	GmtModifyTime   *string `json:"GmtModifyTime,omitempty" xml:"GmtModifyTime,omitempty"`
	MountPath       *string `json:"MountPath,omitempty" xml:"MountPath,omitempty"`
	Options         *string `json:"Options,omitempty" xml:"Options,omitempty"`
	Path            *string `json:"Path,omitempty" xml:"Path,omitempty"`
	SmartCacheId    *string `json:"SmartCacheId,omitempty" xml:"SmartCacheId,omitempty"`
	Status          *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Type            *string `json:"Type,omitempty" xml:"Type,omitempty"`
	UserId          *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s SmartCache) String() string {
	return tea.Prettify(s)
}

func (s SmartCache) GoString() string {
	return s.String()
}

func (s *SmartCache) SetCacheWorkerNum(v int64) *SmartCache {
	s.CacheWorkerNum = &v
	return s
}

func (s *SmartCache) SetCacheWorkerSize(v int64) *SmartCache {
	s.CacheWorkerSize = &v
	return s
}

func (s *SmartCache) SetDescription(v string) *SmartCache {
	s.Description = &v
	return s
}

func (s *SmartCache) SetDisplayName(v string) *SmartCache {
	s.DisplayName = &v
	return s
}

func (s *SmartCache) SetDuration(v string) *SmartCache {
	s.Duration = &v
	return s
}

func (s *SmartCache) SetEndpoint(v string) *SmartCache {
	s.Endpoint = &v
	return s
}

func (s *SmartCache) SetFileSystemId(v string) *SmartCache {
	s.FileSystemId = &v
	return s
}

func (s *SmartCache) SetGmtCreateTime(v string) *SmartCache {
	s.GmtCreateTime = &v
	return s
}

func (s *SmartCache) SetGmtModifyTime(v string) *SmartCache {
	s.GmtModifyTime = &v
	return s
}

func (s *SmartCache) SetMountPath(v string) *SmartCache {
	s.MountPath = &v
	return s
}

func (s *SmartCache) SetOptions(v string) *SmartCache {
	s.Options = &v
	return s
}

func (s *SmartCache) SetPath(v string) *SmartCache {
	s.Path = &v
	return s
}

func (s *SmartCache) SetSmartCacheId(v string) *SmartCache {
	s.SmartCacheId = &v
	return s
}

func (s *SmartCache) SetStatus(v string) *SmartCache {
	s.Status = &v
	return s
}

func (s *SmartCache) SetType(v string) *SmartCache {
	s.Type = &v
	return s
}

func (s *SmartCache) SetUserId(v string) *SmartCache {
	s.UserId = &v
	return s
}

type SpotSpec struct {
	SpotDiscountLimit *float32 `json:"SpotDiscountLimit,omitempty" xml:"SpotDiscountLimit,omitempty"`
	SpotStrategy      *string  `json:"SpotStrategy,omitempty" xml:"SpotStrategy,omitempty"`
}

func (s SpotSpec) String() string {
	return tea.Prettify(s)
}

func (s SpotSpec) GoString() string {
	return s.String()
}

func (s *SpotSpec) SetSpotDiscountLimit(v float32) *SpotSpec {
	s.SpotDiscountLimit = &v
	return s
}

func (s *SpotSpec) SetSpotStrategy(v string) *SpotSpec {
	s.SpotStrategy = &v
	return s
}

type StatusTransitionItem struct {
	EndTime       *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	ReasonCode    *string `json:"ReasonCode,omitempty" xml:"ReasonCode,omitempty"`
	ReasonMessage *string `json:"ReasonMessage,omitempty" xml:"ReasonMessage,omitempty"`
	StartTime     *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Status        *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s StatusTransitionItem) String() string {
	return tea.Prettify(s)
}

func (s StatusTransitionItem) GoString() string {
	return s.String()
}

func (s *StatusTransitionItem) SetEndTime(v string) *StatusTransitionItem {
	s.EndTime = &v
	return s
}

func (s *StatusTransitionItem) SetReasonCode(v string) *StatusTransitionItem {
	s.ReasonCode = &v
	return s
}

func (s *StatusTransitionItem) SetReasonMessage(v string) *StatusTransitionItem {
	s.ReasonMessage = &v
	return s
}

func (s *StatusTransitionItem) SetStartTime(v string) *StatusTransitionItem {
	s.StartTime = &v
	return s
}

func (s *StatusTransitionItem) SetStatus(v string) *StatusTransitionItem {
	s.Status = &v
	return s
}

type Tensorboard struct {
	DataSourceId   *string `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	DisplayName    *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	Duration       *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	GmtCreateTime  *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	GmtModifyTime  *string `json:"GmtModifyTime,omitempty" xml:"GmtModifyTime,omitempty"`
	JobId          *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	ReasonCode     *string `json:"ReasonCode,omitempty" xml:"ReasonCode,omitempty"`
	ReasonMessage  *string `json:"ReasonMessage,omitempty" xml:"ReasonMessage,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Status         *string `json:"Status,omitempty" xml:"Status,omitempty"`
	SummaryPath    *string `json:"SummaryPath,omitempty" xml:"SummaryPath,omitempty"`
	TensorboardId  *string `json:"TensorboardId,omitempty" xml:"TensorboardId,omitempty"`
	TensorboardUrl *string `json:"TensorboardUrl,omitempty" xml:"TensorboardUrl,omitempty"`
	UserId         *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s Tensorboard) String() string {
	return tea.Prettify(s)
}

func (s Tensorboard) GoString() string {
	return s.String()
}

func (s *Tensorboard) SetDataSourceId(v string) *Tensorboard {
	s.DataSourceId = &v
	return s
}

func (s *Tensorboard) SetDisplayName(v string) *Tensorboard {
	s.DisplayName = &v
	return s
}

func (s *Tensorboard) SetDuration(v string) *Tensorboard {
	s.Duration = &v
	return s
}

func (s *Tensorboard) SetGmtCreateTime(v string) *Tensorboard {
	s.GmtCreateTime = &v
	return s
}

func (s *Tensorboard) SetGmtModifyTime(v string) *Tensorboard {
	s.GmtModifyTime = &v
	return s
}

func (s *Tensorboard) SetJobId(v string) *Tensorboard {
	s.JobId = &v
	return s
}

func (s *Tensorboard) SetReasonCode(v string) *Tensorboard {
	s.ReasonCode = &v
	return s
}

func (s *Tensorboard) SetReasonMessage(v string) *Tensorboard {
	s.ReasonMessage = &v
	return s
}

func (s *Tensorboard) SetRequestId(v string) *Tensorboard {
	s.RequestId = &v
	return s
}

func (s *Tensorboard) SetStatus(v string) *Tensorboard {
	s.Status = &v
	return s
}

func (s *Tensorboard) SetSummaryPath(v string) *Tensorboard {
	s.SummaryPath = &v
	return s
}

func (s *Tensorboard) SetTensorboardId(v string) *Tensorboard {
	s.TensorboardId = &v
	return s
}

func (s *Tensorboard) SetTensorboardUrl(v string) *Tensorboard {
	s.TensorboardUrl = &v
	return s
}

func (s *Tensorboard) SetUserId(v string) *Tensorboard {
	s.UserId = &v
	return s
}

type Workspace struct {
	Creator         *string    `json:"Creator,omitempty" xml:"Creator,omitempty"`
	GmtCreateTime   *string    `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	GmtModifyTime   *string    `json:"GmtModifyTime,omitempty" xml:"GmtModifyTime,omitempty"`
	Members         []*Member  `json:"Members,omitempty" xml:"Members,omitempty" type:"Repeated"`
	Quotas          []*Quota   `json:"Quotas,omitempty" xml:"Quotas,omitempty" type:"Repeated"`
	TotalResources  *Resources `json:"TotalResources,omitempty" xml:"TotalResources,omitempty"`
	WorkspaceAdmins []*Member  `json:"WorkspaceAdmins,omitempty" xml:"WorkspaceAdmins,omitempty" type:"Repeated"`
	WorkspaceId     *string    `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
	WorkspaceName   *string    `json:"WorkspaceName,omitempty" xml:"WorkspaceName,omitempty"`
}

func (s Workspace) String() string {
	return tea.Prettify(s)
}

func (s Workspace) GoString() string {
	return s.String()
}

func (s *Workspace) SetCreator(v string) *Workspace {
	s.Creator = &v
	return s
}

func (s *Workspace) SetGmtCreateTime(v string) *Workspace {
	s.GmtCreateTime = &v
	return s
}

func (s *Workspace) SetGmtModifyTime(v string) *Workspace {
	s.GmtModifyTime = &v
	return s
}

func (s *Workspace) SetMembers(v []*Member) *Workspace {
	s.Members = v
	return s
}

func (s *Workspace) SetQuotas(v []*Quota) *Workspace {
	s.Quotas = v
	return s
}

func (s *Workspace) SetTotalResources(v *Resources) *Workspace {
	s.TotalResources = v
	return s
}

func (s *Workspace) SetWorkspaceAdmins(v []*Member) *Workspace {
	s.WorkspaceAdmins = v
	return s
}

func (s *Workspace) SetWorkspaceId(v string) *Workspace {
	s.WorkspaceId = &v
	return s
}

func (s *Workspace) SetWorkspaceName(v string) *Workspace {
	s.WorkspaceName = &v
	return s
}

type CreateJobRequest struct {
	CodeSource               *CreateJobRequestCodeSource    `json:"CodeSource,omitempty" xml:"CodeSource,omitempty" type:"Struct"`
	DataSources              []*CreateJobRequestDataSources `json:"DataSources,omitempty" xml:"DataSources,omitempty" type:"Repeated"`
	DebuggerConfigContent    *string                        `json:"DebuggerConfigContent,omitempty" xml:"DebuggerConfigContent,omitempty"`
	DisplayName              *string                        `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	ElasticSpec              *JobElasticSpec                `json:"ElasticSpec,omitempty" xml:"ElasticSpec,omitempty"`
	Envs                     map[string]*string             `json:"Envs,omitempty" xml:"Envs,omitempty"`
	JobMaxRunningTimeMinutes *int64                         `json:"JobMaxRunningTimeMinutes,omitempty" xml:"JobMaxRunningTimeMinutes,omitempty"`
	JobSpecs                 []*JobSpec                     `json:"JobSpecs,omitempty" xml:"JobSpecs,omitempty" type:"Repeated"`
	JobType                  *string                        `json:"JobType,omitempty" xml:"JobType,omitempty"`
	Options                  *string                        `json:"Options,omitempty" xml:"Options,omitempty"`
	Priority                 *int32                         `json:"Priority,omitempty" xml:"Priority,omitempty"`
	ResourceId               *string                        `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	Settings                 *JobSettings                   `json:"Settings,omitempty" xml:"Settings,omitempty"`
	SuccessPolicy            *string                        `json:"SuccessPolicy,omitempty" xml:"SuccessPolicy,omitempty"`
	ThirdpartyLibDir         *string                        `json:"ThirdpartyLibDir,omitempty" xml:"ThirdpartyLibDir,omitempty"`
	ThirdpartyLibs           []*string                      `json:"ThirdpartyLibs,omitempty" xml:"ThirdpartyLibs,omitempty" type:"Repeated"`
	UserCommand              *string                        `json:"UserCommand,omitempty" xml:"UserCommand,omitempty"`
	UserVpc                  *CreateJobRequestUserVpc       `json:"UserVpc,omitempty" xml:"UserVpc,omitempty" type:"Struct"`
	WorkspaceId              *string                        `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s CreateJobRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateJobRequest) GoString() string {
	return s.String()
}

func (s *CreateJobRequest) SetCodeSource(v *CreateJobRequestCodeSource) *CreateJobRequest {
	s.CodeSource = v
	return s
}

func (s *CreateJobRequest) SetDataSources(v []*CreateJobRequestDataSources) *CreateJobRequest {
	s.DataSources = v
	return s
}

func (s *CreateJobRequest) SetDebuggerConfigContent(v string) *CreateJobRequest {
	s.DebuggerConfigContent = &v
	return s
}

func (s *CreateJobRequest) SetDisplayName(v string) *CreateJobRequest {
	s.DisplayName = &v
	return s
}

func (s *CreateJobRequest) SetElasticSpec(v *JobElasticSpec) *CreateJobRequest {
	s.ElasticSpec = v
	return s
}

func (s *CreateJobRequest) SetEnvs(v map[string]*string) *CreateJobRequest {
	s.Envs = v
	return s
}

func (s *CreateJobRequest) SetJobMaxRunningTimeMinutes(v int64) *CreateJobRequest {
	s.JobMaxRunningTimeMinutes = &v
	return s
}

func (s *CreateJobRequest) SetJobSpecs(v []*JobSpec) *CreateJobRequest {
	s.JobSpecs = v
	return s
}

func (s *CreateJobRequest) SetJobType(v string) *CreateJobRequest {
	s.JobType = &v
	return s
}

func (s *CreateJobRequest) SetOptions(v string) *CreateJobRequest {
	s.Options = &v
	return s
}

func (s *CreateJobRequest) SetPriority(v int32) *CreateJobRequest {
	s.Priority = &v
	return s
}

func (s *CreateJobRequest) SetResourceId(v string) *CreateJobRequest {
	s.ResourceId = &v
	return s
}

func (s *CreateJobRequest) SetSettings(v *JobSettings) *CreateJobRequest {
	s.Settings = v
	return s
}

func (s *CreateJobRequest) SetSuccessPolicy(v string) *CreateJobRequest {
	s.SuccessPolicy = &v
	return s
}

func (s *CreateJobRequest) SetThirdpartyLibDir(v string) *CreateJobRequest {
	s.ThirdpartyLibDir = &v
	return s
}

func (s *CreateJobRequest) SetThirdpartyLibs(v []*string) *CreateJobRequest {
	s.ThirdpartyLibs = v
	return s
}

func (s *CreateJobRequest) SetUserCommand(v string) *CreateJobRequest {
	s.UserCommand = &v
	return s
}

func (s *CreateJobRequest) SetUserVpc(v *CreateJobRequestUserVpc) *CreateJobRequest {
	s.UserVpc = v
	return s
}

func (s *CreateJobRequest) SetWorkspaceId(v string) *CreateJobRequest {
	s.WorkspaceId = &v
	return s
}

type CreateJobRequestCodeSource struct {
	Branch       *string `json:"Branch,omitempty" xml:"Branch,omitempty"`
	CodeSourceId *string `json:"CodeSourceId,omitempty" xml:"CodeSourceId,omitempty"`
	Commit       *string `json:"Commit,omitempty" xml:"Commit,omitempty"`
	MountPath    *string `json:"MountPath,omitempty" xml:"MountPath,omitempty"`
}

func (s CreateJobRequestCodeSource) String() string {
	return tea.Prettify(s)
}

func (s CreateJobRequestCodeSource) GoString() string {
	return s.String()
}

func (s *CreateJobRequestCodeSource) SetBranch(v string) *CreateJobRequestCodeSource {
	s.Branch = &v
	return s
}

func (s *CreateJobRequestCodeSource) SetCodeSourceId(v string) *CreateJobRequestCodeSource {
	s.CodeSourceId = &v
	return s
}

func (s *CreateJobRequestCodeSource) SetCommit(v string) *CreateJobRequestCodeSource {
	s.Commit = &v
	return s
}

func (s *CreateJobRequestCodeSource) SetMountPath(v string) *CreateJobRequestCodeSource {
	s.MountPath = &v
	return s
}

type CreateJobRequestDataSources struct {
	DataSourceId *string `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	MountPath    *string `json:"MountPath,omitempty" xml:"MountPath,omitempty"`
	Uri          *string `json:"Uri,omitempty" xml:"Uri,omitempty"`
}

func (s CreateJobRequestDataSources) String() string {
	return tea.Prettify(s)
}

func (s CreateJobRequestDataSources) GoString() string {
	return s.String()
}

func (s *CreateJobRequestDataSources) SetDataSourceId(v string) *CreateJobRequestDataSources {
	s.DataSourceId = &v
	return s
}

func (s *CreateJobRequestDataSources) SetMountPath(v string) *CreateJobRequestDataSources {
	s.MountPath = &v
	return s
}

func (s *CreateJobRequestDataSources) SetUri(v string) *CreateJobRequestDataSources {
	s.Uri = &v
	return s
}

type CreateJobRequestUserVpc struct {
	DefaultRoute    *string   `json:"DefaultRoute,omitempty" xml:"DefaultRoute,omitempty"`
	ExtendedCIDRs   []*string `json:"ExtendedCIDRs,omitempty" xml:"ExtendedCIDRs,omitempty" type:"Repeated"`
	SecurityGroupId *string   `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	SwitchId        *string   `json:"SwitchId,omitempty" xml:"SwitchId,omitempty"`
	VpcId           *string   `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s CreateJobRequestUserVpc) String() string {
	return tea.Prettify(s)
}

func (s CreateJobRequestUserVpc) GoString() string {
	return s.String()
}

func (s *CreateJobRequestUserVpc) SetDefaultRoute(v string) *CreateJobRequestUserVpc {
	s.DefaultRoute = &v
	return s
}

func (s *CreateJobRequestUserVpc) SetExtendedCIDRs(v []*string) *CreateJobRequestUserVpc {
	s.ExtendedCIDRs = v
	return s
}

func (s *CreateJobRequestUserVpc) SetSecurityGroupId(v string) *CreateJobRequestUserVpc {
	s.SecurityGroupId = &v
	return s
}

func (s *CreateJobRequestUserVpc) SetSwitchId(v string) *CreateJobRequestUserVpc {
	s.SwitchId = &v
	return s
}

func (s *CreateJobRequestUserVpc) SetVpcId(v string) *CreateJobRequestUserVpc {
	s.VpcId = &v
	return s
}

type CreateJobResponseBody struct {
	JobId     *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateJobResponseBody) GoString() string {
	return s.String()
}

func (s *CreateJobResponseBody) SetJobId(v string) *CreateJobResponseBody {
	s.JobId = &v
	return s
}

func (s *CreateJobResponseBody) SetRequestId(v string) *CreateJobResponseBody {
	s.RequestId = &v
	return s
}

type CreateJobResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateJobResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateJobResponse) GoString() string {
	return s.String()
}

func (s *CreateJobResponse) SetHeaders(v map[string]*string) *CreateJobResponse {
	s.Headers = v
	return s
}

func (s *CreateJobResponse) SetStatusCode(v int32) *CreateJobResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateJobResponse) SetBody(v *CreateJobResponseBody) *CreateJobResponse {
	s.Body = v
	return s
}

type CreateTensorboardRequest struct {
	Cpu                   *int64            `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	DataSourceId          *string           `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	DataSourceType        *string           `json:"DataSourceType,omitempty" xml:"DataSourceType,omitempty"`
	DataSources           []*DataSourceItem `json:"DataSources,omitempty" xml:"DataSources,omitempty" type:"Repeated"`
	DisplayName           *string           `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	JobId                 *string           `json:"JobId,omitempty" xml:"JobId,omitempty"`
	MaxRunningTimeMinutes *int64            `json:"MaxRunningTimeMinutes,omitempty" xml:"MaxRunningTimeMinutes,omitempty"`
	Memory                *int64            `json:"Memory,omitempty" xml:"Memory,omitempty"`
	Options               *string           `json:"Options,omitempty" xml:"Options,omitempty"`
	SourceId              *string           `json:"SourceId,omitempty" xml:"SourceId,omitempty"`
	SourceType            *string           `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	SummaryPath           *string           `json:"SummaryPath,omitempty" xml:"SummaryPath,omitempty"`
	SummaryRelativePath   *string           `json:"SummaryRelativePath,omitempty" xml:"SummaryRelativePath,omitempty"`
	Uri                   *string           `json:"Uri,omitempty" xml:"Uri,omitempty"`
	WorkspaceId           *string           `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s CreateTensorboardRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateTensorboardRequest) GoString() string {
	return s.String()
}

func (s *CreateTensorboardRequest) SetCpu(v int64) *CreateTensorboardRequest {
	s.Cpu = &v
	return s
}

func (s *CreateTensorboardRequest) SetDataSourceId(v string) *CreateTensorboardRequest {
	s.DataSourceId = &v
	return s
}

func (s *CreateTensorboardRequest) SetDataSourceType(v string) *CreateTensorboardRequest {
	s.DataSourceType = &v
	return s
}

func (s *CreateTensorboardRequest) SetDataSources(v []*DataSourceItem) *CreateTensorboardRequest {
	s.DataSources = v
	return s
}

func (s *CreateTensorboardRequest) SetDisplayName(v string) *CreateTensorboardRequest {
	s.DisplayName = &v
	return s
}

func (s *CreateTensorboardRequest) SetJobId(v string) *CreateTensorboardRequest {
	s.JobId = &v
	return s
}

func (s *CreateTensorboardRequest) SetMaxRunningTimeMinutes(v int64) *CreateTensorboardRequest {
	s.MaxRunningTimeMinutes = &v
	return s
}

func (s *CreateTensorboardRequest) SetMemory(v int64) *CreateTensorboardRequest {
	s.Memory = &v
	return s
}

func (s *CreateTensorboardRequest) SetOptions(v string) *CreateTensorboardRequest {
	s.Options = &v
	return s
}

func (s *CreateTensorboardRequest) SetSourceId(v string) *CreateTensorboardRequest {
	s.SourceId = &v
	return s
}

func (s *CreateTensorboardRequest) SetSourceType(v string) *CreateTensorboardRequest {
	s.SourceType = &v
	return s
}

func (s *CreateTensorboardRequest) SetSummaryPath(v string) *CreateTensorboardRequest {
	s.SummaryPath = &v
	return s
}

func (s *CreateTensorboardRequest) SetSummaryRelativePath(v string) *CreateTensorboardRequest {
	s.SummaryRelativePath = &v
	return s
}

func (s *CreateTensorboardRequest) SetUri(v string) *CreateTensorboardRequest {
	s.Uri = &v
	return s
}

func (s *CreateTensorboardRequest) SetWorkspaceId(v string) *CreateTensorboardRequest {
	s.WorkspaceId = &v
	return s
}

type CreateTensorboardResponseBody struct {
	DataSourceId  *string `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	JobId         *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TensorboardId *string `json:"TensorboardId,omitempty" xml:"TensorboardId,omitempty"`
}

func (s CreateTensorboardResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateTensorboardResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTensorboardResponseBody) SetDataSourceId(v string) *CreateTensorboardResponseBody {
	s.DataSourceId = &v
	return s
}

func (s *CreateTensorboardResponseBody) SetJobId(v string) *CreateTensorboardResponseBody {
	s.JobId = &v
	return s
}

func (s *CreateTensorboardResponseBody) SetRequestId(v string) *CreateTensorboardResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTensorboardResponseBody) SetTensorboardId(v string) *CreateTensorboardResponseBody {
	s.TensorboardId = &v
	return s
}

type CreateTensorboardResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateTensorboardResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateTensorboardResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateTensorboardResponse) GoString() string {
	return s.String()
}

func (s *CreateTensorboardResponse) SetHeaders(v map[string]*string) *CreateTensorboardResponse {
	s.Headers = v
	return s
}

func (s *CreateTensorboardResponse) SetStatusCode(v int32) *CreateTensorboardResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateTensorboardResponse) SetBody(v *CreateTensorboardResponseBody) *CreateTensorboardResponse {
	s.Body = v
	return s
}

type DeleteJobResponseBody struct {
	JobId     *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteJobResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteJobResponseBody) SetJobId(v string) *DeleteJobResponseBody {
	s.JobId = &v
	return s
}

func (s *DeleteJobResponseBody) SetRequestId(v string) *DeleteJobResponseBody {
	s.RequestId = &v
	return s
}

type DeleteJobResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteJobResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteJobResponse) GoString() string {
	return s.String()
}

func (s *DeleteJobResponse) SetHeaders(v map[string]*string) *DeleteJobResponse {
	s.Headers = v
	return s
}

func (s *DeleteJobResponse) SetStatusCode(v int32) *DeleteJobResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteJobResponse) SetBody(v *DeleteJobResponseBody) *DeleteJobResponse {
	s.Body = v
	return s
}

type DeleteTensorboardRequest struct {
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s DeleteTensorboardRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteTensorboardRequest) GoString() string {
	return s.String()
}

func (s *DeleteTensorboardRequest) SetWorkspaceId(v string) *DeleteTensorboardRequest {
	s.WorkspaceId = &v
	return s
}

type DeleteTensorboardResponseBody struct {
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TensorboardId *string `json:"TensorboardId,omitempty" xml:"TensorboardId,omitempty"`
}

func (s DeleteTensorboardResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteTensorboardResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteTensorboardResponseBody) SetRequestId(v string) *DeleteTensorboardResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteTensorboardResponseBody) SetTensorboardId(v string) *DeleteTensorboardResponseBody {
	s.TensorboardId = &v
	return s
}

type DeleteTensorboardResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteTensorboardResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteTensorboardResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteTensorboardResponse) GoString() string {
	return s.String()
}

func (s *DeleteTensorboardResponse) SetHeaders(v map[string]*string) *DeleteTensorboardResponse {
	s.Headers = v
	return s
}

func (s *DeleteTensorboardResponse) SetStatusCode(v int32) *DeleteTensorboardResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteTensorboardResponse) SetBody(v *DeleteTensorboardResponseBody) *DeleteTensorboardResponse {
	s.Body = v
	return s
}

type GetJobRequest struct {
	NeedDetail *bool `json:"NeedDetail,omitempty" xml:"NeedDetail,omitempty"`
}

func (s GetJobRequest) String() string {
	return tea.Prettify(s)
}

func (s GetJobRequest) GoString() string {
	return s.String()
}

func (s *GetJobRequest) SetNeedDetail(v bool) *GetJobRequest {
	s.NeedDetail = &v
	return s
}

type GetJobResponseBody struct {
	ClusterId        *string                          `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	CodeSource       *GetJobResponseBodyCodeSource    `json:"CodeSource,omitempty" xml:"CodeSource,omitempty" type:"Struct"`
	DataSources      []*GetJobResponseBodyDataSources `json:"DataSources,omitempty" xml:"DataSources,omitempty" type:"Repeated"`
	DisplayName      *string                          `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	Duration         *int64                           `json:"Duration,omitempty" xml:"Duration,omitempty"`
	ElasticSpec      *JobElasticSpec                  `json:"ElasticSpec,omitempty" xml:"ElasticSpec,omitempty"`
	EnabledDebugger  *bool                            `json:"EnabledDebugger,omitempty" xml:"EnabledDebugger,omitempty"`
	Envs             map[string]*string               `json:"Envs,omitempty" xml:"Envs,omitempty"`
	GmtCreateTime    *string                          `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	GmtFailedTime    *string                          `json:"GmtFailedTime,omitempty" xml:"GmtFailedTime,omitempty"`
	GmtFinishTime    *string                          `json:"GmtFinishTime,omitempty" xml:"GmtFinishTime,omitempty"`
	GmtRunningTime   *string                          `json:"GmtRunningTime,omitempty" xml:"GmtRunningTime,omitempty"`
	GmtStoppedTime   *string                          `json:"GmtStoppedTime,omitempty" xml:"GmtStoppedTime,omitempty"`
	GmtSubmittedTime *string                          `json:"GmtSubmittedTime,omitempty" xml:"GmtSubmittedTime,omitempty"`
	GmtSuccessedTime *string                          `json:"GmtSuccessedTime,omitempty" xml:"GmtSuccessedTime,omitempty"`
	JobId            *string                          `json:"JobId,omitempty" xml:"JobId,omitempty"`
	JobSpecs         []*JobSpec                       `json:"JobSpecs,omitempty" xml:"JobSpecs,omitempty" type:"Repeated"`
	JobType          *string                          `json:"JobType,omitempty" xml:"JobType,omitempty"`
	Pods             []*GetJobResponseBodyPods        `json:"Pods,omitempty" xml:"Pods,omitempty" type:"Repeated"`
	Priority         *int32                           `json:"Priority,omitempty" xml:"Priority,omitempty"`
	ReasonCode       *string                          `json:"ReasonCode,omitempty" xml:"ReasonCode,omitempty"`
	ReasonMessage    *string                          `json:"ReasonMessage,omitempty" xml:"ReasonMessage,omitempty"`
	RequestId        *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResourceId       *string                          `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceLevel    *string                          `json:"ResourceLevel,omitempty" xml:"ResourceLevel,omitempty"`
	ResourceType     *string                          `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	RestartTimes     *string                          `json:"RestartTimes,omitempty" xml:"RestartTimes,omitempty"`
	Settings         *JobSettings                     `json:"Settings,omitempty" xml:"Settings,omitempty"`
	Status           *string                          `json:"Status,omitempty" xml:"Status,omitempty"`
	StatusHistory    []*StatusTransitionItem          `json:"StatusHistory,omitempty" xml:"StatusHistory,omitempty" type:"Repeated"`
	SubStatus        *string                          `json:"SubStatus,omitempty" xml:"SubStatus,omitempty"`
	TenantId         *string                          `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	ThirdpartyLibDir *string                          `json:"ThirdpartyLibDir,omitempty" xml:"ThirdpartyLibDir,omitempty"`
	ThirdpartyLibs   []*string                        `json:"ThirdpartyLibs,omitempty" xml:"ThirdpartyLibs,omitempty" type:"Repeated"`
	UserCommand      *string                          `json:"UserCommand,omitempty" xml:"UserCommand,omitempty"`
	UserId           *string                          `json:"UserId,omitempty" xml:"UserId,omitempty"`
	WorkspaceId      *string                          `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
	WorkspaceName    *string                          `json:"WorkspaceName,omitempty" xml:"WorkspaceName,omitempty"`
}

func (s GetJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetJobResponseBody) GoString() string {
	return s.String()
}

func (s *GetJobResponseBody) SetClusterId(v string) *GetJobResponseBody {
	s.ClusterId = &v
	return s
}

func (s *GetJobResponseBody) SetCodeSource(v *GetJobResponseBodyCodeSource) *GetJobResponseBody {
	s.CodeSource = v
	return s
}

func (s *GetJobResponseBody) SetDataSources(v []*GetJobResponseBodyDataSources) *GetJobResponseBody {
	s.DataSources = v
	return s
}

func (s *GetJobResponseBody) SetDisplayName(v string) *GetJobResponseBody {
	s.DisplayName = &v
	return s
}

func (s *GetJobResponseBody) SetDuration(v int64) *GetJobResponseBody {
	s.Duration = &v
	return s
}

func (s *GetJobResponseBody) SetElasticSpec(v *JobElasticSpec) *GetJobResponseBody {
	s.ElasticSpec = v
	return s
}

func (s *GetJobResponseBody) SetEnabledDebugger(v bool) *GetJobResponseBody {
	s.EnabledDebugger = &v
	return s
}

func (s *GetJobResponseBody) SetEnvs(v map[string]*string) *GetJobResponseBody {
	s.Envs = v
	return s
}

func (s *GetJobResponseBody) SetGmtCreateTime(v string) *GetJobResponseBody {
	s.GmtCreateTime = &v
	return s
}

func (s *GetJobResponseBody) SetGmtFailedTime(v string) *GetJobResponseBody {
	s.GmtFailedTime = &v
	return s
}

func (s *GetJobResponseBody) SetGmtFinishTime(v string) *GetJobResponseBody {
	s.GmtFinishTime = &v
	return s
}

func (s *GetJobResponseBody) SetGmtRunningTime(v string) *GetJobResponseBody {
	s.GmtRunningTime = &v
	return s
}

func (s *GetJobResponseBody) SetGmtStoppedTime(v string) *GetJobResponseBody {
	s.GmtStoppedTime = &v
	return s
}

func (s *GetJobResponseBody) SetGmtSubmittedTime(v string) *GetJobResponseBody {
	s.GmtSubmittedTime = &v
	return s
}

func (s *GetJobResponseBody) SetGmtSuccessedTime(v string) *GetJobResponseBody {
	s.GmtSuccessedTime = &v
	return s
}

func (s *GetJobResponseBody) SetJobId(v string) *GetJobResponseBody {
	s.JobId = &v
	return s
}

func (s *GetJobResponseBody) SetJobSpecs(v []*JobSpec) *GetJobResponseBody {
	s.JobSpecs = v
	return s
}

func (s *GetJobResponseBody) SetJobType(v string) *GetJobResponseBody {
	s.JobType = &v
	return s
}

func (s *GetJobResponseBody) SetPods(v []*GetJobResponseBodyPods) *GetJobResponseBody {
	s.Pods = v
	return s
}

func (s *GetJobResponseBody) SetPriority(v int32) *GetJobResponseBody {
	s.Priority = &v
	return s
}

func (s *GetJobResponseBody) SetReasonCode(v string) *GetJobResponseBody {
	s.ReasonCode = &v
	return s
}

func (s *GetJobResponseBody) SetReasonMessage(v string) *GetJobResponseBody {
	s.ReasonMessage = &v
	return s
}

func (s *GetJobResponseBody) SetRequestId(v string) *GetJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetJobResponseBody) SetResourceId(v string) *GetJobResponseBody {
	s.ResourceId = &v
	return s
}

func (s *GetJobResponseBody) SetResourceLevel(v string) *GetJobResponseBody {
	s.ResourceLevel = &v
	return s
}

func (s *GetJobResponseBody) SetResourceType(v string) *GetJobResponseBody {
	s.ResourceType = &v
	return s
}

func (s *GetJobResponseBody) SetRestartTimes(v string) *GetJobResponseBody {
	s.RestartTimes = &v
	return s
}

func (s *GetJobResponseBody) SetSettings(v *JobSettings) *GetJobResponseBody {
	s.Settings = v
	return s
}

func (s *GetJobResponseBody) SetStatus(v string) *GetJobResponseBody {
	s.Status = &v
	return s
}

func (s *GetJobResponseBody) SetStatusHistory(v []*StatusTransitionItem) *GetJobResponseBody {
	s.StatusHistory = v
	return s
}

func (s *GetJobResponseBody) SetSubStatus(v string) *GetJobResponseBody {
	s.SubStatus = &v
	return s
}

func (s *GetJobResponseBody) SetTenantId(v string) *GetJobResponseBody {
	s.TenantId = &v
	return s
}

func (s *GetJobResponseBody) SetThirdpartyLibDir(v string) *GetJobResponseBody {
	s.ThirdpartyLibDir = &v
	return s
}

func (s *GetJobResponseBody) SetThirdpartyLibs(v []*string) *GetJobResponseBody {
	s.ThirdpartyLibs = v
	return s
}

func (s *GetJobResponseBody) SetUserCommand(v string) *GetJobResponseBody {
	s.UserCommand = &v
	return s
}

func (s *GetJobResponseBody) SetUserId(v string) *GetJobResponseBody {
	s.UserId = &v
	return s
}

func (s *GetJobResponseBody) SetWorkspaceId(v string) *GetJobResponseBody {
	s.WorkspaceId = &v
	return s
}

func (s *GetJobResponseBody) SetWorkspaceName(v string) *GetJobResponseBody {
	s.WorkspaceName = &v
	return s
}

type GetJobResponseBodyCodeSource struct {
	Branch       *string `json:"Branch,omitempty" xml:"Branch,omitempty"`
	CodeSourceId *string `json:"CodeSourceId,omitempty" xml:"CodeSourceId,omitempty"`
	Commit       *string `json:"Commit,omitempty" xml:"Commit,omitempty"`
	MountPath    *string `json:"MountPath,omitempty" xml:"MountPath,omitempty"`
}

func (s GetJobResponseBodyCodeSource) String() string {
	return tea.Prettify(s)
}

func (s GetJobResponseBodyCodeSource) GoString() string {
	return s.String()
}

func (s *GetJobResponseBodyCodeSource) SetBranch(v string) *GetJobResponseBodyCodeSource {
	s.Branch = &v
	return s
}

func (s *GetJobResponseBodyCodeSource) SetCodeSourceId(v string) *GetJobResponseBodyCodeSource {
	s.CodeSourceId = &v
	return s
}

func (s *GetJobResponseBodyCodeSource) SetCommit(v string) *GetJobResponseBodyCodeSource {
	s.Commit = &v
	return s
}

func (s *GetJobResponseBodyCodeSource) SetMountPath(v string) *GetJobResponseBodyCodeSource {
	s.MountPath = &v
	return s
}

type GetJobResponseBodyDataSources struct {
	DataSourceId *string `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	MountPath    *string `json:"MountPath,omitempty" xml:"MountPath,omitempty"`
	Uri          *string `json:"Uri,omitempty" xml:"Uri,omitempty"`
}

func (s GetJobResponseBodyDataSources) String() string {
	return tea.Prettify(s)
}

func (s GetJobResponseBodyDataSources) GoString() string {
	return s.String()
}

func (s *GetJobResponseBodyDataSources) SetDataSourceId(v string) *GetJobResponseBodyDataSources {
	s.DataSourceId = &v
	return s
}

func (s *GetJobResponseBodyDataSources) SetMountPath(v string) *GetJobResponseBodyDataSources {
	s.MountPath = &v
	return s
}

func (s *GetJobResponseBodyDataSources) SetUri(v string) *GetJobResponseBodyDataSources {
	s.Uri = &v
	return s
}

type GetJobResponseBodyPods struct {
	GmtCreateTime *string                              `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	GmtFinishTime *string                              `json:"GmtFinishTime,omitempty" xml:"GmtFinishTime,omitempty"`
	GmtStartTime  *string                              `json:"GmtStartTime,omitempty" xml:"GmtStartTime,omitempty"`
	HistoryPods   []*GetJobResponseBodyPodsHistoryPods `json:"HistoryPods,omitempty" xml:"HistoryPods,omitempty" type:"Repeated"`
	Ip            *string                              `json:"Ip,omitempty" xml:"Ip,omitempty"`
	PodId         *string                              `json:"PodId,omitempty" xml:"PodId,omitempty"`
	PodUid        *string                              `json:"PodUid,omitempty" xml:"PodUid,omitempty"`
	ResourceType  *string                              `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Status        *string                              `json:"Status,omitempty" xml:"Status,omitempty"`
	SubStatus     *string                              `json:"SubStatus,omitempty" xml:"SubStatus,omitempty"`
	Type          *string                              `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetJobResponseBodyPods) String() string {
	return tea.Prettify(s)
}

func (s GetJobResponseBodyPods) GoString() string {
	return s.String()
}

func (s *GetJobResponseBodyPods) SetGmtCreateTime(v string) *GetJobResponseBodyPods {
	s.GmtCreateTime = &v
	return s
}

func (s *GetJobResponseBodyPods) SetGmtFinishTime(v string) *GetJobResponseBodyPods {
	s.GmtFinishTime = &v
	return s
}

func (s *GetJobResponseBodyPods) SetGmtStartTime(v string) *GetJobResponseBodyPods {
	s.GmtStartTime = &v
	return s
}

func (s *GetJobResponseBodyPods) SetHistoryPods(v []*GetJobResponseBodyPodsHistoryPods) *GetJobResponseBodyPods {
	s.HistoryPods = v
	return s
}

func (s *GetJobResponseBodyPods) SetIp(v string) *GetJobResponseBodyPods {
	s.Ip = &v
	return s
}

func (s *GetJobResponseBodyPods) SetPodId(v string) *GetJobResponseBodyPods {
	s.PodId = &v
	return s
}

func (s *GetJobResponseBodyPods) SetPodUid(v string) *GetJobResponseBodyPods {
	s.PodUid = &v
	return s
}

func (s *GetJobResponseBodyPods) SetResourceType(v string) *GetJobResponseBodyPods {
	s.ResourceType = &v
	return s
}

func (s *GetJobResponseBodyPods) SetStatus(v string) *GetJobResponseBodyPods {
	s.Status = &v
	return s
}

func (s *GetJobResponseBodyPods) SetSubStatus(v string) *GetJobResponseBodyPods {
	s.SubStatus = &v
	return s
}

func (s *GetJobResponseBodyPods) SetType(v string) *GetJobResponseBodyPods {
	s.Type = &v
	return s
}

type GetJobResponseBodyPodsHistoryPods struct {
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	GmtFinishTime *string `json:"GmtFinishTime,omitempty" xml:"GmtFinishTime,omitempty"`
	GmtStartTime  *string `json:"GmtStartTime,omitempty" xml:"GmtStartTime,omitempty"`
	Ip            *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	PodId         *string `json:"PodId,omitempty" xml:"PodId,omitempty"`
	PodUid        *string `json:"PodUid,omitempty" xml:"PodUid,omitempty"`
	ResourceType  *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Status        *string `json:"Status,omitempty" xml:"Status,omitempty"`
	SubStatus     *string `json:"SubStatus,omitempty" xml:"SubStatus,omitempty"`
	Type          *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetJobResponseBodyPodsHistoryPods) String() string {
	return tea.Prettify(s)
}

func (s GetJobResponseBodyPodsHistoryPods) GoString() string {
	return s.String()
}

func (s *GetJobResponseBodyPodsHistoryPods) SetGmtCreateTime(v string) *GetJobResponseBodyPodsHistoryPods {
	s.GmtCreateTime = &v
	return s
}

func (s *GetJobResponseBodyPodsHistoryPods) SetGmtFinishTime(v string) *GetJobResponseBodyPodsHistoryPods {
	s.GmtFinishTime = &v
	return s
}

func (s *GetJobResponseBodyPodsHistoryPods) SetGmtStartTime(v string) *GetJobResponseBodyPodsHistoryPods {
	s.GmtStartTime = &v
	return s
}

func (s *GetJobResponseBodyPodsHistoryPods) SetIp(v string) *GetJobResponseBodyPodsHistoryPods {
	s.Ip = &v
	return s
}

func (s *GetJobResponseBodyPodsHistoryPods) SetPodId(v string) *GetJobResponseBodyPodsHistoryPods {
	s.PodId = &v
	return s
}

func (s *GetJobResponseBodyPodsHistoryPods) SetPodUid(v string) *GetJobResponseBodyPodsHistoryPods {
	s.PodUid = &v
	return s
}

func (s *GetJobResponseBodyPodsHistoryPods) SetResourceType(v string) *GetJobResponseBodyPodsHistoryPods {
	s.ResourceType = &v
	return s
}

func (s *GetJobResponseBodyPodsHistoryPods) SetStatus(v string) *GetJobResponseBodyPodsHistoryPods {
	s.Status = &v
	return s
}

func (s *GetJobResponseBodyPodsHistoryPods) SetSubStatus(v string) *GetJobResponseBodyPodsHistoryPods {
	s.SubStatus = &v
	return s
}

func (s *GetJobResponseBodyPodsHistoryPods) SetType(v string) *GetJobResponseBodyPodsHistoryPods {
	s.Type = &v
	return s
}

type GetJobResponse struct {
	Headers    map[string]*string  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetJobResponse) String() string {
	return tea.Prettify(s)
}

func (s GetJobResponse) GoString() string {
	return s.String()
}

func (s *GetJobResponse) SetHeaders(v map[string]*string) *GetJobResponse {
	s.Headers = v
	return s
}

func (s *GetJobResponse) SetStatusCode(v int32) *GetJobResponse {
	s.StatusCode = &v
	return s
}

func (s *GetJobResponse) SetBody(v *GetJobResponseBody) *GetJobResponse {
	s.Body = v
	return s
}

type GetJobEventsRequest struct {
	EndTime      *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	MaxEventsNum *int32  `json:"MaxEventsNum,omitempty" xml:"MaxEventsNum,omitempty"`
	StartTime    *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s GetJobEventsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetJobEventsRequest) GoString() string {
	return s.String()
}

func (s *GetJobEventsRequest) SetEndTime(v string) *GetJobEventsRequest {
	s.EndTime = &v
	return s
}

func (s *GetJobEventsRequest) SetMaxEventsNum(v int32) *GetJobEventsRequest {
	s.MaxEventsNum = &v
	return s
}

func (s *GetJobEventsRequest) SetStartTime(v string) *GetJobEventsRequest {
	s.StartTime = &v
	return s
}

type GetJobEventsResponseBody struct {
	Events    []*string `json:"Events,omitempty" xml:"Events,omitempty" type:"Repeated"`
	JobId     *string   `json:"JobId,omitempty" xml:"JobId,omitempty"`
	RequestId *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetJobEventsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetJobEventsResponseBody) GoString() string {
	return s.String()
}

func (s *GetJobEventsResponseBody) SetEvents(v []*string) *GetJobEventsResponseBody {
	s.Events = v
	return s
}

func (s *GetJobEventsResponseBody) SetJobId(v string) *GetJobEventsResponseBody {
	s.JobId = &v
	return s
}

func (s *GetJobEventsResponseBody) SetRequestId(v string) *GetJobEventsResponseBody {
	s.RequestId = &v
	return s
}

type GetJobEventsResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetJobEventsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetJobEventsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetJobEventsResponse) GoString() string {
	return s.String()
}

func (s *GetJobEventsResponse) SetHeaders(v map[string]*string) *GetJobEventsResponse {
	s.Headers = v
	return s
}

func (s *GetJobEventsResponse) SetStatusCode(v int32) *GetJobEventsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetJobEventsResponse) SetBody(v *GetJobEventsResponseBody) *GetJobEventsResponse {
	s.Body = v
	return s
}

type GetJobMetricsRequest struct {
	EndTime    *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	MetricType *string `json:"MetricType,omitempty" xml:"MetricType,omitempty"`
	StartTime  *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	TimeStep   *string `json:"TimeStep,omitempty" xml:"TimeStep,omitempty"`
	Token      *string `json:"Token,omitempty" xml:"Token,omitempty"`
}

func (s GetJobMetricsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetJobMetricsRequest) GoString() string {
	return s.String()
}

func (s *GetJobMetricsRequest) SetEndTime(v string) *GetJobMetricsRequest {
	s.EndTime = &v
	return s
}

func (s *GetJobMetricsRequest) SetMetricType(v string) *GetJobMetricsRequest {
	s.MetricType = &v
	return s
}

func (s *GetJobMetricsRequest) SetStartTime(v string) *GetJobMetricsRequest {
	s.StartTime = &v
	return s
}

func (s *GetJobMetricsRequest) SetTimeStep(v string) *GetJobMetricsRequest {
	s.TimeStep = &v
	return s
}

func (s *GetJobMetricsRequest) SetToken(v string) *GetJobMetricsRequest {
	s.Token = &v
	return s
}

type GetJobMetricsResponseBody struct {
	JobId      *string      `json:"JobId,omitempty" xml:"JobId,omitempty"`
	PodMetrics []*PodMetric `json:"PodMetrics,omitempty" xml:"PodMetrics,omitempty" type:"Repeated"`
	RequestId  *string      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetJobMetricsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetJobMetricsResponseBody) GoString() string {
	return s.String()
}

func (s *GetJobMetricsResponseBody) SetJobId(v string) *GetJobMetricsResponseBody {
	s.JobId = &v
	return s
}

func (s *GetJobMetricsResponseBody) SetPodMetrics(v []*PodMetric) *GetJobMetricsResponseBody {
	s.PodMetrics = v
	return s
}

func (s *GetJobMetricsResponseBody) SetRequestId(v string) *GetJobMetricsResponseBody {
	s.RequestId = &v
	return s
}

type GetJobMetricsResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetJobMetricsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetJobMetricsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetJobMetricsResponse) GoString() string {
	return s.String()
}

func (s *GetJobMetricsResponse) SetHeaders(v map[string]*string) *GetJobMetricsResponse {
	s.Headers = v
	return s
}

func (s *GetJobMetricsResponse) SetStatusCode(v int32) *GetJobMetricsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetJobMetricsResponse) SetBody(v *GetJobMetricsResponseBody) *GetJobMetricsResponse {
	s.Body = v
	return s
}

type GetJobSanityCheckResultRequest struct {
	SanityCheckNumber *int32  `json:"SanityCheckNumber,omitempty" xml:"SanityCheckNumber,omitempty"`
	SanityCheckPhase  *string `json:"SanityCheckPhase,omitempty" xml:"SanityCheckPhase,omitempty"`
	Token             *string `json:"Token,omitempty" xml:"Token,omitempty"`
}

func (s GetJobSanityCheckResultRequest) String() string {
	return tea.Prettify(s)
}

func (s GetJobSanityCheckResultRequest) GoString() string {
	return s.String()
}

func (s *GetJobSanityCheckResultRequest) SetSanityCheckNumber(v int32) *GetJobSanityCheckResultRequest {
	s.SanityCheckNumber = &v
	return s
}

func (s *GetJobSanityCheckResultRequest) SetSanityCheckPhase(v string) *GetJobSanityCheckResultRequest {
	s.SanityCheckPhase = &v
	return s
}

func (s *GetJobSanityCheckResultRequest) SetToken(v string) *GetJobSanityCheckResultRequest {
	s.Token = &v
	return s
}

type GetJobSanityCheckResultResponseBody struct {
	JobId             *string                  `json:"JobId,omitempty" xml:"JobId,omitempty"`
	RequestID         *string                  `json:"RequestID,omitempty" xml:"RequestID,omitempty"`
	SanityCheckResult []*SanityCheckResultItem `json:"SanityCheckResult,omitempty" xml:"SanityCheckResult,omitempty" type:"Repeated"`
}

func (s GetJobSanityCheckResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetJobSanityCheckResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetJobSanityCheckResultResponseBody) SetJobId(v string) *GetJobSanityCheckResultResponseBody {
	s.JobId = &v
	return s
}

func (s *GetJobSanityCheckResultResponseBody) SetRequestID(v string) *GetJobSanityCheckResultResponseBody {
	s.RequestID = &v
	return s
}

func (s *GetJobSanityCheckResultResponseBody) SetSanityCheckResult(v []*SanityCheckResultItem) *GetJobSanityCheckResultResponseBody {
	s.SanityCheckResult = v
	return s
}

type GetJobSanityCheckResultResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetJobSanityCheckResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetJobSanityCheckResultResponse) String() string {
	return tea.Prettify(s)
}

func (s GetJobSanityCheckResultResponse) GoString() string {
	return s.String()
}

func (s *GetJobSanityCheckResultResponse) SetHeaders(v map[string]*string) *GetJobSanityCheckResultResponse {
	s.Headers = v
	return s
}

func (s *GetJobSanityCheckResultResponse) SetStatusCode(v int32) *GetJobSanityCheckResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GetJobSanityCheckResultResponse) SetBody(v *GetJobSanityCheckResultResponseBody) *GetJobSanityCheckResultResponse {
	s.Body = v
	return s
}

type GetPodEventsRequest struct {
	EndTime      *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	MaxEventsNum *int32  `json:"MaxEventsNum,omitempty" xml:"MaxEventsNum,omitempty"`
	PodUid       *string `json:"PodUid,omitempty" xml:"PodUid,omitempty"`
	StartTime    *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s GetPodEventsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetPodEventsRequest) GoString() string {
	return s.String()
}

func (s *GetPodEventsRequest) SetEndTime(v string) *GetPodEventsRequest {
	s.EndTime = &v
	return s
}

func (s *GetPodEventsRequest) SetMaxEventsNum(v int32) *GetPodEventsRequest {
	s.MaxEventsNum = &v
	return s
}

func (s *GetPodEventsRequest) SetPodUid(v string) *GetPodEventsRequest {
	s.PodUid = &v
	return s
}

func (s *GetPodEventsRequest) SetStartTime(v string) *GetPodEventsRequest {
	s.StartTime = &v
	return s
}

type GetPodEventsResponseBody struct {
	Events    []*string `json:"Events,omitempty" xml:"Events,omitempty" type:"Repeated"`
	JobId     *string   `json:"JobId,omitempty" xml:"JobId,omitempty"`
	PodId     *string   `json:"PodId,omitempty" xml:"PodId,omitempty"`
	PodUid    *string   `json:"PodUid,omitempty" xml:"PodUid,omitempty"`
	RequestId *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetPodEventsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetPodEventsResponseBody) GoString() string {
	return s.String()
}

func (s *GetPodEventsResponseBody) SetEvents(v []*string) *GetPodEventsResponseBody {
	s.Events = v
	return s
}

func (s *GetPodEventsResponseBody) SetJobId(v string) *GetPodEventsResponseBody {
	s.JobId = &v
	return s
}

func (s *GetPodEventsResponseBody) SetPodId(v string) *GetPodEventsResponseBody {
	s.PodId = &v
	return s
}

func (s *GetPodEventsResponseBody) SetPodUid(v string) *GetPodEventsResponseBody {
	s.PodUid = &v
	return s
}

func (s *GetPodEventsResponseBody) SetRequestId(v string) *GetPodEventsResponseBody {
	s.RequestId = &v
	return s
}

type GetPodEventsResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPodEventsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPodEventsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetPodEventsResponse) GoString() string {
	return s.String()
}

func (s *GetPodEventsResponse) SetHeaders(v map[string]*string) *GetPodEventsResponse {
	s.Headers = v
	return s
}

func (s *GetPodEventsResponse) SetStatusCode(v int32) *GetPodEventsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPodEventsResponse) SetBody(v *GetPodEventsResponseBody) *GetPodEventsResponse {
	s.Body = v
	return s
}

type GetPodLogsRequest struct {
	DownloadToFile *bool   `json:"DownloadToFile,omitempty" xml:"DownloadToFile,omitempty"`
	EndTime        *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	MaxLines       *int32  `json:"MaxLines,omitempty" xml:"MaxLines,omitempty"`
	PodUid         *string `json:"PodUid,omitempty" xml:"PodUid,omitempty"`
	StartTime      *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s GetPodLogsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetPodLogsRequest) GoString() string {
	return s.String()
}

func (s *GetPodLogsRequest) SetDownloadToFile(v bool) *GetPodLogsRequest {
	s.DownloadToFile = &v
	return s
}

func (s *GetPodLogsRequest) SetEndTime(v string) *GetPodLogsRequest {
	s.EndTime = &v
	return s
}

func (s *GetPodLogsRequest) SetMaxLines(v int32) *GetPodLogsRequest {
	s.MaxLines = &v
	return s
}

func (s *GetPodLogsRequest) SetPodUid(v string) *GetPodLogsRequest {
	s.PodUid = &v
	return s
}

func (s *GetPodLogsRequest) SetStartTime(v string) *GetPodLogsRequest {
	s.StartTime = &v
	return s
}

type GetPodLogsResponseBody struct {
	JobId     *string   `json:"JobId,omitempty" xml:"JobId,omitempty"`
	Logs      []*string `json:"Logs,omitempty" xml:"Logs,omitempty" type:"Repeated"`
	PodId     *string   `json:"PodId,omitempty" xml:"PodId,omitempty"`
	PodUid    *string   `json:"PodUid,omitempty" xml:"PodUid,omitempty"`
	RequestId *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetPodLogsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetPodLogsResponseBody) GoString() string {
	return s.String()
}

func (s *GetPodLogsResponseBody) SetJobId(v string) *GetPodLogsResponseBody {
	s.JobId = &v
	return s
}

func (s *GetPodLogsResponseBody) SetLogs(v []*string) *GetPodLogsResponseBody {
	s.Logs = v
	return s
}

func (s *GetPodLogsResponseBody) SetPodId(v string) *GetPodLogsResponseBody {
	s.PodId = &v
	return s
}

func (s *GetPodLogsResponseBody) SetPodUid(v string) *GetPodLogsResponseBody {
	s.PodUid = &v
	return s
}

func (s *GetPodLogsResponseBody) SetRequestId(v string) *GetPodLogsResponseBody {
	s.RequestId = &v
	return s
}

type GetPodLogsResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPodLogsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPodLogsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetPodLogsResponse) GoString() string {
	return s.String()
}

func (s *GetPodLogsResponse) SetHeaders(v map[string]*string) *GetPodLogsResponse {
	s.Headers = v
	return s
}

func (s *GetPodLogsResponse) SetStatusCode(v int32) *GetPodLogsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPodLogsResponse) SetBody(v *GetPodLogsResponseBody) *GetPodLogsResponse {
	s.Body = v
	return s
}

type GetTensorboardRequest struct {
	JodId       *string `json:"JodId,omitempty" xml:"JodId,omitempty"`
	Token       *string `json:"Token,omitempty" xml:"Token,omitempty"`
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s GetTensorboardRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTensorboardRequest) GoString() string {
	return s.String()
}

func (s *GetTensorboardRequest) SetJodId(v string) *GetTensorboardRequest {
	s.JodId = &v
	return s
}

func (s *GetTensorboardRequest) SetToken(v string) *GetTensorboardRequest {
	s.Token = &v
	return s
}

func (s *GetTensorboardRequest) SetWorkspaceId(v string) *GetTensorboardRequest {
	s.WorkspaceId = &v
	return s
}

type GetTensorboardResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *Tensorboard       `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTensorboardResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTensorboardResponse) GoString() string {
	return s.String()
}

func (s *GetTensorboardResponse) SetHeaders(v map[string]*string) *GetTensorboardResponse {
	s.Headers = v
	return s
}

func (s *GetTensorboardResponse) SetStatusCode(v int32) *GetTensorboardResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTensorboardResponse) SetBody(v *Tensorboard) *GetTensorboardResponse {
	s.Body = v
	return s
}

type GetTensorboardSharedUrlRequest struct {
	ExpireTimeSeconds *string `json:"ExpireTimeSeconds,omitempty" xml:"ExpireTimeSeconds,omitempty"`
}

func (s GetTensorboardSharedUrlRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTensorboardSharedUrlRequest) GoString() string {
	return s.String()
}

func (s *GetTensorboardSharedUrlRequest) SetExpireTimeSeconds(v string) *GetTensorboardSharedUrlRequest {
	s.ExpireTimeSeconds = &v
	return s
}

type GetTensorboardSharedUrlResponseBody struct {
	RequestId            *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TensorboardSharedUrl *string `json:"TensorboardSharedUrl,omitempty" xml:"TensorboardSharedUrl,omitempty"`
}

func (s GetTensorboardSharedUrlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTensorboardSharedUrlResponseBody) GoString() string {
	return s.String()
}

func (s *GetTensorboardSharedUrlResponseBody) SetRequestId(v string) *GetTensorboardSharedUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTensorboardSharedUrlResponseBody) SetTensorboardSharedUrl(v string) *GetTensorboardSharedUrlResponseBody {
	s.TensorboardSharedUrl = &v
	return s
}

type GetTensorboardSharedUrlResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTensorboardSharedUrlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTensorboardSharedUrlResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTensorboardSharedUrlResponse) GoString() string {
	return s.String()
}

func (s *GetTensorboardSharedUrlResponse) SetHeaders(v map[string]*string) *GetTensorboardSharedUrlResponse {
	s.Headers = v
	return s
}

func (s *GetTensorboardSharedUrlResponse) SetStatusCode(v int32) *GetTensorboardSharedUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTensorboardSharedUrlResponse) SetBody(v *GetTensorboardSharedUrlResponseBody) *GetTensorboardSharedUrlResponse {
	s.Body = v
	return s
}

type GetTokenRequest struct {
	ExpireTime *int64  `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	TargetId   *string `json:"TargetId,omitempty" xml:"TargetId,omitempty"`
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
}

func (s GetTokenRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTokenRequest) GoString() string {
	return s.String()
}

func (s *GetTokenRequest) SetExpireTime(v int64) *GetTokenRequest {
	s.ExpireTime = &v
	return s
}

func (s *GetTokenRequest) SetTargetId(v string) *GetTokenRequest {
	s.TargetId = &v
	return s
}

func (s *GetTokenRequest) SetTargetType(v string) *GetTokenRequest {
	s.TargetType = &v
	return s
}

type GetTokenResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Token     *string `json:"Token,omitempty" xml:"Token,omitempty"`
}

func (s GetTokenResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTokenResponseBody) GoString() string {
	return s.String()
}

func (s *GetTokenResponseBody) SetRequestId(v string) *GetTokenResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTokenResponseBody) SetToken(v string) *GetTokenResponseBody {
	s.Token = &v
	return s
}

type GetTokenResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTokenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTokenResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTokenResponse) GoString() string {
	return s.String()
}

func (s *GetTokenResponse) SetHeaders(v map[string]*string) *GetTokenResponse {
	s.Headers = v
	return s
}

func (s *GetTokenResponse) SetStatusCode(v int32) *GetTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTokenResponse) SetBody(v *GetTokenResponseBody) *GetTokenResponse {
	s.Body = v
	return s
}

type GetWebTerminalRequest struct {
	IsShared *bool `json:"IsShared,omitempty" xml:"IsShared,omitempty"`
	// Pod UID。
	PodUid *string `json:"PodUid,omitempty" xml:"PodUid,omitempty"`
}

func (s GetWebTerminalRequest) String() string {
	return tea.Prettify(s)
}

func (s GetWebTerminalRequest) GoString() string {
	return s.String()
}

func (s *GetWebTerminalRequest) SetIsShared(v bool) *GetWebTerminalRequest {
	s.IsShared = &v
	return s
}

func (s *GetWebTerminalRequest) SetPodUid(v string) *GetWebTerminalRequest {
	s.PodUid = &v
	return s
}

type GetWebTerminalResponseBody struct {
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	WebTerminalUrl *string `json:"WebTerminalUrl,omitempty" xml:"WebTerminalUrl,omitempty"`
}

func (s GetWebTerminalResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetWebTerminalResponseBody) GoString() string {
	return s.String()
}

func (s *GetWebTerminalResponseBody) SetRequestId(v string) *GetWebTerminalResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetWebTerminalResponseBody) SetWebTerminalUrl(v string) *GetWebTerminalResponseBody {
	s.WebTerminalUrl = &v
	return s
}

type GetWebTerminalResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetWebTerminalResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetWebTerminalResponse) String() string {
	return tea.Prettify(s)
}

func (s GetWebTerminalResponse) GoString() string {
	return s.String()
}

func (s *GetWebTerminalResponse) SetHeaders(v map[string]*string) *GetWebTerminalResponse {
	s.Headers = v
	return s
}

func (s *GetWebTerminalResponse) SetStatusCode(v int32) *GetWebTerminalResponse {
	s.StatusCode = &v
	return s
}

func (s *GetWebTerminalResponse) SetBody(v *GetWebTerminalResponseBody) *GetWebTerminalResponse {
	s.Body = v
	return s
}

type ListEcsSpecsRequest struct {
	AcceleratorType *string `json:"AcceleratorType,omitempty" xml:"AcceleratorType,omitempty"`
	InstanceTypes   *string `json:"InstanceTypes,omitempty" xml:"InstanceTypes,omitempty"`
	Order           *string `json:"Order,omitempty" xml:"Order,omitempty"`
	PageNumber      *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize        *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceType    *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	SortBy          *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
}

func (s ListEcsSpecsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListEcsSpecsRequest) GoString() string {
	return s.String()
}

func (s *ListEcsSpecsRequest) SetAcceleratorType(v string) *ListEcsSpecsRequest {
	s.AcceleratorType = &v
	return s
}

func (s *ListEcsSpecsRequest) SetInstanceTypes(v string) *ListEcsSpecsRequest {
	s.InstanceTypes = &v
	return s
}

func (s *ListEcsSpecsRequest) SetOrder(v string) *ListEcsSpecsRequest {
	s.Order = &v
	return s
}

func (s *ListEcsSpecsRequest) SetPageNumber(v int32) *ListEcsSpecsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListEcsSpecsRequest) SetPageSize(v int32) *ListEcsSpecsRequest {
	s.PageSize = &v
	return s
}

func (s *ListEcsSpecsRequest) SetResourceType(v string) *ListEcsSpecsRequest {
	s.ResourceType = &v
	return s
}

func (s *ListEcsSpecsRequest) SetSortBy(v string) *ListEcsSpecsRequest {
	s.SortBy = &v
	return s
}

type ListEcsSpecsResponseBody struct {
	EcsSpecs   []*EcsSpec `json:"EcsSpecs,omitempty" xml:"EcsSpecs,omitempty" type:"Repeated"`
	RequestId  *string    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int64     `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListEcsSpecsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListEcsSpecsResponseBody) GoString() string {
	return s.String()
}

func (s *ListEcsSpecsResponseBody) SetEcsSpecs(v []*EcsSpec) *ListEcsSpecsResponseBody {
	s.EcsSpecs = v
	return s
}

func (s *ListEcsSpecsResponseBody) SetRequestId(v string) *ListEcsSpecsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListEcsSpecsResponseBody) SetTotalCount(v int64) *ListEcsSpecsResponseBody {
	s.TotalCount = &v
	return s
}

type ListEcsSpecsResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListEcsSpecsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListEcsSpecsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListEcsSpecsResponse) GoString() string {
	return s.String()
}

func (s *ListEcsSpecsResponse) SetHeaders(v map[string]*string) *ListEcsSpecsResponse {
	s.Headers = v
	return s
}

func (s *ListEcsSpecsResponse) SetStatusCode(v int32) *ListEcsSpecsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListEcsSpecsResponse) SetBody(v *ListEcsSpecsResponseBody) *ListEcsSpecsResponse {
	s.Body = v
	return s
}

type ListJobSanityCheckResultsRequest struct {
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
}

func (s ListJobSanityCheckResultsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListJobSanityCheckResultsRequest) GoString() string {
	return s.String()
}

func (s *ListJobSanityCheckResultsRequest) SetOrder(v string) *ListJobSanityCheckResultsRequest {
	s.Order = &v
	return s
}

type ListJobSanityCheckResultsResponseBody struct {
	RequestID          *string                    `json:"RequestID,omitempty" xml:"RequestID,omitempty"`
	SanityCheckResults [][]*SanityCheckResultItem `json:"SanityCheckResults,omitempty" xml:"SanityCheckResults,omitempty" type:"Repeated"`
	TotalCount         *int32                     `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListJobSanityCheckResultsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListJobSanityCheckResultsResponseBody) GoString() string {
	return s.String()
}

func (s *ListJobSanityCheckResultsResponseBody) SetRequestID(v string) *ListJobSanityCheckResultsResponseBody {
	s.RequestID = &v
	return s
}

func (s *ListJobSanityCheckResultsResponseBody) SetSanityCheckResults(v [][]*SanityCheckResultItem) *ListJobSanityCheckResultsResponseBody {
	s.SanityCheckResults = v
	return s
}

func (s *ListJobSanityCheckResultsResponseBody) SetTotalCount(v int32) *ListJobSanityCheckResultsResponseBody {
	s.TotalCount = &v
	return s
}

type ListJobSanityCheckResultsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListJobSanityCheckResultsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListJobSanityCheckResultsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListJobSanityCheckResultsResponse) GoString() string {
	return s.String()
}

func (s *ListJobSanityCheckResultsResponse) SetHeaders(v map[string]*string) *ListJobSanityCheckResultsResponse {
	s.Headers = v
	return s
}

func (s *ListJobSanityCheckResultsResponse) SetStatusCode(v int32) *ListJobSanityCheckResultsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListJobSanityCheckResultsResponse) SetBody(v *ListJobSanityCheckResultsResponseBody) *ListJobSanityCheckResultsResponse {
	s.Body = v
	return s
}

type ListJobsRequest struct {
	BusinessUserId    *string            `json:"BusinessUserId,omitempty" xml:"BusinessUserId,omitempty"`
	Caller            *string            `json:"Caller,omitempty" xml:"Caller,omitempty"`
	DisplayName       *string            `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	EndTime           *string            `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	FromAllWorkspaces *bool              `json:"FromAllWorkspaces,omitempty" xml:"FromAllWorkspaces,omitempty"`
	JobId             *string            `json:"JobId,omitempty" xml:"JobId,omitempty"`
	JobType           *string            `json:"JobType,omitempty" xml:"JobType,omitempty"`
	Order             *string            `json:"Order,omitempty" xml:"Order,omitempty"`
	PageNumber        *int32             `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize          *int32             `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PipelineId        *string            `json:"PipelineId,omitempty" xml:"PipelineId,omitempty"`
	ResourceId        *string            `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ShowOwn           *bool              `json:"ShowOwn,omitempty" xml:"ShowOwn,omitempty"`
	SortBy            *string            `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	StartTime         *string            `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Status            *string            `json:"Status,omitempty" xml:"Status,omitempty"`
	Tags              map[string]*string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	UserIdForFilter   *string            `json:"UserIdForFilter,omitempty" xml:"UserIdForFilter,omitempty"`
	Username          *string            `json:"Username,omitempty" xml:"Username,omitempty"`
	WorkspaceId       *string            `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListJobsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListJobsRequest) GoString() string {
	return s.String()
}

func (s *ListJobsRequest) SetBusinessUserId(v string) *ListJobsRequest {
	s.BusinessUserId = &v
	return s
}

func (s *ListJobsRequest) SetCaller(v string) *ListJobsRequest {
	s.Caller = &v
	return s
}

func (s *ListJobsRequest) SetDisplayName(v string) *ListJobsRequest {
	s.DisplayName = &v
	return s
}

func (s *ListJobsRequest) SetEndTime(v string) *ListJobsRequest {
	s.EndTime = &v
	return s
}

func (s *ListJobsRequest) SetFromAllWorkspaces(v bool) *ListJobsRequest {
	s.FromAllWorkspaces = &v
	return s
}

func (s *ListJobsRequest) SetJobId(v string) *ListJobsRequest {
	s.JobId = &v
	return s
}

func (s *ListJobsRequest) SetJobType(v string) *ListJobsRequest {
	s.JobType = &v
	return s
}

func (s *ListJobsRequest) SetOrder(v string) *ListJobsRequest {
	s.Order = &v
	return s
}

func (s *ListJobsRequest) SetPageNumber(v int32) *ListJobsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListJobsRequest) SetPageSize(v int32) *ListJobsRequest {
	s.PageSize = &v
	return s
}

func (s *ListJobsRequest) SetPipelineId(v string) *ListJobsRequest {
	s.PipelineId = &v
	return s
}

func (s *ListJobsRequest) SetResourceId(v string) *ListJobsRequest {
	s.ResourceId = &v
	return s
}

func (s *ListJobsRequest) SetShowOwn(v bool) *ListJobsRequest {
	s.ShowOwn = &v
	return s
}

func (s *ListJobsRequest) SetSortBy(v string) *ListJobsRequest {
	s.SortBy = &v
	return s
}

func (s *ListJobsRequest) SetStartTime(v string) *ListJobsRequest {
	s.StartTime = &v
	return s
}

func (s *ListJobsRequest) SetStatus(v string) *ListJobsRequest {
	s.Status = &v
	return s
}

func (s *ListJobsRequest) SetTags(v map[string]*string) *ListJobsRequest {
	s.Tags = v
	return s
}

func (s *ListJobsRequest) SetUserIdForFilter(v string) *ListJobsRequest {
	s.UserIdForFilter = &v
	return s
}

func (s *ListJobsRequest) SetUsername(v string) *ListJobsRequest {
	s.Username = &v
	return s
}

func (s *ListJobsRequest) SetWorkspaceId(v string) *ListJobsRequest {
	s.WorkspaceId = &v
	return s
}

type ListJobsShrinkRequest struct {
	BusinessUserId    *string `json:"BusinessUserId,omitempty" xml:"BusinessUserId,omitempty"`
	Caller            *string `json:"Caller,omitempty" xml:"Caller,omitempty"`
	DisplayName       *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	EndTime           *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	FromAllWorkspaces *bool   `json:"FromAllWorkspaces,omitempty" xml:"FromAllWorkspaces,omitempty"`
	JobId             *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	JobType           *string `json:"JobType,omitempty" xml:"JobType,omitempty"`
	Order             *string `json:"Order,omitempty" xml:"Order,omitempty"`
	PageNumber        *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize          *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PipelineId        *string `json:"PipelineId,omitempty" xml:"PipelineId,omitempty"`
	ResourceId        *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ShowOwn           *bool   `json:"ShowOwn,omitempty" xml:"ShowOwn,omitempty"`
	SortBy            *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	StartTime         *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Status            *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TagsShrink        *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	UserIdForFilter   *string `json:"UserIdForFilter,omitempty" xml:"UserIdForFilter,omitempty"`
	Username          *string `json:"Username,omitempty" xml:"Username,omitempty"`
	WorkspaceId       *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListJobsShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListJobsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListJobsShrinkRequest) SetBusinessUserId(v string) *ListJobsShrinkRequest {
	s.BusinessUserId = &v
	return s
}

func (s *ListJobsShrinkRequest) SetCaller(v string) *ListJobsShrinkRequest {
	s.Caller = &v
	return s
}

func (s *ListJobsShrinkRequest) SetDisplayName(v string) *ListJobsShrinkRequest {
	s.DisplayName = &v
	return s
}

func (s *ListJobsShrinkRequest) SetEndTime(v string) *ListJobsShrinkRequest {
	s.EndTime = &v
	return s
}

func (s *ListJobsShrinkRequest) SetFromAllWorkspaces(v bool) *ListJobsShrinkRequest {
	s.FromAllWorkspaces = &v
	return s
}

func (s *ListJobsShrinkRequest) SetJobId(v string) *ListJobsShrinkRequest {
	s.JobId = &v
	return s
}

func (s *ListJobsShrinkRequest) SetJobType(v string) *ListJobsShrinkRequest {
	s.JobType = &v
	return s
}

func (s *ListJobsShrinkRequest) SetOrder(v string) *ListJobsShrinkRequest {
	s.Order = &v
	return s
}

func (s *ListJobsShrinkRequest) SetPageNumber(v int32) *ListJobsShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListJobsShrinkRequest) SetPageSize(v int32) *ListJobsShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListJobsShrinkRequest) SetPipelineId(v string) *ListJobsShrinkRequest {
	s.PipelineId = &v
	return s
}

func (s *ListJobsShrinkRequest) SetResourceId(v string) *ListJobsShrinkRequest {
	s.ResourceId = &v
	return s
}

func (s *ListJobsShrinkRequest) SetShowOwn(v bool) *ListJobsShrinkRequest {
	s.ShowOwn = &v
	return s
}

func (s *ListJobsShrinkRequest) SetSortBy(v string) *ListJobsShrinkRequest {
	s.SortBy = &v
	return s
}

func (s *ListJobsShrinkRequest) SetStartTime(v string) *ListJobsShrinkRequest {
	s.StartTime = &v
	return s
}

func (s *ListJobsShrinkRequest) SetStatus(v string) *ListJobsShrinkRequest {
	s.Status = &v
	return s
}

func (s *ListJobsShrinkRequest) SetTagsShrink(v string) *ListJobsShrinkRequest {
	s.TagsShrink = &v
	return s
}

func (s *ListJobsShrinkRequest) SetUserIdForFilter(v string) *ListJobsShrinkRequest {
	s.UserIdForFilter = &v
	return s
}

func (s *ListJobsShrinkRequest) SetUsername(v string) *ListJobsShrinkRequest {
	s.Username = &v
	return s
}

func (s *ListJobsShrinkRequest) SetWorkspaceId(v string) *ListJobsShrinkRequest {
	s.WorkspaceId = &v
	return s
}

type ListJobsResponseBody struct {
	Jobs       []*JobItem `json:"Jobs,omitempty" xml:"Jobs,omitempty" type:"Repeated"`
	RequestId  *string    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int64     `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListJobsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListJobsResponseBody) GoString() string {
	return s.String()
}

func (s *ListJobsResponseBody) SetJobs(v []*JobItem) *ListJobsResponseBody {
	s.Jobs = v
	return s
}

func (s *ListJobsResponseBody) SetRequestId(v string) *ListJobsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListJobsResponseBody) SetTotalCount(v int64) *ListJobsResponseBody {
	s.TotalCount = &v
	return s
}

type ListJobsResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListJobsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListJobsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListJobsResponse) GoString() string {
	return s.String()
}

func (s *ListJobsResponse) SetHeaders(v map[string]*string) *ListJobsResponse {
	s.Headers = v
	return s
}

func (s *ListJobsResponse) SetStatusCode(v int32) *ListJobsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListJobsResponse) SetBody(v *ListJobsResponseBody) *ListJobsResponse {
	s.Body = v
	return s
}

type ListTensorboardsRequest struct {
	DisplayName   *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	EndTime       *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	JobId         *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	Order         *string `json:"Order,omitempty" xml:"Order,omitempty"`
	PageNumber    *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize      *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ShowOwn       *bool   `json:"ShowOwn,omitempty" xml:"ShowOwn,omitempty"`
	SortBy        *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	SourceId      *string `json:"SourceId,omitempty" xml:"SourceId,omitempty"`
	SourceType    *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	StartTime     *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Status        *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TensorboardId *string `json:"TensorboardId,omitempty" xml:"TensorboardId,omitempty"`
	Verbose       *bool   `json:"Verbose,omitempty" xml:"Verbose,omitempty"`
	WorkspaceId   *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListTensorboardsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTensorboardsRequest) GoString() string {
	return s.String()
}

func (s *ListTensorboardsRequest) SetDisplayName(v string) *ListTensorboardsRequest {
	s.DisplayName = &v
	return s
}

func (s *ListTensorboardsRequest) SetEndTime(v string) *ListTensorboardsRequest {
	s.EndTime = &v
	return s
}

func (s *ListTensorboardsRequest) SetJobId(v string) *ListTensorboardsRequest {
	s.JobId = &v
	return s
}

func (s *ListTensorboardsRequest) SetOrder(v string) *ListTensorboardsRequest {
	s.Order = &v
	return s
}

func (s *ListTensorboardsRequest) SetPageNumber(v int32) *ListTensorboardsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListTensorboardsRequest) SetPageSize(v int32) *ListTensorboardsRequest {
	s.PageSize = &v
	return s
}

func (s *ListTensorboardsRequest) SetShowOwn(v bool) *ListTensorboardsRequest {
	s.ShowOwn = &v
	return s
}

func (s *ListTensorboardsRequest) SetSortBy(v string) *ListTensorboardsRequest {
	s.SortBy = &v
	return s
}

func (s *ListTensorboardsRequest) SetSourceId(v string) *ListTensorboardsRequest {
	s.SourceId = &v
	return s
}

func (s *ListTensorboardsRequest) SetSourceType(v string) *ListTensorboardsRequest {
	s.SourceType = &v
	return s
}

func (s *ListTensorboardsRequest) SetStartTime(v string) *ListTensorboardsRequest {
	s.StartTime = &v
	return s
}

func (s *ListTensorboardsRequest) SetStatus(v string) *ListTensorboardsRequest {
	s.Status = &v
	return s
}

func (s *ListTensorboardsRequest) SetTensorboardId(v string) *ListTensorboardsRequest {
	s.TensorboardId = &v
	return s
}

func (s *ListTensorboardsRequest) SetVerbose(v bool) *ListTensorboardsRequest {
	s.Verbose = &v
	return s
}

func (s *ListTensorboardsRequest) SetWorkspaceId(v string) *ListTensorboardsRequest {
	s.WorkspaceId = &v
	return s
}

type ListTensorboardsResponseBody struct {
	RequestId    *string        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Tensorboards []*Tensorboard `json:"Tensorboards,omitempty" xml:"Tensorboards,omitempty" type:"Repeated"`
	TotalCount   *int64         `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListTensorboardsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTensorboardsResponseBody) GoString() string {
	return s.String()
}

func (s *ListTensorboardsResponseBody) SetRequestId(v string) *ListTensorboardsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTensorboardsResponseBody) SetTensorboards(v []*Tensorboard) *ListTensorboardsResponseBody {
	s.Tensorboards = v
	return s
}

func (s *ListTensorboardsResponseBody) SetTotalCount(v int64) *ListTensorboardsResponseBody {
	s.TotalCount = &v
	return s
}

type ListTensorboardsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTensorboardsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTensorboardsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTensorboardsResponse) GoString() string {
	return s.String()
}

func (s *ListTensorboardsResponse) SetHeaders(v map[string]*string) *ListTensorboardsResponse {
	s.Headers = v
	return s
}

func (s *ListTensorboardsResponse) SetStatusCode(v int32) *ListTensorboardsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTensorboardsResponse) SetBody(v *ListTensorboardsResponseBody) *ListTensorboardsResponse {
	s.Body = v
	return s
}

type StartTensorboardRequest struct {
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s StartTensorboardRequest) String() string {
	return tea.Prettify(s)
}

func (s StartTensorboardRequest) GoString() string {
	return s.String()
}

func (s *StartTensorboardRequest) SetWorkspaceId(v string) *StartTensorboardRequest {
	s.WorkspaceId = &v
	return s
}

type StartTensorboardResponseBody struct {
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TensorboardId *string `json:"TensorboardId,omitempty" xml:"TensorboardId,omitempty"`
}

func (s StartTensorboardResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartTensorboardResponseBody) GoString() string {
	return s.String()
}

func (s *StartTensorboardResponseBody) SetRequestId(v string) *StartTensorboardResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartTensorboardResponseBody) SetTensorboardId(v string) *StartTensorboardResponseBody {
	s.TensorboardId = &v
	return s
}

type StartTensorboardResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartTensorboardResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartTensorboardResponse) String() string {
	return tea.Prettify(s)
}

func (s StartTensorboardResponse) GoString() string {
	return s.String()
}

func (s *StartTensorboardResponse) SetHeaders(v map[string]*string) *StartTensorboardResponse {
	s.Headers = v
	return s
}

func (s *StartTensorboardResponse) SetStatusCode(v int32) *StartTensorboardResponse {
	s.StatusCode = &v
	return s
}

func (s *StartTensorboardResponse) SetBody(v *StartTensorboardResponseBody) *StartTensorboardResponse {
	s.Body = v
	return s
}

type StopJobResponseBody struct {
	JobId     *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopJobResponseBody) GoString() string {
	return s.String()
}

func (s *StopJobResponseBody) SetJobId(v string) *StopJobResponseBody {
	s.JobId = &v
	return s
}

func (s *StopJobResponseBody) SetRequestId(v string) *StopJobResponseBody {
	s.RequestId = &v
	return s
}

type StopJobResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopJobResponse) String() string {
	return tea.Prettify(s)
}

func (s StopJobResponse) GoString() string {
	return s.String()
}

func (s *StopJobResponse) SetHeaders(v map[string]*string) *StopJobResponse {
	s.Headers = v
	return s
}

func (s *StopJobResponse) SetStatusCode(v int32) *StopJobResponse {
	s.StatusCode = &v
	return s
}

func (s *StopJobResponse) SetBody(v *StopJobResponseBody) *StopJobResponse {
	s.Body = v
	return s
}

type StopTensorboardRequest struct {
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s StopTensorboardRequest) String() string {
	return tea.Prettify(s)
}

func (s StopTensorboardRequest) GoString() string {
	return s.String()
}

func (s *StopTensorboardRequest) SetWorkspaceId(v string) *StopTensorboardRequest {
	s.WorkspaceId = &v
	return s
}

type StopTensorboardResponseBody struct {
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TensorboardId *string `json:"TensorboardId,omitempty" xml:"TensorboardId,omitempty"`
}

func (s StopTensorboardResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopTensorboardResponseBody) GoString() string {
	return s.String()
}

func (s *StopTensorboardResponseBody) SetRequestId(v string) *StopTensorboardResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopTensorboardResponseBody) SetTensorboardId(v string) *StopTensorboardResponseBody {
	s.TensorboardId = &v
	return s
}

type StopTensorboardResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopTensorboardResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopTensorboardResponse) String() string {
	return tea.Prettify(s)
}

func (s StopTensorboardResponse) GoString() string {
	return s.String()
}

func (s *StopTensorboardResponse) SetHeaders(v map[string]*string) *StopTensorboardResponse {
	s.Headers = v
	return s
}

func (s *StopTensorboardResponse) SetStatusCode(v int32) *StopTensorboardResponse {
	s.StatusCode = &v
	return s
}

func (s *StopTensorboardResponse) SetBody(v *StopTensorboardResponseBody) *StopTensorboardResponse {
	s.Body = v
	return s
}

type UpdateJobRequest struct {
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
}

func (s UpdateJobRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateJobRequest) GoString() string {
	return s.String()
}

func (s *UpdateJobRequest) SetPriority(v int32) *UpdateJobRequest {
	s.Priority = &v
	return s
}

type UpdateJobResponseBody struct {
	JobId     *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateJobResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateJobResponseBody) SetJobId(v string) *UpdateJobResponseBody {
	s.JobId = &v
	return s
}

func (s *UpdateJobResponseBody) SetRequestId(v string) *UpdateJobResponseBody {
	s.RequestId = &v
	return s
}

type UpdateJobResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateJobResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateJobResponse) GoString() string {
	return s.String()
}

func (s *UpdateJobResponse) SetHeaders(v map[string]*string) *UpdateJobResponse {
	s.Headers = v
	return s
}

func (s *UpdateJobResponse) SetStatusCode(v int32) *UpdateJobResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateJobResponse) SetBody(v *UpdateJobResponseBody) *UpdateJobResponse {
	s.Body = v
	return s
}

type UpdateTensorboardRequest struct {
	MaxRunningTimeMinutes *int64  `json:"MaxRunningTimeMinutes,omitempty" xml:"MaxRunningTimeMinutes,omitempty"`
	WorkspaceId           *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s UpdateTensorboardRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateTensorboardRequest) GoString() string {
	return s.String()
}

func (s *UpdateTensorboardRequest) SetMaxRunningTimeMinutes(v int64) *UpdateTensorboardRequest {
	s.MaxRunningTimeMinutes = &v
	return s
}

func (s *UpdateTensorboardRequest) SetWorkspaceId(v string) *UpdateTensorboardRequest {
	s.WorkspaceId = &v
	return s
}

type UpdateTensorboardResponseBody struct {
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TensorboardId *string `json:"TensorboardId,omitempty" xml:"TensorboardId,omitempty"`
}

func (s UpdateTensorboardResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateTensorboardResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateTensorboardResponseBody) SetRequestId(v string) *UpdateTensorboardResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateTensorboardResponseBody) SetTensorboardId(v string) *UpdateTensorboardResponseBody {
	s.TensorboardId = &v
	return s
}

type UpdateTensorboardResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateTensorboardResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateTensorboardResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateTensorboardResponse) GoString() string {
	return s.String()
}

func (s *UpdateTensorboardResponse) SetHeaders(v map[string]*string) *UpdateTensorboardResponse {
	s.Headers = v
	return s
}

func (s *UpdateTensorboardResponse) SetStatusCode(v int32) *UpdateTensorboardResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateTensorboardResponse) SetBody(v *UpdateTensorboardResponseBody) *UpdateTensorboardResponse {
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
	client.SignatureAlgorithm = tea.String("v2")
	client.EndpointRule = tea.String("regional")
	client.EndpointMap = map[string]*string{
		"ap-northeast-1":              tea.String("pai-dlc.aliyuncs.com"),
		"ap-northeast-2-pop":          tea.String("pai-dlc.aliyuncs.com"),
		"ap-south-1":                  tea.String("pai-dlc.aliyuncs.com"),
		"ap-southeast-2":              tea.String("pai-dlc.aliyuncs.com"),
		"ap-southeast-3":              tea.String("pai-dlc.aliyuncs.com"),
		"ap-southeast-5":              tea.String("pai-dlc.aliyuncs.com"),
		"cn-beijing-finance-1":        tea.String("pai-dlc.aliyuncs.com"),
		"cn-beijing-finance-pop":      tea.String("pai-dlc.aliyuncs.com"),
		"cn-beijing-gov-1":            tea.String("pai-dlc.aliyuncs.com"),
		"cn-beijing-nu16-b01":         tea.String("pai-dlc.aliyuncs.com"),
		"cn-chengdu":                  tea.String("pai-dlc.aliyuncs.com"),
		"cn-edge-1":                   tea.String("pai-dlc.aliyuncs.com"),
		"cn-fujian":                   tea.String("pai-dlc.aliyuncs.com"),
		"cn-haidian-cm12-c01":         tea.String("pai-dlc.aliyuncs.com"),
		"cn-hangzhou-bj-b01":          tea.String("pai-dlc.aliyuncs.com"),
		"cn-hangzhou-finance":         tea.String("pai-dlc.aliyuncs.com"),
		"cn-hangzhou-internal-prod-1": tea.String("pai-dlc.aliyuncs.com"),
		"cn-hangzhou-internal-test-1": tea.String("pai-dlc.aliyuncs.com"),
		"cn-hangzhou-internal-test-2": tea.String("pai-dlc.aliyuncs.com"),
		"cn-hangzhou-internal-test-3": tea.String("pai-dlc.aliyuncs.com"),
		"cn-hangzhou-test-306":        tea.String("pai-dlc.aliyuncs.com"),
		"cn-hongkong-finance-pop":     tea.String("pai-dlc.aliyuncs.com"),
		"cn-huhehaote":                tea.String("pai-dlc.aliyuncs.com"),
		"cn-huhehaote-nebula-1":       tea.String("pai-dlc.aliyuncs.com"),
		"cn-north-2-gov-1":            tea.String("pai-dlc.aliyuncs.com"),
		"cn-qingdao":                  tea.String("pai-dlc.aliyuncs.com"),
		"cn-qingdao-nebula":           tea.String("pai-dlc.aliyuncs.com"),
		"cn-shanghai-et15-b01":        tea.String("pai-dlc.aliyuncs.com"),
		"cn-shanghai-et2-b01":         tea.String("pai-dlc.aliyuncs.com"),
		"cn-shanghai-inner":           tea.String("pai-dlc.aliyuncs.com"),
		"cn-shanghai-internal-test-1": tea.String("pai-dlc.aliyuncs.com"),
		"cn-shenzhen-finance-1":       tea.String("pai-dlc.aliyuncs.com"),
		"cn-shenzhen-inner":           tea.String("pai-dlc.aliyuncs.com"),
		"cn-shenzhen-st4-d01":         tea.String("pai-dlc.aliyuncs.com"),
		"cn-shenzhen-su18-b01":        tea.String("pai-dlc.aliyuncs.com"),
		"cn-wuhan":                    tea.String("pai-dlc.aliyuncs.com"),
		"cn-wulanchabu":               tea.String("pai-dlc.aliyuncs.com"),
		"cn-yushanfang":               tea.String("pai-dlc.aliyuncs.com"),
		"cn-zhangbei":                 tea.String("pai-dlc.aliyuncs.com"),
		"cn-zhangbei-na61-b01":        tea.String("pai-dlc.aliyuncs.com"),
		"cn-zhangjiakou":              tea.String("pai-dlc.aliyuncs.com"),
		"cn-zhangjiakou-na62-a01":     tea.String("pai-dlc.aliyuncs.com"),
		"cn-zhengzhou-nebula-1":       tea.String("pai-dlc.aliyuncs.com"),
		"eu-west-1":                   tea.String("pai-dlc.aliyuncs.com"),
		"eu-west-1-oxs":               tea.String("pai-dlc.aliyuncs.com"),
		"me-east-1":                   tea.String("pai-dlc.aliyuncs.com"),
		"rus-west-1-pop":              tea.String("pai-dlc.aliyuncs.com"),
		"us-east-1":                   tea.String("pai-dlc.aliyuncs.com"),
		"us-west-1":                   tea.String("pai-dlc.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("pai-dlc"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) CreateJobWithOptions(request *CreateJobRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CodeSource)) {
		body["CodeSource"] = request.CodeSource
	}

	if !tea.BoolValue(util.IsUnset(request.DataSources)) {
		body["DataSources"] = request.DataSources
	}

	if !tea.BoolValue(util.IsUnset(request.DebuggerConfigContent)) {
		body["DebuggerConfigContent"] = request.DebuggerConfigContent
	}

	if !tea.BoolValue(util.IsUnset(request.DisplayName)) {
		body["DisplayName"] = request.DisplayName
	}

	if !tea.BoolValue(util.IsUnset(request.ElasticSpec)) {
		body["ElasticSpec"] = request.ElasticSpec
	}

	if !tea.BoolValue(util.IsUnset(request.Envs)) {
		body["Envs"] = request.Envs
	}

	if !tea.BoolValue(util.IsUnset(request.JobMaxRunningTimeMinutes)) {
		body["JobMaxRunningTimeMinutes"] = request.JobMaxRunningTimeMinutes
	}

	if !tea.BoolValue(util.IsUnset(request.JobSpecs)) {
		body["JobSpecs"] = request.JobSpecs
	}

	if !tea.BoolValue(util.IsUnset(request.JobType)) {
		body["JobType"] = request.JobType
	}

	if !tea.BoolValue(util.IsUnset(request.Options)) {
		body["Options"] = request.Options
	}

	if !tea.BoolValue(util.IsUnset(request.Priority)) {
		body["Priority"] = request.Priority
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		body["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.Settings)) {
		body["Settings"] = request.Settings
	}

	if !tea.BoolValue(util.IsUnset(request.SuccessPolicy)) {
		body["SuccessPolicy"] = request.SuccessPolicy
	}

	if !tea.BoolValue(util.IsUnset(request.ThirdpartyLibDir)) {
		body["ThirdpartyLibDir"] = request.ThirdpartyLibDir
	}

	if !tea.BoolValue(util.IsUnset(request.ThirdpartyLibs)) {
		body["ThirdpartyLibs"] = request.ThirdpartyLibs
	}

	if !tea.BoolValue(util.IsUnset(request.UserCommand)) {
		body["UserCommand"] = request.UserCommand
	}

	if !tea.BoolValue(util.IsUnset(request.UserVpc)) {
		body["UserVpc"] = request.UserVpc
	}

	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateJob"),
		Version:     tea.String("2020-12-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/jobs"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateJob(request *CreateJobRequest) (_result *CreateJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateJobResponse{}
	_body, _err := client.CreateJobWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateTensorboardWithOptions(request *CreateTensorboardRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateTensorboardResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Cpu)) {
		body["Cpu"] = request.Cpu
	}

	if !tea.BoolValue(util.IsUnset(request.DataSourceId)) {
		body["DataSourceId"] = request.DataSourceId
	}

	if !tea.BoolValue(util.IsUnset(request.DataSourceType)) {
		body["DataSourceType"] = request.DataSourceType
	}

	if !tea.BoolValue(util.IsUnset(request.DataSources)) {
		body["DataSources"] = request.DataSources
	}

	if !tea.BoolValue(util.IsUnset(request.DisplayName)) {
		body["DisplayName"] = request.DisplayName
	}

	if !tea.BoolValue(util.IsUnset(request.JobId)) {
		body["JobId"] = request.JobId
	}

	if !tea.BoolValue(util.IsUnset(request.MaxRunningTimeMinutes)) {
		body["MaxRunningTimeMinutes"] = request.MaxRunningTimeMinutes
	}

	if !tea.BoolValue(util.IsUnset(request.Memory)) {
		body["Memory"] = request.Memory
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

	if !tea.BoolValue(util.IsUnset(request.SummaryPath)) {
		body["SummaryPath"] = request.SummaryPath
	}

	if !tea.BoolValue(util.IsUnset(request.SummaryRelativePath)) {
		body["SummaryRelativePath"] = request.SummaryRelativePath
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
		Action:      tea.String("CreateTensorboard"),
		Version:     tea.String("2020-12-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/tensorboards"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateTensorboardResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateTensorboard(request *CreateTensorboardRequest) (_result *CreateTensorboardResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateTensorboardResponse{}
	_body, _err := client.CreateTensorboardWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteJobWithOptions(JobId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteJobResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteJob"),
		Version:     tea.String("2020-12-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/jobs/" + tea.StringValue(openapiutil.GetEncodeParam(JobId))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteJob(JobId *string) (_result *DeleteJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteJobResponse{}
	_body, _err := client.DeleteJobWithOptions(JobId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteTensorboardWithOptions(TensorboardId *string, request *DeleteTensorboardRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteTensorboardResponse, _err error) {
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
		Action:      tea.String("DeleteTensorboard"),
		Version:     tea.String("2020-12-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/tensorboards/" + tea.StringValue(openapiutil.GetEncodeParam(TensorboardId))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteTensorboardResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteTensorboard(TensorboardId *string, request *DeleteTensorboardRequest) (_result *DeleteTensorboardResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteTensorboardResponse{}
	_body, _err := client.DeleteTensorboardWithOptions(TensorboardId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetJobWithOptions(JobId *string, request *GetJobRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NeedDetail)) {
		query["NeedDetail"] = request.NeedDetail
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetJob"),
		Version:     tea.String("2020-12-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/jobs/" + tea.StringValue(openapiutil.GetEncodeParam(JobId))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetJob(JobId *string, request *GetJobRequest) (_result *GetJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetJobResponse{}
	_body, _err := client.GetJobWithOptions(JobId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetJobEventsWithOptions(JobId *string, request *GetJobEventsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetJobEventsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.MaxEventsNum)) {
		query["MaxEventsNum"] = request.MaxEventsNum
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetJobEvents"),
		Version:     tea.String("2020-12-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/jobs/" + tea.StringValue(openapiutil.GetEncodeParam(JobId)) + "/events"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetJobEventsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetJobEvents(JobId *string, request *GetJobEventsRequest) (_result *GetJobEventsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetJobEventsResponse{}
	_body, _err := client.GetJobEventsWithOptions(JobId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetJobMetricsWithOptions(JobId *string, request *GetJobMetricsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetJobMetricsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.MetricType)) {
		query["MetricType"] = request.MetricType
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.TimeStep)) {
		query["TimeStep"] = request.TimeStep
	}

	if !tea.BoolValue(util.IsUnset(request.Token)) {
		query["Token"] = request.Token
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetJobMetrics"),
		Version:     tea.String("2020-12-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/jobs/" + tea.StringValue(openapiutil.GetEncodeParam(JobId)) + "/metrics"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetJobMetricsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetJobMetrics(JobId *string, request *GetJobMetricsRequest) (_result *GetJobMetricsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetJobMetricsResponse{}
	_body, _err := client.GetJobMetricsWithOptions(JobId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetJobSanityCheckResultWithOptions(JobId *string, request *GetJobSanityCheckResultRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetJobSanityCheckResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SanityCheckNumber)) {
		query["SanityCheckNumber"] = request.SanityCheckNumber
	}

	if !tea.BoolValue(util.IsUnset(request.SanityCheckPhase)) {
		query["SanityCheckPhase"] = request.SanityCheckPhase
	}

	if !tea.BoolValue(util.IsUnset(request.Token)) {
		query["Token"] = request.Token
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetJobSanityCheckResult"),
		Version:     tea.String("2020-12-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/jobs/" + tea.StringValue(openapiutil.GetEncodeParam(JobId)) + "/sanitycheckresult"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetJobSanityCheckResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetJobSanityCheckResult(JobId *string, request *GetJobSanityCheckResultRequest) (_result *GetJobSanityCheckResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetJobSanityCheckResultResponse{}
	_body, _err := client.GetJobSanityCheckResultWithOptions(JobId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetPodEventsWithOptions(JobId *string, PodId *string, request *GetPodEventsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetPodEventsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.MaxEventsNum)) {
		query["MaxEventsNum"] = request.MaxEventsNum
	}

	if !tea.BoolValue(util.IsUnset(request.PodUid)) {
		query["PodUid"] = request.PodUid
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetPodEvents"),
		Version:     tea.String("2020-12-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/jobs/" + tea.StringValue(openapiutil.GetEncodeParam(JobId)) + "/pods/" + tea.StringValue(openapiutil.GetEncodeParam(PodId)) + "/events"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetPodEventsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetPodEvents(JobId *string, PodId *string, request *GetPodEventsRequest) (_result *GetPodEventsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetPodEventsResponse{}
	_body, _err := client.GetPodEventsWithOptions(JobId, PodId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetPodLogsWithOptions(JobId *string, PodId *string, request *GetPodLogsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetPodLogsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DownloadToFile)) {
		query["DownloadToFile"] = request.DownloadToFile
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.MaxLines)) {
		query["MaxLines"] = request.MaxLines
	}

	if !tea.BoolValue(util.IsUnset(request.PodUid)) {
		query["PodUid"] = request.PodUid
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetPodLogs"),
		Version:     tea.String("2020-12-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/jobs/" + tea.StringValue(openapiutil.GetEncodeParam(JobId)) + "/pods/" + tea.StringValue(openapiutil.GetEncodeParam(PodId)) + "/logs"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetPodLogsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetPodLogs(JobId *string, PodId *string, request *GetPodLogsRequest) (_result *GetPodLogsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetPodLogsResponse{}
	_body, _err := client.GetPodLogsWithOptions(JobId, PodId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetTensorboardWithOptions(TensorboardId *string, request *GetTensorboardRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetTensorboardResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JodId)) {
		query["JodId"] = request.JodId
	}

	if !tea.BoolValue(util.IsUnset(request.Token)) {
		query["Token"] = request.Token
	}

	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetTensorboard"),
		Version:     tea.String("2020-12-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/tensorboards/" + tea.StringValue(openapiutil.GetEncodeParam(TensorboardId))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetTensorboardResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetTensorboard(TensorboardId *string, request *GetTensorboardRequest) (_result *GetTensorboardResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetTensorboardResponse{}
	_body, _err := client.GetTensorboardWithOptions(TensorboardId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetTensorboardSharedUrlWithOptions(TensorboardId *string, request *GetTensorboardSharedUrlRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetTensorboardSharedUrlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ExpireTimeSeconds)) {
		query["ExpireTimeSeconds"] = request.ExpireTimeSeconds
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetTensorboardSharedUrl"),
		Version:     tea.String("2020-12-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/tensorboards/" + tea.StringValue(openapiutil.GetEncodeParam(TensorboardId)) + "/sharedurl"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetTensorboardSharedUrlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetTensorboardSharedUrl(TensorboardId *string, request *GetTensorboardSharedUrlRequest) (_result *GetTensorboardSharedUrlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetTensorboardSharedUrlResponse{}
	_body, _err := client.GetTensorboardSharedUrlWithOptions(TensorboardId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetTokenWithOptions(request *GetTokenRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetTokenResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ExpireTime)) {
		query["ExpireTime"] = request.ExpireTime
	}

	if !tea.BoolValue(util.IsUnset(request.TargetId)) {
		query["TargetId"] = request.TargetId
	}

	if !tea.BoolValue(util.IsUnset(request.TargetType)) {
		query["TargetType"] = request.TargetType
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetToken"),
		Version:     tea.String("2020-12-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/tokens"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetTokenResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetToken(request *GetTokenRequest) (_result *GetTokenResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetTokenResponse{}
	_body, _err := client.GetTokenWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetWebTerminalWithOptions(JobId *string, PodId *string, request *GetWebTerminalRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetWebTerminalResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IsShared)) {
		query["IsShared"] = request.IsShared
	}

	if !tea.BoolValue(util.IsUnset(request.PodUid)) {
		query["PodUid"] = request.PodUid
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetWebTerminal"),
		Version:     tea.String("2020-12-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/jobs/" + tea.StringValue(openapiutil.GetEncodeParam(JobId)) + "/pods/" + tea.StringValue(openapiutil.GetEncodeParam(PodId)) + "/webterminal"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetWebTerminalResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetWebTerminal(JobId *string, PodId *string, request *GetWebTerminalRequest) (_result *GetWebTerminalResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetWebTerminalResponse{}
	_body, _err := client.GetWebTerminalWithOptions(JobId, PodId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListEcsSpecsWithOptions(request *ListEcsSpecsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListEcsSpecsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AcceleratorType)) {
		query["AcceleratorType"] = request.AcceleratorType
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceTypes)) {
		query["InstanceTypes"] = request.InstanceTypes
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

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.SortBy)) {
		query["SortBy"] = request.SortBy
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListEcsSpecs"),
		Version:     tea.String("2020-12-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/ecsspecs"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListEcsSpecsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListEcsSpecs(request *ListEcsSpecsRequest) (_result *ListEcsSpecsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListEcsSpecsResponse{}
	_body, _err := client.ListEcsSpecsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListJobSanityCheckResultsWithOptions(JobId *string, request *ListJobSanityCheckResultsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListJobSanityCheckResultsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Order)) {
		query["Order"] = request.Order
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListJobSanityCheckResults"),
		Version:     tea.String("2020-12-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/jobs/" + tea.StringValue(openapiutil.GetEncodeParam(JobId)) + "/sanitycheckresults"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListJobSanityCheckResultsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListJobSanityCheckResults(JobId *string, request *ListJobSanityCheckResultsRequest) (_result *ListJobSanityCheckResultsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListJobSanityCheckResultsResponse{}
	_body, _err := client.ListJobSanityCheckResultsWithOptions(JobId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListJobsWithOptions(tmpReq *ListJobsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListJobsResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListJobsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Tags)) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, tea.String("Tags"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BusinessUserId)) {
		query["BusinessUserId"] = request.BusinessUserId
	}

	if !tea.BoolValue(util.IsUnset(request.Caller)) {
		query["Caller"] = request.Caller
	}

	if !tea.BoolValue(util.IsUnset(request.DisplayName)) {
		query["DisplayName"] = request.DisplayName
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.FromAllWorkspaces)) {
		query["FromAllWorkspaces"] = request.FromAllWorkspaces
	}

	if !tea.BoolValue(util.IsUnset(request.JobId)) {
		query["JobId"] = request.JobId
	}

	if !tea.BoolValue(util.IsUnset(request.JobType)) {
		query["JobType"] = request.JobType
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

	if !tea.BoolValue(util.IsUnset(request.PipelineId)) {
		query["PipelineId"] = request.PipelineId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ShowOwn)) {
		query["ShowOwn"] = request.ShowOwn
	}

	if !tea.BoolValue(util.IsUnset(request.SortBy)) {
		query["SortBy"] = request.SortBy
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.TagsShrink)) {
		query["Tags"] = request.TagsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.UserIdForFilter)) {
		query["UserIdForFilter"] = request.UserIdForFilter
	}

	if !tea.BoolValue(util.IsUnset(request.Username)) {
		query["Username"] = request.Username
	}

	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListJobs"),
		Version:     tea.String("2020-12-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/jobs"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListJobsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListJobs(request *ListJobsRequest) (_result *ListJobsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListJobsResponse{}
	_body, _err := client.ListJobsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListTensorboardsWithOptions(request *ListTensorboardsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListTensorboardsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DisplayName)) {
		query["DisplayName"] = request.DisplayName
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.JobId)) {
		query["JobId"] = request.JobId
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

	if !tea.BoolValue(util.IsUnset(request.ShowOwn)) {
		query["ShowOwn"] = request.ShowOwn
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

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.TensorboardId)) {
		query["TensorboardId"] = request.TensorboardId
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
		Action:      tea.String("ListTensorboards"),
		Version:     tea.String("2020-12-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/tensorboards"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTensorboardsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListTensorboards(request *ListTensorboardsRequest) (_result *ListTensorboardsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListTensorboardsResponse{}
	_body, _err := client.ListTensorboardsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StartTensorboardWithOptions(TensorboardId *string, request *StartTensorboardRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *StartTensorboardResponse, _err error) {
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
		Action:      tea.String("StartTensorboard"),
		Version:     tea.String("2020-12-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/tensorboards/" + tea.StringValue(openapiutil.GetEncodeParam(TensorboardId)) + "/start"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &StartTensorboardResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StartTensorboard(TensorboardId *string, request *StartTensorboardRequest) (_result *StartTensorboardResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &StartTensorboardResponse{}
	_body, _err := client.StartTensorboardWithOptions(TensorboardId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StopJobWithOptions(JobId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *StopJobResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("StopJob"),
		Version:     tea.String("2020-12-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/jobs/" + tea.StringValue(openapiutil.GetEncodeParam(JobId)) + "/stop"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &StopJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StopJob(JobId *string) (_result *StopJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &StopJobResponse{}
	_body, _err := client.StopJobWithOptions(JobId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StopTensorboardWithOptions(TensorboardId *string, request *StopTensorboardRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *StopTensorboardResponse, _err error) {
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
		Action:      tea.String("StopTensorboard"),
		Version:     tea.String("2020-12-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/tensorboards/" + tea.StringValue(openapiutil.GetEncodeParam(TensorboardId)) + "/stop"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &StopTensorboardResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StopTensorboard(TensorboardId *string, request *StopTensorboardRequest) (_result *StopTensorboardResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &StopTensorboardResponse{}
	_body, _err := client.StopTensorboardWithOptions(TensorboardId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateJobWithOptions(JobId *string, request *UpdateJobRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Priority)) {
		body["Priority"] = request.Priority
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateJob"),
		Version:     tea.String("2020-12-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/jobs/" + tea.StringValue(openapiutil.GetEncodeParam(JobId))),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateJob(JobId *string, request *UpdateJobRequest) (_result *UpdateJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateJobResponse{}
	_body, _err := client.UpdateJobWithOptions(JobId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateTensorboardWithOptions(TensorboardId *string, request *UpdateTensorboardRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateTensorboardResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MaxRunningTimeMinutes)) {
		query["MaxRunningTimeMinutes"] = request.MaxRunningTimeMinutes
	}

	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateTensorboard"),
		Version:     tea.String("2020-12-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/tensorboards/" + tea.StringValue(openapiutil.GetEncodeParam(TensorboardId))),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateTensorboardResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateTensorboard(TensorboardId *string, request *UpdateTensorboardRequest) (_result *UpdateTensorboardResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateTensorboardResponse{}
	_body, _err := client.UpdateTensorboardWithOptions(TensorboardId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
