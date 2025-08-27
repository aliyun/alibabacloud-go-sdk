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
  // example:
  // 
  // SUCCESS
  Code *string `json:"code,omitempty" xml:"code,omitempty"`
  // example:
  // 
  // 200
  HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
  Message *string `json:"message,omitempty" xml:"message,omitempty"`
  Module *ExternalUserQueryResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
  // example:
  // 
  // 407543AF-2BD9-5890-BD92-9D1AB7218B27
  RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
  // example:
  // 
  // true
  Success *bool `json:"success,omitempty" xml:"success,omitempty"`
  // traceId
  // 
  // example:
  // 
  // 210e847f16611516748613869de4f6
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
  return dara.Validate(s)
}

type ExternalUserQueryResponseBodyModule struct {
  // example:
  // 
  // 2000-01-02
  Birthday *string `json:"birthday,omitempty" xml:"birthday,omitempty"`
  // example:
  // 
  // btrip123456
  CorpId *string `json:"corp_id,omitempty" xml:"corp_id,omitempty"`
  // example:
  // 
  // zhangsan@alibaba-inc.com
  Email *string `json:"email,omitempty" xml:"email,omitempty"`
  // example:
  // 
  // 000001
  ExternalUserId *string `json:"external_user_id,omitempty" xml:"external_user_id,omitempty"`
  // example:
  // 
  // 13438009765
  Phone *string `json:"phone,omitempty" xml:"phone,omitempty"`
  RealName *string `json:"real_name,omitempty" xml:"real_name,omitempty"`
  // example:
  // 
  // zhang/san
  RealNameEn *string `json:"real_name_en,omitempty" xml:"real_name_en,omitempty"`
  // example:
  // 
  // e1$12345678
  UserId *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
  UserNick *string `json:"user_nick,omitempty" xml:"user_nick,omitempty"`
  // example:
  // 
  // 1
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

