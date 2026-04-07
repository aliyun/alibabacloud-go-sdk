// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMigrationOperationsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBusinessStatusShrink(v string) *ListMigrationOperationsShrinkRequest
	GetBusinessStatusShrink() *string
	SetFilter(v string) *ListMigrationOperationsShrinkRequest
	GetFilter() *string
	SetInstanceId(v string) *ListMigrationOperationsShrinkRequest
	GetInstanceId() *string
	SetOperationStatusShrink(v string) *ListMigrationOperationsShrinkRequest
	GetOperationStatusShrink() *string
	SetOperationType(v string) *ListMigrationOperationsShrinkRequest
	GetOperationType() *string
	SetPageNumber(v int32) *ListMigrationOperationsShrinkRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListMigrationOperationsShrinkRequest
	GetPageSize() *int32
}

type ListMigrationOperationsShrinkRequest struct {
	BusinessStatusShrink *string `json:"businessStatus,omitempty" xml:"businessStatus,omitempty"`
	// example:
	//
	// xx
	Filter *string `json:"filter,omitempty" xml:"filter,omitempty"`
	// example:
	//
	// rmq-cn-zxu37eiqq0x
	InstanceId            *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	OperationStatusShrink *string `json:"operationStatus,omitempty" xml:"operationStatus,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// IMPORT_CONSUMER_GROUP
	OperationType *string `json:"operationType,omitempty" xml:"operationType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s ListMigrationOperationsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListMigrationOperationsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListMigrationOperationsShrinkRequest) GetBusinessStatusShrink() *string {
	return s.BusinessStatusShrink
}

func (s *ListMigrationOperationsShrinkRequest) GetFilter() *string {
	return s.Filter
}

func (s *ListMigrationOperationsShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListMigrationOperationsShrinkRequest) GetOperationStatusShrink() *string {
	return s.OperationStatusShrink
}

func (s *ListMigrationOperationsShrinkRequest) GetOperationType() *string {
	return s.OperationType
}

func (s *ListMigrationOperationsShrinkRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListMigrationOperationsShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListMigrationOperationsShrinkRequest) SetBusinessStatusShrink(v string) *ListMigrationOperationsShrinkRequest {
	s.BusinessStatusShrink = &v
	return s
}

func (s *ListMigrationOperationsShrinkRequest) SetFilter(v string) *ListMigrationOperationsShrinkRequest {
	s.Filter = &v
	return s
}

func (s *ListMigrationOperationsShrinkRequest) SetInstanceId(v string) *ListMigrationOperationsShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *ListMigrationOperationsShrinkRequest) SetOperationStatusShrink(v string) *ListMigrationOperationsShrinkRequest {
	s.OperationStatusShrink = &v
	return s
}

func (s *ListMigrationOperationsShrinkRequest) SetOperationType(v string) *ListMigrationOperationsShrinkRequest {
	s.OperationType = &v
	return s
}

func (s *ListMigrationOperationsShrinkRequest) SetPageNumber(v int32) *ListMigrationOperationsShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListMigrationOperationsShrinkRequest) SetPageSize(v int32) *ListMigrationOperationsShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListMigrationOperationsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
