// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecCrossProjectPipelineRunResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetData(v *ExecCrossProjectPipelineRunResponseBodyData) *ExecCrossProjectPipelineRunResponseBody
  GetData() *ExecCrossProjectPipelineRunResponseBodyData 
  SetRequestId(v string) *ExecCrossProjectPipelineRunResponseBody
  GetRequestId() *string 
  SetSuccess(v bool) *ExecCrossProjectPipelineRunResponseBody
  GetSuccess() *bool 
}

type ExecCrossProjectPipelineRunResponseBody struct {
  // The business response.
  // 
  // example:
  // 
  // {"RequestId":"735894D1-D5E5-50B8-8A6D-041C90A98B23"}
  Data *ExecCrossProjectPipelineRunResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
  // The request ID, which is used to locate and troubleshoot this API call.
  // 
  // example:
  // 
  // 735894D1-D5E5-50B8-8A6D-041C90A98B23
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
  // Indicates whether the request was successful.
  // 
  // example:
  // 
  // true
  Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ExecCrossProjectPipelineRunResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ExecCrossProjectPipelineRunResponseBody) GoString() string {
  return s.String()
}

func (s *ExecCrossProjectPipelineRunResponseBody) GetData() *ExecCrossProjectPipelineRunResponseBodyData  {
  return s.Data
}

func (s *ExecCrossProjectPipelineRunResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *ExecCrossProjectPipelineRunResponseBody) GetSuccess() *bool  {
  return s.Success
}

func (s *ExecCrossProjectPipelineRunResponseBody) SetData(v *ExecCrossProjectPipelineRunResponseBodyData) *ExecCrossProjectPipelineRunResponseBody {
  s.Data = v
  return s
}

func (s *ExecCrossProjectPipelineRunResponseBody) SetRequestId(v string) *ExecCrossProjectPipelineRunResponseBody {
  s.RequestId = &v
  return s
}

func (s *ExecCrossProjectPipelineRunResponseBody) SetSuccess(v bool) *ExecCrossProjectPipelineRunResponseBody {
  s.Success = &v
  return s
}

func (s *ExecCrossProjectPipelineRunResponseBody) Validate() error {
  if s.Data != nil {
    if err := s.Data.Validate(); err != nil {
      return err
    }
  }
  return nil
}

type ExecCrossProjectPipelineRunResponseBodyData struct {
  // The request ID.
  // 
  // example:
  // 
  // 735894D1-D5E5-50B8-8A6D-041C90A98B23
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ExecCrossProjectPipelineRunResponseBodyData) String() string {
  return dara.Prettify(s)
}

func (s ExecCrossProjectPipelineRunResponseBodyData) GoString() string {
  return s.String()
}

func (s *ExecCrossProjectPipelineRunResponseBodyData) GetRequestId() *string  {
  return s.RequestId
}

func (s *ExecCrossProjectPipelineRunResponseBodyData) SetRequestId(v string) *ExecCrossProjectPipelineRunResponseBodyData {
  s.RequestId = &v
  return s
}

func (s *ExecCrossProjectPipelineRunResponseBodyData) Validate() error {
  return dara.Validate(s)
}

