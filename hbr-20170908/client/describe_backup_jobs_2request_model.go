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
	// example:
	//
	// STANDARD
	Edition *string `json:"Edition,omitempty" xml:"Edition,omitempty"`
	// The keys that you want to match in the filter.
	Filters []*DescribeBackupJobs2RequestFilters `json:"Filters,omitempty" xml:"Filters,omitempty" type:"Repeated"`
	// The number of the page to return. Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Valid values: 1 to 99. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The order in which you want to sort the results. Valid values:
	//
	// 	- **ASCEND**: sorts the results in ascending order
	//
	// 	- **DESCEND*	- (default value): sorts the results in descending order
	//
	// example:
	//
	// DESCEND
	SortDirection *string `json:"SortDirection,omitempty" xml:"SortDirection,omitempty"`
	// The type of the data source. Valid values:
	//
	// 	- **ECS_FILE**: Elastic Compute Service (ECS) files
	//
	// 	- **OSS**: Object Storage Service (OSS) buckets
	//
	// 	- **NAS**: Apsara File Storage NAS file systems
	//
	// 	- **OTS**: Tablestore instances
	//
	// 	- **UDM_ECS**: ECS instances
	//
	// 	- **UDM_ECS_DISK**: ECS disks
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
	return dara.Validate(s)
}

type DescribeBackupJobs2RequestFilters struct {
	// The keys in the filter. Valid values:
	//
	// 	- **RegionId**: the ID of a region
	//
	// 	- **PlanId**: the ID of a backup plan
	//
	// 	- **JobId**: the ID of a backup job
	//
	// 	- **VaultId**: the ID of a backup vault
	//
	// 	- **InstanceId**: the ID of an ECS instance
	//
	// 	- **Bucket**: the name of an OSS bucket
	//
	// 	- **FileSystemId**: the ID of a file system
	//
	// 	- **Status**: the status of a backup job
	//
	// 	- **CreatedTime**: the start time of a backup job
	//
	// 	- **CompleteTime**: the end time of a backup job
	//
	// 	- **instanceName**: the name of a Tablestore instance
	//
	// example:
	//
	// VaultId
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The matching method. Default value: IN. This parameter specifies the operator that you want to use to match a key and a value in the filter. Valid values:
	//
	// 	- **EQUAL**: equal to
	//
	// 	- **NOT_EQUAL**: not equal to
	//
	// 	- **GREATER_THAN**: greater than
	//
	// 	- **GREATER_THAN_OR_EQUAL**: greater than or equal to
	//
	// 	- **LESS_THAN**: less than
	//
	// 	- **LESS_THAN_OR_EQUAL**: less than or equal to
	//
	// 	- **BETWEEN**: specifies a JSON array as a range. The results must fall within the range in the `[Minimum value,maximum value]` format.
	//
	// 	- **IN**: specifies an array as a collection. The results must fall within the collection.
	//
	// >  If you specify **CompleteTime*	- as a key to query backup jobs, you cannot use the IN operator to perform a match.
	//
	// example:
	//
	// IN
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	// The values that you want to match in the filter.
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
