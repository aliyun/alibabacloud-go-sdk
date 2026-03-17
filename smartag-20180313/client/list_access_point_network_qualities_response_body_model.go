// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAccessPointNetworkQualitiesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessPointNetworkQualities(v []*ListAccessPointNetworkQualitiesResponseBodyAccessPointNetworkQualities) *ListAccessPointNetworkQualitiesResponseBody
	GetAccessPointNetworkQualities() []*ListAccessPointNetworkQualitiesResponseBodyAccessPointNetworkQualities
	SetRequestId(v string) *ListAccessPointNetworkQualitiesResponseBody
	GetRequestId() *string
}

type ListAccessPointNetworkQualitiesResponseBody struct {
	// The network quality of the endpoint.
	AccessPointNetworkQualities []*ListAccessPointNetworkQualitiesResponseBodyAccessPointNetworkQualities `json:"AccessPointNetworkQualities,omitempty" xml:"AccessPointNetworkQualities,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 8E8E8C86-1404-122A-A1BB-84BBC2E9A4B1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListAccessPointNetworkQualitiesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAccessPointNetworkQualitiesResponseBody) GoString() string {
	return s.String()
}

func (s *ListAccessPointNetworkQualitiesResponseBody) GetAccessPointNetworkQualities() []*ListAccessPointNetworkQualitiesResponseBodyAccessPointNetworkQualities {
	return s.AccessPointNetworkQualities
}

func (s *ListAccessPointNetworkQualitiesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAccessPointNetworkQualitiesResponseBody) SetAccessPointNetworkQualities(v []*ListAccessPointNetworkQualitiesResponseBodyAccessPointNetworkQualities) *ListAccessPointNetworkQualitiesResponseBody {
	s.AccessPointNetworkQualities = v
	return s
}

func (s *ListAccessPointNetworkQualitiesResponseBody) SetRequestId(v string) *ListAccessPointNetworkQualitiesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAccessPointNetworkQualitiesResponseBody) Validate() error {
	if s.AccessPointNetworkQualities != nil {
		for _, item := range s.AccessPointNetworkQualities {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAccessPointNetworkQualitiesResponseBodyAccessPointNetworkQualities struct {
	// The ID of the endpoint.
	//
	// example:
	//
	// 115
	Id *int32 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The packet loss rate.
	//
	// example:
	//
	// 0.00%
	Loss *string `json:"Loss,omitempty" xml:"Loss,omitempty"`
	// The network latency. Unit: milliseconds.
	//
	// example:
	//
	// 4.98
	Rtt *string `json:"Rtt,omitempty" xml:"Rtt,omitempty"`
}

func (s ListAccessPointNetworkQualitiesResponseBodyAccessPointNetworkQualities) String() string {
	return dara.Prettify(s)
}

func (s ListAccessPointNetworkQualitiesResponseBodyAccessPointNetworkQualities) GoString() string {
	return s.String()
}

func (s *ListAccessPointNetworkQualitiesResponseBodyAccessPointNetworkQualities) GetId() *int32 {
	return s.Id
}

func (s *ListAccessPointNetworkQualitiesResponseBodyAccessPointNetworkQualities) GetLoss() *string {
	return s.Loss
}

func (s *ListAccessPointNetworkQualitiesResponseBodyAccessPointNetworkQualities) GetRtt() *string {
	return s.Rtt
}

func (s *ListAccessPointNetworkQualitiesResponseBodyAccessPointNetworkQualities) SetId(v int32) *ListAccessPointNetworkQualitiesResponseBodyAccessPointNetworkQualities {
	s.Id = &v
	return s
}

func (s *ListAccessPointNetworkQualitiesResponseBodyAccessPointNetworkQualities) SetLoss(v string) *ListAccessPointNetworkQualitiesResponseBodyAccessPointNetworkQualities {
	s.Loss = &v
	return s
}

func (s *ListAccessPointNetworkQualitiesResponseBodyAccessPointNetworkQualities) SetRtt(v string) *ListAccessPointNetworkQualitiesResponseBodyAccessPointNetworkQualities {
	s.Rtt = &v
	return s
}

func (s *ListAccessPointNetworkQualitiesResponseBodyAccessPointNetworkQualities) Validate() error {
	return dara.Validate(s)
}
