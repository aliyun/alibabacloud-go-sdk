// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEntVerifyResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetCode(v string) *EntVerifyResponseBody
  GetCode() *string 
  SetMessage(v string) *EntVerifyResponseBody
  GetMessage() *string 
  SetRequestId(v string) *EntVerifyResponseBody
  GetRequestId() *string 
  SetResult(v *EntVerifyResponseBodyResult) *EntVerifyResponseBody
  GetResult() *EntVerifyResponseBodyResult 
}

type EntVerifyResponseBody struct {
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
  // 成功
  Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
  // ID of the request
  // 
  // example:
  // 
  // 473469C7-A***B-A3DC0DE3C83E
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
  // Result
  Result *EntVerifyResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s EntVerifyResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EntVerifyResponseBody) GoString() string {
  return s.String()
}

func (s *EntVerifyResponseBody) GetCode() *string  {
  return s.Code
}

func (s *EntVerifyResponseBody) GetMessage() *string  {
  return s.Message
}

func (s *EntVerifyResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EntVerifyResponseBody) GetResult() *EntVerifyResponseBodyResult  {
  return s.Result
}

func (s *EntVerifyResponseBody) SetCode(v string) *EntVerifyResponseBody {
  s.Code = &v
  return s
}

func (s *EntVerifyResponseBody) SetMessage(v string) *EntVerifyResponseBody {
  s.Message = &v
  return s
}

func (s *EntVerifyResponseBody) SetRequestId(v string) *EntVerifyResponseBody {
  s.RequestId = &v
  return s
}

func (s *EntVerifyResponseBody) SetResult(v *EntVerifyResponseBodyResult) *EntVerifyResponseBody {
  s.Result = v
  return s
}

func (s *EntVerifyResponseBody) Validate() error {
  if s.Result != nil {
    if err := s.Result.Validate(); err != nil {
      return err
    }
  }
  return nil
}

type EntVerifyResponseBodyResult struct {
  // Enterprise risk verification result
  RiskVerifyResult *EntVerifyResponseBodyResultRiskVerifyResult `json:"RiskVerifyResult,omitempty" xml:"RiskVerifyResult,omitempty" type:"Struct"`
}

func (s EntVerifyResponseBodyResult) String() string {
  return dara.Prettify(s)
}

func (s EntVerifyResponseBodyResult) GoString() string {
  return s.String()
}

func (s *EntVerifyResponseBodyResult) GetRiskVerifyResult() *EntVerifyResponseBodyResultRiskVerifyResult  {
  return s.RiskVerifyResult
}

func (s *EntVerifyResponseBodyResult) SetRiskVerifyResult(v *EntVerifyResponseBodyResultRiskVerifyResult) *EntVerifyResponseBodyResult {
  s.RiskVerifyResult = v
  return s
}

func (s *EntVerifyResponseBodyResult) Validate() error {
  if s.RiskVerifyResult != nil {
    if err := s.RiskVerifyResult.Validate(); err != nil {
      return err
    }
  }
  return nil
}

type EntVerifyResponseBodyResultRiskVerifyResult struct {
  // Whether found
  // 
  // example:
  // 
  // true
  Found *bool `json:"Found,omitempty" xml:"Found,omitempty"`
  // List of enterprise risk verification model results
  ModelResults []*EntVerifyResponseBodyResultRiskVerifyResultModelResults `json:"ModelResults,omitempty" xml:"ModelResults,omitempty" type:"Repeated"`
}

func (s EntVerifyResponseBodyResultRiskVerifyResult) String() string {
  return dara.Prettify(s)
}

func (s EntVerifyResponseBodyResultRiskVerifyResult) GoString() string {
  return s.String()
}

func (s *EntVerifyResponseBodyResultRiskVerifyResult) GetFound() *bool  {
  return s.Found
}

func (s *EntVerifyResponseBodyResultRiskVerifyResult) GetModelResults() []*EntVerifyResponseBodyResultRiskVerifyResultModelResults  {
  return s.ModelResults
}

func (s *EntVerifyResponseBodyResultRiskVerifyResult) SetFound(v bool) *EntVerifyResponseBodyResultRiskVerifyResult {
  s.Found = &v
  return s
}

func (s *EntVerifyResponseBodyResultRiskVerifyResult) SetModelResults(v []*EntVerifyResponseBodyResultRiskVerifyResultModelResults) *EntVerifyResponseBodyResultRiskVerifyResult {
  s.ModelResults = v
  return s
}

func (s *EntVerifyResponseBodyResultRiskVerifyResult) Validate() error {
  if s.ModelResults != nil {
    for _, item := range s.ModelResults {
      if item != nil {
        if err := item.Validate(); err != nil {
          return err
        }
      }
    }
  }
  return nil
}

type EntVerifyResponseBodyResultRiskVerifyResultModelResults struct {
  // Model name
  // 
  // example:
  // 
  // model_1
  ModelName *string `json:"ModelName,omitempty" xml:"ModelName,omitempty"`
  // Model result
  // 
  // example:
  // 
  // 5
  Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s EntVerifyResponseBodyResultRiskVerifyResultModelResults) String() string {
  return dara.Prettify(s)
}

func (s EntVerifyResponseBodyResultRiskVerifyResultModelResults) GoString() string {
  return s.String()
}

func (s *EntVerifyResponseBodyResultRiskVerifyResultModelResults) GetModelName() *string  {
  return s.ModelName
}

func (s *EntVerifyResponseBodyResultRiskVerifyResultModelResults) GetResult() *string  {
  return s.Result
}

func (s *EntVerifyResponseBodyResultRiskVerifyResultModelResults) SetModelName(v string) *EntVerifyResponseBodyResultRiskVerifyResultModelResults {
  s.ModelName = &v
  return s
}

func (s *EntVerifyResponseBodyResultRiskVerifyResultModelResults) SetResult(v string) *EntVerifyResponseBodyResultRiskVerifyResultModelResults {
  s.Result = &v
  return s
}

func (s *EntVerifyResponseBodyResultRiskVerifyResultModelResults) Validate() error {
  return dara.Validate(s)
}

