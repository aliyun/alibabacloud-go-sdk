// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountName(v string) *GetAgentRequest
	GetAccountName() *string
	SetClientToken(v string) *GetAgentRequest
	GetClientToken() *string
	SetInstanceId(v string) *GetAgentRequest
	GetInstanceId() *string
}

type GetAgentRequest struct {
	// Agent account name, which is the phone number or mailbox entered during account registration. Unique within the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// username@example.com
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// Unique customer request ID. Used for idempotency validation. You can generate it using UUID.
	//
	// example:
	//
	// 46c1341e-2648-447a-9b11-70b6a298d94d
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Artificial Intelligence Cloud Call Service (AICCS) instance ID.
	//
	// You can obtain it from **Instance Management*	- in the left-side navigation pane of the [Artificial Intelligence Cloud Call Service console](https://aiccs.console.aliyun.com/overview).
	//
	// This parameter is required.
	//
	// example:
	//
	// ccc_xp_pre-cn-***
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetAgentRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAgentRequest) GoString() string {
	return s.String()
}

func (s *GetAgentRequest) GetAccountName() *string {
	return s.AccountName
}

func (s *GetAgentRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *GetAgentRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetAgentRequest) SetAccountName(v string) *GetAgentRequest {
	s.AccountName = &v
	return s
}

func (s *GetAgentRequest) SetClientToken(v string) *GetAgentRequest {
	s.ClientToken = &v
	return s
}

func (s *GetAgentRequest) SetInstanceId(v string) *GetAgentRequest {
	s.InstanceId = &v
	return s
}

func (s *GetAgentRequest) Validate() error {
	return dara.Validate(s)
}
