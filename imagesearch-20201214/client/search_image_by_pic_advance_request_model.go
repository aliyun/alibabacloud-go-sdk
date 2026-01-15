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
	Region         *string `json:"Region,omitempty" xml:"Region,omitempty"`
	ScoreThreshold *string `json:"ScoreThreshold,omitempty" xml:"ScoreThreshold,omitempty"`
	// The number of the image to return. Valid values: 0 to 499. Default value: 0.
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
