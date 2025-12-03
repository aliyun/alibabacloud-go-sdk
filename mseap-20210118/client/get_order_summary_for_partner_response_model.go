// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOrderSummaryForPartnerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetOrderSummaryForPartnerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetOrderSummaryForPartnerResponse
	GetStatusCode() *int32
	SetBody(v *GetOrderSummaryForPartnerResponseBody) *GetOrderSummaryForPartnerResponse
	GetBody() *GetOrderSummaryForPartnerResponseBody
}

type GetOrderSummaryForPartnerResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetOrderSummaryForPartnerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetOrderSummaryForPartnerResponse) String() string {
	return dara.Prettify(s)
}

func (s GetOrderSummaryForPartnerResponse) GoString() string {
	return s.String()
}

func (s *GetOrderSummaryForPartnerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetOrderSummaryForPartnerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetOrderSummaryForPartnerResponse) GetBody() *GetOrderSummaryForPartnerResponseBody {
	return s.Body
}

func (s *GetOrderSummaryForPartnerResponse) SetHeaders(v map[string]*string) *GetOrderSummaryForPartnerResponse {
	s.Headers = v
	return s
}

func (s *GetOrderSummaryForPartnerResponse) SetStatusCode(v int32) *GetOrderSummaryForPartnerResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOrderSummaryForPartnerResponse) SetBody(v *GetOrderSummaryForPartnerResponseBody) *GetOrderSummaryForPartnerResponse {
	s.Body = v
	return s
}

func (s *GetOrderSummaryForPartnerResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
