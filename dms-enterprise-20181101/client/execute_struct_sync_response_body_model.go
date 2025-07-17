// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteStructSyncResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetErrorCode(v string) *ExecuteStructSyncResponseBody
  GetErrorCode() *string 
  SetErrorMessage(v string) *ExecuteStructSyncResponseBody
  GetErrorMessage() *string 
  SetRequestId(v string) *ExecuteStructSyncResponseBody
  GetRequestId() *string 
  SetSuccess(v bool) *ExecuteStructSyncResponseBody
  GetSuccess() *bool 
}

type ExecuteStructSyncResponseBody struct {
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
  // UnknownError
  ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
  // The ID of the request.
  // 
  // example:
  // 
  // 4E1D2B4D-3E53-4ABC-999D-1D2520B3471A
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
  // Indicates whether the request was successful.
  // 
  // example:
  // 
  // true
  Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ExecuteStructSyncResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ExecuteStructSyncResponseBody) GoString() string {
  return s.String()
}

func (s *ExecuteStructSyncResponseBody) GetErrorCode() *string  {
  return s.ErrorCode
}

func (s *ExecuteStructSyncResponseBody) GetErrorMessage() *string  {
  return s.ErrorMessage
}

func (s *ExecuteStructSyncResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *ExecuteStructSyncResponseBody) GetSuccess() *bool  {
  return s.Success
}

func (s *ExecuteStructSyncResponseBody) SetErrorCode(v string) *ExecuteStructSyncResponseBody {
  s.ErrorCode = &v
  return s
}

func (s *ExecuteStructSyncResponseBody) SetErrorMessage(v string) *ExecuteStructSyncResponseBody {
  s.ErrorMessage = &v
  return s
}

func (s *ExecuteStructSyncResponseBody) SetRequestId(v string) *ExecuteStructSyncResponseBody {
  s.RequestId = &v
  return s
}

func (s *ExecuteStructSyncResponseBody) SetSuccess(v bool) *ExecuteStructSyncResponseBody {
  s.Success = &v
  return s
}

func (s *ExecuteStructSyncResponseBody) Validate() error {
  return dara.Validate(s)
}

