// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBClusterMigrationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetComment(v string) *DescribeDBClusterMigrationResponseBody
	GetComment() *string
	SetDBClusterEndpointList(v []*DescribeDBClusterMigrationResponseBodyDBClusterEndpointList) *DescribeDBClusterMigrationResponseBody
	GetDBClusterEndpointList() []*DescribeDBClusterMigrationResponseBodyDBClusterEndpointList
	SetDBClusterId(v string) *DescribeDBClusterMigrationResponseBody
	GetDBClusterId() *string
	SetDBClusterReadWriteMode(v string) *DescribeDBClusterMigrationResponseBody
	GetDBClusterReadWriteMode() *string
	SetDelayedSeconds(v int32) *DescribeDBClusterMigrationResponseBody
	GetDelayedSeconds() *int32
	SetDtsInstanceId(v string) *DescribeDBClusterMigrationResponseBody
	GetDtsInstanceId() *string
	SetExpiredTime(v string) *DescribeDBClusterMigrationResponseBody
	GetExpiredTime() *string
	SetMigrationStatus(v string) *DescribeDBClusterMigrationResponseBody
	GetMigrationStatus() *string
	SetRdsEndpointList(v []*DescribeDBClusterMigrationResponseBodyRdsEndpointList) *DescribeDBClusterMigrationResponseBody
	GetRdsEndpointList() []*DescribeDBClusterMigrationResponseBodyRdsEndpointList
	SetRdsReadWriteMode(v string) *DescribeDBClusterMigrationResponseBody
	GetRdsReadWriteMode() *string
	SetRequestId(v string) *DescribeDBClusterMigrationResponseBody
	GetRequestId() *string
	SetSourceRDSDBInstanceId(v string) *DescribeDBClusterMigrationResponseBody
	GetSourceRDSDBInstanceId() *string
	SetSrcDbType(v string) *DescribeDBClusterMigrationResponseBody
	GetSrcDbType() *string
	SetTopologies(v string) *DescribeDBClusterMigrationResponseBody
	GetTopologies() *string
}

type DescribeDBClusterMigrationResponseBody struct {
	// The comments on the migration exception. If no exception occurs during the migration, an empty value is returned.
	//
	// example:
	//
	// test
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The details of the PolarDB endpoints.
	DBClusterEndpointList []*DescribeDBClusterMigrationResponseBodyDBClusterEndpointList `json:"DBClusterEndpointList,omitempty" xml:"DBClusterEndpointList,omitempty" type:"Repeated"`
	// The cluster ID.
	//
	// example:
	//
	// pc-****************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The read/write mode of the cluster. Valid values:
	//
	// - **rw**: Read and write.
	//
	// - **ro**: Read-only.
	//
	// example:
	//
	// ro
	DBClusterReadWriteMode *string `json:"DBClusterReadWriteMode,omitempty" xml:"DBClusterReadWriteMode,omitempty"`
	// The replication delay between the ApsaraDB RDS instance and the PolarDB cluster, in seconds.
	//
	// example:
	//
	// 0
	DelayedSeconds *int32 `json:"DelayedSeconds,omitempty" xml:"DelayedSeconds,omitempty"`
	// The ID of the sync task.
	//
	// example:
	//
	// dts**********618bs
	DtsInstanceId *string `json:"DtsInstanceId,omitempty" xml:"DtsInstanceId,omitempty"`
	// The time when the replication relationship between the ApsaraDB RDS instance and the PolarDB cluster expires. The time is in the `YYYY-MM-DDThh:mm:ssZ` format and is displayed in UTC.
	//
	// example:
	//
	// 2020-06-17T01:56:36Z
	ExpiredTime *string `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	// The migration status of the PolarDB cluster. Valid values:
	//
	// - **NO_MIGRATION**: No migration task is created.
	//
	// - **RDS2POLARDB_CLONING**: Data is being cloned.
	//
	// - **RDS2POLARDB_SYNCING**: Data is being synchronized. In this state, the PolarDB cluster is read-only, and the ApsaraDB RDS instance is read-write.
	//
	// - **SWITCHING**: The database is being switched.
	//
	// - **POLARDB2RDS_SYNCING**: The database switch is complete. In this state, the PolarDB cluster is read-write, and the ApsaraDB RDS instance is read-only. Change the endpoints in your application.
	//
	// - **ROLLBACK**: The migration is being rolled back. After the rollback is complete, the migration status changes to **RDS2POLARDB_SYNCING**.
	//
	// - **CLOSING_MIGRATION**: The migration task is being shut down.
	//
	// example:
	//
	// RDS2POLARDB_SYNCING
	MigrationStatus *string `json:"MigrationStatus,omitempty" xml:"MigrationStatus,omitempty"`
	// The details of the ApsaraDB RDS endpoints.
	RdsEndpointList []*DescribeDBClusterMigrationResponseBodyRdsEndpointList `json:"RdsEndpointList,omitempty" xml:"RdsEndpointList,omitempty" type:"Repeated"`
	// The read/write mode of the source ApsaraDB RDS instance. Valid values:
	//
	// - **rw**: Read and write.
	//
	// - **ro**: Read-only.
	//
	// example:
	//
	// rw
	RdsReadWriteMode *string `json:"RdsReadWriteMode,omitempty" xml:"RdsReadWriteMode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// F2A9EFA7-915F-4572-8299-85A307******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the source ApsaraDB RDS instance.
	//
	// example:
	//
	// rm-************
	SourceRDSDBInstanceId *string `json:"SourceRDSDBInstanceId,omitempty" xml:"SourceRDSDBInstanceId,omitempty"`
	// The type of the source database. Valid values:
	//
	// - **PolarDBMySQL**: The source database for a major version upgrade of a PolarDB cluster.
	//
	// - **RDS**: The source database for migrating data from an ApsaraDB RDS instance to a PolarDB for MySQL cluster.
	//
	// example:
	//
	// PolarDBMySQL
	SrcDbType *string `json:"SrcDbType,omitempty" xml:"SrcDbType,omitempty"`
	// The data synchronization relationship. Valid values:
	//
	// - **RDS2POLARDB**: Data is synchronized from the ApsaraDB RDS instance to the PolarDB cluster.
	//
	// - **POLARDB2RDS**: Data is synchronized from the PolarDB cluster to the ApsaraDB RDS instance.
	//
	// example:
	//
	// RDS2POLARDB
	Topologies *string `json:"Topologies,omitempty" xml:"Topologies,omitempty"`
}

func (s DescribeDBClusterMigrationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterMigrationResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterMigrationResponseBody) GetComment() *string {
	return s.Comment
}

func (s *DescribeDBClusterMigrationResponseBody) GetDBClusterEndpointList() []*DescribeDBClusterMigrationResponseBodyDBClusterEndpointList {
	return s.DBClusterEndpointList
}

func (s *DescribeDBClusterMigrationResponseBody) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeDBClusterMigrationResponseBody) GetDBClusterReadWriteMode() *string {
	return s.DBClusterReadWriteMode
}

func (s *DescribeDBClusterMigrationResponseBody) GetDelayedSeconds() *int32 {
	return s.DelayedSeconds
}

func (s *DescribeDBClusterMigrationResponseBody) GetDtsInstanceId() *string {
	return s.DtsInstanceId
}

func (s *DescribeDBClusterMigrationResponseBody) GetExpiredTime() *string {
	return s.ExpiredTime
}

func (s *DescribeDBClusterMigrationResponseBody) GetMigrationStatus() *string {
	return s.MigrationStatus
}

func (s *DescribeDBClusterMigrationResponseBody) GetRdsEndpointList() []*DescribeDBClusterMigrationResponseBodyRdsEndpointList {
	return s.RdsEndpointList
}

func (s *DescribeDBClusterMigrationResponseBody) GetRdsReadWriteMode() *string {
	return s.RdsReadWriteMode
}

func (s *DescribeDBClusterMigrationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDBClusterMigrationResponseBody) GetSourceRDSDBInstanceId() *string {
	return s.SourceRDSDBInstanceId
}

func (s *DescribeDBClusterMigrationResponseBody) GetSrcDbType() *string {
	return s.SrcDbType
}

func (s *DescribeDBClusterMigrationResponseBody) GetTopologies() *string {
	return s.Topologies
}

func (s *DescribeDBClusterMigrationResponseBody) SetComment(v string) *DescribeDBClusterMigrationResponseBody {
	s.Comment = &v
	return s
}

func (s *DescribeDBClusterMigrationResponseBody) SetDBClusterEndpointList(v []*DescribeDBClusterMigrationResponseBodyDBClusterEndpointList) *DescribeDBClusterMigrationResponseBody {
	s.DBClusterEndpointList = v
	return s
}

func (s *DescribeDBClusterMigrationResponseBody) SetDBClusterId(v string) *DescribeDBClusterMigrationResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *DescribeDBClusterMigrationResponseBody) SetDBClusterReadWriteMode(v string) *DescribeDBClusterMigrationResponseBody {
	s.DBClusterReadWriteMode = &v
	return s
}

func (s *DescribeDBClusterMigrationResponseBody) SetDelayedSeconds(v int32) *DescribeDBClusterMigrationResponseBody {
	s.DelayedSeconds = &v
	return s
}

func (s *DescribeDBClusterMigrationResponseBody) SetDtsInstanceId(v string) *DescribeDBClusterMigrationResponseBody {
	s.DtsInstanceId = &v
	return s
}

func (s *DescribeDBClusterMigrationResponseBody) SetExpiredTime(v string) *DescribeDBClusterMigrationResponseBody {
	s.ExpiredTime = &v
	return s
}

func (s *DescribeDBClusterMigrationResponseBody) SetMigrationStatus(v string) *DescribeDBClusterMigrationResponseBody {
	s.MigrationStatus = &v
	return s
}

func (s *DescribeDBClusterMigrationResponseBody) SetRdsEndpointList(v []*DescribeDBClusterMigrationResponseBodyRdsEndpointList) *DescribeDBClusterMigrationResponseBody {
	s.RdsEndpointList = v
	return s
}

func (s *DescribeDBClusterMigrationResponseBody) SetRdsReadWriteMode(v string) *DescribeDBClusterMigrationResponseBody {
	s.RdsReadWriteMode = &v
	return s
}

func (s *DescribeDBClusterMigrationResponseBody) SetRequestId(v string) *DescribeDBClusterMigrationResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBClusterMigrationResponseBody) SetSourceRDSDBInstanceId(v string) *DescribeDBClusterMigrationResponseBody {
	s.SourceRDSDBInstanceId = &v
	return s
}

func (s *DescribeDBClusterMigrationResponseBody) SetSrcDbType(v string) *DescribeDBClusterMigrationResponseBody {
	s.SrcDbType = &v
	return s
}

func (s *DescribeDBClusterMigrationResponseBody) SetTopologies(v string) *DescribeDBClusterMigrationResponseBody {
	s.Topologies = &v
	return s
}

func (s *DescribeDBClusterMigrationResponseBody) Validate() error {
	if s.DBClusterEndpointList != nil {
		for _, item := range s.DBClusterEndpointList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.RdsEndpointList != nil {
		for _, item := range s.RdsEndpointList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDBClusterMigrationResponseBodyDBClusterEndpointList struct {
	// The details of the connection strings.
	AddressItems []*DescribeDBClusterMigrationResponseBodyDBClusterEndpointListAddressItems `json:"AddressItems,omitempty" xml:"AddressItems,omitempty" type:"Repeated"`
	// The endpoint ID.
	//
	// example:
	//
	// pe-***********
	DBEndpointId *string `json:"DBEndpointId,omitempty" xml:"DBEndpointId,omitempty"`
	// The type of the endpoint. Valid values:
	//
	// - **Cluster**: The default cluster endpoint.
	//
	// - **Primary**: The primary endpoint.
	//
	// - **Custom**: A custom cluster endpoint.
	//
	// example:
	//
	// Cluster
	EndpointType *string `json:"EndpointType,omitempty" xml:"EndpointType,omitempty"`
	// The read/write mode. Valid values:
	//
	// - ReadWrite: Read and write (automatic read/write splitting).
	//
	// - ReadOnly (Default): Read-only.
	//
	// example:
	//
	// ReadOnly
	ReadWriteMode *string `json:"ReadWriteMode,omitempty" xml:"ReadWriteMode,omitempty"`
}

func (s DescribeDBClusterMigrationResponseBodyDBClusterEndpointList) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterMigrationResponseBodyDBClusterEndpointList) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterMigrationResponseBodyDBClusterEndpointList) GetAddressItems() []*DescribeDBClusterMigrationResponseBodyDBClusterEndpointListAddressItems {
	return s.AddressItems
}

func (s *DescribeDBClusterMigrationResponseBodyDBClusterEndpointList) GetDBEndpointId() *string {
	return s.DBEndpointId
}

func (s *DescribeDBClusterMigrationResponseBodyDBClusterEndpointList) GetEndpointType() *string {
	return s.EndpointType
}

func (s *DescribeDBClusterMigrationResponseBodyDBClusterEndpointList) GetReadWriteMode() *string {
	return s.ReadWriteMode
}

func (s *DescribeDBClusterMigrationResponseBodyDBClusterEndpointList) SetAddressItems(v []*DescribeDBClusterMigrationResponseBodyDBClusterEndpointListAddressItems) *DescribeDBClusterMigrationResponseBodyDBClusterEndpointList {
	s.AddressItems = v
	return s
}

func (s *DescribeDBClusterMigrationResponseBodyDBClusterEndpointList) SetDBEndpointId(v string) *DescribeDBClusterMigrationResponseBodyDBClusterEndpointList {
	s.DBEndpointId = &v
	return s
}

func (s *DescribeDBClusterMigrationResponseBodyDBClusterEndpointList) SetEndpointType(v string) *DescribeDBClusterMigrationResponseBodyDBClusterEndpointList {
	s.EndpointType = &v
	return s
}

func (s *DescribeDBClusterMigrationResponseBodyDBClusterEndpointList) SetReadWriteMode(v string) *DescribeDBClusterMigrationResponseBodyDBClusterEndpointList {
	s.ReadWriteMode = &v
	return s
}

func (s *DescribeDBClusterMigrationResponseBodyDBClusterEndpointList) Validate() error {
	if s.AddressItems != nil {
		for _, item := range s.AddressItems {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDBClusterMigrationResponseBodyDBClusterEndpointListAddressItems struct {
	// The connection string.
	//
	// example:
	//
	// pc-**************.rwlb.rds.aliyuncs.com
	ConnectionString *string `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	// The IP address.
	//
	// example:
	//
	// 192.***.***.10
	IPAddress *string `json:"IPAddress,omitempty" xml:"IPAddress,omitempty"`
	// The network type of the endpoint. Valid values:
	//
	// - **Public**: An endpoint for the Internet.
	//
	// - **Private**: A private endpoint.
	//
	// - **Inner**: A private endpoint in a classic network.
	//
	// example:
	//
	// Private
	NetType *string `json:"NetType,omitempty" xml:"NetType,omitempty"`
	// The port.
	//
	// example:
	//
	// 3306
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// Indicates whether Secure Sockets Layer (SSL) encryption is enabled. Valid values:
	//
	// - **Enabled**: SSL encryption is enabled.
	//
	// - **Disabled**: SSL encryption is disabled.
	//
	// example:
	//
	// Enabled
	SSLEnabled *string `json:"SSLEnabled,omitempty" xml:"SSLEnabled,omitempty"`
	// The ID of the virtual private cloud (VPC).
	//
	// example:
	//
	// vpc-**********
	VPCId *string `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	// The ID of the virtual switch.
	//
	// example:
	//
	// vsw-**********
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
}

func (s DescribeDBClusterMigrationResponseBodyDBClusterEndpointListAddressItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterMigrationResponseBodyDBClusterEndpointListAddressItems) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterMigrationResponseBodyDBClusterEndpointListAddressItems) GetConnectionString() *string {
	return s.ConnectionString
}

func (s *DescribeDBClusterMigrationResponseBodyDBClusterEndpointListAddressItems) GetIPAddress() *string {
	return s.IPAddress
}

func (s *DescribeDBClusterMigrationResponseBodyDBClusterEndpointListAddressItems) GetNetType() *string {
	return s.NetType
}

func (s *DescribeDBClusterMigrationResponseBodyDBClusterEndpointListAddressItems) GetPort() *string {
	return s.Port
}

func (s *DescribeDBClusterMigrationResponseBodyDBClusterEndpointListAddressItems) GetSSLEnabled() *string {
	return s.SSLEnabled
}

func (s *DescribeDBClusterMigrationResponseBodyDBClusterEndpointListAddressItems) GetVPCId() *string {
	return s.VPCId
}

func (s *DescribeDBClusterMigrationResponseBodyDBClusterEndpointListAddressItems) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribeDBClusterMigrationResponseBodyDBClusterEndpointListAddressItems) SetConnectionString(v string) *DescribeDBClusterMigrationResponseBodyDBClusterEndpointListAddressItems {
	s.ConnectionString = &v
	return s
}

func (s *DescribeDBClusterMigrationResponseBodyDBClusterEndpointListAddressItems) SetIPAddress(v string) *DescribeDBClusterMigrationResponseBodyDBClusterEndpointListAddressItems {
	s.IPAddress = &v
	return s
}

func (s *DescribeDBClusterMigrationResponseBodyDBClusterEndpointListAddressItems) SetNetType(v string) *DescribeDBClusterMigrationResponseBodyDBClusterEndpointListAddressItems {
	s.NetType = &v
	return s
}

func (s *DescribeDBClusterMigrationResponseBodyDBClusterEndpointListAddressItems) SetPort(v string) *DescribeDBClusterMigrationResponseBodyDBClusterEndpointListAddressItems {
	s.Port = &v
	return s
}

func (s *DescribeDBClusterMigrationResponseBodyDBClusterEndpointListAddressItems) SetSSLEnabled(v string) *DescribeDBClusterMigrationResponseBodyDBClusterEndpointListAddressItems {
	s.SSLEnabled = &v
	return s
}

func (s *DescribeDBClusterMigrationResponseBodyDBClusterEndpointListAddressItems) SetVPCId(v string) *DescribeDBClusterMigrationResponseBodyDBClusterEndpointListAddressItems {
	s.VPCId = &v
	return s
}

func (s *DescribeDBClusterMigrationResponseBodyDBClusterEndpointListAddressItems) SetVSwitchId(v string) *DescribeDBClusterMigrationResponseBodyDBClusterEndpointListAddressItems {
	s.VSwitchId = &v
	return s
}

func (s *DescribeDBClusterMigrationResponseBodyDBClusterEndpointListAddressItems) Validate() error {
	return dara.Validate(s)
}

type DescribeDBClusterMigrationResponseBodyRdsEndpointList struct {
	// The details of the connection strings.
	AddressItems []*DescribeDBClusterMigrationResponseBodyRdsEndpointListAddressItems `json:"AddressItems,omitempty" xml:"AddressItems,omitempty" type:"Repeated"`
	// The instance type.
	//
	// example:
	//
	// ReadOnly
	//
	// Maxscale
	//
	// Primary
	CustinsType *string `json:"CustinsType,omitempty" xml:"CustinsType,omitempty"`
	// The endpoint ID.
	//
	// example:
	//
	// rm-************-normal
	DBEndpointId *string `json:"DBEndpointId,omitempty" xml:"DBEndpointId,omitempty"`
	// The type of the endpoint. Valid values:
	//
	// - **Normal**: A regular endpoint.
	//
	// - **ReadWriteSplitting**: A read/write splitting endpoint.
	//
	// example:
	//
	// Normal
	EndpointType *string `json:"EndpointType,omitempty" xml:"EndpointType,omitempty"`
}

func (s DescribeDBClusterMigrationResponseBodyRdsEndpointList) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterMigrationResponseBodyRdsEndpointList) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterMigrationResponseBodyRdsEndpointList) GetAddressItems() []*DescribeDBClusterMigrationResponseBodyRdsEndpointListAddressItems {
	return s.AddressItems
}

func (s *DescribeDBClusterMigrationResponseBodyRdsEndpointList) GetCustinsType() *string {
	return s.CustinsType
}

func (s *DescribeDBClusterMigrationResponseBodyRdsEndpointList) GetDBEndpointId() *string {
	return s.DBEndpointId
}

func (s *DescribeDBClusterMigrationResponseBodyRdsEndpointList) GetEndpointType() *string {
	return s.EndpointType
}

func (s *DescribeDBClusterMigrationResponseBodyRdsEndpointList) SetAddressItems(v []*DescribeDBClusterMigrationResponseBodyRdsEndpointListAddressItems) *DescribeDBClusterMigrationResponseBodyRdsEndpointList {
	s.AddressItems = v
	return s
}

func (s *DescribeDBClusterMigrationResponseBodyRdsEndpointList) SetCustinsType(v string) *DescribeDBClusterMigrationResponseBodyRdsEndpointList {
	s.CustinsType = &v
	return s
}

func (s *DescribeDBClusterMigrationResponseBodyRdsEndpointList) SetDBEndpointId(v string) *DescribeDBClusterMigrationResponseBodyRdsEndpointList {
	s.DBEndpointId = &v
	return s
}

func (s *DescribeDBClusterMigrationResponseBodyRdsEndpointList) SetEndpointType(v string) *DescribeDBClusterMigrationResponseBodyRdsEndpointList {
	s.EndpointType = &v
	return s
}

func (s *DescribeDBClusterMigrationResponseBodyRdsEndpointList) Validate() error {
	if s.AddressItems != nil {
		for _, item := range s.AddressItems {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDBClusterMigrationResponseBodyRdsEndpointListAddressItems struct {
	// The connection string.
	//
	// example:
	//
	// rm-***********.mysql.rds.aliyuncs.com
	ConnectionString *string `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	// The IP address.
	//
	// example:
	//
	// 172.***.***.173
	IPAddress *string `json:"IPAddress,omitempty" xml:"IPAddress,omitempty"`
	// The network type of the endpoint. Valid values:
	//
	// - **Public**: An endpoint for the Internet.
	//
	// - **Private**: A private endpoint.
	//
	// - **Inner**: A private endpoint in a classic network.
	//
	// example:
	//
	// Private
	NetType *string `json:"NetType,omitempty" xml:"NetType,omitempty"`
	// The port.
	//
	// example:
	//
	// 3306
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// Indicates whether SSL encryption is enabled. Valid values:
	//
	// - **Enabled**: SSL encryption is enabled.
	//
	// - **Disabled**: SSL encryption is disabled.
	//
	// example:
	//
	// Enabled
	SSLEnabled *string `json:"SSLEnabled,omitempty" xml:"SSLEnabled,omitempty"`
	// The ID of the VPC.
	//
	// example:
	//
	// vpc-************
	VPCId *string `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	// The ID of the vSwitch.
	//
	// example:
	//
	// vsw-**************
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
}

func (s DescribeDBClusterMigrationResponseBodyRdsEndpointListAddressItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterMigrationResponseBodyRdsEndpointListAddressItems) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterMigrationResponseBodyRdsEndpointListAddressItems) GetConnectionString() *string {
	return s.ConnectionString
}

func (s *DescribeDBClusterMigrationResponseBodyRdsEndpointListAddressItems) GetIPAddress() *string {
	return s.IPAddress
}

func (s *DescribeDBClusterMigrationResponseBodyRdsEndpointListAddressItems) GetNetType() *string {
	return s.NetType
}

func (s *DescribeDBClusterMigrationResponseBodyRdsEndpointListAddressItems) GetPort() *string {
	return s.Port
}

func (s *DescribeDBClusterMigrationResponseBodyRdsEndpointListAddressItems) GetSSLEnabled() *string {
	return s.SSLEnabled
}

func (s *DescribeDBClusterMigrationResponseBodyRdsEndpointListAddressItems) GetVPCId() *string {
	return s.VPCId
}

func (s *DescribeDBClusterMigrationResponseBodyRdsEndpointListAddressItems) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribeDBClusterMigrationResponseBodyRdsEndpointListAddressItems) SetConnectionString(v string) *DescribeDBClusterMigrationResponseBodyRdsEndpointListAddressItems {
	s.ConnectionString = &v
	return s
}

func (s *DescribeDBClusterMigrationResponseBodyRdsEndpointListAddressItems) SetIPAddress(v string) *DescribeDBClusterMigrationResponseBodyRdsEndpointListAddressItems {
	s.IPAddress = &v
	return s
}

func (s *DescribeDBClusterMigrationResponseBodyRdsEndpointListAddressItems) SetNetType(v string) *DescribeDBClusterMigrationResponseBodyRdsEndpointListAddressItems {
	s.NetType = &v
	return s
}

func (s *DescribeDBClusterMigrationResponseBodyRdsEndpointListAddressItems) SetPort(v string) *DescribeDBClusterMigrationResponseBodyRdsEndpointListAddressItems {
	s.Port = &v
	return s
}

func (s *DescribeDBClusterMigrationResponseBodyRdsEndpointListAddressItems) SetSSLEnabled(v string) *DescribeDBClusterMigrationResponseBodyRdsEndpointListAddressItems {
	s.SSLEnabled = &v
	return s
}

func (s *DescribeDBClusterMigrationResponseBodyRdsEndpointListAddressItems) SetVPCId(v string) *DescribeDBClusterMigrationResponseBodyRdsEndpointListAddressItems {
	s.VPCId = &v
	return s
}

func (s *DescribeDBClusterMigrationResponseBodyRdsEndpointListAddressItems) SetVSwitchId(v string) *DescribeDBClusterMigrationResponseBodyRdsEndpointListAddressItems {
	s.VSwitchId = &v
	return s
}

func (s *DescribeDBClusterMigrationResponseBodyRdsEndpointListAddressItems) Validate() error {
	return dara.Validate(s)
}
