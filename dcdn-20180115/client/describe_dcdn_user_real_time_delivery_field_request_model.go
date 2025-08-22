// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnUserRealTimeDeliveryFieldRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBusinessType(v string) *DescribeDcdnUserRealTimeDeliveryFieldRequest
	GetBusinessType() *string
}

type DescribeDcdnUserRealTimeDeliveryFieldRequest struct {
	// The type of the collected logs. Default value: cdn_log_access_l1. Valid values:
	//
	// 	- **cdn_log_access_l1**: access logs of L1 Dynamic Content Delivery Network (DCDN) points of presence (POPs)
	//
	// 	- **cdn_log_origin**: back-to-origin logs
	//
	// 	- **cdn_log_er**: EdgeRoutine logs
	//
	// example:
	//
	// cdn_log_access_l1
	BusinessType *string `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
}

func (s DescribeDcdnUserRealTimeDeliveryFieldRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnUserRealTimeDeliveryFieldRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnUserRealTimeDeliveryFieldRequest) GetBusinessType() *string {
	return s.BusinessType
}

func (s *DescribeDcdnUserRealTimeDeliveryFieldRequest) SetBusinessType(v string) *DescribeDcdnUserRealTimeDeliveryFieldRequest {
	s.BusinessType = &v
	return s
}

func (s *DescribeDcdnUserRealTimeDeliveryFieldRequest) Validate() error {
	return dara.Validate(s)
}
