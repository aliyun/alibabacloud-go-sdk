// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPptArtifactExportResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExportTaskId(v string) *GetPptArtifactExportResultRequest
	GetExportTaskId() *string
	SetWorkspaceId(v string) *GetPptArtifactExportResultRequest
	GetWorkspaceId() *string
}

type GetPptArtifactExportResultRequest struct {
	// example:
	//
	// 15aeb61b-cdeb-4b70-94d7-99518040647e
	ExportTaskId *string `json:"ExportTaskId,omitempty" xml:"ExportTaskId,omitempty"`
	// example:
	//
	// llm-xxxx
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s GetPptArtifactExportResultRequest) String() string {
	return dara.Prettify(s)
}

func (s GetPptArtifactExportResultRequest) GoString() string {
	return s.String()
}

func (s *GetPptArtifactExportResultRequest) GetExportTaskId() *string {
	return s.ExportTaskId
}

func (s *GetPptArtifactExportResultRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetPptArtifactExportResultRequest) SetExportTaskId(v string) *GetPptArtifactExportResultRequest {
	s.ExportTaskId = &v
	return s
}

func (s *GetPptArtifactExportResultRequest) SetWorkspaceId(v string) *GetPptArtifactExportResultRequest {
	s.WorkspaceId = &v
	return s
}

func (s *GetPptArtifactExportResultRequest) Validate() error {
	return dara.Validate(s)
}
