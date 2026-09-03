// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVirtualTryOnRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClothImageUrl(v string) *VirtualTryOnRequest
	GetClothImageUrl() *string
	SetClothType(v string) *VirtualTryOnRequest
	GetClothType() *string
	SetModelImageUrl(v string) *VirtualTryOnRequest
	GetModelImageUrl() *string
	SetResolution(v string) *VirtualTryOnRequest
	GetResolution() *string
}

type VirtualTryOnRequest struct {
	// The URL of the garment image. Only one image is supported. The URL must be a publicly accessible `http`/`https` address. The image must be in JPG, JPEG, PNG, BMP, or WEBP format, with a resolution between 256 × 256 and 2049 × 2049 pixels, and a file size of no more than 10 MB.<br>
	//
	// This parameter is required.
	//
	// example:
	//
	// https://ae01.alicdn.com/kf/S342f0070dc9f4be09a6cbed34e90dc8fs.jpg
	ClothImageUrl *string `json:"ClothImageUrl,omitempty" xml:"ClothImageUrl,omitempty"`
	// The garment type. Valid values: tops, bottoms, dresses, tops_and_bottoms, shoes, and hats. If this parameter is not specified, the system automatically identifies the garment type.
	//
	// example:
	//
	// tops
	ClothType *string `json:"ClothType,omitempty" xml:"ClothType,omitempty"`
	// The URL of the model image. Only one image is supported. The URL must be a publicly accessible `http`/`https` address. The image must be in JPG, JPEG, PNG, BMP, or WEBP format, with a resolution between 256 × 256 and 2049 × 2049 pixels, and a file size of no more than 10 MB.<br>
	//
	// This parameter is required.
	//
	// example:
	//
	// https://ae01.alicdn.com/kf/S342f0070dc9f4be09a6cbed34e90dc8fs.jpg
	ModelImageUrl *string `json:"ModelImageUrl,omitempty" xml:"ModelImageUrl,omitempty"`
	// Required. The image resolution. Valid values: 1K and 2K.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1K
	Resolution *string `json:"Resolution,omitempty" xml:"Resolution,omitempty"`
}

func (s VirtualTryOnRequest) String() string {
	return dara.Prettify(s)
}

func (s VirtualTryOnRequest) GoString() string {
	return s.String()
}

func (s *VirtualTryOnRequest) GetClothImageUrl() *string {
	return s.ClothImageUrl
}

func (s *VirtualTryOnRequest) GetClothType() *string {
	return s.ClothType
}

func (s *VirtualTryOnRequest) GetModelImageUrl() *string {
	return s.ModelImageUrl
}

func (s *VirtualTryOnRequest) GetResolution() *string {
	return s.Resolution
}

func (s *VirtualTryOnRequest) SetClothImageUrl(v string) *VirtualTryOnRequest {
	s.ClothImageUrl = &v
	return s
}

func (s *VirtualTryOnRequest) SetClothType(v string) *VirtualTryOnRequest {
	s.ClothType = &v
	return s
}

func (s *VirtualTryOnRequest) SetModelImageUrl(v string) *VirtualTryOnRequest {
	s.ModelImageUrl = &v
	return s
}

func (s *VirtualTryOnRequest) SetResolution(v string) *VirtualTryOnRequest {
	s.Resolution = &v
	return s
}

func (s *VirtualTryOnRequest) Validate() error {
	return dara.Validate(s)
}
