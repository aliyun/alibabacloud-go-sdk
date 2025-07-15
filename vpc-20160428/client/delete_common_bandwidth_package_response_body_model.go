// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCommonBandwidthPackageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteCommonBandwidthPackageResponseBody
	GetRequestId() *string
}

type DeleteCommonBandwidthPackageResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// B400EF57-60E3-4D61-B8FB-7FA8F72DF5A6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteCommonBandwidthPackageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteCommonBandwidthPackageResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCommonBandwidthPackageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteCommonBandwidthPackageResponseBody) SetRequestId(v string) *DeleteCommonBandwidthPackageResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteCommonBandwidthPackageResponseBody) Validate() error {
	return dara.Validate(s)
}
