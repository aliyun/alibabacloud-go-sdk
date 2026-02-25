// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPortProtocol interface {
	dara.Model
	String() string
	GoString() string
	SetPort(v int32) *PortProtocol
	GetPort() *int32
	SetProtocol(v string) *PortProtocol
	GetProtocol() *string
	SetTargetPort(v int32) *PortProtocol
	GetTargetPort() *int32
}

type PortProtocol struct {
	Port       *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
	Protocol   *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	TargetPort *int32  `json:"TargetPort,omitempty" xml:"TargetPort,omitempty"`
}

func (s PortProtocol) String() string {
	return dara.Prettify(s)
}

func (s PortProtocol) GoString() string {
	return s.String()
}

func (s *PortProtocol) GetPort() *int32 {
	return s.Port
}

func (s *PortProtocol) GetProtocol() *string {
	return s.Protocol
}

func (s *PortProtocol) GetTargetPort() *int32 {
	return s.TargetPort
}

func (s *PortProtocol) SetPort(v int32) *PortProtocol {
	s.Port = &v
	return s
}

func (s *PortProtocol) SetProtocol(v string) *PortProtocol {
	s.Protocol = &v
	return s
}

func (s *PortProtocol) SetTargetPort(v int32) *PortProtocol {
	s.TargetPort = &v
	return s
}

func (s *PortProtocol) Validate() error {
	return dara.Validate(s)
}
