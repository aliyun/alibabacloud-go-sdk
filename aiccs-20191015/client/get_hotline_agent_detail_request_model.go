// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHotlineAgentDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountName(v string) *GetHotlineAgentDetailRequest
	GetAccountName() *string
	SetClientToken(v string) *GetHotlineAgentDetailRequest
	GetClientToken() *string
	SetInstanceId(v string) *GetHotlineAgentDetailRequest
	GetInstanceId() *string
}

type GetHotlineAgentDetailRequest struct {
	// Agent account name, which is the phone number or mailbox specified during account registration. It is unique within the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// username@example.com
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// Unique ID of the customer request. Used for idempotency validation. You can generate it by using a UUID.
	//
	// example:
	//
	// 46c1341e-2648-447a-9b11-70b6a298d94d****
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

func (s GetHotlineAgentDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s GetHotlineAgentDetailRequest) GoString() string {
	return s.String()
}

func (s *GetHotlineAgentDetailRequest) GetAccountName() *string {
	return s.AccountName
}

func (s *GetHotlineAgentDetailRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *GetHotlineAgentDetailRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetHotlineAgentDetailRequest) SetAccountName(v string) *GetHotlineAgentDetailRequest {
	s.AccountName = &v
	return s
}

func (s *GetHotlineAgentDetailRequest) SetClientToken(v string) *GetHotlineAgentDetailRequest {
	s.ClientToken = &v
	return s
}

func (s *GetHotlineAgentDetailRequest) SetInstanceId(v string) *GetHotlineAgentDetailRequest {
	s.InstanceId = &v
	return s
}

func (s *GetHotlineAgentDetailRequest) Validate() error {
	return dara.Validate(s)
}
