// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryRefundPriceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryRefundPriceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryRefundPriceResponse
	GetStatusCode() *int32
	SetBody(v *QueryRefundPriceResponseBody) *QueryRefundPriceResponse
	GetBody() *QueryRefundPriceResponseBody
}

type QueryRefundPriceResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryRefundPriceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryRefundPriceResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryRefundPriceResponse) GoString() string {
	return s.String()
}

func (s *QueryRefundPriceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryRefundPriceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryRefundPriceResponse) GetBody() *QueryRefundPriceResponseBody {
	return s.Body
}

func (s *QueryRefundPriceResponse) SetHeaders(v map[string]*string) *QueryRefundPriceResponse {
	s.Headers = v
	return s
}

func (s *QueryRefundPriceResponse) SetStatusCode(v int32) *QueryRefundPriceResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryRefundPriceResponse) SetBody(v *QueryRefundPriceResponseBody) *QueryRefundPriceResponse {
	s.Body = v
	return s
}

func (s *QueryRefundPriceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
