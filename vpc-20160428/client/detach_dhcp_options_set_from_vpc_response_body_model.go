// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachDhcpOptionsSetFromVpcResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DetachDhcpOptionsSetFromVpcResponseBody
	GetRequestId() *string
}

type DetachDhcpOptionsSetFromVpcResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 0ED8D006-F706-4D23-88ED-E11ED28DCAC0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DetachDhcpOptionsSetFromVpcResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DetachDhcpOptionsSetFromVpcResponseBody) GoString() string {
	return s.String()
}

func (s *DetachDhcpOptionsSetFromVpcResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DetachDhcpOptionsSetFromVpcResponseBody) SetRequestId(v string) *DetachDhcpOptionsSetFromVpcResponseBody {
	s.RequestId = &v
	return s
}

func (s *DetachDhcpOptionsSetFromVpcResponseBody) Validate() error {
	return dara.Validate(s)
}
