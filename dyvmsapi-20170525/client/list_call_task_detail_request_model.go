// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCallTaskDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCalledNum(v string) *ListCallTaskDetailRequest
	GetCalledNum() *string
	SetOwnerId(v int64) *ListCallTaskDetailRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *ListCallTaskDetailRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListCallTaskDetailRequest
	GetPageSize() *int32
	SetResourceOwnerAccount(v string) *ListCallTaskDetailRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ListCallTaskDetailRequest
	GetResourceOwnerId() *int64
	SetStatus(v string) *ListCallTaskDetailRequest
	GetStatus() *string
	SetTaskId(v int64) *ListCallTaskDetailRequest
	GetTaskId() *int64
}

type ListCallTaskDetailRequest struct {
	// The called number.
	//
	// example:
	//
	// 1300000****
	CalledNum *string `json:"CalledNum,omitempty" xml:"CalledNum,omitempty"`
	OwnerId   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: **10**.
	//
	// example:
	//
	// 10
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The task state. Valid values:
	//
	// 	- **SUCCESS**: The task is successful.
	//
	// 	- **FAIL**: The task fails.
	//
	// 	- **INIT**: The task is not started.
	//
	// example:
	//
	// SUCCESS
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The task ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 150001****
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ListCallTaskDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCallTaskDetailRequest) GoString() string {
	return s.String()
}

func (s *ListCallTaskDetailRequest) GetCalledNum() *string {
	return s.CalledNum
}

func (s *ListCallTaskDetailRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListCallTaskDetailRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListCallTaskDetailRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListCallTaskDetailRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ListCallTaskDetailRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ListCallTaskDetailRequest) GetStatus() *string {
	return s.Status
}

func (s *ListCallTaskDetailRequest) GetTaskId() *int64 {
	return s.TaskId
}

func (s *ListCallTaskDetailRequest) SetCalledNum(v string) *ListCallTaskDetailRequest {
	s.CalledNum = &v
	return s
}

func (s *ListCallTaskDetailRequest) SetOwnerId(v int64) *ListCallTaskDetailRequest {
	s.OwnerId = &v
	return s
}

func (s *ListCallTaskDetailRequest) SetPageNumber(v int32) *ListCallTaskDetailRequest {
	s.PageNumber = &v
	return s
}

func (s *ListCallTaskDetailRequest) SetPageSize(v int32) *ListCallTaskDetailRequest {
	s.PageSize = &v
	return s
}

func (s *ListCallTaskDetailRequest) SetResourceOwnerAccount(v string) *ListCallTaskDetailRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListCallTaskDetailRequest) SetResourceOwnerId(v int64) *ListCallTaskDetailRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListCallTaskDetailRequest) SetStatus(v string) *ListCallTaskDetailRequest {
	s.Status = &v
	return s
}

func (s *ListCallTaskDetailRequest) SetTaskId(v int64) *ListCallTaskDetailRequest {
	s.TaskId = &v
	return s
}

func (s *ListCallTaskDetailRequest) Validate() error {
	return dara.Validate(s)
}
