// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSagTrafficTopNResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeSagTrafficTopNResponseBody
	GetRequestId() *string
	SetTrafficTopN(v []*DescribeSagTrafficTopNResponseBodyTrafficTopN) *DescribeSagTrafficTopNResponseBody
	GetTrafficTopN() []*DescribeSagTrafficTopNResponseBodyTrafficTopN
}

type DescribeSagTrafficTopNResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// AFF7E5A6-6897-4FDC-A5A8-1978B5B3E545
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the data transfer rate of the SAG instance.
	TrafficTopN []*DescribeSagTrafficTopNResponseBodyTrafficTopN `json:"TrafficTopN,omitempty" xml:"TrafficTopN,omitempty" type:"Repeated"`
}

func (s DescribeSagTrafficTopNResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSagTrafficTopNResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSagTrafficTopNResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSagTrafficTopNResponseBody) GetTrafficTopN() []*DescribeSagTrafficTopNResponseBodyTrafficTopN {
	return s.TrafficTopN
}

func (s *DescribeSagTrafficTopNResponseBody) SetRequestId(v string) *DescribeSagTrafficTopNResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSagTrafficTopNResponseBody) SetTrafficTopN(v []*DescribeSagTrafficTopNResponseBodyTrafficTopN) *DescribeSagTrafficTopNResponseBody {
	s.TrafficTopN = v
	return s
}

func (s *DescribeSagTrafficTopNResponseBody) Validate() error {
	if s.TrafficTopN != nil {
		for _, item := range s.TrafficTopN {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeSagTrafficTopNResponseBodyTrafficTopN struct {
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
	// The data transfer rate of the SAG instance. Unit: bit/s
	//
	// example:
	//
	// 3866.6666666666665
	TrafficRate *string `json:"TrafficRate,omitempty" xml:"TrafficRate,omitempty"`
}

func (s DescribeSagTrafficTopNResponseBodyTrafficTopN) String() string {
	return dara.Prettify(s)
}

func (s DescribeSagTrafficTopNResponseBodyTrafficTopN) GoString() string {
	return s.String()
}

func (s *DescribeSagTrafficTopNResponseBodyTrafficTopN) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeSagTrafficTopNResponseBodyTrafficTopN) GetName() *string {
	return s.Name
}

func (s *DescribeSagTrafficTopNResponseBodyTrafficTopN) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeSagTrafficTopNResponseBodyTrafficTopN) GetTrafficRate() *string {
	return s.TrafficRate
}

func (s *DescribeSagTrafficTopNResponseBodyTrafficTopN) SetInstanceId(v string) *DescribeSagTrafficTopNResponseBodyTrafficTopN {
	s.InstanceId = &v
	return s
}

func (s *DescribeSagTrafficTopNResponseBodyTrafficTopN) SetName(v string) *DescribeSagTrafficTopNResponseBodyTrafficTopN {
	s.Name = &v
	return s
}

func (s *DescribeSagTrafficTopNResponseBodyTrafficTopN) SetRegionId(v string) *DescribeSagTrafficTopNResponseBodyTrafficTopN {
	s.RegionId = &v
	return s
}

func (s *DescribeSagTrafficTopNResponseBodyTrafficTopN) SetTrafficRate(v string) *DescribeSagTrafficTopNResponseBodyTrafficTopN {
	s.TrafficRate = &v
	return s
}

func (s *DescribeSagTrafficTopNResponseBodyTrafficTopN) Validate() error {
	return dara.Validate(s)
}
