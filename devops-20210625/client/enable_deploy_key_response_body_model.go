// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableDeployKeyResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetErrorCode(v string) *EnableDeployKeyResponseBody
  GetErrorCode() *string 
  SetErrorMessage(v string) *EnableDeployKeyResponseBody
  GetErrorMessage() *string 
  SetRequestId(v string) *EnableDeployKeyResponseBody
  GetRequestId() *string 
  SetResult(v *EnableDeployKeyResponseBodyResult) *EnableDeployKeyResponseBody
  GetResult() *EnableDeployKeyResponseBodyResult 
  SetSuccess(v bool) *EnableDeployKeyResponseBody
  GetSuccess() *bool 
}

type EnableDeployKeyResponseBody struct {
  // example:
  // 
  // Openapi.RequestError
  ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
  // example:
  // 
  // ""
  ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
  // example:
  // 
  // ASSDS-ASSASX-XSAXSA-XSAXSAXS
  RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
  Result *EnableDeployKeyResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
  // example:
  // 
  // true
  Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s EnableDeployKeyResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnableDeployKeyResponseBody) GoString() string {
  return s.String()
}

func (s *EnableDeployKeyResponseBody) GetErrorCode() *string  {
  return s.ErrorCode
}

func (s *EnableDeployKeyResponseBody) GetErrorMessage() *string  {
  return s.ErrorMessage
}

func (s *EnableDeployKeyResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnableDeployKeyResponseBody) GetResult() *EnableDeployKeyResponseBodyResult  {
  return s.Result
}

func (s *EnableDeployKeyResponseBody) GetSuccess() *bool  {
  return s.Success
}

func (s *EnableDeployKeyResponseBody) SetErrorCode(v string) *EnableDeployKeyResponseBody {
  s.ErrorCode = &v
  return s
}

func (s *EnableDeployKeyResponseBody) SetErrorMessage(v string) *EnableDeployKeyResponseBody {
  s.ErrorMessage = &v
  return s
}

func (s *EnableDeployKeyResponseBody) SetRequestId(v string) *EnableDeployKeyResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnableDeployKeyResponseBody) SetResult(v *EnableDeployKeyResponseBodyResult) *EnableDeployKeyResponseBody {
  s.Result = v
  return s
}

func (s *EnableDeployKeyResponseBody) SetSuccess(v bool) *EnableDeployKeyResponseBody {
  s.Success = &v
  return s
}

func (s *EnableDeployKeyResponseBody) Validate() error {
  if s.Result != nil {
    if err := s.Result.Validate(); err != nil {
      return err
    }
  }
  return nil
}

type EnableDeployKeyResponseBodyResult struct {
  // example:
  // 
  // true
  Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s EnableDeployKeyResponseBodyResult) String() string {
  return dara.Prettify(s)
}

func (s EnableDeployKeyResponseBodyResult) GoString() string {
  return s.String()
}

func (s *EnableDeployKeyResponseBodyResult) GetResult() *bool  {
  return s.Result
}

func (s *EnableDeployKeyResponseBodyResult) SetResult(v bool) *EnableDeployKeyResponseBodyResult {
  s.Result = &v
  return s
}

func (s *EnableDeployKeyResponseBodyResult) Validate() error {
  return dara.Validate(s)
}

