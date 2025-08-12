// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUnhealthyHostAvailabilityRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v []*int64) *DescribeUnhealthyHostAvailabilityRequest
	GetId() []*int64
	SetRegionId(v string) *DescribeUnhealthyHostAvailabilityRequest
	GetRegionId() *string
}

type DescribeUnhealthyHostAvailabilityRequest struct {
	// The ID of the availability monitoring task. Valid values of N: 1 to 20.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456
	Id       []*int64 `json:"Id,omitempty" xml:"Id,omitempty" type:"Repeated"`
	RegionId *string  `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeUnhealthyHostAvailabilityRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeUnhealthyHostAvailabilityRequest) GoString() string {
	return s.String()
}

func (s *DescribeUnhealthyHostAvailabilityRequest) GetId() []*int64 {
	return s.Id
}

func (s *DescribeUnhealthyHostAvailabilityRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeUnhealthyHostAvailabilityRequest) SetId(v []*int64) *DescribeUnhealthyHostAvailabilityRequest {
	s.Id = v
	return s
}

func (s *DescribeUnhealthyHostAvailabilityRequest) SetRegionId(v string) *DescribeUnhealthyHostAvailabilityRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeUnhealthyHostAvailabilityRequest) Validate() error {
	return dara.Validate(s)
}
