// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddIpamPoolCidrResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCidr(v string) *AddIpamPoolCidrResponseBody
	GetCidr() *string
	SetRequestId(v string) *AddIpamPoolCidrResponseBody
	GetRequestId() *string
}

type AddIpamPoolCidrResponseBody struct {
	// The successfully provisioned CIDR block.
	//
	// example:
	//
	// 192.168.1.0/24
	Cidr *string `json:"Cidr,omitempty" xml:"Cidr,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 558BC336-8B88-53B0-B4AD-980EE900AB01
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddIpamPoolCidrResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddIpamPoolCidrResponseBody) GoString() string {
	return s.String()
}

func (s *AddIpamPoolCidrResponseBody) GetCidr() *string {
	return s.Cidr
}

func (s *AddIpamPoolCidrResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddIpamPoolCidrResponseBody) SetCidr(v string) *AddIpamPoolCidrResponseBody {
	s.Cidr = &v
	return s
}

func (s *AddIpamPoolCidrResponseBody) SetRequestId(v string) *AddIpamPoolCidrResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddIpamPoolCidrResponseBody) Validate() error {
	return dara.Validate(s)
}
