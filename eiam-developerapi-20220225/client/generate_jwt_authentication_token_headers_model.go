// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateJwtAuthenticationTokenHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GenerateJwtAuthenticationTokenHeaders
	GetCommonHeaders() map[string]*string
	SetAuthorization(v string) *GenerateJwtAuthenticationTokenHeaders
	GetAuthorization() *string
}

type GenerateJwtAuthenticationTokenHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Bearer xxxxxx
	Authorization *string `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s GenerateJwtAuthenticationTokenHeaders) String() string {
	return dara.Prettify(s)
}

func (s GenerateJwtAuthenticationTokenHeaders) GoString() string {
	return s.String()
}

func (s *GenerateJwtAuthenticationTokenHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GenerateJwtAuthenticationTokenHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *GenerateJwtAuthenticationTokenHeaders) SetCommonHeaders(v map[string]*string) *GenerateJwtAuthenticationTokenHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GenerateJwtAuthenticationTokenHeaders) SetAuthorization(v string) *GenerateJwtAuthenticationTokenHeaders {
	s.Authorization = &v
	return s
}

func (s *GenerateJwtAuthenticationTokenHeaders) Validate() error {
	return dara.Validate(s)
}
