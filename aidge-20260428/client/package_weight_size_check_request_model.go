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
	// The URL of the annotated image, i.e., the image with blue/red rectangular bounding box lines overlaid on the original image by an operator. The URL must be publicly accessible. The image dimensions must not exceed 4000×4000 pixels, the file size must not exceed 10 MB, and the supported formats are PNG, JPEG, and JPG.
	//
	// This parameter is required.
	//
	// example:
	//
	// https://example.com/annotated_image.png
	AnnotatedImageUrl *string `json:"AnnotatedImageUrl,omitempty" xml:"AnnotatedImageUrl,omitempty"`
	// The URL of the raw image, i.e., the unannotated photo of the parcel on the scanning platform. The URL must be publicly accessible. The image dimensions must not exceed 4000×4000 pixels, the file size must not exceed 10 MB, and the supported formats are PNG, JPEG, and JPG.
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
