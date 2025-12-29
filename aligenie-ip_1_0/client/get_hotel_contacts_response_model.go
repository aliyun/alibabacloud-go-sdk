// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHotelContactsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetHotelContactsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetHotelContactsResponse
	GetStatusCode() *int32
	SetBody(v *GetHotelContactsResponseBody) *GetHotelContactsResponse
	GetBody() *GetHotelContactsResponseBody
}

type GetHotelContactsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetHotelContactsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetHotelContactsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetHotelContactsResponse) GoString() string {
	return s.String()
}

func (s *GetHotelContactsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetHotelContactsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetHotelContactsResponse) GetBody() *GetHotelContactsResponseBody {
	return s.Body
}

func (s *GetHotelContactsResponse) SetHeaders(v map[string]*string) *GetHotelContactsResponse {
	s.Headers = v
	return s
}

func (s *GetHotelContactsResponse) SetStatusCode(v int32) *GetHotelContactsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetHotelContactsResponse) SetBody(v *GetHotelContactsResponseBody) *GetHotelContactsResponse {
	s.Body = v
	return s
}

func (s *GetHotelContactsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
