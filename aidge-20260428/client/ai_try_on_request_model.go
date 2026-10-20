// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAiTryOnRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClothImageUrl(v string) *AiTryOnRequest
	GetClothImageUrl() *string
	SetClothType(v string) *AiTryOnRequest
	GetClothType() *string
	SetModelImageUrl(v string) *AiTryOnRequest
	GetModelImageUrl() *string
	SetResolution(v string) *AiTryOnRequest
	GetResolution() *string
}

type AiTryOnRequest struct {
	// The URL of the garment image. Only one image is supported. The URL must be a publicly accessible `http` or `https` address. The image must be in JPG, JPEG, PNG, BMP, or WEBP format, with a resolution between 256×256 and 2049×2049 pixels, and a file size not exceeding 10 MB.<br>
	//
	// This parameter is required.
	//
	// example:
	//
	// https://ae01.alicdn.com/kf/S342f0070dc9f4be09a6cbed34e90dc8fs.jpg
	ClothImageUrl *string `json:"ClothImageUrl,omitempty" xml:"ClothImageUrl,omitempty"`
	// The type of the garment. Valid values: tops, bottoms, dresses, tops_and_bottoms, shoes, and hats. If this parameter is not specified, the system automatically recognizes the garment type.
	//
	// example:
	//
	// tops
	ClothType *string `json:"ClothType,omitempty" xml:"ClothType,omitempty"`
	// The URL of the model image. Only one image is supported. The URL must be a publicly accessible `http` or `https` address. The image must be in JPG, JPEG, PNG, BMP, or WEBP format, with a resolution between 256×256 and 2049×2049 pixels, and a file size not exceeding 10 MB.<br>
	//
	// This parameter is required.
	//
	// example:
	//
	// https://ae01.alicdn.com/kf/S342f0070dc9f4be09a6cbed34e90dc8fs.jpg
	ModelImageUrl *string `json:"ModelImageUrl,omitempty" xml:"ModelImageUrl,omitempty"`
	// The resolution of the output image. Synchronous API operations only support 1K resolution.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1K
	Resolution *string `json:"Resolution,omitempty" xml:"Resolution,omitempty"`
}

func (s AiTryOnRequest) String() string {
	return dara.Prettify(s)
}

func (s AiTryOnRequest) GoString() string {
	return s.String()
}

func (s *AiTryOnRequest) GetClothImageUrl() *string {
	return s.ClothImageUrl
}

func (s *AiTryOnRequest) GetClothType() *string {
	return s.ClothType
}

func (s *AiTryOnRequest) GetModelImageUrl() *string {
	return s.ModelImageUrl
}

func (s *AiTryOnRequest) GetResolution() *string {
	return s.Resolution
}

func (s *AiTryOnRequest) SetClothImageUrl(v string) *AiTryOnRequest {
	s.ClothImageUrl = &v
	return s
}

func (s *AiTryOnRequest) SetClothType(v string) *AiTryOnRequest {
	s.ClothType = &v
	return s
}

func (s *AiTryOnRequest) SetModelImageUrl(v string) *AiTryOnRequest {
	s.ModelImageUrl = &v
	return s
}

func (s *AiTryOnRequest) SetResolution(v string) *AiTryOnRequest {
	s.Resolution = &v
	return s
}

func (s *AiTryOnRequest) Validate() error {
	return dara.Validate(s)
}
