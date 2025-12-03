// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPipelineRunResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *GetPipelineRunResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetPipelineRunResponseBody
	GetErrorMessage() *string
	SetPipelineRun(v *GetPipelineRunResponseBodyPipelineRun) *GetPipelineRunResponseBody
	GetPipelineRun() *GetPipelineRunResponseBodyPipelineRun
	SetRequestId(v string) *GetPipelineRunResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetPipelineRunResponseBody
	GetSuccess() *bool
}

type GetPipelineRunResponseBody struct {
	// example:
	//
	// ""
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// ""
	ErrorMessage *string                                `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	PipelineRun  *GetPipelineRunResponseBodyPipelineRun `json:"pipelineRun,omitempty" xml:"pipelineRun,omitempty" type:"Struct"`
	// example:
	//
	// ASSDS-ASSASX-XSAXSA-XSAXSAXS
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetPipelineRunResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetPipelineRunResponseBody) GoString() string {
	return s.String()
}

func (s *GetPipelineRunResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetPipelineRunResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetPipelineRunResponseBody) GetPipelineRun() *GetPipelineRunResponseBodyPipelineRun {
	return s.PipelineRun
}

func (s *GetPipelineRunResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetPipelineRunResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetPipelineRunResponseBody) SetErrorCode(v string) *GetPipelineRunResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetPipelineRunResponseBody) SetErrorMessage(v string) *GetPipelineRunResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetPipelineRunResponseBody) SetPipelineRun(v *GetPipelineRunResponseBodyPipelineRun) *GetPipelineRunResponseBody {
	s.PipelineRun = v
	return s
}

func (s *GetPipelineRunResponseBody) SetRequestId(v string) *GetPipelineRunResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPipelineRunResponseBody) SetSuccess(v bool) *GetPipelineRunResponseBody {
	s.Success = &v
	return s
}

func (s *GetPipelineRunResponseBody) Validate() error {
	if s.PipelineRun != nil {
		if err := s.PipelineRun.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetPipelineRunResponseBodyPipelineRun struct {
	// example:
	//
	// 1586863220000
	CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// 1111111111
	CreatorAccountId *string `json:"creatorAccountId,omitempty" xml:"creatorAccountId,omitempty"`
	// example:
	//
	// 11111111111
	ModifierAccountId *string `json:"modifierAccountId,omitempty" xml:"modifierAccountId,omitempty"`
	// example:
	//
	// 1234
	PipelineId *int64 `json:"pipelineId,omitempty" xml:"pipelineId,omitempty"`
	// example:
	//
	// 1
	PipelineRunId *int64                                          `json:"pipelineRunId,omitempty" xml:"pipelineRunId,omitempty"`
	Sources       []*GetPipelineRunResponseBodyPipelineRunSources `json:"sources,omitempty" xml:"sources,omitempty" type:"Repeated"`
	StageGroup    [][]*string                                     `json:"stageGroup,omitempty" xml:"stageGroup,omitempty" type:"Repeated"`
	Stages        []*GetPipelineRunResponseBodyPipelineRunStages  `json:"stages,omitempty" xml:"stages,omitempty" type:"Repeated"`
	// example:
	//
	// SUCCESS
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// 1
	TriggerMode *int32 `json:"triggerMode,omitempty" xml:"triggerMode,omitempty"`
	// example:
	//
	// 1586863220000
	UpdateTime *int64 `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
}

func (s GetPipelineRunResponseBodyPipelineRun) String() string {
	return dara.Prettify(s)
}

func (s GetPipelineRunResponseBodyPipelineRun) GoString() string {
	return s.String()
}

func (s *GetPipelineRunResponseBodyPipelineRun) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *GetPipelineRunResponseBodyPipelineRun) GetCreatorAccountId() *string {
	return s.CreatorAccountId
}

func (s *GetPipelineRunResponseBodyPipelineRun) GetModifierAccountId() *string {
	return s.ModifierAccountId
}

func (s *GetPipelineRunResponseBodyPipelineRun) GetPipelineId() *int64 {
	return s.PipelineId
}

func (s *GetPipelineRunResponseBodyPipelineRun) GetPipelineRunId() *int64 {
	return s.PipelineRunId
}

func (s *GetPipelineRunResponseBodyPipelineRun) GetSources() []*GetPipelineRunResponseBodyPipelineRunSources {
	return s.Sources
}

func (s *GetPipelineRunResponseBodyPipelineRun) GetStageGroup() [][]*string {
	return s.StageGroup
}

func (s *GetPipelineRunResponseBodyPipelineRun) GetStages() []*GetPipelineRunResponseBodyPipelineRunStages {
	return s.Stages
}

func (s *GetPipelineRunResponseBodyPipelineRun) GetStatus() *string {
	return s.Status
}

func (s *GetPipelineRunResponseBodyPipelineRun) GetTriggerMode() *int32 {
	return s.TriggerMode
}

func (s *GetPipelineRunResponseBodyPipelineRun) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *GetPipelineRunResponseBodyPipelineRun) SetCreateTime(v int64) *GetPipelineRunResponseBodyPipelineRun {
	s.CreateTime = &v
	return s
}

func (s *GetPipelineRunResponseBodyPipelineRun) SetCreatorAccountId(v string) *GetPipelineRunResponseBodyPipelineRun {
	s.CreatorAccountId = &v
	return s
}

func (s *GetPipelineRunResponseBodyPipelineRun) SetModifierAccountId(v string) *GetPipelineRunResponseBodyPipelineRun {
	s.ModifierAccountId = &v
	return s
}

func (s *GetPipelineRunResponseBodyPipelineRun) SetPipelineId(v int64) *GetPipelineRunResponseBodyPipelineRun {
	s.PipelineId = &v
	return s
}

func (s *GetPipelineRunResponseBodyPipelineRun) SetPipelineRunId(v int64) *GetPipelineRunResponseBodyPipelineRun {
	s.PipelineRunId = &v
	return s
}

func (s *GetPipelineRunResponseBodyPipelineRun) SetSources(v []*GetPipelineRunResponseBodyPipelineRunSources) *GetPipelineRunResponseBodyPipelineRun {
	s.Sources = v
	return s
}

func (s *GetPipelineRunResponseBodyPipelineRun) SetStageGroup(v [][]*string) *GetPipelineRunResponseBodyPipelineRun {
	s.StageGroup = v
	return s
}

func (s *GetPipelineRunResponseBodyPipelineRun) SetStages(v []*GetPipelineRunResponseBodyPipelineRunStages) *GetPipelineRunResponseBodyPipelineRun {
	s.Stages = v
	return s
}

func (s *GetPipelineRunResponseBodyPipelineRun) SetStatus(v string) *GetPipelineRunResponseBodyPipelineRun {
	s.Status = &v
	return s
}

func (s *GetPipelineRunResponseBodyPipelineRun) SetTriggerMode(v int32) *GetPipelineRunResponseBodyPipelineRun {
	s.TriggerMode = &v
	return s
}

func (s *GetPipelineRunResponseBodyPipelineRun) SetUpdateTime(v int64) *GetPipelineRunResponseBodyPipelineRun {
	s.UpdateTime = &v
	return s
}

func (s *GetPipelineRunResponseBodyPipelineRun) Validate() error {
	if s.Sources != nil {
		for _, item := range s.Sources {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Stages != nil {
		for _, item := range s.Stages {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetPipelineRunResponseBodyPipelineRunSources struct {
	Data *GetPipelineRunResponseBodyPipelineRunSourcesData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// assaaaaaasasasa
	Sign *string `json:"sign,omitempty" xml:"sign,omitempty"`
	// example:
	//
	// Codeup
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetPipelineRunResponseBodyPipelineRunSources) String() string {
	return dara.Prettify(s)
}

func (s GetPipelineRunResponseBodyPipelineRunSources) GoString() string {
	return s.String()
}

func (s *GetPipelineRunResponseBodyPipelineRunSources) GetData() *GetPipelineRunResponseBodyPipelineRunSourcesData {
	return s.Data
}

func (s *GetPipelineRunResponseBodyPipelineRunSources) GetSign() *string {
	return s.Sign
}

func (s *GetPipelineRunResponseBodyPipelineRunSources) GetType() *string {
	return s.Type
}

func (s *GetPipelineRunResponseBodyPipelineRunSources) SetData(v *GetPipelineRunResponseBodyPipelineRunSourcesData) *GetPipelineRunResponseBodyPipelineRunSources {
	s.Data = v
	return s
}

func (s *GetPipelineRunResponseBodyPipelineRunSources) SetSign(v string) *GetPipelineRunResponseBodyPipelineRunSources {
	s.Sign = &v
	return s
}

func (s *GetPipelineRunResponseBodyPipelineRunSources) SetType(v string) *GetPipelineRunResponseBodyPipelineRunSources {
	s.Type = &v
	return s
}

func (s *GetPipelineRunResponseBodyPipelineRunSources) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetPipelineRunResponseBodyPipelineRunSourcesData struct {
	// example:
	//
	// master
	Branch *string `json:"branch,omitempty" xml:"branch,omitempty"`
	// example:
	//
	// {}
	Commint *string `json:"commint,omitempty" xml:"commint,omitempty"`
	// example:
	//
	// http://codeup.aliyun.com/a.git
	Repo *string `json:"repo,omitempty" xml:"repo,omitempty"`
}

func (s GetPipelineRunResponseBodyPipelineRunSourcesData) String() string {
	return dara.Prettify(s)
}

func (s GetPipelineRunResponseBodyPipelineRunSourcesData) GoString() string {
	return s.String()
}

func (s *GetPipelineRunResponseBodyPipelineRunSourcesData) GetBranch() *string {
	return s.Branch
}

func (s *GetPipelineRunResponseBodyPipelineRunSourcesData) GetCommint() *string {
	return s.Commint
}

func (s *GetPipelineRunResponseBodyPipelineRunSourcesData) GetRepo() *string {
	return s.Repo
}

func (s *GetPipelineRunResponseBodyPipelineRunSourcesData) SetBranch(v string) *GetPipelineRunResponseBodyPipelineRunSourcesData {
	s.Branch = &v
	return s
}

func (s *GetPipelineRunResponseBodyPipelineRunSourcesData) SetCommint(v string) *GetPipelineRunResponseBodyPipelineRunSourcesData {
	s.Commint = &v
	return s
}

func (s *GetPipelineRunResponseBodyPipelineRunSourcesData) SetRepo(v string) *GetPipelineRunResponseBodyPipelineRunSourcesData {
	s.Repo = &v
	return s
}

func (s *GetPipelineRunResponseBodyPipelineRunSourcesData) Validate() error {
	return dara.Validate(s)
}

type GetPipelineRunResponseBodyPipelineRunStages struct {
	// example:
	//
	// Java构建
	Name      *string                                               `json:"name,omitempty" xml:"name,omitempty"`
	StageInfo *GetPipelineRunResponseBodyPipelineRunStagesStageInfo `json:"stageInfo,omitempty" xml:"stageInfo,omitempty" type:"Struct"`
}

func (s GetPipelineRunResponseBodyPipelineRunStages) String() string {
	return dara.Prettify(s)
}

func (s GetPipelineRunResponseBodyPipelineRunStages) GoString() string {
	return s.String()
}

func (s *GetPipelineRunResponseBodyPipelineRunStages) GetName() *string {
	return s.Name
}

func (s *GetPipelineRunResponseBodyPipelineRunStages) GetStageInfo() *GetPipelineRunResponseBodyPipelineRunStagesStageInfo {
	return s.StageInfo
}

func (s *GetPipelineRunResponseBodyPipelineRunStages) SetName(v string) *GetPipelineRunResponseBodyPipelineRunStages {
	s.Name = &v
	return s
}

func (s *GetPipelineRunResponseBodyPipelineRunStages) SetStageInfo(v *GetPipelineRunResponseBodyPipelineRunStagesStageInfo) *GetPipelineRunResponseBodyPipelineRunStages {
	s.StageInfo = v
	return s
}

func (s *GetPipelineRunResponseBodyPipelineRunStages) Validate() error {
	if s.StageInfo != nil {
		if err := s.StageInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetPipelineRunResponseBodyPipelineRunStagesStageInfo struct {
	// example:
	//
	// 1586863220000
	EndTime *int64                                                      `json:"endTime,omitempty" xml:"endTime,omitempty"`
	Jobs    []*GetPipelineRunResponseBodyPipelineRunStagesStageInfoJobs `json:"jobs,omitempty" xml:"jobs,omitempty" type:"Repeated"`
	// example:
	//
	// Java构建
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 1586863220000
	StartTime *int64 `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// example:
	//
	// RUNNING
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s GetPipelineRunResponseBodyPipelineRunStagesStageInfo) String() string {
	return dara.Prettify(s)
}

func (s GetPipelineRunResponseBodyPipelineRunStagesStageInfo) GoString() string {
	return s.String()
}

func (s *GetPipelineRunResponseBodyPipelineRunStagesStageInfo) GetEndTime() *int64 {
	return s.EndTime
}

func (s *GetPipelineRunResponseBodyPipelineRunStagesStageInfo) GetJobs() []*GetPipelineRunResponseBodyPipelineRunStagesStageInfoJobs {
	return s.Jobs
}

func (s *GetPipelineRunResponseBodyPipelineRunStagesStageInfo) GetName() *string {
	return s.Name
}

func (s *GetPipelineRunResponseBodyPipelineRunStagesStageInfo) GetStartTime() *int64 {
	return s.StartTime
}

func (s *GetPipelineRunResponseBodyPipelineRunStagesStageInfo) GetStatus() *string {
	return s.Status
}

func (s *GetPipelineRunResponseBodyPipelineRunStagesStageInfo) SetEndTime(v int64) *GetPipelineRunResponseBodyPipelineRunStagesStageInfo {
	s.EndTime = &v
	return s
}

func (s *GetPipelineRunResponseBodyPipelineRunStagesStageInfo) SetJobs(v []*GetPipelineRunResponseBodyPipelineRunStagesStageInfoJobs) *GetPipelineRunResponseBodyPipelineRunStagesStageInfo {
	s.Jobs = v
	return s
}

func (s *GetPipelineRunResponseBodyPipelineRunStagesStageInfo) SetName(v string) *GetPipelineRunResponseBodyPipelineRunStagesStageInfo {
	s.Name = &v
	return s
}

func (s *GetPipelineRunResponseBodyPipelineRunStagesStageInfo) SetStartTime(v int64) *GetPipelineRunResponseBodyPipelineRunStagesStageInfo {
	s.StartTime = &v
	return s
}

func (s *GetPipelineRunResponseBodyPipelineRunStagesStageInfo) SetStatus(v string) *GetPipelineRunResponseBodyPipelineRunStagesStageInfo {
	s.Status = &v
	return s
}

func (s *GetPipelineRunResponseBodyPipelineRunStagesStageInfo) Validate() error {
	if s.Jobs != nil {
		for _, item := range s.Jobs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetPipelineRunResponseBodyPipelineRunStagesStageInfoJobs struct {
	Actions []*GetPipelineRunResponseBodyPipelineRunStagesStageInfoJobsActions `json:"actions,omitempty" xml:"actions,omitempty" type:"Repeated"`
	// example:
	//
	// 1586863220000
	EndTime *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// example:
	//
	// 21212
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// java构建
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// {}
	Params *string `json:"params,omitempty" xml:"params,omitempty"`
	// example:
	//
	// 1586863220000
	StartTime *int64 `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// example:
	//
	// RUNNING
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s GetPipelineRunResponseBodyPipelineRunStagesStageInfoJobs) String() string {
	return dara.Prettify(s)
}

func (s GetPipelineRunResponseBodyPipelineRunStagesStageInfoJobs) GoString() string {
	return s.String()
}

func (s *GetPipelineRunResponseBodyPipelineRunStagesStageInfoJobs) GetActions() []*GetPipelineRunResponseBodyPipelineRunStagesStageInfoJobsActions {
	return s.Actions
}

func (s *GetPipelineRunResponseBodyPipelineRunStagesStageInfoJobs) GetEndTime() *int64 {
	return s.EndTime
}

func (s *GetPipelineRunResponseBodyPipelineRunStagesStageInfoJobs) GetId() *int64 {
	return s.Id
}

func (s *GetPipelineRunResponseBodyPipelineRunStagesStageInfoJobs) GetName() *string {
	return s.Name
}

func (s *GetPipelineRunResponseBodyPipelineRunStagesStageInfoJobs) GetParams() *string {
	return s.Params
}

func (s *GetPipelineRunResponseBodyPipelineRunStagesStageInfoJobs) GetStartTime() *int64 {
	return s.StartTime
}

func (s *GetPipelineRunResponseBodyPipelineRunStagesStageInfoJobs) GetStatus() *string {
	return s.Status
}

func (s *GetPipelineRunResponseBodyPipelineRunStagesStageInfoJobs) SetActions(v []*GetPipelineRunResponseBodyPipelineRunStagesStageInfoJobsActions) *GetPipelineRunResponseBodyPipelineRunStagesStageInfoJobs {
	s.Actions = v
	return s
}

func (s *GetPipelineRunResponseBodyPipelineRunStagesStageInfoJobs) SetEndTime(v int64) *GetPipelineRunResponseBodyPipelineRunStagesStageInfoJobs {
	s.EndTime = &v
	return s
}

func (s *GetPipelineRunResponseBodyPipelineRunStagesStageInfoJobs) SetId(v int64) *GetPipelineRunResponseBodyPipelineRunStagesStageInfoJobs {
	s.Id = &v
	return s
}

func (s *GetPipelineRunResponseBodyPipelineRunStagesStageInfoJobs) SetName(v string) *GetPipelineRunResponseBodyPipelineRunStagesStageInfoJobs {
	s.Name = &v
	return s
}

func (s *GetPipelineRunResponseBodyPipelineRunStagesStageInfoJobs) SetParams(v string) *GetPipelineRunResponseBodyPipelineRunStagesStageInfoJobs {
	s.Params = &v
	return s
}

func (s *GetPipelineRunResponseBodyPipelineRunStagesStageInfoJobs) SetStartTime(v int64) *GetPipelineRunResponseBodyPipelineRunStagesStageInfoJobs {
	s.StartTime = &v
	return s
}

func (s *GetPipelineRunResponseBodyPipelineRunStagesStageInfoJobs) SetStatus(v string) *GetPipelineRunResponseBodyPipelineRunStagesStageInfoJobs {
	s.Status = &v
	return s
}

func (s *GetPipelineRunResponseBodyPipelineRunStagesStageInfoJobs) Validate() error {
	if s.Actions != nil {
		for _, item := range s.Actions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetPipelineRunResponseBodyPipelineRunStagesStageInfoJobsActions struct {
	// example:
	//
	// true
	Disable *bool `json:"disable,omitempty" xml:"disable,omitempty"`
	// example:
	//
	// {}
	Params map[string]interface{} `json:"params,omitempty" xml:"params,omitempty"`
	// example:
	//
	// PassPipelineValidate
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetPipelineRunResponseBodyPipelineRunStagesStageInfoJobsActions) String() string {
	return dara.Prettify(s)
}

func (s GetPipelineRunResponseBodyPipelineRunStagesStageInfoJobsActions) GoString() string {
	return s.String()
}

func (s *GetPipelineRunResponseBodyPipelineRunStagesStageInfoJobsActions) GetDisable() *bool {
	return s.Disable
}

func (s *GetPipelineRunResponseBodyPipelineRunStagesStageInfoJobsActions) GetParams() map[string]interface{} {
	return s.Params
}

func (s *GetPipelineRunResponseBodyPipelineRunStagesStageInfoJobsActions) GetType() *string {
	return s.Type
}

func (s *GetPipelineRunResponseBodyPipelineRunStagesStageInfoJobsActions) SetDisable(v bool) *GetPipelineRunResponseBodyPipelineRunStagesStageInfoJobsActions {
	s.Disable = &v
	return s
}

func (s *GetPipelineRunResponseBodyPipelineRunStagesStageInfoJobsActions) SetParams(v map[string]interface{}) *GetPipelineRunResponseBodyPipelineRunStagesStageInfoJobsActions {
	s.Params = v
	return s
}

func (s *GetPipelineRunResponseBodyPipelineRunStagesStageInfoJobsActions) SetType(v string) *GetPipelineRunResponseBodyPipelineRunStagesStageInfoJobsActions {
	s.Type = &v
	return s
}

func (s *GetPipelineRunResponseBodyPipelineRunStagesStageInfoJobsActions) Validate() error {
	return dara.Validate(s)
}
