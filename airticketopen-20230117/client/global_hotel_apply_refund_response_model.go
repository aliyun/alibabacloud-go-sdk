// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGlobalHotelApplyRefundResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GlobalHotelApplyRefundResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GlobalHotelApplyRefundResponse
	GetStatusCode() *int32
	SetBody(v *GlobalHotelApplyRefundResponseBody) *GlobalHotelApplyRefundResponse
	GetBody() *GlobalHotelApplyRefundResponseBody
}

type GlobalHotelApplyRefundResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GlobalHotelApplyRefundResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GlobalHotelApplyRefundResponse) String() string {
	return dara.Prettify(s)
}

func (s GlobalHotelApplyRefundResponse) GoString() string {
	return s.String()
}

func (s *GlobalHotelApplyRefundResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GlobalHotelApplyRefundResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GlobalHotelApplyRefundResponse) GetBody() *GlobalHotelApplyRefundResponseBody {
	return s.Body
}

func (s *GlobalHotelApplyRefundResponse) SetHeaders(v map[string]*string) *GlobalHotelApplyRefundResponse {
	s.Headers = v
	return s
}

func (s *GlobalHotelApplyRefundResponse) SetStatusCode(v int32) *GlobalHotelApplyRefundResponse {
	s.StatusCode = &v
	return s
}

func (s *GlobalHotelApplyRefundResponse) SetBody(v *GlobalHotelApplyRefundResponseBody) *GlobalHotelApplyRefundResponse {
	s.Body = v
	return s
}

func (s *GlobalHotelApplyRefundResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
