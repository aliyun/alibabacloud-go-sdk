// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCopyImageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDestinationDescription(v string) *CopyImageRequest
	GetDestinationDescription() *string
	SetDestinationImageName(v string) *CopyImageRequest
	GetDestinationImageName() *string
	SetDestinationRegionId(v string) *CopyImageRequest
	GetDestinationRegionId() *string
	SetImageId(v string) *CopyImageRequest
	GetImageId() *string
	SetRegionId(v string) *CopyImageRequest
	GetRegionId() *string
}

type CopyImageRequest struct {
	// The description of the new image in the destination region. The description must be 2 to 256 characters in length and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// This is a test.
	DestinationDescription *string `json:"DestinationDescription,omitempty" xml:"DestinationDescription,omitempty"`
	// The name of the new image. The name must be 2 to 128 characters in length. The name must start with a letter but cannot start with `http://` or `https://`. The name can contain letters, digits, colons (:), underscores (_), and hyphens (-).
	//
	// This parameter is required.
	//
	// example:
	//
	// Office_Shanghai
	DestinationImageName *string `json:"DestinationImageName,omitempty" xml:"DestinationImageName,omitempty"`
	// The ID of the destination region. The ID must be different from the current region ID of the image. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/196646.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shenzhen
	DestinationRegionId *string `json:"DestinationRegionId,omitempty" xml:"DestinationRegionId,omitempty"`
	// The ID of the image that is copied to the destination region.
	//
	// This parameter is required.
	//
	// example:
	//
	// m-gx2x1dhsmusr2****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/196646.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CopyImageRequest) String() string {
	return dara.Prettify(s)
}

func (s CopyImageRequest) GoString() string {
	return s.String()
}

func (s *CopyImageRequest) GetDestinationDescription() *string {
	return s.DestinationDescription
}

func (s *CopyImageRequest) GetDestinationImageName() *string {
	return s.DestinationImageName
}

func (s *CopyImageRequest) GetDestinationRegionId() *string {
	return s.DestinationRegionId
}

func (s *CopyImageRequest) GetImageId() *string {
	return s.ImageId
}

func (s *CopyImageRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CopyImageRequest) SetDestinationDescription(v string) *CopyImageRequest {
	s.DestinationDescription = &v
	return s
}

func (s *CopyImageRequest) SetDestinationImageName(v string) *CopyImageRequest {
	s.DestinationImageName = &v
	return s
}

func (s *CopyImageRequest) SetDestinationRegionId(v string) *CopyImageRequest {
	s.DestinationRegionId = &v
	return s
}

func (s *CopyImageRequest) SetImageId(v string) *CopyImageRequest {
	s.ImageId = &v
	return s
}

func (s *CopyImageRequest) SetRegionId(v string) *CopyImageRequest {
	s.RegionId = &v
	return s
}

func (s *CopyImageRequest) Validate() error {
	return dara.Validate(s)
}
