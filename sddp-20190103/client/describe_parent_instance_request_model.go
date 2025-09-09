// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeParentInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthStatus(v int32) *DescribeParentInstanceRequest
	GetAuthStatus() *int32
	SetCheckStatus(v int32) *DescribeParentInstanceRequest
	GetCheckStatus() *int32
	SetClusterStatus(v string) *DescribeParentInstanceRequest
	GetClusterStatus() *string
	SetCurrentPage(v int32) *DescribeParentInstanceRequest
	GetCurrentPage() *int32
	SetDbName(v string) *DescribeParentInstanceRequest
	GetDbName() *string
	SetEngineType(v string) *DescribeParentInstanceRequest
	GetEngineType() *string
	SetInstanceId(v string) *DescribeParentInstanceRequest
	GetInstanceId() *string
	SetLang(v string) *DescribeParentInstanceRequest
	GetLang() *string
	SetMemberAccount(v int64) *DescribeParentInstanceRequest
	GetMemberAccount() *int64
	SetPageSize(v int32) *DescribeParentInstanceRequest
	GetPageSize() *int32
	SetResourceType(v int64) *DescribeParentInstanceRequest
	GetResourceType() *int64
	SetServiceRegionId(v string) *DescribeParentInstanceRequest
	GetServiceRegionId() *string
}

type DescribeParentInstanceRequest struct {
	// Authorization status of the data asset instance.
	//
	// - **0**: Unauthorized
	//
	// - **1**: Authorized
	//
	// example:
	//
	// 0
	AuthStatus *int32 `json:"AuthStatus,omitempty" xml:"AuthStatus,omitempty"`
	// Connection status of the instance or the database under the instance. Values:
	//
	// - **-3**: Database not created
	//
	// - **-2**: Released
	//
	// - **-1**: Not connected
	//
	// - **2**: Connectivity test in progress
	//
	// - **3**: Connected
	//
	// - **4**: Connection failed
	//
	// example:
	//
	// 3
	CheckStatus *int32 `json:"CheckStatus,omitempty" xml:"CheckStatus,omitempty"`
	// Instance status.
	//
	// - **Running**: Running
	//
	// - **Released**: Released
	//
	// - **DatabaseNotCreated**: Database not created
	//
	// example:
	//
	// Running
	ClusterStatus *string `json:"ClusterStatus,omitempty" xml:"ClusterStatus,omitempty"`
	// When performing a paginated query, set the current page number. Default value: **1**.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// Database name.
	//
	// example:
	//
	// db_**t
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// Engine type. Values:
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
	// The instance ID to which the data in the data asset table belongs.
	//
	// example:
	//
	// rm-*******xx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Language type for request and response messages. Values:
	//
	// - **zh_cn**: Default, Simplified Chinese
	//
	// - **en_us**: English (US)
	//
	// example:
	//
	// zh_cn
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Member account ID.
	//
	// example:
	//
	// **********8103
	MemberAccount *int64 `json:"MemberAccount,omitempty" xml:"MemberAccount,omitempty"`
	// When performing a paginated query, set the number of rows per page. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The product type. Valid values:
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
	// - **6**: SelfDB
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
	ResourceType *int64 `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The region where the asset is located. Values:
	//
	// - **cn-beijing**: China (Beijing)
	//
	// - **cn-zhangjiakou**: China (Zhangjiakou)
	//
	// - **cn-huhehaote**: China (Hohhot)
	//
	// - **cn-hangzhou**: China (Hangzhou)
	//
	// - **cn-shanghai**: China (Shanghai)
	//
	// - **cn-shenzhen**: China (Shenzhen)
	//
	// - **cn-hongkong**:  China (Hong Kong)
	//
	// example:
	//
	// cn-shanghai
	ServiceRegionId *string `json:"ServiceRegionId,omitempty" xml:"ServiceRegionId,omitempty"`
}

func (s DescribeParentInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeParentInstanceRequest) GoString() string {
	return s.String()
}

func (s *DescribeParentInstanceRequest) GetAuthStatus() *int32 {
	return s.AuthStatus
}

func (s *DescribeParentInstanceRequest) GetCheckStatus() *int32 {
	return s.CheckStatus
}

func (s *DescribeParentInstanceRequest) GetClusterStatus() *string {
	return s.ClusterStatus
}

func (s *DescribeParentInstanceRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeParentInstanceRequest) GetDbName() *string {
	return s.DbName
}

func (s *DescribeParentInstanceRequest) GetEngineType() *string {
	return s.EngineType
}

func (s *DescribeParentInstanceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeParentInstanceRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeParentInstanceRequest) GetMemberAccount() *int64 {
	return s.MemberAccount
}

func (s *DescribeParentInstanceRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeParentInstanceRequest) GetResourceType() *int64 {
	return s.ResourceType
}

func (s *DescribeParentInstanceRequest) GetServiceRegionId() *string {
	return s.ServiceRegionId
}

func (s *DescribeParentInstanceRequest) SetAuthStatus(v int32) *DescribeParentInstanceRequest {
	s.AuthStatus = &v
	return s
}

func (s *DescribeParentInstanceRequest) SetCheckStatus(v int32) *DescribeParentInstanceRequest {
	s.CheckStatus = &v
	return s
}

func (s *DescribeParentInstanceRequest) SetClusterStatus(v string) *DescribeParentInstanceRequest {
	s.ClusterStatus = &v
	return s
}

func (s *DescribeParentInstanceRequest) SetCurrentPage(v int32) *DescribeParentInstanceRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeParentInstanceRequest) SetDbName(v string) *DescribeParentInstanceRequest {
	s.DbName = &v
	return s
}

func (s *DescribeParentInstanceRequest) SetEngineType(v string) *DescribeParentInstanceRequest {
	s.EngineType = &v
	return s
}

func (s *DescribeParentInstanceRequest) SetInstanceId(v string) *DescribeParentInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeParentInstanceRequest) SetLang(v string) *DescribeParentInstanceRequest {
	s.Lang = &v
	return s
}

func (s *DescribeParentInstanceRequest) SetMemberAccount(v int64) *DescribeParentInstanceRequest {
	s.MemberAccount = &v
	return s
}

func (s *DescribeParentInstanceRequest) SetPageSize(v int32) *DescribeParentInstanceRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeParentInstanceRequest) SetResourceType(v int64) *DescribeParentInstanceRequest {
	s.ResourceType = &v
	return s
}

func (s *DescribeParentInstanceRequest) SetServiceRegionId(v string) *DescribeParentInstanceRequest {
	s.ServiceRegionId = &v
	return s
}

func (s *DescribeParentInstanceRequest) Validate() error {
	return dara.Validate(s)
}
