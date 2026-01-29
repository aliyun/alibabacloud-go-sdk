// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOrderDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetOrderDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetOrderDetailResponse
	GetStatusCode() *int32
	SetBody(v *GetOrderDetailResponseBody) *GetOrderDetailResponse
	GetBody() *GetOrderDetailResponseBody
}

type GetOrderDetailResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetOrderDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetOrderDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s GetOrderDetailResponse) GoString() string {
	return s.String()
}

func (s *GetOrderDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetOrderDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetOrderDetailResponse) GetBody() *GetOrderDetailResponseBody {
	return s.Body
}

func (s *GetOrderDetailResponse) SetHeaders(v map[string]*string) *GetOrderDetailResponse {
	s.Headers = v
	return s
}

func (s *GetOrderDetailResponse) SetStatusCode(v int32) *GetOrderDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOrderDetailResponse) SetBody(v *GetOrderDetailResponseBody) *GetOrderDetailResponse {
	s.Body = v
	return s
}

func (s *GetOrderDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
