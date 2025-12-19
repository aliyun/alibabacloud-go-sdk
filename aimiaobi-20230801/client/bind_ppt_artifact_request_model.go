// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindPptArtifactRequest interface {
	dara.Model
	String() string
	GoString() string
	SetArtifactId(v int32) *BindPptArtifactRequest
	GetArtifactId() *int32
	SetTaskId(v string) *BindPptArtifactRequest
	GetTaskId() *string
	SetWorkspaceId(v string) *BindPptArtifactRequest
	GetWorkspaceId() *string
}

type BindPptArtifactRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 12342
	ArtifactId *int32 `json:"ArtifactId,omitempty" xml:"ArtifactId,omitempty"`
	// This parameter is required.
	//
	// if can be null:
	// false
	//
	// example:
	//
	// 85da2bfe-6f05-41af-9841-d73c5bbf43a2
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// llm-xgpt3m25qdosdjr3
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s BindPptArtifactRequest) String() string {
	return dara.Prettify(s)
}

func (s BindPptArtifactRequest) GoString() string {
	return s.String()
}

func (s *BindPptArtifactRequest) GetArtifactId() *int32 {
	return s.ArtifactId
}

func (s *BindPptArtifactRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *BindPptArtifactRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *BindPptArtifactRequest) SetArtifactId(v int32) *BindPptArtifactRequest {
	s.ArtifactId = &v
	return s
}

func (s *BindPptArtifactRequest) SetTaskId(v string) *BindPptArtifactRequest {
	s.TaskId = &v
	return s
}

func (s *BindPptArtifactRequest) SetWorkspaceId(v string) *BindPptArtifactRequest {
	s.WorkspaceId = &v
	return s
}

func (s *BindPptArtifactRequest) Validate() error {
	return dara.Validate(s)
}
