// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDedicatedIpNonePoolListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetIps(v []*DedicatedIpNonePoolListResponseBodyIps) *DedicatedIpNonePoolListResponseBody
	GetIps() []*DedicatedIpNonePoolListResponseBodyIps
	SetRequestId(v string) *DedicatedIpNonePoolListResponseBody
	GetRequestId() *string
}

type DedicatedIpNonePoolListResponseBody struct {
	// The IP addresses that are not added to an IP pool.
	Ips []*DedicatedIpNonePoolListResponseBodyIps `json:"Ips,omitempty" xml:"Ips,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// xxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DedicatedIpNonePoolListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DedicatedIpNonePoolListResponseBody) GoString() string {
	return s.String()
}

func (s *DedicatedIpNonePoolListResponseBody) GetIps() []*DedicatedIpNonePoolListResponseBodyIps {
	return s.Ips
}

func (s *DedicatedIpNonePoolListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DedicatedIpNonePoolListResponseBody) SetIps(v []*DedicatedIpNonePoolListResponseBodyIps) *DedicatedIpNonePoolListResponseBody {
	s.Ips = v
	return s
}

func (s *DedicatedIpNonePoolListResponseBody) SetRequestId(v string) *DedicatedIpNonePoolListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DedicatedIpNonePoolListResponseBody) Validate() error {
	if s.Ips != nil {
		for _, item := range s.Ips {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DedicatedIpNonePoolListResponseBodyIps struct {
	// The ID of the purchased instance.
	//
	// example:
	//
	// xxx
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The IP address.
	//
	// example:
	//
	// xxx
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The zone ID.
	//
	// example:
	//
	// xxx
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DedicatedIpNonePoolListResponseBodyIps) String() string {
	return dara.Prettify(s)
}

func (s DedicatedIpNonePoolListResponseBodyIps) GoString() string {
	return s.String()
}

func (s *DedicatedIpNonePoolListResponseBodyIps) GetId() *string {
	return s.Id
}

func (s *DedicatedIpNonePoolListResponseBodyIps) GetIp() *string {
	return s.Ip
}

func (s *DedicatedIpNonePoolListResponseBodyIps) GetZoneId() *string {
	return s.ZoneId
}

func (s *DedicatedIpNonePoolListResponseBodyIps) SetId(v string) *DedicatedIpNonePoolListResponseBodyIps {
	s.Id = &v
	return s
}

func (s *DedicatedIpNonePoolListResponseBodyIps) SetIp(v string) *DedicatedIpNonePoolListResponseBodyIps {
	s.Ip = &v
	return s
}

func (s *DedicatedIpNonePoolListResponseBodyIps) SetZoneId(v string) *DedicatedIpNonePoolListResponseBodyIps {
	s.ZoneId = &v
	return s
}

func (s *DedicatedIpNonePoolListResponseBodyIps) Validate() error {
	return dara.Validate(s)
}
