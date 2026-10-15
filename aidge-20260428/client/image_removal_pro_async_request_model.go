// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImageRemovalProAsyncRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageUrl(v string) *ImageRemovalProAsyncRequest
	GetImageUrl() *string
}

type ImageRemovalProAsyncRequest struct {
	// The URL of the original image. The image must be in JPG, JPEG, PNG, BMP, or WEBP format. The resolution must be between 256×256 and 3000×3000, and the file size must be less than or equal to 10 MB.
	//
	// This parameter is required.
	//
	// example:
	//
	// https://ae01.alicdn.com/kf/S342f0070dc9f4be09a6cbed34e90dc8fs.jpg
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
}

func (s ImageRemovalProAsyncRequest) String() string {
	return dara.Prettify(s)
}

func (s ImageRemovalProAsyncRequest) GoString() string {
	return s.String()
}

func (s *ImageRemovalProAsyncRequest) GetImageUrl() *string {
	return s.ImageUrl
}

func (s *ImageRemovalProAsyncRequest) SetImageUrl(v string) *ImageRemovalProAsyncRequest {
	s.ImageUrl = &v
	return s
}

func (s *ImageRemovalProAsyncRequest) Validate() error {
	return dara.Validate(s)
}
