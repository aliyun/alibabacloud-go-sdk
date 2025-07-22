// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iKillInstanceAllSessionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConsoleContext(v string) *KillInstanceAllSessionRequest
	GetConsoleContext() *string
	SetInstanceId(v string) *KillInstanceAllSessionRequest
	GetInstanceId() *string
}

type KillInstanceAllSessionRequest struct {
	// The reserved parameter.
	//
	// example:
	//
	// None
	ConsoleContext *string `json:"ConsoleContext,omitempty" xml:"ConsoleContext,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// r-8vbcyr4sw0c4yc****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s KillInstanceAllSessionRequest) String() string {
	return dara.Prettify(s)
}

func (s KillInstanceAllSessionRequest) GoString() string {
	return s.String()
}

func (s *KillInstanceAllSessionRequest) GetConsoleContext() *string {
	return s.ConsoleContext
}

func (s *KillInstanceAllSessionRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *KillInstanceAllSessionRequest) SetConsoleContext(v string) *KillInstanceAllSessionRequest {
	s.ConsoleContext = &v
	return s
}

func (s *KillInstanceAllSessionRequest) SetInstanceId(v string) *KillInstanceAllSessionRequest {
	s.InstanceId = &v
	return s
}

func (s *KillInstanceAllSessionRequest) Validate() error {
	return dara.Validate(s)
}
