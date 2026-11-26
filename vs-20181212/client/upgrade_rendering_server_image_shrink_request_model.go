// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeRenderingServerImageShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageId(v string) *UpgradeRenderingServerImageShrinkRequest
	GetImageId() *string
	SetPrecheck(v bool) *UpgradeRenderingServerImageShrinkRequest
	GetPrecheck() *bool
	SetRenderingInstanceIdsShrink(v string) *UpgradeRenderingServerImageShrinkRequest
	GetRenderingInstanceIdsShrink() *string
}

type UpgradeRenderingServerImageShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// m-9timxhrrgopkec8ju
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// example:
	//
	// true
	Precheck *bool `json:"Precheck,omitempty" xml:"Precheck,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// [\"render-30d593fb042946c58568316956624c6b\"]
	RenderingInstanceIdsShrink *string `json:"RenderingInstanceIds,omitempty" xml:"RenderingInstanceIds,omitempty"`
}

func (s UpgradeRenderingServerImageShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpgradeRenderingServerImageShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpgradeRenderingServerImageShrinkRequest) GetImageId() *string {
	return s.ImageId
}

func (s *UpgradeRenderingServerImageShrinkRequest) GetPrecheck() *bool {
	return s.Precheck
}

func (s *UpgradeRenderingServerImageShrinkRequest) GetRenderingInstanceIdsShrink() *string {
	return s.RenderingInstanceIdsShrink
}

func (s *UpgradeRenderingServerImageShrinkRequest) SetImageId(v string) *UpgradeRenderingServerImageShrinkRequest {
	s.ImageId = &v
	return s
}

func (s *UpgradeRenderingServerImageShrinkRequest) SetPrecheck(v bool) *UpgradeRenderingServerImageShrinkRequest {
	s.Precheck = &v
	return s
}

func (s *UpgradeRenderingServerImageShrinkRequest) SetRenderingInstanceIdsShrink(v string) *UpgradeRenderingServerImageShrinkRequest {
	s.RenderingInstanceIdsShrink = &v
	return s
}

func (s *UpgradeRenderingServerImageShrinkRequest) Validate() error {
	return dara.Validate(s)
}
