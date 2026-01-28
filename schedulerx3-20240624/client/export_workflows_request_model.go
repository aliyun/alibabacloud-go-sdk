// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportWorkflowsRequest interface {
  dara.Model
  String() string
  GoString() string
  SetAppName(v string) *ExportWorkflowsRequest
  GetAppName() *string 
  SetClusterId(v string) *ExportWorkflowsRequest
  GetClusterId() *string 
  SetWorkflowId(v []*int64) *ExportWorkflowsRequest
  GetWorkflowId() []*int64 
}

type ExportWorkflowsRequest struct {
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
  WorkflowId []*int64 `json:"WorkflowId,omitempty" xml:"WorkflowId,omitempty" type:"Repeated"`
}

func (s ExportWorkflowsRequest) String() string {
  return dara.Prettify(s)
}

func (s ExportWorkflowsRequest) GoString() string {
  return s.String()
}

func (s *ExportWorkflowsRequest) GetAppName() *string  {
  return s.AppName
}

func (s *ExportWorkflowsRequest) GetClusterId() *string  {
  return s.ClusterId
}

func (s *ExportWorkflowsRequest) GetWorkflowId() []*int64  {
  return s.WorkflowId
}

func (s *ExportWorkflowsRequest) SetAppName(v string) *ExportWorkflowsRequest {
  s.AppName = &v
  return s
}

func (s *ExportWorkflowsRequest) SetClusterId(v string) *ExportWorkflowsRequest {
  s.ClusterId = &v
  return s
}

func (s *ExportWorkflowsRequest) SetWorkflowId(v []*int64) *ExportWorkflowsRequest {
  s.WorkflowId = v
  return s
}

func (s *ExportWorkflowsRequest) Validate() error {
  return dara.Validate(s)
}

