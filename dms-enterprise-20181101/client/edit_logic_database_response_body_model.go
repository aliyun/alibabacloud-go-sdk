// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEditLogicDatabaseResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetErrorCode(v string) *EditLogicDatabaseResponseBody
  GetErrorCode() *string 
  SetErrorMessage(v string) *EditLogicDatabaseResponseBody
  GetErrorMessage() *string 
  SetRequestId(v string) *EditLogicDatabaseResponseBody
  GetRequestId() *string 
  SetSuccess(v bool) *EditLogicDatabaseResponseBody
  GetSuccess() *bool 
}

type EditLogicDatabaseResponseBody struct {
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
  // C51420E3-144A-4A94-B473-8662FCF4AD10
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
  // Indicates whether the request is successful. Valid values:
  // 
  // - true: The request is successful.
  // 
  // - false: The request fails.
  // 
  // example:
  // 
  // true
  Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s EditLogicDatabaseResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EditLogicDatabaseResponseBody) GoString() string {
  return s.String()
}

func (s *EditLogicDatabaseResponseBody) GetErrorCode() *string  {
  return s.ErrorCode
}

func (s *EditLogicDatabaseResponseBody) GetErrorMessage() *string  {
  return s.ErrorMessage
}

func (s *EditLogicDatabaseResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EditLogicDatabaseResponseBody) GetSuccess() *bool  {
  return s.Success
}

func (s *EditLogicDatabaseResponseBody) SetErrorCode(v string) *EditLogicDatabaseResponseBody {
  s.ErrorCode = &v
  return s
}

func (s *EditLogicDatabaseResponseBody) SetErrorMessage(v string) *EditLogicDatabaseResponseBody {
  s.ErrorMessage = &v
  return s
}

func (s *EditLogicDatabaseResponseBody) SetRequestId(v string) *EditLogicDatabaseResponseBody {
  s.RequestId = &v
  return s
}

func (s *EditLogicDatabaseResponseBody) SetSuccess(v bool) *EditLogicDatabaseResponseBody {
  s.Success = &v
  return s
}

func (s *EditLogicDatabaseResponseBody) Validate() error {
  return dara.Validate(s)
}

