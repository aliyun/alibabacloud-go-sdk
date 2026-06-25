// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchImageByFilterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFilter(v string) *SearchImageByFilterRequest
	GetFilter() *string
	SetInstanceName(v string) *SearchImageByFilterRequest
	GetInstanceName() *string
	SetNum(v int32) *SearchImageByFilterRequest
	GetNum() *int32
	SetStart(v int32) *SearchImageByFilterRequest
	GetStart() *int32
}

type SearchImageByFilterRequest struct {
	// The filter conditions. The operators supported by int_attr include in, not in, greater than (>), greater than or equal to (>=), less than (<), less than or equal to (<=), and equal to (=). The operators supported by str_attr include in, not in, equal to (=), and not equal to (!=). Multiple conditions can be connected by AND and OR.
	//
	// Examples:
	//
	// - int_attr >= 100.
	//
	// - str_attr != "value1".
	//
	// - int_attr = 1000 AND str_attr = "value1".
	//
	// >A maximum of 4,096 characters are supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// int_attr=1000 AND str_attr="value1"
	Filter *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	// The name of the Image Search instance. The name can be up to 20 characters in length.
	//
	// If you have purchased an Image Search instance, you can log on to the [Image Search console](https://imagesearch.console.aliyun.com/) to view the instance.
	//
	// If you have not purchased an Image Search instance, see [Activate the service](https://help.aliyun.com/document_detail/179178.html) and [Create an instance](https://help.aliyun.com/document_detail/66569.html).
	//
	// >The instance name is not the instance ID. Make sure that you can distinguish between them.
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
	// The start position of the results to return. Valid values: 0 to 499. Default value: 0.
	//
	// example:
	//
	// 0
	Start *int32 `json:"Start,omitempty" xml:"Start,omitempty"`
}

func (s SearchImageByFilterRequest) String() string {
	return dara.Prettify(s)
}

func (s SearchImageByFilterRequest) GoString() string {
	return s.String()
}

func (s *SearchImageByFilterRequest) GetFilter() *string {
	return s.Filter
}

func (s *SearchImageByFilterRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *SearchImageByFilterRequest) GetNum() *int32 {
	return s.Num
}

func (s *SearchImageByFilterRequest) GetStart() *int32 {
	return s.Start
}

func (s *SearchImageByFilterRequest) SetFilter(v string) *SearchImageByFilterRequest {
	s.Filter = &v
	return s
}

func (s *SearchImageByFilterRequest) SetInstanceName(v string) *SearchImageByFilterRequest {
	s.InstanceName = &v
	return s
}

func (s *SearchImageByFilterRequest) SetNum(v int32) *SearchImageByFilterRequest {
	s.Num = &v
	return s
}

func (s *SearchImageByFilterRequest) SetStart(v int32) *SearchImageByFilterRequest {
	s.Start = &v
	return s
}

func (s *SearchImageByFilterRequest) Validate() error {
	return dara.Validate(s)
}
