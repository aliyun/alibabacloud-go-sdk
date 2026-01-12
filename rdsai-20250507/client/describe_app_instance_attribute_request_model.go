// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAppInstanceAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceName(v string) *DescribeAppInstanceAttributeRequest
	GetInstanceName() *string
	SetRegionId(v string) *DescribeAppInstanceAttributeRequest
	GetRegionId() *string
}

type DescribeAppInstanceAttributeRequest struct {
	// The region ID.
	//
	// example:
	//
	// ra-supabase-8moov5lxba****
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The operation that you want to perform. Set the value to **DescribeAppInstanceAttribute**.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeAppInstanceAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppInstanceAttributeRequest) GoString() string {
	return s.String()
}

func (s *DescribeAppInstanceAttributeRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeAppInstanceAttributeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeAppInstanceAttributeRequest) SetInstanceName(v string) *DescribeAppInstanceAttributeRequest {
	s.InstanceName = &v
	return s
}

func (s *DescribeAppInstanceAttributeRequest) SetRegionId(v string) *DescribeAppInstanceAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAppInstanceAttributeRequest) Validate() error {
	return dara.Validate(s)
}
