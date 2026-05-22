// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRealtimeDeliveryFieldRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBusinessType(v string) *GetRealtimeDeliveryFieldRequest
	GetBusinessType() *string
}

type GetRealtimeDeliveryFieldRequest struct {
	// This parameter is required.
	BusinessType *string `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
}

func (s GetRealtimeDeliveryFieldRequest) String() string {
	return dara.Prettify(s)
}

func (s GetRealtimeDeliveryFieldRequest) GoString() string {
	return s.String()
}

func (s *GetRealtimeDeliveryFieldRequest) GetBusinessType() *string {
	return s.BusinessType
}

func (s *GetRealtimeDeliveryFieldRequest) SetBusinessType(v string) *GetRealtimeDeliveryFieldRequest {
	s.BusinessType = &v
	return s
}

func (s *GetRealtimeDeliveryFieldRequest) Validate() error {
	return dara.Validate(s)
}
