// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryOrderPriceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryOrderPriceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryOrderPriceResponse
	GetStatusCode() *int32
	SetBody(v *QueryOrderPriceResponseBody) *QueryOrderPriceResponse
	GetBody() *QueryOrderPriceResponseBody
}

type QueryOrderPriceResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryOrderPriceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryOrderPriceResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryOrderPriceResponse) GoString() string {
	return s.String()
}

func (s *QueryOrderPriceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryOrderPriceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryOrderPriceResponse) GetBody() *QueryOrderPriceResponseBody {
	return s.Body
}

func (s *QueryOrderPriceResponse) SetHeaders(v map[string]*string) *QueryOrderPriceResponse {
	s.Headers = v
	return s
}

func (s *QueryOrderPriceResponse) SetStatusCode(v int32) *QueryOrderPriceResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryOrderPriceResponse) SetBody(v *QueryOrderPriceResponseBody) *QueryOrderPriceResponse {
	s.Body = v
	return s
}

func (s *QueryOrderPriceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
