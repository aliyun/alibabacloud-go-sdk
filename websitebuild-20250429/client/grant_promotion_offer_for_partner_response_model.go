// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGrantPromotionOfferForPartnerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GrantPromotionOfferForPartnerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GrantPromotionOfferForPartnerResponse
	GetStatusCode() *int32
	SetBody(v *GrantPromotionOfferForPartnerResponseBody) *GrantPromotionOfferForPartnerResponse
	GetBody() *GrantPromotionOfferForPartnerResponseBody
}

type GrantPromotionOfferForPartnerResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GrantPromotionOfferForPartnerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GrantPromotionOfferForPartnerResponse) String() string {
	return dara.Prettify(s)
}

func (s GrantPromotionOfferForPartnerResponse) GoString() string {
	return s.String()
}

func (s *GrantPromotionOfferForPartnerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GrantPromotionOfferForPartnerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GrantPromotionOfferForPartnerResponse) GetBody() *GrantPromotionOfferForPartnerResponseBody {
	return s.Body
}

func (s *GrantPromotionOfferForPartnerResponse) SetHeaders(v map[string]*string) *GrantPromotionOfferForPartnerResponse {
	s.Headers = v
	return s
}

func (s *GrantPromotionOfferForPartnerResponse) SetStatusCode(v int32) *GrantPromotionOfferForPartnerResponse {
	s.StatusCode = &v
	return s
}

func (s *GrantPromotionOfferForPartnerResponse) SetBody(v *GrantPromotionOfferForPartnerResponseBody) *GrantPromotionOfferForPartnerResponse {
	s.Body = v
	return s
}

func (s *GrantPromotionOfferForPartnerResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
