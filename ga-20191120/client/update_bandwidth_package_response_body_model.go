// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateBandwidthPackageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBandwidthPackage(v string) *UpdateBandwidthPackageResponseBody
	GetBandwidthPackage() *string
	SetDescription(v string) *UpdateBandwidthPackageResponseBody
	GetDescription() *string
	SetName(v string) *UpdateBandwidthPackageResponseBody
	GetName() *string
	SetRequestId(v string) *UpdateBandwidthPackageResponseBody
	GetRequestId() *string
}

type UpdateBandwidthPackageResponseBody struct {
	// The bandwidth plan ID.
	//
	// example:
	//
	// gbwp-bp1eo4f57z1kbbcmn****
	BandwidthPackage *string `json:"BandwidthPackage,omitempty" xml:"BandwidthPackage,omitempty"`
	// The description of the bandwidth plan.
	//
	// example:
	//
	// testDescription
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the bandwidth plan.
	//
	// example:
	//
	// testName
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1DF3A3CB-B621-44F8-9870-C20D034D7AB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateBandwidthPackageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateBandwidthPackageResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateBandwidthPackageResponseBody) GetBandwidthPackage() *string {
	return s.BandwidthPackage
}

func (s *UpdateBandwidthPackageResponseBody) GetDescription() *string {
	return s.Description
}

func (s *UpdateBandwidthPackageResponseBody) GetName() *string {
	return s.Name
}

func (s *UpdateBandwidthPackageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateBandwidthPackageResponseBody) SetBandwidthPackage(v string) *UpdateBandwidthPackageResponseBody {
	s.BandwidthPackage = &v
	return s
}

func (s *UpdateBandwidthPackageResponseBody) SetDescription(v string) *UpdateBandwidthPackageResponseBody {
	s.Description = &v
	return s
}

func (s *UpdateBandwidthPackageResponseBody) SetName(v string) *UpdateBandwidthPackageResponseBody {
	s.Name = &v
	return s
}

func (s *UpdateBandwidthPackageResponseBody) SetRequestId(v string) *UpdateBandwidthPackageResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateBandwidthPackageResponseBody) Validate() error {
	return dara.Validate(s)
}
