// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBasicIpSetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetIpSetId(v string) *CreateBasicIpSetResponseBody
	GetIpSetId() *string
	SetRequestId(v string) *CreateBasicIpSetResponseBody
	GetRequestId() *string
}

type CreateBasicIpSetResponseBody struct {
	// The region ID of the GA instance.
	//
	// example:
	//
	// ips-bp11r5jb8ogp122xl****
	IpSetId *string `json:"IpSetId,omitempty" xml:"IpSetId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateBasicIpSetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateBasicIpSetResponseBody) GoString() string {
	return s.String()
}

func (s *CreateBasicIpSetResponseBody) GetIpSetId() *string {
	return s.IpSetId
}

func (s *CreateBasicIpSetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateBasicIpSetResponseBody) SetIpSetId(v string) *CreateBasicIpSetResponseBody {
	s.IpSetId = &v
	return s
}

func (s *CreateBasicIpSetResponseBody) SetRequestId(v string) *CreateBasicIpSetResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateBasicIpSetResponseBody) Validate() error {
	return dara.Validate(s)
}
