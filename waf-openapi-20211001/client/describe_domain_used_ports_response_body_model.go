// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainUsedPortsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeDomainUsedPortsResponseBody
	GetRequestId() *string
	SetUsedPorts(v []*int32) *DescribeDomainUsedPortsResponseBody
	GetUsedPorts() []*int32
}

type DescribeDomainUsedPortsResponseBody struct {
	// example:
	//
	// D7861F61-5B61-*-A47C-*
	RequestId *string  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	UsedPorts []*int32 `json:"UsedPorts,omitempty" xml:"UsedPorts,omitempty" type:"Repeated"`
}

func (s DescribeDomainUsedPortsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainUsedPortsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainUsedPortsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDomainUsedPortsResponseBody) GetUsedPorts() []*int32 {
	return s.UsedPorts
}

func (s *DescribeDomainUsedPortsResponseBody) SetRequestId(v string) *DescribeDomainUsedPortsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainUsedPortsResponseBody) SetUsedPorts(v []*int32) *DescribeDomainUsedPortsResponseBody {
	s.UsedPorts = v
	return s
}

func (s *DescribeDomainUsedPortsResponseBody) Validate() error {
	return dara.Validate(s)
}
