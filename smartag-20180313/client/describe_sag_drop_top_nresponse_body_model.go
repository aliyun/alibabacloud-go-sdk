// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSagDropTopNResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDropTopN(v []*DescribeSagDropTopNResponseBodyDropTopN) *DescribeSagDropTopNResponseBody
	GetDropTopN() []*DescribeSagDropTopNResponseBodyDropTopN
	SetRequestId(v string) *DescribeSagDropTopNResponseBody
	GetRequestId() *string
}

type DescribeSagDropTopNResponseBody struct {
	// The information about packets dropped by the SAG instance.
	DropTopN []*DescribeSagDropTopNResponseBodyDropTopN `json:"DropTopN,omitempty" xml:"DropTopN,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// AFF7E5A6-6897-4FDC-A5A8-1978B5B3E545
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeSagDropTopNResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSagDropTopNResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSagDropTopNResponseBody) GetDropTopN() []*DescribeSagDropTopNResponseBodyDropTopN {
	return s.DropTopN
}

func (s *DescribeSagDropTopNResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSagDropTopNResponseBody) SetDropTopN(v []*DescribeSagDropTopNResponseBodyDropTopN) *DescribeSagDropTopNResponseBody {
	s.DropTopN = v
	return s
}

func (s *DescribeSagDropTopNResponseBody) SetRequestId(v string) *DescribeSagDropTopNResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSagDropTopNResponseBody) Validate() error {
	if s.DropTopN != nil {
		for _, item := range s.DropTopN {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeSagDropTopNResponseBodyDropTopN struct {
	// The packet loss rate of the SAG instance. Unit: packets per second (PPS).
	//
	// example:
	//
	// 0.0
	DropRate *string `json:"DropRate,omitempty" xml:"DropRate,omitempty"`
	// The ID of the SAG instance.
	//
	// example:
	//
	// sag-whfn****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the SAG instance.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the region where the SAG instance is deployed.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeSagDropTopNResponseBodyDropTopN) String() string {
	return dara.Prettify(s)
}

func (s DescribeSagDropTopNResponseBodyDropTopN) GoString() string {
	return s.String()
}

func (s *DescribeSagDropTopNResponseBodyDropTopN) GetDropRate() *string {
	return s.DropRate
}

func (s *DescribeSagDropTopNResponseBodyDropTopN) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeSagDropTopNResponseBodyDropTopN) GetName() *string {
	return s.Name
}

func (s *DescribeSagDropTopNResponseBodyDropTopN) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeSagDropTopNResponseBodyDropTopN) SetDropRate(v string) *DescribeSagDropTopNResponseBodyDropTopN {
	s.DropRate = &v
	return s
}

func (s *DescribeSagDropTopNResponseBodyDropTopN) SetInstanceId(v string) *DescribeSagDropTopNResponseBodyDropTopN {
	s.InstanceId = &v
	return s
}

func (s *DescribeSagDropTopNResponseBodyDropTopN) SetName(v string) *DescribeSagDropTopNResponseBodyDropTopN {
	s.Name = &v
	return s
}

func (s *DescribeSagDropTopNResponseBodyDropTopN) SetRegionId(v string) *DescribeSagDropTopNResponseBodyDropTopN {
	s.RegionId = &v
	return s
}

func (s *DescribeSagDropTopNResponseBodyDropTopN) Validate() error {
	return dara.Validate(s)
}
