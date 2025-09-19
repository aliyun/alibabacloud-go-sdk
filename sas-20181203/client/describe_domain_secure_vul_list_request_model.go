// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainSecureVulListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFrom(v string) *DescribeDomainSecureVulListRequest
	GetFrom() *string
	SetLang(v string) *DescribeDomainSecureVulListRequest
	GetLang() *string
	SetSourceIp(v string) *DescribeDomainSecureVulListRequest
	GetSourceIp() *string
	SetType(v string) *DescribeDomainSecureVulListRequest
	GetType() *string
}

type DescribeDomainSecureVulListRequest struct {
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
	// 111.196.*.*
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// The vulnerability type. Valid values:
	//
	// 	- **app**: application vulnerability.
	//
	// example:
	//
	// app
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeDomainSecureVulListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainSecureVulListRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainSecureVulListRequest) GetFrom() *string {
	return s.From
}

func (s *DescribeDomainSecureVulListRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeDomainSecureVulListRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeDomainSecureVulListRequest) GetType() *string {
	return s.Type
}

func (s *DescribeDomainSecureVulListRequest) SetFrom(v string) *DescribeDomainSecureVulListRequest {
	s.From = &v
	return s
}

func (s *DescribeDomainSecureVulListRequest) SetLang(v string) *DescribeDomainSecureVulListRequest {
	s.Lang = &v
	return s
}

func (s *DescribeDomainSecureVulListRequest) SetSourceIp(v string) *DescribeDomainSecureVulListRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeDomainSecureVulListRequest) SetType(v string) *DescribeDomainSecureVulListRequest {
	s.Type = &v
	return s
}

func (s *DescribeDomainSecureVulListRequest) Validate() error {
	return dara.Validate(s)
}
