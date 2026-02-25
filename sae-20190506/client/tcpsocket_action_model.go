// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTCPSocketAction interface {
	dara.Model
	String() string
	GoString() string
	SetHost(v string) *TCPSocketAction
	GetHost() *string
	SetPort(v int32) *TCPSocketAction
	GetPort() *int32
}

type TCPSocketAction struct {
	// Optional. The hostname to which you want to connect. The default value is the IP address of the pod.
	//
	// example:
	//
	// 172.16.XX.XX
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// The port number that you want to access on the container. Valid values: 1 to 65535.
	//
	// example:
	//
	// 80
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
}

func (s TCPSocketAction) String() string {
	return dara.Prettify(s)
}

func (s TCPSocketAction) GoString() string {
	return s.String()
}

func (s *TCPSocketAction) GetHost() *string {
	return s.Host
}

func (s *TCPSocketAction) GetPort() *int32 {
	return s.Port
}

func (s *TCPSocketAction) SetHost(v string) *TCPSocketAction {
	s.Host = &v
	return s
}

func (s *TCPSocketAction) SetPort(v int32) *TCPSocketAction {
	s.Port = &v
	return s
}

func (s *TCPSocketAction) Validate() error {
	return dara.Validate(s)
}
