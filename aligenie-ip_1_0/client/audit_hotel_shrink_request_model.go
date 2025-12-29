// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAuditHotelShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuditHotelReqShrink(v string) *AuditHotelShrinkRequest
	GetAuditHotelReqShrink() *string
}

type AuditHotelShrinkRequest struct {
	// This parameter is required.
	AuditHotelReqShrink *string `json:"AuditHotelReq,omitempty" xml:"AuditHotelReq,omitempty"`
}

func (s AuditHotelShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s AuditHotelShrinkRequest) GoString() string {
	return s.String()
}

func (s *AuditHotelShrinkRequest) GetAuditHotelReqShrink() *string {
	return s.AuditHotelReqShrink
}

func (s *AuditHotelShrinkRequest) SetAuditHotelReqShrink(v string) *AuditHotelShrinkRequest {
	s.AuditHotelReqShrink = &v
	return s
}

func (s *AuditHotelShrinkRequest) Validate() error {
	return dara.Validate(s)
}
