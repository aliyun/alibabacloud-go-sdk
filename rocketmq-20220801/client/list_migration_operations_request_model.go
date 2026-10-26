// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMigrationOperationsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBusinessStatus(v []*string) *ListMigrationOperationsRequest
	GetBusinessStatus() []*string
	SetFilter(v string) *ListMigrationOperationsRequest
	GetFilter() *string
	SetInstanceId(v string) *ListMigrationOperationsRequest
	GetInstanceId() *string
	SetOperationStatus(v []*string) *ListMigrationOperationsRequest
	GetOperationStatus() []*string
	SetOperationType(v string) *ListMigrationOperationsRequest
	GetOperationType() *string
	SetPageNumber(v int32) *ListMigrationOperationsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListMigrationOperationsRequest
	GetPageSize() *int32
}

type ListMigrationOperationsRequest struct {
	// The operation business status.
	BusinessStatus []*string `json:"businessStatus,omitempty" xml:"businessStatus,omitempty" type:"Repeated"`
	// The filter condition for the query.
	//
	// >
	//
	// > - Supports filtering by operationKey, which is the topic name.
	//
	// example:
	//
	// xx
	Filter *string `json:"filter,omitempty" xml:"filter,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// rmq-cn-zxu37eiqq0x
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// The operation status.
	OperationStatus []*string `json:"operationStatus,omitempty" xml:"operationStatus,omitempty" type:"Repeated"`
	// The operation type.
	//
	// This parameter is required.
	//
	// example:
	//
	// MIGRATE_MESSAGE
	OperationType *string `json:"operationType,omitempty" xml:"operationType,omitempty"`
	// The page number of the results to return.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The page size, which specifies the maximum number of results to return per page.
	//
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

func (s *ListMigrationOperationsRequest) GetBusinessStatus() []*string {
	return s.BusinessStatus
}

func (s *ListMigrationOperationsRequest) GetFilter() *string {
	return s.Filter
}

func (s *ListMigrationOperationsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListMigrationOperationsRequest) GetOperationStatus() []*string {
	return s.OperationStatus
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

func (s *ListMigrationOperationsRequest) SetBusinessStatus(v []*string) *ListMigrationOperationsRequest {
	s.BusinessStatus = v
	return s
}

func (s *ListMigrationOperationsRequest) SetFilter(v string) *ListMigrationOperationsRequest {
	s.Filter = &v
	return s
}

func (s *ListMigrationOperationsRequest) SetInstanceId(v string) *ListMigrationOperationsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListMigrationOperationsRequest) SetOperationStatus(v []*string) *ListMigrationOperationsRequest {
	s.OperationStatus = v
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
