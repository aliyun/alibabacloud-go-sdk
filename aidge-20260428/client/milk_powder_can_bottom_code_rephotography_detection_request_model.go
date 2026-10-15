// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMilkPowderCanBottomCodeRephotographyDetectionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageUrl(v string) *MilkPowderCanBottomCodeRephotographyDetectionRequest
	GetImageUrl() *string
}

type MilkPowderCanBottomCodeRephotographyDetectionRequest struct {
	// The HTTP(S) URL of the image to be detected.
	//
	// This parameter is required.
	//
	// example:
	//
	// https://example.com/image.jpg
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
}

func (s MilkPowderCanBottomCodeRephotographyDetectionRequest) String() string {
	return dara.Prettify(s)
}

func (s MilkPowderCanBottomCodeRephotographyDetectionRequest) GoString() string {
	return s.String()
}

func (s *MilkPowderCanBottomCodeRephotographyDetectionRequest) GetImageUrl() *string {
	return s.ImageUrl
}

func (s *MilkPowderCanBottomCodeRephotographyDetectionRequest) SetImageUrl(v string) *MilkPowderCanBottomCodeRephotographyDetectionRequest {
	s.ImageUrl = &v
	return s
}

func (s *MilkPowderCanBottomCodeRephotographyDetectionRequest) Validate() error {
	return dara.Validate(s)
}
