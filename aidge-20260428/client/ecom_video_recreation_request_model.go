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
  // The input parameters for video variation generation.
  // 
  // This parameter is required.
  Input *EcomVideoRecreationRequestInput `json:"Input,omitempty" xml:"Input,omitempty" type:"Struct"`
  // The output specifications for the generated video.
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
  // The description or supplementary constraints for the target person in person replacement. The length is 1 to 500 characters. This parameter is required when PersonReferenceImageUrls is not provided.   
  // 
  // Example value: The target person is an adult male, retaining the original clothing and actions.
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
  // The URLs of the reference images for the target person in person replacement. You can provide 1 to 5 images, which must be of the same person. We recommend that you arrange the images in the order of facial close-up, front, 45-degree angle, side, and back.  
  // 
  // Example value: ["https://example.com/person.jpg"]
  PersonReferenceImageUrls []*string `json:"PersonReferenceImageUrls,omitempty" xml:"PersonReferenceImageUrls,omitempty" type:"Repeated"`
  // The URLs of the target product images. This parameter is required for product replacement. You can upload 1 to 8 images for the same SKU. We recommend that the main subject is clear, unobstructed, and has a clean background.   
  // 
  // Example value: ["https://example.com/product.png"]
  ProductImageUrls []*string `json:"ProductImageUrls,omitempty" xml:"ProductImageUrls,omitempty" type:"Repeated"`
  // The target product information. We recommend that you specify this parameter to improve the accuracy of the voiceover.
  ProductInfo *EcomVideoRecreationRequestInputProductInfo `json:"ProductInfo,omitempty" xml:"ProductInfo,omitempty" type:"Struct"`
  // The generation mode. Default value: `reference`. Valid values: `reference` (directly references the original video) and `rebuild` (breaks down the original video before regenerating it).
  // 
  // example:
  // 
  // reference
  RecreateMode *string `json:"RecreateMode,omitempty" xml:"RecreateMode,omitempty"`
  // The HTTP(S) URL of the reference video. The duration ranges from 2 to 360 seconds. The URL must remain accessible during task execution. We recommend that the validity period is no less than 24 hours.
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

func (s *EcomVideoRecreationRequestInput) GetRecreateMode() *string  {
  return s.RecreateMode
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

func (s *EcomVideoRecreationRequestInput) SetRecreateMode(v string) *EcomVideoRecreationRequestInput {
  s.RecreateMode = &v
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
  // Example value: Women\\"s clothing/Sun protection clothing.
  // 
  // example:
  // 
  // Women\\"s clothing/sun protection clothing
  Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
  // The actual product information, such as SKU, brand, color, material, size, specifications, logo, and usage instructions. This information is used to constrain the facts in the voiceover.  
  // 
  // Example value: Pale moon yellow, cool-touch fabric, sun protection down to the back of the hand, UPF50+
  // 
  // example:
  // 
  // Light moon yellow, cool-touch fabric, sun protection extending to the back of the hand, UPF50+
  Detail *string `json:"Detail,omitempty" xml:"Detail,omitempty"`
  // The name of the target product. This parameter is required for product replacement. The maximum length is 200 characters.  
  // 
  // Example value: Pale moon yellow cool-touch sun protection clothing.
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
  // The target duration in seconds. In `reference` mode, the input value is ignored and forced to -1, which follows the duration of the original video. In `rebuild` mode, you can pass an integer from 5 to 60, or pass -1 to indicate an automatic duration.
  // 
  // example:
  // 
  // -1
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

