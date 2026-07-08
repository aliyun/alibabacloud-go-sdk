// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportAnalysisTagDetailByTaskIdShrinkRequest interface {
  dara.Model
  String() string
  GoString() string
  SetCategoriesShrink(v string) *ExportAnalysisTagDetailByTaskIdShrinkRequest
  GetCategoriesShrink() *string 
  SetTaskId(v string) *ExportAnalysisTagDetailByTaskIdShrinkRequest
  GetTaskId() *string 
  SetWorkspaceId(v string) *ExportAnalysisTagDetailByTaskIdShrinkRequest
  GetWorkspaceId() *string 
}

type ExportAnalysisTagDetailByTaskIdShrinkRequest struct {
  // Category filter list.
  CategoriesShrink *string `json:"Categories,omitempty" xml:"Categories,omitempty"`
  // Unique task ID.
  // 
  // > By default, do not specify TaskId. The system automatically generates it. If subsequent tasks use the same TaskId, they belong to the same conversation group.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // xxx
  TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
  // Unique identifier for Alibaba Cloud Model Studio workspace. For more information, see [Get Workspace ID](https://help.aliyun.com/document_detail/2782167.html).
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // xxxx
  WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ExportAnalysisTagDetailByTaskIdShrinkRequest) String() string {
  return dara.Prettify(s)
}

func (s ExportAnalysisTagDetailByTaskIdShrinkRequest) GoString() string {
  return s.String()
}

func (s *ExportAnalysisTagDetailByTaskIdShrinkRequest) GetCategoriesShrink() *string  {
  return s.CategoriesShrink
}

func (s *ExportAnalysisTagDetailByTaskIdShrinkRequest) GetTaskId() *string  {
  return s.TaskId
}

func (s *ExportAnalysisTagDetailByTaskIdShrinkRequest) GetWorkspaceId() *string  {
  return s.WorkspaceId
}

func (s *ExportAnalysisTagDetailByTaskIdShrinkRequest) SetCategoriesShrink(v string) *ExportAnalysisTagDetailByTaskIdShrinkRequest {
  s.CategoriesShrink = &v
  return s
}

func (s *ExportAnalysisTagDetailByTaskIdShrinkRequest) SetTaskId(v string) *ExportAnalysisTagDetailByTaskIdShrinkRequest {
  s.TaskId = &v
  return s
}

func (s *ExportAnalysisTagDetailByTaskIdShrinkRequest) SetWorkspaceId(v string) *ExportAnalysisTagDetailByTaskIdShrinkRequest {
  s.WorkspaceId = &v
  return s
}

func (s *ExportAnalysisTagDetailByTaskIdShrinkRequest) Validate() error {
  return dara.Validate(s)
}

