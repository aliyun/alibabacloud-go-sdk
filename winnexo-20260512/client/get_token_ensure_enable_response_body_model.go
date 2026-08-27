// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTokenEnsureEnableResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetTokenEnsureEnableResponseBody
	GetCode() *string
	SetGmtCreate(v string) *GetTokenEnsureEnableResponseBody
	GetGmtCreate() *string
	SetMessage(v string) *GetTokenEnsureEnableResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetTokenEnsureEnableResponseBody
	GetRequestId() *string
	SetToken(v string) *GetTokenEnsureEnableResponseBody
	GetToken() *string
	SetTokenMasked(v string) *GetTokenEnsureEnableResponseBody
	GetTokenMasked() *string
}

type GetTokenEnsureEnableResponseBody struct {
	// The status code.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2026-08-25T10:00:00+08:00
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// The description of the status code.
	//
	// example:
	//
	// successful
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The temporary access credential for the data catalog.
	//
	// example:
	//
	// example_token_value
	Token *string `json:"token,omitempty" xml:"token,omitempty"`
	// The masked token value.
	//
	// example:
	//
	// string_value
	TokenMasked *string `json:"tokenMasked,omitempty" xml:"tokenMasked,omitempty"`
}

func (s GetTokenEnsureEnableResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTokenEnsureEnableResponseBody) GoString() string {
	return s.String()
}

func (s *GetTokenEnsureEnableResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetTokenEnsureEnableResponseBody) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *GetTokenEnsureEnableResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetTokenEnsureEnableResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTokenEnsureEnableResponseBody) GetToken() *string {
	return s.Token
}

func (s *GetTokenEnsureEnableResponseBody) GetTokenMasked() *string {
	return s.TokenMasked
}

func (s *GetTokenEnsureEnableResponseBody) SetCode(v string) *GetTokenEnsureEnableResponseBody {
	s.Code = &v
	return s
}

func (s *GetTokenEnsureEnableResponseBody) SetGmtCreate(v string) *GetTokenEnsureEnableResponseBody {
	s.GmtCreate = &v
	return s
}

func (s *GetTokenEnsureEnableResponseBody) SetMessage(v string) *GetTokenEnsureEnableResponseBody {
	s.Message = &v
	return s
}

func (s *GetTokenEnsureEnableResponseBody) SetRequestId(v string) *GetTokenEnsureEnableResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTokenEnsureEnableResponseBody) SetToken(v string) *GetTokenEnsureEnableResponseBody {
	s.Token = &v
	return s
}

func (s *GetTokenEnsureEnableResponseBody) SetTokenMasked(v string) *GetTokenEnsureEnableResponseBody {
	s.TokenMasked = &v
	return s
}

func (s *GetTokenEnsureEnableResponseBody) Validate() error {
	return dara.Validate(s)
}
