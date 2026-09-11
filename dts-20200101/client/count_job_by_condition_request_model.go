// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCountJobByConditionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDestDbType(v string) *CountJobByConditionRequest
	GetDestDbType() *string
	SetGroupId(v string) *CountJobByConditionRequest
	GetGroupId() *string
	SetJobType(v string) *CountJobByConditionRequest
	GetJobType() *string
	SetParams(v string) *CountJobByConditionRequest
	GetParams() *string
	SetRegion(v string) *CountJobByConditionRequest
	GetRegion() *string
	SetRegionId(v string) *CountJobByConditionRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CountJobByConditionRequest
	GetResourceGroupId() *string
	SetSrcDbType(v string) *CountJobByConditionRequest
	GetSrcDbType() *string
	SetStatus(v string) *CountJobByConditionRequest
	GetStatus() *string
	SetType(v string) *CountJobByConditionRequest
	GetType() *string
}

type CountJobByConditionRequest struct {
	// The type of the destination database.
	//
	// example:
	//
	// MongoDB
	DestDbType *string `json:"DestDbType,omitempty" xml:"DestDbType,omitempty"`
	// The parent task ID of a distributed synchronization task.
	//
	// example:
	//
	// pk13r731m****
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The task type. Valid values:
	//
	// - **MIGRATION**: data migration.
	//
	// - **SYNC**: data synchronization.
	//
	// - **SUBSCRIBE**: change tracking.
	//
	// example:
	//
	// SYNC
	JobType *string `json:"JobType,omitempty" xml:"JobType,omitempty"`
	// The query value that corresponds to JobType.
	//
	// example:
	//
	// dtspk3f13r731m****
	Params *string `json:"Params,omitempty" xml:"Params,omitempty"`
	// The region ID used as a filter condition. For more information, see [Supported regions](https://help.aliyun.com/document_detail/141033.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The ID of the region where the DTS instance resides. For more information, see [Supported regions](https://help.aliyun.com/document_detail/141033.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID. This is a global parameter and does not need to be passed for this API operation.
	//
	// example:
	//
	// 资源组ID，全局参数，当前API无需传入。
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The type of the source database.
	//
	// example:
	//
	// MongoDB
	SrcDbType *string `json:"SrcDbType,omitempty" xml:"SrcDbType,omitempty"`
	// The instance status of the DTS instance. Valid values:
	//
	// Data migration node statuses:
	//
	// - **NotStarted**: not started.
	//
	// - **Prechecking**: running a dry run.
	//
	// - **PrecheckFailed**: dry run failed.
	//
	// - **PreCheckPass**: dry run passed.
	//
	// - **NotConfigured**: not configured.
	//
	// - **Migrating**: migrating.
	//
	// - **Suspending**: paused.
	//
	// - **MigrationFailed**: migration failed.
	//
	// - **Finished**: finished.
	//
	// - **Retrying**: retrying.
	//
	// - **Upgrade**: upgrading.
	//
	// - **Locked**: locked.
	//
	// - **Downgrade**: downgrading.
	//
	// Data synchronization node statuses:
	//
	// - **NotStarted**: not started.
	//
	// - **Prechecking**: running a dry run.
	//
	// - **PrecheckFailed**: dry run failed.
	//
	// - **PreCheckPass**: dry run passed.
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
	// - **Modifying**: modifying sub-objects.
	//
	// - **Finished**: finished.
	//
	// - **Retrying**: retrying.
	//
	// - **Upgrade**: upgrading.
	//
	// - **Locked**: locked.
	//
	// - **Downgrade**: downgrading.
	//
	// Subscribe node statuses:
	//
	// - **NotConfigured**: not configured.
	//
	// - **NotStarted**: not started.
	//
	// - **Prechecking**: running a dry run.
	//
	// - **PrecheckFailed**: dry run failed.
	//
	// - **PreCheckPass**: dry run passed.
	//
	// - **Starting**: starting.
	//
	// - **Normal**: Normal.
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
	// Finished
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The query type. Valid values:
	//
	// - **name**: queries by job name.
	//
	// - **rds**: queries by destination instance ID.
	//
	// - **instance**: queries by DTS instance ID.
	//
	// - **srcRds**: queries by source instance ID.
	//
	// > This parameter corresponds to the **JobType*	- parameter.
	//
	// example:
	//
	// name/instance
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CountJobByConditionRequest) String() string {
	return dara.Prettify(s)
}

func (s CountJobByConditionRequest) GoString() string {
	return s.String()
}

func (s *CountJobByConditionRequest) GetDestDbType() *string {
	return s.DestDbType
}

func (s *CountJobByConditionRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *CountJobByConditionRequest) GetJobType() *string {
	return s.JobType
}

func (s *CountJobByConditionRequest) GetParams() *string {
	return s.Params
}

func (s *CountJobByConditionRequest) GetRegion() *string {
	return s.Region
}

func (s *CountJobByConditionRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CountJobByConditionRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CountJobByConditionRequest) GetSrcDbType() *string {
	return s.SrcDbType
}

func (s *CountJobByConditionRequest) GetStatus() *string {
	return s.Status
}

func (s *CountJobByConditionRequest) GetType() *string {
	return s.Type
}

func (s *CountJobByConditionRequest) SetDestDbType(v string) *CountJobByConditionRequest {
	s.DestDbType = &v
	return s
}

func (s *CountJobByConditionRequest) SetGroupId(v string) *CountJobByConditionRequest {
	s.GroupId = &v
	return s
}

func (s *CountJobByConditionRequest) SetJobType(v string) *CountJobByConditionRequest {
	s.JobType = &v
	return s
}

func (s *CountJobByConditionRequest) SetParams(v string) *CountJobByConditionRequest {
	s.Params = &v
	return s
}

func (s *CountJobByConditionRequest) SetRegion(v string) *CountJobByConditionRequest {
	s.Region = &v
	return s
}

func (s *CountJobByConditionRequest) SetRegionId(v string) *CountJobByConditionRequest {
	s.RegionId = &v
	return s
}

func (s *CountJobByConditionRequest) SetResourceGroupId(v string) *CountJobByConditionRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CountJobByConditionRequest) SetSrcDbType(v string) *CountJobByConditionRequest {
	s.SrcDbType = &v
	return s
}

func (s *CountJobByConditionRequest) SetStatus(v string) *CountJobByConditionRequest {
	s.Status = &v
	return s
}

func (s *CountJobByConditionRequest) SetType(v string) *CountJobByConditionRequest {
	s.Type = &v
	return s
}

func (s *CountJobByConditionRequest) Validate() error {
	return dara.Validate(s)
}
