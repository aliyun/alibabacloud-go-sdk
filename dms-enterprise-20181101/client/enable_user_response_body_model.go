// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableUserResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetErrorCode(v string) *EnableUserResponseBody
  GetErrorCode() *string 
  SetErrorMessage(v string) *EnableUserResponseBody
  GetErrorMessage() *string 
  SetRequestId(v string) *EnableUserResponseBody
  GetRequestId() *string 
  SetSuccess(v bool) *EnableUserResponseBody
  GetSuccess() *bool 
}

type EnableUserResponseBody struct {
  // The error code returned if the request fails.
  // 
  // example:
  // 
  // 403
  ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
  // The error message returned if the request fails.
  // 
  // example:
  // 
  // The specified user not exists.
  ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
  // The ID of the request.
  // 
  // example:
  // 
  // 34E01EDD-6A16-4CF0-9541-C644D1BE01AA
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

func (s EnableUserResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnableUserResponseBody) GoString() string {
  return s.String()
}

func (s *EnableUserResponseBody) GetErrorCode() *string  {
  return s.ErrorCode
}

func (s *EnableUserResponseBody) GetErrorMessage() *string  {
  return s.ErrorMessage
}

func (s *EnableUserResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnableUserResponseBody) GetSuccess() *bool  {
  return s.Success
}

func (s *EnableUserResponseBody) SetErrorCode(v string) *EnableUserResponseBody {
  s.ErrorCode = &v
  return s
}

func (s *EnableUserResponseBody) SetErrorMessage(v string) *EnableUserResponseBody {
  s.ErrorMessage = &v
  return s
}

func (s *EnableUserResponseBody) SetRequestId(v string) *EnableUserResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnableUserResponseBody) SetSuccess(v bool) *EnableUserResponseBody {
  s.Success = &v
  return s
}

func (s *EnableUserResponseBody) Validate() error {
  return dara.Validate(s)
}

