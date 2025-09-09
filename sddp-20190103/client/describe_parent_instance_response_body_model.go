// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeParentInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *DescribeParentInstanceResponseBody
	GetCurrentPage() *int32
	SetItems(v []*DescribeParentInstanceResponseBodyItems) *DescribeParentInstanceResponseBody
	GetItems() []*DescribeParentInstanceResponseBodyItems
	SetPageSize(v int32) *DescribeParentInstanceResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeParentInstanceResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeParentInstanceResponseBody
	GetTotalCount() *int32
}

type DescribeParentInstanceResponseBody struct {
	// When performing a paginated query, set the current page number. Default value: **1**.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The assets.
	Items []*DescribeParentInstanceResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// When performing a paginated query, set the maximum number of data asset instances displayed per page. Default value: **10**.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Request ID of the result.
	//
	// example:
	//
	// ACEF9334-BB50-525D-8CF3-6CF504E4D1B3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Total number of data items in the result.
	//
	// example:
	//
	// 3
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeParentInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeParentInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeParentInstanceResponseBody) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeParentInstanceResponseBody) GetItems() []*DescribeParentInstanceResponseBodyItems {
	return s.Items
}

func (s *DescribeParentInstanceResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeParentInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeParentInstanceResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeParentInstanceResponseBody) SetCurrentPage(v int32) *DescribeParentInstanceResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *DescribeParentInstanceResponseBody) SetItems(v []*DescribeParentInstanceResponseBodyItems) *DescribeParentInstanceResponseBody {
	s.Items = v
	return s
}

func (s *DescribeParentInstanceResponseBody) SetPageSize(v int32) *DescribeParentInstanceResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeParentInstanceResponseBody) SetRequestId(v string) *DescribeParentInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeParentInstanceResponseBody) SetTotalCount(v int32) *DescribeParentInstanceResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeParentInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeParentInstanceResponseBodyItems struct {
	// Audit authorization status. The values are as follows:
	//
	// - **1**: Authorized
	//
	// - **0**: Unauthorized
	//
	// example:
	//
	// 1
	AuditStatus *int32 `json:"AuditStatus,omitempty" xml:"AuditStatus,omitempty"`
	// Authorization status of the data asset instance.
	//
	// - **0**: Unauthorized
	//
	// - **1**: Authorized
	//
	// example:
	//
	// 1
	AuthStatus *int32 `json:"AuthStatus,omitempty" xml:"AuthStatus,omitempty"`
	// Instance authorization time, in milliseconds.
	//
	// example:
	//
	// 1719882941000
	AuthTime *int64 `json:"AuthTime,omitempty" xml:"AuthTime,omitempty"`
	// Instance status.
	//
	// example:
	//
	// Running
	ClusterStatus *string `json:"ClusterStatus,omitempty" xml:"ClusterStatus,omitempty"`
	// Connection node type, valid only for MongoDB assets.
	//
	// example:
	//
	// Primary
	ConnectNode *string `json:"ConnectNode,omitempty" xml:"ConnectNode,omitempty"`
	// Number of databases under the instance.
	//
	// example:
	//
	// 3
	DbNum *string `json:"DbNum,omitempty" xml:"DbNum,omitempty"`
	// The engine type. Valid values:
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
	// Description of the instance.
	//
	// example:
	//
	// instance description
	InstanceDescription *string `json:"InstanceDescription,omitempty" xml:"InstanceDescription,omitempty"`
	// Instance ID.
	//
	// example:
	//
	// rm-*******t2vz
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Instance space size, valid only for OSS assets. Unit: bytes.
	//
	// example:
	//
	// 409600
	InstanceSize *int64 `json:"InstanceSize,omitempty" xml:"InstanceSize,omitempty"`
	// Region name. The values are as follows:
	//
	// - **China (Hangzhou)**
	//
	// - **China (Shanghai)**
	//
	// - **China (Beijing)**
	//
	// - **China (Zhangjiakou)**
	//
	// - **China (Shenzhen)**
	//
	// - **China (Guangzhou)**
	//
	// - **China (Hong Kong)**
	//
	// - **Singapore**
	//
	// - **US (Silicon Valley)**
	//
	// example:
	//
	// China (Hangzhou)
	LocalName *string `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	// Member account ID.
	//
	// example:
	//
	// **********8103
	MemberAccount *int64 `json:"MemberAccount,omitempty" xml:"MemberAccount,omitempty"`
	// Identifier for the authorized asset. For structured data, it is identified by `instanceID.databaseName`.
	//
	// example:
	//
	// rm-******xxx.**st
	ParentId *string `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
	// The region in which the asset resides.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Asset type name. The values are as follows:
	//
	// - **MaxCompute**
	//
	// - **OSS**
	//
	// - **ADB-MYSQL**
	//
	// - **TableStore**
	//
	// - **RDS**
	//
	// - **SelfDB**
	//
	// - **PolarDB-X**
	//
	// - **PolarDB**
	//
	// - **ADB-PG**
	//
	// - **OceanBase**
	//
	// - **MongoDB**
	//
	// - **Redis**
	//
	// example:
	//
	// RDS
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// Supported connection nodes, separated by commas.
	//
	// example:
	//
	// Primary,Secondary
	SupportConnectNodes *string `json:"SupportConnectNodes,omitempty" xml:"SupportConnectNodes,omitempty"`
	// Tenant ID, valid only for OceanBase assets.
	//
	// example:
	//
	// HB***-zh_CN
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// Tenant name, valid only for OceanBase assets.
	//
	// example:
	//
	// user1
	TenantName *string `json:"TenantName,omitempty" xml:"TenantName,omitempty"`
	// Number of unconnected databases under the instance.
	//
	// example:
	//
	// 1
	UnConnectDbCount *string `json:"UnConnectDbCount,omitempty" xml:"UnConnectDbCount,omitempty"`
	// Reason for not supporting one-click authorization.
	//
	// example:
	//
	// engine type not support
	UnSupportOneClickAuthReason *string `json:"UnSupportOneClickAuthReason,omitempty" xml:"UnSupportOneClickAuthReason,omitempty"`
}

func (s DescribeParentInstanceResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeParentInstanceResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeParentInstanceResponseBodyItems) GetAuditStatus() *int32 {
	return s.AuditStatus
}

func (s *DescribeParentInstanceResponseBodyItems) GetAuthStatus() *int32 {
	return s.AuthStatus
}

func (s *DescribeParentInstanceResponseBodyItems) GetAuthTime() *int64 {
	return s.AuthTime
}

func (s *DescribeParentInstanceResponseBodyItems) GetClusterStatus() *string {
	return s.ClusterStatus
}

func (s *DescribeParentInstanceResponseBodyItems) GetConnectNode() *string {
	return s.ConnectNode
}

func (s *DescribeParentInstanceResponseBodyItems) GetDbNum() *string {
	return s.DbNum
}

func (s *DescribeParentInstanceResponseBodyItems) GetEngineType() *string {
	return s.EngineType
}

func (s *DescribeParentInstanceResponseBodyItems) GetInstanceDescription() *string {
	return s.InstanceDescription
}

func (s *DescribeParentInstanceResponseBodyItems) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeParentInstanceResponseBodyItems) GetInstanceSize() *int64 {
	return s.InstanceSize
}

func (s *DescribeParentInstanceResponseBodyItems) GetLocalName() *string {
	return s.LocalName
}

func (s *DescribeParentInstanceResponseBodyItems) GetMemberAccount() *int64 {
	return s.MemberAccount
}

func (s *DescribeParentInstanceResponseBodyItems) GetParentId() *string {
	return s.ParentId
}

func (s *DescribeParentInstanceResponseBodyItems) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeParentInstanceResponseBodyItems) GetResourceType() *string {
	return s.ResourceType
}

func (s *DescribeParentInstanceResponseBodyItems) GetSupportConnectNodes() *string {
	return s.SupportConnectNodes
}

func (s *DescribeParentInstanceResponseBodyItems) GetTenantId() *string {
	return s.TenantId
}

func (s *DescribeParentInstanceResponseBodyItems) GetTenantName() *string {
	return s.TenantName
}

func (s *DescribeParentInstanceResponseBodyItems) GetUnConnectDbCount() *string {
	return s.UnConnectDbCount
}

func (s *DescribeParentInstanceResponseBodyItems) GetUnSupportOneClickAuthReason() *string {
	return s.UnSupportOneClickAuthReason
}

func (s *DescribeParentInstanceResponseBodyItems) SetAuditStatus(v int32) *DescribeParentInstanceResponseBodyItems {
	s.AuditStatus = &v
	return s
}

func (s *DescribeParentInstanceResponseBodyItems) SetAuthStatus(v int32) *DescribeParentInstanceResponseBodyItems {
	s.AuthStatus = &v
	return s
}

func (s *DescribeParentInstanceResponseBodyItems) SetAuthTime(v int64) *DescribeParentInstanceResponseBodyItems {
	s.AuthTime = &v
	return s
}

func (s *DescribeParentInstanceResponseBodyItems) SetClusterStatus(v string) *DescribeParentInstanceResponseBodyItems {
	s.ClusterStatus = &v
	return s
}

func (s *DescribeParentInstanceResponseBodyItems) SetConnectNode(v string) *DescribeParentInstanceResponseBodyItems {
	s.ConnectNode = &v
	return s
}

func (s *DescribeParentInstanceResponseBodyItems) SetDbNum(v string) *DescribeParentInstanceResponseBodyItems {
	s.DbNum = &v
	return s
}

func (s *DescribeParentInstanceResponseBodyItems) SetEngineType(v string) *DescribeParentInstanceResponseBodyItems {
	s.EngineType = &v
	return s
}

func (s *DescribeParentInstanceResponseBodyItems) SetInstanceDescription(v string) *DescribeParentInstanceResponseBodyItems {
	s.InstanceDescription = &v
	return s
}

func (s *DescribeParentInstanceResponseBodyItems) SetInstanceId(v string) *DescribeParentInstanceResponseBodyItems {
	s.InstanceId = &v
	return s
}

func (s *DescribeParentInstanceResponseBodyItems) SetInstanceSize(v int64) *DescribeParentInstanceResponseBodyItems {
	s.InstanceSize = &v
	return s
}

func (s *DescribeParentInstanceResponseBodyItems) SetLocalName(v string) *DescribeParentInstanceResponseBodyItems {
	s.LocalName = &v
	return s
}

func (s *DescribeParentInstanceResponseBodyItems) SetMemberAccount(v int64) *DescribeParentInstanceResponseBodyItems {
	s.MemberAccount = &v
	return s
}

func (s *DescribeParentInstanceResponseBodyItems) SetParentId(v string) *DescribeParentInstanceResponseBodyItems {
	s.ParentId = &v
	return s
}

func (s *DescribeParentInstanceResponseBodyItems) SetRegionId(v string) *DescribeParentInstanceResponseBodyItems {
	s.RegionId = &v
	return s
}

func (s *DescribeParentInstanceResponseBodyItems) SetResourceType(v string) *DescribeParentInstanceResponseBodyItems {
	s.ResourceType = &v
	return s
}

func (s *DescribeParentInstanceResponseBodyItems) SetSupportConnectNodes(v string) *DescribeParentInstanceResponseBodyItems {
	s.SupportConnectNodes = &v
	return s
}

func (s *DescribeParentInstanceResponseBodyItems) SetTenantId(v string) *DescribeParentInstanceResponseBodyItems {
	s.TenantId = &v
	return s
}

func (s *DescribeParentInstanceResponseBodyItems) SetTenantName(v string) *DescribeParentInstanceResponseBodyItems {
	s.TenantName = &v
	return s
}

func (s *DescribeParentInstanceResponseBodyItems) SetUnConnectDbCount(v string) *DescribeParentInstanceResponseBodyItems {
	s.UnConnectDbCount = &v
	return s
}

func (s *DescribeParentInstanceResponseBodyItems) SetUnSupportOneClickAuthReason(v string) *DescribeParentInstanceResponseBodyItems {
	s.UnSupportOneClickAuthReason = &v
	return s
}

func (s *DescribeParentInstanceResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
