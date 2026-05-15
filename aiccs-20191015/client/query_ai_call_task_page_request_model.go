// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAiCallTaskPageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentName(v string) *QueryAiCallTaskPageRequest
	GetAgentName() *string
	SetApplicationCode(v string) *QueryAiCallTaskPageRequest
	GetApplicationCode() *string
	SetOwnerId(v int64) *QueryAiCallTaskPageRequest
	GetOwnerId() *int64
	SetPageNo(v int64) *QueryAiCallTaskPageRequest
	GetPageNo() *int64
	SetPageSize(v int64) *QueryAiCallTaskPageRequest
	GetPageSize() *int64
	SetResourceOwnerAccount(v string) *QueryAiCallTaskPageRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *QueryAiCallTaskPageRequest
	GetResourceOwnerId() *int64
	SetSource(v int64) *QueryAiCallTaskPageRequest
	GetSource() *int64
	SetStatus(v string) *QueryAiCallTaskPageRequest
	GetStatus() *string
	SetTaskId(v string) *QueryAiCallTaskPageRequest
	GetTaskId() *string
	SetTaskName(v string) *QueryAiCallTaskPageRequest
	GetTaskName() *string
}

type QueryAiCallTaskPageRequest struct {
	// example:
	//
	// 示例值示例值
	AgentName *string `json:"AgentName,omitempty" xml:"AgentName,omitempty"`
	// example:
	//
	// 示例值示例值示例值
	ApplicationCode *string `json:"ApplicationCode,omitempty" xml:"ApplicationCode,omitempty"`
	OwnerId         *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// 1
	PageNo *int64 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// example:
	//
	// 20
	PageSize             *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// example:
	//
	// 68
	Source *int64 `json:"Source,omitempty" xml:"Source,omitempty"`
	// example:
	//
	// INIT
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 112212312*****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// 示例值示例值
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
}

func (s QueryAiCallTaskPageRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryAiCallTaskPageRequest) GoString() string {
	return s.String()
}

func (s *QueryAiCallTaskPageRequest) GetAgentName() *string {
	return s.AgentName
}

func (s *QueryAiCallTaskPageRequest) GetApplicationCode() *string {
	return s.ApplicationCode
}

func (s *QueryAiCallTaskPageRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *QueryAiCallTaskPageRequest) GetPageNo() *int64 {
	return s.PageNo
}

func (s *QueryAiCallTaskPageRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *QueryAiCallTaskPageRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *QueryAiCallTaskPageRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *QueryAiCallTaskPageRequest) GetSource() *int64 {
	return s.Source
}

func (s *QueryAiCallTaskPageRequest) GetStatus() *string {
	return s.Status
}

func (s *QueryAiCallTaskPageRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *QueryAiCallTaskPageRequest) GetTaskName() *string {
	return s.TaskName
}

func (s *QueryAiCallTaskPageRequest) SetAgentName(v string) *QueryAiCallTaskPageRequest {
	s.AgentName = &v
	return s
}

func (s *QueryAiCallTaskPageRequest) SetApplicationCode(v string) *QueryAiCallTaskPageRequest {
	s.ApplicationCode = &v
	return s
}

func (s *QueryAiCallTaskPageRequest) SetOwnerId(v int64) *QueryAiCallTaskPageRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryAiCallTaskPageRequest) SetPageNo(v int64) *QueryAiCallTaskPageRequest {
	s.PageNo = &v
	return s
}

func (s *QueryAiCallTaskPageRequest) SetPageSize(v int64) *QueryAiCallTaskPageRequest {
	s.PageSize = &v
	return s
}

func (s *QueryAiCallTaskPageRequest) SetResourceOwnerAccount(v string) *QueryAiCallTaskPageRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryAiCallTaskPageRequest) SetResourceOwnerId(v int64) *QueryAiCallTaskPageRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QueryAiCallTaskPageRequest) SetSource(v int64) *QueryAiCallTaskPageRequest {
	s.Source = &v
	return s
}

func (s *QueryAiCallTaskPageRequest) SetStatus(v string) *QueryAiCallTaskPageRequest {
	s.Status = &v
	return s
}

func (s *QueryAiCallTaskPageRequest) SetTaskId(v string) *QueryAiCallTaskPageRequest {
	s.TaskId = &v
	return s
}

func (s *QueryAiCallTaskPageRequest) SetTaskName(v string) *QueryAiCallTaskPageRequest {
	s.TaskName = &v
	return s
}

func (s *QueryAiCallTaskPageRequest) Validate() error {
	return dara.Validate(s)
}
