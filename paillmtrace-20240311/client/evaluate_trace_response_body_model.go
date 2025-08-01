// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEvaluateTraceResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetCode(v string) *EvaluateTraceResponseBody
  GetCode() *string 
  SetEvaluationId(v string) *EvaluateTraceResponseBody
  GetEvaluationId() *string 
  SetMessage(v string) *EvaluateTraceResponseBody
  GetMessage() *string 
  SetRequestId(v string) *EvaluateTraceResponseBody
  GetRequestId() *string 
}

type EvaluateTraceResponseBody struct {
  // The internal error code. This parameter is returned if an exception occurred.
  // 
  // example:
  // 
  // InvalidInputParams
  Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
  // the task ID of the evaluation task to which the trace belongs.
  // 
  // example:
  // 
  // 6000043e103011f0922edec44617e03c
  EvaluationId *string `json:"EvaluationId,omitempty" xml:"EvaluationId,omitempty"`
  // The error message. This parameter is returned if an exception occurred.
  // 
  // example:
  // 
  // eval_request missing dataset id or times
  Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
  // Id of the request
  // 
  // example:
  // 
  // F1AB295E-0D1F-5ECE-9FFA-98ABB4CB5DF5
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EvaluateTraceResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EvaluateTraceResponseBody) GoString() string {
  return s.String()
}

func (s *EvaluateTraceResponseBody) GetCode() *string  {
  return s.Code
}

func (s *EvaluateTraceResponseBody) GetEvaluationId() *string  {
  return s.EvaluationId
}

func (s *EvaluateTraceResponseBody) GetMessage() *string  {
  return s.Message
}

func (s *EvaluateTraceResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EvaluateTraceResponseBody) SetCode(v string) *EvaluateTraceResponseBody {
  s.Code = &v
  return s
}

func (s *EvaluateTraceResponseBody) SetEvaluationId(v string) *EvaluateTraceResponseBody {
  s.EvaluationId = &v
  return s
}

func (s *EvaluateTraceResponseBody) SetMessage(v string) *EvaluateTraceResponseBody {
  s.Message = &v
  return s
}

func (s *EvaluateTraceResponseBody) SetRequestId(v string) *EvaluateTraceResponseBody {
  s.RequestId = &v
  return s
}

func (s *EvaluateTraceResponseBody) Validate() error {
  return dara.Validate(s)
}

