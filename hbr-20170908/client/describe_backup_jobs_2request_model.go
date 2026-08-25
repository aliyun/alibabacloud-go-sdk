// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBackupJobs2Request interface {
	dara.Model
	String() string
	GoString() string
	SetEdition(v string) *DescribeBackupJobs2Request
	GetEdition() *string
	SetFilters(v []*DescribeBackupJobs2RequestFilters) *DescribeBackupJobs2Request
	GetFilters() []*DescribeBackupJobs2RequestFilters
	SetPageNumber(v int32) *DescribeBackupJobs2Request
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeBackupJobs2Request
	GetPageSize() *int32
	SetSortDirection(v string) *DescribeBackupJobs2Request
	GetSortDirection() *string
	SetSourceType(v string) *DescribeBackupJobs2Request
	GetSourceType() *string
}

type DescribeBackupJobs2Request struct {
	// The edition. Valid values: BASIC and STANDARD. The default value is STANDARD.
	//
	// example:
	//
	// STANDARD
	Edition *string `json:"Edition,omitempty" xml:"Edition,omitempty"`
	// The key-value pairs of the filter.
	Filters []*DescribeBackupJobs2RequestFilters `json:"Filters,omitempty" xml:"Filters,omitempty" type:"Repeated"`
	// The page number. Pages start from page 1. The default value is 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values: 1 to 99. The default value is 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The sort direction. Valid values:
	//
	// - **ASCEND**: Ascending order.
	//
	// - **DESCEND*	- (Default): Descending order.
	//
	// example:
	//
	// DESCEND
	SortDirection *string `json:"SortDirection,omitempty" xml:"SortDirection,omitempty"`
	// The type of the data source. Valid values:
	//
	// - **ECS_FILE**: Backs up Elastic Compute Service (ECS) files.
	//
	// - **OSS**: Backs up Alibaba Cloud Object Storage Service (OSS) buckets.
	//
	// - **NAS**: Backs up Alibaba Cloud Apsara File Storage NAS (NAS) file systems.
	//
	// - **OTS**: Backs up Alibaba Cloud Tablestore instances.
	//
	// - **UDM_ECS**: Backs up entire ECS instances.
	//
	// - **UDM_ECS_DISK**: A sub-task for disk backup in an ECS instance backup job.
	//
	// - **COMMON_NAS**: A generic NAS data source. This includes archive NAS and on-premises NAS data sources. Use the Values parameter of Filters to specify the data source type.
	//
	// - **File**: Backs up on-premises files.
	//
	// - **SYNC**: Data synchronization.
	//
	// example:
	//
	// ECS_FILE
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
}

func (s DescribeBackupJobs2Request) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupJobs2Request) GoString() string {
	return s.String()
}

func (s *DescribeBackupJobs2Request) GetEdition() *string {
	return s.Edition
}

func (s *DescribeBackupJobs2Request) GetFilters() []*DescribeBackupJobs2RequestFilters {
	return s.Filters
}

func (s *DescribeBackupJobs2Request) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeBackupJobs2Request) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeBackupJobs2Request) GetSortDirection() *string {
	return s.SortDirection
}

func (s *DescribeBackupJobs2Request) GetSourceType() *string {
	return s.SourceType
}

func (s *DescribeBackupJobs2Request) SetEdition(v string) *DescribeBackupJobs2Request {
	s.Edition = &v
	return s
}

func (s *DescribeBackupJobs2Request) SetFilters(v []*DescribeBackupJobs2RequestFilters) *DescribeBackupJobs2Request {
	s.Filters = v
	return s
}

func (s *DescribeBackupJobs2Request) SetPageNumber(v int32) *DescribeBackupJobs2Request {
	s.PageNumber = &v
	return s
}

func (s *DescribeBackupJobs2Request) SetPageSize(v int32) *DescribeBackupJobs2Request {
	s.PageSize = &v
	return s
}

func (s *DescribeBackupJobs2Request) SetSortDirection(v string) *DescribeBackupJobs2Request {
	s.SortDirection = &v
	return s
}

func (s *DescribeBackupJobs2Request) SetSourceType(v string) *DescribeBackupJobs2Request {
	s.SourceType = &v
	return s
}

func (s *DescribeBackupJobs2Request) Validate() error {
	if s.Filters != nil {
		for _, item := range s.Filters {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeBackupJobs2RequestFilters struct {
	// The key of the filter. Valid values:
	//
	// - **RegionId**: The region ID.
	//
	// - **PlanId**: The backup plan ID.
	//
	// - **JobId**: The backup job ID.
	//
	// - **VaultId**: The repository ID.
	//
	// - **InstanceId**: The ECS instance ID.
	//
	// - **Bucket**: The name of the OSS bucket.
	//
	// - **FileSystemId**: The file system ID.
	//
	// - **Status**: The job status.
	//
	// - **CreatedTime**: The start time of the job.
	//
	// - **CompleteTime**: The end time of the job.
	//
	// - **InstanceName**: The name of the Tablestore instance.
	//
	// - **BackupType**: The backup job. This parameter is required only when SourceType is set to COMMON_NAS.
	//
	// - **ParentId**: The ID of the parent job. This parameter is required when you query sub-tasks. For example, if you set SourceType to UDM_ECS_DISK, you must specify the ID of the UDM_ECS job.
	//
	// example:
	//
	// VaultId
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The matching operator. The default value is IN. This parameter specifies the operator to use for matching the Key and Value. Valid values:
	//
	// - **EQUAL**: Equal to.
	//
	// - **NOT_EQUAL**: Not equal to.
	//
	// - **GREATER_THAN**: Greater than.
	//
	// - **GREATER_THAN_OR_EQUAL**: Greater than or equal to.
	//
	// - **LESS_THAN**: Less than.
	//
	// - **LESS_THAN_OR_EQUAL**: Less than or equal to.
	//
	// - **BETWEEN**: The value is a JSON array in the format of `[start,end]`.
	//
	// - **IN**: The value is an array.
	//
	// > The IN operator is not supported when you use **CompleteTime*	- as the key for a query.
	//
	// example:
	//
	// IN
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	// The value of the filter.
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s DescribeBackupJobs2RequestFilters) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupJobs2RequestFilters) GoString() string {
	return s.String()
}

func (s *DescribeBackupJobs2RequestFilters) GetKey() *string {
	return s.Key
}

func (s *DescribeBackupJobs2RequestFilters) GetOperator() *string {
	return s.Operator
}

func (s *DescribeBackupJobs2RequestFilters) GetValues() []*string {
	return s.Values
}

func (s *DescribeBackupJobs2RequestFilters) SetKey(v string) *DescribeBackupJobs2RequestFilters {
	s.Key = &v
	return s
}

func (s *DescribeBackupJobs2RequestFilters) SetOperator(v string) *DescribeBackupJobs2RequestFilters {
	s.Operator = &v
	return s
}

func (s *DescribeBackupJobs2RequestFilters) SetValues(v []*string) *DescribeBackupJobs2RequestFilters {
	s.Values = v
	return s
}

func (s *DescribeBackupJobs2RequestFilters) Validate() error {
	return dara.Validate(s)
}
