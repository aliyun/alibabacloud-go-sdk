// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCallTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizType(v string) *ListCallTaskRequest
	GetBizType() *string
	SetOwnerId(v int64) *ListCallTaskRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *ListCallTaskRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListCallTaskRequest
	GetPageSize() *int32
	SetResourceOwnerAccount(v string) *ListCallTaskRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ListCallTaskRequest
	GetResourceOwnerId() *int64
	SetStatus(v string) *ListCallTaskRequest
	GetStatus() *string
	SetTaskId(v string) *ListCallTaskRequest
	GetTaskId() *string
	SetTaskName(v string) *ListCallTaskRequest
	GetTaskName() *string
	SetTemplateName(v string) *ListCallTaskRequest
	GetTemplateName() *string
}

type ListCallTaskRequest struct {
	// The type of the task template. Valid values:
	//
	// 	- **VMS_VOICE_TTS**: the text-to-speech (TTS) notification template.
	//
	// 	- **VMS_VOICE_CODE**: the voice notification template.
	//
	// 	- **VMS_TTS**: the voice verification code template.
	//
	// example:
	//
	// VMS_VOICE_CODE
	BizType *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
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
	// 	- **INIT**: The task is in the initial state.
	//
	// 	- **RELEASE**: The task is being parsed.
	//
	// 	- **RUNNING**: The task is running.
	//
	// 	- **STOP**: The task is suspended.
	//
	// 	- **SYSTEM_STOP**: The task is suspended by the system.
	//
	// 	- **CANCEL**: The task is canceled.
	//
	// 	- **SYSTEM_CANCEL**: The task is canceled by the system.
	//
	// 	- **DONE**: The task is complete.
	//
	// example:
	//
	// DONE
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The task ID.
	//
	// example:
	//
	// 151001****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The task name.
	//
	// example:
	//
	// Aliyun
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// The template name.
	//
	// example:
	//
	// Test Template
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
}

func (s ListCallTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCallTaskRequest) GoString() string {
	return s.String()
}

func (s *ListCallTaskRequest) GetBizType() *string {
	return s.BizType
}

func (s *ListCallTaskRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListCallTaskRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListCallTaskRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListCallTaskRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ListCallTaskRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ListCallTaskRequest) GetStatus() *string {
	return s.Status
}

func (s *ListCallTaskRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *ListCallTaskRequest) GetTaskName() *string {
	return s.TaskName
}

func (s *ListCallTaskRequest) GetTemplateName() *string {
	return s.TemplateName
}

func (s *ListCallTaskRequest) SetBizType(v string) *ListCallTaskRequest {
	s.BizType = &v
	return s
}

func (s *ListCallTaskRequest) SetOwnerId(v int64) *ListCallTaskRequest {
	s.OwnerId = &v
	return s
}

func (s *ListCallTaskRequest) SetPageNumber(v int32) *ListCallTaskRequest {
	s.PageNumber = &v
	return s
}

func (s *ListCallTaskRequest) SetPageSize(v int32) *ListCallTaskRequest {
	s.PageSize = &v
	return s
}

func (s *ListCallTaskRequest) SetResourceOwnerAccount(v string) *ListCallTaskRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListCallTaskRequest) SetResourceOwnerId(v int64) *ListCallTaskRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListCallTaskRequest) SetStatus(v string) *ListCallTaskRequest {
	s.Status = &v
	return s
}

func (s *ListCallTaskRequest) SetTaskId(v string) *ListCallTaskRequest {
	s.TaskId = &v
	return s
}

func (s *ListCallTaskRequest) SetTaskName(v string) *ListCallTaskRequest {
	s.TaskName = &v
	return s
}

func (s *ListCallTaskRequest) SetTemplateName(v string) *ListCallTaskRequest {
	s.TemplateName = &v
	return s
}

func (s *ListCallTaskRequest) Validate() error {
	return dara.Validate(s)
}
