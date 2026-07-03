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
	// The language of the response. Valid values:
	//
	// - **zh*	- (default): Chinese.
	//
	// - **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The maximum number of entries to return.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that is used to start the next query. You do not need to specify this parameter for the first query. For subsequent queries, set this parameter to the \\`NextToken\\` value that is returned from the previous API call.
	//
	// example:
	//
	// AAAAAUqcj6VO4E3ECWIrFczs****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The list of product IDs.
	ProductIds []*string `json:"ProductIds,omitempty" xml:"ProductIds,omitempty" type:"Repeated"`
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
	// The region of the Data Management center for threat analysis. Select the region for the Management Hub based on the region where your assets are located. Valid values:
	//
	// - cn-hangzhou: Your assets are in the Chinese mainland or China (Hong Kong).
	//
	// - ap-southeast-1: Your assets are in regions outside China.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The user ID of the member. An administrator can specify this parameter to switch to the perspective of this member.
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
