// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDtsJobsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDedicatedClusterId(v string) *DescribeDtsJobsRequest
	GetDedicatedClusterId() *string
	SetDestProductType(v string) *DescribeDtsJobsRequest
	GetDestProductType() *string
	SetDtsBisLabel(v string) *DescribeDtsJobsRequest
	GetDtsBisLabel() *string
	SetDtsInstanceId(v string) *DescribeDtsJobsRequest
	GetDtsInstanceId() *string
	SetDtsJobId(v string) *DescribeDtsJobsRequest
	GetDtsJobId() *string
	SetGroupId(v string) *DescribeDtsJobsRequest
	GetGroupId() *string
	SetInstanceId(v string) *DescribeDtsJobsRequest
	GetInstanceId() *string
	SetInstanceType(v string) *DescribeDtsJobsRequest
	GetInstanceType() *string
	SetJobType(v string) *DescribeDtsJobsRequest
	GetJobType() *string
	SetOrderColumn(v string) *DescribeDtsJobsRequest
	GetOrderColumn() *string
	SetOrderDirection(v string) *DescribeDtsJobsRequest
	GetOrderDirection() *string
	SetOwnerId(v string) *DescribeDtsJobsRequest
	GetOwnerId() *string
	SetPageNumber(v int32) *DescribeDtsJobsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeDtsJobsRequest
	GetPageSize() *int32
	SetParams(v string) *DescribeDtsJobsRequest
	GetParams() *string
	SetRegion(v string) *DescribeDtsJobsRequest
	GetRegion() *string
	SetRegionId(v string) *DescribeDtsJobsRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeDtsJobsRequest
	GetResourceGroupId() *string
	SetSrcProductType(v string) *DescribeDtsJobsRequest
	GetSrcProductType() *string
	SetStatus(v string) *DescribeDtsJobsRequest
	GetStatus() *string
	SetTags(v string) *DescribeDtsJobsRequest
	GetTags() *string
	SetType(v string) *DescribeDtsJobsRequest
	GetType() *string
	SetWithoutDbList(v bool) *DescribeDtsJobsRequest
	GetWithoutDbList() *bool
	SetZeroEtlJob(v bool) *DescribeDtsJobsRequest
	GetZeroEtlJob() *bool
}

type DescribeDtsJobsRequest struct {
	// The ID of the DTS dedicated cluster.
	//
	// example:
	//
	// dtscluster_atyl3b5214uk***
	DedicatedClusterId *string `json:"DedicatedClusterId,omitempty" xml:"DedicatedClusterId,omitempty"`
	// The type of the source database instance.
	//
	// example:
	//
	// RDS
	DestProductType *string `json:"DestProductType,omitempty" xml:"DestProductType,omitempty"`
	// The environment label of the DTS instance. Valid values:
	//
	// - **normal**: normal
	//
	// - **online**: online
	//
	// example:
	//
	// normal
	DtsBisLabel *string `json:"DtsBisLabel,omitempty" xml:"DtsBisLabel,omitempty"`
	// The ID of the data migration, data synchronization, or change tracking instance.
	//
	// > Separate multiple instance IDs with commas (,). Make sure that the **JobType*	- parameter is set as expected.
	//
	// example:
	//
	// dtsi03e3zty16i****
	DtsInstanceId *string `json:"DtsInstanceId,omitempty" xml:"DtsInstanceId,omitempty"`
	// The ID of the data migration, data synchronization, or change tracking task.
	//
	// > Separate multiple task IDs with commas (,). Make sure that the **JobType*	- parameter is set as expected.
	//
	// example:
	//
	// qa110wq5r93hb49
	DtsJobId *string `json:"DtsJobId,omitempty" xml:"DtsJobId,omitempty"`
	// The DTS task ID.
	//
	// > In most cases, you do not need to set this parameter.
	//
	// example:
	//
	// pk13r731m****
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The ID of the source or destination database instance that corresponds to the **InstanceType*	- request parameter.
	//
	// example:
	//
	// rm-bp1966yuut4w3****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The type of the source or destination database instance.
	//
	// example:
	//
	// RDS
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The task type of the DTS instance. Valid values:
	//
	// - **MIGRATION**: data migration (default).
	//
	// - **SYNC**: data synchronization.
	//
	// - **SUBSCRIBE**: change tracking.
	//
	// example:
	//
	// MIGRATION
	JobType *string `json:"JobType,omitempty" xml:"JobType,omitempty"`
	// The sort criterion when the response contains multiple DTS instances. Valid values:
	//
	// - **CreateTime**: sorts by task creation time.
	//
	// - **FinishTime**: sorts by task completion time.
	//
	// - **duLimit*	- (dedicated cluster tasks): sorts by the upper limit of DU usage for DTS tasks. This value is supported only for dedicated clusters.
	//
	// > You can also specify **OrderDirection*	- to set the sort order to ascending or descending.
	//
	// example:
	//
	// CreateTime
	OrderColumn *string `json:"OrderColumn,omitempty" xml:"OrderColumn,omitempty"`
	// The sort order of instances. Valid values:
	//
	// - **ASC**: ascending order. This is the default value.
	//
	// - **DESC**: descending order.
	//
	// example:
	//
	// ASC
	OrderDirection *string `json:"OrderDirection,omitempty" xml:"OrderDirection,omitempty"`
	OwnerId        *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number. The value must be a positive integer that does not exceed the maximum value of the Integer data type. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of records per page. Valid values: **10**, **20**, and **30**. Default value: **20**. Maximum value: **30**.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The specific content of the query condition.
	//
	// > Specify **Type*	- in advance to define the query condition.
	//
	// example:
	//
	// dtspk3f13r731m****
	Params *string `json:"Params,omitempty" xml:"Params,omitempty"`
	// The region in which the DTS instance resides. For more information, see [List of supported regions](https://help.aliyun.com/document_detail/141033.html).
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// Deprecated parameter.
	//
	// example:
	//
	// 无
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmzawhxxc****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The type of the destination database instance.
	//
	// example:
	//
	// RDS
	SrcProductType *string `json:"SrcProductType,omitempty" xml:"SrcProductType,omitempty"`
	// The instance status of the DTS instance. Valid values:
	//
	// Data migration task statuses:
	//
	// - **NotStarted**: not started.
	//
	// - **Prechecking**: running a precheck.
	//
	// - **PrecheckFailed**: precheck failed.
	//
	// - **PreCheckPass**: precheck passed.
	//
	// - **NotConfigured**: not configured.
	//
	// - **Migrating**: migrating.
	//
	// - **Suspending**: paused.
	//
	// - **MigrationFailed**: migration failed.
	//
	// - **Finished**: completed.
	//
	// - **Retrying**: retrying.
	//
	// - **Upgrade**: upgrading.
	//
	// - **Locked**: locked.
	//
	// - **Downgrade**: downgrading.
	//
	// Data synchronization task statuses:
	//
	// - **NotStarted**: not started.
	//
	// - **Prechecking**: running a precheck.
	//
	// - **PrecheckFailed**: precheck failed.
	//
	// - **PreCheckPass**: precheck passed.
	//
	// - **NotConfigured**: not configured.
	//
	// - **Initializing**: performing initial synchronization.
	//
	// - **InitializeFailed**: initial synchronization failed.
	//
	// - **Synchronizing**: synchronizing.
	//
	// - **Failed**: synchronization failed.
	//
	// - **Suspending**: paused.
	//
	// - **Modifying**: modifying synchronization objects.
	//
	// - **Finished**: completed.
	//
	// - **Retrying**: retrying.
	//
	// - **Upgrade**: upgrading.
	//
	// - **Locked**: locked.
	//
	// - **Downgrade**: downgrading.
	//
	// Change tracking task statuses:
	//
	// - **NotConfigured**: not configured.
	//
	// - **NotStarted**: not started.
	//
	// - **Prechecking**: running a precheck.
	//
	// - **PrecheckFailed**: precheck failed.
	//
	// - **PreCheckPass**: precheck passed.
	//
	// - **Starting**: starting.
	//
	// - **Normal**: normal.
	//
	// - **Retrying**: retrying.
	//
	// - **Abnormal**: abnormal.
	//
	// - **Upgrade**: upgrading.
	//
	// - **Locked**: locked.
	//
	// - **Downgrade**: downgrading.
	//
	// example:
	//
	// Migrating
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tag-based search condition in JSON format.
	//
	// > You can call the **ListTagResources*	- operation to query tag keys and values.
	//
	// example:
	//
	// [     {         \\"key\\": \\"testK\\",         \\"value\\": \\"testV\\"     }  ]
	Tags *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The conditional query parameter. Valid values:
	//
	// - **instance**: queries by DTS instance ID.
	//
	// - **name**: queries by DTS instance name. Fuzzy match is supported.
	//
	// - **srcRds**: queries by the ID of the source instance (ApsaraDB RDS).
	//
	// - **rds**: queries by the ID of the destination instance (ApsaraDB RDS).
	//
	// > Specify the **Params*	- parameter to provide the specific content of the query condition.
	//
	// example:
	//
	// instance
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// Specifies whether to exclude task objects from the response (not return the **DbObject*	- parameter). Valid values:
	//
	// - **true**: excludes **DbObject*	- from the response.
	//
	// - **false**: includes **DbObject*	- in the response, which can improve the response speed.
	//
	// example:
	//
	// true
	WithoutDbList *bool `json:"WithoutDbList,omitempty" xml:"WithoutDbList,omitempty"`
	// Specifies whether the node is a seamless integration (Zero-ETL) node. Valid values:
	//
	// - **false**: No.
	//
	// - **true**: Yes.
	//
	// example:
	//
	// false
	ZeroEtlJob *bool `json:"ZeroEtlJob,omitempty" xml:"ZeroEtlJob,omitempty"`
}

func (s DescribeDtsJobsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDtsJobsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDtsJobsRequest) GetDedicatedClusterId() *string {
	return s.DedicatedClusterId
}

func (s *DescribeDtsJobsRequest) GetDestProductType() *string {
	return s.DestProductType
}

func (s *DescribeDtsJobsRequest) GetDtsBisLabel() *string {
	return s.DtsBisLabel
}

func (s *DescribeDtsJobsRequest) GetDtsInstanceId() *string {
	return s.DtsInstanceId
}

func (s *DescribeDtsJobsRequest) GetDtsJobId() *string {
	return s.DtsJobId
}

func (s *DescribeDtsJobsRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *DescribeDtsJobsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeDtsJobsRequest) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribeDtsJobsRequest) GetJobType() *string {
	return s.JobType
}

func (s *DescribeDtsJobsRequest) GetOrderColumn() *string {
	return s.OrderColumn
}

func (s *DescribeDtsJobsRequest) GetOrderDirection() *string {
	return s.OrderDirection
}

func (s *DescribeDtsJobsRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *DescribeDtsJobsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDtsJobsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDtsJobsRequest) GetParams() *string {
	return s.Params
}

func (s *DescribeDtsJobsRequest) GetRegion() *string {
	return s.Region
}

func (s *DescribeDtsJobsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDtsJobsRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeDtsJobsRequest) GetSrcProductType() *string {
	return s.SrcProductType
}

func (s *DescribeDtsJobsRequest) GetStatus() *string {
	return s.Status
}

func (s *DescribeDtsJobsRequest) GetTags() *string {
	return s.Tags
}

func (s *DescribeDtsJobsRequest) GetType() *string {
	return s.Type
}

func (s *DescribeDtsJobsRequest) GetWithoutDbList() *bool {
	return s.WithoutDbList
}

func (s *DescribeDtsJobsRequest) GetZeroEtlJob() *bool {
	return s.ZeroEtlJob
}

func (s *DescribeDtsJobsRequest) SetDedicatedClusterId(v string) *DescribeDtsJobsRequest {
	s.DedicatedClusterId = &v
	return s
}

func (s *DescribeDtsJobsRequest) SetDestProductType(v string) *DescribeDtsJobsRequest {
	s.DestProductType = &v
	return s
}

func (s *DescribeDtsJobsRequest) SetDtsBisLabel(v string) *DescribeDtsJobsRequest {
	s.DtsBisLabel = &v
	return s
}

func (s *DescribeDtsJobsRequest) SetDtsInstanceId(v string) *DescribeDtsJobsRequest {
	s.DtsInstanceId = &v
	return s
}

func (s *DescribeDtsJobsRequest) SetDtsJobId(v string) *DescribeDtsJobsRequest {
	s.DtsJobId = &v
	return s
}

func (s *DescribeDtsJobsRequest) SetGroupId(v string) *DescribeDtsJobsRequest {
	s.GroupId = &v
	return s
}

func (s *DescribeDtsJobsRequest) SetInstanceId(v string) *DescribeDtsJobsRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeDtsJobsRequest) SetInstanceType(v string) *DescribeDtsJobsRequest {
	s.InstanceType = &v
	return s
}

func (s *DescribeDtsJobsRequest) SetJobType(v string) *DescribeDtsJobsRequest {
	s.JobType = &v
	return s
}

func (s *DescribeDtsJobsRequest) SetOrderColumn(v string) *DescribeDtsJobsRequest {
	s.OrderColumn = &v
	return s
}

func (s *DescribeDtsJobsRequest) SetOrderDirection(v string) *DescribeDtsJobsRequest {
	s.OrderDirection = &v
	return s
}

func (s *DescribeDtsJobsRequest) SetOwnerId(v string) *DescribeDtsJobsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDtsJobsRequest) SetPageNumber(v int32) *DescribeDtsJobsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDtsJobsRequest) SetPageSize(v int32) *DescribeDtsJobsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDtsJobsRequest) SetParams(v string) *DescribeDtsJobsRequest {
	s.Params = &v
	return s
}

func (s *DescribeDtsJobsRequest) SetRegion(v string) *DescribeDtsJobsRequest {
	s.Region = &v
	return s
}

func (s *DescribeDtsJobsRequest) SetRegionId(v string) *DescribeDtsJobsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDtsJobsRequest) SetResourceGroupId(v string) *DescribeDtsJobsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDtsJobsRequest) SetSrcProductType(v string) *DescribeDtsJobsRequest {
	s.SrcProductType = &v
	return s
}

func (s *DescribeDtsJobsRequest) SetStatus(v string) *DescribeDtsJobsRequest {
	s.Status = &v
	return s
}

func (s *DescribeDtsJobsRequest) SetTags(v string) *DescribeDtsJobsRequest {
	s.Tags = &v
	return s
}

func (s *DescribeDtsJobsRequest) SetType(v string) *DescribeDtsJobsRequest {
	s.Type = &v
	return s
}

func (s *DescribeDtsJobsRequest) SetWithoutDbList(v bool) *DescribeDtsJobsRequest {
	s.WithoutDbList = &v
	return s
}

func (s *DescribeDtsJobsRequest) SetZeroEtlJob(v bool) *DescribeDtsJobsRequest {
	s.ZeroEtlJob = &v
	return s
}

func (s *DescribeDtsJobsRequest) Validate() error {
	return dara.Validate(s)
}
