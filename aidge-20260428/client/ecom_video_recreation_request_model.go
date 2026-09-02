// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEcomVideoRecreationRequest interface {
  dara.Model
  String() string
  GoString() string
  SetInput(v *EcomVideoRecreationRequestInput) *EcomVideoRecreationRequest
  GetInput() *EcomVideoRecreationRequestInput 
  SetOutput(v *EcomVideoRecreationRequestOutput) *EcomVideoRecreationRequest
  GetOutput() *EcomVideoRecreationRequestOutput 
}

type EcomVideoRecreationRequest struct {
  // The input parameters for video remix.
  // 
  // This parameter is required.
  Input *EcomVideoRecreationRequestInput `json:"Input,omitempty" xml:"Input,omitempty" type:"Struct"`
  // The output specifications for the final video.
  Output *EcomVideoRecreationRequestOutput `json:"Output,omitempty" xml:"Output,omitempty" type:"Struct"`
}

func (s EcomVideoRecreationRequest) String() string {
  return dara.Prettify(s)
}

func (s EcomVideoRecreationRequest) GoString() string {
  return s.String()
}

func (s *EcomVideoRecreationRequest) GetInput() *EcomVideoRecreationRequestInput  {
  return s.Input
}

func (s *EcomVideoRecreationRequest) GetOutput() *EcomVideoRecreationRequestOutput  {
  return s.Output
}

func (s *EcomVideoRecreationRequest) SetInput(v *EcomVideoRecreationRequestInput) *EcomVideoRecreationRequest {
  s.Input = v
  return s
}

func (s *EcomVideoRecreationRequest) SetOutput(v *EcomVideoRecreationRequestOutput) *EcomVideoRecreationRequest {
  s.Output = v
  return s
}

func (s *EcomVideoRecreationRequest) Validate() error {
  if s.Input != nil {
    if err := s.Input.Validate(); err != nil {
      return err
    }
  }
  if s.Output != nil {
    if err := s.Output.Validate(); err != nil {
      return err
    }
  }
  return nil
}

type EcomVideoRecreationRequestInput struct {
  // The description or supplementary constraints for the target person in person replacement mode. 1 to 500 characters. Required when PersonReferenceImageUrls is not provided.   
  // 
  // Example: The target person is an adult male. Retain the original clothing and actions.
  // 
  // example:
  // 
  // The target character is an adult male, with the original costume and movements preserved.
  ChangeDescription *string `json:"ChangeDescription,omitempty" xml:"ChangeDescription,omitempty"`
  // The replacement mode. Valid values: `product_replacement` (default) and `person_replacement`.
  // 
  // example:
  // 
  // product_replacement
  Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
  // The URLs of target person reference images for person replacement. 1 to 5 images of the same person are supported. Arrange images in the following order: face close-up, front view, 45-degree angle, side view, and back view.  
  // 
  // Example: ["https://example.com/person.jpg"]
  PersonReferenceImageUrls []*string `json:"PersonReferenceImageUrls,omitempty" xml:"PersonReferenceImageUrls,omitempty" type:"Repeated"`
  // The URL of the target product image. Required for product replacement. Exactly one image must be provided. A clear subject with no occlusion and a clean background is recommended.  
  // 
  // Example: ["https://example.com/product.png"]
  ProductImageUrls []*string `json:"ProductImageUrls,omitempty" xml:"ProductImageUrls,omitempty" type:"Repeated"`
  // The target product information. Provide this parameter to improve voiceover accuracy.
  ProductInfo *EcomVideoRecreationRequestInputProductInfo `json:"ProductInfo,omitempty" xml:"ProductInfo,omitempty" type:"Struct"`
  // The HTTP(S) URL of the reference video. The video duration must be in the range of 2 to 360 seconds. The URL must remain accessible during task execution. Set the URL validity period to at least 24 hours.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // https://example.com/source.mp4
  SourceVideoUrl *string `json:"SourceVideoUrl,omitempty" xml:"SourceVideoUrl,omitempty"`
}

func (s EcomVideoRecreationRequestInput) String() string {
  return dara.Prettify(s)
}

func (s EcomVideoRecreationRequestInput) GoString() string {
  return s.String()
}

func (s *EcomVideoRecreationRequestInput) GetChangeDescription() *string  {
  return s.ChangeDescription
}

func (s *EcomVideoRecreationRequestInput) GetMode() *string  {
  return s.Mode
}

func (s *EcomVideoRecreationRequestInput) GetPersonReferenceImageUrls() []*string  {
  return s.PersonReferenceImageUrls
}

func (s *EcomVideoRecreationRequestInput) GetProductImageUrls() []*string  {
  return s.ProductImageUrls
}

func (s *EcomVideoRecreationRequestInput) GetProductInfo() *EcomVideoRecreationRequestInputProductInfo  {
  return s.ProductInfo
}

func (s *EcomVideoRecreationRequestInput) GetSourceVideoUrl() *string  {
  return s.SourceVideoUrl
}

func (s *EcomVideoRecreationRequestInput) SetChangeDescription(v string) *EcomVideoRecreationRequestInput {
  s.ChangeDescription = &v
  return s
}

func (s *EcomVideoRecreationRequestInput) SetMode(v string) *EcomVideoRecreationRequestInput {
  s.Mode = &v
  return s
}

func (s *EcomVideoRecreationRequestInput) SetPersonReferenceImageUrls(v []*string) *EcomVideoRecreationRequestInput {
  s.PersonReferenceImageUrls = v
  return s
}

func (s *EcomVideoRecreationRequestInput) SetProductImageUrls(v []*string) *EcomVideoRecreationRequestInput {
  s.ProductImageUrls = v
  return s
}

func (s *EcomVideoRecreationRequestInput) SetProductInfo(v *EcomVideoRecreationRequestInputProductInfo) *EcomVideoRecreationRequestInput {
  s.ProductInfo = v
  return s
}

func (s *EcomVideoRecreationRequestInput) SetSourceVideoUrl(v string) *EcomVideoRecreationRequestInput {
  s.SourceVideoUrl = &v
  return s
}

func (s *EcomVideoRecreationRequestInput) Validate() error {
  if s.ProductInfo != nil {
    if err := s.ProductInfo.Validate(); err != nil {
      return err
    }
  }
  return nil
}

type EcomVideoRecreationRequestInputProductInfo struct {
  // The product category.  
  // 
  // Example: Women\\"s Clothing/Sun Protection Jacket
  // 
  // example:
  // 
  // Women\\"s clothing/sun protection clothing
  Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
  // The actual product information (SKU, brand, color, material, size, specifications, logo, and usage), used to constrain voiceover facts.  
  // 
  // Example: Light moon yellow, cool-touch fabric, sun protection to the back of the hand, UPF50+
  // 
  // example:
  // 
  // Light moon yellow, cool-touch fabric, sun protection extending to the back of the hand, UPF50+
  Detail *string `json:"Detail,omitempty" xml:"Detail,omitempty"`
  // Required for product replacement. The name of the target product. Maximum length: 200 characters.  
  // 
  // Example: Light Moon Yellow Cool-touch Sun Protection Jacket
  // 
  // example:
  // 
  // Light moon yellow cool-touch sun protection clothing
  Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s EcomVideoRecreationRequestInputProductInfo) String() string {
  return dara.Prettify(s)
}

func (s EcomVideoRecreationRequestInputProductInfo) GoString() string {
  return s.String()
}

func (s *EcomVideoRecreationRequestInputProductInfo) GetCategory() *string  {
  return s.Category
}

func (s *EcomVideoRecreationRequestInputProductInfo) GetDetail() *string  {
  return s.Detail
}

func (s *EcomVideoRecreationRequestInputProductInfo) GetTitle() *string  {
  return s.Title
}

func (s *EcomVideoRecreationRequestInputProductInfo) SetCategory(v string) *EcomVideoRecreationRequestInputProductInfo {
  s.Category = &v
  return s
}

func (s *EcomVideoRecreationRequestInputProductInfo) SetDetail(v string) *EcomVideoRecreationRequestInputProductInfo {
  s.Detail = &v
  return s
}

func (s *EcomVideoRecreationRequestInputProductInfo) SetTitle(v string) *EcomVideoRecreationRequestInputProductInfo {
  s.Title = &v
  return s
}

func (s *EcomVideoRecreationRequestInputProductInfo) Validate() error {
  return dara.Validate(s)
}

type EcomVideoRecreationRequestOutput struct {
  // The target duration in seconds. `"auto"` (default): determined by the system. For product replacement, an integer from 5 to 60 can be specified. For person replacement, only `"auto"` is supported.
  // 
  // example:
  // 
  // auto
  Duration *int32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
  // The output resolution. Default value: `720p`.
  // 
  // example:
  // 
  // 720p、1080p
  Quality *string `json:"Quality,omitempty" xml:"Quality,omitempty"`
  // The output aspect ratio. Default value: `auto` (automatically matches the original video).
  // 
  // example:
  // 
  // auto、9:16、3:4、1:1、4:3、16:9
  Ratio *string `json:"Ratio,omitempty" xml:"Ratio,omitempty"`
}

func (s EcomVideoRecreationRequestOutput) String() string {
  return dara.Prettify(s)
}

func (s EcomVideoRecreationRequestOutput) GoString() string {
  return s.String()
}

func (s *EcomVideoRecreationRequestOutput) GetDuration() *int32  {
  return s.Duration
}

func (s *EcomVideoRecreationRequestOutput) GetQuality() *string  {
  return s.Quality
}

func (s *EcomVideoRecreationRequestOutput) GetRatio() *string  {
  return s.Ratio
}

func (s *EcomVideoRecreationRequestOutput) SetDuration(v int32) *EcomVideoRecreationRequestOutput {
  s.Duration = &v
  return s
}

func (s *EcomVideoRecreationRequestOutput) SetQuality(v string) *EcomVideoRecreationRequestOutput {
  s.Quality = &v
  return s
}

func (s *EcomVideoRecreationRequestOutput) SetRatio(v string) *EcomVideoRecreationRequestOutput {
  s.Ratio = &v
  return s
}

func (s *EcomVideoRecreationRequestOutput) Validate() error {
  return dara.Validate(s)
}

