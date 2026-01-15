// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRefundAppInstanceForPartnerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RefundAppInstanceForPartnerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RefundAppInstanceForPartnerResponse
	GetStatusCode() *int32
	SetBody(v *RefundAppInstanceForPartnerResponseBody) *RefundAppInstanceForPartnerResponse
	GetBody() *RefundAppInstanceForPartnerResponseBody
}

type RefundAppInstanceForPartnerResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RefundAppInstanceForPartnerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RefundAppInstanceForPartnerResponse) String() string {
	return dara.Prettify(s)
}

func (s RefundAppInstanceForPartnerResponse) GoString() string {
	return s.String()
}

func (s *RefundAppInstanceForPartnerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RefundAppInstanceForPartnerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RefundAppInstanceForPartnerResponse) GetBody() *RefundAppInstanceForPartnerResponseBody {
	return s.Body
}

func (s *RefundAppInstanceForPartnerResponse) SetHeaders(v map[string]*string) *RefundAppInstanceForPartnerResponse {
	s.Headers = v
	return s
}

func (s *RefundAppInstanceForPartnerResponse) SetStatusCode(v int32) *RefundAppInstanceForPartnerResponse {
	s.StatusCode = &v
	return s
}

func (s *RefundAppInstanceForPartnerResponse) SetBody(v *RefundAppInstanceForPartnerResponseBody) *RefundAppInstanceForPartnerResponse {
	s.Body = v
	return s
}

func (s *RefundAppInstanceForPartnerResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
