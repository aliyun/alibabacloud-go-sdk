// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCloudVendorAccountAKListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthIds(v string) *DescribeCloudVendorAccountAKListRequest
	GetAuthIds() *string
	SetCurrentPage(v int32) *DescribeCloudVendorAccountAKListRequest
	GetCurrentPage() *int32
	SetLang(v string) *DescribeCloudVendorAccountAKListRequest
	GetLang() *string
	SetPageSize(v int32) *DescribeCloudVendorAccountAKListRequest
	GetPageSize() *int32
	SetStatus(v int32) *DescribeCloudVendorAccountAKListRequest
	GetStatus() *int32
	SetSubAccountName(v string) *DescribeCloudVendorAccountAKListRequest
	GetSubAccountName() *string
	SetVendor(v string) *DescribeCloudVendorAccountAKListRequest
	GetVendor() *string
	SetVendorAuthAlias(v string) *DescribeCloudVendorAccountAKListRequest
	GetVendorAuthAlias() *string
}

type DescribeCloudVendorAccountAKListRequest struct {
	// The unique ID of the AccessKey pair.
	//
	// example:
	//
	// 2624
	AuthIds *string `json:"AuthIds,omitempty" xml:"AuthIds,omitempty"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The language of the content within the request and response. Default value: **zh**. Valid values:
	//
	// 	- **zh**: Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The number of entries per page. Default value: 20.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The status of the AccessKey pair. Valid values:
	//
	// 	- **0**: enabled
	//
	// 	- **1**: disabled
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The username of the sub-account of the cloud service provider to which the AccessKey pair belongs.
	//
	// example:
	//
	// AlibabaCloud_***
	SubAccountName *string `json:"SubAccountName,omitempty" xml:"SubAccountName,omitempty"`
	Vendor         *string `json:"Vendor,omitempty" xml:"Vendor,omitempty"`
	// The name of the AccessKey pair.
	//
	// example:
	//
	// test
	VendorAuthAlias *string `json:"VendorAuthAlias,omitempty" xml:"VendorAuthAlias,omitempty"`
}

func (s DescribeCloudVendorAccountAKListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudVendorAccountAKListRequest) GoString() string {
	return s.String()
}

func (s *DescribeCloudVendorAccountAKListRequest) GetAuthIds() *string {
	return s.AuthIds
}

func (s *DescribeCloudVendorAccountAKListRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeCloudVendorAccountAKListRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeCloudVendorAccountAKListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeCloudVendorAccountAKListRequest) GetStatus() *int32 {
	return s.Status
}

func (s *DescribeCloudVendorAccountAKListRequest) GetSubAccountName() *string {
	return s.SubAccountName
}

func (s *DescribeCloudVendorAccountAKListRequest) GetVendor() *string {
	return s.Vendor
}

func (s *DescribeCloudVendorAccountAKListRequest) GetVendorAuthAlias() *string {
	return s.VendorAuthAlias
}

func (s *DescribeCloudVendorAccountAKListRequest) SetAuthIds(v string) *DescribeCloudVendorAccountAKListRequest {
	s.AuthIds = &v
	return s
}

func (s *DescribeCloudVendorAccountAKListRequest) SetCurrentPage(v int32) *DescribeCloudVendorAccountAKListRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeCloudVendorAccountAKListRequest) SetLang(v string) *DescribeCloudVendorAccountAKListRequest {
	s.Lang = &v
	return s
}

func (s *DescribeCloudVendorAccountAKListRequest) SetPageSize(v int32) *DescribeCloudVendorAccountAKListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeCloudVendorAccountAKListRequest) SetStatus(v int32) *DescribeCloudVendorAccountAKListRequest {
	s.Status = &v
	return s
}

func (s *DescribeCloudVendorAccountAKListRequest) SetSubAccountName(v string) *DescribeCloudVendorAccountAKListRequest {
	s.SubAccountName = &v
	return s
}

func (s *DescribeCloudVendorAccountAKListRequest) SetVendor(v string) *DescribeCloudVendorAccountAKListRequest {
	s.Vendor = &v
	return s
}

func (s *DescribeCloudVendorAccountAKListRequest) SetVendorAuthAlias(v string) *DescribeCloudVendorAccountAKListRequest {
	s.VendorAuthAlias = &v
	return s
}

func (s *DescribeCloudVendorAccountAKListRequest) Validate() error {
	return dara.Validate(s)
}
