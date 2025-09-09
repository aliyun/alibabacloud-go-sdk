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
	// When performing a paginated query, set the current page number. Default value: **1**.
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
	// Set the unique identifier ID of the data object to be queried.
	//
	// > You can obtain the identifier ID by calling [DescribeDataObjects](https://help.aliyun.com/document_detail/2399253.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// 13456723343
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The language type for the request and response messages, default is **zh_cn**. Values:
	//
	// - **zh_cn**: Chinese.
	//
	// - **en_us**: English.
	//
	// example:
	//
	// zh_cn
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// When performing a paginated query, set the maximum number of data asset instances to display per page. Default value: **10**.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID corresponding to the product name of the data object. Values:
	//
	// - **1**: MaxCompute
	//
	// - **2**: OSS
	//
	// - **3**: ADB-MYSQL
	//
	// - **4**: TableStore
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
	// Industry template ID.
	//
	// > You can obtain the industry template identifier ID by calling [DescribeDataObjects](https://help.aliyun.com/document_detail/2399253.html).
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
