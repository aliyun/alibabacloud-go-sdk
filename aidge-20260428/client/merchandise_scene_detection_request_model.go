// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMerchandiseSceneDetectionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageUrl(v string) *MerchandiseSceneDetectionRequest
	GetImageUrl() *string
}

type MerchandiseSceneDetectionRequest struct {
	// The HTTP(S) URL of the display image to be detected.
	//
	// This parameter is required.
	//
	// example:
	//
	// https://example.com/image.jpg
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
}

func (s MerchandiseSceneDetectionRequest) String() string {
	return dara.Prettify(s)
}

func (s MerchandiseSceneDetectionRequest) GoString() string {
	return s.String()
}

func (s *MerchandiseSceneDetectionRequest) GetImageUrl() *string {
	return s.ImageUrl
}

func (s *MerchandiseSceneDetectionRequest) SetImageUrl(v string) *MerchandiseSceneDetectionRequest {
	s.ImageUrl = &v
	return s
}

func (s *MerchandiseSceneDetectionRequest) Validate() error {
	return dara.Validate(s)
}
