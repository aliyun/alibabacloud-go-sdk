// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainTopFingerprintResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDomainTopFp(v []*DescribeDomainTopFingerprintResponseBodyDomainTopFp) *DescribeDomainTopFingerprintResponseBody
	GetDomainTopFp() []*DescribeDomainTopFingerprintResponseBodyDomainTopFp
	SetRequestId(v string) *DescribeDomainTopFingerprintResponseBody
	GetRequestId() *string
}

type DescribeDomainTopFingerprintResponseBody struct {
	// The list of client fingerprints for the website.
	DomainTopFp []*DescribeDomainTopFingerprintResponseBodyDomainTopFp `json:"DomainTopFp,omitempty" xml:"DomainTopFp,omitempty" type:"Repeated"`
	// The ID of the request. The ID is a unique identifier that Alibaba Cloud generates for the request and can be used to troubleshoot issues.
	//
	// example:
	//
	// C33EB3D5-AF96-43CA-9C7E-37A81BC06A1E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDomainTopFingerprintResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainTopFingerprintResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainTopFingerprintResponseBody) GetDomainTopFp() []*DescribeDomainTopFingerprintResponseBodyDomainTopFp {
	return s.DomainTopFp
}

func (s *DescribeDomainTopFingerprintResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDomainTopFingerprintResponseBody) SetDomainTopFp(v []*DescribeDomainTopFingerprintResponseBodyDomainTopFp) *DescribeDomainTopFingerprintResponseBody {
	s.DomainTopFp = v
	return s
}

func (s *DescribeDomainTopFingerprintResponseBody) SetRequestId(v string) *DescribeDomainTopFingerprintResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainTopFingerprintResponseBody) Validate() error {
	if s.DomainTopFp != nil {
		for _, item := range s.DomainTopFp {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDomainTopFingerprintResponseBodyDomainTopFp struct {
	// The domain name of the website.
	//
	// example:
	//
	// example.aliyundoc.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The client fingerprint.
	//
	// example:
	//
	// 8a374c9724582b14a4cfa58c8c9fb2bc
	Fingerprinting *string `json:"Fingerprinting,omitempty" xml:"Fingerprinting,omitempty"`
	// The number of request matches.
	//
	// example:
	//
	// 22121
	Pv *int64 `json:"Pv,omitempty" xml:"Pv,omitempty"`
}

func (s DescribeDomainTopFingerprintResponseBodyDomainTopFp) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainTopFingerprintResponseBodyDomainTopFp) GoString() string {
	return s.String()
}

func (s *DescribeDomainTopFingerprintResponseBodyDomainTopFp) GetDomain() *string {
	return s.Domain
}

func (s *DescribeDomainTopFingerprintResponseBodyDomainTopFp) GetFingerprinting() *string {
	return s.Fingerprinting
}

func (s *DescribeDomainTopFingerprintResponseBodyDomainTopFp) GetPv() *int64 {
	return s.Pv
}

func (s *DescribeDomainTopFingerprintResponseBodyDomainTopFp) SetDomain(v string) *DescribeDomainTopFingerprintResponseBodyDomainTopFp {
	s.Domain = &v
	return s
}

func (s *DescribeDomainTopFingerprintResponseBodyDomainTopFp) SetFingerprinting(v string) *DescribeDomainTopFingerprintResponseBodyDomainTopFp {
	s.Fingerprinting = &v
	return s
}

func (s *DescribeDomainTopFingerprintResponseBodyDomainTopFp) SetPv(v int64) *DescribeDomainTopFingerprintResponseBodyDomainTopFp {
	s.Pv = &v
	return s
}

func (s *DescribeDomainTopFingerprintResponseBodyDomainTopFp) Validate() error {
	return dara.Validate(s)
}
