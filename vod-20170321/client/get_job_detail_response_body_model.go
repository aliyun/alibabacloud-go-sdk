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
	// The details of the AI task. This field has a value only when TaskType is AI.
	AIJobDetail *GetJobDetailResponseBodyAIJobDetail `json:"AIJobDetail,omitempty" xml:"AIJobDetail,omitempty" type:"Struct"`
	// The task type.
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
	// The details of the snapshot task. This field has a value only when jobType is Snapshot.
	SnapshotJobDetail *GetJobDetailResponseBodySnapshotJobDetail `json:"SnapshotJobDetail,omitempty" xml:"SnapshotJobDetail,omitempty" type:"Struct"`
	// The details of the transcoding task. This field has a value only when jobType is Transcode.
	TranscodeJobDetail *GetJobDetailResponseBodyTranscodeJobDetail `json:"TranscodeJobDetail,omitempty" xml:"TranscodeJobDetail,omitempty" type:"Struct"`
	// The details of the workflow task. This field has a value only when TaskType is Workflow.
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
	// The time when the task was completed.
	//
	// example:
	//
	// 2024-10-14T07:39:46Z
	CompleteTime *string `json:"CompleteTime,omitempty" xml:"CompleteTime,omitempty"`
	// The time when the task was created, in UTC. Format: YYYY-MM-DDTHH:MM:SSZ.
	//
	// example:
	//
	// 2024-10-14T07:39:25Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The task ID.
	//
	// example:
	//
	// 5c9dff751ba**********59d50a967f5
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The AI task type.
	//
	// example:
	//
	// AIVideoCensor
	JobType *string `json:"JobType,omitempty" xml:"JobType,omitempty"`
	// The media asset ID.
	//
	// example:
	//
	// 30e5d7**********bd900764de7c0102
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	// The task status. Valid values:
	//
	// - reserved: submitted.
	//
	// - init: started.
	//
	// - success: execution succeeded.
	//
	// - fail: execution failed.
	//
	// - processing: processing in progress.
	//
	// - analysing: analysis in progress.
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
	// The trigger method. Valid values:
	//
	// - Auto: automatically triggered by a workflow.
	//
	// - Manual: manually submitted.
	//
	// example:
	//
	// Auto
	Trigger *string `json:"Trigger,omitempty" xml:"Trigger,omitempty"`
	// The ID of the user who initiated the task.
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
	// The time when the task was completed.
	//
	// example:
	//
	// 2024-10-14T07:39:45Z
	CompleteTime *string `json:"CompleteTime,omitempty" xml:"CompleteTime,omitempty"`
	// The time when the task was created, in UTC. Format: YYYY-MM-DDTHH:MM:SSZ.
	//
	// example:
	//
	// 2024-10-14T07:39:25Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The task ID.
	//
	// example:
	//
	// 63df12s0**********4hdq249t82kr91
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The normal snapshot configuration.
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
	// The task status. Valid values:
	//
	// - Processing: processing in progress.
	//
	// - Fail: task failed.
	//
	// - Success: task succeeded.
	//
	// example:
	//
	// Success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The trigger method. Valid values:
	//
	// - Auto: automatically triggered by a workflow.
	//
	// - Manual: manually submitted.
	//
	// example:
	//
	// Auto
	Trigger *string `json:"Trigger,omitempty" xml:"Trigger,omitempty"`
	// The ID of the user who initiated the task.
	//
	// example:
	//
	// 139109*****84930
	UserId *int64 `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The media asset ID.
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
	// The time when the task was completed.
	//
	// example:
	//
	// 2024-10-14T07:39:34Z
	CompleteTime *string `json:"CompleteTime,omitempty" xml:"CompleteTime,omitempty"`
	// The time when the task was created, in UTC. Format: YYYY-MM-DDTHH:MM:SSZ.
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
	// The task ID.
	//
	// example:
	//
	// 2dc1634e**********3f1d22d1a0174e
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The task status. Valid values:
	//
	// - Submitted: submitted.
	//
	// - Transcoding: transcoding in progress.
	//
	// - TranscodeSuccess: transcoding succeeded.
	//
	// - TranscodeFail: transcoding failed.
	//
	// - TranscodeCancelled: transcoding canceled.
	//
	// example:
	//
	// TranscodeSuccess
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The template ID.
	//
	// example:
	//
	// dbfaaec9e**********bf0b81219244c
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The ID of the user who initiated the task.
	//
	// example:
	//
	// 139109*****84930
	UserId *int64 `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The media asset ID.
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
	// The processing results of each workflow node, in JSON format.
	//
	// example:
	//
	// {\\"VodSnapshot_123\\":{\\"ActivityId\\":\\"VodSnapshot\\",\\"ActivityInstanceId\\":\\"c8cf62d53bef4e04bf703976bae6d0b9\\",\\"EndTime\\":\\"2025-03-27T08:15:51Z\\",\\"Result\\":\\"{\\\\\\"RequestId\\\\\\":\\\\\\"8B3649AF-5A6B-1099-BEB6-164D81067398\\\\\\",\\\\\\"EventType\\\\\\":\\\\\\"SnapshotComplete\\\\\\",\\\\\\"UserId\\\\\\":1797131669910763,\\\\\\"MessageBody\\\\\\":{\\\\\\"Status\\\\\\":\\\\\\"success\\\\\\",\\\\\\"VideoId\\\\\\":\\\\\\"00f985a50ae371f0ad1c4106e0ea0102\\\\\\",\\\\\\"EventType\\\\\\":\\\\\\"SnapshotComplete\\\\\\",\\\\\\"EventTime\\\\\\":\\\\\\"2025-03-27T08:15:50Z\\\\\\",\\\\\\"TriggerSource\\\\\\":\\\\\\"{\\\\\\\\\\\\\\"ActivityInstanceId\\\\\\\\\\\\\\":\\\\\\\\\\\\\\"c8cf62d53bef4e04bf703976bae6d0b9\\\\\\\\\\\\\\",\\\\\\\\\\\\\\"BizType\\\\\\\\\\\\\\":\\\\\\\\\\\\\\"ice-workflow\\\\\\\\\\\\\\"}\\\\\\"}}\\",\\"StartTime\\":\\"2025-03-27T08:15:47Z\\",\\"Status\\":\\"Succeed\\"},\\"Translate_zh_en\\":{\\"ActivityId\\":\\"VodTranslation\\",\\"ActivityInstanceId\\":\\"c043a872bb044763a3d293a5c2458b50\\",\\"EndTime\\":\\"2025-03-27T08:20:19Z\\",\\"Result\\":\\"{\\\\\\"Type\\\\\\":\\\\\\"VideoTranslationAll\\\\\\",\\\\\\"Success\\\\\\":false}\\",\\"StartTime\\":\\"2025-03-27T08:15:46Z\\",\\"Status\\":\\"Failed\\"},\\"Act_Start\\":{\\"ActivityId\\":\\"start\\",\\"ActivityInstanceId\\":\\"8a9402f4ff064084bf496707fb2d664a\\",\\"Result\\":\\"{\\\\\\"Type\\\\\\":\\\\\\"Media\\\\\\",\\\\\\"bizType\\\\\\":6,\\\\\\"Media\\\\\\":\\\\\\"00f985a50ae371f0ad1c4106e0ea0102\\\\\\",\\\\\\"Title\\\\\\":\\\\\\"2.mp4\\\\\\",\\\\\\"taskInput\\\\\\":\\\\\\"{\\\\\\\\\\\\\\"Type\\\\\\\\\\\\\\":\\\\\\\\\\\\\\"Media\\\\\\\\\\\\\\",\\\\\\\\\\\\\\"Media\\\\\\\\\\\\\\":\\\\\\\\\\\\\\"00f985a50ae371f0ad1c4106e0ea0102\\\\\\\\\\\\\\"}\\\\\\",\\\\\\"userTaskInput\\\\\\":\\\\\\"{\\\\\\\\\\\\\\"Type\\\\\\\\\\\\\\":\\\\\\\\\\\\\\"Media\\\\\\\\\\\\\\",\\\\\\\\\\\\\\"Media\\\\\\\\\\\\\\":\\\\\\\\\\\\\\"00f985a50ae371f0ad1c4106e0ea0102\\\\\\\\\\\\\\",\\\\\\\\\\\\\\"Title\\\\\\\\\\\\\\":\\\\\\\\\\\\\\"2.mp4\\\\\\\\\\\\\\",\\\\\\\\\\\\\\"StorageLocation\\\\\\\\\\\\\\":\\\\\\\\\\\\\\"yiming-pre.oss-cn-shanghai.aliyuncs.com\\\\\\\\\\\\\\"}\\\\\\",\\\\\\"StorageLocation\\\\\\":\\\\\\"yiming-pre.oss-cn-shanghai.aliyuncs.com\\\\\\",\\\\\\"callerUid\\\\\\":1797131669910763,\\\\\\"CUR_NODE_NAME\\\\\\":\\\\\\"Act_Start\\\\\\"}\\",\\"StartTime\\":\\"2025-03-27T08:15:45Z\\",\\"Status\\":\\"Succeed\\"},\\"VodDynamicImage_123\\":{\\"ActivityId\\":\\"VodDynamicImage\\",\\"ActivityInstanceId\\":\\"26e8dab82ab84110b1150f146caf633c\\",\\"EndTime\\":\\"2025-03-27T08:15:55Z\\",\\"Result\\":\\"{\\\\\\"RequestId\\\\\\":\\\\\\"7120B5D5-430F-14AD-8922-577F072DDD64\\\\\\",\\\\\\"EventType\\\\\\":\\\\\\"DynamicImageComplete\\\\\\",\\\\\\"UserId\\\\\\":1797131669910763,\\\\\\"MessageBody\\\\\\":{\\\\\\"Status\\\\\\":\\\\\\"success\\\\\\",\\\\\\"VideoId\\\\\\":\\\\\\"00f985a50ae371f0ad1c4106e0ea0102\\\\\\",\\\\\\"EventType\\\\\\":\\\\\\"DynamicImageComplete\\\\\\",\\\\\\"EventTime\\\\\\":\\\\\\"2025-03-27T08:15:52Z\\\\\\",\\\\\\"TriggerSource\\\\\\":\\\\\\"{\\\\\\\\\\\\\\"ActivityInstanceId\\\\\\\\\\\\\\":\\\\\\\\\\\\\\"26e8dab82ab84110b1150f146caf633c\\\\\\\\\\\\\\",\\\\\\\\\\\\\\"BizType\\\\\\\\\\\\\\":\\\\\\\\\\\\\\"ice-workflow\\\\\\\\\\\\\\"}\\\\\\"}}\\",\\"StartTime\\":\\"2025-03-27T08:15:47Z\\",\\"Status\\":\\"Succeed\\"}}
	ActivityResults *string `json:"ActivityResults,omitempty" xml:"ActivityResults,omitempty"`
	// The time when the task was created, in UTC. Format: YYYY-MM-DDTHH:MM:SSZ.
	//
	// example:
	//
	// 2025-03-27T08:15:46Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The time when the task was completed.
	//
	// example:
	//
	// 2025-03-27T08:20:19Z
	FinishTime *string `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	// The task status. Valid values:
	//
	// - Init: started.
	//
	// - Processing: processing in progress.
	//
	// - Succeed: succeeded.
	//
	// - Failed: failed.
	//
	// - Canceled: canceled.
	//
	// - Skip: skipped.
	//
	// example:
	//
	// Succeed
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The task ID.
	//
	// example:
	//
	// 46ecc024******92c8e26237e51
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The media asset information.
	//
	// example:
	//
	// {\\"Type\\":\\"Media\\",\\"Media\\":\\"00f985a50ae371f0ad1c4106e0ea0102\\",\\"Title\\":\\"2.mp4\\",\\"StorageLocation\\":\\"yiming-pre.oss-cn-shanghai.aliyuncs.com\\"}
	TaskInput *string `json:"TaskInput,omitempty" xml:"TaskInput,omitempty"`
	// The custom settings. The value is a JSON string that supports message callback, upload acceleration, and other settings.
	//
	// example:
	//
	// {"MessageCallback":{"CallbackURL":"http://example.aliyundoc.com"},"Extend":{"localId":"*****","test":"www"}}
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
	// The workflow details.
	Workflow *GetJobDetailResponseBodyWorkflowTaskDetailWorkflow `json:"Workflow,omitempty" xml:"Workflow,omitempty" type:"Struct"`
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
	// The application ID.
	//
	// example:
	//
	// app-1000000
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The time when the workflow was created, in UTC. Format: YYYY-MM-DDTHH:MM:SSZ.
	//
	// example:
	//
	// 2025-03-26T05:50:14Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The time when the workflow was last modified.
	//
	// example:
	//
	// 2025-03-26T07:51:55Z
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// The workflow name.
	//
	// example:
	//
	// All_Activity_New_1_app-1000000
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The workflow status. Valid values:
	//
	// - Active: activated.
	//
	// - Inactive: not activated.
	//
	// - Deleted: deleted.
	//
	// example:
	//
	// Active
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The workflow type. Not populated by default.
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The workflow ID.
	//
	// example:
	//
	// vw_09d6*****b5c5b19a0c891e02
	WorkflowId *string `json:"WorkflowId,omitempty" xml:"WorkflowId,omitempty"`
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
