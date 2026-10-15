// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStorefrontSceneRecognitionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageUrl(v string) *StorefrontSceneRecognitionRequest
	GetImageUrl() *string
}

type StorefrontSceneRecognitionRequest struct {
	// The HTTP(S) URL of the image to be reviewed.
	//
	// This parameter is required.
	//
	// example:
	//
	// https://example.com/image.jpg
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
}

func (s StorefrontSceneRecognitionRequest) String() string {
	return dara.Prettify(s)
}

func (s StorefrontSceneRecognitionRequest) GoString() string {
	return s.String()
}

func (s *StorefrontSceneRecognitionRequest) GetImageUrl() *string {
	return s.ImageUrl
}

func (s *StorefrontSceneRecognitionRequest) SetImageUrl(v string) *StorefrontSceneRecognitionRequest {
	s.ImageUrl = &v
	return s
}

func (s *StorefrontSceneRecognitionRequest) Validate() error {
	return dara.Validate(s)
}
