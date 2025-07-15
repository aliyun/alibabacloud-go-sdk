// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportAuditContentResultRequest interface {
  dara.Model
  String() string
  GoString() string
  SetTaskId(v string) *ExportAuditContentResultRequest
  GetTaskId() *string 
  SetWorkspaceId(v string) *ExportAuditContentResultRequest
  GetWorkspaceId() *string 
}

type ExportAuditContentResultRequest struct {
  // example:
  // 
  // 7AA2AE16-D873-5C5F-9708-15396C382EB1
  TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
  // example:
  // 
  // xxxx
  WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ExportAuditContentResultRequest) String() string {
  return dara.Prettify(s)
}

func (s ExportAuditContentResultRequest) GoString() string {
  return s.String()
}

func (s *ExportAuditContentResultRequest) GetTaskId() *string  {
  return s.TaskId
}

func (s *ExportAuditContentResultRequest) GetWorkspaceId() *string  {
  return s.WorkspaceId
}

func (s *ExportAuditContentResultRequest) SetTaskId(v string) *ExportAuditContentResultRequest {
  s.TaskId = &v
  return s
}

func (s *ExportAuditContentResultRequest) SetWorkspaceId(v string) *ExportAuditContentResultRequest {
  s.WorkspaceId = &v
  return s
}

func (s *ExportAuditContentResultRequest) Validate() error {
  return dara.Validate(s)
}

