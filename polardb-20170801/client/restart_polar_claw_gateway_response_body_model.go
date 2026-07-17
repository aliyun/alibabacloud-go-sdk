// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestartPolarClawGatewayResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *RestartPolarClawGatewayResponseBody
	GetApplicationId() *string
	SetCode(v int32) *RestartPolarClawGatewayResponseBody
	GetCode() *int32
	SetDowntimeMs(v int64) *RestartPolarClawGatewayResponseBody
	GetDowntimeMs() *int64
	SetGatewayVersion(v string) *RestartPolarClawGatewayResponseBody
	GetGatewayVersion() *string
	SetMessage(v string) *RestartPolarClawGatewayResponseBody
	GetMessage() *string
	SetMode(v string) *RestartPolarClawGatewayResponseBody
	GetMode() *string
	SetOk(v bool) *RestartPolarClawGatewayResponseBody
	GetOk() *bool
	SetOperation(v string) *RestartPolarClawGatewayResponseBody
	GetOperation() *string
	SetRequestId(v string) *RestartPolarClawGatewayResponseBody
	GetRequestId() *string
	SetRestarted(v bool) *RestartPolarClawGatewayResponseBody
	GetRestarted() *bool
	SetState(v string) *RestartPolarClawGatewayResponseBody
	GetState() *string
	SetTaskId(v string) *RestartPolarClawGatewayResponseBody
	GetTaskId() *string
}

type RestartPolarClawGatewayResponseBody struct {
	// example:
	//
	// pa-xxx
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 3235
	DowntimeMs *int64 `json:"DowntimeMs,omitempty" xml:"DowntimeMs,omitempty"`
	// example:
	//
	// 2026.5.7
	GatewayVersion *string `json:"GatewayVersion,omitempty" xml:"GatewayVersion,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// in-process
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// example:
	//
	// true
	Ok *bool `json:"Ok,omitempty" xml:"Ok,omitempty"`
	// example:
	//
	// RestartPolarClawGateway
	Operation *string `json:"Operation,omitempty" xml:"Operation,omitempty"`
	// example:
	//
	// F45FFACC-xxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Restarted *bool `json:"Restarted,omitempty" xml:"Restarted,omitempty"`
	// example:
	//
	// pending
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// example:
	//
	// 0ee00f56-f467-4d41-858c-ca4ede2c770e
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s RestartPolarClawGatewayResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RestartPolarClawGatewayResponseBody) GoString() string {
	return s.String()
}

func (s *RestartPolarClawGatewayResponseBody) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *RestartPolarClawGatewayResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *RestartPolarClawGatewayResponseBody) GetDowntimeMs() *int64 {
	return s.DowntimeMs
}

func (s *RestartPolarClawGatewayResponseBody) GetGatewayVersion() *string {
	return s.GatewayVersion
}

func (s *RestartPolarClawGatewayResponseBody) GetMessage() *string {
	return s.Message
}

func (s *RestartPolarClawGatewayResponseBody) GetMode() *string {
	return s.Mode
}

func (s *RestartPolarClawGatewayResponseBody) GetOk() *bool {
	return s.Ok
}

func (s *RestartPolarClawGatewayResponseBody) GetOperation() *string {
	return s.Operation
}

func (s *RestartPolarClawGatewayResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RestartPolarClawGatewayResponseBody) GetRestarted() *bool {
	return s.Restarted
}

func (s *RestartPolarClawGatewayResponseBody) GetState() *string {
	return s.State
}

func (s *RestartPolarClawGatewayResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *RestartPolarClawGatewayResponseBody) SetApplicationId(v string) *RestartPolarClawGatewayResponseBody {
	s.ApplicationId = &v
	return s
}

func (s *RestartPolarClawGatewayResponseBody) SetCode(v int32) *RestartPolarClawGatewayResponseBody {
	s.Code = &v
	return s
}

func (s *RestartPolarClawGatewayResponseBody) SetDowntimeMs(v int64) *RestartPolarClawGatewayResponseBody {
	s.DowntimeMs = &v
	return s
}

func (s *RestartPolarClawGatewayResponseBody) SetGatewayVersion(v string) *RestartPolarClawGatewayResponseBody {
	s.GatewayVersion = &v
	return s
}

func (s *RestartPolarClawGatewayResponseBody) SetMessage(v string) *RestartPolarClawGatewayResponseBody {
	s.Message = &v
	return s
}

func (s *RestartPolarClawGatewayResponseBody) SetMode(v string) *RestartPolarClawGatewayResponseBody {
	s.Mode = &v
	return s
}

func (s *RestartPolarClawGatewayResponseBody) SetOk(v bool) *RestartPolarClawGatewayResponseBody {
	s.Ok = &v
	return s
}

func (s *RestartPolarClawGatewayResponseBody) SetOperation(v string) *RestartPolarClawGatewayResponseBody {
	s.Operation = &v
	return s
}

func (s *RestartPolarClawGatewayResponseBody) SetRequestId(v string) *RestartPolarClawGatewayResponseBody {
	s.RequestId = &v
	return s
}

func (s *RestartPolarClawGatewayResponseBody) SetRestarted(v bool) *RestartPolarClawGatewayResponseBody {
	s.Restarted = &v
	return s
}

func (s *RestartPolarClawGatewayResponseBody) SetState(v string) *RestartPolarClawGatewayResponseBody {
	s.State = &v
	return s
}

func (s *RestartPolarClawGatewayResponseBody) SetTaskId(v string) *RestartPolarClawGatewayResponseBody {
	s.TaskId = &v
	return s
}

func (s *RestartPolarClawGatewayResponseBody) Validate() error {
	return dara.Validate(s)
}
