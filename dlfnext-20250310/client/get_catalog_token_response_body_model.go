// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCatalogTokenResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetExpiresAtMillis(v int64) *GetCatalogTokenResponseBody
	GetExpiresAtMillis() *int64
	SetToken(v map[string]*string) *GetCatalogTokenResponseBody
	GetToken() map[string]*string
}

type GetCatalogTokenResponseBody struct {
	// example:
	//
	// 1749160909000
	ExpiresAtMillis *int64             `json:"expiresAtMillis,omitempty" xml:"expiresAtMillis,omitempty"`
	Token           map[string]*string `json:"token,omitempty" xml:"token,omitempty"`
}

func (s GetCatalogTokenResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetCatalogTokenResponseBody) GoString() string {
	return s.String()
}

func (s *GetCatalogTokenResponseBody) GetExpiresAtMillis() *int64 {
	return s.ExpiresAtMillis
}

func (s *GetCatalogTokenResponseBody) GetToken() map[string]*string {
	return s.Token
}

func (s *GetCatalogTokenResponseBody) SetExpiresAtMillis(v int64) *GetCatalogTokenResponseBody {
	s.ExpiresAtMillis = &v
	return s
}

func (s *GetCatalogTokenResponseBody) SetToken(v map[string]*string) *GetCatalogTokenResponseBody {
	s.Token = v
	return s
}

func (s *GetCatalogTokenResponseBody) Validate() error {
	return dara.Validate(s)
}
