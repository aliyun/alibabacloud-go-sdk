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
  // The name of the application.
  // 
  // example:
  // 
  // test-app
  AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
  // The ID of the cluster where the Workflow is located.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // xxljob-b6ec1xxxx
  ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
  // An array of Workflow IDs to export.
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

