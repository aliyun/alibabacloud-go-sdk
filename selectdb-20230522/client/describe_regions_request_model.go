// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRegionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegion(v string) *DescribeRegionsRequest
	GetRegion() *string
	SetZoneId(v string) *DescribeRegionsRequest
	GetZoneId() *string
}

type DescribeRegionsRequest struct {
	// The region ID.
	//
	// example:
	//
	// cn-beijing
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The zone ID.
	//
	// example:
	//
	// cn-beijing-h
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeRegionsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRegionsRequest) GoString() string {
	return s.String()
}

func (s *DescribeRegionsRequest) GetRegion() *string {
	return s.Region
}

func (s *DescribeRegionsRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeRegionsRequest) SetRegion(v string) *DescribeRegionsRequest {
	s.Region = &v
	return s
}

func (s *DescribeRegionsRequest) SetZoneId(v string) *DescribeRegionsRequest {
	s.ZoneId = &v
	return s
}

func (s *DescribeRegionsRequest) Validate() error {
	return dara.Validate(s)
}
