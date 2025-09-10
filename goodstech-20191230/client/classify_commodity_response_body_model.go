// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClassifyCommodityResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ClassifyCommodityResponseBodyData) *ClassifyCommodityResponseBody
	GetData() *ClassifyCommodityResponseBodyData
	SetRequestId(v string) *ClassifyCommodityResponseBody
	GetRequestId() *string
}

type ClassifyCommodityResponseBody struct {
	Data *ClassifyCommodityResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 87C5AF93-F641-54C2-873D-0501E637489C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ClassifyCommodityResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ClassifyCommodityResponseBody) GoString() string {
	return s.String()
}

func (s *ClassifyCommodityResponseBody) GetData() *ClassifyCommodityResponseBodyData {
	return s.Data
}

func (s *ClassifyCommodityResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ClassifyCommodityResponseBody) SetData(v *ClassifyCommodityResponseBodyData) *ClassifyCommodityResponseBody {
	s.Data = v
	return s
}

func (s *ClassifyCommodityResponseBody) SetRequestId(v string) *ClassifyCommodityResponseBody {
	s.RequestId = &v
	return s
}

func (s *ClassifyCommodityResponseBody) Validate() error {
	return dara.Validate(s)
}

type ClassifyCommodityResponseBodyData struct {
	Categories []*ClassifyCommodityResponseBodyDataCategories `json:"Categories,omitempty" xml:"Categories,omitempty" type:"Repeated"`
}

func (s ClassifyCommodityResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ClassifyCommodityResponseBodyData) GoString() string {
	return s.String()
}

func (s *ClassifyCommodityResponseBodyData) GetCategories() []*ClassifyCommodityResponseBodyDataCategories {
	return s.Categories
}

func (s *ClassifyCommodityResponseBodyData) SetCategories(v []*ClassifyCommodityResponseBodyDataCategories) *ClassifyCommodityResponseBodyData {
	s.Categories = v
	return s
}

func (s *ClassifyCommodityResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type ClassifyCommodityResponseBodyDataCategories struct {
	// example:
	//
	// 584
	CategoryId   *string `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	CategoryName *string `json:"CategoryName,omitempty" xml:"CategoryName,omitempty"`
	// example:
	//
	// 0.417248
	Score *float32 `json:"Score,omitempty" xml:"Score,omitempty"`
}

func (s ClassifyCommodityResponseBodyDataCategories) String() string {
	return dara.Prettify(s)
}

func (s ClassifyCommodityResponseBodyDataCategories) GoString() string {
	return s.String()
}

func (s *ClassifyCommodityResponseBodyDataCategories) GetCategoryId() *string {
	return s.CategoryId
}

func (s *ClassifyCommodityResponseBodyDataCategories) GetCategoryName() *string {
	return s.CategoryName
}

func (s *ClassifyCommodityResponseBodyDataCategories) GetScore() *float32 {
	return s.Score
}

func (s *ClassifyCommodityResponseBodyDataCategories) SetCategoryId(v string) *ClassifyCommodityResponseBodyDataCategories {
	s.CategoryId = &v
	return s
}

func (s *ClassifyCommodityResponseBodyDataCategories) SetCategoryName(v string) *ClassifyCommodityResponseBodyDataCategories {
	s.CategoryName = &v
	return s
}

func (s *ClassifyCommodityResponseBodyDataCategories) SetScore(v float32) *ClassifyCommodityResponseBodyDataCategories {
	s.Score = &v
	return s
}

func (s *ClassifyCommodityResponseBodyDataCategories) Validate() error {
	return dara.Validate(s)
}
