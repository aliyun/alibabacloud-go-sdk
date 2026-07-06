// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePipelineRunShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoRunUntilStage(v string) *CreatePipelineRunShrinkRequest
	GetAutoRunUntilStage() *string
	SetDescription(v string) *CreatePipelineRunShrinkRequest
	GetDescription() *string
	SetObjectIdsShrink(v string) *CreatePipelineRunShrinkRequest
	GetObjectIdsShrink() *string
	SetProjectId(v int64) *CreatePipelineRunShrinkRequest
	GetProjectId() *int64
	SetRunMode(v string) *CreatePipelineRunShrinkRequest
	GetRunMode() *string
	SetType(v string) *CreatePipelineRunShrinkRequest
	GetType() *string
}

type CreatePipelineRunShrinkRequest struct {
	// The code of the stage in the publish process. This parameter takes effect only when RunMode is set to Auto. After the publish process is created, it automatically runs to the specified stage.
	//
	// 	Notice: The specified stage is automatically completed. For example, if you set this parameter to DEV, the automatic run stops after the DEV stage reaches the desired state.
	//
	// example:
	//
	// DEV
	AutoRunUntilStage *string `json:"AutoRunUntilStage,omitempty" xml:"AutoRunUntilStage,omitempty"`
	// The description of the publish process.
	//
	// example:
	//
	// This is a OdpsSQL-node publishing process. The function is XXXX.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The list of entity IDs that you want to publish in this publish process.
	//
	// 	Notice: Only a single entity and its child entities can be published at a time. Only the first entity in this array and its child entities are published. Make sure that the length of this array is 1. Entities beyond the first one are ignored.
	//
	// This parameter is required.
	ObjectIdsShrink *string `json:"ObjectIds,omitempty" xml:"ObjectIds,omitempty"`
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
	// The run mode of the publish process. Default value: Normal. If you set this parameter to Auto, the publish process is automatically driven to the specified stage. This parameter is used together with the AutoRunUntilStage parameter.
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
	// Specifies whether the publish process is used to bring an entity online or offline.
	//
	// - Online: online
	//
	// - Offline: offline
	//
	// This parameter is required.
	//
	// example:
	//
	// Online
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreatePipelineRunShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePipelineRunShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreatePipelineRunShrinkRequest) GetAutoRunUntilStage() *string {
	return s.AutoRunUntilStage
}

func (s *CreatePipelineRunShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *CreatePipelineRunShrinkRequest) GetObjectIdsShrink() *string {
	return s.ObjectIdsShrink
}

func (s *CreatePipelineRunShrinkRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *CreatePipelineRunShrinkRequest) GetRunMode() *string {
	return s.RunMode
}

func (s *CreatePipelineRunShrinkRequest) GetType() *string {
	return s.Type
}

func (s *CreatePipelineRunShrinkRequest) SetAutoRunUntilStage(v string) *CreatePipelineRunShrinkRequest {
	s.AutoRunUntilStage = &v
	return s
}

func (s *CreatePipelineRunShrinkRequest) SetDescription(v string) *CreatePipelineRunShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreatePipelineRunShrinkRequest) SetObjectIdsShrink(v string) *CreatePipelineRunShrinkRequest {
	s.ObjectIdsShrink = &v
	return s
}

func (s *CreatePipelineRunShrinkRequest) SetProjectId(v int64) *CreatePipelineRunShrinkRequest {
	s.ProjectId = &v
	return s
}

func (s *CreatePipelineRunShrinkRequest) SetRunMode(v string) *CreatePipelineRunShrinkRequest {
	s.RunMode = &v
	return s
}

func (s *CreatePipelineRunShrinkRequest) SetType(v string) *CreatePipelineRunShrinkRequest {
	s.Type = &v
	return s
}

func (s *CreatePipelineRunShrinkRequest) Validate() error {
	return dara.Validate(s)
}
