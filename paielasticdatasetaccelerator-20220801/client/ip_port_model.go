// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIpPort interface {
	dara.Model
	String() string
	GoString() string
	SetIp(v string) *IpPort
	GetIp() *string
	SetPort(v string) *IpPort
	GetPort() *string
}

type IpPort struct {
	// example:
	//
	// 10.0.0.5
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// example:
	//
	// 3306
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
}

func (s IpPort) String() string {
	return dara.Prettify(s)
}

func (s IpPort) GoString() string {
	return s.String()
}

func (s *IpPort) GetIp() *string {
	return s.Ip
}

func (s *IpPort) GetPort() *string {
	return s.Port
}

func (s *IpPort) SetIp(v string) *IpPort {
	s.Ip = &v
	return s
}

func (s *IpPort) SetPort(v string) *IpPort {
	s.Port = &v
	return s
}

func (s *IpPort) Validate() error {
	return dara.Validate(s)
}
