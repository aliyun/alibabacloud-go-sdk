// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMerchandisePlacementDetectionLiteRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageUrl(v string) *MerchandisePlacementDetectionLiteRequest
	GetImageUrl() *string
}

type MerchandisePlacementDetectionLiteRequest struct {
	// The HTTPS URL of the display image to be detected.
	//
	// This parameter is required.
	//
	// example:
	//
	// https://example.com/image.jpg
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
}

func (s MerchandisePlacementDetectionLiteRequest) String() string {
	return dara.Prettify(s)
}

func (s MerchandisePlacementDetectionLiteRequest) GoString() string {
	return s.String()
}

func (s *MerchandisePlacementDetectionLiteRequest) GetImageUrl() *string {
	return s.ImageUrl
}

func (s *MerchandisePlacementDetectionLiteRequest) SetImageUrl(v string) *MerchandisePlacementDetectionLiteRequest {
	s.ImageUrl = &v
	return s
}

func (s *MerchandisePlacementDetectionLiteRequest) Validate() error {
	return dara.Validate(s)
}
