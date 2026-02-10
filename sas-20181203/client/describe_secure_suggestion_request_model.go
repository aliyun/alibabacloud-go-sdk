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
	SetResourceDirectoryAccountId(v int64) *DescribeSecureSuggestionRequest
	GetResourceDirectoryAccountId() *int64
	SetSource(v int32) *DescribeSecureSuggestionRequest
	GetSource() *int32
	SetSourceIp(v string) *DescribeSecureSuggestionRequest
	GetSourceIp() *string
}

type DescribeSecureSuggestionRequest struct {
	// Choose to query the new or old version of the security score rules. When the value is **home_security_score**, it queries the new version of the security score rules; otherwise, it defaults to querying the old version of the security score rules.
	//
	// example:
	//
	// home_security_score
	CalType *string `json:"CalType,omitempty" xml:"CalType,omitempty"`
	// The language type for request and response messages, default is **zh**. Values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Resource directory member account ID (Alibaba Cloud account).
	//
	// > You can obtain this parameter by calling the [DescribeMonitorAccounts](~~DescribeMonitorAccounts~~) API.
	//
	// example:
	//
	// 1232428423234****
	ResourceDirectoryAccountId *int64 `json:"ResourceDirectoryAccountId,omitempty" xml:"ResourceDirectoryAccountId,omitempty"`
	// Source of the security score. If left empty, it defaults to Cloud Security Center. Enumerated values:
	//
	// - 0: Cloud Security Center.
	//
	// - 1: Yaochi Console.
	//
	// example:
	//
	// 0
	Source *int32 `json:"Source,omitempty" xml:"Source,omitempty"`
	// The IP address of the access source.
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

func (s *DescribeSecureSuggestionRequest) GetResourceDirectoryAccountId() *int64 {
	return s.ResourceDirectoryAccountId
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

func (s *DescribeSecureSuggestionRequest) SetResourceDirectoryAccountId(v int64) *DescribeSecureSuggestionRequest {
	s.ResourceDirectoryAccountId = &v
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
