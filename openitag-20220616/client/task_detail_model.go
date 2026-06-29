// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTaskDetail interface {
	dara.Model
	String() string
	GoString() string
	SetAdmins(v []*SimpleUser) *TaskDetail
	GetAdmins() []*SimpleUser
	SetAlertTime(v int64) *TaskDetail
	GetAlertTime() *int64
	SetAllowAppendData(v bool) *TaskDetail
	GetAllowAppendData() *bool
	SetArchived(v bool) *TaskDetail
	GetArchived() *bool
	SetArchivedInfos(v string) *TaskDetail
	GetArchivedInfos() *string
	SetAssignConfig(v map[string]interface{}) *TaskDetail
	GetAssignConfig() map[string]interface{}
	SetCreator(v *SimpleUser) *TaskDetail
	GetCreator() *SimpleUser
	SetDatasetProxyRelations(v []*TaskDetailDatasetProxyRelations) *TaskDetail
	GetDatasetProxyRelations() []*TaskDetailDatasetProxyRelations
	SetExif(v map[string]interface{}) *TaskDetail
	GetExif() map[string]interface{}
	SetGmtCreateTime(v string) *TaskDetail
	GetGmtCreateTime() *string
	SetGmtModifiedTime(v string) *TaskDetail
	GetGmtModifiedTime() *string
	SetLabelStyle(v string) *TaskDetail
	GetLabelStyle() *string
	SetMineConfigs(v map[string]interface{}) *TaskDetail
	GetMineConfigs() map[string]interface{}
	SetModifier(v *SimpleUser) *TaskDetail
	GetModifier() *SimpleUser
	SetNoticeConfig(v map[string]interface{}) *TaskDetail
	GetNoticeConfig() map[string]interface{}
	SetPeriodConfig(v map[string]interface{}) *TaskDetail
	GetPeriodConfig() map[string]interface{}
	SetRefTaskId(v string) *TaskDetail
	GetRefTaskId() *string
	SetRelateTaskConfig(v map[string]interface{}) *TaskDetail
	GetRelateTaskConfig() map[string]interface{}
	SetRemark(v string) *TaskDetail
	GetRemark() *string
	SetResultCallbackConfig(v map[string]interface{}) *TaskDetail
	GetResultCallbackConfig() map[string]interface{}
	SetStage(v string) *TaskDetail
	GetStage() *string
	SetStatus(v string) *TaskDetail
	GetStatus() *string
	SetTags(v []*string) *TaskDetail
	GetTags() []*string
	SetTaskId(v string) *TaskDetail
	GetTaskId() *string
	SetTaskName(v string) *TaskDetail
	GetTaskName() *string
	SetTaskTemplateConfig(v *TaskDetailTaskTemplateConfig) *TaskDetail
	GetTaskTemplateConfig() *TaskDetailTaskTemplateConfig
	SetTaskType(v string) *TaskDetail
	GetTaskType() *string
	SetTaskWorkflow(v []*TaskDetailTaskWorkflow) *TaskDetail
	GetTaskWorkflow() []*TaskDetailTaskWorkflow
	SetTemplateId(v string) *TaskDetail
	GetTemplateId() *string
	SetTenantId(v string) *TaskDetail
	GetTenantId() *string
	SetTenantName(v string) *TaskDetail
	GetTenantName() *string
	SetUUID(v string) *TaskDetail
	GetUUID() *string
	SetVoteConfigs(v map[string]interface{}) *TaskDetail
	GetVoteConfigs() map[string]interface{}
	SetWorkflowNodes(v []*string) *TaskDetail
	GetWorkflowNodes() []*string
	SetRunMsg(v string) *TaskDetail
	GetRunMsg() *string
}

type TaskDetail struct {
	// List of job administrators.
	Admins []*SimpleUser `json:"Admins,omitempty" xml:"Admins,omitempty" type:"Repeated"`
	// Alert time.
	//
	// example:
	//
	// 12412312
	AlertTime *int64 `json:"AlertTime,omitempty" xml:"AlertTime,omitempty"`
	// Indicates whether data appending is allowed.
	//
	// example:
	//
	// true
	AllowAppendData *bool `json:"AllowAppendData,omitempty" xml:"AllowAppendData,omitempty"`
	// Indicates whether the job has been archived.
	//
	// example:
	//
	// false
	Archived *bool `json:"Archived,omitempty" xml:"Archived,omitempty"`
	// Data archiving information.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// null
	ArchivedInfos *string `json:"ArchivedInfos,omitempty" xml:"ArchivedInfos,omitempty"`
	// Job assignment configuration.
	//
	// example:
	//
	// null
	AssignConfig map[string]interface{} `json:"AssignConfig,omitempty" xml:"AssignConfig,omitempty"`
	// Creator information.
	Creator *SimpleUser `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// List of Data proxy relationships.
	DatasetProxyRelations []*TaskDetailDatasetProxyRelations `json:"DatasetProxyRelations,omitempty" xml:"DatasetProxyRelations,omitempty" type:"Repeated"`
	// Additional configuration.
	//
	// example:
	//
	// null
	Exif map[string]interface{} `json:"Exif,omitempty" xml:"Exif,omitempty"`
	// Creation UTC time.
	//
	// example:
	//
	// 2022-07-04 11:42:57
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// UTC time of the last modification.
	//
	// example:
	//
	// 2022-08-16 18:38:42
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// Annotation style.
	//
	// example:
	//
	// null
	LabelStyle *string `json:"LabelStyle,omitempty" xml:"LabelStyle,omitempty"`
	// My configuration.
	//
	// example:
	//
	// null
	MineConfigs map[string]interface{} `json:"MineConfigs,omitempty" xml:"MineConfigs,omitempty"`
	// Updated By information.
	Modifier *SimpleUser `json:"Modifier,omitempty" xml:"Modifier,omitempty"`
	// Notice configuration.
	//
	// example:
	//
	// {"DingHook":"","ListenActions":[],"SubTaskAlertGap":""}
	NoticeConfig map[string]interface{} `json:"NoticeConfig,omitempty" xml:"NoticeConfig,omitempty"`
	// Time configuration.
	//
	// example:
	//
	// {"periodConfigId":"","status":""}
	PeriodConfig map[string]interface{} `json:"PeriodConfig,omitempty" xml:"PeriodConfig,omitempty"`
	// Auto triggered task ID.
	//
	// example:
	//
	// 12312414
	RefTaskId *string `json:"RefTaskId,omitempty" xml:"RefTaskId,omitempty"`
	// Related job configuration.
	//
	// example:
	//
	// null
	RelateTaskConfig map[string]interface{} `json:"RelateTaskConfig,omitempty" xml:"RelateTaskConfig,omitempty"`
	// Remark information.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// demo
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// Result callback configuration.
	//
	// example:
	//
	// {"RetMsgMode":"","Exif":{}}
	ResultCallbackConfig map[string]interface{} `json:"ResultCallbackConfig,omitempty" xml:"ResultCallbackConfig,omitempty"`
	// Current edge zone. Possible values:
	//
	// - MARK: Annotating.
	//
	// - CHECK: Checking.
	//
	// - SAMPLING: Validating.
	//
	// example:
	//
	// MARK
	Stage *string `json:"Stage,omitempty" xml:"Stage,omitempty"`
	// Task Status. Possible values:
	//
	// - CREATED
	//
	// - SUCCESS
	//
	// example:
	//
	// CREATED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// List of job tags.
	//
	// if can be null:
	// true
	Tags []*string `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// Job ID.
	//
	// example:
	//
	// 15438***8306500608
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// Task Name.
	//
	// example:
	//
	// 测试任务
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// Supplementary configuration for the Job template.
	TaskTemplateConfig *TaskDetailTaskTemplateConfig `json:"TaskTemplateConfig,omitempty" xml:"TaskTemplateConfig,omitempty" type:"Struct"`
	// Task Type.
	//
	// example:
	//
	// NORMAL
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	// List of job stream configurations.
	TaskWorkflow []*TaskDetailTaskWorkflow `json:"TaskWorkflow,omitempty" xml:"TaskWorkflow,omitempty" type:"Repeated"`
	// Template ID.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// 1529***348342353920
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// Tenant ID.
	//
	// example:
	//
	// GA***W134
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// Tenant name.
	//
	// example:
	//
	// 测试任务202208101424
	TenantName *string `json:"TenantName,omitempty" xml:"TenantName,omitempty"`
	// UUID.
	//
	// example:
	//
	// paiworkspace-0001
	UUID *string `json:"UUID,omitempty" xml:"UUID,omitempty"`
	// Voting configuration.
	//
	// example:
	//
	// {"1":{"VoteNum":1,"MinVote":1}}
	VoteConfigs map[string]interface{} `json:"VoteConfigs,omitempty" xml:"VoteConfigs,omitempty"`
	// List of pipeline nodes.
	WorkflowNodes []*string `json:"WorkflowNodes,omitempty" xml:"WorkflowNodes,omitempty" type:"Repeated"`
	// Run message.
	//
	// example:
	//
	// xxxxx
	RunMsg *string `json:"runMsg,omitempty" xml:"runMsg,omitempty"`
}

func (s TaskDetail) String() string {
	return dara.Prettify(s)
}

func (s TaskDetail) GoString() string {
	return s.String()
}

func (s *TaskDetail) GetAdmins() []*SimpleUser {
	return s.Admins
}

func (s *TaskDetail) GetAlertTime() *int64 {
	return s.AlertTime
}

func (s *TaskDetail) GetAllowAppendData() *bool {
	return s.AllowAppendData
}

func (s *TaskDetail) GetArchived() *bool {
	return s.Archived
}

func (s *TaskDetail) GetArchivedInfos() *string {
	return s.ArchivedInfos
}

func (s *TaskDetail) GetAssignConfig() map[string]interface{} {
	return s.AssignConfig
}

func (s *TaskDetail) GetCreator() *SimpleUser {
	return s.Creator
}

func (s *TaskDetail) GetDatasetProxyRelations() []*TaskDetailDatasetProxyRelations {
	return s.DatasetProxyRelations
}

func (s *TaskDetail) GetExif() map[string]interface{} {
	return s.Exif
}

func (s *TaskDetail) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *TaskDetail) GetGmtModifiedTime() *string {
	return s.GmtModifiedTime
}

func (s *TaskDetail) GetLabelStyle() *string {
	return s.LabelStyle
}

func (s *TaskDetail) GetMineConfigs() map[string]interface{} {
	return s.MineConfigs
}

func (s *TaskDetail) GetModifier() *SimpleUser {
	return s.Modifier
}

func (s *TaskDetail) GetNoticeConfig() map[string]interface{} {
	return s.NoticeConfig
}

func (s *TaskDetail) GetPeriodConfig() map[string]interface{} {
	return s.PeriodConfig
}

func (s *TaskDetail) GetRefTaskId() *string {
	return s.RefTaskId
}

func (s *TaskDetail) GetRelateTaskConfig() map[string]interface{} {
	return s.RelateTaskConfig
}

func (s *TaskDetail) GetRemark() *string {
	return s.Remark
}

func (s *TaskDetail) GetResultCallbackConfig() map[string]interface{} {
	return s.ResultCallbackConfig
}

func (s *TaskDetail) GetStage() *string {
	return s.Stage
}

func (s *TaskDetail) GetStatus() *string {
	return s.Status
}

func (s *TaskDetail) GetTags() []*string {
	return s.Tags
}

func (s *TaskDetail) GetTaskId() *string {
	return s.TaskId
}

func (s *TaskDetail) GetTaskName() *string {
	return s.TaskName
}

func (s *TaskDetail) GetTaskTemplateConfig() *TaskDetailTaskTemplateConfig {
	return s.TaskTemplateConfig
}

func (s *TaskDetail) GetTaskType() *string {
	return s.TaskType
}

func (s *TaskDetail) GetTaskWorkflow() []*TaskDetailTaskWorkflow {
	return s.TaskWorkflow
}

func (s *TaskDetail) GetTemplateId() *string {
	return s.TemplateId
}

func (s *TaskDetail) GetTenantId() *string {
	return s.TenantId
}

func (s *TaskDetail) GetTenantName() *string {
	return s.TenantName
}

func (s *TaskDetail) GetUUID() *string {
	return s.UUID
}

func (s *TaskDetail) GetVoteConfigs() map[string]interface{} {
	return s.VoteConfigs
}

func (s *TaskDetail) GetWorkflowNodes() []*string {
	return s.WorkflowNodes
}

func (s *TaskDetail) GetRunMsg() *string {
	return s.RunMsg
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

func (s *TaskDetail) Validate() error {
	if s.Admins != nil {
		for _, item := range s.Admins {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Creator != nil {
		if err := s.Creator.Validate(); err != nil {
			return err
		}
	}
	if s.DatasetProxyRelations != nil {
		for _, item := range s.DatasetProxyRelations {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Modifier != nil {
		if err := s.Modifier.Validate(); err != nil {
			return err
		}
	}
	if s.TaskTemplateConfig != nil {
		if err := s.TaskTemplateConfig.Validate(); err != nil {
			return err
		}
	}
	if s.TaskWorkflow != nil {
		for _, item := range s.TaskWorkflow {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type TaskDetailDatasetProxyRelations struct {
	// Dataset ID.
	//
	// example:
	//
	// 102***124
	DatasetId *string `json:"DatasetId,omitempty" xml:"DatasetId,omitempty"`
	// Dataset type.
	//
	// example:
	//
	// LABEL
	DatasetType *string `json:"DatasetType,omitempty" xml:"DatasetType,omitempty"`
	// Additional information.
	//
	// example:
	//
	// false
	Exif map[string]interface{} `json:"Exif,omitempty" xml:"Exif,omitempty"`
	// Dataset source.
	//
	// example:
	//
	// PAI
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// Source business ID.
	//
	// example:
	//
	// d-24eyimdzdn4a98jktk
	SourceBizId *string `json:"SourceBizId,omitempty" xml:"SourceBizId,omitempty"`
	// Source dataset ID.
	//
	// example:
	//
	// 2312124
	SourceDatasetId *string `json:"SourceDatasetId,omitempty" xml:"SourceDatasetId,omitempty"`
}

func (s TaskDetailDatasetProxyRelations) String() string {
	return dara.Prettify(s)
}

func (s TaskDetailDatasetProxyRelations) GoString() string {
	return s.String()
}

func (s *TaskDetailDatasetProxyRelations) GetDatasetId() *string {
	return s.DatasetId
}

func (s *TaskDetailDatasetProxyRelations) GetDatasetType() *string {
	return s.DatasetType
}

func (s *TaskDetailDatasetProxyRelations) GetExif() map[string]interface{} {
	return s.Exif
}

func (s *TaskDetailDatasetProxyRelations) GetSource() *string {
	return s.Source
}

func (s *TaskDetailDatasetProxyRelations) GetSourceBizId() *string {
	return s.SourceBizId
}

func (s *TaskDetailDatasetProxyRelations) GetSourceDatasetId() *string {
	return s.SourceDatasetId
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

func (s *TaskDetailDatasetProxyRelations) Validate() error {
	return dara.Validate(s)
}

type TaskDetailTaskTemplateConfig struct {
	// Additional information for template configuration.
	//
	// example:
	//
	// false
	Exif map[string]interface{} `json:"Exif,omitempty" xml:"Exif,omitempty"`
	// Resource key.
	//
	// example:
	//
	// picture1
	ResourceKey *string `json:"ResourceKey,omitempty" xml:"ResourceKey,omitempty"`
	// Robot configuration.
	//
	// example:
	//
	// null
	RobotConfig map[string]interface{} `json:"RobotConfig,omitempty" xml:"RobotConfig,omitempty"`
	// If the number of questions in the Job is less than that in the template, you can select the required questions from the List.
	SelectQuestions []*string `json:"SelectQuestions,omitempty" xml:"SelectQuestions,omitempty" type:"Repeated"`
	// Options configuration.
	//
	// example:
	//
	// {"1":[{"label":"label1","key":"label1"}]}
	TemplateOptionMap map[string]interface{} `json:"TemplateOptionMap,omitempty" xml:"TemplateOptionMap,omitempty"`
	// Template relation ID.
	//
	// example:
	//
	// 200
	TemplateRelationId *string `json:"TemplateRelationId,omitempty" xml:"TemplateRelationId,omitempty"`
}

func (s TaskDetailTaskTemplateConfig) String() string {
	return dara.Prettify(s)
}

func (s TaskDetailTaskTemplateConfig) GoString() string {
	return s.String()
}

func (s *TaskDetailTaskTemplateConfig) GetExif() map[string]interface{} {
	return s.Exif
}

func (s *TaskDetailTaskTemplateConfig) GetResourceKey() *string {
	return s.ResourceKey
}

func (s *TaskDetailTaskTemplateConfig) GetRobotConfig() map[string]interface{} {
	return s.RobotConfig
}

func (s *TaskDetailTaskTemplateConfig) GetSelectQuestions() []*string {
	return s.SelectQuestions
}

func (s *TaskDetailTaskTemplateConfig) GetTemplateOptionMap() map[string]interface{} {
	return s.TemplateOptionMap
}

func (s *TaskDetailTaskTemplateConfig) GetTemplateRelationId() *string {
	return s.TemplateRelationId
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

func (s *TaskDetailTaskTemplateConfig) Validate() error {
	return dara.Validate(s)
}

type TaskDetailTaskWorkflow struct {
	// Extra information.
	//
	// example:
	//
	// false
	Exif map[string]interface{} `json:"Exif,omitempty" xml:"Exif,omitempty"`
	// Group list.
	Groups []*string `json:"Groups,omitempty" xml:"Groups,omitempty" type:"Repeated"`
	// Edge zone name. Possible values:
	//
	// - MARK
	//
	// - CHECK
	//
	// - SAMPLING
	//
	// example:
	//
	// MARK
	NodeName *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	// User List.
	Users []*SimpleUser `json:"Users,omitempty" xml:"Users,omitempty" type:"Repeated"`
}

func (s TaskDetailTaskWorkflow) String() string {
	return dara.Prettify(s)
}

func (s TaskDetailTaskWorkflow) GoString() string {
	return s.String()
}

func (s *TaskDetailTaskWorkflow) GetExif() map[string]interface{} {
	return s.Exif
}

func (s *TaskDetailTaskWorkflow) GetGroups() []*string {
	return s.Groups
}

func (s *TaskDetailTaskWorkflow) GetNodeName() *string {
	return s.NodeName
}

func (s *TaskDetailTaskWorkflow) GetUsers() []*SimpleUser {
	return s.Users
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

func (s *TaskDetailTaskWorkflow) Validate() error {
	if s.Users != nil {
		for _, item := range s.Users {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
