// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainBpsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDomainBps(v []*DescribeDomainBpsResponseBodyDomainBps) *DescribeDomainBpsResponseBody
	GetDomainBps() []*DescribeDomainBpsResponseBodyDomainBps
	SetRequestId(v string) *DescribeDomainBpsResponseBody
	GetRequestId() *string
}

type DescribeDomainBpsResponseBody struct {
	// The bandwidths.
	DomainBps []*DescribeDomainBpsResponseBodyDomainBps `json:"DomainBps,omitempty" xml:"DomainBps,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// C33EB3D5-AF96-43CA-9C7E-37A81BC06A1E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDomainBpsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainBpsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainBpsResponseBody) GetDomainBps() []*DescribeDomainBpsResponseBodyDomainBps {
	return s.DomainBps
}

func (s *DescribeDomainBpsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDomainBpsResponseBody) SetDomainBps(v []*DescribeDomainBpsResponseBodyDomainBps) *DescribeDomainBpsResponseBody {
	s.DomainBps = v
	return s
}

func (s *DescribeDomainBpsResponseBody) SetRequestId(v string) *DescribeDomainBpsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainBpsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDomainBpsResponseBodyDomainBps struct {
	// The inbound bandwidth. Unit: bit/s.
	//
	// example:
	//
	// 0
	InBps *int64 `json:"InBps,omitempty" xml:"InBps,omitempty"`
	// The index number of the returned data.
	//
	// example:
	//
	// 1
	Index *int64 `json:"Index,omitempty" xml:"Index,omitempty"`
	// The outbound bandwidth. Unit: bit/s.
	//
	// example:
	//
	// 0
	OutBps *int64 `json:"OutBps,omitempty" xml:"OutBps,omitempty"`
}

func (s DescribeDomainBpsResponseBodyDomainBps) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainBpsResponseBodyDomainBps) GoString() string {
	return s.String()
}

func (s *DescribeDomainBpsResponseBodyDomainBps) GetInBps() *int64 {
	return s.InBps
}

func (s *DescribeDomainBpsResponseBodyDomainBps) GetIndex() *int64 {
	return s.Index
}

func (s *DescribeDomainBpsResponseBodyDomainBps) GetOutBps() *int64 {
	return s.OutBps
}

func (s *DescribeDomainBpsResponseBodyDomainBps) SetInBps(v int64) *DescribeDomainBpsResponseBodyDomainBps {
	s.InBps = &v
	return s
}

func (s *DescribeDomainBpsResponseBodyDomainBps) SetIndex(v int64) *DescribeDomainBpsResponseBodyDomainBps {
	s.Index = &v
	return s
}

func (s *DescribeDomainBpsResponseBodyDomainBps) SetOutBps(v int64) *DescribeDomainBpsResponseBodyDomainBps {
	s.OutBps = &v
	return s
}

func (s *DescribeDomainBpsResponseBodyDomainBps) Validate() error {
	return dara.Validate(s)
}
