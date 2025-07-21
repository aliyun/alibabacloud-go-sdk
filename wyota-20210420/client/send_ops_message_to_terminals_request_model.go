// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendOpsMessageToTerminalsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDelay(v bool) *SendOpsMessageToTerminalsRequest
	GetDelay() *bool
	SetMsg(v string) *SendOpsMessageToTerminalsRequest
	GetMsg() *string
	SetOpsAction(v string) *SendOpsMessageToTerminalsRequest
	GetOpsAction() *string
	SetUuids(v []*string) *SendOpsMessageToTerminalsRequest
	GetUuids() []*string
	SetWaitForAck(v bool) *SendOpsMessageToTerminalsRequest
	GetWaitForAck() *bool
}

type SendOpsMessageToTerminalsRequest struct {
	Delay      *bool     `json:"Delay,omitempty" xml:"Delay,omitempty"`
	Msg        *string   `json:"Msg,omitempty" xml:"Msg,omitempty"`
	OpsAction  *string   `json:"OpsAction,omitempty" xml:"OpsAction,omitempty"`
	Uuids      []*string `json:"Uuids,omitempty" xml:"Uuids,omitempty" type:"Repeated"`
	WaitForAck *bool     `json:"WaitForAck,omitempty" xml:"WaitForAck,omitempty"`
}

func (s SendOpsMessageToTerminalsRequest) String() string {
	return dara.Prettify(s)
}

func (s SendOpsMessageToTerminalsRequest) GoString() string {
	return s.String()
}

func (s *SendOpsMessageToTerminalsRequest) GetDelay() *bool {
	return s.Delay
}

func (s *SendOpsMessageToTerminalsRequest) GetMsg() *string {
	return s.Msg
}

func (s *SendOpsMessageToTerminalsRequest) GetOpsAction() *string {
	return s.OpsAction
}

func (s *SendOpsMessageToTerminalsRequest) GetUuids() []*string {
	return s.Uuids
}

func (s *SendOpsMessageToTerminalsRequest) GetWaitForAck() *bool {
	return s.WaitForAck
}

func (s *SendOpsMessageToTerminalsRequest) SetDelay(v bool) *SendOpsMessageToTerminalsRequest {
	s.Delay = &v
	return s
}

func (s *SendOpsMessageToTerminalsRequest) SetMsg(v string) *SendOpsMessageToTerminalsRequest {
	s.Msg = &v
	return s
}

func (s *SendOpsMessageToTerminalsRequest) SetOpsAction(v string) *SendOpsMessageToTerminalsRequest {
	s.OpsAction = &v
	return s
}

func (s *SendOpsMessageToTerminalsRequest) SetUuids(v []*string) *SendOpsMessageToTerminalsRequest {
	s.Uuids = v
	return s
}

func (s *SendOpsMessageToTerminalsRequest) SetWaitForAck(v bool) *SendOpsMessageToTerminalsRequest {
	s.WaitForAck = &v
	return s
}

func (s *SendOpsMessageToTerminalsRequest) Validate() error {
	return dara.Validate(s)
}
