// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnRealTimeDeliveryFieldRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBusinessType(v string) *DescribeDcdnRealTimeDeliveryFieldRequest
	GetBusinessType() *string
}

type DescribeDcdnRealTimeDeliveryFieldRequest struct {
	// The type of the collected logs. Default value: cdn_log_access_l1. Valid values:
	//
	// 	- **cdn_log_access_l1**: access logs of Dynamic Content Delivery Network (DCDN) points of presence (POPs)
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

func (s DescribeDcdnRealTimeDeliveryFieldRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnRealTimeDeliveryFieldRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnRealTimeDeliveryFieldRequest) GetBusinessType() *string {
	return s.BusinessType
}

func (s *DescribeDcdnRealTimeDeliveryFieldRequest) SetBusinessType(v string) *DescribeDcdnRealTimeDeliveryFieldRequest {
	s.BusinessType = &v
	return s
}

func (s *DescribeDcdnRealTimeDeliveryFieldRequest) Validate() error {
	return dara.Validate(s)
}
