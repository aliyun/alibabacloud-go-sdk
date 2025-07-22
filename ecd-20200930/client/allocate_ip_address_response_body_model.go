// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAllocateIpAddressResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEipAddress(v string) *AllocateIpAddressResponseBody
	GetEipAddress() *string
	SetEipId(v string) *AllocateIpAddressResponseBody
	GetEipId() *string
	SetRequestId(v string) *AllocateIpAddressResponseBody
	GetRequestId() *string
}

type AllocateIpAddressResponseBody struct {
	EipAddress *string `json:"EipAddress,omitempty" xml:"EipAddress,omitempty"`
	EipId      *string `json:"EipId,omitempty" xml:"EipId,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AllocateIpAddressResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AllocateIpAddressResponseBody) GoString() string {
	return s.String()
}

func (s *AllocateIpAddressResponseBody) GetEipAddress() *string {
	return s.EipAddress
}

func (s *AllocateIpAddressResponseBody) GetEipId() *string {
	return s.EipId
}

func (s *AllocateIpAddressResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AllocateIpAddressResponseBody) SetEipAddress(v string) *AllocateIpAddressResponseBody {
	s.EipAddress = &v
	return s
}

func (s *AllocateIpAddressResponseBody) SetEipId(v string) *AllocateIpAddressResponseBody {
	s.EipId = &v
	return s
}

func (s *AllocateIpAddressResponseBody) SetRequestId(v string) *AllocateIpAddressResponseBody {
	s.RequestId = &v
	return s
}

func (s *AllocateIpAddressResponseBody) Validate() error {
	return dara.Validate(s)
}
