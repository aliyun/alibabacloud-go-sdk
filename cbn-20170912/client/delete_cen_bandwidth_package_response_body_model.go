// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCenBandwidthPackageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteCenBandwidthPackageResponseBody
	GetRequestId() *string
}

type DeleteCenBandwidthPackageResponseBody struct {
	// example:
	//
	// C0245BEF-52AC-44A8-A776-EF96FD26A5CA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteCenBandwidthPackageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteCenBandwidthPackageResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCenBandwidthPackageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteCenBandwidthPackageResponseBody) SetRequestId(v string) *DeleteCenBandwidthPackageResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteCenBandwidthPackageResponseBody) Validate() error {
	return dara.Validate(s)
}
