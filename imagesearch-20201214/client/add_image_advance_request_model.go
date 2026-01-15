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
	// The category of the image. For more information, see [Category reference](https://help.aliyun.com/document_detail/179184.html).
	//
	// > 	- For product image search, if you specify a category for an image, the specified category prevails. If you do not specify a category for an image, the system predicts the category, and returns the prediction result in the response.
	//
	// >	- For generic image search, only 88888888 may be returned for this parameter in the response regardless of whether a category is specified.
	//
	// example:
	//
	// 88888888
	CategoryId *int32 `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	// Specifies whether to identify the subject in the image and search for images based on the subject identification result. Default value: true. Valid values:
	//
	// 	- true: The system identifies the subject in the image, and searches for images based on the subject identification result. The subject identification result is included in the response.
	//
	// 	- false: The system does not identify the subject in the image, and searches for images based on the entire image.
	//
	// example:
	//
	// true
	Crop *bool `json:"Crop,omitempty" xml:"Crop,omitempty"`
	// The user-defined content. The value can be up to 4,096 characters in length.
	//
	// > If you specify this parameter, the response includes this parameter and its value. You can add text such as an image description.
	//
	// example:
	//
	// zidingyi
	CustomContent *string `json:"CustomContent,omitempty" xml:"CustomContent,omitempty"`
	// The name of the Image Search instance. The name can be up to 20 characters in length. If an Image Search instance is purchased, you can log on to the [Image Search console](https://imagesearch.console.aliyun.com/) to view the instance. If no Image Search instance is purchased, you must purchase an instance. For more information, see [Activate Image Search](https://help.aliyun.com/document_detail/179178.html) and [Create an instance](https://help.aliyun.com/document_detail/66569.html).
	//
	// > The instance name is not the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// demoinstance1
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The attribute, which is an integer. The attribute can be used to filter images when you search for images. If you specify this parameter, the response includes this parameter and its value.
	//
	// example:
	//
	// 22
	IntAttr *int32 `json:"IntAttr,omitempty" xml:"IntAttr,omitempty"`
	// The attribute, which is an integer. The attribute can be used to filter images when you search for images. If you specify this parameter, the response includes this parameter and its value.
	//
	// example:
	//
	// 22
	IntAttr2 *int32 `json:"IntAttr2,omitempty" xml:"IntAttr2,omitempty"`
	IntAttr3 *int32 `json:"IntAttr3,omitempty" xml:"IntAttr3,omitempty"`
	IntAttr4 *int32 `json:"IntAttr4,omitempty" xml:"IntAttr4,omitempty"`
	// The image file. The image file is encoded in Base64.
	//
	// 	- The file size of the image cannot exceed 4 MB.
	//
	// 	- The following image formats are supported: PNG, JPG, JPEG, BMP, GIF, WebP, TIFF, and PPM.
	//
	// 	- The transmission timeout period cannot exceed 5 seconds.
	//
	// 	- For product and generic image searches, the length and width of the image must range from 100 pixels to 4,096 pixels.
	//
	// 	- The image cannot contain rotation settings.
	//
	// > 	- If you use SDKs to call this operation, you do not need to specify **PicContent**. The SDKs encapsulate this parameter and automatically encode its value in Base64. For more information about how to use Image Search SDK for Java, see [Java SDK](https://help.aliyun.com/document_detail/179188.html).
	//
	// >	- If you use OpenAPI Explorer to call this operation, you can select only the **2019-03-25*	- version. If you call this operation of other versions, the value of **PicContent*	- cannot be encoded in Base64.
	//
	// This parameter is required.
	//
	// example:
	//
	// AAAANSUhEUgAAAPcAAAEVCAYAAAA8d3NuAAAAAXNSR0IArs......RK5CYII=
	PicContentObject io.Reader `json:"PicContent,omitempty" xml:"PicContent,omitempty"`
	// The name of the image. The name can be up to 512 characters in length.
	//
	// > 	- An image is uniquely identified by the values of ProductId and PicName.
	//
	// >	- If you add an image whose product ID (ProductId) and image name (PicName) are the same as those of an existing image, the newly added image overwrites the existing image.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2092061_1.jpg
	PicName *string `json:"PicName,omitempty" xml:"PicName,omitempty"`
	// The ID of the product. The ID can be up to 512 characters in length.
	//
	// > A product may have multiple images.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2092061_1
	ProductId *string `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	// The subject area of the image, in the format of `x1,x2,y1,y2`. `x1 and y1` represent the position in the upper-left corner, in pixels. `x2 and y2` represent the position in the lower-right corner, in pixels.
	//
	// > 	- If you specify Region, the system searches for images based on the value of Region regardless of the value of Crop.
	//
	// >	- The value of Region does not have a unit. The value is generated based on the length and width of the image. If the length and width of the image are scaled, the value of Region must be proportionally adjusted.
	//
	// example:
	//
	// 280,486,232,351
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The attribute, which is a string. The value can be up to 128 characters in length. The attribute can be used to filter images when you search for images. If you specify this parameter, the response includes this parameter and its value.
	//
	// > The value cannot contain the following special characters: \\ ¥ $ & %
	//
	// example:
	//
	// ss
	StrAttr *string `json:"StrAttr,omitempty" xml:"StrAttr,omitempty"`
	// The attribute, which is a string. The value can be up to 128 characters in length. The attribute can be used to filter images when you search for images. If you specify this parameter, the response includes this parameter and its value.
	//
	// > The value cannot contain the following special characters: \\ ¥ $ & %
	//
	// example:
	//
	// ss
	StrAttr2 *string `json:"StrAttr2,omitempty" xml:"StrAttr2,omitempty"`
	StrAttr3 *string `json:"StrAttr3,omitempty" xml:"StrAttr3,omitempty"`
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
