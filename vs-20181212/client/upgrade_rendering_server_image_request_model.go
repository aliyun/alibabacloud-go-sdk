// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeRenderingServerImageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageId(v string) *UpgradeRenderingServerImageRequest
	GetImageId() *string
	SetPrecheck(v bool) *UpgradeRenderingServerImageRequest
	GetPrecheck() *bool
	SetRenderingInstanceIds(v []*string) *UpgradeRenderingServerImageRequest
	GetRenderingInstanceIds() []*string
}

type UpgradeRenderingServerImageRequest struct {
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
	RenderingInstanceIds []*string `json:"RenderingInstanceIds,omitempty" xml:"RenderingInstanceIds,omitempty" type:"Repeated"`
}

func (s UpgradeRenderingServerImageRequest) String() string {
	return dara.Prettify(s)
}

func (s UpgradeRenderingServerImageRequest) GoString() string {
	return s.String()
}

func (s *UpgradeRenderingServerImageRequest) GetImageId() *string {
	return s.ImageId
}

func (s *UpgradeRenderingServerImageRequest) GetPrecheck() *bool {
	return s.Precheck
}

func (s *UpgradeRenderingServerImageRequest) GetRenderingInstanceIds() []*string {
	return s.RenderingInstanceIds
}

func (s *UpgradeRenderingServerImageRequest) SetImageId(v string) *UpgradeRenderingServerImageRequest {
	s.ImageId = &v
	return s
}

func (s *UpgradeRenderingServerImageRequest) SetPrecheck(v bool) *UpgradeRenderingServerImageRequest {
	s.Precheck = &v
	return s
}

func (s *UpgradeRenderingServerImageRequest) SetRenderingInstanceIds(v []*string) *UpgradeRenderingServerImageRequest {
	s.RenderingInstanceIds = v
	return s
}

func (s *UpgradeRenderingServerImageRequest) Validate() error {
	return dara.Validate(s)
}
