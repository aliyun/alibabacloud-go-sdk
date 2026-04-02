// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEkycVerifyV2ResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetCode(v string) *EkycVerifyV2ResponseBody
  GetCode() *string 
  SetMessage(v string) *EkycVerifyV2ResponseBody
  GetMessage() *string 
  SetRequestId(v string) *EkycVerifyV2ResponseBody
  GetRequestId() *string 
  SetResult(v *EkycVerifyV2ResponseBodyResult) *EkycVerifyV2ResponseBody
  GetResult() *EkycVerifyV2ResponseBodyResult 
}

type EkycVerifyV2ResponseBody struct {
  // example:
  // 
  // Success
  Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
  // example:
  // 
  // success
  Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
  // Id of the request
  // 
  // example:
  // 
  // 4EB356FE-BB6A-5DCC-B4C5-E8051787EBA1
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
  Result *EkycVerifyV2ResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s EkycVerifyV2ResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EkycVerifyV2ResponseBody) GoString() string {
  return s.String()
}

func (s *EkycVerifyV2ResponseBody) GetCode() *string  {
  return s.Code
}

func (s *EkycVerifyV2ResponseBody) GetMessage() *string  {
  return s.Message
}

func (s *EkycVerifyV2ResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EkycVerifyV2ResponseBody) GetResult() *EkycVerifyV2ResponseBodyResult  {
  return s.Result
}

func (s *EkycVerifyV2ResponseBody) SetCode(v string) *EkycVerifyV2ResponseBody {
  s.Code = &v
  return s
}

func (s *EkycVerifyV2ResponseBody) SetMessage(v string) *EkycVerifyV2ResponseBody {
  s.Message = &v
  return s
}

func (s *EkycVerifyV2ResponseBody) SetRequestId(v string) *EkycVerifyV2ResponseBody {
  s.RequestId = &v
  return s
}

func (s *EkycVerifyV2ResponseBody) SetResult(v *EkycVerifyV2ResponseBodyResult) *EkycVerifyV2ResponseBody {
  s.Result = v
  return s
}

func (s *EkycVerifyV2ResponseBody) Validate() error {
  if s.Result != nil {
    if err := s.Result.Validate(); err != nil {
      return err
    }
  }
  return nil
}

type EkycVerifyV2ResponseBodyResult struct {
  // example:
  // 
  // {
  // 
  // "faceAttack": "N",
  // 
  // "faceComparisonScore": 52.57,
  // 
  // "facePassed": "N",
  // 
  // "authorityComparisonScore": 80.39
  // 
  // }
  ExtFaceInfo *string `json:"ExtFaceInfo,omitempty" xml:"ExtFaceInfo,omitempty"`
  ExtIdInfo *string `json:"ExtIdInfo,omitempty" xml:"ExtIdInfo,omitempty"`
  // example:
  // 
  // Y
  Passed *string `json:"Passed,omitempty" xml:"Passed,omitempty"`
  // example:
  // 
  // 205
  SubCode *string `json:"SubCode,omitempty" xml:"SubCode,omitempty"`
  // example:
  // 
  // 4ab0b***cbde97
  TransactionId *string `json:"TransactionId,omitempty" xml:"TransactionId,omitempty"`
}

func (s EkycVerifyV2ResponseBodyResult) String() string {
  return dara.Prettify(s)
}

func (s EkycVerifyV2ResponseBodyResult) GoString() string {
  return s.String()
}

func (s *EkycVerifyV2ResponseBodyResult) GetExtFaceInfo() *string  {
  return s.ExtFaceInfo
}

func (s *EkycVerifyV2ResponseBodyResult) GetExtIdInfo() *string  {
  return s.ExtIdInfo
}

func (s *EkycVerifyV2ResponseBodyResult) GetPassed() *string  {
  return s.Passed
}

func (s *EkycVerifyV2ResponseBodyResult) GetSubCode() *string  {
  return s.SubCode
}

func (s *EkycVerifyV2ResponseBodyResult) GetTransactionId() *string  {
  return s.TransactionId
}

func (s *EkycVerifyV2ResponseBodyResult) SetExtFaceInfo(v string) *EkycVerifyV2ResponseBodyResult {
  s.ExtFaceInfo = &v
  return s
}

func (s *EkycVerifyV2ResponseBodyResult) SetExtIdInfo(v string) *EkycVerifyV2ResponseBodyResult {
  s.ExtIdInfo = &v
  return s
}

func (s *EkycVerifyV2ResponseBodyResult) SetPassed(v string) *EkycVerifyV2ResponseBodyResult {
  s.Passed = &v
  return s
}

func (s *EkycVerifyV2ResponseBodyResult) SetSubCode(v string) *EkycVerifyV2ResponseBodyResult {
  s.SubCode = &v
  return s
}

func (s *EkycVerifyV2ResponseBodyResult) SetTransactionId(v string) *EkycVerifyV2ResponseBodyResult {
  s.TransactionId = &v
  return s
}

func (s *EkycVerifyV2ResponseBodyResult) Validate() error {
  return dara.Validate(s)
}

