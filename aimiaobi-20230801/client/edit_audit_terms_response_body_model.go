// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEditAuditTermsResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetCode(v string) *EditAuditTermsResponseBody
  GetCode() *string 
  SetData(v bool) *EditAuditTermsResponseBody
  GetData() *bool 
  SetHttpStatusCode(v int32) *EditAuditTermsResponseBody
  GetHttpStatusCode() *int32 
  SetMessage(v string) *EditAuditTermsResponseBody
  GetMessage() *string 
  SetRequestId(v string) *EditAuditTermsResponseBody
  GetRequestId() *string 
  SetSuccess(v bool) *EditAuditTermsResponseBody
  GetSuccess() *bool 
}

type EditAuditTermsResponseBody struct {
  // example:
  // 
  // DataNotExists
  Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
  // example:
  // 
  // true
  Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
  // example:
  // 
  // 200
  HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
  // example:
  // 
  // success
  Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
  // Id of the request
  // 
  // example:
  // 
  // F2F366D6-E9FE-1006-BB70-2C650896AAB5
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
  // example:
  // 
  // true
  Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s EditAuditTermsResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EditAuditTermsResponseBody) GoString() string {
  return s.String()
}

func (s *EditAuditTermsResponseBody) GetCode() *string  {
  return s.Code
}

func (s *EditAuditTermsResponseBody) GetData() *bool  {
  return s.Data
}

func (s *EditAuditTermsResponseBody) GetHttpStatusCode() *int32  {
  return s.HttpStatusCode
}

func (s *EditAuditTermsResponseBody) GetMessage() *string  {
  return s.Message
}

func (s *EditAuditTermsResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EditAuditTermsResponseBody) GetSuccess() *bool  {
  return s.Success
}

func (s *EditAuditTermsResponseBody) SetCode(v string) *EditAuditTermsResponseBody {
  s.Code = &v
  return s
}

func (s *EditAuditTermsResponseBody) SetData(v bool) *EditAuditTermsResponseBody {
  s.Data = &v
  return s
}

func (s *EditAuditTermsResponseBody) SetHttpStatusCode(v int32) *EditAuditTermsResponseBody {
  s.HttpStatusCode = &v
  return s
}

func (s *EditAuditTermsResponseBody) SetMessage(v string) *EditAuditTermsResponseBody {
  s.Message = &v
  return s
}

func (s *EditAuditTermsResponseBody) SetRequestId(v string) *EditAuditTermsResponseBody {
  s.RequestId = &v
  return s
}

func (s *EditAuditTermsResponseBody) SetSuccess(v bool) *EditAuditTermsResponseBody {
  s.Success = &v
  return s
}

func (s *EditAuditTermsResponseBody) Validate() error {
  return dara.Validate(s)
}

