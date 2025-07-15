// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyVcoRouteEntryWeightResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyVcoRouteEntryWeightResponseBody
	GetRequestId() *string
}

type ModifyVcoRouteEntryWeightResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 9496F8A-82F4-3130-A51A-2266ACC799B4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyVcoRouteEntryWeightResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyVcoRouteEntryWeightResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyVcoRouteEntryWeightResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyVcoRouteEntryWeightResponseBody) SetRequestId(v string) *ModifyVcoRouteEntryWeightResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyVcoRouteEntryWeightResponseBody) Validate() error {
	return dara.Validate(s)
}
