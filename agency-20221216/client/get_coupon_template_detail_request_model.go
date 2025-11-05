// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCouponTemplateDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTemplateId(v int64) *GetCouponTemplateDetailRequest
	GetTemplateId() *int64
}

type GetCouponTemplateDetailRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 5093156
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s GetCouponTemplateDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s GetCouponTemplateDetailRequest) GoString() string {
	return s.String()
}

func (s *GetCouponTemplateDetailRequest) GetTemplateId() *int64 {
	return s.TemplateId
}

func (s *GetCouponTemplateDetailRequest) SetTemplateId(v int64) *GetCouponTemplateDetailRequest {
	s.TemplateId = &v
	return s
}

func (s *GetCouponTemplateDetailRequest) Validate() error {
	return dara.Validate(s)
}
