// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePipelineRunRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoRunUntilStage(v string) *CreatePipelineRunRequest
	GetAutoRunUntilStage() *string
	SetDescription(v string) *CreatePipelineRunRequest
	GetDescription() *string
	SetObjectIds(v []*string) *CreatePipelineRunRequest
	GetObjectIds() []*string
	SetProjectId(v int64) *CreatePipelineRunRequest
	GetProjectId() *int64
	SetRunMode(v string) *CreatePipelineRunRequest
	GetRunMode() *string
	SetType(v string) *CreatePipelineRunRequest
	GetType() *string
}

type CreatePipelineRunRequest struct {
	// The code of the stage in the deployment process. This parameter takes effect only when RunMode is set to Auto. After the deployment process is created, it automatically runs to the specified stage.
	//
	//
	//
	//
	// > &lt;notice&gt;The specified stage is automatically completed. For example, if you set this parameter to DEV, the automatic run stops after the DEV stage reaches the terminal state.&gt;&lt;/notice&gt;
	//
	// example:
	//
	// DEV
	AutoRunUntilStage *string `json:"AutoRunUntilStage,omitempty" xml:"AutoRunUntilStage,omitempty"`
	// The description of the deployment process.
	//
	// example:
	//
	// This is a OdpsSQL-node publishing process. The function is XXXX.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The list of entity IDs that you want to deploy in this deployment process.
	//
	//
	//
	//
	// > &lt;notice&gt;Only a single entity and its child entities can be deployed at a time. Only the first entity in this array and its child entities are deployed. Make sure that the length of this array is 1. Entities beyond the first one are ignored.&gt;&lt;/notice&gt;
	//
	// This parameter is required.
	ObjectIds []*string `json:"ObjectIds,omitempty" xml:"ObjectIds,omitempty" type:"Repeated"`
	// The ID of the DataWorks workspace. You can log on to the [DataWorks console](https://workbench.data.aliyun.com/console) and go to the workspace configuration page to obtain the workspace ID.
	//
	// This parameter specifies the DataWorks workspace for this API call.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10000
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The run mode of the deployment process. Default value: Normal. If you set this parameter to Auto, the deployment process is automatically driven to the specified stage. This parameter is used together with the AutoRunUntilStage parameter.
	//
	//
	//
	//
	// Valid values:
	//
	//
	//
	//
	// - Normal
	//
	// - Auto
	//
	// example:
	//
	// Normal
	RunMode *string `json:"RunMode,omitempty" xml:"RunMode,omitempty"`
	// Specifies whether the deployment process is used to deploy or undeploy an entity.
	//
	//
	//
	//
	// - Online: deploy
	//
	// - Offline: undeploy
	//
	// This parameter is required.
	//
	// example:
	//
	// Online
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreatePipelineRunRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePipelineRunRequest) GoString() string {
	return s.String()
}

func (s *CreatePipelineRunRequest) GetAutoRunUntilStage() *string {
	return s.AutoRunUntilStage
}

func (s *CreatePipelineRunRequest) GetDescription() *string {
	return s.Description
}

func (s *CreatePipelineRunRequest) GetObjectIds() []*string {
	return s.ObjectIds
}

func (s *CreatePipelineRunRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *CreatePipelineRunRequest) GetRunMode() *string {
	return s.RunMode
}

func (s *CreatePipelineRunRequest) GetType() *string {
	return s.Type
}

func (s *CreatePipelineRunRequest) SetAutoRunUntilStage(v string) *CreatePipelineRunRequest {
	s.AutoRunUntilStage = &v
	return s
}

func (s *CreatePipelineRunRequest) SetDescription(v string) *CreatePipelineRunRequest {
	s.Description = &v
	return s
}

func (s *CreatePipelineRunRequest) SetObjectIds(v []*string) *CreatePipelineRunRequest {
	s.ObjectIds = v
	return s
}

func (s *CreatePipelineRunRequest) SetProjectId(v int64) *CreatePipelineRunRequest {
	s.ProjectId = &v
	return s
}

func (s *CreatePipelineRunRequest) SetRunMode(v string) *CreatePipelineRunRequest {
	s.RunMode = &v
	return s
}

func (s *CreatePipelineRunRequest) SetType(v string) *CreatePipelineRunRequest {
	s.Type = &v
	return s
}

func (s *CreatePipelineRunRequest) Validate() error {
	return dara.Validate(s)
}
