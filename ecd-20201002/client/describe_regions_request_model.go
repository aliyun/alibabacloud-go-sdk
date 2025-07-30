// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRegionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientId(v string) *DescribeRegionsRequest
	GetClientId() *string
	SetRegionId(v string) *DescribeRegionsRequest
	GetRegionId() *string
}

type DescribeRegionsRequest struct {
	// The client ID. The system generates a unique ID for each client.
	//
	// This parameter is required.
	//
	// example:
	//
	// ac4a73ad-789a-449a-a88f-d18571d6****
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeRegionsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRegionsRequest) GoString() string {
	return s.String()
}

func (s *DescribeRegionsRequest) GetClientId() *string {
	return s.ClientId
}

func (s *DescribeRegionsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeRegionsRequest) SetClientId(v string) *DescribeRegionsRequest {
	s.ClientId = &v
	return s
}

func (s *DescribeRegionsRequest) SetRegionId(v string) *DescribeRegionsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeRegionsRequest) Validate() error {
	return dara.Validate(s)
}
