// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGlobalDatabaseNetworkResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConnections(v []*DescribeGlobalDatabaseNetworkResponseBodyConnections) *DescribeGlobalDatabaseNetworkResponseBody
	GetConnections() []*DescribeGlobalDatabaseNetworkResponseBodyConnections
	SetCreateTime(v string) *DescribeGlobalDatabaseNetworkResponseBody
	GetCreateTime() *string
	SetDBClusterId(v string) *DescribeGlobalDatabaseNetworkResponseBody
	GetDBClusterId() *string
	SetDBClusters(v []*DescribeGlobalDatabaseNetworkResponseBodyDBClusters) *DescribeGlobalDatabaseNetworkResponseBody
	GetDBClusters() []*DescribeGlobalDatabaseNetworkResponseBodyDBClusters
	SetDBType(v string) *DescribeGlobalDatabaseNetworkResponseBody
	GetDBType() *string
	SetDBVersion(v string) *DescribeGlobalDatabaseNetworkResponseBody
	GetDBVersion() *string
	SetGDNDescription(v string) *DescribeGlobalDatabaseNetworkResponseBody
	GetGDNDescription() *string
	SetGDNId(v string) *DescribeGlobalDatabaseNetworkResponseBody
	GetGDNId() *string
	SetGDNStatus(v string) *DescribeGlobalDatabaseNetworkResponseBody
	GetGDNStatus() *string
	SetGlobalDomainName(v string) *DescribeGlobalDatabaseNetworkResponseBody
	GetGlobalDomainName() *string
	SetLabels(v *DescribeGlobalDatabaseNetworkResponseBodyLabels) *DescribeGlobalDatabaseNetworkResponseBody
	GetLabels() *DescribeGlobalDatabaseNetworkResponseBodyLabels
	SetRequestId(v string) *DescribeGlobalDatabaseNetworkResponseBody
	GetRequestId() *string
	SetResourceGroupId(v string) *DescribeGlobalDatabaseNetworkResponseBody
	GetResourceGroupId() *string
}

type DescribeGlobalDatabaseNetworkResponseBody struct {
	// The information about the connection to the cluster.
	Connections []*DescribeGlobalDatabaseNetworkResponseBodyConnections `json:"Connections,omitempty" xml:"Connections,omitempty" type:"Repeated"`
	// The time at which the GDN was created.
	//
	// example:
	//
	// 2020-02-24T11:57:54Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the cluster.
	//
	// example:
	//
	// pc-bp1s826a1up******
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The clusters in the GDN.
	DBClusters []*DescribeGlobalDatabaseNetworkResponseBodyDBClusters `json:"DBClusters,omitempty" xml:"DBClusters,omitempty" type:"Repeated"`
	// The type of the database engine. Only MySQL is supported.
	//
	// example:
	//
	// MySQL
	DBType *string `json:"DBType,omitempty" xml:"DBType,omitempty"`
	// The version of the database engine. Only version 8.0 is supported.
	//
	// example:
	//
	// 8.0
	DBVersion *string `json:"DBVersion,omitempty" xml:"DBVersion,omitempty"`
	// The description of the GDN. The description must meet the following requirements:
	//
	// 	- It cannot start with `http://` or `https://`.
	//
	// 	- It must start with a letter.
	//
	// 	- It can contain letters, digits, underscores (_), and hyphens (-).
	//
	// 	- It must be 2 to 126 characters in length.
	//
	// example:
	//
	// GDN-fortest
	GDNDescription *string `json:"GDNDescription,omitempty" xml:"GDNDescription,omitempty"`
	// The ID of the GDN.
	//
	// example:
	//
	// gdn-bp1fttxsrmv*****
	GDNId *string `json:"GDNId,omitempty" xml:"GDNId,omitempty"`
	// The status of the GDN. Valid values:
	//
	// 	- **Creating**: The GDN is being created.
	//
	// 	- **active**: The GDN is running.
	//
	// 	- **deleting**: The GDN is being deleted.
	//
	// 	- **locked**: The GDN is locked. If the GDN is locked, you cannot perform operations on clusters in the GDN.
	//
	// 	- **removing_member**: The secondary cluster is being removed from the GDN.
	//
	// example:
	//
	// active
	GDNStatus *string `json:"GDNStatus,omitempty" xml:"GDNStatus,omitempty"`
	// The global domain name.
	//
	// example:
	//
	// [gdnid].gdn.rds.aliyuncs.com
	GlobalDomainName *string                                          `json:"GlobalDomainName,omitempty" xml:"GlobalDomainName,omitempty"`
	Labels           *DescribeGlobalDatabaseNetworkResponseBodyLabels `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 67F2E75F-AE67-4FB2-821F-A81237EACD15
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-************
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s DescribeGlobalDatabaseNetworkResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeGlobalDatabaseNetworkResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeGlobalDatabaseNetworkResponseBody) GetConnections() []*DescribeGlobalDatabaseNetworkResponseBodyConnections {
	return s.Connections
}

func (s *DescribeGlobalDatabaseNetworkResponseBody) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeGlobalDatabaseNetworkResponseBody) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeGlobalDatabaseNetworkResponseBody) GetDBClusters() []*DescribeGlobalDatabaseNetworkResponseBodyDBClusters {
	return s.DBClusters
}

func (s *DescribeGlobalDatabaseNetworkResponseBody) GetDBType() *string {
	return s.DBType
}

func (s *DescribeGlobalDatabaseNetworkResponseBody) GetDBVersion() *string {
	return s.DBVersion
}

func (s *DescribeGlobalDatabaseNetworkResponseBody) GetGDNDescription() *string {
	return s.GDNDescription
}

func (s *DescribeGlobalDatabaseNetworkResponseBody) GetGDNId() *string {
	return s.GDNId
}

func (s *DescribeGlobalDatabaseNetworkResponseBody) GetGDNStatus() *string {
	return s.GDNStatus
}

func (s *DescribeGlobalDatabaseNetworkResponseBody) GetGlobalDomainName() *string {
	return s.GlobalDomainName
}

func (s *DescribeGlobalDatabaseNetworkResponseBody) GetLabels() *DescribeGlobalDatabaseNetworkResponseBodyLabels {
	return s.Labels
}

func (s *DescribeGlobalDatabaseNetworkResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeGlobalDatabaseNetworkResponseBody) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeGlobalDatabaseNetworkResponseBody) SetConnections(v []*DescribeGlobalDatabaseNetworkResponseBodyConnections) *DescribeGlobalDatabaseNetworkResponseBody {
	s.Connections = v
	return s
}

func (s *DescribeGlobalDatabaseNetworkResponseBody) SetCreateTime(v string) *DescribeGlobalDatabaseNetworkResponseBody {
	s.CreateTime = &v
	return s
}

func (s *DescribeGlobalDatabaseNetworkResponseBody) SetDBClusterId(v string) *DescribeGlobalDatabaseNetworkResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *DescribeGlobalDatabaseNetworkResponseBody) SetDBClusters(v []*DescribeGlobalDatabaseNetworkResponseBodyDBClusters) *DescribeGlobalDatabaseNetworkResponseBody {
	s.DBClusters = v
	return s
}

func (s *DescribeGlobalDatabaseNetworkResponseBody) SetDBType(v string) *DescribeGlobalDatabaseNetworkResponseBody {
	s.DBType = &v
	return s
}

func (s *DescribeGlobalDatabaseNetworkResponseBody) SetDBVersion(v string) *DescribeGlobalDatabaseNetworkResponseBody {
	s.DBVersion = &v
	return s
}

func (s *DescribeGlobalDatabaseNetworkResponseBody) SetGDNDescription(v string) *DescribeGlobalDatabaseNetworkResponseBody {
	s.GDNDescription = &v
	return s
}

func (s *DescribeGlobalDatabaseNetworkResponseBody) SetGDNId(v string) *DescribeGlobalDatabaseNetworkResponseBody {
	s.GDNId = &v
	return s
}

func (s *DescribeGlobalDatabaseNetworkResponseBody) SetGDNStatus(v string) *DescribeGlobalDatabaseNetworkResponseBody {
	s.GDNStatus = &v
	return s
}

func (s *DescribeGlobalDatabaseNetworkResponseBody) SetGlobalDomainName(v string) *DescribeGlobalDatabaseNetworkResponseBody {
	s.GlobalDomainName = &v
	return s
}

func (s *DescribeGlobalDatabaseNetworkResponseBody) SetLabels(v *DescribeGlobalDatabaseNetworkResponseBodyLabels) *DescribeGlobalDatabaseNetworkResponseBody {
	s.Labels = v
	return s
}

func (s *DescribeGlobalDatabaseNetworkResponseBody) SetRequestId(v string) *DescribeGlobalDatabaseNetworkResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeGlobalDatabaseNetworkResponseBody) SetResourceGroupId(v string) *DescribeGlobalDatabaseNetworkResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeGlobalDatabaseNetworkResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeGlobalDatabaseNetworkResponseBodyConnections struct {
	// The endpoint URL of the database service.
	//
	// example:
	//
	// abc.polardb.rds.aliyuncs.com
	ConnectionString *string `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	// The network type for the database connection.
	//
	// example:
	//
	// Private
	NetType *string `json:"NetType,omitempty" xml:"NetType,omitempty"`
	// The port number for the database connection.
	//
	// example:
	//
	// 3306
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
}

func (s DescribeGlobalDatabaseNetworkResponseBodyConnections) String() string {
	return dara.Prettify(s)
}

func (s DescribeGlobalDatabaseNetworkResponseBodyConnections) GoString() string {
	return s.String()
}

func (s *DescribeGlobalDatabaseNetworkResponseBodyConnections) GetConnectionString() *string {
	return s.ConnectionString
}

func (s *DescribeGlobalDatabaseNetworkResponseBodyConnections) GetNetType() *string {
	return s.NetType
}

func (s *DescribeGlobalDatabaseNetworkResponseBodyConnections) GetPort() *string {
	return s.Port
}

func (s *DescribeGlobalDatabaseNetworkResponseBodyConnections) SetConnectionString(v string) *DescribeGlobalDatabaseNetworkResponseBodyConnections {
	s.ConnectionString = &v
	return s
}

func (s *DescribeGlobalDatabaseNetworkResponseBodyConnections) SetNetType(v string) *DescribeGlobalDatabaseNetworkResponseBodyConnections {
	s.NetType = &v
	return s
}

func (s *DescribeGlobalDatabaseNetworkResponseBodyConnections) SetPort(v string) *DescribeGlobalDatabaseNetworkResponseBodyConnections {
	s.Port = &v
	return s
}

func (s *DescribeGlobalDatabaseNetworkResponseBodyConnections) Validate() error {
	return dara.Validate(s)
}

type DescribeGlobalDatabaseNetworkResponseBodyDBClusters struct {
	// The edition of the cluster. Valid values:
	//
	// Normal: Cluster Edition Basic: Single Node Edition Archive: X-Engine Edition NormalMultimaster: Multi-master Cluster Edition SENormal: Standard Edition
	//
	// >
	//
	// 	- PolarDB for PostgreSQL clusters that run the PostgreSQL 11 database engine do not support Single Node Edition.
	//
	// 	- PolarDB for MySQL 8.0 and 5.7 clusters, and PolarDB for PostgreSQL clusters that run the PostgreSQL 14 database engine support Standard Edition.
	//
	// 	- PolarDB for MySQL 8.0 clusters support X-Engine Edition and Multi-master Cluster Edition.
	//
	// example:
	//
	// Normal
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The description of the cluster.
	//
	// example:
	//
	// test
	DBClusterDescription *string `json:"DBClusterDescription,omitempty" xml:"DBClusterDescription,omitempty"`
	// The ID of the cluster.
	//
	// example:
	//
	// pc-wz9fb5nn44u1d****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The status of the cluster. For more information, see [Cluster status table](https://help.aliyun.com/document_detail/99286.html).
	//
	// example:
	//
	// Running
	DBClusterStatus *string `json:"DBClusterStatus,omitempty" xml:"DBClusterStatus,omitempty"`
	// The node specifications of the cluster.
	//
	// example:
	//
	// polar.mysql.x4.large
	DBNodeClass *string `json:"DBNodeClass,omitempty" xml:"DBNodeClass,omitempty"`
	// The nodes of the cluster.
	DBNodes []*DescribeGlobalDatabaseNetworkResponseBodyDBClustersDBNodes `json:"DBNodes,omitempty" xml:"DBNodes,omitempty" type:"Repeated"`
	// The database engine type of the cluster. Only MySQL is supported.
	//
	// example:
	//
	// MySQL
	DBType *string `json:"DBType,omitempty" xml:"DBType,omitempty"`
	// The version of the database engine. Only version 8.0 is supported.
	//
	// example:
	//
	// 8.0
	DBVersion *string `json:"DBVersion,omitempty" xml:"DBVersion,omitempty"`
	// The expiration time of the cluster.
	//
	// >  A specific value is returned only for subscription (**Prepaid**) clusters. No value is returned for pay-as-you-go (**Postpaid**) clusters.
	//
	// example:
	//
	// 2020-11-14T16:00:00Z
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// Indicates whether the cluster has expired. Valid values:
	//
	// 	- **true*	- (default)
	//
	// 	- **false**
	//
	// >  This parameter is returned only for subscription (**Prepaid**) clusters.
	//
	// example:
	//
	// false
	Expired *string `json:"Expired,omitempty" xml:"Expired,omitempty"`
	// The billing method of the cluster. Valid values:
	//
	// 	- **Postpaid**: pay-as-you-go
	//
	// 	- **Prepaid**: subscription
	//
	// example:
	//
	// Prepaid
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// The region ID of the cluster.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The cross-region data replication latency between the primary cluster and secondary clusters. Unit: seconds.
	//
	// example:
	//
	// 1
	ReplicaLag *string `json:"ReplicaLag,omitempty" xml:"ReplicaLag,omitempty"`
	// The role of the cluster. Valid values:
	//
	// 	- **Primary**: the primary cluster
	//
	// 	- **standby**: a secondary cluster
	//
	// >  A GDN consists of one primary cluster and up to four secondary clusters.
	//
	// example:
	//
	// primary
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// Indicates whether the cluster is a serverless cluster. The value is fixed at AgileServerless.
	//
	// >  This parameter is returned only for serverless clusters.
	//
	// example:
	//
	// AgileServerless
	ServerlessType *string `json:"ServerlessType,omitempty" xml:"ServerlessType,omitempty"`
	// The storage usage of the cluster. Unit: bytes.
	//
	// example:
	//
	// 3012558848
	StorageUsed *string `json:"StorageUsed,omitempty" xml:"StorageUsed,omitempty"`
}

func (s DescribeGlobalDatabaseNetworkResponseBodyDBClusters) String() string {
	return dara.Prettify(s)
}

func (s DescribeGlobalDatabaseNetworkResponseBodyDBClusters) GoString() string {
	return s.String()
}

func (s *DescribeGlobalDatabaseNetworkResponseBodyDBClusters) GetCategory() *string {
	return s.Category
}

func (s *DescribeGlobalDatabaseNetworkResponseBodyDBClusters) GetDBClusterDescription() *string {
	return s.DBClusterDescription
}

func (s *DescribeGlobalDatabaseNetworkResponseBodyDBClusters) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeGlobalDatabaseNetworkResponseBodyDBClusters) GetDBClusterStatus() *string {
	return s.DBClusterStatus
}

func (s *DescribeGlobalDatabaseNetworkResponseBodyDBClusters) GetDBNodeClass() *string {
	return s.DBNodeClass
}

func (s *DescribeGlobalDatabaseNetworkResponseBodyDBClusters) GetDBNodes() []*DescribeGlobalDatabaseNetworkResponseBodyDBClustersDBNodes {
	return s.DBNodes
}

func (s *DescribeGlobalDatabaseNetworkResponseBodyDBClusters) GetDBType() *string {
	return s.DBType
}

func (s *DescribeGlobalDatabaseNetworkResponseBodyDBClusters) GetDBVersion() *string {
	return s.DBVersion
}

func (s *DescribeGlobalDatabaseNetworkResponseBodyDBClusters) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *DescribeGlobalDatabaseNetworkResponseBodyDBClusters) GetExpired() *string {
	return s.Expired
}

func (s *DescribeGlobalDatabaseNetworkResponseBodyDBClusters) GetPayType() *string {
	return s.PayType
}

func (s *DescribeGlobalDatabaseNetworkResponseBodyDBClusters) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeGlobalDatabaseNetworkResponseBodyDBClusters) GetReplicaLag() *string {
	return s.ReplicaLag
}

func (s *DescribeGlobalDatabaseNetworkResponseBodyDBClusters) GetRole() *string {
	return s.Role
}

func (s *DescribeGlobalDatabaseNetworkResponseBodyDBClusters) GetServerlessType() *string {
	return s.ServerlessType
}

func (s *DescribeGlobalDatabaseNetworkResponseBodyDBClusters) GetStorageUsed() *string {
	return s.StorageUsed
}

func (s *DescribeGlobalDatabaseNetworkResponseBodyDBClusters) SetCategory(v string) *DescribeGlobalDatabaseNetworkResponseBodyDBClusters {
	s.Category = &v
	return s
}

func (s *DescribeGlobalDatabaseNetworkResponseBodyDBClusters) SetDBClusterDescription(v string) *DescribeGlobalDatabaseNetworkResponseBodyDBClusters {
	s.DBClusterDescription = &v
	return s
}

func (s *DescribeGlobalDatabaseNetworkResponseBodyDBClusters) SetDBClusterId(v string) *DescribeGlobalDatabaseNetworkResponseBodyDBClusters {
	s.DBClusterId = &v
	return s
}

func (s *DescribeGlobalDatabaseNetworkResponseBodyDBClusters) SetDBClusterStatus(v string) *DescribeGlobalDatabaseNetworkResponseBodyDBClusters {
	s.DBClusterStatus = &v
	return s
}

func (s *DescribeGlobalDatabaseNetworkResponseBodyDBClusters) SetDBNodeClass(v string) *DescribeGlobalDatabaseNetworkResponseBodyDBClusters {
	s.DBNodeClass = &v
	return s
}

func (s *DescribeGlobalDatabaseNetworkResponseBodyDBClusters) SetDBNodes(v []*DescribeGlobalDatabaseNetworkResponseBodyDBClustersDBNodes) *DescribeGlobalDatabaseNetworkResponseBodyDBClusters {
	s.DBNodes = v
	return s
}

func (s *DescribeGlobalDatabaseNetworkResponseBodyDBClusters) SetDBType(v string) *DescribeGlobalDatabaseNetworkResponseBodyDBClusters {
	s.DBType = &v
	return s
}

func (s *DescribeGlobalDatabaseNetworkResponseBodyDBClusters) SetDBVersion(v string) *DescribeGlobalDatabaseNetworkResponseBodyDBClusters {
	s.DBVersion = &v
	return s
}

func (s *DescribeGlobalDatabaseNetworkResponseBodyDBClusters) SetExpireTime(v string) *DescribeGlobalDatabaseNetworkResponseBodyDBClusters {
	s.ExpireTime = &v
	return s
}

func (s *DescribeGlobalDatabaseNetworkResponseBodyDBClusters) SetExpired(v string) *DescribeGlobalDatabaseNetworkResponseBodyDBClusters {
	s.Expired = &v
	return s
}

func (s *DescribeGlobalDatabaseNetworkResponseBodyDBClusters) SetPayType(v string) *DescribeGlobalDatabaseNetworkResponseBodyDBClusters {
	s.PayType = &v
	return s
}

func (s *DescribeGlobalDatabaseNetworkResponseBodyDBClusters) SetRegionId(v string) *DescribeGlobalDatabaseNetworkResponseBodyDBClusters {
	s.RegionId = &v
	return s
}

func (s *DescribeGlobalDatabaseNetworkResponseBodyDBClusters) SetReplicaLag(v string) *DescribeGlobalDatabaseNetworkResponseBodyDBClusters {
	s.ReplicaLag = &v
	return s
}

func (s *DescribeGlobalDatabaseNetworkResponseBodyDBClusters) SetRole(v string) *DescribeGlobalDatabaseNetworkResponseBodyDBClusters {
	s.Role = &v
	return s
}

func (s *DescribeGlobalDatabaseNetworkResponseBodyDBClusters) SetServerlessType(v string) *DescribeGlobalDatabaseNetworkResponseBodyDBClusters {
	s.ServerlessType = &v
	return s
}

func (s *DescribeGlobalDatabaseNetworkResponseBodyDBClusters) SetStorageUsed(v string) *DescribeGlobalDatabaseNetworkResponseBodyDBClusters {
	s.StorageUsed = &v
	return s
}

func (s *DescribeGlobalDatabaseNetworkResponseBodyDBClusters) Validate() error {
	return dara.Validate(s)
}

type DescribeGlobalDatabaseNetworkResponseBodyDBClustersDBNodes struct {
	// The time when the node was created.
	//
	// example:
	//
	// 2020-03-23T21:35:43Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The specifications of the node.
	//
	// example:
	//
	// polar.mysql.x4.large
	DBNodeClass *string `json:"DBNodeClass,omitempty" xml:"DBNodeClass,omitempty"`
	// The node ID.
	//
	// example:
	//
	// pi-****************
	DBNodeId *string `json:"DBNodeId,omitempty" xml:"DBNodeId,omitempty"`
	// The role of the node. Valid values:
	//
	// 	- **Writer**: the primary node
	//
	// 	- **Reader**: a read-only node
	//
	// example:
	//
	// Reader
	DBNodeRole *string `json:"DBNodeRole,omitempty" xml:"DBNodeRole,omitempty"`
	// The status of the node. Valid values:
	//
	// 	- **Creating**: The node is being created.
	//
	// 	- **Running**: The node is running.
	//
	// 	- **Deleting**: The node is being deleted.
	//
	// 	- **Rebooting**: The node is restarting.
	//
	// 	- **ClassChanging**: The specifications of the node are being changed.
	//
	// 	- **NetAddressCreating**: The network connection is being created.
	//
	// 	- **NetAddressDeleting**: The network connection is being deleted.
	//
	// 	- **NetAddressModifying**: The network connection is being modified.
	//
	// 	- **MinorVersionUpgrading**: The minor version of the node is being updated.
	//
	// 	- **Maintaining**: The node is being maintained.
	//
	// 	- **Switching**: A failover is being performed.
	//
	// example:
	//
	// Running
	DBNodeStatus *string `json:"DBNodeStatus,omitempty" xml:"DBNodeStatus,omitempty"`
	// The failover priority. Each node is assigned a failover priority. The failover priority determines which node is selected as the primary node when a failover occurs. A larger value indicates a higher priority. Valid values: 1 to 15.
	//
	// example:
	//
	// 1
	FailoverPriority *int32 `json:"FailoverPriority,omitempty" xml:"FailoverPriority,omitempty"`
	// The maximum number of concurrent connections.
	//
	// example:
	//
	// 8000
	MaxConnections *int32 `json:"MaxConnections,omitempty" xml:"MaxConnections,omitempty"`
	// The maximum input/output operations per second (IOPS).
	//
	// example:
	//
	// 32000
	MaxIOPS *int32 `json:"MaxIOPS,omitempty" xml:"MaxIOPS,omitempty"`
	// The zone ID of the node.
	//
	// example:
	//
	// cn-hangzhou-i
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeGlobalDatabaseNetworkResponseBodyDBClustersDBNodes) String() string {
	return dara.Prettify(s)
}

func (s DescribeGlobalDatabaseNetworkResponseBodyDBClustersDBNodes) GoString() string {
	return s.String()
}

func (s *DescribeGlobalDatabaseNetworkResponseBodyDBClustersDBNodes) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeGlobalDatabaseNetworkResponseBodyDBClustersDBNodes) GetDBNodeClass() *string {
	return s.DBNodeClass
}

func (s *DescribeGlobalDatabaseNetworkResponseBodyDBClustersDBNodes) GetDBNodeId() *string {
	return s.DBNodeId
}

func (s *DescribeGlobalDatabaseNetworkResponseBodyDBClustersDBNodes) GetDBNodeRole() *string {
	return s.DBNodeRole
}

func (s *DescribeGlobalDatabaseNetworkResponseBodyDBClustersDBNodes) GetDBNodeStatus() *string {
	return s.DBNodeStatus
}

func (s *DescribeGlobalDatabaseNetworkResponseBodyDBClustersDBNodes) GetFailoverPriority() *int32 {
	return s.FailoverPriority
}

func (s *DescribeGlobalDatabaseNetworkResponseBodyDBClustersDBNodes) GetMaxConnections() *int32 {
	return s.MaxConnections
}

func (s *DescribeGlobalDatabaseNetworkResponseBodyDBClustersDBNodes) GetMaxIOPS() *int32 {
	return s.MaxIOPS
}

func (s *DescribeGlobalDatabaseNetworkResponseBodyDBClustersDBNodes) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeGlobalDatabaseNetworkResponseBodyDBClustersDBNodes) SetCreationTime(v string) *DescribeGlobalDatabaseNetworkResponseBodyDBClustersDBNodes {
	s.CreationTime = &v
	return s
}

func (s *DescribeGlobalDatabaseNetworkResponseBodyDBClustersDBNodes) SetDBNodeClass(v string) *DescribeGlobalDatabaseNetworkResponseBodyDBClustersDBNodes {
	s.DBNodeClass = &v
	return s
}

func (s *DescribeGlobalDatabaseNetworkResponseBodyDBClustersDBNodes) SetDBNodeId(v string) *DescribeGlobalDatabaseNetworkResponseBodyDBClustersDBNodes {
	s.DBNodeId = &v
	return s
}

func (s *DescribeGlobalDatabaseNetworkResponseBodyDBClustersDBNodes) SetDBNodeRole(v string) *DescribeGlobalDatabaseNetworkResponseBodyDBClustersDBNodes {
	s.DBNodeRole = &v
	return s
}

func (s *DescribeGlobalDatabaseNetworkResponseBodyDBClustersDBNodes) SetDBNodeStatus(v string) *DescribeGlobalDatabaseNetworkResponseBodyDBClustersDBNodes {
	s.DBNodeStatus = &v
	return s
}

func (s *DescribeGlobalDatabaseNetworkResponseBodyDBClustersDBNodes) SetFailoverPriority(v int32) *DescribeGlobalDatabaseNetworkResponseBodyDBClustersDBNodes {
	s.FailoverPriority = &v
	return s
}

func (s *DescribeGlobalDatabaseNetworkResponseBodyDBClustersDBNodes) SetMaxConnections(v int32) *DescribeGlobalDatabaseNetworkResponseBodyDBClustersDBNodes {
	s.MaxConnections = &v
	return s
}

func (s *DescribeGlobalDatabaseNetworkResponseBodyDBClustersDBNodes) SetMaxIOPS(v int32) *DescribeGlobalDatabaseNetworkResponseBodyDBClustersDBNodes {
	s.MaxIOPS = &v
	return s
}

func (s *DescribeGlobalDatabaseNetworkResponseBodyDBClustersDBNodes) SetZoneId(v string) *DescribeGlobalDatabaseNetworkResponseBodyDBClustersDBNodes {
	s.ZoneId = &v
	return s
}

func (s *DescribeGlobalDatabaseNetworkResponseBodyDBClustersDBNodes) Validate() error {
	return dara.Validate(s)
}

type DescribeGlobalDatabaseNetworkResponseBodyLabels struct {
	GDNVersion *string `json:"GDNVersion,omitempty" xml:"GDNVersion,omitempty"`
}

func (s DescribeGlobalDatabaseNetworkResponseBodyLabels) String() string {
	return dara.Prettify(s)
}

func (s DescribeGlobalDatabaseNetworkResponseBodyLabels) GoString() string {
	return s.String()
}

func (s *DescribeGlobalDatabaseNetworkResponseBodyLabels) GetGDNVersion() *string {
	return s.GDNVersion
}

func (s *DescribeGlobalDatabaseNetworkResponseBodyLabels) SetGDNVersion(v string) *DescribeGlobalDatabaseNetworkResponseBodyLabels {
	s.GDNVersion = &v
	return s
}

func (s *DescribeGlobalDatabaseNetworkResponseBodyLabels) Validate() error {
	return dara.Validate(s)
}
