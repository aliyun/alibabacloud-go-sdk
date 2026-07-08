// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchCaptureRenderingInstanceScreenshotRequest interface {
	dara.Model
	String() string
	GoString() string
	SetQuality(v int32) *BatchCaptureRenderingInstanceScreenshotRequest
	GetQuality() *int32
	SetRenderingInstanceIds(v []*string) *BatchCaptureRenderingInstanceScreenshotRequest
	GetRenderingInstanceIds() []*string
}

type BatchCaptureRenderingInstanceScreenshotRequest struct {
	// The image quality. Valid values: 1 to 100.
	//
	// example:
	//
	// 60
	Quality *int32 `json:"Quality,omitempty" xml:"Quality,omitempty"`
	// The list of instance IDs. A maximum of 100 instance IDs can be specified.
	//
	// This parameter is required.
	RenderingInstanceIds []*string `json:"RenderingInstanceIds,omitempty" xml:"RenderingInstanceIds,omitempty" type:"Repeated"`
}

func (s BatchCaptureRenderingInstanceScreenshotRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchCaptureRenderingInstanceScreenshotRequest) GoString() string {
	return s.String()
}

func (s *BatchCaptureRenderingInstanceScreenshotRequest) GetQuality() *int32 {
	return s.Quality
}

func (s *BatchCaptureRenderingInstanceScreenshotRequest) GetRenderingInstanceIds() []*string {
	return s.RenderingInstanceIds
}

func (s *BatchCaptureRenderingInstanceScreenshotRequest) SetQuality(v int32) *BatchCaptureRenderingInstanceScreenshotRequest {
	s.Quality = &v
	return s
}

func (s *BatchCaptureRenderingInstanceScreenshotRequest) SetRenderingInstanceIds(v []*string) *BatchCaptureRenderingInstanceScreenshotRequest {
	s.RenderingInstanceIds = v
	return s
}

func (s *BatchCaptureRenderingInstanceScreenshotRequest) Validate() error {
	return dara.Validate(s)
}
