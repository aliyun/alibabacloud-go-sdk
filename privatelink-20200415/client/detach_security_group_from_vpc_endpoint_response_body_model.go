// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachSecurityGroupFromVpcEndpointResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DetachSecurityGroupFromVpcEndpointResponseBody
	GetRequestId() *string
}

type DetachSecurityGroupFromVpcEndpointResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 8D778FF9-7640-4C13-BCD6-9265CA9A2F81
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DetachSecurityGroupFromVpcEndpointResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DetachSecurityGroupFromVpcEndpointResponseBody) GoString() string {
	return s.String()
}

func (s *DetachSecurityGroupFromVpcEndpointResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DetachSecurityGroupFromVpcEndpointResponseBody) SetRequestId(v string) *DetachSecurityGroupFromVpcEndpointResponseBody {
	s.RequestId = &v
	return s
}

func (s *DetachSecurityGroupFromVpcEndpointResponseBody) Validate() error {
	return dara.Validate(s)
}
