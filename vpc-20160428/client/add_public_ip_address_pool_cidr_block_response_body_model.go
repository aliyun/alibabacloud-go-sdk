// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddPublicIpAddressPoolCidrBlockResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCidrBlock(v string) *AddPublicIpAddressPoolCidrBlockResponseBody
	GetCidrBlock() *string
	SetRequestId(v string) *AddPublicIpAddressPoolCidrBlockResponseBody
	GetRequestId() *string
}

type AddPublicIpAddressPoolCidrBlockResponseBody struct {
	// The CIDR block.
	//
	// example:
	//
	// 47.0.XX.XX/28
	CidrBlock *string `json:"CidrBlock,omitempty" xml:"CidrBlock,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 4EC47282-1B74-4534-BD0E-403F3EE64CAF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddPublicIpAddressPoolCidrBlockResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddPublicIpAddressPoolCidrBlockResponseBody) GoString() string {
	return s.String()
}

func (s *AddPublicIpAddressPoolCidrBlockResponseBody) GetCidrBlock() *string {
	return s.CidrBlock
}

func (s *AddPublicIpAddressPoolCidrBlockResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddPublicIpAddressPoolCidrBlockResponseBody) SetCidrBlock(v string) *AddPublicIpAddressPoolCidrBlockResponseBody {
	s.CidrBlock = &v
	return s
}

func (s *AddPublicIpAddressPoolCidrBlockResponseBody) SetRequestId(v string) *AddPublicIpAddressPoolCidrBlockResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddPublicIpAddressPoolCidrBlockResponseBody) Validate() error {
	return dara.Validate(s)
}
