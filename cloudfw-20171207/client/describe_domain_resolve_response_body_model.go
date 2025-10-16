// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainResolveResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeDomainResolveResponseBody
	GetRequestId() *string
	SetResolveResult(v *DescribeDomainResolveResponseBodyResolveResult) *DescribeDomainResolveResponseBody
	GetResolveResult() *DescribeDomainResolveResponseBodyResolveResult
}

type DescribeDomainResolveResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// CBF1E9B7-D6A0-4E9E-AD3E-2B47E6C2837D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The details about the DNS record of the domain name.
	ResolveResult *DescribeDomainResolveResponseBodyResolveResult `json:"ResolveResult,omitempty" xml:"ResolveResult,omitempty" type:"Struct"`
}

func (s DescribeDomainResolveResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainResolveResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainResolveResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDomainResolveResponseBody) GetResolveResult() *DescribeDomainResolveResponseBodyResolveResult {
	return s.ResolveResult
}

func (s *DescribeDomainResolveResponseBody) SetRequestId(v string) *DescribeDomainResolveResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainResolveResponseBody) SetResolveResult(v *DescribeDomainResolveResponseBodyResolveResult) *DescribeDomainResolveResponseBody {
	s.ResolveResult = v
	return s
}

func (s *DescribeDomainResolveResponseBody) Validate() error {
	if s.ResolveResult != nil {
		if err := s.ResolveResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDomainResolveResponseBodyResolveResult struct {
	// The IP address to which the domain name is resolved. Multiple IP addresses are separated by commas (,).
	//
	// example:
	//
	// 11.1.X.X,12.1.X.X
	IpAddrs *string `json:"IpAddrs,omitempty" xml:"IpAddrs,omitempty"`
	// The time when the domain name was resolved. The value of this parameter is a timestamp. Unit: seconds.
	//
	// example:
	//
	// 1579091739
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s DescribeDomainResolveResponseBodyResolveResult) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainResolveResponseBodyResolveResult) GoString() string {
	return s.String()
}

func (s *DescribeDomainResolveResponseBodyResolveResult) GetIpAddrs() *string {
	return s.IpAddrs
}

func (s *DescribeDomainResolveResponseBodyResolveResult) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *DescribeDomainResolveResponseBodyResolveResult) SetIpAddrs(v string) *DescribeDomainResolveResponseBodyResolveResult {
	s.IpAddrs = &v
	return s
}

func (s *DescribeDomainResolveResponseBodyResolveResult) SetUpdateTime(v int64) *DescribeDomainResolveResponseBodyResolveResult {
	s.UpdateTime = &v
	return s
}

func (s *DescribeDomainResolveResponseBodyResolveResult) Validate() error {
	return dara.Validate(s)
}
