// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPptArtifactRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPptArtifactId(v int32) *GetPptArtifactRequest
	GetPptArtifactId() *int32
	SetWorkspaceId(v string) *GetPptArtifactRequest
	GetWorkspaceId() *string
}

type GetPptArtifactRequest struct {
	// example:
	//
	// 5232136
	PptArtifactId *int32 `json:"PptArtifactId,omitempty" xml:"PptArtifactId,omitempty"`
	// example:
	//
	// llm-az2gglxxxx
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s GetPptArtifactRequest) String() string {
	return dara.Prettify(s)
}

func (s GetPptArtifactRequest) GoString() string {
	return s.String()
}

func (s *GetPptArtifactRequest) GetPptArtifactId() *int32 {
	return s.PptArtifactId
}

func (s *GetPptArtifactRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetPptArtifactRequest) SetPptArtifactId(v int32) *GetPptArtifactRequest {
	s.PptArtifactId = &v
	return s
}

func (s *GetPptArtifactRequest) SetWorkspaceId(v string) *GetPptArtifactRequest {
	s.WorkspaceId = &v
	return s
}

func (s *GetPptArtifactRequest) Validate() error {
	return dara.Validate(s)
}
