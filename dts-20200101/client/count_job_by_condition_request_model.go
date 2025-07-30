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
	// The ID of the DTS task.
	//
	// example:
	//
	// pk13r731m****
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The type of the DTS task. Valid values:
	//
	// 	- **MIGRATION**: data migration task
	//
	// 	- **SYNC**: data synchronization task
	//
	// 	- **SUBSCRIBE**: change tracking task
	//
	// example:
	//
	// SYNC
	JobType *string `json:"JobType,omitempty" xml:"JobType,omitempty"`
	// The content of the query condition, which corresponds to the value of the JobType parameter.
	//
	// example:
	//
	// dtspk3f13r731m****
	Params *string `json:"Params,omitempty" xml:"Params,omitempty"`
	// One of the query conditions. The ID of the region. For more information, see [Supported regions](https://help.aliyun.com/document_detail/141033.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The ID of the region in which the DTS instance resides. For more information, see [Supported regions](https://help.aliyun.com/document_detail/141033.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Resource group ID, global parameter that does not need to be passed in by the current API.
	//
	// example:
	//
	// Resource group ID, global parameter that does not need to be passed in by the current API.
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The type of the source database.
	//
	// example:
	//
	// MongoDB
	SrcDbType *string `json:"SrcDbType,omitempty" xml:"SrcDbType,omitempty"`
	// The status of the DTS task.
	//
	// Valid values for a data migration task:
	//
	// 	- **NotStarted**: The task is not started.
	//
	// 	- **Prechecking**: The task is in precheck.
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
	// 	- **MigrationFailed**: The task failed to migrate data.
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
	// 	- **Prechecking**: The task is in precheck.
	//
	// 	- **PrecheckFailed**: The task failed to pass the precheck.
	//
	// 	- **PreCheckPass**: The task passed the precheck.
	//
	// 	- **NotConfigured**: The task is not configured.
	//
	// 	- **Initializing**: The task is performing initial synchronization.
	//
	// 	- **InitializeFailed**: Initial synchronization failed.
	//
	// 	- **Synchronizing**: The task is in progress.
	//
	// 	- **Failed**: The task failed to synchronize data.
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
	// 	- **Prechecking**: The task is in precheck.
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
	// Finished
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The content of the query condition. Valid values:
	//
	// 	- **name**: the name of the task
	//
	// 	- **rds**: the ID of the destination instance
	//
	// 	- **instance**: the ID of the Data Transmission Service (DTS) instance
	//
	// 	- **srcRds**: the ID of the source instance
	//
	// > The value of this parameter corresponds to the value of the **JobType*	- parameter.
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
