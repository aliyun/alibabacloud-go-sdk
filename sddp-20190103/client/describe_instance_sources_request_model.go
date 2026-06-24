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
	// The audit status. Valid values:
	//
	// - **1**: Auditing is enabled.
	//
	// - **0**: Auditing is disabled.
	//
	// example:
	//
	// 1
	AuditStatus *int32 `json:"AuditStatus,omitempty" xml:"AuditStatus,omitempty"`
	// The authorization status of the data asset instance.
	//
	// - **0**: Unauthorized.
	//
	// - **1**: Authorized.
	//
	// example:
	//
	// 0
	AuthStatus *int32 `json:"AuthStatus,omitempty" xml:"AuthStatus,omitempty"`
	// The page number. Default value: **1**.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The database engine type. Valid values:
	//
	// - **MySQL**
	//
	// - **MariaDB**
	//
	// - **Oracle**
	//
	// - **PostgreSQL**
	//
	// - **SQLServer**
	//
	// example:
	//
	// MySQL
	EngineType *string `json:"EngineType,omitempty" xml:"EngineType,omitempty"`
	// This parameter is deprecated.
	//
	// example:
	//
	// 1
	FeatureType *int32 `json:"FeatureType,omitempty" xml:"FeatureType,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// instance-demo-****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The language of the request and response. Valid values:
	//
	// - **zh_cn**: Simplified Chinese. This is the default value.
	//
	// - **en_us**: English.
	//
	// example:
	//
	// zh_cn
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The number of entries per page in a paginated query. Default value: **10**.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The name of the product to query. Valid values: MaxCompute, OSS, ADS, OTS, and RDS.
	//
	// example:
	//
	// MaxCompute
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The product type ID to query. Valid values:
	//
	// - **1**: MaxCompute
	//
	// - **2**: OSS
	//
	// - **3**: ADS
	//
	// - **4**: OTS
	//
	// - **5**: RDS
	//
	// - **6**: SELF_DB
	//
	// example:
	//
	// 1
	ProductId *int64 `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	// The keyword for the fuzzy search of data assets.
	//
	// example:
	//
	// 1
	SearchKey *string `json:"SearchKey,omitempty" xml:"SearchKey,omitempty"`
	// The type of the fuzzy search for data assets. Valid values:
	//
	// - **InstanceId**: The instance ID.
	//
	// - **InstanceName**: The instance name.
	//
	// - **DatabaseName**: The database name.
	//
	// example:
	//
	// InstanceId
	SearchType *string `json:"SearchType,omitempty" xml:"SearchType,omitempty"`
	// The region where the asset is located. For more information, see [Supported regions](https://help.aliyun.com/document_detail/214257.html).
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
