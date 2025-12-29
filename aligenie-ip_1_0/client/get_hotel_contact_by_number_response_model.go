// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHotelContactByNumberResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetHotelContactByNumberResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetHotelContactByNumberResponse
	GetStatusCode() *int32
	SetBody(v *GetHotelContactByNumberResponseBody) *GetHotelContactByNumberResponse
	GetBody() *GetHotelContactByNumberResponseBody
}

type GetHotelContactByNumberResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetHotelContactByNumberResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetHotelContactByNumberResponse) String() string {
	return dara.Prettify(s)
}

func (s GetHotelContactByNumberResponse) GoString() string {
	return s.String()
}

func (s *GetHotelContactByNumberResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetHotelContactByNumberResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetHotelContactByNumberResponse) GetBody() *GetHotelContactByNumberResponseBody {
	return s.Body
}

func (s *GetHotelContactByNumberResponse) SetHeaders(v map[string]*string) *GetHotelContactByNumberResponse {
	s.Headers = v
	return s
}

func (s *GetHotelContactByNumberResponse) SetStatusCode(v int32) *GetHotelContactByNumberResponse {
	s.StatusCode = &v
	return s
}

func (s *GetHotelContactByNumberResponse) SetBody(v *GetHotelContactByNumberResponseBody) *GetHotelContactByNumberResponse {
	s.Body = v
	return s
}

func (s *GetHotelContactByNumberResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
