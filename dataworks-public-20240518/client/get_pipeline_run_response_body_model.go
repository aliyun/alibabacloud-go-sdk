// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPipelineRunResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPipeline(v *GetPipelineRunResponseBodyPipeline) *GetPipelineRunResponseBody
	GetPipeline() *GetPipelineRunResponseBodyPipeline
	SetRequestId(v string) *GetPipelineRunResponseBody
	GetRequestId() *string
}

type GetPipelineRunResponseBody struct {
	// The details of the pipeline run.
	Pipeline *GetPipelineRunResponseBodyPipeline `json:"Pipeline,omitempty" xml:"Pipeline,omitempty" type:"Struct"`
	// The ID of the request. You can use this ID to locate logs and troubleshoot issues.
	//
	// example:
	//
	// 08468352-032C-5262-AEDC-68C9FA05XXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetPipelineRunResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetPipelineRunResponseBody) GoString() string {
	return s.String()
}

func (s *GetPipelineRunResponseBody) GetPipeline() *GetPipelineRunResponseBodyPipeline {
	return s.Pipeline
}

func (s *GetPipelineRunResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetPipelineRunResponseBody) SetPipeline(v *GetPipelineRunResponseBodyPipeline) *GetPipelineRunResponseBody {
	s.Pipeline = v
	return s
}

func (s *GetPipelineRunResponseBody) SetRequestId(v string) *GetPipelineRunResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPipelineRunResponseBody) Validate() error {
	if s.Pipeline != nil {
		if err := s.Pipeline.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetPipelineRunResponseBodyPipeline struct {
	// The time when the pipeline run was created. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1724984066000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the user who created the pipeline run.
	//
	// example:
	//
	// 137946317766XXXX
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// The description of the pipeline run.
	//
	// example:
	//
	// 发布流程描述信息
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the pipeline run.
	//
	// example:
	//
	// a7ef0634-20ec-4a7c-a214-54020f91XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The error message returned if the pipeline run fails.
	//
	// example:
	//
	// Error message
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The time when the pipeline run was last modified. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1724984066000
	ModifyTime *int64 `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The ID of the DataWorks workspace.
	//
	// example:
	//
	// 56160
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The stages in the pipeline run.
	Stages []*GetPipelineRunResponseBodyPipelineStages `json:"Stages,omitempty" xml:"Stages,omitempty" type:"Repeated"`
	// The status of the pipeline run.
	//
	// Valid values:
	//
	// - `Init`: The pipeline run is being initialized.
	//
	// - `Running`: The pipeline run is in progress.
	//
	// - `Success`: The pipeline run succeeded.
	//
	// - `Fail`: The pipeline run failed.
	//
	// - `Terminated`: The pipeline run was terminated.
	//
	// - `Canceled`: The pipeline run was canceled.
	//
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetPipelineRunResponseBodyPipeline) String() string {
	return dara.Prettify(s)
}

func (s GetPipelineRunResponseBodyPipeline) GoString() string {
	return s.String()
}

func (s *GetPipelineRunResponseBodyPipeline) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *GetPipelineRunResponseBodyPipeline) GetCreator() *string {
	return s.Creator
}

func (s *GetPipelineRunResponseBodyPipeline) GetDescription() *string {
	return s.Description
}

func (s *GetPipelineRunResponseBodyPipeline) GetId() *string {
	return s.Id
}

func (s *GetPipelineRunResponseBodyPipeline) GetMessage() *string {
	return s.Message
}

func (s *GetPipelineRunResponseBodyPipeline) GetModifyTime() *int64 {
	return s.ModifyTime
}

func (s *GetPipelineRunResponseBodyPipeline) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetPipelineRunResponseBodyPipeline) GetStages() []*GetPipelineRunResponseBodyPipelineStages {
	return s.Stages
}

func (s *GetPipelineRunResponseBodyPipeline) GetStatus() *string {
	return s.Status
}

func (s *GetPipelineRunResponseBodyPipeline) SetCreateTime(v int64) *GetPipelineRunResponseBodyPipeline {
	s.CreateTime = &v
	return s
}

func (s *GetPipelineRunResponseBodyPipeline) SetCreator(v string) *GetPipelineRunResponseBodyPipeline {
	s.Creator = &v
	return s
}

func (s *GetPipelineRunResponseBodyPipeline) SetDescription(v string) *GetPipelineRunResponseBodyPipeline {
	s.Description = &v
	return s
}

func (s *GetPipelineRunResponseBodyPipeline) SetId(v string) *GetPipelineRunResponseBodyPipeline {
	s.Id = &v
	return s
}

func (s *GetPipelineRunResponseBodyPipeline) SetMessage(v string) *GetPipelineRunResponseBodyPipeline {
	s.Message = &v
	return s
}

func (s *GetPipelineRunResponseBodyPipeline) SetModifyTime(v int64) *GetPipelineRunResponseBodyPipeline {
	s.ModifyTime = &v
	return s
}

func (s *GetPipelineRunResponseBodyPipeline) SetProjectId(v int64) *GetPipelineRunResponseBodyPipeline {
	s.ProjectId = &v
	return s
}

func (s *GetPipelineRunResponseBodyPipeline) SetStages(v []*GetPipelineRunResponseBodyPipelineStages) *GetPipelineRunResponseBodyPipeline {
	s.Stages = v
	return s
}

func (s *GetPipelineRunResponseBodyPipeline) SetStatus(v string) *GetPipelineRunResponseBodyPipeline {
	s.Status = &v
	return s
}

func (s *GetPipelineRunResponseBodyPipeline) Validate() error {
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

type GetPipelineRunResponseBodyPipelineStages struct {
	// The code that identifies the stage.
	//
	// example:
	//
	// DEV_CHECK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The description of the stage.
	//
	// example:
	//
	// Phase description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Detailed information about the stage.
	Detail map[string]interface{} `json:"Detail,omitempty" xml:"Detail,omitempty"`
	// The exception message returned if the stage fails.
	//
	// example:
	//
	// Exception information XXX
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The name of the stage.
	//
	// example:
	//
	// Publish package build
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The status of the stage.
	//
	// Valid values:
	//
	// - `Init`: The stage is being initialized.
	//
	// - `Running`: The stage is in progress.
	//
	// - `Success`: The stage succeeded.
	//
	// - `Fail`: The stage failed.
	//
	// - `Terminated`: The stage was terminated.
	//
	// - `Canceled`: The stage was canceled.
	//
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The sequence number of the stage within the pipeline.
	//
	// example:
	//
	// 1
	Step *int32 `json:"Step,omitempty" xml:"Step,omitempty"`
	// The type of the stage.
	//
	// Valid values:
	//
	// - `Deploy`: A deployment operation.
	//
	// - `Check`: A check operation.
	//
	// - `Offline`: An offline operation.
	//
	// - `Build`: A build operation.
	//
	// - `Delete`: A delete operation.
	//
	// example:
	//
	// Check
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetPipelineRunResponseBodyPipelineStages) String() string {
	return dara.Prettify(s)
}

func (s GetPipelineRunResponseBodyPipelineStages) GoString() string {
	return s.String()
}

func (s *GetPipelineRunResponseBodyPipelineStages) GetCode() *string {
	return s.Code
}

func (s *GetPipelineRunResponseBodyPipelineStages) GetDescription() *string {
	return s.Description
}

func (s *GetPipelineRunResponseBodyPipelineStages) GetDetail() map[string]interface{} {
	return s.Detail
}

func (s *GetPipelineRunResponseBodyPipelineStages) GetMessage() *string {
	return s.Message
}

func (s *GetPipelineRunResponseBodyPipelineStages) GetName() *string {
	return s.Name
}

func (s *GetPipelineRunResponseBodyPipelineStages) GetStatus() *string {
	return s.Status
}

func (s *GetPipelineRunResponseBodyPipelineStages) GetStep() *int32 {
	return s.Step
}

func (s *GetPipelineRunResponseBodyPipelineStages) GetType() *string {
	return s.Type
}

func (s *GetPipelineRunResponseBodyPipelineStages) SetCode(v string) *GetPipelineRunResponseBodyPipelineStages {
	s.Code = &v
	return s
}

func (s *GetPipelineRunResponseBodyPipelineStages) SetDescription(v string) *GetPipelineRunResponseBodyPipelineStages {
	s.Description = &v
	return s
}

func (s *GetPipelineRunResponseBodyPipelineStages) SetDetail(v map[string]interface{}) *GetPipelineRunResponseBodyPipelineStages {
	s.Detail = v
	return s
}

func (s *GetPipelineRunResponseBodyPipelineStages) SetMessage(v string) *GetPipelineRunResponseBodyPipelineStages {
	s.Message = &v
	return s
}

func (s *GetPipelineRunResponseBodyPipelineStages) SetName(v string) *GetPipelineRunResponseBodyPipelineStages {
	s.Name = &v
	return s
}

func (s *GetPipelineRunResponseBodyPipelineStages) SetStatus(v string) *GetPipelineRunResponseBodyPipelineStages {
	s.Status = &v
	return s
}

func (s *GetPipelineRunResponseBodyPipelineStages) SetStep(v int32) *GetPipelineRunResponseBodyPipelineStages {
	s.Step = &v
	return s
}

func (s *GetPipelineRunResponseBodyPipelineStages) SetType(v string) *GetPipelineRunResponseBodyPipelineStages {
	s.Type = &v
	return s
}

func (s *GetPipelineRunResponseBodyPipelineStages) Validate() error {
	return dara.Validate(s)
}
