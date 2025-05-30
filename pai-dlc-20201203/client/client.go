// This file is auto-generated, don't edit it. Thanks.
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

type AssignNodeSpec struct {
	// example:
	//
	// lingjxxxxxxxx
	AntiAffinityNodeNames *string `json:"AntiAffinityNodeNames,omitempty" xml:"AntiAffinityNodeNames,omitempty"`
	// example:
	//
	// true
	EnableAssignNode *bool `json:"EnableAssignNode,omitempty" xml:"EnableAssignNode,omitempty"`
	// example:
	//
	// lingjxxxxxxxx
	NodeNames *string `json:"NodeNames,omitempty" xml:"NodeNames,omitempty"`
}

func (s AssignNodeSpec) String() string {
	return tea.Prettify(s)
}

func (s AssignNodeSpec) GoString() string {
	return s.String()
}

func (s *AssignNodeSpec) SetAntiAffinityNodeNames(v string) *AssignNodeSpec {
	s.AntiAffinityNodeNames = &v
	return s
}

func (s *AssignNodeSpec) SetEnableAssignNode(v bool) *AssignNodeSpec {
	s.EnableAssignNode = &v
	return s
}

func (s *AssignNodeSpec) SetNodeNames(v string) *AssignNodeSpec {
	s.NodeNames = &v
	return s
}

type AssumeUserInfo struct {
	AccessKeyId   *string `json:"AccessKeyId,omitempty" xml:"AccessKeyId,omitempty"`
	Id            *string `json:"Id,omitempty" xml:"Id,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	Type          *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s AssumeUserInfo) String() string {
	return tea.Prettify(s)
}

func (s AssumeUserInfo) GoString() string {
	return s.String()
}

func (s *AssumeUserInfo) SetAccessKeyId(v string) *AssumeUserInfo {
	s.AccessKeyId = &v
	return s
}

func (s *AssumeUserInfo) SetId(v string) *AssumeUserInfo {
	s.Id = &v
	return s
}

func (s *AssumeUserInfo) SetSecurityToken(v string) *AssumeUserInfo {
	s.SecurityToken = &v
	return s
}

func (s *AssumeUserInfo) SetType(v string) *AssumeUserInfo {
	s.Type = &v
	return s
}

type AutoScalingSpec struct {
	ScalingStrategy *string `json:"ScalingStrategy,omitempty" xml:"ScalingStrategy,omitempty"`
}

func (s AutoScalingSpec) String() string {
	return tea.Prettify(s)
}

func (s AutoScalingSpec) GoString() string {
	return s.String()
}

func (s *AutoScalingSpec) SetScalingStrategy(v string) *AutoScalingSpec {
	s.ScalingStrategy = &v
	return s
}

type CodeSourceItem struct {
	// example:
	//
	// master
	CodeBranch *string `json:"CodeBranch,omitempty" xml:"CodeBranch,omitempty"`
	// example:
	//
	// 44da1*******
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
	// code-20210111103721-85qz*****
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
	// 115**********
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
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
	Args    []*string `json:"Args,omitempty" xml:"Args,omitempty" type:"Repeated"`
	Command []*string `json:"Command,omitempty" xml:"Command,omitempty" type:"Repeated"`
	Env     []*EnvVar `json:"Env,omitempty" xml:"Env,omitempty" type:"Repeated"`
	// example:
	//
	// registry.cn-hangzhou.aliyuncs.com/pai-dlc/curl:v1.0.0
	Image *string `json:"Image,omitempty" xml:"Image,omitempty"`
	// example:
	//
	// data-init
	Name      *string               `json:"Name,omitempty" xml:"Name,omitempty"`
	Resources *ResourceRequirements `json:"Resources,omitempty" xml:"Resources,omitempty"`
	// example:
	//
	// /root
	WorkingDir *string `json:"WorkingDir,omitempty" xml:"WorkingDir,omitempty"`
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

type CredentialConfig struct {
	AliyunEnvRoleKey       *string                 `json:"AliyunEnvRoleKey,omitempty" xml:"AliyunEnvRoleKey,omitempty"`
	CredentialConfigItems  []*CredentialConfigItem `json:"CredentialConfigItems,omitempty" xml:"CredentialConfigItems,omitempty" type:"Repeated"`
	EnableCredentialInject *bool                   `json:"EnableCredentialInject,omitempty" xml:"EnableCredentialInject,omitempty"`
}

func (s CredentialConfig) String() string {
	return tea.Prettify(s)
}

func (s CredentialConfig) GoString() string {
	return s.String()
}

func (s *CredentialConfig) SetAliyunEnvRoleKey(v string) *CredentialConfig {
	s.AliyunEnvRoleKey = &v
	return s
}

func (s *CredentialConfig) SetCredentialConfigItems(v []*CredentialConfigItem) *CredentialConfig {
	s.CredentialConfigItems = v
	return s
}

func (s *CredentialConfig) SetEnableCredentialInject(v bool) *CredentialConfig {
	s.EnableCredentialInject = &v
	return s
}

type CredentialConfigItem struct {
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// if can be null:
	// true
	Roles []*CredentialRole `json:"Roles,omitempty" xml:"Roles,omitempty" type:"Repeated"`
	Type  *string           `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CredentialConfigItem) String() string {
	return tea.Prettify(s)
}

func (s CredentialConfigItem) GoString() string {
	return s.String()
}

func (s *CredentialConfigItem) SetKey(v string) *CredentialConfigItem {
	s.Key = &v
	return s
}

func (s *CredentialConfigItem) SetRoles(v []*CredentialRole) *CredentialConfigItem {
	s.Roles = v
	return s
}

func (s *CredentialConfigItem) SetType(v string) *CredentialConfigItem {
	s.Type = &v
	return s
}

type CredentialRole struct {
	AssumeRoleFor  *string         `json:"AssumeRoleFor,omitempty" xml:"AssumeRoleFor,omitempty"`
	AssumeUserInfo *AssumeUserInfo `json:"AssumeUserInfo,omitempty" xml:"AssumeUserInfo,omitempty"`
	Policy         *string         `json:"Policy,omitempty" xml:"Policy,omitempty"`
	RoleArn        *string         `json:"RoleArn,omitempty" xml:"RoleArn,omitempty"`
	RoleType       *string         `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
}

func (s CredentialRole) String() string {
	return tea.Prettify(s)
}

func (s CredentialRole) GoString() string {
	return s.String()
}

func (s *CredentialRole) SetAssumeRoleFor(v string) *CredentialRole {
	s.AssumeRoleFor = &v
	return s
}

func (s *CredentialRole) SetAssumeUserInfo(v *AssumeUserInfo) *CredentialRole {
	s.AssumeUserInfo = v
	return s
}

func (s *CredentialRole) SetPolicy(v string) *CredentialRole {
	s.Policy = &v
	return s
}

func (s *CredentialRole) SetRoleArn(v string) *CredentialRole {
	s.RoleArn = &v
	return s
}

func (s *CredentialRole) SetRoleType(v string) *CredentialRole {
	s.RoleType = &v
	return s
}

type DataSourceItem struct {
	// example:
	//
	// data-20210114104214-vf9lowjt3pso
	DataSourceId *string `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	// example:
	//
	// nas
	DataSourceType *string `json:"DataSourceType,omitempty" xml:"DataSourceType,omitempty"`
	// example:
	//
	// data source of dlc examples
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// nas-data
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// example:
	//
	// oss-cn-beijing-internal.aliyuncs.com
	Endpoint *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	// example:
	//
	// 1ca404****
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// example:
	//
	// 2021-01-12T14:35:00Z
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// example:
	//
	// 2021-01-12T14:36:00Z
	GmtModifyTime *string `json:"GmtModifyTime,omitempty" xml:"GmtModifyTime,omitempty"`
	// example:
	//
	// /root/data/
	MountPath *string `json:"MountPath,omitempty" xml:"MountPath,omitempty"`
	// example:
	//
	// {"key": "value"}
	Options *string `json:"Options,omitempty" xml:"Options,omitempty"`
	// example:
	//
	// oss://mybucket/path/to/dir
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// example:
	//
	// 123456789
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
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
	// example:
	//
	// {\"description\":\"这是一个新的pytorchjob模板\"}
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// dc-vf9lowjt3pso
	DebuggerConfigId *string `json:"DebuggerConfigId,omitempty" xml:"DebuggerConfigId,omitempty"`
	// example:
	//
	// 这是一个Pytorch的基础配置模板
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// Pytorch Experiment Config
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// example:
	//
	// 2021-01-12T14:35:00Z
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// example:
	//
	// 2021-01-12T14:36:00Z
	GmtModifyTime *string `json:"GmtModifyTime,omitempty" xml:"GmtModifyTime,omitempty"`
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
	// example:
	//
	// dlc20210126170216-mtl37ge7gkvdz
	DebuggerJobId *string `json:"DebuggerJobId,omitempty" xml:"DebuggerJobId,omitempty"`
	// example:
	//
	// dlc debugger test
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// example:
	//
	// 2932
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// example:
	//
	// 2021-01-12T14:35:00Z
	GmtCreateTime    *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	GmtFailedTime    *string `json:"GmtFailedTime,omitempty" xml:"GmtFailedTime,omitempty"`
	GmtFinishTime    *string `json:"GmtFinishTime,omitempty" xml:"GmtFinishTime,omitempty"`
	GmtRunningTime   *string `json:"GmtRunningTime,omitempty" xml:"GmtRunningTime,omitempty"`
	GmtStoppedTime   *string `json:"GmtStoppedTime,omitempty" xml:"GmtStoppedTime,omitempty"`
	GmtSubmittedTime *string `json:"GmtSubmittedTime,omitempty" xml:"GmtSubmittedTime,omitempty"`
	GmtSucceedTime   *string `json:"GmtSucceedTime,omitempty" xml:"GmtSucceedTime,omitempty"`
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 12344556
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// example:
	//
	// workspace01
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
	// example:
	//
	// public
	WorkspaceName *string `json:"WorkspaceName,omitempty" xml:"WorkspaceName,omitempty"`
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
	// example:
	//
	// {"Name": "CPUBottleneck",  "Triggered": 10, "Violations": 2,  "Details": "{}"}
	DebuggerJobIssue *string `json:"DebuggerJobIssue,omitempty" xml:"DebuggerJobIssue,omitempty"`
	// example:
	//
	// 2021-01-12T14:35:00Z
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// example:
	//
	// de-826ca1bcfba30
	JobDebuggerIssueId *string `json:"JobDebuggerIssueId,omitempty" xml:"JobDebuggerIssueId,omitempty"`
	// example:
	//
	// dlc-20210126170216-mtl37ge7gkvdz
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// example:
	//
	// 1002300
	ReasonCode *string `json:"ReasonCode,omitempty" xml:"ReasonCode,omitempty"`
	// example:
	//
	// GPU利用率低
	ReasonMessage *string `json:"ReasonMessage,omitempty" xml:"ReasonMessage,omitempty"`
	// example:
	//
	// ProfileReport
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
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
	// example:
	//
	// {\"description\":\"这是一个新的pytorchjob模板\"}
	DebuggerConfigContent *string `json:"DebuggerConfigContent,omitempty" xml:"DebuggerConfigContent,omitempty"`
	// example:
	//
	// { "ProfileReport": { "Name": "CPUBottleneck","Triggered": 10,"Violations": 2,"Details": "{}"}, "LowCPU": { "Name": "CPUBottleneck","Triggered": 10,"Violations": 2,"Details": "{}"}}
	DebuggerJobIssues *string `json:"DebuggerJobIssues,omitempty" xml:"DebuggerJobIssues,omitempty"`
	// example:
	//
	// {"Running": 1, "Failed": 1, "Succeeded": 2}
	DebuggerJobStatus *string `json:"DebuggerJobStatus,omitempty" xml:"DebuggerJobStatus,omitempty"`
	// example:
	//
	// http://xxx.com/debug/report/download/new_xxxx.html
	DebuggerReportURL *string `json:"DebuggerReportURL,omitempty" xml:"DebuggerReportURL,omitempty"`
	// example:
	//
	// dlc debugger test
	JobDisplayName *string `json:"JobDisplayName,omitempty" xml:"JobDisplayName,omitempty"`
	// example:
	//
	// dlc-20210126170216-mtl37ge7gkvdz
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// example:
	//
	// 12344556
	JobUserId *string `json:"JobUserId,omitempty" xml:"JobUserId,omitempty"`
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
	// example:
	//
	// GPU
	AcceleratorType *string `json:"AcceleratorType,omitempty" xml:"AcceleratorType,omitempty"`
	// example:
	//
	// 12
	Cpu *int32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// example:
	//
	// 470.199.02
	DefaultGPUDriver *string `json:"DefaultGPUDriver,omitempty" xml:"DefaultGPUDriver,omitempty"`
	// example:
	//
	// 1
	Gpu *int32 `json:"Gpu,omitempty" xml:"Gpu,omitempty"`
	// example:
	//
	// 80
	GpuMemory *int32 `json:"GpuMemory,omitempty" xml:"GpuMemory,omitempty"`
	// example:
	//
	// NVIDIA v100
	GpuType *string `json:"GpuType,omitempty" xml:"GpuType,omitempty"`
	// example:
	//
	// ecs.gn6e-c12g1.3xlarge
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// example:
	//
	// true
	IsAvailable *bool `json:"IsAvailable,omitempty" xml:"IsAvailable,omitempty"`
	// example:
	//
	// 92
	Memory *int32 `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// example:
	//
	// 0.1
	NonProtectSpotDiscount *float32  `json:"NonProtectSpotDiscount,omitempty" xml:"NonProtectSpotDiscount,omitempty"`
	PaymentTypes           []*string `json:"PaymentTypes,omitempty" xml:"PaymentTypes,omitempty" type:"Repeated"`
	// example:
	//
	// ECS
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// example:
	//
	// WithStock
	SpotStockStatus     *string   `json:"SpotStockStatus,omitempty" xml:"SpotStockStatus,omitempty"`
	SupportedGPUDrivers []*string `json:"SupportedGPUDrivers,omitempty" xml:"SupportedGPUDrivers,omitempty" type:"Repeated"`
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

func (s *EcsSpec) SetGpuMemory(v int32) *EcsSpec {
	s.GpuMemory = &v
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
	// example:
	//
	// ENABLE_DEBUG
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// true
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
	InitContainers []*ContainerSpec `json:"InitContainers,omitempty" xml:"InitContainers,omitempty" type:"Repeated"`
	Lifecycle      *Lifecycle       `json:"Lifecycle,omitempty" xml:"Lifecycle,omitempty"`
	// Deprecated
	PodAnnotations map[string]*string `json:"PodAnnotations,omitempty" xml:"PodAnnotations,omitempty"`
	// Deprecated
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

func (s *ExtraPodSpec) SetLifecycle(v *Lifecycle) *ExtraPodSpec {
	s.Lifecycle = v
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
	ClusterID          *string `json:"ClusterID,omitempty" xml:"ClusterID,omitempty"`
	ClusterName        *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	CrossClusters      *bool   `json:"CrossClusters,omitempty" xml:"CrossClusters,omitempty"`
	EnableFreeResource *bool   `json:"EnableFreeResource,omitempty" xml:"EnableFreeResource,omitempty"`
	// example:
	//
	// frcc-whateversth
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
	// example:
	//
	// 2
	Amount *int32 `json:"Amount,omitempty" xml:"Amount,omitempty"`
	// example:
	//
	// CPU
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
	// example:
	//
	// 2
	AvailableNumber *int64  `json:"AvailableNumber,omitempty" xml:"AvailableNumber,omitempty"`
	ClusterID       *string `json:"ClusterID,omitempty" xml:"ClusterID,omitempty"`
	ClusterName     *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// example:
	//
	// freeres-whateversth
	FreeResourceId *string `json:"FreeResourceId,omitempty" xml:"FreeResourceId,omitempty"`
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
	// inner
	RegionID *string `json:"RegionID,omitempty" xml:"RegionID,omitempty"`
	// example:
	//
	// cpu
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
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
	GPU *string `json:"GPU,omitempty" xml:"GPU,omitempty"`
	// example:
	//
	// Tesla-V100-32G
	GPUType *string `json:"GPUType,omitempty" xml:"GPUType,omitempty"`
	// example:
	//
	// nvidia.com/gpu-tesla-v100-sxm2-16gb
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
	Auth *string `json:"Auth,omitempty" xml:"Auth,omitempty"`
	// example:
	//
	// registry.cn-hangzhou.aliyuncs.com
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
	// example:
	//
	// gpu
	AcceleratorType *string `json:"AcceleratorType,omitempty" xml:"AcceleratorType,omitempty"`
	// example:
	//
	// ken
	AuthorId *string `json:"AuthorId,omitempty" xml:"AuthorId,omitempty"`
	// example:
	//
	// PyTorchJob
	Framework *string `json:"Framework,omitempty" xml:"Framework,omitempty"`
	// example:
	//
	// Community
	ImageProviderType *string `json:"ImageProviderType,omitempty" xml:"ImageProviderType,omitempty"`
	// example:
	//
	// tensorflow-training:2.3-cpu-py36-ubuntu18.04
	ImageTag *string `json:"ImageTag,omitempty" xml:"ImageTag,omitempty"`
	// example:
	//
	// registry.cn-beijing.aliyuncs.com/pai-dlc/tensorflow-training:2.3-cpu-py36-ubuntu18.04
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	// example:
	//
	// registry-vpc.cn-beijing.aliyuncs.com/pai-dlc/tensorflow-training:2.3-cpu-py36-ubuntu18.04
	ImageUrlVpc *string `json:"ImageUrlVpc,omitempty" xml:"ImageUrlVpc,omitempty"`
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
	AIMasterDockerImage *string `json:"AIMasterDockerImage,omitempty" xml:"AIMasterDockerImage,omitempty"`
	AIMasterType        *string `json:"AIMasterType,omitempty" xml:"AIMasterType,omitempty"`
	// example:
	//
	// 16
	EDPMaxParallelism *int32 `json:"EDPMaxParallelism,omitempty" xml:"EDPMaxParallelism,omitempty"`
	// example:
	//
	// 8
	EDPMinParallelism *int32  `json:"EDPMinParallelism,omitempty" xml:"EDPMinParallelism,omitempty"`
	ElasticStrategy   *string `json:"ElasticStrategy,omitempty" xml:"ElasticStrategy,omitempty"`
	EnableAIMaster    *bool   `json:"EnableAIMaster,omitempty" xml:"EnableAIMaster,omitempty"`
	// example:
	//
	// true
	EnableEDP *bool `json:"EnableEDP,omitempty" xml:"EnableEDP,omitempty"`
	// example:
	//
	// true
	EnableElasticTraining *bool `json:"EnableElasticTraining,omitempty" xml:"EnableElasticTraining,omitempty"`
	// example:
	//
	// true
	EnablePsJobElasticPS     *bool `json:"EnablePsJobElasticPS,omitempty" xml:"EnablePsJobElasticPS,omitempty"`
	EnablePsJobElasticWorker *bool `json:"EnablePsJobElasticWorker,omitempty" xml:"EnablePsJobElasticWorker,omitempty"`
	// example:
	//
	// true
	EnablePsResourceEstimate *bool `json:"EnablePsResourceEstimate,omitempty" xml:"EnablePsResourceEstimate,omitempty"`
	// example:
	//
	// 8
	MaxParallelism *int32 `json:"MaxParallelism,omitempty" xml:"MaxParallelism,omitempty"`
	// example:
	//
	// 1
	MinParallelism *int32 `json:"MinParallelism,omitempty" xml:"MinParallelism,omitempty"`
	// example:
	//
	// 10
	PSMaxParallelism *int32 `json:"PSMaxParallelism,omitempty" xml:"PSMaxParallelism,omitempty"`
	// example:
	//
	// 4
	PSMinParallelism *int32 `json:"PSMinParallelism,omitempty" xml:"PSMinParallelism,omitempty"`
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
	// example:
	//
	// PUBLIC
	Accessibility    *string               `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	ClusterId        *string               `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	CodeSource       *JobItemCodeSource    `json:"CodeSource,omitempty" xml:"CodeSource,omitempty" type:"Struct"`
	CredentialConfig *CredentialConfig     `json:"CredentialConfig,omitempty" xml:"CredentialConfig,omitempty"`
	DataSources      []*JobItemDataSources `json:"DataSources,omitempty" xml:"DataSources,omitempty" type:"Repeated"`
	// example:
	//
	// tf-mnist-test
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// example:
	//
	// 3602
	Duration    *int64          `json:"Duration,omitempty" xml:"Duration,omitempty"`
	ElasticSpec *JobElasticSpec `json:"ElasticSpec,omitempty" xml:"ElasticSpec,omitempty"`
	// example:
	//
	// false
	EnablePreemptibleJob *bool `json:"EnablePreemptibleJob,omitempty" xml:"EnablePreemptibleJob,omitempty"`
	// example:
	//
	// false
	EnabledDebugger *bool              `json:"EnabledDebugger,omitempty" xml:"EnabledDebugger,omitempty"`
	Envs            map[string]*string `json:"Envs,omitempty" xml:"Envs,omitempty"`
	// example:
	//
	// 2021-01-12T14:35:01Z
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// example:
	//
	// 2021-01-12T14:35:01Z
	GmtFailedTime *string `json:"GmtFailedTime,omitempty" xml:"GmtFailedTime,omitempty"`
	// example:
	//
	// 2021-01-12T15:36:08Z
	GmtFinishTime *string `json:"GmtFinishTime,omitempty" xml:"GmtFinishTime,omitempty"`
	// example:
	//
	// 2021-01-12T15:36:08Z
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// example:
	//
	// 2021-01-12T14:35:01Z
	GmtRunningTime *string `json:"GmtRunningTime,omitempty" xml:"GmtRunningTime,omitempty"`
	// example:
	//
	// 2021-01-12T14:35:01Z
	GmtStoppedTime *string `json:"GmtStoppedTime,omitempty" xml:"GmtStoppedTime,omitempty"`
	// example:
	//
	// 2021-01-12T14:35:01Z
	GmtSubmittedTime *string `json:"GmtSubmittedTime,omitempty" xml:"GmtSubmittedTime,omitempty"`
	// example:
	//
	// 2021-01-12T14:35:01Z
	GmtSuccessedTime *string `json:"GmtSuccessedTime,omitempty" xml:"GmtSuccessedTime,omitempty"`
	// example:
	//
	// false
	IsDeleted *bool `json:"IsDeleted,omitempty" xml:"IsDeleted,omitempty"`
	// example:
	//
	// dlc-20210126170216-mtl37ge7gkvdz
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// example:
	//
	// 1
	JobMaxRunningTimeMinutes *int64     `json:"JobMaxRunningTimeMinutes,omitempty" xml:"JobMaxRunningTimeMinutes,omitempty"`
	JobSpecs                 []*JobSpec `json:"JobSpecs,omitempty" xml:"JobSpecs,omitempty" type:"Repeated"`
	// example:
	//
	// TFJob
	JobType *string `json:"JobType,omitempty" xml:"JobType,omitempty"`
	// example:
	//
	// 1
	NodeCount *string    `json:"NodeCount,omitempty" xml:"NodeCount,omitempty"`
	NodeNames []*string  `json:"NodeNames,omitempty" xml:"NodeNames,omitempty" type:"Repeated"`
	Pods      []*PodItem `json:"Pods,omitempty" xml:"Pods,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// example:
	//
	// JobStoppedByUser
	ReasonCode *string `json:"ReasonCode,omitempty" xml:"ReasonCode,omitempty"`
	// example:
	//
	// Job is stopped by user.
	ReasonMessage *string `json:"ReasonMessage,omitempty" xml:"ReasonMessage,omitempty"`
	// example:
	//
	// 1
	RequestCPU *int64 `json:"RequestCPU,omitempty" xml:"RequestCPU,omitempty"`
	// example:
	//
	// 1
	RequestGPU *string `json:"RequestGPU,omitempty" xml:"RequestGPU,omitempty"`
	// example:
	//
	// 1Gi
	RequestMemory *string `json:"RequestMemory,omitempty" xml:"RequestMemory,omitempty"`
	// example:
	//
	// dlc-quota
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// example:
	//
	// L0
	ResourceLevel *string `json:"ResourceLevel,omitempty" xml:"ResourceLevel,omitempty"`
	// example:
	//
	// my_resource_group
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	// example:
	//
	// test
	ResourceQuotaName *string `json:"ResourceQuotaName,omitempty" xml:"ResourceQuotaName,omitempty"`
	// example:
	//
	// ECS
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// example:
	//
	// 1
	RestartTimes *string      `json:"RestartTimes,omitempty" xml:"RestartTimes,omitempty"`
	Settings     *JobSettings `json:"Settings,omitempty" xml:"Settings,omitempty"`
	// example:
	//
	// Stopped
	Status        *string                 `json:"Status,omitempty" xml:"Status,omitempty"`
	StatusHistory []*StatusTransitionItem `json:"StatusHistory,omitempty" xml:"StatusHistory,omitempty" type:"Repeated"`
	// example:
	//
	// Restarting
	SubStatus  *string            `json:"SubStatus,omitempty" xml:"SubStatus,omitempty"`
	SystemEnvs map[string]*string `json:"SystemEnvs,omitempty" xml:"SystemEnvs,omitempty"`
	TenantId   *string            `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// example:
	//
	// /root/code/
	ThirdpartyLibDir *string   `json:"ThirdpartyLibDir,omitempty" xml:"ThirdpartyLibDir,omitempty"`
	ThirdpartyLibs   []*string `json:"ThirdpartyLibs,omitempty" xml:"ThirdpartyLibs,omitempty" type:"Repeated"`
	// example:
	//
	// false
	UseOversoldResource *bool `json:"UseOversoldResource,omitempty" xml:"UseOversoldResource,omitempty"`
	// example:
	//
	// python /root/code/mnist.py
	UserCommand *string `json:"UserCommand,omitempty" xml:"UserCommand,omitempty"`
	// example:
	//
	// 123456789
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// example:
	//
	// ls
	UserScript *string `json:"UserScript,omitempty" xml:"UserScript,omitempty"`
	// example:
	//
	// vpc-1
	UserVpc *JobItemUserVpc `json:"UserVpc,omitempty" xml:"UserVpc,omitempty" type:"Struct"`
	// example:
	//
	// pai-dlc-role
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
	// example:
	//
	// /mnt/data
	WorkingDir *string `json:"WorkingDir,omitempty" xml:"WorkingDir,omitempty"`
	// example:
	//
	// 268
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
	// example:
	//
	// dlc-workspace
	WorkspaceName *string `json:"WorkspaceName,omitempty" xml:"WorkspaceName,omitempty"`
}

func (s JobItem) String() string {
	return tea.Prettify(s)
}

func (s JobItem) GoString() string {
	return s.String()
}

func (s *JobItem) SetAccessibility(v string) *JobItem {
	s.Accessibility = &v
	return s
}

func (s *JobItem) SetClusterId(v string) *JobItem {
	s.ClusterId = &v
	return s
}

func (s *JobItem) SetCodeSource(v *JobItemCodeSource) *JobItem {
	s.CodeSource = v
	return s
}

func (s *JobItem) SetCredentialConfig(v *CredentialConfig) *JobItem {
	s.CredentialConfig = v
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

func (s *JobItem) SetElasticSpec(v *JobElasticSpec) *JobItem {
	s.ElasticSpec = v
	return s
}

func (s *JobItem) SetEnablePreemptibleJob(v bool) *JobItem {
	s.EnablePreemptibleJob = &v
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

func (s *JobItem) SetGmtModifiedTime(v string) *JobItem {
	s.GmtModifiedTime = &v
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

func (s *JobItem) SetIsDeleted(v bool) *JobItem {
	s.IsDeleted = &v
	return s
}

func (s *JobItem) SetJobId(v string) *JobItem {
	s.JobId = &v
	return s
}

func (s *JobItem) SetJobMaxRunningTimeMinutes(v int64) *JobItem {
	s.JobMaxRunningTimeMinutes = &v
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

func (s *JobItem) SetNodeCount(v string) *JobItem {
	s.NodeCount = &v
	return s
}

func (s *JobItem) SetNodeNames(v []*string) *JobItem {
	s.NodeNames = v
	return s
}

func (s *JobItem) SetPods(v []*PodItem) *JobItem {
	s.Pods = v
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

func (s *JobItem) SetRequestCPU(v int64) *JobItem {
	s.RequestCPU = &v
	return s
}

func (s *JobItem) SetRequestGPU(v string) *JobItem {
	s.RequestGPU = &v
	return s
}

func (s *JobItem) SetRequestMemory(v string) *JobItem {
	s.RequestMemory = &v
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

func (s *JobItem) SetResourceQuotaName(v string) *JobItem {
	s.ResourceQuotaName = &v
	return s
}

func (s *JobItem) SetResourceType(v string) *JobItem {
	s.ResourceType = &v
	return s
}

func (s *JobItem) SetRestartTimes(v string) *JobItem {
	s.RestartTimes = &v
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

func (s *JobItem) SetStatusHistory(v []*StatusTransitionItem) *JobItem {
	s.StatusHistory = v
	return s
}

func (s *JobItem) SetSubStatus(v string) *JobItem {
	s.SubStatus = &v
	return s
}

func (s *JobItem) SetSystemEnvs(v map[string]*string) *JobItem {
	s.SystemEnvs = v
	return s
}

func (s *JobItem) SetTenantId(v string) *JobItem {
	s.TenantId = &v
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

func (s *JobItem) SetUserScript(v string) *JobItem {
	s.UserScript = &v
	return s
}

func (s *JobItem) SetUserVpc(v *JobItemUserVpc) *JobItem {
	s.UserVpc = v
	return s
}

func (s *JobItem) SetUsername(v string) *JobItem {
	s.Username = &v
	return s
}

func (s *JobItem) SetWorkingDir(v string) *JobItem {
	s.WorkingDir = &v
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
	// example:
	//
	// master
	Branch *string `json:"Branch,omitempty" xml:"Branch,omitempty"`
	// example:
	//
	// code-20210111103721-85qz78ia96lu
	CodeSourceId *string `json:"CodeSourceId,omitempty" xml:"CodeSourceId,omitempty"`
	// example:
	//
	// 44da109b59f8596152987eaa8f3b2487bb72ea63
	Commit *string `json:"Commit,omitempty" xml:"Commit,omitempty"`
	// example:
	//
	// /mnt/data
	MountPath *string `json:"MountPath,omitempty" xml:"MountPath,omitempty"`
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
	// example:
	//
	// data-20210114104214-vf9lowjt3pso
	DataSourceId *string `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	// example:
	//
	// /mnt/data
	MountPath *string `json:"MountPath,omitempty" xml:"MountPath,omitempty"`
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

type JobItemUserVpc struct {
	DefaultRoute    *string   `json:"DefaultRoute,omitempty" xml:"DefaultRoute,omitempty"`
	ExtendedCidrs   []*string `json:"ExtendedCidrs,omitempty" xml:"ExtendedCidrs,omitempty" type:"Repeated"`
	SecurityGroupId *string   `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	SwitchId        *string   `json:"SwitchId,omitempty" xml:"SwitchId,omitempty"`
	VpcId           *string   `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s JobItemUserVpc) String() string {
	return tea.Prettify(s)
}

func (s JobItemUserVpc) GoString() string {
	return s.String()
}

func (s *JobItemUserVpc) SetDefaultRoute(v string) *JobItemUserVpc {
	s.DefaultRoute = &v
	return s
}

func (s *JobItemUserVpc) SetExtendedCidrs(v []*string) *JobItemUserVpc {
	s.ExtendedCidrs = v
	return s
}

func (s *JobItemUserVpc) SetSecurityGroupId(v string) *JobItemUserVpc {
	s.SecurityGroupId = &v
	return s
}

func (s *JobItemUserVpc) SetSwitchId(v string) *JobItemUserVpc {
	s.SwitchId = &v
	return s
}

func (s *JobItemUserVpc) SetVpcId(v string) *JobItemUserVpc {
	s.VpcId = &v
	return s
}

type JobSettings struct {
	AdvancedSettings       map[string]interface{} `json:"AdvancedSettings,omitempty" xml:"AdvancedSettings,omitempty"`
	AllocateAllRDMADevices *bool                  `json:"AllocateAllRDMADevices,omitempty" xml:"AllocateAllRDMADevices,omitempty"`
	// example:
	//
	// 166924
	BusinessUserId *string `json:"BusinessUserId,omitempty" xml:"BusinessUserId,omitempty"`
	// example:
	//
	// SilkFlow
	Caller *string `json:"Caller,omitempty" xml:"Caller,omitempty"`
	// example:
	//
	// false
	DisableEcsStockCheck *bool `json:"DisableEcsStockCheck,omitempty" xml:"DisableEcsStockCheck,omitempty"`
	// example:
	//
	// 535.54.03
	Driver *string `json:"Driver,omitempty" xml:"Driver,omitempty"`
	// example:
	//
	// true
	EnableCPUAffinity *bool `json:"EnableCPUAffinity,omitempty" xml:"EnableCPUAffinity,omitempty"`
	// example:
	//
	// false
	EnableErrorMonitoringInAIMaster *bool `json:"EnableErrorMonitoringInAIMaster,omitempty" xml:"EnableErrorMonitoringInAIMaster,omitempty"`
	// example:
	//
	// true
	EnableOssAppend *bool `json:"EnableOssAppend,omitempty" xml:"EnableOssAppend,omitempty"`
	// example:
	//
	// true
	EnableRDMA *bool `json:"EnableRDMA,omitempty" xml:"EnableRDMA,omitempty"`
	// example:
	//
	// true
	EnableSanityCheck *bool `json:"EnableSanityCheck,omitempty" xml:"EnableSanityCheck,omitempty"`
	// example:
	//
	// true
	EnableTideResource *bool `json:"EnableTideResource,omitempty" xml:"EnableTideResource,omitempty"`
	// example:
	//
	// --enable-log-hang-detection true
	ErrorMonitoringArgs *string `json:"ErrorMonitoringArgs,omitempty" xml:"ErrorMonitoringArgs,omitempty"`
	// example:
	//
	// 30
	JobReservedMinutes *int32 `json:"JobReservedMinutes,omitempty" xml:"JobReservedMinutes,omitempty"`
	// example:
	//
	// Always
	JobReservedPolicy *string `json:"JobReservedPolicy,omitempty" xml:"JobReservedPolicy,omitempty"`
	// example:
	//
	// AcceptQuotaOverSold
	OversoldType *string `json:"OversoldType,omitempty" xml:"OversoldType,omitempty"`
	// example:
	//
	// pid-123456
	PipelineId *string `json:"PipelineId,omitempty" xml:"PipelineId,omitempty"`
	// example:
	//
	// --sanity-check-timing=AfterJobFaultTolerant --sanity-check-timeout-ops=MarkJobFai
	SanityCheckArgs *string            `json:"SanityCheckArgs,omitempty" xml:"SanityCheckArgs,omitempty"`
	Tags            map[string]*string `json:"Tags,omitempty" xml:"Tags,omitempty"`
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

func (s *JobSettings) SetAllocateAllRDMADevices(v bool) *JobSettings {
	s.AllocateAllRDMADevices = &v
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

func (s *JobSettings) SetDisableEcsStockCheck(v bool) *JobSettings {
	s.DisableEcsStockCheck = &v
	return s
}

func (s *JobSettings) SetDriver(v string) *JobSettings {
	s.Driver = &v
	return s
}

func (s *JobSettings) SetEnableCPUAffinity(v bool) *JobSettings {
	s.EnableCPUAffinity = &v
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
	AssignNodeSpec  *AssignNodeSpec  `json:"AssignNodeSpec,omitempty" xml:"AssignNodeSpec,omitempty"`
	AutoScalingSpec *AutoScalingSpec `json:"AutoScalingSpec,omitempty" xml:"AutoScalingSpec,omitempty"`
	// example:
	//
	// ecs.c6.large
	EcsSpec      *string       `json:"EcsSpec,omitempty" xml:"EcsSpec,omitempty"`
	ExtraPodSpec *ExtraPodSpec `json:"ExtraPodSpec,omitempty" xml:"ExtraPodSpec,omitempty"`
	// example:
	//
	// registry.cn-hangzhou.aliyuncs.com/pai-dlc/tensorflow-training:1.12.2PAI-cpu-py27-ubuntu16.04
	Image       *string      `json:"Image,omitempty" xml:"Image,omitempty"`
	ImageConfig *ImageConfig `json:"ImageConfig,omitempty" xml:"ImageConfig,omitempty"`
	// Deprecated
	IsCheif         *bool             `json:"IsCheif,omitempty" xml:"IsCheif,omitempty"`
	IsChief         *bool             `json:"IsChief,omitempty" xml:"IsChief,omitempty"`
	LocalMountSpecs []*LocalMountSpec `json:"LocalMountSpecs,omitempty" xml:"LocalMountSpecs,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PodCount       *int64          `json:"PodCount,omitempty" xml:"PodCount,omitempty"`
	ResourceConfig *ResourceConfig `json:"ResourceConfig,omitempty" xml:"ResourceConfig,omitempty"`
	RestartPolicy  *string         `json:"RestartPolicy,omitempty" xml:"RestartPolicy,omitempty"`
	ServiceSpec    *ServiceSpec    `json:"ServiceSpec,omitempty" xml:"ServiceSpec,omitempty"`
	SpotSpec       *SpotSpec       `json:"SpotSpec,omitempty" xml:"SpotSpec,omitempty"`
	// example:
	//
	// Worker
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// Deprecated
	//
	// example:
	//
	// false
	UseSpotInstance *bool `json:"UseSpotInstance,omitempty" xml:"UseSpotInstance,omitempty"`
}

func (s JobSpec) String() string {
	return tea.Prettify(s)
}

func (s JobSpec) GoString() string {
	return s.String()
}

func (s *JobSpec) SetAssignNodeSpec(v *AssignNodeSpec) *JobSpec {
	s.AssignNodeSpec = v
	return s
}

func (s *JobSpec) SetAutoScalingSpec(v *AutoScalingSpec) *JobSpec {
	s.AutoScalingSpec = v
	return s
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

func (s *JobSpec) SetIsCheif(v bool) *JobSpec {
	s.IsCheif = &v
	return s
}

func (s *JobSpec) SetIsChief(v bool) *JobSpec {
	s.IsChief = &v
	return s
}

func (s *JobSpec) SetLocalMountSpecs(v []*LocalMountSpec) *JobSpec {
	s.LocalMountSpecs = v
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

func (s *JobSpec) SetRestartPolicy(v string) *JobSpec {
	s.RestartPolicy = &v
	return s
}

func (s *JobSpec) SetServiceSpec(v *ServiceSpec) *JobSpec {
	s.ServiceSpec = v
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

type Lifecycle struct {
	PostStart *LifecyclePostStart `json:"PostStart,omitempty" xml:"PostStart,omitempty" type:"Struct"`
	PreStop   *LifecyclePreStop   `json:"PreStop,omitempty" xml:"PreStop,omitempty" type:"Struct"`
}

func (s Lifecycle) String() string {
	return tea.Prettify(s)
}

func (s Lifecycle) GoString() string {
	return s.String()
}

func (s *Lifecycle) SetPostStart(v *LifecyclePostStart) *Lifecycle {
	s.PostStart = v
	return s
}

func (s *Lifecycle) SetPreStop(v *LifecyclePreStop) *Lifecycle {
	s.PreStop = v
	return s
}

type LifecyclePostStart struct {
	Exec *LifecyclePostStartExec `json:"Exec,omitempty" xml:"Exec,omitempty" type:"Struct"`
}

func (s LifecyclePostStart) String() string {
	return tea.Prettify(s)
}

func (s LifecyclePostStart) GoString() string {
	return s.String()
}

func (s *LifecyclePostStart) SetExec(v *LifecyclePostStartExec) *LifecyclePostStart {
	s.Exec = v
	return s
}

type LifecyclePostStartExec struct {
	Command []*string `json:"Command,omitempty" xml:"Command,omitempty" type:"Repeated"`
}

func (s LifecyclePostStartExec) String() string {
	return tea.Prettify(s)
}

func (s LifecyclePostStartExec) GoString() string {
	return s.String()
}

func (s *LifecyclePostStartExec) SetCommand(v []*string) *LifecyclePostStartExec {
	s.Command = v
	return s
}

type LifecyclePreStop struct {
	Exec *LifecyclePreStopExec `json:"Exec,omitempty" xml:"Exec,omitempty" type:"Struct"`
}

func (s LifecyclePreStop) String() string {
	return tea.Prettify(s)
}

func (s LifecyclePreStop) GoString() string {
	return s.String()
}

func (s *LifecyclePreStop) SetExec(v *LifecyclePreStopExec) *LifecyclePreStop {
	s.Exec = v
	return s
}

type LifecyclePreStopExec struct {
	Command []*string `json:"Command,omitempty" xml:"Command,omitempty" type:"Repeated"`
}

func (s LifecyclePreStopExec) String() string {
	return tea.Prettify(s)
}

func (s LifecyclePreStopExec) GoString() string {
	return s.String()
}

func (s *LifecyclePreStopExec) SetCommand(v []*string) *LifecyclePreStopExec {
	s.Command = v
	return s
}

type LocalMountSpec struct {
	LocalPath *string `json:"LocalPath,omitempty" xml:"LocalPath,omitempty"`
	MountMode *string `json:"MountMode,omitempty" xml:"MountMode,omitempty"`
	MountPath *string `json:"MountPath,omitempty" xml:"MountPath,omitempty"`
}

func (s LocalMountSpec) String() string {
	return tea.Prettify(s)
}

func (s LocalMountSpec) GoString() string {
	return s.String()
}

func (s *LocalMountSpec) SetLocalPath(v string) *LocalMountSpec {
	s.LocalPath = &v
	return s
}

func (s *LocalMountSpec) SetMountMode(v string) *LocalMountSpec {
	s.MountMode = &v
	return s
}

func (s *LocalMountSpec) SetMountPath(v string) *LocalMountSpec {
	s.MountPath = &v
	return s
}

type LogInfo struct {
	Content     *string `json:"Content,omitempty" xml:"Content,omitempty"`
	Id          *string `json:"Id,omitempty" xml:"Id,omitempty"`
	IsTruncated *bool   `json:"IsTruncated,omitempty" xml:"IsTruncated,omitempty"`
	PodId       *string `json:"PodId,omitempty" xml:"PodId,omitempty"`
	PodUid      *string `json:"PodUid,omitempty" xml:"PodUid,omitempty"`
	// example:
	//
	// stderr, stdout
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	Time   *string `json:"Time,omitempty" xml:"Time,omitempty"`
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

func (s *LogInfo) SetIsTruncated(v bool) *LogInfo {
	s.IsTruncated = &v
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
	// example:
	//
	// ken_12345
	MemberId *string `json:"MemberId,omitempty" xml:"MemberId,omitempty"`
	// example:
	//
	// WorkspaceAdmin
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
	// example:
	//
	// 1616987726587
	Time *int64 `json:"Time,omitempty" xml:"Time,omitempty"`
	// example:
	//
	// 23.45
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
	Metrics []*Metric `json:"Metrics,omitempty" xml:"Metrics,omitempty" type:"Repeated"`
	// example:
	//
	// asi_xxx
	NodeName *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
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
	// example:
	//
	// 2021-01-12T14:36:01Z
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// example:
	//
	// 2021-01-12T15:36:05Z
	GmtFinishTime *string `json:"GmtFinishTime,omitempty" xml:"GmtFinishTime,omitempty"`
	// example:
	//
	// 2021-01-12T14:36:05Z
	GmtStartTime *string    `json:"GmtStartTime,omitempty" xml:"GmtStartTime,omitempty"`
	HistoryPods  []*PodItem `json:"HistoryPods,omitempty" xml:"HistoryPods,omitempty" type:"Repeated"`
	// example:
	//
	// 10.0.1.2
	Ip       *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	NodeName *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	// example:
	//
	// dlc-20210126170216-mtl37ge7gkvdz-worker-0
	PodId *string `json:"PodId,omitempty" xml:"PodId,omitempty"`
	// example:
	//
	// fe846462-af2c-4521-bd6f-96787a57591d
	PodUid *string `json:"PodUid,omitempty" xml:"PodUid,omitempty"`
	// example:
	//
	// Stopped
	Status    *string `json:"Status,omitempty" xml:"Status,omitempty"`
	SubStatus *string `json:"SubStatus,omitempty" xml:"SubStatus,omitempty"`
	// example:
	//
	// Worker
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
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

func (s *PodItem) SetNodeName(v string) *PodItem {
	s.NodeName = &v
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

func (s *PodItem) SetSubStatus(v string) *PodItem {
	s.SubStatus = &v
	return s
}

func (s *PodItem) SetType(v string) *PodItem {
	s.Type = &v
	return s
}

type PodMetric struct {
	Metrics []*Metric `json:"Metrics,omitempty" xml:"Metrics,omitempty" type:"Repeated"`
	// example:
	//
	// dlc-20210329110128-746bf7cl47pr8-worker-0
	PodId *string `json:"PodId,omitempty" xml:"PodId,omitempty"`
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
	ClusterId   *string      `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	ClusterName *string      `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	QuotaConfig *QuotaConfig `json:"QuotaConfig,omitempty" xml:"QuotaConfig,omitempty"`
	// example:
	//
	// quotamtl37ge7gkvdz
	QuotaId *string `json:"QuotaId,omitempty" xml:"QuotaId,omitempty"`
	// example:
	//
	// dlc-quota
	QuotaName *string `json:"QuotaName,omitempty" xml:"QuotaName,omitempty"`
	// example:
	//
	// asiquota
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
	// example:
	//
	// 2
	CPU *string `json:"CPU,omitempty" xml:"CPU,omitempty"`
	// example:
	//
	// 5
	GPU        *string      `json:"GPU,omitempty" xml:"GPU,omitempty"`
	GPUDetails []*GPUDetail `json:"GPUDetails,omitempty" xml:"GPUDetails,omitempty" type:"Repeated"`
	// example:
	//
	// Tesla-V100
	GPUType *string `json:"GPUType,omitempty" xml:"GPUType,omitempty"`
	// example:
	//
	// nvidia.com/gpu
	GPUTypeFullName *string `json:"GPUTypeFullName,omitempty" xml:"GPUTypeFullName,omitempty"`
	// example:
	//
	// 10Gi
	Memory *string `json:"Memory,omitempty" xml:"Memory,omitempty"`
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
	// example:
	//
	// 10
	CPU *string `json:"CPU,omitempty" xml:"CPU,omitempty"`
	// example:
	//
	// 3
	GPU *string `json:"GPU,omitempty" xml:"GPU,omitempty"`
	// example:
	//
	// Tesla-V100-16G
	GPUType *string `json:"GPUType,omitempty" xml:"GPUType,omitempty"`
	// example:
	//
	// 10Gi
	Memory *string `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// example:
	//
	// 5Gi
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
	// example:
	//
	// 10
	CPU *string `json:"CPU,omitempty" xml:"CPU,omitempty"`
	// example:
	//
	// 8
	GPU *string `json:"GPU,omitempty" xml:"GPU,omitempty"`
	// example:
	//
	// 1024（单位GB）
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
	// example:
	//
	// 1
	CheckNumber *int32 `json:"CheckNumber,omitempty" xml:"CheckNumber,omitempty"`
	// example:
	//
	// ”2023-11-30T16:47:30.378817+08:00"
	FinishedAt *string `json:"FinishedAt,omitempty" xml:"FinishedAt,omitempty"`
	Message    *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// CheckInit
	Phase *string `json:"Phase,omitempty" xml:"Phase,omitempty"`
	// example:
	//
	// ”2023-11-30T16:47:30.343005+08:00“
	StartedAt *string `json:"StartedAt,omitempty" xml:"StartedAt,omitempty"`
	// example:
	//
	// Succeeded
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
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

type SeccompProfile struct {
	// example:
	//
	// my-profiles/profile-allow.json
	LocalhostProfile *string `json:"LocalhostProfile,omitempty" xml:"LocalhostProfile,omitempty"`
	// example:
	//
	// Unconfined
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s SeccompProfile) String() string {
	return tea.Prettify(s)
}

func (s SeccompProfile) GoString() string {
	return s.String()
}

func (s *SeccompProfile) SetLocalhostProfile(v string) *SeccompProfile {
	s.LocalhostProfile = &v
	return s
}

func (s *SeccompProfile) SetType(v string) *SeccompProfile {
	s.Type = &v
	return s
}

type SecurityContext struct {
	Capabilities *SecurityContextCapabilities `json:"Capabilities,omitempty" xml:"Capabilities,omitempty"`
	Privileged   *bool                        `json:"Privileged,omitempty" xml:"Privileged,omitempty"`
	// example:
	//
	// 1000
	RunAsGroup *int64 `json:"RunAsGroup,omitempty" xml:"RunAsGroup,omitempty"`
	// example:
	//
	// 1000
	RunAsUser      *int64          `json:"RunAsUser,omitempty" xml:"RunAsUser,omitempty"`
	SeccompProfile *SeccompProfile `json:"SeccompProfile,omitempty" xml:"SeccompProfile,omitempty"`
}

func (s SecurityContext) String() string {
	return tea.Prettify(s)
}

func (s SecurityContext) GoString() string {
	return s.String()
}

func (s *SecurityContext) SetCapabilities(v *SecurityContextCapabilities) *SecurityContext {
	s.Capabilities = v
	return s
}

func (s *SecurityContext) SetPrivileged(v bool) *SecurityContext {
	s.Privileged = &v
	return s
}

func (s *SecurityContext) SetRunAsGroup(v int64) *SecurityContext {
	s.RunAsGroup = &v
	return s
}

func (s *SecurityContext) SetRunAsUser(v int64) *SecurityContext {
	s.RunAsUser = &v
	return s
}

func (s *SecurityContext) SetSeccompProfile(v *SeccompProfile) *SecurityContext {
	s.SeccompProfile = v
	return s
}

type SecurityContextCapabilities struct {
	Add  []*string `json:"Add,omitempty" xml:"Add,omitempty" type:"Repeated"`
	Drop []*string `json:"Drop,omitempty" xml:"Drop,omitempty" type:"Repeated"`
}

func (s SecurityContextCapabilities) String() string {
	return tea.Prettify(s)
}

func (s SecurityContextCapabilities) GoString() string {
	return s.String()
}

func (s *SecurityContextCapabilities) SetAdd(v []*string) *SecurityContextCapabilities {
	s.Add = v
	return s
}

func (s *SecurityContextCapabilities) SetDrop(v []*string) *SecurityContextCapabilities {
	s.Drop = v
	return s
}

type ServiceSpec struct {
	DefaultPort *int32  `json:"DefaultPort,omitempty" xml:"DefaultPort,omitempty"`
	ExtraPorts  *int32  `json:"ExtraPorts,omitempty" xml:"ExtraPorts,omitempty"`
	ServiceMode *string `json:"ServiceMode,omitempty" xml:"ServiceMode,omitempty"`
}

func (s ServiceSpec) String() string {
	return tea.Prettify(s)
}

func (s ServiceSpec) GoString() string {
	return s.String()
}

func (s *ServiceSpec) SetDefaultPort(v int32) *ServiceSpec {
	s.DefaultPort = &v
	return s
}

func (s *ServiceSpec) SetExtraPorts(v int32) *ServiceSpec {
	s.ExtraPorts = &v
	return s
}

func (s *ServiceSpec) SetServiceMode(v string) *ServiceSpec {
	s.ServiceMode = &v
	return s
}

type SmartCache struct {
	// example:
	//
	// 10
	CacheWorkerNum *int64 `json:"CacheWorkerNum,omitempty" xml:"CacheWorkerNum,omitempty"`
	// example:
	//
	// 100
	CacheWorkerSize *int64 `json:"CacheWorkerSize,omitempty" xml:"CacheWorkerSize,omitempty"`
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// test
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// example:
	//
	// 123456
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// example:
	//
	// oss-cn-beijing-internal.aliyuncs.com
	Endpoint *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	// example:
	//
	// 1ca404****
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// example:
	//
	// 2021-01-12T14:36:01Z
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// example:
	//
	// 2021-01-12T23:36:01Z
	GmtModifyTime *string `json:"GmtModifyTime,omitempty" xml:"GmtModifyTime,omitempty"`
	// example:
	//
	// /root/data/
	MountPath *string `json:"MountPath,omitempty" xml:"MountPath,omitempty"`
	// example:
	//
	// {"num_threads": 32}
	Options *string `json:"Options,omitempty" xml:"Options,omitempty"`
	// example:
	//
	// oss://buc/path/to/dir
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// example:
	//
	// smartcache-20210114104214-vf9lowjt3pso
	SmartCacheId *string `json:"SmartCacheId,omitempty" xml:"SmartCacheId,omitempty"`
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// oss
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// 189xxx
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
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
	SpotPriceLimit    *float32 `json:"SpotPriceLimit,omitempty" xml:"SpotPriceLimit,omitempty"`
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

func (s *SpotSpec) SetSpotPriceLimit(v float32) *SpotSpec {
	s.SpotPriceLimit = &v
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
	Accessibility *string `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	Cpu           *int64  `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// example:
	//
	// datasource-test
	DataSourceId   *string `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	DataSourceType *string `json:"DataSourceType,omitempty" xml:"DataSourceType,omitempty"`
	// example:
	//
	// test
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// example:
	//
	// 1234567
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// example:
	//
	// 2021-01-12T14:35:00Z
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// example:
	//
	// 2021-01-12T14:36:00Z
	GmtFinishTime *string `json:"GmtFinishTime,omitempty" xml:"GmtFinishTime,omitempty"`
	// example:
	//
	// 2021-01-12T14:36:00Z
	GmtModifyTime *string `json:"GmtModifyTime,omitempty" xml:"GmtModifyTime,omitempty"`
	// example:
	//
	// dlc-20210114104214-vf9lowjt3pso
	JobId                 *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	MaxRunningTimeMinutes *int64  `json:"MaxRunningTimeMinutes,omitempty" xml:"MaxRunningTimeMinutes,omitempty"`
	Memory                *int64  `json:"Memory,omitempty" xml:"Memory,omitempty"`
	Options               *string `json:"Options,omitempty" xml:"Options,omitempty"`
	Priority              *string `json:"Priority,omitempty" xml:"Priority,omitempty"`
	QuotaId               *string `json:"QuotaId,omitempty" xml:"QuotaId,omitempty"`
	QuotaName             *string `json:"QuotaName,omitempty" xml:"QuotaName,omitempty"`
	// example:
	//
	// Delete by user
	ReasonCode *string `json:"ReasonCode,omitempty" xml:"ReasonCode,omitempty"`
	// example:
	//
	// Tensorboard is deleted
	ReasonMessage *string `json:"ReasonMessage,omitempty" xml:"ReasonMessage,omitempty"`
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// /root/data
	SummaryPath            *string                      `json:"SummaryPath,omitempty" xml:"SummaryPath,omitempty"`
	SummaryRelativePath    *string                      `json:"SummaryRelativePath,omitempty" xml:"SummaryRelativePath,omitempty"`
	TensorboardDataSources []*TensorboardDataSourceSpec `json:"TensorboardDataSources,omitempty" xml:"TensorboardDataSources,omitempty" type:"Repeated"`
	// example:
	//
	// tensorboard-xxx
	TensorboardId   *string          `json:"TensorboardId,omitempty" xml:"TensorboardId,omitempty"`
	TensorboardSpec *TensorboardSpec `json:"TensorboardSpec,omitempty" xml:"TensorboardSpec,omitempty"`
	// example:
	//
	// http://xxxxxx
	TensorboardUrl *string `json:"TensorboardUrl,omitempty" xml:"TensorboardUrl,omitempty"`
	Token          *string `json:"Token,omitempty" xml:"Token,omitempty"`
	// example:
	//
	// lycxxxxx
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// example:
	//
	// tensorboard.pai
	Username    *string `json:"Username,omitempty" xml:"Username,omitempty"`
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s Tensorboard) String() string {
	return tea.Prettify(s)
}

func (s Tensorboard) GoString() string {
	return s.String()
}

func (s *Tensorboard) SetAccessibility(v string) *Tensorboard {
	s.Accessibility = &v
	return s
}

func (s *Tensorboard) SetCpu(v int64) *Tensorboard {
	s.Cpu = &v
	return s
}

func (s *Tensorboard) SetDataSourceId(v string) *Tensorboard {
	s.DataSourceId = &v
	return s
}

func (s *Tensorboard) SetDataSourceType(v string) *Tensorboard {
	s.DataSourceType = &v
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

func (s *Tensorboard) SetGmtFinishTime(v string) *Tensorboard {
	s.GmtFinishTime = &v
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

func (s *Tensorboard) SetMaxRunningTimeMinutes(v int64) *Tensorboard {
	s.MaxRunningTimeMinutes = &v
	return s
}

func (s *Tensorboard) SetMemory(v int64) *Tensorboard {
	s.Memory = &v
	return s
}

func (s *Tensorboard) SetOptions(v string) *Tensorboard {
	s.Options = &v
	return s
}

func (s *Tensorboard) SetPriority(v string) *Tensorboard {
	s.Priority = &v
	return s
}

func (s *Tensorboard) SetQuotaId(v string) *Tensorboard {
	s.QuotaId = &v
	return s
}

func (s *Tensorboard) SetQuotaName(v string) *Tensorboard {
	s.QuotaName = &v
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

func (s *Tensorboard) SetSummaryRelativePath(v string) *Tensorboard {
	s.SummaryRelativePath = &v
	return s
}

func (s *Tensorboard) SetTensorboardDataSources(v []*TensorboardDataSourceSpec) *Tensorboard {
	s.TensorboardDataSources = v
	return s
}

func (s *Tensorboard) SetTensorboardId(v string) *Tensorboard {
	s.TensorboardId = &v
	return s
}

func (s *Tensorboard) SetTensorboardSpec(v *TensorboardSpec) *Tensorboard {
	s.TensorboardSpec = v
	return s
}

func (s *Tensorboard) SetTensorboardUrl(v string) *Tensorboard {
	s.TensorboardUrl = &v
	return s
}

func (s *Tensorboard) SetToken(v string) *Tensorboard {
	s.Token = &v
	return s
}

func (s *Tensorboard) SetUserId(v string) *Tensorboard {
	s.UserId = &v
	return s
}

func (s *Tensorboard) SetUsername(v string) *Tensorboard {
	s.Username = &v
	return s
}

func (s *Tensorboard) SetWorkspaceId(v string) *Tensorboard {
	s.WorkspaceId = &v
	return s
}

type TensorboardDataSourceSpec struct {
	// example:
	//
	// OSS
	DataSourceType *string `json:"DataSourceType,omitempty" xml:"DataSourceType,omitempty"`
	// example:
	//
	// dlcJobName
	DirectoryName *string `json:"DirectoryName,omitempty" xml:"DirectoryName,omitempty"`
	// example:
	//
	// oss://xxxxx/tensorboard/run1
	FullSummaryPath *string `json:"FullSummaryPath,omitempty" xml:"FullSummaryPath,omitempty"`
	// example:
	//
	// d-vf2fdhxxxxxx
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// dlcJobName
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// datasource
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// example:
	//
	// /tensorboard/run1
	SummaryPath *string `json:"SummaryPath,omitempty" xml:"SummaryPath,omitempty"`
	// example:
	//
	// oss://.oss-cn-shanghai-finance-1.aliyuncs.com/
	Uri *string `json:"Uri,omitempty" xml:"Uri,omitempty"`
}

func (s TensorboardDataSourceSpec) String() string {
	return tea.Prettify(s)
}

func (s TensorboardDataSourceSpec) GoString() string {
	return s.String()
}

func (s *TensorboardDataSourceSpec) SetDataSourceType(v string) *TensorboardDataSourceSpec {
	s.DataSourceType = &v
	return s
}

func (s *TensorboardDataSourceSpec) SetDirectoryName(v string) *TensorboardDataSourceSpec {
	s.DirectoryName = &v
	return s
}

func (s *TensorboardDataSourceSpec) SetFullSummaryPath(v string) *TensorboardDataSourceSpec {
	s.FullSummaryPath = &v
	return s
}

func (s *TensorboardDataSourceSpec) SetId(v string) *TensorboardDataSourceSpec {
	s.Id = &v
	return s
}

func (s *TensorboardDataSourceSpec) SetName(v string) *TensorboardDataSourceSpec {
	s.Name = &v
	return s
}

func (s *TensorboardDataSourceSpec) SetSourceType(v string) *TensorboardDataSourceSpec {
	s.SourceType = &v
	return s
}

func (s *TensorboardDataSourceSpec) SetSummaryPath(v string) *TensorboardDataSourceSpec {
	s.SummaryPath = &v
	return s
}

func (s *TensorboardDataSourceSpec) SetUri(v string) *TensorboardDataSourceSpec {
	s.Uri = &v
	return s
}

type TensorboardSpec struct {
	// example:
	//
	// ecs.g6.large
	EcsType *string `json:"EcsType,omitempty" xml:"EcsType,omitempty"`
	// example:
	//
	// sg-xxxxx
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// example:
	//
	// vsw-xxxx
	SwitchId *string `json:"SwitchId,omitempty" xml:"SwitchId,omitempty"`
	// example:
	//
	// vpc-xxxx
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s TensorboardSpec) String() string {
	return tea.Prettify(s)
}

func (s TensorboardSpec) GoString() string {
	return s.String()
}

func (s *TensorboardSpec) SetEcsType(v string) *TensorboardSpec {
	s.EcsType = &v
	return s
}

func (s *TensorboardSpec) SetSecurityGroupId(v string) *TensorboardSpec {
	s.SecurityGroupId = &v
	return s
}

func (s *TensorboardSpec) SetSwitchId(v string) *TensorboardSpec {
	s.SwitchId = &v
	return s
}

func (s *TensorboardSpec) SetVpcId(v string) *TensorboardSpec {
	s.VpcId = &v
	return s
}

type Workspace struct {
	// example:
	//
	// ken
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// example:
	//
	// 2021-01-12T14:36:01Z
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// example:
	//
	// 2021-01-12T14:36:01Z
	GmtModifyTime   *string    `json:"GmtModifyTime,omitempty" xml:"GmtModifyTime,omitempty"`
	Members         []*Member  `json:"Members,omitempty" xml:"Members,omitempty" type:"Repeated"`
	Quotas          []*Quota   `json:"Quotas,omitempty" xml:"Quotas,omitempty" type:"Repeated"`
	TotalResources  *Resources `json:"TotalResources,omitempty" xml:"TotalResources,omitempty"`
	WorkspaceAdmins []*Member  `json:"WorkspaceAdmins,omitempty" xml:"WorkspaceAdmins,omitempty" type:"Repeated"`
	// example:
	//
	// ws-20210126170216-mtl37ge7gkvdz
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
	// example:
	//
	// dlc-workspace
	WorkspaceName *string `json:"WorkspaceName,omitempty" xml:"WorkspaceName,omitempty"`
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
	// The job visibility. Valid values:
	//
	// 	- PUBLIC: The job is visible to all members in the workspace.
	//
	// 	- PRIVATE: The job is visible only to you and the administrator of the workspace.
	//
	// example:
	//
	// PRIVATE
	Accessibility *string `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	// The code source of the job. Before the node of the job runs, DLC automatically downloads the configured code from the code source and mounts the code to the local path of the container.
	CodeSource *CreateJobRequestCodeSource `json:"CodeSource,omitempty" xml:"CodeSource,omitempty" type:"Struct"`
	// The access 
	CredentialConfig *CredentialConfig `json:"CredentialConfig,omitempty" xml:"CredentialConfig,omitempty"`
	// The data sources for job running.
	DataSources []*CreateJobRequestDataSources `json:"DataSources,omitempty" xml:"DataSources,omitempty" type:"Repeated"`
	// This parameter is not supported.
	//
	// example:
	//
	// “”
	DebuggerConfigContent *string `json:"DebuggerConfigContent,omitempty" xml:"DebuggerConfigContent,omitempty"`
	// The job name. The name must be in the following format:
	//
	// 	- The name must be 1 to 256 characters in length.
	//
	// 	- The name can contain digits, letters, underscores (_), periods (.), and hyphens (-).
	//
	// This parameter is required.
	//
	// example:
	//
	// tf-mnist-test
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// This parameter is not supported.
	ElasticSpec *JobElasticSpec `json:"ElasticSpec,omitempty" xml:"ElasticSpec,omitempty"`
	// The environment variables.
	Envs map[string]*string `json:"Envs,omitempty" xml:"Envs,omitempty"`
	// The maximum running duration of the job. Unit: minutes.
	//
	// example:
	//
	// 1024
	JobMaxRunningTimeMinutes *int64 `json:"JobMaxRunningTimeMinutes,omitempty" xml:"JobMaxRunningTimeMinutes,omitempty"`
	// The configurations for job running, such as the image address, startup command, node resource declaration, and number of replicas.****
	//
	// A DLC job consists of different types of nodes. If nodes of the same type have exactly the same configuration, the configuration is called JobSpec. **JobSpecs*	- specifies the configurations of all types of nodes. The value is of the array type.
	//
	// This parameter is required.
	JobSpecs []*JobSpec `json:"JobSpecs,omitempty" xml:"JobSpecs,omitempty" type:"Repeated"`
	// The job type. The value is case-sensitive. The following job types are supported:
	//
	// 	- TFJob
	//
	// 	- PyTorchJob
	//
	// 	- MPIJob
	//
	// 	- XGBoostJob
	//
	// 	- OneFlowJob
	//
	// 	- ElasticBatchJob
	//
	// 	- SlurmJob
	//
	// 	- RayJob
	//
	// Valid values for each job type:
	//
	// 	- OneFlowJob: OneFlow.
	//
	// 	- PyTorchJob: PyTorch.
	//
	// 	- SlurmJob: Slurm.
	//
	// 	- XGBoostJob: XGBoost.
	//
	// 	- ElasticBatchJob: ElasticBatch.
	//
	// 	- MPIJob: MPIJob.
	//
	// 	- TFJob: Tensorflow.
	//
	// 	- RayJob: Ray.
	//
	// This parameter is required.
	//
	// example:
	//
	// TFJob
	JobType *string `json:"JobType,omitempty" xml:"JobType,omitempty"`
	// The additional configuration of the job. You can use this parameter to adjust the behavior of the attached data source. For example, if the attached data source of the job is of the OSS type, you can use this parameter to add the following configurations to override the default parameters of JindoFS: `fs.oss.download.thread.concurrency=4,fs.oss.download.queue.size=16`.
	//
	// example:
	//
	// key1=value1,key2=value2
	Options *string `json:"Options,omitempty" xml:"Options,omitempty"`
	// The priority of the job. Default value: 1. Valid values: 1 to 9.
	//
	// 	- 1: the lowest priority.
	//
	// 	- 9: the highest priority.
	//
	// example:
	//
	// 8
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The ID of the resource group. This parameter is optional.
	//
	// 	- If you leave this parameter empty, the job is submitted to a public resource group.
	//
	// 	- If a resource quota is bound to the current workspace, you can specify the resource quota ID. For more information about how to query the resource quota ID, see [Manage resource quotas](https://help.aliyun.com/document_detail/2651299.html).
	//
	// example:
	//
	// rs-xxx
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The additional parameter configurations of the job.
	Settings *JobSettings `json:"Settings,omitempty" xml:"Settings,omitempty"`
	// The policy that is used to check whether a distributed multi-node job is successful. Only TensorFlow distributed multi-node jobs are supported.
	//
	// 	- ChiefWorker: If you use this policy, the job is considered successful when the pod on the chief node completes operations.
	//
	// 	- AllWorkers (default): If you use this policy, the job is considered successful when all worker nodes complete operations.
	//
	// example:
	//
	// AllWorkers
	SuccessPolicy *string `json:"SuccessPolicy,omitempty" xml:"SuccessPolicy,omitempty"`
	// The folder in which the third-party Python library file requirements.txt is stored. Before the startup command specified by the UserCommand parameter is run on each node, DLC fetches the requirements.txt file from the folder and runs `pip install -r` to install the required package and library.
	//
	// example:
	//
	// /root/code/
	ThirdpartyLibDir *string `json:"ThirdpartyLibDir,omitempty" xml:"ThirdpartyLibDir,omitempty"`
	// The third-party Python libraries to be installed.
	ThirdpartyLibs []*string `json:"ThirdpartyLibs,omitempty" xml:"ThirdpartyLibs,omitempty" type:"Repeated"`
	// The startup command for all nodes of the job.
	//
	// This parameter is required.
	//
	// example:
	//
	// python /root/code/mnist.py
	UserCommand *string `json:"UserCommand,omitempty" xml:"UserCommand,omitempty"`
	// The VPC settings.
	UserVpc *CreateJobRequestUserVpc `json:"UserVpc,omitempty" xml:"UserVpc,omitempty" type:"Struct"`
	// The workspace ID.
	//
	// example:
	//
	// ws-20210126170216-xxxxxxx
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s CreateJobRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateJobRequest) GoString() string {
	return s.String()
}

func (s *CreateJobRequest) SetAccessibility(v string) *CreateJobRequest {
	s.Accessibility = &v
	return s
}

func (s *CreateJobRequest) SetCodeSource(v *CreateJobRequestCodeSource) *CreateJobRequest {
	s.CodeSource = v
	return s
}

func (s *CreateJobRequest) SetCredentialConfig(v *CredentialConfig) *CreateJobRequest {
	s.CredentialConfig = v
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
	// The branch of the referenced code repository. By default, the branch configured in the code source is used. This parameter is optional.
	//
	// example:
	//
	// master
	Branch *string `json:"Branch,omitempty" xml:"Branch,omitempty"`
	// The ID of the code source.
	//
	// example:
	//
	// code-20210111103721-xxxxxxx
	CodeSourceId *string `json:"CodeSourceId,omitempty" xml:"CodeSourceId,omitempty"`
	// The commit ID of the code to be downloaded. By default, the commit ID configured in the code source is used. This parameter is optional.
	//
	// example:
	//
	// 44da109b5******
	Commit *string `json:"Commit,omitempty" xml:"Commit,omitempty"`
	// The path to which the job is mounted. By default, the mount path configured in the data source is used. This parameter is optional.
	//
	// example:
	//
	// /root/data
	MountPath *string `json:"MountPath,omitempty" xml:"MountPath,omitempty"`
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
	// The data source ID.
	//
	// example:
	//
	// d-cn9dl*******
	DataSourceId      *string `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	DataSourceVersion *string `json:"DataSourceVersion,omitempty" xml:"DataSourceVersion,omitempty"`
	MountAccess       *string `json:"MountAccess,omitempty" xml:"MountAccess,omitempty"`
	// The path to which the job is mounted. By default, the mount path in the data source configuration is used. This parameter is optional.
	//
	// example:
	//
	// /root/data
	MountPath *string `json:"MountPath,omitempty" xml:"MountPath,omitempty"`
	// The mount attribute of the custom dataset. Set the value to OSS.
	//
	// example:
	//
	// {
	//
	//   "fs.oss.download.thread.concurrency": "10",
	//
	//   "fs.oss.upload.thread.concurrency": "10",
	//
	//   "fs.jindo.args": "-oattr_timeout=3 -oentry_timeout=0 -onegative_timeout=0 -oauto_cache -ono_symlink"
	//
	// }
	Options *string `json:"Options,omitempty" xml:"Options,omitempty"`
	// The data source path.
	//
	// example:
	//
	// oss://bucket.oss-cn-hangzhou-internal.aliyuncs.com/path/
	Uri *string `json:"Uri,omitempty" xml:"Uri,omitempty"`
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

func (s *CreateJobRequestDataSources) SetDataSourceVersion(v string) *CreateJobRequestDataSources {
	s.DataSourceVersion = &v
	return s
}

func (s *CreateJobRequestDataSources) SetMountAccess(v string) *CreateJobRequestDataSources {
	s.MountAccess = &v
	return s
}

func (s *CreateJobRequestDataSources) SetMountPath(v string) *CreateJobRequestDataSources {
	s.MountPath = &v
	return s
}

func (s *CreateJobRequestDataSources) SetOptions(v string) *CreateJobRequestDataSources {
	s.Options = &v
	return s
}

func (s *CreateJobRequestDataSources) SetUri(v string) *CreateJobRequestDataSources {
	s.Uri = &v
	return s
}

type CreateJobRequestUserVpc struct {
	// The default route. Default value: false. Valid values:
	//
	// 	- eth0: The default network interface is used to access the Internet through the public gateway.
	//
	// 	- eth1: The user\\"s Elastic Network Interface is used to access the Internet through the private gateway. For more information about the configuration method, see [Enable Internet access for a DSW instance by using a private Internet NAT gateway](https://help.aliyun.com/document_detail/2525343.html).
	//
	// example:
	//
	// eth0
	DefaultRoute *string `json:"DefaultRoute,omitempty" xml:"DefaultRoute,omitempty"`
	// The extended CIDR block.
	//
	// 	- If you leave the SwitchId and ExtendedCIDRs parameters empty, the system automatically obtains all CIDR blocks in a VPC.
	//
	// 	- If you configure the SwitchId and ExtendedCIDRs parameters, we recommend that you specify all CIDR blocks in a VPC.
	ExtendedCIDRs []*string `json:"ExtendedCIDRs,omitempty" xml:"ExtendedCIDRs,omitempty" type:"Repeated"`
	// The ID of the security group.
	//
	// example:
	//
	// sg-abcdef****
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The vSwitch ID. This parameter is optional.
	//
	// 	- If you leave this parameter empty, the system automatically selects a vSwitch based on the inventory status.
	//
	// 	- You can also specify a vSwitch ID.
	//
	// example:
	//
	// vs-abcdef****
	SwitchId *string `json:"SwitchId,omitempty" xml:"SwitchId,omitempty"`
	// The VPC ID.
	//
	// example:
	//
	// vpc-abcdef****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
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
	// The job ID.
	//
	// example:
	//
	// dlc7*******
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The request ID used to troubleshoot issues.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-xxxxxxx
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
	// The visibility of the job. Valid values:
	//
	// 	- PUBLIC: The configuration is public in the workspace.
	//
	// 	- PRIVATE: The configuration is visible only to you and the administrator of the workspace.
	//
	// example:
	//
	// PRIVATE
	Accessibility *string `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	// The number of vCPU cores.
	//
	// example:
	//
	// 1
	Cpu *int64 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// The dataset ID.
	//
	// <props="china">Call [ListDatasets](https://help.aliyun.com/document_detail/457222.html) to get the dataset ID.
	//
	// example:
	//
	// d-xxxxxxxx
	DataSourceId *string `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	// The dataset type. Valid values:
	//
	// 	- OSS
	//
	// 	- NAS
	//
	// example:
	//
	// OSS
	DataSourceType *string `json:"DataSourceType,omitempty" xml:"DataSourceType,omitempty"`
	// The configurations of the data source.
	DataSources []*DataSourceItem `json:"DataSources,omitempty" xml:"DataSources,omitempty" type:"Repeated"`
	// The TensorBoard name
	//
	// example:
	//
	// tensorboard
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The job ID. Call [ListJobs](https://help.aliyun.com/document_detail/459676.html) to get the job ID.
	//
	// example:
	//
	// dlc-20210126170216-mtl37ge7gkvdz
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The maximum running duration. Unit: minutes.
	//
	// example:
	//
	// 240
	MaxRunningTimeMinutes *int64 `json:"MaxRunningTimeMinutes,omitempty" xml:"MaxRunningTimeMinutes,omitempty"`
	// The memory size. Unit: GB.
	//
	// example:
	//
	// 1000
	Memory *int64 `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// The extended fields of the dataset are in the JSON format. MountPath: the path to mount the dataset.
	//
	// example:
	//
	// {"mountpath":"/root/data/"}
	Options *string `json:"Options,omitempty" xml:"Options,omitempty"`
	// The priority of the job. Default value: 1. Valid values: 1 to 9.
	//
	// 	- 1 is the lowest priority.
	//
	// 	- 9 is the highest priority.
	//
	// example:
	//
	// 1
	Priority *string `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The resource quota ID. This parameter is required when you create a TensorBoard job by using a resource quota. <props="china">Call [ListQuotas](https://help.aliyun.com/document_detail/2628071.html) to get the quota ID.
	//
	// This feature is currently limited to whitelisted users. If you need to use this feature, contact us.
	//
	// example:
	//
	// quota12345
	QuotaId *string `json:"QuotaId,omitempty" xml:"QuotaId,omitempty"`
	// The source ID.
	//
	// example:
	//
	// dlc-xxxxxx
	SourceId *string `json:"SourceId,omitempty" xml:"SourceId,omitempty"`
	// The source type.
	//
	// example:
	//
	// job
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// The directory of summary.
	//
	// example:
	//
	// /root/data/
	SummaryPath *string `json:"SummaryPath,omitempty" xml:"SummaryPath,omitempty"`
	// The relative path of summary.
	//
	// example:
	//
	// /summary/
	SummaryRelativePath *string `json:"SummaryRelativePath,omitempty" xml:"SummaryRelativePath,omitempty"`
	// The configurations of datasets mounted with the TensorBoard job.
	TensorboardDataSources []*TensorboardDataSourceSpec `json:"TensorboardDataSources,omitempty" xml:"TensorboardDataSources,omitempty" type:"Repeated"`
	// The pay-as-you-go configuration of TensorBoard, which is used to create TensorBoard jobs that use pay-as-you-go resources.
	TensorboardSpec *TensorboardSpec `json:"TensorboardSpec,omitempty" xml:"TensorboardSpec,omitempty"`
	// The dataset URI.
	//
	// 	- Value format when DataSourceType is set to OSS: `oss://[oss-bucket].[endpoint]/[path]`.
	//
	// 	- Value format when DataSourceType is set to NAS:`nas://[nas-filesystem-id].[region]/[path]`.
	//
	// example:
	//
	// oss://.oss-cn-shanghai-finance-1.aliyuncs.com/
	Uri *string `json:"Uri,omitempty" xml:"Uri,omitempty"`
	// The workspace ID.
	//
	// <props="china">Call [ListWorkspaces](https://help.aliyun.com/document_detail/449124.html) to obtain the workspace ID.
	//
	// example:
	//
	// 123***
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s CreateTensorboardRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateTensorboardRequest) GoString() string {
	return s.String()
}

func (s *CreateTensorboardRequest) SetAccessibility(v string) *CreateTensorboardRequest {
	s.Accessibility = &v
	return s
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

func (s *CreateTensorboardRequest) SetPriority(v string) *CreateTensorboardRequest {
	s.Priority = &v
	return s
}

func (s *CreateTensorboardRequest) SetQuotaId(v string) *CreateTensorboardRequest {
	s.QuotaId = &v
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

func (s *CreateTensorboardRequest) SetTensorboardDataSources(v []*TensorboardDataSourceSpec) *CreateTensorboardRequest {
	s.TensorboardDataSources = v
	return s
}

func (s *CreateTensorboardRequest) SetTensorboardSpec(v *TensorboardSpec) *CreateTensorboardRequest {
	s.TensorboardSpec = v
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
	// The dataset ID.
	//
	// example:
	//
	// ds-20210126170216-xxxxxxxx
	DataSourceId *string `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	// The job ID.
	//
	// example:
	//
	// dlc-20210126170216-xxxxxxxx
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// TensorBoard ID
	//
	// example:
	//
	// tbxxxxxxxx
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
	// The job ID.
	//
	// example:
	//
	// dlc*************
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The request ID. You can troubleshoot issues based on the request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
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
	// The workspace ID.
	//
	// <props="china">For more information about how to obtain the workspace ID, see [ListWorkspaces](https://help.aliyun.com/document_detail/449124.html).
	//
	// example:
	//
	// 46099
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
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The TensorBoard ID.
	//
	// example:
	//
	// tensorboard-20210114104214-vf9lowjt3pso
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
	// Specifies whether to return the job details. Default value: true.
	//
	// example:
	//
	// true
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
	// The visibility of the job. Valid values:
	//
	// 	- PUBLIC: The code is public in the workspace.
	//
	// 	- PRIVATE: The workspace is visible only to you and the administrator of the workspace. This is the default value.
	//
	// example:
	//
	// PRIVATE
	Accessibility *string `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	// The cluster ID.
	//
	// example:
	//
	// a*****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The code source.
	CodeSource *GetJobResponseBodyCodeSource `json:"CodeSource,omitempty" xml:"CodeSource,omitempty" type:"Struct"`
	// The access credential configurations.
	CredentialConfig *CredentialConfig `json:"CredentialConfig,omitempty" xml:"CredentialConfig,omitempty"`
	// The data sources.
	DataSources []*GetJobResponseBodyDataSources `json:"DataSources,omitempty" xml:"DataSources,omitempty" type:"Repeated"`
	// The job name.
	//
	// example:
	//
	// tf-mnist-test
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The duration of the job (seconds).
	//
	// example:
	//
	// 3602
	Duration *int64 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The elastic job parameters.
	ElasticSpec *JobElasticSpec `json:"ElasticSpec,omitempty" xml:"ElasticSpec,omitempty"`
	// Specifies whether to enable the debugger job.
	//
	// example:
	//
	// false
	EnabledDebugger *bool `json:"EnabledDebugger,omitempty" xml:"EnabledDebugger,omitempty"`
	// The configurations of environment variables.
	Envs map[string]*string `json:"Envs,omitempty" xml:"Envs,omitempty"`
	// The time when the job was created (UTC).
	//
	// example:
	//
	// 2021-01-12T14:35:01Z
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// The time of the job failed (UTC).
	//
	// example:
	//
	// 2021-01-12T15:36:08Z
	GmtFailedTime *string `json:"GmtFailedTime,omitempty" xml:"GmtFailedTime,omitempty"`
	// The time when the job ended (UTC).
	//
	// example:
	//
	// 2021-01-12T15:36:08Z
	GmtFinishTime *string `json:"GmtFinishTime,omitempty" xml:"GmtFinishTime,omitempty"`
	// The start time of the job (UTC).
	//
	// example:
	//
	// 2021-01-12T14:36:21Z
	GmtRunningTime *string `json:"GmtRunningTime,omitempty" xml:"GmtRunningTime,omitempty"`
	// The time when the job stopped (UTC).
	//
	// example:
	//
	// 2021-01-12T15:36:08Z
	GmtStoppedTime *string `json:"GmtStoppedTime,omitempty" xml:"GmtStoppedTime,omitempty"`
	// The time when the job was submitted to the cluster (UTC).
	//
	// example:
	//
	// 2021-01-12T14:36:01Z
	GmtSubmittedTime *string `json:"GmtSubmittedTime,omitempty" xml:"GmtSubmittedTime,omitempty"`
	// The time when the job succeeded (UTC).
	//
	// example:
	//
	// 2021-01-12T15:36:08Z
	GmtSuccessedTime *string `json:"GmtSuccessedTime,omitempty" xml:"GmtSuccessedTime,omitempty"`
	// The job ID.
	//
	// example:
	//
	// dlc*******
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The node configurations of the job, which is **JobSpecs*	- in the CreateJob operation.
	JobSpecs []*JobSpec `json:"JobSpecs,omitempty" xml:"JobSpecs,omitempty" type:"Repeated"`
	// The job type. Specified by the JobType parameter of the [CreateJob](https://help.aliyun.com/document_detail/459672.html) operation.
	//
	// example:
	//
	// TFJob
	JobType *string `json:"JobType,omitempty" xml:"JobType,omitempty"`
	// All running nodes of the job.
	Pods []*GetJobResponseBodyPods `json:"Pods,omitempty" xml:"Pods,omitempty" type:"Repeated"`
	// The priority of the job. Valid values: 1 to 9.
	//
	// example:
	//
	// 1
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The status detail code, which is a sub-status under the current status.
	//
	// example:
	//
	// JobStoppedByUser
	ReasonCode *string `json:"ReasonCode,omitempty" xml:"ReasonCode,omitempty"`
	// The description of the status detail code.
	//
	// example:
	//
	// Job is stopped by user.
	ReasonMessage *string `json:"ReasonMessage,omitempty" xml:"ReasonMessage,omitempty"`
	// The request ID, which can be used for troubleshooting.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-xxxxxxxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the resource group to which the job belongs.
	//
	// example:
	//
	// r******
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The resource level that the job uses.
	//
	// example:
	//
	// L0
	ResourceLevel *string `json:"ResourceLevel,omitempty" xml:"ResourceLevel,omitempty"`
	// The resource type. Valid values: ECS, Lingjun, and ACS.
	//
	// example:
	//
	// ECS
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The number of retries and the maximum number of retries used by the job.
	//
	// example:
	//
	// 0/10
	RestartTimes *string `json:"RestartTimes,omitempty" xml:"RestartTimes,omitempty"`
	// The settings of the additional parameters of the job.
	Settings *JobSettings `json:"Settings,omitempty" xml:"Settings,omitempty"`
	// The status of the job. Valid values:
	//
	// 	- Creating
	//
	// 	- Queuing
	//
	// 	- Bidding (Only for Lingjun preemptible jobs)
	//
	// 	- EnvPreparing
	//
	// 	- SanityChecking
	//
	// 	- Running
	//
	// 	- Restarting
	//
	// 	- Stopping
	//
	// 	- SucceededReserving
	//
	// 	- FailedReserving
	//
	// 	- Succeeded
	//
	// 	- Failed
	//
	// 	- Stopped
	//
	// example:
	//
	// Stopped
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The status history.
	StatusHistory []*StatusTransitionItem `json:"StatusHistory,omitempty" xml:"StatusHistory,omitempty" type:"Repeated"`
	// The sub-status of the job, such as its preemption status.
	//
	// example:
	//
	// Restarting
	SubStatus *string `json:"SubStatus,omitempty" xml:"SubStatus,omitempty"`
	// The tenant ID.
	//
	// example:
	//
	// GAR***W134
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// The directory that contains requirements.txt.
	//
	// example:
	//
	// /root/code/
	ThirdpartyLibDir *string `json:"ThirdpartyLibDir,omitempty" xml:"ThirdpartyLibDir,omitempty"`
	// The third-party Python libraries to be installed.
	ThirdpartyLibs []*string `json:"ThirdpartyLibs,omitempty" xml:"ThirdpartyLibs,omitempty" type:"Repeated"`
	// The command that is run to start each node.
	//
	// example:
	//
	// python /root/code/mnist.py
	UserCommand *string `json:"UserCommand,omitempty" xml:"UserCommand,omitempty"`
	// The UID of the Alibaba Cloud account who submitted the job.
	//
	// example:
	//
	// 12*********
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The VPC of the user.
	UserVpc *GetJobResponseBodyUserVpc `json:"UserVpc,omitempty" xml:"UserVpc,omitempty" type:"Struct"`
	// The ID of the workspace to which the job belongs.
	//
	// example:
	//
	// 268
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
	// The name of the workspace to which the job belongs.
	//
	// example:
	//
	// dlc-workspace
	WorkspaceName *string `json:"WorkspaceName,omitempty" xml:"WorkspaceName,omitempty"`
}

func (s GetJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetJobResponseBody) GoString() string {
	return s.String()
}

func (s *GetJobResponseBody) SetAccessibility(v string) *GetJobResponseBody {
	s.Accessibility = &v
	return s
}

func (s *GetJobResponseBody) SetClusterId(v string) *GetJobResponseBody {
	s.ClusterId = &v
	return s
}

func (s *GetJobResponseBody) SetCodeSource(v *GetJobResponseBodyCodeSource) *GetJobResponseBody {
	s.CodeSource = v
	return s
}

func (s *GetJobResponseBody) SetCredentialConfig(v *CredentialConfig) *GetJobResponseBody {
	s.CredentialConfig = v
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

func (s *GetJobResponseBody) SetUserVpc(v *GetJobResponseBodyUserVpc) *GetJobResponseBody {
	s.UserVpc = v
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
	// The code branch.
	//
	// example:
	//
	// master
	Branch *string `json:"Branch,omitempty" xml:"Branch,omitempty"`
	// The code source ID.
	//
	// example:
	//
	// code******
	CodeSourceId *string `json:"CodeSourceId,omitempty" xml:"CodeSourceId,omitempty"`
	// The code commit ID
	//
	// example:
	//
	// 44da109b59f8596152987eaa8f3b2487xxxxxx
	Commit *string `json:"Commit,omitempty" xml:"Commit,omitempty"`
	// The local mount path.
	//
	// example:
	//
	// /mnt/data
	MountPath *string `json:"MountPath,omitempty" xml:"MountPath,omitempty"`
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
	// The data source ID.
	//
	// example:
	//
	// d*******
	DataSourceId *string `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	// The local mount path. This parameter is optional. The default value is empty, which specifies that the mount path in the data source is used.
	//
	// example:
	//
	// /mnt/data/
	MountPath *string `json:"MountPath,omitempty" xml:"MountPath,omitempty"`
	// The data source URL.
	//
	// example:
	//
	// oss://bucket.oss-cn-hangzhou-internal.aliyuncs.com/path/
	Uri *string `json:"Uri,omitempty" xml:"Uri,omitempty"`
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
	// The time when the node was created (UTC).
	//
	// example:
	//
	// 2021-01-12T14:36:01Z
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// The end time of the node (UTC).
	//
	// example:
	//
	// 2021-01-12T15:36:05Z
	GmtFinishTime *string `json:"GmtFinishTime,omitempty" xml:"GmtFinishTime,omitempty"`
	// The start time of the node (UTC).
	//
	// example:
	//
	// 2021-01-12T14:36:01Z
	GmtStartTime *string `json:"GmtStartTime,omitempty" xml:"GmtStartTime,omitempty"`
	// The historical nodes.
	HistoryPods []*GetJobResponseBodyPodsHistoryPods `json:"HistoryPods,omitempty" xml:"HistoryPods,omitempty" type:"Repeated"`
	// The IP address of the node.
	//
	// example:
	//
	// 10.0.1.2
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The node ID. It can be used in the GetPodLogs and GetPodEvents operations to obtain the detailed logs and events of the node.
	//
	// example:
	//
	// Worker
	PodId *string `json:"PodId,omitempty" xml:"PodId,omitempty"`
	// The UID of the node.
	//
	// example:
	//
	// fe846462-af2c-4521-bd6f-96787a57591d
	PodUid *string `json:"PodUid,omitempty" xml:"PodUid,omitempty"`
	// The resource type of the node.
	//
	// example:
	//
	// Normal
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The status of the node. Valid values:
	//
	// 	- Pending
	//
	// 	- Running
	//
	// 	- Succeeded
	//
	// 	- Failed
	//
	// 	- Unknown
	//
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The sub-status of the node, such as its preemption status. Valid values:
	//
	// 	- Normal
	//
	// 	- Evicted
	//
	// example:
	//
	// Normal
	SubStatus *string `json:"SubStatus,omitempty" xml:"SubStatus,omitempty"`
	// The node type, which corresponds to a specific JobSpec in JobSpecs of the CreateJob operation.
	//
	// example:
	//
	// Worker
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
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
	// The time when the node was created (UTC).
	//
	// example:
	//
	// 2021-01-12T14:36:01Z
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// The end time of the node (UTC).
	//
	// example:
	//
	// 2021-01-12T14:36:01Z
	GmtFinishTime *string `json:"GmtFinishTime,omitempty" xml:"GmtFinishTime,omitempty"`
	// The start time of the node (UTC).
	//
	// example:
	//
	// 2021-01-12T14:36:01Z
	GmtStartTime *string `json:"GmtStartTime,omitempty" xml:"GmtStartTime,omitempty"`
	// The IP address of the node.
	//
	// example:
	//
	// 10.0.1.3
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The ID of the node.
	//
	// example:
	//
	// Worker
	PodId *string `json:"PodId,omitempty" xml:"PodId,omitempty"`
	// The UID of the node.
	//
	// example:
	//
	// fe846462-af2c-4521-bd6f-96787a57591d
	PodUid *string `json:"PodUid,omitempty" xml:"PodUid,omitempty"`
	// The resource type of the node.
	//
	// example:
	//
	// Normal
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The status of the node.
	//
	// example:
	//
	// Failed
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The sub-status of the node, such as its preemption status. Valid values:
	//
	// 	- Normal
	//
	// 	- Evicted
	//
	// example:
	//
	// Normal
	SubStatus *string `json:"SubStatus,omitempty" xml:"SubStatus,omitempty"`
	// The type of the node.
	//
	// example:
	//
	// Worker
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
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

type GetJobResponseBodyUserVpc struct {
	// The default router. This parameter is valid only for general-purpose computing resources. Valid values:
	//
	// eth0: The default network interface is used to access the Internet through the public gateway. eth1: The user\\"s Elastic Network Interface is used to access the Internet through the private gateway.
	DefaultRoute *string `json:"DefaultRoute,omitempty" xml:"DefaultRoute,omitempty"`
	// The extended CIDR block. Example: 192.168.0.1/24.
	ExtendedCidrs []*string `json:"ExtendedCidrs,omitempty" xml:"ExtendedCidrs,omitempty" type:"Repeated"`
	// The security group ID.
	//
	// example:
	//
	// sg-abcdef****
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The vSwitch ID.
	//
	// example:
	//
	// vs-abcdef****
	SwitchId *string `json:"SwitchId,omitempty" xml:"SwitchId,omitempty"`
	// The VPC ID.
	//
	// example:
	//
	// vpc-abcdef****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s GetJobResponseBodyUserVpc) String() string {
	return tea.Prettify(s)
}

func (s GetJobResponseBodyUserVpc) GoString() string {
	return s.String()
}

func (s *GetJobResponseBodyUserVpc) SetDefaultRoute(v string) *GetJobResponseBodyUserVpc {
	s.DefaultRoute = &v
	return s
}

func (s *GetJobResponseBodyUserVpc) SetExtendedCidrs(v []*string) *GetJobResponseBodyUserVpc {
	s.ExtendedCidrs = v
	return s
}

func (s *GetJobResponseBodyUserVpc) SetSecurityGroupId(v string) *GetJobResponseBodyUserVpc {
	s.SecurityGroupId = &v
	return s
}

func (s *GetJobResponseBodyUserVpc) SetSwitchId(v string) *GetJobResponseBodyUserVpc {
	s.SwitchId = &v
	return s
}

func (s *GetJobResponseBodyUserVpc) SetVpcId(v string) *GetJobResponseBodyUserVpc {
	s.VpcId = &v
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
	// The end time (UTC) of the time range for querying events. The default value is the current time.
	//
	// example:
	//
	// 2020-11-08T18:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The maximum number of events that can be returned. Default value: 2000.
	//
	// example:
	//
	// 100
	MaxEventsNum *int32 `json:"MaxEventsNum,omitempty" xml:"MaxEventsNum,omitempty"`
	// The start time (UTC) of the time range for querying events. The default value is 7 days ago.
	//
	// example:
	//
	// 2020-11-08T16:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
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
	// The events.
	Events []*string `json:"Events,omitempty" xml:"Events,omitempty" type:"Repeated"`
	// The job ID.
	//
	// example:
	//
	// dlc-20210126170216-******
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The request ID, which can be used for troubleshooting.
	//
	// example:
	//
	// 78F6FCE2-278F-4C4A-A6B7-DD8ECEA9C456
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// The end time of the time range to query monitoring data. The time is displayed in UTC. The default value is the current time.
	//
	// example:
	//
	// 2020-11-09T16:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The type of the monitoring metrics. Valid values:
	//
	// 	- GpuCoreUsage: GPU utilization
	//
	// 	- GpuMemoryUsage: GPU memory utilization
	//
	// 	- CpuCoreUsage: CPU utilization
	//
	// 	- MemoryUsage: memory utilization
	//
	// 	- NetworkInputRate: the network write in rate.
	//
	// 	- NetworkOutputRate: the network write out rate
	//
	// 	- DiskReadRate: the disk read rate
	//
	// 	- DiskWriteRate: the disk write rate
	//
	// This parameter is required.
	//
	// example:
	//
	// GpuMemoryUsage
	MetricType *string `json:"MetricType,omitempty" xml:"MetricType,omitempty"`
	// The beginning of the time range to query monitoring data. The time is displayed in UTC. The default value is the time 1 hour before the current time.
	//
	// example:
	//
	// 2020-11-08T16:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The interval at which monitoring data is returned. Default value: 5. Unit: minutes.
	//
	// example:
	//
	// 5m
	TimeStep *string `json:"TimeStep,omitempty" xml:"TimeStep,omitempty"`
	// The temporary token used for authentication.
	//
	// example:
	//
	// eyXXXX-XXXX.XXXXX
	Token *string `json:"Token,omitempty" xml:"Token,omitempty"`
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
	// The job ID.
	//
	// example:
	//
	// dlc-20210126170216-*******
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The monitoring metrics of the job.
	PodMetrics []*PodMetric `json:"PodMetrics,omitempty" xml:"PodMetrics,omitempty" type:"Repeated"`
	// The request ID. You can troubleshoot issues based on the request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// The nth time for which the job sanity check is performed.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	SanityCheckNumber *int32 `json:"SanityCheckNumber,omitempty" xml:"SanityCheckNumber,omitempty"`
	// The phase in which the job sanity check is performed.
	//
	// 	- CheckInit
	//
	// 	- DeviceCheck
	//
	// 	- SingleNodeCommCheck
	//
	// 	- TwoNodeCommCheck
	//
	// 	- AllNodeCommCheck
	//
	// example:
	//
	// DeviceCheck
	SanityCheckPhase *string `json:"SanityCheckPhase,omitempty" xml:"SanityCheckPhase,omitempty"`
	// The token information for job sharing. For more information about how to obtain the token information, see [GetToken](https://help.aliyun.com/document_detail/2557812.html).
	//
	// example:
	//
	// eyJhbG******zI1NiIsInR5cCI6IkpXVCJ9.eyJle****jE3MDk1Mzk0NDIsImlhdCI6MTcwODkzNDY0MiwidXNlcl9pZCI6IjE3NTgwNTQxNjI0Mzg2NTUiLCJ0YXJnZXRfaWQiOiJkbGM1OGh1a2xyYzZwdGMyIiwidGFyZ2V0X3R5cGUiOiJqb2IifQ.GNL7jo6****mgKKv0QeGIYgvBufSU-PH_EQttX****
	Token *string `json:"Token,omitempty" xml:"Token,omitempty"`
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
	// The job ID.
	//
	// example:
	//
	// dlc-20210126170216-xxxxxx
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// B3789344-F1xxxBE-5xx2-A04D-xxxxx
	RequestID *string `json:"RequestID,omitempty" xml:"RequestID,omitempty"`
	// The job sanity check result.
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
	// The end time (UTC).
	//
	// example:
	//
	// 2020-11-09T16:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The maximum number of events that can be returned.
	//
	// example:
	//
	// 100
	MaxEventsNum *int32 `json:"MaxEventsNum,omitempty" xml:"MaxEventsNum,omitempty"`
	// The node UID. Call [GetJob](https://help.aliyun.com/document_detail/459677.html) to get the node UID.
	//
	// example:
	//
	// dlc-20210126170216-*****-chief-0
	PodUid *string `json:"PodUid,omitempty" xml:"PodUid,omitempty"`
	// The start time (UTC).
	//
	// example:
	//
	// 2020-11-08T16:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
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
	// The events returned.
	Events []*string `json:"Events,omitempty" xml:"Events,omitempty" type:"Repeated"`
	// The job ID.
	//
	// example:
	//
	// dlc-20210126170216-*****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The node ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// dlc-20210126170216-*****-chief-0
	PodId *string `json:"PodId,omitempty" xml:"PodId,omitempty"`
	// The node UID.
	//
	// example:
	//
	// 94a7cc7c-0033-48b5-85bd-71c63592c268
	PodUid *string `json:"PodUid,omitempty" xml:"PodUid,omitempty"`
	// The request ID, which can be used for troubleshooting.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// Specifies whether to download the log file. Default value: false. Valid values:
	//
	// 	- false
	//
	// 	- true
	//
	// example:
	//
	// true
	DownloadToFile *bool `json:"DownloadToFile,omitempty" xml:"DownloadToFile,omitempty"`
	// The end time of the query. Default value: current time.
	//
	// example:
	//
	// 2020-11-08T17:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The maximum number of log entries. Default value: 2000.
	//
	// example:
	//
	// 100
	MaxLines *int32 `json:"MaxLines,omitempty" xml:"MaxLines,omitempty"`
	// The node UID. For more information about how to obtain a node UID, see [GetJob](https://help.aliyun.com/document_detail/459677.html).
	//
	// example:
	//
	// fe846462-af2c-4521-bd6f-96787a57****
	PodUid *string `json:"PodUid,omitempty" xml:"PodUid,omitempty"`
	// The start time of the query. Default value: 7 days ago.
	//
	// example:
	//
	// 2020-11-08T16:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
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
	// The job ID.
	//
	// example:
	//
	// dlc-20210126170216-******
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The logs.
	Logs []*string `json:"Logs,omitempty" xml:"Logs,omitempty" type:"Repeated"`
	// The node ID.
	//
	// example:
	//
	// dlc-20210126170216-****-chief-0
	PodId *string `json:"PodId,omitempty" xml:"PodId,omitempty"`
	// The instance UID.
	//
	// example:
	//
	// 94a7cc7c-0033-48b5-85bd-71c63592c268
	PodUid *string `json:"PodUid,omitempty" xml:"PodUid,omitempty"`
	// The request ID which is used for diagnostics and Q\\&A.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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

type GetRayDashboardRequest struct {
	// Specifies whether the link is a sharing link. If yes, a token is required.
	//
	// Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// false
	IsShared *bool `json:"isShared,omitempty" xml:"isShared,omitempty"`
	// The token obtained from GetToken
	//
	// example:
	//
	// some_token_value
	Token *string `json:"token,omitempty" xml:"token,omitempty"`
}

func (s GetRayDashboardRequest) String() string {
	return tea.Prettify(s)
}

func (s GetRayDashboardRequest) GoString() string {
	return s.String()
}

func (s *GetRayDashboardRequest) SetIsShared(v bool) *GetRayDashboardRequest {
	s.IsShared = &v
	return s
}

func (s *GetRayDashboardRequest) SetToken(v string) *GetRayDashboardRequest {
	s.Token = &v
	return s
}

type GetRayDashboardResponseBody struct {
	// Indicates whether the dashboard has been integrated with CloudMonitor and supports ray metrics
	//
	// example:
	//
	// true
	MetricsEnabled *string `json:"metricsEnabled,omitempty" xml:"metricsEnabled,omitempty"`
	// The Ray Dashboard URL
	//
	// example:
	//
	// https://pre-pai-dlc-proxy-cn-hangzhou.aliyun.com/ray/dashboard/dlc1k7426goc7bvy
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s GetRayDashboardResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetRayDashboardResponseBody) GoString() string {
	return s.String()
}

func (s *GetRayDashboardResponseBody) SetMetricsEnabled(v string) *GetRayDashboardResponseBody {
	s.MetricsEnabled = &v
	return s
}

func (s *GetRayDashboardResponseBody) SetUrl(v string) *GetRayDashboardResponseBody {
	s.Url = &v
	return s
}

type GetRayDashboardResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRayDashboardResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRayDashboardResponse) String() string {
	return tea.Prettify(s)
}

func (s GetRayDashboardResponse) GoString() string {
	return s.String()
}

func (s *GetRayDashboardResponse) SetHeaders(v map[string]*string) *GetRayDashboardResponse {
	s.Headers = v
	return s
}

func (s *GetRayDashboardResponse) SetStatusCode(v int32) *GetRayDashboardResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRayDashboardResponse) SetBody(v *GetRayDashboardResponseBody) *GetRayDashboardResponse {
	s.Body = v
	return s
}

type GetTensorboardRequest struct {
	// The job ID. For more information about how to query the job ID, see [ListJob](https://help.aliyun.com/document_detail/459676.html).
	//
	// example:
	//
	// dlc-xxxxxxxx
	JodId *string `json:"JodId,omitempty" xml:"JodId,omitempty"`
	// The information about the shared token. You can specify this parameter to obtain the permission to view a TensorBoard job based on the shared token information. You can execute [GetTensorboardSharedUrl](https://help.aliyun.com/document_detail/2557813.html) and extract the shared token from the obtained information.
	//
	// example:
	//
	// eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.e
	//
	// yJleHAiOjE2OTUyODA0NTMsImlhdCI6MTY5NTE5NDA1MywidXNlcl9pZCI6IjExN
	//
	// Tc3MDMyNzA5OTQ5MDEiLCJ0YXJnZXRfaWQiOiJ0YjRrOGxjNXhmdTM2b3B0Iiw
	//
	// idGFyZ2V0X3R5cGUiOiJ0ZW5zb3Jib2FyZCJ9.6eT68J-KMBwwfN2d7fj7u6vyPcf0erfqYeizd2N****
	Token *string `json:"Token,omitempty" xml:"Token,omitempty"`
	// The workspace ID.
	//
	// <props="china">For more information about how to query the workspace ID, see [ListWorkspaces](https://help.aliyun.com/document_detail/449124.html).
	//
	// example:
	//
	// 46099
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
	// The validity period of the shareable link. Unit: seconds. Maximum value: 604800.
	//
	// example:
	//
	// 86400
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
	// The request ID which is used for troubleshooting.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The shareable link of the TensorBoard task.
	//
	// example:
	//
	// http://pai-dlc-proxy-xxx.alicyuncs.com/xxx/xxx/token/
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
	// The time when the share link expires. Default value: 604800. Minimum value: 0. Unit: seconds.
	//
	// example:
	//
	// 60
	ExpireTime *int64 `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The ID of the job to be shared.
	//
	// example:
	//
	// dlc*******
	TargetId *string `json:"TargetId,omitempty" xml:"TargetId,omitempty"`
	// The type of the job that you want to share. Valid values: job and tensorboard.
	//
	// example:
	//
	// job
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
	// The request ID, which is used to troubleshoot issues.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-xxxxxxxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The sharing token, used to view the information about the shared job.
	//
	// example:
	//
	// eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9*****
	Token *string `json:"Token,omitempty" xml:"Token,omitempty"`
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
	// Specifies whether to create a shareable link to access the container. Valid values:
	//
	// 	- true: returns a shareable link to access the container. The link will expire after 30 seconds and can only be used once. After you access the container by using the link, other requests that use this link to access the container become invalid.
	//
	// 	- false: returns a common shareable link to access the container. If you use a common shareable link to access a container, Alibaba Cloud identity authentication is required. The link will expire after 30 seconds.
	//
	// example:
	//
	// true
	IsShared *bool `json:"IsShared,omitempty" xml:"IsShared,omitempty"`
	// The pod UID.
	//
	// example:
	//
	// 94a7cc7c-0033-48b5-85bd-71c63592c268
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
	// The request ID which is used for diagnostics and Q\\&A.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The WebSocket URI for accessing the container. You must build a WebSocket client. For more information about the communication format, see the following code:
	//
	//     ws = new WebSocket(
	//
	//       `wss://xxxxx`,
	//
	//     );
	//
	//     ws.onopen = function open() {
	//
	//       console.warn(\\"connected\\");
	//
	//       term.write(\\"\\r\\");
	//
	//     };
	//
	//     ws.onclose = function close() {
	//
	//       console.warn(\\"disconnected\\");
	//
	//       term.write(\\"Connection closed\\");
	//
	//     };
	//
	//     // Return the following information in the backend.
	//
	//     ws.onmessage = function incoming(event) {
	//
	//       const msg = JSON.parse(event.data);
	//
	//       console.warn(msg);
	//
	//       if (msg.operation === \\"stdout\\") {
	//
	//         term.write(msg.data);
	//
	//       } else {
	//
	//         console.warn(\\"invalid msg operation: \\" + msg);
	//
	//       }
	//
	//     };
	//
	//     // Enter the following code in the console.
	//
	//     term.onData(data => {
	//
	//       const msg = { operation: \\"stdin\\", data: data };
	//
	//       ws.send(JSON.stringify(msg));
	//
	//     });
	//
	//     term.onResize(size => {
	//
	//       const msg = { operation: \\"resize\\", cols: size.cols, rows: size.rows };
	//
	//       ws.send(JSON.stringify(msg));
	//
	//     });
	//
	//     fitAddon.fit();
	//
	//     };
	//
	// example:
	//
	// wss://*****
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
	// Filter by accelerator type. Valid values:
	//
	// 	- CPU
	//
	// 	- GPU
	//
	// example:
	//
	// GPU
	AcceleratorType *string `json:"AcceleratorType,omitempty" xml:"AcceleratorType,omitempty"`
	// The instance types to query. Separate the types with commas (,).
	//
	// example:
	//
	// ecs.g6.large,ecs.g6.xlarge
	InstanceTypes *string `json:"InstanceTypes,omitempty" xml:"InstanceTypes,omitempty"`
	// The sorting order. Valid values:
	//
	// 	- desc: descending order.
	//
	// 	- asc: ascending order.
	//
	// example:
	//
	// desc
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// The number of the page to query. The start value is 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The type of the resource. Valid values:
	//
	// 	- ECS
	//
	// 	- Lingjun
	//
	// example:
	//
	// ECS
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The field based on which the results are sorted. Valid values:
	//
	// 	- CPU
	//
	// 	- GPU
	//
	// 	- Memory
	//
	// 	- GmtCreateTime
	//
	// example:
	//
	// Gpu
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
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
	// The list of ECS specifications.
	EcsSpecs []*EcsSpec `json:"EcsSpecs,omitempty" xml:"EcsSpecs,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of types that meet the filter conditions.
	//
	// example:
	//
	// 10
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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
	// The sorting order:
	//
	// 	- desc: descending order
	//
	// 	- asc: ascending order
	//
	// example:
	//
	// desc
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
	// The request ID.
	//
	// example:
	//
	// 1AC9xxx-3xxx-5xxx2-xxxx-FA5
	RequestID *string `json:"RequestID,omitempty" xml:"RequestID,omitempty"`
	// The sanity check results.
	SanityCheckResults [][]*SanityCheckResultItem `json:"SanityCheckResults,omitempty" xml:"SanityCheckResults,omitempty" type:"Repeated"`
	// The total number of results that meet the filter conditions.
	//
	// example:
	//
	// 10
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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
	// The job visibility. Valid values:
	//
	// 	- PUBLIC: The job is visible to all members in the workspace.
	//
	// 	- PRIVATE: The job is visible only to you and the administrator of the workspace.
	//
	// example:
	//
	// PRIVATE
	Accessibility *string `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	// The ID of the user associated with the job.
	//
	// example:
	//
	// 16****
	BusinessUserId *string `json:"BusinessUserId,omitempty" xml:"BusinessUserId,omitempty"`
	// The caller.
	//
	// example:
	//
	// local
	Caller *string `json:"Caller,omitempty" xml:"Caller,omitempty"`
	// The job name. Fuzzy query is supported. The name is case-insensitive. Wildcards are not supported. For example, if you enter test, test-job1, job-test, job-test2, or job-test can be matched, and job-t1 cannot be matched. The default value null indicates any job name.
	//
	// example:
	//
	// tf-mnist-test
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The end time of the query. Use the job creation time to filter data. The default value is the current time.
	//
	// example:
	//
	// 2020-11-09T14:45:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// Specifies whether to query a list of jobs across workspaces. This parameter must be used together with `ShowOwn=true`. You can use this parameter to query a list of jobs recently submitted by the current user.
	//
	// example:
	//
	// false
	FromAllWorkspaces *bool `json:"FromAllWorkspaces,omitempty" xml:"FromAllWorkspaces,omitempty"`
	// The job ID. Fuzzy query is supported. The name is case-insensitive. Wildcards are not supported. The default value null indicates any job ID.
	//
	// example:
	//
	// dlc********
	JobId  *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	JobIds *string `json:"JobIds,omitempty" xml:"JobIds,omitempty"`
	// The job type. The default value null indicates any type. Valid values:
	//
	// 	- TFJob
	//
	// 	- PyTorchJob
	//
	// 	- XGBoostJob
	//
	// 	- OneFlowJob
	//
	// 	- ElasticBatchJob
	//
	// example:
	//
	// TFJob
	JobType *string `json:"JobType,omitempty" xml:"JobType,omitempty"`
	// The sorting order. Valid values:
	//
	// 	- desc (default)
	//
	// 	- asc
	//
	// example:
	//
	// desc
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// The Idle resource information. Valid values:
	//
	// 	- ForbiddenQuotaOverSold
	//
	// 	- ForceQuotaOverSold
	//
	// 	- AcceptQuotaOverSold-true (true indicates that the job uses idle resources.)
	//
	// 	- AcceptQuotaOverSold-false (false indicates that the job uses guaranteed resources.)
	//
	// example:
	//
	// ForbiddenQuotaOverSold
	OversoldInfo *string `json:"OversoldInfo,omitempty" xml:"OversoldInfo,omitempty"`
	// The number of the page to return for the current query. Minimum value: 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of jobs per page.
	//
	// example:
	//
	// 50
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The type of the resource. Valid values:
	//
	// 	- PrePaid: Resource quota
	//
	// 	- Spot: Preemptible resources
	//
	// 	- PostPaid: Public resources
	//
	// example:
	//
	// PostPaid
	PaymentType *string `json:"PaymentType,omitempty" xml:"PaymentType,omitempty"`
	// The specific pipeline ID used to filter jobs.
	//
	// example:
	//
	// flow-*******
	PipelineId *string `json:"PipelineId,omitempty" xml:"PipelineId,omitempty"`
	// The resource group ID. For information about how to obtain the ID of a dedicated resource group, see [Manage resource quota](https://help.aliyun.com/document_detail/2651299.html).
	//
	// example:
	//
	// r*****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The resource quota name used to filter jobs. Fuzzy search is supported. Wildcards are not supported. The default value null indicates that jobs are not filtered by resource quota name.
	//
	// example:
	//
	// quota***
	ResourceQuotaName *string `json:"ResourceQuotaName,omitempty" xml:"ResourceQuotaName,omitempty"`
	// Specifies whether to query only the jobs submitted by the current user.
	//
	// example:
	//
	// true
	ShowOwn *bool `json:"ShowOwn,omitempty" xml:"ShowOwn,omitempty"`
	// The sorting field. Valid values:
	//
	// 	- DisplayName
	//
	// 	- JobType
	//
	// 	- Status
	//
	// 	- GmtCreateTime
	//
	// 	- GmtFinishTime
	//
	// example:
	//
	// GmtCreateTime
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// The start time of the query. Use the job creation time to filter data. The default value is the current time minus seven days. In other words, if you do not configure the StartTime and EndTime parameters, the system queries the job list in the last seven days.
	//
	// example:
	//
	// 2020-11-08T16:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The job status. Valid values:
	//
	// 	- Creating
	//
	// 	- Queuing
	//
	// 	- Bidding (only available for spot jobs that use Lingjun resources)
	//
	// 	- EnvPreparing
	//
	// 	- SanityChecking
	//
	// 	- Running
	//
	// 	- Restarting
	//
	// 	- Stopping
	//
	// 	- SucceededReserving
	//
	// 	- FailedReserving
	//
	// 	- Succeeded
	//
	// 	- Failed
	//
	// 	- Stopped
	//
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tags.
	Tags map[string]*string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The user ID used to filter jobs.
	//
	// example:
	//
	// 20**************
	UserIdForFilter *string `json:"UserIdForFilter,omitempty" xml:"UserIdForFilter,omitempty"`
	// The username used to filter jobs. Fuzzy search is supported. Wildcards are not supported. The default value null indicates that jobs are not filtered by username.
	//
	// example:
	//
	// test***
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
	// The workspace ID.
	//
	// example:
	//
	// 1****
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListJobsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListJobsRequest) GoString() string {
	return s.String()
}

func (s *ListJobsRequest) SetAccessibility(v string) *ListJobsRequest {
	s.Accessibility = &v
	return s
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

func (s *ListJobsRequest) SetJobIds(v string) *ListJobsRequest {
	s.JobIds = &v
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

func (s *ListJobsRequest) SetOversoldInfo(v string) *ListJobsRequest {
	s.OversoldInfo = &v
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

func (s *ListJobsRequest) SetPaymentType(v string) *ListJobsRequest {
	s.PaymentType = &v
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

func (s *ListJobsRequest) SetResourceQuotaName(v string) *ListJobsRequest {
	s.ResourceQuotaName = &v
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
	// The job visibility. Valid values:
	//
	// 	- PUBLIC: The job is visible to all members in the workspace.
	//
	// 	- PRIVATE: The job is visible only to you and the administrator of the workspace.
	//
	// example:
	//
	// PRIVATE
	Accessibility *string `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	// The ID of the user associated with the job.
	//
	// example:
	//
	// 16****
	BusinessUserId *string `json:"BusinessUserId,omitempty" xml:"BusinessUserId,omitempty"`
	// The caller.
	//
	// example:
	//
	// local
	Caller *string `json:"Caller,omitempty" xml:"Caller,omitempty"`
	// The job name. Fuzzy query is supported. The name is case-insensitive. Wildcards are not supported. For example, if you enter test, test-job1, job-test, job-test2, or job-test can be matched, and job-t1 cannot be matched. The default value null indicates any job name.
	//
	// example:
	//
	// tf-mnist-test
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The end time of the query. Use the job creation time to filter data. The default value is the current time.
	//
	// example:
	//
	// 2020-11-09T14:45:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// Specifies whether to query a list of jobs across workspaces. This parameter must be used together with `ShowOwn=true`. You can use this parameter to query a list of jobs recently submitted by the current user.
	//
	// example:
	//
	// false
	FromAllWorkspaces *bool `json:"FromAllWorkspaces,omitempty" xml:"FromAllWorkspaces,omitempty"`
	// The job ID. Fuzzy query is supported. The name is case-insensitive. Wildcards are not supported. The default value null indicates any job ID.
	//
	// example:
	//
	// dlc********
	JobId  *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	JobIds *string `json:"JobIds,omitempty" xml:"JobIds,omitempty"`
	// The job type. The default value null indicates any type. Valid values:
	//
	// 	- TFJob
	//
	// 	- PyTorchJob
	//
	// 	- XGBoostJob
	//
	// 	- OneFlowJob
	//
	// 	- ElasticBatchJob
	//
	// example:
	//
	// TFJob
	JobType *string `json:"JobType,omitempty" xml:"JobType,omitempty"`
	// The sorting order. Valid values:
	//
	// 	- desc (default)
	//
	// 	- asc
	//
	// example:
	//
	// desc
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// The Idle resource information. Valid values:
	//
	// 	- ForbiddenQuotaOverSold
	//
	// 	- ForceQuotaOverSold
	//
	// 	- AcceptQuotaOverSold-true (true indicates that the job uses idle resources.)
	//
	// 	- AcceptQuotaOverSold-false (false indicates that the job uses guaranteed resources.)
	//
	// example:
	//
	// ForbiddenQuotaOverSold
	OversoldInfo *string `json:"OversoldInfo,omitempty" xml:"OversoldInfo,omitempty"`
	// The number of the page to return for the current query. Minimum value: 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of jobs per page.
	//
	// example:
	//
	// 50
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The type of the resource. Valid values:
	//
	// 	- PrePaid: Resource quota
	//
	// 	- Spot: Preemptible resources
	//
	// 	- PostPaid: Public resources
	//
	// example:
	//
	// PostPaid
	PaymentType *string `json:"PaymentType,omitempty" xml:"PaymentType,omitempty"`
	// The specific pipeline ID used to filter jobs.
	//
	// example:
	//
	// flow-*******
	PipelineId *string `json:"PipelineId,omitempty" xml:"PipelineId,omitempty"`
	// The resource group ID. For information about how to obtain the ID of a dedicated resource group, see [Manage resource quota](https://help.aliyun.com/document_detail/2651299.html).
	//
	// example:
	//
	// r*****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The resource quota name used to filter jobs. Fuzzy search is supported. Wildcards are not supported. The default value null indicates that jobs are not filtered by resource quota name.
	//
	// example:
	//
	// quota***
	ResourceQuotaName *string `json:"ResourceQuotaName,omitempty" xml:"ResourceQuotaName,omitempty"`
	// Specifies whether to query only the jobs submitted by the current user.
	//
	// example:
	//
	// true
	ShowOwn *bool `json:"ShowOwn,omitempty" xml:"ShowOwn,omitempty"`
	// The sorting field. Valid values:
	//
	// 	- DisplayName
	//
	// 	- JobType
	//
	// 	- Status
	//
	// 	- GmtCreateTime
	//
	// 	- GmtFinishTime
	//
	// example:
	//
	// GmtCreateTime
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// The start time of the query. Use the job creation time to filter data. The default value is the current time minus seven days. In other words, if you do not configure the StartTime and EndTime parameters, the system queries the job list in the last seven days.
	//
	// example:
	//
	// 2020-11-08T16:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The job status. Valid values:
	//
	// 	- Creating
	//
	// 	- Queuing
	//
	// 	- Bidding (only available for spot jobs that use Lingjun resources)
	//
	// 	- EnvPreparing
	//
	// 	- SanityChecking
	//
	// 	- Running
	//
	// 	- Restarting
	//
	// 	- Stopping
	//
	// 	- SucceededReserving
	//
	// 	- FailedReserving
	//
	// 	- Succeeded
	//
	// 	- Failed
	//
	// 	- Stopped
	//
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tags.
	TagsShrink *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The user ID used to filter jobs.
	//
	// example:
	//
	// 20**************
	UserIdForFilter *string `json:"UserIdForFilter,omitempty" xml:"UserIdForFilter,omitempty"`
	// The username used to filter jobs. Fuzzy search is supported. Wildcards are not supported. The default value null indicates that jobs are not filtered by username.
	//
	// example:
	//
	// test***
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
	// The workspace ID.
	//
	// example:
	//
	// 1****
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListJobsShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListJobsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListJobsShrinkRequest) SetAccessibility(v string) *ListJobsShrinkRequest {
	s.Accessibility = &v
	return s
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

func (s *ListJobsShrinkRequest) SetJobIds(v string) *ListJobsShrinkRequest {
	s.JobIds = &v
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

func (s *ListJobsShrinkRequest) SetOversoldInfo(v string) *ListJobsShrinkRequest {
	s.OversoldInfo = &v
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

func (s *ListJobsShrinkRequest) SetPaymentType(v string) *ListJobsShrinkRequest {
	s.PaymentType = &v
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

func (s *ListJobsShrinkRequest) SetResourceQuotaName(v string) *ListJobsShrinkRequest {
	s.ResourceQuotaName = &v
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
	// The jobs.
	Jobs []*JobItem `json:"Jobs,omitempty" xml:"Jobs,omitempty" type:"Repeated"`
	// The request ID used to troubleshoot issues.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of jobs that meet the filter conditions.
	//
	// example:
	//
	// 2
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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
	// The instance visibility.
	//
	// 	- PUBLIC: TensorBoard instances are visible to all members in the workspace.
	//
	// 	- PRIVATE: TensorBoard instances are visible only to you and the administrator of the workspace.
	//
	// example:
	//
	// PRIVATE
	Accessibility *string `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	// The TensorBoard instance name.
	//
	// example:
	//
	// TestTensorboard
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The end time of the query. Use the UTC time when the TensorBoard instance is created to filter data. If you leave this parameter empty, the default value is the current time.
	//
	// example:
	//
	// 2020-11-09T14:45:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The job ID used to filter TensorBoard instances. For more information about how to obtain the ID of a job, see [ListJobs](https://help.aliyun.com/document_detail/459676.html).
	//
	// example:
	//
	// dlc-xxx
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The sorting order.
	//
	// 	- desc
	//
	// 	- asc
	//
	// example:
	//
	// desc
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// The page number. Minimum value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of TensorBoard instances per page.
	//
	// example:
	//
	// 50
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The billing method of TensorBoard instances.
	//
	// 	- Free: the TensorBoard instance that uses free resources.
	//
	// 	- Postpaid: the TensorBoard instance that uses pay-as-you-go resources.
	//
	// example:
	//
	// Postpaid
	PaymentType *string `json:"PaymentType,omitempty" xml:"PaymentType,omitempty"`
	// The resource quota ID.
	//
	// >
	//
	// 	- Only whitelisted users can use resource quotas to create TensorBoard instances. If you want to use this feature, contact us.
	//
	// 	- This parameter takes effect only when TensorBoard instances use resource quotas.
	//
	// example:
	//
	// quota12***
	QuotaId *string `json:"QuotaId,omitempty" xml:"QuotaId,omitempty"`
	// Specifies whether to return only the TensorBoard instances created by the current logon account.
	//
	// example:
	//
	// false
	ShowOwn *bool `json:"ShowOwn,omitempty" xml:"ShowOwn,omitempty"`
	// The returned field used to sort TensorBoard instances.
	//
	// 	- DisplayName: the name of the TensorBoard instance.
	//
	// 	- GmtCreateTime: the time when the TensorBoard instance is created.
	//
	// example:
	//
	// GmtCreateTime
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// The data source ID. For more information about how to obtain the ID of a job, see [ListJobs](https://help.aliyun.com/document_detail/459676.html).
	//
	// example:
	//
	// dlc-xxxxxx
	SourceId *string `json:"SourceId,omitempty" xml:"SourceId,omitempty"`
	// The data source associated with the TensorBoard instance. This parameter is no longer used. Only Deep Learning Containers (DLC) training jobs are supported.
	//
	// example:
	//
	// job
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// The start time of the query. Use the UTC time when the TensorBoard instance is created to filter data. If you leave this parameter empty, the default value is seven days before the current time.
	//
	// example:
	//
	// 2020-11-08T16:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The TensorBoard instance status. Valid values:
	//
	// 	- Creating
	//
	// 	- Running
	//
	// 	- Stopped
	//
	// 	- Succeeded
	//
	// 	- Failed
	//
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The TensorBoard instance ID used to filter TensorBoard instances.
	//
	// example:
	//
	// tensorboard-xxx
	TensorboardId *string `json:"TensorboardId,omitempty" xml:"TensorboardId,omitempty"`
	// The user ID.
	//
	// example:
	//
	// 161****3000
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The username.
	//
	// example:
	//
	// she****mo
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
	// Specifies whether to return the information about the TensorBoard instance.
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Verbose *bool `json:"Verbose,omitempty" xml:"Verbose,omitempty"`
	// The workspace ID. Obtain a list of TensorBoard instances based on the workspace ID.
	//
	// <props="china">For more information, see [ListWorkspaces](https://help.aliyun.com/document_detail/449124.html).
	//
	// example:
	//
	// 380
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListTensorboardsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTensorboardsRequest) GoString() string {
	return s.String()
}

func (s *ListTensorboardsRequest) SetAccessibility(v string) *ListTensorboardsRequest {
	s.Accessibility = &v
	return s
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

func (s *ListTensorboardsRequest) SetPaymentType(v string) *ListTensorboardsRequest {
	s.PaymentType = &v
	return s
}

func (s *ListTensorboardsRequest) SetQuotaId(v string) *ListTensorboardsRequest {
	s.QuotaId = &v
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

func (s *ListTensorboardsRequest) SetUserId(v string) *ListTensorboardsRequest {
	s.UserId = &v
	return s
}

func (s *ListTensorboardsRequest) SetUsername(v string) *ListTensorboardsRequest {
	s.Username = &v
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
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The TensorBoard instances.
	Tensorboards []*Tensorboard `json:"Tensorboards,omitempty" xml:"Tensorboards,omitempty" type:"Repeated"`
	// The total number of data sources that meet the conditions.
	//
	// example:
	//
	// 100
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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
	// The workspace ID.
	//
	// example:
	//
	// 380
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
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The TensorBoard instance ID.
	//
	// example:
	//
	// tensorboard-20210114104214-vf9lowjt3pso
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
	// The job ID.
	//
	// example:
	//
	// dlc-20210126170216-xxxxxxx
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The request ID. You can troubleshoot issues based on the request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-xxxxxx
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
	// The workspace ID.
	//
	// <props="china">For more information about how to query the workspace ID, see [ListWorkspaces](https://help.aliyun.com/document_detail/449124.html).
	//
	// example:
	//
	// 380
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
	// The ID of the request.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the TensorBoard instance.
	//
	// example:
	//
	// tensorboard-20210114104214-xxxxxxxx
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
	// The job visibility. Valid values:
	//
	// 	- PUBLIC: The job is visible to all members in the workspace.
	//
	// 	- PRIVATE: The job is visible only to you and the administrator of the workspace.
	//
	// example:
	//
	// PRIVATE
	Accessibility *string `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	// The job priority. Valid values: 1 to 9.
	//
	// 	- 1: the lowest priority.
	//
	// 	- 9: the highest priority.
	//
	// example:
	//
	// 5
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
}

func (s UpdateJobRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateJobRequest) GoString() string {
	return s.String()
}

func (s *UpdateJobRequest) SetAccessibility(v string) *UpdateJobRequest {
	s.Accessibility = &v
	return s
}

func (s *UpdateJobRequest) SetPriority(v int32) *UpdateJobRequest {
	s.Priority = &v
	return s
}

type UpdateJobResponseBody struct {
	// The job ID.
	//
	// example:
	//
	// dlc*************
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The request ID, which is used for diagnostics and Q\\&A.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
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
	// The visibility of the jobs. Valid values:
	//
	// 	- PUBLIC: The jobs are public in the workspace.
	//
	// 	- PRIVATE: The jobs are visible only to you and the administrator of the workspace.
	//
	// example:
	//
	// PRIVATE
	Accessibility *string `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	// The maximum running time. Unit: minutes.
	//
	// example:
	//
	// 300
	MaxRunningTimeMinutes *int64  `json:"MaxRunningTimeMinutes,omitempty" xml:"MaxRunningTimeMinutes,omitempty"`
	Priority              *string `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The workspace ID.
	//
	// <props="china">For more information about how to query the workspace ID, see [ListWorkspaces](https://help.aliyun.com/document_detail/449124.html).
	//
	// example:
	//
	// 380
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s UpdateTensorboardRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateTensorboardRequest) GoString() string {
	return s.String()
}

func (s *UpdateTensorboardRequest) SetAccessibility(v string) *UpdateTensorboardRequest {
	s.Accessibility = &v
	return s
}

func (s *UpdateTensorboardRequest) SetMaxRunningTimeMinutes(v int64) *UpdateTensorboardRequest {
	s.MaxRunningTimeMinutes = &v
	return s
}

func (s *UpdateTensorboardRequest) SetPriority(v string) *UpdateTensorboardRequest {
	s.Priority = &v
	return s
}

func (s *UpdateTensorboardRequest) SetWorkspaceId(v string) *UpdateTensorboardRequest {
	s.WorkspaceId = &v
	return s
}

type UpdateTensorboardResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the TensorBoard instance.
	//
	// example:
	//
	// tensorboard-20210114104214-xxxxxxxx
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
	client.EndpointRule = tea.String("regional")
	client.EndpointMap = map[string]*string{
		"ap-northeast-2-pop":          tea.String("pai-dlc.aliyuncs.com"),
		"ap-south-1":                  tea.String("pai-dlc.aliyuncs.com"),
		"ap-southeast-2":              tea.String("pai-dlc.aliyuncs.com"),
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

// Summary:
//
// Creates a job that runs in a cluster. You can configure the data source, code source, startup command, and computing resources of each node on which a job runs.
//
// Description:
//
// Before you call this operation, make sure that you understand the billing methods and [pricing](https://help.aliyun.com/document_detail/171758.html) of Deep Learning Containers (DLC) of Platform for AI (PAI).
//
// @param request - CreateJobRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateJobResponse
func (client *Client) CreateJobWithOptions(request *CreateJobRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Accessibility)) {
		body["Accessibility"] = request.Accessibility
	}

	if !tea.BoolValue(util.IsUnset(request.CodeSource)) {
		body["CodeSource"] = request.CodeSource
	}

	if !tea.BoolValue(util.IsUnset(request.CredentialConfig)) {
		body["CredentialConfig"] = request.CredentialConfig
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

// Summary:
//
// Creates a job that runs in a cluster. You can configure the data source, code source, startup command, and computing resources of each node on which a job runs.
//
// Description:
//
// Before you call this operation, make sure that you understand the billing methods and [pricing](https://help.aliyun.com/document_detail/171758.html) of Deep Learning Containers (DLC) of Platform for AI (PAI).
//
// @param request - CreateJobRequest
//
// @return CreateJobResponse
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

// Summary:
//
// Creates a TensorBoard by using a job or specifying a data source configuration.
//
// @param request - CreateTensorboardRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTensorboardResponse
func (client *Client) CreateTensorboardWithOptions(request *CreateTensorboardRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateTensorboardResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Accessibility)) {
		body["Accessibility"] = request.Accessibility
	}

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

	if !tea.BoolValue(util.IsUnset(request.Priority)) {
		body["Priority"] = request.Priority
	}

	if !tea.BoolValue(util.IsUnset(request.QuotaId)) {
		body["QuotaId"] = request.QuotaId
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

	if !tea.BoolValue(util.IsUnset(request.TensorboardDataSources)) {
		body["TensorboardDataSources"] = request.TensorboardDataSources
	}

	if !tea.BoolValue(util.IsUnset(request.TensorboardSpec)) {
		body["TensorboardSpec"] = request.TensorboardSpec
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

// Summary:
//
// Creates a TensorBoard by using a job or specifying a data source configuration.
//
// @param request - CreateTensorboardRequest
//
// @return CreateTensorboardResponse
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

// Summary:
//
// Deletes a completed or stopped job.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteJobResponse
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

// Summary:
//
// Deletes a completed or stopped job.
//
// @return DeleteJobResponse
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

// Summary:
//
// Deletes a stopped TensorBoard.
//
// @param request - DeleteTensorboardRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteTensorboardResponse
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

// Summary:
//
// Deletes a stopped TensorBoard.
//
// @param request - DeleteTensorboardRequest
//
// @return DeleteTensorboardResponse
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

// Summary:
//
// Obtains the configuration and runtime information of a job.
//
// @param request - GetJobRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetJobResponse
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

// Summary:
//
// Obtains the configuration and runtime information of a job.
//
// @param request - GetJobRequest
//
// @return GetJobResponse
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

// Summary:
//
// Obtains the system events of a job.
//
// @param request - GetJobEventsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetJobEventsResponse
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

// Summary:
//
// Obtains the system events of a job.
//
// @param request - GetJobEventsRequest
//
// @return GetJobEventsResponse
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

// Summary:
//
// Obtains the monitoring data of a job, including the CPU, GPU, and memory utilization, network, and disk read/write rate.
//
// @param request - GetJobMetricsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetJobMetricsResponse
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

// Summary:
//
// Obtains the monitoring data of a job, including the CPU, GPU, and memory utilization, network, and disk read/write rate.
//
// @param request - GetJobMetricsRequest
//
// @return GetJobMetricsResponse
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

// Summary:
//
// Obtains specified job sanity check result in a Deep Learning Containers (DLC) job.
//
// @param request - GetJobSanityCheckResultRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetJobSanityCheckResultResponse
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

// Summary:
//
// Obtains specified job sanity check result in a Deep Learning Containers (DLC) job.
//
// @param request - GetJobSanityCheckResultRequest
//
// @return GetJobSanityCheckResultResponse
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

// Summary:
//
// Obtains the system events of a specific node in a job to locate and troubleshoot issues.
//
// @param request - GetPodEventsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPodEventsResponse
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

// Summary:
//
// Obtains the system events of a specific node in a job to locate and troubleshoot issues.
//
// @param request - GetPodEventsRequest
//
// @return GetPodEventsResponse
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

// Summary:
//
// Obtains or downloads the logs of a node for a task. The logs are from the stdout and stderr of the system and user scripts.
//
// @param request - GetPodLogsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPodLogsResponse
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

// Summary:
//
// Obtains or downloads the logs of a node for a task. The logs are from the stdout and stderr of the system and user scripts.
//
// @param request - GetPodLogsRequest
//
// @return GetPodLogsResponse
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

// Summary:
//
// Obtains a Ray Dashboard URL.
//
// Description:
//
// Before you call this operation, make sure that you understand the billing methods and [pricing](https://help.aliyun.com/document_detail/171758.html) of Deep Learning Containers (DLC) of Platform for AI (PAI).
//
// @param request - GetRayDashboardRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRayDashboardResponse
func (client *Client) GetRayDashboardWithOptions(jobId *string, request *GetRayDashboardRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetRayDashboardResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IsShared)) {
		query["isShared"] = request.IsShared
	}

	if !tea.BoolValue(util.IsUnset(request.Token)) {
		query["token"] = request.Token
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetRayDashboard"),
		Version:     tea.String("2020-12-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/jobs/" + tea.StringValue(openapiutil.GetEncodeParam(jobId)) + "/rayDashboard"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetRayDashboardResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains a Ray Dashboard URL.
//
// Description:
//
// Before you call this operation, make sure that you understand the billing methods and [pricing](https://help.aliyun.com/document_detail/171758.html) of Deep Learning Containers (DLC) of Platform for AI (PAI).
//
// @param request - GetRayDashboardRequest
//
// @return GetRayDashboardResponse
func (client *Client) GetRayDashboard(jobId *string, request *GetRayDashboardRequest) (_result *GetRayDashboardResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetRayDashboardResponse{}
	_body, _err := client.GetRayDashboardWithOptions(jobId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information of a TensorBoard instance.
//
// @param request - GetTensorboardRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTensorboardResponse
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

// Summary:
//
// Queries the information of a TensorBoard instance.
//
// @param request - GetTensorboardRequest
//
// @return GetTensorboardResponse
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

// Summary:
//
// Obtains the shareable link of a TensorBoard task. The link contains digital tokens. You can use a shareable link to access a TensorBoard task.
//
// @param request - GetTensorboardSharedUrlRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTensorboardSharedUrlResponse
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

// Summary:
//
// Obtains the shareable link of a TensorBoard task. The link contains digital tokens. You can use a shareable link to access a TensorBoard task.
//
// @param request - GetTensorboardSharedUrlRequest
//
// @return GetTensorboardSharedUrlResponse
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

// Summary:
//
// Obtains the sharing token of a DLC job. This token is used to view the information about the shared job.
//
// @param request - GetTokenRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTokenResponse
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

// Summary:
//
// Obtains the sharing token of a DLC job. This token is used to view the information about the shared job.
//
// @param request - GetTokenRequest
//
// @return GetTokenResponse
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

// Summary:
//
// Provides methods and steps to obtain a HTTP link for accessing a container.
//
// @param request - GetWebTerminalRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetWebTerminalResponse
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

// Summary:
//
// Provides methods and steps to obtain a HTTP link for accessing a container.
//
// @param request - GetWebTerminalRequest
//
// @return GetWebTerminalResponse
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

// Summary:
//
// Queries the list of supported instance types.
//
// @param request - ListEcsSpecsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListEcsSpecsResponse
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

// Summary:
//
// Queries the list of supported instance types.
//
// @param request - ListEcsSpecsRequest
//
// @return ListEcsSpecsResponse
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

// Summary:
//
// Obtains the results of all sanity checks for a DLC job.
//
// @param request - ListJobSanityCheckResultsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListJobSanityCheckResultsResponse
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

// Summary:
//
// Obtains the results of all sanity checks for a DLC job.
//
// @param request - ListJobSanityCheckResultsRequest
//
// @return ListJobSanityCheckResultsResponse
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

// Summary:
//
// Queries a list of jobs and supports pagination, sorting, and filtering by conditions.
//
// @param tmpReq - ListJobsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListJobsResponse
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
	if !tea.BoolValue(util.IsUnset(request.Accessibility)) {
		query["Accessibility"] = request.Accessibility
	}

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

	if !tea.BoolValue(util.IsUnset(request.JobIds)) {
		query["JobIds"] = request.JobIds
	}

	if !tea.BoolValue(util.IsUnset(request.JobType)) {
		query["JobType"] = request.JobType
	}

	if !tea.BoolValue(util.IsUnset(request.Order)) {
		query["Order"] = request.Order
	}

	if !tea.BoolValue(util.IsUnset(request.OversoldInfo)) {
		query["OversoldInfo"] = request.OversoldInfo
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.PaymentType)) {
		query["PaymentType"] = request.PaymentType
	}

	if !tea.BoolValue(util.IsUnset(request.PipelineId)) {
		query["PipelineId"] = request.PipelineId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceQuotaName)) {
		query["ResourceQuotaName"] = request.ResourceQuotaName
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

// Summary:
//
// Queries a list of jobs and supports pagination, sorting, and filtering by conditions.
//
// @param request - ListJobsRequest
//
// @return ListJobsResponse
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

// Summary:
//
// Queries a list of TensorBoard instances.
//
// @param request - ListTensorboardsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTensorboardsResponse
func (client *Client) ListTensorboardsWithOptions(request *ListTensorboardsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListTensorboardsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Accessibility)) {
		query["Accessibility"] = request.Accessibility
	}

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

	if !tea.BoolValue(util.IsUnset(request.PaymentType)) {
		query["PaymentType"] = request.PaymentType
	}

	if !tea.BoolValue(util.IsUnset(request.QuotaId)) {
		query["QuotaId"] = request.QuotaId
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

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["UserId"] = request.UserId
	}

	if !tea.BoolValue(util.IsUnset(request.Username)) {
		query["Username"] = request.Username
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

// Summary:
//
// Queries a list of TensorBoard instances.
//
// @param request - ListTensorboardsRequest
//
// @return ListTensorboardsResponse
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

// Summary:
//
// Starts a TensorBoard instance.
//
// @param request - StartTensorboardRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartTensorboardResponse
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

// Summary:
//
// Starts a TensorBoard instance.
//
// @param request - StartTensorboardRequest
//
// @return StartTensorboardResponse
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

// Summary:
//
// Stops a running job.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopJobResponse
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

// Summary:
//
// Stops a running job.
//
// @return StopJobResponse
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

// Summary:
//
// Stops a TensorBoard instance.
//
// @param request - StopTensorboardRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopTensorboardResponse
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

// Summary:
//
// Stops a TensorBoard instance.
//
// @param request - StopTensorboardRequest
//
// @return StopTensorboardResponse
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

// Summary:
//
// Updates the configuration information of a job. For example, you can modify the priority of a job in a queue.
//
// @param request - UpdateJobRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateJobResponse
func (client *Client) UpdateJobWithOptions(JobId *string, request *UpdateJobRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Accessibility)) {
		body["Accessibility"] = request.Accessibility
	}

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

// Summary:
//
// Updates the configuration information of a job. For example, you can modify the priority of a job in a queue.
//
// @param request - UpdateJobRequest
//
// @return UpdateJobResponse
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

// Summary:
//
// Updates a TensorBoard instance.
//
// @param request - UpdateTensorboardRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateTensorboardResponse
func (client *Client) UpdateTensorboardWithOptions(TensorboardId *string, request *UpdateTensorboardRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateTensorboardResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Accessibility)) {
		query["Accessibility"] = request.Accessibility
	}

	if !tea.BoolValue(util.IsUnset(request.MaxRunningTimeMinutes)) {
		query["MaxRunningTimeMinutes"] = request.MaxRunningTimeMinutes
	}

	if !tea.BoolValue(util.IsUnset(request.Priority)) {
		query["Priority"] = request.Priority
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

// Summary:
//
// Updates a TensorBoard instance.
//
// @param request - UpdateTensorboardRequest
//
// @return UpdateTensorboardResponse
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
