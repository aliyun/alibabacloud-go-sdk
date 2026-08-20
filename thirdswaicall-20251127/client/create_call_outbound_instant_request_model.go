// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCallOutboundInstantRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCalledNumber(v string) *CreateCallOutboundInstantRequest
	GetCalledNumber() *string
	SetCallerNumber(v string) *CreateCallOutboundInstantRequest
	GetCallerNumber() *string
	SetCallerUacAccountId(v string) *CreateCallOutboundInstantRequest
	GetCallerUacAccountId() *string
	SetCurrentWorkspaceId(v string) *CreateCallOutboundInstantRequest
	GetCurrentWorkspaceId() *string
	SetCustomerLineCode(v string) *CreateCallOutboundInstantRequest
	GetCustomerLineCode() *string
	SetCustomerName(v string) *CreateCallOutboundInstantRequest
	GetCustomerName() *string
	SetEncryptCall(v bool) *CreateCallOutboundInstantRequest
	GetEncryptCall() *bool
	SetPromptVariables(v string) *CreateCallOutboundInstantRequest
	GetPromptVariables() *string
	SetTaskId(v int64) *CreateCallOutboundInstantRequest
	GetTaskId() *int64
}

type CreateCallOutboundInstantRequest struct {
	// example:
	//
	// 13800138000
	CalledNumber *string `json:"CalledNumber,omitempty" xml:"CalledNumber,omitempty"`
	// example:
	//
	// 057188888888
	CallerNumber *string `json:"CallerNumber,omitempty" xml:"CallerNumber,omitempty"`
	// example:
	//
	// abc123***
	CallerUacAccountId *string `json:"CallerUacAccountId,omitempty" xml:"CallerUacAccountId,omitempty"`
	// example:
	//
	// abc123***
	CurrentWorkspaceId *string `json:"CurrentWorkspaceId,omitempty" xml:"CurrentWorkspaceId,omitempty"`
	// example:
	//
	// line_001
	CustomerLineCode *string `json:"CustomerLineCode,omitempty" xml:"CustomerLineCode,omitempty"`
	// example:
	//
	// 张三
	CustomerName *string `json:"CustomerName,omitempty" xml:"CustomerName,omitempty"`
	// example:
	//
	// false
	EncryptCall *bool `json:"EncryptCall,omitempty" xml:"EncryptCall,omitempty"`
	// example:
	//
	// {"start":"2220"}
	PromptVariables *string `json:"PromptVariables,omitempty" xml:"PromptVariables,omitempty"`
	// example:
	//
	// 132
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CreateCallOutboundInstantRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateCallOutboundInstantRequest) GoString() string {
	return s.String()
}

func (s *CreateCallOutboundInstantRequest) GetCalledNumber() *string {
	return s.CalledNumber
}

func (s *CreateCallOutboundInstantRequest) GetCallerNumber() *string {
	return s.CallerNumber
}

func (s *CreateCallOutboundInstantRequest) GetCallerUacAccountId() *string {
	return s.CallerUacAccountId
}

func (s *CreateCallOutboundInstantRequest) GetCurrentWorkspaceId() *string {
	return s.CurrentWorkspaceId
}

func (s *CreateCallOutboundInstantRequest) GetCustomerLineCode() *string {
	return s.CustomerLineCode
}

func (s *CreateCallOutboundInstantRequest) GetCustomerName() *string {
	return s.CustomerName
}

func (s *CreateCallOutboundInstantRequest) GetEncryptCall() *bool {
	return s.EncryptCall
}

func (s *CreateCallOutboundInstantRequest) GetPromptVariables() *string {
	return s.PromptVariables
}

func (s *CreateCallOutboundInstantRequest) GetTaskId() *int64 {
	return s.TaskId
}

func (s *CreateCallOutboundInstantRequest) SetCalledNumber(v string) *CreateCallOutboundInstantRequest {
	s.CalledNumber = &v
	return s
}

func (s *CreateCallOutboundInstantRequest) SetCallerNumber(v string) *CreateCallOutboundInstantRequest {
	s.CallerNumber = &v
	return s
}

func (s *CreateCallOutboundInstantRequest) SetCallerUacAccountId(v string) *CreateCallOutboundInstantRequest {
	s.CallerUacAccountId = &v
	return s
}

func (s *CreateCallOutboundInstantRequest) SetCurrentWorkspaceId(v string) *CreateCallOutboundInstantRequest {
	s.CurrentWorkspaceId = &v
	return s
}

func (s *CreateCallOutboundInstantRequest) SetCustomerLineCode(v string) *CreateCallOutboundInstantRequest {
	s.CustomerLineCode = &v
	return s
}

func (s *CreateCallOutboundInstantRequest) SetCustomerName(v string) *CreateCallOutboundInstantRequest {
	s.CustomerName = &v
	return s
}

func (s *CreateCallOutboundInstantRequest) SetEncryptCall(v bool) *CreateCallOutboundInstantRequest {
	s.EncryptCall = &v
	return s
}

func (s *CreateCallOutboundInstantRequest) SetPromptVariables(v string) *CreateCallOutboundInstantRequest {
	s.PromptVariables = &v
	return s
}

func (s *CreateCallOutboundInstantRequest) SetTaskId(v int64) *CreateCallOutboundInstantRequest {
	s.TaskId = &v
	return s
}

func (s *CreateCallOutboundInstantRequest) Validate() error {
	return dara.Validate(s)
}
