// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSecureSuggestionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCalType(v string) *DescribeSecureSuggestionRequest
	GetCalType() *string
	SetLang(v string) *DescribeSecureSuggestionRequest
	GetLang() *string
	SetSource(v int32) *DescribeSecureSuggestionRequest
	GetSource() *int32
	SetSourceIp(v string) *DescribeSecureSuggestionRequest
	GetSourceIp() *string
}

type DescribeSecureSuggestionRequest struct {
	// The old or new version of the security score rule. If you set this parameter to **home_security_score**, the new version of the security score rule is returned. Otherwise, the old version of the security score rule is returned.
	//
	// example:
	//
	// home_security_score
	CalType *string `json:"CalType,omitempty" xml:"CalType,omitempty"`
	// The language of the content within the request and response. Default value: **zh**. Valid values:
	//
	// 	- **zh**: Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Source of security score, default is Cloud Security Center if left empty. Enum values:
	//
	// - 0:Cloud Security Center.
	//
	// - 1:Yaochi Console.
	//
	// example:
	//
	// 0
	Source *int32 `json:"Source,omitempty" xml:"Source,omitempty"`
	// The source IP address of the request.
	//
	// example:
	//
	// 192.168.XX.XX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DescribeSecureSuggestionRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSecureSuggestionRequest) GoString() string {
	return s.String()
}

func (s *DescribeSecureSuggestionRequest) GetCalType() *string {
	return s.CalType
}

func (s *DescribeSecureSuggestionRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeSecureSuggestionRequest) GetSource() *int32 {
	return s.Source
}

func (s *DescribeSecureSuggestionRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeSecureSuggestionRequest) SetCalType(v string) *DescribeSecureSuggestionRequest {
	s.CalType = &v
	return s
}

func (s *DescribeSecureSuggestionRequest) SetLang(v string) *DescribeSecureSuggestionRequest {
	s.Lang = &v
	return s
}

func (s *DescribeSecureSuggestionRequest) SetSource(v int32) *DescribeSecureSuggestionRequest {
	s.Source = &v
	return s
}

func (s *DescribeSecureSuggestionRequest) SetSourceIp(v string) *DescribeSecureSuggestionRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeSecureSuggestionRequest) Validate() error {
	return dara.Validate(s)
}
