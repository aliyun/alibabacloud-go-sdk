// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeRenderingInstanceImageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageId(v string) *UpgradeRenderingInstanceImageRequest
	GetImageId() *string
	SetRenderingInstanceIds(v []*string) *UpgradeRenderingInstanceImageRequest
	GetRenderingInstanceIds() []*string
}

type UpgradeRenderingInstanceImageRequest struct {
	// The ID of the image.
	//
	// This parameter is required.
	//
	// example:
	//
	// m-bp15om9lg9zb20magg86
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The IDs of the cloud application service instances. A maximum of 100 instance IDs can be specified.
	//
	// This parameter is required.
	RenderingInstanceIds []*string `json:"RenderingInstanceIds,omitempty" xml:"RenderingInstanceIds,omitempty" type:"Repeated"`
}

func (s UpgradeRenderingInstanceImageRequest) String() string {
	return dara.Prettify(s)
}

func (s UpgradeRenderingInstanceImageRequest) GoString() string {
	return s.String()
}

func (s *UpgradeRenderingInstanceImageRequest) GetImageId() *string {
	return s.ImageId
}

func (s *UpgradeRenderingInstanceImageRequest) GetRenderingInstanceIds() []*string {
	return s.RenderingInstanceIds
}

func (s *UpgradeRenderingInstanceImageRequest) SetImageId(v string) *UpgradeRenderingInstanceImageRequest {
	s.ImageId = &v
	return s
}

func (s *UpgradeRenderingInstanceImageRequest) SetRenderingInstanceIds(v []*string) *UpgradeRenderingInstanceImageRequest {
	s.RenderingInstanceIds = v
	return s
}

func (s *UpgradeRenderingInstanceImageRequest) Validate() error {
	return dara.Validate(s)
}
