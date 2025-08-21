// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCloudDiskTypesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnsRegionId(v string) *DescribeCloudDiskTypesShrinkRequest
	GetEnsRegionId() *string
	SetEnsRegionIdsShrink(v string) *DescribeCloudDiskTypesShrinkRequest
	GetEnsRegionIdsShrink() *string
}

type DescribeCloudDiskTypesShrinkRequest struct {
	// The ID of the edge node.
	//
	// example:
	//
	// cn-chongqing-cmcc
	EnsRegionId *string `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty"`
	// The edge nodes.
	EnsRegionIdsShrink *string `json:"EnsRegionIds,omitempty" xml:"EnsRegionIds,omitempty"`
}

func (s DescribeCloudDiskTypesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudDiskTypesShrinkRequest) GoString() string {
	return s.String()
}

func (s *DescribeCloudDiskTypesShrinkRequest) GetEnsRegionId() *string {
	return s.EnsRegionId
}

func (s *DescribeCloudDiskTypesShrinkRequest) GetEnsRegionIdsShrink() *string {
	return s.EnsRegionIdsShrink
}

func (s *DescribeCloudDiskTypesShrinkRequest) SetEnsRegionId(v string) *DescribeCloudDiskTypesShrinkRequest {
	s.EnsRegionId = &v
	return s
}

func (s *DescribeCloudDiskTypesShrinkRequest) SetEnsRegionIdsShrink(v string) *DescribeCloudDiskTypesShrinkRequest {
	s.EnsRegionIdsShrink = &v
	return s
}

func (s *DescribeCloudDiskTypesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
