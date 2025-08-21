// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDDosEventMaxResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCps(v int64) *DescribeDDosEventMaxResponseBody
	GetCps() *int64
	SetMbps(v int64) *DescribeDDosEventMaxResponseBody
	GetMbps() *int64
	SetQps(v int64) *DescribeDDosEventMaxResponseBody
	GetQps() *int64
	SetRequestId(v string) *DescribeDDosEventMaxResponseBody
	GetRequestId() *string
}

type DescribeDDosEventMaxResponseBody struct {
	// The peak of connection flood attacks. Unit: connections per seconds (CPS).
	//
	// example:
	//
	// 1302
	Cps *int64 `json:"Cps,omitempty" xml:"Cps,omitempty"`
	// The peak of volumetric attacks. Unit: Mbit/s.
	//
	// example:
	//
	// 6809
	Mbps *int64 `json:"Mbps,omitempty" xml:"Mbps,omitempty"`
	// The peak of resource exhaustion attacks. Unit: queries per second (QPS).
	//
	// example:
	//
	// 26314
	Qps *int64 `json:"Qps,omitempty" xml:"Qps,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 5AE2FC86-C840-41AE-9F1A-3A2747C7C1DF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDDosEventMaxResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDDosEventMaxResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDDosEventMaxResponseBody) GetCps() *int64 {
	return s.Cps
}

func (s *DescribeDDosEventMaxResponseBody) GetMbps() *int64 {
	return s.Mbps
}

func (s *DescribeDDosEventMaxResponseBody) GetQps() *int64 {
	return s.Qps
}

func (s *DescribeDDosEventMaxResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDDosEventMaxResponseBody) SetCps(v int64) *DescribeDDosEventMaxResponseBody {
	s.Cps = &v
	return s
}

func (s *DescribeDDosEventMaxResponseBody) SetMbps(v int64) *DescribeDDosEventMaxResponseBody {
	s.Mbps = &v
	return s
}

func (s *DescribeDDosEventMaxResponseBody) SetQps(v int64) *DescribeDDosEventMaxResponseBody {
	s.Qps = &v
	return s
}

func (s *DescribeDDosEventMaxResponseBody) SetRequestId(v string) *DescribeDDosEventMaxResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDDosEventMaxResponseBody) Validate() error {
	return dara.Validate(s)
}
