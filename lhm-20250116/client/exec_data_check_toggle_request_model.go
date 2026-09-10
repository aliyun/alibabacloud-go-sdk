// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecDataCheckToggleRequest interface {
  dara.Model
  String() string
  GoString() string
  SetParams(v []*ExecDataCheckToggleRequestParams) *ExecDataCheckToggleRequest
  GetParams() []*ExecDataCheckToggleRequestParams 
}

type ExecDataCheckToggleRequest struct {
  // The task scheduling parameter list. Each item must contain id, lastBatchId, and isScheduled.
  // 
  // This parameter is required.
  Params []*ExecDataCheckToggleRequestParams `json:"params,omitempty" xml:"params,omitempty" type:"Repeated"`
}

func (s ExecDataCheckToggleRequest) String() string {
  return dara.Prettify(s)
}

func (s ExecDataCheckToggleRequest) GoString() string {
  return s.String()
}

func (s *ExecDataCheckToggleRequest) GetParams() []*ExecDataCheckToggleRequestParams  {
  return s.Params
}

func (s *ExecDataCheckToggleRequest) SetParams(v []*ExecDataCheckToggleRequestParams) *ExecDataCheckToggleRequest {
  s.Params = v
  return s
}

func (s *ExecDataCheckToggleRequest) Validate() error {
  if s.Params != nil {
    for _, item := range s.Params {
      if item != nil {
        if err := item.Validate(); err != nil {
          return err
        }
      }
    }
  }
  return nil
}

type ExecDataCheckToggleRequestParams struct {
  // The task ID.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // 10001
  Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
  // Specifies whether to enable scheduling. Valid values:
  // 
  // - 0: Disabled.
  // 
  // - 1: Enabled.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // 0
  IsScheduled *int32 `json:"isScheduled,omitempty" xml:"isScheduled,omitempty"`
  // The most recent batch number.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // 20001
  LastBatchId *int64 `json:"lastBatchId,omitempty" xml:"lastBatchId,omitempty"`
}

func (s ExecDataCheckToggleRequestParams) String() string {
  return dara.Prettify(s)
}

func (s ExecDataCheckToggleRequestParams) GoString() string {
  return s.String()
}

func (s *ExecDataCheckToggleRequestParams) GetId() *int64  {
  return s.Id
}

func (s *ExecDataCheckToggleRequestParams) GetIsScheduled() *int32  {
  return s.IsScheduled
}

func (s *ExecDataCheckToggleRequestParams) GetLastBatchId() *int64  {
  return s.LastBatchId
}

func (s *ExecDataCheckToggleRequestParams) SetId(v int64) *ExecDataCheckToggleRequestParams {
  s.Id = &v
  return s
}

func (s *ExecDataCheckToggleRequestParams) SetIsScheduled(v int32) *ExecDataCheckToggleRequestParams {
  s.IsScheduled = &v
  return s
}

func (s *ExecDataCheckToggleRequestParams) SetLastBatchId(v int64) *ExecDataCheckToggleRequestParams {
  s.LastBatchId = &v
  return s
}

func (s *ExecDataCheckToggleRequestParams) Validate() error {
  return dara.Validate(s)
}

