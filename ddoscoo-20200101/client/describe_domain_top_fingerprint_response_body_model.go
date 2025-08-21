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
	// The information about the fingerprints of the clients.
	DomainTopFp []*DescribeDomainTopFingerprintResponseBodyDomainTopFp `json:"DomainTopFp,omitempty" xml:"DomainTopFp,omitempty" type:"Repeated"`
	// The request ID.
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
	return dara.Validate(s)
}

type DescribeDomainTopFingerprintResponseBodyDomainTopFp struct {
	// The domain name of the website.
	//
	// example:
	//
	// example.aliyundoc.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The fingerprint of the client.
	//
	// example:
	//
	// 8a374c9724582b14a4cfa58c8c9fb2bc
	Fingerprinting *string `json:"Fingerprinting,omitempty" xml:"Fingerprinting,omitempty"`
	// The page views.
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
