// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMerchandisePriceTagDetectionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageUrl(v string) *MerchandisePriceTagDetectionRequest
	GetImageUrl() *string
}

type MerchandisePriceTagDetectionRequest struct {
	// The HTTP(S) URL of the image to be detected.
	//
	// This parameter is required.
	//
	// example:
	//
	// https://example.com/image.jpg
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
}

func (s MerchandisePriceTagDetectionRequest) String() string {
	return dara.Prettify(s)
}

func (s MerchandisePriceTagDetectionRequest) GoString() string {
	return s.String()
}

func (s *MerchandisePriceTagDetectionRequest) GetImageUrl() *string {
	return s.ImageUrl
}

func (s *MerchandisePriceTagDetectionRequest) SetImageUrl(v string) *MerchandisePriceTagDetectionRequest {
	s.ImageUrl = &v
	return s
}

func (s *MerchandisePriceTagDetectionRequest) Validate() error {
	return dara.Validate(s)
}
