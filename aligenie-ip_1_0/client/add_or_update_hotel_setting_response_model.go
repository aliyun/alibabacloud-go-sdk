// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddOrUpdateHotelSettingResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddOrUpdateHotelSettingResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddOrUpdateHotelSettingResponse
	GetStatusCode() *int32
	SetBody(v *AddOrUpdateHotelSettingResponseBody) *AddOrUpdateHotelSettingResponse
	GetBody() *AddOrUpdateHotelSettingResponseBody
}

type AddOrUpdateHotelSettingResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddOrUpdateHotelSettingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddOrUpdateHotelSettingResponse) String() string {
	return dara.Prettify(s)
}

func (s AddOrUpdateHotelSettingResponse) GoString() string {
	return s.String()
}

func (s *AddOrUpdateHotelSettingResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddOrUpdateHotelSettingResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddOrUpdateHotelSettingResponse) GetBody() *AddOrUpdateHotelSettingResponseBody {
	return s.Body
}

func (s *AddOrUpdateHotelSettingResponse) SetHeaders(v map[string]*string) *AddOrUpdateHotelSettingResponse {
	s.Headers = v
	return s
}

func (s *AddOrUpdateHotelSettingResponse) SetStatusCode(v int32) *AddOrUpdateHotelSettingResponse {
	s.StatusCode = &v
	return s
}

func (s *AddOrUpdateHotelSettingResponse) SetBody(v *AddOrUpdateHotelSettingResponseBody) *AddOrUpdateHotelSettingResponse {
	s.Body = v
	return s
}

func (s *AddOrUpdateHotelSettingResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
