// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainH2FingerprintResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDomainH2Fp(v []*DescribeDomainH2FingerprintResponseBodyDomainH2Fp) *DescribeDomainH2FingerprintResponseBody
	GetDomainH2Fp() []*DescribeDomainH2FingerprintResponseBodyDomainH2Fp
	SetRequestId(v string) *DescribeDomainH2FingerprintResponseBody
	GetRequestId() *string
}

type DescribeDomainH2FingerprintResponseBody struct {
	// The list of top N HTTP/2.0 fingerprints for the domain name.
	DomainH2Fp []*DescribeDomainH2FingerprintResponseBodyDomainH2Fp `json:"DomainH2Fp,omitempty" xml:"DomainH2Fp,omitempty" type:"Repeated"`
	// The ID of the request. The ID is a unique identifier that Alibaba Cloud generates for the request and can be used to troubleshoot issues.
	//
	// example:
	//
	// 112777CC-2AD6-46FC-A263-00B931406FCD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDomainH2FingerprintResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainH2FingerprintResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainH2FingerprintResponseBody) GetDomainH2Fp() []*DescribeDomainH2FingerprintResponseBodyDomainH2Fp {
	return s.DomainH2Fp
}

func (s *DescribeDomainH2FingerprintResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDomainH2FingerprintResponseBody) SetDomainH2Fp(v []*DescribeDomainH2FingerprintResponseBodyDomainH2Fp) *DescribeDomainH2FingerprintResponseBody {
	s.DomainH2Fp = v
	return s
}

func (s *DescribeDomainH2FingerprintResponseBody) SetRequestId(v string) *DescribeDomainH2FingerprintResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainH2FingerprintResponseBody) Validate() error {
	if s.DomainH2Fp != nil {
		for _, item := range s.DomainH2Fp {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDomainH2FingerprintResponseBodyDomainH2Fp struct {
	// The domain name of the website.
	//
	// example:
	//
	// example.aliyundoc.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The HTTP/2.0 fingerprint.
	//
	// example:
	//
	// 532501bc316d02c8b1a007db76f2c796
	H2Fingerprint *string `json:"H2Fingerprint,omitempty" xml:"H2Fingerprint,omitempty"`
	// The number of times the HTTP/2.0 fingerprint is matched for the domain name.
	//
	// example:
	//
	// 471755
	Pv *int64 `json:"Pv,omitempty" xml:"Pv,omitempty"`
}

func (s DescribeDomainH2FingerprintResponseBodyDomainH2Fp) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainH2FingerprintResponseBodyDomainH2Fp) GoString() string {
	return s.String()
}

func (s *DescribeDomainH2FingerprintResponseBodyDomainH2Fp) GetDomain() *string {
	return s.Domain
}

func (s *DescribeDomainH2FingerprintResponseBodyDomainH2Fp) GetH2Fingerprint() *string {
	return s.H2Fingerprint
}

func (s *DescribeDomainH2FingerprintResponseBodyDomainH2Fp) GetPv() *int64 {
	return s.Pv
}

func (s *DescribeDomainH2FingerprintResponseBodyDomainH2Fp) SetDomain(v string) *DescribeDomainH2FingerprintResponseBodyDomainH2Fp {
	s.Domain = &v
	return s
}

func (s *DescribeDomainH2FingerprintResponseBodyDomainH2Fp) SetH2Fingerprint(v string) *DescribeDomainH2FingerprintResponseBodyDomainH2Fp {
	s.H2Fingerprint = &v
	return s
}

func (s *DescribeDomainH2FingerprintResponseBodyDomainH2Fp) SetPv(v int64) *DescribeDomainH2FingerprintResponseBodyDomainH2Fp {
	s.Pv = &v
	return s
}

func (s *DescribeDomainH2FingerprintResponseBodyDomainH2Fp) Validate() error {
	return dara.Validate(s)
}
