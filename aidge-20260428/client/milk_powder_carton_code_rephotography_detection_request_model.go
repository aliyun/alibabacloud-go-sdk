// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMilkPowderCartonCodeRephotographyDetectionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageUrl(v string) *MilkPowderCartonCodeRephotographyDetectionRequest
	GetImageUrl() *string
}

type MilkPowderCartonCodeRephotographyDetectionRequest struct {
	// The HTTP(S) URL of the image to be analyzed.
	//
	// This parameter is required.
	//
	// example:
	//
	// https://example.com/image.jpg
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
}

func (s MilkPowderCartonCodeRephotographyDetectionRequest) String() string {
	return dara.Prettify(s)
}

func (s MilkPowderCartonCodeRephotographyDetectionRequest) GoString() string {
	return s.String()
}

func (s *MilkPowderCartonCodeRephotographyDetectionRequest) GetImageUrl() *string {
	return s.ImageUrl
}

func (s *MilkPowderCartonCodeRephotographyDetectionRequest) SetImageUrl(v string) *MilkPowderCartonCodeRephotographyDetectionRequest {
	s.ImageUrl = &v
	return s
}

func (s *MilkPowderCartonCodeRephotographyDetectionRequest) Validate() error {
	return dara.Validate(s)
}
