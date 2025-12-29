// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHotelScreenSaverStyleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetHotelScreenSaverStyleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetHotelScreenSaverStyleResponse
	GetStatusCode() *int32
	SetBody(v *GetHotelScreenSaverStyleResponseBody) *GetHotelScreenSaverStyleResponse
	GetBody() *GetHotelScreenSaverStyleResponseBody
}

type GetHotelScreenSaverStyleResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetHotelScreenSaverStyleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetHotelScreenSaverStyleResponse) String() string {
	return dara.Prettify(s)
}

func (s GetHotelScreenSaverStyleResponse) GoString() string {
	return s.String()
}

func (s *GetHotelScreenSaverStyleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetHotelScreenSaverStyleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetHotelScreenSaverStyleResponse) GetBody() *GetHotelScreenSaverStyleResponseBody {
	return s.Body
}

func (s *GetHotelScreenSaverStyleResponse) SetHeaders(v map[string]*string) *GetHotelScreenSaverStyleResponse {
	s.Headers = v
	return s
}

func (s *GetHotelScreenSaverStyleResponse) SetStatusCode(v int32) *GetHotelScreenSaverStyleResponse {
	s.StatusCode = &v
	return s
}

func (s *GetHotelScreenSaverStyleResponse) SetBody(v *GetHotelScreenSaverStyleResponseBody) *GetHotelScreenSaverStyleResponse {
	s.Body = v
	return s
}

func (s *GetHotelScreenSaverStyleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
