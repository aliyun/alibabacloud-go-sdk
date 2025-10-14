// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProductsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *ListProductsShrinkRequest
	GetLang() *string
	SetMaxResults(v int32) *ListProductsShrinkRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListProductsShrinkRequest
	GetNextToken() *string
	SetProductIdsShrink(v string) *ListProductsShrinkRequest
	GetProductIdsShrink() *string
	SetProductName(v string) *ListProductsShrinkRequest
	GetProductName() *string
	SetProductType(v string) *ListProductsShrinkRequest
	GetProductType() *string
	SetRegionId(v string) *ListProductsShrinkRequest
	GetRegionId() *string
	SetRoleFor(v int64) *ListProductsShrinkRequest
	GetRoleFor() *int64
	SetVendorId(v string) *ListProductsShrinkRequest
	GetVendorId() *string
}

type ListProductsShrinkRequest struct {
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// if can be null:
	// true
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// AAAAAUqcj6VO4E3ECWIrFczs****
	NextToken        *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	ProductIdsShrink *string `json:"ProductIds,omitempty" xml:"ProductIds,omitempty"`
	// example:
	//
	// sas
	ProductName *string `json:"ProductName,omitempty" xml:"ProductName,omitempty"`
	// example:
	//
	// preset
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// 173326*******
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// example:
	//
	// vd-qlsw5eocx94w9
	VendorId *string `json:"VendorId,omitempty" xml:"VendorId,omitempty"`
}

func (s ListProductsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListProductsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListProductsShrinkRequest) GetLang() *string {
	return s.Lang
}

func (s *ListProductsShrinkRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListProductsShrinkRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListProductsShrinkRequest) GetProductIdsShrink() *string {
	return s.ProductIdsShrink
}

func (s *ListProductsShrinkRequest) GetProductName() *string {
	return s.ProductName
}

func (s *ListProductsShrinkRequest) GetProductType() *string {
	return s.ProductType
}

func (s *ListProductsShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListProductsShrinkRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *ListProductsShrinkRequest) GetVendorId() *string {
	return s.VendorId
}

func (s *ListProductsShrinkRequest) SetLang(v string) *ListProductsShrinkRequest {
	s.Lang = &v
	return s
}

func (s *ListProductsShrinkRequest) SetMaxResults(v int32) *ListProductsShrinkRequest {
	s.MaxResults = &v
	return s
}

func (s *ListProductsShrinkRequest) SetNextToken(v string) *ListProductsShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *ListProductsShrinkRequest) SetProductIdsShrink(v string) *ListProductsShrinkRequest {
	s.ProductIdsShrink = &v
	return s
}

func (s *ListProductsShrinkRequest) SetProductName(v string) *ListProductsShrinkRequest {
	s.ProductName = &v
	return s
}

func (s *ListProductsShrinkRequest) SetProductType(v string) *ListProductsShrinkRequest {
	s.ProductType = &v
	return s
}

func (s *ListProductsShrinkRequest) SetRegionId(v string) *ListProductsShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *ListProductsShrinkRequest) SetRoleFor(v int64) *ListProductsShrinkRequest {
	s.RoleFor = &v
	return s
}

func (s *ListProductsShrinkRequest) SetVendorId(v string) *ListProductsShrinkRequest {
	s.VendorId = &v
	return s
}

func (s *ListProductsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
