// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyVpnPbrRouteEntryAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyVpnPbrRouteEntryAttributeResponseBody
	GetRequestId() *string
}

type ModifyVpnPbrRouteEntryAttributeResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 5BE01CD7-5A50-472D-AC14-CA181C5C03BE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyVpnPbrRouteEntryAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyVpnPbrRouteEntryAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyVpnPbrRouteEntryAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyVpnPbrRouteEntryAttributeResponseBody) SetRequestId(v string) *ModifyVpnPbrRouteEntryAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyVpnPbrRouteEntryAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}
