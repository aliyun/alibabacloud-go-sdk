// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPriceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *GetPriceRequest
	GetRegionId() *string
	SetResourceAttributes(v map[string]interface{}) *GetPriceRequest
	GetResourceAttributes() map[string]interface{}
}

type GetPriceRequest struct {
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
	ResourceAttributes map[string]interface{} `json:"resourceAttributes,omitempty" xml:"resourceAttributes,omitempty"`
}

func (s GetPriceRequest) String() string {
	return dara.Prettify(s)
}

func (s GetPriceRequest) GoString() string {
	return s.String()
}

func (s *GetPriceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetPriceRequest) GetResourceAttributes() map[string]interface{} {
	return s.ResourceAttributes
}

func (s *GetPriceRequest) SetRegionId(v string) *GetPriceRequest {
	s.RegionId = &v
	return s
}

func (s *GetPriceRequest) SetResourceAttributes(v map[string]interface{}) *GetPriceRequest {
	s.ResourceAttributes = v
	return s
}

func (s *GetPriceRequest) Validate() error {
	return dara.Validate(s)
}
