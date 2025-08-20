// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApsHiveWorkloadResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApsWorkload(v *DescribeApsHiveWorkloadResponseBodyApsWorkload) *DescribeApsHiveWorkloadResponseBody
	GetApsWorkload() *DescribeApsHiveWorkloadResponseBodyApsWorkload
	SetRequestId(v string) *DescribeApsHiveWorkloadResponseBody
	GetRequestId() *string
}

type DescribeApsHiveWorkloadResponseBody struct {
	// The queried job.
	//
	// example:
	//
	// -
	ApsWorkload *DescribeApsHiveWorkloadResponseBodyApsWorkload `json:"ApsWorkload,omitempty" xml:"ApsWorkload,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 86F92D26-B774-5FA1-8E53-******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeApsHiveWorkloadResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeApsHiveWorkloadResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeApsHiveWorkloadResponseBody) GetApsWorkload() *DescribeApsHiveWorkloadResponseBodyApsWorkload {
	return s.ApsWorkload
}

func (s *DescribeApsHiveWorkloadResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeApsHiveWorkloadResponseBody) SetApsWorkload(v *DescribeApsHiveWorkloadResponseBodyApsWorkload) *DescribeApsHiveWorkloadResponseBody {
	s.ApsWorkload = v
	return s
}

func (s *DescribeApsHiveWorkloadResponseBody) SetRequestId(v string) *DescribeApsHiveWorkloadResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeApsHiveWorkloadResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeApsHiveWorkloadResponseBodyApsWorkload struct {
	// The advanced configurations.
	//
	// example:
	//
	// test.adv.config=value
	AdvancedConfig *string `json:"AdvancedConfig,omitempty" xml:"AdvancedConfig,omitempty"`
	// The policy to handle tables with the same name in the destination cluster.
	//
	// example:
	//
	// Intercept
	ConflictStrategy *string `json:"ConflictStrategy,omitempty" xml:"ConflictStrategy,omitempty"`
	// The time when the workload was created.
	//
	// example:
	//
	// -
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The cluster ID.
	//
	// example:
	//
	// amv-*******
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The data source ID.
	//
	// example:
	//
	// 8
	DatasourceId *int64 `json:"DatasourceId,omitempty" xml:"DatasourceId,omitempty"`
	// The name of the data source.
	//
	// example:
	//
	// sls-******
	DatasourceName *string `json:"DatasourceName,omitempty" xml:"DatasourceName,omitempty"`
	// The ID of the E-MapReduce (EMR) cluster.
	//
	// example:
	//
	// -
	EmrClusterId *string `json:"EmrClusterId,omitempty" xml:"EmrClusterId,omitempty"`
	// The number of AnalyticDB compute units (ACUs) required for migration.
	//
	// example:
	//
	// 16
	FullComputeUnit *string `json:"FullComputeUnit,omitempty" xml:"FullComputeUnit,omitempty"`
	// The URL of the Hive Metastore.
	//
	// example:
	//
	// -
	MetaStoreUri *string `json:"MetaStoreUri,omitempty" xml:"MetaStoreUri,omitempty"`
	// The Object Storage Service (OSS) URL of the AnalyticDB for MySQL cluster data.
	//
	// example:
	//
	// oss://******
	OssLocation *string `json:"OssLocation,omitempty" xml:"OssLocation,omitempty"`
	// The number of tasks that are allowed in parallel.
	//
	// example:
	//
	// 2
	Parallelism *int64 `json:"Parallelism,omitempty" xml:"Parallelism,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group to which the SQL statement belongs.
	//
	// example:
	//
	// test
	ResourceGroup *string `json:"ResourceGroup,omitempty" xml:"ResourceGroup,omitempty"`
	// The status of the workload.
	//
	// example:
	//
	// COMPLETED
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The expression that manually matches the source database table whitelist.
	//
	// example:
	//
	// abc
	SyncAllowExpression *string `json:"SyncAllowExpression,omitempty" xml:"SyncAllowExpression,omitempty"`
	// Manually match the blacklist expressions for source database tables.
	//
	// example:
	//
	// def
	SyncDenyExpression *string `json:"SyncDenyExpression,omitempty" xml:"SyncDenyExpression,omitempty"`
	// The destination type.
	//
	// example:
	//
	// OSS
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
	// The name of the vSwitch.
	//
	// example:
	//
	// vsw-******
	Vswitch *string `json:"Vswitch,omitempty" xml:"Vswitch,omitempty"`
	// The job ID.
	//
	// example:
	//
	// aps-******
	WorkloadId *string `json:"WorkloadId,omitempty" xml:"WorkloadId,omitempty"`
	// The name of the workload.
	//
	// example:
	//
	// test
	WorkloadName *string `json:"WorkloadName,omitempty" xml:"WorkloadName,omitempty"`
	// The name of the workload.
	//
	// example:
	//
	// test
	WorkloadTypeName *string `json:"WorkloadTypeName,omitempty" xml:"WorkloadTypeName,omitempty"`
}

func (s DescribeApsHiveWorkloadResponseBodyApsWorkload) String() string {
	return dara.Prettify(s)
}

func (s DescribeApsHiveWorkloadResponseBodyApsWorkload) GoString() string {
	return s.String()
}

func (s *DescribeApsHiveWorkloadResponseBodyApsWorkload) GetAdvancedConfig() *string {
	return s.AdvancedConfig
}

func (s *DescribeApsHiveWorkloadResponseBodyApsWorkload) GetConflictStrategy() *string {
	return s.ConflictStrategy
}

func (s *DescribeApsHiveWorkloadResponseBodyApsWorkload) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeApsHiveWorkloadResponseBodyApsWorkload) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeApsHiveWorkloadResponseBodyApsWorkload) GetDatasourceId() *int64 {
	return s.DatasourceId
}

func (s *DescribeApsHiveWorkloadResponseBodyApsWorkload) GetDatasourceName() *string {
	return s.DatasourceName
}

func (s *DescribeApsHiveWorkloadResponseBodyApsWorkload) GetEmrClusterId() *string {
	return s.EmrClusterId
}

func (s *DescribeApsHiveWorkloadResponseBodyApsWorkload) GetFullComputeUnit() *string {
	return s.FullComputeUnit
}

func (s *DescribeApsHiveWorkloadResponseBodyApsWorkload) GetMetaStoreUri() *string {
	return s.MetaStoreUri
}

func (s *DescribeApsHiveWorkloadResponseBodyApsWorkload) GetOssLocation() *string {
	return s.OssLocation
}

func (s *DescribeApsHiveWorkloadResponseBodyApsWorkload) GetParallelism() *int64 {
	return s.Parallelism
}

func (s *DescribeApsHiveWorkloadResponseBodyApsWorkload) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeApsHiveWorkloadResponseBodyApsWorkload) GetResourceGroup() *string {
	return s.ResourceGroup
}

func (s *DescribeApsHiveWorkloadResponseBodyApsWorkload) GetState() *string {
	return s.State
}

func (s *DescribeApsHiveWorkloadResponseBodyApsWorkload) GetSyncAllowExpression() *string {
	return s.SyncAllowExpression
}

func (s *DescribeApsHiveWorkloadResponseBodyApsWorkload) GetSyncDenyExpression() *string {
	return s.SyncDenyExpression
}

func (s *DescribeApsHiveWorkloadResponseBodyApsWorkload) GetTargetType() *string {
	return s.TargetType
}

func (s *DescribeApsHiveWorkloadResponseBodyApsWorkload) GetVswitch() *string {
	return s.Vswitch
}

func (s *DescribeApsHiveWorkloadResponseBodyApsWorkload) GetWorkloadId() *string {
	return s.WorkloadId
}

func (s *DescribeApsHiveWorkloadResponseBodyApsWorkload) GetWorkloadName() *string {
	return s.WorkloadName
}

func (s *DescribeApsHiveWorkloadResponseBodyApsWorkload) GetWorkloadTypeName() *string {
	return s.WorkloadTypeName
}

func (s *DescribeApsHiveWorkloadResponseBodyApsWorkload) SetAdvancedConfig(v string) *DescribeApsHiveWorkloadResponseBodyApsWorkload {
	s.AdvancedConfig = &v
	return s
}

func (s *DescribeApsHiveWorkloadResponseBodyApsWorkload) SetConflictStrategy(v string) *DescribeApsHiveWorkloadResponseBodyApsWorkload {
	s.ConflictStrategy = &v
	return s
}

func (s *DescribeApsHiveWorkloadResponseBodyApsWorkload) SetCreateTime(v string) *DescribeApsHiveWorkloadResponseBodyApsWorkload {
	s.CreateTime = &v
	return s
}

func (s *DescribeApsHiveWorkloadResponseBodyApsWorkload) SetDBClusterId(v string) *DescribeApsHiveWorkloadResponseBodyApsWorkload {
	s.DBClusterId = &v
	return s
}

func (s *DescribeApsHiveWorkloadResponseBodyApsWorkload) SetDatasourceId(v int64) *DescribeApsHiveWorkloadResponseBodyApsWorkload {
	s.DatasourceId = &v
	return s
}

func (s *DescribeApsHiveWorkloadResponseBodyApsWorkload) SetDatasourceName(v string) *DescribeApsHiveWorkloadResponseBodyApsWorkload {
	s.DatasourceName = &v
	return s
}

func (s *DescribeApsHiveWorkloadResponseBodyApsWorkload) SetEmrClusterId(v string) *DescribeApsHiveWorkloadResponseBodyApsWorkload {
	s.EmrClusterId = &v
	return s
}

func (s *DescribeApsHiveWorkloadResponseBodyApsWorkload) SetFullComputeUnit(v string) *DescribeApsHiveWorkloadResponseBodyApsWorkload {
	s.FullComputeUnit = &v
	return s
}

func (s *DescribeApsHiveWorkloadResponseBodyApsWorkload) SetMetaStoreUri(v string) *DescribeApsHiveWorkloadResponseBodyApsWorkload {
	s.MetaStoreUri = &v
	return s
}

func (s *DescribeApsHiveWorkloadResponseBodyApsWorkload) SetOssLocation(v string) *DescribeApsHiveWorkloadResponseBodyApsWorkload {
	s.OssLocation = &v
	return s
}

func (s *DescribeApsHiveWorkloadResponseBodyApsWorkload) SetParallelism(v int64) *DescribeApsHiveWorkloadResponseBodyApsWorkload {
	s.Parallelism = &v
	return s
}

func (s *DescribeApsHiveWorkloadResponseBodyApsWorkload) SetRegionId(v string) *DescribeApsHiveWorkloadResponseBodyApsWorkload {
	s.RegionId = &v
	return s
}

func (s *DescribeApsHiveWorkloadResponseBodyApsWorkload) SetResourceGroup(v string) *DescribeApsHiveWorkloadResponseBodyApsWorkload {
	s.ResourceGroup = &v
	return s
}

func (s *DescribeApsHiveWorkloadResponseBodyApsWorkload) SetState(v string) *DescribeApsHiveWorkloadResponseBodyApsWorkload {
	s.State = &v
	return s
}

func (s *DescribeApsHiveWorkloadResponseBodyApsWorkload) SetSyncAllowExpression(v string) *DescribeApsHiveWorkloadResponseBodyApsWorkload {
	s.SyncAllowExpression = &v
	return s
}

func (s *DescribeApsHiveWorkloadResponseBodyApsWorkload) SetSyncDenyExpression(v string) *DescribeApsHiveWorkloadResponseBodyApsWorkload {
	s.SyncDenyExpression = &v
	return s
}

func (s *DescribeApsHiveWorkloadResponseBodyApsWorkload) SetTargetType(v string) *DescribeApsHiveWorkloadResponseBodyApsWorkload {
	s.TargetType = &v
	return s
}

func (s *DescribeApsHiveWorkloadResponseBodyApsWorkload) SetVswitch(v string) *DescribeApsHiveWorkloadResponseBodyApsWorkload {
	s.Vswitch = &v
	return s
}

func (s *DescribeApsHiveWorkloadResponseBodyApsWorkload) SetWorkloadId(v string) *DescribeApsHiveWorkloadResponseBodyApsWorkload {
	s.WorkloadId = &v
	return s
}

func (s *DescribeApsHiveWorkloadResponseBodyApsWorkload) SetWorkloadName(v string) *DescribeApsHiveWorkloadResponseBodyApsWorkload {
	s.WorkloadName = &v
	return s
}

func (s *DescribeApsHiveWorkloadResponseBodyApsWorkload) SetWorkloadTypeName(v string) *DescribeApsHiveWorkloadResponseBodyApsWorkload {
	s.WorkloadTypeName = &v
	return s
}

func (s *DescribeApsHiveWorkloadResponseBodyApsWorkload) Validate() error {
	return dara.Validate(s)
}
