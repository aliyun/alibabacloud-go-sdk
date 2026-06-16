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
	// example:
	//
	// disabled
	CrossRegionReplication *string `json:"CrossRegionReplication,omitempty" xml:"CrossRegionReplication,omitempty"`
	// The edition of the license. Valid values:
	//
	// - free: Free edition.
	//
	// - trial: Trial edition.
	//
	// - scalability: Scalability edition.
	//
	// - standard: Standard edition.
	//
	// - enterprise: Enterprise edition.
	//
	// example:
	//
	// free
	Edition *string `json:"Edition,omitempty" xml:"Edition,omitempty"`
	// Instance ID list.
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// Page number.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// Page size.
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Instance status. Valid values:
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
