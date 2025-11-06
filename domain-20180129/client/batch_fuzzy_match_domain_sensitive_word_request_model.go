// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchFuzzyMatchDomainSensitiveWordRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKeyword(v string) *BatchFuzzyMatchDomainSensitiveWordRequest
	GetKeyword() *string
	SetLang(v string) *BatchFuzzyMatchDomainSensitiveWordRequest
	GetLang() *string
	SetUserClientIp(v string) *BatchFuzzyMatchDomainSensitiveWordRequest
	GetUserClientIp() *string
}

type BatchFuzzyMatchDomainSensitiveWordRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// example.com,aliyundoc.com
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

func (s BatchFuzzyMatchDomainSensitiveWordRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchFuzzyMatchDomainSensitiveWordRequest) GoString() string {
	return s.String()
}

func (s *BatchFuzzyMatchDomainSensitiveWordRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *BatchFuzzyMatchDomainSensitiveWordRequest) GetLang() *string {
	return s.Lang
}

func (s *BatchFuzzyMatchDomainSensitiveWordRequest) GetUserClientIp() *string {
	return s.UserClientIp
}

func (s *BatchFuzzyMatchDomainSensitiveWordRequest) SetKeyword(v string) *BatchFuzzyMatchDomainSensitiveWordRequest {
	s.Keyword = &v
	return s
}

func (s *BatchFuzzyMatchDomainSensitiveWordRequest) SetLang(v string) *BatchFuzzyMatchDomainSensitiveWordRequest {
	s.Lang = &v
	return s
}

func (s *BatchFuzzyMatchDomainSensitiveWordRequest) SetUserClientIp(v string) *BatchFuzzyMatchDomainSensitiveWordRequest {
	s.UserClientIp = &v
	return s
}

func (s *BatchFuzzyMatchDomainSensitiveWordRequest) Validate() error {
	return dara.Validate(s)
}
