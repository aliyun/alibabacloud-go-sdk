// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPromotionActivitiesForPartnerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListPromotionActivitiesForPartnerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListPromotionActivitiesForPartnerResponse
	GetStatusCode() *int32
	SetBody(v *ListPromotionActivitiesForPartnerResponseBody) *ListPromotionActivitiesForPartnerResponse
	GetBody() *ListPromotionActivitiesForPartnerResponseBody
}

type ListPromotionActivitiesForPartnerResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPromotionActivitiesForPartnerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPromotionActivitiesForPartnerResponse) String() string {
	return dara.Prettify(s)
}

func (s ListPromotionActivitiesForPartnerResponse) GoString() string {
	return s.String()
}

func (s *ListPromotionActivitiesForPartnerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListPromotionActivitiesForPartnerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListPromotionActivitiesForPartnerResponse) GetBody() *ListPromotionActivitiesForPartnerResponseBody {
	return s.Body
}

func (s *ListPromotionActivitiesForPartnerResponse) SetHeaders(v map[string]*string) *ListPromotionActivitiesForPartnerResponse {
	s.Headers = v
	return s
}

func (s *ListPromotionActivitiesForPartnerResponse) SetStatusCode(v int32) *ListPromotionActivitiesForPartnerResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPromotionActivitiesForPartnerResponse) SetBody(v *ListPromotionActivitiesForPartnerResponseBody) *ListPromotionActivitiesForPartnerResponse {
	s.Body = v
	return s
}

func (s *ListPromotionActivitiesForPartnerResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
