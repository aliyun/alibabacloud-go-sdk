// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportCustomSourceAnalysisTaskRequest interface {
  dara.Model
  String() string
  GoString() string
  SetTaskId(v string) *ExportCustomSourceAnalysisTaskRequest
  GetTaskId() *string 
  SetWorkspaceId(v string) *ExportCustomSourceAnalysisTaskRequest
  GetWorkspaceId() *string 
}

type ExportCustomSourceAnalysisTaskRequest struct {
  // This parameter is required.
  // 
  // example:
  // 
  // c9f226b02cca4f42a84c5e955c39dfd2
  TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
  // This parameter is required.
  // 
  // example:
  // 
  // xxxxx_p_efm
  WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ExportCustomSourceAnalysisTaskRequest) String() string {
  return dara.Prettify(s)
}

func (s ExportCustomSourceAnalysisTaskRequest) GoString() string {
  return s.String()
}

func (s *ExportCustomSourceAnalysisTaskRequest) GetTaskId() *string  {
  return s.TaskId
}

func (s *ExportCustomSourceAnalysisTaskRequest) GetWorkspaceId() *string  {
  return s.WorkspaceId
}

func (s *ExportCustomSourceAnalysisTaskRequest) SetTaskId(v string) *ExportCustomSourceAnalysisTaskRequest {
  s.TaskId = &v
  return s
}

func (s *ExportCustomSourceAnalysisTaskRequest) SetWorkspaceId(v string) *ExportCustomSourceAnalysisTaskRequest {
  s.WorkspaceId = &v
  return s
}

func (s *ExportCustomSourceAnalysisTaskRequest) Validate() error {
  return dara.Validate(s)
}

