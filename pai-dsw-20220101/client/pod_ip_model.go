// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPodIp interface {
	dara.Model
	String() string
	GoString() string
	SetInterfaceName(v string) *PodIp
	GetInterfaceName() *string
	SetIp(v string) *PodIp
	GetIp() *string
	SetType(v string) *PodIp
	GetType() *string
}

type PodIp struct {
	InterfaceName *string `json:"InterfaceName,omitempty" xml:"InterfaceName,omitempty"`
	Ip            *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	Type          *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s PodIp) String() string {
	return dara.Prettify(s)
}

func (s PodIp) GoString() string {
	return s.String()
}

func (s *PodIp) GetInterfaceName() *string {
	return s.InterfaceName
}

func (s *PodIp) GetIp() *string {
	return s.Ip
}

func (s *PodIp) GetType() *string {
	return s.Type
}

func (s *PodIp) SetInterfaceName(v string) *PodIp {
	s.InterfaceName = &v
	return s
}

func (s *PodIp) SetIp(v string) *PodIp {
	s.Ip = &v
	return s
}

func (s *PodIp) SetType(v string) *PodIp {
	s.Type = &v
	return s
}

func (s *PodIp) Validate() error {
	return dara.Validate(s)
}
