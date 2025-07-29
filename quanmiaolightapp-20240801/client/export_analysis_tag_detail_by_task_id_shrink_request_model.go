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
  SetCategory(v string) *ExportAnalysisTagDetailByTaskIdShrinkRequest
  GetCategory() *string 
  SetTaskId(v string) *ExportAnalysisTagDetailByTaskIdShrinkRequest
  GetTaskId() *string 
}

type ExportAnalysisTagDetailByTaskIdShrinkRequest struct {
  CategoriesShrink *string `json:"categories,omitempty" xml:"categories,omitempty"`
  Category *string `json:"category,omitempty" xml:"category,omitempty"`
  // This parameter is required.
  // 
  // example:
  // 
  // a3d1c2ac-f086-4a21-9069-f5631542f5a2
  TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
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

func (s *ExportAnalysisTagDetailByTaskIdShrinkRequest) GetCategory() *string  {
  return s.Category
}

func (s *ExportAnalysisTagDetailByTaskIdShrinkRequest) GetTaskId() *string  {
  return s.TaskId
}

func (s *ExportAnalysisTagDetailByTaskIdShrinkRequest) SetCategoriesShrink(v string) *ExportAnalysisTagDetailByTaskIdShrinkRequest {
  s.CategoriesShrink = &v
  return s
}

func (s *ExportAnalysisTagDetailByTaskIdShrinkRequest) SetCategory(v string) *ExportAnalysisTagDetailByTaskIdShrinkRequest {
  s.Category = &v
  return s
}

func (s *ExportAnalysisTagDetailByTaskIdShrinkRequest) SetTaskId(v string) *ExportAnalysisTagDetailByTaskIdShrinkRequest {
  s.TaskId = &v
  return s
}

func (s *ExportAnalysisTagDetailByTaskIdShrinkRequest) Validate() error {
  return dara.Validate(s)
}

