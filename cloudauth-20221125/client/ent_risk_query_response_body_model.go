// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEntRiskQueryResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetCode(v string) *EntRiskQueryResponseBody
  GetCode() *string 
  SetMessage(v string) *EntRiskQueryResponseBody
  GetMessage() *string 
  SetRequestId(v string) *EntRiskQueryResponseBody
  GetRequestId() *string 
  SetResult(v *EntRiskQueryResponseBodyResult) *EntRiskQueryResponseBody
  GetResult() *EntRiskQueryResponseBodyResult 
}

type EntRiskQueryResponseBody struct {
  // Error code. For details about error codes, see **[Error Codes](https://help.aliyun.com/document_detail/215420.html)**.
  // 
  // example:
  // 
  // 200
  Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
  // Response message for the request information.
  // 
  // example:
  // 
  // success
  Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
  // Request ID
  // 
  // example:
  // 
  // 8FC3D6AC-9FED-4311-8DA7-C4BF47D9F260
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
  // Result
  Result *EntRiskQueryResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s EntRiskQueryResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EntRiskQueryResponseBody) GoString() string {
  return s.String()
}

func (s *EntRiskQueryResponseBody) GetCode() *string  {
  return s.Code
}

func (s *EntRiskQueryResponseBody) GetMessage() *string  {
  return s.Message
}

func (s *EntRiskQueryResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EntRiskQueryResponseBody) GetResult() *EntRiskQueryResponseBodyResult  {
  return s.Result
}

func (s *EntRiskQueryResponseBody) SetCode(v string) *EntRiskQueryResponseBody {
  s.Code = &v
  return s
}

func (s *EntRiskQueryResponseBody) SetMessage(v string) *EntRiskQueryResponseBody {
  s.Message = &v
  return s
}

func (s *EntRiskQueryResponseBody) SetRequestId(v string) *EntRiskQueryResponseBody {
  s.RequestId = &v
  return s
}

func (s *EntRiskQueryResponseBody) SetResult(v *EntRiskQueryResponseBodyResult) *EntRiskQueryResponseBody {
  s.Result = v
  return s
}

func (s *EntRiskQueryResponseBody) Validate() error {
  if s.Result != nil {
    if err := s.Result.Validate(); err != nil {
      return err
    }
  }
  return nil
}

type EntRiskQueryResponseBodyResult struct {
  // Query result
  // 
  // 0: Normal business operation
  // 
  // 1: Abnormal business operation
  // 
  // 2: Not found
  // 
  // example:
  // 
  // 1
  BizCode *string `json:"BizCode,omitempty" xml:"BizCode,omitempty"`
  // List of abnormal information
  RiskList []*EntRiskQueryResponseBodyResultRiskList `json:"RiskList,omitempty" xml:"RiskList,omitempty" type:"Repeated"`
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

func (s EntRiskQueryResponseBodyResult) String() string {
  return dara.Prettify(s)
}

func (s EntRiskQueryResponseBodyResult) GoString() string {
  return s.String()
}

func (s *EntRiskQueryResponseBodyResult) GetBizCode() *string  {
  return s.BizCode
}

func (s *EntRiskQueryResponseBodyResult) GetRiskList() []*EntRiskQueryResponseBodyResultRiskList  {
  return s.RiskList
}

func (s *EntRiskQueryResponseBodyResult) GetStatus() *string  {
  return s.Status
}

func (s *EntRiskQueryResponseBodyResult) SetBizCode(v string) *EntRiskQueryResponseBodyResult {
  s.BizCode = &v
  return s
}

func (s *EntRiskQueryResponseBodyResult) SetRiskList(v []*EntRiskQueryResponseBodyResultRiskList) *EntRiskQueryResponseBodyResult {
  s.RiskList = v
  return s
}

func (s *EntRiskQueryResponseBodyResult) SetStatus(v string) *EntRiskQueryResponseBodyResult {
  s.Status = &v
  return s
}

func (s *EntRiskQueryResponseBodyResult) Validate() error {
  if s.RiskList != nil {
    for _, item := range s.RiskList {
      if item != nil {
        if err := item.Validate(); err != nil {
          return err
        }
      }
    }
  }
  return nil
}

type EntRiskQueryResponseBodyResultRiskList struct {
  // Unified Social Credit Code
  // 
  // example:
  // 
  // 92500112MA5UHU****
  CreditCode *string `json:"CreditCode,omitempty" xml:"CreditCode,omitempty"`
  // Company name.
  // 
  // example:
  // 
  // 杭州****
  EntName *string `json:"EntName,omitempty" xml:"EntName,omitempty"`
  // Date listed as abnormal
  // 
  // Example: 2023-02-03
  // 
  // example:
  // 
  // 2023-02-03
  ListedDate *string `json:"ListedDate,omitempty" xml:"ListedDate,omitempty"`
  // Reason for being listed as abnormal
  // 
  // example:
  // 
  // 未按照《个体工商户年度报告暂行办法》***
  ListedReason *string `json:"ListedReason,omitempty" xml:"ListedReason,omitempty"`
  // Authority that handled the inclusion
  // 
  // example:
  // 
  // ***市场监督***
  OperationOrg *string `json:"OperationOrg,omitempty" xml:"OperationOrg,omitempty"`
  // Business registration number
  // 
  // example:
  // 
  // 50011260996****
  RegNo *string `json:"RegNo,omitempty" xml:"RegNo,omitempty"`
  // Date removed from the abnormal list
  // 
  // Example: 2023-02-03
  // 
  // example:
  // 
  // 2023-02-06
  RemovedDate *string `json:"RemovedDate,omitempty" xml:"RemovedDate,omitempty"`
  // Authority that handled the removal
  // 
  // example:
  // 
  // ***市场监督***
  RemovedOrg *string `json:"RemovedOrg,omitempty" xml:"RemovedOrg,omitempty"`
  // Reason for being removed from the abnormal list
  // 
  // example:
  // 
  // 根据《个体工商户年度报告暂行办法》第十三条的规定******
  RemovedReason *string `json:"RemovedReason,omitempty" xml:"RemovedReason,omitempty"`
}

func (s EntRiskQueryResponseBodyResultRiskList) String() string {
  return dara.Prettify(s)
}

func (s EntRiskQueryResponseBodyResultRiskList) GoString() string {
  return s.String()
}

func (s *EntRiskQueryResponseBodyResultRiskList) GetCreditCode() *string  {
  return s.CreditCode
}

func (s *EntRiskQueryResponseBodyResultRiskList) GetEntName() *string  {
  return s.EntName
}

func (s *EntRiskQueryResponseBodyResultRiskList) GetListedDate() *string  {
  return s.ListedDate
}

func (s *EntRiskQueryResponseBodyResultRiskList) GetListedReason() *string  {
  return s.ListedReason
}

func (s *EntRiskQueryResponseBodyResultRiskList) GetOperationOrg() *string  {
  return s.OperationOrg
}

func (s *EntRiskQueryResponseBodyResultRiskList) GetRegNo() *string  {
  return s.RegNo
}

func (s *EntRiskQueryResponseBodyResultRiskList) GetRemovedDate() *string  {
  return s.RemovedDate
}

func (s *EntRiskQueryResponseBodyResultRiskList) GetRemovedOrg() *string  {
  return s.RemovedOrg
}

func (s *EntRiskQueryResponseBodyResultRiskList) GetRemovedReason() *string  {
  return s.RemovedReason
}

func (s *EntRiskQueryResponseBodyResultRiskList) SetCreditCode(v string) *EntRiskQueryResponseBodyResultRiskList {
  s.CreditCode = &v
  return s
}

func (s *EntRiskQueryResponseBodyResultRiskList) SetEntName(v string) *EntRiskQueryResponseBodyResultRiskList {
  s.EntName = &v
  return s
}

func (s *EntRiskQueryResponseBodyResultRiskList) SetListedDate(v string) *EntRiskQueryResponseBodyResultRiskList {
  s.ListedDate = &v
  return s
}

func (s *EntRiskQueryResponseBodyResultRiskList) SetListedReason(v string) *EntRiskQueryResponseBodyResultRiskList {
  s.ListedReason = &v
  return s
}

func (s *EntRiskQueryResponseBodyResultRiskList) SetOperationOrg(v string) *EntRiskQueryResponseBodyResultRiskList {
  s.OperationOrg = &v
  return s
}

func (s *EntRiskQueryResponseBodyResultRiskList) SetRegNo(v string) *EntRiskQueryResponseBodyResultRiskList {
  s.RegNo = &v
  return s
}

func (s *EntRiskQueryResponseBodyResultRiskList) SetRemovedDate(v string) *EntRiskQueryResponseBodyResultRiskList {
  s.RemovedDate = &v
  return s
}

func (s *EntRiskQueryResponseBodyResultRiskList) SetRemovedOrg(v string) *EntRiskQueryResponseBodyResultRiskList {
  s.RemovedOrg = &v
  return s
}

func (s *EntRiskQueryResponseBodyResultRiskList) SetRemovedReason(v string) *EntRiskQueryResponseBodyResultRiskList {
  s.RemovedReason = &v
  return s
}

func (s *EntRiskQueryResponseBodyResultRiskList) Validate() error {
  return dara.Validate(s)
}

