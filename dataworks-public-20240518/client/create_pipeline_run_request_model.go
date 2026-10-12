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
	// The code of the stage to which the deployment process runs. This parameter takes effect only when the RunMode parameter is set to Auto. After the deployment process is created, it automatically runs to the specified stage.
	//
	// 	Notice: The specified stage is automatically completed. For example, if you set this parameter to DEV, the automatic running stops after the DEV stage reaches the desired state.
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
	// The list of entity IDs to be deployed in this deployment process.
	//
	// 	Notice: Currently, only a single entity and its child entities can be deployed. Therefore, only the first entity and its child entities in the array are deployed. Make sure that the array contains only one element. Additional elements are ignored.
	//
	// This parameter is required.
	ObjectIds []*string `json:"ObjectIds,omitempty" xml:"ObjectIds,omitempty" type:"Repeated"`
	// The ID of the DataWorks workspace. You can log on to the [DataWorks console](https://workbench.data.aliyun.com/console) and go to the Workspace Configuration page to obtain the workspace ID.
	//
	// This parameter specifies the DataWorks workspace on which the API operation is performed.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10000
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The running mode of the deployment process. Default value: Normal. If you set this parameter to Auto, the deployment process automatically runs to the specified stage. This parameter must be used together with the AutoRunUntilStage parameter.
	//
	// Valid values:
	//
	// - Normal
	//
	// - Auto
	//
	// example:
	//
	// Normal
	RunMode *string `json:"RunMode,omitempty" xml:"RunMode,omitempty"`
	// Specifies the type of the deployment process.
	//
	// Valid values:
	//
	// - Online: Goes online.
	//
	// - Offline: Goes offline.
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
