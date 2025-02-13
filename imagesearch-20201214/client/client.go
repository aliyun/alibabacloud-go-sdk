// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	openplatform "github.com/alibabacloud-go/openplatform-20191219/v2/client"
	fileform "github.com/alibabacloud-go/tea-fileform/service"
	oss "github.com/alibabacloud-go/tea-oss-sdk/client"
	ossutil "github.com/alibabacloud-go/tea-oss-utils/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
	"io"
)

type AddImageRequest struct {
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
	PicContent *string `json:"PicContent,omitempty" xml:"PicContent,omitempty"`
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

func (s AddImageRequest) String() string {
	return tea.Prettify(s)
}

func (s AddImageRequest) GoString() string {
	return s.String()
}

func (s *AddImageRequest) SetCategoryId(v int32) *AddImageRequest {
	s.CategoryId = &v
	return s
}

func (s *AddImageRequest) SetCrop(v bool) *AddImageRequest {
	s.Crop = &v
	return s
}

func (s *AddImageRequest) SetCustomContent(v string) *AddImageRequest {
	s.CustomContent = &v
	return s
}

func (s *AddImageRequest) SetInstanceName(v string) *AddImageRequest {
	s.InstanceName = &v
	return s
}

func (s *AddImageRequest) SetIntAttr(v int32) *AddImageRequest {
	s.IntAttr = &v
	return s
}

func (s *AddImageRequest) SetIntAttr2(v int32) *AddImageRequest {
	s.IntAttr2 = &v
	return s
}

func (s *AddImageRequest) SetIntAttr3(v int32) *AddImageRequest {
	s.IntAttr3 = &v
	return s
}

func (s *AddImageRequest) SetIntAttr4(v int32) *AddImageRequest {
	s.IntAttr4 = &v
	return s
}

func (s *AddImageRequest) SetPicContent(v string) *AddImageRequest {
	s.PicContent = &v
	return s
}

func (s *AddImageRequest) SetPicName(v string) *AddImageRequest {
	s.PicName = &v
	return s
}

func (s *AddImageRequest) SetProductId(v string) *AddImageRequest {
	s.ProductId = &v
	return s
}

func (s *AddImageRequest) SetRegion(v string) *AddImageRequest {
	s.Region = &v
	return s
}

func (s *AddImageRequest) SetStrAttr(v string) *AddImageRequest {
	s.StrAttr = &v
	return s
}

func (s *AddImageRequest) SetStrAttr2(v string) *AddImageRequest {
	s.StrAttr2 = &v
	return s
}

func (s *AddImageRequest) SetStrAttr3(v string) *AddImageRequest {
	s.StrAttr3 = &v
	return s
}

func (s *AddImageRequest) SetStrAttr4(v string) *AddImageRequest {
	s.StrAttr4 = &v
	return s
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
	return tea.Prettify(s)
}

func (s AddImageAdvanceRequest) GoString() string {
	return s.String()
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

type AddImageResponseBody struct {
	// The code returned.
	//
	// 	- A value of 0 indicates that the request was successful.
	//
	// 	- Values other than 0 indicate that the request failed.
	//
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message returned if the request failed.
	//
	// > No value is returned if the request was successful, and an error message is returned if the request failed.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The results of category prediction and subject identification.
	PicInfo *AddImageResponseBodyPicInfo `json:"PicInfo,omitempty" xml:"PicInfo,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// E0845DE6-52AF-4B50-9F15-51ED4044E6AB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddImageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddImageResponseBody) GoString() string {
	return s.String()
}

func (s *AddImageResponseBody) SetCode(v int32) *AddImageResponseBody {
	s.Code = &v
	return s
}

func (s *AddImageResponseBody) SetMessage(v string) *AddImageResponseBody {
	s.Message = &v
	return s
}

func (s *AddImageResponseBody) SetPicInfo(v *AddImageResponseBodyPicInfo) *AddImageResponseBody {
	s.PicInfo = v
	return s
}

func (s *AddImageResponseBody) SetRequestId(v string) *AddImageResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddImageResponseBody) SetSuccess(v bool) *AddImageResponseBody {
	s.Success = &v
	return s
}

type AddImageResponseBodyPicInfo struct {
	// The result of category prediction. If a category is specified in the request, the specified category prevails.
	//
	// example:
	//
	// 88888888
	CategoryId *int32 `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	// The result of subject identification. The subject area of the image is in the format of `x1,x2,y1,y2`. `x1 and y1` represent the position in the upper-left corner, in pixels. `x2 and y2` represent the position in the lower-right corner, in pixels. If a subject area is specified in the request, the specified subject area prevails.
	//
	// example:
	//
	// 94,691,206,650
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s AddImageResponseBodyPicInfo) String() string {
	return tea.Prettify(s)
}

func (s AddImageResponseBodyPicInfo) GoString() string {
	return s.String()
}

func (s *AddImageResponseBodyPicInfo) SetCategoryId(v int32) *AddImageResponseBodyPicInfo {
	s.CategoryId = &v
	return s
}

func (s *AddImageResponseBodyPicInfo) SetRegion(v string) *AddImageResponseBodyPicInfo {
	s.Region = &v
	return s
}

type AddImageResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddImageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddImageResponse) String() string {
	return tea.Prettify(s)
}

func (s AddImageResponse) GoString() string {
	return s.String()
}

func (s *AddImageResponse) SetHeaders(v map[string]*string) *AddImageResponse {
	s.Headers = v
	return s
}

func (s *AddImageResponse) SetStatusCode(v int32) *AddImageResponse {
	s.StatusCode = &v
	return s
}

func (s *AddImageResponse) SetBody(v *AddImageResponseBody) *AddImageResponse {
	s.Body = v
	return s
}

type CompareSimilarByImageRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// demoinstance1
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// AAAANSUhEUgAAAPcAAAEVCAYAAAA8d3NuAAAAAXNSR0IArs......RK5CYII=
	PrimaryPicContent *string `json:"PrimaryPicContent,omitempty" xml:"PrimaryPicContent,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// AAAANSUhEUgAAAPcAAAEVCAYAAAA8d3NuAAAAAXNSR0IArs......RK5CYII=
	SecondaryPicContent *string `json:"SecondaryPicContent,omitempty" xml:"SecondaryPicContent,omitempty"`
}

func (s CompareSimilarByImageRequest) String() string {
	return tea.Prettify(s)
}

func (s CompareSimilarByImageRequest) GoString() string {
	return s.String()
}

func (s *CompareSimilarByImageRequest) SetInstanceName(v string) *CompareSimilarByImageRequest {
	s.InstanceName = &v
	return s
}

func (s *CompareSimilarByImageRequest) SetPrimaryPicContent(v string) *CompareSimilarByImageRequest {
	s.PrimaryPicContent = &v
	return s
}

func (s *CompareSimilarByImageRequest) SetSecondaryPicContent(v string) *CompareSimilarByImageRequest {
	s.SecondaryPicContent = &v
	return s
}

type CompareSimilarByImageAdvanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// demoinstance1
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// AAAANSUhEUgAAAPcAAAEVCAYAAAA8d3NuAAAAAXNSR0IArs......RK5CYII=
	PrimaryPicContentObject io.Reader `json:"PrimaryPicContent,omitempty" xml:"PrimaryPicContent,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// AAAANSUhEUgAAAPcAAAEVCAYAAAA8d3NuAAAAAXNSR0IArs......RK5CYII=
	SecondaryPicContentObject io.Reader `json:"SecondaryPicContent,omitempty" xml:"SecondaryPicContent,omitempty"`
}

func (s CompareSimilarByImageAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s CompareSimilarByImageAdvanceRequest) GoString() string {
	return s.String()
}

func (s *CompareSimilarByImageAdvanceRequest) SetInstanceName(v string) *CompareSimilarByImageAdvanceRequest {
	s.InstanceName = &v
	return s
}

func (s *CompareSimilarByImageAdvanceRequest) SetPrimaryPicContentObject(v io.Reader) *CompareSimilarByImageAdvanceRequest {
	s.PrimaryPicContentObject = v
	return s
}

func (s *CompareSimilarByImageAdvanceRequest) SetSecondaryPicContentObject(v io.Reader) *CompareSimilarByImageAdvanceRequest {
	s.SecondaryPicContentObject = v
	return s
}

type CompareSimilarByImageResponseBody struct {
	AccessDeniedDetail *CompareSimilarByImageResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// success
	Msg *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	// example:
	//
	// B3137727-7D6E-488C-BA21-0E034C38A879
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 0.85
	Score *float64 `json:"Score,omitempty" xml:"Score,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CompareSimilarByImageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CompareSimilarByImageResponseBody) GoString() string {
	return s.String()
}

func (s *CompareSimilarByImageResponseBody) SetAccessDeniedDetail(v *CompareSimilarByImageResponseBodyAccessDeniedDetail) *CompareSimilarByImageResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *CompareSimilarByImageResponseBody) SetCode(v int32) *CompareSimilarByImageResponseBody {
	s.Code = &v
	return s
}

func (s *CompareSimilarByImageResponseBody) SetMsg(v string) *CompareSimilarByImageResponseBody {
	s.Msg = &v
	return s
}

func (s *CompareSimilarByImageResponseBody) SetRequestId(v string) *CompareSimilarByImageResponseBody {
	s.RequestId = &v
	return s
}

func (s *CompareSimilarByImageResponseBody) SetScore(v float64) *CompareSimilarByImageResponseBody {
	s.Score = &v
	return s
}

func (s *CompareSimilarByImageResponseBody) SetSuccess(v bool) *CompareSimilarByImageResponseBody {
	s.Success = &v
	return s
}

type CompareSimilarByImageResponseBodyAccessDeniedDetail struct {
	AuthAction               *string `json:"AuthAction,omitempty" xml:"AuthAction,omitempty"`
	AuthPrincipalDisplayName *string `json:"AuthPrincipalDisplayName,omitempty" xml:"AuthPrincipalDisplayName,omitempty"`
	AuthPrincipalOwnerId     *string `json:"AuthPrincipalOwnerId,omitempty" xml:"AuthPrincipalOwnerId,omitempty"`
	AuthPrincipalType        *string `json:"AuthPrincipalType,omitempty" xml:"AuthPrincipalType,omitempty"`
	EncodedDiagnosticMessage *string `json:"EncodedDiagnosticMessage,omitempty" xml:"EncodedDiagnosticMessage,omitempty"`
	NoPermissionType         *string `json:"NoPermissionType,omitempty" xml:"NoPermissionType,omitempty"`
	PolicyType               *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s CompareSimilarByImageResponseBodyAccessDeniedDetail) String() string {
	return tea.Prettify(s)
}

func (s CompareSimilarByImageResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *CompareSimilarByImageResponseBodyAccessDeniedDetail) SetAuthAction(v string) *CompareSimilarByImageResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *CompareSimilarByImageResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *CompareSimilarByImageResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *CompareSimilarByImageResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *CompareSimilarByImageResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *CompareSimilarByImageResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *CompareSimilarByImageResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *CompareSimilarByImageResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *CompareSimilarByImageResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *CompareSimilarByImageResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *CompareSimilarByImageResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *CompareSimilarByImageResponseBodyAccessDeniedDetail) SetPolicyType(v string) *CompareSimilarByImageResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

type CompareSimilarByImageResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CompareSimilarByImageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CompareSimilarByImageResponse) String() string {
	return tea.Prettify(s)
}

func (s CompareSimilarByImageResponse) GoString() string {
	return s.String()
}

func (s *CompareSimilarByImageResponse) SetHeaders(v map[string]*string) *CompareSimilarByImageResponse {
	s.Headers = v
	return s
}

func (s *CompareSimilarByImageResponse) SetStatusCode(v int32) *CompareSimilarByImageResponse {
	s.StatusCode = &v
	return s
}

func (s *CompareSimilarByImageResponse) SetBody(v *CompareSimilarByImageResponseBody) *CompareSimilarByImageResponse {
	s.Body = v
	return s
}

type DeleteImageRequest struct {
	Filter *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	// The name of the Image Search instance. The name can be up to 20 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// demoinstance1
	InstanceName     *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	IsDeleteByFilter *bool   `json:"IsDeleteByFilter,omitempty" xml:"IsDeleteByFilter,omitempty"`
	// The name of the image.
	//
	// 	- If this parameter is not set, the system deletes all the images that correspond to the specified ProductId parameter.
	//
	// 	- If this parameter is set, the system deletes only the image that is specified by the ProductId and PicName parameters.
	//
	// example:
	//
	// 2092061_1.jpg
	PicName *string `json:"PicName,omitempty" xml:"PicName,omitempty"`
	// The ID of the commodity.
	//
	// >  A commodity may have multiple images.
	//
	// example:
	//
	// 2092061_1
	ProductId *string `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
}

func (s DeleteImageRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteImageRequest) GoString() string {
	return s.String()
}

func (s *DeleteImageRequest) SetFilter(v string) *DeleteImageRequest {
	s.Filter = &v
	return s
}

func (s *DeleteImageRequest) SetInstanceName(v string) *DeleteImageRequest {
	s.InstanceName = &v
	return s
}

func (s *DeleteImageRequest) SetIsDeleteByFilter(v bool) *DeleteImageRequest {
	s.IsDeleteByFilter = &v
	return s
}

func (s *DeleteImageRequest) SetPicName(v string) *DeleteImageRequest {
	s.PicName = &v
	return s
}

func (s *DeleteImageRequest) SetProductId(v string) *DeleteImageRequest {
	s.ProductId = &v
	return s
}

type DeleteImageResponseBody struct {
	// The error code returned.
	//
	// 	- A value of 0 indicates that the operation is successful.
	//
	// 	- Values other than 0 indicate errors.
	//
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The information about the deleted images.
	Data *DeleteImageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 0703956F-9BCC-48FA-99F7-96C0BF449C69
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteImageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteImageResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteImageResponseBody) SetCode(v int32) *DeleteImageResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteImageResponseBody) SetData(v *DeleteImageResponseBodyData) *DeleteImageResponseBody {
	s.Data = v
	return s
}

func (s *DeleteImageResponseBody) SetMessage(v string) *DeleteImageResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteImageResponseBody) SetRequestId(v string) *DeleteImageResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteImageResponseBody) SetSuccess(v bool) *DeleteImageResponseBody {
	s.Success = &v
	return s
}

type DeleteImageResponseBodyData struct {
	// The name (PicName) of the deleted image.
	//
	// example:
	//
	// 5555.jpg
	PicNames []*string `json:"PicNames,omitempty" xml:"PicNames,omitempty" type:"Repeated"`
}

func (s DeleteImageResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DeleteImageResponseBodyData) GoString() string {
	return s.String()
}

func (s *DeleteImageResponseBodyData) SetPicNames(v []*string) *DeleteImageResponseBodyData {
	s.PicNames = v
	return s
}

type DeleteImageResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteImageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteImageResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteImageResponse) GoString() string {
	return s.String()
}

func (s *DeleteImageResponse) SetHeaders(v map[string]*string) *DeleteImageResponse {
	s.Headers = v
	return s
}

func (s *DeleteImageResponse) SetStatusCode(v int32) *DeleteImageResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteImageResponse) SetBody(v *DeleteImageResponseBody) *DeleteImageResponse {
	s.Body = v
	return s
}

type DetailRequest struct {
	// The name of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// imagesearchName
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
}

func (s DetailRequest) String() string {
	return tea.Prettify(s)
}

func (s DetailRequest) GoString() string {
	return s.String()
}

func (s *DetailRequest) SetInstanceName(v string) *DetailRequest {
	s.InstanceName = &v
	return s
}

type DetailResponseBody struct {
	// The details about the instance.
	Instance *DetailResponseBodyInstance `json:"Instance,omitempty" xml:"Instance,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 36C43E96-8F68-44AA-B1AF-B1F7AB94A6C1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DetailResponseBody) GoString() string {
	return s.String()
}

func (s *DetailResponseBody) SetInstance(v *DetailResponseBodyInstance) *DetailResponseBody {
	s.Instance = v
	return s
}

func (s *DetailResponseBody) SetRequestId(v string) *DetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *DetailResponseBody) SetSuccess(v bool) *DetailResponseBody {
	s.Success = &v
	return s
}

type DetailResponseBodyInstance struct {
	// The capacity of the plan. Unit: × 10,000 images.
	//
	// example:
	//
	// 10
	Capacity *int32 `json:"Capacity,omitempty" xml:"Capacity,omitempty"`
	// The name of the instance.
	//
	// example:
	//
	// imagesearchName
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The number of queries per second supported by the plan.
	//
	// example:
	//
	// 1
	Qps *int32 `json:"Qps,omitempty" xml:"Qps,omitempty"`
	// The information about the region.
	//
	// example:
	//
	// cn-shanghai
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The Image Search model.
	//
	// 0: commodity search. 1: generic image search.
	//
	// example:
	//
	// 0
	ServiceType *int32 `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
	// The number of images.
	//
	// example:
	//
	// 10063
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The time when the instance was created. Unit: milliseconds.
	//
	// example:
	//
	// 1620382716000
	UtcCreate *string `json:"UtcCreate,omitempty" xml:"UtcCreate,omitempty"`
	// The time when the instance expires. Unit: milliseconds.
	//
	// example:
	//
	// 1623081600000
	UtcExpireTime *string `json:"UtcExpireTime,omitempty" xml:"UtcExpireTime,omitempty"`
}

func (s DetailResponseBodyInstance) String() string {
	return tea.Prettify(s)
}

func (s DetailResponseBodyInstance) GoString() string {
	return s.String()
}

func (s *DetailResponseBodyInstance) SetCapacity(v int32) *DetailResponseBodyInstance {
	s.Capacity = &v
	return s
}

func (s *DetailResponseBodyInstance) SetName(v string) *DetailResponseBodyInstance {
	s.Name = &v
	return s
}

func (s *DetailResponseBodyInstance) SetQps(v int32) *DetailResponseBodyInstance {
	s.Qps = &v
	return s
}

func (s *DetailResponseBodyInstance) SetRegion(v string) *DetailResponseBodyInstance {
	s.Region = &v
	return s
}

func (s *DetailResponseBodyInstance) SetServiceType(v int32) *DetailResponseBodyInstance {
	s.ServiceType = &v
	return s
}

func (s *DetailResponseBodyInstance) SetTotalCount(v int64) *DetailResponseBodyInstance {
	s.TotalCount = &v
	return s
}

func (s *DetailResponseBodyInstance) SetUtcCreate(v string) *DetailResponseBodyInstance {
	s.UtcCreate = &v
	return s
}

func (s *DetailResponseBodyInstance) SetUtcExpireTime(v string) *DetailResponseBodyInstance {
	s.UtcExpireTime = &v
	return s
}

type DetailResponse struct {
	Headers    map[string]*string  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DetailResponse) String() string {
	return tea.Prettify(s)
}

func (s DetailResponse) GoString() string {
	return s.String()
}

func (s *DetailResponse) SetHeaders(v map[string]*string) *DetailResponse {
	s.Headers = v
	return s
}

func (s *DetailResponse) SetStatusCode(v int32) *DetailResponse {
	s.StatusCode = &v
	return s
}

func (s *DetailResponse) SetBody(v *DetailResponseBody) *DetailResponse {
	s.Body = v
	return s
}

type DumpMetaRequest struct {
	// The name of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// imagesearchName
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
}

func (s DumpMetaRequest) String() string {
	return tea.Prettify(s)
}

func (s DumpMetaRequest) GoString() string {
	return s.String()
}

func (s *DumpMetaRequest) SetInstanceName(v string) *DumpMetaRequest {
	s.InstanceName = &v
	return s
}

type DumpMetaResponseBody struct {
	// The information about the export task.
	Data *DumpMetaResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 36C43E96-8F68-44AA-B1AF-B1F7AB94A6C1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DumpMetaResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DumpMetaResponseBody) GoString() string {
	return s.String()
}

func (s *DumpMetaResponseBody) SetData(v *DumpMetaResponseBodyData) *DumpMetaResponseBody {
	s.Data = v
	return s
}

func (s *DumpMetaResponseBody) SetRequestId(v string) *DumpMetaResponseBody {
	s.RequestId = &v
	return s
}

func (s *DumpMetaResponseBody) SetSuccess(v bool) *DumpMetaResponseBody {
	s.Success = &v
	return s
}

type DumpMetaResponseBodyData struct {
	// The status of the export task.
	//
	// 	- PROCESSING: in progress
	//
	// 	- FAIL: failed
	//
	// 	- SUCCESS: successful
	//
	// example:
	//
	// PROCESSING
	DumpMetaStatus *string `json:"DumpMetaStatus,omitempty" xml:"DumpMetaStatus,omitempty"`
	// The ID of the export task.
	//
	// example:
	//
	// 500
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DumpMetaResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DumpMetaResponseBodyData) GoString() string {
	return s.String()
}

func (s *DumpMetaResponseBodyData) SetDumpMetaStatus(v string) *DumpMetaResponseBodyData {
	s.DumpMetaStatus = &v
	return s
}

func (s *DumpMetaResponseBodyData) SetId(v string) *DumpMetaResponseBodyData {
	s.Id = &v
	return s
}

type DumpMetaResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DumpMetaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DumpMetaResponse) String() string {
	return tea.Prettify(s)
}

func (s DumpMetaResponse) GoString() string {
	return s.String()
}

func (s *DumpMetaResponse) SetHeaders(v map[string]*string) *DumpMetaResponse {
	s.Headers = v
	return s
}

func (s *DumpMetaResponse) SetStatusCode(v int32) *DumpMetaResponse {
	s.StatusCode = &v
	return s
}

func (s *DumpMetaResponse) SetBody(v *DumpMetaResponseBody) *DumpMetaResponse {
	s.Body = v
	return s
}

type DumpMetaListRequest struct {
	// The ID of the export task.
	//
	// example:
	//
	// 500
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the Image Search instance. The name can be up to 20 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// demoinstance1
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The number of the page to return. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of images to return on each page. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DumpMetaListRequest) String() string {
	return tea.Prettify(s)
}

func (s DumpMetaListRequest) GoString() string {
	return s.String()
}

func (s *DumpMetaListRequest) SetId(v int64) *DumpMetaListRequest {
	s.Id = &v
	return s
}

func (s *DumpMetaListRequest) SetInstanceName(v string) *DumpMetaListRequest {
	s.InstanceName = &v
	return s
}

func (s *DumpMetaListRequest) SetPageNumber(v int32) *DumpMetaListRequest {
	s.PageNumber = &v
	return s
}

func (s *DumpMetaListRequest) SetPageSize(v int32) *DumpMetaListRequest {
	s.PageSize = &v
	return s
}

type DumpMetaListResponseBody struct {
	// The information about the task that is used to export metadata.
	Data *DumpMetaListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// B3137727-7D6E-488C-BA21-0E034C38A879
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DumpMetaListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DumpMetaListResponseBody) GoString() string {
	return s.String()
}

func (s *DumpMetaListResponseBody) SetData(v *DumpMetaListResponseBodyData) *DumpMetaListResponseBody {
	s.Data = v
	return s
}

func (s *DumpMetaListResponseBody) SetRequestId(v string) *DumpMetaListResponseBody {
	s.RequestId = &v
	return s
}

type DumpMetaListResponseBodyData struct {
	// A list of tasks that are used to export metadata.
	DumpMetaList []*DumpMetaListResponseBodyDataDumpMetaList `json:"DumpMetaList,omitempty" xml:"DumpMetaList,omitempty" type:"Repeated"`
	// The number of the page to return.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of tasks.
	//
	// example:
	//
	// 15
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DumpMetaListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DumpMetaListResponseBodyData) GoString() string {
	return s.String()
}

func (s *DumpMetaListResponseBodyData) SetDumpMetaList(v []*DumpMetaListResponseBodyDataDumpMetaList) *DumpMetaListResponseBodyData {
	s.DumpMetaList = v
	return s
}

func (s *DumpMetaListResponseBodyData) SetPageNumber(v int32) *DumpMetaListResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *DumpMetaListResponseBodyData) SetPageSize(v int32) *DumpMetaListResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *DumpMetaListResponseBodyData) SetTotalCount(v int64) *DumpMetaListResponseBodyData {
	s.TotalCount = &v
	return s
}

type DumpMetaListResponseBodyDataDumpMetaList struct {
	// The error code returned.
	//
	// 	- A value of 0 indicates that the operation is successful.
	//
	// 	- Values other than 0 indicate errors.
	//
	// example:
	//
	// 0
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The ID of the task.
	//
	// example:
	//
	// 500
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The address where you can download the metadata. The address is valid for 2 hours.
	//
	// example:
	//
	// https://imagesearchname.oss-cn-shanghai.aliyuncs.com/xxx
	MetaUrl *string `json:"MetaUrl,omitempty" xml:"MetaUrl,omitempty"`
	// The error message returned.
	//
	// example:
	//
	// success
	Msg *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	// The status of the export task.
	//
	// 	- PROCESSING: in progress
	//
	// 	- FAIL: failed
	//
	// 	- SUCCESS: successful
	//
	// example:
	//
	// SUCCESS
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The time when the task was created. Unit: milliseconds.
	//
	// example:
	//
	// 1629095713000
	UtcCreate *string `json:"UtcCreate,omitempty" xml:"UtcCreate,omitempty"`
	// The time when the task was updated. Unit: milliseconds.
	//
	// example:
	//
	// 1629095760000
	UtcModified *int64 `json:"UtcModified,omitempty" xml:"UtcModified,omitempty"`
}

func (s DumpMetaListResponseBodyDataDumpMetaList) String() string {
	return tea.Prettify(s)
}

func (s DumpMetaListResponseBodyDataDumpMetaList) GoString() string {
	return s.String()
}

func (s *DumpMetaListResponseBodyDataDumpMetaList) SetCode(v string) *DumpMetaListResponseBodyDataDumpMetaList {
	s.Code = &v
	return s
}

func (s *DumpMetaListResponseBodyDataDumpMetaList) SetId(v int64) *DumpMetaListResponseBodyDataDumpMetaList {
	s.Id = &v
	return s
}

func (s *DumpMetaListResponseBodyDataDumpMetaList) SetMetaUrl(v string) *DumpMetaListResponseBodyDataDumpMetaList {
	s.MetaUrl = &v
	return s
}

func (s *DumpMetaListResponseBodyDataDumpMetaList) SetMsg(v string) *DumpMetaListResponseBodyDataDumpMetaList {
	s.Msg = &v
	return s
}

func (s *DumpMetaListResponseBodyDataDumpMetaList) SetStatus(v string) *DumpMetaListResponseBodyDataDumpMetaList {
	s.Status = &v
	return s
}

func (s *DumpMetaListResponseBodyDataDumpMetaList) SetUtcCreate(v string) *DumpMetaListResponseBodyDataDumpMetaList {
	s.UtcCreate = &v
	return s
}

func (s *DumpMetaListResponseBodyDataDumpMetaList) SetUtcModified(v int64) *DumpMetaListResponseBodyDataDumpMetaList {
	s.UtcModified = &v
	return s
}

type DumpMetaListResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DumpMetaListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DumpMetaListResponse) String() string {
	return tea.Prettify(s)
}

func (s DumpMetaListResponse) GoString() string {
	return s.String()
}

func (s *DumpMetaListResponse) SetHeaders(v map[string]*string) *DumpMetaListResponse {
	s.Headers = v
	return s
}

func (s *DumpMetaListResponse) SetStatusCode(v int32) *DumpMetaListResponse {
	s.StatusCode = &v
	return s
}

func (s *DumpMetaListResponse) SetBody(v *DumpMetaListResponseBody) *DumpMetaListResponse {
	s.Body = v
	return s
}

type IncreaseInstanceRequest struct {
	// The name of the Object Storage Service (OSS) bucket.
	//
	// >  The bucket must be in the same region as the Image Search instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// bucketName
	BucketName *string `json:"BucketName,omitempty" xml:"BucketName,omitempty"`
	// The callback address.
	//
	// example:
	//
	// http://xxxxx
	CallbackAddress *string `json:"CallbackAddress,omitempty" xml:"CallbackAddress,omitempty"`
	// The name of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// imagesearchName
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The absolute path to the increment.meta file in the bucket. The path must start with a forward slash (/) and cannot end with a forward slash (/).
	//
	// This parameter is required.
	//
	// example:
	//
	// /xxx/xxx
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
}

func (s IncreaseInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s IncreaseInstanceRequest) GoString() string {
	return s.String()
}

func (s *IncreaseInstanceRequest) SetBucketName(v string) *IncreaseInstanceRequest {
	s.BucketName = &v
	return s
}

func (s *IncreaseInstanceRequest) SetCallbackAddress(v string) *IncreaseInstanceRequest {
	s.CallbackAddress = &v
	return s
}

func (s *IncreaseInstanceRequest) SetInstanceName(v string) *IncreaseInstanceRequest {
	s.InstanceName = &v
	return s
}

func (s *IncreaseInstanceRequest) SetPath(v string) *IncreaseInstanceRequest {
	s.Path = &v
	return s
}

type IncreaseInstanceResponseBody struct {
	// The information about the task.
	Data *IncreaseInstanceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 36C43E96-8F68-44AA-B1AF-B1F7AB94A6C1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s IncreaseInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s IncreaseInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *IncreaseInstanceResponseBody) SetData(v *IncreaseInstanceResponseBodyData) *IncreaseInstanceResponseBody {
	s.Data = v
	return s
}

func (s *IncreaseInstanceResponseBody) SetRequestId(v string) *IncreaseInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *IncreaseInstanceResponseBody) SetSuccess(v bool) *IncreaseInstanceResponseBody {
	s.Success = &v
	return s
}

type IncreaseInstanceResponseBodyData struct {
	// The ID of the task.
	//
	// example:
	//
	// 500
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The status of the task.
	//
	// 	- PROCESSING: in progress
	//
	// 	- FAIL: failed
	//
	// 	- SUCCESS: successful
	//
	// example:
	//
	// PROCESSING
	IncrementStatus *string `json:"IncrementStatus,omitempty" xml:"IncrementStatus,omitempty"`
}

func (s IncreaseInstanceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s IncreaseInstanceResponseBodyData) GoString() string {
	return s.String()
}

func (s *IncreaseInstanceResponseBodyData) SetId(v string) *IncreaseInstanceResponseBodyData {
	s.Id = &v
	return s
}

func (s *IncreaseInstanceResponseBodyData) SetIncrementStatus(v string) *IncreaseInstanceResponseBodyData {
	s.IncrementStatus = &v
	return s
}

type IncreaseInstanceResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *IncreaseInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s IncreaseInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s IncreaseInstanceResponse) GoString() string {
	return s.String()
}

func (s *IncreaseInstanceResponse) SetHeaders(v map[string]*string) *IncreaseInstanceResponse {
	s.Headers = v
	return s
}

func (s *IncreaseInstanceResponse) SetStatusCode(v int32) *IncreaseInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *IncreaseInstanceResponse) SetBody(v *IncreaseInstanceResponseBody) *IncreaseInstanceResponse {
	s.Body = v
	return s
}

type IncreaseListRequest struct {
	// The name of the Object Storage Service (OSS) bucket.
	//
	// example:
	//
	// bucketName
	BucketName *string `json:"BucketName,omitempty" xml:"BucketName,omitempty"`
	// The ID of the batch task.
	//
	// example:
	//
	// 500
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the Image Search instance. The name can be up to 20 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// demoinstance1
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The number of the page to return. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of images to return on each page. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The absolute path to the increment.meta file in the bucket. The path must start with a forward slash (/) and cannot end with a forward slash (/).
	//
	// example:
	//
	// /xxx/xxx
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
}

func (s IncreaseListRequest) String() string {
	return tea.Prettify(s)
}

func (s IncreaseListRequest) GoString() string {
	return s.String()
}

func (s *IncreaseListRequest) SetBucketName(v string) *IncreaseListRequest {
	s.BucketName = &v
	return s
}

func (s *IncreaseListRequest) SetId(v int64) *IncreaseListRequest {
	s.Id = &v
	return s
}

func (s *IncreaseListRequest) SetInstanceName(v string) *IncreaseListRequest {
	s.InstanceName = &v
	return s
}

func (s *IncreaseListRequest) SetPageNumber(v int32) *IncreaseListRequest {
	s.PageNumber = &v
	return s
}

func (s *IncreaseListRequest) SetPageSize(v int32) *IncreaseListRequest {
	s.PageSize = &v
	return s
}

func (s *IncreaseListRequest) SetPath(v string) *IncreaseListRequest {
	s.Path = &v
	return s
}

type IncreaseListResponseBody struct {
	// The information about the batch task.
	Data *IncreaseListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// B3137727-7D6E-488C-BA21-0E034C38A879
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s IncreaseListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s IncreaseListResponseBody) GoString() string {
	return s.String()
}

func (s *IncreaseListResponseBody) SetData(v *IncreaseListResponseBodyData) *IncreaseListResponseBody {
	s.Data = v
	return s
}

func (s *IncreaseListResponseBody) SetRequestId(v string) *IncreaseListResponseBody {
	s.RequestId = &v
	return s
}

type IncreaseListResponseBodyData struct {
	// A list of batch tasks.
	Increments *IncreaseListResponseBodyDataIncrements `json:"Increments,omitempty" xml:"Increments,omitempty" type:"Struct"`
	// The number of the page to return.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of tasks.
	//
	// example:
	//
	// 15
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s IncreaseListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s IncreaseListResponseBodyData) GoString() string {
	return s.String()
}

func (s *IncreaseListResponseBodyData) SetIncrements(v *IncreaseListResponseBodyDataIncrements) *IncreaseListResponseBodyData {
	s.Increments = v
	return s
}

func (s *IncreaseListResponseBodyData) SetPageNumber(v int32) *IncreaseListResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *IncreaseListResponseBodyData) SetPageSize(v int32) *IncreaseListResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *IncreaseListResponseBodyData) SetTotalCount(v int64) *IncreaseListResponseBodyData {
	s.TotalCount = &v
	return s
}

type IncreaseListResponseBodyDataIncrements struct {
	Instance []*IncreaseListResponseBodyDataIncrementsInstance `json:"Instance,omitempty" xml:"Instance,omitempty" type:"Repeated"`
}

func (s IncreaseListResponseBodyDataIncrements) String() string {
	return tea.Prettify(s)
}

func (s IncreaseListResponseBodyDataIncrements) GoString() string {
	return s.String()
}

func (s *IncreaseListResponseBodyDataIncrements) SetInstance(v []*IncreaseListResponseBodyDataIncrementsInstance) *IncreaseListResponseBodyDataIncrements {
	s.Instance = v
	return s
}

type IncreaseListResponseBodyDataIncrementsInstance struct {
	// The name of the Object Storage Service (OSS) bucket.
	//
	// example:
	//
	// bucketName
	BucketName *string `json:"BucketName,omitempty" xml:"BucketName,omitempty"`
	// The callback address.
	//
	// example:
	//
	// http://xxxxx
	CallbackAddress *string `json:"CallbackAddress,omitempty" xml:"CallbackAddress,omitempty"`
	// The error code returned.
	//
	// 	- A value of 0 indicates that the operation is successful.
	//
	// 	- Values other than 0 indicate errors.
	//
	// example:
	//
	// 0
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The address where you can download the result. The address is valid for 2 hours.
	//
	// example:
	//
	// https://imagesearchname.oss-cn-shanghai.aliyuncs.com/xxx
	ErrorUrl *string `json:"ErrorUrl,omitempty" xml:"ErrorUrl,omitempty"`
	// The ID of the task.
	//
	// example:
	//
	// 500
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The error message returned.
	//
	// example:
	//
	// sucess
	Msg *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	// The absolute path to the increment.meta file in the bucket. The path must start with a forward slash (/) and cannot end with a forward slash (/).
	//
	// example:
	//
	// /xx/xx
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// The status of the batch task.
	//
	// 	- PROCESSING: in progress
	//
	// 	- FAIL: failed
	//
	// 	- SUCCESS: successful
	//
	// example:
	//
	// SUCCESS
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The time when the task was created. Unit: milliseconds.
	//
	// example:
	//
	// 1629095713000
	UtcCreate *string `json:"UtcCreate,omitempty" xml:"UtcCreate,omitempty"`
	// The time when the task was updated. Unit: milliseconds.
	//
	// example:
	//
	// 1629095760000
	UtcModified *int64 `json:"UtcModified,omitempty" xml:"UtcModified,omitempty"`
}

func (s IncreaseListResponseBodyDataIncrementsInstance) String() string {
	return tea.Prettify(s)
}

func (s IncreaseListResponseBodyDataIncrementsInstance) GoString() string {
	return s.String()
}

func (s *IncreaseListResponseBodyDataIncrementsInstance) SetBucketName(v string) *IncreaseListResponseBodyDataIncrementsInstance {
	s.BucketName = &v
	return s
}

func (s *IncreaseListResponseBodyDataIncrementsInstance) SetCallbackAddress(v string) *IncreaseListResponseBodyDataIncrementsInstance {
	s.CallbackAddress = &v
	return s
}

func (s *IncreaseListResponseBodyDataIncrementsInstance) SetCode(v string) *IncreaseListResponseBodyDataIncrementsInstance {
	s.Code = &v
	return s
}

func (s *IncreaseListResponseBodyDataIncrementsInstance) SetErrorUrl(v string) *IncreaseListResponseBodyDataIncrementsInstance {
	s.ErrorUrl = &v
	return s
}

func (s *IncreaseListResponseBodyDataIncrementsInstance) SetId(v int64) *IncreaseListResponseBodyDataIncrementsInstance {
	s.Id = &v
	return s
}

func (s *IncreaseListResponseBodyDataIncrementsInstance) SetMsg(v string) *IncreaseListResponseBodyDataIncrementsInstance {
	s.Msg = &v
	return s
}

func (s *IncreaseListResponseBodyDataIncrementsInstance) SetPath(v string) *IncreaseListResponseBodyDataIncrementsInstance {
	s.Path = &v
	return s
}

func (s *IncreaseListResponseBodyDataIncrementsInstance) SetStatus(v string) *IncreaseListResponseBodyDataIncrementsInstance {
	s.Status = &v
	return s
}

func (s *IncreaseListResponseBodyDataIncrementsInstance) SetUtcCreate(v string) *IncreaseListResponseBodyDataIncrementsInstance {
	s.UtcCreate = &v
	return s
}

func (s *IncreaseListResponseBodyDataIncrementsInstance) SetUtcModified(v int64) *IncreaseListResponseBodyDataIncrementsInstance {
	s.UtcModified = &v
	return s
}

type IncreaseListResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *IncreaseListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s IncreaseListResponse) String() string {
	return tea.Prettify(s)
}

func (s IncreaseListResponse) GoString() string {
	return s.String()
}

func (s *IncreaseListResponse) SetHeaders(v map[string]*string) *IncreaseListResponse {
	s.Headers = v
	return s
}

func (s *IncreaseListResponse) SetStatusCode(v int32) *IncreaseListResponse {
	s.StatusCode = &v
	return s
}

func (s *IncreaseListResponse) SetBody(v *IncreaseListResponseBody) *IncreaseListResponse {
	s.Body = v
	return s
}

type SearchImageByNameRequest struct {
	// The category of the product. For more information, see [Category references](https://help.aliyun.com/document_detail/179184.html).
	//
	// 	- For product search: If a category is specified, the specified category prevails. If no category is specified, the system estimates and selects a category. The category selected by the system is included in the response.
	//
	// 	- For generic search: The parameter value is set to 88888888 regardless of whether a category is specified.
	//
	// example:
	//
	// 88888888
	CategoryId        *int32 `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	DistinctProductId *bool  `json:"DistinctProductId,omitempty" xml:"DistinctProductId,omitempty"`
	// The filter conditions. int_attr supports the following operators: >, >=, <, <=, and =. str_attr supports the following operators: = and !=. You can set the logical operator between conditions to AND or OR.
	//
	// Examples:
	//
	// 	- int_attr>=100
	//
	// 	- str_attr!="value1"
	//
	// 	- int_attr=1000 AND str_attr="value1"
	//
	// example:
	//
	// int_attr>=100
	Filter *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	// The name of the Image Search instance. The name can be up to 20 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// demoinstance1
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The number of images to return on each page. Valid values: 1 to 100. Default value: 10.
	//
	// example:
	//
	// 10
	Num *int32 `json:"Num,omitempty" xml:"Num,omitempty"`
	// The name of the image.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2092061_1.jpg
	PicName *string `json:"PicName,omitempty" xml:"PicName,omitempty"`
	// The ID of the product.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2092061_1
	ProductId *string `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	// The number of the image to return. Valid values: 0 to 499. Default value: 0.
	//
	// example:
	//
	// 0
	Start *int32 `json:"Start,omitempty" xml:"Start,omitempty"`
}

func (s SearchImageByNameRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchImageByNameRequest) GoString() string {
	return s.String()
}

func (s *SearchImageByNameRequest) SetCategoryId(v int32) *SearchImageByNameRequest {
	s.CategoryId = &v
	return s
}

func (s *SearchImageByNameRequest) SetDistinctProductId(v bool) *SearchImageByNameRequest {
	s.DistinctProductId = &v
	return s
}

func (s *SearchImageByNameRequest) SetFilter(v string) *SearchImageByNameRequest {
	s.Filter = &v
	return s
}

func (s *SearchImageByNameRequest) SetInstanceName(v string) *SearchImageByNameRequest {
	s.InstanceName = &v
	return s
}

func (s *SearchImageByNameRequest) SetNum(v int32) *SearchImageByNameRequest {
	s.Num = &v
	return s
}

func (s *SearchImageByNameRequest) SetPicName(v string) *SearchImageByNameRequest {
	s.PicName = &v
	return s
}

func (s *SearchImageByNameRequest) SetProductId(v string) *SearchImageByNameRequest {
	s.ProductId = &v
	return s
}

func (s *SearchImageByNameRequest) SetStart(v int32) *SearchImageByNameRequest {
	s.Start = &v
	return s
}

type SearchImageByNameResponseBody struct {
	// The product descriptions returned.
	Auctions []*SearchImageByNameResponseBodyAuctions `json:"Auctions,omitempty" xml:"Auctions,omitempty" type:"Repeated"`
	// The error code returned.
	//
	// 	- A value of 0 indicates that the operation is successful.
	//
	// 	- Values other than 0 indicate errors.
	//
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The summary of the search result.
	Head *SearchImageByNameResponseBodyHead `json:"Head,omitempty" xml:"Head,omitempty" type:"Struct"`
	// The error message returned.
	//
	// example:
	//
	// success
	Msg *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	// The information such as the system-selected category and result of subject recognition.
	PicInfo *SearchImageByNameResponseBodyPicInfo `json:"PicInfo,omitempty" xml:"PicInfo,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 36C43E96-8F68-44AA-B1AF-B1F7AB94A6C1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SearchImageByNameResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SearchImageByNameResponseBody) GoString() string {
	return s.String()
}

func (s *SearchImageByNameResponseBody) SetAuctions(v []*SearchImageByNameResponseBodyAuctions) *SearchImageByNameResponseBody {
	s.Auctions = v
	return s
}

func (s *SearchImageByNameResponseBody) SetCode(v int32) *SearchImageByNameResponseBody {
	s.Code = &v
	return s
}

func (s *SearchImageByNameResponseBody) SetHead(v *SearchImageByNameResponseBodyHead) *SearchImageByNameResponseBody {
	s.Head = v
	return s
}

func (s *SearchImageByNameResponseBody) SetMsg(v string) *SearchImageByNameResponseBody {
	s.Msg = &v
	return s
}

func (s *SearchImageByNameResponseBody) SetPicInfo(v *SearchImageByNameResponseBodyPicInfo) *SearchImageByNameResponseBody {
	s.PicInfo = v
	return s
}

func (s *SearchImageByNameResponseBody) SetRequestId(v string) *SearchImageByNameResponseBody {
	s.RequestId = &v
	return s
}

func (s *SearchImageByNameResponseBody) SetSuccess(v bool) *SearchImageByNameResponseBody {
	s.Success = &v
	return s
}

type SearchImageByNameResponseBodyAuctions struct {
	// The category of the image.
	//
	// example:
	//
	// 20
	CategoryId *int32 `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	// The user-defined content.
	//
	// example:
	//
	// zidingyi
	CustomContent *string `json:"CustomContent,omitempty" xml:"CustomContent,omitempty"`
	// The attribute, which is an integer.
	//
	// example:
	//
	// 2
	IntAttr *int32 `json:"IntAttr,omitempty" xml:"IntAttr,omitempty"`
	// example:
	//
	// 20
	IntAttr2 *int32 `json:"IntAttr2,omitempty" xml:"IntAttr2,omitempty"`
	IntAttr3 *int32 `json:"IntAttr3,omitempty" xml:"IntAttr3,omitempty"`
	IntAttr4 *int32 `json:"IntAttr4,omitempty" xml:"IntAttr4,omitempty"`
	// The name of the image.
	//
	// example:
	//
	// 2092061_1.jpg
	PicName *string `json:"PicName,omitempty" xml:"PicName,omitempty"`
	// The ID of the product.
	//
	// example:
	//
	// 2092061_1
	ProductId *string `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	// The similarity score of the returned image. Valid values: 0 to 1.
	//
	// >  To use this feature, you must upgrade the SDK to version 3.1.1.
	//
	// example:
	//
	// 1
	Score *float32 `json:"Score,omitempty" xml:"Score,omitempty"`
	// The score information about the image.
	//
	// > 	- This parameter is not supported. We recommend that you use the Score parameter.
	//
	// >	- The SortExprValues parameter indicates a 2-tuple in which values are separated by a semicolon (;). The first value indicates the correlation score of the returned image. A greater value indicates a higher correlation with the sample image. Different algorithms are used.
	//
	// >	- If the value of CategoryId is within the value range from 0 to 2, the value range of SortExprValues is from 0 to 7.33136443711219e+24.
	//
	// >	- If the value of CategoryId is not within the value range from 0 to 2, the value range of SortExprValues is from 0 to 5.37633353624177e+24. If the returned image is identical with the sample image, the highest correlation score is generated.
	//
	// example:
	//
	// 5.37633353624177e+24;0
	SortExprValues *string `json:"SortExprValues,omitempty" xml:"SortExprValues,omitempty"`
	// The attribute, which is a string.
	//
	// example:
	//
	// ss
	StrAttr *string `json:"StrAttr,omitempty" xml:"StrAttr,omitempty"`
	// example:
	//
	// test
	StrAttr2 *string `json:"StrAttr2,omitempty" xml:"StrAttr2,omitempty"`
	StrAttr3 *string `json:"StrAttr3,omitempty" xml:"StrAttr3,omitempty"`
	StrAttr4 *string `json:"StrAttr4,omitempty" xml:"StrAttr4,omitempty"`
}

func (s SearchImageByNameResponseBodyAuctions) String() string {
	return tea.Prettify(s)
}

func (s SearchImageByNameResponseBodyAuctions) GoString() string {
	return s.String()
}

func (s *SearchImageByNameResponseBodyAuctions) SetCategoryId(v int32) *SearchImageByNameResponseBodyAuctions {
	s.CategoryId = &v
	return s
}

func (s *SearchImageByNameResponseBodyAuctions) SetCustomContent(v string) *SearchImageByNameResponseBodyAuctions {
	s.CustomContent = &v
	return s
}

func (s *SearchImageByNameResponseBodyAuctions) SetIntAttr(v int32) *SearchImageByNameResponseBodyAuctions {
	s.IntAttr = &v
	return s
}

func (s *SearchImageByNameResponseBodyAuctions) SetIntAttr2(v int32) *SearchImageByNameResponseBodyAuctions {
	s.IntAttr2 = &v
	return s
}

func (s *SearchImageByNameResponseBodyAuctions) SetIntAttr3(v int32) *SearchImageByNameResponseBodyAuctions {
	s.IntAttr3 = &v
	return s
}

func (s *SearchImageByNameResponseBodyAuctions) SetIntAttr4(v int32) *SearchImageByNameResponseBodyAuctions {
	s.IntAttr4 = &v
	return s
}

func (s *SearchImageByNameResponseBodyAuctions) SetPicName(v string) *SearchImageByNameResponseBodyAuctions {
	s.PicName = &v
	return s
}

func (s *SearchImageByNameResponseBodyAuctions) SetProductId(v string) *SearchImageByNameResponseBodyAuctions {
	s.ProductId = &v
	return s
}

func (s *SearchImageByNameResponseBodyAuctions) SetScore(v float32) *SearchImageByNameResponseBodyAuctions {
	s.Score = &v
	return s
}

func (s *SearchImageByNameResponseBodyAuctions) SetSortExprValues(v string) *SearchImageByNameResponseBodyAuctions {
	s.SortExprValues = &v
	return s
}

func (s *SearchImageByNameResponseBodyAuctions) SetStrAttr(v string) *SearchImageByNameResponseBodyAuctions {
	s.StrAttr = &v
	return s
}

func (s *SearchImageByNameResponseBodyAuctions) SetStrAttr2(v string) *SearchImageByNameResponseBodyAuctions {
	s.StrAttr2 = &v
	return s
}

func (s *SearchImageByNameResponseBodyAuctions) SetStrAttr3(v string) *SearchImageByNameResponseBodyAuctions {
	s.StrAttr3 = &v
	return s
}

func (s *SearchImageByNameResponseBodyAuctions) SetStrAttr4(v string) *SearchImageByNameResponseBodyAuctions {
	s.StrAttr4 = &v
	return s
}

type SearchImageByNameResponseBodyHead struct {
	// The number of images returned.
	//
	// example:
	//
	// 10
	DocsFound *int32 `json:"DocsFound,omitempty" xml:"DocsFound,omitempty"`
	// The number of images that match the search conditions on the Image Search instance.
	//
	// example:
	//
	// 10000
	DocsReturn *int32 `json:"DocsReturn,omitempty" xml:"DocsReturn,omitempty"`
	// The time it takes to complete the search process. Unit: milliseconds.
	//
	// example:
	//
	// 95
	SearchTime *int32 `json:"SearchTime,omitempty" xml:"SearchTime,omitempty"`
}

func (s SearchImageByNameResponseBodyHead) String() string {
	return tea.Prettify(s)
}

func (s SearchImageByNameResponseBodyHead) GoString() string {
	return s.String()
}

func (s *SearchImageByNameResponseBodyHead) SetDocsFound(v int32) *SearchImageByNameResponseBodyHead {
	s.DocsFound = &v
	return s
}

func (s *SearchImageByNameResponseBodyHead) SetDocsReturn(v int32) *SearchImageByNameResponseBodyHead {
	s.DocsReturn = &v
	return s
}

func (s *SearchImageByNameResponseBodyHead) SetSearchTime(v int32) *SearchImageByNameResponseBodyHead {
	s.SearchTime = &v
	return s
}

type SearchImageByNameResponseBodyPicInfo struct {
	// The categories that are supported by the system.
	AllCategories []*SearchImageByNameResponseBodyPicInfoAllCategories `json:"AllCategories,omitempty" xml:"AllCategories,omitempty" type:"Repeated"`
	// The category selected by the system.
	//
	// If a category is specified in the request, the specified category prevails.
	//
	// example:
	//
	// 20
	CategoryId *int32 `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	// The recognized subjects.
	MultiRegion []*SearchImageByNameResponseBodyPicInfoMultiRegion `json:"MultiRegion,omitempty" xml:"MultiRegion,omitempty" type:"Repeated"`
	// The result of subject recognition.
	//
	// The subject area of the image, in the format of x1,x2,y1,y2. Specifically, x1 and y1 specify the upper-left pixel, and x2 and y2 specify the lower-right pixel. If a subject area is specified in the request, the specified subject area prevails.
	//
	// example:
	//
	// 280,486,232,351
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s SearchImageByNameResponseBodyPicInfo) String() string {
	return tea.Prettify(s)
}

func (s SearchImageByNameResponseBodyPicInfo) GoString() string {
	return s.String()
}

func (s *SearchImageByNameResponseBodyPicInfo) SetAllCategories(v []*SearchImageByNameResponseBodyPicInfoAllCategories) *SearchImageByNameResponseBodyPicInfo {
	s.AllCategories = v
	return s
}

func (s *SearchImageByNameResponseBodyPicInfo) SetCategoryId(v int32) *SearchImageByNameResponseBodyPicInfo {
	s.CategoryId = &v
	return s
}

func (s *SearchImageByNameResponseBodyPicInfo) SetMultiRegion(v []*SearchImageByNameResponseBodyPicInfoMultiRegion) *SearchImageByNameResponseBodyPicInfo {
	s.MultiRegion = v
	return s
}

func (s *SearchImageByNameResponseBodyPicInfo) SetRegion(v string) *SearchImageByNameResponseBodyPicInfo {
	s.Region = &v
	return s
}

type SearchImageByNameResponseBodyPicInfoAllCategories struct {
	// The ID of the category.
	//
	// example:
	//
	// other
	Id *int32 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the category.
	//
	// example:
	//
	// 88888888
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s SearchImageByNameResponseBodyPicInfoAllCategories) String() string {
	return tea.Prettify(s)
}

func (s SearchImageByNameResponseBodyPicInfoAllCategories) GoString() string {
	return s.String()
}

func (s *SearchImageByNameResponseBodyPicInfoAllCategories) SetId(v int32) *SearchImageByNameResponseBodyPicInfoAllCategories {
	s.Id = &v
	return s
}

func (s *SearchImageByNameResponseBodyPicInfoAllCategories) SetName(v string) *SearchImageByNameResponseBodyPicInfoAllCategories {
	s.Name = &v
	return s
}

type SearchImageByNameResponseBodyPicInfoMultiRegion struct {
	// The result of subject recognition.
	//
	// The subject area of the image, in the format of x1,x2,y1,y2. Specifically, x1 and y1 specify the upper-left pixel, and x2 and y2 specify the lower-right pixel. If a subject area is specified in the request, the specified subject area prevails.
	//
	// example:
	//
	// 280,486,232,351
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s SearchImageByNameResponseBodyPicInfoMultiRegion) String() string {
	return tea.Prettify(s)
}

func (s SearchImageByNameResponseBodyPicInfoMultiRegion) GoString() string {
	return s.String()
}

func (s *SearchImageByNameResponseBodyPicInfoMultiRegion) SetRegion(v string) *SearchImageByNameResponseBodyPicInfoMultiRegion {
	s.Region = &v
	return s
}

type SearchImageByNameResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SearchImageByNameResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SearchImageByNameResponse) String() string {
	return tea.Prettify(s)
}

func (s SearchImageByNameResponse) GoString() string {
	return s.String()
}

func (s *SearchImageByNameResponse) SetHeaders(v map[string]*string) *SearchImageByNameResponse {
	s.Headers = v
	return s
}

func (s *SearchImageByNameResponse) SetStatusCode(v int32) *SearchImageByNameResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchImageByNameResponse) SetBody(v *SearchImageByNameResponseBody) *SearchImageByNameResponse {
	s.Body = v
	return s
}

type SearchImageByPicRequest struct {
	// The category of the product. For more information, see [Category references](https://help.aliyun.com/document_detail/179184.html).
	//
	// 	- For product search: If a category is specified, the specified category prevails. If no category is specified, the system estimates and selects a category. The category selected by the system is included in the response.
	//
	// 	- For generic search: The parameter value is set to 88888888 regardless of whether a category is specified.
	//
	// example:
	//
	// 88888888
	CategoryId *int32 `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	// Specifies whether to recognize the subject in the image and search for images based on the recognized subject. Valid values: true and false. Default value: true.
	//
	// 	- true: The system recognizes the subject in the image, and searches for images based on the recognized subject. The recognition result is included in the response.
	//
	// 	- false: The system does not recognize the subject of the image, and searches for images based on the entire image.
	//
	// example:
	//
	// true
	Crop              *bool `json:"Crop,omitempty" xml:"Crop,omitempty"`
	DistinctProductId *bool `json:"DistinctProductId,omitempty" xml:"DistinctProductId,omitempty"`
	// The filter conditions. int_attr supports the following operators: >, >=, <, <=, and =. str_attr supports the following operators: = and !=. You can set the logical operator between conditions to AND or OR.
	//
	// Examples:
	//
	// 	- int_attr>=100
	//
	// 	- str_attr!="value1"
	//
	// 	- Example: int_attr=1000 AND str_attr="value1"
	//
	// example:
	//
	// int_attr=1000 AND str_attr="value1"
	Filter *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	// The name of the Image Search instance. The name can be up to 20 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// demoinstance1
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The number of images to return on each page. Valid values: 1 to 100. Default value: 10.
	//
	// example:
	//
	// 10
	Num *int32 `json:"Num,omitempty" xml:"Num,omitempty"`
	// The image file. The image file is encoded in Base64.
	//
	// 	- The file size of the image cannot exceed 4 MB.
	//
	// 	- The following image formats are supported: PNG, JPG, JPEG, BMP, GIF, WebP, TIFF, and PPM.
	//
	// 	- The transmission timeout period cannot exceed 5 seconds.
	//
	// 	- For brand image searches, the length and the width of the image must range from 200 pixels to 4,096 pixels.
	//
	// 	- For cloth image searches, the length and the width of the image must range from 448 pixels to 4,096 pixels.
	//
	// 	- For product and generic image searches, the length and the width of the image must range from 100 pixels to 4,096 pixels.
	//
	// 	- The image cannot contain rotation settings.
	//
	// This parameter is required.
	//
	// example:
	//
	// AAAANSUhEUgAAAPcAAAEVCAYAAAA8d3NuAAAAAXNSR0IArs......RK5CYII=
	PicContent *string `json:"PicContent,omitempty" xml:"PicContent,omitempty"`
	// The subject area of the image. Specify the subject area in the following format: `x1,x2,y1,y2`. `x1 and y1` represent the upper-left corner pixel. `x2 and y2` represent the lower-right corner pixel.
	//
	// >	- If you set the Region parameter, the system searches for images based on the value of Region regardless of the value of the Crop parameter.
	//
	// example:
	//
	// 280,486,232,351
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The number of the image to return. Valid values: 0 to 499. Default value: 0.
	//
	// example:
	//
	// 0
	Start *int32 `json:"Start,omitempty" xml:"Start,omitempty"`
}

func (s SearchImageByPicRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchImageByPicRequest) GoString() string {
	return s.String()
}

func (s *SearchImageByPicRequest) SetCategoryId(v int32) *SearchImageByPicRequest {
	s.CategoryId = &v
	return s
}

func (s *SearchImageByPicRequest) SetCrop(v bool) *SearchImageByPicRequest {
	s.Crop = &v
	return s
}

func (s *SearchImageByPicRequest) SetDistinctProductId(v bool) *SearchImageByPicRequest {
	s.DistinctProductId = &v
	return s
}

func (s *SearchImageByPicRequest) SetFilter(v string) *SearchImageByPicRequest {
	s.Filter = &v
	return s
}

func (s *SearchImageByPicRequest) SetInstanceName(v string) *SearchImageByPicRequest {
	s.InstanceName = &v
	return s
}

func (s *SearchImageByPicRequest) SetNum(v int32) *SearchImageByPicRequest {
	s.Num = &v
	return s
}

func (s *SearchImageByPicRequest) SetPicContent(v string) *SearchImageByPicRequest {
	s.PicContent = &v
	return s
}

func (s *SearchImageByPicRequest) SetRegion(v string) *SearchImageByPicRequest {
	s.Region = &v
	return s
}

func (s *SearchImageByPicRequest) SetStart(v int32) *SearchImageByPicRequest {
	s.Start = &v
	return s
}

type SearchImageByPicAdvanceRequest struct {
	// The category of the product. For more information, see [Category references](https://help.aliyun.com/document_detail/179184.html).
	//
	// 	- For product search: If a category is specified, the specified category prevails. If no category is specified, the system estimates and selects a category. The category selected by the system is included in the response.
	//
	// 	- For generic search: The parameter value is set to 88888888 regardless of whether a category is specified.
	//
	// example:
	//
	// 88888888
	CategoryId *int32 `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	// Specifies whether to recognize the subject in the image and search for images based on the recognized subject. Valid values: true and false. Default value: true.
	//
	// 	- true: The system recognizes the subject in the image, and searches for images based on the recognized subject. The recognition result is included in the response.
	//
	// 	- false: The system does not recognize the subject of the image, and searches for images based on the entire image.
	//
	// example:
	//
	// true
	Crop              *bool `json:"Crop,omitempty" xml:"Crop,omitempty"`
	DistinctProductId *bool `json:"DistinctProductId,omitempty" xml:"DistinctProductId,omitempty"`
	// The filter conditions. int_attr supports the following operators: >, >=, <, <=, and =. str_attr supports the following operators: = and !=. You can set the logical operator between conditions to AND or OR.
	//
	// Examples:
	//
	// 	- int_attr>=100
	//
	// 	- str_attr!="value1"
	//
	// 	- Example: int_attr=1000 AND str_attr="value1"
	//
	// example:
	//
	// int_attr=1000 AND str_attr="value1"
	Filter *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	// The name of the Image Search instance. The name can be up to 20 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// demoinstance1
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The number of images to return on each page. Valid values: 1 to 100. Default value: 10.
	//
	// example:
	//
	// 10
	Num *int32 `json:"Num,omitempty" xml:"Num,omitempty"`
	// The image file. The image file is encoded in Base64.
	//
	// 	- The file size of the image cannot exceed 4 MB.
	//
	// 	- The following image formats are supported: PNG, JPG, JPEG, BMP, GIF, WebP, TIFF, and PPM.
	//
	// 	- The transmission timeout period cannot exceed 5 seconds.
	//
	// 	- For brand image searches, the length and the width of the image must range from 200 pixels to 4,096 pixels.
	//
	// 	- For cloth image searches, the length and the width of the image must range from 448 pixels to 4,096 pixels.
	//
	// 	- For product and generic image searches, the length and the width of the image must range from 100 pixels to 4,096 pixels.
	//
	// 	- The image cannot contain rotation settings.
	//
	// This parameter is required.
	//
	// example:
	//
	// AAAANSUhEUgAAAPcAAAEVCAYAAAA8d3NuAAAAAXNSR0IArs......RK5CYII=
	PicContentObject io.Reader `json:"PicContent,omitempty" xml:"PicContent,omitempty"`
	// The subject area of the image. Specify the subject area in the following format: `x1,x2,y1,y2`. `x1 and y1` represent the upper-left corner pixel. `x2 and y2` represent the lower-right corner pixel.
	//
	// >	- If you set the Region parameter, the system searches for images based on the value of Region regardless of the value of the Crop parameter.
	//
	// example:
	//
	// 280,486,232,351
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The number of the image to return. Valid values: 0 to 499. Default value: 0.
	//
	// example:
	//
	// 0
	Start *int32 `json:"Start,omitempty" xml:"Start,omitempty"`
}

func (s SearchImageByPicAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchImageByPicAdvanceRequest) GoString() string {
	return s.String()
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

func (s *SearchImageByPicAdvanceRequest) SetStart(v int32) *SearchImageByPicAdvanceRequest {
	s.Start = &v
	return s
}

type SearchImageByPicResponseBody struct {
	// The product descriptions returned.
	Auctions []*SearchImageByPicResponseBodyAuctions `json:"Auctions,omitempty" xml:"Auctions,omitempty" type:"Repeated"`
	// The error code returned.
	//
	// 	- A value of 0 indicates that the operation is successful.
	//
	// 	- Values other than 0 indicate errors.
	//
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The summary of the search result.
	Head *SearchImageByPicResponseBodyHead `json:"Head,omitempty" xml:"Head,omitempty" type:"Struct"`
	// The error message returned.
	//
	// example:
	//
	// success
	Msg *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	// The information such as the system-selected category and result of subject recognition.
	PicInfo *SearchImageByPicResponseBodyPicInfo `json:"PicInfo,omitempty" xml:"PicInfo,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// B3137727-7D6E-488C-BA21-0E034C38A879
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SearchImageByPicResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SearchImageByPicResponseBody) GoString() string {
	return s.String()
}

func (s *SearchImageByPicResponseBody) SetAuctions(v []*SearchImageByPicResponseBodyAuctions) *SearchImageByPicResponseBody {
	s.Auctions = v
	return s
}

func (s *SearchImageByPicResponseBody) SetCode(v int32) *SearchImageByPicResponseBody {
	s.Code = &v
	return s
}

func (s *SearchImageByPicResponseBody) SetHead(v *SearchImageByPicResponseBodyHead) *SearchImageByPicResponseBody {
	s.Head = v
	return s
}

func (s *SearchImageByPicResponseBody) SetMsg(v string) *SearchImageByPicResponseBody {
	s.Msg = &v
	return s
}

func (s *SearchImageByPicResponseBody) SetPicInfo(v *SearchImageByPicResponseBodyPicInfo) *SearchImageByPicResponseBody {
	s.PicInfo = v
	return s
}

func (s *SearchImageByPicResponseBody) SetRequestId(v string) *SearchImageByPicResponseBody {
	s.RequestId = &v
	return s
}

func (s *SearchImageByPicResponseBody) SetSuccess(v bool) *SearchImageByPicResponseBody {
	s.Success = &v
	return s
}

type SearchImageByPicResponseBodyAuctions struct {
	// The category of the image.
	//
	// example:
	//
	// 8888888
	CategoryId *int32 `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	// The user-defined content.
	//
	// example:
	//
	// zidingyi
	CustomContent *string `json:"CustomContent,omitempty" xml:"CustomContent,omitempty"`
	// The attribute, which is an integer.
	//
	// example:
	//
	// 2
	IntAttr *int32 `json:"IntAttr,omitempty" xml:"IntAttr,omitempty"`
	// example:
	//
	// 20
	IntAttr2 *int32 `json:"IntAttr2,omitempty" xml:"IntAttr2,omitempty"`
	IntAttr3 *int32 `json:"IntAttr3,omitempty" xml:"IntAttr3,omitempty"`
	IntAttr4 *int32 `json:"IntAttr4,omitempty" xml:"IntAttr4,omitempty"`
	// The name of the image.
	//
	// example:
	//
	// 2092061_1.jpg
	PicName *string `json:"PicName,omitempty" xml:"PicName,omitempty"`
	// The ID of the product.
	//
	// example:
	//
	// 2092061_1
	ProductId *string `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	// The similarity score of the searched image. Valid values: 0 to 1.
	//
	// >  To use this feature, you must upgrade the SDK to version 3.1.1.
	//
	// example:
	//
	// 1
	Score *float32 `json:"Score,omitempty" xml:"Score,omitempty"`
	// The score information about the image.
	//
	// > 	- This parameter is not supported. We recommend that you use the Score parameter.
	//
	// >	- The SortExprValues parameter indicates a 2-tuple in which values are separated by a semicolon (;). The first value indicates the correlation score of the returned image. A greater value indicates a higher correlation with the sample image. Different algorithms are used.
	//
	// >	- If the value of CategoryId is within the value range from 0 to 2, the value range of SortExprValues is from 0 to 7.33136443711219e+24.
	//
	// >	- If the value of CategoryId is not within the value range from 0 to 2, the value range of SortExprValues is from 0 to 5.37633353624177e+24. If the returned image is identical with the sample image, the highest correlation score is generated.
	//
	// example:
	//
	// 5.37633353624177e+24;0
	SortExprValues *string `json:"SortExprValues,omitempty" xml:"SortExprValues,omitempty"`
	// The attribute, which is a string.
	//
	// example:
	//
	// 2
	StrAttr *string `json:"StrAttr,omitempty" xml:"StrAttr,omitempty"`
	// example:
	//
	// test
	StrAttr2 *string `json:"StrAttr2,omitempty" xml:"StrAttr2,omitempty"`
	StrAttr3 *string `json:"StrAttr3,omitempty" xml:"StrAttr3,omitempty"`
	StrAttr4 *string `json:"StrAttr4,omitempty" xml:"StrAttr4,omitempty"`
}

func (s SearchImageByPicResponseBodyAuctions) String() string {
	return tea.Prettify(s)
}

func (s SearchImageByPicResponseBodyAuctions) GoString() string {
	return s.String()
}

func (s *SearchImageByPicResponseBodyAuctions) SetCategoryId(v int32) *SearchImageByPicResponseBodyAuctions {
	s.CategoryId = &v
	return s
}

func (s *SearchImageByPicResponseBodyAuctions) SetCustomContent(v string) *SearchImageByPicResponseBodyAuctions {
	s.CustomContent = &v
	return s
}

func (s *SearchImageByPicResponseBodyAuctions) SetIntAttr(v int32) *SearchImageByPicResponseBodyAuctions {
	s.IntAttr = &v
	return s
}

func (s *SearchImageByPicResponseBodyAuctions) SetIntAttr2(v int32) *SearchImageByPicResponseBodyAuctions {
	s.IntAttr2 = &v
	return s
}

func (s *SearchImageByPicResponseBodyAuctions) SetIntAttr3(v int32) *SearchImageByPicResponseBodyAuctions {
	s.IntAttr3 = &v
	return s
}

func (s *SearchImageByPicResponseBodyAuctions) SetIntAttr4(v int32) *SearchImageByPicResponseBodyAuctions {
	s.IntAttr4 = &v
	return s
}

func (s *SearchImageByPicResponseBodyAuctions) SetPicName(v string) *SearchImageByPicResponseBodyAuctions {
	s.PicName = &v
	return s
}

func (s *SearchImageByPicResponseBodyAuctions) SetProductId(v string) *SearchImageByPicResponseBodyAuctions {
	s.ProductId = &v
	return s
}

func (s *SearchImageByPicResponseBodyAuctions) SetScore(v float32) *SearchImageByPicResponseBodyAuctions {
	s.Score = &v
	return s
}

func (s *SearchImageByPicResponseBodyAuctions) SetSortExprValues(v string) *SearchImageByPicResponseBodyAuctions {
	s.SortExprValues = &v
	return s
}

func (s *SearchImageByPicResponseBodyAuctions) SetStrAttr(v string) *SearchImageByPicResponseBodyAuctions {
	s.StrAttr = &v
	return s
}

func (s *SearchImageByPicResponseBodyAuctions) SetStrAttr2(v string) *SearchImageByPicResponseBodyAuctions {
	s.StrAttr2 = &v
	return s
}

func (s *SearchImageByPicResponseBodyAuctions) SetStrAttr3(v string) *SearchImageByPicResponseBodyAuctions {
	s.StrAttr3 = &v
	return s
}

func (s *SearchImageByPicResponseBodyAuctions) SetStrAttr4(v string) *SearchImageByPicResponseBodyAuctions {
	s.StrAttr4 = &v
	return s
}

type SearchImageByPicResponseBodyHead struct {
	// The number of images returned.
	//
	// example:
	//
	// 10
	DocsFound *int32 `json:"DocsFound,omitempty" xml:"DocsFound,omitempty"`
	// The number of images that match the search conditions on the Image Search instance.
	//
	// example:
	//
	// 10000
	DocsReturn *int32 `json:"DocsReturn,omitempty" xml:"DocsReturn,omitempty"`
	// The time it takes to complete the search process. Unit: milliseconds.
	//
	// example:
	//
	// 95
	SearchTime *int32 `json:"SearchTime,omitempty" xml:"SearchTime,omitempty"`
}

func (s SearchImageByPicResponseBodyHead) String() string {
	return tea.Prettify(s)
}

func (s SearchImageByPicResponseBodyHead) GoString() string {
	return s.String()
}

func (s *SearchImageByPicResponseBodyHead) SetDocsFound(v int32) *SearchImageByPicResponseBodyHead {
	s.DocsFound = &v
	return s
}

func (s *SearchImageByPicResponseBodyHead) SetDocsReturn(v int32) *SearchImageByPicResponseBodyHead {
	s.DocsReturn = &v
	return s
}

func (s *SearchImageByPicResponseBodyHead) SetSearchTime(v int32) *SearchImageByPicResponseBodyHead {
	s.SearchTime = &v
	return s
}

type SearchImageByPicResponseBodyPicInfo struct {
	// The categories that are supported by the system.
	AllCategories []*SearchImageByPicResponseBodyPicInfoAllCategories `json:"AllCategories,omitempty" xml:"AllCategories,omitempty" type:"Repeated"`
	// The category selected by the system. If a category is specified in the request, the specified category prevails.
	//
	// example:
	//
	// 88888888
	CategoryId *int32 `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	// The recognized subjects.
	//
	// >  To use this feature, you must upgrade the SDK to version 3.1.1.
	MultiRegion []*SearchImageByPicResponseBodyPicInfoMultiRegion `json:"MultiRegion,omitempty" xml:"MultiRegion,omitempty" type:"Repeated"`
	// The result of subject recognition. The subject area of the image, in the format of x1,x2,y1,y2. Specifically, x1 and y1 specify the upper-left pixel, and x2 and y2 specify the lower-right pixel. If a subject area is specified in the request, the specified subject area prevails.
	//
	// example:
	//
	// 280,486,232,351
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s SearchImageByPicResponseBodyPicInfo) String() string {
	return tea.Prettify(s)
}

func (s SearchImageByPicResponseBodyPicInfo) GoString() string {
	return s.String()
}

func (s *SearchImageByPicResponseBodyPicInfo) SetAllCategories(v []*SearchImageByPicResponseBodyPicInfoAllCategories) *SearchImageByPicResponseBodyPicInfo {
	s.AllCategories = v
	return s
}

func (s *SearchImageByPicResponseBodyPicInfo) SetCategoryId(v int32) *SearchImageByPicResponseBodyPicInfo {
	s.CategoryId = &v
	return s
}

func (s *SearchImageByPicResponseBodyPicInfo) SetMultiRegion(v []*SearchImageByPicResponseBodyPicInfoMultiRegion) *SearchImageByPicResponseBodyPicInfo {
	s.MultiRegion = v
	return s
}

func (s *SearchImageByPicResponseBodyPicInfo) SetRegion(v string) *SearchImageByPicResponseBodyPicInfo {
	s.Region = &v
	return s
}

type SearchImageByPicResponseBodyPicInfoAllCategories struct {
	// The ID of the category.
	//
	// example:
	//
	// 88888888
	Id *int32 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the category.
	//
	// example:
	//
	// other
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s SearchImageByPicResponseBodyPicInfoAllCategories) String() string {
	return tea.Prettify(s)
}

func (s SearchImageByPicResponseBodyPicInfoAllCategories) GoString() string {
	return s.String()
}

func (s *SearchImageByPicResponseBodyPicInfoAllCategories) SetId(v int32) *SearchImageByPicResponseBodyPicInfoAllCategories {
	s.Id = &v
	return s
}

func (s *SearchImageByPicResponseBodyPicInfoAllCategories) SetName(v string) *SearchImageByPicResponseBodyPicInfoAllCategories {
	s.Name = &v
	return s
}

type SearchImageByPicResponseBodyPicInfoMultiRegion struct {
	// The result of subject recognition. The subject area of the image, in the format of x1,x2,y1,y2. Specifically, x1 and y1 specify the upper-left pixel, and x2 and y2 specify the lower-right pixel. If a subject area is specified in the request, the specified subject area prevails.
	//
	// example:
	//
	// 280,486,232,351
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s SearchImageByPicResponseBodyPicInfoMultiRegion) String() string {
	return tea.Prettify(s)
}

func (s SearchImageByPicResponseBodyPicInfoMultiRegion) GoString() string {
	return s.String()
}

func (s *SearchImageByPicResponseBodyPicInfoMultiRegion) SetRegion(v string) *SearchImageByPicResponseBodyPicInfoMultiRegion {
	s.Region = &v
	return s
}

type SearchImageByPicResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SearchImageByPicResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SearchImageByPicResponse) String() string {
	return tea.Prettify(s)
}

func (s SearchImageByPicResponse) GoString() string {
	return s.String()
}

func (s *SearchImageByPicResponse) SetHeaders(v map[string]*string) *SearchImageByPicResponse {
	s.Headers = v
	return s
}

func (s *SearchImageByPicResponse) SetStatusCode(v int32) *SearchImageByPicResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchImageByPicResponse) SetBody(v *SearchImageByPicResponseBody) *SearchImageByPicResponse {
	s.Body = v
	return s
}

type UpdateImageRequest struct {
	// The user-defined content. The value can be up to 4,096 characters in length.
	//
	// >  If you set this parameter, the response includes this parameter and its value. You can add text such as an image description.
	//
	// example:
	//
	// zidingyi
	CustomContent *string `json:"CustomContent,omitempty" xml:"CustomContent,omitempty"`
	// The name of the Image Search instance. The name can be up to 20 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// demoinstance1
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The attribute, which is an integer. The attribute can be used to filter images when you search for images. If you set this parameter, the response includes this parameter and its value.
	//
	// example:
	//
	// 2
	IntAttr *int32 `json:"IntAttr,omitempty" xml:"IntAttr,omitempty"`
	// example:
	//
	// 20
	IntAttr2 *int32 `json:"IntAttr2,omitempty" xml:"IntAttr2,omitempty"`
	IntAttr3 *int32 `json:"IntAttr3,omitempty" xml:"IntAttr3,omitempty"`
	IntAttr4 *int32 `json:"IntAttr4,omitempty" xml:"IntAttr4,omitempty"`
	// The name of the image. The name can be up to 512 characters in length.
	//
	// > 	- An image is uniquely identified by the values of the ProductId and PicName parameters.
	//
	// >	- If you add an image whose product ID (ProductId) and image name (PicName) are the same as those of an existing image, the newly added image overwrites the existing image.
	//
	// This parameter is required.
	//
	// example:
	//
	// namexxx.jpg
	PicName *string `json:"PicName,omitempty" xml:"PicName,omitempty"`
	// The ID of the product. The ID can be up to 512 characters in length.
	//
	// >  A product may have multiple images.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2092061_1
	ProductId *string `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	// The attribute, which is a string. The value can be up to 128 characters in length. The attribute can be used to filter images. If you set this parameter, the response includes this parameter and its value.
	//
	// example:
	//
	// ss
	StrAttr *string `json:"StrAttr,omitempty" xml:"StrAttr,omitempty"`
	// example:
	//
	// test
	StrAttr2 *string `json:"StrAttr2,omitempty" xml:"StrAttr2,omitempty"`
	StrAttr3 *string `json:"StrAttr3,omitempty" xml:"StrAttr3,omitempty"`
	StrAttr4 *string `json:"StrAttr4,omitempty" xml:"StrAttr4,omitempty"`
}

func (s UpdateImageRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateImageRequest) GoString() string {
	return s.String()
}

func (s *UpdateImageRequest) SetCustomContent(v string) *UpdateImageRequest {
	s.CustomContent = &v
	return s
}

func (s *UpdateImageRequest) SetInstanceName(v string) *UpdateImageRequest {
	s.InstanceName = &v
	return s
}

func (s *UpdateImageRequest) SetIntAttr(v int32) *UpdateImageRequest {
	s.IntAttr = &v
	return s
}

func (s *UpdateImageRequest) SetIntAttr2(v int32) *UpdateImageRequest {
	s.IntAttr2 = &v
	return s
}

func (s *UpdateImageRequest) SetIntAttr3(v int32) *UpdateImageRequest {
	s.IntAttr3 = &v
	return s
}

func (s *UpdateImageRequest) SetIntAttr4(v int32) *UpdateImageRequest {
	s.IntAttr4 = &v
	return s
}

func (s *UpdateImageRequest) SetPicName(v string) *UpdateImageRequest {
	s.PicName = &v
	return s
}

func (s *UpdateImageRequest) SetProductId(v string) *UpdateImageRequest {
	s.ProductId = &v
	return s
}

func (s *UpdateImageRequest) SetStrAttr(v string) *UpdateImageRequest {
	s.StrAttr = &v
	return s
}

func (s *UpdateImageRequest) SetStrAttr2(v string) *UpdateImageRequest {
	s.StrAttr2 = &v
	return s
}

func (s *UpdateImageRequest) SetStrAttr3(v string) *UpdateImageRequest {
	s.StrAttr3 = &v
	return s
}

func (s *UpdateImageRequest) SetStrAttr4(v string) *UpdateImageRequest {
	s.StrAttr4 = &v
	return s
}

type UpdateImageResponseBody struct {
	// The error code returned.
	//
	// 	- A value of 0 indicates that the operation is successful.
	//
	// 	- Values other than 0 indicate errors.
	//
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// Id of the request
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// E0845DE6-52AF-4B50-9F15-51ED4044E6AB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateImageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateImageResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateImageResponseBody) SetCode(v int32) *UpdateImageResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateImageResponseBody) SetMessage(v string) *UpdateImageResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateImageResponseBody) SetRequestId(v string) *UpdateImageResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateImageResponseBody) SetSuccess(v bool) *UpdateImageResponseBody {
	s.Success = &v
	return s
}

type UpdateImageResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateImageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateImageResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateImageResponse) GoString() string {
	return s.String()
}

func (s *UpdateImageResponse) SetHeaders(v map[string]*string) *UpdateImageResponse {
	s.Headers = v
	return s
}

func (s *UpdateImageResponse) SetStatusCode(v int32) *UpdateImageResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateImageResponse) SetBody(v *UpdateImageResponseBody) *UpdateImageResponse {
	s.Body = v
	return s
}

type Client struct {
	openapi.Client
}

func NewClient(config *openapi.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapi.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("imagesearch"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !tea.BoolValue(util.Empty(endpoint)) {
		_result = endpoint
		return _result, _err
	}

	if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(endpointMap[tea.StringValue(regionId)])) {
		_result = endpointMap[tea.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := endpointutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds an image to an Image Search instance.
//
// Description:
//
// You can call this operation to add an image to an Image Search instance.
//
// > If you want to obtain more information about the service and technical support, click [Online Consulting](https://www.aliyun.com/core/online-consult?from=aZgW6LJHr2) or join the DingTalk group (ID 35035130).
//
// ## QPS limits
//
// By default, the concurrency limit for adding an image to instances whose image capacity specifications are 0.1 million images is 1. This means that the system can process up to one request of adding an image every second. By default, the concurrency limit for adding an image to instances of other image capacity specifications is 5. This means that the system can process up to five requests of adding an image every second.
//
// @param request - AddImageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddImageResponse
func (client *Client) AddImageWithOptions(request *AddImageRequest, runtime *util.RuntimeOptions) (_result *AddImageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CategoryId)) {
		body["CategoryId"] = request.CategoryId
	}

	if !tea.BoolValue(util.IsUnset(request.Crop)) {
		body["Crop"] = request.Crop
	}

	if !tea.BoolValue(util.IsUnset(request.CustomContent)) {
		body["CustomContent"] = request.CustomContent
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceName)) {
		body["InstanceName"] = request.InstanceName
	}

	if !tea.BoolValue(util.IsUnset(request.IntAttr)) {
		body["IntAttr"] = request.IntAttr
	}

	if !tea.BoolValue(util.IsUnset(request.IntAttr2)) {
		body["IntAttr2"] = request.IntAttr2
	}

	if !tea.BoolValue(util.IsUnset(request.IntAttr3)) {
		body["IntAttr3"] = request.IntAttr3
	}

	if !tea.BoolValue(util.IsUnset(request.IntAttr4)) {
		body["IntAttr4"] = request.IntAttr4
	}

	if !tea.BoolValue(util.IsUnset(request.PicContent)) {
		body["PicContent"] = request.PicContent
	}

	if !tea.BoolValue(util.IsUnset(request.PicName)) {
		body["PicName"] = request.PicName
	}

	if !tea.BoolValue(util.IsUnset(request.ProductId)) {
		body["ProductId"] = request.ProductId
	}

	if !tea.BoolValue(util.IsUnset(request.Region)) {
		body["Region"] = request.Region
	}

	if !tea.BoolValue(util.IsUnset(request.StrAttr)) {
		body["StrAttr"] = request.StrAttr
	}

	if !tea.BoolValue(util.IsUnset(request.StrAttr2)) {
		body["StrAttr2"] = request.StrAttr2
	}

	if !tea.BoolValue(util.IsUnset(request.StrAttr3)) {
		body["StrAttr3"] = request.StrAttr3
	}

	if !tea.BoolValue(util.IsUnset(request.StrAttr4)) {
		body["StrAttr4"] = request.StrAttr4
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("AddImage"),
		Version:     tea.String("2020-12-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &AddImageResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &AddImageResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Adds an image to an Image Search instance.
//
// Description:
//
// You can call this operation to add an image to an Image Search instance.
//
// > If you want to obtain more information about the service and technical support, click [Online Consulting](https://www.aliyun.com/core/online-consult?from=aZgW6LJHr2) or join the DingTalk group (ID 35035130).
//
// ## QPS limits
//
// By default, the concurrency limit for adding an image to instances whose image capacity specifications are 0.1 million images is 1. This means that the system can process up to one request of adding an image every second. By default, the concurrency limit for adding an image to instances of other image capacity specifications is 5. This means that the system can process up to five requests of adding an image every second.
//
// @param request - AddImageRequest
//
// @return AddImageResponse
func (client *Client) AddImage(request *AddImageRequest) (_result *AddImageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddImageResponse{}
	_body, _err := client.AddImageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddImageAdvance(request *AddImageAdvanceRequest, runtime *util.RuntimeOptions) (_result *AddImageResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	securityToken, _err := client.Credential.GetSecurityToken()
	if _err != nil {
		return _result, _err
	}

	credentialType := client.Credential.GetType()
	openPlatformEndpoint := client.OpenPlatformEndpoint
	if tea.BoolValue(util.Empty(openPlatformEndpoint)) {
		openPlatformEndpoint = tea.String("openplatform.aliyuncs.com")
	}

	if tea.BoolValue(util.IsUnset(credentialType)) {
		credentialType = tea.String("access_key")
	}

	authConfig := &openapi.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		SecurityToken:   securityToken,
		Type:            credentialType,
		Endpoint:        openPlatformEndpoint,
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("ImageSearch"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	ossClient, _err := oss.NewClient(ossConfig)
	if _err != nil {
		return _result, _err
	}

	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	openapiutil.Convert(runtime, ossRuntime)
	addImageReq := &AddImageRequest{}
	openapiutil.Convert(request, addImageReq)
	if !tea.BoolValue(util.IsUnset(request.PicContentObject)) {
		authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
		if _err != nil {
			return _result, _err
		}

		ossConfig.AccessKeyId = authResponse.Body.AccessKeyId
		ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Body.Endpoint, authResponse.Body.UseAccelerate, client.EndpointType)
		ossClient, _err = oss.NewClient(ossConfig)
		if _err != nil {
			return _result, _err
		}

		fileObj = &fileform.FileField{
			Filename:    authResponse.Body.ObjectKey,
			Content:     request.PicContentObject,
			ContentType: tea.String(""),
		}
		ossHeader = &oss.PostObjectRequestHeader{
			AccessKeyId:         authResponse.Body.AccessKeyId,
			Policy:              authResponse.Body.EncodedPolicy,
			Signature:           authResponse.Body.Signature,
			Key:                 authResponse.Body.ObjectKey,
			File:                fileObj,
			SuccessActionStatus: tea.String("201"),
		}
		uploadRequest = &oss.PostObjectRequest{
			BucketName: authResponse.Body.Bucket,
			Header:     ossHeader,
		}
		_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
		if _err != nil {
			return _result, _err
		}
		addImageReq.PicContent = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	addImageResp, _err := client.AddImageWithOptions(addImageReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = addImageResp
	return _result, _err
}

// Summary:
//
// 对比图片相似值
//
// @param request - CompareSimilarByImageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CompareSimilarByImageResponse
func (client *Client) CompareSimilarByImageWithOptions(request *CompareSimilarByImageRequest, runtime *util.RuntimeOptions) (_result *CompareSimilarByImageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceName)) {
		body["InstanceName"] = request.InstanceName
	}

	if !tea.BoolValue(util.IsUnset(request.PrimaryPicContent)) {
		body["PrimaryPicContent"] = request.PrimaryPicContent
	}

	if !tea.BoolValue(util.IsUnset(request.SecondaryPicContent)) {
		body["SecondaryPicContent"] = request.SecondaryPicContent
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CompareSimilarByImage"),
		Version:     tea.String("2020-12-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &CompareSimilarByImageResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &CompareSimilarByImageResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 对比图片相似值
//
// @param request - CompareSimilarByImageRequest
//
// @return CompareSimilarByImageResponse
func (client *Client) CompareSimilarByImage(request *CompareSimilarByImageRequest) (_result *CompareSimilarByImageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CompareSimilarByImageResponse{}
	_body, _err := client.CompareSimilarByImageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CompareSimilarByImageAdvance(request *CompareSimilarByImageAdvanceRequest, runtime *util.RuntimeOptions) (_result *CompareSimilarByImageResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	securityToken, _err := client.Credential.GetSecurityToken()
	if _err != nil {
		return _result, _err
	}

	credentialType := client.Credential.GetType()
	openPlatformEndpoint := client.OpenPlatformEndpoint
	if tea.BoolValue(util.Empty(openPlatformEndpoint)) {
		openPlatformEndpoint = tea.String("openplatform.aliyuncs.com")
	}

	if tea.BoolValue(util.IsUnset(credentialType)) {
		credentialType = tea.String("access_key")
	}

	authConfig := &openapi.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		SecurityToken:   securityToken,
		Type:            credentialType,
		Endpoint:        openPlatformEndpoint,
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("ImageSearch"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	ossClient, _err := oss.NewClient(ossConfig)
	if _err != nil {
		return _result, _err
	}

	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	openapiutil.Convert(runtime, ossRuntime)
	compareSimilarByImageReq := &CompareSimilarByImageRequest{}
	openapiutil.Convert(request, compareSimilarByImageReq)
	if !tea.BoolValue(util.IsUnset(request.PrimaryPicContentObject)) {
		authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
		if _err != nil {
			return _result, _err
		}

		ossConfig.AccessKeyId = authResponse.Body.AccessKeyId
		ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Body.Endpoint, authResponse.Body.UseAccelerate, client.EndpointType)
		ossClient, _err = oss.NewClient(ossConfig)
		if _err != nil {
			return _result, _err
		}

		fileObj = &fileform.FileField{
			Filename:    authResponse.Body.ObjectKey,
			Content:     request.PrimaryPicContentObject,
			ContentType: tea.String(""),
		}
		ossHeader = &oss.PostObjectRequestHeader{
			AccessKeyId:         authResponse.Body.AccessKeyId,
			Policy:              authResponse.Body.EncodedPolicy,
			Signature:           authResponse.Body.Signature,
			Key:                 authResponse.Body.ObjectKey,
			File:                fileObj,
			SuccessActionStatus: tea.String("201"),
		}
		uploadRequest = &oss.PostObjectRequest{
			BucketName: authResponse.Body.Bucket,
			Header:     ossHeader,
		}
		_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
		if _err != nil {
			return _result, _err
		}
		compareSimilarByImageReq.PrimaryPicContent = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	if !tea.BoolValue(util.IsUnset(request.SecondaryPicContentObject)) {
		authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
		if _err != nil {
			return _result, _err
		}

		ossConfig.AccessKeyId = authResponse.Body.AccessKeyId
		ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Body.Endpoint, authResponse.Body.UseAccelerate, client.EndpointType)
		ossClient, _err = oss.NewClient(ossConfig)
		if _err != nil {
			return _result, _err
		}

		fileObj = &fileform.FileField{
			Filename:    authResponse.Body.ObjectKey,
			Content:     request.SecondaryPicContentObject,
			ContentType: tea.String(""),
		}
		ossHeader = &oss.PostObjectRequestHeader{
			AccessKeyId:         authResponse.Body.AccessKeyId,
			Policy:              authResponse.Body.EncodedPolicy,
			Signature:           authResponse.Body.Signature,
			Key:                 authResponse.Body.ObjectKey,
			File:                fileObj,
			SuccessActionStatus: tea.String("201"),
		}
		uploadRequest = &oss.PostObjectRequest{
			BucketName: authResponse.Body.Bucket,
			Header:     ossHeader,
		}
		_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
		if _err != nil {
			return _result, _err
		}
		compareSimilarByImageReq.SecondaryPicContent = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	compareSimilarByImageResp, _err := client.CompareSimilarByImageWithOptions(compareSimilarByImageReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = compareSimilarByImageResp
	return _result, _err
}

// Summary:
//
// This topic describes the syntax of the DeleteImage operation and provides examples of this operation. You can call this operation to delete images from an Image Search instance.
//
// Description:
//
// This operation deletes images from an Image Search instance.
//
// >  A success response is returned even if the specified image does not exist on the instance. Therefore, you cannot determine whether the image exists on the instance based on the response.
//
// ## QPS limits
//
// By default, the maximum number of queries supported by this operation is 20. In this case, the system can process at most 20 requests every second.
//
// @param request - DeleteImageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteImageResponse
func (client *Client) DeleteImageWithOptions(request *DeleteImageRequest, runtime *util.RuntimeOptions) (_result *DeleteImageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Filter)) {
		body["Filter"] = request.Filter
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceName)) {
		body["InstanceName"] = request.InstanceName
	}

	if !tea.BoolValue(util.IsUnset(request.IsDeleteByFilter)) {
		body["IsDeleteByFilter"] = request.IsDeleteByFilter
	}

	if !tea.BoolValue(util.IsUnset(request.PicName)) {
		body["PicName"] = request.PicName
	}

	if !tea.BoolValue(util.IsUnset(request.ProductId)) {
		body["ProductId"] = request.ProductId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteImage"),
		Version:     tea.String("2020-12-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DeleteImageResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DeleteImageResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// This topic describes the syntax of the DeleteImage operation and provides examples of this operation. You can call this operation to delete images from an Image Search instance.
//
// Description:
//
// This operation deletes images from an Image Search instance.
//
// >  A success response is returned even if the specified image does not exist on the instance. Therefore, you cannot determine whether the image exists on the instance based on the response.
//
// ## QPS limits
//
// By default, the maximum number of queries supported by this operation is 20. In this case, the system can process at most 20 requests every second.
//
// @param request - DeleteImageRequest
//
// @return DeleteImageResponse
func (client *Client) DeleteImage(request *DeleteImageRequest) (_result *DeleteImageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteImageResponse{}
	_body, _err := client.DeleteImageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// This topic describes the syntax of the Detail operation and provides examples of this operation. You can call this operation to query instance details.
//
// Description:
//
// This operation queries instance details.
//
// ## QPS limits
//
// By default, the maximum number of queries supported by this operation is 1. In this case, the system can process only 1 request every second.
//
// @param request - DetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DetailResponse
func (client *Client) DetailWithOptions(request *DetailRequest, runtime *util.RuntimeOptions) (_result *DetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceName)) {
		query["InstanceName"] = request.InstanceName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("Detail"),
		Version:     tea.String("2020-12-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DetailResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DetailResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// This topic describes the syntax of the Detail operation and provides examples of this operation. You can call this operation to query instance details.
//
// Description:
//
// This operation queries instance details.
//
// ## QPS limits
//
// By default, the maximum number of queries supported by this operation is 1. In this case, the system can process only 1 request every second.
//
// @param request - DetailRequest
//
// @return DetailResponse
func (client *Client) Detail(request *DetailRequest) (_result *DetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DetailResponse{}
	_body, _err := client.DetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// This topic describes the syntax of the DumpMeta operation and provides examples of this operation. You can call this operation to create a task for exporting metadata from an Image Search instance.
//
// Description:
//
// This operation creates a task for exporting metadata from an Image Search instance.
//
// ## QPS limits
//
// By default, the maximum number of queries supported by this operation is 1. In this case, the system can process at most 1 request every second.
//
// @param request - DumpMetaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DumpMetaResponse
func (client *Client) DumpMetaWithOptions(request *DumpMetaRequest, runtime *util.RuntimeOptions) (_result *DumpMetaResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceName)) {
		query["InstanceName"] = request.InstanceName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DumpMeta"),
		Version:     tea.String("2020-12-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DumpMetaResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DumpMetaResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// This topic describes the syntax of the DumpMeta operation and provides examples of this operation. You can call this operation to create a task for exporting metadata from an Image Search instance.
//
// Description:
//
// This operation creates a task for exporting metadata from an Image Search instance.
//
// ## QPS limits
//
// By default, the maximum number of queries supported by this operation is 1. In this case, the system can process at most 1 request every second.
//
// @param request - DumpMetaRequest
//
// @return DumpMetaResponse
func (client *Client) DumpMeta(request *DumpMetaRequest) (_result *DumpMetaResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DumpMetaResponse{}
	_body, _err := client.DumpMetaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// This topic describes the syntax of the DumpMetaList operation and provides examples of this operation. You can call this operation to query tasks that are used for exporting metadata from an Image Search instance.
//
// Description:
//
// This operation queries tasks that are used for exporting metadata from an Image Search instance.
//
// ## QPS limits
//
// By default, the maximum number of queries supported by this operation is 1. In this case, the system can process at most 1 request every second.
//
// @param request - DumpMetaListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DumpMetaListResponse
func (client *Client) DumpMetaListWithOptions(request *DumpMetaListRequest, runtime *util.RuntimeOptions) (_result *DumpMetaListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceName)) {
		query["InstanceName"] = request.InstanceName
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DumpMetaList"),
		Version:     tea.String("2020-12-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DumpMetaListResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DumpMetaListResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// This topic describes the syntax of the DumpMetaList operation and provides examples of this operation. You can call this operation to query tasks that are used for exporting metadata from an Image Search instance.
//
// Description:
//
// This operation queries tasks that are used for exporting metadata from an Image Search instance.
//
// ## QPS limits
//
// By default, the maximum number of queries supported by this operation is 1. In this case, the system can process at most 1 request every second.
//
// @param request - DumpMetaListRequest
//
// @return DumpMetaListResponse
func (client *Client) DumpMetaList(request *DumpMetaListRequest) (_result *DumpMetaListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DumpMetaListResponse{}
	_body, _err := client.DumpMetaListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// This topic describes the syntax of the IncreaseInstance operation and provides examples of this operation. You can call this operation to create a batch task on an Image Search instance.
//
// Description:
//
// This operation creates a batch task on an Image Search instance.
//
// ## QPS limits
//
// By default, the maximum number of queries supported by this operation is 1. In this case, the system can process at most 1 request every second.
//
// @param request - IncreaseInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return IncreaseInstanceResponse
func (client *Client) IncreaseInstanceWithOptions(request *IncreaseInstanceRequest, runtime *util.RuntimeOptions) (_result *IncreaseInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BucketName)) {
		query["BucketName"] = request.BucketName
	}

	if !tea.BoolValue(util.IsUnset(request.CallbackAddress)) {
		query["CallbackAddress"] = request.CallbackAddress
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceName)) {
		query["InstanceName"] = request.InstanceName
	}

	if !tea.BoolValue(util.IsUnset(request.Path)) {
		query["Path"] = request.Path
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("IncreaseInstance"),
		Version:     tea.String("2020-12-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &IncreaseInstanceResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &IncreaseInstanceResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// This topic describes the syntax of the IncreaseInstance operation and provides examples of this operation. You can call this operation to create a batch task on an Image Search instance.
//
// Description:
//
// This operation creates a batch task on an Image Search instance.
//
// ## QPS limits
//
// By default, the maximum number of queries supported by this operation is 1. In this case, the system can process at most 1 request every second.
//
// @param request - IncreaseInstanceRequest
//
// @return IncreaseInstanceResponse
func (client *Client) IncreaseInstance(request *IncreaseInstanceRequest) (_result *IncreaseInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &IncreaseInstanceResponse{}
	_body, _err := client.IncreaseInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// This topic describes the syntax of the IncreaseList operation and provides examples of this operation. You can call this operation to query batch tasks on an Image Search instance.
//
// Description:
//
// This operation queries batch tasks on an Image Search instance.
//
// ## QPS limits
//
// By default, the maximum number of queries supported by this operation is 1. In this case, the system can process at most 1 request every second.
//
// @param request - IncreaseListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return IncreaseListResponse
func (client *Client) IncreaseListWithOptions(request *IncreaseListRequest, runtime *util.RuntimeOptions) (_result *IncreaseListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BucketName)) {
		query["BucketName"] = request.BucketName
	}

	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceName)) {
		query["InstanceName"] = request.InstanceName
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Path)) {
		query["Path"] = request.Path
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("IncreaseList"),
		Version:     tea.String("2020-12-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &IncreaseListResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &IncreaseListResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// This topic describes the syntax of the IncreaseList operation and provides examples of this operation. You can call this operation to query batch tasks on an Image Search instance.
//
// Description:
//
// This operation queries batch tasks on an Image Search instance.
//
// ## QPS limits
//
// By default, the maximum number of queries supported by this operation is 1. In this case, the system can process at most 1 request every second.
//
// @param request - IncreaseListRequest
//
// @return IncreaseListResponse
func (client *Client) IncreaseList(request *IncreaseListRequest) (_result *IncreaseListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &IncreaseListResponse{}
	_body, _err := client.IncreaseListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// This topic describes the syntax of the SearchByName operation and provides examples of this operation. You can call this operation to search for images by image name on an Image Search instance.
//
// Description:
//
// This operation searches for images by image name on an Image Search instance.
//
// ## QPS limits
//
// The maximum number of queries per second is displayed in the Image Search console. The upper limit is specified when you purchase the instance. You can set the upper limit to 5 QPS or 10 QPS.
//
// @param request - SearchImageByNameRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SearchImageByNameResponse
func (client *Client) SearchImageByNameWithOptions(request *SearchImageByNameRequest, runtime *util.RuntimeOptions) (_result *SearchImageByNameResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CategoryId)) {
		body["CategoryId"] = request.CategoryId
	}

	if !tea.BoolValue(util.IsUnset(request.DistinctProductId)) {
		body["DistinctProductId"] = request.DistinctProductId
	}

	if !tea.BoolValue(util.IsUnset(request.Filter)) {
		body["Filter"] = request.Filter
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceName)) {
		body["InstanceName"] = request.InstanceName
	}

	if !tea.BoolValue(util.IsUnset(request.Num)) {
		body["Num"] = request.Num
	}

	if !tea.BoolValue(util.IsUnset(request.PicName)) {
		body["PicName"] = request.PicName
	}

	if !tea.BoolValue(util.IsUnset(request.ProductId)) {
		body["ProductId"] = request.ProductId
	}

	if !tea.BoolValue(util.IsUnset(request.Start)) {
		body["Start"] = request.Start
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("SearchImageByName"),
		Version:     tea.String("2020-12-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &SearchImageByNameResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &SearchImageByNameResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// This topic describes the syntax of the SearchByName operation and provides examples of this operation. You can call this operation to search for images by image name on an Image Search instance.
//
// Description:
//
// This operation searches for images by image name on an Image Search instance.
//
// ## QPS limits
//
// The maximum number of queries per second is displayed in the Image Search console. The upper limit is specified when you purchase the instance. You can set the upper limit to 5 QPS or 10 QPS.
//
// @param request - SearchImageByNameRequest
//
// @return SearchImageByNameResponse
func (client *Client) SearchImageByName(request *SearchImageByNameRequest) (_result *SearchImageByNameResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SearchImageByNameResponse{}
	_body, _err := client.SearchImageByNameWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// This topic describes the syntax of the SearchByPic operation and provides examples of this operation. You can call this operation to search for images by image on an Image Search Instance.
//
// Description:
//
// This operation searches for images by image name on an Image Search instance.
//
// ## QPS limits
//
// The maximum number of queries per second is displayed in the Image Search console. The upper limit is specified when you purchase the instance. You can set the upper limit to 5 QPS or 10 QPS.
//
// ## SDK release notes
//
// The Image Search SDK has been upgraded to version 3.1.1, which supports multi-subject recognition and similarity scores. For more information, see [Image Search SDK for Java](/help/en/image-search/latest/version-v3-java-sdk).
//
// @param request - SearchImageByPicRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SearchImageByPicResponse
func (client *Client) SearchImageByPicWithOptions(request *SearchImageByPicRequest, runtime *util.RuntimeOptions) (_result *SearchImageByPicResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CategoryId)) {
		body["CategoryId"] = request.CategoryId
	}

	if !tea.BoolValue(util.IsUnset(request.Crop)) {
		body["Crop"] = request.Crop
	}

	if !tea.BoolValue(util.IsUnset(request.DistinctProductId)) {
		body["DistinctProductId"] = request.DistinctProductId
	}

	if !tea.BoolValue(util.IsUnset(request.Filter)) {
		body["Filter"] = request.Filter
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceName)) {
		body["InstanceName"] = request.InstanceName
	}

	if !tea.BoolValue(util.IsUnset(request.Num)) {
		body["Num"] = request.Num
	}

	if !tea.BoolValue(util.IsUnset(request.PicContent)) {
		body["PicContent"] = request.PicContent
	}

	if !tea.BoolValue(util.IsUnset(request.Region)) {
		body["Region"] = request.Region
	}

	if !tea.BoolValue(util.IsUnset(request.Start)) {
		body["Start"] = request.Start
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("SearchImageByPic"),
		Version:     tea.String("2020-12-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &SearchImageByPicResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &SearchImageByPicResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// This topic describes the syntax of the SearchByPic operation and provides examples of this operation. You can call this operation to search for images by image on an Image Search Instance.
//
// Description:
//
// This operation searches for images by image name on an Image Search instance.
//
// ## QPS limits
//
// The maximum number of queries per second is displayed in the Image Search console. The upper limit is specified when you purchase the instance. You can set the upper limit to 5 QPS or 10 QPS.
//
// ## SDK release notes
//
// The Image Search SDK has been upgraded to version 3.1.1, which supports multi-subject recognition and similarity scores. For more information, see [Image Search SDK for Java](/help/en/image-search/latest/version-v3-java-sdk).
//
// @param request - SearchImageByPicRequest
//
// @return SearchImageByPicResponse
func (client *Client) SearchImageByPic(request *SearchImageByPicRequest) (_result *SearchImageByPicResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SearchImageByPicResponse{}
	_body, _err := client.SearchImageByPicWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SearchImageByPicAdvance(request *SearchImageByPicAdvanceRequest, runtime *util.RuntimeOptions) (_result *SearchImageByPicResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	securityToken, _err := client.Credential.GetSecurityToken()
	if _err != nil {
		return _result, _err
	}

	credentialType := client.Credential.GetType()
	openPlatformEndpoint := client.OpenPlatformEndpoint
	if tea.BoolValue(util.Empty(openPlatformEndpoint)) {
		openPlatformEndpoint = tea.String("openplatform.aliyuncs.com")
	}

	if tea.BoolValue(util.IsUnset(credentialType)) {
		credentialType = tea.String("access_key")
	}

	authConfig := &openapi.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		SecurityToken:   securityToken,
		Type:            credentialType,
		Endpoint:        openPlatformEndpoint,
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("ImageSearch"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	ossClient, _err := oss.NewClient(ossConfig)
	if _err != nil {
		return _result, _err
	}

	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	openapiutil.Convert(runtime, ossRuntime)
	searchImageByPicReq := &SearchImageByPicRequest{}
	openapiutil.Convert(request, searchImageByPicReq)
	if !tea.BoolValue(util.IsUnset(request.PicContentObject)) {
		authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
		if _err != nil {
			return _result, _err
		}

		ossConfig.AccessKeyId = authResponse.Body.AccessKeyId
		ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Body.Endpoint, authResponse.Body.UseAccelerate, client.EndpointType)
		ossClient, _err = oss.NewClient(ossConfig)
		if _err != nil {
			return _result, _err
		}

		fileObj = &fileform.FileField{
			Filename:    authResponse.Body.ObjectKey,
			Content:     request.PicContentObject,
			ContentType: tea.String(""),
		}
		ossHeader = &oss.PostObjectRequestHeader{
			AccessKeyId:         authResponse.Body.AccessKeyId,
			Policy:              authResponse.Body.EncodedPolicy,
			Signature:           authResponse.Body.Signature,
			Key:                 authResponse.Body.ObjectKey,
			File:                fileObj,
			SuccessActionStatus: tea.String("201"),
		}
		uploadRequest = &oss.PostObjectRequest{
			BucketName: authResponse.Body.Bucket,
			Header:     ossHeader,
		}
		_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
		if _err != nil {
			return _result, _err
		}
		searchImageByPicReq.PicContent = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	searchImageByPicResp, _err := client.SearchImageByPicWithOptions(searchImageByPicReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = searchImageByPicResp
	return _result, _err
}

// Summary:
//
// This topic describes the syntax of the UpdateImage operation and provides examples of this operation. You can call this operation to update image information on an Image Search instance.
//
// Description:
//
// This operation updates image information on an Image Search instance.
//
// > 	- Limits are imposed on the instance creation time.
//
// >	- This operation is supported by instances that are created in the Singapore (Singapore) region after December 2021. This operation is not supported in other regions.
//
// ## QPS limits
//
// By default, the maximum number of queries supported by this operation is 20. In this case, the system can process at most 20 requests every second.
//
// @param request - UpdateImageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateImageResponse
func (client *Client) UpdateImageWithOptions(request *UpdateImageRequest, runtime *util.RuntimeOptions) (_result *UpdateImageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IntAttr3)) {
		query["IntAttr3"] = request.IntAttr3
	}

	if !tea.BoolValue(util.IsUnset(request.IntAttr4)) {
		query["IntAttr4"] = request.IntAttr4
	}

	if !tea.BoolValue(util.IsUnset(request.StrAttr3)) {
		query["StrAttr3"] = request.StrAttr3
	}

	if !tea.BoolValue(util.IsUnset(request.StrAttr4)) {
		query["StrAttr4"] = request.StrAttr4
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CustomContent)) {
		body["CustomContent"] = request.CustomContent
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceName)) {
		body["InstanceName"] = request.InstanceName
	}

	if !tea.BoolValue(util.IsUnset(request.IntAttr)) {
		body["IntAttr"] = request.IntAttr
	}

	if !tea.BoolValue(util.IsUnset(request.IntAttr2)) {
		body["IntAttr2"] = request.IntAttr2
	}

	if !tea.BoolValue(util.IsUnset(request.PicName)) {
		body["PicName"] = request.PicName
	}

	if !tea.BoolValue(util.IsUnset(request.ProductId)) {
		body["ProductId"] = request.ProductId
	}

	if !tea.BoolValue(util.IsUnset(request.StrAttr)) {
		body["StrAttr"] = request.StrAttr
	}

	if !tea.BoolValue(util.IsUnset(request.StrAttr2)) {
		body["StrAttr2"] = request.StrAttr2
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateImage"),
		Version:     tea.String("2020-12-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &UpdateImageResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &UpdateImageResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// This topic describes the syntax of the UpdateImage operation and provides examples of this operation. You can call this operation to update image information on an Image Search instance.
//
// Description:
//
// This operation updates image information on an Image Search instance.
//
// > 	- Limits are imposed on the instance creation time.
//
// >	- This operation is supported by instances that are created in the Singapore (Singapore) region after December 2021. This operation is not supported in other regions.
//
// ## QPS limits
//
// By default, the maximum number of queries supported by this operation is 20. In this case, the system can process at most 20 requests every second.
//
// @param request - UpdateImageRequest
//
// @return UpdateImageResponse
func (client *Client) UpdateImage(request *UpdateImageRequest) (_result *UpdateImageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateImageResponse{}
	_body, _err := client.UpdateImageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
