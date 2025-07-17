// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteDataExportResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetErrorCode(v string) *ExecuteDataExportResponseBody
  GetErrorCode() *string 
  SetErrorMessage(v string) *ExecuteDataExportResponseBody
  GetErrorMessage() *string 
  SetRequestId(v string) *ExecuteDataExportResponseBody
  GetRequestId() *string 
  SetSuccess(v bool) *ExecuteDataExportResponseBody
  GetSuccess() *bool 
}

type ExecuteDataExportResponseBody struct {
  // The error code returned if the request failed.
  // 
  // example:
  // 
  // UnknownError
  ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
  // The error message returned if the request failed.
  // 
  // example:
  // 
  // UnknownError
  ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
  // The ID of the request.
  // 
  // example:
  // 
  // FE8EE2F1-4880-46BC-A704-5CF63EAF9A04
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
  // Indicates whether the request is successful.
  // 
  // example:
  // 
  // true
  Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ExecuteDataExportResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ExecuteDataExportResponseBody) GoString() string {
  return s.String()
}

func (s *ExecuteDataExportResponseBody) GetErrorCode() *string  {
  return s.ErrorCode
}

func (s *ExecuteDataExportResponseBody) GetErrorMessage() *string  {
  return s.ErrorMessage
}

func (s *ExecuteDataExportResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *ExecuteDataExportResponseBody) GetSuccess() *bool  {
  return s.Success
}

func (s *ExecuteDataExportResponseBody) SetErrorCode(v string) *ExecuteDataExportResponseBody {
  s.ErrorCode = &v
  return s
}

func (s *ExecuteDataExportResponseBody) SetErrorMessage(v string) *ExecuteDataExportResponseBody {
  s.ErrorMessage = &v
  return s
}

func (s *ExecuteDataExportResponseBody) SetRequestId(v string) *ExecuteDataExportResponseBody {
  s.RequestId = &v
  return s
}

func (s *ExecuteDataExportResponseBody) SetSuccess(v bool) *ExecuteDataExportResponseBody {
  s.Success = &v
  return s
}

func (s *ExecuteDataExportResponseBody) Validate() error {
  return dara.Validate(s)
}

