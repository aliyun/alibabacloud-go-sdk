// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGlobalHotelCancelOrRefundResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GlobalHotelCancelOrRefundResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GlobalHotelCancelOrRefundResponse
	GetStatusCode() *int32
	SetBody(v *GlobalHotelCancelOrRefundResponseBody) *GlobalHotelCancelOrRefundResponse
	GetBody() *GlobalHotelCancelOrRefundResponseBody
}

type GlobalHotelCancelOrRefundResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GlobalHotelCancelOrRefundResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GlobalHotelCancelOrRefundResponse) String() string {
	return dara.Prettify(s)
}

func (s GlobalHotelCancelOrRefundResponse) GoString() string {
	return s.String()
}

func (s *GlobalHotelCancelOrRefundResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GlobalHotelCancelOrRefundResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GlobalHotelCancelOrRefundResponse) GetBody() *GlobalHotelCancelOrRefundResponseBody {
	return s.Body
}

func (s *GlobalHotelCancelOrRefundResponse) SetHeaders(v map[string]*string) *GlobalHotelCancelOrRefundResponse {
	s.Headers = v
	return s
}

func (s *GlobalHotelCancelOrRefundResponse) SetStatusCode(v int32) *GlobalHotelCancelOrRefundResponse {
	s.StatusCode = &v
	return s
}

func (s *GlobalHotelCancelOrRefundResponse) SetBody(v *GlobalHotelCancelOrRefundResponseBody) *GlobalHotelCancelOrRefundResponse {
	s.Body = v
	return s
}

func (s *GlobalHotelCancelOrRefundResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
