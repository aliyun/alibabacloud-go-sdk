// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeImageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageInstanceId(v string) *DescribeImageRequest
	GetImageInstanceId() *string
	SetImageRegionId(v string) *DescribeImageRequest
	GetImageRegionId() *string
	SetImageRepoId(v string) *DescribeImageRequest
	GetImageRepoId() *string
	SetImageTag(v string) *DescribeImageRequest
	GetImageTag() *string
}

type DescribeImageRequest struct {
	// The instance ID of the image.
	//
	// >  You can call the [DescribeImageInstances](~~DescribeImageInstances~~) operation to query the IDs of instances.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-hfs6gaawhyu6****
	ImageInstanceId *string `json:"ImageInstanceId,omitempty" xml:"ImageInstanceId,omitempty"`
	// The region ID of the image.
	//
	// >  You can call the [DescribeImageInstances](~~DescribeImageInstances~~) operation to query the IDs of regions.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	ImageRegionId *string `json:"ImageRegionId,omitempty" xml:"ImageRegionId,omitempty"`
	// The ID of the image repository.
	//
	// >  You can call the [DescribeImageInstances](~~DescribeImageInstances~~) operation to query the IDs of image repositories.
	//
	// This parameter is required.
	//
	// example:
	//
	// crr-7i88t7lx3fmf****
	ImageRepoId *string `json:"ImageRepoId,omitempty" xml:"ImageRepoId,omitempty"`
	// The tag that is added to the image.
	//
	// >  You can call the [DescribeImageInstances](~~DescribeImageInstances~~) operation to query tags.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1.8.0.15
	ImageTag *string `json:"ImageTag,omitempty" xml:"ImageTag,omitempty"`
}

func (s DescribeImageRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageRequest) GoString() string {
	return s.String()
}

func (s *DescribeImageRequest) GetImageInstanceId() *string {
	return s.ImageInstanceId
}

func (s *DescribeImageRequest) GetImageRegionId() *string {
	return s.ImageRegionId
}

func (s *DescribeImageRequest) GetImageRepoId() *string {
	return s.ImageRepoId
}

func (s *DescribeImageRequest) GetImageTag() *string {
	return s.ImageTag
}

func (s *DescribeImageRequest) SetImageInstanceId(v string) *DescribeImageRequest {
	s.ImageInstanceId = &v
	return s
}

func (s *DescribeImageRequest) SetImageRegionId(v string) *DescribeImageRequest {
	s.ImageRegionId = &v
	return s
}

func (s *DescribeImageRequest) SetImageRepoId(v string) *DescribeImageRequest {
	s.ImageRepoId = &v
	return s
}

func (s *DescribeImageRequest) SetImageTag(v string) *DescribeImageRequest {
	s.ImageTag = &v
	return s
}

func (s *DescribeImageRequest) Validate() error {
	return dara.Validate(s)
}
