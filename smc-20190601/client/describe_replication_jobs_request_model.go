// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeReplicationJobsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBusinessStatus(v string) *DescribeReplicationJobsRequest
	GetBusinessStatus() *string
	SetInstanceId(v []*string) *DescribeReplicationJobsRequest
	GetInstanceId() []*string
	SetJobId(v []*string) *DescribeReplicationJobsRequest
	GetJobId() []*string
	SetJobType(v int32) *DescribeReplicationJobsRequest
	GetJobType() *int32
	SetName(v string) *DescribeReplicationJobsRequest
	GetName() *string
	SetOwnerId(v int64) *DescribeReplicationJobsRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeReplicationJobsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeReplicationJobsRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeReplicationJobsRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeReplicationJobsRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *DescribeReplicationJobsRequest
	GetResourceOwnerAccount() *string
	SetSourceId(v []*string) *DescribeReplicationJobsRequest
	GetSourceId() []*string
	SetStatus(v string) *DescribeReplicationJobsRequest
	GetStatus() *string
	SetTag(v []*DescribeReplicationJobsRequestTag) *DescribeReplicationJobsRequest
	GetTag() []*DescribeReplicationJobsRequestTag
}

type DescribeReplicationJobsRequest struct {
	// The business status of the migration job. Valid values:
	//
	// 	- Preparing: The migration is being prepared.
	//
	// 	- Syncing: Data is being synchronized.
	//
	// 	- Processing: The migration is in progress.
	//
	// 	- Cleaning: Intermediate resources are being released.
	//
	// example:
	//
	// Preparing
	BusinessStatus *string `json:"BusinessStatus,omitempty" xml:"BusinessStatus,omitempty"`
	// The IDs of the destination Elastic Compute Service (ECS) instances.
	InstanceId []*string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" type:"Repeated"`
	// The IDs of the migration jobs. You can specify a maximum of 50 IDs.
	//
	// example:
	//
	// j-bp19vlwm0tyigbmj****
	JobId []*string `json:"JobId,omitempty" xml:"JobId,omitempty" type:"Repeated"`
	// The type of the migration job. Valid values:
	//
	// 	- 0: server migration.
	//
	// 	- 1: operating system migration.
	//
	// 	- 2: cross-zone migration.
	//
	// 	- 3: agentless migration for a VMware VM.
	//
	// example:
	//
	// 0
	JobType *int32 `json:"JobType,omitempty" xml:"JobType,omitempty"`
	// The name of the migration job.
	//
	// example:
	//
	// testMigrationTaskName
	Name    *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number. Minimum value: 1.
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values: 1 to 50.
	//
	// Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the Alibaba Cloud region to which you want to migrate the source server.
	//
	// For example, if you want to migrate a source server to the China (Hangzhou) region, set this parameter to `cn-hangzhou`. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation to query the latest regions.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmw3ty5y7****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	// The IDs of the source servers. You can specify a maximum of 50 IDs.
	//
	// example:
	//
	// s-bp1e2fsl57knvuug****
	SourceId []*string `json:"SourceId,omitempty" xml:"SourceId,omitempty" type:"Repeated"`
	// The status of the migration job. Valid values:
	//
	// 	- Ready: The migration job is not started.
	//
	// 	- Running: The migration job is running.
	//
	// 	- Stopped: The migration job is paused.
	//
	// 	- InError: An error occurs in the migration job.
	//
	// 	- Finished: The migration job is complete.
	//
	// 	- Waiting: The migration job is waiting to run.
	//
	// 	- Expired: The migration job has expired.
	//
	// 	- Deleting: The migration job is being deleted.
	//
	// example:
	//
	// Ready
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The information about tags that are attached to the SMC resource.
	Tag []*DescribeReplicationJobsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeReplicationJobsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeReplicationJobsRequest) GoString() string {
	return s.String()
}

func (s *DescribeReplicationJobsRequest) GetBusinessStatus() *string {
	return s.BusinessStatus
}

func (s *DescribeReplicationJobsRequest) GetInstanceId() []*string {
	return s.InstanceId
}

func (s *DescribeReplicationJobsRequest) GetJobId() []*string {
	return s.JobId
}

func (s *DescribeReplicationJobsRequest) GetJobType() *int32 {
	return s.JobType
}

func (s *DescribeReplicationJobsRequest) GetName() *string {
	return s.Name
}

func (s *DescribeReplicationJobsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeReplicationJobsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeReplicationJobsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeReplicationJobsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeReplicationJobsRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeReplicationJobsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeReplicationJobsRequest) GetSourceId() []*string {
	return s.SourceId
}

func (s *DescribeReplicationJobsRequest) GetStatus() *string {
	return s.Status
}

func (s *DescribeReplicationJobsRequest) GetTag() []*DescribeReplicationJobsRequestTag {
	return s.Tag
}

func (s *DescribeReplicationJobsRequest) SetBusinessStatus(v string) *DescribeReplicationJobsRequest {
	s.BusinessStatus = &v
	return s
}

func (s *DescribeReplicationJobsRequest) SetInstanceId(v []*string) *DescribeReplicationJobsRequest {
	s.InstanceId = v
	return s
}

func (s *DescribeReplicationJobsRequest) SetJobId(v []*string) *DescribeReplicationJobsRequest {
	s.JobId = v
	return s
}

func (s *DescribeReplicationJobsRequest) SetJobType(v int32) *DescribeReplicationJobsRequest {
	s.JobType = &v
	return s
}

func (s *DescribeReplicationJobsRequest) SetName(v string) *DescribeReplicationJobsRequest {
	s.Name = &v
	return s
}

func (s *DescribeReplicationJobsRequest) SetOwnerId(v int64) *DescribeReplicationJobsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeReplicationJobsRequest) SetPageNumber(v int32) *DescribeReplicationJobsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeReplicationJobsRequest) SetPageSize(v int32) *DescribeReplicationJobsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeReplicationJobsRequest) SetRegionId(v string) *DescribeReplicationJobsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeReplicationJobsRequest) SetResourceGroupId(v string) *DescribeReplicationJobsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeReplicationJobsRequest) SetResourceOwnerAccount(v string) *DescribeReplicationJobsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeReplicationJobsRequest) SetSourceId(v []*string) *DescribeReplicationJobsRequest {
	s.SourceId = v
	return s
}

func (s *DescribeReplicationJobsRequest) SetStatus(v string) *DescribeReplicationJobsRequest {
	s.Status = &v
	return s
}

func (s *DescribeReplicationJobsRequest) SetTag(v []*DescribeReplicationJobsRequestTag) *DescribeReplicationJobsRequest {
	s.Tag = v
	return s
}

func (s *DescribeReplicationJobsRequest) Validate() error {
	return dara.Validate(s)
}

type DescribeReplicationJobsRequestTag struct {
	// The key of the tag N that is added to the SMC resource. Valid values of N: 1 to 20.
	//
	// The tag key can be an empty string. It can be up to 64 characters in length and cannot contain http:// or https://.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N that is added to the SMC resource. Valid values of N: 1 to 20.
	//
	// The tag value can be an empty string. It can be up to 64 characters in length and cannot contain http:// or https://.[](http://https://。)
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeReplicationJobsRequestTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeReplicationJobsRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeReplicationJobsRequestTag) GetKey() *string {
	return s.Key
}

func (s *DescribeReplicationJobsRequestTag) GetValue() *string {
	return s.Value
}

func (s *DescribeReplicationJobsRequestTag) SetKey(v string) *DescribeReplicationJobsRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeReplicationJobsRequestTag) SetValue(v string) *DescribeReplicationJobsRequestTag {
	s.Value = &v
	return s
}

func (s *DescribeReplicationJobsRequestTag) Validate() error {
	return dara.Validate(s)
}
