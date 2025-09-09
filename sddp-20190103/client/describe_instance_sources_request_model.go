// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceSourcesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuditStatus(v int32) *DescribeInstanceSourcesRequest
	GetAuditStatus() *int32
	SetAuthStatus(v int32) *DescribeInstanceSourcesRequest
	GetAuthStatus() *int32
	SetCurrentPage(v int32) *DescribeInstanceSourcesRequest
	GetCurrentPage() *int32
	SetEngineType(v string) *DescribeInstanceSourcesRequest
	GetEngineType() *string
	SetFeatureType(v int32) *DescribeInstanceSourcesRequest
	GetFeatureType() *int32
	SetInstanceId(v string) *DescribeInstanceSourcesRequest
	GetInstanceId() *string
	SetLang(v string) *DescribeInstanceSourcesRequest
	GetLang() *string
	SetPageSize(v int32) *DescribeInstanceSourcesRequest
	GetPageSize() *int32
	SetProductCode(v string) *DescribeInstanceSourcesRequest
	GetProductCode() *string
	SetProductId(v int64) *DescribeInstanceSourcesRequest
	GetProductId() *int64
	SetSearchKey(v string) *DescribeInstanceSourcesRequest
	GetSearchKey() *string
	SetSearchType(v string) *DescribeInstanceSourcesRequest
	GetSearchType() *string
	SetServiceRegionId(v string) *DescribeInstanceSourcesRequest
	GetServiceRegionId() *string
}

type DescribeInstanceSourcesRequest struct {
	// Specifies whether to enable the security audit feature. Valid values:
	//
	// 	- **1**: yes
	//
	// 	- **0**: no
	//
	// example:
	//
	// 1
	AuditStatus *int32 `json:"AuditStatus,omitempty" xml:"AuditStatus,omitempty"`
	// Specifies whether DSC is authorized to access the data asset.
	//
	// 	- **0**: no
	//
	// 	- **1**: yes
	//
	// example:
	//
	// 0
	AuthStatus *int32 `json:"AuthStatus,omitempty" xml:"AuthStatus,omitempty"`
	// The number of the page to return. Default value: **1**.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The engine type. Valid values:
	//
	// 	- **MySQL**
	//
	// 	- **MariaDB**
	//
	// 	- **Oracle**
	//
	// 	- **PostgreSQL**
	//
	// 	- **SQLServer**
	//
	// example:
	//
	// MySQL
	EngineType *string `json:"EngineType,omitempty" xml:"EngineType,omitempty"`
	// This parameter is deprecated.
	//
	// example:
	//
	// 2
	FeatureType *int32 `json:"FeatureType,omitempty" xml:"FeatureType,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// instance-demo-****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The language of the content within the request and response. Valid values:
	//
	// 	- **zh_cn**: Simplified Chinese (default)
	//
	// 	- **en_us**: English
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
	// The name of the service to which the data asset to query belongs. Valid values: **MaxCompute, OSS, ADS, OTS, and RDS**.
	//
	// example:
	//
	// MaxCompute
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The ID of the service to which the asset belongs. Valid values:
	//
	// 	- **1**: MaxCompute
	//
	// 	- **2**: Object Storage Service (OSS)
	//
	// 	- **3**: AnalyticDB for MySQL
	//
	// 	- **4**: Tablestore (OTS)
	//
	// 	- **5**: ApsaraDB RDS
	//
	// 	- **6**: self-managed databases
	//
	// example:
	//
	// 1
	ProductId *int64 `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	// The content based on which a fuzzy search is performed.
	//
	// example:
	//
	// 1
	SearchKey *string `json:"SearchKey,omitempty" xml:"SearchKey,omitempty"`
	// The data asset type based on which a fuzzy search is performed.
	//
	// 	- **InstanceId**: the ID of the instance.
	//
	// 	- **InstanceName**: the name of the instance.
	//
	// 	- **DatabaseName**: the name of the database.
	//
	// example:
	//
	// InstanceId
	SearchType *string `json:"SearchType,omitempty" xml:"SearchType,omitempty"`
	// The region in which the data asset resides. For more information, see [Supported regions](https://help.aliyun.com/document_detail/214257.html).
	//
	// example:
	//
	// cn-hangzhou
	ServiceRegionId *string `json:"ServiceRegionId,omitempty" xml:"ServiceRegionId,omitempty"`
}

func (s DescribeInstanceSourcesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceSourcesRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceSourcesRequest) GetAuditStatus() *int32 {
	return s.AuditStatus
}

func (s *DescribeInstanceSourcesRequest) GetAuthStatus() *int32 {
	return s.AuthStatus
}

func (s *DescribeInstanceSourcesRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeInstanceSourcesRequest) GetEngineType() *string {
	return s.EngineType
}

func (s *DescribeInstanceSourcesRequest) GetFeatureType() *int32 {
	return s.FeatureType
}

func (s *DescribeInstanceSourcesRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeInstanceSourcesRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeInstanceSourcesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeInstanceSourcesRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *DescribeInstanceSourcesRequest) GetProductId() *int64 {
	return s.ProductId
}

func (s *DescribeInstanceSourcesRequest) GetSearchKey() *string {
	return s.SearchKey
}

func (s *DescribeInstanceSourcesRequest) GetSearchType() *string {
	return s.SearchType
}

func (s *DescribeInstanceSourcesRequest) GetServiceRegionId() *string {
	return s.ServiceRegionId
}

func (s *DescribeInstanceSourcesRequest) SetAuditStatus(v int32) *DescribeInstanceSourcesRequest {
	s.AuditStatus = &v
	return s
}

func (s *DescribeInstanceSourcesRequest) SetAuthStatus(v int32) *DescribeInstanceSourcesRequest {
	s.AuthStatus = &v
	return s
}

func (s *DescribeInstanceSourcesRequest) SetCurrentPage(v int32) *DescribeInstanceSourcesRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeInstanceSourcesRequest) SetEngineType(v string) *DescribeInstanceSourcesRequest {
	s.EngineType = &v
	return s
}

func (s *DescribeInstanceSourcesRequest) SetFeatureType(v int32) *DescribeInstanceSourcesRequest {
	s.FeatureType = &v
	return s
}

func (s *DescribeInstanceSourcesRequest) SetInstanceId(v string) *DescribeInstanceSourcesRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstanceSourcesRequest) SetLang(v string) *DescribeInstanceSourcesRequest {
	s.Lang = &v
	return s
}

func (s *DescribeInstanceSourcesRequest) SetPageSize(v int32) *DescribeInstanceSourcesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeInstanceSourcesRequest) SetProductCode(v string) *DescribeInstanceSourcesRequest {
	s.ProductCode = &v
	return s
}

func (s *DescribeInstanceSourcesRequest) SetProductId(v int64) *DescribeInstanceSourcesRequest {
	s.ProductId = &v
	return s
}

func (s *DescribeInstanceSourcesRequest) SetSearchKey(v string) *DescribeInstanceSourcesRequest {
	s.SearchKey = &v
	return s
}

func (s *DescribeInstanceSourcesRequest) SetSearchType(v string) *DescribeInstanceSourcesRequest {
	s.SearchType = &v
	return s
}

func (s *DescribeInstanceSourcesRequest) SetServiceRegionId(v string) *DescribeInstanceSourcesRequest {
	s.ServiceRegionId = &v
	return s
}

func (s *DescribeInstanceSourcesRequest) Validate() error {
	return dara.Validate(s)
}
