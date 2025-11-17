// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPodNetworkInterface interface {
	dara.Model
	String() string
	GoString() string
	SetInterfaceName(v string) *PodNetworkInterface
	GetInterfaceName() *string
	SetIp(v string) *PodNetworkInterface
	GetIp() *string
}

type PodNetworkInterface struct {
	InterfaceName *string `json:"InterfaceName,omitempty" xml:"InterfaceName,omitempty"`
	Ip            *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
}

func (s PodNetworkInterface) String() string {
	return dara.Prettify(s)
}

func (s PodNetworkInterface) GoString() string {
	return s.String()
}

func (s *PodNetworkInterface) GetInterfaceName() *string {
	return s.InterfaceName
}

func (s *PodNetworkInterface) GetIp() *string {
	return s.Ip
}

func (s *PodNetworkInterface) SetInterfaceName(v string) *PodNetworkInterface {
	s.InterfaceName = &v
	return s
}

func (s *PodNetworkInterface) SetIp(v string) *PodNetworkInterface {
	s.Ip = &v
	return s
}

func (s *PodNetworkInterface) Validate() error {
	return dara.Validate(s)
}
