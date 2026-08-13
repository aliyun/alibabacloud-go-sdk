// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableTokenResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetCode(v string) *EnableTokenResponseBody
  GetCode() *string 
  SetMessage(v string) *EnableTokenResponseBody
  GetMessage() *string 
  SetRequestId(v string) *EnableTokenResponseBody
  GetRequestId() *string 
  SetToken(v string) *EnableTokenResponseBody
  GetToken() *string 
  SetTokenMasked(v string) *EnableTokenResponseBody
  GetTokenMasked() *string 
}

type EnableTokenResponseBody struct {
  // 业务状态码：成功为 200，失败为后端错误码（ERR.	- / InvalidParameter.*）
  // 
  // example:
  // 
  // 200
  Code *string `json:"code,omitempty" xml:"code,omitempty"`
  // 错误描述，成功时为空
  Message *string `json:"message,omitempty" xml:"message,omitempty"`
  // 请求追踪 ID
  // 
  // example:
  // 
  // 019FF406-1B10-0065-A97D-2D1920C2A03D
  RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
  // Token 明文（仅本次返回，请妥善保管）
  // 
  // example:
  // 
  // example_token_value
  Token *string `json:"token,omitempty" xml:"token,omitempty"`
  // 脱敏后的 Token 值
  // 
  // example:
  // 
  // string_value
  TokenMasked *string `json:"tokenMasked,omitempty" xml:"tokenMasked,omitempty"`
}

func (s EnableTokenResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnableTokenResponseBody) GoString() string {
  return s.String()
}

func (s *EnableTokenResponseBody) GetCode() *string  {
  return s.Code
}

func (s *EnableTokenResponseBody) GetMessage() *string  {
  return s.Message
}

func (s *EnableTokenResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnableTokenResponseBody) GetToken() *string  {
  return s.Token
}

func (s *EnableTokenResponseBody) GetTokenMasked() *string  {
  return s.TokenMasked
}

func (s *EnableTokenResponseBody) SetCode(v string) *EnableTokenResponseBody {
  s.Code = &v
  return s
}

func (s *EnableTokenResponseBody) SetMessage(v string) *EnableTokenResponseBody {
  s.Message = &v
  return s
}

func (s *EnableTokenResponseBody) SetRequestId(v string) *EnableTokenResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnableTokenResponseBody) SetToken(v string) *EnableTokenResponseBody {
  s.Token = &v
  return s
}

func (s *EnableTokenResponseBody) SetTokenMasked(v string) *EnableTokenResponseBody {
  s.TokenMasked = &v
  return s
}

func (s *EnableTokenResponseBody) Validate() error {
  return dara.Validate(s)
}

