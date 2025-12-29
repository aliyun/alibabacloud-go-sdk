// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHotelInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListHotelInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListHotelInfoResponse
	GetStatusCode() *int32
	SetBody(v *ListHotelInfoResponseBody) *ListHotelInfoResponse
	GetBody() *ListHotelInfoResponseBody
}

type ListHotelInfoResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListHotelInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListHotelInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s ListHotelInfoResponse) GoString() string {
	return s.String()
}

func (s *ListHotelInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListHotelInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListHotelInfoResponse) GetBody() *ListHotelInfoResponseBody {
	return s.Body
}

func (s *ListHotelInfoResponse) SetHeaders(v map[string]*string) *ListHotelInfoResponse {
	s.Headers = v
	return s
}

func (s *ListHotelInfoResponse) SetStatusCode(v int32) *ListHotelInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *ListHotelInfoResponse) SetBody(v *ListHotelInfoResponseBody) *ListHotelInfoResponse {
	s.Body = v
	return s
}

func (s *ListHotelInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
