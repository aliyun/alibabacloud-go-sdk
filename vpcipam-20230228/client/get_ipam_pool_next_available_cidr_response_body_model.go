// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetIpamPoolNextAvailableCidrResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCidrBlock(v string) *GetIpamPoolNextAvailableCidrResponseBody
	GetCidrBlock() *string
	SetRequestId(v string) *GetIpamPoolNextAvailableCidrResponseBody
	GetRequestId() *string
}

type GetIpamPoolNextAvailableCidrResponseBody struct {
	// Available CIDR.
	//
	// example:
	//
	// 172.68.0.0/26
	CidrBlock *string `json:"CidrBlock,omitempty" xml:"CidrBlock,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 29FC6758-9B7C-5CC7-8CBF-4DD846FE7D82
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetIpamPoolNextAvailableCidrResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetIpamPoolNextAvailableCidrResponseBody) GoString() string {
	return s.String()
}

func (s *GetIpamPoolNextAvailableCidrResponseBody) GetCidrBlock() *string {
	return s.CidrBlock
}

func (s *GetIpamPoolNextAvailableCidrResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetIpamPoolNextAvailableCidrResponseBody) SetCidrBlock(v string) *GetIpamPoolNextAvailableCidrResponseBody {
	s.CidrBlock = &v
	return s
}

func (s *GetIpamPoolNextAvailableCidrResponseBody) SetRequestId(v string) *GetIpamPoolNextAvailableCidrResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetIpamPoolNextAvailableCidrResponseBody) Validate() error {
	return dara.Validate(s)
}
