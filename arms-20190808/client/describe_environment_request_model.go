// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEnvironmentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnvironmentId(v string) *DescribeEnvironmentRequest
	GetEnvironmentId() *string
	SetRegionId(v string) *DescribeEnvironmentRequest
	GetRegionId() *string
}

type DescribeEnvironmentRequest struct {
	// The ID of the environment instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// env-xxxxx
	EnvironmentId *string `json:"EnvironmentId,omitempty" xml:"EnvironmentId,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeEnvironmentRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeEnvironmentRequest) GoString() string {
	return s.String()
}

func (s *DescribeEnvironmentRequest) GetEnvironmentId() *string {
	return s.EnvironmentId
}

func (s *DescribeEnvironmentRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeEnvironmentRequest) SetEnvironmentId(v string) *DescribeEnvironmentRequest {
	s.EnvironmentId = &v
	return s
}

func (s *DescribeEnvironmentRequest) SetRegionId(v string) *DescribeEnvironmentRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeEnvironmentRequest) Validate() error {
	return dara.Validate(s)
}
