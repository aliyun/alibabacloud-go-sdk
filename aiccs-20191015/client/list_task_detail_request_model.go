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
	// The callee number. You can view the callee number on the **Detail*	- interface of [**Task Management**](https://aiccs.console.aliyun.com/job/list).
	//
	// example:
	//
	// 186****0000
	Called *string `json:"Called,omitempty" xml:"Called,omitempty"`
	// The detail ID. You can view the detail ID on the **Detail*	- interface of [**Task Management**](https://aiccs.console.aliyun.com/job/list).
	//
	// example:
	//
	// 12****
	Id      *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The current page number. The value must be greater than **0**. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page. The value must be greater than **0**. Default value: **20**.
	//
	// example:
	//
	// 20
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// Detail status. Valid values:
	//
	// - **SUCCESS**: The outbound call succeeded.
	//
	// - **FAIL**: The outbound call failed.
	//
	// - **INIT**: The outbound call has not been made.
	//
	// example:
	//
	// SUCCESS
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The call status code. For more information, see [Call Status Codes](https://help.aliyun.com/document_detail/112804.html) in Voice Service.
	//
	// example:
	//
	// 200100
	StatusCode *string `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
	// The job ID. You can view the job ID on the [Task Management](https://aiccs.console.aliyun.com/job/list) page or obtain it by using the [ListTask](https://help.aliyun.com/document_detail/2718008.html) API.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12****
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
