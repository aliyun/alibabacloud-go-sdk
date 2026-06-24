// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDataObjectColumnDetailV2Request interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *DescribeDataObjectColumnDetailV2Request
	GetCurrentPage() *int32
	SetFeatureType(v int32) *DescribeDataObjectColumnDetailV2Request
	GetFeatureType() *int32
	SetId(v string) *DescribeDataObjectColumnDetailV2Request
	GetId() *string
	SetLang(v string) *DescribeDataObjectColumnDetailV2Request
	GetLang() *string
	SetPageSize(v int32) *DescribeDataObjectColumnDetailV2Request
	GetPageSize() *int32
	SetProductId(v int64) *DescribeDataObjectColumnDetailV2Request
	GetProductId() *int64
	SetTemplateId(v int64) *DescribeDataObjectColumnDetailV2Request
	GetTemplateId() *int64
}

type DescribeDataObjectColumnDetailV2Request struct {
	// The page number. Default value: **1**.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// This parameter is deprecated.
	//
	// example:
	//
	// 2
	FeatureType *int32 `json:"FeatureType,omitempty" xml:"FeatureType,omitempty"`
	// The unique ID of the data object to query.
	//
	// > You can call the [DescribeDataObjects](https://help.aliyun.com/document_detail/2399253.html) operation to obtain the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 13456723343
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The language of the content within the request and response. Default value: **zh_cn**. Valid values:
	//
	// - **zh_cn**: Chinese.
	//
	// - **en_us**: English.
	//
	// example:
	//
	// zh_cn
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The number of entries to return on each page. Default value: **10**.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the product to which the data object belongs. Valid values:
	//
	// - **1**: MaxCompute
	//
	// - **2**: OSS
	//
	// - **3**: ADB-MYSQL
	//
	// - **4**: Tablestore
	//
	// - **5**: RDS
	//
	// - **6**: SELF_DB
	//
	// - **7**: PolarDB-X
	//
	// - **8**: PolarDB
	//
	// - **9**: ADB-PG
	//
	// - **10**: OceanBase
	//
	// - **11**: MongoDB
	//
	// - **25**: Redis
	//
	// example:
	//
	// 5
	ProductId *int64 `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	// The ID of the industry-specific template.
	//
	// > You can call the [DescribeDataObjects](https://help.aliyun.com/document_detail/2399253.html) operation to obtain the ID of the industry-specific template.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s DescribeDataObjectColumnDetailV2Request) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataObjectColumnDetailV2Request) GoString() string {
	return s.String()
}

func (s *DescribeDataObjectColumnDetailV2Request) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeDataObjectColumnDetailV2Request) GetFeatureType() *int32 {
	return s.FeatureType
}

func (s *DescribeDataObjectColumnDetailV2Request) GetId() *string {
	return s.Id
}

func (s *DescribeDataObjectColumnDetailV2Request) GetLang() *string {
	return s.Lang
}

func (s *DescribeDataObjectColumnDetailV2Request) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDataObjectColumnDetailV2Request) GetProductId() *int64 {
	return s.ProductId
}

func (s *DescribeDataObjectColumnDetailV2Request) GetTemplateId() *int64 {
	return s.TemplateId
}

func (s *DescribeDataObjectColumnDetailV2Request) SetCurrentPage(v int32) *DescribeDataObjectColumnDetailV2Request {
	s.CurrentPage = &v
	return s
}

func (s *DescribeDataObjectColumnDetailV2Request) SetFeatureType(v int32) *DescribeDataObjectColumnDetailV2Request {
	s.FeatureType = &v
	return s
}

func (s *DescribeDataObjectColumnDetailV2Request) SetId(v string) *DescribeDataObjectColumnDetailV2Request {
	s.Id = &v
	return s
}

func (s *DescribeDataObjectColumnDetailV2Request) SetLang(v string) *DescribeDataObjectColumnDetailV2Request {
	s.Lang = &v
	return s
}

func (s *DescribeDataObjectColumnDetailV2Request) SetPageSize(v int32) *DescribeDataObjectColumnDetailV2Request {
	s.PageSize = &v
	return s
}

func (s *DescribeDataObjectColumnDetailV2Request) SetProductId(v int64) *DescribeDataObjectColumnDetailV2Request {
	s.ProductId = &v
	return s
}

func (s *DescribeDataObjectColumnDetailV2Request) SetTemplateId(v int64) *DescribeDataObjectColumnDetailV2Request {
	s.TemplateId = &v
	return s
}

func (s *DescribeDataObjectColumnDetailV2Request) Validate() error {
	return dara.Validate(s)
}
