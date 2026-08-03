// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGlobalHotelCreateAndPayResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GlobalHotelCreateAndPayResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GlobalHotelCreateAndPayResponse
	GetStatusCode() *int32
	SetBody(v *GlobalHotelCreateAndPayResponseBody) *GlobalHotelCreateAndPayResponse
	GetBody() *GlobalHotelCreateAndPayResponseBody
}

type GlobalHotelCreateAndPayResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GlobalHotelCreateAndPayResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GlobalHotelCreateAndPayResponse) String() string {
	return dara.Prettify(s)
}

func (s GlobalHotelCreateAndPayResponse) GoString() string {
	return s.String()
}

func (s *GlobalHotelCreateAndPayResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GlobalHotelCreateAndPayResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GlobalHotelCreateAndPayResponse) GetBody() *GlobalHotelCreateAndPayResponseBody {
	return s.Body
}

func (s *GlobalHotelCreateAndPayResponse) SetHeaders(v map[string]*string) *GlobalHotelCreateAndPayResponse {
	s.Headers = v
	return s
}

func (s *GlobalHotelCreateAndPayResponse) SetStatusCode(v int32) *GlobalHotelCreateAndPayResponse {
	s.StatusCode = &v
	return s
}

func (s *GlobalHotelCreateAndPayResponse) SetBody(v *GlobalHotelCreateAndPayResponseBody) *GlobalHotelCreateAndPayResponse {
	s.Body = v
	return s
}

func (s *GlobalHotelCreateAndPayResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
