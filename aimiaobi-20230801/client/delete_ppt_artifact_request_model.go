// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePptArtifactRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPptArtifactId(v string) *DeletePptArtifactRequest
	GetPptArtifactId() *string
	SetWorkspaceId(v string) *DeletePptArtifactRequest
	GetWorkspaceId() *string
}

type DeletePptArtifactRequest struct {
	PptArtifactId *string `json:"PptArtifactId,omitempty" xml:"PptArtifactId,omitempty"`
	// example:
	//
	// llm-xx
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s DeletePptArtifactRequest) String() string {
	return dara.Prettify(s)
}

func (s DeletePptArtifactRequest) GoString() string {
	return s.String()
}

func (s *DeletePptArtifactRequest) GetPptArtifactId() *string {
	return s.PptArtifactId
}

func (s *DeletePptArtifactRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *DeletePptArtifactRequest) SetPptArtifactId(v string) *DeletePptArtifactRequest {
	s.PptArtifactId = &v
	return s
}

func (s *DeletePptArtifactRequest) SetWorkspaceId(v string) *DeletePptArtifactRequest {
	s.WorkspaceId = &v
	return s
}

func (s *DeletePptArtifactRequest) Validate() error {
	return dara.Validate(s)
}
