// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteDataCorrectResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetErrorCode(v string) *ExecuteDataCorrectResponseBody
  GetErrorCode() *string 
  SetErrorMessage(v string) *ExecuteDataCorrectResponseBody
  GetErrorMessage() *string 
  SetRequestId(v string) *ExecuteDataCorrectResponseBody
  GetRequestId() *string 
  SetSuccess(v bool) *ExecuteDataCorrectResponseBody
  GetSuccess() *bool 
}

type ExecuteDataCorrectResponseBody struct {
  // The error code.
  // 
  // example:
  // 
  // UnknownError
  ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
  // The error message.
  // 
  // example:
  // 
  // Unknown server error
  ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
  // The ID of the request.
  // 
  // example:
  // 
  // EADDA791-2809-58CE-A303-743A77FF****
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
  // Indicates whether the request is successful. Valid values:
  // 
  // 	- **true**: The request is successful.
  // 
  // 	- **false**: The request fails.
  // 
  // example:
  // 
  // true
  Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ExecuteDataCorrectResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ExecuteDataCorrectResponseBody) GoString() string {
  return s.String()
}

func (s *ExecuteDataCorrectResponseBody) GetErrorCode() *string  {
  return s.ErrorCode
}

func (s *ExecuteDataCorrectResponseBody) GetErrorMessage() *string  {
  return s.ErrorMessage
}

func (s *ExecuteDataCorrectResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *ExecuteDataCorrectResponseBody) GetSuccess() *bool  {
  return s.Success
}

func (s *ExecuteDataCorrectResponseBody) SetErrorCode(v string) *ExecuteDataCorrectResponseBody {
  s.ErrorCode = &v
  return s
}

func (s *ExecuteDataCorrectResponseBody) SetErrorMessage(v string) *ExecuteDataCorrectResponseBody {
  s.ErrorMessage = &v
  return s
}

func (s *ExecuteDataCorrectResponseBody) SetRequestId(v string) *ExecuteDataCorrectResponseBody {
  s.RequestId = &v
  return s
}

func (s *ExecuteDataCorrectResponseBody) SetSuccess(v bool) *ExecuteDataCorrectResponseBody {
  s.Success = &v
  return s
}

func (s *ExecuteDataCorrectResponseBody) Validate() error {
  return dara.Validate(s)
}

