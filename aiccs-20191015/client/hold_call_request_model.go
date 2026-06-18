// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHoldCallRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountName(v string) *HoldCallRequest
	GetAccountName() *string
	SetCallId(v string) *HoldCallRequest
	GetCallId() *string
	SetClientToken(v string) *HoldCallRequest
	GetClientToken() *string
	SetConnectionId(v string) *HoldCallRequest
	GetConnectionId() *string
	SetInstanceId(v string) *HoldCallRequest
	GetInstanceId() *string
	SetJobId(v string) *HoldCallRequest
	GetJobId() *string
}

type HoldCallRequest struct {
	// Agent account name (agent logon name)
	//
	// This parameter is required.
	//
	// example:
	//
	// 123@123.com
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// acid in WebSocket after an inbound call
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
	// connId in WebSocket after an inbound call
	//
	// example:
	//
	// 7719788
	ConnectionId *string `json:"ConnectionId,omitempty" xml:"ConnectionId,omitempty"`
	// Artificial Intelligence Cloud Call Service (AICCS) instance ID, visible in the Artificial Intelligence Cloud Call Service console
	//
	// This parameter is required.
	//
	// example:
	//
	// ccc_xp_pre-cn-***
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// jobId in WebSocket after an inbound call
	//
	// example:
	//
	// 7719787
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s HoldCallRequest) String() string {
	return dara.Prettify(s)
}

func (s HoldCallRequest) GoString() string {
	return s.String()
}

func (s *HoldCallRequest) GetAccountName() *string {
	return s.AccountName
}

func (s *HoldCallRequest) GetCallId() *string {
	return s.CallId
}

func (s *HoldCallRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *HoldCallRequest) GetConnectionId() *string {
	return s.ConnectionId
}

func (s *HoldCallRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *HoldCallRequest) GetJobId() *string {
	return s.JobId
}

func (s *HoldCallRequest) SetAccountName(v string) *HoldCallRequest {
	s.AccountName = &v
	return s
}

func (s *HoldCallRequest) SetCallId(v string) *HoldCallRequest {
	s.CallId = &v
	return s
}

func (s *HoldCallRequest) SetClientToken(v string) *HoldCallRequest {
	s.ClientToken = &v
	return s
}

func (s *HoldCallRequest) SetConnectionId(v string) *HoldCallRequest {
	s.ConnectionId = &v
	return s
}

func (s *HoldCallRequest) SetInstanceId(v string) *HoldCallRequest {
	s.InstanceId = &v
	return s
}

func (s *HoldCallRequest) SetJobId(v string) *HoldCallRequest {
	s.JobId = &v
	return s
}

func (s *HoldCallRequest) Validate() error {
	return dara.Validate(s)
}
