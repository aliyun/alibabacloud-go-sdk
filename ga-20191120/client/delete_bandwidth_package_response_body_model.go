// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBandwidthPackageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBandwidthPackageId(v string) *DeleteBandwidthPackageResponseBody
	GetBandwidthPackageId() *string
	SetRequestId(v string) *DeleteBandwidthPackageResponseBody
	GetRequestId() *string
}

type DeleteBandwidthPackageResponseBody struct {
	// The bandwidth plan ID.
	//
	// example:
	//
	// gbwp-bp1sgzldyj6b4q7cx****
	BandwidthPackageId *string `json:"BandwidthPackageId,omitempty" xml:"BandwidthPackageId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 6FEA0CF3-D3B9-43E5-A304-D217037876A8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteBandwidthPackageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteBandwidthPackageResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteBandwidthPackageResponseBody) GetBandwidthPackageId() *string {
	return s.BandwidthPackageId
}

func (s *DeleteBandwidthPackageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteBandwidthPackageResponseBody) SetBandwidthPackageId(v string) *DeleteBandwidthPackageResponseBody {
	s.BandwidthPackageId = &v
	return s
}

func (s *DeleteBandwidthPackageResponseBody) SetRequestId(v string) *DeleteBandwidthPackageResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteBandwidthPackageResponseBody) Validate() error {
	return dara.Validate(s)
}
