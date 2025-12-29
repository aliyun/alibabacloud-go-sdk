// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHotelScreenSaverResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetHotelScreenSaverResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetHotelScreenSaverResponse
	GetStatusCode() *int32
	SetBody(v *GetHotelScreenSaverResponseBody) *GetHotelScreenSaverResponse
	GetBody() *GetHotelScreenSaverResponseBody
}

type GetHotelScreenSaverResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetHotelScreenSaverResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetHotelScreenSaverResponse) String() string {
	return dara.Prettify(s)
}

func (s GetHotelScreenSaverResponse) GoString() string {
	return s.String()
}

func (s *GetHotelScreenSaverResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetHotelScreenSaverResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetHotelScreenSaverResponse) GetBody() *GetHotelScreenSaverResponseBody {
	return s.Body
}

func (s *GetHotelScreenSaverResponse) SetHeaders(v map[string]*string) *GetHotelScreenSaverResponse {
	s.Headers = v
	return s
}

func (s *GetHotelScreenSaverResponse) SetStatusCode(v int32) *GetHotelScreenSaverResponse {
	s.StatusCode = &v
	return s
}

func (s *GetHotelScreenSaverResponse) SetBody(v *GetHotelScreenSaverResponseBody) *GetHotelScreenSaverResponse {
	s.Body = v
	return s
}

func (s *GetHotelScreenSaverResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
