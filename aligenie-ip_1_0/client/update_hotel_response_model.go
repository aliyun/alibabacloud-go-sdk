// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateHotelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateHotelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateHotelResponse
	GetStatusCode() *int32
	SetBody(v *UpdateHotelResponseBody) *UpdateHotelResponse
	GetBody() *UpdateHotelResponseBody
}

type UpdateHotelResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateHotelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateHotelResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateHotelResponse) GoString() string {
	return s.String()
}

func (s *UpdateHotelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateHotelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateHotelResponse) GetBody() *UpdateHotelResponseBody {
	return s.Body
}

func (s *UpdateHotelResponse) SetHeaders(v map[string]*string) *UpdateHotelResponse {
	s.Headers = v
	return s
}

func (s *UpdateHotelResponse) SetStatusCode(v int32) *UpdateHotelResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateHotelResponse) SetBody(v *UpdateHotelResponseBody) *UpdateHotelResponse {
	s.Body = v
	return s
}

func (s *UpdateHotelResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
