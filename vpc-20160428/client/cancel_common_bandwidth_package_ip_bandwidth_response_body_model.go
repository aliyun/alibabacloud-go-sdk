// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelCommonBandwidthPackageIpBandwidthResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CancelCommonBandwidthPackageIpBandwidthResponseBody
	GetRequestId() *string
}

type CancelCommonBandwidthPackageIpBandwidthResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 63D187BF-A30A-4DD6-B68D-FF182C96D8A2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CancelCommonBandwidthPackageIpBandwidthResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CancelCommonBandwidthPackageIpBandwidthResponseBody) GoString() string {
	return s.String()
}

func (s *CancelCommonBandwidthPackageIpBandwidthResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CancelCommonBandwidthPackageIpBandwidthResponseBody) SetRequestId(v string) *CancelCommonBandwidthPackageIpBandwidthResponseBody {
	s.RequestId = &v
	return s
}

func (s *CancelCommonBandwidthPackageIpBandwidthResponseBody) Validate() error {
	return dara.Validate(s)
}
