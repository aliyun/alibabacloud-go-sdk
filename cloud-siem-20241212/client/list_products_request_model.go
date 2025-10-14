// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProductsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *ListProductsRequest
	GetLang() *string
	SetMaxResults(v int32) *ListProductsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListProductsRequest
	GetNextToken() *string
	SetProductIds(v []*string) *ListProductsRequest
	GetProductIds() []*string
	SetProductName(v string) *ListProductsRequest
	GetProductName() *string
	SetProductType(v string) *ListProductsRequest
	GetProductType() *string
	SetRegionId(v string) *ListProductsRequest
	GetRegionId() *string
	SetRoleFor(v int64) *ListProductsRequest
	GetRoleFor() *int64
	SetVendorId(v string) *ListProductsRequest
	GetVendorId() *string
}

type ListProductsRequest struct {
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
	NextToken  *string   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	ProductIds []*string `json:"ProductIds,omitempty" xml:"ProductIds,omitempty" type:"Repeated"`
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

func (s ListProductsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListProductsRequest) GoString() string {
	return s.String()
}

func (s *ListProductsRequest) GetLang() *string {
	return s.Lang
}

func (s *ListProductsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListProductsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListProductsRequest) GetProductIds() []*string {
	return s.ProductIds
}

func (s *ListProductsRequest) GetProductName() *string {
	return s.ProductName
}

func (s *ListProductsRequest) GetProductType() *string {
	return s.ProductType
}

func (s *ListProductsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListProductsRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *ListProductsRequest) GetVendorId() *string {
	return s.VendorId
}

func (s *ListProductsRequest) SetLang(v string) *ListProductsRequest {
	s.Lang = &v
	return s
}

func (s *ListProductsRequest) SetMaxResults(v int32) *ListProductsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListProductsRequest) SetNextToken(v string) *ListProductsRequest {
	s.NextToken = &v
	return s
}

func (s *ListProductsRequest) SetProductIds(v []*string) *ListProductsRequest {
	s.ProductIds = v
	return s
}

func (s *ListProductsRequest) SetProductName(v string) *ListProductsRequest {
	s.ProductName = &v
	return s
}

func (s *ListProductsRequest) SetProductType(v string) *ListProductsRequest {
	s.ProductType = &v
	return s
}

func (s *ListProductsRequest) SetRegionId(v string) *ListProductsRequest {
	s.RegionId = &v
	return s
}

func (s *ListProductsRequest) SetRoleFor(v int64) *ListProductsRequest {
	s.RoleFor = &v
	return s
}

func (s *ListProductsRequest) SetVendorId(v string) *ListProductsRequest {
	s.VendorId = &v
	return s
}

func (s *ListProductsRequest) Validate() error {
	return dara.Validate(s)
}
