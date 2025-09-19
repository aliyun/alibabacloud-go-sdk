// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainSecureScoreRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeDomainSecureScoreRequest
	GetLang() *string
	SetSourceIp(v string) *DescribeDomainSecureScoreRequest
	GetSourceIp() *string
}

type DescribeDomainSecureScoreRequest struct {
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
	// The source IP address of the request.
	//
	// example:
	//
	// 123.113.*.*
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DescribeDomainSecureScoreRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainSecureScoreRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainSecureScoreRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeDomainSecureScoreRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeDomainSecureScoreRequest) SetLang(v string) *DescribeDomainSecureScoreRequest {
	s.Lang = &v
	return s
}

func (s *DescribeDomainSecureScoreRequest) SetSourceIp(v string) *DescribeDomainSecureScoreRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeDomainSecureScoreRequest) Validate() error {
	return dara.Validate(s)
}
