// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelRefundOrderResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CancelRefundOrderResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CancelRefundOrderResponse
	GetStatusCode() *int32
	SetBody(v *RefundOrderResult) *CancelRefundOrderResponse
	GetBody() *RefundOrderResult
}

type CancelRefundOrderResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RefundOrderResult `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CancelRefundOrderResponse) String() string {
	return dara.Prettify(s)
}

func (s CancelRefundOrderResponse) GoString() string {
	return s.String()
}

func (s *CancelRefundOrderResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CancelRefundOrderResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CancelRefundOrderResponse) GetBody() *RefundOrderResult {
	return s.Body
}

func (s *CancelRefundOrderResponse) SetHeaders(v map[string]*string) *CancelRefundOrderResponse {
	s.Headers = v
	return s
}

func (s *CancelRefundOrderResponse) SetStatusCode(v int32) *CancelRefundOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelRefundOrderResponse) SetBody(v *RefundOrderResult) *CancelRefundOrderResponse {
	s.Body = v
	return s
}

func (s *CancelRefundOrderResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
