// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLayer7CustomPortsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLayer7CustomPorts(v []*ListLayer7CustomPortsResponseBodyLayer7CustomPorts) *ListLayer7CustomPortsResponseBody
	GetLayer7CustomPorts() []*ListLayer7CustomPortsResponseBodyLayer7CustomPorts
	SetRequestId(v string) *ListLayer7CustomPortsResponseBody
	GetRequestId() *string
}

type ListLayer7CustomPortsResponseBody struct {
	Layer7CustomPorts []*ListLayer7CustomPortsResponseBodyLayer7CustomPorts `json:"Layer7CustomPorts,omitempty" xml:"Layer7CustomPorts,omitempty" type:"Repeated"`
	// example:
	//
	// CF33B4C3-196E-4015-AADD-5CAD00057B80
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListLayer7CustomPortsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListLayer7CustomPortsResponseBody) GoString() string {
	return s.String()
}

func (s *ListLayer7CustomPortsResponseBody) GetLayer7CustomPorts() []*ListLayer7CustomPortsResponseBodyLayer7CustomPorts {
	return s.Layer7CustomPorts
}

func (s *ListLayer7CustomPortsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListLayer7CustomPortsResponseBody) SetLayer7CustomPorts(v []*ListLayer7CustomPortsResponseBodyLayer7CustomPorts) *ListLayer7CustomPortsResponseBody {
	s.Layer7CustomPorts = v
	return s
}

func (s *ListLayer7CustomPortsResponseBody) SetRequestId(v string) *ListLayer7CustomPortsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListLayer7CustomPortsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListLayer7CustomPortsResponseBodyLayer7CustomPorts struct {
	Flag       *string   `json:"Flag,omitempty" xml:"Flag,omitempty"`
	ProxyPorts []*string `json:"ProxyPorts,omitempty" xml:"ProxyPorts,omitempty" type:"Repeated"`
	// example:
	//
	// https
	ProxyType *string `json:"ProxyType,omitempty" xml:"ProxyType,omitempty"`
}

func (s ListLayer7CustomPortsResponseBodyLayer7CustomPorts) String() string {
	return dara.Prettify(s)
}

func (s ListLayer7CustomPortsResponseBodyLayer7CustomPorts) GoString() string {
	return s.String()
}

func (s *ListLayer7CustomPortsResponseBodyLayer7CustomPorts) GetFlag() *string {
	return s.Flag
}

func (s *ListLayer7CustomPortsResponseBodyLayer7CustomPorts) GetProxyPorts() []*string {
	return s.ProxyPorts
}

func (s *ListLayer7CustomPortsResponseBodyLayer7CustomPorts) GetProxyType() *string {
	return s.ProxyType
}

func (s *ListLayer7CustomPortsResponseBodyLayer7CustomPorts) SetFlag(v string) *ListLayer7CustomPortsResponseBodyLayer7CustomPorts {
	s.Flag = &v
	return s
}

func (s *ListLayer7CustomPortsResponseBodyLayer7CustomPorts) SetProxyPorts(v []*string) *ListLayer7CustomPortsResponseBodyLayer7CustomPorts {
	s.ProxyPorts = v
	return s
}

func (s *ListLayer7CustomPortsResponseBodyLayer7CustomPorts) SetProxyType(v string) *ListLayer7CustomPortsResponseBodyLayer7CustomPorts {
	s.ProxyType = &v
	return s
}

func (s *ListLayer7CustomPortsResponseBodyLayer7CustomPorts) Validate() error {
	return dara.Validate(s)
}
