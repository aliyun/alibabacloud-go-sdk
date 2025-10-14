// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOrderDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OrderDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OrderDetailResponse
	GetStatusCode() *int32
	SetBody(v *OrderDetailResponseBody) *OrderDetailResponse
	GetBody() *OrderDetailResponseBody
}

type OrderDetailResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OrderDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OrderDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s OrderDetailResponse) GoString() string {
	return s.String()
}

func (s *OrderDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OrderDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OrderDetailResponse) GetBody() *OrderDetailResponseBody {
	return s.Body
}

func (s *OrderDetailResponse) SetHeaders(v map[string]*string) *OrderDetailResponse {
	s.Headers = v
	return s
}

func (s *OrderDetailResponse) SetStatusCode(v int32) *OrderDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *OrderDetailResponse) SetBody(v *OrderDetailResponseBody) *OrderDetailResponse {
	s.Body = v
	return s
}

func (s *OrderDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
