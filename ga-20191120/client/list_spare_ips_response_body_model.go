// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSpareIpsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListSpareIpsResponseBody
	GetRequestId() *string
	SetSpareIps(v []*ListSpareIpsResponseBodySpareIps) *ListSpareIpsResponseBody
	GetSpareIps() []*ListSpareIpsResponseBodySpareIps
}

type ListSpareIpsResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 6FEA0CF3-D3B9-43E5-A304-D217037876A8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The secondary IP addresses that are associated with the CNAME.
	SpareIps []*ListSpareIpsResponseBodySpareIps `json:"SpareIps,omitempty" xml:"SpareIps,omitempty" type:"Repeated"`
}

func (s ListSpareIpsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSpareIpsResponseBody) GoString() string {
	return s.String()
}

func (s *ListSpareIpsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSpareIpsResponseBody) GetSpareIps() []*ListSpareIpsResponseBodySpareIps {
	return s.SpareIps
}

func (s *ListSpareIpsResponseBody) SetRequestId(v string) *ListSpareIpsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSpareIpsResponseBody) SetSpareIps(v []*ListSpareIpsResponseBodySpareIps) *ListSpareIpsResponseBody {
	s.SpareIps = v
	return s
}

func (s *ListSpareIpsResponseBody) Validate() error {
	if s.SpareIps != nil {
		for _, item := range s.SpareIps {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListSpareIpsResponseBodySpareIps struct {
	// The secondary IP address that is associated with the CNAME. If the acceleration area becomes unavailable, GA redirects traffic to the secondary IP address.
	//
	// example:
	//
	// 47.100.XX.XX
	SpareIp *string `json:"SpareIp,omitempty" xml:"SpareIp,omitempty"`
	// The status of the secondary IP address. Valid values:
	//
	// 	- **active:*	- The secondary IP address is available.
	//
	// 	- **inuse:*	- The secondary IP address is in use.
	//
	// example:
	//
	// active
	State *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s ListSpareIpsResponseBodySpareIps) String() string {
	return dara.Prettify(s)
}

func (s ListSpareIpsResponseBodySpareIps) GoString() string {
	return s.String()
}

func (s *ListSpareIpsResponseBodySpareIps) GetSpareIp() *string {
	return s.SpareIp
}

func (s *ListSpareIpsResponseBodySpareIps) GetState() *string {
	return s.State
}

func (s *ListSpareIpsResponseBodySpareIps) SetSpareIp(v string) *ListSpareIpsResponseBodySpareIps {
	s.SpareIp = &v
	return s
}

func (s *ListSpareIpsResponseBodySpareIps) SetState(v string) *ListSpareIpsResponseBodySpareIps {
	s.State = &v
	return s
}

func (s *ListSpareIpsResponseBodySpareIps) Validate() error {
	return dara.Validate(s)
}
