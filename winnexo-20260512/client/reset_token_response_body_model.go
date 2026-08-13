// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetTokenResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ResetTokenResponseBody
	GetCode() *string
	SetMessage(v string) *ResetTokenResponseBody
	GetMessage() *string
	SetRequestId(v string) *ResetTokenResponseBody
	GetRequestId() *string
	SetToken(v string) *ResetTokenResponseBody
	GetToken() *string
	SetTokenMasked(v string) *ResetTokenResponseBody
	GetTokenMasked() *string
}

type ResetTokenResponseBody struct {
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
	// 新 Token 明文（仅本次返回，请妥善保管）
	//
	// example:
	//
	// example_token_value
	Token *string `json:"token,omitempty" xml:"token,omitempty"`
	// 脱敏后的新 Token 值
	//
	// example:
	//
	// string_value
	TokenMasked *string `json:"tokenMasked,omitempty" xml:"tokenMasked,omitempty"`
}

func (s ResetTokenResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ResetTokenResponseBody) GoString() string {
	return s.String()
}

func (s *ResetTokenResponseBody) GetCode() *string {
	return s.Code
}

func (s *ResetTokenResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ResetTokenResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ResetTokenResponseBody) GetToken() *string {
	return s.Token
}

func (s *ResetTokenResponseBody) GetTokenMasked() *string {
	return s.TokenMasked
}

func (s *ResetTokenResponseBody) SetCode(v string) *ResetTokenResponseBody {
	s.Code = &v
	return s
}

func (s *ResetTokenResponseBody) SetMessage(v string) *ResetTokenResponseBody {
	s.Message = &v
	return s
}

func (s *ResetTokenResponseBody) SetRequestId(v string) *ResetTokenResponseBody {
	s.RequestId = &v
	return s
}

func (s *ResetTokenResponseBody) SetToken(v string) *ResetTokenResponseBody {
	s.Token = &v
	return s
}

func (s *ResetTokenResponseBody) SetTokenMasked(v string) *ResetTokenResponseBody {
	s.TokenMasked = &v
	return s
}

func (s *ResetTokenResponseBody) Validate() error {
	return dara.Validate(s)
}
