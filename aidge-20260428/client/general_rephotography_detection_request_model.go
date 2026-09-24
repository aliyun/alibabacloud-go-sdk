// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGeneralRephotographyDetectionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageUrl(v string) *GeneralRephotographyDetectionRequest
	GetImageUrl() *string
}

type GeneralRephotographyDetectionRequest struct {
	// The HTTPS URL of the original image to recognize. The URL must be accessible and must not contain whitespace or URL-embedded usernames or passwords.
	//
	// This parameter is required.
	//
	// example:
	//
	// https://example.com/image.jpg
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
}

func (s GeneralRephotographyDetectionRequest) String() string {
	return dara.Prettify(s)
}

func (s GeneralRephotographyDetectionRequest) GoString() string {
	return s.String()
}

func (s *GeneralRephotographyDetectionRequest) GetImageUrl() *string {
	return s.ImageUrl
}

func (s *GeneralRephotographyDetectionRequest) SetImageUrl(v string) *GeneralRephotographyDetectionRequest {
	s.ImageUrl = &v
	return s
}

func (s *GeneralRephotographyDetectionRequest) Validate() error {
	return dara.Validate(s)
}
