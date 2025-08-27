// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHotelOrderCreateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *HotelOrderCreateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *HotelOrderCreateResponse
	GetStatusCode() *int32
	SetBody(v *HotelOrderCreateResponseBody) *HotelOrderCreateResponse
	GetBody() *HotelOrderCreateResponseBody
}

type HotelOrderCreateResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *HotelOrderCreateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s HotelOrderCreateResponse) String() string {
	return dara.Prettify(s)
}

func (s HotelOrderCreateResponse) GoString() string {
	return s.String()
}

func (s *HotelOrderCreateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *HotelOrderCreateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *HotelOrderCreateResponse) GetBody() *HotelOrderCreateResponseBody {
	return s.Body
}

func (s *HotelOrderCreateResponse) SetHeaders(v map[string]*string) *HotelOrderCreateResponse {
	s.Headers = v
	return s
}

func (s *HotelOrderCreateResponse) SetStatusCode(v int32) *HotelOrderCreateResponse {
	s.StatusCode = &v
	return s
}

func (s *HotelOrderCreateResponse) SetBody(v *HotelOrderCreateResponseBody) *HotelOrderCreateResponse {
	s.Body = v
	return s
}

func (s *HotelOrderCreateResponse) Validate() error {
	return dara.Validate(s)
}
