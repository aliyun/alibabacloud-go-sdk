// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExternalUserQueryResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetCode(v string) *ExternalUserQueryResponseBody
  GetCode() *string 
  SetHttpStatusCode(v int32) *ExternalUserQueryResponseBody
  GetHttpStatusCode() *int32 
  SetMessage(v string) *ExternalUserQueryResponseBody
  GetMessage() *string 
  SetModule(v *ExternalUserQueryResponseBodyModule) *ExternalUserQueryResponseBody
  GetModule() *ExternalUserQueryResponseBodyModule 
  SetRequestId(v string) *ExternalUserQueryResponseBody
  GetRequestId() *string 
  SetSuccess(v bool) *ExternalUserQueryResponseBody
  GetSuccess() *bool 
  SetTraceId(v string) *ExternalUserQueryResponseBody
  GetTraceId() *string 
}

type ExternalUserQueryResponseBody struct {
  Code *string `json:"code,omitempty" xml:"code,omitempty"`
  HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
  Message *string `json:"message,omitempty" xml:"message,omitempty"`
  Module *ExternalUserQueryResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
  RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
  Success *bool `json:"success,omitempty" xml:"success,omitempty"`
  TraceId *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s ExternalUserQueryResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ExternalUserQueryResponseBody) GoString() string {
  return s.String()
}

func (s *ExternalUserQueryResponseBody) GetCode() *string  {
  return s.Code
}

func (s *ExternalUserQueryResponseBody) GetHttpStatusCode() *int32  {
  return s.HttpStatusCode
}

func (s *ExternalUserQueryResponseBody) GetMessage() *string  {
  return s.Message
}

func (s *ExternalUserQueryResponseBody) GetModule() *ExternalUserQueryResponseBodyModule  {
  return s.Module
}

func (s *ExternalUserQueryResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *ExternalUserQueryResponseBody) GetSuccess() *bool  {
  return s.Success
}

func (s *ExternalUserQueryResponseBody) GetTraceId() *string  {
  return s.TraceId
}

func (s *ExternalUserQueryResponseBody) SetCode(v string) *ExternalUserQueryResponseBody {
  s.Code = &v
  return s
}

func (s *ExternalUserQueryResponseBody) SetHttpStatusCode(v int32) *ExternalUserQueryResponseBody {
  s.HttpStatusCode = &v
  return s
}

func (s *ExternalUserQueryResponseBody) SetMessage(v string) *ExternalUserQueryResponseBody {
  s.Message = &v
  return s
}

func (s *ExternalUserQueryResponseBody) SetModule(v *ExternalUserQueryResponseBodyModule) *ExternalUserQueryResponseBody {
  s.Module = v
  return s
}

func (s *ExternalUserQueryResponseBody) SetRequestId(v string) *ExternalUserQueryResponseBody {
  s.RequestId = &v
  return s
}

func (s *ExternalUserQueryResponseBody) SetSuccess(v bool) *ExternalUserQueryResponseBody {
  s.Success = &v
  return s
}

func (s *ExternalUserQueryResponseBody) SetTraceId(v string) *ExternalUserQueryResponseBody {
  s.TraceId = &v
  return s
}

func (s *ExternalUserQueryResponseBody) Validate() error {
  if s.Module != nil {
    if err := s.Module.Validate(); err != nil {
      return err
    }
  }
  return nil
}

type ExternalUserQueryResponseBodyModule struct {
  Birthday *string `json:"birthday,omitempty" xml:"birthday,omitempty"`
  CorpId *string `json:"corp_id,omitempty" xml:"corp_id,omitempty"`
  Email *string `json:"email,omitempty" xml:"email,omitempty"`
  ExternalUserId *string `json:"external_user_id,omitempty" xml:"external_user_id,omitempty"`
  Phone *string `json:"phone,omitempty" xml:"phone,omitempty"`
  RealName *string `json:"real_name,omitempty" xml:"real_name,omitempty"`
  RealNameEn *string `json:"real_name_en,omitempty" xml:"real_name_en,omitempty"`
  UserId *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
  UserNick *string `json:"user_nick,omitempty" xml:"user_nick,omitempty"`
  UserType *int32 `json:"user_type,omitempty" xml:"user_type,omitempty"`
}

func (s ExternalUserQueryResponseBodyModule) String() string {
  return dara.Prettify(s)
}

func (s ExternalUserQueryResponseBodyModule) GoString() string {
  return s.String()
}

func (s *ExternalUserQueryResponseBodyModule) GetBirthday() *string  {
  return s.Birthday
}

func (s *ExternalUserQueryResponseBodyModule) GetCorpId() *string  {
  return s.CorpId
}

func (s *ExternalUserQueryResponseBodyModule) GetEmail() *string  {
  return s.Email
}

func (s *ExternalUserQueryResponseBodyModule) GetExternalUserId() *string  {
  return s.ExternalUserId
}

func (s *ExternalUserQueryResponseBodyModule) GetPhone() *string  {
  return s.Phone
}

func (s *ExternalUserQueryResponseBodyModule) GetRealName() *string  {
  return s.RealName
}

func (s *ExternalUserQueryResponseBodyModule) GetRealNameEn() *string  {
  return s.RealNameEn
}

func (s *ExternalUserQueryResponseBodyModule) GetUserId() *string  {
  return s.UserId
}

func (s *ExternalUserQueryResponseBodyModule) GetUserNick() *string  {
  return s.UserNick
}

func (s *ExternalUserQueryResponseBodyModule) GetUserType() *int32  {
  return s.UserType
}

func (s *ExternalUserQueryResponseBodyModule) SetBirthday(v string) *ExternalUserQueryResponseBodyModule {
  s.Birthday = &v
  return s
}

func (s *ExternalUserQueryResponseBodyModule) SetCorpId(v string) *ExternalUserQueryResponseBodyModule {
  s.CorpId = &v
  return s
}

func (s *ExternalUserQueryResponseBodyModule) SetEmail(v string) *ExternalUserQueryResponseBodyModule {
  s.Email = &v
  return s
}

func (s *ExternalUserQueryResponseBodyModule) SetExternalUserId(v string) *ExternalUserQueryResponseBodyModule {
  s.ExternalUserId = &v
  return s
}

func (s *ExternalUserQueryResponseBodyModule) SetPhone(v string) *ExternalUserQueryResponseBodyModule {
  s.Phone = &v
  return s
}

func (s *ExternalUserQueryResponseBodyModule) SetRealName(v string) *ExternalUserQueryResponseBodyModule {
  s.RealName = &v
  return s
}

func (s *ExternalUserQueryResponseBodyModule) SetRealNameEn(v string) *ExternalUserQueryResponseBodyModule {
  s.RealNameEn = &v
  return s
}

func (s *ExternalUserQueryResponseBodyModule) SetUserId(v string) *ExternalUserQueryResponseBodyModule {
  s.UserId = &v
  return s
}

func (s *ExternalUserQueryResponseBodyModule) SetUserNick(v string) *ExternalUserQueryResponseBodyModule {
  s.UserNick = &v
  return s
}

func (s *ExternalUserQueryResponseBodyModule) SetUserType(v int32) *ExternalUserQueryResponseBodyModule {
  s.UserType = &v
  return s
}

func (s *ExternalUserQueryResponseBodyModule) Validate() error {
  return dara.Validate(s)
}

