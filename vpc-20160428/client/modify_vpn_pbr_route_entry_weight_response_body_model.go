// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyVpnPbrRouteEntryWeightResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyVpnPbrRouteEntryWeightResponseBody
	GetRequestId() *string
}

type ModifyVpnPbrRouteEntryWeightResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 5BE01CD7-5A50-472D-AC14-CA181C5C03BE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyVpnPbrRouteEntryWeightResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyVpnPbrRouteEntryWeightResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyVpnPbrRouteEntryWeightResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyVpnPbrRouteEntryWeightResponseBody) SetRequestId(v string) *ModifyVpnPbrRouteEntryWeightResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyVpnPbrRouteEntryWeightResponseBody) Validate() error {
	return dara.Validate(s)
}
