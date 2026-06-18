// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAnswerCallRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountName(v string) *AnswerCallRequest
	GetAccountName() *string
	SetCallId(v string) *AnswerCallRequest
	GetCallId() *string
	SetClientToken(v string) *AnswerCallRequest
	GetClientToken() *string
	SetConnectionId(v string) *AnswerCallRequest
	GetConnectionId() *string
	SetInstanceId(v string) *AnswerCallRequest
	GetInstanceId() *string
	SetJobId(v string) *AnswerCallRequest
	GetJobId() *string
}

type AnswerCallRequest struct {
	// Agent account name (agent logon name).
	//
	// This parameter is required.
	//
	// example:
	//
	// 123@123.com
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The acid from the WebSocket after an inbound call.
	//
	// example:
	//
	// 7719786
	CallId *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
	// Unique ID for the customer request. Used for idempotency validation. You can generate it using UUID.
	//
	// example:
	//
	// 46c1341e-2648-447a-9b11-70b6a298d94d
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The connId from the WebSocket after an inbound call.
	//
	// example:
	//
	// 7719788
	ConnectionId *string `json:"ConnectionId,omitempty" xml:"ConnectionId,omitempty"`
	// AICCS instance ID. You can obtain it from the Artificial Intelligence Cloud Call Service console.
	//
	// This parameter is required.
	//
	// example:
	//
	// ccc_xp_pre-cn-***
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The jobId from the WebSocket after an inbound call.
	//
	// example:
	//
	// 7719787
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s AnswerCallRequest) String() string {
	return dara.Prettify(s)
}

func (s AnswerCallRequest) GoString() string {
	return s.String()
}

func (s *AnswerCallRequest) GetAccountName() *string {
	return s.AccountName
}

func (s *AnswerCallRequest) GetCallId() *string {
	return s.CallId
}

func (s *AnswerCallRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *AnswerCallRequest) GetConnectionId() *string {
	return s.ConnectionId
}

func (s *AnswerCallRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *AnswerCallRequest) GetJobId() *string {
	return s.JobId
}

func (s *AnswerCallRequest) SetAccountName(v string) *AnswerCallRequest {
	s.AccountName = &v
	return s
}

func (s *AnswerCallRequest) SetCallId(v string) *AnswerCallRequest {
	s.CallId = &v
	return s
}

func (s *AnswerCallRequest) SetClientToken(v string) *AnswerCallRequest {
	s.ClientToken = &v
	return s
}

func (s *AnswerCallRequest) SetConnectionId(v string) *AnswerCallRequest {
	s.ConnectionId = &v
	return s
}

func (s *AnswerCallRequest) SetInstanceId(v string) *AnswerCallRequest {
	s.InstanceId = &v
	return s
}

func (s *AnswerCallRequest) SetJobId(v string) *AnswerCallRequest {
	s.JobId = &v
	return s
}

func (s *AnswerCallRequest) Validate() error {
	return dara.Validate(s)
}
