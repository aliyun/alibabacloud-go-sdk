// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetectImageBasicInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageUrl(v string) *DetectImageBasicInfoRequest
	GetImageUrl() *string
}

type DetectImageBasicInfoRequest struct {
	// The URL of the image. Only HTTP/HTTPS protocols are supported. Specify either `ImageUrl` or `ObjectKey`. At least one of them is required.
	//
	// example:
	//
	// https://example.com/photo.jpg
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
}

func (s DetectImageBasicInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s DetectImageBasicInfoRequest) GoString() string {
	return s.String()
}

func (s *DetectImageBasicInfoRequest) GetImageUrl() *string {
	return s.ImageUrl
}

func (s *DetectImageBasicInfoRequest) SetImageUrl(v string) *DetectImageBasicInfoRequest {
	s.ImageUrl = &v
	return s
}

func (s *DetectImageBasicInfoRequest) Validate() error {
	return dara.Validate(s)
}
