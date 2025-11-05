// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCouponTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateCouponTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateCouponTemplateResponse
	GetStatusCode() *int32
	SetBody(v *CreateCouponTemplateResponseBody) *CreateCouponTemplateResponse
	GetBody() *CreateCouponTemplateResponseBody
}

type CreateCouponTemplateResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateCouponTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateCouponTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateCouponTemplateResponse) GoString() string {
	return s.String()
}

func (s *CreateCouponTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateCouponTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateCouponTemplateResponse) GetBody() *CreateCouponTemplateResponseBody {
	return s.Body
}

func (s *CreateCouponTemplateResponse) SetHeaders(v map[string]*string) *CreateCouponTemplateResponse {
	s.Headers = v
	return s
}

func (s *CreateCouponTemplateResponse) SetStatusCode(v int32) *CreateCouponTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCouponTemplateResponse) SetBody(v *CreateCouponTemplateResponseBody) *CreateCouponTemplateResponse {
	s.Body = v
	return s
}

func (s *CreateCouponTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
