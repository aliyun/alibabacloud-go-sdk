// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInquiryPriceRefundInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *InquiryPriceRefundInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *InquiryPriceRefundInstanceResponse
	GetStatusCode() *int32
	SetBody(v *InquiryPriceRefundInstanceResponseBody) *InquiryPriceRefundInstanceResponse
	GetBody() *InquiryPriceRefundInstanceResponseBody
}

type InquiryPriceRefundInstanceResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InquiryPriceRefundInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InquiryPriceRefundInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s InquiryPriceRefundInstanceResponse) GoString() string {
	return s.String()
}

func (s *InquiryPriceRefundInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *InquiryPriceRefundInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *InquiryPriceRefundInstanceResponse) GetBody() *InquiryPriceRefundInstanceResponseBody {
	return s.Body
}

func (s *InquiryPriceRefundInstanceResponse) SetHeaders(v map[string]*string) *InquiryPriceRefundInstanceResponse {
	s.Headers = v
	return s
}

func (s *InquiryPriceRefundInstanceResponse) SetStatusCode(v int32) *InquiryPriceRefundInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *InquiryPriceRefundInstanceResponse) SetBody(v *InquiryPriceRefundInstanceResponseBody) *InquiryPriceRefundInstanceResponse {
	s.Body = v
	return s
}

func (s *InquiryPriceRefundInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
