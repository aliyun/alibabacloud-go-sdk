// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTaskDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCalled(v string) *ListTaskDetailRequest
	GetCalled() *string
	SetId(v int64) *ListTaskDetailRequest
	GetId() *int64
	SetOwnerId(v int64) *ListTaskDetailRequest
	GetOwnerId() *int64
	SetPageNo(v int32) *ListTaskDetailRequest
	GetPageNo() *int32
	SetPageSize(v int32) *ListTaskDetailRequest
	GetPageSize() *int32
	SetResourceOwnerAccount(v string) *ListTaskDetailRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ListTaskDetailRequest
	GetResourceOwnerId() *int64
	SetStatus(v string) *ListTaskDetailRequest
	GetStatus() *string
	SetStatusCode(v string) *ListTaskDetailRequest
	GetStatusCode() *string
	SetTaskId(v int64) *ListTaskDetailRequest
	GetTaskId() *int64
}

type ListTaskDetailRequest struct {
	// example:
	//
	// 186****0000
	Called *string `json:"Called,omitempty" xml:"Called,omitempty"`
	// example:
	//
	// 123456
	Id      *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// example:
	//
	// 20
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// example:
	//
	// SUCCESS
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 000001
	StatusCode *string `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123456
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ListTaskDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTaskDetailRequest) GoString() string {
	return s.String()
}

func (s *ListTaskDetailRequest) GetCalled() *string {
	return s.Called
}

func (s *ListTaskDetailRequest) GetId() *int64 {
	return s.Id
}

func (s *ListTaskDetailRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListTaskDetailRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *ListTaskDetailRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListTaskDetailRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ListTaskDetailRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ListTaskDetailRequest) GetStatus() *string {
	return s.Status
}

func (s *ListTaskDetailRequest) GetStatusCode() *string {
	return s.StatusCode
}

func (s *ListTaskDetailRequest) GetTaskId() *int64 {
	return s.TaskId
}

func (s *ListTaskDetailRequest) SetCalled(v string) *ListTaskDetailRequest {
	s.Called = &v
	return s
}

func (s *ListTaskDetailRequest) SetId(v int64) *ListTaskDetailRequest {
	s.Id = &v
	return s
}

func (s *ListTaskDetailRequest) SetOwnerId(v int64) *ListTaskDetailRequest {
	s.OwnerId = &v
	return s
}

func (s *ListTaskDetailRequest) SetPageNo(v int32) *ListTaskDetailRequest {
	s.PageNo = &v
	return s
}

func (s *ListTaskDetailRequest) SetPageSize(v int32) *ListTaskDetailRequest {
	s.PageSize = &v
	return s
}

func (s *ListTaskDetailRequest) SetResourceOwnerAccount(v string) *ListTaskDetailRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListTaskDetailRequest) SetResourceOwnerId(v int64) *ListTaskDetailRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListTaskDetailRequest) SetStatus(v string) *ListTaskDetailRequest {
	s.Status = &v
	return s
}

func (s *ListTaskDetailRequest) SetStatusCode(v string) *ListTaskDetailRequest {
	s.StatusCode = &v
	return s
}

func (s *ListTaskDetailRequest) SetTaskId(v int64) *ListTaskDetailRequest {
	s.TaskId = &v
	return s
}

func (s *ListTaskDetailRequest) Validate() error {
	return dara.Validate(s)
}
