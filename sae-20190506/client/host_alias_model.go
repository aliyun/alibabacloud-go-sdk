// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHostAlias interface {
	dara.Model
	String() string
	GoString() string
	SetHostnames(v []*string) *HostAlias
	GetHostnames() []*string
	SetIp(v string) *HostAlias
	GetIp() *string
}

type HostAlias struct {
	Hostnames []*string `json:"hostnames,omitempty" xml:"hostnames,omitempty" type:"Repeated"`
	Ip        *string   `json:"ip,omitempty" xml:"ip,omitempty"`
}

func (s HostAlias) String() string {
	return dara.Prettify(s)
}

func (s HostAlias) GoString() string {
	return s.String()
}

func (s *HostAlias) GetHostnames() []*string {
	return s.Hostnames
}

func (s *HostAlias) GetIp() *string {
	return s.Ip
}

func (s *HostAlias) SetHostnames(v []*string) *HostAlias {
	s.Hostnames = v
	return s
}

func (s *HostAlias) SetIp(v string) *HostAlias {
	s.Ip = &v
	return s
}

func (s *HostAlias) Validate() error {
	return dara.Validate(s)
}
