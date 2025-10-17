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
	// The mode of the source ApsaraDB RDS instance. Valid values:
	//
	// 	- **rw**: read and write mode
	//
	// 	- **ro**: read-only mode
	//
	// example:
	//
	// test
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The port number.
	DBClusterEndpointList []*DescribeDBClusterMigrationResponseBodyDBClusterEndpointList `json:"DBClusterEndpointList,omitempty" xml:"DBClusterEndpointList,omitempty" type:"Repeated"`
	// The replication latency between the ApsaraDB RDS instance and the PolarDB cluster. Unit: seconds.
	//
	// example:
	//
	// pc-****************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// Details about the endpoints.
	//
	// example:
	//
	// ro
	DBClusterReadWriteMode *string `json:"DBClusterReadWriteMode,omitempty" xml:"DBClusterReadWriteMode,omitempty"`
	// The vSwitch ID.
	//
	// example:
	//
	// 0
	DelayedSeconds *int32 `json:"DelayedSeconds,omitempty" xml:"DelayedSeconds,omitempty"`
	// The network type of the endpoint. Valid values:
	//
	// 	- **Public**: the public endpoint
	//
	// 	- **Private**: the internal endpoint (VPC)
	//
	// 	- **Inner**: the internal endpoint (classic network)
	//
	// example:
	//
	// dts**********618bs
	DtsInstanceId *string `json:"DtsInstanceId,omitempty" xml:"DtsInstanceId,omitempty"`
	// The mode of the PolarDB cluster. Valid values:
	//
	// 	- **rw**: read and write mode
	//
	// 	- **ro**: read-only mode
	//
	// example:
	//
	// 2020-06-17T01:56:36Z
	ExpiredTime *string `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	// The endpoint.
	//
	// example:
	//
	// RDS2POLARDB_SYNCING
	MigrationStatus *string `json:"MigrationStatus,omitempty" xml:"MigrationStatus,omitempty"`
	// The endpoints of the ApsaraDB RDS instance.
	RdsEndpointList []*DescribeDBClusterMigrationResponseBodyRdsEndpointList `json:"RdsEndpointList,omitempty" xml:"RdsEndpointList,omitempty" type:"Repeated"`
	// The ID of the synchronous task.
	//
	// example:
	//
	// rw
	RdsReadWriteMode *string `json:"RdsReadWriteMode,omitempty" xml:"RdsReadWriteMode,omitempty"`
	// The ID of the source ApsaraDB RDS instance.
	//
	// example:
	//
	// F2A9EFA7-915F-4572-8299-85A307******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The endpoints of the ApsaraDB RDS instance.
	//
	// example:
	//
	// rm-************
	SourceRDSDBInstanceId *string `json:"SourceRDSDBInstanceId,omitempty" xml:"SourceRDSDBInstanceId,omitempty"`
	// The type of the source database. Valid values:
	//
	// - **PolarDBMySQL**: The source database is a PolarDB for MySQL database when the major version of your PolarDB cluster is upgraded.
	//
	// - **RDS**: The source database is an ApsaraDB RDS database when data is migrated from ApsaraDB RDS to PolarDB for MySQL.
	//
	// example:
	//
	// PolarDBMySQL
	SrcDbType *string `json:"SrcDbType,omitempty" xml:"SrcDbType,omitempty"`
	// The migration state of the PolarDB cluster. Valid values:
	//
	// 	- **NO_MIGRATION**: No migration task is running.
	//
	// 	- **RDS2POLARDB_CLONING**: Data is being replicated.
	//
	// 	- **RDS2POLARDB_SYNCING**: Data is being replicated. During the replication, the PolarDB cluster is running in read-only mode and the source ApsaraDB RDS instance is running in read and write mode.
	//
	// 	- **SWITCHING**: Databases are being switched.
	//
	// 	- **POLARDB2RDS_SYNCING**: Databases are switched. The PolarDB cluster is running in read and write mode and the source ApsaraDB RDS instance is running in read-only mode. In this state, you can modify the endpoints for your applications.
	//
	// 	- **ROLLBACK**: The migration is being rolled back. After the rollback is complete, the value **RDS2POLARDB_SYNCING*	- is returned.
	//
	// 	- **CLOSING_MIGRATION**: The migration task is being terminated.
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
	// The VPC ID.
	AddressItems []*DescribeDBClusterMigrationResponseBodyDBClusterEndpointListAddressItems `json:"AddressItems,omitempty" xml:"AddressItems,omitempty" type:"Repeated"`
	// The expiration time of the replication between ApsaraDB RDS and PolarDB. The time is in the `YYYY-MM-DDThh:mm:ssZ` format. The time is displayed in UTC.
	//
	// example:
	//
	// pe-***********
	DBEndpointId *string `json:"DBEndpointId,omitempty" xml:"DBEndpointId,omitempty"`
	// The ID of the cluster.
	//
	// example:
	//
	// Cluster
	EndpointType *string `json:"EndpointType,omitempty" xml:"EndpointType,omitempty"`
	// The synchronization direction. Valid values:
	//
	// 	- **RDS2POLARDB**: Data is replicated from an ApsaraDB RDS instance to a PolarDB cluster.
	//
	// 	- **POLARDB2RDS**: Data is replicated from a PolarDB cluster to an ApsaraDB RDS instance.
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
	// Indicates whether SSL encryption is enabled. Valid values:
	//
	// 	- **Enabled**
	//
	// 	- **Disabled**
	//
	// example:
	//
	// pc-**************.rwlb.rds.aliyuncs.com
	ConnectionString *string `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	// The description of a migration exception. If no exception occurs during the migration, an empty string is returned.
	//
	// example:
	//
	// 192.***.***.10
	IPAddress *string `json:"IPAddress,omitempty" xml:"IPAddress,omitempty"`
	// The ID of the endpoint.
	//
	// example:
	//
	// Private
	NetType *string `json:"NetType,omitempty" xml:"NetType,omitempty"`
	// The type of the endpoint. Valid values:
	//
	// 	- **Cluster**: the default cluster endpoint
	//
	// 	- **Primary**: the primary endpoint
	//
	// 	- **Custom**: the custom endpoint
	//
	// example:
	//
	// 3306
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// Enabled
	SSLEnabled *string `json:"SSLEnabled,omitempty" xml:"SSLEnabled,omitempty"`
	// The read/write mode. Valid values:
	//
	// 	- ReadWrite: receives and forwards read and write requests (automatic read-write splitting).
	//
	// 	- ReadOnly (default): receives and forwards read requests only.
	//
	// example:
	//
	// vpc-**********
	VPCId *string `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	// The IP address of the endpoint.
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
	// The VPC ID.
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
	// The ID of the endpoint.
	//
	// example:
	//
	// rm-************-normal
	DBEndpointId *string `json:"DBEndpointId,omitempty" xml:"DBEndpointId,omitempty"`
	// The type of the endpoint. Valid values:
	//
	// - **Normal**: the standard endpoint
	//
	// - **ReadWriteSplitting**: the read/write splitting endpoint
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
	// Indicates whether SSL encryption is enabled. Valid values:
	//
	// 	- **Enabled**
	//
	// 	- **Disabled**
	//
	// example:
	//
	// rm-***********.mysql.rds.aliyuncs.com
	ConnectionString *string `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	// The type of the source database. Valid values:
	//
	// 	- **PolarDBMySQL**: The source database is a PolarDB for MySQL database when the major version of your PolarDB cluster is upgraded.
	//
	// 	- **RDS**: The source database is an ApsaraDB RDS database when data is migrated from ApsaraDB RDS to PolarDB for MySQL.
	//
	// example:
	//
	// 172.***.***.173
	IPAddress *string `json:"IPAddress,omitempty" xml:"IPAddress,omitempty"`
	// The ID of the endpoint.
	//
	// example:
	//
	// Private
	NetType *string `json:"NetType,omitempty" xml:"NetType,omitempty"`
	// The type of the endpoint. Valid values:
	//
	// 	- **Normal**: the standard endpoint
	//
	// 	- **ReadWriteSplitting**: the read/write splitting endpoint
	//
	// example:
	//
	// 3306
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// Indicates whether SSL encryption is enabled. Valid values:
	//
	// - **Enabled**
	//
	// - **Disabled**
	//
	// example:
	//
	// Enabled
	SSLEnabled *string `json:"SSLEnabled,omitempty" xml:"SSLEnabled,omitempty"`
	// The instance type.
	//
	// example:
	//
	// vpc-************
	VPCId *string `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	// The IP address of the endpoint.
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
