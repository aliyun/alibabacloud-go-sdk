// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSensitiveSwitchRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCategoryName(v string) *DescribeSensitiveSwitchRequest
	GetCategoryName() *string
	SetCurrentPage(v int32) *DescribeSensitiveSwitchRequest
	GetCurrentPage() *int32
	SetLang(v string) *DescribeSensitiveSwitchRequest
	GetLang() *string
	SetPageSize(v int32) *DescribeSensitiveSwitchRequest
	GetPageSize() *int32
	SetParentCategory(v string) *DescribeSensitiveSwitchRequest
	GetParentCategory() *string
	SetSensitiveCategory(v string) *DescribeSensitiveSwitchRequest
	GetSensitiveCategory() *string
	SetSensitiveLevel(v string) *DescribeSensitiveSwitchRequest
	GetSensitiveLevel() *string
	SetSwitchStatus(v string) *DescribeSensitiveSwitchRequest
	GetSwitchStatus() *string
}

type DescribeSensitiveSwitchRequest struct {
	CategoryName *string `json:"CategoryName,omitempty" xml:"CategoryName,omitempty"`
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// universal_industry_template
	ParentCategory *string `json:"ParentCategory,omitempty" xml:"ParentCategory,omitempty"`
	// example:
	//
	// id_card
	SensitiveCategory *string `json:"SensitiveCategory,omitempty" xml:"SensitiveCategory,omitempty"`
	// example:
	//
	// S3
	SensitiveLevel *string `json:"SensitiveLevel,omitempty" xml:"SensitiveLevel,omitempty"`
	// example:
	//
	// 1
	SwitchStatus *string `json:"SwitchStatus,omitempty" xml:"SwitchStatus,omitempty"`
}

func (s DescribeSensitiveSwitchRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSensitiveSwitchRequest) GoString() string {
	return s.String()
}

func (s *DescribeSensitiveSwitchRequest) GetCategoryName() *string {
	return s.CategoryName
}

func (s *DescribeSensitiveSwitchRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeSensitiveSwitchRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeSensitiveSwitchRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeSensitiveSwitchRequest) GetParentCategory() *string {
	return s.ParentCategory
}

func (s *DescribeSensitiveSwitchRequest) GetSensitiveCategory() *string {
	return s.SensitiveCategory
}

func (s *DescribeSensitiveSwitchRequest) GetSensitiveLevel() *string {
	return s.SensitiveLevel
}

func (s *DescribeSensitiveSwitchRequest) GetSwitchStatus() *string {
	return s.SwitchStatus
}

func (s *DescribeSensitiveSwitchRequest) SetCategoryName(v string) *DescribeSensitiveSwitchRequest {
	s.CategoryName = &v
	return s
}

func (s *DescribeSensitiveSwitchRequest) SetCurrentPage(v int32) *DescribeSensitiveSwitchRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeSensitiveSwitchRequest) SetLang(v string) *DescribeSensitiveSwitchRequest {
	s.Lang = &v
	return s
}

func (s *DescribeSensitiveSwitchRequest) SetPageSize(v int32) *DescribeSensitiveSwitchRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeSensitiveSwitchRequest) SetParentCategory(v string) *DescribeSensitiveSwitchRequest {
	s.ParentCategory = &v
	return s
}

func (s *DescribeSensitiveSwitchRequest) SetSensitiveCategory(v string) *DescribeSensitiveSwitchRequest {
	s.SensitiveCategory = &v
	return s
}

func (s *DescribeSensitiveSwitchRequest) SetSensitiveLevel(v string) *DescribeSensitiveSwitchRequest {
	s.SensitiveLevel = &v
	return s
}

func (s *DescribeSensitiveSwitchRequest) SetSwitchStatus(v string) *DescribeSensitiveSwitchRequest {
	s.SwitchStatus = &v
	return s
}

func (s *DescribeSensitiveSwitchRequest) Validate() error {
	return dara.Validate(s)
}
