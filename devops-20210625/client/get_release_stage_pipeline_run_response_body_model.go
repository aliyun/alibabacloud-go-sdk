// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetReleaseStagePipelineRunResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPipelineRun(v *GetReleaseStagePipelineRunResponseBodyPipelineRun) *GetReleaseStagePipelineRunResponseBody
	GetPipelineRun() *GetReleaseStagePipelineRunResponseBodyPipelineRun
}

type GetReleaseStagePipelineRunResponseBody struct {
	PipelineRun *GetReleaseStagePipelineRunResponseBodyPipelineRun `json:"pipelineRun,omitempty" xml:"pipelineRun,omitempty" type:"Struct"`
}

func (s GetReleaseStagePipelineRunResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetReleaseStagePipelineRunResponseBody) GoString() string {
	return s.String()
}

func (s *GetReleaseStagePipelineRunResponseBody) GetPipelineRun() *GetReleaseStagePipelineRunResponseBodyPipelineRun {
	return s.PipelineRun
}

func (s *GetReleaseStagePipelineRunResponseBody) SetPipelineRun(v *GetReleaseStagePipelineRunResponseBodyPipelineRun) *GetReleaseStagePipelineRunResponseBody {
	s.PipelineRun = v
	return s
}

func (s *GetReleaseStagePipelineRunResponseBody) Validate() error {
	if s.PipelineRun != nil {
		if err := s.PipelineRun.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetReleaseStagePipelineRunResponseBodyPipelineRun struct {
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
	PipelineRunId *int64                                                      `json:"pipelineRunId,omitempty" xml:"pipelineRunId,omitempty"`
	Sources       []*GetReleaseStagePipelineRunResponseBodyPipelineRunSources `json:"sources,omitempty" xml:"sources,omitempty" type:"Repeated"`
	StageGroup    [][]*string                                                 `json:"stageGroup,omitempty" xml:"stageGroup,omitempty" type:"Repeated"`
	Stages        []*GetReleaseStagePipelineRunResponseBodyPipelineRunStages  `json:"stages,omitempty" xml:"stages,omitempty" type:"Repeated"`
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

func (s GetReleaseStagePipelineRunResponseBodyPipelineRun) String() string {
	return dara.Prettify(s)
}

func (s GetReleaseStagePipelineRunResponseBodyPipelineRun) GoString() string {
	return s.String()
}

func (s *GetReleaseStagePipelineRunResponseBodyPipelineRun) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *GetReleaseStagePipelineRunResponseBodyPipelineRun) GetCreatorAccountId() *string {
	return s.CreatorAccountId
}

func (s *GetReleaseStagePipelineRunResponseBodyPipelineRun) GetModifierAccountId() *string {
	return s.ModifierAccountId
}

func (s *GetReleaseStagePipelineRunResponseBodyPipelineRun) GetPipelineId() *int64 {
	return s.PipelineId
}

func (s *GetReleaseStagePipelineRunResponseBodyPipelineRun) GetPipelineRunId() *int64 {
	return s.PipelineRunId
}

func (s *GetReleaseStagePipelineRunResponseBodyPipelineRun) GetSources() []*GetReleaseStagePipelineRunResponseBodyPipelineRunSources {
	return s.Sources
}

func (s *GetReleaseStagePipelineRunResponseBodyPipelineRun) GetStageGroup() [][]*string {
	return s.StageGroup
}

func (s *GetReleaseStagePipelineRunResponseBodyPipelineRun) GetStages() []*GetReleaseStagePipelineRunResponseBodyPipelineRunStages {
	return s.Stages
}

func (s *GetReleaseStagePipelineRunResponseBodyPipelineRun) GetStatus() *string {
	return s.Status
}

func (s *GetReleaseStagePipelineRunResponseBodyPipelineRun) GetTriggerMode() *int32 {
	return s.TriggerMode
}

func (s *GetReleaseStagePipelineRunResponseBodyPipelineRun) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *GetReleaseStagePipelineRunResponseBodyPipelineRun) SetCreateTime(v int64) *GetReleaseStagePipelineRunResponseBodyPipelineRun {
	s.CreateTime = &v
	return s
}

func (s *GetReleaseStagePipelineRunResponseBodyPipelineRun) SetCreatorAccountId(v string) *GetReleaseStagePipelineRunResponseBodyPipelineRun {
	s.CreatorAccountId = &v
	return s
}

func (s *GetReleaseStagePipelineRunResponseBodyPipelineRun) SetModifierAccountId(v string) *GetReleaseStagePipelineRunResponseBodyPipelineRun {
	s.ModifierAccountId = &v
	return s
}

func (s *GetReleaseStagePipelineRunResponseBodyPipelineRun) SetPipelineId(v int64) *GetReleaseStagePipelineRunResponseBodyPipelineRun {
	s.PipelineId = &v
	return s
}

func (s *GetReleaseStagePipelineRunResponseBodyPipelineRun) SetPipelineRunId(v int64) *GetReleaseStagePipelineRunResponseBodyPipelineRun {
	s.PipelineRunId = &v
	return s
}

func (s *GetReleaseStagePipelineRunResponseBodyPipelineRun) SetSources(v []*GetReleaseStagePipelineRunResponseBodyPipelineRunSources) *GetReleaseStagePipelineRunResponseBodyPipelineRun {
	s.Sources = v
	return s
}

func (s *GetReleaseStagePipelineRunResponseBodyPipelineRun) SetStageGroup(v [][]*string) *GetReleaseStagePipelineRunResponseBodyPipelineRun {
	s.StageGroup = v
	return s
}

func (s *GetReleaseStagePipelineRunResponseBodyPipelineRun) SetStages(v []*GetReleaseStagePipelineRunResponseBodyPipelineRunStages) *GetReleaseStagePipelineRunResponseBodyPipelineRun {
	s.Stages = v
	return s
}

func (s *GetReleaseStagePipelineRunResponseBodyPipelineRun) SetStatus(v string) *GetReleaseStagePipelineRunResponseBodyPipelineRun {
	s.Status = &v
	return s
}

func (s *GetReleaseStagePipelineRunResponseBodyPipelineRun) SetTriggerMode(v int32) *GetReleaseStagePipelineRunResponseBodyPipelineRun {
	s.TriggerMode = &v
	return s
}

func (s *GetReleaseStagePipelineRunResponseBodyPipelineRun) SetUpdateTime(v int64) *GetReleaseStagePipelineRunResponseBodyPipelineRun {
	s.UpdateTime = &v
	return s
}

func (s *GetReleaseStagePipelineRunResponseBodyPipelineRun) Validate() error {
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

type GetReleaseStagePipelineRunResponseBodyPipelineRunSources struct {
	Data *GetReleaseStagePipelineRunResponseBodyPipelineRunSourcesData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// assaaaaaasasasa
	Sign *string `json:"sign,omitempty" xml:"sign,omitempty"`
	// example:
	//
	// Codeup
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetReleaseStagePipelineRunResponseBodyPipelineRunSources) String() string {
	return dara.Prettify(s)
}

func (s GetReleaseStagePipelineRunResponseBodyPipelineRunSources) GoString() string {
	return s.String()
}

func (s *GetReleaseStagePipelineRunResponseBodyPipelineRunSources) GetData() *GetReleaseStagePipelineRunResponseBodyPipelineRunSourcesData {
	return s.Data
}

func (s *GetReleaseStagePipelineRunResponseBodyPipelineRunSources) GetSign() *string {
	return s.Sign
}

func (s *GetReleaseStagePipelineRunResponseBodyPipelineRunSources) GetType() *string {
	return s.Type
}

func (s *GetReleaseStagePipelineRunResponseBodyPipelineRunSources) SetData(v *GetReleaseStagePipelineRunResponseBodyPipelineRunSourcesData) *GetReleaseStagePipelineRunResponseBodyPipelineRunSources {
	s.Data = v
	return s
}

func (s *GetReleaseStagePipelineRunResponseBodyPipelineRunSources) SetSign(v string) *GetReleaseStagePipelineRunResponseBodyPipelineRunSources {
	s.Sign = &v
	return s
}

func (s *GetReleaseStagePipelineRunResponseBodyPipelineRunSources) SetType(v string) *GetReleaseStagePipelineRunResponseBodyPipelineRunSources {
	s.Type = &v
	return s
}

func (s *GetReleaseStagePipelineRunResponseBodyPipelineRunSources) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetReleaseStagePipelineRunResponseBodyPipelineRunSourcesData struct {
	// example:
	//
	// master
	Branch *string `json:"branch,omitempty" xml:"branch,omitempty"`
	// example:
	//
	// {}
	Commit *string `json:"commit,omitempty" xml:"commit,omitempty"`
	// example:
	//
	// http://codeup.aliyun.com/a.git
	Repo *string `json:"repo,omitempty" xml:"repo,omitempty"`
}

func (s GetReleaseStagePipelineRunResponseBodyPipelineRunSourcesData) String() string {
	return dara.Prettify(s)
}

func (s GetReleaseStagePipelineRunResponseBodyPipelineRunSourcesData) GoString() string {
	return s.String()
}

func (s *GetReleaseStagePipelineRunResponseBodyPipelineRunSourcesData) GetBranch() *string {
	return s.Branch
}

func (s *GetReleaseStagePipelineRunResponseBodyPipelineRunSourcesData) GetCommit() *string {
	return s.Commit
}

func (s *GetReleaseStagePipelineRunResponseBodyPipelineRunSourcesData) GetRepo() *string {
	return s.Repo
}

func (s *GetReleaseStagePipelineRunResponseBodyPipelineRunSourcesData) SetBranch(v string) *GetReleaseStagePipelineRunResponseBodyPipelineRunSourcesData {
	s.Branch = &v
	return s
}

func (s *GetReleaseStagePipelineRunResponseBodyPipelineRunSourcesData) SetCommit(v string) *GetReleaseStagePipelineRunResponseBodyPipelineRunSourcesData {
	s.Commit = &v
	return s
}

func (s *GetReleaseStagePipelineRunResponseBodyPipelineRunSourcesData) SetRepo(v string) *GetReleaseStagePipelineRunResponseBodyPipelineRunSourcesData {
	s.Repo = &v
	return s
}

func (s *GetReleaseStagePipelineRunResponseBodyPipelineRunSourcesData) Validate() error {
	return dara.Validate(s)
}

type GetReleaseStagePipelineRunResponseBodyPipelineRunStages struct {
	Name      *string                                                           `json:"name,omitempty" xml:"name,omitempty"`
	StageInfo *GetReleaseStagePipelineRunResponseBodyPipelineRunStagesStageInfo `json:"stageInfo,omitempty" xml:"stageInfo,omitempty" type:"Struct"`
}

func (s GetReleaseStagePipelineRunResponseBodyPipelineRunStages) String() string {
	return dara.Prettify(s)
}

func (s GetReleaseStagePipelineRunResponseBodyPipelineRunStages) GoString() string {
	return s.String()
}

func (s *GetReleaseStagePipelineRunResponseBodyPipelineRunStages) GetName() *string {
	return s.Name
}

func (s *GetReleaseStagePipelineRunResponseBodyPipelineRunStages) GetStageInfo() *GetReleaseStagePipelineRunResponseBodyPipelineRunStagesStageInfo {
	return s.StageInfo
}

func (s *GetReleaseStagePipelineRunResponseBodyPipelineRunStages) SetName(v string) *GetReleaseStagePipelineRunResponseBodyPipelineRunStages {
	s.Name = &v
	return s
}

func (s *GetReleaseStagePipelineRunResponseBodyPipelineRunStages) SetStageInfo(v *GetReleaseStagePipelineRunResponseBodyPipelineRunStagesStageInfo) *GetReleaseStagePipelineRunResponseBodyPipelineRunStages {
	s.StageInfo = v
	return s
}

func (s *GetReleaseStagePipelineRunResponseBodyPipelineRunStages) Validate() error {
	if s.StageInfo != nil {
		if err := s.StageInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetReleaseStagePipelineRunResponseBodyPipelineRunStagesStageInfo struct {
	// example:
	//
	// 1586863220000
	EndTime *int64                                                                  `json:"endTime,omitempty" xml:"endTime,omitempty"`
	Jobs    []*GetReleaseStagePipelineRunResponseBodyPipelineRunStagesStageInfoJobs `json:"jobs,omitempty" xml:"jobs,omitempty" type:"Repeated"`
	Name    *string                                                                 `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 1586863220000
	StartTime *int64 `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// example:
	//
	// RUNNING
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s GetReleaseStagePipelineRunResponseBodyPipelineRunStagesStageInfo) String() string {
	return dara.Prettify(s)
}

func (s GetReleaseStagePipelineRunResponseBodyPipelineRunStagesStageInfo) GoString() string {
	return s.String()
}

func (s *GetReleaseStagePipelineRunResponseBodyPipelineRunStagesStageInfo) GetEndTime() *int64 {
	return s.EndTime
}

func (s *GetReleaseStagePipelineRunResponseBodyPipelineRunStagesStageInfo) GetJobs() []*GetReleaseStagePipelineRunResponseBodyPipelineRunStagesStageInfoJobs {
	return s.Jobs
}

func (s *GetReleaseStagePipelineRunResponseBodyPipelineRunStagesStageInfo) GetName() *string {
	return s.Name
}

func (s *GetReleaseStagePipelineRunResponseBodyPipelineRunStagesStageInfo) GetStartTime() *int64 {
	return s.StartTime
}

func (s *GetReleaseStagePipelineRunResponseBodyPipelineRunStagesStageInfo) GetStatus() *string {
	return s.Status
}

func (s *GetReleaseStagePipelineRunResponseBodyPipelineRunStagesStageInfo) SetEndTime(v int64) *GetReleaseStagePipelineRunResponseBodyPipelineRunStagesStageInfo {
	s.EndTime = &v
	return s
}

func (s *GetReleaseStagePipelineRunResponseBodyPipelineRunStagesStageInfo) SetJobs(v []*GetReleaseStagePipelineRunResponseBodyPipelineRunStagesStageInfoJobs) *GetReleaseStagePipelineRunResponseBodyPipelineRunStagesStageInfo {
	s.Jobs = v
	return s
}

func (s *GetReleaseStagePipelineRunResponseBodyPipelineRunStagesStageInfo) SetName(v string) *GetReleaseStagePipelineRunResponseBodyPipelineRunStagesStageInfo {
	s.Name = &v
	return s
}

func (s *GetReleaseStagePipelineRunResponseBodyPipelineRunStagesStageInfo) SetStartTime(v int64) *GetReleaseStagePipelineRunResponseBodyPipelineRunStagesStageInfo {
	s.StartTime = &v
	return s
}

func (s *GetReleaseStagePipelineRunResponseBodyPipelineRunStagesStageInfo) SetStatus(v string) *GetReleaseStagePipelineRunResponseBodyPipelineRunStagesStageInfo {
	s.Status = &v
	return s
}

func (s *GetReleaseStagePipelineRunResponseBodyPipelineRunStagesStageInfo) Validate() error {
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

type GetReleaseStagePipelineRunResponseBodyPipelineRunStagesStageInfoJobs struct {
	Actions []*GetReleaseStagePipelineRunResponseBodyPipelineRunStagesStageInfoJobsActions `json:"actions,omitempty" xml:"actions,omitempty" type:"Repeated"`
	// example:
	//
	// 1586863220000
	EndTime *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// example:
	//
	// 21212
	Id   *int64  `json:"id,omitempty" xml:"id,omitempty"`
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

func (s GetReleaseStagePipelineRunResponseBodyPipelineRunStagesStageInfoJobs) String() string {
	return dara.Prettify(s)
}

func (s GetReleaseStagePipelineRunResponseBodyPipelineRunStagesStageInfoJobs) GoString() string {
	return s.String()
}

func (s *GetReleaseStagePipelineRunResponseBodyPipelineRunStagesStageInfoJobs) GetActions() []*GetReleaseStagePipelineRunResponseBodyPipelineRunStagesStageInfoJobsActions {
	return s.Actions
}

func (s *GetReleaseStagePipelineRunResponseBodyPipelineRunStagesStageInfoJobs) GetEndTime() *int64 {
	return s.EndTime
}

func (s *GetReleaseStagePipelineRunResponseBodyPipelineRunStagesStageInfoJobs) GetId() *int64 {
	return s.Id
}

func (s *GetReleaseStagePipelineRunResponseBodyPipelineRunStagesStageInfoJobs) GetName() *string {
	return s.Name
}

func (s *GetReleaseStagePipelineRunResponseBodyPipelineRunStagesStageInfoJobs) GetParams() *string {
	return s.Params
}

func (s *GetReleaseStagePipelineRunResponseBodyPipelineRunStagesStageInfoJobs) GetStartTime() *int64 {
	return s.StartTime
}

func (s *GetReleaseStagePipelineRunResponseBodyPipelineRunStagesStageInfoJobs) GetStatus() *string {
	return s.Status
}

func (s *GetReleaseStagePipelineRunResponseBodyPipelineRunStagesStageInfoJobs) SetActions(v []*GetReleaseStagePipelineRunResponseBodyPipelineRunStagesStageInfoJobsActions) *GetReleaseStagePipelineRunResponseBodyPipelineRunStagesStageInfoJobs {
	s.Actions = v
	return s
}

func (s *GetReleaseStagePipelineRunResponseBodyPipelineRunStagesStageInfoJobs) SetEndTime(v int64) *GetReleaseStagePipelineRunResponseBodyPipelineRunStagesStageInfoJobs {
	s.EndTime = &v
	return s
}

func (s *GetReleaseStagePipelineRunResponseBodyPipelineRunStagesStageInfoJobs) SetId(v int64) *GetReleaseStagePipelineRunResponseBodyPipelineRunStagesStageInfoJobs {
	s.Id = &v
	return s
}

func (s *GetReleaseStagePipelineRunResponseBodyPipelineRunStagesStageInfoJobs) SetName(v string) *GetReleaseStagePipelineRunResponseBodyPipelineRunStagesStageInfoJobs {
	s.Name = &v
	return s
}

func (s *GetReleaseStagePipelineRunResponseBodyPipelineRunStagesStageInfoJobs) SetParams(v string) *GetReleaseStagePipelineRunResponseBodyPipelineRunStagesStageInfoJobs {
	s.Params = &v
	return s
}

func (s *GetReleaseStagePipelineRunResponseBodyPipelineRunStagesStageInfoJobs) SetStartTime(v int64) *GetReleaseStagePipelineRunResponseBodyPipelineRunStagesStageInfoJobs {
	s.StartTime = &v
	return s
}

func (s *GetReleaseStagePipelineRunResponseBodyPipelineRunStagesStageInfoJobs) SetStatus(v string) *GetReleaseStagePipelineRunResponseBodyPipelineRunStagesStageInfoJobs {
	s.Status = &v
	return s
}

func (s *GetReleaseStagePipelineRunResponseBodyPipelineRunStagesStageInfoJobs) Validate() error {
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

type GetReleaseStagePipelineRunResponseBodyPipelineRunStagesStageInfoJobsActions struct {
	// example:
	//
	// true
	Disable *bool `json:"disable,omitempty" xml:"disable,omitempty"`
	// example:
	//
	// {}
	Params interface{} `json:"params,omitempty" xml:"params,omitempty"`
	// example:
	//
	// PassPipelineValidate
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetReleaseStagePipelineRunResponseBodyPipelineRunStagesStageInfoJobsActions) String() string {
	return dara.Prettify(s)
}

func (s GetReleaseStagePipelineRunResponseBodyPipelineRunStagesStageInfoJobsActions) GoString() string {
	return s.String()
}

func (s *GetReleaseStagePipelineRunResponseBodyPipelineRunStagesStageInfoJobsActions) GetDisable() *bool {
	return s.Disable
}

func (s *GetReleaseStagePipelineRunResponseBodyPipelineRunStagesStageInfoJobsActions) GetParams() interface{} {
	return s.Params
}

func (s *GetReleaseStagePipelineRunResponseBodyPipelineRunStagesStageInfoJobsActions) GetType() *string {
	return s.Type
}

func (s *GetReleaseStagePipelineRunResponseBodyPipelineRunStagesStageInfoJobsActions) SetDisable(v bool) *GetReleaseStagePipelineRunResponseBodyPipelineRunStagesStageInfoJobsActions {
	s.Disable = &v
	return s
}

func (s *GetReleaseStagePipelineRunResponseBodyPipelineRunStagesStageInfoJobsActions) SetParams(v interface{}) *GetReleaseStagePipelineRunResponseBodyPipelineRunStagesStageInfoJobsActions {
	s.Params = v
	return s
}

func (s *GetReleaseStagePipelineRunResponseBodyPipelineRunStagesStageInfoJobsActions) SetType(v string) *GetReleaseStagePipelineRunResponseBodyPipelineRunStagesStageInfoJobsActions {
	s.Type = &v
	return s
}

func (s *GetReleaseStagePipelineRunResponseBodyPipelineRunStagesStageInfoJobsActions) Validate() error {
	return dara.Validate(s)
}
