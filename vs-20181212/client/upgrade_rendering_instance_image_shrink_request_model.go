// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeRenderingInstanceImageShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageId(v string) *UpgradeRenderingInstanceImageShrinkRequest
	GetImageId() *string
	SetRenderingInstanceIdsShrink(v string) *UpgradeRenderingInstanceImageShrinkRequest
	GetRenderingInstanceIdsShrink() *string
}

type UpgradeRenderingInstanceImageShrinkRequest struct {
	// The image ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// m-bp15om9lg9zb20magg86
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The list of cloud application service instance IDs. A maximum of 100 IDs can be specified.
	//
	// This parameter is required.
	RenderingInstanceIdsShrink *string `json:"RenderingInstanceIds,omitempty" xml:"RenderingInstanceIds,omitempty"`
}

func (s UpgradeRenderingInstanceImageShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpgradeRenderingInstanceImageShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpgradeRenderingInstanceImageShrinkRequest) GetImageId() *string {
	return s.ImageId
}

func (s *UpgradeRenderingInstanceImageShrinkRequest) GetRenderingInstanceIdsShrink() *string {
	return s.RenderingInstanceIdsShrink
}

func (s *UpgradeRenderingInstanceImageShrinkRequest) SetImageId(v string) *UpgradeRenderingInstanceImageShrinkRequest {
	s.ImageId = &v
	return s
}

func (s *UpgradeRenderingInstanceImageShrinkRequest) SetRenderingInstanceIdsShrink(v string) *UpgradeRenderingInstanceImageShrinkRequest {
	s.RenderingInstanceIdsShrink = &v
	return s
}

func (s *UpgradeRenderingInstanceImageShrinkRequest) Validate() error {
	return dara.Validate(s)
}
