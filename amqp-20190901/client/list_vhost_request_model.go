// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVhostRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConsoleSessionId(v string) *ListVhostRequest
	GetConsoleSessionId() *string
	SetInstanceId(v string) *ListVhostRequest
	GetInstanceId() *string
	SetVhostNamePrefix(v string) *ListVhostRequest
	GetVhostNamePrefix() *string
}

type ListVhostRequest struct {
	ConsoleSessionId *string `json:"ConsoleSessionId,omitempty" xml:"ConsoleSessionId,omitempty"`
	InstanceId       *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	VhostNamePrefix  *string `json:"VhostNamePrefix,omitempty" xml:"VhostNamePrefix,omitempty"`
}

func (s ListVhostRequest) String() string {
	return dara.Prettify(s)
}

func (s ListVhostRequest) GoString() string {
	return s.String()
}

func (s *ListVhostRequest) GetConsoleSessionId() *string {
	return s.ConsoleSessionId
}

func (s *ListVhostRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListVhostRequest) GetVhostNamePrefix() *string {
	return s.VhostNamePrefix
}

func (s *ListVhostRequest) SetConsoleSessionId(v string) *ListVhostRequest {
	s.ConsoleSessionId = &v
	return s
}

func (s *ListVhostRequest) SetInstanceId(v string) *ListVhostRequest {
	s.InstanceId = &v
	return s
}

func (s *ListVhostRequest) SetVhostNamePrefix(v string) *ListVhostRequest {
	s.VhostNamePrefix = &v
	return s
}

func (s *ListVhostRequest) Validate() error {
	return dara.Validate(s)
}
