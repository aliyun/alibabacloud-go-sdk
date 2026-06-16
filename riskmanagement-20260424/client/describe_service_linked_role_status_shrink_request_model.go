// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeServiceLinkedRoleStatusShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *DescribeServiceLinkedRoleStatusShrinkRequest
	GetRegionId() *string
	SetSdkRequestShrink(v string) *DescribeServiceLinkedRoleStatusShrinkRequest
	GetSdkRequestShrink() *string
}

type DescribeServiceLinkedRoleStatusShrinkRequest struct {
	RegionId         *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SdkRequestShrink *string `json:"SdkRequest,omitempty" xml:"SdkRequest,omitempty"`
}

func (s DescribeServiceLinkedRoleStatusShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeServiceLinkedRoleStatusShrinkRequest) GoString() string {
	return s.String()
}

func (s *DescribeServiceLinkedRoleStatusShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeServiceLinkedRoleStatusShrinkRequest) GetSdkRequestShrink() *string {
	return s.SdkRequestShrink
}

func (s *DescribeServiceLinkedRoleStatusShrinkRequest) SetRegionId(v string) *DescribeServiceLinkedRoleStatusShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeServiceLinkedRoleStatusShrinkRequest) SetSdkRequestShrink(v string) *DescribeServiceLinkedRoleStatusShrinkRequest {
	s.SdkRequestShrink = &v
	return s
}

func (s *DescribeServiceLinkedRoleStatusShrinkRequest) Validate() error {
	return dara.Validate(s)
}
