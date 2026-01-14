// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateIpSetsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAcceleratorId(v string) *CreateIpSetsResponseBody
	GetAcceleratorId() *string
	SetIpSets(v []*CreateIpSetsResponseBodyIpSets) *CreateIpSetsResponseBody
	GetIpSets() []*CreateIpSetsResponseBodyIpSets
	SetRequestId(v string) *CreateIpSetsResponseBody
	GetRequestId() *string
}

type CreateIpSetsResponseBody struct {
	// The GA instance ID.
	//
	// example:
	//
	// ga-bp1yeeq8yfoyszmq****
	AcceleratorId *string `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	// The details about the acceleration regions.
	IpSets []*CreateIpSetsResponseBodyIpSets `json:"IpSets,omitempty" xml:"IpSets,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 1F4B6A4A-C89E-489E-BAF1-52777EE148EF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateIpSetsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateIpSetsResponseBody) GoString() string {
	return s.String()
}

func (s *CreateIpSetsResponseBody) GetAcceleratorId() *string {
	return s.AcceleratorId
}

func (s *CreateIpSetsResponseBody) GetIpSets() []*CreateIpSetsResponseBodyIpSets {
	return s.IpSets
}

func (s *CreateIpSetsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateIpSetsResponseBody) SetAcceleratorId(v string) *CreateIpSetsResponseBody {
	s.AcceleratorId = &v
	return s
}

func (s *CreateIpSetsResponseBody) SetIpSets(v []*CreateIpSetsResponseBodyIpSets) *CreateIpSetsResponseBody {
	s.IpSets = v
	return s
}

func (s *CreateIpSetsResponseBody) SetRequestId(v string) *CreateIpSetsResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateIpSetsResponseBody) Validate() error {
	if s.IpSets != nil {
		for _, item := range s.IpSets {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateIpSetsResponseBodyIpSets struct {
	// The acceleration region ID.
	//
	// example:
	//
	// cn-qingdao
	AccelerateRegionId *string `json:"AccelerateRegionId,omitempty" xml:"AccelerateRegionId,omitempty"`
	// The bandwidth allocated to the acceleration region. Unit: **Mbit/s**.
	//
	// example:
	//
	// 2
	Bandwidth *int32 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The acceleration region ID.
	//
	// example:
	//
	// ips-bp11r5jb8ogp122xl****
	IpSetId *string `json:"IpSetId,omitempty" xml:"IpSetId,omitempty"`
	// The line type of the EIP in the acceleration region.
	//
	// example:
	//
	// BGP
	IspType *string `json:"IspType,omitempty" xml:"IspType,omitempty"`
}

func (s CreateIpSetsResponseBodyIpSets) String() string {
	return dara.Prettify(s)
}

func (s CreateIpSetsResponseBodyIpSets) GoString() string {
	return s.String()
}

func (s *CreateIpSetsResponseBodyIpSets) GetAccelerateRegionId() *string {
	return s.AccelerateRegionId
}

func (s *CreateIpSetsResponseBodyIpSets) GetBandwidth() *int32 {
	return s.Bandwidth
}

func (s *CreateIpSetsResponseBodyIpSets) GetIpSetId() *string {
	return s.IpSetId
}

func (s *CreateIpSetsResponseBodyIpSets) GetIspType() *string {
	return s.IspType
}

func (s *CreateIpSetsResponseBodyIpSets) SetAccelerateRegionId(v string) *CreateIpSetsResponseBodyIpSets {
	s.AccelerateRegionId = &v
	return s
}

func (s *CreateIpSetsResponseBodyIpSets) SetBandwidth(v int32) *CreateIpSetsResponseBodyIpSets {
	s.Bandwidth = &v
	return s
}

func (s *CreateIpSetsResponseBodyIpSets) SetIpSetId(v string) *CreateIpSetsResponseBodyIpSets {
	s.IpSetId = &v
	return s
}

func (s *CreateIpSetsResponseBodyIpSets) SetIspType(v string) *CreateIpSetsResponseBodyIpSets {
	s.IspType = &v
	return s
}

func (s *CreateIpSetsResponseBodyIpSets) Validate() error {
	return dara.Validate(s)
}
