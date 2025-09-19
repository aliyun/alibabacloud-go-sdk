// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainSecureStatisticsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeDomainSecureStatisticsRequest
	GetLang() *string
	SetSourceIp(v string) *DescribeDomainSecureStatisticsRequest
	GetSourceIp() *string
}

type DescribeDomainSecureStatisticsRequest struct {
	// The language of the content within the request and response. Valid values: Default value: **zh**. Valid values:
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
	// 113.87.*.*
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DescribeDomainSecureStatisticsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainSecureStatisticsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainSecureStatisticsRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeDomainSecureStatisticsRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeDomainSecureStatisticsRequest) SetLang(v string) *DescribeDomainSecureStatisticsRequest {
	s.Lang = &v
	return s
}

func (s *DescribeDomainSecureStatisticsRequest) SetSourceIp(v string) *DescribeDomainSecureStatisticsRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeDomainSecureStatisticsRequest) Validate() error {
	return dara.Validate(s)
}
