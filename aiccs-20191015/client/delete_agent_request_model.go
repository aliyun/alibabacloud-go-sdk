// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAgentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountName(v string) *DeleteAgentRequest
	GetAccountName() *string
	SetClientToken(v string) *DeleteAgentRequest
	GetClientToken() *string
	SetInstanceId(v string) *DeleteAgentRequest
	GetInstanceId() *string
}

type DeleteAgentRequest struct {
	// Agent account name, which is the phone number or mailbox specified during account registration. It is unique within the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// username@example.com
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// Unique ID for the customer request. Used for idempotency validation. You can generate it using UUID.
	//
	// example:
	//
	// 46c1341e-2648-447a-9b11-70b6a298****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// AICCS instance ID.
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

func (s DeleteAgentRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAgentRequest) GoString() string {
	return s.String()
}

func (s *DeleteAgentRequest) GetAccountName() *string {
	return s.AccountName
}

func (s *DeleteAgentRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteAgentRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteAgentRequest) SetAccountName(v string) *DeleteAgentRequest {
	s.AccountName = &v
	return s
}

func (s *DeleteAgentRequest) SetClientToken(v string) *DeleteAgentRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteAgentRequest) SetInstanceId(v string) *DeleteAgentRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteAgentRequest) Validate() error {
	return dara.Validate(s)
}
