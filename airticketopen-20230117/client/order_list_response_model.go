// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOrderListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OrderListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OrderListResponse
	GetStatusCode() *int32
	SetBody(v *OrderListResponseBody) *OrderListResponse
	GetBody() *OrderListResponseBody
}

type OrderListResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OrderListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OrderListResponse) String() string {
	return dara.Prettify(s)
}

func (s OrderListResponse) GoString() string {
	return s.String()
}

func (s *OrderListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OrderListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OrderListResponse) GetBody() *OrderListResponseBody {
	return s.Body
}

func (s *OrderListResponse) SetHeaders(v map[string]*string) *OrderListResponse {
	s.Headers = v
	return s
}

func (s *OrderListResponse) SetStatusCode(v int32) *OrderListResponse {
	s.StatusCode = &v
	return s
}

func (s *OrderListResponse) SetBody(v *OrderListResponseBody) *OrderListResponse {
	s.Body = v
	return s
}

func (s *OrderListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
