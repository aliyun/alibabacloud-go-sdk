// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEcologyOpennessSendVerificationCodeResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetCode(v int32) *EcologyOpennessSendVerificationCodeResponseBody
  GetCode() *int32 
  SetMessage(v string) *EcologyOpennessSendVerificationCodeResponseBody
  GetMessage() *string 
  SetRequestId(v string) *EcologyOpennessSendVerificationCodeResponseBody
  GetRequestId() *string 
  SetResult(v *EcologyOpennessSendVerificationCodeResponseBodyResult) *EcologyOpennessSendVerificationCodeResponseBody
  GetResult() *EcologyOpennessSendVerificationCodeResponseBodyResult 
  SetSuccess(v bool) *EcologyOpennessSendVerificationCodeResponseBody
  GetSuccess() *bool 
}

type EcologyOpennessSendVerificationCodeResponseBody struct {
  // example:
  // 
  // 200
  Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
  // example:
  // 
  // OK
  Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
  // example:
  // 
  // 73C67BD9-175A-1324-8202-9FAABBB3E6FA
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
  Result *EcologyOpennessSendVerificationCodeResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
  Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s EcologyOpennessSendVerificationCodeResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EcologyOpennessSendVerificationCodeResponseBody) GoString() string {
  return s.String()
}

func (s *EcologyOpennessSendVerificationCodeResponseBody) GetCode() *int32  {
  return s.Code
}

func (s *EcologyOpennessSendVerificationCodeResponseBody) GetMessage() *string  {
  return s.Message
}

func (s *EcologyOpennessSendVerificationCodeResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EcologyOpennessSendVerificationCodeResponseBody) GetResult() *EcologyOpennessSendVerificationCodeResponseBodyResult  {
  return s.Result
}

func (s *EcologyOpennessSendVerificationCodeResponseBody) GetSuccess() *bool  {
  return s.Success
}

func (s *EcologyOpennessSendVerificationCodeResponseBody) SetCode(v int32) *EcologyOpennessSendVerificationCodeResponseBody {
  s.Code = &v
  return s
}

func (s *EcologyOpennessSendVerificationCodeResponseBody) SetMessage(v string) *EcologyOpennessSendVerificationCodeResponseBody {
  s.Message = &v
  return s
}

func (s *EcologyOpennessSendVerificationCodeResponseBody) SetRequestId(v string) *EcologyOpennessSendVerificationCodeResponseBody {
  s.RequestId = &v
  return s
}

func (s *EcologyOpennessSendVerificationCodeResponseBody) SetResult(v *EcologyOpennessSendVerificationCodeResponseBodyResult) *EcologyOpennessSendVerificationCodeResponseBody {
  s.Result = v
  return s
}

func (s *EcologyOpennessSendVerificationCodeResponseBody) SetSuccess(v bool) *EcologyOpennessSendVerificationCodeResponseBody {
  s.Success = &v
  return s
}

func (s *EcologyOpennessSendVerificationCodeResponseBody) Validate() error {
  return dara.Validate(s)
}

type EcologyOpennessSendVerificationCodeResponseBodyResult struct {
  // example:
  // 
  // 900
  ExpireIn *int32 `json:"ExpireIn,omitempty" xml:"ExpireIn,omitempty"`
  // example:
  // 
  // 60
  RepeatInterval *int32 `json:"RepeatInterval,omitempty" xml:"RepeatInterval,omitempty"`
}

func (s EcologyOpennessSendVerificationCodeResponseBodyResult) String() string {
  return dara.Prettify(s)
}

func (s EcologyOpennessSendVerificationCodeResponseBodyResult) GoString() string {
  return s.String()
}

func (s *EcologyOpennessSendVerificationCodeResponseBodyResult) GetExpireIn() *int32  {
  return s.ExpireIn
}

func (s *EcologyOpennessSendVerificationCodeResponseBodyResult) GetRepeatInterval() *int32  {
  return s.RepeatInterval
}

func (s *EcologyOpennessSendVerificationCodeResponseBodyResult) SetExpireIn(v int32) *EcologyOpennessSendVerificationCodeResponseBodyResult {
  s.ExpireIn = &v
  return s
}

func (s *EcologyOpennessSendVerificationCodeResponseBodyResult) SetRepeatInterval(v int32) *EcologyOpennessSendVerificationCodeResponseBodyResult {
  s.RepeatInterval = &v
  return s
}

func (s *EcologyOpennessSendVerificationCodeResponseBodyResult) Validate() error {
  return dara.Validate(s)
}

