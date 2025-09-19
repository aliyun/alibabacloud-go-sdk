// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainSecureSuggestsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeDomainSecureSuggestsRequest
	GetLang() *string
	SetSourceIp(v string) *DescribeDomainSecureSuggestsRequest
	GetSourceIp() *string
}

type DescribeDomainSecureSuggestsRequest struct {
	// Sets the language type for requests and received messages, default is **zh**. Values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The IP address of the access source.
	//
	// example:
	//
	// 218.249.*.*
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DescribeDomainSecureSuggestsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainSecureSuggestsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainSecureSuggestsRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeDomainSecureSuggestsRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeDomainSecureSuggestsRequest) SetLang(v string) *DescribeDomainSecureSuggestsRequest {
	s.Lang = &v
	return s
}

func (s *DescribeDomainSecureSuggestsRequest) SetSourceIp(v string) *DescribeDomainSecureSuggestsRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeDomainSecureSuggestsRequest) Validate() error {
	return dara.Validate(s)
}
