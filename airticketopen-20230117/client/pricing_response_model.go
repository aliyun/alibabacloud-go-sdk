// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPricingResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PricingResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PricingResponse
	GetStatusCode() *int32
	SetBody(v *PricingResponseBody) *PricingResponse
	GetBody() *PricingResponseBody
}

type PricingResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PricingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PricingResponse) String() string {
	return dara.Prettify(s)
}

func (s PricingResponse) GoString() string {
	return s.String()
}

func (s *PricingResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PricingResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PricingResponse) GetBody() *PricingResponseBody {
	return s.Body
}

func (s *PricingResponse) SetHeaders(v map[string]*string) *PricingResponse {
	s.Headers = v
	return s
}

func (s *PricingResponse) SetStatusCode(v int32) *PricingResponse {
	s.StatusCode = &v
	return s
}

func (s *PricingResponse) SetBody(v *PricingResponseBody) *PricingResponse {
	s.Body = v
	return s
}

func (s *PricingResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
