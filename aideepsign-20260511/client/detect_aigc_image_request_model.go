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
	SetObjectKey(v string) *DetectAigcImageRequest
	GetObjectKey() *string
}

type DetectAigcImageRequest struct {
	// The URL of the image to detect. Only HTTP and HTTPS protocols are supported. You must provide at least one of ImageUrl and ObjectKey.
	//
	// example:
	//
	// https://example.com/photo.jpg
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	// The ObjectKey of the image to detect in OSS. When you use ObjectKey, make sure that the key belongs to the namespace of the current caller. You must provide at least one of ImageUrl and ObjectKey.
	//
	// example:
	//
	// deepsign/123456789/image-generation/abc12345-def6-7890-abcd-ef1234567890.png
	ObjectKey *string `json:"ObjectKey,omitempty" xml:"ObjectKey,omitempty"`
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

func (s *DetectAigcImageRequest) GetObjectKey() *string {
	return s.ObjectKey
}

func (s *DetectAigcImageRequest) SetImageUrl(v string) *DetectAigcImageRequest {
	s.ImageUrl = &v
	return s
}

func (s *DetectAigcImageRequest) SetObjectKey(v string) *DetectAigcImageRequest {
	s.ObjectKey = &v
	return s
}

func (s *DetectAigcImageRequest) Validate() error {
	return dara.Validate(s)
}
