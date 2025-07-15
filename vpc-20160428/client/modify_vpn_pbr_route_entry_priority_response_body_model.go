// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyVpnPbrRouteEntryPriorityResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyVpnPbrRouteEntryPriorityResponseBody
	GetRequestId() *string
}

type ModifyVpnPbrRouteEntryPriorityResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 5BE01CD7-5A50-472D-AC14-CA181C5C03BE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyVpnPbrRouteEntryPriorityResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyVpnPbrRouteEntryPriorityResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyVpnPbrRouteEntryPriorityResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyVpnPbrRouteEntryPriorityResponseBody) SetRequestId(v string) *ModifyVpnPbrRouteEntryPriorityResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyVpnPbrRouteEntryPriorityResponseBody) Validate() error {
	return dara.Validate(s)
}
