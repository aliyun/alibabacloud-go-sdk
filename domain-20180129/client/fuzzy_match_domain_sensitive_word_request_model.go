// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFuzzyMatchDomainSensitiveWordRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKeyword(v string) *FuzzyMatchDomainSensitiveWordRequest
	GetKeyword() *string
	SetLang(v string) *FuzzyMatchDomainSensitiveWordRequest
	GetLang() *string
	SetUserClientIp(v string) *FuzzyMatchDomainSensitiveWordRequest
	GetUserClientIp() *string
}

type FuzzyMatchDomainSensitiveWordRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// xxx**.cn
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// 127.0.0.1
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s FuzzyMatchDomainSensitiveWordRequest) String() string {
	return dara.Prettify(s)
}

func (s FuzzyMatchDomainSensitiveWordRequest) GoString() string {
	return s.String()
}

func (s *FuzzyMatchDomainSensitiveWordRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *FuzzyMatchDomainSensitiveWordRequest) GetLang() *string {
	return s.Lang
}

func (s *FuzzyMatchDomainSensitiveWordRequest) GetUserClientIp() *string {
	return s.UserClientIp
}

func (s *FuzzyMatchDomainSensitiveWordRequest) SetKeyword(v string) *FuzzyMatchDomainSensitiveWordRequest {
	s.Keyword = &v
	return s
}

func (s *FuzzyMatchDomainSensitiveWordRequest) SetLang(v string) *FuzzyMatchDomainSensitiveWordRequest {
	s.Lang = &v
	return s
}

func (s *FuzzyMatchDomainSensitiveWordRequest) SetUserClientIp(v string) *FuzzyMatchDomainSensitiveWordRequest {
	s.UserClientIp = &v
	return s
}

func (s *FuzzyMatchDomainSensitiveWordRequest) Validate() error {
	return dara.Validate(s)
}
