// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveUserFromVpcEndpointServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RemoveUserFromVpcEndpointServiceResponseBody
	GetRequestId() *string
}

type RemoveUserFromVpcEndpointServiceResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 0ED8D006-F706-4D23-88ED-E11ED28DCAC0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RemoveUserFromVpcEndpointServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RemoveUserFromVpcEndpointServiceResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveUserFromVpcEndpointServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RemoveUserFromVpcEndpointServiceResponseBody) SetRequestId(v string) *RemoveUserFromVpcEndpointServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveUserFromVpcEndpointServiceResponseBody) Validate() error {
	return dara.Validate(s)
}
