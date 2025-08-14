// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnassociateCenBandwidthPackageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UnassociateCenBandwidthPackageResponseBody
	GetRequestId() *string
}

type UnassociateCenBandwidthPackageResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 0C2EE7A8-74D4-4081-8236-CEBDE3BBCF50
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UnassociateCenBandwidthPackageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UnassociateCenBandwidthPackageResponseBody) GoString() string {
	return s.String()
}

func (s *UnassociateCenBandwidthPackageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UnassociateCenBandwidthPackageResponseBody) SetRequestId(v string) *UnassociateCenBandwidthPackageResponseBody {
	s.RequestId = &v
	return s
}

func (s *UnassociateCenBandwidthPackageResponseBody) Validate() error {
	return dara.Validate(s)
}
