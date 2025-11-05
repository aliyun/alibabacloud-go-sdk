// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCoupondeductProductCodeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *GetCoupondeductProductCodeRequest
	GetAcceptLanguage() *string
}

type GetCoupondeductProductCodeRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// zh-CN
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
}

func (s GetCoupondeductProductCodeRequest) String() string {
	return dara.Prettify(s)
}

func (s GetCoupondeductProductCodeRequest) GoString() string {
	return s.String()
}

func (s *GetCoupondeductProductCodeRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *GetCoupondeductProductCodeRequest) SetAcceptLanguage(v string) *GetCoupondeductProductCodeRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *GetCoupondeductProductCodeRequest) Validate() error {
	return dara.Validate(s)
}
