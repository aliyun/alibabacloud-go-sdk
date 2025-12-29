// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateHotelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateHotelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateHotelResponse
	GetStatusCode() *int32
	SetBody(v *CreateHotelResponseBody) *CreateHotelResponse
	GetBody() *CreateHotelResponseBody
}

type CreateHotelResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateHotelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateHotelResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateHotelResponse) GoString() string {
	return s.String()
}

func (s *CreateHotelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateHotelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateHotelResponse) GetBody() *CreateHotelResponseBody {
	return s.Body
}

func (s *CreateHotelResponse) SetHeaders(v map[string]*string) *CreateHotelResponse {
	s.Headers = v
	return s
}

func (s *CreateHotelResponse) SetStatusCode(v int32) *CreateHotelResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateHotelResponse) SetBody(v *CreateHotelResponseBody) *CreateHotelResponse {
	s.Body = v
	return s
}

func (s *CreateHotelResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
