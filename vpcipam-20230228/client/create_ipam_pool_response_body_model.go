// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateIpamPoolResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetIpamPoolId(v string) *CreateIpamPoolResponseBody
	GetIpamPoolId() *string
	SetRequestId(v string) *CreateIpamPoolResponseBody
	GetRequestId() *string
}

type CreateIpamPoolResponseBody struct {
	// The ID of the IPAM pool.
	//
	// example:
	//
	// ipam-pool-6rcq3tobayc20t****
	IpamPoolId *string `json:"IpamPoolId,omitempty" xml:"IpamPoolId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// BB2C39DE-CEB8-595A-981A-F2EFCBE7324E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateIpamPoolResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateIpamPoolResponseBody) GoString() string {
	return s.String()
}

func (s *CreateIpamPoolResponseBody) GetIpamPoolId() *string {
	return s.IpamPoolId
}

func (s *CreateIpamPoolResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateIpamPoolResponseBody) SetIpamPoolId(v string) *CreateIpamPoolResponseBody {
	s.IpamPoolId = &v
	return s
}

func (s *CreateIpamPoolResponseBody) SetRequestId(v string) *CreateIpamPoolResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateIpamPoolResponseBody) Validate() error {
	return dara.Validate(s)
}
