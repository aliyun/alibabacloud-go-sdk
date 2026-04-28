// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAdvisorChecksFoPagesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAssumeAliyunId(v int64) *DescribeAdvisorChecksFoPagesShrinkRequest
	GetAssumeAliyunId() *int64
	SetBizCategory(v string) *DescribeAdvisorChecksFoPagesShrinkRequest
	GetBizCategory() *string
	SetCategory(v string) *DescribeAdvisorChecksFoPagesShrinkRequest
	GetCategory() *string
	SetCheckTypesShrink(v string) *DescribeAdvisorChecksFoPagesShrinkRequest
	GetCheckTypesShrink() *string
	SetLanguage(v string) *DescribeAdvisorChecksFoPagesShrinkRequest
	GetLanguage() *string
	SetName(v string) *DescribeAdvisorChecksFoPagesShrinkRequest
	GetName() *string
	SetPageNumber(v int32) *DescribeAdvisorChecksFoPagesShrinkRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeAdvisorChecksFoPagesShrinkRequest
	GetPageSize() *int32
	SetProduct(v string) *DescribeAdvisorChecksFoPagesShrinkRequest
	GetProduct() *string
	SetSource(v string) *DescribeAdvisorChecksFoPagesShrinkRequest
	GetSource() *string
	SetStatus(v string) *DescribeAdvisorChecksFoPagesShrinkRequest
	GetStatus() *string
	SetToken(v string) *DescribeAdvisorChecksFoPagesShrinkRequest
	GetToken() *string
}

type DescribeAdvisorChecksFoPagesShrinkRequest struct {
	// example:
	//
	// 11*********35
	AssumeAliyunId *int64 `json:"AssumeAliyunId,omitempty" xml:"AssumeAliyunId,omitempty"`
	// example:
	//
	// 2
	BizCategory *string `json:"BizCategory,omitempty" xml:"BizCategory,omitempty"`
	// example:
	//
	// *
	Category         *string `json:"Category,omitempty" xml:"Category,omitempty"`
	CheckTypesShrink *string `json:"CheckTypes,omitempty" xml:"CheckTypes,omitempty"`
	Language         *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// example:
	//
	// ****
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// ecs
	Product *string `json:"Product,omitempty" xml:"Product,omitempty"`
	// example:
	//
	// *
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// example:
	//
	// enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// ***
	Token *string `json:"Token,omitempty" xml:"Token,omitempty"`
}

func (s DescribeAdvisorChecksFoPagesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAdvisorChecksFoPagesShrinkRequest) GoString() string {
	return s.String()
}

func (s *DescribeAdvisorChecksFoPagesShrinkRequest) GetAssumeAliyunId() *int64 {
	return s.AssumeAliyunId
}

func (s *DescribeAdvisorChecksFoPagesShrinkRequest) GetBizCategory() *string {
	return s.BizCategory
}

func (s *DescribeAdvisorChecksFoPagesShrinkRequest) GetCategory() *string {
	return s.Category
}

func (s *DescribeAdvisorChecksFoPagesShrinkRequest) GetCheckTypesShrink() *string {
	return s.CheckTypesShrink
}

func (s *DescribeAdvisorChecksFoPagesShrinkRequest) GetLanguage() *string {
	return s.Language
}

func (s *DescribeAdvisorChecksFoPagesShrinkRequest) GetName() *string {
	return s.Name
}

func (s *DescribeAdvisorChecksFoPagesShrinkRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeAdvisorChecksFoPagesShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeAdvisorChecksFoPagesShrinkRequest) GetProduct() *string {
	return s.Product
}

func (s *DescribeAdvisorChecksFoPagesShrinkRequest) GetSource() *string {
	return s.Source
}

func (s *DescribeAdvisorChecksFoPagesShrinkRequest) GetStatus() *string {
	return s.Status
}

func (s *DescribeAdvisorChecksFoPagesShrinkRequest) GetToken() *string {
	return s.Token
}

func (s *DescribeAdvisorChecksFoPagesShrinkRequest) SetAssumeAliyunId(v int64) *DescribeAdvisorChecksFoPagesShrinkRequest {
	s.AssumeAliyunId = &v
	return s
}

func (s *DescribeAdvisorChecksFoPagesShrinkRequest) SetBizCategory(v string) *DescribeAdvisorChecksFoPagesShrinkRequest {
	s.BizCategory = &v
	return s
}

func (s *DescribeAdvisorChecksFoPagesShrinkRequest) SetCategory(v string) *DescribeAdvisorChecksFoPagesShrinkRequest {
	s.Category = &v
	return s
}

func (s *DescribeAdvisorChecksFoPagesShrinkRequest) SetCheckTypesShrink(v string) *DescribeAdvisorChecksFoPagesShrinkRequest {
	s.CheckTypesShrink = &v
	return s
}

func (s *DescribeAdvisorChecksFoPagesShrinkRequest) SetLanguage(v string) *DescribeAdvisorChecksFoPagesShrinkRequest {
	s.Language = &v
	return s
}

func (s *DescribeAdvisorChecksFoPagesShrinkRequest) SetName(v string) *DescribeAdvisorChecksFoPagesShrinkRequest {
	s.Name = &v
	return s
}

func (s *DescribeAdvisorChecksFoPagesShrinkRequest) SetPageNumber(v int32) *DescribeAdvisorChecksFoPagesShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeAdvisorChecksFoPagesShrinkRequest) SetPageSize(v int32) *DescribeAdvisorChecksFoPagesShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeAdvisorChecksFoPagesShrinkRequest) SetProduct(v string) *DescribeAdvisorChecksFoPagesShrinkRequest {
	s.Product = &v
	return s
}

func (s *DescribeAdvisorChecksFoPagesShrinkRequest) SetSource(v string) *DescribeAdvisorChecksFoPagesShrinkRequest {
	s.Source = &v
	return s
}

func (s *DescribeAdvisorChecksFoPagesShrinkRequest) SetStatus(v string) *DescribeAdvisorChecksFoPagesShrinkRequest {
	s.Status = &v
	return s
}

func (s *DescribeAdvisorChecksFoPagesShrinkRequest) SetToken(v string) *DescribeAdvisorChecksFoPagesShrinkRequest {
	s.Token = &v
	return s
}

func (s *DescribeAdvisorChecksFoPagesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
