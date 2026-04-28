// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAdvisorChecksFoPagesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAssumeAliyunId(v int64) *DescribeAdvisorChecksFoPagesRequest
	GetAssumeAliyunId() *int64
	SetBizCategory(v string) *DescribeAdvisorChecksFoPagesRequest
	GetBizCategory() *string
	SetCategory(v string) *DescribeAdvisorChecksFoPagesRequest
	GetCategory() *string
	SetCheckTypes(v []*int64) *DescribeAdvisorChecksFoPagesRequest
	GetCheckTypes() []*int64
	SetLanguage(v string) *DescribeAdvisorChecksFoPagesRequest
	GetLanguage() *string
	SetName(v string) *DescribeAdvisorChecksFoPagesRequest
	GetName() *string
	SetPageNumber(v int32) *DescribeAdvisorChecksFoPagesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeAdvisorChecksFoPagesRequest
	GetPageSize() *int32
	SetProduct(v string) *DescribeAdvisorChecksFoPagesRequest
	GetProduct() *string
	SetSource(v string) *DescribeAdvisorChecksFoPagesRequest
	GetSource() *string
	SetStatus(v string) *DescribeAdvisorChecksFoPagesRequest
	GetStatus() *string
	SetToken(v string) *DescribeAdvisorChecksFoPagesRequest
	GetToken() *string
}

type DescribeAdvisorChecksFoPagesRequest struct {
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
	Category   *string  `json:"Category,omitempty" xml:"Category,omitempty"`
	CheckTypes []*int64 `json:"CheckTypes,omitempty" xml:"CheckTypes,omitempty" type:"Repeated"`
	Language   *string  `json:"Language,omitempty" xml:"Language,omitempty"`
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

func (s DescribeAdvisorChecksFoPagesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAdvisorChecksFoPagesRequest) GoString() string {
	return s.String()
}

func (s *DescribeAdvisorChecksFoPagesRequest) GetAssumeAliyunId() *int64 {
	return s.AssumeAliyunId
}

func (s *DescribeAdvisorChecksFoPagesRequest) GetBizCategory() *string {
	return s.BizCategory
}

func (s *DescribeAdvisorChecksFoPagesRequest) GetCategory() *string {
	return s.Category
}

func (s *DescribeAdvisorChecksFoPagesRequest) GetCheckTypes() []*int64 {
	return s.CheckTypes
}

func (s *DescribeAdvisorChecksFoPagesRequest) GetLanguage() *string {
	return s.Language
}

func (s *DescribeAdvisorChecksFoPagesRequest) GetName() *string {
	return s.Name
}

func (s *DescribeAdvisorChecksFoPagesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeAdvisorChecksFoPagesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeAdvisorChecksFoPagesRequest) GetProduct() *string {
	return s.Product
}

func (s *DescribeAdvisorChecksFoPagesRequest) GetSource() *string {
	return s.Source
}

func (s *DescribeAdvisorChecksFoPagesRequest) GetStatus() *string {
	return s.Status
}

func (s *DescribeAdvisorChecksFoPagesRequest) GetToken() *string {
	return s.Token
}

func (s *DescribeAdvisorChecksFoPagesRequest) SetAssumeAliyunId(v int64) *DescribeAdvisorChecksFoPagesRequest {
	s.AssumeAliyunId = &v
	return s
}

func (s *DescribeAdvisorChecksFoPagesRequest) SetBizCategory(v string) *DescribeAdvisorChecksFoPagesRequest {
	s.BizCategory = &v
	return s
}

func (s *DescribeAdvisorChecksFoPagesRequest) SetCategory(v string) *DescribeAdvisorChecksFoPagesRequest {
	s.Category = &v
	return s
}

func (s *DescribeAdvisorChecksFoPagesRequest) SetCheckTypes(v []*int64) *DescribeAdvisorChecksFoPagesRequest {
	s.CheckTypes = v
	return s
}

func (s *DescribeAdvisorChecksFoPagesRequest) SetLanguage(v string) *DescribeAdvisorChecksFoPagesRequest {
	s.Language = &v
	return s
}

func (s *DescribeAdvisorChecksFoPagesRequest) SetName(v string) *DescribeAdvisorChecksFoPagesRequest {
	s.Name = &v
	return s
}

func (s *DescribeAdvisorChecksFoPagesRequest) SetPageNumber(v int32) *DescribeAdvisorChecksFoPagesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeAdvisorChecksFoPagesRequest) SetPageSize(v int32) *DescribeAdvisorChecksFoPagesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeAdvisorChecksFoPagesRequest) SetProduct(v string) *DescribeAdvisorChecksFoPagesRequest {
	s.Product = &v
	return s
}

func (s *DescribeAdvisorChecksFoPagesRequest) SetSource(v string) *DescribeAdvisorChecksFoPagesRequest {
	s.Source = &v
	return s
}

func (s *DescribeAdvisorChecksFoPagesRequest) SetStatus(v string) *DescribeAdvisorChecksFoPagesRequest {
	s.Status = &v
	return s
}

func (s *DescribeAdvisorChecksFoPagesRequest) SetToken(v string) *DescribeAdvisorChecksFoPagesRequest {
	s.Token = &v
	return s
}

func (s *DescribeAdvisorChecksFoPagesRequest) Validate() error {
	return dara.Validate(s)
}
