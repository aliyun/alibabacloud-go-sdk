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
	// This parameter is required.
	//
	// example:
	//
	// 123@123.com
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
