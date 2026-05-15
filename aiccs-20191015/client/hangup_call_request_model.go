// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHangupCallRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountName(v string) *HangupCallRequest
	GetAccountName() *string
	SetCallId(v string) *HangupCallRequest
	GetCallId() *string
	SetClientToken(v string) *HangupCallRequest
	GetClientToken() *string
	SetConnectionId(v string) *HangupCallRequest
	GetConnectionId() *string
	SetInstanceId(v string) *HangupCallRequest
	GetInstanceId() *string
	SetJobId(v string) *HangupCallRequest
	GetJobId() *string
}

type HangupCallRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 123@123.com
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// example:
	//
	// 7719786
	CallId *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
	// example:
	//
	// 46c1341e-2648-447a-9b11-70b6a298d94d
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// example:
	//
	// 7719788
	ConnectionId *string `json:"ConnectionId,omitempty" xml:"ConnectionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ccc_xp_pre-cn-***
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 7719787
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s HangupCallRequest) String() string {
	return dara.Prettify(s)
}

func (s HangupCallRequest) GoString() string {
	return s.String()
}

func (s *HangupCallRequest) GetAccountName() *string {
	return s.AccountName
}

func (s *HangupCallRequest) GetCallId() *string {
	return s.CallId
}

func (s *HangupCallRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *HangupCallRequest) GetConnectionId() *string {
	return s.ConnectionId
}

func (s *HangupCallRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *HangupCallRequest) GetJobId() *string {
	return s.JobId
}

func (s *HangupCallRequest) SetAccountName(v string) *HangupCallRequest {
	s.AccountName = &v
	return s
}

func (s *HangupCallRequest) SetCallId(v string) *HangupCallRequest {
	s.CallId = &v
	return s
}

func (s *HangupCallRequest) SetClientToken(v string) *HangupCallRequest {
	s.ClientToken = &v
	return s
}

func (s *HangupCallRequest) SetConnectionId(v string) *HangupCallRequest {
	s.ConnectionId = &v
	return s
}

func (s *HangupCallRequest) SetInstanceId(v string) *HangupCallRequest {
	s.InstanceId = &v
	return s
}

func (s *HangupCallRequest) SetJobId(v string) *HangupCallRequest {
	s.JobId = &v
	return s
}

func (s *HangupCallRequest) Validate() error {
	return dara.Validate(s)
}
