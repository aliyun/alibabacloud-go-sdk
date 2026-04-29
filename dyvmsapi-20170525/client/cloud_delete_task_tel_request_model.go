// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudDeleteTaskTelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnterpriseId(v int64) *CloudDeleteTaskTelRequest
	GetEnterpriseId() *int64
	SetFileId(v int64) *CloudDeleteTaskTelRequest
	GetFileId() *int64
	SetFileName(v string) *CloudDeleteTaskTelRequest
	GetFileName() *string
	SetOwnerId(v int64) *CloudDeleteTaskTelRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *CloudDeleteTaskTelRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CloudDeleteTaskTelRequest
	GetResourceOwnerId() *int64
	SetStatus(v int64) *CloudDeleteTaskTelRequest
	GetStatus() *int64
	SetTaskId(v int64) *CloudDeleteTaskTelRequest
	GetTaskId() *int64
	SetTels(v string) *CloudDeleteTaskTelRequest
	GetTels() *string
}

type CloudDeleteTaskTelRequest struct {
	// 呼叫中心 id
	//
	// This parameter is required.
	//
	// example:
	//
	// 7000002
	EnterpriseId *int64 `json:"EnterpriseId,omitempty" xml:"EnterpriseId,omitempty"`
	// 批次Id
	//
	// example:
	//
	// 1
	FileId *int64 `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// 批次名称；按照名称删除会删除最近添加的批次
	//
	// example:
	//
	// test_1
	FileName             *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// 号码状态；0:未呼叫 1:已呼叫
	//
	// example:
	//
	// 0
	Status *int64 `json:"Status,omitempty" xml:"Status,omitempty"`
	// 任务Id
	//
	// This parameter is required.
	//
	// example:
	//
	// 35
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// 电话号码；支持多个，多个号码用英文逗号","分隔；任务在运行中，号码也会被删除，最多支持10W个号码；支持加密号码删除，加密号码传参时需URL Encode
	//
	// example:
	//
	// 18360957135
	Tels *string `json:"Tels,omitempty" xml:"Tels,omitempty"`
}

func (s CloudDeleteTaskTelRequest) String() string {
	return dara.Prettify(s)
}

func (s CloudDeleteTaskTelRequest) GoString() string {
	return s.String()
}

func (s *CloudDeleteTaskTelRequest) GetEnterpriseId() *int64 {
	return s.EnterpriseId
}

func (s *CloudDeleteTaskTelRequest) GetFileId() *int64 {
	return s.FileId
}

func (s *CloudDeleteTaskTelRequest) GetFileName() *string {
	return s.FileName
}

func (s *CloudDeleteTaskTelRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CloudDeleteTaskTelRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CloudDeleteTaskTelRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CloudDeleteTaskTelRequest) GetStatus() *int64 {
	return s.Status
}

func (s *CloudDeleteTaskTelRequest) GetTaskId() *int64 {
	return s.TaskId
}

func (s *CloudDeleteTaskTelRequest) GetTels() *string {
	return s.Tels
}

func (s *CloudDeleteTaskTelRequest) SetEnterpriseId(v int64) *CloudDeleteTaskTelRequest {
	s.EnterpriseId = &v
	return s
}

func (s *CloudDeleteTaskTelRequest) SetFileId(v int64) *CloudDeleteTaskTelRequest {
	s.FileId = &v
	return s
}

func (s *CloudDeleteTaskTelRequest) SetFileName(v string) *CloudDeleteTaskTelRequest {
	s.FileName = &v
	return s
}

func (s *CloudDeleteTaskTelRequest) SetOwnerId(v int64) *CloudDeleteTaskTelRequest {
	s.OwnerId = &v
	return s
}

func (s *CloudDeleteTaskTelRequest) SetResourceOwnerAccount(v string) *CloudDeleteTaskTelRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CloudDeleteTaskTelRequest) SetResourceOwnerId(v int64) *CloudDeleteTaskTelRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CloudDeleteTaskTelRequest) SetStatus(v int64) *CloudDeleteTaskTelRequest {
	s.Status = &v
	return s
}

func (s *CloudDeleteTaskTelRequest) SetTaskId(v int64) *CloudDeleteTaskTelRequest {
	s.TaskId = &v
	return s
}

func (s *CloudDeleteTaskTelRequest) SetTels(v string) *CloudDeleteTaskTelRequest {
	s.Tels = &v
	return s
}

func (s *CloudDeleteTaskTelRequest) Validate() error {
	return dara.Validate(s)
}
