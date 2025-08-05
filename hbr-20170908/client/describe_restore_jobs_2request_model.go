// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRestoreJobs2Request interface {
	dara.Model
	String() string
	GoString() string
	SetFilters(v []*DescribeRestoreJobs2RequestFilters) *DescribeRestoreJobs2Request
	GetFilters() []*DescribeRestoreJobs2RequestFilters
	SetPageNumber(v int32) *DescribeRestoreJobs2Request
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeRestoreJobs2Request
	GetPageSize() *int32
	SetRestoreType(v string) *DescribeRestoreJobs2Request
	GetRestoreType() *string
}

type DescribeRestoreJobs2Request struct {
	// The keys in the filter.
	Filters []*DescribeRestoreJobs2RequestFilters `json:"Filters,omitempty" xml:"Filters,omitempty" type:"Repeated"`
	// The page number. Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values: 1 to 99. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The type of the data source. Valid values:
	//
	// 	- **ECS_FILE**: Elastic Compute Service (ECS) files
	//
	// 	- **OSS**: Object Storage Service (OSS) buckets
	//
	// 	- **NAS**: Apsara File Storage NAS file systems
	//
	// 	- **OTS_TABLE**: Tablestore instances
	//
	// 	- **UDM_ECS_ROLLBACK**: ECS instances
	//
	// example:
	//
	// ECS_FILE
	RestoreType *string `json:"RestoreType,omitempty" xml:"RestoreType,omitempty"`
}

func (s DescribeRestoreJobs2Request) String() string {
	return dara.Prettify(s)
}

func (s DescribeRestoreJobs2Request) GoString() string {
	return s.String()
}

func (s *DescribeRestoreJobs2Request) GetFilters() []*DescribeRestoreJobs2RequestFilters {
	return s.Filters
}

func (s *DescribeRestoreJobs2Request) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeRestoreJobs2Request) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeRestoreJobs2Request) GetRestoreType() *string {
	return s.RestoreType
}

func (s *DescribeRestoreJobs2Request) SetFilters(v []*DescribeRestoreJobs2RequestFilters) *DescribeRestoreJobs2Request {
	s.Filters = v
	return s
}

func (s *DescribeRestoreJobs2Request) SetPageNumber(v int32) *DescribeRestoreJobs2Request {
	s.PageNumber = &v
	return s
}

func (s *DescribeRestoreJobs2Request) SetPageSize(v int32) *DescribeRestoreJobs2Request {
	s.PageSize = &v
	return s
}

func (s *DescribeRestoreJobs2Request) SetRestoreType(v string) *DescribeRestoreJobs2Request {
	s.RestoreType = &v
	return s
}

func (s *DescribeRestoreJobs2Request) Validate() error {
	return dara.Validate(s)
}

type DescribeRestoreJobs2RequestFilters struct {
	// The key in the filter. Valid values:
	//
	// 	- **RegionId**: the region ID
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
	// 	- **CompleteTime**: the end time of a backup job
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
	// 	- **BETWEEN**: specifies a JSON array as a range. The results must fall within the range in the `[Minimum value,Maximum value]` format.
	//
	// 	- **IN**: specifies an array as a collection. The results must fall within the collection.
	//
	// > If you specify the **CompleteTime*	- parameter as a key to query backup jobs, you cannot use the IN operator to perform a match.
	//
	// example:
	//
	// IN
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	// The values that you want to match in the filter.
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s DescribeRestoreJobs2RequestFilters) String() string {
	return dara.Prettify(s)
}

func (s DescribeRestoreJobs2RequestFilters) GoString() string {
	return s.String()
}

func (s *DescribeRestoreJobs2RequestFilters) GetKey() *string {
	return s.Key
}

func (s *DescribeRestoreJobs2RequestFilters) GetOperator() *string {
	return s.Operator
}

func (s *DescribeRestoreJobs2RequestFilters) GetValues() []*string {
	return s.Values
}

func (s *DescribeRestoreJobs2RequestFilters) SetKey(v string) *DescribeRestoreJobs2RequestFilters {
	s.Key = &v
	return s
}

func (s *DescribeRestoreJobs2RequestFilters) SetOperator(v string) *DescribeRestoreJobs2RequestFilters {
	s.Operator = &v
	return s
}

func (s *DescribeRestoreJobs2RequestFilters) SetValues(v []*string) *DescribeRestoreJobs2RequestFilters {
	s.Values = v
	return s
}

func (s *DescribeRestoreJobs2RequestFilters) Validate() error {
	return dara.Validate(s)
}
