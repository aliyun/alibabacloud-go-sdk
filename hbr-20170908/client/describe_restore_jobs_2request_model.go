// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRestoreJobs2Request interface {
	dara.Model
	String() string
	GoString() string
	SetEdition(v string) *DescribeRestoreJobs2Request
	GetEdition() *string
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
	// The edition. Valid values: `BASIC` and `STANDARD`. Default value: `STANDARD`.
	//
	// example:
	//
	// STANDARD
	Edition *string `json:"Edition,omitempty" xml:"Edition,omitempty"`
	// The filter conditions.
	Filters []*DescribeRestoreJobs2RequestFilters `json:"Filters,omitempty" xml:"Filters,omitempty" type:"Repeated"`
	// The page number. Pages start from 1. Default value: 1.
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
	// The data source type. Valid values:
	//
	// - **ECS_FILE**: Restores ECS files.
	//
	// - **OSS**: Restores OSS objects.
	//
	// - **NAS**: Restores NAS files.
	//
	// - **COMMON_FILE_SYSTEM**: Restores data to a CPFS file system.
	//
	// - **OTS_TABLE**: Restores an OTS table.
	//
	// - **UDM_ECS_ROLLBACK**: Restores an entire ECS instance.
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

func (s *DescribeRestoreJobs2Request) GetEdition() *string {
	return s.Edition
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

func (s *DescribeRestoreJobs2Request) SetEdition(v string) *DescribeRestoreJobs2Request {
	s.Edition = &v
	return s
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

type DescribeRestoreJobs2RequestFilters struct {
	// The filter key. Valid values:
	//
	// - **RegionId**: region ID
	//
	// - **PlanId**: backup plan ID
	//
	// - **JobId**: backup job ID
	//
	// - **VaultId**: vault ID
	//
	// - **InstanceId**: ECS instance ID
	//
	// - **Bucket**: OSS bucket name
	//
	// - **FileSystemId**: file system ID
	//
	// - **Status**: job status
	//
	// - **CompleteTime**: completion time
	//
	// example:
	//
	// VaultId
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The matching method. The default value is IN. Valid values:
	//
	// - **EQUAL**: Equal to
	//
	// - **NOT_EQUAL**: Not equal to
	//
	// - **GREATER_THAN**: Greater than
	//
	// - **GREATER_THAN_OR_EQUAL**: Greater than or equal to
	//
	// - **LESS_THAN**: Less than
	//
	// - **LESS_THAN_OR_EQUAL**: Less than or equal to
	//
	// - **BETWEEN**: The value is within a specified range. The `Values` parameter must be a JSON array in the `[min, max]` format.
	//
	// - **IN**: The value is in a specified set. The `Values` parameter must be an array.
	//
	// > The IN operator is not supported when `Key` is **CompleteTime**.
	//
	// example:
	//
	// IN
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	// An array of values for the specified filter key.
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
