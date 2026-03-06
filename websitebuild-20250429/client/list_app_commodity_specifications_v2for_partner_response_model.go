// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAppCommoditySpecificationsV2ForPartnerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAppCommoditySpecificationsV2ForPartnerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAppCommoditySpecificationsV2ForPartnerResponse
	GetStatusCode() *int32
	SetBody(v *ListAppCommoditySpecificationsV2ForPartnerResponseBody) *ListAppCommoditySpecificationsV2ForPartnerResponse
	GetBody() *ListAppCommoditySpecificationsV2ForPartnerResponseBody
}

type ListAppCommoditySpecificationsV2ForPartnerResponse struct {
	Headers    map[string]*string                                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAppCommoditySpecificationsV2ForPartnerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAppCommoditySpecificationsV2ForPartnerResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAppCommoditySpecificationsV2ForPartnerResponse) GoString() string {
	return s.String()
}

func (s *ListAppCommoditySpecificationsV2ForPartnerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAppCommoditySpecificationsV2ForPartnerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAppCommoditySpecificationsV2ForPartnerResponse) GetBody() *ListAppCommoditySpecificationsV2ForPartnerResponseBody {
	return s.Body
}

func (s *ListAppCommoditySpecificationsV2ForPartnerResponse) SetHeaders(v map[string]*string) *ListAppCommoditySpecificationsV2ForPartnerResponse {
	s.Headers = v
	return s
}

func (s *ListAppCommoditySpecificationsV2ForPartnerResponse) SetStatusCode(v int32) *ListAppCommoditySpecificationsV2ForPartnerResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAppCommoditySpecificationsV2ForPartnerResponse) SetBody(v *ListAppCommoditySpecificationsV2ForPartnerResponseBody) *ListAppCommoditySpecificationsV2ForPartnerResponse {
	s.Body = v
	return s
}

func (s *ListAppCommoditySpecificationsV2ForPartnerResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
