// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableVodRealtimeLogDeliveryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DisableVodRealtimeLogDeliveryRequest
	GetDomainName() *string
	SetOwnerId(v int64) *DisableVodRealtimeLogDeliveryRequest
	GetOwnerId() *int64
}

type DisableVodRealtimeLogDeliveryRequest struct {
	// This parameter is required.
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s DisableVodRealtimeLogDeliveryRequest) String() string {
	return dara.Prettify(s)
}

func (s DisableVodRealtimeLogDeliveryRequest) GoString() string {
	return s.String()
}

func (s *DisableVodRealtimeLogDeliveryRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DisableVodRealtimeLogDeliveryRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DisableVodRealtimeLogDeliveryRequest) SetDomainName(v string) *DisableVodRealtimeLogDeliveryRequest {
	s.DomainName = &v
	return s
}

func (s *DisableVodRealtimeLogDeliveryRequest) SetOwnerId(v int64) *DisableVodRealtimeLogDeliveryRequest {
	s.OwnerId = &v
	return s
}

func (s *DisableVodRealtimeLogDeliveryRequest) Validate() error {
	return dara.Validate(s)
}
