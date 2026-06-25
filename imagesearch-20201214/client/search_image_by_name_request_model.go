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
	// The product category. For more information, see [Category reference](https://help.aliyun.com/document_detail/179184.html).
	//
	//  - For product image search, if you specify a category, the specified category is used. If you do not specify a category, the system predicts the category. You can obtain the predicted category from the response.
	//
	// <props="china">
	//
	//  - For fabric, trademark, generic, home furnishing, and industrial hardware searches, the system sets the category to 88888888 regardless of whether you specify a category.
	//
	// <props="intl">
	//
	// - For generic image search, the system sets the category to 88888888 regardless of whether you specify a category..
	//
	// example:
	//
	// 88888888
	CategoryId *int32 `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	// Specifies whether to deduplicate results based on ProductId.
	//
	// > Set this parameter to true to enable deduplication.
	//
	// example:
	//
	// false
	DistinctProductId *bool `json:"DistinctProductId,omitempty" xml:"DistinctProductId,omitempty"`
	// The filter condition. The int_attr field supports the following operators: in, not in, greater than (>), greater than or equal to (>=), less than (<), less than or equal to (<=), and equal to (=). The str_attr field supports the following operators: in, not in, equal to (=), and not equal to (!=). Multiple conditions can be connected by using AND and OR.
	//
	// Examples:
	//
	// - int_attr>=100.
	//
	// - str_attr!="value1".
	//
	// - int_attr=1000 AND str_attr="value1".
	//
	// >The maximum length is 4,096 characters.
	//
	// example:
	//
	// int_attr>=100
	Filter *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	// The name of the Image Search instance. The name can be up to 20 characters in length.
	//
	// If you have purchased an Image Search instance, log on to the [Image Search console](https://imagesearch.console.aliyun.com/) to view the instance name.
	//
	// If you have not purchased an Image Search instance, refer to [Activate the service](https://help.aliyun.com/document_detail/179178.html) and [Create an instance](https://help.aliyun.com/document_detail/66569.html).
	//
	// >The instance name is not the instance ID. Make sure you distinguish between them.
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
	// The image name.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2092061_1.jpg
	PicName *string `json:"PicName,omitempty" xml:"PicName,omitempty"`
	// The product ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2092061_1
	ProductId *string `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	// The similarity score threshold. If you specify this parameter, only images with a similarity score greater than or equal to the threshold are returned. Valid values: 0.00 to 1.00. Up to two decimal places are supported. Default value: 0.00.
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
