// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddZoneToVpcEndpointResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AddZoneToVpcEndpointResponseBody
	GetRequestId() *string
}

type AddZoneToVpcEndpointResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 0c593ea1-3bea-11e9-b96b-88e9fe637760
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddZoneToVpcEndpointResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddZoneToVpcEndpointResponseBody) GoString() string {
	return s.String()
}

func (s *AddZoneToVpcEndpointResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddZoneToVpcEndpointResponseBody) SetRequestId(v string) *AddZoneToVpcEndpointResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddZoneToVpcEndpointResponseBody) Validate() error {
	return dara.Validate(s)
}
