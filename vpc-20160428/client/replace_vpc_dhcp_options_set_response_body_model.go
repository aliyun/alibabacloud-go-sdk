// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReplaceVpcDhcpOptionsSetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ReplaceVpcDhcpOptionsSetResponseBody
	GetRequestId() *string
}

type ReplaceVpcDhcpOptionsSetResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 0ED8D006-F706-4D23-88ED-E11ED28DCAC0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ReplaceVpcDhcpOptionsSetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ReplaceVpcDhcpOptionsSetResponseBody) GoString() string {
	return s.String()
}

func (s *ReplaceVpcDhcpOptionsSetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ReplaceVpcDhcpOptionsSetResponseBody) SetRequestId(v string) *ReplaceVpcDhcpOptionsSetResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReplaceVpcDhcpOptionsSetResponseBody) Validate() error {
	return dara.Validate(s)
}
