// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBackupPlansRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEdition(v string) *DescribeBackupPlansRequest
	GetEdition() *string
	SetFilters(v []*DescribeBackupPlansRequestFilters) *DescribeBackupPlansRequest
	GetFilters() []*DescribeBackupPlansRequestFilters
	SetPageNumber(v int32) *DescribeBackupPlansRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeBackupPlansRequest
	GetPageSize() *int32
	SetSourceType(v string) *DescribeBackupPlansRequest
	GetSourceType() *string
}

type DescribeBackupPlansRequest struct {
	// example:
	//
	// STANDARD
	Edition *string `json:"Edition,omitempty" xml:"Edition,omitempty"`
	// The filters.
	Filters []*DescribeBackupPlansRequestFilters `json:"Filters,omitempty" xml:"Filters,omitempty" type:"Repeated"`
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
	// 	- **NAS**: File Storage NAS (NAS) file systems
	//
	// 	- **OTS**: Tablestore instances
	//
	// 	- **UDM_ECS**: ECS instances
	//
	// 	- **SYNC**: data synchronization
	//
	// example:
	//
	// ECS_FILE
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
}

func (s DescribeBackupPlansRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupPlansRequest) GoString() string {
	return s.String()
}

func (s *DescribeBackupPlansRequest) GetEdition() *string {
	return s.Edition
}

func (s *DescribeBackupPlansRequest) GetFilters() []*DescribeBackupPlansRequestFilters {
	return s.Filters
}

func (s *DescribeBackupPlansRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeBackupPlansRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeBackupPlansRequest) GetSourceType() *string {
	return s.SourceType
}

func (s *DescribeBackupPlansRequest) SetEdition(v string) *DescribeBackupPlansRequest {
	s.Edition = &v
	return s
}

func (s *DescribeBackupPlansRequest) SetFilters(v []*DescribeBackupPlansRequestFilters) *DescribeBackupPlansRequest {
	s.Filters = v
	return s
}

func (s *DescribeBackupPlansRequest) SetPageNumber(v int32) *DescribeBackupPlansRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeBackupPlansRequest) SetPageSize(v int32) *DescribeBackupPlansRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeBackupPlansRequest) SetSourceType(v string) *DescribeBackupPlansRequest {
	s.SourceType = &v
	return s
}

func (s *DescribeBackupPlansRequest) Validate() error {
	return dara.Validate(s)
}

type DescribeBackupPlansRequestFilters struct {
	// The keys in the filter. Valid values:
	//
	// 	- **regionId**: the ID of a region
	//
	// 	- **planId**: the ID of a backup plan
	//
	// 	- **sourceType**: the type of a data source
	//
	// 	- **vaultId**: the ID of a backup vault
	//
	// 	- **instanceName**: the name of an instance
	//
	// 	- **instanceId**: the ID of an instance
	//
	// 	- **planName**: the name of a backup plan
	//
	// example:
	//
	// vaultId
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The values that you want to match in the filter.
	//
	// example:
	//
	// ["v-*********************"]
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s DescribeBackupPlansRequestFilters) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupPlansRequestFilters) GoString() string {
	return s.String()
}

func (s *DescribeBackupPlansRequestFilters) GetKey() *string {
	return s.Key
}

func (s *DescribeBackupPlansRequestFilters) GetValues() []*string {
	return s.Values
}

func (s *DescribeBackupPlansRequestFilters) SetKey(v string) *DescribeBackupPlansRequestFilters {
	s.Key = &v
	return s
}

func (s *DescribeBackupPlansRequestFilters) SetValues(v []*string) *DescribeBackupPlansRequestFilters {
	s.Values = v
	return s
}

func (s *DescribeBackupPlansRequestFilters) Validate() error {
	return dara.Validate(s)
}
