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
  SetCategory(v string) *ExportAnalysisTagDetailByTaskIdRequest
  GetCategory() *string 
  SetTaskId(v string) *ExportAnalysisTagDetailByTaskIdRequest
  GetTaskId() *string 
}

type ExportAnalysisTagDetailByTaskIdRequest struct {
  Categories []*string `json:"categories,omitempty" xml:"categories,omitempty" type:"Repeated"`
  Category *string `json:"category,omitempty" xml:"category,omitempty"`
  // This parameter is required.
  // 
  // example:
  // 
  // a3d1c2ac-f086-4a21-9069-f5631542f5a2
  TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
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

func (s *ExportAnalysisTagDetailByTaskIdRequest) GetCategory() *string  {
  return s.Category
}

func (s *ExportAnalysisTagDetailByTaskIdRequest) GetTaskId() *string  {
  return s.TaskId
}

func (s *ExportAnalysisTagDetailByTaskIdRequest) SetCategories(v []*string) *ExportAnalysisTagDetailByTaskIdRequest {
  s.Categories = v
  return s
}

func (s *ExportAnalysisTagDetailByTaskIdRequest) SetCategory(v string) *ExportAnalysisTagDetailByTaskIdRequest {
  s.Category = &v
  return s
}

func (s *ExportAnalysisTagDetailByTaskIdRequest) SetTaskId(v string) *ExportAnalysisTagDetailByTaskIdRequest {
  s.TaskId = &v
  return s
}

func (s *ExportAnalysisTagDetailByTaskIdRequest) Validate() error {
  return dara.Validate(s)
}

