// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetCenInterRegionBandwidthLimitResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SetCenInterRegionBandwidthLimitResponseBody
	GetRequestId() *string
}

type SetCenInterRegionBandwidthLimitResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 530BC816-F575-412A-AAB2-435125D26328
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetCenInterRegionBandwidthLimitResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetCenInterRegionBandwidthLimitResponseBody) GoString() string {
	return s.String()
}

func (s *SetCenInterRegionBandwidthLimitResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetCenInterRegionBandwidthLimitResponseBody) SetRequestId(v string) *SetCenInterRegionBandwidthLimitResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetCenInterRegionBandwidthLimitResponseBody) Validate() error {
	return dara.Validate(s)
}
