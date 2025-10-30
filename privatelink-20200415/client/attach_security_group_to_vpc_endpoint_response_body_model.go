// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachSecurityGroupToVpcEndpointResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AttachSecurityGroupToVpcEndpointResponseBody
	GetRequestId() *string
}

type AttachSecurityGroupToVpcEndpointResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 8D778FF9-7640-4C13-BCD6-9265CA9A2F81
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AttachSecurityGroupToVpcEndpointResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AttachSecurityGroupToVpcEndpointResponseBody) GoString() string {
	return s.String()
}

func (s *AttachSecurityGroupToVpcEndpointResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AttachSecurityGroupToVpcEndpointResponseBody) SetRequestId(v string) *AttachSecurityGroupToVpcEndpointResponseBody {
	s.RequestId = &v
	return s
}

func (s *AttachSecurityGroupToVpcEndpointResponseBody) Validate() error {
	return dara.Validate(s)
}
