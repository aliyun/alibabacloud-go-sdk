// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPushHotelMessageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PushHotelMessageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PushHotelMessageResponse
	GetStatusCode() *int32
	SetBody(v *PushHotelMessageResponseBody) *PushHotelMessageResponse
	GetBody() *PushHotelMessageResponseBody
}

type PushHotelMessageResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PushHotelMessageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PushHotelMessageResponse) String() string {
	return dara.Prettify(s)
}

func (s PushHotelMessageResponse) GoString() string {
	return s.String()
}

func (s *PushHotelMessageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PushHotelMessageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PushHotelMessageResponse) GetBody() *PushHotelMessageResponseBody {
	return s.Body
}

func (s *PushHotelMessageResponse) SetHeaders(v map[string]*string) *PushHotelMessageResponse {
	s.Headers = v
	return s
}

func (s *PushHotelMessageResponse) SetStatusCode(v int32) *PushHotelMessageResponse {
	s.StatusCode = &v
	return s
}

func (s *PushHotelMessageResponse) SetBody(v *PushHotelMessageResponseBody) *PushHotelMessageResponse {
	s.Body = v
	return s
}

func (s *PushHotelMessageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
