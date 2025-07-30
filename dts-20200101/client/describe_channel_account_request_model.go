// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeChannelAccountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDtsJobId(v string) *DescribeChannelAccountRequest
	GetDtsJobId() *string
	SetOwnerId(v string) *DescribeChannelAccountRequest
	GetOwnerId() *string
	SetRegion(v string) *DescribeChannelAccountRequest
	GetRegion() *string
	SetRegionId(v string) *DescribeChannelAccountRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeChannelAccountRequest
	GetResourceGroupId() *string
	SetTags(v string) *DescribeChannelAccountRequest
	GetTags() *string
	SetType(v string) *DescribeChannelAccountRequest
	GetType() *string
}

type DescribeChannelAccountRequest struct {
	// This parameter is required.
	DtsJobId        *string `json:"DtsJobId,omitempty" xml:"DtsJobId,omitempty"`
	OwnerId         *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	Region          *string `json:"Region,omitempty" xml:"Region,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Tags            *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	Type            *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeChannelAccountRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeChannelAccountRequest) GoString() string {
	return s.String()
}

func (s *DescribeChannelAccountRequest) GetDtsJobId() *string {
	return s.DtsJobId
}

func (s *DescribeChannelAccountRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *DescribeChannelAccountRequest) GetRegion() *string {
	return s.Region
}

func (s *DescribeChannelAccountRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeChannelAccountRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeChannelAccountRequest) GetTags() *string {
	return s.Tags
}

func (s *DescribeChannelAccountRequest) GetType() *string {
	return s.Type
}

func (s *DescribeChannelAccountRequest) SetDtsJobId(v string) *DescribeChannelAccountRequest {
	s.DtsJobId = &v
	return s
}

func (s *DescribeChannelAccountRequest) SetOwnerId(v string) *DescribeChannelAccountRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeChannelAccountRequest) SetRegion(v string) *DescribeChannelAccountRequest {
	s.Region = &v
	return s
}

func (s *DescribeChannelAccountRequest) SetRegionId(v string) *DescribeChannelAccountRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeChannelAccountRequest) SetResourceGroupId(v string) *DescribeChannelAccountRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeChannelAccountRequest) SetTags(v string) *DescribeChannelAccountRequest {
	s.Tags = &v
	return s
}

func (s *DescribeChannelAccountRequest) SetType(v string) *DescribeChannelAccountRequest {
	s.Type = &v
	return s
}

func (s *DescribeChannelAccountRequest) Validate() error {
	return dara.Validate(s)
}
