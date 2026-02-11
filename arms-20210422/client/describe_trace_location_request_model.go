// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTraceLocationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *DescribeTraceLocationRequest
	GetRegionId() *string
}

type DescribeTraceLocationRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeTraceLocationRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeTraceLocationRequest) GoString() string {
	return s.String()
}

func (s *DescribeTraceLocationRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeTraceLocationRequest) SetRegionId(v string) *DescribeTraceLocationRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeTraceLocationRequest) Validate() error {
	return dara.Validate(s)
}
