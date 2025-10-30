// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableVpcEndpointZoneConnectionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DisableVpcEndpointZoneConnectionResponseBody
	GetRequestId() *string
}

type DisableVpcEndpointZoneConnectionResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 0ED8D006-F706-4D23-88ED-E11ED28DCAC0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisableVpcEndpointZoneConnectionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DisableVpcEndpointZoneConnectionResponseBody) GoString() string {
	return s.String()
}

func (s *DisableVpcEndpointZoneConnectionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DisableVpcEndpointZoneConnectionResponseBody) SetRequestId(v string) *DisableVpcEndpointZoneConnectionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisableVpcEndpointZoneConnectionResponseBody) Validate() error {
	return dara.Validate(s)
}
