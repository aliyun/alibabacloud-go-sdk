// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSDGsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceIds(v []*string) *DescribeSDGsRequest
	GetInstanceIds() []*string
	SetSDGIds(v []*string) *DescribeSDGsRequest
	GetSDGIds() []*string
}

type DescribeSDGsRequest struct {
	// The list of AIC instance IDs to query.
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// The list of SDG IDs to query.
	SDGIds []*string `json:"SDGIds,omitempty" xml:"SDGIds,omitempty" type:"Repeated"`
}

func (s DescribeSDGsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSDGsRequest) GoString() string {
	return s.String()
}

func (s *DescribeSDGsRequest) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *DescribeSDGsRequest) GetSDGIds() []*string {
	return s.SDGIds
}

func (s *DescribeSDGsRequest) SetInstanceIds(v []*string) *DescribeSDGsRequest {
	s.InstanceIds = v
	return s
}

func (s *DescribeSDGsRequest) SetSDGIds(v []*string) *DescribeSDGsRequest {
	s.SDGIds = v
	return s
}

func (s *DescribeSDGsRequest) Validate() error {
	return dara.Validate(s)
}
