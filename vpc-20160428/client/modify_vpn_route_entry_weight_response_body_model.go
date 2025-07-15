// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyVpnRouteEntryWeightResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyVpnRouteEntryWeightResponseBody
	GetRequestId() *string
}

type ModifyVpnRouteEntryWeightResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 5BE01CD7-5A50-472D-AC14-CA181C5C03BE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyVpnRouteEntryWeightResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyVpnRouteEntryWeightResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyVpnRouteEntryWeightResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyVpnRouteEntryWeightResponseBody) SetRequestId(v string) *ModifyVpnRouteEntryWeightResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyVpnRouteEntryWeightResponseBody) Validate() error {
	return dara.Validate(s)
}
