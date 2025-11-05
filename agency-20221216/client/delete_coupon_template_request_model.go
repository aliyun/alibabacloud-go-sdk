// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCouponTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTemplateId(v int64) *DeleteCouponTemplateRequest
	GetTemplateId() *int64
}

type DeleteCouponTemplateRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 6558410265670417297
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s DeleteCouponTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteCouponTemplateRequest) GoString() string {
	return s.String()
}

func (s *DeleteCouponTemplateRequest) GetTemplateId() *int64 {
	return s.TemplateId
}

func (s *DeleteCouponTemplateRequest) SetTemplateId(v int64) *DeleteCouponTemplateRequest {
	s.TemplateId = &v
	return s
}

func (s *DeleteCouponTemplateRequest) Validate() error {
	return dara.Validate(s)
}
