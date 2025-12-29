// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHotelSettingResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetHotelSettingResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetHotelSettingResponse
	GetStatusCode() *int32
	SetBody(v *GetHotelSettingResponseBody) *GetHotelSettingResponse
	GetBody() *GetHotelSettingResponseBody
}

type GetHotelSettingResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetHotelSettingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetHotelSettingResponse) String() string {
	return dara.Prettify(s)
}

func (s GetHotelSettingResponse) GoString() string {
	return s.String()
}

func (s *GetHotelSettingResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetHotelSettingResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetHotelSettingResponse) GetBody() *GetHotelSettingResponseBody {
	return s.Body
}

func (s *GetHotelSettingResponse) SetHeaders(v map[string]*string) *GetHotelSettingResponse {
	s.Headers = v
	return s
}

func (s *GetHotelSettingResponse) SetStatusCode(v int32) *GetHotelSettingResponse {
	s.StatusCode = &v
	return s
}

func (s *GetHotelSettingResponse) SetBody(v *GetHotelSettingResponseBody) *GetHotelSettingResponse {
	s.Body = v
	return s
}

func (s *GetHotelSettingResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
