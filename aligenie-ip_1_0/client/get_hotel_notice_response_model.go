// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHotelNoticeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetHotelNoticeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetHotelNoticeResponse
	GetStatusCode() *int32
	SetBody(v *GetHotelNoticeResponseBody) *GetHotelNoticeResponse
	GetBody() *GetHotelNoticeResponseBody
}

type GetHotelNoticeResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetHotelNoticeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetHotelNoticeResponse) String() string {
	return dara.Prettify(s)
}

func (s GetHotelNoticeResponse) GoString() string {
	return s.String()
}

func (s *GetHotelNoticeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetHotelNoticeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetHotelNoticeResponse) GetBody() *GetHotelNoticeResponseBody {
	return s.Body
}

func (s *GetHotelNoticeResponse) SetHeaders(v map[string]*string) *GetHotelNoticeResponse {
	s.Headers = v
	return s
}

func (s *GetHotelNoticeResponse) SetStatusCode(v int32) *GetHotelNoticeResponse {
	s.StatusCode = &v
	return s
}

func (s *GetHotelNoticeResponse) SetBody(v *GetHotelNoticeResponseBody) *GetHotelNoticeResponse {
	s.Body = v
	return s
}

func (s *GetHotelNoticeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
