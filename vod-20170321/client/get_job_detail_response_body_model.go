// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetJobDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAIJobDetail(v *GetJobDetailResponseBodyAIJobDetail) *GetJobDetailResponseBody
	GetAIJobDetail() *GetJobDetailResponseBodyAIJobDetail
	SetJobType(v string) *GetJobDetailResponseBody
	GetJobType() *string
	SetRequestId(v string) *GetJobDetailResponseBody
	GetRequestId() *string
	SetSnapshotJobDetail(v *GetJobDetailResponseBodySnapshotJobDetail) *GetJobDetailResponseBody
	GetSnapshotJobDetail() *GetJobDetailResponseBodySnapshotJobDetail
	SetTranscodeJobDetail(v *GetJobDetailResponseBodyTranscodeJobDetail) *GetJobDetailResponseBody
	GetTranscodeJobDetail() *GetJobDetailResponseBodyTranscodeJobDetail
	SetWorkflowTaskDetail(v *GetJobDetailResponseBodyWorkflowTaskDetail) *GetJobDetailResponseBody
	GetWorkflowTaskDetail() *GetJobDetailResponseBodyWorkflowTaskDetail
}

type GetJobDetailResponseBody struct {
	// The details of the AI task. This parameter takes effect only when the TaskType parameter is set to AI.
	AIJobDetail *GetJobDetailResponseBodyAIJobDetail `json:"AIJobDetail,omitempty" xml:"AIJobDetail,omitempty" type:"Struct"`
	// The type of the task. Valid values:
	//
	// example:
	//
	// transcode
	JobType *string `json:"JobType,omitempty" xml:"JobType,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 6708D849-F109-1A6C-AC91-************
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The details of the snapshot task. This parameter takes effect only when the jobType parameter is set to Snapshot.
	SnapshotJobDetail *GetJobDetailResponseBodySnapshotJobDetail `json:"SnapshotJobDetail,omitempty" xml:"SnapshotJobDetail,omitempty" type:"Struct"`
	// The details of the transcoding task. This parameter takes effect only when the jobType parameter is set to Transcode.
	TranscodeJobDetail *GetJobDetailResponseBodyTranscodeJobDetail `json:"TranscodeJobDetail,omitempty" xml:"TranscodeJobDetail,omitempty" type:"Struct"`
	WorkflowTaskDetail *GetJobDetailResponseBodyWorkflowTaskDetail `json:"WorkflowTaskDetail,omitempty" xml:"WorkflowTaskDetail,omitempty" type:"Struct"`
}

func (s GetJobDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetJobDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetJobDetailResponseBody) GetAIJobDetail() *GetJobDetailResponseBodyAIJobDetail {
	return s.AIJobDetail
}

func (s *GetJobDetailResponseBody) GetJobType() *string {
	return s.JobType
}

func (s *GetJobDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetJobDetailResponseBody) GetSnapshotJobDetail() *GetJobDetailResponseBodySnapshotJobDetail {
	return s.SnapshotJobDetail
}

func (s *GetJobDetailResponseBody) GetTranscodeJobDetail() *GetJobDetailResponseBodyTranscodeJobDetail {
	return s.TranscodeJobDetail
}

func (s *GetJobDetailResponseBody) GetWorkflowTaskDetail() *GetJobDetailResponseBodyWorkflowTaskDetail {
	return s.WorkflowTaskDetail
}

func (s *GetJobDetailResponseBody) SetAIJobDetail(v *GetJobDetailResponseBodyAIJobDetail) *GetJobDetailResponseBody {
	s.AIJobDetail = v
	return s
}

func (s *GetJobDetailResponseBody) SetJobType(v string) *GetJobDetailResponseBody {
	s.JobType = &v
	return s
}

func (s *GetJobDetailResponseBody) SetRequestId(v string) *GetJobDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetJobDetailResponseBody) SetSnapshotJobDetail(v *GetJobDetailResponseBodySnapshotJobDetail) *GetJobDetailResponseBody {
	s.SnapshotJobDetail = v
	return s
}

func (s *GetJobDetailResponseBody) SetTranscodeJobDetail(v *GetJobDetailResponseBodyTranscodeJobDetail) *GetJobDetailResponseBody {
	s.TranscodeJobDetail = v
	return s
}

func (s *GetJobDetailResponseBody) SetWorkflowTaskDetail(v *GetJobDetailResponseBodyWorkflowTaskDetail) *GetJobDetailResponseBody {
	s.WorkflowTaskDetail = v
	return s
}

func (s *GetJobDetailResponseBody) Validate() error {
	if s.AIJobDetail != nil {
		if err := s.AIJobDetail.Validate(); err != nil {
			return err
		}
	}
	if s.SnapshotJobDetail != nil {
		if err := s.SnapshotJobDetail.Validate(); err != nil {
			return err
		}
	}
	if s.TranscodeJobDetail != nil {
		if err := s.TranscodeJobDetail.Validate(); err != nil {
			return err
		}
	}
	if s.WorkflowTaskDetail != nil {
		if err := s.WorkflowTaskDetail.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetJobDetailResponseBodyAIJobDetail struct {
	// The end time of the task.
	//
	// example:
	//
	// 2024-10-14T07:39:46Z
	CompleteTime *string `json:"CompleteTime,omitempty" xml:"CompleteTime,omitempty"`
	// The time when the task was created. The time follows the ISO 8601 standard in the YYYY-MM-DDTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2024-10-14T07:39:25Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the task.
	//
	// example:
	//
	// 5c9dff751ba**********59d50a967f5
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The type of the AI task.
	//
	// example:
	//
	// AIVideoCensor
	JobType *string `json:"JobType,omitempty" xml:"JobType,omitempty"`
	// The ID of the media asset.
	//
	// example:
	//
	// 30e5d7**********bd900764de7c0102
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	// The status of the task. Valid values:
	//
	// 	- reserved
	//
	// 	- init
	//
	// 	- success
	//
	// 	- fail
	//
	// 	- processing
	//
	// 	- analysing
	//
	// example:
	//
	// success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The template configuration.
	//
	// example:
	//
	// {"AuditRange":["video","image-cover","text-title"],"AuditContent":["screen"],"AuditItem":["terrorism","porn"],"AuditAutoBlock":"no"}
	TemplateConfig *string `json:"TemplateConfig,omitempty" xml:"TemplateConfig,omitempty"`
	// The trigger mode. Valid values:
	//
	// 	- Auto
	//
	// 	- Manual
	//
	// example:
	//
	// Auto
	Trigger *string `json:"Trigger,omitempty" xml:"Trigger,omitempty"`
	// The ID of the user who submitted the task.
	//
	// example:
	//
	// 139109*****84930
	UserId *int64 `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s GetJobDetailResponseBodyAIJobDetail) String() string {
	return dara.Prettify(s)
}

func (s GetJobDetailResponseBodyAIJobDetail) GoString() string {
	return s.String()
}

func (s *GetJobDetailResponseBodyAIJobDetail) GetCompleteTime() *string {
	return s.CompleteTime
}

func (s *GetJobDetailResponseBodyAIJobDetail) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetJobDetailResponseBodyAIJobDetail) GetJobId() *string {
	return s.JobId
}

func (s *GetJobDetailResponseBodyAIJobDetail) GetJobType() *string {
	return s.JobType
}

func (s *GetJobDetailResponseBodyAIJobDetail) GetMediaId() *string {
	return s.MediaId
}

func (s *GetJobDetailResponseBodyAIJobDetail) GetStatus() *string {
	return s.Status
}

func (s *GetJobDetailResponseBodyAIJobDetail) GetTemplateConfig() *string {
	return s.TemplateConfig
}

func (s *GetJobDetailResponseBodyAIJobDetail) GetTrigger() *string {
	return s.Trigger
}

func (s *GetJobDetailResponseBodyAIJobDetail) GetUserId() *int64 {
	return s.UserId
}

func (s *GetJobDetailResponseBodyAIJobDetail) SetCompleteTime(v string) *GetJobDetailResponseBodyAIJobDetail {
	s.CompleteTime = &v
	return s
}

func (s *GetJobDetailResponseBodyAIJobDetail) SetCreateTime(v string) *GetJobDetailResponseBodyAIJobDetail {
	s.CreateTime = &v
	return s
}

func (s *GetJobDetailResponseBodyAIJobDetail) SetJobId(v string) *GetJobDetailResponseBodyAIJobDetail {
	s.JobId = &v
	return s
}

func (s *GetJobDetailResponseBodyAIJobDetail) SetJobType(v string) *GetJobDetailResponseBodyAIJobDetail {
	s.JobType = &v
	return s
}

func (s *GetJobDetailResponseBodyAIJobDetail) SetMediaId(v string) *GetJobDetailResponseBodyAIJobDetail {
	s.MediaId = &v
	return s
}

func (s *GetJobDetailResponseBodyAIJobDetail) SetStatus(v string) *GetJobDetailResponseBodyAIJobDetail {
	s.Status = &v
	return s
}

func (s *GetJobDetailResponseBodyAIJobDetail) SetTemplateConfig(v string) *GetJobDetailResponseBodyAIJobDetail {
	s.TemplateConfig = &v
	return s
}

func (s *GetJobDetailResponseBodyAIJobDetail) SetTrigger(v string) *GetJobDetailResponseBodyAIJobDetail {
	s.Trigger = &v
	return s
}

func (s *GetJobDetailResponseBodyAIJobDetail) SetUserId(v int64) *GetJobDetailResponseBodyAIJobDetail {
	s.UserId = &v
	return s
}

func (s *GetJobDetailResponseBodyAIJobDetail) Validate() error {
	return dara.Validate(s)
}

type GetJobDetailResponseBodySnapshotJobDetail struct {
	// The time when the task was complete.
	//
	// example:
	//
	// 2024-10-14T07:39:45Z
	CompleteTime *string `json:"CompleteTime,omitempty" xml:"CompleteTime,omitempty"`
	// The time when the task was created. The time follows the ISO 8601 standard in the YYYY-MM-DDTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2024-10-14T07:39:25Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the task.
	//
	// example:
	//
	// 63df12s0**********4hdq249t82kr91
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// Configuration of normal snapshots.
	//
	// example:
	//
	// {"inl":0,"num":32,"tm":5,"wd":"352","ft":"normal","hg":"640"}
	NormalConfig *string `json:"NormalConfig,omitempty" xml:"NormalConfig,omitempty"`
	// The sprite configuration.
	//
	// example:
	//
	// {"pad":"0","lines":"10","mgin":"0","cols":"10","ikcp":"false","hg":"68"}
	SpriteConfig *string `json:"SpriteConfig,omitempty" xml:"SpriteConfig,omitempty"`
	// The status of the task. Valid values:
	//
	// 	- Processing
	//
	// 	- Fail
	//
	// 	- Success
	//
	// example:
	//
	// Success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The trigger mode. Valid values:
	//
	// 	- Auto
	//
	// 	- Manual
	//
	// example:
	//
	// Auto
	Trigger *string `json:"Trigger,omitempty" xml:"Trigger,omitempty"`
	// The ID of the user who submitted the task.
	//
	// example:
	//
	// 139109*****84930
	UserId *int64 `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The ID of the media asset.
	//
	// example:
	//
	// 30e5d7**********bd900764de7c0102
	VideoId *string `json:"VideoId,omitempty" xml:"VideoId,omitempty"`
}

func (s GetJobDetailResponseBodySnapshotJobDetail) String() string {
	return dara.Prettify(s)
}

func (s GetJobDetailResponseBodySnapshotJobDetail) GoString() string {
	return s.String()
}

func (s *GetJobDetailResponseBodySnapshotJobDetail) GetCompleteTime() *string {
	return s.CompleteTime
}

func (s *GetJobDetailResponseBodySnapshotJobDetail) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetJobDetailResponseBodySnapshotJobDetail) GetJobId() *string {
	return s.JobId
}

func (s *GetJobDetailResponseBodySnapshotJobDetail) GetNormalConfig() *string {
	return s.NormalConfig
}

func (s *GetJobDetailResponseBodySnapshotJobDetail) GetSpriteConfig() *string {
	return s.SpriteConfig
}

func (s *GetJobDetailResponseBodySnapshotJobDetail) GetStatus() *string {
	return s.Status
}

func (s *GetJobDetailResponseBodySnapshotJobDetail) GetTrigger() *string {
	return s.Trigger
}

func (s *GetJobDetailResponseBodySnapshotJobDetail) GetUserId() *int64 {
	return s.UserId
}

func (s *GetJobDetailResponseBodySnapshotJobDetail) GetVideoId() *string {
	return s.VideoId
}

func (s *GetJobDetailResponseBodySnapshotJobDetail) SetCompleteTime(v string) *GetJobDetailResponseBodySnapshotJobDetail {
	s.CompleteTime = &v
	return s
}

func (s *GetJobDetailResponseBodySnapshotJobDetail) SetCreateTime(v string) *GetJobDetailResponseBodySnapshotJobDetail {
	s.CreateTime = &v
	return s
}

func (s *GetJobDetailResponseBodySnapshotJobDetail) SetJobId(v string) *GetJobDetailResponseBodySnapshotJobDetail {
	s.JobId = &v
	return s
}

func (s *GetJobDetailResponseBodySnapshotJobDetail) SetNormalConfig(v string) *GetJobDetailResponseBodySnapshotJobDetail {
	s.NormalConfig = &v
	return s
}

func (s *GetJobDetailResponseBodySnapshotJobDetail) SetSpriteConfig(v string) *GetJobDetailResponseBodySnapshotJobDetail {
	s.SpriteConfig = &v
	return s
}

func (s *GetJobDetailResponseBodySnapshotJobDetail) SetStatus(v string) *GetJobDetailResponseBodySnapshotJobDetail {
	s.Status = &v
	return s
}

func (s *GetJobDetailResponseBodySnapshotJobDetail) SetTrigger(v string) *GetJobDetailResponseBodySnapshotJobDetail {
	s.Trigger = &v
	return s
}

func (s *GetJobDetailResponseBodySnapshotJobDetail) SetUserId(v int64) *GetJobDetailResponseBodySnapshotJobDetail {
	s.UserId = &v
	return s
}

func (s *GetJobDetailResponseBodySnapshotJobDetail) SetVideoId(v string) *GetJobDetailResponseBodySnapshotJobDetail {
	s.VideoId = &v
	return s
}

func (s *GetJobDetailResponseBodySnapshotJobDetail) Validate() error {
	return dara.Validate(s)
}

type GetJobDetailResponseBodyTranscodeJobDetail struct {
	// The time when the task was complete.
	//
	// example:
	//
	// 2024-10-14T07:39:34Z
	CompleteTime *string `json:"CompleteTime,omitempty" xml:"CompleteTime,omitempty"`
	// The time when the task was created. The time follows the ISO 8601 standard in the YYYY-MM-DDTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2024-10-14T07:39:25Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The definition.
	//
	// example:
	//
	// HD
	Definition *string `json:"Definition,omitempty" xml:"Definition,omitempty"`
	// The ID of the task.
	//
	// example:
	//
	// 2dc1634e**********3f1d22d1a0174e
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The status of the task. Valid values:
	//
	// 	- Submitted
	//
	// 	- Transcoding
	//
	// 	- TranscodeSuccess
	//
	// 	- TranscodeFail
	//
	// 	- TranscodeCancelled
	//
	// example:
	//
	// TranscodeSuccess
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the template.
	//
	// example:
	//
	// dbfaaec9e**********bf0b81219244c
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The ID of the user who submitted the task.
	//
	// example:
	//
	// 139109*****84930
	UserId *int64 `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The ID of the media asset.
	//
	// example:
	//
	// 30e5d7**********bd900764de7c0102
	VideoId *string `json:"VideoId,omitempty" xml:"VideoId,omitempty"`
}

func (s GetJobDetailResponseBodyTranscodeJobDetail) String() string {
	return dara.Prettify(s)
}

func (s GetJobDetailResponseBodyTranscodeJobDetail) GoString() string {
	return s.String()
}

func (s *GetJobDetailResponseBodyTranscodeJobDetail) GetCompleteTime() *string {
	return s.CompleteTime
}

func (s *GetJobDetailResponseBodyTranscodeJobDetail) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetJobDetailResponseBodyTranscodeJobDetail) GetDefinition() *string {
	return s.Definition
}

func (s *GetJobDetailResponseBodyTranscodeJobDetail) GetJobId() *string {
	return s.JobId
}

func (s *GetJobDetailResponseBodyTranscodeJobDetail) GetStatus() *string {
	return s.Status
}

func (s *GetJobDetailResponseBodyTranscodeJobDetail) GetTemplateId() *string {
	return s.TemplateId
}

func (s *GetJobDetailResponseBodyTranscodeJobDetail) GetUserId() *int64 {
	return s.UserId
}

func (s *GetJobDetailResponseBodyTranscodeJobDetail) GetVideoId() *string {
	return s.VideoId
}

func (s *GetJobDetailResponseBodyTranscodeJobDetail) SetCompleteTime(v string) *GetJobDetailResponseBodyTranscodeJobDetail {
	s.CompleteTime = &v
	return s
}

func (s *GetJobDetailResponseBodyTranscodeJobDetail) SetCreateTime(v string) *GetJobDetailResponseBodyTranscodeJobDetail {
	s.CreateTime = &v
	return s
}

func (s *GetJobDetailResponseBodyTranscodeJobDetail) SetDefinition(v string) *GetJobDetailResponseBodyTranscodeJobDetail {
	s.Definition = &v
	return s
}

func (s *GetJobDetailResponseBodyTranscodeJobDetail) SetJobId(v string) *GetJobDetailResponseBodyTranscodeJobDetail {
	s.JobId = &v
	return s
}

func (s *GetJobDetailResponseBodyTranscodeJobDetail) SetStatus(v string) *GetJobDetailResponseBodyTranscodeJobDetail {
	s.Status = &v
	return s
}

func (s *GetJobDetailResponseBodyTranscodeJobDetail) SetTemplateId(v string) *GetJobDetailResponseBodyTranscodeJobDetail {
	s.TemplateId = &v
	return s
}

func (s *GetJobDetailResponseBodyTranscodeJobDetail) SetUserId(v int64) *GetJobDetailResponseBodyTranscodeJobDetail {
	s.UserId = &v
	return s
}

func (s *GetJobDetailResponseBodyTranscodeJobDetail) SetVideoId(v string) *GetJobDetailResponseBodyTranscodeJobDetail {
	s.VideoId = &v
	return s
}

func (s *GetJobDetailResponseBodyTranscodeJobDetail) Validate() error {
	return dara.Validate(s)
}

type GetJobDetailResponseBodyWorkflowTaskDetail struct {
	ActivityResults *string                                             `json:"ActivityResults,omitempty" xml:"ActivityResults,omitempty"`
	CreateTime      *string                                             `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	FinishTime      *string                                             `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	Status          *string                                             `json:"Status,omitempty" xml:"Status,omitempty"`
	TaskId          *string                                             `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	TaskInput       *string                                             `json:"TaskInput,omitempty" xml:"TaskInput,omitempty"`
	UserData        *string                                             `json:"UserData,omitempty" xml:"UserData,omitempty"`
	Workflow        *GetJobDetailResponseBodyWorkflowTaskDetailWorkflow `json:"Workflow,omitempty" xml:"Workflow,omitempty" type:"Struct"`
}

func (s GetJobDetailResponseBodyWorkflowTaskDetail) String() string {
	return dara.Prettify(s)
}

func (s GetJobDetailResponseBodyWorkflowTaskDetail) GoString() string {
	return s.String()
}

func (s *GetJobDetailResponseBodyWorkflowTaskDetail) GetActivityResults() *string {
	return s.ActivityResults
}

func (s *GetJobDetailResponseBodyWorkflowTaskDetail) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetJobDetailResponseBodyWorkflowTaskDetail) GetFinishTime() *string {
	return s.FinishTime
}

func (s *GetJobDetailResponseBodyWorkflowTaskDetail) GetStatus() *string {
	return s.Status
}

func (s *GetJobDetailResponseBodyWorkflowTaskDetail) GetTaskId() *string {
	return s.TaskId
}

func (s *GetJobDetailResponseBodyWorkflowTaskDetail) GetTaskInput() *string {
	return s.TaskInput
}

func (s *GetJobDetailResponseBodyWorkflowTaskDetail) GetUserData() *string {
	return s.UserData
}

func (s *GetJobDetailResponseBodyWorkflowTaskDetail) GetWorkflow() *GetJobDetailResponseBodyWorkflowTaskDetailWorkflow {
	return s.Workflow
}

func (s *GetJobDetailResponseBodyWorkflowTaskDetail) SetActivityResults(v string) *GetJobDetailResponseBodyWorkflowTaskDetail {
	s.ActivityResults = &v
	return s
}

func (s *GetJobDetailResponseBodyWorkflowTaskDetail) SetCreateTime(v string) *GetJobDetailResponseBodyWorkflowTaskDetail {
	s.CreateTime = &v
	return s
}

func (s *GetJobDetailResponseBodyWorkflowTaskDetail) SetFinishTime(v string) *GetJobDetailResponseBodyWorkflowTaskDetail {
	s.FinishTime = &v
	return s
}

func (s *GetJobDetailResponseBodyWorkflowTaskDetail) SetStatus(v string) *GetJobDetailResponseBodyWorkflowTaskDetail {
	s.Status = &v
	return s
}

func (s *GetJobDetailResponseBodyWorkflowTaskDetail) SetTaskId(v string) *GetJobDetailResponseBodyWorkflowTaskDetail {
	s.TaskId = &v
	return s
}

func (s *GetJobDetailResponseBodyWorkflowTaskDetail) SetTaskInput(v string) *GetJobDetailResponseBodyWorkflowTaskDetail {
	s.TaskInput = &v
	return s
}

func (s *GetJobDetailResponseBodyWorkflowTaskDetail) SetUserData(v string) *GetJobDetailResponseBodyWorkflowTaskDetail {
	s.UserData = &v
	return s
}

func (s *GetJobDetailResponseBodyWorkflowTaskDetail) SetWorkflow(v *GetJobDetailResponseBodyWorkflowTaskDetailWorkflow) *GetJobDetailResponseBodyWorkflowTaskDetail {
	s.Workflow = v
	return s
}

func (s *GetJobDetailResponseBodyWorkflowTaskDetail) Validate() error {
	if s.Workflow != nil {
		if err := s.Workflow.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetJobDetailResponseBodyWorkflowTaskDetailWorkflow struct {
	AppId        *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	CreateTime   *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Type         *string `json:"Type,omitempty" xml:"Type,omitempty"`
	WorkflowId   *string `json:"WorkflowId,omitempty" xml:"WorkflowId,omitempty"`
}

func (s GetJobDetailResponseBodyWorkflowTaskDetailWorkflow) String() string {
	return dara.Prettify(s)
}

func (s GetJobDetailResponseBodyWorkflowTaskDetailWorkflow) GoString() string {
	return s.String()
}

func (s *GetJobDetailResponseBodyWorkflowTaskDetailWorkflow) GetAppId() *string {
	return s.AppId
}

func (s *GetJobDetailResponseBodyWorkflowTaskDetailWorkflow) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetJobDetailResponseBodyWorkflowTaskDetailWorkflow) GetModifiedTime() *string {
	return s.ModifiedTime
}

func (s *GetJobDetailResponseBodyWorkflowTaskDetailWorkflow) GetName() *string {
	return s.Name
}

func (s *GetJobDetailResponseBodyWorkflowTaskDetailWorkflow) GetStatus() *string {
	return s.Status
}

func (s *GetJobDetailResponseBodyWorkflowTaskDetailWorkflow) GetType() *string {
	return s.Type
}

func (s *GetJobDetailResponseBodyWorkflowTaskDetailWorkflow) GetWorkflowId() *string {
	return s.WorkflowId
}

func (s *GetJobDetailResponseBodyWorkflowTaskDetailWorkflow) SetAppId(v string) *GetJobDetailResponseBodyWorkflowTaskDetailWorkflow {
	s.AppId = &v
	return s
}

func (s *GetJobDetailResponseBodyWorkflowTaskDetailWorkflow) SetCreateTime(v string) *GetJobDetailResponseBodyWorkflowTaskDetailWorkflow {
	s.CreateTime = &v
	return s
}

func (s *GetJobDetailResponseBodyWorkflowTaskDetailWorkflow) SetModifiedTime(v string) *GetJobDetailResponseBodyWorkflowTaskDetailWorkflow {
	s.ModifiedTime = &v
	return s
}

func (s *GetJobDetailResponseBodyWorkflowTaskDetailWorkflow) SetName(v string) *GetJobDetailResponseBodyWorkflowTaskDetailWorkflow {
	s.Name = &v
	return s
}

func (s *GetJobDetailResponseBodyWorkflowTaskDetailWorkflow) SetStatus(v string) *GetJobDetailResponseBodyWorkflowTaskDetailWorkflow {
	s.Status = &v
	return s
}

func (s *GetJobDetailResponseBodyWorkflowTaskDetailWorkflow) SetType(v string) *GetJobDetailResponseBodyWorkflowTaskDetailWorkflow {
	s.Type = &v
	return s
}

func (s *GetJobDetailResponseBodyWorkflowTaskDetailWorkflow) SetWorkflowId(v string) *GetJobDetailResponseBodyWorkflowTaskDetailWorkflow {
	s.WorkflowId = &v
	return s
}

func (s *GetJobDetailResponseBodyWorkflowTaskDetailWorkflow) Validate() error {
	return dara.Validate(s)
}
