// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeLindormInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterStorage(v int32) *UpgradeLindormInstanceRequest
	GetClusterStorage() *int32
	SetColdStorage(v int32) *UpgradeLindormInstanceRequest
	GetColdStorage() *int32
	SetCoreSingleStorage(v int32) *UpgradeLindormInstanceRequest
	GetCoreSingleStorage() *int32
	SetFilestoreNum(v int32) *UpgradeLindormInstanceRequest
	GetFilestoreNum() *int32
	SetFilestoreSpec(v string) *UpgradeLindormInstanceRequest
	GetFilestoreSpec() *string
	SetInstanceId(v string) *UpgradeLindormInstanceRequest
	GetInstanceId() *string
	SetLindormNum(v int32) *UpgradeLindormInstanceRequest
	GetLindormNum() *int32
	SetLindormSpec(v string) *UpgradeLindormInstanceRequest
	GetLindormSpec() *string
	SetLogNum(v int32) *UpgradeLindormInstanceRequest
	GetLogNum() *int32
	SetLogSingleStorage(v int32) *UpgradeLindormInstanceRequest
	GetLogSingleStorage() *int32
	SetLogSpec(v string) *UpgradeLindormInstanceRequest
	GetLogSpec() *string
	SetLtsCoreNum(v int32) *UpgradeLindormInstanceRequest
	GetLtsCoreNum() *int32
	SetLtsCoreSpec(v string) *UpgradeLindormInstanceRequest
	GetLtsCoreSpec() *string
	SetOwnerAccount(v string) *UpgradeLindormInstanceRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *UpgradeLindormInstanceRequest
	GetOwnerId() *int64
	SetRegionId(v string) *UpgradeLindormInstanceRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *UpgradeLindormInstanceRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *UpgradeLindormInstanceRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *UpgradeLindormInstanceRequest
	GetSecurityToken() *string
	SetSolrNum(v int32) *UpgradeLindormInstanceRequest
	GetSolrNum() *int32
	SetSolrSpec(v string) *UpgradeLindormInstanceRequest
	GetSolrSpec() *string
	SetStreamNum(v int32) *UpgradeLindormInstanceRequest
	GetStreamNum() *int32
	SetStreamSpec(v string) *UpgradeLindormInstanceRequest
	GetStreamSpec() *string
	SetTsdbNum(v int32) *UpgradeLindormInstanceRequest
	GetTsdbNum() *int32
	SetTsdbSpec(v string) *UpgradeLindormInstanceRequest
	GetTsdbSpec() *string
	SetUpgradeType(v string) *UpgradeLindormInstanceRequest
	GetUpgradeType() *string
	SetZoneId(v string) *UpgradeLindormInstanceRequest
	GetZoneId() *string
}

type UpgradeLindormInstanceRequest struct {
	// The storage capacity of the instance after it is upgraded. Unit: GB. Valid values: **480*	- to **1017600**.
	//
	// example:
	//
	// 480
	ClusterStorage *int32 `json:"ClusterStorage,omitempty" xml:"ClusterStorage,omitempty"`
	// The cold storage capacity of the instance after it is upgraded. Unit: GB. Valid values: **800*	- to **1000000**.
	//
	// example:
	//
	// 800
	ColdStorage *int32 `json:"ColdStorage,omitempty" xml:"ColdStorage,omitempty"`
	// The storage capacity of a single core node in the instance after the instance upgraded. This parameter is available only if the instance you want to upgrade is a multi-zone instance. Unit: GB. Valid values: 400 to 64000. **This parameter is optional**.
	//
	// example:
	//
	// 400
	CoreSingleStorage *int32 `json:"CoreSingleStorage,omitempty" xml:"CoreSingleStorage,omitempty"`
	// The number of LindormDFS nodes in the instance after the instance is upgraded. Valid values: integers from **0*	- to **60**.
	//
	// example:
	//
	// 2
	FilestoreNum *int32 `json:"FilestoreNum,omitempty" xml:"FilestoreNum,omitempty"`
	// The specification of LindormDFS nodes in the instance after the instance is upgraded. Valid values:
	//
	// 	- **lindorm.g.xlarge**: Each node has 4 dedicated CPU cores and 16 GB of dedicated memory.
	//
	// 	- **lindorm.g.2xlarge**: Each node has 8 dedicated CPU cores and 32 GB of dedicated memory.
	//
	// 	- **lindorm.g.4xlarge**: Each node has 16 dedicated CPU cores and 64 GB of dedicated memory.
	//
	// 	- **lindorm.g.8xlarge**: Each node has 32 dedicated CPU cores and 128 GB of dedicated memory.
	//
	// example:
	//
	// lindorm.g.xlarge
	FilestoreSpec *string `json:"FilestoreSpec,omitempty" xml:"FilestoreSpec,omitempty"`
	// The ID of the instance that you want to upgrade, scale up, or enable cold storage. You can call the [GetLindormInstanceList](https://help.aliyun.com/document_detail/426069.html) operation to query the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ld-bp1o3y0yme2i2****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The number of LindormTable nodes in the instance after the instance is upgraded. Valid values: integers from **0*	- to **90**.
	//
	// > This parameter must be specified together with the LindormSpec parameter.
	//
	// example:
	//
	// 2
	LindormNum *int32 `json:"LindormNum,omitempty" xml:"LindormNum,omitempty"`
	// The specification of LindormTable nodes in the instance after the instance is upgraded. Valid values:
	//
	// 	- **lindorm.c.xlarge**: Each node has 4 dedicated CPU cores and 8 GB of dedicated memory.
	//
	// 	- **lindorm.c.2xlarge**: Each node has 8 dedicated CPU cores and 16 GB of dedicated memory.
	//
	// 	- **lindorm.c.4xlarge**: Each node has 16 dedicated CPU cores and 32 GB of dedicated memory.
	//
	// 	- **lindorm.c.8xlarge**: Each node has 32 dedicated CPU cores and 64 GB of dedicated memory.
	//
	// example:
	//
	// lindorm.c.xlarge
	LindormSpec *string `json:"LindormSpec,omitempty" xml:"LindormSpec,omitempty"`
	// The number of log nodes in the instance after the instance is upgraded. This parameter is available only if the instance you want to upgrade is a multi-zone instance. **This parameter is optional**.
	//
	// example:
	//
	// 4
	LogNum *int32 `json:"LogNum,omitempty" xml:"LogNum,omitempty"`
	// The storage capacity of a single log node in the instance after the instance upgraded. This parameter is available only if the instance you want to upgrade is a multi-zone instance. **This parameter is optional**.
	//
	// example:
	//
	// 400
	LogSingleStorage *int32 `json:"LogSingleStorage,omitempty" xml:"LogSingleStorage,omitempty"`
	// The specification of log nodes in the instance after the instance is upgraded. This parameter is available only if the instance you want to upgrade is a multi-zone instance. Valid values:
	//
	// 	- **lindorm.sn1.large**: Each node has 4 dedicated CPU cores and 8 GB of dedicated memory.
	//
	// 	- **lindorm.sn1.2xlarge**: Each node has 8 dedicated CPU cores and 16 GB of dedicated memory.
	//
	// **This parameter is optional**.
	//
	// example:
	//
	// lindorm.sn1.large
	LogSpec *string `json:"LogSpec,omitempty" xml:"LogSpec,omitempty"`
	// The number of LTS nodes in the instance after the instance is upgraded. Valid values: integers from **0*	- to **50**.
	//
	// example:
	//
	// 2
	LtsCoreNum *int32 `json:"LtsCoreNum,omitempty" xml:"LtsCoreNum,omitempty"`
	// The specification of Lindorm Tunnel Service (LTS) nodes in the instance after the instance is upgraded. Valid values:
	//
	// 	- **lindorm.g.xlarge**: Each node has 4 dedicated CPU cores and 16 GB of dedicated memory.
	//
	// 	- **lindorm.g.2xlarge**: Each node has 8 dedicated CPU cores and 32 GB of dedicated memory.
	//
	// example:
	//
	// lindorm.g.xlarge
	LtsCoreSpec  *string `json:"LtsCoreSpec,omitempty" xml:"LtsCoreSpec,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region in which the instance that you want to upgrade, scale up, or enable cold storage is located. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/426062.html) operation to query the region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The number of LindormSearch nodes in the instance after the instance is upgraded. Valid values: integers from **0*	- to **60**.
	//
	// example:
	//
	// 2
	SolrNum *int32 `json:"SolrNum,omitempty" xml:"SolrNum,omitempty"`
	// The specification of LindormSearch nodes in the instance after the instance is upgraded. Valid values:
	//
	// 	- **lindorm.g.xlarge**: Each node has 4 dedicated CPU cores and 16 GB of dedicated memory.
	//
	// 	- **lindorm.g.2xlarge**: Each node has 8 dedicated CPU cores and 32 GB of dedicated memory.
	//
	// 	- **lindorm.g.4xlarge**: Each node has 16 dedicated CPU cores and 64 GB of dedicated memory.
	//
	// 	- **lindorm.g.8xlarge**: Each node has 32 dedicated CPU cores and 128 GB of dedicated memory.
	//
	// example:
	//
	// lindorm.g.xlarge
	SolrSpec *string `json:"SolrSpec,omitempty" xml:"SolrSpec,omitempty"`
	// The number of LindormStream nodes in the instance after the instance is upgraded. Valid values: integers from **0*	- to **60**.
	//
	// example:
	//
	// 2
	StreamNum *int32 `json:"StreamNum,omitempty" xml:"StreamNum,omitempty"`
	// The specification of LindormStream nodes in the instance after the instance is upgraded. Valid values:
	//
	// 	- **lindorm.g.xlarge**: Each node has 4 dedicated CPU cores and 16 GB of dedicated memory.
	//
	// 	- **lindorm.g.2xlarge**: Each node has 8 dedicated CPU cores and 32 GB of dedicated memory.
	//
	// 	- **lindorm.g.4xlarge**: Each node has 16 dedicated CPU cores and 64 GB of dedicated memory.
	//
	// 	- **lindorm.g.8xlarge**: Each node has 32 dedicated CPU cores and 128 GB of dedicated memory.
	//
	// example:
	//
	// lindorm.g.xlarge
	StreamSpec *string `json:"StreamSpec,omitempty" xml:"StreamSpec,omitempty"`
	// The number of LindormTSDB nodes in the instance after the instance is upgraded. Valid values: integers from **0*	- to **24**.
	//
	// example:
	//
	// 2
	TsdbNum *int32 `json:"TsdbNum,omitempty" xml:"TsdbNum,omitempty"`
	// The specification of LindormTSDB nodes in the instance after the instance is upgraded. Valid values:
	//
	// 	- **lindorm.g.xlarge**: Each node has 4 dedicated CPU cores and 16 GB of dedicated memory.
	//
	// 	- **lindorm.g.2xlarge**: Each node has 8 dedicated CPU cores and 32 GB of dedicated memory.
	//
	// 	- **lindorm.g.4xlarge**: Each node has 16 dedicated CPU cores and 64 GB of dedicated memory.
	//
	// 	- **lindorm.g.8xlarge**: Each node has 32 dedicated CPU cores and 128 GB of dedicated memory.
	//
	// example:
	//
	// lindorm.g.xlarge
	TsdbSpec *string `json:"TsdbSpec,omitempty" xml:"TsdbSpec,omitempty"`
	// The upgrade type of the operation. For more information about upgrade types, see the UpgradeType parameters section.
	//
	// This parameter is required.
	//
	// example:
	//
	// upgrade-cold-storage
	UpgradeType *string `json:"UpgradeType,omitempty" xml:"UpgradeType,omitempty"`
	// The ID of the zone in which the instance that you want to upgrade, scale up, or enable cold storage is located. You can call the [GetLindormInstance](https://help.aliyun.com/document_detail/426067.html) operation to query the zone ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai-f
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s UpgradeLindormInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s UpgradeLindormInstanceRequest) GoString() string {
	return s.String()
}

func (s *UpgradeLindormInstanceRequest) GetClusterStorage() *int32 {
	return s.ClusterStorage
}

func (s *UpgradeLindormInstanceRequest) GetColdStorage() *int32 {
	return s.ColdStorage
}

func (s *UpgradeLindormInstanceRequest) GetCoreSingleStorage() *int32 {
	return s.CoreSingleStorage
}

func (s *UpgradeLindormInstanceRequest) GetFilestoreNum() *int32 {
	return s.FilestoreNum
}

func (s *UpgradeLindormInstanceRequest) GetFilestoreSpec() *string {
	return s.FilestoreSpec
}

func (s *UpgradeLindormInstanceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpgradeLindormInstanceRequest) GetLindormNum() *int32 {
	return s.LindormNum
}

func (s *UpgradeLindormInstanceRequest) GetLindormSpec() *string {
	return s.LindormSpec
}

func (s *UpgradeLindormInstanceRequest) GetLogNum() *int32 {
	return s.LogNum
}

func (s *UpgradeLindormInstanceRequest) GetLogSingleStorage() *int32 {
	return s.LogSingleStorage
}

func (s *UpgradeLindormInstanceRequest) GetLogSpec() *string {
	return s.LogSpec
}

func (s *UpgradeLindormInstanceRequest) GetLtsCoreNum() *int32 {
	return s.LtsCoreNum
}

func (s *UpgradeLindormInstanceRequest) GetLtsCoreSpec() *string {
	return s.LtsCoreSpec
}

func (s *UpgradeLindormInstanceRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *UpgradeLindormInstanceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UpgradeLindormInstanceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpgradeLindormInstanceRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *UpgradeLindormInstanceRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *UpgradeLindormInstanceRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *UpgradeLindormInstanceRequest) GetSolrNum() *int32 {
	return s.SolrNum
}

func (s *UpgradeLindormInstanceRequest) GetSolrSpec() *string {
	return s.SolrSpec
}

func (s *UpgradeLindormInstanceRequest) GetStreamNum() *int32 {
	return s.StreamNum
}

func (s *UpgradeLindormInstanceRequest) GetStreamSpec() *string {
	return s.StreamSpec
}

func (s *UpgradeLindormInstanceRequest) GetTsdbNum() *int32 {
	return s.TsdbNum
}

func (s *UpgradeLindormInstanceRequest) GetTsdbSpec() *string {
	return s.TsdbSpec
}

func (s *UpgradeLindormInstanceRequest) GetUpgradeType() *string {
	return s.UpgradeType
}

func (s *UpgradeLindormInstanceRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *UpgradeLindormInstanceRequest) SetClusterStorage(v int32) *UpgradeLindormInstanceRequest {
	s.ClusterStorage = &v
	return s
}

func (s *UpgradeLindormInstanceRequest) SetColdStorage(v int32) *UpgradeLindormInstanceRequest {
	s.ColdStorage = &v
	return s
}

func (s *UpgradeLindormInstanceRequest) SetCoreSingleStorage(v int32) *UpgradeLindormInstanceRequest {
	s.CoreSingleStorage = &v
	return s
}

func (s *UpgradeLindormInstanceRequest) SetFilestoreNum(v int32) *UpgradeLindormInstanceRequest {
	s.FilestoreNum = &v
	return s
}

func (s *UpgradeLindormInstanceRequest) SetFilestoreSpec(v string) *UpgradeLindormInstanceRequest {
	s.FilestoreSpec = &v
	return s
}

func (s *UpgradeLindormInstanceRequest) SetInstanceId(v string) *UpgradeLindormInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *UpgradeLindormInstanceRequest) SetLindormNum(v int32) *UpgradeLindormInstanceRequest {
	s.LindormNum = &v
	return s
}

func (s *UpgradeLindormInstanceRequest) SetLindormSpec(v string) *UpgradeLindormInstanceRequest {
	s.LindormSpec = &v
	return s
}

func (s *UpgradeLindormInstanceRequest) SetLogNum(v int32) *UpgradeLindormInstanceRequest {
	s.LogNum = &v
	return s
}

func (s *UpgradeLindormInstanceRequest) SetLogSingleStorage(v int32) *UpgradeLindormInstanceRequest {
	s.LogSingleStorage = &v
	return s
}

func (s *UpgradeLindormInstanceRequest) SetLogSpec(v string) *UpgradeLindormInstanceRequest {
	s.LogSpec = &v
	return s
}

func (s *UpgradeLindormInstanceRequest) SetLtsCoreNum(v int32) *UpgradeLindormInstanceRequest {
	s.LtsCoreNum = &v
	return s
}

func (s *UpgradeLindormInstanceRequest) SetLtsCoreSpec(v string) *UpgradeLindormInstanceRequest {
	s.LtsCoreSpec = &v
	return s
}

func (s *UpgradeLindormInstanceRequest) SetOwnerAccount(v string) *UpgradeLindormInstanceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *UpgradeLindormInstanceRequest) SetOwnerId(v int64) *UpgradeLindormInstanceRequest {
	s.OwnerId = &v
	return s
}

func (s *UpgradeLindormInstanceRequest) SetRegionId(v string) *UpgradeLindormInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *UpgradeLindormInstanceRequest) SetResourceOwnerAccount(v string) *UpgradeLindormInstanceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpgradeLindormInstanceRequest) SetResourceOwnerId(v int64) *UpgradeLindormInstanceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UpgradeLindormInstanceRequest) SetSecurityToken(v string) *UpgradeLindormInstanceRequest {
	s.SecurityToken = &v
	return s
}

func (s *UpgradeLindormInstanceRequest) SetSolrNum(v int32) *UpgradeLindormInstanceRequest {
	s.SolrNum = &v
	return s
}

func (s *UpgradeLindormInstanceRequest) SetSolrSpec(v string) *UpgradeLindormInstanceRequest {
	s.SolrSpec = &v
	return s
}

func (s *UpgradeLindormInstanceRequest) SetStreamNum(v int32) *UpgradeLindormInstanceRequest {
	s.StreamNum = &v
	return s
}

func (s *UpgradeLindormInstanceRequest) SetStreamSpec(v string) *UpgradeLindormInstanceRequest {
	s.StreamSpec = &v
	return s
}

func (s *UpgradeLindormInstanceRequest) SetTsdbNum(v int32) *UpgradeLindormInstanceRequest {
	s.TsdbNum = &v
	return s
}

func (s *UpgradeLindormInstanceRequest) SetTsdbSpec(v string) *UpgradeLindormInstanceRequest {
	s.TsdbSpec = &v
	return s
}

func (s *UpgradeLindormInstanceRequest) SetUpgradeType(v string) *UpgradeLindormInstanceRequest {
	s.UpgradeType = &v
	return s
}

func (s *UpgradeLindormInstanceRequest) SetZoneId(v string) *UpgradeLindormInstanceRequest {
	s.ZoneId = &v
	return s
}

func (s *UpgradeLindormInstanceRequest) Validate() error {
	return dara.Validate(s)
}
