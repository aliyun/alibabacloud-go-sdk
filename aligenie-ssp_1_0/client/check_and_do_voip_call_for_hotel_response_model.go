// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckAndDoVoipCallForHotelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CheckAndDoVoipCallForHotelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CheckAndDoVoipCallForHotelResponse
	GetStatusCode() *int32
	SetBody(v *CheckAndDoVoipCallForHotelResponseBody) *CheckAndDoVoipCallForHotelResponse
	GetBody() *CheckAndDoVoipCallForHotelResponseBody
}

type CheckAndDoVoipCallForHotelResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckAndDoVoipCallForHotelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckAndDoVoipCallForHotelResponse) String() string {
	return dara.Prettify(s)
}

func (s CheckAndDoVoipCallForHotelResponse) GoString() string {
	return s.String()
}

func (s *CheckAndDoVoipCallForHotelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CheckAndDoVoipCallForHotelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CheckAndDoVoipCallForHotelResponse) GetBody() *CheckAndDoVoipCallForHotelResponseBody {
	return s.Body
}

func (s *CheckAndDoVoipCallForHotelResponse) SetHeaders(v map[string]*string) *CheckAndDoVoipCallForHotelResponse {
	s.Headers = v
	return s
}

func (s *CheckAndDoVoipCallForHotelResponse) SetStatusCode(v int32) *CheckAndDoVoipCallForHotelResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckAndDoVoipCallForHotelResponse) SetBody(v *CheckAndDoVoipCallForHotelResponseBody) *CheckAndDoVoipCallForHotelResponse {
	s.Body = v
	return s
}

func (s *CheckAndDoVoipCallForHotelResponse) Validate() error {
	return dara.Validate(s)
}
