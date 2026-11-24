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
	// The list of product image URLs. This parameter is required for single-item and multi-item outfit scenarios.
	//
	// 	- Single-item try-on: Pass images of the same product from different angles. The standard edition supports up to two product images, but requires that all images contain one and only one main product. We recommend that you use product images with a white background.
	//
	// 	- Multi-item outfit: Pass images of all products to be worn. The standard edition supports up to two product images (each image can contain multiple products). A single outfit can include up to six clothing items and accessories.
	//
	// Note: For the clothing replacement, face replacement, and posture replacement modes, specify only one mode per call. Do not upload parameters for multiple modes at the same time.
	//
	// example:
	//
	// https://ae01.alicdn.com/kf/S3470dc9f4be09a6cbed34e90dc8fs.jpg
	ClothImageUrl *string `json:"ClothImageUrl,omitempty" xml:"ClothImageUrl,omitempty"`
	// Specifies the types of products to try on. You can specify multiple types at a time. If you do not specify this parameter, the system automatically identifies all supported categories in the product image.
	//
	// Valid values: tops, outwear, bottoms, dresses, tops and bottoms, socks, shoes, hats, earrings, necklaces, braceles, scarves, bags, and glasses.
	//
	// example:
	//
	// tops
	ClothType *string `json:"ClothType,omitempty" xml:"ClothType,omitempty"`
	// The try-on mode, which determines the maximum number of product images and outfit items. Currently, only `lite` is supported.
	//
	// example:
	//
	// lite
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// The URL of the model scene image. You can specify only one image.
	//
	// 	- The image must support public network access over HTTP or HTTPS.
	//
	// 	- Supported formats: JPG, JPEG, PNG, BMP, and WEBP.
	//
	// 	- The width and height must be between 256 and 2049 pixels.
	//
	// 	- The size of a single file cannot exceed 10 MB.
	//
	// This parameter is required.
	//
	// example:
	//
	// https://example.com/model.jpg\\
	ModelImageUrl *string `json:"ModelImageUrl,omitempty" xml:"ModelImageUrl,omitempty"`
	// The URL of the model template image. You can specify only one image. This image provides reference identity information such as gender, skin color, age, and body type.
	//
	// This parameter is required for the model skin replacement scenario. Do not specify this parameter for the clothing replacement or posture replacement scenarios. For the clothing replacement, face replacement, and posture replacement modes, specify only one mode per call. Do not upload parameters for multiple modes at the same time.
	//
	// We recommend that you use the official model templates. Experience URL: [E-commerce Operation Assistant Aidge console](https://ecoa.console.aliyun.com/translate?category=assetOptimize&apiType=VirtualTryOn)
	//
	// example:
	//
	// https://xxxx
	ModelTemplateUrl *string `json:"ModelTemplateUrl,omitempty" xml:"ModelTemplateUrl,omitempty"`
	// The URL of the posture template image. You can specify only one image. This image provides a reference for the person\\"s posture.
	//
	// This parameter is required for the posture fission scenario.
	//
	// If this parameter is specified, the image is generated based on the template posture.
	//
	// If this parameter is not specified, the posture in the scene image is used if a model scene image is provided. Otherwise, the posture in the model source image is preserved.
	//
	// Note: For the clothing replacement, face replacement, and posture replacement modes, specify only one mode per call. Do not upload parameters for multiple modes at the same time.
	//
	// We recommend that you use the official posture templates. Experience URL: [E-commerce Operation Assistant Aidge console](https://ecoa.console.aliyun.com/translate?category=assetOptimize&apiType=VirtualTryOn)
	//
	// example:
	//
	// https://xxxx
	PoseTemplateUrl *string `json:"PoseTemplateUrl,omitempty" xml:"PoseTemplateUrl,omitempty"`
	// The aspect ratio of the output image. By default, the aspect ratio is the same as that of the original image specified by ModelImageUrl. Valid values: 1:1, 2:3, 3:4, and 9:16.
	//
	// example:
	//
	// 1:1
	Ratio *string `json:"Ratio,omitempty" xml:"Ratio,omitempty"`
	// The resolution of the output image.
	//
	// 	- Synchronous API: supports 1K.
	//
	// 	- Asynchronous API: supports 1K and 2K. The asynchronous API only supports single-item outfit scenarios.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1K
	Resolution *string `json:"Resolution,omitempty" xml:"Resolution,omitempty"`
	// This parameter takes effect only in try-on scenarios and specifies the number of products to try on. Default value: `singletryon`.
	//
	// Valid values:
	//
	// 	- `singletryon`: single-item try-on.
	//
	// 	- `multitryon`: multi-item outfit.
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
