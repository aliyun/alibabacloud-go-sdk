// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecDataCheckSaveTaskRequest interface {
  dara.Model
  String() string
  GoString() string
  SetCheckGlobalParams(v string) *ExecDataCheckSaveTaskRequest
  GetCheckGlobalParams() *string 
  SetFullTableCount(v int32) *ExecDataCheckSaveTaskRequest
  GetFullTableCount() *int32 
  SetSourceGlobalParams(v string) *ExecDataCheckSaveTaskRequest
  GetSourceGlobalParams() *string 
  SetStartImmediately(v int32) *ExecDataCheckSaveTaskRequest
  GetStartImmediately() *int32 
  SetTargetGlobalParams(v string) *ExecDataCheckSaveTaskRequest
  GetTargetGlobalParams() *string 
  SetTaskId(v int64) *ExecDataCheckSaveTaskRequest
  GetTaskId() *int64 
  SetTotalCountThreshold(v float32) *ExecDataCheckSaveTaskRequest
  GetTotalCountThreshold() *float32 
}

type ExecDataCheckSaveTaskRequest struct {
  // The global parameters for the validation phase. Separate multiple parameters with a line feed (`
  // 
  // `).
  // 
  // example:
  // 
  // {}
  CheckGlobalParams *string `json:"checkGlobalParams,omitempty" xml:"checkGlobalParams,omitempty"`
  // Specifies whether to perform full-table validation. Valid values:
  // 
  // - 0: Partition-level validation. This is the default value.
  // 
  // - 1: Full-table validation.
  // 
  // example:
  // 
  // 0
  FullTableCount *int32 `json:"fullTableCount,omitempty" xml:"fullTableCount,omitempty"`
  // The global parameters for the source. Separate multiple parameters with a line feed (`
  // 
  // `).
  // 
  // example:
  // 
  // {}
  SourceGlobalParams *string `json:"sourceGlobalParams,omitempty" xml:"sourceGlobalParams,omitempty"`
  // Specifies whether to execute immediately after saving. Valid values:
  // 
  // - 0: No. This is the default value.
  // 
  // - 1: Yes.
  // 
  // example:
  // 
  // 0
  StartImmediately *int32 `json:"startImmediately,omitempty" xml:"startImmediately,omitempty"`
  // The global parameters for the target. Separate multiple parameters with a line feed (`
  // 
  // `).
  // 
  // example:
  // 
  // {}
  TargetGlobalParams *string `json:"targetGlobalParams,omitempty" xml:"targetGlobalParams,omitempty"`
  // The ID of the validation task.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // 10001
  TaskId *int64 `json:"taskId,omitempty" xml:"taskId,omitempty"`
  // The total data volume comparison threshold, used to determine whether the data volume difference between the source and target is within an acceptable range.
  // 
  // example:
  // 
  // 0.5
  TotalCountThreshold *float32 `json:"totalCountThreshold,omitempty" xml:"totalCountThreshold,omitempty"`
}

func (s ExecDataCheckSaveTaskRequest) String() string {
  return dara.Prettify(s)
}

func (s ExecDataCheckSaveTaskRequest) GoString() string {
  return s.String()
}

func (s *ExecDataCheckSaveTaskRequest) GetCheckGlobalParams() *string  {
  return s.CheckGlobalParams
}

func (s *ExecDataCheckSaveTaskRequest) GetFullTableCount() *int32  {
  return s.FullTableCount
}

func (s *ExecDataCheckSaveTaskRequest) GetSourceGlobalParams() *string  {
  return s.SourceGlobalParams
}

func (s *ExecDataCheckSaveTaskRequest) GetStartImmediately() *int32  {
  return s.StartImmediately
}

func (s *ExecDataCheckSaveTaskRequest) GetTargetGlobalParams() *string  {
  return s.TargetGlobalParams
}

func (s *ExecDataCheckSaveTaskRequest) GetTaskId() *int64  {
  return s.TaskId
}

func (s *ExecDataCheckSaveTaskRequest) GetTotalCountThreshold() *float32  {
  return s.TotalCountThreshold
}

func (s *ExecDataCheckSaveTaskRequest) SetCheckGlobalParams(v string) *ExecDataCheckSaveTaskRequest {
  s.CheckGlobalParams = &v
  return s
}

func (s *ExecDataCheckSaveTaskRequest) SetFullTableCount(v int32) *ExecDataCheckSaveTaskRequest {
  s.FullTableCount = &v
  return s
}

func (s *ExecDataCheckSaveTaskRequest) SetSourceGlobalParams(v string) *ExecDataCheckSaveTaskRequest {
  s.SourceGlobalParams = &v
  return s
}

func (s *ExecDataCheckSaveTaskRequest) SetStartImmediately(v int32) *ExecDataCheckSaveTaskRequest {
  s.StartImmediately = &v
  return s
}

func (s *ExecDataCheckSaveTaskRequest) SetTargetGlobalParams(v string) *ExecDataCheckSaveTaskRequest {
  s.TargetGlobalParams = &v
  return s
}

func (s *ExecDataCheckSaveTaskRequest) SetTaskId(v int64) *ExecDataCheckSaveTaskRequest {
  s.TaskId = &v
  return s
}

func (s *ExecDataCheckSaveTaskRequest) SetTotalCountThreshold(v float32) *ExecDataCheckSaveTaskRequest {
  s.TotalCountThreshold = &v
  return s
}

func (s *ExecDataCheckSaveTaskRequest) Validate() error {
  return dara.Validate(s)
}

