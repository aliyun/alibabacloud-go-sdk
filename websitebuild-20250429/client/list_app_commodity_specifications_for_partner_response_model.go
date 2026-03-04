// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAppCommoditySpecificationsForPartnerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAppCommoditySpecificationsForPartnerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAppCommoditySpecificationsForPartnerResponse
	GetStatusCode() *int32
	SetBody(v *ListAppCommoditySpecificationsForPartnerResponseBody) *ListAppCommoditySpecificationsForPartnerResponse
	GetBody() *ListAppCommoditySpecificationsForPartnerResponseBody
}

type ListAppCommoditySpecificationsForPartnerResponse struct {
	Headers    map[string]*string                                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAppCommoditySpecificationsForPartnerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAppCommoditySpecificationsForPartnerResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAppCommoditySpecificationsForPartnerResponse) GoString() string {
	return s.String()
}

func (s *ListAppCommoditySpecificationsForPartnerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAppCommoditySpecificationsForPartnerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAppCommoditySpecificationsForPartnerResponse) GetBody() *ListAppCommoditySpecificationsForPartnerResponseBody {
	return s.Body
}

func (s *ListAppCommoditySpecificationsForPartnerResponse) SetHeaders(v map[string]*string) *ListAppCommoditySpecificationsForPartnerResponse {
	s.Headers = v
	return s
}

func (s *ListAppCommoditySpecificationsForPartnerResponse) SetStatusCode(v int32) *ListAppCommoditySpecificationsForPartnerResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAppCommoditySpecificationsForPartnerResponse) SetBody(v *ListAppCommoditySpecificationsForPartnerResponseBody) *ListAppCommoditySpecificationsForPartnerResponse {
	s.Body = v
	return s
}

func (s *ListAppCommoditySpecificationsForPartnerResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
