// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchCaptureRenderingInstanceScreenshotShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetQuality(v int32) *BatchCaptureRenderingInstanceScreenshotShrinkRequest
	GetQuality() *int32
	SetRenderingInstanceIdsShrink(v string) *BatchCaptureRenderingInstanceScreenshotShrinkRequest
	GetRenderingInstanceIdsShrink() *string
}

type BatchCaptureRenderingInstanceScreenshotShrinkRequest struct {
	// The image quality. Valid values: 1 to 100.
	//
	// example:
	//
	// 60
	Quality *int32 `json:"Quality,omitempty" xml:"Quality,omitempty"`
	// The list of instance IDs. A maximum of 100 instance IDs can be specified.
	//
	// This parameter is required.
	RenderingInstanceIdsShrink *string `json:"RenderingInstanceIds,omitempty" xml:"RenderingInstanceIds,omitempty"`
}

func (s BatchCaptureRenderingInstanceScreenshotShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchCaptureRenderingInstanceScreenshotShrinkRequest) GoString() string {
	return s.String()
}

func (s *BatchCaptureRenderingInstanceScreenshotShrinkRequest) GetQuality() *int32 {
	return s.Quality
}

func (s *BatchCaptureRenderingInstanceScreenshotShrinkRequest) GetRenderingInstanceIdsShrink() *string {
	return s.RenderingInstanceIdsShrink
}

func (s *BatchCaptureRenderingInstanceScreenshotShrinkRequest) SetQuality(v int32) *BatchCaptureRenderingInstanceScreenshotShrinkRequest {
	s.Quality = &v
	return s
}

func (s *BatchCaptureRenderingInstanceScreenshotShrinkRequest) SetRenderingInstanceIdsShrink(v string) *BatchCaptureRenderingInstanceScreenshotShrinkRequest {
	s.RenderingInstanceIdsShrink = &v
	return s
}

func (s *BatchCaptureRenderingInstanceScreenshotShrinkRequest) Validate() error {
	return dara.Validate(s)
}
