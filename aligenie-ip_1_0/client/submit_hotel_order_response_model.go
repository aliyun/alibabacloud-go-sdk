// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitHotelOrderResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitHotelOrderResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitHotelOrderResponse
	GetStatusCode() *int32
	SetBody(v *SubmitHotelOrderResponseBody) *SubmitHotelOrderResponse
	GetBody() *SubmitHotelOrderResponseBody
}

type SubmitHotelOrderResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitHotelOrderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitHotelOrderResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitHotelOrderResponse) GoString() string {
	return s.String()
}

func (s *SubmitHotelOrderResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitHotelOrderResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitHotelOrderResponse) GetBody() *SubmitHotelOrderResponseBody {
	return s.Body
}

func (s *SubmitHotelOrderResponse) SetHeaders(v map[string]*string) *SubmitHotelOrderResponse {
	s.Headers = v
	return s
}

func (s *SubmitHotelOrderResponse) SetStatusCode(v int32) *SubmitHotelOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitHotelOrderResponse) SetBody(v *SubmitHotelOrderResponseBody) *SubmitHotelOrderResponse {
	s.Body = v
	return s
}

func (s *SubmitHotelOrderResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
