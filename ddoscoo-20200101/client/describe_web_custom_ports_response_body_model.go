// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeWebCustomPortsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeWebCustomPortsResponseBody
	GetRequestId() *string
	SetWebCustomPorts(v []*DescribeWebCustomPortsResponseBodyWebCustomPorts) *DescribeWebCustomPortsResponseBody
	GetWebCustomPorts() []*DescribeWebCustomPortsResponseBodyWebCustomPorts
}

type DescribeWebCustomPortsResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 0bcf28g5-d57c-11e7-9bs0-d89d6717dxbc
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// An array consisting of information about supported custom ports that are used by a website.
	WebCustomPorts []*DescribeWebCustomPortsResponseBodyWebCustomPorts `json:"WebCustomPorts,omitempty" xml:"WebCustomPorts,omitempty" type:"Repeated"`
}

func (s DescribeWebCustomPortsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeWebCustomPortsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeWebCustomPortsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeWebCustomPortsResponseBody) GetWebCustomPorts() []*DescribeWebCustomPortsResponseBodyWebCustomPorts {
	return s.WebCustomPorts
}

func (s *DescribeWebCustomPortsResponseBody) SetRequestId(v string) *DescribeWebCustomPortsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeWebCustomPortsResponseBody) SetWebCustomPorts(v []*DescribeWebCustomPortsResponseBodyWebCustomPorts) *DescribeWebCustomPortsResponseBody {
	s.WebCustomPorts = v
	return s
}

func (s *DescribeWebCustomPortsResponseBody) Validate() error {
	if s.WebCustomPorts != nil {
		for _, item := range s.WebCustomPorts {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeWebCustomPortsResponseBodyWebCustomPorts struct {
	// An array that consists of supported custom ports.
	ProxyPorts []*string `json:"ProxyPorts,omitempty" xml:"ProxyPorts,omitempty" type:"Repeated"`
	// The type of the protocol. Valid values:
	//
	// 	- **http**
	//
	// 	- **https**
	//
	// example:
	//
	// http
	ProxyType *string `json:"ProxyType,omitempty" xml:"ProxyType,omitempty"`
}

func (s DescribeWebCustomPortsResponseBodyWebCustomPorts) String() string {
	return dara.Prettify(s)
}

func (s DescribeWebCustomPortsResponseBodyWebCustomPorts) GoString() string {
	return s.String()
}

func (s *DescribeWebCustomPortsResponseBodyWebCustomPorts) GetProxyPorts() []*string {
	return s.ProxyPorts
}

func (s *DescribeWebCustomPortsResponseBodyWebCustomPorts) GetProxyType() *string {
	return s.ProxyType
}

func (s *DescribeWebCustomPortsResponseBodyWebCustomPorts) SetProxyPorts(v []*string) *DescribeWebCustomPortsResponseBodyWebCustomPorts {
	s.ProxyPorts = v
	return s
}

func (s *DescribeWebCustomPortsResponseBodyWebCustomPorts) SetProxyType(v string) *DescribeWebCustomPortsResponseBodyWebCustomPorts {
	s.ProxyType = &v
	return s
}

func (s *DescribeWebCustomPortsResponseBodyWebCustomPorts) Validate() error {
	return dara.Validate(s)
}
