// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchImageByTextRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDistinctProductId(v bool) *SearchImageByTextRequest
	GetDistinctProductId() *bool
	SetFilter(v string) *SearchImageByTextRequest
	GetFilter() *string
	SetInstanceName(v string) *SearchImageByTextRequest
	GetInstanceName() *string
	SetNum(v int32) *SearchImageByTextRequest
	GetNum() *int32
	SetScoreThreshold(v string) *SearchImageByTextRequest
	GetScoreThreshold() *string
	SetStart(v int32) *SearchImageByTextRequest
	GetStart() *int32
	SetText(v string) *SearchImageByTextRequest
	GetText() *string
}

type SearchImageByTextRequest struct {
	// If this parameter is set to true, duplicate data is removed based on the ProductId field during the search.
	//
	// example:
	//
	// false
	DistinctProductId *bool `json:"DistinctProductId,omitempty" xml:"DistinctProductId,omitempty"`
	// The filter condition. The int_attr field supports the following operators: in, not in, greater than (>), greater than or equal to (>=), less than (<), less than or equal to (<=), and equal to (=). The str_attr field supports the following operators: in, not in, equal to (=), and not equal to (!=). You can use AND and OR to connect multiple conditions.
	//
	// Examples:
	//
	// - int_attr >= 100.
	//
	// - str_attr != "value1".
	//
	// - int_attr = 1000 AND str_attr = "value1".
	//
	// >The filter condition can be up to 4,096 characters in length.
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
	// The similarity score threshold. After you specify this parameter, only images whose similarity scores are greater than or equal to the threshold are returned. Valid values: 0.00 to 1.00. The value supports up to two decimal places. Default value: 0.00.
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
	// The description text for searching images. Chinese and English are supported.
	//
	// >The text can be up to 512 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// xxxx
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s SearchImageByTextRequest) String() string {
	return dara.Prettify(s)
}

func (s SearchImageByTextRequest) GoString() string {
	return s.String()
}

func (s *SearchImageByTextRequest) GetDistinctProductId() *bool {
	return s.DistinctProductId
}

func (s *SearchImageByTextRequest) GetFilter() *string {
	return s.Filter
}

func (s *SearchImageByTextRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *SearchImageByTextRequest) GetNum() *int32 {
	return s.Num
}

func (s *SearchImageByTextRequest) GetScoreThreshold() *string {
	return s.ScoreThreshold
}

func (s *SearchImageByTextRequest) GetStart() *int32 {
	return s.Start
}

func (s *SearchImageByTextRequest) GetText() *string {
	return s.Text
}

func (s *SearchImageByTextRequest) SetDistinctProductId(v bool) *SearchImageByTextRequest {
	s.DistinctProductId = &v
	return s
}

func (s *SearchImageByTextRequest) SetFilter(v string) *SearchImageByTextRequest {
	s.Filter = &v
	return s
}

func (s *SearchImageByTextRequest) SetInstanceName(v string) *SearchImageByTextRequest {
	s.InstanceName = &v
	return s
}

func (s *SearchImageByTextRequest) SetNum(v int32) *SearchImageByTextRequest {
	s.Num = &v
	return s
}

func (s *SearchImageByTextRequest) SetScoreThreshold(v string) *SearchImageByTextRequest {
	s.ScoreThreshold = &v
	return s
}

func (s *SearchImageByTextRequest) SetStart(v int32) *SearchImageByTextRequest {
	s.Start = &v
	return s
}

func (s *SearchImageByTextRequest) SetText(v string) *SearchImageByTextRequest {
	s.Text = &v
	return s
}

func (s *SearchImageByTextRequest) Validate() error {
	return dara.Validate(s)
}
