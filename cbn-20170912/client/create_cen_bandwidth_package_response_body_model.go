// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCenBandwidthPackageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCenBandwidthPackageId(v string) *CreateCenBandwidthPackageResponseBody
	GetCenBandwidthPackageId() *string
	SetCenBandwidthPackageOrderId(v string) *CreateCenBandwidthPackageResponseBody
	GetCenBandwidthPackageOrderId() *string
	SetRequestId(v string) *CreateCenBandwidthPackageResponseBody
	GetRequestId() *string
}

type CreateCenBandwidthPackageResponseBody struct {
	// The ID of the bandwidth plan.
	//
	// example:
	//
	// cenbwp-4c2zaavbvh5fx****
	CenBandwidthPackageId *string `json:"CenBandwidthPackageId,omitempty" xml:"CenBandwidthPackageId,omitempty"`
	// The ID of the order for the bandwidth plan.
	//
	// example:
	//
	// 20156420004****
	CenBandwidthPackageOrderId *string `json:"CenBandwidthPackageOrderId,omitempty" xml:"CenBandwidthPackageOrderId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// E4B345CD-2CBA-4881-AF6D-E5D9BAE1CA7B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateCenBandwidthPackageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateCenBandwidthPackageResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCenBandwidthPackageResponseBody) GetCenBandwidthPackageId() *string {
	return s.CenBandwidthPackageId
}

func (s *CreateCenBandwidthPackageResponseBody) GetCenBandwidthPackageOrderId() *string {
	return s.CenBandwidthPackageOrderId
}

func (s *CreateCenBandwidthPackageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateCenBandwidthPackageResponseBody) SetCenBandwidthPackageId(v string) *CreateCenBandwidthPackageResponseBody {
	s.CenBandwidthPackageId = &v
	return s
}

func (s *CreateCenBandwidthPackageResponseBody) SetCenBandwidthPackageOrderId(v string) *CreateCenBandwidthPackageResponseBody {
	s.CenBandwidthPackageOrderId = &v
	return s
}

func (s *CreateCenBandwidthPackageResponseBody) SetRequestId(v string) *CreateCenBandwidthPackageResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateCenBandwidthPackageResponseBody) Validate() error {
	return dara.Validate(s)
}
