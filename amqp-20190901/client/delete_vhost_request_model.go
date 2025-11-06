// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVhostRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConsoleSessionId(v string) *DeleteVhostRequest
	GetConsoleSessionId() *string
	SetInstanceId(v string) *DeleteVhostRequest
	GetInstanceId() *string
	SetVhostName(v string) *DeleteVhostRequest
	GetVhostName() *string
	SetVhostNames(v map[string]interface{}) *DeleteVhostRequest
	GetVhostNames() map[string]interface{}
}

type DeleteVhostRequest struct {
	ConsoleSessionId *string                `json:"ConsoleSessionId,omitempty" xml:"ConsoleSessionId,omitempty"`
	InstanceId       *string                `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	VhostName        *string                `json:"VhostName,omitempty" xml:"VhostName,omitempty"`
	VhostNames       map[string]interface{} `json:"VhostNames,omitempty" xml:"VhostNames,omitempty"`
}

func (s DeleteVhostRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteVhostRequest) GoString() string {
	return s.String()
}

func (s *DeleteVhostRequest) GetConsoleSessionId() *string {
	return s.ConsoleSessionId
}

func (s *DeleteVhostRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteVhostRequest) GetVhostName() *string {
	return s.VhostName
}

func (s *DeleteVhostRequest) GetVhostNames() map[string]interface{} {
	return s.VhostNames
}

func (s *DeleteVhostRequest) SetConsoleSessionId(v string) *DeleteVhostRequest {
	s.ConsoleSessionId = &v
	return s
}

func (s *DeleteVhostRequest) SetInstanceId(v string) *DeleteVhostRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteVhostRequest) SetVhostName(v string) *DeleteVhostRequest {
	s.VhostName = &v
	return s
}

func (s *DeleteVhostRequest) SetVhostNames(v map[string]interface{}) *DeleteVhostRequest {
	s.VhostNames = v
	return s
}

func (s *DeleteVhostRequest) Validate() error {
	return dara.Validate(s)
}
