// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCandidateInstanceTypeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCandidateInstanceTypes(v []*string) *DescribeCandidateInstanceTypeResponseBody
	GetCandidateInstanceTypes() []*string
	SetCandidateZoneIds(v []*string) *DescribeCandidateInstanceTypeResponseBody
	GetCandidateZoneIds() []*string
	SetRequestId(v string) *DescribeCandidateInstanceTypeResponseBody
	GetRequestId() *string
}

type DescribeCandidateInstanceTypeResponseBody struct {
	// A list of candidate ECS instance types.
	CandidateInstanceTypes []*string `json:"CandidateInstanceTypes,omitempty" xml:"CandidateInstanceTypes,omitempty" type:"Repeated"`
	// A list of candidate availability zones.
	CandidateZoneIds []*string `json:"CandidateZoneIds,omitempty" xml:"CandidateZoneIds,omitempty" type:"Repeated"`
	// The unique ID of the request.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeCandidateInstanceTypeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCandidateInstanceTypeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCandidateInstanceTypeResponseBody) GetCandidateInstanceTypes() []*string {
	return s.CandidateInstanceTypes
}

func (s *DescribeCandidateInstanceTypeResponseBody) GetCandidateZoneIds() []*string {
	return s.CandidateZoneIds
}

func (s *DescribeCandidateInstanceTypeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCandidateInstanceTypeResponseBody) SetCandidateInstanceTypes(v []*string) *DescribeCandidateInstanceTypeResponseBody {
	s.CandidateInstanceTypes = v
	return s
}

func (s *DescribeCandidateInstanceTypeResponseBody) SetCandidateZoneIds(v []*string) *DescribeCandidateInstanceTypeResponseBody {
	s.CandidateZoneIds = v
	return s
}

func (s *DescribeCandidateInstanceTypeResponseBody) SetRequestId(v string) *DescribeCandidateInstanceTypeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCandidateInstanceTypeResponseBody) Validate() error {
	return dara.Validate(s)
}
