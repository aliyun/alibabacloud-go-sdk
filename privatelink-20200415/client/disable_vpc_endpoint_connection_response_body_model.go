// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableVpcEndpointConnectionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DisableVpcEndpointConnectionResponseBody
	GetRequestId() *string
}

type DisableVpcEndpointConnectionResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 0ED8D006-F706-4D23-88ED-E11ED28DCAC0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisableVpcEndpointConnectionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DisableVpcEndpointConnectionResponseBody) GoString() string {
	return s.String()
}

func (s *DisableVpcEndpointConnectionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DisableVpcEndpointConnectionResponseBody) SetRequestId(v string) *DisableVpcEndpointConnectionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisableVpcEndpointConnectionResponseBody) Validate() error {
	return dara.Validate(s)
}
