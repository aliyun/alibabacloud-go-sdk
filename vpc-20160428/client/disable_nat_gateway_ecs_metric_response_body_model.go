// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableNatGatewayEcsMetricResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DisableNatGatewayEcsMetricResponseBody
	GetRequestId() *string
}

type DisableNatGatewayEcsMetricResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 0ED8D006-F706-4D23-88ED-E11ED28DCAC0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisableNatGatewayEcsMetricResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DisableNatGatewayEcsMetricResponseBody) GoString() string {
	return s.String()
}

func (s *DisableNatGatewayEcsMetricResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DisableNatGatewayEcsMetricResponseBody) SetRequestId(v string) *DisableNatGatewayEcsMetricResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisableNatGatewayEcsMetricResponseBody) Validate() error {
	return dara.Validate(s)
}
