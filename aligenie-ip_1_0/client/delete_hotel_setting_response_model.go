// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteHotelSettingResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteHotelSettingResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteHotelSettingResponse
	GetStatusCode() *int32
	SetBody(v *DeleteHotelSettingResponseBody) *DeleteHotelSettingResponse
	GetBody() *DeleteHotelSettingResponseBody
}

type DeleteHotelSettingResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteHotelSettingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteHotelSettingResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteHotelSettingResponse) GoString() string {
	return s.String()
}

func (s *DeleteHotelSettingResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteHotelSettingResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteHotelSettingResponse) GetBody() *DeleteHotelSettingResponseBody {
	return s.Body
}

func (s *DeleteHotelSettingResponse) SetHeaders(v map[string]*string) *DeleteHotelSettingResponse {
	s.Headers = v
	return s
}

func (s *DeleteHotelSettingResponse) SetStatusCode(v int32) *DeleteHotelSettingResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteHotelSettingResponse) SetBody(v *DeleteHotelSettingResponseBody) *DeleteHotelSettingResponse {
	s.Body = v
	return s
}

func (s *DeleteHotelSettingResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
