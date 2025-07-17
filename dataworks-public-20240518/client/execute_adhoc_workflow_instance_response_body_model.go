// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteAdhocWorkflowInstanceResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetRequestId(v string) *ExecuteAdhocWorkflowInstanceResponseBody
  GetRequestId() *string 
  SetWorkflowInstanceId(v int64) *ExecuteAdhocWorkflowInstanceResponseBody
  GetWorkflowInstanceId() *int64 
}

type ExecuteAdhocWorkflowInstanceResponseBody struct {
  // The request ID.
  // 
  // example:
  // 
  // 22C97E95-F023-56B5-8852-B1A77A17XXXX
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
  // The workflow instance ID.
  // 
  // example:
  // 
  // 1234
  WorkflowInstanceId *int64 `json:"WorkflowInstanceId,omitempty" xml:"WorkflowInstanceId,omitempty"`
}

func (s ExecuteAdhocWorkflowInstanceResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ExecuteAdhocWorkflowInstanceResponseBody) GoString() string {
  return s.String()
}

func (s *ExecuteAdhocWorkflowInstanceResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *ExecuteAdhocWorkflowInstanceResponseBody) GetWorkflowInstanceId() *int64  {
  return s.WorkflowInstanceId
}

func (s *ExecuteAdhocWorkflowInstanceResponseBody) SetRequestId(v string) *ExecuteAdhocWorkflowInstanceResponseBody {
  s.RequestId = &v
  return s
}

func (s *ExecuteAdhocWorkflowInstanceResponseBody) SetWorkflowInstanceId(v int64) *ExecuteAdhocWorkflowInstanceResponseBody {
  s.WorkflowInstanceId = &v
  return s
}

func (s *ExecuteAdhocWorkflowInstanceResponseBody) Validate() error {
  return dara.Validate(s)
}

