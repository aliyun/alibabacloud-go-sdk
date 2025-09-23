// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMigrationOperationsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFilter(v string) *ListMigrationOperationsRequest
	GetFilter() *string
	SetInstanceId(v string) *ListMigrationOperationsRequest
	GetInstanceId() *string
	SetOperationType(v string) *ListMigrationOperationsRequest
	GetOperationType() *string
	SetPageNumber(v int32) *ListMigrationOperationsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListMigrationOperationsRequest
	GetPageSize() *int32
}

type ListMigrationOperationsRequest struct {
	// example:
	//
	// xx
	Filter *string `json:"filter,omitempty" xml:"filter,omitempty"`
	// example:
	//
	// rmq-cn-zxu37eiqq0x
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
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

func (s ListMigrationOperationsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListMigrationOperationsRequest) GoString() string {
	return s.String()
}

func (s *ListMigrationOperationsRequest) GetFilter() *string {
	return s.Filter
}

func (s *ListMigrationOperationsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListMigrationOperationsRequest) GetOperationType() *string {
	return s.OperationType
}

func (s *ListMigrationOperationsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListMigrationOperationsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListMigrationOperationsRequest) SetFilter(v string) *ListMigrationOperationsRequest {
	s.Filter = &v
	return s
}

func (s *ListMigrationOperationsRequest) SetInstanceId(v string) *ListMigrationOperationsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListMigrationOperationsRequest) SetOperationType(v string) *ListMigrationOperationsRequest {
	s.OperationType = &v
	return s
}

func (s *ListMigrationOperationsRequest) SetPageNumber(v int32) *ListMigrationOperationsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListMigrationOperationsRequest) SetPageSize(v int32) *ListMigrationOperationsRequest {
	s.PageSize = &v
	return s
}

func (s *ListMigrationOperationsRequest) Validate() error {
	return dara.Validate(s)
}
