// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDtsJobShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *ModifyDtsJobShrinkRequest
	GetClientToken() *string
	SetDataInitialization(v bool) *ModifyDtsJobShrinkRequest
	GetDataInitialization() *bool
	SetDataSynchronization(v bool) *ModifyDtsJobShrinkRequest
	GetDataSynchronization() *bool
	SetDbListShrink(v string) *ModifyDtsJobShrinkRequest
	GetDbListShrink() *string
	SetDtsInstanceId(v string) *ModifyDtsJobShrinkRequest
	GetDtsInstanceId() *string
	SetDtsJobId(v string) *ModifyDtsJobShrinkRequest
	GetDtsJobId() *string
	SetEtlOperatorColumnReference(v string) *ModifyDtsJobShrinkRequest
	GetEtlOperatorColumnReference() *string
	SetFileOssUrl(v string) *ModifyDtsJobShrinkRequest
	GetFileOssUrl() *string
	SetFilterTableName(v string) *ModifyDtsJobShrinkRequest
	GetFilterTableName() *string
	SetModifyTypeEnum(v string) *ModifyDtsJobShrinkRequest
	GetModifyTypeEnum() *string
	SetRegionId(v string) *ModifyDtsJobShrinkRequest
	GetRegionId() *string
	SetReserved(v string) *ModifyDtsJobShrinkRequest
	GetReserved() *string
	SetResourceGroupId(v string) *ModifyDtsJobShrinkRequest
	GetResourceGroupId() *string
	SetStructureInitialization(v bool) *ModifyDtsJobShrinkRequest
	GetStructureInitialization() *bool
	SetSynchronizationDirection(v string) *ModifyDtsJobShrinkRequest
	GetSynchronizationDirection() *string
	SetZeroEtlJob(v bool) *ModifyDtsJobShrinkRequest
	GetZeroEtlJob() *bool
}

type ModifyDtsJobShrinkRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the value, but you must make sure that it is unique among different requests. **ClientToken*	- can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 0c593ea1-3bea-11e9-b96b-88e9fe63****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform full data migration or initial full data synchronization. Valid values:
	//
	// - **true**: yes.
	//
	// - **false**: no.
	//
	// example:
	//
	// true
	DataInitialization *bool `json:"DataInitialization,omitempty" xml:"DataInitialization,omitempty"`
	// Specifies whether to perform incremental data migration or synchronization. Valid values:
	//
	// - **false**: no.
	//
	// - **true**: yes.
	//
	// example:
	//
	// true
	DataSynchronization *bool `json:"DataSynchronization,omitempty" xml:"DataSynchronization,omitempty"`
	// The modified synchronization objects, in JSON format. For more information about the definition, see [Objects of DTS tasks](https://help.aliyun.com/document_detail/209545.html).
	//
	// - The original DbList is overwritten by the new DbList. Make sure that the new DbList contains all the objects that need to be synchronized. Otherwise, synchronization objects may be lost. Modify this parameter with caution.
	//
	// - Call [DescribeDtsJobDetail](https://help.aliyun.com/document_detail/208925.html) to query the current synchronization objects before you modify them based on your business requirements. For example, if the current objects are tables A and B, and you want to add table C, specify tables A, B, and C in this parameter.
	//
	// - The maximum size of DbList is 1 MB.
	//
	// - If DbList contains filter conditions, the total length of DbList (including filter conditions) cannot exceed 1 MB.
	//
	// - For distributed tasks (such as migration or synchronization tasks whose source is PolarDB-X 1.0), DbList is split based on physical shards and multiple subtasks are generated. The maximum size of DbList for each subtask is 1 MB.
	//
	// example:
	//
	// {"dtstest":{"name":"dtstest","all":true}}
	DbListShrink *string `json:"DbList,omitempty" xml:"DbList,omitempty"`
	// The instance ID of the data synchronization instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// dtsbi6e22ay243****
	DtsInstanceId *string `json:"DtsInstanceId,omitempty" xml:"DtsInstanceId,omitempty"`
	// The ID of the synchronization task. You can call [DescribeDtsJobs](https://help.aliyun.com/document_detail/209702.html) to query the task ID.
	//
	// example:
	//
	// fpx1149rw7p***
	DtsJobId *string `json:"DtsJobId,omitempty" xml:"DtsJobId,omitempty"`
	// A field dedicated to T+1 business scenarios. This is an ETL operator and a business-specific field.
	//
	// example:
	//
	// {
	//
	//   "configKeyMap": [
	//
	//     {
	//
	//       "moduleCode": "03",
	//
	//       "etlOperatorColumnReference": "etlOperatorColumnReference",
	//
	//       "etlOperatorSetting": "etlOperatorSetting"
	//
	//     },
	//
	//     {
	//
	//       "moduleCode": "07",
	//
	//       "etlOperatorColumnReference": "etlOperatorColumnReference",
	//
	//       "etlOperatorSetting": "etlOperatorSetting"
	//
	//     }
	//
	//   ],
	//
	//   "amp.increment.replicator.compare.all.columns.enable": false,
	//
	//   "srcNetType": "VPC",
	//
	//   "etlOperatorColumnReference": {
	//
	//     "MODIFY_TIME": [
	//
	//       {
	//
	//         "datasynchron.test_timing_user_000": "user_create_date"
	//
	//       },
	//
	//       {
	//
	//         "datasynchron.test_timing_user_001": "user_create_date"
	//
	//       },
	//
	//       {
	//
	//         "datasynchron.test_timing_user_002": "user_create_date"
	//
	//       },
	//
	//       {
	//
	//         "datasynchron.test_timing_user_003": "user_create_date"
	//
	//       }
	//
	//     ]
	//
	//   },
	//
	//   "destNetType": "VPC",
	//
	//   "originalSrcDbInst": "sg-sit-db2-primary.mysql.polardb.rds.aliyuncs.com",
	//
	//   "etlOperatorSetting": "adb_time_travel",
	//
	//   "sjobUseJobTask": "1",
	//
	//   "srcVpcNetMappingInst": "172.19.XXX.XXX:4261",
	//
	//   "destVpcNetMappingInst": "172.19.XXX.XXX:4260",
	//
	//   "useJobTask": "1"
	//
	// }
	EtlOperatorColumnReference *string `json:"EtlOperatorColumnReference,omitempty" xml:"EtlOperatorColumnReference,omitempty"`
	// The OSS URL of the synchronization file.
	//
	// example:
	//
	// http://db-list-os-file.oss-cn-shanghai.aliyuncs.com/8e42_12185******43989_************
	FileOssUrl *string `json:"FileOssUrl,omitempty" xml:"FileOssUrl,omitempty"`
	// The name of the table to be filtered.
	//
	// example:
	//
	// testtable
	FilterTableName *string `json:"FilterTableName,omitempty" xml:"FilterTableName,omitempty"`
	// The method used to modify the synchronization task. If this parameter is not specified, the synchronization objects are modified by default. Set this parameter to UPDATE_RESERVED to modify reserved parameters.
	//
	// example:
	//
	// UPDATE_RESERVED
	ModifyTypeEnum *string `json:"ModifyTypeEnum,omitempty" xml:"ModifyTypeEnum,omitempty"`
	// The region in which the instance resides. For more information, see [Supported regions](https://help.aliyun.com/document_detail/141033.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The reserved parameters of DTS. The update method is append rather than overwrite. The value is in the MAP JSON format. You can specify this parameter to meet special requirements, such as whether to automatically start a precheck. The usage is similar to that of the Reserve parameter. For details, see [Reserve parameter description](https://help.aliyun.com/document_detail/273111.html).
	//
	// example:
	//
	// {"definer": false,"syncArchitecture": "oneway","whitelist.dms.online.ddl.enable": false,"destSSL": "0","triggerMode": "manual","sqlparser.dms.original.ddl": true,"whitelist.ghost.online.ddl.enable": false,"privilegeMigration": false,"maxRetryTime": 43200,"srcSSL": "0","autoStartModulesAfterConfig": "none"}
	Reserved *string `json:"Reserved,omitempty" xml:"Reserved,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmzawhxxc****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// Specifies whether to perform schema migration or initial schema synchronization. Valid values:
	//
	// - **true**: yes.
	//
	// - **false**: no.
	//
	// example:
	//
	// true
	StructureInitialization *bool `json:"StructureInitialization,omitempty" xml:"StructureInitialization,omitempty"`
	// The synchronization direction. Valid values:
	//
	// - **Forward**: forward.
	//
	// - **Reverse**: reverse.
	//
	// > - Default value: **Forward**.
	//
	// - This parameter is required only when the synchronization topology of the data synchronization instance is two-way synchronization.
	//
	// example:
	//
	// Forward
	SynchronizationDirection *string `json:"SynchronizationDirection,omitempty" xml:"SynchronizationDirection,omitempty"`
	// Specifies whether this is a seamless integration (zero-ETL) node. Valid values:
	//
	// - **true**: yes.
	//
	// - **false**: no.
	//
	// example:
	//
	// false
	ZeroEtlJob *bool `json:"ZeroEtlJob,omitempty" xml:"ZeroEtlJob,omitempty"`
}

func (s ModifyDtsJobShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDtsJobShrinkRequest) GoString() string {
	return s.String()
}

func (s *ModifyDtsJobShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ModifyDtsJobShrinkRequest) GetDataInitialization() *bool {
	return s.DataInitialization
}

func (s *ModifyDtsJobShrinkRequest) GetDataSynchronization() *bool {
	return s.DataSynchronization
}

func (s *ModifyDtsJobShrinkRequest) GetDbListShrink() *string {
	return s.DbListShrink
}

func (s *ModifyDtsJobShrinkRequest) GetDtsInstanceId() *string {
	return s.DtsInstanceId
}

func (s *ModifyDtsJobShrinkRequest) GetDtsJobId() *string {
	return s.DtsJobId
}

func (s *ModifyDtsJobShrinkRequest) GetEtlOperatorColumnReference() *string {
	return s.EtlOperatorColumnReference
}

func (s *ModifyDtsJobShrinkRequest) GetFileOssUrl() *string {
	return s.FileOssUrl
}

func (s *ModifyDtsJobShrinkRequest) GetFilterTableName() *string {
	return s.FilterTableName
}

func (s *ModifyDtsJobShrinkRequest) GetModifyTypeEnum() *string {
	return s.ModifyTypeEnum
}

func (s *ModifyDtsJobShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyDtsJobShrinkRequest) GetReserved() *string {
	return s.Reserved
}

func (s *ModifyDtsJobShrinkRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ModifyDtsJobShrinkRequest) GetStructureInitialization() *bool {
	return s.StructureInitialization
}

func (s *ModifyDtsJobShrinkRequest) GetSynchronizationDirection() *string {
	return s.SynchronizationDirection
}

func (s *ModifyDtsJobShrinkRequest) GetZeroEtlJob() *bool {
	return s.ZeroEtlJob
}

func (s *ModifyDtsJobShrinkRequest) SetClientToken(v string) *ModifyDtsJobShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyDtsJobShrinkRequest) SetDataInitialization(v bool) *ModifyDtsJobShrinkRequest {
	s.DataInitialization = &v
	return s
}

func (s *ModifyDtsJobShrinkRequest) SetDataSynchronization(v bool) *ModifyDtsJobShrinkRequest {
	s.DataSynchronization = &v
	return s
}

func (s *ModifyDtsJobShrinkRequest) SetDbListShrink(v string) *ModifyDtsJobShrinkRequest {
	s.DbListShrink = &v
	return s
}

func (s *ModifyDtsJobShrinkRequest) SetDtsInstanceId(v string) *ModifyDtsJobShrinkRequest {
	s.DtsInstanceId = &v
	return s
}

func (s *ModifyDtsJobShrinkRequest) SetDtsJobId(v string) *ModifyDtsJobShrinkRequest {
	s.DtsJobId = &v
	return s
}

func (s *ModifyDtsJobShrinkRequest) SetEtlOperatorColumnReference(v string) *ModifyDtsJobShrinkRequest {
	s.EtlOperatorColumnReference = &v
	return s
}

func (s *ModifyDtsJobShrinkRequest) SetFileOssUrl(v string) *ModifyDtsJobShrinkRequest {
	s.FileOssUrl = &v
	return s
}

func (s *ModifyDtsJobShrinkRequest) SetFilterTableName(v string) *ModifyDtsJobShrinkRequest {
	s.FilterTableName = &v
	return s
}

func (s *ModifyDtsJobShrinkRequest) SetModifyTypeEnum(v string) *ModifyDtsJobShrinkRequest {
	s.ModifyTypeEnum = &v
	return s
}

func (s *ModifyDtsJobShrinkRequest) SetRegionId(v string) *ModifyDtsJobShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyDtsJobShrinkRequest) SetReserved(v string) *ModifyDtsJobShrinkRequest {
	s.Reserved = &v
	return s
}

func (s *ModifyDtsJobShrinkRequest) SetResourceGroupId(v string) *ModifyDtsJobShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ModifyDtsJobShrinkRequest) SetStructureInitialization(v bool) *ModifyDtsJobShrinkRequest {
	s.StructureInitialization = &v
	return s
}

func (s *ModifyDtsJobShrinkRequest) SetSynchronizationDirection(v string) *ModifyDtsJobShrinkRequest {
	s.SynchronizationDirection = &v
	return s
}

func (s *ModifyDtsJobShrinkRequest) SetZeroEtlJob(v bool) *ModifyDtsJobShrinkRequest {
	s.ZeroEtlJob = &v
	return s
}

func (s *ModifyDtsJobShrinkRequest) Validate() error {
	return dara.Validate(s)
}
