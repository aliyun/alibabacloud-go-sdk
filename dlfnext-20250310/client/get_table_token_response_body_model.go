// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTableTokenResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetExpiresAtMillis(v int64) *GetTableTokenResponseBody
	GetExpiresAtMillis() *int64
	SetToken(v map[string]*string) *GetTableTokenResponseBody
	GetToken() map[string]*string
}

type GetTableTokenResponseBody struct {
	// example:
	//
	// 1749160909000
	ExpiresAtMillis *int64             `json:"expiresAtMillis,omitempty" xml:"expiresAtMillis,omitempty"`
	Token           map[string]*string `json:"token,omitempty" xml:"token,omitempty"`
}

func (s GetTableTokenResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTableTokenResponseBody) GoString() string {
	return s.String()
}

func (s *GetTableTokenResponseBody) GetExpiresAtMillis() *int64 {
	return s.ExpiresAtMillis
}

func (s *GetTableTokenResponseBody) GetToken() map[string]*string {
	return s.Token
}

func (s *GetTableTokenResponseBody) SetExpiresAtMillis(v int64) *GetTableTokenResponseBody {
	s.ExpiresAtMillis = &v
	return s
}

func (s *GetTableTokenResponseBody) SetToken(v map[string]*string) *GetTableTokenResponseBody {
	s.Token = v
	return s
}

func (s *GetTableTokenResponseBody) Validate() error {
	return dara.Validate(s)
}
