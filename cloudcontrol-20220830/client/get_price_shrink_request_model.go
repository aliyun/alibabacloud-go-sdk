// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPriceShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *GetPriceShrinkRequest
	GetRegionId() *string
	SetResourceAttributesShrink(v string) *GetPriceShrinkRequest
	GetResourceAttributesShrink() *string
}

type GetPriceShrinkRequest struct {
	// The region ID. This parameter is required if the cloud product is deployed in a region.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// The attributes based on which the price is queried (in JSON format).
	//
	// example:
	//
	// {
	//
	//         "LoadBalancerName": "cc-test",
	//
	//         "LoadBalancerSpec": "slb.s3.small",
	//
	//         "InternetChargeType": "paybybandwidth",
	//
	//         "AddressType": "internet",
	//
	//         "PaymentType": "PayAsYouGo",
	//
	//         "Bandwidth": 6
	//
	//       }
	ResourceAttributesShrink *string `json:"resourceAttributes,omitempty" xml:"resourceAttributes,omitempty"`
}

func (s GetPriceShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetPriceShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetPriceShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetPriceShrinkRequest) GetResourceAttributesShrink() *string {
	return s.ResourceAttributesShrink
}

func (s *GetPriceShrinkRequest) SetRegionId(v string) *GetPriceShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *GetPriceShrinkRequest) SetResourceAttributesShrink(v string) *GetPriceShrinkRequest {
	s.ResourceAttributesShrink = &v
	return s
}

func (s *GetPriceShrinkRequest) Validate() error {
	return dara.Validate(s)
}
