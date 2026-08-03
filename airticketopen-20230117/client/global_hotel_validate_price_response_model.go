// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGlobalHotelValidatePriceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GlobalHotelValidatePriceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GlobalHotelValidatePriceResponse
	GetStatusCode() *int32
	SetBody(v *GlobalHotelValidatePriceResponseBody) *GlobalHotelValidatePriceResponse
	GetBody() *GlobalHotelValidatePriceResponseBody
}

type GlobalHotelValidatePriceResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GlobalHotelValidatePriceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GlobalHotelValidatePriceResponse) String() string {
	return dara.Prettify(s)
}

func (s GlobalHotelValidatePriceResponse) GoString() string {
	return s.String()
}

func (s *GlobalHotelValidatePriceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GlobalHotelValidatePriceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GlobalHotelValidatePriceResponse) GetBody() *GlobalHotelValidatePriceResponseBody {
	return s.Body
}

func (s *GlobalHotelValidatePriceResponse) SetHeaders(v map[string]*string) *GlobalHotelValidatePriceResponse {
	s.Headers = v
	return s
}

func (s *GlobalHotelValidatePriceResponse) SetStatusCode(v int32) *GlobalHotelValidatePriceResponse {
	s.StatusCode = &v
	return s
}

func (s *GlobalHotelValidatePriceResponse) SetBody(v *GlobalHotelValidatePriceResponseBody) *GlobalHotelValidatePriceResponse {
	s.Body = v
	return s
}

func (s *GlobalHotelValidatePriceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
