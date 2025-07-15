// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteIpv6InternetBandwidthResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteIpv6InternetBandwidthResponseBody
	GetRequestId() *string
}

type DeleteIpv6InternetBandwidthResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// E07E0FE6-5C21-405F-AF82-7613AA81EF92
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteIpv6InternetBandwidthResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteIpv6InternetBandwidthResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteIpv6InternetBandwidthResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteIpv6InternetBandwidthResponseBody) SetRequestId(v string) *DeleteIpv6InternetBandwidthResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteIpv6InternetBandwidthResponseBody) Validate() error {
	return dara.Validate(s)
}
