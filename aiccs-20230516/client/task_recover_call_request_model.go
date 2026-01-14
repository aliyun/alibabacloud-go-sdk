// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTaskRecoverCallRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBeginImportTime(v string) *TaskRecoverCallRequest
	GetBeginImportTime() *string
	SetEndImportTime(v string) *TaskRecoverCallRequest
	GetEndImportTime() *string
	SetNumbers(v []*string) *TaskRecoverCallRequest
	GetNumbers() []*string
	SetOwnerId(v int64) *TaskRecoverCallRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *TaskRecoverCallRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *TaskRecoverCallRequest
	GetResourceOwnerId() *int64
	SetTags(v []*string) *TaskRecoverCallRequest
	GetTags() []*string
	SetTaskId(v int64) *TaskRecoverCallRequest
	GetTaskId() *int64
}

type TaskRecoverCallRequest struct {
	// 查询开始导入时间
	//
	// example:
	//
	// "2023-01-09 18:58:19"
	BeginImportTime *string `json:"BeginImportTime,omitempty" xml:"BeginImportTime,omitempty"`
	// 查询结束导入时间
	//
	// example:
	//
	// "2023-01-09 18:58:19"
	EndImportTime *string `json:"EndImportTime,omitempty" xml:"EndImportTime,omitempty"`
	// 号码列表
	Numbers              []*string `json:"Numbers,omitempty" xml:"Numbers,omitempty" type:"Repeated"`
	OwnerId              *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64    `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// 用户自定义标签列表
	Tags []*string `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// 任务ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 93
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s TaskRecoverCallRequest) String() string {
	return dara.Prettify(s)
}

func (s TaskRecoverCallRequest) GoString() string {
	return s.String()
}

func (s *TaskRecoverCallRequest) GetBeginImportTime() *string {
	return s.BeginImportTime
}

func (s *TaskRecoverCallRequest) GetEndImportTime() *string {
	return s.EndImportTime
}

func (s *TaskRecoverCallRequest) GetNumbers() []*string {
	return s.Numbers
}

func (s *TaskRecoverCallRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *TaskRecoverCallRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *TaskRecoverCallRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *TaskRecoverCallRequest) GetTags() []*string {
	return s.Tags
}

func (s *TaskRecoverCallRequest) GetTaskId() *int64 {
	return s.TaskId
}

func (s *TaskRecoverCallRequest) SetBeginImportTime(v string) *TaskRecoverCallRequest {
	s.BeginImportTime = &v
	return s
}

func (s *TaskRecoverCallRequest) SetEndImportTime(v string) *TaskRecoverCallRequest {
	s.EndImportTime = &v
	return s
}

func (s *TaskRecoverCallRequest) SetNumbers(v []*string) *TaskRecoverCallRequest {
	s.Numbers = v
	return s
}

func (s *TaskRecoverCallRequest) SetOwnerId(v int64) *TaskRecoverCallRequest {
	s.OwnerId = &v
	return s
}

func (s *TaskRecoverCallRequest) SetResourceOwnerAccount(v string) *TaskRecoverCallRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *TaskRecoverCallRequest) SetResourceOwnerId(v int64) *TaskRecoverCallRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *TaskRecoverCallRequest) SetTags(v []*string) *TaskRecoverCallRequest {
	s.Tags = v
	return s
}

func (s *TaskRecoverCallRequest) SetTaskId(v int64) *TaskRecoverCallRequest {
	s.TaskId = &v
	return s
}

func (s *TaskRecoverCallRequest) Validate() error {
	return dara.Validate(s)
}
