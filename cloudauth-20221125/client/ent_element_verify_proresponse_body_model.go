// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEntElementVerifyPROResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetCode(v string) *EntElementVerifyPROResponseBody
  GetCode() *string 
  SetMessage(v string) *EntElementVerifyPROResponseBody
  GetMessage() *string 
  SetRequestId(v string) *EntElementVerifyPROResponseBody
  GetRequestId() *string 
  SetResult(v *EntElementVerifyPROResponseBodyResult) *EntElementVerifyPROResponseBody
  GetResult() *EntElementVerifyPROResponseBodyResult 
}

type EntElementVerifyPROResponseBody struct {
  // example:
  // 
  // 200
  Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
  // example:
  // 
  // success
  Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
  // Id of the request
  // 
  // example:
  // 
  // 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
  Result *EntElementVerifyPROResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s EntElementVerifyPROResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EntElementVerifyPROResponseBody) GoString() string {
  return s.String()
}

func (s *EntElementVerifyPROResponseBody) GetCode() *string  {
  return s.Code
}

func (s *EntElementVerifyPROResponseBody) GetMessage() *string  {
  return s.Message
}

func (s *EntElementVerifyPROResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EntElementVerifyPROResponseBody) GetResult() *EntElementVerifyPROResponseBodyResult  {
  return s.Result
}

func (s *EntElementVerifyPROResponseBody) SetCode(v string) *EntElementVerifyPROResponseBody {
  s.Code = &v
  return s
}

func (s *EntElementVerifyPROResponseBody) SetMessage(v string) *EntElementVerifyPROResponseBody {
  s.Message = &v
  return s
}

func (s *EntElementVerifyPROResponseBody) SetRequestId(v string) *EntElementVerifyPROResponseBody {
  s.RequestId = &v
  return s
}

func (s *EntElementVerifyPROResponseBody) SetResult(v *EntElementVerifyPROResponseBodyResult) *EntElementVerifyPROResponseBody {
  s.Result = v
  return s
}

func (s *EntElementVerifyPROResponseBody) Validate() error {
  if s.Result != nil {
    if err := s.Result.Validate(); err != nil {
      return err
    }
  }
  return nil
}

type EntElementVerifyPROResponseBodyResult struct {
  // example:
  // 
  // 1
  BizCode *string `json:"BizCode,omitempty" xml:"BizCode,omitempty"`
  // example:
  // 
  // 2018-09-25/9999-09-09
  OpenTime *string `json:"OpenTime,omitempty" xml:"OpenTime,omitempty"`
  // example:
  // 
  // LegalPersonNameFlag,LegalPersonCertNoFlag
  ReasonDetail *string `json:"ReasonDetail,omitempty" xml:"ReasonDetail,omitempty"`
  // example:
  // 
  // 1
  Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s EntElementVerifyPROResponseBodyResult) String() string {
  return dara.Prettify(s)
}

func (s EntElementVerifyPROResponseBodyResult) GoString() string {
  return s.String()
}

func (s *EntElementVerifyPROResponseBodyResult) GetBizCode() *string  {
  return s.BizCode
}

func (s *EntElementVerifyPROResponseBodyResult) GetOpenTime() *string  {
  return s.OpenTime
}

func (s *EntElementVerifyPROResponseBodyResult) GetReasonDetail() *string  {
  return s.ReasonDetail
}

func (s *EntElementVerifyPROResponseBodyResult) GetStatus() *string  {
  return s.Status
}

func (s *EntElementVerifyPROResponseBodyResult) SetBizCode(v string) *EntElementVerifyPROResponseBodyResult {
  s.BizCode = &v
  return s
}

func (s *EntElementVerifyPROResponseBodyResult) SetOpenTime(v string) *EntElementVerifyPROResponseBodyResult {
  s.OpenTime = &v
  return s
}

func (s *EntElementVerifyPROResponseBodyResult) SetReasonDetail(v string) *EntElementVerifyPROResponseBodyResult {
  s.ReasonDetail = &v
  return s
}

func (s *EntElementVerifyPROResponseBodyResult) SetStatus(v string) *EntElementVerifyPROResponseBodyResult {
  s.Status = &v
  return s
}

func (s *EntElementVerifyPROResponseBodyResult) Validate() error {
  return dara.Validate(s)
}

