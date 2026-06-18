// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHotlineAgentStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountName(v string) *GetHotlineAgentStatusRequest
	GetAccountName() *string
	SetInstanceId(v string) *GetHotlineAgentStatusRequest
	GetInstanceId() *string
}

type GetHotlineAgentStatusRequest struct {
	// Agent account name, which is the phone number or mailbox entered during account registration. It is unique within the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// username@example.com
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
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

func (s GetHotlineAgentStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s GetHotlineAgentStatusRequest) GoString() string {
	return s.String()
}

func (s *GetHotlineAgentStatusRequest) GetAccountName() *string {
	return s.AccountName
}

func (s *GetHotlineAgentStatusRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetHotlineAgentStatusRequest) SetAccountName(v string) *GetHotlineAgentStatusRequest {
	s.AccountName = &v
	return s
}

func (s *GetHotlineAgentStatusRequest) SetInstanceId(v string) *GetHotlineAgentStatusRequest {
	s.InstanceId = &v
	return s
}

func (s *GetHotlineAgentStatusRequest) Validate() error {
	return dara.Validate(s)
}
