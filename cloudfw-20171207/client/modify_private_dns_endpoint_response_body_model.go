// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyPrivateDnsEndpointResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyPrivateDnsEndpointResponseBody
	GetRequestId() *string
}

type ModifyPrivateDnsEndpointResponseBody struct {
	// example:
	//
	// 6B780BD6-282C-51A9-A8E6-59F636BAFA54
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyPrivateDnsEndpointResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyPrivateDnsEndpointResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyPrivateDnsEndpointResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyPrivateDnsEndpointResponseBody) SetRequestId(v string) *ModifyPrivateDnsEndpointResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyPrivateDnsEndpointResponseBody) Validate() error {
	return dara.Validate(s)
}
