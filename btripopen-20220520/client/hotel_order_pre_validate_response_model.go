// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHotelOrderPreValidateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *HotelOrderPreValidateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *HotelOrderPreValidateResponse
	GetStatusCode() *int32
	SetBody(v *HotelOrderPreValidateResponseBody) *HotelOrderPreValidateResponse
	GetBody() *HotelOrderPreValidateResponseBody
}

type HotelOrderPreValidateResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *HotelOrderPreValidateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s HotelOrderPreValidateResponse) String() string {
	return dara.Prettify(s)
}

func (s HotelOrderPreValidateResponse) GoString() string {
	return s.String()
}

func (s *HotelOrderPreValidateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *HotelOrderPreValidateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *HotelOrderPreValidateResponse) GetBody() *HotelOrderPreValidateResponseBody {
	return s.Body
}

func (s *HotelOrderPreValidateResponse) SetHeaders(v map[string]*string) *HotelOrderPreValidateResponse {
	s.Headers = v
	return s
}

func (s *HotelOrderPreValidateResponse) SetStatusCode(v int32) *HotelOrderPreValidateResponse {
	s.StatusCode = &v
	return s
}

func (s *HotelOrderPreValidateResponse) SetBody(v *HotelOrderPreValidateResponseBody) *HotelOrderPreValidateResponse {
	s.Body = v
	return s
}

func (s *HotelOrderPreValidateResponse) Validate() error {
	return dara.Validate(s)
}
