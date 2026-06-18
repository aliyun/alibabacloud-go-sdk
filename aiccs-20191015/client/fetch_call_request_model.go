// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFetchCallRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountName(v string) *FetchCallRequest
	GetAccountName() *string
	SetCallId(v string) *FetchCallRequest
	GetCallId() *string
	SetClientToken(v string) *FetchCallRequest
	GetClientToken() *string
	SetConnectionId(v string) *FetchCallRequest
	GetConnectionId() *string
	SetHoldConnectionId(v string) *FetchCallRequest
	GetHoldConnectionId() *string
	SetInstanceId(v string) *FetchCallRequest
	GetInstanceId() *string
	SetJobId(v string) *FetchCallRequest
	GetJobId() *string
}

type FetchCallRequest struct {
	// Agent account name (agent logon name)
	//
	// This parameter is required.
	//
	// example:
	//
	// 123@123.com
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The acid from WebSocket after an inbound call
	//
	// example:
	//
	// 7719786
	CallId *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
	// Unique ID of the customer request, used for idempotency validation; can be generated using UUID
	//
	// example:
	//
	// 46c1341e-2648-447a-9b11-70b6a298d94d
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The connId from WebSocket after an inbound call
	//
	// example:
	//
	// 7719788
	ConnectionId *string `json:"ConnectionId,omitempty" xml:"ConnectionId,omitempty"`
	// The holdConnId from WebSocket after an inbound call (only provided during two-step transfer)
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
	// The jobId from WebSocket after an inbound call
	//
	// example:
	//
	// 7719787
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s FetchCallRequest) String() string {
	return dara.Prettify(s)
}

func (s FetchCallRequest) GoString() string {
	return s.String()
}

func (s *FetchCallRequest) GetAccountName() *string {
	return s.AccountName
}

func (s *FetchCallRequest) GetCallId() *string {
	return s.CallId
}

func (s *FetchCallRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *FetchCallRequest) GetConnectionId() *string {
	return s.ConnectionId
}

func (s *FetchCallRequest) GetHoldConnectionId() *string {
	return s.HoldConnectionId
}

func (s *FetchCallRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *FetchCallRequest) GetJobId() *string {
	return s.JobId
}

func (s *FetchCallRequest) SetAccountName(v string) *FetchCallRequest {
	s.AccountName = &v
	return s
}

func (s *FetchCallRequest) SetCallId(v string) *FetchCallRequest {
	s.CallId = &v
	return s
}

func (s *FetchCallRequest) SetClientToken(v string) *FetchCallRequest {
	s.ClientToken = &v
	return s
}

func (s *FetchCallRequest) SetConnectionId(v string) *FetchCallRequest {
	s.ConnectionId = &v
	return s
}

func (s *FetchCallRequest) SetHoldConnectionId(v string) *FetchCallRequest {
	s.HoldConnectionId = &v
	return s
}

func (s *FetchCallRequest) SetInstanceId(v string) *FetchCallRequest {
	s.InstanceId = &v
	return s
}

func (s *FetchCallRequest) SetJobId(v string) *FetchCallRequest {
	s.JobId = &v
	return s
}

func (s *FetchCallRequest) Validate() error {
	return dara.Validate(s)
}
