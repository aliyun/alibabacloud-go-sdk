// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPromotionOfferRecordsForPartnerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListPromotionOfferRecordsForPartnerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListPromotionOfferRecordsForPartnerResponse
	GetStatusCode() *int32
	SetBody(v *ListPromotionOfferRecordsForPartnerResponseBody) *ListPromotionOfferRecordsForPartnerResponse
	GetBody() *ListPromotionOfferRecordsForPartnerResponseBody
}

type ListPromotionOfferRecordsForPartnerResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPromotionOfferRecordsForPartnerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPromotionOfferRecordsForPartnerResponse) String() string {
	return dara.Prettify(s)
}

func (s ListPromotionOfferRecordsForPartnerResponse) GoString() string {
	return s.String()
}

func (s *ListPromotionOfferRecordsForPartnerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListPromotionOfferRecordsForPartnerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListPromotionOfferRecordsForPartnerResponse) GetBody() *ListPromotionOfferRecordsForPartnerResponseBody {
	return s.Body
}

func (s *ListPromotionOfferRecordsForPartnerResponse) SetHeaders(v map[string]*string) *ListPromotionOfferRecordsForPartnerResponse {
	s.Headers = v
	return s
}

func (s *ListPromotionOfferRecordsForPartnerResponse) SetStatusCode(v int32) *ListPromotionOfferRecordsForPartnerResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPromotionOfferRecordsForPartnerResponse) SetBody(v *ListPromotionOfferRecordsForPartnerResponseBody) *ListPromotionOfferRecordsForPartnerResponse {
	s.Body = v
	return s
}

func (s *ListPromotionOfferRecordsForPartnerResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
