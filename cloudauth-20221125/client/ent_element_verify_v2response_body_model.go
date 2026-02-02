// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEntElementVerifyV2ResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetCode(v string) *EntElementVerifyV2ResponseBody
  GetCode() *string 
  SetMessage(v string) *EntElementVerifyV2ResponseBody
  GetMessage() *string 
  SetRequestId(v string) *EntElementVerifyV2ResponseBody
  GetRequestId() *string 
  SetResult(v *EntElementVerifyV2ResponseBodyResult) *EntElementVerifyV2ResponseBody
  GetResult() *EntElementVerifyV2ResponseBodyResult 
}

type EntElementVerifyV2ResponseBody struct {
  // Return code
  // 
  // example:
  // 
  // Success
  Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
  // Error message
  // 
  // example:
  // 
  // success
  Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
  // Id of the request
  // 
  // example:
  // 
  // 473469C7***B-A3DC0DE3C83E
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
  // Result
  Result *EntElementVerifyV2ResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s EntElementVerifyV2ResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EntElementVerifyV2ResponseBody) GoString() string {
  return s.String()
}

func (s *EntElementVerifyV2ResponseBody) GetCode() *string  {
  return s.Code
}

func (s *EntElementVerifyV2ResponseBody) GetMessage() *string  {
  return s.Message
}

func (s *EntElementVerifyV2ResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EntElementVerifyV2ResponseBody) GetResult() *EntElementVerifyV2ResponseBodyResult  {
  return s.Result
}

func (s *EntElementVerifyV2ResponseBody) SetCode(v string) *EntElementVerifyV2ResponseBody {
  s.Code = &v
  return s
}

func (s *EntElementVerifyV2ResponseBody) SetMessage(v string) *EntElementVerifyV2ResponseBody {
  s.Message = &v
  return s
}

func (s *EntElementVerifyV2ResponseBody) SetRequestId(v string) *EntElementVerifyV2ResponseBody {
  s.RequestId = &v
  return s
}

func (s *EntElementVerifyV2ResponseBody) SetResult(v *EntElementVerifyV2ResponseBodyResult) *EntElementVerifyV2ResponseBody {
  s.Result = v
  return s
}

func (s *EntElementVerifyV2ResponseBody) Validate() error {
  if s.Result != nil {
    if err := s.Result.Validate(); err != nil {
      return err
    }
  }
  return nil
}

type EntElementVerifyV2ResponseBodyResult struct {
  // Verification result code.
  // 
  // - 1: Verification consistent
  // 
  // - 2: Verification inconsistent
  // 
  // - 3: Not found
  // 
  // example:
  // 
  // 1
  BizCode *string `json:"BizCode,omitempty" xml:"BizCode,omitempty"`
  // Business term: start and end time of operations.
  // 
  // - Format: yyyy-MM-dd/yyyy-MM-dd.
  // 
  // - Example: 2018-09-25/9999-09-09.
  // 
  // example:
  // 
  // 2018-09-25/9999-09-09
  OpenTime *string `json:"OpenTime,omitempty" xml:"OpenTime,omitempty"`
  // Details of inconsistencies, multiple inconsistencies will be separated by commas.
  // 
  // - LegalPersonNameFlag: Legal person\\"s name does not match
  // 
  // - LegalPersonCertNoFlag: Legal person\\"s ID number does not match
  // 
  // - EntNameFlag: Enterprise name does not match
  // 
  // - LicenseNoFlag: Business license number does not match
  // 
  // example:
  // 
  // LegalPersonNameFlag,LegalPersonCertNoFlag
  ReasonDetail *string `json:"ReasonDetail,omitempty" xml:"ReasonDetail,omitempty"`
  // Business operation status.
  // 
  // - 1: In operation (open)
  // 
  // - 2: Relocated
  // 
  // - 3: Deregistered
  // 
  // - 4: Revoked
  // 
  // - 5: Canceled
  // 
  // - 6: Suspended
  // 
  // - 0: Other
  // 
  // example:
  // 
  // 1
  Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s EntElementVerifyV2ResponseBodyResult) String() string {
  return dara.Prettify(s)
}

func (s EntElementVerifyV2ResponseBodyResult) GoString() string {
  return s.String()
}

func (s *EntElementVerifyV2ResponseBodyResult) GetBizCode() *string  {
  return s.BizCode
}

func (s *EntElementVerifyV2ResponseBodyResult) GetOpenTime() *string  {
  return s.OpenTime
}

func (s *EntElementVerifyV2ResponseBodyResult) GetReasonDetail() *string  {
  return s.ReasonDetail
}

func (s *EntElementVerifyV2ResponseBodyResult) GetStatus() *string  {
  return s.Status
}

func (s *EntElementVerifyV2ResponseBodyResult) SetBizCode(v string) *EntElementVerifyV2ResponseBodyResult {
  s.BizCode = &v
  return s
}

func (s *EntElementVerifyV2ResponseBodyResult) SetOpenTime(v string) *EntElementVerifyV2ResponseBodyResult {
  s.OpenTime = &v
  return s
}

func (s *EntElementVerifyV2ResponseBodyResult) SetReasonDetail(v string) *EntElementVerifyV2ResponseBodyResult {
  s.ReasonDetail = &v
  return s
}

func (s *EntElementVerifyV2ResponseBodyResult) SetStatus(v string) *EntElementVerifyV2ResponseBodyResult {
  s.Status = &v
  return s
}

func (s *EntElementVerifyV2ResponseBodyResult) Validate() error {
  return dara.Validate(s)
}

