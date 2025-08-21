// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainH2FingerprintRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomain(v string) *DescribeDomainH2FingerprintRequest
	GetDomain() *string
	SetEndTime(v int64) *DescribeDomainH2FingerprintRequest
	GetEndTime() *int64
	SetLimit(v int64) *DescribeDomainH2FingerprintRequest
	GetLimit() *int64
	SetStartTime(v int64) *DescribeDomainH2FingerprintRequest
	GetStartTime() *int64
}

type DescribeDomainH2FingerprintRequest struct {
	// The domain name of the website.
	//
	// >  You can call the [DescribeDomains](https://help.aliyun.com/document_detail/91724.html) operation to query the domain names of all websites that are protected by Anti-DDoS Proxy.
	//
	// example:
	//
	// example.aliyundoc.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The end of the time range to query. The value is a UNIX timestamp. Unit: seconds.
	//
	// >  This UNIX timestamp must indicate a point in time that is accurate to the minute.
	//
	// example:
	//
	// 1726318200
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The maximum number of entries to return.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20
	Limit *int64 `json:"Limit,omitempty" xml:"Limit,omitempty"`
	// The beginning of the time range to query. The value is a UNIX timestamp. Unit: seconds.
	//
	// >  This UNIX timestamp must indicate a point in time that is accurate to the minute.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1716435180
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDomainH2FingerprintRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainH2FingerprintRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainH2FingerprintRequest) GetDomain() *string {
	return s.Domain
}

func (s *DescribeDomainH2FingerprintRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeDomainH2FingerprintRequest) GetLimit() *int64 {
	return s.Limit
}

func (s *DescribeDomainH2FingerprintRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeDomainH2FingerprintRequest) SetDomain(v string) *DescribeDomainH2FingerprintRequest {
	s.Domain = &v
	return s
}

func (s *DescribeDomainH2FingerprintRequest) SetEndTime(v int64) *DescribeDomainH2FingerprintRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainH2FingerprintRequest) SetLimit(v int64) *DescribeDomainH2FingerprintRequest {
	s.Limit = &v
	return s
}

func (s *DescribeDomainH2FingerprintRequest) SetStartTime(v int64) *DescribeDomainH2FingerprintRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainH2FingerprintRequest) Validate() error {
	return dara.Validate(s)
}
