// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetProductListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLanguage(v string) *GetProductListRequest
	GetLanguage() *string
	SetToken(v string) *GetProductListRequest
	GetToken() *string
}

type GetProductListRequest struct {
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// example:
	//
	// ****
	Token *string `json:"Token,omitempty" xml:"Token,omitempty"`
}

func (s GetProductListRequest) String() string {
	return dara.Prettify(s)
}

func (s GetProductListRequest) GoString() string {
	return s.String()
}

func (s *GetProductListRequest) GetLanguage() *string {
	return s.Language
}

func (s *GetProductListRequest) GetToken() *string {
	return s.Token
}

func (s *GetProductListRequest) SetLanguage(v string) *GetProductListRequest {
	s.Language = &v
	return s
}

func (s *GetProductListRequest) SetToken(v string) *GetProductListRequest {
	s.Token = &v
	return s
}

func (s *GetProductListRequest) Validate() error {
	return dara.Validate(s)
}
