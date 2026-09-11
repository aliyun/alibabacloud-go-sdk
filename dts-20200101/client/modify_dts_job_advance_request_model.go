// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
	"io"
)

type iModifyDtsJobAdvanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *ModifyDtsJobAdvanceRequest
	GetClientToken() *string
	SetDataInitialization(v bool) *ModifyDtsJobAdvanceRequest
	GetDataInitialization() *bool
	SetDataSynchronization(v bool) *ModifyDtsJobAdvanceRequest
	GetDataSynchronization() *bool
	SetDbList(v map[string]interface{}) *ModifyDtsJobAdvanceRequest
	GetDbList() map[string]interface{}
	SetDtsInstanceId(v string) *ModifyDtsJobAdvanceRequest
	GetDtsInstanceId() *string
	SetDtsJobId(v string) *ModifyDtsJobAdvanceRequest
	GetDtsJobId() *string
	SetEtlOperatorColumnReference(v string) *ModifyDtsJobAdvanceRequest
	GetEtlOperatorColumnReference() *string
	SetFileOssUrlObject(v io.Reader) *ModifyDtsJobAdvanceRequest
	GetFileOssUrlObject() io.Reader
	SetFilterTableName(v string) *ModifyDtsJobAdvanceRequest
	GetFilterTableName() *string
	SetModifyTypeEnum(v string) *ModifyDtsJobAdvanceRequest
	GetModifyTypeEnum() *string
	SetRegionId(v string) *ModifyDtsJobAdvanceRequest
	GetRegionId() *string
	SetReserved(v string) *ModifyDtsJobAdvanceRequest
	GetReserved() *string
	SetResourceGroupId(v string) *ModifyDtsJobAdvanceRequest
	GetResourceGroupId() *string
	SetStructureInitialization(v bool) *ModifyDtsJobAdvanceRequest
	GetStructureInitialization() *bool
	SetSynchronizationDirection(v string) *ModifyDtsJobAdvanceRequest
	GetSynchronizationDirection() *string
	SetZeroEtlJob(v bool) *ModifyDtsJobAdvanceRequest
	GetZeroEtlJob() *bool
}

type ModifyDtsJobAdvanceRequest struct {
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
	FileOssUrlObject io.Reader `json:"FileOssUrl,omitempty" xml:"FileOssUrl,omitempty"`
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

func (s ModifyDtsJobAdvanceRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDtsJobAdvanceRequest) GoString() string {
	return s.String()
}

func (s *ModifyDtsJobAdvanceRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ModifyDtsJobAdvanceRequest) GetDataInitialization() *bool {
	return s.DataInitialization
}

func (s *ModifyDtsJobAdvanceRequest) GetDataSynchronization() *bool {
	return s.DataSynchronization
}

func (s *ModifyDtsJobAdvanceRequest) GetDbList() map[string]interface{} {
	return s.DbList
}

func (s *ModifyDtsJobAdvanceRequest) GetDtsInstanceId() *string {
	return s.DtsInstanceId
}

func (s *ModifyDtsJobAdvanceRequest) GetDtsJobId() *string {
	return s.DtsJobId
}

func (s *ModifyDtsJobAdvanceRequest) GetEtlOperatorColumnReference() *string {
	return s.EtlOperatorColumnReference
}

func (s *ModifyDtsJobAdvanceRequest) GetFileOssUrlObject() io.Reader {
	return s.FileOssUrlObject
}

func (s *ModifyDtsJobAdvanceRequest) GetFilterTableName() *string {
	return s.FilterTableName
}

func (s *ModifyDtsJobAdvanceRequest) GetModifyTypeEnum() *string {
	return s.ModifyTypeEnum
}

func (s *ModifyDtsJobAdvanceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyDtsJobAdvanceRequest) GetReserved() *string {
	return s.Reserved
}

func (s *ModifyDtsJobAdvanceRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ModifyDtsJobAdvanceRequest) GetStructureInitialization() *bool {
	return s.StructureInitialization
}

func (s *ModifyDtsJobAdvanceRequest) GetSynchronizationDirection() *string {
	return s.SynchronizationDirection
}

func (s *ModifyDtsJobAdvanceRequest) GetZeroEtlJob() *bool {
	return s.ZeroEtlJob
}

func (s *ModifyDtsJobAdvanceRequest) SetClientToken(v string) *ModifyDtsJobAdvanceRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyDtsJobAdvanceRequest) SetDataInitialization(v bool) *ModifyDtsJobAdvanceRequest {
	s.DataInitialization = &v
	return s
}

func (s *ModifyDtsJobAdvanceRequest) SetDataSynchronization(v bool) *ModifyDtsJobAdvanceRequest {
	s.DataSynchronization = &v
	return s
}

func (s *ModifyDtsJobAdvanceRequest) SetDbList(v map[string]interface{}) *ModifyDtsJobAdvanceRequest {
	s.DbList = v
	return s
}

func (s *ModifyDtsJobAdvanceRequest) SetDtsInstanceId(v string) *ModifyDtsJobAdvanceRequest {
	s.DtsInstanceId = &v
	return s
}

func (s *ModifyDtsJobAdvanceRequest) SetDtsJobId(v string) *ModifyDtsJobAdvanceRequest {
	s.DtsJobId = &v
	return s
}

func (s *ModifyDtsJobAdvanceRequest) SetEtlOperatorColumnReference(v string) *ModifyDtsJobAdvanceRequest {
	s.EtlOperatorColumnReference = &v
	return s
}

func (s *ModifyDtsJobAdvanceRequest) SetFileOssUrlObject(v io.Reader) *ModifyDtsJobAdvanceRequest {
	s.FileOssUrlObject = v
	return s
}

func (s *ModifyDtsJobAdvanceRequest) SetFilterTableName(v string) *ModifyDtsJobAdvanceRequest {
	s.FilterTableName = &v
	return s
}

func (s *ModifyDtsJobAdvanceRequest) SetModifyTypeEnum(v string) *ModifyDtsJobAdvanceRequest {
	s.ModifyTypeEnum = &v
	return s
}

func (s *ModifyDtsJobAdvanceRequest) SetRegionId(v string) *ModifyDtsJobAdvanceRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyDtsJobAdvanceRequest) SetReserved(v string) *ModifyDtsJobAdvanceRequest {
	s.Reserved = &v
	return s
}

func (s *ModifyDtsJobAdvanceRequest) SetResourceGroupId(v string) *ModifyDtsJobAdvanceRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ModifyDtsJobAdvanceRequest) SetStructureInitialization(v bool) *ModifyDtsJobAdvanceRequest {
	s.StructureInitialization = &v
	return s
}

func (s *ModifyDtsJobAdvanceRequest) SetSynchronizationDirection(v string) *ModifyDtsJobAdvanceRequest {
	s.SynchronizationDirection = &v
	return s
}

func (s *ModifyDtsJobAdvanceRequest) SetZeroEtlJob(v bool) *ModifyDtsJobAdvanceRequest {
	s.ZeroEtlJob = &v
	return s
}

func (s *ModifyDtsJobAdvanceRequest) Validate() error {
	return dara.Validate(s)
}
