// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOutboundPhoneNumberRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountName(v string) *ListOutboundPhoneNumberRequest
	GetAccountName() *string
	SetClientToken(v string) *ListOutboundPhoneNumberRequest
	GetClientToken() *string
	SetInstanceId(v string) *ListOutboundPhoneNumberRequest
	GetInstanceId() *string
}

type ListOutboundPhoneNumberRequest struct {
	// Agent account name, which is the phone number or mailbox entered during account registration. It is unique within the instance.
	//
	// example:
	//
	// username@example.com
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// Unique customer request ID. Used for idempotency validation. You can generate it using UUID.
	//
	// example:
	//
	// 46c1341e-2648-447a-9b11-70b6a298d****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Artificial Intelligence Cloud Call Service (AICCS) instance ID.
	//
	// You can obtain it in the <b>Instance Management</b> section of the left-side navigation pane in the [Artificial Intelligence Cloud Call Service console](https://aiccs.console.aliyun.com/overview).
	//
	// This parameter is required.
	//
	// example:
	//
	// ccc_xp_pre-cn-***
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ListOutboundPhoneNumberRequest) String() string {
	return dara.Prettify(s)
}

func (s ListOutboundPhoneNumberRequest) GoString() string {
	return s.String()
}

func (s *ListOutboundPhoneNumberRequest) GetAccountName() *string {
	return s.AccountName
}

func (s *ListOutboundPhoneNumberRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ListOutboundPhoneNumberRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListOutboundPhoneNumberRequest) SetAccountName(v string) *ListOutboundPhoneNumberRequest {
	s.AccountName = &v
	return s
}

func (s *ListOutboundPhoneNumberRequest) SetClientToken(v string) *ListOutboundPhoneNumberRequest {
	s.ClientToken = &v
	return s
}

func (s *ListOutboundPhoneNumberRequest) SetInstanceId(v string) *ListOutboundPhoneNumberRequest {
	s.InstanceId = &v
	return s
}

func (s *ListOutboundPhoneNumberRequest) Validate() error {
	return dara.Validate(s)
}
