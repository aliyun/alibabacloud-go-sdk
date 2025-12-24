// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVendorApiListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiName(v string) *DescribeVendorApiListRequest
	GetApiName() *string
	SetKeyWord(v string) *DescribeVendorApiListRequest
	GetKeyWord() *string
	SetPageNumber(v int32) *DescribeVendorApiListRequest
	GetPageNumber() *int32
	SetPageSize(v int64) *DescribeVendorApiListRequest
	GetPageSize() *int64
	SetProductCode(v string) *DescribeVendorApiListRequest
	GetProductCode() *string
	SetVendorCode(v string) *DescribeVendorApiListRequest
	GetVendorCode() *string
}

type DescribeVendorApiListRequest struct {
	// The name of the Alibaba Cloud product interface, supporting fuzzy search.
	//
	// example:
	//
	// AddAssetCleanConfig
	ApiName *string `json:"ApiName,omitempty" xml:"ApiName,omitempty"`
	// Keyword.
	//
	// example:
	//
	// Create
	KeyWord *string `json:"KeyWord,omitempty" xml:"KeyWord,omitempty"`
	// The current page number for pagination. The default value is 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of items per page when displaying paginated results. The default is 10.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Vendor\\"s product identifier:
	//
	// - **waf**: Web Application Firewall.
	//
	// - **cfw**: Cloud Firewall.
	//
	// example:
	//
	// waf
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// Vendor code:
	//
	// - **Tencent**: Tencent.
	//
	// - **HUAWEICLOUD**: Huawei.
	//
	// - **Azure**: Microsoft Azure.
	//
	// - **AWS**: Amazon Web Services.
	//
	// - **Chaitin**: Chaitin.
	//
	// example:
	//
	// Azure
	VendorCode *string `json:"VendorCode,omitempty" xml:"VendorCode,omitempty"`
}

func (s DescribeVendorApiListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVendorApiListRequest) GoString() string {
	return s.String()
}

func (s *DescribeVendorApiListRequest) GetApiName() *string {
	return s.ApiName
}

func (s *DescribeVendorApiListRequest) GetKeyWord() *string {
	return s.KeyWord
}

func (s *DescribeVendorApiListRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeVendorApiListRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeVendorApiListRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *DescribeVendorApiListRequest) GetVendorCode() *string {
	return s.VendorCode
}

func (s *DescribeVendorApiListRequest) SetApiName(v string) *DescribeVendorApiListRequest {
	s.ApiName = &v
	return s
}

func (s *DescribeVendorApiListRequest) SetKeyWord(v string) *DescribeVendorApiListRequest {
	s.KeyWord = &v
	return s
}

func (s *DescribeVendorApiListRequest) SetPageNumber(v int32) *DescribeVendorApiListRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeVendorApiListRequest) SetPageSize(v int64) *DescribeVendorApiListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeVendorApiListRequest) SetProductCode(v string) *DescribeVendorApiListRequest {
	s.ProductCode = &v
	return s
}

func (s *DescribeVendorApiListRequest) SetVendorCode(v string) *DescribeVendorApiListRequest {
	s.VendorCode = &v
	return s
}

func (s *DescribeVendorApiListRequest) Validate() error {
	return dara.Validate(s)
}
