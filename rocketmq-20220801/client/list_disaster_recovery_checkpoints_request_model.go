// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDisasterRecoveryCheckpointsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFilter(v string) *ListDisasterRecoveryCheckpointsRequest
	GetFilter() *string
	SetInstanceId(v string) *ListDisasterRecoveryCheckpointsRequest
	GetInstanceId() *string
	SetPageNumber(v int32) *ListDisasterRecoveryCheckpointsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListDisasterRecoveryCheckpointsRequest
	GetPageSize() *int32
}

type ListDisasterRecoveryCheckpointsRequest struct {
	// Filter Condition
	//
	// example:
	//
	// topic_test
	Filter *string `json:"filter,omitempty" xml:"filter,omitempty"`
	// Source Instance ID
	//
	// This parameter is required.
	//
	// example:
	//
	// rmq-cn-7e22ody****
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// Current page number, starting from 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// Page size, the maximum number of results returned per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s ListDisasterRecoveryCheckpointsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDisasterRecoveryCheckpointsRequest) GoString() string {
	return s.String()
}

func (s *ListDisasterRecoveryCheckpointsRequest) GetFilter() *string {
	return s.Filter
}

func (s *ListDisasterRecoveryCheckpointsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListDisasterRecoveryCheckpointsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListDisasterRecoveryCheckpointsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDisasterRecoveryCheckpointsRequest) SetFilter(v string) *ListDisasterRecoveryCheckpointsRequest {
	s.Filter = &v
	return s
}

func (s *ListDisasterRecoveryCheckpointsRequest) SetInstanceId(v string) *ListDisasterRecoveryCheckpointsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListDisasterRecoveryCheckpointsRequest) SetPageNumber(v int32) *ListDisasterRecoveryCheckpointsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListDisasterRecoveryCheckpointsRequest) SetPageSize(v int32) *ListDisasterRecoveryCheckpointsRequest {
	s.PageSize = &v
	return s
}

func (s *ListDisasterRecoveryCheckpointsRequest) Validate() error {
	return dara.Validate(s)
}
