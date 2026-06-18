// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iJoinThirdCallRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountName(v string) *JoinThirdCallRequest
	GetAccountName() *string
	SetCallId(v string) *JoinThirdCallRequest
	GetCallId() *string
	SetClientToken(v string) *JoinThirdCallRequest
	GetClientToken() *string
	SetConnectionId(v string) *JoinThirdCallRequest
	GetConnectionId() *string
	SetHoldConnectionId(v string) *JoinThirdCallRequest
	GetHoldConnectionId() *string
	SetInstanceId(v string) *JoinThirdCallRequest
	GetInstanceId() *string
	SetJobId(v string) *JoinThirdCallRequest
	GetJobId() *string
}

type JoinThirdCallRequest struct {
	// Agent account name (agent logon name)
	//
	// This parameter is required.
	//
	// example:
	//
	// 123@123.com
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// acid from WebSocket after an inbound call
	//
	// example:
	//
	// 7719786
	CallId *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
	// Unique customer request ID used for idempotency validation; can be generated using UUID
	//
	// example:
	//
	// 46c1341e-2648-447a-9b11-70b6a298d94d
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// connId from WebSocket after an inbound call
	//
	// example:
	//
	// 7719788
	ConnectionId *string `json:"ConnectionId,omitempty" xml:"ConnectionId,omitempty"`
	// holdConnId from WebSocket after an inbound call (only provided during two-step transfer)
	//
	// example:
	//
	// 0
	HoldConnectionId *string `json:"HoldConnectionId,omitempty" xml:"HoldConnectionId,omitempty"`
	// AICCS instance ID, visible in the Artificial Intelligence Cloud Call Service console
	//
	// This parameter is required.
	//
	// example:
	//
	// ccc_xp_pre-cn-***
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// jobId from WebSocket after an inbound call
	//
	// example:
	//
	// 7719787
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s JoinThirdCallRequest) String() string {
	return dara.Prettify(s)
}

func (s JoinThirdCallRequest) GoString() string {
	return s.String()
}

func (s *JoinThirdCallRequest) GetAccountName() *string {
	return s.AccountName
}

func (s *JoinThirdCallRequest) GetCallId() *string {
	return s.CallId
}

func (s *JoinThirdCallRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *JoinThirdCallRequest) GetConnectionId() *string {
	return s.ConnectionId
}

func (s *JoinThirdCallRequest) GetHoldConnectionId() *string {
	return s.HoldConnectionId
}

func (s *JoinThirdCallRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *JoinThirdCallRequest) GetJobId() *string {
	return s.JobId
}

func (s *JoinThirdCallRequest) SetAccountName(v string) *JoinThirdCallRequest {
	s.AccountName = &v
	return s
}

func (s *JoinThirdCallRequest) SetCallId(v string) *JoinThirdCallRequest {
	s.CallId = &v
	return s
}

func (s *JoinThirdCallRequest) SetClientToken(v string) *JoinThirdCallRequest {
	s.ClientToken = &v
	return s
}

func (s *JoinThirdCallRequest) SetConnectionId(v string) *JoinThirdCallRequest {
	s.ConnectionId = &v
	return s
}

func (s *JoinThirdCallRequest) SetHoldConnectionId(v string) *JoinThirdCallRequest {
	s.HoldConnectionId = &v
	return s
}

func (s *JoinThirdCallRequest) SetInstanceId(v string) *JoinThirdCallRequest {
	s.InstanceId = &v
	return s
}

func (s *JoinThirdCallRequest) SetJobId(v string) *JoinThirdCallRequest {
	s.JobId = &v
	return s
}

func (s *JoinThirdCallRequest) Validate() error {
	return dara.Validate(s)
}
