// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveZoneFromVpcEndpointResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RemoveZoneFromVpcEndpointResponseBody
	GetRequestId() *string
}

type RemoveZoneFromVpcEndpointResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 0c593ea1-3bea-11e9-b96b-88e9fe637760
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RemoveZoneFromVpcEndpointResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RemoveZoneFromVpcEndpointResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveZoneFromVpcEndpointResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RemoveZoneFromVpcEndpointResponseBody) SetRequestId(v string) *RemoveZoneFromVpcEndpointResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveZoneFromVpcEndpointResponseBody) Validate() error {
	return dara.Validate(s)
}
