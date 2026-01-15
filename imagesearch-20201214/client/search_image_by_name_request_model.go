// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchImageByNameRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCategoryId(v int32) *SearchImageByNameRequest
	GetCategoryId() *int32
	SetDistinctProductId(v bool) *SearchImageByNameRequest
	GetDistinctProductId() *bool
	SetFilter(v string) *SearchImageByNameRequest
	GetFilter() *string
	SetInstanceName(v string) *SearchImageByNameRequest
	GetInstanceName() *string
	SetNum(v int32) *SearchImageByNameRequest
	GetNum() *int32
	SetPicName(v string) *SearchImageByNameRequest
	GetPicName() *string
	SetProductId(v string) *SearchImageByNameRequest
	GetProductId() *string
	SetScoreThreshold(v string) *SearchImageByNameRequest
	GetScoreThreshold() *string
	SetStart(v int32) *SearchImageByNameRequest
	GetStart() *int32
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
	ProductId      *string `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	ScoreThreshold *string `json:"ScoreThreshold,omitempty" xml:"ScoreThreshold,omitempty"`
	// The number of the image to return. Valid values: 0 to 499. Default value: 0.
	//
	// example:
	//
	// 0
	Start *int32 `json:"Start,omitempty" xml:"Start,omitempty"`
}

func (s SearchImageByNameRequest) String() string {
	return dara.Prettify(s)
}

func (s SearchImageByNameRequest) GoString() string {
	return s.String()
}

func (s *SearchImageByNameRequest) GetCategoryId() *int32 {
	return s.CategoryId
}

func (s *SearchImageByNameRequest) GetDistinctProductId() *bool {
	return s.DistinctProductId
}

func (s *SearchImageByNameRequest) GetFilter() *string {
	return s.Filter
}

func (s *SearchImageByNameRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *SearchImageByNameRequest) GetNum() *int32 {
	return s.Num
}

func (s *SearchImageByNameRequest) GetPicName() *string {
	return s.PicName
}

func (s *SearchImageByNameRequest) GetProductId() *string {
	return s.ProductId
}

func (s *SearchImageByNameRequest) GetScoreThreshold() *string {
	return s.ScoreThreshold
}

func (s *SearchImageByNameRequest) GetStart() *int32 {
	return s.Start
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

func (s *SearchImageByNameRequest) SetScoreThreshold(v string) *SearchImageByNameRequest {
	s.ScoreThreshold = &v
	return s
}

func (s *SearchImageByNameRequest) SetStart(v int32) *SearchImageByNameRequest {
	s.Start = &v
	return s
}

func (s *SearchImageByNameRequest) Validate() error {
	return dara.Validate(s)
}
