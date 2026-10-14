// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetectAigcImageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageUrl(v string) *DetectAigcImageRequest
	GetImageUrl() *string
}

type DetectAigcImageRequest struct {
	// The URL of the image to detect. Only HTTP/HTTPS protocols are supported. You must provide at least one of `ImageUrl` and `ObjectKey`.
	//
	// example:
	//
	// https://example.com/photo.jpg
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
}

func (s DetectAigcImageRequest) String() string {
	return dara.Prettify(s)
}

func (s DetectAigcImageRequest) GoString() string {
	return s.String()
}

func (s *DetectAigcImageRequest) GetImageUrl() *string {
	return s.ImageUrl
}

func (s *DetectAigcImageRequest) SetImageUrl(v string) *DetectAigcImageRequest {
	s.ImageUrl = &v
	return s
}

func (s *DetectAigcImageRequest) Validate() error {
	return dara.Validate(s)
}
