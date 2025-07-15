// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddCommonBandwidthPackageIpsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AddCommonBandwidthPackageIpsResponseBody
	GetRequestId() *string
}

type AddCommonBandwidthPackageIpsResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 01FDDD49-C4B7-4D2A-A8E5-A93915C450A6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddCommonBandwidthPackageIpsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddCommonBandwidthPackageIpsResponseBody) GoString() string {
	return s.String()
}

func (s *AddCommonBandwidthPackageIpsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddCommonBandwidthPackageIpsResponseBody) SetRequestId(v string) *AddCommonBandwidthPackageIpsResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddCommonBandwidthPackageIpsResponseBody) Validate() error {
	return dara.Validate(s)
}
