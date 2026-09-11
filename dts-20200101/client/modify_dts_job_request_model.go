// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDtsJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *ModifyDtsJobRequest
	GetClientToken() *string
	SetDataInitialization(v bool) *ModifyDtsJobRequest
	GetDataInitialization() *bool
	SetDataSynchronization(v bool) *ModifyDtsJobRequest
	GetDataSynchronization() *bool
	SetDbList(v map[string]interface{}) *ModifyDtsJobRequest
	GetDbList() map[string]interface{}
	SetDtsInstanceId(v string) *ModifyDtsJobRequest
	GetDtsInstanceId() *string
	SetDtsJobId(v string) *ModifyDtsJobRequest
	GetDtsJobId() *string
	SetEtlOperatorColumnReference(v string) *ModifyDtsJobRequest
	GetEtlOperatorColumnReference() *string
	SetFileOssUrl(v string) *ModifyDtsJobRequest
	GetFileOssUrl() *string
	SetFilterTableName(v string) *ModifyDtsJobRequest
	GetFilterTableName() *string
	SetModifyTypeEnum(v string) *ModifyDtsJobRequest
	GetModifyTypeEnum() *string
	SetRegionId(v string) *ModifyDtsJobRequest
	GetRegionId() *string
	SetReserved(v string) *ModifyDtsJobRequest
	GetReserved() *string
	SetResourceGroupId(v string) *ModifyDtsJobRequest
	GetResourceGroupId() *string
	SetStructureInitialization(v bool) *ModifyDtsJobRequest
	GetStructureInitialization() *bool
	SetSynchronizationDirection(v string) *ModifyDtsJobRequest
	GetSynchronizationDirection() *string
	SetZeroEtlJob(v bool) *ModifyDtsJobRequest
	GetZeroEtlJob() *bool
}

type ModifyDtsJobRequest struct {
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
	DbList map[string]interface{} `json:"DbList,omitempty" xml:"DbList,omitempty"`
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

func (s ModifyDtsJobRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDtsJobRequest) GoString() string {
	return s.String()
}

func (s *ModifyDtsJobRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ModifyDtsJobRequest) GetDataInitialization() *bool {
	return s.DataInitialization
}

func (s *ModifyDtsJobRequest) GetDataSynchronization() *bool {
	return s.DataSynchronization
}

func (s *ModifyDtsJobRequest) GetDbList() map[string]interface{} {
	return s.DbList
}

func (s *ModifyDtsJobRequest) GetDtsInstanceId() *string {
	return s.DtsInstanceId
}

func (s *ModifyDtsJobRequest) GetDtsJobId() *string {
	return s.DtsJobId
}

func (s *ModifyDtsJobRequest) GetEtlOperatorColumnReference() *string {
	return s.EtlOperatorColumnReference
}

func (s *ModifyDtsJobRequest) GetFileOssUrl() *string {
	return s.FileOssUrl
}

func (s *ModifyDtsJobRequest) GetFilterTableName() *string {
	return s.FilterTableName
}

func (s *ModifyDtsJobRequest) GetModifyTypeEnum() *string {
	return s.ModifyTypeEnum
}

func (s *ModifyDtsJobRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyDtsJobRequest) GetReserved() *string {
	return s.Reserved
}

func (s *ModifyDtsJobRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ModifyDtsJobRequest) GetStructureInitialization() *bool {
	return s.StructureInitialization
}

func (s *ModifyDtsJobRequest) GetSynchronizationDirection() *string {
	return s.SynchronizationDirection
}

func (s *ModifyDtsJobRequest) GetZeroEtlJob() *bool {
	return s.ZeroEtlJob
}

func (s *ModifyDtsJobRequest) SetClientToken(v string) *ModifyDtsJobRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyDtsJobRequest) SetDataInitialization(v bool) *ModifyDtsJobRequest {
	s.DataInitialization = &v
	return s
}

func (s *ModifyDtsJobRequest) SetDataSynchronization(v bool) *ModifyDtsJobRequest {
	s.DataSynchronization = &v
	return s
}

func (s *ModifyDtsJobRequest) SetDbList(v map[string]interface{}) *ModifyDtsJobRequest {
	s.DbList = v
	return s
}

func (s *ModifyDtsJobRequest) SetDtsInstanceId(v string) *ModifyDtsJobRequest {
	s.DtsInstanceId = &v
	return s
}

func (s *ModifyDtsJobRequest) SetDtsJobId(v string) *ModifyDtsJobRequest {
	s.DtsJobId = &v
	return s
}

func (s *ModifyDtsJobRequest) SetEtlOperatorColumnReference(v string) *ModifyDtsJobRequest {
	s.EtlOperatorColumnReference = &v
	return s
}

func (s *ModifyDtsJobRequest) SetFileOssUrl(v string) *ModifyDtsJobRequest {
	s.FileOssUrl = &v
	return s
}

func (s *ModifyDtsJobRequest) SetFilterTableName(v string) *ModifyDtsJobRequest {
	s.FilterTableName = &v
	return s
}

func (s *ModifyDtsJobRequest) SetModifyTypeEnum(v string) *ModifyDtsJobRequest {
	s.ModifyTypeEnum = &v
	return s
}

func (s *ModifyDtsJobRequest) SetRegionId(v string) *ModifyDtsJobRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyDtsJobRequest) SetReserved(v string) *ModifyDtsJobRequest {
	s.Reserved = &v
	return s
}

func (s *ModifyDtsJobRequest) SetResourceGroupId(v string) *ModifyDtsJobRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ModifyDtsJobRequest) SetStructureInitialization(v bool) *ModifyDtsJobRequest {
	s.StructureInitialization = &v
	return s
}

func (s *ModifyDtsJobRequest) SetSynchronizationDirection(v string) *ModifyDtsJobRequest {
	s.SynchronizationDirection = &v
	return s
}

func (s *ModifyDtsJobRequest) SetZeroEtlJob(v bool) *ModifyDtsJobRequest {
	s.ZeroEtlJob = &v
	return s
}

func (s *ModifyDtsJobRequest) Validate() error {
	return dara.Validate(s)
}
