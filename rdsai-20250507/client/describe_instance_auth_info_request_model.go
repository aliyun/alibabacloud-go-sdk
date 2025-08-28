// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceAuthInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceName(v string) *DescribeInstanceAuthInfoRequest
	GetInstanceName() *string
	SetRegionId(v string) *DescribeInstanceAuthInfoRequest
	GetRegionId() *string
}

type DescribeInstanceAuthInfoRequest struct {
	// example:
	//
	// ra-supabase-8moov5lxba****
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeInstanceAuthInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceAuthInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceAuthInfoRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeInstanceAuthInfoRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeInstanceAuthInfoRequest) SetInstanceName(v string) *DescribeInstanceAuthInfoRequest {
	s.InstanceName = &v
	return s
}

func (s *DescribeInstanceAuthInfoRequest) SetRegionId(v string) *DescribeInstanceAuthInfoRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeInstanceAuthInfoRequest) Validate() error {
	return dara.Validate(s)
}
