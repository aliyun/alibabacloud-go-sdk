// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRCInstanceTypeFamiliesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *DescribeRCInstanceTypeFamiliesRequest
	GetRegionId() *string
}

type DescribeRCInstanceTypeFamiliesRequest struct {
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeRCInstanceTypeFamiliesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRCInstanceTypeFamiliesRequest) GoString() string {
	return s.String()
}

func (s *DescribeRCInstanceTypeFamiliesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeRCInstanceTypeFamiliesRequest) SetRegionId(v string) *DescribeRCInstanceTypeFamiliesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeRCInstanceTypeFamiliesRequest) Validate() error {
	return dara.Validate(s)
}
