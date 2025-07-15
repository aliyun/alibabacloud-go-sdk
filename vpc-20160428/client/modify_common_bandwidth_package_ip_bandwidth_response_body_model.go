// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCommonBandwidthPackageIpBandwidthResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyCommonBandwidthPackageIpBandwidthResponseBody
	GetRequestId() *string
}

type ModifyCommonBandwidthPackageIpBandwidthResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 63D187BF-A30A-4DD6-B68D-FF182C96D8A2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyCommonBandwidthPackageIpBandwidthResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyCommonBandwidthPackageIpBandwidthResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyCommonBandwidthPackageIpBandwidthResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyCommonBandwidthPackageIpBandwidthResponseBody) SetRequestId(v string) *ModifyCommonBandwidthPackageIpBandwidthResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyCommonBandwidthPackageIpBandwidthResponseBody) Validate() error {
	return dara.Validate(s)
}
