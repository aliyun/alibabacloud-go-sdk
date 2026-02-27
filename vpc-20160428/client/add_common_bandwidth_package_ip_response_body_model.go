// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddCommonBandwidthPackageIpResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AddCommonBandwidthPackageIpResponseBody
	GetRequestId() *string
}

type AddCommonBandwidthPackageIpResponseBody struct {
	// example:
	//
	// 01FDDD49-C4B7-4D2A-A8E5-A93915C450A6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddCommonBandwidthPackageIpResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddCommonBandwidthPackageIpResponseBody) GoString() string {
	return s.String()
}

func (s *AddCommonBandwidthPackageIpResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddCommonBandwidthPackageIpResponseBody) SetRequestId(v string) *AddCommonBandwidthPackageIpResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddCommonBandwidthPackageIpResponseBody) Validate() error {
	return dara.Validate(s)
}
