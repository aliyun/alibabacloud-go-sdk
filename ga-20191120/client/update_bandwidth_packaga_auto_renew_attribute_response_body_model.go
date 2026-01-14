// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateBandwidthPackagaAutoRenewAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *UpdateBandwidthPackagaAutoRenewAttributeResponseBody
	GetInstanceId() *string
	SetRequestId(v string) *UpdateBandwidthPackagaAutoRenewAttributeResponseBody
	GetRequestId() *string
}

type UpdateBandwidthPackagaAutoRenewAttributeResponseBody struct {
	// The ID of the bandwidth plan.
	//
	// example:
	//
	// gbwp-bp1sgzldyj6b4q7cx****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// F591955F-5CB5-4CCE-A75D-17CF2085CE22
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateBandwidthPackagaAutoRenewAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateBandwidthPackagaAutoRenewAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateBandwidthPackagaAutoRenewAttributeResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateBandwidthPackagaAutoRenewAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateBandwidthPackagaAutoRenewAttributeResponseBody) SetInstanceId(v string) *UpdateBandwidthPackagaAutoRenewAttributeResponseBody {
	s.InstanceId = &v
	return s
}

func (s *UpdateBandwidthPackagaAutoRenewAttributeResponseBody) SetRequestId(v string) *UpdateBandwidthPackagaAutoRenewAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateBandwidthPackagaAutoRenewAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}
