// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportWorkflowsShrinkRequest interface {
  dara.Model
  String() string
  GoString() string
  SetAppName(v string) *ExportWorkflowsShrinkRequest
  GetAppName() *string 
  SetClusterId(v string) *ExportWorkflowsShrinkRequest
  GetClusterId() *string 
  SetWorkflowIdShrink(v string) *ExportWorkflowsShrinkRequest
  GetWorkflowIdShrink() *string 
}

type ExportWorkflowsShrinkRequest struct {
  // example:
  // 
  // test-app
  AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
  // A short description of struct
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // xxljob-b6ec1xxxx
  ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
  WorkflowIdShrink *string `json:"WorkflowId,omitempty" xml:"WorkflowId,omitempty"`
}

func (s ExportWorkflowsShrinkRequest) String() string {
  return dara.Prettify(s)
}

func (s ExportWorkflowsShrinkRequest) GoString() string {
  return s.String()
}

func (s *ExportWorkflowsShrinkRequest) GetAppName() *string  {
  return s.AppName
}

func (s *ExportWorkflowsShrinkRequest) GetClusterId() *string  {
  return s.ClusterId
}

func (s *ExportWorkflowsShrinkRequest) GetWorkflowIdShrink() *string  {
  return s.WorkflowIdShrink
}

func (s *ExportWorkflowsShrinkRequest) SetAppName(v string) *ExportWorkflowsShrinkRequest {
  s.AppName = &v
  return s
}

func (s *ExportWorkflowsShrinkRequest) SetClusterId(v string) *ExportWorkflowsShrinkRequest {
  s.ClusterId = &v
  return s
}

func (s *ExportWorkflowsShrinkRequest) SetWorkflowIdShrink(v string) *ExportWorkflowsShrinkRequest {
  s.WorkflowIdShrink = &v
  return s
}

func (s *ExportWorkflowsShrinkRequest) Validate() error {
  return dara.Validate(s)
}

