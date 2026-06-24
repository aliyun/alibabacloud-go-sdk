// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDataObjectColumnDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *DescribeDataObjectColumnDetailRequest
	GetCurrentPage() *int32
	SetFeatureType(v int32) *DescribeDataObjectColumnDetailRequest
	GetFeatureType() *int32
	SetId(v int64) *DescribeDataObjectColumnDetailRequest
	GetId() *int64
	SetLang(v string) *DescribeDataObjectColumnDetailRequest
	GetLang() *string
	SetPageSize(v int32) *DescribeDataObjectColumnDetailRequest
	GetPageSize() *int32
	SetProductId(v int64) *DescribeDataObjectColumnDetailRequest
	GetProductId() *int64
	SetTemplateId(v int64) *DescribeDataObjectColumnDetailRequest
	GetTemplateId() *int64
}

type DescribeDataObjectColumnDetailRequest struct {
	// The page number to return. Default value: **1**.
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
	// The unique ID of the data object that you want to query.
	//
	// > Call the [DescribeDataObjects](https://help.aliyun.com/document_detail/2399253.html) operation to obtain the ID.
	//
	// example:
	//
	// 318248
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The language of the content within the request and response. Default value: **zh_cn**. Valid values:
	//
	// - **zh_cn**: Chinese
	//
	// - **en_us**: English
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
	// The ID of the service to which the data object belongs. Valid values:
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
	// > Call the [DescribeDataObjects](https://help.aliyun.com/document_detail/2399253.html) operation to obtain the ID of the industry-specific template.
	//
	// example:
	//
	// 1
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s DescribeDataObjectColumnDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataObjectColumnDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribeDataObjectColumnDetailRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeDataObjectColumnDetailRequest) GetFeatureType() *int32 {
	return s.FeatureType
}

func (s *DescribeDataObjectColumnDetailRequest) GetId() *int64 {
	return s.Id
}

func (s *DescribeDataObjectColumnDetailRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeDataObjectColumnDetailRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDataObjectColumnDetailRequest) GetProductId() *int64 {
	return s.ProductId
}

func (s *DescribeDataObjectColumnDetailRequest) GetTemplateId() *int64 {
	return s.TemplateId
}

func (s *DescribeDataObjectColumnDetailRequest) SetCurrentPage(v int32) *DescribeDataObjectColumnDetailRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeDataObjectColumnDetailRequest) SetFeatureType(v int32) *DescribeDataObjectColumnDetailRequest {
	s.FeatureType = &v
	return s
}

func (s *DescribeDataObjectColumnDetailRequest) SetId(v int64) *DescribeDataObjectColumnDetailRequest {
	s.Id = &v
	return s
}

func (s *DescribeDataObjectColumnDetailRequest) SetLang(v string) *DescribeDataObjectColumnDetailRequest {
	s.Lang = &v
	return s
}

func (s *DescribeDataObjectColumnDetailRequest) SetPageSize(v int32) *DescribeDataObjectColumnDetailRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDataObjectColumnDetailRequest) SetProductId(v int64) *DescribeDataObjectColumnDetailRequest {
	s.ProductId = &v
	return s
}

func (s *DescribeDataObjectColumnDetailRequest) SetTemplateId(v int64) *DescribeDataObjectColumnDetailRequest {
	s.TemplateId = &v
	return s
}

func (s *DescribeDataObjectColumnDetailRequest) Validate() error {
	return dara.Validate(s)
}
