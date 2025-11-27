// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVSwitchListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v string) *DescribeVSwitchListRequest
	GetPageNumber() *string
	SetPageSize(v string) *DescribeVSwitchListRequest
	GetPageSize() *string
	SetRegionId(v string) *DescribeVSwitchListRequest
	GetRegionId() *string
	SetVSwitchIds(v []*string) *DescribeVSwitchListRequest
	GetVSwitchIds() []*string
	SetVpcId(v string) *DescribeVSwitchListRequest
	GetVpcId() *string
}

type DescribeVSwitchListRequest struct {
	PageNumber *string   `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *string   `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId   *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	VSwitchIds []*string `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty" type:"Repeated"`
	VpcId      *string   `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DescribeVSwitchListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVSwitchListRequest) GoString() string {
	return s.String()
}

func (s *DescribeVSwitchListRequest) GetPageNumber() *string {
	return s.PageNumber
}

func (s *DescribeVSwitchListRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeVSwitchListRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeVSwitchListRequest) GetVSwitchIds() []*string {
	return s.VSwitchIds
}

func (s *DescribeVSwitchListRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeVSwitchListRequest) SetPageNumber(v string) *DescribeVSwitchListRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeVSwitchListRequest) SetPageSize(v string) *DescribeVSwitchListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeVSwitchListRequest) SetRegionId(v string) *DescribeVSwitchListRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeVSwitchListRequest) SetVSwitchIds(v []*string) *DescribeVSwitchListRequest {
	s.VSwitchIds = v
	return s
}

func (s *DescribeVSwitchListRequest) SetVpcId(v string) *DescribeVSwitchListRequest {
	s.VpcId = &v
	return s
}

func (s *DescribeVSwitchListRequest) Validate() error {
	return dara.Validate(s)
}
