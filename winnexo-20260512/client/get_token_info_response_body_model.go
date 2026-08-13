// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTokenInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetTokenInfoResponseBody
	GetCode() *string
	SetEnabled(v bool) *GetTokenInfoResponseBody
	GetEnabled() *bool
	SetGmtCreate(v string) *GetTokenInfoResponseBody
	GetGmtCreate() *string
	SetMessage(v string) *GetTokenInfoResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetTokenInfoResponseBody
	GetRequestId() *string
	SetTokenMasked(v string) *GetTokenInfoResponseBody
	GetTokenMasked() *string
}

type GetTokenInfoResponseBody struct {
	// 业务状态码：成功为 200，失败为后端错误码（ERR.	- / InvalidParameter.*）
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// Token 是否开启（存在 ACTIVE 状态的 Token）
	//
	// example:
	//
	// true
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// Token 创建时间（ISO 8601）
	//
	// example:
	//
	// string_value
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// 错误描述，成功时为空
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// 请求追踪 ID
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// 脱敏后的 Token 值
	//
	// example:
	//
	// string_value
	TokenMasked *string `json:"tokenMasked,omitempty" xml:"tokenMasked,omitempty"`
}

func (s GetTokenInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTokenInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetTokenInfoResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetTokenInfoResponseBody) GetEnabled() *bool {
	return s.Enabled
}

func (s *GetTokenInfoResponseBody) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *GetTokenInfoResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetTokenInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTokenInfoResponseBody) GetTokenMasked() *string {
	return s.TokenMasked
}

func (s *GetTokenInfoResponseBody) SetCode(v string) *GetTokenInfoResponseBody {
	s.Code = &v
	return s
}

func (s *GetTokenInfoResponseBody) SetEnabled(v bool) *GetTokenInfoResponseBody {
	s.Enabled = &v
	return s
}

func (s *GetTokenInfoResponseBody) SetGmtCreate(v string) *GetTokenInfoResponseBody {
	s.GmtCreate = &v
	return s
}

func (s *GetTokenInfoResponseBody) SetMessage(v string) *GetTokenInfoResponseBody {
	s.Message = &v
	return s
}

func (s *GetTokenInfoResponseBody) SetRequestId(v string) *GetTokenInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTokenInfoResponseBody) SetTokenMasked(v string) *GetTokenInfoResponseBody {
	s.TokenMasked = &v
	return s
}

func (s *GetTokenInfoResponseBody) Validate() error {
	return dara.Validate(s)
}
