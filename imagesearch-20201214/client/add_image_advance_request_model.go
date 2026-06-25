// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
	"io"
)

type iAddImageAdvanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCategoryId(v int32) *AddImageAdvanceRequest
	GetCategoryId() *int32
	SetCrop(v bool) *AddImageAdvanceRequest
	GetCrop() *bool
	SetCustomContent(v string) *AddImageAdvanceRequest
	GetCustomContent() *string
	SetInstanceName(v string) *AddImageAdvanceRequest
	GetInstanceName() *string
	SetIntAttr(v int32) *AddImageAdvanceRequest
	GetIntAttr() *int32
	SetIntAttr2(v int32) *AddImageAdvanceRequest
	GetIntAttr2() *int32
	SetIntAttr3(v int32) *AddImageAdvanceRequest
	GetIntAttr3() *int32
	SetIntAttr4(v int32) *AddImageAdvanceRequest
	GetIntAttr4() *int32
	SetPicContentObject(v io.Reader) *AddImageAdvanceRequest
	GetPicContentObject() io.Reader
	SetPicName(v string) *AddImageAdvanceRequest
	GetPicName() *string
	SetProductId(v string) *AddImageAdvanceRequest
	GetProductId() *string
	SetRegion(v string) *AddImageAdvanceRequest
	GetRegion() *string
	SetStrAttr(v string) *AddImageAdvanceRequest
	GetStrAttr() *string
	SetStrAttr2(v string) *AddImageAdvanceRequest
	GetStrAttr2() *string
	SetStrAttr3(v string) *AddImageAdvanceRequest
	GetStrAttr3() *string
	SetStrAttr4(v string) *AddImageAdvanceRequest
	GetStrAttr4() *string
}

type AddImageAdvanceRequest struct {
	// The image category. For more information, refer to [Category reference](https://help.aliyun.com/document_detail/179184.html).
	//
	// > - For product image search, if you specify a category, the specified category is used. If you do not specify a category, the system predicts the category. The predicted category result can be obtained from the response.
	//
	// <props="china">
	//
	// - For fabric, trademark, generic, furniture, and industrial hardware image search, the system sets the category to 88888888 regardless of whether you specify a category.
	//
	// - For generic image search, the system sets the category to 88888888 regardless of whether you specify a category.
	//
	// example:
	//
	// 88888888
	CategoryId *int32 `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	// Specifies whether to perform subject identification. Default value: true.
	//
	//  - true: The system performs subject identification and searches based on the identified subject. The subject identification result can be obtained from the response.
	//
	// - false: The system does not perform subject identification and searches based on the entire image.
	//
	// <props="china">For fabric image search, this parameter is ignored. The system searches based on the entire image..
	//
	// example:
	//
	// true
	Crop *bool `json:"Crop,omitempty" xml:"Crop,omitempty"`
	// The custom content defined by the user. The content can be up to 4,096 characters in length.
	//
	// >This field is returned in query results. For example, you can add a text description of the image.
	//
	// example:
	//
	// zidingyi
	CustomContent *string `json:"CustomContent,omitempty" xml:"CustomContent,omitempty"`
	// The name of the Image Search instance. The name can be up to 20 characters in length.
	//
	// If you have purchased an Image Search instance, log on to the [Image Search console](https://imagesearch.console.aliyun.com/) to view the instance name.
	//
	// If you have not purchased an Image Search instance, refer to [Activate the service](https://help.aliyun.com/document_detail/179178.html) and [Create an instance](https://help.aliyun.com/document_detail/66569.html).
	//
	// >The instance name is not the instance ID. Do not confuse the two.
	//
	// This parameter is required.
	//
	// example:
	//
	// demoinstance1
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The integer attribute. This attribute can be used to filter query results and is returned in query results.
	//
	// example:
	//
	// 22
	IntAttr *int32 `json:"IntAttr,omitempty" xml:"IntAttr,omitempty"`
	// The integer attribute. This attribute can be used to filter query results and is returned in query results.
	//
	// example:
	//
	// 22
	IntAttr2 *int32 `json:"IntAttr2,omitempty" xml:"IntAttr2,omitempty"`
	// The integer attribute. This attribute can be used to filter query results and is returned in query results.
	//
	// example:
	//
	// 33
	IntAttr3 *int32 `json:"IntAttr3,omitempty" xml:"IntAttr3,omitempty"`
	// The integer attribute. This attribute can be used to filter query results and is returned in query results.
	//
	// example:
	//
	// 44
	IntAttr4 *int32 `json:"IntAttr4,omitempty" xml:"IntAttr4,omitempty"`
	// The image content.
	//
	//  - The image size cannot exceed 4 MB.
	//
	//  - Image formats: PNG, JPG, JPEG, BMP, GIF, WEBP, TIFF, and PPM.
	//
	//  - The transmission wait time cannot exceed 5 seconds.
	//
	// <props="china">
	//
	//  - For product image search, generic image search, furniture image search, and industrial hardware image search, the image width and height must be at least 100 pixels and at most 4,096 pixels.
	//
	//   For trademark image search, the image width and height must be at least 200 pixels and less than 4,096 pixels.
	//
	//  For fabric image search, the image width and height must be at least 448 pixels and at most 4,096 pixels.
	//
	// <props="intl">
	//
	//  - For product image search and generic image search, the image width and height must be at least 100 pixels and at most 4,096 pixels.
	//
	// - The image must not contain rotation information.
	//
	// > - **When calling by using an SDK:**
	//
	//   - If you use a V3 SDK, you do not need to set the PicContent field. The SDK encapsulates this field as PicContentObject and automatically converts it to Base64 encoding. For specific examples, refer to [Java SDK](https://help.aliyun.com/document_detail/179188.html).
	//
	//   - The SDK does not support passing image URLs directly. The V3 SDK provides an alternative method to upload images by URL. For specific examples, refer to [Java SDK](https://help.aliyun.com/document_detail/179188.html).
	//
	// - **When calling by using the Alibaba Cloud OpenAPI platform:**
	//
	//   - If you use the **2019-03-25*	- version, set the **PicContent*	- field to the **Base64*	- encoding of the image.
	//
	//   - If you use the **2020-12-14*	- version, click to upload the image directly in the **PicContent*	- field.
	//
	// This parameter is required.
	//
	// example:
	//
	// AAAANSUhEUgAAAPcAAAEVCAYAAAA8d3NuAAAAAXNSR0IArs......RK5CYII=
	PicContentObject io.Reader `json:"PicContent,omitempty" xml:"PicContent,omitempty"`
	// The image name. The name can be up to 256 characters in length.
	//
	// > - ProductId and PicName uniquely identify an image.
	//
	// - If you add multiple images with the same ProductId and PicName, only the last added image is retained. Previously added images are overwritten.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2092061_1.jpg
	PicName *string `json:"PicName,omitempty" xml:"PicName,omitempty"`
	// The product ID. The ID can be up to 256 characters in length.
	//
	// >A product can have multiple images.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2092061_1
	ProductId *string `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	// The subject region of the image, in the format of `x1,x2,y1,y2`, where `x1,y1` is the upper-left point and `x2,y2` is the lower-right point.
	//
	// > - If you specify Region, the system searches based on the specified Region regardless of the Crop parameter value.
	//
	// <props="china">
	//
	// - For fabric image search, this parameter is ignored. The system searches based on the entire image.
	//
	// - The Region parameter has no unit. The values are based on the pixel dimensions of the image. If the image is scaled, the Region parameter values must be scaled proportionally.
	//
	// example:
	//
	// 280,486,232,351
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The string attribute. The attribute can be up to 128 characters in length. This attribute can be used to filter query results and is returned in query results.
	//
	// > Special characters such as \\¥$&% are not supported.
	//
	// example:
	//
	// ss
	StrAttr *string `json:"StrAttr,omitempty" xml:"StrAttr,omitempty"`
	// The string attribute. The attribute can be up to 128 characters in length. This attribute can be used to filter query results and is returned in query results.
	//
	// > Special characters such as \\¥$&% are not supported.
	//
	// example:
	//
	// ss
	StrAttr2 *string `json:"StrAttr2,omitempty" xml:"StrAttr2,omitempty"`
	// The string attribute. The attribute can be up to 128 characters in length. This attribute can be used to filter query results and is returned in query results.
	//
	// > Special characters such as \\¥$&% are not supported.
	//
	// example:
	//
	// ss
	StrAttr3 *string `json:"StrAttr3,omitempty" xml:"StrAttr3,omitempty"`
	// The string attribute. The attribute can be up to 128 characters in length. This attribute can be used to filter query results and is returned in query results.
	//
	// > Special characters such as \\¥$&% are not supported.
	//
	// example:
	//
	// ss
	StrAttr4 *string `json:"StrAttr4,omitempty" xml:"StrAttr4,omitempty"`
}

func (s AddImageAdvanceRequest) String() string {
	return dara.Prettify(s)
}

func (s AddImageAdvanceRequest) GoString() string {
	return s.String()
}

func (s *AddImageAdvanceRequest) GetCategoryId() *int32 {
	return s.CategoryId
}

func (s *AddImageAdvanceRequest) GetCrop() *bool {
	return s.Crop
}

func (s *AddImageAdvanceRequest) GetCustomContent() *string {
	return s.CustomContent
}

func (s *AddImageAdvanceRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *AddImageAdvanceRequest) GetIntAttr() *int32 {
	return s.IntAttr
}

func (s *AddImageAdvanceRequest) GetIntAttr2() *int32 {
	return s.IntAttr2
}

func (s *AddImageAdvanceRequest) GetIntAttr3() *int32 {
	return s.IntAttr3
}

func (s *AddImageAdvanceRequest) GetIntAttr4() *int32 {
	return s.IntAttr4
}

func (s *AddImageAdvanceRequest) GetPicContentObject() io.Reader {
	return s.PicContentObject
}

func (s *AddImageAdvanceRequest) GetPicName() *string {
	return s.PicName
}

func (s *AddImageAdvanceRequest) GetProductId() *string {
	return s.ProductId
}

func (s *AddImageAdvanceRequest) GetRegion() *string {
	return s.Region
}

func (s *AddImageAdvanceRequest) GetStrAttr() *string {
	return s.StrAttr
}

func (s *AddImageAdvanceRequest) GetStrAttr2() *string {
	return s.StrAttr2
}

func (s *AddImageAdvanceRequest) GetStrAttr3() *string {
	return s.StrAttr3
}

func (s *AddImageAdvanceRequest) GetStrAttr4() *string {
	return s.StrAttr4
}

func (s *AddImageAdvanceRequest) SetCategoryId(v int32) *AddImageAdvanceRequest {
	s.CategoryId = &v
	return s
}

func (s *AddImageAdvanceRequest) SetCrop(v bool) *AddImageAdvanceRequest {
	s.Crop = &v
	return s
}

func (s *AddImageAdvanceRequest) SetCustomContent(v string) *AddImageAdvanceRequest {
	s.CustomContent = &v
	return s
}

func (s *AddImageAdvanceRequest) SetInstanceName(v string) *AddImageAdvanceRequest {
	s.InstanceName = &v
	return s
}

func (s *AddImageAdvanceRequest) SetIntAttr(v int32) *AddImageAdvanceRequest {
	s.IntAttr = &v
	return s
}

func (s *AddImageAdvanceRequest) SetIntAttr2(v int32) *AddImageAdvanceRequest {
	s.IntAttr2 = &v
	return s
}

func (s *AddImageAdvanceRequest) SetIntAttr3(v int32) *AddImageAdvanceRequest {
	s.IntAttr3 = &v
	return s
}

func (s *AddImageAdvanceRequest) SetIntAttr4(v int32) *AddImageAdvanceRequest {
	s.IntAttr4 = &v
	return s
}

func (s *AddImageAdvanceRequest) SetPicContentObject(v io.Reader) *AddImageAdvanceRequest {
	s.PicContentObject = v
	return s
}

func (s *AddImageAdvanceRequest) SetPicName(v string) *AddImageAdvanceRequest {
	s.PicName = &v
	return s
}

func (s *AddImageAdvanceRequest) SetProductId(v string) *AddImageAdvanceRequest {
	s.ProductId = &v
	return s
}

func (s *AddImageAdvanceRequest) SetRegion(v string) *AddImageAdvanceRequest {
	s.Region = &v
	return s
}

func (s *AddImageAdvanceRequest) SetStrAttr(v string) *AddImageAdvanceRequest {
	s.StrAttr = &v
	return s
}

func (s *AddImageAdvanceRequest) SetStrAttr2(v string) *AddImageAdvanceRequest {
	s.StrAttr2 = &v
	return s
}

func (s *AddImageAdvanceRequest) SetStrAttr3(v string) *AddImageAdvanceRequest {
	s.StrAttr3 = &v
	return s
}

func (s *AddImageAdvanceRequest) SetStrAttr4(v string) *AddImageAdvanceRequest {
	s.StrAttr4 = &v
	return s
}

func (s *AddImageAdvanceRequest) Validate() error {
	return dara.Validate(s)
}
