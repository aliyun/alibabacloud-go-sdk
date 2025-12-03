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
	// example:
	//
	// home_security_score
	CalType *string `json:"CalType,omitempty" xml:"CalType,omitempty"`
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The Alibaba Cloud account ID of the member in the resource directory.
	//
	// >  You can call the [DescribeMonitorAccounts](~~DescribeMonitorAccounts~~) operation to obtain the IDs.
	//
	// example:
	//
	// 1232428423234****
	ResourceDirectoryAccountId *int64 `json:"ResourceDirectoryAccountId,omitempty" xml:"ResourceDirectoryAccountId,omitempty"`
	// example:
	//
	// 0
	Source *int32 `json:"Source,omitempty" xml:"Source,omitempty"`
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
