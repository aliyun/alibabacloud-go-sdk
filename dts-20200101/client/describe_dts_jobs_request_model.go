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
	// The ID of the DTS dedicated cluster on which the task runs.
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
	// The environment tag of the DTS instance. Valid values:
	//
	// - **normal**
	//
	// - **online**
	//
	// example:
	//
	// normal
	DtsBisLabel *string `json:"DtsBisLabel,omitempty" xml:"DtsBisLabel,omitempty"`
	// The ID of the data migration, data synchronization, or change tracking instance.
	//
	// example:
	//
	// dtsi03e3zty16i****
	DtsInstanceId *string `json:"DtsInstanceId,omitempty" xml:"DtsInstanceId,omitempty"`
	// The ID of the data migration, data synchronization, or change tracking task.
	//
	// example:
	//
	// qa110wq5r93hb49
	DtsJobId *string `json:"DtsJobId,omitempty" xml:"DtsJobId,omitempty"`
	// The ID of the parent task.
	//
	// >  In most cases, you do not need to specify this parameter.
	//
	// example:
	//
	// pk13r731m****
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The ID of the source or target database instance corresponding to the request parameter **InstanceType**.
	//
	// example:
	//
	// rm-bp1966yuut4w3****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The type of the source or target database instance.
	//
	// example:
	//
	// RDS
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The type of the DTS task. Valid values:
	//
	// 	- **MIGRATION**: data migration. This is the default value.
	//
	// 	- **SYNC**: data synchronization.
	//
	// 	- **SUBSCRIBE**: change tracking.
	//
	// example:
	//
	// MIGRATION
	JobType *string `json:"JobType,omitempty" xml:"JobType,omitempty"`
	// The basis on which the returned DTS tasks are sorted. Valid values:
	//
	// 	- **CreateTime**: sorts the DTS tasks based on the points in time when the DTS tasks are created.
	//
	// 	- **FinishTime**: sorts the DTS tasks based on the points in time when the DTS tasks are complete.
	//
	// 	- **duLimit*	- sorts the DTS tasks based on the upper limits on DTS Units (DUs) that the DTS tasks can use. This option applies only to the DTS tasks that are run on a DTS dedicated cluster.
	//
	// >  You can also set the **OrderDirection*	- parameter to specify whether to sort the DTS tasks in ascending or descending order.
	//
	// example:
	//
	// CreateTime
	OrderColumn *string `json:"OrderColumn,omitempty" xml:"OrderColumn,omitempty"`
	// The order in which the returned DTS tasks are sorted. Valid values:
	//
	// 	- **ASC**: sorts the DTS tasks in ascending order. This is the default value.
	//
	// 	- **DESC**: sorts the DTS tasks in descending order.
	//
	// example:
	//
	// ASC
	OrderDirection *string `json:"OrderDirection,omitempty" xml:"OrderDirection,omitempty"`
	OwnerId        *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number. Pages start from page **1**. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values: **20**, **30**, **50**, and **100**. Default value: **20**.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The content of the query condition.
	//
	// >  You must set the **Type*	- parameter to specify the type of the query condition.
	//
	// example:
	//
	// dtspk3f13r731m****
	Params *string `json:"Params,omitempty" xml:"Params,omitempty"`
	// The ID of the region in which the DTS instance resides. For more information, see [List of supported regions](https://help.aliyun.com/document_detail/141033.html).
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// This parameter is deprecated.
	//
	// Valid values:
	//
	// 	- cn-hangzhou
	//
	// 	- cn-shanghai
	//
	// 	- cn-beijing
	//
	// 	- cn-guangzhou
	//
	// 	- cn-shenzhen
	//
	// 	- cn-chengdu
	//
	// 	- cn-heyuan
	//
	// 	- cn-hongkong
	//
	// 	- cn-qingdao
	//
	// 	- cn-zhangbei
	//
	// 	- cn-zhangjiakou
	//
	// 	- us-east-1
	//
	// 	- us-west-1
	//
	// 	- cn-hangzhou-finance
	//
	// 	- cn-shanghai-finance
	//
	// 	- cn-shanghai-finance-1
	//
	// 	- cn-shenzhen-finance
	//
	// 	- cn-shenzhen-finance-1
	//
	// 	- cn-beijing-finance-1
	//
	// 	- cn-huhehaote
	//
	// 	- cn-north-2-gov-1
	//
	// 	- eu-central-1
	//
	// 	- eu-west-1
	//
	// 	- me-central-1
	//
	// 	- me-east-1
	//
	// 	- ap-northeast-1
	//
	// 	- ap-northeast-2
	//
	// 	- ap-southeast-1
	//
	// 	- ap-southeast-2
	//
	// 	- ap-southeast-3
	//
	// 	- ap-southeast-5
	//
	// 	- ap-southeast-6
	//
	// 	- ap-southeast-7
	//
	// 	- cn-wulanchabu
	//
	// 	- cn-zhengzhou-jva
	//
	// 	- cn-wuhan-lr
	//
	// example:
	//
	// cn-hangzhou
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
	// The state of the DTS task.
	//
	// Valid values for a data migration task:
	//
	// 	- **NotStarted**: The task is not started.
	//
	// 	- **Prechecking**: The task is being prechecked.
	//
	// 	- **PrecheckFailed**: The task failed to pass the precheck.
	//
	// 	- **PreCheckPass**: The task passed the precheck.
	//
	// 	- **NotConfigured**: The task is not configured.
	//
	// 	- **Migrating**: The task is in progress.
	//
	// 	- **Suspending**: The task is paused.
	//
	// 	- **MigrationFailed**: The task failed.
	//
	// 	- **Finished**: The task is complete.
	//
	// 	- **Retrying**: The task is being retried.
	//
	// 	- **Upgrade**: The task is being upgraded.
	//
	// 	- **Locked**: The task is locked.
	//
	// 	- **Downgrade**: The task is being downgraded.
	//
	// Valid values for a data synchronization task:
	//
	// 	- **NotStarted**: The task is not started.
	//
	// 	- **Prechecking**: The task is being prechecked.
	//
	// 	- **PrecheckFailed**: The task failed to pass the precheck.
	//
	// 	- **PreCheckPass**: The task passed the precheck.
	//
	// 	- **NotConfigured**: The task is not configured.
	//
	// 	- **Initializing**: The task is being initialized.
	//
	// 	- **InitializeFailed**: Initialization failed.
	//
	// 	- **Synchronizing**: The task is in progress.
	//
	// 	- **Failed**: The task failed.
	//
	// 	- **Suspending**: The task is paused.
	//
	// 	- **Modifying**: The objects in the task are being modified.
	//
	// 	- **Finished**: The task is complete.
	//
	// 	- **Retrying**: The task is being retried.
	//
	// 	- **Upgrade**: The task is being upgraded.
	//
	// 	- **Locked**: The task is locked.
	//
	// 	- **Downgrade**: The task is being downgraded.
	//
	// Valid values for a change tracking task:
	//
	// 	- **NotConfigured**: The task is not configured.
	//
	// 	- **NotStarted**: The task is not started.
	//
	// 	- **Prechecking**: The task is being prechecked.
	//
	// 	- **PrecheckFailed**: The task failed to pass the precheck.
	//
	// 	- **PreCheckPass**: The task passed the precheck.
	//
	// 	- **Starting**: The task is being started.
	//
	// 	- **Normal**: The task is running as expected.
	//
	// 	- **Retrying**: The task is being retried.
	//
	// 	- **Abnormal**: The task is not running as expected.
	//
	// 	- **Upgrade**: The task is being upgraded.
	//
	// 	- **Locked**: The task is locked.
	//
	// 	- **Downgrade**: The task is being downgraded.
	//
	// example:
	//
	// Migrating
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tags of the DTS task to be queried. Specify tags in the JSON format.
	//
	// >  You can call the **ListTagResources*	- operation to query the tag key and tag value.
	//
	// example:
	//
	// [     {         \\"key\\": \\"testK\\",         \\"value\\": \\"testV\\"     }  ]
	Tags *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The type of the query condition. Valid values:
	//
	// 	- **instance**: queries DTS tasks based on the ID of a DTS instance.
	//
	// 	- **name**: queries DTS tasks based on the name of a DTS instance. Fuzzy match is supported.
	//
	// 	- **srcRds**: queries DTS tasks based on the ID of an ApsaraDB RDS instance. The ApsaraDB RDS instance is the source instance of a DTS task.
	//
	// 	- **rds**: queries DTS tasks based on the ID of an ApsaraDB RDS instance. The ApsaraDB RDS instance is the destination instance of a DTS task.
	//
	// >  You must set the **Params*	- parameter to specify the content of the query condition.
	//
	// example:
	//
	// instance
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// Specifies whether to skip the **DbObject*	- parameter in the response. The DbObject parameter specifies the objects of the data migration, data synchronization, or change tracking task. Valid values:
	//
	// - **true**: does not return **DbObject**.
	//
	// - **false**: returns **DbObject**. If you set this parameter to false, the response time is shortened.
	//
	// example:
	//
	// true
	WithoutDbList *bool `json:"WithoutDbList,omitempty" xml:"WithoutDbList,omitempty"`
	// Whether it is a seamless integration (Zero-ETL) task, the value can be:
	//
	// - **false**: No. - **true**: Yes.
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
