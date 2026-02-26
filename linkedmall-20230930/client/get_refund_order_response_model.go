// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRefundOrderResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetRefundOrderResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetRefundOrderResponse
	GetStatusCode() *int32
	SetBody(v *RefundResult) *GetRefundOrderResponse
	GetBody() *RefundResult
}

type GetRefundOrderResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RefundResult      `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRefundOrderResponse) String() string {
	return dara.Prettify(s)
}

func (s GetRefundOrderResponse) GoString() string {
	return s.String()
}

func (s *GetRefundOrderResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetRefundOrderResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetRefundOrderResponse) GetBody() *RefundResult {
	return s.Body
}

func (s *GetRefundOrderResponse) SetHeaders(v map[string]*string) *GetRefundOrderResponse {
	s.Headers = v
	return s
}

func (s *GetRefundOrderResponse) SetStatusCode(v int32) *GetRefundOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRefundOrderResponse) SetBody(v *RefundResult) *GetRefundOrderResponse {
	s.Body = v
	return s
}

func (s *GetRefundOrderResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
