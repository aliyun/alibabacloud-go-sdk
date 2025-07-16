// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainTopClientIpVisitResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetClientIpList(v []*DescribeDomainTopClientIpVisitResponseBodyClientIpList) *DescribeDomainTopClientIpVisitResponseBody
	GetClientIpList() []*DescribeDomainTopClientIpVisitResponseBodyClientIpList
	SetRequestId(v string) *DescribeDomainTopClientIpVisitResponseBody
	GetRequestId() *string
}

type DescribeDomainTopClientIpVisitResponseBody struct {
	// A list of client IP addresses.
	ClientIpList []*DescribeDomainTopClientIpVisitResponseBodyClientIpList `json:"ClientIpList,omitempty" xml:"ClientIpList,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 64D28B53-5902-409B-94F6-FD46680144FE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDomainTopClientIpVisitResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainTopClientIpVisitResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainTopClientIpVisitResponseBody) GetClientIpList() []*DescribeDomainTopClientIpVisitResponseBodyClientIpList {
	return s.ClientIpList
}

func (s *DescribeDomainTopClientIpVisitResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDomainTopClientIpVisitResponseBody) SetClientIpList(v []*DescribeDomainTopClientIpVisitResponseBodyClientIpList) *DescribeDomainTopClientIpVisitResponseBody {
	s.ClientIpList = v
	return s
}

func (s *DescribeDomainTopClientIpVisitResponseBody) SetRequestId(v string) *DescribeDomainTopClientIpVisitResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainTopClientIpVisitResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDomainTopClientIpVisitResponseBodyClientIpList struct {
	// The total number of requests.
	//
	// example:
	//
	// 256
	Acc *int64 `json:"Acc,omitempty" xml:"Acc,omitempty"`
	// The client IP address returned. Only IPv4 addressed are supported.
	//
	// example:
	//
	// 1.1.xxx
	ClientIp *string `json:"ClientIp,omitempty" xml:"ClientIp,omitempty"`
	// The ranking of the client IP address returned.
	//
	// example:
	//
	// 2
	Rank *int32 `json:"Rank,omitempty" xml:"Rank,omitempty"`
	// The total amount of network traffic consumed. Unit: bytes.
	//
	// example:
	//
	// 1024
	Traffic *int64 `json:"Traffic,omitempty" xml:"Traffic,omitempty"`
}

func (s DescribeDomainTopClientIpVisitResponseBodyClientIpList) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainTopClientIpVisitResponseBodyClientIpList) GoString() string {
	return s.String()
}

func (s *DescribeDomainTopClientIpVisitResponseBodyClientIpList) GetAcc() *int64 {
	return s.Acc
}

func (s *DescribeDomainTopClientIpVisitResponseBodyClientIpList) GetClientIp() *string {
	return s.ClientIp
}

func (s *DescribeDomainTopClientIpVisitResponseBodyClientIpList) GetRank() *int32 {
	return s.Rank
}

func (s *DescribeDomainTopClientIpVisitResponseBodyClientIpList) GetTraffic() *int64 {
	return s.Traffic
}

func (s *DescribeDomainTopClientIpVisitResponseBodyClientIpList) SetAcc(v int64) *DescribeDomainTopClientIpVisitResponseBodyClientIpList {
	s.Acc = &v
	return s
}

func (s *DescribeDomainTopClientIpVisitResponseBodyClientIpList) SetClientIp(v string) *DescribeDomainTopClientIpVisitResponseBodyClientIpList {
	s.ClientIp = &v
	return s
}

func (s *DescribeDomainTopClientIpVisitResponseBodyClientIpList) SetRank(v int32) *DescribeDomainTopClientIpVisitResponseBodyClientIpList {
	s.Rank = &v
	return s
}

func (s *DescribeDomainTopClientIpVisitResponseBodyClientIpList) SetTraffic(v int64) *DescribeDomainTopClientIpVisitResponseBodyClientIpList {
	s.Traffic = &v
	return s
}

func (s *DescribeDomainTopClientIpVisitResponseBodyClientIpList) Validate() error {
	return dara.Validate(s)
}
