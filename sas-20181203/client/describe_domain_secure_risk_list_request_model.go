// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainSecureRiskListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFrom(v string) *DescribeDomainSecureRiskListRequest
	GetFrom() *string
	SetLang(v string) *DescribeDomainSecureRiskListRequest
	GetLang() *string
	SetSourceIp(v string) *DescribeDomainSecureRiskListRequest
	GetSourceIp() *string
}

type DescribeDomainSecureRiskListRequest struct {
	// The identifier of the request source. Set the value to sas.
	//
	// example:
	//
	// sas
	From *string `json:"From,omitempty" xml:"From,omitempty"`
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
	// The source IP address.
	//
	// example:
	//
	// 124.78.*.*
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DescribeDomainSecureRiskListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainSecureRiskListRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainSecureRiskListRequest) GetFrom() *string {
	return s.From
}

func (s *DescribeDomainSecureRiskListRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeDomainSecureRiskListRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeDomainSecureRiskListRequest) SetFrom(v string) *DescribeDomainSecureRiskListRequest {
	s.From = &v
	return s
}

func (s *DescribeDomainSecureRiskListRequest) SetLang(v string) *DescribeDomainSecureRiskListRequest {
	s.Lang = &v
	return s
}

func (s *DescribeDomainSecureRiskListRequest) SetSourceIp(v string) *DescribeDomainSecureRiskListRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeDomainSecureRiskListRequest) Validate() error {
	return dara.Validate(s)
}
