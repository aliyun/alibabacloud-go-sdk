// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCloudCenterInstancesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *DescribeCloudCenterInstancesShrinkRequest
	GetRegionId() *string
	SetSdkRequestShrink(v string) *DescribeCloudCenterInstancesShrinkRequest
	GetSdkRequestShrink() *string
}

type DescribeCloudCenterInstancesShrinkRequest struct {
	// The ID of the region in which the instance resides.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The Security Center SDK request.
	SdkRequestShrink *string `json:"SdkRequest,omitempty" xml:"SdkRequest,omitempty"`
}

func (s DescribeCloudCenterInstancesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudCenterInstancesShrinkRequest) GoString() string {
	return s.String()
}

func (s *DescribeCloudCenterInstancesShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeCloudCenterInstancesShrinkRequest) GetSdkRequestShrink() *string {
	return s.SdkRequestShrink
}

func (s *DescribeCloudCenterInstancesShrinkRequest) SetRegionId(v string) *DescribeCloudCenterInstancesShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeCloudCenterInstancesShrinkRequest) SetSdkRequestShrink(v string) *DescribeCloudCenterInstancesShrinkRequest {
	s.SdkRequestShrink = &v
	return s
}

func (s *DescribeCloudCenterInstancesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
