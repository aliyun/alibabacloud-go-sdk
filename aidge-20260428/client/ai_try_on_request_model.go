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
	SetMode(v string) *AiTryOnRequest
	GetMode() *string
	SetModelImageUrl(v string) *AiTryOnRequest
	GetModelImageUrl() *string
	SetModelTemplateUrl(v string) *AiTryOnRequest
	GetModelTemplateUrl() *string
	SetPoseTemplateUrl(v string) *AiTryOnRequest
	GetPoseTemplateUrl() *string
	SetRatio(v string) *AiTryOnRequest
	GetRatio() *string
	SetResolution(v string) *AiTryOnRequest
	GetResolution() *string
	SetTryonType(v string) *AiTryOnRequest
	GetTryonType() *string
}

type AiTryOnRequest struct {
	// The URL of the garment image. Only one image is supported. The URL must be a publicly accessible `http` or `https` address. The image must be in JPG, JPEG, PNG, BMP, or WEBP format. The resolution must be between 256×256 and 2049×2049 pixels, and the file size cannot exceed 10 MB.<br>
	//
	// example:
	//
	// https://ae01.alicdn.com/kf/S342f0070dc9f4be09a6cbed34e90dc8fs.jpg
	ClothImageUrl *string `json:"ClothImageUrl,omitempty" xml:"ClothImageUrl,omitempty"`
	// The clothing type. If specified, valid values are tops, bottoms, dresses, tops_and_bottoms, shoes, and hats. If this parameter is not specified, the clothing type is automatically identified.
	//
	// example:
	//
	// tops
	ClothType *string `json:"ClothType,omitempty" xml:"ClothType,omitempty"`
	// The model tier. Set the value to lite. This parameter is required.
	//
	// example:
	//
	// lite
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// The URL of the model image. Only one image is supported. The URL must be a publicly accessible `http` or `https` address. The image must be in JPG, JPEG, PNG, BMP, or WEBP format. The resolution must be between 256×256 and 2049×2049 pixels, and the file size cannot exceed 10 MB.<br>
	//
	// This parameter is required.
	//
	// example:
	//
	// https://ae01.alicdn.com/kf/S342f0070dc9f4be09a6cbed34e90dc8fs.jpg
	ModelImageUrl *string `json:"ModelImageUrl,omitempty" xml:"ModelImageUrl,omitempty"`
	// The URL of the identity template image. Only one image is supported. This parameter is mutually exclusive with PoseTemplateUrl.
	//
	// example:
	//
	// https://xxxx
	ModelTemplateUrl *string `json:"ModelTemplateUrl,omitempty" xml:"ModelTemplateUrl,omitempty"`
	// The URL of the pose template image. Only one image is supported. This parameter is mutually exclusive with ModelTemplateUrl.
	//
	// example:
	//
	// https://xxxx
	PoseTemplateUrl *string `json:"PoseTemplateUrl,omitempty" xml:"PoseTemplateUrl,omitempty"`
	// The aspect ratio of the output image. Valid values: 1:1, 2:3, 3:4, and 9:16. This parameter is optional. If this parameter is not specified, the aspect ratio of the original model image is used.
	//
	// example:
	//
	// 1:1
	Ratio *string `json:"Ratio,omitempty" xml:"Ratio,omitempty"`
	// The resolution of the output image. Synchronous API operations only support 1K resolution.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1K
	Resolution *string `json:"Resolution,omitempty" xml:"Resolution,omitempty"`
	// The try-on type. Valid values: singletryon and multitryon. This parameter is optional. If this parameter is not specified, the workflow defaults to singletryon based on whether a product image is provided.
	//
	// example:
	//
	// singletryon
	TryonType *string `json:"TryonType,omitempty" xml:"TryonType,omitempty"`
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

func (s *AiTryOnRequest) GetMode() *string {
	return s.Mode
}

func (s *AiTryOnRequest) GetModelImageUrl() *string {
	return s.ModelImageUrl
}

func (s *AiTryOnRequest) GetModelTemplateUrl() *string {
	return s.ModelTemplateUrl
}

func (s *AiTryOnRequest) GetPoseTemplateUrl() *string {
	return s.PoseTemplateUrl
}

func (s *AiTryOnRequest) GetRatio() *string {
	return s.Ratio
}

func (s *AiTryOnRequest) GetResolution() *string {
	return s.Resolution
}

func (s *AiTryOnRequest) GetTryonType() *string {
	return s.TryonType
}

func (s *AiTryOnRequest) SetClothImageUrl(v string) *AiTryOnRequest {
	s.ClothImageUrl = &v
	return s
}

func (s *AiTryOnRequest) SetClothType(v string) *AiTryOnRequest {
	s.ClothType = &v
	return s
}

func (s *AiTryOnRequest) SetMode(v string) *AiTryOnRequest {
	s.Mode = &v
	return s
}

func (s *AiTryOnRequest) SetModelImageUrl(v string) *AiTryOnRequest {
	s.ModelImageUrl = &v
	return s
}

func (s *AiTryOnRequest) SetModelTemplateUrl(v string) *AiTryOnRequest {
	s.ModelTemplateUrl = &v
	return s
}

func (s *AiTryOnRequest) SetPoseTemplateUrl(v string) *AiTryOnRequest {
	s.PoseTemplateUrl = &v
	return s
}

func (s *AiTryOnRequest) SetRatio(v string) *AiTryOnRequest {
	s.Ratio = &v
	return s
}

func (s *AiTryOnRequest) SetResolution(v string) *AiTryOnRequest {
	s.Resolution = &v
	return s
}

func (s *AiTryOnRequest) SetTryonType(v string) *AiTryOnRequest {
	s.TryonType = &v
	return s
}

func (s *AiTryOnRequest) Validate() error {
	return dara.Validate(s)
}
