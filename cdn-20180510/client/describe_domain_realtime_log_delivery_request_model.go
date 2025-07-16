// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainRealtimeLogDeliveryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomain(v string) *DescribeDomainRealtimeLogDeliveryRequest
	GetDomain() *string
}

type DescribeDomainRealtimeLogDeliveryRequest struct {
	// The accelerated domain name for which real-time log delivery is enabled. You can specify only one domain name.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
}

func (s DescribeDomainRealtimeLogDeliveryRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainRealtimeLogDeliveryRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainRealtimeLogDeliveryRequest) GetDomain() *string {
	return s.Domain
}

func (s *DescribeDomainRealtimeLogDeliveryRequest) SetDomain(v string) *DescribeDomainRealtimeLogDeliveryRequest {
	s.Domain = &v
	return s
}

func (s *DescribeDomainRealtimeLogDeliveryRequest) Validate() error {
	return dara.Validate(s)
}
