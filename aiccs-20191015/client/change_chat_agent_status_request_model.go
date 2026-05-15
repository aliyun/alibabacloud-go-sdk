// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeChatAgentStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountName(v string) *ChangeChatAgentStatusRequest
	GetAccountName() *string
	SetClientToken(v string) *ChangeChatAgentStatusRequest
	GetClientToken() *string
	SetInstanceId(v string) *ChangeChatAgentStatusRequest
	GetInstanceId() *string
	SetMethod(v string) *ChangeChatAgentStatusRequest
	GetMethod() *string
}

type ChangeChatAgentStatusRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// account1
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// example:
	//
	// 46c1341e-2648-447a-9b11-70b6a298d94d
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ccc_xp_pre-cn-***
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// requestLogout
	Method *string `json:"Method,omitempty" xml:"Method,omitempty"`
}

func (s ChangeChatAgentStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s ChangeChatAgentStatusRequest) GoString() string {
	return s.String()
}

func (s *ChangeChatAgentStatusRequest) GetAccountName() *string {
	return s.AccountName
}

func (s *ChangeChatAgentStatusRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ChangeChatAgentStatusRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ChangeChatAgentStatusRequest) GetMethod() *string {
	return s.Method
}

func (s *ChangeChatAgentStatusRequest) SetAccountName(v string) *ChangeChatAgentStatusRequest {
	s.AccountName = &v
	return s
}

func (s *ChangeChatAgentStatusRequest) SetClientToken(v string) *ChangeChatAgentStatusRequest {
	s.ClientToken = &v
	return s
}

func (s *ChangeChatAgentStatusRequest) SetInstanceId(v string) *ChangeChatAgentStatusRequest {
	s.InstanceId = &v
	return s
}

func (s *ChangeChatAgentStatusRequest) SetMethod(v string) *ChangeChatAgentStatusRequest {
	s.Method = &v
	return s
}

func (s *ChangeChatAgentStatusRequest) Validate() error {
	return dara.Validate(s)
}
