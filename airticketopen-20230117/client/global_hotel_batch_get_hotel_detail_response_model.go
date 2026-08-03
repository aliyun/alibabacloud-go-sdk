// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGlobalHotelBatchGetHotelDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GlobalHotelBatchGetHotelDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GlobalHotelBatchGetHotelDetailResponse
	GetStatusCode() *int32
	SetBody(v *GlobalHotelBatchGetHotelDetailResponseBody) *GlobalHotelBatchGetHotelDetailResponse
	GetBody() *GlobalHotelBatchGetHotelDetailResponseBody
}

type GlobalHotelBatchGetHotelDetailResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GlobalHotelBatchGetHotelDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GlobalHotelBatchGetHotelDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s GlobalHotelBatchGetHotelDetailResponse) GoString() string {
	return s.String()
}

func (s *GlobalHotelBatchGetHotelDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GlobalHotelBatchGetHotelDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GlobalHotelBatchGetHotelDetailResponse) GetBody() *GlobalHotelBatchGetHotelDetailResponseBody {
	return s.Body
}

func (s *GlobalHotelBatchGetHotelDetailResponse) SetHeaders(v map[string]*string) *GlobalHotelBatchGetHotelDetailResponse {
	s.Headers = v
	return s
}

func (s *GlobalHotelBatchGetHotelDetailResponse) SetStatusCode(v int32) *GlobalHotelBatchGetHotelDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *GlobalHotelBatchGetHotelDetailResponse) SetBody(v *GlobalHotelBatchGetHotelDetailResponseBody) *GlobalHotelBatchGetHotelDetailResponse {
	s.Body = v
	return s
}

func (s *GlobalHotelBatchGetHotelDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
