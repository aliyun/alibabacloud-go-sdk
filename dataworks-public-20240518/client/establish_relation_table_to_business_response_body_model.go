// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEstablishRelationTableToBusinessResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetErrorCode(v string) *EstablishRelationTableToBusinessResponseBody
  GetErrorCode() *string 
  SetErrorMessage(v string) *EstablishRelationTableToBusinessResponseBody
  GetErrorMessage() *string 
  SetHttpStatusCode(v int32) *EstablishRelationTableToBusinessResponseBody
  GetHttpStatusCode() *int32 
  SetRequestId(v string) *EstablishRelationTableToBusinessResponseBody
  GetRequestId() *string 
  SetSuccess(v bool) *EstablishRelationTableToBusinessResponseBody
  GetSuccess() *bool 
}

type EstablishRelationTableToBusinessResponseBody struct {
  // The error code.
  // 
  // example:
  // 
  // Invalid.Tenant.ConnectionNotExists
  ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
  // The error message.
  // 
  // example:
  // 
  // The connection does not exist.
  ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
  // The HTTP status code.
  // 
  // example:
  // 
  // 200
  HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
  // The unique ID of this request. You can use this ID to troubleshoot issues if errors occur.
  // 
  // example:
  // 
  // 0000-ABCD-EFG
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
  // Indicates whether the request succeeded. Valid values:
  // 
  // 	- true
  // 
  // 	- false
  // 
  // example:
  // 
  // true
  Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s EstablishRelationTableToBusinessResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EstablishRelationTableToBusinessResponseBody) GoString() string {
  return s.String()
}

func (s *EstablishRelationTableToBusinessResponseBody) GetErrorCode() *string  {
  return s.ErrorCode
}

func (s *EstablishRelationTableToBusinessResponseBody) GetErrorMessage() *string  {
  return s.ErrorMessage
}

func (s *EstablishRelationTableToBusinessResponseBody) GetHttpStatusCode() *int32  {
  return s.HttpStatusCode
}

func (s *EstablishRelationTableToBusinessResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EstablishRelationTableToBusinessResponseBody) GetSuccess() *bool  {
  return s.Success
}

func (s *EstablishRelationTableToBusinessResponseBody) SetErrorCode(v string) *EstablishRelationTableToBusinessResponseBody {
  s.ErrorCode = &v
  return s
}

func (s *EstablishRelationTableToBusinessResponseBody) SetErrorMessage(v string) *EstablishRelationTableToBusinessResponseBody {
  s.ErrorMessage = &v
  return s
}

func (s *EstablishRelationTableToBusinessResponseBody) SetHttpStatusCode(v int32) *EstablishRelationTableToBusinessResponseBody {
  s.HttpStatusCode = &v
  return s
}

func (s *EstablishRelationTableToBusinessResponseBody) SetRequestId(v string) *EstablishRelationTableToBusinessResponseBody {
  s.RequestId = &v
  return s
}

func (s *EstablishRelationTableToBusinessResponseBody) SetSuccess(v bool) *EstablishRelationTableToBusinessResponseBody {
  s.Success = &v
  return s
}

func (s *EstablishRelationTableToBusinessResponseBody) Validate() error {
  return dara.Validate(s)
}

