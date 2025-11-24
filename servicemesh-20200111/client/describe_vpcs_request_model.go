// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVpcsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *DescribeVpcsRequest
	GetRegionId() *string
}

type DescribeVpcsRequest struct {
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeVpcsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcsRequest) GoString() string {
	return s.String()
}

func (s *DescribeVpcsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeVpcsRequest) SetRegionId(v string) *DescribeVpcsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeVpcsRequest) Validate() error {
	return dara.Validate(s)
}
