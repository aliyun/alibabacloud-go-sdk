// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainResolveRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomain(v string) *DescribeDomainResolveRequest
	GetDomain() *string
	SetIpVersion(v string) *DescribeDomainResolveRequest
	GetIpVersion() *string
	SetLang(v string) *DescribeDomainResolveRequest
	GetLang() *string
	SetSourceIp(v string) *DescribeDomainResolveRequest
	GetSourceIp() *string
}

type DescribeDomainResolveRequest struct {
	// The domain name whose DNS record you want to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// www.aliyundoc.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The IP version of the asset that is protected by Cloud Firewall. Valid values:
	//
	// 	- **4**: IPv4 (default)
	//
	// 	- **6**: IPv6
	//
	// example:
	//
	// 6
	IpVersion *string `json:"IpVersion,omitempty" xml:"IpVersion,omitempty"`
	// The language of the content within the response.
	//
	// Valid values:
	//
	// 	- **zh**: Chinese (default)
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Deprecated
	//
	// The source IP address of the request.
	//
	// example:
	//
	// 192.0.XX.XX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DescribeDomainResolveRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainResolveRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainResolveRequest) GetDomain() *string {
	return s.Domain
}

func (s *DescribeDomainResolveRequest) GetIpVersion() *string {
	return s.IpVersion
}

func (s *DescribeDomainResolveRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeDomainResolveRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeDomainResolveRequest) SetDomain(v string) *DescribeDomainResolveRequest {
	s.Domain = &v
	return s
}

func (s *DescribeDomainResolveRequest) SetIpVersion(v string) *DescribeDomainResolveRequest {
	s.IpVersion = &v
	return s
}

func (s *DescribeDomainResolveRequest) SetLang(v string) *DescribeDomainResolveRequest {
	s.Lang = &v
	return s
}

func (s *DescribeDomainResolveRequest) SetSourceIp(v string) *DescribeDomainResolveRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeDomainResolveRequest) Validate() error {
	return dara.Validate(s)
}
