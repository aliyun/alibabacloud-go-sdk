// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchGetHotelDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BatchGetHotelDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BatchGetHotelDetailResponse
	GetStatusCode() *int32
	SetBody(v *BatchGetHotelDetailResponseBody) *BatchGetHotelDetailResponse
	GetBody() *BatchGetHotelDetailResponseBody
}

type BatchGetHotelDetailResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchGetHotelDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchGetHotelDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s BatchGetHotelDetailResponse) GoString() string {
	return s.String()
}

func (s *BatchGetHotelDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BatchGetHotelDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BatchGetHotelDetailResponse) GetBody() *BatchGetHotelDetailResponseBody {
	return s.Body
}

func (s *BatchGetHotelDetailResponse) SetHeaders(v map[string]*string) *BatchGetHotelDetailResponse {
	s.Headers = v
	return s
}

func (s *BatchGetHotelDetailResponse) SetStatusCode(v int32) *BatchGetHotelDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchGetHotelDetailResponse) SetBody(v *BatchGetHotelDetailResponseBody) *BatchGetHotelDetailResponse {
	s.Body = v
	return s
}

func (s *BatchGetHotelDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
