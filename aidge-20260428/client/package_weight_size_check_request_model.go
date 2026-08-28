// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPackageWeightSizeCheckRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAnnotatedImageUrl(v string) *PackageWeightSizeCheckRequest
	GetAnnotatedImageUrl() *string
	SetRawImageUrl(v string) *PackageWeightSizeCheckRequest
	GetRawImageUrl() *string
}

type PackageWeightSizeCheckRequest struct {
	// The URL of the annotated image with manual bounding box markings, which is the original image overlaid with blue or red rectangular bounding box lines. The URL must be publicly accessible. The image must not exceed 4000 × 4000 pixels or 10 MB in size. Supported formats: png, jpeg, and jpg.
	//
	// This parameter is required.
	//
	// example:
	//
	// https://example.com/annotated_image.png
	AnnotatedImageUrl *string `json:"AnnotatedImageUrl,omitempty" xml:"AnnotatedImageUrl,omitempty"`
	// The URL of the raw image, which is the unannotated photo of the parcel on the scanning platform. The URL must be publicly accessible. The image must not exceed 4000 × 4000 pixels or 10 MB in size. Supported formats: png, jpeg, and jpg.
	//
	// This parameter is required.
	//
	// example:
	//
	// https://example.com/raw_image.png
	RawImageUrl *string `json:"RawImageUrl,omitempty" xml:"RawImageUrl,omitempty"`
}

func (s PackageWeightSizeCheckRequest) String() string {
	return dara.Prettify(s)
}

func (s PackageWeightSizeCheckRequest) GoString() string {
	return s.String()
}

func (s *PackageWeightSizeCheckRequest) GetAnnotatedImageUrl() *string {
	return s.AnnotatedImageUrl
}

func (s *PackageWeightSizeCheckRequest) GetRawImageUrl() *string {
	return s.RawImageUrl
}

func (s *PackageWeightSizeCheckRequest) SetAnnotatedImageUrl(v string) *PackageWeightSizeCheckRequest {
	s.AnnotatedImageUrl = &v
	return s
}

func (s *PackageWeightSizeCheckRequest) SetRawImageUrl(v string) *PackageWeightSizeCheckRequest {
	s.RawImageUrl = &v
	return s
}

func (s *PackageWeightSizeCheckRequest) Validate() error {
	return dara.Validate(s)
}
