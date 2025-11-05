// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCouponTemplateDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetCouponTemplateDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetCouponTemplateDetailResponse
	GetStatusCode() *int32
	SetBody(v *GetCouponTemplateDetailResponseBody) *GetCouponTemplateDetailResponse
	GetBody() *GetCouponTemplateDetailResponseBody
}

type GetCouponTemplateDetailResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCouponTemplateDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCouponTemplateDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s GetCouponTemplateDetailResponse) GoString() string {
	return s.String()
}

func (s *GetCouponTemplateDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetCouponTemplateDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetCouponTemplateDetailResponse) GetBody() *GetCouponTemplateDetailResponseBody {
	return s.Body
}

func (s *GetCouponTemplateDetailResponse) SetHeaders(v map[string]*string) *GetCouponTemplateDetailResponse {
	s.Headers = v
	return s
}

func (s *GetCouponTemplateDetailResponse) SetStatusCode(v int32) *GetCouponTemplateDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCouponTemplateDetailResponse) SetBody(v *GetCouponTemplateDetailResponseBody) *GetCouponTemplateDetailResponse {
	s.Body = v
	return s
}

func (s *GetCouponTemplateDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
