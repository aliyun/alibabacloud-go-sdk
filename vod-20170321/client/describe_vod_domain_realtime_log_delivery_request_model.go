// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodDomainRealtimeLogDeliveryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeVodDomainRealtimeLogDeliveryRequest
	GetDomainName() *string
	SetOwnerId(v int64) *DescribeVodDomainRealtimeLogDeliveryRequest
	GetOwnerId() *int64
}

type DescribeVodDomainRealtimeLogDeliveryRequest struct {
	// This parameter is required.
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s DescribeVodDomainRealtimeLogDeliveryRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainRealtimeLogDeliveryRequest) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainRealtimeLogDeliveryRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeVodDomainRealtimeLogDeliveryRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeVodDomainRealtimeLogDeliveryRequest) SetDomainName(v string) *DescribeVodDomainRealtimeLogDeliveryRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeVodDomainRealtimeLogDeliveryRequest) SetOwnerId(v int64) *DescribeVodDomainRealtimeLogDeliveryRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeVodDomainRealtimeLogDeliveryRequest) Validate() error {
	return dara.Validate(s)
}
