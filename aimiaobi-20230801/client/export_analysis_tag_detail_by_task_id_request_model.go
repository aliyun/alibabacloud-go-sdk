// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportAnalysisTagDetailByTaskIdRequest interface {
  dara.Model
  String() string
  GoString() string
  SetCategories(v []*string) *ExportAnalysisTagDetailByTaskIdRequest
  GetCategories() []*string 
  SetTaskId(v string) *ExportAnalysisTagDetailByTaskIdRequest
  GetTaskId() *string 
  SetWorkspaceId(v string) *ExportAnalysisTagDetailByTaskIdRequest
  GetWorkspaceId() *string 
}

type ExportAnalysisTagDetailByTaskIdRequest struct {
  Categories []*string `json:"Categories,omitempty" xml:"Categories,omitempty" type:"Repeated"`
  // This parameter is required.
  // 
  // example:
  // 
  // xxx
  TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
  // This parameter is required.
  // 
  // example:
  // 
  // xxxx
  WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ExportAnalysisTagDetailByTaskIdRequest) String() string {
  return dara.Prettify(s)
}

func (s ExportAnalysisTagDetailByTaskIdRequest) GoString() string {
  return s.String()
}

func (s *ExportAnalysisTagDetailByTaskIdRequest) GetCategories() []*string  {
  return s.Categories
}

func (s *ExportAnalysisTagDetailByTaskIdRequest) GetTaskId() *string  {
  return s.TaskId
}

func (s *ExportAnalysisTagDetailByTaskIdRequest) GetWorkspaceId() *string  {
  return s.WorkspaceId
}

func (s *ExportAnalysisTagDetailByTaskIdRequest) SetCategories(v []*string) *ExportAnalysisTagDetailByTaskIdRequest {
  s.Categories = v
  return s
}

func (s *ExportAnalysisTagDetailByTaskIdRequest) SetTaskId(v string) *ExportAnalysisTagDetailByTaskIdRequest {
  s.TaskId = &v
  return s
}

func (s *ExportAnalysisTagDetailByTaskIdRequest) SetWorkspaceId(v string) *ExportAnalysisTagDetailByTaskIdRequest {
  s.WorkspaceId = &v
  return s
}

func (s *ExportAnalysisTagDetailByTaskIdRequest) Validate() error {
  return dara.Validate(s)
}

