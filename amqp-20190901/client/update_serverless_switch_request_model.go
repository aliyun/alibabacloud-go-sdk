// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateServerlessSwitchRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConsoleSessionId(v string) *UpdateServerlessSwitchRequest
	GetConsoleSessionId() *string
	SetInstanceId(v string) *UpdateServerlessSwitchRequest
	GetInstanceId() *string
	SetServerlessSwitch(v bool) *UpdateServerlessSwitchRequest
	GetServerlessSwitch() *bool
}

type UpdateServerlessSwitchRequest struct {
	ConsoleSessionId *string `json:"ConsoleSessionId,omitempty" xml:"ConsoleSessionId,omitempty"`
	// This parameter is required.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	ServerlessSwitch *bool `json:"ServerlessSwitch,omitempty" xml:"ServerlessSwitch,omitempty"`
}

func (s UpdateServerlessSwitchRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateServerlessSwitchRequest) GoString() string {
	return s.String()
}

func (s *UpdateServerlessSwitchRequest) GetConsoleSessionId() *string {
	return s.ConsoleSessionId
}

func (s *UpdateServerlessSwitchRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateServerlessSwitchRequest) GetServerlessSwitch() *bool {
	return s.ServerlessSwitch
}

func (s *UpdateServerlessSwitchRequest) SetConsoleSessionId(v string) *UpdateServerlessSwitchRequest {
	s.ConsoleSessionId = &v
	return s
}

func (s *UpdateServerlessSwitchRequest) SetInstanceId(v string) *UpdateServerlessSwitchRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateServerlessSwitchRequest) SetServerlessSwitch(v bool) *UpdateServerlessSwitchRequest {
	s.ServerlessSwitch = &v
	return s
}

func (s *UpdateServerlessSwitchRequest) Validate() error {
	return dara.Validate(s)
}
