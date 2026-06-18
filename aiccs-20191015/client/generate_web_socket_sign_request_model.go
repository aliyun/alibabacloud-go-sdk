// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateWebSocketSignRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountName(v string) *GenerateWebSocketSignRequest
	GetAccountName() *string
	SetClientToken(v string) *GenerateWebSocketSignRequest
	GetClientToken() *string
	SetInstanceId(v string) *GenerateWebSocketSignRequest
	GetInstanceId() *string
}

type GenerateWebSocketSignRequest struct {
	// Agent account name, which is the phone number or mailbox entered during account registration. It is unique within the instance.
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
	// 46c1341e-2648-447a-****-70b6a298d94d
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

func (s GenerateWebSocketSignRequest) String() string {
	return dara.Prettify(s)
}

func (s GenerateWebSocketSignRequest) GoString() string {
	return s.String()
}

func (s *GenerateWebSocketSignRequest) GetAccountName() *string {
	return s.AccountName
}

func (s *GenerateWebSocketSignRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *GenerateWebSocketSignRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GenerateWebSocketSignRequest) SetAccountName(v string) *GenerateWebSocketSignRequest {
	s.AccountName = &v
	return s
}

func (s *GenerateWebSocketSignRequest) SetClientToken(v string) *GenerateWebSocketSignRequest {
	s.ClientToken = &v
	return s
}

func (s *GenerateWebSocketSignRequest) SetInstanceId(v string) *GenerateWebSocketSignRequest {
	s.InstanceId = &v
	return s
}

func (s *GenerateWebSocketSignRequest) Validate() error {
	return dara.Validate(s)
}
