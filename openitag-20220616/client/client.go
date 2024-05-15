// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type CreateTaskDetail struct {
	Admins          *CreateTaskDetailAdmins `json:"Admins,omitempty" xml:"Admins,omitempty" type:"Struct"`
	AllowAppendData *bool                   `json:"AllowAppendData,omitempty" xml:"AllowAppendData,omitempty"`
	// This parameter is required.
	AssignConfig *TaskAssginConfig `json:"AssignConfig,omitempty" xml:"AssignConfig,omitempty"`
	// This parameter is required.
	DatasetProxyRelations []*DatasetProxyConfig  `json:"DatasetProxyRelations,omitempty" xml:"DatasetProxyRelations,omitempty" type:"Repeated"`
	Exif                  map[string]interface{} `json:"Exif,omitempty" xml:"Exif,omitempty"`
	Tags                  []*string              `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// This parameter is required.
	TaskName           *string             `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	TaskTemplateConfig *TaskTemplateConfig `json:"TaskTemplateConfig,omitempty" xml:"TaskTemplateConfig,omitempty"`
	// This parameter is required.
	TaskWorkflow []*CreateTaskDetailTaskWorkflow `json:"TaskWorkflow,omitempty" xml:"TaskWorkflow,omitempty" type:"Repeated"`
	// This parameter is required.
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// This parameter is required.
	UUID        *string                              `json:"UUID,omitempty" xml:"UUID,omitempty"`
	VoteConfigs map[string]*CreateTaskDetailVoteInfo `json:"VoteConfigs,omitempty" xml:"VoteConfigs,omitempty"`
}

func (s CreateTaskDetail) String() string {
	return tea.Prettify(s)
}

func (s CreateTaskDetail) GoString() string {
	return s.String()
}

func (s *CreateTaskDetail) SetAdmins(v *CreateTaskDetailAdmins) *CreateTaskDetail {
	s.Admins = v
	return s
}

func (s *CreateTaskDetail) SetAllowAppendData(v bool) *CreateTaskDetail {
	s.AllowAppendData = &v
	return s
}

func (s *CreateTaskDetail) SetAssignConfig(v *TaskAssginConfig) *CreateTaskDetail {
	s.AssignConfig = v
	return s
}

func (s *CreateTaskDetail) SetDatasetProxyRelations(v []*DatasetProxyConfig) *CreateTaskDetail {
	s.DatasetProxyRelations = v
	return s
}

func (s *CreateTaskDetail) SetExif(v map[string]interface{}) *CreateTaskDetail {
	s.Exif = v
	return s
}

func (s *CreateTaskDetail) SetTags(v []*string) *CreateTaskDetail {
	s.Tags = v
	return s
}

func (s *CreateTaskDetail) SetTaskName(v string) *CreateTaskDetail {
	s.TaskName = &v
	return s
}

func (s *CreateTaskDetail) SetTaskTemplateConfig(v *TaskTemplateConfig) *CreateTaskDetail {
	s.TaskTemplateConfig = v
	return s
}

func (s *CreateTaskDetail) SetTaskWorkflow(v []*CreateTaskDetailTaskWorkflow) *CreateTaskDetail {
	s.TaskWorkflow = v
	return s
}

func (s *CreateTaskDetail) SetTemplateId(v string) *CreateTaskDetail {
	s.TemplateId = &v
	return s
}

func (s *CreateTaskDetail) SetUUID(v string) *CreateTaskDetail {
	s.UUID = &v
	return s
}

func (s *CreateTaskDetail) SetVoteConfigs(v map[string]*CreateTaskDetailVoteInfo) *CreateTaskDetail {
	s.VoteConfigs = v
	return s
}

type CreateTaskDetailAdmins struct {
	Users []*SimpleUser `json:"Users,omitempty" xml:"Users,omitempty" type:"Repeated"`
}

func (s CreateTaskDetailAdmins) String() string {
	return tea.Prettify(s)
}

func (s CreateTaskDetailAdmins) GoString() string {
	return s.String()
}

func (s *CreateTaskDetailAdmins) SetUsers(v []*SimpleUser) *CreateTaskDetailAdmins {
	s.Users = v
	return s
}

type CreateTaskDetailTaskWorkflow struct {
	NodeName *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
}

func (s CreateTaskDetailTaskWorkflow) String() string {
	return tea.Prettify(s)
}

func (s CreateTaskDetailTaskWorkflow) GoString() string {
	return s.String()
}

func (s *CreateTaskDetailTaskWorkflow) SetNodeName(v string) *CreateTaskDetailTaskWorkflow {
	s.NodeName = &v
	return s
}

type CreateTaskDetailVoteInfo struct {
	// example:
	//
	// 3
	MinVote *int64 `json:"MinVote,omitempty" xml:"MinVote,omitempty"`
	// example:
	//
	// 3
	VoteNum *int64 `json:"VoteNum,omitempty" xml:"VoteNum,omitempty"`
}

func (s CreateTaskDetailVoteInfo) String() string {
	return tea.Prettify(s)
}

func (s CreateTaskDetailVoteInfo) GoString() string {
	return s.String()
}

func (s *CreateTaskDetailVoteInfo) SetMinVote(v int64) *CreateTaskDetailVoteInfo {
	s.MinVote = &v
	return s
}

func (s *CreateTaskDetailVoteInfo) SetVoteNum(v int64) *CreateTaskDetailVoteInfo {
	s.VoteNum = &v
	return s
}

type DatasetProxyConfig struct {
	// example:
	//
	// LABEL
	DatasetType *string `json:"DatasetType,omitempty" xml:"DatasetType,omitempty"`
	// example:
	//
	// PAI
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// This parameter is required.
	SourceDatasetId *string `json:"SourceDatasetId,omitempty" xml:"SourceDatasetId,omitempty"`
}

func (s DatasetProxyConfig) String() string {
	return tea.Prettify(s)
}

func (s DatasetProxyConfig) GoString() string {
	return s.String()
}

func (s *DatasetProxyConfig) SetDatasetType(v string) *DatasetProxyConfig {
	s.DatasetType = &v
	return s
}

func (s *DatasetProxyConfig) SetSource(v string) *DatasetProxyConfig {
	s.Source = &v
	return s
}

func (s *DatasetProxyConfig) SetSourceDatasetId(v string) *DatasetProxyConfig {
	s.SourceDatasetId = &v
	return s
}

type FlowJobInfo struct {
	Display     *bool   `json:"Display,omitempty" xml:"Display,omitempty"`
	JobId       *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	JobType     *string `json:"JobType,omitempty" xml:"JobType,omitempty"`
	MessageId   *string `json:"MessageId,omitempty" xml:"MessageId,omitempty"`
	ProcessType *string `json:"ProcessType,omitempty" xml:"ProcessType,omitempty"`
	TaskId      *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s FlowJobInfo) String() string {
	return tea.Prettify(s)
}

func (s FlowJobInfo) GoString() string {
	return s.String()
}

func (s *FlowJobInfo) SetDisplay(v bool) *FlowJobInfo {
	s.Display = &v
	return s
}

func (s *FlowJobInfo) SetJobId(v string) *FlowJobInfo {
	s.JobId = &v
	return s
}

func (s *FlowJobInfo) SetJobType(v string) *FlowJobInfo {
	s.JobType = &v
	return s
}

func (s *FlowJobInfo) SetMessageId(v string) *FlowJobInfo {
	s.MessageId = &v
	return s
}

func (s *FlowJobInfo) SetProcessType(v string) *FlowJobInfo {
	s.ProcessType = &v
	return s
}

func (s *FlowJobInfo) SetTaskId(v string) *FlowJobInfo {
	s.TaskId = &v
	return s
}

type Job struct {
	Creator         *SimpleUser   `json:"Creator,omitempty" xml:"Creator,omitempty"`
	GmtCreateTime   *string       `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	GmtModifiedTime *string       `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	JobId           *string       `json:"JobId,omitempty" xml:"JobId,omitempty"`
	JobResult       *JobJobResult `json:"JobResult,omitempty" xml:"JobResult,omitempty" type:"Struct"`
	JobType         *string       `json:"JobType,omitempty" xml:"JobType,omitempty"`
	Status          *string       `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s Job) String() string {
	return tea.Prettify(s)
}

func (s Job) GoString() string {
	return s.String()
}

func (s *Job) SetCreator(v *SimpleUser) *Job {
	s.Creator = v
	return s
}

func (s *Job) SetGmtCreateTime(v string) *Job {
	s.GmtCreateTime = &v
	return s
}

func (s *Job) SetGmtModifiedTime(v string) *Job {
	s.GmtModifiedTime = &v
	return s
}

func (s *Job) SetJobId(v string) *Job {
	s.JobId = &v
	return s
}

func (s *Job) SetJobResult(v *JobJobResult) *Job {
	s.JobResult = v
	return s
}

func (s *Job) SetJobType(v string) *Job {
	s.JobType = &v
	return s
}

func (s *Job) SetStatus(v string) *Job {
	s.Status = &v
	return s
}

type JobJobResult struct {
	ResultLink *string `json:"ResultLink,omitempty" xml:"ResultLink,omitempty"`
}

func (s JobJobResult) String() string {
	return tea.Prettify(s)
}

func (s JobJobResult) GoString() string {
	return s.String()
}

func (s *JobJobResult) SetResultLink(v string) *JobJobResult {
	s.ResultLink = &v
	return s
}

type MarkResult struct {
	// example:
	//
	// False
	IsNeedVoteJudge *bool `json:"IsNeedVoteJudge,omitempty" xml:"IsNeedVoteJudge,omitempty"`
	// example:
	//
	// b
	MarkResult *string `json:"MarkResult,omitempty" xml:"MarkResult,omitempty"`
	// example:
	//
	// 1500758849089597440
	MarkResultId *string `json:"MarkResultId,omitempty" xml:"MarkResultId,omitempty"`
	// example:
	//
	// Mon Mar 07 17:02:48 CST 2022
	MarkTime *string `json:"MarkTime,omitempty" xml:"MarkTime,omitempty"`
	// example:
	//
	// 单选
	MarkTitle *string `json:"MarkTitle,omitempty" xml:"MarkTitle,omitempty"`
	// example:
	//
	// None
	Progress *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// example:
	//
	// 1
	QuestionId *string `json:"QuestionId,omitempty" xml:"QuestionId,omitempty"`
	// example:
	//
	// RADIO
	ResultType *string `json:"ResultType,omitempty" xml:"ResultType,omitempty"`
	// example:
	//
	// 1500758849358032896
	UserMarkResultId *string `json:"UserMarkResultId,omitempty" xml:"UserMarkResultId,omitempty"`
	// example:
	//
	// 1646643768468
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s MarkResult) String() string {
	return tea.Prettify(s)
}

func (s MarkResult) GoString() string {
	return s.String()
}

func (s *MarkResult) SetIsNeedVoteJudge(v bool) *MarkResult {
	s.IsNeedVoteJudge = &v
	return s
}

func (s *MarkResult) SetMarkResult(v string) *MarkResult {
	s.MarkResult = &v
	return s
}

func (s *MarkResult) SetMarkResultId(v string) *MarkResult {
	s.MarkResultId = &v
	return s
}

func (s *MarkResult) SetMarkTime(v string) *MarkResult {
	s.MarkTime = &v
	return s
}

func (s *MarkResult) SetMarkTitle(v string) *MarkResult {
	s.MarkTitle = &v
	return s
}

func (s *MarkResult) SetProgress(v string) *MarkResult {
	s.Progress = &v
	return s
}

func (s *MarkResult) SetQuestionId(v string) *MarkResult {
	s.QuestionId = &v
	return s
}

func (s *MarkResult) SetResultType(v string) *MarkResult {
	s.ResultType = &v
	return s
}

func (s *MarkResult) SetUserMarkResultId(v string) *MarkResult {
	s.UserMarkResultId = &v
	return s
}

func (s *MarkResult) SetVersion(v string) *MarkResult {
	s.Version = &v
	return s
}

type OpenDatasetProxyAppendDataRequest struct {
	DataMeta []map[string]*string `json:"DataMeta,omitempty" xml:"DataMeta,omitempty" type:"Repeated"`
	TaskId   *string              `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	TraceId  *string              `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
	UUID     *string              `json:"UUID,omitempty" xml:"UUID,omitempty"`
}

func (s OpenDatasetProxyAppendDataRequest) String() string {
	return tea.Prettify(s)
}

func (s OpenDatasetProxyAppendDataRequest) GoString() string {
	return s.String()
}

func (s *OpenDatasetProxyAppendDataRequest) SetDataMeta(v []map[string]*string) *OpenDatasetProxyAppendDataRequest {
	s.DataMeta = v
	return s
}

func (s *OpenDatasetProxyAppendDataRequest) SetTaskId(v string) *OpenDatasetProxyAppendDataRequest {
	s.TaskId = &v
	return s
}

func (s *OpenDatasetProxyAppendDataRequest) SetTraceId(v string) *OpenDatasetProxyAppendDataRequest {
	s.TraceId = &v
	return s
}

func (s *OpenDatasetProxyAppendDataRequest) SetUUID(v string) *OpenDatasetProxyAppendDataRequest {
	s.UUID = &v
	return s
}

type QuestionOption struct {
	Children []*QuestionOption `json:"Children,omitempty" xml:"Children,omitempty" type:"Repeated"`
	Color    *string           `json:"Color,omitempty" xml:"Color,omitempty"`
	// This parameter is required.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// This parameter is required.
	Label    *string `json:"Label,omitempty" xml:"Label,omitempty"`
	Remark   *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	Shortcut *string `json:"Shortcut,omitempty" xml:"Shortcut,omitempty"`
}

func (s QuestionOption) String() string {
	return tea.Prettify(s)
}

func (s QuestionOption) GoString() string {
	return s.String()
}

func (s *QuestionOption) SetChildren(v []*QuestionOption) *QuestionOption {
	s.Children = v
	return s
}

func (s *QuestionOption) SetColor(v string) *QuestionOption {
	s.Color = &v
	return s
}

func (s *QuestionOption) SetKey(v string) *QuestionOption {
	s.Key = &v
	return s
}

func (s *QuestionOption) SetLabel(v string) *QuestionOption {
	s.Label = &v
	return s
}

func (s *QuestionOption) SetRemark(v string) *QuestionOption {
	s.Remark = &v
	return s
}

func (s *QuestionOption) SetShortcut(v string) *QuestionOption {
	s.Shortcut = &v
	return s
}

type QuestionPlugin struct {
	// example:
	//
	// False
	CanSelect *bool             `json:"CanSelect,omitempty" xml:"CanSelect,omitempty"`
	Children  []*QuestionPlugin `json:"Children,omitempty" xml:"Children,omitempty" type:"Repeated"`
	// example:
	//
	// None
	DefaultResult *string `json:"DefaultResult,omitempty" xml:"DefaultResult,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// True
	Display *bool                  `json:"Display,omitempty" xml:"Display,omitempty"`
	Exif    map[string]interface{} `json:"Exif,omitempty" xml:"Exif,omitempty"`
	// example:
	//
	// None
	HotKeyMap *string `json:"HotKeyMap,omitempty" xml:"HotKeyMap,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 内部单选
	MarkTitle *string `json:"MarkTitle,omitempty" xml:"MarkTitle,omitempty"`
	// example:
	//
	// None
	MarkTitleAlias *string `json:"MarkTitleAlias,omitempty" xml:"MarkTitleAlias,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// False
	MustFill *bool `json:"MustFill,omitempty" xml:"MustFill,omitempty"`
	// This parameter is required.
	Options    []*QuestionOption `json:"Options,omitempty" xml:"Options,omitempty" type:"Repeated"`
	PreOptions []*string         `json:"PreOptions,omitempty" xml:"PreOptions,omitempty" type:"Repeated"`
	// This parameter is required.
	QuestionId *string `json:"QuestionId,omitempty" xml:"QuestionId,omitempty"`
	// example:
	//
	// None
	Rule *string `json:"Rule,omitempty" xml:"Rule,omitempty"`
	// example:
	//
	// None
	SelectGroup *string `json:"SelectGroup,omitempty" xml:"SelectGroup,omitempty"`
	// example:
	//
	// False
	Selected *bool `json:"Selected,omitempty" xml:"Selected,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// RADIO
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s QuestionPlugin) String() string {
	return tea.Prettify(s)
}

func (s QuestionPlugin) GoString() string {
	return s.String()
}

func (s *QuestionPlugin) SetCanSelect(v bool) *QuestionPlugin {
	s.CanSelect = &v
	return s
}

func (s *QuestionPlugin) SetChildren(v []*QuestionPlugin) *QuestionPlugin {
	s.Children = v
	return s
}

func (s *QuestionPlugin) SetDefaultResult(v string) *QuestionPlugin {
	s.DefaultResult = &v
	return s
}

func (s *QuestionPlugin) SetDisplay(v bool) *QuestionPlugin {
	s.Display = &v
	return s
}

func (s *QuestionPlugin) SetExif(v map[string]interface{}) *QuestionPlugin {
	s.Exif = v
	return s
}

func (s *QuestionPlugin) SetHotKeyMap(v string) *QuestionPlugin {
	s.HotKeyMap = &v
	return s
}

func (s *QuestionPlugin) SetMarkTitle(v string) *QuestionPlugin {
	s.MarkTitle = &v
	return s
}

func (s *QuestionPlugin) SetMarkTitleAlias(v string) *QuestionPlugin {
	s.MarkTitleAlias = &v
	return s
}

func (s *QuestionPlugin) SetMustFill(v bool) *QuestionPlugin {
	s.MustFill = &v
	return s
}

func (s *QuestionPlugin) SetOptions(v []*QuestionOption) *QuestionPlugin {
	s.Options = v
	return s
}

func (s *QuestionPlugin) SetPreOptions(v []*string) *QuestionPlugin {
	s.PreOptions = v
	return s
}

func (s *QuestionPlugin) SetQuestionId(v string) *QuestionPlugin {
	s.QuestionId = &v
	return s
}

func (s *QuestionPlugin) SetRule(v string) *QuestionPlugin {
	s.Rule = &v
	return s
}

func (s *QuestionPlugin) SetSelectGroup(v string) *QuestionPlugin {
	s.SelectGroup = &v
	return s
}

func (s *QuestionPlugin) SetSelected(v bool) *QuestionPlugin {
	s.Selected = &v
	return s
}

func (s *QuestionPlugin) SetType(v string) *QuestionPlugin {
	s.Type = &v
	return s
}

type SimpleSubtask struct {
	Items []*SimpleSubtaskItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// example:
	//
	// FINISHED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 1500682457270333440
	SubtaskId *int64 `json:"SubtaskId,omitempty" xml:"SubtaskId,omitempty"`
}

func (s SimpleSubtask) String() string {
	return tea.Prettify(s)
}

func (s SimpleSubtask) GoString() string {
	return s.String()
}

func (s *SimpleSubtask) SetItems(v []*SimpleSubtaskItems) *SimpleSubtask {
	s.Items = v
	return s
}

func (s *SimpleSubtask) SetStatus(v string) *SimpleSubtask {
	s.Status = &v
	return s
}

func (s *SimpleSubtask) SetSubtaskId(v int64) *SimpleSubtask {
	s.SubtaskId = &v
	return s
}

type SimpleSubtaskItems struct {
	// example:
	//
	// False
	AbandonFlag *bool `json:"AbandonFlag,omitempty" xml:"AbandonFlag,omitempty"`
	// example:
	//
	// None
	AbandonRemark *string `json:"AbandonRemark,omitempty" xml:"AbandonRemark,omitempty"`
	// example:
	//
	// 1957578084
	DataId *string `json:"DataId,omitempty" xml:"DataId,omitempty"`
	// example:
	//
	// False
	FeedbackFlag *bool `json:"FeedbackFlag,omitempty" xml:"FeedbackFlag,omitempty"`
	// example:
	//
	// None
	FeedbackRemark *string `json:"FeedbackRemark,omitempty" xml:"FeedbackRemark,omitempty"`
	// example:
	//
	// False
	FixedFlag *bool `json:"FixedFlag,omitempty" xml:"FixedFlag,omitempty"`
	// if can be null:
	// true
	ItemId *int64 `json:"ItemId,omitempty" xml:"ItemId,omitempty"`
	// example:
	//
	// 0
	Mine *int64 `json:"Mine,omitempty" xml:"Mine,omitempty"`
	// example:
	//
	// False
	RejectFlag *bool `json:"RejectFlag,omitempty" xml:"RejectFlag,omitempty"`
	// example:
	//
	// HANDLING
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// example:
	//
	// 311011
	Weight *int64 `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s SimpleSubtaskItems) String() string {
	return tea.Prettify(s)
}

func (s SimpleSubtaskItems) GoString() string {
	return s.String()
}

func (s *SimpleSubtaskItems) SetAbandonFlag(v bool) *SimpleSubtaskItems {
	s.AbandonFlag = &v
	return s
}

func (s *SimpleSubtaskItems) SetAbandonRemark(v string) *SimpleSubtaskItems {
	s.AbandonRemark = &v
	return s
}

func (s *SimpleSubtaskItems) SetDataId(v string) *SimpleSubtaskItems {
	s.DataId = &v
	return s
}

func (s *SimpleSubtaskItems) SetFeedbackFlag(v bool) *SimpleSubtaskItems {
	s.FeedbackFlag = &v
	return s
}

func (s *SimpleSubtaskItems) SetFeedbackRemark(v string) *SimpleSubtaskItems {
	s.FeedbackRemark = &v
	return s
}

func (s *SimpleSubtaskItems) SetFixedFlag(v bool) *SimpleSubtaskItems {
	s.FixedFlag = &v
	return s
}

func (s *SimpleSubtaskItems) SetItemId(v int64) *SimpleSubtaskItems {
	s.ItemId = &v
	return s
}

func (s *SimpleSubtaskItems) SetMine(v int64) *SimpleSubtaskItems {
	s.Mine = &v
	return s
}

func (s *SimpleSubtaskItems) SetRejectFlag(v bool) *SimpleSubtaskItems {
	s.RejectFlag = &v
	return s
}

func (s *SimpleSubtaskItems) SetState(v string) *SimpleSubtaskItems {
	s.State = &v
	return s
}

func (s *SimpleSubtaskItems) SetWeight(v int64) *SimpleSubtaskItems {
	s.Weight = &v
	return s
}

type SimpleTask struct {
	Archived *bool `json:"Archived,omitempty" xml:"Archived,omitempty"`
	// if can be null:
	// true
	ArchivedInfos   *string     `json:"ArchivedInfos,omitempty" xml:"ArchivedInfos,omitempty"`
	Creator         *SimpleUser `json:"Creator,omitempty" xml:"Creator,omitempty"`
	GmtCreateTime   *string     `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	GmtModifiedTime *string     `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	LabelStyle      *string     `json:"LabelStyle,omitempty" xml:"LabelStyle,omitempty"`
	Modifier        *SimpleUser `json:"Modifier,omitempty" xml:"Modifier,omitempty"`
	RefTaskId       *string     `json:"RefTaskId,omitempty" xml:"RefTaskId,omitempty"`
	// if can be null:
	// true
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	Stage  *string `json:"Stage,omitempty" xml:"Stage,omitempty"`
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// if can be null:
	// true
	Tags     []*string `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	TaskId   *string   `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	TaskName *string   `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	TaskType *string   `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	// if can be null:
	// true
	TemplateId    *string   `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	TenantId      *string   `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	UUID          *string   `json:"UUID,omitempty" xml:"UUID,omitempty"`
	WorkflowNodes []*string `json:"WorkflowNodes,omitempty" xml:"WorkflowNodes,omitempty" type:"Repeated"`
}

func (s SimpleTask) String() string {
	return tea.Prettify(s)
}

func (s SimpleTask) GoString() string {
	return s.String()
}

func (s *SimpleTask) SetArchived(v bool) *SimpleTask {
	s.Archived = &v
	return s
}

func (s *SimpleTask) SetArchivedInfos(v string) *SimpleTask {
	s.ArchivedInfos = &v
	return s
}

func (s *SimpleTask) SetCreator(v *SimpleUser) *SimpleTask {
	s.Creator = v
	return s
}

func (s *SimpleTask) SetGmtCreateTime(v string) *SimpleTask {
	s.GmtCreateTime = &v
	return s
}

func (s *SimpleTask) SetGmtModifiedTime(v string) *SimpleTask {
	s.GmtModifiedTime = &v
	return s
}

func (s *SimpleTask) SetLabelStyle(v string) *SimpleTask {
	s.LabelStyle = &v
	return s
}

func (s *SimpleTask) SetModifier(v *SimpleUser) *SimpleTask {
	s.Modifier = v
	return s
}

func (s *SimpleTask) SetRefTaskId(v string) *SimpleTask {
	s.RefTaskId = &v
	return s
}

func (s *SimpleTask) SetRemark(v string) *SimpleTask {
	s.Remark = &v
	return s
}

func (s *SimpleTask) SetStage(v string) *SimpleTask {
	s.Stage = &v
	return s
}

func (s *SimpleTask) SetStatus(v string) *SimpleTask {
	s.Status = &v
	return s
}

func (s *SimpleTask) SetTags(v []*string) *SimpleTask {
	s.Tags = v
	return s
}

func (s *SimpleTask) SetTaskId(v string) *SimpleTask {
	s.TaskId = &v
	return s
}

func (s *SimpleTask) SetTaskName(v string) *SimpleTask {
	s.TaskName = &v
	return s
}

func (s *SimpleTask) SetTaskType(v string) *SimpleTask {
	s.TaskType = &v
	return s
}

func (s *SimpleTask) SetTemplateId(v string) *SimpleTask {
	s.TemplateId = &v
	return s
}

func (s *SimpleTask) SetTenantId(v string) *SimpleTask {
	s.TenantId = &v
	return s
}

func (s *SimpleTask) SetUUID(v string) *SimpleTask {
	s.UUID = &v
	return s
}

func (s *SimpleTask) SetWorkflowNodes(v []*string) *SimpleTask {
	s.WorkflowNodes = v
	return s
}

type SimpleTemplate struct {
	// example:
	//
	// None
	AbandonReasons *string `json:"AbandonReasons,omitempty" xml:"AbandonReasons,omitempty"`
	Description    *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 2022-07-12 14:21:08
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// example:
	//
	// 2022-07-12 14:21:08
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// example:
	//
	// None
	SharedMode *string `json:"SharedMode,omitempty" xml:"SharedMode,omitempty"`
	// example:
	//
	// DRAFT
	Status *string   `json:"Status,omitempty" xml:"Status,omitempty"`
	Tags   []*string `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// example:
	//
	// 1546741431673270272
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// example:
	//
	// 图片分割组合77aa
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// example:
	//
	// GARDAW134
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// example:
	//
	// CUSTOM
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s SimpleTemplate) String() string {
	return tea.Prettify(s)
}

func (s SimpleTemplate) GoString() string {
	return s.String()
}

func (s *SimpleTemplate) SetAbandonReasons(v string) *SimpleTemplate {
	s.AbandonReasons = &v
	return s
}

func (s *SimpleTemplate) SetDescription(v string) *SimpleTemplate {
	s.Description = &v
	return s
}

func (s *SimpleTemplate) SetGmtCreateTime(v string) *SimpleTemplate {
	s.GmtCreateTime = &v
	return s
}

func (s *SimpleTemplate) SetGmtModifiedTime(v string) *SimpleTemplate {
	s.GmtModifiedTime = &v
	return s
}

func (s *SimpleTemplate) SetSharedMode(v string) *SimpleTemplate {
	s.SharedMode = &v
	return s
}

func (s *SimpleTemplate) SetStatus(v string) *SimpleTemplate {
	s.Status = &v
	return s
}

func (s *SimpleTemplate) SetTags(v []*string) *SimpleTemplate {
	s.Tags = v
	return s
}

func (s *SimpleTemplate) SetTemplateId(v string) *SimpleTemplate {
	s.TemplateId = &v
	return s
}

func (s *SimpleTemplate) SetTemplateName(v string) *SimpleTemplate {
	s.TemplateName = &v
	return s
}

func (s *SimpleTemplate) SetTenantId(v string) *SimpleTemplate {
	s.TenantId = &v
	return s
}

func (s *SimpleTemplate) SetType(v string) *SimpleTemplate {
	s.Type = &v
	return s
}

type SimpleTenant struct {
	Creator         *SimpleUser `json:"Creator,omitempty" xml:"Creator,omitempty"`
	Description     *string     `json:"Description,omitempty" xml:"Description,omitempty"`
	GmtCreateTime   *string     `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	GmtModifiedTime *string     `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	Modifier        *SimpleUser `json:"Modifier,omitempty" xml:"Modifier,omitempty"`
	Role            *string     `json:"Role,omitempty" xml:"Role,omitempty"`
	TenantId        *string     `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	TenantName      *string     `json:"TenantName,omitempty" xml:"TenantName,omitempty"`
	UUID            *string     `json:"UUID,omitempty" xml:"UUID,omitempty"`
}

func (s SimpleTenant) String() string {
	return tea.Prettify(s)
}

func (s SimpleTenant) GoString() string {
	return s.String()
}

func (s *SimpleTenant) SetCreator(v *SimpleUser) *SimpleTenant {
	s.Creator = v
	return s
}

func (s *SimpleTenant) SetDescription(v string) *SimpleTenant {
	s.Description = &v
	return s
}

func (s *SimpleTenant) SetGmtCreateTime(v string) *SimpleTenant {
	s.GmtCreateTime = &v
	return s
}

func (s *SimpleTenant) SetGmtModifiedTime(v string) *SimpleTenant {
	s.GmtModifiedTime = &v
	return s
}

func (s *SimpleTenant) SetModifier(v *SimpleUser) *SimpleTenant {
	s.Modifier = v
	return s
}

func (s *SimpleTenant) SetRole(v string) *SimpleTenant {
	s.Role = &v
	return s
}

func (s *SimpleTenant) SetTenantId(v string) *SimpleTenant {
	s.TenantId = &v
	return s
}

func (s *SimpleTenant) SetTenantName(v string) *SimpleTenant {
	s.TenantName = &v
	return s
}

func (s *SimpleTenant) SetUUID(v string) *SimpleTenant {
	s.UUID = &v
	return s
}

type SimpleUser struct {
	AccountNo *string `json:"AccountNo,omitempty" xml:"AccountNo,omitempty"`
	// example:
	//
	// BUC
	AccountType *string `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
	// if can be null:
	// true
	Role     *string `json:"Role,omitempty" xml:"Role,omitempty"`
	UserId   *int64  `json:"UserId,omitempty" xml:"UserId,omitempty"`
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s SimpleUser) String() string {
	return tea.Prettify(s)
}

func (s SimpleUser) GoString() string {
	return s.String()
}

func (s *SimpleUser) SetAccountNo(v string) *SimpleUser {
	s.AccountNo = &v
	return s
}

func (s *SimpleUser) SetAccountType(v string) *SimpleUser {
	s.AccountType = &v
	return s
}

func (s *SimpleUser) SetRole(v string) *SimpleUser {
	s.Role = &v
	return s
}

func (s *SimpleUser) SetUserId(v int64) *SimpleUser {
	s.UserId = &v
	return s
}

func (s *SimpleUser) SetUserName(v string) *SimpleUser {
	s.UserName = &v
	return s
}

type SimpleWorkforce struct {
	UserIds    []*int64 `json:"UserIds,omitempty" xml:"UserIds,omitempty" type:"Repeated"`
	WorkNodeId *int32   `json:"WorkNodeId,omitempty" xml:"WorkNodeId,omitempty"`
}

func (s SimpleWorkforce) String() string {
	return tea.Prettify(s)
}

func (s SimpleWorkforce) GoString() string {
	return s.String()
}

func (s *SimpleWorkforce) SetUserIds(v []*int64) *SimpleWorkforce {
	s.UserIds = v
	return s
}

func (s *SimpleWorkforce) SetWorkNodeId(v int32) *SimpleWorkforce {
	s.WorkNodeId = &v
	return s
}

type SingleTenant struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Status      *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TenantId    *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	TenantName  *string `json:"TenantName,omitempty" xml:"TenantName,omitempty"`
	UUID        *string `json:"UUID,omitempty" xml:"UUID,omitempty"`
}

func (s SingleTenant) String() string {
	return tea.Prettify(s)
}

func (s SingleTenant) GoString() string {
	return s.String()
}

func (s *SingleTenant) SetDescription(v string) *SingleTenant {
	s.Description = &v
	return s
}

func (s *SingleTenant) SetStatus(v string) *SingleTenant {
	s.Status = &v
	return s
}

func (s *SingleTenant) SetTenantId(v string) *SingleTenant {
	s.TenantId = &v
	return s
}

func (s *SingleTenant) SetTenantName(v string) *SingleTenant {
	s.TenantName = &v
	return s
}

func (s *SingleTenant) SetUUID(v string) *SingleTenant {
	s.UUID = &v
	return s
}

type SubtaskDetail struct {
	CanDiscard  *bool `json:"CanDiscard,omitempty" xml:"CanDiscard,omitempty"`
	CanReassign *bool `json:"CanReassign,omitempty" xml:"CanReassign,omitempty"`
	CanRelease  *bool `json:"CanRelease,omitempty" xml:"CanRelease,omitempty"`
	// example:
	//
	// MARK
	CurrentWorkNode *string `json:"CurrentWorkNode,omitempty" xml:"CurrentWorkNode,omitempty"`
	// if can be null:
	// true
	ExtConfigs *string               `json:"ExtConfigs,omitempty" xml:"ExtConfigs,omitempty"`
	Items      []*SubtaskDetailItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// example:
	//
	// FINISHED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 1500682457270333440
	SubtaskId *string `json:"SubtaskId,omitempty" xml:"SubtaskId,omitempty"`
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	Weight    *int64  `json:"Weight,omitempty" xml:"Weight,omitempty"`
	// example:
	//
	// FINISHED
	WorkNodeState *string      `json:"WorkNodeState,omitempty" xml:"WorkNodeState,omitempty"`
	Workforce     []*Workforce `json:"Workforce,omitempty" xml:"Workforce,omitempty" type:"Repeated"`
}

func (s SubtaskDetail) String() string {
	return tea.Prettify(s)
}

func (s SubtaskDetail) GoString() string {
	return s.String()
}

func (s *SubtaskDetail) SetCanDiscard(v bool) *SubtaskDetail {
	s.CanDiscard = &v
	return s
}

func (s *SubtaskDetail) SetCanReassign(v bool) *SubtaskDetail {
	s.CanReassign = &v
	return s
}

func (s *SubtaskDetail) SetCanRelease(v bool) *SubtaskDetail {
	s.CanRelease = &v
	return s
}

func (s *SubtaskDetail) SetCurrentWorkNode(v string) *SubtaskDetail {
	s.CurrentWorkNode = &v
	return s
}

func (s *SubtaskDetail) SetExtConfigs(v string) *SubtaskDetail {
	s.ExtConfigs = &v
	return s
}

func (s *SubtaskDetail) SetItems(v []*SubtaskDetailItems) *SubtaskDetail {
	s.Items = v
	return s
}

func (s *SubtaskDetail) SetStatus(v string) *SubtaskDetail {
	s.Status = &v
	return s
}

func (s *SubtaskDetail) SetSubtaskId(v string) *SubtaskDetail {
	s.SubtaskId = &v
	return s
}

func (s *SubtaskDetail) SetTaskId(v string) *SubtaskDetail {
	s.TaskId = &v
	return s
}

func (s *SubtaskDetail) SetWeight(v int64) *SubtaskDetail {
	s.Weight = &v
	return s
}

func (s *SubtaskDetail) SetWorkNodeState(v string) *SubtaskDetail {
	s.WorkNodeState = &v
	return s
}

func (s *SubtaskDetail) SetWorkforce(v []*Workforce) *SubtaskDetail {
	s.Workforce = v
	return s
}

type SubtaskDetailItems struct {
	// example:
	//
	// False
	AbandonFlag *bool `json:"AbandonFlag,omitempty" xml:"AbandonFlag,omitempty"`
	// example:
	//
	// None
	AbandonRemark *string `json:"AbandonRemark,omitempty" xml:"AbandonRemark,omitempty"`
	// example:
	//
	// 1957578084
	DataId *string `json:"DataId,omitempty" xml:"DataId,omitempty"`
	// example:
	//
	// False
	FeedbackFlag *bool `json:"FeedbackFlag,omitempty" xml:"FeedbackFlag,omitempty"`
	// example:
	//
	// None
	FeedbackRemark *string `json:"FeedbackRemark,omitempty" xml:"FeedbackRemark,omitempty"`
	// example:
	//
	// False
	FixedFlag *bool `json:"FixedFlag,omitempty" xml:"FixedFlag,omitempty"`
	// example:
	//
	// 0
	Mine *int64 `json:"Mine,omitempty" xml:"Mine,omitempty"`
	// example:
	//
	// False
	RejectFlag *bool `json:"RejectFlag,omitempty" xml:"RejectFlag,omitempty"`
	// example:
	//
	// HANDLING
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// example:
	//
	// 311011
	Weight *int64 `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s SubtaskDetailItems) String() string {
	return tea.Prettify(s)
}

func (s SubtaskDetailItems) GoString() string {
	return s.String()
}

func (s *SubtaskDetailItems) SetAbandonFlag(v bool) *SubtaskDetailItems {
	s.AbandonFlag = &v
	return s
}

func (s *SubtaskDetailItems) SetAbandonRemark(v string) *SubtaskDetailItems {
	s.AbandonRemark = &v
	return s
}

func (s *SubtaskDetailItems) SetDataId(v string) *SubtaskDetailItems {
	s.DataId = &v
	return s
}

func (s *SubtaskDetailItems) SetFeedbackFlag(v bool) *SubtaskDetailItems {
	s.FeedbackFlag = &v
	return s
}

func (s *SubtaskDetailItems) SetFeedbackRemark(v string) *SubtaskDetailItems {
	s.FeedbackRemark = &v
	return s
}

func (s *SubtaskDetailItems) SetFixedFlag(v bool) *SubtaskDetailItems {
	s.FixedFlag = &v
	return s
}

func (s *SubtaskDetailItems) SetMine(v int64) *SubtaskDetailItems {
	s.Mine = &v
	return s
}

func (s *SubtaskDetailItems) SetRejectFlag(v bool) *SubtaskDetailItems {
	s.RejectFlag = &v
	return s
}

func (s *SubtaskDetailItems) SetState(v string) *SubtaskDetailItems {
	s.State = &v
	return s
}

func (s *SubtaskDetailItems) SetWeight(v int64) *SubtaskDetailItems {
	s.Weight = &v
	return s
}

type SubtaskItemDetail struct {
	Annotations []*SubtaskItemDetailAnnotations `json:"Annotations,omitempty" xml:"Annotations,omitempty" type:"Repeated"`
	// example:
	//
	// None
	DataSource map[string]interface{} `json:"DataSource,omitempty" xml:"DataSource,omitempty"`
	// example:
	//
	// 1500758847176994816
	ItemId *int64 `json:"ItemId,omitempty" xml:"ItemId,omitempty"`
}

func (s SubtaskItemDetail) String() string {
	return tea.Prettify(s)
}

func (s SubtaskItemDetail) GoString() string {
	return s.String()
}

func (s *SubtaskItemDetail) SetAnnotations(v []*SubtaskItemDetailAnnotations) *SubtaskItemDetail {
	s.Annotations = v
	return s
}

func (s *SubtaskItemDetail) SetDataSource(v map[string]interface{}) *SubtaskItemDetail {
	s.DataSource = v
	return s
}

func (s *SubtaskItemDetail) SetItemId(v int64) *SubtaskItemDetail {
	s.ItemId = &v
	return s
}

type SubtaskItemDetailAnnotations struct {
	// example:
	//
	// False
	AbandonFlag *bool `json:"AbandonFlag,omitempty" xml:"AbandonFlag,omitempty"`
	// example:
	//
	// None
	AbandonRemark *string `json:"AbandonRemark,omitempty" xml:"AbandonRemark,omitempty"`
	// example:
	//
	// 1957578084
	DataId *string `json:"DataId,omitempty" xml:"DataId,omitempty"`
	// example:
	//
	// False
	FeedbackFlag *bool `json:"FeedbackFlag,omitempty" xml:"FeedbackFlag,omitempty"`
	// example:
	//
	// None
	FeedbackRemark *string `json:"FeedbackRemark,omitempty" xml:"FeedbackRemark,omitempty"`
	// example:
	//
	// False
	FixedFlag *bool `json:"FixedFlag,omitempty" xml:"FixedFlag,omitempty"`
	// example:
	//
	// 0
	Mine *int64 `json:"Mine,omitempty" xml:"Mine,omitempty"`
	// example:
	//
	// False
	RejectFlag *bool `json:"RejectFlag,omitempty" xml:"RejectFlag,omitempty"`
	// example:
	//
	// HANDLING
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// example:
	//
	// 311011
	Weight *int64 `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s SubtaskItemDetailAnnotations) String() string {
	return tea.Prettify(s)
}

func (s SubtaskItemDetailAnnotations) GoString() string {
	return s.String()
}

func (s *SubtaskItemDetailAnnotations) SetAbandonFlag(v bool) *SubtaskItemDetailAnnotations {
	s.AbandonFlag = &v
	return s
}

func (s *SubtaskItemDetailAnnotations) SetAbandonRemark(v string) *SubtaskItemDetailAnnotations {
	s.AbandonRemark = &v
	return s
}

func (s *SubtaskItemDetailAnnotations) SetDataId(v string) *SubtaskItemDetailAnnotations {
	s.DataId = &v
	return s
}

func (s *SubtaskItemDetailAnnotations) SetFeedbackFlag(v bool) *SubtaskItemDetailAnnotations {
	s.FeedbackFlag = &v
	return s
}

func (s *SubtaskItemDetailAnnotations) SetFeedbackRemark(v string) *SubtaskItemDetailAnnotations {
	s.FeedbackRemark = &v
	return s
}

func (s *SubtaskItemDetailAnnotations) SetFixedFlag(v bool) *SubtaskItemDetailAnnotations {
	s.FixedFlag = &v
	return s
}

func (s *SubtaskItemDetailAnnotations) SetMine(v int64) *SubtaskItemDetailAnnotations {
	s.Mine = &v
	return s
}

func (s *SubtaskItemDetailAnnotations) SetRejectFlag(v bool) *SubtaskItemDetailAnnotations {
	s.RejectFlag = &v
	return s
}

func (s *SubtaskItemDetailAnnotations) SetState(v string) *SubtaskItemDetailAnnotations {
	s.State = &v
	return s
}

func (s *SubtaskItemDetailAnnotations) SetWeight(v int64) *SubtaskItemDetailAnnotations {
	s.Weight = &v
	return s
}

type TaskAssginConfig struct {
	AssignCount        *int64  `json:"AssignCount,omitempty" xml:"AssignCount,omitempty"`
	AssignField        *string `json:"AssignField,omitempty" xml:"AssignField,omitempty"`
	AssignSubTaskCount *string `json:"AssignSubTaskCount,omitempty" xml:"AssignSubTaskCount,omitempty"`
	AssignType         *string `json:"AssignType,omitempty" xml:"AssignType,omitempty"`
}

func (s TaskAssginConfig) String() string {
	return tea.Prettify(s)
}

func (s TaskAssginConfig) GoString() string {
	return s.String()
}

func (s *TaskAssginConfig) SetAssignCount(v int64) *TaskAssginConfig {
	s.AssignCount = &v
	return s
}

func (s *TaskAssginConfig) SetAssignField(v string) *TaskAssginConfig {
	s.AssignField = &v
	return s
}

func (s *TaskAssginConfig) SetAssignSubTaskCount(v string) *TaskAssginConfig {
	s.AssignSubTaskCount = &v
	return s
}

func (s *TaskAssginConfig) SetAssignType(v string) *TaskAssginConfig {
	s.AssignType = &v
	return s
}

type TaskDetail struct {
	Admins          []*SimpleUser `json:"Admins,omitempty" xml:"Admins,omitempty" type:"Repeated"`
	AlertTime       *int64        `json:"AlertTime,omitempty" xml:"AlertTime,omitempty"`
	AllowAppendData *bool         `json:"AllowAppendData,omitempty" xml:"AllowAppendData,omitempty"`
	Archived        *bool         `json:"Archived,omitempty" xml:"Archived,omitempty"`
	// if can be null:
	// true
	ArchivedInfos         *string                            `json:"ArchivedInfos,omitempty" xml:"ArchivedInfos,omitempty"`
	AssignConfig          map[string]interface{}             `json:"AssignConfig,omitempty" xml:"AssignConfig,omitempty"`
	Creator               *SimpleUser                        `json:"Creator,omitempty" xml:"Creator,omitempty"`
	DatasetProxyRelations []*TaskDetailDatasetProxyRelations `json:"DatasetProxyRelations,omitempty" xml:"DatasetProxyRelations,omitempty" type:"Repeated"`
	Exif                  map[string]interface{}             `json:"Exif,omitempty" xml:"Exif,omitempty"`
	GmtCreateTime         *string                            `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	GmtModifiedTime       *string                            `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	LabelStyle            *string                            `json:"LabelStyle,omitempty" xml:"LabelStyle,omitempty"`
	MineConfigs           map[string]interface{}             `json:"MineConfigs,omitempty" xml:"MineConfigs,omitempty"`
	Modifier              *SimpleUser                        `json:"Modifier,omitempty" xml:"Modifier,omitempty"`
	NoticeConfig          map[string]interface{}             `json:"NoticeConfig,omitempty" xml:"NoticeConfig,omitempty"`
	PeriodConfig          map[string]interface{}             `json:"PeriodConfig,omitempty" xml:"PeriodConfig,omitempty"`
	RefTaskId             *string                            `json:"RefTaskId,omitempty" xml:"RefTaskId,omitempty"`
	RelateTaskConfig      map[string]interface{}             `json:"RelateTaskConfig,omitempty" xml:"RelateTaskConfig,omitempty"`
	// if can be null:
	// true
	Remark               *string                `json:"Remark,omitempty" xml:"Remark,omitempty"`
	ResultCallbackConfig map[string]interface{} `json:"ResultCallbackConfig,omitempty" xml:"ResultCallbackConfig,omitempty"`
	Stage                *string                `json:"Stage,omitempty" xml:"Stage,omitempty"`
	Status               *string                `json:"Status,omitempty" xml:"Status,omitempty"`
	// if can be null:
	// true
	Tags               []*string                     `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	TaskId             *string                       `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	TaskName           *string                       `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	TaskTemplateConfig *TaskDetailTaskTemplateConfig `json:"TaskTemplateConfig,omitempty" xml:"TaskTemplateConfig,omitempty" type:"Struct"`
	TaskType           *string                       `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	TaskWorkflow       []*TaskDetailTaskWorkflow     `json:"TaskWorkflow,omitempty" xml:"TaskWorkflow,omitempty" type:"Repeated"`
	// if can be null:
	// true
	TemplateId    *string                `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	TenantId      *string                `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	TenantName    *string                `json:"TenantName,omitempty" xml:"TenantName,omitempty"`
	UUID          *string                `json:"UUID,omitempty" xml:"UUID,omitempty"`
	VoteConfigs   map[string]interface{} `json:"VoteConfigs,omitempty" xml:"VoteConfigs,omitempty"`
	WorkflowNodes []*string              `json:"WorkflowNodes,omitempty" xml:"WorkflowNodes,omitempty" type:"Repeated"`
	RunMsg        *string                `json:"runMsg,omitempty" xml:"runMsg,omitempty"`
}

func (s TaskDetail) String() string {
	return tea.Prettify(s)
}

func (s TaskDetail) GoString() string {
	return s.String()
}

func (s *TaskDetail) SetAdmins(v []*SimpleUser) *TaskDetail {
	s.Admins = v
	return s
}

func (s *TaskDetail) SetAlertTime(v int64) *TaskDetail {
	s.AlertTime = &v
	return s
}

func (s *TaskDetail) SetAllowAppendData(v bool) *TaskDetail {
	s.AllowAppendData = &v
	return s
}

func (s *TaskDetail) SetArchived(v bool) *TaskDetail {
	s.Archived = &v
	return s
}

func (s *TaskDetail) SetArchivedInfos(v string) *TaskDetail {
	s.ArchivedInfos = &v
	return s
}

func (s *TaskDetail) SetAssignConfig(v map[string]interface{}) *TaskDetail {
	s.AssignConfig = v
	return s
}

func (s *TaskDetail) SetCreator(v *SimpleUser) *TaskDetail {
	s.Creator = v
	return s
}

func (s *TaskDetail) SetDatasetProxyRelations(v []*TaskDetailDatasetProxyRelations) *TaskDetail {
	s.DatasetProxyRelations = v
	return s
}

func (s *TaskDetail) SetExif(v map[string]interface{}) *TaskDetail {
	s.Exif = v
	return s
}

func (s *TaskDetail) SetGmtCreateTime(v string) *TaskDetail {
	s.GmtCreateTime = &v
	return s
}

func (s *TaskDetail) SetGmtModifiedTime(v string) *TaskDetail {
	s.GmtModifiedTime = &v
	return s
}

func (s *TaskDetail) SetLabelStyle(v string) *TaskDetail {
	s.LabelStyle = &v
	return s
}

func (s *TaskDetail) SetMineConfigs(v map[string]interface{}) *TaskDetail {
	s.MineConfigs = v
	return s
}

func (s *TaskDetail) SetModifier(v *SimpleUser) *TaskDetail {
	s.Modifier = v
	return s
}

func (s *TaskDetail) SetNoticeConfig(v map[string]interface{}) *TaskDetail {
	s.NoticeConfig = v
	return s
}

func (s *TaskDetail) SetPeriodConfig(v map[string]interface{}) *TaskDetail {
	s.PeriodConfig = v
	return s
}

func (s *TaskDetail) SetRefTaskId(v string) *TaskDetail {
	s.RefTaskId = &v
	return s
}

func (s *TaskDetail) SetRelateTaskConfig(v map[string]interface{}) *TaskDetail {
	s.RelateTaskConfig = v
	return s
}

func (s *TaskDetail) SetRemark(v string) *TaskDetail {
	s.Remark = &v
	return s
}

func (s *TaskDetail) SetResultCallbackConfig(v map[string]interface{}) *TaskDetail {
	s.ResultCallbackConfig = v
	return s
}

func (s *TaskDetail) SetStage(v string) *TaskDetail {
	s.Stage = &v
	return s
}

func (s *TaskDetail) SetStatus(v string) *TaskDetail {
	s.Status = &v
	return s
}

func (s *TaskDetail) SetTags(v []*string) *TaskDetail {
	s.Tags = v
	return s
}

func (s *TaskDetail) SetTaskId(v string) *TaskDetail {
	s.TaskId = &v
	return s
}

func (s *TaskDetail) SetTaskName(v string) *TaskDetail {
	s.TaskName = &v
	return s
}

func (s *TaskDetail) SetTaskTemplateConfig(v *TaskDetailTaskTemplateConfig) *TaskDetail {
	s.TaskTemplateConfig = v
	return s
}

func (s *TaskDetail) SetTaskType(v string) *TaskDetail {
	s.TaskType = &v
	return s
}

func (s *TaskDetail) SetTaskWorkflow(v []*TaskDetailTaskWorkflow) *TaskDetail {
	s.TaskWorkflow = v
	return s
}

func (s *TaskDetail) SetTemplateId(v string) *TaskDetail {
	s.TemplateId = &v
	return s
}

func (s *TaskDetail) SetTenantId(v string) *TaskDetail {
	s.TenantId = &v
	return s
}

func (s *TaskDetail) SetTenantName(v string) *TaskDetail {
	s.TenantName = &v
	return s
}

func (s *TaskDetail) SetUUID(v string) *TaskDetail {
	s.UUID = &v
	return s
}

func (s *TaskDetail) SetVoteConfigs(v map[string]interface{}) *TaskDetail {
	s.VoteConfigs = v
	return s
}

func (s *TaskDetail) SetWorkflowNodes(v []*string) *TaskDetail {
	s.WorkflowNodes = v
	return s
}

func (s *TaskDetail) SetRunMsg(v string) *TaskDetail {
	s.RunMsg = &v
	return s
}

type TaskDetailDatasetProxyRelations struct {
	DatasetId       *string                `json:"DatasetId,omitempty" xml:"DatasetId,omitempty"`
	DatasetType     *string                `json:"DatasetType,omitempty" xml:"DatasetType,omitempty"`
	Exif            map[string]interface{} `json:"Exif,omitempty" xml:"Exif,omitempty"`
	Source          *string                `json:"Source,omitempty" xml:"Source,omitempty"`
	SourceBizId     *string                `json:"SourceBizId,omitempty" xml:"SourceBizId,omitempty"`
	SourceDatasetId *string                `json:"SourceDatasetId,omitempty" xml:"SourceDatasetId,omitempty"`
}

func (s TaskDetailDatasetProxyRelations) String() string {
	return tea.Prettify(s)
}

func (s TaskDetailDatasetProxyRelations) GoString() string {
	return s.String()
}

func (s *TaskDetailDatasetProxyRelations) SetDatasetId(v string) *TaskDetailDatasetProxyRelations {
	s.DatasetId = &v
	return s
}

func (s *TaskDetailDatasetProxyRelations) SetDatasetType(v string) *TaskDetailDatasetProxyRelations {
	s.DatasetType = &v
	return s
}

func (s *TaskDetailDatasetProxyRelations) SetExif(v map[string]interface{}) *TaskDetailDatasetProxyRelations {
	s.Exif = v
	return s
}

func (s *TaskDetailDatasetProxyRelations) SetSource(v string) *TaskDetailDatasetProxyRelations {
	s.Source = &v
	return s
}

func (s *TaskDetailDatasetProxyRelations) SetSourceBizId(v string) *TaskDetailDatasetProxyRelations {
	s.SourceBizId = &v
	return s
}

func (s *TaskDetailDatasetProxyRelations) SetSourceDatasetId(v string) *TaskDetailDatasetProxyRelations {
	s.SourceDatasetId = &v
	return s
}

type TaskDetailTaskTemplateConfig struct {
	Exif               map[string]interface{} `json:"Exif,omitempty" xml:"Exif,omitempty"`
	ResourceKey        *string                `json:"ResourceKey,omitempty" xml:"ResourceKey,omitempty"`
	RobotConfig        map[string]interface{} `json:"RobotConfig,omitempty" xml:"RobotConfig,omitempty"`
	SelectQuestions    []*string              `json:"SelectQuestions,omitempty" xml:"SelectQuestions,omitempty" type:"Repeated"`
	TemplateOptionMap  map[string]interface{} `json:"TemplateOptionMap,omitempty" xml:"TemplateOptionMap,omitempty"`
	TemplateRelationId *string                `json:"TemplateRelationId,omitempty" xml:"TemplateRelationId,omitempty"`
}

func (s TaskDetailTaskTemplateConfig) String() string {
	return tea.Prettify(s)
}

func (s TaskDetailTaskTemplateConfig) GoString() string {
	return s.String()
}

func (s *TaskDetailTaskTemplateConfig) SetExif(v map[string]interface{}) *TaskDetailTaskTemplateConfig {
	s.Exif = v
	return s
}

func (s *TaskDetailTaskTemplateConfig) SetResourceKey(v string) *TaskDetailTaskTemplateConfig {
	s.ResourceKey = &v
	return s
}

func (s *TaskDetailTaskTemplateConfig) SetRobotConfig(v map[string]interface{}) *TaskDetailTaskTemplateConfig {
	s.RobotConfig = v
	return s
}

func (s *TaskDetailTaskTemplateConfig) SetSelectQuestions(v []*string) *TaskDetailTaskTemplateConfig {
	s.SelectQuestions = v
	return s
}

func (s *TaskDetailTaskTemplateConfig) SetTemplateOptionMap(v map[string]interface{}) *TaskDetailTaskTemplateConfig {
	s.TemplateOptionMap = v
	return s
}

func (s *TaskDetailTaskTemplateConfig) SetTemplateRelationId(v string) *TaskDetailTaskTemplateConfig {
	s.TemplateRelationId = &v
	return s
}

type TaskDetailTaskWorkflow struct {
	Exif     map[string]interface{} `json:"Exif,omitempty" xml:"Exif,omitempty"`
	Groups   []*string              `json:"Groups,omitempty" xml:"Groups,omitempty" type:"Repeated"`
	NodeName *string                `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	Users    []*SimpleUser          `json:"Users,omitempty" xml:"Users,omitempty" type:"Repeated"`
}

func (s TaskDetailTaskWorkflow) String() string {
	return tea.Prettify(s)
}

func (s TaskDetailTaskWorkflow) GoString() string {
	return s.String()
}

func (s *TaskDetailTaskWorkflow) SetExif(v map[string]interface{}) *TaskDetailTaskWorkflow {
	s.Exif = v
	return s
}

func (s *TaskDetailTaskWorkflow) SetGroups(v []*string) *TaskDetailTaskWorkflow {
	s.Groups = v
	return s
}

func (s *TaskDetailTaskWorkflow) SetNodeName(v string) *TaskDetailTaskWorkflow {
	s.NodeName = &v
	return s
}

func (s *TaskDetailTaskWorkflow) SetUsers(v []*SimpleUser) *TaskDetailTaskWorkflow {
	s.Users = v
	return s
}

type TaskStatistic struct {
	AcceptItemCount      *float32 `json:"AcceptItemCount,omitempty" xml:"AcceptItemCount,omitempty"`
	CheckAbandon         *float32 `json:"CheckAbandon,omitempty" xml:"CheckAbandon,omitempty"`
	CheckAccuracy        *float32 `json:"CheckAccuracy,omitempty" xml:"CheckAccuracy,omitempty"`
	CheckEfficiency      *float32 `json:"CheckEfficiency,omitempty" xml:"CheckEfficiency,omitempty"`
	CheckedAccuracy      *float32 `json:"CheckedAccuracy,omitempty" xml:"CheckedAccuracy,omitempty"`
	CheckedError         *float32 `json:"CheckedError,omitempty" xml:"CheckedError,omitempty"`
	CheckedRejectCount   *float32 `json:"CheckedRejectCount,omitempty" xml:"CheckedRejectCount,omitempty"`
	FinalAbandonCount    *float32 `json:"FinalAbandonCount,omitempty" xml:"FinalAbandonCount,omitempty"`
	FinishedItemCount    *int64   `json:"FinishedItemCount,omitempty" xml:"FinishedItemCount,omitempty"`
	FinishedSubtaskCount *int64   `json:"FinishedSubtaskCount,omitempty" xml:"FinishedSubtaskCount,omitempty"`
	MarkEfficiency       *float32 `json:"MarkEfficiency,omitempty" xml:"MarkEfficiency,omitempty"`
	// if can be null:
	// true
	PreMarkFixedCount  *float32 `json:"PreMarkFixedCount,omitempty" xml:"PreMarkFixedCount,omitempty"`
	SampledAccuracy    *float32 `json:"SampledAccuracy,omitempty" xml:"SampledAccuracy,omitempty"`
	SampledErrorCount  *float32 `json:"SampledErrorCount,omitempty" xml:"SampledErrorCount,omitempty"`
	SampledRejectCount *float32 `json:"SampledRejectCount,omitempty" xml:"SampledRejectCount,omitempty"`
	SamplingAccuracy   *float32 `json:"SamplingAccuracy,omitempty" xml:"SamplingAccuracy,omitempty"`
	TotalCheckCount    *float32 `json:"TotalCheckCount,omitempty" xml:"TotalCheckCount,omitempty"`
	TotalCheckTime     *float32 `json:"TotalCheckTime,omitempty" xml:"TotalCheckTime,omitempty"`
	TotalCheckedCount  *float32 `json:"TotalCheckedCount,omitempty" xml:"TotalCheckedCount,omitempty"`
	TotalItemCount     *int64   `json:"TotalItemCount,omitempty" xml:"TotalItemCount,omitempty"`
	TotalMarkTime      *float32 `json:"TotalMarkTime,omitempty" xml:"TotalMarkTime,omitempty"`
	TotalSampledCount  *float32 `json:"TotalSampledCount,omitempty" xml:"TotalSampledCount,omitempty"`
	TotalSamplingCount *float32 `json:"TotalSamplingCount,omitempty" xml:"TotalSamplingCount,omitempty"`
	TotalSubtaskCount  *int64   `json:"TotalSubtaskCount,omitempty" xml:"TotalSubtaskCount,omitempty"`
	TotalWorkTime      *float32 `json:"TotalWorkTime,omitempty" xml:"TotalWorkTime,omitempty"`
}

func (s TaskStatistic) String() string {
	return tea.Prettify(s)
}

func (s TaskStatistic) GoString() string {
	return s.String()
}

func (s *TaskStatistic) SetAcceptItemCount(v float32) *TaskStatistic {
	s.AcceptItemCount = &v
	return s
}

func (s *TaskStatistic) SetCheckAbandon(v float32) *TaskStatistic {
	s.CheckAbandon = &v
	return s
}

func (s *TaskStatistic) SetCheckAccuracy(v float32) *TaskStatistic {
	s.CheckAccuracy = &v
	return s
}

func (s *TaskStatistic) SetCheckEfficiency(v float32) *TaskStatistic {
	s.CheckEfficiency = &v
	return s
}

func (s *TaskStatistic) SetCheckedAccuracy(v float32) *TaskStatistic {
	s.CheckedAccuracy = &v
	return s
}

func (s *TaskStatistic) SetCheckedError(v float32) *TaskStatistic {
	s.CheckedError = &v
	return s
}

func (s *TaskStatistic) SetCheckedRejectCount(v float32) *TaskStatistic {
	s.CheckedRejectCount = &v
	return s
}

func (s *TaskStatistic) SetFinalAbandonCount(v float32) *TaskStatistic {
	s.FinalAbandonCount = &v
	return s
}

func (s *TaskStatistic) SetFinishedItemCount(v int64) *TaskStatistic {
	s.FinishedItemCount = &v
	return s
}

func (s *TaskStatistic) SetFinishedSubtaskCount(v int64) *TaskStatistic {
	s.FinishedSubtaskCount = &v
	return s
}

func (s *TaskStatistic) SetMarkEfficiency(v float32) *TaskStatistic {
	s.MarkEfficiency = &v
	return s
}

func (s *TaskStatistic) SetPreMarkFixedCount(v float32) *TaskStatistic {
	s.PreMarkFixedCount = &v
	return s
}

func (s *TaskStatistic) SetSampledAccuracy(v float32) *TaskStatistic {
	s.SampledAccuracy = &v
	return s
}

func (s *TaskStatistic) SetSampledErrorCount(v float32) *TaskStatistic {
	s.SampledErrorCount = &v
	return s
}

func (s *TaskStatistic) SetSampledRejectCount(v float32) *TaskStatistic {
	s.SampledRejectCount = &v
	return s
}

func (s *TaskStatistic) SetSamplingAccuracy(v float32) *TaskStatistic {
	s.SamplingAccuracy = &v
	return s
}

func (s *TaskStatistic) SetTotalCheckCount(v float32) *TaskStatistic {
	s.TotalCheckCount = &v
	return s
}

func (s *TaskStatistic) SetTotalCheckTime(v float32) *TaskStatistic {
	s.TotalCheckTime = &v
	return s
}

func (s *TaskStatistic) SetTotalCheckedCount(v float32) *TaskStatistic {
	s.TotalCheckedCount = &v
	return s
}

func (s *TaskStatistic) SetTotalItemCount(v int64) *TaskStatistic {
	s.TotalItemCount = &v
	return s
}

func (s *TaskStatistic) SetTotalMarkTime(v float32) *TaskStatistic {
	s.TotalMarkTime = &v
	return s
}

func (s *TaskStatistic) SetTotalSampledCount(v float32) *TaskStatistic {
	s.TotalSampledCount = &v
	return s
}

func (s *TaskStatistic) SetTotalSamplingCount(v float32) *TaskStatistic {
	s.TotalSamplingCount = &v
	return s
}

func (s *TaskStatistic) SetTotalSubtaskCount(v int64) *TaskStatistic {
	s.TotalSubtaskCount = &v
	return s
}

func (s *TaskStatistic) SetTotalWorkTime(v float32) *TaskStatistic {
	s.TotalWorkTime = &v
	return s
}

type TaskTemplateConfig struct {
	Exif               map[string]*string                   `json:"Exif,omitempty" xml:"Exif,omitempty"`
	ResourceKey        *string                              `json:"ResourceKey,omitempty" xml:"ResourceKey,omitempty"`
	SelectQuestions    []*string                            `json:"SelectQuestions,omitempty" xml:"SelectQuestions,omitempty" type:"Repeated"`
	TemplateOptionMap  map[string]*TaskTemplateOptionConfig `json:"TemplateOptionMap,omitempty" xml:"TemplateOptionMap,omitempty"`
	TemplateRelationId *string                              `json:"TemplateRelationId,omitempty" xml:"TemplateRelationId,omitempty"`
}

func (s TaskTemplateConfig) String() string {
	return tea.Prettify(s)
}

func (s TaskTemplateConfig) GoString() string {
	return s.String()
}

func (s *TaskTemplateConfig) SetExif(v map[string]*string) *TaskTemplateConfig {
	s.Exif = v
	return s
}

func (s *TaskTemplateConfig) SetResourceKey(v string) *TaskTemplateConfig {
	s.ResourceKey = &v
	return s
}

func (s *TaskTemplateConfig) SetSelectQuestions(v []*string) *TaskTemplateConfig {
	s.SelectQuestions = v
	return s
}

func (s *TaskTemplateConfig) SetTemplateOptionMap(v map[string]*TaskTemplateOptionConfig) *TaskTemplateConfig {
	s.TemplateOptionMap = v
	return s
}

func (s *TaskTemplateConfig) SetTemplateRelationId(v string) *TaskTemplateConfig {
	s.TemplateRelationId = &v
	return s
}

type TaskTemplateOptionConfig struct {
	// if can be null:
	// false
	DefaultResult *string `json:"DefaultResult,omitempty" xml:"DefaultResult,omitempty"`
	// if can be null:
	// false
	Options []*QuestionOption `json:"Options,omitempty" xml:"Options,omitempty" type:"Repeated"`
	// if can be null:
	// false
	PreOptions []*string `json:"PreOptions,omitempty" xml:"PreOptions,omitempty" type:"Repeated"`
	// if can be null:
	// false
	Rule *string `json:"Rule,omitempty" xml:"Rule,omitempty"`
}

func (s TaskTemplateOptionConfig) String() string {
	return tea.Prettify(s)
}

func (s TaskTemplateOptionConfig) GoString() string {
	return s.String()
}

func (s *TaskTemplateOptionConfig) SetDefaultResult(v string) *TaskTemplateOptionConfig {
	s.DefaultResult = &v
	return s
}

func (s *TaskTemplateOptionConfig) SetOptions(v []*QuestionOption) *TaskTemplateOptionConfig {
	s.Options = v
	return s
}

func (s *TaskTemplateOptionConfig) SetPreOptions(v []*string) *TaskTemplateOptionConfig {
	s.PreOptions = v
	return s
}

func (s *TaskTemplateOptionConfig) SetRule(v string) *TaskTemplateOptionConfig {
	s.Rule = &v
	return s
}

type TemplateDTO struct {
	Classify    *string                `json:"Classify,omitempty" xml:"Classify,omitempty"`
	Description *string                `json:"Description,omitempty" xml:"Description,omitempty"`
	Exif        map[string]interface{} `json:"Exif,omitempty" xml:"Exif,omitempty"`
	// This parameter is required.
	QuestionConfigs []*QuestionPlugin        `json:"QuestionConfigs,omitempty" xml:"QuestionConfigs,omitempty" type:"Repeated"`
	RobotConfigs    []map[string]interface{} `json:"RobotConfigs,omitempty" xml:"RobotConfigs,omitempty" type:"Repeated"`
	SharedMode      *string                  `json:"SharedMode,omitempty" xml:"SharedMode,omitempty"`
	Tags            []*string                `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	TemplateId      *string                  `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// This parameter is required.
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// This parameter is required.
	ViewConfigs *TemplateDTOViewConfigs `json:"ViewConfigs,omitempty" xml:"ViewConfigs,omitempty" type:"Struct"`
}

func (s TemplateDTO) String() string {
	return tea.Prettify(s)
}

func (s TemplateDTO) GoString() string {
	return s.String()
}

func (s *TemplateDTO) SetClassify(v string) *TemplateDTO {
	s.Classify = &v
	return s
}

func (s *TemplateDTO) SetDescription(v string) *TemplateDTO {
	s.Description = &v
	return s
}

func (s *TemplateDTO) SetExif(v map[string]interface{}) *TemplateDTO {
	s.Exif = v
	return s
}

func (s *TemplateDTO) SetQuestionConfigs(v []*QuestionPlugin) *TemplateDTO {
	s.QuestionConfigs = v
	return s
}

func (s *TemplateDTO) SetRobotConfigs(v []map[string]interface{}) *TemplateDTO {
	s.RobotConfigs = v
	return s
}

func (s *TemplateDTO) SetSharedMode(v string) *TemplateDTO {
	s.SharedMode = &v
	return s
}

func (s *TemplateDTO) SetTags(v []*string) *TemplateDTO {
	s.Tags = v
	return s
}

func (s *TemplateDTO) SetTemplateId(v string) *TemplateDTO {
	s.TemplateId = &v
	return s
}

func (s *TemplateDTO) SetTemplateName(v string) *TemplateDTO {
	s.TemplateName = &v
	return s
}

func (s *TemplateDTO) SetViewConfigs(v *TemplateDTOViewConfigs) *TemplateDTO {
	s.ViewConfigs = v
	return s
}

type TemplateDTOViewConfigs struct {
	ViewPlugins []*ViewPlugin `json:"ViewPlugins,omitempty" xml:"ViewPlugins,omitempty" type:"Repeated"`
}

func (s TemplateDTOViewConfigs) String() string {
	return tea.Prettify(s)
}

func (s TemplateDTOViewConfigs) GoString() string {
	return s.String()
}

func (s *TemplateDTOViewConfigs) SetViewPlugins(v []*ViewPlugin) *TemplateDTOViewConfigs {
	s.ViewPlugins = v
	return s
}

type TemplateDetail struct {
	AbandonReasons  []*string                  `json:"AbandonReasons,omitempty" xml:"AbandonReasons,omitempty" type:"Repeated"`
	Classify        *string                    `json:"Classify,omitempty" xml:"Classify,omitempty"`
	Creator         *SimpleUser                `json:"Creator,omitempty" xml:"Creator,omitempty"`
	Description     *string                    `json:"Description,omitempty" xml:"Description,omitempty"`
	Exif            map[string]interface{}     `json:"Exif,omitempty" xml:"Exif,omitempty"`
	GmtCreateTime   *string                    `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	GmtModifiedTime *string                    `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	Modifier        *SimpleUser                `json:"Modifier,omitempty" xml:"Modifier,omitempty"`
	QuestionConfigs []*QuestionPlugin          `json:"QuestionConfigs,omitempty" xml:"QuestionConfigs,omitempty" type:"Repeated"`
	SharedMode      *string                    `json:"SharedMode,omitempty" xml:"SharedMode,omitempty"`
	Status          *string                    `json:"Status,omitempty" xml:"Status,omitempty"`
	Tags            []*string                  `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	TemplateId      *string                    `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	TemplateName    *string                    `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	TenantId        *string                    `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	Type            *string                    `json:"Type,omitempty" xml:"Type,omitempty"`
	ViewConfigs     *TemplateDetailViewConfigs `json:"ViewConfigs,omitempty" xml:"ViewConfigs,omitempty" type:"Struct"`
}

func (s TemplateDetail) String() string {
	return tea.Prettify(s)
}

func (s TemplateDetail) GoString() string {
	return s.String()
}

func (s *TemplateDetail) SetAbandonReasons(v []*string) *TemplateDetail {
	s.AbandonReasons = v
	return s
}

func (s *TemplateDetail) SetClassify(v string) *TemplateDetail {
	s.Classify = &v
	return s
}

func (s *TemplateDetail) SetCreator(v *SimpleUser) *TemplateDetail {
	s.Creator = v
	return s
}

func (s *TemplateDetail) SetDescription(v string) *TemplateDetail {
	s.Description = &v
	return s
}

func (s *TemplateDetail) SetExif(v map[string]interface{}) *TemplateDetail {
	s.Exif = v
	return s
}

func (s *TemplateDetail) SetGmtCreateTime(v string) *TemplateDetail {
	s.GmtCreateTime = &v
	return s
}

func (s *TemplateDetail) SetGmtModifiedTime(v string) *TemplateDetail {
	s.GmtModifiedTime = &v
	return s
}

func (s *TemplateDetail) SetModifier(v *SimpleUser) *TemplateDetail {
	s.Modifier = v
	return s
}

func (s *TemplateDetail) SetQuestionConfigs(v []*QuestionPlugin) *TemplateDetail {
	s.QuestionConfigs = v
	return s
}

func (s *TemplateDetail) SetSharedMode(v string) *TemplateDetail {
	s.SharedMode = &v
	return s
}

func (s *TemplateDetail) SetStatus(v string) *TemplateDetail {
	s.Status = &v
	return s
}

func (s *TemplateDetail) SetTags(v []*string) *TemplateDetail {
	s.Tags = v
	return s
}

func (s *TemplateDetail) SetTemplateId(v string) *TemplateDetail {
	s.TemplateId = &v
	return s
}

func (s *TemplateDetail) SetTemplateName(v string) *TemplateDetail {
	s.TemplateName = &v
	return s
}

func (s *TemplateDetail) SetTenantId(v string) *TemplateDetail {
	s.TenantId = &v
	return s
}

func (s *TemplateDetail) SetType(v string) *TemplateDetail {
	s.Type = &v
	return s
}

func (s *TemplateDetail) SetViewConfigs(v *TemplateDetailViewConfigs) *TemplateDetail {
	s.ViewConfigs = v
	return s
}

type TemplateDetailViewConfigs struct {
	ViewPlugins []*ViewPlugin `json:"ViewPlugins,omitempty" xml:"ViewPlugins,omitempty" type:"Repeated"`
}

func (s TemplateDetailViewConfigs) String() string {
	return tea.Prettify(s)
}

func (s TemplateDetailViewConfigs) GoString() string {
	return s.String()
}

func (s *TemplateDetailViewConfigs) SetViewPlugins(v []*ViewPlugin) *TemplateDetailViewConfigs {
	s.ViewPlugins = v
	return s
}

type TemplateQuestion struct {
	Children []*TemplateQuestion    `json:"Children,omitempty" xml:"Children,omitempty" type:"Repeated"`
	Exif     map[string]interface{} `json:"Exif,omitempty" xml:"Exif,omitempty"`
	// example:
	//
	// 题目1
	MarkTitle  *string           `json:"MarkTitle,omitempty" xml:"MarkTitle,omitempty"`
	Options    []*QuestionOption `json:"Options,omitempty" xml:"Options,omitempty" type:"Repeated"`
	PreOptions []*string         `json:"PreOptions,omitempty" xml:"PreOptions,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	QuestionId *int64 `json:"QuestionId,omitempty" xml:"QuestionId,omitempty"`
	// example:
	//
	// RADIO/GROUP
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s TemplateQuestion) String() string {
	return tea.Prettify(s)
}

func (s TemplateQuestion) GoString() string {
	return s.String()
}

func (s *TemplateQuestion) SetChildren(v []*TemplateQuestion) *TemplateQuestion {
	s.Children = v
	return s
}

func (s *TemplateQuestion) SetExif(v map[string]interface{}) *TemplateQuestion {
	s.Exif = v
	return s
}

func (s *TemplateQuestion) SetMarkTitle(v string) *TemplateQuestion {
	s.MarkTitle = &v
	return s
}

func (s *TemplateQuestion) SetOptions(v []*QuestionOption) *TemplateQuestion {
	s.Options = v
	return s
}

func (s *TemplateQuestion) SetPreOptions(v []*string) *TemplateQuestion {
	s.PreOptions = v
	return s
}

func (s *TemplateQuestion) SetQuestionId(v int64) *TemplateQuestion {
	s.QuestionId = &v
	return s
}

func (s *TemplateQuestion) SetType(v string) *TemplateQuestion {
	s.Type = &v
	return s
}

type TemplateView struct {
	Fields []*TemplateViewFields `json:"Fields,omitempty" xml:"Fields,omitempty" type:"Repeated"`
}

func (s TemplateView) String() string {
	return tea.Prettify(s)
}

func (s TemplateView) GoString() string {
	return s.String()
}

func (s *TemplateView) SetFields(v []*TemplateViewFields) *TemplateView {
	s.Fields = v
	return s
}

type TemplateViewFields struct {
	// example:
	//
	// True
	DisplayOriImg *bool `json:"DisplayOriImg,omitempty" xml:"DisplayOriImg,omitempty"`
	// example:
	//
	// url
	FieldName *string `json:"FieldName,omitempty" xml:"FieldName,omitempty"`
	// example:
	//
	// IMG
	Type      *string                      `json:"Type,omitempty" xml:"Type,omitempty"`
	VisitInfo *TemplateViewFieldsVisitInfo `json:"VisitInfo,omitempty" xml:"VisitInfo,omitempty" type:"Struct"`
}

func (s TemplateViewFields) String() string {
	return tea.Prettify(s)
}

func (s TemplateViewFields) GoString() string {
	return s.String()
}

func (s *TemplateViewFields) SetDisplayOriImg(v bool) *TemplateViewFields {
	s.DisplayOriImg = &v
	return s
}

func (s *TemplateViewFields) SetFieldName(v string) *TemplateViewFields {
	s.FieldName = &v
	return s
}

func (s *TemplateViewFields) SetType(v string) *TemplateViewFields {
	s.Type = &v
	return s
}

func (s *TemplateViewFields) SetVisitInfo(v *TemplateViewFieldsVisitInfo) *TemplateViewFields {
	s.VisitInfo = v
	return s
}

type TemplateViewFieldsVisitInfo struct {
	AftsConf map[string]interface{} `json:"AftsConf,omitempty" xml:"AftsConf,omitempty"`
	OssConf  map[string]interface{} `json:"OssConf,omitempty" xml:"OssConf,omitempty"`
}

func (s TemplateViewFieldsVisitInfo) String() string {
	return tea.Prettify(s)
}

func (s TemplateViewFieldsVisitInfo) GoString() string {
	return s.String()
}

func (s *TemplateViewFieldsVisitInfo) SetAftsConf(v map[string]interface{}) *TemplateViewFieldsVisitInfo {
	s.AftsConf = v
	return s
}

func (s *TemplateViewFieldsVisitInfo) SetOssConf(v map[string]interface{}) *TemplateViewFieldsVisitInfo {
	s.OssConf = v
	return s
}

type UpdateTaskDTO struct {
	Exif     map[string]*string `json:"Exif,omitempty" xml:"Exif,omitempty"`
	Remark   *string            `json:"Remark,omitempty" xml:"Remark,omitempty"`
	Tags     []*string          `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	TaskName *string            `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
}

func (s UpdateTaskDTO) String() string {
	return tea.Prettify(s)
}

func (s UpdateTaskDTO) GoString() string {
	return s.String()
}

func (s *UpdateTaskDTO) SetExif(v map[string]*string) *UpdateTaskDTO {
	s.Exif = v
	return s
}

func (s *UpdateTaskDTO) SetRemark(v string) *UpdateTaskDTO {
	s.Remark = &v
	return s
}

func (s *UpdateTaskDTO) SetTags(v []*string) *UpdateTaskDTO {
	s.Tags = v
	return s
}

func (s *UpdateTaskDTO) SetTaskName(v string) *UpdateTaskDTO {
	s.TaskName = &v
	return s
}

type UserStatistic struct {
	AcceptedMarkItemsCount *float32 `json:"AcceptedMarkItemsCount,omitempty" xml:"AcceptedMarkItemsCount,omitempty"`
	CheckCount             *float32 `json:"CheckCount,omitempty" xml:"CheckCount,omitempty"`
	CheckedAcceptedCount   *float32 `json:"CheckedAcceptedCount,omitempty" xml:"CheckedAcceptedCount,omitempty"`
	CheckedAccuracy        *float32 `json:"CheckedAccuracy,omitempty" xml:"CheckedAccuracy,omitempty"`
	MarkEfficiency         *float32 `json:"MarkEfficiency,omitempty" xml:"MarkEfficiency,omitempty"`
	MarkTime               *float32 `json:"MarkTime,omitempty" xml:"MarkTime,omitempty"`
	SamplingAccuracy       *float32 `json:"SamplingAccuracy,omitempty" xml:"SamplingAccuracy,omitempty"`
	SamplingCount          *float32 `json:"SamplingCount,omitempty" xml:"SamplingCount,omitempty"`
	SamplingErrorCount     *float32 `json:"SamplingErrorCount,omitempty" xml:"SamplingErrorCount,omitempty"`
	TotalMarkItemsCount    *float32 `json:"TotalMarkItemsCount,omitempty" xml:"TotalMarkItemsCount,omitempty"`
	UserId                 *string  `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s UserStatistic) String() string {
	return tea.Prettify(s)
}

func (s UserStatistic) GoString() string {
	return s.String()
}

func (s *UserStatistic) SetAcceptedMarkItemsCount(v float32) *UserStatistic {
	s.AcceptedMarkItemsCount = &v
	return s
}

func (s *UserStatistic) SetCheckCount(v float32) *UserStatistic {
	s.CheckCount = &v
	return s
}

func (s *UserStatistic) SetCheckedAcceptedCount(v float32) *UserStatistic {
	s.CheckedAcceptedCount = &v
	return s
}

func (s *UserStatistic) SetCheckedAccuracy(v float32) *UserStatistic {
	s.CheckedAccuracy = &v
	return s
}

func (s *UserStatistic) SetMarkEfficiency(v float32) *UserStatistic {
	s.MarkEfficiency = &v
	return s
}

func (s *UserStatistic) SetMarkTime(v float32) *UserStatistic {
	s.MarkTime = &v
	return s
}

func (s *UserStatistic) SetSamplingAccuracy(v float32) *UserStatistic {
	s.SamplingAccuracy = &v
	return s
}

func (s *UserStatistic) SetSamplingCount(v float32) *UserStatistic {
	s.SamplingCount = &v
	return s
}

func (s *UserStatistic) SetSamplingErrorCount(v float32) *UserStatistic {
	s.SamplingErrorCount = &v
	return s
}

func (s *UserStatistic) SetTotalMarkItemsCount(v float32) *UserStatistic {
	s.TotalMarkItemsCount = &v
	return s
}

func (s *UserStatistic) SetUserId(v string) *UserStatistic {
	s.UserId = &v
	return s
}

type ViewPlugin struct {
	// This parameter is required.
	BindField *string `json:"BindField,omitempty" xml:"BindField,omitempty"`
	Convertor *string `json:"Convertor,omitempty" xml:"Convertor,omitempty"`
	// This parameter is required.
	CorsProxy *bool `json:"CorsProxy,omitempty" xml:"CorsProxy,omitempty"`
	// This parameter is required.
	DisplayOriImg       *bool                  `json:"DisplayOriImg,omitempty" xml:"DisplayOriImg,omitempty"`
	Exif                map[string]interface{} `json:"Exif,omitempty" xml:"Exif,omitempty"`
	Hide                *bool                  `json:"Hide,omitempty" xml:"Hide,omitempty"`
	Plugins             map[string]interface{} `json:"Plugins,omitempty" xml:"Plugins,omitempty"`
	RelationQuestionIds []*string              `json:"RelationQuestionIds,omitempty" xml:"RelationQuestionIds,omitempty" type:"Repeated"`
	// This parameter is required.
	Type      *string              `json:"Type,omitempty" xml:"Type,omitempty"`
	VisitInfo *ViewPluginVisitInfo `json:"VisitInfo,omitempty" xml:"VisitInfo,omitempty" type:"Struct"`
}

func (s ViewPlugin) String() string {
	return tea.Prettify(s)
}

func (s ViewPlugin) GoString() string {
	return s.String()
}

func (s *ViewPlugin) SetBindField(v string) *ViewPlugin {
	s.BindField = &v
	return s
}

func (s *ViewPlugin) SetConvertor(v string) *ViewPlugin {
	s.Convertor = &v
	return s
}

func (s *ViewPlugin) SetCorsProxy(v bool) *ViewPlugin {
	s.CorsProxy = &v
	return s
}

func (s *ViewPlugin) SetDisplayOriImg(v bool) *ViewPlugin {
	s.DisplayOriImg = &v
	return s
}

func (s *ViewPlugin) SetExif(v map[string]interface{}) *ViewPlugin {
	s.Exif = v
	return s
}

func (s *ViewPlugin) SetHide(v bool) *ViewPlugin {
	s.Hide = &v
	return s
}

func (s *ViewPlugin) SetPlugins(v map[string]interface{}) *ViewPlugin {
	s.Plugins = v
	return s
}

func (s *ViewPlugin) SetRelationQuestionIds(v []*string) *ViewPlugin {
	s.RelationQuestionIds = v
	return s
}

func (s *ViewPlugin) SetType(v string) *ViewPlugin {
	s.Type = &v
	return s
}

func (s *ViewPlugin) SetVisitInfo(v *ViewPluginVisitInfo) *ViewPlugin {
	s.VisitInfo = v
	return s
}

type ViewPluginVisitInfo struct {
	AftsConf map[string]interface{} `json:"aftsConf,omitempty" xml:"aftsConf,omitempty"`
	OssConf  map[string]interface{} `json:"ossConf,omitempty" xml:"ossConf,omitempty"`
}

func (s ViewPluginVisitInfo) String() string {
	return tea.Prettify(s)
}

func (s ViewPluginVisitInfo) GoString() string {
	return s.String()
}

func (s *ViewPluginVisitInfo) SetAftsConf(v map[string]interface{}) *ViewPluginVisitInfo {
	s.AftsConf = v
	return s
}

func (s *ViewPluginVisitInfo) SetOssConf(v map[string]interface{}) *ViewPluginVisitInfo {
	s.OssConf = v
	return s
}

type Workforce struct {
	NodeType   *string       `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
	Users      []*SimpleUser `json:"Users,omitempty" xml:"Users,omitempty" type:"Repeated"`
	WorkNodeId *int32        `json:"WorkNodeId,omitempty" xml:"WorkNodeId,omitempty"`
}

func (s Workforce) String() string {
	return tea.Prettify(s)
}

func (s Workforce) GoString() string {
	return s.String()
}

func (s *Workforce) SetNodeType(v string) *Workforce {
	s.NodeType = &v
	return s
}

func (s *Workforce) SetUsers(v []*SimpleUser) *Workforce {
	s.Users = v
	return s
}

func (s *Workforce) SetWorkNodeId(v int32) *Workforce {
	s.WorkNodeId = &v
	return s
}

type AddWorkNodeWorkforceRequest struct {
	UserIds []*int64 `json:"UserIds,omitempty" xml:"UserIds,omitempty" type:"Repeated"`
}

func (s AddWorkNodeWorkforceRequest) String() string {
	return tea.Prettify(s)
}

func (s AddWorkNodeWorkforceRequest) GoString() string {
	return s.String()
}

func (s *AddWorkNodeWorkforceRequest) SetUserIds(v []*int64) *AddWorkNodeWorkforceRequest {
	s.UserIds = v
	return s
}

type AddWorkNodeWorkforceResponseBody struct {
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// null
	Details *string `json:"Details,omitempty" xml:"Details,omitempty"`
	// example:
	//
	// ""
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 90ABA848-AD74-1F6E-84BC-4182A7F1F29E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddWorkNodeWorkforceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddWorkNodeWorkforceResponseBody) GoString() string {
	return s.String()
}

func (s *AddWorkNodeWorkforceResponseBody) SetCode(v int32) *AddWorkNodeWorkforceResponseBody {
	s.Code = &v
	return s
}

func (s *AddWorkNodeWorkforceResponseBody) SetDetails(v string) *AddWorkNodeWorkforceResponseBody {
	s.Details = &v
	return s
}

func (s *AddWorkNodeWorkforceResponseBody) SetErrorCode(v string) *AddWorkNodeWorkforceResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *AddWorkNodeWorkforceResponseBody) SetMessage(v string) *AddWorkNodeWorkforceResponseBody {
	s.Message = &v
	return s
}

func (s *AddWorkNodeWorkforceResponseBody) SetRequestId(v string) *AddWorkNodeWorkforceResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddWorkNodeWorkforceResponseBody) SetSuccess(v bool) *AddWorkNodeWorkforceResponseBody {
	s.Success = &v
	return s
}

type AddWorkNodeWorkforceResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddWorkNodeWorkforceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddWorkNodeWorkforceResponse) String() string {
	return tea.Prettify(s)
}

func (s AddWorkNodeWorkforceResponse) GoString() string {
	return s.String()
}

func (s *AddWorkNodeWorkforceResponse) SetHeaders(v map[string]*string) *AddWorkNodeWorkforceResponse {
	s.Headers = v
	return s
}

func (s *AddWorkNodeWorkforceResponse) SetStatusCode(v int32) *AddWorkNodeWorkforceResponse {
	s.StatusCode = &v
	return s
}

func (s *AddWorkNodeWorkforceResponse) SetBody(v *AddWorkNodeWorkforceResponseBody) *AddWorkNodeWorkforceResponse {
	s.Body = v
	return s
}

type AppendAllDataToTaskRequest struct {
	Body *OpenDatasetProxyAppendDataRequest `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AppendAllDataToTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s AppendAllDataToTaskRequest) GoString() string {
	return s.String()
}

func (s *AppendAllDataToTaskRequest) SetBody(v *OpenDatasetProxyAppendDataRequest) *AppendAllDataToTaskRequest {
	s.Body = v
	return s
}

type AppendAllDataToTaskResponseBody struct {
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// null
	Details *string `json:"Details,omitempty" xml:"Details,omitempty"`
	// example:
	//
	// ""
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 90ABA848-AD74-1F6E-84BC-4182A7F1F29E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AppendAllDataToTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AppendAllDataToTaskResponseBody) GoString() string {
	return s.String()
}

func (s *AppendAllDataToTaskResponseBody) SetCode(v int32) *AppendAllDataToTaskResponseBody {
	s.Code = &v
	return s
}

func (s *AppendAllDataToTaskResponseBody) SetDetails(v string) *AppendAllDataToTaskResponseBody {
	s.Details = &v
	return s
}

func (s *AppendAllDataToTaskResponseBody) SetErrorCode(v string) *AppendAllDataToTaskResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *AppendAllDataToTaskResponseBody) SetMessage(v string) *AppendAllDataToTaskResponseBody {
	s.Message = &v
	return s
}

func (s *AppendAllDataToTaskResponseBody) SetRequestId(v string) *AppendAllDataToTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *AppendAllDataToTaskResponseBody) SetSuccess(v bool) *AppendAllDataToTaskResponseBody {
	s.Success = &v
	return s
}

type AppendAllDataToTaskResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AppendAllDataToTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AppendAllDataToTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s AppendAllDataToTaskResponse) GoString() string {
	return s.String()
}

func (s *AppendAllDataToTaskResponse) SetHeaders(v map[string]*string) *AppendAllDataToTaskResponse {
	s.Headers = v
	return s
}

func (s *AppendAllDataToTaskResponse) SetStatusCode(v int32) *AppendAllDataToTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *AppendAllDataToTaskResponse) SetBody(v *AppendAllDataToTaskResponseBody) *AppendAllDataToTaskResponse {
	s.Body = v
	return s
}

type CreateTaskRequest struct {
	// This parameter is required.
	Body *CreateTaskDetail `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateTaskRequest) SetBody(v *CreateTaskDetail) *CreateTaskRequest {
	s.Body = v
	return s
}

type CreateTaskResponseBody struct {
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// null
	Details *string `json:"Details,omitempty" xml:"Details,omitempty"`
	// example:
	//
	// ""
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 90ABA848-AD74-1F6E-84BC-4182A7F1F29E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 154***2518306500608
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CreateTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTaskResponseBody) SetCode(v int32) *CreateTaskResponseBody {
	s.Code = &v
	return s
}

func (s *CreateTaskResponseBody) SetDetails(v string) *CreateTaskResponseBody {
	s.Details = &v
	return s
}

func (s *CreateTaskResponseBody) SetErrorCode(v string) *CreateTaskResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateTaskResponseBody) SetMessage(v string) *CreateTaskResponseBody {
	s.Message = &v
	return s
}

func (s *CreateTaskResponseBody) SetRequestId(v string) *CreateTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTaskResponseBody) SetSuccess(v bool) *CreateTaskResponseBody {
	s.Success = &v
	return s
}

func (s *CreateTaskResponseBody) SetTaskId(v string) *CreateTaskResponseBody {
	s.TaskId = &v
	return s
}

type CreateTaskResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateTaskResponse) SetHeaders(v map[string]*string) *CreateTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateTaskResponse) SetStatusCode(v int32) *CreateTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateTaskResponse) SetBody(v *CreateTaskResponseBody) *CreateTaskResponse {
	s.Body = v
	return s
}

type CreateTemplateRequest struct {
	// This parameter is required.
	Body *TemplateDTO `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateTemplateRequest) GoString() string {
	return s.String()
}

func (s *CreateTemplateRequest) SetBody(v *TemplateDTO) *CreateTemplateRequest {
	s.Body = v
	return s
}

type CreateTemplateResponseBody struct {
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// -
	Details *string `json:"Details,omitempty" xml:"Details,omitempty"`
	// example:
	//
	// -
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 90ABA848-AD74-1F6E-84BC-4182A7F1F29E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 152***0348342353920
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s CreateTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTemplateResponseBody) SetCode(v int32) *CreateTemplateResponseBody {
	s.Code = &v
	return s
}

func (s *CreateTemplateResponseBody) SetDetails(v string) *CreateTemplateResponseBody {
	s.Details = &v
	return s
}

func (s *CreateTemplateResponseBody) SetErrorCode(v string) *CreateTemplateResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateTemplateResponseBody) SetMessage(v string) *CreateTemplateResponseBody {
	s.Message = &v
	return s
}

func (s *CreateTemplateResponseBody) SetRequestId(v string) *CreateTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTemplateResponseBody) SetSuccess(v bool) *CreateTemplateResponseBody {
	s.Success = &v
	return s
}

func (s *CreateTemplateResponseBody) SetTemplateId(v string) *CreateTemplateResponseBody {
	s.TemplateId = &v
	return s
}

type CreateTemplateResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *CreateTemplateResponse) SetStatusCode(v int32) *CreateTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateTemplateResponse) SetBody(v *CreateTemplateResponseBody) *CreateTemplateResponse {
	s.Body = v
	return s
}

type CreateUserRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 166***980757310
	AccountNo *string `json:"AccountNo,omitempty" xml:"AccountNo,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ALIYUN
	AccountType *string `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ADMIN
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// example:
	//
	// user1
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s CreateUserRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateUserRequest) GoString() string {
	return s.String()
}

func (s *CreateUserRequest) SetAccountNo(v string) *CreateUserRequest {
	s.AccountNo = &v
	return s
}

func (s *CreateUserRequest) SetAccountType(v string) *CreateUserRequest {
	s.AccountType = &v
	return s
}

func (s *CreateUserRequest) SetRole(v string) *CreateUserRequest {
	s.Role = &v
	return s
}

func (s *CreateUserRequest) SetUserName(v string) *CreateUserRequest {
	s.UserName = &v
	return s
}

type CreateUserResponseBody struct {
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// null
	Details *string `json:"Details,omitempty" xml:"Details,omitempty"`
	// example:
	//
	// ""
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 90ABA848-AD74-1F6E-84BC-4182A7F1F29E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 1662339980757311
	UserId *int64 `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s CreateUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateUserResponseBody) GoString() string {
	return s.String()
}

func (s *CreateUserResponseBody) SetCode(v int32) *CreateUserResponseBody {
	s.Code = &v
	return s
}

func (s *CreateUserResponseBody) SetDetails(v string) *CreateUserResponseBody {
	s.Details = &v
	return s
}

func (s *CreateUserResponseBody) SetErrorCode(v string) *CreateUserResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateUserResponseBody) SetMessage(v string) *CreateUserResponseBody {
	s.Message = &v
	return s
}

func (s *CreateUserResponseBody) SetRequestId(v string) *CreateUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateUserResponseBody) SetSuccess(v bool) *CreateUserResponseBody {
	s.Success = &v
	return s
}

func (s *CreateUserResponseBody) SetUserId(v int64) *CreateUserResponseBody {
	s.UserId = &v
	return s
}

type CreateUserResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *CreateUserResponse) SetStatusCode(v int32) *CreateUserResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateUserResponse) SetBody(v *CreateUserResponseBody) *CreateUserResponse {
	s.Body = v
	return s
}

type DeleteTaskResponseBody struct {
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// null
	Details *string `json:"Details,omitempty" xml:"Details,omitempty"`
	// example:
	//
	// ""
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 90ABA848-AD74-1F6E-84BC-4182A7F1F29E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteTaskResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteTaskResponseBody) SetCode(v int32) *DeleteTaskResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteTaskResponseBody) SetDetails(v string) *DeleteTaskResponseBody {
	s.Details = &v
	return s
}

func (s *DeleteTaskResponseBody) SetErrorCode(v string) *DeleteTaskResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteTaskResponseBody) SetMessage(v string) *DeleteTaskResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteTaskResponseBody) SetRequestId(v string) *DeleteTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteTaskResponseBody) SetSuccess(v bool) *DeleteTaskResponseBody {
	s.Success = &v
	return s
}

type DeleteTaskResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteTaskResponse) GoString() string {
	return s.String()
}

func (s *DeleteTaskResponse) SetHeaders(v map[string]*string) *DeleteTaskResponse {
	s.Headers = v
	return s
}

func (s *DeleteTaskResponse) SetStatusCode(v int32) *DeleteTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteTaskResponse) SetBody(v *DeleteTaskResponseBody) *DeleteTaskResponse {
	s.Body = v
	return s
}

type DeleteTemplateResponseBody struct {
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// ""
	Details *string `json:"Details,omitempty" xml:"Details,omitempty"`
	// example:
	//
	// ""
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 90ABA848-AD74-1F6E-84BC-4182A7F1F29E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 152***348342353920
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s DeleteTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteTemplateResponseBody) SetCode(v int32) *DeleteTemplateResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteTemplateResponseBody) SetDetails(v string) *DeleteTemplateResponseBody {
	s.Details = &v
	return s
}

func (s *DeleteTemplateResponseBody) SetErrorCode(v string) *DeleteTemplateResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteTemplateResponseBody) SetMessage(v string) *DeleteTemplateResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteTemplateResponseBody) SetRequestId(v string) *DeleteTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteTemplateResponseBody) SetSuccess(v bool) *DeleteTemplateResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteTemplateResponseBody) SetTemplateId(v string) *DeleteTemplateResponseBody {
	s.TemplateId = &v
	return s
}

type DeleteTemplateResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DeleteTemplateResponse) SetStatusCode(v int32) *DeleteTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteTemplateResponse) SetBody(v *DeleteTemplateResponseBody) *DeleteTemplateResponse {
	s.Body = v
	return s
}

type DeleteUserResponseBody struct {
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// null
	Details *string `json:"Details,omitempty" xml:"Details,omitempty"`
	// example:
	//
	// ""
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 90ABA848-AD74-1F6E-84BC-4182A7F1F29E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteUserResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteUserResponseBody) SetCode(v int32) *DeleteUserResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteUserResponseBody) SetDetails(v string) *DeleteUserResponseBody {
	s.Details = &v
	return s
}

func (s *DeleteUserResponseBody) SetErrorCode(v string) *DeleteUserResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteUserResponseBody) SetMessage(v string) *DeleteUserResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteUserResponseBody) SetRequestId(v string) *DeleteUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteUserResponseBody) SetSuccess(v bool) *DeleteUserResponseBody {
	s.Success = &v
	return s
}

type DeleteUserResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteUserResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteUserResponse) GoString() string {
	return s.String()
}

func (s *DeleteUserResponse) SetHeaders(v map[string]*string) *DeleteUserResponse {
	s.Headers = v
	return s
}

func (s *DeleteUserResponse) SetStatusCode(v int32) *DeleteUserResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteUserResponse) SetBody(v *DeleteUserResponseBody) *DeleteUserResponse {
	s.Body = v
	return s
}

type ExportAnnotationsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// oss://***-hz-oss.oss-cn-hangzhou.aliyuncs.com/output/
	OssPath *string `json:"OssPath,omitempty" xml:"OssPath,omitempty"`
	// example:
	//
	// true
	RegisterDataset *string `json:"RegisterDataset,omitempty" xml:"RegisterDataset,omitempty"`
	// example:
	//
	// PAI
	Target *string `json:"Target,omitempty" xml:"Target,omitempty"`
}

func (s ExportAnnotationsRequest) String() string {
	return tea.Prettify(s)
}

func (s ExportAnnotationsRequest) GoString() string {
	return s.String()
}

func (s *ExportAnnotationsRequest) SetOssPath(v string) *ExportAnnotationsRequest {
	s.OssPath = &v
	return s
}

func (s *ExportAnnotationsRequest) SetRegisterDataset(v string) *ExportAnnotationsRequest {
	s.RegisterDataset = &v
	return s
}

func (s *ExportAnnotationsRequest) SetTarget(v string) *ExportAnnotationsRequest {
	s.Target = &v
	return s
}

type ExportAnnotationsResponseBody struct {
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// null
	Details *string `json:"Details,omitempty" xml:"Details,omitempty"`
	// example:
	//
	// ""
	ErrorCode *string      `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	FlowJob   *FlowJobInfo `json:"FlowJob,omitempty" xml:"FlowJob,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 90ABA848-AD74-1F6E-84BC-4182A7F1F29E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ExportAnnotationsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ExportAnnotationsResponseBody) GoString() string {
	return s.String()
}

func (s *ExportAnnotationsResponseBody) SetCode(v int32) *ExportAnnotationsResponseBody {
	s.Code = &v
	return s
}

func (s *ExportAnnotationsResponseBody) SetDetails(v string) *ExportAnnotationsResponseBody {
	s.Details = &v
	return s
}

func (s *ExportAnnotationsResponseBody) SetErrorCode(v string) *ExportAnnotationsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ExportAnnotationsResponseBody) SetFlowJob(v *FlowJobInfo) *ExportAnnotationsResponseBody {
	s.FlowJob = v
	return s
}

func (s *ExportAnnotationsResponseBody) SetMessage(v string) *ExportAnnotationsResponseBody {
	s.Message = &v
	return s
}

func (s *ExportAnnotationsResponseBody) SetRequestId(v string) *ExportAnnotationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ExportAnnotationsResponseBody) SetSuccess(v bool) *ExportAnnotationsResponseBody {
	s.Success = &v
	return s
}

type ExportAnnotationsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ExportAnnotationsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExportAnnotationsResponse) String() string {
	return tea.Prettify(s)
}

func (s ExportAnnotationsResponse) GoString() string {
	return s.String()
}

func (s *ExportAnnotationsResponse) SetHeaders(v map[string]*string) *ExportAnnotationsResponse {
	s.Headers = v
	return s
}

func (s *ExportAnnotationsResponse) SetStatusCode(v int32) *ExportAnnotationsResponse {
	s.StatusCode = &v
	return s
}

func (s *ExportAnnotationsResponse) SetBody(v *ExportAnnotationsResponseBody) *ExportAnnotationsResponse {
	s.Body = v
	return s
}

type GetJobRequest struct {
	// example:
	//
	// DOWNLOWD_MARKRESULT_FLOW
	JobType *string `json:"JobType,omitempty" xml:"JobType,omitempty"`
}

func (s GetJobRequest) String() string {
	return tea.Prettify(s)
}

func (s GetJobRequest) GoString() string {
	return s.String()
}

func (s *GetJobRequest) SetJobType(v string) *GetJobRequest {
	s.JobType = &v
	return s
}

type GetJobResponseBody struct {
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// null
	Details *string `json:"Details,omitempty" xml:"Details,omitempty"`
	// example:
	//
	// ""
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Job       *Job    `json:"Job,omitempty" xml:"Job,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 90ABA848-AD74-1F6E-84BC-4182A7F1F29E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetJobResponseBody) GoString() string {
	return s.String()
}

func (s *GetJobResponseBody) SetCode(v int32) *GetJobResponseBody {
	s.Code = &v
	return s
}

func (s *GetJobResponseBody) SetDetails(v string) *GetJobResponseBody {
	s.Details = &v
	return s
}

func (s *GetJobResponseBody) SetErrorCode(v string) *GetJobResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetJobResponseBody) SetJob(v *Job) *GetJobResponseBody {
	s.Job = v
	return s
}

func (s *GetJobResponseBody) SetMessage(v string) *GetJobResponseBody {
	s.Message = &v
	return s
}

func (s *GetJobResponseBody) SetRequestId(v string) *GetJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetJobResponseBody) SetSuccess(v bool) *GetJobResponseBody {
	s.Success = &v
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

type GetSubtaskResponseBody struct {
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// null
	Details *string `json:"Details,omitempty" xml:"Details,omitempty"`
	// example:
	//
	// ""
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 90ABA848-AD74-1F6E-84BC-4182A7F1F29E
	RequestId *string        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Subtask   *SimpleSubtask `json:"Subtask,omitempty" xml:"Subtask,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetSubtaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSubtaskResponseBody) GoString() string {
	return s.String()
}

func (s *GetSubtaskResponseBody) SetCode(v int32) *GetSubtaskResponseBody {
	s.Code = &v
	return s
}

func (s *GetSubtaskResponseBody) SetDetails(v string) *GetSubtaskResponseBody {
	s.Details = &v
	return s
}

func (s *GetSubtaskResponseBody) SetErrorCode(v string) *GetSubtaskResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetSubtaskResponseBody) SetMessage(v string) *GetSubtaskResponseBody {
	s.Message = &v
	return s
}

func (s *GetSubtaskResponseBody) SetRequestId(v string) *GetSubtaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSubtaskResponseBody) SetSubtask(v *SimpleSubtask) *GetSubtaskResponseBody {
	s.Subtask = v
	return s
}

func (s *GetSubtaskResponseBody) SetSuccess(v bool) *GetSubtaskResponseBody {
	s.Success = &v
	return s
}

type GetSubtaskResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSubtaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSubtaskResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSubtaskResponse) GoString() string {
	return s.String()
}

func (s *GetSubtaskResponse) SetHeaders(v map[string]*string) *GetSubtaskResponse {
	s.Headers = v
	return s
}

func (s *GetSubtaskResponse) SetStatusCode(v int32) *GetSubtaskResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSubtaskResponse) SetBody(v *GetSubtaskResponseBody) *GetSubtaskResponse {
	s.Body = v
	return s
}

type GetSubtaskItemResponseBody struct {
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// null
	Details *string `json:"Details,omitempty" xml:"Details,omitempty"`
	// example:
	//
	// ""
	ErrorCode *string            `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Item      *SubtaskItemDetail `json:"Item,omitempty" xml:"Item,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 90ABA848-AD74-1F6E-84BC-4182A7F1F29E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetSubtaskItemResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSubtaskItemResponseBody) GoString() string {
	return s.String()
}

func (s *GetSubtaskItemResponseBody) SetCode(v int32) *GetSubtaskItemResponseBody {
	s.Code = &v
	return s
}

func (s *GetSubtaskItemResponseBody) SetDetails(v string) *GetSubtaskItemResponseBody {
	s.Details = &v
	return s
}

func (s *GetSubtaskItemResponseBody) SetErrorCode(v string) *GetSubtaskItemResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetSubtaskItemResponseBody) SetItem(v *SubtaskItemDetail) *GetSubtaskItemResponseBody {
	s.Item = v
	return s
}

func (s *GetSubtaskItemResponseBody) SetMessage(v string) *GetSubtaskItemResponseBody {
	s.Message = &v
	return s
}

func (s *GetSubtaskItemResponseBody) SetRequestId(v string) *GetSubtaskItemResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSubtaskItemResponseBody) SetSuccess(v bool) *GetSubtaskItemResponseBody {
	s.Success = &v
	return s
}

type GetSubtaskItemResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSubtaskItemResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSubtaskItemResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSubtaskItemResponse) GoString() string {
	return s.String()
}

func (s *GetSubtaskItemResponse) SetHeaders(v map[string]*string) *GetSubtaskItemResponse {
	s.Headers = v
	return s
}

func (s *GetSubtaskItemResponse) SetStatusCode(v int32) *GetSubtaskItemResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSubtaskItemResponse) SetBody(v *GetSubtaskItemResponseBody) *GetSubtaskItemResponse {
	s.Body = v
	return s
}

type GetTaskResponseBody struct {
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// null
	Details *string `json:"Details,omitempty" xml:"Details,omitempty"`
	// example:
	//
	// ""
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 0F01E603-8A9F-18ED-AD43-D52B5030AFA2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool       `json:"Success,omitempty" xml:"Success,omitempty"`
	Task    *TaskDetail `json:"Task,omitempty" xml:"Task,omitempty"`
}

func (s GetTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTaskResponseBody) GoString() string {
	return s.String()
}

func (s *GetTaskResponseBody) SetCode(v int32) *GetTaskResponseBody {
	s.Code = &v
	return s
}

func (s *GetTaskResponseBody) SetDetails(v string) *GetTaskResponseBody {
	s.Details = &v
	return s
}

func (s *GetTaskResponseBody) SetErrorCode(v string) *GetTaskResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetTaskResponseBody) SetMessage(v string) *GetTaskResponseBody {
	s.Message = &v
	return s
}

func (s *GetTaskResponseBody) SetRequestId(v string) *GetTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTaskResponseBody) SetSuccess(v bool) *GetTaskResponseBody {
	s.Success = &v
	return s
}

func (s *GetTaskResponseBody) SetTask(v *TaskDetail) *GetTaskResponseBody {
	s.Task = v
	return s
}

type GetTaskResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTaskResponse) GoString() string {
	return s.String()
}

func (s *GetTaskResponse) SetHeaders(v map[string]*string) *GetTaskResponse {
	s.Headers = v
	return s
}

func (s *GetTaskResponse) SetStatusCode(v int32) *GetTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTaskResponse) SetBody(v *GetTaskResponseBody) *GetTaskResponse {
	s.Body = v
	return s
}

type GetTaskStatisticsRequest struct {
	// example:
	//
	// Item
	StatType *string `json:"StatType,omitempty" xml:"StatType,omitempty"`
}

func (s GetTaskStatisticsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTaskStatisticsRequest) GoString() string {
	return s.String()
}

func (s *GetTaskStatisticsRequest) SetStatType(v string) *GetTaskStatisticsRequest {
	s.StatType = &v
	return s
}

type GetTaskStatisticsResponseBody struct {
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// null
	Details *string `json:"Details,omitempty" xml:"Details,omitempty"`
	// example:
	//
	// ""
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 90ABA848-AD74-1F6E-84BC-4182A7F1F29E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success        *bool          `json:"Success,omitempty" xml:"Success,omitempty"`
	TaskStatistics *TaskStatistic `json:"TaskStatistics,omitempty" xml:"TaskStatistics,omitempty"`
}

func (s GetTaskStatisticsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTaskStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *GetTaskStatisticsResponseBody) SetCode(v int32) *GetTaskStatisticsResponseBody {
	s.Code = &v
	return s
}

func (s *GetTaskStatisticsResponseBody) SetDetails(v string) *GetTaskStatisticsResponseBody {
	s.Details = &v
	return s
}

func (s *GetTaskStatisticsResponseBody) SetErrorCode(v string) *GetTaskStatisticsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetTaskStatisticsResponseBody) SetMessage(v string) *GetTaskStatisticsResponseBody {
	s.Message = &v
	return s
}

func (s *GetTaskStatisticsResponseBody) SetRequestId(v string) *GetTaskStatisticsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTaskStatisticsResponseBody) SetSuccess(v bool) *GetTaskStatisticsResponseBody {
	s.Success = &v
	return s
}

func (s *GetTaskStatisticsResponseBody) SetTaskStatistics(v *TaskStatistic) *GetTaskStatisticsResponseBody {
	s.TaskStatistics = v
	return s
}

type GetTaskStatisticsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTaskStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTaskStatisticsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTaskStatisticsResponse) GoString() string {
	return s.String()
}

func (s *GetTaskStatisticsResponse) SetHeaders(v map[string]*string) *GetTaskStatisticsResponse {
	s.Headers = v
	return s
}

func (s *GetTaskStatisticsResponse) SetStatusCode(v int32) *GetTaskStatisticsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTaskStatisticsResponse) SetBody(v *GetTaskStatisticsResponseBody) *GetTaskStatisticsResponse {
	s.Body = v
	return s
}

type GetTaskStatusResponseBody struct {
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// ""
	Details *string `json:"Details,omitempty" xml:"Details,omitempty"`
	// example:
	//
	// ""
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 90ABA848-AD74-1F6E-84BC-4182A7F1F29E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// SUCCESS
	TaskStatus *string `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
}

func (s GetTaskStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTaskStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetTaskStatusResponseBody) SetCode(v int32) *GetTaskStatusResponseBody {
	s.Code = &v
	return s
}

func (s *GetTaskStatusResponseBody) SetDetails(v string) *GetTaskStatusResponseBody {
	s.Details = &v
	return s
}

func (s *GetTaskStatusResponseBody) SetErrorCode(v string) *GetTaskStatusResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetTaskStatusResponseBody) SetMessage(v string) *GetTaskStatusResponseBody {
	s.Message = &v
	return s
}

func (s *GetTaskStatusResponseBody) SetRequestId(v string) *GetTaskStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTaskStatusResponseBody) SetSuccess(v bool) *GetTaskStatusResponseBody {
	s.Success = &v
	return s
}

func (s *GetTaskStatusResponseBody) SetTaskStatus(v string) *GetTaskStatusResponseBody {
	s.TaskStatus = &v
	return s
}

type GetTaskStatusResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTaskStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTaskStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTaskStatusResponse) GoString() string {
	return s.String()
}

func (s *GetTaskStatusResponse) SetHeaders(v map[string]*string) *GetTaskStatusResponse {
	s.Headers = v
	return s
}

func (s *GetTaskStatusResponse) SetStatusCode(v int32) *GetTaskStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTaskStatusResponse) SetBody(v *GetTaskStatusResponseBody) *GetTaskStatusResponse {
	s.Body = v
	return s
}

type GetTaskTemplateResponseBody struct {
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// ""
	Details *string `json:"Details,omitempty" xml:"Details,omitempty"`
	// example:
	//
	// ""
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 90ABA848-AD74-1F6E-84BC-4182A7F1F29E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success  *bool           `json:"Success,omitempty" xml:"Success,omitempty"`
	Template *TemplateDetail `json:"Template,omitempty" xml:"Template,omitempty"`
}

func (s GetTaskTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTaskTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *GetTaskTemplateResponseBody) SetCode(v int32) *GetTaskTemplateResponseBody {
	s.Code = &v
	return s
}

func (s *GetTaskTemplateResponseBody) SetDetails(v string) *GetTaskTemplateResponseBody {
	s.Details = &v
	return s
}

func (s *GetTaskTemplateResponseBody) SetErrorCode(v string) *GetTaskTemplateResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetTaskTemplateResponseBody) SetMessage(v string) *GetTaskTemplateResponseBody {
	s.Message = &v
	return s
}

func (s *GetTaskTemplateResponseBody) SetRequestId(v string) *GetTaskTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTaskTemplateResponseBody) SetSuccess(v bool) *GetTaskTemplateResponseBody {
	s.Success = &v
	return s
}

func (s *GetTaskTemplateResponseBody) SetTemplate(v *TemplateDetail) *GetTaskTemplateResponseBody {
	s.Template = v
	return s
}

type GetTaskTemplateResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTaskTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTaskTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTaskTemplateResponse) GoString() string {
	return s.String()
}

func (s *GetTaskTemplateResponse) SetHeaders(v map[string]*string) *GetTaskTemplateResponse {
	s.Headers = v
	return s
}

func (s *GetTaskTemplateResponse) SetStatusCode(v int32) *GetTaskTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTaskTemplateResponse) SetBody(v *GetTaskTemplateResponseBody) *GetTaskTemplateResponse {
	s.Body = v
	return s
}

type GetTaskTemplateQuestionsResponseBody struct {
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// null
	Details *string `json:"Details,omitempty" xml:"Details,omitempty"`
	// example:
	//
	// ""
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// success
	Message   *string           `json:"Message,omitempty" xml:"Message,omitempty"`
	Questions []*QuestionPlugin `json:"Questions,omitempty" xml:"Questions,omitempty" type:"Repeated"`
	// example:
	//
	// 90ABA848-AD74-1F6E-84BC-4182A7F1F29E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetTaskTemplateQuestionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTaskTemplateQuestionsResponseBody) GoString() string {
	return s.String()
}

func (s *GetTaskTemplateQuestionsResponseBody) SetCode(v int32) *GetTaskTemplateQuestionsResponseBody {
	s.Code = &v
	return s
}

func (s *GetTaskTemplateQuestionsResponseBody) SetDetails(v string) *GetTaskTemplateQuestionsResponseBody {
	s.Details = &v
	return s
}

func (s *GetTaskTemplateQuestionsResponseBody) SetErrorCode(v string) *GetTaskTemplateQuestionsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetTaskTemplateQuestionsResponseBody) SetMessage(v string) *GetTaskTemplateQuestionsResponseBody {
	s.Message = &v
	return s
}

func (s *GetTaskTemplateQuestionsResponseBody) SetQuestions(v []*QuestionPlugin) *GetTaskTemplateQuestionsResponseBody {
	s.Questions = v
	return s
}

func (s *GetTaskTemplateQuestionsResponseBody) SetRequestId(v string) *GetTaskTemplateQuestionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTaskTemplateQuestionsResponseBody) SetSuccess(v bool) *GetTaskTemplateQuestionsResponseBody {
	s.Success = &v
	return s
}

type GetTaskTemplateQuestionsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTaskTemplateQuestionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTaskTemplateQuestionsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTaskTemplateQuestionsResponse) GoString() string {
	return s.String()
}

func (s *GetTaskTemplateQuestionsResponse) SetHeaders(v map[string]*string) *GetTaskTemplateQuestionsResponse {
	s.Headers = v
	return s
}

func (s *GetTaskTemplateQuestionsResponse) SetStatusCode(v int32) *GetTaskTemplateQuestionsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTaskTemplateQuestionsResponse) SetBody(v *GetTaskTemplateQuestionsResponseBody) *GetTaskTemplateQuestionsResponse {
	s.Body = v
	return s
}

type GetTaskTemplateViewsResponseBody struct {
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// null
	Details *string `json:"Details,omitempty" xml:"Details,omitempty"`
	// example:
	//
	// ""
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 90ABA848-AD74-1F6E-84BC-4182A7F1F29E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool                                  `json:"Success,omitempty" xml:"Success,omitempty"`
	Views   *GetTaskTemplateViewsResponseBodyViews `json:"Views,omitempty" xml:"Views,omitempty" type:"Struct"`
}

func (s GetTaskTemplateViewsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTaskTemplateViewsResponseBody) GoString() string {
	return s.String()
}

func (s *GetTaskTemplateViewsResponseBody) SetCode(v int32) *GetTaskTemplateViewsResponseBody {
	s.Code = &v
	return s
}

func (s *GetTaskTemplateViewsResponseBody) SetDetails(v string) *GetTaskTemplateViewsResponseBody {
	s.Details = &v
	return s
}

func (s *GetTaskTemplateViewsResponseBody) SetErrorCode(v string) *GetTaskTemplateViewsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetTaskTemplateViewsResponseBody) SetMessage(v string) *GetTaskTemplateViewsResponseBody {
	s.Message = &v
	return s
}

func (s *GetTaskTemplateViewsResponseBody) SetRequestId(v string) *GetTaskTemplateViewsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTaskTemplateViewsResponseBody) SetSuccess(v bool) *GetTaskTemplateViewsResponseBody {
	s.Success = &v
	return s
}

func (s *GetTaskTemplateViewsResponseBody) SetViews(v *GetTaskTemplateViewsResponseBodyViews) *GetTaskTemplateViewsResponseBody {
	s.Views = v
	return s
}

type GetTaskTemplateViewsResponseBodyViews struct {
	ViewPlugins []*ViewPlugin `json:"ViewPlugins,omitempty" xml:"ViewPlugins,omitempty" type:"Repeated"`
}

func (s GetTaskTemplateViewsResponseBodyViews) String() string {
	return tea.Prettify(s)
}

func (s GetTaskTemplateViewsResponseBodyViews) GoString() string {
	return s.String()
}

func (s *GetTaskTemplateViewsResponseBodyViews) SetViewPlugins(v []*ViewPlugin) *GetTaskTemplateViewsResponseBodyViews {
	s.ViewPlugins = v
	return s
}

type GetTaskTemplateViewsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTaskTemplateViewsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTaskTemplateViewsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTaskTemplateViewsResponse) GoString() string {
	return s.String()
}

func (s *GetTaskTemplateViewsResponse) SetHeaders(v map[string]*string) *GetTaskTemplateViewsResponse {
	s.Headers = v
	return s
}

func (s *GetTaskTemplateViewsResponse) SetStatusCode(v int32) *GetTaskTemplateViewsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTaskTemplateViewsResponse) SetBody(v *GetTaskTemplateViewsResponseBody) *GetTaskTemplateViewsResponse {
	s.Body = v
	return s
}

type GetTaskWorkforceResponseBody struct {
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// null
	Details *string `json:"Details,omitempty" xml:"Details,omitempty"`
	// example:
	//
	// ""
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 90ABA848-AD74-1F6E-84BC-4182A7F1F29E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success   *bool        `json:"Success,omitempty" xml:"Success,omitempty"`
	Workforce []*Workforce `json:"Workforce,omitempty" xml:"Workforce,omitempty" type:"Repeated"`
}

func (s GetTaskWorkforceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTaskWorkforceResponseBody) GoString() string {
	return s.String()
}

func (s *GetTaskWorkforceResponseBody) SetCode(v int32) *GetTaskWorkforceResponseBody {
	s.Code = &v
	return s
}

func (s *GetTaskWorkforceResponseBody) SetDetails(v string) *GetTaskWorkforceResponseBody {
	s.Details = &v
	return s
}

func (s *GetTaskWorkforceResponseBody) SetErrorCode(v string) *GetTaskWorkforceResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetTaskWorkforceResponseBody) SetMessage(v string) *GetTaskWorkforceResponseBody {
	s.Message = &v
	return s
}

func (s *GetTaskWorkforceResponseBody) SetRequestId(v string) *GetTaskWorkforceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTaskWorkforceResponseBody) SetSuccess(v bool) *GetTaskWorkforceResponseBody {
	s.Success = &v
	return s
}

func (s *GetTaskWorkforceResponseBody) SetWorkforce(v []*Workforce) *GetTaskWorkforceResponseBody {
	s.Workforce = v
	return s
}

type GetTaskWorkforceResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTaskWorkforceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTaskWorkforceResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTaskWorkforceResponse) GoString() string {
	return s.String()
}

func (s *GetTaskWorkforceResponse) SetHeaders(v map[string]*string) *GetTaskWorkforceResponse {
	s.Headers = v
	return s
}

func (s *GetTaskWorkforceResponse) SetStatusCode(v int32) *GetTaskWorkforceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTaskWorkforceResponse) SetBody(v *GetTaskWorkforceResponseBody) *GetTaskWorkforceResponse {
	s.Body = v
	return s
}

type GetTaskWorkforceStatisticRequest struct {
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
	// Item
	StatType *string `json:"StatType,omitempty" xml:"StatType,omitempty"`
}

func (s GetTaskWorkforceStatisticRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTaskWorkforceStatisticRequest) GoString() string {
	return s.String()
}

func (s *GetTaskWorkforceStatisticRequest) SetPageNumber(v int32) *GetTaskWorkforceStatisticRequest {
	s.PageNumber = &v
	return s
}

func (s *GetTaskWorkforceStatisticRequest) SetPageSize(v int32) *GetTaskWorkforceStatisticRequest {
	s.PageSize = &v
	return s
}

func (s *GetTaskWorkforceStatisticRequest) SetStatType(v string) *GetTaskWorkforceStatisticRequest {
	s.StatType = &v
	return s
}

type GetTaskWorkforceStatisticResponseBody struct {
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// null
	Details *string `json:"Details,omitempty" xml:"Details,omitempty"`
	// example:
	//
	// ""
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
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
	// 90ABA848-AD74-1F6E-84BC-4182A7F1F29E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 22
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// example:
	//
	// 2
	TotalPage      *int32           `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
	UsersStatistic []*UserStatistic `json:"UsersStatistic,omitempty" xml:"UsersStatistic,omitempty" type:"Repeated"`
}

func (s GetTaskWorkforceStatisticResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTaskWorkforceStatisticResponseBody) GoString() string {
	return s.String()
}

func (s *GetTaskWorkforceStatisticResponseBody) SetCode(v int32) *GetTaskWorkforceStatisticResponseBody {
	s.Code = &v
	return s
}

func (s *GetTaskWorkforceStatisticResponseBody) SetDetails(v string) *GetTaskWorkforceStatisticResponseBody {
	s.Details = &v
	return s
}

func (s *GetTaskWorkforceStatisticResponseBody) SetErrorCode(v string) *GetTaskWorkforceStatisticResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetTaskWorkforceStatisticResponseBody) SetMessage(v string) *GetTaskWorkforceStatisticResponseBody {
	s.Message = &v
	return s
}

func (s *GetTaskWorkforceStatisticResponseBody) SetPageNumber(v int32) *GetTaskWorkforceStatisticResponseBody {
	s.PageNumber = &v
	return s
}

func (s *GetTaskWorkforceStatisticResponseBody) SetPageSize(v int32) *GetTaskWorkforceStatisticResponseBody {
	s.PageSize = &v
	return s
}

func (s *GetTaskWorkforceStatisticResponseBody) SetRequestId(v string) *GetTaskWorkforceStatisticResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTaskWorkforceStatisticResponseBody) SetSuccess(v bool) *GetTaskWorkforceStatisticResponseBody {
	s.Success = &v
	return s
}

func (s *GetTaskWorkforceStatisticResponseBody) SetTotalCount(v int32) *GetTaskWorkforceStatisticResponseBody {
	s.TotalCount = &v
	return s
}

func (s *GetTaskWorkforceStatisticResponseBody) SetTotalPage(v int32) *GetTaskWorkforceStatisticResponseBody {
	s.TotalPage = &v
	return s
}

func (s *GetTaskWorkforceStatisticResponseBody) SetUsersStatistic(v []*UserStatistic) *GetTaskWorkforceStatisticResponseBody {
	s.UsersStatistic = v
	return s
}

type GetTaskWorkforceStatisticResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTaskWorkforceStatisticResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTaskWorkforceStatisticResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTaskWorkforceStatisticResponse) GoString() string {
	return s.String()
}

func (s *GetTaskWorkforceStatisticResponse) SetHeaders(v map[string]*string) *GetTaskWorkforceStatisticResponse {
	s.Headers = v
	return s
}

func (s *GetTaskWorkforceStatisticResponse) SetStatusCode(v int32) *GetTaskWorkforceStatisticResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTaskWorkforceStatisticResponse) SetBody(v *GetTaskWorkforceStatisticResponseBody) *GetTaskWorkforceStatisticResponse {
	s.Body = v
	return s
}

type GetTemplateResponseBody struct {
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// null
	Details *string `json:"Details,omitempty" xml:"Details,omitempty"`
	// example:
	//
	// ""
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 90ABA848-AD74-1F6E-84BC-4182A7F1F29E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success  *bool           `json:"Success,omitempty" xml:"Success,omitempty"`
	Template *TemplateDetail `json:"Template,omitempty" xml:"Template,omitempty"`
}

func (s GetTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *GetTemplateResponseBody) SetCode(v int32) *GetTemplateResponseBody {
	s.Code = &v
	return s
}

func (s *GetTemplateResponseBody) SetDetails(v string) *GetTemplateResponseBody {
	s.Details = &v
	return s
}

func (s *GetTemplateResponseBody) SetErrorCode(v string) *GetTemplateResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetTemplateResponseBody) SetMessage(v string) *GetTemplateResponseBody {
	s.Message = &v
	return s
}

func (s *GetTemplateResponseBody) SetRequestId(v string) *GetTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTemplateResponseBody) SetSuccess(v bool) *GetTemplateResponseBody {
	s.Success = &v
	return s
}

func (s *GetTemplateResponseBody) SetTemplate(v *TemplateDetail) *GetTemplateResponseBody {
	s.Template = v
	return s
}

type GetTemplateResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *GetTemplateResponse) SetStatusCode(v int32) *GetTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTemplateResponse) SetBody(v *GetTemplateResponseBody) *GetTemplateResponse {
	s.Body = v
	return s
}

type GetTemplateQuestionsResponseBody struct {
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// null
	Details *string `json:"Details,omitempty" xml:"Details,omitempty"`
	// example:
	//
	// ""
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// success
	Message         *string           `json:"Message,omitempty" xml:"Message,omitempty"`
	QuestionConfigs []*QuestionPlugin `json:"QuestionConfigs,omitempty" xml:"QuestionConfigs,omitempty" type:"Repeated"`
	// example:
	//
	// 90ABA848-AD74-1F6E-84BC-4182A7F1F29E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetTemplateQuestionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTemplateQuestionsResponseBody) GoString() string {
	return s.String()
}

func (s *GetTemplateQuestionsResponseBody) SetCode(v int32) *GetTemplateQuestionsResponseBody {
	s.Code = &v
	return s
}

func (s *GetTemplateQuestionsResponseBody) SetDetails(v string) *GetTemplateQuestionsResponseBody {
	s.Details = &v
	return s
}

func (s *GetTemplateQuestionsResponseBody) SetErrorCode(v string) *GetTemplateQuestionsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetTemplateQuestionsResponseBody) SetMessage(v string) *GetTemplateQuestionsResponseBody {
	s.Message = &v
	return s
}

func (s *GetTemplateQuestionsResponseBody) SetQuestionConfigs(v []*QuestionPlugin) *GetTemplateQuestionsResponseBody {
	s.QuestionConfigs = v
	return s
}

func (s *GetTemplateQuestionsResponseBody) SetRequestId(v string) *GetTemplateQuestionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTemplateQuestionsResponseBody) SetSuccess(v bool) *GetTemplateQuestionsResponseBody {
	s.Success = &v
	return s
}

type GetTemplateQuestionsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTemplateQuestionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTemplateQuestionsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTemplateQuestionsResponse) GoString() string {
	return s.String()
}

func (s *GetTemplateQuestionsResponse) SetHeaders(v map[string]*string) *GetTemplateQuestionsResponse {
	s.Headers = v
	return s
}

func (s *GetTemplateQuestionsResponse) SetStatusCode(v int32) *GetTemplateQuestionsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTemplateQuestionsResponse) SetBody(v *GetTemplateQuestionsResponseBody) *GetTemplateQuestionsResponse {
	s.Body = v
	return s
}

type GetTemplateViewResponseBody struct {
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// ""
	Details *string `json:"Details,omitempty" xml:"Details,omitempty"`
	// example:
	//
	// ""
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 90ABA848-AD74-1F6E-84BC-4182A7F1F29E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success     *bool                                   `json:"Success,omitempty" xml:"Success,omitempty"`
	ViewConfigs *GetTemplateViewResponseBodyViewConfigs `json:"ViewConfigs,omitempty" xml:"ViewConfigs,omitempty" type:"Struct"`
}

func (s GetTemplateViewResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTemplateViewResponseBody) GoString() string {
	return s.String()
}

func (s *GetTemplateViewResponseBody) SetCode(v int32) *GetTemplateViewResponseBody {
	s.Code = &v
	return s
}

func (s *GetTemplateViewResponseBody) SetDetails(v string) *GetTemplateViewResponseBody {
	s.Details = &v
	return s
}

func (s *GetTemplateViewResponseBody) SetErrorCode(v string) *GetTemplateViewResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetTemplateViewResponseBody) SetMessage(v string) *GetTemplateViewResponseBody {
	s.Message = &v
	return s
}

func (s *GetTemplateViewResponseBody) SetRequestId(v string) *GetTemplateViewResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTemplateViewResponseBody) SetSuccess(v bool) *GetTemplateViewResponseBody {
	s.Success = &v
	return s
}

func (s *GetTemplateViewResponseBody) SetViewConfigs(v *GetTemplateViewResponseBodyViewConfigs) *GetTemplateViewResponseBody {
	s.ViewConfigs = v
	return s
}

type GetTemplateViewResponseBodyViewConfigs struct {
	ViewPlugins []*ViewPlugin `json:"ViewPlugins,omitempty" xml:"ViewPlugins,omitempty" type:"Repeated"`
}

func (s GetTemplateViewResponseBodyViewConfigs) String() string {
	return tea.Prettify(s)
}

func (s GetTemplateViewResponseBodyViewConfigs) GoString() string {
	return s.String()
}

func (s *GetTemplateViewResponseBodyViewConfigs) SetViewPlugins(v []*ViewPlugin) *GetTemplateViewResponseBodyViewConfigs {
	s.ViewPlugins = v
	return s
}

type GetTemplateViewResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTemplateViewResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTemplateViewResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTemplateViewResponse) GoString() string {
	return s.String()
}

func (s *GetTemplateViewResponse) SetHeaders(v map[string]*string) *GetTemplateViewResponse {
	s.Headers = v
	return s
}

func (s *GetTemplateViewResponse) SetStatusCode(v int32) *GetTemplateViewResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTemplateViewResponse) SetBody(v *GetTemplateViewResponseBody) *GetTemplateViewResponse {
	s.Body = v
	return s
}

type GetTenantResponseBody struct {
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// -
	Details *string `json:"Details,omitempty" xml:"Details,omitempty"`
	// example:
	//
	// -
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 90ABA848-AD74-1F6E-84BC-4182A7F1F29E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool         `json:"Success,omitempty" xml:"Success,omitempty"`
	Tenant  *SingleTenant `json:"Tenant,omitempty" xml:"Tenant,omitempty"`
}

func (s GetTenantResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTenantResponseBody) GoString() string {
	return s.String()
}

func (s *GetTenantResponseBody) SetCode(v int32) *GetTenantResponseBody {
	s.Code = &v
	return s
}

func (s *GetTenantResponseBody) SetDetails(v string) *GetTenantResponseBody {
	s.Details = &v
	return s
}

func (s *GetTenantResponseBody) SetErrorCode(v string) *GetTenantResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetTenantResponseBody) SetMessage(v string) *GetTenantResponseBody {
	s.Message = &v
	return s
}

func (s *GetTenantResponseBody) SetRequestId(v string) *GetTenantResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTenantResponseBody) SetSuccess(v bool) *GetTenantResponseBody {
	s.Success = &v
	return s
}

func (s *GetTenantResponseBody) SetTenant(v *SingleTenant) *GetTenantResponseBody {
	s.Tenant = v
	return s
}

type GetTenantResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTenantResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTenantResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTenantResponse) GoString() string {
	return s.String()
}

func (s *GetTenantResponse) SetHeaders(v map[string]*string) *GetTenantResponse {
	s.Headers = v
	return s
}

func (s *GetTenantResponse) SetStatusCode(v int32) *GetTenantResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTenantResponse) SetBody(v *GetTenantResponseBody) *GetTenantResponse {
	s.Body = v
	return s
}

type GetUserResponseBody struct {
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// null
	Details *string `json:"Details,omitempty" xml:"Details,omitempty"`
	// example:
	//
	// ""
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 90ABA848-AD74-1F6E-84BC-4182A7F1F29E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool       `json:"Success,omitempty" xml:"Success,omitempty"`
	User    *SimpleUser `json:"User,omitempty" xml:"User,omitempty"`
}

func (s GetUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetUserResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserResponseBody) SetCode(v int32) *GetUserResponseBody {
	s.Code = &v
	return s
}

func (s *GetUserResponseBody) SetDetails(v string) *GetUserResponseBody {
	s.Details = &v
	return s
}

func (s *GetUserResponseBody) SetErrorCode(v string) *GetUserResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetUserResponseBody) SetMessage(v string) *GetUserResponseBody {
	s.Message = &v
	return s
}

func (s *GetUserResponseBody) SetRequestId(v string) *GetUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetUserResponseBody) SetSuccess(v bool) *GetUserResponseBody {
	s.Success = &v
	return s
}

func (s *GetUserResponseBody) SetUser(v *SimpleUser) *GetUserResponseBody {
	s.User = v
	return s
}

type GetUserResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetUserResponse) String() string {
	return tea.Prettify(s)
}

func (s GetUserResponse) GoString() string {
	return s.String()
}

func (s *GetUserResponse) SetHeaders(v map[string]*string) *GetUserResponse {
	s.Headers = v
	return s
}

func (s *GetUserResponse) SetStatusCode(v int32) *GetUserResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUserResponse) SetBody(v *GetUserResponseBody) *GetUserResponse {
	s.Body = v
	return s
}

type ListJobsRequest struct {
	// example:
	//
	// DOWNLOWD_MARKRESULT_FLOW
	JobType *string `json:"JobType,omitempty" xml:"JobType,omitempty"`
	// example:
	//
	// 20
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 1
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListJobsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListJobsRequest) GoString() string {
	return s.String()
}

func (s *ListJobsRequest) SetJobType(v string) *ListJobsRequest {
	s.JobType = &v
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

type ListJobsResponseBody struct {
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// null
	Details *string `json:"Details,omitempty" xml:"Details,omitempty"`
	// example:
	//
	// ""
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Jobs      []*Job  `json:"Jobs,omitempty" xml:"Jobs,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
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
	// 90ABA848-AD74-1F6E-84BC-4182A7F1F29E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 22
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// example:
	//
	// 2
	TotalPage *int32 `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
}

func (s ListJobsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListJobsResponseBody) GoString() string {
	return s.String()
}

func (s *ListJobsResponseBody) SetCode(v int32) *ListJobsResponseBody {
	s.Code = &v
	return s
}

func (s *ListJobsResponseBody) SetDetails(v string) *ListJobsResponseBody {
	s.Details = &v
	return s
}

func (s *ListJobsResponseBody) SetErrorCode(v string) *ListJobsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListJobsResponseBody) SetJobs(v []*Job) *ListJobsResponseBody {
	s.Jobs = v
	return s
}

func (s *ListJobsResponseBody) SetMessage(v string) *ListJobsResponseBody {
	s.Message = &v
	return s
}

func (s *ListJobsResponseBody) SetPageNumber(v int32) *ListJobsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListJobsResponseBody) SetPageSize(v int32) *ListJobsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListJobsResponseBody) SetRequestId(v string) *ListJobsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListJobsResponseBody) SetSuccess(v bool) *ListJobsResponseBody {
	s.Success = &v
	return s
}

func (s *ListJobsResponseBody) SetTotalCount(v int32) *ListJobsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListJobsResponseBody) SetTotalPage(v int32) *ListJobsResponseBody {
	s.TotalPage = &v
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

type ListSubtaskItemsRequest struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListSubtaskItemsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSubtaskItemsRequest) GoString() string {
	return s.String()
}

func (s *ListSubtaskItemsRequest) SetPageNumber(v int32) *ListSubtaskItemsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListSubtaskItemsRequest) SetPageSize(v int32) *ListSubtaskItemsRequest {
	s.PageSize = &v
	return s
}

type ListSubtaskItemsResponseBody struct {
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// ""
	Details *string `json:"Details,omitempty" xml:"Details,omitempty"`
	// example:
	//
	// ""
	ErrorCode *string              `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Items     []*SubtaskItemDetail `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
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
	// 90ABA848-AD74-1F6E-84BC-4182A7F1F29E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 22
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// example:
	//
	// 2
	TotalPage *int32 `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
}

func (s ListSubtaskItemsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSubtaskItemsResponseBody) GoString() string {
	return s.String()
}

func (s *ListSubtaskItemsResponseBody) SetCode(v int32) *ListSubtaskItemsResponseBody {
	s.Code = &v
	return s
}

func (s *ListSubtaskItemsResponseBody) SetDetails(v string) *ListSubtaskItemsResponseBody {
	s.Details = &v
	return s
}

func (s *ListSubtaskItemsResponseBody) SetErrorCode(v string) *ListSubtaskItemsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListSubtaskItemsResponseBody) SetItems(v []*SubtaskItemDetail) *ListSubtaskItemsResponseBody {
	s.Items = v
	return s
}

func (s *ListSubtaskItemsResponseBody) SetMessage(v string) *ListSubtaskItemsResponseBody {
	s.Message = &v
	return s
}

func (s *ListSubtaskItemsResponseBody) SetPageNumber(v int32) *ListSubtaskItemsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListSubtaskItemsResponseBody) SetPageSize(v int32) *ListSubtaskItemsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListSubtaskItemsResponseBody) SetRequestId(v string) *ListSubtaskItemsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSubtaskItemsResponseBody) SetSuccess(v bool) *ListSubtaskItemsResponseBody {
	s.Success = &v
	return s
}

func (s *ListSubtaskItemsResponseBody) SetTotalCount(v int32) *ListSubtaskItemsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListSubtaskItemsResponseBody) SetTotalPage(v int32) *ListSubtaskItemsResponseBody {
	s.TotalPage = &v
	return s
}

type ListSubtaskItemsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSubtaskItemsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSubtaskItemsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSubtaskItemsResponse) GoString() string {
	return s.String()
}

func (s *ListSubtaskItemsResponse) SetHeaders(v map[string]*string) *ListSubtaskItemsResponse {
	s.Headers = v
	return s
}

func (s *ListSubtaskItemsResponse) SetStatusCode(v int32) *ListSubtaskItemsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSubtaskItemsResponse) SetBody(v *ListSubtaskItemsResponseBody) *ListSubtaskItemsResponse {
	s.Body = v
	return s
}

type ListSubtasksRequest struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListSubtasksRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSubtasksRequest) GoString() string {
	return s.String()
}

func (s *ListSubtasksRequest) SetPageNumber(v int32) *ListSubtasksRequest {
	s.PageNumber = &v
	return s
}

func (s *ListSubtasksRequest) SetPageSize(v int32) *ListSubtasksRequest {
	s.PageSize = &v
	return s
}

type ListSubtasksResponseBody struct {
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// null
	Details *string `json:"Details,omitempty" xml:"Details,omitempty"`
	// example:
	//
	// ""
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
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
	// 90ABA848-AD74-1F6E-84BC-4182A7F1F29E
	RequestId *string          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Subtasks  []*SubtaskDetail `json:"Subtasks,omitempty" xml:"Subtasks,omitempty" type:"Repeated"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 22
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// example:
	//
	// 2
	TotalPage *int32 `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
}

func (s ListSubtasksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSubtasksResponseBody) GoString() string {
	return s.String()
}

func (s *ListSubtasksResponseBody) SetCode(v int32) *ListSubtasksResponseBody {
	s.Code = &v
	return s
}

func (s *ListSubtasksResponseBody) SetDetails(v string) *ListSubtasksResponseBody {
	s.Details = &v
	return s
}

func (s *ListSubtasksResponseBody) SetErrorCode(v string) *ListSubtasksResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListSubtasksResponseBody) SetMessage(v string) *ListSubtasksResponseBody {
	s.Message = &v
	return s
}

func (s *ListSubtasksResponseBody) SetPageNumber(v int32) *ListSubtasksResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListSubtasksResponseBody) SetPageSize(v int32) *ListSubtasksResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListSubtasksResponseBody) SetRequestId(v string) *ListSubtasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSubtasksResponseBody) SetSubtasks(v []*SubtaskDetail) *ListSubtasksResponseBody {
	s.Subtasks = v
	return s
}

func (s *ListSubtasksResponseBody) SetSuccess(v bool) *ListSubtasksResponseBody {
	s.Success = &v
	return s
}

func (s *ListSubtasksResponseBody) SetTotalCount(v int32) *ListSubtasksResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListSubtasksResponseBody) SetTotalPage(v int32) *ListSubtasksResponseBody {
	s.TotalPage = &v
	return s
}

type ListSubtasksResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSubtasksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSubtasksResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSubtasksResponse) GoString() string {
	return s.String()
}

func (s *ListSubtasksResponse) SetHeaders(v map[string]*string) *ListSubtasksResponse {
	s.Headers = v
	return s
}

func (s *ListSubtasksResponse) SetStatusCode(v int32) *ListSubtasksResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSubtasksResponse) SetBody(v *ListSubtasksResponseBody) *ListSubtasksResponse {
	s.Body = v
	return s
}

type ListTasksRequest struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListTasksRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTasksRequest) GoString() string {
	return s.String()
}

func (s *ListTasksRequest) SetPageNumber(v int32) *ListTasksRequest {
	s.PageNumber = &v
	return s
}

func (s *ListTasksRequest) SetPageSize(v int32) *ListTasksRequest {
	s.PageSize = &v
	return s
}

type ListTasksResponseBody struct {
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// null
	Details *string `json:"Details,omitempty" xml:"Details,omitempty"`
	// example:
	//
	// ""
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
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
	// 90ABA848-AD74-1F6E-84BC-4182A7F1F29E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool         `json:"Success,omitempty" xml:"Success,omitempty"`
	Tasks   []*SimpleTask `json:"Tasks,omitempty" xml:"Tasks,omitempty" type:"Repeated"`
	// example:
	//
	// 22
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// example:
	//
	// 2
	TotalPage *int32 `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
}

func (s ListTasksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTasksResponseBody) GoString() string {
	return s.String()
}

func (s *ListTasksResponseBody) SetCode(v int32) *ListTasksResponseBody {
	s.Code = &v
	return s
}

func (s *ListTasksResponseBody) SetDetails(v string) *ListTasksResponseBody {
	s.Details = &v
	return s
}

func (s *ListTasksResponseBody) SetErrorCode(v string) *ListTasksResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListTasksResponseBody) SetMessage(v string) *ListTasksResponseBody {
	s.Message = &v
	return s
}

func (s *ListTasksResponseBody) SetPageNumber(v int32) *ListTasksResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListTasksResponseBody) SetPageSize(v int32) *ListTasksResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListTasksResponseBody) SetRequestId(v string) *ListTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTasksResponseBody) SetSuccess(v bool) *ListTasksResponseBody {
	s.Success = &v
	return s
}

func (s *ListTasksResponseBody) SetTasks(v []*SimpleTask) *ListTasksResponseBody {
	s.Tasks = v
	return s
}

func (s *ListTasksResponseBody) SetTotalCount(v int32) *ListTasksResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListTasksResponseBody) SetTotalPage(v int32) *ListTasksResponseBody {
	s.TotalPage = &v
	return s
}

type ListTasksResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTasksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTasksResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTasksResponse) GoString() string {
	return s.String()
}

func (s *ListTasksResponse) SetHeaders(v map[string]*string) *ListTasksResponse {
	s.Headers = v
	return s
}

func (s *ListTasksResponse) SetStatusCode(v int32) *ListTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTasksResponse) SetBody(v *ListTasksResponseBody) *ListTasksResponse {
	s.Body = v
	return s
}

type ListTemplatesRequest struct {
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
	// demo
	SearchKey *string   `json:"SearchKey,omitempty" xml:"SearchKey,omitempty"`
	Types     []*string `json:"Types,omitempty" xml:"Types,omitempty" type:"Repeated"`
}

func (s ListTemplatesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTemplatesRequest) GoString() string {
	return s.String()
}

func (s *ListTemplatesRequest) SetPageNumber(v int32) *ListTemplatesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListTemplatesRequest) SetPageSize(v int32) *ListTemplatesRequest {
	s.PageSize = &v
	return s
}

func (s *ListTemplatesRequest) SetSearchKey(v string) *ListTemplatesRequest {
	s.SearchKey = &v
	return s
}

func (s *ListTemplatesRequest) SetTypes(v []*string) *ListTemplatesRequest {
	s.Types = v
	return s
}

type ListTemplatesShrinkRequest struct {
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
	// demo
	SearchKey   *string `json:"SearchKey,omitempty" xml:"SearchKey,omitempty"`
	TypesShrink *string `json:"Types,omitempty" xml:"Types,omitempty"`
}

func (s ListTemplatesShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTemplatesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListTemplatesShrinkRequest) SetPageNumber(v int32) *ListTemplatesShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListTemplatesShrinkRequest) SetPageSize(v int32) *ListTemplatesShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListTemplatesShrinkRequest) SetSearchKey(v string) *ListTemplatesShrinkRequest {
	s.SearchKey = &v
	return s
}

func (s *ListTemplatesShrinkRequest) SetTypesShrink(v string) *ListTemplatesShrinkRequest {
	s.TypesShrink = &v
	return s
}

type ListTemplatesResponseBody struct {
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// null
	Details *string `json:"Details,omitempty" xml:"Details,omitempty"`
	// example:
	//
	// ""
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
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
	// 90ABA848-AD74-1F6E-84BC-4182A7F1F29E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success   *bool             `json:"Success,omitempty" xml:"Success,omitempty"`
	Templates []*SimpleTemplate `json:"Templates,omitempty" xml:"Templates,omitempty" type:"Repeated"`
	// example:
	//
	// 22
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// example:
	//
	// 2
	TotalPage *int32 `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
}

func (s ListTemplatesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *ListTemplatesResponseBody) SetCode(v int32) *ListTemplatesResponseBody {
	s.Code = &v
	return s
}

func (s *ListTemplatesResponseBody) SetDetails(v string) *ListTemplatesResponseBody {
	s.Details = &v
	return s
}

func (s *ListTemplatesResponseBody) SetErrorCode(v string) *ListTemplatesResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListTemplatesResponseBody) SetMessage(v string) *ListTemplatesResponseBody {
	s.Message = &v
	return s
}

func (s *ListTemplatesResponseBody) SetPageNumber(v int32) *ListTemplatesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListTemplatesResponseBody) SetPageSize(v int32) *ListTemplatesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListTemplatesResponseBody) SetRequestId(v string) *ListTemplatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTemplatesResponseBody) SetSuccess(v bool) *ListTemplatesResponseBody {
	s.Success = &v
	return s
}

func (s *ListTemplatesResponseBody) SetTemplates(v []*SimpleTemplate) *ListTemplatesResponseBody {
	s.Templates = v
	return s
}

func (s *ListTemplatesResponseBody) SetTotalCount(v int32) *ListTemplatesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListTemplatesResponseBody) SetTotalPage(v int32) *ListTemplatesResponseBody {
	s.TotalPage = &v
	return s
}

type ListTemplatesResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTemplatesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *ListTemplatesResponse) SetStatusCode(v int32) *ListTemplatesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTemplatesResponse) SetBody(v *ListTemplatesResponseBody) *ListTemplatesResponse {
	s.Body = v
	return s
}

type ListTenantsRequest struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListTenantsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTenantsRequest) GoString() string {
	return s.String()
}

func (s *ListTenantsRequest) SetPageNumber(v int32) *ListTenantsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListTenantsRequest) SetPageSize(v int32) *ListTenantsRequest {
	s.PageSize = &v
	return s
}

type ListTenantsResponseBody struct {
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// -
	Details *string `json:"Details,omitempty" xml:"Details,omitempty"`
	// example:
	//
	// -
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
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
	// 90ABA848-AD74-1F6E-84BC-4182A7F1F29E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool           `json:"Success,omitempty" xml:"Success,omitempty"`
	Tenants []*SimpleTenant `json:"Tenants,omitempty" xml:"Tenants,omitempty" type:"Repeated"`
	// example:
	//
	// 22
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// example:
	//
	// 2
	TotalPage *int32 `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
}

func (s ListTenantsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTenantsResponseBody) GoString() string {
	return s.String()
}

func (s *ListTenantsResponseBody) SetCode(v int32) *ListTenantsResponseBody {
	s.Code = &v
	return s
}

func (s *ListTenantsResponseBody) SetDetails(v string) *ListTenantsResponseBody {
	s.Details = &v
	return s
}

func (s *ListTenantsResponseBody) SetErrorCode(v string) *ListTenantsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListTenantsResponseBody) SetMessage(v string) *ListTenantsResponseBody {
	s.Message = &v
	return s
}

func (s *ListTenantsResponseBody) SetPageNumber(v int32) *ListTenantsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListTenantsResponseBody) SetPageSize(v int32) *ListTenantsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListTenantsResponseBody) SetRequestId(v string) *ListTenantsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTenantsResponseBody) SetSuccess(v bool) *ListTenantsResponseBody {
	s.Success = &v
	return s
}

func (s *ListTenantsResponseBody) SetTenants(v []*SimpleTenant) *ListTenantsResponseBody {
	s.Tenants = v
	return s
}

func (s *ListTenantsResponseBody) SetTotalCount(v int32) *ListTenantsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListTenantsResponseBody) SetTotalPage(v int32) *ListTenantsResponseBody {
	s.TotalPage = &v
	return s
}

type ListTenantsResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTenantsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTenantsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTenantsResponse) GoString() string {
	return s.String()
}

func (s *ListTenantsResponse) SetHeaders(v map[string]*string) *ListTenantsResponse {
	s.Headers = v
	return s
}

func (s *ListTenantsResponse) SetStatusCode(v int32) *ListTenantsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTenantsResponse) SetBody(v *ListTenantsResponseBody) *ListTenantsResponse {
	s.Body = v
	return s
}

type ListUsersRequest struct {
	// example:
	//
	// 20
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 1
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListUsersRequest) String() string {
	return tea.Prettify(s)
}

func (s ListUsersRequest) GoString() string {
	return s.String()
}

func (s *ListUsersRequest) SetPageNumber(v int32) *ListUsersRequest {
	s.PageNumber = &v
	return s
}

func (s *ListUsersRequest) SetPageSize(v int32) *ListUsersRequest {
	s.PageSize = &v
	return s
}

type ListUsersResponseBody struct {
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// null
	Details *string `json:"Details,omitempty" xml:"Details,omitempty"`
	// example:
	//
	// ""
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
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
	// 90ABA848-AD74-1F6E-84BC-4182A7F1F29E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 22
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// example:
	//
	// 2
	TotalPage *int32        `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
	Users     []*SimpleUser `json:"Users,omitempty" xml:"Users,omitempty" type:"Repeated"`
}

func (s ListUsersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListUsersResponseBody) GoString() string {
	return s.String()
}

func (s *ListUsersResponseBody) SetCode(v int32) *ListUsersResponseBody {
	s.Code = &v
	return s
}

func (s *ListUsersResponseBody) SetDetails(v string) *ListUsersResponseBody {
	s.Details = &v
	return s
}

func (s *ListUsersResponseBody) SetErrorCode(v string) *ListUsersResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListUsersResponseBody) SetMessage(v string) *ListUsersResponseBody {
	s.Message = &v
	return s
}

func (s *ListUsersResponseBody) SetPageNumber(v int32) *ListUsersResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListUsersResponseBody) SetPageSize(v int32) *ListUsersResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListUsersResponseBody) SetRequestId(v string) *ListUsersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUsersResponseBody) SetSuccess(v bool) *ListUsersResponseBody {
	s.Success = &v
	return s
}

func (s *ListUsersResponseBody) SetTotalCount(v int32) *ListUsersResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListUsersResponseBody) SetTotalPage(v int32) *ListUsersResponseBody {
	s.TotalPage = &v
	return s
}

func (s *ListUsersResponseBody) SetUsers(v []*SimpleUser) *ListUsersResponseBody {
	s.Users = v
	return s
}

type ListUsersResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListUsersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *ListUsersResponse) SetStatusCode(v int32) *ListUsersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListUsersResponse) SetBody(v *ListUsersResponseBody) *ListUsersResponse {
	s.Body = v
	return s
}

type RemoveWorkNodeWorkforceRequest struct {
	UserIds []*int64 `json:"UserIds,omitempty" xml:"UserIds,omitempty" type:"Repeated"`
}

func (s RemoveWorkNodeWorkforceRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveWorkNodeWorkforceRequest) GoString() string {
	return s.String()
}

func (s *RemoveWorkNodeWorkforceRequest) SetUserIds(v []*int64) *RemoveWorkNodeWorkforceRequest {
	s.UserIds = v
	return s
}

type RemoveWorkNodeWorkforceResponseBody struct {
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// ""
	Details *string `json:"Details,omitempty" xml:"Details,omitempty"`
	// example:
	//
	// ""
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 90ABA848-AD74-1F6E-84BC-4182A7F1F29E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RemoveWorkNodeWorkforceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveWorkNodeWorkforceResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveWorkNodeWorkforceResponseBody) SetCode(v int32) *RemoveWorkNodeWorkforceResponseBody {
	s.Code = &v
	return s
}

func (s *RemoveWorkNodeWorkforceResponseBody) SetDetails(v string) *RemoveWorkNodeWorkforceResponseBody {
	s.Details = &v
	return s
}

func (s *RemoveWorkNodeWorkforceResponseBody) SetErrorCode(v string) *RemoveWorkNodeWorkforceResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *RemoveWorkNodeWorkforceResponseBody) SetMessage(v string) *RemoveWorkNodeWorkforceResponseBody {
	s.Message = &v
	return s
}

func (s *RemoveWorkNodeWorkforceResponseBody) SetRequestId(v string) *RemoveWorkNodeWorkforceResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveWorkNodeWorkforceResponseBody) SetSuccess(v bool) *RemoveWorkNodeWorkforceResponseBody {
	s.Success = &v
	return s
}

type RemoveWorkNodeWorkforceResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveWorkNodeWorkforceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveWorkNodeWorkforceResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveWorkNodeWorkforceResponse) GoString() string {
	return s.String()
}

func (s *RemoveWorkNodeWorkforceResponse) SetHeaders(v map[string]*string) *RemoveWorkNodeWorkforceResponse {
	s.Headers = v
	return s
}

func (s *RemoveWorkNodeWorkforceResponse) SetStatusCode(v int32) *RemoveWorkNodeWorkforceResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveWorkNodeWorkforceResponse) SetBody(v *RemoveWorkNodeWorkforceResponseBody) *RemoveWorkNodeWorkforceResponse {
	s.Body = v
	return s
}

type UpdateTaskRequest struct {
	// This parameter is required.
	Body *UpdateTaskDTO `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateTaskRequest) GoString() string {
	return s.String()
}

func (s *UpdateTaskRequest) SetBody(v *UpdateTaskDTO) *UpdateTaskRequest {
	s.Body = v
	return s
}

type UpdateTaskResponseBody struct {
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// null
	Details *string `json:"Details,omitempty" xml:"Details,omitempty"`
	// example:
	//
	// ""
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 90ABA848-AD74-1F6E-84BC-4182A7F1F29E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateTaskResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateTaskResponseBody) SetCode(v int32) *UpdateTaskResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateTaskResponseBody) SetDetails(v string) *UpdateTaskResponseBody {
	s.Details = &v
	return s
}

func (s *UpdateTaskResponseBody) SetErrorCode(v string) *UpdateTaskResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateTaskResponseBody) SetMessage(v string) *UpdateTaskResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateTaskResponseBody) SetRequestId(v string) *UpdateTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateTaskResponseBody) SetSuccess(v bool) *UpdateTaskResponseBody {
	s.Success = &v
	return s
}

type UpdateTaskResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateTaskResponse) GoString() string {
	return s.String()
}

func (s *UpdateTaskResponse) SetHeaders(v map[string]*string) *UpdateTaskResponse {
	s.Headers = v
	return s
}

func (s *UpdateTaskResponse) SetStatusCode(v int32) *UpdateTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateTaskResponse) SetBody(v *UpdateTaskResponseBody) *UpdateTaskResponse {
	s.Body = v
	return s
}

type UpdateTaskWorkforceRequest struct {
	Workforce []*SimpleWorkforce `json:"Workforce,omitempty" xml:"Workforce,omitempty" type:"Repeated"`
}

func (s UpdateTaskWorkforceRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateTaskWorkforceRequest) GoString() string {
	return s.String()
}

func (s *UpdateTaskWorkforceRequest) SetWorkforce(v []*SimpleWorkforce) *UpdateTaskWorkforceRequest {
	s.Workforce = v
	return s
}

type UpdateTaskWorkforceResponseBody struct {
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// ""
	Details *string `json:"Details,omitempty" xml:"Details,omitempty"`
	// example:
	//
	// ""
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 90ABA848-AD74-1F6E-84BC-4182A7F1F29E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateTaskWorkforceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateTaskWorkforceResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateTaskWorkforceResponseBody) SetCode(v int32) *UpdateTaskWorkforceResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateTaskWorkforceResponseBody) SetDetails(v string) *UpdateTaskWorkforceResponseBody {
	s.Details = &v
	return s
}

func (s *UpdateTaskWorkforceResponseBody) SetErrorCode(v string) *UpdateTaskWorkforceResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateTaskWorkforceResponseBody) SetMessage(v string) *UpdateTaskWorkforceResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateTaskWorkforceResponseBody) SetRequestId(v string) *UpdateTaskWorkforceResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateTaskWorkforceResponseBody) SetSuccess(v bool) *UpdateTaskWorkforceResponseBody {
	s.Success = &v
	return s
}

type UpdateTaskWorkforceResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateTaskWorkforceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateTaskWorkforceResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateTaskWorkforceResponse) GoString() string {
	return s.String()
}

func (s *UpdateTaskWorkforceResponse) SetHeaders(v map[string]*string) *UpdateTaskWorkforceResponse {
	s.Headers = v
	return s
}

func (s *UpdateTaskWorkforceResponse) SetStatusCode(v int32) *UpdateTaskWorkforceResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateTaskWorkforceResponse) SetBody(v *UpdateTaskWorkforceResponseBody) *UpdateTaskWorkforceResponse {
	s.Body = v
	return s
}

type UpdateTemplateRequest struct {
	Body *TemplateDTO `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateTemplateRequest) GoString() string {
	return s.String()
}

func (s *UpdateTemplateRequest) SetBody(v *TemplateDTO) *UpdateTemplateRequest {
	s.Body = v
	return s
}

type UpdateTemplateResponseBody struct {
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// null
	Details *string `json:"Details,omitempty" xml:"Details,omitempty"`
	// example:
	//
	// ""
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 90ABA848-AD74-1F6E-84BC-4182A7F1F29E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 1529360348342353920
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s UpdateTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateTemplateResponseBody) SetCode(v int32) *UpdateTemplateResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateTemplateResponseBody) SetDetails(v string) *UpdateTemplateResponseBody {
	s.Details = &v
	return s
}

func (s *UpdateTemplateResponseBody) SetErrorCode(v string) *UpdateTemplateResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateTemplateResponseBody) SetMessage(v string) *UpdateTemplateResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateTemplateResponseBody) SetRequestId(v string) *UpdateTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateTemplateResponseBody) SetSuccess(v bool) *UpdateTemplateResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateTemplateResponseBody) SetTemplateId(v string) *UpdateTemplateResponseBody {
	s.TemplateId = &v
	return s
}

type UpdateTemplateResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *UpdateTemplateResponse) SetStatusCode(v int32) *UpdateTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateTemplateResponse) SetBody(v *UpdateTemplateResponseBody) *UpdateTemplateResponse {
	s.Body = v
	return s
}

type UpdateTenantRequest struct {
	// example:
	//
	// demo
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	TenantName  *string `json:"TenantName,omitempty" xml:"TenantName,omitempty"`
}

func (s UpdateTenantRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateTenantRequest) GoString() string {
	return s.String()
}

func (s *UpdateTenantRequest) SetDescription(v string) *UpdateTenantRequest {
	s.Description = &v
	return s
}

func (s *UpdateTenantRequest) SetTenantName(v string) *UpdateTenantRequest {
	s.TenantName = &v
	return s
}

type UpdateTenantResponseBody struct {
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// null
	Details *string `json:"Details,omitempty" xml:"Details,omitempty"`
	// example:
	//
	// ""
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 90ABA848-AD74-1F6E-84BC-4182A7F1F29E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateTenantResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateTenantResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateTenantResponseBody) SetCode(v int32) *UpdateTenantResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateTenantResponseBody) SetDetails(v string) *UpdateTenantResponseBody {
	s.Details = &v
	return s
}

func (s *UpdateTenantResponseBody) SetErrorCode(v string) *UpdateTenantResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateTenantResponseBody) SetMessage(v string) *UpdateTenantResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateTenantResponseBody) SetRequestId(v string) *UpdateTenantResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateTenantResponseBody) SetSuccess(v bool) *UpdateTenantResponseBody {
	s.Success = &v
	return s
}

type UpdateTenantResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateTenantResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateTenantResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateTenantResponse) GoString() string {
	return s.String()
}

func (s *UpdateTenantResponse) SetHeaders(v map[string]*string) *UpdateTenantResponse {
	s.Headers = v
	return s
}

func (s *UpdateTenantResponse) SetStatusCode(v int32) *UpdateTenantResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateTenantResponse) SetBody(v *UpdateTenantResponseBody) *UpdateTenantResponse {
	s.Body = v
	return s
}

type UpdateUserRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ADMIN
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// user1
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s UpdateUserRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateUserRequest) GoString() string {
	return s.String()
}

func (s *UpdateUserRequest) SetRole(v string) *UpdateUserRequest {
	s.Role = &v
	return s
}

func (s *UpdateUserRequest) SetUserName(v string) *UpdateUserRequest {
	s.UserName = &v
	return s
}

type UpdateUserResponseBody struct {
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// null
	Details *string `json:"Details,omitempty" xml:"Details,omitempty"`
	// example:
	//
	// ""
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 90ABA848-AD74-1F6E-84BC-4182A7F1F29E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 166***980757311
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s UpdateUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateUserResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateUserResponseBody) SetCode(v int32) *UpdateUserResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateUserResponseBody) SetDetails(v string) *UpdateUserResponseBody {
	s.Details = &v
	return s
}

func (s *UpdateUserResponseBody) SetErrorCode(v string) *UpdateUserResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateUserResponseBody) SetMessage(v string) *UpdateUserResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateUserResponseBody) SetRequestId(v string) *UpdateUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateUserResponseBody) SetSuccess(v bool) *UpdateUserResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateUserResponseBody) SetUserId(v string) *UpdateUserResponseBody {
	s.UserId = &v
	return s
}

type UpdateUserResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateUserResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateUserResponse) GoString() string {
	return s.String()
}

func (s *UpdateUserResponse) SetHeaders(v map[string]*string) *UpdateUserResponse {
	s.Headers = v
	return s
}

func (s *UpdateUserResponse) SetStatusCode(v int32) *UpdateUserResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateUserResponse) SetBody(v *UpdateUserResponseBody) *UpdateUserResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("openitag"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// 增加结点任务人力
//
// @param request - AddWorkNodeWorkforceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddWorkNodeWorkforceResponse
func (client *Client) AddWorkNodeWorkforceWithOptions(TenantId *string, TaskId *string, WorkNodeId *string, request *AddWorkNodeWorkforceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *AddWorkNodeWorkforceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UserIds)) {
		body["UserIds"] = request.UserIds
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("AddWorkNodeWorkforce"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/api/v1/tenants/" + tea.StringValue(openapiutil.GetEncodeParam(TenantId)) + "/tasks/" + tea.StringValue(openapiutil.GetEncodeParam(TaskId)) + "/worknodes/" + tea.StringValue(openapiutil.GetEncodeParam(WorkNodeId)) + "/workforce"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &AddWorkNodeWorkforceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 增加结点任务人力
//
// @param request - AddWorkNodeWorkforceRequest
//
// @return AddWorkNodeWorkforceResponse
func (client *Client) AddWorkNodeWorkforce(TenantId *string, TaskId *string, WorkNodeId *string, request *AddWorkNodeWorkforceRequest) (_result *AddWorkNodeWorkforceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &AddWorkNodeWorkforceResponse{}
	_body, _err := client.AddWorkNodeWorkforceWithOptions(TenantId, TaskId, WorkNodeId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 数据追加
//
// @param request - AppendAllDataToTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AppendAllDataToTaskResponse
func (client *Client) AppendAllDataToTaskWithOptions(TenantId *string, TaskId *string, request *AppendAllDataToTaskRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *AppendAllDataToTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapi.Params{
		Action:      tea.String("AppendAllDataToTask"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/api/v1/tenants/" + tea.StringValue(openapiutil.GetEncodeParam(TenantId)) + "/tasks/" + tea.StringValue(openapiutil.GetEncodeParam(TaskId)) + "/appendAllDataToTask"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &AppendAllDataToTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 数据追加
//
// @param request - AppendAllDataToTaskRequest
//
// @return AppendAllDataToTaskResponse
func (client *Client) AppendAllDataToTask(TenantId *string, TaskId *string, request *AppendAllDataToTaskRequest) (_result *AppendAllDataToTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &AppendAllDataToTaskResponse{}
	_body, _err := client.AppendAllDataToTaskWithOptions(TenantId, TaskId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建标注任务
//
// @param request - CreateTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTaskResponse
func (client *Client) CreateTaskWithOptions(TenantId *string, request *CreateTaskRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateTask"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/api/v1/tenants/" + tea.StringValue(openapiutil.GetEncodeParam(TenantId)) + "/tasks"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建标注任务
//
// @param request - CreateTaskRequest
//
// @return CreateTaskResponse
func (client *Client) CreateTask(TenantId *string, request *CreateTaskRequest) (_result *CreateTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateTaskResponse{}
	_body, _err := client.CreateTaskWithOptions(TenantId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建标注模版
//
// @param request - CreateTemplateRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTemplateResponse
func (client *Client) CreateTemplateWithOptions(TenantId *string, request *CreateTemplateRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateTemplate"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/api/v1/tenants/" + tea.StringValue(openapiutil.GetEncodeParam(TenantId)) + "/templates"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建标注模版
//
// @param request - CreateTemplateRequest
//
// @return CreateTemplateResponse
func (client *Client) CreateTemplate(TenantId *string, request *CreateTemplateRequest) (_result *CreateTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateTemplateResponse{}
	_body, _err := client.CreateTemplateWithOptions(TenantId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建租户内用户
//
// @param request - CreateUserRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateUserResponse
func (client *Client) CreateUserWithOptions(TenantId *string, request *CreateUserRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountNo)) {
		body["AccountNo"] = request.AccountNo
	}

	if !tea.BoolValue(util.IsUnset(request.AccountType)) {
		body["AccountType"] = request.AccountType
	}

	if !tea.BoolValue(util.IsUnset(request.Role)) {
		body["Role"] = request.Role
	}

	if !tea.BoolValue(util.IsUnset(request.UserName)) {
		body["UserName"] = request.UserName
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateUser"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/api/v1/tenants/" + tea.StringValue(openapiutil.GetEncodeParam(TenantId)) + "/users"),
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

// Summary:
//
// 创建租户内用户
//
// @param request - CreateUserRequest
//
// @return CreateUserResponse
func (client *Client) CreateUser(TenantId *string, request *CreateUserRequest) (_result *CreateUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateUserResponse{}
	_body, _err := client.CreateUserWithOptions(TenantId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除任务
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteTaskResponse
func (client *Client) DeleteTaskWithOptions(TenantId *string, TaskId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteTaskResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteTask"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/api/v1/tenants/" + tea.StringValue(openapiutil.GetEncodeParam(TenantId)) + "/tasks/" + tea.StringValue(openapiutil.GetEncodeParam(TaskId))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除任务
//
// @return DeleteTaskResponse
func (client *Client) DeleteTask(TenantId *string, TaskId *string) (_result *DeleteTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteTaskResponse{}
	_body, _err := client.DeleteTaskWithOptions(TenantId, TaskId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除租户下的单个模板
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteTemplateResponse
func (client *Client) DeleteTemplateWithOptions(TenantId *string, TemplateId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteTemplateResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteTemplate"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/api/v1/tenants/" + tea.StringValue(openapiutil.GetEncodeParam(TenantId)) + "/templates/" + tea.StringValue(openapiutil.GetEncodeParam(TemplateId))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除租户下的单个模板
//
// @return DeleteTemplateResponse
func (client *Client) DeleteTemplate(TenantId *string, TemplateId *string) (_result *DeleteTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteTemplateResponse{}
	_body, _err := client.DeleteTemplateWithOptions(TenantId, TemplateId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除用户
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteUserResponse
func (client *Client) DeleteUserWithOptions(TenantId *string, UserId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteUserResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteUser"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/api/v1/tenants/" + tea.StringValue(openapiutil.GetEncodeParam(TenantId)) + "/users/" + tea.StringValue(openapiutil.GetEncodeParam(UserId))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除用户
//
// @return DeleteUserResponse
func (client *Client) DeleteUser(TenantId *string, UserId *string) (_result *DeleteUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteUserResponse{}
	_body, _err := client.DeleteUserWithOptions(TenantId, UserId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取任务导出结果
//
// @param request - ExportAnnotationsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExportAnnotationsResponse
func (client *Client) ExportAnnotationsWithOptions(TenantId *string, TaskId *string, request *ExportAnnotationsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ExportAnnotationsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OssPath)) {
		query["OssPath"] = request.OssPath
	}

	if !tea.BoolValue(util.IsUnset(request.RegisterDataset)) {
		query["RegisterDataset"] = request.RegisterDataset
	}

	if !tea.BoolValue(util.IsUnset(request.Target)) {
		query["Target"] = request.Target
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ExportAnnotations"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/api/v1/tenants/" + tea.StringValue(openapiutil.GetEncodeParam(TenantId)) + "/tasks/" + tea.StringValue(openapiutil.GetEncodeParam(TaskId)) + "/annotations/export"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ExportAnnotationsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取任务导出结果
//
// @param request - ExportAnnotationsRequest
//
// @return ExportAnnotationsResponse
func (client *Client) ExportAnnotations(TenantId *string, TaskId *string, request *ExportAnnotationsRequest) (_result *ExportAnnotationsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ExportAnnotationsResponse{}
	_body, _err := client.ExportAnnotationsWithOptions(TenantId, TaskId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取异步任务Job
//
// @param request - GetJobRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetJobResponse
func (client *Client) GetJobWithOptions(TenantId *string, JobId *string, request *GetJobRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JobType)) {
		query["JobType"] = request.JobType
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetJob"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/api/v1/tenants/" + tea.StringValue(openapiutil.GetEncodeParam(TenantId)) + "/jobs/" + tea.StringValue(openapiutil.GetEncodeParam(JobId))),
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
// 获取异步任务Job
//
// @param request - GetJobRequest
//
// @return GetJobResponse
func (client *Client) GetJob(TenantId *string, JobId *string, request *GetJobRequest) (_result *GetJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetJobResponse{}
	_body, _err := client.GetJobWithOptions(TenantId, JobId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取单个子任务信息
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSubtaskResponse
func (client *Client) GetSubtaskWithOptions(TenantId *string, TaskID *string, SubtaskId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetSubtaskResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetSubtask"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/api/v1/tenants/" + tea.StringValue(openapiutil.GetEncodeParam(TenantId)) + "/tasks/" + tea.StringValue(openapiutil.GetEncodeParam(TaskID)) + "/subtasks/" + tea.StringValue(openapiutil.GetEncodeParam(SubtaskId))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetSubtaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取单个子任务信息
//
// @return GetSubtaskResponse
func (client *Client) GetSubtask(TenantId *string, TaskID *string, SubtaskId *string) (_result *GetSubtaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetSubtaskResponse{}
	_body, _err := client.GetSubtaskWithOptions(TenantId, TaskID, SubtaskId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取子任务单个ITEM信息
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSubtaskItemResponse
func (client *Client) GetSubtaskItemWithOptions(TenantId *string, TaskId *string, SubtaskId *string, ItemId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetSubtaskItemResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetSubtaskItem"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/api/v1/tenants/" + tea.StringValue(openapiutil.GetEncodeParam(TenantId)) + "/tasks/" + tea.StringValue(openapiutil.GetEncodeParam(TaskId)) + "/subtasks/" + tea.StringValue(openapiutil.GetEncodeParam(SubtaskId)) + "/items/" + tea.StringValue(openapiutil.GetEncodeParam(ItemId))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetSubtaskItemResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取子任务单个ITEM信息
//
// @return GetSubtaskItemResponse
func (client *Client) GetSubtaskItem(TenantId *string, TaskId *string, SubtaskId *string, ItemId *string) (_result *GetSubtaskItemResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetSubtaskItemResponse{}
	_body, _err := client.GetSubtaskItemWithOptions(TenantId, TaskId, SubtaskId, ItemId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取任务状态信息
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTaskResponse
func (client *Client) GetTaskWithOptions(TenantId *string, TaskId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetTaskResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetTask"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/api/v1/tenants/" + tea.StringValue(openapiutil.GetEncodeParam(TenantId)) + "/tasks/" + tea.StringValue(openapiutil.GetEncodeParam(TaskId))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取任务状态信息
//
// @return GetTaskResponse
func (client *Client) GetTask(TenantId *string, TaskId *string) (_result *GetTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetTaskResponse{}
	_body, _err := client.GetTaskWithOptions(TenantId, TaskId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取任务统计信息
//
// @param request - GetTaskStatisticsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTaskStatisticsResponse
func (client *Client) GetTaskStatisticsWithOptions(TenantId *string, TaskId *string, request *GetTaskStatisticsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetTaskStatisticsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.StatType)) {
		query["StatType"] = request.StatType
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetTaskStatistics"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/api/v1/tenants/" + tea.StringValue(openapiutil.GetEncodeParam(TenantId)) + "/tasks/" + tea.StringValue(openapiutil.GetEncodeParam(TaskId)) + "/statistics"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetTaskStatisticsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取任务统计信息
//
// @param request - GetTaskStatisticsRequest
//
// @return GetTaskStatisticsResponse
func (client *Client) GetTaskStatistics(TenantId *string, TaskId *string, request *GetTaskStatisticsRequest) (_result *GetTaskStatisticsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetTaskStatisticsResponse{}
	_body, _err := client.GetTaskStatisticsWithOptions(TenantId, TaskId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取任务状态信息
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTaskStatusResponse
func (client *Client) GetTaskStatusWithOptions(TenantId *string, TaskId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetTaskStatusResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetTaskStatus"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/api/v1/tenants/" + tea.StringValue(openapiutil.GetEncodeParam(TenantId)) + "/tasks/" + tea.StringValue(openapiutil.GetEncodeParam(TaskId)) + "/status"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetTaskStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取任务状态信息
//
// @return GetTaskStatusResponse
func (client *Client) GetTaskStatus(TenantId *string, TaskId *string) (_result *GetTaskStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetTaskStatusResponse{}
	_body, _err := client.GetTaskStatusWithOptions(TenantId, TaskId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取任务模版信息
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTaskTemplateResponse
func (client *Client) GetTaskTemplateWithOptions(TenantId *string, TaskId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetTaskTemplateResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetTaskTemplate"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/api/v1/tenants/" + tea.StringValue(openapiutil.GetEncodeParam(TenantId)) + "/tasks/" + tea.StringValue(openapiutil.GetEncodeParam(TaskId)) + "/template"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetTaskTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取任务模版信息
//
// @return GetTaskTemplateResponse
func (client *Client) GetTaskTemplate(TenantId *string, TaskId *string) (_result *GetTaskTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetTaskTemplateResponse{}
	_body, _err := client.GetTaskTemplateWithOptions(TenantId, TaskId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取任务题目信息
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTaskTemplateQuestionsResponse
func (client *Client) GetTaskTemplateQuestionsWithOptions(TenantId *string, TaskId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetTaskTemplateQuestionsResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetTaskTemplateQuestions"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/api/v1/tenants/" + tea.StringValue(openapiutil.GetEncodeParam(TenantId)) + "/tasks/" + tea.StringValue(openapiutil.GetEncodeParam(TaskId)) + "/template/questions"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetTaskTemplateQuestionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取任务题目信息
//
// @return GetTaskTemplateQuestionsResponse
func (client *Client) GetTaskTemplateQuestions(TenantId *string, TaskId *string) (_result *GetTaskTemplateQuestionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetTaskTemplateQuestionsResponse{}
	_body, _err := client.GetTaskTemplateQuestionsWithOptions(TenantId, TaskId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取任务题目信息
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTaskTemplateViewsResponse
func (client *Client) GetTaskTemplateViewsWithOptions(TenantId *string, TaskId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetTaskTemplateViewsResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetTaskTemplateViews"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/api/v1/tenants/" + tea.StringValue(openapiutil.GetEncodeParam(TenantId)) + "/tasks/" + tea.StringValue(openapiutil.GetEncodeParam(TaskId)) + "/template/views"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetTaskTemplateViewsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取任务题目信息
//
// @return GetTaskTemplateViewsResponse
func (client *Client) GetTaskTemplateViews(TenantId *string, TaskId *string) (_result *GetTaskTemplateViewsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetTaskTemplateViewsResponse{}
	_body, _err := client.GetTaskTemplateViewsWithOptions(TenantId, TaskId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取任务人力
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTaskWorkforceResponse
func (client *Client) GetTaskWorkforceWithOptions(TenantId *string, TaskId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetTaskWorkforceResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetTaskWorkforce"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/api/v1/tenants/" + tea.StringValue(openapiutil.GetEncodeParam(TenantId)) + "/tasks/" + tea.StringValue(openapiutil.GetEncodeParam(TaskId)) + "/workforce"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetTaskWorkforceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取任务人力
//
// @return GetTaskWorkforceResponse
func (client *Client) GetTaskWorkforce(TenantId *string, TaskId *string) (_result *GetTaskWorkforceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetTaskWorkforceResponse{}
	_body, _err := client.GetTaskWorkforceWithOptions(TenantId, TaskId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取任务人力统计信息
//
// @param request - GetTaskWorkforceStatisticRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTaskWorkforceStatisticResponse
func (client *Client) GetTaskWorkforceStatisticWithOptions(TenantId *string, TaskId *string, request *GetTaskWorkforceStatisticRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetTaskWorkforceStatisticResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.StatType)) {
		query["StatType"] = request.StatType
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetTaskWorkforceStatistic"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/api/v1/tenants/" + tea.StringValue(openapiutil.GetEncodeParam(TenantId)) + "/tasks/" + tea.StringValue(openapiutil.GetEncodeParam(TaskId)) + "/workforce/statistic"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetTaskWorkforceStatisticResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取任务人力统计信息
//
// @param request - GetTaskWorkforceStatisticRequest
//
// @return GetTaskWorkforceStatisticResponse
func (client *Client) GetTaskWorkforceStatistic(TenantId *string, TaskId *string, request *GetTaskWorkforceStatisticRequest) (_result *GetTaskWorkforceStatisticResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetTaskWorkforceStatisticResponse{}
	_body, _err := client.GetTaskWorkforceStatisticWithOptions(TenantId, TaskId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取租户下单个模板
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTemplateResponse
func (client *Client) GetTemplateWithOptions(TenantId *string, TemplateId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetTemplateResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetTemplate"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/api/v1/tenants/" + tea.StringValue(openapiutil.GetEncodeParam(TenantId)) + "/templates/" + tea.StringValue(openapiutil.GetEncodeParam(TemplateId))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取租户下单个模板
//
// @return GetTemplateResponse
func (client *Client) GetTemplate(TenantId *string, TemplateId *string) (_result *GetTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetTemplateResponse{}
	_body, _err := client.GetTemplateWithOptions(TenantId, TemplateId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取租户下单个模板问题
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTemplateQuestionsResponse
func (client *Client) GetTemplateQuestionsWithOptions(TenantId *string, TemplateId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetTemplateQuestionsResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetTemplateQuestions"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/api/v1/tenants/" + tea.StringValue(openapiutil.GetEncodeParam(TenantId)) + "/templates/" + tea.StringValue(openapiutil.GetEncodeParam(TemplateId)) + "/questions"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetTemplateQuestionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取租户下单个模板问题
//
// @return GetTemplateQuestionsResponse
func (client *Client) GetTemplateQuestions(TenantId *string, TemplateId *string) (_result *GetTemplateQuestionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetTemplateQuestionsResponse{}
	_body, _err := client.GetTemplateQuestionsWithOptions(TenantId, TemplateId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取租户下模板视图
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTemplateViewResponse
func (client *Client) GetTemplateViewWithOptions(TenantId *string, TemplateId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetTemplateViewResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetTemplateView"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/api/v1/tenants/" + tea.StringValue(openapiutil.GetEncodeParam(TenantId)) + "/templates/" + tea.StringValue(openapiutil.GetEncodeParam(TemplateId)) + "/views"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetTemplateViewResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取租户下模板视图
//
// @return GetTemplateViewResponse
func (client *Client) GetTemplateView(TenantId *string, TemplateId *string) (_result *GetTemplateViewResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetTemplateViewResponse{}
	_body, _err := client.GetTemplateViewWithOptions(TenantId, TemplateId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取租户信息
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTenantResponse
func (client *Client) GetTenantWithOptions(TenantId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetTenantResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetTenant"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/api/v1/tenants/" + tea.StringValue(openapiutil.GetEncodeParam(TenantId))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetTenantResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取租户信息
//
// @return GetTenantResponse
func (client *Client) GetTenant(TenantId *string) (_result *GetTenantResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetTenantResponse{}
	_body, _err := client.GetTenantWithOptions(TenantId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取用户
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUserResponse
func (client *Client) GetUserWithOptions(TenantId *string, UserId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetUserResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetUser"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/api/v1/tenants/" + tea.StringValue(openapiutil.GetEncodeParam(TenantId)) + "/users/" + tea.StringValue(openapiutil.GetEncodeParam(UserId))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取用户
//
// @return GetUserResponse
func (client *Client) GetUser(TenantId *string, UserId *string) (_result *GetUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetUserResponse{}
	_body, _err := client.GetUserWithOptions(TenantId, UserId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取异步任务Job列表
//
// @param request - ListJobsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListJobsResponse
func (client *Client) ListJobsWithOptions(TenantId *string, request *ListJobsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListJobsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JobType)) {
		query["JobType"] = request.JobType
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListJobs"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/api/v1/tenants/" + tea.StringValue(openapiutil.GetEncodeParam(TenantId)) + "/jobs"),
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
// 获取异步任务Job列表
//
// @param request - ListJobsRequest
//
// @return ListJobsResponse
func (client *Client) ListJobs(TenantId *string, request *ListJobsRequest) (_result *ListJobsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListJobsResponse{}
	_body, _err := client.ListJobsWithOptions(TenantId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取子任务ITEM列表页信息
//
// @param request - ListSubtaskItemsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSubtaskItemsResponse
func (client *Client) ListSubtaskItemsWithOptions(TenantId *string, TaskID *string, SubtaskId *string, request *ListSubtaskItemsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListSubtaskItemsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListSubtaskItems"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/api/v1/tenants/" + tea.StringValue(openapiutil.GetEncodeParam(TenantId)) + "/tasks/" + tea.StringValue(openapiutil.GetEncodeParam(TaskID)) + "/subtasks/" + tea.StringValue(openapiutil.GetEncodeParam(SubtaskId)) + "/items"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListSubtaskItemsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取子任务ITEM列表页信息
//
// @param request - ListSubtaskItemsRequest
//
// @return ListSubtaskItemsResponse
func (client *Client) ListSubtaskItems(TenantId *string, TaskID *string, SubtaskId *string, request *ListSubtaskItemsRequest) (_result *ListSubtaskItemsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListSubtaskItemsResponse{}
	_body, _err := client.ListSubtaskItemsWithOptions(TenantId, TaskID, SubtaskId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取子任务列表页信息
//
// @param request - ListSubtasksRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSubtasksResponse
func (client *Client) ListSubtasksWithOptions(TenantId *string, TaskID *string, request *ListSubtasksRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListSubtasksResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListSubtasks"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/api/v1/tenants/" + tea.StringValue(openapiutil.GetEncodeParam(TenantId)) + "/tasks/" + tea.StringValue(openapiutil.GetEncodeParam(TaskID)) + "/subtasks"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListSubtasksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取子任务列表页信息
//
// @param request - ListSubtasksRequest
//
// @return ListSubtasksResponse
func (client *Client) ListSubtasks(TenantId *string, TaskID *string, request *ListSubtasksRequest) (_result *ListSubtasksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListSubtasksResponse{}
	_body, _err := client.ListSubtasksWithOptions(TenantId, TaskID, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取任务列表页信息
//
// @param request - ListTasksRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTasksResponse
func (client *Client) ListTasksWithOptions(TenantId *string, request *ListTasksRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListTasksResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTasks"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/api/v1/tenants/" + tea.StringValue(openapiutil.GetEncodeParam(TenantId)) + "/tasks"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTasksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取任务列表页信息
//
// @param request - ListTasksRequest
//
// @return ListTasksResponse
func (client *Client) ListTasks(TenantId *string, request *ListTasksRequest) (_result *ListTasksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListTasksResponse{}
	_body, _err := client.ListTasksWithOptions(TenantId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取租户模板信息列表
//
// @param tmpReq - ListTemplatesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTemplatesResponse
func (client *Client) ListTemplatesWithOptions(TenantId *string, tmpReq *ListTemplatesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListTemplatesResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListTemplatesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Types)) {
		request.TypesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Types, tea.String("Types"), tea.String("simple"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SearchKey)) {
		query["SearchKey"] = request.SearchKey
	}

	if !tea.BoolValue(util.IsUnset(request.TypesShrink)) {
		query["Types"] = request.TypesShrink
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTemplates"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/api/v1/tenants/" + tea.StringValue(openapiutil.GetEncodeParam(TenantId)) + "/templates"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTemplatesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取租户模板信息列表
//
// @param request - ListTemplatesRequest
//
// @return ListTemplatesResponse
func (client *Client) ListTemplates(TenantId *string, request *ListTemplatesRequest) (_result *ListTemplatesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListTemplatesResponse{}
	_body, _err := client.ListTemplatesWithOptions(TenantId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取租户列表
//
// @param request - ListTenantsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTenantsResponse
func (client *Client) ListTenantsWithOptions(request *ListTenantsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListTenantsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTenants"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/api/v1/tenants"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTenantsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取租户列表
//
// @param request - ListTenantsRequest
//
// @return ListTenantsResponse
func (client *Client) ListTenants(request *ListTenantsRequest) (_result *ListTenantsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListTenantsResponse{}
	_body, _err := client.ListTenantsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取用户列表
//
// @param request - ListUsersRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUsersResponse
func (client *Client) ListUsersWithOptions(TenantId *string, request *ListUsersRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListUsersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListUsers"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/api/v1/tenants/" + tea.StringValue(openapiutil.GetEncodeParam(TenantId)) + "/users"),
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

// Summary:
//
// 获取用户列表
//
// @param request - ListUsersRequest
//
// @return ListUsersResponse
func (client *Client) ListUsers(TenantId *string, request *ListUsersRequest) (_result *ListUsersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListUsersResponse{}
	_body, _err := client.ListUsersWithOptions(TenantId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除结点人力
//
// @param request - RemoveWorkNodeWorkforceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveWorkNodeWorkforceResponse
func (client *Client) RemoveWorkNodeWorkforceWithOptions(TenantId *string, TaskId *string, WorkNodeId *string, request *RemoveWorkNodeWorkforceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *RemoveWorkNodeWorkforceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UserIds)) {
		body["UserIds"] = request.UserIds
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RemoveWorkNodeWorkforce"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/api/v1/tenants/" + tea.StringValue(openapiutil.GetEncodeParam(TenantId)) + "/tasks/" + tea.StringValue(openapiutil.GetEncodeParam(TaskId)) + "/worknodes/" + tea.StringValue(openapiutil.GetEncodeParam(WorkNodeId)) + "/workforce"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &RemoveWorkNodeWorkforceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除结点人力
//
// @param request - RemoveWorkNodeWorkforceRequest
//
// @return RemoveWorkNodeWorkforceResponse
func (client *Client) RemoveWorkNodeWorkforce(TenantId *string, TaskId *string, WorkNodeId *string, request *RemoveWorkNodeWorkforceRequest) (_result *RemoveWorkNodeWorkforceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RemoveWorkNodeWorkforceResponse{}
	_body, _err := client.RemoveWorkNodeWorkforceWithOptions(TenantId, TaskId, WorkNodeId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新标注任务基础信息
//
// @param request - UpdateTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateTaskResponse
func (client *Client) UpdateTaskWithOptions(TenantId *string, TaskId *string, request *UpdateTaskRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateTask"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/api/v1/tenants/" + tea.StringValue(openapiutil.GetEncodeParam(TenantId)) + "/tasks/" + tea.StringValue(openapiutil.GetEncodeParam(TaskId))),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新标注任务基础信息
//
// @param request - UpdateTaskRequest
//
// @return UpdateTaskResponse
func (client *Client) UpdateTask(TenantId *string, TaskId *string, request *UpdateTaskRequest) (_result *UpdateTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateTaskResponse{}
	_body, _err := client.UpdateTaskWithOptions(TenantId, TaskId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新任务人力
//
// @param request - UpdateTaskWorkforceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateTaskWorkforceResponse
func (client *Client) UpdateTaskWorkforceWithOptions(TenantId *string, TaskId *string, request *UpdateTaskWorkforceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateTaskWorkforceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Workforce)) {
		body["Workforce"] = request.Workforce
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateTaskWorkforce"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/api/v1/tenants/" + tea.StringValue(openapiutil.GetEncodeParam(TenantId)) + "/tasks/" + tea.StringValue(openapiutil.GetEncodeParam(TaskId)) + "/workforce"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateTaskWorkforceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新任务人力
//
// @param request - UpdateTaskWorkforceRequest
//
// @return UpdateTaskWorkforceResponse
func (client *Client) UpdateTaskWorkforce(TenantId *string, TaskId *string, request *UpdateTaskWorkforceRequest) (_result *UpdateTaskWorkforceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateTaskWorkforceResponse{}
	_body, _err := client.UpdateTaskWorkforceWithOptions(TenantId, TaskId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新标注模版
//
// @param request - UpdateTemplateRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateTemplateResponse
func (client *Client) UpdateTemplateWithOptions(TenantId *string, TemplateId *string, request *UpdateTemplateRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateTemplate"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/api/v1/tenants/" + tea.StringValue(openapiutil.GetEncodeParam(TenantId)) + "/templates/" + tea.StringValue(openapiutil.GetEncodeParam(TemplateId))),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新标注模版
//
// @param request - UpdateTemplateRequest
//
// @return UpdateTemplateResponse
func (client *Client) UpdateTemplate(TenantId *string, TemplateId *string, request *UpdateTemplateRequest) (_result *UpdateTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateTemplateResponse{}
	_body, _err := client.UpdateTemplateWithOptions(TenantId, TemplateId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新租户信息
//
// @param request - UpdateTenantRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateTenantResponse
func (client *Client) UpdateTenantWithOptions(TenantId *string, request *UpdateTenantRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateTenantResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.TenantName)) {
		body["TenantName"] = request.TenantName
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateTenant"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/api/v1/tenants/" + tea.StringValue(openapiutil.GetEncodeParam(TenantId))),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateTenantResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新租户信息
//
// @param request - UpdateTenantRequest
//
// @return UpdateTenantResponse
func (client *Client) UpdateTenant(TenantId *string, request *UpdateTenantRequest) (_result *UpdateTenantResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateTenantResponse{}
	_body, _err := client.UpdateTenantWithOptions(TenantId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新用户信息
//
// @param request - UpdateUserRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateUserResponse
func (client *Client) UpdateUserWithOptions(TenantId *string, UserId *string, request *UpdateUserRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Role)) {
		body["Role"] = request.Role
	}

	if !tea.BoolValue(util.IsUnset(request.UserName)) {
		body["UserName"] = request.UserName
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateUser"),
		Version:     tea.String("2022-06-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/api/v1/tenants/" + tea.StringValue(openapiutil.GetEncodeParam(TenantId)) + "/users/" + tea.StringValue(openapiutil.GetEncodeParam(UserId))),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新用户信息
//
// @param request - UpdateUserRequest
//
// @return UpdateUserResponse
func (client *Client) UpdateUser(TenantId *string, UserId *string, request *UpdateUserRequest) (_result *UpdateUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateUserResponse{}
	_body, _err := client.UpdateUserWithOptions(TenantId, UserId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
