// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExistMcubeRsaKeyResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetCheckRsaKeyResult(v *ExistMcubeRsaKeyResponseBodyCheckRsaKeyResult) *ExistMcubeRsaKeyResponseBody
  GetCheckRsaKeyResult() *ExistMcubeRsaKeyResponseBodyCheckRsaKeyResult 
  SetRequestId(v string) *ExistMcubeRsaKeyResponseBody
  GetRequestId() *string 
  SetResultCode(v string) *ExistMcubeRsaKeyResponseBody
  GetResultCode() *string 
  SetResultMessage(v string) *ExistMcubeRsaKeyResponseBody
  GetResultMessage() *string 
}

type ExistMcubeRsaKeyResponseBody struct {
  CheckRsaKeyResult *ExistMcubeRsaKeyResponseBodyCheckRsaKeyResult `json:"CheckRsaKeyResult,omitempty" xml:"CheckRsaKeyResult,omitempty" type:"Struct"`
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
  ResultCode *string `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
  ResultMessage *string `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
}

func (s ExistMcubeRsaKeyResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ExistMcubeRsaKeyResponseBody) GoString() string {
  return s.String()
}

func (s *ExistMcubeRsaKeyResponseBody) GetCheckRsaKeyResult() *ExistMcubeRsaKeyResponseBodyCheckRsaKeyResult  {
  return s.CheckRsaKeyResult
}

func (s *ExistMcubeRsaKeyResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *ExistMcubeRsaKeyResponseBody) GetResultCode() *string  {
  return s.ResultCode
}

func (s *ExistMcubeRsaKeyResponseBody) GetResultMessage() *string  {
  return s.ResultMessage
}

func (s *ExistMcubeRsaKeyResponseBody) SetCheckRsaKeyResult(v *ExistMcubeRsaKeyResponseBodyCheckRsaKeyResult) *ExistMcubeRsaKeyResponseBody {
  s.CheckRsaKeyResult = v
  return s
}

func (s *ExistMcubeRsaKeyResponseBody) SetRequestId(v string) *ExistMcubeRsaKeyResponseBody {
  s.RequestId = &v
  return s
}

func (s *ExistMcubeRsaKeyResponseBody) SetResultCode(v string) *ExistMcubeRsaKeyResponseBody {
  s.ResultCode = &v
  return s
}

func (s *ExistMcubeRsaKeyResponseBody) SetResultMessage(v string) *ExistMcubeRsaKeyResponseBody {
  s.ResultMessage = &v
  return s
}

func (s *ExistMcubeRsaKeyResponseBody) Validate() error {
  return dara.Validate(s)
}

type ExistMcubeRsaKeyResponseBodyCheckRsaKeyResult struct {
  Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
  ResultMsg *string `json:"ResultMsg,omitempty" xml:"ResultMsg,omitempty"`
  Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ExistMcubeRsaKeyResponseBodyCheckRsaKeyResult) String() string {
  return dara.Prettify(s)
}

func (s ExistMcubeRsaKeyResponseBodyCheckRsaKeyResult) GoString() string {
  return s.String()
}

func (s *ExistMcubeRsaKeyResponseBodyCheckRsaKeyResult) GetData() *string  {
  return s.Data
}

func (s *ExistMcubeRsaKeyResponseBodyCheckRsaKeyResult) GetResultMsg() *string  {
  return s.ResultMsg
}

func (s *ExistMcubeRsaKeyResponseBodyCheckRsaKeyResult) GetSuccess() *bool  {
  return s.Success
}

func (s *ExistMcubeRsaKeyResponseBodyCheckRsaKeyResult) SetData(v string) *ExistMcubeRsaKeyResponseBodyCheckRsaKeyResult {
  s.Data = &v
  return s
}

func (s *ExistMcubeRsaKeyResponseBodyCheckRsaKeyResult) SetResultMsg(v string) *ExistMcubeRsaKeyResponseBodyCheckRsaKeyResult {
  s.ResultMsg = &v
  return s
}

func (s *ExistMcubeRsaKeyResponseBodyCheckRsaKeyResult) SetSuccess(v bool) *ExistMcubeRsaKeyResponseBodyCheckRsaKeyResult {
  s.Success = &v
  return s
}

func (s *ExistMcubeRsaKeyResponseBodyCheckRsaKeyResult) Validate() error {
  return dara.Validate(s)
}

