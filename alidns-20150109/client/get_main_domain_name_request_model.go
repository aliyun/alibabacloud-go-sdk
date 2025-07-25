// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMainDomainNameRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInputString(v string) *GetMainDomainNameRequest
	GetInputString() *string
	SetLang(v string) *GetMainDomainNameRequest
	GetLang() *string
}

type GetMainDomainNameRequest struct {
	// The string. The string can be up to 128 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// www.example.com
	InputString *string `json:"InputString,omitempty" xml:"InputString,omitempty"`
	// The language.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s GetMainDomainNameRequest) String() string {
	return dara.Prettify(s)
}

func (s GetMainDomainNameRequest) GoString() string {
	return s.String()
}

func (s *GetMainDomainNameRequest) GetInputString() *string {
	return s.InputString
}

func (s *GetMainDomainNameRequest) GetLang() *string {
	return s.Lang
}

func (s *GetMainDomainNameRequest) SetInputString(v string) *GetMainDomainNameRequest {
	s.InputString = &v
	return s
}

func (s *GetMainDomainNameRequest) SetLang(v string) *GetMainDomainNameRequest {
	s.Lang = &v
	return s
}

func (s *GetMainDomainNameRequest) Validate() error {
	return dara.Validate(s)
}
