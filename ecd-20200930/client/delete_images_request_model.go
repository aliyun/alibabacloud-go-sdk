// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteImagesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeleteCascadedBundle(v bool) *DeleteImagesRequest
	GetDeleteCascadedBundle() *bool
	SetImageId(v []*string) *DeleteImagesRequest
	GetImageId() []*string
	SetRegionId(v string) *DeleteImagesRequest
	GetRegionId() *string
}

type DeleteImagesRequest struct {
	// Specifies whether to delete the associated template.
	//
	// example:
	//
	// true
	DeleteCascadedBundle *bool `json:"DeleteCascadedBundle,omitempty" xml:"DeleteCascadedBundle,omitempty"`
	// The image IDs. You can specify 1 to 100 image IDs.
	//
	// This parameter is required.
	ImageId []*string `json:"ImageId,omitempty" xml:"ImageId,omitempty" type:"Repeated"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/196646.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteImagesRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteImagesRequest) GoString() string {
	return s.String()
}

func (s *DeleteImagesRequest) GetDeleteCascadedBundle() *bool {
	return s.DeleteCascadedBundle
}

func (s *DeleteImagesRequest) GetImageId() []*string {
	return s.ImageId
}

func (s *DeleteImagesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteImagesRequest) SetDeleteCascadedBundle(v bool) *DeleteImagesRequest {
	s.DeleteCascadedBundle = &v
	return s
}

func (s *DeleteImagesRequest) SetImageId(v []*string) *DeleteImagesRequest {
	s.ImageId = v
	return s
}

func (s *DeleteImagesRequest) SetRegionId(v string) *DeleteImagesRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteImagesRequest) Validate() error {
	return dara.Validate(s)
}
