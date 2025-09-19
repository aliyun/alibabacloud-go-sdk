// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeUserLangRequest interface {
	dara.Model
	String() string
	GoString() string
	SetUserLang(v string) *ChangeUserLangRequest
	GetUserLang() *string
}

type ChangeUserLangRequest struct {
	// The new language. Valid values:
	//
	// 	- zh: Chinese
	//
	// 	- en: English
	//
	// example:
	//
	// zh
	UserLang *string `json:"UserLang,omitempty" xml:"UserLang,omitempty"`
}

func (s ChangeUserLangRequest) String() string {
	return dara.Prettify(s)
}

func (s ChangeUserLangRequest) GoString() string {
	return s.String()
}

func (s *ChangeUserLangRequest) GetUserLang() *string {
	return s.UserLang
}

func (s *ChangeUserLangRequest) SetUserLang(v string) *ChangeUserLangRequest {
	s.UserLang = &v
	return s
}

func (s *ChangeUserLangRequest) Validate() error {
	return dara.Validate(s)
}
