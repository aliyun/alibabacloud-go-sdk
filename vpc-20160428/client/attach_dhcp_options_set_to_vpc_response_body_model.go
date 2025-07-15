// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachDhcpOptionsSetToVpcResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AttachDhcpOptionsSetToVpcResponseBody
	GetRequestId() *string
}

type AttachDhcpOptionsSetToVpcResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 0ED8D006-F706-4D23-88ED-E11ED28DCAC0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AttachDhcpOptionsSetToVpcResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AttachDhcpOptionsSetToVpcResponseBody) GoString() string {
	return s.String()
}

func (s *AttachDhcpOptionsSetToVpcResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AttachDhcpOptionsSetToVpcResponseBody) SetRequestId(v string) *AttachDhcpOptionsSetToVpcResponseBody {
	s.RequestId = &v
	return s
}

func (s *AttachDhcpOptionsSetToVpcResponseBody) Validate() error {
	return dara.Validate(s)
}
