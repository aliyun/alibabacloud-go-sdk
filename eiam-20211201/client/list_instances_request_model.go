// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCrossRegionReplication(v string) *ListInstancesRequest
	GetCrossRegionReplication() *string
	SetEdition(v string) *ListInstancesRequest
	GetEdition() *string
	SetInstanceIds(v []*string) *ListInstancesRequest
	GetInstanceIds() []*string
	SetPageNumber(v int64) *ListInstancesRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *ListInstancesRequest
	GetPageSize() *int64
	SetStatus(v string) *ListInstancesRequest
	GetStatus() *string
}

type ListInstancesRequest struct {
	// The cross-region replication status.
	//
	// example:
	//
	// disabled
	CrossRegionReplication *string `json:"CrossRegionReplication,omitempty" xml:"CrossRegionReplication,omitempty"`
	// The license edition. Valid values:
	//
	// - free: Free Edition.
	//
	// - trial: Trial Edition.
	//
	// - scalability: Scalability Edition.
	//
	// - standard: Standard Edition.
	//
	// - enterprise: Enterprise Edition.
	//
	// example:
	//
	// free
	Edition *string `json:"Edition,omitempty" xml:"Edition,omitempty"`
	// The list of instance IDs.
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The page size.
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The instance status. Valid values:
	//
	// - creating: Being created.
	//
	// - running: Running.
	//
	// example:
	//
	// running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListInstancesRequest) GoString() string {
	return s.String()
}

func (s *ListInstancesRequest) GetCrossRegionReplication() *string {
	return s.CrossRegionReplication
}

func (s *ListInstancesRequest) GetEdition() *string {
	return s.Edition
}

func (s *ListInstancesRequest) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *ListInstancesRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListInstancesRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListInstancesRequest) GetStatus() *string {
	return s.Status
}

func (s *ListInstancesRequest) SetCrossRegionReplication(v string) *ListInstancesRequest {
	s.CrossRegionReplication = &v
	return s
}

func (s *ListInstancesRequest) SetEdition(v string) *ListInstancesRequest {
	s.Edition = &v
	return s
}

func (s *ListInstancesRequest) SetInstanceIds(v []*string) *ListInstancesRequest {
	s.InstanceIds = v
	return s
}

func (s *ListInstancesRequest) SetPageNumber(v int64) *ListInstancesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListInstancesRequest) SetPageSize(v int64) *ListInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *ListInstancesRequest) SetStatus(v string) *ListInstancesRequest {
	s.Status = &v
	return s
}

func (s *ListInstancesRequest) Validate() error {
	return dara.Validate(s)
}
