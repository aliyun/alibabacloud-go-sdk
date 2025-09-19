// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeDomainDetailRequest
	GetDomainName() *string
	SetSourceIp(v string) *DescribeDomainDetailRequest
	GetSourceIp() *string
}

type DescribeDomainDetailRequest struct {
	// The domain name or the name of website that you want to query.
	//
	// >  Fuzzy match is not supported. You must enter a complete domain name or a website.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The source IP address of the request.
	//
	// example:
	//
	// 1.2.XX.XX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DescribeDomainDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainDetailRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDomainDetailRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeDomainDetailRequest) SetDomainName(v string) *DescribeDomainDetailRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainDetailRequest) SetSourceIp(v string) *DescribeDomainDetailRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeDomainDetailRequest) Validate() error {
	return dara.Validate(s)
}
