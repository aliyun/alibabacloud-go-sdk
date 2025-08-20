// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApsJobDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApsJobId(v string) *DescribeApsJobDetailRequest
	GetApsJobId() *string
	SetRegionId(v string) *DescribeApsJobDetailRequest
	GetRegionId() *string
}

type DescribeApsJobDetailRequest struct {
	// The job ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// aps-******
	ApsJobId *string `json:"ApsJobId,omitempty" xml:"ApsJobId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeApsJobDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeApsJobDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribeApsJobDetailRequest) GetApsJobId() *string {
	return s.ApsJobId
}

func (s *DescribeApsJobDetailRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeApsJobDetailRequest) SetApsJobId(v string) *DescribeApsJobDetailRequest {
	s.ApsJobId = &v
	return s
}

func (s *DescribeApsJobDetailRequest) SetRegionId(v string) *DescribeApsJobDetailRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeApsJobDetailRequest) Validate() error {
	return dara.Validate(s)
}
