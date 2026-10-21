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
	// The language of the response. Valid values:
	//
	// - **zh*	- (default): Chinese.
	//
	// - **en**: English.
	//
	// Note: This parameter is required in actual calls. If this parameter is not specified, the service returns a 400 error: "The specified Lang can not be empty." Explicitly pass a language code.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The maximum number of entries to return in this request.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token. You do not need to specify this parameter for the first request or if no more results exist. If more results exist, set this parameter to the NextToken value returned by the previous API call.
	//
	// example:
	//
	// AAAAAUqcj6VO4E3ECWIrFczs****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The list of product IDs.
	ProductIdsShrink *string `json:"ProductIds,omitempty" xml:"ProductIds,omitempty"`
	// The product name.
	//
	// example:
	//
	// sas
	ProductName *string `json:"ProductName,omitempty" xml:"ProductName,omitempty"`
	// The product type. Valid values:
	//
	// - preset
	//
	// - custom
	//
	// example:
	//
	// preset
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// The region where the threat analysis data management center is located. Specify the management center based on the region of your assets. Valid values:
	//
	// - cn-hangzhou: The assets are located in the Chinese mainland or Hong Kong (China).
	//
	// - ap-southeast-1: The assets are located outside China.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The user ID of the member to which the administrator switches the view.
	//
	// example:
	//
	// 173326*******
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// The vendor ID.
	//
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
