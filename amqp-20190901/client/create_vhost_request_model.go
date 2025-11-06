// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVhostRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConsoleSessionId(v string) *CreateVhostRequest
	GetConsoleSessionId() *string
	SetInstanceId(v string) *CreateVhostRequest
	GetInstanceId() *string
	SetVhostName(v string) *CreateVhostRequest
	GetVhostName() *string
}

type CreateVhostRequest struct {
	ConsoleSessionId *string `json:"ConsoleSessionId,omitempty" xml:"ConsoleSessionId,omitempty"`
	InstanceId       *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	VhostName *string `json:"VhostName,omitempty" xml:"VhostName,omitempty"`
}

func (s CreateVhostRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateVhostRequest) GoString() string {
	return s.String()
}

func (s *CreateVhostRequest) GetConsoleSessionId() *string {
	return s.ConsoleSessionId
}

func (s *CreateVhostRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateVhostRequest) GetVhostName() *string {
	return s.VhostName
}

func (s *CreateVhostRequest) SetConsoleSessionId(v string) *CreateVhostRequest {
	s.ConsoleSessionId = &v
	return s
}

func (s *CreateVhostRequest) SetInstanceId(v string) *CreateVhostRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateVhostRequest) SetVhostName(v string) *CreateVhostRequest {
	s.VhostName = &v
	return s
}

func (s *CreateVhostRequest) Validate() error {
	return dara.Validate(s)
}
