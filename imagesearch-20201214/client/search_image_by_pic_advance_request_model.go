// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
	"io"
)

type iSearchImageByPicAdvanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCategoryId(v int32) *SearchImageByPicAdvanceRequest
	GetCategoryId() *int32
	SetCrop(v bool) *SearchImageByPicAdvanceRequest
	GetCrop() *bool
	SetDistinctProductId(v bool) *SearchImageByPicAdvanceRequest
	GetDistinctProductId() *bool
	SetFilter(v string) *SearchImageByPicAdvanceRequest
	GetFilter() *string
	SetInstanceName(v string) *SearchImageByPicAdvanceRequest
	GetInstanceName() *string
	SetNum(v int32) *SearchImageByPicAdvanceRequest
	GetNum() *int32
	SetPicContentObject(v io.Reader) *SearchImageByPicAdvanceRequest
	GetPicContentObject() io.Reader
	SetRegion(v string) *SearchImageByPicAdvanceRequest
	GetRegion() *string
	SetScoreThreshold(v string) *SearchImageByPicAdvanceRequest
	GetScoreThreshold() *string
	SetStart(v int32) *SearchImageByPicAdvanceRequest
	GetStart() *int32
}

type SearchImageByPicAdvanceRequest struct {
	// The product category. For more information, see [Category reference](https://help.aliyun.com/document_detail/179184.html).
	//
	//  - For product image search, if you specify a category, the specified category is used. If you do not specify a category, the system predicts the category. You can obtain the predicted category from the response.
	//
	// <props="china">
	//
	//  - For fabric, trademark, generic furniture, and industrial hardware image search, the system sets the category to 88888888 regardless of whether you specify a category.
	//
	// <props="intl">
	//
	//  - For generic image search, the system sets the category to 88888888 regardless of whether you specify a category.
	//
	// .
	//
	// example:
	//
	// 88888888
	CategoryId *int32 `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	// Specifies whether to perform subject identification. Default value: true.
	//
	//  - If this parameter is set to true, the system performs subject identification and searches based on the identified subject. You can obtain the subject identification result from the response.
	//
	//  - If this parameter is set to false, the system does not perform subject identification and searches based on the entire image.
	//
	// <props="china">
	//
	// - For fabric image search, this parameter is ignored and the system searches based on the entire image.
	//
	// .
	//
	// example:
	//
	// true
	Crop *bool `json:"Crop,omitempty" xml:"Crop,omitempty"`
	// Specifies whether to deduplicate results based on the ProductId field during the search. If this parameter is set to true, deduplication is performed.
	//
	// example:
	//
	// false
	DistinctProductId *bool `json:"DistinctProductId,omitempty" xml:"DistinctProductId,omitempty"`
	// The filter condition. The int_attr field supports the following operators: in, not in, greater than (>), greater than or equal to (>=), less than (<), less than or equal to (<=), and equal to (=). The str_attr field supports the following operators: in, not in, equal to (=), and not equal to (!=). Multiple conditions can be connected by using AND and OR.
	//
	// Examples:
	//
	// - int_attr >= 100.
	//
	// - str_attr != "value1".
	//
	// - int_attr = 1000 AND str_attr = "value1".
	//
	// >A maximum of 4096 characters are supported.
	//
	// example:
	//
	// int_attr=1000 AND str_attr="value1"
	Filter *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	// The name of the Image Search instance. The name can be up to 20 characters in length.
	//
	// If you have purchased an Image Search instance, log on to the [Image Search console](https://imagesearch.console.aliyun.com/) to view the instance name.
	//
	// If you have not purchased an Image Search instance, see [Activate the service](https://help.aliyun.com/document_detail/179178.html) and [Create an instance](https://help.aliyun.com/document_detail/66569.html).
	//
	// >The instance name is not the instance ID. Make sure that you distinguish between them.
	//
	// This parameter is required.
	//
	// example:
	//
	// demoinstance1
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The number of results to return. Valid values: 1 to 100. Default value: 10.
	//
	// example:
	//
	// 10
	Num *int32 `json:"Num,omitempty" xml:"Num,omitempty"`
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
	//  - For product image search, generic image search, furniture image search, and industrial hardware image search, the image width and height must be greater than or equal to 100 px and less than or equal to 4096 px.
	//
	// For trademark image search, the image width and height must be greater than or equal to 200 px and less than 4096 px.
	//
	// For fabric image search, the image width and height must be greater than or equal to 448 px and less than or equal to 4096 px.
	//
	// <props="intl">
	//
	//  - For product image search and generic image search, the image width and height must be greater than or equal to 100 px and less than or equal to 4096 px.
	//
	// - The image cannot contain rotation information.
	//
	// > - **Call by using the SDK:**
	//
	//   - If you use the V3 SDK, you do not need to specify the PicContent field. The SDK encapsulates this field as PicContentObject and automatically converts it to Base64 encoding. For specific examples, see [Java SDK](https://help.aliyun.com/document_detail/179188.html).
	//
	//   - The SDK does not support directly passing an image URL. The V3 SDK provides an alternative method to upload images by URL. For specific examples, see [Java SDK](https://help.aliyun.com/document_detail/179188.html).
	//
	// - **Call by using the Alibaba Cloud OpenAPI platform:**
	//
	//   - If you use the **2019-03-25*	- version, set the **PicContent*	- field to the **Base64**-encoded string of the image.
	//
	//   - If you use the **2020-12-14*	- version, click the upload button in the **PicContent*	- field to upload the image.
	//
	// This parameter is required.
	//
	// example:
	//
	// AAAANSUhEUgAAAPcAAAEVCAYAAAA8d3NuAAAAAXNSR0IArs......RK5CYII=
	PicContentObject io.Reader `json:"PicContent,omitempty" xml:"PicContent,omitempty"`
	// The subject region of the image, in the format of `x1,x2,y1,y2`, where `x1,y1` is the upper-left point and `x2,y2` is the lower-right point.
	//
	// > - If you specify Region, the system searches based on the specified Region regardless of the value of the Crop parameter.
	//
	// <props="china">
	//
	// - For fabric image search, this parameter is ignored and the system searches based on the entire image.
	//
	// .
	//
	// example:
	//
	// 280,486,232,351
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The similarity score threshold. After you specify this threshold, only images with a similarity score greater than or equal to the threshold are returned. Valid values: 0.00 to 1.00. Up to two decimal places are supported. Default value: 0.00.
	//
	// example:
	//
	// 0.50
	ScoreThreshold *string `json:"ScoreThreshold,omitempty" xml:"ScoreThreshold,omitempty"`
	// The start position of the results to return. Valid values: 0 to 499. Default value: 0.
	//
	// example:
	//
	// 0
	Start *int32 `json:"Start,omitempty" xml:"Start,omitempty"`
}

func (s SearchImageByPicAdvanceRequest) String() string {
	return dara.Prettify(s)
}

func (s SearchImageByPicAdvanceRequest) GoString() string {
	return s.String()
}

func (s *SearchImageByPicAdvanceRequest) GetCategoryId() *int32 {
	return s.CategoryId
}

func (s *SearchImageByPicAdvanceRequest) GetCrop() *bool {
	return s.Crop
}

func (s *SearchImageByPicAdvanceRequest) GetDistinctProductId() *bool {
	return s.DistinctProductId
}

func (s *SearchImageByPicAdvanceRequest) GetFilter() *string {
	return s.Filter
}

func (s *SearchImageByPicAdvanceRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *SearchImageByPicAdvanceRequest) GetNum() *int32 {
	return s.Num
}

func (s *SearchImageByPicAdvanceRequest) GetPicContentObject() io.Reader {
	return s.PicContentObject
}

func (s *SearchImageByPicAdvanceRequest) GetRegion() *string {
	return s.Region
}

func (s *SearchImageByPicAdvanceRequest) GetScoreThreshold() *string {
	return s.ScoreThreshold
}

func (s *SearchImageByPicAdvanceRequest) GetStart() *int32 {
	return s.Start
}

func (s *SearchImageByPicAdvanceRequest) SetCategoryId(v int32) *SearchImageByPicAdvanceRequest {
	s.CategoryId = &v
	return s
}

func (s *SearchImageByPicAdvanceRequest) SetCrop(v bool) *SearchImageByPicAdvanceRequest {
	s.Crop = &v
	return s
}

func (s *SearchImageByPicAdvanceRequest) SetDistinctProductId(v bool) *SearchImageByPicAdvanceRequest {
	s.DistinctProductId = &v
	return s
}

func (s *SearchImageByPicAdvanceRequest) SetFilter(v string) *SearchImageByPicAdvanceRequest {
	s.Filter = &v
	return s
}

func (s *SearchImageByPicAdvanceRequest) SetInstanceName(v string) *SearchImageByPicAdvanceRequest {
	s.InstanceName = &v
	return s
}

func (s *SearchImageByPicAdvanceRequest) SetNum(v int32) *SearchImageByPicAdvanceRequest {
	s.Num = &v
	return s
}

func (s *SearchImageByPicAdvanceRequest) SetPicContentObject(v io.Reader) *SearchImageByPicAdvanceRequest {
	s.PicContentObject = v
	return s
}

func (s *SearchImageByPicAdvanceRequest) SetRegion(v string) *SearchImageByPicAdvanceRequest {
	s.Region = &v
	return s
}

func (s *SearchImageByPicAdvanceRequest) SetScoreThreshold(v string) *SearchImageByPicAdvanceRequest {
	s.ScoreThreshold = &v
	return s
}

func (s *SearchImageByPicAdvanceRequest) SetStart(v int32) *SearchImageByPicAdvanceRequest {
	s.Start = &v
	return s
}

func (s *SearchImageByPicAdvanceRequest) Validate() error {
	return dara.Validate(s)
}
