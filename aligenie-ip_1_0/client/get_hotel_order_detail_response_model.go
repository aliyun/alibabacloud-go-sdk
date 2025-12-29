// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHotelOrderDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetHotelOrderDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetHotelOrderDetailResponse
	GetStatusCode() *int32
	SetBody(v *GetHotelOrderDetailResponseBody) *GetHotelOrderDetailResponse
	GetBody() *GetHotelOrderDetailResponseBody
}

type GetHotelOrderDetailResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetHotelOrderDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetHotelOrderDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s GetHotelOrderDetailResponse) GoString() string {
	return s.String()
}

func (s *GetHotelOrderDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetHotelOrderDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetHotelOrderDetailResponse) GetBody() *GetHotelOrderDetailResponseBody {
	return s.Body
}

func (s *GetHotelOrderDetailResponse) SetHeaders(v map[string]*string) *GetHotelOrderDetailResponse {
	s.Headers = v
	return s
}

func (s *GetHotelOrderDetailResponse) SetStatusCode(v int32) *GetHotelOrderDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *GetHotelOrderDetailResponse) SetBody(v *GetHotelOrderDetailResponseBody) *GetHotelOrderDetailResponse {
	s.Body = v
	return s
}

func (s *GetHotelOrderDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
