// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRelateDbForHBaseHaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *RelateDbForHBaseHaRequest
	GetClusterId() *string
	SetHaActive(v string) *RelateDbForHBaseHaRequest
	GetHaActive() *string
	SetHaActiveClusterKey(v string) *RelateDbForHBaseHaRequest
	GetHaActiveClusterKey() *string
	SetHaActiveDBType(v string) *RelateDbForHBaseHaRequest
	GetHaActiveDBType() *string
	SetHaActiveHbaseFsDir(v string) *RelateDbForHBaseHaRequest
	GetHaActiveHbaseFsDir() *string
	SetHaActiveHdfsUri(v string) *RelateDbForHBaseHaRequest
	GetHaActiveHdfsUri() *string
	SetHaActivePassword(v string) *RelateDbForHBaseHaRequest
	GetHaActivePassword() *string
	SetHaActiveUser(v string) *RelateDbForHBaseHaRequest
	GetHaActiveUser() *string
	SetHaActiveVersion(v string) *RelateDbForHBaseHaRequest
	GetHaActiveVersion() *string
	SetHaMigrateType(v string) *RelateDbForHBaseHaRequest
	GetHaMigrateType() *string
	SetHaStandby(v string) *RelateDbForHBaseHaRequest
	GetHaStandby() *string
	SetHaStandbyClusterKey(v string) *RelateDbForHBaseHaRequest
	GetHaStandbyClusterKey() *string
	SetHaStandbyDBType(v string) *RelateDbForHBaseHaRequest
	GetHaStandbyDBType() *string
	SetHaStandbyHbaseFsDir(v string) *RelateDbForHBaseHaRequest
	GetHaStandbyHbaseFsDir() *string
	SetHaStandbyHdfsUri(v string) *RelateDbForHBaseHaRequest
	GetHaStandbyHdfsUri() *string
	SetHaStandbyPassword(v string) *RelateDbForHBaseHaRequest
	GetHaStandbyPassword() *string
	SetHaStandbyUser(v string) *RelateDbForHBaseHaRequest
	GetHaStandbyUser() *string
	SetHaStandbyVersion(v string) *RelateDbForHBaseHaRequest
	GetHaStandbyVersion() *string
	SetHaTables(v string) *RelateDbForHBaseHaRequest
	GetHaTables() *string
	SetIsActiveStandard(v bool) *RelateDbForHBaseHaRequest
	GetIsActiveStandard() *bool
	SetIsStandbyStandard(v bool) *RelateDbForHBaseHaRequest
	GetIsStandbyStandard() *bool
}

type RelateDbForHBaseHaRequest struct {
	// The ID of the BDS cluster. You can call the [DescribeInstances](https://help.aliyun.com/document_detail/144595.html) operation to obtain the cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// bds-t4nj9v2x85******
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The instance ID of the primary instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// hb-bp1w6krp539******
	HaActive *string `json:"HaActive,omitempty" xml:"HaActive,omitempty"`
	// The ZooKeeper address of the primary instance. This parameter is required when the primary instance is a non-standard instance (IsActiveStandard is set to false).
	//
	// example:
	//
	// hb-t4naqsay5gn******-master1-001.hbase.singapore.rds.aliyuncs.com,hb-t4naqsay5gn******-master3-001.hbase.singapore.rds.aliyuncs.com,hb-t4naqsay5gn******-master2-001.hbase.singapore.rds.aliyuncs.com:2181:/hbase
	HaActiveClusterKey *string `json:"HaActiveClusterKey,omitempty" xml:"HaActiveClusterKey,omitempty"`
	// The cluster type of the primary instance. Valid values: **HBase*	- and **HBaseue**.
	//
	// This parameter is required.
	//
	// example:
	//
	// hbase
	HaActiveDBType *string `json:"HaActiveDBType,omitempty" xml:"HaActiveDBType,omitempty"`
	// The HDFS directory of the primary instance. This parameter is required when the primary instance is a non-standard instance (IsActiveStandard is set to false).
	//
	// example:
	//
	// /hbase
	HaActiveHbaseFsDir *string `json:"HaActiveHbaseFsDir,omitempty" xml:"HaActiveHbaseFsDir,omitempty"`
	// The HDFS URI of the primary instance. This parameter is required when the primary instance is a non-standard instance (IsActiveStandard is set to false).
	//
	// example:
	//
	// hdfs://hb-t4naqsay5gn******-master1-001.hbase.rds.aliyuncs.com:8020,hb-t4naqsay5gn******-master2-001.hbase.rds.aliyuncs.com:8020
	HaActiveHdfsUri *string `json:"HaActiveHdfsUri,omitempty" xml:"HaActiveHdfsUri,omitempty"`
	// The password that corresponds to the username of the primary instance. This parameter is required when the primary instance is **HBaseue**.
	//
	// example:
	//
	// root
	HaActivePassword *string `json:"HaActivePassword,omitempty" xml:"HaActivePassword,omitempty"`
	// The username of the primary instance. This parameter is required when the primary instance is **HBaseue**.
	//
	// example:
	//
	// root
	HaActiveUser *string `json:"HaActiveUser,omitempty" xml:"HaActiveUser,omitempty"`
	// The database engine version of the primary instance. This parameter is required when the primary instance is a non-standard instance (IsActiveStandard is set to false). Valid values:
	//
	// - **HBase1x**: HBase 1.x.
	//
	// - **HBase2x**: HBase 2.x.
	//
	// - **HBaseUE**: HBaseue.
	//
	// example:
	//
	// HBase2x
	HaActiveVersion *string `json:"HaActiveVersion,omitempty" xml:"HaActiveVersion,omitempty"`
	// The synchronization type. Valid values:
	//
	// - **CLUSTER**: instance-level synchronization.
	//
	// - **TABLE**: table-level synchronization.
	//
	// - **SKIP**: no synchronization required.
	//
	// This parameter is required.
	//
	// example:
	//
	// CLUSTER
	HaMigrateType *string `json:"HaMigrateType,omitempty" xml:"HaMigrateType,omitempty"`
	// The ID of the secondary instance cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// hb-bp1bl7iqzka******
	HaStandby *string `json:"HaStandby,omitempty" xml:"HaStandby,omitempty"`
	// The ZooKeeper address of the secondary instance. This parameter is required when the secondary instance is a non-standard instance (IsStandbyStandard is set to false).
	//
	// example:
	//
	// hb-bp1w6krp539******-master1-001.hbase.singapore.rds.aliyuncs.com,hb-bp1w6krp539******-master3-001.hbase.singapore.rds.aliyuncs.com,hb-t4naqsay5gn******-master2-001.hbase.singapore.rds.aliyuncs.com:2181:/hbase
	HaStandbyClusterKey *string `json:"HaStandbyClusterKey,omitempty" xml:"HaStandbyClusterKey,omitempty"`
	// The cluster type of the secondary instance. Valid values: **HBase*	- and **HBaseue**.
	//
	// This parameter is required.
	//
	// example:
	//
	// hbase
	HaStandbyDBType *string `json:"HaStandbyDBType,omitempty" xml:"HaStandbyDBType,omitempty"`
	// The HDFS directory of the secondary instance. This parameter is required when the secondary instance is a non-standard instance (IsStandbyStandard is set to false).
	//
	// example:
	//
	// /hbase
	HaStandbyHbaseFsDir *string `json:"HaStandbyHbaseFsDir,omitempty" xml:"HaStandbyHbaseFsDir,omitempty"`
	// The HDFS URI of the secondary instance. This parameter is required when the secondary instance is a non-standard instance (IsStandbyStandard is set to false).
	//
	// example:
	//
	// hdfs://hb-bp1w6krp539******-master1-001.hbase.rds.aliyuncs.com:8020,hb-bp1w6krp539******-master2-001.hbase.rds.aliyuncs.com:8020
	HaStandbyHdfsUri *string `json:"HaStandbyHdfsUri,omitempty" xml:"HaStandbyHdfsUri,omitempty"`
	// The password that corresponds to the username of the secondary instance. This parameter is required when the secondary instance is **hbaseue**.
	//
	// example:
	//
	// root
	HaStandbyPassword *string `json:"HaStandbyPassword,omitempty" xml:"HaStandbyPassword,omitempty"`
	// The username of the secondary instance. This parameter is required when the secondary instance is **hbaseue**.
	//
	// example:
	//
	// root
	HaStandbyUser *string `json:"HaStandbyUser,omitempty" xml:"HaStandbyUser,omitempty"`
	// The database engine version of the secondary instance. This parameter is required when the secondary instance is a non-standard instance (IsStandbyStandard is set to false). Valid values:
	//
	// - **HBase1x**: HBase 1.x.
	//
	// - **HBase2x**: HBase 2.x.
	//
	// - **HBaseUE**: HBaseue.
	//
	// example:
	//
	// HBase2x
	HaStandbyVersion *string `json:"HaStandbyVersion,omitempty" xml:"HaStandbyVersion,omitempty"`
	// The tables to synchronize. This parameter is required when HaMigrateType is set to TABLE. Separate multiple tables with commas (,).
	//
	// example:
	//
	// test,test1
	HaTables *string `json:"HaTables,omitempty" xml:"HaTables,omitempty"`
	// Specifies whether the primary instance is a standard instance. Set this parameter to **true*	- for a standard instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	IsActiveStandard *bool `json:"IsActiveStandard,omitempty" xml:"IsActiveStandard,omitempty"`
	// Specifies whether the secondary instance is a standard instance. Set this parameter to **true*	- for a standard instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	IsStandbyStandard *bool `json:"IsStandbyStandard,omitempty" xml:"IsStandbyStandard,omitempty"`
}

func (s RelateDbForHBaseHaRequest) String() string {
	return dara.Prettify(s)
}

func (s RelateDbForHBaseHaRequest) GoString() string {
	return s.String()
}

func (s *RelateDbForHBaseHaRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *RelateDbForHBaseHaRequest) GetHaActive() *string {
	return s.HaActive
}

func (s *RelateDbForHBaseHaRequest) GetHaActiveClusterKey() *string {
	return s.HaActiveClusterKey
}

func (s *RelateDbForHBaseHaRequest) GetHaActiveDBType() *string {
	return s.HaActiveDBType
}

func (s *RelateDbForHBaseHaRequest) GetHaActiveHbaseFsDir() *string {
	return s.HaActiveHbaseFsDir
}

func (s *RelateDbForHBaseHaRequest) GetHaActiveHdfsUri() *string {
	return s.HaActiveHdfsUri
}

func (s *RelateDbForHBaseHaRequest) GetHaActivePassword() *string {
	return s.HaActivePassword
}

func (s *RelateDbForHBaseHaRequest) GetHaActiveUser() *string {
	return s.HaActiveUser
}

func (s *RelateDbForHBaseHaRequest) GetHaActiveVersion() *string {
	return s.HaActiveVersion
}

func (s *RelateDbForHBaseHaRequest) GetHaMigrateType() *string {
	return s.HaMigrateType
}

func (s *RelateDbForHBaseHaRequest) GetHaStandby() *string {
	return s.HaStandby
}

func (s *RelateDbForHBaseHaRequest) GetHaStandbyClusterKey() *string {
	return s.HaStandbyClusterKey
}

func (s *RelateDbForHBaseHaRequest) GetHaStandbyDBType() *string {
	return s.HaStandbyDBType
}

func (s *RelateDbForHBaseHaRequest) GetHaStandbyHbaseFsDir() *string {
	return s.HaStandbyHbaseFsDir
}

func (s *RelateDbForHBaseHaRequest) GetHaStandbyHdfsUri() *string {
	return s.HaStandbyHdfsUri
}

func (s *RelateDbForHBaseHaRequest) GetHaStandbyPassword() *string {
	return s.HaStandbyPassword
}

func (s *RelateDbForHBaseHaRequest) GetHaStandbyUser() *string {
	return s.HaStandbyUser
}

func (s *RelateDbForHBaseHaRequest) GetHaStandbyVersion() *string {
	return s.HaStandbyVersion
}

func (s *RelateDbForHBaseHaRequest) GetHaTables() *string {
	return s.HaTables
}

func (s *RelateDbForHBaseHaRequest) GetIsActiveStandard() *bool {
	return s.IsActiveStandard
}

func (s *RelateDbForHBaseHaRequest) GetIsStandbyStandard() *bool {
	return s.IsStandbyStandard
}

func (s *RelateDbForHBaseHaRequest) SetClusterId(v string) *RelateDbForHBaseHaRequest {
	s.ClusterId = &v
	return s
}

func (s *RelateDbForHBaseHaRequest) SetHaActive(v string) *RelateDbForHBaseHaRequest {
	s.HaActive = &v
	return s
}

func (s *RelateDbForHBaseHaRequest) SetHaActiveClusterKey(v string) *RelateDbForHBaseHaRequest {
	s.HaActiveClusterKey = &v
	return s
}

func (s *RelateDbForHBaseHaRequest) SetHaActiveDBType(v string) *RelateDbForHBaseHaRequest {
	s.HaActiveDBType = &v
	return s
}

func (s *RelateDbForHBaseHaRequest) SetHaActiveHbaseFsDir(v string) *RelateDbForHBaseHaRequest {
	s.HaActiveHbaseFsDir = &v
	return s
}

func (s *RelateDbForHBaseHaRequest) SetHaActiveHdfsUri(v string) *RelateDbForHBaseHaRequest {
	s.HaActiveHdfsUri = &v
	return s
}

func (s *RelateDbForHBaseHaRequest) SetHaActivePassword(v string) *RelateDbForHBaseHaRequest {
	s.HaActivePassword = &v
	return s
}

func (s *RelateDbForHBaseHaRequest) SetHaActiveUser(v string) *RelateDbForHBaseHaRequest {
	s.HaActiveUser = &v
	return s
}

func (s *RelateDbForHBaseHaRequest) SetHaActiveVersion(v string) *RelateDbForHBaseHaRequest {
	s.HaActiveVersion = &v
	return s
}

func (s *RelateDbForHBaseHaRequest) SetHaMigrateType(v string) *RelateDbForHBaseHaRequest {
	s.HaMigrateType = &v
	return s
}

func (s *RelateDbForHBaseHaRequest) SetHaStandby(v string) *RelateDbForHBaseHaRequest {
	s.HaStandby = &v
	return s
}

func (s *RelateDbForHBaseHaRequest) SetHaStandbyClusterKey(v string) *RelateDbForHBaseHaRequest {
	s.HaStandbyClusterKey = &v
	return s
}

func (s *RelateDbForHBaseHaRequest) SetHaStandbyDBType(v string) *RelateDbForHBaseHaRequest {
	s.HaStandbyDBType = &v
	return s
}

func (s *RelateDbForHBaseHaRequest) SetHaStandbyHbaseFsDir(v string) *RelateDbForHBaseHaRequest {
	s.HaStandbyHbaseFsDir = &v
	return s
}

func (s *RelateDbForHBaseHaRequest) SetHaStandbyHdfsUri(v string) *RelateDbForHBaseHaRequest {
	s.HaStandbyHdfsUri = &v
	return s
}

func (s *RelateDbForHBaseHaRequest) SetHaStandbyPassword(v string) *RelateDbForHBaseHaRequest {
	s.HaStandbyPassword = &v
	return s
}

func (s *RelateDbForHBaseHaRequest) SetHaStandbyUser(v string) *RelateDbForHBaseHaRequest {
	s.HaStandbyUser = &v
	return s
}

func (s *RelateDbForHBaseHaRequest) SetHaStandbyVersion(v string) *RelateDbForHBaseHaRequest {
	s.HaStandbyVersion = &v
	return s
}

func (s *RelateDbForHBaseHaRequest) SetHaTables(v string) *RelateDbForHBaseHaRequest {
	s.HaTables = &v
	return s
}

func (s *RelateDbForHBaseHaRequest) SetIsActiveStandard(v bool) *RelateDbForHBaseHaRequest {
	s.IsActiveStandard = &v
	return s
}

func (s *RelateDbForHBaseHaRequest) SetIsStandbyStandard(v bool) *RelateDbForHBaseHaRequest {
	s.IsStandbyStandard = &v
	return s
}

func (s *RelateDbForHBaseHaRequest) Validate() error {
	return dara.Validate(s)
}
