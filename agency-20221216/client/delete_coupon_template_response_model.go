// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCouponTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteCouponTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteCouponTemplateResponse
	GetStatusCode() *int32
	SetBody(v *DeleteCouponTemplateResponseBody) *DeleteCouponTemplateResponse
	GetBody() *DeleteCouponTemplateResponseBody
}

type DeleteCouponTemplateResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteCouponTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteCouponTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteCouponTemplateResponse) GoString() string {
	return s.String()
}

func (s *DeleteCouponTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteCouponTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteCouponTemplateResponse) GetBody() *DeleteCouponTemplateResponseBody {
	return s.Body
}

func (s *DeleteCouponTemplateResponse) SetHeaders(v map[string]*string) *DeleteCouponTemplateResponse {
	s.Headers = v
	return s
}

func (s *DeleteCouponTemplateResponse) SetStatusCode(v int32) *DeleteCouponTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCouponTemplateResponse) SetBody(v *DeleteCouponTemplateResponseBody) *DeleteCouponTemplateResponse {
	s.Body = v
	return s
}

func (s *DeleteCouponTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
