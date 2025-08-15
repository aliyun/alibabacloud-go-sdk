// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDisasterRecoveryPlansRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFilter(v string) *ListDisasterRecoveryPlansRequest
	GetFilter() *string
	SetInstanceId(v string) *ListDisasterRecoveryPlansRequest
	GetInstanceId() *string
	SetPageNumber(v int32) *ListDisasterRecoveryPlansRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListDisasterRecoveryPlansRequest
	GetPageSize() *int32
}

type ListDisasterRecoveryPlansRequest struct {
	// The filter condition. Global Replicator tasks are filtered by task name or description.
	//
	// example:
	//
	// xxx
	Filter *string `json:"filter,omitempty" xml:"filter,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// rmq-cn-778xxx
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s ListDisasterRecoveryPlansRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDisasterRecoveryPlansRequest) GoString() string {
	return s.String()
}

func (s *ListDisasterRecoveryPlansRequest) GetFilter() *string {
	return s.Filter
}

func (s *ListDisasterRecoveryPlansRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListDisasterRecoveryPlansRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListDisasterRecoveryPlansRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDisasterRecoveryPlansRequest) SetFilter(v string) *ListDisasterRecoveryPlansRequest {
	s.Filter = &v
	return s
}

func (s *ListDisasterRecoveryPlansRequest) SetInstanceId(v string) *ListDisasterRecoveryPlansRequest {
	s.InstanceId = &v
	return s
}

func (s *ListDisasterRecoveryPlansRequest) SetPageNumber(v int32) *ListDisasterRecoveryPlansRequest {
	s.PageNumber = &v
	return s
}

func (s *ListDisasterRecoveryPlansRequest) SetPageSize(v int32) *ListDisasterRecoveryPlansRequest {
	s.PageSize = &v
	return s
}

func (s *ListDisasterRecoveryPlansRequest) Validate() error {
	return dara.Validate(s)
}
