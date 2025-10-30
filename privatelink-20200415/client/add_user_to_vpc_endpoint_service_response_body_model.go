// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddUserToVpcEndpointServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AddUserToVpcEndpointServiceResponseBody
	GetRequestId() *string
}

type AddUserToVpcEndpointServiceResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 0ED8D006-F706-4D23-88ED-E11ED28DCAC0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddUserToVpcEndpointServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddUserToVpcEndpointServiceResponseBody) GoString() string {
	return s.String()
}

func (s *AddUserToVpcEndpointServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddUserToVpcEndpointServiceResponseBody) SetRequestId(v string) *AddUserToVpcEndpointServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddUserToVpcEndpointServiceResponseBody) Validate() error {
	return dara.Validate(s)
}
