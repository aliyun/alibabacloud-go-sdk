// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHangupThirdCallRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountName(v string) *HangupThirdCallRequest
	GetAccountName() *string
	SetCallId(v string) *HangupThirdCallRequest
	GetCallId() *string
	SetClientToken(v string) *HangupThirdCallRequest
	GetClientToken() *string
	SetConnectionId(v string) *HangupThirdCallRequest
	GetConnectionId() *string
	SetInstanceId(v string) *HangupThirdCallRequest
	GetInstanceId() *string
	SetJobId(v string) *HangupThirdCallRequest
	GetJobId() *string
}

type HangupThirdCallRequest struct {
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

func (s HangupThirdCallRequest) String() string {
	return dara.Prettify(s)
}

func (s HangupThirdCallRequest) GoString() string {
	return s.String()
}

func (s *HangupThirdCallRequest) GetAccountName() *string {
	return s.AccountName
}

func (s *HangupThirdCallRequest) GetCallId() *string {
	return s.CallId
}

func (s *HangupThirdCallRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *HangupThirdCallRequest) GetConnectionId() *string {
	return s.ConnectionId
}

func (s *HangupThirdCallRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *HangupThirdCallRequest) GetJobId() *string {
	return s.JobId
}

func (s *HangupThirdCallRequest) SetAccountName(v string) *HangupThirdCallRequest {
	s.AccountName = &v
	return s
}

func (s *HangupThirdCallRequest) SetCallId(v string) *HangupThirdCallRequest {
	s.CallId = &v
	return s
}

func (s *HangupThirdCallRequest) SetClientToken(v string) *HangupThirdCallRequest {
	s.ClientToken = &v
	return s
}

func (s *HangupThirdCallRequest) SetConnectionId(v string) *HangupThirdCallRequest {
	s.ConnectionId = &v
	return s
}

func (s *HangupThirdCallRequest) SetInstanceId(v string) *HangupThirdCallRequest {
	s.InstanceId = &v
	return s
}

func (s *HangupThirdCallRequest) SetJobId(v string) *HangupThirdCallRequest {
	s.JobId = &v
	return s
}

func (s *HangupThirdCallRequest) Validate() error {
	return dara.Validate(s)
}
