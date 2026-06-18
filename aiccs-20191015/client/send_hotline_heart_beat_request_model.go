// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendHotlineHeartBeatRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountName(v string) *SendHotlineHeartBeatRequest
	GetAccountName() *string
	SetClientToken(v string) *SendHotlineHeartBeatRequest
	GetClientToken() *string
	SetInstanceId(v string) *SendHotlineHeartBeatRequest
	GetInstanceId() *string
	SetToken(v string) *SendHotlineHeartBeatRequest
	GetToken() *string
}

type SendHotlineHeartBeatRequest struct {
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
	// 46c1341e-2648-447a-9b11-70b6a298d94d
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
	// Heartbeat signature.
	//
	// You can obtain the token by invoking the [StartHotlineService](https://help.aliyun.com/document_detail/2718045.html) API.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0079e7a845e373****
	Token *string `json:"Token,omitempty" xml:"Token,omitempty"`
}

func (s SendHotlineHeartBeatRequest) String() string {
	return dara.Prettify(s)
}

func (s SendHotlineHeartBeatRequest) GoString() string {
	return s.String()
}

func (s *SendHotlineHeartBeatRequest) GetAccountName() *string {
	return s.AccountName
}

func (s *SendHotlineHeartBeatRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *SendHotlineHeartBeatRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *SendHotlineHeartBeatRequest) GetToken() *string {
	return s.Token
}

func (s *SendHotlineHeartBeatRequest) SetAccountName(v string) *SendHotlineHeartBeatRequest {
	s.AccountName = &v
	return s
}

func (s *SendHotlineHeartBeatRequest) SetClientToken(v string) *SendHotlineHeartBeatRequest {
	s.ClientToken = &v
	return s
}

func (s *SendHotlineHeartBeatRequest) SetInstanceId(v string) *SendHotlineHeartBeatRequest {
	s.InstanceId = &v
	return s
}

func (s *SendHotlineHeartBeatRequest) SetToken(v string) *SendHotlineHeartBeatRequest {
	s.Token = &v
	return s
}

func (s *SendHotlineHeartBeatRequest) Validate() error {
	return dara.Validate(s)
}
