// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEntElementVerifyResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetCode(v string) *EntElementVerifyResponseBody
  GetCode() *string 
  SetMessage(v string) *EntElementVerifyResponseBody
  GetMessage() *string 
  SetRequestId(v string) *EntElementVerifyResponseBody
  GetRequestId() *string 
  SetResult(v *EntElementVerifyResponseBodyResult) *EntElementVerifyResponseBody
  GetResult() *EntElementVerifyResponseBodyResult 
}

type EntElementVerifyResponseBody struct {
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
  // ID of the request
  // 
  // example:
  // 
  // 473469C7***B-A3DC0DE3C83E
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
  // Result
  Result *EntElementVerifyResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s EntElementVerifyResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EntElementVerifyResponseBody) GoString() string {
  return s.String()
}

func (s *EntElementVerifyResponseBody) GetCode() *string  {
  return s.Code
}

func (s *EntElementVerifyResponseBody) GetMessage() *string  {
  return s.Message
}

func (s *EntElementVerifyResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EntElementVerifyResponseBody) GetResult() *EntElementVerifyResponseBodyResult  {
  return s.Result
}

func (s *EntElementVerifyResponseBody) SetCode(v string) *EntElementVerifyResponseBody {
  s.Code = &v
  return s
}

func (s *EntElementVerifyResponseBody) SetMessage(v string) *EntElementVerifyResponseBody {
  s.Message = &v
  return s
}

func (s *EntElementVerifyResponseBody) SetRequestId(v string) *EntElementVerifyResponseBody {
  s.RequestId = &v
  return s
}

func (s *EntElementVerifyResponseBody) SetResult(v *EntElementVerifyResponseBodyResult) *EntElementVerifyResponseBody {
  s.Result = v
  return s
}

func (s *EntElementVerifyResponseBody) Validate() error {
  if s.Result != nil {
    if err := s.Result.Validate(); err != nil {
      return err
    }
  }
  return nil
}

type EntElementVerifyResponseBodyResult struct {
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
  // Operating period: start and end time of operations.
  // 
  // - Format: yyyy-MM-dd/yyyy-MM-dd.
  // 
  // - Example: 2018-09-25/9999-09-09.
  // 
  // example:
  // 
  // 2018-09-25/9999-09-09
  OpenTime *string `json:"OpenTime,omitempty" xml:"OpenTime,omitempty"`
  // ReasonCode explanation:
  // 
  // - 100: Verification consistent
  // 
  // - 201: Inconsistent between person and enterprise
  // 
  // - 202: Inconsistent in two elements of the enterprise
  // 
  // - 301: No enterprise found, unable to verify
  // 
  // - 302: Legal representative does not exist in the database
  // 
  // 	Warning: This field will be discontinued on June 15, 2023. Existing customers are not affected.
  // 
  // example:
  // 
  // 100
  ReasonCode *string `json:"ReasonCode,omitempty" xml:"ReasonCode,omitempty"`
  // Details of inconsistencies, multiple inconsistencies will be separated by commas.
  // 
  // - LegalPersonNameFlag: Inconsistent legal representative name
  // 
  // - LegalPersonCertNoFlag: Inconsistent legal representative ID number
  // 
  // - EntNameFlag: Inconsistent enterprise name
  // 
  // - LicenseNoFlag: Inconsistent business license number
  // 
  // example:
  // 
  // LegalPersonNameFlag,LegalPersonCertNoFlag
  ReasonDetail *string `json:"ReasonDetail,omitempty" xml:"ReasonDetail,omitempty"`
  // Enterprise operating status.
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

func (s EntElementVerifyResponseBodyResult) String() string {
  return dara.Prettify(s)
}

func (s EntElementVerifyResponseBodyResult) GoString() string {
  return s.String()
}

func (s *EntElementVerifyResponseBodyResult) GetBizCode() *string  {
  return s.BizCode
}

func (s *EntElementVerifyResponseBodyResult) GetOpenTime() *string  {
  return s.OpenTime
}

func (s *EntElementVerifyResponseBodyResult) GetReasonCode() *string  {
  return s.ReasonCode
}

func (s *EntElementVerifyResponseBodyResult) GetReasonDetail() *string  {
  return s.ReasonDetail
}

func (s *EntElementVerifyResponseBodyResult) GetStatus() *string  {
  return s.Status
}

func (s *EntElementVerifyResponseBodyResult) SetBizCode(v string) *EntElementVerifyResponseBodyResult {
  s.BizCode = &v
  return s
}

func (s *EntElementVerifyResponseBodyResult) SetOpenTime(v string) *EntElementVerifyResponseBodyResult {
  s.OpenTime = &v
  return s
}

func (s *EntElementVerifyResponseBodyResult) SetReasonCode(v string) *EntElementVerifyResponseBodyResult {
  s.ReasonCode = &v
  return s
}

func (s *EntElementVerifyResponseBodyResult) SetReasonDetail(v string) *EntElementVerifyResponseBodyResult {
  s.ReasonDetail = &v
  return s
}

func (s *EntElementVerifyResponseBodyResult) SetStatus(v string) *EntElementVerifyResponseBodyResult {
  s.Status = &v
  return s
}

func (s *EntElementVerifyResponseBodyResult) Validate() error {
  return dara.Validate(s)
}

