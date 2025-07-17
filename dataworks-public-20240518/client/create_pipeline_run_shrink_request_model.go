// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePipelineRunShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreatePipelineRunShrinkRequest
	GetDescription() *string
	SetObjectIdsShrink(v string) *CreatePipelineRunShrinkRequest
	GetObjectIdsShrink() *string
	SetProjectId(v int64) *CreatePipelineRunShrinkRequest
	GetProjectId() *int64
	SetType(v string) *CreatePipelineRunShrinkRequest
	GetType() *string
}

type CreatePipelineRunShrinkRequest struct {
	// The description of the process.
	//
	// example:
	//
	// This is a OdpsSQL-node publishing process. The function is XXXX.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The IDs of entities to which you want to apply the process.
	//
	// >  A process can be applied to only a single entity and its child entities. If you specify multiple entities in the array, the process is applied only to the first entity in the array and its child entities. Make sure that the array in your request contains only one element. Extra elements will be ignored.
	//
	// This parameter is required.
	ObjectIdsShrink *string `json:"ObjectIds,omitempty" xml:"ObjectIds,omitempty"`
	// The DataWorks workspace ID. You can log on to the [DataWorks console](https://workbench.data.aliyun.com/console) and go to the Workspace page to query the ID. You must configure this parameter to specify the DataWorks workspace to which the API operation is applied.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10000
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// Specifies whether to deploy or undeploy the entity. Valid values:
	//
	// 	- Online: deploys the entity.
	//
	// 	- Offline: undeploys the entity.
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

func (s *CreatePipelineRunShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *CreatePipelineRunShrinkRequest) GetObjectIdsShrink() *string {
	return s.ObjectIdsShrink
}

func (s *CreatePipelineRunShrinkRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *CreatePipelineRunShrinkRequest) GetType() *string {
	return s.Type
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

func (s *CreatePipelineRunShrinkRequest) SetType(v string) *CreatePipelineRunShrinkRequest {
	s.Type = &v
	return s
}

func (s *CreatePipelineRunShrinkRequest) Validate() error {
	return dara.Validate(s)
}
