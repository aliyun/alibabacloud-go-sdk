// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVnoBatchRefundOrderResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *VnoBatchRefundOrderResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *VnoBatchRefundOrderResponse
	GetStatusCode() *int32
	SetBody(v *VnoBatchRefundOrderResponseBody) *VnoBatchRefundOrderResponse
	GetBody() *VnoBatchRefundOrderResponseBody
}

type VnoBatchRefundOrderResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *VnoBatchRefundOrderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s VnoBatchRefundOrderResponse) String() string {
	return dara.Prettify(s)
}

func (s VnoBatchRefundOrderResponse) GoString() string {
	return s.String()
}

func (s *VnoBatchRefundOrderResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *VnoBatchRefundOrderResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *VnoBatchRefundOrderResponse) GetBody() *VnoBatchRefundOrderResponseBody {
	return s.Body
}

func (s *VnoBatchRefundOrderResponse) SetHeaders(v map[string]*string) *VnoBatchRefundOrderResponse {
	s.Headers = v
	return s
}

func (s *VnoBatchRefundOrderResponse) SetStatusCode(v int32) *VnoBatchRefundOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *VnoBatchRefundOrderResponse) SetBody(v *VnoBatchRefundOrderResponseBody) *VnoBatchRefundOrderResponse {
	s.Body = v
	return s
}

func (s *VnoBatchRefundOrderResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
