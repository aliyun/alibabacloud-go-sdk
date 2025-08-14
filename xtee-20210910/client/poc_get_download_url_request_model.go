// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPocGetDownloadUrlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *PocGetDownloadUrlRequest
	GetLang() *string
	SetToken(v string) *PocGetDownloadUrlRequest
	GetToken() *string
}

type PocGetDownloadUrlRequest struct {
	// Sets the language type for requests and received messages, with a default value of **zh**. Values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Task token.
	//
	// example:
	//
	// d83221f51752805880
	Token *string `json:"Token,omitempty" xml:"Token,omitempty"`
}

func (s PocGetDownloadUrlRequest) String() string {
	return dara.Prettify(s)
}

func (s PocGetDownloadUrlRequest) GoString() string {
	return s.String()
}

func (s *PocGetDownloadUrlRequest) GetLang() *string {
	return s.Lang
}

func (s *PocGetDownloadUrlRequest) GetToken() *string {
	return s.Token
}

func (s *PocGetDownloadUrlRequest) SetLang(v string) *PocGetDownloadUrlRequest {
	s.Lang = &v
	return s
}

func (s *PocGetDownloadUrlRequest) SetToken(v string) *PocGetDownloadUrlRequest {
	s.Token = &v
	return s
}

func (s *PocGetDownloadUrlRequest) Validate() error {
	return dara.Validate(s)
}
